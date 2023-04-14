package zend

func ZEND_TYPE_CHECK_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int = 0
	value = opline.Const1()
	if (opline.GetExtendedValue() >> uint32(value.GetType()) & 1) != 0 {
	type_check_resource:
		if value.GetType() != types.IS_RESOURCE || nil != ZendRsrcListGetRsrcType(value.GetRes()) {
			result = 1
		}
	} else {
	}

	{
		ZEND_VM_SMART_BRANCH(result, 0)
		types.ZVAL_BOOL(opline.Result(), result != 0)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func ZEND_TYPE_CHECK_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int = 0
	var free_op1 ZendFreeOp
	value = opline.Op1Ptr(&free_op1)
	if (opline.GetExtendedValue() >> uint32(value.GetType()) & 1) != 0 {
	type_check_resource:
		if value.GetType() != types.IS_RESOURCE || nil != ZendRsrcListGetRsrcType(value.GetRes()) {
			result = 1
		}
	} else if value.IsReference() {
		value = types.Z_REFVAL_P(value)
		if (opline.GetExtendedValue() >> uint32(value.GetType()) & 1) != 0 {
			goto type_check_resource
		}
	} else if value.IsUndef() {
		result = (1 << types.IS_NULL & opline.GetExtendedValue()) != 0
		ZVAL_UNDEFINED_OP1(executeData)
		if EG__().GetException() != nil {
			opline.Result().SetUndef()
			return 0
		}
	}
	{
		ZvalPtrDtorNogc(free_op1)
		ZEND_VM_SMART_BRANCH(result, 1)
		types.ZVAL_BOOL(opline.Result(), result != 0)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}

}
func ZEND_TYPE_CHECK_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int = 0
	value = opline.Op1()
	if (opline.GetExtendedValue() >> uint32(value.GetType()) & 1) != 0 {
	type_check_resource:
		if value.GetType() != types.IS_RESOURCE || nil != ZendRsrcListGetRsrcType(value.GetRes()) {
			result = 1
		}
	} else if value.IsReference() {
		value = types.Z_REFVAL_P(value)
		if (opline.GetExtendedValue() >> uint32(value.GetType()) & 1) != 0 {
			goto type_check_resource
		}
	} else if value.IsUndef() {
		result = (1 << types.IS_NULL & opline.GetExtendedValue()) != 0
		ZVAL_UNDEFINED_OP1(executeData)
		if EG__().GetException() != nil {
			opline.Result().SetUndef()
			return 0
		}
	}
	{
		ZEND_VM_SMART_BRANCH(result, 1)
		types.ZVAL_BOOL(opline.Result(), result != 0)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}

}
