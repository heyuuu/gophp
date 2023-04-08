package streams

import (
	"github.com/heyuuu/gophp/core"
)

var PhpGetUidByName func(name *byte, uid *uid_t) int
var PhpGetGidByName func(name *byte, gid *gid_t) int

var PhpStreamStdioOps core.PhpStreamOps = core.MakePhpStreamOps(PhpStdiopWrite, PhpStdiopRead, PhpStdiopClose, PhpStdiopFlush, "STDIO", PhpStdiopSeek, PhpStdiopCast, PhpStdiopStat, PhpStdiopSetOption)
var PhpPlainFilesDirstreamOps core.PhpStreamOps = core.MakePhpStreamOps(nil, PhpPlainFilesDirstreamRead, PhpPlainFilesDirstreamClose, nil, "dir", PhpPlainFilesDirstreamRewind, nil, nil, nil)
var PhpPlainFilesWrapperOps core.PhpStreamWrapperOps = core.MakePhpStreamWrapperOps(PhpPlainFilesStreamOpener, nil, nil, PhpPlainFilesUrlStater, PhpPlainFilesDirOpener, "plainfile", PhpPlainFilesUnlink, PhpPlainFilesRename, PhpPlainFilesMkdir, PhpPlainFilesRmdir, PhpPlainFilesMetadata)

/* TODO: We have to make php_plain_files_wrapper writable to support SWOOLE */

var PhpPlainFilesWrapper core.PhpStreamWrapper = core.MakePhpStreamWrapper(&PhpPlainFilesWrapperOps, nil, 0)
