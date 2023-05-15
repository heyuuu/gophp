package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

type ClassTable = *types.Table[*types.ClassEntry]
type FunctionTable = *types.Table[types.IFunction]
type ConstantTable = *types.Table[*ZendConstant]
type IniDirectives = *types.Table[*ZendIniEntry]
type ResourceTable = *types.Table[*types.ZendResource]

/**
 * ZendCompilerGlobals
 */
type ZendCompilerGlobals struct {
	loop_var_stack     ZendStack
	active_class_entry *types.ClassEntry
	compiled_filename  *types.String
	zend_lineno        int
	active_op_array    *types.ZendOpArray
	functionTable      FunctionTable
	classTable         ClassTable
	filenamesTable     map[string]string //filenames_table              HashTable

	auto_globals                 *types.Array
	parse_error                  types.ZendBool
	in_compilation               types.ZendBool
	short_tags                   types.ZendBool
	unclean_shutdown             types.ZendBool
	ini_parser_unbuffered_errors types.ZendBool
	open_files                   ZendLlist
	ini_parser_param             *ZendIniParserParam
	skip_shebang                 bool
	increment_lineno             types.ZendBool
	doc_comment                  *types.String
	extra_fn_flags               uint32
	compiler_options             uint32
	context                      ZendOparrayContext
	file_context                 ZendFileContext
	arena                        *ZendArena

	ast                          *ZendAst
	ast_arena                    *ZendArena
	delayed_oplines_stack        ZendStack
	memoized_exprs               *types.Array
	memoize_mode                 int
	map_ptr_base                 any
	map_ptr_size                 int
	map_ptr_last                 int
	delayed_variance_obligations *types.Array
	delayed_autoloads            *types.Array
	rtd_key_counter              uint32
}

func (this *ZendCompilerGlobals) InitTables() {
	this.functionTable = types.NewLcTable[types.IFunction](ZendFunctionDtorEx)
	this.classTable = types.NewLcTable[*types.ClassEntry](nil)
	this.auto_globals = types.NewArray(8)
}

func (this *ZendCompilerGlobals) DestroyTables() {
	this.functionTable.Destroy()
	this.classTable.Destroy()
	this.auto_globals.Destroy()
}

// class table
func (this *ZendCompilerGlobals) ClassTable() ClassTable {
	return this.classTable
}

func (this *ZendCompilerGlobals) FunctionTable() FunctionTable { return this.functionTable }

