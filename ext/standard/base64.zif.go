package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifBase64Encode
var DefZifBase64Encode = def.DefFunc("base64_encode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifBase64Encode(str)
	returnValue.SetString(ret)
})

// generate by ZifBase64Decode
var DefZifBase64Decode = def.DefFunc("base64_decode", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "strict"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	strict := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifBase64Decode(str, nil, strict)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})
