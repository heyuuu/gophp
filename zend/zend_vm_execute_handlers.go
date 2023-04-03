package zend

import "github.com/heyuuu/gophp/zend/types"

// ZEND_NOP
func vmNopHandler(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

// ZEND_ADD
func vmAddHandler(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval = opline.Op1Ex()
	var op2 *types.Zval = opline.Op2Ex()

	// fast
	switch TYPE_PAIR(op1.GetType(), op2.GetType()) {
	case TYPE_PAIR(types.IS_LONG, types.IS_LONG):
		result := opline.GetResultZval()
		FastLongAddFunction(result, op1, op2)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	case TYPE_PAIR(types.IS_DOUBLE, types.IS_DOUBLE),
		TYPE_PAIR(types.IS_LONG, types.IS_DOUBLE),
		TYPE_PAIR(types.IS_DOUBLE, types.IS_LONG):
		var d1, d2 float64
		if op1.IsLong() {
			d1 = float64(op1.GetLval())
		} else {
			d1 = op1.GetDval()
		}
		if op2.IsLong() {
			d2 = float64(op2.GetLval())
		} else {
			d2 = op2.GetDval()
		}
		result := opline.GetResultZval()
		result.SetDouble(d1 + d2)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	// zend_add_helper_SPEC
	if op1.IsUndef() {
		op1 = ZVAL_UNDEFINED_OP1()
	}
	if op2.IsUndef() {
		op2 = ZVAL_UNDEFINED_OP2()
	}
	AddFunction(opline.GetResultZval(), op1, op2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

// ZEND_SUB
func vmSubHandler(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval = opline.Op1Ex()
	var op2 *types.Zval = opline.Op2Ex()

	// fast
	switch TYPE_PAIR(op1.GetType(), op2.GetType()) {
	case TYPE_PAIR(types.IS_LONG, types.IS_LONG):
		result := opline.GetResultZval()
		FastLongSubFunction(result, op1, op2)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	case TYPE_PAIR(types.IS_DOUBLE, types.IS_DOUBLE),
		TYPE_PAIR(types.IS_LONG, types.IS_DOUBLE),
		TYPE_PAIR(types.IS_DOUBLE, types.IS_LONG):
		var d1, d2 float64
		if op1.IsLong() {
			d1 = float64(op1.GetLval())
		} else {
			d1 = op1.GetDval()
		}
		if op2.IsLong() {
			d2 = float64(op2.GetLval())
		} else {
			d2 = op2.GetDval()
		}
		result := opline.GetResultZval()
		result.SetDouble(d1 - d2)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	//  zend_sub_helper_SPEC
	if op1.IsUndef() {
		op1 = ZVAL_UNDEFINED_OP1()
	}
	if op2.IsUndef() {
		op2 = ZVAL_UNDEFINED_OP2()
	}
	SubFunction(opline.GetResultZval(), op1, op2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

// ZEND_MUL
func vmMulHandler(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval = opline.Op1Ex()
	var op2 *types.Zval = opline.Op2Ex()

	// fast
	switch TYPE_PAIR(op1.GetType(), op2.GetType()) {
	case TYPE_PAIR(types.IS_LONG, types.IS_LONG):
		var overflow ZendLong
		result := opline.GetResultZval()
		ZEND_SIGNED_MULTIPLY_LONG(op1.GetLval(), op2.GetLval(), result.GetLval(), result.GetDval(), overflow)
		if overflow != 0 {
			result.SetTypeInfo(types.IS_DOUBLE)
		} else {
			result.SetTypeInfo(types.IS_LONG)
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	case TYPE_PAIR(types.IS_DOUBLE, types.IS_DOUBLE),
		TYPE_PAIR(types.IS_LONG, types.IS_DOUBLE),
		TYPE_PAIR(types.IS_DOUBLE, types.IS_LONG):
		var d1, d2 float64
		if op1.IsLong() {
			d1 = float64(op1.GetLval())
		} else {
			d1 = op1.GetDval()
		}
		if op2.IsLong() {
			d2 = float64(op2.GetLval())
		} else {
			d2 = op2.GetDval()
		}
		result := opline.GetResultZval()
		result.SetDouble(d1 * d2)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	// common
	if op1.IsUndef() {
		op1 = ZVAL_UNDEFINED_OP1()
	}
	if op2.IsUndef() {
		op2 = ZVAL_UNDEFINED_OP2()
	}
	MulFunction(opline.GetResultZval(), op1, op2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op2)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

// ZEND_DIV
func vmDivHandler(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var freeOp1, freeOp2 ZendFreeOp
	var op1 *types.Zval = opline.Op1ExEx(&freeOp1)
	var op2 *types.Zval = opline.Op2ExEx(&freeOp2)
	FastDivFunction(opline.GetResultZval(), op1, op2)
	if freeOp1 != nil {
		ZvalPtrDtorNogc(freeOp1)
	}
	if freeOp2 != nil {
		ZvalPtrDtorNogc(freeOp2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

// ZEND_MOD
func vmModHandler(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval = opline.Op1Ex()
	var op2 *types.Zval = opline.Op2Ex()

	// fast
	if op1.IsLong() && op2.IsLong() {
		result := opline.GetResultZval()
		if op2.GetLval() == 0 {
			return zend_mod_by_zero_helper_SPEC(executeData)
		} else if op2.GetLval() == -1 {
			/* Prevent overflow error/crash if op1==ZEND_LONG_MIN */
			result.SetLong(0)
		} else {
			result.SetLong(op1.GetLval() % op2.GetLval())
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	return zend_mod_helper_SPEC(op1, op2, executeData)
}

// ZEND_SL
func vmSlHandler(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval = opline.Op1Ex()
	var op2 *types.Zval = opline.Op2Ex()

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.GetLval() < SIZEOF_ZEND_LONG*8) != 0 {
		/* Perform shift on unsigned numbers to get well-defined wrap behavior. */
		opline.GetResultZval().SetLong(ZendLong(ZendUlong(op1.GetLval() << op2.GetLval())))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_shift_left_helper_SPEC(op1, op2, executeData)
}

// ZEND_SR
func getSrHandler(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval = opline.Op1Ex()
	var op2 *types.Zval = opline.Op2Ex()

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.GetLval() < SIZEOF_ZEND_LONG*8) != 0 {
		opline.GetResultZval().SetLong(op1.GetLval() >> op2.GetLval())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	return zend_shift_right_helper_SPEC(op1, op2, executeData)
}
