package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_FETCH_DIM_W_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = opline.GetOp1Zval()
	zend_fetch_dimension_address_W(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData), IS_TMP_VAR|IS_VAR, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	{
		var result *types.Zval = opline.GetResultZval()
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_RW_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = opline.GetOp1Zval()
	zend_fetch_dimension_address_RW(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData), IS_TMP_VAR|IS_VAR, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	{
		var result *types.Zval = opline.GetResultZval()
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_IS_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = opline.GetOp1Zval()
	zend_fetch_dimension_address_read_IS(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData), IS_TMP_VAR|IS_VAR, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_DIM_W_SPEC_CV_TMPVAR_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_CV_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_UNSET_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = opline.GetOp1Zval()
	zend_fetch_dimension_address_UNSET(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData), IS_TMP_VAR|IS_VAR, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	{
		var result *types.Zval = opline.GetResultZval()
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_R_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var free_op2 ZendFreeOp
	var offset *types.Zval
	var cache_slot *any = nil
	container = opline.GetOp1Zval()
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	if container.GetType() != types.IS_OBJECT {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			if offset.IsUndef() {
				ZVAL_UNDEFINED_OP2()
			}
			ZendWrongPropertyRead(offset)
			opline.GetResultZval().SetNull()
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval

	if offset.IsUndef() {
		ZVAL_UNDEFINED_OP2()
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, opline.GetResultZval())
	if retval != opline.GetResultZval() {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(opline.GetResultZval(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_W_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = opline.GetOp1Zval()
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = opline.GetResultZval()
	ZendFetchPropertyAddress(result, container, IS_CV, property, IS_TMP_VAR|IS_VAR, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_FETCH_OBJ_FLAGS) }, nil), BP_VAR_W, opline.GetExtendedValue()&ZEND_FETCH_OBJ_FLAGS, 1, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	{
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_RW_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = opline.GetOp1Zval()
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = opline.GetResultZval()
	ZendFetchPropertyAddress(result, container, IS_CV, property, IS_TMP_VAR|IS_VAR, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_RW, 0, 1, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	{
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_IS_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var free_op2 ZendFreeOp
	var offset *types.Zval
	var cache_slot *any = nil
	container = _get_zval_ptr_cv_BP_VAR_IS(opline.GetOp1().GetVar(), executeData)
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	if container.GetType() != types.IS_OBJECT {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			opline.GetResultZval().SetNull()
			goto fetch_obj_is_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_IS, cache_slot, opline.GetResultZval())
	if retval != opline.GetResultZval() {
	fetch_obj_is_copy:
		types.ZVAL_COPY_DEREF(opline.GetResultZval(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_is_finish:
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_OBJ_W_SPEC_CV_TMPVAR_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CV_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	var property *types.Zval
	var result *types.Zval
	container = opline.GetOp1Zval()
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = opline.GetResultZval()
	ZendFetchPropertyAddress(result, container, IS_CV, property, IS_TMP_VAR|IS_VAR, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_UNSET, 0, 1, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	{
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = opline.GetOp1Zval()
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
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
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}
exit_assign_obj:
	ZvalPtrDtorNogc(free_op2)

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = opline.GetOp1Zval()
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
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
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}
	ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:
	ZvalPtrDtorNogc(free_op2)

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = opline.GetOp1Zval()
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
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
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}
	ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:
	ZvalPtrDtorNogc(free_op2)

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = opline.GetOp1Zval()
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
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
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.GetResultZval(), value)
	}
exit_assign_obj:
	ZvalPtrDtorNogc(free_op2)

	/* assign_obj has two opcodes! */

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
	object_ptr = opline.GetOp1Zval()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.GetResultZval(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
			value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
				value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.GetResultZval().SetNull()
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
	object_ptr = opline.GetOp1Zval()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.GetResultZval(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
			value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
				value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
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
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.GetResultZval().SetNull()
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
	object_ptr = opline.GetOp1Zval()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.GetResultZval(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.IsString() {

			{
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
				value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
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
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
		assign_dim_error:
			ZvalPtrDtorNogc(EX_VAR((opline + 1).GetOp1().GetVar()))
			if RETURN_VALUE_USED(opline) {
				opline.GetResultZval().SetNull()
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
	object_ptr = opline.GetOp1Zval()
	orig_object_ptr = object_ptr
	if object_ptr.IsArray() {
	try_assign_dim_array:
		value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
		types.SEPARATE_ARRAY(object_ptr)

		{
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

			{
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetArr(), dim, executeData)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, IS_CV, executeData.IsCallUseStrictTypes())
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(opline.GetResultZval(), value)
		}
	} else {
		if object_ptr.IsReference() {
			object_ptr = types.Z_REFVAL_P(object_ptr)
			if object_ptr.IsArray() {
				goto try_assign_dim_array
			}
		}
		if object_ptr.IsObject() {
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
			ZendAssignToObjectDim(object_ptr, dim, value, opline, executeData)
		} else if object_ptr.IsString() {

			{
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
				value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), executeData)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, executeData)
			}
		} else if object_ptr.GetType() <= types.IS_FALSE {
			if orig_object_ptr.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(orig_object_ptr.GetRef()) && ZendVerifyRefArrayAssignable(orig_object_ptr.GetRef()) == 0 {
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
				UNDEF_RESULT()
			} else {
				object_ptr.SetArray(types.NewArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				opline.GetResultZval().SetNull()
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
func ZEND_ASSIGN_OBJ_REF_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = opline.GetOp1Zval()
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)

	{

		{
			ZendAssignToPropertyReferenceVarVar(container, property, value_ptr, opline, executeData)
		}
	}
	ZvalPtrDtorNogc(free_op2)
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = opline.GetOp1Zval()
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), executeData)

	{

		{
			ZendAssignToPropertyReferenceVarVar(container, property, value_ptr, opline, executeData)
		}
	}
	ZvalPtrDtorNogc(free_op2)
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_FAST_CONCAT_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
	op1 = opline.GetOp1Zval()
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	if (op1.IsString()) && (op2.IsString()) {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String
		if op1_str.GetLen() == 0 {
			{
				opline.GetResultZval().SetStringCopy(op2_str)
			}

			{
				types.ZendStringReleaseEx(op1_str, 0)
			}
		} else if op2_str.GetLen() == 0 {
			{
				opline.GetResultZval().SetStringCopy(op1_str)
			}

			{
				types.ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			opline.GetResultZval().SetString(str)
			{
				types.ZendStringReleaseEx(op1_str, 0)
			}
			{
				types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	if op1.IsString() {
		op1_str = op1.GetStr().Copy()
	} else {
		if op1.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		op1_str = ZvalGetStringFunc(op1)
	}

	if op2.IsString() {
		op2_str = op2.GetStr().Copy()
	} else {
		if op2.IsUndef() {
			ZVAL_UNDEFINED_OP2()
		}
		op2_str = ZvalGetStringFunc(op2)
	}
	for {
		{
			if op1_str.GetLen() == 0 {
				opline.GetResultZval().SetString(op2_str)
				types.ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		{
			if op2_str.GetLen() == 0 {
				opline.GetResultZval().SetString(op1_str)
				types.ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		opline.GetResultZval().SetString(str)
		{
			types.ZendStringReleaseEx(op1_str, 0)
		}
		{
			types.ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INIT_METHOD_CALL_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = opline.GetOp1Zval()
	{
		function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	}
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
					return 0
				}
			}
			faults.ThrowError(nil, "Method name must be a string")
			ZvalPtrDtorNogc(free_op2)
			return 0
			break
		}
	}
	{
		for {
			if object.GetType() != types.IS_OBJECT {
				if object.IsReference() {
					object = types.Z_REFVAL_P(object)
					if object.IsObject() {
						break
					}
				}
				if object.IsUndef() {
					object = ZVAL_UNDEFINED_OP1()
					if EG__().GetException() != nil {
						{
							ZvalPtrDtorNogc(free_op2)
						}
						return 0
					}
				}
				ZendInvalidMethodCall(object, function_name)
				ZvalPtrDtorNogc(free_op2)
				return 0
			}
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()

	{
		var orig_obj *types.ZendObject = obj
		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *types.Zval { return RT_CONSTANT(opline, opline.GetOp2()) + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			ZvalPtrDtorNogc(free_op2)
			return 0
		}
		if obj != orig_obj {

			/* Reset "object" to trigger reference counting */

			object = nil

			/* Reset "object" to trigger reference counting */

		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	{
		ZvalPtrDtorNogc(free_op2)
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		if EG__().GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*types.ZendObject)(called_scope)
		call_info = ZEND_CALL_NESTED_FUNCTION
	} else {
		{
			obj.AddRefcount()
		}

		/* CV may be changed indirectly (e.g. when it's a reference) */

		call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS | ZEND_CALL_RELEASE_THIS

		/* CV may be changed indirectly (e.g. when it's a reference) */

	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
