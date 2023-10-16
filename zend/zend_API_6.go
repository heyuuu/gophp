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

func ModuleDestructor(module *ModuleEntry) {
	module.ModuleShutdown()
}
func ZendActivateModules() {
	globals.G().EachModuleReserve(func(module *ModuleEntry) {
		if !module.RequestStartup() {
			faults.Error(faults.E_WARNING, "request_startup() for %s module failed", module.Name())
			exit(1)
		}
	})
}
func ZendDeactivateModules() {
	EG__().SetCurrentExecuteData(nil)
	faults.Try(func() {
		globals.G().EachModuleReserve(func(module *ModuleEntry) {
			module.RequestShutdown()
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
func ZendRegisterClassAliasEx(name string, ce *types.ClassEntry) bool {
	/* TODO: Move this out of here in 7.4. */
	name = trimClassName(name)
	ZendAssertValidClassName(name)
	return CG__().ClassTable().Add(name, ce)
}
func ZifDisplayDisabledFunction(executeData *ZendExecuteData, return_value *types.Zval) {
	faults.Error(faults.E_WARNING, "%s() has been disabled for security reasons", CurrEX().FunctionName())
}
func ZendDisableFunction(functionName string) int {
	f := CG__().FunctionTable().Get(functionName)
	if f != nil {
		func_ := f.(*types.InternalFunction)
		//ZendFreeInternalArgInfo(func_)
		func_.SubFnFlags(types.AccVariadic | types.AccHasTypeHints | types.AccHasReturnType)
		func_.SetNumArgs(0)
		func_.SetArgInfo(nil)
		func_.SetHandler(ZifDisplayDisabledFunction)
		return types.SUCCESS
	}
	return types.FAILURE
}
