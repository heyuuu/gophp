package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/kits/strkit"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"strings"
)

func ZendThrowOrError(fetchType int, exceptionCe *types.ClassEntry, message string) {
	if (fetchType & ZEND_FETCH_CLASS_EXCEPTION) != 0 {
		faults.ThrowError(exceptionCe, "%s", message)
	} else {
		faults.Error(faults.E_ERROR, "%s", message)
	}
}
func ShutdownDestructors() {
	faults.TryCatch(func() {
		// notice: 无需主动调用析构函数，使用自动析构代替
		EG__().GetSymbolTable().Clean()
		// notice: 无需主动调用析构函数，使用自动析构代替
		//ZendObjectsStoreCallDestructors(EG__().GetObjectsStore())
	}, func() {
		/* if we couldn't destruct cleanly, mark all objects as destructed anyway */
		// notice: 无需主动调用析构函数，使用自动析构代替
		//ZendObjectsStoreMarkDestructed(EG__().GetObjectsStore())
	})
}

func ZendGetExecutedFilename() string {
	var ex *ZendExecuteData = CurrEX()
	for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
		ex = ex.GetPrevExecuteData()
	}
	if ex != nil {
		return ex.GetFunc().GetOpArray().GetFilename()
	} else {
		return "[no active file]"
	}
}
func ZendGetExecutedFilenameEx() *types.String {
	var ex *ZendExecuteData = CurrEX()
	for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
		ex = ex.GetPrevExecuteData()
	}
	if ex != nil {
		filename := ex.GetFunc().GetOpArray().GetFilename()
		return types.NewString(filename)
	} else {
		return nil
	}
}
func ZendGetExecutedLineno() int {
	var ex *ZendExecuteData = CurrEX()
	for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
		ex = ex.GetPrevExecuteData()
	}
	if ex != nil {
		if EG__().HasException() && ex.GetOpline().GetOpcode() == ZEND_HANDLE_EXCEPTION && ex.GetOpline().GetLineno() == 0 && EG__().GetOplineBeforeException() != nil {
			return EG__().GetOplineBeforeException().GetLineno()
		}
		return ex.GetOpline().GetLineno()
	} else {
		return 0
	}
}
func ZendGetExecutedScope() *types.ClassEntry {
	var ex *ZendExecuteData = CurrEX()
	for {
		if ex == nil {
			return nil
		} else if ex.GetFunc() != nil && (ZEND_USER_CODE(ex.GetFunc().GetType()) || ex.GetFunc().GetScope() != nil) {
			return ex.GetFunc().GetScope()
		}
		ex = ex.GetPrevExecuteData()
	}
}
func ZendIsExecuting() bool {
	return CurrEX() != nil
}
func ZendUseUndefinedConstant(name string, attr ZendAstAttr, result *types.Zval) bool {
	if EG__().HasException() {
		return false
	} else if strings.ContainsRune(name, ':') {
		faults.ThrowError(nil, fmt.Sprintf("Undefined class constant '%s'", name))
		return false
	} else if (attr & IS_CONSTANT_UNQUALIFIED) == 0 {
		faults.ThrowError(nil, fmt.Sprintf("Undefined constant '%s'", name))
		return false
	} else {
		_, actualStr, _ := strkit.LastCut(name, "\\")
		faults.Error(faults.E_WARNING, fmt.Sprintf("Use of undefined constant %s - assumed '%s' (this will throw an Error in a future version of PHP)", actualStr, actualStr))
		if EG__().HasException() {
			return false
		} else {
			result.SetString(actualStr)
		}
	}
	return true
}
func ZvalUpdateConstantEx(p *types.Zval, scope *types.ClassEntry) bool {
	if p.IsConstantAst() {
		var ast *ZendAst = types.Z_ASTVAL_P(p)
		if ast.Kind() == ZEND_AST_CONSTANT {
			var name = ZendAstGetConstantName(ast)
			var zv *types.Zval = ZendGetConstantEx(name, scope, ast.Attr())
			if zv == nil {
				return ZendUseUndefinedConstant(name, ast.Attr(), p)
			}
			types.ZVAL_COPY_OR_DUP(p, zv)
		} else {
			var tmp types.Zval
			if ZendAstEvaluate(&tmp, ast, scope) != types.SUCCESS {
				return false
			}
			types.ZVAL_COPY_VALUE(p, &tmp)
		}
	}
	return true
}
func ZendCallFunction(fci *types.ZendFcallInfo, fciCache *types.ZendFcallInfoCache) int {
	var i uint32
	var call *ZendExecuteData
	var dummy_execute_data ZendExecuteData
	var fci_cache_local types.ZendFcallInfoCache
	var func_ types.IFunction
	var call_info uint32
	var object_or_called_scope any
	fci.GetRetval().SetUndef()
	if !EG__().IsActive() {
		return types.FAILURE
	}
	if EG__().HasException() {
		return types.FAILURE
	}
	b.Assert(fci.IsInit())

	/* Initialize executeData */

	if CurrEX() == nil {

		/* This only happens when we're called outside any execute()'s
		 * It shouldn't be strictly necessary to NULL executeData out,
		 * but it may make bugs easier to spot
		 */
		memset(&dummy_execute_data, 0, b.SizeOf("zend_execute_data"))
		EG__().SetCurrentExecuteData(&dummy_execute_data)
	} else if CurrEX().GetFunc() != nil && ZEND_USER_CODE(CurrEX().GetFunc().GetType()) && CurrEX().GetOpline().GetOpcode() != ZEND_DO_FCALL && CurrEX().GetOpline().GetOpcode() != ZEND_DO_ICALL && CurrEX().GetOpline().GetOpcode() != ZEND_DO_UCALL && CurrEX().GetOpline().GetOpcode() != ZEND_DO_FCALL_BY_NAME {

		/* Insert fake frame in case of include or magic calls */

		dummy_execute_data = *EG__().GetCurrentExecuteData()
		dummy_execute_data.SetPrevExecuteData(CurrEX())
		dummy_execute_data.SetCall(nil)
		dummy_execute_data.SetOpline(nil)
		dummy_execute_data.SetFunc(nil)
		EG__().SetCurrentExecuteData(&dummy_execute_data)
	}
	if fciCache == nil || fciCache.GetFunctionHandler() == nil {
		var error *byte = nil
		if fciCache == nil {
			fciCache = &fci_cache_local
		}
		if !ZendIsCallableEx(fci.GetFunctionName(), fci.GetObject(), IS_CALLABLE_CHECK_SILENT, nil, fciCache, &error) {
			if error != nil {
				var callableName = GetCallableName(fci.GetFunctionName(), fci.GetObject())
				faults.Error(faults.E_WARNING, "Invalid callback %s, %s", callableName, error)
				Efree(error)
			}
			if CurrEX() == &dummy_execute_data {
				EG__().SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
			}
			return types.FAILURE
		} else if error != nil {

			/* Capitalize the first latter of the error message */

			if error[0] >= 'a' && error[0] <= 'z' {
				error[0] += 'A' - 'a'
			}
			faults.Error(faults.E_DEPRECATED, "%s", error)
			Efree(error)
			if EG__().HasException() {
				if CurrEX() == &dummy_execute_data {
					EG__().SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
				}
				return types.FAILURE
			}
		}
	}
	func_ = fciCache.GetFunctionHandler()
	if func_.IsStatic() || fciCache.GetObject() == nil {
		fci.SetObject(nil)
		object_or_called_scope = fciCache.GetCalledScope()
		call_info = ZEND_CALL_TOP_FUNCTION | ZEND_CALL_DYNAMIC
	} else {
		fci.SetObject(fciCache.GetObject())
		object_or_called_scope = fci.GetObject()
		call_info = ZEND_CALL_TOP_FUNCTION | ZEND_CALL_DYNAMIC | ZEND_CALL_HAS_THIS
	}
	call = ZendVmStackPushCallFrame(call_info, func_, fci.GetParamCount(), object_or_called_scope)
	if func_.IsDeprecated() {
		faults.Error(faults.E_DEPRECATED, "Function %s%s%s() is deprecated", lang.CondF1(func_.GetScope() != nil, func() []byte { return func_.GetScope().Name() }, ""), lang.Cond(func_.GetScope() != nil, "::", ""), func_.FunctionName())
		if EG__().HasException() {
			ZendVmStackFreeCallFrame(call)
			if CurrEX() == &dummy_execute_data {
				EG__().SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
				faults.RethrowException(CurrEX())
			}
			return types.FAILURE
		}
	}
	for i = 0; i < fci.GetParamCount(); i++ {
		var param *types.Zval
		var arg *types.Zval = fci.GetParams()[i]
		var must_wrap bool = 0
		if ARG_SHOULD_BE_SENT_BY_REF(func_, i+1) != 0 {
			if !(arg.IsRef()) {
				if fci.GetNoSeparation() == 0 {
					/* Separation is enabled -- create a ref */
					arg.SetNewRef(arg)
				} else if ARG_MAY_BE_SENT_BY_REF(func_, i+1) == 0 {

					/* By-value send is not allowed -- emit a warning,
					 * and perform the call with the value wrapped in a reference. */
					faults.Error(faults.E_WARNING, "Parameter %d to %s%s%s() expected to be a reference, value given", i+1, lang.CondF1(func_.GetScope() != nil, func() []byte { return func_.GetScope().Name() }, ""), lang.Cond(func_.GetScope() != nil, "::", ""), func_.FunctionName())
					must_wrap = 1
					if EG__().HasException() {
						call.NumArgs() = i
						ZendVmStackFreeArgs(call)
						ZendVmStackFreeCallFrame(call)
						if CurrEX() == &dummy_execute_data {
							EG__().SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
						}
						return types.FAILURE
					}
				}
			}
		} else {
			if arg.IsRef() && !func_.IsCallViaTrampoline() {

				/* don't separate references for __call */

				arg = types.Z_REFVAL_P(arg)

				/* don't separate references for __call */

			}
		}
		param = call.Arg(i + 1)
		if must_wrap == 0 {
			types.ZVAL_COPY(param, arg)
		} else {
			// arg.TryAddRefcount()
			param.SetNewRef(arg)
		}
	}
	if func_.GetOpArray().IsClosure() {
		var call_info uint32
		//ZEND_CLOSURE_OBJECT(func_).AddRefcount()
		call_info = ZEND_CALL_CLOSURE
		if func_.IsFakeClosure() {
			call_info |= ZEND_CALL_FAKE_CLOSURE
		}
		ZEND_ADD_CALL_FLAG(call, call_info)
	}
	if func_.GetType() == ZEND_USER_FUNCTION {
		var call_via_handler int = func_.IsCallViaTrampoline()
		var current_opline_before_exception *types.ZendOp = EG__().GetOplineBeforeException()
		ZendInitFuncExecuteData(call, func_.GetOpArray(), fci.GetRetval())
		ZendExecuteEx(call)
		EG__().SetOplineBeforeException(current_opline_before_exception)
		if call_via_handler != 0 {

			/* We must re-initialize function again */

			fciCache.SetFunctionHandler(nil)

			/* We must re-initialize function again */

		}
	} else if func_.GetType() == ZEND_INTERNAL_FUNCTION {
		var call_via_handler int = func_.IsCallViaTrampoline()
		fci.GetRetval().SetNull()
		call.SetPrevExecuteData(CurrEX())
		EG__().SetCurrentExecuteData(call)
		if ZendExecuteInternal == nil {

			/* saves one function call if zend_execute_internal is not used */

			func_.GetInternalFunction().GetHandler()(call, fci.GetRetval())

			/* saves one function call if zend_execute_internal is not used */

		} else {
			ZendExecuteInternal(call, fci.GetRetval())
		}
		EG__().SetCurrentExecuteData(call.GetPrevExecuteData())
		ZendVmStackFreeArgs(call)
		if EG__().HasException() {
			// ZvalPtrDtor(fci.GetRetval())
			fci.GetRetval().SetUndef()
		}
		if call_via_handler != 0 {

			/* We must re-initialize function again */

			fciCache.SetFunctionHandler(nil)

			/* We must re-initialize function again */

		}
	} else {
		fci.GetRetval().SetNull()

		/* Not sure what should be done here if it's a static method */

		if fci.GetObject() != nil {
			call.SetPrevExecuteData(CurrEX())
			EG__().SetCurrentExecuteData(call)
			fci.GetObject().CallMethod(func_.FunctionName(), fci.GetObject(), call, fci.GetRetval())
			EG__().SetCurrentExecuteData(call.GetPrevExecuteData())
		} else {
			faults.ThrowError(nil, "Cannot call overloaded function for non-object")
		}
		ZendVmStackFreeArgs(call)
		if func_.GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY {
			// types.ZendStringReleaseEx(func_.GetFunctionName(), 0)
		}
		Efree(func_)
		if EG__().HasException() {
			// ZvalPtrDtor(fci.GetRetval())
			fci.GetRetval().SetUndef()
		}
	}
	ZendVmStackFreeCallFrame(call)
	if CurrEX() == &dummy_execute_data {
		EG__().SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
	}
	if EG__().HasException() {
		if CurrEX() == nil {
			faults.ThrowExceptionInternal(nil)
		} else if CurrEX().GetFunc() != nil && ZEND_USER_CODE(CurrEX().GetFunc().GetType()) {
			faults.RethrowException(CurrEX())
		}
	}
	return types.SUCCESS
}
func isValidClassName(name string) bool {
	return strkit.IndexAnyExcept(name, validClassNameChars) < 0
}
func trimClassName(name string) string {
	if name != "" && name[0] == '\\' {
		return name[1:]
	}
	return name
}

