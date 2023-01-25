// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/sapi/cli"
	"sik/zend"
)

// Source: <ext/standard/browscap.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Author: Zeev Suraski <zeev@php.net>                                  |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_browscap.h"

// # include "php_ini.h"

// # include "php_string.h"

// failed # include "ext/pcre/php_pcre.h"

// # include "zend_ini_scanner.h"

// # include "zend_globals.h"

const BROWSCAP_NUM_CONTAINS = 5

/* browser data defined in startup phase, eagerly loaded in MINIT */

var GlobalBdata BrowserData = BrowserData{0}

/* browser data defined in activation phase, lazily loaded in get_browser.
 * Per request and per thread, if applicable */

var BrowscapGlobals ZendBrowscapGlobals

func BROWSCAP_G(v __auto__) __auto__ { return BrowscapGlobals.v }

const DEFAULT_SECTION_NAME = "Default Browser Capability Settings"

/* OBJECTS_FIXME: This whole extension needs going through. The use of objects looks pretty broken here */

func BrowscapEntryDtor(zvalue *zend.Zval) {
	var entry *BrowscapEntry = zend.Z_PTR_P(zvalue)
	zend.ZendStringReleaseEx(entry.GetPattern(), 0)
	if entry.GetParent() != nil {
		zend.ZendStringReleaseEx(entry.GetParent(), 0)
	}
	zend.Efree(entry)
}
func BrowscapEntryDtorPersistent(zvalue *zend.Zval) {
	var entry *BrowscapEntry = zend.Z_PTR_P(zvalue)
	zend.ZendStringReleaseEx(entry.GetPattern(), 1)
	if entry.GetParent() != nil {
		zend.ZendStringReleaseEx(entry.GetParent(), 1)
	}
	zend.Pefree(entry, 1)
}
func IsPlaceholder(c byte) zend.ZendBool { return c == '?' || c == '*' }

/* Length of prefix not containing any wildcards */

func BrowscapComputePrefixLen(pattern *zend.ZendString) uint8 {
	var i int
	for i = 0; i < zend.ZSTR_LEN(pattern); i++ {
		if IsPlaceholder(zend.ZSTR_VAL(pattern)[i]) != 0 {
			break
		}
	}
	return uint8(cli.MIN(i, UINT8_MAX))
}
func BrowscapComputeContains(pattern *zend.ZendString, start_pos int, contains_start *uint16, contains_len *uint8) int {
	var i int = start_pos

	/* Find first non-placeholder character after prefix */

	for ; i < zend.ZSTR_LEN(pattern); i++ {
		if IsPlaceholder(zend.ZSTR_VAL(pattern)[i]) == 0 {

			/* Skip the case of a single non-placeholder character.
			 * Let's try to find something longer instead. */

			if i+1 < zend.ZSTR_LEN(pattern) && IsPlaceholder(zend.ZSTR_VAL(pattern)[i+1]) == 0 {
				break
			}

			/* Skip the case of a single non-placeholder character.
			 * Let's try to find something longer instead. */

		}
	}
	*contains_start = uint16(i)

	/* Find first placeholder character after that */

	for ; i < zend.ZSTR_LEN(pattern); i++ {
		if IsPlaceholder(zend.ZSTR_VAL(pattern)[i]) != 0 {
			break
		}
	}
	*contains_len = uint8(cli.MIN(i-(*contains_start), UINT8_MAX))
	return i
}

/* Length of regex, including escapes, anchors, etc. */

