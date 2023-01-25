// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_closures.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Christian Seiler <chris_se@gmx.net>                         |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define ZEND_CLOSURES_H

/* This macro depends on zend_closure structure layout */

func ZEND_CLOSURE_OBJECT(op_array *ZendFunction) *ZendObject {
	return (*ZendObject)((*byte)(op_array - b.SizeOf("zend_object")))
}

var ZendCeClosure *ZendClassEntry

// Source: <Zend/zend_closures.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Christian Seiler <chris_se@gmx.net>                         |
   |          Dmitry Stogov <dmitry@php.net>                              |
   |          Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_API.h"

// # include "zend_closures.h"

// # include "zend_exceptions.h"

// # include "zend_interfaces.h"

// # include "zend_objects.h"

// # include "zend_objects_API.h"

// # include "zend_globals.h"

const ZEND_CLOSURE_PRINT_NAME = "Closure object"

func ZEND_CLOSURE_PROPERTY_ERROR() {
	ZendThrowError(nil, "Closure object cannot have properties")
}

/* non-static since it needs to be referenced */

var ClosureHandlers ZendObjectHandlers

func zim_Closure___invoke(execute_data *ZendExecuteData, return_value *Zval) {
	var func_ *ZendFunction = EX(func_)
	var arguments *Zval = ZEND_CALL_ARG(execute_data, 1)
	if CallUserFunction(CompilerGlobals.GetFunctionTable(), nil, ZEND_THIS, return_value, ZEND_NUM_ARGS(), arguments) == FAILURE {
		RETVAL_FALSE
	}

	/* destruct the function also, then - we have allocated it in get_method */

	ZendStringReleaseEx(func_.GetInternalFunction().GetFunctionName(), 0)
	Efree(func_)
}

/* }}} */

