// <<generate>>

package zend

// Source: <Zend/zend_strtod_int.h>

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

// #define ZEND_STRTOD_INT_H

// # include < stddef . h >

// # include < stdio . h >

// # include < ctype . h >

// # include < stdarg . h >

// # include < math . h >

// # include < sys / types . h >

/* TODO check to undef this option, this might
   make more perf. destroy_freelist()
   should be adapted then. */

const Omit_Private_Memory = 1

/* HEX strings aren't supported as per
   https://wiki.php.net/rfc/remove_hex_support_in_numeric_strings */

const NO_HEX_FP = 1

// # include < inttypes . h >

// #define NO_INFNAN_CHECK

// #define NO_ERRNO

const IEEE_LITTLE_ENDIAN = 1
const IEEE_8087 = 1
