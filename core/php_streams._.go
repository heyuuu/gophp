package core

import (
	"sik/core/streams"
	"sik/zend"
	"sik/zend/types"
)

// Source: <main/php_streams.h>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Author: Wez Furlong (wez@thebrainroom.com)                           |
   +----------------------------------------------------------------------+
*/

var PhpFileLeStream func() int
var PhpFileLePstream func() int
var PhpFileLeStreamFilter func() int

/* {{{ Streams memory debugging stuff */

/* these functions relay the file/line number information. They are depth aware, so they will pass
 * the ultimate ancestor, which is useful, because there can be several layers of calls */

type PhpStreamContext = streams.PhpStreamContext
type PhpStreamFilter = streams.PhpStreamFilter

/* operations on streams that are file-handles */

const PHP_STREAM_FLAG_NO_SEEK = 0x1
const PHP_STREAM_FLAG_NO_BUFFER = 0x2
const PHP_STREAM_FLAG_EOL_UNIX = 0x0
const PHP_STREAM_FLAG_DETECT_EOL = 0x4
const PHP_STREAM_FLAG_EOL_MAC = 0x8

/* set this when the stream might represent "interactive" data.
 * When set, the read buffer will avoid certain operations that
 * might otherwise cause the read to block for much longer than
 * is strictly required. */

const PHP_STREAM_FLAG_AVOID_BLOCKING = 0x10
const PHP_STREAM_FLAG_NO_CLOSE = 0x20
const PHP_STREAM_FLAG_IS_DIR = 0x40
const PHP_STREAM_FLAG_NO_FCLOSE = 0x80
const PHP_STREAM_FLAG_WAS_WRITTEN = 0x80000000

/* state definitions when closing down; these are private to streams.c */

const PHP_STREAM_FCLOSE_NONE = 0
const PHP_STREAM_FCLOSE_FDOPEN = 1
const PHP_STREAM_FCLOSE_FOPENCOOKIE = 2

/* allocate a new stream for a particular ops */

var _phpStreamAlloc func(ops *PhpStreamOps, abstract any, persistent_id *byte, mode *byte) *PhpStream

/* use this to tell the stream that it is OK if we don't explicitly close it */

/* use this to assign the stream to a zval and tell the stream that is
 * has been exported to the engine; it will expect to be closed automatically
 * when the resources are auto-destructed */

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
var PhpStreamReadToStr func(stream *PhpStream, len_ int) *types.String
var _phpStreamWrite func(stream *PhpStream, buf *byte, count int) ssize_t
var _phpStreamFillReadBuffer func(stream *PhpStream, size int) int
var _phpStreamPrintf func(stream *PhpStream, fmt *byte, _ ...any) ssize_t

/* php_stream_printf macro & function require */

const PhpStreamPrintf = _phpStreamPrintf

var _phpStreamEof func(stream *PhpStream) int
var _phpStreamGetc func(stream *PhpStream) int
var _phpStreamPutc func(stream *PhpStream, c int) int
var _phpStreamFlush func(stream *PhpStream, closing int) int
var _phpStreamGetLine func(stream *PhpStream, buf *byte, maxlen int, returned_len *int) *byte
var PhpStreamGetRecord func(stream *PhpStream, maxlen int, delim *byte, delim_len int) *types.String

/* CAREFUL! this is equivalent to puts NOT fputs! */

var _phpStreamPuts func(stream *PhpStream, buf *byte) int
var _phpStreamStat func(stream *PhpStream, ssb *PhpStreamStatbuf) int
var _phpStreamStatPath func(path *byte, flags int, ssb *PhpStreamStatbuf, context *PhpStreamContext) int
var _phpStreamMkdir func(path *byte, mode int, options int, context *PhpStreamContext) int
var _phpStreamRmdir func(path *byte, options int, context *PhpStreamContext) int
var _phpStreamOpendir func(path *byte, options int, context *PhpStreamContext) *PhpStream
var _phpStreamReaddir func(dirstream *PhpStream, ent *PhpStreamDirent) *PhpStreamDirent
var PhpStreamDirentAlphasort func(a **types.String, b **types.String) int
var PhpStreamDirentAlphasortr func(a **types.String, b **types.String) int
var _phpStreamScandir func(dirname *byte, namelist []**types.String, flags int, context *PhpStreamContext, compare func(a **types.String, b **types.String) int) int
var _phpStreamSetOption func(stream *PhpStream, option int, value int, ptrparam any) int

/* Flags for mkdir method in wrapper ops */

