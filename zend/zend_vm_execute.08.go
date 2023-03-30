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
	array = EX_VAR(opline.GetResult().GetVar())
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
	offset = EX_VAR(opline.GetOp2().GetVar())
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
				types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
				ZEND_VM_NEXT_OPCODE()
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
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
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
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = RT_CONSTANT(opline, opline.GetOp1())
	subject = EX_VAR(opline.GetOp2().GetVar())
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
	EX_VAR(opline.GetResult().GetVar()).SetTypeInfo(result)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
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
		if (executeData.GetFunc().op_array.fn_flags & ZEND_ACC_RETURN_REFERENCE) != 0 {

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
func ZEND_QM_ASSIGN_LONG_SPEC_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = EX_VAR(opline.GetOp1().GetVar())
	EX_VAR(opline.GetResult().GetVar()).SetLong(value.GetLval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_QM_ASSIGN_DOUBLE_SPEC_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = EX_VAR(opline.GetOp1().GetVar())
	EX_VAR(opline.GetResult().GetVar()).SetDouble(value.GetDval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_QM_ASSIGN_NOREF_SPEC_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = EX_VAR(opline.GetOp1().GetVar())
	types.ZVAL_COPY_VALUE(EX_VAR(opline.GetResult().GetVar()), value)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ADD_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			result = EX_VAR(opline.GetResult().GetVar())
			FastLongAddFunction(result, op1, op2)
			ZEND_VM_NEXT_OPCODE()
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto add_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		add_double:
			result = EX_VAR(opline.GetResult().GetVar())
			result.SetDouble(d1 + d2)
			ZEND_VM_NEXT_OPCODE()
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto add_double
		}
	}
	return zend_add_helper_SPEC(op1, op2, executeData)
}
func ZEND_SUB_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			result = EX_VAR(opline.GetResult().GetVar())
			FastLongSubFunction(result, op1, op2)
			ZEND_VM_NEXT_OPCODE()
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto sub_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		sub_double:
			result = EX_VAR(opline.GetResult().GetVar())
			result.SetDouble(d1 - d2)
			ZEND_VM_NEXT_OPCODE()
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto sub_double
		}
	}
	return zend_sub_helper_SPEC(op1, op2, executeData)
}
func ZEND_MUL_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			var overflow ZendLong
			result = EX_VAR(opline.GetResult().GetVar())
			ZEND_SIGNED_MULTIPLY_LONG(op1.GetLval(), op2.GetLval(), result.GetLval(), result.GetDval(), overflow)
			if overflow != 0 {
				result.SetTypeInfo(types.IS_DOUBLE)
			} else {
				result.SetTypeInfo(types.IS_LONG)
			}
			ZEND_VM_NEXT_OPCODE()
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto mul_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		mul_double:
			result = EX_VAR(opline.GetResult().GetVar())
			result.SetDouble(d1 * d2)
			ZEND_VM_NEXT_OPCODE()
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto mul_double
		}
	}
	return zend_mul_helper_SPEC(op1, op2, executeData)
}
func ZEND_MOD_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			result = EX_VAR(opline.GetResult().GetVar())
			if op2.GetLval() == 0 {
				return zend_mod_by_zero_helper_SPEC(executeData)
			} else if op2.GetLval() == -1 {

				/* Prevent overflow error/crash if op1==ZEND_LONG_MIN */

				result.SetLong(0)

				/* Prevent overflow error/crash if op1==ZEND_LONG_MIN */

			} else {
				result.SetLong(op1.GetLval() % op2.GetLval())
			}
			ZEND_VM_NEXT_OPCODE()
		}
	}
	return zend_mod_helper_SPEC(op1, op2, executeData)
}
func ZEND_SL_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.GetLval() < SIZEOF_ZEND_LONG*8) != 0 {

		/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

		EX_VAR(opline.GetResult().GetVar()).SetLong(zend_long(ZendUlong(op1.GetLval() << op2.GetLval())))
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_shift_left_helper_SPEC(op1, op2, executeData)
}
func ZEND_SR_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.GetLval() < SIZEOF_ZEND_LONG*8) != 0 {
		EX_VAR(opline.GetResult().GetVar()).SetLong(op1.GetLval() >> op2.GetLval())
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_shift_right_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() < op2.GetLval() {
			is_smaller_true:
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_smaller_false:
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
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
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() < op2.GetLval() {
			is_smaller_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_smaller_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
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
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() < op2.GetLval() {
			is_smaller_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_smaller_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
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
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() <= op2.GetLval() {
			is_smaller_or_equal_true:
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_smaller_or_equal_false:
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
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
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() <= op2.GetLval() {
			is_smaller_or_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_smaller_or_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
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
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() <= op2.GetLval() {
			is_smaller_or_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_smaller_or_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
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
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() && op2.IsLong() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(op1.GetLval() | op2.GetLval())
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_bw_or_helper_SPEC(op1, op2, executeData)
}
func ZEND_BW_AND_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() && op2.IsLong() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(op1.GetLval() & op2.GetLval())
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_bw_and_helper_SPEC(op1, op2, executeData)
}
func ZEND_BW_XOR_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() && op2.IsLong() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(op1.GetLval() ^ op2.GetLval())
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_bw_xor_helper_SPEC(op1, op2, executeData)
}
func ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	zend_fetch_dimension_address_LIST_r(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_SWITCH_LONG_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op *types.Zval
	var jump_zv *types.Zval
	var jumptable *types.Array
	op = EX_VAR(opline.GetOp1().GetVar())
	jumptable = RT_CONSTANT(opline, opline.GetOp2()).GetArr()
	if op.GetType() != types.IS_LONG {
		op = types.ZVAL_DEREF(op)
		if op.GetType() != types.IS_LONG {

			/* Wrong type, fall back to ZEND_CASE chain */

			ZEND_VM_NEXT_OPCODE()

			/* Wrong type, fall back to ZEND_CASE chain */

		}
	}
	jump_zv = jumptable.IndexFind(op.GetLval())
	if jump_zv != nil {
		ZEND_VM_SET_RELATIVE_OPCODE(opline, jump_zv.GetLval())
		return 0
	} else {

		/* default */

		ZEND_VM_SET_RELATIVE_OPCODE(opline, opline.GetExtendedValue())
		return 0
	}
}
func ZEND_SWITCH_STRING_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op *types.Zval
	var jump_zv *types.Zval
	var jumptable *types.Array
	op = EX_VAR(opline.GetOp1().GetVar())
	jumptable = RT_CONSTANT(opline, opline.GetOp2()).GetArr()
	if op.GetType() != types.IS_STRING {

		/* Wrong type, fall back to ZEND_CASE chain */

		/* Wrong type, fall back to ZEND_CASE chain */

		{
			op = types.ZVAL_DEREF(op)
			if op.GetType() != types.IS_STRING {

				/* Wrong type, fall back to ZEND_CASE chain */

				ZEND_VM_NEXT_OPCODE()

				/* Wrong type, fall back to ZEND_CASE chain */

			}
		}
	}
	jump_zv = jumptable.KeyFind(op.GetStr().GetStr())
	if jump_zv != nil {
		ZEND_VM_SET_RELATIVE_OPCODE(opline, jump_zv.GetLval())
		return 0
	} else {

		/* default */

		ZEND_VM_SET_RELATIVE_OPCODE(opline, opline.GetExtendedValue())
		return 0
	}
}
func ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	result.SetLong(op1.GetLval() + op2.GetLval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ADD_LONG_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	FastLongAddFunction(result, op1, op2)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ADD_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	result.SetDouble(op1.GetDval() + op2.GetDval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	result.SetLong(op1.GetLval() - op2.GetLval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SUB_LONG_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	FastLongSubFunction(result, op1, op2)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SUB_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	result.SetDouble(op1.GetDval() - op2.GetDval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	result.SetLong(op1.GetLval() * op2.GetLval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_MUL_LONG_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	var overflow ZendLong
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	ZEND_SIGNED_MULTIPLY_LONG(op1.GetLval(), op2.GetLval(), result.GetLval(), result.GetDval(), overflow)
	if overflow != 0 {
		result.SetTypeInfo(types.IS_DOUBLE)
	} else {
		result.SetTypeInfo(types.IS_LONG)
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_MUL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = EX_VAR(opline.GetResult().GetVar())
	result.SetDouble(op1.GetDval() * op2.GetDval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetLval() == op2.GetLval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetLval() == op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetLval() == op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetDval() == op2.GetDval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetDval() == op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetDval() == op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetLval() != op2.GetLval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetLval() != op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
