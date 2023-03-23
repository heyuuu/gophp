package types

/**
 * Open methods
 */
func (ht *Array) RealInit() { ht.clearData() } // todo remove 无需init操作

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
