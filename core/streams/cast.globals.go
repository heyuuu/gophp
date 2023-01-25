// <<generate>>

package streams

const PHP_FPOS_T = fpos_t
const HAVE_FOPENCOOKIE = 1
const PHP_EMULATE_FOPENCOOKIE = 1
const PHP_STREAM_COOKIE_FUNCTIONS *COOKIE_IO_FUNCTIONS_T = &StreamCookieFunctions

var StreamCookieFunctions COOKIE_IO_FUNCTIONS_T = COOKIE_IO_FUNCTIONS_T{StreamCookieReader, StreamCookieWriter, StreamCookieSeeker, StreamCookieCloser}
