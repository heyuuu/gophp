// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	r "sik/runtime"
	"sik/zend"
)

func PHP_UU_ENC(c __auto__) __auto__ {
	if c {
		return (c & 077) + ' '
	} else {
		return '`'
	}
}
func PHP_UU_ENC_C2(c int) __auto__ {
	return PHP_UU_ENC((*c)<<4&060 | (*(c + 1))>>4&017)
}
func PHP_UU_ENC_C3(c int) __auto__ {
	return PHP_UU_ENC((*(c + 1))<<2&074 | (*(c + 2))>>6&3)
}
func PHP_UU_DEC(c char) int { return c - ' '&077 }
func PhpUuencode(src *byte, src_len int) *zend.ZendString {
	var len_ int = 45
	var p *uint8
	var s *uint8
	var e *uint8
	var ee *uint8
	var dest *zend.ZendString

	/* encoded length is ~ 38% greater than the original
	   Use 1.5 for easier calculation.
	*/

	dest = zend.ZendStringSafeAlloc(src_len/2, 3, 46, 0)
	p = (*uint8)(zend.ZSTR_VAL(dest))
	s = (*uint8)(src)
	e = s + src_len
	for s+3 < e {
		ee = s + len_
		if ee > e {
			ee = e
			len_ = ee - s
			if len_%3 != 0 {
				ee = s + int(floor(float64(len_/3))*3)
			}
		}
		b.PostInc(&(*p)) = PHP_UU_ENC(len_)
		for s < ee {
			b.PostInc(&(*p)) = PHP_UU_ENC((*s) >> 2)
			b.PostInc(&(*p)) = PHP_UU_ENC_C2(s)
			b.PostInc(&(*p)) = PHP_UU_ENC_C3(s)
			b.PostInc(&(*p)) = PHP_UU_ENC((*(s + 2)) & 077)
			s += 3
		}
		if len_ == 45 {
			b.PostInc(&(*p)) = '\n'
		}
	}
	if s < e {
		if len_ == 45 {
			b.PostInc(&(*p)) = PHP_UU_ENC(e - s)
			len_ = 0
		}
		b.PostInc(&(*p)) = PHP_UU_ENC((*s) >> 2)
		b.PostInc(&(*p)) = PHP_UU_ENC_C2(s)
		if e-s > 1 {
			b.PostInc(&(*p)) = PHP_UU_ENC_C3(s)
		} else {
			b.PostInc(&(*p)) = PHP_UU_ENC('0')
		}
		if e-s > 2 {
			b.PostInc(&(*p)) = PHP_UU_ENC((*(s + 2)) & 077)
		} else {
			b.PostInc(&(*p)) = PHP_UU_ENC('0')
		}
	}
	if len_ < 45 {
		b.PostInc(&(*p)) = '\n'
	}
	b.PostInc(&(*p)) = PHP_UU_ENC('0')
	b.PostInc(&(*p)) = '\n'
	*p = '0'
	dest = zend.ZendStringTruncate(dest, (*byte)(p-zend.ZSTR_VAL(dest)), 0)
	return dest
}
func PhpUudecode(src *byte, src_len int) *zend.ZendString {
	var len_ int
	var total_len int = 0
	var s *byte
	var e *byte
	var p *byte
	var ee *byte
	var dest *zend.ZendString
	dest = zend.ZendStringAlloc(int(ceil(src_len*0.75)), 0)
	p = zend.ZSTR_VAL(dest)
	s = src
	e = src + src_len
	for s < e {
		if b.Assign(&len_, PHP_UU_DEC(b.PostInc(&(*s)))) == 0 {
			break
		}

		/* sanity check */

		if len_ > src_len {
			goto err
		}
		total_len += len_
		ee = s + b.CondF2(len_ == 45, 60, func() int { return int(floor(len_ * 1.33)) })

		/* sanity check */

		if ee > e {
			goto err
		}
		for s < ee {
			if s+4 > e {
				goto err
			}
			b.PostInc(&(*p)) = PHP_UU_DEC(*s)<<2 | PHP_UU_DEC(*(s + 1))>>4
			b.PostInc(&(*p)) = PHP_UU_DEC(*(s + 1))<<4 | PHP_UU_DEC(*(s + 2))>>2
			b.PostInc(&(*p)) = PHP_UU_DEC(*(s + 2))<<6 | PHP_UU_DEC(*(s + 3))
			s += 4
		}
		if len_ < 45 {
			break
		}

		/* skip \n */

		s++

		/* skip \n */

	}
	r.Assert(p >= zend.ZSTR_VAL(dest))
	if b.Assign(&len_, total_len) > size_t(p-zend.ZSTR_VAL(dest)) {
		b.PostInc(&(*p)) = PHP_UU_DEC(*s)<<2 | PHP_UU_DEC(*(s + 1))>>4
		if len_ > 1 {
			b.PostInc(&(*p)) = PHP_UU_DEC(*(s + 1))<<4 | PHP_UU_DEC(*(s + 2))>>2
			if len_ > 2 {
				b.PostInc(&(*p)) = PHP_UU_DEC(*(s + 2))<<6 | PHP_UU_DEC(*(s + 3))
			}
		}
	}
	zend.ZSTR_LEN(dest) = total_len
	zend.ZSTR_VAL(dest)[zend.ZSTR_LEN(dest)] = '0'
	return dest
err:
	zend.ZendStringEfree(dest)
	return nil
}
func ZifConvertUuencode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var src *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &src, 0) == 0) {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if zend.ZSTR_LEN(src) < 1 {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STR(PhpUuencode(zend.ZSTR_VAL(src), zend.ZSTR_LEN(src)))
	return
}
func ZifConvertUudecode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var src *zend.ZendString
	var dest *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &src, 0) == 0) {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if zend.ZSTR_LEN(src) < 1 {
		zend.RETVAL_FALSE
		return
	}
	if b.Assign(&dest, PhpUudecode(zend.ZSTR_VAL(src), zend.ZSTR_LEN(src))) == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The given parameter is not a valid uuencoded string")
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STR(dest)
	return
}
