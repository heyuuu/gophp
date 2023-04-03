package zend

func ZEND_SEND_REF_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var varptr *types.Zval
	var arg *types.Zval
	varptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	if varptr.IsError() {
		arg.SetNewEmptyRef()
		types.Z_REFVAL_P(arg).SetNull()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	if varptr.IsReference() {
		varptr.AddRefcount()
	} else {
		types.ZVAL_MAKE_REF_EX(varptr, 2)
	}
	arg.SetReference(varptr.GetRef())
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_SEND_REF_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	varptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	if varptr.IsError() {
		arg.SetNewEmptyRef()
		types.Z_REFVAL_P(arg).SetNull()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	if varptr.IsReference() {
		varptr.AddRefcount()
	} else {
		types.ZVAL_MAKE_REF_EX(varptr, 2)
	}
	arg.SetReference(varptr.GetRef())
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
