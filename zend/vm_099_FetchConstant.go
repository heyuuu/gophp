package zend

func ZEND_FETCH_CONSTANT_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var c *ZendConstant
	c = CACHED_PTR(opline.GetExtendedValue())
	if c != nil && IS_SPECIAL_CACHE_VAL(c) == 0 {
		types.ZVAL_COPY_OR_DUP(opline.Result(), c.Value())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	ZendQuickGetConstant(RT_CONSTANT(opline, opline.GetOp2())+1, opline.GetOp1().GetNum(), opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
