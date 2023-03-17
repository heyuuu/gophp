// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
)

func RandomGlobalsCtor(random_globals_p *PhpRandomGlobals) { random_globals_p.SetFd(-1) }
func RandomGlobalsDtor(random_globals_p *PhpRandomGlobals) {
	if random_globals_p.GetFd() > 0 {
		close(random_globals_p.GetFd())
		random_globals_p.SetFd(-1)
	}
}
func ZmStartupRandom(type_ int, module_number int) int {
	RandomGlobalsCtor(&RandomGlobals)
	return zend.SUCCESS
}
func ZmShutdownRandom(type_ int, module_number int) int {
	RandomGlobalsDtor(&RandomGlobals)
	return zend.SUCCESS
}
func PhpRandomBytes(bytes any, size int, should_throw zend.ZendBool) int {
	var read_bytes int = 0
	var n ssize_t
	if read_bytes < size {
		var fd int = RANDOM_G(fd)
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
			RANDOM_G(fd) = fd
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
func ZifRandomBytes(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var size zend.ZendLong
	var bytes *zend.ZendString
	for {
		var _flags int = zend.ZEND_PARSE_PARAMS_THROW
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
			if zend.ZendParseArgLong(_arg, &size, &_dummy, 0, 0) == 0 {
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
	if size < 1 {
		zend.ZendThrowException(zend.ZendCeError, "Length must be greater than 0", 0)
		return
	}
	bytes = zend.ZendStringAlloc(size, 0)
	if PhpRandomBytesThrow(bytes.GetVal(), size) == zend.FAILURE {
		zend.ZendStringReleaseEx(bytes, 0)
		return
	}
	bytes.GetVal()[size] = '0'
	return_value.SetString(bytes)
	return
}
func PhpRandomInt(min zend.ZendLong, max zend.ZendLong, result *zend.ZendLong, should_throw zend.ZendBool) int {
	var umax zend.ZendUlong
	var trial zend.ZendUlong
	if min == max {
		*result = min
		return zend.SUCCESS
	}
	umax = zend.ZendUlong(max - zend.ZendUlong(min))
	if PhpRandomBytes(&trial, b.SizeOf("trial"), should_throw) == zend.FAILURE {
		return zend.FAILURE
	}

	/* Special case where no modulus is required */

	if umax == zend.ZEND_ULONG_MAX {
		*result = zend.ZendLong(trial)
		return zend.SUCCESS
	}

	/* Increment the max so the range is inclusive of max */

	umax++

	/* Powers of two are not biased */

	if (umax&umax - 1) != 0 {

		/* Ceiling under which ZEND_LONG_MAX % max == 0 */

		var limit zend.ZendUlong = zend.ZEND_ULONG_MAX - zend.ZEND_ULONG_MAX%umax - 1

		/* Discard numbers over the limit to avoid modulo bias */

		for trial > limit {
			if PhpRandomBytes(&trial, b.SizeOf("trial"), should_throw) == zend.FAILURE {
				return zend.FAILURE
			}
		}

		/* Discard numbers over the limit to avoid modulo bias */

	}
	*result = zend_long(trial%umax + min)
	return zend.SUCCESS
}
func ZifRandomInt(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var min zend.ZendLong
	var max zend.ZendLong
	var result zend.ZendLong
	for {
		var _flags int = zend.ZEND_PARSE_PARAMS_THROW
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
			if zend.ZendParseArgLong(_arg, &min, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &max, &_dummy, 0, 0) == 0 {
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
	if min > max {
		zend.ZendThrowException(zend.ZendCeError, "Minimum value must be less than or equal to the maximum value", 0)
		return
	}
	if PhpRandomIntThrow(min, max, &result) == zend.FAILURE {
		return
	}
	return_value.SetLong(result)
	return
}
