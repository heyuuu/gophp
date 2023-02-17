// <<generate>>

package zend

import b "sik/builtin"

/**
 * ZendCompilerGlobals
 */
type ZendCompilerGlobals struct {
	loop_var_stack               ZendStack
	active_class_entry           *ZendClassEntry
	compiled_filename            *ZendString
	zend_lineno                  int
	active_op_array              *ZendOpArray
	function_table               *HashTable
	class_table                  *HashTable
	filenames_table              HashTable
	auto_globals                 *HashTable
	parse_error                  ZendBool
	in_compilation               ZendBool
	short_tags                   ZendBool
	unclean_shutdown             ZendBool
	ini_parser_unbuffered_errors ZendBool
	open_files                   ZendLlist
	ini_parser_param             *ZendIniParserParam
	skip_shebang                 ZendBool
	increment_lineno             ZendBool
	doc_comment                  *ZendString
	extra_fn_flags               uint32
	compiler_options             uint32
	context                      ZendOparrayContext
	file_context                 ZendFileContext
	arena                        *ZendArena

	InternedStrings *InternedStrings /* request 专用内部字符串 */

	ast                          *ZendAst
	ast_arena                    *ZendArena
	delayed_oplines_stack        ZendStack
	memoized_exprs               *HashTable
	memoize_mode                 int
	map_ptr_base                 any
	map_ptr_size                 int
	map_ptr_last                 int
	delayed_variance_obligations *HashTable
	delayed_autoloads            *HashTable
	rtd_key_counter              uint32
}

