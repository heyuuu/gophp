package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

// ModuleData
type ModuleData interface {
	Name() string
	Version() string
	Functions() []types.FunctionEntry
	ModuleStartup(moduleNumber int) bool
	ModuleShutdown(moduleNumber int) bool
	RequestStartup(moduleNumber int) bool
	RequestShutdown(moduleNumber int) bool
}

// ModuleEntry
type ModuleEntry struct {
	moduleNumber  int
	moduleStarted bool
	data          ModuleData
	infoFunc      func(zendModule *ModuleEntry)
}

func MakeZendModuleEntry(data ModuleData, infoFunc func(*ModuleEntry)) ModuleEntry {
	return ModuleEntry{data: data, infoFunc: infoFunc}
}

func (m *ModuleEntry) SetInfoFunc(value func(zendModule *ModuleEntry)) { m.infoFunc = value }
func (m *ModuleEntry) GetInfoFunc() func(zend_module *ModuleEntry)     { return m.infoFunc }
func (m *ModuleEntry) IsModuleStarted() bool                           { return m.moduleStarted }
func (m *ModuleEntry) ModuleNumber() int                               { return m.moduleNumber }

func (m *ModuleEntry) Name() string                     { return m.data.Name() }
func (m *ModuleEntry) Version() string                  { return m.data.Version() }
func (m *ModuleEntry) Functions() []types.FunctionEntry { return m.data.Functions() }
func (m *ModuleEntry) Init(moduleNumber int)            { m.moduleNumber = moduleNumber }
func (m *ModuleEntry) ModuleStartup() bool {
	m.moduleStarted = true
	return m.data.ModuleStartup(m.moduleNumber)
}
func (m *ModuleEntry) ModuleShutdown() bool {
	if m.moduleStarted {
		m.data.ModuleShutdown(m.moduleNumber)
	}
	m.moduleStarted = false
	return true
}
func (m *ModuleEntry) RequestStartup() bool {
	return m.data.RequestStartup(m.moduleNumber)
}
func (m *ModuleEntry) RequestShutdown() bool {
	return m.data.RequestShutdown(m.moduleNumber)
}
