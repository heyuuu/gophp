// <<generate>>

package core

import (
	"sik/core/streams"
	r "sik/runtime"
	"sik/zend"
)

/**
 * PhpStreamStatbuf
 */
type PhpStreamStatbuf struct {
	sb zend.ZendStatT
}

// func MakePhpStreamStatbuf(sb zend.ZendStatT) PhpStreamStatbuf {
//     return PhpStreamStatbuf{
//         sb:sb,
//     }
// }
func (this *PhpStreamStatbuf) GetSb() zend.ZendStatT { return this.sb }

// func (this *PhpStreamStatbuf) SetSb(value zend.ZendStatT) { this.sb = value }

/**
 * PhpStreamDirent
 */
type PhpStreamDirent struct {
	d_name []byte
}

// func MakePhpStreamDirent(d_name []byte) PhpStreamDirent {
//     return PhpStreamDirent{
//         d_name:d_name,
//     }
// }
func (this *PhpStreamDirent) GetDName() []byte { return this.d_name }

// func (this *PhpStreamDirent) SetDName(value []byte) { this.d_name = value }

/**
 * PhpStreamOps
 */
type PhpStreamOps struct {
	write      func(stream *PhpStream, buf *byte, count int) ssize_t
	read       func(stream *PhpStream, buf *byte, count int) ssize_t
	close      func(stream *PhpStream, close_handle int) int
	flush      func(stream *PhpStream) int
	label      *byte
	seek       func(stream *PhpStream, offset zend.ZendOffT, whence int, newoffset *zend.ZendOffT) int
	cast       func(stream *PhpStream, castas int, ret *any) int
	stat       func(stream *PhpStream, ssb *PhpStreamStatbuf) int
	set_option func(stream *PhpStream, option int, value int, ptrparam any) int
}

func MakePhpStreamOps(
	write func(stream *PhpStream, buf *byte, count int) ssize_t,
	read func(stream *PhpStream, buf *byte, count int) ssize_t,
	close func(stream *PhpStream, close_handle int) int,
	flush func(stream *PhpStream) int,
	label *byte,
	seek func(stream *PhpStream, offset zend.ZendOffT, whence int, newoffset *zend.ZendOffT) int,
	cast func(stream *PhpStream, castas int, ret *any) int,
	stat func(stream *PhpStream, ssb *PhpStreamStatbuf) int,
	set_option func(stream *PhpStream, option int, value int, ptrparam any) int,
) PhpStreamOps {
	return PhpStreamOps{
		write:      write,
		read:       read,
		close:      close,
		flush:      flush,
		label:      label,
		seek:       seek,
		cast:       cast,
		stat:       stat,
		set_option: set_option,
	}
}
func (this *PhpStreamOps) GetWrite() func(stream *PhpStream, buf *byte, count int) ssize_t {
	return this.write
}

// func (this *PhpStreamOps) SetWrite(value func(stream *PhpStream, buf *byte, count int) ssize_t) { this.write = value }
func (this *PhpStreamOps) GetRead() func(stream *PhpStream, buf *byte, count int) ssize_t {
	return this.read
}

// func (this *PhpStreamOps) SetRead(value func(stream *PhpStream, buf *byte, count int) ssize_t) { this.read = value }
func (this *PhpStreamOps) GetClose() func(stream *PhpStream, close_handle int) int { return this.close }

// func (this *PhpStreamOps) SetClose(value func(stream *PhpStream, close_handle int) int) { this.close = value }
func (this *PhpStreamOps) GetFlush() func(stream *PhpStream) int { return this.flush }

// func (this *PhpStreamOps) SetFlush(value func(stream *PhpStream) int) { this.flush = value }
func (this *PhpStreamOps) GetLabel() *byte { return this.label }

// func (this *PhpStreamOps) SetLabel(value *byte) { this.label = value }
func (this *PhpStreamOps) GetSeek() func(stream *PhpStream, offset zend.ZendOffT, whence int, newoffset *zend.ZendOffT) int {
	return this.seek
}

// func (this *PhpStreamOps) SetSeek(value func(stream *PhpStream, offset zend.ZendOffT, whence int, newoffset *zend.ZendOffT) int) { this.seek = value }
func (this *PhpStreamOps) GetCast() func(stream *PhpStream, castas int, ret *any) int {
	return this.cast
}

