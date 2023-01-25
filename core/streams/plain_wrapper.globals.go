// <<generate>>

package streams

import (
	"sik/core"
)

var PhpGetUidByName func(name *byte, uid *uid_t) int
var PhpGetGidByName func(name *byte, gid *gid_t) int
var PhpStreamStdioOps core.PhpStreamOps = core.PhpStreamOps{PhpStdiopWrite, PhpStdiopRead, PhpStdiopClose, PhpStdiopFlush, "STDIO", PhpStdiopSeek, PhpStdiopCast, PhpStdiopStat, PhpStdiopSetOption}
var PhpPlainFilesDirstreamOps core.PhpStreamOps = core.PhpStreamOps{nil, PhpPlainFilesDirstreamRead, PhpPlainFilesDirstreamClose, nil, "dir", PhpPlainFilesDirstreamRewind, nil, nil, nil}
var PhpPlainFilesWrapperOps core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{PhpPlainFilesStreamOpener, nil, nil, PhpPlainFilesUrlStater, PhpPlainFilesDirOpener, "plainfile", PhpPlainFilesUnlink, PhpPlainFilesRename, PhpPlainFilesMkdir, PhpPlainFilesRmdir, PhpPlainFilesMetadata}
var PhpPlainFilesWrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&PhpPlainFilesWrapperOps, nil, 0}
