// <<generate>>

package streams

import (
	"sik/core"
	"sik/ext/standard"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/streams/memory.c>

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

// #define _GNU_SOURCE

// # include "php.h"

// # include "ext/standard/base64.h"

var PhpUrlDecode func(str *byte, len_ int) int

/* Memory streams use a dynamic memory buffer to emulate a stream.
 * You can use php_stream_memory_open to create a readonly stream
 * from an existing memory buffer.
 */

// @type PhpStreamMemoryData struct

/* {{{ */

func PhpStreamMemoryWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var ms *PhpStreamMemoryData = (*PhpStreamMemoryData)(stream.abstract)
	r.Assert(ms != nil)
	if (ms.GetMode() & 0x1) != 0 {
		return ssize_t - 1
	} else if (ms.GetMode() & 0x4) != 0 {
		ms.SetFpos(ms.GetFsize())
	}
	if ms.GetFpos()+count > ms.GetFsize() {
		var tmp *byte
		if ms.GetData() == nil {
			tmp = zend._emalloc(ms.GetFpos() + count)
		} else {
			tmp = zend._erealloc(ms.GetData(), ms.GetFpos()+count)
		}
		ms.SetData(tmp)
		ms.SetFsize(ms.GetFpos() + count)
	}
	if ms.GetData() == nil {
		count = 0
	}
	if count != 0 {
		r.Assert(buf != nil)
		memcpy(ms.GetData()+ms.GetFpos(), (*byte)(buf), count)
		ms.SetFpos(ms.GetFpos() + count)
	}
	return count
}

/* }}} */

func PhpStreamMemoryRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var ms *PhpStreamMemoryData = (*PhpStreamMemoryData)(stream.abstract)
	r.Assert(ms != nil)
	if ms.GetFpos() == ms.GetFsize() {
		stream.eof = 1
		count = 0
	} else {
		if ms.GetFpos()+count >= ms.GetFsize() {
			count = ms.GetFsize() - ms.GetFpos()
		}
		if count != 0 {
			r.Assert(ms.GetData() != nil)
			r.Assert(buf != nil)
			memcpy(buf, ms.GetData()+ms.GetFpos(), count)
			ms.SetFpos(ms.GetFpos() + count)
		}
	}
	return count
}

/* }}} */

func PhpStreamMemoryClose(stream *core.PhpStream, close_handle int) int {
	var ms *PhpStreamMemoryData = (*PhpStreamMemoryData)(stream.abstract)
	r.Assert(ms != nil)
	if ms.GetData() != nil && close_handle != 0 && ms.GetMode() != 0x1 {
		zend._efree(ms.GetData())
	}
	zend._efree(ms)
	return 0
}

/* }}} */

func PhpStreamMemoryFlush(stream *core.PhpStream) int {
	/* nothing to do here */

	return 0

	/* nothing to do here */
}

/* }}} */

func PhpStreamMemorySeek(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	var ms *PhpStreamMemoryData = (*PhpStreamMemoryData)(stream.abstract)
	r.Assert(ms != nil)
	switch whence {
	case 1:
		if offset < 0 {
			if ms.GetFpos() < size_t(-offset) {
				ms.SetFpos(0)
				*newoffs = -1
				return -1
			} else {
				ms.SetFpos(ms.GetFpos() + offset)
				*newoffs = ms.GetFpos()
				stream.eof = 0
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
				stream.eof = 0
				return 0
			}
		}
	case 0:
		if ms.GetFsize() < size_t(offset) {
			ms.SetFpos(ms.GetFsize())
			*newoffs = -1
			return -1
		} else {
			ms.SetFpos(offset)
			*newoffs = ms.GetFpos()
			stream.eof = 0
			return 0
		}
	case 2:
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
			stream.eof = 0
			return 0
		}
	default:
		*newoffs = ms.GetFpos()
		return -1
	}
}

/* }}} */

func PhpStreamMemoryCast(stream *core.PhpStream, castas int, ret *any) int { return zend.FAILURE }

/* }}} */

func PhpStreamMemoryStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	var timestamp int64 = 0
	var ms *PhpStreamMemoryData = (*PhpStreamMemoryData)(stream.abstract)
	r.Assert(ms != nil)
	memset(ssb, 0, g.SizeOf("php_stream_statbuf"))

	/* read-only across the board */

	if (ms.GetMode() & 0x1) != 0 {
		ssb.sb.st_mode = 0444
	} else {
		ssb.sb.st_mode = 0666
	}
	ssb.sb.st_size = ms.GetFsize()
	ssb.sb.st_mode |= S_IFREG
	ssb.sb.st_mtime = timestamp
	ssb.sb.st_atime = timestamp
	ssb.sb.st_ctime = timestamp
	ssb.sb.st_nlink = 1
	ssb.sb.st_rdev = -1

	/* this is only for APC, so use /dev/null device - no chance of conflict there! */

	ssb.sb.st_dev = 0xc

	/* generate unique inode number for alias/filename, so no phars will conflict */

	ssb.sb.st_ino = 0
	ssb.sb.st_blksize = -1
	ssb.sb.st_blocks = -1
	return 0
}

/* }}} */

func PhpStreamMemorySetOption(stream *core.PhpStream, option int, value int, ptrparam any) int {
	var ms *PhpStreamMemoryData = (*PhpStreamMemoryData)(stream.abstract)
	var newsize int
	switch option {
	case 10:
		switch value {
		case 0:
			return 0
		case 1:
			if (ms.GetMode() & 0x1) != 0 {
				return -1
			}
			newsize = *((*int)(ptrparam))
			if newsize <= ms.GetFsize() {
				if newsize < ms.GetFpos() {
					ms.SetFpos(newsize)
				}
			} else {
				ms.SetData(zend._erealloc(ms.GetData(), newsize))
				memset(ms.GetData()+ms.GetFsize(), 0, newsize-ms.GetFsize())
				ms.SetFsize(newsize)
			}
			ms.SetFsize(newsize)
			return 0
		}
	default:
		return -2
	}
}

/* }}} */

var PhpStreamMemoryOps core.PhpStreamOps = core.PhpStreamOps{PhpStreamMemoryWrite, PhpStreamMemoryRead, PhpStreamMemoryClose, PhpStreamMemoryFlush, "MEMORY", PhpStreamMemorySeek, PhpStreamMemoryCast, PhpStreamMemoryStat, PhpStreamMemorySetOption}

/* {{{ */

func PhpStreamModeFromStr(mode *byte) int {
	if strpbrk(mode, "a") {
		return 0x4
	} else if strpbrk(mode, "w+") {
		return 0x0
	}
	return 0x1
}

/* }}} */

func _phpStreamModeToStr(mode int) *byte {
	if mode == 0x1 {
		return "rb"
	} else if mode == 0x4 {
		return "a+b"
	}
	return "w+b"
}

/* }}} */

func _phpStreamMemoryCreate(mode int) *core.PhpStream {
	var self *PhpStreamMemoryData
	var stream *core.PhpStream
	self = zend._emalloc(g.SizeOf("* self"))
	self.SetData(nil)
	self.SetFpos(0)
	self.SetFsize(0)
	self.SetSmax(^0)
	self.SetMode(mode)
	stream = _phpStreamAlloc(&PhpStreamMemoryOps, self, 0, _phpStreamModeToStr(mode))
	stream.flags |= 0x2
	return stream
}

/* }}} */

func _phpStreamMemoryOpen(mode int, buf *byte, length int) *core.PhpStream {
	var stream *core.PhpStream
	var ms *PhpStreamMemoryData
	if g.Assign(&stream, _phpStreamMemoryCreate(mode)) != nil {
		ms = (*PhpStreamMemoryData)(stream.abstract)
		if mode == 0x1 || mode == 0x2 {

			/* use the buffer directly */

			ms.SetData(buf)
			ms.SetFsize(length)
		} else {
			if length != 0 {
				r.Assert(buf != nil)
				_phpStreamWrite(stream, buf, length)
			}
		}
	}
	return stream
}

/* }}} */

func _phpStreamMemoryGetBuffer(stream *core.PhpStream, length *int) *byte {
	var ms *PhpStreamMemoryData = (*PhpStreamMemoryData)(stream.abstract)
	r.Assert(ms != nil)
	r.Assert(length != 0)
	*length = ms.GetFsize()
	return ms.GetData()
}

