package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZendTryAssignTypedRefNull(ref *types.ZendReference) int {
	var tmp types.Zval
	tmp.SetNull()
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefBool(ref *types.ZendReference, val types.ZendBool) int {
	var tmp types.Zval
	types.ZVAL_BOOL(&tmp, val != 0)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefLong(ref *types.ZendReference, lval ZendLong) int {
	var tmp types.Zval
	tmp.SetLong(lval)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefDouble(ref *types.ZendReference, dval float64) int {
	var tmp types.Zval
	tmp.SetDouble(dval)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefEmptyString(ref *types.ZendReference) int {
	var tmp types.Zval
	ZVAL_EMPTY_STRING(&tmp)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefStr(ref *types.ZendReference, str *types.String) int {
	var tmp types.Zval
	tmp.SetString(str)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefString(ref *types.ZendReference, string *byte) int {
	var tmp types.Zval
	tmp.SetStringVal(b.CastStrAuto(string))
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefStringl(ref *types.ZendReference, string *byte, len_ int) int {
	var tmp types.Zval
	tmp.SetStringVal(b.CastStr(string, len_))
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefArr(ref *types.ZendReference, arr *types.Array) int {
	var tmp types.Zval
	tmp.SetArray(arr)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefRes(ref *types.ZendReference, res *types.ZendResource) int {
	var tmp types.Zval
	tmp.SetResource(res)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefZval(ref *types.ZendReference, zv *types.Zval) int {
	var tmp types.Zval
	types.ZVAL_COPY_VALUE(&tmp, zv)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefZvalEx(ref *types.ZendReference, zv *types.Zval, strict types.ZendBool) int {
	var tmp types.Zval
	types.ZVAL_COPY_VALUE(&tmp, zv)
	return ZendTryAssignTypedRefEx(ref, &tmp, strict)
}
func ZendDeclarePropertyEx(ce *types.ClassEntry, name *types.String, property *types.Zval, access_type int, doc_comment *types.String) int {
	return ZendDeclareTypedProperty(ce, name, property, access_type, doc_comment, 0)
}
func ZendDeclareProperty(ce *types.ClassEntry, name *byte, name_length int, property *types.Zval, access_type int) int {
	var key *types.String = types.NewString(b.CastStr(name, name_length))
	var ret int = ZendDeclarePropertyEx(ce, key, property, access_type, nil)
	types.ZendStringRelease(key)
	return ret
}
func ZendDeclarePropertyNull(ce *types.ClassEntry, name string, name_length int, access_type int) int {
	var property types.Zval
	property.SetNull()
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyBool(ce *types.ClassEntry, name *byte, name_length int, value ZendLong, access_type int) int {
	var property types.Zval
	types.ZVAL_BOOL(&property, value != 0)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyLong(ce *types.ClassEntry, name string, name_length int, value ZendLong, access_type int) int {
	var property types.Zval
	property.SetLong(value)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyDouble(ce *types.ClassEntry, name *byte, name_length int, value float64, access_type int) int {
	var property types.Zval
	property.SetDouble(value)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyString(ce *types.ClassEntry, name string, name_length int, value string, access_type int) int {
	var property types.Zval
	property.SetString(types.NewString(value))
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyStringl(
	ce *types.ClassEntry,
	name *byte,
	name_length int,
	value *byte,
	value_len int,
	access_type int,
) int {
	var property types.Zval
	property.SetString(types.NewString(b.CastStr(value, value_len)))
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclareClassConstantEx(ce *types.ClassEntry, name *types.String, value *types.Zval, access_type int, doc_comment *types.String) int {
	var c *ZendClassConstant
	if ce.IsInterface() {
		if access_type != ZEND_ACC_PUBLIC {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Access type for interface constant %s::%s must be public", ce.GetName().GetVal(), name.GetVal())
		}
	}
	if types.ZendStringEqualsLiteralCi(name, "class") {
		faults.ErrorNoreturn(b.Cond(ce.GetType() == ZEND_INTERNAL_CLASS, faults.E_CORE_ERROR, faults.E_COMPILE_ERROR), "A class constant must not be called 'class'; it is reserved for class name fetching")
	}
	if value.IsString() {
		ZvalMakeInternedString(value)
	}
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		c = Pemalloc(b.SizeOf("zend_class_constant"), 1)
	} else {
		c = ZendArenaAlloc(CG__().GetArena(), b.SizeOf("zend_class_constant"))
	}
	types.ZVAL_COPY_VALUE(c.GetValue(), value)
	c.GetValue().GetAccessFlags() = access_type
	c.SetDocComment(doc_comment)
	c.SetCe(ce)
	if value.IsConstant() {
		ce.SetIsConstantsUpdated(false)
	}
	if !(types.ZendHashAddPtr(ce.GetConstantsTable(), name.GetStr(), c)) {
		faults.ErrorNoreturn(b.Cond(ce.GetType() == ZEND_INTERNAL_CLASS, faults.E_CORE_ERROR, faults.E_COMPILE_ERROR), "Cannot redefine class __special__  constant %s::%s", ce.GetName().GetVal(), name.GetVal())
	}
	return types.SUCCESS
}
func ZendDeclareClassConstant(ce *types.ClassEntry, name *byte, name_length int, value *types.Zval) int {
	var ret int
	var key *types.String
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		key = types.ZendStringInitInterned(name, name_length, 1)
	} else {
		key = types.NewString(b.CastStr(name, name_length))
	}
	ret = ZendDeclareClassConstantEx(ce, key, value, ZEND_ACC_PUBLIC, nil)
	types.ZendStringRelease(key)
	return ret
}
func ZendDeclareClassConstantNull(ce *types.ClassEntry, name *byte, name_length int) int {
	var constant types.Zval
	constant.SetNull()
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantLong(ce *types.ClassEntry, name string, name_length int, value ZendLong) int {
	var constant types.Zval
	constant.SetLong(value)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantBool(ce *types.ClassEntry, name *byte, name_length int, value types.ZendBool) int {
	var constant types.Zval
	types.ZVAL_BOOL(&constant, value != 0)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantDouble(ce *types.ClassEntry, name *byte, name_length int, value float64) int {
	var constant types.Zval
	constant.SetDouble(value)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantStringl(ce *types.ClassEntry, name *byte, name_length int, value *byte, value_length int) int {
	var constant types.Zval
	constant.SetString(types.NewString(b.CastStr(value, value_length)))
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantString(ce *types.ClassEntry, name *byte, name_length int, value *byte) int {
	return ZendDeclareClassConstantStringl(ce, name, name_length, value, strlen(value))
}
func ZendUpdatePropertyEx(scope *types.ClassEntry, object *types.Zval, name *types.String, value *types.Zval) {
	var property types.Zval
	var old_scope *types.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property.SetString(name)
	types.Z_OBJ_HT_P(object).GetWriteProperty()(object, &property, value, nil)
	EG__().SetFakeScope(old_scope)
}
func ZendUpdateProperty(scope *types.ClassEntry, object *types.Zval, name *byte, name_length int, value *types.Zval) {
	var property types.Zval
	var old_scope *types.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property.SetStringVal(b.CastStr(name, name_length))
	types.Z_OBJ_HT_P(object).GetWriteProperty()(object, &property, value, nil)
	ZvalPtrDtor(&property)
	EG__().SetFakeScope(old_scope)
}
func ZendUpdatePropertyNull(scope *types.ClassEntry, object *types.Zval, name *byte, name_length int) {
	var tmp types.Zval
	tmp.SetNull()
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUnsetProperty(scope *types.ClassEntry, object *types.Zval, name string, name_length int) {
	var property types.Zval
	var old_scope *types.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property.SetStringVal(b.CastStr(name, name_length))
	types.Z_OBJ_HT_P(object).GetUnsetProperty()(object, &property, 0)
	ZvalPtrDtor(&property)
	EG__().SetFakeScope(old_scope)
}
func ZendUpdatePropertyBool(scope *types.ClassEntry, object *types.Zval, name *byte, name_length int, value ZendLong) {
	var tmp types.Zval
	types.ZVAL_BOOL(&tmp, value != 0)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyLong(scope *types.ClassEntry, object *types.Zval, name *byte, name_length int, value ZendLong) {
	var tmp types.Zval
	tmp.SetLong(value)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyDouble(scope *types.ClassEntry, object *types.Zval, name *byte, name_length int, value float64) {
	var tmp types.Zval
	tmp.SetDouble(value)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyStr(scope *types.ClassEntry, object *types.Zval, name *byte, name_length int, value *types.String) {
	var tmp types.Zval
	tmp.SetString(value)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyString(scope *types.ClassEntry, object *types.Zval, name *byte, name_length int, value *byte) {
	var tmp types.Zval
	tmp.SetStringVal(b.CastStrAuto(value))
	tmp.SetRefcount(0)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyStringl(
	scope *types.ClassEntry,
	object *types.Zval,
	name *byte,
	name_length int,
	value *byte,
	value_len int,
) {
	var tmp types.Zval
	tmp.SetStringVal(b.CastStr(value, value_len))
	tmp.SetRefcount(0)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyEx(scope *types.ClassEntry, name *types.String, value *types.Zval) int {
	var property *types.Zval
	var tmp types.Zval
	var prop_info *ZendPropertyInfo
	var old_scope *types.ClassEntry = EG__().GetFakeScope()
	if !scope.IsConstantsUpdated() {
		if ZendUpdateClassConstants(scope) != types.SUCCESS {
			return types.FAILURE
		}
	}
	EG__().SetFakeScope(scope)
	property = ZendStdGetStaticPropertyWithInfo(scope, name, BP_VAR_W, &prop_info)
	EG__().SetFakeScope(old_scope)
	if property == nil {
		return types.FAILURE
	}
	b.Assert(!(value.IsReference()))
	value.TryAddRefcount()
	if prop_info.GetType() != 0 {
		types.ZVAL_COPY_VALUE(&tmp, value)
		if ZendVerifyPropertyType(prop_info, &tmp, 0) == 0 {
			value.TryDelRefcount()
			return types.FAILURE
		}
		value = &tmp
	}
	ZendAssignToVariable(property, value, IS_TMP_VAR, 0)
	return types.SUCCESS
}
func ZendUpdateStaticProperty(scope *types.ClassEntry, name *byte, name_length int, value *types.Zval) int {
	var key *types.String = types.NewString(b.CastStr(name, name_length))
	var retval int = ZendUpdateStaticPropertyEx(scope, key, value)
	types.ZendStringEfree(key)
	return retval
}
func ZendUpdateStaticPropertyNull(scope *types.ClassEntry, name *byte, name_length int) int {
	var tmp types.Zval
	tmp.SetNull()
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyBool(scope *types.ClassEntry, name *byte, name_length int, value ZendLong) int {
	var tmp types.Zval
	types.ZVAL_BOOL(&tmp, value != 0)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyLong(scope *types.ClassEntry, name *byte, name_length int, value ZendLong) int {
	var tmp types.Zval
	tmp.SetLong(value)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyDouble(scope *types.ClassEntry, name *byte, name_length int, value float64) int {
	var tmp types.Zval
	tmp.SetDouble(value)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyString(scope *types.ClassEntry, name *byte, name_length int, value *byte) int {
	var tmp types.Zval
	tmp.SetStringVal(b.CastStrAuto(value))
	tmp.SetRefcount(0)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyStringl(scope *types.ClassEntry, name *byte, name_length int, value *byte, value_len int) int {
	var tmp types.Zval
	tmp.SetStringVal(b.CastStr(value, value_len))
	tmp.SetRefcount(0)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendReadPropertyEx(scope *types.ClassEntry, object *types.Zval, name *types.String, silent types.ZendBool, rv *types.Zval) *types.Zval {
	var property types.Zval
	var value *types.Zval
	var old_scope *types.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property.SetString(name)
	value = types.Z_OBJ_HT_P(object).GetReadProperty()(object, &property, b.Cond(silent != 0, BP_VAR_IS, BP_VAR_R), nil, rv)
	EG__().SetFakeScope(old_scope)
	return value
}
func ZendReadProperty(
	scope *types.ClassEntry,
	object *types.Zval,
	name string,
	name_length int,
	silent types.ZendBool,
	rv *types.Zval,
) *types.Zval {
	var value *types.Zval
	var str *types.String
	str = types.NewString(b.CastStr(name, name_length))
	value = ZendReadPropertyEx(scope, object, str, silent, rv)
	types.ZendStringReleaseEx(str, 0)
	return value
}
func ZendReadStaticPropertyEx(scope *types.ClassEntry, name *types.String, silent types.ZendBool) *types.Zval {
	var property *types.Zval
	var old_scope *types.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property = ZendStdGetStaticProperty(scope, name, b.Cond(silent != 0, BP_VAR_IS, BP_VAR_R))
	EG__().SetFakeScope(old_scope)
	return property
}
func ZendReadStaticProperty(scope *types.ClassEntry, name *byte, name_length int, silent types.ZendBool) *types.Zval {
	var key *types.String = types.NewString(b.CastStr(name, name_length))
	var property *types.Zval = ZendReadStaticPropertyEx(scope, key, silent)
	types.ZendStringEfree(key)
	return property
}
func ZendSaveErrorHandling(current *ZendErrorHandling) {
	current.SetHandling(EG__().GetErrorHandling())
	current.SetException(EG__().GetExceptionClass())
	current.GetUserHandler().SetUndef()
}
func ZendReplaceErrorHandling(error_handling ZendErrorHandlingT, exception_class *types.ClassEntry, current *ZendErrorHandling) {
	if current != nil {
		ZendSaveErrorHandling(current)
	}
	b.Assert(error_handling == EH_THROW || exception_class == nil)
	EG__().SetErrorHandling(error_handling)
	EG__().SetExceptionClass(exception_class)
}
func ZendRestoreErrorHandling(saved *ZendErrorHandling) {
	EG__().SetErrorHandling(saved.GetHandling())
	EG__().SetExceptionClass(saved.GetException())
}
func ZendFindAliasName(ce *types.ClassEntry, name *types.String) *types.String {
	var alias *ZendTraitAlias
	var alias_ptr **ZendTraitAlias
	if b.Assign(&alias_ptr, ce.GetTraitAliases()) {
		alias = *alias_ptr
		for alias != nil {
			if alias.GetAlias() != nil && types.ZendStringEqualsCi(alias.GetAlias(), name) {
				return alias.GetAlias()
			}
			alias_ptr++
			alias = *alias_ptr
		}
	}
	return name
}
func ZendResolveMethodName(ce *types.ClassEntry, f *ZendFunction) *types.String {
	var func_ *ZendFunction
	var function_table *types.Array
	var name *types.String
	if f.GetCommonType() != ZEND_USER_FUNCTION || f.GetOpArray().GetRefcount() != nil && (*(f.GetOpArray().GetRefcount())) < 2 || f.GetScope() == nil || f.GetScope().GetTraitAliases() == nil {
		return f.GetFunctionName()
	}
	function_table = ce.GetFunctionTable()
	var __ht *types.Array = function_table
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

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
func ZendGetObjectType(ce *types.ClassEntry) *byte {
	if ce.IsTrait() {
		return "trait"
	} else if ce.IsInterface() {
		return "interface"
	} else {
		return "class"
	}
}
func ZendIsIterable(iterable *types.Zval) types.ZendBool {
	switch iterable.GetType() {
	case types.IS_ARRAY:
		return 1
	case types.IS_OBJECT:
		return InstanceofFunction(types.Z_OBJCE_P(iterable), ZendCeTraversable)
	default:
		return 0
	}
}
func ZendIsCountable(countable *types.Zval) types.ZendBool {
	switch countable.GetType() {
	case types.IS_ARRAY:
		return 1
	case types.IS_OBJECT:
		if types.Z_OBJ_HT_P(countable).GetCountElements() != nil {
			return 1
		}
		return InstanceofFunction(types.Z_OBJCE_P(countable), ZendCeCountable)
	default:
		return 0
	}
}
