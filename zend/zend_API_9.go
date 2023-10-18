package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZendTryAssignTypedRefNull(ref *types.Reference) int {
	var tmp types.Zval
	tmp.SetNull()
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefLong(ref *types.Reference, lval ZendLong) int {
	var tmp types.Zval
	tmp.SetLong(lval)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefDouble(ref *types.Reference, dval float64) int {
	var tmp types.Zval
	tmp.SetDouble(dval)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefEmptyString(ref *types.Reference) int {
	zv := types.NewZvalString("")
	return ZendTryAssignTypedRef(ref, zv)
}
func ZendTryAssignTypedRefStr(ref *types.Reference, str *types.String) int {
	var tmp types.Zval
	tmp.SetStringEx(str)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefString(ref *types.Reference, string *byte) int {
	var tmp types.Zval
	tmp.SetString(b.CastStrAuto(string))
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefStringl(ref *types.Reference, string *byte, len_ int) int {
	var tmp types.Zval
	tmp.SetString(b.CastStr(string, len_))
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefArr(ref *types.Reference, arr *types.Array) int {
	return ZendTryAssignTypedRef(ref, types.NewZvalArray(arr))
}
func ZendTryAssignTypedRefZvalEx(ref *types.Reference, zv *types.Zval, strict bool) int {
	return ZendTryAssignTypedRefEx(ref, zv.CopyValue(), strict)
}
func ZendDeclareProperty(ce *types.ClassEntry, name string, property *types.Zval, accessType uint32) int {
	return ZendDeclareTypedProperty(ce, name, property, accessType, "", nil)
}
func ZendDeclarePropertyNull(ce *types.ClassEntry, name string, accessType uint32) int {
	return ZendDeclareProperty(ce, name, types.NewZvalNull(), accessType)
}
func ZendDeclarePropertyLong(ce *types.ClassEntry, name string, value int, accessType uint32) int {
	return ZendDeclareProperty(ce, name, types.NewZvalLong(value), accessType)
}
func ZendDeclarePropertyString(ce *types.ClassEntry, name string, value string, accessType uint32) int {
	return ZendDeclareProperty(ce, name, types.NewZvalString(value), accessType)
}
func ZendDeclareClassConstantEx(ce *types.ClassEntry, name *types.String, value *types.Zval, accessType uint32, docComment string) int {
	if ce.IsInterface() {
		if accessType != types.AccPublic {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, fmt.Sprintf("Access type for interface constant %s::%s must be public", ce.Name(), name.GetStr()))
		}
	}
	if ascii.StrCaseEquals(name.GetStr(), "class") {
		faults.ErrorNoreturn(lang.Cond(ce.IsInternalClass(), faults.E_CORE_ERROR, faults.E_COMPILE_ERROR), "A class constant must not be called 'class'; it is reserved for class name fetching")
	}

	var c = types.NewClassConstant(ce, value, docComment, accessType)
	if value.IsConstantAst() {
		ce.SetIsConstantsUpdated(false)
	}
	if !ce.ConstantsTable().Add(name.GetStr(), c) {
		faults.ErrorNoreturn(lang.Cond(ce.IsInternalClass(), faults.E_CORE_ERROR, faults.E_COMPILE_ERROR), fmt.Sprintf("Cannot redefine class __special__  constant %s::%s", ce.Name(), name.GetStr()))
	}
	return types.SUCCESS
}
func ZendDeclareClassConstant(ce *types.ClassEntry, name string, value *types.Zval) int {
	key := types.NewString(name)
	ret := ZendDeclareClassConstantEx(ce, key, value, types.AccPublic, "")
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
	object.Object().WritePropertyEx(property, value)
	EG__().SetFakeScope(oldScope)
}
func ZendUnsetProperty(scope *types.ClassEntry, object *types.Zval, name string) {
	var oldScope *types.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	property := types.NewZvalString(name)
	object.Object().UnsetPropertyEx(property)
	EG__().SetFakeScope(oldScope)
}
func ZendReadProperty(scope *types.ClassEntry, object *types.Zval, name string, silent bool, rv *types.Zval) *types.Zval {
	var oldScope *types.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(scope)
	value := object.Object().ReadPropertyEx(types.NewZvalString(name), lang.Cond(silent, BP_VAR_IS, BP_VAR_R), rv)
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
func ZendFindAliasName(ce *types.ClassEntry, name string) string {
	for _, alias := range ce.GetTraitAliases() {
		if alias.GetAlias() != "" && ascii.StrCaseEquals(alias.GetAlias(), name) {
			return alias.GetAlias()
		}
	}
	return name
}

func ZendResolveMethodName(ce *types.ClassEntry, f types.IFunction) string {
	if f.GetType() != ZEND_USER_FUNCTION || f.GetOpArray().GetRefcount() != nil && (*(f.GetOpArray().GetRefcount())) < 2 || f.GetScope() == nil || f.GetScope().GetTraitAliases() == nil {
		return f.FunctionName()
	}

	var ret = f.FunctionName()
	ce.FunctionTable().ForeachEx(func(name string, func_ types.IFunction) bool {
		if func_ != f {
			return true
		}

		if name != "" && !ascii.StrCaseEquals(name, f.FunctionName()) {
			ret = ZendFindAliasName(f.GetScope(), name)
		}
		return false
	})

	return ret
}
func ZendGetObjectType(ce *types.ClassEntry) string {
	if ce.IsTrait() {
		return "trait"
	} else if ce.IsInterface() {
		return "interface"
	} else {
		return "class"
	}
}
func ZendIsIterable(iterable *types.Zval) bool {
	switch iterable.Type() {
	case types.IsArray:
		return true
	case types.IsObject:
		return operators.InstanceofFunction(types.Z_OBJCE_P(iterable), ZendCeTraversable)
	default:
		return false
	}
}
func ZendIsCountable(countable *types.Zval) bool {
	switch countable.Type() {
	case types.IsArray:
		return true
	case types.IsObject:
		if countable.Object().CanCountElements() {
			return true
		}
		return operators.InstanceofFunction(types.Z_OBJCE_P(countable), ZendCeCountable)
	default:
		return false
	}
}
