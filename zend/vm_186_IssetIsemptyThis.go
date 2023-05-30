package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_ISSET_ISEMPTY_THIS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	opline.Result().SetBool((opline.GetExtendedValue()&ZEND_ISEMPTY ^ executeData.GetThis().IsObject()) != 0)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