// func (this *PhpStreamOps) SetCast(value func(stream *PhpStream, castas int, ret *any) int) { this.cast = value }
func (this *PhpStreamOps) GetStat() func(stream *PhpStream, ssb *PhpStreamStatbuf) int {
	return this.stat
}

// func (this *PhpStreamOps) SetStat(value func(stream *PhpStream, ssb *PhpStreamStatbuf) int) { this.stat = value }
func (this *PhpStreamOps) GetSetOption() func(stream *PhpStream, option int, value int, ptrparam any) int {
	return this.set_option
}

// func (this *PhpStreamOps) SetSetOption(value func(stream *PhpStream, option int, value int, ptrparam any) int) { this.set_option = value }

/**
 * PhpStreamWrapperOps
 */
type PhpStreamWrapperOps struct {
	stream_opener func(
		wrapper *PhpStreamWrapper,
		filename *byte,
		mode *byte,
		options int,
		opened_path **zend.ZendString,
		context *PhpStreamContext,
	) *PhpStream
	stream_closer func(wrapper *PhpStreamWrapper, stream *PhpStream) int
	stream_stat   func(wrapper *PhpStreamWrapper, stream *PhpStream, ssb *PhpStreamStatbuf) int
	url_stat      func(wrapper *PhpStreamWrapper, url *byte, flags int, ssb *PhpStreamStatbuf, context *PhpStreamContext) int
	dir_opener    func(
		wrapper *PhpStreamWrapper,
		filename *byte,
		mode *byte,
		options int,
		opened_path **zend.ZendString,
		context *PhpStreamContext,
	) *PhpStream
	label           *byte
	unlink          func(wrapper *PhpStreamWrapper, url *byte, options int, context *PhpStreamContext) int
	rename          func(wrapper *PhpStreamWrapper, url_from *byte, url_to *byte, options int, context *PhpStreamContext) int
	stream_mkdir    func(wrapper *PhpStreamWrapper, url *byte, mode int, options int, context *PhpStreamContext) int
	stream_rmdir    func(wrapper *PhpStreamWrapper, url *byte, options int, context *PhpStreamContext) int
	stream_metadata func(wrapper *PhpStreamWrapper, url *byte, options int, value any, context *PhpStreamContext) int
}

func MakePhpStreamWrapperOps(
	stream_opener func(
		wrapper *PhpStreamWrapper,
		filename *byte,
		mode *byte,
		options int,
		opened_path **zend.ZendString,
		context *PhpStreamContext,
	) *PhpStream,
	stream_closer func(wrapper *PhpStreamWrapper, stream *PhpStream) int,
	stream_stat func(wrapper *PhpStreamWrapper, stream *PhpStream, ssb *PhpStreamStatbuf) int,
	url_stat func(wrapper *PhpStreamWrapper, url *byte, flags int, ssb *PhpStreamStatbuf, context *PhpStreamContext) int,
	dir_opener func(
		wrapper *PhpStreamWrapper,
		filename *byte,
		mode *byte,
		options int,
		opened_path **zend.ZendString,
		context *PhpStreamContext,
	) *PhpStream,
	label *byte,
	unlink func(wrapper *PhpStreamWrapper, url *byte, options int, context *PhpStreamContext) int,
	rename func(wrapper *PhpStreamWrapper, url_from *byte, url_to *byte, options int, context *PhpStreamContext) int,
	stream_mkdir func(wrapper *PhpStreamWrapper, url *byte, mode int, options int, context *PhpStreamContext) int,
	stream_rmdir func(wrapper *PhpStreamWrapper, url *byte, options int, context *PhpStreamContext) int,
	stream_metadata func(wrapper *PhpStreamWrapper, url *byte, options int, value any, context *PhpStreamContext) int,
) PhpStreamWrapperOps {
	return PhpStreamWrapperOps{
		stream_opener:   stream_opener,
		stream_closer:   stream_closer,
		stream_stat:     stream_stat,
		url_stat:        url_stat,
		dir_opener:      dir_opener,
		label:           label,
		unlink:          unlink,
		rename:          rename,
		stream_mkdir:    stream_mkdir,
		stream_rmdir:    stream_rmdir,
		stream_metadata: stream_metadata,
	}
}
func (this *PhpStreamWrapperOps) GetStreamOpener() func(
	wrapper *PhpStreamWrapper,
	filename *byte,
	mode *byte,
	options int,
	opened_path **zend.ZendString,
	context *PhpStreamContext,
) *PhpStream {
	return this.stream_opener
}

