package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZEND_JMPZ_EX_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	var ret int
	val = opline.Const1()
	if val.IsTrue() {
		opline.Result().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		opline.Result().SetFalse()
		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ret = operators.IZendIsTrue(val)
	if ret != 0 {
		opline.Result().SetTrue()
		opline++
	} else {
		opline.Result().SetFalse()
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	return ZEND_VM_JMP(executeData, opline)
}
func ZEND_JMPZ_EX_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var val *types.Zval
	var ret int
	val = opline.Op1()
	if val.IsTrue() {
		opline.Result().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		opline.Result().SetFalse()
		if val.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
			if EG__().GetException() != nil {
				return 0
			}
		}
		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ret = operators.IZendIsTrue(val)
	// ZvalPtrDtorNogc(free_op1)
	if ret != 0 {
		opline.Result().SetTrue()
		opline++
	} else {
		opline.Result().SetFalse()
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	return ZEND_VM_JMP(executeData, opline)
}
func ZEND_JMPZ_EX_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	var ret int
	val = opline.Op1()
	if val.IsTrue() {
		opline.Result().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		opline.Result().SetFalse()
		if val.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
			if EG__().GetException() != nil {
				return 0
			}
		}
		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ret = operators.IZendIsTrue(val)
	if ret != 0 {
		opline.Result().SetTrue()
		opline++
	} else {
		opline.Result().SetFalse()
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	return ZEND_VM_JMP(executeData, opline)
}
