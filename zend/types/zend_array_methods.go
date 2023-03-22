package types

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/faults"
	"sort"
)

func (this *Array) Sort(comparer func(a *Bucket, b *Bucket) bool, renumber bool) {
	this.assertRc1()

	if this.nNumOfElements == 0 || (this.nNumOfElements == 1 && !renumber) {
		return
	}

	this.removeHolesForce()
	this.SetNInternalPointer(0)

	sort.SliceStable(this.data, func(i, j int) bool {
		return comparer(&this.data[i], &this.data[i])
	})

	if renumber {
		this.eachBucket(func(pos uint32, p *Bucket) {
			p.SetIndexKey(int(pos))
		})
		this.nNextFreeElement = int(this.DataSize())
	}

	this.Rehash()
}

func (this *Array) SortCompatible(comparer CompareFuncT, renumber ZendBool) int {
	this.Sort(func(a *Bucket, b *Bucket) bool {
		var compareResult = comparer(a, b)
		return compareResult > 0
	}, renumber != 0)
	return SUCCESS
}

func (this *Array) SortCompatibleEx(sort_ SortFuncT) int {
	// todo sort 转 sortFunc 需要订制处理
	var sortFunc = *b.Cast[func([]Bucket)](&sort_)

	// 正常 sort 逻辑，除 sortFunc 部分外和 this.Sort() 逻辑一致
	this.assertRc1()

	if this.nNumOfElements <= 1 {
		return SUCCESS
	}

	this.removeHolesForce()
	this.SetNInternalPointer(0)

	sortFunc(this.data)

	this.Rehash()

	return SUCCESS
}

// ---

func (this *Array) posBucket(p *Bucket) (uint32, bool) {
	if p.IsStrKey() {
		if pos, ok := this.keyMap[p.StrKey()]; ok {
			return pos, true
		}
		return 0, false
	} else {
		if pos, ok := this.indexMap[p.IndexKey()]; ok {
			return pos, true
		}
		return 0, false
	}
}

func (this *Array) SetBucketKey(b *Bucket, key string) *Zval {
	this.assertRc1()

	// 若已存在此key，与设置值相同则返回 val；否则返回 nil (设置失败)
	var p = this.KeyFindBucket(key)
	if p != nil {
		if p == b {
			return p.GetVal()
		} else {
			return nil
		}
	}

	// 定义 bucket 位置；若 bucket 不在数据内，返回 nil
	var pos, ok = this.posBucket(b)
	if !ok {
		return nil
	}

	/* del from hash */
	this.deleteHash(b.key)

	/* add to hash */
	b.SetStrKey(key)
	this.addHash(b.key, pos)

	return b.GetVal()
}

func (this *Array) addHash(key ArrayKey, pos uint32) {
	if key.IsStrKey() {
		this.keyMap[key.KeyKey()] = pos
	} else {
		this.indexMap[key.IndexKey()] = pos
	}
}

func (this *Array) deleteHash(key ArrayKey) {
	if key.IsStrKey() {
		delete(this.keyMap, key.KeyKey())
	} else {
		delete(this.indexMap, key.IndexKey())
	}
}

func (this *Array) eachBucket(handler func(pos uint32, p *Bucket)) {
	var size = uint32(len(this.data))
	for i := uint32(0); i < size; i++ {
		var p = &this.data[i]
		handler(i, p)
	}
}

func (this *Array) eachValidBucket(handler func(pos uint32, p *Bucket)) {
	var size = uint32(len(this.data))
	for i := uint32(0); i < size; i++ {
		var p = &this.data[i]
		if p.IsValid() {
			continue
		}
		handler(i, p)
	}
}

