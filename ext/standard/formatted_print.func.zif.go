package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifUserSprintf
var DefZifUserSprintf = def.DefFunc("user_sprintf", 1, -1, []def.ArgInfo{{name: "format"}, {name: "args"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	format := fp.ParseZval()
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifUserSprintf(executeData, returnValue, format, nil, args)
})

// generate by ZifUserPrintf
var DefZifUserPrintf = def.DefFunc("user_printf", 1, -1, []def.ArgInfo{{name: "format"}, {name: "args"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	format := fp.ParseZval()
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifUserPrintf(executeData, returnValue, format, nil, args)
})

// generate by ZifFprintf
var DefZifFprintf = def.DefFunc("fprintf", 2, -1, []def.ArgInfo{{name: "stream"}, {name: "format"}, {name: "args"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, -1, 0)
	stream := fp.ParseZval()
	format := fp.ParseZval()
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifFprintf(executeData, returnValue, stream, format, nil, args)
})
