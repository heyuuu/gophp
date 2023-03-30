package tokenizer

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifTokenGetAll
var DefZifTokenGetAll = def.DefFunc("token_get_all", 1, 2, []def.ArgInfo{{Name: "source"}, {Name: "flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	source := fp.ParseStringVal()
	fp.StartOptional()
	flags := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ZifTokenGetAll(returnValue, source, nil, flags)
})

// generate by ZifTokenName
var DefZifTokenName = def.DefFunc("token_name", 1, 1, []def.ArgInfo{{Name: "token"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	token := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifTokenName(token)
	returnValue.SetStringVal(ret)
})
