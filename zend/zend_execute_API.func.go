package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZendExtensionActivator(extension *ZendExtension) {
	if extension.GetActivate() != nil {
		extension.GetActivate()()
	}
}
func ZendExtensionDeactivator(extension *ZendExtension) {
	if extension.GetDeactivate() != nil {
		extension.GetDeactivate()()
	}
}
func CleanNonPersistentConstantFull(zv *types.Zval) int {
	var c *ZendConstant = zv.GetPtr()
	if (ZEND_CONSTANT_FLAGS(c) & CONST_PERSISTENT) != 0 {
		return types.ArrayApplyKeep
	} else {
		return types.ArrayApplyRemove
	}
}
func InitExecutor() {
	ZendInitFpu()
	EG__().GetUninitializedZval().SetNull()
	EG__().GetErrorZval().IsError()

	/* destroys stack frame, therefore makes core dumps worthless */

	EG__().SetSymtableCachePtr(EG__().GetSymtableCache())
	EG__().SetSymtableCacheLimit(EG__().GetSymtableCache() + SYMTABLE_CACHE_SIZE)
	EG__().SetNoExtensions(0)
	EG__().SetFunctionTable(CG__().FunctionTable())
	EG__().SetClassTable(CG__().ClassTable())
	EG__().SetInAutoload(nil)
	EG__().SetAutoloadFunc(nil)
	EG__().SetErrorHandling(EH_NORMAL)
	EG__().SetFlags(EG_FLAGS_INITIAL)
	ZendVmStackInit()
	EG__().GetSymbolTable() = types.MakeArrayEx(64, ZVAL_PTR_DTOR, 0)
	ZendExtensions.Apply(LlistApplyFuncT(ZendExtensionActivator))
	EG__().GetIncludedFiles() = types.MakeArrayEx(8, nil, 0)
	EG__().SetTicksCount(0)
	EG__().GetUserErrorHandler().SetUndef()
	EG__().GetUserExceptionHandler().SetUndef()
	EG__().SetCurrentExecuteData(nil)
	EG__().GetUserErrorHandlersErrorReporting().Init()
	EG__().GetUserErrorHandlers().Init()
	EG__().GetUserExceptionHandlers().Init()
	ZendObjectsStoreInit(EG__().GetObjectsStore(), 1024)
	EG__().SetVmInterrupt(0)
	EG__().SetTimedOut(0)
	EG__().SetException(nil)
	EG__().SetPrevException(nil)
	EG__().SetFakeScope(nil)
	EG__().GetTrampoline().SetFunctionName(nil)
	EG__().SetHtIteratorsCount(b.SizeOf("EG ( ht_iterators_slots )") / b.SizeOf("HashTableIterator"))
	EG__().SetHtIteratorsUsed(0)
	EG__().SetHtIterators(EG__().GetHtIteratorsSlots())
	memset(EG__().GetHtIterators(), 0, b.SizeOf("EG ( ht_iterators_slots )"))
	EG__().SetEachDeprecationThrown(0)
	EG__().SetPersistentConstantsCount(EG__().GetZendConstants().GetNNumUsed())
	EG__().SetPersistentFunctionsCount(uint32(EG__().FunctionTable().Len()))
	EG__().SetPersistentClassesCount(uint32(EG__().ClassTable().Len()))
	ZendWeakrefsInit()
	EG__().SetActive(1)
}
func ZvalCallDestructor(zv *types.Zval) int {
	if zv.IsIndirect() {
		zv = zv.GetZv()
	}
	if zv.IsObject() && zv.GetRefcount() == 1 {
		return types.ArrayApplyRemove
	} else {
		return types.ArrayApplyKeep
	}
}
func ZendUncleanZvalPtrDtor(zv *types.Zval) {
	if zv.IsIndirect() {
		zv = zv.GetZv()
	}
	IZvalPtrDtor(zv)
}
func ZendThrowOrError(fetch_type int, exception_ce *types.ClassEntry, format string, args ...any) {
	message := ZendSprintf(format, args)
	if (fetch_type & ZEND_FETCH_CLASS_EXCEPTION) != 0 {
		faults.ThrowError(exception_ce, "%s", message)
	} else {
		faults.Error(faults.E_ERROR, "%s", message)
	}
}
func ShutdownDestructors() {
	if CG__().GetUncleanShutdown() != 0 {
		EG__().GetSymbolTable().SetPDestructor(ZendUncleanZvalPtrDtor)
	}
	faults.TryCatch(func() {
		var symbols uint32
		for {
			symbols = EG__().GetSymbolTable().Len()
			types.ZendHashReverseApply(EG__().GetSymbolTable(), types.ApplyFuncT(ZvalCallDestructor))
			if symbols == EG__().GetSymbolTable().Len() {
				break
			}
		}
		ZendObjectsStoreCallDestructors(EG__().GetObjectsStore())
	}, func() {
		/* if we couldn't destruct cleanly, mark all objects as destructed anyway */
		ZendObjectsStoreMarkDestructed(EG__().GetObjectsStore())
	})
}
func ShutdownExecutor() {
	var key *types.String
	var zv *types.Zval
	var fast_shutdown types.ZendBool = IsZendMm() != 0 && EG__().GetFullTablesCleanup() == 0

	faults.Try(func() {
		CG__().GetOpenFiles().Destroy()
	})

	EG__().SetIsInResourceShutdown(true)

	faults.Try(func() {
		ZendCloseRsrcList(EG__().GetRegularList())
	})

	/* No PHP callback functions should be called after this point. */

	EG__().SetActive(0)
	if fast_shutdown == 0 {
		EG__().GetSymbolTable().GracefulReverseDestroy()

		/* Release static properties and static variables prior to the final GC run,
		 * as they may hold GC roots. */
		EG__().FunctionTable().ForeachReserve(func(_ string, f types.IFunction) {
			if f.GetType() == ZEND_INTERNAL_FUNCTION {
				return
			}

			opArray := f.GetOpArray()
			if opArray.GetStaticVariables() != nil {
				var ht *types.Array = ZEND_MAP_PTR_GET(opArray.static_variables_ptr)
				if ht != nil {
					if (ht.GetGcFlags()&types.IS_ARRAY_IMMUTABLE) == 0 && ht.DelRefcount() == 0 {
						ht.DestroyEx()
					}
					ZEND_MAP_PTR_SET(opArray.static_variables_ptr, nil)
				}
			}
		})

		EG__().ClassTable().ForeachReserve(func(_ string, ce *types.ClassEntry) {
			if ce.GetDefaultStaticMembersCount() != 0 {
				ZendCleanupInternalClassData(ce)
			}
			if ce.IsHasStaticInMethods() {
				var op_array *types.ZendOpArray
				var __ht *types.Array = ce.GetFunctionTable()
				for _, _p := range __ht.ForeachData() {
					var _z *types.Zval = _p.GetVal()

					op_array = _z.GetPtr()
					if op_array.GetType() == ZEND_USER_FUNCTION {
						if op_array.GetStaticVariables() != nil {
							var ht *types.Array = ZEND_MAP_PTR_GET(op_array.static_variables_ptr)
							if ht != nil {
								if (ht.GetGcFlags()&types.IS_ARRAY_IMMUTABLE) == 0 && ht.DelRefcount() == 0 {
									ht.DestroyEx()
								}
								ZEND_MAP_PTR_SET(op_array.static_variables_ptr, nil)
							}
						}
					}
				}
			}
		})

		/* Also release error and exception handlers, which may hold objects. */

		if EG__().GetUserErrorHandler().IsNotUndef() {
			ZvalPtrDtor(EG__().GetUserErrorHandler())
			EG__().GetUserErrorHandler().SetUndef()
		}
		if EG__().GetUserExceptionHandler().IsNotUndef() {
			ZvalPtrDtor(EG__().GetUserExceptionHandler())
			EG__().GetUserExceptionHandler().SetUndef()
		}
		ZendStackClean(EG__().GetUserErrorHandlersErrorReporting(), nil, 1)
		ZendStackClean(EG__().GetUserErrorHandlers(), (func(any))(ZVAL_PTR_DTOR), 1)
		ZendStackClean(EG__().GetUserExceptionHandlers(), (func(any))(ZVAL_PTR_DTOR), 1)
	}
	ZendObjectsStoreFreeObjectStorage(EG__().GetObjectsStore(), fast_shutdown)
	ZendWeakrefsShutdown()

	faults.Try(func() {
		ZendExtensions.Apply(LlistApplyFuncT(ZendExtensionDeactivator))
	})

	if fast_shutdown != 0 {

		/* Fast Request Shutdown
		 * =====================
		 * Zend Memory Manager frees memory by its own. We don't have to free
		 * each allocated block separately.
		 */

		EG__().GetZendConstants().Discard(EG__().GetPersistentConstantsCount())
		EG__().FunctionTable().FilterReserve(func(_ string, f types.IFunction) bool {
			return f.GetType() == ZEND_INTERNAL_FUNCTION
		})
		EG__().ClassTable().FilterReserve(func(_ string, ce *types.ClassEntry) bool {
			return ce.GetType() == ZEND_INTERNAL_CLASS
		})
		ZendCleanupInternalClasses()
	} else {
		ZendVmStackDestroy()
		if EG__().GetFullTablesCleanup() != 0 {
			types.ZendHashReverseApply(EG__().GetZendConstants(), CleanNonPersistentConstantFull)
			EG__().FunctionTable().FilterReserve(func(_ string, f types.IFunction) bool {
				return f.GetType() == ZEND_INTERNAL_FUNCTION
			})
			EG__().ClassTable().FilterReserve(func(_ string, ce *types.ClassEntry) bool {
				return ce.GetType() == ZEND_INTERNAL_CLASS
			})

		} else {
			var __ht *types.Array = EG__().GetZendConstants()
			for _, _p := range __ht.ForeachDataReserve() {
				var _z types.Zval = _p.GetVal()

				key = _p.GetKey()
				zv = _z
				var c *ZendConstant = zv.GetPtr()
				if _idx == EG__().GetPersistentConstantsCount() {
					break
				}
				ZvalPtrDtorNogc(c.Value())
				if c.GetName() != nil {
					// types.ZendStringReleaseEx(c.GetName(), 0)
				}
				Efree(c)
				// types.ZendStringReleaseEx(key, 0)
				__ht.Len()--
				var j uint32 = types.HT_IDX_TO_HASH(_idx - 1)
				var nIndex uint32 = _p.GetH() | __ht.GetNTableMask()
				var i uint32 = types.HT_HASH(__ht, nIndex)
				if j != i {
					var prev *types.Bucket = __ht.Bucket(i)
					for prev.GetVal().GetNext() != j {
						i = prev.GetVal().GetNext()
						prev = __ht.Bucket(i)
					}
					prev.GetVal().GetNext() = _p.GetVal().GetNext()
				} else {
					types.HT_HASH(__ht, nIndex) = _p.GetVal().GetNext()
				}
			}
			__ht.SetNNumUsed(_idx)

			EG__().FunctionTable().FilterReserve(func(key string, f types.IFunction) bool {
				if f.GetType() == ZEND_INTERNAL_FUNCTION {
					return true
				}

				DestroyOpArray(f.GetOpArray())
				return false
			})

			EG__().ClassTable().FilterReserve(func(_ string, ce *types.ClassEntry) bool {
				if ce.GetType() == ZEND_INTERNAL_CLASS {
					return true
				}

				DestroyZendClass(zv)
				return false
			})
		}
		for EG__().GetSymtableCachePtr() > EG__().GetSymtableCache() {
			EG__().GetSymtableCachePtr()--
			(*EG__)().symtable_cache_ptr.Destroy()
			FREE_HASHTABLE((*EG__)().symtable_cache_ptr)
		}
		EG__().GetIncludedFiles().Destroy()
		EG__().GetUserErrorHandlersErrorReporting().Destroy()
		EG__().GetUserErrorHandlers().Destroy()
		EG__().GetUserExceptionHandlers().Destroy()
		ZendObjectsStoreDestroy(EG__().GetObjectsStore())
		if EG__().GetInAutoload() != nil {
			EG__().GetInAutoload().Destroy()
			FREE_HASHTABLE(EG__().GetInAutoload())
		}
		if EG__().GetHtIterators() != EG__().GetHtIteratorsSlots() {
			Efree(EG__().GetHtIterators())
		}
	}
	EG__().SetHtIteratorsUsed(0)
	ZendShutdownFpu()
}

