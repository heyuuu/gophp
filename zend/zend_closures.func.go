// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZEND_CLOSURE_OBJECT(op_array *ZendFunction) *ZendObject {
	return (*ZendObject)((*byte)(op_array - b.SizeOf("zend_object")))
}
func ZEND_CLOSURE_PROPERTY_ERROR() {
	ZendThrowError(nil, "Closure object cannot have properties")
}
func zim_Closure___invoke(execute_data *ZendExecuteData, return_value *Zval) {
	var func_ *ZendFunction = EX(func_)
	var arguments *Zval = ZEND_CALL_ARG(execute_data, 1)
	if CallUserFunction(CG__().GetFunctionTable(), nil, ZEND_THIS, return_value, ZEND_NUM_ARGS(), arguments) == FAILURE {
		return_value.SetFalse()
	}

	/* destruct the function also, then - we have allocated it in get_method */

	ZendStringReleaseEx(func_.GetInternalFunction().GetFunctionName(), 0)
	Efree(func_)
}
func ZendValidClosureBinding(closure *ZendClosure, newthis *Zval, scope *ZendClassEntry) ZendBool {
	var func_ *ZendFunction = closure.GetFunc()
	var is_fake_closure ZendBool = func_.IsFakeClosure()
	if newthis != nil {
		if func_.IsStatic() {
			ZendError(E_WARNING, "Cannot bind an instance to a static closure")
			return 0
		}
		if is_fake_closure != 0 && func_.GetScope() != nil && InstanceofFunction(Z_OBJCE_P(newthis), func_.GetScope()) == 0 {

			/* Binding incompatible $this to an internal method is not supported. */

			ZendError(E_WARNING, "Cannot bind method %s::%s() to object of class %s", func_.GetScope().GetName().GetVal(), func_.GetFunctionName().GetVal(), Z_OBJCE_P(newthis).GetName().GetVal())
			return 0
		}
	} else if is_fake_closure != 0 && func_.GetScope() != nil && !func_.IsStatic() {
		if func_.GetType() == ZEND_INTERNAL_FUNCTION {
			ZendError(E_WARNING, "Cannot unbind $this of internal method")
			return 0
		} else {
			ZendError(E_DEPRECATED, "Unbinding $this of a method is deprecated")
		}
	} else if is_fake_closure == 0 && !(closure.GetThisPtr().IsUndef()) && func_.IsUsesThis() {

		// TODO: Only deprecate if it had $this *originally*?

		ZendError(E_DEPRECATED, "Unbinding $this of closure is deprecated")

		// TODO: Only deprecate if it had $this *originally*?

	}
	if scope != nil && scope != func_.GetScope() && scope.GetType() == ZEND_INTERNAL_CLASS {

		/* rebinding to internal class is not allowed */

		ZendError(E_WARNING, "Cannot bind closure to scope of internal class %s", scope.GetName().GetVal())
		return 0
	}
	if is_fake_closure != 0 && scope != func_.GetScope() {
		if func_.GetScope() == nil {
			ZendError(E_WARNING, "Cannot rebind scope of closure created from function")
		} else {
			ZendError(E_WARNING, "Cannot rebind scope of closure created from method")
		}
		return 0
	}
	return 1
}
func zim_Closure_call(execute_data *ZendExecuteData, return_value *Zval) {
	var newthis *Zval
	var closure_result Zval
	var closure *ZendClosure
	var fci ZendFcallInfo
	var fci_cache ZendFcallInfoCache
	var my_function ZendFunction
	var newobj *ZendObject
	fci.SetParamCount(0)
	fci.SetParams(nil)
	if ZendParseParameters(ZEND_NUM_ARGS(), "o*", &newthis, fci.GetParams(), fci.GetParamCount()) == FAILURE {
		return
	}
	closure = (*ZendClosure)(ZEND_THIS.GetObj())
	newobj = newthis.GetObj()
	if ZendValidClosureBinding(closure, newthis, Z_OBJCE_P(newthis)) == 0 {
		return
	}
	if closure.GetFunc().IsGenerator() {
		var new_closure Zval
		ZendCreateClosure(&new_closure, closure.GetFunc(), Z_OBJCE_P(newthis), closure.GetCalledScope(), newthis)
		closure = (*ZendClosure)(new_closure.GetObj())
		fci_cache.SetFunctionHandler(closure.GetFunc())
	} else {
		memcpy(&my_function, closure.GetFunc(), b.CondF(closure.GetFunc().GetType() == ZEND_USER_FUNCTION, func() __auto__ { return b.SizeOf("zend_op_array") }, func() __auto__ { return b.SizeOf("zend_internal_function") }))
		my_function.SetIsClosure(false)

		/* use scope of passed object */

		my_function.SetScope(Z_OBJCE_P(newthis))
		if closure.GetFunc().GetType() == ZEND_INTERNAL_FUNCTION {
			my_function.GetInternalFunction().SetHandler(closure.GetOrigInternalHandler())
		}
		fci_cache.SetFunctionHandler(&my_function)

		/* Runtime cache relies on bound scope to be immutable, hence we need a separate rt cache in case scope changed */

		if ZEND_USER_CODE(my_function.GetType()) && (closure.GetFunc().GetScope() != Z_OBJCE_P(newthis) || closure.GetFunc().IsHeapRtCache()) {
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
	if ZendCallFunction(&fci, &fci_cache) == SUCCESS && closure_result.GetType() != IS_UNDEF {
		if closure_result.IsReference() {
			ZendUnwrapReference(&closure_result)
		}
		ZVAL_COPY_VALUE(return_value, &closure_result)
	}
	if fci_cache.GetFunctionHandler().IsGenerator() {

		/* copied upon generator creation */

		closure.GetStd().DelRefcount()

		/* copied upon generator creation */

	} else if ZEND_USER_CODE(my_function.GetType()) && fci_cache.GetFunctionHandler().IsHeapRtCache() {
		Efree(my_function.GetOpArray().GetRunTimeCachePtr())
	}
}
func zim_Closure_bind(execute_data *ZendExecuteData, return_value *Zval) {
	var newthis *Zval
	var zclosure *Zval
	var scope_arg *Zval = nil
	var closure *ZendClosure
	var ce *ZendClassEntry
	var called_scope *ZendClassEntry
	if ZendParseMethodParameters(ZEND_NUM_ARGS(), getThis(), "Oo!|z", &zclosure, ZendCeClosure, &newthis, &scope_arg) == FAILURE {
		return
	}
	closure = (*ZendClosure)(zclosure.GetObj())
	if scope_arg != nil {
		if scope_arg.IsObject() {
			ce = Z_OBJCE_P(scope_arg)
		} else if scope_arg.IsNull() {
			ce = nil
		} else {
			var tmp_class_name *ZendString
			var class_name *ZendString = ZvalGetTmpString(scope_arg, &tmp_class_name)
			if ZendStringEqualsLiteral(class_name, "static") {
				ce = closure.GetFunc().GetScope()
			} else if b.Assign(&ce, ZendLookupClass(class_name)) == nil {
				ZendError(E_WARNING, "Class '%s' not found", class_name.GetVal())
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
		called_scope = Z_OBJCE_P(newthis)
	} else {
		called_scope = ce
	}
	ZendCreateClosure(return_value, closure.GetFunc(), ce, called_scope, newthis)
}
func ZendClosureCallMagic(execute_data *ZendExecuteData, return_value *Zval) {
	var fci ZendFcallInfo
	var fcc ZendFcallInfoCache
	var params []Zval
	memset(&fci, 0, b.SizeOf("zend_fcall_info"))
	memset(&fcc, 0, b.SizeOf("zend_fcall_info_cache"))
	fci.SetSize(b.SizeOf("zend_fcall_info"))
	fci.SetRetval(return_value)
	if (EX(func_).internal_function.fn_flags & ZEND_ACC_STATIC) != 0 {
		fcc.SetFunctionHandler(EX(func_).internal_function.scope.__callstatic)
	} else {
		fcc.SetFunctionHandler(EX(func_).internal_function.scope.__call)
	}
	fci.SetParams(params)
	fci.SetParamCount(2)
	fci.GetParams()[0].SetString(EX(func_).common.function_name)
	if ZEND_NUM_ARGS() != 0 {
		ArrayInitSize(fci.GetParams()[1], ZEND_NUM_ARGS())
		ZendCopyParametersArray(ZEND_NUM_ARGS(), fci.GetParams()[1])
	} else {
		ZVAL_EMPTY_ARRAY(fci.GetParams()[1])
	}
	fci.SetObject(ZEND_THIS.GetObj())
	fcc.SetObject(fci.GetObject())
	fcc.SetCalledScope(ZendGetCalledScope(EG__().GetCurrentExecuteData()))
	ZendCallFunction(&fci, &fcc)
	ZvalPtrDtor(fci.GetParams()[1])
}
func ZendCreateClosureFromCallable(return_value *Zval, callable *Zval, error **byte) int {
	var fcc ZendFcallInfoCache
	var mptr *ZendFunction
	var instance Zval
	var call ZendInternalFunction
	if ZendIsCallableEx(callable, nil, 0, nil, &fcc, error) == 0 {
		return FAILURE
	}
	mptr = fcc.GetFunctionHandler()
	if mptr.IsCallViaTrampoline() {

		/* For Closure::fromCallable([$closure, "__invoke"]) return $closure. */

		if fcc.GetObject() != nil && fcc.GetObject().GetCe() == ZendCeClosure && ZendStringEqualsLiteral(mptr.GetFunctionName(), "__invoke") {
			return_value.SetObject(fcc.GetObject())
			fcc.GetObject().AddRefcount()
			ZendFreeTrampoline(mptr)
			return SUCCESS
		}
		if mptr.GetScope() == nil {
			return FAILURE
		}
		if mptr.IsStatic() {
			if mptr.GetScope().GetCallstatic() == nil {
				return FAILURE
			}
		} else {
			if mptr.GetScope().GetCall() == nil {
				return FAILURE
			}
		}
		memset(&call, 0, b.SizeOf("zend_internal_function"))
		call.SetType(ZEND_INTERNAL_FUNCTION)
		call.SetFnFlags(mptr.GetFnFlags() & ZEND_ACC_STATIC)
		call.SetHandler(ZendClosureCallMagic)
		call.SetFunctionName(mptr.GetFunctionName())
		call.SetScope(mptr.GetScope())
		ZendFreeTrampoline(mptr)
		mptr = (*ZendFunction)(&call)
	}
	if fcc.GetObject() != nil {
		instance.SetObject(fcc.GetObject())
		ZendCreateFakeClosure(return_value, mptr, mptr.GetScope(), fcc.GetCalledScope(), &instance)
	} else {
		ZendCreateFakeClosure(return_value, mptr, mptr.GetScope(), fcc.GetCalledScope(), nil)
	}
	return SUCCESS
}
func zim_Closure_fromCallable(execute_data *ZendExecuteData, return_value *Zval) {
	var callable *Zval
	var success int
	var error *byte = nil
	if ZendParseParameters(ZEND_NUM_ARGS(), "z", &callable) == FAILURE {
		return
	}
	if callable.IsObject() && InstanceofFunction(Z_OBJCE_P(callable), ZendCeClosure) != 0 {

		/* It's already a closure */

		ZVAL_ZVAL(return_value, callable, 1, 0)
		return
	}

	/* create closure as if it were called from parent scope */

	EG__().SetCurrentExecuteData(EX(prev_execute_data))
	success = ZendCreateClosureFromCallable(return_value, callable, &error)
	EG__().SetCurrentExecuteData(execute_data)
	if success == FAILURE || error != nil {
		if error != nil {
			ZendTypeError("Failed to create closure from callable: %s", error)
			Efree(error)
		} else {
			ZendTypeError("Failed to create closure from callable")
		}
	}
}
func ZendClosureGetConstructor(object *ZendObject) *ZendFunction {
	ZendThrowError(nil, "Instantiation of 'Closure' is not allowed")
	return nil
}
func ZendClosureCompareObjects(o1 *Zval, o2 *Zval) int { return o1.GetObj() != o2.GetObj() }
func ZendGetClosureInvokeMethod(object *ZendObject) *ZendFunction {
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
	invoke.GetInternalFunction().SetFunctionName(ZSTR_KNOWN(ZEND_STR_MAGIC_INVOKE))
	return invoke
}
func ZendGetClosureMethodDef(obj *Zval) *ZendFunction {
	var closure *ZendClosure = (*ZendClosure)(obj.GetObj())
	return closure.GetFunc()
}
func ZendGetClosureThisPtr(obj *Zval) *Zval {
	var closure *ZendClosure = (*ZendClosure)(obj.GetObj())
	return closure.GetThisPtr()
}
func ZendClosureGetMethod(object **ZendObject, method *ZendString, key *Zval) *ZendFunction {
	if ZendStringEqualsLiteralCi(method, ZEND_INVOKE_FUNC_NAME) {
		return ZendGetClosureInvokeMethod(*object)
	}
	return ZendStdGetMethod(object, method, key)
}
func ZendClosureReadProperty(object *Zval, member *Zval, type_ int, cache_slot *any, rv *Zval) *Zval {
	ZEND_CLOSURE_PROPERTY_ERROR()
	return EG__().GetUninitializedZval()
}
func ZendClosureWriteProperty(object *Zval, member *Zval, value *Zval, cache_slot *any) *Zval {
	ZEND_CLOSURE_PROPERTY_ERROR()
	return EG__().GetErrorZval()
}
func ZendClosureGetPropertyPtrPtr(object *Zval, member *Zval, type_ int, cache_slot *any) *Zval {
	ZEND_CLOSURE_PROPERTY_ERROR()
	return nil
}
func ZendClosureHasProperty(object *Zval, member *Zval, has_set_exists int, cache_slot *any) int {
	if has_set_exists != ZEND_PROPERTY_EXISTS {
		ZEND_CLOSURE_PROPERTY_ERROR()
	}
	return 0
}
func ZendClosureUnsetProperty(object *Zval, member *Zval, cache_slot *any) {
	ZEND_CLOSURE_PROPERTY_ERROR()
}
func ZendClosureFreeStorage(object *ZendObject) {
	var closure *ZendClosure = (*ZendClosure)(object)
	ZendObjectStdDtor(closure.GetStd())
	if closure.GetFunc().GetType() == ZEND_USER_FUNCTION {
		DestroyOpArray(closure.GetFunc().GetOpArray())
	} else if closure.GetOrigInternalHandler() == ZendClosureCallMagic {
		ZendStringRelease(closure.GetFunc().GetFunctionName())
	}
	if closure.GetThisPtr().GetType() != IS_UNDEF {
		ZvalPtrDtor(closure.GetThisPtr())
	}
}
func ZendClosureNew(class_type *ZendClassEntry) *ZendObject {
	var closure *ZendClosure
	closure = Emalloc(b.SizeOf("zend_closure"))
	memset(closure, 0, b.SizeOf("zend_closure"))
	ZendObjectStdInit(closure.GetStd(), class_type)
	closure.GetStd().SetHandlers(&ClosureHandlers)
	return (*ZendObject)(closure)
}
func ZendClosureClone(zobject *Zval) *ZendObject {
	var closure *ZendClosure = (*ZendClosure)(zobject.GetObj())
	var result Zval
	ZendCreateClosure(&result, closure.GetFunc(), closure.GetFunc().GetScope(), closure.GetCalledScope(), closure.GetThisPtr())
	return result.GetObj()
}
func ZendClosureGetClosure(obj *Zval, ce_ptr **ZendClassEntry, fptr_ptr **ZendFunction, obj_ptr **ZendObject) int {
	var closure *ZendClosure = (*ZendClosure)(obj.GetObj())
	*fptr_ptr = closure.GetFunc()
	*ce_ptr = closure.GetCalledScope()
	if closure.GetThisPtr().GetType() != IS_UNDEF {
		*obj_ptr = closure.GetThisPtr().GetObj()
	} else {
		*obj_ptr = nil
	}
	return SUCCESS
}
func ZendClosureGetDebugInfo(object *Zval, is_temp *int) *HashTable {
	var closure *ZendClosure = (*ZendClosure)(object.GetObj())
	var val Zval
	var arg_info *ZendArgInfo = closure.GetFunc().GetArgInfo()
	var debug_info *HashTable
	var zstr_args ZendBool = closure.GetFunc().GetType() == ZEND_USER_FUNCTION || closure.GetFunc().IsUserArgInfo()
	*is_temp = 1
	debug_info = ZendNewArray(8)
	if closure.GetFunc().GetType() == ZEND_USER_FUNCTION && closure.GetFunc().GetOpArray().GetStaticVariables() != nil {
		var var_ *Zval
		var static_variables *HashTable = ZEND_MAP_PTR_GET(closure.GetFunc().GetOpArray().static_variables_ptr)
		val.SetArray(ZendArrayDup(static_variables))
		debug_info.KeyUpdate(ZSTR_KNOWN(ZEND_STR_STATIC).GetStr(), &val)
		var __ht *HashTable = val.GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			var_ = _z
			if var_.IsConstant() {
				ZvalPtrDtor(var_)
				ZVAL_STRING(var_, "<constant ast>")
			}
		}
	}
	if closure.GetThisPtr().GetType() != IS_UNDEF {
		closure.GetThisPtr().AddRefcount()
		debug_info.KeyUpdate(ZSTR_KNOWN(ZEND_STR_THIS).GetStr(), closure.GetThisPtr())
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
			var name *ZendString
			var info Zval
			if arg_info.GetName() != nil {
				if zstr_args != 0 {
					name = ZendStrpprintf(0, "%s$%s", b.Cond(arg_info.GetPassByReference() != 0, "&", ""), arg_info.GetName().GetVal())
				} else {
					name = ZendStrpprintf(0, "%s$%s", b.Cond(arg_info.GetPassByReference() != 0, "&", ""), (*ZendInternalArgInfo)(arg_info).GetName())
				}
			} else {
				name = ZendStrpprintf(0, "%s$param%d", b.Cond(arg_info.GetPassByReference() != 0, "&", ""), i+1)
			}
			info.SetString(ZendStrpprintf(0, "%s", b.Cond(i >= required, "<optional>", "<required>")))
			val.GetArr().KeyUpdate(name.GetStr(), &info)
			ZendStringReleaseEx(name, 0)
			arg_info++
		}
		debug_info.KeyUpdate("parameter", &val)
	}
	return debug_info
}
func ZendClosureGetGc(obj *Zval, table **Zval, n *int) *HashTable {
	var closure *ZendClosure = (*ZendClosure)(obj.GetObj())
	if closure.GetThisPtr().GetType() != IS_NULL {
		*table = closure.GetThisPtr()
	} else {
		*table = nil
	}
	if closure.GetThisPtr().GetType() != IS_NULL {
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
func zim_Closure___construct(execute_data *ZendExecuteData, return_value *Zval) {
	ZendThrowError(nil, "Instantiation of 'Closure' is not allowed")
}
func ZendRegisterClosureCe() {
	var ce ZendClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Closure", b.SizeOf("\"Closure\"")-1, 1))
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
func ZendClosureInternalHandler(execute_data *ZendExecuteData, return_value *Zval) {
	var closure *ZendClosure = (*ZendClosure)(ZEND_CLOSURE_OBJECT(EX(func_)))
	closure.GetOrigInternalHandler()(execute_data, return_value)
	OBJ_RELEASE((*ZendObject)(closure))
	EX(func_) = nil
}
func ZendCreateClosure(res *Zval, func_ *ZendFunction, scope *ZendClassEntry, called_scope *ZendClassEntry, this_ptr *Zval) {
	var closure *ZendClosure
	ObjectInitEx(res, ZendCeClosure)
	closure = (*ZendClosure)(res.GetObj())
	if scope == nil && this_ptr != nil && this_ptr.GetType() != IS_UNDEF {

		/* use dummy scope if we're binding an object without specifying a scope */

		scope = ZendCeClosure

		/* use dummy scope if we're binding an object without specifying a scope */

	}
	if func_.GetType() == ZEND_USER_FUNCTION {
		memcpy(closure.GetFunc(), func_, b.SizeOf("zend_op_array"))
		closure.GetFunc().SetIsClosure(true)
		closure.GetFunc().SetIsImmutable(false)
		if closure.GetFunc().GetOpArray().GetStaticVariables() != nil {
			closure.GetFunc().GetOpArray().SetStaticVariables(ZendArrayDup(closure.GetFunc().GetOpArray().GetStaticVariables()))
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
			ZEND_ASSERT(nested.GetStd().GetCe() == ZendCeClosure)
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
func ZendCreateFakeClosure(res *Zval, func_ *ZendFunction, scope *ZendClassEntry, called_scope *ZendClassEntry, this_ptr *Zval) {
	var closure *ZendClosure
	ZendCreateClosure(res, func_, scope, called_scope, this_ptr)
	closure = (*ZendClosure)(res.GetObj())
	closure.GetFunc().SetIsFakeClosure(true)
}
func ZendClosureBindVar(closure_zv *Zval, var_name *ZendString, var_ *Zval) {
	var closure *ZendClosure = (*ZendClosure)(closure_zv.GetObj())
	var static_variables *HashTable = ZEND_MAP_PTR_GET(closure.GetFunc().GetOpArray().static_variables_ptr)
	static_variables.KeyUpdate(var_name.GetStr(), var_)
}
func ZendClosureBindVarEx(closure_zv *Zval, offset uint32, val *Zval) {
	var closure *ZendClosure = (*ZendClosure)(closure_zv.GetObj())
	var static_variables *HashTable = ZEND_MAP_PTR_GET(closure.GetFunc().GetOpArray().static_variables_ptr)
	var var_ *Zval = (*Zval)((*byte)(static_variables.GetArData() + offset))
	ZvalPtrDtor(var_)
	ZVAL_COPY_VALUE(var_, val)
}
