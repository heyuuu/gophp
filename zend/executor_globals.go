package zend

import (
	"github.com/heyuuu/gophp/php/contracts"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

var ExecutorGlobals ZendExecutorGlobals

func EG__() *ZendExecutorGlobals { return &ExecutorGlobals }

// ZendExecutorGlobals
const SYMTABLE_CACHE_SIZE = 32

type ZendExecutorGlobals struct {
	errorZval        types.Zval
	symtableCache    [SYMTABLE_CACHE_SIZE]*types.Array
	symtableCacheIdx int
	symbolTable      *types.Array
	includedFiles    *types.Array
	errorReporting   int
	exitStatus       int
	functionTable    FunctionTable
	classTable       ClassTable
	constantTable    ConstantTable
	vmStack          VmStack

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

/**
 * 生命周期
 */
var _ contracts.ModuleLifeCycle = (*ZendExecutorGlobals)(nil)

func (eg *ZendExecutorGlobals) StartUp() {
	eg.constantTable = types.NewTable[*ZendConstant](nil)
}
func (eg *ZendExecutorGlobals) Shutdown() {
	eg.constantTable.Destroy()
}
func (eg *ZendExecutorGlobals) Activate() {
	ZendInitFpu()

	eg.errorZval.SetIsError()
	//eg.symtable_cache = [SYMTABLE_CACHE_SIZE]*types.Array{}
	eg.symtableCacheIdx = 0
	eg.symbolTable = types.NewArrayCap(64)

	eg.no_extensions = false
	eg.functionTable = CG__().FunctionTable()
	eg.classTable = CG__().ClassTable()
	eg.SetInAutoload(nil)
	eg.SetAutoloadFunc(nil)
	eg.SetErrorHandling(EH_NORMAL)
	eg.SetFlags(EG_FLAGS_INITIAL)
	eg.VmStack().Reset()
	ZendExtensions.Apply(func(ext *ZendExtension) {
		if ext.GetActivate() != nil {
			ext.GetActivate()()
		}
	})

	eg.includedFiles = types.NewArrayCap(8)
	eg.GetUserErrorHandler().SetUndef()
	eg.GetUserExceptionHandler().SetUndef()
	eg.SetCurrentExecuteData(nil)
	eg.GetUserErrorHandlersErrorReporting().Init()
	eg.GetUserErrorHandlers().Init()
	eg.GetUserExceptionHandlers().Init()
	eg.SetVmInterrupt(0)
	eg.SetTimedOut(0)
	eg.SetException(nil)
	eg.SetPrevException(nil)
	eg.SetFakeScope(nil)
	eg.GetTrampoline().SetFunctionName("")
	eg.ResetArrayIterators()
	eg.SetEachDeprecationThrown(0)
	eg.SetPersistentConstantsCount(uint32(eg.ConstantTable().Len()))
	eg.SetPersistentFunctionsCount(uint32(eg.FunctionTable().Len()))
	eg.SetPersistentClassesCount(uint32(eg.ClassTable().Len()))

	eg.active = true
}
func (eg *ZendExecutorGlobals) Deactivate() {
	var fastShutdown bool = IsZendMm() != 0 && !EG__().GetFullTablesCleanup()

	faults.Try(func() {
		CG__().GetOpenFiles().Destroy()
	})

	eg.SetIsInResourceShutdown(true)

	faults.Try(func() {
		eg.RegularList().ForeachReserve(func(_ string, res *types.Resource) {
			if res.GetType() >= 0 {
				ZendResourceDtor(res)
			}
		})
	})

	/* No PHP callback functions should be called after this point. */
	eg.active = false
	if !fastShutdown {
		eg.GetSymbolTable().Clean()

		/* Release static properties and static variables prior to the final GC run,
		 * as they may hold GC roots. */
		eg.FunctionTable().ForeachReserve(func(_ string, f types.IFunction) {
			if f.GetType() == ZEND_INTERNAL_FUNCTION {
				return
			}

			opArray := f.GetOpArray()
			if opArray.GetStaticVariables() != nil {
				var ht *types.Array = opArray.GetStaticVariablesPtr()
				if ht != nil {
					//if ht.DelRefcount() == 0 {
					//	ht.Destroy()
					//}
					opArray.SetStaticVariablesPtr(nil)
				}
			}
		})

		eg.ClassTable().ForeachReserve(func(_ string, ce *types.ClassEntry) {
			if ce.GetDefaultStaticMembersCount() != 0 {
				//ZendCleanupInternalClassData(ce)
			}
			if ce.IsHasStaticInMethods() {
				ce.FunctionTable().Foreach(func(_ string, f types.IFunction) {
					if f.GetType() == ZEND_USER_FUNCTION {
						opArray := f.GetOpArray()
						if opArray.GetStaticVariables() != nil {
							var ht *types.Array = opArray.GetStaticVariablesPtr()
							if ht != nil {
								//if ht.DelRefcount() == 0 {
								//	ht.Destroy()
								//}
								opArray.SetStaticVariablesPtr(nil)
							}
						}
					}
				})
			}
		})

		/* Also release error and exception handlers, which may hold objects. */
		eg.GetUserErrorHandler().SetUndef()
		eg.GetUserExceptionHandler().SetUndef()
		ZendStackClean(eg.GetUserErrorHandlersErrorReporting(), nil, 1)
		ZendStackClean(eg.GetUserErrorHandlers(), nil, 1)
		ZendStackClean(eg.GetUserExceptionHandlers(), nil, 1)
	}

	// notice: 无需主动调用析构函数，使用自动析构代替
	//ZendObjectsStoreFreeObjectStorage(eg.GetObjectsStore(), fast_shutdown)
	//ZendWeakrefsShutdown()
	faults.Try(func() {
		ZendExtensions.Apply(func(ext *ZendExtension) {
			if ext.GetDeactivate() != nil {
				ext.GetDeactivate()()
			}
		})
	})

	if fastShutdown {
		/* Fast Request Shutdown
		 * =====================
		 * Zend Memory Manager frees memory by its own. We don't have to free
		 * each allocated block separately.
		 */
		eg.ConstantTable().FilterReserve(func(_ string, c *ZendConstant) bool {
			return c.IsPersistent()
		})
		eg.FunctionTable().FilterReserve(func(_ string, f types.IFunction) bool {
			return f.GetType() == ZEND_INTERNAL_FUNCTION
		})
		eg.ClassTable().FilterReserve(func(_ string, ce *types.ClassEntry) bool {
			return ce.IsInternalClass()
		})
		ZendCleanupInternalClasses()
	} else {
		eg.VmStack().Reset()
		if eg.GetFullTablesCleanup() {
			eg.ConstantTable().FilterReserve(func(_ string, c *ZendConstant) bool {
				return c.IsPersistent()
			})
			eg.FunctionTable().FilterReserve(func(_ string, f types.IFunction) bool {
				return f.GetType() == ZEND_INTERNAL_FUNCTION
			})
			eg.ClassTable().FilterReserve(func(_ string, ce *types.ClassEntry) bool {
				return ce.IsInternalClass()
			})
		} else {
			eg.ConstantTable().FilterReserve(func(_ string, c *ZendConstant) bool {
				if c.IsPersistent() {
					return true
				}

				// ZvalPtrDtorNogc(c.Value())
				return false
			})

			eg.FunctionTable().FilterReserve(func(key string, f types.IFunction) bool {
				if f.GetType() == ZEND_INTERNAL_FUNCTION {
					return true
				}
				return false
			})

			eg.ClassTable().FilterReserve(func(_ string, ce *types.ClassEntry) bool {
				if ce.IsInternalClass() {
					return true
				}

				//DestroyZendClass(zv)
				return false
			})
		}
		for eg.symtableCacheIdx > 0 {
			eg.symtableCacheIdx--
			eg.symtableCache[eg.symtableCacheIdx].Destroy()
		}
		eg.GetIncludedFiles().Destroy()
		eg.GetUserErrorHandlersErrorReporting().Destroy()
		eg.GetUserErrorHandlers().Destroy()
		eg.GetUserExceptionHandlers().Destroy()
		//eg.GetObjectsStore().Destroy()
		if eg.GetInAutoload() != nil {
			eg.GetInAutoload().Destroy()
		}
	}
	eg.ResetArrayIterators()
	ZendShutdownFpu()
}

// symtable cache
func (eg *ZendExecutorGlobals) PopSymbolTable(cap int) (symbolTable *types.Array) {
	if eg.symtableCacheIdx > 0 {
		eg.symtableCacheIdx--
		symbolTable = eg.symtableCache[eg.symtableCacheIdx]
	} else {
		symbolTable = types.NewArrayCap(cap)
	}
	return symbolTable
}

func (eg *ZendExecutorGlobals) CleanAndCacheSymbolTable(symbolTable *types.Array) {
	/* Clean before putting into the cache, since clean could call dtors,
	 * which could use the cached hash. Also do this before the check for
	 * available cache slots, as those may be used by a dtor as well. */
	symbolTable.Clean()
	if eg.symtableCacheIdx >= len(eg.symtableCache) {
		symbolTable.Destroy()
	} else {
		eg.symtableCache[eg.symtableCacheIdx] = symbolTable
		eg.symtableCacheIdx++
	}
}

/**
 * 辅助方法
 */
func (eg *ZendExecutorGlobals) ArrayIterators() []*types.ArrayIterator {
	return eg.arrayIterators
}
func (eg *ZendExecutorGlobals) ResetArrayIterators() {
	eg.arrayIterators = nil
}
func (eg *ZendExecutorGlobals) AddArrayIterator(ht *types.Array) uint32 {
	eg.arrayIterators = append(eg.arrayIterators, ht.Iterator())
	return uint32(len(eg.arrayIterators) - 1)
}
func (eg *ZendExecutorGlobals) GetArrayIterator(idx uint32) *types.ArrayIterator {
	len_ := uint32(len(eg.arrayIterators))
	if idx >= len_ {
		return nil
	}
	return eg.arrayIterators[idx]
}
func (eg *ZendExecutorGlobals) SetArrayIterator(idx uint32, iterator *types.ArrayIterator) {
	len_ := uint32(len(eg.arrayIterators))
	for len_ <= idx {
		eg.arrayIterators = append(eg.arrayIterators, nil)
		len_++
	}
	eg.arrayIterators[idx] = iterator
}
func (eg *ZendExecutorGlobals) DelArrayIterator(idx uint32) {
	len_ := uint32(len(eg.arrayIterators))
	if idx >= len_ {
		return
	}
	eg.arrayIterators[idx] = nil
	// tail
	if idx == len_-1 {
		for idx > 0 && eg.arrayIterators[idx-1] == nil {
			idx--
		}
		eg.arrayIterators = eg.arrayIterators[:idx]
	}
}

func (eg *ZendExecutorGlobals) GetHtIteratorsUsed() uint32           { return eg.ht_iterators_used }
func (eg *ZendExecutorGlobals) GetHtIterators() *types.ArrayIterator { return eg.ht_iterators }

func (eg *ZendExecutorGlobals) ClassTable() ClassTable       { return eg.classTable }
func (eg *ZendExecutorGlobals) FunctionTable() FunctionTable { return eg.functionTable }
func (eg *ZendExecutorGlobals) ConstantTable() ConstantTable { return eg.constantTable }

func (eg *ZendExecutorGlobals) IniDirectives() IniDirectives {
	return eg.iniDirectives
}
func (eg *ZendExecutorGlobals) InitIniDirectives() {
	eg.iniDirectives = types.NewTable[*ZendIniEntry](nil)
}
func (eg *ZendExecutorGlobals) ModifiedIniDirectives() IniDirectives {
	return eg.modifiedIniDirectives
}
func (eg *ZendExecutorGlobals) InitModifiedIniDirectives() {
	eg.modifiedIniDirectives = types.NewTable[*ZendIniEntry](nil)
}

func (eg *ZendExecutorGlobals) VmStack() *VmStack { return &eg.vmStack }

// llist
func (eg *ZendExecutorGlobals) GetRegularList() *types.Array { return &eg.regular_list }

func (eg *ZendExecutorGlobals) InitRegularList() {
	eg.persistentList = types.NewTable[*types.Resource](ListEntryDtor)
}
func (eg *ZendExecutorGlobals) RegularList() ResourceTable { return eg.regularList }

func (eg *ZendExecutorGlobals) InitPersistentList() {
	eg.persistentList = types.NewTable[*types.Resource](PlistEntryDtor)
}
func (eg *ZendExecutorGlobals) PersistentList() ResourceTable {
	return eg.persistentList
}

func (eg *ZendExecutorGlobals) IsActive() bool { return eg.active }

/**
 * 以下是自动生成的方法
 */

func (eg *ZendExecutorGlobals) GetErrorZval() *types.Zval        { return &eg.errorZval }
func (eg *ZendExecutorGlobals) GetSymtableCache() []*types.Array { return eg.symtableCache }
func (eg *ZendExecutorGlobals) GetSymtableCachePtr() **types.Array {
	return eg.symtable_cache_ptr
}
func (eg *ZendExecutorGlobals) GetSymbolTable() *types.Array        { return eg.symbolTable }
func (eg *ZendExecutorGlobals) SetSymbolTable(value *types.Array)   { eg.symbolTable = value }
func (eg *ZendExecutorGlobals) GetIncludedFiles() *types.Array      { return eg.includedFiles }
func (eg *ZendExecutorGlobals) SetIncludedFiles(value *types.Array) { eg.includedFiles = value }
func (eg *ZendExecutorGlobals) GetErrorReporting() int              { return eg.errorReporting }
func (eg *ZendExecutorGlobals) SetErrorReporting(value int)         { eg.errorReporting = value }
func (eg *ZendExecutorGlobals) GetExitStatus() int                  { return eg.exitStatus }
func (eg *ZendExecutorGlobals) SetExitStatus(value int)             { eg.exitStatus = value }
func (eg *ZendExecutorGlobals) GetCurrentExecuteData() *ZendExecuteData {
	return eg.current_execute_data
}
func (eg *ZendExecutorGlobals) SetCurrentExecuteData(value *ZendExecuteData) {
	eg.current_execute_data = value
}
func (eg *ZendExecutorGlobals) GetFakeScope() *types.ClassEntry      { return eg.fake_scope }
func (eg *ZendExecutorGlobals) SetFakeScope(value *types.ClassEntry) { eg.fake_scope = value }
func (eg *ZendExecutorGlobals) GetPrecision() ZendLong               { return eg.precision }
func (eg *ZendExecutorGlobals) SetPrecision(value ZendLong)          { eg.precision = value }
func (eg *ZendExecutorGlobals) GetPersistentConstantsCount() uint32 {
	return eg.persistent_constants_count
}
func (eg *ZendExecutorGlobals) SetPersistentConstantsCount(value uint32) {
	eg.persistent_constants_count = value
}
func (eg *ZendExecutorGlobals) GetPersistentFunctionsCount() uint32 {
	return eg.persistent_functions_count
}
func (eg *ZendExecutorGlobals) SetPersistentFunctionsCount(value uint32) {
	eg.persistent_functions_count = value
}
func (eg *ZendExecutorGlobals) GetPersistentClassesCount() uint32 {
	return eg.persistent_classes_count
}
func (eg *ZendExecutorGlobals) SetPersistentClassesCount(value uint32) {
	eg.persistent_classes_count = value
}
func (eg *ZendExecutorGlobals) GetInAutoload() *types.Array      { return eg.in_autoload }
func (eg *ZendExecutorGlobals) SetInAutoload(value *types.Array) { eg.in_autoload = value }
func (eg *ZendExecutorGlobals) GetAutoloadFunc() types.IFunction { return eg.autoload_func }
func (eg *ZendExecutorGlobals) SetAutoloadFunc(value types.IFunction) {
	eg.autoload_func = value
}
func (eg *ZendExecutorGlobals) GetFullTablesCleanup() bool    { return 0 }
func (eg *ZendExecutorGlobals) GetNoExtensions() bool         { return eg.no_extensions }
func (eg *ZendExecutorGlobals) SetNoExtensions(value bool)    { eg.no_extensions = value }
func (eg *ZendExecutorGlobals) GetVmInterrupt() bool          { return eg.vm_interrupt }
func (eg *ZendExecutorGlobals) SetVmInterrupt(value bool)     { eg.vm_interrupt = value }
func (eg *ZendExecutorGlobals) GetTimedOut() bool             { return eg.timed_out }
func (eg *ZendExecutorGlobals) SetTimedOut(value bool)        { eg.timed_out = value }
func (eg *ZendExecutorGlobals) GetHardTimeout() ZendLong      { return eg.hard_timeout }
func (eg *ZendExecutorGlobals) SetHardTimeout(value ZendLong) { eg.hard_timeout = value }
func (eg *ZendExecutorGlobals) GetUserErrorHandlerErrorReporting() int {
	return eg.user_error_handler_error_reporting
}
func (eg *ZendExecutorGlobals) SetUserErrorHandlerErrorReporting(value int) {
	eg.user_error_handler_error_reporting = value
}
func (eg *ZendExecutorGlobals) GetUserErrorHandler() *types.Zval { return &eg.user_error_handler }
func (eg *ZendExecutorGlobals) SetUserErrorHandler(value types.Zval) {
	eg.user_error_handler = value
}
func (eg *ZendExecutorGlobals) GetUserExceptionHandler() *types.Zval {
	return eg.user_exception_handler
}
func (eg *ZendExecutorGlobals) SetUserExceptionHandler(value *types.Zval) {
	eg.user_exception_handler = value
}
func (eg *ZendExecutorGlobals) GetUserErrorHandlersErrorReporting() ZendStack {
	return eg.user_error_handlers_error_reporting
}
func (eg *ZendExecutorGlobals) SetUserErrorHandlersErrorReporting(value ZendStack) {
	eg.user_error_handlers_error_reporting = value
}
func (eg *ZendExecutorGlobals) GetUserErrorHandlers() *ZendStack { return eg.user_error_handlers }
func (eg *ZendExecutorGlobals) GetUserExceptionHandlers() *ZendStack {
	return eg.user_exception_handlers
}
func (eg *ZendExecutorGlobals) GetErrorHandling() ZendErrorHandlingT { return eg.error_handling }
func (eg *ZendExecutorGlobals) SetErrorHandling(value ZendErrorHandlingT) {
	eg.error_handling = value
}
func (eg *ZendExecutorGlobals) GetExceptionClass() *types.ClassEntry { return eg.exception_class }
func (eg *ZendExecutorGlobals) SetExceptionClass(value *types.ClassEntry) {
	eg.exception_class = value
}
func (eg *ZendExecutorGlobals) GetTimeoutSeconds() ZendLong      { return eg.timeout_seconds }
func (eg *ZendExecutorGlobals) SetTimeoutSeconds(value ZendLong) { eg.timeout_seconds = value }
func (eg *ZendExecutorGlobals) GetLambdaCount() int              { return eg.lambda_count }
func (eg *ZendExecutorGlobals) SetLambdaCount(value int)         { eg.lambda_count = value }

func (eg *ZendExecutorGlobals) GetErrorReportingIniEntry() *ZendIniEntry {
	return eg.error_reporting_ini_entry
}
func (eg *ZendExecutorGlobals) SetErrorReportingIniEntry(value *ZendIniEntry) {
	eg.error_reporting_ini_entry = value
}
func (eg *ZendExecutorGlobals) GetException() *types.Object      { return eg.exception }
func (eg *ZendExecutorGlobals) SetException(value *types.Object) { eg.exception = value }
func (eg *ZendExecutorGlobals) GetPrevException() **types.Object { return eg.prev_exception }
func (eg *ZendExecutorGlobals) SetPrevException(value **types.Object) {
	eg.prev_exception = value
}
func (eg *ZendExecutorGlobals) GetOplineBeforeException() *types.ZendOp {
	return eg.opline_before_exception
}
func (eg *ZendExecutorGlobals) SetOplineBeforeException(value *types.ZendOp) {
	eg.opline_before_exception = value
}
func (eg *ZendExecutorGlobals) GetExceptionOp() *[3]types.ZendOp { return &eg.exception_op }
func (eg *ZendExecutorGlobals) GetCurrentModule() *ModuleEntry   { return eg.current_module }
func (eg *ZendExecutorGlobals) SetCurrentModule(value *ModuleEntry) {
	eg.current_module = value
}
func (eg *ZendExecutorGlobals) GetFlags() uint8                      { return eg.flags }
func (eg *ZendExecutorGlobals) SetFlags(value uint8)                 { eg.flags = value }
func (eg *ZendExecutorGlobals) GetAssertions() ZendLong              { return eg.assertions }
func (eg *ZendExecutorGlobals) SetAssertions(value ZendLong)         { eg.assertions = value }
func (eg *ZendExecutorGlobals) GetSavedFpuCwPtr() any                { return eg.saved_fpu_cw_ptr }
func (eg *ZendExecutorGlobals) SetSavedFpuCwPtr(value any)           { eg.saved_fpu_cw_ptr = value }
func (eg *ZendExecutorGlobals) GetTrampoline() types.IFunction       { return eg.trampoline }
func (eg *ZendExecutorGlobals) SetTrampoline(value types.IFunction)  { eg.trampoline = value }
func (eg *ZendExecutorGlobals) GetCallTrampolineOp() *types.ZendOp   { return eg.callTrampolineOp }
func (eg *ZendExecutorGlobals) SetCallTrampolineOp(op *types.ZendOp) { eg.callTrampolineOp = op }
func (eg *ZendExecutorGlobals) GetEachDeprecationThrown() bool {
	return eg.each_deprecation_thrown
}
func (eg *ZendExecutorGlobals) SetEachDeprecationThrown(value bool) {
	eg.each_deprecation_thrown = value
}
func (eg *ZendExecutorGlobals) GetWeakrefs() types.Array      { return eg.weakrefs }
func (eg *ZendExecutorGlobals) SetWeakrefs(value types.Array) { eg.weakrefs = value }
func (eg *ZendExecutorGlobals) GetExceptionIgnoreArgs() bool {
	return eg.exception_ignore_args
}
func (eg *ZendExecutorGlobals) SetExceptionIgnoreArgs(value bool) {
	eg.exception_ignore_args = value
}
func (eg *ZendExecutorGlobals) GetReserved() []any      { return eg.reserved }
func (eg *ZendExecutorGlobals) SetReserved(value []any) { eg.reserved = value }

/* ZendExecutorGlobals.flags */
func (eg *ZendExecutorGlobals) AddFlags(value uint8)      { eg.flags |= value }
func (eg *ZendExecutorGlobals) SubFlags(value uint8)      { eg.flags &^= value }
func (eg *ZendExecutorGlobals) HasFlags(value uint8) bool { return eg.flags&value != 0 }
func (eg *ZendExecutorGlobals) SwitchFlags(value uint8, cond bool) {
	if cond {
		eg.AddFlags(value)
	} else {
		eg.SubFlags(value)
	}
}
func (eg ZendExecutorGlobals) IsObjectStoreNoReuse() bool {
	return eg.HasFlags(EG_FLAGS_OBJECT_STORE_NO_REUSE)
}
func (eg ZendExecutorGlobals) IsInResourceShutdown() bool {
	return eg.HasFlags(EG_FLAGS_IN_RESOURCE_SHUTDOWN)
}
func (eg *ZendExecutorGlobals) SetIsObjectStoreNoReuse(cond bool) {
	eg.SwitchFlags(EG_FLAGS_OBJECT_STORE_NO_REUSE, cond)
}
func (eg *ZendExecutorGlobals) SetIsInResourceShutdown(cond bool) {
	eg.SwitchFlags(EG_FLAGS_IN_RESOURCE_SHUTDOWN, cond)
}
