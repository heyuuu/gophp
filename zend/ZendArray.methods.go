package zend

import (
	b "sik/builtin"
)

func (this *HashTable) posBucket(p *Bucket) (uint32, bool) {
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

func (this *HashTable) SetBucketKey(b *Bucket, key string) *Zval {
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

	this.SubUFlags(HASH_FLAG_STATIC_KEYS)

	/* del from hash */
	this.deleteHash(b.key)

	/* add to hash */
	b.SetStrKey(key)
	this.addHash(b.key, pos)

	return b.GetVal()
}

func (this *ZendArray) addHash(key ZendArrayKey, pos uint32) {
	if key.IsStrKey() {
		this.keyMap[key.GetKey()] = pos
	} else {
		this.indexMap[key.GetIndex()] = pos
	}
}

func (this *ZendArray) deleteHash(key ZendArrayKey) {
	if key.IsStrKey() {
		delete(this.keyMap, key.GetKey())
	} else {
		delete(this.indexMap, key.GetIndex())
	}
}

func (this *HashTable) eachBucket(handler func(pos uint32, p *Bucket)) {
	var size = uint32(len(this.data))
	for i := uint32(0); i < size; i++ {
		var p = &this.data[i]
		handler(i, p)
	}
}

func (this *HashTable) eachValidBucket(handler func(pos uint32, p *Bucket)) {
	var size = uint32(len(this.data))
	for i := uint32(0); i < size; i++ {
		var p = &this.data[i]
		if p.IsValid() {
			continue
		}
		handler(i, p)
	}
}

func (this *HashTable) foreachData() []*Bucket {
	// todo 逐渐替换为 eachBucket 或其他更高效代码
	var data = make([]*Bucket, 0)
	this.eachValidBucket(func(_ uint32, p *Bucket) {
		data = append(data, p)
	})
	return data
}

func (this *HashTable) foreachDataReserve() []*Bucket {
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
func (this *HashTable) _moveBucket(pos uint32, newPos uint32) {
	ZEND_ASSERT(newPos <= pos)
	if newPos == pos {
		return
	}
	(&this.data[newPos]).CopyFrom(&this.data[pos])
	if this.nInternalPointer == pos {
		this.nInternalPointer = newPos
	}
}

// 移除 this.data 数据中的 holes, 返回是否移动 bucket
func (this *HashTable) removeHoles() bool {
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

	ZEND_ASSERT(this.IsWithoutHoles())

	return true
}

// 清除 data 队尾无用数据
func (this *HashTable) removeInvalidTail() {
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

func (this *HashTable) Rehash() {
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
		_zendHashIteratorsUpdate(this, oldNumUsed, this.GetNNumUsed())
	}
}

func (this *HashTable) ifFullDoResize() {
	if this.DataSize() >= this.nTableSize {
		this.doResize()
	}
}

func (this *HashTable) doResize() {
	this.assertRc1()

	if this.DataSize() > this.nNumOfElements+(this.nNumOfElements>>5) {
		this.Rehash()
	} else if this.nTableSize < HT_MAX_SIZE {
		// 无内存复制，仅扩充尺寸标识
		this.nTableSize *= 2
	} else {
		ZendErrorNoreturn(E_ERROR, "Possible integer overflow in memory allocation (%d)", this.nTableSize*2)
	}
}

func (this *HashTable) Extend(nSize uint32) {
	// todo remove 无需手动扩展
	this.assertRc1()
	if nSize > this.nTableSize {
		// 无内存复制，仅扩充尺寸标识
		this.nTableSize = ZendHashCheckSize(nSize)
	}
}

func (this *HashTable) Discard(nNumUsed uint32) {
	if nNumUsed < this.DataSize() {
		// 裁剪数据，重新映射
		this.data = this.data[:nNumUsed]
		this.Rehash()
		this.nNumOfElements = this.DataSize()
	}
}

// 重新计算有效元素个数(与 nnNumOfElements 不同，它考虑 IS_INDIRECT 元素为 IS_UNDEF 的情况)
func (this *HashTable) RecalcElements() uint32 {
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

func (this *HashTable) Count() uint32 {
	var num uint32
	if this.HasUFlags(HASH_FLAG_HAS_EMPTY_IND) {
		num = this.RecalcElements()
		if this.nNumOfElements == num {
			this.SubUFlags(HASH_FLAG_HAS_EMPTY_IND)
		}
	} else if this == __EG().GetSymbolTable() {
		num = this.RecalcElements()
	} else {
		num = this.GetNNumOfElements()
	}
	return num
}

func (this *HashTable) currentPos() (uint32, bool) {
	return this.validPos(this.nInternalPointer)
}

func (this *HashTable) currentPosVal() uint32 {
	var pos, _ = this.currentPos()
	return pos
}

func (this *HashTable) validPosVal(pos uint32) uint32 {
	pos, _ = this.validPos(pos)
	return pos
}

func (this *HashTable) validPos(pos uint32) (uint32, bool) {
	var dataSize = this.DataSize()
	for ; pos < dataSize; pos++ {
		if this.IsValidPos(pos) {
			return pos, true
		}
	}
	// 没有有效pos，此时 pos == this.DataSize()
	return pos, false
}

func (this *HashTable) IsValidPos(pos uint32) bool {
	ZEND_ASSERT(pos < this.DataSize())
	return !this.data[pos].GetVal().IsType(IS_UNDEF)
}

// ----

func (this *ZendArray) IsWithoutHoles() bool { return this.GetNNumUsed() == this.nNumOfElements }

func (this *ZendArray) appendBucket(bucket *Bucket) *Bucket {
	// 尝试 resize
	this.ifFullDoResize()

	// 添加到 data
	var idx = uint32(len(this.data))
	this.nNumOfElements++
	this.data = append(this.data, *bucket)

	// 更新 map
	this.addHash(bucket.key, idx)

	if bucket.IsIndexKey() {
		var indexKey = bucket.IndexKey()
		// 更新 nNextFreeElement
		if indexKey > this.nNextFreeElement {
			if indexKey < ZEND_LONG_MAX {
				this.nNextFreeElement = indexKey + 1
			} else {
				this.nNextFreeElement = ZEND_LONG_MAX
			}
		}
	}

	return &this.data[idx]
}

func (this *ZendArray) appendBucketStr(strKey string, zv *Zval) *Bucket {
	var bucket = NewBucketStr(strKey, zv)
	return this.appendBucket(bucket)
}

func (this *ZendArray) appendBucketIndex(indexKey int, zv *Zval) *Bucket {
	var bucket = NewBucketIndex(indexKey, zv)
	return this.appendBucket(bucket)
}
