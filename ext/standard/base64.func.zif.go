package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifBase64Encode
var DefZifBase64Encode = def.DefFunc("base64_encode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifBase64Encode(executeData, returnValue, str)
})

// generate by ZifBase64Decode
var DefZifBase64Decode = def.DefFunc("base64_decode", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "strict"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	strict := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifBase64Decode(executeData, returnValue, str, nil, strict)
})
