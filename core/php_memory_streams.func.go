package core

import (
	"github.com/heyuuu/gophp/zend"
)

func PhpStreamMemoryCreate(mode int) *PhpStream    { return _phpStreamMemoryCreate(mode) }
func PhpStreamMemoryCreateRel(mode int) *PhpStream { return _phpStreamMemoryCreate(mode) }
func PhpStreamMemoryOpen(mode int, buf *byte, length int) *PhpStream {
	return _phpStreamMemoryOpen(mode, buf, length)
}
func PhpStreamMemoryGetBuffer(stream *PhpStream, length *int) *byte {
	return _phpStreamMemoryGetBuffer(stream, length)
}
func PhpStreamTempNew() *PhpStream {
	return PhpStreamTempCreate(TEMP_STREAM_DEFAULT, PHP_STREAM_MAX_MEM)
}
func PhpStreamTempCreate(mode int, max_memory_usage zend.ZendLong) *PhpStream {
	return _phpStreamTempCreate(mode, max_memory_usage)
}
func PhpStreamTempCreateEx(mode int, max_memory_usage int, tmpdir *byte) *PhpStream {
	return _phpStreamTempCreateEx(mode, max_memory_usage, tmpdir)
}
func PhpStreamTempCreateRel(mode int, max_memory_usage int) *PhpStream {
	return _phpStreamTempCreate(mode, max_memory_usage)
}
func PhpStreamTempOpen(mode int, max_memory_usage int, buf *byte, length int) *PhpStream {
	return _phpStreamTempOpen(mode, max_memory_usage, buf, length)
}
