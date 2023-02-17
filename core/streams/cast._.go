package streams

const PHP_FPOS_T = fpos_t
const PHP_STREAM_COOKIE_FUNCTIONS *COOKIE_IO_FUNCTIONS_T = &StreamCookieFunctions

/* {{{ STDIO with fopencookie */

/* use our fopencookie emulation */

var StreamCookieFunctions COOKIE_IO_FUNCTIONS_T = MakeCOOKIE_IO_FUNCTIONS_T(StreamCookieReader, StreamCookieWriter, StreamCookieSeeker, StreamCookieCloser)
