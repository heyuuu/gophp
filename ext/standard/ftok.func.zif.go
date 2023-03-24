package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifFtok
var DefZifFtok = def.DefFunc("ftok", 2, 2, []def.ArgInfo{{name: "pathname"}, {name: "proj"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	pathname := fp.ParseZval()
	proj := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFtok(executeData, returnValue, pathname, proj)
})
