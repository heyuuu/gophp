package standard

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/def"
	"github.com/heyuuu/gophp/php/zpp"
)

var zifFunctions = []def.FuncType{DefZifUtf8Encode, DefZifVarDump}

// generate by ZifUtf8Encode
var DefZifUtf8Encode = def.DefFunc("utf8_encode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := php.NewParser(executeData, 1, 1, 0)
	data := fp.ParseString()
	if fp.HasError() {
		return
	}
	ret := ZifUtf8Encode(data)
	returnValue.SetString(ret)
})

// generate by ZifVarDump
var DefZifVarDump = def.DefFunc("var_dump", 0, -1, []def.ArgInfo{{Name: "vars"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !php.CheckNumArgs(executeData, 0, -1, 0) {
		return
	}
	fp := php.NewParser(executeData, 0, -1, 0)
	vars := fp.ParseVariadic(0)
	if fp.HasError() {
		return
	}
	ZifVarDump(executeData.Ctx(), vars)
})