const PHP_STREAM_MKDIR_RECURSIVE = 1

/* define REPORT __special__  ERRORS 8 (below) */

const PHP_STREAM_URL_STAT_LINK = 1
const PHP_STREAM_URL_STAT_QUIET = 2
const PHP_STREAM_URL_STAT_NOCACHE = 4

/* change the blocking mode of stream: value == 1 => blocking, value == 0 => non-blocking. */

const PHP_STREAM_OPTION_BLOCKING = 1

/* change the buffering mode of stream. value is a PHP_STREAM_BUFFER_XXXX value, ptrparam is a ptr to a size_t holding
 * the required buffer size */

const PHP_STREAM_OPTION_READ_BUFFER = 2
const PHP_STREAM_OPTION_WRITE_BUFFER = 3
const PHP_STREAM_BUFFER_NONE = 0
const PHP_STREAM_BUFFER_LINE = 1
const PHP_STREAM_BUFFER_FULL = 2

/* set the timeout duration for reads on the stream. ptrparam is a pointer to a struct timeval * */

const PHP_STREAM_OPTION_READ_TIMEOUT = 4
const PHP_STREAM_OPTION_SET_CHUNK_SIZE = 5

/* set or release lock on a stream */

const PHP_STREAM_OPTION_LOCKING = 6

/* whether or not locking is supported */

const PHP_STREAM_LOCK_SUPPORTED = 1

/* option code used by the php_stream_xport_XXX api */

const PHP_STREAM_OPTION_XPORT_API = 7
const PHP_STREAM_OPTION_CRYPTO_API = 8
const PHP_STREAM_OPTION_MMAP_API = 9
const PHP_STREAM_OPTION_TRUNCATE_API = 10
const PHP_STREAM_TRUNCATE_SUPPORTED = 0
const PHP_STREAM_TRUNCATE_SET_SIZE = 1

var _phpStreamTruncateSetSize func(stream *PhpStream, newsize int) int

const PHP_STREAM_OPTION_META_DATA_API = 11

/* Check if the stream is still "live"; for sockets/pipes this means the socket
 * is still connected; for files, this does not really have meaning */

const PHP_STREAM_OPTION_CHECK_LIVENESS = 12

/* Enable/disable blocking reads on anonymous pipes on Windows. */

const PHP_STREAM_OPTION_PIPE_BLOCKING = 13
const PHP_STREAM_OPTION_RETURN_OK = 0
const PHP_STREAM_OPTION_RETURN_ERR = -1
const PHP_STREAM_OPTION_RETURN_NOTIMPL = -2

/* copy up to maxlen bytes from src to dest.  If maxlen is PHP_STREAM_COPY_ALL,
 * copy until eof(src). */

const PHP_STREAM_COPY_ALL zend.ZendLong = size_t - 1

var _phpStreamCopyToStream func(src *PhpStream, dest *PhpStream, maxlen int) int
var _phpStreamCopyToStreamEx func(src *PhpStream, dest *PhpStream, maxlen int, len_ *int) int

/* read all data from stream and put into a buffer. Caller must free buffer
 * when done. */

var _phpStreamCopyToMem func(src *PhpStream, maxlen int, persistent int) *types.String

/* output all data from a stream */

var _phpStreamPassthru func(src *PhpStream) ssize_t

/* coerce the stream into some other form */

const PHP_STREAM_AS_STDIO = 0

/* cast as a POSIX fd or socketd */

const PHP_STREAM_AS_FD = 1

/* cast as a socketd */

const PHP_STREAM_AS_SOCKETD = 2

/* cast as fd/socket for select purposes */

const PHP_STREAM_AS_FD_FOR_SELECT = 3

/* try really, really hard to make sure the cast happens (avoid using this flag if possible) */

const PHP_STREAM_CAST_TRY_HARD = 0x80000000
const PHP_STREAM_CAST_RELEASE = 0x40000000
const PHP_STREAM_CAST_INTERNAL = 0x20000000
const PHP_STREAM_CAST_MASK = PHP_STREAM_CAST_TRY_HARD | PHP_STREAM_CAST_RELEASE | PHP_STREAM_CAST_INTERNAL

var _phpStreamCast func(stream *PhpStream, castas int, ret *any, show_err int) int

/* use this to check if a stream can be cast into another form */

/* use this to check if a stream is of a particular type:
 * PHPAPI int php_stream_is(php_stream *stream, php_stream_ops *ops); */

const PHP_STREAM_IS_STDIO *PhpStreamOps = &streams.PhpStreamStdioOps

