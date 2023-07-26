package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifFlock
var DefZifFlock = def.DefFunc("flock", 2, 3, []def.ArgInfo{{Name: "fp"}, {Name: "operation"}, {Name: "wouldblock"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	fp := fp.ParseZval()
	operation := fp.ParseZval()
	fp.StartOptional()
	wouldblock := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifFlock(executeData, returnValue, fp, operation, nil, wouldblock)
})

// generate by ZifGetMetaTags
var DefZifGetMetaTags = def.DefFunc("get_meta_tags", 1, 2, []def.ArgInfo{{Name: "filename"}, {Name: "use_include_path"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	use_include_path := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetMetaTags(executeData, returnValue, filename, nil, use_include_path)
})

// generate by ZifFileGetContents
var DefZifFileGetContents = def.DefFunc("file_get_contents", 1, 5, []def.ArgInfo{{Name: "filename"}, {Name: "flags"}, {Name: "context"}, {Name: "offset"}, {Name: "maxlen"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 5, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	context := fp.ParseZval()
	offset := fp.ParseZval()
	maxlen := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFileGetContents(executeData, returnValue, filename, nil, flags, context, offset, maxlen)
})

// generate by ZifFilePutContents
var DefZifFilePutContents = def.DefFunc("file_put_contents", 2, 4, []def.ArgInfo{{Name: "filename"}, {Name: "data"}, {Name: "flags"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	filename := fp.ParseZval()
	data := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFilePutContents(executeData, returnValue, filename, data, nil, flags, context)
})

// generate by ZifFile
var DefZifFile = def.DefFunc("file", 1, 3, []def.ArgInfo{{Name: "filename"}, {Name: "flags"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFile(executeData, returnValue, filename, nil, flags, context)
})

// generate by ZifTempnam
var DefZifTempnam = def.DefFunc("tempnam", 2, 2, []def.ArgInfo{{Name: "dir"}, {Name: "prefix"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	dir := fp.ParseZval()
	prefix := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTempnam(executeData, returnValue, dir, prefix)
})

// generate by ZifTmpfile
var DefZifTmpfile = def.DefFunc("tmpfile", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifTmpfile(executeData, returnValue)
})

// generate by ZifOpen
var DefZifOpen = def.DefFunc("open", 2, 4, []def.ArgInfo{{Name: "filename"}, {Name: "mode"}, {Name: "use_include_path"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	filename := fp.ParseZval()
	mode := fp.ParseZval()
	fp.StartOptional()
	use_include_path := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifOpen(executeData, returnValue, filename, mode, nil, use_include_path, context)
})

// generate by ZifFclose
var DefZifFclose = def.DefFunc("fclose", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFclose(executeData, returnValue, fp)
})

// generate by ZifPopen
var DefZifPopen = def.DefFunc("popen", 2, 2, []def.ArgInfo{{Name: "command"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	command := fp.ParseZval()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPopen(executeData, returnValue, command, mode)
})

// generate by ZifPclose
var DefZifPclose = def.DefFunc("pclose", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPclose(executeData, returnValue, fp)
})

// generate by ZifFeof
var DefZifFeof = def.DefFunc("feof", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFeof(executeData, returnValue, fp)
})

// generate by ZifFgets
var DefZifFgets = def.DefFunc("fgets", 1, 2, []def.ArgInfo{{Name: "fp"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	fp := fp.ParseResource()
	fp.StartOptional()
	length := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifFgets(fp, nil, length)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifFgetc
var DefZifFgetc = def.DefFunc("fgetc", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFgetc(executeData, returnValue, fp)
})

// generate by ZifFgetss
var DefZifFgetss = def.DefFunc("fgetss", 1, 3, []def.ArgInfo{{Name: "fp"}, {Name: "length"}, {Name: "allowable_tags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	fp := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseLongNullable()
	allowable_tags := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifFgetss(fp, nil, length, allowable_tags)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifFwrite
var DefZifFwrite = def.DefFunc("fwrite", 2, 3, []def.ArgInfo{{Name: "fp"}, {Name: "str"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	fp := fp.ParseZval()
	str := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFwrite(executeData, returnValue, fp, str, nil, length)
})

// generate by ZifFwrite
var DefZifFputs = def.DefFunc("fputs", 2, 3, []def.ArgInfo{{Name: "fp"}, {Name: "str"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	fp := fp.ParseZval()
	str := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFwrite(executeData, returnValue, fp, str, nil, length)
})

// generate by ZifFflush
var DefZifFflush = def.DefFunc("fflush", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFflush(executeData, returnValue, fp)
})

// generate by ZifRewind
var DefZifRewind = def.DefFunc("rewind", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRewind(executeData, returnValue, fp)
})

// generate by ZifFtell
var DefZifFtell = def.DefFunc("ftell", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFtell(executeData, returnValue, fp)
})

// generate by ZifFseek
var DefZifFseek = def.DefFunc("fseek", 2, 3, []def.ArgInfo{{Name: "fp"}, {Name: "offset"}, {Name: "whence"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	fp := fp.ParseZval()
	offset := fp.ParseZval()
	fp.StartOptional()
	whence := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFseek(executeData, returnValue, fp, offset, nil, whence)
})

// generate by ZifMkdir
var DefZifMkdir = def.DefFunc("mkdir", 1, 4, []def.ArgInfo{{Name: "pathname"}, {Name: "mode"}, {Name: "recursive"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	pathname := fp.ParseZval()
	fp.StartOptional()
	mode := fp.ParseZval()
	recursive := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMkdir(executeData, returnValue, pathname, nil, mode, recursive, context)
})

// generate by ZifRmdir
var DefZifRmdir = def.DefFunc("rmdir", 1, 2, []def.ArgInfo{{Name: "dirname"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	dirname := fp.ParseZval()
	fp.StartOptional()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRmdir(executeData, returnValue, dirname, nil, context)
})

// generate by ZifReadfile
var DefZifReadfile = def.DefFunc("readfile", 1, 3, []def.ArgInfo{{Name: "filename"}, {Name: "flags"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifReadfile(executeData, returnValue, filename, nil, flags, context)
})

// generate by ZifUmask
var DefZifUmask = def.DefFunc("umask", 0, 1, []def.ArgInfo{{Name: "mask"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	mask := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUmask(executeData, returnValue, nil, mask)
})

// generate by ZifFpassthru
var DefZifFpassthru = def.DefFunc("fpassthru", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFpassthru(executeData, returnValue, fp)
})

// generate by ZifRename
var DefZifRename = def.DefFunc("rename", 2, 3, []def.ArgInfo{{Name: "old_name"}, {Name: "new_name"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	old_name := fp.ParseZval()
	new_name := fp.ParseZval()
	fp.StartOptional()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRename(executeData, returnValue, old_name, new_name, nil, context)
})

// generate by ZifUnlink
var DefZifUnlink = def.DefFunc("unlink", 1, 2, []def.ArgInfo{{Name: "filename"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUnlink(executeData, returnValue, filename, nil, context)
})

// generate by ZifFtruncate
var DefZifFtruncate = def.DefFunc("ftruncate", 2, 2, []def.ArgInfo{{Name: "fp"}, {Name: "size"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	fp := fp.ParseZval()
	size := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFtruncate(executeData, returnValue, fp, size)
})

// generate by ZifFstat
var DefZifFstat = def.DefFunc("fstat", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFstat(executeData, returnValue, fp)
})

// generate by ZifCopy
var DefZifCopy = def.DefFunc("copy", 2, 3, []def.ArgInfo{{Name: "source_file"}, {Name: "destination_file"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	source_file := fp.ParseZval()
	destination_file := fp.ParseZval()
	fp.StartOptional()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCopy(executeData, returnValue, source_file, destination_file, nil, context)
})

// generate by ZifFread
var DefZifFread = def.DefFunc("fread", 2, 2, []def.ArgInfo{{Name: "fp"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	fp := fp.ParseZval()
	length := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFread(executeData, returnValue, fp, length)
})

// generate by ZifFputcsv
var DefZifFputcsv = def.DefFunc("fputcsv", 2, 5, []def.ArgInfo{{Name: "fp"}, {Name: "fields"}, {Name: "delimiter"}, {Name: "enclosure"}, {Name: "escape_char"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 5, 0)
	fp := fp.ParseZval()
	fields := fp.ParseZval()
	fp.StartOptional()
	delimiter := fp.ParseZval()
	enclosure := fp.ParseZval()
	escape_char := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFputcsv(executeData, returnValue, fp, fields, nil, delimiter, enclosure, escape_char)
})

// generate by ZifFgetcsv
var DefZifFgetcsv = def.DefFunc("fgetcsv", 1, 5, []def.ArgInfo{{Name: "fp"}, {Name: "length"}, {Name: "delimiter"}, {Name: "enclosure"}, {Name: "escape"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 5, 0)
	fp := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZval()
	delimiter := fp.ParseZval()
	enclosure := fp.ParseZval()
	escape := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFgetcsv(executeData, returnValue, fp, nil, length, delimiter, enclosure, escape)
})

// generate by ZifRealpath
var DefZifRealpath = def.DefFunc("realpath", 1, 1, []def.ArgInfo{{Name: "path"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	path := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRealpath(executeData, returnValue, path)
})

// generate by ZifFnmatch
var DefZifFnmatch = def.DefFunc("fnmatch", 2, 3, []def.ArgInfo{{Name: "pattern"}, {Name: "filename"}, {Name: "flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	pattern := fp.ParseZval()
	filename := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFnmatch(executeData, returnValue, pattern, filename, nil, flags)
})

// generate by ZifSysGetTempDir
var DefZifSysGetTempDir = def.DefFunc("sys_get_temp_dir", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifSysGetTempDir(executeData, returnValue)
})
