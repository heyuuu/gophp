package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifEzmlmHash
var DefZifEzmlmHash = def.DefFunc("ezmlm_hash", 1, 1, []def.ArgInfo{{name: "addr"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	addr := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifEzmlmHash(executeData, returnValue, addr)
})

// generate by ZifMail
var DefZifMail = def.DefFunc("mail", 3, 5, []def.ArgInfo{{name: "to"}, {name: "subject"}, {name: "message"}, {name: "additional_headers"}, {name: "additional_parameters"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 3, 5, 0)
	to := fp.ParseZval()
	subject := fp.ParseZval()
	message := fp.ParseZval()
	fp.StartOptional()
	additional_headers := fp.ParseZval()
	additional_parameters := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMail(executeData, returnValue, to, subject, message, nil, additional_headers, additional_parameters)
})
