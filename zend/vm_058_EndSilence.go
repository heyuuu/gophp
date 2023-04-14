package zend

func ZEND_END_SILENCE_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if EG__().GetErrorReporting() == 0 && opline.Op1().Long() != 0 {
		EG__().SetErrorReporting(opline.Op1().Long())
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
