package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/slicekit"
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
	vmStack          []*ZendExecuteData

	currentExecuteData *ZendExecuteData
	fakeScope          *types.ClassEntry
	precision          int
	inAutoload         *types.Array
	autoloadFunc       types.IFunction
	vmInterrupt        bool
	timedOut           bool
	hardTimeout        int

	regular_list    types.Array
	persistent_list types.Array
	regularList     ResourceTable
	persistentList  ResourceTable

	userErrorHandlerErrorReporting  int
	userErrorHandler                *types.Zval
	userExceptionHandler            *types.Zval
	userErrorHandlersErrorReporting []int
	userErrorHandlers               []*types.Zval
	userExceptionHandlers           []*types.Zval
	errorHandling                   ZendErrorHandlingT
	exceptionClass                  *types.ClassEntry
	timeoutSeconds                  int

	iniDirectives         IniDirectives
	modifiedIniDirectives IniDirectives

	errorReportingIniEntry *ZendIniEntry
	exception              *types.Object
	prevException          *types.Object
	oplineBeforeException  *types.ZendOp
	exceptionOp            [3]types.ZendOp
	currentModule          *ModuleEntry
	active                 bool
	flags                  uint8
	assertions             int

	htIteratorsCount uint32
	htIteratorsUsed  uint32
	htIterators      *types.ArrayIterator
	htIteratorsSlots []types.ArrayIterator
	arrayIterators   []*types.ArrayIterator

	trampoline            types.IFunction
	callTrampolineOp      *types.ZendOp
	eachDeprecationThrown bool
	exceptionIgnoreArgs   bool
	reserved              []any
}

/**
 * 生命周期
 */
var _ contracts.ModuleLifeCycle = (*ZendExecutorGlobals)(nil)

func (eg *ZendExecutorGlobals) StartUp() {
	eg.constantTable = types.NewTable[*ZendConstant]()
}
func (eg *ZendExecutorGlobals) Shutdown() {
	eg.constantTable.Destroy()
}
func (eg *ZendExecutorGlobals) Activate() {
	eg.errorZval.SetIsError()
	eg.symbolTable = types.NewArrayCap(64)
	eg.includedFiles = types.NewArrayCap(8)

	eg.functionTable = CG__().FunctionTable()
	eg.classTable = CG__().ClassTable()
	eg.vmStack = nil

	eg.inAutoload = nil
	eg.autoloadFunc = nil
	eg.errorHandling = EH_NORMAL
	eg.flags = EG_FLAGS_INITIAL
	eg.userErrorHandler = nil
	eg.userExceptionHandler = nil
	eg.currentExecuteData = nil
	eg.userErrorHandlersErrorReporting = nil
	eg.userErrorHandlers = nil
	eg.userExceptionHandlers = nil
	eg.vmInterrupt = false
	eg.timedOut = false
	eg.exception = nil
	eg.prevException = nil
	eg.fakeScope = nil
	eg.GetTrampoline().SetFunctionName("")
	eg.ResetArrayIterators()
	eg.eachDeprecationThrown = false

	eg.active = true
}
func (eg *ZendExecutorGlobals) Deactivate() {
	var fastShutdown bool = IsZendMm() != 0

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
		eg.userErrorHandlersErrorReporting = nil
		eg.userErrorHandlers = nil
		eg.userExceptionHandlers = nil
	}

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
		eg.vmStack = nil

		eg.ConstantTable().FilterReserve(func(_ string, c *ZendConstant) bool {
			return c.IsPersistent()
		})
		eg.FunctionTable().FilterReserve(func(_ string, f types.IFunction) bool {
			return f.GetType() == ZEND_INTERNAL_FUNCTION
		})
		eg.ClassTable().FilterReserve(func(_ string, ce *types.ClassEntry) bool {
			return ce.IsInternalClass()
		})

		for eg.symtableCacheIdx > 0 {
			eg.symtableCacheIdx--
			eg.symtableCache[eg.symtableCacheIdx].Destroy()
		}
		eg.GetIncludedFiles().Destroy()
		eg.userErrorHandlersErrorReporting = nil
		eg.userErrorHandlers = nil
		eg.userExceptionHandlers = nil
		//eg.GetObjectsStore().Destroy()
		if eg.GetInAutoload() != nil {
			eg.GetInAutoload().Destroy()
		}
	}
	eg.ResetArrayIterators()
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

