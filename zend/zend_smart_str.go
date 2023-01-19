// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
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

// #define smart_str_appends_ex(dest,src,what) smart_str_appendl_ex ( ( dest ) , ( src ) , strlen ( src ) , ( what ) )

// #define smart_str_appends(dest,src) smart_str_appendl ( ( dest ) , ( src ) , strlen ( src ) )

// #define smart_str_extend(dest,len) smart_str_extend_ex ( ( dest ) , ( len ) , 0 )

// #define smart_str_appendc(dest,c) smart_str_appendc_ex ( ( dest ) , ( c ) , 0 )

// #define smart_str_appendl(dest,src,len) smart_str_appendl_ex ( ( dest ) , ( src ) , ( len ) , 0 )

// #define smart_str_append(dest,src) smart_str_append_ex ( ( dest ) , ( src ) , 0 )

// #define smart_str_append_smart_str(dest,src) smart_str_append_smart_str_ex ( ( dest ) , ( src ) , 0 )

// #define smart_str_sets(dest,src) smart_str_setl ( ( dest ) , ( src ) , strlen ( src ) ) ;

// #define smart_str_append_long(dest,val) smart_str_append_long_ex ( ( dest ) , ( val ) , 0 )

// #define smart_str_append_unsigned(dest,val) smart_str_append_unsigned_ex ( ( dest ) , ( val ) , 0 )

// #define smart_str_free(dest) smart_str_free_ex ( ( dest ) , 0 )

