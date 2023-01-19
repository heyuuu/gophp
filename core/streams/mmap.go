// <<generate>>

package streams

import (
	"sik/core"
	"sik/zend"
)

// Source: <main/streams/mmap.c>

/*
  +----------------------------------------------------------------------+
  | PHP Version 7                                                        |
  +----------------------------------------------------------------------+
  | Copyright (c) The PHP Group                                          |
  +----------------------------------------------------------------------+
  | This source file is subject to version 3.01 of the PHP license,      |
  | that is bundled with this package in the file LICENSE, and is        |
  | available through the world-wide-web at the following url:           |
  | http://www.php.net/license/3_01.txt                                  |
  | If you did not receive a copy of the PHP license and are unable to   |
  | obtain it through the world-wide-web, please send a note to          |
  | license@php.net so we can mail you a copy immediately.               |
  +----------------------------------------------------------------------+
  | Author: Wez Furlong <wez@thebrainroom.com>                           |
  +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_streams_int.h"

func _phpStreamMmapRange(stream *core.PhpStream, offset int, length int, mode PhpStreamMmapAccessT, mapped_len *int) *byte {
	var range_ PhpStreamMmapRange
	range_.SetOffset(offset)
	range_.SetLength(length)
	range_.SetMode(mode)
	range_.SetMapped(nil)
	if 0 == _phpStreamSetOption(stream, 9, PHP_STREAM_MMAP_MAP_RANGE, &range_) {
		if mapped_len != nil {
			*mapped_len = range_.GetLength()
		}
		return range_.GetMapped()
	}
	return nil
}
func _phpStreamMmapUnmap(stream *core.PhpStream) int {
	return _phpStreamSetOption(stream, 9, PHP_STREAM_MMAP_UNMAP, nil) == 0
}
func _phpStreamMmapUnmapEx(stream *core.PhpStream, readden zend.ZendOffT) int {
	var ret int = 1
	if _phpStreamSeek(stream, readden, SEEK_CUR) != 0 {
		ret = 0
	}
	if _phpStreamMmapUnmap(stream) == 0 {
		ret = 0
	}
	return ret
}
