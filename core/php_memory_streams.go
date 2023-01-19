// <<generate>>

package core

// Source: <main/php_memory_streams.h>

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
   | Author: Marcus Boerger <helly@php.net>                               |
   +----------------------------------------------------------------------+
*/

// #define PHP_MEMORY_STREAM_H

// # include "php_streams.h"

// #define PHP_STREAM_MAX_MEM       2 * 1024 * 1024

// #define TEMP_STREAM_DEFAULT       0x0

// #define TEMP_STREAM_READONLY       0x1

// #define TEMP_STREAM_TAKE_BUFFER       0x2

// #define TEMP_STREAM_APPEND       0x4

// #define php_stream_memory_create(mode) _php_stream_memory_create ( ( mode ) STREAMS_CC )

// #define php_stream_memory_create_rel(mode) _php_stream_memory_create ( ( mode ) STREAMS_REL_CC )

// #define php_stream_memory_open(mode,buf,length) _php_stream_memory_open ( ( mode ) , ( buf ) , ( length ) STREAMS_CC )

// #define php_stream_memory_get_buffer(stream,length) _php_stream_memory_get_buffer ( ( stream ) , ( length ) STREAMS_CC )

// #define php_stream_temp_new() php_stream_temp_create ( TEMP_STREAM_DEFAULT , PHP_STREAM_MAX_MEM )

// #define php_stream_temp_create(mode,max_memory_usage) _php_stream_temp_create ( ( mode ) , ( max_memory_usage ) STREAMS_CC )

// #define php_stream_temp_create_ex(mode,max_memory_usage,tmpdir) _php_stream_temp_create_ex ( ( mode ) , ( max_memory_usage ) , ( tmpdir ) STREAMS_CC )

// #define php_stream_temp_create_rel(mode,max_memory_usage) _php_stream_temp_create ( ( mode ) , ( max_memory_usage ) STREAMS_REL_CC )

// #define php_stream_temp_open(mode,max_memory_usage,buf,length) _php_stream_temp_open ( ( mode ) , ( max_memory_usage ) , ( buf ) , ( length ) STREAMS_CC )

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

// #define PHP_STREAM_IS_MEMORY       & php_stream_memory_ops

// #define PHP_STREAM_IS_TEMP       & php_stream_temp_ops
