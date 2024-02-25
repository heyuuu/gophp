package streams

import (
	"errors"
	"io"
)

type OpFlag uint8

const (
	OpWritable OpFlag = 1 << iota
	OpReadable
	OpClosable
	OpFlushable
	OpSeekable
	OpCastable
	OpCanCast
	OpCanSetOption

	OpAll = OpWritable | OpReadable | OpClosable | OpFlushable | OpSeekable | OpCastable | OpCanCast | OpCanSetOption
)

var OpSimpleError = errors.New("")

// StreamOp
type StreamOp interface {
	Flag() OpFlag
	Label() string
	io.Writer
	io.Reader
	io.Closer
	CloseHandle()
	Flush() (int, error)
	io.Seeker
	Cast(as int) (ret any, err error)
	Stat() StreamStatInfo
	SetOption(option int, value int, param any) int
}

// EmptyStreamOp
type EmptyStreamOp struct{}

func (op *EmptyStreamOp) Write(p []byte) (n int, err error) { return 0, OpSimpleError }
func (op *EmptyStreamOp) Read(p []byte) (n int, err error)  { return 0, OpSimpleError }
func (op *EmptyStreamOp) Close() error                      { return nil }
func (op *EmptyStreamOp) CloseHandle()                      {}
func (op *EmptyStreamOp) Flush() (int, error)               { return 0, nil }
func (op *EmptyStreamOp) Seek(offset int64, whence int) (int64, error) {
	return 0, OpSimpleError
}
func (op *EmptyStreamOp) Cast(as int) (ret any, err error) {
	return nil, OpSimpleError
}
func (op *EmptyStreamOp) Stat() StreamStatInfo {
	return nil
}
func (op *EmptyStreamOp) SetOption(option int, value int, param any) int {
	return PHP_STREAM_OPTION_RETURN_NOTIMPL
}
