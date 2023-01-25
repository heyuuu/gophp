// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendRegisterIteratorWrapper() {
	memset(&ZendIteratorClassEntry, 0, b.SizeOf("zend_class_entry"))
	ZendIteratorClassEntry.SetName(ZendStringInitInterned("__iterator_wrapper", b.SizeOf("\"__iterator_wrapper\"")-1, 1))
	ZendIteratorClassEntry.SetBuiltinFunctions(nil)
}
func IterWrapperFree(object *ZendObject) {
	var iter *ZendObjectIterator = (*ZendObjectIterator)(object)
	iter.GetFuncs().GetDtor()(iter)
}
func IterWrapperDtor(object *ZendObject) {}
func IterWrapperGetGc(object *Zval, table **Zval, n *int) *HashTable {
	/* TODO: We need a get_gc iterator handler */

	*table = nil
	*n = 0
	return nil
}
func ZendIteratorInit(iter *ZendObjectIterator) {
	ZendObjectStdInit(&iter.GetStd(), &ZendIteratorClassEntry)
	iter.GetStd().SetHandlers(&IteratorObjectHandlers)
}
func ZendIteratorDtor(iter *ZendObjectIterator) {
	if GC_DELREF(&iter.GetStd()) > 0 {
		return
	}
	ZendObjectsStoreDel(&iter.GetStd())
}
func ZendIteratorUnwrap(array_ptr *Zval) *ZendObjectIterator {
	ZEND_ASSERT(Z_TYPE_P(array_ptr) == IS_OBJECT)
	if Z_OBJ_HT_P(array_ptr) == &IteratorObjectHandlers {
		return (*ZendObjectIterator)(Z_OBJ_P(array_ptr))
	}
	return nil
}
