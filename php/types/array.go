package types

import (
	"errors"
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
func (k ArrayKey) ToZval() Zval {
	if k.IsStrKey() {
		return ZvalString(k.StrKey())
	} else {
		return ZvalLong(k.IdxKey())
	}
}

// ArrayPair
// Array内的键值对。纯数值无引用，直接修改不会影响数组
type ArrayPair struct {
	Key ArrayKey
	Val Zval
}

func MakeArrayPair(key ArrayKey, val Zval) ArrayPair {
	return ArrayPair{Key: key, Val: val}
}
func (p ArrayPair) IsStrKey() bool { return p.Key.IsStrKey() }
func (p ArrayPair) StrKey() string { return p.Key.StrKey() }
func (p ArrayPair) IdxKey() int    { return p.Key.IdxKey() }
func (p ArrayPair) IsValid() bool  { return !p.Val.IsUndef() }

// ArrayPosition
type ArrayPosition = int

const InvalidArrayPos = math.MinInt
const maxArrayPos = math.MaxInt

// Array
type Array struct {
	data ArrayData

	// flags todo 待合并
	flags     uint8
	protected bool
	writable  bool // todo 完成写锁逻辑
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
func NewSymtableArray(cap int) *Array {
	data := newSymtableData(cap)
	return &Array{data: data, writable: true}
}
func NewArrayOf(values ...Zval) *Array {
	data := NewArrayDataList(
		values,
		func(v Zval) Zval { return v },
		func(v Zval) (Zval, bool) { return v, true },
	)
	return &Array{data: data, writable: true}
}
func NewArrayOfZval(values []Zval) *Array {
	return NewArrayOf(values...)
}
func NewArrayOfInt(values []int) *Array {
	data := NewArrayDataList(
		values,
		func(v int) Zval { return ZvalLong(v) },
		func(v Zval) (int, bool) {
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
		func(v string) Zval { return ZvalString(v) },
		func(v Zval) (string, bool) {
			if v.IsString() {
				return v.String(), true
			}
			return "", false
		},
	)
	return &Array{data: data, writable: true}
}
func NewArrayOfPairs(pairs []ArrayPair) *Array {
	data := newArrayDataHtByData(pairs)
	return &Array{data: data, writable: true}
}

func (ht *Array) Clone() *Array {
	// todo 懒复制逻辑
	return ht
}

func (ht *Array) Clean() {
	ht.assertWritable()
	ht.data = emptyArrayData
	ht.protected = false
}

// 常用读操作

func (ht *Array) Len() int                 { return ht.data.Len() }
func (ht *Array) Cap() int                 { return ht.data.Cap() }
func (ht *Array) Count() int               { return ht.data.Count() }
func (ht *Array) Exists(key ArrayKey) bool { return ht.data.Exists(key) }
func (ht *Array) Find(key ArrayKey) Zval   { val, _ := ht.data.Find(key); return val }

func (ht *Array) Each(handler func(key ArrayKey, value Zval)) {
	_ = ht.data.Each(func(key ArrayKey, value Zval) error {
		handler(key, value)
		return nil
	})
}

func (ht *Array) EachEx(handler func(key ArrayKey, value Zval) error) error {
	return ht.data.Each(handler)
}

func (ht *Array) EachReserve(handler func(key ArrayKey, value Zval)) {
	_ = ht.data.EachReserve(func(key ArrayKey, value Zval) error {
		handler(key, value)
		return nil
	})
}

// 常用写操作

func (ht *Array) Add(key ArrayKey, value Zval) bool {
	assert(value.IsNotUndef())
	ht.assertWritable()
	ret, err := ht.data.Add(key, value)
	if errors.Is(err, arrayDataUnsupported) {
		ret, _ = ht.operableData().Add(key, value)
	}
	return ret
}
func (ht *Array) Update(key ArrayKey, value Zval) {
	assert(value.IsNotUndef())
	ht.assertWritable()
	err := ht.data.Update(key, value)
	if errors.Is(err, arrayDataUnsupported) {
		_ = ht.operableData().Update(key, value)
	}
}
func (ht *Array) Delete(key ArrayKey) bool {
	ht.assertWritable()
	ret, err := ht.data.Delete(key)
	if errors.Is(err, arrayDataUnsupported) {
		ret, _ = ht.operableData().Delete(key)
	}
	return ret
}
func (ht *Array) Append(value Zval) int {
	assert(value.IsNotUndef())
	ht.assertWritable()
	ret, err := ht.data.Append(value)
	if errors.Is(err, arrayDataUnsupported) {
		ret, _ = ht.operableData().Append(value)
	}
	return ret
}

func (ht *Array) assertWritable() { assert(ht.writable) }

// 返回支持所有操作的 ArrayData (隐式转换为 ht.data 为 *ArrayDataHt)
func (ht *Array) operableData() *ArrayDataHt {
	// 已经是 ArrayDataHt，就无法再转化了
	if d, ok := ht.data.(*ArrayDataHt); ok {
		return d
	}

	// 构建 ArrayDataHt 类型的 data 并拷贝已有数据
	data := ht.data
	newData := newArrayDataHt(data.Len())
	_ = data.Each(func(key ArrayKey, value Zval) error {
		_, _ = newData.Add(key, value)
		return nil
	})
	ht.data = newData
	return newData
}

// Methods use idx key

func (ht *Array) IndexExists(idx int) bool          { return ht.Exists(IdxKey(idx)) }
func (ht *Array) IndexFind(idx int) Zval            { return ht.Find(IdxKey(idx)) }
func (ht *Array) IndexAdd(idx int, value Zval) bool { return ht.Add(IdxKey(idx), value) }
func (ht *Array) IndexUpdate(idx int, value Zval)   { ht.Update(IdxKey(idx), value) }
func (ht *Array) IndexDelete(idx int) bool          { return ht.Delete(IdxKey(idx)) }

// Methods use string key

func (ht *Array) KeyExists(key string) bool          { return ht.Exists(StrKey(key)) }
func (ht *Array) KeyFind(key string) Zval            { return ht.Find(StrKey(key)) }
func (ht *Array) KeyAdd(key string, value Zval) bool { return ht.Add(StrKey(key), value) }
func (ht *Array) KeyUpdate(key string, value Zval)   { ht.Update(StrKey(key), value) }
func (ht *Array) KeyDelete(key string) bool          { return ht.Delete(StrKey(key)) }

// recursive

func (ht *Array) IsRecursive() bool   { return ht.protected }
func (ht *Array) ProtectRecursive()   { ht.protected = true }
func (ht *Array) UnprotectRecursive() { ht.protected = false }

// sort
type ArrayComparer interface {
	Compare(p1, p2 ArrayPair) int
}

type ArrayPairComparer func(p1, p2 ArrayPair) int

func (c ArrayPairComparer) Compare(p1, p2 ArrayPair) int { return c(p1, p2) }

type ArrayKeyComparer func(k1, k2 ArrayKey) int

func (c ArrayKeyComparer) Compare(p1, p2 ArrayPair) int { return c(p1.Key, p2.Key) }

type ArrayValueComparer func(v1, v2 Zval) int

func (c ArrayValueComparer) Compare(p1, p2 ArrayPair) int { return c(p1.Val, p2.Val) }

func (ht *Array) Sort(comparer ArrayComparer, renumber bool) {
	ht.assertWritable()

	if ht.Len() == 0 || (ht.Len() == 1 && !renumber) {
		return
	}

	// 将 ht.data 转成 *ArrayDataHt 后排序
	// todo 细分可优化情况单独处理 (例如，预判是否 IsSorted 以跳过排序、对 List 类型不保留key的排序可直接操作等 )
	ht.operableData().Sort(comparer, renumber)
}

// fast methods

func (ht *Array) Keys() []ArrayKey {
	var keys = make([]ArrayKey, 0, ht.Len())
	ht.Each(func(key ArrayKey, value Zval) {
		keys = append(keys, key)
	})
	return keys
}
func (ht *Array) Values() []Zval {
	var values = make([]Zval, 0, ht.Len())
	ht.Each(func(key ArrayKey, value Zval) {
		values = append(values, value)
	})
	return values
}
func (ht *Array) Pairs() []ArrayPair {
	var pairs = make([]ArrayPair, 0, ht.Len())
	ht.Each(func(key ArrayKey, value Zval) {
		pairs = append(pairs, MakeArrayPair(key, value))
	})
	return pairs
}
