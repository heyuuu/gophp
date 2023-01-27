package builtin

import "unsafe"

type StrArg struct {
	str  *byte
	len_ uint
}

func NewStrArg(str *byte, len_ uint) StrArg { return StrArg{str, len_} }

func (this StrArg) Len() uint     { return this.len_ }
func (this StrArg) StrPtr() *byte { return this.str }
func (this StrArg) Str() string {
	// todo 此段代码仅表意，实际不应依赖此实现 (因为无法保证 *byte 后续内存有效)
	var bytes = make([]byte, this.len_)
	var ptr = uintptr(unsafe.Pointer(this.str))
	for i := uint(0); i < this.len_; i++ {
		bytes[i] = *(*byte)(unsafe.Pointer(ptr + uintptr(i)))
	}

	return string(bytes)
}

func (this StrArg) Hash() uint {
	return HashStr(this.Str())
}
