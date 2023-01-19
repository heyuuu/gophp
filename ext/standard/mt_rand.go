// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/mt_rand.c>

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
   | Authors: Rasmus Lerdorf <rasmus@php.net>                             |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Pedro Melo <melo@ip.pt>                                     |
   |          Sterling Hughes <sterling@php.net>                          |
   |                                                                      |
   | Based on code from: Richard J. Wagner <rjwagner@writeme.com>         |
   |                     Makoto Matsumoto <matumoto@math.keio.ac.jp>      |
   |                     Takuji Nishimura                                 |
   |                     Shawn Cokus <Cokus@math.washington.edu>          |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_rand.h"

// # include "php_mt_rand.h"

/* MT RAND FUNCTIONS */

// #define N       MT_N

// #define M       ( 397 )

// #define hiBit(u) ( ( u ) & 0x80000000U )

// #define loBit(u) ( ( u ) & 0x00000001U )

// #define loBits(u) ( ( u ) & 0x7FFFFFFFU )

// #define mixBits(u,v) ( hiBit ( u ) | loBits ( v ) )

// #define twist(m,u,v) ( m ^ ( mixBits ( u , v ) >> 1 ) ^ ( ( uint32_t ) ( - ( int32_t ) ( loBit ( v ) ) ) & 0x9908b0dfU ) )

// #define twist_php(m,u,v) ( m ^ ( mixBits ( u , v ) >> 1 ) ^ ( ( uint32_t ) ( - ( int32_t ) ( loBit ( u ) ) ) & 0x9908b0dfU ) )

/* {{{ php_mt_initialize
 */

func PhpMtInitialize(seed uint32, state *uint32) {
	/* Initialize generator state with seed
	   See Knuth TAOCP Vol 2, 3rd Ed, p.106 for multiplier.
	   In previous versions, most significant bits (MSBs) of the seed affect
	   only MSBs of the state array.  Modified 9 Jan 2002 by Makoto Matsumoto. */

	var s *uint32 = state
	var r *uint32 = state
	var i int = 1
	g.PostInc(&(*s)) = seed & 0xffffffff
	for ; i < 624; i++ {
		g.PostInc(&(*s)) = 1812433253*((*r)^(*r)>>30) + i&0xffffffff
		r++
	}
}

/* }}} */

func PhpMtReload() {
	/* Generate N new values in state
	   Made clearer and faster by Matthew Bellew (matthew.bellew@home.com) */

	var state *uint32 = BasicGlobals.GetState()
	var p *uint32 = state
	var i int
	if BasicGlobals.GetMtRandMode() == 0 {
		for i = 624 - 397; g.PostDec(&i); p++ {
			*p = p[397] ^ (p[0]&0x80000000|p[1]&0x7fffffff)>>1 ^ uint32(-(int32(p[1]&0x1)))&0x9908b0df
		}
		for i = 397; g.PreDec(&i); p++ {
			*p = p[397-624] ^ (p[0]&0x80000000|p[1]&0x7fffffff)>>1 ^ uint32(-(int32(p[1]&0x1)))&0x9908b0df
		}
		*p = p[397-624] ^ (p[0]&0x80000000|state[0]&0x7fffffff)>>1 ^ uint32(-(int32(state[0]&0x1)))&0x9908b0df
	} else {
		for i = 624 - 397; g.PostDec(&i); p++ {
			*p = p[397] ^ (p[0]&0x80000000|p[1]&0x7fffffff)>>1 ^ uint32(-(int32(p[0]&0x1)))&0x9908b0df
		}
		for i = 397; g.PreDec(&i); p++ {
			*p = p[397-624] ^ (p[0]&0x80000000|p[1]&0x7fffffff)>>1 ^ uint32(-(int32(p[0]&0x1)))&0x9908b0df
		}
		*p = p[397-624] ^ (p[0]&0x80000000|state[0]&0x7fffffff)>>1 ^ uint32(-(int32(p[0]&0x1)))&0x9908b0df
	}
	BasicGlobals.SetLeft(624)
	BasicGlobals.SetNext(state)
}

/* }}} */

func PhpMtSrand(seed uint32) {
	/* Seed the generator with a simple uint32 */

	PhpMtInitialize(seed, BasicGlobals.GetState())
	PhpMtReload()

	/* Seed only once */

	BasicGlobals.SetMtRandIsSeeded(1)

	/* Seed only once */
}

