package zend

import b "sik/builtin"

func (this *ZendArray) IsWithoutHoles() bool { return this.nNumUsed == this.nNumOfElements }

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

func (this *ZendArray) StrFindBucket(key string) *Bucket {
	if pos, ok := this.keyMap[key]; ok {
		return &this.data[pos]
	}
	return nil
}

func (this *ZendArray) IndexFindBucket(key int) *Bucket {
	if pos, ok := this.indexMap[key]; ok {
		return &this.data[pos]
	}
	return nil
}

func (this *ZendArray) _addBucket(strKey string, zv *Zval) *Bucket {
	var bucket = NewBucketStr(strKey, zv)
	var idx = this.nNumUsed

	this.nNumUsed++
	this.nNumOfElements++
	this.data = append(this.data, *bucket)

	this.keyMap[strKey] = idx

	return &this.data[idx]
}

func (this *ZendArray) _addBucketIndex(indexKey int, zv *Zval) *Bucket {
	var bucket = NewBucketIndex(indexKey, zv)
	var idx = this.nNumUsed

	this.nNumUsed++
	this.nNumOfElements++
	this.data = append(this.data, *bucket)

	this.indexMap[indexKey] = idx

	// 更新 nNextFreeElement
	if indexKey > this.nNextFreeElement {
		if indexKey < ZEND_LONG_MAX {
			this.nNextFreeElement = indexKey + 1
		} else {
			this.nNextFreeElement = ZEND_LONG_MAX
		}
	}

	return &this.data[idx]
}

func (this *ZendArray) addOrUpdate(strKey string, pData *Zval, flag uint32) *Zval {
	this.assertRc1()

	var isAddNew = b.FlagMatch(flag, HASH_ADD_NEW)
	var isAdd = b.FlagMatch(flag, HASH_ADD)
	var isUpdateIndirect = b.FlagMatch(flag, HASH_UPDATE_INDIRECT)

	if !isAddNew {
		var p = this.StrFindBucket(strKey)
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
	ZEND_HASH_IF_FULL_DO_RESIZE(this)

	var p = this._addBucket(strKey, pData)
	return p.GetVal()
}

func (this *ZendArray) indexAddOrUpdate(indexKey int, pData *Zval, flag uint32) *Zval {
	this.assertRc1()

	var isAddNew = b.FlagMatch(flag, HASH_ADD_NEW)
	var isAdd = b.FlagMatch(flag, HASH_ADD)

	if !isAddNew {
		var p = this.IndexFindBucket(indexKey)
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
	ZEND_HASH_IF_FULL_DO_RESIZE(this)

	var p = this._addBucketIndex(indexKey, pData)

	return p.GetVal()
}
