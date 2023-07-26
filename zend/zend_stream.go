package zend

import (
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/kits/oskit"
	"os"
)

/**
 * Stream 抽象接口
 */
type IStream interface {
	Read(len_ int) []byte
	Close()
	Isatty() bool
	FileSize() int
	GetHandle() any
}

type StdStream struct {
	fp *r.File
}

func NewStdStream(fp *r.File) *StdStream {
	return &StdStream{fp: fp}
}

var _ IStream = (*StdStream)(nil)

func (s *StdStream) Isatty() bool {
	return s.fp.Isatty()
}

func (s *StdStream) Read(len_ int) []byte {
	var buf []byte
	size, _ := s.fp.Read(buf)
	return buf[:size]
}

func (s *StdStream) Close() {
	if s.fp != nil && s.fp.IsStdin() {
		s.fp.Close()
	}
}

func (s *StdStream) FileSize() int {
	if s.fp != nil {
		stat, err := s.fp.Stat()
		if err == nil {
			return int(stat.Size())
		}
	}
	return -1
}

func (s *StdStream) GetHandle() any {
	return s.fp
}

type GoFileStream struct {
	fp *os.File
}

func NewGoFileStream(fp *os.File) *GoFileStream {
	return &GoFileStream{fp: fp}
}

var _ IStream = (*GoFileStream)(nil)

func (s *GoFileStream) Isatty() bool { return oskit.Isatty(s.fp) }

func (s *GoFileStream) Read(len_ int) []byte {
	var buf []byte
	size, _ := s.fp.Read(buf)
	return buf[:size]
}

func (s *GoFileStream) Close() {
	s.fp.Close()
}

func (s *GoFileStream) FileSize() int {
	stat, err := s.fp.Stat()
	if err == nil {
		return int(stat.Size())
	}
	return -1
}

func (s *GoFileStream) GetHandle() any {
	return s.fp
}
