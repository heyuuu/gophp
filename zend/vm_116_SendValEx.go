package zend

func ZEND_SEND_VAL_EX_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
	send_val_by_ref:
		return zend_cannot_pass_by_ref_helper_SPEC(executeData)
	}
	value = opline.Const1()
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	arg.CopyValueFrom(value)
	{

		// arg.TryAddRefcount()

	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_SEND_VAL_EX_SPEC_CONST_QUICK_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		goto send_val_by_ref
	}
	value = opline.Const1()
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	arg.CopyValueFrom(value)
	{

		// arg.TryAddRefcount()

	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_SEND_VAL_EX_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
	send_val_by_ref:
		return zend_cannot_pass_by_ref_helper_SPEC(executeData)
	}
	value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	arg.CopyValueFrom(value)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_SEND_VAL_EX_SPEC_TMP_QUICK_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		goto send_val_by_ref
	}
	value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	arg.CopyValueFrom(value)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
