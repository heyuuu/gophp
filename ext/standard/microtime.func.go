// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
)

func _phpGettimeofday(executeData *zend.ZendExecuteData, return_value *zend.Zval, mode int) {
	var get_as_float zend.ZendBool = 0
	var tp __struct__timeval = __struct__timeval{0}
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
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
			if zend.ZendParseArgBool(_arg, &get_as_float, &_dummy, 0) == 0 {
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
	if gettimeofday(&tp, nil) {
		zend.ZEND_ASSERT(false)
	}
	if get_as_float != 0 {
		return_value.SetDouble(float64(tp.tv_sec + tp.tv_usec/MICRO_IN_SEC))
		return
	}
	if mode != 0 {
		var offset *timelib_time_offset
		offset = timelib_get_time_zone_info(tp.tv_sec, get_timezone_info())
		zend.ArrayInit(return_value)
		zend.AddAssocLong(return_value, "sec", tp.tv_sec)
		zend.AddAssocLong(return_value, "usec", tp.tv_usec)
		zend.AddAssocLong(return_value, "minuteswest", -(offset.offset)/SEC_IN_MIN)
		zend.AddAssocLong(return_value, "dsttime", offset.is_dst)
		timelib_time_offset_dtor(offset)
	} else {
		return_value.SetString(zend.ZendStrpprintf(0, "%.8F %ld", tp.tv_usec/MICRO_IN_SEC, long(tp.tv_sec)))
		return
	}
}
func ZifMicrotime(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	_phpGettimeofday(executeData, return_value, 0)
}
func ZifGettimeofday(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	_phpGettimeofday(executeData, return_value, 1)
}
func ZifGetrusage(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var usg __struct__rusage
	var pwho zend.ZendLong = 0
	var who int = RUSAGE_SELF
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
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
			if zend.ZendParseArgLong(_arg, &pwho, &_dummy, 0, 0) == 0 {
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
	if pwho == 1 {
		who = RUSAGE_CHILDREN
	}
	memset(&usg, 0, b.SizeOf("struct rusage"))
	if getrusage(who, &usg) == -1 {
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)

	// #define PHP_RUSAGE_PARA(a) add_assoc_long ( return_value , # a , usg . a )

	zend.AddAssocLong(return_value, "ru_oublock", usg.ru_oublock)
	zend.AddAssocLong(return_value, "ru_inblock", usg.ru_inblock)
	zend.AddAssocLong(return_value, "ru_msgsnd", usg.ru_msgsnd)
	zend.AddAssocLong(return_value, "ru_msgrcv", usg.ru_msgrcv)
	zend.AddAssocLong(return_value, "ru_maxrss", usg.ru_maxrss)
	zend.AddAssocLong(return_value, "ru_ixrss", usg.ru_ixrss)
	zend.AddAssocLong(return_value, "ru_idrss", usg.ru_idrss)
	zend.AddAssocLong(return_value, "ru_minflt", usg.ru_minflt)
	zend.AddAssocLong(return_value, "ru_majflt", usg.ru_majflt)
	zend.AddAssocLong(return_value, "ru_nsignals", usg.ru_nsignals)
	zend.AddAssocLong(return_value, "ru_nvcsw", usg.ru_nvcsw)
	zend.AddAssocLong(return_value, "ru_nivcsw", usg.ru_nivcsw)
	zend.AddAssocLong(return_value, "ru_nswap", usg.ru_nswap)
	zend.AddAssocLong(return_value, "ru_utime . tv_usec", usg.ru_utime.tv_usec)
	zend.AddAssocLong(return_value, "ru_utime . tv_sec", usg.ru_utime.tv_sec)
	zend.AddAssocLong(return_value, "ru_stime . tv_usec", usg.ru_stime.tv_usec)
	zend.AddAssocLong(return_value, "ru_stime . tv_sec", usg.ru_stime.tv_sec)
}
