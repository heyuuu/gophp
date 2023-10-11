package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_EXT_FCALL_END_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
