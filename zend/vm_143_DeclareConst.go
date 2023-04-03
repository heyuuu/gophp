package zend

func ZEND_DECLARE_CONST_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var name *types.Zval
	var val *types.Zval
	var c ZendConstant
	name = opline.Const1()
	val = opline.Const2()
	types.ZVAL_COPY(c.Value(), val)
	if c.Value().IsConstant() {
		if ZvalUpdateConstantEx(c.Value(), executeData.GetFunc().op_array.scope) != types.SUCCESS {
			ZvalPtrDtorNogc(c.Value())
			return 0
		}
	}

	/* non persistent, case sensitive */

	ZEND_CONSTANT_SET_FLAGS(&c, CONST_CS, PHP_USER_CONSTANT)
	c.SetName(name.GetStr().Copy())
	if ZendRegisterConstant(&c) == types.FAILURE {
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
