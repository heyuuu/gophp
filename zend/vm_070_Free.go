package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_FREE_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	// ZvalPtrDtorNogc(opline.Op1())
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