/* }}} */

func PhpMtRand() uint32 {
	/* Pull a 32-bit integer from the generator state
	   Every other access function simply transforms the numbers extracted here */

	var s1 uint32
	if BasicGlobals.GetMtRandIsSeeded() == 0 {
		PhpMtSrand(zend_long(time(0)*getpid()) ^ zend_long(1000000.0*PhpCombinedLcg()))
	}
	if BasicGlobals.GetLeft() == 0 {
		PhpMtReload()
	}
	BasicGlobals.GetLeft()--
	*(BasicGlobals.GetNext())++
	s1 = (*(BasicGlobals.GetNext())) - 1
	s1 ^= s1 >> 11
	s1 ^= s1 << 7 & 0x9d2c5680
	s1 ^= s1 << 15 & 0xefc60000
	return s1 ^ s1>>18
}

/* }}} */

func ZifMtSrand(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var seed zend.ZendLong = 0
	var mode zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
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

			if zend.ZendParseArgLong(_arg, &seed, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgLong(_arg, &mode, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if execute_data.This.u2.num_args == 0 {
		seed = zend_long(time(0)*getpid()) ^ zend_long(1000000.0*PhpCombinedLcg())
	}
	switch mode {
	case 1:
		BasicGlobals.SetMtRandMode(1)
		break
	default:
		BasicGlobals.SetMtRandMode(0)
	}
	PhpMtSrand(seed)
}

/* }}} */

func RandRange32(umax uint32) uint32 {
	var result uint32
	var limit uint32
	result = PhpMtRand()

	/* Special case where no modulus is required */

	if umax == UINT32_MAX {
		return result
	}

	/* Increment the max so the range is inclusive of max */

	umax++

	/* Powers of two are not biased */

	if (umax&umax - 1) == 0 {
		return result&umax - 1
	}

	/* Ceiling under which UINT32_MAX % max == 0 */

	limit = UINT32_MAX - UINT32_MAX%umax - 1

	/* Discard numbers over the limit to avoid modulo bias */

	for result > limit {
		result = PhpMtRand()
	}
	return result % umax
}

/* {{{ php_mt_rand_range
 */

func PhpMtRandRange(min zend.ZendLong, max zend.ZendLong) zend.ZendLong {
	var umax zend.ZendUlong = max - min
	return zend_long(RandRange32(umax) + min)
}

/* }}} */

func PhpMtRandCommon(min zend.ZendLong, max zend.ZendLong) zend.ZendLong {
	var n int64
	if BasicGlobals.GetMtRandMode() == 0 {
		return PhpMtRandRange(min, max)
	}

	/* Legacy mode deliberately not inside php_mt_rand_range()
	 * to prevent other functions being affected */

	n = int64(PhpMtRand() >> 1)
	n = min + zend_long(float64(float64(max-min+1.0)*(n/(zend_long(0x7fffffff)+1.0))))
	return n
}

/* }}} */

func ZifMtRand(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var min zend.ZendLong
	var max zend.ZendLong
	var argc int = execute_data.This.u2.num_args
	if argc == 0 {

		// genrand_int31 in mt19937ar.c performs a right shift

		var __z *zend.Zval = return_value
		__z.value.lval = PhpMtRand() >> 1
		__z.u1.type_info = 4
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 2
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

			if zend.ZendParseArgLong(_arg, &min, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgLong(_arg, &max, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if max < min {
		core.PhpErrorDocref(nil, 1<<1, "max("+"%"+"lld"+") is smaller than min("+"%"+"lld"+")", max, min)
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = PhpMtRandCommon(min, max)
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifMtGetrandmax(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}

	/*
	 * Melo: it could be 2^^32 but we only use 2^^31 to maintain
	 * compatibility with the previous php_rand
	 */

	var __z *zend.Zval = return_value
	__z.value.lval = zend_long(0x7fffffff)
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZmStartupMtRand(type_ int, module_number int) int {
	zend.ZendRegisterLongConstant("MT_RAND_MT19937", g.SizeOf("\"MT_RAND_MT19937\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("MT_RAND_PHP", g.SizeOf("\"MT_RAND_PHP\"")-1, 1, 1<<0|1<<1, module_number)
	return zend.SUCCESS
}
