// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

func ZendTryAssignTypedRefNull(ref *ZendReference) int {
	var tmp Zval
	tmp.SetNull()
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefBool(ref *ZendReference, val ZendBool) int {
	var tmp Zval
	ZVAL_BOOL(&tmp, val != 0)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefLong(ref *ZendReference, lval ZendLong) int {
	var tmp Zval
	tmp.SetLong(lval)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefDouble(ref *ZendReference, dval float64) int {
	var tmp Zval
	tmp.SetDouble(dval)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefEmptyString(ref *ZendReference) int {
	var tmp Zval
	ZVAL_EMPTY_STRING(&tmp)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefStr(ref *ZendReference, str *ZendString) int {
	var tmp Zval
	tmp.SetString(str)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefString(ref *ZendReference, string *byte) int {
	var tmp Zval
	ZVAL_STRING(&tmp, string)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefStringl(ref *ZendReference, string *byte, len_ int) int {
	var tmp Zval
	ZVAL_STRINGL(&tmp, string, len_)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefArr(ref *ZendReference, arr *ZendArray) int {
	var tmp Zval
	tmp.SetArray(arr)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefRes(ref *ZendReference, res *ZendResource) int {
	var tmp Zval
	tmp.SetResource(res)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefZval(ref *ZendReference, zv *Zval) int {
	var tmp Zval
	ZVAL_COPY_VALUE(&tmp, zv)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefZvalEx(ref *ZendReference, zv *Zval, strict ZendBool) int {
	var tmp Zval
	ZVAL_COPY_VALUE(&tmp, zv)
	return ZendTryAssignTypedRefEx(ref, &tmp, strict)
}
func ZendDeclarePropertyEx(ce *ZendClassEntry, name *ZendString, property *Zval, access_type int, doc_comment *ZendString) int {
	return ZendDeclareTypedProperty(ce, name, property, access_type, doc_comment, 0)
}
func ZendDeclareProperty(ce *ZendClassEntry, name *byte, name_length int, property *Zval, access_type int) int {
	var key *ZendString = ZendStringInit(name, name_length, IsPersistentClass(ce))
	var ret int = ZendDeclarePropertyEx(ce, key, property, access_type, nil)
	ZendStringRelease(key)
	return ret
}
func ZendDeclarePropertyNull(ce *ZendClassEntry, name string, name_length int, access_type int) int {
	var property Zval
	property.SetNull()
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyBool(ce *ZendClassEntry, name *byte, name_length int, value ZendLong, access_type int) int {
	var property Zval
	ZVAL_BOOL(&property, value != 0)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyLong(ce *ZendClassEntry, name string, name_length int, value ZendLong, access_type int) int {
	var property Zval
	property.SetLong(value)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyDouble(ce *ZendClassEntry, name *byte, name_length int, value float64, access_type int) int {
	var property Zval
	property.SetDouble(value)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyString(ce *ZendClassEntry, name string, name_length int, value string, access_type int) int {
	var property Zval
	property.SetString(ZendStringInit(value, strlen(value), ce.GetType()&ZEND_INTERNAL_CLASS))
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyStringl(
	ce *ZendClassEntry,
	name *byte,
	name_length int,
	value *byte,
	value_len int,
	access_type int,
) int {
	var property Zval
	property.SetString(ZendStringInit(value, value_len, ce.GetType()&ZEND_INTERNAL_CLASS))
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclareClassConstantEx(ce *ZendClassEntry, name *ZendString, value *Zval, access_type int, doc_comment *ZendString) int {
	var c *ZendClassConstant
	if ce.IsInterface() {
		if access_type != ZEND_ACC_PUBLIC {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Access type for interface constant %s::%s must be public", ce.GetName().GetVal(), name.GetVal())
		}
	}
	if ZendStringEqualsLiteralCi(name, "class") {
		ZendErrorNoreturn(b.Cond(ce.GetType() == ZEND_INTERNAL_CLASS, E_CORE_ERROR, E_COMPILE_ERROR), "A class constant must not be called 'class'; it is reserved for class name fetching")
	}
	if value.IsString() {
		ZvalMakeInternedString(value)
	}
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		c = Pemalloc(b.SizeOf("zend_class_constant"), 1)
	} else {
		c = ZendArenaAlloc(CG__().GetArena(), b.SizeOf("zend_class_constant"))
	}
	ZVAL_COPY_VALUE(c.GetValue(), value)
	c.GetValue().GetAccessFlags() = access_type
	c.SetDocComment(doc_comment)
	c.SetCe(ce)
	if value.IsConstant() {
		ce.SetIsConstantsUpdated(false)
	}
	if !(ZendHashAddPtr(ce.GetConstantsTable(), name, c)) {
		ZendErrorNoreturn(b.Cond(ce.GetType() == ZEND_INTERNAL_CLASS, E_CORE_ERROR, E_COMPILE_ERROR), "Cannot redefine class __special__  constant %s::%s", ce.GetName().GetVal(), name.GetVal())
	}
	return SUCCESS
}
func ZendDeclareClassConstant(ce *ZendClassEntry, name *byte, name_length int, value *Zval) int {
	var ret int
	var key *ZendString
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		key = ZendStringInitInterned(name, name_length, 1)
	} else {
		key = ZendStringInit(name, name_length, 0)
	}
	ret = ZendDeclareClassConstantEx(ce, key, value, ZEND_ACC_PUBLIC, nil)
	ZendStringRelease(key)
	return ret
}
func ZendDeclareClassConstantNull(ce *ZendClassEntry, name *byte, name_length int) int {
	var constant Zval
	constant.SetNull()
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantLong(ce *ZendClassEntry, name string, name_length int, value ZendLong) int {
	var constant Zval
	constant.SetLong(value)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantBool(ce *ZendClassEntry, name *byte, name_length int, value ZendBool) int {
	var constant Zval
	ZVAL_BOOL(&constant, value != 0)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantDouble(ce *ZendClassEntry, name *byte, name_length int, value float64) int {
	var constant Zval
	constant.SetDouble(value)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantStringl(ce *ZendClassEntry, name *byte, name_length int, value *byte, value_length int) int {
	var constant Zval
	constant.SetString(ZendStringInit(value, value_length, ce.GetType()&ZEND_INTERNAL_CLASS))
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantString(ce *ZendClassEntry, name *byte, name_length int, value *byte) int {
	return ZendDeclareClassConstantStringl(ce, name, name_length, value, strlen(value))
}
func ZendUpdatePropertyEx(scope *ZendClassEntry, object *Zval, name *ZendString, value *Zval) {
	var property Zval
	var old_scope *ZendClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property.SetString(name)
	Z_OBJ_HT_P(object).GetWriteProperty()(object, &property, value, nil)
	EG__().SetFakeScope(old_scope)
}
func ZendUpdateProperty(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value *Zval) {
	var property Zval
	var old_scope *ZendClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	ZVAL_STRINGL(&property, name, name_length)
	Z_OBJ_HT_P(object).GetWriteProperty()(object, &property, value, nil)
	ZvalPtrDtor(&property)
	EG__().SetFakeScope(old_scope)
}
func ZendUpdatePropertyNull(scope *ZendClassEntry, object *Zval, name *byte, name_length int) {
	var tmp Zval
	tmp.SetNull()
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUnsetProperty(scope *ZendClassEntry, object *Zval, name string, name_length int) {
	var property Zval
	var old_scope *ZendClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	ZVAL_STRINGL(&property, name, name_length)
	Z_OBJ_HT_P(object).GetUnsetProperty()(object, &property, 0)
	ZvalPtrDtor(&property)
	EG__().SetFakeScope(old_scope)
}
func ZendUpdatePropertyBool(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value ZendLong) {
	var tmp Zval
	ZVAL_BOOL(&tmp, value != 0)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyLong(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value ZendLong) {
	var tmp Zval
	tmp.SetLong(value)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyDouble(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value float64) {
	var tmp Zval
	tmp.SetDouble(value)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyStr(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value *ZendString) {
	var tmp Zval
	tmp.SetString(value)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyString(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value *byte) {
	var tmp Zval
	ZVAL_STRING(&tmp, value)
	tmp.SetRefcount(0)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyStringl(
	scope *ZendClassEntry,
	object *Zval,
	name *byte,
	name_length int,
	value *byte,
	value_len int,
) {
	var tmp Zval
	ZVAL_STRINGL(&tmp, value, value_len)
	tmp.SetRefcount(0)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyEx(scope *ZendClassEntry, name *ZendString, value *Zval) int {
	var property *Zval
	var tmp Zval
	var prop_info *ZendPropertyInfo
	var old_scope *ZendClassEntry = EG__().GetFakeScope()
	if !scope.IsConstantsUpdated() {
		if ZendUpdateClassConstants(scope) != SUCCESS {
			return FAILURE
		}
	}
	EG__().SetFakeScope(scope)
	property = ZendStdGetStaticPropertyWithInfo(scope, name, BP_VAR_W, &prop_info)
	EG__().SetFakeScope(old_scope)
	if property == nil {
		return FAILURE
	}
	ZEND_ASSERT(!(value.IsReference()))
	value.TryAddRefcount()
	if prop_info.GetType() != 0 {
		ZVAL_COPY_VALUE(&tmp, value)
		if ZendVerifyPropertyType(prop_info, &tmp, 0) == 0 {
			value.TryDelRefcount()
			return FAILURE
		}
		value = &tmp
	}
	ZendAssignToVariable(property, value, IS_TMP_VAR, 0)
	return SUCCESS
}
func ZendUpdateStaticProperty(scope *ZendClassEntry, name *byte, name_length int, value *Zval) int {
	var key *ZendString = ZendStringInit(name, name_length, 0)
	var retval int = ZendUpdateStaticPropertyEx(scope, key, value)
	ZendStringEfree(key)
	return retval
}
func ZendUpdateStaticPropertyNull(scope *ZendClassEntry, name *byte, name_length int) int {
	var tmp Zval
	tmp.SetNull()
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyBool(scope *ZendClassEntry, name *byte, name_length int, value ZendLong) int {
	var tmp Zval
	ZVAL_BOOL(&tmp, value != 0)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyLong(scope *ZendClassEntry, name *byte, name_length int, value ZendLong) int {
	var tmp Zval
	tmp.SetLong(value)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyDouble(scope *ZendClassEntry, name *byte, name_length int, value float64) int {
	var tmp Zval
	tmp.SetDouble(value)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyString(scope *ZendClassEntry, name *byte, name_length int, value *byte) int {
	var tmp Zval
	ZVAL_STRING(&tmp, value)
	tmp.SetRefcount(0)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyStringl(scope *ZendClassEntry, name *byte, name_length int, value *byte, value_len int) int {
	var tmp Zval
	ZVAL_STRINGL(&tmp, value, value_len)
	tmp.SetRefcount(0)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendReadPropertyEx(scope *ZendClassEntry, object *Zval, name *ZendString, silent ZendBool, rv *Zval) *Zval {
	var property Zval
	var value *Zval
	var old_scope *ZendClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property.SetString(name)
	value = Z_OBJ_HT_P(object).GetReadProperty()(object, &property, b.Cond(silent != 0, BP_VAR_IS, BP_VAR_R), nil, rv)
	EG__().SetFakeScope(old_scope)
	return value
}
func ZendReadProperty(
	scope *ZendClassEntry,
	object *Zval,
	name string,
	name_length int,
	silent ZendBool,
	rv *Zval,
) *Zval {
	var value *Zval
	var str *ZendString
	str = ZendStringInit(name, name_length, 0)
	value = ZendReadPropertyEx(scope, object, str, silent, rv)
	ZendStringReleaseEx(str, 0)
	return value
}
func ZendReadStaticPropertyEx(scope *ZendClassEntry, name *ZendString, silent ZendBool) *Zval {
	var property *Zval
	var old_scope *ZendClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property = ZendStdGetStaticProperty(scope, name, b.Cond(silent != 0, BP_VAR_IS, BP_VAR_R))
	EG__().SetFakeScope(old_scope)
	return property
}
func ZendReadStaticProperty(scope *ZendClassEntry, name *byte, name_length int, silent ZendBool) *Zval {
	var key *ZendString = ZendStringInit(name, name_length, 0)
	var property *Zval = ZendReadStaticPropertyEx(scope, key, silent)
	ZendStringEfree(key)
	return property
}
func ZendSaveErrorHandling(current *ZendErrorHandling) {
	current.SetHandling(EG__().GetErrorHandling())
	current.SetException(EG__().GetExceptionClass())
	current.GetUserHandler().SetUndef()
}
func ZendReplaceErrorHandling(error_handling ZendErrorHandlingT, exception_class *ZendClassEntry, current *ZendErrorHandling) {
	if current != nil {
		ZendSaveErrorHandling(current)
	}
	ZEND_ASSERT(error_handling == EH_THROW || exception_class == nil)
	EG__().SetErrorHandling(error_handling)
	EG__().SetExceptionClass(exception_class)
}
func ZendRestoreErrorHandling(saved *ZendErrorHandling) {
	EG__().SetErrorHandling(saved.GetHandling())
	EG__().SetExceptionClass(saved.GetException())
}
func ZendFindAliasName(ce *ZendClassEntry, name *ZendString) *ZendString {
	var alias *ZendTraitAlias
	var alias_ptr **ZendTraitAlias
	if b.Assign(&alias_ptr, ce.GetTraitAliases()) {
		alias = *alias_ptr
		for alias != nil {
			if alias.GetAlias() != nil && ZendStringEqualsCi(alias.GetAlias(), name) {
				return alias.GetAlias()
			}
			alias_ptr++
			alias = *alias_ptr
		}
	}
	return name
}
func ZendResolveMethodName(ce *ZendClassEntry, f *ZendFunction) *ZendString {
	var func_ *ZendFunction
	var function_table *HashTable
	var name *ZendString
	if f.GetCommonType() != ZEND_USER_FUNCTION || f.GetOpArray().GetRefcount() != nil && (*(f.GetOpArray().GetRefcount())) < 2 || f.GetScope() == nil || f.GetScope().GetTraitAliases() == nil {
		return f.GetFunctionName()
	}
	function_table = ce.GetFunctionTable()
	var __ht *HashTable = function_table
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		name = _p.GetKey()
		func_ = _z.GetPtr()
		if func_ == f {
			if name == nil {
				return f.GetFunctionName()
			}
			if name.GetLen() == f.GetFunctionName().GetLen() && !(strncasecmp(name.GetVal(), f.GetFunctionName().GetVal(), f.GetFunctionName().GetLen())) {
				return f.GetFunctionName()
			}
			return ZendFindAliasName(f.GetScope(), name)
		}
	}
	return f.GetFunctionName()
}
func ZendGetObjectType(ce *ZendClassEntry) *byte {
	if ce.IsTrait() {
		return "trait"
	} else if ce.IsInterface() {
		return "interface"
	} else {
		return "class"
	}
}
func ZendIsIterable(iterable *Zval) ZendBool {
	switch iterable.GetType() {
	case IS_ARRAY:
		return 1
	case IS_OBJECT:
		return InstanceofFunction(Z_OBJCE_P(iterable), ZendCeTraversable)
	default:
		return 0
	}
}
func ZendIsCountable(countable *Zval) ZendBool {
	switch countable.GetType() {
	case IS_ARRAY:
		return 1
	case IS_OBJECT:
		if Z_OBJ_HT_P(countable).GetCountElements() != nil {
			return 1
		}
		return InstanceofFunction(Z_OBJCE_P(countable), ZendCeCountable)
	default:
		return 0
	}
}