func (eg *ZendExecutorGlobals) GetHtIteratorsUsed() uint32           { return eg.htIteratorsUsed }
func (eg *ZendExecutorGlobals) GetHtIterators() *types.ArrayIterator { return eg.htIterators }

func (eg *ZendExecutorGlobals) ClassTable() ClassTable       { return eg.classTable }
func (eg *ZendExecutorGlobals) FunctionTable() FunctionTable { return eg.functionTable }
func (eg *ZendExecutorGlobals) ConstantTable() ConstantTable { return eg.constantTable }

func (eg *ZendExecutorGlobals) IniDirectives() IniDirectives {
	return eg.iniDirectives
}
func (eg *ZendExecutorGlobals) InitIniDirectives() {
	eg.iniDirectives = types.NewTable[*ZendIniEntry]()
}
func (eg *ZendExecutorGlobals) ModifiedIniDirectives() IniDirectives {
	return eg.modifiedIniDirectives
}
func (eg *ZendExecutorGlobals) InitModifiedIniDirectives() {
	eg.modifiedIniDirectives = types.NewTable[*ZendIniEntry]()
}

func (eg *ZendExecutorGlobals) VmStackPush(ex *ZendExecuteData) {
	eg.vmStack = append(eg.vmStack, ex)
}
func (eg *ZendExecutorGlobals) VmStackPop() *ZendExecuteData {
	ex, _ := slicekit.Pop(&eg.vmStack)
	return ex
}
func (eg *ZendExecutorGlobals) VmStackPopCheck(ex *ZendExecuteData) {
	b.Assert(eg.VmStackPop() == ex)
}

// llist
func (eg *ZendExecutorGlobals) GetRegularList() *types.Array { return &eg.regular_list }

func (eg *ZendExecutorGlobals) InitRegularList() {
	eg.persistentList = types.NewTableEx[*types.Resource](ListEntryDtor)
}
func (eg *ZendExecutorGlobals) RegularList() ResourceTable { return eg.regularList }

func (eg *ZendExecutorGlobals) InitPersistentList() {
	eg.persistentList = types.NewTableEx[*types.Resource](PlistEntryDtor)
}
func (eg *ZendExecutorGlobals) PersistentList() ResourceTable {
	return eg.persistentList
}

func (eg *ZendExecutorGlobals) IsActive() bool { return eg.active }

// error && exception
func (eg *ZendExecutorGlobals) SetUserErrorHandler(errorHandler *types.Zval, errorLevels int) {
	slicekit.Push(&eg.userErrorHandlersErrorReporting, eg.userErrorHandlerErrorReporting)
	slicekit.Push(&eg.userErrorHandlers, eg.userErrorHandler)

	eg.userErrorHandler = nil
	if !errorHandler.IsNull() {
		eg.userErrorHandler = errorHandler.Copy()
		eg.userErrorHandlerErrorReporting = errorLevels
	}
}
func (eg *ZendExecutorGlobals) RestoreUserErrorHandler() {
	eg.userErrorHandler = nil
	if len(eg.userErrorHandlers) > 0 {
		eg.userErrorHandlerErrorReporting, _ = slicekit.Pop(&eg.userErrorHandlersErrorReporting)
		eg.userErrorHandler, _ = slicekit.Pop(&eg.userErrorHandlers)
	}
}

func (eg *ZendExecutorGlobals) SetUserExceptionHandler(exceptionHandler *types.Zval) {
	slicekit.Push(&eg.userExceptionHandlers, eg.userExceptionHandler)

	eg.userExceptionHandler = nil
	if !exceptionHandler.IsNull() {
		eg.userExceptionHandler = exceptionHandler.Copy()
	}
}
func (eg *ZendExecutorGlobals) RestoreUserExceptionHandler() {
	eg.userExceptionHandler = nil
	if len(eg.userExceptionHandlers) > 0 {
		eg.userExceptionHandler, _ = slicekit.Pop(&eg.userExceptionHandlers)
	}
}

