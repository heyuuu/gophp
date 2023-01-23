// <<generate>>

package zend

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
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

// #define ZSTR_VAL(zstr) ( zstr ) -> val

// #define ZSTR_LEN(zstr) ( zstr ) -> len

// #define ZSTR_H(zstr) ( zstr ) -> h

// #define ZSTR_HASH(zstr) zend_string_hash_val ( zstr )

/* Compatibility macros */

// #define IS_INTERNED(s) ZSTR_IS_INTERNED ( s )

// #define STR_EMPTY_ALLOC() ZSTR_EMPTY_ALLOC ( )

// #define _STR_HEADER_SIZE       _ZSTR_HEADER_SIZE

// #define STR_ALLOCA_ALLOC(str,_len,use_heap) ZSTR_ALLOCA_ALLOC ( str , _len , use_heap )

// #define STR_ALLOCA_INIT(str,s,len,use_heap) ZSTR_ALLOCA_INIT ( str , s , len , use_heap )

// #define STR_ALLOCA_FREE(str,use_heap) ZSTR_ALLOCA_FREE ( str , use_heap )

/*---*/

// #define ZSTR_IS_INTERNED(s) ( GC_FLAGS ( s ) & IS_STR_INTERNED )

// #define ZSTR_EMPTY_ALLOC() zend_empty_string

// #define ZSTR_CHAR(c) zend_one_char_string [ c ]

// #define ZSTR_KNOWN(idx) zend_known_strings [ idx ]

// #define _ZSTR_HEADER_SIZE       XtOffsetOf ( zend_string , val )

// #define _ZSTR_STRUCT_SIZE(len) ( _ZSTR_HEADER_SIZE + len + 1 )

// #define ZSTR_MAX_OVERHEAD       ( ZEND_MM_ALIGNED_SIZE ( _ZSTR_HEADER_SIZE + 1 ) )

// #define ZSTR_MAX_LEN       ( SIZE_MAX - ZSTR_MAX_OVERHEAD )

// #define ZSTR_ALLOCA_ALLOC(str,_len,use_heap) do { ( str ) = ( zend_string * ) do_alloca ( ZEND_MM_ALIGNED_SIZE_EX ( _ZSTR_STRUCT_SIZE ( _len ) , 8 ) , ( use_heap ) ) ; GC_SET_REFCOUNT ( str , 1 ) ; GC_TYPE_INFO ( str ) = IS_STRING ; ZSTR_H ( str ) = 0 ; ZSTR_LEN ( str ) = _len ; } while ( 0 )

// #define ZSTR_ALLOCA_INIT(str,s,len,use_heap) do { ZSTR_ALLOCA_ALLOC ( str , len , use_heap ) ; memcpy ( ZSTR_VAL ( str ) , ( s ) , ( len ) ) ; ZSTR_VAL ( str ) [ ( len ) ] = '\0' ; } while ( 0 )

// #define ZSTR_ALLOCA_FREE(str,use_heap) free_alloca ( str , use_heap )

/*---*/

