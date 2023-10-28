package php

// Engine
type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) NewContext() *Context {
	return &Context{engine: e}
}

func (e *Engine) Start() error {
	// todo
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
