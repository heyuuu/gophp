// <<generate>>

package streams

// Source: <main/streams/cast.c>

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
   | Authors: Wez Furlong <wez@thebrainroom.com>                          |
   +----------------------------------------------------------------------+
*/

/* Under BSD, emulate fopencookie using funopen */

/* NetBSD 6.0+ uses off_t instead of fpos_t in funopen */

const PHP_FPOS_T = fpos_t
const HAVE_FOPENCOOKIE = 1
const PHP_EMULATE_FOPENCOOKIE = 1
const PHP_STREAM_COOKIE_FUNCTIONS *COOKIE_IO_FUNCTIONS_T = &StreamCookieFunctions

/* {{{ STDIO with fopencookie */

/* use our fopencookie emulation */

var StreamCookieFunctions COOKIE_IO_FUNCTIONS_T = COOKIE_IO_FUNCTIONS_T{StreamCookieReader, StreamCookieWriter, StreamCookieSeeker, StreamCookieCloser}
