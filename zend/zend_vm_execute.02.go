package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func zend_cannot_pass_by_ref_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	faults.ThrowError(nil, fmt.Sprintf("Cannot pass parameter %d by reference", arg_num))
	FREE_UNFETCHED_OP(opline.GetOp1Type(), opline.GetOp1().GetVar())
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	arg.SetUndef()
	return 0
}
func zend_case_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.CompareFunction(opline.Result(), op_1, op_2)
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		// ZvalPtrDtorNogc(op_2)
	}
	if EG__().HasException() {
		return 0
	}
	if opline.Result().Long() == 0 {
		ZEND_VM_SMART_BRANCH_TRUE()
		opline.Result().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		ZEND_VM_SMART_BRANCH_FALSE()
		opline.Result().SetFalse()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func zend_dispatch_try_catch_finally_helper_SPEC(try_catch_offset uint32, op_num uint32, executeData *ZendExecuteData) int {
	/* May be NULL during generator closing (only finally blocks are executed) */

	var ex *types.Object = EG__().GetException()

	/* Walk try/catch/finally structures upwards, performing the necessary actions */

	for try_catch_offset != uint32-1 {
		var try_catch *ZendTryCatchElement = executeData.GetFunc().GetOpArray().try_catch_array[try_catch_offset]
		if op_num < try_catch.GetCatchOp() && ex != nil {

			/* Go to catch block */

			CleanupLiveVars(executeData, op_num, try_catch.GetCatchOp())
			return ZEND_VM_JMP_EX(executeData, executeData.GetFunc().GetOpArray().opcodes[try_catch.GetCatchOp()], 0)
		} else if op_num < try_catch.GetFinallyOp() {

			/* Go to finally block */

			var fast_call *types.Zval = EX_VAR(executeData.GetFunc().GetOpArray().opcodes[try_catch.GetFinallyEnd()].op1.GetVar())
			CleanupLiveVars(executeData, op_num, try_catch.GetFinallyOp())
			fast_call.SetObject(EG__().GetException())
			EG__().SetException(nil)
			fast_call.SetOplineNum(uint32 - 1)
			return ZEND_VM_JMP_EX(executeData, executeData.GetFunc().GetOpArray().opcodes[try_catch.GetFinallyOp()], 0)
		} else if op_num < try_catch.GetFinallyEnd() {
			var fast_call *types.Zval = EX_VAR(executeData.GetFunc().GetOpArray().opcodes[try_catch.GetFinallyEnd()].op1.GetVar())

			/* cleanup incomplete RETURN statement */

			if fast_call.GetOplineNum() != uint32-1 && (executeData.GetFunc().GetOpArray().opcodes[fast_call.GetOplineNum()].op2_type&(IS_TMP_VAR|IS_VAR)) != 0 {
				var return_value *types.Zval = EX_VAR(executeData.GetFunc().GetOpArray().opcodes[fast_call.GetOplineNum()].op2.GetVar())
				// ZvalPtrDtor(return_value)
			}

			/* Chain potential exception from wrapping finally block */

			if fast_call.Object() != nil {
				if ex != nil {
					faults.ExceptionSetPrevious(ex, fast_call.Object())
				} else {
					EG__().SetException(fast_call.Object())
				}
				ex = fast_call.Object()
			}

			/* Chain potential exception from wrapping finally block */

		}
		try_catch_offset--
	}

	/* Uncaught exception */

	CleanupLiveVars(executeData, op_num, 0)
	if (EX_CALL_INFO() & ZEND_CALL_GENERATOR) != 0 {
		var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
		ZendGeneratorClose(generator, 1)
		return -1
	} else {

		/* We didn't execute RETURN, and have to initialize return_value */

		if executeData.GetReturnValue() {
			executeData.GetReturnValue().
				SetUndef()
		}
		return zend_leave_helper_SPEC(executeData)
	}
}
func zend_yield_in_closed_generator_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	faults.ThrowError(nil, "Cannot yield from finally in a force-closed generator")
	FREE_UNFETCHED_OP(opline.GetOp2Type(), opline.GetOp2().GetVar())
	FREE_UNFETCHED_OP(opline.GetOp1Type(), opline.GetOp1().GetVar())
	UNDEF_RESULT()
	return 0
}
func zend_interrupt_helper_SPEC(executeData *ZendExecuteData) int {
	EG__().SetVmInterrupt(0)
	if EG__().GetTimedOut() != 0 {
		ZendTimeout(0)
	}
	return 0
}
