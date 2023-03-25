package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifLcgValue
var DefZifLcgValue = def.DefFunc("lcg_value", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifLcgValue(executeData, returnValue)
})
