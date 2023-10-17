package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ZEND_CLOSURE_OBJECT(op_array types.IFunction) *types.Object {
	return (*types.Object)((*byte)(op_array - b.SizeOf("zend_object")))
}
func ZEND_CLOSURE_PROPERTY_ERROR() {
	faults.ThrowError(nil, "Closure object cannot have properties")
}
func zim_Closure___invoke(executeData *ZendExecuteData, return_value *types.Zval) {
	var func_ types.IFunction = executeData.GetFunc()
	var arguments *types.Zval = executeData.Arg(1)
	if CallUserFunction(nil, executeData.ThisObjectZval(), return_value, executeData.NumArgs(), arguments) == types.FAILURE {
		return_value.SetFalse()
	}

	/* destruct the function also, then - we have allocated it in get_method */

	// types.ZendStringReleaseEx(func_.GetInternalFunction().GetFunctionName(), 0)
	Efree(func_)
}
func ZendValidClosureBinding(closure *ZendClosure, newthis *types.Zval, scope *types.ClassEntry) bool {
	var func_ types.IFunction = closure.GetFunc()
	var is_fake_closure bool = func_.IsFakeClosure()
	if newthis != nil {
		if func_.IsStatic() {
			faults.Error(faults.E_WARNING, "Cannot bind an instance to a static closure")
			return 0
		}
		if is_fake_closure != 0 && func_.GetScope() != nil && operators.InstanceofFunction(types.Z_OBJCE_P(newthis), func_.GetScope()) == 0 {

			/* Binding incompatible $this to an internal method is not supported. */

			faults.Error(faults.E_WARNING, fmt.Sprintf("Cannot bind method %s::%s() to object of class %s", func_.GetScope().Name(), func_.FunctionName(), types.Z_OBJCE_P(newthis).Name()))
			return 0
		}
	} else if is_fake_closure != 0 && func_.GetScope() != nil && !func_.IsStatic() {
		if func_.GetType() == ZEND_INTERNAL_FUNCTION {
			faults.Error(faults.E_WARNING, "Cannot unbind $this of internal method")
			return 0
		} else {
			faults.Error(faults.E_DEPRECATED, "Unbinding $this of a method is deprecated")
		}
	} else if is_fake_closure == 0 && !(closure.GetThisPtr().IsUndef()) && func_.IsUsesThis() {

		// TODO: Only deprecate if it had $this *originally*?

		faults.Error(faults.E_DEPRECATED, "Unbinding $this of closure is deprecated")

		// TODO: Only deprecate if it had $this *originally*?

	}
	if scope != nil && scope != func_.GetScope() && scope.IsInternalClass() {

		/* rebinding to internal class is not allowed */

		faults.Error(faults.E_WARNING, fmt.Sprintf("Cannot bind closure to scope of internal class %s", scope.Name()))
		return 0
	}
	if is_fake_closure != 0 && scope != func_.GetScope() {
		if func_.GetScope() == nil {
			faults.Error(faults.E_WARNING, "Cannot rebind scope of closure created from function")
		} else {
			faults.Error(faults.E_WARNING, "Cannot rebind scope of closure created from method")
		}
		return 0
	}
	return 1
}