func (this *ZendCompilerGlobals) GetLoopVarStack() ZendStack      { return this.loop_var_stack }
func (this *ZendCompilerGlobals) SetLoopVarStack(value ZendStack) { this.loop_var_stack = value }
func (this *ZendCompilerGlobals) GetActiveClassEntry() *ZendClassEntry {
	return this.active_class_entry
}
func (this *ZendCompilerGlobals) SetActiveClassEntry(value *ZendClassEntry) {
	this.active_class_entry = value
}
func (this *ZendCompilerGlobals) GetCompiledFilename() *ZendString { return this.compiled_filename }
func (this *ZendCompilerGlobals) SetCompiledFilename(value *ZendString) {
	this.compiled_filename = value
}
func (this *ZendCompilerGlobals) GetZendLineno() int                  { return this.zend_lineno }
func (this *ZendCompilerGlobals) SetZendLineno(value int)             { this.zend_lineno = value }
func (this *ZendCompilerGlobals) GetActiveOpArray() *ZendOpArray      { return this.active_op_array }
func (this *ZendCompilerGlobals) SetActiveOpArray(value *ZendOpArray) { this.active_op_array = value }
func (this *ZendCompilerGlobals) GetFunctionTable() *HashTable        { return this.function_table }
func (this *ZendCompilerGlobals) SetFunctionTable(value *HashTable)   { this.function_table = value }
func (this *ZendCompilerGlobals) GetClassTable() *HashTable           { return this.class_table }
func (this *ZendCompilerGlobals) SetClassTable(value *HashTable)      { this.class_table = value }
func (this *ZendCompilerGlobals) GetFilenamesTable() HashTable        { return this.filenames_table }
func (this *ZendCompilerGlobals) SetFilenamesTable(value HashTable)   { this.filenames_table = value }
func (this *ZendCompilerGlobals) GetAutoGlobals() *HashTable          { return this.auto_globals }
func (this *ZendCompilerGlobals) SetAutoGlobals(value *HashTable)     { this.auto_globals = value }
func (this *ZendCompilerGlobals) GetParseError() ZendBool             { return this.parse_error }
func (this *ZendCompilerGlobals) SetParseError(value ZendBool)        { this.parse_error = value }
func (this *ZendCompilerGlobals) GetInCompilation() ZendBool          { return this.in_compilation }
func (this *ZendCompilerGlobals) SetInCompilation(value ZendBool)     { this.in_compilation = value }
func (this *ZendCompilerGlobals) GetShortTags() ZendBool              { return this.short_tags }
func (this *ZendCompilerGlobals) SetShortTags(value ZendBool)         { this.short_tags = value }
func (this *ZendCompilerGlobals) GetUncleanShutdown() ZendBool        { return this.unclean_shutdown }
func (this *ZendCompilerGlobals) SetUncleanShutdown(value ZendBool)   { this.unclean_shutdown = value }
func (this *ZendCompilerGlobals) GetIniParserUnbufferedErrors() ZendBool {
	return this.ini_parser_unbuffered_errors
}
func (this *ZendCompilerGlobals) SetIniParserUnbufferedErrors(value ZendBool) {
	this.ini_parser_unbuffered_errors = value
}
func (this *ZendCompilerGlobals) GetOpenFiles() ZendLlist      { return this.open_files }
func (this *ZendCompilerGlobals) SetOpenFiles(value ZendLlist) { this.open_files = value }
func (this *ZendCompilerGlobals) GetIniParserParam() *ZendIniParserParam {
	return this.ini_parser_param
}
func (this *ZendCompilerGlobals) SetIniParserParam(value *ZendIniParserParam) {
	this.ini_parser_param = value
}
func (this *ZendCompilerGlobals) GetSkipShebang() ZendBool             { return this.skip_shebang }
func (this *ZendCompilerGlobals) SetSkipShebang(value ZendBool)        { this.skip_shebang = value }
func (this *ZendCompilerGlobals) GetIncrementLineno() ZendBool         { return this.increment_lineno }
func (this *ZendCompilerGlobals) SetIncrementLineno(value ZendBool)    { this.increment_lineno = value }
func (this *ZendCompilerGlobals) GetDocComment() *ZendString           { return this.doc_comment }
func (this *ZendCompilerGlobals) SetDocComment(value *ZendString)      { this.doc_comment = value }
func (this *ZendCompilerGlobals) GetExtraFnFlags() uint32              { return this.extra_fn_flags }
func (this *ZendCompilerGlobals) SetExtraFnFlags(value uint32)         { this.extra_fn_flags = value }
func (this *ZendCompilerGlobals) GetCompilerOptions() uint32           { return this.compiler_options }
func (this *ZendCompilerGlobals) SetCompilerOptions(value uint32)      { this.compiler_options = value }
func (this *ZendCompilerGlobals) GetContext() ZendOparrayContext       { return this.context }
func (this *ZendCompilerGlobals) SetContext(value ZendOparrayContext)  { this.context = value }
func (this *ZendCompilerGlobals) GetFileContext() ZendFileContext      { return this.file_context }
func (this *ZendCompilerGlobals) SetFileContext(value ZendFileContext) { this.file_context = value }
func (this *ZendCompilerGlobals) GetArena() *ZendArena                 { return this.arena }
func (this *ZendCompilerGlobals) SetArena(value *ZendArena)            { this.arena = value }
func (this *ZendCompilerGlobals) GetAst() *ZendAst                     { return this.ast }
func (this *ZendCompilerGlobals) SetAst(value *ZendAst)                { this.ast = value }
func (this *ZendCompilerGlobals) GetAstArena() *ZendArena              { return this.ast_arena }
func (this *ZendCompilerGlobals) SetAstArena(value *ZendArena)         { this.ast_arena = value }
func (this *ZendCompilerGlobals) GetDelayedOplinesStack() ZendStack {
	return this.delayed_oplines_stack
}
func (this *ZendCompilerGlobals) SetDelayedOplinesStack(value ZendStack) {
	this.delayed_oplines_stack = value
}
func (this *ZendCompilerGlobals) GetMemoizedExprs() *HashTable      { return this.memoized_exprs }
func (this *ZendCompilerGlobals) SetMemoizedExprs(value *HashTable) { this.memoized_exprs = value }
func (this *ZendCompilerGlobals) GetMemoizeMode() int               { return this.memoize_mode }
func (this *ZendCompilerGlobals) SetMemoizeMode(value int)          { this.memoize_mode = value }
func (this *ZendCompilerGlobals) GetMapPtrBase() any                { return this.map_ptr_base }
func (this *ZendCompilerGlobals) SetMapPtrBase(value any)           { this.map_ptr_base = value }
func (this *ZendCompilerGlobals) GetMapPtrSize() int                { return this.map_ptr_size }
func (this *ZendCompilerGlobals) SetMapPtrSize(value int)           { this.map_ptr_size = value }
func (this *ZendCompilerGlobals) GetMapPtrLast() int                { return this.map_ptr_last }
func (this *ZendCompilerGlobals) SetMapPtrLast(value int)           { this.map_ptr_last = value }
func (this *ZendCompilerGlobals) GetDelayedVarianceObligations() *HashTable {
	return this.delayed_variance_obligations
}
func (this *ZendCompilerGlobals) SetDelayedVarianceObligations(value *HashTable) {
	this.delayed_variance_obligations = value
}
func (this *ZendCompilerGlobals) GetDelayedAutoloads() *HashTable { return this.delayed_autoloads }
func (this *ZendCompilerGlobals) SetDelayedAutoloads(value *HashTable) {
	this.delayed_autoloads = value
}
func (this *ZendCompilerGlobals) GetRtdKeyCounter() uint32      { return this.rtd_key_counter }
func (this *ZendCompilerGlobals) SetRtdKeyCounter(value uint32) { this.rtd_key_counter = value }

