package zend

import "sik/zend/types"

// generate by ZifZendVersion
var DefZifZendVersion = DefFunc(DefFuncOpts{name: "zend_version", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	ret := ZifZendVersion()
	returnValue.SetRawString(ret)
}})

// generate by ZifGcMemCaches
var DefZifGcMemCaches = DefFunc(DefFuncOpts{name: "gc_mem_caches", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	ret := ZifGcMemCaches()
	returnValue.SetLong(ret)
}})

// generate by ZifGcCollectCycles
var DefZifGcCollectCycles = DefFunc(DefFuncOpts{name: "gc_collect_cycles", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	ret := ZifGcCollectCycles()
	returnValue.SetLong(ret)
}})

// generate by ZifGcEnabled
var DefZifGcEnabled = DefFunc(DefFuncOpts{name: "gc_enabled", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	ret := ZifGcEnabled()
	returnValue.SetBool(ret)
}})

// generate by ZifGcEnable
var DefZifGcEnable = DefFunc(DefFuncOpts{name: "gc_enable", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	ZifGcEnable()
}})

// generate by ZifGcDisable
var DefZifGcDisable = DefFunc(DefFuncOpts{name: "gc_disable", handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	ZifGcDisable()
}})
