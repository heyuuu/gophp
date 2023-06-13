package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifVarDump
var DefZifVarDump = def.DefFunc("var_dump", 0, -1, []def.ArgInfo{{Name: "vars"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	vars := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifVarDump(vars)
})

// generate by ZifDebugZvalDump
var DefZifDebugZvalDump = def.DefFunc("debug_zval_dump", 0, -1, []def.ArgInfo{{Name: "vars"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	vars := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifDebugZvalDump(vars)
})

// generate by ZifVarExport
var DefZifVarExport = def.DefFunc("var_export", 1, 2, []def.ArgInfo{{Name: "var"}, {Name: "return"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	var_ := fp.ParseZval()
	fp.StartOptional()
	return_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifVarExport(executeData, returnValue, var_, nil, return_)
})

// generate by ZifSerialize
var DefZifSerialize = def.DefFunc("serialize", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSerialize(var_)
})

// generate by ZifUnserialize
var DefZifUnserialize = def.DefFunc("unserialize", 1, 2, []def.ArgInfo{{Name: "variable_representation"}, {Name: "allowed_classes"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	variable_representation := fp.ParseZval()
	fp.StartOptional()
	allowed_classes := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUnserialize(executeData, returnValue, variable_representation, nil, allowed_classes)
})

// generate by ZifMemoryGetUsage
var DefZifMemoryGetUsage = def.DefFunc("memory_get_usage", 0, 1, []def.ArgInfo{{Name: "real_usage"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	real_usage := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMemoryGetUsage(executeData, returnValue, nil, real_usage)
})

// generate by ZifMemoryGetPeakUsage
var DefZifMemoryGetPeakUsage = def.DefFunc("memory_get_peak_usage", 0, 1, []def.ArgInfo{{Name: "real_usage"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	real_usage := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMemoryGetPeakUsage(executeData, returnValue, nil, real_usage)
})
