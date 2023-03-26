package zend

import (
	"sik/zend/types"
)

/**
 * ZendModuleEntry
 */
type ZendModuleEntry struct {
	name                  string
	functions             []types.ZendFunctionEntry
	module_startup_func   func(type_ int, module_number int) int
	module_shutdown_func  func(type_ int, module_number int) int
	request_startup_func  func(type_ int, module_number int) int
	request_shutdown_func func(type_ int, module_number int) int
	info_func             func(zend_module *ZendModuleEntry)
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
	functions []types.ZendFunctionEntry,
	module_startup_func func(type_ int, module_number int) int,
	module_shutdown_func func(type_ int, module_number int) int,
	request_startup_func func(type_ int, module_number int) int,
	request_shutdown_func func(type_ int, module_number int) int,
	info_func func(zend_module *ZendModuleEntry),
	version string,

	globals_size int,
	globals_ptr any,
	globals_ctor func(global any),
	globals_dtor func(global any),

) ZendModuleEntry {
	return ZendModuleEntry{
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

func (this *ZendModuleEntry) SetInfoFunc(value func(zend_module *ZendModuleEntry)) {
	this.info_func = value
}
func (this *ZendModuleEntry) SetHandle(value any)        { this.handle = value }
func (this *ZendModuleEntry) SetModuleNumber(value int)  { this.module_number = value }
func (this *ZendModuleEntry) SetModuleStarted(value int) { this.module_started = value }

func (this *ZendModuleEntry) GetName() string                         { return this.name }
func (this *ZendModuleEntry) GetFunctions() []types.ZendFunctionEntry { return this.functions }
func (this *ZendModuleEntry) GetModuleStartupFunc() func(type_ int, module_number int) int {
	return this.module_startup_func
}
func (this *ZendModuleEntry) GetModuleShutdownFunc() func(type_ int, module_number int) int {
	return this.module_shutdown_func
}
func (this *ZendModuleEntry) GetRequestStartupFunc() func(type_ int, module_number int) int {
	return this.request_startup_func
}
func (this *ZendModuleEntry) GetRequestShutdownFunc() func(type_ int, module_number int) int {
	return this.request_shutdown_func
}
func (this *ZendModuleEntry) GetInfoFunc() func(zend_module *ZendModuleEntry) { return this.info_func }
func (this *ZendModuleEntry) GetVersion() string                              { return this.version }
func (this *ZendModuleEntry) GetGlobalsSize() int                             { return this.globals_size }
func (this *ZendModuleEntry) GetGlobalsPtr() any                              { return this.globals_ptr }
func (this *ZendModuleEntry) GetGlobalsCtor() func(global any)                { return this.globals_ctor }
func (this *ZendModuleEntry) GetGlobalsDtor() func(global any)                { return this.globals_dtor }
func (this *ZendModuleEntry) GetModuleStarted() int                           { return this.module_started }
func (this *ZendModuleEntry) GetType() uint8                                  { return MODULE_PERSISTENT }
func (this *ZendModuleEntry) GetHandle() any                                  { return this.handle }
func (this *ZendModuleEntry) GetModuleNumber() int                            { return this.module_number }
