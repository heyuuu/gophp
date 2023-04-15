package zend

func ZEND_ROPE_INIT_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var rope **types.String
	var var_ *types.Zval

	/* Compiler allocates the necessary number of zval slots to keep the rope */

	rope = (**types.String)(opline.Result())
	{
		var_ = opline.Const2()
		rope[0] = var_.GetStr()

		// var_.TryAddRefcount()

	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_ROPE_INIT_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var rope **types.String
	var var_ *types.Zval

	/* Compiler allocates the necessary number of zval slots to keep the rope */

	rope = (**types.String)(opline.Result())

	{
		var_ = opline.Op2Ptr(&free_op2)
		if var_.IsString() {
			{
				rope[0] = var_.GetStr().Copy()
			}

		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			rope[0] = ZvalGetStringFunc(var_)
			// ZvalPtrDtorNogc(free_op2)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_ROPE_INIT_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var rope **types.String
	var var_ *types.Zval

	/* Compiler allocates the necessary number of zval slots to keep the rope */

	rope = (**types.String)(opline.Result())

	{
		var_ = opline.Op2()
		if var_.IsString() {
			{
				rope[0] = var_.GetStr().Copy()
			}

		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			rope[0] = ZvalGetStringFunc(var_)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
