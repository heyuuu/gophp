package spl

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifClassParents
var DefZifClassParents = def.DefFunc("class_parents", 1, 2, []def.ArgInfo{{Name: "instance"}, {Name: "autoload"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	instance := fp.ParseZval()
	fp.StartOptional()
	autoload_ := fp.ParseBoolValNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifClassParents(instance, nil, autoload_)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifClassImplements
var DefZifClassImplements = def.DefFunc("class_implements", 1, 2, []def.ArgInfo{{Name: "instance"}, {Name: "autoload"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	instance := fp.ParseZval()
	fp.StartOptional()
	autoload_ := fp.ParseBoolValNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifClassImplements(instance, nil, autoload_)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifClassUses
var DefZifClassUses = def.DefFunc("class_uses", 1, 2, []def.ArgInfo{{Name: "instance"}, {Name: "autoload"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	instance := fp.ParseZval()
	fp.StartOptional()
	autoload_ := fp.ParseBoolValNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifClassUses(instance, nil, autoload_)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifSplClasses
var DefZifSplClasses = def.DefFunc("spl_classes", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifSplClasses(executeData, returnValue)
})

// generate by ZifSplAutoload
var DefZifSplAutoload = def.DefFunc("spl_autoload", 1, 2, []def.ArgInfo{{Name: "class_name"}, {Name: "file_extensions"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	class_name := fp.ParseStringVal()
	fp.StartOptional()
	file_extensions := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ZifSplAutoload(class_name, nil, file_extensions)
})

// generate by ZifSplAutoloadExtensions
var DefZifSplAutoloadExtensions = def.DefFunc("spl_autoload_extensions", 0, 1, []def.ArgInfo{{Name: "file_extensions"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	file_extensions := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifSplAutoloadExtensions(nil, file_extensions)
	returnValue.SetStringVal(ret)
})

// generate by ZifSplAutoloadCall
var DefZifSplAutoloadCall = def.DefFunc("spl_autoload_call", 1, 1, []def.ArgInfo{{Name: "class_name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	class_name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSplAutoloadCall(executeData, returnValue, class_name)
})

// generate by ZifSplAutoloadRegister
var DefZifSplAutoloadRegister = def.DefFunc("spl_autoload_register", 0, 3, []def.ArgInfo{{Name: "autoload_function"}, {Name: "throw"}, {Name: "prepend"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifSplAutoloadUnregister = def.DefFunc("spl_autoload_unregister", 1, 1, []def.ArgInfo{{Name: "autoload_function"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	autoload_function := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSplAutoloadUnregister(executeData, returnValue, autoload_function)
})

// generate by ZifSplAutoloadFunctions
var DefZifSplAutoloadFunctions = def.DefFunc("spl_autoload_functions", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifSplAutoloadFunctions(executeData, returnValue)
})

// generate by ZifSplObjectHash
var DefZifSplObjectHash = def.DefFunc("spl_object_hash", 1, 1, []def.ArgInfo{{Name: "obj"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, zpp.FlagOldMode)
	obj := fp.ParseObject()
	if fp.HasError() {
		return
	}
	ret := ZifSplObjectHash(obj)
	returnValue.SetStringVal(ret)
})

// generate by ZifSplObjectId
var DefZifSplObjectId = def.DefFunc("spl_object_id", 1, 1, []def.ArgInfo{{Name: "obj"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	obj := fp.ParseObject()
	if fp.HasError() {
		return
	}
	ret := ZifSplObjectId(obj)
	returnValue.SetLong(ret)
})
