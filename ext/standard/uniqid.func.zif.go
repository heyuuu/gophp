package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifUniqid
var DefZifUniqid = def.DefFunc("uniqid", 0, 2, []def.ArgInfo{{Name: "prefix"}, {Name: "more_entropy"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	prefix := fp.ParseZval()
	more_entropy := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUniqid(executeData, returnValue, nil, prefix, more_entropy)
})
