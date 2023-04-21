package types

import (
	"fmt"
	"math"
	"runtime"
)

/**
 * ArrayKey
 */
type ArrayKey struct {
	index int
	str   *string
}

func StrKey(str string) ArrayKey  { return ArrayKey{0, &str} }
func IndexKey(index int) ArrayKey { return ArrayKey{index, nil} }

func (k ArrayKey) IsStrKey() bool { return k.str != nil }
func (k ArrayKey) IndexKey() int  { return k.index }
func (k ArrayKey) StrKey() string { return *k.str }
func (k ArrayKey) Keys() (index int, key string, isStrKey bool) {
	if k.IsStrKey() {
		return 0, *k.str, true
	} else {
		return k.index, "", false
	}
}
func (k ArrayKey) ToZval() *Zval {
	if k.IsStrKey() {
		return NewZvalString(k.StrKey())
	} else {
		return NewZvalLong(k.IndexKey())
	}
}

/**
 * ArrayPair
 */
type ArrayPair struct {
	key ArrayKey
	val *Zval
}

func MakeArrayPair(key ArrayKey, val *Zval) ArrayPair {
	return ArrayPair{key: key, val: val}
}
func (p ArrayPair) GetKey() ArrayKey { return p.key }
func (p ArrayPair) GetVal() *Zval    { return p.val }

/**
 * Bucket
 */
type Bucket struct {
	val Zval
	key ArrayKey
}

func NewBucket(key ArrayKey, zval *Zval) *Bucket {
	var bucket = &Bucket{key: key}
	ZVAL_COPY_VALUE(&bucket.val, zval)
	return bucket
}

func NewStrKeyBucket(strKey string, zval *Zval) *Bucket {
	var key = StrKey(strKey)
	return NewBucket(key, zval)
}

func NewIndexBucket(indexKey int, zval *Zval) *Bucket {
	var key = IndexKey(indexKey)
	return NewBucket(key, zval)
}

func (this *Bucket) IsStrKey() bool            { return this.key.IsStrKey() }
func (this *Bucket) StrKey() string            { return this.key.StrKey() }
func (this *Bucket) IndexKey() int             { return this.key.IndexKey() }
func (this *Bucket) Keys() (int, string, bool) { return this.key.Keys() }

func (this *Bucket) SetStrKey(key string)  { this.key = StrKey(key) }
func (this *Bucket) SetIndexKey(index int) { this.key = IndexKey(index) }
func (this *Bucket) GetArrayKey() ArrayKey { return this.key }

func (this *Bucket) GetVal() *Zval     { return &this.val }
func (this *Bucket) SetVal(zval *Zval) { ZVAL_COPY_VALUE(&this.val, zval) }

func (this *Bucket) IsValid() bool { return !this.val.IsUndef() }
func (this *Bucket) MarkInvalid()  { this.val.SetUndef() }

func (this *Bucket) CopyFrom(from *Bucket) {
	this.SetVal(from.GetVal())
	this.key = from.key
}

/**
 * Array
 * HashTable Data Layout
 * =====================
 *
 *                 +=============================+
 *                 | HT_HASH(ht, ht->nTableMask) |
 *                 | ...                         |
 *                 | HT_HASH(ht, -1)             |
 *                 +-----------------------------+
 * ht->arData ---> | Bucket[0]                   |
 *                 | ...                         |
 *                 | Bucket[ht->tableSize-1]    |
 *                 +=============================+
 */
type ArrayPosition = uint32
type Array struct {
	ZendRefcounted
	flags           ZendUchar
	elementsCount   uint32
	internalPointer uint32
	nextFreeElement int
	destructor      DtorFuncT

	data     []Bucket          // 实际存储数据的地方
	indexMap map[int]uint32    // 数字索引到具体位置的映射
	keyMap   map[string]uint32 // 字符串索引到具体位置的映射

	arData *Bucket // C 源码中存储数据的地方，实际不使用

	keys  []ArrayKey
	data0 map[ArrayKey]*Zval
}

var _ IRefcounted = &Array{}

/**
 * Constructor && Init
 */
