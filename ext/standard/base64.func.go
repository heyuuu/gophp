// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
)

func PhpBase64EncodeStr(str *zend.ZendString) *zend.ZendString {
	return PhpBase64Encode((*uint8)(str.GetVal()), str.GetLen())
}
func PhpBase64Decode(str *uint8, len_ int) *zend.ZendString { return PhpBase64DecodeEx(str, len_, 0) }
func PhpBase64DecodeStr(str *zend.ZendString) *zend.ZendString {
	return PhpBase64DecodeEx((*uint8)(str.GetVal()), str.GetLen(), 0)
}
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
		case 1:
			out[b.PostInc(&j)] |= ch >> 4
			out[j] = (ch & 0xf) << 4
		case 2:
			out[b.PostInc(&j)] |= ch >> 2
			out[j] = (ch & 0x3) << 6
		case 3:
			out[b.PostInc(&j)] |= ch
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
func PhpBase64Encode(str *uint8, length int) *zend.ZendString {
	var p *uint8
	var result *zend.ZendString
	result = zend.ZendStringSafeAlloc((length+2)/3, 4*b.SizeOf("char"), 0, 0)
	p = (*uint8)(result.GetVal())
	p = PhpBase64EncodeImpl(str, length, p)
	result.SetLen(p - (*uint8)(result.GetVal()))
	return result
}
func PhpBase64DecodeEx(str *uint8, length int, strict zend.ZendBool) *zend.ZendString {
	var result *zend.ZendString
	var outl int = 0
	result = zend.ZendStringAlloc(length, 0)
	if PhpBase64DecodeImpl(str, length, (*uint8)(result.GetVal()), &outl, strict) == 0 {
		zend.ZendStringEfree(result)
		return nil
	}
	result.SetLen(outl)
	return result
}
func ZifBase64Encode(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
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
			if zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	result = PhpBase64Encode((*uint8)(str), str_len)
	return_value.SetString(result)
	return
}
func ZifBase64Decode(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
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
			if zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &strict, &_dummy, 0) == 0 {
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
	result = PhpBase64DecodeEx((*uint8)(str), str_len, strict)
	if result != nil {
		return_value.SetString(result)
		return
	} else {
		return_value.SetFalse()
		return
	}
}
