package zend

import (
	b "sik/builtin"
)

func (this *ZendArray) addHash(key ZendArrayKey, pos uint32) {
	if key.IsStrKey() {
		this.keyMap[key.GetKey()] = pos
	} else {
		this.indexMap[key.GetIndex()] = pos
	}
}

func (this *HashTable) eachBucket(handler func(uint32, *Bucket)) {
	var size = uint32(len(this.data))
	for i := uint32(0); i < size; i++ {
		var p = &this.data[i]
		handler(i, p)
	}
}

func (this *HashTable) eachUsedBucket(handler func(uint32, *Bucket)) {
	var size = uint32(len(this.data))
	for i := uint32(0); i < size; i++ {
		var p = &this.data[i]
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		handler(i, p)
	}
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

		this.eachUsedBucket(func(pos uint32, p *Bucket) {
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
		this.eachUsedBucket(func(pos uint32, p *Bucket) {
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

func (this *HashTable) validPos(pos uint32) (uint32, bool) {
	var dataSize = this.DataSize()
	for ; pos < dataSize; pos++ {
		if !this.data[pos].GetVal().IsType(IS_UNDEF) {
			return pos, true
		}
	}
	// 没有有效pos，此时 pos == this.DataSize()
	return pos, false
}

// ----

func (this *ZendArray) IsWithoutHoles() bool { return this.GetNNumUsed() == this.nNumOfElements }

func (this *ZendArray) findPos(key ZendArrayKey) (uint32, bool) {
	if key.IsStrKey() {
		if pos, ok := this.keyMap[key.GetKey()]; ok {
			return pos, true
		}
	} else {
		if pos, ok := this.indexMap[key.GetIndex()]; ok {
			return pos, true
		}
	}

	return 0, false
}

func (this *ZendArray) FindBucket(key ZendArrayKey) *Bucket {
	if pos, ok := this.findPos(key); ok {
		return &this.data[pos]
	}
	return nil
}

func (this *ZendArray) Find(key ZendArrayKey) *Zval {
	if pos, ok := this.findPos(key); ok {
		return this.data[pos].GetVal()
	}
	return nil
}

func (this *ZendArray) Exists(key ZendArrayKey) bool {
	if _, ok := this.findPos(key); ok {
		return ok
	}
	return false
}

func (this *ZendArray) appendBucket(bucket *Bucket) *Bucket {
	var idx = uint32(len(this.data))

	this.nNumOfElements++
	this.data = append(this.data, *bucket)

	if bucket.IsStrKey() {
		var strKey = bucket.StrKey()
		this.keyMap[strKey] = idx
	} else {
		var indexKey = bucket.IndexKey()
		this.indexMap[indexKey] = idx

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

func (this *ZendArray) addOrUpdate(strKey string, pData *Zval, flag uint32) *Zval {
	this.assertRc1()

	var isAddNew = b.FlagMatch(flag, HASH_ADD_NEW)
	var isAdd = b.FlagMatch(flag, HASH_ADD)
	var isUpdateIndirect = b.FlagMatch(flag, HASH_UPDATE_INDIRECT)

	if !isAddNew {
		var p = this.FindBucketByStr(strKey)
		if p != nil {
			var data *Zval
			if isAdd {
				if !isUpdateIndirect {
					return nil
				}
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
			} else {
				ZEND_ASSERT(p.GetVal() != pData)
				data = p.GetVal()
				if isUpdateIndirect && data.IsType(IS_INDIRECT) {
					data = data.GetZv()
				}
			}
			if this.GetPDestructor() != nil {
				this.GetPDestructor()(data)
			}
			ZVAL_COPY_VALUE(data, pData)
			return data
		}
	}

	this.SubUFlags(HASH_FLAG_STATIC_KEYS)
	this.ifFullDoResize()

	var p = this.appendBucketStr(strKey, pData)
	return p.GetVal()
}

func (this *ZendArray) indexAddOrUpdate(indexKey int, pData *Zval, flag uint32) *Zval {
	this.assertRc1()

	var isAddNew = b.FlagMatch(flag, HASH_ADD_NEW)
	var isAdd = b.FlagMatch(flag, HASH_ADD)

	if !isAddNew {
		var p = this.FindBucketByIndex(indexKey)
		if p != nil {
			if isAdd {
				return nil
			}
			if this.pDestructor != nil {
				this.pDestructor(p.GetVal())
			}
			ZVAL_COPY_VALUE(p.GetVal(), pData)
			return p.GetVal()
		}
	}
	this.ifFullDoResize()

	var p = this.appendBucketIndex(indexKey, pData)

	return p.GetVal()
}
