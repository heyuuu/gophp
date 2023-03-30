package core

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifSetTimeLimit
var DefZifSetTimeLimit = def.DefFunc("set_time_limit", 1, 1, []def.ArgInfo{{Name: "seconds"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	seconds := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifSetTimeLimit(seconds)
	returnValue.SetBool(ret)
})
