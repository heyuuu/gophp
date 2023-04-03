package zend

func ZEND_DECLARE_FUNCTION_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	DoBindFunction(RT_CONSTANT(opline, opline.GetOp1()))
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
