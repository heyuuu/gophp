package zend

func ZEND_FREE_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	ZvalPtrDtorNogc(opline.GetOp1Zval())
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
