package zend

import "github.com/heyuuu/gophp/zend/types"

func ZEND_POST_INC_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *types.Zval
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if var_ptr.IsLong() {
		opline.Result().SetLong(var_ptr.Long()())
		FastLongIncrementFunction(var_ptr)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_post_inc_helper_SPEC_VAR(executeData)
}
func ZEND_POST_INC_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var_ptr = opline.Op1()
	if var_ptr.IsLong() {
		opline.Result().SetLong(var_ptr.Long()())
		FastLongIncrementFunction(var_ptr)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_post_inc_helper_SPEC_CV(executeData)
}
