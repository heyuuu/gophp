package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

// ZEND_BW_OR
func getBwOrHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)

	if op1.IsLong() && op2.IsLong() {
		opline.Result().SetLong(op1.Long() | op2.Long())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	return zend_bw_or_helper_SPEC(op1, op2, executeData)
}

// ZEND_BW_AND
func getBwAndHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)

	if op1.IsLong() && op2.IsLong() {
		opline.Result().SetLong(op1.Long() & op2.Long())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_bw_and_helper_SPEC(op1, op2, executeData)
}

// ZEND_BW_XOR
func getBwXorHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)

	if op1.IsLong() && op2.IsLong() {
		opline.Result().SetLong(op1.Long() ^ op2.Long())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	return zend_bw_xor_helper_SPEC(op1, op2, executeData)
}

// ZEND_POW
func getPowHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode2)
	var op2 *types.Zval = executeData.Op2(opline, opMode2)

	operators.PowFunction(opline.Result(), op1, op2)

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

// ZEND_BW_NOT
func getBwNotHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)

	if op1.IsLong() {
		opline.Result().SetLong(^(op1.Long()))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	if op1.IsUndef() {
		op1 = ZVAL_UNDEFINED_OP1(executeData)
	}
	operators.BitwiseNotFunction(opline.Result(), op1)

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

// ZEND_BOOL_NOT
func getBoolNotHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)

	if op1.IsTrue() {
		opline.Result().SetFalse()
	} else if op1.IsSignType() {
		opline.Result().SetTrue()
		if op1.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	} else {
		opline.Result().SetBool(!operators.ZvalIsTrue(op1))
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

// ZEND_BOOL_XOR
func getBoolXorHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode2)
	var op2 *types.Zval = executeData.Op2(opline, opMode2)

	operators.BooleanXorFunction(opline.Result(), op1, op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

// ZEND_IS_IDENTICAL
func getIsIdenticalHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode0)
	var op2 *types.Zval = executeData.Op2(opline, opMode0)

	result := operators.FastIsIdenticalFunction(op1, op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

// ZEND_IS_NOT_IDENTICAL
func getIsNotIdenticalHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode0)
	var op2 *types.Zval = executeData.Op2(opline, opMode0)

	result := operators.FastIsNotIdenticalFunction(op1, op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != 0)

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}

// ZEND_IS_EQUAL
func getIsEqualHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)

	var result bool
	switch operators.TypePair(op1.Type(), op2.Type()) {
	case operators.TypePair(types.IsLong, types.IsLong):
		result = op1.Long() == op2.Long()
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
		result = d1 == d2
	case operators.TypePair(types.IsString, types.IsString):
		result = operators.ZendFastEqualStringsEx(op1.String(), op2.String())
	default:
		return zend_is_equal_helper_SPEC(op1, op2, executeData)
	}

	nextOpCode := opline.Offset(1).GetOpcode()
	switch nextOpCode {
	case ZEND_JMPZ:
		if result {
			ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
		} else {
			ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
		}
	case ZEND_JMPNZ:
		if result {
			ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
		} else {
			ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
		}
	}

	opline.Result().SetBool(result)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

// ZEND_IS_NOT_EQUAL
func getIsNotEqualHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)

	var result bool
	switch operators.TypePair(op1.Type(), op2.Type()) {
	case operators.TypePair(types.IsLong, types.IsLong):
		result = op1.Long() != op2.Long()
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
		result = d1 != d2
	case operators.TypePair(types.IsString, types.IsString):
		result = !operators.ZendFastEqualStringsEx(op1.String(), op2.String())
	default:
		return zend_is_not_equal_helper_SPEC(op1, op2, executeData)
	}

	nextOpCode := opline.Offset(1).GetOpcode()
	switch nextOpCode {
	case ZEND_JMPZ:
		if result {
			ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
		} else {
			ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
		}
	case ZEND_JMPNZ:
		if result {
			ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
		} else {
			ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
		}
	}

	opline.Result().SetBool(result)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

