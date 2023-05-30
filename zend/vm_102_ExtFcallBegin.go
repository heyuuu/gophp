package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_EXT_FCALL_BEGIN_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if EG__().GetNoExtensions() == 0 {
		ZendExtensions.ApplyWithArgument(LlistApplyWithArgFuncT(ZendExtensionFcallBeginHandler), executeData)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
