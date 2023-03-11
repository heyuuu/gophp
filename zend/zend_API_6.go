// <<generate>>

package zend

import (
	b "sik/builtin"
	"strings"
)

func ZendUnregisterFunctions(functions []ZendFunctionEntry, count int, functionTable *HashTable) {
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

func CleanModuleClass(el *Zval, arg any) int {
	var ce *ZendClassEntry = (*ZendClassEntry)(el.GetPtr())
	var module_number int = *((*int)(arg))
	if ce.GetType() == ZEND_INTERNAL_CLASS && ce.GetModule().GetModuleNumber() == module_number {
		return ZEND_HASH_APPLY_REMOVE
	} else {
		return ZEND_HASH_APPLY_KEEP
	}
}
func CleanModuleClasses(module_number int) {
	ZendHashApplyWithArgument(EG__().GetClassTable(), CleanModuleClass, any(&module_number))
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
		if module.GetRequestStartupFunc()(module.GetType(), module.GetModuleNumber()) == FAILURE {
			ZendError(E_WARNING, "request_startup() for %s module failed", module.GetName())
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
			var __ht *HashTable = &ModuleRegistry
			for _, _p := range __ht.foreachDataReserve() {
				var _z Zval = _p.GetVal()

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
	var p **ZendClassEntry = ClassCleanupHandlers
	for (*p) != nil {
		ZendCleanupInternalClassData(*p)
		p++
	}
}
func ZendPostDeactivateModules() {
	if EG__().GetFullTablesCleanup() != 0 {
		var module *ZendModuleEntry
		var zv *Zval
		var key *ZendString
		var __ht *HashTable = &ModuleRegistry
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			module = _z.GetPtr()
			if module.GetPostDeactivateFunc() != nil {
				module.GetPostDeactivateFunc()()
			}
		}
		var __ht__1 *HashTable = &ModuleRegistry
		for _, _p := range __ht__1.foreachDataReserve() {
			var _z Zval = _p.GetVal()

			key = _p.GetKey()
			zv = _z
			module = zv.GetPtr()
			if module.GetType() != MODULE_TEMPORARY {
				break
			}
			ModuleDestructor(module)
			Free(module)
			ZendStringReleaseEx(key, 0)
			__ht__1.GetNNumOfElements()--
			var j uint32 = HT_IDX_TO_HASH(_idx - 1)
			var nIndex uint32 = _p.GetH() | __ht__1.GetNTableMask()
			var i uint32 = HT_HASH(__ht__1, nIndex)
			if j != i {
				var prev *Bucket = __ht__1.Bucket(i)
				for prev.GetVal().GetNext() != j {
					i = prev.GetVal().GetNext()
					prev = __ht__1.Bucket(i)
				}
				prev.GetVal().GetNext() = _p.GetVal().GetNext()
			} else {
				HT_HASH(__ht__1, nIndex) = _p.GetVal().GetNext()
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
	return ModuleRegistry.GetNNumOfElements() + 1
}
func DoRegisterInternalClass(orig_class_entry *ZendClassEntry, ce_flags uint32) *ZendClassEntry {
	var class_entry *ZendClassEntry = Malloc(b.SizeOf("zend_class_entry"))
	var lowercase_name *ZendString
	*class_entry = *orig_class_entry
	class_entry.SetType(ZEND_INTERNAL_CLASS)
	ZendInitializeClassData(class_entry, 0)
	class_entry.SetCeFlags(ce_flags | ZEND_ACC_CONSTANTS_UPDATED | ZEND_ACC_LINKED | ZEND_ACC_RESOLVED_PARENT | ZEND_ACC_RESOLVED_INTERFACES)
	class_entry.SetModule(EG__().GetCurrentModule())
	if class_entry.GetBuiltinFunctions() != nil {
		ZendRegisterFunctions(class_entry, class_entry.GetBuiltinFunctions(), class_entry.GetFunctionTable(), EG__().GetCurrentModule().GetType())
	}
	lowercase_name = ZendStringTolowerEx(orig_class_entry.GetName(), EG__().GetCurrentModule().GetType() == MODULE_PERSISTENT)
	lowercase_name = ZendNewInternedString(lowercase_name)
	ZendHashUpdatePtr(CG__().GetClassTable(), lowercase_name, class_entry)
	ZendStringReleaseEx(lowercase_name, 1)
	return class_entry
}
func ZendRegisterInternalClassEx(class_entry *ZendClassEntry, parent_ce *ZendClassEntry) *ZendClassEntry {
	var register_class *ZendClassEntry
	register_class = ZendRegisterInternalClass(class_entry)
	if parent_ce != nil {
		ZendDoInheritance(register_class, parent_ce)
		ZendBuildPropertiesInfoTable(register_class)
	}
	return register_class
}
func ZendClassImplements(class_entry *ZendClassEntry, num_interfaces int, _ ...any) {
	var interface_entry *ZendClassEntry
	var interface_list va_list
	va_start(interface_list, num_interfaces)
	for b.PostDec(&num_interfaces) {
		interface_entry = __va_arg(interface_list, (*ZendClassEntry)(_))
		ZendDoImplementInterface(class_entry, interface_entry)
	}
	va_end(interface_list)
}
func ZendRegisterInternalClass(orig_class_entry *ZendClassEntry) *ZendClassEntry {
	return DoRegisterInternalClass(orig_class_entry, 0)
}
func ZendRegisterInternalInterface(orig_class_entry *ZendClassEntry) *ZendClassEntry {
	return DoRegisterInternalClass(orig_class_entry, ZEND_ACC_INTERFACE)
}
func ZendRegisterClassAliasEx(name *byte, name_len int, ce *ZendClassEntry, persistent int) int {
	var lcname *ZendString
	var zv Zval
	var ret *Zval

	/* TODO: Move this out of here in 7.4. */

	if persistent != 0 && EG__().GetCurrentModule() != nil && EG__().GetCurrentModule().GetType() == MODULE_TEMPORARY {
		persistent = 0
	}
	if name[0] == '\\' {
		lcname = ZendStringAlloc(name_len-1, persistent)
		ZendStrTolowerCopy(lcname.GetVal(), name+1, name_len-1)
	} else {
		lcname = ZendStringAlloc(name_len, persistent)
		ZendStrTolowerCopy(lcname.GetVal(), name, name_len)
	}
	ZendAssertValidClassName(lcname)
	lcname = ZendNewInternedString(lcname)
	ZVAL_ALIAS_PTR(&zv, ce)
	ret = CG__().GetClassTable().KeyAdd(lcname.GetStr(), &zv)
	ZendStringReleaseEx(lcname, 0)
	if ret != nil {
		if !ce.IsImmutable() {
			ce.GetRefcount()++
		}
		return SUCCESS
	}
	return FAILURE
}
func ZendSetHashSymbol(
	symbol *Zval,
	name *byte,
	name_length int,
	is_ref ZendBool,
	num_symbol_tables int,
	_ ...any,
) int {
	var symbol_table *HashTable
	var symbol_table_list va_list
	if num_symbol_tables <= 0 {
		return FAILURE
	}
	if is_ref != 0 {
		ZVAL_MAKE_REF(symbol)
	}
	va_start(symbol_table_list, num_symbol_tables)
	for b.PostDec(&num_symbol_tables) > 0 {
		symbol_table = __va_arg(symbol_table_list, (*HashTable)(_))
		symbol_table.KeyUpdate(b.CastStr(name, name_length), symbol)
		symbol.TryAddRefcount()
	}
	va_end(symbol_table_list)
	return SUCCESS
}
func ZifDisplayDisabledFunction(executeData *ZendExecuteData, return_value *Zval) {
	ZendError(E_WARNING, "%s() has been disabled for security reasons", GetActiveFunctionName())
}
func ZendDisableFunction(function_name *byte, function_name_length int) int {
	var func_ *ZendInternalFunction
	if b.Assign(&func_, ZendHashStrFindPtr(CG__().GetFunctionTable(), function_name, function_name_length)) {
		ZendFreeInternalArgInfo(func_)
		func_.SubFnFlags(ZEND_ACC_VARIADIC | ZEND_ACC_HAS_TYPE_HINTS | ZEND_ACC_HAS_RETURN_TYPE)
		func_.SetNumArgs(0)
		func_.SetArgInfo(nil)
		func_.SetHandler(ZifDisplayDisabledFunction)
		return SUCCESS
	}
	return FAILURE
}
