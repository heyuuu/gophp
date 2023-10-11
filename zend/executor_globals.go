package zend

import "github.com/heyuuu/gophp/php/types"

// ZendExecutorGlobals

type ZendExecutorGlobals struct {
	error_zval           types.Zval
	symtable_cache       []*types.Array
	symtable_cache_limit **types.Array
	symtable_cache_ptr   **types.Array
	symbol_table         *types.Array
	included_files       *types.Array
	error_reporting      int
	exit_status          int
	functionTable        FunctionTable
	classTable           ClassTable
	constantTable        ConstantTable
	vmStack              VmStack

	current_execute_data                *ZendExecuteData
	fake_scope                          *types.ClassEntry
	precision                           ZendLong
	persistent_constants_count          uint32
	persistent_functions_count          uint32
	persistent_classes_count            uint32
	in_autoload                         *types.Array
	autoload_func                       types.IFunction
	no_extensions                       bool
	vm_interrupt                        bool
	timed_out                           bool
	hard_timeout                        ZendLong
	regular_list                        types.Array
	persistent_list                     types.Array
	regularList                         ResourceTable
	persistentList                      ResourceTable
	user_error_handler_error_reporting  int
	user_error_handler                  types.Zval
	user_exception_handler              *types.Zval
	user_error_handlers_error_reporting ZendStack
	user_error_handlers                 ZendStack
	user_exception_handlers             ZendStack
	error_handling                      ZendErrorHandlingT
	exception_class                     *types.ClassEntry
	timeout_seconds                     ZendLong
	lambda_count                        int

	iniDirectives         IniDirectives
	modifiedIniDirectives IniDirectives

	error_reporting_ini_entry *ZendIniEntry
	exception                 *types.Object
	prev_exception            **types.Object
	opline_before_exception   *types.ZendOp
	exception_op              [3]types.ZendOp
	current_module            *ModuleEntry
	active                    bool
	flags                     uint8
	assertions                ZendLong

	ht_iterators_count uint32
	ht_iterators_used  uint32
	ht_iterators       *types.ArrayIterator
	ht_iterators_slots []types.ArrayIterator
	arrayIterators     []*types.ArrayIterator

	saved_fpu_cw_ptr        any
	trampoline              types.IFunction
	callTrampolineOp        *types.ZendOp
	each_deprecation_thrown bool
	weakrefs                types.Array
	exception_ignore_args   bool
	reserved                []any
}

func (this *ZendExecutorGlobals) InitTables() {
	this.constantTable = types.NewTable[*ZendConstant](nil)
}
func (this *ZendExecutorGlobals) DestroyTables() {
	this.constantTable.Destroy()
}

/**
 * 辅助方法
 */
func (this *ZendExecutorGlobals) ArrayIterators() []*types.ArrayIterator {
	return this.arrayIterators
}
func (this *ZendExecutorGlobals) ResetArrayIterators() {
	this.arrayIterators = nil
}
func (this *ZendExecutorGlobals) AddArrayIterator(ht *types.Array) uint32 {
	this.arrayIterators = append(this.arrayIterators, ht.Iterator())
	return uint32(len(this.arrayIterators) - 1)
}
func (this *ZendExecutorGlobals) GetArrayIterator(idx uint32) *types.ArrayIterator {
	len_ := uint32(len(this.arrayIterators))
	if idx >= len_ {
		return nil
	}
	return this.arrayIterators[idx]
}
func (this *ZendExecutorGlobals) SetArrayIterator(idx uint32, iterator *types.ArrayIterator) {
	len_ := uint32(len(this.arrayIterators))
	for len_ <= idx {
		this.arrayIterators = append(this.arrayIterators, nil)
		len_++
	}
	this.arrayIterators[idx] = iterator
}
func (this *ZendExecutorGlobals) DelArrayIterator(idx uint32) {
	len_ := uint32(len(this.arrayIterators))
	if idx >= len_ {
		return
	}
	this.arrayIterators[idx] = nil
	// tail
	if idx == len_-1 {
		for idx > 0 && this.arrayIterators[idx-1] == nil {
			idx--
		}
		this.arrayIterators = this.arrayIterators[:idx]
	}
}

func (this *ZendExecutorGlobals) GetHtIteratorsUsed() uint32           { return this.ht_iterators_used }
func (this *ZendExecutorGlobals) GetHtIterators() *types.ArrayIterator { return this.ht_iterators }

