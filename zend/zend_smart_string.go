// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
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

// #define smart_string_appends_ex(str,src,what) smart_string_appendl_ex ( ( str ) , ( src ) , strlen ( src ) , ( what ) )

// #define smart_string_appends(str,src) smart_string_appendl ( ( str ) , ( src ) , strlen ( src ) )

// #define smart_string_append_ex(str,src,what) smart_string_appendl_ex ( ( str ) , ( ( smart_string * ) ( src ) ) -> c , ( ( smart_string * ) ( src ) ) -> len , ( what ) ) ;

// #define smart_string_sets(str,src) smart_string_setl ( ( str ) , ( src ) , strlen ( src ) ) ;

// #define smart_string_appendc(str,c) smart_string_appendc_ex ( ( str ) , ( c ) , 0 )

// #define smart_string_free(s) smart_string_free_ex ( ( s ) , 0 )

// #define smart_string_appendl(str,src,len) smart_string_appendl_ex ( ( str ) , ( src ) , ( len ) , 0 )

// #define smart_string_append(str,src) smart_string_append_ex ( ( str ) , ( src ) , 0 )

// #define smart_string_append_long(str,val) smart_string_append_long_ex ( ( str ) , ( val ) , 0 )

// #define smart_string_append_unsigned(str,val) smart_string_append_unsigned_ex ( ( str ) , ( val ) , 0 )

func SmartStringAlloc(str *SmartString, len_ int, persistent ZendBool) int {
	if str.GetC() == nil || len_ >= str.GetA()-str.GetLen() {
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
		g.CondF(persistent != 0, func() { return Free(str.GetC()) }, func() { return _efree(str.GetC()) })
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
	var result *byte = ZendPrintLongToBuf(buf+g.SizeOf("buf")-1, num)
	SmartStringAppendlEx(dest, result, buf+g.SizeOf("buf")-1-result, persistent)
}
func SmartStringAppendUnsignedEx(dest *SmartString, num ZendUlong, persistent ZendBool) {
	var buf []byte
	var result *byte = ZendPrintUlongToBuf(buf+g.SizeOf("buf")-1, num)
	SmartStringAppendlEx(dest, result, buf+g.SizeOf("buf")-1-result, persistent)
}
func SmartStringSetl(dest *SmartString, src *byte, len_ int) {
	dest.SetLen(len_)
	dest.SetA(len_ + 1)
	dest.SetC(src)
}
func SmartStringReset(str *SmartString) { str.SetLen(0) }
