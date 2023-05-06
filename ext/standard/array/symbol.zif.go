package array

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifExtract
var DefZifExtract = def.DefFunc("extract", 1, 3, []def.ArgInfo{{Name: "array"}, {Name: "flags"}, {Name: "prefix"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	array := fp.ParseArrayEx2(false, true, false)
	fp.StartOptional()
	flags := fp.ParseLong()
	prefix := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifExtract(array, nil, flags, prefix)
	returnValue.SetLong(ret)
})

// generate by ZifCompact
var DefZifCompact = def.DefFunc("compact", 0, -1, []def.ArgInfo{{Name: "var_names"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	var_names := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret := ZifCompact(var_names)
	returnValue.SetArray(ret)
})
