// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func MakeSha1Digest(sha1str *byte, digest *uint8) { MakeDigestEx(sha1str, digest, 20) }
func ZifSha1(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *zend.ZendString
	var raw_output zend.ZendBool = 0
	var context PHP_SHA1_CTX
	var digest []uint8
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
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
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
			if zend.ZendParseArgStr(_arg, &arg, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &raw_output, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
	PHP_SHA1Init(&context)
	PHP_SHA1Update(&context, (*uint8)(arg.GetVal()), arg.GetLen())
	PHP_SHA1Final(digest, &context)
	if raw_output != 0 {
		zend.ZVAL_STRINGL(return_value, (*byte)(digest), 20)
		return
	} else {
		return_value.SetString(zend.ZendStringAlloc(40, 0))
		MakeDigestEx(zend.Z_STRVAL_P(return_value), digest, 20)
	}
}
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
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
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
			if zend.ZendParseArgPath(_arg, &arg, &arg_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &raw_output, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
	stream = core.PhpStreamOpenWrapper(arg, "rb", core.REPORT_ERRORS, nil)
	if stream == nil {
		return_value.SetFalse()
		return
	}
	PHP_SHA1Init(&context)
	for b.Assign(&n, core.PhpStreamRead(stream, (*byte)(buf), b.SizeOf("buf"))) > 0 {
		PHP_SHA1Update(&context, buf, n)
	}
	PHP_SHA1Final(digest, &context)
	core.PhpStreamClose(stream)
	if raw_output != 0 {
		zend.ZVAL_STRINGL(return_value, (*byte)(digest), 20)
		return
	} else {
		return_value.SetString(zend.ZendStringAlloc(40, 0))
		MakeDigestEx(zend.Z_STRVAL_P(return_value), digest, 20)
	}
}
func ROTATE_LEFT(x uint32, n int) int { return x<<n | x>>32 - n }
func W(i int) __auto__ {
	tmp = x[i-3&15] ^ x[i-8&15] ^ x[i-14&15] ^ x[i&15]
	x[i&15] = ROTATE_LEFT(tmp, 1)
	return x[i&15]
}
func FF(
	a uint32,
	b uint32,
	c uint32,
	d uint32,
	e __auto__,
	w int,
) {
	e += F(b, c, d) + w + uint32(0x5a827999)
	e += ROTATE_LEFT(a, 5)
	b = ROTATE_LEFT(b, 30)
}
func GG(
	a uint32,
	b uint32,
	c uint32,
	d uint32,
	e __auto__,
	w int,
) {
	e += G(b, c, d) + w + uint32(0x6ed9eba1)
	e += ROTATE_LEFT(a, 5)
	b = ROTATE_LEFT(b, 30)
}
func HH(
	a uint32,
	b uint32,
	c uint32,
	d uint32,
	e __auto__,
	w int,
) {
	e += H(b, c, d) + w + uint32(0x8f1bbcdc)
	e += ROTATE_LEFT(a, 5)
	b = ROTATE_LEFT(b, 30)
}
func II(
	a uint32,
	b uint32,
	c uint32,
	d uint32,
	e __auto__,
	w int,
) {
	e += I(b, c, d) + w + uint32(0xca62c1d6)
	e += ROTATE_LEFT(a, 5)
	b = ROTATE_LEFT(b, 30)
}
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
func PHP_SHA1Update(context *PHP_SHA1_CTX, input *uint8, inputLen int) {
	var i uint
	var index uint
	var partLen uint

	/* Compute number of bytes mod 64 */

	index = uint(context.GetCount()[0] >> 3 & 0x3f)

	/* Update number of bits */

	if b.AssignOp(&context.GetCount()[0], "+=", uint32(inputLen<<3)) < uint32(inputLen<<3) {
		context.GetCount()[1]++
	}
	context.GetCount()[1] += uint32(inputLen >> 29)
	partLen = 64 - index

	/* Transform as many times as possible.
	 */

	if inputLen >= partLen {
		memcpy((*uint8)(context.GetBuffer()[index]), (*uint8)(input), partLen)
		SHA1Transform(context.GetState(), context.GetBuffer())
		for i = partLen; i+63 < inputLen; i += 64 {
			SHA1Transform(context.GetState(), &input[i])
		}
		index = 0
	} else {
		i = 0
	}

	/* Buffer remaining input */

	memcpy((*uint8)(context.GetBuffer()[index]), (*uint8)(&input[i]), inputLen-i)

	/* Buffer remaining input */
}
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

	zend.ZEND_SECURE_ZERO((*uint8)(context), b.SizeOf("* context"))

	/* Zeroize sensitive information.
	 */
}
