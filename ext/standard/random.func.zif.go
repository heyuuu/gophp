package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifRandomBytes
var DefZifRandomBytes = def.DefFunc("random_bytes", 1, 1, []def.ArgInfo{{Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	length := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifRandomBytes(length)
	returnValue.SetStringVal(ret)
})

// generate by ZifRandomInt
var DefZifRandomInt = def.DefFunc("random_int", 2, 2, []def.ArgInfo{{Name: "min"}, {Name: "max"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	min := fp.ParseLong()
	max := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifRandomInt(min, max)
	returnValue.SetLong(ret)
})
