package types

import (
	"sik/zend"
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
func (this ArrayKey) KeyKey() string { return *this.key }

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

func (this *Bucket) IsStrKey() bool        { return this.key.IsStrKey() }
func (this *Bucket) StrKey() string        { return this.key.KeyKey() }
func (this *Bucket) IndexKey() int         { return this.key.IndexKey() }
func (this *Bucket) SetStrKey(key string)  { this.key = MakeStrKey(key) }
func (this *Bucket) SetIndexKey(index int) { this.key = MakeIndexKey(index) }
func (this *Bucket) GetArrayKey() ArrayKey { return this.key }

func (this *Bucket) GetVal() *Zval     { return &this.val }
func (this *Bucket) SetVal(zval *Zval) { ZVAL_COPY_VALUE(&this.val, zval) }

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
type HashPosition = uint32
type Array struct {
	ZendRefcounted
	flags           ZendUchar
	iteratorsCount  ZendUchar
	elementsCount   uint32
	tableSize       uint32
	internalPointer uint32
	nextFreeElement int
	destructor      DtorFuncT

	data []Bucket // 实际存储数据的地方

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

/**
 * Constructor && Init
 */
func NewArray(size int) *Array {
	return NewArrayEx(size, zend.ZVAL_PTR_DTOR, false)
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

func (ht *Array) CopyFrom(arr *Array) {
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

func (ht *Array) Cap() int { return cap(ht.data) }

/** Array.flags */
func (ht *Array) CopyFlags(arr *Array) { ht.flags = arr.flags }

func (ht *Array) IsPacked() bool      { return ht.flags&HASH_FLAG_PACKED != 0 }
func (ht *Array) SetIsPacked()        { ht.flags |= HASH_FLAG_PACKED }
func (ht *Array) UnsetIsPacked()      { ht.flags &^= HASH_FLAG_PACKED }
func (ht *Array) IsHasEmptyInd() bool { return ht.flags&HASH_FLAG_HAS_EMPTY_IND != 0 }
func (ht *Array) SetIsHasEmptyInd()   { ht.flags |= HASH_FLAG_HAS_EMPTY_IND }
func (ht *Array) UnsetIsHasEmptyInd() { ht.flags &^= HASH_FLAG_HAS_EMPTY_IND }

/** Array.iteratorsCount */
func (ht *Array) GetIteratorsCount() ZendUchar      { return ht.iteratorsCount }
func (ht *Array) SetIteratorsCount(value ZendUchar) { ht.iteratorsCount = value }
func (ht *Array) IncIteratorsCount()                { ht.iteratorsCount++ }
func (ht *Array) DecIteratorsCount()                { ht.iteratorsCount-- }
func (ht *Array) HasIterators() bool                { return ht.iteratorsCount != 0 }
func (ht *Array) IsIteratorsOverflow() bool         { return ht.iteratorsCount == 0xff }
