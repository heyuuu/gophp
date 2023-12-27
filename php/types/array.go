package types

import (
	"math"
)

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

func (k ArrayKey) IsIdxKey() bool { return k.str == "" }
func (k ArrayKey) IsStrKey() bool { return k.str != "" }
func (k ArrayKey) IdxKey() int    { return k.idx }
func (k ArrayKey) StrKey() string {
	if k.str != "" && k.str[0] == arrayKeySign {
		return k.str[1:]
	} else {
		return k.str
	}
}

// ArrayPosition
type ArrayPosition = int

const InvalidArrayPos = math.MinInt
const maxArrayPos = math.MaxInt

// Array
type Array struct {
	data     ArrayData
	writable bool // todo 完成写锁逻辑
}

/**
 * Constructor && Init
 */
func NewArray() *Array {
	return NewArrayCap(0)
}
func NewArrayCap(cap int) *Array {
	var data ArrayData = emptyArrayData
	if cap > 0 {
		data = newArrayDataHt(cap)
	}

	return &Array{data: data, writable: true}
}
func NewArrayOf(values ...*Zval) *Array {
	data := NewArrayDataList(
		values,
		func(v *Zval) *Zval { return v },
		func(v *Zval) (*Zval, bool) { return v, true },
	)
	return &Array{data: data, writable: true}
}
func NewArrayOfZval(values []*Zval) *Array {
	return NewArrayOf(values...)
}
func NewArrayOfInt(values []int) *Array {
	data := NewArrayDataList(
		values,
		func(v int) *Zval { return NewZvalLong(v) },
		func(v *Zval) (int, bool) {
			if v.IsLong() {
				return v.Long(), true
			}
			return 0, false
		},
	)
	return &Array{data: data, writable: true}
}
func NewArrayOfString(values []string) *Array {
	data := NewArrayDataList(
		values,
		func(v string) *Zval { return NewZvalString(v) },
		func(v *Zval) (string, bool) {
			if v.IsString() {
				return v.String(), true
			}
			return "", false
		},
	)
	return &Array{data: data, writable: true}
}

// 常用读操作

func (ht *Array) Len() int                 { return ht.data.Len() }
func (ht *Array) Cap() int                 { return ht.data.Cap() }
func (ht *Array) Count() int               { return ht.data.Count() }
func (ht *Array) Exists(key ArrayKey) bool { return ht.data.Exists(key) }
func (ht *Array) Find(key ArrayKey) *Zval  { val, _ := ht.data.Find(key); return val }

// 常用写操作

func (ht *Array) Add(key ArrayKey, value *Zval) bool {
	ht.assertWritable()
	ret, err := ht.data.Add(key, value)
	if err == arrayDataUnsupported && ht.makeOperable() {
		ret, _ = ht.data.Add(key, value)
	}
	return ret
}
func (ht *Array) Update(key ArrayKey, value *Zval) {
	ht.assertWritable()
	err := ht.data.Update(key, value)
	if err == arrayDataUnsupported && ht.makeOperable() {
		_ = ht.data.Update(key, value)
	}
}
func (ht *Array) Delete(key ArrayKey) bool {
	ht.assertWritable()
	ret, err := ht.data.Delete(key)
	if err == arrayDataUnsupported && ht.makeOperable() {
		ret, _ = ht.data.Delete(key)
	}
	return ret
}
func (ht *Array) Append(value *Zval) int {
	ht.assertWritable()
	ret, err := ht.data.Append(value)
	if err == arrayDataUnsupported && ht.makeOperable() {
		ret, _ = ht.data.Append(value)
	}
	return ret
}

func (ht *Array) assertWritable() { assert(ht.writable) }

// 使 ArrayData 支持所有操作
func (ht *Array) makeOperable() bool {
	// 已经是 ArrayDataHt，就无法再转化了
	if _, ok := ht.data.(*ArrayDataHt); ok {
		return false
	}

	// 构建 ArrayDataHt 类型的 data 并拷贝已有数据
	data := ht.data
	newData := newArrayDataHt(data.Len())
	_ = data.Each(func(key ArrayKey, value *Zval) error {
		_, _ = newData.Add(key, value)
		return nil
	})
	ht.data = newData
	return true
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
