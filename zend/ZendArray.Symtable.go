package zend

func (this *HashTable) SymtableClean() {
	// todo 这里可能不会严格对等，需要处理一下
	ZEND_ASSERT(this.pDestructor == ZVAL_PTR_DTOR)

	this.Clean()
}
