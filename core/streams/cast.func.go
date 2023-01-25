// <<generate>>

package streams

import (
	b "sik/builtin"
	"sik/core"
	r "sik/runtime"
	"sik/zend"
)

func Fopencookie(cookie any, mode *byte, funcs *COOKIE_IO_FUNCTIONS_T) *r.FILE {
	return funopen(cookie, funcs.GetReader(), funcs.GetWriter(), funcs.GetSeeker(), funcs.GetCloser())
}
func StreamCookieReader(cookie any, buffer *byte, size int) int {
	var ret int
	ret = core.PhpStreamRead((*core.PhpStream)(cookie), buffer, size)
	return ret
}
func StreamCookieWriter(cookie any, buffer *byte, size int) int {
	return core.PhpStreamWrite((*core.PhpStream)(cookie), (*byte)(buffer), size)
}
func StreamCookieSeeker(cookie any, position zend.ZendOffT, whence int) PHP_FPOS_T {
	return PHP_FPOS_T(core.PhpStreamSeek((*core.PhpStream)(cookie), position, whence))
}
func StreamCookieCloser(cookie any) int {
	var stream *core.PhpStream = (*core.PhpStream)(cookie)

	/* prevent recursion */

	stream.fclose_stdiocast = core.PHP_STREAM_FCLOSE_NONE
	return core.PhpStreamFree(stream, core.PHP_STREAM_FREE_CLOSE|core.PHP_STREAM_FREE_KEEP_RSRC|core.PHP_STREAM_FREE_RSRC_DTOR)
}
func PhpStreamModeSanitizeFdopenFopencookie(stream *core.PhpStream, result *byte) {
	/* replace modes not supported by fdopen and fopencookie, but supported
	 * by PHP's fread(), so that their calls won't fail */

	var cur_mode *byte = stream.mode
	var has_plus int = 0
	var has_bin int = 0
	var i int
	var res_curs int = 0
	if cur_mode[0] == 'r' || cur_mode[0] == 'w' || cur_mode[0] == 'a' {
		result[b.PostInc(&res_curs)] = cur_mode[0]
	} else {

		/* assume cur_mode[0] is 'c' or 'x'; substitute by 'w', which should not
		 * truncate anything in fdopen/fopencookie */

		result[b.PostInc(&res_curs)] = 'w'

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
		result[b.PostInc(&res_curs)] = 'b'
	}
	if has_plus != 0 {
		result[b.PostInc(&res_curs)] = '+'
	}
	result[res_curs] = '0'
}
func _phpStreamCast(stream *core.PhpStream, castas int, ret *any, show_err int) int {
	var flags int = castas & core.PHP_STREAM_CAST_MASK
	castas &= ^core.PHP_STREAM_CAST_MASK

	/* synchronize our buffer (if possible) */

	if ret != nil && castas != core.PHP_STREAM_AS_FD_FOR_SELECT {
		core.PhpStreamFlush(stream)
		if stream.ops.seek != nil && (stream.flags&core.PHP_STREAM_FLAG_NO_SEEK) == 0 {
			var dummy zend.ZendOffT
			stream.ops.seek(stream, stream.position, r.SEEK_SET, &dummy)
			stream.writepos = 0
			stream.readpos = stream.writepos
		}
	}

	/* filtered streams can only be cast as stdio, and only when fopencookie is present */

	if castas == core.PHP_STREAM_AS_STDIO {
		if stream.stdiocast != nil {
			if ret != nil {
				*((**r.FILE)(ret)) = stream.stdiocast
			}
			goto exit_success
		}

		/* if the stream is a stdio stream let's give it a chance to respond
		 * first, to avoid doubling up the layers of stdio with an fopencookie */

		if core.PhpStreamIs(stream, core.PHP_STREAM_IS_STDIO) && stream.ops.cast != nil && !(PhpStreamIsFiltered(stream)) && stream.ops.cast(stream, castas, ret) == zend.SUCCESS {
			goto exit_success
		}

		/* if just checking, say yes we can be a FILE*, but don't actually create it yet */

		if ret == nil {
			goto exit_success
		}
		var fixed_mode []byte
		PhpStreamModeSanitizeFdopenFopencookie(stream, fixed_mode)
		*((**r.FILE)(ret)) = Fopencookie(stream, fixed_mode, PHP_STREAM_COOKIE_FUNCTIONS)
		if (*ret) != nil {
			var pos zend.ZendOffT
			stream.fclose_stdiocast = core.PHP_STREAM_FCLOSE_FOPENCOOKIE

			/* If the stream position is not at the start, we need to force
			 * the stdio layer to believe it's real location. */

			pos = core.PhpStreamTell(stream)
			if pos > 0 {
				zend.ZendFseek(*ret, pos, r.SEEK_SET)
			}
			goto exit_success
		}

		/* must be either:
		   a) programmer error
		   b) no memory
		   -> lets bail
		*/

		core.PhpErrorDocref(nil, zend.E_ERROR, "fopencookie failed")
		return zend.FAILURE
		if !(PhpStreamIsFiltered(stream)) && stream.ops.cast != nil && stream.ops.cast(stream, castas, nil) == zend.SUCCESS {
			if zend.FAILURE == stream.ops.cast(stream, castas, ret) {
				return zend.FAILURE
			}
			goto exit_success
		} else if (flags & core.PHP_STREAM_CAST_TRY_HARD) != 0 {
			var newstream *core.PhpStream
			newstream = _phpStreamFopenTmpfile(0)
			if newstream != nil {
				var retcopy int = core.PhpStreamCopyToStreamEx(stream, newstream, core.PHP_STREAM_COPY_ALL, nil)
				if retcopy != zend.SUCCESS {
					core.PhpStreamClose(newstream)
				} else {
					var retcast int = core.PhpStreamCast(newstream, castas|flags, (*any)(ret), show_err)
					if retcast == zend.SUCCESS {
						r.Rewind(*((**r.FILE)(ret)))
					}

					/* do some specialized cleanup */

					if (flags & core.PHP_STREAM_CAST_RELEASE) != 0 {
						core.PhpStreamFree(stream, core.PHP_STREAM_FREE_CLOSE_CASTED)
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
	if PhpStreamIsFiltered(stream) {
		if show_err != 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "cannot cast a filtered stream on this system")
		}
		return zend.FAILURE
	} else if stream.ops.cast != nil && stream.ops.cast(stream, castas, ret) == zend.SUCCESS {
		goto exit_success
	}
	if show_err != 0 {

		/* these names depend on the values of the PHP_STREAM_AS_XXX defines in php_streams.h */

		var cast_names []*byte = []*byte{"STDIO FILE*", "File Descriptor", "Socket Descriptor", "select()able descriptor"}
		core.PhpErrorDocref(nil, zend.E_WARNING, "cannot represent a stream of type %s as a %s", stream.ops.label, cast_names[castas])
	}
	return zend.FAILURE
exit_success:
	if stream.writepos-stream.readpos > 0 && stream.fclose_stdiocast != core.PHP_STREAM_FCLOSE_FOPENCOOKIE && (flags&core.PHP_STREAM_CAST_INTERNAL) == 0 {

		/* the data we have buffered will be lost to the third party library that
		 * will be accessing the stream.  Emit a warning so that the end-user will
		 * know that they should try something else */

		core.PhpErrorDocref(nil, zend.E_WARNING, zend.ZEND_LONG_FMT+" bytes of buffered data lost during stream conversion!", zend_long(stream.writepos-stream.readpos))

		/* the data we have buffered will be lost to the third party library that
		 * will be accessing the stream.  Emit a warning so that the end-user will
		 * know that they should try something else */

	}
	if castas == core.PHP_STREAM_AS_STDIO && ret != nil {
		stream.stdiocast = *((**r.FILE)(ret))
	}
	if (flags & core.PHP_STREAM_CAST_RELEASE) != 0 {
		core.PhpStreamFree(stream, core.PHP_STREAM_FREE_CLOSE_CASTED)
	}
	return zend.SUCCESS
}
func _phpStreamOpenWrapperAsFile(path *byte, mode string, options int, opened_path **zend.ZendString) *r.FILE {
	var fp *r.FILE = nil
	var stream *core.PhpStream = nil
	stream = core.PhpStreamOpenWrapperRel(path, mode, options|core.STREAM_WILL_CAST, opened_path)
	if stream == nil {
		return nil
	}
	if core.PhpStreamCast(stream, core.PHP_STREAM_AS_STDIO|core.PHP_STREAM_CAST_TRY_HARD|core.PHP_STREAM_CAST_RELEASE, (*any)(&fp), core.REPORT_ERRORS) == zend.FAILURE {
		core.PhpStreamClose(stream)
		if opened_path != nil && (*opened_path) != nil {
			zend.ZendStringReleaseEx(*opened_path, 0)
		}
		return nil
	}
	return fp
}
func _phpStreamMakeSeekable(origstream *core.PhpStream, newstream **core.PhpStream, flags int) int {
	if newstream == nil {
		return core.PHP_STREAM_FAILED
	}
	*newstream = nil
	if (flags&core.PHP_STREAM_FORCE_CONVERSION) == 0 && origstream.ops.seek != nil {
		*newstream = origstream
		return core.PHP_STREAM_UNCHANGED
	}

	/* Use a tmpfile and copy the old streams contents into it */

	if (flags & core.PHP_STREAM_PREFER_STDIO) != 0 {
		*newstream = _phpStreamFopenTmpfile(0)
	} else {
		*newstream = core.PhpStreamTempNew()
	}
	if (*newstream) == nil {
		return core.PHP_STREAM_FAILED
	}
	if core.PhpStreamCopyToStreamEx(origstream, *newstream, core.PHP_STREAM_COPY_ALL, nil) != zend.SUCCESS {
		core.PhpStreamClose(*newstream)
		*newstream = nil
		return core.PHP_STREAM_CRITICAL
	}
	core.PhpStreamClose(origstream)
	core.PhpStreamSeek(*newstream, 0, r.SEEK_SET)
	return core.PHP_STREAM_RELEASED
}
