// <<generate>>

package zend

import (
	"fmt"
	b "sik/builtin"
	r "sik/builtin/file"
	"sik/zend/faults"
	"sik/zend/types"
)

func USED_RET() bool {
	return !(executeData.GetPrevExecuteData()) || !(ZEND_USER_CODE(executeData.GetPrevExecuteData().func_.common.type_)) || executeData.GetPrevExecuteData().opline.result_type != IS_UNUSED
}
func ZEND_PUTS(str string) int { return ZendWrite(str) }
func ZEND_PUTC(c byte) int     { return ZendWrite(string([]byte{c})) }

func OnUpdateErrorReportingEx(entry *ZendIniEntry, newValue *string, stage int) bool {
	if newValue == nil {
		EG__().SetErrorReporting(faults.E_ALL & ^faults.E_NOTICE & ^faults.E_STRICT & ^faults.E_DEPRECATED)
	} else {
		EG__().SetErrorReporting(b.Atoi(*newValue))
	}
	return true
}
func OnUpdateAssertionsEx(entry *ZendIniEntry, new_value *string, stage int) bool {
	if new_value == nil {
		// todo
		return true
	}

	assertions := EG__().assertions
	val := ZendAtolEx(*new_value)
	if stage != ZEND_INI_STAGE_STARTUP && stage != ZEND_INI_STAGE_SHUTDOWN && assertions != val && (assertions < 0 || val < 0) {
		faults.Error(faults.E_WARNING, "zend.assertions may be completely enabled or disabled only in php.ini")
		return false
	}
	EG__().assertions = val
	return true
}

// 替代各种 sprintf 方法(限制长度)
func ZendSprintfEx(maxLen int, format string, args ...any) string {
	result := ZendSprintf(format, args...)
	if maxLen != 0 && len(result) > maxLen {
		return result[:maxLen]
	}
	return result
}

// 替代各种 sprintf 方法
func ZendSprintf(format string, args ...any) string {
	var buf = SmartStr{}
	ZendPrintfToSmartStr(&buf, format, args...)
	return buf.GetStr()
}

func ZendVspprintf(pbuf *string, max_len int, format string, args ...any) int {
	/* since there are places where (v)spprintf called without checking for null,
	   a bit of defensive coding here */
	if pbuf == nil {
		return 0
	}
	result := ZendSprintfEx(max_len, format, args...)
	*pbuf = result
	return len(result)
}

func ZendSpprintf(message *string, max_len int, format string, args ...any) int {
	result := ZendSprintfEx(max_len, format, args...)
	*message = result
	return len(result)
}

func ZendStrpprintf(max_len int, format string, args ...any) *types.ZendString {
	result := ZendSprintfEx(max_len, format, args...)
	return types.NewZendString(result)
}

