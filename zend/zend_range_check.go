// <<generate>>

package zend

import (
	"sik/core"
)

// Source: <Zend/zend_range_check.h>

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

// #define ZEND_RANGE_CHECK_H

// # include "zend_long.h"

/* Flag macros for basic range recognition. Notable is that
   always sizeof(signed) == sizeof(unsigned), so no need to
   overcomplicate things. */

const ZEND_LONG_CAN_OVFL_INT = 1
const ZEND_LONG_CAN_OVFL_UINT = 1

/* size_t can always overflow signed int on the same platform.
   Furthermore, by the current design, size_t can always
   overflow zend_long. */

const ZEND_SIZE_T_CAN_OVFL_UINT = 1

/* zend_long vs. (unsigned) int checks. */

func ZEND_LONG_INT_OVFL(zlong __auto__) __auto__ {
	return UNEXPECTED(zlong > ZendLong(core.INT_MAX))
}
func ZEND_LONG_INT_UDFL(zlong __auto__) __auto__ {
	return UNEXPECTED(zlong < ZendLong(core.INT_MIN))
}
func ZEND_LONG_EXCEEDS_INT(zlong __auto__) __auto__ {
	return UNEXPECTED(ZEND_LONG_INT_OVFL(zlong) || ZEND_LONG_INT_UDFL(zlong))
}
func ZEND_LONG_UINT_OVFL(zlong __auto__) __auto__ {
	return UNEXPECTED(zlong < 0 || zlong > ZendLong(UINT_MAX))
}

/* size_t vs (unsigned) int checks. */

func ZEND_SIZE_T_INT_OVFL(size int) __auto__ {
	return UNEXPECTED(size > int(core.INT_MAX))
}
func ZEND_SIZE_T_UINT_OVFL(size __auto__) __auto__ { return UNEXPECTED(size > int(UINT_MAX)) }

/* Comparison zend_long vs size_t */

func ZEND_SIZE_T_GT_ZEND_LONG(size __auto__, zlong __auto__) bool {
	return zlong < 0 || size > size_t(zlong)
}
func ZEND_SIZE_T_GTE_ZEND_LONG(size __auto__, zlong __auto__) bool {
	return zlong < 0 || size >= size_t(zlong)
}
func ZEND_SIZE_T_LT_ZEND_LONG(size __auto__, zlong __auto__) bool {
	return zlong >= 0 && size < size_t(zlong)
}
func ZEND_SIZE_T_LTE_ZEND_LONG(size __auto__, zlong __auto__) bool {
	return zlong >= 0 && size <= size_t(zlong)
}
