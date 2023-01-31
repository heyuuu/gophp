// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	r "sik/runtime"
	"sik/sapi/cgi"
	"sik/sapi/cli"
	"sik/zend"
)

func PhpStreamOutputWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	core.PHPWRITE(buf, count)
	return count
}
func PhpStreamOutputRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	stream.SetEof(1)
	return -1
}
func PhpStreamOutputClose(stream *core.PhpStream, close_handle int) int        { return 0 }
func PhpStreamInputWrite(stream *core.PhpStream, buf *byte, count int) ssize_t { return -1 }
func PhpStreamInputRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var input *PhpStreamInputT = stream.GetAbstract()
	var read ssize_t
	if !(core.SG(post_read)) && core.SG(read_post_bytes) < int64(input.GetPosition()+count) {

		/* read requested data from SAPI */

		var read_bytes int = core.SapiReadPostBlock(buf, count)
		if read_bytes > 0 {
			core.PhpStreamSeek(input.GetBody(), 0, r.SEEK_END)
			core.PhpStreamWrite(input.GetBody(), buf, read_bytes)
		}
	}
	if input.GetBody().GetReadfilters().GetHead() == nil {

		/* If the input stream contains filters, it's not really seekable. The
		   input->position is likely to be wrong for unfiltered data. */

		core.PhpStreamSeek(input.GetBody(), input.GetPosition(), r.SEEK_SET)

		/* If the input stream contains filters, it's not really seekable. The
		   input->position is likely to be wrong for unfiltered data. */

	}
	read = core.PhpStreamRead(input.GetBody(), buf, count)
	if !read || read == size_t-1 {
		stream.SetEof(1)
	} else {
		input.SetPosition(input.GetPosition() + read)
	}
	return read
}
func PhpStreamInputClose(stream *core.PhpStream, close_handle int) int {
	zend.Efree(stream.GetAbstract())
	stream.SetAbstract(nil)
	return 0
}
func PhpStreamInputFlush(stream *core.PhpStream) int { return -1 }
func PhpStreamInputSeek(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffset *zend.ZendOffT) int {
	var input *PhpStreamInputT = stream.GetAbstract()
	if input.GetBody() != nil {
		var sought int = core.PhpStreamSeek(input.GetBody(), offset, whence)
		input.SetPosition(input.GetBody().GetPosition())
		*newoffset = input.GetPosition()
		return sought
	}
	return -1
}
func PhpStreamApplyFilterList(stream *core.PhpStream, filterlist *byte, read_chain int, write_chain int) {
	var p *byte
	var token *byte = nil
	var temp_filter *core.PhpStreamFilter
	p = core.PhpStrtokR(filterlist, "|", &token)
	for p != nil {
		PhpUrlDecode(p, strlen(p))
		if read_chain != 0 {
			if b.Assign(&temp_filter, streams.PhpStreamFilterCreate(p, nil, stream.GetIsPersistent())) {
				streams.PhpStreamFilterAppend(stream.GetReadfilters(), temp_filter)
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to create filter (%s)", p)
			}
		}
		if write_chain != 0 {
			if b.Assign(&temp_filter, streams.PhpStreamFilterCreate(p, nil, stream.GetIsPersistent())) {
				streams.PhpStreamFilterAppend(stream.GetWritefilters(), temp_filter)
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to create filter (%s)", p)
			}
		}
		p = core.PhpStrtokR(nil, "|", &token)
	}
}
func PhpStreamUrlWrapPhp(wrapper *core.PhpStreamWrapper, path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
	var fd int = -1
	var mode_rw int = 0
	var stream *core.PhpStream = nil
	var p *byte
	var token *byte = nil
	var pathdup *byte
	var max_memory zend.ZendLong
	var file *r.FILE = nil
	if !(strncasecmp(path, "php://", 6)) {
		path += 6
	}
	if !(strncasecmp(path, "temp", 4)) {
		path += 4
		max_memory = core.PHP_STREAM_MAX_MEM
		if !(strncasecmp(path, "/maxmemory:", 11)) {
			path += 11
			max_memory = zend.ZEND_STRTOL(path, nil, 10)
			if max_memory < 0 {
				zend.ZendThrowError(nil, "Max memory must be >= 0")
				return nil
			}
		}
		mode_rw = streams.PhpStreamModeFromStr(mode)
		return core.PhpStreamTempCreate(mode_rw, max_memory)
	}
	if !(strcasecmp(path, "memory")) {
		mode_rw = streams.PhpStreamModeFromStr(mode)
		return core.PhpStreamMemoryCreate(mode_rw)
	}
	if !(strcasecmp(path, "output")) {
		return core.PhpStreamAlloc(&PhpStreamOutputOps, nil, 0, "wb")
	}
	if !(strcasecmp(path, "input")) {
		var input *PhpStreamInputT
		if (options&core.STREAM_OPEN_FOR_INCLUDE) != 0 && !(core.PG(allow_url_include)) {
			if (options & core.REPORT_ERRORS) != 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "URL file-access is disabled in the server configuration")
			}
			return nil
		}
		input = zend.Ecalloc(1, b.SizeOf("* input"))
		if b.Assign(&(input.GetBody()), core.SG(request_info).request_body) {
			core.PhpStreamRewind(input.GetBody())
		} else {
			input.SetBody(core.PhpStreamTempCreateEx(core.TEMP_STREAM_DEFAULT, core.SAPI_POST_BLOCK_SIZE, core.PG(upload_tmp_dir)))
			core.SG(request_info).request_body = input.GetBody()
		}
		return core.PhpStreamAlloc(&PhpStreamInputOps, input, 0, "rb")
	}
	if !(strcasecmp(path, "stdin")) {
		if (options&core.STREAM_OPEN_FOR_INCLUDE) != 0 && !(core.PG(allow_url_include)) {
			if (options & core.REPORT_ERRORS) != 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "URL file-access is disabled in the server configuration")
			}
			return nil
		}
		if !(strcmp(core.sapi_module.GetName(), "cli")) {
			var cli_in int = 0
			fd = cgi.STDIN_FILENO
			if cli_in != 0 {
				fd = dup(fd)
			} else {
				cli_in = 1
				file = stdin
			}
		} else {
			fd = dup(cgi.STDIN_FILENO)
		}
	} else if !(strcasecmp(path, "stdout")) {
		if !(strcmp(core.sapi_module.GetName(), "cli")) {
			var cli_out int = 0
			fd = cli.STDOUT_FILENO
			if b.PostInc(&cli_out) {
				fd = dup(fd)
			} else {
				cli_out = 1
				file = stdout
			}
		} else {
			fd = dup(cli.STDOUT_FILENO)
		}
	} else if !(strcasecmp(path, "stderr")) {
		if !(strcmp(core.sapi_module.GetName(), "cli")) {
			var cli_err int = 0
			fd = cli.STDERR_FILENO
			if b.PostInc(&cli_err) {
				fd = dup(fd)
			} else {
				cli_err = 1
				file = stderr
			}
		} else {
			fd = dup(cli.STDERR_FILENO)
		}
	} else if !(strncasecmp(path, "fd/", 3)) {
		var start *byte
		var end *byte
		var fildes_ori zend.ZendLong
		var dtablesize int
		if strcmp(core.sapi_module.GetName(), "cli") {
			if (options & core.REPORT_ERRORS) != 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Direct access to file descriptors is only available from command-line PHP")
			}
			return nil
		}
		if (options&core.STREAM_OPEN_FOR_INCLUDE) != 0 && !(core.PG(allow_url_include)) {
			if (options & core.REPORT_ERRORS) != 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "URL file-access is disabled in the server configuration")
			}
			return nil
		}
		start = &path[3]
		fildes_ori = zend.ZEND_STRTOL(start, &end, 10)
		if end == start || (*end) != '0' {
			streams.PhpStreamWrapperLogError(wrapper, options, "php://fd/ stream must be specified in the form php://fd/<orig fd>")
			return nil
		}
		dtablesize = getdtablesize()
		if fildes_ori < 0 || fildes_ori >= dtablesize {
			streams.PhpStreamWrapperLogError(wrapper, options, "The file descriptors must be non-negative numbers smaller than %d", dtablesize)
			return nil
		}
		fd = dup(int(fildes_ori))
		if fd == -1 {
			streams.PhpStreamWrapperLogError(wrapper, options, "Error duping file descriptor "+zend.ZEND_LONG_FMT+"; possibly it doesn't exist: "+"[%d]: %s", fildes_ori, errno, strerror(errno))
			return nil
		}
	} else if !(strncasecmp(path, "filter/", 7)) {

		/* Save time/memory when chain isn't specified */

		if strchr(mode, 'r') || strchr(mode, '+') {
			mode_rw |= streams.PHP_STREAM_FILTER_READ
		}
		if strchr(mode, 'w') || strchr(mode, '+') || strchr(mode, 'a') {
			mode_rw |= streams.PHP_STREAM_FILTER_WRITE
		}
		pathdup = zend.Estrndup(path+6, strlen(path+6))
		p = strstr(pathdup, "/resource=")
		if p == nil {
			zend.ZendThrowError(nil, "No URL resource specified")
			zend.Efree(pathdup)
			return nil
		}
		if !(b.Assign(&stream, core.PhpStreamOpenWrapper(p+10, mode, options, opened_path))) {
			zend.Efree(pathdup)
			return nil
		}
		*p = '0'
		p = core.PhpStrtokR(pathdup+1, "/", &token)
		for p != nil {
			if !(strncasecmp(p, "read=", 5)) {
				PhpStreamApplyFilterList(stream, p+5, 1, 0)
			} else if !(strncasecmp(p, "write=", 6)) {
				PhpStreamApplyFilterList(stream, p+6, 0, 1)
			} else {
				PhpStreamApplyFilterList(stream, p, mode_rw&streams.PHP_STREAM_FILTER_READ, mode_rw&streams.PHP_STREAM_FILTER_WRITE)
			}
			p = core.PhpStrtokR(nil, "/", &token)
		}
		zend.Efree(pathdup)
		if zend.EG__().GetException() != nil {
			core.PhpStreamClose(stream)
			return nil
		}
		return stream
	} else {

		/* invalid php://thingy */

		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid php:// URL specified")
		return nil
	}

	/* must be stdin, stderr or stdout */

	if fd == -1 {

		/* failed to dup */

		return nil

		/* failed to dup */

	}
	if file != nil {
		stream = streams.PhpStreamFopenFromFile(file, mode)
	} else {
		stream = streams.PhpStreamFopenFromFd(fd, mode, nil)
		if stream == nil {
			close(fd)
		}
	}
	return stream
}
