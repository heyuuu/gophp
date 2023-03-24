package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifImageTypeToExtension
var DefZifImageTypeToExtension = def.DefFunc("image_type_to_extension", 1, 2, []def.ArgInfo{{name: "imagetype"}, {name: "include_dot"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	imagetype := fp.ParseZval()
	fp.StartOptional()
	include_dot := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifImageTypeToExtension(executeData, returnValue, imagetype, nil, include_dot)
})

// generate by ZifGetimagesize
var DefZifGetimagesize = def.DefFunc("getimagesize", 1, 2, []def.ArgInfo{{name: "imagefile"}, {name: "info"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	imagefile := fp.ParseZval()
	fp.StartOptional()
	info := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifGetimagesize(executeData, returnValue, imagefile, nil, info)
})

// generate by ZifGetimagesizefromstring
var DefZifGetimagesizefromstring = def.DefFunc("getimagesizefromstring", 1, 2, []def.ArgInfo{{name: "imagefile"}, {name: "info"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	imagefile := fp.ParseZval()
	fp.StartOptional()
	info := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifGetimagesizefromstring(executeData, returnValue, imagefile, nil, info)
})
