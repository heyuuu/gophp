package php

import "net/http"

// Engine
type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) NewContext(request *http.Request, response http.ResponseWriter) *Context {
	return NewContext(e, request, response)
}

func (e *Engine) Start() error {
	// todo
	return nil
}

func (e *Engine) HandleContext(ctx *Context) {

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
