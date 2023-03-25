package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifStrptime
var DefZifStrptime = def.DefFunc("strptime", 2, 2, []def.ArgInfo{{Name: "timestamp"}, {Name: "format"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	timestamp := fp.ParseZval()
	format := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrptime(executeData, returnValue, timestamp, format)
})
