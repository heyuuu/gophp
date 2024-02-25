package standard

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/streams"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"time"
)

var LeStreamContext = php.InitResourceType("stream-context")

// -- FG(FileGlobals) --

type FileGlobals struct {
	pcloseRet                 int                    `prop:""`
	defChunkSize              int                    `prop:""`
	autoDetectLineEndings     bool                   `prop:""`
	defaultSocketTimeout      time.Duration          `prop:""`
	userAgent                 string                 `prop:""`
	fromAddress               string                 `prop:""`
	userStreamCurrentFilename string                 `prop:""`
	defaultContext            *streams.StreamContext `prop:""`
}

func NewPhpFileGlobals(chunkSize int) *FileGlobals {
	return &FileGlobals{defChunkSize: chunkSize}
}

const fileGlobalKey = "ext.standard.file_global"

func FG(ctx *php.Context) *FileGlobals {
	return php.ContextGetOrInit(ctx, fileGlobalKey, func() *FileGlobals {
		return NewPhpFileGlobals(php.PHP_SOCK_CHUNK_SIZE)
	})
}
func UnsetFG(ctx *php.Context) {
	php.ContextDel(ctx, fileGlobalKey)
}

func ZifFopen(ctx *php.Context, filename zpp.Path, mode string, _ zpp.Opt, useIncludePath bool, context_ zpp.ResourceNullable) (*types.Resource, bool) {
	context := StreamContextFromResource(ctx, context_, false)
	options := streams.REPORT_ERRORS
	if useIncludePath {
		options |= streams.USE_PATH
	}
	stream := streams.StreamOpenWrapper(ctx, filename, mode, options, context)
	if stream == nil {
		return nil, false
	}
	return stream.Resource(), true
}
