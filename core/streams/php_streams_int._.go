package streams

import (
	"github.com/heyuuu/gophp/core"
)

const STREAM_DEBUG = 0
const STREAM_WRAPPER_PLAIN_FILES = (*core.PhpStreamWrapper)(-1)
const MAP_FAILED = any(-1)
const CHUNK_SIZE = 8192

/* This functions transforms the first char to 'w' if it's not 'r', 'a' or 'w'
 * and strips any subsequent chars except '+' and 'b'.
 * Use this to sanitize stream->mode if you call e.g. fdopen, fopencookie or
 * any other function that expects standard modes and you allow non-standard
 * ones. result should be a char[5]. */
