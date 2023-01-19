// <<generate>>

package streams

import (
	"sik/core"
	"sik/ext/standard"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/streams/plain_wrapper.c>

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

// # include "php.h"

// # include "php_globals.h"

// # include "php_network.h"

// # include "php_open_temporary_file.h"

// # include "ext/standard/file.h"

// # include "ext/standard/flock_compat.h"

// # include "ext/standard/php_filestat.h"

// # include < stddef . h >

// # include < fcntl . h >

// # include < sys / wait . h >

// # include < sys / file . h >

// # include < sys / mman . h >

// # include "SAPI.h"

// # include "php_streams_int.h"

// #define php_stream_fopen_from_fd_int(fd,mode,persistent_id) _php_stream_fopen_from_fd_int ( ( fd ) , ( mode ) , ( persistent_id ) STREAMS_CC )

// #define php_stream_fopen_from_fd_int_rel(fd,mode,persistent_id) _php_stream_fopen_from_fd_int ( ( fd ) , ( mode ) , ( persistent_id ) STREAMS_REL_CC )

// #define php_stream_fopen_from_file_int(file,mode) _php_stream_fopen_from_file_int ( ( file ) , ( mode ) STREAMS_CC )

// #define php_stream_fopen_from_file_int_rel(file,mode) _php_stream_fopen_from_file_int ( ( file ) , ( mode ) STREAMS_REL_CC )

var PhpGetUidByName func(name *byte, uid *uid_t) int
var PhpGetGidByName func(name *byte, gid *gid_t) int

// #define PLAIN_WRAP_BUF_SIZE(st) ( st )

/* parse standard "fopen" modes into open() flags */

func PhpStreamParseFopenModes(mode *byte, open_flags *int) int {
	var flags int
	switch mode[0] {
	case 'r':
		flags = 0
		break
	case 'w':
		flags = O_TRUNC | O_CREAT
		break
	case 'a':
		flags = O_CREAT | O_APPEND
		break
	case 'x':
		flags = O_CREAT | O_EXCL
		break
	case 'c':
		flags = O_CREAT
		break
	default:

		/* unknown mode */

		return zend.FAILURE
	}
	if strchr(mode, '+') {
		flags |= O_RDWR
	} else if flags != 0 {
		flags |= O_WRONLY
	} else {
		flags |= O_RDONLY
	}
	*open_flags = flags
	return zend.SUCCESS
}

/* {{{ ------- STDIO stream implementation -------*/

// @type PhpStdioStreamData struct

// #define PHP_STDIOP_GET_FD(anfd,data) anfd = ( data ) -> file ? fileno ( ( data ) -> file ) : ( data ) -> fd

func DoFstat(d *PhpStdioStreamData, force int) int {
	if !(d.GetCachedFstat()) || force != 0 && !(d.GetNoForcedFstat()) {
		var fd int
		var r int
		if d.GetFile() != nil {
			fd = fileno(d.GetFile())
		} else {
			fd = d.GetFd()
		}
		r = fstat(fd, &d.sb)
		d.SetCachedFstat(r == 0)
		return r
	}
	return 0
}
func _phpStreamFopenFromFdInt(fd int, mode *byte, persistent_id *byte) *core.PhpStream {
	var self *PhpStdioStreamData
	if persistent_id != nil {
		self = zend.__zendMalloc(g.SizeOf("* self"))
	} else {
		self = zend._emalloc(g.SizeOf("* self"))
	}
	memset(self, 0, g.SizeOf("* self"))
	self.SetFile(nil)
	self.SetIsSeekable(1)
	self.SetIsPipe(0)
	self.SetLockFlag(LOCK_UN)
	self.SetIsProcessPipe(0)
	self.SetTempName(nil)
	self.SetFd(fd)
	return _phpStreamAlloc(&PhpStreamStdioOps, self, persistent_id, mode)
}
func _phpStreamFopenFromFileInt(file *FILE, mode *byte) *core.PhpStream {
	var self *PhpStdioStreamData
	self = zend._emalloc(g.SizeOf("* self"))
	memset(self, 0, g.SizeOf("* self"))
	self.SetFile(file)
	self.SetIsSeekable(1)
	self.SetIsPipe(0)
	self.SetLockFlag(LOCK_UN)
	self.SetIsProcessPipe(0)
	self.SetTempName(nil)
	self.SetFd(fileno(file))
	return _phpStreamAlloc(&PhpStreamStdioOps, self, 0, mode)
}
func _phpStreamFopenTemporaryFile(dir *byte, pfx string, opened_path_ptr **zend.ZendString) *core.PhpStream {
	var opened_path *zend.ZendString = nil
	var fd int
	fd = core.PhpOpenTemporaryFd(dir, pfx, &opened_path)
	if fd != -1 {
		var stream *core.PhpStream
		if opened_path_ptr != nil {
			*opened_path_ptr = opened_path
		}
		stream = _phpStreamFopenFromFdInt(fd, "r+b", nil)
		if stream != nil {
			var self *PhpStdioStreamData = (*PhpStdioStreamData)(stream.abstract)
			stream.wrapper = (*core.PhpStreamWrapper)(&PhpPlainFilesWrapper)
			stream.orig_path = zend._estrndup(opened_path.val, opened_path.len_)
			self.SetTempName(opened_path)
			self.SetLockFlag(LOCK_UN)
			return stream
		}
		close(fd)
		core.PhpErrorDocref(nil, 1<<1, "unable to allocate stream")
		return nil
	}
	return nil
}
func _phpStreamFopenTmpfile(dummy int) *core.PhpStream {
	return _phpStreamFopenTemporaryFile(nil, "php", nil)
}
func DetectIsSeekable(self *PhpStdioStreamData) {}
func _phpStreamFopenFromFd(fd int, mode *byte, persistent_id *byte) *core.PhpStream {
	var stream *core.PhpStream = _phpStreamFopenFromFdInt(fd, mode, persistent_id)
	if stream != nil {
		var self *PhpStdioStreamData = (*PhpStdioStreamData)(stream.abstract)
		DetectIsSeekable(self)
		if !(self.GetIsSeekable()) {
			stream.flags |= 0x1
			stream.position = -1
		} else {
			stream.position = lseek(self.GetFd(), 0, SEEK_CUR)
		}
	}
	return stream
}
func _phpStreamFopenFromFile(file *FILE, mode *byte) *core.PhpStream {
	var stream *core.PhpStream = _phpStreamFopenFromFileInt(file, mode)
	if stream != nil {
		var self *PhpStdioStreamData = (*PhpStdioStreamData)(stream.abstract)
		DetectIsSeekable(self)
		if !(self.GetIsSeekable()) {
			stream.flags |= 0x1
			stream.position = -1
		} else {
			stream.position = ftell(file)
		}
	}
	return stream
}
func _phpStreamFopenFromPipe(file *FILE, mode *byte) *core.PhpStream {
	var self *PhpStdioStreamData
	var stream *core.PhpStream
	self = zend._emalloc(g.SizeOf("* self"))
	memset(self, 0, g.SizeOf("* self"))
	self.SetFile(file)
	self.SetIsSeekable(0)
	self.SetIsPipe(1)
	self.SetLockFlag(LOCK_UN)
	self.SetIsProcessPipe(1)
	self.SetFd(fileno(file))
	self.SetTempName(nil)
	stream = _phpStreamAlloc(&PhpStreamStdioOps, self, 0, mode)
	stream.flags |= 0x1
	return stream
}
func PhpStdiopWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.abstract)
	assert(data != nil)
	if data.GetFd() >= 0 {
		var bytes_written ssize_t = write(data.GetFd(), buf, count)
		if bytes_written < 0 {
			if errno == EAGAIN || errno == EAGAIN {
				return 0
			}
			if errno == EINTR {

				/* TODO: Should this be treated as a proper error or not? */

				return bytes_written

				/* TODO: Should this be treated as a proper error or not? */

			}
			core.PhpErrorDocref(nil, 1<<3, "write of %zu bytes failed with errno=%d %s", count, errno, strerror(errno))
		}
		return bytes_written
	} else {
		if data.GetIsSeekable() && data.GetLastOp() == 'r' {
			fseek(data.GetFile(), 0, SEEK_CUR)
		}
		data.SetLastOp('w')
		return ssize_t(fwrite(buf, 1, count, data.GetFile()))
	}
}
func PhpStdiopRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.abstract)
	var ret ssize_t
	assert(data != nil)
	if data.GetFd() >= 0 {
		ret = read(data.GetFd(), buf, count)
		if ret == size_t-1 && errno == EINTR {

			/* Read was interrupted, retry once,
			   If read still fails, giveup with feof==0
			   so script can retry if desired */

			ret = read(data.GetFd(), buf, count)

			/* Read was interrupted, retry once,
			   If read still fails, giveup with feof==0
			   so script can retry if desired */

		}
		if ret < 0 {
			if errno == EAGAIN || errno == EAGAIN {

				/* Not an error. */

				ret = 0

				/* Not an error. */

			} else if errno == EINTR {

			} else {
				core.PhpErrorDocref(nil, 1<<3, "read of %zu bytes failed with errno=%d %s", count, errno, strerror(errno))

				/* TODO: Remove this special-case? */

				if errno != EBADF {
					stream.eof = 1
				}

				/* TODO: Remove this special-case? */

			}
		} else if ret == 0 {
			stream.eof = 1
		}
	} else {
		if data.GetIsSeekable() && data.GetLastOp() == 'w' {
			fseek(data.GetFile(), 0, SEEK_CUR)
		}
		data.SetLastOp('r')
		ret = fread(buf, 1, count, data.GetFile())
		stream.eof = feof(data.GetFile())
	}
	return ret
}
func PhpStdiopClose(stream *core.PhpStream, close_handle int) int {
	var ret int
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.abstract)
	assert(data != nil)
	if data.GetLastMappedAddr() != nil {
		munmap(data.GetLastMappedAddr(), data.GetLastMappedLen())
		data.SetLastMappedAddr(nil)
	}
	if close_handle != 0 {
		if data.GetFile() != nil {
			if data.GetIsProcessPipe() {
				errno = 0
				ret = pclose(data.GetFile())
				if WIFEXITED(ret) {
					ret = WEXITSTATUS(ret)
				}
			} else {
				ret = fclose(data.GetFile())
				data.SetFile(nil)
			}
		} else if data.GetFd() != -1 {
			ret = close(data.GetFd())
			data.SetFd(-1)
		} else {
			return 0
		}
		if data.GetTempName() != nil {
			unlink(data.GetTempName().val)

			/* temporary streams are never persistent */

			zend.ZendStringReleaseEx(data.GetTempName(), 0)
			data.SetTempName(nil)
		}
	} else {
		ret = 0
		data.SetFile(nil)
		data.SetFd(-1)
	}
	g.CondF(stream.is_persistent != 0, func() { return zend.Free(data) }, func() { return zend._efree(data) })
	return ret
}
func PhpStdiopFlush(stream *core.PhpStream) int {
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.abstract)
	assert(data != nil)

	/*
	 * stdio buffers data in user land. By calling fflush(3), this
	 * data is send to the kernel using write(2). fsync'ing is
	 * something completely different.
	 */

	if data.GetFile() != nil {
		return fflush(data.GetFile())
	}
	return 0
}
func PhpStdiopSeek(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffset *zend.ZendOffT) int {
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.abstract)
	var ret int
	assert(data != nil)
	if !(data.GetIsSeekable()) {
		core.PhpErrorDocref(nil, 1<<1, "cannot seek on this stream")
		return -1
	}
	if data.GetFd() >= 0 {
		var result zend.ZendOffT
		result = lseek(data.GetFd(), offset, whence)
		if result == zend_off_t-1 {
			return -1
		}
		*newoffset = result
		return 0
	} else {
		ret = fseek(data.GetFile(), offset, whence)
		*newoffset = ftell(data.GetFile())
		return ret
	}
}
func PhpStdiopCast(stream *core.PhpStream, castas int, ret *any) int {
	var fd core.PhpSocketT
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.abstract)
	assert(data != nil)

	/* as soon as someone touches the stdio layer, buffering may ensue,
	 * so we need to stop using the fd directly in that case */

	switch castas {
	case 0:
		if ret != nil {
			if data.GetFile() == nil {

				/* we were opened as a plain file descriptor, so we
				 * need fdopen now */

				var fixed_mode []byte
				PhpStreamModeSanitizeFdopenFopencookie(stream, fixed_mode)
				data.SetFile(fdopen(data.GetFd(), fixed_mode))
				if data.GetFile() == nil {
					return zend.FAILURE
				}
			}
			*((**FILE)(ret)) = data.GetFile()
			data.SetFd(-1)
		}
		return zend.SUCCESS
	case 3:
		if data.GetFile() != nil {
			fd = fileno(data.GetFile())
		} else {
			fd = data.GetFd()
		}
		if -1 == fd {
			return zend.FAILURE
		}
		if ret != nil {
			*((*core.PhpSocketT)(ret)) = fd
		}
		return zend.SUCCESS
	case 1:
		if data.GetFile() != nil {
			fd = fileno(data.GetFile())
		} else {
			fd = data.GetFd()
		}
		if -1 == fd {
			return zend.FAILURE
		}
		if data.GetFile() != nil {
			fflush(data.GetFile())
		}
		if ret != nil {
			*((*core.PhpSocketT)(ret)) = fd
		}
		return zend.SUCCESS
	default:
		return zend.FAILURE
	}

	/* as soon as someone touches the stdio layer, buffering may ensue,
	 * so we need to stop using the fd directly in that case */
}
func PhpStdiopStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	var ret int
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.abstract)
	assert(data != nil)
	if g.Assign(&ret, DoFstat(data, 1)) == 0 {
		memcpy(&ssb.sb, &data.sb, g.SizeOf("ssb -> sb"))
	}
	return ret
}
func PhpStdiopSetOption(stream *core.PhpStream, option int, value int, ptrparam any) int {
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.abstract)
	var size int
	var fd int
	if data.GetFile() != nil {
		fd = fileno(data.GetFile())
	} else {
		fd = data.GetFd()
	}
	switch option {
	case 1:
		if fd == -1 {
			return -1
		}
		return -1
	case 3:
		if data.GetFile() == nil {
			return -1
		}
		if ptrparam {
			size = *((*int)(ptrparam))
		} else {
			size = BUFSIZ
		}
		switch value {
		case 0:
			return setvbuf(data.GetFile(), nil, _IONBF, 0)
		case 1:
			return setvbuf(data.GetFile(), nil, _IOLBF, size)
		case 2:
			return setvbuf(data.GetFile(), nil, _IOFBF, size)
		default:
			return -1
		}
		break
	case 6:
		if fd == -1 {
			return -1
		}
		if zend.ZendUintptrT(ptrparam == 1) != 0 {
			return 0
		}
		if !(flock(fd, value)) {
			data.SetLockFlag(value)
			return 0
		} else {
			return -1
		}
		break
	case 9:
		var range_ *PhpStreamMmapRange = (*PhpStreamMmapRange)(ptrparam)
		var prot int
		var flags int
		switch value {
		case PHP_STREAM_MMAP_SUPPORTED:
			if fd == -1 {
				return -1
			} else {
				return 0
			}
		case PHP_STREAM_MMAP_MAP_RANGE:
			if DoFstat(data, 1) != 0 {
				return -1
			}
			if range_.GetOffset() > data.sb.st_size {
				range_.SetOffset(data.sb.st_size)
			}
			if range_.GetLength() == 0 || range_.GetLength() > data.sb.st_size-range_.GetOffset() {
				range_.SetLength(data.sb.st_size - range_.GetOffset())
			}
			switch range_.GetMode() {
			case PHP_STREAM_MAP_MODE_READONLY:
				prot = PROT_READ
				flags = MAP_PRIVATE
				break
			case PHP_STREAM_MAP_MODE_READWRITE:
				prot = PROT_READ | PROT_WRITE
				flags = MAP_PRIVATE
				break
			case PHP_STREAM_MAP_MODE_SHARED_READONLY:
				prot = PROT_READ
				flags = MAP_SHARED
				break
			case PHP_STREAM_MAP_MODE_SHARED_READWRITE:
				prot = PROT_READ | PROT_WRITE
				flags = MAP_SHARED
				break
			default:
				return -1
			}
			range_.SetMapped((*byte)(mmap(nil, range_.GetLength(), prot, flags, fd, range_.GetOffset())))
			if range_.GetMapped() == (*byte)(any(-1)) {
				range_.SetMapped(nil)
				return -1
			}

			/* remember the mapping */

			data.SetLastMappedAddr(range_.GetMapped())
			data.SetLastMappedLen(range_.GetLength())
			return 0
		case PHP_STREAM_MMAP_UNMAP:
			if data.GetLastMappedAddr() != nil {
				munmap(data.GetLastMappedAddr(), data.GetLastMappedLen())
				data.SetLastMappedAddr(nil)
				return 0
			}
			return -1
		}
		return -2
	case 10:
		switch value {
		case 0:
			if fd == -1 {
				return -1
			} else {
				return 0
			}
		case 1:
			var new_size ptrdiff_t = *((*ptrdiff_t)(ptrparam))
			if new_size < 0 {
				return -1
			}
			if ftruncate(fd, new_size) == 0 {
				return 0
			} else {
				return -1
			}
		}
	case 11:
		if fd == -1 {
			return -1
		}
		return -1
	default:
		return -2
	}
}

