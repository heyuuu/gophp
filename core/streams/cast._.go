package streams

type PHP_FPOS_T = int64

var PHP_STREAM_COOKIE_FUNCTIONS *COOKIE_IO_FUNCTIONS_T = &StreamCookieFunctions
var StreamCookieFunctions COOKIE_IO_FUNCTIONS_T = COOKIE_IO_FUNCTIONS_T{
	reader: StreamCookieReader,
	writer: StreamCookieWriter,
	seeker: StreamCookieSeeker,
	closer: StreamCookieCloser,
}
