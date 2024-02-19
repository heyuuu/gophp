package php

import (
	"fmt"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"net/http"
)

// Engine
type Engine struct {
	modules *types.Table[*Module]
	host    string   `get:""`
	port    int      `get:""`
	baseCtx *Context `get:""`
}

func NewEngine() *Engine {
	CheckIsBoot()
	engine := &Engine{}
	engine.init()
	return engine
}

func (engine *Engine) init() {
	engine.modules = types.NewTable[*Module]()
	engine.baseCtx = initBaseContext(engine)

	moduleStartupRegisterConstants(engine.baseCtx)
	for _, entry := range builtinModuleEntries {
		engine.RegisterModule(entry)
	}
}

func (engine *Engine) Start() (err error) {
	PhpInitConfig(engine.BaseCtx())

	engine.baseCtx.INI().RegisterIniEntries(0, MainIniEntries)

	// todo
	err = engine.modules.EachEx(func(_ string, m *Module) error {
		if !m.ModuleStartup(engine.baseCtx) {
			return perr.Internal("module start failed: " + m.Name())
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

/* lifecycle */
func (engine *Engine) NewContext(request *http.Request, response http.ResponseWriter) *Context {
	return initContext(engine, engine.baseCtx, request, response)
}

func (engine *Engine) HandleContext(ctx *Context, handler func(ctx *Context)) {
	ctx.Start()
	defer func() {
		ctx.Finish()
	}()
	handler(ctx)
}

func (engine *Engine) HttpServe(host string, port int, handler func(ctx *Context)) error {
	engine.host = host
	engine.port = port
	addr := fmt.Sprintf("%s:%d", host, port)
	err := http.ListenAndServe(addr, http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := engine.NewContext(req, res)
		engine.HandleContext(ctx, handler)
	}))
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

/* modules */

func (engine *Engine) RegisterModule(entry ModuleEntry) *Module {
	lcName := ascii.StrToLower(entry.Name)
	// 若已注册，返回nil
	if engine.modules.Exists(lcName) {
		return nil
	}

	// 复制值，返回新地址
	moduleNumber := engine.modules.Len()
	module := NewModule(moduleNumber, entry)
	engine.modules.Add(lcName, module)

	// 注册模块函数
	RegisterModuleFunctions(engine.baseCtx, module, entry.Functions)

	return module
}
