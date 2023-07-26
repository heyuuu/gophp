package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_ASSIGN_REF_SPEC_VAR_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var variable_ptr *types.Zval
	var value_ptr *types.Zval
	value_ptr = _getZvalPtrPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		variable_ptr = UninitializedZval()
	} else if opline.Op1().Type() != types.IsIndirect {
		faults.ThrowError(nil, "Cannot assign by reference to an array dimension of an object")
		variable_ptr = UninitializedZval()
	} else if value_ptr.IsError() {
		variable_ptr = UninitializedZval()
	} else if opline.GetExtendedValue() == ZEND_RETURNS_FUNCTION && !(value_ptr.IsRef()) {
		variable_ptr = ZendWrongAssignToVariableReference(variable_ptr, value_ptr, opline, executeData)
	} else {
		ZendAssignToVariableReference(variable_ptr, value_ptr)
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), variable_ptr)
	}
	if free_op2 != nil {
		// ZvalPtrDtorNogc(free_op2)
	}
	if free_op1 != nil {
		// ZvalPtrDtorNogc(free_op1)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_REF_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var variable_ptr *types.Zval
	var value_ptr *types.Zval
	value_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp2().GetVar(), executeData)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		variable_ptr = UninitializedZval()
	} else if opline.Op1().Type() != types.IsIndirect {
		faults.ThrowError(nil, "Cannot assign by reference to an array dimension of an object")
		variable_ptr = UninitializedZval()
	} else if value_ptr.IsError() {
		variable_ptr = UninitializedZval()
	} else if opline.GetExtendedValue() == ZEND_RETURNS_FUNCTION && !(value_ptr.IsRef()) {
		variable_ptr = ZendWrongAssignToVariableReference(variable_ptr, value_ptr, opline, executeData)
	} else {
		ZendAssignToVariableReference(variable_ptr, value_ptr)
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), variable_ptr)
	}
	if free_op1 != nil {
		// ZvalPtrDtorNogc(free_op1)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_REF_SPEC_CV_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var variable_ptr *types.Zval
	var value_ptr *types.Zval
	value_ptr = _getZvalPtrPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = opline.Op1()
	if variable_ptr.IsError() {
		variable_ptr = UninitializedZval()
	} else if opline.Op1().Type() != types.IsIndirect {
		faults.ThrowError(nil, "Cannot assign by reference to an array dimension of an object")
		variable_ptr = UninitializedZval()
	} else if value_ptr.IsError() {
		variable_ptr = UninitializedZval()
	} else if opline.GetExtendedValue() == ZEND_RETURNS_FUNCTION && !(value_ptr.IsRef()) {
		variable_ptr = ZendWrongAssignToVariableReference(variable_ptr, value_ptr, opline, executeData)
	} else {
		ZendAssignToVariableReference(variable_ptr, value_ptr)
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), variable_ptr)
	}
	if free_op2 != nil {
		// ZvalPtrDtorNogc(free_op2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_REF_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var variable_ptr *types.Zval
	var value_ptr *types.Zval
	value_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp2().GetVar(), executeData)
	variable_ptr = opline.Op1()
	if variable_ptr.IsError() {
		variable_ptr = UninitializedZval()
	} else if opline.Op1().Type() != types.IsIndirect {
		faults.ThrowError(nil, "Cannot assign by reference to an array dimension of an object")
		variable_ptr = UninitializedZval()
	} else if value_ptr.IsError() {
		variable_ptr = UninitializedZval()
	} else if opline.GetExtendedValue() == ZEND_RETURNS_FUNCTION && !(value_ptr.IsRef()) {
		variable_ptr = ZendWrongAssignToVariableReference(variable_ptr, value_ptr, opline, executeData)
	} else {
		ZendAssignToVariableReference(variable_ptr, value_ptr)
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), variable_ptr)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
