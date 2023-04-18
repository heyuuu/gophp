package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_FE_RESET_R_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var array_ptr *types2.Zval
	var result *types2.Zval
	array_ptr = opline.Const1()
	if array_ptr.IsArray() {
		result = opline.Result()
		result.CopyValueFrom(array_ptr)
		if result.IsRefcounted() {
			// 			array_ptr.AddRefcount()
		}
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
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var array_ptr *types2.Zval
	var result *types2.Zval
	array_ptr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if array_ptr.IsArray() {
		result = opline.Result()
		result.CopyValueFrom(array_ptr)
		result.SetFePos(0)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if array_ptr.IsObject() {
		if types2.Z_OBJCE_P(array_ptr).GetGetIterator() == nil {
			var properties *types2.Array
			if types2.Z_OBJ_P(array_ptr).GetProperties() != nil && types2.Z_OBJ_P(array_ptr).GetProperties().GetRefcount() > 1 {
				if (types2.Z_OBJ_P(array_ptr).GetProperties().GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) == 0 {
					types2.Z_OBJ_P(array_ptr).GetProperties().DelRefcount()
				}
				types2.Z_OBJ_P(array_ptr).SetProperties(types2.ZendArrayDup(types2.Z_OBJ_P(array_ptr).GetProperties()))
			}
			properties = types2.Z_OBJPROP_P(array_ptr)
			result = opline.Result()
			result.CopyValueFrom(array_ptr)
			if properties.Len() == 0 {
				result.SetFeIterIdx(uint32 - 1)
				return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
			}
			result.SetFeIterIdx(types2.ZendHashIteratorAdd(properties, 0))
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		} else {
			var is_empty types2.ZendBool = ZendFeResetIterator(array_ptr, 0, opline, executeData)
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
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var array_ptr *types2.Zval
	var result *types2.Zval
	array_ptr = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	if array_ptr.IsArray() {
		result = opline.Result()
		result.CopyValueFrom(array_ptr)
		result.SetFePos(0)
		// ZvalPtrDtorNogc(free_op1)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if array_ptr.IsObject() {
		if types2.Z_OBJCE_P(array_ptr).GetGetIterator() == nil {
			var properties *types2.Array
			if types2.Z_OBJ_P(array_ptr).GetProperties() != nil && types2.Z_OBJ_P(array_ptr).GetProperties().GetRefcount() > 1 {
				if (types2.Z_OBJ_P(array_ptr).GetProperties().GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) == 0 {
					types2.Z_OBJ_P(array_ptr).GetProperties().DelRefcount()
				}
				types2.Z_OBJ_P(array_ptr).SetProperties(types2.ZendArrayDup(types2.Z_OBJ_P(array_ptr).GetProperties()))
			}
			properties = types2.Z_OBJPROP_P(array_ptr)
			result = opline.Result()
			result.CopyValueFrom(array_ptr)
			if properties.Len() == 0 {
				result.SetFeIterIdx(uint32 - 1)
				// ZvalPtrDtorNogc(free_op1)
				return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
			}
			result.SetFeIterIdx(types2.ZendHashIteratorAdd(properties, 0))
			// ZvalPtrDtorNogc(free_op1)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		} else {
			var is_empty types2.ZendBool = ZendFeResetIterator(array_ptr, 0, opline, executeData)
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
	var opline *ZendOp = executeData.GetOpline()
	var array_ptr *types2.Zval
	var result *types2.Zval
	array_ptr = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	if array_ptr.IsArray() {
		result = opline.Result()
		result.CopyValueFrom(array_ptr)
		result.SetFePos(0)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if array_ptr.IsObject() {
		if types2.Z_OBJCE_P(array_ptr).GetGetIterator() == nil {
			var properties *types2.Array
			if types2.Z_OBJ_P(array_ptr).GetProperties() != nil && types2.Z_OBJ_P(array_ptr).GetProperties().GetRefcount() > 1 {
				if (types2.Z_OBJ_P(array_ptr).GetProperties().GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) == 0 {
					types2.Z_OBJ_P(array_ptr).GetProperties().DelRefcount()
				}
				types2.Z_OBJ_P(array_ptr).SetProperties(types2.ZendArrayDup(types2.Z_OBJ_P(array_ptr).GetProperties()))
			}
			properties = types2.Z_OBJPROP_P(array_ptr)
			result = opline.Result()
			result.CopyValueFrom(array_ptr)
			if properties.Len() == 0 {
				result.SetFeIterIdx(uint32 - 1)
				return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
			}
			result.SetFeIterIdx(types2.ZendHashIteratorAdd(properties, 0))
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		} else {
			var is_empty types2.ZendBool = ZendFeResetIterator(array_ptr, 0, opline, executeData)
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
