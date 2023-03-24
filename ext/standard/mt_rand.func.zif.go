package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifMtSrand
var DefZifMtSrand = def.DefFunc("mt_srand", 0, 2, []def.ArgInfo{{name: "seed"}, {name: "mode"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifMtRand = def.DefFunc("mt_rand", 0, 2, []def.ArgInfo{{name: "min"}, {name: "max"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifMtGetrandmax = def.DefFunc("mt_getrandmax", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifMtGetrandmax(executeData, returnValue)
})
