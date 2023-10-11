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
	if opArray, ok := func_.(*types.ZendOpArray); ok {
		usedStack += uint32(opArray.GetLastVar()) + opArray.GetT() - b.Min(opArray.GetNumArgs(), numArgs)
	}
	return usedStack
}

func ZendVmStackPushCallFrameEx(runtimeCacheSize uint32, callInfo uint32, func_ types.IFunction, numArgs uint32, objectOrCalledScope any) *ZendExecuteData {
	var call = NewExecuteData(callInfo, func_, numArgs, objectOrCalledScope, runtimeCacheSize)
	EG__().VmStackPush(call)
	return call
}
