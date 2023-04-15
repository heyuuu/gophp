package zend

func ZEND_COALESCE_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var ref *types.Zval = nil
	value = opline.Const1()
	if value.GetType() > types.IS_NULL {
		var result *types.Zval = opline.Result()
		result.CopyValueFrom(value)
		{

			// result.TryAddRefcount()

		}

		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_COALESCE_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var ref *types.Zval = nil
	value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if value.IsReference() {
		{
			ref = value
		}
		value = types.Z_REFVAL_P(value)
	}
	if value.GetType() > types.IS_NULL {
		var result *types.Zval = opline.Result()
		result.CopyValueFrom(value)

		{

			// result.TryAddRefcount()

		}

		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_COALESCE_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var ref *types.Zval = nil
	value = opline.Op1Ptr(&free_op1)
	if value.IsReference() {
		{
			ref = value
		}
		value = types.Z_REFVAL_P(value)
	}
	if value.GetType() > types.IS_NULL {
		var result *types.Zval = opline.Result()
		result.CopyValueFrom(value)

		{

			// result.TryAddRefcount()

		}

		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_COALESCE_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var ref *types.Zval = nil
	value = opline.Op1()
	if value.IsReference() {
		{
			ref = value
		}
		value = types.Z_REFVAL_P(value)
	}
	if value.GetType() > types.IS_NULL {
		var result *types.Zval = opline.Result()
		result.CopyValueFrom(value)

		{

			// result.TryAddRefcount()

		}

		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