//zif -old "o*"
func zim_Closure_call(executeData *ZendExecuteData, returnValue *types.Zval) {
	// zpp parse start
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	newthis := fp.ParseObject()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	// zpp parse finish

	var closure_result types.Zval
	var closure *ZendClosure
	var fci_cache types.ZendFcallInfoCache
	var my_function types.IFunction
	var newobj *types.Object
	closure = (*ZendClosure)(executeData.ThisObject())
	newobj = newthis.Object()
	if ZendValidClosureBinding(closure, newthis, types.Z_OBJCE_P(newthis)) == 0 {
		return
	}
	if closure.GetFunc().IsGenerator() {
		var new_closure types.Zval
		ZendCreateClosure(&new_closure, closure.GetFunc(), types.Z_OBJCE_P(newthis), closure.GetCalledScope(), newthis)
		closure = (*ZendClosure)(new_closure.Object())
		fci_cache.SetFunctionHandler(closure.GetFunc())
	} else {
		memcpy(&my_function, closure.GetFunc(), lang.CondF(closure.GetFunc().GetType() == ZEND_USER_FUNCTION, func() __auto__ { return b.SizeOf("zend_op_array") }, func() __auto__ { return b.SizeOf("zend_internal_function") }))
		my_function.SetIsClosure(false)

		/* use scope of passed object */

		my_function.SetScope(types.Z_OBJCE_P(newthis))
		if closure.GetFunc().GetType() == ZEND_INTERNAL_FUNCTION {
			my_function.GetInternalFunction().SetHandler(closure.GetOrigInternalHandler())
		}
		fci_cache.SetFunctionHandler(&my_function)

		/* Runtime cache relies on bound scope to be immutable, hence we need a separate rt cache in case scope changed */

		if ZEND_USER_CODE(my_function.GetType()) && (closure.GetFunc().GetScope() != types.Z_OBJCE_P(newthis) || closure.GetFunc().IsHeapRtCache()) {
			my_function.GetOpArray().SetIsHeapRtCache(true)
			my_function.GetOpArray().InitRunTimeCache()
		}

		/* Runtime cache relies on bound scope to be immutable, hence we need a separate rt cache in case scope changed */

	}
	var fci *types.ZendFcallInfo = types.InitFCallInfo(newobj, &closure_result, args...)
	fci.GetFunctionName().SetObject(closure.GetStd())

	fci_cache.SetCalledScope(newobj.GetCe())
	fci_cache.SetObject(fci.GetObject())
	if ZendCallFunction(fci, &fci_cache) == types.SUCCESS && closure_result.IsNotUndef() {
		if closure_result.IsRef() {
			operators.ZendUnwrapReference(&closure_result)
		}
		types.ZVAL_COPY_VALUE(returnValue, &closure_result)
	}
}
func zim_Closure_bind(executeData *ZendExecuteData, return_value *types.Zval) {
	var newthis *types.Zval
	var zclosure *types.Zval
	var scope_arg *types.Zval = nil
	var closure *ZendClosure
	var ce *types.ClassEntry
	var called_scope *types.ClassEntry
	if ZendParseMethodParameters(executeData.NumArgs(), executeData.ThisObjectZval(), "Oo!|z", &zclosure, ZendCeClosure, &newthis, &scope_arg) == types.FAILURE {
		return
	}
	closure = (*ZendClosure)(zclosure.Object())
	if scope_arg != nil {
		if scope_arg.IsObject() {
			ce = types.Z_OBJCE_P(scope_arg)
		} else if scope_arg.IsNull() {
			ce = nil
		} else {
			var class_name *types.String = operators.ZvalGetString(scope_arg)
			if class_name.GetStr() == "static" {
				ce = closure.GetFunc().GetScope()
			} else if lang.Assign(&ce, ZendLookupClass(class_name.GetStr())) == nil {
				faults.Error(faults.E_WARNING, fmt.Sprintf("Class '%s' not found", class_name.GetVal()))
				return_value.SetNull()
				return
			}
		}
	} else {
		ce = closure.GetFunc().GetScope()
	}
	if ZendValidClosureBinding(closure, newthis, ce) == 0 {
		return
	}
	if newthis != nil {
		called_scope = types.Z_OBJCE_P(newthis)
	} else {
		called_scope = ce
	}
	ZendCreateClosure(return_value, closure.GetFunc(), ce, called_scope, newthis)
}
func ZendClosureCallMagic(executeData *ZendExecuteData, returnValue *types.Zval) {
	// init fci
	obj := executeData.ThisObject()
	param1 := types.NewZvalString(executeData.GetFunc().FunctionName())
	var param2 *types.Zval
	if executeData.NumArgs() != 0 {
		param2 = types.NewZvalArray(types.NewArrayCap(executeData.NumArgs()))
		ZendCopyParametersArray(executeData.NumArgs(), param2)
	} else {
		param2 = types.NewZvalArray(nil)
	}
	fci := types.InitFCallInfo(obj, returnValue, param1, param2)

	// init fcc
	var fcc types.ZendFcallInfoCache
	fcc.SetObject(fci.GetObject())
	fcc.SetCalledScope(ZendGetCalledScope(CurrEX()))
	if (executeData.GetFunc().GetInternalFunction().GetFnFlags() & types.AccStatic) != 0 {
		fcc.SetFunctionHandler(executeData.GetFunc().GetInternalFunction().GetScope().GetCallstatic())
	} else {
		fcc.SetFunctionHandler(executeData.GetFunc().GetInternalFunction().GetScope().GetCall())
	}

	ZendCallFunction(fci, &fcc)
}
func ZendCreateClosureFromCallable(return_value *types.Zval, callable *types.Zval, error **byte) int {
	var fcc types.ZendFcallInfoCache
	var mptr types.IFunction
	var instance types.Zval
	if ZendIsCallableEx(callable, nil, 0, nil, &fcc, error) == 0 {
		return types.FAILURE
	}
	mptr = fcc.GetFunctionHandler()
	if mptr.IsCallViaTrampoline() {
		/* For Closure::fromCallable([$closure, "__invoke"]) return $closure. */
		if fcc.GetObject() != nil && fcc.GetObject().GetCe() == ZendCeClosure && mptr.FunctionName() == "__invoke" {
			return_value.SetObject(fcc.GetObject())
			//fcc.GetObject().AddRefcount()
			ZendFreeTrampoline(mptr)
			return types.SUCCESS
		}
		if mptr.GetScope() == nil {
			return types.FAILURE
		}
		if mptr.IsStatic() {
			if mptr.GetScope().GetCallstatic() == nil {
				return types.FAILURE
			}
		} else {
			if mptr.GetScope().GetCall() == nil {
				return types.FAILURE
			}
		}
		call := types.NewInternalFunctionEx(mptr.FunctionName(), ZendClosureCallMagic)
		call.SetFnFlags(mptr.GetFnFlags() & types.AccStatic)
		call.SetScope(mptr.GetScope())
		ZendFreeTrampoline(mptr)
		mptr = call
	}
	if fcc.GetObject() != nil {
		instance.SetObject(fcc.GetObject())
		ZendCreateFakeClosure(return_value, mptr, mptr.GetScope(), fcc.GetCalledScope(), &instance)
	} else {
		ZendCreateFakeClosure(return_value, mptr, mptr.GetScope(), fcc.GetCalledScope(), nil)
	}
	return types.SUCCESS
}
func zim_Closure_fromCallable(executeData *ZendExecuteData, return_value *types.Zval) {
	var callable *types.Zval
	var success int
	var error *byte = nil
	if ZendParseParameters(executeData.NumArgs(), "z", &callable) == types.FAILURE {
		return
	}
	if callable.IsObject() && operators.InstanceofFunction(types.Z_OBJCE_P(callable), ZendCeClosure) != 0 {

		/* It's already a closure */

		ZVAL_ZVAL(return_value, callable, 1, 0)
		return
	}

	/* create closure as if it were called from parent scope */

	EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
	success = ZendCreateClosureFromCallable(return_value, callable, &error)
	EG__().SetCurrentExecuteData(executeData)
	if success == types.FAILURE || error != nil {
		if error != nil {
			faults.TypeError(fmt.Sprintf("Failed to create closure from callable: %s", error))
			Efree(error)
		} else {
			faults.TypeError("Failed to create closure from callable")
		}
	}
}
func ZendClosureGetConstructor(object *types.Object) types.IFunction {
	faults.ThrowError(nil, "Instantiation of 'Closure' is not allowed")
	return nil
}
func ZendClosureCompareObjects(o1 *types.Zval, o2 *types.Zval) int {
	return o1.Object() != o2.Object()
}
func ZendGetClosureInvokeMethod(object *types.Object) types.IFunction {
	var closure *ZendClosure = (*ZendClosure)(object)
	var invoke types.IFunction = types.NewInternalFunction()

	(types.IFunction)(Emalloc(b.SizeOf("zend_function")))
	var keep_flags uint32 = types.AccReturnReference | types.AccVariadic | types.AccHasReturnType
	invoke.SetCommon(closure.GetFunc().GetCommon())

	/* We return ZEND_INTERNAL_FUNCTION, but arg_info representation is the
	 * same as for ZEND_USER_FUNCTION (uses zend_string* instead of char*).
	 * This is not a problem, because ZEND_ACC_HAS_TYPE_HINTS is never set,
	 * and we won't check arguments on internal function. We also set
	 * ZEND_ACC_USER_ARG_INFO flag to prevent invalid usage by Reflection */

	invoke.SetType(ZEND_INTERNAL_FUNCTION)
	invoke.GetInternalFunction().SetFnFlags(types.AccPublic | types.AccCallViaHandler | closure.GetFunc().GetFnFlags()&keep_flags)
	if closure.GetFunc().GetType() != ZEND_INTERNAL_FUNCTION || closure.GetFunc().IsUserArgInfo() {
		invoke.GetInternalFunction().SetIsUserArgInfo(true)
	}
	invoke.GetInternalFunction().SetHandler(zim_Closure___invoke)
	invoke.GetInternalFunction().SetModule(nil)
	invoke.GetInternalFunction().SetScope(ZendCeClosure)
	invoke.GetInternalFunction().SetFunctionName(types.STR_MAGIC_INVOKE)
	return invoke
}
func ZendClosureGetMethod(object **types.Object, method *types.String, key *types.Zval) types.IFunction {
	if ascii.StrCaseEquals(method.GetStr(), ZEND_INVOKE_FUNC_NAME) {
		return ZendGetClosureInvokeMethod(*object)
	}
	return ZendStdGetMethod(object, method, key)
}
func ZendClosureReadProperty(object *types.Zval, member *types.Zval, type_ int, cache_slot *any, rv *types.Zval) *types.Zval {
	ZEND_CLOSURE_PROPERTY_ERROR()
	return UninitializedZval()
}
func ZendClosureWriteProperty(object *types.Zval, member *types.Zval, value *types.Zval, cache_slot *any) *types.Zval {
	ZEND_CLOSURE_PROPERTY_ERROR()
	return EG__().GetErrorZval()
}
func ZendClosureGetPropertyPtrPtr(object *types.Zval, member *types.Zval, type_ int, cache_slot *any) *types.Zval {
	ZEND_CLOSURE_PROPERTY_ERROR()
	return nil
}
func ZendClosureHasProperty(object *types.Zval, member *types.Zval, has_set_exists int, cache_slot *any) int {
	if has_set_exists != ZEND_PROPERTY_EXISTS {
		ZEND_CLOSURE_PROPERTY_ERROR()
	}
	return 0
}
func ZendClosureUnsetProperty(object *types.Zval, member *types.Zval, cache_slot *any) {
	ZEND_CLOSURE_PROPERTY_ERROR()
}
func ZendClosureFreeStorage(object *types.Object) {
	var closure *ZendClosure = (*ZendClosure)(object)
	ZendObjectStdDtor(closure.GetStd())
	if closure.GetFunc().GetType() == ZEND_USER_FUNCTION {
		//DestroyOpArray(closure.GetFunc().GetOpArray())
	} else if closure.GetOrigInternalHandler() == ZendClosureCallMagic {
		// types.ZendStringRelease(closure.GetFunc().GetFunctionName())
	}
	if closure.GetThisPtr().IsNotUndef() {
		// ZvalPtrDtor(closure.GetThisPtr())
	}
}
func ZendClosureNew(class_type *types.ClassEntry) *types.Object {
	var closure = NewZendClosure(class_type)
	return (*types.Object)(closure)
}
func ZendClosureClone(zobject *types.Zval) *types.Object {
	var closure *ZendClosure = (*ZendClosure)(zobject.Object())
	var result types.Zval
	ZendCreateClosure(&result, closure.GetFunc(), closure.GetFunc().GetScope(), closure.GetCalledScope(), closure.GetThisPtr())
	return result.Object()
}
func ZendClosureGetClosure(obj *types.Zval, ce_ptr **types.ClassEntry, fptr_ptr *types.IFunction, obj_ptr **types.Object) int {
	var closure *ZendClosure = (*ZendClosure)(obj.Object())
	*fptr_ptr = closure.GetFunc()
	*ce_ptr = closure.GetCalledScope()
	if closure.GetThisPtr().IsNotUndef() {
		*obj_ptr = closure.GetThisPtr().Object()
	} else {
		*obj_ptr = nil
	}
	return types.SUCCESS
}
func zim_Closure___construct(executeData *ZendExecuteData, return_value *types.Zval) {
	faults.ThrowError(nil, "Instantiation of 'Closure' is not allowed")
}
func ZendRegisterClosureCe() {
	ZendCeClosure = RegisterClass(&types.InternalClassDecl{
		Name:         "Closure",
		Functions:    ClosureFunctions,
		CreateObject: ZendClosureNew,
	})
	ZendCeClosure.SetIsFinal(true)
	ZendCeClosure.SetSerialize(ZendClassSerializeDeny)
	ZendCeClosure.SetUnserialize(ZendClassUnserializeDeny)

	ClosureHandlers = *types.NewObjectHandlersEx(StdObjectHandlersPtr, types.ObjectHandlersSetting{
		FreeObj:           ZendClosureFreeStorage,
		GetConstructor:    ZendClosureGetConstructor,
		GetMethod:         ZendClosureGetMethod,
		WriteProperty:     ZendClosureWriteProperty,
		ReadProperty:      ZendClosureReadProperty,
		GetPropertyPtrPtr: ZendClosureGetPropertyPtrPtr,
		HasProperty:       ZendClosureHasProperty,
		UnsetProperty:     ZendClosureUnsetProperty,
		CompareObjects:    ZendClosureCompareObjects,
		CloneObj:          ZendClosureClone,
		GetClosure:        ZendClosureGetClosure,
	})
}
func ZendClosureInternalHandler(executeData *ZendExecuteData, return_value *types.Zval) {
	var closure *ZendClosure = (*ZendClosure)(ZEND_CLOSURE_OBJECT(executeData.GetFunc()))
	closure.GetOrigInternalHandler()(executeData, return_value)
	// OBJ_RELEASE((*types.ZendObject)(closure))
	executeData.GetFunc() = nil
}
func ZendCreateClosure(res *types.Zval, func_ types.IFunction, scope *types.ClassEntry, called_scope *types.ClassEntry, this_ptr *types.Zval) {
	var thisPtr *types.Object
	if this_ptr != nil && this_ptr.IsNotUndef() {
		thisPtr = this_ptr.Object()
	}
	ZendCreateClosureEx(res, func_, scope, called_scope, thisPtr)
}
func ZendCreateClosureEx(res *types.Zval, func_ types.IFunction, scope *types.ClassEntry, called_scope *types.ClassEntry, thisPtr *types.Object) {
	var closure *ZendClosure
	ObjectInitEx(res, ZendCeClosure)
	closure = (*ZendClosure)(res.Object())
	if scope == nil && thisPtr != nil {
		/* use dummy scope if we're binding an object without specifying a scope */
		scope = ZendCeClosure
	}
	if func_.GetType() == ZEND_USER_FUNCTION {
		memcpy(closure.GetFunc(), func_, b.SizeOf("zend_op_array"))
		closure.GetFunc().SetIsClosure(true)
		closure.GetFunc().SetIsImmutable(false)
		if closure.GetFunc().GetOpArray().GetStaticVariables() != nil {
			closure.GetFunc().GetOpArray().SetStaticVariables(types.ZendArrayDup(closure.GetFunc().GetOpArray().GetStaticVariables()))
		}
		ZEND_MAP_PTR_INIT(closure.GetFunc().GetOpArray().static_variables_ptr, closure.GetFunc().GetOpArray().GetStaticVariables())

		/* Runtime cache is scope-dependent, so we cannot reuse it if the scope changed */

		if !closure.GetFunc().GetOpArray().HasInitRunTimeCache() || func_.GetScope() != scope || func_.IsHeapRtCache() {
			if func_.GetOpArray().HasInitRunTimeCache() && func_.IsClosure() && (func_.GetScope() == scope || !func_.IsImmutable()) {

				/* If a real closure is used for the first time, we create a shared runtime cache
				 * and remember which scope it is for. */

				if func_.GetScope() != scope {
					func_.SetScope(scope)
				}
				closure.GetFunc().GetOpArray().InitRunTimeCacheEx(false)
			} else {
				/* Otherwise, we use a non-shared runtime cache */
				closure.GetFunc().GetOpArray().InitRunTimeCacheEx(true)
			}
		}
		closure.GetFunc().GetOpArray().TryIncRefCount()
	} else {
		memcpy(closure.GetFunc(), func_, b.SizeOf("zend_internal_function"))
		closure.GetFunc().SetIsClosure(true)

		/* wrap internal function handler to avoid memory leak */

		if closure.GetFunc().GetInternalFunction().GetHandler() == ZendClosureInternalHandler {

			/* avoid infinity recursion, by taking handler from nested closure */

			var nested *ZendClosure = (*ZendClosure)((*byte)(func_ - zend_long((*byte)(&((*ZendClosure)(nil).GetFunc()))-(*byte)(nil))))
			b.Assert(nested.GetStd().GetCe() == ZendCeClosure)
			closure.SetOrigInternalHandler(nested.GetOrigInternalHandler())
		} else {
			closure.SetOrigInternalHandler(closure.GetFunc().GetInternalFunction().GetHandler())
		}
		closure.GetFunc().GetInternalFunction().SetHandler(ZendClosureInternalHandler)
		if func_.GetScope() == nil {
			/* if it's a free function, we won't set scope & this since they're meaningless */
			thisPtr = nil
			scope = nil
		}
	}
	closure.GetThisPtr().SetUndef()

	/* Invariant:
	 * If the closure is unscoped or static, it has no bound object. */

	closure.GetFunc().SetScope(scope)
	closure.SetCalledScope(called_scope)
	if scope != nil {
		closure.GetFunc().SetIsPublic(true)
		if thisPtr != nil && !closure.GetFunc().IsStatic() {
			closure.GetThisPtr().SetObject(thisPtr)
		}
	}
}
func ZendCreateFakeClosure(res *types.Zval, func_ types.IFunction, scope *types.ClassEntry, called_scope *types.ClassEntry, this_ptr *types.Zval) {
	var closure *ZendClosure
	ZendCreateClosure(res, func_, scope, called_scope, this_ptr)
	closure = (*ZendClosure)(res.Object())
	closure.GetFunc().SetIsFakeClosure(true)
}
func ZendClosureBindVarEx(closure_zv *types.Zval, offset uint32, val *types.Zval) {
	var closure *ZendClosure = (*ZendClosure)(closure_zv.Object())
	var static_variables *types.Array = closure.GetFunc().GetOpArray().GetStaticVariablesPtr()
	var var_ *types.Zval = (*types.Zval)((*byte)(static_variables.Bucket(offset)))
	// ZvalPtrDtor(var_)
	var_.CopyValueFrom(val)
}
