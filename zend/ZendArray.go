// <<generate>>

package zend

import b "sik/builtin"

/**
 * ZendHashKey
 */
type ZendHashKey struct {
	h   ZendUlong
	key *ZendString
}

func (this *ZendHashKey) GetH() ZendUlong          { return this.h }
func (this *ZendHashKey) SetH(value ZendUlong)     { this.h = value }
func (this *ZendHashKey) GetKey() *ZendString      { return this.key }
func (this *ZendHashKey) SetKey(value *ZendString) { this.key = value }

/**
 * ZendArrayKey
 * 新增类型，表示 ZendArray 的 Key。与原类型 ZendHashKey 作用类似，后续会取代 ZendHashKey。
 */
type ZendArrayKey struct {
	index int
	key   *string
}

func NewStrKey(str string) ZendArrayKey  { return ZendArrayKey{0, &str} }
func NewIndexKey(index int) ZendArrayKey { return ZendArrayKey{index, nil} }
func (this ZendArrayKey) GetIndex() int  { return this.index }
func (this ZendArrayKey) GetKey() string { return *this.key }
func (this ZendArrayKey) IsStrKey() bool { return this.key != nil }
func (this ZendArrayKey) GetH() ZendUlong {
	// todo remove
	if this.key != nil {
		return b.HashStr(*this.key)
	} else {
		return uint(this.index)
	}
}
func (this ZendArrayKey) GetZendStringKey() *ZendString {
	// todo remove
	if this.key != nil {
		return ZendStringNew(*this.key, false)
	} else {
		return nil
	}
}

/**
 * Bucket
 */
type Bucket struct {
	val Zval
	key ZendArrayKey
	// todo remove
	h ZendUlong
}

func NewBucket(key ZendArrayKey, zval *Zval) *Bucket {
	var bucket = &Bucket{key: key}
	ZVAL_COPY_VALUE(&bucket.val, zval)
	return bucket
}

func NewBucketStr(strKey string, zval *Zval) *Bucket {
	var key = NewStrKey(strKey)
	return NewBucket(key, zval)
}

func NewBucketIndex(indexKey int, zval *Zval) *Bucket {
	var key = NewIndexKey(indexKey)
	return NewBucket(key, zval)
}

func (this *Bucket) GetVal() *Zval       { return &this.val }
func (this *Bucket) SetVal(zval *Zval)   { ZVAL_COPY_VALUE(&this.val, zval) }
func (this *Bucket) GetH() ZendUlong     { return this.key.GetH() }
func (this *Bucket) GetKey() *ZendString { return this.key.GetZendStringKey() }

func (this *Bucket) GetStrKey() string { return this.key.GetKey() }
func (this *Bucket) GetIndexKey() int  { return this.key.GetIndex() }

func (this *Bucket) SetH(value ZendUlong) {
	// todo remove
	this.h = value
}
func (this *Bucket) SetKey(value *ZendString) {
	// todo 此方法应被替换
	ZEND_ASSERT(false)
}

/**
 * ZendArray
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
type HashTable = ZendArray
type ZendArray struct {
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
	nNumUsed         uint32
	nNumOfElements   uint32
	nTableSize       uint32
	nInternalPointer uint32
	nNextFreeElement ZendLong
	pDestructor      DtorFuncT

	//
	arData *Bucket // C 源码中存储数据的地方，实际不使用

	data     []Bucket          // 实际存储数据的地方
	indexMap map[int]uint32    // 数字索引到具体位置的映射
	keyMap   map[string]uint32 // 字符串索引到具体位置的映射
}

var _ IRefcounted = &ZendArray{}

func (this *ZendArray) GetArData() *Bucket                 { return this.arData }
func (this *ZendArray) SetArData(value *Bucket)            { this.arData = value }
func (this *ZendArray) GetNNumUsed() uint32                { return this.nNumUsed }
func (this *ZendArray) SetNNumUsed(value uint32)           { this.nNumUsed = value }
func (this *ZendArray) GetNNumOfElements() uint32          { return this.nNumOfElements }
func (this *ZendArray) SetNNumOfElements(value uint32)     { this.nNumOfElements = value }
func (this *ZendArray) GetNTableSize() uint32              { return this.nTableSize }
func (this *ZendArray) SetNTableSize(value uint32)         { this.nTableSize = value }
func (this *ZendArray) GetNInternalPointer() uint32        { return this.nInternalPointer }
func (this *ZendArray) SetNInternalPointer(value uint32)   { this.nInternalPointer = value }
func (this *ZendArray) GetNNextFreeElement() ZendLong      { return this.nNextFreeElement }
func (this *ZendArray) SetNNextFreeElement(value ZendLong) { this.nNextFreeElement = value }
func (this *ZendArray) GetPDestructor() DtorFuncT          { return this.pDestructor }
func (this *ZendArray) SetPDestructor(value DtorFuncT)     { this.pDestructor = value }

func (this *ZendArray) GetNTableMask() uint32 {
	return HT_SIZE_TO_MASK(this.nTableSize)
}
func (this *ZendArray) SetNTableMask(value uint32) {
	ZEND_ASSERT(this.GetNTableMask() == value)
}

/* ZendArray.u.v.flags */
func (this *ZendArray) GetFlags() ZendUchar           { return this.u.v.flags }
func (this *ZendArray) SetFlags(value ZendUchar)      { this.u.v.flags = value }
func (this *ZendArray) AddFlags(value ZendUchar)      { this.u.v.flags |= value }
func (this *ZendArray) SubFlags(value ZendUchar)      { this.u.v.flags &^= value }
func (this *ZendArray) HasFlags(value ZendUchar) bool { return this.u.v.flags&value != 0 }
func (this *ZendArray) SwitchFlags(value ZendUchar, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}

