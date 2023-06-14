package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_GENERATOR_CREATE_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var return_value *types.Zval = executeData.GetReturnValue()
	if return_value != nil {
		var opline *types.ZendOp = executeData.GetOpline()
		var generator *ZendGenerator
		var gen_execute_data *ZendExecuteData
		var num_args uint32
		var used_stack uint32
		var call_info uint32
		ObjectInitEx(return_value, ZendCeGenerator)

		/*
		 * Normally the executeData is allocated on the VM stack (because it does
		 * not actually do any allocation and thus is faster). For generators
		 * though this behavior would be suboptimal, because the (rather large)
		 * structure would have to be copied back and forth every time execution is
		 * suspended or resumed. That's why for generators the execution context
		 * is allocated on heap.
		 */

		num_args = executeData.NumArgs()
		if num_args <= executeData.GetFunc().GetOpArray().GetNumArgs() {
			used_stack = (ZEND_CALL_FRAME_SLOT + executeData.GetFunc().GetOpArray().GetLastVar() + executeData.GetFunc().GetOpArray().T) * b.SizeOf("zval")
			gen_execute_data = (*ZendExecuteData)(Emalloc(used_stack))
			used_stack = (ZEND_CALL_FRAME_SLOT + executeData.GetFunc().GetOpArray().GetLastVar()) * b.SizeOf("zval")
		} else {
			used_stack = (ZEND_CALL_FRAME_SLOT + num_args + executeData.GetFunc().GetOpArray().GetLastVar() + executeData.GetFunc().GetOpArray().T - executeData.GetFunc().GetOpArray().GetNumArgs()) * b.SizeOf("zval")
			gen_execute_data = (*ZendExecuteData)(Emalloc(used_stack))
		}
		memcpy(gen_execute_data, executeData, used_stack)

		/* Save execution context in generator object. */

		generator = (*ZendGenerator)(executeData.GetReturnValue().Object())
		generator.SetExecuteData(gen_execute_data)
		generator.SetFrozenCallStack(nil)
		generator.GetExecuteFake().SetOpline(nil)
		generator.GetExecuteFake().SetFunc(nil)
		generator.GetExecuteFake().SetPrevExecuteData(nil)
		generator.GetExecuteFake().GetThis().SetObject((*types.ZendObject)(generator))
		gen_execute_data.SetOpline(opline + 1)

		/* EX(return_value) keeps pointer to zend_object (not a real zval) */
		gen_execute_data.SetReturnValue((*types.Zval)(generator))

		call_info = executeData.CallInfo()
		if (call_info&types.Z_TYPE_MASK) == types.IS_OBJECT && ((call_info&(ZEND_CALL_CLOSURE|ZEND_CALL_RELEASE_THIS)) == 0 || ZendExecuteEx != ExecuteEx) {
			ZEND_ADD_CALL_FLAG_EX(call_info, ZEND_CALL_RELEASE_THIS)
		}
		ZEND_ADD_CALL_FLAG_EX(call_info, ZEND_CALL_TOP_FUNCTION|ZEND_CALL_ALLOCATED|ZEND_CALL_GENERATOR)
		gen_execute_data.SetCallInfo(call_info)
		gen_execute_data.SetPrevExecuteData(nil)

		call_info = executeData.CallInfo()
		EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())

		if (call_info & (ZEND_CALL_TOP | ZEND_CALL_ALLOCATED)) == 0 {
			EG__().VmStack().PopCheck(executeData)

			executeData = executeData.GetPrevExecuteData()
			ZEND_VM_INC_OPCODE(executeData)
			return 2
		} else if (call_info & ZEND_CALL_TOP) == 0 {
			EG__().VmStack().PopCheck(executeData)

			executeData = executeData.GetPrevExecuteData()
			ZEND_VM_INC_OPCODE(executeData)
			return 2
		} else {
			return -1
		}
	} else {
		return zend_leave_helper_SPEC(executeData)
	}
}
