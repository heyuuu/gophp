package zend

func ZEND_BOOL_NOT_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = RT_CONSTANT(opline, opline.GetOp1())
	if val.IsTrue() {
		opline.GetResultZval().SetFalse()
	} else if val.GetTypeInfo() <= types.IS_TRUE {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		opline.GetResultZval().SetTrue()
	} else {
		types.ZVAL_BOOL(opline.GetResultZval(), IZendIsTrue(val) == 0)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_BOOL_NOT_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	var free_op1 ZendFreeOp
	val = opline.getZvalPtrVar1(&free_op1)
	if val.IsTrue() {
		opline.GetResultZval().SetFalse()
	} else if val.GetTypeInfo() <= types.IS_TRUE {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		opline.GetResultZval().SetTrue()
		if orig_val_type == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1()
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	} else {
		types.ZVAL_BOOL(opline.GetResultZval(), IZendIsTrue(val) == 0)
		ZvalPtrDtorNogc(free_op1)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_BOOL_NOT_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = opline.GetOp1Zval()
	if val.IsTrue() {
		opline.GetResultZval().SetFalse()
	} else if val.GetTypeInfo() <= types.IS_TRUE {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		opline.GetResultZval().SetTrue()
		if orig_val_type == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1()
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	} else {
		types.ZVAL_BOOL(opline.GetResultZval(), IZendIsTrue(val) == 0)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
