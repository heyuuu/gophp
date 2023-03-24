package standard

import (
	"sik/zend/types"
	"sik/zend/zpp"
	"sik/zend/def"
)

// generate by ZifUniqid
var DefZifUniqid = def.DefFunc("uniqid", 0, 2, []def.ArgInfo{{name: "prefix"}, {name: "more_entropy"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	prefix := fp.ParseZval()
	more_entropy := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUniqid(executeData, returnValue, nil, prefix, more_entropy)
})
