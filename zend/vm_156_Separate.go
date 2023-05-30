package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_SEPARATE_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
