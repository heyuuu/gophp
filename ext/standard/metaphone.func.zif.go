package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifMetaphone
var DefZifMetaphone = def.DefFunc("metaphone", 1, 2, []def.ArgInfo{{name: "text"}, {name: "phones"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	text := fp.ParseZval()
	fp.StartOptional()
	phones := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMetaphone(executeData, returnValue, text, nil, phones)
})
