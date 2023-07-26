package streams

import (
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	b "github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

func Fopencookie(cookie any, mode *byte, funcs *COOKIE_IO_FUNCTIONS_T) *r.File {
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

	stream.SetFcloseStdiocast(core.PHP_STREAM_FCLOSE_NONE)
	return core.PhpStreamFree(stream, core.PHP_STREAM_FREE_CLOSE|core.PHP_STREAM_FREE_KEEP_RSRC|core.PHP_STREAM_FREE_RSRC_DTOR)
}
func PhpStreamModeSanitizeFdopenFopencookie(stream *core.PhpStream, result *byte) {
	/* replace modes not supported by fdopen and fopencookie, but supported
	 * by PHP's fread(), so that their calls won't fail */

	var cur_mode *byte = stream.GetMode()
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
func O_phpStreamCast(stream *core.PhpStream, castas int, ret *any, show_err int) int {
	var flags int = castas & core.PHP_STREAM_CAST_MASK
	castas &= ^core.PHP_STREAM_CAST_MASK

	/* synchronize our buffer (if possible) */

	if ret != nil && castas != core.PHP_STREAM_AS_FD_FOR_SELECT {
		core.PhpStreamFlush(stream)
		if stream.GetOps().GetSeek() != nil && !stream.HasFlags(core.PHP_STREAM_FLAG_NO_SEEK) {
			var dummy zend.ZendOffT
			stream.GetOps().GetSeek()(stream, stream.GetPosition(), r.SEEK_SET, &dummy)
			stream.SetWritepos(0)
			stream.SetReadpos(stream.GetWritepos())
		}
	}

	/* filtered streams can only be cast as stdio, and only when fopencookie is present */

	if castas == core.PHP_STREAM_AS_STDIO {
		if stream.GetStdiocast() != nil {
			if ret != nil {
				*((**r.File)(ret)) = stream.GetStdiocast()
			}
			goto exit_success
		}

		/* if the stream is a stdio stream let's give it a chance to respond
		 * first, to avoid doubling up the layers of stdio with an fopencookie */

		if core.PhpStreamIs(stream, core.PHP_STREAM_IS_STDIO) && stream.GetOps().GetCast() != nil && !(PhpStreamIsFiltered(stream)) && stream.GetOps().GetCast()(stream, castas, ret) == types.SUCCESS {
			goto exit_success
		}

		/* if just checking, say yes we can be a FILE*, but don't actually create it yet */

		if ret == nil {
			goto exit_success
		}
		var fixed_mode []byte
		PhpStreamModeSanitizeFdopenFopencookie(stream, fixed_mode)
		*((**r.File)(ret)) = Fopencookie(stream, fixed_mode, PHP_STREAM_COOKIE_FUNCTIONS)
		if (*ret) != nil {
			var pos zend.ZendOffT
			stream.SetFcloseStdiocast(core.PHP_STREAM_FCLOSE_FOPENCOOKIE)

			/* If the stream position is not at the start, we need to force
			 * the stdio layer to believe it's real location. */

			pos = stream.GetPosition()
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

		core.PhpErrorDocref(nil, faults.E_ERROR, "fopencookie failed")
		return types.FAILURE
	}
	if PhpStreamIsFiltered(stream) {
		if show_err != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "cannot cast a filtered stream on this system")
		}
		return types.FAILURE
	} else if stream.GetOps().GetCast() != nil && stream.GetOps().GetCast()(stream, castas, ret) == types.SUCCESS {
		goto exit_success
	}
	if show_err != 0 {

		/* these names depend on the values of the PHP_STREAM_AS_XXX defines in php_streams.h */

		var cast_names []*byte = []*byte{"STDIO FILE*", "File Descriptor", "Socket Descriptor", "select()able descriptor"}
		core.PhpErrorDocref(nil, faults.E_WARNING, "cannot represent a stream of type %s as a %s", stream.GetOps().GetLabel(), cast_names[castas])
	}
	return types.FAILURE
exit_success:
	if stream.GetWritepos()-stream.GetReadpos() > 0 && stream.GetFcloseStdiocast() != core.PHP_STREAM_FCLOSE_FOPENCOOKIE && (flags&core.PHP_STREAM_CAST_INTERNAL) == 0 {

		/* the data we have buffered will be lost to the third party library that
		 * will be accessing the stream.  Emit a warning so that the end-user will
		 * know that they should try something else */

		core.PhpErrorDocref(nil, faults.E_WARNING, zend.ZEND_LONG_FMT+" bytes of buffered data lost during stream conversion!", zend_long(stream.GetWritepos()-stream.GetReadpos()))

		/* the data we have buffered will be lost to the third party library that
		 * will be accessing the stream.  Emit a warning so that the end-user will
		 * know that they should try something else */

	}
	if castas == core.PHP_STREAM_AS_STDIO && ret != nil {
		stream.SetStdiocast(*((**r.File)(ret)))
	}
	if (flags & core.PHP_STREAM_CAST_RELEASE) != 0 {
		core.PhpStreamFree(stream, core.PHP_STREAM_FREE_CLOSE_CASTED)
	}
	return types.SUCCESS
}
func _phpStreamOpenWrapperAsFile(path *byte, mode string, options int, opened_path **types.String) *r.File {
	var fp *r.File = nil
	var stream *core.PhpStream = nil
	stream = core.PhpStreamOpenWrapperRel(path, mode, options|core.STREAM_WILL_CAST, opened_path)
	if stream == nil {
		return nil
	}
	if core.PhpStreamCast(stream, core.PHP_STREAM_AS_STDIO|core.PHP_STREAM_CAST_TRY_HARD|core.PHP_STREAM_CAST_RELEASE, (*any)(&fp), core.REPORT_ERRORS) == types.FAILURE {
		core.PhpStreamClose(stream)
		if opened_path != nil && (*opened_path) != nil {
			// types.ZendStringReleaseEx(*opened_path, 0)
		}
		return nil
	}
	return fp
}
