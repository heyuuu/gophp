package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_IS_SMALLER_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = executeData.GetOp1(opline)
	op2 = executeData.GetOp2(opline)
	return zend_is_smaller_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = executeData.GetOp1(opline)
	op2 = opline.Op2()

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.Long() < op2.Long() {
			is_smaller_true:
				opline.Result().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_false:
				opline.Result().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.Long())
			d2 = op2.Double()
			goto is_smaller_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.Double()
			d2 = op2.Double()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.IsLong() {
			d1 = op1.Double()
			d2 = float64(op2.Long())
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
	op1 = executeData.GetOp1(opline)
	op2 = opline.Op2()

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.Long() < op2.Long() {
			is_smaller_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				opline.Result().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				opline.Result().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.Long())
			d2 = op2.Double()
			goto is_smaller_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.Double()
			d2 = op2.Double()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.IsLong() {
			d1 = op1.Double()
			d2 = float64(op2.Long())
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
	op1 = executeData.GetOp1(opline)
	op2 = opline.Op2()

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.Long() < op2.Long() {
			is_smaller_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
				opline.Result().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
				opline.Result().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.Long())
			d2 = op2.Double()
			goto is_smaller_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.Double()
			d2 = op2.Double()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.IsLong() {
			d1 = op1.Double()
			d2 = float64(op2.Long())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = opline.Op1()
	op2 = executeData.GetOp2(opline)

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.Long() < op2.Long() {
			is_smaller_true:
				opline.Result().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_false:
				opline.Result().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.Long())
			d2 = op2.Double()
			goto is_smaller_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.Double()
			d2 = op2.Double()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.IsLong() {
			d1 = op1.Double()
			d2 = float64(op2.Long())
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
	op1 = opline.Op1()
	op2 = executeData.GetOp2(opline)

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.Long() < op2.Long() {
			is_smaller_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				opline.Result().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				opline.Result().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.Long())
			d2 = op2.Double()
			goto is_smaller_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.Double()
			d2 = op2.Double()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.IsLong() {
			d1 = op1.Double()
			d2 = float64(op2.Long())
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
	op1 = opline.Op1()
	op2 = executeData.GetOp2(opline)

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.Long() < op2.Long() {
			is_smaller_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
				opline.Result().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
				opline.Result().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.Long())
			d2 = op2.Double()
			goto is_smaller_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.Double()
			d2 = op2.Double()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.IsLong() {
			d1 = op1.Double()
			d2 = float64(op2.Long())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = opline.Op1()
	op2 = opline.Op2()

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.Long() < op2.Long() {
			is_smaller_true:
				opline.Result().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_false:
				opline.Result().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.Long())
			d2 = op2.Double()
			goto is_smaller_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.Double()
			d2 = op2.Double()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.IsLong() {
			d1 = op1.Double()
			d2 = float64(op2.Long())
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
	op1 = opline.Op1()
	op2 = opline.Op2()

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.Long() < op2.Long() {
			is_smaller_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				opline.Result().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				opline.Result().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.Long())
			d2 = op2.Double()
			goto is_smaller_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.Double()
			d2 = op2.Double()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.IsLong() {
			d1 = op1.Double()
			d2 = float64(op2.Long())
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
	op1 = opline.Op1()
	op2 = opline.Op2()

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.Long() < op2.Long() {
			is_smaller_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
				opline.Result().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			is_smaller_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
				opline.Result().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.Long())
			d2 = op2.Double()
			goto is_smaller_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.Double()
			d2 = op2.Double()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.IsLong() {
			d1 = op1.Double()
			d2 = float64(op2.Long())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, executeData)
}
