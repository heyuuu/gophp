package zend

func ZEND_FETCH_DIM_R_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = opline.Const1()
	dim = opline.Const2()

	{
		zend_fetch_dimension_address_read_R(container, dim, IS_CONST, opline, executeData)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_R_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = opline.Const1()
	dim = opline.Op2Ptr(&free_op2)

	{
		zend_fetch_dimension_address_read_R(container, dim, IS_TMP_VAR|IS_VAR, opline, executeData)
	}
	// ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_R_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = opline.Const1()
	dim = opline.Op2()

	{
		zend_fetch_dimension_address_read_R(container, dim, IS_CV, opline, executeData)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_R_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = opline.Op1Ptr(&free_op1)
	dim = opline.Const2()
	{
		if container.IsArray() {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetArr(), dim, IS_CONST, BP_VAR_R, executeData)
			types.ZVAL_COPY_DEREF(opline.Result(), value)
		} else if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto fetch_dim_r_array
			} else {
				goto fetch_dim_r_slow
			}
		} else {
		fetch_dim_r_slow:
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			zend_fetch_dimension_address_read_R_slow(container, dim, opline, executeData)
		}
	}

	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = opline.Op1Ptr(&free_op1)
	dim = opline.Op2Ptr(&free_op2)
	{
		if container.IsArray() {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetArr(), dim, IS_TMP_VAR|IS_VAR, BP_VAR_R, executeData)
			types.ZVAL_COPY_DEREF(opline.Result(), value)
		} else if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto fetch_dim_r_array
			} else {
				goto fetch_dim_r_slow
			}
		} else {
		fetch_dim_r_slow:
			zend_fetch_dimension_address_read_R_slow(container, dim, opline, executeData)
		}
	}

	// ZvalPtrDtorNogc(free_op2)
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_R_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = opline.Op1Ptr(&free_op1)
	dim = opline.Op2()
	{
		if container.IsArray() {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetArr(), dim, IS_CV, BP_VAR_R, executeData)
			types.ZVAL_COPY_DEREF(opline.Result(), value)
		} else if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto fetch_dim_r_array
			} else {
				goto fetch_dim_r_slow
			}
		} else {
		fetch_dim_r_slow:
			zend_fetch_dimension_address_read_R_slow(container, dim, opline, executeData)
		}
	}

	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_R_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = opline.Op1()
	dim = opline.Const2()
	{
		if container.IsArray() {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetArr(), dim, IS_CONST, BP_VAR_R, executeData)
			types.ZVAL_COPY_DEREF(opline.Result(), value)
		} else if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto fetch_dim_r_array
			} else {
				goto fetch_dim_r_slow
			}
		} else {
		fetch_dim_r_slow:
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			zend_fetch_dimension_address_read_R_slow(container, dim, opline, executeData)
		}
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_R_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = opline.Op1()
	dim = opline.Op2Ptr(&free_op2)
	{
		if container.IsArray() {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetArr(), dim, IS_TMP_VAR|IS_VAR, BP_VAR_R, executeData)
			types.ZVAL_COPY_DEREF(opline.Result(), value)
		} else if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto fetch_dim_r_array
			} else {
				goto fetch_dim_r_slow
			}
		} else {
		fetch_dim_r_slow:
			zend_fetch_dimension_address_read_R_slow(container, dim, opline, executeData)
		}
	}

	// ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_R_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = opline.Op1()
	dim = opline.Op2()
	{
		if container.IsArray() {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetArr(), dim, IS_CV, BP_VAR_R, executeData)
			types.ZVAL_COPY_DEREF(opline.Result(), value)
		} else if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto fetch_dim_r_array
			} else {
				goto fetch_dim_r_slow
			}
		} else {
		fetch_dim_r_slow:
			zend_fetch_dimension_address_read_R_slow(container, dim, opline, executeData)
		}
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