/**
 * 以下是自动生成的方法
 */
func (eg *ZendExecutorGlobals) GetErrorZval() *types.Zval           { return &eg.errorZval }
func (eg *ZendExecutorGlobals) GetSymbolTable() *types.Array        { return eg.symbolTable }
func (eg *ZendExecutorGlobals) SetSymbolTable(value *types.Array)   { eg.symbolTable = value }
func (eg *ZendExecutorGlobals) GetIncludedFiles() *types.Array      { return eg.includedFiles }
func (eg *ZendExecutorGlobals) SetIncludedFiles(value *types.Array) { eg.includedFiles = value }
func (eg *ZendExecutorGlobals) GetErrorReporting() int              { return eg.errorReporting }
func (eg *ZendExecutorGlobals) SetErrorReporting(value int)         { eg.errorReporting = value }
func (eg *ZendExecutorGlobals) GetExitStatus() int                  { return eg.exitStatus }
func (eg *ZendExecutorGlobals) SetExitStatus(value int)             { eg.exitStatus = value }
func (eg *ZendExecutorGlobals) GetCurrentExecuteData() *ZendExecuteData {
	return eg.currentExecuteData
}
func (eg *ZendExecutorGlobals) SetCurrentExecuteData(value *ZendExecuteData) {
	eg.currentExecuteData = value
}
func (eg *ZendExecutorGlobals) GetFakeScope() *types.ClassEntry      { return eg.fakeScope }
func (eg *ZendExecutorGlobals) SetFakeScope(value *types.ClassEntry) { eg.fakeScope = value }
func (eg *ZendExecutorGlobals) GetPrecision() int                    { return eg.precision }
func (eg *ZendExecutorGlobals) SetPrecision(value int)               { eg.precision = value }
func (eg *ZendExecutorGlobals) GetInAutoload() *types.Array          { return eg.inAutoload }
func (eg *ZendExecutorGlobals) SetInAutoload(value *types.Array)     { eg.inAutoload = value }
func (eg *ZendExecutorGlobals) GetAutoloadFunc() types.IFunction     { return eg.autoloadFunc }
func (eg *ZendExecutorGlobals) SetAutoloadFunc(value types.IFunction) {
	eg.autoloadFunc = value
}
func (eg *ZendExecutorGlobals) GetVmInterrupt() bool      { return eg.vmInterrupt }
func (eg *ZendExecutorGlobals) SetVmInterrupt(value bool) { eg.vmInterrupt = value }
func (eg *ZendExecutorGlobals) GetTimedOut() bool         { return eg.timedOut }
func (eg *ZendExecutorGlobals) SetTimedOut(value bool)    { eg.timedOut = value }
func (eg *ZendExecutorGlobals) GetHardTimeout() int       { return eg.hardTimeout }
func (eg *ZendExecutorGlobals) SetHardTimeout(value int)  { eg.hardTimeout = value }
func (eg *ZendExecutorGlobals) GetUserErrorHandlerErrorReporting() int {
	return eg.userErrorHandlerErrorReporting
}
func (eg *ZendExecutorGlobals) SetUserErrorHandlerErrorReporting(value int) {
	eg.userErrorHandlerErrorReporting = value
}
func (eg *ZendExecutorGlobals) GetUserErrorHandler() *types.Zval     { return eg.userErrorHandler }
func (eg *ZendExecutorGlobals) GetUserExceptionHandler() *types.Zval { return eg.userExceptionHandler }

func (eg *ZendExecutorGlobals) GetErrorHandling() ZendErrorHandlingT      { return eg.errorHandling }
func (eg *ZendExecutorGlobals) SetErrorHandling(value ZendErrorHandlingT) { eg.errorHandling = value }
func (eg *ZendExecutorGlobals) GetExceptionClass() *types.ClassEntry      { return eg.exceptionClass }
func (eg *ZendExecutorGlobals) SetExceptionClass(value *types.ClassEntry) { eg.exceptionClass = value }
func (eg *ZendExecutorGlobals) GetTimeoutSeconds() int                    { return eg.timeoutSeconds }
func (eg *ZendExecutorGlobals) SetTimeoutSeconds(value int)               { eg.timeoutSeconds = value }

