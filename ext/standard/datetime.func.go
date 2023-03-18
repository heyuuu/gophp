// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/types"
)

func PhpStdDate(t int64) *byte {
	var tm1 *__struct__tm
	var tmbuf __struct__tm
	var str *byte
	tm1 = core.PhpGmtimeR(&t, &tmbuf)
	str = zend.Emalloc(81)
	str[0] = '0'
	if tm1 == nil {
		return str
	}
	core.Snprintf(str, 80, "%s, %02d %s %04d %02d:%02d:%02d GMT", DayShortNames[tm1.tm_wday], tm1.tm_mday, MonShortNames[tm1.tm_mon], tm1.tm_year+1900, tm1.tm_hour, tm1.tm_min, tm1.tm_sec)
	str[79] = 0
	return str
}
func ZifStrptime(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var ts *byte
	var ts_length int
	var format *byte
	var format_length int
	var parsed_time __struct__tm
	var unparsed_part *byte
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
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
			if zend.ZendParseArgString(_arg, &ts, &ts_length, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &format, &format_length, 0) == 0 {
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
	memset(&parsed_time, 0, b.SizeOf("parsed_time"))
	unparsed_part = strptime(ts, format, &parsed_time)
	if unparsed_part == nil {
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	zend.AddAssocLong(return_value, "tm_sec", parsed_time.tm_sec)
	zend.AddAssocLong(return_value, "tm_min", parsed_time.tm_min)
	zend.AddAssocLong(return_value, "tm_hour", parsed_time.tm_hour)
	zend.AddAssocLong(return_value, "tm_mday", parsed_time.tm_mday)
	zend.AddAssocLong(return_value, "tm_mon", parsed_time.tm_mon)
	zend.AddAssocLong(return_value, "tm_year", parsed_time.tm_year)
	zend.AddAssocLong(return_value, "tm_wday", parsed_time.tm_wday)
	zend.AddAssocLong(return_value, "tm_yday", parsed_time.tm_yday)
	zend.AddAssocString(return_value, "unparsed", unparsed_part)
}