func ZendStringHashVal(s *ZendString) ZendUlong {
	if s.GetH() != 0 {
		return s.GetH()
	} else {
		return ZendStringHashFunc(s)
	}
}
func ZendStringForgetHashVal(s *ZendString) {
	s.SetH(0)
	s.GetGc().SetTypeInfo(s.GetGc().GetTypeInfo() &^ (1 << 9 << 0))
}
func ZendStringRefcount(s *ZendString) uint32 {
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		return ZendGcRefcount(&s.gc)
	}
	return 1
}
func ZendStringAddref(s *ZendString) uint32 {
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		return ZendGcAddref(&s.gc)
	}
	return 1
}
func ZendStringDelref(s *ZendString) uint32 {
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		return ZendGcDelref(&s.gc)
	}
	return 1
}
func ZendStringAlloc(len_ int, persistent int) *ZendString {
	var ret *ZendString = (*ZendString)(g.CondF(persistent != 0, func() any {
		return __zendMalloc(zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + len_ + 1 + 8 - 1 & ^(8-1))
	}, func() any {
		return _emalloc(zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + len_ + 1 + 8 - 1 & ^(8-1))
	}))
	ZendGcSetRefcount(&ret.gc, 1)
	ret.GetGc().SetTypeInfo(6 | g.Cond(persistent != 0, 1<<7, 0)<<0)
	ret.SetH(0)
	ret.SetLen(len_)
	return ret
}
func ZendStringSafeAlloc(n int, m int, l int, persistent int) *ZendString {
	var ret *ZendString = (*ZendString)(g.CondF(persistent != 0, func() any {
		return _safeMalloc(n, m, zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+l+1+8 - 1 & ^(8-1))
	}, func() any {
		return _safeEmalloc(n, m, zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+l+1+8 - 1 & ^(8-1))
	}))
	ZendGcSetRefcount(&ret.gc, 1)
	ret.GetGc().SetTypeInfo(6 | g.Cond(persistent != 0, 1<<7, 0)<<0)
	ret.SetH(0)
	ret.SetLen(n*m + l)
	return ret
}
func ZendStringInit(str *byte, len_ int, persistent int) *ZendString {
	var ret *ZendString = ZendStringAlloc(len_, persistent)
	memcpy(ret.GetVal(), str, len_)
	ret.GetVal()[len_] = '0'
	return ret
}
func ZendStringCopy(s *ZendString) *ZendString {
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ZendGcAddref(&s.gc)
	}
	return s
}
func ZendStringDup(s *ZendString, persistent int) *ZendString {
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		return s
	} else {
		return ZendStringInit(s.GetVal(), s.GetLen(), persistent)
	}
}
func ZendStringRealloc(s *ZendString, len_ int, persistent int) *ZendString {
	var ret *ZendString
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		if ZendGcRefcount(&s.gc) == 1 {
			ret = (*ZendString)(g.CondF(persistent != 0, func() any {
				return __zendRealloc(s, zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+len_+1+8 - 1 & ^(8-1))
			}, func() any {
				return _erealloc(s, zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+len_+1+8 - 1 & ^(8-1))
			}))
			ret.SetLen(len_)
			ZendStringForgetHashVal(ret)
			return ret
		}
	}
	ret = ZendStringAlloc(len_, persistent)
	memcpy(ret.GetVal(), s.GetVal(), g.CondF2(len_ < s.GetLen(), len_, func() int { return s.GetLen() })+1)
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ZendGcDelref(&s.gc)
	}
	return ret
}
func ZendStringExtend(s *ZendString, len_ int, persistent int) *ZendString {
	var ret *ZendString
	r.Assert(len_ >= s.GetLen())
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		if ZendGcRefcount(&s.gc) == 1 {
			ret = (*ZendString)(g.CondF(persistent != 0, func() any {
				return __zendRealloc(s, zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+len_+1+8 - 1 & ^(8-1))
			}, func() any {
				return _erealloc(s, zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+len_+1+8 - 1 & ^(8-1))
			}))
			ret.SetLen(len_)
			ZendStringForgetHashVal(ret)
			return ret
		}
	}
	ret = ZendStringAlloc(len_, persistent)
	memcpy(ret.GetVal(), s.GetVal(), s.GetLen()+1)
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ZendGcDelref(&s.gc)
	}
	return ret
}
func ZendStringTruncate(s *ZendString, len_ int, persistent int) *ZendString {
	var ret *ZendString
	r.Assert(len_ <= s.GetLen())
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		if ZendGcRefcount(&s.gc) == 1 {
			ret = (*ZendString)(g.CondF(persistent != 0, func() any {
				return __zendRealloc(s, zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+len_+1+8 - 1 & ^(8-1))
			}, func() any {
				return _erealloc(s, zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+len_+1+8 - 1 & ^(8-1))
			}))
			ret.SetLen(len_)
			ZendStringForgetHashVal(ret)
			return ret
		}
	}
	ret = ZendStringAlloc(len_, persistent)
	memcpy(ret.GetVal(), s.GetVal(), len_+1)
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ZendGcDelref(&s.gc)
	}
	return ret
}
func ZendStringSafeRealloc(s *ZendString, n int, m int, l int, persistent int) *ZendString {
	var ret *ZendString
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		if ZendGcRefcount(&s.gc) == 1 {
			ret = (*ZendString)(g.CondF(persistent != 0, func() any {
				return _safeRealloc(s, n, m, zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+l+1+8 - 1 & ^(8-1))
			}, func() any {
				return _safeErealloc(s, n, m, zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+l+1+8 - 1 & ^(8-1))
			}))
			ret.SetLen(n*m + l)
			ZendStringForgetHashVal(ret)
			return ret
		}
	}
	ret = ZendStringSafeAlloc(n, m, l, persistent)
	memcpy(ret.GetVal(), s.GetVal(), g.CondF2(n*m+l < s.GetLen(), n*m+l, func() int { return s.GetLen() })+1)
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ZendGcDelref(&s.gc)
	}
	return ret
}
func ZendStringFree(s *ZendString) {
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		r.Assert(ZendGcRefcount(&s.gc) <= 1)
		g.CondF((ZvalGcFlags(s.GetGc().GetTypeInfo())&1<<7) != 0, func() { return Free(s) }, func() { return _efree(s) })
	}
}
func ZendStringEfree(s *ZendString) {
	r.Assert((ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0)
	r.Assert(ZendGcRefcount(&s.gc) <= 1)
	r.Assert((ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 7) == 0)
	_efree(s)
}
func ZendStringRelease(s *ZendString) {
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		if ZendGcDelref(&s.gc) == 0 {
			g.CondF((ZvalGcFlags(s.GetGc().GetTypeInfo())&1<<7) != 0, func() { return Free(s) }, func() { return _efree(s) })
		}
	}
}
func ZendStringReleaseEx(s *ZendString, persistent int) {
	if (ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		if ZendGcDelref(&s.gc) == 0 {
			if persistent != 0 {
				r.Assert((ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 7) != 0)
				Free(s)
			} else {
				r.Assert((ZvalGcFlags(s.GetGc().GetTypeInfo()) & 1 << 7) == 0)
				_efree(s)
			}
		}
	}
}
func ZendStringEqualVal(s1 *ZendString, s2 *ZendString) ZendBool {
	return !(memcmp(s1.GetVal(), s2.GetVal(), s1.GetLen()))
}
func ZendStringEqualContent(s1 *ZendString, s2 *ZendString) ZendBool {
	return s1.GetLen() == s2.GetLen() && ZendStringEqualVal(s1, s2) != 0
}
func ZendStringEquals(s1 *ZendString, s2 *ZendString) ZendBool {
	return s1 == s2 || ZendStringEqualContent(s1, s2) != 0
}

