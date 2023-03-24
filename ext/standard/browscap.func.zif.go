package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifGetBrowser
var DefZifGetBrowser = def.DefFunc("get_browser", 0, 2, []def.ArgInfo{{name: "browser_name"}, {name: "return_array"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	browser_name := fp.ParseZval()
	return_array := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetBrowser(executeData, returnValue, nil, browser_name, return_array)
})
