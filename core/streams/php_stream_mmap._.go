// <<generate>>

package streams

// Source: <main/streams/php_stream_mmap.h>

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

type PhpStreamMmapOperationT = int

const (
	PHP_STREAM_MMAP_SUPPORTED = iota
	PHP_STREAM_MMAP_MAP_RANGE
	PHP_STREAM_MMAP_UNMAP
)

type PhpStreamMmapAccessT = int

const (
	PHP_STREAM_MAP_MODE_READONLY = iota
	PHP_STREAM_MAP_MODE_READWRITE
	PHP_STREAM_MAP_MODE_SHARED_READONLY
	PHP_STREAM_MAP_MODE_SHARED_READWRITE
)
const PHP_STREAM_MMAP_ALL = 0
const PHP_STREAM_MMAP_MAX = 512 * 1024 * 1024

/* Returns 1 if the stream in its current state can be memory mapped,
 * 0 otherwise */

/* un-maps the last mapped range */
