package zend

import (
	"bytes"
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"os"
)

/**
 * FileHandle
 */
type FileHandle struct {
	stream IStream

	// 属性
	filename   string
	openedPath string
	buf        []byte
}

func NewFileHandleByFilename(filename string) *FileHandle {
	return &FileHandle{
		stream:   nil,
		filename: filename,
	}
}

func NewFileHandleByFp(filename string, fp *r.File) *FileHandle {
	return &FileHandle{
		stream:   NewStdStream(fp),
		filename: filename,
	}
}

func NewFileHandleByGoFile(filename string, fp *os.File) *FileHandle {
	return &FileHandle{
		stream:   NewGoFileStream(fp),
		filename: filename,
	}
}

func NewFileHandleByStream(filename string, openedPath string, stream *core.PhpStreamForZend) *FileHandle {
	return &FileHandle{
		stream:     stream,
		filename:   filename,
		openedPath: openedPath,
	}
}

func NewFileHandleForStdin() *FileHandle {
	return NewFileHandleByGoFile("Standard input code", os.Stdin)
}

func NewFileHandleByOpenFile(filename string) *FileHandle {
	fp, err := os.Open(filename)
	if err != nil {
		return nil
	}
	return NewFileHandleByGoFile(filename, fp)
}

func NewFileHandleByOpenStream(filename string) *FileHandle {
	fh := NewFileHandleByFilename(filename)
	if fh.openStream() {
		return fh
	}
	return nil
}

func (fh *FileHandle) IsTypeHandleFileName() bool {
	return fh.stream == nil
}

func (fh *FileHandle) fileSize() int {
	if fh.stream.Isatty() {
		return 0
	}
	return fh.stream.FileSize()
}

func (fh *FileHandle) openStream() bool {
	b.Assert(fh.stream == nil)
	stream, openedPath := core.PhpStreamOpenForZend(fh.filename)
	if stream != nil {
		fh.stream = stream
		fh.openedPath = openedPath
		return true
	}
	return false
}

func (fh *FileHandle) Fixup() ([]byte, bool) {
	if fh.buf != nil {
		return fh.buf, true
	}

	if fh.stream == nil {
		if !fh.openStream() {
			return nil, false
		}
	}
	if fh.stream == nil {
		return nil, false
	}
	fileSize := fh.fileSize()
	if fileSize == -1 {
		return nil, false
	}
	if fileSize != 0 {
		buf := fh.read(fileSize)
		if len(buf) == 0 {
			return nil, false
		}
		fh.buf = buf
	} else {
		buf := fh.readAll()
		if len(buf) == 0 {
			return nil, false
		}
		fh.buf = buf
	}

	return fh.buf, true
}

func (fh *FileHandle) readAll() []byte {
	var buf bytes.Buffer
	chunkSize := 4 * 1024
	for {
		if chunk := fh.read(chunkSize); len(chunk) > 0 {
			buf.Write(chunk)
		} else {
			break
		}
	}
	return buf.Bytes()
}

func (fh *FileHandle) readByte() (byte, bool) {
	cs := fh.stream.Read(1)
	if len(cs) > 0 {
		return cs[0], true
	}
	return 0, false
}

func (fh *FileHandle) read(len_ int) []byte {
	if fh.stream.Isatty() {
		var buf bytes.Buffer
		for n := 0; n < len_; n++ {
			if c, ok := fh.readByte(); ok {
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

	return fh.stream.Read(len_)
}

func (fh *FileHandle) Close() {
	if fh.stream != nil {
		fh.stream.Close()
	}
}

func (fh *FileHandle) Destroy() {
	if fh.stream != nil {
		fh.stream.Close()
	}

	fh.openedPath = ""
	fh.buf = nil
	fh.filename = ""
}

func IsFileHandlesEquals(fh1 *FileHandle, fh2 *FileHandle) bool {
	if fh1.stream != nil && fh2.stream != nil {
		return fh1.stream.GetHandle() == fh2.stream.GetHandle()
	} else if fh1.filename != "" && fh2.filename != "" {
		return fh1.filename == fh2.filename
	} else {
		return false
	}
}

func (fh *FileHandle) GetStream() IStream         { return fh.stream }
func (fh *FileHandle) GetFilename() string        { return fh.filename }
func (fh *FileHandle) GetOpenedPath() string      { return fh.openedPath }
func (fh *FileHandle) SetOpenedPath(value string) { fh.openedPath = value }
