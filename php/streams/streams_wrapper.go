package streams

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/perr"
)

func StreamOpenWrapper(ctx *php.Context, path string, mode string, options int, context *StreamContext) *Stream {
	stream, _ := StreamOpenWrapperEx(ctx, path, mode, options, context)
	return stream
}

// @see: *_php_stream_open_wrapper_ex
func StreamOpenWrapperEx(ctx *php.Context, path string, mode string, options int, context *StreamContext) (stream *Stream, openedPath string) {
	if path == "" {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Filename cannot be empty")
		return nil, ""
	}

	return PhpPlainFilesWrapper.Open(ctx, path, mode, options, context)
}
