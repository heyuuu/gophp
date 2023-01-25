// <<generate>>

package streams

// Source: <main/streams/php_streams_int.h>

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

/* This functions transforms the first char to 'w' if it's not 'r', 'a' or 'w'
 * and strips any subsequent chars except '+' and 'b'.
 * Use this to sanitize stream->mode if you call e.g. fdopen, fopencookie or
 * any other function that expects standard modes and you allow non-standard
 * ones. result should be a char[5]. */
