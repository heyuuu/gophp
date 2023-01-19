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

// #define php_stream_fopen(filename,mode,opened) _php_stream_fopen ( ( filename ) , ( mode ) , ( opened ) , 0 STREAMS_CC )

// #define php_stream_fopen_with_path(filename,mode,path,opened) _php_stream_fopen_with_path ( ( filename ) , ( mode ) , ( path ) , ( opened ) , 0 STREAMS_CC )

// #define php_stream_fopen_from_file(file,mode) _php_stream_fopen_from_file ( ( file ) , ( mode ) STREAMS_CC )

// #define php_stream_fopen_from_fd(fd,mode,persistent_id) _php_stream_fopen_from_fd ( ( fd ) , ( mode ) , ( persistent_id ) STREAMS_CC )

// #define php_stream_fopen_from_pipe(file,mode) _php_stream_fopen_from_pipe ( ( file ) , ( mode ) STREAMS_CC )

// #define php_stream_fopen_tmpfile() _php_stream_fopen_tmpfile ( 0 STREAMS_CC )

// #define php_stream_fopen_temporary_file(dir,pfx,opened_path) _php_stream_fopen_temporary_file ( ( dir ) , ( pfx ) , ( opened_path ) STREAMS_CC )

/* This is a utility API for extensions that are opening a stream, converting it
 * to a FILE* and then closing it again.  Be warned that fileno() on the result
 * will most likely fail on systems with fopencookie. */

// #define php_stream_open_wrapper_as_file(path,mode,options,opened_path) _php_stream_open_wrapper_as_file ( ( path ) , ( mode ) , ( options ) , ( opened_path ) STREAMS_CC )

/* parse standard "fopen" modes into open() flags */
