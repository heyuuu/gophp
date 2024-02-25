package streams

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
)

type StreamBindable interface {
	Bind(*Stream)
}

type Stream struct {
	op              StreamOp        `get:""`
	isPersistent    bool            `get:""`
	mode            string          `get:""`
	resource        *types.Resource `get:""`
	contextResource *types.Resource `get:""`
}

var _ types.ResourceCloser = (*Stream)(nil)

func NewStream(ctx *php.Context, op StreamOp, mode string) *Stream {
	return NewStreamEx(ctx, op, mode, false)
}

func NewStreamEx(ctx *php.Context, op StreamOp, mode string, isPersistent bool) *Stream {
	s := &Stream{
		op:           op,
		isPersistent: isPersistent,
		mode:         mode,
	}

	// register
	if isPersistent {
		s.resource = php.RegisterResource(ctx, s, LePstream)
	} else {
		s.resource = php.RegisterResource(ctx, s, LeStream)
	}

	return s
}

func (t *Stream) ResourceClose() int {
	//TODO implement me
	panic("implement me")
}
