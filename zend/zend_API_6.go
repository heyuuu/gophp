package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/globals"
)

func ZendUnregisterFunctions(functions []types.FunctionEntry, count int, functionTable FunctionTable) {
	targetFunctionTable := functionTable
	if targetFunctionTable == nil {
		targetFunctionTable = CG__().FunctionTable()
	}

	for i, ptr := range functions {
		// count 为 -1 不限制；否则限制方法个数
		if count == -1 || i >= count {
			break
		}
		functionTable.Del(ptr.FuncName())
	}
}

func CleanModuleClasses(moduleNumber int) {
	EG__().ClassTable().Filter(func(_ string, ce *types.ClassEntry) bool {
		needClean := ce.IsInternalClass() && ce.ModuleNumber() == moduleNumber
		return !needClean
	})
}
func ModuleDestructor(module *ModuleEntry) {
	if module.GetType() == MODULE_TEMPORARY {
		ZendCleanModuleRsrcDtors(module.GetModuleNumber())
		CleanModuleConstants(module.GetModuleNumber())
		CleanModuleClasses(module.GetModuleNumber())
	}
	if module.GetModuleStarted() != 0 && module.GetModuleShutdownFunc() != nil {
		module.GetModuleShutdownFunc()(module.GetType(), module.GetModuleNumber())
	}
	if module.GetModuleStarted() != 0 && module.GetModuleShutdownFunc() == nil && module.GetType() == MODULE_TEMPORARY {
		ZendUnregisterIniEntries(module.GetModuleNumber())
	}

	/* Deinitilaise module globals */
	if module.GetGlobalsSize() != 0 {
		if module.GetGlobalsDtor() != nil {
			module.GetGlobalsDtor()(module.GetGlobalsPtr())
		}
	}
	module.SetModuleStarted(0)
	if module.GetType() == MODULE_TEMPORARY && module.GetFunctions() != nil {
		ZendUnregisterFunctions(module.GetFunctions(), -1, nil)
	}
	if module.GetHandle() && !(getenv("ZEND_DONT_UNLOAD_MODULES")) {
		DL_UNLOAD(module.GetHandle())
	}
}
func ZendActivateModules() {
	globals.G().EachModuleReserve(func(module *ModuleEntry) {
		if module.GetModuleStartupFunc() == nil {
			return
		}
		if module.GetRequestStartupFunc()(module.GetType(), module.GetModuleNumber()) == types.FAILURE {
			faults.Error(faults.E_WARNING, "request_startup() for %s module failed", module.GetName())
			exit(1)
		}
	})
}
func ZendDeactivateModules() {
	EG__().SetCurrentExecuteData(nil)
	faults.Try(func() {
		globals.G().EachModuleReserve(func(module *ModuleEntry) {
			if module.GetRequestShutdownFunc() != nil {
				module.GetRequestShutdownFunc()(module.GetType(), module.GetModuleNumber())
			}
		})
	})
}
func ZendCleanupInternalClasses() {
	var p **types.ClassEntry = ClassCleanupHandlers
	for (*p) != nil {
		//ZendCleanupInternalClassData(*p)
		p++
	}
}
func ZendNextFreeModule() int {
	return globals.G().CountModules() + 1
}
func ZendClassImplements(classEntry *types.ClassEntry, _ int, interfaces ...*types.ClassEntry) {
	for _, iface := range interfaces {
		ZendDoImplementInterface(classEntry, iface)
	}
}
func ZendRegisterClassAliasEx(name string, ce *types.ClassEntry, persistent int) int {
	/* TODO: Move this out of here in 7.4. */
	if persistent != 0 && EG__().GetCurrentModule() != nil && EG__().GetCurrentModule().GetType() == MODULE_TEMPORARY {
		persistent = 0
	}
	if name[0] == '\\' {
		name = name[1:]
	}
	ZendAssertValidClassName(name)
	if CG__().ClassTable().Add(name, ce) {
		return types.SUCCESS
	}
	return types.FAILURE
}
func ZifDisplayDisabledFunction(executeData *ZendExecuteData, return_value *types.Zval) {
	faults.Error(faults.E_WARNING, "%s() has been disabled for security reasons", GetActiveFunctionName())
}
func ZendDisableFunction(functionName string) int {
	f := CG__().FunctionTable().Get(functionName)
	if f != nil {
		func_ := f.(*types.InternalFunction)
		//ZendFreeInternalArgInfo(func_)
		func_.SubFnFlags(AccVariadic | AccHasTypeHints | AccHasReturnType)
		func_.SetNumArgs(0)
		func_.SetArgInfo(nil)
		func_.SetHandler(ZifDisplayDisabledFunction)
		return types.SUCCESS
	}
	return types.FAILURE
}
