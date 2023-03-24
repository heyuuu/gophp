package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifGetmyuid
var DefZifGetmyuid = def.DefFunc("getmyuid", 0, 0, []def.ArgInfo{}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetmyuid(executeData, returnValue)
})

// generate by ZifGetmygid
var DefZifGetmygid = def.DefFunc("getmygid", 0, 0, []def.ArgInfo{}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetmygid(executeData, returnValue)
})

// generate by ZifGetmypid
var DefZifGetmypid = def.DefFunc("getmypid", 0, 0, []def.ArgInfo{}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetmypid(executeData, returnValue)
})

// generate by ZifGetmyinode
var DefZifGetmyinode = def.DefFunc("getmyinode", 0, 0, []def.ArgInfo{}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetmyinode(executeData, returnValue)
})

// generate by ZifGetlastmod
var DefZifGetlastmod = def.DefFunc("getlastmod", 0, 0, []def.ArgInfo{}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetlastmod(executeData, returnValue)
})
