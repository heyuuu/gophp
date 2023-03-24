package standard

import (
	"sik/zend/zpp"
	"sik/zend/def"
	"sik/zend/types"
)

// generate by ZifOpendir
var DefZifOpendir = def.DefFunc("opendir", 1, 2, []def.ArgInfo{{name: "path"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	path := fp.ParseZval()
	fp.StartOptional()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifOpendir(executeData, returnValue, path, nil, context)
})

// generate by ZifGetdir
var DefZifGetdir = def.DefFunc("getdir", 1, 2, []def.ArgInfo{{name: "directory"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	directory := fp.ParseZval()
	fp.StartOptional()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetdir(executeData, returnValue, directory, nil, context)
})

// generate by ZifClosedir
var DefZifClosedir = def.DefFunc("closedir", 0, 1, []def.ArgInfo{{name: "dir_handle"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	dir_handle := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifClosedir(executeData, returnValue, nil, dir_handle)
})

// generate by ZifRewinddir
var DefZifRewinddir = def.DefFunc("rewinddir", 0, 1, []def.ArgInfo{{name: "dir_handle"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	dir_handle := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRewinddir(executeData, returnValue, nil, dir_handle)
})

// generate by ZifGlob
var DefZifGlob = def.DefFunc("glob", 1, 2, []def.ArgInfo{{name: "pattern"}, {name: "flags"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	pattern := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGlob(executeData, returnValue, pattern, nil, flags)
})

// generate by ZifScandir
var DefZifScandir = def.DefFunc("scandir", 1, 3, []def.ArgInfo{{name: "dir"}, {name: "sorting_order"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	dir := fp.ParseZval()
	fp.StartOptional()
	sorting_order := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifScandir(executeData, returnValue, dir, nil, sorting_order, context)
})
