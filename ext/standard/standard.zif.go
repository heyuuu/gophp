package standard

import (
	"github.com/heyuuu/gophp/php/def"
	"github.com/heyuuu/gophp/php/zpp"
)

var zifFunctions = []def.FuncType{DefZifVarDump}

// generate by ZifVarDump
var DefZifVarDump = def.DefFunc("var_dump", 0, -1, []def.ArgInfo{{Name: "ctx"}, {Name: "vars"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgs(executeData, 0, -1, 0) {
		return
	}
	fp := zpp.NewParser(executeData, 0, -1, 0)
	vars := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifVarDump(executeData.Ctx(), vars)
})
