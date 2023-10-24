package php

// Context
type Context struct {
	cg CompilerGlobals
	eg ExecutorGlobals
}

func (c *Context) CG() *CompilerGlobals { return &c.cg }
func (c *Context) EG() *ExecutorGlobals { return &c.eg }
