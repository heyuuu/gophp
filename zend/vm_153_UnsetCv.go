package zend

import "github.com/heyuuu/gophp/zend/types"

func ZEND_UNSET_CV_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ *types.Zval = opline.Op1()
	if var_.IsRefcounted() {
		var garbage *types.ZendRefcounted = var_.GetCounted()
		var_.SetUndef()
		if garbage.DelRefcount() == 0 {
			RcDtorFunc(garbage)
		} else {
			//GcCheckPossibleRoot(garbage)
		}
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	} else {
		var_.SetUndef()
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
