package zend

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"sort"
)

var currEngine *Engine

func G() *Engine { return currEngine }

type Engine struct {
	modules map[string]*ModuleEntry
}

/* module register */
func (g *Engine) InitModules() {
	g.modules = make(map[string]*ModuleEntry, 32)
}

func (g *Engine) RegisterModule(m *ModuleEntry) *ModuleEntry {
	lcName := ascii.StrToLower(m.Name())
	// 若已注册，返回nil
	if _, ok := g.modules[lcName]; ok {
		return nil
	}

	// 复制值，返回新地址
	tmp := *m
	g.modules[lcName] = &tmp
	return &tmp
}

func (g *Engine) GetModule(name string) *ModuleEntry {
	lcName := ascii.StrToLower(name)
	if m, ok := g.modules[lcName]; ok {
		return m
	}
	return nil
}

func (g *Engine) DelModule(name string) {
	lcName := ascii.StrToLower(name)
	delete(g.modules, lcName)
}

func (g *Engine) CountModules() int { return len(g.modules) }

func (g *Engine) GetSortedModules() []*ModuleEntry {
	var modules []*ModuleEntry
	for _, module := range g.modules {
		modules = append(modules, module)
	}
	sort.Slice(modules, func(i, j int) bool {
		return ascii.StrCaseCompare(modules[i].Name(), modules[j].Name()) < 0
	})
	return modules
}

func (g *Engine) EachModule(handler func(module *ModuleEntry)) {
	for _, module := range g.modules {
		handler(module)
	}
}

func (g *Engine) EachModuleReserve(handler func(module *ModuleEntry)) {
	// todo 确认 reserve 方式是否符合原意
	modules := g.GetSortedModules()
	for i := len(modules) - 1; i >= 0; i-- {
		module := modules[i]
		handler(module)
	}
}
func (g *Engine) DestroyModules() {
	g.EachModuleReserve(func(module *ModuleEntry) {
		ModuleDestructor(module)
	})
	g.InitModules()
}
