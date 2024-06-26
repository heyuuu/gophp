package php

import (
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/php/perr"
	"io"
	"net/http"
)

// Context
type Context struct {
	engine *Engine
	cg     CompilerGlobals
	eg     ExecutorGlobals
	og     OutputGlobals
	pg     PhpGlobals
	ini    IniGlobals

	values map[string]any

	eh ErrorHandling

	executor *Executor `get:""`
}

func MockContext() *Context {
	return &Context{}
}

func initBaseContext(e *Engine) *Context {
	assert.Assert(e != nil)
	ctx := &Context{engine: e, values: map[string]any{}}

	ctx.eg.InitBase(ctx)
	ctx.ini.Init(ctx, nil)

	return ctx
}

func initContext(e *Engine, baseCtx *Context, request *http.Request, response http.ResponseWriter) *Context {
	assert.Assert(e != nil && baseCtx != nil)

	ctx := &Context{engine: e, values: map[string]any{}}

	ctx.cg.Init()
	ctx.eg.Init(ctx, baseCtx.EG())
	ctx.og.Init()
	ctx.pg.Init(ctx, baseCtx.PG())
	ctx.ini.Init(ctx, baseCtx.INI())

	ctx.executor = NewExecutor(ctx)

	return ctx
}

/* lifecycle */
func (c *Context) Start()  {}
func (c *Context) Finish() {}

func (c *Context) Engine() *Engine { return c.engine }

func (c *Context) CG() *CompilerGlobals { return &c.cg }
func (c *Context) EG() *ExecutorGlobals { return &c.eg }
func (c *Context) OG() *OutputGlobals   { return &c.og }
func (c *Context) PG() *PhpGlobals      { return &c.pg }
func (c *Context) INI() *IniGlobals     { return &c.ini }

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

func ContextGetOrInit[T any](ctx *Context, key string, initializer func() T) T {
	if v, ok := ctx.values[key]; ok {
		if result, typeMatch := v.(T); typeMatch {
			return result
		} else {
			var tmp T
			panic(perr.Internalf("php.ContextGetOrInit() fail, expect type %T, got type %T, key=%s", tmp, v, key))
		}
	} else {
		result := initializer()
		if ctx.values == nil {
			ctx.values = map[string]any{}
		}
		ctx.values[key] = result
		return result
	}
}

func ContextGetOrInitPersistent[T any](ctx *Context, key string, initializer func() T) T {
	// todo 考虑持久化 globals 和非持久化 globals 之间的区分和生命周期问题
	return ContextGetOrInit(ctx, key, initializer)
}

func ContextDel(ctx *Context, key string) {
	delete(ctx.values, key)
}
