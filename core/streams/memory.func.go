package streams

import (
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func PhpStreamMemoryWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var ms *PhpStreamMemoryData = (*PhpStreamMemoryData)(stream.GetAbstract())
	b.Assert(ms != nil)
	if (ms.GetMode() & core.TEMP_STREAM_READONLY) != 0 {
		return ssize_t - 1
	} else if (ms.GetMode() & core.TEMP_STREAM_APPEND) != 0 {
		ms.SetFpos(ms.GetFsize())
	}
	if ms.GetFpos()+count > ms.GetFsize() {
		var tmp *byte
		if ms.GetData() == nil {
			tmp = zend.Emalloc(ms.GetFpos() + count)
		} else {
			tmp = zend.Erealloc(ms.GetData(), ms.GetFpos()+count)
		}
		ms.SetData(tmp)
		ms.SetFsize(ms.GetFpos() + count)
	}
	if ms.GetData() == nil {
		count = 0
	}
	if count != 0 {
		b.Assert(buf != nil)
		memcpy(ms.GetData()+ms.GetFpos(), (*byte)(buf), count)
		ms.SetFpos(ms.GetFpos() + count)
	}
	return count
}
func PhpStreamMemoryRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var ms *PhpStreamMemoryData = (*PhpStreamMemoryData)(stream.GetAbstract())
	b.Assert(ms != nil)
	if ms.GetFpos() == ms.GetFsize() {
		stream.SetEof(1)
		count = 0
	} else {
		if ms.GetFpos()+count >= ms.GetFsize() {
			count = ms.GetFsize() - ms.GetFpos()
		}
		if count != 0 {
			b.Assert(ms.GetData() != nil)
			b.Assert(buf != nil)
			memcpy(buf, ms.GetData()+ms.GetFpos(), count)
			ms.SetFpos(ms.GetFpos() + count)
		}
	}
	return count
}
func PhpStreamMemoryClose(stream *core.PhpStream, close_handle int) int {
	var ms *PhpStreamMemoryData = (*PhpStreamMemoryData)(stream.GetAbstract())
	b.Assert(ms != nil)
	if ms.GetData() != nil && close_handle != 0 && ms.GetMode() != core.TEMP_STREAM_READONLY {
		zend.Efree(ms.GetData())
	}
	zend.Efree(ms)
	return 0
}
func PhpStreamMemoryFlush(stream *core.PhpStream) int {
	/* nothing to do here */

	return 0

	/* nothing to do here */
}
func PhpStreamMemorySeek(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	var ms *PhpStreamMemoryData = (*PhpStreamMemoryData)(stream.GetAbstract())
	b.Assert(ms != nil)
	switch whence {
	case r.SEEK_CUR:
		if offset < 0 {
			if ms.GetFpos() < size_t(-offset) {
				ms.SetFpos(0)
				*newoffs = -1
				return -1
			} else {
				ms.SetFpos(ms.GetFpos() + offset)
				*newoffs = ms.GetFpos()
				stream.SetEof(0)
				return 0
			}
		} else {
			if ms.GetFpos()+size_t(offset) > ms.GetFsize() {
				ms.SetFpos(ms.GetFsize())
				*newoffs = -1
				return -1
			} else {
				ms.SetFpos(ms.GetFpos() + offset)
				*newoffs = ms.GetFpos()
				stream.SetEof(0)
				return 0
			}
		}
		fallthrough
	case r.SEEK_SET:
		if ms.GetFsize() < size_t(offset) {
			ms.SetFpos(ms.GetFsize())
			*newoffs = -1
			return -1
		} else {
			ms.SetFpos(offset)
			*newoffs = ms.GetFpos()
			stream.SetEof(0)
			return 0
		}
		fallthrough
	case r.SEEK_END:
		if offset > 0 {
			ms.SetFpos(ms.GetFsize())
			*newoffs = -1
			return -1
		} else if ms.GetFsize() < size_t(-offset) {
			ms.SetFpos(0)
			*newoffs = -1
			return -1
		} else {
			ms.SetFpos(ms.GetFsize() + offset)
			*newoffs = ms.GetFpos()
			stream.SetEof(0)
			return 0
		}
		fallthrough
	default:
		*newoffs = ms.GetFpos()
		return -1
	}
}
func PhpStreamMemoryCast(stream *core.PhpStream, castas int, ret *any) int { return types.FAILURE }
func PhpStreamMemoryStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	var timestamp int64 = 0
	var ms *PhpStreamMemoryData = (*PhpStreamMemoryData)(stream.GetAbstract())
	b.Assert(ms != nil)
	memset(ssb, 0, b.SizeOf("php_stream_statbuf"))

	/* read-only across the board */

	if (ms.GetMode() & core.TEMP_STREAM_READONLY) != 0 {
		ssb.GetSb().st_mode = 0444
	} else {
		ssb.GetSb().st_mode = 0666
	}
	ssb.GetSb().st_size = ms.GetFsize()
	ssb.GetSb().st_mode |= S_IFREG
	ssb.GetSb().st_mtime = timestamp
	ssb.GetSb().st_atime = timestamp
	ssb.GetSb().st_ctime = timestamp
	ssb.GetSb().st_nlink = 1
	ssb.GetSb().st_rdev = -1

	/* this is only for APC, so use /dev/null device - no chance of conflict there! */

	ssb.GetSb().st_dev = 0xc

	/* generate unique inode number for alias/filename, so no phars will conflict */

	ssb.GetSb().st_ino = 0
	ssb.GetSb().st_blksize = -1
	ssb.GetSb().st_blocks = -1
	return 0
}
func PhpStreamMemorySetOption(stream *core.PhpStream, option int, value int, ptrparam any) int {
	var ms *PhpStreamMemoryData = (*PhpStreamMemoryData)(stream.GetAbstract())
	var newsize int
	switch option {
	case core.PHP_STREAM_OPTION_TRUNCATE_API:
		switch value {
		case core.PHP_STREAM_TRUNCATE_SUPPORTED:
			return core.PHP_STREAM_OPTION_RETURN_OK
		case core.PHP_STREAM_TRUNCATE_SET_SIZE:
			if (ms.GetMode() & core.TEMP_STREAM_READONLY) != 0 {
				return core.PHP_STREAM_OPTION_RETURN_ERR
			}
			newsize = *((*int)(ptrparam))
			if newsize <= ms.GetFsize() {
				if newsize < ms.GetFpos() {
					ms.SetFpos(newsize)
				}
			} else {
				ms.SetData(zend.Erealloc(ms.GetData(), newsize))
				memset(ms.GetData()+ms.GetFsize(), 0, newsize-ms.GetFsize())
				ms.SetFsize(newsize)
			}
			ms.SetFsize(newsize)
			return core.PHP_STREAM_OPTION_RETURN_OK
		}
		fallthrough
	default:
		return core.PHP_STREAM_OPTION_RETURN_NOTIMPL
	}
}
func PhpStreamModeFromStr(mode *byte) int {
	if strpbrk(mode, "a") {
		return core.TEMP_STREAM_APPEND
	} else if strpbrk(mode, "w+") {
		return core.TEMP_STREAM_DEFAULT
	}
	return core.TEMP_STREAM_READONLY
}
func PhpStreamTempWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.GetAbstract())
	b.Assert(ts != nil)
	if ts.GetInnerstream() == nil {
		return -1
	}
	if core.PhpStreamIs(ts.GetInnerstream(), core.PHP_STREAM_IS_MEMORY) {
		var memsize int
		var membuf *byte = core.PhpStreamMemoryGetBuffer(ts.GetInnerstream(), &memsize)
		if memsize+count >= ts.GetSmax() {
			var file *core.PhpStream = PhpStreamFopenTemporaryFile(ts.GetTmpdir(), "php", nil)
			if file == nil {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to create temporary file, Check permissions in temporary files directory.")
				return 0
			}
			core.PhpStreamWrite(file, membuf, memsize)
			core.PhpStreamFreeEnclosed(ts.GetInnerstream(), core.PHP_STREAM_FREE_CLOSE)
			ts.SetInnerstream(file)
			PhpStreamEncloses(stream, ts.GetInnerstream())
		}
	}
	return core.PhpStreamWrite(ts.GetInnerstream(), buf, count)
}
func PhpStreamTempRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.GetAbstract())
	var got int
	b.Assert(ts != nil)
	if ts.GetInnerstream() == nil {
		return -1
	}
	got = core.PhpStreamRead(ts.GetInnerstream(), buf, count)
	stream.SetEof(ts.GetInnerstream().GetEof())
	return got
}
func PhpStreamTempClose(stream *core.PhpStream, close_handle int) int {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.GetAbstract())
	var ret int
	b.Assert(ts != nil)
	if ts.GetInnerstream() != nil {
		ret = core.PhpStreamFreeEnclosed(ts.GetInnerstream(), core.PHP_STREAM_FREE_CLOSE|b.Cond(close_handle != 0, 0, core.PHP_STREAM_FREE_PRESERVE_HANDLE))
	} else {
		ret = 0
	}
	// zend.ZvalPtrDtor(ts.GetMeta())
	if ts.GetTmpdir() != nil {
		zend.Efree(ts.GetTmpdir())
	}
	zend.Efree(ts)
	return ret
}
func PhpStreamTempFlush(stream *core.PhpStream) int {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.GetAbstract())
	b.Assert(ts != nil)
	if ts.GetInnerstream() != nil {
		return core.PhpStreamFlush(ts.GetInnerstream())
	} else {
		return -1
	}
}
func PhpStreamTempSeek(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.GetAbstract())
	var ret int
	b.Assert(ts != nil)
	if ts.GetInnerstream() == nil {
		*newoffs = -1
		return -1
	}
	ret = core.PhpStreamSeek(ts.GetInnerstream(), offset, whence)
	*newoffs = ts.GetInnerstream().GetPosition()
	stream.SetEof(ts.GetInnerstream().GetEof())
	return ret
}
func PhpStreamTempCast(stream *core.PhpStream, castas int, ret *any) int {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.GetAbstract())
	var file *core.PhpStream
	var memsize int
	var membuf *byte
	var pos zend.ZendOffT
	b.Assert(ts != nil)
	if ts.GetInnerstream() == nil {
		return types.FAILURE
	}
	if core.PhpStreamIs(ts.GetInnerstream(), core.PHP_STREAM_IS_STDIO) {
		return core.PhpStreamCast(ts.GetInnerstream(), castas, ret, 0)
	}

	/* we are still using a memory based backing. If they are if we can be
	 * a FILE*, say yes because we can perform the conversion.
	 * If they actually want to perform the conversion, we need to switch
	 * the memory stream to a tmpfile stream */

	if ret == nil && castas == core.PHP_STREAM_AS_STDIO {
		return types.SUCCESS
	}

	/* say "no" to other stream forms */

	if ret == nil {
		return types.FAILURE
	}
	file = _phpStreamFopenTmpfile(0)
	if file == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to create temporary file.")
		return types.FAILURE
	}

	/* perform the conversion and then pass the request on to the innerstream */

	membuf = core.PhpStreamMemoryGetBuffer(ts.GetInnerstream(), &memsize)
	core.PhpStreamWrite(file, membuf, memsize)
	pos = ts.GetInnerstream().GetPosition()
	core.PhpStreamFreeEnclosed(ts.GetInnerstream(), core.PHP_STREAM_FREE_CLOSE)
	ts.SetInnerstream(file)
	PhpStreamEncloses(stream, ts.GetInnerstream())
	core.PhpStreamSeek(ts.GetInnerstream(), pos, r.SEEK_SET)
	return core.PhpStreamCast(ts.GetInnerstream(), castas, ret, 1)
}
func PhpStreamTempStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.GetAbstract())
	if ts == nil || ts.GetInnerstream() == nil {
		return -1
	}
	return core.PhpStreamStat(ts.GetInnerstream(), ssb)
}
func PhpStreamTempSetOption(stream *core.PhpStream, option int, value int, ptrparam any) int {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.GetAbstract())
	switch option {
	case core.PHP_STREAM_OPTION_META_DATA_API:
		if ts.GetMeta().IsNotUndef() {
			types.ZendHashCopy((*types.Zval)(ptrparam).Array(), ts.GetMeta().Array(), zend.ZvalAddRef)
		}
		return core.PHP_STREAM_OPTION_RETURN_OK
	default:
		if ts.GetInnerstream() != nil {
			return core.PhpStreamSetOption(ts.GetInnerstream(), option, value, ptrparam)
		}
		return core.PHP_STREAM_OPTION_RETURN_NOTIMPL
	}
}
func PhpStreamUrlWrapRfc2397(
	wrapper *core.PhpStreamWrapper,
	path *byte,
	mode *byte,
	options int,
	opened_path **types.String,
	context *core.PhpStreamContext,
) *core.PhpStream {
	var stream *core.PhpStream
	var ts *PhpStreamTempData
	var comma *byte
	var semi *byte
	var sep *byte
	var mlen int
	var dlen int
	var plen int
	var vlen int
	var ilen int
	var newoffs zend.ZendOffT
	var meta types.Zval
	var base64 int = 0
	var base64_comma *types.String = nil
	meta.SetNull()
	if memcmp(path, "data:", 5) {
		return nil
	}
	path += 5
	dlen = strlen(path)
	if dlen >= 2 && path[0] == '/' && path[1] == '/' {
		dlen -= 2
		path += 2
	}
	if b.Assign(&comma, memchr(path, ',', dlen)) == nil {
		PhpStreamWrapperLogError(wrapper, options, "rfc2397: no comma in URL")
		return nil
	}
	if comma != path {

		/* meta info */

		mlen = comma - path
		dlen -= mlen
		semi = memchr(path, ';', mlen)
		sep = memchr(path, '/', mlen)
		if semi == nil && sep == nil {
			PhpStreamWrapperLogError(wrapper, options, "rfc2397: illegal media type")
			return nil
		}
		zend.ArrayInit(&meta)
		if semi == nil {
			zend.AddAssocStringl(&meta, "mediatype", (*byte)(path), mlen)
			mlen = 0
		} else if sep != nil && sep < semi {
			plen = semi - path
			zend.AddAssocStringl(&meta, "mediatype", (*byte)(path), plen)
			mlen -= plen
			path += plen
		} else if semi != path || mlen != b.SizeOf("\";base64\"")-1 || memcmp(path, ";base64", b.SizeOf("\";base64\"")-1) {
			// zend.ZvalPtrDtor(&meta)
			PhpStreamWrapperLogError(wrapper, options, "rfc2397: illegal media type")
			return nil
		}

		/* get parameters and potentially ';base64' */

		for semi != nil && semi == path {
			path++
			mlen--
			sep = memchr(path, '=', mlen)
			semi = memchr(path, ';', mlen)
			if sep == nil || semi != nil && semi < sep {
				if mlen != b.SizeOf("\"base64\"")-1 || memcmp(path, "base64", b.SizeOf("\"base64\"")-1) {

					/* must be error since parameters are only allowed after mediatype and we have no '=' sign */

					// zend.ZvalPtrDtor(&meta)
					PhpStreamWrapperLogError(wrapper, options, "rfc2397: illegal parameter")
					return nil
				}
				base64 = 1
				mlen -= b.SizeOf("\"base64\"") - 1
				path += b.SizeOf("\"base64\"") - 1
				break
			}

			/* found parameter ... the heart of cs ppl lies in +1/-1 or was it +2 this time? */

			plen = sep - path
			vlen = b.CondF1(semi != nil, func() __auto__ { return size_t(semi - sep) }, mlen-plen) - 1
			if plen != b.SizeOf("\"mediatype\"")-1 || memcmp(path, "mediatype", b.SizeOf("\"mediatype\"")-1) {
				zend.AddAssocStringlEx(&meta, b.CastStr(path, plen), b.CastStr(sep+1, vlen))
			}
			plen += vlen + 1
			mlen -= plen
			path += plen
		}
		if mlen != 0 {
			// zend.ZvalPtrDtor(&meta)
			PhpStreamWrapperLogError(wrapper, options, "rfc2397: illegal URL")
			return nil
		}
	} else {
		zend.ArrayInit(&meta)
	}
	zend.AddAssocBool(&meta, "base64", base64)

	/* skip ',' */

	comma++
	dlen--
	if base64 != 0 {
		if ret, ok := standard.PhpBase64DecodeEx(b.CastStr(comma, dlen), true); ok {
			base64_comma = types.NewString(ret)
		} else {
			// zend.ZvalPtrDtor(&meta)
			PhpStreamWrapperLogError(wrapper, options, "rfc2397: unable to decode")
			return nil
		}
		comma = base64_comma.GetVal()
		ilen = base64_comma.GetLen()
	} else {
		comma = zend.Estrndup(comma, dlen)
		dlen = PhpUrlDecode(comma, dlen)
		ilen = dlen
	}
	if b.Assign(&stream, core.PhpStreamTempCreateRel(0, ^0)) != nil {

		/* store data */

		PhpStreamTempWrite(stream, comma, ilen)
		PhpStreamTempSeek(stream, 0, r.SEEK_SET, &newoffs)

		/* set special stream stuff (enforce exact mode) */

		vlen = strlen(mode)
		if vlen >= b.SizeOf("stream -> mode") {
			vlen = b.SizeOf("stream -> mode") - 1
		}
		memcpy(stream.GetMode(), mode, vlen)
		stream.GetMode()[vlen] = '0'
		stream.SetOps(&PhpStreamRfc2397Ops)
		ts = (*PhpStreamTempData)(stream.GetAbstract())
		b.Assert(ts != nil)
		if mode != nil && mode[0] == 'r' && mode[1] != '+' {
			ts.SetMode(core.TEMP_STREAM_READONLY)
		} else {
			ts.SetMode(0)
		}
		types.ZVAL_COPY_VALUE(ts.GetMeta(), &meta)
	}
	if base64_comma != nil {
		//types.ZendStringFree(base64_comma)
	} else {
		zend.Efree(comma)
	}
	return stream
}
