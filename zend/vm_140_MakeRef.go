package zend

func ZEND_MAKE_REF_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval = opline.GetOp1Zval()
	{
		if op1.IsUndef() {
			op1.SetNewEmptyRef()
			op1.SetRefcount(2)
			types.Z_REFVAL_P(op1).SetNull()
			opline.GetResultZval().SetReference(op1.GetRef())
		} else {
			if op1.IsReference() {
				op1.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(op1, 2)
			}
			opline.GetResultZval().SetReference(op1.GetRef())
		}
	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_MAKE_REF_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval = opline.GetOp1Zval()
	{
		if op1.IsUndef() {
			op1.SetNewEmptyRef()
			op1.SetRefcount(2)
			types.Z_REFVAL_P(op1).SetNull()
			opline.GetResultZval().SetReference(op1.GetRef())
		} else {
			if op1.IsReference() {
				op1.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(op1, 2)
			}
			opline.GetResultZval().SetReference(op1.GetRef())
		}
	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
