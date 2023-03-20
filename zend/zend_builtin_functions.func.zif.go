package zend

import (
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifZendVersion
var DefZifZendVersion = DefFunc(DefFuncOpts{name: "zend_version", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifZendVersion()
	returnValue.SetRawString(ret)
}})

// generate by ZifGcMemCaches
var DefZifGcMemCaches = DefFunc(DefFuncOpts{name: "gc_mem_caches", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGcMemCaches()
	returnValue.SetLong(ret)
}})

// generate by ZifGcCollectCycles
var DefZifGcCollectCycles = DefFunc(DefFuncOpts{name: "gc_collect_cycles", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGcCollectCycles()
	returnValue.SetLong(ret)
}})

// generate by ZifGcEnabled
var DefZifGcEnabled = DefFunc(DefFuncOpts{name: "gc_enabled", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGcEnabled()
	returnValue.SetBool(ret)
}})

// generate by ZifGcEnable
var DefZifGcEnable = DefFunc(DefFuncOpts{name: "gc_enable", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGcEnable()
}})

// generate by ZifGcDisable
var DefZifGcDisable = DefFunc(DefFuncOpts{name: "gc_disable", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGcDisable()
}})

// generate by ZifGcStatus
var DefZifGcStatus = DefFunc(DefFuncOpts{name: "gc_status", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGcStatus(returnValue)
}})

// generate by ZifFuncNumArgs
var DefZifFuncNumArgs = DefFunc(DefFuncOpts{name: "func_num_args", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifFuncNumArgs(executeData)
	returnValue.SetLong(ret)
}})

// generate by ZifFuncGetArg
var DefZifFuncGetArg = DefFunc(DefFuncOpts{name: "func_get_arg", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifFuncGetArg(executeData, returnValue)
}})
