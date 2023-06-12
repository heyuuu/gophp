package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_FE_RESET_R_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var array_ptr *types.Zval
	var result *types.Zval
	array_ptr = opline.Const1()
	if array_ptr.IsArray() {
		result = opline.Result()
		result.CopyValueFrom(array_ptr)
		result.SetFePos(0)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		faults.Error(faults.E_WARNING, "Invalid argument supplied for foreach()")
		opline.Result().SetUndef()
		opline.Result().SetFeIterIdx(uint32 - 1)
		return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
	}
}
func ZEND_FE_RESET_R_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var array_ptr *types.Zval
	var result *types.Zval
	array_ptr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if array_ptr.IsArray() {
		result = opline.Result()
		result.CopyValueFrom(array_ptr)
		result.SetFePos(0)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if array_ptr.IsObject() {
		if types.Z_OBJCE_P(array_ptr).GetGetIterator() == nil {
			var properties *types.Array
			if array_ptr.Object().GetProperties() != nil {
				array_ptr.Object().DupProperties()
			}
			properties = types.Z_OBJPROP_P(array_ptr)
			result = opline.Result()
			result.CopyValueFrom(array_ptr)
			if properties.Len() == 0 {
				result.SetFeIterIdx(uint32 - 1)
				return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
			}
			result.SetFeIterIdx(EG__().AddArrayIterator(properties))
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		} else {
			var is_empty bool = ZendFeResetIterator(array_ptr, 0, opline, executeData)
			// ZvalPtrDtorNogc(free_op1)
			if EG__().GetException() != nil {
				return 0
			} else if is_empty != 0 {
				return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
			} else {
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		}
	} else {
		faults.Error(faults.E_WARNING, "Invalid argument supplied for foreach()")
		opline.Result().SetUndef()
		opline.Result().SetFeIterIdx(uint32 - 1)
		// ZvalPtrDtorNogc(free_op1)
		return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
	}
}
func ZEND_FE_RESET_R_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var array_ptr *types.Zval
	var result *types.Zval
	array_ptr = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	if array_ptr.IsArray() {
		result = opline.Result()
		result.CopyValueFrom(array_ptr)
		result.SetFePos(0)
		// ZvalPtrDtorNogc(free_op1)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if array_ptr.IsObject() {
		if types.Z_OBJCE_P(array_ptr).GetGetIterator() == nil {
			var properties *types.Array
			if array_ptr.Object().GetProperties() != nil {
				array_ptr.Object().DupProperties()
			}
			properties = types.Z_OBJPROP_P(array_ptr)
			result = opline.Result()
			result.CopyValueFrom(array_ptr)
			if properties.Len() == 0 {
				result.SetFeIterIdx(uint32 - 1)
				// ZvalPtrDtorNogc(free_op1)
				return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
			}
			result.SetFeIterIdx(EG__().AddArrayIterator(properties))
			// ZvalPtrDtorNogc(free_op1)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		} else {
			var is_empty bool = ZendFeResetIterator(array_ptr, 0, opline, executeData)
			// ZvalPtrDtorNogc(free_op1)
			if EG__().GetException() != nil {
				return 0
			} else if is_empty != 0 {
				return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
			} else {
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		}
	} else {
		faults.Error(faults.E_WARNING, "Invalid argument supplied for foreach()")
		opline.Result().SetUndef()
		opline.Result().SetFeIterIdx(uint32 - 1)
		// ZvalPtrDtorNogc(free_op1)
		return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
	}
}
func ZEND_FE_RESET_R_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var array_ptr *types.Zval
	var result *types.Zval
	array_ptr = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	if array_ptr.IsArray() {
		result = opline.Result()
		result.CopyValueFrom(array_ptr)
		result.SetFePos(0)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if array_ptr.IsObject() {
		if types.Z_OBJCE_P(array_ptr).GetGetIterator() == nil {
			var properties *types.Array
			if array_ptr.Object().GetProperties() != nil {
				array_ptr.Object().DupProperties()
			}
			properties = types.Z_OBJPROP_P(array_ptr)
			result = opline.Result()
			result.CopyValueFrom(array_ptr)
			if properties.Len() == 0 {
				result.SetFeIterIdx(uint32 - 1)
				return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
			}
			result.SetFeIterIdx(EG__().AddArrayIterator(properties))
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		} else {
			var is_empty bool = ZendFeResetIterator(array_ptr, 0, opline, executeData)
			if EG__().GetException() != nil {
				return 0
			} else if is_empty != 0 {
				return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
			} else {
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		}
	} else {
		faults.Error(faults.E_WARNING, "Invalid argument supplied for foreach()")
		opline.Result().SetUndef()
		opline.Result().SetFeIterIdx(uint32 - 1)
		return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
	}
}
