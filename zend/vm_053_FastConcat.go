package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

func fastConcatHelper(executeData *ZendExecuteData, op1, op2 *types.Zval) int {
	var op1Str, op2Str string
	var checkExcpetion bool = false

	if op1.IsString() {
		op1Str = op1.StringVal()
	} else {
		checkExcpetion = true
		if op1.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
		}
		op1Str = operators.ZvalGetStrVal(op1)
	}

	if op2.IsString() {
		op2Str = op2.StringVal()
	} else {
		checkExcpetion = true
		if op2.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
		}
		op2Str = operators.ZvalGetStrVal(op2)
	}

	str := op1Str + op2Str

	//
	var opline *types.ZendOp = executeData.GetOpline()
	opline.Result().SetString(str)
	if checkExcpetion {
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	} else {
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}

func ZEND_FAST_CONCAT_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = executeData.GetOp1(opline)
	op2 = executeData.GetOp2(opline)

	return fastConcatHelper(executeData, op1, op2)
}
func ZEND_FAST_CONCAT_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = executeData.GetOp1(opline)
	op2 = opline.Op2()
	return fastConcatHelper(executeData, op1, op2)
}
func ZEND_FAST_CONCAT_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = executeData.GetOp1(opline)
	op2 = opline.Op2()
	return fastConcatHelper(executeData, op1, op2)
}
func ZEND_FAST_CONCAT_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1()
	op2 = executeData.GetOp2(opline)

	return fastConcatHelper(executeData, op1, op2)
}
func ZEND_FAST_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1()
	op2 = opline.Op2()

	return fastConcatHelper(executeData, op1, op2)
}
func ZEND_FAST_CONCAT_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1()
	op2 = opline.Op2()

	return fastConcatHelper(executeData, op1, op2)
}
func ZEND_FAST_CONCAT_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1()
	op2 = executeData.GetOp2(opline)

	return fastConcatHelper(executeData, op1, op2)
}
func ZEND_FAST_CONCAT_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1()
	op2 = opline.Op2()
	return fastConcatHelper(executeData, op1, op2)
}
func ZEND_FAST_CONCAT_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1()
	op2 = opline.Op2()

	return fastConcatHelper(executeData, op1, op2)
}
