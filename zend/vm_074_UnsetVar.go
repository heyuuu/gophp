package zend

func ZEND_UNSET_VAR_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	varname = opline.Const1()
	{
		name = varname.GetStr()
	}

	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	types.ZendHashDelInd(target_symbol_table, name.GetStr())
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_VAR_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	var free_op1 ZendFreeOp
	varname = opline.Op1Ptr(&free_op1)

	if varname.IsString() {
		name = varname.GetStr()
		tmp_name = nil
	} else {
		if varname.IsUndef() {
			varname = ZVAL_UNDEFINED_OP1()
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			ZvalPtrDtorNogc(free_op1)
			return 0
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	types.ZendHashDelInd(target_symbol_table, name.GetStr())
	{
		ZendTmpStringRelease(tmp_name)
	}
	ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_VAR_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	varname = opline.Op1()

	if varname.IsString() {
		name = varname.GetStr()
		tmp_name = nil
	} else {
		if varname.IsUndef() {
			varname = ZVAL_UNDEFINED_OP1()
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			return 0
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	types.ZendHashDelInd(target_symbol_table, name.GetStr())
	{
		ZendTmpStringRelease(tmp_name)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
