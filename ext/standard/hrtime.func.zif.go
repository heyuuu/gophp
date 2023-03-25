package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifHrtime
var DefZifHrtime = def.DefFunc("hrtime", 1, 1, []def.ArgInfo{{Name: "get_as_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	get_as_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHrtime(executeData, returnValue, get_as_number)
})
