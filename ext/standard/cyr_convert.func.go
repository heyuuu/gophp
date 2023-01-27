// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
)

func PhpConvertCyrString(str *uint8, length int, from byte, to byte) *byte {
	var from_table *uint8
	var to_table *uint8
	var tmp uint8
	var i int
	from_table = nil
	to_table = nil
	switch toupper(int(uint8(from))) {
	case 'W':
		from_table = _cyrWin1251
		break
	case 'A':

	case 'D':
		from_table = _cyrCp866
		break
	case 'I':
		from_table = _cyrIso88595
		break
	case 'M':
		from_table = _cyrMac
		break
	case 'K':
		break
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unknown source charset: %c", from)
		break
	}
	switch toupper(int(uint8(to))) {
	case 'W':
		to_table = _cyrWin1251
		break
	case 'A':

	case 'D':
		to_table = _cyrCp866
		break
	case 'I':
		to_table = _cyrIso88595
		break
	case 'M':
		to_table = _cyrMac
		break
	case 'K':
		break
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unknown destination charset: %c", to)
		break
	}
	if str == nil {
		return (*byte)(str)
	}
	for i = 0; i < length; i++ {
		if from_table == nil {
			tmp = str[i]
		} else {
			tmp = from_table[str[i]]
		}
		if to_table == nil {
			str[i] = tmp
		} else {
			str[i] = to_table[tmp+256]
		}
	}
	return (*byte)(str)
}
func ZifConvertCyrString(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *byte
	var fr_cs *byte
	var to_cs *byte
	var input_len int
	var fr_cs_len int
	var to_cs_len int
	var str *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 3
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
			if zend.ZendParseArgString(_arg, &input, &input_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &fr_cs, &fr_cs_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &to_cs, &to_cs_len, 0) == 0 {
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
	str = zend.ZendStringInit(input, input_len, 0)
	PhpConvertCyrString((*uint8)(str.GetVal()), str.GetLen(), fr_cs[0], to_cs[0])
	zend.RETVAL_NEW_STR(str)
}
