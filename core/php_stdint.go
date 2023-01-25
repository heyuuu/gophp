// <<generate>>

package core

// Source: <main/php_stdint.h>

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
   | Author: Michael Wallner <mike@php.net>                               |
   +----------------------------------------------------------------------+
*/

// #define PHP_STDINT_H

/* C99 requires these for C++ to get the definitions
 * of INT64_MAX and other macros used by Zend/zend_long.h
 * C11 drops this requirement, so these effectively
 * just backport that piece of behavior.
 *
 * These defines are placed here instead of
 * with the include below, because sys/types
 * and inttypes may include stdint themselves.
 * And these definitions MUST come first.
 */

// # include "php_config.h"

// # include < sys / types . h >

// # include < inttypes . h >

// # include < stdint . h >

func INT8_C(c __auto__) __auto__   { return c }
func UINT8_C(c __auto__) __auto__  { return c }
func INT16_C(c __auto__) __auto__  { return c }
func UINT16_C(c __auto__) __auto__ { return c }

// #define UINT32_C(c) c ## U
