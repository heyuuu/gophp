package streams

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
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

// generate by ZifStreamWrapperUnregister
var DefZifStreamWrapperUnregister = def.DefFunc("stream_wrapper_unregister", 1, 1, []def.ArgInfo{{name: "protocol"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	protocol := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamWrapperUnregister(executeData, returnValue, protocol)
})

// generate by ZifStreamWrapperRestore
var DefZifStreamWrapperRestore = def.DefFunc("stream_wrapper_restore", 1, 1, []def.ArgInfo{{name: "protocol"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	protocol := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamWrapperRestore(executeData, returnValue, protocol)
})
