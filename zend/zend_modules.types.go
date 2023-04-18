package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * ModuleEntry
 */
type ModuleEntry struct {
	name                  string
	functions             []types.FunctionEntry
	module_startup_func   func(type_ int, module_number int) int
	module_shutdown_func  func(type_ int, module_number int) int
	request_startup_func  func(type_ int, module_number int) int
	request_shutdown_func func(type_ int, module_number int) int
	info_func             func(zend_module *ModuleEntry)
	version               string
	globals_size          int
	globals_ptr           any
	globals_ctor          func(global any)
	globals_dtor          func(global any)
	post_deactivate_func  func() int
	module_started        int
	handle                any
	module_number         int
}

func MakeZendModuleEntry(
	name string,
	functions []types.FunctionEntry,
	module_startup_func func(type_ int, module_number int) int,
	module_shutdown_func func(type_ int, module_number int) int,
	request_startup_func func(type_ int, module_number int) int,
	request_shutdown_func func(type_ int, module_number int) int,
	info_func func(zend_module *ModuleEntry),
	version string,

	globals_size int,
	globals_ptr any,
	globals_ctor func(global any),
	globals_dtor func(global any),

) ModuleEntry {
	return ModuleEntry{
		name:                  name,
		functions:             functions,
		module_startup_func:   module_startup_func,
		module_shutdown_func:  module_shutdown_func,
		request_startup_func:  request_startup_func,
		request_shutdown_func: request_shutdown_func,
		info_func:             info_func,
		version:               version,
		globals_size:          globals_size,
		globals_ptr:           globals_ptr,
		globals_ctor:          globals_ctor,
		globals_dtor:          globals_dtor,
	}
}

func (this *ModuleEntry) SetInfoFunc(value func(zend_module *ModuleEntry)) {
	this.info_func = value
}
func (this *ModuleEntry) SetHandle(value any)        { this.handle = value }
func (this *ModuleEntry) SetModuleNumber(value int)  { this.module_number = value }
func (this *ModuleEntry) SetModuleStarted(value int) { this.module_started = value }

func (this *ModuleEntry) GetName() string                     { return this.name }
func (this *ModuleEntry) GetFunctions() []types.FunctionEntry { return this.functions }
func (this *ModuleEntry) GetModuleStartupFunc() func(type_ int, module_number int) int {
	return this.module_startup_func
}
func (this *ModuleEntry) GetModuleShutdownFunc() func(type_ int, module_number int) int {
	return this.module_shutdown_func
}
func (this *ModuleEntry) GetRequestStartupFunc() func(type_ int, module_number int) int {
	return this.request_startup_func
}
func (this *ModuleEntry) GetRequestShutdownFunc() func(type_ int, module_number int) int {
	return this.request_shutdown_func
}
func (this *ModuleEntry) GetInfoFunc() func(zend_module *ModuleEntry) { return this.info_func }
func (this *ModuleEntry) GetVersion() string                          { return this.version }
func (this *ModuleEntry) GetGlobalsSize() int                         { return this.globals_size }
func (this *ModuleEntry) GetGlobalsPtr() any                          { return this.globals_ptr }
func (this *ModuleEntry) GetGlobalsCtor() func(global any)            { return this.globals_ctor }
func (this *ModuleEntry) GetGlobalsDtor() func(global any)            { return this.globals_dtor }
func (this *ModuleEntry) GetModuleStarted() int                       { return this.module_started }
func (this *ModuleEntry) GetType() uint8                              { return MODULE_PERSISTENT }
func (this *ModuleEntry) GetHandle() any                              { return this.handle }
func (this *ModuleEntry) GetModuleNumber() int                        { return this.module_number }
