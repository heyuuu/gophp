// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/php_fopen_wrapper.c>

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
   |          Jim Winstead <jimw@php.net>                                 |
   |          Hartmut Holzgraefe <hholzgra@php.net>                       |
   +----------------------------------------------------------------------+
*/

// # include < stdio . h >

// # include < stdlib . h >

// # include "php.h"

// # include "php_globals.h"

// # include "php_standard.h"

// # include "php_memory_streams.h"

// # include "php_fopen_wrappers.h"

// # include "SAPI.h"

func PhpStreamOutputWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	core.PhpOutputWrite(buf, count)
	return count
}

/* }}} */

func PhpStreamOutputRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	stream.eof = 1
	return -1
}

/* }}} */

func PhpStreamOutputClose(stream *core.PhpStream, close_handle int) int { return 0 }

/* }}} */

var PhpStreamOutputOps core.PhpStreamOps = core.PhpStreamOps{PhpStreamOutputWrite, PhpStreamOutputRead, PhpStreamOutputClose, nil, "Output", nil, nil, nil, nil}

// @type PhpStreamInputT struct
type PhpStreamInput = PhpStreamInputT

/* }}} */

func PhpStreamInputWrite(stream *core.PhpStream, buf *byte, count int) ssize_t { return -1 }

/* }}} */

func PhpStreamInputRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var input *PhpStreamInputT = stream.abstract
	var read ssize_t
	if core.sapi_globals.post_read == 0 && core.sapi_globals.read_post_bytes < int64(input.GetPosition()+count) {

		/* read requested data from SAPI */

		var read_bytes int = core.SapiReadPostBlock(buf, count)
		if read_bytes > 0 {
			streams._phpStreamSeek(input.GetBody(), 0, SEEK_END)
			streams._phpStreamWrite(input.GetBody(), buf, read_bytes)
		}
	}
	if input.GetBody().readfilters.head == nil {

		/* If the input stream contains filters, it's not really seekable. The
		   input->position is likely to be wrong for unfiltered data. */

		streams._phpStreamSeek(input.GetBody(), input.GetPosition(), SEEK_SET)

		/* If the input stream contains filters, it's not really seekable. The
		   input->position is likely to be wrong for unfiltered data. */

	}
	read = streams._phpStreamRead(input.GetBody(), buf, count)
	if !read || read == size_t-1 {
		stream.eof = 1
	} else {
		input.SetPosition(input.GetPosition() + read)
	}
	return read
}

/* }}} */

func PhpStreamInputClose(stream *core.PhpStream, close_handle int) int {
	zend._efree(stream.abstract)
	stream.abstract = nil
	return 0
}

/* }}} */

func PhpStreamInputFlush(stream *core.PhpStream) int { return -1 }

/* }}} */

func PhpStreamInputSeek(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffset *zend.ZendOffT) int {
	var input *PhpStreamInputT = stream.abstract
	if input.GetBody() != nil {
		var sought int = streams._phpStreamSeek(input.GetBody(), offset, whence)
		input.SetPosition(input.GetBody().position)
		*newoffset = input.GetPosition()
		return sought
	}
	return -1
}

/* }}} */

var PhpStreamInputOps core.PhpStreamOps = core.PhpStreamOps{PhpStreamInputWrite, PhpStreamInputRead, PhpStreamInputClose, PhpStreamInputFlush, "Input", PhpStreamInputSeek, nil, nil, nil}

func PhpStreamApplyFilterList(stream *core.PhpStream, filterlist *byte, read_chain int, write_chain int) {
	var p *byte
	var token *byte = nil
	var temp_filter *core.PhpStreamFilter
	p = strtok_r(filterlist, "|", &token)
	for p != nil {
		PhpUrlDecode(p, strlen(p))
		if read_chain != 0 {
			if g.Assign(&temp_filter, streams.PhpStreamFilterCreate(p, nil, stream.is_persistent)) {
				streams._phpStreamFilterAppend(&stream.readfilters, temp_filter)
			} else {
				core.PhpErrorDocref(nil, 1<<1, "Unable to create filter (%s)", p)
			}
		}
		if write_chain != 0 {
			if g.Assign(&temp_filter, streams.PhpStreamFilterCreate(p, nil, stream.is_persistent)) {
				streams._phpStreamFilterAppend(&stream.writefilters, temp_filter)
			} else {
				core.PhpErrorDocref(nil, 1<<1, "Unable to create filter (%s)", p)
			}
		}
		p = strtok_r(nil, "|", &token)
	}
}

/* }}} */

