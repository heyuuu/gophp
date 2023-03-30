package zend

import (
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetLval() != op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetDval() != op2.GetDval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetDval() != op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetDval() != op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetLval() < op2.GetLval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetLval() < op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetLval() < op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetDval() < op2.GetDval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetDval() < op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetDval() < op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetLval() <= op2.GetLval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetLval() <= op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetLval() <= op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetDval() <= op2.GetDval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetDval() <= op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = op1.GetDval() <= op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

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
func ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

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
func ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

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
func ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

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
func ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.GetLval() < SIZEOF_ZEND_LONG*8) != 0 {

		/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

		EX_VAR(opline.GetResult().GetVar()).SetLong(zend_long(ZendUlong(op1.GetLval() << op2.GetLval())))
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_shift_left_helper_SPEC(op1, op2, executeData)
}
func ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.GetLval() < SIZEOF_ZEND_LONG*8) != 0 {
		EX_VAR(opline.GetResult().GetVar()).SetLong(op1.GetLval() >> op2.GetLval())
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_shift_right_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

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
func ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

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
func ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

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
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

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
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

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
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

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
func ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

	if op1.IsLong() && op2.IsLong() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(op1.GetLval() | op2.GetLval())
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_bw_or_helper_SPEC(op1, op2, executeData)
}
func ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

	if op1.IsLong() && op2.IsLong() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(op1.GetLval() & op2.GetLval())
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_bw_and_helper_SPEC(op1, op2, executeData)
}
func ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())

	if op1.IsLong() && op2.IsLong() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(op1.GetLval() ^ op2.GetLval())
		ZEND_VM_NEXT_OPCODE()
	}
	return zend_bw_xor_helper_SPEC(op1, op2, executeData)
}
func ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = EX_VAR(opline.GetResult().GetVar())
	result.SetLong(op1.GetLval() + op2.GetLval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = EX_VAR(opline.GetResult().GetVar())
	FastLongAddFunction(result, op1, op2)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = EX_VAR(opline.GetResult().GetVar())
	result.SetDouble(op1.GetDval() + op2.GetDval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = EX_VAR(opline.GetResult().GetVar())
	result.SetLong(op1.GetLval() - op2.GetLval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = EX_VAR(opline.GetResult().GetVar())
	FastLongSubFunction(result, op1, op2)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = EX_VAR(opline.GetResult().GetVar())
	result.SetDouble(op1.GetDval() - op2.GetDval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = EX_VAR(opline.GetResult().GetVar())
	result.SetLong(op1.GetLval() * op2.GetLval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	var overflow ZendLong
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = EX_VAR(opline.GetResult().GetVar())
	ZEND_SIGNED_MULTIPLY_LONG(op1.GetLval(), op2.GetLval(), result.GetLval(), result.GetDval(), overflow)
	if overflow != 0 {
		result.SetTypeInfo(types.IS_DOUBLE)
	} else {
		result.SetTypeInfo(types.IS_LONG)
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = EX_VAR(opline.GetResult().GetVar())
	result.SetDouble(op1.GetDval() * op2.GetDval())
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetLval() == op2.GetLval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetLval() == op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetLval() == op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetDval() == op2.GetDval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetDval() == op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetDval() == op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetLval() != op2.GetLval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetLval() != op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetLval() != op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetDval() != op2.GetDval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetDval() != op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetDval() != op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetLval() < op2.GetLval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetLval() < op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetLval() < op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetDval() < op2.GetDval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetDval() < op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetDval() < op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetLval() <= op2.GetLval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetLval() <= op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
