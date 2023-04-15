package zend

func ZEND_SEND_VAR_NO_REF_EX_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var varptr *types.Zval
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) == 0 {
		return ZEND_SEND_VAR_SPEC_VAR_HANDLER(executeData)
	}
	varptr = opline.Op1()
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	arg.CopyValueFrom(varptr)
	if varptr.IsReference() || ARG_MAY_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	arg.SetNewRef(arg)
	faults.Error(faults.E_NOTICE, "Only variables should be passed by reference")
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_SEND_VAR_NO_REF_EX_SPEC_VAR_QUICK_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var varptr *types.Zval
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) == 0 {
		return ZEND_SEND_VAR_SPEC_VAR_INLINE_HANDLER(executeData)
	}
	varptr = opline.Op1()
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	arg.CopyValueFrom(varptr)
	if varptr.IsReference() || QUICK_ARG_MAY_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	arg.SetNewRef(arg)
	faults.Error(faults.E_NOTICE, "Only variables should be passed by reference")
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
