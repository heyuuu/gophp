package zend

func ZEND_QM_ASSIGN_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result *types.Zval = opline.GetResultZval()
	value = RT_CONSTANT(opline, opline.GetOp1())
	{
		types.ZVAL_COPY_VALUE(result, value)
		{

			result.TryAddRefcount()

		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_QM_ASSIGN_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var result *types.Zval = opline.GetResultZval()
	value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if value.IsUndef() {
		ZVAL_UNDEFINED_OP1()
		result.SetNull()
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	{
		types.ZVAL_COPY_DEREF(result, value)
	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_QM_ASSIGN_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var result *types.Zval = opline.GetResultZval()
	value = opline.getZvalPtrVar1(&free_op1)
	if value.IsUndef() {
		ZVAL_UNDEFINED_OP1()
		result.SetNull()
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	{
		types.ZVAL_COPY_DEREF(result, value)
	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_QM_ASSIGN_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result *types.Zval = opline.GetResultZval()
	value = opline.GetOp1Zval()
	if value.IsUndef() {
		ZVAL_UNDEFINED_OP1()
		result.SetNull()
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	{
		types.ZVAL_COPY_DEREF(result, value)
	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
