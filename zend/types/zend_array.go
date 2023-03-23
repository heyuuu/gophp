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
	// todo remove
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

func (ht *Array) GetArData() *Bucket { return ht.arData }

func (ht *Array) DataSize() uint32 { return uint32(len(ht.data)) }
func (ht *Array) LastPos() uint32  { return ht.DataSize() - 1 }

func (ht *Array) GetNNumUsed() uint32      { return ht.DataSize() }
func (ht *Array) SetNNumUsed(value uint32) {} // todo remove

func (ht *Array) SetNNumOfElements(value uint32)          { ht.elementsCount = value }
func (ht *Array) GetNInternalPointer() uint32             { return ht.internalPointer }
func (ht *Array) SetNInternalPointer(value uint32)        { ht.internalPointer = value }
func (ht *Array) GetNNextFreeElement() zend.ZendLong      { return ht.nextFreeElement }
func (ht *Array) SetNNextFreeElement(value zend.ZendLong) { ht.nextFreeElement = value }
func (ht *Array) GetPDestructor() DtorFuncT               { return ht.destructor }
func (ht *Array) SetPDestructor(value DtorFuncT)          { ht.destructor = value }

func (ht *Array) GetNTableMask() uint32 { return 0 } // todo remove

/**
 * Bucket 相关读接口
 */

func (ht *Array) keyFindBucket(key string) *Bucket {
	if pos, ok := ht.keyMap[key]; ok {
		return &ht.data[pos]
	}
	return nil
}

func (ht *Array) KeyFind(key string) *Zval {
	var p = ht.keyFindBucket(key)
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

	// 若找到则更新
	if zv := ht.IndexFind(index); zv != nil {
		if ht.destructor != nil {
			ht.destructor(zv)
		}
		ZVAL_COPY_VALUE(zv, pData)
		return zv
	}

	// 插入后返回
	return ht.appendBucketIndex(index, pData).GetVal()
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

	var p = ht.keyFindBucket(strKey)
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

	var p = ht.keyFindBucket(key)
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

	var p = ht.keyFindBucket(key)
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