/* }}} */

// @type PhpStreamTempData struct

/* {{{ */

func PhpStreamTempWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.abstract)
	r.Assert(ts != nil)
	if ts.GetInnerstream() == nil {
		return -1
	}
	if ts.GetInnerstream().ops == &PhpStreamMemoryOps {
		var memsize int
		var membuf *byte = _phpStreamMemoryGetBuffer(ts.GetInnerstream(), &memsize)
		if memsize+count >= ts.GetSmax() {
			var file *core.PhpStream = _phpStreamFopenTemporaryFile(ts.GetTmpdir(), "php", nil)
			if file == nil {
				core.PhpErrorDocref(nil, 1<<1, "Unable to create temporary file, Check permissions in temporary files directory.")
				return 0
			}
			_phpStreamWrite(file, membuf, memsize)
			_phpStreamFreeEnclosed(ts.GetInnerstream(), 1|2)
			ts.SetInnerstream(file)
			PhpStreamEncloses(stream, ts.GetInnerstream())
		}
	}
	return _phpStreamWrite(ts.GetInnerstream(), buf, count)
}

/* }}} */

func PhpStreamTempRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.abstract)
	var got int
	r.Assert(ts != nil)
	if ts.GetInnerstream() == nil {
		return -1
	}
	got = _phpStreamRead(ts.GetInnerstream(), buf, count)
	stream.eof = ts.GetInnerstream().eof
	return got
}

/* }}} */

func PhpStreamTempClose(stream *core.PhpStream, close_handle int) int {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.abstract)
	var ret int
	r.Assert(ts != nil)
	if ts.GetInnerstream() != nil {
		ret = _phpStreamFreeEnclosed(ts.GetInnerstream(), 1|2|g.Cond(close_handle != 0, 0, 4))
	} else {
		ret = 0
	}
	zend.ZvalPtrDtor(&ts.meta)
	if ts.GetTmpdir() != nil {
		zend._efree(ts.GetTmpdir())
	}
	zend._efree(ts)
	return ret
}

/* }}} */

func PhpStreamTempFlush(stream *core.PhpStream) int {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.abstract)
	r.Assert(ts != nil)
	if ts.GetInnerstream() != nil {
		return _phpStreamFlush(ts.GetInnerstream(), 0)
	} else {
		return -1
	}
}

/* }}} */

func PhpStreamTempSeek(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.abstract)
	var ret int
	r.Assert(ts != nil)
	if ts.GetInnerstream() == nil {
		*newoffs = -1
		return -1
	}
	ret = _phpStreamSeek(ts.GetInnerstream(), offset, whence)
	*newoffs = _phpStreamTell(ts.GetInnerstream())
	stream.eof = ts.GetInnerstream().eof
	return ret
}

/* }}} */

func PhpStreamTempCast(stream *core.PhpStream, castas int, ret *any) int {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.abstract)
	var file *core.PhpStream
	var memsize int
	var membuf *byte
	var pos zend.ZendOffT
	r.Assert(ts != nil)
	if ts.GetInnerstream() == nil {
		return zend.FAILURE
	}
	if ts.GetInnerstream().ops == &PhpStreamStdioOps {
		return _phpStreamCast(ts.GetInnerstream(), castas, ret, 0)
	}

	/* we are still using a memory based backing. If they are if we can be
	 * a FILE*, say yes because we can perform the conversion.
	 * If they actually want to perform the conversion, we need to switch
	 * the memory stream to a tmpfile stream */

	if ret == nil && castas == 0 {
		return zend.SUCCESS
	}

	/* say "no" to other stream forms */

	if ret == nil {
		return zend.FAILURE
	}
	file = _phpStreamFopenTmpfile(0)
	if file == nil {
		core.PhpErrorDocref(nil, 1<<1, "Unable to create temporary file.")
		return zend.FAILURE
	}

	/* perform the conversion and then pass the request on to the innerstream */

	membuf = _phpStreamMemoryGetBuffer(ts.GetInnerstream(), &memsize)
	_phpStreamWrite(file, membuf, memsize)
	pos = _phpStreamTell(ts.GetInnerstream())
	_phpStreamFreeEnclosed(ts.GetInnerstream(), 1|2)
	ts.SetInnerstream(file)
	PhpStreamEncloses(stream, ts.GetInnerstream())
	_phpStreamSeek(ts.GetInnerstream(), pos, 0)
	return _phpStreamCast(ts.GetInnerstream(), castas, ret, 1)
}

