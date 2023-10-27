package php

// Engine
type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
}

//func (e *Engine) Start() error {
//}

//func (e *Engine) HandleRequest() error {
//	err := e.RequestStartup()
//	if err != nil {
//		return err
//	}
//}
//
//func (e *Engine) RequestStartup() error {
//}

func (e *Engine) ExecuteScript(primaryFile *FileHandle) error {

}
