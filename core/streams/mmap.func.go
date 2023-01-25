// <<generate>>

package streams

import (
	"sik/core"
	r "sik/runtime"
	"sik/zend"
)

func _phpStreamMmapRange(stream *core.PhpStream, offset int, length int, mode PhpStreamMmapAccessT, mapped_len *int) *byte {
	var range_ PhpStreamMmapRange
	range_.SetOffset(offset)
	range_.SetLength(length)
	range_.SetMode(mode)
	range_.SetMapped(nil)
	if core.PHP_STREAM_OPTION_RETURN_OK == core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_MMAP_API, PHP_STREAM_MMAP_MAP_RANGE, &range_) {
		if mapped_len != nil {
			*mapped_len = range_.GetLength()
		}
		return range_.GetMapped()
	}
	return nil
}
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
