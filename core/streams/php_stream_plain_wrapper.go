// <<generate>>

package streams

import (
	"sik/core"
	r "sik/runtime"
	"sik/zend"
)

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

func PhpStreamFopenFromFile(file *r.FILE, mode *byte) *core.PhpStream {
	return _phpStreamFopenFromFile(file, mode)
}
func PhpStreamFopenFromFd(fd int, mode *byte, persistent_id *byte) *core.PhpStream {
	return _phpStreamFopenFromFd(fd, mode, persistent_id)
}
func PhpStreamFopenFromPipe(file *r.FILE, mode *byte) *core.PhpStream {
	return _phpStreamFopenFromPipe(file, mode)
}

// #define php_stream_fopen_tmpfile() _php_stream_fopen_tmpfile ( 0 STREAMS_CC )

func PhpStreamFopenTemporaryFile(dir *byte, pfx string, opened_path **zend.ZendString) *core.PhpStream {
	return _phpStreamFopenTemporaryFile(dir, pfx, opened_path)
}

/* This is a utility API for extensions that are opening a stream, converting it
 * to a FILE* and then closing it again.  Be warned that fileno() on the result
 * will most likely fail on systems with fopencookie. */

func PhpStreamOpenWrapperAsFile(path *byte, mode string, options int, opened_path **zend.ZendString) *r.FILE {
	return _phpStreamOpenWrapperAsFile(path, mode, options, opened_path)
}

/* parse standard "fopen" modes into open() flags */
