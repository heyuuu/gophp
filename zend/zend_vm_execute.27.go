// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/types"
)

func ZEND_IS_NOT_EQUAL_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() != op2.GetLval() {
			is_not_equal_true:
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_not_equal_false:
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_not_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_not_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{
				ZvalPtrDtorStr(op1)
			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_NOT_EQUAL_SPEC_CV_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() != op2.GetLval() {
			is_not_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_not_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_not_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_not_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{
				ZvalPtrDtorStr(op1)
			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_NOT_EQUAL_SPEC_CV_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() != op2.GetLval() {
			is_not_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_not_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_not_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_not_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{
				ZvalPtrDtorStr(op1)
			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_SPACESHIP_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	CompareFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_BOOL_XOR_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	BooleanXorFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_OBJ_OP_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = EX_VAR(opline.GetOp1().GetVar())
	property = RT_CONSTANT(opline, opline.GetOp2())
	for {
		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data)
		if object.GetType() != types.IS_OBJECT {
			if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
				object = types.Z_REFVAL_P(object)
				goto assign_op_object
			}
			if object.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			object = MakeRealObject(object, property, opline, executeData)
			if object == nil {
				break
			}
		}
	assign_op_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = CACHE_ADDR((opline + 1).GetExtendedValue())
		}

		if b.Assign(&zptr, types.Z_OBJ_HT_P(object).GetGetPropertyPtrPtr()(object, property, BP_VAR_RW, cache_slot)) != nil {
			if zptr.IsError() {
				if RETURN_VALUE_USED(opline) {
					EX_VAR(opline.GetResult().GetVar()).SetNull()
				}
			} else {
				var orig_zptr *types.Zval = zptr
				var ref *types.ZendReference
				for {
					if zptr.IsReference() {
						ref = zptr.GetRef()
						zptr = types.Z_REFVAL_P(zptr)
						if ZEND_REF_HAS_TYPE_SOURCES(ref) {
							ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
							break
						}
					}
					{
						prop_info = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
					}

					if prop_info != nil {

						/* special case for typed properties */

						ZendBinaryAssignOpTypedProp(prop_info, zptr, value, opline, executeData)

						/* special case for typed properties */

					} else {
						ZendBinaryOp(zptr, zptr, value, opline)
					}
					break
				}
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), zptr)
				}
			}
		} else {
			ZendAssignOpOverloadedProperty(object, property, cache_slot, value, opline, executeData)
		}
		break
	}
	FREE_OP(free_op_data)

	/* assign_obj has two opcodes! */

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
	container = EX_VAR(opline.GetOp1().GetVar())
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
			types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), var_ptr)
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
			container.SetArray(types.NewZendArray(8))
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, executeData)
		assign_dim_op_ret_null:
			FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
			if RETURN_VALUE_USED(opline) {
				EX_VAR(opline.GetResult().GetVar()).SetNull()
			}
		}
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OP_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var var_ptr *types.Zval
	var value *types.Zval
	value = RT_CONSTANT(opline, opline.GetOp2())
	var_ptr = _get_zval_ptr_cv_BP_VAR_RW(opline.GetOp1().GetVar(), executeData)
	if var_ptr.IsError() {
		if RETURN_VALUE_USED(opline) {
			EX_VAR(opline.GetResult().GetVar()).SetNull()
		}
	} else {
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
			types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), var_ptr)
		}
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_PRE_INC_OBJ_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = EX_VAR(opline.GetOp1().GetVar())
	property = RT_CONSTANT(opline, opline.GetOp2())
	for {
		if object.GetType() != types.IS_OBJECT {
			if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
				object = types.Z_REFVAL_P(object)
				goto pre_incdec_object
			}
			if object.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			object = MakeRealObject(object, property, opline, executeData)
			if object == nil {
				break
			}
		}
	pre_incdec_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = CACHE_ADDR(opline.GetExtendedValue())
		}

		if b.Assign(&zptr, types.Z_OBJ_HT_P(object).GetGetPropertyPtrPtr()(object, property, BP_VAR_RW, cache_slot)) != nil {
			if zptr.IsError() {
				if RETURN_VALUE_USED(opline) {
					EX_VAR(opline.GetResult().GetVar()).SetNull()
				}
			} else {
				{
					prop_info = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
				}

				ZendPreIncdecPropertyZval(zptr, prop_info, opline, executeData)
			}
		} else {
			ZendPreIncdecOverloadedProperty(object, property, cache_slot, opline, executeData)
		}
		break
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POST_INC_OBJ_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = EX_VAR(opline.GetOp1().GetVar())
	property = RT_CONSTANT(opline, opline.GetOp2())
	for {
		if object.GetType() != types.IS_OBJECT {
			if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
				object = types.Z_REFVAL_P(object)
				goto post_incdec_object
			}
			if object.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			object = MakeRealObject(object, property, opline, executeData)
			if object == nil {
				break
			}
		}
	post_incdec_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = CACHE_ADDR(opline.GetExtendedValue())
		}

		if b.Assign(&zptr, types.Z_OBJ_HT_P(object).GetGetPropertyPtrPtr()(object, property, BP_VAR_RW, cache_slot)) != nil {
			if zptr.IsError() {
				EX_VAR(opline.GetResult().GetVar()).SetNull()
			} else {
				{
					prop_info = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_slot + 2))
				}

				ZendPostIncdecPropertyZval(zptr, prop_info, opline, executeData)
			}
		} else {
			ZendPostIncdecOverloadedProperty(object, property, cache_slot, opline, executeData)
		}
		break
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_R_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	dim = RT_CONSTANT(opline, opline.GetOp2())
	{
		if container.IsArray() {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetArr(), dim, IS_CONST, BP_VAR_R, executeData)
			types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), value)
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

	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_W_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	zend_fetch_dimension_address_W(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	{
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_RW_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	zend_fetch_dimension_address_RW(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	{
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_IS_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	zend_fetch_dimension_address_read_IS(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_DIM_W_SPEC_CV_CONST_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_CV_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_UNSET_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	zend_fetch_dimension_address_UNSET(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	{
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_R_SPEC_CV_CONST_INLINE_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = EX_VAR(opline.GetOp1().GetVar())
	offset = RT_CONSTANT(opline, opline.GetOp2())
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
			ZendWrongPropertyRead(offset)
			EX_VAR(opline.GetResult().GetVar()).SetNull()
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval
	{
		cache_slot = CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_FETCH_REF)
		if zobj.GetCe() == CACHED_PTR_EX(cache_slot) {
			var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
			if IS_VALID_PROPERTY_OFFSET(prop_offset) {
				retval = OBJ_PROP(zobj, prop_offset)
				if !retval.IsUndef() {
					{
						goto fetch_obj_r_copy
					}

				}
			} else if zobj.GetProperties() != nil {
				if !(IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(prop_offset)) {
					var idx uintPtr = ZEND_DECODE_DYN_PROP_OFFSET(prop_offset)
					if idx < zobj.GetProperties().GetNNumUsed()*b.SizeOf("Bucket") {
						var p *types.Bucket = (*types.Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
						if p.GetVal().IsNotUndef() && (p.GetKey() == offset.GetStr() || (p.IsStrKey() && p.StrKey() == offset.GetStrVal())) {
							retval = p.GetVal()
							{
								goto fetch_obj_r_copy
							}

						}
					}
					CACHE_PTR_EX(cache_slot+1, any(ZEND_DYNAMIC_PROPERTY_OFFSET))
				}
				retval = zobj.GetProperties().KeyFind(offset.GetStr().GetStr())
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					CACHE_PTR_EX(cache_slot+1, any(ZEND_ENCODE_DYN_PROP_OFFSET(idx)))
					{
						goto fetch_obj_r_copy
					}

				}
			}
		}
	}

	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, EX_VAR(opline.GetResult().GetVar()))
	if retval != EX_VAR(opline.GetResult().GetVar()) {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_R_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	return ZEND_FETCH_OBJ_R_SPEC_CV_CONST_INLINE_HANDLER(executeData)
}
func ZEND_FETCH_OBJ_W_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	property = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_CV, property, IS_CONST, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_FETCH_OBJ_FLAGS) }, nil), BP_VAR_W, opline.GetExtendedValue()&ZEND_FETCH_OBJ_FLAGS, 1, opline, executeData)
	{
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_RW_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var result *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	property = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_CV, property, IS_CONST, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_RW, 0, 1, opline, executeData)
	{
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_IS_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = _get_zval_ptr_cv_BP_VAR_IS(opline.GetOp1().GetVar(), executeData)
	offset = RT_CONSTANT(opline, opline.GetOp2())
	if container.GetType() != types.IS_OBJECT {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			EX_VAR(opline.GetResult().GetVar()).SetNull()
			goto fetch_obj_is_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval
	{
		cache_slot = CACHE_ADDR(opline.GetExtendedValue())
		if zobj.GetCe() == CACHED_PTR_EX(cache_slot) {
			var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
			if IS_VALID_PROPERTY_OFFSET(prop_offset) {
				retval = OBJ_PROP(zobj, prop_offset)
				if retval.IsNotUndef() {
					{
						goto fetch_obj_is_copy
					}

				}
			} else if zobj.GetProperties() != nil {
				if !(IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(prop_offset)) {
					var idx uintPtr = ZEND_DECODE_DYN_PROP_OFFSET(prop_offset)
					if idx < zobj.GetProperties().GetNNumUsed()*b.SizeOf("Bucket") {
						var p *types.Bucket = (*types.Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
						if p.GetVal().IsNotUndef() && (p.GetKey() == offset.GetStr() || (p.IsStrKey() && p.StrKey() == offset.GetStrVal())) {
							retval = p.GetVal()
							{
								goto fetch_obj_is_copy
							}

						}
					}
					CACHE_PTR_EX(cache_slot+1, any(ZEND_DYNAMIC_PROPERTY_OFFSET))
				}
				retval = zobj.GetProperties().KeyFind(offset.GetStr().GetStr())
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					CACHE_PTR_EX(cache_slot+1, any(ZEND_ENCODE_DYN_PROP_OFFSET(idx)))
					{
						goto fetch_obj_is_copy
					}

				}
			}
		}
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_IS, cache_slot, EX_VAR(opline.GetResult().GetVar()))
	if retval != EX_VAR(opline.GetResult().GetVar()) {
	fetch_obj_is_copy:
		types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_is_finish:
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_OBJ_W_SPEC_CV_CONST_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CV_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var property *types.Zval
	var result *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	property = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	ZendFetchPropertyAddress(result, container, IS_CV, property, IS_CONST, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil), BP_VAR_UNSET, 0, 1, opline, executeData)
	{
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = EX_VAR(opline.GetOp1().GetVar())
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
			if property_val.IsNotUndef() {
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

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = EX_VAR(opline.GetOp1().GetVar())
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
			if property_val.IsNotUndef() {
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

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = EX_VAR(opline.GetOp1().GetVar())
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
			if property_val.IsNotUndef() {
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

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
