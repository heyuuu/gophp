package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

func ZEND_INIT_NS_FCALL_BY_NAME_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var func_name *types2.Zval
	var fbc types2.IFunction
	var call *ZendExecuteData
	fbc = CACHED_PTR(opline.GetResult().GetNum())
	if fbc == nil {
		func_name = (*types2.Zval)(opline.Const2())
		var func_name_1 *types2.Zval = func_name + 1
		fbc = EG__().FunctionTable().Get(func_name_1.StringVal())
		if fbc == nil {
			var func_name_2 *types2.Zval = func_name + 2
			fbc = EG__().FunctionTable().Get(func_name_2.StringVal())
			if fbc == nil {
				return zend_undefined_function_helper_SPEC(executeData)
			}
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		CACHE_PTR(opline.GetResult().GetNum(), fbc)
	}
	call = ZendVmStackPushCallFrame(ZEND_CALL_NESTED_FUNCTION, fbc, opline.GetExtendedValue(), nil)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
