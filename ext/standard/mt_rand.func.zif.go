package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifMtSrand
var DefZifMtSrand = def.DefFunc("mt_srand", 0, 2, []def.ArgInfo{{Name: "seed"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	seed := fp.ParseZval()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMtSrand(executeData, returnValue, nil, seed, mode)
})

// generate by ZifMtSrand
var DefZifSrand = def.DefFunc("srand", 0, 2, []def.ArgInfo{{Name: "seed"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	seed := fp.ParseZval()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMtSrand(executeData, returnValue, nil, seed, mode)
})

// generate by ZifMtRand
var DefZifMtRand = def.DefFunc("mt_rand", 0, 2, []def.ArgInfo{{Name: "min"}, {Name: "max"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	min := fp.ParseZval()
	max := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMtRand(executeData, returnValue, nil, min, max)
})

// generate by ZifMtGetrandmax
var DefZifMtGetrandmax = def.DefFunc("mt_getrandmax", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifMtGetrandmax(executeData, returnValue)
})

// generate by ZifMtGetrandmax
var DefZifGetrandmax = def.DefFunc("getrandmax", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifMtGetrandmax(executeData, returnValue)
})
