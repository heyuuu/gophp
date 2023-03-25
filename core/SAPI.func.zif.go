package core

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifHeaderRegisterCallback
var DefZifHeaderRegisterCallback = def.DefFunc("header_register_callback", 1, 1, []def.ArgInfo{{Name: "callback"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	callback := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifHeaderRegisterCallback(callback)
	returnValue.SetBool(ret)
})
