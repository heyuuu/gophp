package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifFtok
var DefZifFtok = def.DefFunc("ftok", 2, 2, []def.ArgInfo{{Name: "pathname"}, {Name: "proj"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	pathname := fp.ParseZval()
	proj := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFtok(executeData, returnValue, pathname, proj)
})
