package types

import (
	b "sik/builtin"
	"sik/zend"
)

/**
 * Open methods
 */
func (ht *Array) RealInit()                               { ht.clearData() } // todo remove 无需init操作
func (ht *Array) GetArData() *Bucket                      { return ht.arData }
func (ht *Array) DataSize() uint32                        { return uint32(len(ht.data)) }
func (ht *Array) LastPos() uint32                         { return ht.DataSize() - 1 }
func (ht *Array) GetNNumUsed() uint32                     { return ht.DataSize() }
func (ht *Array) SetNNumUsed(value uint32)                {} // todo remove
func (ht *Array) SetNNumOfElements(value uint32)          { ht.elementsCount = value }
func (ht *Array) GetNInternalPointer() uint32             { return ht.internalPointer }
func (ht *Array) SetNInternalPointer(value uint32)        { ht.internalPointer = value }
func (ht *Array) GetNNextFreeElement() zend.ZendLong      { return ht.nextFreeElement }
func (ht *Array) SetNNextFreeElement(value zend.ZendLong) { ht.nextFreeElement = value }
func (ht *Array) GetPDestructor() DtorFuncT               { return ht.destructor }
func (ht *Array) SetPDestructor(value DtorFuncT)          { ht.destructor = value }
func (ht *Array) GetNTableMask() uint32                   { return 0 } // todo remove

/**
 * Clean && Destroy
 */
func (ht *Array) Clean() {
	ht.assertRc1()
	if ht.elementsCount != 0 && ht.destructor != nil {
		ht.eachValidBucket(func(pos uint32, p *Bucket) {
			ht.destructor(p.GetVal())
		})
	}
	ht.clearData()
}

func (ht *Array) Destroy() {
	if ht.elementsCount != 0 && ht.destructor != nil {
		ht.eachValidBucket(func(pos uint32, p *Bucket) {
			ht.destructor(p.GetVal())
		})
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
		pos := idx - 1
		p := &ht.data[pos]
		if p.IsValid() {
			ht.deleteBucket(idx)
		}
	}
}

/**
 * Methods use index key
 */
func (ht *Array) IndexFind(index int) *Zval {
	if pos, ok := ht.indexMap[index]; ok {
		return &ht.data[pos].val
	}
	return nil
}
func (ht *Array) IndexExists(index int) bool {
	_, ok := ht.indexMap[index]
	return ok
}
func (ht *Array) IndexAdd(index int, pData *Zval) *Zval {
	ht.assertRc1()
	if ht.IndexExists(index) {
		return nil
	}
	return ht.appendBucketIndex(index, pData).GetVal()
}
func (ht *Array) IndexAddNew(index int, pData *Zval) *Zval {
	ht.assertRc1()
	return ht.appendBucketIndex(index, pData).GetVal()
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
func (ht *Array) NextIndexInsert(pData *Zval) *Zval {
	ht.assertRc1()

	var index = ht.nextFreeElement

	if ht.IndexExists(index) {
		return nil
	}

	var p = ht.appendBucketIndex(index, pData)
	return p.GetVal()
}
func (ht *Array) NextIndexInsertNew(pData *Zval) *Zval {
	ht.assertRc1()

	var index = ht.nextFreeElement
	var p = ht.appendBucketIndex(index, pData)
	return p.GetVal()
}
func (ht *Array) IndexDelete(index int) bool {
	if idx, ok := ht.indexMap[index]; ok {
		ht.deleteBucket(idx)
		return true
	}
	return false
}

/**
 * Methods use string key
 */

func (ht *Array) KeyFind(key string) *Zval {
	if pos, ok := ht.keyMap[key]; ok {
		return &ht.data[pos].val
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
	_, ok := ht.keyMap[key]
	return ok
}
func (ht *Array) KeyExistsIndirect(key string) bool {
	var zv = ht.KeyFind(key)
	if zv == nil {
		return false
	}

	if zv.IsUndef() && zv.GetZv().IsUndef() {
		return false
	}

	return true
}
func (ht *Array) KeyAdd(key string, pData *Zval) *Zval {
	ht.assertRc1()
	if ht.KeyExists(key) {
		return nil
	}

	var p = ht.appendBucketStr(key, pData)
	return p.GetVal()
}
func (ht *Array) KeyAddNew(key string, pData *Zval) *Zval {
	ht.assertRc1()

	var p = ht.appendBucketStr(key, pData)
	return p.GetVal()
}
func (ht *Array) KeyAddIndirect(key string, pData *Zval) *Zval {
	ht.assertRc1()

	if data := ht.KeyFind(key); data != nil {
		b.Assert(data != pData)
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

	return ht.appendBucketStr(key, pData).GetVal()
}
func (ht *Array) KeyUpdate(key string, pData *Zval) *Zval {
	ht.assertRc1()

	if data := ht.KeyFind(key); data != nil {
		b.Assert(data != pData)
		if ht.GetPDestructor() != nil {
			ht.GetPDestructor()(data)
		}
		ZVAL_COPY_VALUE(data, pData)
		return data
	}

	return ht.appendBucketStr(key, pData).GetVal()
}
func (ht *Array) KeyUpdateIndirect(key string, pData *Zval) *Zval {
	ht.assertRc1()

	if data := ht.KeyFind(key); data != nil {
		b.Assert(data != pData)
		if data.IsType(IS_INDIRECT) {
			data = data.GetZv()
		}
		if ht.GetPDestructor() != nil {
			ht.GetPDestructor()(data)
		}
		ZVAL_COPY_VALUE(data, pData)
		return data
	}

	return ht.appendBucketStr(key, pData).GetVal()
}
func (ht *Array) KeyDelete(key string) bool {
	if pos, ok := ht.keyMap[key]; ok {
		ht.deleteBucket(pos)
		return true
	}
	return false
}

func (ht *Array) KeyDeleteIndirect(key string) bool {
	ht.assertRc1()
	if pos, ok := ht.keyMap[key]; ok {
		var p = &ht.data[pos]
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
				ht.MarkHasEmptyIndex()
			}
		} else {
			ht.deleteBucket(pos)
		}
		return true
	}
	return false
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
 * each
 */
func (ht *Array) Foreach(handler func(key ArrayKey, value *Zval)) {
	for i, _ := range ht.data {
		handler(ht.data[i].GetArrayKey(), &ht.data[i].val)
	}
}

func (ht *Array) eachBucket(handler func(pos uint32, p *Bucket)) {
	var size = uint32(len(ht.data))
	for i := uint32(0); i < size; i++ {
		var p = &ht.data[i]
		handler(i, p)
	}
}
func (ht *Array) eachValidBucket(handler func(pos uint32, p *Bucket)) {
	var size = uint32(len(ht.data))
	for i := uint32(0); i < size; i++ {
		var p = &ht.data[i]
		if p.IsValid() {
			continue
		}
		handler(i, p)
	}
}
func (ht *Array) eachValidBucketIndirect(handler func(pos uint32, p *Bucket, data *Zval)) {
	var size = uint32(len(ht.data))
	for i := uint32(0); i < size; i++ {
		var p = &ht.data[i]
		var data = p.GetVal()

		if data.IsIndirect() {
			data = data.GetZv()
		}
		if data.IsUndef() {
			return
		}

		handler(i, p, data)
	}
}

/**
 * Internal methods
 */
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