// #define zend_string_equals_ci(s1,s2) ( ZSTR_LEN ( s1 ) == ZSTR_LEN ( s2 ) && ! zend_binary_strcasecmp ( ZSTR_VAL ( s1 ) , ZSTR_LEN ( s1 ) , ZSTR_VAL ( s2 ) , ZSTR_LEN ( s2 ) ) )

// #define zend_string_equals_literal_ci(str,c) ( ZSTR_LEN ( str ) == sizeof ( c ) - 1 && ! zend_binary_strcasecmp ( ZSTR_VAL ( str ) , ZSTR_LEN ( str ) , ( c ) , sizeof ( c ) - 1 ) )

// #define zend_string_equals_literal(str,literal) ( ZSTR_LEN ( str ) == sizeof ( literal ) - 1 && ! memcmp ( ZSTR_VAL ( str ) , literal , sizeof ( literal ) - 1 ) )

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
	var hash ZendUlong = 5381

	/* variant with the hash unrolled eight times */

	for ; len_ >= 8; len_ -= 8 {
		hash = (hash << 5) + hash + g.PostInc(&(*str))
		hash = (hash << 5) + hash + g.PostInc(&(*str))
		hash = (hash << 5) + hash + g.PostInc(&(*str))
		hash = (hash << 5) + hash + g.PostInc(&(*str))
		hash = (hash << 5) + hash + g.PostInc(&(*str))
		hash = (hash << 5) + hash + g.PostInc(&(*str))
		hash = (hash << 5) + hash + g.PostInc(&(*str))
		hash = (hash << 5) + hash + g.PostInc(&(*str))
	}
	switch len_ {
	case 7:
		hash = (hash << 5) + hash + g.PostInc(&(*str))
	case 6:
		hash = (hash << 5) + hash + g.PostInc(&(*str))
	case 5:
		hash = (hash << 5) + hash + g.PostInc(&(*str))
	case 4:
		hash = (hash << 5) + hash + g.PostInc(&(*str))
	case 3:
		hash = (hash << 5) + hash + g.PostInc(&(*str))
	case 2:
		hash = (hash << 5) + hash + g.PostInc(&(*str))
	case 1:
		hash = (hash << 5) + hash + g.PostInc(&(*str))
		break
	case 0:
		break
	default:
		break
	}

	/* Hash value can't be zero, so we always set the high bit */

	return hash | -0x8000000000000000

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
	str.SetH(ZendHashFunc(str.GetVal(), str.GetLen()))
	return str.GetH()
}
func ZendHashFunc(str *byte, len_ int) ZendUlong { return ZendInlineHashFunc(str, len_) }
func _strDtor(zv *Zval) {
	var str *ZendString = zv.GetValue().GetStr()
	g.CondF((ZvalGcFlags(str.GetGc().GetTypeInfo())&1<<7) != 0, func() { return Free(str) }, func() { return _efree(str) })
}

