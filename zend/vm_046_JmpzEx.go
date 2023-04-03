package zend

func ZEND_JMPZ_EX_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	var ret int
	val = RT_CONSTANT(opline, opline.GetOp1())
	if val.IsTrue() {
		opline.GetResultZval().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		opline.GetResultZval().SetFalse()
		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ret = IZendIsTrue(val)
	if ret != 0 {
		opline.GetResultZval().SetTrue()
		opline++
	} else {
		opline.GetResultZval().SetFalse()
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	return ZEND_VM_JMP(executeData, opline)
}
func ZEND_JMPZ_EX_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var val *types.Zval
	var ret int
	val = opline.getZvalPtrVar1(&free_op1)
	if val.IsTrue() {
		opline.GetResultZval().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		opline.GetResultZval().SetFalse()
		if val.IsUndef() {
			ZVAL_UNDEFINED_OP1()
			if EG__().GetException() != nil {
				return 0
			}
		}
		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ret = IZendIsTrue(val)
	ZvalPtrDtorNogc(free_op1)
	if ret != 0 {
		opline.GetResultZval().SetTrue()
		opline++
	} else {
		opline.GetResultZval().SetFalse()
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	return ZEND_VM_JMP(executeData, opline)
}
func ZEND_JMPZ_EX_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	var ret int
	val = opline.GetOp1Zval()
	if val.IsTrue() {
		opline.GetResultZval().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		opline.GetResultZval().SetFalse()
		if val.IsUndef() {
			ZVAL_UNDEFINED_OP1()
			if EG__().GetException() != nil {
				return 0
			}
		}
		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ret = IZendIsTrue(val)
	if ret != 0 {
		opline.GetResultZval().SetTrue()
		opline++
	} else {
		opline.GetResultZval().SetFalse()
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	return ZEND_VM_JMP(executeData, opline)
}
