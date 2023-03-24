package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifStrptime
var DefZifStrptime = def.DefFunc("strptime", 2, 2, []def.ArgInfo{{name: "timestamp"}, {name: "format"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	timestamp := fp.ParseZval()
	format := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrptime(executeData, returnValue, timestamp, format)
})
