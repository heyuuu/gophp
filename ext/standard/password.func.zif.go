package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifPasswordGetInfo
var DefZifPasswordGetInfo = def.DefFunc("password_get_info", 1, 1, []def.ArgInfo{{Name: "hash"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	hash_ := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifPasswordGetInfo(hash_)
	returnValue.SetBy(ret)
})

// generate by ZifPasswordNeedsRehash
var DefZifPasswordNeedsRehash = def.DefFunc("password_needs_rehash", 2, 3, []def.ArgInfo{{Name: "hash"}, {Name: "algo"}, {Name: "options"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	hash_ := fp.ParseStringVal()
	algo_ := fp.ParseZval()
	fp.StartOptional()
	options := fp.ParseArrayOrObjectHt()
	if fp.HasError() {
		return
	}
	ret := ZifPasswordNeedsRehash(hash_, algo_, nil, options)
	returnValue.SetBool(ret)
})

// generate by ZifPasswordVerify
var DefZifPasswordVerify = def.DefFunc("password_verify", 2, 2, []def.ArgInfo{{Name: "password"}, {Name: "hash"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	password := fp.ParseStringVal()
	hash := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifPasswordVerify(password, hash)
	returnValue.SetBool(ret)
})

// generate by ZifPasswordHash
var DefZifPasswordHash = def.DefFunc("password_hash", 2, 3, []def.ArgInfo{{Name: "password"}, {Name: "algo"}, {Name: "options"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	password := fp.ParseStringVal()
	algo_ := fp.ParseZval()
	fp.StartOptional()
	options := fp.ParseArrayOrObjectHt()
	if fp.HasError() {
		return
	}
	ret := ZifPasswordHash(password, algo_, nil, options)
	returnValue.SetBy(ret)
})

// generate by ZifPasswordAlgos
var DefZifPasswordAlgos = def.DefFunc("password_algos", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifPasswordAlgos()
	returnValue.SetArrayOfString(ret)
})
