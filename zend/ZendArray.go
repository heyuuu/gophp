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

func (this ZendArrayKey) GetZendHashKey() ZendHashKey {
	return ZendHashKey{key: this.GetZendStringKey(), h: this.GetH()}
}

/**
 * Bucket
 */
type Bucket struct {
	val Zval
	key ZendArrayKey
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

func NewBucketPtr(key ZendArrayKey, ptr any) *Bucket {
	var bucket = &Bucket{key: key}
	ZVAL_PTR(&bucket.val, ptr)
	return bucket
}

func NewBucketIndirect(key ZendArrayKey, ptr *Zval) *Bucket {
	var bucket = &Bucket{key: key}
	ZVAL_INDIRECT(&bucket.val, ptr)
	return bucket
}

func (this *Bucket) GetVal() *Zval       { return &this.val }
func (this *Bucket) SetVal(zval *Zval)   { ZVAL_COPY_VALUE(&this.val, zval) }
func (this *Bucket) GetH() ZendUlong     { return this.key.GetH() }
func (this *Bucket) GetKey() *ZendString { return this.key.GetZendStringKey() }

func (this *Bucket) IsStrKey() bool   { return this.key.IsStrKey() }
func (this *Bucket) IsIndexKey() bool { return !this.key.IsStrKey() }
func (this *Bucket) StrKey() string   { return this.key.GetKey() }
func (this *Bucket) IndexKey() int    { return this.key.GetIndex() }
func (this *Bucket) SetStrKey(key string) {
	this.key = NewStrKey(key)
}
func (this *Bucket) SetIndexKey(index int) {
	this.key = NewIndexKey(index)
}

func (this *Bucket) SetH(value ZendUlong) {
	// todo remove
	ZEND_ASSERT(false)
}
func (this *Bucket) SetKey(value *ZendString) {
	// todo 此方法应被替换
	ZEND_ASSERT(false)
}

func (this *Bucket) CopyFrom(from *Bucket) {
	ZVAL_COPY_VALUE(this.GetVal(), from.GetVal())
	this.key = from.key
}

func (this *Bucket) IsValid() bool {
	return this.val.IsType(IS_UNDEF)
}

func (this *Bucket) SetInvalid() {
	this.val.SetTypeInfo(IS_UNDEF)
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

func (this *ZendArray) GetArData() *Bucket      { return this.arData }
func (this *ZendArray) SetArData(value *Bucket) { this.arData = value }

func (this *ZendArray) DataSize() uint32 { return uint32(len(this.data)) }

func (this *ZendArray) GetNNumUsed() uint32 { return this.DataSize() }
func (this *ZendArray) SetNNumUsed(value uint32) {
	// todo remove
}

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

func (this *ZendArray) clearData() {
	this.assertRc1()

	this.nNumOfElements = 0
	this.data = nil
	this.indexMap = make(map[int]uint32)
	this.keyMap = make(map[string]uint32)
	this.nNextFreeElement = 0
	this.nInternalPointer = 0

	// todo 确认有多少 flags
	this.SetIsStaticKeys()
}

func (this *ZendArray) RealInit() {
	this.clearData()
}

func (this *ZendArray) resetHash() {
	this.assertRc1()
	this.indexMap = make(map[int]uint32)
	this.keyMap = make(map[string]uint32)
}

/**
 * Bucket 相关读接口
 */
func (this *ZendArray) Bucket(pos uint32) *Bucket { return &this.data[pos] }

func (this *ZendArray) IndexFindBucket(index int) *Bucket {
	if pos, ok := this.indexMap[index]; ok {
		return &this.data[pos]
	}
	return nil
}

func (this *ZendArray) KeyFindBucket(key string) *Bucket {
	if pos, ok := this.keyMap[key]; ok {
		return &this.data[pos]
	}
	return nil
}

func (this *ZendArray) IndexFindH(h ZendUlong) *Zval {
	return this.IndexFind(int(h))
}
func (this *ZendArray) IndexFind(index int) *Zval {
	var p = this.IndexFindBucket(index)
	if p != nil {
		return p.GetVal()
	}
	return nil
}

func (this *ZendArray) KeyFind(key string) *Zval {
	var p = this.KeyFindBucket(key)
	if p != nil {
		return p.GetVal()
	}
	return nil
}

func (this *ZendArray) IndexExists(index int) bool {
	if _, ok := this.indexMap[index]; ok {
		return true
	}
	return false
}

func (this *ZendArray) KeyExists(key string) bool {
	if _, ok := this.keyMap[key]; ok {
		return true
	}
	return false
}

/**
 * Add / Update by IndexKey
 */

// IndexAdd
func (this *ZendArray) IndexAddH(h ZendUlong, pData *Zval) *Zval {
	return this.IndexAdd(int(h), pData)
}
func (this *ZendArray) IndexAdd(index int, pData *Zval) *Zval {
	this.assertRc1()

	if this.IndexExists(index) {
		return nil
	}

	var p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}

// IndexAddNew
func (this *ZendArray) IndexAddNewH(h ZendUlong, pData *Zval) *Zval {
	return this.IndexAddNew(int(h), pData)
}
func (this *ZendArray) IndexAddNew(index int, pData *Zval) *Zval {
	this.assertRc1()

	var p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}

// IndexUpdate
func (this *ZendArray) IndexUpdateH(h ZendUlong, pData *Zval) *Zval {
	return this.IndexUpdate(int(h), pData)
}
func (this *ZendArray) IndexUpdate(index int, pData *Zval) *Zval {
	this.assertRc1()

	var p *Bucket

	p = this.IndexFindBucket(index)
	if p != nil {
		if this.pDestructor != nil {
			this.pDestructor(p.GetVal())
		}
		ZVAL_COPY_VALUE(p.GetVal(), pData)
		return p.GetVal()
	}

	p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}

// NextIndexInsert
func (this *ZendArray) NextIndexInsert(pData *Zval) *Zval {
	this.assertRc1()

	var index = this.nNextFreeElement

	if this.IndexExists(index) {
		return nil
	}

	var p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}

// NextIndexInsertNew
func (this *ZendArray) NextIndexInsertNew(pData *Zval) *Zval {
	this.assertRc1()

	var index = this.nNextFreeElement
	var p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}

/**
 * Add / Update by StringKey
 */

// KeyAdd
func (this *ZendArray) KeyAdd(key string, pData *Zval) *Zval {
	this.assertRc1()
	if this.KeyExists(key) {
		return nil
	}

	this.SubUFlags(HASH_FLAG_STATIC_KEYS)
	var p = this.appendBucketStr(key, pData)
	return p.GetVal()
}

// KeyAddNew
func (this *ZendArray) KeyAddNew(key string, pData *Zval) *Zval {
	this.assertRc1()

	this.SubUFlags(HASH_FLAG_STATIC_KEYS)
	var p = this.appendBucketStr(key, pData)
	return p.GetVal()
}

func (this *ZendArray) KeyAddIndirect(strKey string, pData *Zval) *Zval {
	this.assertRc1()

	var p = this.KeyFindBucket(strKey)
	if p != nil {
		var data *Zval
		ZEND_ASSERT(p.GetVal() != pData)
		data = p.GetVal()
		if data.IsType(IS_INDIRECT) {
			data = data.GetZv()
			if data.GetType() != IS_UNDEF {
				return nil
			}
		} else {
			return nil
		}
		if this.GetPDestructor() != nil {
			this.GetPDestructor()(data)
		}
		ZVAL_COPY_VALUE(data, pData)
		return data
	}

	this.SubUFlags(HASH_FLAG_STATIC_KEYS)
	p = this.appendBucketStr(strKey, pData)
	return p.GetVal()
}

// KeyUpdate
func (this *ZendArray) KeyUpdate(key string, pData *Zval) *Zval {
	this.assertRc1()

	var p = this.KeyFindBucket(key)
	if p != nil {
		var data *Zval
		ZEND_ASSERT(p.GetVal() != pData)
		data = p.GetVal()
		if this.GetPDestructor() != nil {
			this.GetPDestructor()(data)
		}
		ZVAL_COPY_VALUE(data, pData)
		return data
	}

	this.SubUFlags(HASH_FLAG_STATIC_KEYS)
	p = this.appendBucketStr(key, pData)
	return p.GetVal()
}

// KeyUpdateIndirect
func (this *ZendArray) KeyUpdateIndirect(key string, pData *Zval) *Zval {
	this.assertRc1()

	var p = this.KeyFindBucket(key)
	if p != nil {
		var data *Zval
		ZEND_ASSERT(p.GetVal() != pData)
		data = p.GetVal()
		if data.IsType(IS_INDIRECT) {
			data = data.GetZv()
		}
		if this.GetPDestructor() != nil {
			this.GetPDestructor()(data)
		}
		ZVAL_COPY_VALUE(data, pData)
		return data
	}

	this.SubUFlags(HASH_FLAG_STATIC_KEYS)
	p = this.appendBucketStr(key, pData)
	return p.GetVal()
}

/**
 * Add / Update by ZendArrayKey
 */
func (this *HashTable) Update(key ZendArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return this.KeyUpdate(key.GetKey(), pData)
	} else {
		return this.IndexUpdate(key.GetIndex(), pData)
	}
}

/**
 * Delete
 */
func (this *HashTable) KeyDelete(key string) bool {
	if idx, ok := this.keyMap[key]; ok {
		this.deleteBucket(idx)
		return true
	}
	return false
}

func (this *HashTable) KeyDeleteIndirect(key string) bool {
	this.assertRc1()
	if idx, ok := this.keyMap[key]; ok {
		var p = &this.data[idx]
		if p.GetVal().IsType(IS_INDIRECT) {
			var data *Zval = p.GetVal().GetZv()
			if data.IsType(IS_UNDEF) {
				return false
			} else {
				if this.GetPDestructor() != nil {
					var tmp Zval
					ZVAL_COPY_VALUE(&tmp, data)
					ZVAL_UNDEF(data)
					this.GetPDestructor()(&tmp)
				} else {
					ZVAL_UNDEF(data)
				}
				this.AddUFlags(HASH_FLAG_HAS_EMPTY_IND)
			}
		} else {
			this.deleteBucket(idx)
		}
		return true
	}
	return false
}

func (this *HashTable) IndexDelete(index int) bool {
	if idx, ok := this.indexMap[index]; ok {
		this.deleteBucket(idx)
		return true
	}
	return false
}

func (this *ZendArray) deleteBucket(pos uint32) {
	this.assertRc1()
	ZEND_ASSERT(pos < this.DataSize())

	var p = &this.data[pos]
	ZEND_ASSERT(p.IsValid())

	// 移除映射
	this.deleteHash(p.key)

	// 减少有效元素
	this.nNumOfElements--

	// 更新内部指针和遍历器指针
	if this.nInternalPointer == pos || this.HasIterators() {
		var newIdx = this.validPosVal(pos + 1)
		if this.nInternalPointer == pos {
			this.nInternalPointer = newIdx
		}
		ZendHashIteratorsUpdate(this, pos, newIdx)
	}

	// 析构函数
	if this.pDestructor != nil {
		var tmp Zval
		ZVAL_COPY_VALUE(&tmp, p.GetVal())
		this.GetPDestructor()(&tmp)
	}

	// 设置数据不可用
	p.SetInvalid()

	// 若删除队尾元素，尝试清除 data 队尾无用数据
	if this.DataSize()-1 == pos {
		this.removeInvalidTail()
	}
}

/**
 * Clean && Destroy
 */
func (this *HashTable) Clean() {
	this.assertRc1()
	if this.GetNNumUsed() != 0 {
		if this.pDestructor != nil {
			this.eachValidBucket(func(pos uint32, p *Bucket) {
				this.pDestructor(p.GetVal())
			})
		}
	}
	this.clearData()
}

func (this *HashTable) Destroy() {
	if this.DataSize() != 0 {
		if this.pDestructor != nil {
			this.eachValidBucket(func(pos uint32, p *Bucket) {
				this.pDestructor(p.GetVal())
			})
		}
	}
	ZendHashIteratorsRemove(this)
}

func (this *HashTable) DestroyEx() {
	/* break possible cycles */
	GC_REMOVE_FROM_BUFFER(this)
	this.SetGcTypeInfo(IS_NULL)

	this.Destroy()
}

func (this *HashTable) GracefulReverseDestroy() {
	this.assertRc1()
	for idx := this.DataSize(); idx > 0; idx-- {
		var p = &this.data[idx-1]
		if p.IsValid() {
			this.deleteBucket(idx)
		}
	}
}
