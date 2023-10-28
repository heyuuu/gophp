package php

// Context
type Context struct {
	engine *Engine
	cg     CompilerGlobals
	eg     ExecutorGlobals
}

func (c *Context) Engine() *Engine      { return c.engine }
func (c *Context) CG() *CompilerGlobals { return &c.cg }
func (c *Context) EG() *ExecutorGlobals { return &c.eg }