/* This should be "const", but phpdbg overwrite it */

var PhpStreamStdioOps core.PhpStreamOps = core.PhpStreamOps{PhpStdiopWrite, PhpStdiopRead, PhpStdiopClose, PhpStdiopFlush, "STDIO", PhpStdiopSeek, PhpStdiopCast, PhpStdiopStat, PhpStdiopSetOption}

/* }}} */

func PhpPlainFilesDirstreamRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var dir *DIR = (*DIR)(stream.abstract)
	var result *__struct__dirent
	var ent *core.PhpStreamDirent = (*core.PhpStreamDirent)(buf)

	/* avoid problems if someone mis-uses the stream */

	if count != g.SizeOf("php_stream_dirent") {
		return -1
	}
	result = readdir(dir)
	if result != nil {
		var php_str_len int
		if strlen(result.d_name) >= g.SizeOf("ent -> d_name") {
			php_str_len = g.SizeOf("ent -> d_name") - 1
		} else {
			php_str_len = strlen(result.d_name)
		}
		memcpy(ent.d_name, result.d_name, php_str_len)
		ent.d_name[php_str_len] = '0'
		return g.SizeOf("php_stream_dirent")
	}
	return 0
}
func PhpPlainFilesDirstreamClose(stream *core.PhpStream, close_handle int) int {
	return closedir((*DIR)(stream.abstract))
}
func PhpPlainFilesDirstreamRewind(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	rewinddir((*DIR)(stream.abstract))
	return 0
}

