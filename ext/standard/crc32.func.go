// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/types"
	"sik/zend/zpp"
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

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			p, nr = fp.ParseString()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
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
