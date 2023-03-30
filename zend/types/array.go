package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"math"
)

/**
 * ArrayKey
 */
type ArrayKey struct {
	index int
	key   *string
}

func MakeStrKey(str string) ArrayKey  { return ArrayKey{0, &str} }
func MakeIndexKey(index int) ArrayKey { return ArrayKey{index, nil} }

func (this ArrayKey) IsStrKey() bool { return this.key != nil }
func (this ArrayKey) IndexKey() int  { return this.index }
func (this ArrayKey) StrKey() string { return *this.key }
func (this ArrayKey) Keys() (index int, key string, isStrKey bool) {
	if this.IsStrKey() {
		return 0, *this.key, true
	} else {
		return this.index, "", false
	}
}

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
	var key = MakeStrKey(strKey)
	return NewBucket(key, zval)
}

func NewIndexBucket(indexKey int, zval *Zval) *Bucket {
	var key = MakeIndexKey(indexKey)
	return NewBucket(key, zval)
}

func (this *Bucket) IsStrKey() bool            { return this.key.IsStrKey() }
func (this *Bucket) StrKey() string            { return this.key.StrKey() }
func (this *Bucket) IndexKey() int             { return this.key.IndexKey() }
func (this *Bucket) Keys() (int, string, bool) { return this.key.Keys() }

func (this *Bucket) SetStrKey(key string)  { this.key = MakeStrKey(key) }
func (this *Bucket) SetIndexKey(index int) { this.key = MakeIndexKey(index) }
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
	iteratorsCount  ZendUchar
	elementsCount   uint32
	internalPointer uint32
	nextFreeElement int
	destructor      DtorFuncT

	data     []Bucket          // 实际存储数据的地方
	indexMap map[int]uint32    // 数字索引到具体位置的映射
	keyMap   map[string]uint32 // 字符串索引到具体位置的映射

	arData *Bucket // C 源码中存储数据的地方，实际不使用
	//u struct /* union */ {
	//	v struct {
	//		flags           ZendUchar
	//		_unused         ZendUchar
	//		nIteratorsCount ZendUchar
	//		_unused2        ZendUchar
	//	}
	//	flags uint32
	//}
}

var _ IRefcounted = &Array{}

/**
 * Constructor && Init
 */
func NewArray(size int) *Array {
	return NewArrayEx(size, zend.ZVAL_PTR_DTOR, false)
}
func MakeArrayEx(nSize int, pDestructor DtorFuncT, persistent ZendBool) Array {
	return *NewArrayEx(nSize, pDestructor, persistent != 0)
}
func NewArrayEx(size int, pDestructor DtorFuncT, persistent bool) *Array {
	var data []Bucket
	if size > 0 {
		data = make([]Bucket, 0, size)
	}

	var ht = &Array{
		destructor: pDestructor,

		// 数据存储
		data:     data,
		indexMap: make(map[int]uint32),    // todo 改为 nil，延迟初始化
		keyMap:   make(map[string]uint32), // todo 改为 nil，延迟初始化
	}

	// GC 信息
	ht.SetRefcount(1)
	ht.SetGcTypeInfo(IS_ARRAY)
	if persistent {
		ht.AddGcFlags(GC_PERSISTENT)
	} else {
		ht.AddGcFlags(GC_COLLECTABLE)
	}

	return ht
}