func BrowscapComputeRegexLen(pattern *zend.ZendString) int {
	var i int
	var len_ int = zend.ZSTR_LEN(pattern)
	for i = 0; i < zend.ZSTR_LEN(pattern); i++ {
		switch zend.ZSTR_VAL(pattern)[i] {
		case '*':

		case '.':

		case '\\':

		case '(':

		case ')':

		case '~':

		case '+':
			len_++
			break
		}
	}
	return len_ + b.SizeOf("\"~^$~\"") - 1
}
func BrowscapConvertPattern(pattern *zend.ZendString, persistent int) *zend.ZendString {
	var i int
	var j int = 0
	var t *byte
	var res *zend.ZendString
	var lc_pattern *byte
	res = zend.ZendStringAlloc(BrowscapComputeRegexLen(pattern), persistent)
	t = zend.ZSTR_VAL(res)
	lc_pattern = zend.DoAlloca(zend.ZSTR_LEN(pattern)+1, use_heap)
	zend.ZendStrTolowerCopy(lc_pattern, zend.ZSTR_VAL(pattern), zend.ZSTR_LEN(pattern))
	t[b.PostInc(&j)] = '~'
	t[b.PostInc(&j)] = '^'
	for i = 0; i < zend.ZSTR_LEN(pattern); {
		switch lc_pattern[i] {
		case '?':
			t[j] = '.'
			break
		case '*':
			t[b.PostInc(&j)] = '.'
			t[j] = '*'
			break
		case '.':
			t[b.PostInc(&j)] = '\\'
			t[j] = '.'
			break
		case '\\':
			t[b.PostInc(&j)] = '\\'
			t[j] = '\\'
			break
		case '(':
			t[b.PostInc(&j)] = '\\'
			t[j] = '('
			break
		case ')':
			t[b.PostInc(&j)] = '\\'
			t[j] = ')'
			break
		case '~':
			t[b.PostInc(&j)] = '\\'
			t[j] = '~'
			break
		case '+':
			t[b.PostInc(&j)] = '\\'
			t[j] = '+'
			break
		default:
			t[j] = lc_pattern[i]
			break
		}
		i++
		j++
	}
	t[b.PostInc(&j)] = '$'
	t[b.PostInc(&j)] = '~'
	t[j] = 0
	zend.ZSTR_LEN(res) = j
	zend.FreeAlloca(lc_pattern, use_heap)
	return res
}

/* }}} */

