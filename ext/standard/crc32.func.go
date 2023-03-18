// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
)

func CRC32(crc int, ch __auto__) int {
	crc = crc>>8 ^ Crc32tab[(crc^ch)&0xff]
	return crc
}
func HasCrc32Insn() int {
	/* Only go through the runtime detection once. */

	var res int = -1
	if res != -1 {
		return res
	}
	res = 0
	return res
}
func Crc32Aarch64(crc uint32, p *byte, nr int) uint32 {
	for nr >= b.SizeOf("uint64_t") {
		crc = __crc32d(crc, *((*uint64)(p)))
		p += b.SizeOf("uint64_t")
		nr -= b.SizeOf("uint64_t")
	}
	if nr >= b.SizeOf("int32_t") {
		crc = __crc32w(crc, *((*uint32)(p)))
		p += b.SizeOf("uint32_t")
		nr -= b.SizeOf("uint32_t")
	}
	if nr >= b.SizeOf("int16_t") {
		crc = __crc32h(crc, *((*uint16)(p)))
		p += b.SizeOf("uint16_t")
		nr -= b.SizeOf("uint16_t")
	}
	if nr != 0 {
		crc = __crc32b(crc, *p)
	}
	return crc
}
func PhpIfCrc32(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var p *byte
	var nr int
	var crcinit uint32 = 0
	var crc uint32
	for {
		var _flags int = 0
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
			if zend.ZendParseArgString(_arg, &p, &nr, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
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
	crc = crcinit ^ 0xffffffff
	if HasCrc32Insn() != 0 {
		crc = Crc32Aarch64(crc, p, nr)
		return_value.SetLong(crc ^ 0xffffffff)
		return
	}
	for ; b.PostDec(&nr); p++ {
		crc = crc>>8&0xffffff ^ Crc32tab[(crc^(*p))&0xff]
	}
	return_value.SetLong(crc ^ 0xffffffff)
	return
}
