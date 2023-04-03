package zend

func ZEND_CHECK_FUNC_ARG_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		ZEND_ADD_CALL_FLAG(executeData.GetCall(), ZEND_CALL_SEND_ARG_BY_REF)
	} else {
		ZEND_DEL_CALL_FLAG(executeData.GetCall(), ZEND_CALL_SEND_ARG_BY_REF)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_CHECK_FUNC_ARG_SPEC_UNUSED_QUICK_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		ZEND_ADD_CALL_FLAG(executeData.GetCall(), ZEND_CALL_SEND_ARG_BY_REF)
	} else {
		ZEND_DEL_CALL_FLAG(executeData.GetCall(), ZEND_CALL_SEND_ARG_BY_REF)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
