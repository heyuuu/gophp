package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/sapi/cli"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
)

func BROWSCAP_G(v __auto__) __auto__      { return BrowscapGlobals.v }
func IsPlaceholder(c byte) types.ZendBool { return c == '?' || c == '*' }
func BrowscapComputePrefixLen(pattern *types.String) uint8 {
	var i int
	for i = 0; i < pattern.GetLen(); i++ {
		if IsPlaceholder(pattern.GetVal()[i]) != 0 {
			break
		}
	}
	return uint8(cli.MIN(i, UINT8_MAX))
}
func BrowscapComputeContains(pattern *types.String, start_pos int, contains_start *uint16, contains_len *uint8) int {
	var i int = start_pos

	/* Find first non-placeholder character after prefix */

	for ; i < pattern.GetLen(); i++ {
		if IsPlaceholder(pattern.GetVal()[i]) == 0 {

			/* Skip the case of a single non-placeholder character.
			 * Let's try to find something longer instead. */

			if i+1 < pattern.GetLen() && IsPlaceholder(pattern.GetVal()[i+1]) == 0 {
				break
			}

			/* Skip the case of a single non-placeholder character.
			 * Let's try to find something longer instead. */

		}
	}
	*contains_start = uint16(i)

	/* Find first placeholder character after that */

	for ; i < pattern.GetLen(); i++ {
		if IsPlaceholder(pattern.GetVal()[i]) != 0 {
			break
		}
	}
	*contains_len = uint8(cli.MIN(i-(*contains_start), UINT8_MAX))
	return i
}
func BrowscapComputeRegexLen(pattern *types.String) int {
	var i int
	var len_ int = pattern.GetLen()
	for i = 0; i < pattern.GetLen(); i++ {
		switch pattern.GetVal()[i] {
		case '*':
			fallthrough
		case '.':
			fallthrough
		case '\\':
			fallthrough
		case '(':
			fallthrough
		case ')':
			fallthrough
		case '~':
			fallthrough
		case '+':
			len_++
		}
	}
	return len_ + b.SizeOf("\"~^$~\"") - 1
}
func BrowscapConvertPattern(pattern *types.String, persistent int) *types.String {
	var i int
	var j int = 0
	var t *byte
	var res *types.String
	var lc_pattern *byte
	res = types.ZendStringAlloc(BrowscapComputeRegexLen(pattern), persistent)
	t = res.GetVal()
	lc_pattern = zend.DoAlloca(pattern.GetLen()+1, use_heap)
	zend.ZendStrTolowerCopy(lc_pattern, pattern.GetVal(), pattern.GetLen())
	t[b.PostInc(&j)] = '~'
	t[b.PostInc(&j)] = '^'
	for i = 0; i < pattern.GetLen(); {
		switch lc_pattern[i] {
		case '?':
			t[j] = '.'
		case '*':
			t[b.PostInc(&j)] = '.'
			t[j] = '*'
		case '.':
			t[b.PostInc(&j)] = '\\'
			t[j] = '.'
		case '\\':
			t[b.PostInc(&j)] = '\\'
			t[j] = '\\'
		case '(':
			t[b.PostInc(&j)] = '\\'
			t[j] = '('
		case ')':
			t[b.PostInc(&j)] = '\\'
			t[j] = ')'
		case '~':
			t[b.PostInc(&j)] = '\\'
			t[j] = '~'
		case '+':
			t[b.PostInc(&j)] = '\\'
			t[j] = '+'
		default:
			t[j] = lc_pattern[i]
		}
		i++
		j++
	}
	t[b.PostInc(&j)] = '$'
	t[b.PostInc(&j)] = '~'
	t[j] = 0
	res.SetLen(j)
	zend.FreeAlloca(lc_pattern, use_heap)
	return res
}
func BrowscapInternStr(ctx *BrowscapParserCtx, str *types.String) *types.String {
	return ctx.GetInternedStr(str.GetStr())
}
func BrowscapInternStrCi(ctx *BrowscapParserCtx, str *types.String) *types.String {
	lcName := ascii.StrToLower(str.GetStr())
	return ctx.GetInternedStr(lcName)
}
func BrowscapAddKv(bdata *BrowserData, key *types.String, value *types.String, persistent types.ZendBool) {
	if bdata.GetKvUsed() == bdata.GetKvSize() {
		bdata.SetKvSize(bdata.GetKvSize() * 2)
		bdata.SetKv(zend.SafePerealloc(bdata.GetKv(), bdata.GetKvSize()))
	}
	bdata.GetKv()[bdata.GetKvUsed()].SetKey(key)
	bdata.GetKv()[bdata.GetKvUsed()].SetValue(value)
	bdata.GetKvUsed()++
}
func BrowscapEntryToArray(bdata *BrowserData, entry *BrowscapEntry) *types.Array {
	var tmp types.Zval
	var i uint32
	var ht *types.Array = types.NewArray(8)
	tmp.SetString(BrowscapConvertPattern(entry.GetPattern(), 0))
	ht.KeyAdd("browser_name_regex", &tmp)
	tmp.SetStringCopy(entry.GetPattern())
	ht.KeyAdd("browser_name_pattern", &tmp)
	if entry.GetParent() != nil {
		tmp.SetStringCopy(entry.GetParent())
		ht.KeyAdd("parent", &tmp)
	}
	for i = entry.GetKvStart(); i < entry.GetKvEnd(); i++ {
		tmp.SetStringCopy(bdata.GetKv()[i].GetValue())
		ht.KeyAdd(bdata.GetKv()[i].GetKey().GetStr(), &tmp)
	}
	return ht
}
func PhpBrowscapParserCb(arg1 *types.Zval, arg2 *types.Zval, arg3 *types.Zval, callback_type int, arg any) {
	var ctx *BrowscapParserCtx = arg
	var bdata *BrowserData = ctx.GetBdata()
	var persistent int = bdata.GetHtab().GetGcFlags() & types.IS_ARRAY_PERSISTENT
	if arg1 == nil {
		return
	}
	switch callback_type {
	case zend.ZEND_INI_PARSER_ENTRY:
		if ctx.GetCurrentEntry() != nil && arg2 != nil {
			var new_key *types.String
			var new_value *types.String

			/* Set proper value for true/false settings */

			if arg2.String().GetLen() == 2 && !(strncasecmp(arg2.String().GetVal(), "on", b.SizeOf("\"on\"")-1)) || arg2.String().GetLen() == 3 && !(strncasecmp(arg2.String().GetVal(), "yes", b.SizeOf("\"yes\"")-1)) || arg2.String().GetLen() == 4 && !(strncasecmp(arg2.String().GetVal(), "true", b.SizeOf("\"true\"")-1)) {
				new_value = types.NewString("1")
			} else if arg2.String().GetLen() == 2 && !(strncasecmp(arg2.String().GetVal(), "no", b.SizeOf("\"no\"")-1)) || arg2.String().GetLen() == 3 && !(strncasecmp(arg2.String().GetVal(), "off", b.SizeOf("\"off\"")-1)) || arg2.String().GetLen() == 4 && !(strncasecmp(arg2.String().GetVal(), "none", b.SizeOf("\"none\"")-1)) || arg2.String().GetLen() == 5 && !(strncasecmp(arg2.String().GetVal(), "false", b.SizeOf("\"false\"")-1)) {
				new_value = types.NewString("")
			} else {
				new_value = BrowscapInternStr(ctx, arg2.String())
			}
			if !(strcasecmp(arg1.String().GetVal(), "parent")) {

				/* parent entry can not be same as current section -> causes infinite loop! */

				if ctx.GetCurrentSectionName() != nil && !(strcasecmp(ctx.GetCurrentSectionName().GetVal(), arg2.String().GetVal())) {
					faults.Error(faults.E_CORE_ERROR, "Invalid browscap ini file: "+"'Parent' value cannot be same as the section name: %s "+"(in file %s)", ctx.GetCurrentSectionName().GetVal(), zend.INI_STR("browscap"))
					return
				}
				if ctx.GetCurrentEntry().GetParent() != nil {
					// types.ZendStringRelease(ctx.GetCurrentEntry().GetParent())
				}
				ctx.GetCurrentEntry().SetParent(new_value)
			} else {
				new_key = BrowscapInternStrCi(ctx, arg1.String())
				BrowscapAddKv(bdata, new_key, new_value, persistent)
				ctx.GetCurrentEntry().SetKvEnd(bdata.GetKvUsed())
			}
		}
	case zend.ZEND_INI_PARSER_SECTION:
		var entry *BrowscapEntry
		var pattern *types.String = arg1.String()
		var pos int
		var i int
		if pattern.GetLen() > UINT16_MAX {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Skipping excessively long pattern of length %zd", pattern.GetLen())
			break
		}
		//if persistent != 0 {
		//	pattern = types.ZendNewInternedString(pattern.Copy())
		//	// types.ZendStringRelease(pattern)
		//}
		ctx.SetCurrentEntry(zend.Pemalloc(b.SizeOf("browscap_entry")))
		entry = ctx.GetCurrentEntry()
		types.ZendHashUpdatePtr(bdata.GetHtab(), pattern.GetStr(), entry)
		if ctx.GetCurrentSectionName() != nil {
			// types.ZendStringRelease(ctx.GetCurrentSectionName())
		}
		ctx.SetCurrentSectionName(pattern.Copy())
		entry.SetPattern(pattern.Copy())
		entry.SetKvStart(bdata.GetKvUsed())
		entry.SetKvEnd(entry.GetKvStart())
		entry.SetParent(nil)
		entry.SetPrefixLen(BrowscapComputePrefixLen(pattern))
		pos = entry.GetPrefixLen()
		for i = 0; i < BROWSCAP_NUM_CONTAINS; i++ {
			pos = BrowscapComputeContains(pattern, pos, entry.GetContainsStart()[i], entry.GetContainsLen()[i])
		}
	}
}
func BrowscapReadFile(filename *byte, browdata *BrowserData, persistent int) int {
	var fh zend.ZendFileHandle
	if filename == nil || filename[0] == '0' {
		return types.FAILURE
	}
	fh.InitFp(zend.VCWD_FOPEN(filename, "r"), filename)
	if fh.GetFp() == nil {
		faults.Error(faults.E_CORE_WARNING, "Cannot open '%s' for reading", filename)
		return types.FAILURE
	}
	browdata.SetHtab(types.NewArray(0))
	browdata.SetKvSize(16 * 1024)
	browdata.SetKvUsed(0)
	browdata.SetKv(zend.Pemalloc(b.SizeOf("browscap_kv") * browdata.GetKvSize()))

	/* Create parser context */
	var ctx = NewBrowscapParserCtx(browdata)
	zend.ZendParseIniFile(&fh, 1, zend.ZEND_INI_SCANNER_RAW, zend.ZendIniParserCbT(PhpBrowscapParserCb), ctx)

	return types.SUCCESS
}
func BrowscapBdataDtor(bdata *BrowserData, persistent int) {
	if bdata.GetHtab() != nil {
		bdata.GetHtab().Destroy()
		zend.Pefree(bdata.GetHtab(), persistent)
		bdata.SetHtab(nil)
		zend.Pefree(bdata.GetKv(), persistent)
		bdata.SetKv(nil)
	}
	bdata.GetFilename()[0] = '0'
}
func ZmStartupBrowscap(type_ int, module_number int) int {
	var browscap *byte = zend.INI_STR("browscap")

	/* ctor call not really needed for non-ZTS */

	if browscap != nil && browscap[0] {
		if BrowscapReadFile(browscap, &GlobalBdata, 1) == types.FAILURE {
			return types.FAILURE
		}
	}
	return types.SUCCESS
}
func ZmDeactivateBrowscap(type_ int, module_number int) int {
	var bdata *BrowserData = &(BROWSCAP_G(activation_bdata))
	if bdata.GetFilename()[0] != '0' {
		BrowscapBdataDtor(bdata, 0)
	}
	return types.SUCCESS
}
func ZmShutdownBrowscap(type_ int, module_number int) int {
	BrowscapBdataDtor(&GlobalBdata, 1)
	return types.SUCCESS
}
func BrowscapGetMinimumLength(entry *BrowscapEntry) int {
	var len_ int = entry.GetPrefixLen()
	var i int
	for i = 0; i < BROWSCAP_NUM_CONTAINS; i++ {
		len_ += entry.GetContainsLen()[i]
	}
	return len_
}
func BrowserRegCompare(entry *BrowscapEntry, agent_name *types.String, found_entry_ptr **BrowscapEntry) int {
	var found_entry *BrowscapEntry = *found_entry_ptr
	var pattern_lc *types.String
	var regex *types.String
	var cur *byte
	var i int
	var re *pcre2_code
	var match_data *pcre2_match_data
	var capture_count uint32
	var rc int

	/* Agent name too short */

	if agent_name.GetLen() < BrowscapGetMinimumLength(entry) {
		return 0
	}

	/* Quickly discard patterns where the prefix doesn't match. */

	if zend.ZendBinaryStrcasecmp(b.CastStr(agent_name.GetVal(), entry.GetPrefixLen()), b.CastStr(entry.GetPattern().GetVal(), entry.GetPrefixLen())) != 0 {
		return 0
	}

	/* Lowercase the pattern, the agent name is already lowercase */
	pattern_lc = types.ZendStringAlloc(entry.GetPattern().GetLen(), 0)
	zend.ZendStrTolowerCopy(pattern_lc.GetVal(), entry.GetPattern().GetVal(), entry.GetPattern().GetLen())

	/* Check if the agent contains the "contains" portions */

	cur = agent_name.GetVal() + entry.GetPrefixLen()
	for i = 0; i < BROWSCAP_NUM_CONTAINS; i++ {
		if entry.GetContainsLen()[i] != 0 {
			cur = zend.ZendMemnstr(cur, pattern_lc.GetVal()+entry.GetContainsStart()[i], entry.GetContainsLen()[i], agent_name.GetVal()+agent_name.GetLen())
			if cur == nil {
				//pattern_lc.Free()
				return 0
			}
			cur += entry.GetContainsLen()[i]
		}
	}

	/* See if we have an exact match, if so, we're done... */

	if agent_name.GetStr() == pattern_lc.GetStr() {
		*found_entry_ptr = entry
		//pattern_lc.Free()
		return 1
	}
	regex = BrowscapConvertPattern(entry.GetPattern(), 0)
	re = pcre_get_compiled_regex(regex, &capture_count)
	if re == nil {
		//pattern_lc.Free()
		// types.ZendStringRelease(regex)
		return 0
	}
	match_data = php_pcre_create_match_data(capture_count, re)
	if match_data == nil {
		//pattern_lc.Free()
		// types.ZendStringRelease(regex)
		return 0
	}
	rc = pcre2_match(re, PCRE2_SPTR(agent_name.GetVal()), agent_name.GetLen(), 0, 0, match_data, php_pcre_mctx())
	php_pcre_free_match_data(match_data)
	if PCRE2_ERROR_NOMATCH != rc {

		/* If we've found a possible browser, we need to do a comparison of the
		   number of characters changed in the user agent being checked versus
		   the previous match found and the current match. */

		if found_entry != nil {
			var i int
			var prev_len int = 0
			var curr_len int = 0
			var previous_match *types.String = found_entry.GetPattern()
			var current_match *types.String = entry.GetPattern()
			for i = 0; i < previous_match.GetLen(); i++ {
				switch previous_match.GetVal()[i] {
				case '?':
					fallthrough
				case '*':

				/* do nothing, ignore these characters in the count */

				default:
					prev_len++
				}
			}
			for i = 0; i < current_match.GetLen(); i++ {
				switch current_match.GetVal()[i] {
				case '?':
					fallthrough
				case '*':

				/* do nothing, ignore these characters in the count */

				default:
					curr_len++
				}
			}

			/* Pick which browser pattern replaces the least amount of
			   characters when compared to the original user agent string... */

			if prev_len < curr_len {
				*found_entry_ptr = entry
			}

			/* Pick which browser pattern replaces the least amount of
			   characters when compared to the original user agent string... */

		} else {
			*found_entry_ptr = entry
		}

		/* If we've found a possible browser, we need to do a comparison of the
		   number of characters changed in the user agent being checked versus
		   the previous match found and the current match. */

	}
	//pattern_lc.Free()
	//// types.ZendStringRelease(regex)
	return 0
}
func BrowscapZvalCopyCtor(p *types.Zval) {
	if p.IsRefcounted() {
		var str *types.String
		b.Assert(p.IsType(types.IS_STRING))
		str = p.String().Copy()
		p.SetString(str)
	}
}
func ZifGetBrowser(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, browserName *types.Zval, returnArray *types.Zval) {
	var agent_name *types.String = nil
	var lookup_browser_name *types.String
	var return_array types.ZendBool = 0
	var bdata *BrowserData
	var found_entry *BrowscapEntry = nil
	var agent_ht *types.Array
	if BROWSCAP_G(activation_bdata).filename[0] != '0' {
		bdata = &(BROWSCAP_G(activation_bdata))
		if bdata.GetHtab() == nil {
			if BrowscapReadFile(bdata.GetFilename(), bdata, 0) == types.FAILURE {
				return_value.SetFalse()
				return
			}
		}
	} else {
		if GlobalBdata.GetHtab() == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "browscap ini directive not set")
			return_value.SetFalse()
			return
		}
		bdata = &GlobalBdata
	}
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 2, 0)
			fp.StartOptional()
			agent_name = fp.ParseStrEx(true, false)
			return_array = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if agent_name == nil {
		var http_user_agent *types.Zval = nil
		if core.PG__().http_globals[core.TRACK_VARS_SERVER].GetType() == types.IS_ARRAY || zend.ZendIsAutoGlobalStr("_SERVER") != 0 {
			http_user_agent = core.PG__().http_globals[core.TRACK_VARS_SERVER].Array().KeyFind("HTTP_USER_AGENT")
		}
		if http_user_agent == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "HTTP_USER_AGENT variable is not set, cannot determine user agent name")
			return_value.SetFalse()
			return
		}
		agent_name = http_user_agent.String()
	}
	lookup_browser_name = zend.ZendStringTolower(agent_name)
	found_entry = types.ZendHashFindPtr(bdata.GetHtab(), lookup_browser_name.GetStr())
	if found_entry == nil {
		var entry *BrowscapEntry
		var __ht *types.Array = bdata.GetHtab()
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			entry = _z.Ptr()
			if BrowserRegCompare(entry, lookup_browser_name, &found_entry) != 0 {
				break
			}
		}
		if found_entry == nil {
			found_entry = types.ZendHashStrFindPtr(bdata.GetHtab(), DEFAULT_SECTION_NAME)
			if found_entry == nil {
				// types.ZendStringRelease(lookup_browser_name)
				return_value.SetFalse()
				return
			}
		}
	}
	agent_ht = BrowscapEntryToArray(bdata, found_entry)
	if return_array != 0 {
		return_value.SetArray(agent_ht)
	} else {
		zend.ObjectAndPropertiesInit(return_value, zend.ZendStandardClassDef, agent_ht)
	}
	for found_entry.GetParent() != nil {
		found_entry = types.ZendHashFindPtr(bdata.GetHtab(), found_entry.GetParent().GetStr())
		if found_entry == nil {
			break
		}
		agent_ht = BrowscapEntryToArray(bdata, found_entry)
		if return_array != 0 {
			types.ZendHashMerge(return_value.Array(), agent_ht, types.CopyCtorFuncT(BrowscapZvalCopyCtor), 0)
		} else {
			types.ZendHashMerge(types.Z_OBJPROP_P(return_value), agent_ht, types.CopyCtorFuncT(BrowscapZvalCopyCtor), 0)
		}
		agent_ht.Destroy()
		zend.Efree(agent_ht)
	}
	// types.ZendStringReleaseEx(lookup_browser_name, 0)
}