func PhpStreamUrlWrapPhp(wrapper *core.PhpStreamWrapper, path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
	var fd int = -1
	var mode_rw int = 0
	var stream *core.PhpStream = nil
	var p *byte
	var token *byte = nil
	var pathdup *byte
	var max_memory zend.ZendLong
	var file *FILE = nil
	if !(strncasecmp(path, "php://", 6)) {
		path += 6
	}
	if !(strncasecmp(path, "temp", 4)) {
		path += 4
		max_memory = 2 * 1024 * 1024
		if !(strncasecmp(path, "/maxmemory:", 11)) {
			path += 11
			max_memory = strtoll(path, nil, 10)
			if max_memory < 0 {
				zend.ZendThrowError(nil, "Max memory must be >= 0")
				return nil
			}
		}
		mode_rw = streams.PhpStreamModeFromStr(mode)
		return streams._phpStreamTempCreate(mode_rw, max_memory)
	}
	if !(strcasecmp(path, "memory")) {
		mode_rw = streams.PhpStreamModeFromStr(mode)
		return streams._phpStreamMemoryCreate(mode_rw)
	}
	if !(strcasecmp(path, "output")) {
		return streams._phpStreamAlloc(&PhpStreamOutputOps, nil, 0, "wb")
	}
	if !(strcasecmp(path, "input")) {
		var input *PhpStreamInputT
		if (options&0x80) != 0 && core.CoreGlobals.allow_url_include == 0 {
			if (options & 0x8) != 0 {
				core.PhpErrorDocref(nil, 1<<1, "URL file-access is disabled in the server configuration")
			}
			return nil
		}
		input = zend._ecalloc(1, g.SizeOf("* input"))
		if g.Assign(&(input.GetBody()), core.sapi_globals.request_info.request_body) {
			streams._phpStreamSeek(input.GetBody(), 0, SEEK_SET)
		} else {
			input.SetBody(streams._phpStreamTempCreateEx(0x0, 0x4000, core.CoreGlobals.upload_tmp_dir))
			core.sapi_globals.request_info.request_body = input.GetBody()
		}
		return streams._phpStreamAlloc(&PhpStreamInputOps, input, 0, "rb")
	}
	if !(strcasecmp(path, "stdin")) {
		if (options&0x80) != 0 && core.CoreGlobals.allow_url_include == 0 {
			if (options & 0x8) != 0 {
				core.PhpErrorDocref(nil, 1<<1, "URL file-access is disabled in the server configuration")
			}
			return nil
		}
		if !(strcmp(core.sapi_module.name, "cli")) {
			var cli_in int = 0
			fd = STDIN_FILENO
			if cli_in != 0 {
				fd = dup(fd)
			} else {
				cli_in = 1
				file = stdin
			}
		} else {
			fd = dup(STDIN_FILENO)
		}
	} else if !(strcasecmp(path, "stdout")) {
		if !(strcmp(core.sapi_module.name, "cli")) {
			var cli_out int = 0
			fd = STDOUT_FILENO
			if g.PostInc(&cli_out) {
				fd = dup(fd)
			} else {
				cli_out = 1
				file = stdout
			}
		} else {
			fd = dup(STDOUT_FILENO)
		}
	} else if !(strcasecmp(path, "stderr")) {
		if !(strcmp(core.sapi_module.name, "cli")) {
			var cli_err int = 0
			fd = STDERR_FILENO
			if g.PostInc(&cli_err) {
				fd = dup(fd)
			} else {
				cli_err = 1
				file = stderr
			}
		} else {
			fd = dup(STDERR_FILENO)
		}
	} else if !(strncasecmp(path, "fd/", 3)) {
		var start *byte
		var end *byte
		var fildes_ori zend.ZendLong
		var dtablesize int
		if strcmp(core.sapi_module.name, "cli") {
			if (options & 0x8) != 0 {
				core.PhpErrorDocref(nil, 1<<1, "Direct access to file descriptors is only available from command-line PHP")
			}
			return nil
		}
		if (options&0x80) != 0 && core.CoreGlobals.allow_url_include == 0 {
			if (options & 0x8) != 0 {
				core.PhpErrorDocref(nil, 1<<1, "URL file-access is disabled in the server configuration")
			}
			return nil
		}
		start = &path[3]
		fildes_ori = strtoll(start, &end, 10)
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
			streams.PhpStreamWrapperLogError(wrapper, options, "Error duping file descriptor "+"%"+"lld"+"; possibly it doesn't exist: "+"[%d]: %s", fildes_ori, errno, strerror(errno))
			return nil
		}
	} else if !(strncasecmp(path, "filter/", 7)) {

		/* Save time/memory when chain isn't specified */

		if strchr(mode, 'r') || strchr(mode, '+') {
			mode_rw |= 0x1
		}
		if strchr(mode, 'w') || strchr(mode, '+') || strchr(mode, 'a') {
			mode_rw |= 0x2
		}
		pathdup = zend._estrndup(path+6, strlen(path+6))
		p = strstr(pathdup, "/resource=")
		if p == nil {
			zend.ZendThrowError(nil, "No URL resource specified")
			zend._efree(pathdup)
			return nil
		}
		if !(g.Assign(&stream, streams._phpStreamOpenWrapperEx(p+10, mode, options, opened_path, nil))) {
			zend._efree(pathdup)
			return nil
		}
		*p = '0'
		p = strtok_r(pathdup+1, "/", &token)
		for p != nil {
			if !(strncasecmp(p, "read=", 5)) {
				PhpStreamApplyFilterList(stream, p+5, 1, 0)
			} else if !(strncasecmp(p, "write=", 6)) {
				PhpStreamApplyFilterList(stream, p+6, 0, 1)
			} else {
				PhpStreamApplyFilterList(stream, p, mode_rw&0x1, mode_rw&0x2)
			}
			p = strtok_r(nil, "/", &token)
		}
		zend._efree(pathdup)
		if zend.EG.exception != nil {
			streams._phpStreamFree(stream, 1|2)
			return nil
		}
		return stream
	} else {

		/* invalid php://thingy */

		core.PhpErrorDocref(nil, 1<<1, "Invalid php:// URL specified")
		return nil
	}

	/* must be stdin, stderr or stdout */

	if fd == -1 {

		/* failed to dup */

		return nil

		/* failed to dup */

	}
	if file != nil {
		stream = streams._phpStreamFopenFromFile(file, mode)
	} else {
		stream = streams._phpStreamFopenFromFd(fd, mode, nil)
		if stream == nil {
			close(fd)
		}
	}
	return stream
}

/* }}} */

var PhpStdioWops core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{PhpStreamUrlWrapPhp, nil, nil, nil, nil, "PHP", nil, nil, nil, nil, nil}
var PhpStreamPhpWrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&PhpStdioWops, nil, 0}
