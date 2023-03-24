package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifEzmlmHash
var DefZifEzmlmHash = def.DefFunc("ezmlm_hash", 1, 1, []def.ArgInfo{{Name: "addr"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	addr := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifEzmlmHash(executeData, returnValue, addr)
})

// generate by ZifMail
var DefZifMail = def.DefFunc("mail", 3, 5, []def.ArgInfo{{Name: "to"}, {Name: "subject"}, {Name: "message"}, {Name: "additional_headers"}, {Name: "additional_parameters"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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
