package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
	"math"
	"strings"
)

func IsPlaceholder(c byte) bool { return c == '?' || c == '*' }
func BrowscapComputePrefixLen(pattern string) uint8 {
	max := uint8(b.Min(len(pattern), math.MaxUint8))

	var i uint8
	for i = 0; i < max; i++ {
		if IsPlaceholder(pattern[i]) {
			break
		}
	}
	return i
}
func BrowscapComputeContains(pattern *types.String, start_pos int, contains_start *uint16, contains_len *uint8) int {
	var i int = start_pos

	/* Find first non-placeholder character after prefix */
	for ; i < pattern.GetLen(); i++ {
		if !IsPlaceholder(pattern.GetStr()[i]) {
			/* Skip the case of a single non-placeholder character.
			 * Let's try to find something longer instead. */
			if i+1 < pattern.GetLen() && !IsPlaceholder(pattern.GetStr()[i+1]) {
				break
			}
		}
	}
	*contains_start = uint16(i)

	/* Find first placeholder character after that */

	for ; i < pattern.GetLen(); i++ {
		if IsPlaceholder(pattern.GetStr()[i]) {
			break
		}
	}
	*contains_len = uint8(b.Min(i-(*contains_start), math.MaxUint8))
	return i
}
func BrowscapConvertPatternEx(pattern string) string {
	pattern = ascii.StrToLower(pattern)

	var buf strings.Builder
	buf.WriteString("~^")
	for _, c := range []byte(pattern) {
		switch c {
		case '?':
			buf.WriteByte('.')
		case '*':
			buf.WriteByte('.')
			buf.WriteByte('*')
		case '.', '\\', '(', ')', '~', '+':
			buf.WriteByte('\\')
			buf.WriteByte(c)
		default:
			buf.WriteByte(c)
		}
	}
	buf.WriteString("$~")
	return buf.String()
}
func BrowscapInternStr(ctx *BrowscapParserCtx, str *types.String) *types.String {
	return ctx.GetInternedStr(str.GetStr())
}
func BrowscapInternStrCi(ctx *BrowscapParserCtx, str *types.String) *types.String {
	lcName := ascii.StrToLower(str.GetStr())
	return ctx.GetInternedStr(lcName)
}
func BrowscapEntryToArray(bdata *BrowserData, entry *BrowscapEntry) *types.Array {
	var ht = types.NewArray(8)
	ht.KeyAdd("browser_name_regex", types.NewZvalString(BrowscapConvertPatternEx(entry.GetPattern().GetStr())))
	ht.KeyAdd("browser_name_pattern", types.NewZvalString(entry.GetPattern().GetStr()))
	if entry.GetParent() != nil {
		ht.KeyAdd("parent", types.NewZvalString(entry.GetPattern().GetStr()))
	}
	bdata.EachKv(func(key string, value string) {
		ht.KeyAdd(key, types.NewZvalString(value))
	})
	return ht
}
func PhpBrowscapParserCb(arg1 *types.Zval, arg2 *types.Zval, arg3 *types.Zval, callback_type int, arg any) {
	var ctx *BrowscapParserCtx = arg
	var bdata *BrowserData = ctx.GetBdata()
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
				bdata.AddKv(new_key.GetStr(), new_value.GetStr())
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
		ctx.SetCurrentEntry(zend.Pemalloc(b.SizeOf("browscap_entry")))
		entry = ctx.GetCurrentEntry()
		bdata.GetHtab().KeyUpdate(pattern.GetStr(), types.NewZvalPtr(entry))
		ctx.SetCurrentSectionName(pattern.Copy())
		entry.SetPattern(pattern.Copy())
		entry.SetKvStart(bdata.GetKvUsed())
		entry.SetKvEnd(entry.GetKvStart())
		entry.SetParent(nil)
		entry.SetPrefixLen(BrowscapComputePrefixLen(pattern.GetStr()))
		pos = entry.GetPrefixLen()
		for i = 0; i < BROWSCAP_NUM_CONTAINS; i++ {
			pos = BrowscapComputeContains(pattern, pos, entry.GetContainsStart()[i], entry.GetContainsLen()[i])
		}
	}
}
func BrowscapReadFileEx(filename string) *BrowserData {
	if filename == "" {
		return nil
	}
	var fh = zend.NewFileHandleByOpenFile(filename)
	if fh == nil {
		faults.Error(faults.E_CORE_WARNING, "Cannot open '%s' for reading", filename)
		return nil
	}

	browserData := NewBrowserData(16 * 1024)

	/* Create parser context */
	var ctx = NewBrowscapParserCtx(browserData)
	zend.ZendParseIniFile(fh, 1, zend.ZEND_INI_SCANNER_RAW, PhpBrowscapParserCb, ctx)

	return browserData
}
func ZmStartupBrowscap(type_ int, module_number int) int {
	var browscap = zend.INI_STRING("browscap")

	/* ctor call not really needed for non-ZTS */
	if browscap != "" {
		GlobalBdata = BrowscapReadFileEx(browscap)
		if GlobalBdata == nil {
			return types.FAILURE
		}
	}
	return types.SUCCESS
}
func ZmShutdownBrowscap(type_ int, module_number int) int {
	GlobalBdata = nil
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
	var regex string
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
	if !ascii.StrCaseHasPrefix(agent_name.GetStr(), entry.GetPattern().GetStr()) {
		return 0
	}

	/* Lowercase the pattern, the agent name is already lowercase */
	pattern_lc = types.NewString(ascii.StrToLower(pattern_lc.GetStr()))

	/* Check if the agent contains the "contains" portions */

	cur = agent_name.GetVal() + entry.GetPrefixLen()
	for i = 0; i < BROWSCAP_NUM_CONTAINS; i++ {
		if entry.GetContainsLen()[i] != 0 {
			cur = operators.ZendMemnstr(cur, pattern_lc.GetVal()+entry.GetContainsStart()[i], entry.GetContainsLen()[i], agent_name.GetVal()+agent_name.GetLen())
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
	regex = BrowscapConvertPatternEx(entry.GetPattern().GetStr())
	re = pcre_get_compiled_regex(regex, &capture_count)
	if re == nil {
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
				switch previous_match.GetStr()[i] {
				case '?':
					fallthrough
				case '*':

				/* do nothing, ignore these characters in the count */

				default:
					prev_len++
				}
			}
			for i = 0; i < current_match.GetLen(); i++ {
				switch current_match.GetStr()[i] {
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
func ZifGetBrowser(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, browserName *string, returnArray bool) {
	var agent_name *string = browserName
	var lookup_browser_name *types.String
	var return_array bool = 0
	var bdata *BrowserData
	var found_entry *BrowscapEntry = nil
	var agent_ht *types.Array

	if GlobalBdata.GetHtab() == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "browscap ini directive not set")
		return_value.SetFalse()
		return
	}
	bdata = GlobalBdata

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
		if core.PG__().http_globals[core.TRACK_VARS_SERVER].IsArray() || zend.ZendIsAutoGlobal("_SERVER") {
			http_user_agent = core.PG__().http_globals[core.TRACK_VARS_SERVER].Array().KeyFind("HTTP_USER_AGENT")
		}
		if http_user_agent == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "HTTP_USER_AGENT variable is not set, cannot determine user agent name")
			return_value.SetFalse()
			return
		}
		agent_name = http_user_agent.String()
	}
	lookup_browser_name = operators.ZendStringTolower(agent_name)
	found_entry = types.ZendHashFindPtr(bdata.GetHtab(), lookup_browser_name.GetStr())
	if found_entry == nil {
		bdata.GetHtab().ForeachEx(func(_ types.ArrayKey, value *types.Zval) bool {
			var entry *BrowscapEntry = value.Ptr()
			if BrowserRegCompare(entry, lookup_browser_name, &found_entry) != 0 {
				return false
			}
			return true
		})
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
			types.ZendHashMerge(return_value.Array(), agent_ht, false)
		} else {
			types.ZendHashMerge(types.Z_OBJPROP_P(return_value), agent_ht, false)
		}
		agent_ht.Destroy()
		zend.Efree(agent_ht)
	}
}
