package zend

func ZEND_ROPE_ADD_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var rope **types.String
	var var_ *types.Zval

	/* op1 and result are the same */

	rope = (**types.String)(opline.Op1())
	{
		var_ = opline.Const2()
		rope[opline.GetExtendedValue()] = var_.GetStr()

		// var_.TryAddRefcount()

	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_ROPE_ADD_SPEC_TMP_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var rope **types.String
	var var_ *types.Zval

	/* op1 and result are the same */

	rope = (**types.String)(opline.Op1())

	{
		var_ = opline.Op2Ptr(&free_op2)
		if var_.IsString() {
			{
				rope[opline.GetExtendedValue()] = var_.GetStr().Copy()
			}

		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			rope[opline.GetExtendedValue()] = ZvalGetStringFunc(var_)
			// ZvalPtrDtorNogc(free_op2)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_ROPE_ADD_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var rope **types.String
	var var_ *types.Zval

	/* op1 and result are the same */

	rope = (**types.String)(opline.Op1())

	{
		var_ = opline.Op2()
		if var_.IsString() {
			{
				rope[opline.GetExtendedValue()] = var_.GetStr().Copy()
			}

		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			rope[opline.GetExtendedValue()] = ZvalGetStringFunc(var_)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
