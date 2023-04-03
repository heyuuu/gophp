package zend

func ZEND_FETCH_LIST_R_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	zend_fetch_dimension_address_LIST_r(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_LIST_R_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	zend_fetch_dimension_address_LIST_r(container, opline.getZvalPtrVar2(&free_op2), IS_TMP_VAR|IS_VAR, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_LIST_R_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	zend_fetch_dimension_address_LIST_r(container, opline.Op2(), IS_CV, opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = opline.Op1()
	zend_fetch_dimension_address_LIST_r(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = opline.Op1()
	zend_fetch_dimension_address_LIST_r(container, opline.getZvalPtrVar2(&free_op2), IS_TMP_VAR|IS_VAR, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = opline.Op1()
	zend_fetch_dimension_address_LIST_r(container, opline.Op2(), IS_CV, opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
