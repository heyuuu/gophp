package streams

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifStreamWrapperRegister
var DefZifStreamWrapperRegister = def.DefFunc("stream_wrapper_register", 2, 3, []def.ArgInfo{{Name: "protocol"}, {Name: "classname"}, {Name: "flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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

// generate by ZifStreamWrapperRegister
var DefZifStreamRegisterWrapper = def.DefFunc("stream_register_wrapper", 2, 3, []def.ArgInfo{{Name: "protocol"}, {Name: "classname"}, {Name: "flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifStreamWrapperUnregister = def.DefFunc("stream_wrapper_unregister", 1, 1, []def.ArgInfo{{Name: "protocol"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	protocol := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamWrapperUnregister(executeData, returnValue, protocol)
})

// generate by ZifStreamWrapperRestore
var DefZifStreamWrapperRestore = def.DefFunc("stream_wrapper_restore", 1, 1, []def.ArgInfo{{Name: "protocol"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	protocol := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStreamWrapperRestore(executeData, returnValue, protocol)
	returnValue.SetBool(ret)
})
