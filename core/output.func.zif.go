package core

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifObStart
var DefZifObStart = def.DefFunc("ob_start", 0, 3, []def.ArgInfo{{name: "user_function"}, {name: "chunk_size"}, {name: "flags"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 3, 0)
	fp.StartOptional()
	user_function := fp.ParseZval()
	chunk_size := fp.ParseZval()
	flags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifObStart(executeData, returnValue, nil, user_function, chunk_size, flags)
})

// generate by ZifObGetStatus
var DefZifObGetStatus = def.DefFunc("ob_get_status", 0, 1, []def.ArgInfo{{name: "full_status"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	full_status := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifObGetStatus(executeData, returnValue, nil, full_status)
})

// generate by ZifObImplicitFlush
var DefZifObImplicitFlush = def.DefFunc("ob_implicit_flush", 0, 1, []def.ArgInfo{{name: "flag"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	flag := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifObImplicitFlush(executeData, returnValue, nil, flag)
})
