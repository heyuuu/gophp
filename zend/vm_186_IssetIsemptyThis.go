package zend

func ZEND_ISSET_ISEMPTY_THIS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	types.ZVAL_BOOL(opline.GetResultZval(), (opline.GetExtendedValue()&ZEND_ISEMPTY^executeData.GetThis().IsObject()) != 0)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
