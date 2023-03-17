// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
)

func ZifSoundex(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte
	var i int
	var _small int
	var str_len int
	var code int
	var last int
	var soundex []byte
	var soundex_table []byte = []byte{0, '1', '2', '3', 0, '1', '2', 0, 0, '2', '2', '4', '5', '5', 0, '1', '2', '6', '2', '3', 0, '1', 0, '2', 0, '2'}
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
	if str_len == 0 {
		return_value.SetFalse()
		return
	}

	/* build soundex string */

	last = -1
	i = 0
	_small = 0
	for ; i < str_len && _small < 4; i++ {

		/* convert chars to upper case and strip non-letter chars */

		code = toupper(int(uint8(str[i])))
		if code >= 'A' && code <= 'Z' {
			if _small == 0 {

				/* remember first valid char */

				soundex[b.PostInc(&_small)] = byte(code)
				last = soundex_table[code-'A']
			} else {

				/* ignore sequences of consonants with same soundex */

				code = soundex_table[code-'A']
				if code != last {
					if code != 0 {
						soundex[b.PostInc(&_small)] = byte(code)
					}
					last = code
				}
			}
		}
	}

	/* pad with '0' and terminate with 0 ;-) */

	for _small < 4 {
		soundex[b.PostInc(&_small)] = '0'
	}
	soundex[_small] = '0'
	return_value.SetRawString(b.CastStr(soundex, _small))
	return
}