var PhpPlainFilesDirstreamOps core.PhpStreamOps = core.PhpStreamOps{nil, PhpPlainFilesDirstreamRead, PhpPlainFilesDirstreamClose, nil, "dir", PhpPlainFilesDirstreamRewind, nil, nil, nil}

func PhpPlainFilesDirOpener(wrapper *core.PhpStreamWrapper, path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
	var dir *DIR = nil
	var stream *core.PhpStream = nil
	if (options & 0x1000) != 0 {
		return PhpGlobStreamWrapper.wops.dir_opener((*core.PhpStreamWrapper)(&PhpGlobStreamWrapper), path, mode, options, opened_path, context)
	}
	if (options&0x400) == 0 && core.PhpCheckOpenBasedir(path) != 0 {
		return nil
	}
	dir = opendir(path)
	if dir != nil {
		stream = _phpStreamAlloc(&PhpPlainFilesDirstreamOps, dir, 0, mode)
		if stream == nil {
			closedir(dir)
		}
	}
	return stream
}

/* }}} */

func _phpStreamFopen(filename *byte, mode *byte, opened_path **zend.ZendString, options int) *core.PhpStream {
	var realpath []byte
	var open_flags int
	var fd int
	var ret *core.PhpStream
	var persistent int = options & 0x800
	var persistent_id *byte = nil
	if zend.FAILURE == PhpStreamParseFopenModes(mode, &open_flags) {
		PhpStreamWrapperLogError(&PhpPlainFilesWrapper, options, "`%s' is not a valid mode for fopen", mode)
		return nil
	}
	if (options & 0x4000) != 0 {
		strlcpy(realpath, filename, g.SizeOf("realpath"))
	} else {
		if core.ExpandFilepath(filename, realpath) == nil {
			return nil
		}
	}
	if persistent != 0 {
		zend.ZendSpprintf(&persistent_id, 0, "streams_stdio_%d_%s", open_flags, realpath)
		switch PhpStreamFromPersistentId(persistent_id, &ret) {
		case 0:
			if opened_path != nil {

				//TODO: avoid reallocation???

				*opened_path = zend.ZendStringInit(realpath, strlen(realpath), 0)

				//TODO: avoid reallocation???

			}
		case 1:
			zend._efree(persistent_id)
			return ret
		}
	}
	fd = open(realpath, open_flags, 0666)
	if fd != -1 {
		if (options & 0x80) != 0 {
			ret = _phpStreamFopenFromFdInt(fd, mode, persistent_id)
		} else {
			ret = _phpStreamFopenFromFd(fd, mode, persistent_id)
		}
		if ret != nil {
			if opened_path != nil {
				*opened_path = zend.ZendStringInit(realpath, strlen(realpath), 0)
			}
			if persistent_id != nil {
				zend._efree(persistent_id)
			}

			/* WIN32 always set ISREG flag */

			/* sanity checks for include/require.
			 * We check these after opening the stream, so that we save
			 * on fstat() syscalls */

			if (options & 0x80) != 0 {
				var self *PhpStdioStreamData = (*PhpStdioStreamData)(ret.abstract)
				var r int
				r = DoFstat(self, 0)
				if r == 0 && (self.sb.st_mode&S_IFMT) != S_IFREG {
					if opened_path != nil {
						zend.ZendStringReleaseEx(*opened_path, 0)
						*opened_path = nil
					}
					_phpStreamFree(ret, 1|2)
					return nil
				}

				/* Make sure the fstat result is reused when we later try to get the
				 * file size. */

				self.SetNoForcedFstat(1)

				/* Make sure the fstat result is reused when we later try to get the
				 * file size. */

			}
			if (options & 0x8000) != 0 {
				var self *PhpStdioStreamData = (*PhpStdioStreamData)(ret.abstract)
				self.SetIsPipeBlocking(1)
			}
			return ret
		}
		close(fd)
	}
	if persistent_id != nil {
		zend._efree(persistent_id)
	}
	return nil
}

