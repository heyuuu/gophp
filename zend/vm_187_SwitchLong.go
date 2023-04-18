package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

func ZEND_SWITCH_LONG_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op *types2.Zval
	var jump_zv *types2.Zval
	var jumptable *types2.Array
	op = opline.Const1()
	jumptable = opline.Const2().Array()
	if op.GetType() != types2.IS_LONG {
		op = types2.ZVAL_DEREF(op)
		if op.GetType() != types2.IS_LONG {

			/* Wrong type, fall back to ZEND_CASE chain */

			return ZEND_VM_NEXT_OPCODE(executeData, opline)

			/* Wrong type, fall back to ZEND_CASE chain */

		}
	}
	jump_zv = jumptable.IndexFind(op.Long()())
	if jump_zv != nil {
		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, jump_zv.Long()())
		return 0
	} else {

		/* default */

		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, opline.GetExtendedValue())
		return 0
	}
}
func ZEND_SWITCH_LONG_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op *types2.Zval
	var jump_zv *types2.Zval
	var jumptable *types2.Array
	op = opline.Op1()
	jumptable = opline.Const2().Array()
	if op.GetType() != types2.IS_LONG {
		op = types2.ZVAL_DEREF(op)
		if op.GetType() != types2.IS_LONG {

			/* Wrong type, fall back to ZEND_CASE chain */

			return ZEND_VM_NEXT_OPCODE(executeData, opline)

			/* Wrong type, fall back to ZEND_CASE chain */

		}
	}
	jump_zv = jumptable.IndexFind(op.Long()())
	if jump_zv != nil {
		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, jump_zv.Long()())
		return 0
	} else {

		/* default */

		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, opline.GetExtendedValue())
		return 0
	}
}
