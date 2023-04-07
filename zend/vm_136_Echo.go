package zend

func ZEND_ECHO_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var z *types.Zval
	z = opline.Const1()
	if z.IsString() {
		var str *types.String = z.GetStr()
		if str.GetLen() != 0 {
			ZendWrite(str.GetStr())
		}
	} else {
		var str *types.String = ZvalGetStringFunc(z)
		if str.GetLen() != 0 {
			ZendWrite(str.GetStr())
		} else {
		}

		// types.ZendStringReleaseEx(str, 0)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ECHO_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var z *types.Zval
	z = opline.Op1Ptr(&free_op1)
	if z.IsString() {
		var str *types.String = z.GetStr()
		if str.GetLen() != 0 {
			ZendWrite(str.GetStr())
		}
	} else {
		var str *types.String = ZvalGetStringFunc(z)
		if str.GetLen() != 0 {
			ZendWrite(str.GetStr())
		} else if z.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		// types.ZendStringReleaseEx(str, 0)
	}
	ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ECHO_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var z *types.Zval
	z = opline.Op1()
	if z.IsString() {
		var str *types.String = z.GetStr()
		if str.GetLen() != 0 {
			ZendWrite(str.GetStr())
		}
	} else {
		var str *types.String = ZvalGetStringFunc(z)
		if str.GetLen() != 0 {
			ZendWrite(str.GetStr())
		} else if z.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		// types.ZendStringReleaseEx(str, 0)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
