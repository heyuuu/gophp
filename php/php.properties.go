package php

import "github.com/heyuuu/gophp/php/types"

// properties for Context
func (c *Context) Executor() *Executor {
	return c.executor
}

// properties for StdInternObject
func (o *StdInternObject) Ctx() *Context {
	return o.ctx
}
func (o *StdInternObject) Obj() *types.Object {
	return o.obj
}
func (o *StdInternObject) Class() *types.Class {
	return o.class
}
