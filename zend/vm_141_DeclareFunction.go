package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_DECLARE_FUNCTION_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	DoBindFunction(opline.Const1())
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
