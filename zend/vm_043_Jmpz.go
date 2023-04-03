package zend

func ZEND_JMPZ_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = RT_CONSTANT(opline, opline.GetOp1())
	if val.IsTrue() {
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	if IZendIsTrue(val) != 0 {
		opline++
	} else {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	return ZEND_VM_JMP(executeData, opline)
}
func ZEND_JMPZ_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var val *types.Zval
	val = opline.getZvalPtrVar1(&free_op1)
	if val.IsTrue() {
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		if val.IsUndef() {
			ZVAL_UNDEFINED_OP1()
			if EG__().GetException() != nil {
				return 0
			}
		}
		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	if IZendIsTrue(val) != 0 {
		opline++
	} else {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_JMP(executeData, opline)
}
func ZEND_JMPZ_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = opline.Op1()
	if val.IsTrue() {
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		if val.IsUndef() {
			ZVAL_UNDEFINED_OP1()
			if EG__().GetException() != nil {
				return 0
			}
		}
		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	if IZendIsTrue(val) != 0 {
		opline++
	} else {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	return ZEND_VM_JMP(executeData, opline)
}
