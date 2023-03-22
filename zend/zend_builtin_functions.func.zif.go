package zend

import (
	"sik/zend/zpp"
	"sik/zend/def"
	"sik/zend/types"
)

// generate by ZifZendVersion
var DefZifZendVersion = def.DefFunc("zend_version", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifZendVersion()
	returnValue.SetStringVal(ret)
})

// generate by ZifGcMemCaches
var DefZifGcMemCaches = def.DefFunc("gc_mem_caches", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGcMemCaches()
	returnValue.SetLong(ret)
})

// generate by ZifGcCollectCycles
var DefZifGcCollectCycles = def.DefFunc("gc_collect_cycles", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGcCollectCycles()
	returnValue.SetLong(ret)
})

// generate by ZifGcEnabled
var DefZifGcEnabled = def.DefFunc("gc_enabled", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGcEnabled()
	returnValue.SetBool(ret)
})

// generate by ZifGcEnable
var DefZifGcEnable = def.DefFunc("gc_enable", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGcEnable()
})

// generate by ZifGcDisable
var DefZifGcDisable = def.DefFunc("gc_disable", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGcDisable()
})

// generate by ZifGcStatus
var DefZifGcStatus = def.DefFunc("gc_status", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGcStatus(returnValue)
})

// generate by ZifFuncNumArgs
var DefZifFuncNumArgs = def.DefFunc("func_num_args", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifFuncNumArgs(executeData)
	returnValue.SetLong(ret)
})

// generate by ZifFuncGetArg
var DefZifFuncGetArg = def.DefFunc("func_get_arg", 1, 1, []def.ArgInfo{{name: "arg_num"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, zpp.FlagOldMode)
	arg_num := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ZifFuncGetArg(executeData, returnValue, arg_num)
})

// generate by ZifFuncGetArgs
var DefZifFuncGetArgs = def.DefFunc("func_get_args", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret, ok := ZifFuncGetArgs(executeData)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrlen
var DefZifStrlen = def.DefFunc("strlen", 1, 1, []def.ArgInfo{{name: "str"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStr()
	if fp.HasError() {
		return
	}
	ret := ZifStrlen(str)
	returnValue.SetLong(ret)
})

// generate by ZifStrcmp
var DefZifStrcmp = def.DefFunc("strcmp", 2, 2, []def.ArgInfo{{name: "str1"}, {name: "str2"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	str1 := fp.ParseStringVal()
	str2 := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStrcmp(str1, str2)
	returnValue.SetLong(ret)
})

// generate by ZifStrncmp
var DefZifStrncmp = def.DefFunc("strncmp", 3, 3, []def.ArgInfo{{name: "str1"}, {name: "str2"}, {name: "len_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	str1 := fp.ParseStringVal()
	str2 := fp.ParseStringVal()
	len_ := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrncmp(str1, str2, len_)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrcasecmp
var DefZifStrcasecmp = def.DefFunc("strcasecmp", 2, 2, []def.ArgInfo{{name: "str1"}, {name: "str2"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	str1 := fp.ParseStringVal()
	str2 := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStrcasecmp(str1, str2)
	returnValue.SetLong(ret)
})

// generate by ZifStrncasecmp
var DefZifStrncasecmp = def.DefFunc("strncasecmp", 3, 3, []def.ArgInfo{{name: "str1"}, {name: "str2"}, {name: "len_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	str1 := fp.ParseStringVal()
	str2 := fp.ParseStringVal()
	len_ := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrncasecmp(str1, str2, len_)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifErrorReporting
var DefZifErrorReporting = def.DefFunc("error_reporting", 0, 1, []def.ArgInfo{{name: "new_error_level"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	new_error_level := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifErrorReporting(returnValue, nil, new_error_level)
})

// generate by ZifDefine
var DefZifDefine = def.DefFunc("define", 2, 3, []def.ArgInfo{{name: "constant_name"}, {name: "value"}, {name: "case_insensitive"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	constant_name := fp.ParseStringVal()
	value := fp.ParseZval()
	fp.StartOptional()
	case_insensitive := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifDefine(constant_name, value, nil, case_insensitive)
	returnValue.SetBool(ret)
})

// generate by ZifDefined
var DefZifDefined = def.DefFunc("defined", 1, 1, []def.ArgInfo{{name: "constant_name"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	constant_name := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifDefined(constant_name)
	returnValue.SetBool(ret)
})

// generate by ZifGetCalledClass
var DefZifGetCalledClass = def.DefFunc("get_called_class", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret, ok := ZifGetCalledClass(executeData)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})
