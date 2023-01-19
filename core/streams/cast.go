// <<generate>>

package streams

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

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

// #define _GNU_SOURCE

// # include "php.h"

// # include "php_globals.h"

// # include "php_network.h"

// # include "php_open_temporary_file.h"

// # include "ext/standard/file.h"

// # include < stddef . h >

// # include < fcntl . h >

// # include "php_streams_int.h"

/* Under BSD, emulate fopencookie using funopen */

/* NetBSD 6.0+ uses off_t instead of fpos_t in funopen */

// #define PHP_FPOS_T       fpos_t

// @type COOKIE_IO_FUNCTIONS_T struct
func Fopencookie(cookie any, mode *byte, funcs *COOKIE_IO_FUNCTIONS_T) *FILE {
	return funopen(cookie, funcs.GetReader(), funcs.GetWriter(), funcs.GetSeeker(), funcs.GetCloser())
}

// #define HAVE_FOPENCOOKIE       1

// #define PHP_EMULATE_FOPENCOOKIE       1

// #define PHP_STREAM_COOKIE_FUNCTIONS       & stream_cookie_functions

/* {{{ STDIO with fopencookie */

/* use our fopencookie emulation */

func StreamCookieReader(cookie any, buffer *byte, size int) int {
	var ret int
	ret = _phpStreamRead((*core.PhpStream)(cookie), buffer, size)
	return ret
}
func StreamCookieWriter(cookie any, buffer *byte, size int) int {
	return _phpStreamWrite((*core.PhpStream)(cookie), (*byte)(buffer), size)
}
func StreamCookieSeeker(cookie any, position zend.ZendOffT, whence int) fpos_t {
	return fpos_t(_phpStreamSeek((*core.PhpStream)(cookie), position, whence))
}
func StreamCookieCloser(cookie any) int {
	var stream *core.PhpStream = (*core.PhpStream)(cookie)

	/* prevent recursion */

	stream.fclose_stdiocast = 0
	return _phpStreamFree(stream, 1|2|64|8)
}

var StreamCookieFunctions COOKIE_IO_FUNCTIONS_T = COOKIE_IO_FUNCTIONS_T{StreamCookieReader, StreamCookieWriter, StreamCookieSeeker, StreamCookieCloser}

/* }}} */

func PhpStreamModeSanitizeFdopenFopencookie(stream *core.PhpStream, result *byte) {
	/* replace modes not supported by fdopen and fopencookie, but supported
	 * by PHP's fread(), so that their calls won't fail */

	var cur_mode *byte = stream.mode
	var has_plus int = 0
	var has_bin int = 0
	var i int
	var res_curs int = 0
	if cur_mode[0] == 'r' || cur_mode[0] == 'w' || cur_mode[0] == 'a' {
		result[g.PostInc(&res_curs)] = cur_mode[0]
	} else {

		/* assume cur_mode[0] is 'c' or 'x'; substitute by 'w', which should not
		 * truncate anything in fdopen/fopencookie */

		result[g.PostInc(&res_curs)] = 'w'

		/* assume cur_mode[0] is 'c' or 'x'; substitute by 'w', which should not
		 * truncate anything in fdopen/fopencookie */

	}

	/* assume current mode has at most length 4 (e.g. wbn+) */

	for i = 1; i < 4 && cur_mode[i] != '0'; i++ {
		if cur_mode[i] == 'b' {
			has_bin = 1
		} else if cur_mode[i] == '+' {
			has_plus = 1
		}
	}
	if has_bin != 0 {
		result[g.PostInc(&res_curs)] = 'b'
	}
	if has_plus != 0 {
		result[g.PostInc(&res_curs)] = '+'
	}
	result[res_curs] = '0'
}

/* }}} */

