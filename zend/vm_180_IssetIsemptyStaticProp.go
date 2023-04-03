package zend

func ZEND_ISSET_ISEMPTY_STATIC_PROP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int
	result = ZendFetchStaticPropertyAddress(&value, nil, opline.GetExtendedValue() & ^ZEND_ISEMPTY, BP_VAR_IS, 0, opline, executeData)
	if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {
		result = result == types.SUCCESS && value.GetType() > types.IS_NULL && (!(value.IsReference()) || types.Z_REFVAL_P(value).GetType() != types.IS_NULL)
	} else {
		result = result != types.SUCCESS || IZendIsTrue(value) == 0
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.Result(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
