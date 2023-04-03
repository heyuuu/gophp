package zend

func ZEND_POW_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Const1()
	op2 = opline.Const2()
	PowFunction(opline.Result(), op1, op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_POW_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Const1()
	op2 = opline.Op2Ptr(&free_op2)
	PowFunction(opline.Result(), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_POW_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Const1()
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	PowFunction(opline.Result(), op1, op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_POW_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1Ptr(&free_op1)
	op2 = opline.Const2()
	PowFunction(opline.Result(), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_POW_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1Ptr(&free_op1)
	op2 = opline.Op2Ptr(&free_op2)
	PowFunction(opline.Result(), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_POW_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1Ptr(&free_op1)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	PowFunction(opline.Result(), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_POW_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = opline.Const2()
	PowFunction(opline.Result(), op1, op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_POW_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = opline.Op2Ptr(&free_op2)
	PowFunction(opline.Result(), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_POW_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	PowFunction(opline.Result(), op1, op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
