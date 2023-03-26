package streams

import (
	"sik/zend"
)

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

const PSFS_FLAG_NORMAL = 0
const PSFS_FLAG_FLUSH_INC = 1
const PSFS_FLAG_FLUSH_CLOSE = 2

/* stack filter onto a stream */
