// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_string.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define ZEND_STRING_H

// # include "zend.h"

type ZendStringCopyStorageFuncT func()
type ZendNewInternedStringFuncT func(str *ZendString) *ZendString
type ZendStringInitInternedFuncT func(str *byte, size int, permanent int) *ZendString

var ZendNewInternedString ZendNewInternedStringFuncT
var ZendStringInitInterned ZendStringInitInternedFuncT
var ZendOneCharString []*ZendString

/* Shortcuts */

func ZSTR_VAL(zstr *ZendString) []byte     { return zstr.GetVal() }
func ZSTR_LEN(zstr *ZendString) int        { return zstr.GetLen() }
func ZSTR_H(zstr *ZendString) ZendUlong    { return zstr.GetH() }
func ZSTR_HASH(zstr *ZendString) ZendUlong { return ZendStringHashVal(zstr) }

/* Compatibility macros */

func IS_INTERNED(s *ZendString) int { return ZSTR_IS_INTERNED(s) }
func STR_EMPTY_ALLOC() *ZendString  { return ZSTR_EMPTY_ALLOC() }

const _STR_HEADER_SIZE = _ZSTR_HEADER_SIZE

func STR_ALLOCA_ALLOC(str *ZendString, _len int, use_heap __auto__) {
	ZSTR_ALLOCA_ALLOC(str, _len, use_heap)
}
func STR_ALLOCA_INIT(str *ZendString, s __auto__, len_ int, use_heap __auto__) {
	ZSTR_ALLOCA_INIT(str, s, len_, use_heap)
}
func STR_ALLOCA_FREE(str any, use_heap __auto__) { ZSTR_ALLOCA_FREE(str, use_heap) }

/*---*/

func ZSTR_IS_INTERNED(s *ZendString) int           { return GC_FLAGS(s) & IS_STR_INTERNED }
func ZSTR_EMPTY_ALLOC() *ZendString                { return ZendEmptyString }
func ZSTR_CHAR(c int) *ZendString                  { return ZendOneCharString[c] }
func ZSTR_KNOWN(idx ZendKnownStringId) *ZendString { return ZendKnownStrings[idx] }

const _ZSTR_HEADER_SIZE = zend_long((*byte)(&((*ZendString)(nil).GetVal())) - (*byte)(nil))

func _ZSTR_STRUCT_SIZE(len_ int) int { return _ZSTR_HEADER_SIZE + len_ + 1 }

const ZSTR_MAX_OVERHEAD = ZEND_MM_ALIGNED_SIZE(_ZSTR_HEADER_SIZE + 1)
const ZSTR_MAX_LEN = SIZE_MAX - ZSTR_MAX_OVERHEAD

func ZSTR_ALLOCA_ALLOC(str *ZendString, _len int, use_heap __auto__) {
	str = (*ZendString)(DoAlloca(ZEND_MM_ALIGNED_SIZE_EX(_ZSTR_STRUCT_SIZE(_len), 8), use_heap))
	GC_SET_REFCOUNT(str, 1)
	GC_TYPE_INFO(str) = IS_STRING
	ZSTR_H(str) = 0
	ZSTR_LEN(str) = _len
}
func ZSTR_ALLOCA_INIT(str *ZendString, s __auto__, len_ int, use_heap __auto__) {
	ZSTR_ALLOCA_ALLOC(str, len_, use_heap)
	memcpy(ZSTR_VAL(str), s, len_)
	ZSTR_VAL(str)[len_] = '0'
}
func ZSTR_ALLOCA_FREE(str any, use_heap __auto__) { FreeAlloca(str, use_heap) }

/*---*/

