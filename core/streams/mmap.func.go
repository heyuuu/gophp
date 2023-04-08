package streams

import (
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend"
)

func _phpStreamMmapUnmap(stream *core.PhpStream) int {
	return core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_MMAP_API, PHP_STREAM_MMAP_UNMAP, nil) == core.PHP_STREAM_OPTION_RETURN_OK
}
func _phpStreamMmapUnmapEx(stream *core.PhpStream, readden zend.ZendOffT) int {
	var ret int = 1
	if core.PhpStreamSeek(stream, readden, r.SEEK_CUR) != 0 {
		ret = 0
	}
	if PhpStreamMmapUnmap(stream) == 0 {
		ret = 0
	}
	return ret
}
