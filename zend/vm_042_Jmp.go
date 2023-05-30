package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_JMP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp1()), 0)
}
