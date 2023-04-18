package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

func IncompleteClassMessage(object *types2.Zval, error_type int) {
	var class_name *types2.String
	class_name = PhpLookupClassName(object)
	if class_name != nil {
		core.PhpErrorDocref(nil, error_type, INCOMPLETE_CLASS_MSG, class_name.GetVal())
		// types.ZendStringReleaseEx(class_name, 0)
	} else {
		core.PhpErrorDocref(nil, error_type, INCOMPLETE_CLASS_MSG, "unknown")
	}
}
func IncompleteClassGetProperty(object *types2.Zval, member *types2.Zval, type_ int, cache_slot *any, rv *types2.Zval) *types2.Zval {
	IncompleteClassMessage(object, faults.E_NOTICE)
	if type_ == zend.BP_VAR_W || type_ == zend.BP_VAR_RW {
		rv.IsError()
		return rv
	} else {
		return zend.EG__().GetUninitializedZval()
	}
}
func IncompleteClassWriteProperty(object *types2.Zval, member *types2.Zval, value *types2.Zval, cache_slot *any) *types2.Zval {
	IncompleteClassMessage(object, faults.E_NOTICE)
	return value
}
func IncompleteClassGetPropertyPtrPtr(object *types2.Zval, member *types2.Zval, type_ int, cache_slot *any) *types2.Zval {
	IncompleteClassMessage(object, faults.E_NOTICE)
	return zend.EG__().GetErrorZval()
}
func IncompleteClassUnsetProperty(object *types2.Zval, member *types2.Zval, cache_slot *any) {
	IncompleteClassMessage(object, faults.E_NOTICE)
}
func IncompleteClassHasProperty(object *types2.Zval, member *types2.Zval, check_empty int, cache_slot *any) int {
	IncompleteClassMessage(object, faults.E_NOTICE)
	return 0
}
func IncompleteClassGetMethod(object **types2.ZendObject, method *types2.String, key *types2.Zval) types2.IFunction {
	var zobject types2.Zval
	zobject.SetObject(*object)
	IncompleteClassMessage(&zobject, faults.E_ERROR)
	return nil
}
func PhpCreateIncompleteObject(class_type *types2.ClassEntry) *types2.ZendObject {
	var object *types2.ZendObject
	object = zend.ZendObjectsNew(class_type)
	object.SetHandlers(&PhpIncompleteObjectHandlers)
	zend.ObjectPropertiesInit(object, class_type)
	return object
}
func PhpCreateIncompleteClass() *types2.ClassEntry {
	var incomplete_class types2.ClassEntry
	memset(&incomplete_class, 0, b.SizeOf("zend_class_entry"))
	incomplete_class.SetNameVal(INCOMPLETE_CLASS)
	incomplete_class.SetBuiltinFunctions(nil)
	incomplete_class.SetCreateObject(PhpCreateIncompleteObject)
	memcpy(&PhpIncompleteObjectHandlers, zend.StdObjectHandlersPtr, b.SizeOf("zend_object_handlers"))
	PhpIncompleteObjectHandlers.SetReadProperty(IncompleteClassGetProperty)
	PhpIncompleteObjectHandlers.SetHasProperty(IncompleteClassHasProperty)
	PhpIncompleteObjectHandlers.SetUnsetProperty(IncompleteClassUnsetProperty)
	PhpIncompleteObjectHandlers.SetWriteProperty(IncompleteClassWriteProperty)
	PhpIncompleteObjectHandlers.SetGetPropertyPtrPtr(IncompleteClassGetPropertyPtrPtr)
	PhpIncompleteObjectHandlers.SetGetMethod(IncompleteClassGetMethod)
	return zend.ZendRegisterInternalClass(&incomplete_class)
}
func PhpLookupClassName(object *types2.Zval) *types2.String {
	var val *types2.Zval
	var object_properties *types2.Array
	object_properties = types2.Z_OBJPROP_P(object)
	if b.Assign(&val, object_properties.KeyFind(b.CastStrAuto(MAGIC_MEMBER))) != nil && val.IsType(types2.IS_STRING) {
		return val.String().Copy()
	}
	return nil
}
func PhpStoreClassName(object *types2.Zval, name *byte, len_ int) {
	var val types2.Zval
	val.SetStringVal(b.CastStr(name, len_))
	types2.Z_OBJPROP_P(object).KeyUpdate(b.CastStrAuto(MAGIC_MEMBER), &val)
}
