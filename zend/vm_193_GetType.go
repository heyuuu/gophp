package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

func ZEND_GET_TYPE_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types2.Zval
	op1 = executeData.GetOp1(opline)

	typ := types2.ZvalGetType(op1)
	opline.Result().SetStringVal(typ)

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_GET_TYPE_SPEC_TMP_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types2.Zval
	op1 = executeData.GetVarOp1(opline)

	typ := types2.ZvalGetType(op1)
	opline.Result().SetStringVal(typ)

	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_GET_TYPE_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types2.Zval
	op1 = executeData.GetVarOp1(opline).DeRef()

	typ := types2.ZvalGetType(op1)
	opline.Result().SetStringVal(typ)

	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_GET_TYPE_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types2.Zval
	op1 = executeData.GetCvOp1(opline)

	typ := types2.ZvalGetType(op1)
	opline.Result().SetStringVal(typ)

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