func (this *ZendExecutorGlobals) ClassTable() ClassTable     { return this.classTable }
func (this *ZendExecutorGlobals) SetClassTable(t ClassTable) { this.classTable = t }

func (this *ZendExecutorGlobals) FunctionTable() FunctionTable     { return this.functionTable }
func (this *ZendExecutorGlobals) SetFunctionTable(t FunctionTable) { this.functionTable = t }

func (this *ZendExecutorGlobals) ConstantTable() ConstantTable { return this.constantTable }

func (this *ZendExecutorGlobals) IniDirectives() IniDirectives {
	return this.iniDirectives
}
func (this *ZendExecutorGlobals) InitIniDirectives() {
	this.iniDirectives = types.NewTable[*ZendIniEntry](nil)
}
func (this *ZendExecutorGlobals) ModifiedIniDirectives() IniDirectives {
	return this.modifiedIniDirectives
}
func (this *ZendExecutorGlobals) InitModifiedIniDirectives() {
	this.modifiedIniDirectives = types.NewTable[*ZendIniEntry](nil)
}

func (this *ZendExecutorGlobals) VmStack() *VmStack { return &this.vmStack }

// llist
func (this *ZendExecutorGlobals) GetRegularList() *types.Array { return &this.regular_list }

func (this *ZendExecutorGlobals) InitRegularList() {
	this.persistentList = types.NewTable[*types.Resource](ListEntryDtor)
}
func (this *ZendExecutorGlobals) RegularList() ResourceTable { return this.regularList }

func (this *ZendExecutorGlobals) InitPersistentList() {
	this.persistentList = types.NewTable[*types.Resource](PlistEntryDtor)
}
func (this *ZendExecutorGlobals) PersistentList() ResourceTable {
	return this.persistentList
}

/**
 * 以下是自动生成的方法
 */

