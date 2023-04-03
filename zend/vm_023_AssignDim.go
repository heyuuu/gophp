package zend

func ZEND_ASSIGN_DIM_SPEC_VAR_CONST_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = RT_CONSTANT(opline, opline.GetOp2())
			{
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetArr(), dim, executeData)
			}

			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = RT_CONSTANT(opline, opline.GetOp2())
			value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = RT_CONSTANT(opline, opline.GetOp2())
				value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = RT_CONSTANT(opline, opline.GetOp2())
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = RT_CONSTANT(opline, opline.GetOp2())
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_CONST_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = RT_CONSTANT(opline, opline.GetOp2())
			{
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetArr(), dim, executeData)
			}

			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = RT_CONSTANT(opline, opline.GetOp2())
			value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = RT_CONSTANT(opline, opline.GetOp2())
				value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = RT_CONSTANT(opline, opline.GetOp2())
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = RT_CONSTANT(opline, opline.GetOp2())
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_CONST_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = RT_CONSTANT(opline, opline.GetOp2())
			{
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetArr(), dim, executeData)
			}

			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = RT_CONSTANT(opline, opline.GetOp2())
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = RT_CONSTANT(opline, opline.GetOp2())
				value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = RT_CONSTANT(opline, opline.GetOp2())
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = RT_CONSTANT(opline, opline.GetOp2())
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_CONST_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = RT_CONSTANT(opline, opline.GetOp2())
			{
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetArr(), dim, executeData)
			}

			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = RT_CONSTANT(opline, opline.GetOp2())
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = RT_CONSTANT(opline, opline.GetOp2())
				value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = RT_CONSTANT(opline, opline.GetOp2())
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = RT_CONSTANT(opline, opline.GetOp2())
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.getZvalPtrVar2(&free_op2)

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = opline.getZvalPtrVar2(&free_op2)
			value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = opline.getZvalPtrVar2(&free_op2)
				value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = opline.getZvalPtrVar2(&free_op2)
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = opline.getZvalPtrVar2(&free_op2)
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	{
		ZvalPtrDtorNogc(free_op2)
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.getZvalPtrVar2(&free_op2)

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = opline.getZvalPtrVar2(&free_op2)
			value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = opline.getZvalPtrVar2(&free_op2)
				value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = opline.getZvalPtrVar2(&free_op2)
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = opline.getZvalPtrVar2(&free_op2)
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	{
		ZvalPtrDtorNogc(free_op2)
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.getZvalPtrVar2(&free_op2)

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = opline.getZvalPtrVar2(&free_op2)
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = opline.getZvalPtrVar2(&free_op2)
				value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = opline.getZvalPtrVar2(&free_op2)
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = opline.getZvalPtrVar2(&free_op2)
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	{
		ZvalPtrDtorNogc(free_op2)
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.getZvalPtrVar2(&free_op2)

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = opline.getZvalPtrVar2(&free_op2)
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = opline.getZvalPtrVar2(&free_op2)
				value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = opline.getZvalPtrVar2(&free_op2)
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = opline.getZvalPtrVar2(&free_op2)
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	{
		ZvalPtrDtorNogc(free_op2)
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
		types.SEPARATE_ARRAY(object_ptr)
		{
			variable_ptr = object_ptr.GetArr().NextIndexInsert(value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else {

				value.TryAddRefcount()

			}
		}

		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = nil
			value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {
			{
				ZendUseNewElementForString()
				UNDEF_RESULT()
			}

		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = nil
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)
		{
			{
				value = types.ZVAL_DEREF(value)
			}
			variable_ptr = object_ptr.GetArr().NextIndexInsert(value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else {

				value.TryAddRefcount()

			}

		}

		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = nil
			value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {
			{
				ZendUseNewElementForString()
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			}

		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = nil
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)
		{
			{
				value = types.ZVAL_DEREF(value)
			}
			variable_ptr = object_ptr.GetArr().NextIndexInsert(value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else {

				value.TryAddRefcount()

			}

		}

		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = nil
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {
			{
				ZendUseNewElementForString()
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			}

		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = nil
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
		types.SEPARATE_ARRAY(object_ptr)
		{
			{
				value = types.ZVAL_DEREF(value)
			}
			variable_ptr = object_ptr.GetArr().NextIndexInsert(value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else {

				value.TryAddRefcount()

			}

		}

		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = nil
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {
			{
				ZendUseNewElementForString()
				UNDEF_RESULT()
			}

		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = nil
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.Op2()

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
			value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.Op2()

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
			value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.Op2()

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.Op2()

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = RT_CONSTANT(opline, opline.GetOp2())
			{
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetArr(), dim, executeData)
			}

			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = RT_CONSTANT(opline, opline.GetOp2())
			value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = RT_CONSTANT(opline, opline.GetOp2())
				value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = RT_CONSTANT(opline, opline.GetOp2())
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = RT_CONSTANT(opline, opline.GetOp2())
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = RT_CONSTANT(opline, opline.GetOp2())
			{
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetArr(), dim, executeData)
			}

			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = RT_CONSTANT(opline, opline.GetOp2())
			value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = RT_CONSTANT(opline, opline.GetOp2())
				value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = RT_CONSTANT(opline, opline.GetOp2())
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = RT_CONSTANT(opline, opline.GetOp2())
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = RT_CONSTANT(opline, opline.GetOp2())
			{
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetArr(), dim, executeData)
			}

			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = RT_CONSTANT(opline, opline.GetOp2())
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = RT_CONSTANT(opline, opline.GetOp2())
				value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = RT_CONSTANT(opline, opline.GetOp2())
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = RT_CONSTANT(opline, opline.GetOp2())
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = RT_CONSTANT(opline, opline.GetOp2())
			{
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetArr(), dim, executeData)
			}

			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = RT_CONSTANT(opline, opline.GetOp2())
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = RT_CONSTANT(opline, opline.GetOp2())
				value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = RT_CONSTANT(opline, opline.GetOp2())
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = RT_CONSTANT(opline, opline.GetOp2())
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.getZvalPtrVar2(&free_op2)

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = opline.getZvalPtrVar2(&free_op2)
			value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = opline.getZvalPtrVar2(&free_op2)
				value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = opline.getZvalPtrVar2(&free_op2)
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = opline.getZvalPtrVar2(&free_op2)
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	{
		ZvalPtrDtorNogc(free_op2)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.getZvalPtrVar2(&free_op2)

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = opline.getZvalPtrVar2(&free_op2)
			value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = opline.getZvalPtrVar2(&free_op2)
				value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = opline.getZvalPtrVar2(&free_op2)
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = opline.getZvalPtrVar2(&free_op2)
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	{
		ZvalPtrDtorNogc(free_op2)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.getZvalPtrVar2(&free_op2)

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = opline.getZvalPtrVar2(&free_op2)
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = opline.getZvalPtrVar2(&free_op2)
				value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = opline.getZvalPtrVar2(&free_op2)
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = opline.getZvalPtrVar2(&free_op2)
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	{
		ZvalPtrDtorNogc(free_op2)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.getZvalPtrVar2(&free_op2)

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = opline.getZvalPtrVar2(&free_op2)
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = opline.getZvalPtrVar2(&free_op2)
				value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = opline.getZvalPtrVar2(&free_op2)
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = opline.getZvalPtrVar2(&free_op2)
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	{
		ZvalPtrDtorNogc(free_op2)
	}

	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
		types.SEPARATE_ARRAY(object_ptr)
		{
			variable_ptr = object_ptr.GetArr().NextIndexInsert(value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else {

				value.TryAddRefcount()

			}
		}

		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = nil
			value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {
			{
				ZendUseNewElementForString()
				UNDEF_RESULT()
			}

		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = nil
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)
		{
			{
				value = types.ZVAL_DEREF(value)
			}
			variable_ptr = object_ptr.GetArr().NextIndexInsert(value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else {

				value.TryAddRefcount()

			}

		}

		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = nil
			value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {
			{
				ZendUseNewElementForString()
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			}

		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = nil
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)
		{
			{
				value = types.ZVAL_DEREF(value)
			}
			variable_ptr = object_ptr.GetArr().NextIndexInsert(value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else {

				value.TryAddRefcount()

			}

		}

		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = nil
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {
			{
				ZendUseNewElementForString()
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			}

		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = nil
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
		types.SEPARATE_ARRAY(object_ptr)
		{
			{
				value = types.ZVAL_DEREF(value)
			}
			variable_ptr = object_ptr.GetArr().NextIndexInsert(value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else {

				value.TryAddRefcount()

			}

		}

		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = nil
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {
			{
				ZendUseNewElementForString()
				UNDEF_RESULT()
			}

		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = nil
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.Op2()

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
			value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.Op2()

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
			value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.Op2()

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = opline.Op1()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = opline.Op2()

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.Result(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.Result().SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
