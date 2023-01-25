// <<generate>>

package zend

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
	info_func             func(ZEND_MODULE_INFO_FUNC_ARGS)
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

func (this ZendModuleEntry) GetSize() uint16                        { return this.size }
func (this *ZendModuleEntry) SetSize(value uint16)                  { this.size = value }
func (this ZendModuleEntry) GetZendApi() uint                       { return this.zend_api }
func (this *ZendModuleEntry) SetZendApi(value uint)                 { this.zend_api = value }
func (this ZendModuleEntry) GetZendDebug() uint8                    { return this.zend_debug }
func (this *ZendModuleEntry) SetZendDebug(value uint8)              { this.zend_debug = value }
func (this ZendModuleEntry) GetZts() uint8                          { return this.zts }
func (this *ZendModuleEntry) SetZts(value uint8)                    { this.zts = value }
func (this ZendModuleEntry) GetIniEntry() *ZendIniEntry             { return this.ini_entry }
func (this *ZendModuleEntry) SetIniEntry(value *ZendIniEntry)       { this.ini_entry = value }
func (this ZendModuleEntry) GetDeps() *ZendModuleDep                { return this.deps }
func (this *ZendModuleEntry) SetDeps(value *ZendModuleDep)          { this.deps = value }
func (this ZendModuleEntry) GetName() *byte                         { return this.name }
func (this *ZendModuleEntry) SetName(value *byte)                   { this.name = value }
func (this ZendModuleEntry) GetFunctions() *ZendFunctionEntry       { return this.functions }
func (this *ZendModuleEntry) SetFunctions(value *ZendFunctionEntry) { this.functions = value }
func (this ZendModuleEntry) GetModuleStartupFunc() func(type_ int, module_number int) int {
	return this.module_startup_func
}
func (this *ZendModuleEntry) SetModuleStartupFunc(value func(type_ int, module_number int) int) {
	this.module_startup_func = value
}
func (this ZendModuleEntry) GetModuleShutdownFunc() func(type_ int, module_number int) int {
	return this.module_shutdown_func
}
func (this *ZendModuleEntry) SetModuleShutdownFunc(value func(type_ int, module_number int) int) {
	this.module_shutdown_func = value
}
func (this ZendModuleEntry) GetRequestStartupFunc() func(type_ int, module_number int) int {
	return this.request_startup_func
}
func (this *ZendModuleEntry) SetRequestStartupFunc(value func(type_ int, module_number int) int) {
	this.request_startup_func = value
}
func (this ZendModuleEntry) GetRequestShutdownFunc() func(type_ int, module_number int) int {
	return this.request_shutdown_func
}
func (this *ZendModuleEntry) SetRequestShutdownFunc(value func(type_ int, module_number int) int) {
	this.request_shutdown_func = value
}
func (this ZendModuleEntry) GetInfoFunc() func(ZEND_MODULE_INFO_FUNC_ARGS) { return this.info_func }
func (this *ZendModuleEntry) SetInfoFunc(value func(ZEND_MODULE_INFO_FUNC_ARGS)) {
	this.info_func = value
}
func (this ZendModuleEntry) GetVersion() *byte                      { return this.version }
func (this *ZendModuleEntry) SetVersion(value *byte)                { this.version = value }
func (this ZendModuleEntry) GetGlobalsSize() int                    { return this.globals_size }
func (this *ZendModuleEntry) SetGlobalsSize(value int)              { this.globals_size = value }
func (this ZendModuleEntry) GetGlobalsPtr() any                     { return this.globals_ptr }
func (this *ZendModuleEntry) SetGlobalsPtr(value any)               { this.globals_ptr = value }
func (this ZendModuleEntry) GetGlobalsCtor() func(global any)       { return this.globals_ctor }
func (this *ZendModuleEntry) SetGlobalsCtor(value func(global any)) { this.globals_ctor = value }
func (this ZendModuleEntry) GetGlobalsDtor() func(global any)       { return this.globals_dtor }
func (this *ZendModuleEntry) SetGlobalsDtor(value func(global any)) { this.globals_dtor = value }
func (this ZendModuleEntry) GetPostDeactivateFunc() func() int      { return this.post_deactivate_func }
func (this *ZendModuleEntry) SetPostDeactivateFunc(value func() int) {
	this.post_deactivate_func = value
}
func (this ZendModuleEntry) GetModuleStarted() int       { return this.module_started }
func (this *ZendModuleEntry) SetModuleStarted(value int) { this.module_started = value }
func (this ZendModuleEntry) GetType() uint8              { return this.type_ }
func (this *ZendModuleEntry) SetType(value uint8)        { this.type_ = value }
func (this ZendModuleEntry) GetHandle() any              { return this.handle }
func (this *ZendModuleEntry) SetHandle(value any)        { this.handle = value }
func (this ZendModuleEntry) GetModuleNumber() int        { return this.module_number }
func (this *ZendModuleEntry) SetModuleNumber(value int)  { this.module_number = value }
func (this ZendModuleEntry) GetBuildId() *byte           { return this.build_id }
func (this *ZendModuleEntry) SetBuildId(value *byte)     { this.build_id = value }

/**
 * ZendModuleDep
 */
type ZendModuleDep struct {
	name    *byte
	rel     *byte
	version *byte
	type_   uint8
}

func (this ZendModuleDep) GetName() *byte          { return this.name }
func (this *ZendModuleDep) SetName(value *byte)    { this.name = value }
func (this ZendModuleDep) GetRel() *byte           { return this.rel }
func (this *ZendModuleDep) SetRel(value *byte)     { this.rel = value }
func (this ZendModuleDep) GetVersion() *byte       { return this.version }
func (this *ZendModuleDep) SetVersion(value *byte) { this.version = value }
func (this ZendModuleDep) GetType() uint8          { return this.type_ }
func (this *ZendModuleDep) SetType(value uint8)    { this.type_ = value }
