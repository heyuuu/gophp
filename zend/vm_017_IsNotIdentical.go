package zend

import "github.com/heyuuu/gophp/zend/types"

func ZEND_IS_NOT_IDENTICAL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = executeData.GetOp1(opline)
	op2 = executeData.GetOp2(opline)
	result = FastIsNotIdenticalFunction(op1, op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_NOT_IDENTICAL_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = executeData.GetVarOp1(opline)
	op2 = executeData.GetOp2(opline)
	result = FastIsNotIdenticalFunction(op1, op2)
	// ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_NOT_IDENTICAL_SPEC_TMP_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = executeData.GetVarOp1(opline)
	op2 = executeData.GetVarOp2(opline)
	result = FastIsNotIdenticalFunction(op1, op2)
	// ZvalPtrDtorNogc(free_op1)
	// ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_NOT_IDENTICAL_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = executeData.GetVarOp1(opline).DeRef()
	op2 = executeData.GetOp2(opline)
	result = FastIsNotIdenticalFunction(op1, op2)
	// ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_NOT_IDENTICAL_SPEC_VAR_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = executeData.GetVarOp1(opline).DeRef()
	op2 = executeData.GetVarOp2(opline)
	result = FastIsNotIdenticalFunction(op1, op2)
	// ZvalPtrDtorNogc(free_op1)
	// ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_NOT_IDENTICAL_SPEC_VAR_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = executeData.GetVarOp1(opline).DeRef()
	op2 = executeData.GetVarOp2(opline).DeRef()
	result = FastIsNotIdenticalFunction(op1, op2)
	// ZvalPtrDtorNogc(free_op1)
	// ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_NOT_IDENTICAL_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = executeData.GetCvOp1(opline)
	op2 = executeData.GetOp2(opline)
	result = FastIsNotIdenticalFunction(op1, op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_NOT_IDENTICAL_SPEC_CV_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = executeData.GetCvOp1(opline)
	op2 = executeData.GetVarOp2(opline)
	result = FastIsNotIdenticalFunction(op1, op2)
	// ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_NOT_IDENTICAL_SPEC_CV_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = executeData.GetCvOp1(opline)
	op2 = executeData.GetVarOp2(opline).DeRef()
	result = FastIsNotIdenticalFunction(op1, op2)
	// ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_NOT_IDENTICAL_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = executeData.GetCvOp1(opline)
	op2 = executeData.GetCvOp2(opline)
	result = FastIsNotIdenticalFunction(op1, op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
