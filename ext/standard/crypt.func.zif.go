package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifCrypt
var DefZifCrypt = def.DefFunc("crypt", 1, 2, []def.ArgInfo{{name: "str"}, {name: "salt"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	salt := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCrypt(executeData, returnValue, str, nil, salt)
})
