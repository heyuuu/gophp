package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

// ZEND_NOP
func vmNopHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

// ZEND_ADD
func vmAddHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)

	// fast
	switch operators.TypePair(op1.GetType(), op2.GetType()) {
	case operators.TypePair(types.IsLong, types.IsLong):
		result := opline.Result()
		operators.FastLongAddFunction(result, op1, op2)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	case operators.TypePair(types.IsDouble, types.IsDouble),
		operators.TypePair(types.IsLong, types.IsDouble),
		operators.TypePair(types.IsDouble, types.IsLong):
		var d1, d2 float64
		if op1.IsLong() {
			d1 = float64(op1.Long())
		} else {
			d1 = op1.Double()
		}
		if op2.IsLong() {
			d2 = float64(op2.Long())
		} else {
			d2 = op2.Double()
		}
		result := opline.Result()
		result.SetDouble(d1 + d2)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	// zend_add_helper_SPEC
	if op1.IsUndef() {
		op1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op2.IsUndef() {
		op2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.AddFunction(opline.Result(), op1, op2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		// ZvalPtrDtorNogc(op1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		// ZvalPtrDtorNogc(op2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

// ZEND_SUB
func vmSubHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)

	// fast
	switch operators.TypePair(op1.GetType(), op2.GetType()) {
	case operators.TypePair(types.IsLong, types.IsLong):
		result := opline.Result()
		operators.FastLongSubFunction(result, op1, op2)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	case operators.TypePair(types.IsDouble, types.IsDouble),
		operators.TypePair(types.IsLong, types.IsDouble),
		operators.TypePair(types.IsDouble, types.IsLong):
		var d1, d2 float64
		if op1.IsLong() {
			d1 = float64(op1.Long())
		} else {
			d1 = op1.Double()
		}
		if op2.IsLong() {
			d2 = float64(op2.Long())
		} else {
			d2 = op2.Double()
		}
		result := opline.Result()
		result.SetDouble(d1 - d2)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	//  zend_sub_helper_SPEC
	if op1.IsUndef() {
		op1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op2.IsUndef() {
		op2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.SubFunction(opline.Result(), op1, op2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		// ZvalPtrDtorNogc(op1)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		// ZvalPtrDtorNogc(op2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

// ZEND_MUL
func vmMulHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)

	// fast
	switch operators.TypePair(op1.GetType(), op2.GetType()) {
	case operators.TypePair(types.IsLong, types.IsLong):
		result := opline.Result()
		if iVal, dVal, overflow := SignedMultiplyLong(op1.Long(), op2.Long()); overflow {
			result.SetDouble(dVal)
		} else {
			result.SetLong(iVal)
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	case operators.TypePair(types.IsDouble, types.IsDouble),
		operators.TypePair(types.IsLong, types.IsDouble),
		operators.TypePair(types.IsDouble, types.IsLong):
		var d1, d2 float64
		if op1.IsLong() {
			d1 = float64(op1.Long())
		} else {
			d1 = op1.Double()
		}
		if op2.IsLong() {
			d2 = float64(op2.Long())
		} else {
			d2 = op2.Double()
		}
		result := opline.Result()
		result.SetDouble(d1 * d2)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	// common
	if op1.IsUndef() {
		op1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op2.IsUndef() {
		op2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.MulFunction(opline.Result(), op1, op2)
	if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		// ZvalPtrDtorNogc(op2)
	}
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		// ZvalPtrDtorNogc(op2)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

// ZEND_DIV
func vmDivHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var freeOp1, freeOp2 ZendFreeOp
	var op1 *types.Zval = executeData.Op1(opline, opMode2)
	var op2 *types.Zval = executeData.Op2(opline, opMode2)
	operators.FastDivFunction(opline.Result(), op1, op2)

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

// ZEND_MOD
func vmModHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)

	// fast
	if op1.IsLong() && op2.IsLong() {
		result := opline.Result()
		if op2.Long() == 0 {
			return zend_mod_by_zero_helper_SPEC(executeData)
		} else if op2.Long() == -1 {
			/* Prevent overflow error/crash if op1==ZEND_LONG_MIN */
			result.SetLong(0)
		} else {
			result.SetLong(op1.Long() % op2.Long())
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	return zend_mod_helper_SPEC(op1, op2, executeData)
}

// ZEND_SL
func vmSlHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.Long() < SIZEOF_ZEND_LONG*8) != 0 {
		/* Perform shift on unsigned numbers to get well-defined wrap behavior. */
		opline.Result().SetLong(ZendLong(ZendUlong(op1.Long() << op2.Long())))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_shift_left_helper_SPEC(op1, op2, executeData)
}

// ZEND_SR
func getSrHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)

	if op1.IsLong() && op2.IsLong() && ZendUlong(op2.Long() < SIZEOF_ZEND_LONG*8) != 0 {
		opline.Result().SetLong(op1.Long() >> op2.Long())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	return zend_shift_right_helper_SPEC(op1, op2, executeData)
}

// ZEND_CONCAT
func getConcatHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var freeOp1, freeOp2 ZendFreeOp
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)

	// fast
	if op1.IsString() && op2.IsString() {
		opline.Result().SetStringVal(op1.StringVal() + op2.StringVal())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	// common
	if op1.IsUndef() {
		op1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	if op2.IsUndef() {
		op2 = ZVAL_UNDEFINED_OP2(executeData)
	}
	operators.ConcatFunction(opline.Result(), op1, op2)

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
