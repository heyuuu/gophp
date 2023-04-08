package zend

import "github.com/heyuuu/gophp/zend/types"

func ZEND_INIT_FCALL_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var fname *types.Zval
	var fbc types.IFunction
	var call *ZendExecuteData
	fbc = CACHED_PTR(opline.GetResult().GetNum())
	if fbc == nil {
		fname = (*types.Zval)(opline.Const2())
		fbc = EG__().FunctionTable().Get(fname.GetStr().GetStr())
		if fbc == nil {
			return zend_undefined_function_helper_SPEC(executeData)
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		CACHE_PTR(opline.GetResult().GetNum(), fbc)
	}
	call = _zendVmStackPushCallFrameEx(opline.GetOp1().GetNum(), ZEND_CALL_NESTED_FUNCTION, fbc, opline.GetExtendedValue(), nil)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
