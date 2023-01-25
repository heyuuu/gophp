// <<generate>>

package streams

// Source: <main/streams/streams.c>

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
   | Borrowed code from:                                                  |
   |          Rasmus Lerdorf <rasmus@lerdorf.on.ca>                       |
   |          Jim Winstead <jimw@php.net>                                 |
   +----------------------------------------------------------------------+
*/

/* {{{ resource and registration code */

/* }}} */

/* {{{ wrapper error reporting */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* Like php_stream_read(), but reading into a zend_string buffer. This has some similarity
 * to the copy_to_mem() operation, but only performs a single direct read. */

/* If buf == NULL, the buffer will be allocated automatically and will be of an
 * appropriate length to hold the line, regardless of the line length, memory
 * permitting */

/* Writes a buffer directly to a stream, using multiple of the chunk size */

/* push some data through the write filter chain.
 * buf may be NULL, if flags are set to indicate a flush.
 * This may trigger a real write to the stream.
 * Returns the number of bytes consumed from buf by the first filter in the chain.
 * */

/* Returns SUCCESS/FAILURE and sets *len to the number of bytes moved */

/* Returns the number of bytes moved.
 * Returns 1 when source len is 0.
 * Deprecated in favor of php_stream_copy_to_stream_ex() */

/* }}} */

/* Validate protocol scheme names during registration
 * Must conform to /^[a-zA-Z0-9+.-]+$/
 */

/* API for registering GLOBAL wrappers */

/* API for registering VOLATILE wrappers */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */
