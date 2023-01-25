// <<generate>>

package core

import (
	"sik/core/streams"
	"sik/zend"
)

var PhpFileLeStream func() int
var PhpFileLePstream func() int
var PhpFileLeStreamFilter func() int

const PHP_STREAM_FLAG_NO_SEEK = 0x1
const PHP_STREAM_FLAG_NO_BUFFER = 0x2
const PHP_STREAM_FLAG_EOL_UNIX = 0x0
const PHP_STREAM_FLAG_DETECT_EOL = 0x4
const PHP_STREAM_FLAG_EOL_MAC = 0x8
const PHP_STREAM_FLAG_AVOID_BLOCKING = 0x10
const PHP_STREAM_FLAG_NO_CLOSE = 0x20
const PHP_STREAM_FLAG_IS_DIR = 0x40
const PHP_STREAM_FLAG_NO_FCLOSE = 0x80
const PHP_STREAM_FLAG_WAS_WRITTEN = 0x80000000
const PHP_STREAM_FCLOSE_NONE = 0
const PHP_STREAM_FCLOSE_FDOPEN = 1
const PHP_STREAM_FCLOSE_FOPENCOOKIE = 2

var _phpStreamAlloc func(ops *PhpStreamOps, abstract any, persistent_id *byte, mode *byte) *PhpStream
var PhpStreamEncloses func(enclosing *PhpStream, enclosed *PhpStream) *PhpStream
var _phpStreamFreeEnclosed func(stream_enclosed *PhpStream, close_options int) int
var PhpStreamFromPersistentId func(persistent_id *byte, stream **PhpStream) int

const PHP_STREAM_PERSISTENT_SUCCESS = 0
const PHP_STREAM_PERSISTENT_FAILURE = 1
const PHP_STREAM_PERSISTENT_NOT_EXIST = 2
const PHP_STREAM_FREE_CALL_DTOR = 1
const PHP_STREAM_FREE_RELEASE_STREAM = 2
const PHP_STREAM_FREE_PRESERVE_HANDLE = 4
const PHP_STREAM_FREE_RSRC_DTOR = 8
const PHP_STREAM_FREE_PERSISTENT = 16
const PHP_STREAM_FREE_IGNORE_ENCLOSING = 32
const PHP_STREAM_FREE_KEEP_RSRC = 64
const PHP_STREAM_FREE_CLOSE = PHP_STREAM_FREE_CALL_DTOR | PHP_STREAM_FREE_RELEASE_STREAM
const PHP_STREAM_FREE_CLOSE_CASTED = PHP_STREAM_FREE_CLOSE | PHP_STREAM_FREE_PRESERVE_HANDLE
const PHP_STREAM_FREE_CLOSE_PERSISTENT = PHP_STREAM_FREE_CLOSE | PHP_STREAM_FREE_PERSISTENT

var _phpStreamFree func(stream *PhpStream, close_options int) int
var _phpStreamSeek func(stream *PhpStream, offset zend.ZendOffT, whence int) int
var _phpStreamTell func(stream *PhpStream) zend.ZendOffT
var _phpStreamRead func(stream *PhpStream, buf *byte, count int) ssize_t
var PhpStreamReadToStr func(stream *PhpStream, len_ int) *zend.ZendString
var _phpStreamWrite func(stream *PhpStream, buf *byte, count int) ssize_t
var _phpStreamFillReadBuffer func(stream *PhpStream, size int) int
var _phpStreamPrintf func(stream *PhpStream, fmt *byte, _ ...any) ssize_t

const PhpStreamPrintf = _phpStreamPrintf

