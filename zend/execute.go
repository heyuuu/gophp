package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/types"
)

func ZendVmStackPushCallFrame(callInfo uint32, func_ types.IFunction, numArgs uint32, objectOrCalledScope any) *ZendExecuteData {
	var runtimeCacheSize uint32 = numArgs
	if ZEND_USER_CODE(func_.GetType()) {
		runtimeCacheSize += func_.GetOpArray().GetLastVar() + func_.GetOpArray().GetT() - b.Min(func_.GetOpArray().GetNumArgs(), numArgs)
	}
	return ZendVmStackPushCallFrameEx(runtimeCacheSize, callInfo, func_, numArgs, objectOrCalledScope)
}

func ZendVmCalcUsedStack(numArgs uint32, func_ types.IFunction) uint32 {
	var usedStack uint32 = ZEND_CALL_FRAME_SLOT + numArgs
	if ZEND_USER_CODE(func_.GetType()) {
		usedStack += func_.GetOpArray().GetLastVar() + func_.GetOpArray().GetT() - b.Min(func_.GetOpArray().GetNumArgs(), numArgs)
	}
	return usedStack * b.SizeOf("zval")
}

func ZendVmStackPushCallFrameEx(usedStack uint32, callInfo uint32, func_ types.IFunction, numArgs uint32, objectOrCalledScope any) *ZendExecuteData {
	runtimeCacheSize := usedStack/b.SizeOf("zval") - ZEND_CALL_FRAME_SLOT
	//callInfo |= ZEND_CALL_ALLOCATED
	return ZendVmStackPushCallFrameExEx(runtimeCacheSize, callInfo, func_, numArgs, objectOrCalledScope)
}

func ZendVmStackPushCallFrameExEx(runtimeCacheSize uint32, callInfo uint32, func_ types.IFunction, numArgs uint32, objectOrCalledScope any) *ZendExecuteData {
	var call *ZendExecuteData = NewExecuteData(callInfo, func_, numArgs, objectOrCalledScope, runtimeCacheSize)
	EG__().VmStack().Push(call)
	return call
}
