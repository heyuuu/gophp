package zend

func ZEND_FE_FREE_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var var_ *types.Zval
	var opline *ZendOp = executeData.GetOpline()
	var_ = opline.GetOp1Zval()
	if var_.GetType() != types.IS_ARRAY && var_.GetFeIterIdx() != uint32-1 {
		types.ZendHashIteratorDel(var_.GetFeIterIdx())
	}
	ZvalPtrDtorNogc(var_)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
