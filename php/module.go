package php

import "github.com/heyuuu/gophp/php/types"

// ModuleEntry
type ModuleEntry struct {
	Name            string
	Version         string
	Functions       []types.FunctionDecl
	ModuleStartup   func(ctx *Context, moduleNumber int) bool
	ModuleShutdown  func(ctx *Context, moduleNumber int) bool
	RequestStartup  func(ctx *Context, moduleNumber int) bool
	RequestShutdown func(ctx *Context, moduleNumber int) bool
}

// Module
type Module struct {
	moduleNumber  int
	moduleStarted bool
	data          ModuleEntry
}

func NewModule(moduleNumber int, data ModuleEntry) *Module {
	return &Module{moduleNumber: moduleNumber, data: data}
}

func (m *Module) IsModuleStarted() bool { return m.moduleStarted }
func (m *Module) ModuleNumber() int     { return m.moduleNumber }
func (m *Module) Data() ModuleEntry     { return m.data }

func (m *Module) Name() string    { return m.data.Name }
func (m *Module) Version() string { return m.data.Version }
func (m *Module) ModuleStartup(ctx *Context) bool {
	m.moduleStarted = true
	if m.data.ModuleStartup == nil {
		return true
	}
	return m.data.ModuleStartup(ctx, m.moduleNumber)
}
func (m *Module) ModuleShutdown(ctx *Context) bool {
	if m.moduleStarted {
		if m.data.ModuleShutdown != nil {
			m.data.ModuleShutdown(ctx, m.moduleNumber)
		}
	}
	m.moduleStarted = false
	return true
}
func (m *Module) RequestStartup(ctx *Context) bool {
	if m.data.RequestStartup == nil {
		return true
	}
	return m.data.RequestStartup(ctx, m.moduleNumber)
}
func (m *Module) RequestShutdown(ctx *Context) bool {
	if m.data.RequestShutdown == nil {
		return true
	}
	return m.data.RequestShutdown(ctx, m.moduleNumber)
}
