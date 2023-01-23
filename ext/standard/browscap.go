// <<generate>>

package standard

import (
	"sik/core"
	r "sik/runtime"
	g "sik/runtime/grammar"
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

// #define BROWSCAP_NUM_CONTAINS       5

/* browser data defined in startup phase, eagerly loaded in MINIT */

var GlobalBdata BrowserData = BrowserData{0}

/* browser data defined in activation phase, lazily loaded in get_browser.
 * Per request and per thread, if applicable */

var BrowscapGlobals ZendBrowscapGlobals

// #define BROWSCAP_G(v) ZEND_MODULE_GLOBALS_ACCESSOR ( browscap , v )

// #define DEFAULT_SECTION_NAME       "Default Browser Capability Settings"

/* OBJECTS_FIXME: This whole extension needs going through. The use of objects looks pretty broken here */

func BrowscapEntryDtor(zvalue *zend.Zval) {
	var entry *BrowscapEntry = zvalue.value.ptr
	zend.ZendStringReleaseEx(entry.GetPattern(), 0)
	if entry.GetParent() != nil {
		zend.ZendStringReleaseEx(entry.GetParent(), 0)
	}
	zend._efree(entry)
}
func BrowscapEntryDtorPersistent(zvalue *zend.Zval) {
	var entry *BrowscapEntry = zvalue.value.ptr
	zend.ZendStringReleaseEx(entry.GetPattern(), 1)
	if entry.GetParent() != nil {
		zend.ZendStringReleaseEx(entry.GetParent(), 1)
	}
	g.CondF(true, func() { return zend.Free(entry) }, func() { return zend._efree(entry) })
}
func IsPlaceholder(c byte) zend.ZendBool { return c == '?' || c == '*' }

/* Length of prefix not containing any wildcards */

func BrowscapComputePrefixLen(pattern *zend.ZendString) uint8 {
	var i int
	for i = 0; i < pattern.len_; i++ {
		if IsPlaceholder(pattern.val[i]) != 0 {
			break
		}
	}
	return uint8_t(g.Cond(i < UINT8_MAX, i, UINT8_MAX))
}
func BrowscapComputeContains(pattern *zend.ZendString, start_pos int, contains_start *uint16, contains_len *uint8) int {
	var i int = start_pos

	/* Find first non-placeholder character after prefix */

	for ; i < pattern.len_; i++ {
		if IsPlaceholder(pattern.val[i]) == 0 {

			/* Skip the case of a single non-placeholder character.
			 * Let's try to find something longer instead. */

			if i+1 < pattern.len_ && IsPlaceholder(pattern.val[i+1]) == 0 {
				break
			}

			/* Skip the case of a single non-placeholder character.
			 * Let's try to find something longer instead. */

		}
	}
	*contains_start = uint16(i)

	/* Find first placeholder character after that */

	for ; i < pattern.len_; i++ {
		if IsPlaceholder(pattern.val[i]) != 0 {
			break
		}
	}
	*contains_len = uint8_t(g.Cond(i-(*contains_start) < UINT8_MAX, i-(*contains_start), UINT8_MAX))
	return i
}

/* Length of regex, including escapes, anchors, etc. */

func BrowscapComputeRegexLen(pattern *zend.ZendString) int {
	var i int
	var len_ int = pattern.len_
	for i = 0; i < pattern.len_; i++ {
		switch pattern.val[i] {
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
	return len_ + g.SizeOf("\"~^$~\"") - 1
}
func BrowscapConvertPattern(pattern *zend.ZendString, persistent int) *zend.ZendString {
	var i int
	var j int = 0
	var t *byte
	var res *zend.ZendString
	var lc_pattern *byte
	res = zend.ZendStringAlloc(BrowscapComputeRegexLen(pattern), persistent)
	t = res.val
	lc_pattern = zend._emalloc(pattern.len_ + 1)
	zend.ZendStrTolowerCopy(lc_pattern, pattern.val, pattern.len_)
	t[g.PostInc(&j)] = '~'
	t[g.PostInc(&j)] = '^'
	for i = 0; i < pattern.len_; {
		switch lc_pattern[i] {
		case '?':
			t[j] = '.'
			break
		case '*':
			t[g.PostInc(&j)] = '.'
			t[j] = '*'
			break
		case '.':
			t[g.PostInc(&j)] = '\\'
			t[j] = '.'
			break
		case '\\':
			t[g.PostInc(&j)] = '\\'
			t[j] = '\\'
			break
		case '(':
			t[g.PostInc(&j)] = '\\'
			t[j] = '('
			break
		case ')':
			t[g.PostInc(&j)] = '\\'
			t[j] = ')'
			break
		case '~':
			t[g.PostInc(&j)] = '\\'
			t[j] = '~'
			break
		case '+':
			t[g.PostInc(&j)] = '\\'
			t[j] = '+'
			break
		default:
			t[j] = lc_pattern[i]
			break
		}
		i++
		j++
	}
	t[g.PostInc(&j)] = '$'
	t[g.PostInc(&j)] = '~'
	t[j] = 0
	res.len_ = j
	zend._efree(lc_pattern)
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
	lcname = (*zend.ZendString)(zend._emalloc(zend_long((*byte)(&((*zend.ZendString)(nil).val))-(*byte)(nil)) + str.len_ + 1 + (8-1) & ^(8-1)))
	zend.ZendGcSetRefcount(&lcname.gc, 1)
	lcname.gc.u.type_info = 6
	lcname.h = 0
	lcname.len_ = str.len_
	zend.ZendStrTolowerCopy(lcname.val, str.val, str.len_)
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
	zend._efree(lcname)
	return interned
}
func BrowscapAddKv(bdata *BrowserData, key *zend.ZendString, value *zend.ZendString, persistent zend.ZendBool) {
	if bdata.GetKvUsed() == bdata.GetKvSize() {
		bdata.SetKvSize(bdata.GetKvSize() * 2)
		if persistent != 0 {
			bdata.SetKv(zend._safeRealloc(bdata.GetKv(), g.SizeOf("browscap_kv"), bdata.GetKvSize(), 0))
		} else {
			bdata.SetKv(zend._safeErealloc(bdata.GetKv(), g.SizeOf("browscap_kv"), bdata.GetKvSize(), 0))
		}
	}
	bdata.GetKv()[bdata.GetKvUsed()].SetKey(key)
	bdata.GetKv()[bdata.GetKvUsed()].SetValue(value)
	bdata.GetKvUsed()++
}
func BrowscapEntryToArray(bdata *BrowserData, entry *BrowscapEntry) *zend.HashTable {
	var tmp zend.Zval
	var i uint32
	var ht *zend.HashTable = zend._zendNewArray(8)
	var __z *zend.Zval = &tmp
	var __s *zend.ZendString = BrowscapConvertPattern(entry.GetPattern(), 0)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	zend.ZendHashStrAdd(ht, "browser_name_regex", g.SizeOf("\"browser_name_regex\"")-1, &tmp)
	var __z *zend.Zval = &tmp
	var __s *zend.ZendString = entry.GetPattern()
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		zend.ZendGcAddref(&__s.gc)
		__z.u1.type_info = 6 | 1<<0<<8
	}
	zend.ZendHashStrAdd(ht, "browser_name_pattern", g.SizeOf("\"browser_name_pattern\"")-1, &tmp)
	if entry.GetParent() != nil {
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = entry.GetParent()
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		zend.ZendHashStrAdd(ht, "parent", g.SizeOf("\"parent\"")-1, &tmp)
	}
	for i = entry.GetKvStart(); i < entry.GetKvEnd(); i++ {
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = bdata.GetKv()[i].GetValue()
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		zend.ZendHashAdd(ht, bdata.GetKv()[i].GetKey(), &tmp)
	}
	return ht
}
func PhpBrowscapParserCb(arg1 *zend.Zval, arg2 *zend.Zval, arg3 *zend.Zval, callback_type int, arg any) {
	var ctx *BrowscapParserCtx = arg
	var bdata *BrowserData = ctx.GetBdata()
	var persistent int = zend.ZvalGcFlags(bdata.GetHtab().gc.u.type_info) & 1 << 7
	if arg1 == nil {
		return
	}
	switch callback_type {
	case 1:
		if ctx.GetCurrentEntry() != nil && arg2 != nil {
			var new_key *zend.ZendString
			var new_value *zend.ZendString

			/* Set proper value for true/false settings */

			if arg2.value.str.len_ == 2 && !(strncasecmp(arg2.value.str.val, "on", g.SizeOf("\"on\"")-1)) || arg2.value.str.len_ == 3 && !(strncasecmp(arg2.value.str.val, "yes", g.SizeOf("\"yes\"")-1)) || arg2.value.str.len_ == 4 && !(strncasecmp(arg2.value.str.val, "true", g.SizeOf("\"true\"")-1)) {
				new_value = zend.ZendOneCharString['1']
			} else if arg2.value.str.len_ == 2 && !(strncasecmp(arg2.value.str.val, "no", g.SizeOf("\"no\"")-1)) || arg2.value.str.len_ == 3 && !(strncasecmp(arg2.value.str.val, "off", g.SizeOf("\"off\"")-1)) || arg2.value.str.len_ == 4 && !(strncasecmp(arg2.value.str.val, "none", g.SizeOf("\"none\"")-1)) || arg2.value.str.len_ == 5 && !(strncasecmp(arg2.value.str.val, "false", g.SizeOf("\"false\"")-1)) {
				new_value = zend.ZendEmptyString
			} else {
				new_value = BrowscapInternStr(ctx, arg2.value.str, persistent)
			}
			if !(strcasecmp(arg1.value.str.val, "parent")) {

				/* parent entry can not be same as current section -> causes infinite loop! */

				if ctx.GetCurrentSectionName() != nil && !(strcasecmp(ctx.GetCurrentSectionName().val, arg2.value.str.val)) {
					zend.ZendError(1<<4, "Invalid browscap ini file: "+"'Parent' value cannot be same as the section name: %s "+"(in file %s)", ctx.GetCurrentSectionName().val, zend.ZendIniStringEx("browscap", g.SizeOf("\"browscap\"")-1, 0, nil))
					return
				}
				if ctx.GetCurrentEntry().GetParent() != nil {
					zend.ZendStringRelease(ctx.GetCurrentEntry().GetParent())
				}
				ctx.GetCurrentEntry().SetParent(new_value)
			} else {
				new_key = BrowscapInternStrCi(ctx, arg1.value.str, persistent)
				BrowscapAddKv(bdata, new_key, new_value, persistent)
				ctx.GetCurrentEntry().SetKvEnd(bdata.GetKvUsed())
			}
		}
		break
	case 2:
		var entry *BrowscapEntry
		var pattern *zend.ZendString = arg1.value.str
		var pos int
		var i int
		if pattern.len_ > UINT16_MAX {
			core.PhpErrorDocref(nil, 1<<1, "Skipping excessively long pattern of length %zd", pattern.len_)
			break
		}
		if persistent != 0 {
			pattern = zend.ZendNewInternedString(zend.ZendStringCopy(pattern))
			if (zend.ZvalGcFlags(pattern.gc.u.type_info) & 1 << 6) != 0 {
				arg1.u1.v.type_flags = 0
			} else {
				zend.ZendStringRelease(pattern)
			}
		}
		if persistent != 0 {
			ctx.SetCurrentEntry(zend.__zendMalloc(g.SizeOf("browscap_entry")))
		} else {
			ctx.SetCurrentEntry(zend._emalloc(g.SizeOf("browscap_entry")))
		}
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
		for i = 0; i < 5; i++ {
			pos = BrowscapComputeContains(pattern, pos, &entry.contains_start[i], &entry.contains_len[i])
		}
		break
	}
}

/* }}} */

func StrInternedDtor(zv *zend.Zval) { zend.ZendStringRelease(zv.value.str) }
func BrowscapReadFile(filename *byte, browdata *BrowserData, persistent int) int {
	var fh zend.ZendFileHandle
	var ctx BrowscapParserCtx = BrowscapParserCtx{0}
	if filename == nil || filename[0] == '0' {
		return zend.FAILURE
	}
	zend.ZendStreamInitFp(&fh, r.Fopen(filename, "r"), filename)
	if fh.handle.fp == nil {
		zend.ZendError(1<<5, "Cannot open '%s' for reading", filename)
		return zend.FAILURE
	}
	if persistent != 0 {
		browdata.SetHtab(zend.__zendMalloc(sizeof * browdata.GetHtab()))
	} else {
		browdata.SetHtab(zend._emalloc(sizeof * browdata.GetHtab()))
	}
	zend._zendHashInit(browdata.GetHtab(), 0, g.Cond(persistent != 0, BrowscapEntryDtorPersistent, BrowscapEntryDtor), persistent)
	browdata.SetKvSize(16 * 1024)
	browdata.SetKvUsed(0)
	if persistent != 0 {
		browdata.SetKv(zend.__zendMalloc(g.SizeOf("browscap_kv") * browdata.GetKvSize()))
	} else {
		browdata.SetKv(zend._emalloc(g.SizeOf("browscap_kv") * browdata.GetKvSize()))
	}

	/* Create parser context */

	ctx.SetBdata(browdata)
	ctx.SetCurrentEntry(nil)
	ctx.SetCurrentSectionName(nil)
	zend._zendHashInit(&ctx.str_interned, 8, StrInternedDtor, persistent)
	zend.ZendParseIniFile(&fh, 1, 1, zend.ZendIniParserCbT(PhpBrowscapParserCb), &ctx)

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
		g.CondF(persistent != 0, func() { return zend.Free(bdata.GetHtab()) }, func() { return zend._efree(bdata.GetHtab()) })
		bdata.SetHtab(nil)
		for i = 0; i < bdata.GetKvUsed(); i++ {
			zend.ZendStringRelease(bdata.GetKv()[i].GetKey())
			zend.ZendStringRelease(bdata.GetKv()[i].GetValue())
		}
		g.CondF(persistent != 0, func() { return zend.Free(bdata.GetKv()) }, func() { return zend._efree(bdata.GetKv()) })
		bdata.SetKv(nil)
	}
	bdata.GetFilename()[0] = '0'
}

/* }}} */

func OnChangeBrowscap(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	if stage == 1<<0 {

		/* value handled in browscap.c's MINIT */

		return zend.SUCCESS

		/* value handled in browscap.c's MINIT */

	} else if stage == 1<<2 {
		var bdata *BrowserData = &(BrowscapGlobals.GetActivationBdata())
		if bdata.GetFilename()[0] != '0' {
			BrowscapBdataDtor(bdata, 0)
		}
		if zend.TsrmRealpath(new_value.val, bdata.GetFilename()) == nil {
			return zend.FAILURE
		}
		return zend.SUCCESS
	}
	return zend.FAILURE
}

/* }}} */

func ZmStartupBrowscap(type_ int, module_number int) int {
	var browscap *byte = zend.ZendIniStringEx("browscap", g.SizeOf("\"browscap\"")-1, 0, nil)

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
	var bdata *BrowserData = &(BrowscapGlobals.GetActivationBdata())
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
	for i = 0; i < 5; i++ {
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

	if agent_name.len_ < BrowscapGetMinimumLength(entry) {
		return 0
	}

	/* Quickly discard patterns where the prefix doesn't match. */

	if zend.ZendBinaryStrcasecmp(agent_name.val, entry.GetPrefixLen(), entry.GetPattern().val, entry.GetPrefixLen()) != 0 {
		return 0
	}

	/* Lowercase the pattern, the agent name is already lowercase */

	pattern_lc = (*zend.ZendString)(zend._emalloc(zend_long((*byte)(&((*zend.ZendString)(nil).val))-(*byte)(nil)) + entry.GetPattern().len_ + 1 + (8-1) & ^(8-1)))
	zend.ZendGcSetRefcount(&pattern_lc.gc, 1)
	pattern_lc.gc.u.type_info = 6
	pattern_lc.h = 0
	pattern_lc.len_ = entry.GetPattern().len_
	zend.ZendStrTolowerCopy(pattern_lc.val, entry.GetPattern().val, entry.GetPattern().len_)

	/* Check if the agent contains the "contains" portions */

	cur = agent_name.val + entry.GetPrefixLen()
	for i = 0; i < 5; i++ {
		if entry.GetContainsLen()[i] != 0 {
			cur = zend.ZendMemnstr(cur, pattern_lc.val+entry.GetContainsStart()[i], entry.GetContainsLen()[i], agent_name.val+agent_name.len_)
			if cur == nil {
				zend._efree(pattern_lc)
				return 0
			}
			cur += entry.GetContainsLen()[i]
		}
	}

	/* See if we have an exact match, if so, we're done... */

	if zend.ZendStringEquals(agent_name, pattern_lc) != 0 {
		*found_entry_ptr = entry
		zend._efree(pattern_lc)
		return 1
	}
	regex = BrowscapConvertPattern(entry.GetPattern(), 0)
	re = pcre_get_compiled_regex(regex, &capture_count)
	if re == nil {
		zend._efree(pattern_lc)
		zend.ZendStringRelease(regex)
		return 0
	}
	match_data = php_pcre_create_match_data(capture_count, re)
	if match_data == nil {
		zend._efree(pattern_lc)
		zend.ZendStringRelease(regex)
		return 0
	}
	rc = pcre2_match(re, PCRE2_SPTR(agent_name).val, agent_name.len_, 0, 0, match_data, php_pcre_mctx())
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
			for i = 0; i < previous_match.len_; i++ {
				switch previous_match.val[i] {
				case '?':

				case '*':

					/* do nothing, ignore these characters in the count */

					break
				default:
					prev_len++
				}
			}
			for i = 0; i < current_match.len_; i++ {
				switch current_match.val[i] {
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
	zend._efree(pattern_lc)
	zend.ZendStringRelease(regex)
	return 0
}

/* }}} */

func BrowscapZvalCopyCtor(p *zend.Zval) {
	if p.u1.v.type_flags != 0 {
		var str *zend.ZendString
		r.Assert(p.u1.v.type_ == 6)
		str = p.value.str
		if (zend.ZvalGcFlags(str.gc.u.type_info) & 1 << 7) == 0 {
			zend.ZendGcAddref(&str.gc)
		} else {
			var __z *zend.Zval = p
			var __s *zend.ZendString = zend.ZendStringInit(str.val, str.len_, 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
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
	if BrowscapGlobals.GetActivationBdata().GetFilename()[0] != '0' {
		bdata = &(BrowscapGlobals.GetActivationBdata())
		if bdata.GetHtab() == nil {
			if BrowscapReadFile(bdata.GetFilename(), bdata, 0) == zend.FAILURE {
				return_value.u1.type_info = 2
				return
			}
		}
	} else {
		if GlobalBdata.GetHtab() == nil {
			core.PhpErrorDocref(nil, 1<<1, "browscap ini directive not set")
			return_value.u1.type_info = 2
			return
		}
		bdata = &GlobalBdata
	}
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_optional = 1
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &agent_name, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &return_array, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
		if core.CoreGlobals.http_globals[3].u1.v.type_ == 7 || zend.ZendIsAutoGlobalStr("_SERVER", g.SizeOf("\"_SERVER\"")-1) != 0 {
			http_user_agent = zend.ZendHashStrFind(&core.CoreGlobals.http_globals[3].value.arr, "HTTP_USER_AGENT", g.SizeOf("\"HTTP_USER_AGENT\"")-1)
		}
		if http_user_agent == nil {
			core.PhpErrorDocref(nil, 1<<1, "HTTP_USER_AGENT variable is not set, cannot determine user agent name")
			return_value.u1.type_info = 2
			return
		}
		agent_name = http_user_agent.value.str
	}
	lookup_browser_name = zend.ZendStringTolowerEx(agent_name, 0)
	found_entry = zend.ZendHashFindPtr(bdata.GetHtab(), lookup_browser_name)
	if found_entry == nil {
		var entry *BrowscapEntry
		for {
			var __ht *zend.HashTable = bdata.GetHtab()
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				entry = _z.value.ptr
				if BrowserRegCompare(entry, lookup_browser_name, &found_entry) != 0 {
					break
				}
			}
			break
		}
		if found_entry == nil {
			found_entry = zend.ZendHashStrFindPtr(bdata.GetHtab(), "Default Browser Capability Settings", g.SizeOf("DEFAULT_SECTION_NAME")-1)
			if found_entry == nil {
				zend.ZendStringRelease(lookup_browser_name)
				return_value.u1.type_info = 2
				return
			}
		}
	}
	agent_ht = BrowscapEntryToArray(bdata, found_entry)
	if return_array != 0 {
		var __arr *zend.ZendArray = agent_ht
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
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
			zend.ZendHashMerge(return_value.value.arr, agent_ht, zend.CopyCtorFuncT(BrowscapZvalCopyCtor), 0)
		} else {
			zend.ZendHashMerge(return_value.value.obj.handlers.get_properties(&(*return_value)), agent_ht, zend.CopyCtorFuncT(BrowscapZvalCopyCtor), 0)
		}
		zend.ZendHashDestroy(agent_ht)
		zend._efree(agent_ht)
	}
	zend.ZendStringReleaseEx(lookup_browser_name, 0)
}

/* }}} */
