// <<generate>>

package zend

import (
	"bytes"
	b "sik/builtin"
	r "sik/runtime"
)

/**
 * ZendStream
 */
type ZendStream struct {
	handle *any
	isatty int
	reader ZendStreamReaderT
	fsizer ZendStreamFsizerT
	closer ZendStreamCloserT
}

var _ IStream = (*ZendStream)(nil)

func MakeZendStream(handle any, isatty int, reader ZendStreamReaderT, fsizer ZendStreamFsizerT, closer ZendStreamCloserT) ZendStream {
	return ZendStream{handle: handle, isatty: isatty, reader: reader, fsizer: fsizer, closer: closer}
}

func (z *ZendStream) Read(len_ int) []byte {
	buf := make([]byte, len_)
	size := z.reader(z.handle, buf, len_)
	return buf[:size]
}

func (z *ZendStream) Close() {
	if z.closer != nil && z.handle != nil {
		z.closer(z.handle)
	}
	z.handle = nil
}

func (z *ZendStream) Isatty() bool {
	return z.isatty != 0
}

func (z *ZendStream) FileSize() int {
	return z.fsizer(z.handle)
}

/**
 * ZendFileHandle
 */
type ZendFileHandle struct {
	// union 二选一，待合并为 IStream
	fp     *r.FILE
	stream ZendStream

	// 属性
	type_      ZendStreamType
	filename   string
	openedPath string
	buf        []byte
}

func (this *ZendFileHandle) reset() {
	this.fp = nil
	this.stream = nil
	this.type_ = 0
	this.filename = ""
	this.openedPath = ""
	this.buf = nil
}

func (this *ZendFileHandle) FileSize() int {
	ZEND_ASSERT(this.type_ == ZEND_HANDLE_STREAM)
	if this.stream.Isatty() {
		return 0
	}
	return this.stream.FileSize()
}

func (this *ZendFileHandle) InitFp(fp *r.FILE, filename string) {
	this.reset()
	this.type_ = ZEND_HANDLE_FP
	this.fp = fp
	this.filename = filename
}

func (this *ZendFileHandle) InitFilename(filename string) {
	this.reset()
	this.type_ = ZEND_HANDLE_FILENAME
	this.filename = filename
}

func (this *ZendFileHandle) Open(filename string) bool {
	if ZendStreamOpenFunction != nil {
		return ZendStreamOpenFunctionEx(filename, this)
	}

	var openedPath string
	fp := ZendFopen(filename, &openedPath)
	this.InitFp(fp, filename)
	this.openedPath = openedPath

	return this.fp != nil
}

func (this *ZendFileHandle) GetC() (byte, bool) {
	var buf byte
	if this.stream.reader(this.stream.handle, &buf, 0) {
		return buf, true
	}
	return 0, false
}

func (this *ZendFileHandle) ReadAll() []byte {
	var buf bytes.Buffer
	chunkSize := 4 * 1024
	for {
		if chunk := this.Read(chunkSize); len(chunk) > 0 {
			buf.Write(chunk)
		} else {
			break
		}
	}
	return buf.Bytes()
}

func (this *ZendFileHandle) Read(len_ int) []byte {
	if this.stream.isatty != 0 {
		var buf bytes.Buffer
		for n := 0; n < len_; n++ {
			if c, ok := this.GetC(); ok {
				buf.WriteByte(c)
				if c == '\n' {
					break
				}
			} else {
				// EOF
				break
			}
		}
		return buf.Bytes()
	}

	return this.stream.Read(len_)
}

func (this *ZendFileHandle) Fixup() ([]byte, bool) {
	if this.buf != nil {
		return this.buf, true
	}

	if this.type_ == ZEND_HANDLE_FILENAME {
		if ok := this.Open(this.filename); !ok {
			return nil, false
		}
	}
	if this.type_ == ZEND_HANDLE_FP {
		if this.fp != nil {
			return nil, false
		}
		this.type_ = ZEND_HANDLE_STREAM
		this.stream = MakeZendStream(
			this.fp,
			Isatty(this.fp),
			ZendStreamReaderT(ZendStreamStdioReader),
			ZendStreamCloserT(ZendStreamStdioCloser),
			ZendStreamFsizerT(ZendStreamStdioFsizer),
		)
	}
	fileSize := this.FileSize()
	if fileSize == -1 {
		return nil, false
	}
	if fileSize != 0 {
		buf := this.Read(fileSize)
		if len(buf) == 0 {
			return nil, false
		}
		this.buf = buf
	} else {
		buf := this.ReadAll()
		if len(buf) == 0 {
			return nil, false
		}
		this.buf = buf
	}

	return this.buf, true
}

func (this *ZendFileHandle) Destroy() {
	switch this.type_ {
	case ZEND_HANDLE_FP:
		r.Fclose(this.fp)
	case ZEND_HANDLE_STREAM:
		this.stream.Close()
	case ZEND_HANDLE_FILENAME:
		/* We're only supposed to get here when destructing the used_files hash,
		 * which doesn't really contain open files, but references to their names/paths
		 */
	}

	this.openedPath = ""
	this.buf = nil
	this.filename = ""
}

func IsFileHandlesEquals(fh1 *ZendFileHandle, fh2 *ZendFileHandle) bool {
	if fh1.type_ != fh2.type_ {
		return false
	}
	switch fh1.type_ {
	case ZEND_HANDLE_FILENAME:
		return fh1.filename == fh2.filename
	case ZEND_HANDLE_FP:
		return fh1.fp == fh2.fp
	case ZEND_HANDLE_STREAM:
		return fh1.stream.handle == fh2.stream.handle
	default:
		return false
	}
}

/**
 * generate
 */

func (this *ZendFileHandle) GetFp() *r.FILE     { return this.fp }
func (this *ZendFileHandle) GetStream() IStream { return &this.stream }
func (this *ZendFileHandle) SetStream(stream ZendStream) {
	this.stream = stream
}

func (this *ZendFileHandle) GetFilenameStr() string { return b.CastStrAuto(this.filename) }

func (this *ZendFileHandle) GetFilename() *byte              { return this.filename }
func (this *ZendFileHandle) SetFilename(value *byte)         { this.filename = value }
func (this *ZendFileHandle) GetOpenedPath() *ZendString      { return this.openedPath }
func (this *ZendFileHandle) SetOpenedPath(value *ZendString) { this.openedPath = value }
func (this *ZendFileHandle) GetType() ZendStreamType         { return this.type_ }
func (this *ZendFileHandle) SetType(value ZendStreamType)    { this.type_ = value }
