// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
	"strings"
)

func ZendUnregisterFunctions(functions []types.ZendFunctionEntry, count int, functionTable *types.Array) {
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
func ModuleDestructor(module *ZendModuleEntry) {
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
	var p **ZendModuleEntry = ModuleRequestStartupHandlers
	for (*p) != nil {
		var module *ZendModuleEntry = *p
		if module.GetRequestStartupFunc()(module.GetType(), module.GetModuleNumber()) == types.FAILURE {
			faults.Error(faults.E_WARNING, "request_startup() for %s module failed", module.GetName())
			exit(1)
		}
		p++
	}
}
func ZendDeactivateModules() {
	EG__().SetCurrentExecuteData(nil)
	var __orig_bailout *JMP_BUF = EG__().GetBailout()
	var __bailout JMP_BUF
	EG__().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		if EG__().GetFullTablesCleanup() != 0 {
			var module *ZendModuleEntry
			var __ht *types.Array = &ModuleRegistry
			for _, _p := range __ht.ForeachDataReserve() {
				var _z types.Zval = _p.GetVal()

				module = _z.GetPtr()
				if module.GetRequestShutdownFunc() != nil {
					module.GetRequestShutdownFunc()(module.GetType(), module.GetModuleNumber())
				}
			}
		} else {
			var p **ZendModuleEntry = ModuleRequestShutdownHandlers
			for (*p) != nil {
				var module *ZendModuleEntry = *p
				module.GetRequestShutdownFunc()(module.GetType(), module.GetModuleNumber())
				p++
			}
		}
	}
	EG__().SetBailout(__orig_bailout)
}
func ZendCleanupInternalClasses() {
	var p **types.ClassEntry = ClassCleanupHandlers
	for (*p) != nil {
		ZendCleanupInternalClassData(*p)
		p++
	}
}
func ZendPostDeactivateModules() {
	if EG__().GetFullTablesCleanup() != 0 {
		var module *ZendModuleEntry
		var zv *types.Zval
		var key *types.String
		var __ht *types.Array = &ModuleRegistry
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			module = _z.GetPtr()
			if module.GetPostDeactivateFunc() != nil {
				module.GetPostDeactivateFunc()()
			}
		}
		var __ht__1 *types.Array = &ModuleRegistry
		for _, _p := range __ht__1.ForeachDataReserve() {
			var _z types.Zval = _p.GetVal()

			key = _p.GetKey()
			zv = _z
			module = zv.GetPtr()
			if module.GetType() != MODULE_TEMPORARY {
				break
			}
			ModuleDestructor(module)
			Free(module)
			types.ZendStringReleaseEx(key, 0)
			__ht__1.Len()--
			var j uint32 = types.HT_IDX_TO_HASH(_idx - 1)
			var nIndex uint32 = _p.GetH() | __ht__1.GetNTableMask()
			var i uint32 = types.HT_HASH(__ht__1, nIndex)
			if j != i {
				var prev *types.Bucket = __ht__1.Bucket(i)
				for prev.GetVal().GetNext() != j {
					i = prev.GetVal().GetNext()
					prev = __ht__1.Bucket(i)
				}
				prev.GetVal().GetNext() = _p.GetVal().GetNext()
			} else {
				types.HT_HASH(__ht__1, nIndex) = _p.GetVal().GetNext()
			}
		}
		__ht__1.SetNNumUsed(_idx)
	} else {
		var p **ZendModuleEntry = ModulePostDeactivateHandlers
		for (*p) != nil {
			var module *ZendModuleEntry = *p
			module.GetPostDeactivateFunc()()
			p++
		}
	}
}
func ZendNextFreeModule() int {
	return ModuleRegistry.Len() + 1
}
func DoRegisterInternalClass(orig_class_entry *types.ClassEntry, ce_flags uint32) *types.ClassEntry {
	var class_entry *types.ClassEntry = Malloc(b.SizeOf("zend_class_entry"))
	var lowercase_name *types.String
	*class_entry = *orig_class_entry
	class_entry.SetType(ZEND_INTERNAL_CLASS)
	ZendInitializeClassData(class_entry, 0)
	class_entry.SetCeFlags(ce_flags | ZEND_ACC_CONSTANTS_UPDATED | ZEND_ACC_LINKED | ZEND_ACC_RESOLVED_PARENT | ZEND_ACC_RESOLVED_INTERFACES)
	class_entry.SetModule(EG__().GetCurrentModule())
	if class_entry.GetBuiltinFunctions() != nil {
		ZendRegisterFunctions(class_entry, class_entry.GetBuiltinFunctions(), class_entry.GetFunctionTable(), EG__().GetCurrentModule().GetType())
	}
	lowercase_name = ZendStringTolowerEx(orig_class_entry.GetName(), EG__().GetCurrentModule().GetType() == MODULE_PERSISTENT)
	lowercase_name = types.ZendNewInternedString(lowercase_name)
	types.ZendHashUpdatePtr(CG__().GetClassTable(), lowercase_name.GetStr(), class_entry)
	types.ZendStringReleaseEx(lowercase_name, 1)
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
	return DoRegisterInternalClass(orig_class_entry, ZEND_ACC_INTERFACE)
}
func ZendRegisterClassAliasEx(name *byte, name_len int, ce *types.ClassEntry, persistent int) int {
	var lcname *types.String
	var zv types.Zval
	var ret *types.Zval

	/* TODO: Move this out of here in 7.4. */

	if persistent != 0 && EG__().GetCurrentModule() != nil && EG__().GetCurrentModule().GetType() == MODULE_TEMPORARY {
		persistent = 0
	}
	if name[0] == '\\' {
		lcname = types.ZendStringAlloc(name_len-1, persistent)
		ZendStrTolowerCopy(lcname.GetVal(), name+1, name_len-1)
	} else {
		lcname = types.ZendStringAlloc(name_len, persistent)
		ZendStrTolowerCopy(lcname.GetVal(), name, name_len)
	}
	ZendAssertValidClassName(lcname)
	lcname = types.ZendNewInternedString(lcname)
	types.ZVAL_ALIAS_PTR(&zv, ce)
	ret = CG__().GetClassTable().KeyAdd(lcname.GetStr(), &zv)
	types.ZendStringReleaseEx(lcname, 0)
	if ret != nil {
		if !ce.IsImmutable() {
			ce.GetRefcount()++
		}
		return types.SUCCESS
	}
	return types.FAILURE
}
func ZendSetHashSymbol(
	symbol *types.Zval,
	name *byte,
	name_length int,
	is_ref types.ZendBool,
	num_symbol_tables int,
	_ ...any,
) int {
	var symbol_table *types.Array
	var symbol_table_list va_list
	if num_symbol_tables <= 0 {
		return types.FAILURE
	}
	if is_ref != 0 {
		types.ZVAL_MAKE_REF(symbol)
	}
	va_start(symbol_table_list, num_symbol_tables)
	for b.PostDec(&num_symbol_tables) > 0 {
		symbol_table = __va_arg(symbol_table_list, (*types.Array)(_))
		symbol_table.KeyUpdate(b.CastStr(name, name_length), symbol)
		symbol.TryAddRefcount()
	}
	va_end(symbol_table_list)
	return types.SUCCESS
}
func ZifDisplayDisabledFunction(executeData *ZendExecuteData, return_value *types.Zval) {
	faults.Error(faults.E_WARNING, "%s() has been disabled for security reasons", GetActiveFunctionName())
}
func ZendDisableFunction(function_name *byte, function_name_length int) int {
	var func_ *ZendInternalFunction
	if b.Assign(&func_, types.ZendHashStrFindPtr(CG__().GetFunctionTable(), b.CastStr(function_name, function_name_length))) {
		ZendFreeInternalArgInfo(func_)
		func_.SubFnFlags(ZEND_ACC_VARIADIC | ZEND_ACC_HAS_TYPE_HINTS | ZEND_ACC_HAS_RETURN_TYPE)
		func_.SetNumArgs(0)
		func_.SetArgInfo(nil)
		func_.SetHandler(ZifDisplayDisabledFunction)
		return types.SUCCESS
	}
	return types.FAILURE
}