/* Wrappers support */

const IGNORE_PATH = 0x0
const USE_PATH = 0x1
const IGNORE_URL = 0x2
const REPORT_ERRORS = 0x8

/* If you don't need to write to the stream, but really need to
 * be able to seek, use this flag in your options. */

const STREAM_MUST_SEEK = 0x10

/* If you are going to end up casting the stream into a FILE* or
 * a socket, pass this flag and the streams/wrappers will not use
 * buffering mechanisms while reading the headers, so that HTTP
 * wrapped streams will work consistently.
 * If you omit this flag, streams will use buffering and should end
 * up working more optimally.
 * */

const STREAM_WILL_CAST = 0x20

/* this flag applies to php_stream_locate_url_wrapper */

const STREAM_LOCATE_WRAPPERS_ONLY = 0x40

/* this flag is only used by include/require functions */

const STREAM_OPEN_FOR_INCLUDE = 0x80

/* this flag tells streams to ONLY open urls */

const STREAM_USE_URL = 0x100

/* this flag is used when only the headers from HTTP request are to be fetched */

const STREAM_ONLY_GET_HEADERS = 0x200

/* don't apply open_basedir checks */

const STREAM_DISABLE_OPEN_BASEDIR = 0x400

/* get (or create) a persistent version of the stream */

const STREAM_OPEN_PERSISTENT = 0x800

/* use glob stream for directory open in plain files stream */

const STREAM_USE_GLOB_DIR_OPEN = 0x1000

/* don't check allow_url_fopen and allow_url_include */

const STREAM_DISABLE_URL_PROTECTION = 0x2000

/* assume the path passed in exists and is fully expanded, avoiding syscalls */

const STREAM_ASSUME_REALPATH = 0x4000

/* Allow blocking reads on anonymous pipes on Windows. */

const STREAM_USE_BLOCKING_PIPE = 0x8000

/* Antique - no longer has meaning */

const IGNORE_URL_WIN = 0

var PhpInitStreamWrappers func(module_number int) int
var PhpShutdownStreamWrappers func(module_number int) int
var PhpShutdownStreamHashes func()
var ZmDeactivateStreams func(type_ int, module_number int) int
var PhpRegisterUrlStreamWrapper func(protocol *byte, wrapper *PhpStreamWrapper) int
var PhpUnregisterUrlStreamWrapper func(protocol *byte) int
var PhpRegisterUrlStreamWrapperVolatile func(protocol *types.String, wrapper *PhpStreamWrapper) int
var PhpUnregisterUrlStreamWrapperVolatile func(protocol *types.String) int
var _phpStreamOpenWrapperEx func(path *byte, mode *byte, options int, opened_path **types.String, context *PhpStreamContext) *PhpStream
var PhpStreamLocateUrlWrapper func(path *byte, path_for_open **byte, options int) *PhpStreamWrapper
var PhpStreamLocateEol func(stream *PhpStream, buf *types.String) *byte

/* pushes an error message onto the stack for a wrapper instance */

var PhpStreamWrapperLogError func(wrapper *PhpStreamWrapper, options int, fmt *byte, _ ...any)

const PHP_STREAM_UNCHANGED = 0
const PHP_STREAM_RELEASED = 1
const PHP_STREAM_FAILED = 2
const PHP_STREAM_CRITICAL = 3
const PHP_STREAM_NO_PREFERENCE = 0
const PHP_STREAM_PREFER_STDIO = 1
const PHP_STREAM_FORCE_CONVERSION = 2

/* DO NOT call this on streams that are referenced by resources! */

var _phpStreamMakeSeekable func(origstream *PhpStream, newstream **PhpStream, flags int) int

/* Give other modules access to the url_stream_wrappers_hash and stream_filters_hash */

var _phpStreamGetUrlStreamWrappersHash func() *types.Array
var PhpStreamGetUrlStreamWrappersHashGlobal func() *types.Array
var _phpGetStreamFiltersHash func() *types.Array
var PhpGetStreamFiltersHashGlobal func() *types.Array
var PhpStreamUserWrapperOps *PhpStreamWrapperOps

/* Definitions for user streams */

const PHP_STREAM_IS_URL = 1

/* Stream metadata definitions */

const PHP_STREAM_META_TOUCH = 1
const PHP_STREAM_META_OWNER_NAME = 2
const PHP_STREAM_META_OWNER = 3
const PHP_STREAM_META_GROUP_NAME = 4
const PHP_STREAM_META_GROUP = 5
const PHP_STREAM_META_ACCESS = 6