/* ZendCompilerGlobals.extra_fn_flags */
func (this *ZendCompilerGlobals) AddExtraFnFlags(value uint32) { this.extra_fn_flags |= value }
func (this *ZendCompilerGlobals) SubExtraFnFlags(value uint32) { this.extra_fn_flags &^= value }
func (this *ZendCompilerGlobals) HasExtraFnFlags(value uint32) bool {
	return this.extra_fn_flags&value != 0
}
func (this *ZendCompilerGlobals) SwitchExtraFnFlags(value uint32, cond bool) {
	if cond {
		this.AddExtraFnFlags(value)
	} else {
		this.SubExtraFnFlags(value)
	}
}

/**
 * ZendExecutorGlobals
 */
type ZendExecutorGlobals struct {
	uninitialized_zval                  Zval
	error_zval                          Zval
	symtable_cache                      []*ZendArray
	symtable_cache_limit                **ZendArray
	symtable_cache_ptr                  **ZendArray
	symbol_table                        ZendArray
	included_files                      HashTable
	bailout                             *JMP_BUF
	error_reporting                     int
	exit_status                         int
	function_table                      *HashTable
	class_table                         *HashTable
	zend_constants                      *HashTable
	vm_stack_top                        *Zval
	vm_stack_end                        *Zval
	vm_stack                            ZendVmStack
	vm_stack_page_size                  int
	current_execute_data                *ZendExecuteData
	fake_scope                          *ZendClassEntry
	precision                           ZendLong
	ticks_count                         int
	persistent_constants_count          uint32
	persistent_functions_count          uint32
	persistent_classes_count            uint32
	in_autoload                         *HashTable
	autoload_func                       *ZendFunction
	full_tables_cleanup                 ZendBool
	no_extensions                       ZendBool
	vm_interrupt                        ZendBool
	timed_out                           ZendBool
	hard_timeout                        ZendLong
	regular_list                        HashTable
	persistent_list                     HashTable
	user_error_handler_error_reporting  int
	user_error_handler                  Zval
	user_exception_handler              Zval
	user_error_handlers_error_reporting ZendStack
	user_error_handlers                 ZendStack
	user_exception_handlers             ZendStack
	error_handling                      ZendErrorHandlingT
	exception_class                     *ZendClassEntry
	timeout_seconds                     ZendLong
	lambda_count                        int
	ini_directives                      *HashTable
	modified_ini_directives             *HashTable
	error_reporting_ini_entry           *ZendIniEntry
	objects_store                       ZendObjectsStore
	exception                           *ZendObject
	prev_exception                      **ZendObject
	opline_before_exception             *ZendOp
	exception_op                        []ZendOp
	current_module                      *ZendModuleEntry
	active                              ZendBool
	flags                               ZendUchar
	assertions                          ZendLong
	ht_iterators_count                  uint32
	ht_iterators_used                   uint32
	ht_iterators                        *HashTableIterator
	ht_iterators_slots                  []HashTableIterator
	saved_fpu_cw_ptr                    any
	trampoline                          ZendFunction
	call_trampoline_op                  ZendOp
	each_deprecation_thrown             ZendBool
	weakrefs                            HashTable
	exception_ignore_args               ZendBool
	reserved                            []any
}

