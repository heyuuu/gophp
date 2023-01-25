// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_multibyte.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at                              |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Masaki Fujimoto <fujimoto@php.net>                          |
   |          Rui Hirokawa <hirokawa@php.net>                             |
   +----------------------------------------------------------------------+
*/

// #define ZEND_MULTIBYTE_H

type ZendEncoding = __struct___zend_encoding
type ZendEncodingFilter func(str **uint8, str_length *int, buf *uint8, length int) int
type ZendEncodingFetcher func(encoding_name *byte) *ZendEncoding
type ZendEncodingNameGetter func(encoding *ZendEncoding) *byte
type ZendEncodingLexerCompatibilityChecker func(encoding *ZendEncoding) int
type ZendEncodingDetector func(string *uint8, length int, list **ZendEncoding, list_size int) *ZendEncoding
type ZendEncodingConverter func(to **uint8, to_length *int, from *uint8, from_length int, encoding_to *ZendEncoding, encoding_from *ZendEncoding) int
type ZendEncodingListParser func(encoding_list *byte, encoding_list_len int, return_list ***ZendEncoding, return_size *int, persistent int) int
type ZendEncodingInternalEncodingGetter func() *ZendEncoding
type ZendEncodingInternalEncodingSetter func(encoding *ZendEncoding) int

/*
 * zend multibyte APIs
 */

/* multibyte utility functions */

// Source: <Zend/zend_multibyte.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at                              |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Masaki Fujimoto <fujimoto@php.net>                          |
   |          Rui Hirokawa <hirokawa@php.net>                             |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_compile.h"

// # include "zend_operators.h"

// # include "zend_multibyte.h"

// # include "zend_ini.h"

func DummyEncodingFetcher(encoding_name *byte) *ZendEncoding            { return nil }
func DummyEncodingNameGetter(encoding *ZendEncoding) *byte              { return (*byte)(encoding) }
func DummyEncodingLexerCompatibilityChecker(encoding *ZendEncoding) int { return 0 }
func DummyEncodingDetector(string *uint8, length int, list **ZendEncoding, list_size int) *ZendEncoding {
	return nil
}
func DummyEncodingConverter(to **uint8, to_length *int, from *uint8, from_length int, encoding_to *ZendEncoding, encoding_from *ZendEncoding) int {
	return size_t - 1
}
func DummyEncodingListParser(encoding_list *byte, encoding_list_len int, return_list ***ZendEncoding, return_size *int, persistent int) int {
	*return_list = Pemalloc(0, persistent)
	*return_size = 0
	return SUCCESS
}
func DummyInternalEncodingGetter() *ZendEncoding             { return nil }
func DummyInternalEncodingSetter(encoding *ZendEncoding) int { return FAILURE }

var MultibyteFunctionsDummy ZendMultibyteFunctions
var MultibyteFunctions ZendMultibyteFunctions = ZendMultibyteFunctions{nil, DummyEncodingFetcher, DummyEncodingNameGetter, DummyEncodingLexerCompatibilityChecker, DummyEncodingDetector, DummyEncodingConverter, DummyEncodingListParser, DummyInternalEncodingGetter, DummyInternalEncodingSetter}
var ZendMultibyteEncodingUtf32be *ZendEncoding = (*ZendEncoding)("UTF-32BE")
var ZendMultibyteEncodingUtf32le *ZendEncoding = (*ZendEncoding)("UTF-32LE")
var ZendMultibyteEncodingUtf16be *ZendEncoding = (*ZendEncoding)("UTF-16BE")
var ZendMultibyteEncodingUtf16le *ZendEncoding = (*ZendEncoding)("UTF-32LE")
var ZendMultibyteEncodingUtf8 *ZendEncoding = (*ZendEncoding)("UTF-8")

