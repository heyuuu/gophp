// <<generate>>

package streams

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

// #define PHP_STREAM_FILTER_READ       0x0001

// #define PHP_STREAM_FILTER_WRITE       0x0002

// #define PHP_STREAM_FILTER_ALL       ( PHP_STREAM_FILTER_READ | PHP_STREAM_FILTER_WRITE )

type PhpStreamFilterStatusT = int

const (
	PSFS_ERR_FATAL = iota
	PSFS_FEED_ME
	PSFS_PASS_ON
)

/* Buckets API. */

// #define php_stream_bucket_addref(bucket) ( bucket ) -> refcount ++

// #define PSFS_FLAG_NORMAL       0

// #define PSFS_FLAG_FLUSH_INC       1

// #define PSFS_FLAG_FLUSH_CLOSE       2

/* stack filter onto a stream */

// #define php_stream_filter_alloc(fops,thisptr,persistent) _php_stream_filter_alloc ( ( fops ) , ( thisptr ) , ( persistent ) STREAMS_CC )

// #define php_stream_filter_alloc_rel(fops,thisptr,persistent) _php_stream_filter_alloc ( ( fops ) , ( thisptr ) , ( persistent ) STREAMS_REL_CC )

// #define php_stream_filter_prepend(chain,filter) _php_stream_filter_prepend ( ( chain ) , ( filter ) )

// #define php_stream_filter_append(chain,filter) _php_stream_filter_append ( ( chain ) , ( filter ) )

// #define php_stream_filter_flush(filter,finish) _php_stream_filter_flush ( ( filter ) , ( finish ) )

// #define php_stream_is_filtered(stream) ( ( stream ) -> readfilters . head || ( stream ) -> writefilters . head )