func ZendLookupClassEx(name string, key string, flags uint32) *types.ClassEntry {
	if name == "" && key == "" {
		return nil
	}

	var lcName = key
	if lcName == "" {
		lcName = ascii.StrToLower(trimClassName(name))
	}

	if ce := EG__().ClassTable().Get(lcName); ce != nil {
		if !ce.IsLinked() {
			if (flags&ZEND_FETCH_CLASS_ALLOW_UNLINKED) != 0 || (flags&ZEND_FETCH_CLASS_ALLOW_NEARLY_LINKED) != 0 && ce.IsNearlyLinked() {
				ce.SetIsHasUnlinkedUses(true)
				return ce
			}
			return nil
		}
		return ce
	}

	/* The compiler is not-reentrant. Make sure we __autoload() only during run-time
	 * (doesn't impact functionality of __autoload()
	 */
	if (flags&ZEND_FETCH_CLASS_NO_AUTOLOAD) != 0 || ZendIsCompiling() {
		return nil
	}
	if EG__().GetAutoloadFunc() == nil {
		var func_ types.IFunction = ZendFetchFunctionStr(types.STR_MAGIC_AUTOLOAD)
		if func_ != nil {
			EG__().SetAutoloadFunc(func_)
		} else {
			return nil
		}
	}

	/* Verify class name before passing it to __autoload() */
	if lcName == "" && isValidClassName(name) {
		return nil
	}
	if EG__().GetInAutoload() == nil {
		EG__().SetInAutoload(types.NewArray())
	}
	if types.ZendHashAddEmptyElement(EG__().GetInAutoload(), lcName) == nil {
		return nil
	}

	// init fci
	var arg0 = trimClassName(name)
	var fci = types.InitFCallInfo(nil, nil, types.NewZvalString(arg0))
	fci.SetFunctionName(EG__().GetAutoloadFunc().FunctionName())

	// init fcc
	var fcc types.ZendFcallInfoCache
	fcc.SetFunctionHandler(EG__().GetAutoloadFunc())
	fcc.SetCalledScope(nil)
	fcc.SetObject(nil)

	origFakeScope := EG__().GetFakeScope()
	EG__().SetFakeScope(nil)
	EG__().ExceptionSave()

	var ce *types.ClassEntry = nil
	if ZendCallFunction(fci, &fcc) == types.SUCCESS && EG__().NoException() {
		ce = EG__().ClassTable().Get(lcName)
	}

	EG__().ExceptionRestore()
	EG__().SetFakeScope(origFakeScope)
	EG__().GetInAutoload().KeyDelete(lcName)
	return ce
}
func ZendLookupClass(name string) *types.ClassEntry {
	return ZendLookupClassEx(name, "", 0)
}
func ZendGetCalledScope(ex *ZendExecuteData) *types.ClassEntry {
	for ex != nil {
		if ce := ex.ThisClass(); ce != nil {
			return ce
		} else if ex.GetFunc() != nil {
			if ex.GetFunc().GetType() != ZEND_INTERNAL_FUNCTION || ex.GetFunc().GetScope() != nil {
				return nil
			}
		}
		ex = ex.GetPrevExecuteData()
	}
	return nil
}
func ZendGetThisObject(ex *ZendExecuteData) *types.Object {
	for ex != nil {
		if obj := ex.ThisObject(); obj != nil {
			return obj
		} else if ex.GetFunc() != nil {
			if !ex.GetFunc().IsInternalFunction() || ex.GetFunc().GetScope() != nil {
				return nil
			}
		}
		ex = ex.GetPrevExecuteData()
	}
	return nil
}

