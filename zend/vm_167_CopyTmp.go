package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_COPY_TMP_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval = opline.Op1()
	var result *types.Zval = opline.Result()
	types.ZVAL_COPY(result, value)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