func ZendStringHashVal(s *ZendString) ZendUlong {
	if ZSTR_H(s) != 0 {
		return ZSTR_H(s)
	} else {
		return ZendStringHashFunc(s)
	}
}
func ZendStringForgetHashVal(s *ZendString) {
	ZSTR_H(s) = 0
	GC_DEL_FLAGS(s, IS_STR_VALID_UTF8)
}
func ZendStringRefcount(s *ZendString) uint32 {
	if ZSTR_IS_INTERNED(s) == 0 {
		return GC_REFCOUNT(s)
	}
	return 1
}
func ZendStringAddref(s *ZendString) uint32 {
	if ZSTR_IS_INTERNED(s) == 0 {
		return GC_ADDREF(s)
	}
	return 1
}
func ZendStringDelref(s *ZendString) uint32 {
	if ZSTR_IS_INTERNED(s) == 0 {
		return GC_DELREF(s)
	}
	return 1
}
func ZendStringAlloc(len_ int, persistent int) *ZendString {
	var ret *ZendString = (*ZendString)(Pemalloc(ZEND_MM_ALIGNED_SIZE(_ZSTR_STRUCT_SIZE(len_)), persistent))
	GC_SET_REFCOUNT(ret, 1)
	GC_TYPE_INFO(ret) = IS_STRING | b.Cond(persistent != 0, IS_STR_PERSISTENT, 0)<<GC_FLAGS_SHIFT
	ZSTR_H(ret) = 0
	ZSTR_LEN(ret) = len_
	return ret
}
func ZendStringSafeAlloc(n int, m int, l int, persistent int) *ZendString {
	var ret *ZendString = (*ZendString)(SafePemalloc(n, m, ZEND_MM_ALIGNED_SIZE(_ZSTR_STRUCT_SIZE(l)), persistent))
	GC_SET_REFCOUNT(ret, 1)
	GC_TYPE_INFO(ret) = IS_STRING | b.Cond(persistent != 0, IS_STR_PERSISTENT, 0)<<GC_FLAGS_SHIFT
	ZSTR_H(ret) = 0
	ZSTR_LEN(ret) = n*m + l
	return ret
}
func ZendStringInit(str *byte, len_ int, persistent int) *ZendString {
	var ret *ZendString = ZendStringAlloc(len_, persistent)
	memcpy(ZSTR_VAL(ret), str, len_)
	ZSTR_VAL(ret)[len_] = '0'
	return ret
}
func ZendStringCopy(s *ZendString) *ZendString {
	if ZSTR_IS_INTERNED(s) == 0 {
		GC_ADDREF(s)
	}
	return s
}
func ZendStringDup(s *ZendString, persistent int) *ZendString {
	if ZSTR_IS_INTERNED(s) != 0 {
		return s
	} else {
		return ZendStringInit(ZSTR_VAL(s), ZSTR_LEN(s), persistent)
	}
}
func ZendStringRealloc(s *ZendString, len_ int, persistent int) *ZendString {
	var ret *ZendString
	if ZSTR_IS_INTERNED(s) == 0 {
		if EXPECTED(GC_REFCOUNT(s) == 1) {
			ret = (*ZendString)(Perealloc(s, ZEND_MM_ALIGNED_SIZE(_ZSTR_STRUCT_SIZE(len_)), persistent))
			ZSTR_LEN(ret) = len_
			ZendStringForgetHashVal(ret)
			return ret
		}
	}
	ret = ZendStringAlloc(len_, persistent)
	memcpy(ZSTR_VAL(ret), ZSTR_VAL(s), MIN(len_, ZSTR_LEN(s))+1)
	if ZSTR_IS_INTERNED(s) == 0 {
		GC_DELREF(s)
	}
	return ret
}
func ZendStringExtend(s *ZendString, len_ int, persistent int) *ZendString {
	var ret *ZendString
	ZEND_ASSERT(len_ >= ZSTR_LEN(s))
	if ZSTR_IS_INTERNED(s) == 0 {
		if EXPECTED(GC_REFCOUNT(s) == 1) {
			ret = (*ZendString)(Perealloc(s, ZEND_MM_ALIGNED_SIZE(_ZSTR_STRUCT_SIZE(len_)), persistent))
			ZSTR_LEN(ret) = len_
			ZendStringForgetHashVal(ret)
			return ret
		}
	}
	ret = ZendStringAlloc(len_, persistent)
	memcpy(ZSTR_VAL(ret), ZSTR_VAL(s), ZSTR_LEN(s)+1)
	if ZSTR_IS_INTERNED(s) == 0 {
		GC_DELREF(s)
	}
	return ret
}
func ZendStringTruncate(s *ZendString, len_ int, persistent int) *ZendString {
	var ret *ZendString
	ZEND_ASSERT(len_ <= ZSTR_LEN(s))
	if ZSTR_IS_INTERNED(s) == 0 {
		if EXPECTED(GC_REFCOUNT(s) == 1) {
			ret = (*ZendString)(Perealloc(s, ZEND_MM_ALIGNED_SIZE(_ZSTR_STRUCT_SIZE(len_)), persistent))
			ZSTR_LEN(ret) = len_
			ZendStringForgetHashVal(ret)
			return ret
		}
	}
	ret = ZendStringAlloc(len_, persistent)
	memcpy(ZSTR_VAL(ret), ZSTR_VAL(s), len_+1)
	if ZSTR_IS_INTERNED(s) == 0 {
		GC_DELREF(s)
	}
	return ret
}
func ZendStringSafeRealloc(s *ZendString, n int, m int, l int, persistent int) *ZendString {
	var ret *ZendString
	if ZSTR_IS_INTERNED(s) == 0 {
		if GC_REFCOUNT(s) == 1 {
			ret = (*ZendString)(SafePerealloc(s, n, m, ZEND_MM_ALIGNED_SIZE(_ZSTR_STRUCT_SIZE(l)), persistent))
			ZSTR_LEN(ret) = n*m + l
			ZendStringForgetHashVal(ret)
			return ret
		}
	}
	ret = ZendStringSafeAlloc(n, m, l, persistent)
	memcpy(ZSTR_VAL(ret), ZSTR_VAL(s), MIN(n*m+l, ZSTR_LEN(s))+1)
	if ZSTR_IS_INTERNED(s) == 0 {
		GC_DELREF(s)
	}
	return ret
}
func ZendStringFree(s *ZendString) {
	if ZSTR_IS_INTERNED(s) == 0 {
		ZEND_ASSERT(GC_REFCOUNT(s) <= 1)
		Pefree(s, GC_FLAGS(s)&IS_STR_PERSISTENT)
	}
}
func ZendStringEfree(s *ZendString) {
	ZEND_ASSERT(ZSTR_IS_INTERNED(s) == 0)
	ZEND_ASSERT(GC_REFCOUNT(s) <= 1)
	ZEND_ASSERT((GC_FLAGS(s) & IS_STR_PERSISTENT) == 0)
	Efree(s)
}
func ZendStringRelease(s *ZendString) {
	if ZSTR_IS_INTERNED(s) == 0 {
		if GC_DELREF(s) == 0 {
			Pefree(s, GC_FLAGS(s)&IS_STR_PERSISTENT)
		}
	}
}
func ZendStringReleaseEx(s *ZendString, persistent int) {
	if ZSTR_IS_INTERNED(s) == 0 {
		if GC_DELREF(s) == 0 {
			if persistent != 0 {
				ZEND_ASSERT((GC_FLAGS(s) & IS_STR_PERSISTENT) != 0)
				Free(s)
			} else {
				ZEND_ASSERT((GC_FLAGS(s) & IS_STR_PERSISTENT) == 0)
				Efree(s)
			}
		}
	}
}
func ZendStringEqualVal(s1 *ZendString, s2 *ZendString) ZendBool {
	return !(memcmp(ZSTR_VAL(s1), ZSTR_VAL(s2), ZSTR_LEN(s1)))
}
func ZendStringEqualContent(s1 *ZendString, s2 *ZendString) ZendBool {
	return ZSTR_LEN(s1) == ZSTR_LEN(s2) && ZendStringEqualVal(s1, s2) != 0
}
func ZendStringEquals(s1 *ZendString, s2 *ZendString) ZendBool {
	return s1 == s2 || ZendStringEqualContent(s1, s2) != 0
}
func ZendStringEqualsCi(s1 *ZendString, s2 *ZendString) bool {
	return ZSTR_LEN(s1) == ZSTR_LEN(s2) && ZendBinaryStrcasecmp(ZSTR_VAL(s1), ZSTR_LEN(s1), ZSTR_VAL(s2), ZSTR_LEN(s2)) == 0
}
func ZendStringEqualsLiteralCi(str *ZendString, c string) bool {
	return ZSTR_LEN(str) == b.SizeOf("c")-1 && ZendBinaryStrcasecmp(ZSTR_VAL(str), ZSTR_LEN(str), c, b.SizeOf("c")-1) == 0
}
func ZendStringEqualsLiteral(str *ZendString, literal string) bool {
	return ZSTR_LEN(str) == b.SizeOf("literal")-1 && !(memcmp(ZSTR_VAL(str), literal, b.SizeOf("literal")-1))
}

