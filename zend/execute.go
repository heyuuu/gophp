package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/types"
)

func ZendVmStackPushCallFrame(callInfo uint32, func_ types.IFunction, numArgs uint32, objectOrCalledScope any) *ZendExecuteData {
	var usedStack uint32 = ZendVmCalcUsedStack(numArgs, func_)
	return ZendVmStackPushCallFrameEx(usedStack, callInfo, func_, numArgs, objectOrCalledScope)
}

func ZendVmCalcUsedStack(numArgs uint32, func_ types.IFunction) uint32 {
	var usedStack uint32 = ZEND_CALL_FRAME_SLOT + numArgs
	if ZEND_USER_CODE(func_.GetType()) {
		usedStack += func_.GetOpArray().GetLastVar() + func_.GetOpArray().GetT() - b.Min(func_.GetOpArray().GetNumArgs(), numArgs)
	}
	return usedStack * b.SizeOf("zval")
}

func ZendVmStackPushCallFrameEx(usedStack uint32, callInfo uint32, func_ types.IFunction, numArgs uint32, objectOrCalledScope any) *ZendExecuteData {
	var call *ZendExecuteData = (*ZendExecuteData)(EG__().GetVmStackTop())
	if usedStack > size_t((*byte)(EG__().GetVmStackEnd())-(*byte)(call)) {
		call = (*ZendExecuteData)(ZendVmStackExtend(usedStack))
		callInfo |= ZEND_CALL_ALLOCATED
	} else {
		EG__().SetVmStackTop((*types.Zval)((*byte)(call + usedStack)))
	}

	ZendVmInitCallFrame(call, callInfo, func_, numArgs, objectOrCalledScope)
	return call
}

func ZendVmInitCallFrame(call *ZendExecuteData, call_info uint32, func_ types.IFunction, num_args uint32, objectOrCalledScope any) {
	call.Init(call_info, func_, num_args, objectOrCalledScope)
}
