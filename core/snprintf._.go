// <<generate>>

package core

// Source: <main/snprintf.h>

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
   | Author: Stig Sæther Bakken <ssb@php.net>                             |
   |         Marcus Boerger <helly@php.net>                               |
   +----------------------------------------------------------------------+
*/

type BoolInt = int
type BooleanE = int

const (
	NO  = 0
	YES = 1
)

var Php0cvt func(value float64, ndigit int, dec_point byte, exponent byte, buf *byte) *byte

const Slprintf = ApPhpSlprintf
const Vslprintf = ApPhpVslprintf
const Snprintf = ApPhpSnprintf
const Vsnprintf = ApPhpVsnprintf

type LengthModifierE = int

const (
	LM_STD = 0
	LM_INTMAX_T
	LM_PTRDIFF_T
	LM_LONG_LONG
	LM_SIZE_T
	LM_LONG
	LM_LONG_DOUBLE
	LM_PHP_INT_T
)

type WideInt = long__long
type UWideInt = unsigned__long__long

/* The maximum precision that's allowed for float conversion. Does not include
 * decimal separator, exponent, sign, terminator. Currently does not affect
 * the modes e/f, only g/k/H, as those have a different limit enforced at
 * another level (see NDIG in php_conv_fp()).
 * Applies to the formatting functions of both spprintf.c and snprintf.c, which
 * use equally sized buffers of MAX_BUF_SIZE = 512 to hold the result of the
 * call to php_gcvt().
 * This should be reasonably smaller than MAX_BUF_SIZE (I think MAX_BUF_SIZE - 9
 * should be enough, but let's give some more space) */

const FORMAT_CONV_MAX_PRECISION = 500

// Source: <main/snprintf.c>

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
  | Author:                                                              |
  +----------------------------------------------------------------------+
*/

const LCONV_DECIMAL_POINT = (*lconv).decimal_point

/*
 * Copyright (c) 2002, 2006 Todd C. Miller <Todd.Miller@courtesan.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 *
 * Sponsored in part by the Defense Advanced Research Projects
 * Agency (DARPA) and Air Force Research Laboratory, Air Force
 * Materiel Command, USAF, under agreement number F39502-99-1-0512.
 */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

const FALSE = 0
const TRUE = 1
const NUL = '0'
const INT_NULL = (*int)(0)
const S_NULL = "(null)"
const S_NULL_LEN = 6
const FLOAT_DIGITS = 6
const EXPONENT_LENGTH = 10

/*
 * Convert num to its decimal format.
 * Return value:
 *   - a pointer to a string containing the number (no sign)
 *   - len contains the length of the string
 *   - is_negative is set to TRUE or FALSE depending on the sign
 *     of the number (always set to FALSE if is_unsigned is TRUE)
 *
 * The caller provides a buffer for the string: that is the buf_end argument
 * which is a pointer to the END of the buffer + 1 (i.e. if the buffer
 * is declared as buf[ 100 ], buf_end should be &buf[ 100 ])
 */

/* }}} */

const NDIG = 320

/*
 * Convert a floating point number to a string formats 'f', 'e' or 'E'.
 * The result is placed in buf, and len denotes the length of the string
 * The sign is returned in the is_negative argument (and is not placed
 * in buf).
 */

/* }}} */

/* }}} */

const NUM_BUF_SIZE = 2048

/*
 * Descriptor for buffer area
 */

type Buffy = BufArea

/*
 * The INS_CHAR macro inserts a character in the buffer and writes
 * the buffer back to disk if necessary
 * It uses the char pointers sp and bep:
 *      sp points to the next available character in the buffer
 *      bep points to the end-of-buffer+1
 * While using this macro, note that the nextb pointer is NOT updated.
 *
 * NOTE: Evaluation of the c argument should not have any side-effects
 */

/*
 * This macro does zero padding so that the precision
 * requirement is satisfied. The padding is done by
 * adding '0's to the left of the string that is going
 * to be printed.
 */

/*
 * Macro that does padding. The padding is done by printing
 * the character ch.
 */

/*
 * Prefix the character ch to the string str
 * Increase length
 * Set the has_prefix flag
 */

/*
 * Do format conversion placing the output in buffer
 */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */
