package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_FE_RESET_RW_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var array_ptr *types.Zval
	var array_ref *types.Zval

	{
		array_ptr = opline.Const1()
		array_ref = array_ptr
	}
	if array_ptr.IsArray() {

		{
			array_ref = opline.Result()
			array_ref.SetNewRef(array_ptr)
			array_ptr = types.Z_REFVAL_P(array_ref)
		}
		{
			array_ptr.SetArray(types.ZendArrayDup(array_ptr.GetArr()))
		}

		opline.Result().SetFeIterIdx(EG__().AddArrayIterator(array_ptr.Array()))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		faults.Error(faults.E_WARNING, "Invalid argument supplied for foreach()")
		opline.Result().SetUndef()
		opline.Result().SetFeIterIdx(uint32 - 1)
		return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
	}
}
func ZEND_FE_RESET_RW_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var array_ptr *types.Zval
	var array_ref *types.Zval
	{
		array_ptr = nil
		array_ref = array_ptr
		if array_ref.IsReference() {
			array_ptr = types.Z_REFVAL_P(array_ref)
		}
	}

	if array_ptr.IsArray() {
		{
			if array_ptr == array_ref {
				array_ref.SetNewRef(array_ref)
				array_ptr = types.Z_REFVAL_P(array_ref)
			}
			// 			array_ref.AddRefcount()
			types.ZVAL_COPY_VALUE(opline.Result(), array_ref)
		}

		{
			types.SeparateArray(array_ptr)
		}
		opline.Result().SetFeIterIdx(EG__().AddArrayIterator(array_ptr.Array()))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if array_ptr.IsObject() {
		if types.Z_OBJCE_P(array_ptr).GetGetIterator() == nil {
			var properties *types.Array
			{
				if array_ptr == array_ref {
					array_ref.SetNewRef(array_ref)
					array_ptr = types.Z_REFVAL_P(array_ref)
				}
				// 				array_ref.AddRefcount()
				types.ZVAL_COPY_VALUE(opline.Result(), array_ref)
			}

			if array_ptr.Object().GetProperties() != nil {
				array_ptr.Object().DupProperties()
			}
			properties = types.Z_OBJPROP_P(array_ptr)
			if properties.Len() == 0 {
				opline.Result().SetFeIterIdx(uint32 - 1)
				return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
			}
			opline.Result().SetFeIterIdx(EG__().AddArrayIterator(properties))
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		} else {
			var is_empty bool = ZendFeResetIterator(array_ptr, 1, opline, executeData)
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
func ZEND_FE_RESET_RW_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var array_ptr *types.Zval
	var array_ref *types.Zval
	{
		array_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
		array_ref = array_ptr
		if array_ref.IsReference() {
			array_ptr = types.Z_REFVAL_P(array_ref)
		}
	}

	if array_ptr.IsArray() {
		{
			if array_ptr == array_ref {
				array_ref.SetNewRef(array_ref)
				array_ptr = types.Z_REFVAL_P(array_ref)
			}
			// 			array_ref.AddRefcount()
			types.ZVAL_COPY_VALUE(opline.Result(), array_ref)
		}

		{
			types.SeparateArray(array_ptr)
		}
		opline.Result().SetFeIterIdx(EG__().AddArrayIterator(array_ptr.Array()))
		{
			if free_op1 != nil {
				// ZvalPtrDtorNogc(free_op1)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if array_ptr.IsObject() {
		if types.Z_OBJCE_P(array_ptr).GetGetIterator() == nil {
			var properties *types.Array
			{
				if array_ptr == array_ref {
					array_ref.SetNewRef(array_ref)
					array_ptr = types.Z_REFVAL_P(array_ref)
				}
				// 				array_ref.AddRefcount()
				types.ZVAL_COPY_VALUE(opline.Result(), array_ref)
			}

			if array_ptr.Object().GetProperties() != nil {
				array_ptr.Object().DupProperties()
			}
			properties = types.Z_OBJPROP_P(array_ptr)
			if properties.Len() == 0 {
				opline.Result().SetFeIterIdx(uint32 - 1)
				if free_op1 != nil {
					// ZvalPtrDtorNogc(free_op1)
				}
				return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
			}
			opline.Result().SetFeIterIdx(EG__().AddArrayIterator(properties))
			if free_op1 != nil {
				// ZvalPtrDtorNogc(free_op1)
			}
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		} else {
			var is_empty bool = ZendFeResetIterator(array_ptr, 1, opline, executeData)
			{
				if free_op1 != nil {
					// ZvalPtrDtorNogc(free_op1)
				}
			}

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
		{
			if free_op1 != nil {
				// ZvalPtrDtorNogc(free_op1)
			}
		}

		return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
	}
}
func ZEND_FE_RESET_RW_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var array_ptr *types.Zval
	var array_ref *types.Zval
	{
		array_ptr = opline.Cv1OrUndef()
		array_ref = array_ptr
		if array_ref.IsReference() {
			array_ptr = types.Z_REFVAL_P(array_ref)
		}
	}

	if array_ptr.IsArray() {
		{
			if array_ptr == array_ref {
				array_ref.SetNewRef(array_ref)
				array_ptr = types.Z_REFVAL_P(array_ref)
			}
			// 			array_ref.AddRefcount()
			types.ZVAL_COPY_VALUE(opline.Result(), array_ref)
		}

		{
			types.SeparateArray(array_ptr)
		}
		opline.Result().SetFeIterIdx(EG__().AddArrayIterator(array_ptr.Array()))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if array_ptr.IsObject() {
		if types.Z_OBJCE_P(array_ptr).GetGetIterator() == nil {
			var properties *types.Array
			{
				if array_ptr == array_ref {
					array_ref.SetNewRef(array_ref)
					array_ptr = types.Z_REFVAL_P(array_ref)
				}
				// 				array_ref.AddRefcount()
				types.ZVAL_COPY_VALUE(opline.Result(), array_ref)
			}

			if array_ptr.Object().GetProperties() != nil {
				array_ptr.Object().DupProperties()
			}
			properties = types.Z_OBJPROP_P(array_ptr)
			if properties.Len() == 0 {
				opline.Result().SetFeIterIdx(uint32 - 1)
				return ZEND_VM_JMP(executeData, OP_JMP_ADDR(opline, opline.GetOp2()))
			}
			opline.Result().SetFeIterIdx(EG__().AddArrayIterator(properties))
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		} else {
			var is_empty bool = ZendFeResetIterator(array_ptr, 1, opline, executeData)
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
