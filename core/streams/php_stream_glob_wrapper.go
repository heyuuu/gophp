// <<generate>>

package streams

import (
	"sik/core"
)

// Source: <main/streams/php_stream_glob_wrapper.h>

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
   | Author: Marcus Boerger <helly@php.net>                               |
   +----------------------------------------------------------------------+
*/

func PhpGlobStreamGetPath(stream *core.PhpStream, plen *int) *byte {
	return _phpGlobStreamGetPath(stream, plen)
}
func PhpGlobStreamGetPattern(stream *core.PhpStream, plen *int) *byte {
	return _phpGlobStreamGetPattern(stream, plen)
}
func PhpGlobStreamGetCount(stream *core.PhpStream, pflags *int) int {
	return _phpGlobStreamGetCount(stream, pflags)
}
