package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifConvertUuencode
var DefZifConvertUuencode = def.DefFunc("convert_uuencode", 1, 1, []def.ArgInfo{{name: "data"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifConvertUuencode(executeData, returnValue, data)
})

// generate by ZifConvertUudecode
var DefZifConvertUudecode = def.DefFunc("convert_uudecode", 1, 1, []def.ArgInfo{{name: "data"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifConvertUudecode(executeData, returnValue, data)
})
