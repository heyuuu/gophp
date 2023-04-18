package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_FETCH_DIM_IS_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = opline.Const1()
	zend_fetch_dimension_address_read_IS(container, opline.Const2(), IS_CONST, opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_IS_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = opline.Const1()
	zend_fetch_dimension_address_read_IS(container, opline.Op2(), IS_TMP_VAR|IS_VAR, opline, executeData)
	// ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_IS_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = opline.Const1()
	zend_fetch_dimension_address_read_IS(container, opline.Op2(), IS_CV, opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_IS_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = opline.Op1()
	zend_fetch_dimension_address_read_IS(container, opline.Const2(), IS_CONST, opline, executeData)
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_IS_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = opline.Op1()
	zend_fetch_dimension_address_read_IS(container, opline.Op2(), IS_TMP_VAR|IS_VAR, opline, executeData)
	// ZvalPtrDtorNogc(free_op2)
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_IS_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = opline.Op1()
	zend_fetch_dimension_address_read_IS(container, opline.Op2(), IS_CV, opline, executeData)
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_IS_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = opline.Op1()
	zend_fetch_dimension_address_read_IS(container, opline.Const2(), IS_CONST, opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_IS_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = opline.Op1()
	zend_fetch_dimension_address_read_IS(container, opline.Op2(), IS_TMP_VAR|IS_VAR, opline, executeData)
	// ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_IS_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = opline.Op1()
	zend_fetch_dimension_address_read_IS(container, opline.Op2(), IS_CV, opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
