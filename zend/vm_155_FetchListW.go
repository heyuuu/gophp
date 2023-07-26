package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_FETCH_LIST_W_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	dim = opline.Const2()
	if opline.Op1().Type() != types.IsIndirect && !(container.IsRef()) {
		faults.Error(faults.E_NOTICE, "Attempting to set reference to non referenceable value")
		zend_fetch_dimension_address_LIST_r(container, dim, IS_CONST, opline, executeData)
	} else {
		zend_fetch_dimension_address_W(container, dim, IS_CONST, opline, executeData)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_LIST_W_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	dim = opline.Op2()
	if opline.Op1().Type() != types.IsIndirect && !(container.IsRef()) {
		faults.Error(faults.E_NOTICE, "Attempting to set reference to non referenceable value")
		zend_fetch_dimension_address_LIST_r(container, dim, IS_TMP_VAR|IS_VAR, opline, executeData)
	} else {
		zend_fetch_dimension_address_W(container, dim, IS_TMP_VAR|IS_VAR, opline, executeData)
	}
	// ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_LIST_W_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	dim = opline.Op2()
	if opline.Op1().Type() != types.IsIndirect && !(container.IsRef()) {
		faults.Error(faults.E_NOTICE, "Attempting to set reference to non referenceable value")
		zend_fetch_dimension_address_LIST_r(container, dim, IS_CV, opline, executeData)
	} else {
		zend_fetch_dimension_address_W(container, dim, IS_CV, opline, executeData)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
