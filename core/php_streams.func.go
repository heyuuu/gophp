package core

import (
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core/streams"
	b "github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

func PhpStreamAllocRel(ops *PhpStreamOps, thisptr any, persistent *byte, mode string) *PhpStream {
	return _phpStreamAlloc(ops, thisptr, persistent, mode)
}
func PhpStreamFopenRel(filename *byte, mode *byte, opened **types.String, options int) *PhpStream {
	return streams._phpStreamFopen(filename, mode, opened, options)
}
func PhpStreamFopenFromFdRel(fd int, mode *byte, persistent_id *byte) *PhpStream {
	return streams._phpStreamFopenFromFd(fd, mode, persistent_id)
}
func PhpStreamOpenWrapperRel(path *byte, mode string, options int, opened **types.String) *PhpStream {
	return _phpStreamOpenWrapperEx(path, mode, options, opened, nil)
}
func PHP_STREAM_CONTEXT(stream *PhpStream) *PhpStreamContext {
	return (*PhpStreamContext)(b.CondF1(stream.GetCtx() != nil, func() any { return stream.GetCtx().GetPtr() }, nil))
}
func PhpStreamAlloc(ops *PhpStreamOps, thisptr any, persistent_id int, mode *byte) *PhpStream {
	return _phpStreamAlloc(ops, thisptr, persistent_id, mode)
}
func PhpStreamAutoCleanup(stream *PhpStream) { stream.SetExposed(1) }
func PhpStreamToZval(stream *PhpStream, zval *types.Zval) {
	zval.SetResource(stream.GetRes())
	stream.SetExposed(1)
}
func PhpStreamFromZval(xstr *PhpStream, pzval *types.Zval) {
	if b.Assign(&xstr, (*PhpStream)(zend.ZendFetchResource2Ex(pzval, "stream", PhpFileLeStream(), PhpFileLePstream()))) == nil {
		return_value.SetFalse()
		return
	}
}
func PhpStreamFromRes(xstr *PhpStream, res *types.ZendResource) {
	if b.Assign(&xstr, (*PhpStream)(zend.ZendFetchResource2(res, "stream", PhpFileLeStream(), PhpFileLePstream()))) == nil {
		return_value.SetFalse()
		return
	}
}
func PhpStreamFromZvalNoVerify(xstr *PhpStream, pzval *types.Zval) *PhpStream {
	xstr = (*PhpStream)(zend.ZendFetchResource2Ex(pzval, "stream", PhpFileLeStream(), PhpFileLePstream()))
	return xstr
}
func PhpStreamFreeEnclosed(stream_enclosed *PhpStream, close_options int) int {
	return _phpStreamFreeEnclosed(stream_enclosed, close_options)
}
func PhpStreamFree(stream *PhpStream, close_options int) int {
	return _phpStreamFree(stream, close_options)
}
func PhpStreamClose(stream *PhpStream) int {
	return _phpStreamFree(stream, PHP_STREAM_FREE_CLOSE)
}
func PhpStreamPclose(stream *PhpStream) int {
	return _phpStreamFree(stream, PHP_STREAM_FREE_CLOSE_PERSISTENT)
}
func PhpStreamRewind(stream *PhpStream) int {
	return _phpStreamSeek(stream, 0, r.SEEK_SET)
}
func PhpStreamSeek(stream *PhpStream, offset zend.ZendLong, whence int) int {
	return _phpStreamSeek(stream, offset, whence)
}
func PhpStreamRead(stream *PhpStream, buf *byte, count int) ssize_t {
	return _phpStreamRead(stream, buf, count)
}

func PhpStreamReadStr(stream *PhpStream, count int) *string {
	var buf = make([]byte, count)
	var size int = PhpStreamRead(stream, buf, count)
	if size < 0 {
		return nil
	}
	str := string(buf[:size])
	return &str
}

func PhpStreamWriteString(stream *PhpStream, str string) ssize_t {
	return _phpStreamWrite(stream, str, strlen(str))
}
func PhpStreamWrite(stream *PhpStream, buf *byte, count int) ssize_t {
	return _phpStreamWrite(stream, buf, count)
}
func PhpStreamFillReadBuffer(stream *PhpStream, size int) int {
	return _phpStreamFillReadBuffer(stream, size)
}
func PhpStreamEof(stream *PhpStream) int   { return _phpStreamEof(stream) }
func PhpStreamGetc(stream *PhpStream) int  { return _phpStreamGetc(stream) }
func PhpStreamFlush(stream *PhpStream) int { return _phpStreamFlush(stream, 0) }
func PhpStreamGets(stream *PhpStream, buf *byte, maxlen int) *byte {
	return _phpStreamGetLine(stream, buf, maxlen, nil)
}
func PhpStreamGetLine(stream *PhpStream, buf *byte, maxlen int, retlen *int) *byte {
	return _phpStreamGetLine(stream, buf, maxlen, retlen)
}

func PhpStreamGetLineStr(stream *PhpStream, maxlen int) *string {
	var buf = make([]byte, maxlen)
	var retlen int
	ret := PhpStreamGetLine(stream, buf, maxlen, &retlen)
	if ret == nil {
		return nil
	}
	str := string(buf[:retlen])
	return &str
}

func PhpStreamStat(stream *PhpStream, ssb *PhpStreamStatbuf) int { return _phpStreamStat(stream, ssb) }
func PhpStreamStatPathEx(path *byte, flags int, ssb *PhpStreamStatbuf, context *PhpStreamContext) int {
	return _phpStreamStatPath(path, flags, ssb, context)
}
func PhpStreamMkdir(path *byte, mode int, options int, context *PhpStreamContext) int {
	return _phpStreamMkdir(path, mode, options, context)
}
func PhpStreamRmdir(path *byte, options int, context *PhpStreamContext) int {
	return _phpStreamRmdir(path, options, context)
}
func PhpStreamOpendir(path *byte, options int, context *PhpStreamContext) *PhpStream {
	return _phpStreamOpendir(path, options, context)
}
func PhpStreamReaddir(dirstream *PhpStream, dirent PhpStreamDirent) *PhpStreamDirent {
	return _phpStreamReaddir(dirstream, dirent)
}
func PhpStreamRewinddir(dirstream *PhpStream) int { return PhpStreamRewind(dirstream) }
func PhpStreamScandir(dirname *byte, namelist ***types.String, context *PhpStreamContext, compare any) int {
	return _phpStreamScandir(dirname, namelist, 0, context, compare)
}
func PhpStreamSetOption(stream *PhpStream, option int, value int, ptrvalue any) int {
	return _phpStreamSetOption(stream, option, value, ptrvalue)
}
func PhpStreamSetChunkSize(stream *PhpStream, size int) int {
	return _phpStreamSetOption(stream, PHP_STREAM_OPTION_SET_CHUNK_SIZE, size, nil)
}
func PhpStreamSupportsLock(stream *PhpStream) int {
	if _phpStreamSetOption(stream, PHP_STREAM_OPTION_LOCKING, 0, any(PHP_STREAM_LOCK_SUPPORTED)) == 0 {
		return 1
	} else {
		return 0
	}
}
func PhpStreamLock(stream *PhpStream, mode int) int {
	return _phpStreamSetOption(stream, PHP_STREAM_OPTION_LOCKING, mode, any(nil))
}
func PhpStreamTruncateSupported(stream *PhpStream) int {
	if _phpStreamSetOption(stream, PHP_STREAM_OPTION_TRUNCATE_API, PHP_STREAM_TRUNCATE_SUPPORTED, nil) == PHP_STREAM_OPTION_RETURN_OK {
		return 1
	} else {
		return 0
	}
}
func PhpStreamTruncateSetSize(stream *PhpStream, size zend.ZendLong) int {
	return _phpStreamTruncateSetSize(stream, size)
}
func PhpStreamPopulateMetaData(stream *PhpStream, zv *types.Zval) int {
	if _phpStreamSetOption(stream, PHP_STREAM_OPTION_META_DATA_API, 0, zv) == PHP_STREAM_OPTION_RETURN_OK {
		return 1
	} else {
		return 0
	}
}
func PhpStreamCopyToStreamEx(src *PhpStream, dest *PhpStream, maxlen zend.ZendLong, len_ *int) int {
	return _phpStreamCopyToStreamEx(src, dest, maxlen, len_)
}
func PhpStreamCopyToMem(src *PhpStream, maxlen zend.ZendLong, persistent int) *types.String {
	return _phpStreamCopyToMem(src, maxlen, persistent)
}
func PhpStreamPassthru(stream *PhpStream) ssize_t { return _phpStreamPassthru(stream) }
func PhpStreamCanCast(stream *PhpStream, as int) int {
	return _phpStreamCast(stream, as, nil, 0)
}
func PhpStreamCast(stream *PhpStream, as int, ret *any, show_err int) int {
	return _phpStreamCast(stream, as, ret, show_err)
}
func PhpStreamIs(stream *PhpStream, anops *PhpStreamOps) bool { return stream.GetOps() == anops }
func PhpStreamOpenWrapper(path *byte, mode string, options int, opened **types.String) *PhpStream {
	return _phpStreamOpenWrapperEx(path, mode, options, opened, nil)
}
func PhpStreamOpenWrapperEx(path string, mode string, options int, opened **types.String, context *PhpStreamContext) *PhpStream {
	return _phpStreamOpenWrapperEx(path, mode, options, opened, context)
}
func PhpStreamGetUrlStreamWrappersHash() map[string]*PhpStreamWrapper {
	return _phpStreamGetUrlStreamWrappersHash()
}
func PhpGetStreamFiltersHash() *types.Array { return _phpGetStreamFiltersHash() }
