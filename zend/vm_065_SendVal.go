package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_SEND_VAL_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	value = opline.Const1()
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	arg.CopyValueFrom(value)
	{

		// arg.TryAddRefcount()

	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_SEND_VAL_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	value = opline.Op1()
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	arg.CopyValueFrom(value)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
