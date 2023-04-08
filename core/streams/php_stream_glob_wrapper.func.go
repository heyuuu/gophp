package streams

import (
	"github.com/heyuuu/gophp/core"
)

func PhpGlobStreamGetPath(stream *core.PhpStream, plen *int) *byte {
	return _phpGlobStreamGetPath(stream, plen)
}
func PhpGlobStreamGetCount(stream *core.PhpStream, pflags *int) int {
	return _phpGlobStreamGetCount(stream, pflags)
}
