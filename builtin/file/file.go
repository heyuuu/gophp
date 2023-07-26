package file

import (
	"errors"
	"github.com/heyuuu/gophp/kits/oskit"
	"io"
	"io/fs"
	"os"
)

type handle interface {
	fs.File
	io.ReadWriteCloser
	io.Seeker
	Sync() error
}

/**
 * FILE
 */
type File struct {
	h   handle
	eof bool
}

var _ handle = (*File)(nil)

func NewFile(file handle) *File { return &File{h: file} }

func (f *File) Stat() (fs.FileInfo, error)        { return f.h.Stat() }
func (f *File) Read(bytes []byte) (int, error)    { return f.h.Read(bytes) }
func (f *File) Write(p []byte) (n int, err error) { return f.h.Write(p) }
func (f *File) Close() error                      { return f.h.Close() }

func (f *File) Eof() bool {
	return f.eof
}
func (f *File) Flush() int {
	if err := f.h.Sync(); err != nil {
		return EOF
	}
	return 0
}

func (f *File) GetC() (byte, bool) {
	buf := make([]byte, 1)
	_, err := f.Read(buf)
	if err == nil && len(buf) > 0 {
		return buf[0], true
	}
	return 0, false
}

func (f *File) WriteString(p string) (n int, err error) {
	return io.WriteString(f, p)
}

func (f *File) Isatty() bool {
	if fp, ok := f.h.(*os.File); ok {
		return oskit.Isatty(fp)
	} else {
		return false
	}
}

func (f *File) IsStdin() bool {
	return f.h == os.Stdin
}

const BUFSIZ = 1024
const EOF = -1

/* must be == _POSIX_STREAM_MAX <limits.h> */

const SEEK_SET = 0
const SEEK_CUR = 1
const SEEK_END = 2

var modeMap = map[string]int{
	"r":   os.O_RDONLY,
	"r+":  os.O_RDWR,
	"rb+": os.O_RDWR,
	"rt+": os.O_RDWR,
	"w":   os.O_WRONLY | os.O_CREATE,
	"w+":  os.O_RDWR | os.O_CREATE,
	"a":   os.O_WRONLY | os.O_CREATE | os.O_APPEND,
	"a+":  os.O_WRONLY | os.O_CREATE | os.O_APPEND,
}

func Fopen(filename string, mode string) (*File, error) {
	flags, ok := modeMap[mode]
	if !ok {
		return nil, errors.New("mode is not supported: " + mode)
	}

	f, err := os.OpenFile(filename, flags, 0)
	if err != nil {
		return nil, err
	}
	return NewFile(f), nil
}

var Fseek func(*File, int64, int) int
var Ftell func(*File) int64
var Setvbuf func(*File, *byte, int, int) int
