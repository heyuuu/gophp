package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZEND_BOOL_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = opline.Const1()
	if val.IsTrue() {
		opline.Result().SetTrue()
	} else if val.IsSignType() {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		opline.Result().SetFalse()
	} else {
		opline.Result().SetBool(operators.ZvalIsTrue(val))
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_BOOL_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var val *types.Zval
	var free_op1 ZendFreeOp
	val = opline.Op1()
	if val.IsTrue() {
		opline.Result().SetTrue()
	} else if val.IsSignType() {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		opline.Result().SetFalse()
		if orig_val_type == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1(executeData)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	} else {
		opline.Result().SetBool(operators.ZvalIsTrue(val))
		// ZvalPtrDtorNogc(free_op1)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_BOOL_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var val *types.Zval
	val = opline.Op1()
	if val.IsTrue() {
		opline.Result().SetTrue()
	} else if val.IsSignType() {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		opline.Result().SetFalse()
		if orig_val_type == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1(executeData)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	} else {
		opline.Result().SetBool(operators.ZvalIsTrue(val))
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
