// <<generate>>

package standard

import (
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/microtime.h>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Author: Paul Panotzki - Bunyip Information Systems                   |
   +----------------------------------------------------------------------+
*/

// #define MICROTIME_H

// Source: <ext/standard/microtime.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Author: Paul Panotzki - Bunyip Information Systems                   |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include < sys / types . h >

// # include < sys / time . h >

// # include < sys / resource . h >

// # include < unistd . h >

// # include < stdlib . h >

// # include < string . h >

// # include < stdio . h >

// # include < errno . h >

// # include "microtime.h"

// failed # include "ext/date/php_date.h"

// #define NUL       '\0'

// #define MICRO_IN_SEC       1000000.00

// #define SEC_IN_MIN       60

func _phpGettimeofday(execute_data *zend.ZendExecuteData, return_value *zend.Zval, mode int) {
	var get_as_float zend.ZendBool = 0
	var tp __struct__timeval = __struct__timeval{0}
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &get_as_float, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
		assert(false)
	}
	if get_as_float != 0 {
		var __z *zend.Zval = return_value
		__z.value.dval = float64(tp.tv_sec + tp.tv_usec/1000000.0)
		__z.u1.type_info = 5
		return
	}
	if mode != 0 {
		var offset *timelib_time_offset
		offset = timelib_get_time_zone_info(tp.tv_sec, get_timezone_info())
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		zend.AddAssocLongEx(return_value, "sec", strlen("sec"), tp.tv_sec)
		zend.AddAssocLongEx(return_value, "usec", strlen("usec"), tp.tv_usec)
		zend.AddAssocLongEx(return_value, "minuteswest", strlen("minuteswest"), -(offset.offset)/60)
		zend.AddAssocLongEx(return_value, "dsttime", strlen("dsttime"), offset.is_dst)
		timelib_time_offset_dtor(offset)
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStrpprintf(0, "%.8F %ld", tp.tv_usec/1000000.0, long(tp.tv_sec))
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
}

/* {{{ proto mixed microtime([bool get_as_float])
   Returns either a string or a float containing the current time in seconds and microseconds */

func ZifMicrotime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	_phpGettimeofday(execute_data, return_value, 0)
}

/* }}} */

func ZifGettimeofday(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	_phpGettimeofday(execute_data, return_value, 1)
}

/* }}} */

/* {{{ proto array getrusage([int who])
   Returns an array of usage statistics */

func ZifGetrusage(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var usg __struct__rusage
	var pwho zend.ZendLong = 0
	var who int = RUSAGE_SELF
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &pwho, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	memset(&usg, 0, g.SizeOf("struct rusage"))
	if getrusage(who, &usg) == -1 {
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

	// #define PHP_RUSAGE_PARA(a) add_assoc_long ( return_value , # a , usg . a )

	zend.AddAssocLongEx(return_value, "ru_oublock", strlen("ru_oublock"), usg.ru_oublock)
	zend.AddAssocLongEx(return_value, "ru_inblock", strlen("ru_inblock"), usg.ru_inblock)
	zend.AddAssocLongEx(return_value, "ru_msgsnd", strlen("ru_msgsnd"), usg.ru_msgsnd)
	zend.AddAssocLongEx(return_value, "ru_msgrcv", strlen("ru_msgrcv"), usg.ru_msgrcv)
	zend.AddAssocLongEx(return_value, "ru_maxrss", strlen("ru_maxrss"), usg.ru_maxrss)
	zend.AddAssocLongEx(return_value, "ru_ixrss", strlen("ru_ixrss"), usg.ru_ixrss)
	zend.AddAssocLongEx(return_value, "ru_idrss", strlen("ru_idrss"), usg.ru_idrss)
	zend.AddAssocLongEx(return_value, "ru_minflt", strlen("ru_minflt"), usg.ru_minflt)
	zend.AddAssocLongEx(return_value, "ru_majflt", strlen("ru_majflt"), usg.ru_majflt)
	zend.AddAssocLongEx(return_value, "ru_nsignals", strlen("ru_nsignals"), usg.ru_nsignals)
	zend.AddAssocLongEx(return_value, "ru_nvcsw", strlen("ru_nvcsw"), usg.ru_nvcsw)
	zend.AddAssocLongEx(return_value, "ru_nivcsw", strlen("ru_nivcsw"), usg.ru_nivcsw)
	zend.AddAssocLongEx(return_value, "ru_nswap", strlen("ru_nswap"), usg.ru_nswap)
	zend.AddAssocLongEx(return_value, "ru_utime . tv_usec", strlen("ru_utime . tv_usec"), usg.ru_utime.tv_usec)
	zend.AddAssocLongEx(return_value, "ru_utime . tv_sec", strlen("ru_utime . tv_sec"), usg.ru_utime.tv_sec)
	zend.AddAssocLongEx(return_value, "ru_stime . tv_usec", strlen("ru_stime . tv_usec"), usg.ru_stime.tv_usec)
	zend.AddAssocLongEx(return_value, "ru_stime . tv_sec", strlen("ru_stime . tv_sec"), usg.ru_stime.tv_sec)
}

/* }}} */
