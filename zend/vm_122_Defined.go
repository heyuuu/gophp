package zend

func ZEND_DEFINED_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var c *ZendConstant
	c = CACHED_PTR(opline.GetExtendedValue())
	if c != nil {
		if IS_SPECIAL_CACHE_VAL(c) == 0 {
		defined_true:
			ZEND_VM_SMART_BRANCH_TRUE()
			opline.GetResultZval().SetTrue()
			return ZEND_VM_NEXT_OPCODE(executeData, opline)
		} else if EG__().GetZendConstants().Len() == DECODE_SPECIAL_CACHE_NUM(c) {
		defined_false:
			ZEND_VM_SMART_BRANCH_FALSE()
			opline.GetResultZval().SetFalse()
			return ZEND_VM_NEXT_OPCODE(executeData, opline)
		}
	}
	if ZendQuickCheckConstant(RT_CONSTANT(opline, opline.GetOp1()), opline, executeData) != types.SUCCESS {
		CACHE_PTR(opline.GetExtendedValue(), ENCODE_SPECIAL_CACHE_NUM(EG__().GetZendConstants().Len()))
		goto defined_false
	} else {
		goto defined_true
	}
}
