// <<generate>>

package zend

import b "sik/builtin"

/**
 * ZendModuleEntry
 */
type ZendModuleEntry struct {
	size                  uint16
	zend_api              uint
	zend_debug            uint8
	zts                   uint8
	ini_entry             *ZendIniEntry
	deps                  *ZendModuleDep
	name                  *byte
	functions             *ZendFunctionEntry
	module_startup_func   func(type_ int, module_number int) int
	module_shutdown_func  func(type_ int, module_number int) int
	request_startup_func  func(type_ int, module_number int) int
	request_shutdown_func func(type_ int, module_number int) int
	info_func             func(zend_module *ZendModuleEntry)
	version               *byte
	globals_size          int
	globals_ptr           any
	globals_ctor          func(global any)
	globals_dtor          func(global any)
	post_deactivate_func  func() int
	module_started        int
	type_                 uint8
	handle                any
	module_number         int
	build_id              *byte
}

func MakeZendModuleEntry(
	size uint16,
	zend_api uint,
	zend_debug uint8,
	zts uint8,
	ini_entry *ZendIniEntry,
	deps *ZendModuleDep,
	name *byte,
	functions *ZendFunctionEntry,
	module_startup_func func(type_ int, module_number int) int,
	module_shutdown_func func(type_ int, module_number int) int,
	request_startup_func func(type_ int, module_number int) int,
	request_shutdown_func func(type_ int, module_number int) int,
	info_func func(zend_module *ZendModuleEntry),
	version *byte,
	globals_size int,
	globals_ptr any,
	globals_ctor func(global any),
	globals_dtor func(global any),
	post_deactivate_func func() int,
	module_started int,
	type_ uint8,
	handle any,
	module_number int,
	build_id *byte,
) ZendModuleEntry {
	return ZendModuleEntry{
		zend_api:              zend_api,
		zend_debug:            zend_debug,
		zts:                   zts,
		ini_entry:             ini_entry,
		deps:                  deps,
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
		post_deactivate_func:  post_deactivate_func,
		module_started:        module_started,
		type_:                 type_,
		handle:                handle,
		module_number:         module_number,
		build_id:              build_id,
	}
}

func (this *ZendModuleEntry) GetZendApi() uint                 { return this.zend_api }
func (this *ZendModuleEntry) GetDeps() *ZendModuleDep          { return this.deps }
func (this *ZendModuleEntry) GetName() *byte                   { return this.name }
func (this *ZendModuleEntry) GetNameStr() string               { return b.CastStrAuto(this.name) }
func (this *ZendModuleEntry) GetFunctions() *ZendFunctionEntry { return this.functions }
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
func (this *ZendModuleEntry) SetInfoFunc(value func(zend_module *ZendModuleEntry)) {
	this.info_func = value
}
func (this *ZendModuleEntry) GetVersion() *byte                 { return this.version }
func (this *ZendModuleEntry) SetVersion(value *byte)            { this.version = value }
func (this *ZendModuleEntry) GetGlobalsSize() int               { return this.globals_size }
func (this *ZendModuleEntry) GetGlobalsPtr() any                { return this.globals_ptr }
func (this *ZendModuleEntry) GetGlobalsCtor() func(global any)  { return this.globals_ctor }
func (this *ZendModuleEntry) GetGlobalsDtor() func(global any)  { return this.globals_dtor }
func (this *ZendModuleEntry) GetPostDeactivateFunc() func() int { return this.post_deactivate_func }
func (this *ZendModuleEntry) GetModuleStarted() int             { return this.module_started }
func (this *ZendModuleEntry) SetModuleStarted(value int)        { this.module_started = value }
func (this *ZendModuleEntry) GetType() uint8                    { return this.type_ }
func (this *ZendModuleEntry) SetType(value uint8)               { this.type_ = value }
func (this *ZendModuleEntry) GetHandle() any                    { return this.handle }
func (this *ZendModuleEntry) SetHandle(value any)               { this.handle = value }
func (this *ZendModuleEntry) GetModuleNumber() int              { return this.module_number }
func (this *ZendModuleEntry) SetModuleNumber(value int)         { this.module_number = value }
func (this *ZendModuleEntry) GetBuildId() *byte                 { return this.build_id }

/**
 * ZendModuleDep
 */
type ZendModuleDep struct {
	name    *byte
	rel     *byte
	version *byte
	type_   uint8
}

func MakeZendModuleDep(name *byte, rel *byte, version *byte, type_ uint8) ZendModuleDep {
	return ZendModuleDep{
		name:    name,
		rel:     rel,
		version: version,
		type_:   type_,
	}
}
func (this *ZendModuleDep) GetName() *byte { return this.name }
func (this *ZendModuleDep) GetType() uint8 { return this.type_ }
