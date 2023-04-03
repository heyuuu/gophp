package zend

func ZEND_ASSIGN_STATIC_PROP_REF_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var prop *types.Zval
	var value_ptr *types.Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue() & ^ZEND_RETURNS_FUNCTION, BP_VAR_W, 0, opline, executeData) != types.SUCCESS {
		FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
		UNDEF_RESULT()
		return 0
	}
	value_ptr = GetZvalPtrPtr((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data, BP_VAR_W)
	if (opline+1).GetOp1Type() == IS_VAR && value_ptr.IsError() {
		prop = EG__().GetUninitializedZval()
	} else if (opline+1).GetOp1Type() == IS_VAR && (opline.GetExtendedValue()&ZEND_RETURNS_FUNCTION) != 0 && !(value_ptr.IsReference()) {
		if ZendWrongAssignToVariableReference(prop, value_ptr, opline, executeData) == nil {
			prop = EG__().GetUninitializedZval()
		}
	} else if prop_info.GetType() != 0 {
		prop = ZendAssignToTypedPropertyReference(prop_info, prop, value_ptr, executeData)
	} else {
		ZendAssignToVariableReference(prop, value_ptr)
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), prop)
	}
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
