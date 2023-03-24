package standard

import (
	"sik/zend/zpp"
	"sik/zend/def"
	"sik/zend/types"
)

// generate by ZifFlock
var DefZifFlock = def.DefFunc("flock", 2, 3, []def.ArgInfo{{name: "fp"}, {name: "operation"}, {name: "wouldblock"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifGetMetaTags = def.DefFunc("get_meta_tags", 1, 2, []def.ArgInfo{{name: "filename"}, {name: "use_include_path"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifFileGetContents = def.DefFunc("file_get_contents", 1, 5, []def.ArgInfo{{name: "filename"}, {name: "flags"}, {name: "context"}, {name: "offset"}, {name: "maxlen"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifFilePutContents = def.DefFunc("file_put_contents", 2, 4, []def.ArgInfo{{name: "filename"}, {name: "data"}, {name: "flags"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifFile = def.DefFunc("file", 1, 3, []def.ArgInfo{{name: "filename"}, {name: "flags"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifFgets
var DefZifFgets = def.DefFunc("fgets", 1, 2, []def.ArgInfo{{name: "fp"}, {name: "length"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	fp := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFgets(executeData, returnValue, fp, nil, length)
})

// generate by ZifFgetss
var DefZifFgetss = def.DefFunc("fgetss", 1, 3, []def.ArgInfo{{name: "fp"}, {name: "length"}, {name: "allowable_tags"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	fp := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZval()
	allowable_tags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFgetss(executeData, returnValue, fp, nil, length, allowable_tags)
})

// generate by ZifFwrite
var DefZifFwrite = def.DefFunc("fwrite", 2, 3, []def.ArgInfo{{name: "fp"}, {name: "str"}, {name: "length"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifFseek
var DefZifFseek = def.DefFunc("fseek", 2, 3, []def.ArgInfo{{name: "fp"}, {name: "offset"}, {name: "whence"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifMkdir = def.DefFunc("mkdir", 1, 4, []def.ArgInfo{{name: "pathname"}, {name: "mode"}, {name: "recursive"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifRmdir = def.DefFunc("rmdir", 1, 2, []def.ArgInfo{{name: "dirname"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifReadfile = def.DefFunc("readfile", 1, 3, []def.ArgInfo{{name: "filename"}, {name: "flags"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifUmask = def.DefFunc("umask", 0, 1, []def.ArgInfo{{name: "mask"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	mask := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUmask(executeData, returnValue, nil, mask)
})

// generate by ZifRename
var DefZifRename = def.DefFunc("rename", 2, 3, []def.ArgInfo{{name: "old_name"}, {name: "new_name"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifUnlink = def.DefFunc("unlink", 1, 2, []def.ArgInfo{{name: "filename"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUnlink(executeData, returnValue, filename, nil, context)
})

// generate by ZifCopy
var DefZifCopy = def.DefFunc("copy", 2, 3, []def.ArgInfo{{name: "source_file"}, {name: "destination_file"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifFputcsv
var DefZifFputcsv = def.DefFunc("fputcsv", 2, 5, []def.ArgInfo{{name: "fp"}, {name: "fields"}, {name: "delimiter"}, {name: "enclosure"}, {name: "escape_char"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
var DefZifFgetcsv = def.DefFunc("fgetcsv", 1, 5, []def.ArgInfo{{name: "fp"}, {name: "length"}, {name: "delimiter"}, {name: "enclosure"}, {name: "escape"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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

// generate by ZifFnmatch
var DefZifFnmatch = def.DefFunc("fnmatch", 2, 3, []def.ArgInfo{{name: "pattern"}, {name: "filename"}, {name: "flags"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
