package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

func IncompleteClassMessage(object *types.Zval, error_type int) {
	var class_name *types.String
	class_name = PhpLookupClassName(object)
	if class_name != nil {
		core.PhpErrorDocref(nil, error_type, INCOMPLETE_CLASS_MSG, class_name.GetVal())
		// types.ZendStringReleaseEx(class_name, 0)
	} else {
		core.PhpErrorDocref(nil, error_type, INCOMPLETE_CLASS_MSG, "unknown")
	}
}
func IncompleteClassGetProperty(object *types.Zval, member *types.Zval, type_ int, cache_slot *any, rv *types.Zval) *types.Zval {
	IncompleteClassMessage(object, faults.E_NOTICE)
	if type_ == zend.BP_VAR_W || type_ == zend.BP_VAR_RW {
		rv.IsError()
		return rv
	} else {
		return zend.UninitializedZval()
	}
}
func IncompleteClassWriteProperty(object *types.Zval, member *types.Zval, value *types.Zval, cache_slot *any) *types.Zval {
	IncompleteClassMessage(object, faults.E_NOTICE)
	return value
}
func IncompleteClassGetPropertyPtrPtr(object *types.Zval, member *types.Zval, type_ int, cache_slot *any) *types.Zval {
	IncompleteClassMessage(object, faults.E_NOTICE)
	return zend.EG__().GetErrorZval()
}
func IncompleteClassUnsetProperty(object *types.Zval, member *types.Zval, cache_slot *any) {
	IncompleteClassMessage(object, faults.E_NOTICE)
}
func IncompleteClassHasProperty(object *types.Zval, member *types.Zval, check_empty int, cache_slot *any) int {
	IncompleteClassMessage(object, faults.E_NOTICE)
	return 0
}
func IncompleteClassGetMethod(object **types.Object, method *types.String, key *types.Zval) types.IFunction {
	var zobject types.Zval
	zobject.SetObject(*object)
	IncompleteClassMessage(&zobject, faults.E_ERROR)
	return nil
}
func PhpCreateIncompleteObject(ce *types.ClassEntry) *types.Object {
	return types.NewObject(ce, &PhpIncompleteObjectHandlers)
}
func PhpCreateIncompleteClass() *types.ClassEntry {
	PhpIncompleteObjectHandlers = *types.NewObjectHandlersEx(zend.StdObjectHandlersPtr, types.ObjectHandlersSetting{
		ReadProperty:      IncompleteClassGetProperty,
		HasProperty:       IncompleteClassHasProperty,
		UnsetProperty:     IncompleteClassUnsetProperty,
		WriteProperty:     IncompleteClassWriteProperty,
		GetPropertyPtrPtr: IncompleteClassGetPropertyPtrPtr,
		GetMethod:         IncompleteClassGetMethod,
	})

	return zend.RegisterClass(INCOMPLETE_CLASS, PhpCreateIncompleteObject, nil)
}
func PhpLookupClassName(object *types.Zval) *types.String {
	var val *types.Zval
	var object_properties *types.Array
	object_properties = types.Z_OBJPROP_P(object)
	if lang.Assign(&val, object_properties.KeyFind(b.CastStrAuto(MAGIC_MEMBER))) != nil && val.IsString() {
		return val.String().Copy()
	}
	return nil
}
func PhpStoreClassName(object *types.Zval, name *byte, len_ int) {
	var val types.Zval
	val.SetString(b.CastStr(name, len_))
	types.Z_OBJPROP_P(object).KeyUpdate(b.CastStrAuto(MAGIC_MEMBER), &val)
}
