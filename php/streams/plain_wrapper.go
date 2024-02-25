package streams

import (
	"github.com/heyuuu/gophp/php"
	"os"
	"strings"
)

// PlainFilesWrapper
var PhpPlainFilesWrapper = NewStreamWrapper(StreamWrapperOps{
	Label: "plainfile",
	Open:  plainFilesOpen,
	//UrlStat:  plainFilesUrlStat,
	//DirOpen:  plainFilesDirOpen,
	//Unlink:   plainFilesUnlink,
	//Rename:   plainFilesRename,
	//Mkdir:    plainFilesMkdir,
	//Rmdir:    plainFilesRmdir,
	//Metadata: plainFilesMetadata,
})

func IsPlainFilesWrapper(wrapper *StreamWrapper) bool {
	return wrapper == PhpPlainFilesWrapper
}

var plainFopenModes = [2]map[byte]int{
	{
		'r': os.O_RDONLY,
		'w': os.O_WRONLY | os.O_CREATE | os.O_TRUNC,
		'a': os.O_WRONLY | os.O_CREATE | os.O_APPEND,
		'x': os.O_WRONLY | os.O_CREATE | os.O_EXCL,
		'c': os.O_WRONLY | os.O_CREATE,
	},
	{
		'r': os.O_RDWR,
		'w': os.O_RDWR | os.O_CREATE | os.O_TRUNC,
		'a': os.O_RDWR | os.O_CREATE | os.O_APPEND,
		'x': os.O_RDWR | os.O_CREATE | os.O_EXCL,
		'c': os.O_RDWR | os.O_CREATE,
	},
}

func plainParseFopenMode(mode string) (flags int, ok bool) {
	if mode == "" {
		return 0, false
	}
	var c = mode[0]
	switch c {
	case 'r', 'w', 'a', 'x', 'c':
		if strings.IndexByte(mode, '+') < 0 {
			return plainFopenModes[0][c], true
		} else {
			return plainFopenModes[1][c], true
		}
	default:
		return 0, false
	}
}

// @see: _php_stream_fopen
func plainFilesOpen(ctx *php.Context, w *StreamWrapper, filename string, mode string, options int, context *StreamContext) (stream *Stream, openedPath string) {
	openFlags, ok := plainParseFopenMode(mode)
	if !ok {
		//PhpStreamWrapperLogError(ctx, PhpPlainFilesWrapper, options, fmt.Sprintf("`%s' is not a valid mode for fopen", mode))
		return nil, ""
	}

	realpath := filename

	var persistentId string

	fp, err := os.OpenFile(realpath, openFlags, 0666)
	if err != nil {
		return nil, ""
	}

	streamOp := NewStreamStdioFromFile(fp)
	if streamOp == nil {
		return nil, ""
	}

	stream = NewStreamEx(ctx, streamOp, mode, persistentId != "")

	return stream, realpath
}
