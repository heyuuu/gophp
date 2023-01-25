// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/sapi/cli"
	"sik/zend"
)

func RegisterStringConstants(type_ int, module_number int) {
	zend.REGISTER_LONG_CONSTANT("STR_PAD_LEFT", STR_PAD_LEFT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STR_PAD_RIGHT", STR_PAD_RIGHT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STR_PAD_BOTH", STR_PAD_BOTH, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PATHINFO_DIRNAME", PHP_PATHINFO_DIRNAME, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PATHINFO_BASENAME", PHP_PATHINFO_BASENAME, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PATHINFO_EXTENSION", PHP_PATHINFO_EXTENSION, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PATHINFO_FILENAME", PHP_PATHINFO_FILENAME, zend.CONST_CS|zend.CONST_PERSISTENT)

	/* If last members of struct lconv equal CHAR_MAX, no grouping is done */

	zend.REGISTER_LONG_CONSTANT("CHAR_MAX", CHAR_MAX, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("LC_CTYPE", LC_CTYPE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("LC_NUMERIC", LC_NUMERIC, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("LC_TIME", LC_TIME, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("LC_COLLATE", LC_COLLATE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("LC_MONETARY", LC_MONETARY, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("LC_ALL", LC_ALL, zend.CONST_CS|zend.CONST_PERSISTENT)
}
func PhpBin2hex(old *uint8, oldlen int) *zend.ZendString {
	var result *zend.ZendString
	var i int
	var j int
	result = zend.ZendStringSafeAlloc(oldlen, 2*b.SizeOf("char"), 0, 0)
	j = 0
	i = j
	for ; i < oldlen; i++ {
		zend.ZSTR_VAL(result)[b.PostInc(&j)] = Hexconvtab[old[i]>>4]
		zend.ZSTR_VAL(result)[b.PostInc(&j)] = Hexconvtab[old[i]&15]
	}
	zend.ZSTR_VAL(result)[j] = '0'
	return result
}
func PhpHex2bin(old *uint8, oldlen int) *zend.ZendString {
	var target_length int = oldlen >> 1
	var str *zend.ZendString = zend.ZendStringAlloc(target_length, 0)
	var ret *uint8 = (*uint8)(zend.ZSTR_VAL(str))
	var i int
	var j int
	j = 0
	i = j
	for ; i < target_length; i++ {
		var c uint8 = old[b.PostInc(&j)]
		var l uint8 = c & ^0x20
		var is_letter int = uint(l-'A'^l-'F'-1)>>8*b.SizeOf("unsigned int") - 1
		var d uint8

		/* basically (c >= '0' && c <= '9') || (l >= 'A' && l <= 'F') */

		if zend.EXPECTED(((c ^ '0') - 10>>8*b.SizeOf("unsigned int") - 1 | is_letter) != 0) {
			d = l - 0x10 - 0x27*is_letter<<4
		} else {
			zend.ZendStringEfree(str)
			return nil
		}
		c = old[b.PostInc(&j)]
		l = c & ^0x20
		is_letter = uint(l-'A'^l-'F'-1)>>8*b.SizeOf("unsigned int") - 1
		if zend.EXPECTED(((c ^ '0') - 10>>8*b.SizeOf("unsigned int") - 1 | is_letter) != 0) {
			d |= l - 0x10 - 0x27*is_letter
		} else {
			zend.ZendStringEfree(str)
			return nil
		}
		ret[i] = d
	}
	ret[i] = '0'
	return str
}
func LocaleconvR(out *__struct__lconv) *__struct__lconv {
	/*  cur->locinfo is struct __crt_locale_info which implementation is
	    hidden in vc14. TODO revisit this and check if a workaround available
	    and needed. */

	/* localeconv doesn't return an error condition */

	*out = (*localeconv)()
	return out
}
func ZifBin2hex(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var result *zend.ZendString
	var data *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &data, 0) == 0) {
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
	result = PhpBin2hex((*uint8)(zend.ZSTR_VAL(data)), zend.ZSTR_LEN(data))
	if result == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STR(result)
	return
}
func ZifHex2bin(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var result *zend.ZendString
	var data *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &data, 0) == 0) {
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
	if zend.ZSTR_LEN(data)%2 != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Hexadecimal input string must have an even length")
		zend.RETVAL_FALSE
		return
	}
	result = PhpHex2bin((*uint8)(zend.ZSTR_VAL(data)), zend.ZSTR_LEN(data))
	if result == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Input string must be hexadecimal string")
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STR(result)
}
func PhpSpnCommonHandler(execute_data *zend.ZendExecuteData, return_value *zend.Zval, behavior int) {
	var s11 *zend.ZendString
	var s22 *zend.ZendString
	var start zend.ZendLong = 0
	var len_ zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 4
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &s11, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &s22, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &start, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &len_, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if zend.ZEND_NUM_ARGS() < 4 {
		len_ = zend.ZSTR_LEN(s11)
	}

	/* look at substr() function for more information */

	if start < 0 {
		start += zend.ZendLong(zend.ZSTR_LEN(s11))
		if start < 0 {
			start = 0
		}
	} else if int(start > zend.ZSTR_LEN(s11)) != 0 {
		zend.RETVAL_FALSE
		return
	}
	if len_ < 0 {
		len_ += zend.ZSTR_LEN(s11) - start
		if len_ < 0 {
			len_ = 0
		}
	}
	if len_ > zend.ZendLong(zend.ZSTR_LEN(s11)-start) {
		len_ = zend.ZSTR_LEN(s11) - start
	}
	if len_ == 0 {
		zend.RETVAL_LONG(0)
		return
	}
	if behavior == STR_STRSPN {
		zend.RETVAL_LONG(PhpStrspn(zend.ZSTR_VAL(s11)+start, zend.ZSTR_VAL(s22), zend.ZSTR_VAL(s11)+start+len_, zend.ZSTR_VAL(s22)+zend.ZSTR_LEN(s22)))
		return
	} else if behavior == STR_STRCSPN {
		zend.RETVAL_LONG(PhpStrcspn(zend.ZSTR_VAL(s11)+start, zend.ZSTR_VAL(s22), zend.ZSTR_VAL(s11)+start+len_, zend.ZSTR_VAL(s22)+zend.ZSTR_LEN(s22)))
		return
	}
}
func ZifStrspn(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpSpnCommonHandler(execute_data, return_value, STR_STRSPN)
}
func ZifStrcspn(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpSpnCommonHandler(execute_data, return_value, STR_STRCSPN)
}
func ZifStrcoll(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var s1 *zend.ZendString
	var s2 *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &s1, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &s2, 0) == 0) {
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
	zend.RETVAL_LONG(strcoll((*byte)(zend.ZSTR_VAL(s1)), (*byte)(zend.ZSTR_VAL(s2))))
	return
}
func PhpCharmask(input *uint8, len_ int, mask *byte) int {
	var end *uint8
	var c uint8
	var result int = zend.SUCCESS
	memset(mask, 0, 256)
	for end = input + len_; input < end; input++ {
		c = *input
		if input+3 < end && input[1] == '.' && input[2] == '.' && input[3] >= c {
			memset(mask+c, 1, input[3]-c+1)
			input += 3
		} else if input+1 < end && input[0] == '.' && input[1] == '.' {

			/* Error, try to be as helpful as possible:
			   (a range ending/starting with '.' won't be captured here) */

			if end-len_ >= input {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid '..'-range, no character to the left of '..'")
				result = zend.FAILURE
				continue
			}
			if input+2 >= end {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid '..'-range, no character to the right of '..'")
				result = zend.FAILURE
				continue
			}
			if input[-1] > input[2] {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid '..'-range, '..'-range needs to be incrementing")
				result = zend.FAILURE
				continue
			}

			/* FIXME: better error (a..b..c is the only left possibility?) */

			core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid '..'-range")
			result = zend.FAILURE
			continue
		} else {
			mask[c] = 1
		}
	}
	return result
}
func PhpTrimInt(str *zend.ZendString, what *byte, what_len int, mode int) *zend.ZendString {
	var start *byte = zend.ZSTR_VAL(str)
	var end *byte = start + zend.ZSTR_LEN(str)
	var mask []byte
	if what != nil {
		if what_len == 1 {
			var p byte = *what
			if (mode & 1) != 0 {
				for start != end {
					if (*start) == p {
						start++
					} else {
						break
					}
				}
			}
			if (mode & 2) != 0 {
				for start != end {
					if (*(end - 1)) == p {
						end--
					} else {
						break
					}
				}
			}
		} else {
			PhpCharmask((*uint8)(what), what_len, mask)
			if (mode & 1) != 0 {
				for start != end {
					if mask[uint8(*start)] {
						start++
					} else {
						break
					}
				}
			}
			if (mode & 2) != 0 {
				for start != end {
					if mask[uint8(*(end - 1))] {
						end--
					} else {
						break
					}
				}
			}
		}
	} else {
		if (mode & 1) != 0 {
			for start != end {
				var c uint8 = uint8(*start)
				if c <= ' ' && (c == ' ' || c == '\n' || c == '\r' || c == '\t' || c == 'v' || c == '0') {
					start++
				} else {
					break
				}
			}
		}
		if (mode & 2) != 0 {
			for start != end {
				var c uint8 = uint8(*(end - 1))
				if c <= ' ' && (c == ' ' || c == '\n' || c == '\r' || c == '\t' || c == 'v' || c == '0') {
					end--
				} else {
					break
				}
			}
		}
	}
	if zend.ZSTR_LEN(str) == end-start {
		return zend.ZendStringCopy(str)
	} else if end-start == 0 {
		return zend.ZSTR_EMPTY_ALLOC()
	} else {
		return zend.ZendStringInit(start, end-start, 0)
	}
}
func PhpTrim(str *zend.ZendString, what *byte, what_len int, mode int) *zend.ZendString {
	return PhpTrimInt(str, what, what_len, mode)
}
func PhpDoTrim(execute_data *zend.ZendExecuteData, return_value *zend.Zval, mode int) {
	var str *zend.ZendString
	var what *zend.ZendString = nil
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &what, 0) == 0) {
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
	zend.ZVAL_STR(return_value, PhpTrimInt(str, b.CondF1(what != nil, func() []byte { return zend.ZSTR_VAL(what) }, nil), b.CondF1(what != nil, func() int { return zend.ZSTR_LEN(what) }, 0), mode))
}
func ZifTrim(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpDoTrim(execute_data, return_value, 3)
}
func ZifRtrim(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpDoTrim(execute_data, return_value, 2)
}
func ZifLtrim(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpDoTrim(execute_data, return_value, 1)
}
func ZifWordwrap(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var text *zend.ZendString
	var breakchar *byte = "\n"
	var newtextlen int
	var chk int
	var breakchar_len int = 1
	var alloced int
	var current zend.ZendLong = 0
	var laststart zend.ZendLong = 0
	var lastspace zend.ZendLong = 0
	var linelength zend.ZendLong = 75
	var docut zend.ZendBool = 0
	var newtext *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 4
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &text, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &linelength, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &breakchar, &breakchar_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &docut, &_dummy, 0) == 0) {
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
	if zend.ZSTR_LEN(text) == 0 {
		zend.RETVAL_EMPTY_STRING()
		return
	}
	if breakchar_len == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Break string cannot be empty")
		zend.RETVAL_FALSE
		return
	}
	if linelength == 0 && docut != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Can't force cut when width is zero")
		zend.RETVAL_FALSE
		return
	}

	/* Special case for a single-character break as it needs no
	   additional storage space */

	if breakchar_len == 1 && docut == 0 {
		newtext = zend.ZendStringInit(zend.ZSTR_VAL(text), zend.ZSTR_LEN(text), 0)
		lastspace = 0
		laststart = lastspace
		for current = 0; current < zend.ZendLong(zend.ZSTR_LEN(text)); current++ {
			if zend.ZSTR_VAL(text)[current] == breakchar[0] {
				lastspace = current + 1
				laststart = lastspace
			} else if zend.ZSTR_VAL(text)[current] == ' ' {
				if current-laststart >= linelength {
					zend.ZSTR_VAL(newtext)[current] = breakchar[0]
					laststart = current + 1
				}
				lastspace = current
			} else if current-laststart >= linelength && laststart != lastspace {
				zend.ZSTR_VAL(newtext)[lastspace] = breakchar[0]
				laststart = lastspace + 1
			}
		}
		zend.RETVAL_NEW_STR(newtext)
		return
	} else {

		/* Multiple character line break or forced cut */

		if linelength > 0 {
			chk = size_t(zend.ZSTR_LEN(text)/linelength + 1)
			newtext = zend.ZendStringSafeAlloc(chk, breakchar_len, zend.ZSTR_LEN(text), 0)
			alloced = zend.ZSTR_LEN(text) + chk*breakchar_len + 1
		} else {
			chk = zend.ZSTR_LEN(text)
			alloced = zend.ZSTR_LEN(text)*(breakchar_len+1) + 1
			newtext = zend.ZendStringSafeAlloc(zend.ZSTR_LEN(text), breakchar_len+1, 0, 0)
		}

		/* now keep track of the actual new text length */

		newtextlen = 0
		lastspace = 0
		laststart = lastspace
		for current = 0; current < zend.ZendLong(zend.ZSTR_LEN(text)); current++ {
			if chk == 0 {
				alloced += size_t(((zend.ZSTR_LEN(text)-current+1)/linelength+1)*breakchar_len) + 1
				newtext = zend.ZendStringExtend(newtext, alloced, 0)
				chk = size_t((zend.ZSTR_LEN(text)-current)/linelength) + 1
			}

			/* when we hit an existing break, copy to new buffer, and
			 * fix up laststart and lastspace */

			if zend.ZSTR_VAL(text)[current] == breakchar[0] && current+breakchar_len < zend.ZSTR_LEN(text) && !(strncmp(zend.ZSTR_VAL(text)+current, breakchar, breakchar_len)) {
				memcpy(zend.ZSTR_VAL(newtext)+newtextlen, zend.ZSTR_VAL(text)+laststart, current-laststart+breakchar_len)
				newtextlen += current - laststart + breakchar_len
				current += breakchar_len - 1
				lastspace = current + 1
				laststart = lastspace
				chk--
			} else if zend.ZSTR_VAL(text)[current] == ' ' {
				if current-laststart >= linelength {
					memcpy(zend.ZSTR_VAL(newtext)+newtextlen, zend.ZSTR_VAL(text)+laststart, current-laststart)
					newtextlen += current - laststart
					memcpy(zend.ZSTR_VAL(newtext)+newtextlen, breakchar, breakchar_len)
					newtextlen += breakchar_len
					laststart = current + 1
					chk--
				}
				lastspace = current
			} else if current-laststart >= linelength && docut != 0 && laststart >= lastspace {
				memcpy(zend.ZSTR_VAL(newtext)+newtextlen, zend.ZSTR_VAL(text)+laststart, current-laststart)
				newtextlen += current - laststart
				memcpy(zend.ZSTR_VAL(newtext)+newtextlen, breakchar, breakchar_len)
				newtextlen += breakchar_len
				lastspace = current
				laststart = lastspace
				chk--
			} else if current-laststart >= linelength && laststart < lastspace {
				memcpy(zend.ZSTR_VAL(newtext)+newtextlen, zend.ZSTR_VAL(text)+laststart, lastspace-laststart)
				newtextlen += lastspace - laststart
				memcpy(zend.ZSTR_VAL(newtext)+newtextlen, breakchar, breakchar_len)
				newtextlen += breakchar_len
				lastspace = lastspace + 1
				laststart = lastspace
				chk--
			}

			/* when we hit an existing break, copy to new buffer, and
			 * fix up laststart and lastspace */

		}

		/* copy over any stragglers */

		if laststart != current {
			memcpy(zend.ZSTR_VAL(newtext)+newtextlen, zend.ZSTR_VAL(text)+laststart, current-laststart)
			newtextlen += current - laststart
		}
		zend.ZSTR_VAL(newtext)[newtextlen] = '0'

		/* free unused memory */

		newtext = zend.ZendStringTruncate(newtext, newtextlen, 0)
		zend.RETVAL_NEW_STR(newtext)
		return
	}

	/* Special case for a single-character break as it needs no
	   additional storage space */
}
func PhpExplode(delim *zend.ZendString, str *zend.ZendString, return_value *zend.Zval, limit zend.ZendLong) {
	var p1 *byte = zend.ZSTR_VAL(str)
	var endp *byte = zend.ZSTR_VAL(str) + zend.ZSTR_LEN(str)
	var p2 *byte = core.PhpMemnstr(zend.ZSTR_VAL(str), zend.ZSTR_VAL(delim), zend.ZSTR_LEN(delim), endp)
	var tmp zend.Zval
	if p2 == nil {
		zend.ZVAL_STR_COPY(&tmp, str)
		zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), &tmp)
	} else {
		for {
			var l int = p2 - p1
			if l == 0 {
				zend.ZVAL_EMPTY_STRING(&tmp)
			} else if l == 1 {
				zend.ZVAL_INTERNED_STR(&tmp, zend.ZSTR_CHAR(zend_uchar(*p1)))
			} else {
				zend.ZVAL_STRINGL(&tmp, p1, p2-p1)
			}
			zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), &tmp)
			p1 = p2 + zend.ZSTR_LEN(delim)
			p2 = core.PhpMemnstr(p1, zend.ZSTR_VAL(delim), zend.ZSTR_LEN(delim), endp)
			if !(p2 != nil && b.PreDec(&limit) > 1) {
				break
			}
		}
		if p1 <= endp {
			zend.ZVAL_STRINGL(&tmp, p1, endp-p1)
			zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), &tmp)
		}
	}
}
func PhpExplodeNegativeLimit(delim *zend.ZendString, str *zend.ZendString, return_value *zend.Zval, limit zend.ZendLong) {
	// #define EXPLODE_ALLOC_STEP       64

	var p1 *byte = zend.ZSTR_VAL(str)
	var endp *byte = zend.ZSTR_VAL(str) + zend.ZSTR_LEN(str)
	var p2 *byte = core.PhpMemnstr(zend.ZSTR_VAL(str), zend.ZSTR_VAL(delim), zend.ZSTR_LEN(delim), endp)
	var tmp zend.Zval
	if p2 == nil {

	} else {
		var allocated int = 64
		var found int = 0
		var i zend.ZendLong
		var to_return zend.ZendLong
		var positions **byte = zend.Emalloc(allocated * b.SizeOf("char *"))
		positions[b.PostInc(&found)] = p1
		for {
			if found >= allocated {
				allocated = found + 64
				positions = zend.Erealloc(positions, allocated*b.SizeOf("char *"))
			}
			p1 = p2 + zend.ZSTR_LEN(delim)
			positions[b.PostInc(&found)] = p1
			p2 = core.PhpMemnstr(p1, zend.ZSTR_VAL(delim), zend.ZSTR_LEN(delim), endp)
			if p2 == nil {
				break
			}
		}
		to_return = limit + found

		/* limit is at least -1 therefore no need of bounds checking : i will be always less than found */

		for i = 0; i < to_return; i++ {
			zend.ZVAL_STRINGL(&tmp, positions[i], positions[i+1]-zend.ZSTR_LEN(delim)-positions[i])
			zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), &tmp)
		}
		zend.Efree(any(positions))
	}
}
func ZifExplode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var delim *zend.ZendString
	var limit zend.ZendLong = zend.ZEND_LONG_MAX
	var tmp zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &delim, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &limit, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if zend.ZSTR_LEN(delim) == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Empty delimiter")
		zend.RETVAL_FALSE
		return
	}
	zend.ArrayInit(return_value)
	if zend.ZSTR_LEN(str) == 0 {
		if limit >= 0 {
			zend.ZVAL_EMPTY_STRING(&tmp)
			zend.ZendHashIndexAddNew(zend.Z_ARRVAL_P(return_value), 0, &tmp)
		}
		return
	}
	if limit > 1 {
		PhpExplode(delim, str, return_value, limit)
	} else if limit < 0 {
		PhpExplodeNegativeLimit(delim, str, return_value, limit)
	} else {
		zend.ZVAL_STR_COPY(&tmp, str)
		zend.ZendHashIndexAddNew(zend.Z_ARRVAL_P(return_value), 0, &tmp)
	}
}
func PhpImplode(glue *zend.ZendString, pieces *zend.Zval, return_value *zend.Zval) {
	var tmp *zend.Zval
	var numelems int
	var str *zend.ZendString
	var cptr *byte
	var len_ int = 0
	var strings *struct {
		str  *zend.ZendString
		lval zend.ZendLong
	}
	var ptr *struct {
		str  *zend.ZendString
		lval zend.ZendLong
	}
	numelems = zend.ZendHashNumElements(zend.Z_ARRVAL_P(pieces))
	if numelems == 0 {
		zend.RETVAL_EMPTY_STRING()
		return
	} else if numelems == 1 {

		/* loop to search the first not undefined element... */

		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(pieces)
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
					_z = zend.Z_INDIRECT_P(_z)
				}
				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				tmp = _z
				zend.RETVAL_STR(zend.ZvalGetString(tmp))
				return
			}
			break
		}

		/* loop to search the first not undefined element... */

	}
	strings = zend.DoAlloca(b.SizeOf("* strings")*numelems, use_heap)
	ptr = strings
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL_P(pieces)
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			tmp = _z
			if zend.EXPECTED(zend.Z_TYPE_P(tmp) == zend.IS_STRING) {
				ptr.str = zend.Z_STR_P(tmp)
				len_ += zend.ZSTR_LEN(ptr.str)
				ptr.lval = 0
				ptr++
			} else if zend.UNEXPECTED(zend.Z_TYPE_P(tmp) == zend.IS_LONG) {
				var val zend.ZendLong = zend.Z_LVAL_P(tmp)
				ptr.str = nil
				ptr.lval = val
				ptr++
				if val <= 0 {
					len_++
				}
				for val != 0 {
					val /= 10
					len_++
				}
			} else {
				ptr.str = zend.ZvalGetStringFunc(tmp)
				len_ += zend.ZSTR_LEN(ptr.str)
				ptr.lval = 1
				ptr++
			}
		}
		break
	}

	/* numelems can not be 0, we checked above */

	str = zend.ZendStringSafeAlloc(numelems-1, zend.ZSTR_LEN(glue), len_, 0)
	cptr = zend.ZSTR_VAL(str) + zend.ZSTR_LEN(str)
	*cptr = 0
	for true {
		ptr--
		if zend.EXPECTED(ptr.str != nil) {
			cptr -= zend.ZSTR_LEN(ptr.str)
			memcpy(cptr, zend.ZSTR_VAL(ptr.str), zend.ZSTR_LEN(ptr.str))
			if ptr.lval != 0 {
				zend.ZendStringReleaseEx(ptr.str, 0)
			}
		} else {
			var oldPtr *byte = cptr
			var oldVal byte = *cptr
			cptr = zend.ZendPrintLongToBuf(cptr, ptr.lval)
			*oldPtr = oldVal
		}
		if ptr == strings {
			break
		}
		cptr -= zend.ZSTR_LEN(glue)
		memcpy(cptr, zend.ZSTR_VAL(glue), zend.ZSTR_LEN(glue))
	}
	zend.FreeAlloca(strings, use_heap)
	zend.RETVAL_NEW_STR(str)
	return
}
func ZifImplode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg1 *zend.Zval
	var arg2 *zend.Zval = nil
	var pieces *zend.Zval
	var glue *zend.ZendString
	var tmp_glue *zend.ZendString
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
			zend.ZendParseArgZvalDeref(_arg, &arg1, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &arg2, 0)
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
	if arg2 == nil {
		if zend.Z_TYPE_P(arg1) != zend.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Argument must be an array")
			return
		}
		glue = zend.ZSTR_EMPTY_ALLOC()
		tmp_glue = nil
		pieces = arg1
	} else {
		if zend.Z_TYPE_P(arg1) == zend.IS_ARRAY {
			glue = zend.ZvalGetTmpString(arg2, &tmp_glue)
			pieces = arg1
			core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Passing glue string after array is deprecated. Swap the parameters")
		} else if zend.Z_TYPE_P(arg2) == zend.IS_ARRAY {
			glue = zend.ZvalGetTmpString(arg1, &tmp_glue)
			pieces = arg2
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid arguments passed")
			return
		}
	}
	PhpImplode(glue, pieces, return_value)
	zend.ZendTmpStringRelease(tmp_glue)
}
func STRTOK_TABLE(p *byte) __auto__ { return BG(strtok_table)[uint8(*p)] }
func ZifStrtok(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var tok *zend.ZendString = nil
	var token *byte
	var token_end *byte
	var p *byte
	var pe *byte
	var skipped int = 0
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &tok, 0) == 0) {
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
	if zend.ZEND_NUM_ARGS() == 1 {
		tok = str
	} else {
		zend.ZvalPtrDtor(&BG(strtok_zval))
		zend.ZVAL_STRINGL(&BG(strtok_zval), zend.ZSTR_VAL(str), zend.ZSTR_LEN(str))
		BG(strtok_string) = zend.Z_STRVAL(BG(strtok_zval))
		BG(strtok_last) = BG(strtok_string)
		BG(strtok_len) = zend.ZSTR_LEN(str)
	}
	p = BG(strtok_last)
	pe = BG(strtok_string) + BG(strtok_len)
	if p == nil || p >= pe {
		zend.RETVAL_FALSE
		return
	}
	token = zend.ZSTR_VAL(tok)
	token_end = token + zend.ZSTR_LEN(tok)
	for token < token_end {
		STRTOK_TABLE(b.PostInc(&token)) = 1
	}

	/* Skip leading delimiters */

	for STRTOK_TABLE(p) {
		if b.PreInc(&p) >= pe {

			/* no other chars left */

			BG(strtok_last) = nil
			zend.RETVAL_FALSE
			goto restore
		}
		skipped++
	}

	/* We know at this place that *p is no delimiter, so skip it */

	for b.PreInc(&p) < pe {
		if STRTOK_TABLE(p) {
			goto return_token
		}
	}
	if p-BG(strtok_last) != 0 {
	return_token:
		zend.RETVAL_STRINGL(BG(strtok_last)+skipped, p-BG(strtok_last)-skipped)
		BG(strtok_last) = p + 1
	} else {
		zend.RETVAL_FALSE
		BG(strtok_last) = nil
	}

	/* Restore table -- usually faster then memset'ing the table on every invocation */

restore:
	token = zend.ZSTR_VAL(tok)
	for token < token_end {
		STRTOK_TABLE(b.PostInc(&token)) = 0
	}
}
func PhpStrtoupper(s *byte, len_ int) *byte {
	var c *uint8
	var e *uint8
	c = (*uint8)(s)
	e = (*uint8)(c + len_)
	for c < e {
		*c = toupper(*c)
		c++
	}
	return s
}
func PhpStringToupper(s *zend.ZendString) *zend.ZendString {
	var c *uint8
	var e *uint8
	c = (*uint8)(zend.ZSTR_VAL(s))
	e = c + zend.ZSTR_LEN(s)
	for c < e {
		if islower(*c) {
			var r *uint8
			var res *zend.ZendString = zend.ZendStringAlloc(zend.ZSTR_LEN(s), 0)
			if c != (*uint8)(zend.ZSTR_VAL(s)) {
				memcpy(zend.ZSTR_VAL(res), zend.ZSTR_VAL(s), c-(*uint8)(zend.ZSTR_VAL(s)))
			}
			r = c + (zend.ZSTR_VAL(res) - zend.ZSTR_VAL(s))
			for c < e {
				*r = toupper(*c)
				r++
				c++
			}
			*r = '0'
			return res
		}
		c++
	}
	return zend.ZendStringCopy(s)
}
func ZifStrtoupper(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &arg, 0) == 0) {
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
	zend.RETVAL_STR(PhpStringToupper(arg))
	return
}
func PhpStrtolower(s *byte, len_ int) *byte {
	var c *uint8
	var e *uint8
	c = (*uint8)(s)
	e = c + len_
	for c < e {
		*c = tolower(*c)
		c++
	}
	return s
}
func PhpStringTolower(s *zend.ZendString) *zend.ZendString {
	var c *uint8
	var e *uint8
	c = (*uint8)(zend.ZSTR_VAL(s))
	e = c + zend.ZSTR_LEN(s)
	for c < e {
		if isupper(*c) {
			var r *uint8
			var res *zend.ZendString = zend.ZendStringAlloc(zend.ZSTR_LEN(s), 0)
			if c != (*uint8)(zend.ZSTR_VAL(s)) {
				memcpy(zend.ZSTR_VAL(res), zend.ZSTR_VAL(s), c-(*uint8)(zend.ZSTR_VAL(s)))
			}
			r = c + (zend.ZSTR_VAL(res) - zend.ZSTR_VAL(s))
			for c < e {
				*r = tolower(*c)
				r++
				c++
			}
			*r = '0'
			return res
		}
		c++
	}
	return zend.ZendStringCopy(s)
}
func ZifStrtolower(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
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
	zend.RETVAL_STR(PhpStringTolower(str))
	return
}
func PhpBasename(s *byte, len_ int, suffix *byte, sufflen int) *zend.ZendString {
	var c *byte
	var comp *byte
	var cend *byte
	var inc_len int
	var cnt int
	var state int
	var ret *zend.ZendString
	c = (*byte)(s)
	cend = c
	comp = cend
	cnt = len_
	state = 0
	for cnt > 0 {
		if (*c) == '0' {
			inc_len = 1
		} else {
			inc_len = PhpMblen(c, cnt)
		}
		switch inc_len {
		case -2:

		case -1:
			inc_len = 1
			core.PhpIgnoreValue(mblen(nil, 0))
			break
		case 0:
			goto quit_loop
		case 1:
			if (*c) == '/' {
				if state == 1 {
					state = 0
					cend = c
				}
			} else {
				if state == 0 {
					comp = c
					state = 1
				}
			}
			break
		default:
			if state == 0 {
				comp = c
				state = 1
			}
			break
		}
		c += inc_len
		cnt -= inc_len
	}
quit_loop:
	if state == 1 {
		cend = c
	}
	if suffix != nil && sufflen < size_t(cend-comp) && memcmp(cend-sufflen, suffix, sufflen) == 0 {
		cend -= sufflen
	}
	len_ = cend - comp
	ret = zend.ZendStringInit(comp, len_, 0)
	return ret
}
func ZifBasename(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var string *byte
	var suffix *byte = nil
	var string_len int
	var suffix_len int = 0
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
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &string, &string_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &suffix, &suffix_len, 0) == 0) {
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
	zend.RETVAL_STR(PhpBasename(string, string_len, suffix, suffix_len))
	return
}
func PhpDirname(path *byte, len_ int) int { return zend.ZendDirname(path, len_) }
func ZifDirname(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte
	var str_len int
	var ret *zend.ZendString
	var levels zend.ZendLong = 1
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
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &levels, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	ret = zend.ZendStringInit(str, str_len, 0)
	if levels == 1 {

		/* Default case */

		zend.ZSTR_LEN(ret) = zend.ZendDirname(zend.ZSTR_VAL(ret), str_len)

		/* Default case */

	} else if levels < 1 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid argument, levels must be >= 1")
		zend.ZendStringEfree(ret)
		return
	} else {

		/* Some levels up */

		for {
			zend.ZSTR_LEN(ret) = zend.ZendDirname(zend.ZSTR_VAL(ret), b.Assign(&str_len, zend.ZSTR_LEN(ret)))
			if !(zend.ZSTR_LEN(ret) < str_len && b.PreDec(&levels)) {
				break
			}
		}

		/* Some levels up */

	}
	zend.RETVAL_NEW_STR(ret)
	return
}
func ZifPathinfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var tmp zend.Zval
	var path *byte
	var dirname *byte
	var path_len int
	var have_basename int
	var opt zend.ZendLong = PHP_PATHINFO_ALL
	var ret *zend.ZendString = nil
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
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &path, &path_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &opt, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	have_basename = (opt & PHP_PATHINFO_BASENAME) == PHP_PATHINFO_BASENAME
	zend.ArrayInit(&tmp)
	if (opt & PHP_PATHINFO_DIRNAME) == PHP_PATHINFO_DIRNAME {
		dirname = zend.Estrndup(path, path_len)
		PhpDirname(dirname, path_len)
		if *dirname {
			zend.AddAssocString(&tmp, "dirname", dirname)
		}
		zend.Efree(dirname)
	}
	if have_basename != 0 {
		ret = PhpBasename(path, path_len, nil, 0)
		zend.AddAssocStr(&tmp, "basename", zend.ZendStringCopy(ret))
	}
	if (opt & PHP_PATHINFO_EXTENSION) == PHP_PATHINFO_EXTENSION {
		var p *byte
		var idx ptrdiff_t
		if have_basename == 0 {
			ret = PhpBasename(path, path_len, nil, 0)
		}
		p = zend.ZendMemrchr(zend.ZSTR_VAL(ret), '.', zend.ZSTR_LEN(ret))
		if p != nil {
			idx = p - zend.ZSTR_VAL(ret)
			zend.AddAssocStringl(&tmp, "extension", zend.ZSTR_VAL(ret)+idx+1, zend.ZSTR_LEN(ret)-idx-1)
		}
	}
	if (opt & PHP_PATHINFO_FILENAME) == PHP_PATHINFO_FILENAME {
		var p *byte
		var idx ptrdiff_t

		/* Have we already looked up the basename? */

		if have_basename == 0 && ret == nil {
			ret = PhpBasename(path, path_len, nil, 0)
		}
		p = zend.ZendMemrchr(zend.ZSTR_VAL(ret), '.', zend.ZSTR_LEN(ret))
		if p != nil {
			idx = p - zend.ZSTR_VAL(ret)
		} else {
			idx = ptrdiff_t(zend.ZSTR_LEN(ret))
		}
		zend.AddAssocStringl(&tmp, "filename", zend.ZSTR_VAL(ret), idx)
	}
	if ret != nil {
		zend.ZendStringReleaseEx(ret, 0)
	}
	if opt == PHP_PATHINFO_ALL {
		zend.ZVAL_COPY_VALUE(return_value, &tmp)
	} else {
		var element *zend.Zval
		if b.Assign(&element, zend.ZendHashGetCurrentData(zend.Z_ARRVAL(tmp))) != nil {
			zend.ZVAL_COPY_DEREF(return_value, element)
		} else {
			zend.ZVAL_EMPTY_STRING(return_value)
		}
		zend.ZvalPtrDtor(&tmp)
	}
}
func PhpStristr(s *byte, t *byte, s_len int, t_len int) *byte {
	PhpStrtolower(s, s_len)
	PhpStrtolower(t, t_len)
	return (*byte)(core.PhpMemnstr(s, t, t_len, s+s_len))
}
func PhpStrspn(s1 *byte, s2 *byte, s1_end *byte, s2_end *byte) int {
	var p *byte = s1
	var spanp *byte
	var c byte = *p
cont:
	for spanp = s2; p != s1_end && spanp != s2_end; {
		if b.PostInc(&(*spanp)) == c {
			c = *(b.PreInc(&p))
			goto cont
		}
	}
	return p - s1
}
func PhpStrcspn(s1 *byte, s2 *byte, s1_end *byte, s2_end *byte) int {
	var p *byte
	var spanp *byte
	var c byte = *s1
	for p = s1; ; {
		spanp = s2
		for {
			if (*spanp) == c || p == s1_end {
				return p - s1
			}
			if b.PostInc(&spanp) >= s2_end-1 {
				break
			}
		}
		c = *(b.PreInc(&p))
	}
}
func PhpNeedleChar(needle *zend.Zval, target *byte) int {
	switch zend.Z_TYPE_P(needle) {
	case zend.IS_LONG:
		*target = byte(zend.Z_LVAL_P(needle))
		return zend.SUCCESS
	case zend.IS_NULL:

	case zend.IS_FALSE:
		*target = '0'
		return zend.SUCCESS
	case zend.IS_TRUE:
		*target = '1'
		return zend.SUCCESS
	case zend.IS_DOUBLE:

	case zend.IS_OBJECT:
		*target = byte(zend.ZvalGetLong(needle))
		return zend.SUCCESS
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "needle is not a string or an integer")
		return zend.FAILURE
	}
}
func ZifStristr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var needle *zend.Zval
	var haystack *zend.ZendString
	var found *byte = nil
	var found_offset int
	var haystack_dup *byte
	var needle_char []byte
	var part zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &haystack, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &needle, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &part, &_dummy, 0) == 0) {
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
	haystack_dup = zend.Estrndup(zend.ZSTR_VAL(haystack), zend.ZSTR_LEN(haystack))
	if zend.Z_TYPE_P(needle) == zend.IS_STRING {
		var orig_needle *byte
		if zend.Z_STRLEN_P(needle) == 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Empty needle")
			zend.Efree(haystack_dup)
			zend.RETVAL_FALSE
			return
		}
		orig_needle = zend.Estrndup(zend.Z_STRVAL_P(needle), zend.Z_STRLEN_P(needle))
		found = PhpStristr(haystack_dup, orig_needle, zend.ZSTR_LEN(haystack), zend.Z_STRLEN_P(needle))
		zend.Efree(orig_needle)
	} else {
		if PhpNeedleChar(needle, needle_char) != zend.SUCCESS {
			zend.Efree(haystack_dup)
			zend.RETVAL_FALSE
			return
		}
		needle_char[1] = 0
		core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		found = PhpStristr(haystack_dup, needle_char, zend.ZSTR_LEN(haystack), 1)
	}
	if found != nil {
		found_offset = found - haystack_dup
		if part != 0 {
			zend.RETVAL_STRINGL(zend.ZSTR_VAL(haystack), found_offset)
		} else {
			zend.RETVAL_STRINGL(zend.ZSTR_VAL(haystack)+found_offset, zend.ZSTR_LEN(haystack)-found_offset)
		}
	} else {
		zend.RETVAL_FALSE
	}
	zend.Efree(haystack_dup)
}
func ZifStrstr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var needle *zend.Zval
	var haystack *zend.ZendString
	var found *byte = nil
	var needle_char []byte
	var found_offset zend.ZendLong
	var part zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &haystack, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &needle, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &part, &_dummy, 0) == 0) {
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
	if zend.Z_TYPE_P(needle) == zend.IS_STRING {
		if zend.Z_STRLEN_P(needle) == 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Empty needle")
			zend.RETVAL_FALSE
			return
		}
		found = core.PhpMemnstr(zend.ZSTR_VAL(haystack), zend.Z_STRVAL_P(needle), zend.Z_STRLEN_P(needle), zend.ZSTR_VAL(haystack)+zend.ZSTR_LEN(haystack))
	} else {
		if PhpNeedleChar(needle, needle_char) != zend.SUCCESS {
			zend.RETVAL_FALSE
			return
		}
		needle_char[1] = 0
		core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		found = core.PhpMemnstr(zend.ZSTR_VAL(haystack), needle_char, 1, zend.ZSTR_VAL(haystack)+zend.ZSTR_LEN(haystack))
	}
	if found != nil {
		found_offset = found - zend.ZSTR_VAL(haystack)
		if part != 0 {
			zend.RETVAL_STRINGL(zend.ZSTR_VAL(haystack), found_offset)
			return
		} else {
			zend.RETVAL_STRINGL(found, zend.ZSTR_LEN(haystack)-found_offset)
			return
		}
	}
	zend.RETVAL_FALSE
	return
}
func ZifStrpos(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var needle *zend.Zval
	var haystack *zend.ZendString
	var found *byte = nil
	var needle_char []byte
	var offset zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &haystack, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &needle, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if offset < 0 {
		offset += zend.ZendLong(zend.ZSTR_LEN(haystack))
	}
	if offset < 0 || int(offset > zend.ZSTR_LEN(haystack)) != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Offset not contained in string")
		zend.RETVAL_FALSE
		return
	}
	if zend.Z_TYPE_P(needle) == zend.IS_STRING {
		if zend.Z_STRLEN_P(needle) == 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Empty needle")
			zend.RETVAL_FALSE
			return
		}
		found = (*byte)(core.PhpMemnstr(zend.ZSTR_VAL(haystack)+offset, zend.Z_STRVAL_P(needle), zend.Z_STRLEN_P(needle), zend.ZSTR_VAL(haystack)+zend.ZSTR_LEN(haystack)))
	} else {
		if PhpNeedleChar(needle, needle_char) != zend.SUCCESS {
			zend.RETVAL_FALSE
			return
		}
		needle_char[1] = 0
		core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		found = (*byte)(core.PhpMemnstr(zend.ZSTR_VAL(haystack)+offset, needle_char, 1, zend.ZSTR_VAL(haystack)+zend.ZSTR_LEN(haystack)))
	}
	if found != nil {
		zend.RETVAL_LONG(found - zend.ZSTR_VAL(haystack))
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func ZifStripos(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var found *byte = nil
	var haystack *zend.ZendString
	var offset zend.ZendLong = 0
	var needle_char []byte
	var needle *zend.Zval
	var needle_dup *zend.ZendString = nil
	var haystack_dup *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &haystack, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &needle, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if offset < 0 {
		offset += zend.ZendLong(zend.ZSTR_LEN(haystack))
	}
	if offset < 0 || int(offset > zend.ZSTR_LEN(haystack)) != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Offset not contained in string")
		zend.RETVAL_FALSE
		return
	}
	if zend.ZSTR_LEN(haystack) == 0 {
		zend.RETVAL_FALSE
		return
	}
	if zend.Z_TYPE_P(needle) == zend.IS_STRING {
		if zend.Z_STRLEN_P(needle) == 0 || zend.Z_STRLEN_P(needle) > zend.ZSTR_LEN(haystack) {
			zend.RETVAL_FALSE
			return
		}
		haystack_dup = PhpStringTolower(haystack)
		needle_dup = PhpStringTolower(zend.Z_STR_P(needle))
		found = (*byte)(core.PhpMemnstr(zend.ZSTR_VAL(haystack_dup)+offset, zend.ZSTR_VAL(needle_dup), zend.ZSTR_LEN(needle_dup), zend.ZSTR_VAL(haystack_dup)+zend.ZSTR_LEN(haystack)))
	} else {
		if PhpNeedleChar(needle, needle_char) != zend.SUCCESS {
			zend.RETVAL_FALSE
			return
		}
		core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		haystack_dup = PhpStringTolower(haystack)
		needle_char[0] = tolower(needle_char[0])
		needle_char[1] = '0'
		found = (*byte)(core.PhpMemnstr(zend.ZSTR_VAL(haystack_dup)+offset, needle_char, b.SizeOf("needle_char")-1, zend.ZSTR_VAL(haystack_dup)+zend.ZSTR_LEN(haystack)))
	}
	if found != nil {
		zend.RETVAL_LONG(found - zend.ZSTR_VAL(haystack_dup))
	} else {
		zend.RETVAL_FALSE
	}
	zend.ZendStringReleaseEx(haystack_dup, 0)
	if needle_dup != nil {
		zend.ZendStringReleaseEx(needle_dup, 0)
	}
}
func ZifStrrpos(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zneedle *zend.Zval
	var haystack *zend.ZendString
	var needle_len int
	var offset zend.ZendLong = 0
	var ord_needle []byte
	var p *byte
	var e *byte
	var found *byte
	var needle *byte
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &haystack, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zneedle, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if zend.Z_TYPE_P(zneedle) == zend.IS_STRING {
		needle = zend.Z_STRVAL_P(zneedle)
		needle_len = zend.Z_STRLEN_P(zneedle)
	} else {
		if PhpNeedleChar(zneedle, ord_needle) != zend.SUCCESS {
			zend.RETVAL_FALSE
			return
		}
		core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		ord_needle[1] = '0'
		needle = ord_needle
		needle_len = 1
	}
	if zend.ZSTR_LEN(haystack) == 0 || needle_len == 0 {
		zend.RETVAL_FALSE
		return
	}
	if offset >= 0 {
		if int(offset > zend.ZSTR_LEN(haystack)) != 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Offset is greater than the length of haystack string")
			zend.RETVAL_FALSE
			return
		}
		p = zend.ZSTR_VAL(haystack) + int(offset)
		e = zend.ZSTR_VAL(haystack) + zend.ZSTR_LEN(haystack)
	} else {
		if offset < -core.INT_MAX || size_t(-offset) > zend.ZSTR_LEN(haystack) {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Offset is greater than the length of haystack string")
			zend.RETVAL_FALSE
			return
		}
		p = zend.ZSTR_VAL(haystack)
		if size_t-offset < needle_len {
			e = zend.ZSTR_VAL(haystack) + zend.ZSTR_LEN(haystack)
		} else {
			e = zend.ZSTR_VAL(haystack) + zend.ZSTR_LEN(haystack) + offset + needle_len
		}
	}
	if b.Assign(&found, zend.ZendMemnrstr(p, needle, needle_len, e)) {
		zend.RETVAL_LONG(found - zend.ZSTR_VAL(haystack))
		return
	}
	zend.RETVAL_FALSE
	return
}
func ZifStrripos(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zneedle *zend.Zval
	var needle *zend.ZendString
	var haystack *zend.ZendString
	var offset zend.ZendLong = 0
	var p *byte
	var e *byte
	var found *byte
	var needle_dup *zend.ZendString
	var haystack_dup *zend.ZendString
	var ord_needle *zend.ZendString = nil
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &haystack, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zneedle, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	zend.ZSTR_ALLOCA_ALLOC(ord_needle, 1, use_heap)
	if zend.Z_TYPE_P(zneedle) == zend.IS_STRING {
		needle = zend.Z_STR_P(zneedle)
	} else {
		if PhpNeedleChar(zneedle, zend.ZSTR_VAL(ord_needle)) != zend.SUCCESS {
			zend.ZSTR_ALLOCA_FREE(ord_needle, use_heap)
			zend.RETVAL_FALSE
			return
		}
		core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		zend.ZSTR_VAL(ord_needle)[1] = '0'
		needle = ord_needle
	}
	if zend.ZSTR_LEN(haystack) == 0 || zend.ZSTR_LEN(needle) == 0 {
		zend.ZSTR_ALLOCA_FREE(ord_needle, use_heap)
		zend.RETVAL_FALSE
		return
	}
	if zend.ZSTR_LEN(needle) == 1 {

		/* Single character search can shortcut memcmps
		   Can also avoid tolower emallocs */

		if offset >= 0 {
			if int(offset > zend.ZSTR_LEN(haystack)) != 0 {
				zend.ZSTR_ALLOCA_FREE(ord_needle, use_heap)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Offset is greater than the length of haystack string")
				zend.RETVAL_FALSE
				return
			}
			p = zend.ZSTR_VAL(haystack) + int(offset)
			e = zend.ZSTR_VAL(haystack) + zend.ZSTR_LEN(haystack) - 1
		} else {
			p = zend.ZSTR_VAL(haystack)
			if offset < -core.INT_MAX || size_t(-offset) > zend.ZSTR_LEN(haystack) {
				zend.ZSTR_ALLOCA_FREE(ord_needle, use_heap)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Offset is greater than the length of haystack string")
				zend.RETVAL_FALSE
				return
			}
			e = zend.ZSTR_VAL(haystack) + (zend.ZSTR_LEN(haystack) + int(offset))
		}

		/* Borrow that ord_needle buffer to avoid repeatedly tolower()ing needle */

		(*zend.ZSTR_VAL)(ord_needle) = tolower((*zend.ZSTR_VAL)(needle))
		for e >= p {
			if tolower(*e) == (*zend.ZSTR_VAL)(ord_needle) {
				zend.ZSTR_ALLOCA_FREE(ord_needle, use_heap)
				zend.RETVAL_LONG(e - p + b.Cond(offset > 0, offset, 0))
				return
			}
			e--
		}
		zend.ZSTR_ALLOCA_FREE(ord_needle, use_heap)
		zend.RETVAL_FALSE
		return
	}
	haystack_dup = PhpStringTolower(haystack)
	if offset >= 0 {
		if int(offset > zend.ZSTR_LEN(haystack)) != 0 {
			zend.ZendStringReleaseEx(haystack_dup, 0)
			zend.ZSTR_ALLOCA_FREE(ord_needle, use_heap)
			core.PhpErrorDocref(nil, zend.E_WARNING, "Offset is greater than the length of haystack string")
			zend.RETVAL_FALSE
			return
		}
		p = zend.ZSTR_VAL(haystack_dup) + offset
		e = zend.ZSTR_VAL(haystack_dup) + zend.ZSTR_LEN(haystack)
	} else {
		if offset < -core.INT_MAX || size_t(-offset) > zend.ZSTR_LEN(haystack) {
			zend.ZendStringReleaseEx(haystack_dup, 0)
			zend.ZSTR_ALLOCA_FREE(ord_needle, use_heap)
			core.PhpErrorDocref(nil, zend.E_WARNING, "Offset is greater than the length of haystack string")
			zend.RETVAL_FALSE
			return
		}
		p = zend.ZSTR_VAL(haystack_dup)
		if size_t-offset < zend.ZSTR_LEN(needle) {
			e = zend.ZSTR_VAL(haystack_dup) + zend.ZSTR_LEN(haystack)
		} else {
			e = zend.ZSTR_VAL(haystack_dup) + zend.ZSTR_LEN(haystack) + offset + zend.ZSTR_LEN(needle)
		}
	}
	needle_dup = PhpStringTolower(needle)
	if b.Assign(&found, (*byte)(zend.ZendMemnrstr(p, zend.ZSTR_VAL(needle_dup), zend.ZSTR_LEN(needle_dup), e))) {
		zend.RETVAL_LONG(found - zend.ZSTR_VAL(haystack_dup))
		zend.ZendStringReleaseEx(needle_dup, 0)
		zend.ZendStringReleaseEx(haystack_dup, 0)
		zend.ZSTR_ALLOCA_FREE(ord_needle, use_heap)
	} else {
		zend.ZendStringReleaseEx(needle_dup, 0)
		zend.ZendStringReleaseEx(haystack_dup, 0)
		zend.ZSTR_ALLOCA_FREE(ord_needle, use_heap)
		zend.RETVAL_FALSE
		return
	}
}
func ZifStrrchr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var needle *zend.Zval
	var haystack *zend.ZendString
	var found *byte = nil
	var found_offset zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &haystack, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &needle, 0)
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
	if zend.Z_TYPE_P(needle) == zend.IS_STRING {
		found = zend.ZendMemrchr(zend.ZSTR_VAL(haystack), (*zend.Z_STRVAL_P)(needle), zend.ZSTR_LEN(haystack))
	} else {
		var needle_chr byte
		if PhpNeedleChar(needle, &needle_chr) != zend.SUCCESS {
			zend.RETVAL_FALSE
			return
		}
		core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		found = zend.ZendMemrchr(zend.ZSTR_VAL(haystack), needle_chr, zend.ZSTR_LEN(haystack))
	}
	if found != nil {
		found_offset = found - zend.ZSTR_VAL(haystack)
		zend.RETVAL_STRINGL(found, zend.ZSTR_LEN(haystack)-found_offset)
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func PhpChunkSplit(src *byte, srclen int, end *byte, endlen int, chunklen int) *zend.ZendString {
	var q *byte
	var p *byte
	var chunks int
	var restlen int
	var out_len int
	var dest *zend.ZendString
	chunks = srclen / chunklen
	restlen = srclen - chunks*chunklen
	if chunks > core.INT_MAX-1 {
		return nil
	}
	out_len = chunks + 1
	if endlen != 0 && out_len > core.INT_MAX/endlen {
		return nil
	}
	out_len *= endlen
	if out_len > core.INT_MAX-srclen-1 {
		return nil
	}
	out_len += srclen + 1
	dest = zend.ZendStringAlloc(out_len*b.SizeOf("char"), 0)
	p = src
	q = zend.ZSTR_VAL(dest)
	for p < src+srclen-chunklen+1 {
		memcpy(q, p, chunklen)
		q += chunklen
		memcpy(q, end, endlen)
		q += endlen
		p += chunklen
	}
	if restlen != 0 {
		memcpy(q, p, restlen)
		q += restlen
		memcpy(q, end, endlen)
		q += endlen
	}
	*q = '0'
	zend.ZSTR_LEN(dest) = q - zend.ZSTR_VAL(dest)
	return dest
}
func ZifChunkSplit(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var end *byte = "\r\n"
	var endlen int = 2
	var chunklen zend.ZendLong = 76
	var result *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &chunklen, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &end, &endlen, 0) == 0) {
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
	if chunklen <= 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Chunk length should be greater than zero")
		zend.RETVAL_FALSE
		return
	}
	if int(chunklen > zend.ZSTR_LEN(str)) != 0 {

		/* to maintain BC, we must return original string + ending */

		result = zend.ZendStringSafeAlloc(zend.ZSTR_LEN(str), 1, endlen, 0)
		memcpy(zend.ZSTR_VAL(result), zend.ZSTR_VAL(str), zend.ZSTR_LEN(str))
		memcpy(zend.ZSTR_VAL(result)+zend.ZSTR_LEN(str), end, endlen)
		zend.ZSTR_VAL(result)[zend.ZSTR_LEN(result)] = '0'
		zend.RETVAL_NEW_STR(result)
		return
	}
	if zend.ZSTR_LEN(str) == 0 {
		zend.RETVAL_EMPTY_STRING()
		return
	}
	result = PhpChunkSplit(zend.ZSTR_VAL(str), zend.ZSTR_LEN(str), end, endlen, int(chunklen))
	if result != nil {
		zend.RETVAL_STR(result)
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func ZifSubstr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var l zend.ZendLong = 0
	var f zend.ZendLong
	var argc int = zend.ZEND_NUM_ARGS()
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &f, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &l, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if f > zend.ZendLong(zend.ZSTR_LEN(str)) {
		zend.RETVAL_FALSE
		return
	} else if f < 0 {

		/* if "from" position is negative, count start position from the end
		 * of the string
		 */

		if size_t-f > zend.ZSTR_LEN(str) {
			f = 0
		} else {
			f = zend.ZendLong(zend.ZSTR_LEN(str) + f)
		}
		if argc > 2 {
			if l < 0 {

				/* if "length" position is negative, set it to the length
				 * needed to stop that many chars from the end of the string
				 */

				if size_t(-l) > zend.ZSTR_LEN(str)-int(f) {
					if size_t(-l) > zend.ZSTR_LEN(str) {
						zend.RETVAL_FALSE
						return
					} else {
						l = 0
					}
				} else {
					l = zend.ZendLong(zend.ZSTR_LEN(str) - f + l)
				}

				/* if "length" position is negative, set it to the length
				 * needed to stop that many chars from the end of the string
				 */

			} else if int(l > zend.ZSTR_LEN(str)-int(f)) != 0 {
				goto truncate_len
			}
		} else {
			goto truncate_len
		}
	} else if argc > 2 {
		if l < 0 {

			/* if "length" position is negative, set it to the length
			 * needed to stop that many chars from the end of the string
			 */

			if size_t(-l) > zend.ZSTR_LEN(str)-int(f) {
				zend.RETVAL_FALSE
				return
			} else {
				l = zend.ZendLong(zend.ZSTR_LEN(str) - f + l)
			}

			/* if "length" position is negative, set it to the length
			 * needed to stop that many chars from the end of the string
			 */

		} else if int(l > zend.ZSTR_LEN(str)-int(f)) != 0 {
			goto truncate_len
		}
	} else {
	truncate_len:
		l = zend.ZendLong(zend.ZSTR_LEN(str) - f)
	}
	if l == 0 {
		zend.RETVAL_EMPTY_STRING()
		return
	} else if l == 1 {
		zend.RETVAL_INTERNED_STR(zend.ZSTR_CHAR(zend_uchar(zend.ZSTR_VAL(str)[f])))
		return
	} else if l == zend.ZSTR_LEN(str) {
		zend.RETVAL_STR_COPY(str)
		return
	}
	zend.RETVAL_STRINGL(zend.ZSTR_VAL(str)+f, l)
	return
}
func ZifSubstrReplace(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.Zval
	var from *zend.Zval
	var len_ *zend.Zval = nil
	var repl *zend.Zval
	var l zend.ZendLong = 0
	var f zend.ZendLong
	var argc int = zend.ZEND_NUM_ARGS()
	var result *zend.ZendString
	var from_idx zend.HashPosition
	var repl_idx zend.HashPosition
	var len_idx zend.HashPosition
	var tmp_str *zend.Zval = nil
	var tmp_from *zend.Zval = nil
	var tmp_repl *zend.Zval = nil
	var tmp_len *zend.Zval = nil
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 4
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
			zend.ZendParseArgZvalDeref(_arg, &str, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &repl, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &from, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &len_, 0)
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
	if zend.Z_TYPE_P(str) != zend.IS_ARRAY {
		zend.ConvertToStringEx(str)
	}
	if zend.Z_TYPE_P(repl) != zend.IS_ARRAY {
		zend.ConvertToStringEx(repl)
	}
	if zend.Z_TYPE_P(from) != zend.IS_ARRAY {
		if zend.Z_TYPE_P(from) != zend.IS_LONG {
			zend.ConvertToLong(from)
		}
	}
	if zend.ExecutorGlobals.exception != nil {
		return
	}
	if argc > 3 {
		if zend.Z_TYPE_P(len_) != zend.IS_ARRAY {
			if zend.Z_TYPE_P(len_) != zend.IS_LONG {
				zend.ConvertToLong(len_)
			}
			l = zend.Z_LVAL_P(len_)
		}
	} else {
		if zend.Z_TYPE_P(str) != zend.IS_ARRAY {
			l = zend.Z_STRLEN_P(str)
		}
	}
	if zend.Z_TYPE_P(str) == zend.IS_STRING {
		if argc == 3 && zend.Z_TYPE_P(from) == zend.IS_ARRAY || argc == 4 && zend.Z_TYPE_P(from) != zend.Z_TYPE_P(len_) {
			core.PhpErrorDocref(nil, zend.E_WARNING, "'start' and 'length' should be of same type - numerical or array ")
			zend.RETVAL_STR_COPY(zend.Z_STR_P(str))
			return
		}
		if argc == 4 && zend.Z_TYPE_P(from) == zend.IS_ARRAY {
			if zend.ZendHashNumElements(zend.Z_ARRVAL_P(from)) != zend.ZendHashNumElements(zend.Z_ARRVAL_P(len_)) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "'start' and 'length' should have the same number of elements")
				zend.RETVAL_STR_COPY(zend.Z_STR_P(str))
				return
			}
		}
	}
	if zend.Z_TYPE_P(str) != zend.IS_ARRAY {
		if zend.Z_TYPE_P(from) != zend.IS_ARRAY {
			var repl_str *zend.ZendString
			var tmp_repl_str *zend.ZendString = nil
			f = zend.Z_LVAL_P(from)

			/* if "from" position is negative, count start position from the end
			 * of the string
			 */

			if f < 0 {
				f = zend.ZendLong(zend.Z_STRLEN_P(str) + f)
				if f < 0 {
					f = 0
				}
			} else if int(f > zend.Z_STRLEN_P(str)) != 0 {
				f = zend.Z_STRLEN_P(str)
			}

			/* if "length" position is negative, set it to the length
			 * needed to stop that many chars from the end of the string
			 */

			if l < 0 {
				l = zend.ZendLong(zend.Z_STRLEN_P(str)-f) + l
				if l < 0 {
					l = 0
				}
			}
			if int(l > zend.Z_STRLEN_P(str) || l < 0 && size_t(-l) > zend.Z_STRLEN_P(str)) != 0 {
				l = zend.Z_STRLEN_P(str)
			}
			if f+l > zend.ZendLong(zend.Z_STRLEN_P(str)) {
				l = zend.Z_STRLEN_P(str) - f
			}
			if zend.Z_TYPE_P(repl) == zend.IS_ARRAY {
				repl_idx = 0
				for repl_idx < zend.Z_ARRVAL_P(repl).nNumUsed {
					tmp_repl = &zend.Z_ARRVAL_P(repl).arData[repl_idx].val
					if zend.Z_TYPE_P(tmp_repl) != zend.IS_UNDEF {
						break
					}
					repl_idx++
				}
				if repl_idx < zend.Z_ARRVAL_P(repl).nNumUsed {
					repl_str = zend.ZvalGetTmpString(tmp_repl, &tmp_repl_str)
				} else {
					repl_str = zend.STR_EMPTY_ALLOC()
				}
			} else {
				repl_str = zend.Z_STR_P(repl)
			}
			result = zend.ZendStringSafeAlloc(1, zend.Z_STRLEN_P(str)-l+zend.ZSTR_LEN(repl_str), 0, 0)
			memcpy(zend.ZSTR_VAL(result), zend.Z_STRVAL_P(str), f)
			if zend.ZSTR_LEN(repl_str) != 0 {
				memcpy(zend.ZSTR_VAL(result)+f, zend.ZSTR_VAL(repl_str), zend.ZSTR_LEN(repl_str))
			}
			memcpy(zend.ZSTR_VAL(result)+f+zend.ZSTR_LEN(repl_str), zend.Z_STRVAL_P(str)+f+l, zend.Z_STRLEN_P(str)-f-l)
			zend.ZSTR_VAL(result)[zend.ZSTR_LEN(result)] = '0'
			zend.ZendTmpStringRelease(tmp_repl_str)
			zend.RETVAL_NEW_STR(result)
			return
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Functionality of 'start' and 'length' as arrays is not implemented")
			zend.RETVAL_STR_COPY(zend.Z_STR_P(str))
			return
		}
	} else {
		var str_index *zend.ZendString = nil
		var result_len int
		var num_index zend.ZendUlong
		zend.ArrayInit(return_value)
		repl_idx = 0
		len_idx = repl_idx
		from_idx = len_idx
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(str)
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
					_z = zend.Z_INDIRECT_P(_z)
				}
				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				num_index = _p.h
				str_index = _p.key
				tmp_str = _z
				var tmp_orig_str *zend.ZendString
				var orig_str *zend.ZendString = zend.ZvalGetTmpString(tmp_str, &tmp_orig_str)
				if zend.Z_TYPE_P(from) == zend.IS_ARRAY {
					for from_idx < zend.Z_ARRVAL_P(from).nNumUsed {
						tmp_from = &zend.Z_ARRVAL_P(from).arData[from_idx].val
						if zend.Z_TYPE_P(tmp_from) != zend.IS_UNDEF {
							break
						}
						from_idx++
					}
					if from_idx < zend.Z_ARRVAL_P(from).nNumUsed {
						f = zend.ZvalGetLong(tmp_from)
						if f < 0 {
							f = zend.ZendLong(zend.ZSTR_LEN(orig_str) + f)
							if f < 0 {
								f = 0
							}
						} else if f > zend.ZendLong(zend.ZSTR_LEN(orig_str)) {
							f = zend.ZSTR_LEN(orig_str)
						}
						from_idx++
					} else {
						f = 0
					}
				} else {
					f = zend.Z_LVAL_P(from)
					if f < 0 {
						f = zend.ZendLong(zend.ZSTR_LEN(orig_str) + f)
						if f < 0 {
							f = 0
						}
					} else if f > zend.ZendLong(zend.ZSTR_LEN(orig_str)) {
						f = zend.ZSTR_LEN(orig_str)
					}
				}
				if argc > 3 && zend.Z_TYPE_P(len_) == zend.IS_ARRAY {
					for len_idx < zend.Z_ARRVAL_P(len_).nNumUsed {
						tmp_len = &zend.Z_ARRVAL_P(len_).arData[len_idx].val
						if zend.Z_TYPE_P(tmp_len) != zend.IS_UNDEF {
							break
						}
						len_idx++
					}
					if len_idx < zend.Z_ARRVAL_P(len_).nNumUsed {
						l = zend.ZvalGetLong(tmp_len)
						len_idx++
					} else {
						l = zend.ZSTR_LEN(orig_str)
					}
				} else if argc > 3 {
					l = zend.Z_LVAL_P(len_)
				} else {
					l = zend.ZSTR_LEN(orig_str)
				}
				if l < 0 {
					l = zend.ZSTR_LEN(orig_str) - f + l
					if l < 0 {
						l = 0
					}
				}
				zend.ZEND_ASSERT(0 <= f && f <= zend.ZEND_LONG_MAX)
				zend.ZEND_ASSERT(0 <= l && l <= zend.ZEND_LONG_MAX)
				if int(f+l) > zend.ZSTR_LEN(orig_str) {
					l = zend.ZSTR_LEN(orig_str) - f
				}
				result_len = zend.ZSTR_LEN(orig_str) - l
				if zend.Z_TYPE_P(repl) == zend.IS_ARRAY {
					for repl_idx < zend.Z_ARRVAL_P(repl).nNumUsed {
						tmp_repl = &zend.Z_ARRVAL_P(repl).arData[repl_idx].val
						if zend.Z_TYPE_P(tmp_repl) != zend.IS_UNDEF {
							break
						}
						repl_idx++
					}
					if repl_idx < zend.Z_ARRVAL_P(repl).nNumUsed {
						var tmp_repl_str *zend.ZendString
						var repl_str *zend.ZendString = zend.ZvalGetTmpString(tmp_repl, &tmp_repl_str)
						result_len += zend.ZSTR_LEN(repl_str)
						repl_idx++
						result = zend.ZendStringSafeAlloc(1, result_len, 0, 0)
						memcpy(zend.ZSTR_VAL(result), zend.ZSTR_VAL(orig_str), f)
						memcpy(zend.ZSTR_VAL(result)+f, zend.ZSTR_VAL(repl_str), zend.ZSTR_LEN(repl_str))
						memcpy(zend.ZSTR_VAL(result)+f+zend.ZSTR_LEN(repl_str), zend.ZSTR_VAL(orig_str)+f+l, zend.ZSTR_LEN(orig_str)-f-l)
						zend.ZendTmpStringRelease(tmp_repl_str)
					} else {
						result = zend.ZendStringSafeAlloc(1, result_len, 0, 0)
						memcpy(zend.ZSTR_VAL(result), zend.ZSTR_VAL(orig_str), f)
						memcpy(zend.ZSTR_VAL(result)+f, zend.ZSTR_VAL(orig_str)+f+l, zend.ZSTR_LEN(orig_str)-f-l)
					}
				} else {
					result_len += zend.Z_STRLEN_P(repl)
					result = zend.ZendStringSafeAlloc(1, result_len, 0, 0)
					memcpy(zend.ZSTR_VAL(result), zend.ZSTR_VAL(orig_str), f)
					memcpy(zend.ZSTR_VAL(result)+f, zend.Z_STRVAL_P(repl), zend.Z_STRLEN_P(repl))
					memcpy(zend.ZSTR_VAL(result)+f+zend.Z_STRLEN_P(repl), zend.ZSTR_VAL(orig_str)+f+l, zend.ZSTR_LEN(orig_str)-f-l)
				}
				zend.ZSTR_VAL(result)[zend.ZSTR_LEN(result)] = '0'
				if str_index != nil {
					var tmp zend.Zval
					zend.ZVAL_NEW_STR(&tmp, result)
					zend.ZendSymtableUpdate(zend.Z_ARRVAL_P(return_value), str_index, &tmp)
				} else {
					zend.AddIndexStr(return_value, num_index, result)
				}
				zend.ZendTmpStringRelease(tmp_orig_str)
			}
			break
		}
	}
}
func ZifQuotemeta(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var old *zend.ZendString
	var old_end *byte
	var p *byte
	var q *byte
	var c byte
	var str *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &old, 0) == 0) {
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
	old_end = zend.ZSTR_VAL(old) + zend.ZSTR_LEN(old)
	if zend.ZSTR_VAL(old) == old_end {
		zend.RETVAL_FALSE
		return
	}
	str = zend.ZendStringSafeAlloc(2, zend.ZSTR_LEN(old), 0, 0)
	p = zend.ZSTR_VAL(old)
	q = zend.ZSTR_VAL(str)
	for ; p != old_end; p++ {
		c = *p
		switch c {
		case '.':

		case '\\':

		case '+':

		case '*':

		case '?':

		case '[':

		case '^':

		case ']':

		case '$':

		case '(':

		case ')':
			b.PostInc(&(*q)) = '\\'
		default:
			b.PostInc(&(*q)) = c
		}
	}
	*q = '0'
	zend.RETVAL_NEW_STR(zend.ZendStringTruncate(str, q-zend.ZSTR_VAL(str), 0))
	return
}
func ZifOrd(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
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
	zend.RETVAL_LONG(uint8(zend.ZSTR_VAL(str)[0]))
	return
}
func ZifChr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var c zend.ZendLong
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
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &c, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
			c = 0
		}
		break
	}
	c &= 0xff
	zend.ZVAL_INTERNED_STR(return_value, zend.ZSTR_CHAR(c))
}
func PhpUcfirst(str *zend.ZendString) *zend.ZendString {
	var ch uint8 = zend.ZSTR_VAL(str)[0]
	var r uint8 = toupper(ch)
	if r == ch {
		return zend.ZendStringCopy(str)
	} else {
		var s *zend.ZendString = zend.ZendStringInit(zend.ZSTR_VAL(str), zend.ZSTR_LEN(str), 0)
		zend.ZSTR_VAL(s)[0] = r
		return s
	}
}
func ZifUcfirst(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
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
	if zend.ZSTR_LEN(str) == 0 {
		zend.RETVAL_EMPTY_STRING()
		return
	}
	zend.RETVAL_STR(PhpUcfirst(str))
	return
}
func PhpLcfirst(str *zend.ZendString) *zend.ZendString {
	var r uint8 = tolower(zend.ZSTR_VAL(str)[0])
	if r == zend.ZSTR_VAL(str)[0] {
		return zend.ZendStringCopy(str)
	} else {
		var s *zend.ZendString = zend.ZendStringInit(zend.ZSTR_VAL(str), zend.ZSTR_LEN(str), 0)
		zend.ZSTR_VAL(s)[0] = r
		return s
	}
}
func ZifLcfirst(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
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
	if zend.ZSTR_LEN(str) == 0 {
		zend.RETVAL_EMPTY_STRING()
		return
	}
	zend.RETVAL_STR(PhpLcfirst(str))
	return
}
func ZifUcwords(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var delims *byte = " \t\r\nfv"
	var r *byte
	var r_end *byte
	var delims_len int = 6
	var mask []byte
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &delims, &delims_len, 0) == 0) {
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
	if zend.ZSTR_LEN(str) == 0 {
		zend.RETVAL_EMPTY_STRING()
		return
	}
	PhpCharmask((*uint8)(delims), delims_len, mask)
	zend.ZVAL_STRINGL(return_value, zend.ZSTR_VAL(str), zend.ZSTR_LEN(str))
	r = zend.Z_STRVAL_P(return_value)
	*r = toupper(uint8(*r))
	for r_end = r + zend.Z_STRLEN_P(return_value) - 1; r < r_end; {
		if mask[uint8(b.PostInc(&(*r)))] {
			*r = toupper(uint8(*r))
		}
	}
}
func PhpStrtr(str *byte, len_ int, str_from *byte, str_to *byte, trlen int) *byte {
	var i int
	if zend.UNEXPECTED(trlen < 1) {
		return str
	} else if trlen == 1 {
		var ch_from byte = *str_from
		var ch_to byte = *str_to
		for i = 0; i < len_; i++ {
			if str[i] == ch_from {
				str[i] = ch_to
			}
		}
	} else {
		var xlat []uint8
		var j uint8 = 0
		for {
			xlat[j] = j
			if b.PreInc(&j) == 0 {
				break
			}
		}
		for i = 0; i < trlen; i++ {
			xlat[int(uint8(str_from[i]))] = str_to[i]
		}
		for i = 0; i < len_; i++ {
			str[i] = xlat[int(uint8(str[i]))]
		}
	}
	return str
}
func PhpStrtrEx(str *zend.ZendString, str_from *byte, str_to *byte, trlen int) *zend.ZendString {
	var new_str *zend.ZendString = nil
	var i int
	if zend.UNEXPECTED(trlen < 1) {
		return zend.ZendStringCopy(str)
	} else if trlen == 1 {
		var ch_from byte = *str_from
		var ch_to byte = *str_to
		for i = 0; i < zend.ZSTR_LEN(str); i++ {
			if zend.ZSTR_VAL(str)[i] == ch_from {
				new_str = zend.ZendStringAlloc(zend.ZSTR_LEN(str), 0)
				memcpy(zend.ZSTR_VAL(new_str), zend.ZSTR_VAL(str), i)
				zend.ZSTR_VAL(new_str)[i] = ch_to
				break
			}
		}
		for ; i < zend.ZSTR_LEN(str); i++ {
			if zend.ZSTR_VAL(str)[i] != ch_from {
				zend.ZSTR_VAL(new_str)[i] = zend.ZSTR_VAL(str)[i]
			} else {
				zend.ZSTR_VAL(new_str)[i] = ch_to
			}
		}
	} else {
		var xlat []uint8
		var j uint8 = 0
		for {
			xlat[j] = j
			if b.PreInc(&j) == 0 {
				break
			}
		}
		for i = 0; i < trlen; i++ {
			xlat[int(uint8(str_from[i]))] = str_to[i]
		}
		for i = 0; i < zend.ZSTR_LEN(str); i++ {
			if zend.ZSTR_VAL(str)[i] != xlat[int(uint8(zend.ZSTR_VAL(str)[i]))] {
				new_str = zend.ZendStringAlloc(zend.ZSTR_LEN(str), 0)
				memcpy(zend.ZSTR_VAL(new_str), zend.ZSTR_VAL(str), i)
				zend.ZSTR_VAL(new_str)[i] = xlat[int(uint8(zend.ZSTR_VAL(str)[i]))]
				break
			}
		}
		for ; i < zend.ZSTR_LEN(str); i++ {
			zend.ZSTR_VAL(new_str)[i] = xlat[int(uint8(zend.ZSTR_VAL(str)[i]))]
		}
	}
	if new_str == nil {
		return zend.ZendStringCopy(str)
	}
	zend.ZSTR_VAL(new_str)[zend.ZSTR_LEN(new_str)] = 0
	return new_str
}
func PhpStrtrArray(return_value *zend.Zval, input *zend.ZendString, pats *zend.HashTable) {
	var str *byte = zend.ZSTR_VAL(input)
	var slen int = zend.ZSTR_LEN(input)
	var num_key zend.ZendUlong
	var str_key *zend.ZendString
	var len_ int
	var pos int
	var old_pos int
	var num_keys int = 0
	var minlen int = 128 * 1024
	var maxlen int = 0
	var str_hash zend.HashTable
	var entry *zend.Zval
	var key *byte
	var result zend.SmartStr = zend.SmartStr{0}
	var bitset []zend.ZendUlong
	var num_bitset *zend.ZendUlong

	/* we will collect all possible key lengths */

	num_bitset = zend.Ecalloc((slen+b.SizeOf("zend_ulong"))/b.SizeOf("zend_ulong"), b.SizeOf("zend_ulong"))
	memset(bitset, 0, b.SizeOf("bitset"))

	/* check if original array has numeric keys */

	for {
		var __ht *zend.HashTable = pats
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			str_key = _p.key
			if zend.UNEXPECTED(str_key == nil) {
				num_keys = 1
			} else {
				len_ = zend.ZSTR_LEN(str_key)
				if zend.UNEXPECTED(len_ < 1) {
					zend.Efree(num_bitset)
					zend.RETVAL_FALSE
					return
				} else if zend.UNEXPECTED(len_ > slen) {

					/* skip long patterns */

					continue

					/* skip long patterns */

				}
				if len_ > maxlen {
					maxlen = len_
				}
				if len_ < minlen {
					minlen = len_
				}

				/* remember possible key length */

				num_bitset[len_/b.SizeOf("zend_ulong")] |= uint64(1) << len_ % b.SizeOf("zend_ulong")
				bitset[uint8(zend.ZSTR_VAL(str_key)[0])/b.SizeOf("zend_ulong")] |= uint64(1) << uint8(zend.ZSTR_VAL(str_key)[0]) % b.SizeOf("zend_ulong")
			}
		}
		break
	}
	if zend.UNEXPECTED(num_keys != 0) {
		var key_used *zend.ZendString

		/* we have to rebuild HashTable with numeric keys */

		zend.ZendHashInit(&str_hash, zend.ZendHashNumElements(pats), nil, nil, 0)
		for {
			var __ht *zend.HashTable = pats
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
					_z = zend.Z_INDIRECT_P(_z)
				}
				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				num_key = _p.h
				str_key = _p.key
				entry = _z
				if zend.UNEXPECTED(str_key == nil) {
					key_used = zend.ZendLongToStr(num_key)
					len_ = zend.ZSTR_LEN(key_used)
					if zend.UNEXPECTED(len_ > slen) {

						/* skip long patterns */

						zend.ZendStringRelease(key_used)
						continue
					}
					if len_ > maxlen {
						maxlen = len_
					}
					if len_ < minlen {
						minlen = len_
					}

					/* remember possible key length */

					num_bitset[len_/b.SizeOf("zend_ulong")] |= uint64(1) << len_ % b.SizeOf("zend_ulong")
					bitset[uint8(zend.ZSTR_VAL(key_used)[0])/b.SizeOf("zend_ulong")] |= uint64(1) << uint8(zend.ZSTR_VAL(key_used)[0]) % b.SizeOf("zend_ulong")
				} else {
					key_used = str_key
					len_ = zend.ZSTR_LEN(key_used)
					if zend.UNEXPECTED(len_ > slen) {

						/* skip long patterns */

						continue

						/* skip long patterns */

					}
				}
				zend.ZendHashAdd(&str_hash, key_used, entry)
				if zend.UNEXPECTED(str_key == nil) {
					zend.ZendStringReleaseEx(key_used, 0)
				}
			}
			break
		}
		pats = &str_hash
	}
	if zend.UNEXPECTED(minlen > maxlen) {

		/* return the original string */

		if pats == &str_hash {
			zend.ZendHashDestroy(&str_hash)
		}
		zend.Efree(num_bitset)
		zend.RETVAL_STR_COPY(input)
		return
	}
	pos = 0
	old_pos = pos
	for pos <= slen-minlen {
		key = str + pos
		if (bitset[uint8(key[0])/b.SizeOf("zend_ulong")] & uint64(1) << uint8(key[0]) % b.SizeOf("zend_ulong")) != 0 {
			len_ = maxlen
			if len_ > slen-pos {
				len_ = slen - pos
			}
			for len_ >= minlen {
				if (num_bitset[len_/b.SizeOf("zend_ulong")] & uint64(1) << len_ % b.SizeOf("zend_ulong")) != 0 {
					entry = zend.ZendHashStrFind(pats, key, len_)
					if entry != nil {
						var tmp *zend.ZendString
						var s *zend.ZendString = zend.ZvalGetTmpString(entry, &tmp)
						zend.SmartStrAppendl(&result, str+old_pos, pos-old_pos)
						zend.SmartStrAppend(&result, s)
						old_pos = pos + len_
						pos = old_pos - 1
						zend.ZendTmpStringRelease(tmp)
						break
					}
				}
				len_--
			}
		}
		pos++
	}
	if result.s != nil {
		zend.SmartStrAppendl(&result, str+old_pos, slen-old_pos)
		zend.SmartStr0(&result)
		zend.RETVAL_NEW_STR(result.s)
	} else {
		zend.SmartStrFree(&result)
		zend.RETVAL_STR_COPY(input)
	}
	if pats == &str_hash {
		zend.ZendHashDestroy(&str_hash)
	}
	zend.Efree(num_bitset)
}
func PhpCharToStrEx(str *zend.ZendString, from byte, to *byte, to_len int, case_sensitivity int, replace_count *zend.ZendLong) *zend.ZendString {
	var result *zend.ZendString
	var char_count int = 0
	var lc_from int = 0
	var source *byte
	var source_end *byte = zend.ZSTR_VAL(str) + zend.ZSTR_LEN(str)
	var target *byte
	if case_sensitivity != 0 {
		var p *byte = zend.ZSTR_VAL(str)
		var e *byte = p + zend.ZSTR_LEN(str)
		for b.Assign(&p, memchr(p, from, e-p)) {
			char_count++
			p++
		}
	} else {
		lc_from = tolower(from)
		for source = zend.ZSTR_VAL(str); source < source_end; source++ {
			if tolower(*source) == lc_from {
				char_count++
			}
		}
	}
	if char_count == 0 {
		return zend.ZendStringCopy(str)
	}
	if to_len > 0 {
		result = zend.ZendStringSafeAlloc(char_count, to_len-1, zend.ZSTR_LEN(str), 0)
	} else {
		result = zend.ZendStringAlloc(zend.ZSTR_LEN(str)-char_count, 0)
	}
	target = zend.ZSTR_VAL(result)
	if case_sensitivity != 0 {
		var p *byte = zend.ZSTR_VAL(str)
		var e *byte = p + zend.ZSTR_LEN(str)
		var s *byte = zend.ZSTR_VAL(str)
		for b.Assign(&p, memchr(p, from, e-p)) {
			memcpy(target, s, p-s)
			target += p - s
			memcpy(target, to, to_len)
			target += to_len
			p++
			s = p
			if replace_count != nil {
				*replace_count += 1
			}
		}
		if s < e {
			memcpy(target, s, e-s)
			target += e - s
		}
	} else {
		for source = zend.ZSTR_VAL(str); source < source_end; source++ {
			if tolower(*source) == lc_from {
				if replace_count != nil {
					*replace_count += 1
				}
				memcpy(target, to, to_len)
				target += to_len
			} else {
				*target = *source
				target++
			}
		}
	}
	*target = 0
	return result
}
func PhpStrToStrEx(haystack *zend.ZendString, needle *byte, needle_len int, str *byte, str_len int, replace_count *zend.ZendLong) *zend.ZendString {
	var new_str *zend.ZendString
	if needle_len < zend.ZSTR_LEN(haystack) {
		var end *byte
		var p *byte
		var r *byte
		var e *byte
		if needle_len == str_len {
			new_str = nil
			end = zend.ZSTR_VAL(haystack) + zend.ZSTR_LEN(haystack)
			for p = zend.ZSTR_VAL(haystack); b.Assign(&r, (*byte)(core.PhpMemnstr(p, needle, needle_len, end))); p = r + needle_len {
				if new_str == nil {
					new_str = zend.ZendStringInit(zend.ZSTR_VAL(haystack), zend.ZSTR_LEN(haystack), 0)
				}
				memcpy(zend.ZSTR_VAL(new_str)+(r-zend.ZSTR_VAL(haystack)), str, str_len)
				*replace_count++
			}
			if new_str == nil {
				goto nothing_todo
			}
			return new_str
		} else {
			var count int = 0
			var o *byte = zend.ZSTR_VAL(haystack)
			var n *byte = needle
			var endp *byte = o + zend.ZSTR_LEN(haystack)
			for b.Assign(&o, (*byte)(core.PhpMemnstr(o, n, needle_len, endp))) {
				o += needle_len
				count++
			}
			if count == 0 {

				/* Needle doesn't occur, shortcircuit the actual replacement. */

				goto nothing_todo

				/* Needle doesn't occur, shortcircuit the actual replacement. */

			}
			if str_len > needle_len {
				new_str = zend.ZendStringSafeAlloc(count, str_len-needle_len, zend.ZSTR_LEN(haystack), 0)
			} else {
				new_str = zend.ZendStringAlloc(count*(str_len-needle_len)+zend.ZSTR_LEN(haystack), 0)
			}
			e = zend.ZSTR_VAL(new_str)
			end = zend.ZSTR_VAL(haystack) + zend.ZSTR_LEN(haystack)
			for p = zend.ZSTR_VAL(haystack); b.Assign(&r, (*byte)(core.PhpMemnstr(p, needle, needle_len, end))); p = r + needle_len {
				memcpy(e, p, r-p)
				e += r - p
				memcpy(e, str, str_len)
				e += str_len
				*replace_count++
			}
			if p < end {
				memcpy(e, p, end-p)
				e += end - p
			}
			*e = '0'
			return new_str
		}
	} else if needle_len > zend.ZSTR_LEN(haystack) || memcmp(zend.ZSTR_VAL(haystack), needle, zend.ZSTR_LEN(haystack)) {
	nothing_todo:
		return zend.ZendStringCopy(haystack)
	} else {
		if str_len == 0 {
			new_str = zend.ZSTR_EMPTY_ALLOC()
		} else if str_len == 1 {
			new_str = zend.ZSTR_CHAR(zend_uchar(*str))
		} else {
			new_str = zend.ZendStringInit(str, str_len, 0)
		}
		*replace_count++
		return new_str
	}
}
func PhpStrToStrIEx(haystack *zend.ZendString, lc_haystack *byte, needle *zend.ZendString, str *byte, str_len int, replace_count *zend.ZendLong) *zend.ZendString {
	var new_str *zend.ZendString = nil
	var lc_needle *zend.ZendString
	if zend.ZSTR_LEN(needle) < zend.ZSTR_LEN(haystack) {
		var end *byte
		var p *byte
		var r *byte
		var e *byte
		if zend.ZSTR_LEN(needle) == str_len {
			lc_needle = PhpStringTolower(needle)
			end = lc_haystack + zend.ZSTR_LEN(haystack)
			for p = lc_haystack; b.Assign(&r, (*byte)(core.PhpMemnstr(p, zend.ZSTR_VAL(lc_needle), zend.ZSTR_LEN(lc_needle), end))); p = r + zend.ZSTR_LEN(lc_needle) {
				if new_str == nil {
					new_str = zend.ZendStringInit(zend.ZSTR_VAL(haystack), zend.ZSTR_LEN(haystack), 0)
				}
				memcpy(zend.ZSTR_VAL(new_str)+(r-lc_haystack), str, str_len)
				*replace_count++
			}
			zend.ZendStringReleaseEx(lc_needle, 0)
			if new_str == nil {
				goto nothing_todo
			}
			return new_str
		} else {
			var count int = 0
			var o *byte = lc_haystack
			var n *byte
			var endp *byte = o + zend.ZSTR_LEN(haystack)
			lc_needle = PhpStringTolower(needle)
			n = zend.ZSTR_VAL(lc_needle)
			for b.Assign(&o, (*byte)(core.PhpMemnstr(o, n, zend.ZSTR_LEN(lc_needle), endp))) {
				o += zend.ZSTR_LEN(lc_needle)
				count++
			}
			if count == 0 {

				/* Needle doesn't occur, shortcircuit the actual replacement. */

				zend.ZendStringReleaseEx(lc_needle, 0)
				goto nothing_todo
			}
			if str_len > zend.ZSTR_LEN(lc_needle) {
				new_str = zend.ZendStringSafeAlloc(count, str_len-zend.ZSTR_LEN(lc_needle), zend.ZSTR_LEN(haystack), 0)
			} else {
				new_str = zend.ZendStringAlloc(count*(str_len-zend.ZSTR_LEN(lc_needle))+zend.ZSTR_LEN(haystack), 0)
			}
			e = zend.ZSTR_VAL(new_str)
			end = lc_haystack + zend.ZSTR_LEN(haystack)
			for p = lc_haystack; b.Assign(&r, (*byte)(core.PhpMemnstr(p, zend.ZSTR_VAL(lc_needle), zend.ZSTR_LEN(lc_needle), end))); p = r + zend.ZSTR_LEN(lc_needle) {
				memcpy(e, zend.ZSTR_VAL(haystack)+(p-lc_haystack), r-p)
				e += r - p
				memcpy(e, str, str_len)
				e += str_len
				*replace_count++
			}
			if p < end {
				memcpy(e, zend.ZSTR_VAL(haystack)+(p-lc_haystack), end-p)
				e += end - p
			}
			*e = '0'
			zend.ZendStringReleaseEx(lc_needle, 0)
			return new_str
		}
	} else if zend.ZSTR_LEN(needle) > zend.ZSTR_LEN(haystack) {
	nothing_todo:
		return zend.ZendStringCopy(haystack)
	} else {
		lc_needle = PhpStringTolower(needle)
		if memcmp(lc_haystack, zend.ZSTR_VAL(lc_needle), zend.ZSTR_LEN(lc_needle)) {
			zend.ZendStringReleaseEx(lc_needle, 0)
			goto nothing_todo
		}
		zend.ZendStringReleaseEx(lc_needle, 0)
		new_str = zend.ZendStringInit(str, str_len, 0)
		*replace_count++
		return new_str
	}
}
func PhpStrToStr(haystack *byte, length int, needle string, needle_len int, str string, str_len int) *zend.ZendString {
	var new_str *zend.ZendString
	if needle_len < length {
		var end *byte
		var s *byte
		var p *byte
		var e *byte
		var r *byte
		if needle_len == str_len {
			new_str = zend.ZendStringInit(haystack, length, 0)
			end = zend.ZSTR_VAL(new_str) + length
			for p = zend.ZSTR_VAL(new_str); b.Assign(&r, (*byte)(core.PhpMemnstr(p, needle, needle_len, end))); p = r + needle_len {
				memcpy(r, str, str_len)
			}
			return new_str
		} else {
			if str_len < needle_len {
				new_str = zend.ZendStringAlloc(length, 0)
			} else {
				var count int = 0
				var o *byte = haystack
				var n *byte = needle
				var endp *byte = o + length
				for b.Assign(&o, (*byte)(core.PhpMemnstr(o, n, needle_len, endp))) {
					o += needle_len
					count++
				}
				if count == 0 {

					/* Needle doesn't occur, shortcircuit the actual replacement. */

					new_str = zend.ZendStringInit(haystack, length, 0)
					return new_str
				} else {
					if str_len > needle_len {
						new_str = zend.ZendStringSafeAlloc(count, str_len-needle_len, length, 0)
					} else {
						new_str = zend.ZendStringAlloc(count*(str_len-needle_len)+length, 0)
					}
				}
			}
			e = zend.ZSTR_VAL(new_str)
			s = e
			end = haystack + length
			for p = haystack; b.Assign(&r, (*byte)(core.PhpMemnstr(p, needle, needle_len, end))); p = r + needle_len {
				memcpy(e, p, r-p)
				e += r - p
				memcpy(e, str, str_len)
				e += str_len
			}
			if p < end {
				memcpy(e, p, end-p)
				e += end - p
			}
			*e = '0'
			new_str = zend.ZendStringTruncate(new_str, e-s, 0)
			return new_str
		}
	} else if needle_len > length || memcmp(haystack, needle, length) {
		new_str = zend.ZendStringInit(haystack, length, 0)
		return new_str
	} else {
		new_str = zend.ZendStringInit(str, str_len, 0)
		return new_str
	}
}
func ZifStrtr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var from *zend.Zval
	var str *zend.ZendString
	var to *byte = nil
	var to_len int = 0
	var ac int = zend.ZEND_NUM_ARGS()
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &from, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &to, &to_len, 0) == 0) {
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
	if ac == 2 && zend.Z_TYPE_P(from) != zend.IS_ARRAY {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The second argument is not an array")
		zend.RETVAL_FALSE
		return
	}

	/* shortcut for empty string */

	if zend.ZSTR_LEN(str) == 0 {
		zend.RETVAL_EMPTY_STRING()
		return
	}
	if ac == 2 {
		var pats *zend.HashTable = zend.Z_ARRVAL_P(from)
		if zend.ZendHashNumElements(pats) < 1 {
			zend.RETVAL_STR_COPY(str)
			return
		} else if zend.ZendHashNumElements(pats) == 1 {
			var num_key zend.ZendLong
			var str_key *zend.ZendString
			var tmp_str *zend.ZendString
			var replace *zend.ZendString
			var tmp_replace *zend.ZendString
			var entry *zend.Zval
			for {
				var __ht *zend.HashTable = pats
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
						_z = zend.Z_INDIRECT_P(_z)
					}
					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					num_key = _p.h
					str_key = _p.key
					entry = _z
					tmp_str = nil
					if zend.UNEXPECTED(str_key == nil) {
						tmp_str = zend.ZendLongToStr(num_key)
						str_key = tmp_str
					}
					replace = zend.ZvalGetTmpString(entry, &tmp_replace)
					if zend.ZSTR_LEN(str_key) < 1 {
						zend.RETVAL_STR_COPY(str)
					} else if zend.ZSTR_LEN(str_key) == 1 {
						zend.RETVAL_STR(PhpCharToStrEx(str, zend.ZSTR_VAL(str_key)[0], zend.ZSTR_VAL(replace), zend.ZSTR_LEN(replace), 1, nil))
					} else {
						var dummy zend.ZendLong
						zend.RETVAL_STR(PhpStrToStrEx(str, zend.ZSTR_VAL(str_key), zend.ZSTR_LEN(str_key), zend.ZSTR_VAL(replace), zend.ZSTR_LEN(replace), &dummy))
					}
					zend.ZendTmpStringRelease(tmp_str)
					zend.ZendTmpStringRelease(tmp_replace)
					return
				}
				break
			}
		} else {
			PhpStrtrArray(return_value, str, pats)
		}
	} else {
		if zend.TryConvertToString(from) == 0 {
			return
		}
		zend.RETVAL_STR(PhpStrtrEx(str, zend.Z_STRVAL_P(from), to, cli.MIN(zend.Z_STRLEN_P(from), to_len)))
		return
	}
}
func ZifStrrev(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var s *byte
	var e *byte
	var p *byte
	var n *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
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
	n = zend.ZendStringAlloc(zend.ZSTR_LEN(str), 0)
	p = zend.ZSTR_VAL(n)
	s = zend.ZSTR_VAL(str)
	e = s + zend.ZSTR_LEN(str)
	e--
	for e >= s {
		*e--
		b.PostInc(&(*p)) = (*e) + 1
	}
	*p = '0'
	zend.RETVAL_NEW_STR(n)
}
func PhpSimilarStr(txt1 *byte, len1 int, txt2 *byte, len2 int, pos1 *int, pos2 *int, max *int, count *int) {
	var p *byte
	var q *byte
	var end1 *byte = (*byte)(txt1 + len1)
	var end2 *byte = (*byte)(txt2 + len2)
	var l int
	*max = 0
	*count = 0
	for p = (*byte)(txt1); p < end1; p++ {
		for q = (*byte)(txt2); q < end2; q++ {
			for l = 0; p+l < end1 && q+l < end2 && p[l] == q[l]; l++ {

			}
			if l > (*max) {
				*max = l
				*count += 1
				*pos1 = p - txt1
				*pos2 = q - txt2
			}
		}
	}
}
func PhpSimilarChar(txt1 *byte, len1 int, txt2 *byte, len2 int) int {
	var sum int
	var pos1 int = 0
	var pos2 int = 0
	var max int
	var count int
	PhpSimilarStr(txt1, len1, txt2, len2, &pos1, &pos2, &max, &count)
	if b.Assign(&sum, max) {
		if pos1 != 0 && pos2 != 0 && count > 1 {
			sum += PhpSimilarChar(txt1, pos1, txt2, pos2)
		}
		if pos1+max < len1 && pos2+max < len2 {
			sum += PhpSimilarChar(txt1+pos1+max, len1-pos1-max, txt2+pos2+max, len2-pos2-max)
		}
	}
	return sum
}
func ZifSimilarText(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var t1 *zend.ZendString
	var t2 *zend.ZendString
	var percent *zend.Zval = nil
	var ac int = zend.ZEND_NUM_ARGS()
	var sim int
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &t1, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &t2, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &percent, 0)
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
	if zend.ZSTR_LEN(t1)+zend.ZSTR_LEN(t2) == 0 {
		if ac > 2 {
			zend.ZEND_TRY_ASSIGN_REF_DOUBLE(percent, 0)
		}
		zend.RETVAL_LONG(0)
		return
	}
	sim = PhpSimilarChar(zend.ZSTR_VAL(t1), zend.ZSTR_LEN(t1), zend.ZSTR_VAL(t2), zend.ZSTR_LEN(t2))
	if ac > 2 {
		zend.ZEND_TRY_ASSIGN_REF_DOUBLE(percent, sim*200.0/(zend.ZSTR_LEN(t1)+zend.ZSTR_LEN(t2)))
	}
	zend.RETVAL_LONG(sim)
	return
}
func ZifAddcslashes(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var what *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &what, 0) == 0) {
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
	if zend.ZSTR_LEN(str) == 0 {
		zend.RETVAL_EMPTY_STRING()
		return
	}
	if zend.ZSTR_LEN(what) == 0 {
		zend.RETVAL_STR_COPY(str)
		return
	}
	zend.RETVAL_STR(PhpAddcslashesStr(zend.ZSTR_VAL(str), zend.ZSTR_LEN(str), zend.ZSTR_VAL(what), zend.ZSTR_LEN(what)))
	return
}
func ZifAddslashes(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
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
	if zend.ZSTR_LEN(str) == 0 {
		zend.RETVAL_EMPTY_STRING()
		return
	}
	zend.RETVAL_STR(PhpAddslashes(str))
	return
}
func ZifStripcslashes(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
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
	zend.ZVAL_STRINGL(return_value, zend.ZSTR_VAL(str), zend.ZSTR_LEN(str))
	PhpStripcslashes(zend.Z_STR_P(return_value))
}
func ZifStripslashes(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
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
	zend.ZVAL_STRINGL(return_value, zend.ZSTR_VAL(str), zend.ZSTR_LEN(str))
	PhpStripslashes(zend.Z_STR_P(return_value))
}
func PhpStripcslashes(str *zend.ZendString) {
	var source *byte
	var end *byte
	var target *byte
	var nlen int = zend.ZSTR_LEN(str)
	var i int
	var numtmp []byte
	source = (*byte)(zend.ZSTR_VAL(str))
	end = source + zend.ZSTR_LEN(str)
	target = zend.ZSTR_VAL(str)
	for ; source < end; source++ {
		if (*source) == '\\' && source+1 < end {
			source++
			switch *source {
			case 'n':
				b.PostInc(&(*target)) = '\n'
				nlen--
				break
			case 'r':
				b.PostInc(&(*target)) = '\r'
				nlen--
				break
			case 'a':
				b.PostInc(&(*target)) = 'a'
				nlen--
				break
			case 't':
				b.PostInc(&(*target)) = '\t'
				nlen--
				break
			case 'v':
				b.PostInc(&(*target)) = 'v'
				nlen--
				break
			case 'b':
				b.PostInc(&(*target)) = 'b'
				nlen--
				break
			case 'f':
				b.PostInc(&(*target)) = 'f'
				nlen--
				break
			case '\\':
				b.PostInc(&(*target)) = '\\'
				nlen--
				break
			case 'x':
				if source+1 < end && isxdigit(int(*(source + 1))) {
					numtmp[0] = *(b.PreInc(&source))
					if source+1 < end && isxdigit(int(*(source + 1))) {
						numtmp[1] = *(b.PreInc(&source))
						numtmp[2] = '0'
						nlen -= 3
					} else {
						numtmp[1] = '0'
						nlen -= 2
					}
					b.PostInc(&(*target)) = byte(strtol(numtmp, nil, 16))
					break
				}
			default:
				i = 0
				for source < end && (*source) >= '0' && (*source) <= '7' && i < 3 {
					*source++
					numtmp[b.PostInc(&i)] = (*source) - 1
				}
				if i != 0 {
					numtmp[i] = '0'
					b.PostInc(&(*target)) = byte(strtol(numtmp, nil, 8))
					nlen -= i
					source--
				} else {
					b.PostInc(&(*target)) = *source
					nlen--
				}
			}
		} else {
			b.PostInc(&(*target)) = *source
		}
	}
	if nlen != 0 {
		*target = '0'
	}
	zend.ZSTR_LEN(str) = nlen
}
func PhpAddcslashesStr(str *byte, len_ int, what *byte, wlength int) *zend.ZendString {
	var flags []byte
	var target *byte
	var source *byte
	var end *byte
	var c byte
	var newlen int
	var new_str *zend.ZendString = zend.ZendStringSafeAlloc(4, len_, 0, 0)
	PhpCharmask((*uint8)(what), wlength, flags)
	source = str
	end = source + len_
	target = zend.ZSTR_VAL(new_str)
	for ; source < end; source++ {
		c = *source
		if flags[uint8(c)] {
			if uint8(c < 32 || uint8(c > 126) != 0) != 0 {
				b.PostInc(&(*target)) = '\\'
				switch c {
				case '\n':
					b.PostInc(&(*target)) = 'n'
					break
				case '\t':
					b.PostInc(&(*target)) = 't'
					break
				case '\r':
					b.PostInc(&(*target)) = 'r'
					break
				case 'a':
					b.PostInc(&(*target)) = 'a'
					break
				case 'v':
					b.PostInc(&(*target)) = 'v'
					break
				case 'b':
					b.PostInc(&(*target)) = 'b'
					break
				case 'f':
					b.PostInc(&(*target)) = 'f'
					break
				default:
					target += sprintf(target, "%03o", uint8(c))
				}
				continue
			}
			b.PostInc(&(*target)) = '\\'
		}
		b.PostInc(&(*target)) = c
	}
	*target = 0
	newlen = target - zend.ZSTR_VAL(new_str)
	if newlen < len_*4 {
		new_str = zend.ZendStringTruncate(new_str, newlen, 0)
	}
	return new_str
}
func PhpAddcslashes(str *zend.ZendString, what string, wlength int) *zend.ZendString {
	return PhpAddcslashesStr(zend.ZSTR_VAL(str), zend.ZSTR_LEN(str), what, wlength)
}
func PhpAddslashes(str *zend.ZendString) *zend.ZendString {
	/* maximum string length, worst case situation */

	var target *byte
	var source *byte
	var end *byte
	var offset int
	var new_str *zend.ZendString
	if str == nil {
		return zend.ZSTR_EMPTY_ALLOC()
	}
	source = zend.ZSTR_VAL(str)
	end = source + zend.ZSTR_LEN(str)
	for source < end {
		switch *source {
		case '0':

		case '\'':

		case '"':

		case '\\':
			goto do_escape
		default:
			source++
			break
		}
	}
	return zend.ZendStringCopy(str)
do_escape:
	offset = source - (*byte)(zend.ZSTR_VAL(str))
	new_str = zend.ZendStringSafeAlloc(2, zend.ZSTR_LEN(str)-offset, offset, 0)
	memcpy(zend.ZSTR_VAL(new_str), zend.ZSTR_VAL(str), offset)
	target = zend.ZSTR_VAL(new_str) + offset
	for source < end {
		switch *source {
		case '0':
			b.PostInc(&(*target)) = '\\'
			b.PostInc(&(*target)) = '0'
			break
		case '\'':

		case '"':

		case '\\':
			b.PostInc(&(*target)) = '\\'
		default:
			b.PostInc(&(*target)) = *source
			break
		}
		source++
	}
	*target = '0'
	if zend.ZSTR_LEN(new_str)-(target-zend.ZSTR_VAL(new_str)) > 16 {
		new_str = zend.ZendStringTruncate(new_str, target-zend.ZSTR_VAL(new_str), 0)
	} else {
		zend.ZSTR_LEN(new_str) = target - zend.ZSTR_VAL(new_str)
	}
	return new_str
}
func PhpStripslashesImpl(str *byte, out *byte, len_ int) *byte {
	for len_ > 0 {
		if (*str) == '\\' {
			str++
			len_--
			if len_ > 0 {
				if (*str) == '0' {
					b.PostInc(&(*out)) = '0'
					str++
				} else {
					*str++
					b.PostInc(&(*out)) = (*str) - 1
				}
				len_--
			}
		} else {
			*str++
			b.PostInc(&(*out)) = (*str) - 1
			len_--
		}
	}
	return out
}
func PhpStripslashes(str *zend.ZendString) {
	var t *byte = PhpStripslashesImpl(zend.ZSTR_VAL(str), zend.ZSTR_VAL(str), zend.ZSTR_LEN(str))
	if t != zend.ZSTR_VAL(str)+zend.ZSTR_LEN(str) {
		zend.ZSTR_LEN(str) = t - zend.ZSTR_VAL(str)
		zend.ZSTR_VAL(str)[zend.ZSTR_LEN(str)] = '0'
	}
}
func Isheb(c __auto__) int {
	if uint8(c) >= 224 && uint8(c) <= 250 {
		return 1
	} else {
		return 0
	}
}
func _isblank(c __auto__) int {
	if uint8(c) == ' ' || uint8(c) == '\t' {
		return 1
	} else {
		return 0
	}
}
func _isnewline(c byte) int {
	if uint8(c) == '\n' || uint8(c) == '\r' {
		return 1
	} else {
		return 0
	}
}
func PhpStrReplaceInSubject(search *zend.Zval, replace *zend.Zval, subject *zend.Zval, result *zend.Zval, case_sensitivity int) zend.ZendLong {
	var search_entry *zend.Zval
	var tmp_result *zend.ZendString
	var tmp_subject_str *zend.ZendString
	var replace_value *byte = nil
	var replace_len int = 0
	var replace_count zend.ZendLong = 0
	var subject_str *zend.ZendString
	var lc_subject_str *zend.ZendString = nil
	var replace_idx uint32

	/* Make sure we're dealing with strings. */

	subject_str = zend.ZvalGetTmpString(subject, &tmp_subject_str)
	if zend.ZSTR_LEN(subject_str) == 0 {
		zend.ZendTmpStringRelease(tmp_subject_str)
		zend.ZVAL_EMPTY_STRING(result)
		return 0
	}

	/* If search is an array */

	if zend.Z_TYPE_P(search) == zend.IS_ARRAY {

		/* Duplicate subject string for repeated replacement */

		zend.ZendStringAddref(subject_str)
		if zend.Z_TYPE_P(replace) == zend.IS_ARRAY {
			replace_idx = 0
		} else {

			/* Set replacement value to the passed one */

			replace_value = zend.Z_STRVAL_P(replace)
			replace_len = zend.Z_STRLEN_P(replace)
		}

		/* For each entry in the search array, get the entry */

		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(search)
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
					_z = zend.Z_INDIRECT_P(_z)
				}
				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				search_entry = _z

				/* Make sure we're dealing with strings. */

				var tmp_search_str *zend.ZendString
				var search_str *zend.ZendString = zend.ZvalGetTmpString(search_entry, &tmp_search_str)
				var replace_entry_str *zend.ZendString
				var tmp_replace_entry_str *zend.ZendString = nil

				/* If replace is an array. */

				if zend.Z_TYPE_P(replace) == zend.IS_ARRAY {

					/* Get current entry */

					var replace_entry *zend.Zval = nil
					for replace_idx < zend.Z_ARRVAL_P(replace).nNumUsed {
						replace_entry = &zend.Z_ARRVAL_P(replace).arData[replace_idx].val
						if zend.Z_TYPE_P(replace_entry) != zend.IS_UNDEF {
							break
						}
						replace_idx++
					}
					if replace_idx < zend.Z_ARRVAL_P(replace).nNumUsed {

						/* Make sure we're dealing with strings. */

						replace_entry_str = zend.ZvalGetTmpString(replace_entry, &tmp_replace_entry_str)

						/* Set replacement value to the one we got from array */

						replace_value = zend.ZSTR_VAL(replace_entry_str)
						replace_len = zend.ZSTR_LEN(replace_entry_str)
						replace_idx++
					} else {

						/* We've run out of replacement strings, so use an empty one. */

						replace_value = ""
						replace_len = 0
					}
				}
				if zend.ZSTR_LEN(search_str) == 1 {
					var old_replace_count zend.ZendLong = replace_count
					tmp_result = PhpCharToStrEx(subject_str, zend.ZSTR_VAL(search_str)[0], replace_value, replace_len, case_sensitivity, &replace_count)
					if lc_subject_str != nil && replace_count != old_replace_count {
						zend.ZendStringReleaseEx(lc_subject_str, 0)
						lc_subject_str = nil
					}
				} else if zend.ZSTR_LEN(search_str) > 1 {
					if case_sensitivity != 0 {
						tmp_result = PhpStrToStrEx(subject_str, zend.ZSTR_VAL(search_str), zend.ZSTR_LEN(search_str), replace_value, replace_len, &replace_count)
					} else {
						var old_replace_count zend.ZendLong = replace_count
						if lc_subject_str == nil {
							lc_subject_str = PhpStringTolower(subject_str)
						}
						tmp_result = PhpStrToStrIEx(subject_str, zend.ZSTR_VAL(lc_subject_str), search_str, replace_value, replace_len, &replace_count)
						if replace_count != old_replace_count {
							zend.ZendStringReleaseEx(lc_subject_str, 0)
							lc_subject_str = nil
						}
					}
				} else {
					zend.ZendTmpStringRelease(tmp_search_str)
					zend.ZendTmpStringRelease(tmp_replace_entry_str)
					continue
				}
				zend.ZendTmpStringRelease(tmp_search_str)
				zend.ZendTmpStringRelease(tmp_replace_entry_str)
				if subject_str == tmp_result {
					zend.ZendStringDelref(subject_str)
				} else {
					zend.ZendStringReleaseEx(subject_str, 0)
					subject_str = tmp_result
					if zend.ZSTR_LEN(subject_str) == 0 {
						zend.ZendStringReleaseEx(subject_str, 0)
						zend.ZVAL_EMPTY_STRING(result)
						if lc_subject_str != nil {
							zend.ZendStringReleaseEx(lc_subject_str, 0)
						}
						zend.ZendTmpStringRelease(tmp_subject_str)
						return replace_count
					}
				}
			}
			break
		}
		zend.ZVAL_STR(result, subject_str)
		if lc_subject_str != nil {
			zend.ZendStringReleaseEx(lc_subject_str, 0)
		}
	} else {
		zend.ZEND_ASSERT(zend.Z_TYPE_P(search) == zend.IS_STRING)
		if zend.Z_STRLEN_P(search) == 1 {
			zend.ZVAL_STR(result, PhpCharToStrEx(subject_str, zend.Z_STRVAL_P(search)[0], zend.Z_STRVAL_P(replace), zend.Z_STRLEN_P(replace), case_sensitivity, &replace_count))
		} else if zend.Z_STRLEN_P(search) > 1 {
			if case_sensitivity != 0 {
				zend.ZVAL_STR(result, PhpStrToStrEx(subject_str, zend.Z_STRVAL_P(search), zend.Z_STRLEN_P(search), zend.Z_STRVAL_P(replace), zend.Z_STRLEN_P(replace), &replace_count))
			} else {
				lc_subject_str = PhpStringTolower(subject_str)
				zend.ZVAL_STR(result, PhpStrToStrIEx(subject_str, zend.ZSTR_VAL(lc_subject_str), zend.Z_STR_P(search), zend.Z_STRVAL_P(replace), zend.Z_STRLEN_P(replace), &replace_count))
				zend.ZendStringReleaseEx(lc_subject_str, 0)
			}
		} else {
			zend.ZVAL_STR_COPY(result, subject_str)
		}
	}
	zend.ZendTmpStringRelease(tmp_subject_str)
	return replace_count
}
func PhpStrReplaceCommon(execute_data *zend.ZendExecuteData, return_value *zend.Zval, case_sensitivity int) {
	var subject *zend.Zval
	var search *zend.Zval
	var replace *zend.Zval
	var subject_entry *zend.Zval
	var zcount *zend.Zval = nil
	var result zend.Zval
	var string_key *zend.ZendString
	var num_key zend.ZendUlong
	var count zend.ZendLong = 0
	var argc int = zend.ZEND_NUM_ARGS()
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 4
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
			zend.ZendParseArgZvalDeref(_arg, &search, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &replace, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &subject, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zcount, 0)
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

	/* Make sure we're dealing with strings and do the replacement. */

	if zend.Z_TYPE_P(search) != zend.IS_ARRAY {
		zend.ConvertToStringEx(search)
		if zend.Z_TYPE_P(replace) != zend.IS_STRING {
			zend.ConvertToStringEx(replace)
		}
	} else if zend.Z_TYPE_P(replace) != zend.IS_ARRAY {
		zend.ConvertToStringEx(replace)
	}
	if zend.ExecutorGlobals.exception != nil {
		return
	}

	/* if subject is an array */

	if zend.Z_TYPE_P(subject) == zend.IS_ARRAY {
		zend.ArrayInit(return_value)

		/* For each subject entry, convert it to string, then perform replacement
		   and add the result to the return_value array. */

		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(subject)
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
					_z = zend.Z_INDIRECT_P(_z)
				}
				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				num_key = _p.h
				string_key = _p.key
				subject_entry = _z
				zend.ZVAL_DEREF(subject_entry)
				if zend.Z_TYPE_P(subject_entry) != zend.IS_ARRAY && zend.Z_TYPE_P(subject_entry) != zend.IS_OBJECT {
					count += PhpStrReplaceInSubject(search, replace, subject_entry, &result, case_sensitivity)
				} else {
					zend.ZVAL_COPY(&result, subject_entry)
				}

				/* Add to return array */

				if string_key != nil {
					zend.ZendHashAddNew(zend.Z_ARRVAL_P(return_value), string_key, &result)
				} else {
					zend.ZendHashIndexAddNew(zend.Z_ARRVAL_P(return_value), num_key, &result)
				}

				/* Add to return array */

			}
			break
		}

		/* For each subject entry, convert it to string, then perform replacement
		   and add the result to the return_value array. */

	} else {
		count = PhpStrReplaceInSubject(search, replace, subject, return_value, case_sensitivity)
	}
	if argc > 3 {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zcount, count)
	}
}
func ZifStrReplace(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpStrReplaceCommon(execute_data, return_value, 1)
}
func ZifStrIreplace(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpStrReplaceCommon(execute_data, return_value, 0)
}
func PhpHebrev(execute_data *zend.ZendExecuteData, return_value *zend.Zval, convert_newlines int) {
	var str *byte
	var heb_str *byte
	var target *byte
	var tmp *byte
	var block_start int
	var block_end int
	var block_type int
	var block_length int
	var i int
	var max_chars zend.ZendLong = 0
	var char_count zend.ZendLong
	var begin int
	var end int
	var orig_begin int
	var str_len int
	var broken_str *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &max_chars, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if str_len == 0 {
		zend.RETVAL_FALSE
		return
	}
	tmp = str
	block_end = 0
	block_start = block_end
	heb_str = (*byte)(zend.Emalloc(str_len + 1))
	target = heb_str + str_len
	*target = 0
	target--
	block_length = 0
	if Isheb(*tmp) != 0 {
		block_type = _HEB_BLOCK_TYPE_HEB
	} else {
		block_type = _HEB_BLOCK_TYPE_ENG
	}
	for {
		if block_type == _HEB_BLOCK_TYPE_HEB {
			for (Isheb(int(*(tmp + 1))) != 0 || _isblank(int(*(tmp + 1))) != 0 || ispunct(int(*(tmp + 1))) || int((*(tmp + 1)) == '\n') != 0) && block_end < str_len-1 {
				tmp++
				block_end++
				block_length++
			}
			for i = block_start + 1; i <= block_end+1; i++ {
				*target = str[i-1]
				switch *target {
				case '(':
					*target = ')'
					break
				case ')':
					*target = '('
					break
				case '[':
					*target = ']'
					break
				case ']':
					*target = '['
					break
				case '{':
					*target = '}'
					break
				case '}':
					*target = '{'
					break
				case '<':
					*target = '>'
					break
				case '>':
					*target = '<'
					break
				case '\\':
					*target = '/'
					break
				case '/':
					*target = '\\'
					break
				default:
					break
				}
				target--
			}
			block_type = _HEB_BLOCK_TYPE_ENG
		} else {
			for Isheb(*(tmp + 1)) == 0 && int((*(tmp + 1)) != '\n' && block_end < str_len-1) != 0 {
				tmp++
				block_end++
				block_length++
			}
			for (_isblank(int(*tmp)) != 0 || ispunct(int(*tmp))) && (*tmp) != '/' && (*tmp) != '-' && block_end > block_start {
				tmp--
				block_end--
			}
			for i = block_end + 1; i >= block_start+1; i-- {
				*target = str[i-1]
				target--
			}
			block_type = _HEB_BLOCK_TYPE_HEB
		}
		block_start = block_end + 1
		if block_end >= str_len-1 {
			break
		}
	}
	broken_str = zend.ZendStringAlloc(str_len, 0)
	end = str_len - 1
	begin = end
	target = zend.ZSTR_VAL(broken_str)
	for true {
		char_count = 0
		for (max_chars == 0 || max_chars > 0 && char_count < max_chars) && begin > 0 {
			char_count++
			begin--
			if _isnewline(heb_str[begin]) != 0 {
				for begin > 0 && _isnewline(heb_str[begin-1]) != 0 {
					begin--
					char_count++
				}
				break
			}
		}
		if max_chars >= 0 && char_count == max_chars {
			var new_char_count int = char_count
			var new_begin int = begin
			for new_char_count > 0 {
				if _isblank(heb_str[new_begin]) != 0 || _isnewline(heb_str[new_begin]) != 0 {
					break
				}
				new_begin++
				new_char_count--
			}
			if new_char_count > 0 {
				begin = new_begin
			}
		}
		orig_begin = begin
		if _isblank(heb_str[begin]) != 0 {
			heb_str[begin] = '\n'
		}
		for begin <= end && _isnewline(heb_str[begin]) != 0 {
			begin++
		}
		for i = begin; i <= end; i++ {
			*target = heb_str[i]
			target++
		}
		for i = orig_begin; i <= end && _isnewline(heb_str[i]) != 0; i++ {
			*target = heb_str[i]
			target++
		}
		begin = orig_begin
		if begin == 0 {
			*target = 0
			break
		}
		begin--
		end = begin
	}
	zend.Efree(heb_str)
	if convert_newlines != 0 {
		zend.RETVAL_STR(PhpCharToStrEx(broken_str, '\n', "<br />\n", 7, 1, nil))
		zend.ZendStringReleaseEx(broken_str, 0)
	} else {
		zend.RETVAL_NEW_STR(broken_str)
		return
	}
}
func ZifHebrev(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpHebrev(execute_data, return_value, 0)
}
func ZifHebrevc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpHebrev(execute_data, return_value, 1)
}
func ZifNl2br(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	/* in brief this inserts <br /> or <br> before matched regexp \n\r?|\r\n? */

	var tmp *byte
	var end *byte
	var str *zend.ZendString
	var target *byte
	var repl_cnt int = 0
	var is_xhtml zend.ZendBool = 1
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &is_xhtml, &_dummy, 0) == 0) {
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
	tmp = zend.ZSTR_VAL(str)
	end = zend.ZSTR_VAL(str) + zend.ZSTR_LEN(str)

	/* it is really faster to scan twice and allocate mem once instead of scanning once
	   and constantly reallocing */

	for tmp < end {
		if (*tmp) == '\r' {
			if (*(tmp + 1)) == '\n' {
				tmp++
			}
			repl_cnt++
		} else if (*tmp) == '\n' {
			if (*(tmp + 1)) == '\r' {
				tmp++
			}
			repl_cnt++
		}
		tmp++
	}
	if repl_cnt == 0 {
		zend.RETVAL_STR_COPY(str)
		return
	}
	var repl_len int = b.CondF(is_xhtml != 0, func() int { return b.SizeOf("\"<br />\"") - 1 }, func() int { return b.SizeOf("\"<br>\"") - 1 })
	result = zend.ZendStringSafeAlloc(repl_cnt, repl_len, zend.ZSTR_LEN(str), 0)
	target = zend.ZSTR_VAL(result)
	tmp = zend.ZSTR_VAL(str)
	for tmp < end {
		switch *tmp {
		case '\r':

		case '\n':
			b.PostInc(&(*target)) = '<'
			b.PostInc(&(*target)) = 'b'
			b.PostInc(&(*target)) = 'r'
			if is_xhtml != 0 {
				b.PostInc(&(*target)) = ' '
				b.PostInc(&(*target)) = '/'
			}
			b.PostInc(&(*target)) = '>'
			if (*tmp) == '\r' && (*(tmp + 1)) == '\n' || (*tmp) == '\n' && (*(tmp + 1)) == '\r' {
				*tmp++
				b.PostInc(&(*target)) = (*tmp) - 1
			}
		default:
			b.PostInc(&(*target)) = *tmp
		}
		tmp++
	}
	*target = '0'
	zend.RETVAL_NEW_STR(result)
	return
}
func ZifStripTags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var buf *zend.ZendString
	var str *zend.ZendString
	var allow *zend.Zval = nil
	var allowed_tags *byte = nil
	var allowed_tags_len int = 0
	var tags_ss zend.SmartStr = zend.SmartStr{0}
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &allow, 0)
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
	if allow != nil {
		if zend.Z_TYPE_P(allow) == zend.IS_ARRAY {
			var tmp *zend.Zval
			var tag *zend.ZendString
			for {
				var __ht *zend.HashTable = zend.Z_ARRVAL_P(allow)
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					tmp = _z
					tag = zend.ZvalGetString(tmp)
					zend.SmartStrAppendc(&tags_ss, '<')
					zend.SmartStrAppend(&tags_ss, tag)
					zend.SmartStrAppendc(&tags_ss, '>')
					zend.ZendStringRelease(tag)
				}
				break
			}
			if tags_ss.s != nil {
				zend.SmartStr0(&tags_ss)
				allowed_tags = zend.ZSTR_VAL(tags_ss.s)
				allowed_tags_len = zend.ZSTR_LEN(tags_ss.s)
			}
		} else {

			/* To maintain a certain BC, we allow anything for the second parameter and return original string */

			zend.ConvertToString(allow)
			allowed_tags = zend.Z_STRVAL_P(allow)
			allowed_tags_len = zend.Z_STRLEN_P(allow)
		}
	}
	buf = zend.ZendStringInit(zend.ZSTR_VAL(str), zend.ZSTR_LEN(str), 0)
	zend.ZSTR_LEN(buf) = PhpStripTagsEx(zend.ZSTR_VAL(buf), zend.ZSTR_LEN(str), nil, allowed_tags, allowed_tags_len, 0)
	zend.SmartStrFree(&tags_ss)
	zend.RETVAL_NEW_STR(buf)
	return
}
func ZifSetlocale(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval = nil
	var plocale *zend.Zval
	var loc *zend.ZendString
	var retval *byte
	var cat zend.ZendLong
	var num_args int
	var i int = 0
	var idx uint32
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = -1
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
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &cat, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				args = _real_arg + 1
				num_args = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				num_args = 0
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
	idx = 0
	for true {
		if zend.Z_TYPE(args[0]) == zend.IS_ARRAY {
			for idx < zend.Z_ARRVAL(args[0]).nNumUsed {
				plocale = &zend.Z_ARRVAL(args[0]).arData[idx].val
				if zend.Z_TYPE_P(plocale) != zend.IS_UNDEF {
					break
				}
				idx++
			}
			if idx >= zend.Z_ARRVAL(args[0]).nNumUsed {
				break
			}
		} else {
			plocale = &args[i]
		}
		loc = zend.ZvalTryGetString(plocale)
		if zend.UNEXPECTED(loc == nil) {
			return
		}
		if !(strcmp("0", zend.ZSTR_VAL(loc))) {
			zend.ZendStringReleaseEx(loc, 0)
			loc = nil
		} else {
			if zend.ZSTR_LEN(loc) >= 255 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Specified locale name is too long")
				zend.ZendStringReleaseEx(loc, 0)
				break
			}
		}
		retval = PhpMySetlocale(cat, b.CondF1(loc != nil, func() []byte { return zend.ZSTR_VAL(loc) }, nil))
		if retval != nil {
			if loc != nil {

				/* Remember if locale was changed */

				var len_ int = strlen(retval)
				BG(locale_changed) = 1
				if cat == LC_CTYPE || cat == LC_ALL {
					if BG(locale_string) {
						zend.ZendStringReleaseEx(BG(locale_string), 0)
					}
					if len_ == zend.ZSTR_LEN(loc) && !(memcmp(zend.ZSTR_VAL(loc), retval, len_)) {
						BG(locale_string) = zend.ZendStringCopy(loc)
						zend.RETVAL_STR(BG(locale_string))
						return
					} else {
						BG(locale_string) = zend.ZendStringInit(retval, len_, 0)
						zend.ZendStringReleaseEx(loc, 0)
						zend.RETVAL_STR_COPY(BG(locale_string))
						return
					}
				} else if len_ == zend.ZSTR_LEN(loc) && !(memcmp(zend.ZSTR_VAL(loc), retval, len_)) {
					zend.RETVAL_STR(loc)
					return
				}
				zend.ZendStringReleaseEx(loc, 0)
			}
			zend.RETVAL_STRING(retval)
			return
		}
		if loc != nil {
			zend.ZendStringReleaseEx(loc, 0)
		}
		if zend.Z_TYPE(args[0]) == zend.IS_ARRAY {
			idx++
		} else {
			if b.PreInc(&i) >= num_args {
				break
			}
		}
	}
	zend.RETVAL_FALSE
	return
}
func ZifParseStr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *byte
	var arrayArg *zend.Zval = nil
	var res *byte = nil
	var arglen int
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
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &arg, &arglen, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &arrayArg, 0)
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
	res = zend.Estrndup(arg, arglen)
	if arrayArg == nil {
		var tmp zend.Zval
		var symbol_table *zend.ZendArray
		if zend.ZendForbidDynamicCall("parse_str() with a single argument") == zend.FAILURE {
			zend.Efree(res)
			return
		}
		core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Calling parse_str() without the result argument is deprecated")
		symbol_table = zend.ZendRebuildSymbolTable()
		zend.ZVAL_ARR(&tmp, symbol_table)
		core.sapi_module.treat_data(core.PARSE_STRING, res, &tmp)
		if zend.UNEXPECTED(zend.ZendHashDel(symbol_table, zend.ZSTR_KNOWN(zend.ZEND_STR_THIS)) == zend.SUCCESS) {
			zend.ZendThrowError(nil, "Cannot re-assign $this")
		}
	} else {
		arrayArg = zend.ZendTryArrayInit(arrayArg)
		if arrayArg == nil {
			zend.Efree(res)
			return
		}
		core.sapi_module.treat_data(core.PARSE_STRING, res, arrayArg)
	}
}
func PhpTagFind(tag *byte, len_ int, set *byte) int {
	var c byte
	var n *byte
	var t *byte
	var state int = 0
	var done int = 0
	var norm *byte
	if len_ == 0 {
		return 0
	}
	norm = zend.Emalloc(len_ + 1)
	n = norm
	t = tag
	c = tolower(*t)

	/*
	   normalize the tag removing leading and trailing whitespace
	   and turn any <a whatever...> into just <a> and any </tag>
	   into <tag>
	*/

	for done == 0 {
		switch c {
		case '<':
			*(b.PostInc(&n)) = c
			break
		case '>':
			done = 1
			break
		default:
			if !(isspace(int(c))) {
				if state == 0 {
					state = 1
				}
				if c != '/' || (*(t - 1)) != '<' && (*(t + 1)) != '>' {
					*(b.PostInc(&n)) = c
				}
			} else {
				if state == 1 {
					done = 1
				}
			}
			break
		}
		c = tolower(*(b.PreInc(&t)))
	}
	*(b.PostInc(&n)) = '>'
	*n = '0'
	if strstr(set, norm) {
		done = 1
	} else {
		done = 0
	}
	zend.Efree(norm)
	return done
}
func PhpStripTags(rbuf *byte, len_ int, stateptr *uint8, allow *byte, allow_len int) int {
	return PhpStripTagsEx(rbuf, len_, stateptr, allow, allow_len, 0)
}
func PhpStripTagsEx(rbuf *byte, len_ int, stateptr *uint8, allow *byte, allow_len int, allow_tag_spaces zend.ZendBool) int {
	var tbuf *byte
	var tp *byte
	var rp *byte
	var c byte
	var lc byte
	var buf *byte
	var p *byte
	var end *byte
	var br int
	var depth int = 0
	var in_q int = 0
	var state uint8 = 0
	var pos int
	var allow_free *byte = nil
	var is_xml byte = 0
	buf = zend.Estrndup(rbuf, len_)
	end = buf + len_
	lc = '0'
	p = buf
	rp = rbuf
	br = 0
	if allow != nil {
		allow_free = zend.ZendStrTolowerDupEx(allow, allow_len)
		if allow_free != nil {
			allow = allow_free
		} else {
			allow = allow
		}
		tbuf = zend.Emalloc(PHP_TAG_BUF_SIZE + 1)
		tp = tbuf
	} else {
		tp = nil
		tbuf = tp
	}
	if stateptr != nil {
		state = *stateptr
		switch state {
		case 1:
			goto state_1
		case 2:
			goto state_2
		case 3:
			goto state_3
		case 4:
			goto state_4
		default:
			break
		}
	}
state_0:
	if p >= end {
		goto finish
	}
	c = *p
	switch c {
	case '0':
		break
	case '<':
		if in_q != 0 {
			break
		}
		if isspace(*(p + 1)) && allow_tag_spaces == 0 {
			*(b.PostInc(&rp)) = c
			break
		}
		lc = '<'
		state = 1
		if allow != nil {
			if tp-tbuf >= PHP_TAG_BUF_SIZE {
				pos = tp - tbuf
				tbuf = zend.Erealloc(tbuf, tp-tbuf+PHP_TAG_BUF_SIZE+1)
				tp = tbuf + pos
			}
			*(b.PostInc(&tp)) = '<'
		}
		p++
		goto state_1
	case '>':
		if depth != 0 {
			depth--
			break
		}
		if in_q != 0 {
			break
		}
		*(b.PostInc(&rp)) = c
		break
	default:
		*(b.PostInc(&rp)) = c
		break
	}
	p++
	goto state_0
state_1:
	if p >= end {
		goto finish
	}
	c = *p
	switch c {
	case '0':
		break
	case '<':
		if in_q != 0 {
			break
		}
		if isspace(*(p + 1)) && allow_tag_spaces == 0 {
			goto reg_char_1
		}
		depth++
		break
	case '>':
		if depth != 0 {
			depth--
			break
		}
		if in_q != 0 {
			break
		}
		lc = '>'
		if is_xml && p >= buf+1 && (*(p - 1)) == '-' {
			break
		}
		is_xml = 0
		state = is_xml
		in_q = state
		if allow != nil {
			if tp-tbuf >= PHP_TAG_BUF_SIZE {
				pos = tp - tbuf
				tbuf = zend.Erealloc(tbuf, tp-tbuf+PHP_TAG_BUF_SIZE+1)
				tp = tbuf + pos
			}
			*(b.PostInc(&tp)) = '>'
			*tp = '0'
			if PhpTagFind(tbuf, tp-tbuf, allow) != 0 {
				memcpy(rp, tbuf, tp-tbuf)
				rp += tp - tbuf
			}
			tp = tbuf
		}
		p++
		goto state_0
	case '"':

	case '\'':
		if p != buf && (in_q == 0 || (*p) == in_q) {
			if in_q != 0 {
				in_q = 0
			} else {
				in_q = *p
			}
		}
		goto reg_char_1
	case '!':

		/* JavaScript & Other HTML scripting languages */

		if p >= buf+1 && (*(p - 1)) == '<' {
			state = 3
			lc = c
			p++
			goto state_3
		} else {
			goto reg_char_1
		}
		break
	case '?':
		if p >= buf+1 && (*(p - 1)) == '<' {
			br = 0
			state = 2
			p++
			goto state_2
		} else {
			goto reg_char_1
		}
		break
	default:
	reg_char_1:
		if allow != nil {
			if tp-tbuf >= PHP_TAG_BUF_SIZE {
				pos = tp - tbuf
				tbuf = zend.Erealloc(tbuf, tp-tbuf+PHP_TAG_BUF_SIZE+1)
				tp = tbuf + pos
			}
			*(b.PostInc(&tp)) = c
		}
		break
	}
	p++
	goto state_1
state_2:
	if p >= end {
		goto finish
	}
	c = *p
	switch c {
	case '(':
		if lc != '"' && lc != '\'' {
			lc = '('
			br++
		}
		break
	case ')':
		if lc != '"' && lc != '\'' {
			lc = ')'
			br--
		}
		break
	case '>':
		if depth != 0 {
			depth--
			break
		}
		if in_q != 0 {
			break
		}
		if br == 0 && p >= buf+1 && lc != '"' && (*(p - 1)) == '?' {
			state = 0
			in_q = state
			tp = tbuf
			p++
			goto state_0
		}
		break
	case '"':

	case '\'':
		if p >= buf+1 && (*(p - 1)) != '\\' {
			if lc == c {
				lc = '0'
			} else if lc != '\\' {
				lc = c
			}
			if p != buf && (in_q == 0 || (*p) == in_q) {
				if in_q != 0 {
					in_q = 0
				} else {
					in_q = *p
				}
			}
		}
		break
	case 'l':

	case 'L':

		/* swm: If we encounter '<?xml' then we shouldn't be in
		 * state == 2 (PHP). Switch back to HTML.
		 */

		if state == 2 && p > buf+4 && ((*(p - 1)) == 'm' || (*(p - 1)) == 'M') && ((*(p - 2)) == 'x' || (*(p - 2)) == 'X') && (*(p - 3)) == '?' && (*(p - 4)) == '<' {
			state = 1
			is_xml = 1
			p++
			goto state_1
		}
		break
	default:
		break
	}
	p++
	goto state_2
state_3:
	if p >= end {
		goto finish
	}
	c = *p
	switch c {
	case '>':
		if depth != 0 {
			depth--
			break
		}
		if in_q != 0 {
			break
		}
		state = 0
		in_q = state
		tp = tbuf
		p++
		goto state_0
	case '"':

	case '\'':
		if p != buf && (*(p - 1)) != '\\' && (in_q == 0 || (*p) == in_q) {
			if in_q != 0 {
				in_q = 0
			} else {
				in_q = *p
			}
		}
		break
	case '-':
		if p >= buf+2 && (*(p - 1)) == '-' && (*(p - 2)) == '!' {
			state = 4
			p++
			goto state_4
		}
		break
	case 'E':

	case 'e':

		/* !DOCTYPE exception */

		if p > buf+6 && ((*(p - 1)) == 'p' || (*(p - 1)) == 'P') && ((*(p - 2)) == 'y' || (*(p - 2)) == 'Y') && ((*(p - 3)) == 't' || (*(p - 3)) == 'T') && ((*(p - 4)) == 'c' || (*(p - 4)) == 'C') && ((*(p - 5)) == 'o' || (*(p - 5)) == 'O') && ((*(p - 6)) == 'd' || (*(p - 6)) == 'D') {
			state = 1
			p++
			goto state_1
		}
		break
	default:
		break
	}
	p++
	goto state_3
state_4:
	for p < end {
		c = *p
		if c == '>' && in_q == 0 {
			if p >= buf+2 && (*(p - 1)) == '-' && (*(p - 2)) == '-' {
				state = 0
				in_q = state
				tp = tbuf
				p++
				goto state_0
			}
		}
		p++
	}
finish:
	if rp < rbuf+len_ {
		*rp = '0'
	}
	zend.Efree(any(buf))
	if tbuf != nil {
		zend.Efree(tbuf)
	}
	if allow_free != nil {
		zend.Efree(allow_free)
	}
	if stateptr != nil {
		*stateptr = state
	}
	return size_t(rp - rbuf)
}
func ZifStrGetcsv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var delim byte = ','
	var enc byte = '"'
	var esc int = uint8('\\')
	var delim_str *byte = nil
	var enc_str *byte = nil
	var esc_str *byte = nil
	var delim_len int = 0
	var enc_len int = 0
	var esc_len int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 4
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &delim_str, &delim_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &enc_str, &enc_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &esc_str, &esc_len, 0) == 0) {
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
	if delim_len != 0 {
		delim = delim_str[0]
	} else {
		delim = delim
	}
	if enc_len != 0 {
		enc = enc_str[0]
	} else {
		enc = enc
	}
	if esc_str != nil {
		if esc_len != 0 {
			esc = uint8(esc_str[0])
		} else {
			esc = PHP_CSV_NO_ESCAPE
		}
	}
	PhpFgetcsv(nil, delim, enc, esc, zend.ZSTR_LEN(str), zend.ZSTR_VAL(str), return_value)
}
func ZifStrRepeat(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input_str *zend.ZendString
	var mult zend.ZendLong
	var result *zend.ZendString
	var result_len int
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &input_str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &mult, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if mult < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Second argument has to be greater than or equal to 0")
		return
	}

	/* Don't waste our time if it's empty */

	if zend.ZSTR_LEN(input_str) == 0 || mult == 0 {
		zend.RETVAL_EMPTY_STRING()
		return
	}

	/* Initialize the result string */

	result = zend.ZendStringSafeAlloc(zend.ZSTR_LEN(input_str), mult, 0, 0)
	result_len = zend.ZSTR_LEN(input_str) * mult

	/* Heavy optimization for situations where input string is 1 byte long */

	if zend.ZSTR_LEN(input_str) == 1 {
		memset(zend.ZSTR_VAL(result), (*zend.ZSTR_VAL)(input_str), mult)
	} else {
		var s *byte
		var ee *byte
		var e *byte
		var l ptrdiff_t = 0
		memcpy(zend.ZSTR_VAL(result), zend.ZSTR_VAL(input_str), zend.ZSTR_LEN(input_str))
		s = zend.ZSTR_VAL(result)
		e = zend.ZSTR_VAL(result) + zend.ZSTR_LEN(input_str)
		ee = zend.ZSTR_VAL(result) + result_len
		for e < ee {
			if e-s < ee-e {
				l = e - s
			} else {
				l = ee - e
			}
			memmove(e, s, l)
			e += l
		}
	}
	zend.ZSTR_VAL(result)[result_len] = '0'
	zend.RETVAL_NEW_STR(result)
	return
}
func ZifCountChars(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.ZendString
	var chars []int
	var mymode zend.ZendLong = 0
	var buf *uint8
	var inx int
	var retstr []byte
	var retlen int = 0
	var tmp int = 0
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &input, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &mymode, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if mymode < 0 || mymode > 4 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unknown mode")
		zend.RETVAL_FALSE
		return
	}
	buf = (*uint8)(zend.ZSTR_VAL(input))
	memset(any(chars), 0, b.SizeOf("chars"))
	for tmp < zend.ZSTR_LEN(input) {
		chars[*buf]++
		buf++
		tmp++
	}
	if mymode < 3 {
		zend.ArrayInit(return_value)
	}
	for inx = 0; inx < 256; inx++ {
		switch mymode {
		case 0:
			zend.AddIndexLong(return_value, inx, chars[inx])
			break
		case 1:
			if chars[inx] != 0 {
				zend.AddIndexLong(return_value, inx, chars[inx])
			}
			break
		case 2:
			if chars[inx] == 0 {
				zend.AddIndexLong(return_value, inx, chars[inx])
			}
			break
		case 3:
			if chars[inx] != 0 {
				retstr[b.PostInc(&retlen)] = inx
			}
			break
		case 4:
			if chars[inx] == 0 {
				retstr[b.PostInc(&retlen)] = inx
			}
			break
		}
	}
	if mymode >= 3 && mymode <= 4 {
		zend.RETVAL_STRINGL(retstr, retlen)
		return
	}
}
func PhpStrnatcmp(execute_data *zend.ZendExecuteData, return_value *zend.Zval, fold_case int) {
	var s1 *zend.ZendString
	var s2 *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &s1, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &s2, 0) == 0) {
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
	zend.RETVAL_LONG(StrnatcmpEx(zend.ZSTR_VAL(s1), zend.ZSTR_LEN(s1), zend.ZSTR_VAL(s2), zend.ZSTR_LEN(s2), fold_case))
	return
}
func StringNaturalCompareFunctionEx(result *zend.Zval, op1 *zend.Zval, op2 *zend.Zval, case_insensitive zend.ZendBool) int {
	var tmp_str1 *zend.ZendString
	var tmp_str2 *zend.ZendString
	var str1 *zend.ZendString = zend.ZvalGetTmpString(op1, &tmp_str1)
	var str2 *zend.ZendString = zend.ZvalGetTmpString(op2, &tmp_str2)
	zend.ZVAL_LONG(result, StrnatcmpEx(zend.ZSTR_VAL(str1), zend.ZSTR_LEN(str1), zend.ZSTR_VAL(str2), zend.ZSTR_LEN(str2), case_insensitive))
	zend.ZendTmpStringRelease(tmp_str1)
	zend.ZendTmpStringRelease(tmp_str2)
	return zend.SUCCESS
}
func StringNaturalCaseCompareFunction(result *zend.Zval, op1 *zend.Zval, op2 *zend.Zval) int {
	return StringNaturalCompareFunctionEx(result, op1, op2, 1)
}
func StringNaturalCompareFunction(result *zend.Zval, op1 *zend.Zval, op2 *zend.Zval) int {
	return StringNaturalCompareFunctionEx(result, op1, op2, 0)
}
func ZifStrnatcmp(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpStrnatcmp(execute_data, return_value, 0)
}
func ZifLocaleconv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var grouping zend.Zval
	var mon_grouping zend.Zval
	var len_ int
	var i int

	/* We don't need no stinkin' parameters... */

	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ArrayInit(return_value)
	zend.ArrayInit(&grouping)
	zend.ArrayInit(&mon_grouping)
	var currlocdata __struct__lconv
	LocaleconvR(&currlocdata)

	/* Grab the grouping data out of the array */

	len_ = int(strlen(currlocdata.grouping))
	for i = 0; i < len_; i++ {
		zend.AddIndexLong(&grouping, i, currlocdata.grouping[i])
	}

	/* Grab the monetary grouping data out of the array */

	len_ = int(strlen(currlocdata.mon_grouping))
	for i = 0; i < len_; i++ {
		zend.AddIndexLong(&mon_grouping, i, currlocdata.mon_grouping[i])
	}
	zend.AddAssocString(return_value, "decimal_point", currlocdata.decimal_point)
	zend.AddAssocString(return_value, "thousands_sep", currlocdata.thousands_sep)
	zend.AddAssocString(return_value, "int_curr_symbol", currlocdata.int_curr_symbol)
	zend.AddAssocString(return_value, "currency_symbol", currlocdata.currency_symbol)
	zend.AddAssocString(return_value, "mon_decimal_point", currlocdata.mon_decimal_point)
	zend.AddAssocString(return_value, "mon_thousands_sep", currlocdata.mon_thousands_sep)
	zend.AddAssocString(return_value, "positive_sign", currlocdata.positive_sign)
	zend.AddAssocString(return_value, "negative_sign", currlocdata.negative_sign)
	zend.AddAssocLong(return_value, "int_frac_digits", currlocdata.int_frac_digits)
	zend.AddAssocLong(return_value, "frac_digits", currlocdata.frac_digits)
	zend.AddAssocLong(return_value, "p_cs_precedes", currlocdata.p_cs_precedes)
	zend.AddAssocLong(return_value, "p_sep_by_space", currlocdata.p_sep_by_space)
	zend.AddAssocLong(return_value, "n_cs_precedes", currlocdata.n_cs_precedes)
	zend.AddAssocLong(return_value, "n_sep_by_space", currlocdata.n_sep_by_space)
	zend.AddAssocLong(return_value, "p_sign_posn", currlocdata.p_sign_posn)
	zend.AddAssocLong(return_value, "n_sign_posn", currlocdata.n_sign_posn)
	zend.ZendHashStrUpdate(zend.Z_ARRVAL_P(return_value), "grouping", b.SizeOf("\"grouping\"")-1, &grouping)
	zend.ZendHashStrUpdate(zend.Z_ARRVAL_P(return_value), "mon_grouping", b.SizeOf("\"mon_grouping\"")-1, &mon_grouping)
}
func ZifStrnatcasecmp(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpStrnatcmp(execute_data, return_value, 1)
}
func ZifSubstrCount(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var haystack *byte
	var needle *byte
	var offset zend.ZendLong = 0
	var length zend.ZendLong = 0
	var ac int = zend.ZEND_NUM_ARGS()
	var count zend.ZendLong = 0
	var haystack_len int
	var needle_len int
	var p *byte
	var endp *byte
	var cmp byte
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 4
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
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &haystack, &haystack_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &needle, &needle_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &length, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if needle_len == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Empty substring")
		zend.RETVAL_FALSE
		return
	}
	p = haystack
	endp = p + haystack_len
	if offset < 0 {
		offset += zend.ZendLong(haystack_len)
	}
	if offset < 0 || int(offset > haystack_len) != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Offset not contained in string")
		zend.RETVAL_FALSE
		return
	}
	p += offset
	if ac == 4 {
		if length < 0 {
			length += haystack_len - offset
		}
		if length < 0 || int(length > haystack_len-offset) != 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid length value")
			zend.RETVAL_FALSE
			return
		}
		endp = p + length
	}
	if needle_len == 1 {
		cmp = needle[0]
		for b.Assign(&p, memchr(p, cmp, endp-p)) {
			count++
			p++
		}
	} else {
		for b.Assign(&p, (*byte)(core.PhpMemnstr(p, needle, needle_len, endp))) {
			p += needle_len
			count++
		}
	}
	zend.RETVAL_LONG(count)
	return
}
func ZifStrPad(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	/* Input arguments */

	var input *zend.ZendString
	var pad_length zend.ZendLong

	/* Helper variables */

	var num_pad_chars int
	var pad_str *byte = " "
	var pad_str_len int = 1
	var pad_type_val zend.ZendLong = STR_PAD_RIGHT
	var i int
	var left_pad int = 0
	var right_pad int = 0
	var result *zend.ZendString = nil
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 4
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &input, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &pad_length, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &pad_str, &pad_str_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &pad_type_val, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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

	/* If resulting string turns out to be shorter than input string,
	   we simply copy the input and return. */

	if pad_length < 0 || int(pad_length <= zend.ZSTR_LEN(input)) != 0 {
		zend.RETVAL_STR_COPY(input)
		return
	}
	if pad_str_len == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Padding string cannot be empty")
		return
	}
	if pad_type_val < STR_PAD_LEFT || pad_type_val > STR_PAD_BOTH {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Padding type has to be STR_PAD_LEFT, STR_PAD_RIGHT, or STR_PAD_BOTH")
		return
	}
	num_pad_chars = pad_length - zend.ZSTR_LEN(input)
	if num_pad_chars >= core.INT_MAX {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Padding length is too long")
		return
	}
	result = zend.ZendStringSafeAlloc(1, zend.ZSTR_LEN(input), num_pad_chars, 0)
	zend.ZSTR_LEN(result) = 0

	/* We need to figure out the left/right padding lengths. */

	switch pad_type_val {
	case STR_PAD_RIGHT:
		left_pad = 0
		right_pad = num_pad_chars
		break
	case STR_PAD_LEFT:
		left_pad = num_pad_chars
		right_pad = 0
		break
	case STR_PAD_BOTH:
		left_pad = num_pad_chars / 2
		right_pad = num_pad_chars - left_pad
		break
	}

	/* First we pad on the left. */

	for i = 0; i < left_pad; i++ {
		zend.ZSTR_VAL(result)[b.PostInc(&(zend.ZSTR_LEN(result)))] = pad_str[i%pad_str_len]
	}

	/* Then we copy the input string. */

	memcpy(zend.ZSTR_VAL(result)+zend.ZSTR_LEN(result), zend.ZSTR_VAL(input), zend.ZSTR_LEN(input))
	zend.ZSTR_LEN(result) += zend.ZSTR_LEN(input)

	/* Finally, we pad on the right. */

	for i = 0; i < right_pad; i++ {
		zend.ZSTR_VAL(result)[b.PostInc(&(zend.ZSTR_LEN(result)))] = pad_str[i%pad_str_len]
	}
	zend.ZSTR_VAL(result)[zend.ZSTR_LEN(result)] = '0'
	zend.RETVAL_NEW_STR(result)
	return
}
func ZifSscanf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval = nil
	var str *byte
	var format *byte
	var str_len int
	var format_len int
	var result int
	var num_args int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = -1
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
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &format, &format_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				args = _real_arg + 1
				num_args = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				num_args = 0
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
	result = PhpSscanfInternal(str, format, num_args, args, 0, return_value)
	if SCAN_ERROR_WRONG_PARAM_COUNT == result {
		zend.WRONG_PARAM_COUNT
	}
}
func PhpStrRot13(str *zend.ZendString) *zend.ZendString {
	var ret *zend.ZendString
	var p *byte
	var e *byte
	var target *byte
	if zend.UNEXPECTED(zend.ZSTR_LEN(str) == 0) {
		return zend.ZSTR_EMPTY_ALLOC()
	}
	ret = zend.ZendStringAlloc(zend.ZSTR_LEN(str), 0)
	p = zend.ZSTR_VAL(str)
	e = p + zend.ZSTR_LEN(str)
	target = zend.ZSTR_VAL(ret)
	for p < e {
		if (*p) >= 'a' && (*p) <= 'z' {
			b.PostInc(&(*target)) = 'a' + (b.PostInc(&(*p))-'a'+13)%26
		} else if (*p) >= 'A' && (*p) <= 'Z' {
			b.PostInc(&(*target)) = 'A' + (b.PostInc(&(*p))-'A'+13)%26
		} else {
			*p++
			b.PostInc(&(*target)) = (*p) - 1
		}
	}
	*target = '0'
	return ret
}
func ZifStrRot13(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &arg, 0) == 0) {
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
	zend.RETVAL_STR(PhpStrRot13(arg))
	return
}
func PhpStringShuffle(str *byte, len_ zend.ZendLong) {
	var n_elems zend.ZendLong
	var rnd_idx zend.ZendLong
	var n_left zend.ZendLong
	var temp byte

	/* The implementation is stolen from array_data_shuffle       */

	n_elems = len_
	if n_elems <= 1 {
		return
	}
	n_left = n_elems
	for b.PreDec(&n_left) {
		rnd_idx = PhpMtRandRange(0, n_left)
		if rnd_idx != n_left {
			temp = str[n_left]
			str[n_left] = str[rnd_idx]
			str[rnd_idx] = temp
		}
	}
}
func ZifStrShuffle(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *zend.ZendString
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &arg, 0) == 0) {
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
	zend.RETVAL_STRINGL(zend.ZSTR_VAL(arg), zend.ZSTR_LEN(arg))
	if zend.Z_STRLEN_P(return_value) > 1 {
		PhpStringShuffle(zend.Z_STRVAL_P(return_value), zend.ZendLong(zend.Z_STRLEN_P(return_value)))
	}
}
func ZifStrWordCount(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var char_list *byte = nil
	var ch []*byte
	var p *byte
	var e *byte
	var s *byte
	var char_list_len int = 0
	var word_count int = 0
	var type_ zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &type_, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &char_list, &char_list_len, 0) == 0) {
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
	switch type_ {
	case 1:

	case 2:
		zend.ArrayInit(return_value)
		if zend.ZSTR_LEN(str) == 0 {
			return
		}
		break
	case 0:
		if zend.ZSTR_LEN(str) == 0 {
			zend.RETVAL_LONG(0)
			return
		}

		/* nothing to be done */

		break
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid format value "+zend.ZEND_LONG_FMT, type_)
		zend.RETVAL_FALSE
		return
	}
	if char_list != nil {
		PhpCharmask((*uint8)(char_list), char_list_len, ch)
	}
	p = zend.ZSTR_VAL(str)
	e = zend.ZSTR_VAL(str) + zend.ZSTR_LEN(str)

	/* first character cannot be ' or -, unless explicitly allowed by the user */

	if (*p) == '\'' && (char_list == nil || ch['\''] == nil) || (*p) == '-' && (char_list == nil || ch['-'] == nil) {
		p++
	}

	/* last character cannot be -, unless explicitly allowed by the user */

	if (*(e - 1)) == '-' && (char_list == nil || ch['-'] == nil) {
		e--
	}
	for p < e {
		s = p
		for p < e && (isalpha(uint8(*p)) || char_list != nil && ch[uint8(*p)] != nil || (*p) == '\'' || (*p) == '-') {
			p++
		}
		if p > s {
			switch type_ {
			case 1:
				zend.AddNextIndexStringl(return_value, s, p-s)
				break
			case 2:
				zend.AddIndexStringl(return_value, s-zend.ZSTR_VAL(str), s, p-s)
				break
			default:
				word_count++
				break
			}
		}
		p++
	}
	if type_ == 0 {
		zend.RETVAL_LONG(word_count)
		return
	}
}
func ZifMoneyFormat(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var format_len int = 0
	var format *byte
	var p *byte
	var e *byte
	var value float64
	var check zend.ZendBool = 0
	var str *zend.ZendString
	var res_len ssize_t
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &format, &format_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgDouble(_arg, &value, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_DOUBLE
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
	p = format
	e = p + format_len
	for b.Assign(&p, memchr(p, '%', e-p)) {
		if (*(p + 1)) == '%' {
			p += 2
		} else if check == 0 {
			check = 1
			p++
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Only a single %%i or %%n token can be used")
			zend.RETVAL_FALSE
			return
		}
	}
	str = zend.ZendStringSafeAlloc(format_len, 1, 1024, 0)
	if b.Assign(&res_len, strfmon(zend.ZSTR_VAL(str), zend.ZSTR_LEN(str), format, value)) < 0 {
		zend.ZendStringEfree(str)
		zend.RETVAL_FALSE
		return
	}
	zend.ZSTR_LEN(str) = int(res_len)
	zend.ZSTR_VAL(str)[zend.ZSTR_LEN(str)] = '0'
	zend.RETVAL_NEW_STR(zend.ZendStringTruncate(str, zend.ZSTR_LEN(str), 0))
	return
}
func ZifStrSplit(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var split_length zend.ZendLong = 1
	var p *byte
	var n_reg_segments int
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &str, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &split_length, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if split_length <= 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The length of each segment must be greater than zero")
		zend.RETVAL_FALSE
		return
	}
	if 0 == zend.ZSTR_LEN(str) || int(split_length >= zend.ZSTR_LEN(str)) != 0 {
		zend.ArrayInitSize(return_value, 1)
		zend.AddNextIndexStringl(return_value, zend.ZSTR_VAL(str), zend.ZSTR_LEN(str))
		return
	}
	zend.ArrayInitSize(return_value, uint32_t((zend.ZSTR_LEN(str)-1)/split_length+1))
	n_reg_segments = zend.ZSTR_LEN(str) / split_length
	p = zend.ZSTR_VAL(str)
	for b.PostDec(&n_reg_segments) > 0 {
		zend.AddNextIndexStringl(return_value, p, split_length)
		p += split_length
	}
	if p != zend.ZSTR_VAL(str)+zend.ZSTR_LEN(str) {
		zend.AddNextIndexStringl(return_value, p, zend.ZSTR_VAL(str)+zend.ZSTR_LEN(str)-p)
	}
}
func ZifStrpbrk(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var haystack *zend.ZendString
	var char_list *zend.ZendString
	var haystack_ptr *byte
	var cl_ptr *byte
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &haystack, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &char_list, 0) == 0) {
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
	if zend.ZSTR_LEN(char_list) == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The character list cannot be empty")
		zend.RETVAL_FALSE
		return
	}
	for haystack_ptr = zend.ZSTR_VAL(haystack); haystack_ptr < zend.ZSTR_VAL(haystack)+zend.ZSTR_LEN(haystack); haystack_ptr++ {
		for cl_ptr = zend.ZSTR_VAL(char_list); cl_ptr < zend.ZSTR_VAL(char_list)+zend.ZSTR_LEN(char_list); cl_ptr++ {
			if (*cl_ptr) == (*haystack_ptr) {
				zend.RETVAL_STRINGL(haystack_ptr, zend.ZSTR_VAL(haystack)+zend.ZSTR_LEN(haystack)-haystack_ptr)
				return
			}
		}
	}
	zend.RETVAL_FALSE
	return
}
func ZifSubstrCompare(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var s1 *zend.ZendString
	var s2 *zend.ZendString
	var offset zend.ZendLong
	var len_ zend.ZendLong = 0
	var len_is_default zend.ZendBool = 1
	var cs zend.ZendBool = 0
	var cmp_len int
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 5
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
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &s1, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &s2, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &len_, &len_is_default, 1, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &cs, &_dummy, 0) == 0) {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if len_is_default == 0 && len_ <= 0 {
		if len_ == 0 {
			zend.RETVAL_LONG(0)
			return
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "The length must be greater than or equal to zero")
			zend.RETVAL_FALSE
			return
		}
	}
	if offset < 0 {
		offset = zend.ZSTR_LEN(s1) + offset
		if offset < 0 {
			offset = 0
		} else {
			offset = offset
		}
	}
	if int(offset > zend.ZSTR_LEN(s1)) != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The start position cannot exceed initial string length")
		zend.RETVAL_FALSE
		return
	}
	if len_ != 0 {
		cmp_len = int(len_)
	} else {
		cmp_len = zend.MAX(zend.ZSTR_LEN(s2), zend.ZSTR_LEN(s1)-offset)
	}
	if cs == 0 {
		zend.RETVAL_LONG(zend.ZendBinaryStrncmp(zend.ZSTR_VAL(s1)+offset, zend.ZSTR_LEN(s1)-offset, zend.ZSTR_VAL(s2), zend.ZSTR_LEN(s2), cmp_len))
		return
	} else {
		zend.RETVAL_LONG(zend.ZendBinaryStrncasecmpL(zend.ZSTR_VAL(s1)+offset, zend.ZSTR_LEN(s1)-offset, zend.ZSTR_VAL(s2), zend.ZSTR_LEN(s2), cmp_len))
		return
	}
}
func PhpUtf8Encode(s *byte, len_ int) *zend.ZendString {
	var pos int = len_
	var str *zend.ZendString
	var c uint8
	str = zend.ZendStringSafeAlloc(len_, 2, 0, 0)
	zend.ZSTR_LEN(str) = 0
	for pos > 0 {

		/* The lower 256 codepoints of Unicode are identical to Latin-1,
		 * so we don't need to do any mapping here. */

		c = uint8(*s)
		if c < 0x80 {
			zend.ZSTR_VAL(str)[b.PostInc(&(zend.ZSTR_LEN(str)))] = byte(c)
		} else {
			zend.ZSTR_VAL(str)[b.PostInc(&(zend.ZSTR_LEN(str)))] = 0xc0 | c>>6
			zend.ZSTR_VAL(str)[b.PostInc(&(zend.ZSTR_LEN(str)))] = 0x80 | c&0x3f
		}
		pos--
		s++
	}
	zend.ZSTR_VAL(str)[zend.ZSTR_LEN(str)] = '0'
	str = zend.ZendStringTruncate(str, zend.ZSTR_LEN(str), 0)
	return str
}
func PhpUtf8Decode(s *byte, len_ int) *zend.ZendString {
	var pos int = 0
	var c uint
	var str *zend.ZendString
	str = zend.ZendStringAlloc(len_, 0)
	zend.ZSTR_LEN(str) = 0
	for pos < len_ {
		var status int = zend.FAILURE
		c = PhpNextUtf8Char((*uint8)(s), int(len_), &pos, &status)

		/* The lower 256 codepoints of Unicode are identical to Latin-1,
		 * so we don't need to do any mapping here beyond replacing non-Latin-1
		 * characters. */

		if status == zend.FAILURE || c > 0xff {
			c = '?'
		}
		zend.ZSTR_VAL(str)[b.PostInc(&(zend.ZSTR_LEN(str)))] = c
	}
	zend.ZSTR_VAL(str)[zend.ZSTR_LEN(str)] = '0'
	if zend.ZSTR_LEN(str) < len_ {
		str = zend.ZendStringTruncate(str, zend.ZSTR_LEN(str), 0)
	}
	return str
}
func ZifUtf8Encode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *byte
	var arg_len int
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
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &arg, &arg_len, 0) == 0) {
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
	zend.RETVAL_STR(PhpUtf8Encode(arg, arg_len))
	return
}
func ZifUtf8Decode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *byte
	var arg_len int
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
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &arg, &arg_len, 0) == 0) {
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
	zend.RETVAL_STR(PhpUtf8Decode(arg, arg_len))
	return
}
