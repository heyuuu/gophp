package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_DIM_W_SPEC_VAR_CV_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_TMPVAR_CV_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_UNSET_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	zend_fetch_dimension_address_UNSET(container, EX_VAR(opline.GetOp2().GetVar()), IS_CV, opline, executeData)
	{
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_W_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_VAR, property, IS_CV, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_FETCH_OBJ_FLAGS) }, nil), BP_VAR_W, opline.GetExtendedValue()&ZEND_FETCH_OBJ_FLAGS, 1, opline, executeData)
	{
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_RW_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_VAR, property, IS_CV, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_RW, 0, 1, opline, executeData)
	{
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_OBJ_W_SPEC_VAR_CV_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CV_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var property *types.Zval
	var result *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_VAR, property, IS_CV, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_UNSET, 0, 1, opline, executeData)
	{
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_LIST_W_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	dim = EX_VAR(opline.GetOp2().GetVar())
	if EX_VAR(opline.GetOp1().GetVar()).GetType() != types.IS_INDIRECT && !(container.IsReference()) {
		faults.Error(faults.E_NOTICE, "Attempting to set reference to non referenceable value")
		zend_fetch_dimension_address_LIST_r(container, dim, IS_CV, opline, executeData)
	} else {
		zend_fetch_dimension_address_W(container, dim, IS_CV, opline, executeData)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
	if object.GetType() != types.IS_OBJECT {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = EG__().GetUninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
exit_assign_obj:
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if object.GetType() != types.IS_OBJECT {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = EG__().GetUninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
	ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if object.GetType() != types.IS_OBJECT {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = EG__().GetUninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
	ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
	if object.GetType() != types.IS_OBJECT {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = EG__().GetUninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
exit_assign_obj:
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

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
			dim = EX_VAR(opline.GetOp2().GetVar())

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
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
				EX_VAR(opline.GetResult().GetVar()).SetNull()
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
			dim = EX_VAR(opline.GetOp2().GetVar())

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
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
				EX_VAR(opline.GetResult().GetVar()).SetNull()
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
			dim = EX_VAR(opline.GetOp2().GetVar())

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
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
				EX_VAR(opline.GetResult().GetVar()).SetNull()
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
			dim = EX_VAR(opline.GetOp2().GetVar())

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
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
				EX_VAR(opline.GetResult().GetVar()).SetNull()
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
func ZEND_ASSIGN_SPEC_VAR_CV_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())

		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_SPEC_VAR_CV_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		EX_VAR(opline.GetResult().GetVar()).SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_REF_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var variable_ptr *types.Zval
	var value_ptr *types.Zval
	value_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp2().GetVar(), executeData)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		variable_ptr = EG__().GetUninitializedZval()
	} else if EX_VAR(opline.GetOp1().GetVar()).GetType() != types.IS_INDIRECT {
		faults.ThrowError(nil, "Cannot assign by reference to an array dimension of an object")
		variable_ptr = EG__().GetUninitializedZval()
	} else if value_ptr.IsError() {
		variable_ptr = EG__().GetUninitializedZval()
	} else if opline.GetExtendedValue() == ZEND_RETURNS_FUNCTION && !(value_ptr.IsReference()) {
		variable_ptr = ZendWrongAssignToVariableReference(variable_ptr, value_ptr, opline, executeData)
	} else {
		ZendAssignToVariableReference(variable_ptr, value_ptr)
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), variable_ptr)
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CV_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)

	{

		{
			ZendAssignToPropertyReferenceVarVar(container, property, value_ptr, opline, executeData)
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CV_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), executeData)

	{

		{
			ZendAssignToPropertyReferenceVarVar(container, property, value_ptr, opline, executeData)
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = EX_VAR(opline.GetOp1().GetVar()).GetCe()
	}

	{
		function_name = EX_VAR(opline.GetOp2().GetVar())
		{
			if function_name.GetType() != types.IS_STRING {
				for {
					if function_name.IsReference() {
						function_name = types.Z_REFVAL_P(function_name)
						if function_name.IsString() {
							break
						}
					} else if function_name.IsUndef() {
						ZVAL_UNDEFINED_OP2()
						if EG__().GetException() != nil {
							HANDLE_EXCEPTION()
						}
					}
					faults.ThrowError(nil, "Function name must be a string")
					HANDLE_EXCEPTION()
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), b.CondF1(IS_CV == IS_CONST, func() *types.Zval { return RT_CONSTANT(opline, opline.GetOp2()) + 1 }, nil))
		}
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			HANDLE_EXCEPTION()
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}

	if !fbc.IsStatic() {
		if executeData.GetThis().IsObject() && InstanceofFunction(types.Z_OBJCE(executeData.GetThis()), ce) != 0 {
			ce = (*types.ClassEntry)(executeData.GetThis().GetObj())
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:
		/* previous opcode is ZEND_FETCH_CLASS */

		call_info = ZEND_CALL_NESTED_FUNCTION
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
		if expr_ptr.IsReference() {
			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	} else {
		expr_ptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	}
	{
		var offset *types.Zval = EX_VAR(opline.GetOp2().GetVar())
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
			{
				if types.HandleNumericStr(str.GetStr(), &hval) {
					goto num_index
				}
			}
		str_index:
			EX_VAR(opline.GetResult().GetVar()).GetArr().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.GetLval()
		num_index:
			EX_VAR(opline.GetResult().GetVar()).GetArr().IndexUpdate(hval, expr_ptr)
		} else if offset.IsReference() {
			offset = types.Z_REFVAL_P(offset)
			goto add_again
		} else if offset.IsNull() {
			str = types.ZSTR_EMPTY_ALLOC()
			goto str_index
		} else if offset.IsDouble() {
			hval = DvalToLval(offset.GetDval())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = types.Z_RES_HANDLE_P(offset)
			goto num_index
		} else if offset.IsUndef() {
			ZVAL_UNDEFINED_OP2()
			str = types.ZSTR_EMPTY_ALLOC()
			goto str_index
		} else {
			ZendIllegalOffset()
			ZvalPtrDtorNogc(expr_ptr)
		}
	}

	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_ARRAY_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = EX_VAR(opline.GetResult().GetVar())
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			types.ZendHashRealInitMixed(array.GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_CV_HANDLER(executeData)
	}

}
func ZEND_UNSET_DIM_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	var hval ZendUlong
	var key *types.String
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	offset = EX_VAR(opline.GetOp2().GetVar())
	for {
		if container.IsArray() {
			var ht *types.Array
		unset_dim_array:
			types.SEPARATE_ARRAY(container)
			ht = container.GetArr()
		offset_again:
			if offset.IsString() {
				key = offset.GetStr()
				{
					if types.HandleNumericStr(key.GetStr(), &hval) {
						goto num_index_dim
					}
				}
			str_index_dim:
				if ht == EG__().GetSymbolTable() {
					ZendDeleteGlobalVariable(key)
				} else {
					types.ZendHashDel(ht, key.GetStr())
				}
			} else if offset.IsLong() {
				hval = offset.GetLval()
			num_index_dim:
				types.ZendHashIndexDel(ht, hval)
			} else if offset.IsReference() {
				offset = types.Z_REFVAL_P(offset)
				goto offset_again
			} else if offset.IsDouble() {
				hval = DvalToLval(offset.GetDval())
				goto num_index_dim
			} else if offset.IsNull() {
				key = types.ZSTR_EMPTY_ALLOC()
				goto str_index_dim
			} else if offset.IsFalse() {
				hval = 0
				goto num_index_dim
			} else if offset.IsTrue() {
				hval = 1
				goto num_index_dim
			} else if offset.IsResource() {
				hval = types.Z_RES_HANDLE_P(offset)
				goto num_index_dim
			} else if offset.IsUndef() {
				ZVAL_UNDEFINED_OP2()
				key = types.ZSTR_EMPTY_ALLOC()
				goto str_index_dim
			} else {
				faults.Error(faults.E_WARNING, "Illegal offset type in unset")
			}
			break
		} else if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto unset_dim_array
			}
		}
		if container.IsUndef() {
			container = ZVAL_UNDEFINED_OP1()
		}
		if offset.IsUndef() {
			offset = ZVAL_UNDEFINED_OP2()
		}
		if container.IsObject() {
			types.Z_OBJ_HT_P(container).GetUnsetDimension()(container, offset)
		} else if container.IsString() {
			faults.ThrowError(nil, "Cannot unset string offsets")
		}
		break
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_UNSET_OBJ_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	for {
		if container.GetType() != types.IS_OBJECT {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.GetType() != types.IS_OBJECT {
					if container.IsUndef() {
						ZVAL_UNDEFINED_OP1()
					}
					break
				}
			} else {
				break
			}
		}
		types.Z_OBJ_HT_P(container).GetUnsetProperty()(container, offset, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
		break
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
