package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_UNSET_CV_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var var_ *types.Zval = opline.Op1()
	if var_.IsRefcounted() {
		var_.SetUndef()
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	} else {
		var_.SetUndef()
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