func NewArray(size int) *Array {
	return NewArrayEx(size, nil)
}
func NewArrayEx(size int, pDestructor DtorFuncT) *Array {
	var ht = new(Array)
	ht.Init(size, pDestructor)
	return ht
}
func (ht *Array) Init(size int, pDestructor DtorFuncT) {
	var data []Bucket
	if size > 0 {
		data = make([]Bucket, 0, size)
	}

	*ht = Array{
		destructor: pDestructor,

		// 数据存储
		data:     data,
		indexMap: make(map[int]uint32),    // todo 改为 nil，延迟初始化
		keyMap:   make(map[string]uint32), // todo 改为 nil，延迟初始化
	}

	// GC 信息
	ht.SetRefcount(1)
	ht.SetGcTypeInfo(uint32(IS_ARRAY))

	// 析构函数
	if pDestructor != nil {
		runtime.SetFinalizer(ht, ht.DestroyEx)
	}
}

/* init */
func (ht *Array) SetBy(arr *Array) {
	ht.flags = arr.flags
	ht.elementsCount = arr.elementsCount
	ht.nextFreeElement = arr.nextFreeElement
	ht.destructor = arr.destructor

	ht.data = arr.data
	ht.indexMap = arr.indexMap
	ht.keyMap = arr.keyMap

	ZendHashInternalPointerReset(ht)
}

// 实际元素个数，从使用者角度的数组大小
func (ht *Array) Len() int                  { return int(ht.elementsCount) }
func (ht *Array) Cap() int                  { return cap(ht.data) }
func (ht *Array) Bucket(pos uint32) *Bucket { return &ht.data[pos] }

/* data -> Array.data */
func (ht *Array) clearData() {
	ht.assertRc1()

	ht.elementsCount = 0
	ht.internalPointer = 0
	ht.nextFreeElement = 0
	ht.data = nil
	ht.indexMap = make(map[int]uint32)
	ht.keyMap = make(map[string]uint32)
}
func (ht *Array) appendBucketStr(strKey string, zv *Zval) *Bucket {
	var bucket = NewStrKeyBucket(strKey, zv)
	return ht.appendBucket(bucket)
}

func (ht *Array) appendBucketIndex(indexKey int, zv *Zval) *Bucket {
	var bucket = NewIndexBucket(indexKey, zv)
	return ht.appendBucket(bucket)
}
func (ht *Array) appendBucket(bucket *Bucket) *Bucket {
	// 尝试 resize
	ht.resizeIfFull()

	// 添加到 data
	var idx = uint32(len(ht.data))
	ht.elementsCount++
	ht.data = append(ht.data, *bucket)

	// 更新 map
	if bucket.IsStrKey() {
		ht.keyMap[bucket.StrKey()] = idx
	} else {
		ht.indexMap[bucket.IndexKey()] = idx
	}

	if !bucket.IsStrKey() {
		var indexKey = bucket.IndexKey()
		// 更新 nextFreeElement
		if indexKey > ht.nextFreeElement {
			if indexKey < MaxLong {
				ht.nextFreeElement = indexKey + 1
			} else {
				ht.nextFreeElement = MaxLong
			}
		}
	}

	return &ht.data[idx]
}
func (ht *Array) resizeIfFull() {
	dataSize := len(ht.data)
	if dataSize == cap(ht.data) {
		// 若空隙率过高，重新压缩；否则，跳过扩容 (后面会由 append(ht.data) 触发自动扩容)
		if dataSize > int(ht.elementsCount+(ht.elementsCount>>5)) {
			ht.Rehash()
		} else if dataSize >= math.MaxInt32 {
			triggerError(fmt.Sprintf("Possible integer overflow in memory allocation (%d)", dataSize*2))
		}
	}
}
func (ht *Array) deleteBucket(pos uint32) {
	ht.assertRc1()
	assert(int(pos) < len(ht.data))

	var p = &ht.data[pos]
	assert(p.IsValid())

	// 移除映射
	if p.key.IsStrKey() {
		delete(ht.keyMap, p.key.StrKey())
	} else {
		delete(ht.indexMap, p.key.IndexKey())
	}

	// 减少有效元素
	ht.elementsCount--

	// 更新内部指针和遍历器指针
	if ht.internalPointer == pos {
		var newIdx = ht.validPosVal(pos + 1)
		if ht.internalPointer == pos {
			ht.internalPointer = newIdx
		}
	}

	// 析构函数
	if ht.destructor != nil {
		var tmp Zval
		ZVAL_COPY_VALUE(&tmp, p.GetVal())
		ht.destructor(&tmp)
	}

	// 设置数据不可用
	p.MarkInvalid()

	// 若删除队尾元素，尝试清除 data 队尾无用数据
	dataSize := uint32(len(ht.data))
	if pos == dataSize-1 {
		newDataSize := dataSize
		for newDataSize > 0 && !ht.data[newDataSize-1].IsValid() {
			newDataSize--
		}

		ht.data = ht.data[:newDataSize]
		if ht.internalPointer > newDataSize {
			ht.internalPointer = newDataSize
		}
	}
}

