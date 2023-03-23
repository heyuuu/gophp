package types

import (
	b "sik/builtin"
	"sik/zend"
)

func (this *Bucket) GetH() uint {
	if this.IsStrKey() {
		return b.HashStr(this.key.KeyKey())
	} else {
		return uint(this.key.index)
	}
}
func (this *Bucket) GetKey() *String {
	if this.IsStrKey() {
		return NewString(this.key.KeyKey())
	} else {
		return nil
	}
}
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

func (ht *Array) GetArData() *Bucket      { return ht.arData }
func (ht *Array) SetArData(value *Bucket) { ht.arData = value }

func (ht *Array) DataSize() uint32 { return uint32(len(ht.data)) }
func (ht *Array) LastPos() uint32  { return ht.DataSize() - 1 }

func (ht *Array) GetNNumUsed() uint32 { return ht.DataSize() }
func (ht *Array) SetNNumUsed(value uint32) {
	// todo remove
}

func (ht *Array) CountElements() uint32                   { return ht.elementsCount }
func (ht *Array) SetNNumOfElements(value uint32)          { ht.elementsCount = value }
func (ht *Array) GetNInternalPointer() uint32             { return ht.internalPointer }
func (ht *Array) SetNInternalPointer(value uint32)        { ht.internalPointer = value }
func (ht *Array) GetNNextFreeElement() zend.ZendLong      { return ht.nextFreeElement }
func (ht *Array) SetNNextFreeElement(value zend.ZendLong) { ht.nextFreeElement = value }
func (ht *Array) GetPDestructor() DtorFuncT               { return ht.destructor }
func (ht *Array) SetPDestructor(value DtorFuncT)          { ht.destructor = value }

func (ht *Array) GetNTableMask() uint32 { return 0 } // todo remove

/**
 * Constructor && Init
 */

func (ht *Array) assertRc1() {
	b.Assert(ht.GetRefcount() == 1)
}

func (ht *Array) resetDataAndHash(dataSize uint32) {
	ht.data = make([]Bucket, dataSize)
	ht.indexMap = make(map[int]uint32)
	ht.keyMap = make(map[string]uint32)
}

func (ht *Array) copyDataAndHash(source *Array) {
	ht.data = make([]Bucket, len(source.data))
	copy(ht.data, source.data)

	ht.indexMap = make(map[int]uint32)
	for i, pos := range source.indexMap {
		ht.indexMap[i] = pos
	}

	ht.keyMap = make(map[string]uint32)
	for i, pos := range source.keyMap {
		ht.keyMap[i] = pos
	}
}

func (ht *Array) clearData() {
	ht.assertRc1()

	ht.elementsCount = 0
	ht.data = nil
	ht.indexMap = make(map[int]uint32)
	ht.keyMap = make(map[string]uint32)
	ht.nextFreeElement = 0
	ht.internalPointer = 0
}

func (ht *Array) RealInit() {
	ht.clearData()
}

func (ht *Array) resetHash() {
	ht.assertRc1()
	ht.indexMap = make(map[int]uint32)
	ht.keyMap = make(map[string]uint32)
}

/**
 * Bucket 相关读接口
 */
func (ht *Array) Bucket(pos uint32) *Bucket { return &ht.data[pos] }

func (ht *Array) IndexFindBucket(index int) *Bucket {
	if pos, ok := ht.indexMap[index]; ok {
		return &ht.data[pos]
	}
	return nil
}

func (ht *Array) KeyFindBucket(key string) *Bucket {
	if pos, ok := ht.keyMap[key]; ok {
		return &ht.data[pos]
	}
	return nil
}

func (ht *Array) IndexFindH(h zend.ZendUlong) *Zval {
	return ht.IndexFind(int(h))
}
func (ht *Array) IndexFind(index int) *Zval {
	var p = ht.IndexFindBucket(index)
	if p != nil {
		return p.GetVal()
	}
	return nil
}

func (ht *Array) KeyFind(key string) *Zval {
	var p = ht.KeyFindBucket(key)
	if p != nil {
		return p.GetVal()
	}
	return nil
}

