package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifRand
var DefZifRand = def.DefFunc("rand", 0, 2, []def.ArgInfo{{Name: "min"}, {Name: "max"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	min := fp.ParseZval()
	max := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRand(executeData, returnValue, nil, min, max)
})
