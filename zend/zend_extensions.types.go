// <<generate>>

package zend

/**
 * ZendExtensionVersionInfo
 */
type ZendExtensionVersionInfo struct {
	zend_extension_api_no int
	build_id              *byte
}

func (this *ZendExtensionVersionInfo) GetZendExtensionApiNo() int { return this.zend_extension_api_no }
func (this *ZendExtensionVersionInfo) SetZendExtensionApiNo(value int) {
	this.zend_extension_api_no = value
}
func (this *ZendExtensionVersionInfo) GetBuildId() *byte      { return this.build_id }
func (this *ZendExtensionVersionInfo) SetBuildId(value *byte) { this.build_id = value }

/**
 * ZendExtension
 */
type ZendExtension struct {
	name                  *byte
	version               *byte
	author                *byte
	URL                   *byte
	copyright             *byte
	startup               StartupFuncT
	shutdown              ShutdownFuncT
	activate              ActivateFuncT
	deactivate            DeactivateFuncT
	message_handler       MessageHandlerFuncT
	op_array_handler      OpArrayHandlerFuncT
	statement_handler     StatementHandlerFuncT
	fcall_begin_handler   FcallBeginHandlerFuncT
	fcall_end_handler     FcallEndHandlerFuncT
	op_array_ctor         OpArrayCtorFuncT
	op_array_dtor         OpArrayDtorFuncT
	api_no_check          func(api_no int) int
	build_id_check        func(build_id *byte) int
	op_array_persist_calc OpArrayPersistCalcFuncT
	op_array_persist      OpArrayPersistFuncT
	reserved5             any
	reserved6             any
	reserved7             any
	reserved8             any
	handle                any
	resource_number       int
}

func (this *ZendExtension) GetName() *byte                              { return this.name }
func (this *ZendExtension) SetName(value *byte)                         { this.name = value }
func (this *ZendExtension) GetVersion() *byte                           { return this.version }
func (this *ZendExtension) SetVersion(value *byte)                      { this.version = value }
func (this *ZendExtension) GetAuthor() *byte                            { return this.author }
func (this *ZendExtension) SetAuthor(value *byte)                       { this.author = value }
func (this *ZendExtension) GetURL() *byte                               { return this.URL }
func (this *ZendExtension) SetURL(value *byte)                          { this.URL = value }
func (this *ZendExtension) GetCopyright() *byte                         { return this.copyright }
func (this *ZendExtension) SetCopyright(value *byte)                    { this.copyright = value }
func (this *ZendExtension) GetStartup() StartupFuncT                    { return this.startup }
func (this *ZendExtension) SetStartup(value StartupFuncT)               { this.startup = value }
func (this *ZendExtension) GetShutdown() ShutdownFuncT                  { return this.shutdown }
func (this *ZendExtension) SetShutdown(value ShutdownFuncT)             { this.shutdown = value }
func (this *ZendExtension) GetActivate() ActivateFuncT                  { return this.activate }
func (this *ZendExtension) SetActivate(value ActivateFuncT)             { this.activate = value }
func (this *ZendExtension) GetDeactivate() DeactivateFuncT              { return this.deactivate }
func (this *ZendExtension) SetDeactivate(value DeactivateFuncT)         { this.deactivate = value }
func (this *ZendExtension) GetMessageHandler() MessageHandlerFuncT      { return this.message_handler }
func (this *ZendExtension) SetMessageHandler(value MessageHandlerFuncT) { this.message_handler = value }
func (this *ZendExtension) GetOpArrayHandler() OpArrayHandlerFuncT      { return this.op_array_handler }
func (this *ZendExtension) SetOpArrayHandler(value OpArrayHandlerFuncT) {
	this.op_array_handler = value
}
func (this *ZendExtension) GetStatementHandler() StatementHandlerFuncT { return this.statement_handler }
func (this *ZendExtension) SetStatementHandler(value StatementHandlerFuncT) {
	this.statement_handler = value
}
func (this *ZendExtension) GetFcallBeginHandler() FcallBeginHandlerFuncT {
	return this.fcall_begin_handler
}
func (this *ZendExtension) SetFcallBeginHandler(value FcallBeginHandlerFuncT) {
	this.fcall_begin_handler = value
}
func (this *ZendExtension) GetFcallEndHandler() FcallEndHandlerFuncT { return this.fcall_end_handler }
func (this *ZendExtension) SetFcallEndHandler(value FcallEndHandlerFuncT) {
	this.fcall_end_handler = value
}
func (this *ZendExtension) GetOpArrayCtor() OpArrayCtorFuncT          { return this.op_array_ctor }
func (this *ZendExtension) SetOpArrayCtor(value OpArrayCtorFuncT)     { this.op_array_ctor = value }
func (this *ZendExtension) GetOpArrayDtor() OpArrayDtorFuncT          { return this.op_array_dtor }
func (this *ZendExtension) SetOpArrayDtor(value OpArrayDtorFuncT)     { this.op_array_dtor = value }
func (this *ZendExtension) GetApiNoCheck() func(api_no int) int       { return this.api_no_check }
func (this *ZendExtension) SetApiNoCheck(value func(api_no int) int)  { this.api_no_check = value }
func (this *ZendExtension) GetBuildIdCheck() func(build_id *byte) int { return this.build_id_check }
func (this *ZendExtension) SetBuildIdCheck(value func(build_id *byte) int) {
	this.build_id_check = value
}
func (this *ZendExtension) GetOpArrayPersistCalc() OpArrayPersistCalcFuncT {
	return this.op_array_persist_calc
}
func (this *ZendExtension) SetOpArrayPersistCalc(value OpArrayPersistCalcFuncT) {
	this.op_array_persist_calc = value
}
func (this *ZendExtension) GetOpArrayPersist() OpArrayPersistFuncT { return this.op_array_persist }
func (this *ZendExtension) SetOpArrayPersist(value OpArrayPersistFuncT) {
	this.op_array_persist = value
}
func (this *ZendExtension) GetReserved5() any           { return this.reserved5 }
func (this *ZendExtension) SetReserved5(value any)      { this.reserved5 = value }
func (this *ZendExtension) GetReserved6() any           { return this.reserved6 }
func (this *ZendExtension) SetReserved6(value any)      { this.reserved6 = value }
func (this *ZendExtension) GetReserved7() any           { return this.reserved7 }
func (this *ZendExtension) SetReserved7(value any)      { this.reserved7 = value }
func (this *ZendExtension) GetReserved8() any           { return this.reserved8 }
func (this *ZendExtension) SetReserved8(value any)      { this.reserved8 = value }
func (this *ZendExtension) GetHandle() any              { return this.handle }
func (this *ZendExtension) SetHandle(value any)         { this.handle = value }
func (this *ZendExtension) GetResourceNumber() int      { return this.resource_number }
func (this *ZendExtension) SetResourceNumber(value int) { this.resource_number = value }

/**
 * ZendExtensionPersistData
 */
type ZendExtensionPersistData struct {
	op_array *ZendOpArray
	size     int
	mem      *byte
}

func (this *ZendExtensionPersistData) GetOpArray() *ZendOpArray      { return this.op_array }
func (this *ZendExtensionPersistData) SetOpArray(value *ZendOpArray) { this.op_array = value }
func (this *ZendExtensionPersistData) GetSize() int                  { return this.size }
func (this *ZendExtensionPersistData) SetSize(value int)             { this.size = value }
func (this *ZendExtensionPersistData) GetMem() *byte                 { return this.mem }
func (this *ZendExtensionPersistData) SetMem(value *byte)            { this.mem = value }
