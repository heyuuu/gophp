// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/md5.h>

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
   | Author: Alexander Peslyak (Solar Designer) <solar at openwall.com>   |
   |         Rasmus Lerdorf <rasmus@lerdorf.on.ca>                        |
   +----------------------------------------------------------------------+
*/

// #define MD5_H

// # include "ext/standard/basic_functions.h"

/*
 * This is an OpenSSL-compatible implementation of the RSA Data Security,
 * Inc. MD5 Message-Digest Algorithm (RFC 1321).
 *
 * Written by Solar Designer <solar at openwall.com> in 2001, and placed
 * in the public domain.  There's absolutely no warranty.
 *
 * See md5.c for more information.
 */

// @type PHP_MD5_CTX struct

// Source: <ext/standard/md5.c>

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
   | Author: Alexander Peslyak (Solar Designer) <solar at openwall.com>   |
   |         Lachlan Roche                                                |
   |         Alessandro Astarita <aleast@capri.it>                        |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "md5.h"

func MakeDigest(md5str *byte, digest *uint8) { MakeDigestEx(md5str, digest, 16) }

/* }}} */

func MakeDigestEx(md5str *byte, digest *uint8, len_ int) {
	var hexits []byte = "0123456789abcdef"
	var i int
	for i = 0; i < len_; i++ {
		md5str[i*2] = hexits[digest[i]>>4]
		md5str[i*2+1] = hexits[digest[i]&0xf]
	}
	md5str[len_*2] = '0'
}

/* }}} */

