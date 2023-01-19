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

// #define ZEND_LONG_H

// # include "main/php_stdint.h"

/* This is the heart of the whole int64 enablement in zval. */

// #define ZEND_ENABLE_ZVAL_LONG64       1

/* Integer types. */

type ZendLong = int64
type ZendUlong = uint64
type ZendOffT = int64

// #define ZEND_LONG_MAX       INT64_MAX

// #define ZEND_LONG_MIN       INT64_MIN

// #define ZEND_ULONG_MAX       UINT64_MAX

// #define Z_L(i) INT64_C ( i )

// #define Z_UL(i) UINT64_C ( i )

// #define SIZEOF_ZEND_LONG       8

/* Conversion macros. */

// #define ZEND_LTOA_BUF_LEN       65

// #define ZEND_LONG_FMT       "%" PRId64

// #define ZEND_ULONG_FMT       "%" PRIu64

// #define ZEND_XLONG_FMT       "%" PRIx64

// #define ZEND_LONG_FMT_SPEC       PRId64

// #define ZEND_ULONG_FMT_SPEC       PRIu64

// #define ZEND_LTOA(i,s,len) do { int st = snprintf ( ( s ) , ( len ) , ZEND_LONG_FMT , ( i ) ) ; ( s ) [ st ] = '\0' ; } while ( 0 )

// #define ZEND_ATOL(i,s) ( i ) = atoll ( ( s ) )

// #define ZEND_STRTOL(s0,s1,base) strtoll ( ( s0 ) , ( s1 ) , ( base ) )

// #define ZEND_STRTOUL(s0,s1,base) strtoull ( ( s0 ) , ( s1 ) , ( base ) )

// #define ZEND_STRTOL_PTR       strtoll

// #define ZEND_STRTOUL_PTR       strtoull

// #define ZEND_ABS       imaxabs

// #define MAX_LENGTH_OF_LONG       20

// #define LONG_MIN_DIGITS       "9223372036854775808"

var LongMinDigits []byte = "9223372036854775808"

// #define ZEND_ADDR_FMT       "0x%016zx"
