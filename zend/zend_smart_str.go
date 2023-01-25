// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_smart_str.h>

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
   | Author: Sascha Schumann <sascha@schumann.cx>                         |
   +----------------------------------------------------------------------+
*/

// #define ZEND_SMART_STR_H

// # include < zend . h >

// # include "zend_globals.h"

// # include "zend_smart_str_public.h"

func SmartStrAppendsEx(dest *SmartStr, src *byte, what ZendBool) {
	SmartStrAppendlEx(dest, src, strlen(src), what)
}
func SmartStrAppends(dest *SmartStr, src *byte)            { SmartStrAppendl(dest, src, strlen(src)) }
func SmartStrExtend(dest *SmartStr, len_ int) *byte        { return SmartStrExtendEx(dest, len_, 0) }
func SmartStrAppendc(dest *SmartStr, c byte)               { SmartStrAppendcEx(dest, c, 0) }
func SmartStrAppendl(dest *SmartStr, src *byte, len_ int)  { SmartStrAppendlEx(dest, src, len_, 0) }
func SmartStrAppend(dest *SmartStr, src *ZendString)       { SmartStrAppendEx(dest, src, 0) }
func SmartStrAppendSmartStr(dest *SmartStr, src *SmartStr) { SmartStrAppendSmartStrEx(dest, src, 0) }
func SmartStrSets(dest *SmartStr, src *byte)               { SmartStrSetl(dest, src, strlen(src)) }
func SmartStrAppendLong(dest *SmartStr, val ZendLong)      { SmartStrAppendLongEx(dest, val, 0) }
func SmartStrAppendUnsigned(dest *SmartStr, val ZendUlong) { SmartStrAppendUnsignedEx(dest, val, 0) }
func SmartStrFree(dest *SmartStr)                          { SmartStrFreeEx(dest, 0) }
func SmartStrAlloc(str *SmartStr, len_ int, persistent ZendBool) int {
	if UNEXPECTED(str.GetS() == nil) {
		goto do_smart_str_realloc
	} else {
		len_ += ZSTR_LEN(str.GetS())
		if UNEXPECTED(len_ >= str.GetA()) {
		do_smart_str_realloc:
			if persistent != 0 {
				SmartStrRealloc(str, len_)
			} else {
				SmartStrErealloc(str, len_)
			}
		}
	}
	return len_
}
func SmartStrExtendEx(dest *SmartStr, len_ int, persistent ZendBool) *byte {
	var new_len int = SmartStrAlloc(dest, len_, persistent)
	var ret *byte = ZSTR_VAL(dest.GetS()) + ZSTR_LEN(dest.GetS())
	ZSTR_LEN(dest.GetS()) = new_len
	return ret
}
func SmartStrFreeEx(str *SmartStr, persistent ZendBool) {
	if str.GetS() != nil {
		ZendStringReleaseEx(str.GetS(), persistent)
		str.SetS(nil)
	}
	str.SetA(0)
}
func SmartStr0(str *SmartStr) {
	if str.GetS() != nil {
		ZSTR_VAL(str.GetS())[ZSTR_LEN(str.GetS())] = '0'
	}
}
func SmartStrGetLen(str *SmartStr) int {
	if str.GetS() != nil {
		return ZSTR_LEN(str.GetS())
	} else {
		return 0
	}
}
func SmartStrExtract(str *SmartStr) *ZendString {
	if str.GetS() != nil {
		var res *ZendString
		SmartStr0(str)
		res = str.GetS()
		str.SetS(nil)
		return res
	} else {
		return ZSTR_EMPTY_ALLOC()
	}
}
func SmartStrAppendcEx(dest *SmartStr, ch byte, persistent ZendBool) {
	var new_len int = SmartStrAlloc(dest, 1, persistent)
	ZSTR_VAL(dest.GetS())[new_len-1] = ch
	ZSTR_LEN(dest.GetS()) = new_len
}
func SmartStrAppendlEx(dest *SmartStr, str string, len_ int, persistent ZendBool) {
	var new_len int = SmartStrAlloc(dest, len_, persistent)
	memcpy(ZSTR_VAL(dest.GetS())+ZSTR_LEN(dest.GetS()), str, len_)
	ZSTR_LEN(dest.GetS()) = new_len
}
func SmartStrAppendEx(dest *SmartStr, src *ZendString, persistent ZendBool) {
	SmartStrAppendlEx(dest, ZSTR_VAL(src), ZSTR_LEN(src), persistent)
}
func SmartStrAppendSmartStrEx(dest *SmartStr, src *SmartStr, persistent ZendBool) {
	if src.GetS() != nil && ZSTR_LEN(src.GetS()) != 0 {
		SmartStrAppendEx(dest, src.GetS(), persistent)
	}
}
func SmartStrAppendLongEx(dest *SmartStr, num ZendLong, persistent ZendBool) {
	var buf []byte
	var result *byte = ZendPrintLongToBuf(buf+b.SizeOf("buf")-1, num)
	SmartStrAppendlEx(dest, result, buf+b.SizeOf("buf")-1-result, persistent)
}
func SmartStrAppendUnsignedEx(dest *SmartStr, num ZendUlong, persistent ZendBool) {
	var buf []byte
	var result *byte = ZendPrintUlongToBuf(buf+b.SizeOf("buf")-1, num)
	SmartStrAppendlEx(dest, result, buf+b.SizeOf("buf")-1-result, persistent)
}
func SmartStrSetl(dest *SmartStr, src *byte, len_ int) {
	SmartStrFree(dest)
	SmartStrAppendl(dest, src, len_)
}