func GetActiveCalleeName() string   { return CurrEX().CalleeName() }
func GetActiveFunctionName() string { return CurrEX().FunctionName() }
func ZendGetExecutedFilename() string {
	var ex *ZendExecuteData = CurrEX()
	for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
		ex = ex.GetPrevExecuteData()
	}
	if ex != nil {
		return ex.GetFunc().GetOpArray().GetFilename().GetStr()
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
		return ex.GetFunc().GetOpArray().GetFilename()
	} else {
		return nil
	}
}
func ZendGetExecutedLineno() uint32 {
	var ex *ZendExecuteData = CurrEX()
	for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
		ex = ex.GetPrevExecuteData()
	}
	if ex != nil {
		if EG__().GetException() != nil && ex.GetOpline().GetOpcode() == ZEND_HANDLE_EXCEPTION && ex.GetOpline().GetLineno() == 0 && EG__().GetOplineBeforeException() != nil {
			return EG__().GetOplineBeforeException().GetLineno()
		}
		return ex.GetOpline().GetLineno()
	} else {
		return 0
	}
}
func ZendGetExecutedScope() *types.ClassEntry {
	var ex *ZendExecuteData = CurrEX()
	for true {
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
func ZendUseUndefinedConstant(name *types.String, attr ZendAstAttr, result *types.Zval) int {
	var colon *byte
	if EG__().GetException() != nil {
		return types.FAILURE
	} else if b.Assign(&colon, (*byte)(ZendMemrchr(name.GetVal(), ':', name.GetLen()))) {
		faults.ThrowError(nil, "Undefined class constant '%s'", name.GetVal())
		return types.FAILURE
	} else if (attr & IS_CONSTANT_UNQUALIFIED) == 0 {
		faults.ThrowError(nil, "Undefined constant '%s'", name.GetVal())
		return types.FAILURE
	} else {
		var actual *byte = name.GetVal()
		var actual_len int = name.GetLen()
		var slash *byte = (*byte)(ZendMemrchr(actual, '\\', actual_len))
		if slash != nil {
			actual = slash + 1
			actual_len -= actual - name.GetVal()
		}
		faults.Error(faults.E_WARNING, "Use of undefined constant %s - assumed '%s' (this will throw an Error in a future version of PHP)", actual, actual)
		if EG__().GetException() != nil {
			return types.FAILURE
		} else {
			var result_str *types.String = types.NewString(b.CastStr(actual, actual_len))
			ZvalPtrDtorNogc(result)
			result.SetString(result_str)
		}
	}
	return types.SUCCESS
}
func ZvalUpdateConstantEx(p *types.Zval, scope *types.ClassEntry) int {
	if p.IsConstant() {
		var ast *ZendAst = types.Z_ASTVAL_P(p)
		if ast.GetKind() == ZEND_AST_CONSTANT {
			var name *types.String = ZendAstGetConstantName(ast)
			var zv *types.Zval = ZendGetConstantEx(name, scope, ast.GetAttr())
			if zv == nil {
				return ZendUseUndefinedConstant(name, ast.GetAttr(), p)
			}
			ZvalPtrDtorNogc(p)
			types.ZVAL_COPY_OR_DUP(p, zv)
		} else {
			var tmp types.Zval
			if ZendAstEvaluate(&tmp, ast, scope) != types.SUCCESS {
				return types.FAILURE
			}
			ZvalPtrDtorNogc(p)
			types.ZVAL_COPY_VALUE(p, &tmp)
		}
	}
	return types.SUCCESS
}
func ZendCallFunction(fci *types.ZendFcallInfo, fci_cache *types.ZendFcallInfoCache) int {
	var i uint32
	var call *ZendExecuteData
	var dummy_execute_data ZendExecuteData
	var fci_cache_local types.ZendFcallInfoCache
	var func_ types.IFunction
	var call_info uint32
	var object_or_called_scope any
	fci.GetRetval().SetUndef()
	if EG__().GetActive() == 0 {
		return types.FAILURE
	}
	if EG__().GetException() != nil {
		return types.FAILURE
	}
	b.Assert(fci.GetSize() == b.SizeOf("zend_fcall_info"))

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

		dummy_execute_data = (*EG__)().current_execute_data
		dummy_execute_data.SetPrevExecuteData(CurrEX())
		dummy_execute_data.SetCall(nil)
		dummy_execute_data.SetOpline(nil)
		dummy_execute_data.SetFunc(nil)
		EG__().SetCurrentExecuteData(&dummy_execute_data)
	}
	if fci_cache == nil || fci_cache.GetFunctionHandler() == nil {
		var error *byte = nil
		if fci_cache == nil {
			fci_cache = &fci_cache_local
		}
		if ZendIsCallableEx(fci.GetFunctionName(), fci.GetObject(), IS_CALLABLE_CHECK_SILENT, nil, fci_cache, &error) == 0 {
			if error != nil {
				var callable_name *types.String = ZendGetCallableNameEx(fci.GetFunctionName(), fci.GetObject())
				faults.Error(faults.E_WARNING, "Invalid callback %s, %s", callable_name.GetVal(), error)
				Efree(error)
				// types.ZendStringReleaseEx(callable_name, 0)
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
			if EG__().GetException() != nil {
				if CurrEX() == &dummy_execute_data {
					EG__().SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
				}
				return types.FAILURE
			}
		}
	}
	func_ = fci_cache.GetFunctionHandler()
	if func_.IsStatic() || fci_cache.GetObject() == nil {
		fci.SetObject(nil)
		object_or_called_scope = fci_cache.GetCalledScope()
		call_info = ZEND_CALL_TOP_FUNCTION | ZEND_CALL_DYNAMIC
	} else {
		fci.SetObject(fci_cache.GetObject())
		object_or_called_scope = fci.GetObject()
		call_info = ZEND_CALL_TOP_FUNCTION | ZEND_CALL_DYNAMIC | ZEND_CALL_HAS_THIS
	}
	call = ZendVmStackPushCallFrame(call_info, func_, fci.GetParamCount(), object_or_called_scope)
	if func_.IsDeprecated() {
		faults.Error(faults.E_DEPRECATED, "Function %s%s%s() is deprecated", b.CondF1(func_.GetScope() != nil, func() []byte { return func_.GetScope().GetName().GetVal() }, ""), b.Cond(func_.GetScope() != nil, "::", ""), func_.GetFunctionName().GetVal())
		if EG__().GetException() != nil {
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
		var must_wrap types.ZendBool = 0
		if ARG_SHOULD_BE_SENT_BY_REF(func_, i+1) != 0 {
			if !(arg.IsReference()) {
				if fci.GetNoSeparation() == 0 {

					/* Separation is enabled -- create a ref */

					arg.SetNewRef(arg)

					/* Separation is enabled -- create a ref */

				} else if ARG_MAY_BE_SENT_BY_REF(func_, i+1) == 0 {

					/* By-value send is not allowed -- emit a warning,
					 * and perform the call with the value wrapped in a reference. */

					faults.Error(faults.E_WARNING, "Parameter %d to %s%s%s() expected to be a reference, value given", i+1, b.CondF1(func_.GetScope() != nil, func() []byte { return func_.GetScope().GetName().GetVal() }, ""), b.Cond(func_.GetScope() != nil, "::", ""), func_.GetFunctionName().GetVal())
					must_wrap = 1
					if EG__().GetException() != nil {
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
			if arg.IsReference() && !func_.IsCallViaTrampoline() {

				/* don't separate references for __call */

				arg = types.Z_REFVAL_P(arg)

				/* don't separate references for __call */

			}
		}
		param = call.Arg(i + 1)
		if must_wrap == 0 {
			types.ZVAL_COPY(param, arg)
		} else {
			arg.TryAddRefcount()
			param.SetNewRef(arg)
		}
	}
	if func_.GetOpArray().IsClosure() {
		var call_info uint32
		ZEND_CLOSURE_OBJECT(func_).AddRefcount()
		call_info = ZEND_CALL_CLOSURE
		if func_.IsFakeClosure() {
			call_info |= ZEND_CALL_FAKE_CLOSURE
		}
		ZEND_ADD_CALL_FLAG(call, call_info)
	}
	if func_.GetType() == ZEND_USER_FUNCTION {
		var call_via_handler int = func_.IsCallViaTrampoline()
		var current_opline_before_exception *ZendOp = EG__().GetOplineBeforeException()
		ZendInitFuncExecuteData(call, func_.GetOpArray(), fci.GetRetval())
		ZendExecuteEx(call)
		EG__().SetOplineBeforeException(current_opline_before_exception)
		if call_via_handler != 0 {

			/* We must re-initialize function again */

			fci_cache.SetFunctionHandler(nil)

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
		if EG__().GetException() != nil {
			ZvalPtrDtor(fci.GetRetval())
			fci.GetRetval().SetUndef()
		}
		if call_via_handler != 0 {

			/* We must re-initialize function again */

			fci_cache.SetFunctionHandler(nil)

			/* We must re-initialize function again */

		}
	} else {
		fci.GetRetval().SetNull()

		/* Not sure what should be done here if it's a static method */

		if fci.GetObject() != nil {
			call.SetPrevExecuteData(CurrEX())
			EG__().SetCurrentExecuteData(call)
			fci.GetObject().GetHandlers().GetCallMethod()(func_.GetFunctionName(), fci.GetObject(), call, fci.GetRetval())
			EG__().SetCurrentExecuteData(call.GetPrevExecuteData())
		} else {
			faults.ThrowError(nil, "Cannot call overloaded function for non-object")
		}
		ZendVmStackFreeArgs(call)
		if func_.GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY {
			// types.ZendStringReleaseEx(func_.GetFunctionName(), 0)
		}
		Efree(func_)
		if EG__().GetException() != nil {
			ZvalPtrDtor(fci.GetRetval())
			fci.GetRetval().SetUndef()
		}
	}
	ZendVmStackFreeCallFrame(call)
	if CurrEX() == &dummy_execute_data {
		EG__().SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
	}
	if EG__().GetException() != nil {
		if CurrEX() == nil {
			faults.ThrowExceptionInternal(nil)
		} else if CurrEX().GetFunc() != nil && ZEND_USER_CODE(CurrEX().GetFunc().GetType()) {
			faults.RethrowException(CurrEX())
		}
	}
	return types.SUCCESS
}
func ZendLookupClassEx(name *types.String, key *types.String, flags uint32) *types.ClassEntry {
	var args []types.Zval
	var local_retval types.Zval
	var lc_name *types.String
	var fcall_info types.ZendFcallInfo
	var fcall_cache types.ZendFcallInfoCache
	var orig_fake_scope *types.ClassEntry
	if key != nil {
		lc_name = key
	} else {
		if name == nil || name.GetLen() == 0 {
			return nil
		}
		if name.GetVal()[0] == '\\' {
			lc_name = types.ZendStringAlloc(name.GetLen()-1, 0)
			ZendStrTolowerCopy(lc_name.GetVal(), name.GetVal()+1, name.GetLen()-1)
		} else {
			lc_name = ZendStringTolower(name)
		}
	}

	if ce := EG__().ClassTable().Get(lc_name.GetStr()); ce != nil {
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

	if (flags&ZEND_FETCH_CLASS_NO_AUTOLOAD) != 0 || ZendIsCompiling() != 0 {
		if key == nil {
			// types.ZendStringReleaseEx(lc_name, 0)
		}
		return nil
	}
	if EG__().GetAutoloadFunc() == nil {
		var func_ types.IFunction = ZendFetchFunction(types.ZSTR_MAGIC_AUTOLOAD)
		if func_ != nil {
			EG__().SetAutoloadFunc(func_)
		} else {
			if key == nil {
				// types.ZendStringReleaseEx(lc_name, 0)
			}
			return nil
		}
	}

	/* Verify class name before passing it to __autoload() */

	if key == nil && strspn(name.GetVal(), "0123456789_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ200201202203204205206207210211212213214215216217220221222223224225226227230231232233234235236237240241242243244245246247250251252253254255256257260261262263264265266267270271272273274275276277300301302303304305306307310311312313314315316317320321322323324325326327330331332333334335336337340341342343344345346347350351352353354355356357360361362363364365366367370371372373374375376377\\") != name.GetLen() {
		// types.ZendStringReleaseEx(lc_name, 0)
		return nil
	}
	if EG__().GetInAutoload() == nil {
		ALLOC_HASHTABLE(EG__().GetInAutoload())
		EG__().GetInAutoload() = types.MakeArrayEx(8, nil, 0)
	}
	if types.ZendHashAddEmptyElement(EG__().GetInAutoload(), lc_name.GetStr()) == nil {
		if key == nil {
			// types.ZendStringReleaseEx(lc_name, 0)
		}
		return nil
	}
	local_retval.SetUndef()
	if name.GetVal()[0] == '\\' {
		args[0].SetStringVal(b.CastStr(name.GetVal()+1, name.GetLen()-1))
	} else {
		args[0].SetStringCopy(name)
	}
	fcall_info.SetSize(b.SizeOf("fcall_info"))
	fcall_info.GetFunctionName().SetStringCopy(EG__().GetAutoloadFunc().GetFunctionName())
	fcall_info.SetRetval(&local_retval)
	fcall_info.SetParamCount(1)
	fcall_info.SetParams(args)
	fcall_info.SetObject(nil)
	fcall_info.SetNoSeparation(1)
	fcall_cache.SetFunctionHandler(EG__().GetAutoloadFunc())
	fcall_cache.SetCalledScope(nil)
	fcall_cache.SetObject(nil)
	orig_fake_scope = EG__().GetFakeScope()
	EG__().SetFakeScope(nil)
	faults.ExceptionSave()

	var ce *types.ClassEntry = nil
	if ZendCallFunction(&fcall_info, &fcall_cache) == types.SUCCESS && EG__().GetException() == nil {
		ce = EG__().ClassTable().Get(lc_name.GetStr())
	}

	faults.ExceptionRestore()
	EG__().SetFakeScope(orig_fake_scope)
	ZvalPtrDtor(&args[0])
	types.ZendHashDel(EG__().GetInAutoload(), lc_name.GetStr())
	ZvalPtrDtor(&local_retval)
	return ce
}
func ZendLookupClass(name *types.String) *types.ClassEntry {
	return ZendLookupClassEx(name, nil, 0)
}
func ZendGetCalledScope(ex *ZendExecuteData) *types.ClassEntry {
	for ex != nil {
		if ex.GetThis().IsObject() {
			return types.Z_OBJCE(ex.GetThis())
		} else if ex.GetThis().GetCe() != nil {
			return ex.GetThis().GetCe()
		} else if ex.GetFunc() != nil {
			if ex.GetFunc().GetType() != ZEND_INTERNAL_FUNCTION || ex.GetFunc().GetScope() != nil {
				return nil
			}
		}
		ex = ex.GetPrevExecuteData()
	}
	return nil
}
func ZendGetThisObject(ex *ZendExecuteData) *types.ZendObject {
	for ex != nil {
		if ex.GetThis().IsObject() {
			return ex.GetThis().GetObj()
		} else if ex.GetFunc() != nil {
			if ex.GetFunc().GetType() != ZEND_INTERNAL_FUNCTION || ex.GetFunc().GetScope() != nil {
				return nil
			}
		}
		ex = ex.GetPrevExecuteData()
	}
	return nil
}
func ZendEvalStringl(str *byte, str_len int, retval_ptr *types.Zval, string_name *byte) int {
	var pv types.Zval
	var new_op_array *types.ZendOpArray
	var original_compiler_options uint32
	var retval int
	if retval_ptr != nil {
		pv.SetString(types.ZendStringAlloc(str_len+b.SizeOf("\"return ;\"")-1, 0))
		memcpy(pv.GetStr().GetVal(), "return ", b.SizeOf("\"return \"")-1)
		memcpy(pv.GetStr().GetVal()+b.SizeOf("\"return \"")-1, str, str_len)
		pv.GetStr().GetVal()[pv.GetStr().GetLen()-1] = ';'
		pv.GetStr().GetVal()[pv.GetStr().GetLen()] = '0'
	} else {
		/*printf("Evaluating '%s'\n", pv.value.str.val);*/
		pv.SetStringVal(b.CastStr(str, str_len))
	}

	original_compiler_options = CG__().GetCompilerOptions()
	CG__().SetCompilerOptions(ZEND_COMPILE_DEFAULT_FOR_EVAL)
	new_op_array = ZendCompileString(&pv, string_name)
	CG__().SetCompilerOptions(original_compiler_options)
	if new_op_array != nil {
		var local_retval types.Zval
		EG__().SetNoExtensions(1)
		new_op_array.SetScope(ZendGetExecutedScope())

		faults.TryCatch(func() {
			local_retval.SetUndef()
			ZendExecute(new_op_array, &local_retval)
		}, func() {
			DestroyOpArray(new_op_array)
			EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
			faults.Bailout()
		})

		if local_retval.IsNotUndef() {
			if retval_ptr != nil {
				types.ZVAL_COPY_VALUE(retval_ptr, &local_retval)
			} else {
				ZvalPtrDtor(&local_retval)
			}
		} else {
			if retval_ptr != nil {
				retval_ptr.SetNull()
			}
		}
		EG__().SetNoExtensions(0)
		DestroyOpArray(new_op_array)
		EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
		retval = types.SUCCESS
	} else {
		retval = types.FAILURE
	}

	return retval
}
func ZendEvalString(str *byte, retval_ptr *types.Zval, string_name *byte) int {
	return ZendEvalStringl(str, strlen(str), retval_ptr, string_name)
}
func ZendEvalStringlEx(str *byte, str_len int, retval_ptr *types.Zval, string_name *byte, handle_exceptions int) int {
	var result int
	result = ZendEvalStringl(str, str_len, retval_ptr, string_name)
	if handle_exceptions != 0 && EG__().GetException() != nil {
		faults.ExceptionError(EG__().GetException(), faults.E_ERROR)
		result = types.FAILURE
	}
	return result
}
func ZendEvalStringEx(str *byte, retval_ptr *types.Zval, string_name string, handle_exceptions int) int {
	return ZendEvalStringlEx(str, strlen(str), retval_ptr, string_name, handle_exceptions)
}
func ZendTimeout(dummy int) {
	EG__().SetTimedOut(0)
	ZendSetTimeoutEx(0, 1)
	faults.ErrorNoreturn(faults.E_ERROR, "Maximum execution time of "+ZEND_LONG_FMT+" second%s exceeded", EG__().GetTimeoutSeconds(), b.Cond(EG__().GetTimeoutSeconds() == 1, "", "s"))
}
func ZendTimeoutHandler(dummy int) {
	if EG__().GetTimedOut() != 0 {

		/* Die on hard timeout */

		var error_filename *byte = nil
		var error_lineno uint32 = 0
		var log_buffer []byte
		var output_len int = 0
		if ZendIsCompiling() != 0 {
			error_filename = ZendGetCompiledFilename().GetVal()
			error_lineno = ZendGetCompiledLineno()
		} else if ZendIsExecuting() != 0 {
			error_filename = ZendGetExecutedFilename()
			if error_filename[0] == '[' {
				error_filename = nil
				error_lineno = 0
			} else {
				error_lineno = ZendGetExecutedLineno()
			}
		}
		if error_filename == nil {
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
func ZendFetchClass(class_name *types.String, fetch_type int) *types.ClassEntry {
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var fetch_sub_type int = fetch_type & ZEND_FETCH_CLASS_MASK
check_fetch_type:
	switch fetch_sub_type {
	case ZEND_FETCH_CLASS_SELF:
		scope = ZendGetExecutedScope()
		if scope == nil {
			ZendThrowOrError(fetch_type, nil, "Cannot access self:: when no class scope is active")
		}
		return scope
	case ZEND_FETCH_CLASS_PARENT:
		scope = ZendGetExecutedScope()
		if scope == nil {
			ZendThrowOrError(fetch_type, nil, "Cannot access parent:: when no class scope is active")
			return nil
		}
		if !(scope.GetParent()) {
			ZendThrowOrError(fetch_type, nil, "Cannot access parent:: when current class scope has no parent")
		}
		return scope.GetParent()
	case ZEND_FETCH_CLASS_STATIC:
		ce = ZendGetCalledScope(CurrEX())
		if ce == nil {
			ZendThrowOrError(fetch_type, nil, "Cannot access static:: when no class scope is active")
			return nil
		}
		return ce
	case ZEND_FETCH_CLASS_AUTO:
		fetch_sub_type = ZendGetClassFetchType(class_name)
		if fetch_sub_type != ZEND_FETCH_CLASS_DEFAULT {
			goto check_fetch_type
		}
	}
	if (fetch_type & ZEND_FETCH_CLASS_NO_AUTOLOAD) != 0 {
		return ZendLookupClassEx(class_name, nil, fetch_type)
	} else if b.Assign(&ce, ZendLookupClassEx(class_name, nil, fetch_type)) == nil {
		if (fetch_type&ZEND_FETCH_CLASS_SILENT) == 0 && EG__().GetException() == nil {
			if fetch_sub_type == ZEND_FETCH_CLASS_INTERFACE {
				ZendThrowOrError(fetch_type, nil, "Interface '%s' not found", class_name.GetVal())
			} else if fetch_sub_type == ZEND_FETCH_CLASS_TRAIT {
				ZendThrowOrError(fetch_type, nil, "Trait '%s' not found", class_name.GetVal())
			} else {
				ZendThrowOrError(fetch_type, nil, "Class '%s' not found", class_name.GetVal())
			}
		}
		return nil
	}
	return ce
}
func ZendFetchClassByName(class_name *types.String, key *types.String, fetch_type int) *types.ClassEntry {
	var ce *types.ClassEntry
	if (fetch_type & ZEND_FETCH_CLASS_NO_AUTOLOAD) != 0 {
		return ZendLookupClassEx(class_name, key, fetch_type)
	} else if b.Assign(&ce, ZendLookupClassEx(class_name, key, fetch_type)) == nil {
		if (fetch_type & ZEND_FETCH_CLASS_SILENT) != 0 {
			return nil
		}
		if EG__().GetException() != nil {
			if (fetch_type & ZEND_FETCH_CLASS_EXCEPTION) == 0 {
				var exception_str *types.String
				var exception_zv types.Zval
				exception_zv.SetObject(EG__().GetException())
				exception_zv.AddRefcount()
				faults.ClearException()
				exception_str = ZvalGetString(&exception_zv)
				faults.ErrorNoreturn(faults.E_ERROR, "During class fetch: Uncaught %s", exception_str.GetVal())
			}
			return nil
		}
		if (fetch_type & ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_INTERFACE {
			ZendThrowOrError(fetch_type, nil, "Interface '%s' not found", class_name.GetVal())
		} else if (fetch_type & ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_TRAIT {
			ZendThrowOrError(fetch_type, nil, "Trait '%s' not found", class_name.GetVal())
		} else {
			ZendThrowOrError(fetch_type, nil, "Class '%s' not found", class_name.GetVal())
		}
		return nil
	}
	return ce
}
func ZendDeleteGlobalVariable(name *types.String) int {
	return types.ZendHashDelInd(EG__().GetSymbolTable(), name.GetStr())
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
	if EG__().GetSymtableCachePtr() > EG__().GetSymtableCache() {
		ex.SetSymbolTable(*(b.PreDec(&(EG__().GetSymtableCachePtr()))))
		symbol_table = ex.GetSymbolTable()
		if ex.GetFunc().GetOpArray().GetLastVar() == 0 {
			return symbol_table
		}
		symbol_table.Extend(ex.GetFunc().GetOpArray().GetLastVar())
	} else {
		ex.SetSymbolTable(types.NewArray(ex.GetFunc().GetOpArray().GetLastVar()))
		symbol_table = ex.GetSymbolTable()
		if ex.GetFunc().GetOpArray().GetLastVar() == 0 {
			return symbol_table
		}
		types.ZendHashRealInitMixed(symbol_table)
	}
	if ex.GetFunc().GetOpArray().GetLastVar() != 0 {
		var str **types.String = ex.GetFunc().GetOpArray().GetVars()
		var end **types.String = str + ex.GetFunc().GetOpArray().GetLastVar()
		var var_ *types.Zval = ex.VarNum(0)
		for {
			types._zendHashAppendInd(symbol_table, *str, var_)
			str++
			var_++
			if str == end {
				break
			}
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
		var str **types.String = op_array.GetVars()
		var end **types.String = str + op_array.GetLastVar()
		var var_ *types.Zval = executeData.VarNum(0)
		for {
			var zv *types.Zval = ht.KeyFind(str.GetStr())
			if zv != nil {
				if zv.IsIndirect() {
					var val *types.Zval = zv.GetZv()
					types.ZVAL_COPY_VALUE(var_, val)
				} else {
					types.ZVAL_COPY_VALUE(var_, zv)
				}
			} else {
				var_.SetUndef()
				zv = ht.KeyAddNew(str.GetStr(), var_)
			}
			zv.SetIndirect(var_)
			str++
			var_++
			if str == end {
				break
			}
		}
	}

	/* copy real values from symbol table into CV slots and create
	   INDIRECT references to CV in symbol table  */
}
func ZendDetachSymbolTable(executeData *ZendExecuteData) {
	var op_array *types.ZendOpArray = executeData.GetFunc().GetOpArray()
	var ht *types.Array = executeData.GetSymbolTable()

	/* copy real values from CV slots into symbol table */

	if op_array.GetLastVar() != 0 {
		var str **types.String = op_array.GetVars()
		var end **types.String = str + op_array.GetLastVar()
		var var_ *types.Zval = executeData.VarNum(0)
		for {
			if var_.IsUndef() {
				types.ZendHashDel(ht, (*str).GetStr())
			} else {
				ht.KeyUpdate(str.GetStr(), var_)
				var_.SetUndef()
			}
			str++
			var_++
			if str == end {
				break
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
			var h ZendUlong = b.HashStr(name)
			var op_array *types.ZendOpArray = executeData.GetFunc().GetOpArray()
			if op_array.GetLastVar() != 0 {
				var str **types.String = op_array.GetVars()
				var end **types.String = str + op_array.GetLastVar()
				for {
					if (*str).GetStr() == name {
						var var_ *types.Zval = executeData.VarNum(str - op_array.GetVars())
						ZvalPtrDtor(var_)
						types.ZVAL_COPY_VALUE(var_, value)
						return types.SUCCESS
					}
					str++
					if str == end {
						break
					}
				}
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
