package standard

import (
	"github.com/heyuuu/gophp/php/def"
	"github.com/heyuuu/gophp/php/zpp"
)

var zifFunctions = []def.FuncType{DefZifVarDump}

// generate by ZifVarDump
var DefZifVarDump = def.DefFunc("var_dump", 1, 1, []def.ArgInfo{{Name: "ctx"}, {Name: "value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgs(executeData, 1, 1, 0) {
		return
	}
	fp := zpp.NewParser(executeData, 1, 1, 0)
	value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifVarDump(executeData.Ctx(), executeData, value)
})
