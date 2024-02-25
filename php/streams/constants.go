package streams

/* operations on streams that are file-handles */

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

/* state definitions when closing down; these are private to streams.c */
const PHP_STREAM_FCLOSE_NONE = 0
const PHP_STREAM_FCLOSE_FDOPEN = 1
const PHP_STREAM_FCLOSE_FOPENCOOKIE = 2

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

/* CAREFUL! this is equivalent to puts NOT fputs! */
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

const PHP_STREAM_COPY_ALL int = -1

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

/* pushes an error message onto the stack for a wrapper instance */

const PHP_STREAM_UNCHANGED = 0
const PHP_STREAM_RELEASED = 1
const PHP_STREAM_FAILED = 2
const PHP_STREAM_CRITICAL = 3
const PHP_STREAM_NO_PREFERENCE = 0
const PHP_STREAM_PREFER_STDIO = 1
const PHP_STREAM_FORCE_CONVERSION = 2

const PHP_STREAM_IS_URL = 1

/* Stream metadata definitions */

const PHP_STREAM_META_TOUCH = 1
const PHP_STREAM_META_OWNER_NAME = 2
const PHP_STREAM_META_OWNER = 3
const PHP_STREAM_META_GROUP_NAME = 4
const PHP_STREAM_META_GROUP = 5
const PHP_STREAM_META_ACCESS = 6
