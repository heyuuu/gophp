// <<generate>>

package zend

import (
	"bytes"
	b "sik/builtin"
)

/**
 * SmartString
 */
type SmartString struct {
	c    *byte
	len_ int
	a    int // cap ?

	buffer bytes.Buffer
}

func (this *SmartString) GetC() *byte      { return this.c }
func (this *SmartString) SetC(value *byte) { this.c = value }
func (this *SmartString) GetLen() int      { return this.buffer.Len() }
func (this *SmartString) SetLen(value int) { this.len_ = value }
func (this *SmartString) GetA() int        { return this.a }
func (this *SmartString) SetA(value int)   { this.a = value }

// 分配内存，确认至少有 len_ 的未使用内存
func (this *SmartString) alloc(len_ int) int {
	// 当前实现无需手动扩展内存
	return this.len_ + len_
}

func (this *SmartString) AppendS(src *byte) {
	var str = b.CastStrAuto(src)
	this.buffer.WriteString(str)
}

func (this *SmartString) AppendL(src *byte, len_ int) {
	var str = b.CastStr(src, len_)
	this.buffer.WriteString(str)
}

func (this *SmartString) AppendC(c byte) {
	this.buffer.WriteByte(c)
}

func (this *SmartString) Reset() {
	this.buffer.Reset()
}

func (this *SmartString) Free() {
	this.Reset()
}

func (this *SmartString) ZeroTail() {
	// c 字符串尾部设置0
}