var KnownStrings []*byte = []*byte{"file", "line", "function", "class", "object", "type", "->", "::", "args", "unknown", "eval", "include", "require", "include_once", "require_once", "scalar", "error_reporting", "static", "this", "value", "key", "__autoload", "__invoke", "previous", "code", "message", "severity", "string", "trace", "scheme", "host", "port", "user", "pass", "path", "query", "fragment", "NULL", "boolean", "integer", "double", "array", "resource", "resource (closed)", "name", "argv", "argc", "Array", nil}

func ZendInitInternedStringsHt(interned_strings *HashTable, permanent int) {
	_zendHashInit(interned_strings, 1024, _strDtor, permanent)
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

	str = ZendStringAlloc(g.SizeOf("\"\"")-1, 1)
	str.GetVal()[0] = '0'
	ZendEmptyString = ZendNewInternedStringPermanent(str)
	s[1] = 0
	for i = 0; i < 256; i++ {
		s[0] = i
		ZendOneCharString[i] = ZendNewInternedStringPermanent(ZendStringInit(s, 1, 1))
	}

	/* known strings */

	ZendKnownStrings = __zendMalloc(g.SizeOf("zend_string *") * (g.SizeOf("known_strings")/g.SizeOf("known_strings [ 0 ]") - 1))
	for i = 0; i < g.SizeOf("known_strings")/g.SizeOf("known_strings [ 0 ]")-1; i++ {
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
	idx = (*uint32)(interned_strings.GetArData())[int32(nIndex)]
	for idx != uint32-1 {
		p = interned_strings.GetArData() + idx
		if p.GetH() == h && p.GetKey().GetLen() == size {
			if !(memcmp(p.GetKey().GetVal(), str, size)) {
				return p.GetKey()
			}
		}
		idx = p.GetVal().GetNext()
	}
	return nil
}
func ZendInternedStringHtLookup(str *ZendString, interned_strings *HashTable) *ZendString {
	var h ZendUlong = str.GetH()
	var nIndex uint32
	var idx uint32
	var p *Bucket
	nIndex = h | interned_strings.GetNTableMask()
	idx = (*uint32)(interned_strings.GetArData())[int32(nIndex)]
	for idx != uint32-1 {
		p = interned_strings.GetArData() + idx
		if p.GetH() == h && ZendStringEqualContent(p.GetKey(), str) != 0 {
			return p.GetKey()
		}
		idx = p.GetVal().GetNext()
	}
	return nil
}

/* This function might be not thread safe at least because it would update the
   hash val in the passed string. Be sure it is called in the appropriate context. */

func ZendAddInternedString(str *ZendString, interned_strings *HashTable, flags uint32) *ZendString {
	var val Zval
	ZendGcSetRefcount(&str.gc, 1)
	str.GetGc().SetTypeInfo(str.GetGc().GetTypeInfo() | (1<<6|flags)<<0)
	var __z *Zval = &val
	var __s *ZendString = str
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6)
	ZendHashAddNew(interned_strings, str, &val)
	return str
}
func ZendInternedStringFindPermanent(str *ZendString) *ZendString {
	ZendStringHashVal(str)
	return ZendInternedStringHtLookup(str, &InternedStringsPermanent)
}
func ZendNewInternedStringPermanent(str *ZendString) *ZendString {
	var ret *ZendString
	if (ZvalGcFlags(str.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		return str
	}
	ZendStringHashVal(str)
	ret = ZendInternedStringHtLookup(str, &InternedStringsPermanent)
	if ret != nil {
		ZendStringRelease(str)
		return ret
	}
	r.Assert((ZvalGcFlags(str.GetGc().GetTypeInfo()) & 1 << 7) != 0)
	if ZendGcRefcount(&str.gc) > 1 {
		var h ZendUlong = str.GetH()
		ZendStringDelref(str)
		str = ZendStringInit(str.GetVal(), str.GetLen(), 1)
		str.SetH(h)
	}
	return ZendAddInternedString(str, &InternedStringsPermanent, 1<<8)
}
func ZendNewInternedStringRequest(str *ZendString) *ZendString {
	var ret *ZendString
	if (ZvalGcFlags(str.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		return str
	}
	ZendStringHashVal(str)

	/* Check for permanent strings, the table is readonly at this point. */

	ret = ZendInternedStringHtLookup(str, &InternedStringsPermanent)
	if ret != nil {
		ZendStringRelease(str)
		return ret
	}
	ret = ZendInternedStringHtLookup(str, &CG.interned_strings)
	if ret != nil {
		ZendStringRelease(str)
		return ret
	}

	/* Create a short living interned, freed after the request. */

	if ZendGcRefcount(&str.gc) > 1 {
		var h ZendUlong = str.GetH()
		ZendStringDelref(str)
		str = ZendStringInit(str.GetVal(), str.GetLen(), 0)
		str.SetH(h)
	}
	ret = ZendAddInternedString(str, &CG.interned_strings, 0)
	return ret
}
func ZendStringInitInternedPermanent(str *byte, size int, permanent int) *ZendString {
	var ret *ZendString
	var h ZendUlong = ZendInlineHashFunc(str, size)
	ret = ZendInternedStringHtLookupEx(h, str, size, &InternedStringsPermanent)
	if ret != nil {
		return ret
	}
	r.Assert(permanent != 0)
	ret = ZendStringInit(str, size, permanent)
	ret.SetH(h)
	return ZendAddInternedString(ret, &InternedStringsPermanent, 1<<8)
}
func ZendStringInitInternedRequest(str *byte, size int, permanent int) *ZendString {
	var ret *ZendString
	var h ZendUlong = ZendInlineHashFunc(str, size)

	/* Check for permanent strings, the table is readonly at this point. */

	ret = ZendInternedStringHtLookupEx(h, str, size, &InternedStringsPermanent)
	if ret != nil {
		return ret
	}
	ret = ZendInternedStringHtLookupEx(h, str, size, &CG.interned_strings)
	if ret != nil {
		return ret
	}
	ret = ZendStringInit(str, size, permanent)
	ret.SetH(h)

	/* Create a short living interned, freed after the request. */

	return ZendAddInternedString(ret, &CG.interned_strings, 0)

	/* Create a short living interned, freed after the request. */
}
func ZendInternedStringsActivate() {
	ZendInitInternedStringsHt(&CG.interned_strings, 0)
}
func ZendInternedStringsDeactivate() { ZendHashDestroy(&CG.interned_strings) }
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
