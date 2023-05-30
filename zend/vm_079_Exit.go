package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_EXIT_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if opline.GetOp1Type() != IS_UNUSED {
		var free_op1 ZendFreeOp
		var ptr *types.Zval = GetZvalPtr(opline.GetOp1Type(), opline.GetOp1(), &free_op1, BP_VAR_R)
		for {
			if ptr.IsLong() {
				EG__().SetExitStatus(ptr.Long())
			} else {
				if (opline.GetOp1Type()&(IS_VAR|IS_CV)) != 0 && ptr.IsReference() {
					ptr = types.Z_REFVAL_P(ptr)
					if ptr.IsLong() {
						EG__().SetExitStatus(ptr.Long())
						break
					}
				}
				ZendPrintZval(ptr)
			}
			break
		}
		// 		FREE_OP(free_op1)
	}
	faults.Bailout()
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