/*
 * DJBX33A (Daniel J. Bernstein, Times 33 with Addition)
 *
 * This is Daniel J. Bernstein's popular `times 33' hash function as
 * posted by him years ago on comp.lang.c. It basically uses a function
 * like ``hash(i) = hash(i-1) * 33 + str[i]''. This is one of the best
 * known hash functions for strings. Because it is both computed very
 * fast and distributes very well.
 *
 * The magic of number 33, i.e. why it works better than many other
 * constants, prime or not, has never been adequately explained by
 * anyone. So I try an explanation: if one experimentally tests all
 * multipliers between 1 and 256 (as RSE did now) one detects that even
 * numbers are not usable at all. The remaining 128 odd numbers
 * (except for the number 1) work more or less all equally well. They
 * all distribute in an acceptable way and this way fill a hash table
 * with an average percent of approx. 86%.
 *
 * If one compares the Chi^2 values of the variants, the number 33 not
 * even has the best value. But the number 33 and a few other equally
 * good numbers like 17, 31, 63, 127 and 129 have nevertheless a great
 * advantage to the remaining numbers in the large set of possible
 * multipliers: their multiply operation can be replaced by a faster
 * operation based on just one shift plus either a single addition
 * or subtraction operation. And because a hash function has to both
 * distribute good _and_ has to be very fast to compute, those few
 * numbers should be preferred and seems to be the reason why Daniel J.
 * Bernstein also preferred it.
 *
 *
 *                  -- Ralf S. Engelschall <rse@engelschall.com>
 */

