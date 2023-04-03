package zend

func ZEND_SEND_USER_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg *types.Zval
	var param *types.Zval
	if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, opline.GetOp2().GetNum()) != 0 {
		ZendParamMustBeRef(executeData.GetCall().func_, opline.GetOp2().GetNum())
	}
	arg = opline.Const1()
	param = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY(param, arg)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_SEND_USER_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg *types.Zval
	var param *types.Zval
	var free_op1 ZendFreeOp
	if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, opline.GetOp2().GetNum()) != 0 {
		ZendParamMustBeRef(executeData.GetCall().func_, opline.GetOp2().GetNum())
	}
	arg = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	param = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY(param, arg)
	ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_SEND_USER_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg *types.Zval
	var param *types.Zval
	var free_op1 ZendFreeOp
	if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, opline.GetOp2().GetNum()) != 0 {
		ZendParamMustBeRef(executeData.GetCall().func_, opline.GetOp2().GetNum())
	}
	arg = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	param = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY(param, arg)
	ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_SEND_USER_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg *types.Zval
	var param *types.Zval
	if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, opline.GetOp2().GetNum()) != 0 {
		ZendParamMustBeRef(executeData.GetCall().func_, opline.GetOp2().GetNum())
	}
	arg = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	param = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY(param, arg)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
