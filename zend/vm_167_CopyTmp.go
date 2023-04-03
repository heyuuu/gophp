package zend

func ZEND_COPY_TMP_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval = opline.getZvalPtrVar1(&free_op1)
	var result *types.Zval = opline.Result()
	types.ZVAL_COPY(result, value)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
