package zend

func ZEND_FETCH_CLASS_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var class_name *types.Zval
	var opline *ZendOp = executeData.GetOpline()

	{
		var ce *types.ClassEntry = CACHED_PTR(opline.GetExtendedValue())
		if ce == nil {
			class_name = RT_CONSTANT(opline, opline.GetOp2())
			ce = ZendFetchClassByName(class_name.GetStr(), (class_name + 1).GetStr(), opline.GetOp1().GetNum())
			CACHE_PTR(opline.GetExtendedValue(), ce)
		}
		opline.Result().SetCe(ce)
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_CLASS_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var free_op2 ZendFreeOp
	var class_name *types.Zval
	var opline *ZendOp = executeData.GetOpline()

	{
		class_name = opline.Op2Ptr(&free_op2)
	try_class_name:
		if class_name.IsObject() {
			opline.Result().SetCe(types.Z_OBJCE_P(class_name))
		} else if class_name.IsString() {
			opline.Result().SetCe(ZendFetchClass(class_name.GetStr(), opline.GetOp1().GetNum()))
		} else if class_name.IsReference() {
			class_name = types.Z_REFVAL_P(class_name)
			goto try_class_name
		} else {
			if class_name.IsUndef() {
				ZVAL_UNDEFINED_OP2()
				if EG__().GetException() != nil {
					return 0
				}
			}
			faults.ThrowError(nil, "Class name must be a valid object or a string")
		}
	}
	ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_CLASS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var class_name *types.Zval
	var opline *ZendOp = executeData.GetOpline()
	{
		opline.Result().SetCe(ZendFetchClass(nil, opline.GetOp1().GetNum()))
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_CLASS_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var class_name *types.Zval
	var opline *ZendOp = executeData.GetOpline()

	{
		class_name = opline.Op2()
	try_class_name:
		if class_name.IsObject() {
			opline.Result().SetCe(types.Z_OBJCE_P(class_name))
		} else if class_name.IsString() {
			opline.Result().SetCe(ZendFetchClass(class_name.GetStr(), opline.GetOp1().GetNum()))
		} else if class_name.IsReference() {
			class_name = types.Z_REFVAL_P(class_name)
			goto try_class_name
		} else {
			if class_name.IsUndef() {
				ZVAL_UNDEFINED_OP2()
				if EG__().GetException() != nil {
					return 0
				}
			}
			faults.ThrowError(nil, "Class name must be a valid object or a string")
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
