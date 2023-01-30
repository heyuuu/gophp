package zend

// IndexAdd
func (this *ZendArray) IndexAddH(h ZendUlong, pData *Zval) *Zval {
	return this.IndexAdd(int(h), pData)
}
func (this *ZendArray) IndexAdd(index int, pData *Zval) *Zval {
	this.assertRc1()

	if this.ExistsByIndex(index) {
		return nil
	}

	var p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}

// IndexAddNew
func (this *ZendArray) IndexAddNewH(h ZendUlong, pData *Zval) *Zval {
	return this.IndexAddNew(int(h), pData)
}
func (this *ZendArray) IndexAddNew(index int, pData *Zval) *Zval {
	this.assertRc1()

	var p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}

// IndexUpdate
func (this *ZendArray) IndexUpdateH(h ZendUlong, pData *Zval) *Zval {
	return this.IndexUpdate(int(h), pData)
}
func (this *ZendArray) IndexUpdate(index int, pData *Zval) *Zval {
	this.assertRc1()

	var p *Bucket

	p = this.FindBucketByIndex(index)
	if p != nil {
		if this.pDestructor != nil {
			this.pDestructor(p.GetVal())
		}
		ZVAL_COPY_VALUE(p.GetVal(), pData)
		return p.GetVal()
	}

	p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}

// NextIndexInsert
func (this *ZendArray) NextIndexInsert(pData *Zval) *Zval {
	this.assertRc1()

	var index = this.nNextFreeElement

	if this.ExistsByIndex(index) {
		return nil
	}

	var p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}

// NextIndexInsertNew
func (this *ZendArray) NextIndexInsertNew(pData *Zval) *Zval {
	this.assertRc1()

	var index = this.nNextFreeElement
	var p = this.appendBucketIndex(index, pData)
	return p.GetVal()
}