/* init */
func (ht *Array) SetBy(arr *Array) {
	ht.flags = arr.flags
	ht.iteratorsCount = arr.iteratorsCount
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
	ht.addHash(bucket.key, idx)

	if !bucket.IsStrKey() {
		var indexKey = bucket.IndexKey()
		// 更新 nextFreeElement
		if indexKey > ht.nextFreeElement {
			if indexKey < zend.ZEND_LONG_MAX {
				ht.nextFreeElement = indexKey + 1
			} else {
				ht.nextFreeElement = zend.ZEND_LONG_MAX
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
			faults.ErrorNoreturn(faults.E_ERROR, "Possible integer overflow in memory allocation (%d)", dataSize*2)
		}
	}
}
func (ht *Array) deleteBucket(pos uint32) {
	ht.assertRc1()
	b.Assert(int(pos) < len(ht.data))

	var p = &ht.data[pos]
	b.Assert(p.IsValid())

	// 移除映射
	ht.deleteHash(p.key)

	// 减少有效元素
	ht.elementsCount--

	// 更新内部指针和遍历器指针
	if ht.internalPointer == pos || ht.HasIterators() {
		var newIdx = ht.validPosVal(pos + 1)
		if ht.internalPointer == pos {
			ht.internalPointer = newIdx
		}
		ZendHashIteratorsUpdate(ht, pos, newIdx)
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
	if ht.DataSize()-1 == pos {
		ht.removeInvalidTail()
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
	b.Assert(newPos <= pos)
	if newPos == pos {
		return
	}
	(&ht.data[newPos]).CopyFrom(&ht.data[pos])
	if ht.internalPointer == pos {
		ht.internalPointer = newPos
	}
}

// todo 一般情况无需主动扩展
func (ht *Array) Extend(size uint32) {
	ht.assertRc1()
	if size > uint32(len(ht.data)) {
		// 扩展数组 cap
		newData := make([]Bucket, 0, size)
		if len(ht.data) > 0 {
			copy(newData, ht.data)
		}
		ht.data = newData
	}
}

// todo 一般情况无需主动扩展,确认使用目前是缩减内存占用还是裁剪元素
func (ht *Array) Discard(nNumUsed uint32) {
	if nNumUsed < ht.DataSize() {
		// 裁剪数据，重新映射
		ht.data = ht.data[:nNumUsed]
		ht.Rehash()
		// rehash 清理了所有 holes, 此时元素数就是 ht.data 长度
		ht.elementsCount = ht.DataSize()
	}
}

/* hash -> Array.indexMap & Array.keyMap */
func (ht *Array) resetHash() {
	ht.assertRc1()
	ht.indexMap = make(map[int]uint32)
	ht.keyMap = make(map[string]uint32)
}
func (ht *Array) addHash(key ArrayKey, pos uint32) {
	if key.IsStrKey() {
		ht.keyMap[key.StrKey()] = pos
	} else {
		ht.indexMap[key.IndexKey()] = pos
	}
}
func (ht *Array) deleteHash(key ArrayKey) {
	if key.IsStrKey() {
		delete(ht.keyMap, key.StrKey())
	} else {
		delete(ht.indexMap, key.IndexKey())
	}
}
func (ht *Array) Rehash() {
	// 空数组快速清空
	if ht.elementsCount == 0 {
		ht.resetHash()
		ht.data = nil
		return
	}

	// 移除 data 中的空位
	var oldNumUsed = ht.GetNNumUsed()
	ht.removeHoles()

	// 重建 hash
	ht.resetHash()
	ht.eachBucket(func(pos uint32, p *Bucket) {
		ht.addHash(p.key, pos)
	})

	/* Migrate pointer to one past the end of the array to the new one past the end, so that
	 * newly inserted elements are picked up correctly. */
	if ht.HasIterators() {
		_zendHashIteratorsUpdate(ht, oldNumUsed, ht.GetNNumUsed())
	}
}

/* misc */
func (ht *Array) assertRc1() { b.Assert(ht.GetRefcount() == 1) }

/** Array.flags */
func (ht *Array) CopyFlags(arr *Array) { ht.flags = arr.flags }

func (ht *Array) IsPacked() bool       { return ht.flags&HASH_FLAG_PACKED != 0 }
func (ht *Array) MarkPacked()          { ht.flags |= HASH_FLAG_PACKED }
func (ht *Array) UnmarkIsPacked()      { ht.flags &^= HASH_FLAG_PACKED }
func (ht *Array) HasEmptyIndex() bool  { return ht.flags&HASH_FLAG_HAS_EMPTY_IND != 0 }
func (ht *Array) MarkHasEmptyIndex()   { ht.flags |= HASH_FLAG_HAS_EMPTY_IND }
func (ht *Array) UnmarkHasEmptyIndex() { ht.flags &^= HASH_FLAG_HAS_EMPTY_IND }

/** Array.iteratorsCount */
func (ht *Array) GetIteratorsCount() ZendUchar      { return ht.iteratorsCount }
func (ht *Array) SetIteratorsCount(value ZendUchar) { ht.iteratorsCount = value }
func (ht *Array) IncIteratorsCount()                { ht.iteratorsCount++ }
func (ht *Array) DecIteratorsCount()                { ht.iteratorsCount-- }
func (ht *Array) HasIterators() bool                { return ht.iteratorsCount != 0 }
func (ht *Array) IsIteratorsOverflow() bool         { return ht.iteratorsCount == 0xff }
