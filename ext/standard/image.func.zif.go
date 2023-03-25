package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifImageTypeToMimeType
var DefZifImageTypeToMimeType = def.DefFunc("image_type_to_mime_type", 1, 1, []def.ArgInfo{{Name: "imagetype"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	imagetype := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifImageTypeToMimeType(executeData, returnValue, imagetype)
})

// generate by ZifImageTypeToExtension
var DefZifImageTypeToExtension = def.DefFunc("image_type_to_extension", 1, 2, []def.ArgInfo{{Name: "imagetype"}, {Name: "include_dot"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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
var DefZifGetimagesize = def.DefFunc("getimagesize", 1, 2, []def.ArgInfo{{Name: "imagefile"}, {Name: "info"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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
var DefZifGetimagesizefromstring = def.DefFunc("getimagesizefromstring", 1, 2, []def.ArgInfo{{Name: "imagefile"}, {Name: "info"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	imagefile := fp.ParseZval()
	fp.StartOptional()
	info := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifGetimagesizefromstring(executeData, returnValue, imagefile, nil, info)
})