func ZendInlineHashFunc(str *byte, len_ int) ZendUlong {
	var hash ZendUlong = Z_UL(5381)

	/* variant with the hash unrolled eight times */

	for ; len_ >= 8; len_ -= 8 {
		hash = (hash << 5) + hash + b.PostInc(&(*str))
		hash = (hash << 5) + hash + b.PostInc(&(*str))
		hash = (hash << 5) + hash + b.PostInc(&(*str))
		hash = (hash << 5) + hash + b.PostInc(&(*str))
		hash = (hash << 5) + hash + b.PostInc(&(*str))
		hash = (hash << 5) + hash + b.PostInc(&(*str))
		hash = (hash << 5) + hash + b.PostInc(&(*str))
		hash = (hash << 5) + hash + b.PostInc(&(*str))
	}
	switch len_ {
	case 7:
		hash = (hash << 5) + hash + b.PostInc(&(*str))
	case 6:
		hash = (hash << 5) + hash + b.PostInc(&(*str))
	case 5:
		hash = (hash << 5) + hash + b.PostInc(&(*str))
	case 4:
		hash = (hash << 5) + hash + b.PostInc(&(*str))
	case 3:
		hash = (hash << 5) + hash + b.PostInc(&(*str))
	case 2:
		hash = (hash << 5) + hash + b.PostInc(&(*str))
	case 1:
		hash = (hash << 5) + hash + b.PostInc(&(*str))
		break
	case 0:
		break
	default:
		break
	}

	/* Hash value can't be zero, so we always set the high bit */

	return hash | Z_UL(-0x8000000000000000)

	/* Hash value can't be zero, so we always set the high bit */
}

