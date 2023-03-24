package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifPasswordGetInfo
var DefZifPasswordGetInfo = def.DefFunc("password_get_info", 1, 1, []def.ArgInfo{{Name: "hash"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	hash := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPasswordGetInfo(executeData, returnValue, hash)
})

// generate by ZifPasswordNeedsRehash
var DefZifPasswordNeedsRehash = def.DefFunc("password_needs_rehash", 2, 3, []def.ArgInfo{{Name: "hash"}, {Name: "algo"}, {Name: "options"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	hash := fp.ParseZval()
	algo := fp.ParseZval()
	fp.StartOptional()
	options := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPasswordNeedsRehash(executeData, returnValue, hash, algo, nil, options)
})

// generate by ZifPasswordVerify
var DefZifPasswordVerify = def.DefFunc("password_verify", 2, 2, []def.ArgInfo{{Name: "password"}, {Name: "hash"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	password := fp.ParseZval()
	hash := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPasswordVerify(executeData, returnValue, password, hash)
})

// generate by ZifPasswordHash
var DefZifPasswordHash = def.DefFunc("password_hash", 2, 3, []def.ArgInfo{{Name: "password"}, {Name: "algo"}, {Name: "options"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	password := fp.ParseZval()
	algo := fp.ParseZval()
	fp.StartOptional()
	options := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPasswordHash(executeData, returnValue, password, algo, nil, options)
})

// generate by ZifPasswordAlgos
var DefZifPasswordAlgos = def.DefFunc("password_algos", 0, 0, []def.ArgInfo{}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifPasswordAlgos(executeData, returnValue)
})
