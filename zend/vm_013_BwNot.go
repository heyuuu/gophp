package zend

func ZEND_BW_NOT_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	if op1.IsLong() {
		opline.Result().SetLong(^(op1.GetLval()))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	BitwiseNotFunction(opline.Result(), op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_BW_NOT_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	op1 = opline.Op1Ptr(&free_op1)
	if op1.IsLong() {
		opline.Result().SetLong(^(op1.GetLval()))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	if op1.IsUndef() {
		op1 = ZVAL_UNDEFINED_OP1()
	}
	BitwiseNotFunction(opline.Result(), op1)
	ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_BW_NOT_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	op1 = opline.Op1()
	if op1.IsLong() {
		opline.Result().SetLong(^(op1.GetLval()))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	if op1.IsUndef() {
		op1 = ZVAL_UNDEFINED_OP1()
	}
	BitwiseNotFunction(opline.Result(), op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
