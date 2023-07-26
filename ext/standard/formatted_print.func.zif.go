package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifUserSprintf
var DefZifSprintf = def.DefFunc("sprintf", 1, -1, []def.ArgInfo{{Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	format := fp.ParseZval()
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret, ok := ZifUserSprintf(format, nil, args)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifVsprintf
var DefZifVsprintf = def.DefFunc("vsprintf", 2, 2, []def.ArgInfo{{Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	format := fp.ParseZval()
	args := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifVsprintf(executeData, returnValue, format, args)
})

// generate by ZifUserPrintf
var DefZifPrintf = def.DefFunc("printf", 1, -1, []def.ArgInfo{{Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	format := fp.ParseZval()
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifUserPrintf(executeData, returnValue, format, nil, args)
})

// generate by ZifVprintf
var DefZifVprintf = def.DefFunc("vprintf", 2, 2, []def.ArgInfo{{Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	format := fp.ParseZval()
	args := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifVprintf(executeData, returnValue, format, args)
})

// generate by ZifFprintf
var DefZifFprintf = def.DefFunc("fprintf", 2, -1, []def.ArgInfo{{Name: "stream"}, {Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, -1, 0)
	stream := fp.ParseZval()
	format := fp.ParseZval()
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifFprintf(executeData, returnValue, stream, format, nil, args)
})

// generate by ZifVfprintf
var DefZifVfprintf = def.DefFunc("vfprintf", 3, 3, []def.ArgInfo{{Name: "stream"}, {Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	stream := fp.ParseZval()
	format := fp.ParseZval()
	args := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifVfprintf(executeData, returnValue, stream, format, args)
})
