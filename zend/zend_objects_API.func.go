package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
)

func ZendObjectStoreCtorFailed(obj *types2.ZendObject) {
	obj.AddGcFlags(types2.IS_OBJ_DESTRUCTOR_CALLED)
}
func ZendObjectPropertiesSize(ce *types2.ClassEntry) int {
	return b.SizeOf("zval") * (ce.GetDefaultPropertiesCount() - b.Cond(ce.IsUseGuards(), 0, 1))
}
func ZendObjectAlloc(obj_size int, ce *types2.ClassEntry) any {
	var obj any = Emalloc(obj_size + ZendObjectPropertiesSize(ce))

	/* Subtraction of sizeof(zval) is necessary, because zend_object_properties_size() may be
	 * -sizeof(zval), if the object has no properties. */

	memset(obj, 0, obj_size-b.SizeOf("zval"))
	return obj
}
func ZendGetPropertyInfoForSlot(obj *types2.ZendObject, slot *types2.Zval) *ZendPropertyInfo {
	var table **ZendPropertyInfo = obj.GetCe().GetPropertiesInfoTable()
	var prop_num intPtr = slot - obj.GetPropertiesTable()
	b.Assert(prop_num >= 0 && prop_num < obj.GetCe().GetDefaultPropertiesCount())
	return table[prop_num]
}
func ZendGetTypedPropertyInfoForSlot(obj *types2.ZendObject, slot *types2.Zval) *ZendPropertyInfo {
	var prop_info *ZendPropertyInfo = ZendGetPropertyInfoForSlot(obj, slot)
	if prop_info != nil && prop_info.GetType() != 0 {
		return prop_info
	}
	return nil
}
