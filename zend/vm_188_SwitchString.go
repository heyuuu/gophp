package zend

import "github.com/heyuuu/gophp/zend/types"

func ZEND_SWITCH_STRING_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op *types.Zval
	var jump_zv *types.Zval
	var jumptable *types.Array
	op = opline.Const1()
	jumptable = opline.Const2().Array()
	if op.GetType() != types.IS_STRING {
		{

			/* Wrong type, fall back to ZEND_CASE chain */

			return ZEND_VM_NEXT_OPCODE(executeData, opline)

			/* Wrong type, fall back to ZEND_CASE chain */

		}

		/* Wrong type, fall back to ZEND_CASE chain */

		/* Wrong type, fall back to ZEND_CASE chain */

	}
	jump_zv = jumptable.KeyFind(op.StringVal())
	if jump_zv != nil {
		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, jump_zv.Long()())
		return 0
	} else {

		/* default */

		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, opline.GetExtendedValue())
		return 0
	}
}
func ZEND_SWITCH_STRING_SPEC_TMPVARCV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op *types.Zval
	var jump_zv *types.Zval
	var jumptable *types.Array
	op = opline.Op1()
	jumptable = opline.Const2().Array()
	if op.GetType() != types.IS_STRING {

		/* Wrong type, fall back to ZEND_CASE chain */

		/* Wrong type, fall back to ZEND_CASE chain */

		{
			op = types.ZVAL_DEREF(op)
			if op.GetType() != types.IS_STRING {

				/* Wrong type, fall back to ZEND_CASE chain */

				return ZEND_VM_NEXT_OPCODE(executeData, opline)

				/* Wrong type, fall back to ZEND_CASE chain */

			}
		}
	}
	jump_zv = jumptable.KeyFind(op.StringVal())
	if jump_zv != nil {
		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, jump_zv.Long()())
		return 0
	} else {

		/* default */

		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, opline.GetExtendedValue())
		return 0
	}
}
