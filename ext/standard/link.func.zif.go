package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifReadlink
var DefZifReadlink = def.DefFunc("readlink", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifReadlink(executeData, returnValue, filename)
})

// generate by ZifLinkinfo
var DefZifLinkinfo = def.DefFunc("linkinfo", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLinkinfo(executeData, returnValue, filename)
})

// generate by ZifSymlink
var DefZifSymlink = def.DefFunc("symlink", 2, 2, []def.ArgInfo{{Name: "target"}, {Name: "link"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	target := fp.ParseZval()
	link := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSymlink(executeData, returnValue, target, link)
})

// generate by ZifLink
var DefZifLink = def.DefFunc("link", 2, 2, []def.ArgInfo{{Name: "target"}, {Name: "link"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	target := fp.ParseZval()
	link := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLink(executeData, returnValue, target, link)
})
