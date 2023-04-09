package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/types"
)

func ZendRegisterIteratorWrapper() {
	memset(&ZendIteratorClassEntry, 0, b.SizeOf("zend_class_entry"))
	ZendIteratorClassEntry.SetName(types.NewString("__iterator_wrapper"))
	ZendIteratorClassEntry.SetBuiltinFunctions(nil)
}
func IterWrapperFree(object *types.ZendObject) {
	var iter *ZendObjectIterator = (*ZendObjectIterator)(object)
	iter.GetFuncs().GetDtor()(iter)
}
func IterWrapperDtor(object *types.ZendObject) {}
func IterWrapperGetGc(object *types.Zval, table **types.Zval, n *int) *types.Array {
	/* TODO: We need a get_gc iterator handler */

	*table = nil
	*n = 0
	return nil
}
func ZendIteratorInit(iter *ZendObjectIterator) {
	ZendObjectStdInit(iter.GetStd(), &ZendIteratorClassEntry)
	iter.GetStd().SetHandlers(&IteratorObjectHandlers)
}
func ZendIteratorDtor(iter *ZendObjectIterator) {
	if iter.GetStd().DelRefcount() > 0 {
		return
	}
	ZendObjectsStoreDel(iter.GetStd())
}
func ZendIteratorUnwrap(array_ptr *types.Zval) *ZendObjectIterator {
	b.Assert(array_ptr.IsObject())
	if types.Z_OBJ_HT_P(array_ptr) == &IteratorObjectHandlers {
		return (*ZendObjectIterator)(array_ptr.GetObj())
	}
	return nil
}