/* }}} */

func PhpStreamTempStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.abstract)
	if ts == nil || ts.GetInnerstream() == nil {
		return -1
	}
	return _phpStreamStat(ts.GetInnerstream(), ssb)
}

/* }}} */

func PhpStreamTempSetOption(stream *core.PhpStream, option int, value int, ptrparam any) int {
	var ts *PhpStreamTempData = (*PhpStreamTempData)(stream.abstract)
	switch option {
	case 11:
		if ts.meta.u1.v.type_ != 0 {
			zend.ZendHashCopy((*zend.Zval)(ptrparam).value.arr, ts.meta.value.arr, zend.ZvalAddRef)
		}
		return 0
	default:
		if ts.GetInnerstream() != nil {
			return _phpStreamSetOption(ts.GetInnerstream(), option, value, ptrparam)
		}
		return -2
	}
}

/* }}} */

var PhpStreamTempOps core.PhpStreamOps = core.PhpStreamOps{PhpStreamTempWrite, PhpStreamTempRead, PhpStreamTempClose, PhpStreamTempFlush, "TEMP", PhpStreamTempSeek, PhpStreamTempCast, PhpStreamTempStat, PhpStreamTempSetOption}

/* }}} */

func _phpStreamTempCreateEx(mode int, max_memory_usage int, tmpdir *byte) *core.PhpStream {
	var self *PhpStreamTempData
	var stream *core.PhpStream
	self = zend._ecalloc(1, g.SizeOf("* self"))
	self.SetSmax(max_memory_usage)
	self.SetMode(mode)
	&self.meta.u1.type_info = 0
	if tmpdir != nil {
		self.SetTmpdir(zend._estrdup(tmpdir))
	}
	stream = _phpStreamAlloc(&PhpStreamTempOps, self, 0, _phpStreamModeToStr(mode))
	stream.flags |= 0x2
	self.SetInnerstream(_phpStreamMemoryCreate(mode))
	PhpStreamEncloses(stream, self.GetInnerstream())
	return stream
}

/* }}} */

func _phpStreamTempCreate(mode int, max_memory_usage int) *core.PhpStream {
	return _phpStreamTempCreateEx(mode, max_memory_usage, nil)
}

/* }}} */

func _phpStreamTempOpen(mode int, max_memory_usage int, buf *byte, length int) *core.PhpStream {
	var stream *core.PhpStream
	var ts *PhpStreamTempData
	var newoffs zend.ZendOffT
	if g.Assign(&stream, _phpStreamTempCreate(mode, max_memory_usage)) != nil {
		if length != 0 {
			r.Assert(buf != nil)
			PhpStreamTempWrite(stream, buf, length)
			PhpStreamTempSeek(stream, 0, 0, &newoffs)
		}
		ts = (*PhpStreamTempData)(stream.abstract)
		r.Assert(ts != nil)
		ts.SetMode(mode)
	}
	return stream
}

/* }}} */

var PhpStreamRfc2397Ops core.PhpStreamOps = core.PhpStreamOps{PhpStreamTempWrite, PhpStreamTempRead, PhpStreamTempClose, PhpStreamTempFlush, "RFC2397", PhpStreamTempSeek, PhpStreamTempCast, PhpStreamTempStat, PhpStreamTempSetOption}

