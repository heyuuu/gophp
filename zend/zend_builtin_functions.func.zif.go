package zend

// generate by ZifZendVersion
var ZifZendVersionDef = DefFunc(DefFuncOpts{name: "zend_version", handler: func(executeData *ZendExecuteData, returnValue *Zval) {
	if !executeData.CheckNumArgs(0, 0, false) {
		return
	}
	ret := ZifZendVersion()
	returnValue.SetRawString(ret)
}})

// generate by ZifGcMemCaches
var ZifGcMemCachesDef = DefFunc(DefFuncOpts{name: "gc_mem_caches", handler: func(executeData *ZendExecuteData, returnValue *Zval) {
	if !executeData.CheckNumArgs(0, 0, false) {
		return
	}
	ret := ZifGcMemCaches()
	returnValue.SetLong(ret)
}})

// generate by ZifGcCollectCycles
var ZifGcCollectCyclesDef = DefFunc(DefFuncOpts{name: "gc_collect_cycles", handler: func(executeData *ZendExecuteData, returnValue *Zval) {
	if !executeData.CheckNumArgs(0, 0, false) {
		return
	}
	ret := ZifGcCollectCycles()
	returnValue.SetLong(ret)
}})

// generate by ZifGcEnabled
var ZifGcEnabledDef = DefFunc(DefFuncOpts{name: "gc_enabled", handler: func(executeData *ZendExecuteData, returnValue *Zval) {
	if !executeData.CheckNumArgs(0, 0, false) {
		return
	}
	ret := ZifGcEnabled()
	returnValue.SetBool(ret)
}})
