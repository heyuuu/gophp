// <<generate>>

package zend

import (
	"sik/core"
)

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

// #define ZEND_LONG_H

// # include "main/php_stdint.h"

/* This is the heart of the whole int64 enablement in zval. */

const ZEND_ENABLE_ZVAL_LONG64 = 1

/* Integer types. */

type ZendLong = int64
type ZendUlong = uint64
type ZendOffT = int64

const ZEND_LONG_MAX ZendLong = INT64_MAX
const ZEND_LONG_MIN float = INT64_MIN
const ZEND_ULONG_MAX = UINT64_MAX

func Z_L(i int) __auto__  { return int64(i) }
func Z_UL(i int) __auto__ { return uint64(i) }

const SIZEOF_ZEND_LONG = 8

/* Conversion macros. */

const ZEND_LTOA_BUF_LEN = 65
const ZEND_LONG_FMT string = "%" + "lld"
const ZEND_ULONG_FMT *byte = "%" + "llu"
const ZEND_XLONG_FMT = "%" + PRIx64
const ZEND_LONG_FMT_SPEC = "lld"
const ZEND_ULONG_FMT_SPEC = "llu"

func ZEND_LTOA(i __auto__, s []char, len_ __auto__) {
	var st int = core.Snprintf(s, len_, ZEND_LONG_FMT, i)
	s[st] = '0'
}
func ZEND_ATOL(i __auto__, s __auto__) __auto__ {
	i = atoll(s)
	return i
}
func ZEND_STRTOL(s0 __auto__, s1 **byte, base int) __auto__  { return strtoll(s0, s1, base) }
func ZEND_STRTOUL(s0 __auto__, s1 **byte, base int) __auto__ { return strtoull(s0, s1, base) }

const ZEND_STRTOL_PTR = strtoll
const ZEND_STRTOUL_PTR = strtoull
const ZEND_ABS = imaxabs
const MAX_LENGTH_OF_LONG = 20
const LONG_MIN_DIGITS = "9223372036854775808"

var LongMinDigits []byte = LONG_MIN_DIGITS

const ZEND_ADDR_FMT = "0x%016zx"
