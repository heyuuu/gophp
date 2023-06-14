package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_ISSET_ISEMPTY_THIS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()

	cond1 := opline.GetExtendedValue()&ZEND_ISEMPTY != 0
	cond2 := executeData.InScope()
	result := b.Xor(cond1, cond2)

	opline.Result().SetBool(result)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
