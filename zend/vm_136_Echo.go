package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func vmEchoHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	// todo
	var z *types.Zval = opline.Const1()
	var z *types.Zval = opline.Op1()

	str := z.String()
	if len(str) > 0 {
		ZendWrite(str)
	} else if z.IsUndef() {
		ZVAL_UNDEFINED_OP1(executeData)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
