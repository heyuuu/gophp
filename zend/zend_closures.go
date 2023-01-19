// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
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

// #define ZEND_CLOSURE_OBJECT(op_array) ( ( zend_object * ) ( ( char * ) ( op_array ) - sizeof ( zend_object ) ) )

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

// #define ZEND_CLOSURE_PRINT_NAME       "Closure object"

// #define ZEND_CLOSURE_PROPERTY_ERROR() zend_throw_error ( NULL , "Closure object cannot have properties" )

// @type ZendClosure struct

/* non-static since it needs to be referenced */

var ClosureHandlers ZendObjectHandlers

func zim_Closure___invoke(execute_data *ZendExecuteData, return_value *Zval) {
	var func_ *ZendFunction = execute_data.GetFunc()
	var arguments *Zval = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(1)-1))
	if _callUserFunctionEx(nil, &(execute_data.GetThis()), return_value, execute_data.GetThis().GetNumArgs(), arguments, 1) == FAILURE {
		return_value.SetTypeInfo(2)
	}

	/* destruct the function also, then - we have allocated it in get_method */

	ZendStringReleaseEx(func_.GetInternalFunction().GetFunctionName(), 0)
	_efree(func_)
}

/* }}} */

func ZendValidClosureBinding(closure *ZendClosure, newthis *Zval, scope *ZendClassEntry) ZendBool {
	var func_ *ZendFunction = &closure.func_
	var is_fake_closure ZendBool = (func_.GetFnFlags() & 1 << 21) != 0
	if newthis != nil {
		if (func_.GetFnFlags() & 1 << 4) != 0 {
			ZendError(1<<1, "Cannot bind an instance to a static closure")
			return 0
		}
		if is_fake_closure != 0 && func_.GetScope() != nil && InstanceofFunction(newthis.GetValue().GetObj().GetCe(), func_.GetScope()) == 0 {

			/* Binding incompatible $this to an internal method is not supported. */

			ZendError(1<<1, "Cannot bind method %s::%s() to object of class %s", func_.GetScope().GetName().GetVal(), func_.GetFunctionName().GetVal(), newthis.GetValue().GetObj().GetCe().GetName().GetVal())
			return 0
		}
	} else if is_fake_closure != 0 && func_.GetScope() != nil && (func_.GetFnFlags()&1<<4) == 0 {
		if func_.GetType() == 1 {
			ZendError(1<<1, "Cannot unbind $this of internal method")
			return 0
		} else {
			ZendError(1<<13, "Unbinding $this of a method is deprecated")
		}
	} else if is_fake_closure == 0 && closure.GetThisPtr().GetType() != 0 && (func_.GetFnFlags()&1<<30) != 0 {

		// TODO: Only deprecate if it had $this *originally*?

		ZendError(1<<13, "Unbinding $this of closure is deprecated")

		// TODO: Only deprecate if it had $this *originally*?

	}
	if scope != nil && scope != func_.GetScope() && scope.GetType() == 1 {

		/* rebinding to internal class is not allowed */

		ZendError(1<<1, "Cannot bind closure to scope of internal class %s", scope.GetName().GetVal())
		return 0
	}
	if is_fake_closure != 0 && scope != func_.GetScope() {
		if func_.GetScope() == nil {
			ZendError(1<<1, "Cannot rebind scope of closure created from function")
		} else {
			ZendError(1<<1, "Cannot rebind scope of closure created from method")
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
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "o*", &newthis, &fci.params, &fci.param_count) == FAILURE {
		return
	}
	closure = (*ZendClosure)(&(execute_data.GetThis()).GetValue().GetObj())
	newobj = newthis.GetValue().GetObj()
	if ZendValidClosureBinding(closure, newthis, newthis.GetValue().GetObj().GetCe()) == 0 {
		return
	}
	if (closure.GetFunc().GetFnFlags() & 1 << 24) != 0 {
		var new_closure Zval
		ZendCreateClosure(&new_closure, &closure.func_, newthis.GetValue().GetObj().GetCe(), closure.GetCalledScope(), newthis)
		closure = (*ZendClosure)(new_closure.GetValue().GetObj())
		fci_cache.SetFunctionHandler(&closure.func_)
	} else {
		memcpy(&my_function, &closure.func_, g.CondF(closure.GetFunc().GetType() == 2, func() __auto__ { return g.SizeOf("zend_op_array") }, func() __auto__ { return g.SizeOf("zend_internal_function") }))
		my_function.SetFnFlags(my_function.GetFnFlags() &^ (1 << 20))

		/* use scope of passed object */

		my_function.SetScope(newthis.GetValue().GetObj().GetCe())
		if closure.GetFunc().GetType() == 1 {
			my_function.GetInternalFunction().SetHandler(closure.GetOrigInternalHandler())
		}
		fci_cache.SetFunctionHandler(&my_function)

		/* Runtime cache relies on bound scope to be immutable, hence we need a separate rt cache in case scope changed */

		if (my_function.GetType()&1) == 0 && (closure.GetFunc().GetScope() != newthis.GetValue().GetObj().GetCe() || (closure.GetFunc().GetFnFlags()&1<<22) != 0) {
			var ptr any
			my_function.GetOpArray().SetFnFlags(my_function.GetOpArray().GetFnFlags() | 1<<22)
			ptr = _emalloc(g.SizeOf("void *") + my_function.GetOpArray().GetCacheSize())
			my_function.GetOpArray().SetRunTimeCachePtr(ptr)
			ptr = (*byte)(ptr + g.SizeOf("void *"))
			if (uintPtr(my_function.GetOpArray().GetRunTimeCachePtr()) & 1) != 0 {
				*((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(my_function.GetOpArray().GetRunTimeCachePtr()-1)))) = ptr
			} else {
				*(my_function.GetOpArray().GetRunTimeCachePtr()) = ptr
			}
			memset(ptr, 0, my_function.GetOpArray().GetCacheSize())
		}

		/* Runtime cache relies on bound scope to be immutable, hence we need a separate rt cache in case scope changed */

	}
	fci_cache.SetCalledScope(newobj.GetCe())
	fci.SetObject(newobj)
	fci_cache.SetObject(fci.GetObject())
	fci.SetSize(g.SizeOf("fci"))
	var __z *Zval = &fci.function_name
	__z.GetValue().SetObj(&closure.std)
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	fci.SetRetval(&closure_result)
	fci.SetNoSeparation(1)
	if ZendCallFunction(&fci, &fci_cache) == SUCCESS && closure_result.GetType() != 0 {
		if closure_result.GetType() == 10 {
			ZendUnwrapReference(&closure_result)
		}
		var _z1 *Zval = return_value
		var _z2 *Zval = &closure_result
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	}
	if (fci_cache.GetFunctionHandler().GetFnFlags() & 1 << 24) != 0 {

		/* copied upon generator creation */

		ZendGcDelref(&(&closure.std).GetGc())

		/* copied upon generator creation */

	} else if (my_function.GetType()&1) == 0 && (fci_cache.GetFunctionHandler().GetFnFlags()&1<<22) != 0 {
		_efree(my_function.GetOpArray().GetRunTimeCachePtr())
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
	if ZendParseMethodParameters(execute_data.GetThis().GetNumArgs(), g.CondF1(&(execute_data.GetThis()).GetType() == 8, func() *Zval { return &(execute_data.GetThis()) }, nil), "Oo!|z", &zclosure, ZendCeClosure, &newthis, &scope_arg) == FAILURE {
		return
	}
	closure = (*ZendClosure)(zclosure.GetValue().GetObj())
	if scope_arg != nil {
		if scope_arg.GetType() == 8 {
			ce = scope_arg.GetValue().GetObj().GetCe()
		} else if scope_arg.GetType() == 1 {
			ce = nil
		} else {
			var tmp_class_name *ZendString
			var class_name *ZendString = ZvalGetTmpString(scope_arg, &tmp_class_name)
			if class_name.GetLen() == g.SizeOf("\"static\"")-1 && !(memcmp(class_name.GetVal(), "static", g.SizeOf("\"static\"")-1)) {
				ce = closure.GetFunc().GetScope()
			} else if g.Assign(&ce, ZendLookupClass(class_name)) == nil {
				ZendError(1<<1, "Class '%s' not found", class_name.GetVal())
				ZendTmpStringRelease(tmp_class_name)
				return_value.SetTypeInfo(1)
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
		called_scope = newthis.GetValue().GetObj().GetCe()
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
	memset(&fci, 0, g.SizeOf("zend_fcall_info"))
	memset(&fcc, 0, g.SizeOf("zend_fcall_info_cache"))
	fci.SetSize(g.SizeOf("zend_fcall_info"))
	fci.SetRetval(return_value)
	if (execute_data.GetFunc().GetInternalFunction().GetFnFlags() & 1 << 4) != 0 {
		fcc.SetFunctionHandler(execute_data.GetFunc().GetInternalFunction().GetScope().GetCallstatic())
	} else {
		fcc.SetFunctionHandler(execute_data.GetFunc().GetInternalFunction().GetScope().GetCall())
	}
	fci.SetParams(params)
	fci.SetParamCount(2)
	var __z *Zval = &fci.params[0]
	var __s *ZendString = execute_data.GetFunc().GetFunctionName()
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	if execute_data.GetThis().GetNumArgs() != 0 {
		var __arr *ZendArray = _zendNewArray(execute_data.GetThis().GetNumArgs())
		var __z *Zval = &fci.params[1]
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		ZendCopyParametersArray(execute_data.GetThis().GetNumArgs(), &fci.params[1])
	} else {
		var __z *Zval = &fci.params[1]
		__z.GetValue().SetArr((*ZendArray)(&ZendEmptyArray))
		__z.SetTypeInfo(7)
	}
	fci.SetObject(&(execute_data.GetThis()).GetValue().GetObj())
	fcc.SetObject(fci.GetObject())
	fcc.SetCalledScope(ZendGetCalledScope(EG.GetCurrentExecuteData()))
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
	if (mptr.GetFnFlags() & 1 << 18) != 0 {

		/* For Closure::fromCallable([$closure, "__invoke"]) return $closure. */

		if fcc.GetObject() != nil && fcc.GetObject().GetCe() == ZendCeClosure && (mptr.GetFunctionName().GetLen() == g.SizeOf("\"__invoke\"")-1 && !(memcmp(mptr.GetFunctionName().GetVal(), "__invoke", g.SizeOf("\"__invoke\"")-1))) {
			var __z *Zval = return_value
			__z.GetValue().SetObj(fcc.GetObject())
			__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
			ZendGcAddref(&(fcc.GetObject()).gc)
			if mptr == &EG.trampoline {
				EG.GetTrampoline().SetFunctionName(nil)
			} else {
				_efree(mptr)
			}
			return SUCCESS
		}
		if mptr.GetScope() == nil {
			return FAILURE
		}
		if (mptr.GetFnFlags() & 1 << 4) != 0 {
			if mptr.GetScope().GetCallstatic() == nil {
				return FAILURE
			}
		} else {
			if mptr.GetScope().GetCall() == nil {
				return FAILURE
			}
		}
		memset(&call, 0, g.SizeOf("zend_internal_function"))
		call.SetType(1)
		call.SetFnFlags(mptr.GetFnFlags() & 1 << 4)
		call.SetHandler(ZendClosureCallMagic)
		call.SetFunctionName(mptr.GetFunctionName())
		call.SetScope(mptr.GetScope())
		if mptr == &EG.trampoline {
			EG.GetTrampoline().SetFunctionName(nil)
		} else {
			_efree(mptr)
		}
		mptr = (*ZendFunction)(&call)
	}
	if fcc.GetObject() != nil {
		var __z *Zval = &instance
		__z.GetValue().SetObj(fcc.GetObject())
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
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
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "z", &callable) == FAILURE {
		return
	}
	if callable.GetType() == 8 && InstanceofFunction(callable.GetValue().GetObj().GetCe(), ZendCeClosure) != 0 {

		/* It's already a closure */

		var __z *Zval = return_value
		var __zv *Zval = callable
		if __zv.GetType() != 10 {
			var _z1 *Zval = __z
			var _z2 *Zval = __zv
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		} else {
			var _z1 *Zval = __z
			var _z2 *Zval = &(*__zv).value.GetRef().GetVal()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}

		}
		return
	}

	/* create closure as if it were called from parent scope */

	EG.SetCurrentExecuteData(execute_data.GetPrevExecuteData())
	success = ZendCreateClosureFromCallable(return_value, callable, &error)
	EG.SetCurrentExecuteData(execute_data)
	if success == FAILURE || error != nil {
		if error != nil {
			ZendTypeError("Failed to create closure from callable: %s", error)
			_efree(error)
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

func ZendClosureCompareObjects(o1 *Zval, o2 *Zval) int {
	return o1.GetValue().GetObj() != o2.GetValue().GetObj()
}

/* }}} */

func ZendGetClosureInvokeMethod(object *ZendObject) *ZendFunction {
	var closure *ZendClosure = (*ZendClosure)(object)
	var invoke *ZendFunction = (*ZendFunction)(_emalloc(g.SizeOf("zend_function")))
	var keep_flags uint32 = 1<<12 | 1<<14 | 1<<13
	invoke.SetCommon(closure.GetFunc().GetCommon())

	/* We return ZEND_INTERNAL_FUNCTION, but arg_info representation is the
	 * same as for ZEND_USER_FUNCTION (uses zend_string* instead of char*).
	 * This is not a problem, because ZEND_ACC_HAS_TYPE_HINTS is never set,
	 * and we won't check arguments on internal function. We also set
	 * ZEND_ACC_USER_ARG_INFO flag to prevent invalid usage by Reflection */

	invoke.SetType(1)
	invoke.GetInternalFunction().SetFnFlags(1<<0 | 1<<18 | closure.GetFunc().GetFnFlags()&keep_flags)
	if closure.GetFunc().GetType() != 1 || (closure.GetFunc().GetFnFlags()&1<<22) != 0 {
		invoke.GetInternalFunction().SetFnFlags(invoke.GetInternalFunction().GetFnFlags() | 1<<22)
	}
	invoke.GetInternalFunction().SetHandler(zim_Closure___invoke)
	invoke.GetInternalFunction().SetModule(0)
	invoke.GetInternalFunction().SetScope(ZendCeClosure)
	invoke.GetInternalFunction().SetFunctionName(ZendKnownStrings[ZEND_STR_MAGIC_INVOKE])
	return invoke
}

/* }}} */

func ZendGetClosureMethodDef(obj *Zval) *ZendFunction {
	var closure *ZendClosure = (*ZendClosure)(obj.GetValue().GetObj())
	return &closure.func_
}

/* }}} */

func ZendGetClosureThisPtr(obj *Zval) *Zval {
	var closure *ZendClosure = (*ZendClosure)(obj.GetValue().GetObj())
	return &closure.this_ptr
}

/* }}} */

func ZendClosureGetMethod(object **ZendObject, method *ZendString, key *Zval) *ZendFunction {
	if method.GetLen() == g.SizeOf("ZEND_INVOKE_FUNC_NAME")-1 && ZendBinaryStrcasecmp(method.GetVal(), method.GetLen(), "__invoke", g.SizeOf("ZEND_INVOKE_FUNC_NAME")-1) == 0 {
		return ZendGetClosureInvokeMethod(*object)
	}
	return ZendStdGetMethod(object, method, key)
}

/* }}} */

func ZendClosureReadProperty(object *Zval, member *Zval, type_ int, cache_slot *any, rv *Zval) *Zval {
	ZendThrowError(nil, "Closure object cannot have properties")
	return &EG.uninitialized_zval
}

/* }}} */

func ZendClosureWriteProperty(object *Zval, member *Zval, value *Zval, cache_slot *any) *Zval {
	ZendThrowError(nil, "Closure object cannot have properties")
	return &EG.error_zval
}

/* }}} */

func ZendClosureGetPropertyPtrPtr(object *Zval, member *Zval, type_ int, cache_slot *any) *Zval {
	ZendThrowError(nil, "Closure object cannot have properties")
	return nil
}

/* }}} */

func ZendClosureHasProperty(object *Zval, member *Zval, has_set_exists int, cache_slot *any) int {
	if has_set_exists != 0x2 {
		ZendThrowError(nil, "Closure object cannot have properties")
	}
	return 0
}

/* }}} */

func ZendClosureUnsetProperty(object *Zval, member *Zval, cache_slot *any) {
	ZendThrowError(nil, "Closure object cannot have properties")
}

/* }}} */

func ZendClosureFreeStorage(object *ZendObject) {
	var closure *ZendClosure = (*ZendClosure)(object)
	ZendObjectStdDtor(&closure.std)
	if closure.GetFunc().GetType() == 2 {
		DestroyOpArray(&closure.func_.GetOpArray())
	} else if closure.GetOrigInternalHandler() == ZendClosureCallMagic {
		ZendStringRelease(closure.GetFunc().GetFunctionName())
	}
	if closure.GetThisPtr().GetType() != 0 {
		ZvalPtrDtor(&closure.this_ptr)
	}
}

/* }}} */

func ZendClosureNew(class_type *ZendClassEntry) *ZendObject {
	var closure *ZendClosure
	closure = _emalloc(g.SizeOf("zend_closure"))
	memset(closure, 0, g.SizeOf("zend_closure"))
	ZendObjectStdInit(&closure.std, class_type)
	closure.GetStd().SetHandlers(&ClosureHandlers)
	return (*ZendObject)(closure)
}

/* }}} */

func ZendClosureClone(zobject *Zval) *ZendObject {
	var closure *ZendClosure = (*ZendClosure)(zobject.GetValue().GetObj())
	var result Zval
	ZendCreateClosure(&result, &closure.func_, closure.GetFunc().GetScope(), closure.GetCalledScope(), &closure.this_ptr)
	return result.GetValue().GetObj()
}

/* }}} */

func ZendClosureGetClosure(obj *Zval, ce_ptr **ZendClassEntry, fptr_ptr **ZendFunction, obj_ptr **ZendObject) int {
	var closure *ZendClosure = (*ZendClosure)(obj.GetValue().GetObj())
	*fptr_ptr = &closure.func_
	*ce_ptr = closure.GetCalledScope()
	if closure.GetThisPtr().GetType() != 0 {
		*obj_ptr = closure.GetThisPtr().GetValue().GetObj()
	} else {
		*obj_ptr = nil
	}
	return SUCCESS
}

/* }}} */

func ZendClosureGetDebugInfo(object *Zval, is_temp *int) *HashTable {
	var closure *ZendClosure = (*ZendClosure)(object.GetValue().GetObj())
	var val Zval
	var arg_info *ZendArgInfo = closure.GetFunc().GetArgInfo()
	var debug_info *HashTable
	var zstr_args ZendBool = closure.GetFunc().GetType() == 2 || (closure.GetFunc().GetFnFlags()&1<<22) != 0
	*is_temp = 1
	debug_info = _zendNewArray(8)
	if closure.GetFunc().GetType() == 2 && closure.GetFunc().GetOpArray().GetStaticVariables() != nil {
		var var_ *Zval
		var static_variables *HashTable = g.CondF((uintPtr(closure.GetFunc().GetOpArray().GetStaticVariablesPtrPtr())&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(closure.GetFunc().GetOpArray().GetStaticVariablesPtrPtr()-1))))
		}, func() any { return any(*(closure.GetFunc().GetOpArray().GetStaticVariablesPtrPtr())) })
		var __arr *ZendArray = ZendArrayDup(static_variables)
		var __z *Zval = &val
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		ZendHashUpdate(debug_info, ZendKnownStrings[ZEND_STR_STATIC], &val)
		for {
			var __ht *HashTable = val.GetValue().GetArr()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				var_ = _z
				if var_.GetType() == 11 {
					ZvalPtrDtor(var_)
					var _s *byte = "<constant ast>"
					var __z *Zval = var_
					var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
					__z.GetValue().SetStr(__s)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			break
		}
	}
	if closure.GetThisPtr().GetType() != 0 {
		ZvalAddrefP(&(closure.GetThisPtr()))
		ZendHashUpdate(debug_info, ZendKnownStrings[ZEND_STR_THIS], &closure.this_ptr)
	}
	if arg_info != nil && (closure.GetFunc().GetNumArgs() != 0 || (closure.GetFunc().GetFnFlags()&1<<14) != 0) {
		var i uint32
		var num_args uint32
		var required uint32 = closure.GetFunc().GetRequiredNumArgs()
		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = &val
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		num_args = closure.GetFunc().GetNumArgs()
		if (closure.GetFunc().GetFnFlags() & 1 << 14) != 0 {
			num_args++
		}
		for i = 0; i < num_args; i++ {
			var name *ZendString
			var info Zval
			if arg_info.GetName() != nil {
				if zstr_args != 0 {
					name = ZendStrpprintf(0, "%s$%s", g.Cond(arg_info.GetPassByReference() != 0, "&", ""), arg_info.GetName().GetVal())
				} else {
					name = ZendStrpprintf(0, "%s$%s", g.Cond(arg_info.GetPassByReference() != 0, "&", ""), (*ZendInternalArgInfo)(arg_info).GetName())
				}
			} else {
				name = ZendStrpprintf(0, "%s$param%d", g.Cond(arg_info.GetPassByReference() != 0, "&", ""), i+1)
			}
			var __z *Zval = &info
			var __s *ZendString = ZendStrpprintf(0, "%s", g.Cond(i >= required, "<optional>", "<required>"))
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			ZendHashUpdate(val.GetValue().GetArr(), name, &info)
			ZendStringReleaseEx(name, 0)
			arg_info++
		}
		ZendHashStrUpdate(debug_info, "parameter", g.SizeOf("\"parameter\"")-1, &val)
	}
	return debug_info
}

/* }}} */

func ZendClosureGetGc(obj *Zval, table **Zval, n *int) *HashTable {
	var closure *ZendClosure = (*ZendClosure)(obj.GetValue().GetObj())
	if closure.GetThisPtr().GetType() != 1 {
		*table = &closure.this_ptr
	} else {
		*table = nil
	}
	if closure.GetThisPtr().GetType() != 1 {
		*n = 1
	} else {
		*n = 0
	}
	if closure.GetFunc().GetType() == 2 {
		if (uintPtr(closure.GetFunc().GetOpArray().GetStaticVariablesPtrPtr()) & 1) != 0 {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(closure.GetFunc().GetOpArray().GetStaticVariablesPtrPtr()-1))))
		} else {
			return any(*(closure.GetFunc().GetOpArray().GetStaticVariablesPtrPtr()))
		}
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
var ClosureFunctions []ZendFunctionEntry = []ZendFunctionEntry{{"__construct", zim_Closure___construct, nil, uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1 << 2}, {"bind", zim_Closure_bind, ArginfoClosureBind, uint32(g.SizeOf("arginfo_closure_bind")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<4}, {"bindTo", zim_Closure_bind, ArginfoClosureBindto, uint32(g.SizeOf("arginfo_closure_bindto")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1 << 0}, {"call", zim_Closure_call, ArginfoClosureCall, uint32(g.SizeOf("arginfo_closure_call")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1 << 0}, {"fromCallable", zim_Closure_fromCallable, ArginfoClosureFromcallable, uint32(g.SizeOf("arginfo_closure_fromcallable")/g.SizeOf("struct _zend_internal_arg_info") - 1), 1<<0 | 1<<4}, {nil, nil, nil, 0, 0}}

func ZendRegisterClosureCe() {
	var ce ZendClassEntry
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Closure", g.SizeOf("\"Closure\"")-1, 1))
	ce.SetBuiltinFunctions(ClosureFunctions)
	ZendCeClosure = ZendRegisterInternalClass(&ce)
	ZendCeClosure.SetCeFlags(ZendCeClosure.GetCeFlags() | 1<<5)
	ZendCeClosure.create_object = ZendClosureNew
	ZendCeClosure.SetSerialize(ZendClassSerializeDeny)
	ZendCeClosure.SetUnserialize(ZendClassUnserializeDeny)
	memcpy(&ClosureHandlers, &StdObjectHandlers, g.SizeOf("zend_object_handlers"))
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
	var closure *ZendClosure = (*ZendClosure)((*ZendObject)((*byte)(execute_data.GetFunc() - g.SizeOf("zend_object"))))
	closure.GetOrigInternalHandler()(execute_data, return_value)
	ZendObjectRelease((*ZendObject)(closure))
	execute_data.SetFunc(nil)
}

/* }}} */

func ZendCreateClosure(res *Zval, func_ *ZendFunction, scope *ZendClassEntry, called_scope *ZendClassEntry, this_ptr *Zval) {
	var closure *ZendClosure
	ObjectInitEx(res, ZendCeClosure)
	closure = (*ZendClosure)(res.GetValue().GetObj())
	if scope == nil && this_ptr != nil && this_ptr.GetType() != 0 {

		/* use dummy scope if we're binding an object without specifying a scope */

		scope = ZendCeClosure

		/* use dummy scope if we're binding an object without specifying a scope */

	}
	if func_.GetType() == 2 {
		memcpy(&closure.func_, func_, g.SizeOf("zend_op_array"))
		closure.GetFunc().SetFnFlags(closure.GetFunc().GetFnFlags() | 1<<20)
		closure.GetFunc().SetFnFlags(closure.GetFunc().GetFnFlags() &^ (1 << 7))
		if closure.GetFunc().GetOpArray().GetStaticVariables() != nil {
			closure.GetFunc().GetOpArray().SetStaticVariables(ZendArrayDup(closure.GetFunc().GetOpArray().GetStaticVariables()))
		}
		closure.GetFunc().GetOpArray().SetStaticVariablesPtrPtr(&closure.func_.GetOpArray().GetStaticVariables())

		/* Runtime cache is scope-dependent, so we cannot reuse it if the scope changed */

		if !(g.CondF((uintPtr(closure.GetFunc().GetOpArray().GetRunTimeCachePtr())&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(closure.GetFunc().GetOpArray().GetRunTimeCachePtr()-1))))
		}, func() any { return any(*(closure.GetFunc().GetOpArray().GetRunTimeCachePtr())) })) || func_.GetScope() != scope || (func_.GetFnFlags()&1<<22) != 0 {
			var ptr any
			if !(g.CondF((uintPtr(func_.GetOpArray().GetRunTimeCachePtr())&1) != 0, func() any {
				return *((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(func_.GetOpArray().GetRunTimeCachePtr()-1))))
			}, func() any { return any(*(func_.GetOpArray().GetRunTimeCachePtr())) })) && (func_.GetFnFlags()&1<<20) != 0 && (func_.GetScope() == scope || (func_.GetFnFlags()&1<<7) == 0) {

				/* If a real closure is used for the first time, we create a shared runtime cache
				 * and remember which scope it is for. */

				if func_.GetScope() != scope {
					func_.SetScope(scope)
				}
				closure.GetFunc().GetOpArray().SetFnFlags(closure.GetFunc().GetOpArray().GetFnFlags() &^ (1 << 22))
				ptr = ZendArenaAlloc(&CG.arena, func_.GetOpArray().GetCacheSize())
				if (uintPtr(func_.GetOpArray().GetRunTimeCachePtr()) & 1) != 0 {
					*((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(func_.GetOpArray().GetRunTimeCachePtr()-1)))) = ptr
				} else {
					*(func_.GetOpArray().GetRunTimeCachePtr()) = ptr
				}
				if (uintPtr(closure.GetFunc().GetOpArray().GetRunTimeCachePtr()) & 1) != 0 {
					*((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(closure.GetFunc().GetOpArray().GetRunTimeCachePtr()-1)))) = ptr
				} else {
					*(closure.GetFunc().GetOpArray().GetRunTimeCachePtr()) = ptr
				}
			} else {

				/* Otherwise, we use a non-shared runtime cache */

				closure.GetFunc().GetOpArray().SetFnFlags(closure.GetFunc().GetOpArray().GetFnFlags() | 1<<22)
				ptr = _emalloc(g.SizeOf("void *") + func_.GetOpArray().GetCacheSize())
				closure.GetFunc().GetOpArray().SetRunTimeCachePtr(ptr)
				ptr = (*byte)(ptr + g.SizeOf("void *"))
				if (uintPtr(closure.GetFunc().GetOpArray().GetRunTimeCachePtr()) & 1) != 0 {
					*((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(closure.GetFunc().GetOpArray().GetRunTimeCachePtr()-1)))) = ptr
				} else {
					*(closure.GetFunc().GetOpArray().GetRunTimeCachePtr()) = ptr
				}
			}
			memset(ptr, 0, func_.GetOpArray().GetCacheSize())
		}
		if closure.GetFunc().GetOpArray().GetRefcount() != nil {
			(*closure).func_.op_array.refcount++
		}
	} else {
		memcpy(&closure.func_, func_, g.SizeOf("zend_internal_function"))
		closure.GetFunc().SetFnFlags(closure.GetFunc().GetFnFlags() | 1<<20)

		/* wrap internal function handler to avoid memory leak */

		if closure.GetFunc().GetInternalFunction().GetHandler() == ZendClosureInternalHandler {

			/* avoid infinity recursion, by taking handler from nested closure */

			var nested *ZendClosure = (*ZendClosure)((*byte)(func_ - zend_long((*byte)(&((*ZendClosure)(nil).GetFunc()))-(*byte)(nil))))
			assert(nested.GetStd().GetCe() == ZendCeClosure)
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
	&closure.this_ptr.u1.type_info = 0

	/* Invariant:
	 * If the closure is unscoped or static, it has no bound object. */

	closure.GetFunc().SetScope(scope)
	closure.SetCalledScope(called_scope)
	if scope != nil {
		closure.GetFunc().SetFnFlags(closure.GetFunc().GetFnFlags() | 1<<0)
		if this_ptr != nil && this_ptr.GetType() == 8 && (closure.GetFunc().GetFnFlags()&1<<4) == 0 {
			ZvalAddrefP(this_ptr)
			var __z *Zval = &closure.this_ptr
			__z.GetValue().SetObj(this_ptr.GetValue().GetObj())
			__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
		}
	}
}

/* }}} */

func ZendCreateFakeClosure(res *Zval, func_ *ZendFunction, scope *ZendClassEntry, called_scope *ZendClassEntry, this_ptr *Zval) {
	var closure *ZendClosure
	ZendCreateClosure(res, func_, scope, called_scope, this_ptr)
	closure = (*ZendClosure)(res.GetValue().GetObj())
	closure.GetFunc().SetFnFlags(closure.GetFunc().GetFnFlags() | 1<<21)
}

/* }}} */

func ZendClosureBindVar(closure_zv *Zval, var_name *ZendString, var_ *Zval) {
	var closure *ZendClosure = (*ZendClosure)(closure_zv.GetValue().GetObj())
	var static_variables *HashTable = g.CondF((uintPtr(closure.GetFunc().GetOpArray().GetStaticVariablesPtrPtr())&1) != 0, func() any {
		return *((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(closure.GetFunc().GetOpArray().GetStaticVariablesPtrPtr()-1))))
	}, func() any { return any(*(closure.GetFunc().GetOpArray().GetStaticVariablesPtrPtr())) })
	ZendHashUpdate(static_variables, var_name, var_)
}

/* }}} */

func ZendClosureBindVarEx(closure_zv *Zval, offset uint32, val *Zval) {
	var closure *ZendClosure = (*ZendClosure)(closure_zv.GetValue().GetObj())
	var static_variables *HashTable = g.CondF((uintPtr(closure.GetFunc().GetOpArray().GetStaticVariablesPtrPtr())&1) != 0, func() any {
		return *((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(closure.GetFunc().GetOpArray().GetStaticVariablesPtrPtr()-1))))
	}, func() any { return any(*(closure.GetFunc().GetOpArray().GetStaticVariablesPtrPtr())) })
	var var_ *Zval = (*Zval)((*byte)(static_variables.GetArData() + offset))
	ZvalPtrDtor(var_)
	var _z1 *Zval = var_
	var _z2 *Zval = val
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
}

/* }}} */
