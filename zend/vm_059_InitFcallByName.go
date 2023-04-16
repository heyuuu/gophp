package zend

import "github.com/heyuuu/gophp/zend/types"

func ZEND_INIT_FCALL_BY_NAME_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var fbc types.IFunction
	var function_name *types.Zval
	var call *ZendExecuteData
	fbc = CACHED_PTR(opline.GetResult().GetNum())
	if fbc == nil {
		function_name = (*types.Zval)(opline.Const2())
		var function_name_1 *types.Zval = function_name + 1
		fbc = EG__().FunctionTable().Get(function_name_1.StringVal())
		if fbc == nil {
			return zend_undefined_function_helper_SPEC(executeData)
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
