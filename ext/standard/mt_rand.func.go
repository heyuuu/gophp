// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func hiBit(u uint32) int             { return u & 0x80000000 }
func loBit(u uint32) int             { return u & 0x1 }
func loBits(u uint32) int            { return u & 0x7fffffff }
func mixBits(u uint32, v uint32) int { return hiBit(u) | loBits(v) }
func Twist(m uint32, u uint32, v uint32) int {
	return m ^ mixBits(u, v)>>1 ^ uint32(-(int32(loBit(v))))&0x9908b0df
}
func TwistPhp(m uint32, u uint32, v uint32) int {
	return m ^ mixBits(u, v)>>1 ^ uint32(-(int32(loBit(u))))&0x9908b0df
}
func PhpMtInitialize(seed uint32, state *uint32) {
	/* Initialize generator state with seed
	   See Knuth TAOCP Vol 2, 3rd Ed, p.106 for multiplier.
	   In previous versions, most significant bits (MSBs) of the seed affect
	   only MSBs of the state array.  Modified 9 Jan 2002 by Makoto Matsumoto. */

	var s *uint32 = state
	var r *uint32 = state
	var i int = 1
	b.PostInc(&(*s)) = seed & 0xffffffff
	for ; i < N; i++ {
		b.PostInc(&(*s)) = 1812433253*((*r)^(*r)>>30) + i&0xffffffff
		r++
	}
}
func PhpMtReload() {
	/* Generate N new values in state
	   Made clearer and faster by Matthew Bellew (matthew.bellew@home.com) */

	var state *uint32 = BG(state)
	var p *uint32 = state
	var i int
	if BG(mt_rand_mode) == MT_RAND_MT19937 {
		for i = N - M; b.PostDec(&i); p++ {
			*p = Twist(p[M], p[0], p[1])
		}
		for i = M; b.PreDec(&i); p++ {
			*p = Twist(p[M-N], p[0], p[1])
		}
		*p = Twist(p[M-N], p[0], state[0])
	} else {
		for i = N - M; b.PostDec(&i); p++ {
			*p = TwistPhp(p[M], p[0], p[1])
		}
		for i = M; b.PreDec(&i); p++ {
			*p = TwistPhp(p[M-N], p[0], p[1])
		}
		*p = TwistPhp(p[M-N], p[0], state[0])
	}
	BG(left) = N
	BG(next) = state
}
func PhpMtSrand(seed uint32) {
	/* Seed the generator with a simple uint32 */

	PhpMtInitialize(seed, BG(state))
	PhpMtReload()

	/* Seed only once */

	BG(mt_rand_is_seeded) = 1

	/* Seed only once */
}
func PhpMtRand() uint32 {
	/* Pull a 32-bit integer from the generator state
	   Every other access function simply transforms the numbers extracted here */

	var s1 uint32
	if !(BG(mt_rand_is_seeded)) {
		PhpMtSrand(GENERATE_SEED())
	}
	if BG(left) == 0 {
		PhpMtReload()
	}
	BG(left)--
	(*BG)(next)++
	s1 = (*BG)(next) - 1
	s1 ^= s1 >> 11
	s1 ^= s1 << 7 & 0x9d2c5680
	s1 ^= s1 << 15 & 0xefc60000
	return s1 ^ s1>>18
}
func ZifMtSrand(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var seed zend.ZendLong = 0
	var mode zend.ZendLong = MT_RAND_MT19937
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
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
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = executeData.Arg(0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &seed, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &mode, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
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
	if executeData.NumArgs() == 0 {
		seed = GENERATE_SEED()
	}
	switch mode {
	case MT_RAND_PHP:
		BG(mt_rand_mode) = MT_RAND_PHP
	default:
		BG(mt_rand_mode) = MT_RAND_MT19937
	}
	PhpMtSrand(seed)
}
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
func PhpMtRandRange(min zend.ZendLong, max zend.ZendLong) zend.ZendLong {
	var umax zend.ZendUlong = max - min
	return zend_long(RandRange32(umax) + min)
}
func PhpMtRandCommon(min zend.ZendLong, max zend.ZendLong) zend.ZendLong {
	var n int64
	if BG(mt_rand_mode) == MT_RAND_MT19937 {
		return PhpMtRandRange(min, max)
	}

	/* Legacy mode deliberately not inside php_mt_rand_range()
	 * to prevent other functions being affected */

	n = int64(PhpMtRand() >> 1)
	RAND_RANGE_BADSCALING(n, min, max, PHP_MT_RAND_MAX)
	return n
}
func ZifMtRand(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var min zend.ZendLong
	var max zend.ZendLong
	var argc int = executeData.NumArgs()
	if argc == 0 {

		// genrand_int31 in mt19937ar.c performs a right shift

		return_value.SetLong(PhpMtRand() >> 1)
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
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
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = executeData.Arg(0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &min, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &max, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
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
	if max < min {
		core.PhpErrorDocref(nil, zend.E_WARNING, "max("+zend.ZEND_LONG_FMT+") is smaller than min("+zend.ZEND_LONG_FMT+")", max, min)
		return_value.SetFalse()
		return
	}
	return_value.SetLong(PhpMtRandCommon(min, max))
	return
}
func ZifMtGetrandmax(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}

	/*
	 * Melo: it could be 2^^32 but we only use 2^^31 to maintain
	 * compatibility with the previous php_rand
	 */

	return_value.SetLong(PHP_MT_RAND_MAX)
	return
}
func ZmStartupMtRand(type_ int, module_number int) int {
	zend.REGISTER_LONG_CONSTANT("MT_RAND_MT19937", MT_RAND_MT19937, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("MT_RAND_PHP", MT_RAND_PHP, zend.CONST_CS|zend.CONST_PERSISTENT)
	return zend.SUCCESS
}
