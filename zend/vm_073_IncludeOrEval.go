package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_INCLUDE_OR_EVAL_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var new_op_array *types.ZendOpArray
	var inc_filename *types.Zval
	inc_filename = opline.Const1()
	new_op_array = ZendIncludeOrEval(inc_filename, opline.GetExtendedValue())
	if EG__().GetException() != nil {
		if new_op_array != ZEND_FAKE_OP_ARRAY && new_op_array != nil {
			DestroyOpArray(new_op_array)
			EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
		}
		UNDEF_RESULT()
		return 0
	} else if new_op_array == ZEND_FAKE_OP_ARRAY {
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetTrue()
		}
	} else if new_op_array != nil {
		var return_value *types.Zval = nil
		var call *ZendExecuteData
		if RETURN_VALUE_USED(opline) {
			return_value = opline.Result()
		}
		new_op_array.SetScope(executeData.GetFunc().GetOpArray().scope)
		call = ZendVmStackPushCallFrame(executeData.GetThis().GetTypeInfo()&ZEND_CALL_HAS_THIS|ZEND_CALL_NESTED_CODE|ZEND_CALL_HAS_SYMBOL_TABLE, (types.IFunction)(new_op_array), 0, executeData.GetThis().Ptr())
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
			call.SetSymbolTable(executeData.GetSymbolTable())
		} else {
			call.SetSymbolTable(ZendRebuildSymbolTable())
		}
		call.SetPrevExecuteData(executeData)
		IInitCodeExecuteData(call, new_op_array, return_value)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			ZEND_ADD_CALL_FLAG(call, ZEND_CALL_TOP)
			ZendExecuteEx(call)
			ZendVmStackFreeCallFrame(call)
		}
		DestroyOpArray(new_op_array)
		EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
		if EG__().GetException() != nil {
			faults.RethrowException(executeData)
			UNDEF_RESULT()
			return 0
		}
	} else if RETURN_VALUE_USED(opline) {
		opline.Result().SetFalse()
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INCLUDE_OR_EVAL_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var new_op_array *types.ZendOpArray
	var free_op1 ZendFreeOp
	var inc_filename *types.Zval
	inc_filename = opline.Op1()
	new_op_array = ZendIncludeOrEval(inc_filename, opline.GetExtendedValue())
	// ZvalPtrDtorNogc(free_op1)
	if EG__().GetException() != nil {
		if new_op_array != ZEND_FAKE_OP_ARRAY && new_op_array != nil {
			DestroyOpArray(new_op_array)
			EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
		}
		UNDEF_RESULT()
		return 0
	} else if new_op_array == ZEND_FAKE_OP_ARRAY {
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetTrue()
		}
	} else if new_op_array != nil {
		var return_value *types.Zval = nil
		var call *ZendExecuteData
		if RETURN_VALUE_USED(opline) {
			return_value = opline.Result()
		}
		new_op_array.SetScope(executeData.GetFunc().GetOpArray().scope)
		call = ZendVmStackPushCallFrame(executeData.GetThis().GetTypeInfo()&ZEND_CALL_HAS_THIS|ZEND_CALL_NESTED_CODE|ZEND_CALL_HAS_SYMBOL_TABLE, (types.IFunction)(new_op_array), 0, executeData.GetThis().Ptr())
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
			call.SetSymbolTable(executeData.GetSymbolTable())
		} else {
			call.SetSymbolTable(ZendRebuildSymbolTable())
		}
		call.SetPrevExecuteData(executeData)
		IInitCodeExecuteData(call, new_op_array, return_value)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			ZEND_ADD_CALL_FLAG(call, ZEND_CALL_TOP)
			ZendExecuteEx(call)
			ZendVmStackFreeCallFrame(call)
		}
		DestroyOpArray(new_op_array)
		EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
		if EG__().GetException() != nil {
			faults.RethrowException(executeData)
			UNDEF_RESULT()
			return 0
		}
	} else if RETURN_VALUE_USED(opline) {
		opline.Result().SetFalse()
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INCLUDE_OR_EVAL_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var new_op_array *types.ZendOpArray
	var inc_filename *types.Zval
	inc_filename = opline.Cv1OrUndef()
	new_op_array = ZendIncludeOrEval(inc_filename, opline.GetExtendedValue())
	if EG__().GetException() != nil {
		if new_op_array != ZEND_FAKE_OP_ARRAY && new_op_array != nil {
			DestroyOpArray(new_op_array)
			EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
		}
		UNDEF_RESULT()
		return 0
	} else if new_op_array == ZEND_FAKE_OP_ARRAY {
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetTrue()
		}
	} else if new_op_array != nil {
		var return_value *types.Zval = nil
		var call *ZendExecuteData
		if RETURN_VALUE_USED(opline) {
			return_value = opline.Result()
		}
		new_op_array.SetScope(executeData.GetFunc().GetOpArray().scope)
		call = ZendVmStackPushCallFrame(executeData.GetThis().GetTypeInfo()&ZEND_CALL_HAS_THIS|ZEND_CALL_NESTED_CODE|ZEND_CALL_HAS_SYMBOL_TABLE, (types.IFunction)(new_op_array), 0, executeData.GetThis().Ptr())
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
			call.SetSymbolTable(executeData.GetSymbolTable())
		} else {
			call.SetSymbolTable(ZendRebuildSymbolTable())
		}
		call.SetPrevExecuteData(executeData)
		IInitCodeExecuteData(call, new_op_array, return_value)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			ZEND_ADD_CALL_FLAG(call, ZEND_CALL_TOP)
			ZendExecuteEx(call)
			ZendVmStackFreeCallFrame(call)
		}
		DestroyOpArray(new_op_array)
		EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
		if EG__().GetException() != nil {
			faults.RethrowException(executeData)
			UNDEF_RESULT()
			return 0
		}
	} else if RETURN_VALUE_USED(opline) {
		opline.Result().SetFalse()
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
