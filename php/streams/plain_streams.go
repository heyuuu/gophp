package streams

import "os"

// StreamStdio
// @see: php_stream_stdio_ops
type StreamStdio struct {
	outer    *Stream
	file     *os.File
	seekable bool `get:""`
}

var _ StreamOp = (*StreamStdio)(nil)
var _ StreamBindable = (*StreamStdio)(nil)

func NewStreamStdioFromFile(file *os.File) *StreamStdio {
	return &StreamStdio{
		file:     file,
		seekable: true,
	}
}

func IsStreamStdio(stream *Stream) bool {
	_, ok := stream.Op().(*StreamStdio)
	return ok
}
func AsStreamStdio(stream *Stream) (*StreamStdio, bool) {
	streamOp, ok := stream.Op().(*StreamStdio)
	return streamOp, ok
}

func (s *StreamStdio) Bind(stream *Stream) { s.outer = stream }
func (s *StreamStdio) Label() string       { return "STDIO" }
func (s *StreamStdio) Flag() OpFlag        { return OpAll }
func (s *StreamStdio) Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *StreamStdio) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *StreamStdio) Close() error {
	//TODO implement me
	panic("implement me")
}

func (s *StreamStdio) CloseHandle() {
	//TODO implement me
	panic("implement me")
}

func (s *StreamStdio) Flush() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StreamStdio) Seek(offset int64, whence int) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StreamStdio) Cast(as int) (ret any, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *StreamStdio) Stat() StreamStatInfo {
	//TODO implement me
	panic("implement me")
}

func (s *StreamStdio) SetOption(option int, value int, param any) int {
	//TODO implement me
	panic("implement me")
}
