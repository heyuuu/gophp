package zend

func ZEND_ASSIGN_DIM_OP_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data1 ZendFreeOp
	var var_ptr *types.Zval
	var value *types.Zval
	var container *types.Zval
	var dim *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if container.IsArray() {
	assign_dim_op_array:
		types.SEPARATE_ARRAY(container)
	assign_dim_op_new_array:
		dim = RT_CONSTANT(opline, opline.GetOp2())

		{
			{
				var_ptr = zend_fetch_dimension_address_inner_RW_CONST(container.GetArr(), dim, executeData)
			}

			if var_ptr == nil {
				goto assign_dim_op_ret_null
			}
		}
		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
		for {
			if var_ptr.IsReference() {
				var ref *types.ZendReference = var_ptr.GetRef()
				var_ptr = types.Z_REFVAL_P(var_ptr)
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), var_ptr)
		}
		FREE_OP(free_op_data1)
	} else {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto assign_dim_op_array
			}
		}
		dim = RT_CONSTANT(opline, opline.GetOp2())
		if container.IsObject() {
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			ZendBinaryAssignOpObjDim(container, dim, opline, executeData)
		} else if container.GetType() <= types.IS_FALSE {
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			container.SetArray(types.NewArray(8))
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, executeData)
		assign_dim_op_ret_null:
			FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_OP_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var free_op_data1 ZendFreeOp
	var var_ptr *types.Zval
	var value *types.Zval
	var container *types.Zval
	var dim *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if container.IsArray() {
	assign_dim_op_array:
		types.SEPARATE_ARRAY(container)
	assign_dim_op_new_array:
		dim = opline.Op2Ptr(&free_op2)

		{

			{
				var_ptr = zend_fetch_dimension_address_inner_RW(container.GetArr(), dim, executeData)
			}
			if var_ptr == nil {
				goto assign_dim_op_ret_null
			}
		}
		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
		for {
			if var_ptr.IsReference() {
				var ref *types.ZendReference = var_ptr.GetRef()
				var_ptr = types.Z_REFVAL_P(var_ptr)
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), var_ptr)
		}
		FREE_OP(free_op_data1)
	} else {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto assign_dim_op_array
			}
		}
		dim = opline.Op2Ptr(&free_op2)
		if container.IsObject() {
			ZendBinaryAssignOpObjDim(container, dim, opline, executeData)
		} else if container.GetType() <= types.IS_FALSE {
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			container.SetArray(types.NewArray(8))
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, executeData)
		assign_dim_op_ret_null:
			FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_OP_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data1 ZendFreeOp
	var var_ptr *types.Zval
	var value *types.Zval
	var container *types.Zval
	var dim *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if container.IsArray() {
	assign_dim_op_array:
		types.SEPARATE_ARRAY(container)
	assign_dim_op_new_array:
		dim = nil
		{
			var_ptr = container.GetArr().NextIndexInsert(EG__().GetUninitializedZval())
			if var_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_op_ret_null
			}
		}

		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
		for {
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), var_ptr)
		}
		FREE_OP(free_op_data1)
	} else {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto assign_dim_op_array
			}
		}
		dim = nil
		if container.IsObject() {
			ZendBinaryAssignOpObjDim(container, dim, opline, executeData)
		} else if container.GetType() <= types.IS_FALSE {
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			container.SetArray(types.NewArray(8))
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, executeData)
		assign_dim_op_ret_null:
			FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_OP_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data1 ZendFreeOp
	var var_ptr *types.Zval
	var value *types.Zval
	var container *types.Zval
	var dim *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if container.IsArray() {
	assign_dim_op_array:
		types.SEPARATE_ARRAY(container)
	assign_dim_op_new_array:
		dim = opline.Op2()

		{

			{
				var_ptr = zend_fetch_dimension_address_inner_RW(container.GetArr(), dim, executeData)
			}
			if var_ptr == nil {
				goto assign_dim_op_ret_null
			}
		}
		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
		for {
			if var_ptr.IsReference() {
				var ref *types.ZendReference = var_ptr.GetRef()
				var_ptr = types.Z_REFVAL_P(var_ptr)
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), var_ptr)
		}
		FREE_OP(free_op_data1)
	} else {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto assign_dim_op_array
			}
		}
		dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
		if container.IsObject() {
			ZendBinaryAssignOpObjDim(container, dim, opline, executeData)
		} else if container.GetType() <= types.IS_FALSE {
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			container.SetArray(types.NewArray(8))
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, executeData)
		assign_dim_op_ret_null:
			FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_OP_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data1 ZendFreeOp
	var var_ptr *types.Zval
	var value *types.Zval
	var container *types.Zval
	var dim *types.Zval
	container = opline.Op1()
	if container.IsArray() {
	assign_dim_op_array:
		types.SEPARATE_ARRAY(container)
	assign_dim_op_new_array:
		dim = RT_CONSTANT(opline, opline.GetOp2())

		{
			{
				var_ptr = zend_fetch_dimension_address_inner_RW_CONST(container.GetArr(), dim, executeData)
			}

			if var_ptr == nil {
				goto assign_dim_op_ret_null
			}
		}
		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
		for {
			if var_ptr.IsReference() {
				var ref *types.ZendReference = var_ptr.GetRef()
				var_ptr = types.Z_REFVAL_P(var_ptr)
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), var_ptr)
		}
		FREE_OP(free_op_data1)
	} else {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto assign_dim_op_array
			}
		}
		dim = RT_CONSTANT(opline, opline.GetOp2())
		if container.IsObject() {
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			ZendBinaryAssignOpObjDim(container, dim, opline, executeData)
		} else if container.GetType() <= types.IS_FALSE {
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			container.SetArray(types.NewArray(8))
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, executeData)
		assign_dim_op_ret_null:
			FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_OP_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data1 ZendFreeOp
	var var_ptr *types.Zval
	var value *types.Zval
	var container *types.Zval
	var dim *types.Zval
	container = opline.Op1()
	if container.IsArray() {
	assign_dim_op_array:
		types.SEPARATE_ARRAY(container)
	assign_dim_op_new_array:
		dim = opline.Op2Ptr(&free_op2)

		{

			{
				var_ptr = zend_fetch_dimension_address_inner_RW(container.GetArr(), dim, executeData)
			}
			if var_ptr == nil {
				goto assign_dim_op_ret_null
			}
		}
		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
		for {
			if var_ptr.IsReference() {
				var ref *types.ZendReference = var_ptr.GetRef()
				var_ptr = types.Z_REFVAL_P(var_ptr)
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), var_ptr)
		}
		FREE_OP(free_op_data1)
	} else {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto assign_dim_op_array
			}
		}
		dim = opline.Op2Ptr(&free_op2)
		if container.IsObject() {
			ZendBinaryAssignOpObjDim(container, dim, opline, executeData)
		} else if container.GetType() <= types.IS_FALSE {
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			container.SetArray(types.NewArray(8))
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, executeData)
		assign_dim_op_ret_null:
			FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	ZvalPtrDtorNogc(free_op2)
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_OP_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data1 ZendFreeOp
	var var_ptr *types.Zval
	var value *types.Zval
	var container *types.Zval
	var dim *types.Zval
	container = opline.Op1()
	if container.IsArray() {
	assign_dim_op_array:
		types.SEPARATE_ARRAY(container)
	assign_dim_op_new_array:
		dim = nil
		{
			var_ptr = container.GetArr().NextIndexInsert(EG__().GetUninitializedZval())
			if var_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_op_ret_null
			}
		}

		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
		for {
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), var_ptr)
		}
		FREE_OP(free_op_data1)
	} else {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto assign_dim_op_array
			}
		}
		dim = nil
		if container.IsObject() {
			ZendBinaryAssignOpObjDim(container, dim, opline, executeData)
		} else if container.GetType() <= types.IS_FALSE {
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			container.SetArray(types.NewArray(8))
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, executeData)
		assign_dim_op_ret_null:
			FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_OP_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data1 ZendFreeOp
	var var_ptr *types.Zval
	var value *types.Zval
	var container *types.Zval
	var dim *types.Zval
	container = opline.Op1()
	if container.IsArray() {
	assign_dim_op_array:
		types.SEPARATE_ARRAY(container)
	assign_dim_op_new_array:
		dim = opline.Op2()

		{

			{
				var_ptr = zend_fetch_dimension_address_inner_RW(container.GetArr(), dim, executeData)
			}
			if var_ptr == nil {
				goto assign_dim_op_ret_null
			}
		}
		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
		for {
			if var_ptr.IsReference() {
				var ref *types.ZendReference = var_ptr.GetRef()
				var_ptr = types.Z_REFVAL_P(var_ptr)
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), var_ptr)
		}
		FREE_OP(free_op_data1)
	} else {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto assign_dim_op_array
			}
		}
		dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
		if container.IsObject() {
			ZendBinaryAssignOpObjDim(container, dim, opline, executeData)
		} else if container.GetType() <= types.IS_FALSE {
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			container.SetArray(types.NewArray(8))
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, executeData)
		assign_dim_op_ret_null:
			FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