// getter/setter
func (this *ZendCompilerGlobals) GetLoopVarStack() ZendStack      { return this.loop_var_stack }
func (this *ZendCompilerGlobals) SetLoopVarStack(value ZendStack) { this.loop_var_stack = value }
func (this *ZendCompilerGlobals) GetActiveClassEntry() *types.ClassEntry {
	return this.active_class_entry
}
func (this *ZendCompilerGlobals) SetActiveClassEntry(value *types.ClassEntry) {
	this.active_class_entry = value
}
func (this *ZendCompilerGlobals) GetCompiledFilename() *types.String {
	return this.compiled_filename
}
func (this *ZendCompilerGlobals) SetCompiledFilename(value *types.String) {
	this.compiled_filename = value
}
func (this *ZendCompilerGlobals) GetZendLineno() int                   { return this.zend_lineno }
func (this *ZendCompilerGlobals) SetZendLineno(value int)              { this.zend_lineno = value }
func (this *ZendCompilerGlobals) GetActiveOpArray() *types.ZendOpArray { return this.active_op_array }
func (this *ZendCompilerGlobals) SetActiveOpArray(value *types.ZendOpArray) {
	this.active_op_array = value
}
func (this *ZendCompilerGlobals) GetAutoGlobals() *types.Array          { return this.auto_globals }
func (this *ZendCompilerGlobals) SetAutoGlobals(value *types.Array)     { this.auto_globals = value }
func (this *ZendCompilerGlobals) GetParseError() types.ZendBool         { return this.parse_error }
func (this *ZendCompilerGlobals) SetParseError(value types.ZendBool)    { this.parse_error = value }
func (this *ZendCompilerGlobals) GetInCompilation() types.ZendBool      { return this.in_compilation }
func (this *ZendCompilerGlobals) SetInCompilation(value types.ZendBool) { this.in_compilation = value }
func (this *ZendCompilerGlobals) GetShortTags() types.ZendBool          { return this.short_tags }
func (this *ZendCompilerGlobals) SetShortTags(value types.ZendBool)     { this.short_tags = value }
func (this *ZendCompilerGlobals) GetUncleanShutdown() types.ZendBool    { return this.unclean_shutdown }
func (this *ZendCompilerGlobals) SetUncleanShutdown(value types.ZendBool) {
	this.unclean_shutdown = value
}
func (this *ZendCompilerGlobals) GetIniParserUnbufferedErrors() types.ZendBool {
	return this.ini_parser_unbuffered_errors
}
func (this *ZendCompilerGlobals) SetIniParserUnbufferedErrors(value types.ZendBool) {
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
func (this *ZendCompilerGlobals) GetSkipShebang() types.ZendBool      { return this.skip_shebang }
func (this *ZendCompilerGlobals) SetSkipShebang(value types.ZendBool) { this.skip_shebang = value }
func (this *ZendCompilerGlobals) GetIncrementLineno() types.ZendBool  { return this.increment_lineno }
func (this *ZendCompilerGlobals) SetIncrementLineno(value types.ZendBool) {
	this.increment_lineno = value
}
func (this *ZendCompilerGlobals) GetDocComment() *types.String         { return this.doc_comment }
func (this *ZendCompilerGlobals) SetDocComment(value *types.String)    { this.doc_comment = value }
func (this *ZendCompilerGlobals) GetExtraFnFlags() uint32              { return this.extra_fn_flags }
func (this *ZendCompilerGlobals) SetExtraFnFlags(value uint32)         { this.extra_fn_flags = value }
func (this *ZendCompilerGlobals) GetCompilerOptions() uint32           { return this.compiler_options }
func (this *ZendCompilerGlobals) SetCompilerOptions(value uint32)      { this.compiler_options = value }
func (this *ZendCompilerGlobals) GetContext() *ZendOparrayContext      { return &this.context }
func (this *ZendCompilerGlobals) SetContext(value ZendOparrayContext)  { this.context = value }
func (this *ZendCompilerGlobals) GetFileContext() *ZendFileContext     { return &this.file_context }
func (this *ZendCompilerGlobals) SetFileContext(value ZendFileContext) { this.file_context = value }
func (this *ZendCompilerGlobals) GetArena() *ZendArena                 { return this.arena }
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
func (this *ZendCompilerGlobals) GetMemoizedExprs() *types.Array { return this.memoized_exprs }
func (this *ZendCompilerGlobals) SetMemoizedExprs(value *types.Array) {
	this.memoized_exprs = value
}
func (this *ZendCompilerGlobals) GetMemoizeMode() int      { return this.memoize_mode }
func (this *ZendCompilerGlobals) SetMemoizeMode(value int) { this.memoize_mode = value }
func (this *ZendCompilerGlobals) GetMapPtrBase() any       { return this.map_ptr_base }
func (this *ZendCompilerGlobals) SetMapPtrBase(value any)  { this.map_ptr_base = value }
func (this *ZendCompilerGlobals) GetMapPtrSize() int       { return this.map_ptr_size }
func (this *ZendCompilerGlobals) SetMapPtrSize(value int)  { this.map_ptr_size = value }
func (this *ZendCompilerGlobals) GetMapPtrLast() int       { return this.map_ptr_last }
func (this *ZendCompilerGlobals) SetMapPtrLast(value int)  { this.map_ptr_last = value }
func (this *ZendCompilerGlobals) GetDelayedVarianceObligations() *types.Array {
	return this.delayed_variance_obligations
}
func (this *ZendCompilerGlobals) SetDelayedVarianceObligations(value *types.Array) {
	this.delayed_variance_obligations = value
}
func (this *ZendCompilerGlobals) GetDelayedAutoloads() *types.Array {
	return this.delayed_autoloads
}
func (this *ZendCompilerGlobals) SetDelayedAutoloads(value *types.Array) {
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
	uninitialized_zval   types.Zval
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
	no_extensions                       types.ZendBool
	vm_interrupt                        types.ZendBool
	timed_out                           types.ZendBool
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
	exception                 *types.ZendObject
	prev_exception            **types.ZendObject
	opline_before_exception   *ZendOp
	exception_op              [3]ZendOp
	current_module            *ModuleEntry
	active                    types.ZendBool
	flags                     uint8
	assertions                ZendLong

	ht_iterators_count uint32
	ht_iterators_used  uint32
	ht_iterators       *types.ArrayIterator
	ht_iterators_slots []types.ArrayIterator
	arrayIterators     []*types.ArrayIterator

	saved_fpu_cw_ptr        any
	trampoline              types.IFunction
	call_trampoline_op      ZendOp
	each_deprecation_thrown types.ZendBool
	weakrefs                types.Array
	exception_ignore_args   types.ZendBool
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
	this.persistentList = types.NewTable[*types.ZendResource](ListEntryDtor)
}
func (this *ZendExecutorGlobals) RegularList() ResourceTable { return this.regularList }

func (this *ZendExecutorGlobals) InitPersistentList() {
	this.persistentList = types.NewTable[*types.ZendResource](PlistEntryDtor)
}
func (this *ZendExecutorGlobals) PersistentList() ResourceTable {
	return this.persistentList
}

/**
 * 以下是自动生成的方法
 */

func (this *ZendExecutorGlobals) GetUninitializedZval() types.Zval { return this.uninitialized_zval }
func (this *ZendExecutorGlobals) SetUninitializedZval(value types.Zval) {
	this.uninitialized_zval = value
}
func (this *ZendExecutorGlobals) GetErrorZval() types.Zval         { return this.error_zval }
func (this *ZendExecutorGlobals) SetErrorZval(value types.Zval)    { this.error_zval = value }
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
func (this *ZendExecutorGlobals) GetFullTablesCleanup() types.ZendBool { return 0 }
func (this *ZendExecutorGlobals) GetNoExtensions() types.ZendBool      { return this.no_extensions }
func (this *ZendExecutorGlobals) SetNoExtensions(value types.ZendBool) { this.no_extensions = value }
func (this *ZendExecutorGlobals) GetVmInterrupt() types.ZendBool       { return this.vm_interrupt }
func (this *ZendExecutorGlobals) SetVmInterrupt(value types.ZendBool)  { this.vm_interrupt = value }
func (this *ZendExecutorGlobals) GetTimedOut() types.ZendBool          { return this.timed_out }
func (this *ZendExecutorGlobals) SetTimedOut(value types.ZendBool)     { this.timed_out = value }
func (this *ZendExecutorGlobals) GetHardTimeout() ZendLong             { return this.hard_timeout }
func (this *ZendExecutorGlobals) SetHardTimeout(value ZendLong)        { this.hard_timeout = value }
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
func (this *ZendExecutorGlobals) GetException() *types.ZendObject      { return this.exception }
func (this *ZendExecutorGlobals) SetException(value *types.ZendObject) { this.exception = value }
func (this *ZendExecutorGlobals) GetPrevException() **types.ZendObject { return this.prev_exception }
func (this *ZendExecutorGlobals) SetPrevException(value **types.ZendObject) {
	this.prev_exception = value
}
func (this *ZendExecutorGlobals) GetOplineBeforeException() *ZendOp {
	return this.opline_before_exception
}
func (this *ZendExecutorGlobals) SetOplineBeforeException(value *ZendOp) {
	this.opline_before_exception = value
}
func (this *ZendExecutorGlobals) GetExceptionOp() *[3]ZendOp     { return &this.exception_op }
func (this *ZendExecutorGlobals) GetCurrentModule() *ModuleEntry { return this.current_module }
func (this *ZendExecutorGlobals) SetCurrentModule(value *ModuleEntry) {
	this.current_module = value
}
func (this *ZendExecutorGlobals) GetActive() types.ZendBool           { return this.active }
func (this *ZendExecutorGlobals) SetActive(value types.ZendBool)      { this.active = value }
func (this *ZendExecutorGlobals) GetFlags() uint8                     { return this.flags }
func (this *ZendExecutorGlobals) SetFlags(value uint8)                { this.flags = value }
func (this *ZendExecutorGlobals) GetAssertions() ZendLong             { return this.assertions }
func (this *ZendExecutorGlobals) SetAssertions(value ZendLong)        { this.assertions = value }
func (this *ZendExecutorGlobals) GetSavedFpuCwPtr() any               { return this.saved_fpu_cw_ptr }
func (this *ZendExecutorGlobals) SetSavedFpuCwPtr(value any)          { this.saved_fpu_cw_ptr = value }
func (this *ZendExecutorGlobals) GetTrampoline() types.IFunction      { return this.trampoline }
func (this *ZendExecutorGlobals) SetTrampoline(value types.IFunction) { this.trampoline = value }
func (this *ZendExecutorGlobals) GetCallTrampolineOp() ZendOp         { return this.call_trampoline_op }
func (this *ZendExecutorGlobals) SetCallTrampolineOp(value ZendOp)    { this.call_trampoline_op = value }
func (this *ZendExecutorGlobals) GetEachDeprecationThrown() types.ZendBool {
	return this.each_deprecation_thrown
}
func (this *ZendExecutorGlobals) SetEachDeprecationThrown(value types.ZendBool) {
	this.each_deprecation_thrown = value
}
func (this *ZendExecutorGlobals) GetWeakrefs() types.Array      { return this.weakrefs }
func (this *ZendExecutorGlobals) SetWeakrefs(value types.Array) { this.weakrefs = value }
func (this *ZendExecutorGlobals) GetExceptionIgnoreArgs() types.ZendBool {
	return this.exception_ignore_args
}
func (this *ZendExecutorGlobals) SetExceptionIgnoreArgs(value types.ZendBool) {
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

/**
 * ZendIniScannerGlobals
 */
type ZendIniScannerGlobals struct {
	yy_in        *FileHandle
	yy_out       *FileHandle
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

func (this *ZendIniScannerGlobals) GetYyIn() *FileHandle          { return this.yy_in }
func (this *ZendIniScannerGlobals) SetYyIn(value *FileHandle)     { this.yy_in = value }
func (this *ZendIniScannerGlobals) GetYyOut() *FileHandle         { return this.yy_out }
func (this *ZendIniScannerGlobals) SetYyOut(value *FileHandle)    { this.yy_out = value }
func (this *ZendIniScannerGlobals) GetYyLeng() uint               { return this.yy_leng }
func (this *ZendIniScannerGlobals) SetYyLeng(value uint)          { this.yy_leng = value }
func (this *ZendIniScannerGlobals) GetYyStart() *uint8            { return this.yy_start }
func (this *ZendIniScannerGlobals) SetYyStart(value *uint8)       { this.yy_start = value }
func (this *ZendIniScannerGlobals) GetYyText() *uint8             { return this.yy_text }
func (this *ZendIniScannerGlobals) SetYyText(value *uint8)        { this.yy_text = value }
func (this *ZendIniScannerGlobals) GetYyCursor() *uint8           { return this.yy_cursor }
func (this *ZendIniScannerGlobals) SetYyCursor(value *uint8)      { this.yy_cursor = value }
func (this *ZendIniScannerGlobals) GetYyMarker() *uint8           { return this.yy_marker }
func (this *ZendIniScannerGlobals) SetYyMarker(value *uint8)      { this.yy_marker = value }
func (this *ZendIniScannerGlobals) GetYyLimit() *uint8            { return this.yy_limit }
func (this *ZendIniScannerGlobals) SetYyLimit(value *uint8)       { this.yy_limit = value }
func (this *ZendIniScannerGlobals) GetYyState() int               { return this.yy_state }
func (this *ZendIniScannerGlobals) SetYyState(value int)          { this.yy_state = value }
func (this *ZendIniScannerGlobals) GetStateStack() ZendStack      { return this.state_stack }
func (this *ZendIniScannerGlobals) SetStateStack(value ZendStack) { this.state_stack = value }
func (this *ZendIniScannerGlobals) GetFilename() *byte            { return this.filename }
func (this *ZendIniScannerGlobals) SetFilename(value *byte)       { this.filename = value }
func (this *ZendIniScannerGlobals) GetLineno() int                { return this.lineno }
func (this *ZendIniScannerGlobals) SetLineno(value int)           { this.lineno = value }
func (this *ZendIniScannerGlobals) GetScannerMode() int           { return this.scanner_mode }
func (this *ZendIniScannerGlobals) SetScannerMode(value int)      { this.scanner_mode = value }

/**
 * ZendPhpScannerGlobals
 */
type ZendPhpScannerGlobals struct {
	yy_in                           *FileHandle
	yy_out                          *FileHandle
	yy_leng                         uint
	yy_start                        *uint8
	yy_text                         *uint8
	yy_cursor                       *uint8
	yy_marker                       *uint8
	yy_limit                        *uint8
	yy_state                        int
	state_stack                     b.Stack[int]
	heredoc_label_stack             b.Stack[*ZendHeredocLabel]
	heredoc_scan_ahead              types.ZendBool
	heredoc_indentation             int
	heredoc_indentation_uses_spaces types.ZendBool
	script_org                      *uint8
	script_org_size                 int
	script_filtered                 *uint8
	script_filtered_size            int
	script_encoding                 *ZendEncoding
	scanned_string_len              int
}
