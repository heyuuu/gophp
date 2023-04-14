package zend

import "github.com/heyuuu/gophp/zend/types"

func ZEND_FAST_CALL_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var fast_call *types.Zval = opline.Result()
	fast_call.SetObject(nil)

	/* set return address */
	fast_call.SetOplineNum(opline - executeData.GetFunc().GetOpArray().opcodes)
	return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp1()), 0)
}