func (this *Array) eachValidBucketIndirect(handler func(pos uint32, p *Bucket, data *Zval)) {
	var size = uint32(len(this.data))
	for i := uint32(0); i < size; i++ {
		var p = &this.data[i]
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

func (this *Array) applyValidBucket(apply_func func(p *Bucket) int) {
	var size = this.DataSize()
	for idx := uint32(0); idx < size; idx++ {
		var p = &this.data[idx]
		if !p.IsValid() {
			continue
		}
		var result = apply_func(p)
		if b.FlagMatch(result, ZEND_HASH_APPLY_REMOVE) {
			this.deleteBucket(idx)
		}
		if b.FlagMatch(result, ZEND_HASH_APPLY_STOP) {
			break
		}
	}
}
func (this *Array) applyValidBucketReserve(apply_func func(p *Bucket) int) {
	for idx := this.DataSize(); idx > 0; idx-- {
		var p = &this.data[idx-1]
		if !p.IsValid() {
			continue
		}
		var result = apply_func(p)
		if b.FlagMatch(result, ZEND_HASH_APPLY_REMOVE) {
			this.deleteBucket(idx - 1)
		}
		if b.FlagMatch(result, ZEND_HASH_APPLY_STOP) {
			break
		}
	}
}

func (this *Array) foreachData() []*Bucket {
	// todo 逐渐替换为 eachBucket 或其他更高效代码
	var data = make([]*Bucket, 0)
	this.eachValidBucket(func(_ uint32, p *Bucket) {
		data = append(data, p)
	})
	return data
}

func (this *Array) foreachDataReserve() []*Bucket {
	// todo 逐渐替换为 eachBucket 或其他更高效代码
	var data = make([]*Bucket, 0)

	for i := len(this.data) - 1; i >= 0; i-- {
		var p = &this.data[i]
		if p.IsValid() {
			continue
		}
		data = append(data, p)
	}

	return data
}

// 移动 bucket 到新位置
func (this *Array) _moveBucket(pos uint32, newPos uint32) {
	b.Assert(newPos <= pos)
	if newPos == pos {
		return
	}
	(&this.data[newPos]).CopyFrom(&this.data[pos])
	if this.nInternalPointer == pos {
		this.nInternalPointer = newPos
	}
}

// 移除 this.data 数据中的 holes, 返回是否移动 bucket
func (this *Array) removeHoles() bool {
	var newPos uint32 = 0

	if this.IsWithoutHoles() {
		return false
	}

	if this.HasIterators() {
		var iterPos = ZendHashIteratorsLowerPos(this, 0)

		this.eachValidBucket(func(pos uint32, p *Bucket) {
			// 移动 bucket 到新位置
			this._moveBucket(pos, newPos)
			if pos != newPos {
				if pos >= iterPos {
					for {
						ZendHashIteratorsUpdate(this, iterPos, newPos)
						iterPos = ZendHashIteratorsLowerPos(this, iterPos+1)
						if iterPos >= pos {
							break
						}
					}
				}
			}
			newPos++
		})
	} else {
		this.eachValidBucket(func(pos uint32, p *Bucket) {
			this._moveBucket(pos, newPos)
			newPos++
		})
	}

	// 截取数据，记录有效元素数
	this.data = this.data[:newPos]
	this.nNumOfElements = newPos

	b.Assert(this.IsWithoutHoles())

	return true
}

// 移除 data 的 holes, 不考虑 nInternalPointer 和 Iterators 内的 pos 指针
func (this *Array) removeHolesForce() bool {
	var newPos uint32 = 0

	if this.IsWithoutHoles() {
		return false
	}

	this.eachValidBucket(func(pos uint32, p *Bucket) {
		if newPos != pos {
			// todo 考虑下实现细节的区别
			//(&this.data[newPos]).CopyFrom(&this.data[pos])
			this.data[newPos] = this.data[pos]
		}
		newPos++
	})

	// 截取数据，记录有效元素数
	this.data = this.data[:newPos]
	this.nNumOfElements = newPos

	b.Assert(this.IsWithoutHoles())

	return true
}

// 清除 data 队尾无用数据
func (this *Array) removeInvalidTail() {
	var dataSize = this.DataSize()

	// 从队尾依次判断是否为无效数据，若是则缩短
	var newDataSize = dataSize
	for newDataSize > 0 && !this.data[newDataSize-1].IsValid() {
		newDataSize--
	}

	// 若长度改变，调整 data
	if newDataSize < dataSize {
		this.data = this.data[:newDataSize]
		this.nInternalPointer = b.Min(this.nInternalPointer, newDataSize)
	}
}

func (this *Array) Rehash() {
	// 空数组快速清空
	if this.nNumOfElements == 0 {
		this.resetHash()
		this.data = nil
		return
	}

	// 移除 data 中的空位
	var oldNumUsed = this.GetNNumUsed()
	this.removeHoles()

	// 重建 hash
	this.resetHash()
	this.eachBucket(func(pos uint32, p *Bucket) {
		this.addHash(p.key, pos)
	})

	/* Migrate pointer to one past the end of the array to the new one past the end, so that
	 * newly inserted elements are picked up correctly. */
	if this.HasIterators() {
		zend._zendHashIteratorsUpdate(this, oldNumUsed, this.GetNNumUsed())
	}
}

func (this *Array) ifFullDoResize() {
	if this.DataSize() >= this.nTableSize {
		this.doResize()
	}
}

func (this *Array) doResize() {
	this.assertRc1()

	if this.DataSize() > this.nNumOfElements+(this.nNumOfElements>>5) {
		this.Rehash()
	} else if this.nTableSize < HT_MAX_SIZE {
		// 无内存复制，仅扩充尺寸标识
		this.nTableSize *= 2
	} else {
		faults.ErrorNoreturn(faults.E_ERROR, "Possible integer overflow in memory allocation (%d)", this.nTableSize*2)
	}
}

func (this *Array) Extend(nSize uint32) {
	// todo remove 无需手动扩展
	this.assertRc1()
	if nSize > this.nTableSize {
		// 无内存复制，仅扩充尺寸标识
		this.nTableSize = ZendHashCheckSize(nSize)
	}
}

func (this *Array) Discard(nNumUsed uint32) {
	if nNumUsed < this.DataSize() {
		// 裁剪数据，重新映射
		this.data = this.data[:nNumUsed]
		this.Rehash()
		this.nNumOfElements = this.DataSize()
	}
}

// 重新计算有效元素个数(与 nnNumOfElements 不同，它考虑 IS_INDIRECT 元素为 IS_UNDEF 的情况)
func (this *Array) RecalcElements() uint32 {
	var num uint32 = 0
	for _, bucket := range this.data {
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

func (this *Array) Count() uint32 {
	var num uint32
	if this.HasUFlags(HASH_FLAG_HAS_EMPTY_IND) {
		num = this.RecalcElements()
		if this.nNumOfElements == num {
			this.SubUFlags(HASH_FLAG_HAS_EMPTY_IND)
		}
	} else if this == zend.EG__().GetSymbolTable() {
		num = this.RecalcElements()
	} else {
		num = this.GetNNumOfElements()
	}
	return num
}

func (this *Array) currentPos() (uint32, bool) {
	return this.validPos(this.nInternalPointer)
}

func (this *Array) currentPosVal() uint32 {
	var pos, _ = this.currentPos()
	return pos
}

func (this *Array) validPosVal(pos uint32) uint32 {
	pos, _ = this.validPos(pos)
	return pos
}

func (this *Array) validPos(pos uint32) (uint32, bool) {
	var dataSize = this.DataSize()
	for ; pos < dataSize; pos++ {
		if this.IsValidPos(pos) {
			return pos, true
		}
	}
	// 没有有效pos，此时 pos == this.DataSize()
	return pos, false
}

func (this *Array) IsValidPos(pos uint32) bool {
	b.Assert(pos < this.DataSize())
	return !this.data[pos].GetVal().IsType(IS_UNDEF)
}

// ----

func (this *Array) IsWithoutHoles() bool { return this.GetNNumUsed() == this.nNumOfElements }

func (this *Array) appendBucket(bucket *Bucket) *Bucket {
	// 尝试 resize
	this.ifFullDoResize()

	// 添加到 data
	var idx = uint32(len(this.data))
	this.nNumOfElements++
	this.data = append(this.data, *bucket)

	// 更新 map
	this.addHash(bucket.key, idx)

	if !bucket.IsStrKey() {
		var indexKey = bucket.IndexKey()
		// 更新 nNextFreeElement
		if indexKey > this.nNextFreeElement {
			if indexKey < zend.ZEND_LONG_MAX {
				this.nNextFreeElement = indexKey + 1
			} else {
				this.nNextFreeElement = zend.ZEND_LONG_MAX
			}
		}
	}

	return &this.data[idx]
}

func (this *Array) appendBucketStr(strKey string, zv *Zval) *Bucket {
	var bucket = NewStrKeyBucket(strKey, zv)
	return this.appendBucket(bucket)
}

func (this *Array) appendBucketIndex(indexKey int, zv *Zval) *Bucket {
	var bucket = NewIndexBucket(indexKey, zv)
	return this.appendBucket(bucket)
}