// #define ZEND_KNOWN_STRINGS(_) _ ( ZEND_STR_FILE , "file" ) _ ( ZEND_STR_LINE , "line" ) _ ( ZEND_STR_FUNCTION , "function" ) _ ( ZEND_STR_CLASS , "class" ) _ ( ZEND_STR_OBJECT , "object" ) _ ( ZEND_STR_TYPE , "type" ) _ ( ZEND_STR_OBJECT_OPERATOR , "->" ) _ ( ZEND_STR_PAAMAYIM_NEKUDOTAYIM , "::" ) _ ( ZEND_STR_ARGS , "args" ) _ ( ZEND_STR_UNKNOWN , "unknown" ) _ ( ZEND_STR_EVAL , "eval" ) _ ( ZEND_STR_INCLUDE , "include" ) _ ( ZEND_STR_REQUIRE , "require" ) _ ( ZEND_STR_INCLUDE_ONCE , "include_once" ) _ ( ZEND_STR_REQUIRE_ONCE , "require_once" ) _ ( ZEND_STR_SCALAR , "scalar" ) _ ( ZEND_STR_ERROR_REPORTING , "error_reporting" ) _ ( ZEND_STR_STATIC , "static" ) _ ( ZEND_STR_THIS , "this" ) _ ( ZEND_STR_VALUE , "value" ) _ ( ZEND_STR_KEY , "key" ) _ ( ZEND_STR_MAGIC_AUTOLOAD , "__autoload" ) _ ( ZEND_STR_MAGIC_INVOKE , "__invoke" ) _ ( ZEND_STR_PREVIOUS , "previous" ) _ ( ZEND_STR_CODE , "code" ) _ ( ZEND_STR_MESSAGE , "message" ) _ ( ZEND_STR_SEVERITY , "severity" ) _ ( ZEND_STR_STRING , "string" ) _ ( ZEND_STR_TRACE , "trace" ) _ ( ZEND_STR_SCHEME , "scheme" ) _ ( ZEND_STR_HOST , "host" ) _ ( ZEND_STR_PORT , "port" ) _ ( ZEND_STR_USER , "user" ) _ ( ZEND_STR_PASS , "pass" ) _ ( ZEND_STR_PATH , "path" ) _ ( ZEND_STR_QUERY , "query" ) _ ( ZEND_STR_FRAGMENT , "fragment" ) _ ( ZEND_STR_NULL , "NULL" ) _ ( ZEND_STR_BOOLEAN , "boolean" ) _ ( ZEND_STR_INTEGER , "integer" ) _ ( ZEND_STR_DOUBLE , "double" ) _ ( ZEND_STR_ARRAY , "array" ) _ ( ZEND_STR_RESOURCE , "resource" ) _ ( ZEND_STR_CLOSED_RESOURCE , "resource (closed)" ) _ ( ZEND_STR_NAME , "name" ) _ ( ZEND_STR_ARGV , "argv" ) _ ( ZEND_STR_ARGC , "argc" ) _ ( ZEND_STR_ARRAY_CAPITALIZED , "Array" )

type ZendKnownStringId = int

