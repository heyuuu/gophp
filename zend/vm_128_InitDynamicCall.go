package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_INIT_DYNAMIC_CALL_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var call *ZendExecuteData
	function_name = opline.Const2()
try_function_name:

	if function_name.IsArray() {
		call = ZendInitDynamicCallArray(function_name.GetArr(), opline.GetExtendedValue())
	} else {
		faults.ThrowError(nil, "Function name must be a string")
		call = nil
	}
	if call == nil {
		return 0
	}
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INIT_DYNAMIC_CALL_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var function_name *types.Zval
	var call *ZendExecuteData
	function_name = opline.Op2()
try_function_name:
	if function_name.IsString() {
		call = ZendInitDynamicCallString(function_name.GetStr(), opline.GetExtendedValue())
	} else if function_name.IsObject() {
		call = ZendInitDynamicCallObject(function_name, opline.GetExtendedValue())
	} else if function_name.IsArray() {
		call = ZendInitDynamicCallArray(function_name.GetArr(), opline.GetExtendedValue())
	} else if function_name.IsRef() {
		function_name = types.Z_REFVAL_P(function_name)
		goto try_function_name
	} else {
		if function_name.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
			if EG__().HasException() {
				return 0
			}
		}
		faults.ThrowError(nil, "Function name must be a string")
		call = nil
	}
	// ZvalPtrDtorNogc(free_op2)
	if call == nil {
		return 0
	}
	{
		if EG__().HasException() {
			if call != nil {
				if call.GetFunc().IsCallViaTrampoline() {
					// types.ZendStringReleaseEx(call.GetFunc().GetFunctionName(), 0)
					ZendFreeTrampoline(call.GetFunc())
				}
				ZendVmStackFreeCallFrame(call)
			}
			return 0
		}
	}
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INIT_DYNAMIC_CALL_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var call *ZendExecuteData
	function_name = opline.Op2()
try_function_name:
	if function_name.IsString() {
		call = ZendInitDynamicCallString(function_name.GetStr(), opline.GetExtendedValue())
	} else if function_name.IsObject() {
		call = ZendInitDynamicCallObject(function_name, opline.GetExtendedValue())
	} else if function_name.IsArray() {
		call = ZendInitDynamicCallArray(function_name.GetArr(), opline.GetExtendedValue())
	} else if function_name.IsRef() {
		function_name = types.Z_REFVAL_P(function_name)
		goto try_function_name
	} else {
		if function_name.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
			if EG__().HasException() {
				return 0
			}
		}
		faults.ThrowError(nil, "Function name must be a string")
		call = nil
	}
	if call == nil {
		return 0
	}
	{
		if EG__().HasException() {
			if call != nil {
				if call.GetFunc().IsCallViaTrampoline() {
					// types.ZendStringReleaseEx(call.GetFunc().GetFunctionName(), 0)
					ZendFreeTrampoline(call.GetFunc())
				}
				ZendVmStackFreeCallFrame(call)
			}
			return 0
		}
	}
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
