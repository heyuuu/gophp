package streams

import (
	"sik/zend/types"
)

/* {{{ resource and registration code */

var UrlStreamWrappersHash types.Array
var LeStream int = types.FAILURE
var LePstream int = types.FAILURE
var LeStreamFilter int = types.FAILURE

/* {{{ wrapper error reporting */

/* Like php_stream_read(), but reading into a zend_string buffer. This has some similarity
 * to the copy_to_mem() operation, but only performs a single direct read. */

/* If buf == NULL, the buffer will be allocated automatically and will be of an
 * appropriate length to hold the line, regardless of the line length, memory
 * permitting */

/* Writes a buffer directly to a stream, using multiple of the chunk size */

/* push some data through the write filter chain.
 * buf may be NULL, if flags are set to indicate a flush.
 * This may trigger a real write to the stream.
 * Returns the number of bytes consumed from buf by the first filter in the chain.
 * */

/* Returns SUCCESS/FAILURE and sets *len to the number of bytes moved */

/* Returns the number of bytes moved.
 * Returns 1 when source len is 0.
 * Deprecated in favor of php_stream_copy_to_stream_ex() */

/* Validate protocol scheme names during registration
 * Must conform to /^[a-zA-Z0-9+.-]+$/
 */

/* API for registering GLOBAL wrappers */

/* API for registering VOLATILE wrappers */
