package zend

func ZEND_STRLEN_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = RT_CONSTANT(opline, opline.GetOp1())
	if value.IsString() {
		opline.Result().SetLong(value.GetStr().GetLen())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		var strict types.ZendBool
		strict = executeData.IsCallUseStrictTypes()
		for {
			if strict == 0 {
				var str *types.String
				var tmp types.Zval
				types.ZVAL_COPY(&tmp, value)
				if zpp.ZendParseArgStrWeak(&tmp, &str) != 0 {
					opline.Result().SetLong(str.GetLen())
					ZvalPtrDtor(&tmp)
					break
				}
				ZvalPtrDtor(&tmp)
			}
			if EG__().GetException() == nil {
				faults.InternalTypeError(strict, "strlen() expects parameter 1 to be string, %s given", types.ZendGetTypeByConst(value.GetType()))
			}
			opline.Result().SetNull()
			break
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_STRLEN_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var free_op1 ZendFreeOp
	value = opline.getZvalPtrVar1(&free_op1)
	if value.IsString() {
		opline.Result().SetLong(value.GetStr().GetLen())
		ZvalPtrDtorNogc(free_op1)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		var strict types.ZendBool
		if value.IsReference() {
			value = types.Z_REFVAL_P(value)
			if value.IsString() {
				opline.Result().SetLong(value.GetStr().GetLen())
				ZvalPtrDtorNogc(free_op1)
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		}
		if value.IsUndef() {
			value = ZVAL_UNDEFINED_OP1()
		}
		strict = executeData.IsCallUseStrictTypes()
		for {
			if strict == 0 {
				var str *types.String
				var tmp types.Zval
				types.ZVAL_COPY(&tmp, value)
				if zpp.ZendParseArgStrWeak(&tmp, &str) != 0 {
					opline.Result().SetLong(str.GetLen())
					ZvalPtrDtor(&tmp)
					break
				}
				ZvalPtrDtor(&tmp)
			}
			if EG__().GetException() == nil {
				faults.InternalTypeError(strict, "strlen() expects parameter 1 to be string, %s given", types.ZendGetTypeByConst(value.GetType()))
			}
			opline.Result().SetNull()
			break
		}
	}
	ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_STRLEN_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = opline.Op1()
	if value.IsString() {
		opline.Result().SetLong(value.GetStr().GetLen())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		var strict types.ZendBool
		if value.IsReference() {
			value = types.Z_REFVAL_P(value)
			if value.IsString() {
				opline.Result().SetLong(value.GetStr().GetLen())
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		}
		if value.IsUndef() {
			value = ZVAL_UNDEFINED_OP1()
		}
		strict = executeData.IsCallUseStrictTypes()
		for {
			if strict == 0 {
				var str *types.String
				var tmp types.Zval
				types.ZVAL_COPY(&tmp, value)
				if zpp.ZendParseArgStrWeak(&tmp, &str) != 0 {
					opline.Result().SetLong(str.GetLen())
					ZvalPtrDtor(&tmp)
					break
				}
				ZvalPtrDtor(&tmp)
			}
			if EG__().GetException() == nil {
				faults.InternalTypeError(strict, "strlen() expects parameter 1 to be string, %s given", types.ZendGetTypeByConst(value.GetType()))
			}
			opline.Result().SetNull()
			break
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