/* }}} */

func PhpPlainFilesStreamOpener(wrapper *core.PhpStreamWrapper, path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
	if (options&0x400) == 0 && core.PhpCheckOpenBasedir(path) != 0 {
		return nil
	}
	return _phpStreamFopen(path, mode, opened_path, options)
}
func PhpPlainFilesUrlStater(wrapper *core.PhpStreamWrapper, url *byte, flags int, ssb *core.PhpStreamStatbuf, context *core.PhpStreamContext) int {
	if strncasecmp(url, "file://", g.SizeOf("\"file://\"")-1) == 0 {
		url += g.SizeOf("\"file://\"") - 1
	}
	if core.PhpCheckOpenBasedirEx(url, g.Cond((flags&2) != 0, 0, 1)) != 0 {
		return -1
	}
	if (flags & 1) != 0 {
		return lstat(url, &ssb.sb)
	} else {
		return stat(url, &ssb.sb)
	}
}
func PhpPlainFilesUnlink(wrapper *core.PhpStreamWrapper, url *byte, options int, context *core.PhpStreamContext) int {
	var ret int
	if strncasecmp(url, "file://", g.SizeOf("\"file://\"")-1) == 0 {
		url += g.SizeOf("\"file://\"") - 1
	}
	if core.PhpCheckOpenBasedir(url) != 0 {
		return 0
	}
	ret = unlink(url)
	if ret == -1 {
		if (options & 0x8) != 0 {
			core.PhpErrorDocref1(nil, url, 1<<1, "%s", strerror(errno))
		}
		return 0
	}

	/* Clear stat cache (and realpath cache) */

	standard.PhpClearStatCache(1, nil, 0)
	return 1
}
func PhpPlainFilesRename(wrapper *core.PhpStreamWrapper, url_from *byte, url_to *byte, options int, context *core.PhpStreamContext) int {
	var ret int
	if url_from == nil || url_to == nil {
		return 0
	}
	if strncasecmp(url_from, "file://", g.SizeOf("\"file://\"")-1) == 0 {
		url_from += g.SizeOf("\"file://\"") - 1
	}
	if strncasecmp(url_to, "file://", g.SizeOf("\"file://\"")-1) == 0 {
		url_to += g.SizeOf("\"file://\"") - 1
	}
	if core.PhpCheckOpenBasedir(url_from) != 0 || core.PhpCheckOpenBasedir(url_to) != 0 {
		return 0
	}
	ret = rename(url_from, url_to)
	if ret == -1 {
		core.PhpErrorDocref2(nil, url_from, url_to, 1<<1, "%s", strerror(errno))
		return 0
	}

	/* Clear stat cache (and realpath cache) */

	standard.PhpClearStatCache(1, nil, 0)
	return 1
}
func PhpPlainFilesMkdir(wrapper *core.PhpStreamWrapper, dir *byte, mode int, options int, context *core.PhpStreamContext) int {
	var ret int
	var recursive int = options & 1
	var p *byte
	if strncasecmp(dir, "file://", g.SizeOf("\"file://\"")-1) == 0 {
		dir += g.SizeOf("\"file://\"") - 1
	}
	if recursive == 0 {
		ret = standard.PhpMkdir(dir, mode)
	} else {

		/* we look for directory separator from the end of string, thus hopefuly reducing our work load */

		var e *byte
		var sb zend.ZendStatT
		var dir_len int = strlen(dir)
		var offset int = 0
		var buf []byte
		if core.ExpandFilepathWithMode(dir, buf, nil, 0, 0) == nil {
			core.PhpErrorDocref(nil, 1<<1, "Invalid path")
			return 0
		}
		e = buf + strlen(buf)
		if g.Assign(&p, memchr(buf, '/', dir_len)) {
			offset = p - buf + 1
		}
		if p != nil && dir_len == 1 {

		} else {

			/* find a top level directory we need to create */

			for g.Assign(&p, strrchr(buf+offset, '/')) || offset != 1 && g.Assign(&p, strrchr(buf, '/')) {
				var n int = 0
				*p = '0'
				for p > buf && (*(p - 1)) == '/' {
					n++
					p--
					*p = '0'
				}
				if stat(buf, &sb) == 0 {
					for true {
						*p = '/'
						if n == 0 {
							break
						}
						n--
						p++
					}
					break
				}
			}

			/* find a top level directory we need to create */

		}
		if p == buf {
			ret = standard.PhpMkdir(dir, mode)
		} else if !(g.Assign(&ret, standard.PhpMkdir(buf, mode))) {
			if p == nil {
				p = buf
			}

			/* create any needed directories if the creation of the 1st directory worked */

			for g.PreInc(&p) != e {
				if (*p) == '0' {
					*p = '/'
					if (*(p + 1)) != '0' && g.Assign(&ret, mkdir(buf, mode_t(mode))) < 0 {
						if (options & 0x8) != 0 {
							core.PhpErrorDocref(nil, 1<<1, "%s", strerror(errno))
						}
						break
					}
				}
			}

			/* create any needed directories if the creation of the 1st directory worked */

		}
	}
	if ret < 0 {

		/* Failure */

		return 0

		/* Failure */

	} else {

		/* Success */

		return 1

		/* Success */

	}
}
func PhpPlainFilesRmdir(wrapper *core.PhpStreamWrapper, url *byte, options int, context *core.PhpStreamContext) int {
	if strncasecmp(url, "file://", g.SizeOf("\"file://\"")-1) == 0 {
		url += g.SizeOf("\"file://\"") - 1
	}
	if core.PhpCheckOpenBasedir(url) != 0 {
		return 0
	}
	if rmdir(url) < 0 {
		core.PhpErrorDocref1(nil, url, 1<<1, "%s", strerror(errno))
		return 0
	}

	/* Clear stat cache (and realpath cache) */

	standard.PhpClearStatCache(1, nil, 0)
	return 1
}
func PhpPlainFilesMetadata(wrapper *core.PhpStreamWrapper, url *byte, option int, value any, context *core.PhpStreamContext) int {
	var newtime *__struct__utimbuf
	var uid uid_t
	var gid gid_t
	var mode mode_t
	var ret int = 0
	if strncasecmp(url, "file://", g.SizeOf("\"file://\"")-1) == 0 {
		url += g.SizeOf("\"file://\"") - 1
	}
	if core.PhpCheckOpenBasedir(url) != 0 {
		return 0
	}
	switch option {
	case 1:
		newtime = (*__struct__utimbuf)(value)
		if access(url, F_OK) != 0 {
			var file *FILE = fopen(url, "w")
			if file == nil {
				core.PhpErrorDocref1(nil, url, 1<<1, "Unable to create file %s because %s", url, strerror(errno))
				return 0
			}
			fclose(file)
		}
		ret = utime(url, newtime)
		break
	case 2:

	case 3:
		if option == 2 {
			if PhpGetUidByName((*byte)(value), &uid) != zend.SUCCESS {
				core.PhpErrorDocref1(nil, url, 1<<1, "Unable to find uid for %s", (*byte)(value))
				return 0
			}
		} else {
			uid = uid_t * (*long)(value)
		}
		ret = chown(url, uid, -1)
		break
	case 5:

	case 4:
		if option == 4 {
			if PhpGetGidByName((*byte)(value), &gid) != zend.SUCCESS {
				core.PhpErrorDocref1(nil, url, 1<<1, "Unable to find gid for %s", (*byte)(value))
				return 0
			}
		} else {
			gid = gid_t * (*long)(value)
		}
		ret = chown(url, -1, gid)
		break
	case 6:
		mode = mode_t * (*zend.ZendLong)(value)
		ret = chmod(url, mode)
		break
	default:
		core.PhpErrorDocref1(nil, url, 1<<1, "Unknown option %d for stream_metadata", option)
		return 0
	}
	if ret == -1 {
		core.PhpErrorDocref1(nil, url, 1<<1, "Operation failed: %s", strerror(errno))
		return 0
	}
	standard.PhpClearStatCache(0, nil, 0)
	return 1
}

