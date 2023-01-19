// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/syslog.c>

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
   | Author: Stig Sæther Bakken <ssb@php.net>                             |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_ini.h"

// # include "zend_globals.h"

// # include < stdlib . h >

// # include < unistd . h >

// # include < string . h >

// # include < errno . h >

// # include < stdio . h >

// # include "basic_functions.h"

// # include "php_ext_syslog.h"

/* {{{ PHP_MINIT_FUNCTION
 */

func ZmStartupSyslog(type_ int, module_number int) int {
	/* error levels */

	zend.ZendRegisterLongConstant("LOG_EMERG", g.SizeOf("\"LOG_EMERG\"")-1, LOG_EMERG, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_ALERT", g.SizeOf("\"LOG_ALERT\"")-1, LOG_ALERT, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_CRIT", g.SizeOf("\"LOG_CRIT\"")-1, LOG_CRIT, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_ERR", g.SizeOf("\"LOG_ERR\"")-1, LOG_ERR, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_WARNING", g.SizeOf("\"LOG_WARNING\"")-1, LOG_WARNING, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_NOTICE", g.SizeOf("\"LOG_NOTICE\"")-1, LOG_NOTICE, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_INFO", g.SizeOf("\"LOG_INFO\"")-1, LOG_INFO, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_DEBUG", g.SizeOf("\"LOG_DEBUG\"")-1, LOG_DEBUG, 1<<0|1<<1, module_number)

	/* facility: type of program logging the message */

	zend.ZendRegisterLongConstant("LOG_KERN", g.SizeOf("\"LOG_KERN\"")-1, LOG_KERN, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_USER", g.SizeOf("\"LOG_USER\"")-1, LOG_USER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_MAIL", g.SizeOf("\"LOG_MAIL\"")-1, LOG_MAIL, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_DAEMON", g.SizeOf("\"LOG_DAEMON\"")-1, LOG_DAEMON, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_AUTH", g.SizeOf("\"LOG_AUTH\"")-1, LOG_AUTH, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_SYSLOG", g.SizeOf("\"LOG_SYSLOG\"")-1, LOG_SYSLOG, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_LPR", g.SizeOf("\"LOG_LPR\"")-1, LOG_LPR, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_LOCAL0", g.SizeOf("\"LOG_LOCAL0\"")-1, LOG_LOCAL0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_LOCAL1", g.SizeOf("\"LOG_LOCAL1\"")-1, LOG_LOCAL1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_LOCAL2", g.SizeOf("\"LOG_LOCAL2\"")-1, LOG_LOCAL2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_LOCAL3", g.SizeOf("\"LOG_LOCAL3\"")-1, LOG_LOCAL3, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_LOCAL4", g.SizeOf("\"LOG_LOCAL4\"")-1, LOG_LOCAL4, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_LOCAL5", g.SizeOf("\"LOG_LOCAL5\"")-1, LOG_LOCAL5, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_LOCAL6", g.SizeOf("\"LOG_LOCAL6\"")-1, LOG_LOCAL6, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_LOCAL7", g.SizeOf("\"LOG_LOCAL7\"")-1, LOG_LOCAL7, 1<<0|1<<1, module_number)

	/* options */

	zend.ZendRegisterLongConstant("LOG_PID", g.SizeOf("\"LOG_PID\"")-1, LOG_PID, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_CONS", g.SizeOf("\"LOG_CONS\"")-1, LOG_CONS, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_ODELAY", g.SizeOf("\"LOG_ODELAY\"")-1, LOG_ODELAY, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOG_NDELAY", g.SizeOf("\"LOG_NDELAY\"")-1, LOG_NDELAY, 1<<0|1<<1, module_number)
	BasicGlobals.SetSyslogDevice(nil)
	return zend.SUCCESS
}

/* }}} */

func ZmActivateSyslog(type_ int, module_number int) int {
	BasicGlobals.SetSyslogDevice(nil)
	return zend.SUCCESS
}
func ZmShutdownSyslog(type_ int, module_number int) int {
	if BasicGlobals.GetSyslogDevice() != nil {
		zend.Free(BasicGlobals.GetSyslogDevice())
		BasicGlobals.SetSyslogDevice(nil)
	}
	return zend.SUCCESS
}
func PhpOpenlog(ident *byte, option int, facility int) {
	openlog(ident, option, facility)
	core.CoreGlobals.have_called_openlog = 1
}

/* {{{ proto bool openlog(string ident, int option, int facility)
   Open connection to system logger */

func ZifOpenlog(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var ident *byte
	var option zend.ZendLong
	var facility zend.ZendLong
	var ident_len int
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &ident, &ident_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgLong(_arg, &option, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgLong(_arg, &facility, &_dummy, 0, 0) == 0 {
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
	if BasicGlobals.GetSyslogDevice() != nil {
		zend.Free(BasicGlobals.GetSyslogDevice())
	}
	BasicGlobals.SetSyslogDevice(zend.ZendStrndup(ident, ident_len))
	if BasicGlobals.GetSyslogDevice() == nil {
		return_value.u1.type_info = 2
		return
	}
	PhpOpenlog(BasicGlobals.GetSyslogDevice(), option, facility)
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifCloselog(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	closelog()
	if BasicGlobals.GetSyslogDevice() != nil {
		zend.Free(BasicGlobals.GetSyslogDevice())
		BasicGlobals.SetSyslogDevice(nil)
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifSyslog(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var priority zend.ZendLong
	var message *byte
	var message_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &priority, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgString(_arg, &message, &message_len, 0) == 0 {
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
	core.PhpSyslog(priority, "%s", message)
	return_value.u1.type_info = 3
	return
}

/* }}} */
