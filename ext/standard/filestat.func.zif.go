package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifDiskTotalSpace
var DefZifDiskTotalSpace = def.DefFunc("disk_total_space", 1, 1, []def.ArgInfo{{Name: "path"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	path := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDiskTotalSpace(executeData, returnValue, path)
})

// generate by ZifDiskFreeSpace
var DefZifDiskFreeSpace = def.DefFunc("disk_free_space", 1, 1, []def.ArgInfo{{Name: "path"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	path := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDiskFreeSpace(executeData, returnValue, path)
})

// generate by ZifDiskFreeSpace
var DefZifDiskfreespace = def.DefFunc("diskfreespace", 1, 1, []def.ArgInfo{{Name: "path"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	path := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDiskFreeSpace(executeData, returnValue, path)
})

// generate by ZifChgrp
var DefZifChgrp = def.DefFunc("chgrp", 2, 2, []def.ArgInfo{{Name: "filename"}, {Name: "group"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	filename := fp.ParseZval()
	group := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifChgrp(executeData, returnValue, filename, group)
})

// generate by ZifLchgrp
var DefZifLchgrp = def.DefFunc("lchgrp", 2, 2, []def.ArgInfo{{Name: "filename"}, {Name: "group"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	filename := fp.ParseZval()
	group := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLchgrp(executeData, returnValue, filename, group)
})

// generate by ZifChown
var DefZifChown = def.DefFunc("chown", 2, 2, []def.ArgInfo{{Name: "filename"}, {Name: "user"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	filename := fp.ParseZval()
	user := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifChown(executeData, returnValue, filename, user)
})

// generate by ZifLchown
var DefZifLchown = def.DefFunc("lchown", 2, 2, []def.ArgInfo{{Name: "filename"}, {Name: "user"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	filename := fp.ParseZval()
	user := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLchown(executeData, returnValue, filename, user)
})

// generate by ZifChmod
var DefZifChmod = def.DefFunc("chmod", 2, 2, []def.ArgInfo{{Name: "filename"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	filename := fp.ParseZval()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifChmod(executeData, returnValue, filename, mode)
})

// generate by ZifTouch
var DefZifTouch = def.DefFunc("touch", 1, 3, []def.ArgInfo{{Name: "filename"}, {Name: "time"}, {Name: "atime"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifClearstatcache = def.DefFunc("clearstatcache", 0, 2, []def.ArgInfo{{Name: "clear_realpath_cache"}, {Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	clear_realpath_cache := fp.ParseZval()
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifClearstatcache(executeData, returnValue, nil, clear_realpath_cache, filename)
})

// generate by ZifFileperms
var DefZifFileperms = def.DefFunc("fileperms", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFileperms(executeData, returnValue, filename)
})

// generate by ZifFileinode
var DefZifFileinode = def.DefFunc("fileinode", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFileinode(executeData, returnValue, filename)
})

// generate by ZifFilesize
var DefZifFilesize = def.DefFunc("filesize", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFilesize(executeData, returnValue, filename)
})

// generate by ZifFileowner
var DefZifFileowner = def.DefFunc("fileowner", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFileowner(executeData, returnValue, filename)
})

// generate by ZifFilegroup
var DefZifFilegroup = def.DefFunc("filegroup", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFilegroup(executeData, returnValue, filename)
})

// generate by ZifFileatime
var DefZifFileatime = def.DefFunc("fileatime", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFileatime(executeData, returnValue, filename)
})

// generate by ZifFilemtime
var DefZifFilemtime = def.DefFunc("filemtime", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFilemtime(executeData, returnValue, filename)
})

// generate by ZifFilectime
var DefZifFilectime = def.DefFunc("filectime", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFilectime(executeData, returnValue, filename)
})

// generate by ZifFiletype
var DefZifFiletype = def.DefFunc("filetype", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFiletype(executeData, returnValue, filename)
})

// generate by ZifIsWritable
var DefZifIsWritable = def.DefFunc("is_writable", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsWritable(executeData, returnValue, filename)
})

// generate by ZifIsWritable
var DefZifIsWriteable = def.DefFunc("is_writeable", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsWritable(executeData, returnValue, filename)
})

// generate by ZifIsReadable
var DefZifIsReadable = def.DefFunc("is_readable", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsReadable(executeData, returnValue, filename)
})

// generate by ZifIsExecutable
var DefZifIsExecutable = def.DefFunc("is_executable", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsExecutable(executeData, returnValue, filename)
})

// generate by ZifIsFile
var DefZifIsFile = def.DefFunc("is_file", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsFile(executeData, returnValue, filename)
})

// generate by ZifIsDir
var DefZifIsDir = def.DefFunc("is_dir", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsDir(executeData, returnValue, filename)
})

// generate by ZifIsLink
var DefZifIsLink = def.DefFunc("is_link", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsLink(executeData, returnValue, filename)
})

// generate by ZifFileExists
var DefZifFileExists = def.DefFunc("file_exists", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFileExists(executeData, returnValue, filename)
})

// generate by ZifLstat
var DefZifLstat = def.DefFunc("lstat", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifLstat(executeData, returnValue)
})

// generate by ZifStat
var DefZifStat = def.DefFunc("stat", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifStat(executeData, returnValue)
})

// generate by ZifRealpathCacheSize
var DefZifRealpathCacheSize = def.DefFunc("realpath_cache_size", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifRealpathCacheSize(executeData, returnValue)
})

// generate by ZifRealpathCacheGet
var DefZifRealpathCacheGet = def.DefFunc("realpath_cache_get", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifRealpathCacheGet(executeData, returnValue)
})
