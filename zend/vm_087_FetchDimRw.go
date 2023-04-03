package zend

func ZEND_FETCH_DIM_RW_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	zend_fetch_dimension_address_RW(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	{
		var result *types.Zval = opline.Result()
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_RW_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	zend_fetch_dimension_address_RW(container, opline.Op2Ptr(&free_op2), IS_TMP_VAR|IS_VAR, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	{
		var result *types.Zval = opline.Result()
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_RW_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	zend_fetch_dimension_address_RW(container, nil, IS_UNUSED, opline, executeData)
	{
		var result *types.Zval = opline.Result()
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_RW_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	zend_fetch_dimension_address_RW(container, opline.Op2(), IS_CV, opline, executeData)
	{
		var result *types.Zval = opline.Result()
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_RW_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = opline.Op1()
	zend_fetch_dimension_address_RW(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	{
		var result *types.Zval = opline.Result()
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_RW_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = opline.Op1()
	zend_fetch_dimension_address_RW(container, opline.Op2Ptr(&free_op2), IS_TMP_VAR|IS_VAR, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	{
		var result *types.Zval = opline.Result()
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_RW_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = opline.Op1()
	zend_fetch_dimension_address_RW(container, nil, IS_UNUSED, opline, executeData)
	{
		var result *types.Zval = opline.Result()
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_RW_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = opline.Op1()
	zend_fetch_dimension_address_RW(container, opline.Op2(), IS_CV, opline, executeData)
	{
		var result *types.Zval = opline.Result()
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
