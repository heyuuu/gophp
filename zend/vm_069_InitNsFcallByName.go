package zend

func ZEND_INIT_NS_FCALL_BY_NAME_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var func_name *types.Zval
	var func_ *types.Zval
	var fbc types.IFunction
	var call *ZendExecuteData
	fbc = CACHED_PTR(opline.GetResult().GetNum())
	if fbc == nil {
		func_name = (*types.Zval)(RT_CONSTANT(opline, opline.GetOp2()))
		func_ = EG__().GetFunctionTable().KeyFind((func_name + 1).GetStr().GetStr())
		if func_ == nil {
			func_ = EG__().GetFunctionTable().KeyFind((func_name + 2).GetStr().GetStr())
			if func_ == nil {
				return zend_undefined_function_helper_SPEC(executeData)
			}
		}
		fbc = func_.GetFunc()
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		CACHE_PTR(opline.GetResult().GetNum(), fbc)
	}
	call = _zendVmStackPushCallFrame(ZEND_CALL_NESTED_FUNCTION, fbc, opline.GetExtendedValue(), nil)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
