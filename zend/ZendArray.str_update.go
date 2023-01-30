package zend

// KeyAdd
func (this *ZendArray) KeyAdd(key string, pData *Zval) *Zval {
	this.assertRc1()
	if this.ExistsByStr(key) {
		return nil
	}

	this.SubUFlags(HASH_FLAG_STATIC_KEYS)
	var p = this.appendBucketStr(key, pData)
	return p.GetVal()
}

// KeyAddNew
func (this *ZendArray) KeyAddNew(key string, pData *Zval) *Zval {
	this.assertRc1()

	this.SubUFlags(HASH_FLAG_STATIC_KEYS)
	var p = this.appendBucketStr(key, pData)
	return p.GetVal()
}

func (this *ZendArray) KeyAddIndirect(strKey string, pData *Zval) *Zval {
	this.assertRc1()

	var p = this.FindBucketByStr(strKey)
	if p != nil {
		var data *Zval
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
		if this.GetPDestructor() != nil {
			this.GetPDestructor()(data)
		}
		ZVAL_COPY_VALUE(data, pData)
		return data
	}

	this.SubUFlags(HASH_FLAG_STATIC_KEYS)
	p = this.appendBucketStr(strKey, pData)
	return p.GetVal()
}

// KeyUpdate
func (this *ZendArray) KeyUpdate(key string, pData *Zval) *Zval {
	this.assertRc1()

	var p = this.FindBucketByStr(key)
	if p != nil {
		var data *Zval
		ZEND_ASSERT(p.GetVal() != pData)
		data = p.GetVal()
		if this.GetPDestructor() != nil {
			this.GetPDestructor()(data)
		}
		ZVAL_COPY_VALUE(data, pData)
		return data
	}

	this.SubUFlags(HASH_FLAG_STATIC_KEYS)
	p = this.appendBucketStr(key, pData)
	return p.GetVal()
}

// KeyUpdateIndirect
func (this *ZendArray) KeyUpdateIndirect(key string, pData *Zval) *Zval {
	this.assertRc1()

	var p = this.FindBucketByStr(key)
	if p != nil {
		var data *Zval
		ZEND_ASSERT(p.GetVal() != pData)
		data = p.GetVal()
		if data.IsType(IS_INDIRECT) {
			data = data.GetZv()
		}
		if this.GetPDestructor() != nil {
			this.GetPDestructor()(data)
		}
		ZVAL_COPY_VALUE(data, pData)
		return data
	}

	this.SubUFlags(HASH_FLAG_STATIC_KEYS)
	p = this.appendBucketStr(key, pData)
	return p.GetVal()
}
