package types

import (
	b "sik/builtin"
	"sik/zend"
)

/**
 * ArrayKey
 * 新增类型，表示 Array 的 Key。与原类型 ArrayKey 作用类似，后续会取代 ArrayKey。
 */
type ArrayKey struct {
	index int
	key   *string
}

func NewStrKey(str string) ArrayKey  { return ArrayKey{0, &str} }
func NewIndexKey(index int) ArrayKey { return ArrayKey{index, nil} }
func (this ArrayKey) GetIndex() int  { return this.index }
func (this ArrayKey) GetKey() string { return *this.key }
func (this ArrayKey) IsStrKey() bool { return this.key != nil }
func (this ArrayKey) GetH() zend.ZendUlong {
	// todo remove
	if this.key != nil {
		return b.HashStr(*this.key)
	} else {
		return uint(this.index)
	}
}
func (this ArrayKey) GetZendStringKey() *String {
	// todo remove
	if this.key != nil {
		return NewString(*this.key)
	} else {
		return nil
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

func NewBucketStr(strKey string, zval *Zval) *Bucket {
	var key = NewStrKey(strKey)
	return NewBucket(key, zval)
}

func NewBucketIndex(indexKey int, zval *Zval) *Bucket {
	var key = NewIndexKey(indexKey)
	return NewBucket(key, zval)
}

func NewBucketPtr(key ArrayKey, ptr any) *Bucket {
	var bucket = &Bucket{key: key}
	ZVAL_PTR(&bucket.val, ptr)
	return bucket
}

func NewBucketIndirect(key ArrayKey, ptr *Zval) *Bucket {
	var bucket = &Bucket{key: key}
	ZVAL_INDIRECT(&bucket.val, ptr)
	return bucket
}

func (this *Bucket) GetVal() *Zval           { return &this.val }
func (this *Bucket) SetVal(zval *Zval)       { ZVAL_COPY_VALUE(&this.val, zval) }
func (this *Bucket) GetZendKey() ArrayKey    { return this.key }
func (this *Bucket) SetZendKey(key ArrayKey) { this.key = key }

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

func (this *Bucket) GetH() zend.ZendUlong { return this.key.GetH() }
func (this *Bucket) GetKey() *String      { return this.key.GetZendStringKey() }
func (this *Bucket) SetH(value zend.ZendUlong) {
	// todo remove
	b.Assert(false)
}
func (this *Bucket) SetKey(value *String) {
	// todo 此方法应被替换
	b.Assert(false)
}

func (this *Bucket) CopyFrom(from *Bucket) {
	this.SetVal(from.GetVal())
	this.key = from.key
}

func (this *Bucket) IsValid() bool {
	return !this.val.IsUndef()
}

func (this *Bucket) SetInvalid() {
	this.val.SetUndef()
}

var _ IRefcounted = &Array{}

func (this *Array) GetArData() *Bucket      { return this.arData }
func (this *Array) SetArData(value *Bucket) { this.arData = value }

func (this *Array) DataSize() uint32 { return uint32(len(this.data)) }
func (this *Array) LastPos() uint32  { return this.DataSize() - 1 }

func (this *Array) GetNNumUsed() uint32 { return this.DataSize() }
func (this *Array) SetNNumUsed(value uint32) {
	// todo remove
}

func (this *Array) GetNNumOfElements() uint32               { return this.nNumOfElements }
func (this *Array) SetNNumOfElements(value uint32)          { this.nNumOfElements = value }
func (this *Array) GetNTableSize() uint32                   { return this.nTableSize }
func (this *Array) SetNTableSize(value uint32)              { this.nTableSize = value }
func (this *Array) GetNInternalPointer() uint32             { return this.nInternalPointer }
func (this *Array) SetNInternalPointer(value uint32)        { this.nInternalPointer = value }
func (this *Array) GetNNextFreeElement() zend.ZendLong      { return this.nNextFreeElement }
func (this *Array) SetNNextFreeElement(value zend.ZendLong) { this.nNextFreeElement = value }
func (this *Array) GetPDestructor() DtorFuncT               { return this.pDestructor }
func (this *Array) SetPDestructor(value DtorFuncT)          { this.pDestructor = value }

func (this *Array) GetNTableMask() uint32 {
	// todo 待移除
	return uint32(-(this.nTableSize + this.nTableSize))
}
func (this *Array) SetNTableMask(value uint32) {
	b.Assert(this.GetNTableMask() == value)
}

/* Array.u.v.flags */
func (this *Array) GetFlags() ZendUchar           { return this.u.v.flags }
func (this *Array) SetFlags(value ZendUchar)      { this.u.v.flags = value }
func (this *Array) AddFlags(value ZendUchar)      { this.u.v.flags |= value }
func (this *Array) SubFlags(value ZendUchar)      { this.u.v.flags &^= value }
func (this *Array) HasFlags(value ZendUchar) bool { return this.u.v.flags&value != 0 }
func (this *Array) SwitchFlags(value ZendUchar, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}

const HASH_FLAG_PACKED = 1 << 2
const HASH_FLAG_HAS_EMPTY_IND = 1 << 5

func (this *Array) IsHasEmptyInd() bool { return this.HasFlags(HASH_FLAG_HAS_EMPTY_IND) }
func (this *Array) SetIsHasEmptyInd()   { this.AddFlags(HASH_FLAG_HAS_EMPTY_IND) }

/* Array.u.flags */
func (this *Array) GetUFlags() uint32           { return this.u.flags }
func (this *Array) SetUFlags(value uint32)      { this.u.flags = value }
func (this *Array) AddUFlags(value uint32)      { this.u.flags |= value }
func (this *Array) SubUFlags(value uint32)      { this.u.flags &^= value }
func (this *Array) HasUFlags(value uint32) bool { return this.u.flags&value != 0 }
func (this *Array) SwitchUFlags(value uint32, cond bool) {
	if cond {
		this.AddUFlags(value)
	} else {
		this.SubUFlags(value)
	}
}

// nIteratorsCount
func (this *Array) GetNIteratorsCount() ZendUchar      { return this.u.v.nIteratorsCount }
func (this *Array) SetNIteratorsCount(value ZendUchar) { this.u.v.nIteratorsCount = value }
func (this *Array) IncNIteratorsCount()                { this.u.v.nIteratorsCount++ }
func (this *Array) DecNIteratorsCount()                { this.u.v.nIteratorsCount-- }

func (this *Array) HasIterators() bool        { return this.GetNIteratorsCount() != 0 }
func (this *Array) IsIteratorsOverflow() bool { return this.GetNIteratorsCount() == 0xff }

/**
 * Constructor && Init
 */
func NewZendArray(size int) *Array {
	return NewZendArrayEx(size, zend.ZVAL_PTR_DTOR, false)
}

func NewZendArrayEx(size uint32, pDestructor DtorFuncT, persistent bool) *Array {
	var ht = &Array{
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

func (this *Array) assertRc1() {
	b.Assert(this.GetRefcount() == 1)
}

func (this *Array) resetDataAndHash(dataSize uint32) {
	this.data = make([]Bucket, dataSize)
	this.indexMap = make(map[int]uint32)
	this.keyMap = make(map[string]uint32)
}

func (this *Array) copyDataAndHash(source *Array) {
	this.data = make([]Bucket, len(source.data))
	copy(this.data, source.data)

	this.indexMap = make(map[int]uint32)
	for i, pos := range source.indexMap {
		this.indexMap[i] = pos
	}

	this.keyMap = make(map[string]uint32)
	for i, pos := range source.keyMap {
		this.keyMap[i] = pos
	}
}

func (this *Array) clearData() {
	this.assertRc1()

	this.nNumOfElements = 0
	this.data = nil
	this.indexMap = make(map[int]uint32)
	this.keyMap = make(map[string]uint32)
	this.nNextFreeElement = 0
	this.nInternalPointer = 0
}

func (this *Array) RealInit() {
	this.clearData()
}

func (this *Array) resetHash() {
	this.assertRc1()
	this.indexMap = make(map[int]uint32)
	this.keyMap = make(map[string]uint32)
}

/**
 * Bucket 相关读接口
 */
func (this *Array) Bucket(pos uint32) *Bucket { return &this.data[pos] }

func (this *Array) IndexFindBucket(index int) *Bucket {
	if pos, ok := this.indexMap[index]; ok {
		return &this.data[pos]
	}
	return nil
}

func (this *Array) KeyFindBucket(key string) *Bucket {
	if pos, ok := this.keyMap[key]; ok {
		return &this.data[pos]
	}
	return nil
}

func (this *Array) IndexFindH(h zend.ZendUlong) *Zval {
	return this.IndexFind(int(h))
}
func (this *Array) IndexFind(index int) *Zval {
	var p = this.IndexFindBucket(index)
	if p != nil {
		return p.GetVal()
	}
	return nil
}

func (this *Array) KeyFind(key string) *Zval {
	var p = this.KeyFindBucket(key)
	if p != nil {
		return p.GetVal()
	}
	return nil
}

func (this *Array) KeyFindPtr(key string) any {
	var zv = this.KeyFind(key)
	if zv != nil {
		return zv.GetPtr()
	}
	return nil
}

func (this *Array) IndexExists(index int) bool {
	if _, ok := this.indexMap[index]; ok {
		return true
	}
	return false
}

func (this *Array) KeyExists(key string) bool {
	if _, ok := this.keyMap[key]; ok {
		return true
	}
	return false
}

func (this *Array) KeyExistsInd(key string) bool {
	var zv = this.KeyFind(key)
	if zv == nil {
		return false
	}

	if zv.IsUndef() && zv.GetZv().IsUndef() {
		return false
	}

	return true
}

/**
 * Add / Update by IndexKey
 */

// IndexAdd
func (this *Array) IndexAddH(h zend.ZendUlong, pData *Zval) *Zval {
	return this.IndexAdd(int(h), pData)
}
func (this *Array) IndexAdd(index int, pData *Zval) *Zval {
	this.assertRc1()

	if this.IndexExists(index) {
		return nil
	}

	var p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}

// IndexAddNew
func (this *Array) IndexAddNewH(h zend.ZendUlong, pData *Zval) *Zval {
	return this.IndexAddNew(int(h), pData)
}
func (this *Array) IndexAddNew(index int, pData *Zval) *Zval {
	this.assertRc1()

	var p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}

// IndexUpdate
func (this *Array) IndexUpdateH(h zend.ZendUlong, pData *Zval) *Zval {
	return this.IndexUpdate(int(h), pData)
}
func (this *Array) IndexUpdate(index int, pData *Zval) *Zval {
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
func (this *Array) NextIndexInsert(pData *Zval) *Zval {
	this.assertRc1()

	var index = this.nNextFreeElement

	if this.IndexExists(index) {
		return nil
	}

	var p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}

// NextIndexInsertNew
func (this *Array) NextIndexInsertNew(pData *Zval) *Zval {
	this.assertRc1()

	var index = this.nNextFreeElement
	var p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}

/**
 * Add / Update by StringKey
 */

// KeyAdd
func (this *Array) KeyAdd(key string, pData *Zval) *Zval {
	this.assertRc1()
	if this.KeyExists(key) {
		return nil
	}

	var p = this.appendBucketStr(key, pData)
	return p.GetVal()
}

// KeyAddNew
func (this *Array) KeyAddNew(key string, pData *Zval) *Zval {
	this.assertRc1()

	var p = this.appendBucketStr(key, pData)
	return p.GetVal()
}

func (this *Array) KeyAddIndirect(strKey string, pData *Zval) *Zval {
	this.assertRc1()

	var p = this.KeyFindBucket(strKey)
	if p != nil {
		var data *Zval
		b.Assert(p.GetVal() != pData)
		data = p.GetVal()
		if data.IsIndirect() {
			data = data.GetZv()
			if !data.IsUndef() {
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

	p = this.appendBucketStr(strKey, pData)
	return p.GetVal()
}

// KeyUpdate
func (this *Array) KeyUpdate(key string, pData *Zval) *Zval {
	this.assertRc1()

	var p = this.KeyFindBucket(key)
	if p != nil {
		var data *Zval
		b.Assert(p.GetVal() != pData)
		data = p.GetVal()
		if this.GetPDestructor() != nil {
			this.GetPDestructor()(data)
		}
		ZVAL_COPY_VALUE(data, pData)
		return data
	}

	p = this.appendBucketStr(key, pData)
	return p.GetVal()
}

// KeyUpdateIndirect
func (this *Array) KeyUpdateIndirect(key string, pData *Zval) *Zval {
	this.assertRc1()

	var p = this.KeyFindBucket(key)
	if p != nil {
		var data *Zval
		b.Assert(p.GetVal() != pData)
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

	p = this.appendBucketStr(key, pData)
	return p.GetVal()
}

/**
 * Add / Update by ArrayKey
 */
func (this *Array) Add(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return this.KeyAdd(key.GetKey(), pData)
	} else {
		return this.IndexAdd(key.GetIndex(), pData)
	}
}

func (this *Array) AddIndirect(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return this.KeyAddIndirect(key.GetKey(), pData)
	} else {
		return this.IndexAdd(key.GetIndex(), pData)
	}
}

func (this *Array) Update(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return this.KeyUpdate(key.GetKey(), pData)
	} else {
		return this.IndexUpdate(key.GetIndex(), pData)
	}
}

func (this *Array) UpdateIndirect(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return this.KeyUpdateIndirect(key.GetKey(), pData)
	} else {
		return this.IndexUpdate(key.GetIndex(), pData)
	}
}

/**
 * Delete
 */
func (this *Array) KeyDelete(key string) bool {
	if idx, ok := this.keyMap[key]; ok {
		this.deleteBucket(idx)
		return true
	}
	return false
}

func (this *Array) KeyDeleteIndirect(key string) bool {
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
					data.SetUndef()
					this.GetPDestructor()(&tmp)
				} else {
					data.SetUndef()
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

func (this *Array) IndexDelete(index int) bool {
	if idx, ok := this.indexMap[index]; ok {
		this.deleteBucket(idx)
		return true
	}
	return false
}

func (this *Array) deleteBucket(pos uint32) {
	this.assertRc1()
	b.Assert(pos < this.DataSize())

	var p = &this.data[pos]
	b.Assert(p.IsValid())

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
func (this *Array) Clean() {
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

func (this *Array) Destroy() {
	if this.DataSize() != 0 {
		if this.pDestructor != nil {
			this.eachValidBucket(func(pos uint32, p *Bucket) {
				this.pDestructor(p.GetVal())
			})
		}
	}
	ZendHashIteratorsRemove(this)
}

func (this *Array) DestroyEx() {
	/* break possible cycles */
	//GC_REMOVE_FROM_BUFFER(this)
	this.SetGcTypeInfo(IS_NULL)

	this.Destroy()
}

func (this *Array) GracefulReverseDestroy() {
	this.assertRc1()
	for idx := this.DataSize(); idx > 0; idx-- {
		var p = &this.data[idx-1]
		if p.IsValid() {
			this.deleteBucket(idx)
		}
	}
}