func (this *ZendExecutorGlobals) GetErrorZval() *types.Zval        { return &this.error_zval }
func (this *ZendExecutorGlobals) GetSymtableCache() []*types.Array { return this.symtable_cache }
func (this *ZendExecutorGlobals) SetSymtableCache(value []*types.Array) {
	this.symtable_cache = value
}
func (this *ZendExecutorGlobals) GetSymtableCacheLimit() **types.Array {
	return this.symtable_cache_limit
}
func (this *ZendExecutorGlobals) SetSymtableCacheLimit(value **types.Array) {
	this.symtable_cache_limit = value
}
func (this *ZendExecutorGlobals) GetSymtableCachePtr() **types.Array {
	return this.symtable_cache_ptr
}
func (this *ZendExecutorGlobals) SetSymtableCachePtr(value **types.Array) {
	this.symtable_cache_ptr = value
}
func (this *ZendExecutorGlobals) GetSymbolTable() *types.Array        { return this.symbol_table }
func (this *ZendExecutorGlobals) SetSymbolTable(value *types.Array)   { this.symbol_table = value }
func (this *ZendExecutorGlobals) GetIncludedFiles() *types.Array      { return this.included_files }
func (this *ZendExecutorGlobals) SetIncludedFiles(value *types.Array) { this.included_files = value }
func (this *ZendExecutorGlobals) GetErrorReporting() int              { return this.error_reporting }
func (this *ZendExecutorGlobals) SetErrorReporting(value int)         { this.error_reporting = value }
func (this *ZendExecutorGlobals) GetExitStatus() int                  { return this.exit_status }
func (this *ZendExecutorGlobals) SetExitStatus(value int)             { this.exit_status = value }
func (this *ZendExecutorGlobals) GetCurrentExecuteData() *ZendExecuteData {
	return this.current_execute_data
}
func (this *ZendExecutorGlobals) SetCurrentExecuteData(value *ZendExecuteData) {
	this.current_execute_data = value
}
func (this *ZendExecutorGlobals) GetFakeScope() *types.ClassEntry      { return this.fake_scope }
func (this *ZendExecutorGlobals) SetFakeScope(value *types.ClassEntry) { this.fake_scope = value }
func (this *ZendExecutorGlobals) GetPrecision() ZendLong               { return this.precision }
func (this *ZendExecutorGlobals) SetPrecision(value ZendLong)          { this.precision = value }
func (this *ZendExecutorGlobals) GetPersistentConstantsCount() uint32 {
	return this.persistent_constants_count
}
func (this *ZendExecutorGlobals) SetPersistentConstantsCount(value uint32) {
	this.persistent_constants_count = value
}
func (this *ZendExecutorGlobals) GetPersistentFunctionsCount() uint32 {
	return this.persistent_functions_count
}
func (this *ZendExecutorGlobals) SetPersistentFunctionsCount(value uint32) {
	this.persistent_functions_count = value
}
func (this *ZendExecutorGlobals) GetPersistentClassesCount() uint32 {
	return this.persistent_classes_count
}
func (this *ZendExecutorGlobals) SetPersistentClassesCount(value uint32) {
	this.persistent_classes_count = value
}
func (this *ZendExecutorGlobals) GetInAutoload() *types.Array      { return this.in_autoload }
func (this *ZendExecutorGlobals) SetInAutoload(value *types.Array) { this.in_autoload = value }
func (this *ZendExecutorGlobals) GetAutoloadFunc() types.IFunction { return this.autoload_func }
func (this *ZendExecutorGlobals) SetAutoloadFunc(value types.IFunction) {
	this.autoload_func = value
}
func (this *ZendExecutorGlobals) GetFullTablesCleanup() bool    { return 0 }
func (this *ZendExecutorGlobals) GetNoExtensions() bool         { return this.no_extensions }
func (this *ZendExecutorGlobals) SetNoExtensions(value bool)    { this.no_extensions = value }
func (this *ZendExecutorGlobals) GetVmInterrupt() bool          { return this.vm_interrupt }
func (this *ZendExecutorGlobals) SetVmInterrupt(value bool)     { this.vm_interrupt = value }
func (this *ZendExecutorGlobals) GetTimedOut() bool             { return this.timed_out }
func (this *ZendExecutorGlobals) SetTimedOut(value bool)        { this.timed_out = value }
func (this *ZendExecutorGlobals) GetHardTimeout() ZendLong      { return this.hard_timeout }
func (this *ZendExecutorGlobals) SetHardTimeout(value ZendLong) { this.hard_timeout = value }
func (this *ZendExecutorGlobals) GetUserErrorHandlerErrorReporting() int {
	return this.user_error_handler_error_reporting
}
func (this *ZendExecutorGlobals) SetUserErrorHandlerErrorReporting(value int) {
	this.user_error_handler_error_reporting = value
}
func (this *ZendExecutorGlobals) GetUserErrorHandler() *types.Zval { return &this.user_error_handler }
func (this *ZendExecutorGlobals) SetUserErrorHandler(value types.Zval) {
	this.user_error_handler = value
}
func (this *ZendExecutorGlobals) GetUserExceptionHandler() *types.Zval {
	return this.user_exception_handler
}
func (this *ZendExecutorGlobals) SetUserExceptionHandler(value *types.Zval) {
	this.user_exception_handler = value
}
func (this *ZendExecutorGlobals) GetUserErrorHandlersErrorReporting() ZendStack {
	return this.user_error_handlers_error_reporting
}
func (this *ZendExecutorGlobals) SetUserErrorHandlersErrorReporting(value ZendStack) {
	this.user_error_handlers_error_reporting = value
}
func (this *ZendExecutorGlobals) GetUserErrorHandlers() *ZendStack { return this.user_error_handlers }
func (this *ZendExecutorGlobals) GetUserExceptionHandlers() *ZendStack {
	return this.user_exception_handlers
}
func (this *ZendExecutorGlobals) GetErrorHandling() ZendErrorHandlingT { return this.error_handling }
func (this *ZendExecutorGlobals) SetErrorHandling(value ZendErrorHandlingT) {
	this.error_handling = value
}
func (this *ZendExecutorGlobals) GetExceptionClass() *types.ClassEntry { return this.exception_class }
func (this *ZendExecutorGlobals) SetExceptionClass(value *types.ClassEntry) {
	this.exception_class = value
}
func (this *ZendExecutorGlobals) GetTimeoutSeconds() ZendLong      { return this.timeout_seconds }
func (this *ZendExecutorGlobals) SetTimeoutSeconds(value ZendLong) { this.timeout_seconds = value }
func (this *ZendExecutorGlobals) GetLambdaCount() int              { return this.lambda_count }
func (this *ZendExecutorGlobals) SetLambdaCount(value int)         { this.lambda_count = value }

