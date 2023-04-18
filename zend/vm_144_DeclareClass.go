package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_DECLARE_CLASS_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	DoBindClass(opline.Const1(), b.CondF1(opline.GetOp2Type() == IS_CONST, func() *types.String { return opline.Const2().String() }, nil))
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
