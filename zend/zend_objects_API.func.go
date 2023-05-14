package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZendObjectStoreCtorFailed(obj *types.ZendObject) {
	obj.MarkObjDtorCalled()
}

func ZendGetPropertyInfoForSlot(obj *types.ZendObject, slot *types.Zval) *types.PropertyInfo {
	ce := obj.GetCe()
	propNum := slot - obj.GetPropertiesTable()
	return ce.GetPropertyInfo(propNum)
}

func ZendGetTypedPropertyInfoForSlot(obj *types.ZendObject, slot *types.Zval) *types.PropertyInfo {
	var propInfo *types.PropertyInfo = ZendGetPropertyInfoForSlot(obj, slot)
	if propInfo != nil && propInfo.GetType() != 0 {
		return propInfo
	}
	return nil
}
