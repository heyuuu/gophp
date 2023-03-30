package globals

import (
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/zend"
	"sort"
)

func G() *Globals { return currGlobals }

var currGlobals *Globals

type Globals struct {
	modules map[string]*zend.ModuleEntry
}

/* module register */
func (g *Globals) InitModules() {
	g.modules = make(map[string]*zend.ModuleEntry, 32)
}

func (g *Globals) RegisterModule(m *zend.ModuleEntry) *zend.ModuleEntry {
	lcName := ascii.StrToLower(m.GetName())
	// 若已注册，返回nil
	if _, ok := g.modules[lcName]; ok {
		return nil
	}

	// 复制值，返回新地址
	tmp := *m
	g.modules[lcName] = &tmp
	return &tmp
}

func (g *Globals) GetModule(name string) *zend.ModuleEntry {
	lcName := ascii.StrToLower(name)
	if m, ok := g.modules[lcName]; ok {
		return m
	}
	return nil
}

func (g *Globals) DelModule(name string) {
	lcName := ascii.StrToLower(name)
	delete(g.modules, lcName)
}

func (g *Globals) CountModules() int { return len(g.modules) }

func (g *Globals) GetSortedModules() []*zend.ModuleEntry {
	var modules []*zend.ModuleEntry
	for _, module := range g.modules {
		modules = append(modules, module)
	}
	sort.Slice(modules, func(i, j int) bool {
		return ascii.StrCaseCompare(modules[i].GetName(), modules[j].GetName()) < 0
	})
	return modules
}

func (g *Globals) EachModule(handler func(module *zend.ModuleEntry)) {
	for _, module := range g.modules {
		handler(module)
	}
}

func (g *Globals) EachModuleReserve(handler func(module *zend.ModuleEntry)) {
	// todo 确认 reserve 方式是否符合原意
	modules := g.GetSortedModules()
	for i := len(modules) - 1; i >= 0; i-- {
		module := modules[i]
		handler(module)
	}
}
func (g *Globals) DestroyModules() {
	g.EachModuleReserve(func(module *zend.ModuleEntry) {
		zend.ModuleDestructor(module)
	})
	g.InitModules()
}
