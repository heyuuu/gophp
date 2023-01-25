// <<generate>>

package standard

// Source: <ext/standard/sha1.h>

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
   | Author: Stefan Esser <sesser@php.net>                                |
   +----------------------------------------------------------------------+
*/

/* SHA1 context. */

// Source: <ext/standard/sha1.c>

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
   | Author: Stefan Esser <sesser@php.net>                                |
   +----------------------------------------------------------------------+
*/

/* This code is heavily based on the PHP md5 implementation */

/* {{{ proto string sha1(string str [, bool raw_output])
   Calculate the sha1 hash of a string */

/* }}} */

/* }}} */

/* F, G, H and I are basic SHA1 functions.
 */

/* ROTATE_LEFT rotates x left n bits.
 */

/* W[i]
 */

/* FF, GG, HH, and II transformations for rounds 1, 2, 3, and 4.
 */

/* {{{ PHP_SHA1Init
 * SHA1 initialization. Begins an SHA1 operation, writing a new context.
 */

/* }}} */

/* }}} */

/* }}} */
