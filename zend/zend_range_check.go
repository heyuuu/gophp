// <<generate>>

package zend

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

// #define ZEND_LONG_CAN_OVFL_INT       1

// #define ZEND_LONG_CAN_OVFL_UINT       1

/* size_t can always overflow signed int on the same platform.
   Furthermore, by the current design, size_t can always
   overflow zend_long. */

// #define ZEND_SIZE_T_CAN_OVFL_UINT       1

/* zend_long vs. (unsigned) int checks. */

// #define ZEND_LONG_INT_OVFL(zlong) UNEXPECTED ( ( zlong ) > ( zend_long ) INT_MAX )

// #define ZEND_LONG_INT_UDFL(zlong) UNEXPECTED ( ( zlong ) < ( zend_long ) INT_MIN )

// #define ZEND_LONG_EXCEEDS_INT(zlong) UNEXPECTED ( ZEND_LONG_INT_OVFL ( zlong ) || ZEND_LONG_INT_UDFL ( zlong ) )

// #define ZEND_LONG_UINT_OVFL(zlong) UNEXPECTED ( ( zlong ) < 0 || ( zlong ) > ( zend_long ) UINT_MAX )

/* size_t vs (unsigned) int checks. */

// #define ZEND_SIZE_T_INT_OVFL(size) UNEXPECTED ( ( size ) > ( size_t ) INT_MAX )

// #define ZEND_SIZE_T_UINT_OVFL(size) UNEXPECTED ( ( size ) > ( size_t ) UINT_MAX )

/* Comparison zend_long vs size_t */

// #define ZEND_SIZE_T_GT_ZEND_LONG(size,zlong) ( ( zlong ) < 0 || ( size ) > ( size_t ) ( zlong ) )

// #define ZEND_SIZE_T_GTE_ZEND_LONG(size,zlong) ( ( zlong ) < 0 || ( size ) >= ( size_t ) ( zlong ) )

// #define ZEND_SIZE_T_LT_ZEND_LONG(size,zlong) ( ( zlong ) >= 0 && ( size ) < ( size_t ) ( zlong ) )

// #define ZEND_SIZE_T_LTE_ZEND_LONG(size,zlong) ( ( zlong ) >= 0 && ( size ) <= ( size_t ) ( zlong ) )
