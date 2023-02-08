// <<generate>>

package zend

import (
	r "sik/runtime"
)

/**
 * ZendStream
 */
type ZendStream struct {
	handle any
	isatty int
	reader ZendStreamReaderT
	fsizer ZendStreamFsizerT
	closer ZendStreamCloserT
}

// func MakeZendStream(handle any, isatty int, reader ZendStreamReaderT, fsizer ZendStreamFsizerT, closer ZendStreamCloserT) ZendStream {
//     return ZendStream{
//         handle:handle,
//         isatty:isatty,
//         reader:reader,
//         fsizer:fsizer,
//         closer:closer,
//     }
// }
func (this *ZendStream) GetHandle() any                    { return this.handle }
func (this *ZendStream) SetHandle(value any)               { this.handle = value }
func (this *ZendStream) GetIsatty() int                    { return this.isatty }
func (this *ZendStream) SetIsatty(value int)               { this.isatty = value }
func (this *ZendStream) GetReader() ZendStreamReaderT      { return this.reader }
func (this *ZendStream) SetReader(value ZendStreamReaderT) { this.reader = value }
func (this *ZendStream) GetFsizer() ZendStreamFsizerT      { return this.fsizer }
func (this *ZendStream) SetFsizer(value ZendStreamFsizerT) { this.fsizer = value }
func (this *ZendStream) GetCloser() ZendStreamCloserT      { return this.closer }
func (this *ZendStream) SetCloser(value ZendStreamCloserT) { this.closer = value }

/**
 * ZendFileHandle
 */
type ZendFileHandle struct {
	handle struct /* union */ {
		fp     *r.FILE
		stream ZendStream
	}
	filename      *byte
	opened_path   *ZendString
	type_         ZendStreamType
	free_filename ZendBool
	buf           *byte
	len_          int
}

func (this *ZendFileHandle) GetFp() *r.FILE        { return this.handle.fp }
func (this *ZendFileHandle) SetFp(value *r.FILE)   { this.handle.fp = value }
func (this *ZendFileHandle) GetStream() ZendStream { return this.handle.stream }

// func (this *ZendFileHandle) SetStream(value ZendStream) { this.handle.stream = value }
func (this *ZendFileHandle) GetFilename() *byte              { return this.filename }
func (this *ZendFileHandle) SetFilename(value *byte)         { this.filename = value }
func (this *ZendFileHandle) GetOpenedPath() *ZendString      { return this.opened_path }
func (this *ZendFileHandle) SetOpenedPath(value *ZendString) { this.opened_path = value }
func (this *ZendFileHandle) GetType() ZendStreamType         { return this.type_ }
func (this *ZendFileHandle) SetType(value ZendStreamType)    { this.type_ = value }
func (this *ZendFileHandle) GetFreeFilename() ZendBool       { return this.free_filename }

// func (this *ZendFileHandle) SetFreeFilename(value ZendBool) { this.free_filename = value }
func (this *ZendFileHandle) GetBuf() *byte      { return this.buf }
func (this *ZendFileHandle) SetBuf(value *byte) { this.buf = value }
func (this *ZendFileHandle) GetLen() int        { return this.len_ }
func (this *ZendFileHandle) SetLen(value int)   { this.len_ = value }
