package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

func caseHandlerHelper(op1 *types.Zval, op2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	if op1.IsLong() {
		if op2.IsLong() {
			if op1.Long() == op2.Long() {
				goto caseTrue
			} else {
				goto caseFalse
			}
		} else if op2.IsDouble() {
			d1 := float64(op1.Long())
			d2 := op2.Double()
			if d1 == d2 {
				goto caseTrue
			} else {
				goto caseFalse
			}
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 := op1.Double()
			d2 := op2.Double()
			if d1 == d2 {
				goto caseTrue
			} else {
				goto caseFalse
			}
		} else if op2.IsLong() {
			d1 := op1.Double()
			d2 := float64(op2.Long())
			if d1 == d2 {
				goto caseTrue
			} else {
				goto caseFalse
			}
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result = operators.ZendFastEqualStringsEx(op1.StringVal(), op2.StringVal())
			if result {
				goto caseTrue
			} else {
				goto caseFalse
			}
		}
	}
	return zend_case_helper_SPEC(op1, op2, executeData)
caseTrue:
	ZEND_VM_SMART_BRANCH_TRUE()
	opline.Result().SetTrue()
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
caseFalse:
	ZEND_VM_SMART_BRANCH_FALSE()
	opline.Result().SetFalse()
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}

func ZEND_CASE_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1()
	op2 = executeData.GetOp2(opline)

	return caseHandlerHelper(op1, op2, executeData)
}

func ZEND_CASE_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1()
	op2 = opline.Op2()

	return caseHandlerHelper(op1, op2, executeData)
}
func ZEND_CASE_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = opline.Op1()
	op2 = opline.Op2()

	return caseHandlerHelper(op1, op2, executeData)
}
