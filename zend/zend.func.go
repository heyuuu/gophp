package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func USED_RET(executeData *ZendExecuteData) bool {
	prev := executeData.GetPrevExecuteData()
	if prev == nil {
		return true
	}
	if !(ZEND_USER_CODE(prev.func_.GetType())) {
		return true
	}

	return prev.opline.GetResultType() != IS_UNUSED
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
	val := StrToLongWithUnit(*new_value)
	if stage != ZEND_INI_STAGE_STARTUP && stage != ZEND_INI_STAGE_SHUTDOWN && assertions != val && (assertions < 0 || val < 0) {
		faults.Error(faults.E_WARNING, "zend.assertions may be completely enabled or disabled only in php.ini")
		return false
	}
	EG__().assertions = val
	return true
}

func PrintHash(buf *SmartStr, ht *types.Array, indent int, is_object types.ZendBool) {
	for i := 0; i < indent; i++ {
		buf.AppendByte(' ')
	}
	buf.AppendString("(\n")
	indent += PRINT_ZVAL_INDENT
	ht.ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
		for i := 0; i < indent; i++ {
			buf.AppendByte(' ')
		}
		buf.AppendByte('[')
		if key.IsStrKey() {
			if is_object != 0 {
				className, propName, mangled := ZendUnmanglePropertyName_Ex(key.StrKey())
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
				buf.AppendString(key.StrKey())
			}
		} else {
			buf.AppendLong(key.IdxKey())
		}
		buf.AppendString("] => ")
		ZendPrintZvalRToBuf(buf, value, indent+PRINT_ZVAL_INDENT)
		buf.AppendString("\n")
	})
	indent -= PRINT_ZVAL_INDENT
	for i := 0; i < indent; i++ {
		buf.AppendByte(' ')
	}
	buf.AppendString(")\n")
}
func PrintFlatHash(ht *types.Array) {
	var tmp *types.Zval
	var string_key *types.String
	var num_key ZendUlong
	var i int = 0
	var __ht *types.Array = ht
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
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
func PrintZval(expr *types.Zval) string {
	if expr.IsString() {
		return expr.StringVal()
	} else {
		return operators.ZvalGetString(expr).GetStr()
	}
}

func ZendMakePrintableZval(expr *types.Zval, expr_copy *types.Zval) int {
	if expr.IsString() {
		return 0
	} else {
		expr_copy.SetString(operators.ZvalGetString(expr))
		return 1
	}
}
func ZendPrintZval(expr *types.Zval) int {
	var str = operators.ZvalGetStrVal(expr)
	ZendWrite(str)
	return len(str)
}
func ZendPrintFlatZvalR(expr *types.Zval) {
	switch expr.GetType() {
	case types.IS_ARRAY:
		ZEND_PUTS("Array (")
		if (expr.Array().GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			if expr.Array().IsRecursive() {
				ZEND_PUTS(" *RECURSION*")
				return
			}
			expr.Array().ProtectRecursive()
		}
		PrintFlatHash(expr.Array())
		ZEND_PUTS(")")
		if (expr.Array().GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			expr.Array().UnprotectRecursive()
		}
		break
	case types.IS_OBJECT:
		var properties *types.Array
		ZendPrintf("%s Object (", expr.Object().ClassName())
		// types.ZendStringReleaseEx(class_name, 0)
		if expr.RefCounted().IsRecursive() {
			ZEND_PUTS(" *RECURSION*")
			return
		}
		properties = types.Z_OBJPROP_P(expr)
		if properties != nil {
			expr.Object().ProtectRecursive()
			PrintFlatHash(properties)
			expr.Object().UnprotectRecursive()
		}
		ZEND_PUTS(")")
		break
	case types.IS_REFERENCE:
		ZendPrintFlatZvalR(types.Z_REFVAL_P(expr))
		break
	default:
		ZendPrintZval(expr)
		break
	}
}
func ZendPrintZvalRToBuf(buf *SmartStr, expr *types.Zval, indent int) {
	switch expr.GetType() {
	case types.IS_ARRAY:
		buf.AppendString("Array\n")
		if (expr.Array().GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			if expr.Array().IsRecursive() {
				buf.AppendString(" *RECURSION*")
				return
			}
			expr.Array().ProtectRecursive()
		}
		PrintHash(buf, expr.Array(), indent, 0)
		if (expr.Array().GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			expr.Array().UnprotectRecursive()
		}
		break
	case types.IS_OBJECT:
		var properties *types.Array
		buf.AppendString(expr.Object().ClassName())
		buf.AppendString(" Object\n")
		if expr.Object().IsRecursive() {
			buf.AppendString(" *RECURSION*")
			return
		}
		if b.Assign(&properties, ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_DEBUG)) == nil {
			break
		}
		expr.Object().ProtectRecursive()
		PrintHash(buf, properties, indent, 1)
		expr.Object().UnprotectRecursive()
		ZendReleaseProperties(properties)
		break
	case types.IS_LONG:
		buf.AppendLong(expr.Long())
		break
	case types.IS_REFERENCE:
		ZendPrintZvalRToBuf(buf, types.Z_REFVAL_P(expr), indent)
		break
	case types.IS_STRING:
		buf.AppendString(expr.String().GetStr())
		break
	default:
		var str *types.String = operators.ZvalGetString(expr)
		buf.AppendString(str.GetStr())
		// types.ZendStringReleaseEx(str, 0)
		break
	}
}
func ZendPrintZvalRToStr(expr *types.Zval, indent int) *types.String {
	var buf SmartStr = SmartStr{}
	ZendPrintZvalRToBuf(&buf, expr, indent)
	buf.ZeroTail()
	return buf.GetS()
}
func ZendPrintZvalR(expr *types.Zval, indent int) {
	var str *types.String = ZendPrintZvalRToStr(expr, indent)
	ZendWrite(str.GetStr())
	// types.ZendStringReleaseEx(str, 0)
}
func ZendSetDefaultCompileTimeValues() {
	/* default compile-time values */

	CG__().SetShortTags(ShortTagsDefault)
	CG__().SetCompilerOptions(CompilerOptionsDefault)
	CG__().SetRtdKeyCounter(0)
}
func ZendInitExceptionOp() {
	exceptionOps := EG__().GetExceptionOp()
	*exceptionOps = [3]ZendOp{}
	exceptionOps[0].SetOpcode(ZEND_HANDLE_EXCEPTION)
	ZendVmSetOpcodeHandler(&exceptionOps[0])
	exceptionOps[1].SetOpcode(ZEND_HANDLE_EXCEPTION)
	ZendVmSetOpcodeHandler(&exceptionOps[1])
	exceptionOps[2].SetOpcode(ZEND_HANDLE_EXCEPTION)
	ZendVmSetOpcodeHandler(&exceptionOps[2])
}
func ZendInitCallTrampolineOp() {
	memset(EG__().GetCallTrampolineOp(), 0, b.SizeOf("EG ( call_trampoline_op )"))
	EG__().GetCallTrampolineOp().SetOpcode(ZEND_CALL_TRAMPOLINE)
	ZendVmSetOpcodeHandler(EG__().GetCallTrampolineOp())
}
func IniScannerGlobalsCtor(scanner_globals_p *ZendIniScannerGlobals) {
	memset(scanner_globals_p, 0, b.SizeOf("* scanner_globals_p"))
}
func PhpScannerGlobalsCtor(scanner_globals_p *ZendPhpScannerGlobals) {
	memset(scanner_globals_p, 0, b.SizeOf("* scanner_globals_p"))
}
func PhpAutoGlobalsCreateGlobals(name *types.String) types.ZendBool {
	var globals types.Zval

	/* IS_ARRAY, but with ref-counter 1 and not IS_TYPE_REFCOUNTED */

	globals.SetImmutableArray(EG__().GetSymbolTable())
	globals.SetNewRef(&globals)
	EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &globals)
	return 0
}
func ZendStartup() int {
	var ini_scanner_globals ZendIniScannerGlobals
	var language_scanner_globals ZendPhpScannerGlobals
	//ZendCpuStartup()
	StartMemoryManager()
	VirtualCwdStartup()
	//ZendStartupStrtod()
	ZendStartupExtensionsMechanism()

	ZendVersionInfo = ZEND_CORE_VERSION_INFO

	CG__().InitTables()
	EG__().InitTables()

	ZendInitRsrcListDtors()
	IniScannerGlobalsCtor(&ini_scanner_globals)
	PhpScannerGlobalsCtor(&language_scanner_globals)
	ZendSetDefaultCompileTimeValues()

	/* Map region is going to be created and resized at run-time. */
	CG__().SetMapPtrBase(nil)
	CG__().SetMapPtrSize(0)
	CG__().SetMapPtrLast(0)
	EG__().SetErrorReporting(faults.E_ALL & ^faults.E_NOTICE)
	ZendStartupBuiltinFunctions()
	ZendRegisterStandardConstants()
	ZendRegisterAutoGlobal(types.NewString("GLOBALS"), 1, PhpAutoGlobalsCreateGlobals)
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
	CG__().ClassTable().Foreach(func(_ string, ce *types.ClassEntry) {
		if ce.GetType() != ZEND_INTERNAL_CLASS {
			return
		}
		if ZEND_CLASS_HAS_TYPE_HINTS(ce) {
			ce.PropertyTable().Foreach(func(key string, prop_info *ZendPropertyInfo) {
				if prop_info.GetType().IsName() {
					var type_name = prop_info.GetType().Name().GetStr()
					var prop_ce *types.ClassEntry = CG__().ClassTable().Get(type_name)
					b.Assert(prop_ce != nil && prop_ce.GetType() == ZEND_INTERNAL_CLASS)
					prop_info.SetType(types.ZEND_TYPE_ENCODE_CE(prop_ce, prop_info.GetType().AllowNull()))
				}
			})
		}
		ce.SetIsPropertyTypesResolved(true)
	})
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
	EG__().PersistentList().DestroyReverse()

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
	faults.Try(func() {
		ShutdownDestructors()
	})
}
func ZendDeactivate() {
	/* we're no longer executing anything */
	EG__().SetCurrentExecuteData(nil)

	faults.Try(func() { ShutdownScanner() })

	/* shutdown_executor() takes care of its own bailout handling */
	ShutdownExecutor()

	faults.Try(func() { ZendIniDeactivate() })
	faults.Try(func() { ShutdownCompiler() })

	EG__().RegularList().DestroyReverse()
}
func ZendMessageDispatcher(message ZendLong, data any) {
	if ZendMessageDispatcherP != nil {
		ZendMessageDispatcherP(message, data)
	}
}
func ZendGetConfigurationDirective(name *types.String) *types.Zval {
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
		// ZvalPtrDtor(&retval2)
		if EG__().GetException() != nil {
			// OBJ_RELEASE(EG__().GetException())
			EG__().SetException(nil)
		}
		// OBJ_RELEASE(old_exception)
	} else {
		EG__().SetException(old_exception)
	}
}
func ZendExecuteScriptsEx(typ int, retval *types.Zval, files ...*FileHandle) bool {
	for _, fileHandle := range files {
		if fileHandle == nil {
			continue
		}
		opArray := CompileFile(fileHandle, typ)
		if fileHandle.GetOpenedPath() != "" {
			types.ZendHashAddEmptyElement(EG__().GetIncludedFiles(), fileHandle.GetOpenedPath())
		}
		ZendDestroyFileHandle(fileHandle)
		if opArray != nil {
			ZendExecute(opArray, retval)
			faults.ExceptionRestore()
			if EG__().GetException() != nil {
				if EG__().GetUserExceptionHandler().IsNotUndef() {
					ZendUserExceptionHandler()
				}
				if EG__().GetException() != nil {
					faults.ExceptionError(EG__().GetException(), faults.E_ERROR)
				}
			}
			DestroyOpArray(opArray)
		} else if typ == ZEND_REQUIRE {
			return false
		}
	}
	return true
}

func ZendMakeCompiledStringDescription(name string) string {
	var cur_filename *byte
	var cur_lineno int
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
	return ZendSprintf(COMPILED_STRING_DESCRIPTION_FORMAT, cur_filename, cur_lineno, name)
}
func FreeEstring(str_p **byte) { Efree(*str_p) }
func ZendMapPtrNew() any {
	var ptr *any
	if CG__().GetMapPtrLast() >= CG__().GetMapPtrSize() {

		/* Grow map_ptr table */

		CG__().SetMapPtrSize(ZEND_MM_ALIGNED_SIZE_EX(CG__().GetMapPtrLast()+1, 4096))
		CG__().SetMapPtrBase(Perealloc(CG__().GetMapPtrBase(), CG__().GetMapPtrSize()*b.SizeOf("void *")))
	}
	ptr = (*any)(CG__().GetMapPtrBase() + CG__().GetMapPtrLast())
	*ptr = nil
	CG__().GetMapPtrLast()++
	return ZEND_MAP_PTR_PTR2OFFSET(ptr)
}
