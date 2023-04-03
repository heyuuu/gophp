package zend

func ZEND_RETURN_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types.Zval
	var return_value *types.Zval
	var free_op1 ZendFreeOp
	retval_ptr = RT_CONSTANT(opline, opline.GetOp1())
	return_value = executeData.GetReturnValue()

	if return_value == nil {
	} else {
		{
			types.ZVAL_COPY_VALUE(return_value, retval_ptr)
			{

				return_value.TryAddRefcount()

			}
		}

	}
	return zend_leave_helper_SPEC(executeData)
}
func ZEND_RETURN_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types.Zval
	var return_value *types.Zval
	var free_op1 ZendFreeOp
	retval_ptr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	return_value = executeData.GetReturnValue()
	if retval_ptr.IsUndef() {
		retval_ptr = ZVAL_UNDEFINED_OP1()
		if return_value != nil {
			return_value.SetNull()
		}
	} else if return_value == nil {
		{
			if free_op1.IsRefcounted() && free_op1.DelRefcount() == 0 {
				RcDtorFunc(free_op1.GetCounted())
			}
		}
	} else {
		{
			types.ZVAL_COPY_VALUE(return_value, retval_ptr)
		}

	}
	return zend_leave_helper_SPEC(executeData)
}
func ZEND_RETURN_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types.Zval
	var return_value *types.Zval
	var free_op1 ZendFreeOp
	retval_ptr = opline.Op1Ptr(&free_op1)
	return_value = executeData.GetReturnValue()
	if retval_ptr.IsUndef() {
		retval_ptr = ZVAL_UNDEFINED_OP1()
		if return_value != nil {
			return_value.SetNull()
		}
	} else if return_value == nil {
		{
			if free_op1.IsRefcounted() && free_op1.DelRefcount() == 0 {
				RcDtorFunc(free_op1.GetCounted())
			}
		}
	} else {
		{
			types.ZVAL_COPY_VALUE(return_value, retval_ptr)
		}

	}
	return zend_leave_helper_SPEC(executeData)
}
func ZEND_RETURN_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types.Zval
	var return_value *types.Zval
	var free_op1 ZendFreeOp
	retval_ptr = opline.Op1()
	return_value = executeData.GetReturnValue()
	if retval_ptr.IsUndef() {
		retval_ptr = ZVAL_UNDEFINED_OP1()
		if return_value != nil {
			return_value.SetNull()
		}
	} else if return_value == nil {
		{
			if free_op1.IsRefcounted() && free_op1.DelRefcount() == 0 {
				RcDtorFunc(free_op1.GetCounted())
			}
		}
	} else {
		{
			types.ZVAL_COPY_VALUE(return_value, retval_ptr)
		}

	}
	return zend_leave_helper_SPEC(executeData)
}
