// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
)

// Source: <ext/standard/base64.h>

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
   | Author: Jim Winstead <jimw@php.net>                                  |
   |         Xinchen Hui <laruence@php.net>                               |
   +----------------------------------------------------------------------+
*/

// #define BASE64_H

/*
 * NEON implementation is based on https://github.com/WojciechMula/base64simd
 * which is copyrighted to:
 * Copyright (c) 2015-2018, Wojciech Mula
 * All rights reserved.
 *
 * SSSE3 and AVX2 implementation are based on https://github.com/aklomp/base64
 * which is copyrighted to:
 *
 * Copyright (c) 2005-2007, Nick Galbreath
 * Copyright (c) 2013-2017, Alfred Klomp
 * Copyright (c) 2015-2017, Wojciech Mula
 * Copyright (c) 2016-2017, Matthieu Darbois
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 * - Redistributions of source code must retain the above copyright notice,
 *   this list of conditions and the following disclaimer.
 *
 * - Redistributions in binary form must reproduce the above copyright
 *   notice, this list of conditions and the following disclaimer in the
 *   documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS
 * IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED
 * TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A
 * PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED
 * TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 * PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 * LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 * NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

func PhpBase64EncodeStr(str *zend.ZendString) *zend.ZendString {
	return PhpBase64Encode((*uint8)(zend.ZSTR_VAL(str)), zend.ZSTR_LEN(str))
}
func PhpBase64Decode(str *uint8, len_ int) *zend.ZendString { return PhpBase64DecodeEx(str, len_, 0) }
func PhpBase64DecodeStr(str *zend.ZendString) *zend.ZendString {
	return PhpBase64DecodeEx((*uint8)(zend.ZSTR_VAL(str)), zend.ZSTR_LEN(str), 0)
}

// Source: <ext/standard/base64.c>

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
   | Author: Jim Winstead <jimw@php.net>                                  |
   |         Xinchen Hui <laruence@php.net>                               |
   +----------------------------------------------------------------------+
*/

// # include < string . h >

// # include "php.h"

// # include "base64.h"

/* {{{ base64 tables */

var Base64Table []byte = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/', '0'}
var Base64Pad byte = '='
var Base64ReverseTable []short = []short{-2, -2, -2, -2, -2, -2, -2, -2, -2, -1, -1, -2, -2, -1, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -1, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, 62, -2, -2, -2, 63, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, -2, -2, -2, -2, -2, -2, -2, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, -2, -2, -2, -2, -2, -2, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2}

/* }}} */

func PhpBase64EncodeImpl(in *uint8, inl int, out *uint8) *uint8 {
	for inl > 2 {
		b.PostInc(&(*out)) = Base64Table[in[0]>>2]
		b.PostInc(&(*out)) = Base64Table[((in[0]&0x3)<<4)+(in[1]>>4)]
		b.PostInc(&(*out)) = Base64Table[((in[1]&0xf)<<2)+(in[2]>>6)]
		b.PostInc(&(*out)) = Base64Table[in[2]&0x3f]
		in += 3
		inl -= 3
	}

	/* now deal with the tail end of things */

	if inl != 0 {
		b.PostInc(&(*out)) = Base64Table[in[0]>>2]
		if inl > 1 {
			b.PostInc(&(*out)) = Base64Table[((in[0]&0x3)<<4)+(in[1]>>4)]
			b.PostInc(&(*out)) = Base64Table[(in[1]&0xf)<<2]
			b.PostInc(&(*out)) = Base64Pad
		} else {
			b.PostInc(&(*out)) = Base64Table[(in[0]&0x3)<<4]
			b.PostInc(&(*out)) = Base64Pad
			b.PostInc(&(*out)) = Base64Pad
		}
	}
	*out = '0'
	return out
}

/* }}} */

func PhpBase64DecodeImpl(in *uint8, inl int, out *uint8, outl *int, strict zend.ZendBool) int {
	var ch int
	var i int = 0
	var padding int = 0
	var j int = *outl

	/* run through the whole string, converting as we go */

	for b.PostDec(&inl) > 0 {
		*in++
		ch = (*in) - 1
		if ch == Base64Pad {
			padding++
			continue
		}
		ch = Base64ReverseTable[ch]
		if strict == 0 {

			/* skip unknown characters and whitespace */

			if ch < 0 {
				continue
			}

			/* skip unknown characters and whitespace */

		} else {

			/* skip whitespace */

			if ch == -1 {
				continue
			}

			/* fail on bad characters or if any data follows padding */

			if ch == -2 || padding != 0 {
				goto fail
			}

			/* fail on bad characters or if any data follows padding */

		}
		switch i % 4 {
		case 0:
			out[j] = ch << 2
			break
		case 1:
			out[b.PostInc(&j)] |= ch >> 4
			out[j] = (ch & 0xf) << 4
			break
		case 2:
			out[b.PostInc(&j)] |= ch >> 2
			out[j] = (ch & 0x3) << 6
			break
		case 3:
			out[b.PostInc(&j)] |= ch
			break
		}
		i++
	}

	/* fail if the input is truncated (only one char in last group) */

	if strict != 0 && i%4 == 1 {
		goto fail
	}

	/* fail if the padding length is wrong (not VV==, VVV=), but accept zero padding
	 * RFC 4648: "In some circumstances, the use of padding [--] is not required" */

	if strict != 0 && padding != 0 && (padding > 2 || (i+padding)%4 != 0) {
		goto fail
	}
	*outl = j
	out[j] = '0'
	return 1
fail:
	return 0
}

/* }}} */

/* }}} */

func PhpBase64Encode(str *uint8, length int) *zend.ZendString {
	var p *uint8
	var result *zend.ZendString
	result = zend.ZendStringSafeAlloc((length+2)/3, 4*b.SizeOf("char"), 0, 0)
	p = (*uint8)(zend.ZSTR_VAL(result))
	p = PhpBase64EncodeImpl(str, length, p)
	zend.ZSTR_LEN(result) = p - (*uint8)(zend.ZSTR_VAL(result))
	return result
}
func PhpBase64DecodeEx(str *uint8, length int, strict zend.ZendBool) *zend.ZendString {
	var result *zend.ZendString
	var outl int = 0
	result = zend.ZendStringAlloc(length, 0)
	if PhpBase64DecodeImpl(str, length, (*uint8)(zend.ZSTR_VAL(result)), &outl, strict) == 0 {
		zend.ZendStringEfree(result)
		return nil
	}
	zend.ZSTR_LEN(result) = outl
	return result
}

/* }}} */

func ZifBase64Encode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte
	var str_len int
	var result *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
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
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
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
	result = PhpBase64Encode((*uint8)(str), str_len)
	zend.RETVAL_STR(result)
	return
}

/* }}} */

func ZifBase64Decode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte
	var strict zend.ZendBool = 0
	var str_len int
	var result *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &strict, &_dummy, 0) == 0) {
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
	result = PhpBase64DecodeEx((*uint8)(str), str_len, strict)
	if result != nil {
		zend.RETVAL_STR(result)
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}

/* }}} */
