package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_DO_UCALL_SPEC_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	ret = nil

	call.SetPrevExecuteData(executeData)
	executeData = call
	IInitFuncExecuteData(fbc.GetOpArray(), ret, 0, executeData)
	return 1
}
func ZEND_DO_UCALL_SPEC_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	ret = nil
	ret = opline.Result()
	call.SetPrevExecuteData(executeData)
	executeData = call
	IInitFuncExecuteData(fbc.GetOpArray(), ret, 0, executeData)
	return 1
}
