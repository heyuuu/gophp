// <<generate>>

package streams

import (
	"sik/core"
	"sik/zend"
)

// Source: <main/streams/php_stream_filter_api.h>

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
   | With suggestions from:                                               |
   |      Moriyoshi Koizumi <moriyoshi@at.wakwak.com>                     |
   |      Sara Golemon      <pollita@php.net>                             |
   +----------------------------------------------------------------------+
*/

const PHP_STREAM_FILTER_READ = 0x1
const PHP_STREAM_FILTER_WRITE = 0x2
const PHP_STREAM_FILTER_ALL zend.ZendLong = PHP_STREAM_FILTER_READ | PHP_STREAM_FILTER_WRITE

type PhpStreamFilterStatusT = int

const (
	PSFS_ERR_FATAL = iota
	PSFS_FEED_ME
	PSFS_PASS_ON
)

/* Buckets API. */

func PhpStreamBucketAddref(bucket __auto__) int {
	bucket.refcount++
	return bucket.refcount - 1
}

const PSFS_FLAG_NORMAL = 0
const PSFS_FLAG_FLUSH_INC = 1
const PSFS_FLAG_FLUSH_CLOSE = 2

/* stack filter onto a stream */

func PhpStreamFilterAlloc(fops *PhpStreamFilterOps, thisptr any, persistent uint8) *core.PhpStreamFilter {
	return _phpStreamFilterAlloc(fops, thisptr, persistent)
}
func PhpStreamFilterAllocRel(fops *PhpStreamFilterOps, thisptr any, persistent uint8) *core.PhpStreamFilter {
	return _phpStreamFilterAlloc(fops, thisptr, persistent)
}
func PhpStreamFilterPrepend(chain *PhpStreamFilterChain, filter *core.PhpStreamFilter) {
	_phpStreamFilterPrepend(chain, filter)
}
func PhpStreamFilterAppend(chain PhpStreamFilterChain, filter *core.PhpStreamFilter) {
	_phpStreamFilterAppend(chain, filter)
}
func PhpStreamFilterFlush(filter *core.PhpStreamFilter, finish int) int {
	return _phpStreamFilterFlush(filter, finish)
}
func PhpStreamIsFiltered(stream *core.PhpStream) bool {
	return stream.readfilters.GetHead() != nil || stream.writefilters.GetHead() != nil
}
