package php

import (
	"github.com/heyuuu/gophp/php/operators"
	"net/http"
)

// Context
type Context struct {
	engine *Engine
	cg     CompilerGlobals
	eg     ExecutorGlobals
	og     OutputGlobals

	operator *operators.Operator
}

func NewContext(e *Engine, request *http.Request, response http.ResponseWriter) *Context {
	ctx := &Context{engine: e}
	ctx.cg.Init()
	ctx.eg.Init()
	ctx.og.Init()
	return ctx
}

/* lifecycle */
func (c *Context) Start()  {}
func (c *Context) Finish() {}

func (c *Context) Engine() *Engine { return c.engine }

func (c *Context) CG() *CompilerGlobals { return &c.cg }
func (c *Context) EG() *ExecutorGlobals { return &c.eg }
func (c *Context) OG() *OutputGlobals   { return &c.og }

func (c *Context) Operator() *operators.Operator { return c.operator }

// output
func (c *Context) Write(data []byte)        { c.og.Write(data) }
func (c *Context) WriteString(str string)   { c.og.WriteString(str) }
func (c *Context) WriteStringUb(str string) { c.og.WriteStringUnbuffered(str) }
