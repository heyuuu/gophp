// <<generate>>

package streams

import (
	"sik/core"
)

var LeProtocols int
var UserStreamWops core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{UserWrapperOpener, nil, nil, UserWrapperStatUrl, UserWrapperOpendir, "user-space", UserWrapperUnlink, UserWrapperRename, UserWrapperMkdir, UserWrapperRmdir, UserWrapperMetadata}

type PhpUserstreamDataT = _phpUserstreamData

const USERSTREAM_OPEN = "stream_open"
const USERSTREAM_CLOSE = "stream_close"
const USERSTREAM_READ = "stream_read"
const USERSTREAM_WRITE = "stream_write"
const USERSTREAM_FLUSH = "stream_flush"
const USERSTREAM_SEEK = "stream_seek"
const USERSTREAM_TELL = "stream_tell"
const USERSTREAM_EOF = "stream_eof"
const USERSTREAM_STAT = "stream_stat"
const USERSTREAM_STATURL = "url_stat"
const USERSTREAM_UNLINK = "unlink"
const USERSTREAM_RENAME = "rename"
const USERSTREAM_MKDIR = "mkdir"
const USERSTREAM_RMDIR = "rmdir"
const USERSTREAM_DIR_OPEN = "dir_opendir"
const USERSTREAM_DIR_READ = "dir_readdir"
const USERSTREAM_DIR_REWIND = "dir_rewinddir"
const USERSTREAM_DIR_CLOSE = "dir_closedir"
const USERSTREAM_LOCK = "stream_lock"
const USERSTREAM_CAST = "stream_cast"
const USERSTREAM_SET_OPTION = "stream_set_option"
const USERSTREAM_TRUNCATE = "stream_truncate"
const USERSTREAM_METADATA = "stream_metadata"

var PhpStreamUserspaceOps core.PhpStreamOps = core.PhpStreamOps{PhpUserstreamopWrite, PhpUserstreamopRead, PhpUserstreamopClose, PhpUserstreamopFlush, "user-space", PhpUserstreamopSeek, PhpUserstreamopCast, PhpUserstreamopStat, PhpUserstreamopSetOption}
var PhpStreamUserspaceDirOps core.PhpStreamOps = core.PhpStreamOps{nil, PhpUserstreamopReaddir, PhpUserstreamopClosedir, nil, "user-space-dir", PhpUserstreamopRewinddir, nil, nil, nil}
