package builtin

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/def"
	"github.com/heyuuu/gophp/php/zpp"
)

var zifFunctions = []def.FuncType{DefZifGcMemCaches, DefZifGcCollectCycles, DefZifGcEnabled, DefZifGcStatus, DefZifStrlen, DefZifStrcmp, DefZifStrncmp, DefZifStrcasecmp, DefZifEach, DefZifErrorReporting}

// generate by ZifGcMemCaches
var DefZifGcMemCaches = def.DefFunc("gc_mem_caches", 0, 0, []def.ArgInfo{}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 0, 0)
	fp.CheckNumArgs()
	if fp.HasError() {
		return
	}
	ret := ZifGcMemCaches()
	returnValue.SetLong(ret)
})

// generate by ZifGcCollectCycles
var DefZifGcCollectCycles = def.DefFunc("gc_collect_cycles", 0, 0, []def.ArgInfo{}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 0, 0)
	fp.CheckNumArgs()
	if fp.HasError() {
		return
	}
	ret := ZifGcCollectCycles()
	returnValue.SetLong(ret)
})

// generate by ZifGcEnabled
var DefZifGcEnabled = def.DefFunc("gc_enabled", 0, 0, []def.ArgInfo{}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 0, 0)
	fp.CheckNumArgs()
	if fp.HasError() {
		return
	}
	ret := ZifGcEnabled()
	returnValue.SetBool(ret)
})

// generate by ZifGcStatus
var DefZifGcStatus = def.DefFunc("gc_status", 0, 0, []def.ArgInfo{}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 0, 0)
	fp.CheckNumArgs()
	if fp.HasError() {
		return
	}
	ret := ZifGcStatus()
	returnValue.SetArray(ret)
})

// generate by ZifStrlen
var DefZifStrlen = def.DefFunc("strlen", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, 0)
	fp.CheckNumArgs()
	str := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrlen(str)
	returnValue.SetLong(ret)
})

// generate by ZifStrcmp
var DefZifStrcmp = def.DefFunc("strcmp", 2, 2, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	str1 := fp.ParseString()
	str2 := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrcmp(str1, str2)
	returnValue.SetLong(ret)
})

// generate by ZifStrncmp
var DefZifStrncmp = def.DefFunc("strncmp", 3, 3, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}, {Name: "len"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 3, 3, 0)
	fp.CheckNumArgs()
	str1 := fp.ParseString()
	str2 := fp.ParseString()
	len_ := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrncmp(executeData.Ctx(), str1, str2, len_)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrcasecmp
var DefZifStrcasecmp = def.DefFunc("strcasecmp", 2, 2, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 2, 2, 0)
	fp.CheckNumArgs()
	str1 := fp.ParseString()
	str2 := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifStrcasecmp(str1, str2)
	returnValue.SetLong(ret)
})

// generate by ZifEach
var DefZifEach = def.DefFunc("each", 1, 1, []def.ArgInfo{{Name: "array", ByRef: true}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 1, 1, zpp.FlagOldMode)
	fp.CheckNumArgs()
	array := fp.ParseRefZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifEach(executeData.Ctx(), array)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifErrorReporting
var DefZifErrorReporting = def.DefFunc("error_reporting", 0, 1, []def.ArgInfo{{Name: "new_error_level"}}, func(executeData *php.ExecuteData, returnValue zpp.Ret) {
	fp := php.NewParamParser(executeData, 0, 1, 0)
	fp.CheckNumArgs()
	fp.StartOptional()
	new_error_level := fp.ParseZvalNullable()
	if fp.HasError() {
		return
	}
	ret := ZifErrorReporting(executeData.Ctx(), returnValue, nil, new_error_level)
	returnValue.SetLong(ret)
})
