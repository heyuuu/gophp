package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifPasswordGetInfo
var DefZifPasswordGetInfo = def.DefFunc("password_get_info", 1, 1, []def.ArgInfo{{Name: "hash"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	hash := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPasswordGetInfo(hash)
})

// generate by ZifPasswordNeedsRehash
var DefZifPasswordNeedsRehash = def.DefFunc("password_needs_rehash", 2, 3, []def.ArgInfo{{Name: "hash"}, {Name: "algo"}, {Name: "options"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	hash := fp.ParseZval()
	algo := fp.ParseZval()
	fp.StartOptional()
	options := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPasswordNeedsRehash(hash, algo, nil, options)
})

// generate by ZifPasswordVerify
var DefZifPasswordVerify = def.DefFunc("password_verify", 2, 2, []def.ArgInfo{{Name: "password"}, {Name: "hash"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	password := fp.ParseZval()
	hash := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPasswordVerify(password, hash)
})

// generate by ZifPasswordHash
var DefZifPasswordHash = def.DefFunc("password_hash", 2, 3, []def.ArgInfo{{Name: "password"}, {Name: "algo"}, {Name: "options"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	password := fp.ParseZval()
	algo := fp.ParseZval()
	fp.StartOptional()
	options := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPasswordHash(password, algo, nil, options)
})

// generate by ZifPasswordAlgos
var DefZifPasswordAlgos = def.DefFunc("password_algos", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifPasswordAlgos()
})
