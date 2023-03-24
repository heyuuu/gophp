package streams

import (
	"sik/core"
)

func PhpGlobStreamGetPath(stream *core.PhpStream, plen *int) *byte {
	return _phpGlobStreamGetPath(stream, plen)
}
func PhpGlobStreamGetPattern(stream *core.PhpStream, plen *int) *byte {
	return _phpGlobStreamGetPattern(stream, plen)
}
func PhpGlobStreamGetCount(stream *core.PhpStream, pflags *int) int {
	return _phpGlobStreamGetCount(stream, pflags)
}
