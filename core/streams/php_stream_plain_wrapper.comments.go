// <<generate>>

package streams

// Source: <main/streams/php_stream_plain_wrapper.h>

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
   | Author: Wez Furlong <wez@thebrainroom.com>                           |
   +----------------------------------------------------------------------+
*/

/* like fopen, but returns a stream */

/* This is a utility API for extensions that are opening a stream, converting it
 * to a FILE* and then closing it again.  Be warned that fileno() on the result
 * will most likely fail on systems with fopencookie. */

/* parse standard "fopen" modes into open() flags */
