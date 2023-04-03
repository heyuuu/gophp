package zend

func ZEND_FAST_RET_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var fast_call *types.Zval = opline.GetOp1Zval()
	var current_try_catch_offset uint32
	var current_op_num uint32
	if fast_call.GetOplineNum() != uint32-1 {
		var fast_ret *ZendOp = executeData.GetFunc().op_array.opcodes + fast_call.GetOplineNum()
		return ZEND_VM_JMP_EX(executeData, fast_ret+1, 0)
	}

	/* special case for unhandled exceptions */

	EG__().SetException(fast_call.GetObj())
	fast_call.SetObj(nil)
	current_try_catch_offset = opline.GetOp2().GetNum()
	current_op_num = opline - executeData.GetFunc().op_array.opcodes
	return zend_dispatch_try_catch_finally_helper_SPEC(current_try_catch_offset, current_op_num, executeData)
}
