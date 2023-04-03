package zend

func ZEND_SEND_FUNC_ARG_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		return ZEND_SEND_REF_SPEC_VAR_HANDLER(executeData)
	}
	varptr = opline.Op1Ptr(&free_op1)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	if varptr.IsReference() {
		var ref *types.ZendRefcounted = varptr.GetCounted()
		varptr = types.Z_REFVAL_P(varptr)
		types.ZVAL_COPY_VALUE(arg, varptr)
		if ref.DelRefcount() == 0 {
			EfreeSize(ref, b.SizeOf("zend_reference"))
		} else {
			arg.TryAddRefcount()
		}

	} else {
		types.ZVAL_COPY_VALUE(arg, varptr)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
