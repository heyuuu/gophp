// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZEND_CLOSURE_OBJECT(op_array *ZendFunction) *types.ZendObject {
	return (*types.ZendObject)((*byte)(op_array - b.SizeOf("zend_object")))
}
func ZEND_CLOSURE_PROPERTY_ERROR() {
	faults.ThrowError(nil, "Closure object cannot have properties")
}
func zim_Closure___invoke(executeData *ZendExecuteData, return_value *types.Zval) {
	var func_ *ZendFunction = executeData.GetFunc()
	var arguments *types.Zval = executeData.Arg(1)
	if CallUserFunction(nil, ZEND_THIS(executeData), return_value, executeData.NumArgs(), arguments) == types.FAILURE {
		return_value.SetFalse()
	}

	/* destruct the function also, then - we have allocated it in get_method */

	types.ZendStringReleaseEx(func_.GetInternalFunction().GetFunctionName(), 0)
	Efree(func_)
}
func ZendValidClosureBinding(closure *ZendClosure, newthis *types.Zval, scope *types.ClassEntry) types.ZendBool {
	var func_ *ZendFunction = closure.GetFunc()
	var is_fake_closure types.ZendBool = func_.IsFakeClosure()
	if newthis != nil {
		if func_.IsStatic() {
			faults.Error(faults.E_WARNING, "Cannot bind an instance to a static closure")
			return 0
		}
		if is_fake_closure != 0 && func_.GetScope() != nil && InstanceofFunction(types.Z_OBJCE_P(newthis), func_.GetScope()) == 0 {

			/* Binding incompatible $this to an internal method is not supported. */

			faults.Error(faults.E_WARNING, "Cannot bind method %s::%s() to object of class %s", func_.GetScope().GetName().GetVal(), func_.GetFunctionName().GetVal(), types.Z_OBJCE_P(newthis).GetName().GetVal())
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
	if scope != nil && scope != func_.GetScope() && scope.GetType() == ZEND_INTERNAL_CLASS {

		/* rebinding to internal class is not allowed */

		faults.Error(faults.E_WARNING, "Cannot bind closure to scope of internal class %s", scope.GetName().GetVal())
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
func zim_Closure_call(executeData *ZendExecuteData, return_value *types.Zval) {
	var newthis *types.Zval
	var closure_result types.Zval
	var closure *ZendClosure
	var fci types.ZendFcallInfo
	var fci_cache types.ZendFcallInfoCache
	var my_function ZendFunction
	var newobj *types.ZendObject
	fci.SetParamCount(0)
	fci.SetParams(nil)
	if ZendParseParameters(executeData.NumArgs(), "o*", &newthis, fci.GetParams(), fci.GetParamCount()) == types.FAILURE {
		return
	}
	closure = (*ZendClosure)(ZEND_THIS(executeData).GetObj())
	newobj = newthis.GetObj()
	if ZendValidClosureBinding(closure, newthis, types.Z_OBJCE_P(newthis)) == 0 {
		return
	}
	if closure.GetFunc().IsGenerator() {
		var new_closure types.Zval
		ZendCreateClosure(&new_closure, closure.GetFunc(), types.Z_OBJCE_P(newthis), closure.GetCalledScope(), newthis)
		closure = (*ZendClosure)(new_closure.GetObj())
		fci_cache.SetFunctionHandler(closure.GetFunc())
	} else {
		memcpy(&my_function, closure.GetFunc(), b.CondF(closure.GetFunc().GetType() == ZEND_USER_FUNCTION, func() __auto__ { return b.SizeOf("zend_op_array") }, func() __auto__ { return b.SizeOf("zend_internal_function") }))
		my_function.SetIsClosure(false)

		/* use scope of passed object */

		my_function.SetScope(types.Z_OBJCE_P(newthis))
		if closure.GetFunc().GetType() == ZEND_INTERNAL_FUNCTION {
			my_function.GetInternalFunction().SetHandler(closure.GetOrigInternalHandler())
		}
		fci_cache.SetFunctionHandler(&my_function)

		/* Runtime cache relies on bound scope to be immutable, hence we need a separate rt cache in case scope changed */

		if ZEND_USER_CODE(my_function.GetType()) && (closure.GetFunc().GetScope() != types.Z_OBJCE_P(newthis) || closure.GetFunc().IsHeapRtCache()) {
			var ptr any
			my_function.GetOpArray().SetIsHeapRtCache(true)
			ptr = Emalloc(b.SizeOf("void *") + my_function.GetOpArray().GetCacheSize())
			ZEND_MAP_PTR_INIT(my_function.GetOpArray().run_time_cache, ptr)
			ptr = (*byte)(ptr + b.SizeOf("void *"))
			ZEND_MAP_PTR_SET(my_function.GetOpArray().run_time_cache, ptr)
			memset(ptr, 0, my_function.GetOpArray().GetCacheSize())
		}

		/* Runtime cache relies on bound scope to be immutable, hence we need a separate rt cache in case scope changed */

	}
	fci_cache.SetCalledScope(newobj.GetCe())
	fci.SetObject(newobj)
	fci_cache.SetObject(fci.GetObject())
	fci.SetSize(b.SizeOf("fci"))
	fci.GetFunctionName().SetObject(closure.GetStd())
	fci.SetRetval(&closure_result)
	fci.SetNoSeparation(1)
	if ZendCallFunction(&fci, &fci_cache) == types.SUCCESS && closure_result.IsNotUndef() {
		if closure_result.IsReference() {
			ZendUnwrapReference(&closure_result)
		}
		types.ZVAL_COPY_VALUE(return_value, &closure_result)
	}
	if fci_cache.GetFunctionHandler().IsGenerator() {

		/* copied upon generator creation */

		closure.GetStd().DelRefcount()

		/* copied upon generator creation */

	} else if ZEND_USER_CODE(my_function.GetType()) && fci_cache.GetFunctionHandler().IsHeapRtCache() {
		Efree(my_function.GetOpArray().GetRunTimeCachePtr())
	}
}
func zim_Closure_bind(executeData *ZendExecuteData, return_value *types.Zval) {
	var newthis *types.Zval
	var zclosure *types.Zval
	var scope_arg *types.Zval = nil
	var closure *ZendClosure
	var ce *types.ClassEntry
	var called_scope *types.ClassEntry
	if ZendParseMethodParameters(executeData.NumArgs(), getThis(executeData), "Oo!|z", &zclosure, ZendCeClosure, &newthis, &scope_arg) == types.FAILURE {
		return
	}
	closure = (*ZendClosure)(zclosure.GetObj())
	if scope_arg != nil {
		if scope_arg.IsObject() {
			ce = types.Z_OBJCE_P(scope_arg)
		} else if scope_arg.IsNull() {
			ce = nil
		} else {
			var tmp_class_name *types.String
			var class_name *types.String = ZvalGetTmpString(scope_arg, &tmp_class_name)
			if types.ZendStringEqualsLiteral(class_name, "static") {
				ce = closure.GetFunc().GetScope()
			} else if b.Assign(&ce, ZendLookupClass(class_name)) == nil {
				faults.Error(faults.E_WARNING, "Class '%s' not found", class_name.GetVal())
				ZendTmpStringRelease(tmp_class_name)
				return_value.SetNull()
				return
			}
			ZendTmpStringRelease(tmp_class_name)
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
func ZendClosureCallMagic(executeData *ZendExecuteData, return_value *types.Zval) {
	var fci types.ZendFcallInfo
	var fcc types.ZendFcallInfoCache
	var params []types.Zval
	memset(&fci, 0, b.SizeOf("zend_fcall_info"))
	memset(&fcc, 0, b.SizeOf("zend_fcall_info_cache"))
	fci.SetSize(b.SizeOf("zend_fcall_info"))
	fci.SetRetval(return_value)
	if (executeData.GetFunc().internal_function.fn_flags & ZEND_ACC_STATIC) != 0 {
		fcc.SetFunctionHandler(executeData.GetFunc().internal_function.scope.__callstatic)
	} else {
		fcc.SetFunctionHandler(executeData.GetFunc().internal_function.scope.__call)
	}
	fci.SetParams(params)
	fci.SetParamCount(2)
	fci.GetParams()[0].SetString(executeData.GetFunc().common.function_name)
	if executeData.NumArgs() != 0 {
		ArrayInitSize(fci.GetParams()[1], executeData.NumArgs())
		ZendCopyParametersArray(executeData.NumArgs(), fci.GetParams()[1])
	} else {
		types.ZVAL_EMPTY_ARRAY(fci.GetParams()[1])
	}
	fci.SetObject(ZEND_THIS(executeData).GetObj())
	fcc.SetObject(fci.GetObject())
	fcc.SetCalledScope(ZendGetCalledScope(CurrEX()))
	ZendCallFunction(&fci, &fcc)
	ZvalPtrDtor(fci.GetParams()[1])
}
func ZendCreateClosureFromCallable(return_value *types.Zval, callable *types.Zval, error **byte) int {
	var fcc types.ZendFcallInfoCache
	var mptr *ZendFunction
	var instance types.Zval
	if ZendIsCallableEx(callable, nil, 0, nil, &fcc, error) == 0 {
		return types.FAILURE
	}
	mptr = fcc.GetFunctionHandler()
	if mptr.IsCallViaTrampoline() {

		/* For Closure::fromCallable([$closure, "__invoke"]) return $closure. */

		if fcc.GetObject() != nil && fcc.GetObject().GetCe() == ZendCeClosure && types.ZendStringEqualsLiteral(mptr.GetFunctionName(), "__invoke") {
			return_value.SetObject(fcc.GetObject())
			fcc.GetObject().AddRefcount()
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
		call := NewInternalFunctionEx(mptr.GetFunctionName().GetStr(), ZendClosureCallMagic)
		call.SetFnFlags(mptr.GetFnFlags() & ZEND_ACC_STATIC)
		call.SetScope(mptr.GetScope())
		ZendFreeTrampoline(mptr)
		mptr = NewZendFunctionInternal(call)
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
	if callable.IsObject() && InstanceofFunction(types.Z_OBJCE_P(callable), ZendCeClosure) != 0 {

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
			faults.TypeError("Failed to create closure from callable: %s", error)
			Efree(error)
		} else {
			faults.TypeError("Failed to create closure from callable")
		}
	}
}
func ZendClosureGetConstructor(object *types.ZendObject) *ZendFunction {
	faults.ThrowError(nil, "Instantiation of 'Closure' is not allowed")
	return nil
}
func ZendClosureCompareObjects(o1 *types.Zval, o2 *types.Zval) int { return o1.GetObj() != o2.GetObj() }
func ZendGetClosureInvokeMethod(object *types.ZendObject) *ZendFunction {
	var closure *ZendClosure = (*ZendClosure)(object)
	var invoke *ZendFunction = (*ZendFunction)(Emalloc(b.SizeOf("zend_function")))
	var keep_flags uint32 = ZEND_ACC_RETURN_REFERENCE | ZEND_ACC_VARIADIC | ZEND_ACC_HAS_RETURN_TYPE
	invoke.SetCommon(closure.GetFunc().GetCommon())

	/* We return ZEND_INTERNAL_FUNCTION, but arg_info representation is the
	 * same as for ZEND_USER_FUNCTION (uses zend_string* instead of char*).
	 * This is not a problem, because ZEND_ACC_HAS_TYPE_HINTS is never set,
	 * and we won't check arguments on internal function. We also set
	 * ZEND_ACC_USER_ARG_INFO flag to prevent invalid usage by Reflection */

	invoke.SetType(ZEND_INTERNAL_FUNCTION)
	invoke.GetInternalFunction().SetFnFlags(ZEND_ACC_PUBLIC | ZEND_ACC_CALL_VIA_HANDLER | closure.GetFunc().GetFnFlags()&keep_flags)
	if closure.GetFunc().GetType() != ZEND_INTERNAL_FUNCTION || closure.GetFunc().IsUserArgInfo() {
		invoke.GetInternalFunction().SetIsUserArgInfo(true)
	}
	invoke.GetInternalFunction().SetHandler(zim_Closure___invoke)
	invoke.GetInternalFunction().SetModule(0)
	invoke.GetInternalFunction().SetScope(ZendCeClosure)
	invoke.GetInternalFunction().SetFunctionName(types.ZSTR_MAGIC_INVOKE)
	return invoke
}
func ZendGetClosureMethodDef(obj *types.Zval) *ZendFunction {
	var closure *ZendClosure = (*ZendClosure)(obj.GetObj())
	return closure.GetFunc()
}
func ZendGetClosureThisPtr(obj *types.Zval) *types.Zval {
	var closure *ZendClosure = (*ZendClosure)(obj.GetObj())
	return closure.GetThisPtr()
}
func ZendClosureGetMethod(object **types.ZendObject, method *types.String, key *types.Zval) *ZendFunction {
	if types.ZendStringEqualsLiteralCi(method, ZEND_INVOKE_FUNC_NAME) {
		return ZendGetClosureInvokeMethod(*object)
	}
	return ZendStdGetMethod(object, method, key)
}
func ZendClosureReadProperty(object *types.Zval, member *types.Zval, type_ int, cache_slot *any, rv *types.Zval) *types.Zval {
	ZEND_CLOSURE_PROPERTY_ERROR()
	return EG__().GetUninitializedZval()
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
func ZendClosureFreeStorage(object *types.ZendObject) {
	var closure *ZendClosure = (*ZendClosure)(object)
	ZendObjectStdDtor(closure.GetStd())
	if closure.GetFunc().GetType() == ZEND_USER_FUNCTION {
		DestroyOpArray(closure.GetFunc().GetOpArray())
	} else if closure.GetOrigInternalHandler() == ZendClosureCallMagic {
		types.ZendStringRelease(closure.GetFunc().GetFunctionName())
	}
	if closure.GetThisPtr().IsNotUndef() {
		ZvalPtrDtor(closure.GetThisPtr())
	}
}
func ZendClosureNew(class_type *types.ClassEntry) *types.ZendObject {
	var closure *ZendClosure
	closure = Emalloc(b.SizeOf("zend_closure"))
	memset(closure, 0, b.SizeOf("zend_closure"))
	ZendObjectStdInit(closure.GetStd(), class_type)
	closure.GetStd().SetHandlers(&ClosureHandlers)
	return (*types.ZendObject)(closure)
}
func ZendClosureClone(zobject *types.Zval) *types.ZendObject {
	var closure *ZendClosure = (*ZendClosure)(zobject.GetObj())
	var result types.Zval
	ZendCreateClosure(&result, closure.GetFunc(), closure.GetFunc().GetScope(), closure.GetCalledScope(), closure.GetThisPtr())
	return result.GetObj()
}
func ZendClosureGetClosure(obj *types.Zval, ce_ptr **types.ClassEntry, fptr_ptr **ZendFunction, obj_ptr **types.ZendObject) int {
	var closure *ZendClosure = (*ZendClosure)(obj.GetObj())
	*fptr_ptr = closure.GetFunc()
	*ce_ptr = closure.GetCalledScope()
	if closure.GetThisPtr().IsNotUndef() {
		*obj_ptr = closure.GetThisPtr().GetObj()
	} else {
		*obj_ptr = nil
	}
	return types.SUCCESS
}
func ZendClosureGetDebugInfo(object *types.Zval, is_temp *int) *types.Array {
	var closure *ZendClosure = (*ZendClosure)(object.GetObj())
	var val types.Zval
	var arg_info *ZendArgInfo = closure.GetFunc().GetArgInfo()
	var debug_info *types.Array
	var zstr_args types.ZendBool = closure.GetFunc().GetType() == ZEND_USER_FUNCTION || closure.GetFunc().IsUserArgInfo()
	*is_temp = 1
	debug_info = types.NewZendArray(8)
	if closure.GetFunc().GetType() == ZEND_USER_FUNCTION && closure.GetFunc().GetOpArray().GetStaticVariables() != nil {
		var var_ *types.Zval
		var static_variables *types.Array = ZEND_MAP_PTR_GET(closure.GetFunc().GetOpArray().static_variables_ptr)
		val.SetArray(types.ZendArrayDup(static_variables))
		debug_info.KeyUpdate(types.ZSTR_STATIC.GetStr(), &val)
		var __ht *types.Array = val.GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			var_ = _z
			if var_.IsConstant() {
				ZvalPtrDtor(var_)
				var_.SetRawString("<constant ast>")
			}
		}
	}
	if closure.GetThisPtr().IsNotUndef() {
		closure.GetThisPtr().AddRefcount()
		debug_info.KeyUpdate(types.ZSTR_THIS.GetStr(), closure.GetThisPtr())
	}
	if arg_info != nil && (closure.GetFunc().GetNumArgs() != 0 || closure.GetFunc().IsVariadic()) {
		var i uint32
		var num_args uint32
		var required uint32 = closure.GetFunc().GetRequiredNumArgs()
		ArrayInit(&val)
		num_args = closure.GetFunc().GetNumArgs()
		if closure.GetFunc().IsVariadic() {
			num_args++
		}
		for i = 0; i < num_args; i++ {
			var name *types.String
			var info types.Zval
			if arg_info.GetName() != nil {
				if zstr_args != 0 {
					name = ZendStrpprintf(0, "%s$%s", b.Cond(arg_info.GetPassByReference() != 0, "&", ""), arg_info.GetName().GetVal())
				} else {
					name = ZendStrpprintf(0, "%s$%s", b.Cond(arg_info.GetPassByReference() != 0, "&", ""), (*ArgInfo)(arg_info).Name())
				}
			} else {
				name = ZendStrpprintf(0, "%s$param%d", b.Cond(arg_info.GetPassByReference() != 0, "&", ""), i+1)
			}
			info.SetString(ZendStrpprintf(0, "%s", b.Cond(i >= required, "<optional>", "<required>")))
			val.GetArr().KeyUpdate(name.GetStr(), &info)
			types.ZendStringReleaseEx(name, 0)
			arg_info++
		}
		debug_info.KeyUpdate("parameter", &val)
	}
	return debug_info
}
func ZendClosureGetGc(obj *types.Zval, table **types.Zval, n *int) *types.Array {
	var closure *ZendClosure = (*ZendClosure)(obj.GetObj())
	if closure.GetThisPtr().GetType() != types.IS_NULL {
		*table = closure.GetThisPtr()
	} else {
		*table = nil
	}
	if closure.GetThisPtr().GetType() != types.IS_NULL {
		*n = 1
	} else {
		*n = 0
	}
	if closure.GetFunc().GetType() == ZEND_USER_FUNCTION {
		return ZEND_MAP_PTR_GET(closure.GetFunc().GetOpArray().static_variables_ptr)
	} else {
		return nil
	}
}
func zim_Closure___construct(executeData *ZendExecuteData, return_value *types.Zval) {
	faults.ThrowError(nil, "Instantiation of 'Closure' is not allowed")
}
func ZendRegisterClosureCe() {
	var ce types.ClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("Closure", b.SizeOf("\"Closure\"")-1, 1))
	ce.SetBuiltinFunctions(ClosureFunctions)
	ZendCeClosure = ZendRegisterInternalClass(&ce)
	ZendCeClosure.SetIsFinal(true)
	ZendCeClosure.SetCreateObject(ZendClosureNew)
	ZendCeClosure.SetSerialize(ZendClassSerializeDeny)
	ZendCeClosure.SetUnserialize(ZendClassUnserializeDeny)
	memcpy(&ClosureHandlers, &StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	ClosureHandlers.SetFreeObj(ZendClosureFreeStorage)
	ClosureHandlers.SetGetConstructor(ZendClosureGetConstructor)
	ClosureHandlers.SetGetMethod(ZendClosureGetMethod)
	ClosureHandlers.SetWriteProperty(ZendClosureWriteProperty)
	ClosureHandlers.SetReadProperty(ZendClosureReadProperty)
	ClosureHandlers.SetGetPropertyPtrPtr(ZendClosureGetPropertyPtrPtr)
	ClosureHandlers.SetHasProperty(ZendClosureHasProperty)
	ClosureHandlers.SetUnsetProperty(ZendClosureUnsetProperty)
	ClosureHandlers.SetCompareObjects(ZendClosureCompareObjects)
	ClosureHandlers.SetCloneObj(ZendClosureClone)
	ClosureHandlers.SetGetDebugInfo(ZendClosureGetDebugInfo)
	ClosureHandlers.SetGetClosure(ZendClosureGetClosure)
	ClosureHandlers.SetGetGc(ZendClosureGetGc)
}
func ZendClosureInternalHandler(executeData *ZendExecuteData, return_value *types.Zval) {
	var closure *ZendClosure = (*ZendClosure)(ZEND_CLOSURE_OBJECT(executeData.GetFunc()))
	closure.GetOrigInternalHandler()(executeData, return_value)
	OBJ_RELEASE((*types.ZendObject)(closure))
	executeData.GetFunc() = nil
}
func ZendCreateClosure(res *types.Zval, func_ *ZendFunction, scope *types.ClassEntry, called_scope *types.ClassEntry, this_ptr *types.Zval) {
	var closure *ZendClosure
	ObjectInitEx(res, ZendCeClosure)
	closure = (*ZendClosure)(res.GetObj())
	if scope == nil && this_ptr != nil && this_ptr.IsNotUndef() {

		/* use dummy scope if we're binding an object without specifying a scope */

		scope = ZendCeClosure

		/* use dummy scope if we're binding an object without specifying a scope */

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

		if !(ZEND_MAP_PTR_GET(closure.GetFunc().GetOpArray().run_time_cache)) || func_.GetScope() != scope || func_.IsHeapRtCache() {
			var ptr any
			if !(ZEND_MAP_PTR_GET(func_.GetOpArray().run_time_cache)) && func_.IsClosure() && (func_.GetScope() == scope || !func_.IsImmutable()) {

				/* If a real closure is used for the first time, we create a shared runtime cache
				 * and remember which scope it is for. */

				if func_.GetScope() != scope {
					func_.SetScope(scope)
				}
				closure.GetFunc().GetOpArray().SetIsHeapRtCache(false)
				ptr = ZendArenaAlloc(CG__().GetArena(), func_.GetOpArray().GetCacheSize())
				ZEND_MAP_PTR_SET(func_.GetOpArray().run_time_cache, ptr)
				ZEND_MAP_PTR_SET(closure.GetFunc().GetOpArray().run_time_cache, ptr)
			} else {

				/* Otherwise, we use a non-shared runtime cache */

				closure.GetFunc().GetOpArray().SetIsHeapRtCache(true)
				ptr = Emalloc(b.SizeOf("void *") + func_.GetOpArray().GetCacheSize())
				ZEND_MAP_PTR_INIT(closure.GetFunc().GetOpArray().run_time_cache, ptr)
				ptr = (*byte)(ptr + b.SizeOf("void *"))
				ZEND_MAP_PTR_SET(closure.GetFunc().GetOpArray().run_time_cache, ptr)
			}
			memset(ptr, 0, func_.GetOpArray().GetCacheSize())
		}
		if closure.GetFunc().GetOpArray().GetRefcount() != nil {
			closure.func_.op_array.refcount++
		}
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

			this_ptr = nil
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
		if this_ptr != nil && this_ptr.IsObject() && !closure.GetFunc().IsStatic() {
			this_ptr.AddRefcount()
			closure.GetThisPtr().SetObject(this_ptr.GetObj())
		}
	}
}
func ZendCreateFakeClosure(res *types.Zval, func_ *ZendFunction, scope *types.ClassEntry, called_scope *types.ClassEntry, this_ptr *types.Zval) {
	var closure *ZendClosure
	ZendCreateClosure(res, func_, scope, called_scope, this_ptr)
	closure = (*ZendClosure)(res.GetObj())
	closure.GetFunc().SetIsFakeClosure(true)
}
func ZendClosureBindVar(closure_zv *types.Zval, var_name *types.String, var_ *types.Zval) {
	var closure *ZendClosure = (*ZendClosure)(closure_zv.GetObj())
	var static_variables *types.Array = ZEND_MAP_PTR_GET(closure.GetFunc().GetOpArray().static_variables_ptr)
	static_variables.KeyUpdate(var_name.GetStr(), var_)
}
func ZendClosureBindVarEx(closure_zv *types.Zval, offset uint32, val *types.Zval) {
	var closure *ZendClosure = (*ZendClosure)(closure_zv.GetObj())
	var static_variables *types.Array = ZEND_MAP_PTR_GET(closure.GetFunc().GetOpArray().static_variables_ptr)
	var var_ *types.Zval = (*types.Zval)((*byte)(static_variables.GetArData() + offset))
	ZvalPtrDtor(var_)
	types.ZVAL_COPY_VALUE(var_, val)
}
