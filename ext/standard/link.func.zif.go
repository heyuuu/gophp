package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifReadlink
var DefZifReadlink = def.DefFunc("readlink", 1, 1, []def.ArgInfo{{name: "filename"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifReadlink(executeData, returnValue, filename)
})

// generate by ZifLinkinfo
var DefZifLinkinfo = def.DefFunc("linkinfo", 1, 1, []def.ArgInfo{{name: "filename"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLinkinfo(executeData, returnValue, filename)
})

// generate by ZifSymlink
var DefZifSymlink = def.DefFunc("symlink", 2, 2, []def.ArgInfo{{name: "target"}, {name: "link"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	target := fp.ParseZval()
	link := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSymlink(executeData, returnValue, target, link)
})

// generate by ZifLink
var DefZifLink = def.DefFunc("link", 2, 2, []def.ArgInfo{{name: "target"}, {name: "link"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	target := fp.ParseZval()
	link := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLink(executeData, returnValue, target, link)
})
