// <<generate>>

package zend

import (
	b "sik/builtin"
	"strings"
)

/**
 * SmartStr
 */
type SmartStr struct {
	s *ZendString
	a int

	buffer strings.Builder
}

func (this *SmartStr) SetS(value *ZendString) { this.s = value }
func (this *SmartStr) SetA(value int)         { this.a = value }
func (this *SmartStr) GetS() *ZendString {
	// todo 需要确认是否兼容 ZendStringAlloc() 但未使用时的空 []byte
	return NewZendString(this.GetStr())
}

func (this *SmartStr) GetA() int      { return this.buffer.Cap() }
func (this *SmartStr) GetStr() string { return this.buffer.String() }

func (this *SmartStr) Alloc(len_ int) int {
	// 确保有 len_ 可用
	this.buffer.Grow(len_)
	return this.buffer.Len() + len_
}

func (this *SmartStr) AppendString(str string) {
	this.buffer.WriteString(str)
}

func (this *SmartStr) AppendS(src *byte) {
	var str = b.CastStrAuto(src)
	this.buffer.WriteString(str)
}

func (this *SmartStr) AppendL(src *byte, len_ int) {
	var str = b.CastStr(src, len_)
	this.buffer.WriteString(str)
}

func (this *SmartStr) AppendC(c byte) {
	this.buffer.WriteByte(c)
}

func (this *SmartStr) AppendSmartStr(str *SmartStr) {
	this.buffer.WriteString(str.GetStr())
}

func (this *SmartStr) SetString(str string) {
	this.Reset()
	this.buffer.WriteString(str)
}

func (this *SmartStr) Reset() {
	this.buffer.Reset()
}

func (this *SmartStr) Free() {
	this.Reset()
}

func (this *SmartStr) ZeroTail() {
	// c 字符串尾部设置0
}
