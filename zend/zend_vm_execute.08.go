package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_INIT_ARRAY_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
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
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CV_HANDLER(executeData)
	}

}
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var result int
	var hval ZendUlong
	var offset *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	offset = opline.GetOp2Zval()
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
			value = ht.IndexFind(hval)
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

				ZEND_VM_SMART_BRANCH(result, 0)
				types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else {
			result = value == nil || IZendIsTrue(value) == 0
		}
		goto isset_dim_obj_exit
	} else {
	}
	if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {
		result = ZendIssetDimSlow(container, offset, executeData)
	} else {
		result = ZendIsemptyDimSlow(container, offset, executeData)
	}
isset_dim_obj_exit:
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)
	{

		{
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ types.Z_OBJ_HT_P(container).GetHasProperty()(container, offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CV == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = RT_CONSTANT(opline, opline.GetOp1())
	subject = opline.GetOp2Zval()
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
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.GetResultZval().SetTypeInfo(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_YIELD_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
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
		if (executeData.GetFunc().op_array.fn_flags & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = RT_CONSTANT(opline, opline.GetOp1())
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
				{

					generator.GetValue().TryAddRefcount()

				}
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = RT_CONSTANT(opline, opline.GetOp1())

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)

				generator.GetValue().TryAddRefcount()

			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var key *types.Zval = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), executeData)

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
func ZEND_SL_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.GetOp1Zval()
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.GetLval() < SIZEOF_ZEND_LONG*8) != 0 {

		/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

		opline.GetResultZval().SetLong(zend_long(ZendUlong(op1.GetLval() << op2.GetLval())))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_shift_left_helper_SPEC(op1, op2, executeData)
}
func ZEND_SR_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.GetOp1Zval()
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.GetLval() < SIZEOF_ZEND_LONG*8) != 0 {
		opline.GetResultZval().SetLong(op1.GetLval() >> op2.GetLval())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_shift_right_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = opline.GetOp1Zval()
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() < op2.GetLval() {
			is_smaller_true:
				opline.GetResultZval().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_false:
				opline.GetResultZval().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_smaller_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = opline.GetOp1Zval()
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() < op2.GetLval() {
			is_smaller_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				opline.GetResultZval().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				opline.GetResultZval().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_smaller_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = opline.GetOp1Zval()
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() < op2.GetLval() {
			is_smaller_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
				opline.GetResultZval().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
				opline.GetResultZval().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_smaller_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = opline.GetOp1Zval()
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() <= op2.GetLval() {
			is_smaller_or_equal_true:
				opline.GetResultZval().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_or_equal_false:
				opline.GetResultZval().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_smaller_or_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_smaller_or_equal_double:
			if d1 <= d2 {
				goto is_smaller_or_equal_true
			} else {
				goto is_smaller_or_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_smaller_or_equal_double
		}
	}
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = opline.GetOp1Zval()
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() <= op2.GetLval() {
			is_smaller_or_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				opline.GetResultZval().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_or_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				opline.GetResultZval().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_smaller_or_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_smaller_or_equal_double:
			if d1 <= d2 {
				goto is_smaller_or_equal_true
			} else {
				goto is_smaller_or_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_smaller_or_equal_double
		}
	}
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = opline.GetOp1Zval()
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() <= op2.GetLval() {
			is_smaller_or_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
				opline.GetResultZval().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_or_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
				opline.GetResultZval().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_smaller_or_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_smaller_or_equal_double:
			if d1 <= d2 {
				goto is_smaller_or_equal_true
			} else {
				goto is_smaller_or_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_smaller_or_equal_double
		}
	}
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_BW_OR_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.GetOp1Zval()
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() && op2.IsLong() {
		opline.GetResultZval().SetLong(op1.GetLval() | op2.GetLval())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_bw_or_helper_SPEC(op1, op2, executeData)
}
func ZEND_BW_AND_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.GetOp1Zval()
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() && op2.IsLong() {
		opline.GetResultZval().SetLong(op1.GetLval() & op2.GetLval())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_bw_and_helper_SPEC(op1, op2, executeData)
}
func ZEND_BW_XOR_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.GetOp1Zval()
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() && op2.IsLong() {
		opline.GetResultZval().SetLong(op1.GetLval() ^ op2.GetLval())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_bw_xor_helper_SPEC(op1, op2, executeData)
}
func ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = opline.GetOp1Zval()
	zend_fetch_dimension_address_LIST_r(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_SWITCH_LONG_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op *types.Zval
	var jump_zv *types.Zval
	var jumptable *types.Array
	op = opline.GetOp1Zval()
	jumptable = RT_CONSTANT(opline, opline.GetOp2()).GetArr()
	if op.GetType() != types.IS_LONG {
		op = types.ZVAL_DEREF(op)
		if op.GetType() != types.IS_LONG {

			/* Wrong type, fall back to ZEND_CASE chain */

			return ZEND_VM_NEXT_OPCODE(executeData, opline)

			/* Wrong type, fall back to ZEND_CASE chain */

		}
	}
	jump_zv = jumptable.IndexFind(op.GetLval())
	if jump_zv != nil {
		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, jump_zv.GetLval())
		return 0
	} else {

		/* default */

		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, opline.GetExtendedValue())
		return 0
	}
}
func ZEND_SWITCH_STRING_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op *types.Zval
	var jump_zv *types.Zval
	var jumptable *types.Array
	op = opline.GetOp1Zval()
	jumptable = RT_CONSTANT(opline, opline.GetOp2()).GetArr()
	if op.GetType() != types.IS_STRING {

		/* Wrong type, fall back to ZEND_CASE chain */

		/* Wrong type, fall back to ZEND_CASE chain */

		{
			op = types.ZVAL_DEREF(op)
			if op.GetType() != types.IS_STRING {

				/* Wrong type, fall back to ZEND_CASE chain */

				return ZEND_VM_NEXT_OPCODE(executeData, opline)

				/* Wrong type, fall back to ZEND_CASE chain */

			}
		}
	}
	jump_zv = jumptable.KeyFind(op.GetStr().GetStr())
	if jump_zv != nil {
		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, jump_zv.GetLval())
		return 0
	} else {

		/* default */

		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, opline.GetExtendedValue())
		return 0
	}
}
