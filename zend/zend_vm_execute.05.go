package zend

import (
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = RT_CONSTANT(opline, opline.GetOp1())
	subject = RT_CONSTANT(opline, opline.GetOp2())
	if subject.IsArray() {
	array_key_exists_array:
		ht = subject.GetArr()
		result = ZendArrayKeyExistsFast(ht, key, opline, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, opline, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.GetResultZval().SetTypeInfo(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_DECLARE_CLASS_DELAYED_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var lcname *types.Zval
	var zv *types.Zval
	var ce *types.ClassEntry
	ce = CACHED_PTR(opline.GetExtendedValue())
	if ce == nil {
		lcname = RT_CONSTANT(opline, opline.GetOp1())
		zv = EG__().GetClassTable().KeyFind((lcname + 1).GetStr().GetStr())
		if zv != nil {
			ce = zv.GetCe()
			zv = types.ZendHashSetBucketKey(EG__().GetClassTable(), (*types.Bucket)(zv), lcname.GetStr().GetStr())
			if zv == nil {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ce.GetName().GetVal())
			} else {
				if ZendDoLinkClass(ce, RT_CONSTANT(opline, opline.GetOp2()).GetStr()) == types.FAILURE {

					/* Reload bucket pointer, the hash table may have been reallocated */

					zv = EG__().GetClassTable().KeyFind(lcname.GetStr().GetStr())
					types.ZendHashSetBucketKey(EG__().GetClassTable(), (*types.Bucket)(zv), (lcname + 1).GetStr().GetStr())
					return 0
				}
			}
		}
		CACHE_PTR(opline.GetExtendedValue(), ce)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_DECLARE_CONST_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var name *types.Zval
	var val *types.Zval
	var c ZendConstant
	name = RT_CONSTANT(opline, opline.GetOp1())
	val = RT_CONSTANT(opline, opline.GetOp2())
	types.ZVAL_COPY(c.Value(), val)
	if c.Value().IsConstant() {
		if ZvalUpdateConstantEx(c.Value(), executeData.GetFunc().op_array.scope) != types.SUCCESS {
			ZvalPtrDtorNogc(c.Value())
			return 0
		}
	}

	/* non persistent, case sensitive */

	ZEND_CONSTANT_SET_FLAGS(&c, CONST_CS, PHP_USER_CONSTANT)
	c.SetName(name.GetStr().Copy())
	if ZendRegisterConstant(&c) == types.FAILURE {
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_YIELD_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
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
		var key *types.Zval = RT_CONSTANT(opline, opline.GetOp2())

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)

			generator.GetKey().TryAddRefcount()

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
func ZEND_SWITCH_LONG_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op *types.Zval
	var jump_zv *types.Zval
	var jumptable *types.Array
	op = RT_CONSTANT(opline, opline.GetOp1())
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
func ZEND_SWITCH_STRING_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op *types.Zval
	var jump_zv *types.Zval
	var jumptable *types.Array
	op = RT_CONSTANT(opline, opline.GetOp1())
	jumptable = RT_CONSTANT(opline, opline.GetOp2()).GetArr()
	if op.GetType() != types.IS_STRING {
		{

			/* Wrong type, fall back to ZEND_CASE chain */

			return ZEND_VM_NEXT_OPCODE(executeData, opline)

			/* Wrong type, fall back to ZEND_CASE chain */

		}

		/* Wrong type, fall back to ZEND_CASE chain */

		/* Wrong type, fall back to ZEND_CASE chain */

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
func ZEND_IN_ARRAY_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var ht *types.Array = RT_CONSTANT(opline, opline.GetOp2()).GetArr()
	var result *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	if op1.IsString() {
		result = ht.KeyFind(op1.GetStr().GetStr())
	} else if opline.GetExtendedValue() != 0 {
		if op1.IsLong() {
			result = ht.IndexFind(op1.GetLval())
		} else {
			result = nil
		}
	} else if op1.GetType() <= types.IS_FALSE {
		result = ht.KeyFind(types.ZSTR_EMPTY_ALLOC().GetStr())
	} else {
		var key *types.String
		var key_tmp types.Zval
		var result_tmp types.Zval
		var val *types.Zval
		result = nil
		var __ht *types.Array = ht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			val = _z
			key_tmp.SetString(key)
			CompareFunction(&result_tmp, op1, &key_tmp)
			if result_tmp.GetLval() == 0 {
				result = val
				break
			}
		}
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != nil)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ADD_SPEC_CONST_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = opline.GetOp2Zval()

	if op1.IsLong() {
		if op2.IsLong() {
			result = opline.GetResultZval()
			FastLongAddFunction(result, op1, op2)
			return ZEND_VM_NEXT_OPCODE(executeData, opline)
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
			result = opline.GetResultZval()
			result.SetDouble(d1 + d2)
			return ZEND_VM_NEXT_OPCODE(executeData, opline)
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto add_double
		}
	}
	return zend_add_helper_SPEC(op1, op2, executeData)
}
func ZEND_SUB_SPEC_CONST_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = opline.GetOp2Zval()

	if op1.IsLong() {
		if op2.IsLong() {
			result = opline.GetResultZval()
			FastLongSubFunction(result, op1, op2)
			return ZEND_VM_NEXT_OPCODE(executeData, opline)
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
			result = opline.GetResultZval()
			result.SetDouble(d1 - d2)
			return ZEND_VM_NEXT_OPCODE(executeData, opline)
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto sub_double
		}
	}
	return zend_sub_helper_SPEC(op1, op2, executeData)
}
func ZEND_MOD_SPEC_CONST_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = opline.GetOp2Zval()

	if op1.IsLong() {
		if op2.IsLong() {
			result = opline.GetResultZval()
			if op2.GetLval() == 0 {
				return zend_mod_by_zero_helper_SPEC(executeData)
			} else if op2.GetLval() == -1 {

				/* Prevent overflow error/crash if op1==ZEND_LONG_MIN */

				result.SetLong(0)

				/* Prevent overflow error/crash if op1==ZEND_LONG_MIN */

			} else {
				result.SetLong(op1.GetLval() % op2.GetLval())
			}
			return ZEND_VM_NEXT_OPCODE(executeData, opline)
		}
	}
	return zend_mod_helper_SPEC(op1, op2, executeData)
}
func ZEND_SL_SPEC_CONST_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = opline.GetOp2Zval()

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.GetLval() < SIZEOF_ZEND_LONG*8) != 0 {

		/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

		opline.GetResultZval().SetLong(zend_long(ZendUlong(op1.GetLval() << op2.GetLval())))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_shift_left_helper_SPEC(op1, op2, executeData)
}
func ZEND_SR_SPEC_CONST_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = opline.GetOp2Zval()

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.GetLval() < SIZEOF_ZEND_LONG*8) != 0 {
		opline.GetResultZval().SetLong(op1.GetLval() >> op2.GetLval())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_shift_right_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = opline.GetOp2Zval()

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
func ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = opline.GetOp2Zval()

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
func ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = opline.GetOp2Zval()

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
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = opline.GetOp2Zval()

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
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = opline.GetOp2Zval()

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
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = opline.GetOp2Zval()

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
func ZEND_DIV_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	FastDivFunction(opline.GetResultZval(), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_POW_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	PowFunction(opline.GetResultZval(), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_CONCAT_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	if op2.IsString() {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String

		if op2_str.GetLen() == 0 {
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
				types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		if op2.IsUndef() {
			op2 = ZVAL_UNDEFINED_OP2()
		}
		ConcatFunction(opline.GetResultZval(), op1, op2)
		ZvalPtrDtorNogc(free_op2)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
}
func ZEND_SPACESHIP_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	CompareFunction(opline.GetResultZval(), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_R_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

	{
		zend_fetch_dimension_address_read_R(container, dim, IS_TMP_VAR|IS_VAR, opline, executeData)
	}
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_IS_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	zend_fetch_dimension_address_read_IS(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData), IS_TMP_VAR|IS_VAR, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_CONST_TMPVAR_HANDLER(executeData)
	}
}
