package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * ModuleEntry
 */
type ModuleEntry struct {
	name                string
	functions           []types.FunctionEntry
	moduleStartupFunc   func(persistent bool, moduleNumber int) bool
	moduleShutdownFunc  func(persistent bool, moduleNumber int) bool
	requestStartupFunc  func(persistent bool, moduleNumber int) bool
	requestShutdownFunc func(persistent bool, moduleNumber int) bool
	infoFunc            func(zendModule *ModuleEntry)
	version             string
	globalsSize         int
	globalsPtr          any
	globalsCtor         func(global any)
	globalsDtor         func(global any)
	postDeactivateFunc  func() int
	moduleStarted       bool
	handle              any
	moduleNumber        int
}

func MakeZendModuleEntry(
	name string,
	functions []types.FunctionEntry,
	moduleStartupFunc func(persistent bool, moduleNumber int) bool,
	moduleShutdownFunc func(persistent bool, moduleNumber int) bool,
	requestStartupFunc func(persistent bool, moduleNumber int) bool,
	requestShutdownFunc func(persistent bool, moduleNumber int) bool,
	infoFunc func(zendModule *ModuleEntry),
	version string,
	globalsSize int,
	globalsPtr any,
	globalsCtor func(global any),
	globalsDtor func(global any),

) ModuleEntry {
	return ModuleEntry{
		name:                name,
		functions:           functions,
		moduleStartupFunc:   moduleStartupFunc,
		moduleShutdownFunc:  moduleShutdownFunc,
		requestStartupFunc:  requestStartupFunc,
		requestShutdownFunc: requestShutdownFunc,
		infoFunc:            infoFunc,
		version:             version,
		globalsSize:         globalsSize,
		globalsPtr:          globalsPtr,
		globalsCtor:         globalsCtor,
		globalsDtor:         globalsDtor,
	}
}

func (m *ModuleEntry) IsPersistent() bool { return true }

func (m *ModuleEntry) SetInfoFunc(value func(zendModule *ModuleEntry)) {
	m.infoFunc = value
}
func (m *ModuleEntry) SetHandle(value any)       { m.handle = value }
func (m *ModuleEntry) SetModuleNumber(value int) { m.moduleNumber = value }

func (m *ModuleEntry) Name() string                     { return m.name }
func (m *ModuleEntry) Functions() []types.FunctionEntry { return m.functions }

func (m *ModuleEntry) IsModuleStarted() bool   { return m.moduleStarted }
func (m *ModuleEntry) SetModuleStarted(b bool) { m.moduleStarted = b }

func (m *ModuleEntry) ModuleStartup() bool {
	if m.moduleStartupFunc == nil {
		return true
	}
	return m.moduleStartupFunc(m.IsPersistent(), m.moduleNumber)
}
func (m *ModuleEntry) ModuleShutdown() bool {
	if m.moduleShutdownFunc == nil {
		return true
	}
	return m.moduleShutdownFunc(m.IsPersistent(), m.moduleNumber)
}
func (m *ModuleEntry) RequestStartup() bool {
	if m.requestStartupFunc == nil {
		return true
	}
	return m.requestStartupFunc(m.IsPersistent(), m.moduleNumber)
}
func (m *ModuleEntry) RequestShutdown() bool {
	if m.requestShutdownFunc == nil {
		return true
	}
	return m.requestShutdownFunc(m.IsPersistent(), m.moduleNumber)
}

func (m *ModuleEntry) GetInfoFunc() func(zend_module *ModuleEntry) { return m.infoFunc }
func (m *ModuleEntry) GetVersion() string                          { return m.version }
func (m *ModuleEntry) GetGlobalsSize() int                         { return m.globalsSize }
func (m *ModuleEntry) GetGlobalsPtr() any                          { return m.globalsPtr }
func (m *ModuleEntry) GetGlobalsCtor() func(global any)            { return m.globalsCtor }
func (m *ModuleEntry) GetGlobalsDtor() func(global any)            { return m.globalsDtor }
func (m *ModuleEntry) GetModuleNumber() int                        { return m.moduleNumber }
