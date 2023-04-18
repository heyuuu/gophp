package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_SEPARATE_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = opline.Op1()
	if var_ptr.IsReference() {
		if var_ptr.GetRefcount() == 1 {
			types.ZVAL_UNREF(var_ptr)
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
