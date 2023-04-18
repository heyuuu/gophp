package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZendTryAssignTypedRefNull(ref *types2.ZendReference) int {
	var tmp types2.Zval
	tmp.SetNull()
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefLong(ref *types2.ZendReference, lval ZendLong) int {
	var tmp types2.Zval
	tmp.SetLong(lval)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefDouble(ref *types2.ZendReference, dval float64) int {
	var tmp types2.Zval
	tmp.SetDouble(dval)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefEmptyString(ref *types2.ZendReference) int {
	zv := types2.NewZvalString("")
	return ZendTryAssignTypedRef(ref, zv)
}
func ZendTryAssignTypedRefStr(ref *types2.ZendReference, str *types2.String) int {
	var tmp types2.Zval
	tmp.SetString(str)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefString(ref *types2.ZendReference, string *byte) int {
	var tmp types2.Zval
	tmp.SetStringVal(b.CastStrAuto(string))
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefStringl(ref *types2.ZendReference, string *byte, len_ int) int {
	var tmp types2.Zval
	tmp.SetStringVal(b.CastStr(string, len_))
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefArr(ref *types2.ZendReference, arr *types2.Array) int {
	var tmp types2.Zval
	tmp.SetArray(arr)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefZvalEx(ref *types2.ZendReference, zv *types2.Zval, strict types2.ZendBool) int {
	var tmp types2.Zval
	types2.ZVAL_COPY_VALUE(&tmp, zv)
	return ZendTryAssignTypedRefEx(ref, &tmp, strict)
}
func ZendDeclarePropertyEx(ce *types2.ClassEntry, name *types2.String, property *types2.Zval, access_type int, doc_comment *types2.String) int {
	return ZendDeclareTypedProperty(ce, name, property, access_type, doc_comment, 0)
}
func ZendDeclareProperty(ce *types2.ClassEntry, name *byte, name_length int, property *types2.Zval, access_type int) int {
	var key *types2.String = types2.NewString(b.CastStr(name, name_length))
	var ret int = ZendDeclarePropertyEx(ce, key, property, access_type, nil)
	// types.ZendStringRelease(key)
	return ret
}
func ZendDeclarePropertyNull(ce *types2.ClassEntry, name string, name_length int, access_type int) int {
	var property types2.Zval
	property.SetNull()
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyLong(ce *types2.ClassEntry, name string, name_length int, value ZendLong, access_type int) int {
	var property types2.Zval
	property.SetLong(value)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyString(ce *types2.ClassEntry, name string, name_length int, value string, access_type int) int {
	var property types2.Zval
	property.SetString(types2.NewString(value))
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclareClassConstantEx(ce *types2.ClassEntry, name *types2.String, value *types2.Zval, access_type int, doc_comment *types2.String) int {
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
	return types2.SUCCESS
}
func ZendDeclareClassConstant(ce *types2.ClassEntry, name string, value *types2.Zval) int {
	key := types2.NewString(name)
	ret := ZendDeclareClassConstantEx(ce, key, value, AccPublic, nil)
	return ret
}
func ZendDeclareClassConstantLong(ce *types2.ClassEntry, name string, value ZendLong) int {
	var constant types2.Zval
	constant.SetLong(value)
	return ZendDeclareClassConstant(ce, name, &constant)
}
func ZendUpdatePropertyEx(scope *types2.ClassEntry, object *types2.Zval, name string, value *types2.Zval) {
	var oldScope *types2.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property := types2.NewZvalString(name)
	types2.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, nil)
	EG__().SetFakeScope(oldScope)
}
func ZendUnsetProperty(scope *types2.ClassEntry, object *types2.Zval, name string) {
	var oldScope *types2.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property := types2.NewZvalString(name)
	types2.Z_OBJ_HT_P(object).GetUnsetProperty()(object, property, 0)
	EG__().SetFakeScope(oldScope)
}
func ZendReadPropertyEx(scope *types2.ClassEntry, object *types2.Zval, name *types2.String, silent types2.ZendBool, rv *types2.Zval) *types2.Zval {
	return ZendReadProperty(scope, object, name.GetStr(), silent, rv)
}
func ZendReadProperty(scope *types2.ClassEntry, object *types2.Zval, name string, silent types2.ZendBool, rv *types2.Zval) *types2.Zval {
	var oldScope *types2.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property := types2.NewZvalString(name)
	value := types2.Z_OBJ_HT_P(object).GetReadProperty()(object, property, b.Cond(silent != 0, BP_VAR_IS, BP_VAR_R), nil, rv)
	EG__().SetFakeScope(oldScope)
	return value
}
func ZendSaveErrorHandling(current *ZendErrorHandling) {
	current.SetHandling(EG__().GetErrorHandling())
	current.SetException(EG__().GetExceptionClass())
	current.GetUserHandler().SetUndef()
}
func ZendReplaceErrorHandling(error_handling ZendErrorHandlingT, exception_class *types2.ClassEntry, current *ZendErrorHandling) {
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
func ZendFindAliasName(ce *types2.ClassEntry, name string) *types2.String {
	for _, alias := range ce.GetTraitAliases() {
		if alias.GetAlias() != nil && ascii.StrCaseEquals(alias.GetAlias().GetStr(), name) {
			return alias.GetAlias()
		}
	}
	return types2.NewString(name)
}
func ZendResolveMethodName(ce *types2.ClassEntry, f types2.IFunction) *types2.String {
	if f.GetType() != ZEND_USER_FUNCTION || f.GetOpArray().GetRefcount() != nil && (*(f.GetOpArray().GetRefcount())) < 2 || f.GetScope() == nil || f.GetScope().GetTraitAliases() == nil {
		return f.GetFunctionName()
	}

	var ret = f.GetFunctionName()
	ce.FunctionTable().ForeachEx(func(name string, func_ types2.IFunction) bool {
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
func ZendGetObjectType(ce *types2.ClassEntry) *byte {
	if ce.IsTrait() {
		return "trait"
	} else if ce.IsInterface() {
		return "interface"
	} else {
		return "class"
	}
}
func ZendIsIterable(iterable *types2.Zval) types2.ZendBool {
	switch iterable.GetType() {
	case types2.IS_ARRAY:
		return 1
	case types2.IS_OBJECT:
		return InstanceofFunction(types2.Z_OBJCE_P(iterable), ZendCeTraversable)
	default:
		return 0
	}
}
func ZendIsCountable(countable *types2.Zval) types2.ZendBool {
	switch countable.GetType() {
	case types2.IS_ARRAY:
		return 1
	case types2.IS_OBJECT:
		if types2.Z_OBJ_HT_P(countable).GetCountElements() != nil {
			return 1
		}
		return InstanceofFunction(types2.Z_OBJCE_P(countable), ZendCeCountable)
	default:
		return 0
	}
}
