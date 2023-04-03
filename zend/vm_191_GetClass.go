package zend

func ZEND_GET_CLASS_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()

	{
		var op1 *types.Zval
		op1 = RT_CONSTANT(opline, opline.GetOp1())
		for true {
			if op1.IsObject() {
				opline.Result().SetStringCopy(types.Z_OBJCE_P(op1).GetName())
			} else {
				faults.Error(faults.E_WARNING, "get_class() expects parameter 1 to be object, %s given", types.ZendGetTypeByConst(op1.GetType()))
				opline.Result().SetFalse()
			}
			break
		}
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
}
func ZEND_GET_CLASS_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()

	{
		var free_op1 ZendFreeOp
		var op1 *types.Zval
		op1 = opline.getZvalPtrVar1(&free_op1)
		for true {
			if op1.IsObject() {
				opline.Result().SetStringCopy(types.Z_OBJCE_P(op1).GetName())
			} else if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				continue
			} else {
				if op1.IsUndef() {
					ZVAL_UNDEFINED_OP1()
				}
				faults.Error(faults.E_WARNING, "get_class() expects parameter 1 to be object, %s given", types.ZendGetTypeByConst(op1.GetType()))
				opline.Result().SetFalse()
			}
			break
		}
		ZvalPtrDtorNogc(free_op1)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
}
func ZEND_GET_CLASS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	{
		if !(executeData.GetFunc().common.scope) {
			faults.Error(faults.E_WARNING, "get_class() called without object from outside a class")
			opline.Result().SetFalse()
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		} else {
			opline.Result().SetStringCopy(executeData.GetFunc().common.scope.name)
			return ZEND_VM_NEXT_OPCODE(executeData, opline)
		}
	}

}
func ZEND_GET_CLASS_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()

	{
		var op1 *types.Zval
		op1 = opline.Op1()
		for true {
			if op1.IsObject() {
				opline.Result().SetStringCopy(types.Z_OBJCE_P(op1).GetName())
			} else if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				continue
			} else {
				if op1.IsUndef() {
					ZVAL_UNDEFINED_OP1()
				}
				faults.Error(faults.E_WARNING, "get_class() expects parameter 1 to be object, %s given", types.ZendGetTypeByConst(op1.GetType()))
				opline.Result().SetFalse()
			}
			break
		}
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
}
