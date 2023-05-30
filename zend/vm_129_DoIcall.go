package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_DO_ICALL_SPEC_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	var retval types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	call.SetPrevExecuteData(executeData)
	EG__().SetCurrentExecuteData(call)
	ret = &retval
	ret.SetNull()
	fbc.GetInternalFunction().GetHandler()(call, ret)
	EG__().SetCurrentExecuteData(executeData)
	ZendVmStackFreeArgs(call)
	ZendVmStackFreeCallFrame(call)
	// IZvalPtrDtor(ret)
	if EG__().GetException() != nil {
		faults.RethrowException(executeData)
		return 0
	}
	OPLINE = opline + 1
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
func ZEND_DO_ICALL_SPEC_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	var retval types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	call.SetPrevExecuteData(executeData)
	EG__().SetCurrentExecuteData(call)
	ret = opline.Result()
	ret.SetNull()
	fbc.GetInternalFunction().GetHandler()(call, ret)
	EG__().SetCurrentExecuteData(executeData)
	ZendVmStackFreeArgs(call)
	ZendVmStackFreeCallFrame(call)

	if EG__().GetException() != nil {
		faults.RethrowException(executeData)
		return 0
	}
	OPLINE = opline + 1
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
