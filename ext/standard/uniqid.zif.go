package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifUniqid
var DefZifUniqid = def.DefFunc("uniqid", 0, 2, []def.ArgInfo{{Name: "prefix"}, {Name: "more_entropy"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	prefix := fp.ParseStringVal()
	more_entropy := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifUniqid(nil, prefix, more_entropy)
	returnValue.SetString(ret)
})
