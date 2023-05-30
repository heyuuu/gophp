package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_ASSERT_CHECK_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if EG__().GetAssertions() <= 0 {
		var target *types.ZendOp = OP_JMP_ADDR(opline, opline.GetOp2())
		if RETURN_VALUE_USED(opline) {
			opline.Result().SetTrue()
		}
		return ZEND_VM_JMP_EX(executeData, target, 0)
	} else {
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
