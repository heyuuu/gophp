package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/globals"
	"github.com/heyuuu/gophp/zend/types"
	"strings"
)

func ZendUnregisterFunctions(functions []types.FunctionEntry, count int, functionTable *types.Array) {
	targetFunctionTable := functionTable
	if targetFunctionTable == nil {
		targetFunctionTable = CG__().GetFunctionTable()
	}

	for i, ptr := range functions {
		// count 为 -1 不限制；否则限制方法个数
		if count == -1 || i >= count {
			break
		}
		lcName := strings.ToLower(ptr.FuncName())
		functionTable.KeyDelete(lcName)
	}
}

func CleanModuleClass(el *types.Zval, arg any) int {
	var ce *types.ClassEntry = (*types.ClassEntry)(el.GetPtr())
	var module_number int = *((*int)(arg))
	if ce.GetType() == ZEND_INTERNAL_CLASS && ce.GetModule().GetModuleNumber() == module_number {
		return types.ArrayApplyRemove
	} else {
		return types.ArrayApplyKeep
	}
}
func CleanModuleClasses(module_number int) {
	types.ZendHashApplyWithArgument(EG__().GetClassTable(), CleanModuleClass, any(&module_number))
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
		ZendCleanupInternalClassData(*p)
		p++
	}
}
func ZendPostDeactivateModules() {
	// deleted
}
func ZendNextFreeModule() int {
	return globals.G().CountModules() + 1
}
func DoRegisterInternalClass(orig_class_entry *types.ClassEntry, ce_flags uint32) *types.ClassEntry {
	var class_entry = &types.ClassEntry{}
	*class_entry = *orig_class_entry
	class_entry.SetType(ZEND_INTERNAL_CLASS)
	ZendInitializeClassData(class_entry, 0)
	class_entry.SetCeFlags(ce_flags | AccConstantsUpdated | AccLinked | AccResolvedParent | AccResolvedInterfaces)
	class_entry.SetModule(EG__().GetCurrentModule())
	if class_entry.GetBuiltinFunctions() != nil {
		ZendRegisterFunctions(class_entry, class_entry.GetBuiltinFunctions(), class_entry.GetFunctionTable(), EG__().GetCurrentModule().GetType())
	}
	CG__().ClassTable().Update(class_entry.GetName().GetStr(), class_entry)
	return class_entry
}
func ZendRegisterInternalClassEx(class_entry *types.ClassEntry, parent_ce *types.ClassEntry) *types.ClassEntry {
	var register_class *types.ClassEntry
	register_class = ZendRegisterInternalClass(class_entry)
	if parent_ce != nil {
		ZendDoInheritance(register_class, parent_ce)
		ZendBuildPropertiesInfoTable(register_class)
	}
	return register_class
}
func ZendClassImplements(class_entry *types.ClassEntry, num_interfaces int, _ ...any) {
	var interface_entry *types.ClassEntry
	var interface_list va_list
	va_start(interface_list, num_interfaces)
	for b.PostDec(&num_interfaces) {
		interface_entry = __va_arg(interface_list, (*types.ClassEntry)(_))
		ZendDoImplementInterface(class_entry, interface_entry)
	}
	va_end(interface_list)
}
func ZendRegisterInternalClass(orig_class_entry *types.ClassEntry) *types.ClassEntry {
	return DoRegisterInternalClass(orig_class_entry, 0)
}
func ZendRegisterInternalInterface(orig_class_entry *types.ClassEntry) *types.ClassEntry {
	return DoRegisterInternalClass(orig_class_entry, AccInterface)
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
		if !ce.IsImmutable() {
			ce.GetRefcount()++
		}
		return types.SUCCESS
	}
	return types.FAILURE
}
func ZifDisplayDisabledFunction(executeData *ZendExecuteData, return_value *types.Zval) {
	faults.Error(faults.E_WARNING, "%s() has been disabled for security reasons", GetActiveFunctionName())
}
func ZendDisableFunction(function_name *byte, function_name_length int) int {
	var func_ *types.InternalFunction
	if b.Assign(&func_, types.ZendHashStrFindPtr(CG__().GetFunctionTable(), b.CastStr(function_name, function_name_length))) {
		ZendFreeInternalArgInfo(func_)
		func_.SubFnFlags(AccVariadic | AccHasTypeHints | AccHasReturnType)
		func_.SetNumArgs(0)
		func_.SetArgInfo(nil)
		func_.SetHandler(ZifDisplayDisabledFunction)
		return types.SUCCESS
	}
	return types.FAILURE
}
