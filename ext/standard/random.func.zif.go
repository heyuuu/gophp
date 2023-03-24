package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifRandomBytes
var DefZifRandomBytes = def.DefFunc("random_bytes", 1, 1, []def.ArgInfo{{Name: "length"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	length := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRandomBytes(executeData, returnValue, length)
})

// generate by ZifRandomInt
var DefZifRandomInt = def.DefFunc("random_int", 2, 2, []def.ArgInfo{{Name: "min"}, {Name: "max"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	min := fp.ParseZval()
	max := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRandomInt(executeData, returnValue, min, max)
})
