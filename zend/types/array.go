package types

import "sik/zend"

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
 *                 | Bucket[ht->nTableSize-1]    |
 *                 +=============================+
 */
type Array struct {
	ZendRefcounted
	u struct /* union */ {
		v struct {
			flags           ZendUchar
			_unused         ZendUchar
			nIteratorsCount ZendUchar
			_unused2        ZendUchar
		}
		flags uint32
	}
	nNumOfElements   uint32
	nTableSize       uint32
	nInternalPointer uint32
	nNextFreeElement zend.ZendLong
	pDestructor      DtorFuncT

	//
	arData *Bucket // C 源码中存储数据的地方，实际不使用

	data     []Bucket          // 实际存储数据的地方
	indexMap map[int]uint32    // 数字索引到具体位置的映射
	keyMap   map[string]uint32 // 字符串索引到具体位置的映射
}
