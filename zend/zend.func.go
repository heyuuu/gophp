// <<generate>>

package zend

import (
	"fmt"
	b "sik/builtin"
	r "sik/runtime"
)

func USED_RET() bool {
	return !(EX(prev_execute_data)) || !(ZEND_USER_CODE(EX(prev_execute_data).func_.common.type_)) || EX(prev_execute_data).opline.result_type != IS_UNUSED
}
func ZendBailout()             { _zendBailout(__FILE__, __LINE__) }
func ZEND_PUTS(str string) int { return ZendWrite(str) }
func ZEND_PUTC(c byte) int     { return ZendWrite(string([]byte{c})) }

func OnUpdateErrorReportingEx(entry *ZendIniEntry, newValue *string, stage int) bool {
	if newValue == nil {
		EG__().SetErrorReporting(E_ALL & ^E_NOTICE & ^E_STRICT & ^E_DEPRECATED)
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
		ZendError(E_WARNING, "zend.assertions may be completely enabled or disabled only in php.ini")
		return false
	}
	EG__().assertions = val
	return true
}

func ZendVspprintf(pbuf *string, max_len int, format string, ap ...any) int {
	/* since there are places where (v)spprintf called without checking for null,
	   a bit of defensive coding here */
	if pbuf == nil {
		return 0
	}
	result := __sprintfEx(max_len, format, ap...)
	*pbuf = result
	return len(result)
}

func __sprintfEx(maxLen int, format string, args ...any) string {
	buf := SmartStr{}
	ZendPrintfToSmartStr(&buf, format, args...)
	result := buf.GetStr()
	if maxLen != 0 && len(result) > maxLen {
		return result[:maxLen]
	}
	return result
}

func __sprintf(format string, args ...any) string {
	var buf = SmartStr{}
	ZendPrintfToSmartStr(&buf, format, args...)
	return buf.GetStr()
}