var _phpStreamEof func(stream *PhpStream) int
var _phpStreamGetc func(stream *PhpStream) int
var _phpStreamPutc func(stream *PhpStream, c int) int
var _phpStreamFlush func(stream *PhpStream, closing int) int
var _phpStreamGetLine func(stream *PhpStream, buf *byte, maxlen int, returned_len *int) *byte
var PhpStreamGetRecord func(stream *PhpStream, maxlen int, delim *byte, delim_len int) *zend.ZendString
var _phpStreamPuts func(stream *PhpStream, buf *byte) int
var _phpStreamStat func(stream *PhpStream, ssb *PhpStreamStatbuf) int
var _phpStreamStatPath func(path *byte, flags int, ssb *PhpStreamStatbuf, context *PhpStreamContext) int
var _phpStreamMkdir func(path *byte, mode int, options int, context *PhpStreamContext) int
var _phpStreamRmdir func(path *byte, options int, context *PhpStreamContext) int
var _phpStreamOpendir func(path *byte, options int, context *PhpStreamContext) *PhpStream
var _phpStreamReaddir func(dirstream *PhpStream, ent *PhpStreamDirent) *PhpStreamDirent
var PhpStreamDirentAlphasort func(a **zend.ZendString, b **zend.ZendString) int
var PhpStreamDirentAlphasortr func(a **zend.ZendString, b **zend.ZendString) int
var _phpStreamScandir func(dirname *byte, namelist []**zend.ZendString, flags int, context *PhpStreamContext, compare func(a **zend.ZendString, b **zend.ZendString) int) int
var _phpStreamSetOption func(stream *PhpStream, option int, value int, ptrparam any) int

const PHP_STREAM_MKDIR_RECURSIVE = 1
const PHP_STREAM_URL_STAT_LINK = 1
const PHP_STREAM_URL_STAT_QUIET = 2
const PHP_STREAM_URL_STAT_NOCACHE = 4
const PHP_STREAM_OPTION_BLOCKING = 1
const PHP_STREAM_OPTION_READ_BUFFER = 2
const PHP_STREAM_OPTION_WRITE_BUFFER = 3
const PHP_STREAM_BUFFER_NONE = 0
const PHP_STREAM_BUFFER_LINE = 1
const PHP_STREAM_BUFFER_FULL = 2
const PHP_STREAM_OPTION_READ_TIMEOUT = 4
const PHP_STREAM_OPTION_SET_CHUNK_SIZE = 5
const PHP_STREAM_OPTION_LOCKING = 6
const PHP_STREAM_LOCK_SUPPORTED = 1
const PHP_STREAM_OPTION_XPORT_API = 7
const PHP_STREAM_OPTION_CRYPTO_API = 8
const PHP_STREAM_OPTION_MMAP_API = 9
const PHP_STREAM_OPTION_TRUNCATE_API = 10
const PHP_STREAM_TRUNCATE_SUPPORTED = 0
const PHP_STREAM_TRUNCATE_SET_SIZE = 1

var _phpStreamTruncateSetSize func(stream *PhpStream, newsize int) int

const PHP_STREAM_OPTION_META_DATA_API = 11
const PHP_STREAM_OPTION_CHECK_LIVENESS = 12
const PHP_STREAM_OPTION_PIPE_BLOCKING = 13
const PHP_STREAM_OPTION_RETURN_OK = 0
const PHP_STREAM_OPTION_RETURN_ERR = -1
const PHP_STREAM_OPTION_RETURN_NOTIMPL = -2
const PHP_STREAM_COPY_ALL zend.ZendLong = size_t - 1

var _phpStreamCopyToStream func(src *PhpStream, dest *PhpStream, maxlen int) int
var _phpStreamCopyToStreamEx func(src *PhpStream, dest *PhpStream, maxlen int, len_ *int) int
var _phpStreamCopyToMem func(src *PhpStream, maxlen int, persistent int) *zend.ZendString
var _phpStreamPassthru func(src *PhpStream) ssize_t

