package zend

func ZEND_JMP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp1()), 0)
}