// Source: <Zend/zend_smart_str.c>

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
   | Author: Dmitry Stogov <dmitry@php.net>                               |
   +----------------------------------------------------------------------+
*/

// # include < zend . h >

// # include "zend_smart_str.h"

// # include "zend_smart_string.h"

const SMART_STR_OVERHEAD = ZEND_MM_OVERHEAD + _ZSTR_HEADER_SIZE + 1
const SMART_STR_START_SIZE = 256
const SMART_STR_START_LEN = SMART_STR_START_SIZE - SMART_STR_OVERHEAD
const SMART_STR_PAGE = 4096

func SMART_STR_NEW_LEN(len_ int) int {
	return ZEND_MM_ALIGNED_SIZE_EX(len_+SMART_STR_OVERHEAD, SMART_STR_PAGE) - SMART_STR_OVERHEAD
}
func SmartStrErealloc(str *SmartStr, len_ int) {
	if UNEXPECTED(str.GetS() == nil) {
		if len_ <= SMART_STR_START_LEN {
			str.SetA(SMART_STR_START_LEN)
		} else {
			str.SetA(SMART_STR_NEW_LEN(len_))
		}
		str.SetS(ZendStringAlloc(str.GetA(), 0))
		ZSTR_LEN(str.GetS()) = 0
	} else {
		str.SetA(SMART_STR_NEW_LEN(len_))
		str.SetS((*ZendString)(Erealloc2(str.GetS(), str.GetA()+_ZSTR_HEADER_SIZE+1, _ZSTR_HEADER_SIZE+ZSTR_LEN(str.GetS()))))
	}
}
func SmartStrRealloc(str *SmartStr, len_ int) {
	if UNEXPECTED(str.GetS() == nil) {
		if len_ <= SMART_STR_START_LEN {
			str.SetA(SMART_STR_START_LEN)
		} else {
			str.SetA(SMART_STR_NEW_LEN(len_))
		}
		str.SetS(ZendStringAlloc(str.GetA(), 1))
		ZSTR_LEN(str.GetS()) = 0
	} else {
		str.SetA(SMART_STR_NEW_LEN(len_))
		str.SetS((*ZendString)(Perealloc(str.GetS(), str.GetA()+_ZSTR_HEADER_SIZE+1, 1)))
	}
}

/* Windows uses VK_ESCAPE instead of \e */

const VK_ESCAPE = 'e'