func ZendValidClosureBinding(closure *ZendClosure, newthis *Zval, scope *ZendClassEntry) ZendBool {
	var func_ *ZendFunction = &closure.func_
	var is_fake_closure ZendBool = (func_.GetFnFlags() & ZEND_ACC_FAKE_CLOSURE) != 0
	if newthis != nil {
		if (func_.GetFnFlags() & ZEND_ACC_STATIC) != 0 {
			ZendError(E_WARNING, "Cannot bind an instance to a static closure")
			return 0
		}
		if is_fake_closure != 0 && func_.GetScope() != nil && InstanceofFunction(Z_OBJCE_P(newthis), func_.GetScope()) == 0 {

			/* Binding incompatible $this to an internal method is not supported. */

			ZendError(E_WARNING, "Cannot bind method %s::%s() to object of class %s", ZSTR_VAL(func_.GetScope().GetName()), ZSTR_VAL(func_.GetFunctionName()), ZSTR_VAL(Z_OBJCE_P(newthis).GetName()))
			return 0
		}
	} else if is_fake_closure != 0 && func_.GetScope() != nil && (func_.GetFnFlags()&ZEND_ACC_STATIC) == 0 {
		if func_.GetType() == ZEND_INTERNAL_FUNCTION {
			ZendError(E_WARNING, "Cannot unbind $this of internal method")
			return 0
		} else {
			ZendError(E_DEPRECATED, "Unbinding $this of a method is deprecated")
		}
	} else if is_fake_closure == 0 && !(Z_ISUNDEF(closure.GetThisPtr())) && (func_.GetFnFlags()&ZEND_ACC_USES_THIS) != 0 {

		// TODO: Only deprecate if it had $this *originally*?

		ZendError(E_DEPRECATED, "Unbinding $this of closure is deprecated")

		// TODO: Only deprecate if it had $this *originally*?

	}
	if scope != nil && scope != func_.GetScope() && scope.GetType() == ZEND_INTERNAL_CLASS {

		/* rebinding to internal class is not allowed */

		ZendError(E_WARNING, "Cannot bind closure to scope of internal class %s", ZSTR_VAL(scope.GetName()))
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

/* }}} */

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
	if ZendParseParameters(ZEND_NUM_ARGS(), "o*", &newthis, &fci.params, &fci.param_count) == FAILURE {
		return
	}
	closure = (*ZendClosure)(Z_OBJ_P(ZEND_THIS))
	newobj = Z_OBJ_P(newthis)
	if ZendValidClosureBinding(closure, newthis, Z_OBJCE_P(newthis)) == 0 {
		return
	}
	if (closure.GetFunc().GetFnFlags() & ZEND_ACC_GENERATOR) != 0 {
		var new_closure Zval
		ZendCreateClosure(&new_closure, &closure.func_, Z_OBJCE_P(newthis), closure.GetCalledScope(), newthis)
		closure = (*ZendClosure)(Z_OBJ(new_closure))
		fci_cache.SetFunctionHandler(&closure.func_)
	} else {
		memcpy(&my_function, &closure.func_, b.CondF(closure.GetFunc().GetType() == ZEND_USER_FUNCTION, func() __auto__ { return b.SizeOf("zend_op_array") }, func() __auto__ { return b.SizeOf("zend_internal_function") }))
		my_function.SetFnFlags(my_function.GetFnFlags() &^ ZEND_ACC_CLOSURE)

		/* use scope of passed object */

		my_function.SetScope(Z_OBJCE_P(newthis))
		if closure.GetFunc().GetType() == ZEND_INTERNAL_FUNCTION {
			my_function.GetInternalFunction().SetHandler(closure.GetOrigInternalHandler())
		}
		fci_cache.SetFunctionHandler(&my_function)

		/* Runtime cache relies on bound scope to be immutable, hence we need a separate rt cache in case scope changed */

		if ZEND_USER_CODE(my_function.GetType()) && (closure.GetFunc().GetScope() != Z_OBJCE_P(newthis) || (closure.GetFunc().GetFnFlags()&ZEND_ACC_HEAP_RT_CACHE) != 0) {
			var ptr any
			my_function.GetOpArray().SetFnFlags(my_function.GetOpArray().GetFnFlags() | ZEND_ACC_HEAP_RT_CACHE)
			ptr = Emalloc(b.SizeOf("void *") + my_function.GetOpArray().GetCacheSize())
			ZEND_MAP_PTR_INIT(my_function.op_array.run_time_cache, ptr)
			ptr = (*byte)(ptr + b.SizeOf("void *"))
			ZEND_MAP_PTR_SET(my_function.op_array.run_time_cache, ptr)
			memset(ptr, 0, my_function.GetOpArray().GetCacheSize())
		}

		/* Runtime cache relies on bound scope to be immutable, hence we need a separate rt cache in case scope changed */

	}
	fci_cache.SetCalledScope(newobj.GetCe())
	fci.SetObject(newobj)
	fci_cache.SetObject(fci.GetObject())
	fci.SetSize(b.SizeOf("fci"))
	ZVAL_OBJ(&fci.function_name, &closure.std)
	fci.SetRetval(&closure_result)
	fci.SetNoSeparation(1)
	if ZendCallFunction(&fci, &fci_cache) == SUCCESS && Z_TYPE(closure_result) != IS_UNDEF {
		if Z_ISREF(closure_result) {
			ZendUnwrapReference(&closure_result)
		}
		ZVAL_COPY_VALUE(return_value, &closure_result)
	}
	if (fci_cache.GetFunctionHandler().GetFnFlags() & ZEND_ACC_GENERATOR) != 0 {

		/* copied upon generator creation */

		GC_DELREF(&closure.std)

		/* copied upon generator creation */

	} else if ZEND_USER_CODE(my_function.GetType()) && (fci_cache.GetFunctionHandler().GetFnFlags()&ZEND_ACC_HEAP_RT_CACHE) != 0 {
		Efree(my_function.GetOpArray().GetRunTimeCachePtr())
	}
}

/* }}} */

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
	closure = (*ZendClosure)(Z_OBJ_P(zclosure))
	if scope_arg != nil {
		if Z_TYPE_P(scope_arg) == IS_OBJECT {
			ce = Z_OBJCE_P(scope_arg)
		} else if Z_TYPE_P(scope_arg) == IS_NULL {
			ce = nil
		} else {
			var tmp_class_name *ZendString
			var class_name *ZendString = ZvalGetTmpString(scope_arg, &tmp_class_name)
			if ZendStringEqualsLiteral(class_name, "static") {
				ce = closure.GetFunc().GetScope()
			} else if b.Assign(&ce, ZendLookupClass(class_name)) == nil {
				ZendError(E_WARNING, "Class '%s' not found", ZSTR_VAL(class_name))
				ZendTmpStringRelease(tmp_class_name)
				RETVAL_NULL()
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
	ZendCreateClosure(return_value, &closure.func_, ce, called_scope, newthis)
}

/* }}} */

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
	ZVAL_STR(&fci.params[0], EX(func_).common.function_name)
	if ZEND_NUM_ARGS() != 0 {
		ArrayInitSize(&fci.params[1], ZEND_NUM_ARGS())
		ZendCopyParametersArray(ZEND_NUM_ARGS(), &fci.params[1])
	} else {
		ZVAL_EMPTY_ARRAY(&fci.params[1])
	}
	fci.SetObject(Z_OBJ_P(ZEND_THIS))
	fcc.SetObject(fci.GetObject())
	fcc.SetCalledScope(ZendGetCalledScope(ExecutorGlobals.GetCurrentExecuteData()))
	ZendCallFunction(&fci, &fcc)
	ZvalPtrDtor(&fci.params[1])
}

/* }}} */

func ZendCreateClosureFromCallable(return_value *Zval, callable *Zval, error **byte) int {
	var fcc ZendFcallInfoCache
	var mptr *ZendFunction
	var instance Zval
	var call ZendInternalFunction
	if ZendIsCallableEx(callable, nil, 0, nil, &fcc, error) == 0 {
		return FAILURE
	}
	mptr = fcc.GetFunctionHandler()
	if (mptr.GetFnFlags() & ZEND_ACC_CALL_VIA_TRAMPOLINE) != 0 {

		/* For Closure::fromCallable([$closure, "__invoke"]) return $closure. */

		if fcc.GetObject() != nil && fcc.GetObject().GetCe() == ZendCeClosure && ZendStringEqualsLiteral(mptr.GetFunctionName(), "__invoke") {
			ZVAL_OBJ(return_value, fcc.GetObject())
			GC_ADDREF(fcc.GetObject())
			ZendFreeTrampoline(mptr)
			return SUCCESS
		}
		if mptr.GetScope() == nil {
			return FAILURE
		}
		if (mptr.GetFnFlags() & ZEND_ACC_STATIC) != 0 {
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
		ZVAL_OBJ(&instance, fcc.GetObject())
		ZendCreateFakeClosure(return_value, mptr, mptr.GetScope(), fcc.GetCalledScope(), &instance)
	} else {
		ZendCreateFakeClosure(return_value, mptr, mptr.GetScope(), fcc.GetCalledScope(), nil)
	}
	return SUCCESS
}

/* }}} */

func zim_Closure_fromCallable(execute_data *ZendExecuteData, return_value *Zval) {
	var callable *Zval
	var success int
	var error *byte = nil
	if ZendParseParameters(ZEND_NUM_ARGS(), "z", &callable) == FAILURE {
		return
	}
	if Z_TYPE_P(callable) == IS_OBJECT && InstanceofFunction(Z_OBJCE_P(callable), ZendCeClosure) != 0 {

		/* It's already a closure */

		RETVAL_ZVAL(callable, 1, 0)
		return
	}

	/* create closure as if it were called from parent scope */

	ExecutorGlobals.SetCurrentExecuteData(EX(prev_execute_data))
	success = ZendCreateClosureFromCallable(return_value, callable, &error)
	ExecutorGlobals.SetCurrentExecuteData(execute_data)
	if success == FAILURE || error != nil {
		if error != nil {
			ZendTypeError("Failed to create closure from callable: %s", error)
			Efree(error)
		} else {
			ZendTypeError("Failed to create closure from callable")
		}
	}
}

/* }}} */

func ZendClosureGetConstructor(object *ZendObject) *ZendFunction {
	ZendThrowError(nil, "Instantiation of 'Closure' is not allowed")
	return nil
}

/* }}} */

func ZendClosureCompareObjects(o1 *Zval, o2 *Zval) int { return Z_OBJ_P(o1) != Z_OBJ_P(o2) }

/* }}} */

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
	if closure.GetFunc().GetType() != ZEND_INTERNAL_FUNCTION || (closure.GetFunc().GetFnFlags()&ZEND_ACC_USER_ARG_INFO) != 0 {
		invoke.GetInternalFunction().SetFnFlags(invoke.GetInternalFunction().GetFnFlags() | ZEND_ACC_USER_ARG_INFO)
	}
	invoke.GetInternalFunction().SetHandler(zim_Closure___invoke)
	invoke.GetInternalFunction().SetModule(0)
	invoke.GetInternalFunction().SetScope(ZendCeClosure)
	invoke.GetInternalFunction().SetFunctionName(ZSTR_KNOWN(ZEND_STR_MAGIC_INVOKE))
	return invoke
}

/* }}} */

func ZendGetClosureMethodDef(obj *Zval) *ZendFunction {
	var closure *ZendClosure = (*ZendClosure)(Z_OBJ_P(obj))
	return &closure.func_
}

/* }}} */

func ZendGetClosureThisPtr(obj *Zval) *Zval {
	var closure *ZendClosure = (*ZendClosure)(Z_OBJ_P(obj))
	return &closure.this_ptr
}

/* }}} */

func ZendClosureGetMethod(object **ZendObject, method *ZendString, key *Zval) *ZendFunction {
	if ZendStringEqualsLiteralCi(method, ZEND_INVOKE_FUNC_NAME) {
		return ZendGetClosureInvokeMethod(*object)
	}
	return ZendStdGetMethod(object, method, key)
}

/* }}} */

func ZendClosureReadProperty(object *Zval, member *Zval, type_ int, cache_slot *any, rv *Zval) *Zval {
	ZEND_CLOSURE_PROPERTY_ERROR()
	return &(ExecutorGlobals.GetUninitializedZval())
}

/* }}} */

func ZendClosureWriteProperty(object *Zval, member *Zval, value *Zval, cache_slot *any) *Zval {
	ZEND_CLOSURE_PROPERTY_ERROR()
	return &(ExecutorGlobals.GetErrorZval())
}

/* }}} */

func ZendClosureGetPropertyPtrPtr(object *Zval, member *Zval, type_ int, cache_slot *any) *Zval {
	ZEND_CLOSURE_PROPERTY_ERROR()
	return nil
}

/* }}} */

func ZendClosureHasProperty(object *Zval, member *Zval, has_set_exists int, cache_slot *any) int {
	if has_set_exists != ZEND_PROPERTY_EXISTS {
		ZEND_CLOSURE_PROPERTY_ERROR()
	}
	return 0
}

/* }}} */

func ZendClosureUnsetProperty(object *Zval, member *Zval, cache_slot *any) {
	ZEND_CLOSURE_PROPERTY_ERROR()
}

/* }}} */

func ZendClosureFreeStorage(object *ZendObject) {
	var closure *ZendClosure = (*ZendClosure)(object)
	ZendObjectStdDtor(&closure.std)
	if closure.GetFunc().GetType() == ZEND_USER_FUNCTION {
		DestroyOpArray(&closure.func_.GetOpArray())
	} else if closure.GetOrigInternalHandler() == ZendClosureCallMagic {
		ZendStringRelease(closure.GetFunc().GetFunctionName())
	}
	if Z_TYPE(closure.GetThisPtr()) != IS_UNDEF {
		ZvalPtrDtor(&closure.this_ptr)
	}
}

/* }}} */

func ZendClosureNew(class_type *ZendClassEntry) *ZendObject {
	var closure *ZendClosure
	closure = Emalloc(b.SizeOf("zend_closure"))
	memset(closure, 0, b.SizeOf("zend_closure"))
	ZendObjectStdInit(&closure.std, class_type)
	closure.GetStd().SetHandlers(&ClosureHandlers)
	return (*ZendObject)(closure)
}

/* }}} */

func ZendClosureClone(zobject *Zval) *ZendObject {
	var closure *ZendClosure = (*ZendClosure)(Z_OBJ_P(zobject))
	var result Zval
	ZendCreateClosure(&result, &closure.func_, closure.GetFunc().GetScope(), closure.GetCalledScope(), &closure.this_ptr)
	return Z_OBJ(result)
}

/* }}} */

func ZendClosureGetClosure(obj *Zval, ce_ptr **ZendClassEntry, fptr_ptr **ZendFunction, obj_ptr **ZendObject) int {
	var closure *ZendClosure = (*ZendClosure)(Z_OBJ_P(obj))
	*fptr_ptr = &closure.func_
	*ce_ptr = closure.GetCalledScope()
	if Z_TYPE(closure.GetThisPtr()) != IS_UNDEF {
		*obj_ptr = Z_OBJ(closure.GetThisPtr())
	} else {
		*obj_ptr = nil
	}
	return SUCCESS
}

/* }}} */

func ZendClosureGetDebugInfo(object *Zval, is_temp *int) *HashTable {
	var closure *ZendClosure = (*ZendClosure)(Z_OBJ_P(object))
	var val Zval
	var arg_info *ZendArgInfo = closure.GetFunc().GetArgInfo()
	var debug_info *HashTable
	var zstr_args ZendBool = closure.GetFunc().GetType() == ZEND_USER_FUNCTION || (closure.GetFunc().GetFnFlags()&ZEND_ACC_USER_ARG_INFO) != 0
	*is_temp = 1
	debug_info = ZendNewArray(8)
	if closure.GetFunc().GetType() == ZEND_USER_FUNCTION && closure.GetFunc().GetOpArray().GetStaticVariables() != nil {
		var var_ *Zval
		var static_variables *HashTable = ZEND_MAP_PTR_GET(closure.func_.op_array.static_variables_ptr)
		ZVAL_ARR(&val, ZendArrayDup(static_variables))
		ZendHashUpdate(debug_info, ZSTR_KNOWN(ZEND_STR_STATIC), &val)
		for {
			var __ht *HashTable = Z_ARRVAL(val)
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
					continue
				}
				var_ = _z
				if Z_TYPE_P(var_) == IS_CONSTANT_AST {
					ZvalPtrDtor(var_)
					ZVAL_STRING(var_, "<constant ast>")
				}
			}
			break
		}
	}
	if Z_TYPE(closure.GetThisPtr()) != IS_UNDEF {
		Z_ADDREF(closure.GetThisPtr())
		ZendHashUpdate(debug_info, ZSTR_KNOWN(ZEND_STR_THIS), &closure.this_ptr)
	}
	if arg_info != nil && (closure.GetFunc().GetNumArgs() != 0 || (closure.GetFunc().GetFnFlags()&ZEND_ACC_VARIADIC) != 0) {
		var i uint32
		var num_args uint32
		var required uint32 = closure.GetFunc().GetRequiredNumArgs()
		ArrayInit(&val)
		num_args = closure.GetFunc().GetNumArgs()
		if (closure.GetFunc().GetFnFlags() & ZEND_ACC_VARIADIC) != 0 {
			num_args++
		}
		for i = 0; i < num_args; i++ {
			var name *ZendString
			var info Zval
			if arg_info.GetName() != nil {
				if zstr_args != 0 {
					name = ZendStrpprintf(0, "%s$%s", b.Cond(arg_info.GetPassByReference() != 0, "&", ""), ZSTR_VAL(arg_info.GetName()))
				} else {
					name = ZendStrpprintf(0, "%s$%s", b.Cond(arg_info.GetPassByReference() != 0, "&", ""), (*ZendInternalArgInfo)(arg_info).GetName())
				}
			} else {
				name = ZendStrpprintf(0, "%s$param%d", b.Cond(arg_info.GetPassByReference() != 0, "&", ""), i+1)
			}
			ZVAL_NEW_STR(&info, ZendStrpprintf(0, "%s", b.Cond(i >= required, "<optional>", "<required>")))
			ZendHashUpdate(Z_ARRVAL(val), name, &info)
			ZendStringReleaseEx(name, 0)
			arg_info++
		}
		ZendHashStrUpdate(debug_info, "parameter", b.SizeOf("\"parameter\"")-1, &val)
	}
	return debug_info
}

