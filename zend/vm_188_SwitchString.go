package zend

func ZEND_SWITCH_STRING_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op *types.Zval
	var jump_zv *types.Zval
	var jumptable *types.Array
	op = RT_CONSTANT(opline, opline.GetOp1())
	jumptable = RT_CONSTANT(opline, opline.GetOp2()).GetArr()
	if op.GetType() != types.IS_STRING {
		{

			/* Wrong type, fall back to ZEND_CASE chain */

			return ZEND_VM_NEXT_OPCODE(executeData, opline)

			/* Wrong type, fall back to ZEND_CASE chain */

		}

		/* Wrong type, fall back to ZEND_CASE chain */

		/* Wrong type, fall back to ZEND_CASE chain */

	}
	jump_zv = jumptable.KeyFind(op.GetStr().GetStr())
	if jump_zv != nil {
		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, jump_zv.GetLval())
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
	op = opline.GetOp1Zval()
	jumptable = RT_CONSTANT(opline, opline.GetOp2()).GetArr()
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
	jump_zv = jumptable.KeyFind(op.GetStr().GetStr())
	if jump_zv != nil {
		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, jump_zv.GetLval())
		return 0
	} else {

		/* default */

		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, opline.GetExtendedValue())
		return 0
	}
}
