// <<generate>>

package standard

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/random.c>

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
   | Authors: Sammy Kaye Powers <me@sammyk.me>                            |
   +----------------------------------------------------------------------+
*/

// # include < stdlib . h >

// # include < sys / stat . h >

// # include < fcntl . h >

// # include < math . h >

// # include "php.h"

// # include "zend_exceptions.h"

// # include "php_random.h"

func RandomGlobalsCtor(random_globals_p *PhpRandomGlobals) { random_globals_p.SetFd(-1) }
func RandomGlobalsDtor(random_globals_p *PhpRandomGlobals) {
	if random_globals_p.GetFd() > 0 {
		close(random_globals_p.GetFd())
		random_globals_p.SetFd(-1)
	}
}

/* {{{ */

func ZmStartupRandom(type_ int, module_number int) int {
	RandomGlobalsCtor(&RandomGlobals)
	return zend.SUCCESS
}

/* }}} */

func ZmShutdownRandom(type_ int, module_number int) int {
	RandomGlobalsDtor(&RandomGlobals)
	return zend.SUCCESS
}

/* }}} */

func PhpRandomBytes(bytes any, size int, should_throw zend.ZendBool) int {
	var read_bytes int = 0
	var n ssize_t
	if read_bytes < size {
		var fd int = RandomGlobals.GetFd()
		var st __struct__stat
		if fd < 0 {
			fd = open("/dev/urandom", O_RDONLY)
			if fd < 0 {
				if should_throw != 0 {
					zend.ZendThrowException(zend.ZendCeException, "Cannot open source device", 0)
				}
				return zend.FAILURE
			}

			/* Does the file exist and is it a character device? */

			if fstat(fd, &st) != 0 || !(S_ISCHR(st.st_mode)) {
				close(fd)
				if should_throw != 0 {
					zend.ZendThrowException(zend.ZendCeException, "Error reading from source device", 0)
				}
				return zend.FAILURE
			}
			RandomGlobals.SetFd(fd)
		}
		for read_bytes = 0; read_bytes < size; read_bytes += int(n) {
			n = read(fd, bytes+read_bytes, size-read_bytes)
			if n <= 0 {
				break
			}
		}
		if read_bytes < size {
			if should_throw != 0 {
				zend.ZendThrowException(zend.ZendCeException, "Could not gather sufficient random data", 0)
			}
			return zend.FAILURE
		}
	}
	return zend.SUCCESS
}

/* }}} */

func ZifRandomBytes(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var size zend.ZendLong
	var bytes *zend.ZendString
	for {
		var _flags int = 1 << 2
		var _min_num_args int = 1
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

			if zend.ZendParseArgLong(_arg, &size, &_dummy, 0, 0) == 0 {
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
	if size < 1 {
		zend.ZendThrowException(zend.ZendCeError, "Length must be greater than 0", 0)
		return
	}
	bytes = zend.ZendStringAlloc(size, 0)
	if PhpRandomBytes(bytes.val, size, 1) == zend.FAILURE {
		zend.ZendStringReleaseEx(bytes, 0)
		return
	}
	bytes.val[size] = '0'
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = bytes
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func PhpRandomInt(min zend.ZendLong, max zend.ZendLong, result *zend.ZendLong, should_throw zend.ZendBool) int {
	var umax zend.ZendUlong
	var trial zend.ZendUlong
	if min == max {
		*result = min
		return zend.SUCCESS
	}
	umax = zend.ZendUlong(max - zend.ZendUlong(min))
	if PhpRandomBytes(&trial, g.SizeOf("trial"), should_throw) == zend.FAILURE {
		return zend.FAILURE
	}

	/* Special case where no modulus is required */

	if umax == UINT64_MAX {
		*result = zend.ZendLong(trial)
		return zend.SUCCESS
	}

	/* Increment the max so the range is inclusive of max */

	umax++

	/* Powers of two are not biased */

	if (umax&umax - 1) != 0 {

		/* Ceiling under which ZEND_LONG_MAX % max == 0 */

		var limit zend.ZendUlong = UINT64_MAX - UINT64_MAX%umax - 1

		/* Discard numbers over the limit to avoid modulo bias */

		for trial > limit {
			if PhpRandomBytes(&trial, g.SizeOf("trial"), should_throw) == zend.FAILURE {
				return zend.FAILURE
			}
		}

		/* Discard numbers over the limit to avoid modulo bias */

	}
	*result = zend_long(trial%umax + min)
	return zend.SUCCESS
}

/* }}} */

func ZifRandomInt(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var min zend.ZendLong
	var max zend.ZendLong
	var result zend.ZendLong
	for {
		var _flags int = 1 << 2
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

			if zend.ZendParseArgLong(_arg, &min, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			if zend.ZendParseArgLong(_arg, &max, &_dummy, 0, 0) == 0 {
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
	if min > max {
		zend.ZendThrowException(zend.ZendCeError, "Minimum value must be less than or equal to the maximum value", 0)
		return
	}
	if PhpRandomInt(min, max, &result, 1) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = result
	__z.u1.type_info = 4
	return
}

/* }}} */
