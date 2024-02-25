package streams

import (
	"github.com/heyuuu/gophp/php"
)

type StreamWrapperOps struct {
	Label    string
	IsUrl    bool
	Open     func(ctx *php.Context, w *StreamWrapper, filename string, mode string, options int, context *StreamContext) (stream *Stream, openedPath string)
	Close    func(ctx *php.Context, w *StreamWrapper, stream *Stream) bool
	Stat     func(ctx *php.Context, w *StreamWrapper, stream *Stream) StreamStatInfo
	UrlStat  func(ctx *php.Context, w *StreamWrapper, url string, flags int, context *StreamContext) StreamStatInfo
	DirOpen  func(ctx *php.Context, w *StreamWrapper, filename string, mode string, options int, context *StreamContext) (stream *Stream, openedPath string)
	Unlink   func(ctx *php.Context, w *StreamWrapper, url string, options int, context *StreamContext) bool
	Rename   func(ctx *php.Context, w *StreamWrapper, urlFrom string, urlTo string, options int, context *StreamContext) bool
	Mkdir    func(ctx *php.Context, w *StreamWrapper, url string, mode int, options int, context *StreamContext) bool
	Rmdir    func(ctx *php.Context, w *StreamWrapper, url string, options int, context *StreamContext) bool
	Metadata func(ctx *php.Context, w *StreamWrapper, url string, options int, value any, context *StreamContext) bool
}

// StreamWrapper
type StreamWrapper struct {
	ops StreamWrapperOps
}

func NewStreamWrapper(ops StreamWrapperOps) *StreamWrapper {
	return &StreamWrapper{ops: ops}
}

func (w *StreamWrapper) Label() string { return w.ops.Label }
func (w *StreamWrapper) IsUrl() bool   { return w.ops.IsUrl }
func (w *StreamWrapper) Open(ctx *php.Context, filename string, mode string, options int, context *StreamContext) (stream *Stream, openedPath string) {
	return w.ops.Open(ctx, w, filename, mode, options, context)
}
func (w *StreamWrapper) Close(ctx *php.Context, stream *Stream) bool {
	return w.ops.Close(ctx, w, stream)
}
func (w *StreamWrapper) Stat(ctx *php.Context, stream *Stream) StreamStatInfo {
	return w.ops.Stat(ctx, w, stream)
}
func (w *StreamWrapper) UrlStat(ctx *php.Context, url string, flags int, context *StreamContext) StreamStatInfo {
	return w.ops.UrlStat(ctx, w, url, flags, context)
}
func (w *StreamWrapper) DirOpen(ctx *php.Context, filename string, mode string, options int, context *StreamContext) (stream *Stream, openedPath string) {
	return w.ops.DirOpen(ctx, w, filename, mode, options, context)
}
func (w *StreamWrapper) Unlink(ctx *php.Context, url string, options int, context *StreamContext) bool {
	return w.ops.Unlink(ctx, w, url, options, context)
}
func (w *StreamWrapper) Rename(ctx *php.Context, urlFrom string, urlTo string, options int, context *StreamContext) bool {
	return w.ops.Rename(ctx, w, urlFrom, urlTo, options, context)
}
func (w *StreamWrapper) Mkdir(ctx *php.Context, url string, mode int, options int, context *StreamContext) bool {
	return w.ops.Mkdir(ctx, w, url, mode, options, context)
}
func (w *StreamWrapper) Rmdir(ctx *php.Context, url string, options int, context *StreamContext) bool {
	return w.ops.Rmdir(ctx, w, url, options, context)
}
func (w *StreamWrapper) Metadata(ctx *php.Context, url string, options int, value any, context *StreamContext) bool {
	return w.ops.Metadata(ctx, w, url, options, value, context)
}

// flags

func (w *StreamWrapper) HasOpen() bool     { return w.ops.Open != nil }
func (w *StreamWrapper) HasClose() bool    { return w.ops.Close != nil }
func (w *StreamWrapper) HasStat() bool     { return w.ops.Stat != nil }
func (w *StreamWrapper) HasUrlStat() bool  { return w.ops.UrlStat != nil }
func (w *StreamWrapper) HasDirOpen() bool  { return w.ops.DirOpen != nil }
func (w *StreamWrapper) HasMkdir() bool    { return w.ops.Mkdir != nil }
func (w *StreamWrapper) HasRmdir() bool    { return w.ops.Rmdir != nil }
func (w *StreamWrapper) HasMetadata() bool { return w.ops.Metadata != nil }
func (w *StreamWrapper) HasUnlink() bool   { return w.ops.Unlink != nil }
func (w *StreamWrapper) HasRename() bool   { return w.ops.Rename != nil }
