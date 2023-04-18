package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func ZendVmStackPushCallFrame(callInfo uint32, func_ types.IFunction, numArgs uint32, objectOrCalledScope any) *ZendExecuteData {
	var runtimeCacheSize uint32 = ZendVmCalcUsedStack(numArgs, func_)
	return ZendVmStackPushCallFrameEx(runtimeCacheSize, callInfo, func_, numArgs, objectOrCalledScope)
}

func ZendVmCalcUsedStack(numArgs uint32, func_ types.IFunction) uint32 {
	var usedStack = numArgs
	if ZEND_USER_CODE(func_.GetType()) {
		usedStack += func_.GetOpArray().GetLastVar() + func_.GetOpArray().GetT() - b.Min(func_.GetOpArray().GetNumArgs(), numArgs)
	}
	return usedStack
}

func ZendVmStackPushCallFrameEx(runtimeCacheSize uint32, callInfo uint32, func_ types.IFunction, numArgs uint32, objectOrCalledScope any) *ZendExecuteData {
	var call *ZendExecuteData = NewExecuteData(callInfo, func_, numArgs, objectOrCalledScope, runtimeCacheSize)
	EG__().VmStack().Push(call)
	return call
}