func BrowscapInternStr(ctx *BrowscapParserCtx, str *zend.ZendString, persistent zend.ZendBool) *zend.ZendString {
	var interned *zend.ZendString = zend.ZendHashFindPtr(&ctx.str_interned, str)
	if interned != nil {
		zend.ZendStringAddref(interned)
	} else {
		interned = zend.ZendStringCopy(str)
		if persistent != 0 {
			interned = zend.ZendNewInternedString(str)
		}
		zend.ZendHashAddNewPtr(&ctx.str_interned, interned, interned)
	}
	return interned
}
func BrowscapInternStrCi(ctx *BrowscapParserCtx, str *zend.ZendString, persistent zend.ZendBool) *zend.ZendString {
	var lcname *zend.ZendString
	var interned *zend.ZendString
	zend.ZSTR_ALLOCA_ALLOC(lcname, zend.ZSTR_LEN(str), use_heap)
	zend.ZendStrTolowerCopy(zend.ZSTR_VAL(lcname), zend.ZSTR_VAL(str), zend.ZSTR_LEN(str))
	interned = zend.ZendHashFindPtr(&ctx.str_interned, lcname)
	if interned != nil {
		zend.ZendStringAddref(interned)
	} else {
		interned = zend.ZendStringDup(lcname, persistent)
		if persistent != 0 {
			interned = zend.ZendNewInternedString(interned)
		}
		zend.ZendHashAddNewPtr(&ctx.str_interned, interned, interned)
	}
	zend.ZSTR_ALLOCA_FREE(lcname, use_heap)
	return interned
}
func BrowscapAddKv(bdata *BrowserData, key *zend.ZendString, value *zend.ZendString, persistent zend.ZendBool) {
	if bdata.GetKvUsed() == bdata.GetKvSize() {
		bdata.SetKvSize(bdata.GetKvSize() * 2)
		bdata.SetKv(zend.SafePerealloc(bdata.GetKv(), b.SizeOf("browscap_kv"), bdata.GetKvSize(), 0, persistent))
	}
	bdata.GetKv()[bdata.GetKvUsed()].SetKey(key)
	bdata.GetKv()[bdata.GetKvUsed()].SetValue(value)
	bdata.GetKvUsed()++
}
func BrowscapEntryToArray(bdata *BrowserData, entry *BrowscapEntry) *zend.HashTable {
	var tmp zend.Zval
	var i uint32
	var ht *zend.HashTable = zend.ZendNewArray(8)
	zend.ZVAL_STR(&tmp, BrowscapConvertPattern(entry.GetPattern(), 0))
	zend.ZendHashStrAdd(ht, "browser_name_regex", b.SizeOf("\"browser_name_regex\"")-1, &tmp)
	zend.ZVAL_STR_COPY(&tmp, entry.GetPattern())
	zend.ZendHashStrAdd(ht, "browser_name_pattern", b.SizeOf("\"browser_name_pattern\"")-1, &tmp)
	if entry.GetParent() != nil {
		zend.ZVAL_STR_COPY(&tmp, entry.GetParent())
		zend.ZendHashStrAdd(ht, "parent", b.SizeOf("\"parent\"")-1, &tmp)
	}
	for i = entry.GetKvStart(); i < entry.GetKvEnd(); i++ {
		zend.ZVAL_STR_COPY(&tmp, bdata.GetKv()[i].GetValue())
		zend.ZendHashAdd(ht, bdata.GetKv()[i].GetKey(), &tmp)
	}
	return ht
}
func PhpBrowscapParserCb(arg1 *zend.Zval, arg2 *zend.Zval, arg3 *zend.Zval, callback_type int, arg any) {
	var ctx *BrowscapParserCtx = arg
	var bdata *BrowserData = ctx.GetBdata()
	var persistent int = zend.GC_FLAGS(bdata.GetHtab()) & zend.IS_ARRAY_PERSISTENT
	if arg1 == nil {
		return
	}
	switch callback_type {
	case zend.ZEND_INI_PARSER_ENTRY:
		if ctx.GetCurrentEntry() != nil && arg2 != nil {
			var new_key *zend.ZendString
			var new_value *zend.ZendString

			/* Set proper value for true/false settings */

			if zend.Z_STRLEN_P(arg2) == 2 && !(strncasecmp(zend.Z_STRVAL_P(arg2), "on", b.SizeOf("\"on\"")-1)) || zend.Z_STRLEN_P(arg2) == 3 && !(strncasecmp(zend.Z_STRVAL_P(arg2), "yes", b.SizeOf("\"yes\"")-1)) || zend.Z_STRLEN_P(arg2) == 4 && !(strncasecmp(zend.Z_STRVAL_P(arg2), "true", b.SizeOf("\"true\"")-1)) {
				new_value = zend.ZSTR_CHAR('1')
			} else if zend.Z_STRLEN_P(arg2) == 2 && !(strncasecmp(zend.Z_STRVAL_P(arg2), "no", b.SizeOf("\"no\"")-1)) || zend.Z_STRLEN_P(arg2) == 3 && !(strncasecmp(zend.Z_STRVAL_P(arg2), "off", b.SizeOf("\"off\"")-1)) || zend.Z_STRLEN_P(arg2) == 4 && !(strncasecmp(zend.Z_STRVAL_P(arg2), "none", b.SizeOf("\"none\"")-1)) || zend.Z_STRLEN_P(arg2) == 5 && !(strncasecmp(zend.Z_STRVAL_P(arg2), "false", b.SizeOf("\"false\"")-1)) {
				new_value = zend.ZSTR_EMPTY_ALLOC()
			} else {
				new_value = BrowscapInternStr(ctx, zend.Z_STR_P(arg2), persistent)
			}
			if !(strcasecmp(zend.Z_STRVAL_P(arg1), "parent")) {

				/* parent entry can not be same as current section -> causes infinite loop! */

				if ctx.GetCurrentSectionName() != nil && !(strcasecmp(zend.ZSTR_VAL(ctx.GetCurrentSectionName()), zend.Z_STRVAL_P(arg2))) {
					zend.ZendError(zend.E_CORE_ERROR, "Invalid browscap ini file: "+"'Parent' value cannot be same as the section name: %s "+"(in file %s)", zend.ZSTR_VAL(ctx.GetCurrentSectionName()), zend.INI_STR("browscap"))
					return
				}
				if ctx.GetCurrentEntry().GetParent() != nil {
					zend.ZendStringRelease(ctx.GetCurrentEntry().GetParent())
				}
				ctx.GetCurrentEntry().SetParent(new_value)
			} else {
				new_key = BrowscapInternStrCi(ctx, zend.Z_STR_P(arg1), persistent)
				BrowscapAddKv(bdata, new_key, new_value, persistent)
				ctx.GetCurrentEntry().SetKvEnd(bdata.GetKvUsed())
			}
		}
		break
	case zend.ZEND_INI_PARSER_SECTION:
		var entry *BrowscapEntry
		var pattern *zend.ZendString = zend.Z_STR_P(arg1)
		var pos int
		var i int
		if zend.ZSTR_LEN(pattern) > UINT16_MAX {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Skipping excessively long pattern of length %zd", zend.ZSTR_LEN(pattern))
			break
		}
		if persistent != 0 {
			pattern = zend.ZendNewInternedString(zend.ZendStringCopy(pattern))
			if zend.ZSTR_IS_INTERNED(pattern) != 0 {
				zend.Z_TYPE_FLAGS_P(arg1) = 0
			} else {
				zend.ZendStringRelease(pattern)
			}
		}
		ctx.SetCurrentEntry(zend.Pemalloc(b.SizeOf("browscap_entry"), persistent))
		entry = ctx.GetCurrentEntry()
		zend.ZendHashUpdatePtr(bdata.GetHtab(), pattern, entry)
		if ctx.GetCurrentSectionName() != nil {
			zend.ZendStringRelease(ctx.GetCurrentSectionName())
		}
		ctx.SetCurrentSectionName(zend.ZendStringCopy(pattern))
		entry.SetPattern(zend.ZendStringCopy(pattern))
		entry.SetKvStart(bdata.GetKvUsed())
		entry.SetKvEnd(entry.GetKvStart())
		entry.SetParent(nil)
		entry.SetPrefixLen(BrowscapComputePrefixLen(pattern))
		pos = entry.GetPrefixLen()
		for i = 0; i < BROWSCAP_NUM_CONTAINS; i++ {
			pos = BrowscapComputeContains(pattern, pos, &entry.contains_start[i], &entry.contains_len[i])
		}
		break
	}
}

