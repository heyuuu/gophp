package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_FUNC_NUM_ARGS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	opline.Result().SetLong(executeData.NumArgs())
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
