package zend

import (
	"strconv"
	"strings"
)

const SMART_STR_START_SIZE = 256

/**
 * SmartStr
 */
type SmartString = SmartStr
type SmartStr struct {
	buffer strings.Builder
}

func (this *SmartStr) GetLen() int    { return this.buffer.Len() }
func (this *SmartStr) GetStr() string { return this.buffer.String() }

// 分配内存，确认至少有 len_ 的未使用内存
func (this *SmartStr) Alloc(len_ int) int {
	if this.buffer.Cap() == 0 && len_ < SMART_STR_START_SIZE {
		// 初始化时，最小尺寸为 SMART_STR_START_SIZE，避免小尺寸重复扩展
		this.buffer.Grow(SMART_STR_START_SIZE)
	} else {
		this.buffer.Grow(len_)
	}
	return this.buffer.Len() + len_
}

func (this *SmartStr) AppendString(str string) {
	this.buffer.WriteString(str)
}

func (this *SmartStr) AppendByte(c byte) {
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

/**
 * 快捷方法
 */
func (this *SmartStr) AppendLong(num ZendLong) {
	var str = strconv.FormatInt(int64(num), 10)
	this.AppendString(str)
}

func (this *SmartStr) AppendUlong(num ZendUlong) {
	var str = strconv.FormatUint(uint64(num), 10)
	this.AppendString(str)
}

/**
 * todo 待移除方法
 */

func (this *SmartStr) GetS() *ZendString {
	// todo 需要确认是否兼容 ZendStringAlloc() 但未使用时的空 []byte
	return NewZendString(this.GetStr())
}
func (this *SmartStr) GetC() *byte {
	// todo 仅占位，实际使用需替换
	var str = this.GetStr()
	var char = str[0]
	return &char
}
func (this *SmartStr) SetC(value *byte) {
	/* todo delete */
	if value == nil {
		this.Reset()
	} else {
		ZEND_ASSERT(false)
	}
}
func (this *SmartStr) SetLen(value int) {
	/* todo delete */
	if value == 0 {
		this.Reset()
	} else if value < this.buffer.Len() {
		this.SetString(this.GetStr()[:value])
	} else if value > this.buffer.Len() {
		var appendBytes = make([]byte, value-this.buffer.Len())
		this.buffer.Write(appendBytes)
	}
}
