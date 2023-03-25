package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifConvertUuencode
var DefZifConvertUuencode = def.DefFunc("convert_uuencode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifConvertUuencode(executeData, returnValue, data)
})

// generate by ZifConvertUudecode
var DefZifConvertUudecode = def.DefFunc("convert_uudecode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifConvertUudecode(executeData, returnValue, data)
})
