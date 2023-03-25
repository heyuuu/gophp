package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifOpendir
var DefZifOpendir = def.DefFunc("opendir", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "context"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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
var DefZifGetdir = def.DefFunc("getdir", 1, 2, []def.ArgInfo{{Name: "directory"}, {Name: "context"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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
var DefZifClosedir = def.DefFunc("closedir", 0, 1, []def.ArgInfo{{Name: "dir_handle"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	dir_handle := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifClosedir(executeData, returnValue, nil, dir_handle)
})

// generate by ZifChroot
var DefZifChroot = def.DefFunc("chroot", 1, 1, []def.ArgInfo{{Name: "directory"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	directory := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifChroot(executeData, returnValue, directory)
})

// generate by ZifChdir
var DefZifChdir = def.DefFunc("chdir", 1, 1, []def.ArgInfo{{Name: "directory"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	directory := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifChdir(executeData, returnValue, directory)
})

// generate by ZifGetcwd
var DefZifGetcwd = def.DefFunc("getcwd", 0, 0, []def.ArgInfo{}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetcwd(executeData, returnValue)
})

// generate by ZifRewinddir
var DefZifRewinddir = def.DefFunc("rewinddir", 0, 1, []def.ArgInfo{{Name: "dir_handle"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	dir_handle := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRewinddir(executeData, returnValue, nil, dir_handle)
})

// generate by ZifGlob
var DefZifGlob = def.DefFunc("glob", 1, 2, []def.ArgInfo{{Name: "pattern"}, {Name: "flags"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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
var DefZifScandir = def.DefFunc("scandir", 1, 3, []def.ArgInfo{{Name: "dir"}, {Name: "sorting_order"}, {Name: "context"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
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
