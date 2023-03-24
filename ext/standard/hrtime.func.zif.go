package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifHrtime
var DefZifHrtime = def.DefFunc("hrtime", 1, 1, []def.ArgInfo{{name: "get_as_number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	get_as_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHrtime(executeData, returnValue, get_as_number)
})
