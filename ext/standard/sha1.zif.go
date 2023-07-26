package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifSha1
var DefZifSha1 = def.DefFunc("sha1", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "raw_output"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	raw_output := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifSha1(str, nil, raw_output)
	returnValue.SetString(ret)
})

// generate by ZifSha1File
var DefZifSha1File = def.DefFunc("sha1_file", 1, 2, []def.ArgInfo{{Name: "filename"}, {Name: "raw_output"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	filename := fp.ParseStringVal()
	fp.StartOptional()
	raw_output := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifSha1File(filename, nil, raw_output)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})
