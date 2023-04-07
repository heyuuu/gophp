package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifMd5
var DefZifMd5 = def.DefFunc("md5", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "raw_output"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	raw_output := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifMd5(str, nil, raw_output)
	returnValue.SetStringVal(ret)
})

// generate by ZifMd5File
var DefZifMd5File = def.DefFunc("md5_file", 1, 2, []def.ArgInfo{{Name: "filename"}, {Name: "raw_output"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	filename := fp.ParseStringVal()
	fp.StartOptional()
	raw_output := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifMd5File(filename, nil, raw_output)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})
