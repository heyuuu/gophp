package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifMetaphone
var DefZifMetaphone = def.DefFunc("metaphone", 1, 2, []def.ArgInfo{{Name: "text"}, {Name: "phones"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	text := fp.ParseZval()
	fp.StartOptional()
	phones := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMetaphone(executeData, returnValue, text, nil, phones)
})