const (
	ZEND_STR_FILE = iota
	ZEND_STR_LINE
	ZEND_STR_FUNCTION
	ZEND_STR_CLASS
	ZEND_STR_OBJECT
	ZEND_STR_TYPE
	ZEND_STR_OBJECT_OPERATOR
	ZEND_STR_PAAMAYIM_NEKUDOTAYIM
	ZEND_STR_ARGS
	ZEND_STR_UNKNOWN
	ZEND_STR_EVAL
	ZEND_STR_INCLUDE
	ZEND_STR_REQUIRE
	ZEND_STR_INCLUDE_ONCE
	ZEND_STR_REQUIRE_ONCE
	ZEND_STR_SCALAR
	ZEND_STR_ERROR_REPORTING
	ZEND_STR_STATIC
	ZEND_STR_THIS
	ZEND_STR_VALUE
	ZEND_STR_KEY
	ZEND_STR_MAGIC_AUTOLOAD
	ZEND_STR_MAGIC_INVOKE
	ZEND_STR_PREVIOUS
	ZEND_STR_CODE
	ZEND_STR_MESSAGE
	ZEND_STR_SEVERITY
	ZEND_STR_STRING
	ZEND_STR_TRACE
	ZEND_STR_SCHEME
	ZEND_STR_HOST
	ZEND_STR_PORT
	ZEND_STR_USER
	ZEND_STR_PASS
	ZEND_STR_PATH
	ZEND_STR_QUERY
	ZEND_STR_FRAGMENT
	ZEND_STR_NULL
	ZEND_STR_BOOLEAN
	ZEND_STR_INTEGER
	ZEND_STR_DOUBLE
	ZEND_STR_ARRAY
	ZEND_STR_RESOURCE
	ZEND_STR_CLOSED_RESOURCE
	ZEND_STR_NAME
	ZEND_STR_ARGV
	ZEND_STR_ARGC
	ZEND_STR_ARRAY_CAPITALIZED
	ZEND_STR_LAST_KNOWN
)

// Source: <Zend/zend_string.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_globals.h"

/* Any strings interned in the startup phase. Common to all the threads,
   won't be free'd until process exit. If we want an ability to
   add permanent strings even after startup, it would be still
   possible on costs of locking in the thread safe builds. */

var InternedStringsPermanent HashTable
var InternedStringRequestHandler ZendNewInternedStringFuncT = ZendNewInternedStringRequest
var InternedStringInitRequestHandler ZendStringInitInternedFuncT = ZendStringInitInternedRequest
var ZendEmptyString *ZendString = nil
var ZendKnownStrings **ZendString = nil

func ZendStringHashFunc(str *ZendString) ZendUlong {
	ZSTR_H(str) = ZendHashFunc(ZSTR_VAL(str), ZSTR_LEN(str))
	return ZSTR_H(str)
}
func ZendHashFunc(str *byte, len_ int) ZendUlong { return ZendInlineHashFunc(str, len_) }
func _strDtor(zv *Zval) {
	var str *ZendString = Z_STR_P(zv)
	Pefree(str, GC_FLAGS(str)&IS_STR_PERSISTENT)
}

var KnownStrings []*byte = []*byte{"file", "line", "function", "class", "object", "type", "->", "::", "args", "unknown", "eval", "include", "require", "include_once", "require_once", "scalar", "error_reporting", "static", "this", "value", "key", "__autoload", "__invoke", "previous", "code", "message", "severity", "string", "trace", "scheme", "host", "port", "user", "pass", "path", "query", "fragment", "NULL", "boolean", "integer", "double", "array", "resource", "resource (closed)", "name", "argv", "argc", "Array", nil}