func (ht *Array) KeyFindPtr(key string) any {
	var zv = ht.KeyFind(key)
	if zv != nil {
		return zv.GetPtr()
	}
	return nil
}

func (ht *Array) IndexExists(index int) bool {
	if _, ok := ht.indexMap[index]; ok {
		return true
	}
	return false
}

func (ht *Array) KeyExists(key string) bool {
	if _, ok := ht.keyMap[key]; ok {
		return true
	}
	return false
}

func (ht *Array) KeyExistsInd(key string) bool {
	var zv = ht.KeyFind(key)
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
func (ht *Array) IndexAddH(h zend.ZendUlong, pData *Zval) *Zval {
	return ht.IndexAdd(int(h), pData)
}
func (ht *Array) IndexAdd(index int, pData *Zval) *Zval {
	ht.assertRc1()

	if ht.IndexExists(index) {
		return nil
	}

	var p = ht.appendBucketIndex(index, pData)
	return p.GetVal()
}

// IndexAddNew
func (ht *Array) IndexAddNewH(h zend.ZendUlong, pData *Zval) *Zval {
	return ht.IndexAddNew(int(h), pData)
}
func (ht *Array) IndexAddNew(index int, pData *Zval) *Zval {
	ht.assertRc1()

	var p = ht.appendBucketIndex(index, pData)
	return p.GetVal()
}

// IndexUpdate
func (ht *Array) IndexUpdateH(h zend.ZendUlong, pData *Zval) *Zval {
	return ht.IndexUpdate(int(h), pData)
}
func (ht *Array) IndexUpdate(index int, pData *Zval) *Zval {
	ht.assertRc1()

	var p *Bucket

	p = ht.IndexFindBucket(index)
	if p != nil {
		if ht.destructor != nil {
			ht.destructor(p.GetVal())
		}
		ZVAL_COPY_VALUE(p.GetVal(), pData)
		return p.GetVal()
	}

	p = ht.appendBucketIndex(index, pData)
	return p.GetVal()
}

// NextIndexInsert
func (ht *Array) NextIndexInsert(pData *Zval) *Zval {
	ht.assertRc1()

	var index = ht.nextFreeElement

	if ht.IndexExists(index) {
		return nil
	}

	var p = ht.appendBucketIndex(index, pData)
	return p.GetVal()
}

// NextIndexInsertNew
func (ht *Array) NextIndexInsertNew(pData *Zval) *Zval {
	ht.assertRc1()

	var index = ht.nextFreeElement
	var p = ht.appendBucketIndex(index, pData)
	return p.GetVal()
}

/**
 * Add / Update by StringKey
 */

// KeyAdd
func (ht *Array) KeyAdd(key string, pData *Zval) *Zval {
	ht.assertRc1()
	if ht.KeyExists(key) {
		return nil
	}

	var p = ht.appendBucketStr(key, pData)
	return p.GetVal()
}

// KeyAddNew
func (ht *Array) KeyAddNew(key string, pData *Zval) *Zval {
	ht.assertRc1()

	var p = ht.appendBucketStr(key, pData)
	return p.GetVal()
}

func (ht *Array) KeyAddIndirect(strKey string, pData *Zval) *Zval {
	ht.assertRc1()

	var p = ht.KeyFindBucket(strKey)
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
		if ht.GetPDestructor() != nil {
			ht.GetPDestructor()(data)
		}
		ZVAL_COPY_VALUE(data, pData)
		return data
	}

	p = ht.appendBucketStr(strKey, pData)
	return p.GetVal()
}

// KeyUpdate
func (ht *Array) KeyUpdate(key string, pData *Zval) *Zval {
	ht.assertRc1()

	var p = ht.KeyFindBucket(key)
	if p != nil {
		var data *Zval
		b.Assert(p.GetVal() != pData)
		data = p.GetVal()
		if ht.GetPDestructor() != nil {
			ht.GetPDestructor()(data)
		}
		ZVAL_COPY_VALUE(data, pData)
		return data
	}

	p = ht.appendBucketStr(key, pData)
	return p.GetVal()
}

