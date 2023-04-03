package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

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
				opline.GetResultZval().SetNull()
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
				opline.GetResultZval().SetNull()
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
				opline.GetResultZval().SetNull()
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
func ZEND_ASSIGN_OBJ_REF_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.getZvalPtrVar2(&free_op2)
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)

	{

		{
			ZendAssignToPropertyReferenceVarVar(container, property, value_ptr, opline, executeData)
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZvalPtrDtorNogc(free_op2)
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var property *types.Zval
	var container *types.Zval
	var value_ptr *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.getZvalPtrVar2(&free_op2)
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), executeData)

	{

		{
			ZendAssignToPropertyReferenceVarVar(container, property, value_ptr, opline, executeData)
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZvalPtrDtorNogc(free_op2)
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = opline.GetOp1Zval().GetCe()
	}

	{
		var free_op2 ZendFreeOp
		function_name = opline.getZvalPtrVar2(&free_op2)
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
							return 0
						}
					}
					faults.ThrowError(nil, "Function name must be a string")
					ZvalPtrDtorNogc(free_op2)
					return 0
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *types.Zval { return RT_CONSTANT(opline, opline.GetOp2()) + 1 }, nil))
		}
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			ZvalPtrDtorNogc(free_op2)
			return 0
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		{
			ZvalPtrDtorNogc(free_op2)
		}
	}

	if !fbc.IsStatic() {
		if executeData.GetThis().IsObject() && InstanceofFunction(types.Z_OBJCE(executeData.GetThis()), ce) != 0 {
			ce = (*types.ClassEntry)(executeData.GetThis().GetObj())
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().GetException() != nil {
				return 0
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
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
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
		expr_ptr = opline.getZvalPtrVar1(&free_op1)
	}
	{
		var free_op2 ZendFreeOp
		var offset *types.Zval = opline.getZvalPtrVar2(&free_op2)
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
			opline.GetResultZval().GetArr().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.GetLval()
		num_index:
			opline.GetResultZval().GetArr().IndexUpdate(hval, expr_ptr)
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
		ZvalPtrDtorNogc(free_op2)
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INIT_ARRAY_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.GetResultZval()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			types.ZendHashRealInitMixed(array.GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_TMPVAR_HANDLER(executeData)
	}

}
func ZEND_UNSET_DIM_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	var hval ZendUlong
	var key *types.String
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	offset = opline.getZvalPtrVar2(&free_op2)
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
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_OBJ_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	offset = opline.getZvalPtrVar2(&free_op2)
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
		types.Z_OBJ_HT_P(container).GetUnsetProperty()(container, offset, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue()) }, nil))
		break
	}
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_IDENTICAL_SPEC_VAR_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = FastIsIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_NOT_IDENTICAL_SPEC_VAR_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = FastIsNotIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_VAR_TMP_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)

	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())

		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_VAR_TMP_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)
		opline.GetResultZval().SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(opline.GetResultZval(), value)
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_YIELD_SPEC_VAR_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		var free_op1 ZendFreeOp
		if (executeData.GetFunc().op_array.fn_flags & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.getZvalPtrVar1(&free_op1)
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.getZvalPtrVar1(&free_op1)

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().GetLval() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().GetLval())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.GetResultZval())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_IS_IDENTICAL_SPEC_VAR_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrVarDeref(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = FastIsIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_NOT_IDENTICAL_SPEC_VAR_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = _getZvalPtrVarDeref(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = FastIsNotIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_VAR_VAR_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = opline.getZvalPtrVar2(&free_op2)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)

	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())

		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_SPEC_VAR_VAR_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = opline.getZvalPtrVar2(&free_op2)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)
		opline.GetResultZval().SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(opline.GetResultZval(), value)
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ASSIGN_REF_SPEC_VAR_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var variable_ptr *types.Zval
	var value_ptr *types.Zval
	value_ptr = _getZvalPtrPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if variable_ptr.IsError() {
		variable_ptr = EG__().GetUninitializedZval()
	} else if opline.GetOp1Zval().GetType() != types.IS_INDIRECT {
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
		types.ZVAL_COPY(opline.GetResultZval(), variable_ptr)
	}
	if free_op2 != nil {
		ZvalPtrDtorNogc(free_op2)
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_YIELD_SPEC_VAR_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		var free_op1 ZendFreeOp
		if (executeData.GetFunc().op_array.fn_flags & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.getZvalPtrVar1(&free_op1)
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.getZvalPtrVar1(&free_op1)

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = opline.getZvalPtrVar2(&free_op2)

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().GetLval() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().GetLval())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.GetResultZval())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
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
			types.ZVAL_COPY(opline.GetResultZval(), var_ptr)
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
				opline.GetResultZval().SetNull()
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	OPLINE = executeData.GetOpline() + 2
	return 0
}