func (this *ZendExecutorGlobals) GetErrorReportingIniEntry() *ZendIniEntry {
	return this.error_reporting_ini_entry
}
func (this *ZendExecutorGlobals) SetErrorReportingIniEntry(value *ZendIniEntry) {
	this.error_reporting_ini_entry = value
}
func (this *ZendExecutorGlobals) GetException() *types.Object      { return this.exception }
func (this *ZendExecutorGlobals) SetException(value *types.Object) { this.exception = value }
func (this *ZendExecutorGlobals) GetPrevException() **types.Object { return this.prev_exception }
func (this *ZendExecutorGlobals) SetPrevException(value **types.Object) {
	this.prev_exception = value
}
func (this *ZendExecutorGlobals) GetOplineBeforeException() *types.ZendOp {
	return this.opline_before_exception
}
func (this *ZendExecutorGlobals) SetOplineBeforeException(value *types.ZendOp) {
	this.opline_before_exception = value
}
func (this *ZendExecutorGlobals) GetExceptionOp() *[3]types.ZendOp { return &this.exception_op }
func (this *ZendExecutorGlobals) GetCurrentModule() *ModuleEntry   { return this.current_module }
func (this *ZendExecutorGlobals) SetCurrentModule(value *ModuleEntry) {
	this.current_module = value
}
func (this *ZendExecutorGlobals) GetActive() bool                      { return this.active }
func (this *ZendExecutorGlobals) SetActive(value bool)                 { this.active = value }
func (this *ZendExecutorGlobals) GetFlags() uint8                      { return this.flags }
func (this *ZendExecutorGlobals) SetFlags(value uint8)                 { this.flags = value }
func (this *ZendExecutorGlobals) GetAssertions() ZendLong              { return this.assertions }
func (this *ZendExecutorGlobals) SetAssertions(value ZendLong)         { this.assertions = value }
func (this *ZendExecutorGlobals) GetSavedFpuCwPtr() any                { return this.saved_fpu_cw_ptr }
func (this *ZendExecutorGlobals) SetSavedFpuCwPtr(value any)           { this.saved_fpu_cw_ptr = value }
func (this *ZendExecutorGlobals) GetTrampoline() types.IFunction       { return this.trampoline }
func (this *ZendExecutorGlobals) SetTrampoline(value types.IFunction)  { this.trampoline = value }
func (this *ZendExecutorGlobals) GetCallTrampolineOp() *types.ZendOp   { return this.callTrampolineOp }
func (this *ZendExecutorGlobals) SetCallTrampolineOp(op *types.ZendOp) { this.callTrampolineOp = op }
func (this *ZendExecutorGlobals) GetEachDeprecationThrown() bool {
	return this.each_deprecation_thrown
}
func (this *ZendExecutorGlobals) SetEachDeprecationThrown(value bool) {
	this.each_deprecation_thrown = value
}
func (this *ZendExecutorGlobals) GetWeakrefs() types.Array      { return this.weakrefs }
func (this *ZendExecutorGlobals) SetWeakrefs(value types.Array) { this.weakrefs = value }
func (this *ZendExecutorGlobals) GetExceptionIgnoreArgs() bool {
	return this.exception_ignore_args
}
func (this *ZendExecutorGlobals) SetExceptionIgnoreArgs(value bool) {
	this.exception_ignore_args = value
}
func (this *ZendExecutorGlobals) GetReserved() []any      { return this.reserved }
func (this *ZendExecutorGlobals) SetReserved(value []any) { this.reserved = value }

/* ZendExecutorGlobals.flags */
func (this *ZendExecutorGlobals) AddFlags(value uint8)      { this.flags |= value }
func (this *ZendExecutorGlobals) SubFlags(value uint8)      { this.flags &^= value }
func (this *ZendExecutorGlobals) HasFlags(value uint8) bool { return this.flags&value != 0 }
func (this *ZendExecutorGlobals) SwitchFlags(value uint8, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this ZendExecutorGlobals) IsObjectStoreNoReuse() bool {
	return this.HasFlags(EG_FLAGS_OBJECT_STORE_NO_REUSE)
}
func (this ZendExecutorGlobals) IsInResourceShutdown() bool {
	return this.HasFlags(EG_FLAGS_IN_RESOURCE_SHUTDOWN)
}
func (this *ZendExecutorGlobals) SetIsObjectStoreNoReuse(cond bool) {
	this.SwitchFlags(EG_FLAGS_OBJECT_STORE_NO_REUSE, cond)
}
func (this *ZendExecutorGlobals) SetIsInResourceShutdown(cond bool) {
	this.SwitchFlags(EG_FLAGS_IN_RESOURCE_SHUTDOWN, cond)
}
