package zend

func ZEND_CHECK_VAR_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval = opline.Op1()
	if op1.IsUndef() {
		ZVAL_UNDEFINED_OP1()
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
