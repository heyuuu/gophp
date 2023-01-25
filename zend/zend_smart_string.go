// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_smart_string.h>

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
   |         Xinchen Hui <laruence@php.net>                               |
   +----------------------------------------------------------------------+
*/

// #define PHP_SMART_STRING_H

// # include "zend_smart_string_public.h"

// # include < stdlib . h >

// # include < zend . h >

/* wrapper */

func SmartStringAppendsEx(str *SmartString, src *byte, what ZendBool) {
	SmartStringAppendlEx(str, src, strlen(src), what)
}
func SmartStringAppends(str *SmartString, src *byte) {
	SmartStringAppendl(str, src, strlen(src))
}
func SmartStringAppendEx(str *SmartString, src __auto__, what ZendBool) {
	SmartStringAppendlEx(str, (*SmartString)(src).GetC(), (*SmartString)(src).GetLen(), what)
}
func SmartStringSets(str *SmartString, src *byte) { SmartStringSetl(str, src, strlen(src)) }
func SmartStringAppendc(str *SmartString, c byte) { SmartStringAppendcEx(str, c, 0) }
func SmartStringFree(s *SmartString)              { SmartStringFreeEx(s, 0) }
func SmartStringAppendl(str *SmartString, src string, len_ int) {
	SmartStringAppendlEx(str, src, len_, 0)
}
func SmartStringAppend(str *SmartString, src __auto__)     { SmartStringAppendEx(str, src, 0) }
func SmartStringAppendLong(str *SmartString, val ZendLong) { SmartStringAppendLongEx(str, val, 0) }
func SmartStringAppendUnsigned(str *SmartString, val ZendUlong) {
	SmartStringAppendUnsignedEx(str, val, 0)
}
func SmartStringAlloc(str *SmartString, len_ int, persistent ZendBool) int {
	if UNEXPECTED(str.GetC() == nil) || UNEXPECTED(len_ >= str.GetA()-str.GetLen()) {
		if persistent != 0 {
			_smartStringAllocPersistent(str, len_)
		} else {
			_smartStringAlloc(str, len_)
		}
	}
	return str.GetLen() + len_
}
func SmartStringFreeEx(str *SmartString, persistent ZendBool) {
	if str.GetC() != nil {
		Pefree(str.GetC(), persistent)
		str.SetC(nil)
	}
	str.SetLen(0)
	str.SetA(str.GetLen())
}
func SmartString0(str *SmartString) {
	if str.GetC() != nil {
		str.GetC()[str.GetLen()] = '0'
	}
}
func SmartStringAppendcEx(dest *SmartString, ch byte, persistent ZendBool) {
	dest.SetLen(SmartStringAlloc(dest, 1, persistent))
	dest.GetC()[dest.GetLen()-1] = ch
}
func SmartStringAppendlEx(dest *SmartString, str *byte, len_ int, persistent ZendBool) {
	var new_len int = SmartStringAlloc(dest, len_, persistent)
	memcpy(dest.GetC()+dest.GetLen(), str, len_)
	dest.SetLen(new_len)
}
func SmartStringAppendLongEx(dest *SmartString, num ZendLong, persistent ZendBool) {
	var buf []byte
	var result *byte = ZendPrintLongToBuf(buf+b.SizeOf("buf")-1, num)
	SmartStringAppendlEx(dest, result, buf+b.SizeOf("buf")-1-result, persistent)
}
func SmartStringAppendUnsignedEx(dest *SmartString, num ZendUlong, persistent ZendBool) {
	var buf []byte
	var result *byte = ZendPrintUlongToBuf(buf+b.SizeOf("buf")-1, num)
	SmartStringAppendlEx(dest, result, buf+b.SizeOf("buf")-1-result, persistent)
}
func SmartStringSetl(dest *SmartString, src *byte, len_ int) {
	dest.SetLen(len_)
	dest.SetA(len_ + 1)
	dest.SetC(src)
}
func SmartStringReset(str *SmartString) { str.SetLen(0) }
