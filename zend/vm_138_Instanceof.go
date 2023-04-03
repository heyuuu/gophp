package zend

func ZEND_INSTANCEOF_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr *types.Zval
	var result types.ZendBool
	expr = opline.getZvalPtrVar1(&free_op1)
try_instanceof:
	if expr.IsObject() {
		var ce *types.ClassEntry
		{
			ce = CACHED_PTR(opline.GetExtendedValue())
			if ce == nil {
				ce = ZendFetchClassByName(RT_CONSTANT(opline, opline.GetOp2()).GetStr(), (RT_CONSTANT(opline, opline.GetOp2()) + 1).GetStr(), ZEND_FETCH_CLASS_NO_AUTOLOAD)
				if ce != nil {
					CACHE_PTR(opline.GetExtendedValue(), ce)
				}
			}
		}

		result = ce != nil && InstanceofFunction(types.Z_OBJCE_P(expr), ce) != 0
	} else if expr.IsReference() {
		expr = types.Z_REFVAL_P(expr)
		goto try_instanceof
	} else {
		if expr.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		result = 0
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INSTANCEOF_SPEC_TMPVAR_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr *types.Zval
	var result types.ZendBool
	expr = opline.getZvalPtrVar1(&free_op1)
try_instanceof:
	if expr.IsObject() {
		var ce *types.ClassEntry

		{
			ce = opline.GetOp2Zval().GetCe()
		}
		result = ce != nil && InstanceofFunction(types.Z_OBJCE_P(expr), ce) != 0
	} else if expr.IsReference() {
		expr = types.Z_REFVAL_P(expr)
		goto try_instanceof
	} else {
		if expr.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		result = 0
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INSTANCEOF_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr *types.Zval
	var result types.ZendBool
	expr = opline.getZvalPtrVar1(&free_op1)
try_instanceof:
	if expr.IsObject() {
		var ce *types.ClassEntry

		{
			ce = ZendFetchClass(nil, opline.GetOp2().GetNum())
			if ce == nil {
				b.Assert(EG__().GetException() != nil)
				ZvalPtrDtorNogc(free_op1)
				opline.GetResultZval().SetUndef()
				return 0
			}
		}

		result = ce != nil && InstanceofFunction(types.Z_OBJCE_P(expr), ce) != 0
	} else if expr.IsReference() {
		expr = types.Z_REFVAL_P(expr)
		goto try_instanceof
	} else {
		if expr.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		result = 0
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INSTANCEOF_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr *types.Zval
	var result types.ZendBool
	expr = opline.GetOp1Zval()
try_instanceof:
	if expr.IsObject() {
		var ce *types.ClassEntry
		{
			ce = CACHED_PTR(opline.GetExtendedValue())
			if ce == nil {
				ce = ZendFetchClassByName(RT_CONSTANT(opline, opline.GetOp2()).GetStr(), (RT_CONSTANT(opline, opline.GetOp2()) + 1).GetStr(), ZEND_FETCH_CLASS_NO_AUTOLOAD)
				if ce != nil {
					CACHE_PTR(opline.GetExtendedValue(), ce)
				}
			}
		}

		result = ce != nil && InstanceofFunction(types.Z_OBJCE_P(expr), ce) != 0
	} else if expr.IsReference() {
		expr = types.Z_REFVAL_P(expr)
		goto try_instanceof
	} else {
		if expr.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		result = 0
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INSTANCEOF_SPEC_CV_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr *types.Zval
	var result types.ZendBool
	expr = opline.GetOp1Zval()
try_instanceof:
	if expr.IsObject() {
		var ce *types.ClassEntry

		{
			ce = opline.GetOp2Zval().GetCe()
		}
		result = ce != nil && InstanceofFunction(types.Z_OBJCE_P(expr), ce) != 0
	} else if expr.IsReference() {
		expr = types.Z_REFVAL_P(expr)
		goto try_instanceof
	} else {
		if expr.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		result = 0
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INSTANCEOF_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr *types.Zval
	var result types.ZendBool
	expr = opline.GetOp1Zval()
try_instanceof:
	if expr.IsObject() {
		var ce *types.ClassEntry

		{
			ce = ZendFetchClass(nil, opline.GetOp2().GetNum())
			if ce == nil {
				b.Assert(EG__().GetException() != nil)
				opline.GetResultZval().SetUndef()
				return 0
			}
		}

		result = ce != nil && InstanceofFunction(types.Z_OBJCE_P(expr), ce) != 0
	} else if expr.IsReference() {
		expr = types.Z_REFVAL_P(expr)
		goto try_instanceof
	} else {
		if expr.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		result = 0
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