func PhpIfMd5(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *zend.ZendString
	var raw_output zend.ZendBool = 0
	var context PHP_MD5_CTX
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
	PHP_MD5Init(&context)
	PHP_MD5Update(&context, arg.val, arg.len_)
	PHP_MD5Final(digest, &context)
	if raw_output != 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit((*byte)(digest), 16, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringAlloc(32, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		MakeDigestEx(return_value.value.str.val, digest, 16)
	}
}

/* }}} */

func PhpIfMd5File(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *byte
	var arg_len int
	var raw_output zend.ZendBool = 0
	var buf []uint8
	var digest []uint8
	var context PHP_MD5_CTX
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
	PHP_MD5Init(&context)
	for g.Assign(&n, streams._phpStreamRead(stream, (*byte)(buf), g.SizeOf("buf"))) > 0 {
		PHP_MD5Update(&context, buf, n)
	}

	/* XXX this probably can be improved with some number of retries */

	if streams._phpStreamEof(stream) == 0 {
		streams._phpStreamFree(stream, 1|2)
		PHP_MD5Final(digest, &context)
		return_value.u1.type_info = 2
		return
	}
	streams._phpStreamFree(stream, 1|2)
	PHP_MD5Final(digest, &context)
	if raw_output != 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit((*byte)(digest), 16, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringAlloc(32, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		MakeDigestEx(return_value.value.str.val, digest, 16)
	}
}

/* }}} */

// # include < string . h >

/*
 * The basic MD5 functions.
 *
 * F and G are optimized compared to their RFC 1321 definitions for
 * architectures that lack an AND-NOT instruction, just like in Colin Plumb's
 * implementation.
 */

// #define F(x,y,z) ( ( z ) ^ ( ( x ) & ( ( y ) ^ ( z ) ) ) )

// #define G(x,y,z) ( ( y ) ^ ( ( z ) & ( ( x ) ^ ( y ) ) ) )

// #define H(x,y,z) ( ( x ) ^ ( y ) ^ ( z ) )

// #define I(x,y,z) ( ( y ) ^ ( ( x ) | ~ ( z ) ) )

/*
 * The MD5 transformation for all four rounds.
 */

// #define STEP(f,a,b,c,d,x,t,s) ( a ) += f ( ( b ) , ( c ) , ( d ) ) + ( x ) + ( t ) ; ( a ) = ( ( ( a ) << ( s ) ) | ( ( ( a ) & 0xffffffff ) >> ( 32 - ( s ) ) ) ) ; ( a ) += ( b ) ;

/*
 * SET reads 4 input bytes in little-endian byte order and stores them
 * in a properly aligned word in host byte order.
 *
 * The check for little-endian architectures that tolerate unaligned
 * memory accesses is just an optimization.  Nothing will break if it
 * doesn't work.
 */

// #define SET(n) ( ctx -> block [ ( n ) ] = ( uint32_t ) ptr [ ( n ) * 4 ] | ( ( uint32_t ) ptr [ ( n ) * 4 + 1 ] << 8 ) | ( ( uint32_t ) ptr [ ( n ) * 4 + 2 ] << 16 ) | ( ( uint32_t ) ptr [ ( n ) * 4 + 3 ] << 24 ) )

// #define GET(n) ( ctx -> block [ ( n ) ] )

/*
 * This processes one or more 64-byte data blocks, but does NOT update
 * the bit counters.  There are no alignment requirements.
 */

func Body(ctx *PHP_MD5_CTX, data any, size int) any {
	var ptr *uint8
	var a uint32
	var b uint32
	var c uint32
	var d uint32
	var saved_a uint32
	var saved_b uint32
	var saved_c uint32
	var saved_d uint32
	ptr = data
	a = ctx.GetA()
	b = ctx.GetB()
	c = ctx.GetC()
	d = ctx.GetD()
	for {
		saved_a = a
		saved_b = b
		saved_c = c
		saved_d = d

		/* Round 1 */

		a += (d ^ b&(c^d)) + g.Assign(&ctx.GetBlock()[0], uint32(ptr[0*4]|uint32(ptr[0*4+1]<<8)|uint32(ptr[0*4+2]<<16)|uint32(ptr[0*4+3]<<24))) + 0xd76aa478
		a = a<<7 | (a&0xffffffff)>>32 - 7
		a += b
		d += (c ^ a&(b^c)) + g.Assign(&ctx.GetBlock()[1], uint32(ptr[1*4]|uint32(ptr[1*4+1]<<8)|uint32(ptr[1*4+2]<<16)|uint32(ptr[1*4+3]<<24))) + 0xe8c7b756
		d = d<<12 | (d&0xffffffff)>>32 - 12
		d += a
		c += (b ^ d&(a^b)) + g.Assign(&ctx.GetBlock()[2], uint32(ptr[2*4]|uint32(ptr[2*4+1]<<8)|uint32(ptr[2*4+2]<<16)|uint32(ptr[2*4+3]<<24))) + 0x242070db
		c = c<<17 | (c&0xffffffff)>>32 - 17
		c += d
		b += (a ^ c&(d^a)) + g.Assign(&ctx.GetBlock()[3], uint32(ptr[3*4]|uint32(ptr[3*4+1]<<8)|uint32(ptr[3*4+2]<<16)|uint32(ptr[3*4+3]<<24))) + 0xc1bdceee
		b = b<<22 | (b&0xffffffff)>>32 - 22
		b += c
		a += (d ^ b&(c^d)) + g.Assign(&ctx.GetBlock()[4], uint32(ptr[4*4]|uint32(ptr[4*4+1]<<8)|uint32(ptr[4*4+2]<<16)|uint32(ptr[4*4+3]<<24))) + 0xf57c0faf
		a = a<<7 | (a&0xffffffff)>>32 - 7
		a += b
		d += (c ^ a&(b^c)) + g.Assign(&ctx.GetBlock()[5], uint32(ptr[5*4]|uint32(ptr[5*4+1]<<8)|uint32(ptr[5*4+2]<<16)|uint32(ptr[5*4+3]<<24))) + 0x4787c62a
		d = d<<12 | (d&0xffffffff)>>32 - 12
		d += a
		c += (b ^ d&(a^b)) + g.Assign(&ctx.GetBlock()[6], uint32(ptr[6*4]|uint32(ptr[6*4+1]<<8)|uint32(ptr[6*4+2]<<16)|uint32(ptr[6*4+3]<<24))) + 0xa8304613
		c = c<<17 | (c&0xffffffff)>>32 - 17
		c += d
		b += (a ^ c&(d^a)) + g.Assign(&ctx.GetBlock()[7], uint32(ptr[7*4]|uint32(ptr[7*4+1]<<8)|uint32(ptr[7*4+2]<<16)|uint32(ptr[7*4+3]<<24))) + 0xfd469501
		b = b<<22 | (b&0xffffffff)>>32 - 22
		b += c
		a += (d ^ b&(c^d)) + g.Assign(&ctx.GetBlock()[8], uint32(ptr[8*4]|uint32(ptr[8*4+1]<<8)|uint32(ptr[8*4+2]<<16)|uint32(ptr[8*4+3]<<24))) + 0x698098d8
		a = a<<7 | (a&0xffffffff)>>32 - 7
		a += b
		d += (c ^ a&(b^c)) + g.Assign(&ctx.GetBlock()[9], uint32(ptr[9*4]|uint32(ptr[9*4+1]<<8)|uint32(ptr[9*4+2]<<16)|uint32(ptr[9*4+3]<<24))) + 0x8b44f7af
		d = d<<12 | (d&0xffffffff)>>32 - 12
		d += a
		c += (b ^ d&(a^b)) + g.Assign(&ctx.GetBlock()[10], uint32(ptr[10*4]|uint32(ptr[10*4+1]<<8)|uint32(ptr[10*4+2]<<16)|uint32(ptr[10*4+3]<<24))) + 0xffff5bb1
		c = c<<17 | (c&0xffffffff)>>32 - 17
		c += d
		b += (a ^ c&(d^a)) + g.Assign(&ctx.GetBlock()[11], uint32(ptr[11*4]|uint32(ptr[11*4+1]<<8)|uint32(ptr[11*4+2]<<16)|uint32(ptr[11*4+3]<<24))) + 0x895cd7be
		b = b<<22 | (b&0xffffffff)>>32 - 22
		b += c
		a += (d ^ b&(c^d)) + g.Assign(&ctx.GetBlock()[12], uint32(ptr[12*4]|uint32(ptr[12*4+1]<<8)|uint32(ptr[12*4+2]<<16)|uint32(ptr[12*4+3]<<24))) + 0x6b901122
		a = a<<7 | (a&0xffffffff)>>32 - 7
		a += b
		d += (c ^ a&(b^c)) + g.Assign(&ctx.GetBlock()[13], uint32(ptr[13*4]|uint32(ptr[13*4+1]<<8)|uint32(ptr[13*4+2]<<16)|uint32(ptr[13*4+3]<<24))) + 0xfd987193
		d = d<<12 | (d&0xffffffff)>>32 - 12
		d += a
		c += (b ^ d&(a^b)) + g.Assign(&ctx.GetBlock()[14], uint32(ptr[14*4]|uint32(ptr[14*4+1]<<8)|uint32(ptr[14*4+2]<<16)|uint32(ptr[14*4+3]<<24))) + 0xa679438e
		c = c<<17 | (c&0xffffffff)>>32 - 17
		c += d
		b += (a ^ c&(d^a)) + g.Assign(&ctx.GetBlock()[15], uint32(ptr[15*4]|uint32(ptr[15*4+1]<<8)|uint32(ptr[15*4+2]<<16)|uint32(ptr[15*4+3]<<24))) + 0x49b40821
		b = b<<22 | (b&0xffffffff)>>32 - 22
		b += c

		/* Round 2 */

		a += (c ^ d&(b^c)) + ctx.GetBlock()[1] + 0xf61e2562
		a = a<<5 | (a&0xffffffff)>>32 - 5
		a += b
		d += (b ^ c&(a^b)) + ctx.GetBlock()[6] + 0xc040b340
		d = d<<9 | (d&0xffffffff)>>32 - 9
		d += a
		c += (a ^ b&(d^a)) + ctx.GetBlock()[11] + 0x265e5a51
		c = c<<14 | (c&0xffffffff)>>32 - 14
		c += d
		b += (d ^ a&(c^d)) + ctx.GetBlock()[0] + 0xe9b6c7aa
		b = b<<20 | (b&0xffffffff)>>32 - 20
		b += c
		a += (c ^ d&(b^c)) + ctx.GetBlock()[5] + 0xd62f105d
		a = a<<5 | (a&0xffffffff)>>32 - 5
		a += b
		d += (b ^ c&(a^b)) + ctx.GetBlock()[10] + 0x2441453
		d = d<<9 | (d&0xffffffff)>>32 - 9
		d += a
		c += (a ^ b&(d^a)) + ctx.GetBlock()[15] + 0xd8a1e681
		c = c<<14 | (c&0xffffffff)>>32 - 14
		c += d
		b += (d ^ a&(c^d)) + ctx.GetBlock()[4] + 0xe7d3fbc8
		b = b<<20 | (b&0xffffffff)>>32 - 20
		b += c
		a += (c ^ d&(b^c)) + ctx.GetBlock()[9] + 0x21e1cde6
		a = a<<5 | (a&0xffffffff)>>32 - 5
		a += b
		d += (b ^ c&(a^b)) + ctx.GetBlock()[14] + 0xc33707d6
		d = d<<9 | (d&0xffffffff)>>32 - 9
		d += a
		c += (a ^ b&(d^a)) + ctx.GetBlock()[3] + 0xf4d50d87
		c = c<<14 | (c&0xffffffff)>>32 - 14
		c += d
		b += (d ^ a&(c^d)) + ctx.GetBlock()[8] + 0x455a14ed
		b = b<<20 | (b&0xffffffff)>>32 - 20
		b += c
		a += (c ^ d&(b^c)) + ctx.GetBlock()[13] + 0xa9e3e905
		a = a<<5 | (a&0xffffffff)>>32 - 5
		a += b
		d += (b ^ c&(a^b)) + ctx.GetBlock()[2] + 0xfcefa3f8
		d = d<<9 | (d&0xffffffff)>>32 - 9
		d += a
		c += (a ^ b&(d^a)) + ctx.GetBlock()[7] + 0x676f02d9
		c = c<<14 | (c&0xffffffff)>>32 - 14
		c += d
		b += (d ^ a&(c^d)) + ctx.GetBlock()[12] + 0x8d2a4c8a
		b = b<<20 | (b&0xffffffff)>>32 - 20
		b += c

		/* Round 3 */

		a += (b ^ c ^ d) + ctx.GetBlock()[5] + 0xfffa3942
		a = a<<4 | (a&0xffffffff)>>32 - 4
		a += b
		d += (a ^ b ^ c) + ctx.GetBlock()[8] + 0x8771f681
		d = d<<11 | (d&0xffffffff)>>32 - 11
		d += a
		c += (d ^ a ^ b) + ctx.GetBlock()[11] + 0x6d9d6122
		c = c<<16 | (c&0xffffffff)>>32 - 16
		c += d
		b += (c ^ d ^ a) + ctx.GetBlock()[14] + 0xfde5380c
		b = b<<23 | (b&0xffffffff)>>32 - 23
		b += c
		a += (b ^ c ^ d) + ctx.GetBlock()[1] + 0xa4beea44
		a = a<<4 | (a&0xffffffff)>>32 - 4
		a += b
		d += (a ^ b ^ c) + ctx.GetBlock()[4] + 0x4bdecfa9
		d = d<<11 | (d&0xffffffff)>>32 - 11
		d += a
		c += (d ^ a ^ b) + ctx.GetBlock()[7] + 0xf6bb4b60
		c = c<<16 | (c&0xffffffff)>>32 - 16
		c += d
		b += (c ^ d ^ a) + ctx.GetBlock()[10] + 0xbebfbc70
		b = b<<23 | (b&0xffffffff)>>32 - 23
		b += c
		a += (b ^ c ^ d) + ctx.GetBlock()[13] + 0x289b7ec6
		a = a<<4 | (a&0xffffffff)>>32 - 4
		a += b
		d += (a ^ b ^ c) + ctx.GetBlock()[0] + 0xeaa127fa
		d = d<<11 | (d&0xffffffff)>>32 - 11
		d += a
		c += (d ^ a ^ b) + ctx.GetBlock()[3] + 0xd4ef3085
		c = c<<16 | (c&0xffffffff)>>32 - 16
		c += d
		b += (c ^ d ^ a) + ctx.GetBlock()[6] + 0x4881d05
		b = b<<23 | (b&0xffffffff)>>32 - 23
		b += c
		a += (b ^ c ^ d) + ctx.GetBlock()[9] + 0xd9d4d039
		a = a<<4 | (a&0xffffffff)>>32 - 4
		a += b
		d += (a ^ b ^ c) + ctx.GetBlock()[12] + 0xe6db99e5
		d = d<<11 | (d&0xffffffff)>>32 - 11
		d += a
		c += (d ^ a ^ b) + ctx.GetBlock()[15] + 0x1fa27cf8
		c = c<<16 | (c&0xffffffff)>>32 - 16
		c += d
		b += (c ^ d ^ a) + ctx.GetBlock()[2] + 0xc4ac5665
		b = b<<23 | (b&0xffffffff)>>32 - 23
		b += c

		/* Round 4 */

		a += (c ^ (b | ^d)) + ctx.GetBlock()[0] + 0xf4292244
		a = a<<6 | (a&0xffffffff)>>32 - 6
		a += b
		d += (b ^ (a | ^c)) + ctx.GetBlock()[7] + 0x432aff97
		d = d<<10 | (d&0xffffffff)>>32 - 10
		d += a
		c += (a ^ (d | ^b)) + ctx.GetBlock()[14] + 0xab9423a7
		c = c<<15 | (c&0xffffffff)>>32 - 15
		c += d
		b += (d ^ (c | ^a)) + ctx.GetBlock()[5] + 0xfc93a039
		b = b<<21 | (b&0xffffffff)>>32 - 21
		b += c
		a += (c ^ (b | ^d)) + ctx.GetBlock()[12] + 0x655b59c3
		a = a<<6 | (a&0xffffffff)>>32 - 6
		a += b
		d += (b ^ (a | ^c)) + ctx.GetBlock()[3] + 0x8f0ccc92
		d = d<<10 | (d&0xffffffff)>>32 - 10
		d += a
		c += (a ^ (d | ^b)) + ctx.GetBlock()[10] + 0xffeff47d
		c = c<<15 | (c&0xffffffff)>>32 - 15
		c += d
		b += (d ^ (c | ^a)) + ctx.GetBlock()[1] + 0x85845dd1
		b = b<<21 | (b&0xffffffff)>>32 - 21
		b += c
		a += (c ^ (b | ^d)) + ctx.GetBlock()[8] + 0x6fa87e4f
		a = a<<6 | (a&0xffffffff)>>32 - 6
		a += b
		d += (b ^ (a | ^c)) + ctx.GetBlock()[15] + 0xfe2ce6e0
		d = d<<10 | (d&0xffffffff)>>32 - 10
		d += a
		c += (a ^ (d | ^b)) + ctx.GetBlock()[6] + 0xa3014314
		c = c<<15 | (c&0xffffffff)>>32 - 15
		c += d
		b += (d ^ (c | ^a)) + ctx.GetBlock()[13] + 0x4e0811a1
		b = b<<21 | (b&0xffffffff)>>32 - 21
		b += c
		a += (c ^ (b | ^d)) + ctx.GetBlock()[4] + 0xf7537e82
		a = a<<6 | (a&0xffffffff)>>32 - 6
		a += b
		d += (b ^ (a | ^c)) + ctx.GetBlock()[11] + 0xbd3af235
		d = d<<10 | (d&0xffffffff)>>32 - 10
		d += a
		c += (a ^ (d | ^b)) + ctx.GetBlock()[2] + 0x2ad7d2bb
		c = c<<15 | (c&0xffffffff)>>32 - 15
		c += d
		b += (d ^ (c | ^a)) + ctx.GetBlock()[9] + 0xeb86d391
		b = b<<21 | (b&0xffffffff)>>32 - 21
		b += c
		a += saved_a
		b += saved_b
		c += saved_c
		d += saved_d
		ptr += 64
		if !(g.AssignOp(&size, "-=", 64)) {
			break
		}
	}
	ctx.SetA(a)
	ctx.SetB(b)
	ctx.SetC(c)
	ctx.SetD(d)
	return ptr
}
func PHP_MD5Init(ctx *PHP_MD5_CTX) {
	ctx.SetA(0x67452301)
	ctx.SetB(0xefcdab89)
	ctx.SetC(0x98badcfe)
	ctx.SetD(0x10325476)
	ctx.SetLo(0)
	ctx.SetHi(0)
}
func PHP_MD5Update(ctx *PHP_MD5_CTX, data any, size int) {
	var saved_lo uint32
	var used uint32
	var free uint32
	saved_lo = ctx.GetLo()
	if g.Assign(&(ctx.GetLo()), saved_lo+size&0x1fffffff) < saved_lo {
		ctx.GetHi()++
	}
	ctx.SetHi(ctx.GetHi() + size>>29)
	used = saved_lo & 0x3f
	if used != 0 {
		free = 64 - used
		if size < free {
			memcpy(&ctx.buffer[used], data, size)
			return
		}
		memcpy(&ctx.buffer[used], data, free)
		data = (*uint8)(data + free)
		size -= free
		Body(ctx, ctx.GetBuffer(), 64)
	}
	if size >= 64 {
		data = Body(ctx, data, size & ^int(0x3f))
		size &= 0x3f
	}
	memcpy(ctx.GetBuffer(), data, size)
}
func PHP_MD5Final(result *uint8, ctx *PHP_MD5_CTX) {
	var used uint32
	var free uint32
	used = ctx.GetLo() & 0x3f
	ctx.GetBuffer()[g.PostInc(&used)] = 0x80
	free = 64 - used
	if free < 8 {
		memset(&ctx.buffer[used], 0, free)
		Body(ctx, ctx.GetBuffer(), 64)
		used = 0
		free = 64
	}
	memset(&ctx.buffer[used], 0, free-8)
	ctx.SetLo(ctx.GetLo() << 3)
	ctx.GetBuffer()[56] = ctx.GetLo()
	ctx.GetBuffer()[57] = ctx.GetLo() >> 8
	ctx.GetBuffer()[58] = ctx.GetLo() >> 16
	ctx.GetBuffer()[59] = ctx.GetLo() >> 24
	ctx.GetBuffer()[60] = ctx.GetHi()
	ctx.GetBuffer()[61] = ctx.GetHi() >> 8
	ctx.GetBuffer()[62] = ctx.GetHi() >> 16
	ctx.GetBuffer()[63] = ctx.GetHi() >> 24
	Body(ctx, ctx.GetBuffer(), 64)
	result[0] = ctx.GetA()
	result[1] = ctx.GetA() >> 8
	result[2] = ctx.GetA() >> 16
	result[3] = ctx.GetA() >> 24
	result[4] = ctx.GetB()
	result[5] = ctx.GetB() >> 8
	result[6] = ctx.GetB() >> 16
	result[7] = ctx.GetB() >> 24
	result[8] = ctx.GetC()
	result[9] = ctx.GetC() >> 8
	result[10] = ctx.GetC() >> 16
	result[11] = ctx.GetC() >> 24
	result[12] = ctx.GetD()
	result[13] = ctx.GetD() >> 8
	result[14] = ctx.GetD() >> 16
	result[15] = ctx.GetD() >> 24
	core.PhpExplicitBzero(ctx, g.SizeOf("* ctx"))
}
