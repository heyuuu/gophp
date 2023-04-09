package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifPathinfo
var DefZifPathinfo = def.DefFunc("pathinfo", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "options"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	path := fp.ParseZval()
	fp.StartOptional()
	options := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPathinfo(returnValue, path, nil)
})
