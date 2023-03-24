package standard

import (
	"sik/zend/types"
	"sik/zend/zpp"
	"sik/zend/def"
)

// generate by ZifVersionCompare
var DefZifVersionCompare = def.DefFunc("version_compare", 2, 3, []def.ArgInfo{{name: "ver1"}, {name: "ver2"}, {name: "oper"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
