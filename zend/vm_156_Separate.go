package zend

func ZEND_SEPARATE_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