func _phpStreamCast(stream *core.PhpStream, castas int, ret *any, show_err int) int {
	var flags int = castas & (0x80000000 | 0x40000000 | 0x20000000)
	castas &= ^(0x80000000 | 0x40000000 | 0x20000000)

	/* synchronize our buffer (if possible) */

	if ret != nil && castas != 3 {
		_phpStreamFlush(stream, 0)
		if stream.ops.seek != nil && (stream.flags&0x1) == 0 {
			var dummy zend.ZendOffT
			stream.ops.seek(stream, stream.position, SEEK_SET, &dummy)
			stream.writepos = 0
			stream.readpos = stream.writepos
		}
	}

	/* filtered streams can only be cast as stdio, and only when fopencookie is present */

	if castas == 0 {
		if stream.stdiocast != nil {
			if ret != nil {
				*((**FILE)(ret)) = stream.stdiocast
			}
			goto exit_success
		}

		/* if the stream is a stdio stream let's give it a chance to respond
		 * first, to avoid doubling up the layers of stdio with an fopencookie */

		if stream.ops == &PhpStreamStdioOps && stream.ops.cast != nil && !(stream.readfilters.GetHead() != nil || stream.writefilters.GetHead() != nil) && stream.ops.cast(stream, castas, ret) == zend.SUCCESS {
			goto exit_success
		}

		/* if just checking, say yes we can be a FILE*, but don't actually create it yet */

		if ret == nil {
			goto exit_success
		}
		var fixed_mode []byte
		PhpStreamModeSanitizeFdopenFopencookie(stream, fixed_mode)
		*((**FILE)(ret)) = Fopencookie(stream, fixed_mode, &StreamCookieFunctions)
		if (*ret) != nil {
			var pos zend.ZendOffT
			stream.fclose_stdiocast = 2

			/* If the stream position is not at the start, we need to force
			 * the stdio layer to believe it's real location. */

			pos = _phpStreamTell(stream)
			if pos > 0 {
				fseek(*ret, pos, SEEK_SET)
			}
			goto exit_success
		}

		/* must be either:
		   a) programmer error
		   b) no memory
		   -> lets bail
		*/

		core.PhpErrorDocref(nil, 1<<0, "fopencookie failed")
		return zend.FAILURE
		if !(stream.readfilters.GetHead() != nil || stream.writefilters.GetHead() != nil) && stream.ops.cast != nil && stream.ops.cast(stream, castas, nil) == zend.SUCCESS {
			if zend.FAILURE == stream.ops.cast(stream, castas, ret) {
				return zend.FAILURE
			}
			goto exit_success
		} else if (flags & 0x80000000) != 0 {
			var newstream *core.PhpStream
			newstream = _phpStreamFopenTmpfile(0)
			if newstream != nil {
				var retcopy int = _phpStreamCopyToStreamEx(stream, newstream, size_t-1, nil)
				if retcopy != zend.SUCCESS {
					_phpStreamFree(newstream, 1|2)
				} else {
					var retcast int = _phpStreamCast(newstream, castas|flags, (*any)(ret), show_err)
					if retcast == zend.SUCCESS {
						rewind(*((**FILE)(ret)))
					}

					/* do some specialized cleanup */

					if (flags & 0x40000000) != 0 {
						_phpStreamFree(stream, 1|2|4)
					}

					/* TODO: we probably should be setting .stdiocast and .fclose_stdiocast or
					 * we may be leaking the FILE*. Needs investigation, though. */

					return retcast

					/* TODO: we probably should be setting .stdiocast and .fclose_stdiocast or
					 * we may be leaking the FILE*. Needs investigation, though. */

				}
			}
		}
	}
	if stream.readfilters.GetHead() != nil || stream.writefilters.GetHead() != nil {
		if show_err != 0 {
			core.PhpErrorDocref(nil, 1<<1, "cannot cast a filtered stream on this system")
		}
		return zend.FAILURE
	} else if stream.ops.cast != nil && stream.ops.cast(stream, castas, ret) == zend.SUCCESS {
		goto exit_success
	}
	if show_err != 0 {

		/* these names depend on the values of the PHP_STREAM_AS_XXX defines in php_streams.h */

		var cast_names []*byte = []*byte{"STDIO FILE*", "File Descriptor", "Socket Descriptor", "select()able descriptor"}
		core.PhpErrorDocref(nil, 1<<1, "cannot represent a stream of type %s as a %s", stream.ops.label, cast_names[castas])
	}
	return zend.FAILURE
exit_success:
	if stream.writepos-stream.readpos > 0 && stream.fclose_stdiocast != 2 && (flags&0x20000000) == 0 {

		/* the data we have buffered will be lost to the third party library that
		 * will be accessing the stream.  Emit a warning so that the end-user will
		 * know that they should try something else */

		core.PhpErrorDocref(nil, 1<<1, "%"+"lld"+" bytes of buffered data lost during stream conversion!", zend_long(stream.writepos-stream.readpos))

		/* the data we have buffered will be lost to the third party library that
		 * will be accessing the stream.  Emit a warning so that the end-user will
		 * know that they should try something else */

	}
	if castas == 0 && ret != nil {
		stream.stdiocast = *((**FILE)(ret))
	}
	if (flags & 0x40000000) != 0 {
		_phpStreamFree(stream, 1|2|4)
	}
	return zend.SUCCESS
}

/* }}} */

func _phpStreamOpenWrapperAsFile(path *byte, mode string, options int, opened_path **zend.ZendString) *FILE {
	var fp *FILE = nil
	var stream *core.PhpStream = nil
	stream = _phpStreamOpenWrapperEx(path, mode, options|0x20, opened_path, nil)
	if stream == nil {
		return nil
	}
	if _phpStreamCast(stream, 0|0x80000000|0x40000000, (*any)(&fp), 0x8) == zend.FAILURE {
		_phpStreamFree(stream, 1|2)
		if opened_path != nil && (*opened_path) != nil {
			zend.ZendStringReleaseEx(*opened_path, 0)
		}
		return nil
	}
	return fp
}

/* }}} */

func _phpStreamMakeSeekable(origstream *core.PhpStream, newstream **core.PhpStream, flags int) int {
	if newstream == nil {
		return 2
	}
	*newstream = nil
	if (flags&2) == 0 && origstream.ops.seek != nil {
		*newstream = origstream
		return 0
	}

	/* Use a tmpfile and copy the old streams contents into it */

	if (flags & 1) != 0 {
		*newstream = _phpStreamFopenTmpfile(0)
	} else {
		*newstream = _phpStreamTempCreate(0x0, 2*1024*1024)
	}
	if (*newstream) == nil {
		return 2
	}
	if _phpStreamCopyToStreamEx(origstream, *newstream, size_t-1, nil) != zend.SUCCESS {
		_phpStreamFree(*newstream, 1|2)
		*newstream = nil
		return 3
	}
	_phpStreamFree(origstream, 1|2)
	_phpStreamSeek(*newstream, 0, SEEK_SET)
	return 1
}

/* }}} */
