package zend

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifZendVersion
var DefZifZendVersion = def.DefFunc("zend_version", 0, 0, def.DefFuncOpts{Handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifZendVersion()
	returnValue.SetStringVal(ret)
}})

// generate by ZifGcMemCaches
var DefZifGcMemCaches = def.DefFunc("gc_mem_caches", 0, 0, def.DefFuncOpts{Handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGcMemCaches()
	returnValue.SetLong(ret)
}})

// generate by ZifGcCollectCycles
var DefZifGcCollectCycles = def.DefFunc("gc_collect_cycles", 0, 0, def.DefFuncOpts{Handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGcCollectCycles()
	returnValue.SetLong(ret)
}})

// generate by ZifGcEnabled
var DefZifGcEnabled = def.DefFunc("gc_enabled", 0, 0, def.DefFuncOpts{Handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGcEnabled()
	returnValue.SetBool(ret)
}})

// generate by ZifGcEnable
var DefZifGcEnable = def.DefFunc("gc_enable", 0, 0, def.DefFuncOpts{Handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGcEnable()
}})

// generate by ZifGcDisable
var DefZifGcDisable = def.DefFunc("gc_disable", 0, 0, def.DefFuncOpts{Handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGcDisable()
}})

// generate by ZifGcStatus
var DefZifGcStatus = def.DefFunc("gc_status", 0, 0, def.DefFuncOpts{Handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGcStatus(returnValue)
}})

// generate by ZifFuncNumArgs
var DefZifFuncNumArgs = def.DefFunc("func_num_args", 0, 0, def.DefFuncOpts{Handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifFuncNumArgs(executeData)
	returnValue.SetLong(ret)
}})

// generate by ZifFuncGetArg
var DefZifFuncGetArg = def.DefFunc("func_get_arg", 1, 1, def.DefFuncOpts{ArgInfos: []def.ArgInfo{{name: "arg_num"}}, Handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, zpp.FlagOldMode)
	arg_num := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ZifFuncGetArg(executeData, returnValue, arg_num)
}})

// generate by ZifFuncGetArgs
var DefZifFuncGetArgs = def.DefFunc("func_get_args", 0, 0, def.DefFuncOpts{Handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret, ok := ZifFuncGetArgs(executeData)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
}})

// generate by ZifStrlen
var DefZifStrlen = def.DefFunc("strlen", 1, 1, def.DefFuncOpts{ArgInfos: []def.ArgInfo{{name: "str"}}, Handler: func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStr()
	if fp.HasError() {
		return
	}
	ret := ZifStrlen(str)
	returnValue.SetLong(ret)
}})
