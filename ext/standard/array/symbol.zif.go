package array

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifCompact
var DefZifCompact = def.DefFunc("compact", 0, -1, []def.ArgInfo{{Name: "var_names"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	var_names := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifCompact(executeData, returnValue, var_names)
})
