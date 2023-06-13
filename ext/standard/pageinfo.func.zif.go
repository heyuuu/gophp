package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifGetmyuid
var DefZifGetmyuid = def.DefFunc("getmyuid", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetmyuid()
})

// generate by ZifGetmygid
var DefZifGetmygid = def.DefFunc("getmygid", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetmygid()
})

// generate by ZifGetmypid
var DefZifGetmypid = def.DefFunc("getmypid", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetmypid()
})

// generate by ZifGetmyinode
var DefZifGetmyinode = def.DefFunc("getmyinode", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetmyinode()
})

// generate by ZifGetlastmod
var DefZifGetlastmod = def.DefFunc("getlastmod", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetlastmod(executeData, returnValue)
})
