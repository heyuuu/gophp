package zend

func ZEND_TICKS_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if uint32(b.PreInc(&(EG__().GetTicksCount())) >= opline.GetExtendedValue()) != 0 {
		EG__().SetTicksCount(0)
		if ZendTicksFunction != nil {
			ZendTicksFunction(opline.GetExtendedValue())
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
