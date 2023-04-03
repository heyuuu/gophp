package zend

import "github.com/heyuuu/gophp/zend/types"

// ZEND_ADD
func vmAddHandler(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval = opline.GetOp1ZvalEx()
	var op2 *types.Zval = opline.GetOp2ZvalEx()

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
	var op1 *types.Zval = opline.GetOp1ZvalEx()
	var op2 *types.Zval = opline.GetOp2ZvalEx()

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