func ZendInitInternedStringsHt(interned_strings *HashTable, permanent int) {
	ZendHashInit(interned_strings, 1024, nil, _strDtor, permanent)
	if permanent != 0 {
		ZendHashRealInitMixed(interned_strings)
	}
}
func ZendInternedStringsInit() {
	var s []byte
	var i uint
	var str *ZendString
	InternedStringRequestHandler = ZendNewInternedStringRequest
	InternedStringInitRequestHandler = ZendStringInitInternedRequest
	ZendEmptyString = nil
	ZendKnownStrings = nil
	ZendInitInternedStringsHt(&InternedStringsPermanent, 1)
	ZendNewInternedString = ZendNewInternedStringPermanent
	ZendStringInitInterned = ZendStringInitInternedPermanent

	/* interned empty string */

	str = ZendStringAlloc(b.SizeOf("\"\"")-1, 1)
	ZSTR_VAL(str)[0] = '0'
	ZendEmptyString = ZendNewInternedStringPermanent(str)
	s[1] = 0
	for i = 0; i < 256; i++ {
		s[0] = i
		ZendOneCharString[i] = ZendNewInternedStringPermanent(ZendStringInit(s, 1, 1))
	}

	/* known strings */

	ZendKnownStrings = Pemalloc(b.SizeOf("zend_string *")*(b.SizeOf("known_strings")/b.SizeOf("known_strings [ 0 ]")-1), 1)
	for i = 0; i < b.SizeOf("known_strings")/b.SizeOf("known_strings [ 0 ]")-1; i++ {
		str = ZendStringInit(KnownStrings[i], strlen(KnownStrings[i]), 1)
		ZendKnownStrings[i] = ZendNewInternedStringPermanent(str)
	}
}
func ZendInternedStringsDtor() {
	ZendHashDestroy(&InternedStringsPermanent)
	Free(ZendKnownStrings)
	ZendKnownStrings = nil
}
func ZendInternedStringHtLookupEx(h ZendUlong, str *byte, size int, interned_strings *HashTable) *ZendString {
	var nIndex uint32
	var idx uint32
	var p *Bucket
	nIndex = h | interned_strings.GetNTableMask()
	idx = HT_HASH(interned_strings, nIndex)
	for idx != HT_INVALID_IDX {
		p = HT_HASH_TO_BUCKET(interned_strings, idx)
		if p.GetH() == h && ZSTR_LEN(p.GetKey()) == size {
			if !(memcmp(ZSTR_VAL(p.GetKey()), str, size)) {
				return p.GetKey()
			}
		}
		idx = Z_NEXT(p.GetVal())
	}
	return nil
}
func ZendInternedStringHtLookup(str *ZendString, interned_strings *HashTable) *ZendString {
	var h ZendUlong = ZSTR_H(str)
	var nIndex uint32
	var idx uint32
	var p *Bucket
	nIndex = h | interned_strings.GetNTableMask()
	idx = HT_HASH(interned_strings, nIndex)
	for idx != HT_INVALID_IDX {
		p = HT_HASH_TO_BUCKET(interned_strings, idx)
		if p.GetH() == h && ZendStringEqualContent(p.GetKey(), str) != 0 {
			return p.GetKey()
		}
		idx = Z_NEXT(p.GetVal())
	}
	return nil
}

/* This function might be not thread safe at least because it would update the
   hash val in the passed string. Be sure it is called in the appropriate context. */

