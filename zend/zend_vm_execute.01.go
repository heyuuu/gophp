package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZEND_VM_DISPATCH(executeData *ZendExecuteData, opcode OpCode, opline *types.ZendOp) int {
	return ZendVmGetOpcodeHandler(opcode, opline)(executeData)
}
func zend_mod_by_zero_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	faults.ThrowExceptionEx(faults.ZendCeDivisionByZeroError, 0, "Modulo by zero")
	opline.Result().SetUndef()
	return 0
}
func zend_mod_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.ModFunction(opline.Result(), op_1, op_2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_shift_left_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.ShiftLeftFunction(opline.Result(), op_1, op_2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_shift_right_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.ShiftRightFunction(opline.Result(), op_1, op_2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_is_equal_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.CompareFunction(opline.Result(), op_1, op_2)
	if EG__().GetException() != nil {
		return 0
	}
	if opline.Result().Long() == 0 {
		ZEND_VM_SMART_BRANCH_TRUE()
		opline.Result().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		ZEND_VM_SMART_BRANCH_FALSE()
		opline.Result().SetFalse()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func zend_is_not_equal_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.CompareFunction(opline.Result(), op_1, op_2)
	if EG__().GetException() != nil {
		return 0
	}
	if opline.Result().Long() != 0 {
		ZEND_VM_SMART_BRANCH_TRUE()
		opline.Result().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		ZEND_VM_SMART_BRANCH_FALSE()
		opline.Result().SetFalse()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func zend_is_smaller_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.CompareFunction(opline.Result(), op_1, op_2)
	if EG__().GetException() != nil {
		return 0
	}
	if opline.Result().Long() < 0 {
		ZEND_VM_SMART_BRANCH_TRUE()
		opline.Result().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		ZEND_VM_SMART_BRANCH_FALSE()
		opline.Result().SetFalse()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func zend_is_smaller_or_equal_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.CompareFunction(opline.Result(), op_1, op_2)
	if EG__().GetException() != nil {
		return 0
	}
	if opline.Result().Long() <= 0 {
		ZEND_VM_SMART_BRANCH_TRUE()
		opline.Result().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		ZEND_VM_SMART_BRANCH_FALSE()
		opline.Result().SetFalse()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func zend_bw_or_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.BitwiseOrFunction(opline.Result(), op_1, op_2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_bw_and_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.BitwiseAndFunction(opline.Result(), op_1, op_2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_bw_xor_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.BitwiseXorFunction(opline.Result(), op_1, op_2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_this_not_in_object_context_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	faults.ThrowError(nil, "Using $this when not in object context")
	if (opline + 1).GetOpcode() == ZEND_OP_DATA {
		FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
	}
	FREE_UNFETCHED_OP(opline.GetOp2Type(), opline.GetOp2().GetVar())
	UNDEF_RESULT()
	return 0
}
func zend_undefined_function_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	function_name = opline.Const2()
	faults.ThrowError(nil, "Call to undefined function %s()", function_name.String().GetVal())
	return 0
}
func zend_fetch_static_prop_helper_SPEC(type_ int, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var prop *types.Zval
	if ZendFetchStaticPropertyAddress(&prop, nil, opline.GetExtendedValue() & ^ZEND_FETCH_OBJ_FLAGS, type_, opline.GetExtendedValue()&ZEND_FETCH_OBJ_FLAGS, opline, executeData) != types.SUCCESS {
		b.Assert(EG__().GetException() != nil || type_ == BP_VAR_IS)
		prop = UninitializedZval()
	}
	if type_ == BP_VAR_R || type_ == BP_VAR_IS {
		types.ZVAL_COPY_DEREF(opline.Result(), prop)
	} else {
		opline.Result().SetIndirect(prop)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_use_tmp_in_write_context_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	faults.ThrowError(nil, "Cannot use temporary expression in write context")
	FREE_UNFETCHED_OP(opline.GetOp2Type(), opline.GetOp2().GetVar())
	FREE_UNFETCHED_OP(opline.GetOp1Type(), opline.GetOp1().GetVar())
	opline.Result().SetUndef()
	return 0
}
func zend_use_undef_in_read_context_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	faults.ThrowError(nil, "Cannot use [] for reading")
	FREE_UNFETCHED_OP(opline.GetOp2Type(), opline.GetOp2().GetVar())
	FREE_UNFETCHED_OP(opline.GetOp1Type(), opline.GetOp1().GetVar())
	opline.Result().SetUndef()
	return 0
}
func zend_leave_helper_SPEC(executeData *ZendExecuteData) int {
	var old_execute_data *ZendExecuteData
	var call_info uint32 = EX_CALL_INFO()
	if (call_info & (ZEND_CALL_CODE | ZEND_CALL_TOP | ZEND_CALL_HAS_SYMBOL_TABLE | ZEND_CALL_FREE_EXTRA_ARGS | ZEND_CALL_ALLOCATED)) == 0 {
		EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
		IFreeCompiledVariables(executeData)

		EG__().VmStack().PopCheck(executeData)

		executeData = executeData.GetPrevExecuteData()
		if EG__().GetException() != nil {
			faults.RethrowException(executeData)
			return 2
		}
		ZEND_VM_INC_OPCODE(executeData)
		return 2
	} else if (call_info & (ZEND_CALL_CODE | ZEND_CALL_TOP)) == 0 {
		EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
		IFreeCompiledVariables(executeData)
		if (call_info & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
			/* Free extra args before releasing the closure,
			 * as that may free the op_array. */
			ZendCleanAndCacheSymbolTable(executeData.GetSymbolTable())
		}

		ZendVmStackFreeExtraArgsEx(call_info, executeData)
		if (call_info & ZEND_CALL_RELEASE_THIS) != 0 {
			// OBJ_RELEASE(executeData.GetThis().Object())
		} else if (call_info & ZEND_CALL_CLOSURE) != 0 {
			// OBJ_RELEASE(ZEND_CLOSURE_OBJECT(executeData.GetFunc()))
		}
		old_execute_data = executeData
		executeData = executeData.GetPrevExecuteData()
		ZendVmStackFreeCallFrame(old_execute_data)
		if EG__().GetException() != nil {
			faults.RethrowException(executeData)
			return 2
		}
		ZEND_VM_INC_OPCODE(executeData)
		return 2
	} else if (call_info & ZEND_CALL_TOP) == 0 {
		ZendDetachSymbolTable(executeData)
		//DestroyOpArray(executeData.GetFunc().GetOpArray())
		//EfreeSize(executeData.GetFunc(), b.SizeOf("zend_op_array"))
		old_execute_data = executeData
		EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
		executeData = CurrEX()
		ZendVmStackFreeCallFrame(old_execute_data)
		ZendAttachSymbolTable(executeData)
		if EG__().GetException() != nil {
			faults.RethrowException(executeData)
			return 2
		}
		ZEND_VM_INC_OPCODE(executeData)
		return 2
	} else {
		if (call_info & ZEND_CALL_CODE) == 0 {
			EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
			IFreeCompiledVariables(executeData)
			if (call_info & (ZEND_CALL_HAS_SYMBOL_TABLE | ZEND_CALL_FREE_EXTRA_ARGS)) != 0 {
				if (call_info & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
					ZendCleanAndCacheSymbolTable(executeData.GetSymbolTable())
				}
				ZendVmStackFreeExtraArgsEx(call_info, executeData)
			}
			if (call_info & ZEND_CALL_CLOSURE) != 0 {
				// OBJ_RELEASE(ZEND_CLOSURE_OBJECT(executeData.GetFunc()))
			}
			return -1
		} else {
			var symbol_table *types.Array = executeData.GetSymbolTable()
			ZendDetachSymbolTable(executeData)
			old_execute_data = executeData.GetPrevExecuteData()
			for old_execute_data != nil {
				if old_execute_data.GetFunc() != nil && (ZEND_CALL_INFO(old_execute_data)&ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
					if old_execute_data.GetSymbolTable() == symbol_table {
						ZendAttachSymbolTable(old_execute_data)
					}
					break
				}
				old_execute_data = old_execute_data.GetPrevExecuteData()
			}
			EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
			return -1
		}
	}
}
