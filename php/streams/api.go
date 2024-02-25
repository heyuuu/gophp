package streams

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
)

func StreamFromRes(ctx *php.Context, res *types.Resource) *Stream {
	return StreamFromResEx(ctx, res, "stream")
}
func StreamFromResEx(ctx *php.Context, res *types.Resource, resourceTypeName string) *Stream {
	return php.FetchResource2[Stream](ctx, res, resourceTypeName, LeStream, LePstream)
}

func StreamFromZval(ctx *php.Context, res *types.Zval) *Stream {
	return StreamFromZvalEx(ctx, res, "stream")
}
func StreamFromZvalEx(ctx *php.Context, res *types.Zval, resourceTypeName string) *Stream {
	return php.FetchResource2Ex[Stream](ctx, res, resourceTypeName, LeStream, LePstream)
}

func StreamToZval(stream *Stream) types.Zval {
	return types.ZvalResource(stream.Resource())
}

func LocalStreamFromResEx(ctx *php.Context, res *types.Resource, resourceTypeName string) *Stream {
	return php.FetchResource[Stream](ctx, res, resourceTypeName, LeStream)
}
func LocalStreamFromZvalEx(ctx *php.Context, res *types.Zval, resourceTypeName string) *Stream {
	return php.FetchResourceEx[Stream](ctx, res, resourceTypeName, LeStream)
}

func RegisterStreamFilterResource(ctx *php.Context, filter *StreamFilter) *types.Resource {
	return php.RegisterResource(ctx, filter, LeStreamFilter)
}

func StreamFilterFromRes(ctx *php.Context, res *types.Resource, resourceTypeName string) *StreamFilter {
	return php.FetchResource[StreamFilter](ctx, res, resourceTypeName, LeStreamFilter)
}
