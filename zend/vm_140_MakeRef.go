package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

func ZEND_MAKE_REF_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types2.Zval = opline.Op1()
	{
		if op1.IsUndef() {
			op1.SetNewEmptyRef()
			op1.SetRefcount(2)
			types2.Z_REFVAL_P(op1).SetNull()
			opline.Result().SetReference(op1.Reference())
		} else {
			if op1.IsReference() {
				// 				op1.AddRefcount()
			} else {
				types2.ZVAL_MAKE_REF_EX(op1, 2)
			}
			opline.Result().SetReference(op1.Reference())
		}
	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_MAKE_REF_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types2.Zval = opline.Op1()
	{
		if op1.IsUndef() {
			op1.SetNewEmptyRef()
			op1.SetRefcount(2)
			types2.Z_REFVAL_P(op1).SetNull()
			opline.Result().SetReference(op1.Reference())
		} else {
			if op1.IsReference() {
				// 				op1.AddRefcount()
			} else {
				types2.ZVAL_MAKE_REF_EX(op1, 2)
			}
			opline.Result().SetReference(op1.Reference())
		}
	}

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