func PrintHash(buf *SmartStr, ht *types.HashTable, indent int, is_object types.ZendBool) {
	for i := 0; i < indent; i++ {
		buf.AppendByte(' ')
	}
	buf.AppendString("(\n")
	indent += PRINT_ZVAL_INDENT
	ht.eachValidBucketIndirect(func(_ uint32, p *types.Bucket, z *types.Zval) {
		for i := 0; i < indent; i++ {
			buf.AppendByte(' ')
		}
		buf.AppendByte('[')
		if p.IsStrKey() {
			if is_object != 0 {
				className, propName, mangled := ZendUnmanglePropertyName_Ex(p.StrKey())
				buf.AppendString(propName)
				if className != "" && mangled {
					if className[0] == '*' {
						buf.AppendString(":protected")
					} else {
						buf.AppendString(":")
						buf.AppendString(className)
						buf.AppendString(":private")
					}
				}
			} else {
				buf.AppendString(p.StrKey())
			}
		} else {
			buf.AppendLong(p.IndexKey())
		}
		buf.AppendString("] => ")
		ZendPrintZvalRToBuf(buf, z, indent+PRINT_ZVAL_INDENT)
		buf.AppendString("\n")
	})
	indent -= PRINT_ZVAL_INDENT
	for i := 0; i < indent; i++ {
		buf.AppendByte(' ')
	}
	buf.AppendString(")\n")
}
func PrintFlatHash(ht *types.HashTable) {
	var tmp *types.Zval
	var string_key *types.ZendString
	var num_key ZendUlong
	var i int = 0
	var __ht *types.HashTable = ht
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.GetZv()
			if _z.IsUndef() {
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
			ZendWrite(string_key.GetStr())
		} else {
			ZendPrintf(ZEND_ULONG_FMT, num_key)
		}
		ZEND_PUTS("] => ")
		ZendPrintFlatZvalR(tmp)
	}
}
func ZendMakePrintableZval(expr *types.Zval, expr_copy *types.Zval) int {
	if expr.IsString() {
		return 0
	} else {
		expr_copy.SetString(ZvalGetStringFunc(expr))
		return 1
	}
}
func ZendPrintZval(expr *types.Zval, indent int) int {
	var tmp_str *types.ZendString
	var str *types.ZendString = ZvalGetTmpString(expr, &tmp_str)
	var len_ int = str.GetLen()
	if len_ != 0 {
		ZendWrite(str.GetStr())
	}
	ZendTmpStringRelease(tmp_str)
	return len_
}
func ZendPrintFlatZvalR(expr *types.Zval) {
	switch expr.GetType() {
	case types.IS_ARRAY:
		ZEND_PUTS("Array (")
		if (expr.GetArr().GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			if expr.GetArr().IsRecursive() {
				ZEND_PUTS(" *RECURSION*")
				return
			}
			expr.GetArr().ProtectRecursive()
		}
		PrintFlatHash(expr.GetArr())
		ZEND_PUTS(")")
		if (expr.GetArr().GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			expr.GetArr().UnprotectRecursive()
		}
		break
	case types.IS_OBJECT:
		var properties *types.HashTable
		var class_name *types.ZendString = types.Z_OBJ_HT(*expr).GetGetClassName()(expr.GetObj())
		ZendPrintf("%s Object (", class_name.GetVal())
		types.ZendStringReleaseEx(class_name, 0)
		if expr.GetCounted().IsRecursive() {
			ZEND_PUTS(" *RECURSION*")
			return
		}
		properties = types.Z_OBJPROP_P(expr)
		if properties != nil {
			expr.GetObj().ProtectRecursive()
			PrintFlatHash(properties)
			expr.GetObj().UnprotectRecursive()
		}
		ZEND_PUTS(")")
		break
	case types.IS_REFERENCE:
		ZendPrintFlatZvalR(types.Z_REFVAL_P(expr))
		break
	default:
		ZendPrintZval(expr, 0)
		break
	}
}
func ZendPrintZvalRToBuf(buf *SmartStr, expr *types.Zval, indent int) {
	switch expr.GetType() {
	case types.IS_ARRAY:
		buf.AppendString("Array\n")
		if (expr.GetArr().GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			if expr.GetArr().IsRecursive() {
				buf.AppendString(" *RECURSION*")
				return
			}
			expr.GetArr().ProtectRecursive()
		}
		PrintHash(buf, expr.GetArr(), indent, 0)
		if (expr.GetArr().GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			expr.GetArr().UnprotectRecursive()
		}
		break
	case types.IS_OBJECT:
		var properties *types.HashTable
		var class_name *types.ZendString = types.Z_OBJ_HT(*expr).GetGetClassName()(expr.GetObj())
		buf.AppendString(class_name.GetStr())
		types.ZendStringReleaseEx(class_name, 0)
		buf.AppendString(" Object\n")
		if expr.GetObj().IsRecursive() {
			buf.AppendString(" *RECURSION*")
			return
		}
		if b.Assign(&properties, ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_DEBUG)) == nil {
			break
		}
		expr.GetObj().ProtectRecursive()
		PrintHash(buf, properties, indent, 1)
		expr.GetObj().UnprotectRecursive()
		ZendReleaseProperties(properties)
		break
	case types.IS_LONG:
		buf.AppendLong(expr.GetLval())
		break
	case types.IS_REFERENCE:
		ZendPrintZvalRToBuf(buf, types.Z_REFVAL_P(expr), indent)
		break
	case types.IS_STRING:
		buf.AppendString(expr.GetStr().GetStr())
		break
	default:
		var str *types.ZendString = ZvalGetStringFunc(expr)
		buf.AppendString(str.GetStr())
		types.ZendStringReleaseEx(str, 0)
		break
	}
}
func ZendPrintZvalRToStr(expr *types.Zval, indent int) *types.ZendString {
	var buf SmartStr = SmartStr{}
	ZendPrintZvalRToBuf(&buf, expr, indent)
	buf.ZeroTail()
	return buf.GetS()
}
func ZendPrintZvalR(expr *types.Zval, indent int) {
	var str *types.ZendString = ZendPrintZvalRToStr(expr, indent)
	ZendWrite(str.GetStr())
	types.ZendStringReleaseEx(str, 0)
}
func ZendFopenWrapper(filename string, opened_path *string) *r.FILE {
	if opened_path != nil {
		*opened_path = filename
	}
	return r.Fopen(filename, "rb")
}
func ZendSetDefaultCompileTimeValues() {
	/* default compile-time values */

	CG__().SetShortTags(ShortTagsDefault)
	CG__().SetCompilerOptions(CompilerOptionsDefault)
	CG__().SetRtdKeyCounter(0)
}
func ZendInitExceptionOp() {
	memset(EG__().GetExceptionOp(), 0, b.SizeOf("EG ( exception_op )"))
	EG__().GetExceptionOp()[0].SetOpcode(ZEND_HANDLE_EXCEPTION)
	ZendVmSetOpcodeHandler(EG__().GetExceptionOp())
	EG__().GetExceptionOp()[1].SetOpcode(ZEND_HANDLE_EXCEPTION)
	ZendVmSetOpcodeHandler(EG__().GetExceptionOp() + 1)
	EG__().GetExceptionOp()[2].SetOpcode(ZEND_HANDLE_EXCEPTION)
	ZendVmSetOpcodeHandler(EG__().GetExceptionOp() + 2)
}
func ZendInitCallTrampolineOp() {
	memset(EG__().GetCallTrampolineOp(), 0, b.SizeOf("EG ( call_trampoline_op )"))
	EG__().GetCallTrampolineOp().SetOpcode(ZEND_CALL_TRAMPOLINE)
	ZendVmSetOpcodeHandler(EG__().GetCallTrampolineOp())
}
func AutoGlobalDtor(zv *types.Zval) { Free(zv.GetPtr()) }
func IniScannerGlobalsCtor(scanner_globals_p *ZendIniScannerGlobals) {
	memset(scanner_globals_p, 0, b.SizeOf("* scanner_globals_p"))
}
func PhpScannerGlobalsCtor(scanner_globals_p *ZendPhpScannerGlobals) {
	memset(scanner_globals_p, 0, b.SizeOf("* scanner_globals_p"))
}
func ModuleDestructorZval(zv *types.Zval) {
	var module *ZendModuleEntry = (*ZendModuleEntry)(zv.GetPtr())
	ModuleDestructor(module)
	Free(module)
}
func PhpAutoGlobalsCreateGlobals(name *types.ZendString) types.ZendBool {
	var globals types.Zval

	/* IS_ARRAY, but with ref-counter 1 and not IS_TYPE_REFCOUNTED */

	globals.SetArray(EG__().GetSymbolTable())
	globals.SetTypeFlags(0)
	globals.SetNewRef(&globals)
	EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &globals)
	return 0
}
func ZendStartup(utility_functions *ZendUtilityFunctions) int {
	var ini_scanner_globals ZendIniScannerGlobals
	var language_scanner_globals ZendPhpScannerGlobals
	//ZendCpuStartup()
	StartMemoryManager()
	VirtualCwdStartup()
	//ZendStartupStrtod()
	ZendStartupExtensionsMechanism()

	/* Set up utility functions and values */

	ZendErrorCb = utility_functions.ErrorFunction
	ZendPrintf = utility_functions.PrintfFunction
	ZendWrite = utility_functions.WriteFunction
	ZendFopen = utility_functions.FopenFunction
	if ZendFopen == nil {
		ZendFopen = ZendFopenWrapper
	}
	ZendStreamOpenFunction = utility_functions.StreamOpenFunction
	ZendMessageDispatcherP = utility_functions.MessageHandler
	ZendGetConfigurationDirectiveP = utility_functions.GetConfigurationDirective
	ZendTicksFunction = utility_functions.TicksFunction
	ZendOnTimeout = utility_functions.OnTimeout
	ZendPrintfToSmartStr = utility_functions.PrintfToSmartStrFunction
	ZendResolvePath = utility_functions.ResolvePathFunction
	ZendInterruptFunction = nil
	ZendCompileFile = CompileFile
	ZendExecuteEx = ExecuteEx
	ZendExecuteInternal = nil
	ZendCompileString = CompileString
	faults.ZendThrowExceptionHook = nil

	/* Set up the default garbage collection implementation. */

	//GcCollectCycles = ZendGcCollectCycles
	ZendVmInit()

	/* set up version */

	ZendVersionInfo = ZEND_CORE_VERSION_INFO

	CG__().InitTables()
	EG__().InitTables()
	ModuleRegistry = *types.NewZendArrayEx(32, ModuleDestructorZval, true)

	ZendInitRsrcListDtors()
	IniScannerGlobalsCtor(&ini_scanner_globals)
	PhpScannerGlobalsCtor(&language_scanner_globals)
	ZendSetDefaultCompileTimeValues()

	/* Map region is going to be created and resized at run-time. */

	CG__().SetMapPtrBase(nil)
	CG__().SetMapPtrSize(0)
	CG__().SetMapPtrLast(0)
	EG__().SetErrorReporting(faults.E_ALL & ^faults.E_NOTICE)
	types.ZendInternedStringsInit()
	ZendStartupBuiltinFunctions()
	ZendRegisterStandardConstants()
	ZendRegisterAutoGlobal(types.ZendStringInitInterned("GLOBALS", b.SizeOf("\"GLOBALS\"")-1, 1), 1, PhpAutoGlobalsCreateGlobals)
	ZendInitRsrcPlist()
	ZendInitExceptionOp()
	ZendInitCallTrampolineOp()
	ZendIniStartup()
	return types.SUCCESS
}
func ZendRegisterStandardIniEntries() {
	REGISTER_INI_ENTRIES(0)
}
func ZendResolvePropertyTypes() {
	var ce *types.ClassEntry
	var prop_info *ZendPropertyInfo
	var __ht *types.HashTable = CG__().GetClassTable()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		ce = _z.GetPtr()
		if ce.GetType() != ZEND_INTERNAL_CLASS {
			continue
		}
		if ZEND_CLASS_HAS_TYPE_HINTS(ce) {
			var __ht *types.HashTable = ce.GetPropertiesInfo()
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()

				prop_info = _z.GetPtr()
				if prop_info.GetType().IsName() {
					var type_name *types.ZendString = prop_info.GetType().Name()
					var lc_type_name *types.ZendString = ZendStringTolower(type_name)
					var prop_ce *types.ClassEntry = types.ZendHashFindPtr(CG__().GetClassTable(), lc_type_name)
					b.Assert(prop_ce != nil && prop_ce.GetType() == ZEND_INTERNAL_CLASS)
					prop_info.SetType(types.ZEND_TYPE_ENCODE_CE(prop_ce, prop_info.GetType().AllowNull()))
					types.ZendStringRelease(lc_type_name)
					types.ZendStringRelease(type_name)
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
		if cb() != types.SUCCESS {
			return types.FAILURE
		}
	}
	GlobalMapPtrLast = CG__().GetMapPtrLast()
	return types.SUCCESS
}
func ZendShutdown() {
	ZendVmDtor()
	EG__().GetPersistentList().GracefulReverseDestroy()
	ZendDestroyModules()
	VirtualCwdDeactivate()
	VirtualCwdShutdown()

	CG__().DestroyTables()
	EG__().DestroyTables()

	ZendShutdownExtensions()

	if CG__().GetMapPtrBase() {
		Free(CG__().GetMapPtrBase())
		CG__().SetMapPtrBase(nil)
		CG__().SetMapPtrSize(0)
	}
	ZendDestroyRsrcListDtors()
}
func ZendSetUtilityValues(utility_values *ZendUtilityValues) { ZendUv = *utility_values }
func Zenderror(error *byte) {
	CG__().SetParseError(0)
	if EG__().GetException() != nil {

		/* An exception was thrown in the lexer, don't throw another in the parser. */

		return

		/* An exception was thrown in the lexer, don't throw another in the parser. */

	}
	faults.ThrowException(faults.ZendCeParseError, error, 0)
}
func ZendAppendVersionInfo(extension *ZendExtension) {
	ZendVersionInfo += fmt.Sprintf("    with %s v%s, %s, by %s\n", extension.GetNameStr(), extension.GetVersionStr(), extension.GetCopyrightStr(), extension.GetAuthorStr())
}
func GetZendVersion() string { return ZendVersionInfo }
func ZendActivate() {
	//GcReset()
	InitCompiler()
	InitExecutor()
	StartupScanner()
	if CG__().GetMapPtrLast() != 0 {
		memset(CG__().GetMapPtrBase(), 0, CG__().GetMapPtrLast()*b.SizeOf("void *"))
	}
}
func ZendCallDestructors() {
	var __orig_bailout *JMP_BUF = EG__().GetBailout()
	var __bailout JMP_BUF
	EG__().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ShutdownDestructors()
	}
	EG__().SetBailout(__orig_bailout)
}
func ZendDeactivate() {
	/* we're no longer executing anything */

	EG__().SetCurrentExecuteData(nil)
	var __orig_bailout *JMP_BUF = EG__().bailout
	var __bailout JMP_BUF
	EG__().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ShutdownScanner()
	}
	EG__().SetBailout(__orig_bailout)

	/* shutdown_executor() takes care of its own bailout handling */

	ShutdownExecutor()
	var __orig_bailout *JMP_BUF = EG__().bailout
	var __bailout JMP_BUF
	EG__().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ZendIniDeactivate()
	}
	EG__().SetBailout(__orig_bailout)
	var __orig_bailout *JMP_BUF = EG__().GetBailout()
	var __bailout JMP_BUF
	EG__().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ShutdownCompiler()
	}
	EG__().SetBailout(__orig_bailout)
	EG__().GetRegularList().GracefulReverseDestroy()
}
func ZendMessageDispatcher(message ZendLong, data any) {
	if ZendMessageDispatcherP != nil {
		ZendMessageDispatcherP(message, data)
	}
}
func ZendGetConfigurationDirective(name *types.ZendString) *types.Zval {
	if ZendGetConfigurationDirectiveP != nil {
		return ZendGetConfigurationDirectiveP(name.GetStr())
	} else {
		return nil
	}
}
func SAVE_STACK(stack ZendStack) {
	if CG__().stack.top {
		memcpy(&stack, CG__().stack, b.SizeOf("zend_stack"))
		CG__().stack.max = 0
		CG__().stack.top = CG__().stack.max
		CG__().stack.elements = nil
	} else {
		stack.SetTop(0)
	}
}
func RESTORE_STACK(stack ZendStack) {
	if stack.GetTop() != 0 {
		ZendStackDestroy(CG__().stack)
		memcpy(CG__().stack, &stack, b.SizeOf("zend_stack"))
	}
}
func ZendUserExceptionHandler() {
	var orig_user_exception_handler types.Zval
	var params []types.Zval
	var retval2 types.Zval
	var old_exception *types.ZendObject
	old_exception = EG__().GetException()
	EG__().SetException(nil)
	params[0].SetObject(old_exception)
	types.ZVAL_COPY_VALUE(&orig_user_exception_handler, EG__().GetUserExceptionHandler())
	if CallUserFunction(nil, &orig_user_exception_handler, &retval2, 1, params) == types.SUCCESS {
		ZvalPtrDtor(&retval2)
		if EG__().GetException() != nil {
			OBJ_RELEASE(EG__().GetException())
			EG__().SetException(nil)
		}
		OBJ_RELEASE(old_exception)
	} else {
		EG__().SetException(old_exception)
	}
}
func ZendExecuteScripts(type_ int, retval *types.Zval, file_count int, _ ...any) int {
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
			types.ZendHashAddEmptyElement(EG__().GetIncludedFiles(), file_handle.GetOpenedPath())
		}
		ZendDestroyFileHandle(file_handle)
		if op_array != nil {
			ZendExecute(op_array, retval)
			faults.ExceptionRestore()
			if EG__().GetException() != nil {
				if EG__().GetUserExceptionHandler().GetType() != types.IS_UNDEF {
					ZendUserExceptionHandler()
				}
				if EG__().GetException() != nil {
					faults.ExceptionError(EG__().GetException(), faults.E_ERROR)
				}
			}
			DestroyOpArray(op_array)
			EfreeSize(op_array, b.SizeOf("zend_op_array"))
		} else if type_ == ZEND_REQUIRE {
			va_end(files)
			return types.FAILURE
		}
	}
	va_end(files)
	return types.SUCCESS
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
func ZendMapPtrNew() any {
	var ptr *any
	if CG__().GetMapPtrLast() >= CG__().GetMapPtrSize() {

		/* Grow map_ptr table */

		CG__().SetMapPtrSize(ZEND_MM_ALIGNED_SIZE_EX(CG__().GetMapPtrLast()+1, 4096))
		CG__().SetMapPtrBase(Perealloc(CG__().GetMapPtrBase(), CG__().GetMapPtrSize()*b.SizeOf("void *"), 1))
	}
	ptr = (*any)(CG__().GetMapPtrBase() + CG__().GetMapPtrLast())
	*ptr = nil
	CG__().GetMapPtrLast()++
	return ZEND_MAP_PTR_PTR2OFFSET(ptr)
}