func ZendComputeEscapedStringLen(s *byte, l int) int {
	var i int
	var len_ int = l
	for i = 0; i < l; i++ {
		var c byte = s[i]
		if c == '\n' || c == '\r' || c == '\t' || c == 'f' || c == 'v' || c == '\\' || c == VK_ESCAPE {
			len_ += 1
		} else if c < 32 || c > 126 {
			len_ += 3
		}
	}
	return len_
}
func SmartStrAppendEscaped(str *SmartStr, s *byte, l int) {
	var res *byte
	var i int
	var len_ int = ZendComputeEscapedStringLen(s, l)
	SmartStrAlloc(str, len_, 0)
	res = &ZSTR_VAL(str.GetS())[ZSTR_LEN(str.GetS())]
	ZSTR_LEN(str.GetS()) += len_
	for i = 0; i < l; i++ {
		var c uint8 = s[i]
		if c < 32 || c == '\\' || c > 126 {
			b.PostInc(&(*res)) = '\\'
			switch c {
			case '\n':
				b.PostInc(&(*res)) = 'n'
				break
			case '\r':
				b.PostInc(&(*res)) = 'r'
				break
			case '\t':
				b.PostInc(&(*res)) = 't'
				break
			case 'f':
				b.PostInc(&(*res)) = 'f'
				break
			case 'v':
				b.PostInc(&(*res)) = 'v'
				break
			case '\\':
				b.PostInc(&(*res)) = '\\'
				break
			case VK_ESCAPE:
				b.PostInc(&(*res)) = 'e'
				break
			default:
				b.PostInc(&(*res)) = 'x'
				if c>>4 < 10 {
					b.PostInc(&(*res)) = (c >> 4) + '0'
				} else {
					b.PostInc(&(*res)) = (c >> 4) + 'A' - 10
				}
				if (c & 0xf) < 10 {
					b.PostInc(&(*res)) = (c & 0xf) + '0'
				} else {
					b.PostInc(&(*res)) = (c & 0xf) + 'A' - 10
				}
			}
		} else {
			b.PostInc(&(*res)) = c
		}
	}
}
func SmartStrAppendPrintf(dest *SmartStr, format string, _ ...any) {
	var arg va_list
	va_start(arg, format)
	ZendPrintfToSmartStr(dest, format, arg)
	va_end(arg)
}

const SMART_STRING_OVERHEAD = ZEND_MM_OVERHEAD + 1
const SMART_STRING_START_SIZE = 256
const SMART_STRING_START_LEN = SMART_STRING_START_SIZE - SMART_STRING_OVERHEAD
const SMART_STRING_PAGE = 4096

func _smartStringAllocPersistent(str *SmartString, len_ int) {
	if str.GetC() == nil {
		str.SetLen(0)
		if len_ <= SMART_STRING_START_LEN {
			str.SetA(SMART_STRING_START_LEN)
		} else {
			str.SetA(ZEND_MM_ALIGNED_SIZE_EX(len_+SMART_STRING_OVERHEAD, SMART_STRING_PAGE) - SMART_STRING_OVERHEAD)
		}
		str.SetC(Pemalloc(str.GetA()+1, 1))
	} else {
		if UNEXPECTED(int(len_ > SIZE_MAX-str.GetLen()) != 0) {
			ZendError(E_ERROR, "String size overflow")
		}
		len_ += str.GetLen()
		str.SetA(ZEND_MM_ALIGNED_SIZE_EX(len_+SMART_STRING_OVERHEAD, SMART_STRING_PAGE) - SMART_STRING_OVERHEAD)
		str.SetC(Perealloc(str.GetC(), str.GetA()+1, 1))
	}
}
func _smartStringAlloc(str *SmartString, len_ int) {
	if str.GetC() == nil {
		str.SetLen(0)
		if len_ <= SMART_STRING_START_LEN {
			str.SetA(SMART_STRING_START_LEN)
			str.SetC(Emalloc(SMART_STRING_START_LEN + 1))
		} else {
			str.SetA(ZEND_MM_ALIGNED_SIZE_EX(len_+SMART_STRING_OVERHEAD, SMART_STRING_PAGE) - SMART_STRING_OVERHEAD)
			if EXPECTED(str.GetA() < ZEND_MM_CHUNK_SIZE-SMART_STRING_OVERHEAD) {
				str.SetC(EmallocLarge(str.GetA() + 1))
			} else {

				/* allocate a huge chunk */

				str.SetC(Emalloc(str.GetA() + 1))

				/* allocate a huge chunk */

			}
		}
	} else {
		if UNEXPECTED(int(len_ > SIZE_MAX-str.GetLen()) != 0) {
			ZendError(E_ERROR, "String size overflow")
		}
		len_ += str.GetLen()
		str.SetA(ZEND_MM_ALIGNED_SIZE_EX(len_+SMART_STRING_OVERHEAD, SMART_STRING_PAGE) - SMART_STRING_OVERHEAD)
		str.SetC(Erealloc2(str.GetC(), str.GetA()+1, str.GetLen()))
	}
}
