// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr_ptr *types.Zval
	var new_expr types.Zval
	if (opline.GetExtendedValue() & ZEND_ARRAY_ELEMENT_REF) != 0 {
		expr_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), executeData)
		if expr_ptr.IsReference() {
			expr_ptr.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(expr_ptr, 2)
		}
	} else {
		expr_ptr = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	}
	{
		var free_op2 ZendFreeOp
		var offset *types.Zval = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
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
			EX_VAR(opline.GetResult().GetVar()).GetArr().IndexUpdateH(hval, expr_ptr)
		} else if offset.IsReference() {
			offset = types.Z_REFVAL_P(offset)
			goto add_again
		} else if offset.IsNull() {
			str = types.ZSTR_EMPTY_ALLOC()
			goto str_index
		} else if offset.IsDouble() {
			hval = ZendDvalToLval(offset.GetDval())
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

	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INIT_ARRAY_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = EX_VAR(opline.GetResult().GetVar())
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewZendArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			types.ZendHashRealInitMixed(array.GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_TMPVAR_HANDLER(executeData)
	}

}
func ZEND_UNSET_DIM_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	var hval ZendUlong
	var key *types.String
	container = EX_VAR(opline.GetOp1().GetVar())
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
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
				hval = ZendDvalToLval(offset.GetDval())
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
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_UNSET_OBJ_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
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
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var result int
	var hval ZendUlong
	var offset *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	if container.IsArray() {
		var ht *types.Array
		var value *types.Zval
		var str *types.String
	isset_dim_obj_array:
		ht = container.GetArr()
	isset_again:
		if offset.IsString() {
			str = offset.GetStr()
			{
				if types.HandleNumericStr(str.GetStr(), &hval) {
					goto num_index_prop
				}
			}
			value = types.ZendHashFindInd(ht, str.GetStr())
		} else if offset.IsLong() {
			hval = offset.GetLval()
		num_index_prop:
			value = ht.IndexFindH(hval)
		} else if offset.IsReference() {
			offset = types.Z_REFVAL_P(offset)
			goto isset_again
		} else {
			value = ZendFindArrayDimSlow(ht, offset, executeData)
			if EG__().GetException() != nil {
				result = 0
				goto isset_dim_obj_exit
			}
		}
		if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {

			/* > IS_NULL means not IS_UNDEF and not IS_NULL */

			result = value != nil && value.GetType() > types.IS_NULL && (!(value.IsReference()) || types.Z_REFVAL_P(value).GetType() != types.IS_NULL)
			{

				/* avoid exception check */

				ZvalPtrDtorNogc(free_op2)
				ZEND_VM_SMART_BRANCH(result, 0)
				types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
				ZEND_VM_NEXT_OPCODE()
			}
		} else {
			result = value == nil || IZendIsTrue(value) == 0
		}
		goto isset_dim_obj_exit
	} else if container.IsReference() {
		container = types.Z_REFVAL_P(container)
		if container.IsArray() {
			goto isset_dim_obj_array
		}
	}
	if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {
		result = ZendIssetDimSlow(container, offset, executeData)
	} else {
		result = ZendIsemptyDimSlow(container, offset, executeData)
	}
