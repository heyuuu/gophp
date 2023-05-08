package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_JMPZNZ_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = opline.Const1()
	if val.IsTrue() {
		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, opline.GetExtendedValue())
		return 0
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	if ZvalIsTrue(val) {
		opline = ZEND_OFFSET_TO_OPLINE(opline, opline.GetExtendedValue())
	} else {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	return ZEND_VM_JMP(executeData, opline)
}
func ZEND_JMPZNZ_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = opline.Op1()
	if val.IsTrue() {
		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, opline.GetExtendedValue())
		return 0
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		if val.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
			if EG__().GetException() != nil {
				return 0
			}
		}
		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	if ZvalIsTrue(val) {
		opline = ZEND_OFFSET_TO_OPLINE(opline, opline.GetExtendedValue())
	} else {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	return ZEND_VM_JMP(executeData, opline)
}
func ZEND_JMPZNZ_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = opline.Op1()
	if val.IsTrue() {
		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, opline.GetExtendedValue())
		return 0
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		if val.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
			if EG__().GetException() != nil {
				return 0
			}
		}
		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	if ZvalIsTrue(val) {
		opline = ZEND_OFFSET_TO_OPLINE(opline, opline.GetExtendedValue())
	} else {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	return ZEND_VM_JMP(executeData, opline)
}