func (this *PhpStreamWrapperOps) SetStreamOpener(value func(
	wrapper *PhpStreamWrapper,
	filename *byte,
	mode *byte,
	options int,
	opened_path **zend.ZendString,
	context *PhpStreamContext,
) *PhpStream) {
	this.stream_opener = value
}
func (this *PhpStreamWrapperOps) GetStreamCloser() func(wrapper *PhpStreamWrapper, stream *PhpStream) int {
	return this.stream_closer
}

// func (this *PhpStreamWrapperOps) SetStreamCloser(value func(wrapper *PhpStreamWrapper, stream *PhpStream) int) { this.stream_closer = value }
func (this *PhpStreamWrapperOps) GetStreamStat() func(wrapper *PhpStreamWrapper, stream *PhpStream, ssb *PhpStreamStatbuf) int {
	return this.stream_stat
}

// func (this *PhpStreamWrapperOps) SetStreamStat(value func(wrapper *PhpStreamWrapper, stream *PhpStream, ssb *PhpStreamStatbuf) int) { this.stream_stat = value }
func (this *PhpStreamWrapperOps) GetUrlStat() func(wrapper *PhpStreamWrapper, url *byte, flags int, ssb *PhpStreamStatbuf, context *PhpStreamContext) int {
	return this.url_stat
}

// func (this *PhpStreamWrapperOps) SetUrlStat(value func(wrapper *PhpStreamWrapper, url *byte, flags int, ssb *PhpStreamStatbuf, context *PhpStreamContext) int) { this.url_stat = value }
func (this *PhpStreamWrapperOps) GetDirOpener() func(
	wrapper *PhpStreamWrapper,
	filename *byte,
	mode *byte,
	options int,
	opened_path **zend.ZendString,
	context *PhpStreamContext,
) *PhpStream {
	return this.dir_opener
}

func (this *PhpStreamWrapperOps) SetDirOpener(value func(
	wrapper *PhpStreamWrapper,
	filename *byte,
	mode *byte,
	options int,
	opened_path **zend.ZendString,
	context *PhpStreamContext,
) *PhpStream) {
	this.dir_opener = value
}
func (this *PhpStreamWrapperOps) GetLabel() *byte { return this.label }

// func (this *PhpStreamWrapperOps) SetLabel(value *byte) { this.label = value }
func (this *PhpStreamWrapperOps) GetUnlink() func(wrapper *PhpStreamWrapper, url *byte, options int, context *PhpStreamContext) int {
	return this.unlink
}

// func (this *PhpStreamWrapperOps) SetUnlink(value func(wrapper *PhpStreamWrapper, url *byte, options int, context *PhpStreamContext) int) { this.unlink = value }
func (this *PhpStreamWrapperOps) GetRename() func(wrapper *PhpStreamWrapper, url_from *byte, url_to *byte, options int, context *PhpStreamContext) int {
	return this.rename
}

// func (this *PhpStreamWrapperOps) SetRename(value func(wrapper *PhpStreamWrapper, url_from *byte, url_to *byte, options int, context *PhpStreamContext) int) { this.rename = value }
func (this *PhpStreamWrapperOps) GetStreamMkdir() func(wrapper *PhpStreamWrapper, url *byte, mode int, options int, context *PhpStreamContext) int {
	return this.stream_mkdir
}

// func (this *PhpStreamWrapperOps) SetStreamMkdir(value func(wrapper *PhpStreamWrapper, url *byte, mode int, options int, context *PhpStreamContext) int) { this.stream_mkdir = value }
func (this *PhpStreamWrapperOps) GetStreamRmdir() func(wrapper *PhpStreamWrapper, url *byte, options int, context *PhpStreamContext) int {
	return this.stream_rmdir
}

// func (this *PhpStreamWrapperOps) SetStreamRmdir(value func(wrapper *PhpStreamWrapper, url *byte, options int, context *PhpStreamContext) int) { this.stream_rmdir = value }
func (this *PhpStreamWrapperOps) GetStreamMetadata() func(wrapper *PhpStreamWrapper, url *byte, options int, value any, context *PhpStreamContext) int {
	return this.stream_metadata
}

// func (this *PhpStreamWrapperOps) SetStreamMetadata(value func(wrapper *PhpStreamWrapper, url *byte, options int, value any, context *PhpStreamContext) int) { this.stream_metadata = value }

/**
 * PhpStreamWrapper
 */
type PhpStreamWrapper struct {
	wops     *PhpStreamWrapperOps
	abstract any
	is_url   int
}

