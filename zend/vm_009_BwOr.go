package zend

func ZEND_BW_OR_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	return zend_bw_or_helper_SPEC(op1, op2, executeData)
}
func ZEND_BW_OR_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1()
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() && op2.IsLong() {
		opline.Result().SetLong(op1.GetLval() | op2.GetLval())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_bw_or_helper_SPEC(op1, op2, executeData)
}
func ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1()
	op2 = opline.Op2()

	if op1.IsLong() && op2.IsLong() {
		opline.Result().SetLong(op1.GetLval() | op2.GetLval())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_bw_or_helper_SPEC(op1, op2, executeData)
}