/* }}} */

func StrInternedDtor(zv *zend.Zval) { zend.ZendStringRelease(zend.Z_STR_P(zv)) }
func BrowscapReadFile(filename *byte, browdata *BrowserData, persistent int) int {
	var fh zend.ZendFileHandle
	var ctx BrowscapParserCtx = BrowscapParserCtx{0}
	if filename == nil || filename[0] == '0' {
		return zend.FAILURE
	}
	zend.ZendStreamInitFp(&fh, zend.VCWD_FOPEN(filename, "r"), filename)
	if fh.handle.fp == nil {
		zend.ZendError(zend.E_CORE_WARNING, "Cannot open '%s' for reading", filename)
		return zend.FAILURE
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
	zend.ZendHashInit(&ctx.str_interned, 8, nil, StrInternedDtor, persistent)
	zend.ZendParseIniFile(&fh, 1, zend.ZEND_INI_SCANNER_RAW, zend.ZendIniParserCbT(PhpBrowscapParserCb), &ctx)

	/* Destroy parser context */

	if ctx.GetCurrentSectionName() != nil {
		zend.ZendStringRelease(ctx.GetCurrentSectionName())
	}
	zend.ZendHashDestroy(&ctx.str_interned)
	return zend.SUCCESS
}

/* }}} */

func BrowscapBdataDtor(bdata *BrowserData, persistent int) {
	if bdata.GetHtab() != nil {
		var i uint32
		zend.ZendHashDestroy(bdata.GetHtab())
		zend.Pefree(bdata.GetHtab(), persistent)
		bdata.SetHtab(nil)
		for i = 0; i < bdata.GetKvUsed(); i++ {
			zend.ZendStringRelease(bdata.GetKv()[i].GetKey())
			zend.ZendStringRelease(bdata.GetKv()[i].GetValue())
		}
		zend.Pefree(bdata.GetKv(), persistent)
		bdata.SetKv(nil)
	}
	bdata.GetFilename()[0] = '0'
}

/* }}} */

func OnChangeBrowscap(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	if stage == core.PHP_INI_STAGE_STARTUP {

		/* value handled in browscap.c's MINIT */

		return zend.SUCCESS

		/* value handled in browscap.c's MINIT */

	} else if stage == core.PHP_INI_STAGE_ACTIVATE {
		var bdata *BrowserData = &BROWSCAP_G(activation_bdata)
		if bdata.GetFilename()[0] != '0' {
			BrowscapBdataDtor(bdata, 0)
		}
		if zend.VCWD_REALPATH(zend.ZSTR_VAL(new_value), bdata.GetFilename()) == nil {
			return zend.FAILURE
		}
		return zend.SUCCESS
	}
	return zend.FAILURE
}

/* }}} */

func ZmStartupBrowscap(type_ int, module_number int) int {
	var browscap *byte = zend.INI_STR("browscap")

	/* ctor call not really needed for non-ZTS */

	if browscap != nil && browscap[0] {
		if BrowscapReadFile(browscap, &GlobalBdata, 1) == zend.FAILURE {
			return zend.FAILURE
		}
	}
	return zend.SUCCESS
}

/* }}} */

func ZmDeactivateBrowscap(type_ int, module_number int) int {
	var bdata *BrowserData = &BROWSCAP_G(activation_bdata)
	if bdata.GetFilename()[0] != '0' {
		BrowscapBdataDtor(bdata, 0)
	}
	return zend.SUCCESS
}

/* }}} */

func ZmShutdownBrowscap(type_ int, module_number int) int {
	BrowscapBdataDtor(&GlobalBdata, 1)
	return zend.SUCCESS
}

/* }}} */

func BrowscapGetMinimumLength(entry *BrowscapEntry) int {
	var len_ int = entry.GetPrefixLen()
	var i int
	for i = 0; i < BROWSCAP_NUM_CONTAINS; i++ {
		len_ += entry.GetContainsLen()[i]
	}
	return len_
}
func BrowserRegCompare(entry *BrowscapEntry, agent_name *zend.ZendString, found_entry_ptr **BrowscapEntry) int {
	var found_entry *BrowscapEntry = *found_entry_ptr
	var pattern_lc *zend.ZendString
	var regex *zend.ZendString
	var cur *byte
	var i int
	var re *pcre2_code
	var match_data *pcre2_match_data
	var capture_count uint32
	var rc int

	/* Agent name too short */

	if zend.ZSTR_LEN(agent_name) < BrowscapGetMinimumLength(entry) {
		return 0
	}

	/* Quickly discard patterns where the prefix doesn't match. */

	if zend.ZendBinaryStrcasecmp(zend.ZSTR_VAL(agent_name), entry.GetPrefixLen(), zend.ZSTR_VAL(entry.GetPattern()), entry.GetPrefixLen()) != 0 {
		return 0
	}

	/* Lowercase the pattern, the agent name is already lowercase */

	zend.ZSTR_ALLOCA_ALLOC(pattern_lc, zend.ZSTR_LEN(entry.GetPattern()), use_heap)
	zend.ZendStrTolowerCopy(zend.ZSTR_VAL(pattern_lc), zend.ZSTR_VAL(entry.GetPattern()), zend.ZSTR_LEN(entry.GetPattern()))

	/* Check if the agent contains the "contains" portions */

	cur = zend.ZSTR_VAL(agent_name) + entry.GetPrefixLen()
	for i = 0; i < BROWSCAP_NUM_CONTAINS; i++ {
		if entry.GetContainsLen()[i] != 0 {
			cur = zend.ZendMemnstr(cur, zend.ZSTR_VAL(pattern_lc)+entry.GetContainsStart()[i], entry.GetContainsLen()[i], zend.ZSTR_VAL(agent_name)+zend.ZSTR_LEN(agent_name))
			if cur == nil {
				zend.ZSTR_ALLOCA_FREE(pattern_lc, use_heap)
				return 0
			}
			cur += entry.GetContainsLen()[i]
		}
	}

	/* See if we have an exact match, if so, we're done... */

	if zend.ZendStringEquals(agent_name, pattern_lc) != 0 {
		*found_entry_ptr = entry
		zend.ZSTR_ALLOCA_FREE(pattern_lc, use_heap)
		return 1
	}
	regex = BrowscapConvertPattern(entry.GetPattern(), 0)
	re = pcre_get_compiled_regex(regex, &capture_count)
	if re == nil {
		zend.ZSTR_ALLOCA_FREE(pattern_lc, use_heap)
		zend.ZendStringRelease(regex)
		return 0
	}
	match_data = php_pcre_create_match_data(capture_count, re)
	if match_data == nil {
		zend.ZSTR_ALLOCA_FREE(pattern_lc, use_heap)
		zend.ZendStringRelease(regex)
		return 0
	}
	rc = pcre2_match(re, PCRE2_SPTR(zend.ZSTR_VAL(agent_name)), zend.ZSTR_LEN(agent_name), 0, 0, match_data, php_pcre_mctx())
	php_pcre_free_match_data(match_data)
	if PCRE2_ERROR_NOMATCH != rc {

		/* If we've found a possible browser, we need to do a comparison of the
		   number of characters changed in the user agent being checked versus
		   the previous match found and the current match. */

		if found_entry != nil {
			var i int
			var prev_len int = 0
			var curr_len int = 0
			var previous_match *zend.ZendString = found_entry.GetPattern()
			var current_match *zend.ZendString = entry.GetPattern()
			for i = 0; i < zend.ZSTR_LEN(previous_match); i++ {
				switch zend.ZSTR_VAL(previous_match)[i] {
				case '?':

				case '*':

					/* do nothing, ignore these characters in the count */

					break
				default:
					prev_len++
				}
			}
			for i = 0; i < zend.ZSTR_LEN(current_match); i++ {
				switch zend.ZSTR_VAL(current_match)[i] {
				case '?':

				case '*':

					/* do nothing, ignore these characters in the count */

					break
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
	zend.ZSTR_ALLOCA_FREE(pattern_lc, use_heap)
	zend.ZendStringRelease(regex)
	return 0
}

/* }}} */

func BrowscapZvalCopyCtor(p *zend.Zval) {
	if zend.Z_REFCOUNTED_P(p) {
		var str *zend.ZendString
		zend.ZEND_ASSERT(zend.Z_TYPE_P(p) == zend.IS_STRING)
		str = zend.Z_STR_P(p)
		if (zend.GC_FLAGS(str) & zend.GC_PERSISTENT) == 0 {
			zend.GC_ADDREF(str)
		} else {
			zend.ZVAL_NEW_STR(p, zend.ZendStringInit(zend.ZSTR_VAL(str), zend.ZSTR_LEN(str), 0))
		}
	}
}

/* }}} */

func ZifGetBrowser(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var agent_name *zend.ZendString = nil
	var lookup_browser_name *zend.ZendString
	var return_array zend.ZendBool = 0
	var bdata *BrowserData
	var found_entry *BrowscapEntry = nil
	var agent_ht *zend.HashTable
	if BROWSCAP_G(activation_bdata).filename[0] != '0' {
		bdata = &BROWSCAP_G(activation_bdata)
		if bdata.GetHtab() == nil {
			if BrowscapReadFile(bdata.GetFilename(), bdata, 0) == zend.FAILURE {
				zend.RETVAL_FALSE
				return
			}
		}
	} else {
		if GlobalBdata.GetHtab() == nil {
			core.PhpErrorDocref(nil, zend.E_WARNING, "browscap ini directive not set")
			zend.RETVAL_FALSE
			return
		}
		bdata = &GlobalBdata
	}
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &agent_name, 1) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &return_array, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
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
		var http_user_agent *zend.Zval = nil
		if zend.Z_TYPE(core.PG(http_globals)[core.TRACK_VARS_SERVER]) == zend.IS_ARRAY || zend.ZendIsAutoGlobalStr(zend.ZEND_STRL("_SERVER")) != 0 {
			http_user_agent = zend.ZendHashStrFind(zend.Z_ARRVAL_P(&core.PG(http_globals)[core.TRACK_VARS_SERVER]), "HTTP_USER_AGENT", b.SizeOf("\"HTTP_USER_AGENT\"")-1)
		}
		if http_user_agent == nil {
			core.PhpErrorDocref(nil, zend.E_WARNING, "HTTP_USER_AGENT variable is not set, cannot determine user agent name")
			zend.RETVAL_FALSE
			return
		}
		agent_name = zend.Z_STR_P(http_user_agent)
	}
	lookup_browser_name = zend.ZendStringTolower(agent_name)
	found_entry = zend.ZendHashFindPtr(bdata.GetHtab(), lookup_browser_name)
	if found_entry == nil {
		var entry *BrowscapEntry
		for {
			var __ht *zend.HashTable = bdata.GetHtab()
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				entry = zend.Z_PTR_P(_z)
				if BrowserRegCompare(entry, lookup_browser_name, &found_entry) != 0 {
					break
				}
			}
			break
		}
		if found_entry == nil {
			found_entry = zend.ZendHashStrFindPtr(bdata.GetHtab(), DEFAULT_SECTION_NAME, b.SizeOf("DEFAULT_SECTION_NAME")-1)
			if found_entry == nil {
				zend.ZendStringRelease(lookup_browser_name)
				zend.RETVAL_FALSE
				return
			}
		}
	}
	agent_ht = BrowscapEntryToArray(bdata, found_entry)
	if return_array != 0 {
		zend.RETVAL_ARR(agent_ht)
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
			zend.ZendHashMerge(zend.Z_ARRVAL_P(return_value), agent_ht, zend.CopyCtorFuncT(BrowscapZvalCopyCtor), 0)
		} else {
			zend.ZendHashMerge(zend.Z_OBJPROP_P(return_value), agent_ht, zend.CopyCtorFuncT(BrowscapZvalCopyCtor), 0)
		}
		zend.ZendHashDestroy(agent_ht)
		zend.Efree(agent_ht)
	}
	zend.ZendStringReleaseEx(lookup_browser_name, 0)
}

/* }}} */
