package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifSoundex
var DefZifSoundex = def.DefFunc("soundex", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSoundex(executeData, returnValue, str)
})
