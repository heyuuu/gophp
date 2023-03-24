package types

import (
	b "sik/builtin"
	"sik/zend"
	"sort"
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

func (ht *Array) IsWithoutHoles() bool { return ht.GetNNumUsed() == ht.elementsCount }

func (ht *Array) Count() uint32 {
	var num uint32
	if ht.HasEmptyIndex() {
		num = ht.recalcElements()
		if ht.elementsCount == num {
			ht.UnmarkHasEmptyIndex()
		}
	} else if ht == zend.EG__().GetSymbolTable() { // todo
		num = ht.recalcElements()
	} else {
		num = ht.Len()
	}
	return num
}

// 重新计算有效元素个数(与 elementsCount 不同，它考虑 IS_INDIRECT 元素为 IS_UNDEF 的情况)
func (ht *Array) recalcElements() uint32 {
	var num uint32 = 0
	ht.eachValidBucketIndirect(func(pos uint32, p *Bucket, data *Zval) {
		num++
	})
	return num
}

/**
 * Sort
 */
type ArraySortFunc func(a *Bucket, b *Bucket) bool
type ArraySortExFunc func(a *Bucket, b *Bucket) int

func (ht *Array) Sort(comparer ArraySortFunc, renumber bool) {
	ht.assertRc1()

	if ht.elementsCount == 0 || (ht.elementsCount == 1 && !renumber) {
		return
	}

	ht.removeHolesForce()
	ht.internalPointer = 0

	sort.SliceStable(ht.data, func(i, j int) bool {
		return comparer(&ht.data[i], &ht.data[i])
	})

	if renumber {
		ht.eachBucket(func(pos uint32, p *Bucket) {
			p.SetIndexKey(int(pos))
		})
		ht.nextFreeElement = int(ht.DataSize())
	}

	ht.Rehash()
}

func (ht *Array) SortCompatible(comparer CompareFuncT, renumber ZendBool) bool {
	ht.Sort(func(a *Bucket, b *Bucket) bool {
		var compareResult = comparer(a, b)
		return compareResult > 0
	}, renumber != 0)
	return true
}

func (ht *Array) SortCompatibleEx(sort_ SortFuncT) int {
	// todo sort 转 sortFunc 需要订制处理
	var sortFunc = *b.Cast[func([]Bucket)](&sort_)

	// 正常 sort 逻辑，除 sortFunc 部分外和 ht.Sort() 逻辑一致
	ht.assertRc1()

	if ht.elementsCount <= 1 {
		return SUCCESS
	}

	ht.removeHolesForce()
	ht.SetNInternalPointer(0)

	sortFunc(ht.data)

	ht.Rehash()

	return SUCCESS
}

func (ht *Array) SetBucketKey(b *Bucket, key string) *Zval {
	ht.assertRc1()

	// 若已存在此key，与设置值相同则返回 val；否则返回 nil (设置失败)
	if pos, ok := ht.keyMap[key]; ok {
		p := &ht.data[pos]
		if p == b {
			return p.GetVal()
		} else {
			return nil
		}
	}

	// 定义 bucket 位置；若 bucket 不在数据内，返回 nil
	var pos, ok = ht.posBucket(b)
	if !ok {
		return nil
	}

	/* del from hash */
	ht.deleteHash(b.key)

	/* add to hash */
	b.SetStrKey(key)
	ht.addHash(b.key, pos)

	return b.GetVal()
}

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
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		handler(p.GetArrayKey(), p.GetVal())
	}
}
func (ht *Array) ForeachReserve(handler func(key ArrayKey, value *Zval)) {
	for i := len(ht.data) - 1; i >= 0; i-- {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		handler(p.GetArrayKey(), p.GetVal())
	}
}

// todo 逐渐替换为 Foreach 或其他更高效代码
func (ht *Array) ForeachData() []*Bucket {
	var data = make([]*Bucket, 0)
	ht.eachValidBucket(func(_ uint32, p *Bucket) {
		data = append(data, p)
	})
	return data
}

// todo 逐渐替换为 ForeachReserve 或其他更高效代码
func (ht *Array) ForeachDataReserve() []*Bucket {
	var data = make([]*Bucket, 0)

	for i := len(ht.data) - 1; i >= 0; i-- {
		var p = &ht.data[i]
		if p.IsValid() {
			continue
		}
		data = append(data, p)
	}

	return data
}

