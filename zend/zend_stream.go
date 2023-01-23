// <<generate>>

package zend

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_stream.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Wez Furlong <wez@thebrainroom.com>                          |
   |          Scott MacVicar <scottmac@php.net>                           |
   |          Nuno Lopes <nlopess@php.net>                                |
   |          Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define ZEND_STREAM_H

// # include < sys / types . h >

// # include < sys / stat . h >

/* Lightweight stream implementation for the ZE scanners.
 * These functions are private to the engine.
 * */

type ZendStreamFsizerT func(handle any) int
type ZendStreamReaderT func(handle any, buf *byte, len_ int) ssize_t
type ZendStreamCloserT func(handle any)

// #define ZEND_MMAP_AHEAD       32

type ZendStreamType = int

const (
	ZEND_HANDLE_FILENAME = iota
	ZEND_HANDLE_FP
	ZEND_HANDLE_STREAM
)

// @type ZendStream struct

// @type ZendFileHandle struct

type ZendStatT = __struct__stat

// #define zend_fseek       fseek

// #define zend_ftell       ftell

// #define zend_lseek       lseek

// #define zend_fstat       fstat

// #define zend_stat       stat

// Source: <Zend/zend_stream.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Wez Furlong <wez@thebrainroom.com>                          |
   |          Scott MacVicar <scottmac@php.net>                           |
   |          Nuno Lopes <nlopess@php.net>                                |
   |          Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_compile.h"

// # include "zend_stream.h"

var Isatty func(fd int) int