isset_dim_obj_exit:
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = _get_zval_ptr_cv_BP_VAR_IS(opline.GetOp1().GetVar(), executeData)
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	if container.GetType() != types.IS_OBJECT {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.GetType() != types.IS_OBJECT {
				result = opline.GetExtendedValue() & ZEND_ISEMPTY
				goto isset_object_finish
			}
		} else {
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ types.Z_OBJ_HT_P(container).GetHasProperty()(container, offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = EX_VAR(opline.GetOp1().GetVar())
	subject = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	if subject.IsArray() {
	array_key_exists_array:
		ht = subject.GetArr()
		result = ZendArrayKeyExistsFast(ht, key, opline, executeData)
	} else {
		if subject.IsReference() {
			subject = types.Z_REFVAL_P(subject)
			if subject.IsArray() {
				goto array_key_exists_array
			}
		}
		result = ZendArrayKeyExistsSlow(subject, key, opline, executeData)
	}
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	EX_VAR(opline.GetResult().GetVar()).SetTypeInfo(result)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_IS_IDENTICAL_SPEC_CV_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = FastIsIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_IS_NOT_IDENTICAL_SPEC_CV_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = FastIsNotIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_SPEC_CV_TMP_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = EX_VAR(opline.GetOp1().GetVar())
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)

	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())

	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_SPEC_CV_TMP_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = EX_VAR(opline.GetOp1().GetVar())
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)
		EX_VAR(opline.GetResult().GetVar()).SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_YIELD_SPEC_CV_TMP_HANDLER(executeData *ZendExecuteData) int {
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
		if (executeData.GetFunc().op_array.fn_flags & ZEND_ACC_RETURN_REFERENCE) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)

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

		generator.SetSendTarget(EX_VAR(opline.GetResult().GetVar()))
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE()

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_IS_IDENTICAL_SPEC_CV_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = _getZvalPtrVarDeref(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = FastIsIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_IS_NOT_IDENTICAL_SPEC_CV_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = _getZvalPtrVarDeref(opline.GetOp2().GetVar(), &free_op2, executeData)
	result = FastIsNotIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_SPEC_CV_VAR_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = EX_VAR(opline.GetOp1().GetVar())
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)

	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())

	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_SPEC_CV_VAR_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	value = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = EX_VAR(opline.GetOp1().GetVar())
	if variable_ptr.IsError() {
		ZvalPtrDtorNogc(free_op2)
		EX_VAR(opline.GetResult().GetVar()).SetNull()
	} else {
		value = ZendAssignToVariable(variable_ptr, value, IS_VAR, executeData.IsCallUseStrictTypes())
		types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_REF_SPEC_CV_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var variable_ptr *types.Zval
	var value_ptr *types.Zval
	value_ptr = _getZvalPtrPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	variable_ptr = EX_VAR(opline.GetOp1().GetVar())
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
	if free_op2 != nil {
		ZvalPtrDtorNogc(free_op2)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INSTANCEOF_SPEC_CV_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr *types.Zval
	var result types.ZendBool
	expr = EX_VAR(opline.GetOp1().GetVar())
try_instanceof:
	if expr.IsObject() {
		var ce *types.ClassEntry

		{
			ce = EX_VAR(opline.GetOp2().GetVar()).GetCe()
		}
		result = ce != nil && InstanceofFunction(types.Z_OBJCE_P(expr), ce) != 0
	} else if expr.IsReference() {
		expr = types.Z_REFVAL_P(expr)
		goto try_instanceof
	} else {
		if expr.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		result = 0
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_YIELD_SPEC_CV_VAR_HANDLER(executeData *ZendExecuteData) int {
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
		if (executeData.GetFunc().op_array.fn_flags & ZEND_ACC_RETURN_REFERENCE) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)

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
		var key *types.Zval = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

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

		generator.SetSendTarget(EX_VAR(opline.GetResult().GetVar()))
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE()

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_ASSIGN_DIM_OP_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
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
		dim = nil
		if container.IsObject() {
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
func zend_fetch_var_address_helper_SPEC_CV_UNUSED(type_ int, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var retval *types.Zval
	var name *types.String
	var tmp_name *types.String
	var target_symbol_table *types.Array
	varname = EX_VAR(opline.GetOp1().GetVar())

	if varname.IsString() {
		name = varname.GetStr()
		tmp_name = nil
	} else {
		if varname.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			EX_VAR(opline.GetResult().GetVar()).SetUndef()
			HANDLE_EXCEPTION()
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), executeData)
	retval = target_symbol_table.KeyFind(name.GetStr())
	if retval == nil {
		if types.ZendStringEquals(name, types.ZSTR_THIS) != 0 {
		fetch_this:
			ZendFetchThisVar(type_, opline, executeData)
			{
				ZendTmpStringRelease(tmp_name)
			}
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}
		if type_ == BP_VAR_W {
			retval = target_symbol_table.KeyAddNew(name.GetStr(), EG__().GetUninitializedZval())
		} else if type_ == BP_VAR_IS {
			retval = EG__().GetUninitializedZval()
		} else {
			faults.Error(faults.E_NOTICE, "Undefined variable: %s", name.GetVal())
			if type_ == BP_VAR_RW {
				retval = target_symbol_table.KeyUpdate(name.GetStr(), EG__().GetUninitializedZval())
			} else {
				retval = EG__().GetUninitializedZval()
			}
		}
	} else if retval.IsIndirect() {
		retval = retval.GetZv()
		if retval.IsUndef() {
			if types.ZendStringEquals(name, types.ZSTR_THIS) != 0 {
				goto fetch_this
			}
			if type_ == BP_VAR_W {
				retval.SetNull()
			} else if type_ == BP_VAR_IS {
				retval = EG__().GetUninitializedZval()
			} else {
				faults.Error(faults.E_NOTICE, "Undefined variable: %s", name.GetVal())
				if type_ == BP_VAR_RW {
					retval.SetNull()
				} else {
					retval = EG__().GetUninitializedZval()
				}
			}
		}
	}
	if (opline.GetExtendedValue() & ZEND_FETCH_GLOBAL_LOCK) == 0 {
	}
	{
		ZendTmpStringRelease(tmp_name)
	}
	b.Assert(retval != nil)
	if type_ == BP_VAR_R || type_ == BP_VAR_IS {
		types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), retval)
	} else {
		EX_VAR(opline.GetResult().GetVar()).SetIndirect(retval)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_R_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(BP_VAR_R, executeData)
}
func ZEND_FETCH_W_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(BP_VAR_W, executeData)
}
func ZEND_FETCH_RW_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(BP_VAR_RW, executeData)
}
func ZEND_FETCH_FUNC_ARG_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var fetch_type int = b.Cond((ZEND_CALL_INFO(executeData.GetCall())&ZEND_CALL_SEND_ARG_BY_REF) != 0, BP_VAR_W, BP_VAR_R)
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(fetch_type, executeData)
}
func ZEND_FETCH_UNSET_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(BP_VAR_UNSET, executeData)
}
func ZEND_FETCH_IS_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(BP_VAR_IS, executeData)
}
func ZEND_FETCH_DIM_W_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	zend_fetch_dimension_address_W(container, nil, IS_UNUSED, opline, executeData)
	{
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_RW_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	zend_fetch_dimension_address_RW(container, nil, IS_UNUSED, opline, executeData)
	{
		var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
		FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op1, result)
	}
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_DIM_W_SPEC_CV_UNUSED_HANDLER(executeData)
	} else {
		{
			return zend_use_undef_in_read_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	}
}
func ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var object_ptr *types.Zval
	var orig_object_ptr *types.Zval
	var free_op_data ZendFreeOp
	var value *types.Zval
	var variable_ptr *types.Zval
	var dim *types.Zval
	object_ptr = EX_VAR(opline.GetOp1().GetVar())
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
				object_ptr.SetArray(types.NewZendArray(8))
				goto try_assign_dim_array
			}
		} else {
			if !(object_ptr.IsError()) {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			if RETURN_VALUE_USED(opline) {
				EX_VAR(opline.GetResult().GetVar()).SetNull()
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
	object_ptr = EX_VAR(opline.GetOp1().GetVar())
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
				object_ptr.SetArray(types.NewZendArray(8))
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
				EX_VAR(opline.GetResult().GetVar()).SetNull()
			}
		}
	}
	/* assign_dim has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
