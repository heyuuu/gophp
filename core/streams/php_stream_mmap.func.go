// <<generate>>

package streams

import (
	"sik/core"
	"sik/zend"
)

func PhpStreamMmapSupported(stream *core.PhpStream) int {
	if _phpStreamSetOption(stream, core.PHP_STREAM_OPTION_MMAP_API, PHP_STREAM_MMAP_SUPPORTED, nil) == 0 {
		return 1
	} else {
		return 0
	}
}
func PhpStreamMmapPossible(stream *core.PhpStream) bool {
	return !(PhpStreamIsFiltered(stream)) && PhpStreamMmapSupported(stream) != 0
}
func PhpStreamMmapUnmap(stream *core.PhpStream) int { return _phpStreamMmapUnmap(stream) }
func PhpStreamMmapUnmapEx(stream *core.PhpStream, readden zend.ZendOffT) int {
	return _phpStreamMmapUnmapEx(stream, readden)
}
