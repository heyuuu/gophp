package php

import (
	"fmt"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/types"
	"net/http"
)

// Engine
type Engine struct {
	modules *types.Table[*Module]
	host    string
	port    int
}

func NewEngine() *Engine {
	return &Engine{}
}

func (engine *Engine) Start() error {
	// todo
	return nil
}

/* lifecycle */
func (engine *Engine) NewContext(request *http.Request, response http.ResponseWriter) *Context {
	return NewContext(engine, request, response)
}

func (engine *Engine) HandleContext(ctx *Context, handler func(ctx *Context)) {
	ctx.Start()
	defer ctx.Finish()
	handler(ctx)
}

func (engine *Engine) HttpServe(host string, port int, handler func(ctx *Context)) error {
	engine.host = host
	engine.port = port
	addr := fmt.Sprintf("%s:%d", host, port)
	err := http.ListenAndServe(addr, http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := NewContext(engine, req, res)
		engine.HandleContext(ctx, handler)
	}))
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

/* modules */

func (engine *Engine) RegisterModule(m *Module) *Module {
	lcName := ascii.StrToLower(m.Name())
	// 若已注册，返回nil
	if engine.modules.Exists(lcName) {
		return nil
	}

	// 复制值，返回新地址
	tmp := *m
	tmp.moduleNumber = engine.modules.Len()
	engine.modules.Add(lcName, &tmp)
	return &tmp
}
