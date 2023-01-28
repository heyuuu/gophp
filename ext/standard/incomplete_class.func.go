// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func IncompleteClassMessage(object *zend.Zval, error_type int) {
	var class_name *zend.ZendString
	class_name = PhpLookupClassName(object)
	if class_name != nil {
		core.PhpErrorDocref(nil, error_type, INCOMPLETE_CLASS_MSG, class_name.GetVal())
		zend.ZendStringReleaseEx(class_name, 0)
	} else {
		core.PhpErrorDocref(nil, error_type, INCOMPLETE_CLASS_MSG, "unknown")
	}
}
func IncompleteClassGetProperty(object *zend.Zval, member *zend.Zval, type_ int, cache_slot *any, rv *zend.Zval) *zend.Zval {
	IncompleteClassMessage(object, zend.E_NOTICE)
	if type_ == zend.BP_VAR_W || type_ == zend.BP_VAR_RW {
		zend.ZVAL_ERROR(rv)
		return rv
	} else {
		return &(zend.ExecutorGlobals.GetUninitializedZval())
	}
}
func IncompleteClassWriteProperty(object *zend.Zval, member *zend.Zval, value *zend.Zval, cache_slot *any) *zend.Zval {
	IncompleteClassMessage(object, zend.E_NOTICE)
	return value
}
func IncompleteClassGetPropertyPtrPtr(object *zend.Zval, member *zend.Zval, type_ int, cache_slot *any) *zend.Zval {
	IncompleteClassMessage(object, zend.E_NOTICE)
	return &(zend.ExecutorGlobals.GetErrorZval())
}
func IncompleteClassUnsetProperty(object *zend.Zval, member *zend.Zval, cache_slot *any) {
	IncompleteClassMessage(object, zend.E_NOTICE)
}
func IncompleteClassHasProperty(object *zend.Zval, member *zend.Zval, check_empty int, cache_slot *any) int {
	IncompleteClassMessage(object, zend.E_NOTICE)
	return 0
}
func IncompleteClassGetMethod(object **zend.ZendObject, method *zend.ZendString, key *zend.Zval) *zend.ZendFunction {
	var zobject zend.Zval
	zend.ZVAL_OBJ(&zobject, *object)
	IncompleteClassMessage(&zobject, zend.E_ERROR)
	return nil
}
func PhpCreateIncompleteObject(class_type *zend.ZendClassEntry) *zend.ZendObject {
	var object *zend.ZendObject
	object = zend.ZendObjectsNew(class_type)
	object.SetHandlers(&PhpIncompleteObjectHandlers)
	zend.ObjectPropertiesInit(object, class_type)
	return object
}
func PhpCreateIncompleteClass() *zend.ZendClassEntry {
	var incomplete_class zend.ZendClassEntry
	memset(&incomplete_class, 0, b.SizeOf("zend_class_entry"))
	incomplete_class.SetName(zend.ZendStringInitInterned(INCOMPLETE_CLASS, b.SizeOf("INCOMPLETE_CLASS")-1, 1))
	incomplete_class.SetBuiltinFunctions(nil)
	incomplete_class.create_object = PhpCreateIncompleteObject
	memcpy(&PhpIncompleteObjectHandlers, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	PhpIncompleteObjectHandlers.SetReadProperty(IncompleteClassGetProperty)
	PhpIncompleteObjectHandlers.SetHasProperty(IncompleteClassHasProperty)
	PhpIncompleteObjectHandlers.SetUnsetProperty(IncompleteClassUnsetProperty)
	PhpIncompleteObjectHandlers.SetWriteProperty(IncompleteClassWriteProperty)
	PhpIncompleteObjectHandlers.SetGetPropertyPtrPtr(IncompleteClassGetPropertyPtrPtr)
	PhpIncompleteObjectHandlers.SetGetMethod(IncompleteClassGetMethod)
	return zend.ZendRegisterInternalClass(&incomplete_class)
}
func PhpLookupClassName(object *zend.Zval) *zend.ZendString {
	var val *zend.Zval
	var object_properties *zend.HashTable
	object_properties = zend.Z_OBJPROP_P(object)
	if b.Assign(&val, zend.ZendHashStrFind(object_properties, MAGIC_MEMBER, b.SizeOf("MAGIC_MEMBER")-1)) != nil && val.GetType() == zend.IS_STRING {
		return zend.ZendStringCopy(zend.Z_STR_P(val))
	}
	return nil
}
func PhpStoreClassName(object *zend.Zval, name *byte, len_ int) {
	var val zend.Zval
	zend.ZVAL_STRINGL(&val, name, len_)
	zend.ZendHashStrUpdate(zend.Z_OBJPROP_P(object), MAGIC_MEMBER, b.SizeOf("MAGIC_MEMBER")-1, &val)
}
