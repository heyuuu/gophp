package php

import (
	"fmt"
	"net/http"
)

// Engine
type Engine struct {
	host string
	port int
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

//func (e *Engine) HandleRequest() error {
//	err := e.RequestStartup()
//	if err != nil {
//		return err
//	}
//}
//
//func (e *Engine) RequestStartup() error {
//}