func PhpStreamUrlWrapRfc2397(wrapper *core.PhpStreamWrapper, path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
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
	var meta zend.Zval
	var base64 int = 0
	var base64_comma *zend.ZendString = nil
	&meta.u1.type_info = 1
	if memcmp(path, "data:", 5) {
		return nil
	}
	path += 5
	dlen = strlen(path)
	if dlen >= 2 && path[0] == '/' && path[1] == '/' {
		dlen -= 2
		path += 2
	}
	if g.Assign(&comma, memchr(path, ',', dlen)) == nil {
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
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &meta
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		if semi == nil {
			zend.AddAssocStringlEx(&meta, "mediatype", strlen("mediatype"), (*byte)(path), mlen)
			mlen = 0
		} else if sep != nil && sep < semi {
			plen = semi - path
			zend.AddAssocStringlEx(&meta, "mediatype", strlen("mediatype"), (*byte)(path), plen)
			mlen -= plen
			path += plen
		} else if semi != path || mlen != g.SizeOf("\";base64\"")-1 || memcmp(path, ";base64", g.SizeOf("\";base64\"")-1) {
			zend.ZvalPtrDtor(&meta)
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
				if mlen != g.SizeOf("\"base64\"")-1 || memcmp(path, "base64", g.SizeOf("\"base64\"")-1) {

					/* must be error since parameters are only allowed after mediatype and we have no '=' sign */

					zend.ZvalPtrDtor(&meta)
					PhpStreamWrapperLogError(wrapper, options, "rfc2397: illegal parameter")
					return nil
				}
				base64 = 1
				mlen -= g.SizeOf("\"base64\"") - 1
				path += g.SizeOf("\"base64\"") - 1
				break
			}

			/* found parameter ... the heart of cs ppl lies in +1/-1 or was it +2 this time? */

			plen = sep - path
			vlen = g.CondF1(semi != nil, func() __auto__ { return size_t(semi - sep) }, mlen-plen) - 1
			if plen != g.SizeOf("\"mediatype\"")-1 || memcmp(path, "mediatype", g.SizeOf("\"mediatype\"")-1) {
				zend.AddAssocStringlEx(&meta, path, plen, sep+1, vlen)
			}
			plen += vlen + 1
			mlen -= plen
			path += plen
		}
		if mlen != 0 {
			zend.ZvalPtrDtor(&meta)
			PhpStreamWrapperLogError(wrapper, options, "rfc2397: illegal URL")
			return nil
		}
	} else {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &meta
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	}
	zend.AddAssocBoolEx(&meta, "base64", strlen("base64"), base64)

	/* skip ',' */

	comma++
	dlen--
	if base64 != 0 {
		base64_comma = standard.PhpBase64DecodeEx((*uint8)(comma), dlen, 1)
		if base64_comma == nil {
			zend.ZvalPtrDtor(&meta)
			PhpStreamWrapperLogError(wrapper, options, "rfc2397: unable to decode")
			return nil
		}
		comma = base64_comma.val
		ilen = base64_comma.len_
	} else {
		comma = zend._estrndup(comma, dlen)
		dlen = PhpUrlDecode(comma, dlen)
		ilen = dlen
	}
	if g.Assign(&stream, _phpStreamTempCreate(0, ^0)) != nil {

		/* store data */

		PhpStreamTempWrite(stream, comma, ilen)
		PhpStreamTempSeek(stream, 0, 0, &newoffs)

		/* set special stream stuff (enforce exact mode) */

		vlen = strlen(mode)
		if vlen >= g.SizeOf("stream -> mode") {
			vlen = g.SizeOf("stream -> mode") - 1
		}
		memcpy(stream.mode, mode, vlen)
		stream.mode[vlen] = '0'
		stream.ops = &PhpStreamRfc2397Ops
		ts = (*PhpStreamTempData)(stream.abstract)
		r.Assert(ts != nil)
		if mode != nil && mode[0] == 'r' && mode[1] != '+' {
			ts.SetMode(0x1)
		} else {
			ts.SetMode(0)
		}
		var _z1 *zend.Zval = &ts.meta
		var _z2 *zend.Zval = &meta
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	}
	if base64_comma != nil {
		zend.ZendStringFree(base64_comma)
	} else {
		zend._efree(comma)
	}
	return stream
}

var PhpStreamRfc2397Wops core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{PhpStreamUrlWrapRfc2397, nil, nil, nil, nil, "RFC2397", nil, nil, nil, nil, nil}
var PhpStreamRfc2397Wrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&PhpStreamRfc2397Wops, nil, 1}
