package core

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifHeaderRegisterCallback
var DefZifHeaderRegisterCallback = def.DefFunc("header_register_callback", 1, 1, []def.ArgInfo{{Name: "callback"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	callback := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifHeaderRegisterCallback(callback)
	returnValue.SetBool(ret)
})

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

// generate by ZifObStart
var DefZifObStart = def.DefFunc("ob_start", 0, 3, []def.ArgInfo{{Name: "user_function"}, {Name: "chunk_size"}, {Name: "flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 3, 0)
	fp.StartOptional()
	user_function := fp.ParseZval()
	chunk_size := fp.ParseLong()
	flags_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifObStart(nil, user_function, chunk_size, flags_)
	returnValue.SetBool(ret)
})

// generate by ZifObFlush
var DefZifObFlush = def.DefFunc("ob_flush", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifObFlush(executeData, returnValue)
})

// generate by ZifObClean
var DefZifObClean = def.DefFunc("ob_clean", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifObClean(executeData, returnValue)
})

// generate by ZifObEndFlush
var DefZifObEndFlush = def.DefFunc("ob_end_flush", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifObEndFlush(executeData, returnValue)
})

// generate by ZifObEndClean
var DefZifObEndClean = def.DefFunc("ob_end_clean", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifObEndClean(executeData, returnValue)
})

// generate by ZifObGetFlush
var DefZifObGetFlush = def.DefFunc("ob_get_flush", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifObGetFlush(executeData, returnValue)
})

// generate by ZifObGetClean
var DefZifObGetClean = def.DefFunc("ob_get_clean", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifObGetClean(executeData, returnValue)
})

// generate by ZifObGetContents
var DefZifObGetContents = def.DefFunc("ob_get_contents", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifObGetContents(executeData, returnValue)
})

// generate by ZifObGetLevel
var DefZifObGetLevel = def.DefFunc("ob_get_level", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifObGetLevel(executeData, returnValue)
})

// generate by ZifObGetLength
var DefZifObGetLength = def.DefFunc("ob_get_length", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifObGetLength(executeData, returnValue)
})

// generate by ZifObListHandlers
var DefZifObListHandlers = def.DefFunc("ob_list_handlers", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifObListHandlers(executeData, returnValue)
})

// generate by ZifObGetStatus
var DefZifObGetStatus = def.DefFunc("ob_get_status", 0, 1, []def.ArgInfo{{Name: "full_status"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	full_status := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifObGetStatus(nil, full_status)
	returnValue.SetArray(ret)
})

// generate by ZifObImplicitFlush
var DefZifObImplicitFlush = def.DefFunc("ob_implicit_flush", 0, 1, []def.ArgInfo{{Name: "flag"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	flag := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifObImplicitFlush(executeData, returnValue, nil, flag)
})

// generate by ZifOutputResetRewriteVars
var DefZifOutputResetRewriteVars = def.DefFunc("output_reset_rewrite_vars", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifOutputResetRewriteVars(executeData, returnValue)
})

// generate by ZifOutputAddRewriteVar
var DefZifOutputAddRewriteVar = def.DefFunc("output_add_rewrite_var", 2, 2, []def.ArgInfo{{Name: "name"}, {Name: "value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	name := fp.ParseZval()
	value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifOutputAddRewriteVar(executeData, returnValue, name, value)
})
