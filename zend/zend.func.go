// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
	r "sik/runtime"
)

func HANDLE_BLOCK_INTERRUPTIONS() int {
	SIGG(depth)++
	return SIGG(depth) - 1
}
func HANDLE_UNBLOCK_INTERRUPTIONS() {
	if b.PostDec(&(SIGG(depth))) == SIGG(blocked) {
		ZendSignalHandlerUnblock()
	}
}
func USED_RET() bool {
	return !(EX(prev_execute_data)) || !(ZEND_USER_CODE(EX(prev_execute_data).func_.common.type_)) || EX(prev_execute_data).opline.result_type != IS_UNUSED
}
func ZendBailout()                                          { _zendBailout(__FILE__, __LINE__) }
func ZendPrintVariable(var_ *Zval) int                      { return ZendPrintZval(var_, 0) }
func ZEND_WRITE(str *byte, str_len int) int                 { return ZendWrite(str, str_len) }
func ZEND_WRITE_EX(str __auto__, str_len __auto__) __auto__ { return write_func(str, str_len) }
func ZEND_PUTS(str string) int                              { return ZendWrite(str, strlen(str)) }
func ZEND_PUTS_EX(str __auto__) __auto__                    { return write_func(str, strlen(str)) }
func ZEND_PUTC(c byte) int                                  { return ZendWrite(&c, 1) }
func ZEND_UV(name __auto__) __auto__                        { return ZendUv.name }
func OnUpdateErrorReporting(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	if new_value == nil {
		__EG().SetErrorReporting(E_ALL & ^E_NOTICE & ^E_STRICT & ^E_DEPRECATED)
	} else {
		__EG().SetErrorReporting(atoi(new_value.GetVal()))
	}
	return SUCCESS
}
func OnUpdateGCEnabled(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var val ZendBool
	val = ZendIniParseBool(new_value)
	GcEnable(val)
	return SUCCESS
}
func ZendGcEnabledDisplayerCb(ini_entry *ZendIniEntry, type_ int) {
	if GcEnabled() != 0 {
		ZEND_PUTS("On")
	} else {
		ZEND_PUTS("Off")
	}
}
func OnUpdateScriptEncoding(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	if __CG().GetMultibyte() == 0 {
		return FAILURE
	}
	if ZendMultibyteGetFunctions() == nil {
		return SUCCESS
	}
	return ZendMultibyteSetScriptEncodingByString(b.CondF1(new_value != nil, func() []byte { return new_value.GetVal() }, nil), b.CondF1(new_value != nil, func() int { return new_value.GetLen() }, 0))
}
func OnUpdateAssertions(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p *ZendLong
	var val ZendLong
	var base *byte = (*byte)(mh_arg2)
	p = (*ZendLong)(base + int(mh_arg1))
	val = ZendAtol(new_value.GetVal(), new_value.GetLen())
	if stage != ZEND_INI_STAGE_STARTUP && stage != ZEND_INI_STAGE_SHUTDOWN && (*p) != val && ((*p) < 0 || val < 0) {
		ZendError(E_WARNING, "zend.assertions may be completely enabled or disabled only in php.ini")
		return FAILURE
	}
	*p = val
	return SUCCESS
}
func ZendVspprintf(pbuf **byte, max_len int, format *byte, ap ...any) int {
	var buf SmartString = SmartString{0}

	/* since there are places where (v)spprintf called without checking for null,
	   a bit of defensive coding here */

	if pbuf == nil {
		return 0
	}
	ZendPrintfToSmartString(&buf, format, ap)
	if max_len != 0 && buf.GetLen() > max_len {
		buf.SetLen(max_len)
	}
	SmartString0(&buf)
	if buf.GetC() != nil {
		*pbuf = buf.GetC()
		return buf.GetLen()
	} else {
		*pbuf = Estrndup("", 0)
		return 0
	}
}
func ZendSpprintf(message **byte, max_len int, format string, _ ...any) int {
	var arg va_list
	var len_ int
	va_start(arg, format)
	len_ = ZendVspprintf(message, max_len, format, arg)
	va_end(arg)
	return len_
}
func ZendSpprintfUnchecked(message **byte, max_len int, format *byte, _ ...any) int {
	var arg va_list
	var len_ int
	va_start(arg, format)
	len_ = ZendVspprintf(message, max_len, format, arg)
	va_end(arg)
	return len_
}
func ZendVstrpprintf(max_len int, format *byte, ap ...any) *ZendString {
	var buf SmartStr = SmartStr{0}
	ZendPrintfToSmartStr(&buf, format, ap)
	if buf.GetS() == nil {
		return ZSTR_EMPTY_ALLOC()
	}
	if max_len != 0 && buf.GetS().GetLen() > max_len {
		buf.GetS().GetLen() = max_len
	}
	SmartStr0(&buf)
	return buf.GetS()
}
func ZendStrpprintf(max_len int, format string, _ ...any) *ZendString {
	var arg va_list
	var str *ZendString
	va_start(arg, format)
	str = ZendVstrpprintf(max_len, format, arg)
	va_end(arg)
	return str
}
func ZendStrpprintfUnchecked(max_len int, format string, _ ...any) *ZendString {
	var arg va_list
	var str *ZendString
	va_start(arg, format)
	str = ZendVstrpprintf(max_len, format, arg)
	va_end(arg)
	return str
}
func PrintHash(buf *SmartStr, ht *HashTable, indent int, is_object ZendBool) {
	var tmp *Zval
	var string_key *ZendString
	var num_key ZendUlong
	var i int
	for i = 0; i < indent; i++ {
		SmartStrAppendc(buf, ' ')
	}
	SmartStrAppends(buf, "(\n")
	indent += PRINT_ZVAL_INDENT
	var __ht *HashTable = ht
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()
		if _z.IsType(IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(IS_UNDEF) {
				continue
			}
		}
		num_key = _p.GetH()
		string_key = _p.GetKey()
		tmp = _z
		for i = 0; i < indent; i++ {
			SmartStrAppendc(buf, ' ')
		}
		SmartStrAppendc(buf, '[')
		if string_key != nil {
			if is_object != 0 {
				var prop_name *byte
				var class_name *byte
				var prop_len int
				var mangled int = ZendUnmanglePropertyNameEx(string_key, &class_name, &prop_name, &prop_len)
				SmartStrAppendl(buf, prop_name, prop_len)
				if class_name != nil && mangled == SUCCESS {
					if class_name[0] == '*' {
						SmartStrAppends(buf, ":protected")
					} else {
						SmartStrAppends(buf, ":")
						SmartStrAppends(buf, class_name)
						SmartStrAppends(buf, ":private")
					}
				}
			} else {
				SmartStrAppend(buf, string_key)
			}
		} else {
			SmartStrAppendLong(buf, num_key)
		}
		SmartStrAppends(buf, "] => ")
		ZendPrintZvalRToBuf(buf, tmp, indent+PRINT_ZVAL_INDENT)
		SmartStrAppends(buf, "\n")
	}
	indent -= PRINT_ZVAL_INDENT
	for i = 0; i < indent; i++ {
		SmartStrAppendc(buf, ' ')
	}
	SmartStrAppends(buf, ")\n")
}
func PrintFlatHash(ht *HashTable) {
	var tmp *Zval
	var string_key *ZendString
	var num_key ZendUlong
	var i int = 0
	var __ht *HashTable = ht
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()
		if _z.IsType(IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(IS_UNDEF) {
				continue
			}
		}
		num_key = _p.GetH()
		string_key = _p.GetKey()
		tmp = _z
		if b.PostInc(&i) > 0 {
			ZEND_PUTS(",")
		}
		ZEND_PUTS("[")
		if string_key != nil {
			ZEND_WRITE(string_key.GetVal(), string_key.GetLen())
		} else {
			ZendPrintf(ZEND_ULONG_FMT, num_key)
		}
		ZEND_PUTS("] => ")
		ZendPrintFlatZvalR(tmp)
	}
}
func ZendMakePrintableZval(expr *Zval, expr_copy *Zval) int {
	if expr.IsType(IS_STRING) {
		return 0
	} else {
		ZVAL_STR(expr_copy, ZvalGetStringFunc(expr))
		return 1
	}
}
func ZendPrintZval(expr *Zval, indent int) int {
	var tmp_str *ZendString
	var str *ZendString = ZvalGetTmpString(expr, &tmp_str)
	var len_ int = str.GetLen()
	if len_ != 0 {
		ZendWrite(str.GetVal(), len_)
	}
	ZendTmpStringRelease(tmp_str)
	return len_
}
func ZendPrintFlatZvalR(expr *Zval) {
	switch expr.GetType() {
	case IS_ARRAY:
		ZEND_PUTS("Array (")
		if (expr.GetArr().GetGcFlags() & GC_IMMUTABLE) == 0 {
			if GC_IS_RECURSIVE(expr.GetArr()) != 0 {
				ZEND_PUTS(" *RECURSION*")
				return
			}
			GC_PROTECT_RECURSION(expr.GetArr())
		}
		PrintFlatHash(expr.GetArr())
		ZEND_PUTS(")")
		if (expr.GetArr().GetGcFlags() & GC_IMMUTABLE) == 0 {
			GC_UNPROTECT_RECURSION(expr.GetArr())
		}
		break
	case IS_OBJECT:
		var properties *HashTable
		var class_name *ZendString = Z_OBJ_HT(*expr).GetGetClassName()(expr.GetObj())
		ZendPrintf("%s Object (", class_name.GetVal())
		ZendStringReleaseEx(class_name, 0)
		if GC_IS_RECURSIVE(expr.GetCounted()) != 0 {
			ZEND_PUTS(" *RECURSION*")
			return
		}
		properties = Z_OBJPROP_P(expr)
		if properties != nil {
			GC_PROTECT_RECURSION(expr.GetObj())
			PrintFlatHash(properties)
			GC_UNPROTECT_RECURSION(expr.GetObj())
		}
		ZEND_PUTS(")")
		break
	case IS_REFERENCE:
		ZendPrintFlatZvalR(Z_REFVAL_P(expr))
		break
	default:
		ZendPrintZval(expr, 0)
		break
	}
}
func ZendPrintZvalRToBuf(buf *SmartStr, expr *Zval, indent int) {
	switch expr.GetType() {
	case IS_ARRAY:
		SmartStrAppends(buf, "Array\n")
		if (expr.GetArr().GetGcFlags() & GC_IMMUTABLE) == 0 {
			if GC_IS_RECURSIVE(expr.GetArr()) != 0 {
				SmartStrAppends(buf, " *RECURSION*")
				return
			}
			GC_PROTECT_RECURSION(expr.GetArr())
		}
		PrintHash(buf, expr.GetArr(), indent, 0)
		if (expr.GetArr().GetGcFlags() & GC_IMMUTABLE) == 0 {
			GC_UNPROTECT_RECURSION(expr.GetArr())
		}
		break
	case IS_OBJECT:
		var properties *HashTable
		var class_name *ZendString = Z_OBJ_HT(*expr).GetGetClassName()(expr.GetObj())
		SmartStrAppends(buf, class_name.GetVal())
		ZendStringReleaseEx(class_name, 0)
		SmartStrAppends(buf, " Object\n")
		if GC_IS_RECURSIVE(expr.GetObj()) != 0 {
			SmartStrAppends(buf, " *RECURSION*")
			return
		}
		if b.Assign(&properties, ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_DEBUG)) == nil {
			break
		}
		GC_PROTECT_RECURSION(expr.GetObj())
		PrintHash(buf, properties, indent, 1)
		GC_UNPROTECT_RECURSION(expr.GetObj())
		ZendReleaseProperties(properties)
		break
	case IS_LONG:
		SmartStrAppendLong(buf, expr.GetLval())
		break
	case IS_REFERENCE:
		ZendPrintZvalRToBuf(buf, Z_REFVAL_P(expr), indent)
		break
	case IS_STRING:
		SmartStrAppend(buf, expr.GetStr())
		break
	default:
		var str *ZendString = ZvalGetStringFunc(expr)
		SmartStrAppend(buf, str)
		ZendStringReleaseEx(str, 0)
		break
	}
}
func ZendPrintZvalRToStr(expr *Zval, indent int) *ZendString {
	var buf SmartStr = SmartStr{0}
	ZendPrintZvalRToBuf(&buf, expr, indent)
	SmartStr0(&buf)
	return buf.GetS()
}
func ZendPrintZvalR(expr *Zval, indent int) {
	var str *ZendString = ZendPrintZvalRToStr(expr, indent)
	ZendWrite(str.GetVal(), str.GetLen())
	ZendStringReleaseEx(str, 0)
}
func ZendFopenWrapper(filename *byte, opened_path **ZendString) *r.FILE {
	if opened_path != nil {
		*opened_path = ZendStringInit(filename, strlen(filename), 0)
	}
	return r.Fopen(filename, "rb")
}
func ZendSetDefaultCompileTimeValues() {
	/* default compile-time values */

	__CG().SetShortTags(ShortTagsDefault)
	__CG().SetCompilerOptions(CompilerOptionsDefault)
	__CG().SetRtdKeyCounter(0)
}
func ZendInitExceptionOp() {
	memset(__EG().GetExceptionOp(), 0, b.SizeOf("EG ( exception_op )"))
	__EG().GetExceptionOp()[0].SetOpcode(ZEND_HANDLE_EXCEPTION)
	ZEND_VM_SET_OPCODE_HANDLER(__EG().GetExceptionOp())
	__EG().GetExceptionOp()[1].SetOpcode(ZEND_HANDLE_EXCEPTION)
	ZEND_VM_SET_OPCODE_HANDLER(__EG().GetExceptionOp() + 1)
	__EG().GetExceptionOp()[2].SetOpcode(ZEND_HANDLE_EXCEPTION)
	ZEND_VM_SET_OPCODE_HANDLER(__EG().GetExceptionOp() + 2)
}
func ZendInitCallTrampolineOp() {
	memset(__EG().GetCallTrampolineOp(), 0, b.SizeOf("EG ( call_trampoline_op )"))
	__EG().GetCallTrampolineOp().SetOpcode(ZEND_CALL_TRAMPOLINE)
	ZEND_VM_SET_OPCODE_HANDLER(__EG().GetCallTrampolineOp())
}
func AutoGlobalDtor(zv *Zval) { Free(zv.GetPtr()) }
func IniScannerGlobalsCtor(scanner_globals_p *ZendIniScannerGlobals) {
	memset(scanner_globals_p, 0, b.SizeOf("* scanner_globals_p"))
}
func PhpScannerGlobalsCtor(scanner_globals_p *ZendPhpScannerGlobals) {
	memset(scanner_globals_p, 0, b.SizeOf("* scanner_globals_p"))
}
func ModuleDestructorZval(zv *Zval) {
	var module *ZendModuleEntry = (*ZendModuleEntry)(zv.GetPtr())
	ModuleDestructor(module)
	Free(module)
}
func PhpAutoGlobalsCreateGlobals(name *ZendString) ZendBool {
	var globals Zval

	/* IS_ARRAY, but with ref-counter 1 and not IS_TYPE_REFCOUNTED */

	ZVAL_ARR(&globals, __EG().GetSymbolTable())
	globals.SetTypeFlags(0)
	ZVAL_NEW_REF(&globals, &globals)
	__EG().GetSymbolTable().KeyUpdate(name.GetStr(), &globals)
	return 0
}
func ZendStartup(utility_functions *ZendUtilityFunctions) int {
	var ini_scanner_globals ZendIniScannerGlobals
	var language_scanner_globals ZendPhpScannerGlobals
	ZendCpuStartup()
	StartMemoryManager()
	VirtualCwdStartup()
	ZendStartupStrtod()
	ZendStartupExtensionsMechanism()

	/* Set up utility functions and values */

	ZendErrorCb = utility_functions.GetErrorFunction()
	ZendPrintf = utility_functions.GetPrintfFunction()
	ZendWrite = ZendWriteFuncT(utility_functions.GetWriteFunction())
	ZendFopen = utility_functions.GetFopenFunction()
	if ZendFopen == nil {
		ZendFopen = ZendFopenWrapper
	}
	ZendStreamOpenFunction = utility_functions.GetStreamOpenFunction()
	ZendMessageDispatcherP = utility_functions.GetMessageHandler()
	ZendGetConfigurationDirectiveP = utility_functions.GetGetConfigurationDirective()
	ZendTicksFunction = utility_functions.GetTicksFunction()
	ZendOnTimeout = utility_functions.GetOnTimeout()
	ZendPrintfToSmartString = utility_functions.GetPrintfToSmartStringFunction()
	ZendPrintfToSmartStr = utility_functions.GetPrintfToSmartStrFunction()
	ZendGetenv = utility_functions.GetGetenvFunction()
	ZendResolvePath = utility_functions.GetResolvePathFunction()
	ZendInterruptFunction = nil
	ZendCompileFile = CompileFile
	ZendExecuteEx = ExecuteEx
	ZendExecuteInternal = nil
	ZendCompileString = CompileString
	ZendThrowExceptionHook = nil

	/* Set up the default garbage collection implementation. */

	GcCollectCycles = ZendGcCollectCycles
	ZendVmInit()

	/* set up version */

	ZendVersionInfo = strdup(ZEND_CORE_VERSION_INFO)
	ZendVersionInfoLength = b.SizeOf("ZEND_CORE_VERSION_INFO") - 1
	GLOBAL_FUNCTION_TABLE = (*HashTable)(Malloc(b.SizeOf("HashTable")))
	GLOBAL_CLASS_TABLE = (*HashTable)(Malloc(b.SizeOf("HashTable")))
	GLOBAL_AUTO_GLOBALS_TABLE = (*HashTable)(Malloc(b.SizeOf("HashTable")))
	GLOBAL_CONSTANTS_TABLE = (*HashTable)(Malloc(b.SizeOf("HashTable")))
	ZendHashInitEx(GLOBAL_FUNCTION_TABLE, 1024, nil, ZEND_FUNCTION_DTOR, 1, 0)
	ZendHashInitEx(GLOBAL_CLASS_TABLE, 64, nil, ZEND_CLASS_DTOR, 1, 0)
	ZendHashInitEx(GLOBAL_AUTO_GLOBALS_TABLE, 8, nil, AutoGlobalDtor, 1, 0)
	ZendHashInitEx(GLOBAL_CONSTANTS_TABLE, 128, nil, ZEND_CONSTANT_DTOR, 1, 0)
	ZendHashInitEx(&ModuleRegistry, 32, nil, ModuleDestructorZval, 1, 0)
	ZendInitRsrcListDtors()
	IniScannerGlobalsCtor(&ini_scanner_globals)
	PhpScannerGlobalsCtor(&language_scanner_globals)
	ZendSetDefaultCompileTimeValues()

	/* Map region is going to be created and resized at run-time. */

	__CG().SetMapPtrBase(nil)
	__CG().SetMapPtrSize(0)
	__CG().SetMapPtrLast(0)
	__EG().SetErrorReporting(E_ALL & ^E_NOTICE)
	ZendInternedStringsInit()
	ZendStartupBuiltinFunctions()
	ZendRegisterStandardConstants()
	ZendRegisterAutoGlobal(ZendStringInitInterned("GLOBALS", b.SizeOf("\"GLOBALS\"")-1, 1), 1, PhpAutoGlobalsCreateGlobals)
	ZendInitRsrcPlist()
	ZendInitExceptionOp()
	ZendInitCallTrampolineOp()
	ZendIniStartup()
	return SUCCESS
}
func ZendRegisterStandardIniEntries() {
	var module_number int = 0
	REGISTER_INI_ENTRIES()
}
func ZendResolvePropertyTypes() {
	var ce *ZendClassEntry
	var prop_info *ZendPropertyInfo
	var __ht *HashTable = __CG().GetClassTable()
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		ce = _z.GetPtr()
		if ce.GetType() != ZEND_INTERNAL_CLASS {
			continue
		}
		if ZEND_CLASS_HAS_TYPE_HINTS(ce) {
			var __ht *HashTable = ce.GetPropertiesInfo()
			for _, _p := range __ht.foreachData() {
				var _z *Zval = _p.GetVal()

				prop_info = _z.GetPtr()
				if prop_info.GetType().IsName() {
					var type_name *ZendString = prop_info.GetType().Name()
					var lc_type_name *ZendString = ZendStringTolower(type_name)
					var prop_ce *ZendClassEntry = ZendHashFindPtr(__CG().GetClassTable(), lc_type_name)
					ZEND_ASSERT(prop_ce != nil && prop_ce.GetType() == ZEND_INTERNAL_CLASS)
					prop_info.SetType(ZEND_TYPE_ENCODE_CE(prop_ce, prop_info.GetType().AllowNull()))
					ZendStringRelease(lc_type_name)
					ZendStringRelease(type_name)
				}
			}
		}
		ce.SetIsPropertyTypesResolved(true)
	}
}
func ZendPostStartup() int {
	ZendResolvePropertyTypes()
	if ZendPostStartupCb != nil {
		var cb func() int = ZendPostStartupCb
		ZendPostStartupCb = nil
		if cb() != SUCCESS {
			return FAILURE
		}
	}
	GlobalMapPtrLast = __CG().GetMapPtrLast()
	return SUCCESS
}
func ZendShutdown() {
	ZendVmDtor()
	ZendDestroyRsrcList(__EG().GetPersistentList())
	ZendDestroyModules()
	VirtualCwdDeactivate()
	VirtualCwdShutdown()
	GLOBAL_FUNCTION_TABLE.Destroy()
	GLOBAL_CLASS_TABLE.Destroy()
	GLOBAL_AUTO_GLOBALS_TABLE.Destroy()
	Free(GLOBAL_AUTO_GLOBALS_TABLE)
	ZendShutdownExtensions()
	Free(ZendVersionInfo)
	Free(GLOBAL_FUNCTION_TABLE)
	Free(GLOBAL_CLASS_TABLE)
	GLOBAL_CONSTANTS_TABLE.Destroy()
	Free(GLOBAL_CONSTANTS_TABLE)
	ZendShutdownStrtod()
	if __CG().GetMapPtrBase() {
		Free(__CG().GetMapPtrBase())
		__CG().SetMapPtrBase(nil)
		__CG().SetMapPtrSize(0)
	}
	if __CG().GetScriptEncodingList() != nil {
		Free(__CG().GetScriptEncodingList())
		__CG().SetScriptEncodingList(nil)
		__CG().SetScriptEncodingListSize(0)
	}
	ZendDestroyRsrcListDtors()
}
func ZendSetUtilityValues(utility_values *ZendUtilityValues) { ZendUv = *utility_values }
func Zenderror(error *byte) {
	__CG().SetParseError(0)
	if __EG().GetException() != nil {

		/* An exception was thrown in the lexer, don't throw another in the parser. */

		return

		/* An exception was thrown in the lexer, don't throw another in the parser. */

	}
	ZendThrowException(ZendCeParseError, error, 0)
}
func _zendBailout(filename *byte, lineno uint32) {
	if __EG().GetBailout() == nil {
		ZendOutputDebugString(1, "%s(%d) : Bailed out without a bailout address!", filename, lineno)
		exit(-1)
	}
	GcProtect(1)
	__CG().SetUncleanShutdown(1)
	__CG().SetActiveClassEntry(nil)
	__CG().SetInCompilation(0)
	__EG().SetCurrentExecuteData(nil)
	LONGJMP((*__EG)().bailout, FAILURE)
}
func ZendAppendVersionInfo(extension *ZendExtension) {
	var new_info *byte
	var new_info_length uint32
	new_info_length = uint32(b.SizeOf("\"    with  v, , by \\n\"") + strlen(extension.GetName()) + strlen(extension.GetVersion()) + strlen(extension.GetCopyright()) + strlen(extension.GetAuthor()))
	new_info = (*byte)(Malloc(new_info_length + 1))
	core.Snprintf(new_info, new_info_length, "    with %s v%s, %s, by %s\n", extension.GetName(), extension.GetVersion(), extension.GetCopyright(), extension.GetAuthor())
	ZendVersionInfo = (*byte)(realloc(ZendVersionInfo, ZendVersionInfoLength+new_info_length+1))
	strncat(ZendVersionInfo, new_info, new_info_length)
	ZendVersionInfoLength += new_info_length
	Free(new_info)
}
func GetZendVersion() *byte { return ZendVersionInfo }
func ZendActivate() {
	GcReset()
	InitCompiler()
	InitExecutor()
	StartupScanner()
	if __CG().GetMapPtrLast() != 0 {
		memset(__CG().GetMapPtrBase(), 0, __CG().GetMapPtrLast()*b.SizeOf("void *"))
	}
}
func ZendCallDestructors() {
	var __orig_bailout *JMP_BUF = __EG().GetBailout()
	var __bailout JMP_BUF
	__EG().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ShutdownDestructors()
	}
	__EG().SetBailout(__orig_bailout)
}
func ZendDeactivate() {
	/* we're no longer executing anything */

	__EG().SetCurrentExecuteData(nil)
	var __orig_bailout *JMP_BUF = __EG().bailout
	var __bailout JMP_BUF
	__EG().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ShutdownScanner()
	}
	__EG().SetBailout(__orig_bailout)

	/* shutdown_executor() takes care of its own bailout handling */

	ShutdownExecutor()
	var __orig_bailout *JMP_BUF = __EG().bailout
	var __bailout JMP_BUF
	__EG().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ZendIniDeactivate()
	}
	__EG().SetBailout(__orig_bailout)
	var __orig_bailout *JMP_BUF = __EG().GetBailout()
	var __bailout JMP_BUF
	__EG().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ShutdownCompiler()
	}
	__EG().SetBailout(__orig_bailout)
	ZendDestroyRsrcList(__EG().GetRegularList())
}
func ZendMessageDispatcher(message ZendLong, data any) {
	if ZendMessageDispatcherP != nil {
		ZendMessageDispatcherP(message, data)
	}
}
func ZendGetConfigurationDirective(name *ZendString) *Zval {
	if ZendGetConfigurationDirectiveP != nil {
		return ZendGetConfigurationDirectiveP(name)
	} else {
		return nil
	}
}
func SAVE_STACK(stack ZendStack) {
	if __CG().stack.top {
		memcpy(&stack, __CG().stack, b.SizeOf("zend_stack"))
		__CG().stack.max = 0
		__CG().stack.top = __CG().stack.max
		__CG().stack.elements = nil
	} else {
		stack.SetTop(0)
	}
}
func RESTORE_STACK(stack ZendStack) {
	if stack.GetTop() != 0 {
		ZendStackDestroy(__CG().stack)
		memcpy(__CG().stack, &stack, b.SizeOf("zend_stack"))
	}
}
func ZendErrorVaList(type_ int, error_filename *byte, error_lineno uint32, format *byte, args ...any) {
	var usr_copy va_list
	var params []Zval
	var retval Zval
	var orig_user_error_handler Zval
	var in_compilation ZendBool
	var saved_class_entry *ZendClassEntry
	var loop_var_stack ZendStack
	var delayed_oplines_stack ZendStack
	var symbol_table *ZendArray
	var orig_fake_scope *ZendClassEntry

	/* Report about uncaught exception in case of fatal errors */

	if __EG().GetException() != nil {
		var ex *ZendExecuteData
		var opline *ZendOp
		switch type_ {
		case E_CORE_ERROR:

		case E_ERROR:

		case E_RECOVERABLE_ERROR:

		case E_PARSE:

		case E_COMPILE_ERROR:

		case E_USER_ERROR:
			ex = __EG().GetCurrentExecuteData()
			opline = nil
			for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
				ex = ex.GetPrevExecuteData()
			}
			if ex != nil && ex.GetOpline().GetOpcode() == ZEND_HANDLE_EXCEPTION && __EG().GetOplineBeforeException() != nil {
				opline = __EG().GetOplineBeforeException()
			}
			ZendExceptionError(__EG().GetException(), E_WARNING)
			__EG().SetException(nil)
			if opline != nil {
				ex.SetOpline(opline)
			}
			break
		default:
			break
		}
	}

	/* if we don't have a user defined error handler */

	if __EG().GetUserErrorHandler().IsType(IS_UNDEF) || (__EG().GetUserErrorHandlerErrorReporting()&type_) == 0 || __EG().GetErrorHandling() != EH_NORMAL {
		ZendErrorCb(type_, error_filename, error_lineno, format, args)
	} else {
		switch type_ {
		case E_ERROR:

		case E_PARSE:

		case E_CORE_ERROR:

		case E_CORE_WARNING:

		case E_COMPILE_ERROR:

		case E_COMPILE_WARNING:

			/* The error may not be safe to handle in user-space */

			ZendErrorCb(type_, error_filename, error_lineno, format, args)
			break
		default:

			/* Handle the error in user space */

			VaCopy(usr_copy, args)
			ZVAL_STR(&params[1], ZendVstrpprintf(0, format, usr_copy))
			va_end(usr_copy)
			ZVAL_LONG(&params[0], type_)
			if error_filename != nil {
				ZVAL_STRING(&params[2], error_filename)
			} else {
				ZVAL_NULL(&params[2])
			}
			ZVAL_LONG(&params[3], error_lineno)
			symbol_table = ZendRebuildSymbolTable()

			/* during shutdown the symbol table table can be still null */

			if symbol_table == nil {
				ZVAL_NULL(&params[4])
			} else {
				ZVAL_ARR(&params[4], ZendArrayDup(symbol_table))
			}
			ZVAL_COPY_VALUE(&orig_user_error_handler, __EG().GetUserErrorHandler())
			ZVAL_UNDEF(__EG().GetUserErrorHandler())

			/* User error handler may include() additinal PHP files.
			 * If an error was generated during comilation PHP will compile
			 * such scripts recursively, but some CG() variables may be
			 * inconsistent. */

			in_compilation = __CG().GetInCompilation()
			if in_compilation != 0 {
				saved_class_entry = __CG().GetActiveClassEntry()
				__CG().SetActiveClassEntry(nil)
				SAVE_STACK(loop_var_stack)
				SAVE_STACK(delayed_oplines_stack)
				__CG().SetInCompilation(0)
			}
			orig_fake_scope = __EG().GetFakeScope()
			__EG().SetFakeScope(nil)
			if CallUserFunction(__CG().GetFunctionTable(), nil, &orig_user_error_handler, &retval, 5, params) == SUCCESS {
				if retval.GetType() != IS_UNDEF {
					if retval.IsType(IS_FALSE) {
						ZendErrorCb(type_, error_filename, error_lineno, format, args)
					}
					ZvalPtrDtor(&retval)
				}
			} else if __EG().GetException() == nil {

				/* The user error handler failed, use built-in error handler */

				ZendErrorCb(type_, error_filename, error_lineno, format, args)

				/* The user error handler failed, use built-in error handler */

			}
			__EG().SetFakeScope(orig_fake_scope)
			if in_compilation != 0 {
				__CG().SetActiveClassEntry(saved_class_entry)
				RESTORE_STACK(loop_var_stack)
				RESTORE_STACK(delayed_oplines_stack)
				__CG().SetInCompilation(1)
			}
			ZvalPtrDtor(&params[4])
			ZvalPtrDtor(&params[2])
			ZvalPtrDtor(&params[1])
			if __EG().GetUserErrorHandler().IsType(IS_UNDEF) {
				ZVAL_COPY_VALUE(__EG().GetUserErrorHandler(), &orig_user_error_handler)
			} else {
				ZvalPtrDtor(&orig_user_error_handler)
			}
			break
		}
	}
	if type_ == E_PARSE {

		/* eval() errors do not affect exit_status */

		if !(__EG().GetCurrentExecuteData() != nil && __EG().GetCurrentExecuteData().GetFunc() != nil && ZEND_USER_CODE(__EG().GetCurrentExecuteData().GetFunc().GetType()) && __EG().GetCurrentExecuteData().GetOpline().GetOpcode() == ZEND_INCLUDE_OR_EVAL && __EG().GetCurrentExecuteData().GetOpline().GetExtendedValue() == ZEND_EVAL) {
			__EG().SetExitStatus(255)
		}

		/* eval() errors do not affect exit_status */

	}
}
func GetFilenameLineno(type_ int, filename **byte, lineno *uint32) {
	/* Obtain relevant filename and lineno */

	switch type_ {
	case E_CORE_ERROR:

	case E_CORE_WARNING:
		*filename = nil
		*lineno = 0
		break
	case E_PARSE:

	case E_COMPILE_ERROR:

	case E_COMPILE_WARNING:

	case E_ERROR:

	case E_NOTICE:

	case E_STRICT:

	case E_DEPRECATED:

	case E_WARNING:

	case E_USER_ERROR:

	case E_USER_WARNING:

	case E_USER_NOTICE:

	case E_USER_DEPRECATED:

	case E_RECOVERABLE_ERROR:
		if ZendIsCompiling() != 0 {
			*filename = ZendGetCompiledFilename().GetVal()
			*lineno = ZendGetCompiledLineno()
		} else if ZendIsExecuting() != 0 {
			*filename = ZendGetExecutedFilename()
			if (*filename)[0] == '[' {
				*filename = nil
				*lineno = 0
			} else {
				*lineno = ZendGetExecutedLineno()
			}
		} else {
			*filename = nil
			*lineno = 0
		}
		break
	default:
		*filename = nil
		*lineno = 0
		break
	}
	if (*filename) == nil {
		*filename = "Unknown"
	}
}
func ZendErrorAt(type_ int, filename *byte, lineno uint32, format string, _ ...any) {
	var args va_list
	if filename == nil {
		var dummy_lineno uint32
		GetFilenameLineno(type_, &filename, &dummy_lineno)
	}
	va_start(args, format)
	ZendErrorVaList(type_, filename, lineno, format, args)
	va_end(args)
}
func ZendError(type_ int, format string, _ ...any) {
	var filename *byte
	var lineno uint32
	var args va_list
	GetFilenameLineno(type_, &filename, &lineno)
	va_start(args, format)
	ZendErrorVaList(type_, filename, lineno, format, args)
	va_end(args)
}
func ZendErrorAtNoreturn(type_ int, filename *byte, lineno uint32, format string, _ ...any) {
	var args va_list
	if filename == nil {
		var dummy_lineno uint32
		GetFilenameLineno(type_, &filename, &dummy_lineno)
	}
	va_start(args, format)
	ZendErrorVaList(type_, filename, lineno, format, args)
	va_end(args)

	/* Should never reach this. */

	abort()

	/* Should never reach this. */
}
func ZendErrorNoreturn(type_ int, format string, _ ...any) {
	var filename *byte
	var lineno uint32
	var args va_list
	GetFilenameLineno(type_, &filename, &lineno)
	va_start(args, format)
	ZendErrorVaList(type_, filename, lineno, format, args)
	va_end(args)

	/* Should never reach this. */

	abort()

	/* Should never reach this. */
}
func ZendThrowError(exception_ce *ZendClassEntry, format string, _ ...any) {
	var va va_list
	var message *byte = nil
	if exception_ce != nil {
		if InstanceofFunction(exception_ce, ZendCeError) == 0 {
			ZendError(E_NOTICE, "Error exceptions must be derived from Error")
			exception_ce = ZendCeError
		}
	} else {
		exception_ce = ZendCeError
	}

	/* Marker used to disable exception generation during preloading. */

	if __EG().GetException() == any(uintPtr-1) {
		return
	}
	va_start(va, format)
	ZendVspprintf(&message, 0, format, va)

	//TODO: we can't convert compile-time errors to exceptions yet???

	if __EG().GetCurrentExecuteData() != nil && __CG().GetInCompilation() == 0 {
		ZendThrowException(exception_ce, message, 0)
	} else {
		ZendError(E_ERROR, "%s", message)
	}
	Efree(message)
	va_end(va)
}
func ZendTypeError(format string, _ ...any) {
	var va va_list
	var message *byte = nil
	va_start(va, format)
	ZendVspprintf(&message, 0, format, va)
	ZendThrowException(ZendCeTypeError, message, 0)
	Efree(message)
	va_end(va)
}
func ZendInternalTypeError(throw_exception ZendBool, format string, _ ...any) {
	var va va_list
	var message *byte = nil
	va_start(va, format)
	ZendVspprintf(&message, 0, format, va)
	if throw_exception != 0 {
		ZendThrowException(ZendCeTypeError, message, 0)
	} else {
		ZendError(E_WARNING, "%s", message)
	}
	Efree(message)
	va_end(va)
}
func ZendInternalArgumentCountError(throw_exception ZendBool, format string, _ ...any) {
	var va va_list
	var message *byte = nil
	va_start(va, format)
	ZendVspprintf(&message, 0, format, va)
	if throw_exception != 0 {
		ZendThrowException(ZendCeArgumentCountError, message, 0)
	} else {
		ZendError(E_WARNING, "%s", message)
	}
	Efree(message)
	va_end(va)
}
func ZendOutputDebugString(trigger_break ZendBool, format string, _ ...any) {}
func ZendUserExceptionHandler() {
	var orig_user_exception_handler Zval
	var params []Zval
	var retval2 Zval
	var old_exception *ZendObject
	old_exception = __EG().GetException()
	__EG().SetException(nil)
	ZVAL_OBJ(&params[0], old_exception)
	ZVAL_COPY_VALUE(&orig_user_exception_handler, __EG().GetUserExceptionHandler())
	if CallUserFunction(__CG().GetFunctionTable(), nil, &orig_user_exception_handler, &retval2, 1, params) == SUCCESS {
		ZvalPtrDtor(&retval2)
		if __EG().GetException() != nil {
			OBJ_RELEASE(__EG().GetException())
			__EG().SetException(nil)
		}
		OBJ_RELEASE(old_exception)
	} else {
		__EG().SetException(old_exception)
	}
}
func ZendExecuteScripts(type_ int, retval *Zval, file_count int, _ ...any) int {
	var files va_list
	var i int
	var file_handle *ZendFileHandle
	var op_array *ZendOpArray
	va_start(files, file_count)
	for i = 0; i < file_count; i++ {
		file_handle = __va_arg(files, (*ZendFileHandle)(_))
		if file_handle == nil {
			continue
		}
		op_array = ZendCompileFile(file_handle, type_)
		if file_handle.GetOpenedPath() != nil {
			ZendHashAddEmptyElement(__EG().GetIncludedFiles(), file_handle.GetOpenedPath())
		}
		ZendDestroyFileHandle(file_handle)
		if op_array != nil {
			ZendExecute(op_array, retval)
			ZendExceptionRestore()
			if __EG().GetException() != nil {
				if __EG().GetUserExceptionHandler().GetType() != IS_UNDEF {
					ZendUserExceptionHandler()
				}
				if __EG().GetException() != nil {
					ZendExceptionError(__EG().GetException(), E_ERROR)
				}
			}
			DestroyOpArray(op_array)
			EfreeSize(op_array, b.SizeOf("zend_op_array"))
		} else if type_ == ZEND_REQUIRE {
			va_end(files)
			return FAILURE
		}
	}
	va_end(files)
	return SUCCESS
}
func ZendMakeCompiledStringDescription(name string) *byte {
	var cur_filename *byte
	var cur_lineno int
	var compiled_string_description *byte
	if ZendIsCompiling() != 0 {
		cur_filename = ZendGetCompiledFilename().GetVal()
		cur_lineno = ZendGetCompiledLineno()
	} else if ZendIsExecuting() != 0 {
		cur_filename = ZendGetExecutedFilename()
		cur_lineno = ZendGetExecutedLineno()
	} else {
		cur_filename = "Unknown"
		cur_lineno = 0
	}
	ZendSpprintf(&compiled_string_description, 0, COMPILED_STRING_DESCRIPTION_FORMAT, cur_filename, cur_lineno, name)
	return compiled_string_description
}
func FreeEstring(str_p **byte) { Efree(*str_p) }
func ZendMapPtrReset()         { __CG().SetMapPtrLast(GlobalMapPtrLast) }
func ZendMapPtrNew() any {
	var ptr *any
	if __CG().GetMapPtrLast() >= __CG().GetMapPtrSize() {

		/* Grow map_ptr table */

		__CG().SetMapPtrSize(ZEND_MM_ALIGNED_SIZE_EX(__CG().GetMapPtrLast()+1, 4096))
		__CG().SetMapPtrBase(Perealloc(__CG().GetMapPtrBase(), __CG().GetMapPtrSize()*b.SizeOf("void *"), 1))
	}
	ptr = (*any)(__CG().GetMapPtrBase() + __CG().GetMapPtrLast())
	*ptr = nil
	__CG().GetMapPtrLast()++
	return ZEND_MAP_PTR_PTR2OFFSET(ptr)
}
func ZendMapPtrExtend(last int) {
	if last > __CG().GetMapPtrLast() {
		var ptr *any
		if last >= __CG().GetMapPtrSize() {

			/* Grow map_ptr table */

			__CG().SetMapPtrSize(ZEND_MM_ALIGNED_SIZE_EX(last, 4096))
			__CG().SetMapPtrBase(Perealloc(__CG().GetMapPtrBase(), __CG().GetMapPtrSize()*b.SizeOf("void *"), 1))
		}
		ptr = (*any)(__CG().GetMapPtrBase() + __CG().GetMapPtrLast())
		memset(ptr, 0, (last-__CG().GetMapPtrLast())*b.SizeOf("void *"))
		__CG().SetMapPtrLast(last)
	}
}
