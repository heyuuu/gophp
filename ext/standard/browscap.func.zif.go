package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifGetBrowser
var DefZifGetBrowser = def.DefFunc("get_browser", 0, 2, []def.ArgInfo{{Name: "browser_name"}, {Name: "return_array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	browser_name := fp.ParseZval()
	return_array := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetBrowser(executeData, returnValue, nil, browser_name, return_array)
})