func MakePhpStreamWrapper(wops *PhpStreamWrapperOps, abstract any, is_url int) PhpStreamWrapper {
	return PhpStreamWrapper{
		wops:     wops,
		abstract: abstract,
		is_url:   is_url,
	}
}
func (this *PhpStreamWrapper) GetWops() *PhpStreamWrapperOps      { return this.wops }
func (this *PhpStreamWrapper) SetWops(value *PhpStreamWrapperOps) { this.wops = value }
func (this *PhpStreamWrapper) GetAbstract() any                   { return this.abstract }
func (this *PhpStreamWrapper) SetAbstract(value any)              { this.abstract = value }
func (this *PhpStreamWrapper) GetIsUrl() int                      { return this.is_url }
func (this *PhpStreamWrapper) SetIsUrl(value int)                 { this.is_url = value }

/**
 * PhpStream
 */
type PhpStream struct {
	ops              *PhpStreamOps
	abstract         any
	readfilters      streams.PhpStreamFilterChain
	writefilters     streams.PhpStreamFilterChain
	wrapper          *PhpStreamWrapper
	wrapperthis      any
	wrapperdata      zend.Zval
	is_persistent    uint8
	in_free          uint8
	eof              uint8
	__exposed        uint8
	fclose_stdiocast uint8
	fgetss_state     uint8
	mode             []byte
	flags            uint32
	res              *zend.ZendResource
	stdiocast        *r.FILE
	orig_path        *byte
	ctx              *zend.ZendResource
	position         zend.ZendOffT
	readbuf          *uint8
	readbuflen       int
	readpos          zend.ZendOffT
	writepos         zend.ZendOffT
	chunk_size       int
	enclosing_stream *PhpStream
}

//             func MakePhpStream(
// ops *PhpStreamOps,
// abstract any,
// readfilters streams.PhpStreamFilterChain,
// writefilters streams.PhpStreamFilterChain,
// wrapper *PhpStreamWrapper,
// wrapperthis any,
// wrapperdata zend.Zval,
// is_persistent uint8,
// in_free uint8,
// eof uint8,
// __exposed uint8,
// fclose_stdiocast uint8,
// fgetss_state uint8,
// mode []byte,
// flags uint32,
// res *zend.ZendResource,
// stdiocast *r.FILE,
// orig_path *byte,
// ctx *zend.ZendResource,
// position zend.ZendOffT,
// readbuf *uint8,
// readbuflen int,
// readpos zend.ZendOffT,
// writepos zend.ZendOffT,
// chunk_size int,
// enclosing_stream *PhpStream,
// ) PhpStream {
//                 return PhpStream{
//                     ops:ops,
//                     abstract:abstract,
//                     readfilters:readfilters,
//                     writefilters:writefilters,
//                     wrapper:wrapper,
//                     wrapperthis:wrapperthis,
//                     wrapperdata:wrapperdata,
//                     is_persistent:is_persistent,
//                     in_free:in_free,
//                     eof:eof,
//                     __exposed:__exposed,
//                     fclose_stdiocast:fclose_stdiocast,
//                     fgetss_state:fgetss_state,
//                     mode:mode,
//                     flags:flags,
//                     res:res,
//                     stdiocast:stdiocast,
//                     orig_path:orig_path,
//                     ctx:ctx,
//                     position:position,
//                     readbuf:readbuf,
//                     readbuflen:readbuflen,
//                     readpos:readpos,
//                     writepos:writepos,
//                     chunk_size:chunk_size,
//                     enclosing_stream:enclosing_stream,
//                 }
//             }
func (this *PhpStream) GetOps() *PhpStreamOps                        { return this.ops }
func (this *PhpStream) SetOps(value *PhpStreamOps)                   { this.ops = value }
func (this *PhpStream) GetAbstract() any                             { return this.abstract }
func (this *PhpStream) SetAbstract(value any)                        { this.abstract = value }
func (this *PhpStream) GetReadfilters() streams.PhpStreamFilterChain { return this.readfilters }

// func (this *PhpStream) SetReadfilters(value streams.PhpStreamFilterChain) { this.readfilters = value }
func (this *PhpStream) GetWritefilters() streams.PhpStreamFilterChain { return this.writefilters }