/**
 * 辅助方法
 */
func (this *ZendExecutorGlobals) HtIterators() []HashTableIterator {
	// todo 待调整
	return b.CastSlice(this.ht_iterators, this.ht_iterators_count)
}

/**
 * 以下是自动生成的方法
 */

func (this *ZendExecutorGlobals) GetUninitializedZval() Zval          { return this.uninitialized_zval }
func (this *ZendExecutorGlobals) SetUninitializedZval(value Zval)     { this.uninitialized_zval = value }
func (this *ZendExecutorGlobals) GetErrorZval() Zval                  { return this.error_zval }
func (this *ZendExecutorGlobals) SetErrorZval(value Zval)             { this.error_zval = value }
func (this *ZendExecutorGlobals) GetSymtableCache() []*ZendArray      { return this.symtable_cache }
func (this *ZendExecutorGlobals) SetSymtableCache(value []*ZendArray) { this.symtable_cache = value }
func (this *ZendExecutorGlobals) GetSymtableCacheLimit() **ZendArray {
	return this.symtable_cache_limit
}
func (this *ZendExecutorGlobals) SetSymtableCacheLimit(value **ZendArray) {
	this.symtable_cache_limit = value
}
func (this *ZendExecutorGlobals) GetSymtableCachePtr() **ZendArray { return this.symtable_cache_ptr }
func (this *ZendExecutorGlobals) SetSymtableCachePtr(value **ZendArray) {
	this.symtable_cache_ptr = value
}
func (this *ZendExecutorGlobals) GetSymbolTable() *ZendArray        { return &this.symbol_table }
func (this *ZendExecutorGlobals) SetSymbolTable(value ZendArray)    { this.symbol_table = value }
func (this *ZendExecutorGlobals) GetIncludedFiles() HashTable       { return this.included_files }
func (this *ZendExecutorGlobals) SetIncludedFiles(value HashTable)  { this.included_files = value }
func (this *ZendExecutorGlobals) GetBailout() *JMP_BUF              { return this.bailout }
func (this *ZendExecutorGlobals) SetBailout(value *JMP_BUF)         { this.bailout = value }
func (this *ZendExecutorGlobals) GetErrorReporting() int            { return this.error_reporting }
func (this *ZendExecutorGlobals) SetErrorReporting(value int)       { this.error_reporting = value }
func (this *ZendExecutorGlobals) GetExitStatus() int                { return this.exit_status }
func (this *ZendExecutorGlobals) SetExitStatus(value int)           { this.exit_status = value }
func (this *ZendExecutorGlobals) GetFunctionTable() *HashTable      { return this.function_table }
func (this *ZendExecutorGlobals) SetFunctionTable(value *HashTable) { this.function_table = value }
func (this *ZendExecutorGlobals) GetClassTable() *HashTable         { return this.class_table }
func (this *ZendExecutorGlobals) SetClassTable(value *HashTable)    { this.class_table = value }
func (this *ZendExecutorGlobals) GetZendConstants() *HashTable      { return this.zend_constants }
func (this *ZendExecutorGlobals) SetZendConstants(value *HashTable) { this.zend_constants = value }
func (this *ZendExecutorGlobals) GetVmStackTop() *Zval              { return this.vm_stack_top }
func (this *ZendExecutorGlobals) SetVmStackTop(value *Zval)         { this.vm_stack_top = value }
func (this *ZendExecutorGlobals) GetVmStackEnd() *Zval              { return this.vm_stack_end }
func (this *ZendExecutorGlobals) SetVmStackEnd(value *Zval)         { this.vm_stack_end = value }
func (this *ZendExecutorGlobals) GetVmStack() ZendVmStack           { return this.vm_stack }
func (this *ZendExecutorGlobals) SetVmStack(value ZendVmStack)      { this.vm_stack = value }
func (this *ZendExecutorGlobals) GetVmStackPageSize() int           { return this.vm_stack_page_size }
func (this *ZendExecutorGlobals) SetVmStackPageSize(value int)      { this.vm_stack_page_size = value }
func (this *ZendExecutorGlobals) GetCurrentExecuteData() *ZendExecuteData {
	return this.current_execute_data
}
func (this *ZendExecutorGlobals) SetCurrentExecuteData(value *ZendExecuteData) {
	this.current_execute_data = value
}
func (this *ZendExecutorGlobals) GetFakeScope() *ZendClassEntry      { return this.fake_scope }
func (this *ZendExecutorGlobals) SetFakeScope(value *ZendClassEntry) { this.fake_scope = value }
func (this *ZendExecutorGlobals) GetPrecision() ZendLong             { return this.precision }
func (this *ZendExecutorGlobals) SetPrecision(value ZendLong)        { this.precision = value }
func (this *ZendExecutorGlobals) GetTicksCount() int                 { return this.ticks_count }
func (this *ZendExecutorGlobals) SetTicksCount(value int)            { this.ticks_count = value }
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
func (this *ZendExecutorGlobals) GetInAutoload() *HashTable           { return this.in_autoload }
func (this *ZendExecutorGlobals) SetInAutoload(value *HashTable)      { this.in_autoload = value }
func (this *ZendExecutorGlobals) GetAutoloadFunc() *ZendFunction      { return this.autoload_func }
func (this *ZendExecutorGlobals) SetAutoloadFunc(value *ZendFunction) { this.autoload_func = value }
func (this *ZendExecutorGlobals) GetFullTablesCleanup() ZendBool      { return this.full_tables_cleanup }
func (this *ZendExecutorGlobals) SetFullTablesCleanup(value ZendBool) {
	this.full_tables_cleanup = value
}
func (this *ZendExecutorGlobals) GetNoExtensions() ZendBool         { return this.no_extensions }
func (this *ZendExecutorGlobals) SetNoExtensions(value ZendBool)    { this.no_extensions = value }
func (this *ZendExecutorGlobals) GetVmInterrupt() ZendBool          { return this.vm_interrupt }
func (this *ZendExecutorGlobals) SetVmInterrupt(value ZendBool)     { this.vm_interrupt = value }
func (this *ZendExecutorGlobals) GetTimedOut() ZendBool             { return this.timed_out }
func (this *ZendExecutorGlobals) SetTimedOut(value ZendBool)        { this.timed_out = value }
func (this *ZendExecutorGlobals) GetHardTimeout() ZendLong          { return this.hard_timeout }
func (this *ZendExecutorGlobals) SetHardTimeout(value ZendLong)     { this.hard_timeout = value }
func (this *ZendExecutorGlobals) GetRegularList() HashTable         { return this.regular_list }
func (this *ZendExecutorGlobals) SetRegularList(value HashTable)    { this.regular_list = value }
func (this *ZendExecutorGlobals) GetPersistentList() HashTable      { return this.persistent_list }
func (this *ZendExecutorGlobals) SetPersistentList(value HashTable) { this.persistent_list = value }
func (this *ZendExecutorGlobals) GetUserErrorHandlerErrorReporting() int {
	return this.user_error_handler_error_reporting
}
func (this *ZendExecutorGlobals) SetUserErrorHandlerErrorReporting(value int) {
	this.user_error_handler_error_reporting = value
}
func (this *ZendExecutorGlobals) GetUserErrorHandler() Zval      { return this.user_error_handler }
func (this *ZendExecutorGlobals) SetUserErrorHandler(value Zval) { this.user_error_handler = value }
func (this *ZendExecutorGlobals) GetUserExceptionHandler() Zval  { return this.user_exception_handler }
func (this *ZendExecutorGlobals) SetUserExceptionHandler(value Zval) {
	this.user_exception_handler = value
}
func (this *ZendExecutorGlobals) GetUserErrorHandlersErrorReporting() ZendStack {
	return this.user_error_handlers_error_reporting
}
func (this *ZendExecutorGlobals) SetUserErrorHandlersErrorReporting(value ZendStack) {
	this.user_error_handlers_error_reporting = value
}
func (this *ZendExecutorGlobals) GetUserErrorHandlers() ZendStack { return this.user_error_handlers }
func (this *ZendExecutorGlobals) SetUserErrorHandlers(value ZendStack) {
	this.user_error_handlers = value
}
func (this *ZendExecutorGlobals) GetUserExceptionHandlers() ZendStack {
	return this.user_exception_handlers
}
func (this *ZendExecutorGlobals) SetUserExceptionHandlers(value ZendStack) {
	this.user_exception_handlers = value
}
func (this *ZendExecutorGlobals) GetErrorHandling() ZendErrorHandlingT { return this.error_handling }
func (this *ZendExecutorGlobals) SetErrorHandling(value ZendErrorHandlingT) {
	this.error_handling = value
}
func (this *ZendExecutorGlobals) GetExceptionClass() *ZendClassEntry { return this.exception_class }
func (this *ZendExecutorGlobals) SetExceptionClass(value *ZendClassEntry) {
	this.exception_class = value
}
func (this *ZendExecutorGlobals) GetTimeoutSeconds() ZendLong       { return this.timeout_seconds }
func (this *ZendExecutorGlobals) SetTimeoutSeconds(value ZendLong)  { this.timeout_seconds = value }
func (this *ZendExecutorGlobals) GetLambdaCount() int               { return this.lambda_count }
func (this *ZendExecutorGlobals) SetLambdaCount(value int)          { this.lambda_count = value }
func (this *ZendExecutorGlobals) GetIniDirectives() *HashTable      { return this.ini_directives }
func (this *ZendExecutorGlobals) SetIniDirectives(value *HashTable) { this.ini_directives = value }
func (this *ZendExecutorGlobals) GetModifiedIniDirectives() *HashTable {
	return this.modified_ini_directives
}
func (this *ZendExecutorGlobals) SetModifiedIniDirectives(value *HashTable) {
	this.modified_ini_directives = value
}
func (this *ZendExecutorGlobals) GetErrorReportingIniEntry() *ZendIniEntry {
	return this.error_reporting_ini_entry
}
func (this *ZendExecutorGlobals) SetErrorReportingIniEntry(value *ZendIniEntry) {
	this.error_reporting_ini_entry = value
}
func (this *ZendExecutorGlobals) GetObjectsStore() ZendObjectsStore      { return this.objects_store }
func (this *ZendExecutorGlobals) SetObjectsStore(value ZendObjectsStore) { this.objects_store = value }
func (this *ZendExecutorGlobals) GetException() *ZendObject              { return this.exception }
func (this *ZendExecutorGlobals) SetException(value *ZendObject)         { this.exception = value }
func (this *ZendExecutorGlobals) GetPrevException() **ZendObject         { return this.prev_exception }
func (this *ZendExecutorGlobals) SetPrevException(value **ZendObject)    { this.prev_exception = value }
func (this *ZendExecutorGlobals) GetOplineBeforeException() *ZendOp {
	return this.opline_before_exception
}
func (this *ZendExecutorGlobals) SetOplineBeforeException(value *ZendOp) {
	this.opline_before_exception = value
}
func (this *ZendExecutorGlobals) GetExceptionOp() []ZendOp           { return this.exception_op }
func (this *ZendExecutorGlobals) SetExceptionOp(value []ZendOp)      { this.exception_op = value }
func (this *ZendExecutorGlobals) GetCurrentModule() *ZendModuleEntry { return this.current_module }
func (this *ZendExecutorGlobals) SetCurrentModule(value *ZendModuleEntry) {
	this.current_module = value
}
func (this *ZendExecutorGlobals) GetActive() ZendBool                     { return this.active }
func (this *ZendExecutorGlobals) SetActive(value ZendBool)                { this.active = value }
func (this *ZendExecutorGlobals) GetFlags() ZendUchar                     { return this.flags }
func (this *ZendExecutorGlobals) SetFlags(value ZendUchar)                { this.flags = value }
func (this *ZendExecutorGlobals) GetAssertions() ZendLong                 { return this.assertions }
func (this *ZendExecutorGlobals) SetAssertions(value ZendLong)            { this.assertions = value }
func (this *ZendExecutorGlobals) GetHtIteratorsCount() uint32             { return this.ht_iterators_count }
func (this *ZendExecutorGlobals) SetHtIteratorsCount(value uint32)        { this.ht_iterators_count = value }
func (this *ZendExecutorGlobals) GetHtIteratorsUsed() uint32              { return this.ht_iterators_used }
func (this *ZendExecutorGlobals) SetHtIteratorsUsed(value uint32)         { this.ht_iterators_used = value }
func (this *ZendExecutorGlobals) GetHtIterators() *HashTableIterator      { return this.ht_iterators }
func (this *ZendExecutorGlobals) SetHtIterators(value *HashTableIterator) { this.ht_iterators = value }
func (this *ZendExecutorGlobals) GetHtIteratorsSlots() []HashTableIterator {
	return this.ht_iterators_slots
}
func (this *ZendExecutorGlobals) SetHtIteratorsSlots(value []HashTableIterator) {
	this.ht_iterators_slots = value
}
func (this *ZendExecutorGlobals) GetSavedFpuCwPtr() any            { return this.saved_fpu_cw_ptr }
func (this *ZendExecutorGlobals) SetSavedFpuCwPtr(value any)       { this.saved_fpu_cw_ptr = value }
func (this *ZendExecutorGlobals) GetTrampoline() ZendFunction      { return this.trampoline }
func (this *ZendExecutorGlobals) SetTrampoline(value ZendFunction) { this.trampoline = value }
func (this *ZendExecutorGlobals) GetCallTrampolineOp() ZendOp      { return this.call_trampoline_op }
func (this *ZendExecutorGlobals) SetCallTrampolineOp(value ZendOp) { this.call_trampoline_op = value }
func (this *ZendExecutorGlobals) GetEachDeprecationThrown() ZendBool {
	return this.each_deprecation_thrown
}
func (this *ZendExecutorGlobals) SetEachDeprecationThrown(value ZendBool) {
	this.each_deprecation_thrown = value
}
func (this *ZendExecutorGlobals) GetWeakrefs() HashTable           { return this.weakrefs }
func (this *ZendExecutorGlobals) SetWeakrefs(value HashTable)      { this.weakrefs = value }
func (this *ZendExecutorGlobals) GetExceptionIgnoreArgs() ZendBool { return this.exception_ignore_args }
func (this *ZendExecutorGlobals) SetExceptionIgnoreArgs(value ZendBool) {
	this.exception_ignore_args = value
}
func (this *ZendExecutorGlobals) GetReserved() []any      { return this.reserved }
func (this *ZendExecutorGlobals) SetReserved(value []any) { this.reserved = value }