/* }}} */

func ZendClosureGetGc(obj *Zval, table **Zval, n *int) *HashTable {
	var closure *ZendClosure = (*ZendClosure)(Z_OBJ_P(obj))
	if Z_TYPE(closure.GetThisPtr()) != IS_NULL {
		*table = &closure.this_ptr
	} else {
		*table = nil
	}
	if Z_TYPE(closure.GetThisPtr()) != IS_NULL {
		*n = 1
	} else {
		*n = 0
	}
	if closure.GetFunc().GetType() == ZEND_USER_FUNCTION {
		return ZEND_MAP_PTR_GET(closure.func_.op_array.static_variables_ptr)
	} else {
		return nil
	}
}

/* }}} */

func zim_Closure___construct(execute_data *ZendExecuteData, return_value *Zval) {
	ZendThrowError(nil, "Instantiation of 'Closure' is not allowed")
}

/* }}} */

var ArginfoClosureBindto []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"newthis", 0, 0, 0}, {"newscope", 0, 0, 0}}
var ArginfoClosureBind []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"closure", 0, 0, 0}, {"newthis", 0, 0, 0}, {"newscope", 0, 0, 0}}
var ArginfoClosureCall []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"newthis", 0, 0, 0}, {"parameters", 0, 0, 1}}
var ArginfoClosureFromcallable []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"callable", 0, 0, 0}}
var ClosureFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"__construct",
		zim_Closure___construct,
		nil,
		uint32_t(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PRIVATE,
	},
	{
		"bind",
		zim_Closure_bind,
		ArginfoClosureBind,
		uint32_t(b.SizeOf("arginfo_closure_bind")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_STATIC,
	},
	{
		"bindTo",
		zim_Closure_bind,
		ArginfoClosureBindto,
		uint32_t(b.SizeOf("arginfo_closure_bindto")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"call",
		zim_Closure_call,
		ArginfoClosureCall,
		uint32_t(b.SizeOf("arginfo_closure_call")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"fromCallable",
		zim_Closure_fromCallable,
		ArginfoClosureFromcallable,
		uint32_t(b.SizeOf("arginfo_closure_fromcallable")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_STATIC,
	},
	{nil, nil, nil, 0, 0},
}

func ZendRegisterClosureCe() {
	var ce ZendClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Closure", b.SizeOf("\"Closure\"")-1, 1))
	ce.SetBuiltinFunctions(ClosureFunctions)
	ZendCeClosure = ZendRegisterInternalClass(&ce)
	ZendCeClosure.SetCeFlags(ZendCeClosure.GetCeFlags() | ZEND_ACC_FINAL)
	ZendCeClosure.create_object = ZendClosureNew
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

/* }}} */

func ZendClosureInternalHandler(execute_data *ZendExecuteData, return_value *Zval) {
	var closure *ZendClosure = (*ZendClosure)(ZEND_CLOSURE_OBJECT(EX(func_)))
	closure.GetOrigInternalHandler()(execute_data, return_value)
	OBJ_RELEASE((*ZendObject)(closure))
	EX(func_) = nil
}

/* }}} */

func ZendCreateClosure(res *Zval, func_ *ZendFunction, scope *ZendClassEntry, called_scope *ZendClassEntry, this_ptr *Zval) {
	var closure *ZendClosure
	ObjectInitEx(res, ZendCeClosure)
	closure = (*ZendClosure)(Z_OBJ_P(res))
	if scope == nil && this_ptr != nil && Z_TYPE_P(this_ptr) != IS_UNDEF {

		/* use dummy scope if we're binding an object without specifying a scope */

		scope = ZendCeClosure

		/* use dummy scope if we're binding an object without specifying a scope */

	}
	if func_.GetType() == ZEND_USER_FUNCTION {
		memcpy(&closure.func_, func_, b.SizeOf("zend_op_array"))
		closure.GetFunc().SetFnFlags(closure.GetFunc().GetFnFlags() | ZEND_ACC_CLOSURE)
		closure.GetFunc().SetFnFlags(closure.GetFunc().GetFnFlags() &^ ZEND_ACC_IMMUTABLE)
		if closure.GetFunc().GetOpArray().GetStaticVariables() != nil {
			closure.GetFunc().GetOpArray().SetStaticVariables(ZendArrayDup(closure.GetFunc().GetOpArray().GetStaticVariables()))
		}
		ZEND_MAP_PTR_INIT(closure.func_.op_array.static_variables_ptr, &closure.func_.GetOpArray().GetStaticVariables())

		/* Runtime cache is scope-dependent, so we cannot reuse it if the scope changed */

		if !(ZEND_MAP_PTR_GET(closure.func_.op_array.run_time_cache)) || func_.GetScope() != scope || (func_.GetFnFlags()&ZEND_ACC_HEAP_RT_CACHE) != 0 {
			var ptr any
			if !(ZEND_MAP_PTR_GET(func_.op_array.run_time_cache)) && (func_.GetFnFlags()&ZEND_ACC_CLOSURE) != 0 && (func_.GetScope() == scope || (func_.GetFnFlags()&ZEND_ACC_IMMUTABLE) == 0) {

				/* If a real closure is used for the first time, we create a shared runtime cache
				 * and remember which scope it is for. */

				if func_.GetScope() != scope {
					func_.SetScope(scope)
				}
				closure.GetFunc().GetOpArray().SetFnFlags(closure.GetFunc().GetOpArray().GetFnFlags() &^ ZEND_ACC_HEAP_RT_CACHE)
				ptr = ZendArenaAlloc(&(CompilerGlobals.GetArena()), func_.GetOpArray().GetCacheSize())
				ZEND_MAP_PTR_SET(func_.op_array.run_time_cache, ptr)
				ZEND_MAP_PTR_SET(closure.func_.op_array.run_time_cache, ptr)
			} else {

				/* Otherwise, we use a non-shared runtime cache */

				closure.GetFunc().GetOpArray().SetFnFlags(closure.GetFunc().GetOpArray().GetFnFlags() | ZEND_ACC_HEAP_RT_CACHE)
				ptr = Emalloc(b.SizeOf("void *") + func_.GetOpArray().GetCacheSize())
				ZEND_MAP_PTR_INIT(closure.func_.op_array.run_time_cache, ptr)
				ptr = (*byte)(ptr + b.SizeOf("void *"))
				ZEND_MAP_PTR_SET(closure.func_.op_array.run_time_cache, ptr)
			}
			memset(ptr, 0, func_.GetOpArray().GetCacheSize())
		}
		if closure.GetFunc().GetOpArray().GetRefcount() != nil {
			(*closure).func_.op_array.refcount++
		}
	} else {
		memcpy(&closure.func_, func_, b.SizeOf("zend_internal_function"))
		closure.GetFunc().SetFnFlags(closure.GetFunc().GetFnFlags() | ZEND_ACC_CLOSURE)

		/* wrap internal function handler to avoid memory leak */

		if UNEXPECTED(closure.GetFunc().GetInternalFunction().GetHandler() == ZendClosureInternalHandler) {

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
	ZVAL_UNDEF(&closure.this_ptr)

	/* Invariant:
	 * If the closure is unscoped or static, it has no bound object. */

	closure.GetFunc().SetScope(scope)
	closure.SetCalledScope(called_scope)
	if scope != nil {
		closure.GetFunc().SetFnFlags(closure.GetFunc().GetFnFlags() | ZEND_ACC_PUBLIC)
		if this_ptr != nil && Z_TYPE_P(this_ptr) == IS_OBJECT && (closure.GetFunc().GetFnFlags()&ZEND_ACC_STATIC) == 0 {
			Z_ADDREF_P(this_ptr)
			ZVAL_OBJ(&closure.this_ptr, Z_OBJ_P(this_ptr))
		}
	}
}

/* }}} */

func ZendCreateFakeClosure(res *Zval, func_ *ZendFunction, scope *ZendClassEntry, called_scope *ZendClassEntry, this_ptr *Zval) {
	var closure *ZendClosure
	ZendCreateClosure(res, func_, scope, called_scope, this_ptr)
	closure = (*ZendClosure)(Z_OBJ_P(res))
	closure.GetFunc().SetFnFlags(closure.GetFunc().GetFnFlags() | ZEND_ACC_FAKE_CLOSURE)
}

/* }}} */

func ZendClosureBindVar(closure_zv *Zval, var_name *ZendString, var_ *Zval) {
	var closure *ZendClosure = (*ZendClosure)(Z_OBJ_P(closure_zv))
	var static_variables *HashTable = ZEND_MAP_PTR_GET(closure.func_.op_array.static_variables_ptr)
	ZendHashUpdate(static_variables, var_name, var_)
}

/* }}} */

func ZendClosureBindVarEx(closure_zv *Zval, offset uint32, val *Zval) {
	var closure *ZendClosure = (*ZendClosure)(Z_OBJ_P(closure_zv))
	var static_variables *HashTable = ZEND_MAP_PTR_GET(closure.func_.op_array.static_variables_ptr)
	var var_ *Zval = (*Zval)((*byte)(static_variables.GetArData() + offset))
	ZvalPtrDtor(var_)
	ZVAL_COPY_VALUE(var_, val)
}

/* }}} */
