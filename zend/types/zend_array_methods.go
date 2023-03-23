package types

import (
	"math"
	b "sik/builtin"
	"sik/zend"
	"sik/zend/faults"
	"sort"
)

func (ht *Array) Sort(comparer func(a *Bucket, b *Bucket) bool, renumber bool) {
	ht.assertRc1()

	if ht.elementsCount == 0 || (ht.elementsCount == 1 && !renumber) {
		return
	}

	ht.removeHolesForce()
	ht.SetNInternalPointer(0)

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

func (ht *Array) SortCompatible(comparer CompareFuncT, renumber ZendBool) int {
	ht.Sort(func(a *Bucket, b *Bucket) bool {
		var compareResult = comparer(a, b)
		return compareResult > 0
	}, renumber != 0)
	return SUCCESS
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

// ---

func (ht *Array) posBucket(p *Bucket) (uint32, bool) {
	if p.IsStrKey() {
		if pos, ok := ht.keyMap[p.StrKey()]; ok {
			return pos, true
		}
		return 0, false
	} else {
		if pos, ok := ht.indexMap[p.IndexKey()]; ok {
			return pos, true
		}
		return 0, false
	}
}

func (ht *Array) SetBucketKey(b *Bucket, key string) *Zval {
	ht.assertRc1()

	// 若已存在此key，与设置值相同则返回 val；否则返回 nil (设置失败)
	var p = ht.KeyFindBucket(key)
	if p != nil {
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

func (ht *Array) addHash(key ArrayKey, pos uint32) {
	if key.IsStrKey() {
		ht.keyMap[key.KeyKey()] = pos
	} else {
		ht.indexMap[key.IndexKey()] = pos
	}
}

func (ht *Array) deleteHash(key ArrayKey) {
	if key.IsStrKey() {
		delete(ht.keyMap, key.KeyKey())
	} else {
		delete(ht.indexMap, key.IndexKey())
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

func (ht *Array) applyValidBucket(apply_func func(p *Bucket) int) {
	var size = ht.DataSize()
	for idx := uint32(0); idx < size; idx++ {
		var p = &ht.data[idx]
		if !p.IsValid() {
			continue
		}
		var result = apply_func(p)
		if b.FlagMatch(result, ZEND_HASH_APPLY_REMOVE) {
			ht.deleteBucket(idx)
		}
		if b.FlagMatch(result, ZEND_HASH_APPLY_STOP) {
			break
		}
	}
}
func (ht *Array) applyValidBucketReserve(apply_func func(p *Bucket) int) {
	for idx := ht.DataSize(); idx > 0; idx-- {
		var p = &ht.data[idx-1]
		if !p.IsValid() {
			continue
		}
		var result = apply_func(p)
		if b.FlagMatch(result, ZEND_HASH_APPLY_REMOVE) {
			ht.deleteBucket(idx - 1)
		}
		if b.FlagMatch(result, ZEND_HASH_APPLY_STOP) {
			break
		}
	}
}

func (ht *Array) foreachData() []*Bucket {
	// todo 逐渐替换为 eachBucket 或其他更高效代码
	var data = make([]*Bucket, 0)
	ht.eachValidBucket(func(_ uint32, p *Bucket) {
		data = append(data, p)
	})
	return data
}

func (ht *Array) foreachDataReserve() []*Bucket {
	// todo 逐渐替换为 eachBucket 或其他更高效代码
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

// 移动 bucket 到新位置
func (ht *Array) _moveBucket(pos uint32, newPos uint32) {
	b.Assert(newPos <= pos)
	if newPos == pos {
		return
	}
	(&ht.data[newPos]).CopyFrom(&ht.data[pos])
	if ht.internalPointer == pos {
		ht.internalPointer = newPos
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
			ht._moveBucket(pos, newPos)
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
			ht._moveBucket(pos, newPos)
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
			//(&ht.data[newPos]).CopyFrom(&ht.data[pos])
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

func (ht *Array) Rehash() {
	// 空数组快速清空
	if ht.elementsCount == 0 {
		ht.resetHash()
		ht.data = nil
		return
	}

	// 移除 data 中的空位
	var oldNumUsed = ht.GetNNumUsed()
	ht.removeHoles()

	// 重建 hash
	ht.resetHash()
	ht.eachBucket(func(pos uint32, p *Bucket) {
		ht.addHash(p.key, pos)
	})

	/* Migrate pointer to one past the end of the array to the new one past the end, so that
	 * newly inserted elements are picked up correctly. */
	if ht.HasIterators() {
		_zendHashIteratorsUpdate(ht, oldNumUsed, ht.GetNNumUsed())
	}
}

func (ht *Array) Extend(size uint32) {
	ht.assertRc1()
	if size > uint32(len(ht.data)) {
		// 扩展数组 cap
		newData := make([]Bucket, 0, size)
		if len(ht.data) > 0 {
			copy(newData, ht.data)
		}
		ht.data = newData
	}
}

func (ht *Array) Discard(nNumUsed uint32) {
	if nNumUsed < ht.DataSize() {
		// 裁剪数据，重新映射
		ht.data = ht.data[:nNumUsed]
		ht.Rehash()
		ht.elementsCount = ht.DataSize()
	}
}

// 重新计算有效元素个数(与 nnNumOfElements 不同，它考虑 IS_INDIRECT 元素为 IS_UNDEF 的情况)
func (ht *Array) RecalcElements() uint32 {
	var num uint32 = 0
	for _, bucket := range ht.data {
		var val = bucket.GetVal()
		if val.IsType(IS_UNDEF) {
			continue
		}
		if val.IsType(IS_INDIRECT) && val.GetZv().IsType(IS_UNDEF) {
			continue
		}
		num++
	}
	return num
}

func (ht *Array) Count() uint32 {
	var num uint32
	if ht.IsHasEmptyInd() {
		num = ht.RecalcElements()
		if ht.elementsCount == num {
			ht.UnsetIsHasEmptyInd()
		}
	} else if ht == zend.EG__().GetSymbolTable() {
		num = ht.RecalcElements()
	} else {
		num = ht.CountElements()
	}
	return num
}

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
	var dataSize = ht.DataSize()
	for ; pos < dataSize; pos++ {
		if ht.IsValidPos(pos) {
			return pos, true
		}
	}
	// 没有有效pos，此时 pos == ht.DataSize()
	return pos, false
}

func (ht *Array) IsValidPos(pos uint32) bool {
	b.Assert(pos < ht.DataSize())
	return !ht.data[pos].GetVal().IsType(IS_UNDEF)
}

// ----

func (ht *Array) IsWithoutHoles() bool { return ht.GetNNumUsed() == ht.elementsCount }

func (ht *Array) resizeIfFull() {
	dataSize := len(ht.data)
	if dataSize == cap(ht.data) {
		// 若空隙率过高，重新压缩；否则，跳过扩容 (后面会由 append(ht.data) 触发自动扩容)
		if dataSize > int(ht.elementsCount+(ht.elementsCount>>5)) {
			ht.Rehash()
		} else if dataSize >= math.MaxInt32 {
			faults.ErrorNoreturn(faults.E_ERROR, "Possible integer overflow in memory allocation (%d)", dataSize*2)
		}
	}
}

func (ht *Array) appendBucket(bucket *Bucket) *Bucket {
	// 尝试 resize
	ht.resizeIfFull()

	// 添加到 data
	var idx = uint32(len(ht.data))
	ht.elementsCount++
	ht.data = append(ht.data, *bucket)

	// 更新 map
	ht.addHash(bucket.key, idx)

	if !bucket.IsStrKey() {
		var indexKey = bucket.IndexKey()
		// 更新 nextFreeElement
		if indexKey > ht.nextFreeElement {
			if indexKey < zend.ZEND_LONG_MAX {
				ht.nextFreeElement = indexKey + 1
			} else {
				ht.nextFreeElement = zend.ZEND_LONG_MAX
			}
		}
	}

	return &ht.data[idx]
}

func (ht *Array) appendBucketStr(strKey string, zv *Zval) *Bucket {
	var bucket = NewStrKeyBucket(strKey, zv)
	return ht.appendBucket(bucket)
}

func (ht *Array) appendBucketIndex(indexKey int, zv *Zval) *Bucket {
	var bucket = NewIndexBucket(indexKey, zv)
	return ht.appendBucket(bucket)
}
