// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/sapi/cli"
	"sik/zend"
	"sik/zend/types"
)

func BROWSCAP_G(v __auto__) __auto__ { return BrowscapGlobals.v }
func BrowscapEntryDtor(zvalue *types.Zval) {
	var entry *BrowscapEntry = zvalue.GetPtr()
	types.ZendStringReleaseEx(entry.GetPattern(), 0)
	if entry.GetParent() != nil {
		types.ZendStringReleaseEx(entry.GetParent(), 0)
	}
	zend.Efree(entry)
}
func BrowscapEntryDtorPersistent(zvalue *types.Zval) {
	var entry *BrowscapEntry = zvalue.GetPtr()
	types.ZendStringReleaseEx(entry.GetPattern(), 1)
	if entry.GetParent() != nil {
		types.ZendStringReleaseEx(entry.GetParent(), 1)
	}
	zend.Pefree(entry, 1)
}
func IsPlaceholder(c byte) types.ZendBool { return c == '?' || c == '*' }
func BrowscapComputePrefixLen(pattern *types.ZendString) uint8 {
	var i int
	for i = 0; i < pattern.GetLen(); i++ {
		if IsPlaceholder(pattern.GetVal()[i]) != 0 {
			break
		}
	}
	return uint8(cli.MIN(i, UINT8_MAX))
}
func BrowscapComputeContains(pattern *types.ZendString, start_pos int, contains_start *uint16, contains_len *uint8) int {
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
func BrowscapComputeRegexLen(pattern *types.ZendString) int {
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
func BrowscapConvertPattern(pattern *types.ZendString, persistent int) *types.ZendString {
	var i int
	var j int = 0
	var t *byte
	var res *types.ZendString
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
func BrowscapInternStr(ctx *BrowscapParserCtx, str *types.ZendString, persistent types.ZendBool) *types.ZendString {
	var interned *types.ZendString = zend.ZendHashFindPtr(ctx.GetStrInterned(), str)
	if interned != nil {
		interned.AddRefcount()
	} else {
		interned = str.Copy()
		if persistent != 0 {
			interned = types.ZendNewInternedString(str)
		}
		zend.ZendHashAddNewPtr(ctx.GetStrInterned(), interned, interned)
	}
	return interned
}
func BrowscapInternStrCi(ctx *BrowscapParserCtx, str *types.ZendString, persistent types.ZendBool) *types.ZendString {
	var lcname *types.ZendString
	var interned *types.ZendString
	types.ZSTR_ALLOCA_ALLOC(lcname, str.GetLen())
	zend.ZendStrTolowerCopy(lcname.GetVal(), str.GetVal(), str.GetLen())
	interned = zend.ZendHashFindPtr(ctx.GetStrInterned(), lcname)
	if interned != nil {
		interned.AddRefcount()
	} else {
		interned = lcname.Dup(persistent)
		if persistent != 0 {
			interned = types.ZendNewInternedString(interned)
		}
		zend.ZendHashAddNewPtr(ctx.GetStrInterned(), interned, interned)
	}
	lcname.Free()
	return interned
}
func BrowscapAddKv(bdata *BrowserData, key *types.ZendString, value *types.ZendString, persistent types.ZendBool) {
	if bdata.GetKvUsed() == bdata.GetKvSize() {
		bdata.SetKvSize(bdata.GetKvSize() * 2)
		bdata.SetKv(zend.SafePerealloc(bdata.GetKv(), b.SizeOf("browscap_kv"), bdata.GetKvSize(), 0, persistent))
	}
	bdata.GetKv()[bdata.GetKvUsed()].SetKey(key)
	bdata.GetKv()[bdata.GetKvUsed()].SetValue(value)
	bdata.GetKvUsed()++
}
func BrowscapEntryToArray(bdata *BrowserData, entry *BrowscapEntry) *types.HashTable {
	var tmp types.Zval
	var i uint32
	var ht *types.HashTable = zend.ZendNewArray(8)
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
			var new_key *types.ZendString
			var new_value *types.ZendString

			/* Set proper value for true/false settings */

			if arg2.GetStr().GetLen() == 2 && !(strncasecmp(arg2.GetStr().GetVal(), "on", b.SizeOf("\"on\"")-1)) || arg2.GetStr().GetLen() == 3 && !(strncasecmp(arg2.GetStr().GetVal(), "yes", b.SizeOf("\"yes\"")-1)) || arg2.GetStr().GetLen() == 4 && !(strncasecmp(arg2.GetStr().GetVal(), "true", b.SizeOf("\"true\"")-1)) {
				new_value = types.ZSTR_CHAR('1')
			} else if arg2.GetStr().GetLen() == 2 && !(strncasecmp(arg2.GetStr().GetVal(), "no", b.SizeOf("\"no\"")-1)) || arg2.GetStr().GetLen() == 3 && !(strncasecmp(arg2.GetStr().GetVal(), "off", b.SizeOf("\"off\"")-1)) || arg2.GetStr().GetLen() == 4 && !(strncasecmp(arg2.GetStr().GetVal(), "none", b.SizeOf("\"none\"")-1)) || arg2.GetStr().GetLen() == 5 && !(strncasecmp(arg2.GetStr().GetVal(), "false", b.SizeOf("\"false\"")-1)) {
				new_value = types.ZSTR_EMPTY_ALLOC()
			} else {
				new_value = BrowscapInternStr(ctx, arg2.GetStr(), persistent)
			}
			if !(strcasecmp(arg1.GetStr().GetVal(), "parent")) {

				/* parent entry can not be same as current section -> causes infinite loop! */

				if ctx.GetCurrentSectionName() != nil && !(strcasecmp(ctx.GetCurrentSectionName().GetVal(), arg2.GetStr().GetVal())) {
					zend.ZendError(zend.E_CORE_ERROR, "Invalid browscap ini file: "+"'Parent' value cannot be same as the section name: %s "+"(in file %s)", ctx.GetCurrentSectionName().GetVal(), zend.INI_STR("browscap"))
					return
				}
				if ctx.GetCurrentEntry().GetParent() != nil {
					types.ZendStringRelease(ctx.GetCurrentEntry().GetParent())
				}
				ctx.GetCurrentEntry().SetParent(new_value)
			} else {
				new_key = BrowscapInternStrCi(ctx, arg1.GetStr(), persistent)
				BrowscapAddKv(bdata, new_key, new_value, persistent)
				ctx.GetCurrentEntry().SetKvEnd(bdata.GetKvUsed())
			}
		}
	case zend.ZEND_INI_PARSER_SECTION:
		var entry *BrowscapEntry
		var pattern *types.ZendString = types.Z_STR_P(arg1)
		var pos int
		var i int
		if pattern.GetLen() > UINT16_MAX {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Skipping excessively long pattern of length %zd", pattern.GetLen())
			break
		}
		if persistent != 0 {
			pattern = types.ZendNewInternedString(pattern.Copy())
			types.ZendStringRelease(pattern)
		}
		ctx.SetCurrentEntry(zend.Pemalloc(b.SizeOf("browscap_entry"), persistent))
		entry = ctx.GetCurrentEntry()
		zend.ZendHashUpdatePtr(bdata.GetHtab(), pattern, entry)
		if ctx.GetCurrentSectionName() != nil {
			types.ZendStringRelease(ctx.GetCurrentSectionName())
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
func StrInternedDtor(zv *types.Zval) { types.ZendStringRelease(zv.GetStr()) }
func BrowscapReadFile(filename *byte, browdata *BrowserData, persistent int) int {
	var fh zend.ZendFileHandle
	var ctx BrowscapParserCtx = MakeBrowscapParserCtx(0)
	if filename == nil || filename[0] == '0' {
		return types.FAILURE
	}
	fh.InitFp(zend.VCWD_FOPEN(filename, "r"), filename)
	if fh.GetFp() == nil {
		zend.ZendError(zend.E_CORE_WARNING, "Cannot open '%s' for reading", filename)
		return types.FAILURE
	}
	browdata.SetHtab(zend.Pemalloc(sizeof*browdata.GetHtab(), persistent))
	zend.ZendHashInitEx(browdata.GetHtab(), 0, nil, b.Cond(persistent != 0, BrowscapEntryDtorPersistent, BrowscapEntryDtor), persistent, 0)
	browdata.SetKvSize(16 * 1024)
	browdata.SetKvUsed(0)
	browdata.SetKv(zend.Pemalloc(b.SizeOf("browscap_kv")*browdata.GetKvSize(), persistent))

	/* Create parser context */

	ctx.SetBdata(browdata)
	ctx.SetCurrentEntry(nil)
	ctx.SetCurrentSectionName(nil)
	zend.ZendHashInit(ctx.GetStrInterned(), 8, nil, StrInternedDtor, persistent)
	zend.ZendParseIniFile(&fh, 1, zend.ZEND_INI_SCANNER_RAW, zend.ZendIniParserCbT(PhpBrowscapParserCb), &ctx)

	/* Destroy parser context */

	if ctx.GetCurrentSectionName() != nil {
		types.ZendStringRelease(ctx.GetCurrentSectionName())
	}
	ctx.GetStrInterned().Destroy()
	return types.SUCCESS
}
func BrowscapBdataDtor(bdata *BrowserData, persistent int) {
	if bdata.GetHtab() != nil {
		var i uint32
		bdata.GetHtab().Destroy()
		zend.Pefree(bdata.GetHtab(), persistent)
		bdata.SetHtab(nil)
		for i = 0; i < bdata.GetKvUsed(); i++ {
			types.ZendStringRelease(bdata.GetKv()[i].GetKey())
			types.ZendStringRelease(bdata.GetKv()[i].GetValue())
		}
		zend.Pefree(bdata.GetKv(), persistent)
		bdata.SetKv(nil)
	}
	bdata.GetFilename()[0] = '0'
}
func OnChangeBrowscap(
	entry *zend.ZendIniEntry,
	new_value *types.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	if stage == core.PHP_INI_STAGE_STARTUP {

		/* value handled in browscap.c's MINIT */

		return types.SUCCESS

		/* value handled in browscap.c's MINIT */

	} else if stage == core.PHP_INI_STAGE_ACTIVATE {
		var bdata *BrowserData = &(BROWSCAP_G(activation_bdata))
		if bdata.GetFilename()[0] != '0' {
			BrowscapBdataDtor(bdata, 0)
		}
		if zend.VCWD_REALPATH(new_value.GetVal(), bdata.GetFilename()) == nil {
			return types.FAILURE
		}
		return types.SUCCESS
	}
	return types.FAILURE
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
func BrowserRegCompare(entry *BrowscapEntry, agent_name *types.ZendString, found_entry_ptr **BrowscapEntry) int {
	var found_entry *BrowscapEntry = *found_entry_ptr
	var pattern_lc *types.ZendString
	var regex *types.ZendString
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

	if zend.ZendBinaryStrcasecmp(agent_name.GetVal(), entry.GetPrefixLen(), entry.GetPattern().GetVal(), entry.GetPrefixLen()) != 0 {
		return 0
	}

	/* Lowercase the pattern, the agent name is already lowercase */

	types.ZSTR_ALLOCA_ALLOC(pattern_lc, entry.GetPattern().GetLen())
	zend.ZendStrTolowerCopy(pattern_lc.GetVal(), entry.GetPattern().GetVal(), entry.GetPattern().GetLen())

	/* Check if the agent contains the "contains" portions */

	cur = agent_name.GetVal() + entry.GetPrefixLen()
	for i = 0; i < BROWSCAP_NUM_CONTAINS; i++ {
		if entry.GetContainsLen()[i] != 0 {
			cur = zend.ZendMemnstr(cur, pattern_lc.GetVal()+entry.GetContainsStart()[i], entry.GetContainsLen()[i], agent_name.GetVal()+agent_name.GetLen())
			if cur == nil {
				pattern_lc.Free()
				return 0
			}
			cur += entry.GetContainsLen()[i]
		}
	}

	/* See if we have an exact match, if so, we're done... */

	if types.ZendStringEquals(agent_name, pattern_lc) != 0 {
		*found_entry_ptr = entry
		pattern_lc.Free()
		return 1
	}
	regex = BrowscapConvertPattern(entry.GetPattern(), 0)
	re = pcre_get_compiled_regex(regex, &capture_count)
	if re == nil {
		pattern_lc.Free()
		types.ZendStringRelease(regex)
		return 0
	}
	match_data = php_pcre_create_match_data(capture_count, re)
	if match_data == nil {
		pattern_lc.Free()
		types.ZendStringRelease(regex)
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
			var previous_match *types.ZendString = found_entry.GetPattern()
			var current_match *types.ZendString = entry.GetPattern()
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
	pattern_lc.Free()
	types.ZendStringRelease(regex)
	return 0
}
func BrowscapZvalCopyCtor(p *types.Zval) {
	if p.IsRefcounted() {
		var str *types.ZendString
		b.Assert(p.IsType(types.IS_STRING))
		str = p.GetStr()
		if (str.GetGcFlags() & types.GC_PERSISTENT) == 0 {
			str.AddRefcount()
		} else {
			p.SetString(types.ZendStringInit(str.GetVal(), str.GetLen(), 0))
		}
	}
}
func ZifGetBrowser(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var agent_name *types.ZendString = nil
	var lookup_browser_name *types.ZendString
	var return_array types.ZendBool = 0
	var bdata *BrowserData
	var found_entry *BrowscapEntry = nil
	var agent_ht *types.HashTable
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
			core.PhpErrorDocref(nil, zend.E_WARNING, "browscap ini directive not set")
			return_value.SetFalse()
			return
		}
		bdata = &GlobalBdata
	}
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = executeData.Arg(0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &agent_name, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &return_array, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if agent_name == nil {
		var http_user_agent *types.Zval = nil
		if core.PG(http_globals)[core.TRACK_VARS_SERVER].u1.v.type_ == types.IS_ARRAY || zend.ZendIsAutoGlobalStr(zend.ZEND_STRL("_SERVER")) != 0 {
			http_user_agent = core.PG(http_globals)[core.TRACK_VARS_SERVER].GetArr().KeyFind("HTTP_USER_AGENT")
		}
		if http_user_agent == nil {
			core.PhpErrorDocref(nil, zend.E_WARNING, "HTTP_USER_AGENT variable is not set, cannot determine user agent name")
			return_value.SetFalse()
			return
		}
		agent_name = http_user_agent.GetStr()
	}
	lookup_browser_name = zend.ZendStringTolower(agent_name)
	found_entry = zend.ZendHashFindPtr(bdata.GetHtab(), lookup_browser_name)
	if found_entry == nil {
		var entry *BrowscapEntry
		var __ht *types.HashTable = bdata.GetHtab()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			entry = _z.GetPtr()
			if BrowserRegCompare(entry, lookup_browser_name, &found_entry) != 0 {
				break
			}
		}
		if found_entry == nil {
			found_entry = zend.ZendHashStrFindPtr(bdata.GetHtab(), DEFAULT_SECTION_NAME, b.SizeOf("DEFAULT_SECTION_NAME")-1)
			if found_entry == nil {
				types.ZendStringRelease(lookup_browser_name)
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
		found_entry = zend.ZendHashFindPtr(bdata.GetHtab(), found_entry.GetParent())
		if found_entry == nil {
			break
		}
		agent_ht = BrowscapEntryToArray(bdata, found_entry)
		if return_array != 0 {
			zend.ZendHashMerge(return_value.GetArr(), agent_ht, types.CopyCtorFuncT(BrowscapZvalCopyCtor), 0)
		} else {
			zend.ZendHashMerge(types.Z_OBJPROP_P(return_value), agent_ht, types.CopyCtorFuncT(BrowscapZvalCopyCtor), 0)
		}
		agent_ht.Destroy()
		zend.Efree(agent_ht)
	}
	types.ZendStringReleaseEx(lookup_browser_name, 0)
}
