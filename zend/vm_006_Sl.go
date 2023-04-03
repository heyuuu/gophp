package zend

func ZEND_SL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

	return zend_shift_left_helper_SPEC(op1, op2, executeData)
}
func ZEND_SL_SPEC_CONST_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = opline.GetOp2Zval()

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.GetLval() < SIZEOF_ZEND_LONG*8) != 0 {

		/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

		opline.GetResultZval().SetLong(zend_long(ZendUlong(op1.GetLval() << op2.GetLval())))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_shift_left_helper_SPEC(op1, op2, executeData)
}
func ZEND_SL_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.GetOp1Zval()
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.GetLval() < SIZEOF_ZEND_LONG*8) != 0 {

		/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

		opline.GetResultZval().SetLong(zend_long(ZendUlong(op1.GetLval() << op2.GetLval())))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_shift_left_helper_SPEC(op1, op2, executeData)
}
func ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.GetOp1Zval()
	op2 = opline.GetOp2Zval()

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.GetLval() < SIZEOF_ZEND_LONG*8) != 0 {

		/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

		opline.GetResultZval().SetLong(zend_long(ZendUlong(op1.GetLval() << op2.GetLval())))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_shift_left_helper_SPEC(op1, op2, executeData)
}