var PhpPlainFilesWrapperOps core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{PhpPlainFilesStreamOpener, nil, nil, PhpPlainFilesUrlStater, PhpPlainFilesDirOpener, "plainfile", PhpPlainFilesUnlink, PhpPlainFilesRename, PhpPlainFilesMkdir, PhpPlainFilesRmdir, PhpPlainFilesMetadata}

/* TODO: We have to make php_plain_files_wrapper writable to support SWOOLE */

var PhpPlainFilesWrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&PhpPlainFilesWrapperOps, nil, 0}

/* {{{ php_stream_fopen_with_path */

func _phpStreamFopenWithPath(filename *byte, mode *byte, path *byte, opened_path **zend.ZendString, options int) *core.PhpStream {
	/* code ripped off from fopen_wrappers.c */

	var pathbuf *byte
	var end *byte
	var ptr *byte
	var trypath []byte
	var stream *core.PhpStream
	var filename_length int
	var exec_filename *zend.ZendString
	if opened_path != nil {
		*opened_path = nil
	}
	if filename == nil {
		return nil
	}
	filename_length = strlen(filename)
	void(filename_length)

	/* Relative path open */

	if (*filename) == '.' && (filename[1] == '/' || filename[1] == '.') {

		/* further checks, we could have ....... filenames */

		ptr = filename + 1
		if (*ptr) == '.' {
			for (*(g.PreInc(&ptr))) == '.' {

			}
			if (*ptr) != '/' {
				goto not_relative_path
			}
		}
		if (options&0x400) == 0 && core.PhpCheckOpenBasedir(filename) != 0 {
			return nil
		}
		return _phpStreamFopen(filename, mode, opened_path, options)
	}
not_relative_path:

	/* Absolute path open */

	if filename[0] == '/' {
		if (options&0x400) == 0 && core.PhpCheckOpenBasedir(filename) != 0 {
			return nil
		}
		return _phpStreamFopen(filename, mode, opened_path, options)
	}
	if path == nil || !(*path) {
		return _phpStreamFopen(filename, mode, opened_path, options)
	}

	/* check in provided path */

	if zend.ZendIsExecuting() != 0 && g.Assign(&exec_filename, zend.ZendGetExecutedFilenameEx()) != nil {
		var exec_fname *byte = exec_filename.val
		var exec_fname_length int = exec_filename.len_
		for g.PreDec(&exec_fname_length) < SIZE_MAX && exec_fname[exec_fname_length] != '/' {

		}
		if exec_fname_length <= 0 {

			/* no path */

			pathbuf = zend._estrdup(path)

			/* no path */

		} else {
			var path_length int = strlen(path)
			pathbuf = (*byte)(zend._emalloc(exec_fname_length + path_length + 1 + 1))
			memcpy(pathbuf, path, path_length)
			pathbuf[path_length] = ':'
			memcpy(pathbuf+path_length+1, exec_fname, exec_fname_length)
			pathbuf[path_length+exec_fname_length+1] = '0'
		}
	} else {
		pathbuf = zend._estrdup(path)
	}
	ptr = pathbuf
	for ptr != nil && (*ptr) {
		end = strchr(ptr, ':')
		if end != nil {
			*end = '0'
			end++
		}
		if (*ptr) == '0' {
			goto stream_skip
		}
		if core.ApPhpSnprintf(trypath, 256, "%s/%s", ptr, filename) >= 256 {
			core.PhpErrorDocref(nil, 1<<3, "%s/%s path was truncated to %d", ptr, filename, 256)
		}
		if (options&0x400) == 0 && core.PhpCheckOpenBasedirEx(trypath, 0) != 0 {
			goto stream_skip
		}
		stream = _phpStreamFopen(trypath, mode, opened_path, options)
		if stream != nil {
			zend._efree(pathbuf)
			return stream
		}
	stream_skip:
		ptr = end
	}
	zend._efree(pathbuf)
	return nil
}

/* }}} */
