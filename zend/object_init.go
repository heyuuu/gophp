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
	arg.SetObject(NewStdClassObject(nil))
	return types.SUCCESS
}
func _objectAndPropertiesInit(arg *types.Zval, classType *types.ClassEntry, properties *types.Array) int {
	if classType.HasCeFlags(types.AccInterface | types.AccTrait | types.AccImplicitAbstractClass | types.AccExplicitAbstractClass) {
		if classType.IsInterface() {
			faults.ThrowError(nil, fmt.Sprintf("Cannot instantiate interface %s", classType.Name()))
		} else if classType.IsTrait() {
			faults.ThrowError(nil, fmt.Sprintf("Cannot instantiate trait %s", classType.Name()))
		} else {
			faults.ThrowError(nil, fmt.Sprintf("Cannot instantiate abstract class %s", classType.Name()))
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
		var obj *types.Object
		if properties != nil {
			obj = types.NewStdObjectExEx(classType, properties)
		} else {
			obj = types.NewStdObject(classType)
		}
		arg.SetObject(obj)
	} else {
		arg.SetObject(classType.GetCreateObject()(classType))
	}
	return types.SUCCESS
}
