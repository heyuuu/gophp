package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifNetGetInterfaces
var DefZifNetGetInterfaces = def.DefFunc("net_get_interfaces", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifNetGetInterfaces(executeData, returnValue)
})
