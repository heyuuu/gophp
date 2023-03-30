package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
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
					faults.ThrowException(faults.ZendCeException, "Cannot open source device", 0)
				}
				return types.FAILURE
			}

			/* Does the file exist and is it a character device? */

			if fstat(fd, &st) != 0 || !(S_ISCHR(st.st_mode)) {
				close(fd)
				if should_throw != 0 {
					faults.ThrowException(faults.ZendCeException, "Error reading from source device", 0)
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
				faults.ThrowException(faults.ZendCeException, "Could not gather sufficient random data", 0)
			}
			return types.FAILURE
		}
	}
	return types.SUCCESS
}
func ZifRandomBytes(executeData zpp.Ex, return_value zpp.Ret, length *types.Zval) {
	var size zend.ZendLong
	var bytes *types.String
	for {
		var _flags int = zpp.FlagThrow
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			size = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if size < 1 {
		faults.ThrowException(faults.ZendCeError, "Length must be greater than 0", 0)
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
func ZifRandomInt(executeData zpp.Ex, return_value zpp.Ret, min *types.Zval, max *types.Zval) {
	var min zend.ZendLong
	var max zend.ZendLong
	var result zend.ZendLong
	for {
		var _flags int = zpp.FlagThrow
		var _min_num_args int = 2
		var _max_num_args int = 2

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			min = fp.ParseLong()
			max = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if min > max {
		faults.ThrowException(faults.ZendCeError, "Minimum value must be less than or equal to the maximum value", 0)
		return
	}
	if PhpRandomIntThrow(min, max, &result) == types.FAILURE {
		return
	}
	return_value.SetLong(result)
	return
}