func (ht *Array) eachBucket(handler func(pos uint32, p *Bucket)) {
	for i, _ := range ht.data {
		p := &ht.data[i]
		handler(uint32(i), p)
	}
}
func (ht *Array) eachValidBucket(handler func(pos uint32, p *Bucket)) {
	for i, _ := range ht.data {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		handler(uint32(i), p)
	}
}
func (ht *Array) eachValidBucketIndirect(handler func(pos uint32, p *Bucket, data *Zval)) {
	for i, _ := range ht.data {
		p := &ht.data[i]
		data := p.GetVal()
		if data.IsIndirect() {
			data = data.GetZv()
		}
		if data.IsUndef() {
			return
		}
		handler(uint32(i), p, data)
	}
}

func (ht *Array) applyValidBucket(apply_func func(p *Bucket) int) {
	for i, _ := range ht.data {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		result := apply_func(p)
		if b.FlagMatch(result, ArrayApplyRemove) {
			ht.deleteBucket(uint32(i))
		}
		if b.FlagMatch(result, ArrayApplyStop) {
			break
		}
	}
}
func (ht *Array) applyValidBucketReserve(apply_func func(p *Bucket) int) {
	for i := len(ht.data) - 1; i >= 0; i-- {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		result := apply_func(p)
		if b.FlagMatch(result, ArrayApplyRemove) {
			ht.deleteBucket(uint32(i))
		}
		if b.FlagMatch(result, ArrayApplyStop) {
			break
		}
	}
}

/**
 * Iterator & Pos
 */
func (ht *Array) currentPos() (uint32, bool) {
	return ht.validPos(ht.internalPointer)
}

func (ht *Array) currentPosVal() uint32 {
	var pos, _ = ht.currentPos()
	return pos
}
func (ht *Array) validPosVal(pos uint32) uint32 {
	pos, _ = ht.validPos(pos)
	return pos
}

func (ht *Array) validPos(pos uint32) (uint32, bool) {
	dataSize := uint32(len(ht.data))
	for i := pos; i < dataSize; i++ {
		if ht.data[i].IsValid() {
			return i, true
		}
	}
	// 没有有效pos，此时 pos == ht.DataSize()
	return pos, false
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

// 移除 this.data 数据中的 holes, 返回是否移动 bucket
func (ht *Array) removeHoles() bool {
	var newPos uint32 = 0

	if ht.IsWithoutHoles() {
		return false
	}

	if ht.HasIterators() {
		var iterPos = ZendHashIteratorsLowerPos(ht, 0)

		ht.eachValidBucket(func(pos uint32, p *Bucket) {
			// 移动 bucket 到新位置
			ht.moveBucket(pos, newPos)
			if pos != newPos {
				if pos >= iterPos {
					for {
						ZendHashIteratorsUpdate(ht, iterPos, newPos)
						iterPos = ZendHashIteratorsLowerPos(ht, iterPos+1)
						if iterPos >= pos {
							break
						}
					}
				}
			}
			newPos++
		})
	} else {
		ht.eachValidBucket(func(pos uint32, p *Bucket) {
			ht.moveBucket(pos, newPos)
			newPos++
		})
	}

	// 截取数据，记录有效元素数
	ht.data = ht.data[:newPos]
	ht.elementsCount = newPos

	b.Assert(ht.IsWithoutHoles())

	return true
}

// 移除 data 的 holes, 不考虑 internalPointer 和 Iterators 内的 pos 指针
func (ht *Array) removeHolesForce() bool {
	var newPos uint32 = 0

	if ht.IsWithoutHoles() {
		return false
	}

	ht.eachValidBucket(func(pos uint32, p *Bucket) {
		if newPos != pos {
			// todo 考虑下实现细节的区别
			//(&ht.data[newPos]).SetBy(&ht.data[pos])
			ht.data[newPos] = ht.data[pos]
		}
		newPos++
	})

	// 截取数据，记录有效元素数
	ht.data = ht.data[:newPos]
	ht.elementsCount = newPos

	b.Assert(ht.IsWithoutHoles())

	return true
}

// 清除 data 队尾无用数据
func (ht *Array) removeInvalidTail() {
	var dataSize = ht.DataSize()

	// 从队尾依次判断是否为无效数据，若是则缩短
	var newDataSize = dataSize
	for newDataSize > 0 && !ht.data[newDataSize-1].IsValid() {
		newDataSize--
	}

	// 若长度改变，调整 data
	if newDataSize < dataSize {
		ht.data = ht.data[:newDataSize]
		ht.internalPointer = b.Min(ht.internalPointer, newDataSize)
	}
}