/* ZendExecutorGlobals.flags */
func (this *ZendExecutorGlobals) AddFlags(value ZendUchar)      { this.flags |= value }
func (this *ZendExecutorGlobals) SubFlags(value ZendUchar)      { this.flags &^= value }
func (this *ZendExecutorGlobals) HasFlags(value ZendUchar) bool { return this.flags&value != 0 }
func (this *ZendExecutorGlobals) SwitchFlags(value ZendUchar, cond bool) {
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

/**
 * ZendIniScannerGlobals
 */
type ZendIniScannerGlobals struct {
	yy_in        *ZendFileHandle
	yy_out       *ZendFileHandle
	yy_leng      uint
	yy_start     *uint8
	yy_text      *uint8
	yy_cursor    *uint8
	yy_marker    *uint8
	yy_limit     *uint8
	yy_state     int
	state_stack  ZendStack
	filename     *byte
	lineno       int
	scanner_mode int
}

func (this *ZendIniScannerGlobals) GetYyIn() *ZendFileHandle       { return this.yy_in }
func (this *ZendIniScannerGlobals) SetYyIn(value *ZendFileHandle)  { this.yy_in = value }
func (this *ZendIniScannerGlobals) GetYyOut() *ZendFileHandle      { return this.yy_out }
func (this *ZendIniScannerGlobals) SetYyOut(value *ZendFileHandle) { this.yy_out = value }
func (this *ZendIniScannerGlobals) GetYyLeng() uint                { return this.yy_leng }
func (this *ZendIniScannerGlobals) SetYyLeng(value uint)           { this.yy_leng = value }
func (this *ZendIniScannerGlobals) GetYyStart() *uint8             { return this.yy_start }
func (this *ZendIniScannerGlobals) SetYyStart(value *uint8)        { this.yy_start = value }
func (this *ZendIniScannerGlobals) GetYyText() *uint8              { return this.yy_text }
func (this *ZendIniScannerGlobals) SetYyText(value *uint8)         { this.yy_text = value }
func (this *ZendIniScannerGlobals) GetYyCursor() *uint8            { return this.yy_cursor }
func (this *ZendIniScannerGlobals) SetYyCursor(value *uint8)       { this.yy_cursor = value }
func (this *ZendIniScannerGlobals) GetYyMarker() *uint8            { return this.yy_marker }
func (this *ZendIniScannerGlobals) SetYyMarker(value *uint8)       { this.yy_marker = value }
func (this *ZendIniScannerGlobals) GetYyLimit() *uint8             { return this.yy_limit }
func (this *ZendIniScannerGlobals) SetYyLimit(value *uint8)        { this.yy_limit = value }
func (this *ZendIniScannerGlobals) GetYyState() int                { return this.yy_state }
func (this *ZendIniScannerGlobals) SetYyState(value int)           { this.yy_state = value }
func (this *ZendIniScannerGlobals) GetStateStack() ZendStack       { return this.state_stack }
func (this *ZendIniScannerGlobals) SetStateStack(value ZendStack)  { this.state_stack = value }
func (this *ZendIniScannerGlobals) GetFilename() *byte             { return this.filename }
func (this *ZendIniScannerGlobals) SetFilename(value *byte)        { this.filename = value }
func (this *ZendIniScannerGlobals) GetLineno() int                 { return this.lineno }
func (this *ZendIniScannerGlobals) SetLineno(value int)            { this.lineno = value }
func (this *ZendIniScannerGlobals) GetScannerMode() int            { return this.scanner_mode }
func (this *ZendIniScannerGlobals) SetScannerMode(value int)       { this.scanner_mode = value }

/**
 * ZendPhpScannerGlobals
 */
type ZendPhpScannerGlobals struct {
	yy_in                           *ZendFileHandle
	yy_out                          *ZendFileHandle
	yy_leng                         uint
	yy_start                        *uint8
	yy_text                         *uint8
	yy_cursor                       *uint8
	yy_marker                       *uint8
	yy_limit                        *uint8
	yy_state                        int
	state_stack                     b.Stack[int]
	heredoc_label_stack             b.Stack[*ZendHeredocLabel]
	heredoc_scan_ahead              ZendBool
	heredoc_indentation             int
	heredoc_indentation_uses_spaces ZendBool
	script_org                      *uint8
	script_org_size                 int
	script_filtered                 *uint8
	script_filtered_size            int
	script_encoding                 *ZendEncoding
	scanned_string_len              int
}
