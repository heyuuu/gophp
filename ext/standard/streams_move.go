package standard

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/streams"
	"github.com/heyuuu/gophp/php/types"
)

func StreamContextFromResource(ctx *php.Context, zcontext *types.Resource, noContext bool) *streams.StreamContext {
	if zcontext != nil {
		return php.FetchResource[streams.StreamContext](ctx, zcontext, "Stream-Context", streams.LeStreamContext)
	} else if noContext {
		return nil
	} else if FG(ctx).DefaultContext() != nil {
		return FG(ctx).DefaultContext()
	} else {
		FG(ctx).SetDefaultContext(streams.NewStreamContext(ctx))
		return FG(ctx).DefaultContext()
	}
}

func StreamContextFromZval(ctx *php.Context, zcontext *types.Zval, noContext bool) *streams.StreamContext {
	if zcontext != nil {
		return php.FetchResourceEx[streams.StreamContext](ctx, zcontext, "Stream-Context", streams.LeStreamContext)
	} else if noContext {
		return nil
	} else if FG(ctx).DefaultContext() != nil {
		return FG(ctx).DefaultContext()
	} else {
		FG(ctx).SetDefaultContext(streams.NewStreamContext(ctx))
		return FG(ctx).DefaultContext()
	}
}