func ZendMultibyteSetFunctions(functions *ZendMultibyteFunctions) int {
	ZendMultibyteEncodingUtf32be = functions.GetEncodingFetcher()("UTF-32BE")
	if ZendMultibyteEncodingUtf32be == nil {
		return FAILURE
	}
	ZendMultibyteEncodingUtf32le = functions.GetEncodingFetcher()("UTF-32LE")
	if ZendMultibyteEncodingUtf32le == nil {
		return FAILURE
	}
	ZendMultibyteEncodingUtf16be = functions.GetEncodingFetcher()("UTF-16BE")
	if ZendMultibyteEncodingUtf16be == nil {
		return FAILURE
	}
	ZendMultibyteEncodingUtf16le = functions.GetEncodingFetcher()("UTF-16LE")
	if ZendMultibyteEncodingUtf16le == nil {
		return FAILURE
	}
	ZendMultibyteEncodingUtf8 = functions.GetEncodingFetcher()("UTF-8")
	if ZendMultibyteEncodingUtf8 == nil {
		return FAILURE
	}
	MultibyteFunctionsDummy = MultibyteFunctions
	MultibyteFunctions = *functions

	/* As zend_multibyte_set_functions() gets called after ini settings were
	 * populated, we need to reinitialize script_encoding here.
	 */

	var value *byte = ZendIniString("zend.script_encoding", b.SizeOf("\"zend.script_encoding\"")-1, 0)
	ZendMultibyteSetScriptEncodingByString(value, strlen(value))
	return SUCCESS
}
func ZendMultibyteRestoreFunctions() {
	MultibyteFunctions = MultibyteFunctionsDummy
}
func ZendMultibyteGetFunctions() *ZendMultibyteFunctions {
	if MultibyteFunctions.GetProviderName() != nil {
		return &MultibyteFunctions
	} else {
		return nil
	}
}
func ZendMultibyteFetchEncoding(name *byte) *ZendEncoding {
	return MultibyteFunctions.GetEncodingFetcher()(name)
}
func ZendMultibyteGetEncodingName(encoding *ZendEncoding) *byte {
	return MultibyteFunctions.GetEncodingNameGetter()(encoding)
}
func ZendMultibyteCheckLexerCompatibility(encoding *ZendEncoding) int {
	return MultibyteFunctions.GetLexerCompatibilityChecker()(encoding)
}
func ZendMultibyteEncodingDetector(string *uint8, length int, list **ZendEncoding, list_size int) *ZendEncoding {
	return MultibyteFunctions.GetEncodingDetector()(string, length, list, list_size)
}
func ZendMultibyteEncodingConverter(to **uint8, to_length *int, from *uint8, from_length int, encoding_to *ZendEncoding, encoding_from *ZendEncoding) int {
	return MultibyteFunctions.GetEncodingConverter()(to, to_length, from, from_length, encoding_to, encoding_from)
}
func ZendMultibyteParseEncodingList(encoding_list *byte, encoding_list_len int, return_list ***ZendEncoding, return_size *int, persistent int) int {
	return MultibyteFunctions.GetEncodingListParser()(encoding_list, encoding_list_len, return_list, return_size, persistent)
}
func ZendMultibyteGetInternalEncoding() *ZendEncoding {
	return MultibyteFunctions.GetInternalEncodingGetter()()
}
func ZendMultibyteGetScriptEncoding() *ZendEncoding {
	return LanguageScannerGlobals.GetScriptEncoding()
}
func ZendMultibyteSetScriptEncoding(encoding_list **ZendEncoding, encoding_list_size int) int {
	if CompilerGlobals.GetScriptEncodingList() != nil {
		Free((*byte)(CompilerGlobals.GetScriptEncodingList()))
	}
	CompilerGlobals.SetScriptEncodingList(encoding_list)
	CompilerGlobals.SetScriptEncodingListSize(encoding_list_size)
	return SUCCESS
}
func ZendMultibyteSetInternalEncoding(encoding *ZendEncoding) int {
	return MultibyteFunctions.GetInternalEncodingSetter()(encoding)
}
func ZendMultibyteSetScriptEncodingByString(new_value *byte, new_value_length int) int {
	var list **ZendEncoding = 0
	var size int = 0
	if new_value == nil {
		ZendMultibyteSetScriptEncoding(nil, 0)
		return SUCCESS
	}
	if FAILURE == ZendMultibyteParseEncodingList(new_value, new_value_length, &list, &size, 1) {
		return FAILURE
	}
	if size == 0 {
		Pefree(any(list), 1)
		return FAILURE
	}
	if FAILURE == ZendMultibyteSetScriptEncoding(list, size) {
		return FAILURE
	}
	return SUCCESS
}
