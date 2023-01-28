// <<generate>>

package zend

// Source: <Zend/zend_long.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Anatol Belski <ab@php.net>                                  |
   +----------------------------------------------------------------------+
*/

/* This is the heart of the whole int64 enablement in zval. */

const ZEND_ENABLE_ZVAL_LONG64 = 1

/* Integer types. */

type ZendLong = int64
type ZendUlong = uint64
type ZendOffT = int64

const ZEND_LONG_MAX ZendLong = INT64_MAX
const ZEND_LONG_MIN float = INT64_MIN
const ZEND_ULONG_MAX = UINT64_MAX
const SIZEOF_ZEND_LONG = 8

/* Conversion macros. */

const ZEND_LTOA_BUF_LEN = 65
const ZEND_LONG_FMT string = "%" + "lld"
const ZEND_ULONG_FMT *byte = "%" + "llu"
const ZEND_XLONG_FMT = "%" + PRIx64
const ZEND_LONG_FMT_SPEC = "lld"
const ZEND_ULONG_FMT_SPEC = "llu"
const ZEND_STRTOL_PTR = strtoll
const ZEND_STRTOUL_PTR = strtoull
const ZEND_ABS = imaxabs
const MAX_LENGTH_OF_LONG = 20
const LONG_MIN_DIGITS = "9223372036854775808"

var LongMinDigits []byte = LONG_MIN_DIGITS

const ZEND_ADDR_FMT = "0x%016zx"
