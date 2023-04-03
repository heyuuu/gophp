package zend

func ZEND_DECLARE_LAMBDA_FUNCTION_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var func_ types.IFunction
	var zfunc *types.Zval
	var object *types.Zval
	var called_scope *types.ClassEntry
	func_ = CACHED_PTR(opline.GetExtendedValue())
	if func_ == nil {
		zfunc = EG__().GetFunctionTable().KeyFind(RT_CONSTANT(opline, opline.GetOp1()).GetStr().GetStr())
		b.Assert(zfunc != nil)
		func_ = zfunc.GetFunc()
		b.Assert(func_.GetType() == ZEND_USER_FUNCTION)
		CACHE_PTR(opline.GetExtendedValue(), func_)
	}
	if executeData.GetThis().IsObject() {
		called_scope = types.Z_OBJCE(executeData.GetThis())
		if func_.IsStatic() || (executeData.GetFunc().common.fn_flags&AccStatic) != 0 {
			object = nil
		} else {
			object = &(executeData.GetThis())
		}
	} else {
		called_scope = executeData.GetThis().GetCe()
		object = nil
	}
	ZendCreateClosure(opline.GetResultZval(), func_, executeData.GetFunc().op_array.scope, called_scope, object)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