// ZEND_IS_SMALLER
func getIsSmallerHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_IS_SMALLER_OR_EQUAL
func getIsSmallerOrEqualHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ASSIGN
func getAssignHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ASSIGN_DIM
func getAssignDimHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ASSIGN_OBJ
func getAssignObjHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ASSIGN_STATIC_PROP
func getAssignStaticPropHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ASSIGN_OP
func getAssignOpHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ASSIGN_DIM_OP
func getAssignDimOpHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ASSIGN_OBJ_OP
func getAssignObjOpHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ASSIGN_STATIC_PROP_OP
func getAssignStaticPropOpHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ASSIGN_REF
func getAssignRefHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_QM_ASSIGN
func getQmAssignHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ASSIGN_OBJ_REF
func getAssignObjRefHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ASSIGN_STATIC_PROP_REF
func getAssignStaticPropRefHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_PRE_INC
func getPreIncHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_PRE_DEC
func getPreDecHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_POST_INC
func getPostIncHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_POST_DEC
func getPostDecHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_PRE_INC_STATIC_PROP
func getPreIncStaticPropHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_PRE_DEC_STATIC_PROP
func getPreDecStaticPropHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_POST_INC_STATIC_PROP
func getPostIncStaticPropHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_POST_DEC_STATIC_PROP
func getPostDecStaticPropHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_JMP
func getJmpHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_JMPZ
func getJmpzHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_JMPNZ
func getJmpnzHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_JMPZNZ
func getJmpznzHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_JMPZ_EX
func getJmpzExHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_JMPNZ_EX
func getJmpnzExHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_CASE
func getCaseHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_CHECK_VAR
func getCheckVarHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SEND_VAR_NO_REF_EX
func getSendVarNoRefExHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_CAST
func getCastHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_BOOL
func getBoolHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FAST_CONCAT
func getFastConcatHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ROPE_INIT
func getRopeInitHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ROPE_ADD
func getRopeAddHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ROPE_END
func getRopeEndHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_BEGIN_SILENCE
func getBeginSilenceHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_END_SILENCE
func getEndSilenceHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_INIT_FCALL_BY_NAME
func getInitFcallByNameHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_DO_FCALL
func getDoFcallHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_INIT_FCALL
func getInitFcallHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_RETURN
func getReturnHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_RECV
func getRecvHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_RECV_INIT
func getRecvInitHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SEND_VAL
func getSendValHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SEND_VAR_EX
func getSendVarExHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SEND_REF
func getSendRefHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_NEW
func getNewHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_INIT_NS_FCALL_BY_NAME
func getInitNsFcallByNameHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FREE
func getFreeHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_INIT_ARRAY
func getInitArrayHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ADD_ARRAY_ELEMENT
func getAddArrayElementHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_INCLUDE_OR_EVAL
func getIncludeOrEvalHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_UNSET_VAR
func getUnsetVarHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_UNSET_DIM
func getUnsetDimHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_UNSET_OBJ
func getUnsetObjHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FE_RESET_R
func getFeResetRHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FE_FETCH_R
func getFeFetchRHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_EXIT
func getExitHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_R
func getFetchRHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_DIM_R
func getFetchDimRHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_OBJ_R
func getFetchObjRHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_W
func getFetchWHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_DIM_W
func getFetchDimWHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_OBJ_W
func getFetchObjWHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_RW
func getFetchRwHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_DIM_RW
func getFetchDimRwHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_OBJ_RW
func getFetchObjRwHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_IS
func getFetchIsHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_DIM_IS
func getFetchDimIsHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_OBJ_IS
func getFetchObjIsHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_FUNC_ARG
func getFetchFuncArgHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_DIM_FUNC_ARG
func getFetchDimFuncArgHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_OBJ_FUNC_ARG
func getFetchObjFuncArgHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_UNSET
func getFetchUnsetHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_DIM_UNSET
func getFetchDimUnsetHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_OBJ_UNSET
func getFetchObjUnsetHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_LIST_R
func getFetchListRHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_CONSTANT
func getFetchConstantHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_CHECK_FUNC_ARG
func getCheckFuncArgHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_EXT_STMT
func getExtStmtHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_EXT_FCALL_BEGIN
func getExtFcallBeginHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_EXT_FCALL_END
func getExtFcallEndHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_EXT_NOP
func getExtNopHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SEND_VAR_NO_REF
func getSendVarNoRefHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_CATCH
func getCatchHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_THROW
func getThrowHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_CLASS
func getFetchClassHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_CLONE
func getCloneHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_RETURN_BY_REF
func getReturnByRefHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_INIT_METHOD_CALL
func getInitMethodCallHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_INIT_STATIC_METHOD_CALL
func getInitStaticMethodCallHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ISSET_ISEMPTY_VAR
func getIssetIsemptyVarHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ISSET_ISEMPTY_DIM_OBJ
func getIssetIsemptyDimObjHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SEND_VAL_EX
func getSendValExHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SEND_VAR
func getSendVarHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_INIT_USER_CALL
func getInitUserCallHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SEND_ARRAY
func getSendArrayHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SEND_USER
func getSendUserHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_STRLEN
func getStrlenHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_DEFINED
func getDefinedHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_TYPE_CHECK
func getTypeCheckHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_VERIFY_RETURN_TYPE
func getVerifyReturnTypeHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FE_RESET_RW
func getFeResetRwHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FE_FETCH_RW
func getFeFetchRwHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FE_FREE
func getFeFreeHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_INIT_DYNAMIC_CALL
func getInitDynamicCallHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_DO_ICALL
func getDoIcallHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_DO_UCALL
func getDoUcallHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_DO_FCALL_BY_NAME
func getDoFcallByNameHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_PRE_INC_OBJ
func getPreIncObjHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_PRE_DEC_OBJ
func getPreDecObjHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_POST_INC_OBJ
func getPostIncObjHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_POST_DEC_OBJ
func getPostDecObjHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ECHO
func getEchoHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_OP_DATA
func getOpDataHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_INSTANCEOF
func getInstanceofHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_GENERATOR_CREATE
func getGeneratorCreateHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_MAKE_REF
func getMakeRefHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_DECLARE_FUNCTION
func getDeclareFunctionHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_DECLARE_LAMBDA_FUNCTION
func getDeclareLambdaFunctionHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_DECLARE_CONST
func getDeclareConstHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_DECLARE_CLASS
func getDeclareClassHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_DECLARE_CLASS_DELAYED
func getDeclareClassDelayedHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_DECLARE_ANON_CLASS
func getDeclareAnonClassHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ADD_ARRAY_UNPACK
func getAddArrayUnpackHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ISSET_ISEMPTY_PROP_OBJ
func getIssetIsemptyPropObjHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_HANDLE_EXCEPTION
func getHandleExceptionHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_USER_OPCODE
func getUserOpcodeHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ASSERT_CHECK
func getAssertCheckHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_JMP_SET
func getJmpSetHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_UNSET_CV
func getUnsetCvHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ISSET_ISEMPTY_CV
func getIssetIsemptyCvHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_LIST_W
func getFetchListWHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SEPARATE
func getSeparateHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_CLASS_NAME
func getFetchClassNameHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_CALL_TRAMPOLINE
func getCallTrampolineHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_DISCARD_EXCEPTION
func getDiscardExceptionHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_YIELD
func getYieldHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_GENERATOR_RETURN
func getGeneratorReturnHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FAST_CALL
func getFastCallHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FAST_RET
func getFastRetHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_RECV_VARIADIC
func getRecvVariadicHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SEND_UNPACK
func getSendUnpackHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_YIELD_FROM
func getYieldFromHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_COPY_TMP
func getCopyTmpHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_BIND_GLOBAL
func getBindGlobalHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_COALESCE
func getCoalesceHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SPACESHIP
func getSpaceshipHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FUNC_NUM_ARGS
func getFuncNumArgsHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FUNC_GET_ARGS
func getFuncGetArgsHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_STATIC_PROP_R
func getFetchStaticPropRHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_STATIC_PROP_W
func getFetchStaticPropWHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_STATIC_PROP_RW
func getFetchStaticPropRwHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_STATIC_PROP_IS
func getFetchStaticPropIsHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_STATIC_PROP_FUNC_ARG
func getFetchStaticPropFuncArgHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_STATIC_PROP_UNSET
func getFetchStaticPropUnsetHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_UNSET_STATIC_PROP
func getUnsetStaticPropHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ISSET_ISEMPTY_STATIC_PROP
func getIssetIsemptyStaticPropHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_CLASS_CONSTANT
func getFetchClassConstantHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_BIND_LEXICAL
func getBindLexicalHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_BIND_STATIC
func getBindStaticHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_FETCH_THIS
func getFetchThisHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SEND_FUNC_ARG
func getSendFuncArgHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ISSET_ISEMPTY_THIS
func getIssetIsemptyThisHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SWITCH_LONG
func getSwitchLongHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_SWITCH_STRING
func getSwitchStringHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_IN_ARRAY
func getInArrayHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_COUNT
func getCountHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_GET_CLASS
func getGetClassHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_GET_CALLED_CLASS
func getGetCalledClassHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_GET_TYPE
func getGetTypeHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}

// ZEND_ARRAY_KEY_EXISTS
func getArrayKeyExistsHandler(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval = executeData.Op1(opline, opMode1)
	var op2 *types.Zval = executeData.Op2(opline, opMode1)
	//todo
}
