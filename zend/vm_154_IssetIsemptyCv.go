package zend

func ZEND_ISSET_ISEMPTY_CV_SPEC_CV_UNUSED_SET_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = opline.Op1()
	if value.GetType() > types.IS_NULL && (!(value.IsReference()) || types.Z_REFVAL_P(value).GetType() != types.IS_NULL) {
		ZEND_VM_SMART_BRANCH_TRUE()
		opline.Result().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		ZEND_VM_SMART_BRANCH_FALSE()
		opline.Result().SetFalse()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func ZEND_ISSET_ISEMPTY_CV_SPEC_CV_UNUSED_EMPTY_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = opline.Op1()
	var result int
	result = !(IZendIsTrue(value))
	if EG__().GetException() != nil {
		opline.Result().SetUndef()
		return 0
	}
	ZEND_VM_SMART_BRANCH(result, 0)
	types.ZVAL_BOOL(opline.Result(), result != 0)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
