// <<generate>>

package core

import (
	b "sik/builtin"
	"sik/core/streams"
	r "sik/runtime"
	"sik/zend"
)

func PhpStreamAllocRel(ops *PhpStreamOps, thisptr any, persistent *byte, mode string) *PhpStream {
	return _phpStreamAlloc(ops, thisptr, persistent, mode)
}
func PhpStreamCopyToMemRel(src *PhpStream, maxlen int, persistent __auto__) *zend.ZendString {
	return _phpStreamCopyToMem(src, buf, maxlen, persistent)
}
func PhpStreamFopenRel(filename *byte, mode *byte, opened **zend.ZendString, options int) *PhpStream {
	return streams._phpStreamFopen(filename, mode, opened, options)
}
func PhpStreamFopenWithPathRel(filename *byte, mode *byte, path *byte, opened **zend.ZendString, options int) *PhpStream {
	return streams._phpStreamFopenWithPath(filename, mode, path, opened, options)
}
func PhpStreamFopenFromFdRel(fd int, mode *byte, persistent_id *byte) *PhpStream {
	return streams._phpStreamFopenFromFd(fd, mode, persistent_id)
}
func PhpStreamFopenFromFileRel(file *r.FILE, mode *byte) *PhpStream {
	return streams._phpStreamFopenFromFile(file, mode)
}
func PhpStreamFopenFromPipeRel(file *r.FILE, mode *byte) *PhpStream {
	return streams._phpStreamFopenFromPipe(file, mode)
}
func PhpStreamFopenTemporaryFileRel(dir *byte, pfx *byte, opened_path **zend.ZendString) *PhpStream {
	return streams._phpStreamFopenTemporaryFile(dir, pfx, opened_path)
}
func PhpStreamOpenWrapperRel(path *byte, mode string, options int, opened **zend.ZendString) *PhpStream {
	return _phpStreamOpenWrapperEx(path, mode, options, opened, nil)
}
func PhpStreamOpenWrapperExRel(path *byte, mode *byte, options int, opened **zend.ZendString, context *PhpStreamContext) *PhpStream {
	return _phpStreamOpenWrapperEx(path, mode, options, opened, context)
}
func PhpStreamMakeSeekableRel(origstream *PhpStream, newstream **PhpStream, flags int) int {
	return _phpStreamMakeSeekable(origstream, newstream, flags)
}
func PHP_STREAM_CONTEXT(stream *PhpStream) *PhpStreamContext {
	return (*PhpStreamContext)(b.CondF1(stream.GetCtx() != nil, func() any { return stream.GetCtx().GetPtr() }, nil))
}
func PhpStreamAlloc(ops *PhpStreamOps, thisptr any, persistent_id int, mode *byte) *PhpStream {
	return _phpStreamAlloc(ops, thisptr, persistent_id, mode)
}
func PhpStreamGetResourceId(stream __auto__) int {
	return (*PhpStream)(stream).GetRes().GetHandle()
}
func PhpStreamAutoCleanup(stream *PhpStream) { stream.SetExposed(1) }
func PhpStreamToZval(stream *PhpStream, zval *zend.Zval) {
	zval.SetResource(stream.GetRes())
	stream.SetExposed(1)
}
func PhpStreamFromZval(xstr *PhpStream, pzval *zend.Zval) {
	if b.Assign(&xstr, (*PhpStream)(zend.ZendFetchResource2Ex(pzval, "stream", PhpFileLeStream(), PhpFileLePstream()))) == nil {
		zend.RETVAL_FALSE
		return
	}
}
func PhpStreamFromRes(xstr *PhpStream, res *zend.ZendResource) {
	if b.Assign(&xstr, (*PhpStream)(zend.ZendFetchResource2(res, "stream", PhpFileLeStream(), PhpFileLePstream()))) == nil {
		zend.RETVAL_FALSE
		return
	}
}
func PhpStreamFromResNoVerify(xstr *PhpStream, pzval __auto__) *PhpStream {
	xstr = (*PhpStream)(zend.ZendFetchResource2(res, "stream", PhpFileLeStream(), PhpFileLePstream()))
	return xstr
}
func PhpStreamFromZvalNoVerify(xstr *PhpStream, pzval *zend.Zval) *PhpStream {
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
func PhpStreamTell(stream *PhpStream) zend.ZendOffT { return stream.GetPosition() }
func PhpStreamRead(stream *PhpStream, buf *byte, count int) ssize_t {
	return _phpStreamRead(stream, buf, count)
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
func PhpStreamEof(stream *PhpStream) int         { return _phpStreamEof(stream) }
func PhpStreamGetc(stream *PhpStream) int        { return _phpStreamGetc(stream) }
func PhpStreamPutc(stream *PhpStream, c int) int { return _phpStreamPutc(stream, c) }
func PhpStreamFlush(stream *PhpStream) int       { return _phpStreamFlush(stream, 0) }
func PhpStreamGets(stream *PhpStream, buf *byte, maxlen int) *byte {
	return _phpStreamGetLine(stream, buf, maxlen, nil)
}
func PhpStreamGetLine(stream *PhpStream, buf *byte, maxlen int, retlen *int) *byte {
	return _phpStreamGetLine(stream, buf, maxlen, retlen)
}
func PhpStreamPuts(stream *PhpStream, buf *byte) int             { return _phpStreamPuts(stream, buf) }
func PhpStreamStat(stream *PhpStream, ssb *PhpStreamStatbuf) int { return _phpStreamStat(stream, ssb) }
func PhpStreamStatPath(path *byte, ssb *PhpStreamStatbuf) int {
	return _phpStreamStatPath(path, 0, ssb, nil)
}
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
func PhpStreamClosedir(dirstream *PhpStream) int  { return PhpStreamClose(dirstream) }
func PhpStreamRewinddir(dirstream *PhpStream) int { return PhpStreamRewind(dirstream) }
func PhpStreamScandir(dirname *byte, namelist ***zend.ZendString, context *PhpStreamContext, compare any) int {
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
func PhpStreamPopulateMetaData(stream *PhpStream, zv *zend.Zval) int {
	if _phpStreamSetOption(stream, PHP_STREAM_OPTION_META_DATA_API, 0, zv) == PHP_STREAM_OPTION_RETURN_OK {
		return 1
	} else {
		return 0
	}
}
func PhpStreamCopyToStream(src *PhpStream, dest *PhpStream, maxlen int) int {
	return _phpStreamCopyToStream(src, dest, maxlen)
}
func PhpStreamCopyToStreamEx(src *PhpStream, dest *PhpStream, maxlen zend.ZendLong, len_ *int) int {
	return _phpStreamCopyToStreamEx(src, dest, maxlen, len_)
}
func PhpStreamCopyToMem(src *PhpStream, maxlen zend.ZendLong, persistent int) *zend.ZendString {
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
func PhpStreamIsPersistent(stream *PhpStream) uint8           { return stream.GetIsPersistent() }
func PhpStreamOpenWrapper(path *byte, mode string, options int, opened **zend.ZendString) *PhpStream {
	return _phpStreamOpenWrapperEx(path, mode, options, opened, nil)
}
func PhpStreamOpenWrapperEx(path string, mode string, options int, opened **zend.ZendString, context *PhpStreamContext) *PhpStream {
	return _phpStreamOpenWrapperEx(path, mode, options, opened, context)
}
func PhpStreamMakeSeekable(origstream *PhpStream, newstream **PhpStream, flags int) int {
	return _phpStreamMakeSeekable(origstream, newstream, flags)
}
func PhpStreamGetUrlStreamWrappersHash() *zend.HashTable {
	return _phpStreamGetUrlStreamWrappersHash()
}
func PhpGetStreamFiltersHash() *zend.HashTable { return _phpGetStreamFiltersHash() }
