package streams

import (
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"os"
)

func PhpStreamFopenFromFdIntRel(fd int, mode *byte, persistent_id *byte) *core.PhpStream {
	return _phpStreamFopenFromFdInt(fd, mode, persistent_id)
}
func PhpStreamFopenFromFileIntRel(file *r.File, mode *byte) *core.PhpStream {
	return _phpStreamFopenFromFileInt(file, mode)
}
func PLAIN_WRAP_BUF_SIZE(st int) int { return st }
func PhpStreamParseFopenModes(mode *byte, open_flags *int) int {
	var flags int
	switch mode[0] {
	case 'r':
		flags = 0
	case 'w':
		flags = O_TRUNC | O_CREAT
	case 'a':
		flags = O_CREAT | O_APPEND
	case 'x':
		flags = O_CREAT | O_EXCL
	case 'c':
		flags = O_CREAT
	default:

		/* unknown mode */

		return types.FAILURE
	}
	if strchr(mode, '+') {
		flags |= O_RDWR
	} else if flags != 0 {
		flags |= O_WRONLY
	} else {
		flags |= O_RDONLY
	}
	*open_flags = flags
	return types.SUCCESS
}
func PHP_STDIOP_GET_FD(anfd int, data *PhpStdioStreamData) int {
	if data.GetFile() != nil {
		anfd = fileno(data.GetFile())
	} else {
		anfd = data.GetFd()
	}
	return anfd
}
func DoFstat(d *PhpStdioStreamData, force int) int {
	if !(d.GetCachedFstat()) || force != 0 && !(d.GetNoForcedFstat()) {
		var fd int
		var r int
		PHP_STDIOP_GET_FD(fd, d)
		r = zend.ZendFstat(fd, d.GetSb())
		d.SetCachedFstat(r == 0)
		return r
	}
	return 0
}
func _phpStreamFopenFromFdInt(fd int, mode *byte, persistent_id *byte) *core.PhpStream {
	var self *PhpStdioStreamData
	self = PemallocRelOrig(b.SizeOf("* self"))
	memset(self, 0, b.SizeOf("* self"))
	self.SetFile(nil)
	self.SetIsSeekable(1)
	self.SetIsPipe(0)
	self.SetLockFlag(LOCK_UN)
	self.SetIsProcessPipe(0)
	self.SetTempName(nil)
	self.SetFd(fd)
	return core.PhpStreamAllocRel(&PhpStreamStdioOps, self, persistent_id, mode)
}
func _phpStreamFopenFromFileInt(file *r.File, mode *byte) *core.PhpStream {
	var self *PhpStdioStreamData
	self = EmallocRelOrig(b.SizeOf("* self"))
	memset(self, 0, b.SizeOf("* self"))
	self.SetFile(file)
	self.SetIsSeekable(1)
	self.SetIsPipe(0)
	self.SetLockFlag(LOCK_UN)
	self.SetIsProcessPipe(0)
	self.SetTempName(nil)
	self.SetFd(fileno(file))
	return core.PhpStreamAllocRel(&PhpStreamStdioOps, self, 0, mode)
}
func _phpStreamFopenTemporaryFile(dir *byte, pfx string, opened_path_ptr **types.String) *core.PhpStream {
	var opened_path *types.String = nil
	var fd int
	fd = core.PhpOpenTemporaryFd(dir, pfx, &opened_path)
	if fd != -1 {
		var stream *core.PhpStream
		if opened_path_ptr != nil {
			*opened_path_ptr = opened_path
		}
		stream = PhpStreamFopenFromFdIntRel(fd, "r+b", nil)
		if stream != nil {
			var self *PhpStdioStreamData = (*PhpStdioStreamData)(stream.GetAbstract())
			stream.SetWrapper((*core.PhpStreamWrapper)(&PhpPlainFilesWrapper))
			stream.SetOrigPath(zend.Estrndup(opened_path.GetVal(), opened_path.GetLen()))
			self.SetTempName(opened_path)
			self.SetLockFlag(LOCK_UN)
			return stream
		}
		close(fd)
		core.PhpErrorDocref(nil, faults.E_WARNING, "unable to allocate stream")
		return nil
	}
	return nil
}
func _phpStreamFopenTmpfile(dummy int) *core.PhpStream {
	return PhpStreamFopenTemporaryFile(nil, "php", nil)
}
func DetectIsSeekable(self *PhpStdioStreamData) {}
func _phpStreamFopenFromFd(fd int, mode *byte, persistent_id *byte) *core.PhpStream {
	var stream *core.PhpStream = PhpStreamFopenFromFdIntRel(fd, mode, persistent_id)
	if stream != nil {
		var self *PhpStdioStreamData = (*PhpStdioStreamData)(stream.GetAbstract())
		DetectIsSeekable(self)
		if !(self.GetIsSeekable()) {
			stream.AddFlags(core.PHP_STREAM_FLAG_NO_SEEK)
			stream.SetPosition(-1)
		} else {
			stream.SetPosition(zend.ZendLseek(self.GetFd(), 0, r.SEEK_CUR))
		}
	}
	return stream
}
func _phpStreamFopenFromFile(file *r.File, mode *byte) *core.PhpStream {
	var stream *core.PhpStream = PhpStreamFopenFromFileIntRel(file, mode)
	if stream != nil {
		var self *PhpStdioStreamData = (*PhpStdioStreamData)(stream.GetAbstract())
		DetectIsSeekable(self)
		if !(self.GetIsSeekable()) {
			stream.AddFlags(core.PHP_STREAM_FLAG_NO_SEEK)
			stream.SetPosition(-1)
		} else {
			stream.SetPosition(zend.ZendFtell(file))
		}
	}
	return stream
}
func _phpStreamFopenFromPipe(file *r.File, mode *byte) *core.PhpStream {
	var self *PhpStdioStreamData
	var stream *core.PhpStream
	self = EmallocRelOrig(b.SizeOf("* self"))
	memset(self, 0, b.SizeOf("* self"))
	self.SetFile(file)
	self.SetIsSeekable(0)
	self.SetIsPipe(1)
	self.SetLockFlag(LOCK_UN)
	self.SetIsProcessPipe(1)
	self.SetFd(fileno(file))
	self.SetTempName(nil)
	stream = core.PhpStreamAllocRel(&PhpStreamStdioOps, self, 0, mode)
	stream.AddFlags(core.PHP_STREAM_FLAG_NO_SEEK)
	return stream
}
func PhpStdiopWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.GetAbstract())
	b.Assert(data != nil)
	if data.GetFd() >= 0 {
		var bytes_written ssize_t = write(data.GetFd(), buf, count)
		if bytes_written < 0 {
			if errno == core.EWOULDBLOCK || errno == EAGAIN {
				return 0
			}
			if errno == EINTR {

				/* TODO: Should this be treated as a proper error or not? */

				return bytes_written

				/* TODO: Should this be treated as a proper error or not? */

			}
			core.PhpErrorDocref(nil, faults.E_NOTICE, "write of %zu bytes failed with errno=%d %s", count, errno, strerror(errno))
		}
		return bytes_written
	} else {
		if data.GetIsSeekable() && data.GetLastOp() == 'r' {
			zend.ZendFseek(data.GetFile(), 0, r.SEEK_CUR)
		}
		data.SetLastOp('w')
		size, err := data.GetFile().Write(buf)
		if err != nil {
			return r.EOF
		}
		return size
	}
}
func PhpStdiopRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.GetAbstract())
	var ret ssize_t
	b.Assert(data != nil)
	if data.GetFd() >= 0 {
		ret = read(data.GetFd(), buf, PLAIN_WRAP_BUF_SIZE(count))
		if ret == size_t-1 && errno == EINTR {

			/* read was interrupted, retry once,
			   If read still fails, giveup with feof==0
			   so script can retry if desired */

			ret = read(data.GetFd(), buf, PLAIN_WRAP_BUF_SIZE(count))

			/* read was interrupted, retry once,
			   If read still fails, giveup with feof==0
			   so script can retry if desired */

		}
		if ret < 0 {
			if errno == core.EWOULDBLOCK || errno == EAGAIN {

				/* Not an error. */

				ret = 0

				/* Not an error. */

			} else if errno == EINTR {

			} else {
				core.PhpErrorDocref(nil, faults.E_NOTICE, "read of %zu bytes failed with errno=%d %s", count, errno, strerror(errno))

				/* TODO: Remove this special-case? */

				if errno != EBADF {
					stream.SetEof(1)
				}

				/* TODO: Remove this special-case? */

			}
		} else if ret == 0 {
			stream.SetEof(1)
		}
	} else {
		if data.GetIsSeekable() && data.GetLastOp() == 'w' {
			zend.ZendFseek(data.GetFile(), 0, r.SEEK_CUR)
		}
		data.SetLastOp('r')
		ret, _ = data.GetFile().Read(buf[:count])
		stream.SetEof(data.GetFile().Eof())
	}
	return ret
}
func PhpStdiopClose(stream *core.PhpStream, close_handle int) int {
	var ret int
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.GetAbstract())
	b.Assert(data != nil)
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
				ret = data.GetFile().Close()
				data.SetFile(nil)
			}
		} else if data.GetFd() != -1 {
			ret = close(data.GetFd())
			data.SetFd(-1)
		} else {
			return 0
		}
		if data.GetTempName() != nil {
			unlink(data.GetTempName().GetVal())

			/* temporary streams are never persistent */

			// types.ZendStringReleaseEx(data.GetTempName(), 0)
			data.SetTempName(nil)
		}
	} else {
		ret = 0
		data.SetFile(nil)
		data.SetFd(-1)
	}
	zend.Pefree(data, stream.GetIsPersistent())
	return ret
}
func PhpStdiopFlush(stream *core.PhpStream) int {
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.GetAbstract())
	b.Assert(data != nil)

	/*
	 * stdio buffers data in user land. By calling fflush(3), this
	 * data is send to the kernel using write(2). fsync'ing is
	 * something completely different.
	 */

	if data.GetFile() != nil {
		return data.GetFile().Flush()
	}
	return 0
}
func PhpStdiopSeek(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffset *zend.ZendOffT) int {
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.GetAbstract())
	var ret int
	b.Assert(data != nil)
	if !(data.GetIsSeekable()) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "cannot seek on this stream")
		return -1
	}
	if data.GetFd() >= 0 {
		var result zend.ZendOffT
		result = zend.ZendLseek(data.GetFd(), offset, whence)
		if result == zend_off_t-1 {
			return -1
		}
		*newoffset = result
		return 0
	} else {
		ret = zend.ZendFseek(data.GetFile(), offset, whence)
		*newoffset = zend.ZendFtell(data.GetFile())
		return ret
	}
}
func PhpStdiopCast(stream *core.PhpStream, castas int, ret *any) int {
	var fd core.PhpSocketT
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.GetAbstract())
	b.Assert(data != nil)

	/* as soon as someone touches the stdio layer, buffering may ensue,
	 * so we need to stop using the fd directly in that case */

	switch castas {
	case core.PHP_STREAM_AS_STDIO:
		if ret != nil {
			if data.GetFile() == nil {

				/* we were opened as a plain file descriptor, so we
				 * need fdopen now */

				var fixed_mode []byte
				PhpStreamModeSanitizeFdopenFopencookie(stream, fixed_mode)
				data.SetFile(fdopen(data.GetFd(), fixed_mode))
				if data.GetFile() == nil {
					return types.FAILURE
				}
			}
			*((**r.File)(ret)) = data.GetFile()
			data.SetFd(core.SOCK_ERR)
		}
		return types.SUCCESS
	case core.PHP_STREAM_AS_FD_FOR_SELECT:
		PHP_STDIOP_GET_FD(fd, data)
		if core.SOCK_ERR == fd {
			return types.FAILURE
		}
		if ret != nil {
			*((*core.PhpSocketT)(ret)) = fd
		}
		return types.SUCCESS
	case core.PHP_STREAM_AS_FD:
		PHP_STDIOP_GET_FD(fd, data)
		if core.SOCK_ERR == fd {
			return types.FAILURE
		}
		if data.GetFile() != nil {
			data.GetFile().Flush()
		}
		if ret != nil {
			*((*core.PhpSocketT)(ret)) = fd
		}
		return types.SUCCESS
	default:
		return types.FAILURE
	}

	/* as soon as someone touches the stdio layer, buffering may ensue,
	 * so we need to stop using the fd directly in that case */
}
func PhpStdiopStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	var ret int
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.GetAbstract())
	b.Assert(data != nil)
	if b.Assign(&ret, DoFstat(data, 1)) == 0 {
		memcpy(ssb.GetSb(), data.GetSb(), b.SizeOf("ssb -> sb"))
	}
	return ret
}
func PhpStdiopSetOption(stream *core.PhpStream, option int, value int, ptrparam any) int {
	var data *PhpStdioStreamData = (*PhpStdioStreamData)(stream.GetAbstract())
	var size int
	var fd int
	PHP_STDIOP_GET_FD(fd, data)
	switch option {
	case core.PHP_STREAM_OPTION_BLOCKING:
		if fd == -1 {
			return -1
		}
		return -1
	case core.PHP_STREAM_OPTION_WRITE_BUFFER:
		if data.GetFile() == nil {
			return -1
		}
		if ptrparam {
			size = *((*int)(ptrparam))
		} else {
			size = r.BUFSIZ
		}
		switch value {
		case core.PHP_STREAM_BUFFER_NONE:
			return r.Setvbuf(data.GetFile(), nil, _IONBF, 0)
		case core.PHP_STREAM_BUFFER_LINE:
			return r.Setvbuf(data.GetFile(), nil, _IOLBF, size)
		case core.PHP_STREAM_BUFFER_FULL:
			return r.Setvbuf(data.GetFile(), nil, _IOFBF, size)
		default:
			return -1
		}
	case core.PHP_STREAM_OPTION_LOCKING:
		if fd == -1 {
			return -1
		}
		if types.ZendUintptrT(ptrparam == core.PHP_STREAM_LOCK_SUPPORTED) != 0 {
			return 0
		}
		if !(flock(fd, value)) {
			data.SetLockFlag(value)
			return 0
		} else {
			return -1
		}
	case core.PHP_STREAM_OPTION_MMAP_API:
		var range_ *PhpStreamMmapRange = (*PhpStreamMmapRange)(ptrparam)
		var prot int
		var flags int
		switch value {
		case PHP_STREAM_MMAP_SUPPORTED:
			if fd == -1 {
				return core.PHP_STREAM_OPTION_RETURN_ERR
			} else {
				return core.PHP_STREAM_OPTION_RETURN_OK
			}
			fallthrough
		case PHP_STREAM_MMAP_MAP_RANGE:
			if DoFstat(data, 1) != 0 {
				return core.PHP_STREAM_OPTION_RETURN_ERR
			}
			if range_.GetOffset() > data.GetSb().st_size {
				range_.SetOffset(data.GetSb().st_size)
			}
			if range_.GetLength() == 0 || range_.GetLength() > data.GetSb().st_size-range_.GetOffset() {
				range_.SetLength(data.GetSb().st_size - range_.GetOffset())
			}
			switch range_.GetMode() {
			case PHP_STREAM_MAP_MODE_READONLY:
				prot = PROT_READ
				flags = MAP_PRIVATE
			case PHP_STREAM_MAP_MODE_READWRITE:
				prot = PROT_READ | PROT_WRITE
				flags = MAP_PRIVATE
			case PHP_STREAM_MAP_MODE_SHARED_READONLY:
				prot = PROT_READ
				flags = MAP_SHARED
			case PHP_STREAM_MAP_MODE_SHARED_READWRITE:
				prot = PROT_READ | PROT_WRITE
				flags = MAP_SHARED
			default:
				return core.PHP_STREAM_OPTION_RETURN_ERR
			}
			range_.SetMapped((*byte)(mmap(nil, range_.GetLength(), prot, flags, fd, range_.GetOffset())))
			if range_.GetMapped() == (*byte)(MAP_FAILED) {
				range_.SetMapped(nil)
				return core.PHP_STREAM_OPTION_RETURN_ERR
			}

			/* remember the mapping */

			data.SetLastMappedAddr(range_.GetMapped())
			data.SetLastMappedLen(range_.GetLength())
			return core.PHP_STREAM_OPTION_RETURN_OK
		case PHP_STREAM_MMAP_UNMAP:
			if data.GetLastMappedAddr() != nil {
				munmap(data.GetLastMappedAddr(), data.GetLastMappedLen())
				data.SetLastMappedAddr(nil)
				return core.PHP_STREAM_OPTION_RETURN_OK
			}
			return core.PHP_STREAM_OPTION_RETURN_ERR
		}
		return core.PHP_STREAM_OPTION_RETURN_NOTIMPL
	case core.PHP_STREAM_OPTION_TRUNCATE_API:
		switch value {
		case core.PHP_STREAM_TRUNCATE_SUPPORTED:
			if fd == -1 {
				return core.PHP_STREAM_OPTION_RETURN_ERR
			} else {
				return core.PHP_STREAM_OPTION_RETURN_OK
			}
			fallthrough
		case core.PHP_STREAM_TRUNCATE_SET_SIZE:
			var new_size ptrdiff_t = *((*ptrdiff_t)(ptrparam))
			if new_size < 0 {
				return core.PHP_STREAM_OPTION_RETURN_ERR
			}
			if ftruncate(fd, new_size) == 0 {
				return core.PHP_STREAM_OPTION_RETURN_OK
			} else {
				return core.PHP_STREAM_OPTION_RETURN_ERR
			}
		}
		fallthrough
	case core.PHP_STREAM_OPTION_META_DATA_API:
		if fd == -1 {
			return -1
		}
		return -1
	default:
		return core.PHP_STREAM_OPTION_RETURN_NOTIMPL
	}
}
func PhpPlainFilesDirstreamRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var dir *DIR = (*DIR)(stream.GetAbstract())
	var result *__struct__dirent
	var ent *core.PhpStreamDirent = (*core.PhpStreamDirent)(buf)

	/* avoid problems if someone mis-uses the stream */

	if count != b.SizeOf("php_stream_dirent") {
		return -1
	}
	result = readdir(dir)
	if result != nil {
		core.PHP_STRLCPY(ent.GetDName(), result.d_name, b.SizeOf("ent -> d_name"), strlen(result.d_name))
		return b.SizeOf("php_stream_dirent")
	}
	return 0
}
func PhpPlainFilesDirstreamClose(stream *core.PhpStream, close_handle int) int {
	return closedir((*DIR)(stream.GetAbstract()))
}
func PhpPlainFilesDirstreamRewind(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	rewinddir((*DIR)(stream.GetAbstract()))
	return 0
}
func PhpPlainFilesDirOpener(
	wrapper *core.PhpStreamWrapper,
	path *byte,
	mode *byte,
	options int,
	opened_path **types.String,
	context *core.PhpStreamContext,
) *core.PhpStream {
	var dir *DIR = nil
	var stream *core.PhpStream = nil
	if (options & core.STREAM_USE_GLOB_DIR_OPEN) != 0 {
		return PhpGlobStreamWrapper.GetWops().GetDirOpener()((*core.PhpStreamWrapper)(&PhpGlobStreamWrapper), path, mode, options, opened_path, context)
	}
	if (options&core.STREAM_DISABLE_OPEN_BASEDIR) == 0 && core.PhpCheckOpenBasedir(path) != 0 {
		return nil
	}
	dir = zend.VCWD_OPENDIR(path)
	if dir != nil {
		stream = core.PhpStreamAlloc(&PhpPlainFilesDirstreamOps, dir, 0, mode)
		if stream == nil {
			closedir(dir)
		}
	}
	return stream
}
func _phpStreamFopen(filename *byte, mode *byte, opened_path **types.String, options int) *core.PhpStream {
	var realpath []byte
	var open_flags int
	var fd int
	var ret *core.PhpStream
	var persistent int = options & core.STREAM_OPEN_PERSISTENT
	var persistent_id *byte = nil
	if types.FAILURE == PhpStreamParseFopenModes(mode, &open_flags) {
		PhpStreamWrapperLogError(&PhpPlainFilesWrapper, options, "`%s' is not a valid mode for fopen", mode)
		return nil
	}
	if (options & core.STREAM_ASSUME_REALPATH) != 0 {
		strlcpy(realpath, filename, b.SizeOf("realpath"))
	} else {
		if core.ExpandFilepath(filename, realpath) == nil {
			return nil
		}
	}
	if persistent != 0 {
		core.Spprintf(&persistent_id, 0, "streams_stdio_%d_%s", open_flags, realpath)
		switch PhpStreamFromPersistentId(persistent_id, &ret) {
		case core.PHP_STREAM_PERSISTENT_SUCCESS:
			if opened_path != nil {

				//TODO: avoid reallocation???

				*opened_path = types.NewString(realpath)

				//TODO: avoid reallocation???

			}
			fallthrough
		case core.PHP_STREAM_PERSISTENT_FAILURE:
			zend.Efree(persistent_id)
			return ret
		}
	}
	fd = open(realpath, open_flags, 0666)
	if fd != -1 {
		if (options & core.STREAM_OPEN_FOR_INCLUDE) != 0 {
			ret = PhpStreamFopenFromFdIntRel(fd, mode, persistent_id)
		} else {
			ret = core.PhpStreamFopenFromFdRel(fd, mode, persistent_id)
		}
		if ret != nil {
			if opened_path != nil {
				*opened_path = types.NewString(realpath)
			}
			if persistent_id != nil {
				zend.Efree(persistent_id)
			}

			/* WIN32 always set ISREG flag */

			/* sanity checks for include/require.
			 * We check these after opening the stream, so that we save
			 * on fstat() syscalls */

			if (options & core.STREAM_OPEN_FOR_INCLUDE) != 0 {
				var self *PhpStdioStreamData = (*PhpStdioStreamData)(ret.GetAbstract())
				var r int
				r = DoFstat(self, 0)
				if r == 0 && !(zend.S_ISREG(self.GetSb().st_mode)) {
					if opened_path != nil {
						// types.ZendStringReleaseEx(*opened_path, 0)
						*opened_path = nil
					}
					core.PhpStreamClose(ret)
					return nil
				}

				/* Make sure the fstat result is reused when we later try to get the
				 * file size. */

				self.SetNoForcedFstat(1)

				/* Make sure the fstat result is reused when we later try to get the
				 * file size. */

			}
			if (options & core.STREAM_USE_BLOCKING_PIPE) != 0 {
				var self *PhpStdioStreamData = (*PhpStdioStreamData)(ret.GetAbstract())
				self.SetIsPipeBlocking(1)
			}
			return ret
		}
		close(fd)
	}
	if persistent_id != nil {
		zend.Efree(persistent_id)
	}
	return nil
}
func PhpPlainFilesStreamOpener(
	wrapper *core.PhpStreamWrapper,
	path *byte,
	mode *byte,
	options int,
	opened_path **types.String,
	context *core.PhpStreamContext,
) *core.PhpStream {
	if (options&core.STREAM_DISABLE_OPEN_BASEDIR) == 0 && core.PhpCheckOpenBasedir(path) != 0 {
		return nil
	}
	return core.PhpStreamFopenRel(path, mode, opened_path, options)
}
func PhpPlainFilesUrlStater(wrapper *core.PhpStreamWrapper, url *byte, flags int, ssb *core.PhpStreamStatbuf, context *core.PhpStreamContext) int {
	if strncasecmp(url, "file://", b.SizeOf("\"file://\"")-1) == 0 {
		url += b.SizeOf("\"file://\"") - 1
	}
	if core.PhpCheckOpenBasedirEx(url, b.Cond((flags&core.PHP_STREAM_URL_STAT_QUIET) != 0, 0, 1)) != 0 {
		return -1
	}
	if (flags & core.PHP_STREAM_URL_STAT_LINK) != 0 {
		return zend.VCWD_LSTAT(url, ssb.GetSb())
	} else {
		return zend.VCWD_STAT(url, ssb.GetSb())
	}
}
func PhpPlainFilesUnlink(wrapper *core.PhpStreamWrapper, url *byte, options int, context *core.PhpStreamContext) int {
	var ret int
	if strncasecmp(url, "file://", b.SizeOf("\"file://\"")-1) == 0 {
		url += b.SizeOf("\"file://\"") - 1
	}
	if core.PhpCheckOpenBasedir(url) != 0 {
		return 0
	}
	ret = zend.VCWD_UNLINK(url)
	if ret == -1 {
		if (options & core.REPORT_ERRORS) != 0 {
			core.PhpErrorDocref1(nil, url, faults.E_WARNING, "%s", strerror(errno))
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
	if strncasecmp(url_from, "file://", b.SizeOf("\"file://\"")-1) == 0 {
		url_from += b.SizeOf("\"file://\"") - 1
	}
	if strncasecmp(url_to, "file://", b.SizeOf("\"file://\"")-1) == 0 {
		url_to += b.SizeOf("\"file://\"") - 1
	}
	if core.PhpCheckOpenBasedir(url_from) != 0 || core.PhpCheckOpenBasedir(url_to) != 0 {
		return 0
	}
	if err := os.Rename(url_from, url_to); err != nil {
		core.PhpErrorDocref2(nil, url_from, url_to, faults.E_WARNING, "%s", strerror(errno))
		return 0
	}

	/* Clear stat cache (and realpath cache) */

	standard.PhpClearStatCache(1, nil, 0)
	return 1
}
func PhpPlainFilesMkdir(wrapper *core.PhpStreamWrapper, dir *byte, mode int, options int, context *core.PhpStreamContext) int {
	var ret int
	var recursive int = options & core.PHP_STREAM_MKDIR_RECURSIVE
	var p *byte
	if strncasecmp(dir, "file://", b.SizeOf("\"file://\"")-1) == 0 {
		dir += b.SizeOf("\"file://\"") - 1
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
		if core.ExpandFilepathWithMode(dir, buf, nil, 0, zend.CWD_EXPAND) == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid path")
			return 0
		}
		e = buf + strlen(buf)
		if b.Assign(&p, memchr(buf, zend.DEFAULT_SLASH, dir_len)) {
			offset = p - buf + 1
		}
		if p != nil && dir_len == 1 {

		} else {

			/* find a top level directory we need to create */

			for b.Assign(&p, strrchr(buf+offset, zend.DEFAULT_SLASH)) || offset != 1 && b.Assign(&p, strrchr(buf, zend.DEFAULT_SLASH)) {
				var n int = 0
				*p = '0'
				for p > buf && (*(p - 1)) == zend.DEFAULT_SLASH {
					n++
					p--
					*p = '0'
				}
				if zend.VCWD_STAT(buf, &sb) == 0 {
					for true {
						*p = zend.DEFAULT_SLASH
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
		} else if !(b.Assign(&ret, standard.PhpMkdir(buf, mode))) {
			if p == nil {
				p = buf
			}

			/* create any needed directories if the creation of the 1st directory worked */

			for b.PreInc(&p) != e {
				if (*p) == '0' {
					*p = zend.DEFAULT_SLASH
					if (*(p + 1)) != '0' && b.Assign(&ret, zend.VCWD_MKDIR(buf, mode_t(mode))) < 0 {
						if (options & core.REPORT_ERRORS) != 0 {
							core.PhpErrorDocref(nil, faults.E_WARNING, "%s", strerror(errno))
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
	if strncasecmp(url, "file://", b.SizeOf("\"file://\"")-1) == 0 {
		url += b.SizeOf("\"file://\"") - 1
	}
	if core.PhpCheckOpenBasedir(url) != 0 {
		return 0
	}
	if zend.VCWD_RMDIR(url) < 0 {
		core.PhpErrorDocref1(nil, url, faults.E_WARNING, "%s", strerror(errno))
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
	if strncasecmp(url, "file://", b.SizeOf("\"file://\"")-1) == 0 {
		url += b.SizeOf("\"file://\"") - 1
	}
	if core.PhpCheckOpenBasedir(url) != 0 {
		return 0
	}
	switch option {
	case core.PHP_STREAM_META_TOUCH:
		newtime = (*__struct__utimbuf)(value)
		if zend.VCWD_ACCESS(url, F_OK) != 0 {
			var file *r.File = zend.VCWD_FOPEN(url, "w")
			if file == nil {
				core.PhpErrorDocref1(nil, url, faults.E_WARNING, "Unable to create file %s because %s", url, strerror(errno))
				return 0
			}
			file.Close()
		}
		ret = zend.VCWD_UTIME(url, newtime)
	case core.PHP_STREAM_META_OWNER_NAME:
		fallthrough
	case core.PHP_STREAM_META_OWNER:
		if option == core.PHP_STREAM_META_OWNER_NAME {
			if PhpGetUidByName((*byte)(value), &uid) != types.SUCCESS {
				core.PhpErrorDocref1(nil, url, faults.E_WARNING, "Unable to find uid for %s", (*byte)(value))
				return 0
			}
		} else {
			uid = uid_t * (*long)(value)
		}
		ret = zend.VCWD_CHOWN(url, uid, -1)
	case core.PHP_STREAM_META_GROUP:
		fallthrough
	case core.PHP_STREAM_META_GROUP_NAME:
		if option == core.PHP_STREAM_META_GROUP_NAME {
			if PhpGetGidByName((*byte)(value), &gid) != types.SUCCESS {
				core.PhpErrorDocref1(nil, url, faults.E_WARNING, "Unable to find gid for %s", (*byte)(value))
				return 0
			}
		} else {
			gid = gid_t * (*long)(value)
		}
		ret = zend.VCWD_CHOWN(url, -1, gid)
	case core.PHP_STREAM_META_ACCESS:
		mode = mode_t * (*zend.ZendLong)(value)
		ret = zend.VCWD_CHMOD(url, mode)
	default:
		core.PhpErrorDocref1(nil, url, faults.E_WARNING, "Unknown option %d for stream_metadata", option)
		return 0
	}
	if ret == -1 {
		core.PhpErrorDocref1(nil, url, faults.E_WARNING, "Operation failed: %s", strerror(errno))
		return 0
	}
	standard.PhpClearStatCache(0, nil, 0)
	return 1
}
