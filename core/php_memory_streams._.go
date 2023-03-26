package core

import (
	"sik/zend"
)

const PHP_STREAM_MAX_MEM zend.ZendLong = 2 * 1024 * 1024
const TEMP_STREAM_DEFAULT = 0x0
const TEMP_STREAM_READONLY = 0x1
const TEMP_STREAM_TAKE_BUFFER = 0x2
const TEMP_STREAM_APPEND = 0x4

var _phpStreamMemoryCreate func(mode int) *PhpStream
var _phpStreamMemoryOpen func(mode int, buf *byte, length int) *PhpStream
var _phpStreamMemoryGetBuffer func(stream *PhpStream, length *int) *byte
var _phpStreamTempCreate func(mode int, max_memory_usage int) *PhpStream
var _phpStreamTempCreateEx func(mode int, max_memory_usage int, tmpdir *byte) *PhpStream
var _phpStreamTempOpen func(mode int, max_memory_usage int, buf *byte, length int) *PhpStream
var PhpStreamModeFromStr func(mode *byte) int
var _phpStreamModeToStr func(mode int) *byte
var PhpStreamMemoryOps PhpStreamOps
var PhpStreamTempOps PhpStreamOps
var PhpStreamRfc2397Ops PhpStreamOps
var PhpStreamRfc2397Wrapper PhpStreamWrapper

const PHP_STREAM_IS_MEMORY *PhpStreamOps = &PhpStreamMemoryOps
const PHP_STREAM_IS_TEMP = &PhpStreamTempOps
