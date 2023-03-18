// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
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
	return types.SUCCESS
}
func ZmShutdownRandom(type_ int, module_number int) int {
	RandomGlobalsDtor(&RandomGlobals)
	return types.SUCCESS
}
func PhpRandomBytes(bytes any, size int, should_throw types.ZendBool) int {
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
				return types.FAILURE
			}

			/* Does the file exist and is it a character device? */

			if fstat(fd, &st) != 0 || !(S_ISCHR(st.st_mode)) {
				close(fd)
				if should_throw != 0 {
					zend.ZendThrowException(zend.ZendCeException, "Error reading from source device", 0)
				}
				return types.FAILURE
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
			return types.FAILURE
		}
	}
	return types.SUCCESS
}
func ZifRandomBytes(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var size zend.ZendLong
	var bytes *types.ZendString
	for {
		var _flags int = argparse.ZEND_PARSE_PARAMS_THROW
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if !zend.ZendParseArgLong00(_arg, &size) {
				_expected_type = argparse.Z_EXPECTED_LONG
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != argparse.ZPP_ERROR_OK {
			if (_flags & argparse.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == argparse.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & argparse.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == argparse.ZPP_ERROR_WRONG_CLASS {
					if (_flags & argparse.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == argparse.ZPP_ERROR_WRONG_ARG {
					if (_flags & argparse.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	bytes = types.ZendStringAlloc(size, 0)
	if PhpRandomBytesThrow(bytes.GetVal(), size) == types.FAILURE {
		types.ZendStringReleaseEx(bytes, 0)
		return
	}
	bytes.GetVal()[size] = '0'
	return_value.SetString(bytes)
	return
}
func PhpRandomInt(min zend.ZendLong, max zend.ZendLong, result *zend.ZendLong, should_throw types.ZendBool) int {
	var umax zend.ZendUlong
	var trial zend.ZendUlong
	if min == max {
		*result = min
		return types.SUCCESS
	}
	umax = zend.ZendUlong(max - zend.ZendUlong(min))
	if PhpRandomBytes(&trial, b.SizeOf("trial"), should_throw) == types.FAILURE {
		return types.FAILURE
	}

	/* Special case where no modulus is required */

	if umax == zend.ZEND_ULONG_MAX {
		*result = zend.ZendLong(trial)
		return types.SUCCESS
	}

	/* Increment the max so the range is inclusive of max */

	umax++

	/* Powers of two are not biased */

	if (umax&umax - 1) != 0 {

		/* Ceiling under which ZEND_LONG_MAX % max == 0 */

		var limit zend.ZendUlong = zend.ZEND_ULONG_MAX - zend.ZEND_ULONG_MAX%umax - 1

		/* Discard numbers over the limit to avoid modulo bias */

		for trial > limit {
			if PhpRandomBytes(&trial, b.SizeOf("trial"), should_throw) == types.FAILURE {
				return types.FAILURE
			}
		}

		/* Discard numbers over the limit to avoid modulo bias */

	}
	*result = zend_long(trial%umax + min)
	return types.SUCCESS
}
func ZifRandomInt(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var min zend.ZendLong
	var max zend.ZendLong
	var result zend.ZendLong
	for {
		var _flags int = argparse.ZEND_PARSE_PARAMS_THROW
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if !zend.ZendParseArgLong00(_arg, &min) {
				_expected_type = argparse.Z_EXPECTED_LONG
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if !zend.ZendParseArgLong00(_arg, &max) {
				_expected_type = argparse.Z_EXPECTED_LONG
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != argparse.ZPP_ERROR_OK {
			if (_flags & argparse.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == argparse.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & argparse.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == argparse.ZPP_ERROR_WRONG_CLASS {
					if (_flags & argparse.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == argparse.ZPP_ERROR_WRONG_ARG {
					if (_flags & argparse.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if PhpRandomIntThrow(min, max, &result) == types.FAILURE {
		return
	}
	return_value.SetLong(result)
	return
}
