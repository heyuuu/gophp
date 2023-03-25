package core

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifSetTimeLimit
var DefZifSetTimeLimit = def.DefFunc("set_time_limit", 1, 1, []def.ArgInfo{{Name: "seconds"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	seconds := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifSetTimeLimit(seconds)
	returnValue.SetBool(ret)
})