const PHP_STREAM_AS_STDIO = 0
const PHP_STREAM_AS_FD = 1
const PHP_STREAM_AS_SOCKETD = 2
const PHP_STREAM_AS_FD_FOR_SELECT = 3
const PHP_STREAM_CAST_TRY_HARD = 0x80000000
const PHP_STREAM_CAST_RELEASE = 0x40000000
const PHP_STREAM_CAST_INTERNAL = 0x20000000
const PHP_STREAM_CAST_MASK = PHP_STREAM_CAST_TRY_HARD | PHP_STREAM_CAST_RELEASE | PHP_STREAM_CAST_INTERNAL

var _phpStreamCast func(stream *PhpStream, castas int, ret *any, show_err int) int

const PHP_STREAM_IS_STDIO *PhpStreamOps = &streams.PhpStreamStdioOps
const IGNORE_PATH = 0x0
const USE_PATH = 0x1
const IGNORE_URL = 0x2
const REPORT_ERRORS = 0x8
const STREAM_MUST_SEEK = 0x10
const STREAM_WILL_CAST = 0x20
const STREAM_LOCATE_WRAPPERS_ONLY = 0x40
const STREAM_OPEN_FOR_INCLUDE = 0x80
const STREAM_USE_URL = 0x100
const STREAM_ONLY_GET_HEADERS = 0x200
const STREAM_DISABLE_OPEN_BASEDIR = 0x400
const STREAM_OPEN_PERSISTENT = 0x800
const STREAM_USE_GLOB_DIR_OPEN = 0x1000
const STREAM_DISABLE_URL_PROTECTION = 0x2000
const STREAM_ASSUME_REALPATH = 0x4000
const STREAM_USE_BLOCKING_PIPE = 0x8000
const IGNORE_URL_WIN = 0

var PhpInitStreamWrappers func(module_number int) int
var PhpShutdownStreamWrappers func(module_number int) int
var PhpShutdownStreamHashes func()
var ZmDeactivateStreams func(type_ int, module_number int) int
var PhpRegisterUrlStreamWrapper func(protocol *byte, wrapper *PhpStreamWrapper) int
var PhpUnregisterUrlStreamWrapper func(protocol *byte) int
var PhpRegisterUrlStreamWrapperVolatile func(protocol *zend.ZendString, wrapper *PhpStreamWrapper) int
var PhpUnregisterUrlStreamWrapperVolatile func(protocol *zend.ZendString) int
var _phpStreamOpenWrapperEx func(path *byte, mode *byte, options int, opened_path **zend.ZendString, context *PhpStreamContext) *PhpStream
var PhpStreamLocateUrlWrapper func(path *byte, path_for_open **byte, options int) *PhpStreamWrapper
var PhpStreamLocateEol func(stream *PhpStream, buf *zend.ZendString) *byte
var PhpStreamWrapperLogError func(wrapper *PhpStreamWrapper, options int, fmt *byte, _ ...any)

const PHP_STREAM_UNCHANGED = 0
const PHP_STREAM_RELEASED = 1
const PHP_STREAM_FAILED = 2
const PHP_STREAM_CRITICAL = 3
const PHP_STREAM_NO_PREFERENCE = 0
const PHP_STREAM_PREFER_STDIO = 1
const PHP_STREAM_FORCE_CONVERSION = 2

var _phpStreamMakeSeekable func(origstream *PhpStream, newstream **PhpStream, flags int) int
var _phpStreamGetUrlStreamWrappersHash func() *zend.HashTable
var PhpStreamGetUrlStreamWrappersHashGlobal func() *zend.HashTable
var _phpGetStreamFiltersHash func() *zend.HashTable
var PhpGetStreamFiltersHashGlobal func() *zend.HashTable
var PhpStreamUserWrapperOps *PhpStreamWrapperOps

const PHP_STREAM_IS_URL = 1
const PHP_STREAM_META_TOUCH = 1
const PHP_STREAM_META_OWNER_NAME = 2
const PHP_STREAM_META_OWNER = 3
const PHP_STREAM_META_GROUP_NAME = 4
const PHP_STREAM_META_GROUP = 5
const PHP_STREAM_META_ACCESS = 6
