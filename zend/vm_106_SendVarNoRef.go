package zend

func ZEND_SEND_VAR_NO_REF_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var varptr *types.Zval
	var arg *types.Zval
	varptr = opline.Op1()
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	arg.CopyValueFrom(varptr)
	if varptr.IsReference() {
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	arg.SetNewRef(arg)
	faults.Error(faults.E_NOTICE, "Only variables should be passed by reference")
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
