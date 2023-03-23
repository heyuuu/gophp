// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_DIM_W_SPEC_VAR_CONST_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_TMPVAR_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_UNSET_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	zend_fetch_dimension_address_UNSET(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	{
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_W_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_VAR, property, IS_CONST, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_FETCH_OBJ_FLAGS) }, nil), BP_VAR_W, opline.GetExtendedValue()&ZEND_FETCH_OBJ_FLAGS, 1, opline, executeData)
	{
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_RW_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_VAR, property, IS_CONST, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_RW, 0, 1, opline, executeData)
	{
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_OBJ_W_SPEC_VAR_CONST_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var property *types.Zval
	var result *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_VAR, property, IS_CONST, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_UNSET, 0, 1, opline, executeData)
	{
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_LIST_W_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	dim = RT_CONSTANT(opline, opline.GetOp2())
	if EX_VAR(opline.GetOp1().GetVar()).GetType() != types.IS_INDIRECT && !(container.IsReference()) {
		faults.Error(faults.E_NOTICE, "Attempting to set reference to non referenceable value")
		zend_fetch_dimension_address_LIST_r(container, dim, IS_CONST, opline, executeData)
	} else {
		zend_fetch_dimension_address_W(container, dim, IS_CONST, opline, executeData)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = RT_CONSTANT(opline, opline.GetOp2())
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
	if types.Z_OBJCE_P(object) == CACHED_PTR(opline.GetExtendedValue()) {
		var cache_slot *any = CACHE_ADDR(opline.GetExtendedValue())
		var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
		var zobj *types.ZendObject = object.GetObj()
		var property_val *types.Zval
		if IS_VALID_PROPERTY_OFFSET(prop_offset) {
			property_val = OBJ_PROP(zobj, prop_offset)
			if property_val.GetType() != types.IS_UNDEF {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, executeData)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, IS_CONST, executeData.IsCallUseStrictTypes())
					if RETURN_VALUE_USED(opline) {
						types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if zobj.GetProperties().GetRefcount() > 1 {
					if (zobj.GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
						zobj.GetProperties().DelRefcount()
					}
					zobj.SetProperties(types.ZendArrayDup(zobj.GetProperties()))
				}
				property_val = zobj.GetProperties().KeyFind(property.GetStr().GetStr())
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				{

					value.TryAddRefcount()

				}

				zobj.GetProperties().KeyAddNew(property.GetStr().GetStr(), value)
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
				}
				goto exit_assign_obj
			}
		}
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
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
func ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = RT_CONSTANT(opline, opline.GetOp2())
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
	if types.Z_OBJCE_P(object) == CACHED_PTR(opline.GetExtendedValue()) {
		var cache_slot *any = CACHE_ADDR(opline.GetExtendedValue())
		var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
		var zobj *types.ZendObject = object.GetObj()
		var property_val *types.Zval
		if IS_VALID_PROPERTY_OFFSET(prop_offset) {
			property_val = OBJ_PROP(zobj, prop_offset)
			if property_val.GetType() != types.IS_UNDEF {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, executeData)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
					if RETURN_VALUE_USED(opline) {
						types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if zobj.GetProperties().GetRefcount() > 1 {
					if (zobj.GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
						zobj.GetProperties().DelRefcount()
					}
					zobj.SetProperties(types.ZendArrayDup(zobj.GetProperties()))
				}
				property_val = zobj.GetProperties().KeyFind(property.GetStr().GetStr())
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				zobj.GetProperties().KeyAddNew(property.GetStr().GetStr(), value)
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
				}
				goto exit_assign_obj
			}
		}
	}
	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
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
func ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = RT_CONSTANT(opline, opline.GetOp2())
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
	if types.Z_OBJCE_P(object) == CACHED_PTR(opline.GetExtendedValue()) {
		var cache_slot *any = CACHE_ADDR(opline.GetExtendedValue())
		var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
		var zobj *types.ZendObject = object.GetObj()
		var property_val *types.Zval
		if IS_VALID_PROPERTY_OFFSET(prop_offset) {
			property_val = OBJ_PROP(zobj, prop_offset)
			if property_val.GetType() != types.IS_UNDEF {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, executeData)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, IS_VAR, executeData.IsCallUseStrictTypes())
					if RETURN_VALUE_USED(opline) {
						types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if zobj.GetProperties().GetRefcount() > 1 {
					if (zobj.GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
						zobj.GetProperties().DelRefcount()
					}
					zobj.SetProperties(types.ZendArrayDup(zobj.GetProperties()))
				}
				property_val = zobj.GetProperties().KeyFind(property.GetStr().GetStr())
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				zobj.GetProperties().KeyAddNew(property.GetStr().GetStr(), value)
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
				}
				goto exit_assign_obj
			}
		}
	}
	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
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
func ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = RT_CONSTANT(opline, opline.GetOp2())
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
	if types.Z_OBJCE_P(object) == CACHED_PTR(opline.GetExtendedValue()) {
		var cache_slot *any = CACHE_ADDR(opline.GetExtendedValue())
		var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
		var zobj *types.ZendObject = object.GetObj()
		var property_val *types.Zval
		if IS_VALID_PROPERTY_OFFSET(prop_offset) {
			property_val = OBJ_PROP(zobj, prop_offset)
			if property_val.GetType() != types.IS_UNDEF {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, executeData)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, IS_CV, executeData.IsCallUseStrictTypes())
					if RETURN_VALUE_USED(opline) {
						types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if zobj.GetProperties().GetRefcount() > 1 {
					if (zobj.GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
						zobj.GetProperties().DelRefcount()
					}
					zobj.SetProperties(types.ZendArrayDup(zobj.GetProperties()))
				}
				property_val = zobj.GetProperties().KeyFind(property.GetStr().GetStr())
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				zobj.GetProperties().KeyAddNew(property.GetStr().GetStr(), value)
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
				}
				goto exit_assign_obj
			}
		}
	}
	{
		value = types.ZVAL_DEREF(value)
	}
	value = types.Z_OBJ_HT_P(object).GetWriteProperty()(object, property, value, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
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
				object_ptr.SetArray(types.NewZendArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = RT_CONSTANT(opline, opline.GetOp2())
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
				object_ptr.SetArray(types.NewZendArray(8))
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
				object_ptr.SetArray(types.NewZendArray(8))
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
				object_ptr.SetArray(types.NewZendArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = RT_CONSTANT(opline, opline.GetOp2())
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
func ZEND_ASSIGN_SPEC_VAR_CONST_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = RT_CONSTANT(opline, opline.GetOp2())
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())

		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_SPEC_VAR_CONST_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = RT_CONSTANT(opline, opline.GetOp2())
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		EX_VAR(opline.GetResult().GetVar()).SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_CONST, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CONST_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = RT_CONSTANT(opline, opline.GetOp2())
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)

	{
		{
			ZendAssignToPropertyReferenceVarConst(container, property, value_ptr, opline, executeData)
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
func ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CONST_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = RT_CONSTANT(opline, opline.GetOp2())
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), executeData)

	{
		{
			ZendAssignToPropertyReferenceVarConst(container, property, value_ptr, opline, executeData)
		}

	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc *ZendFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = EX_VAR(opline.GetOp1().GetVar()).GetCe()
	}

	if CACHED_PTR(opline.GetResult().GetNum()) == ce {
		fbc = CACHED_PTR(opline.GetResult().GetNum() + b.SizeOf("void *"))
	} else {
		function_name = RT_CONSTANT(opline, opline.GetOp2())
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), b.CondF1(IS_CONST == IS_CONST, func() *types.Zval { return RT_CONSTANT(opline, opline.GetOp2()) + 1 }, nil))
		}
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			HANDLE_EXCEPTION()
		}
		if fbc.GetType() <= ZEND_USER_FUNCTION && !fbc.HasFnFlags(ZEND_ACC_CALL_VIA_TRAMPOLINE|ZEND_ACC_NEVER_CACHE) {
			CACHE_POLYMORPHIC_PTR(opline.GetResult().GetNum(), ce, fbc)
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
func ZEND_FETCH_CLASS_CONSTANT_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var c *ZendClassConstant
	var value *types.Zval
	var zv *types.Zval
	var opline *ZendOp = executeData.GetOpline()
	for {

		{

			{
				ce = EX_VAR(opline.GetOp1().GetVar()).GetCe()
			}
			if CACHED_PTR(opline.GetExtendedValue()) == ce {
				value = CACHED_PTR(opline.GetExtendedValue() + b.SizeOf("void *"))
				break
			}
		}
		zv = ce.GetConstantsTable().KeyFind(RT_CONSTANT(opline, opline.GetOp2()).GetStr().GetStr())
		if zv != nil {
			c = zv.GetPtr()
			scope = executeData.GetFunc().op_array.scope
			if ZendVerifyConstAccess(c, scope) == 0 {
				faults.ThrowError(nil, "Cannot access %s const %s::%s", ZendVisibilityString(c.GetValue().GetAccessFlags()), ce.GetName().GetVal(), RT_CONSTANT(opline, opline.GetOp2()).GetStr().GetVal())
				EX_VAR(opline.GetResult().GetVar()).SetUndef()
				HANDLE_EXCEPTION()
			}
			value = c.GetValue()
			if value.IsConstant() {
				ZvalUpdateConstantEx(value, c.GetCe())
				if EG__().GetException() != nil {
					EX_VAR(opline.GetResult().GetVar()).SetUndef()
					HANDLE_EXCEPTION()
				}
			}
			CACHE_POLYMORPHIC_PTR(opline.GetExtendedValue(), ce, value)
		} else {
			faults.ThrowError(nil, "Undefined class constant '%s'", RT_CONSTANT(opline, opline.GetOp2()).GetStr().GetVal())
			EX_VAR(opline.GetResult().GetVar()).SetUndef()
			HANDLE_EXCEPTION()
		}
		break
	}
	types.ZVAL_COPY_OR_DUP(EX_VAR(opline.GetResult().GetVar()), value)
	ZEND_VM_NEXT_OPCODE()
}