func SmartStrAlloc(str *SmartStr, len_ int, persistent ZendBool) int {
	if str.GetS() == nil {
		goto do_smart_str_realloc
	} else {
		len_ += str.GetS().GetLen()
		if len_ >= str.GetA() {
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
	var ret *byte = dest.GetS().GetVal() + dest.GetS().GetLen()
	dest.GetS().SetLen(new_len)
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
		str.GetS().GetVal()[str.GetS().GetLen()] = '0'
	}
}
func SmartStrGetLen(str *SmartStr) int {
	if str.GetS() != nil {
		return str.GetS().GetLen()
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
		return ZendEmptyString
	}
}
func SmartStrAppendcEx(dest *SmartStr, ch byte, persistent ZendBool) {
	var new_len int = SmartStrAlloc(dest, 1, persistent)
	dest.GetS().GetVal()[new_len-1] = ch
	dest.GetS().SetLen(new_len)
}
func SmartStrAppendlEx(dest *SmartStr, str *byte, len_ int, persistent ZendBool) {
	var new_len int = SmartStrAlloc(dest, len_, persistent)
	memcpy(dest.GetS().GetVal()+dest.GetS().GetLen(), str, len_)
	dest.GetS().SetLen(new_len)
}
func SmartStrAppendEx(dest *SmartStr, src *ZendString, persistent ZendBool) {
	SmartStrAppendlEx(dest, src.GetVal(), src.GetLen(), persistent)
}
func SmartStrAppendSmartStrEx(dest *SmartStr, src *SmartStr, persistent ZendBool) {
	if src.GetS() != nil && src.GetS().GetLen() != 0 {
		SmartStrAppendEx(dest, src.GetS(), persistent)
	}
}
func SmartStrAppendLongEx(dest *SmartStr, num ZendLong, persistent ZendBool) {
	var buf []byte
	var result *byte = ZendPrintLongToBuf(buf+g.SizeOf("buf")-1, num)
	SmartStrAppendlEx(dest, result, buf+g.SizeOf("buf")-1-result, persistent)
}
func SmartStrAppendUnsignedEx(dest *SmartStr, num ZendUlong, persistent ZendBool) {
	var buf []byte
	var result *byte = ZendPrintUlongToBuf(buf+g.SizeOf("buf")-1, num)
	SmartStrAppendlEx(dest, result, buf+g.SizeOf("buf")-1-result, persistent)
}
func SmartStrSetl(dest *SmartStr, src *byte, len_ int) {
	SmartStrFreeEx(dest, 0)
	SmartStrAppendlEx(dest, src, len_, 0)
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

// #define SMART_STR_OVERHEAD       ( ZEND_MM_OVERHEAD + _ZSTR_HEADER_SIZE + 1 )

// #define SMART_STR_START_SIZE       256

// #define SMART_STR_START_LEN       ( SMART_STR_START_SIZE - SMART_STR_OVERHEAD )

// #define SMART_STR_PAGE       4096

// #define SMART_STR_NEW_LEN(len) ( ZEND_MM_ALIGNED_SIZE_EX ( len + SMART_STR_OVERHEAD , SMART_STR_PAGE ) - SMART_STR_OVERHEAD )

func SmartStrErealloc(str *SmartStr, len_ int) {
	if str.GetS() == nil {
		if len_ <= 256-(0+zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+1) {
			str.SetA(256 - (0 + zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + 1))
		} else {
			str.SetA((len_ + (0 + zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + 1) + (4096-1) & ^(4096-1)) - (0 + zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + 1))
		}
		str.SetS(ZendStringAlloc(str.GetA(), 0))
		str.GetS().SetLen(0)
	} else {
		str.SetA((len_ + (0 + zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + 1) + (4096-1) & ^(4096-1)) - (0 + zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + 1))
		str.SetS((*ZendString)(_erealloc2(str.GetS(), str.GetA()+zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+1, zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+str.GetS().GetLen())))
	}
}
func SmartStrRealloc(str *SmartStr, len_ int) {
	if str.GetS() == nil {
		if len_ <= 256-(0+zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+1) {
			str.SetA(256 - (0 + zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + 1))
		} else {
			str.SetA((len_ + (0 + zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + 1) + (4096-1) & ^(4096-1)) - (0 + zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + 1))
		}
		str.SetS(ZendStringAlloc(str.GetA(), 1))
		str.GetS().SetLen(0)
	} else {
		str.SetA((len_ + (0 + zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + 1) + (4096-1) & ^(4096-1)) - (0 + zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + 1))
		str.SetS((*ZendString)(g.CondF(true, func() any {
			return __zendRealloc(str.GetS(), str.GetA()+zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+1)
		}, func() any {
			return _erealloc(str.GetS(), str.GetA()+zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+1)
		})))
	}
}

/* Windows uses VK_ESCAPE instead of \e */

// #define VK_ESCAPE       '\e'

func ZendComputeEscapedStringLen(s *byte, l int) int {
	var i int
	var len_ int = l
	for i = 0; i < l; i++ {
		var c byte = s[i]
		if c == '\n' || c == '\r' || c == '\t' || c == 'f' || c == 'v' || c == '\\' || c == 'e' {
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
	res = &(str.GetS()).val[str.GetS().GetLen()]
	str.GetS().SetLen(str.GetS().GetLen() + len_)
	for i = 0; i < l; i++ {
		var c uint8 = s[i]
		if c < 32 || c == '\\' || c > 126 {
			g.PostInc(&(*res)) = '\\'
			switch c {
			case '\n':
				g.PostInc(&(*res)) = 'n'
				break
			case '\r':
				g.PostInc(&(*res)) = 'r'
				break
			case '\t':
				g.PostInc(&(*res)) = 't'
				break
			case 'f':
				g.PostInc(&(*res)) = 'f'
				break
			case 'v':
				g.PostInc(&(*res)) = 'v'
				break
			case '\\':
				g.PostInc(&(*res)) = '\\'
				break
			case 'e':
				g.PostInc(&(*res)) = 'e'
				break
			default:
				g.PostInc(&(*res)) = 'x'
				if c>>4 < 10 {
					g.PostInc(&(*res)) = (c >> 4) + '0'
				} else {
					g.PostInc(&(*res)) = (c >> 4) + 'A' - 10
				}
				if (c & 0xf) < 10 {
					g.PostInc(&(*res)) = (c & 0xf) + '0'
				} else {
					g.PostInc(&(*res)) = (c & 0xf) + 'A' - 10
				}
			}
		} else {
			g.PostInc(&(*res)) = c
		}
	}
}
func SmartStrAppendPrintf(dest *SmartStr, format string, _ ...any) {
	var arg va_list
	va_start(arg, format)
	ZendPrintfToSmartStr(dest, format, arg)
	va_end(arg)
}

// #define SMART_STRING_OVERHEAD       ( ZEND_MM_OVERHEAD + 1 )

// #define SMART_STRING_START_SIZE       256

// #define SMART_STRING_START_LEN       ( SMART_STRING_START_SIZE - SMART_STRING_OVERHEAD )

// #define SMART_STRING_PAGE       4096

func _smartStringAllocPersistent(str *SmartString, len_ int) {
	if str.GetC() == nil {
		str.SetLen(0)
		if len_ <= 256-(0+1) {
			str.SetA(256 - (0 + 1))
		} else {
			str.SetA((len_ + (0 + 1) + (4096-1) & ^(4096-1)) - (0 + 1))
		}
		str.SetC(__zendMalloc(str.GetA() + 1))
	} else {
		if int(len_ > SIZE_MAX-str.GetLen()) != 0 {
			ZendError(1<<0, "String size overflow")
		}
		len_ += str.GetLen()
		str.SetA((len_ + (0 + 1) + (4096-1) & ^(4096-1)) - (0 + 1))
		str.SetC(__zendRealloc(str.GetC(), str.GetA()+1))
	}
}
func _smartStringAlloc(str *SmartString, len_ int) {
	if str.GetC() == nil {
		str.SetLen(0)
		if len_ <= 256-(0+1) {
			str.SetA(256 - (0 + 1))
			str.SetC(_emalloc(256 - (0 + 1) + 1))
		} else {
			str.SetA((len_ + (0 + 1) + (4096-1) & ^(4096-1)) - (0 + 1))
			if str.GetA() < 2*1024*1024-(0+1) {
				str.SetC(_emalloc(str.GetA() + 1))
			} else {

				/* allocate a huge chunk */

				str.SetC(_emalloc(str.GetA() + 1))

				/* allocate a huge chunk */

			}
		}
	} else {
		if int(len_ > SIZE_MAX-str.GetLen()) != 0 {
			ZendError(1<<0, "String size overflow")
		}
		len_ += str.GetLen()
		str.SetA((len_ + (0 + 1) + (4096-1) & ^(4096-1)) - (0 + 1))
		str.SetC(_erealloc2(str.GetC(), str.GetA()+1, str.GetLen()))
	}
}
