package zend

func ZEND_PRE_INC_SPEC_VAR_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsLong() {
		FastLongIncrementFunction(var_ptr)

		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_pre_inc_helper_SPEC_VAR(executeData)
}
func ZEND_PRE_INC_SPEC_VAR_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsLong() {
		FastLongIncrementFunction(var_ptr)
		types.ZVAL_COPY_VALUE(opline.GetResultZval(), var_ptr)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_pre_inc_helper_SPEC_VAR(executeData)
}
func ZEND_PRE_INC_SPEC_CV_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = opline.GetOp1Zval()
	if var_ptr.IsLong() {
		FastLongIncrementFunction(var_ptr)

		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_pre_inc_helper_SPEC_CV(executeData)
}
func ZEND_PRE_INC_SPEC_CV_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = opline.GetOp1Zval()
	if var_ptr.IsLong() {
		FastLongIncrementFunction(var_ptr)
		types.ZVAL_COPY_VALUE(opline.GetResultZval(), var_ptr)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_pre_inc_helper_SPEC_CV(executeData)
}