// func (this *PhpStream) SetWritefilters(value streams.PhpStreamFilterChain) { this.writefilters = value }
func (this *PhpStream) GetWrapper() *PhpStreamWrapper      { return this.wrapper }
func (this *PhpStream) SetWrapper(value *PhpStreamWrapper) { this.wrapper = value }
func (this *PhpStream) GetWrapperthis() any                { return this.wrapperthis }
func (this *PhpStream) SetWrapperthis(value any)           { this.wrapperthis = value }
func (this *PhpStream) GetWrapperdata() zend.Zval          { return this.wrapperdata }

// func (this *PhpStream) SetWrapperdata(value zend.Zval) { this.wrapperdata = value }
func (this *PhpStream) GetIsPersistent() uint8      { return this.is_persistent }
func (this *PhpStream) SetIsPersistent(value uint8) { this.is_persistent = value }
func (this *PhpStream) GetInFree() uint8            { return this.in_free }
func (this *PhpStream) SetInFree(value uint8)       { this.in_free = value }
func (this *PhpStream) GetEof() uint8               { return this.eof }
func (this *PhpStream) SetEof(value uint8)          { this.eof = value }

// func (this *PhpStream)  GetExposed() uint8      { return this.__exposed }
func (this *PhpStream) SetExposed(value uint8)         { this.__exposed = value }
func (this *PhpStream) GetFcloseStdiocast() uint8      { return this.fclose_stdiocast }
func (this *PhpStream) SetFcloseStdiocast(value uint8) { this.fclose_stdiocast = value }
func (this *PhpStream) GetFgetssState() uint8          { return this.fgetss_state }

// func (this *PhpStream) SetFgetssState(value uint8) { this.fgetss_state = value }
func (this *PhpStream) GetMode() []byte { return this.mode }

// func (this *PhpStream) SetMode(value []byte) { this.mode = value }
func (this *PhpStream) GetFlags() uint32                    { return this.flags }
func (this *PhpStream) SetFlags(value uint32)               { this.flags = value }
func (this *PhpStream) GetRes() *zend.ZendResource          { return this.res }
func (this *PhpStream) SetRes(value *zend.ZendResource)     { this.res = value }
func (this *PhpStream) GetStdiocast() *r.FILE               { return this.stdiocast }
func (this *PhpStream) SetStdiocast(value *r.FILE)          { this.stdiocast = value }
func (this *PhpStream) GetOrigPath() *byte                  { return this.orig_path }
func (this *PhpStream) SetOrigPath(value *byte)             { this.orig_path = value }
func (this *PhpStream) GetCtx() *zend.ZendResource          { return this.ctx }
func (this *PhpStream) SetCtx(value *zend.ZendResource)     { this.ctx = value }
func (this *PhpStream) GetPosition() zend.ZendOffT          { return this.position }
func (this *PhpStream) SetPosition(value zend.ZendOffT)     { this.position = value }
func (this *PhpStream) GetReadbuf() *uint8                  { return this.readbuf }
func (this *PhpStream) SetReadbuf(value *uint8)             { this.readbuf = value }
func (this *PhpStream) GetReadbuflen() int                  { return this.readbuflen }
func (this *PhpStream) SetReadbuflen(value int)             { this.readbuflen = value }
func (this *PhpStream) GetReadpos() zend.ZendOffT           { return this.readpos }
func (this *PhpStream) SetReadpos(value zend.ZendOffT)      { this.readpos = value }
func (this *PhpStream) GetWritepos() zend.ZendOffT          { return this.writepos }
func (this *PhpStream) SetWritepos(value zend.ZendOffT)     { this.writepos = value }
func (this *PhpStream) GetChunkSize() int                   { return this.chunk_size }
func (this *PhpStream) SetChunkSize(value int)              { this.chunk_size = value }
func (this *PhpStream) GetEnclosingStream() *PhpStream      { return this.enclosing_stream }
func (this *PhpStream) SetEnclosingStream(value *PhpStream) { this.enclosing_stream = value }

/* PhpStream.flags */
func (this *PhpStream) AddFlags(value uint32)      { this.flags |= value }
func (this *PhpStream) SubFlags(value uint32)      { this.flags &^= value }
func (this *PhpStream) HasFlags(value uint32) bool { return this.flags&value != 0 }
func (this *PhpStream) SwitchFlags(value uint32, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this PhpStream) IsAvoidBlocking() bool { return this.HasFlags(PHP_STREAM_FLAG_AVOID_BLOCKING) }
func (this *PhpStream) SetIsAvoidBlocking(cond bool) {
	this.SwitchFlags(PHP_STREAM_FLAG_AVOID_BLOCKING, cond)
}