const HASH_FLAG_PACKED = 1 << 2
const HASH_FLAG_STATIC_KEYS = 1 << 4
const HASH_FLAG_HAS_EMPTY_IND = 1 << 5

func (this *ZendArray) IsStaticKeys() bool  { return this.HasFlags(HASH_FLAG_STATIC_KEYS) }
func (this *ZendArray) IsHasEmptyInd() bool { return this.HasFlags(HASH_FLAG_HAS_EMPTY_IND) }
func (this *ZendArray) SetIsStaticKeys()    { this.AddFlags(HASH_FLAG_STATIC_KEYS) }
func (this *ZendArray) SetIsHasEmptyInd()   { this.AddFlags(HASH_FLAG_HAS_EMPTY_IND) }

/* ZendArray.u.flags */
func (this *ZendArray) GetUFlags() uint32           { return this.u.flags }
func (this *ZendArray) SetUFlags(value uint32)      { this.u.flags = value }
func (this *ZendArray) AddUFlags(value uint32)      { this.u.flags |= value }
func (this *ZendArray) SubUFlags(value uint32)      { this.u.flags &^= value }
func (this *ZendArray) HasUFlags(value uint32) bool { return this.u.flags&value != 0 }
func (this *ZendArray) SwitchUFlags(value uint32, cond bool) {
	if cond {
		this.AddUFlags(value)
	} else {
		this.SubUFlags(value)
	}
}

// nIteratorsCount
func (this *ZendArray) GetNIteratorsCount() ZendUchar      { return this.u.v.nIteratorsCount }
func (this *ZendArray) SetNIteratorsCount(value ZendUchar) { this.u.v.nIteratorsCount = value }
func (this *ZendArray) IncNIteratorsCount()                { this.u.v.nIteratorsCount++ }
func (this *ZendArray) DecNIteratorsCount()                { this.u.v.nIteratorsCount-- }

func (this *ZendArray) HasIterators() bool        { return this.GetNIteratorsCount() != 0 }
func (this *ZendArray) IsIteratorsOverflow() bool { return this.GetNIteratorsCount() == 0xff }

/**
 * Constructor && Init
 */
func NewZendArray(size uint32) *ZendArray {
	return NewZendArrayEx(size, ZVAL_PTR_DTOR, false)
}

func NewZendArrayEx(size uint32, pDestructor DtorFuncT, persistent bool) *ZendArray {
	var ht = &ZendArray{
		nNumUsed:         0,
		nNumOfElements:   0,
		nTableSize:       ZendHashCheckSize(size),
		nInternalPointer: 0,
		nNextFreeElement: 0,
		pDestructor:      pDestructor,

		// 数据存储
		data:     nil,
		indexMap: make(map[int]uint32),
		keyMap:   make(map[string]uint32),
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

func (this *ZendArray) assertRc1() {
	ZEND_ASSERT(this.GetRefcount() == 1)
}

func (this *ZendArray) RealInit() {
	this.assertRc1()

	this.nNumUsed = 0
	this.nNumOfElements = 0
	this.data = nil
	this.indexMap = make(map[int]uint32)
	this.keyMap = make(map[string]uint32)

	this.SetIsStaticKeys()
}

func (this *ZendArray) hashReset() {
	this.assertRc1()
	this.indexMap = make(map[int]uint32)
	this.keyMap = make(map[string]uint32)
}
