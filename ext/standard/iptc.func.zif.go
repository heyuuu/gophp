package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifIptcembed
var DefZifIptcembed = def.DefFunc("iptcembed", 2, 3, []def.ArgInfo{{Name: "iptcdata"}, {Name: "jpeg_file_name"}, {Name: "spool"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	iptcdata := fp.ParseZval()
	jpeg_file_name := fp.ParseZval()
	fp.StartOptional()
	spool := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIptcembed(executeData, returnValue, iptcdata, jpeg_file_name, nil, spool)
})

// generate by ZifIptcparse
var DefZifIptcparse = def.DefFunc("iptcparse", 1, 1, []def.ArgInfo{{Name: "iptcdata"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	iptcdata := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIptcparse(executeData, returnValue, iptcdata)
})
