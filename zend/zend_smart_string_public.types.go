// <<generate>>

package zend

import (
	"strings"
)

/**
 * SmartString
 */
type SmartString struct {
	c    *byte
	len_ int
	//a    int // cap ?

	buffer strings.Builder
}

func (this *SmartString) SetC(value *byte) { this.c = value /* todo delete */ }
func (this *SmartString) SetLen(value int) { this.len_ = value /* todo delete */ }

func (this *SmartString) GetC() *byte { return this.c /* todo 待修改 */ }
func (this *SmartString) GetLen() int { return this.buffer.Len() }
func (this *SmartString) GetA() int   { return this.buffer.Cap() }

// 分配内存，确认至少有 len_ 的未使用内存
func (this *SmartString) Alloc(len_ int) int {
	this.buffer.Grow(len_)
	return this.buffer.Len() + len_
}

func (this *SmartString) AppendString(str string) {
	this.buffer.WriteString(str)
}

func (this *SmartString) AppendByte(c byte) {
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
