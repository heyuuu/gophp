package php

import (
	"io"
	"net/http"
)

// Context
type Context struct {
	engine *Engine
	cg     CompilerGlobals
	eg     ExecutorGlobals
	og     OutputGlobals
	pg     PhpCoreGlobals

	values map[string]any

	eh ErrorHandling
}

func MockContext() *Context {
	return &Context{}
}

func initContext(e *Engine, baseCtx *Context, request *http.Request, response http.ResponseWriter) *Context {
	ctx := &Context{engine: e}
	if baseCtx != nil {
		ctx.cg.Init()
		ctx.eg.Init(ctx, &baseCtx.eg)
		ctx.og.Init()
		ctx.pg.Init()
	} else {
		ctx.cg.Init()
		ctx.eg.Init(ctx, nil)
		ctx.og.Init()
		ctx.pg.Init()
	}

	return ctx
}

/* lifecycle */
func (c *Context) Start()  {}
func (c *Context) Finish() {}

func (c *Context) Engine() *Engine { return c.engine }

func (c *Context) CG() *CompilerGlobals { return &c.cg }
func (c *Context) EG() *ExecutorGlobals { return &c.eg }
func (c *Context) OG() *OutputGlobals   { return &c.og }
func (c *Context) PG() *PhpCoreGlobals  { return &c.pg }

// fast functions
func (c *Context) CurrEX() *ExecuteData { return c.eg.CurrentExecuteData() }

// output
func (c *Context) Write(data []byte)        { c.og.Write(data) }
func (c *Context) WriteString(str string)   { c.og.WriteString(str) }
func (c *Context) WriteStringUb(str string) { c.og.WriteStringUnbuffered(str) }

func (c *Context) AsWriter() io.Writer {
	return ctxWriter{ctx: c}
}

// values
func (c *Context) GetValue(key string) any        { return c.values[key] }
func (c *Context) SetValue(key string, value any) { c.values[key] = value }

func ContextGetOrInit[T any](ctx *Context, key string, initializer func() T) (T, bool) {
	if v, ok := ctx.values[key]; ok {
		if result, typeMatch := v.(T); typeMatch {
			return result, true
		} else {
			return result, false
		}
	} else {
		result := initializer()
		if ctx.values == nil {
			ctx.values = map[string]any{}
		}
		ctx.values[key] = result
		return result, true
	}
}
