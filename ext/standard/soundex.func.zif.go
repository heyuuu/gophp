package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifSoundex
var DefZifSoundex = def.DefFunc("soundex", 1, 1, []def.ArgInfo{{name: "str"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSoundex(executeData, returnValue, str)
})
