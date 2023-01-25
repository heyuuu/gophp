// <<generate>>

package core

// Source: <main/php_streams.h>

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
   | Author: Wez Furlong (wez@thebrainroom.com)                           |
   +----------------------------------------------------------------------+
*/

/* {{{ Streams memory debugging stuff */

/* these functions relay the file/line number information. They are depth aware, so they will pass
 * the ultimate ancestor, which is useful, because there can be several layers of calls */

/* }}} */

/* operations on streams that are file-handles */

/* set this when the stream might represent "interactive" data.
 * When set, the read buffer will avoid certain operations that
 * might otherwise cause the read to block for much longer than
 * is strictly required. */

/* state definitions when closing down; these are private to streams.c */

/* allocate a new stream for a particular ops */

/* use this to tell the stream that it is OK if we don't explicitly close it */

/* use this to assign the stream to a zval and tell the stream that is
 * has been exported to the engine; it will expect to be closed automatically
 * when the resources are auto-destructed */

/* php_stream_printf macro & function require */

/* CAREFUL! this is equivalent to puts NOT fputs! */

/* Flags for mkdir method in wrapper ops */

/* define REPORT __special__  ERRORS 8 (below) */

/* change the blocking mode of stream: value == 1 => blocking, value == 0 => non-blocking. */

/* change the buffering mode of stream. value is a PHP_STREAM_BUFFER_XXXX value, ptrparam is a ptr to a size_t holding
 * the required buffer size */

/* set the timeout duration for reads on the stream. ptrparam is a pointer to a struct timeval * */

/* set or release lock on a stream */

/* whether or not locking is supported */

/* option code used by the php_stream_xport_XXX api */

/* Check if the stream is still "live"; for sockets/pipes this means the socket
 * is still connected; for files, this does not really have meaning */

/* Enable/disable blocking reads on anonymous pipes on Windows. */

/* copy up to maxlen bytes from src to dest.  If maxlen is PHP_STREAM_COPY_ALL,
 * copy until eof(src). */

/* read all data from stream and put into a buffer. Caller must free buffer
 * when done. */

/* output all data from a stream */

/* coerce the stream into some other form */

/* cast as a POSIX fd or socketd */

/* cast as a socketd */

/* cast as fd/socket for select purposes */

/* try really, really hard to make sure the cast happens (avoid using this flag if possible) */

/* use this to check if a stream can be cast into another form */

/* use this to check if a stream is of a particular type:
 * PHPAPI int php_stream_is(php_stream *stream, php_stream_ops *ops); */

/* Wrappers support */

/* If you don't need to write to the stream, but really need to
 * be able to seek, use this flag in your options. */

/* If you are going to end up casting the stream into a FILE* or
 * a socket, pass this flag and the streams/wrappers will not use
 * buffering mechanisms while reading the headers, so that HTTP
 * wrapped streams will work consistently.
 * If you omit this flag, streams will use buffering and should end
 * up working more optimally.
 * */

/* this flag applies to php_stream_locate_url_wrapper */

/* this flag is only used by include/require functions */

/* this flag tells streams to ONLY open urls */

/* this flag is used when only the headers from HTTP request are to be fetched */

/* don't apply open_basedir checks */

/* get (or create) a persistent version of the stream */

/* use glob stream for directory open in plain files stream */

/* don't check allow_url_fopen and allow_url_include */

/* assume the path passed in exists and is fully expanded, avoiding syscalls */

/* Allow blocking reads on anonymous pipes on Windows. */

/* Antique - no longer has meaning */

/* pushes an error message onto the stack for a wrapper instance */

/* DO NOT call this on streams that are referenced by resources! */

/* Give other modules access to the url_stream_wrappers_hash and stream_filters_hash */

/* Definitions for user streams */

/* Stream metadata definitions */