func (ht *Array) posBucket(p *Bucket) (uint32, bool) {
	if p.IsStrKey() {
		if pos, ok := ht.keyMap[p.StrKey()]; ok {
			return pos, true
		}
		return 0, false
	} else {
		if pos, ok := ht.indexMap[p.IndexKey()]; ok {
			return pos, true
		}
		return 0, false
	}
}

// 移动 bucket 到新位置
func (ht *Array) moveBucket(pos uint32, newPos uint32) {
	assert(newPos <= pos)
	if newPos == pos {
		return
	}
	(&ht.data[newPos]).CopyFrom(&ht.data[pos])
	if ht.internalPointer == pos {
		ht.internalPointer = newPos
	}
}

/* hash -> Array.indexMap & Array.keyMap */
func (ht *Array) Rehash() {
	// reset hash
	ht.assertRc1()
	ht.indexMap = make(map[int]uint32)
	ht.keyMap = make(map[string]uint32)

	// 空数组快速清空
	if ht.elementsCount == 0 {
		ht.data = nil
		return
	}

	// 移除 data 中的空位
	ht.removeHoles()

	// 重建 hash
	ht.eachBucket(func(pos uint32, p *Bucket) {
		if p.IsStrKey() {
			ht.keyMap[p.StrKey()] = pos
		} else {
			ht.indexMap[p.IndexKey()] = pos
		}
	})
}

/* misc */
func (ht *Array) assertRc1()      { assert(ht.GetRefcount() == 1) }
func (ht *Array) assertWritable() { assert(ht.GetRefcount() == 1) }

/** Array.flags */
func (ht *Array) CopyFlags(arr *Array) { ht.flags = arr.flags }

func (ht *Array) IsPacked() bool       { return ht.flags&HASH_FLAG_PACKED != 0 }
func (ht *Array) HasEmptyIndex() bool  { return ht.flags&HASH_FLAG_HAS_EMPTY_IND != 0 }
func (ht *Array) MarkHasEmptyIndex()   { ht.flags |= HASH_FLAG_HAS_EMPTY_IND }
func (ht *Array) UnmarkHasEmptyIndex() { ht.flags &^= HASH_FLAG_HAS_EMPTY_IND }

func (ht *Array) Values() []*Zval {
	var values = make([]*Zval, 0, ht.Len())
	ht.Foreach(func(key ArrayKey, value *Zval) {
		values = append(values, value)
	})
	return values
}
func (ht *Array) Pairs() []ArrayPair {
	var pairs = make([]ArrayPair, 0, ht.Len())
	ht.Foreach(func(key ArrayKey, value *Zval) {
		pairs = append(pairs, MakeArrayPair(key, value))
	})
	return pairs
}

func (ht *Array) MapWithKey(mapper func(key ArrayKey, value *Zval) (ArrayKey, *Zval)) *Array {
	// todo 考虑 rehash 等操作 或 对其他属性的处理
	arr := NewArray(ht.Len())
	ht.Foreach(func(key ArrayKey, value *Zval) {
		newKey, newValue := mapper(key, value)
		arr.Add(newKey, newValue)
	})
	return arr
}
