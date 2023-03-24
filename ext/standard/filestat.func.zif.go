package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifTouch
var DefZifTouch = def.DefFunc("touch", 1, 3, []def.ArgInfo{{name: "filename"}, {name: "time"}, {name: "atime"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	time := fp.ParseZval()
	atime := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTouch(executeData, returnValue, filename, nil, time, atime)
})

// generate by ZifClearstatcache
var DefZifClearstatcache = def.DefFunc("clearstatcache", 0, 2, []def.ArgInfo{{name: "clear_realpath_cache"}, {name: "filename"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	clear_realpath_cache := fp.ParseZval()
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifClearstatcache(executeData, returnValue, nil, clear_realpath_cache, filename)
})
