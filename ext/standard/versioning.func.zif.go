package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifVersionCompare
var DefZifVersionCompare = def.DefFunc("version_compare", 2, 3, []def.ArgInfo{{Name: "ver1"}, {Name: "ver2"}, {Name: "oper"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	ver1 := fp.ParseZval()
	ver2 := fp.ParseZval()
	fp.StartOptional()
	oper := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifVersionCompare(executeData, returnValue, ver1, ver2, nil, oper)
})
