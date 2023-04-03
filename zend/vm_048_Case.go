package zend

func ZEND_CASE_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = opline.getZvalPtrVar1(&free_op1)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() == op2.GetLval() {
			case_true:
				ZEND_VM_SMART_BRANCH_TRUE()
				opline.Result().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			case_false:
				ZEND_VM_SMART_BRANCH_FALSE()
				opline.Result().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto case_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		case_double:
			if d1 == d2 {
				goto case_true
			} else {
				goto case_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto case_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			if result != 0 {
				goto case_true
			} else {
				goto case_false
			}
		}
	}
	return zend_case_helper_SPEC(op1, op2, executeData)
}
func ZEND_CASE_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = opline.getZvalPtrVar1(&free_op1)
	op2 = opline.getZvalPtrVar2(&free_op2)
	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() == op2.GetLval() {
			case_true:
				ZEND_VM_SMART_BRANCH_TRUE()
				opline.Result().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			case_false:
				ZEND_VM_SMART_BRANCH_FALSE()
				opline.Result().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto case_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		case_double:
			if d1 == d2 {
				goto case_true
			} else {
				goto case_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto case_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			ZvalPtrDtorNogc(free_op2)
			if result != 0 {
				goto case_true
			} else {
				goto case_false
			}
		}
	}
	return zend_case_helper_SPEC(op1, op2, executeData)
}
func ZEND_CASE_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = opline.getZvalPtrVar1(&free_op1)
	op2 = opline.Op2()
	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() == op2.GetLval() {
			case_true:
				ZEND_VM_SMART_BRANCH_TRUE()
				opline.Result().SetTrue()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			} else {
			case_false:
				ZEND_VM_SMART_BRANCH_FALSE()
				opline.Result().SetFalse()
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto case_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		case_double:
			if d1 == d2 {
				goto case_true
			} else {
				goto case_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto case_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			if result != 0 {
				goto case_true
			} else {
				goto case_false
			}
		}
	}
	return zend_case_helper_SPEC(op1, op2, executeData)
}
