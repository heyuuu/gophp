package streams

import "github.com/heyuuu/gophp/php/types"

// properties for Stream
func (t *Stream) Op() StreamOp {
	return t.op
}
func (t *Stream) IsPersistent() bool {
	return t.isPersistent
}
func (t *Stream) Mode() string {
	return t.mode
}
func (t *Stream) Resource() *types.Resource {
	return t.resource
}
func (t *Stream) ContextResource() *types.Resource {
	return t.contextResource
}

// properties for StreamContext
func (c *StreamContext) Notifier() *StreamNotifier {
	return c.notifier
}
func (c *StreamContext) SetNotifier(v *StreamNotifier) {
	c.notifier = v
}
func (c *StreamContext) Resource() *types.Resource {
	return c.resource
}

// properties for StreamStdio
func (s *StreamStdio) Seekable() bool {
	return s.seekable
}
