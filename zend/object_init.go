package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ObjectAndPropertiesInit(arg *types.Zval, class_type *types.ClassEntry, properties *types.Array) int {
	return _objectAndPropertiesInit(arg, class_type, properties)
}
func ObjectInitEx(arg *types.Zval, class_type *types.ClassEntry) int {
	return _objectAndPropertiesInit(arg, class_type, nil)
}
func ObjectInit(arg *types.Zval) int {
	arg.SetObject(ZendObjectsNew(ZendStandardClassDef))
	return types.SUCCESS
}
func _objectAndPropertiesInit(arg *types.Zval, classType *types.ClassEntry, properties *types.Array) int {
	if classType.HasCeFlags(AccInterface | AccTrait | AccImplicitAbstractClass | AccExplicitAbstractClass) {
		if classType.IsInterface() {
			faults.ThrowError(nil, "Cannot instantiate interface %s", classType.GetName().GetVal())
		} else if classType.IsTrait() {
			faults.ThrowError(nil, "Cannot instantiate trait %s", classType.GetName().GetVal())
		} else {
			faults.ThrowError(nil, "Cannot instantiate abstract class %s", classType.GetName().GetVal())
		}
		arg.SetNull()
		return types.FAILURE
	}
	if !classType.IsConstantsUpdated() {
		if ZendUpdateClassConstants(classType) != types.SUCCESS {
			arg.SetNull()
			return types.FAILURE
		}
	}
	if classType.GetCreateObject() == nil {
		var obj *types.ZendObject = ZendObjectsNew(classType)
		arg.SetObject(obj)
		if properties != nil {
			ObjectPropertiesInitEx(obj, properties)
		} else {
			_objectPropertiesInit(obj, classType)
		}
	} else {
		arg.SetObject(classType.GetCreateObject()(classType))
	}
	return types.SUCCESS
}
