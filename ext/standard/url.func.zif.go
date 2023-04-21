package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifParseUrl
var DefZifParseUrl = def.DefFunc("parse_url", 1, 2, []def.ArgInfo{{Name: "url"}, {Name: "component"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	url := fp.ParseZval()
	fp.StartOptional()
	component := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifParseUrl(executeData, returnValue, url, nil, component)
})

// generate by ZifUrlencode
var DefZifUrlencode = def.DefFunc("urlencode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifUrlencode(str)
	returnValue.SetStringVal(ret)
})

// generate by ZifUrldecode
var DefZifUrldecode = def.DefFunc("urldecode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifUrldecode(str)
	returnValue.SetStringVal(ret)
})

// generate by ZifRawurlencode
var DefZifRawurlencode = def.DefFunc("rawurlencode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRawurlencode(executeData, returnValue, str)
})

// generate by ZifRawurldecode
var DefZifRawurldecode = def.DefFunc("rawurldecode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifRawurldecode(str)
	returnValue.SetStringVal(ret)
})

// generate by ZifGetHeaders
var DefZifGetHeaders = def.DefFunc("get_headers", 1, 3, []def.ArgInfo{{Name: "url"}, {Name: "format"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	url := fp.ParseZval()
	fp.StartOptional()
	format := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetHeaders(executeData, returnValue, url, nil, format, context)
})
