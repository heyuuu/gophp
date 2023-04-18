package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZendTryAssignTypedRefNull(ref *types.ZendReference) int {
	var tmp types.Zval
	tmp.SetNull()
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
	zv := types.NewZvalString("")
	return ZendTryAssignTypedRef(ref, zv)
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
	// types.ZendStringRelease(key)
	return ret
}
func ZendDeclarePropertyNull(ce *types.ClassEntry, name string, name_length int, access_type int) int {
	var property types.Zval
	property.SetNull()
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyLong(ce *types.ClassEntry, name string, name_length int, value ZendLong, access_type int) int {
	var property types.Zval
	property.SetLong(value)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyString(ce *types.ClassEntry, name string, name_length int, value string, access_type int) int {
	var property types.Zval
	property.SetString(types.NewString(value))
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclareClassConstantEx(ce *types.ClassEntry, name *types.String, value *types.Zval, access_type int, doc_comment *types.String) int {
	if ce.IsInterface() {
		if access_type != AccPublic {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Access type for interface constant %s::%s must be public", ce.GetName().GetVal(), name.GetVal())
		}
	}
	if ascii.StrCaseEquals(name.GetStr(), "class") {
		faults.ErrorNoreturn(b.Cond(ce.GetType() == ZEND_INTERNAL_CLASS, faults.E_CORE_ERROR, faults.E_COMPILE_ERROR), "A class constant must not be called 'class'; it is reserved for class name fetching")
	}

	var c *ZendClassConstant = NewClassConstant(ce, value, doc_comment)
	c.GetValue().SetAccessFlags(uint32(access_type))
	if value.IsConstantAst() {
		ce.SetIsConstantsUpdated(false)
	}
	if !ce.ConstantsTable().Add(name.GetStr(), c) {
		faults.ErrorNoreturn(b.Cond(ce.GetType() == ZEND_INTERNAL_CLASS, faults.E_CORE_ERROR, faults.E_COMPILE_ERROR), "Cannot redefine class __special__  constant %s::%s", ce.GetName().GetVal(), name.GetVal())
	}
	return types.SUCCESS
}
func ZendDeclareClassConstant(ce *types.ClassEntry, name string, value *types.Zval) int {
	key := types.NewString(name)
	ret := ZendDeclareClassConstantEx(ce, key, value, AccPublic, nil)
	return ret
}
func ZendDeclareClassConstantLong(ce *types.ClassEntry, name string, value ZendLong) int {
	var constant types.Zval
	constant.SetLong(value)
	return ZendDeclareClassConstant(ce, name, &constant)
}
func ZendUpdatePropertyEx(scope *types.ClassEntry, object *types.Zval, name string, value *types.Zval) {
	var oldScope *types.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property := types.NewZvalString(name)
	types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, nil)
	EG__().SetFakeScope(oldScope)
}
func ZendUnsetProperty(scope *types.ClassEntry, object *types.Zval, name string) {
	var oldScope *types.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property := types.NewZvalString(name)
	types.Z_OBJ_HT_P(object).GetUnsetProperty()(object, property, 0)
	EG__().SetFakeScope(oldScope)
}
func ZendReadPropertyEx(scope *types.ClassEntry, object *types.Zval, name *types.String, silent types.ZendBool, rv *types.Zval) *types.Zval {
	return ZendReadProperty(scope, object, name.GetStr(), silent, rv)
}
func ZendReadProperty(scope *types.ClassEntry, object *types.Zval, name string, silent types.ZendBool, rv *types.Zval) *types.Zval {
	var oldScope *types.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property := types.NewZvalString(name)
	value := types.Z_OBJ_HT_P(object).GetReadProperty()(object, property, b.Cond(silent != 0, BP_VAR_IS, BP_VAR_R), nil, rv)
	EG__().SetFakeScope(oldScope)
	return value
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
func ZendFindAliasName(ce *types.ClassEntry, name string) *types.String {
	for _, alias := range ce.GetTraitAliases() {
		if alias.GetAlias() != nil && ascii.StrCaseEquals(alias.GetAlias().GetStr(), name) {
			return alias.GetAlias()
		}
	}
	return types.NewString(name)
}
func ZendResolveMethodName(ce *types.ClassEntry, f types.IFunction) *types.String {
	if f.GetType() != ZEND_USER_FUNCTION || f.GetOpArray().GetRefcount() != nil && (*(f.GetOpArray().GetRefcount())) < 2 || f.GetScope() == nil || f.GetScope().GetTraitAliases() == nil {
		return f.GetFunctionName()
	}

	var ret = f.GetFunctionName()
	ce.FunctionTable().ForeachEx(func(name string, func_ types.IFunction) bool {
		if func_ != f {
			return true
		}

		if name != "" && !ascii.StrCaseEquals(name, f.GetFunctionName().GetStr()) {
			ret = ZendFindAliasName(f.GetScope(), name)
		}
		return false
	})

	return ret
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
