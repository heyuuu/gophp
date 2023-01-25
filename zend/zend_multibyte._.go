// <<generate>>

package zend

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

var MultibyteFunctionsDummy ZendMultibyteFunctions
var MultibyteFunctions ZendMultibyteFunctions = ZendMultibyteFunctions{nil, DummyEncodingFetcher, DummyEncodingNameGetter, DummyEncodingLexerCompatibilityChecker, DummyEncodingDetector, DummyEncodingConverter, DummyEncodingListParser, DummyInternalEncodingGetter, DummyInternalEncodingSetter}
var ZendMultibyteEncodingUtf32be *ZendEncoding = (*ZendEncoding)("UTF-32BE")
var ZendMultibyteEncodingUtf32le *ZendEncoding = (*ZendEncoding)("UTF-32LE")
var ZendMultibyteEncodingUtf16be *ZendEncoding = (*ZendEncoding)("UTF-16BE")
var ZendMultibyteEncodingUtf16le *ZendEncoding = (*ZendEncoding)("UTF-32LE")
var ZendMultibyteEncodingUtf8 *ZendEncoding = (*ZendEncoding)("UTF-8")
