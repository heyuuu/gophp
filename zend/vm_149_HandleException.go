package zend

func ZEND_HANDLE_EXCEPTION_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var throw_op *ZendOp = EG__().GetOplineBeforeException()
	var throw_op_num uint32 = throw_op - executeData.GetFunc().op_array.opcodes
	var i int
	var current_try_catch_offset int = -1
	if (throw_op.GetOpcode() == ZEND_FREE || throw_op.GetOpcode() == ZEND_FE_FREE) && (throw_op.GetExtendedValue()&ZEND_FREE_ON_RETURN) != 0 {

		/* exceptions thrown because of loop var destruction on return/break/...
		 * are logically thrown at the end of the foreach loop, so adjust the
		 * throw_op_num.
		 */

		var range_ *ZendLiveRange = FindLiveRange(executeData.GetFunc().op_array, throw_op_num, throw_op.GetOp1().GetVar())
		throw_op_num = range_.GetEnd()
	}

	/* Find the innermost try/catch/finally the exception was thrown in */

	for i = 0; i < executeData.GetFunc().op_array.last_try_catch; i++ {
		var try_catch *ZendTryCatchElement = executeData.GetFunc().op_array.try_catch_array[i]
		if try_catch.GetTryOp() > throw_op_num {

			/* further blocks will not be relevant... */

			break

			/* further blocks will not be relevant... */

		}
		if throw_op_num < try_catch.GetCatchOp() || throw_op_num < try_catch.GetFinallyEnd() {
			current_try_catch_offset = i
		}
	}
	CleanupUnfinishedCalls(executeData, throw_op_num)
	if (throw_op.GetResultType() & (IS_VAR | IS_TMP_VAR)) != 0 {
		switch throw_op.GetOpcode() {
		case ZEND_ADD_ARRAY_ELEMENT:
			fallthrough
		case ZEND_ADD_ARRAY_UNPACK:
			fallthrough
		case ZEND_ROPE_INIT:
			fallthrough
		case ZEND_ROPE_ADD:

		case ZEND_FETCH_CLASS:
			fallthrough
		case ZEND_DECLARE_ANON_CLASS:

		default:
			ZvalPtrDtorNogc(EX_VAR(throw_op.GetResult().GetVar()))
		}
	}
	return zend_dispatch_try_catch_finally_helper_SPEC(current_try_catch_offset, throw_op_num, executeData)
}
