package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifSha1
var DefZifSha1 = def.DefFunc("sha1", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "raw_output"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseZval()
	fp.StartOptional()
	raw_output := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSha1(executeData, returnValue, str, nil, raw_output)
})

// generate by ZifSha1File
var DefZifSha1File = def.DefFunc("sha1_file", 1, 2, []def.ArgInfo{{Name: "filename"}, {Name: "raw_output"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	raw_output := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSha1File(executeData, returnValue, filename, nil, raw_output)
})
