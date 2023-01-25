// <<generate>>

package core

// Source: <main/rfc1867.h>

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

// Source: <main/rfc1867.c>

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
   | Authors: Rasmus Lerdorf <rasmus@php.net>                             |
   |          Jani Taskinen <jani@php.net>                                |
   +----------------------------------------------------------------------+
*/

/* The longest property name we use in an uploaded file array */

/* The longest anonymous name */

/* Errors */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/*
 * Fill up the buffer with client data.
 * Returns number of bytes added to buffer.
 */

/* eof if we are out of bytes, or if we hit the final boundary */

/* create new multipart_buffer structure */

/*
 * Gets the next CRLF terminated line from the input buffer.
 * If it doesn't find a CRLF, and the buffer isn't completely full, returns
 * NULL; otherwise, returns the beginning of the null-terminated line,
 * minus the CRLF.
 *
 * Note that we really just look for LF terminated lines. This works
 * around a bug in internet explorer for the macintosh which sends mime
 * boundaries that are only LF terminated when you use an image submit
 * button in a multipart/form-data form.
 */

/* Returns the next CRLF terminated line from the client */

/* Free header entry */

/* finds a boundary */

/* parse headers */

/*
 * Search for a string in a fixed-length byte string.
 * If partial is true, partial matches are allowed at the end of the buffer.
 * Returns NULL if not found, or a pointer to the start of the first match.
 */

/* read until a boundary condition */

/*
  XXX: this is horrible memory-usage-wise, but we only expect
  to do this on small pieces of form data.
*/

/* }}} */

/* }}} */

/* }}} */