// KeyUpdateIndirect
func (ht *Array) KeyUpdateIndirect(key string, pData *Zval) *Zval {
	ht.assertRc1()

	var p = ht.KeyFindBucket(key)
	if p != nil {
		var data *Zval
		b.Assert(p.GetVal() != pData)
		data = p.GetVal()
		if data.IsType(IS_INDIRECT) {
			data = data.GetZv()
		}
		if ht.GetPDestructor() != nil {
			ht.GetPDestructor()(data)
		}
		ZVAL_COPY_VALUE(data, pData)
		return data
	}

	p = ht.appendBucketStr(key, pData)
	return p.GetVal()
}

/**
 * Add / Update by ArrayKey
 */
func (ht *Array) Add(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return ht.KeyAdd(key.KeyKey(), pData)
	} else {
		return ht.IndexAdd(key.IndexKey(), pData)
	}
}

func (ht *Array) AddIndirect(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return ht.KeyAddIndirect(key.KeyKey(), pData)
	} else {
		return ht.IndexAdd(key.IndexKey(), pData)
	}
}

func (ht *Array) Update(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return ht.KeyUpdate(key.KeyKey(), pData)
	} else {
		return ht.IndexUpdate(key.IndexKey(), pData)
	}
}

func (ht *Array) UpdateIndirect(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return ht.KeyUpdateIndirect(key.KeyKey(), pData)
	} else {
		return ht.IndexUpdate(key.IndexKey(), pData)
	}
}

/**
 * Delete
 */
func (ht *Array) KeyDelete(key string) bool {
	if idx, ok := ht.keyMap[key]; ok {
		ht.deleteBucket(idx)
		return true
	}
	return false
}

func (ht *Array) KeyDeleteIndirect(key string) bool {
	ht.assertRc1()
	if idx, ok := ht.keyMap[key]; ok {
		var p = &ht.data[idx]
		if p.GetVal().IsType(IS_INDIRECT) {
			var data *Zval = p.GetVal().GetZv()
			if data.IsType(IS_UNDEF) {
				return false
			} else {
				if ht.GetPDestructor() != nil {
					var tmp Zval
					ZVAL_COPY_VALUE(&tmp, data)
					data.SetUndef()
					ht.GetPDestructor()(&tmp)
				} else {
					data.SetUndef()
				}
				ht.SetIsHasEmptyInd()
			}
		} else {
			ht.deleteBucket(idx)
		}
		return true
	}
	return false
}

func (ht *Array) IndexDelete(index int) bool {
	if idx, ok := ht.indexMap[index]; ok {
		ht.deleteBucket(idx)
		return true
	}
	return false
}

func (ht *Array) deleteBucket(pos uint32) {
	ht.assertRc1()
	b.Assert(pos < ht.DataSize())

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
		ht.GetPDestructor()(&tmp)
	}

	// 设置数据不可用
	p.SetInvalid()

	// 若删除队尾元素，尝试清除 data 队尾无用数据
	if ht.DataSize()-1 == pos {
		ht.removeInvalidTail()
	}
}

/**
 * Clean && Destroy
 */
func (ht *Array) Clean() {
	ht.assertRc1()
	if ht.GetNNumUsed() != 0 {
		if ht.destructor != nil {
			ht.eachValidBucket(func(pos uint32, p *Bucket) {
				ht.destructor(p.GetVal())
			})
		}
	}
	ht.clearData()
}

func (ht *Array) Destroy() {
	if ht.DataSize() != 0 {
		if ht.destructor != nil {
			ht.eachValidBucket(func(pos uint32, p *Bucket) {
				ht.destructor(p.GetVal())
			})
		}
	}
	ZendHashIteratorsRemove(ht)
}

func (ht *Array) DestroyEx() {
	/* break possible cycles */
	//GC_REMOVE_FROM_BUFFER(ht)
	ht.SetGcTypeInfo(IS_NULL)

	ht.Destroy()
}

func (ht *Array) GracefulReverseDestroy() {
	ht.assertRc1()
	for idx := ht.DataSize(); idx > 0; idx-- {
		var p = &ht.data[idx-1]
		if p.IsValid() {
			ht.deleteBucket(idx)
		}
	}
}
