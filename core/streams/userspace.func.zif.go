package streams

import (
	"sik/zend/types"
	"sik/zend/zpp"
	"sik/zend/def"
)

// generate by ZifStreamWrapperRegister
var DefZifStreamWrapperRegister = def.DefFunc("stream_wrapper_register", 2, 3, []def.ArgInfo{{name: "protocol"}, {name: "classname"}, {name: "flags"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	protocol := fp.ParseZval()
	classname := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamWrapperRegister(executeData, returnValue, protocol, classname, nil, flags)
})
