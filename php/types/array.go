package types

import "github.com/heyuuu/gophp/php/lang"

// arrayKeySign ArrayKey 内部标识符，用于标识是否为字符串
// 不再使用 \x00，因为内部会使用 \x00 开始标识一些特殊字符串，较容易冲突 (比如 protected 的方法名等)
const arrayKeySign = '\x01'

/**
 * ArrayKey
 * 	- 可直接比较
 *  - 零值为数字 0
 * 为减少内存占用，省略类型标识字段，采用以下方式确认类型:
 * - str == ""        		时: int 类型，值为 idx
 * - str[0] != arrayKeySign 时: string 类型，值为 str
 * - str[0] == arrayKeySign 时: string 类型，值为 str[1:]
 */
type ArrayKey struct {
	idx int
	str string
}

func IdxKey(idx int) ArrayKey { return ArrayKey{idx, ""} }
func StrKey(str string) ArrayKey {
	if str == "" || str[0] == arrayKeySign {
		str = string(arrayKeySign) + str
	}
	return ArrayKey{0, str}
}

func (k ArrayKey) IsStrKey() bool { return k.str != "" }
func (k ArrayKey) IdxKey() int    { return k.idx }
func (k ArrayKey) StrKey() string {
	if k.str != "" && k.str[0] == arrayKeySign {
		return k.str[1:]
	} else {
		return k.str
	}
}

// Array
type Array struct {
	data0    ArrayData
	writable bool // todo 完成写锁逻辑
}

/**
 * Constructor && Init
 */
func NewArray() *Array {
	return NewArrayCap(0)
}
func NewArrayCap(cap int) *Array {
	return &Array{}
}

func NewArrayOf(values ...*Zval) *Array {
	return NewArrayOfZval(values)
}

func NewArrayOfZval(values []*Zval) *Array {
	// todo 优化
	arr := NewArrayCap(len(values))
	for _, value := range values {
		arr.Append(value)
	}
	return arr
}

func (ht *Array) Len() int { return ht.data0.Len() }

/* misc */
func (ht *Array) assertWritable() { lang.Assert(ht.writable) }

// Methods use ArrayKey

func (ht *Array) Exists(key ArrayKey) bool { return ht.data0.Exists(key) }
func (ht *Array) Find(key ArrayKey) *Zval  { return ht.data0.Find(key) }
func (ht *Array) Add(key ArrayKey, value *Zval) bool {
	return ht.data0.Add(key, value)
}
func (ht *Array) Update(key ArrayKey, value *Zval) {
	ht.data0.Update(key, value)
}
func (ht *Array) Delete(key ArrayKey) bool {
	return ht.data0.Delete(key)
}
func (ht *Array) Append(value *Zval) int {
	return ht.data0.Push(value)
}

// Methods use idx key

func (ht *Array) IndexExists(idx int) bool           { return ht.Exists(IdxKey(idx)) }
func (ht *Array) IndexFind(idx int) *Zval            { return ht.Find(IdxKey(idx)) }
func (ht *Array) IndexAdd(idx int, value *Zval) bool { return ht.Add(IdxKey(idx), value) }
func (ht *Array) IndexUpdate(idx int, value *Zval)   { ht.Update(IdxKey(idx), value) }
func (ht *Array) IndexDelete(idx int) bool           { return ht.Delete(IdxKey(idx)) }

// Methods use string key

func (ht *Array) KeyExists(key string) bool           { return ht.Exists(StrKey(key)) }
func (ht *Array) KeyFind(key string) *Zval            { return ht.Find(StrKey(key)) }
func (ht *Array) KeyAdd(key string, value *Zval) bool { return ht.Add(StrKey(key), value) }
func (ht *Array) KeyUpdate(key string, value *Zval)   { ht.Update(StrKey(key), value) }
func (ht *Array) KeyDelete(key string) bool           { return ht.Delete(StrKey(key)) }
