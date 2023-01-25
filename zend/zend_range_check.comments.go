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

/* Flag macros for basic range recognition. Notable is that
   always sizeof(signed) == sizeof(unsigned), so no need to
   overcomplicate things. */

/* size_t can always overflow signed int on the same platform.
   Furthermore, by the current design, size_t can always
   overflow zend_long. */

/* zend_long vs. (unsigned) int checks. */

/* size_t vs (unsigned) int checks. */

/* Comparison zend_long vs size_t */
