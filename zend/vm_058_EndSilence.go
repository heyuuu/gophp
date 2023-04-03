package zend

func ZEND_END_SILENCE_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if EG__().GetErrorReporting() == 0 && opline.GetOp1Zval().GetLval() != 0 {
		EG__().SetErrorReporting(opline.GetOp1Zval().GetLval())
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
