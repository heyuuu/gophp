package zend

func ZEND_SEND_VAL_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	value = RT_CONSTANT(opline, opline.GetOp1())
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY_VALUE(arg, value)
	{

		arg.TryAddRefcount()

	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_SEND_VAL_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	value = opline.getZvalPtrVar1(&free_op1)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY_VALUE(arg, value)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
