package zend

func ZEND_BOOL_NOT_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = opline.Const1()
	if val.IsTrue() {
		opline.Result().SetFalse()
	} else if val.GetTypeInfo() <= types.IS_TRUE {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		opline.Result().SetTrue()
	} else {
		types.ZVAL_BOOL(opline.Result(), IZendIsTrue(val) == 0)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_BOOL_NOT_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	var free_op1 ZendFreeOp
	val = opline.Op1Ptr(&free_op1)
	if val.IsTrue() {
		opline.Result().SetFalse()
	} else if val.GetTypeInfo() <= types.IS_TRUE {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		opline.Result().SetTrue()
		if orig_val_type == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1()
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	} else {
		types.ZVAL_BOOL(opline.Result(), IZendIsTrue(val) == 0)
		ZvalPtrDtorNogc(free_op1)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_BOOL_NOT_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = opline.Op1()
	if val.IsTrue() {
		opline.Result().SetFalse()
	} else if val.GetTypeInfo() <= types.IS_TRUE {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		opline.Result().SetTrue()
		if orig_val_type == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1()
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	} else {
		types.ZVAL_BOOL(opline.Result(), IZendIsTrue(val) == 0)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
