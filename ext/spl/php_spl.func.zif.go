package spl

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifClassParents
var DefZifClassParents = def.DefFunc("class_parents", 1, 2, []def.ArgInfo{{Name: "instance"}, {Name: "autoload"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	instance := fp.ParseZval()
	fp.StartOptional()
	autoload := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifClassParents(executeData, returnValue, instance, nil, autoload)
})

// generate by ZifClassImplements
var DefZifClassImplements = def.DefFunc("class_implements", 1, 2, []def.ArgInfo{{Name: "what"}, {Name: "autoload"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	what := fp.ParseZval()
	fp.StartOptional()
	autoload := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifClassImplements(executeData, returnValue, what, nil, autoload)
})

// generate by ZifClassUses
var DefZifClassUses = def.DefFunc("class_uses", 1, 2, []def.ArgInfo{{Name: "what"}, {Name: "autoload"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	what := fp.ParseZval()
	fp.StartOptional()
	autoload := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifClassUses(executeData, returnValue, what, nil, autoload)
})

// generate by ZifSplClasses
var DefZifSplClasses = def.DefFunc("spl_classes", 0, 0, []def.ArgInfo{}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifSplClasses(executeData, returnValue)
})

// generate by ZifSplAutoload
var DefZifSplAutoload = def.DefFunc("spl_autoload", 1, 2, []def.ArgInfo{{Name: "class_name"}, {Name: "file_extensions"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	class_name := fp.ParseZval()
	fp.StartOptional()
	file_extensions := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSplAutoload(executeData, returnValue, class_name, nil, file_extensions)
})

// generate by ZifSplAutoloadExtensions
var DefZifSplAutoloadExtensions = def.DefFunc("spl_autoload_extensions", 0, 1, []def.ArgInfo{{Name: "file_extensions"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	file_extensions := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSplAutoloadExtensions(executeData, returnValue, nil, file_extensions)
})

// generate by ZifSplAutoloadCall
var DefZifSplAutoloadCall = def.DefFunc("spl_autoload_call", 1, 1, []def.ArgInfo{{Name: "class_name"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	class_name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSplAutoloadCall(executeData, returnValue, class_name)
})

// generate by ZifSplAutoloadRegister
var DefZifSplAutoloadRegister = def.DefFunc("spl_autoload_register", 0, 3, []def.ArgInfo{{Name: "autoload_function"}, {Name: "throw"}, {Name: "prepend"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 0, 3, 0)
	fp.StartOptional()
	autoload_function := fp.ParseZval()
	throw := fp.ParseZval()
	prepend := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSplAutoloadRegister(executeData, returnValue, nil, autoload_function, throw, prepend)
})

// generate by ZifSplAutoloadUnregister
var DefZifSplAutoloadUnregister = def.DefFunc("spl_autoload_unregister", 1, 1, []def.ArgInfo{{Name: "autoload_function"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	autoload_function := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSplAutoloadUnregister(executeData, returnValue, autoload_function)
})

// generate by ZifSplAutoloadFunctions
var DefZifSplAutoloadFunctions = def.DefFunc("spl_autoload_functions", 0, 0, []def.ArgInfo{}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifSplAutoloadFunctions(executeData, returnValue)
})

// generate by ZifSplObjectHash
var DefZifSplObjectHash = def.DefFunc("spl_object_hash", 1, 1, []def.ArgInfo{{Name: "obj"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	obj := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSplObjectHash(executeData, returnValue, obj)
})

// generate by ZifSplObjectId
var DefZifSplObjectId = def.DefFunc("spl_object_id", 1, 1, []def.ArgInfo{{Name: "obj"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	obj := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSplObjectId(executeData, returnValue, obj)
})