func ZendAddInternedString(str *ZendString, interned_strings *HashTable, flags uint32) *ZendString {
	var val Zval
	GC_SET_REFCOUNT(str, 1)
	GC_ADD_FLAGS(str, IS_STR_INTERNED|flags)
	ZVAL_INTERNED_STR(&val, str)
	ZendHashAddNew(interned_strings, str, &val)
	return str
}
func ZendInternedStringFindPermanent(str *ZendString) *ZendString {
	ZendStringHashVal(str)
	return ZendInternedStringHtLookup(str, &InternedStringsPermanent)
}
func ZendNewInternedStringPermanent(str *ZendString) *ZendString {
	var ret *ZendString
	if ZSTR_IS_INTERNED(str) != 0 {
		return str
	}
	ZendStringHashVal(str)
	ret = ZendInternedStringHtLookup(str, &InternedStringsPermanent)
	if ret != nil {
		ZendStringRelease(str)
		return ret
	}
	ZEND_ASSERT((GC_FLAGS(str) & GC_PERSISTENT) != 0)
	if GC_REFCOUNT(str) > 1 {
		var h ZendUlong = ZSTR_H(str)
		ZendStringDelref(str)
		str = ZendStringInit(ZSTR_VAL(str), ZSTR_LEN(str), 1)
		ZSTR_H(str) = h
	}
	return ZendAddInternedString(str, &InternedStringsPermanent, IS_STR_PERMANENT)
}
func ZendNewInternedStringRequest(str *ZendString) *ZendString {
	var ret *ZendString
	if ZSTR_IS_INTERNED(str) != 0 {
		return str
	}
	ZendStringHashVal(str)

	/* Check for permanent strings, the table is readonly at this point. */

	ret = ZendInternedStringHtLookup(str, &InternedStringsPermanent)
	if ret != nil {
		ZendStringRelease(str)
		return ret
	}
	ret = ZendInternedStringHtLookup(str, &(CompilerGlobals.GetInternedStrings()))
	if ret != nil {
		ZendStringRelease(str)
		return ret
	}

	/* Create a short living interned, freed after the request. */

	if GC_REFCOUNT(str) > 1 {
		var h ZendUlong = ZSTR_H(str)
		ZendStringDelref(str)
		str = ZendStringInit(ZSTR_VAL(str), ZSTR_LEN(str), 0)
		ZSTR_H(str) = h
	}
	ret = ZendAddInternedString(str, &(CompilerGlobals.GetInternedStrings()), 0)
	return ret
}
func ZendStringInitInternedPermanent(str *byte, size int, permanent int) *ZendString {
	var ret *ZendString
	var h ZendUlong = ZendInlineHashFunc(str, size)
	ret = ZendInternedStringHtLookupEx(h, str, size, &InternedStringsPermanent)
	if ret != nil {
		return ret
	}
	ZEND_ASSERT(permanent != 0)
	ret = ZendStringInit(str, size, permanent)
	ZSTR_H(ret) = h
	return ZendAddInternedString(ret, &InternedStringsPermanent, IS_STR_PERMANENT)
}
func ZendStringInitInternedRequest(str *byte, size int, permanent int) *ZendString {
	var ret *ZendString
	var h ZendUlong = ZendInlineHashFunc(str, size)

	/* Check for permanent strings, the table is readonly at this point. */

	ret = ZendInternedStringHtLookupEx(h, str, size, &InternedStringsPermanent)
	if ret != nil {
		return ret
	}
	ret = ZendInternedStringHtLookupEx(h, str, size, &(CompilerGlobals.GetInternedStrings()))
	if ret != nil {
		return ret
	}
	ret = ZendStringInit(str, size, permanent)
	ZSTR_H(ret) = h

	/* Create a short living interned, freed after the request. */

	return ZendAddInternedString(ret, &(CompilerGlobals.GetInternedStrings()), 0)

	/* Create a short living interned, freed after the request. */
}
func ZendInternedStringsActivate() {
	ZendInitInternedStringsHt(&(CompilerGlobals.GetInternedStrings()), 0)
}
func ZendInternedStringsDeactivate() {
	ZendHashDestroy(&(CompilerGlobals.GetInternedStrings()))
}
func ZendInternedStringsSetRequestStorageHandlers(handler ZendNewInternedStringFuncT, init_handler ZendStringInitInternedFuncT) {
	InternedStringRequestHandler = handler
	InternedStringInitRequestHandler = init_handler
}
func ZendInternedStringsSwitchStorage(request ZendBool) {
	if request != 0 {
		ZendNewInternedString = InternedStringRequestHandler
		ZendStringInitInterned = InternedStringInitRequestHandler
	} else {
		ZendNewInternedString = ZendNewInternedStringPermanent
		ZendStringInitInterned = ZendStringInitInternedPermanent
	}
}
