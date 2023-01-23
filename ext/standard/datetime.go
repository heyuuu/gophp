// <<generate>>

package standard

import (
	"sik/core"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/datetime.h>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// #define DATETIME_H

// Source: <ext/standard/datetime.c>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Rasmus Lerdorf <rasmus@php.net>                             |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "zend_operators.h"

// # include "datetime.h"

// # include "php_globals.h"

// # include < time . h >

// # include < sys / time . h >

// # include < stdio . h >

var MonShortNames []*byte = []*byte{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
var DayShortNames []*byte = []*byte{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

/* {{{ PHPAPI char *php_std_date(time_t t)
   Return date string in standard format for http headers */

func PhpStdDate(t int64) *byte {
	var tm1 *__struct__tm
	var tmbuf __struct__tm
	var str *byte
	tm1 = gmtime_r(&t, &tmbuf)
	str = zend._emalloc(81)
	str[0] = '0'
	if tm1 == nil {
		return str
	}
	core.ApPhpSnprintf(str, 80, "%s, %02d %s %04d %02d:%02d:%02d GMT", DayShortNames[tm1.tm_wday], tm1.tm_mday, MonShortNames[tm1.tm_mon], tm1.tm_year+1900, tm1.tm_hour, tm1.tm_min, tm1.tm_sec)
	str[79] = 0
	return str
}

/* }}} */

/* {{{ proto string strptime(string timestamp, string format)
   Parse a time/date generated with strftime() */

func ZifStrptime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
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
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &ts, &ts_length, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &format, &format_length, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	memset(&parsed_time, 0, g.SizeOf("parsed_time"))
	unparsed_part = strptime(ts, format, &parsed_time)
	if unparsed_part == nil {
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	zend.AddAssocLongEx(return_value, "tm_sec", strlen("tm_sec"), parsed_time.tm_sec)
	zend.AddAssocLongEx(return_value, "tm_min", strlen("tm_min"), parsed_time.tm_min)
	zend.AddAssocLongEx(return_value, "tm_hour", strlen("tm_hour"), parsed_time.tm_hour)
	zend.AddAssocLongEx(return_value, "tm_mday", strlen("tm_mday"), parsed_time.tm_mday)
	zend.AddAssocLongEx(return_value, "tm_mon", strlen("tm_mon"), parsed_time.tm_mon)
	zend.AddAssocLongEx(return_value, "tm_year", strlen("tm_year"), parsed_time.tm_year)
	zend.AddAssocLongEx(return_value, "tm_wday", strlen("tm_wday"), parsed_time.tm_wday)
	zend.AddAssocLongEx(return_value, "tm_yday", strlen("tm_yday"), parsed_time.tm_yday)
	zend.AddAssocStringEx(return_value, "unparsed", strlen("unparsed"), unparsed_part)
}

/* }}} */
