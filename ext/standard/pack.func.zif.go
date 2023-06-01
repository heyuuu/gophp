package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifPack
var DefZifPack = def.DefFunc("pack", 1, -1, []def.ArgInfo{{Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	format_ := fp.ParseStringVal()
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret, ok := ZifPack(format_, nil, args)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifUnpack
var DefZifUnpack = def.DefFunc("unpack", 2, 3, []def.ArgInfo{{Name: "format"}, {Name: "input"}, {Name: "offset"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