func (eg *ZendExecutorGlobals) GetErrorReportingIniEntry() *ZendIniEntry {
	return eg.errorReportingIniEntry
}
func (eg *ZendExecutorGlobals) SetErrorReportingIniEntry(value *ZendIniEntry) {
	eg.errorReportingIniEntry = value
}
func (eg *ZendExecutorGlobals) GetOplineBeforeException() *types.ZendOp {
	return eg.oplineBeforeException
}
func (eg *ZendExecutorGlobals) SetOplineBeforeException(value *types.ZendOp) {
	eg.oplineBeforeException = value
}
func (eg *ZendExecutorGlobals) GetExceptionOp() *[3]types.ZendOp { return &eg.exceptionOp }
func (eg *ZendExecutorGlobals) GetCurrentModule() *ModuleEntry   { return eg.currentModule }
func (eg *ZendExecutorGlobals) SetCurrentModule(value *ModuleEntry) {
	eg.currentModule = value
}
func (eg *ZendExecutorGlobals) GetFlags() uint8                      { return eg.flags }
func (eg *ZendExecutorGlobals) SetFlags(value uint8)                 { eg.flags = value }
func (eg *ZendExecutorGlobals) GetAssertions() int                   { return eg.assertions }
func (eg *ZendExecutorGlobals) SetAssertions(value int)              { eg.assertions = value }
func (eg *ZendExecutorGlobals) GetTrampoline() types.IFunction       { return eg.trampoline }
func (eg *ZendExecutorGlobals) SetTrampoline(value types.IFunction)  { eg.trampoline = value }
func (eg *ZendExecutorGlobals) GetCallTrampolineOp() *types.ZendOp   { return eg.callTrampolineOp }
func (eg *ZendExecutorGlobals) SetCallTrampolineOp(op *types.ZendOp) { eg.callTrampolineOp = op }
func (eg *ZendExecutorGlobals) GetEachDeprecationThrown() bool       { return eg.eachDeprecationThrown }
func (eg *ZendExecutorGlobals) SetEachDeprecationThrown(value bool)  { eg.eachDeprecationThrown = value }
func (eg *ZendExecutorGlobals) GetExceptionIgnoreArgs() bool         { return eg.exceptionIgnoreArgs }
func (eg *ZendExecutorGlobals) SetExceptionIgnoreArgs(value bool)    { eg.exceptionIgnoreArgs = value }
func (eg *ZendExecutorGlobals) GetReserved() []any                   { return eg.reserved }
func (eg *ZendExecutorGlobals) SetReserved(value []any)              { eg.reserved = value }

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

/**
 * exceptions && errors
 */
func (eg *ZendExecutorGlobals) GetException() *types.Object      { return eg.exception }
func (eg *ZendExecutorGlobals) SetException(value *types.Object) { eg.exception = value }

func (eg *ZendExecutorGlobals) NoException() bool  { return eg.exception != nil }
func (eg *ZendExecutorGlobals) HasException() bool { return eg.exception != nil }
func (eg *ZendExecutorGlobals) ClearException() {
	eg.prevException = nil
	if eg.exception != nil {
		eg.exception = nil
		if eg.currentExecuteData != nil {
			eg.currentExecuteData.SetOpline(eg.oplineBeforeException)
		}
	}
}
func (eg *ZendExecutorGlobals) ExceptionSave() {
	if eg.prevException != nil {
		faults.ExceptionSetPrevious(eg.exception, eg.prevException)
	}
	if eg.exception != nil {
		eg.prevException = eg.exception
	}
	eg.exception = nil
}

func (eg *ZendExecutorGlobals) ExceptionRestore() {
	if eg.prevException != nil {
		if eg.exception != nil {
			faults.ExceptionSetPrevious(eg.exception, eg.prevException)
		} else {
			eg.exception = eg.prevException
		}
		eg.prevException = nil
	}
}
