package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_SEND_VAR_EX_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
	send_var_by_ref:
		return ZEND_SEND_REF_SPEC_VAR_HANDLER(executeData)
	}
	varptr = opline.Op1()
	if varptr.IsUndef() {
		ZVAL_UNDEFINED_OP1(executeData)
		arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
		arg.SetNull()
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	{
		types.ZVAL_COPY_DEREF(arg, varptr)
	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_SEND_VAR_EX_SPEC_VAR_QUICK_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		goto send_var_by_ref
	}
	varptr = opline.Op1()
	if varptr.IsUndef() {
		ZVAL_UNDEFINED_OP1(executeData)
		arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
		arg.SetNull()
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	{
		types.ZVAL_COPY_DEREF(arg, varptr)
	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_SEND_VAR_EX_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
	send_var_by_ref:
		return ZEND_SEND_REF_SPEC_CV_HANDLER(executeData)
	}
	varptr = opline.Op1()
	if varptr.IsUndef() {
		ZVAL_UNDEFINED_OP1(executeData)
		arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
		arg.SetNull()
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	{
		types.ZVAL_COPY_DEREF(arg, varptr)
	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_SEND_VAR_EX_SPEC_CV_QUICK_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varptr *types.Zval
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if QUICK_ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
		goto send_var_by_ref
	}
	varptr = opline.Op1()
	if varptr.IsUndef() {
		ZVAL_UNDEFINED_OP1(executeData)
		arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
		arg.SetNull()
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	{
		types.ZVAL_COPY_DEREF(arg, varptr)
	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
