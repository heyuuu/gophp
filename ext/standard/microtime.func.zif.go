package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifMicrotime
var DefZifMicrotime = def.DefFunc("microtime", 0, 1, []def.ArgInfo{{Name: "get_as_float"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	get_as_float := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMicrotime(executeData, returnValue, nil, get_as_float)
})

// generate by ZifGettimeofday
var DefZifGettimeofday = def.DefFunc("gettimeofday", 0, 1, []def.ArgInfo{{Name: "get_as_float"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	get_as_float := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGettimeofday(executeData, returnValue, nil, get_as_float)
})

// generate by ZifGetrusage
var DefZifGetrusage = def.DefFunc("getrusage", 0, 1, []def.ArgInfo{{Name: "who"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	who := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetrusage(executeData, returnValue, nil, who)
})
