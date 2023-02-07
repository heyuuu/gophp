// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
)

func ZifUniqid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var prefix *byte = ""
	var more_entropy zend.ZendBool = 0
	var uniqid *zend.ZendString
	var sec int
	var usec int
	var prefix_len int = 0
	var tv __struct__timeval
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &prefix, &prefix_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &more_entropy, &_dummy, 0) == 0 {
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

	/* This implementation needs current microsecond to change,
	 * hence we poll time until it does. This is much faster than
	 * calling usleep(1) which may cause the kernel to schedule
	 * another process, causing a pause of around 10ms.
	 */

	for {
		void(gettimeofday((*__struct__timeval)(&tv), (*__struct__timezone)(nil)))
		if !(tv.tv_sec == PrevTv.tv_sec && tv.tv_usec == PrevTv.tv_usec) {
			break
		}
	}
	PrevTv.tv_sec = tv.tv_sec
	PrevTv.tv_usec = tv.tv_usec
	sec = int(tv.tv_sec)
	usec = int(tv.tv_usec % 0x100000)

	/* The max value usec can have is 0xF423F, so we use only five hex
	 * digits for usecs.
	 */

	if more_entropy != 0 {
		uniqid = core.Strpprintf(0, "%s%08x%05x%.8F", prefix, sec, usec, PhpCombinedLcg()*10)
	} else {
		uniqid = core.Strpprintf(0, "%s%08x%05x", prefix, sec, usec)
	}
	return_value.SetString(uniqid)
	return
}
