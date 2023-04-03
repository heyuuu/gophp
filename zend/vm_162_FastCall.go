package zend

func ZEND_FAST_CALL_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var fast_call *types.Zval = opline.GetResultZval()
	fast_call.SetObj(nil)

	/* set return address */

	fast_call.SetOplineNum(opline - executeData.GetFunc().op_array.opcodes)
	return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp1()), 0)
}