func ZendEvalStringl(str string, retval_ptr *types.Zval, string_name *byte) int {
	var pv types.Zval
	var new_op_array *types.ZendOpArray
	var original_compiler_options uint32
	var retval int
	if retval_ptr != nil {
		pv.SetString("return " + str + ";")
	} else {
		pv.SetString(str)
	}

	original_compiler_options = CG__().GetCompilerOptions()
	CG__().SetCompilerOptions(ZEND_COMPILE_DEFAULT_FOR_EVAL)
	new_op_array = CompileString(&pv, string_name)
	CG__().SetCompilerOptions(original_compiler_options)
	if new_op_array != nil {
		var local_retval types.Zval
		new_op_array.SetScope(ZendGetExecutedScope())

		faults.TryCatch(func() {
			local_retval.SetUndef()
			ZendExecute(new_op_array, &local_retval)
		}, func() {
			faults.Bailout()
		})

		if local_retval.IsNotUndef() {
			if retval_ptr != nil {
				types.ZVAL_COPY_VALUE(retval_ptr, &local_retval)
			}
		} else {
			if retval_ptr != nil {
				retval_ptr.SetNull()
			}
		}
		retval = types.SUCCESS
	} else {
		retval = types.FAILURE
	}

	return retval
}
func ZendEvalStringlEx(str string, retval_ptr *types.Zval, string_name *byte, handle_exceptions int) int {
	var result int
	result = ZendEvalStringl(str, retval_ptr, string_name)
	if handle_exceptions != 0 && EG__().HasException() {
		faults.ExceptionError(EG__().GetException(), faults.E_ERROR)
		result = types.FAILURE
	}
	return result
}
func ZendEvalStringEx(str string, retval_ptr *types.Zval, string_name string, handle_exceptions int) int {
	return ZendEvalStringlEx(str, retval_ptr, string_name, handle_exceptions)
}
func ZendTimeout(dummy int) {
	EG__().SetTimedOut(0)
	ZendSetTimeoutEx(0, 1)
	faults.ErrorNoreturn(faults.E_ERROR, "Maximum execution time of "+ZEND_LONG_FMT+" second%s exceeded", EG__().GetTimeoutSeconds(), lang.Cond(EG__().GetTimeoutSeconds() == 1, "", "s"))
}
func ZendTimeoutHandler(dummy int) {
	if EG__().GetTimedOut() != 0 {

		/* Die on hard timeout */

		var error_filename string = ""
		var error_lineno uint32 = 0
		var log_buffer []byte
		var output_len int = 0
		if ZendIsCompiling() != 0 {
			error_filename = ZendGetCompiledFilename()
			error_lineno = ZendGetCompiledLineno()
		} else if ZendIsExecuting() {
			error_filename = ZendGetExecutedFilename()
			if error_filename[0] == '[' {
				error_filename = nil
				error_lineno = 0
			} else {
				error_lineno = ZendGetExecutedLineno()
			}
		}
		if error_filename == "" {
			error_filename = "Unknown"
		}
		output_len = core.Snprintf(log_buffer, b.SizeOf("log_buffer"), "\nFatal error: Maximum execution time of "+ZEND_LONG_FMT+"+"+ZEND_LONG_FMT+" seconds exceeded (terminated) in %s on line %d\n", EG__().GetTimeoutSeconds(), EG__().GetHardTimeout(), error_filename, error_lineno)
		if output_len > 0 {
			ZendQuietWrite(2, log_buffer, b.Min(output_len, b.SizeOf("log_buffer")))
		}
		_exit(124)
	}
	if ZendOnTimeout != nil {

		/*
		   We got here because we got a timeout signal, so we are in a signal handler
		   at this point. However, we want to be able to timeout any user-supplied
		   shutdown functions, so pretend we are not in a signal handler while we are
		   calling these
		*/

		SIGG(running) = 0
		ZendOnTimeout(EG__().GetTimeoutSeconds())
	}
	EG__().SetTimedOut(1)
	EG__().SetVmInterrupt(1)
	if EG__().GetHardTimeout() > 0 {

		/* Set hard timeout */

		ZendSetTimeoutEx(EG__().GetHardTimeout(), 1)

		/* Set hard timeout */

	}
}
func ZendSetTimeoutEx(seconds ZendLong, reset_signals int) {
	var t_r __struct__itimerval
	var signo int
	if seconds != 0 {
		t_r.it_value.tv_sec = seconds
		t_r.it_interval.tv_usec = 0
		t_r.it_interval.tv_sec = t_r.it_interval.tv_usec
		t_r.it_value.tv_usec = t_r.it_interval.tv_sec
		setitimer(ITIMER_PROF, &t_r, nil)
	}
	signo = SIGPROF
	if reset_signals != 0 {
		ZendSignal(signo, ZendTimeoutHandler)
	}
}
func ZendSetTimeout(seconds ZendLong, reset_signals int) {
	EG__().SetTimeoutSeconds(seconds)
	ZendSetTimeoutEx(seconds, reset_signals)
	EG__().SetTimedOut(0)
}
func ZendUnsetTimeout() {
	if EG__().GetTimeoutSeconds() != 0 {
		var no_timeout __struct__itimerval
		no_timeout.it_interval.tv_usec = 0
		no_timeout.it_interval.tv_sec = no_timeout.it_interval.tv_usec
		no_timeout.it_value.tv_usec = no_timeout.it_interval.tv_sec
		no_timeout.it_value.tv_sec = no_timeout.it_value.tv_usec
		setitimer(ITIMER_PROF, &no_timeout, nil)
	}
	EG__().SetTimedOut(0)
}
func ZendFetchClass(className string, fetchType int) *types.ClassEntry {
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var fetch_sub_type int = fetchType & ZEND_FETCH_CLASS_MASK
check_fetch_type:
	switch fetch_sub_type {
	case ZEND_FETCH_CLASS_SELF:
		scope = ZendGetExecutedScope()
		if scope == nil {
			ZendThrowOrError(fetchType, nil, "Cannot access self:: when no class scope is active")
		}
		return scope
	case ZEND_FETCH_CLASS_PARENT:
		scope = ZendGetExecutedScope()
		if scope == nil {
			ZendThrowOrError(fetchType, nil, "Cannot access parent:: when no class scope is active")
			return nil
		}
		if !(scope.GetParent()) {
			ZendThrowOrError(fetchType, nil, "Cannot access parent:: when current class scope has no parent")
		}
		return scope.GetParent()
	case ZEND_FETCH_CLASS_STATIC:
		ce = ZendGetCalledScope(CurrEX())
		if ce == nil {
			ZendThrowOrError(fetchType, nil, "Cannot access static:: when no class scope is active")
			return nil
		}
		return ce
	case ZEND_FETCH_CLASS_AUTO:
		fetch_sub_type = ZendGetClassFetchType(className)
		if fetch_sub_type != ZEND_FETCH_CLASS_DEFAULT {
			goto check_fetch_type
		}
	}

	ce = ZendLookupClassEx(className, "", fetchType)
	if ce != nil || fetchType&ZEND_FETCH_CLASS_NO_AUTOLOAD != 0 || fetchType&ZEND_FETCH_CLASS_SILENT != 0 {
		return ce
	}
	if EG__().NoException() {
		if fetch_sub_type == ZEND_FETCH_CLASS_INTERFACE {
			ZendThrowOrError(fetchType, nil, fmt.Sprintf("Interface '%s' not found", className))
		} else if fetch_sub_type == ZEND_FETCH_CLASS_TRAIT {
			ZendThrowOrError(fetchType, nil, fmt.Sprintf("Trait '%s' not found", className))
		} else {
			ZendThrowOrError(fetchType, nil, fmt.Sprintf("Class '%s' not found", className))
		}
	}
	return nil
}
func ZendFetchClassByNameEx(className types.ClassName, fetchType int) *types.ClassEntry {
	return ZendFetchClassByName(className.GetName(), className.GetLcName(), fetchType)
}
func ZendFetchClassByName(className string, key string, fetchType int) *types.ClassEntry {
	ce := ZendLookupClassEx(className, key, fetchType)
	if ce != nil || fetchType&ZEND_FETCH_CLASS_NO_AUTOLOAD != 0 || fetchType&ZEND_FETCH_CLASS_SILENT != 0 {
		return ce
	}

	if EG__().HasException() {
		if (fetchType & ZEND_FETCH_CLASS_EXCEPTION) == 0 {
			exceptionZv := types.NewZvalObject(EG__().GetException())
			exceptionStr := operators.ZvalGetStrVal(exceptionZv)
			EG__().ClearException()
			faults.ErrorNoreturn(faults.E_ERROR, "During class fetch: Uncaught %s", exceptionStr)
		}
		return nil
	}
	if (fetchType & ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_INTERFACE {
		ZendThrowOrError(fetchType, nil, fmt.Sprintf("Interface '%s' not found", className))
	} else if (fetchType & ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_TRAIT {
		ZendThrowOrError(fetchType, nil, fmt.Sprintf("Trait '%s' not found", className))
	} else {
		ZendThrowOrError(fetchType, nil, fmt.Sprintf("Class '%s' not found", className))
	}
	return nil
}
func ZendDeleteGlobalVariable(name string) bool {
	return EG__().GetSymbolTable().KeyDeleteIndirect(name)
}
func ZendRebuildSymbolTable() *types.Array {
	var ex *ZendExecuteData
	var symbol_table *types.Array

	/* Search for last called user function */

	ex = CurrEX()
	for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
		ex = ex.GetPrevExecuteData()
	}
	if ex == nil {
		return nil
	}
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
		return ex.GetSymbolTable()
	}
	ZEND_ADD_CALL_FLAG(ex, ZEND_CALL_HAS_SYMBOL_TABLE)

	symbol_table = EG__().PopSymbolTable(ex.GetFunc().GetOpArray().GetLastVar())
	ex.SetSymbolTable(symbol_table)

	if ex.GetFunc().GetOpArray().GetLastVar() != 0 {
		vars := ex.GetFunc().GetOpArray().VarNames()
		for i, varName := range vars {
			var_ := ex.VarNum(i)
			types.ZendHashAppendInd(symbol_table, varName, var_)
		}

	}
	return symbol_table
}
func ZendAttachSymbolTable(executeData *ZendExecuteData) {
	var op_array *types.ZendOpArray = executeData.GetFunc().GetOpArray()
	var ht *types.Array = executeData.GetSymbolTable()

	/* copy real values from symbol table into CV slots and create
	   INDIRECT references to CV in symbol table  */

	if op_array.GetLastVar() != 0 {
		vars := op_array.VarNames()
		for i, varname := range vars {
			var_ := executeData.VarNum(i)
			var zv *types.Zval = ht.KeyFind(varname)
			if zv != nil {
				if zv.IsIndirect() {
					var val *types.Zval = zv.Indirect()
					var_.CopyValueFrom(val)
				} else {
					var_.CopyValueFrom(zv)
				}
			} else {
				var_.SetUndef()
				zv = ht.KeyAddNew(varname, var_)
			}
			zv.SetIndirect(var_)
		}
	}

	/* copy real values from symbol table into CV slots and create
	   INDIRECT references to CV in symbol table  */
}
func ZendDetachSymbolTable(executeData *ZendExecuteData) {
	var op_array *types.ZendOpArray = executeData.GetFunc().GetOpArray()
	var ht *types.Array = executeData.GetSymbolTable()

	/* copy real values from CV slots into symbol table */

	vars := op_array.VarNames()
	if len(vars) != 0 {
		for i, varname := range vars {
			var_ := executeData.VarNum(i)
			if var_.IsUndef() {
				ht.KeyDelete(varname)
			} else {
				ht.KeyUpdate(varname, var_)
				var_.SetUndef()
			}
		}
	}

	/* copy real values from CV slots into symbol table */
}
func ZendSetLocalVarStr(name string, value *types.Zval, force int) int {
	var executeData *ZendExecuteData = CurrEX()
	for executeData != nil && (executeData.GetFunc() == nil || !(ZEND_USER_CODE(executeData.GetFunc().GetType()))) {
		executeData = executeData.GetPrevExecuteData()
	}
	if executeData != nil {
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) == 0 {
			var op_array *types.ZendOpArray = executeData.GetFunc().GetOpArray()

			if i := op_array.FindVarName(name); i >= 0 {
				var var_ = executeData.VarNum(i)
				var_.CopyValueFrom(value)
				return types.SUCCESS
			}

			if force != 0 {
				var symbol_table *types.Array = ZendRebuildSymbolTable()
				if symbol_table != nil {
					symbol_table.KeyUpdate(name, value)
					return types.SUCCESS
				}
			}
		} else {
			executeData.GetSymbolTable().KeyUpdateIndirect(name, value)
			return types.SUCCESS
		}
	}
	return types.FAILURE
}