func ZendSpprintf(message *string, max_len int, format string, args ...any) int {
	len_ := ZendVspprintf(message, max_len, format, args...)
	return len_
}
func ZendVstrpprintf(max_len int, format *byte, ap ...any) *ZendString {
	var buf SmartStr
	ZendPrintfToSmartStr(&buf, format, ap)
	if buf.GetS() == nil {
		return ZSTR_EMPTY_ALLOC()
	}
	if max_len != 0 && buf.GetS().GetLen() > max_len {
		buf.GetS().GetLen() = max_len
	}
	buf.ZeroTail()
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
		buf.AppendByte(' ')
	}
	buf.AppendString("(\n")
	indent += PRINT_ZVAL_INDENT
	var __ht *HashTable = ht
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.GetZv()
			if _z.IsUndef() {
				continue
			}
		}
		num_key = _p.GetH()
		string_key = _p.GetKey()
		tmp = _z
		for i = 0; i < indent; i++ {
			buf.AppendByte(' ')
		}
		buf.AppendByte('[')
		if string_key != nil {
			if is_object != 0 {
				var prop_name *byte
				var class_name *byte
				var prop_len int
				var mangled int = ZendUnmanglePropertyNameEx(string_key, &class_name, &prop_name, &prop_len)
				buf.AppendString(b.CastStr(prop_name, prop_len))
				if class_name != nil && mangled == SUCCESS {
					if class_name[0] == '*' {
						buf.AppendString(":protected")
					} else {
						buf.AppendString(":")
						buf.AppendString(b.CastStrAuto(class_name))
						buf.AppendString(":private")
					}
				}
			} else {
				buf.AppendString(string_key.GetStr())
			}
		} else {
			buf.AppendLong(num_key)
		}
		buf.AppendString("] => ")
		ZendPrintZvalRToBuf(buf, tmp, indent+PRINT_ZVAL_INDENT)
		buf.AppendString("\n")
	}
	indent -= PRINT_ZVAL_INDENT
	for i = 0; i < indent; i++ {
		buf.AppendByte(' ')
	}
	buf.AppendString(")\n")
}
func PrintFlatHash(ht *HashTable) {
	var tmp *Zval
	var string_key *ZendString
	var num_key ZendUlong
	var i int = 0
	var __ht *HashTable = ht
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()
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
func ZendMakePrintableZval(expr *Zval, expr_copy *Zval) int {
	if expr.IsString() {
		return 0
	} else {
		expr_copy.SetString(ZvalGetStringFunc(expr))
		return 1
	}
}
func ZendPrintZval(expr *Zval, indent int) int {
	var tmp_str *ZendString
	var str *ZendString = ZvalGetTmpString(expr, &tmp_str)
	var len_ int = str.GetLen()
	if len_ != 0 {
		ZendWrite(str.GetStr())
	}
	ZendTmpStringRelease(tmp_str)
	return len_
}
func ZendPrintFlatZvalR(expr *Zval) {
	switch expr.GetType() {
	case IS_ARRAY:
		ZEND_PUTS("Array (")
		if (expr.GetArr().GetGcFlags() & GC_IMMUTABLE) == 0 {
			if expr.GetArr().IsRecursive() {
				ZEND_PUTS(" *RECURSION*")
				return
			}
			expr.GetArr().ProtectRecursive()
		}
		PrintFlatHash(expr.GetArr())
		ZEND_PUTS(")")
		if (expr.GetArr().GetGcFlags() & GC_IMMUTABLE) == 0 {
			expr.GetArr().UnprotectRecursive()
		}
		break
	case IS_OBJECT:
		var properties *HashTable
		var class_name *ZendString = Z_OBJ_HT(*expr).GetGetClassName()(expr.GetObj())
		ZendPrintf("%s Object (", class_name.GetVal())
		ZendStringReleaseEx(class_name, 0)
		if expr.GetCounted().IsRecursive() {
			ZEND_PUTS(" *RECURSION*")
			return
		}
		properties = Z_OBJPROP_P(expr)
		if properties != nil {
			expr.GetObj().ProtectRecursive()
			PrintFlatHash(properties)
			expr.GetObj().UnprotectRecursive()
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
		buf.AppendString("Array\n")
		if (expr.GetArr().GetGcFlags() & GC_IMMUTABLE) == 0 {
			if expr.GetArr().IsRecursive() {
				buf.AppendString(" *RECURSION*")
				return
			}
			expr.GetArr().ProtectRecursive()
		}
		PrintHash(buf, expr.GetArr(), indent, 0)
		if (expr.GetArr().GetGcFlags() & GC_IMMUTABLE) == 0 {
			expr.GetArr().UnprotectRecursive()
		}
		break
	case IS_OBJECT:
		var properties *HashTable
		var class_name *ZendString = Z_OBJ_HT(*expr).GetGetClassName()(expr.GetObj())
		buf.AppendString(b.CastStrAuto(class_name.GetVal()))
		ZendStringReleaseEx(class_name, 0)
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
	case IS_LONG:
		buf.AppendLong(expr.GetLval())
		break
	case IS_REFERENCE:
		ZendPrintZvalRToBuf(buf, Z_REFVAL_P(expr), indent)
		break
	case IS_STRING:
		buf.AppendString(expr.GetStr().GetStr())
		break
	default:
		var str *ZendString = ZvalGetStringFunc(expr)
		buf.AppendString(str.GetStr())
		ZendStringReleaseEx(str, 0)
		break
	}
}
func ZendPrintZvalRToStr(expr *Zval, indent int) *ZendString {
	var buf SmartStr = MakeSmartStr(0)
	ZendPrintZvalRToBuf(&buf, expr, indent)
	buf.ZeroTail()
	return buf.GetS()
}
func ZendPrintZvalR(expr *Zval, indent int) {
	var str *ZendString = ZendPrintZvalRToStr(expr, indent)
	ZendWrite(str.GetStr())
	ZendStringReleaseEx(str, 0)
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
	ZendGetenv = utility_functions.GetenvFunction
	ZendResolvePath = utility_functions.ResolvePathFunction
	ZendInterruptFunction = nil
	ZendCompileFile = CompileFile
	ZendExecuteEx = ExecuteEx
	ZendExecuteInternal = nil
	ZendCompileString = CompileString
	ZendThrowExceptionHook = nil

	/* Set up the default garbage collection implementation. */

	//GcCollectCycles = ZendGcCollectCycles
	ZendVmInit()

	/* set up version */

	ZendVersionInfo = ZEND_CORE_VERSION_INFO

	CG__().InitTables()
	EG__().InitTables()
	ModuleRegistry = *NewZendArrayEx(32, ModuleDestructorZval, true)

	ZendInitRsrcListDtors()
	IniScannerGlobalsCtor(&ini_scanner_globals)
	PhpScannerGlobalsCtor(&language_scanner_globals)
	ZendSetDefaultCompileTimeValues()

	/* Map region is going to be created and resized at run-time. */

	CG__().SetMapPtrBase(nil)
	CG__().SetMapPtrSize(0)
	CG__().SetMapPtrLast(0)
	EG__().SetErrorReporting(E_ALL & ^E_NOTICE)
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
	var __ht *HashTable = CG__().GetClassTable()
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
					var prop_ce *ZendClassEntry = ZendHashFindPtr(CG__().GetClassTable(), lc_type_name)
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
	GlobalMapPtrLast = CG__().GetMapPtrLast()
	return SUCCESS
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
	ZendThrowException(ZendCeParseError, error, 0)
}
func _zendBailout(filename *byte, lineno uint32) {
	if EG__().GetBailout() == nil {
		ZendOutputDebugString(1, "%s(%d) : Bailed out without a bailout address!", filename, lineno)
		exit(-1)
	}
	GcProtect(1)
	CG__().SetUncleanShutdown(1)
	CG__().SetActiveClassEntry(nil)
	CG__().SetInCompilation(0)
	EG__().SetCurrentExecuteData(nil)
	LONGJMP((*EG__)().bailout, FAILURE)
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
func ZendGetConfigurationDirective(name *ZendString) *Zval {
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
func ZendErrorVaList(type_ int, error_filename *byte, error_lineno uint32, format string, args ...any) {
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

	if EG__().GetException() != nil {
		var ex *ZendExecuteData
		var opline *ZendOp
		switch type_ {
		case E_CORE_ERROR:

		case E_ERROR:

		case E_RECOVERABLE_ERROR:

		case E_PARSE:

		case E_COMPILE_ERROR:

		case E_USER_ERROR:
			ex = EG__().GetCurrentExecuteData()
			opline = nil
			for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
				ex = ex.GetPrevExecuteData()
			}
			if ex != nil && ex.GetOpline().GetOpcode() == ZEND_HANDLE_EXCEPTION && EG__().GetOplineBeforeException() != nil {
				opline = EG__().GetOplineBeforeException()
			}
			ZendExceptionError(EG__().GetException(), E_WARNING)
			EG__().SetException(nil)
			if opline != nil {
				ex.SetOpline(opline)
			}
			break
		default:
			break
		}
	}

	/* if we don't have a user defined error handler */

	if EG__().GetUserErrorHandler().IsUndef() || (EG__().GetUserErrorHandlerErrorReporting()&type_) == 0 || EG__().GetErrorHandling() != EH_NORMAL {
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
			params[1].SetString(ZendVstrpprintf(0, format, usr_copy))
			va_end(usr_copy)
			params[0].SetLong(type_)
			if error_filename != nil {
				ZVAL_STRING(&params[2], error_filename)
			} else {
				params[2].SetNull()
			}
			params[3].SetLong(error_lineno)
			symbol_table = ZendRebuildSymbolTable()

			/* during shutdown the symbol table table can be still null */

			if symbol_table == nil {
				params[4].SetNull()
			} else {
				params[4].SetArray(ZendArrayDup(symbol_table))
			}
			ZVAL_COPY_VALUE(&orig_user_error_handler, EG__().GetUserErrorHandler())
			EG__().GetUserErrorHandler().SetUndef()

			/* User error handler may include() additinal PHP files.
			 * If an error was generated during comilation PHP will compile
			 * such scripts recursively, but some CG() variables may be
			 * inconsistent. */

			in_compilation = CG__().GetInCompilation()
			if in_compilation != 0 {
				saved_class_entry = CG__().GetActiveClassEntry()
				CG__().SetActiveClassEntry(nil)
				SAVE_STACK(loop_var_stack)
				SAVE_STACK(delayed_oplines_stack)
				CG__().SetInCompilation(0)
			}
			orig_fake_scope = EG__().GetFakeScope()
			EG__().SetFakeScope(nil)
			if CallUserFunction(CG__().GetFunctionTable(), nil, &orig_user_error_handler, &retval, 5, params) == SUCCESS {
				if retval.GetType() != IS_UNDEF {
					if retval.IsFalse() {
						ZendErrorCb(type_, error_filename, error_lineno, format, args)
					}
					ZvalPtrDtor(&retval)
				}
			} else if EG__().GetException() == nil {

				/* The user error handler failed, use built-in error handler */

				ZendErrorCb(type_, error_filename, error_lineno, format, args)

				/* The user error handler failed, use built-in error handler */

			}
			EG__().SetFakeScope(orig_fake_scope)
			if in_compilation != 0 {
				CG__().SetActiveClassEntry(saved_class_entry)
				RESTORE_STACK(loop_var_stack)
				RESTORE_STACK(delayed_oplines_stack)
				CG__().SetInCompilation(1)
			}
			ZvalPtrDtor(&params[4])
			ZvalPtrDtor(&params[2])
			ZvalPtrDtor(&params[1])
			if EG__().GetUserErrorHandler().IsUndef() {
				ZVAL_COPY_VALUE(EG__().GetUserErrorHandler(), &orig_user_error_handler)
			} else {
				ZvalPtrDtor(&orig_user_error_handler)
			}
			break
		}
	}
	if type_ == E_PARSE {

		/* eval() errors do not affect exit_status */

		if !(EG__().GetCurrentExecuteData() != nil && EG__().GetCurrentExecuteData().GetFunc() != nil && ZEND_USER_CODE(EG__().GetCurrentExecuteData().GetFunc().GetType()) && EG__().GetCurrentExecuteData().GetOpline().GetOpcode() == ZEND_INCLUDE_OR_EVAL && EG__().GetCurrentExecuteData().GetOpline().GetExtendedValue() == ZEND_EVAL) {
			EG__().SetExitStatus(255)
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
func ZendError(type_ int, format string, args ...any) {
	var filename *byte
	var lineno uint32
	GetFilenameLineno(type_, &filename, &lineno)
	ZendErrorVaList(type_, filename, lineno, format, args)
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
func ZendThrowError(exception_ce *ZendClassEntry, format string, args ...any) {
	if exception_ce != nil {
		if InstanceofFunction(exception_ce, ZendCeError) == 0 {
			ZendError(E_NOTICE, "Error exceptions must be derived from Error")
			exception_ce = ZendCeError
		}
	} else {
		exception_ce = ZendCeError
	}

	/* Marker used to disable exception generation during preloading. */

	if EG__().GetException() == any(uintPtr-1) {
		return
	}

	message := __sprintf(format, args...)

	//TODO: we can't convert compile-time errors to exceptions yet???

	if EG__().GetCurrentExecuteData() != nil && CG__().GetInCompilation() == 0 {
		ZendThrowException(exception_ce, message, 0)
	} else {
		ZendError(E_ERROR, "%s", message)
	}
}
func ZendTypeError(format string, args ...any) {
	message := __sprintf(format, args...)
	ZendThrowException(ZendCeTypeError, message, 0)
	Efree(message)
}
func ZendInternalTypeError(throw_exception ZendBool, format string, args ...any) {
	message := __sprintf(format, args...)
	if throw_exception != 0 {
		ZendThrowException(ZendCeTypeError, message, 0)
	} else {
		ZendError(E_WARNING, "%s", message)
	}
}
func ZendInternalArgumentCountError(throw_exception ZendBool, format string, args ...any) {
	message := __sprintf(format, args...)
	if throw_exception != 0 {
		ZendThrowException(ZendCeArgumentCountError, message, 0)
	} else {
		ZendError(E_WARNING, "%s", message)
	}
}
func ZendOutputDebugString(trigger_break ZendBool, format string, _ ...any) {}
func ZendUserExceptionHandler() {
	var orig_user_exception_handler Zval
	var params []Zval
	var retval2 Zval
	var old_exception *ZendObject
	old_exception = EG__().GetException()
	EG__().SetException(nil)
	params[0].SetObject(old_exception)
	ZVAL_COPY_VALUE(&orig_user_exception_handler, EG__().GetUserExceptionHandler())
	if CallUserFunction(CG__().GetFunctionTable(), nil, &orig_user_exception_handler, &retval2, 1, params) == SUCCESS {
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
			ZendHashAddEmptyElement(EG__().GetIncludedFiles(), file_handle.GetOpenedPath())
		}
		ZendDestroyFileHandle(file_handle)
		if op_array != nil {
			ZendExecute(op_array, retval)
			ZendExceptionRestore()
			if EG__().GetException() != nil {
				if EG__().GetUserExceptionHandler().GetType() != IS_UNDEF {
					ZendUserExceptionHandler()
				}
				if EG__().GetException() != nil {
					ZendExceptionError(EG__().GetException(), E_ERROR)
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
