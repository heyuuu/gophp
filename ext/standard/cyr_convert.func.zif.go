package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifConvertCyrString
var DefZifConvertCyrString = def.DefFunc("convert_cyr_string", 3, 3, []def.ArgInfo{{Name: "str"}, {Name: "from"}, {Name: "to"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	str := fp.ParseZval()
	from := fp.ParseZval()
	to := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifConvertCyrString(executeData, returnValue, str, from, to)
})
