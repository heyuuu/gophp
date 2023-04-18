package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

func ZEND_UNSET_CV_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ *types2.Zval = opline.Op1()
	if var_.IsRefcounted() {
		var garbage *types2.ZendRefcounted = var_.RefCounted()
		var_.SetUndef()
		if garbage.DelRefcount() == 0 {
			//RcDtorFunc(garbage)
		} else {
			//GcCheckPossibleRoot(garbage)
		}
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	} else {
		var_.SetUndef()
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
