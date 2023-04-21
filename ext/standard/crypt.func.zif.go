package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifCrypt
var DefZifCrypt = def.DefFunc("crypt", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "salt"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str_ := fp.ParseStringVal()
	fp.StartOptional()
	salt_ := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifCrypt(executeData, returnValue, str_, nil, salt_)
	returnValue.SetStringVal(ret)
})
