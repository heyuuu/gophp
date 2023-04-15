package zend

import "github.com/heyuuu/gophp/zend/types"

func ZEND_DISCARD_EXCEPTION_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var fast_call *types.Zval = opline.Op1()

	/* cleanup incomplete RETURN statement */

	if fast_call.GetOplineNum() != uint32-1 && (executeData.GetFunc().GetOpArray().opcodes[fast_call.GetOplineNum()].GetOp2Type()&(IS_TMP_VAR|IS_VAR)) != 0 {
		var return_value *types.Zval = EX_VAR(executeData.GetFunc().GetOpArray().opcodes[fast_call.GetOplineNum()].op2.var_)
		// ZvalPtrDtor(return_value)
	}

	/* cleanup delayed exception */
	if fast_call.Object() != nil {
		/* discard the previously thrown exception */
		// OBJ_RELEASE(fast_call.Object())
		fast_call.SetObject(nil)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
