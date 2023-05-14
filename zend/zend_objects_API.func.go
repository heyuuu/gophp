package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func ZendObjectStoreCtorFailed(obj *types.ZendObject) {
	obj.MarkObjDtorCalled()
}
func ZendObjectPropertiesSize(ce *types.ClassEntry) int {
	return b.SizeOf("zval") * (ce.GetDefaultPropertiesCount() - b.Cond(ce.IsUseGuards(), 0, 1))
}
func ZendObjectAlloc(obj_size int, ce *types.ClassEntry) any {
	var obj any = Emalloc(obj_size + ZendObjectPropertiesSize(ce))

	/* Subtraction of sizeof(zval) is necessary, because zend_object_properties_size() may be
	 * -sizeof(zval), if the object has no properties. */

	memset(obj, 0, obj_size-b.SizeOf("zval"))
	return obj
}

func ZendGetPropertyInfoForSlot(obj *types.ZendObject, slot *types.Zval) *ZendPropertyInfo {
	ce := obj.GetCe()
	propNum := slot - obj.GetPropertiesTable()
	return ce.GetPropertyInfo(propNum)
}

func ZendGetTypedPropertyInfoForSlot(obj *types.ZendObject, slot *types.Zval) *ZendPropertyInfo {
	var propInfo *ZendPropertyInfo = ZendGetPropertyInfoForSlot(obj, slot)
	if propInfo != nil && propInfo.GetType() != 0 {
		return propInfo
	}
	return nil
}
