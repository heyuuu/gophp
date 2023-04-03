package zend

func ZEND_ASSIGN_SPEC_VAR_CONST_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = RT_CONSTANT(opline, opline.GetOp2())
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())

		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_VAR_CONST_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = RT_CONSTANT(opline, opline.GetOp2())
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		opline.GetResultZval().SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(opline.GetResultZval(), value)
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_VAR_TMP_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)

	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())

		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_VAR_TMP_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)
		opline.GetResultZval().SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(opline.GetResultZval(), value)
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_VAR_VAR_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = opline.getZvalPtrVar2(&free_op2)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)

	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())

		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_VAR_VAR_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = opline.getZvalPtrVar2(&free_op2)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)
		opline.GetResultZval().SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(opline.GetResultZval(), value)
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_VAR_CV_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())

		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_VAR_CV_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		opline.GetResultZval().SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(opline.GetResultZval(), value)
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_CV_CONST_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var variable_ptr *types.Zval
	value = RT_CONSTANT(opline, opline.GetOp2())
	variable_ptr = opline.GetOp1Zval()
	if variable_ptr.IsError() {
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())

	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_CV_CONST_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var variable_ptr *types.Zval
	value = RT_CONSTANT(opline, opline.GetOp2())
	variable_ptr = opline.GetOp1Zval()
	if variable_ptr.IsError() {
		opline.GetResultZval().SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_CV_TMP_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = opline.GetOp1Zval()
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)

	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())

	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_CV_TMP_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = opline.GetOp1Zval()
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)
		opline.GetResultZval().SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_CV_VAR_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = opline.getZvalPtrVar2(&free_op2)
	variable_ptr = opline.GetOp1Zval()
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)

	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())

	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_CV_VAR_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = opline.getZvalPtrVar2(&free_op2)
	variable_ptr = opline.GetOp1Zval()
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)
		opline.GetResultZval().SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_CV_CV_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	variable_ptr = opline.GetOp1Zval()
	if variable_ptr.IsError() {
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())

	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_CV_CV_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	variable_ptr = opline.GetOp1Zval()
	if variable_ptr.IsError() {
		opline.GetResultZval().SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
