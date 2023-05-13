package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_INIT_USER_CALL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var fcc types.ZendFcallInfoCache
	var error *byte = nil
	var func_ types.IFunction
	var object_or_called_scope any
	var call *ZendExecuteData
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	function_name = opline.Const2()
	if ZendIsCallableEx(function_name, nil, 0, nil, &fcc, &error) != 0 {
		func_ = fcc.GetFunctionHandler()
		if error != nil {
			Efree(error)

			/* This is the only soft error is_callable() can generate */

			ZendNonStaticMethodCall(func_)
			if EG__().GetException() != nil {
				return 0
			}
		}
		object_or_called_scope = fcc.GetCalledScope()
		if func_.IsClosure() {

			/* Delay closure destruction until its invocation */

			//ZEND_CLOSURE_OBJECT(func_).AddRefcount()
			call_info |= ZEND_CALL_CLOSURE
			if func_.IsFakeClosure() {
				call_info |= ZEND_CALL_FAKE_CLOSURE
			}
			if fcc.GetObject() != nil {
				object_or_called_scope = fcc.GetObject()
				call_info |= ZEND_CALL_HAS_THIS
			}
		} else if fcc.GetObject() != nil {
			//fcc.GetObject().AddRefcount()
			object_or_called_scope = fcc.GetObject()
			call_info |= ZEND_CALL_RELEASE_THIS | ZEND_CALL_HAS_THIS
		}
		if func_.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(func_.GetOpArray())) {
			InitFuncRunTimeCache(func_.GetOpArray())
		}
	} else {
		faults.InternalTypeError(executeData.IsCallUseStrictTypes(), "%s() expects parameter 1 to be a valid callback, %s", opline.Const1().String().GetVal(), error)
		Efree(error)
		if EG__().GetException() != nil {
			return 0
		}
		func_ = (types.IFunction)(&ZendPassFunction)
		object_or_called_scope = nil
	}
	call = ZendVmStackPushCallFrame(call_info, func_, opline.GetExtendedValue(), object_or_called_scope)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_USER_CALL_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var function_name *types.Zval
	var fcc types.ZendFcallInfoCache
	var error *byte = nil
	var func_ types.IFunction
	var object_or_called_scope any
	var call *ZendExecuteData
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	function_name = opline.Op2()
	if ZendIsCallableEx(function_name, nil, 0, nil, &fcc, &error) != 0 {
		func_ = fcc.GetFunctionHandler()
		if error != nil {
			Efree(error)

			/* This is the only soft error is_callable() can generate */

			ZendNonStaticMethodCall(func_)
			if EG__().GetException() != nil {
				// ZvalPtrDtorNogc(free_op2)
				return 0
			}
		}
		object_or_called_scope = fcc.GetCalledScope()
		if func_.IsClosure() {

			/* Delay closure destruction until its invocation */

			//ZEND_CLOSURE_OBJECT(func_).AddRefcount()
			call_info |= ZEND_CALL_CLOSURE
			if func_.IsFakeClosure() {
				call_info |= ZEND_CALL_FAKE_CLOSURE
			}
			if fcc.GetObject() != nil {
				object_or_called_scope = fcc.GetObject()
				call_info |= ZEND_CALL_HAS_THIS
			}
		} else if fcc.GetObject() != nil {
			//fcc.GetObject().AddRefcount()
			object_or_called_scope = fcc.GetObject()
			call_info |= ZEND_CALL_RELEASE_THIS | ZEND_CALL_HAS_THIS
		}
		// ZvalPtrDtorNogc(free_op2)
		if EG__().GetException() != nil {
			if (call_info & ZEND_CALL_CLOSURE) != 0 {
				//ZendObjectRelease(ZEND_CLOSURE_OBJECT(func_))
			} else if (call_info & ZEND_CALL_RELEASE_THIS) != 0 {
				//ZendObjectRelease(fcc.GetObject())
			}
			return 0
		}
		if func_.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(func_.GetOpArray())) {
			InitFuncRunTimeCache(func_.GetOpArray())
		}
	} else {
		faults.InternalTypeError(executeData.IsCallUseStrictTypes(), "%s() expects parameter 1 to be a valid callback, %s", opline.Const1().String().GetVal(), error)
		Efree(error)
		// ZvalPtrDtorNogc(free_op2)
		if EG__().GetException() != nil {
			return 0
		}
		func_ = (types.IFunction)(&ZendPassFunction)
		object_or_called_scope = nil
	}
	call = ZendVmStackPushCallFrame(call_info, func_, opline.GetExtendedValue(), object_or_called_scope)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_USER_CALL_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var fcc types.ZendFcallInfoCache
	var error *byte = nil
	var func_ types.IFunction
	var object_or_called_scope any
	var call *ZendExecuteData
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	function_name = opline.Cv2OrUndef()
	if ZendIsCallableEx(function_name, nil, 0, nil, &fcc, &error) != 0 {
		func_ = fcc.GetFunctionHandler()
		if error != nil {
			Efree(error)

			/* This is the only soft error is_callable() can generate */

			ZendNonStaticMethodCall(func_)
			if EG__().GetException() != nil {
				return 0
			}
		}
		object_or_called_scope = fcc.GetCalledScope()
		if func_.IsClosure() {

			/* Delay closure destruction until its invocation */

			//ZEND_CLOSURE_OBJECT(func_).AddRefcount()
			call_info |= ZEND_CALL_CLOSURE
			if func_.IsFakeClosure() {
				call_info |= ZEND_CALL_FAKE_CLOSURE
			}
			if fcc.GetObject() != nil {
				object_or_called_scope = fcc.GetObject()
				call_info |= ZEND_CALL_HAS_THIS
			}
		} else if fcc.GetObject() != nil {
			//fcc.GetObject().AddRefcount()
			object_or_called_scope = fcc.GetObject()
			call_info |= ZEND_CALL_RELEASE_THIS | ZEND_CALL_HAS_THIS
		}
		if EG__().GetException() != nil {
			if (call_info & ZEND_CALL_CLOSURE) != 0 {
				//ZendObjectRelease(ZEND_CLOSURE_OBJECT(func_))
			} else if (call_info & ZEND_CALL_RELEASE_THIS) != 0 {
				//ZendObjectRelease(fcc.GetObject())
			}
			return 0
		}
		if func_.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(func_.GetOpArray())) {
			InitFuncRunTimeCache(func_.GetOpArray())
		}
	} else {
		faults.InternalTypeError(executeData.IsCallUseStrictTypes(), "%s() expects parameter 1 to be a valid callback, %s", opline.Const1().String().GetVal(), error)
		Efree(error)
		if EG__().GetException() != nil {
			return 0
		}
		func_ = (types.IFunction)(&ZendPassFunction)
		object_or_called_scope = nil
	}
	call = ZendVmStackPushCallFrame(call_info, func_, opline.GetExtendedValue(), object_or_called_scope)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
