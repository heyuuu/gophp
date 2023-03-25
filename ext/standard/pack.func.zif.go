package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifPack
var DefZifPack = def.DefFunc("pack", 1, -1, []def.ArgInfo{{Name: "format"}, {Name: "args"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	format := fp.ParseZval()
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifPack(executeData, returnValue, format, nil, args)
})

// generate by ZifUnpack
var DefZifUnpack = def.DefFunc("unpack", 2, 3, []def.ArgInfo{{Name: "format"}, {Name: "input"}, {Name: "offset"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	format := fp.ParseZval()
	input := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUnpack(executeData, returnValue, format, input, nil, offset)
})
