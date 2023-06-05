package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_FE_FREE_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var var_ *types.Zval
	var opline *types.ZendOp = executeData.GetOpline()
	var_ = opline.Op1()
	if !var_.IsArray() && var_.GetFeIterIdx() != uint32-1 {
		EG__().DelArrayIterator(var_.GetFeIterIdx())
	}
	// ZvalPtrDtorNogc(var_)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