func ZendStreamStdioReader(handle any, buf *byte, len_ int) ssize_t {
	return r.Fread(buf, 1, len_, (*r.FILE)(handle))
}
func ZendStreamStdioCloser(handle any) {
	if handle && (*r.FILE)(handle != stdin) != nil {
		r.Fclose((*r.FILE)(handle))
	}
}
func ZendStreamStdioFsizer(handle any) int {
	var buf ZendStatT
	if handle && fstat(fileno((*r.FILE)(handle)), &buf) == 0 {
		return buf.st_size
	}
	return -1
}
func ZendStreamFsize(file_handle *ZendFileHandle) int {
	r.Assert(file_handle.GetType() == ZEND_HANDLE_STREAM)
	if file_handle.GetStream().GetIsatty() != 0 {
		return 0
	}
	return file_handle.GetStream().GetFsizer()(file_handle.GetStream().GetHandle())
}
func ZendStreamInitFp(handle *ZendFileHandle, fp *r.FILE, filename string) {
	memset(handle, 0, g.SizeOf("zend_file_handle"))
	handle.SetType(ZEND_HANDLE_FP)
	handle.SetFp(fp)
	handle.SetFilename(filename)
}
func ZendStreamInitFilename(handle *ZendFileHandle, filename *byte) {
	memset(handle, 0, g.SizeOf("zend_file_handle"))
	handle.SetType(ZEND_HANDLE_FILENAME)
	handle.SetFilename(filename)
}
func ZendStreamOpen(filename *byte, handle *ZendFileHandle) int {
	var opened_path *ZendString
	if ZendStreamOpenFunction != nil {
		return ZendStreamOpenFunction(filename, handle)
	}
	ZendStreamInitFp(handle, ZendFopen(filename, &opened_path), filename)
	handle.SetOpenedPath(opened_path)
	if handle.GetFp() != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func ZendStreamGetc(file_handle *ZendFileHandle) int {
	var buf byte
	if file_handle.GetStream().GetReader()(file_handle.GetStream().GetHandle(), &buf, g.SizeOf("buf")) {
		return int(buf)
	}
	return -1
}
func ZendStreamRead(file_handle *ZendFileHandle, buf *byte, len_ int) ssize_t {
	if file_handle.GetStream().GetIsatty() != 0 {
		var c int = '*'
		var n int
		for n = 0; n < len_ && g.Assign(&c, ZendStreamGetc(file_handle)) != -1 && c != '\n'; n++ {
			buf[n] = byte(c)
		}
		if c == '\n' {
			buf[g.PostInc(&n)] = byte(c)
		}
		return n
	}
	return file_handle.GetStream().GetReader()(file_handle.GetStream().GetHandle(), buf, len_)
}
func ZendStreamFixup(file_handle *ZendFileHandle, buf **byte, len_ *int) int {
	var file_size int
	if file_handle.GetBuf() != nil {
		*buf = file_handle.GetBuf()
		*len_ = file_handle.GetLen()
		return SUCCESS
	}
	if file_handle.GetType() == ZEND_HANDLE_FILENAME {
		if ZendStreamOpen(file_handle.GetFilename(), file_handle) == FAILURE {
			return FAILURE
		}
	}
	if file_handle.GetType() == ZEND_HANDLE_FP {
		if file_handle.GetFp() == nil {
			return FAILURE
		}
		file_handle.SetType(ZEND_HANDLE_STREAM)
		file_handle.GetStream().SetHandle(file_handle.GetFp())
		file_handle.GetStream().SetIsatty(Isatty(fileno((*r.FILE)(file_handle.GetStream().GetHandle()))))
		file_handle.GetStream().SetReader(ZendStreamReaderT(ZendStreamStdioReader))
		file_handle.GetStream().SetCloser(ZendStreamCloserT(ZendStreamStdioCloser))
		file_handle.GetStream().SetFsizer(ZendStreamFsizerT(ZendStreamStdioFsizer))
	}
	file_size = ZendStreamFsize(file_handle)
	if file_size == size_t-1 {
		return FAILURE
	}
	if file_size != 0 {
		var read ssize_t
		var size int = 0
		*buf = _safeEmalloc(1, file_size, 32)
		for g.Assign(&read, ZendStreamRead(file_handle, (*buf)+size, file_size-size)) > 0 {
			size += read
		}
		if read < 0 {
			_efree(*buf)
			return FAILURE
		}
		file_handle.SetBuf(*buf)
		file_handle.SetLen(size)
	} else {
		var size int = 0
		var remain int = 4 * 1024
		var read ssize_t
		*buf = _emalloc(remain)
		for g.Assign(&read, ZendStreamRead(file_handle, (*buf)+size, remain)) > 0 {
			size += read
			remain -= read
			if remain == 0 {
				*buf = _safeErealloc(*buf, size, 2, 0)
				remain = size
			}
		}
		if read < 0 {
			_efree(*buf)
			return FAILURE
		}
		file_handle.SetLen(size)
		if size != 0 && remain < 32 {
			*buf = _safeErealloc(*buf, size, 1, 32)
		}
		file_handle.SetBuf(*buf)
	}
	if file_handle.GetLen() == 0 {
		*buf = _erealloc(*buf, 32)
		file_handle.SetBuf(*buf)
	}
	memset(file_handle.GetBuf()+file_handle.GetLen(), 0, 32)
	*buf = file_handle.GetBuf()
	*len_ = file_handle.GetLen()
	return SUCCESS
}
func ZendFileHandleDtor(fh *ZendFileHandle) {
	switch fh.GetType() {
	case ZEND_HANDLE_FP:
		r.Fclose(fh.GetFp())
		break
	case ZEND_HANDLE_STREAM:
		if fh.GetStream().GetCloser() != nil && fh.GetStream().GetHandle() {
			fh.GetStream().GetCloser()(fh.GetStream().GetHandle())
		}
		fh.GetStream().SetHandle(nil)
		break
	case ZEND_HANDLE_FILENAME:

		/* We're only supposed to get here when destructing the used_files hash,
		 * which doesn't really contain open files, but references to their names/paths
		 */

		break
	}
	if fh.GetOpenedPath() != nil {
		ZendStringReleaseEx(fh.GetOpenedPath(), 0)
		fh.SetOpenedPath(nil)
	}
	if fh.GetBuf() != nil {
		_efree(fh.GetBuf())
		fh.SetBuf(nil)
	}
	if fh.GetFreeFilename() != 0 && fh.GetFilename() != nil {
		_efree((*byte)(fh.GetFilename()))
		fh.SetFilename(nil)
	}
}

/* }}} */

func ZendCompareFileHandles(fh1 *ZendFileHandle, fh2 *ZendFileHandle) int {
	if fh1.GetType() != fh2.GetType() {
		return 0
	}
	switch fh1.GetType() {
	case ZEND_HANDLE_FILENAME:
		return strcmp(fh1.GetFilename(), fh2.GetFilename()) == 0
	case ZEND_HANDLE_FP:
		return fh1.GetFp() == fh2.GetFp()
	case ZEND_HANDLE_STREAM:
		return fh1.GetStream().GetHandle() == fh2.GetStream().GetHandle()
	default:
		return 0
	}
	return 0
}
