package zend

func ZEND_GET_TYPE_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var type_ *types.String
	op1 = opline.Const1()
	type_ = types.ZendZvalGetType(op1)
	if type_ != nil {
		opline.Result().SetInternedString(type_)
	} else {
		opline.Result().SetStringVal("unknown type")
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_GET_TYPE_SPEC_TMP_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var type_ *types.String
	op1 = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	type_ = types.ZendZvalGetType(op1)
	if type_ != nil {
		opline.Result().SetInternedString(type_)
	} else {
		opline.Result().SetStringVal("unknown type")
	}
	ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_GET_TYPE_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var type_ *types.String
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	type_ = types.ZendZvalGetType(op1)
	if type_ != nil {
		opline.Result().SetInternedString(type_)
	} else {
		opline.Result().SetStringVal("unknown type")
	}
	ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_GET_TYPE_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var type_ *types.String
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	type_ = types.ZendZvalGetType(op1)
	if type_ != nil {
		opline.Result().SetInternedString(type_)
	} else {
		opline.Result().SetStringVal("unknown type")
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
