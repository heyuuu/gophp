package zend

func ZEND_FUNC_NUM_ARGS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	opline.Result().SetLong(executeData.NumArgs())
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
