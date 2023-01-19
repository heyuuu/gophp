// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/sha1.h>

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
   | Author: Stefan Esser <sesser@php.net>                                |
   +----------------------------------------------------------------------+
*/

// #define SHA1_H

// # include "ext/standard/basic_functions.h"

/* SHA1 context. */

// @type PHP_SHA1_CTX struct

// Source: <ext/standard/sha1.c>

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
   | Author: Stefan Esser <sesser@php.net>                                |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

/* This code is heavily based on the PHP md5 implementation */

// # include "sha1.h"

// # include "md5.h"

func MakeSha1Digest(sha1str *byte, digest *uint8) { MakeDigestEx(sha1str, digest, 20) }

/* {{{ proto string sha1(string str [, bool raw_output])
   Calculate the sha1 hash of a string */

func ZifSha1(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *zend.ZendString
	var raw_output zend.ZendBool = 0
	var context PHP_SHA1_CTX
	var digest []uint8
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &arg, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &raw_output, &_dummy, 0) == 0 {
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
	PHP_SHA1Init(&context)
	PHP_SHA1Update(&context, (*uint8)(arg.val), arg.len_)
	PHP_SHA1Final(digest, &context)
	if raw_output != 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit((*byte)(digest), 20, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringAlloc(40, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		MakeDigestEx(return_value.value.str.val, digest, 20)
	}
}

/* }}} */

func ZifSha1File(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *byte
	var arg_len int
	var raw_output zend.ZendBool = 0
	var buf []uint8
	var digest []uint8
	var context PHP_SHA1_CTX
	var n ssize_t
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &arg, &arg_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &raw_output, &_dummy, 0) == 0 {
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
	stream = streams._phpStreamOpenWrapperEx(arg, "rb", 0x8, nil, nil)
	if stream == nil {
		return_value.u1.type_info = 2
		return
	}
	PHP_SHA1Init(&context)
	for g.Assign(&n, streams._phpStreamRead(stream, (*byte)(buf), g.SizeOf("buf"))) > 0 {
		PHP_SHA1Update(&context, buf, n)
	}
	PHP_SHA1Final(digest, &context)
	streams._phpStreamFree(stream, 1|2)
	if raw_output != 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit((*byte)(digest), 20, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringAlloc(40, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		MakeDigestEx(return_value.value.str.val, digest, 20)
	}
}

/* }}} */

var SHA1Transform func([]uint32, []uint8)
var SHA1Encode func(*uint8, *uint32, uint)
var SHA1Decode func(*uint32, *uint8, uint)
var PADDING []uint8 = []uint8{0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

/* F, G, H and I are basic SHA1 functions.
 */

// #define F(x,y,z) ( ( z ) ^ ( ( x ) & ( ( y ) ^ ( z ) ) ) )

// #define G(x,y,z) ( ( x ) ^ ( y ) ^ ( z ) )

// #define H(x,y,z) ( ( ( x ) & ( y ) ) | ( ( z ) & ( ( x ) | ( y ) ) ) )

// #define I(x,y,z) ( ( x ) ^ ( y ) ^ ( z ) )

/* ROTATE_LEFT rotates x left n bits.
 */

// #define ROTATE_LEFT(x,n) ( ( ( x ) << ( n ) ) | ( ( x ) >> ( 32 - ( n ) ) ) )

/* W[i]
 */

// #define W(i) ( tmp = x [ ( i - 3 ) & 15 ] ^ x [ ( i - 8 ) & 15 ] ^ x [ ( i - 14 ) & 15 ] ^ x [ i & 15 ] , ( x [ i & 15 ] = ROTATE_LEFT ( tmp , 1 ) ) )

/* FF, GG, HH, and II transformations for rounds 1, 2, 3, and 4.
 */

// #define FF(a,b,c,d,e,w) { ( e ) += F ( ( b ) , ( c ) , ( d ) ) + ( w ) + ( uint32_t ) ( 0x5A827999 ) ; ( e ) += ROTATE_LEFT ( ( a ) , 5 ) ; ( b ) = ROTATE_LEFT ( ( b ) , 30 ) ; }

// #define GG(a,b,c,d,e,w) { ( e ) += G ( ( b ) , ( c ) , ( d ) ) + ( w ) + ( uint32_t ) ( 0x6ED9EBA1 ) ; ( e ) += ROTATE_LEFT ( ( a ) , 5 ) ; ( b ) = ROTATE_LEFT ( ( b ) , 30 ) ; }

// #define HH(a,b,c,d,e,w) { ( e ) += H ( ( b ) , ( c ) , ( d ) ) + ( w ) + ( uint32_t ) ( 0x8F1BBCDC ) ; ( e ) += ROTATE_LEFT ( ( a ) , 5 ) ; ( b ) = ROTATE_LEFT ( ( b ) , 30 ) ; }

// #define II(a,b,c,d,e,w) { ( e ) += I ( ( b ) , ( c ) , ( d ) ) + ( w ) + ( uint32_t ) ( 0xCA62C1D6 ) ; ( e ) += ROTATE_LEFT ( ( a ) , 5 ) ; ( b ) = ROTATE_LEFT ( ( b ) , 30 ) ; }

/* {{{ PHP_SHA1Init
 * SHA1 initialization. Begins an SHA1 operation, writing a new context.
 */

func PHP_SHA1Init(context *PHP_SHA1_CTX) {
	context.GetCount()[1] = 0
	context.GetCount()[0] = context.GetCount()[1]

	/* Load magic initialization constants.
	 */

	context.GetState()[0] = 0x67452301
	context.GetState()[1] = 0xefcdab89
	context.GetState()[2] = 0x98badcfe
	context.GetState()[3] = 0x10325476
	context.GetState()[4] = 0xc3d2e1f0
}

/* }}} */

func PHP_SHA1Update(context *PHP_SHA1_CTX, input *uint8, inputLen int) {
	var i uint
	var index uint
	var partLen uint

	/* Compute number of bytes mod 64 */

	index = uint(context.GetCount()[0] >> 3 & 0x3f)

	/* Update number of bits */

	if g.AssignOp(&context.GetCount()[0], "+=", uint32(inputLen<<3)) < uint32(inputLen<<3) {
		context.GetCount()[1]++
	}
	context.GetCount()[1] += uint32(inputLen >> 29)
	partLen = 64 - index

	/* Transform as many times as possible.
	 */

	if inputLen >= partLen {
		memcpy((*uint8)(&context.buffer[index]), (*uint8)(input), partLen)
		SHA1Transform(context.GetState(), context.GetBuffer())
		for i = partLen; i+63 < inputLen; i += 64 {
			SHA1Transform(context.GetState(), &input[i])
		}
		index = 0
	} else {
		i = 0
	}

	/* Buffer remaining input */

	memcpy((*uint8)(&context.buffer[index]), (*uint8)(&input[i]), inputLen-i)

	/* Buffer remaining input */
}

/* }}} */

func PHP_SHA1Final(digest []uint8, context *PHP_SHA1_CTX) {
	var bits []uint8
	var index uint
	var padLen uint

	/* Save number of bits */

	bits[7] = context.GetCount()[0] & 0xff
	bits[6] = context.GetCount()[0] >> 8 & 0xff
	bits[5] = context.GetCount()[0] >> 16 & 0xff
	bits[4] = context.GetCount()[0] >> 24 & 0xff
	bits[3] = context.GetCount()[1] & 0xff
	bits[2] = context.GetCount()[1] >> 8 & 0xff
	bits[1] = context.GetCount()[1] >> 16 & 0xff
	bits[0] = context.GetCount()[1] >> 24 & 0xff

	/* Pad out to 56 mod 64.
	 */

	index = uint(context.GetCount()[0] >> 3 & 0x3f)
	if index < 56 {
		padLen = 56 - index
	} else {
		padLen = 120 - index
	}
	PHP_SHA1Update(context, PADDING, padLen)

	/* Append length (before padding) */

	PHP_SHA1Update(context, bits, 8)

	/* Store state in digest */

	SHA1Encode(digest, context.GetState(), 20)

	/* Zeroize sensitive information.
	 */

	core.PhpExplicitBzero((*uint8)(context), g.SizeOf("* context"))

	/* Zeroize sensitive information.
	 */
}

/* }}} */
