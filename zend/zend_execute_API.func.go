// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
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
func CleanNonPersistentConstantFull(zv *Zval) int {
	var c *ZendConstant = zv.GetPtr()
	if (ZEND_CONSTANT_FLAGS(c) & CONST_PERSISTENT) != 0 {
		return ZEND_HASH_APPLY_KEEP
	} else {
		return ZEND_HASH_APPLY_REMOVE
	}
}
func CleanNonPersistentFunctionFull(zv *Zval) int {
	var function *ZendFunction = zv.GetPtr()
	if function.GetType() == ZEND_INTERNAL_FUNCTION {
		return ZEND_HASH_APPLY_KEEP
	} else {
		return ZEND_HASH_APPLY_REMOVE
	}
}
func CleanNonPersistentClassFull(zv *Zval) int {
	var ce *ZendClassEntry = zv.GetPtr()
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		return ZEND_HASH_APPLY_KEEP
	} else {
		return ZEND_HASH_APPLY_REMOVE
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
	EG__().SetFunctionTable(CG__().GetFunctionTable())
	EG__().SetClassTable(CG__().GetClassTable())
	EG__().SetInAutoload(nil)
	EG__().SetAutoloadFunc(nil)
	EG__().SetErrorHandling(EH_NORMAL)
	EG__().SetFlags(EG_FLAGS_INITIAL)
	ZendVmStackInit()
	ZendHashInit(EG__().GetSymbolTable(), 64, nil, ZVAL_PTR_DTOR, 0)
	ZendExtensions.Apply(LlistApplyFuncT(ZendExtensionActivator))
	ZendHashInit(EG__().GetIncludedFiles(), 8, nil, nil, 0)
	EG__().SetTicksCount(0)
	EG__().GetUserErrorHandler().SetUndef()
	EG__().GetUserExceptionHandler().SetUndef()
	EG__().SetCurrentExecuteData(nil)
	EG__().GetUserErrorHandlersErrorReporting().Init()
	EG__().GetUserErrorHandlers().Init()
	EG__().GetUserExceptionHandlers().Init()
	ZendObjectsStoreInit(EG__().GetObjectsStore(), 1024)
	EG__().SetFullTablesCleanup(0)
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
	EG__().SetPersistentFunctionsCount(EG__().GetFunctionTable().GetNNumUsed())
	EG__().SetPersistentClassesCount(EG__().GetClassTable().GetNNumUsed())
	ZendWeakrefsInit()
	EG__().SetActive(1)
}
func ZvalCallDestructor(zv *Zval) int {
	if zv.IsIndirect() {
		zv = zv.GetZv()
	}
	if zv.IsObject() && zv.GetRefcount() == 1 {
		return ZEND_HASH_APPLY_REMOVE
	} else {
		return ZEND_HASH_APPLY_KEEP
	}
}
func ZendUncleanZvalPtrDtor(zv *Zval) {
	if zv.IsIndirect() {
		zv = zv.GetZv()
	}
	IZvalPtrDtor(zv)
}
func ZendThrowOrError(fetch_type int, exception_ce *ZendClassEntry, format string, args ...any) {
	message := ZendSprintf(format, args)
	if (fetch_type & ZEND_FETCH_CLASS_EXCEPTION) != 0 {
		ZendThrowError(exception_ce, "%s", message)
	} else {
		ZendError(E_ERROR, "%s", message)
	}
}
func ShutdownDestructors() {
	if CG__().GetUncleanShutdown() != 0 {
		EG__().GetSymbolTable().SetPDestructor(ZendUncleanZvalPtrDtor)
	}
	var __orig_bailout *JMP_BUF = EG__().GetBailout()
	var __bailout JMP_BUF
	EG__().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		var symbols uint32
		for {
			symbols = EG__().GetSymbolTable().GetNNumOfElements()
			ZendHashReverseApply(EG__().GetSymbolTable(), ApplyFuncT(ZvalCallDestructor))
			if symbols == EG__().GetSymbolTable().GetNNumOfElements() {
				break
			}
		}
		ZendObjectsStoreCallDestructors(EG__().GetObjectsStore())
	} else {
		EG__().SetBailout(__orig_bailout)

		/* if we couldn't destruct cleanly, mark all objects as destructed anyway */

		ZendObjectsStoreMarkDestructed(EG__().GetObjectsStore())

		/* if we couldn't destruct cleanly, mark all objects as destructed anyway */

	}
	EG__().SetBailout(__orig_bailout)
}
func ShutdownExecutor() {
	var key *ZendString
	var zv *Zval
	var fast_shutdown ZendBool = IsZendMm() != 0 && EG__().GetFullTablesCleanup() == 0
	var __orig_bailout *JMP_BUF = EG__().bailout
	var __bailout JMP_BUF
	EG__().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		CG__().GetOpenFiles().Destroy()
	}
	EG__().SetBailout(__orig_bailout)
	EG__().SetIsInResourceShutdown(true)
	var __orig_bailout *JMP_BUF = EG__().bailout
	var __bailout JMP_BUF
	EG__().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ZendCloseRsrcList(EG__().GetRegularList())
	}
	EG__().SetBailout(__orig_bailout)

	/* No PHP callback functions should be called after this point. */

	EG__().SetActive(0)
	if fast_shutdown == 0 {
		EG__().GetSymbolTable().GracefulReverseDestroy()

		/* Release static properties and static variables prior to the final GC run,
		 * as they may hold GC roots. */

		var __ht *HashTable = EG__().GetFunctionTable()
		for _, _p := range __ht.foreachDataReserve() {
			var _z Zval = _p.GetVal()

			zv = _z
			var op_array *ZendOpArray = zv.GetPtr()
			if op_array.GetType() == ZEND_INTERNAL_FUNCTION {
				break
			}
			if op_array.GetStaticVariables() != nil {
				var ht *HashTable = ZEND_MAP_PTR_GET(op_array.static_variables_ptr)
				if ht != nil {
					if (ht.GetGcFlags()&IS_ARRAY_IMMUTABLE) == 0 && ht.DelRefcount() == 0 {
						ht.DestroyEx()
					}
					ZEND_MAP_PTR_SET(op_array.static_variables_ptr, nil)
				}
			}
		}
		var __ht__1 *HashTable = EG__().GetClassTable()
		for _, _p := range __ht__1.foreachDataReserve() {
			var _z Zval = _p.GetVal()

			zv = _z
			var ce *ZendClassEntry = zv.GetPtr()
			if ce.GetDefaultStaticMembersCount() != 0 {
				ZendCleanupInternalClassData(ce)
			}
			if ce.IsHasStaticInMethods() {
				var op_array *ZendOpArray
				var __ht *HashTable = ce.GetFunctionTable()
				for _, _p := range __ht.foreachData() {
					var _z *Zval = _p.GetVal()

					op_array = _z.GetPtr()
					if op_array.GetType() == ZEND_USER_FUNCTION {
						if op_array.GetStaticVariables() != nil {
							var ht *HashTable = ZEND_MAP_PTR_GET(op_array.static_variables_ptr)
							if ht != nil {
								if (ht.GetGcFlags()&IS_ARRAY_IMMUTABLE) == 0 && ht.DelRefcount() == 0 {
									ht.DestroyEx()
								}
								ZEND_MAP_PTR_SET(op_array.static_variables_ptr, nil)
							}
						}
					}
				}
			}
		}

		/* Also release error and exception handlers, which may hold objects. */

		if EG__().GetUserErrorHandler().GetType() != IS_UNDEF {
			ZvalPtrDtor(EG__().GetUserErrorHandler())
			EG__().GetUserErrorHandler().SetUndef()
		}
		if EG__().GetUserExceptionHandler().GetType() != IS_UNDEF {
			ZvalPtrDtor(EG__().GetUserExceptionHandler())
			EG__().GetUserExceptionHandler().SetUndef()
		}
		ZendStackClean(EG__().GetUserErrorHandlersErrorReporting(), nil, 1)
		ZendStackClean(EG__().GetUserErrorHandlers(), (func(any))(ZVAL_PTR_DTOR), 1)
		ZendStackClean(EG__().GetUserExceptionHandlers(), (func(any))(ZVAL_PTR_DTOR), 1)
	}
	ZendObjectsStoreFreeObjectStorage(EG__().GetObjectsStore(), fast_shutdown)
	ZendWeakrefsShutdown()
	var __orig_bailout *JMP_BUF = EG__().GetBailout()
	var __bailout JMP_BUF
	EG__().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ZendExtensions.Apply(LlistApplyFuncT(ZendExtensionDeactivator))
	}
	EG__().SetBailout(__orig_bailout)
	if fast_shutdown != 0 {

		/* Fast Request Shutdown
		 * =====================
		 * Zend Memory Manager frees memory by its own. We don't have to free
		 * each allocated block separately.
		 */

		EG__().GetZendConstants().Discard(EG__().GetPersistentConstantsCount())
		EG__().GetFunctionTable().Discard(EG__().GetPersistentFunctionsCount())
		EG__().GetClassTable().Discard(EG__().GetPersistentClassesCount())
		ZendCleanupInternalClasses()
	} else {
		ZendVmStackDestroy()
		if EG__().GetFullTablesCleanup() != 0 {
			ZendHashReverseApply(EG__().GetZendConstants(), CleanNonPersistentConstantFull)
			ZendHashReverseApply(EG__().GetFunctionTable(), CleanNonPersistentFunctionFull)
			ZendHashReverseApply(EG__().GetClassTable(), CleanNonPersistentClassFull)
		} else {
			var __ht *HashTable = EG__().GetZendConstants()
			for _, _p := range __ht.foreachDataReserve() {
				var _z Zval = _p.GetVal()

				key = _p.GetKey()
				zv = _z
				var c *ZendConstant = zv.GetPtr()
				if _idx == EG__().GetPersistentConstantsCount() {
					break
				}
				ZvalPtrDtorNogc(c.GetValue())
				if c.GetName() != nil {
					ZendStringReleaseEx(c.GetName(), 0)
				}
				Efree(c)
				ZendStringReleaseEx(key, 0)
				__ht.GetNNumOfElements()--
				var j uint32 = HT_IDX_TO_HASH(_idx - 1)
				var nIndex uint32 = _p.GetH() | __ht.GetNTableMask()
				var i uint32 = HT_HASH(__ht, nIndex)
				if j != i {
					var prev *Bucket = __ht.Bucket(i)
					for prev.GetVal().GetNext() != j {
						i = prev.GetVal().GetNext()
						prev = __ht.Bucket(i)
					}
					prev.GetVal().GetNext() = _p.GetVal().GetNext()
				} else {
					HT_HASH(__ht, nIndex) = _p.GetVal().GetNext()
				}
			}
			__ht.SetNNumUsed(_idx)
			var __ht__1 *HashTable = EG__().GetFunctionTable()
			for _, _p := range __ht__1.foreachDataReserve() {
				var _z Zval = _p.GetVal()

				key = _p.GetKey()
				zv = _z
				var func_ *ZendFunction = zv.GetPtr()
				if _idx == EG__().GetPersistentFunctionsCount() {
					break
				}
				DestroyOpArray(func_.GetOpArray())
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
			var __ht__2 *HashTable = EG__().GetClassTable()
			for _, _p := range __ht__2.foreachDataReserve() {
				var _z Zval = _p.GetVal()

				key = _p.GetKey()
				zv = _z
				if _idx == EG__().GetPersistentClassesCount() {
					break
				}
				DestroyZendClass(zv)
				ZendStringReleaseEx(key, 0)
				__ht__2.GetNNumOfElements()--
				var j uint32 = HT_IDX_TO_HASH(_idx - 1)
				var nIndex uint32 = _p.GetH() | __ht__2.GetNTableMask()
				var i uint32 = HT_HASH(__ht__2, nIndex)
				if j != i {
					var prev *Bucket = __ht__2.Bucket(i)
					for prev.GetVal().GetNext() != j {
						i = prev.GetVal().GetNext()
						prev = __ht__2.Bucket(i)
					}
					prev.GetVal().GetNext() = _p.GetVal().GetNext()
				} else {
					HT_HASH(__ht__2, nIndex) = _p.GetVal().GetNext()
				}
			}
			__ht__2.SetNNumUsed(_idx)
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
func GetActiveClassName(space **byte) *byte {
	var func_ *ZendFunction
	if ZendIsExecuting() == 0 {
		if space != nil {
			*space = ""
		}
		return ""
	}
	func_ = EG__().GetCurrentExecuteData().GetFunc()
	switch func_.GetType() {
	case ZEND_USER_FUNCTION:
		fallthrough
	case ZEND_INTERNAL_FUNCTION:
		var ce *ZendClassEntry = func_.GetScope()
		if space != nil {
			if ce != nil {
				*space = "::"
			} else {
				*space = ""
			}
		}
		if ce != nil {
			return ce.GetName().GetVal()
		} else {
			return ""
		}
		fallthrough
	default:
		if space != nil {
			*space = ""
		}
		return ""
	}
}
func GetActiveFunctionName() *byte {
	var func_ *ZendFunction
	if ZendIsExecuting() == 0 {
		return nil
	}
	func_ = EG__().GetCurrentExecuteData().GetFunc()
	switch func_.GetType() {
	case ZEND_USER_FUNCTION:
		var function_name *ZendString = func_.GetFunctionName()
		if function_name != nil {
			return function_name.GetVal()
		} else {
			return "main"
		}
	case ZEND_INTERNAL_FUNCTION:
		return func_.GetFunctionName().GetVal()
	default:
		return nil
	}
}
func ZendGetExecutedFilename() string {
	var ex *ZendExecuteData = EG__().GetCurrentExecuteData()
	for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
		ex = ex.GetPrevExecuteData()
	}
	if ex != nil {
		return ex.GetFunc().GetOpArray().GetFilename().GetStr()
	} else {
		return "[no active file]"
	}
}
func ZendGetExecutedFilenameEx() *ZendString {
	var ex *ZendExecuteData = EG__().GetCurrentExecuteData()
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
	var ex *ZendExecuteData = EG__().GetCurrentExecuteData()
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
func ZendGetExecutedScope() *ZendClassEntry {
	var ex *ZendExecuteData = EG__().GetCurrentExecuteData()
	for true {
		if ex == nil {
			return nil
		} else if ex.GetFunc() != nil && (ZEND_USER_CODE(ex.GetFunc().GetType()) || ex.GetFunc().GetScope() != nil) {
			return ex.GetFunc().GetScope()
		}
		ex = ex.GetPrevExecuteData()
	}
}
func ZendIsExecuting() ZendBool {
	return EG__().GetCurrentExecuteData() != 0
}
func ZendUseUndefinedConstant(name *ZendString, attr ZendAstAttr, result *Zval) int {
	var colon *byte
	if EG__().GetException() != nil {
		return FAILURE
	} else if b.Assign(&colon, (*byte)(ZendMemrchr(name.GetVal(), ':', name.GetLen()))) {
		ZendThrowError(nil, "Undefined class constant '%s'", name.GetVal())
		return FAILURE
	} else if (attr & IS_CONSTANT_UNQUALIFIED) == 0 {
		ZendThrowError(nil, "Undefined constant '%s'", name.GetVal())
		return FAILURE
	} else {
		var actual *byte = name.GetVal()
		var actual_len int = name.GetLen()
		var slash *byte = (*byte)(ZendMemrchr(actual, '\\', actual_len))
		if slash != nil {
			actual = slash + 1
			actual_len -= actual - name.GetVal()
		}
		ZendError(E_WARNING, "Use of undefined constant %s - assumed '%s' (this will throw an Error in a future version of PHP)", actual, actual)
		if EG__().GetException() != nil {
			return FAILURE
		} else {
			var result_str *ZendString = ZendStringInit(actual, actual_len, 0)
			ZvalPtrDtorNogc(result)
			result.SetString(result_str)
		}
	}
	return SUCCESS
}
func ZvalUpdateConstantEx(p *Zval, scope *ZendClassEntry) int {
	if p.IsConstant() {
		var ast *ZendAst = Z_ASTVAL_P(p)
		if ast.GetKind() == ZEND_AST_CONSTANT {
			var name *ZendString = ZendAstGetConstantName(ast)
			var zv *Zval = ZendGetConstantEx(name, scope, ast.GetAttr())
			if zv == nil {
				return ZendUseUndefinedConstant(name, ast.GetAttr(), p)
			}
			ZvalPtrDtorNogc(p)
			ZVAL_COPY_OR_DUP(p, zv)
		} else {
			var tmp Zval
			if ZendAstEvaluate(&tmp, ast, scope) != SUCCESS {
				return FAILURE
			}
			ZvalPtrDtorNogc(p)
			ZVAL_COPY_VALUE(p, &tmp)
		}
	}
	return SUCCESS
}
func ZvalUpdateConstant(pp *Zval) int {
	return ZvalUpdateConstantEx(pp, b.CondF(EG__().GetCurrentExecuteData() != nil, func() *ZendClassEntry { return ZendGetExecutedScope() }, func() *ZendClassEntry { return CG__().GetActiveClassEntry() }))
}
func _callUserFunctionEx(
	object *Zval,
	function_name *Zval,
	retval_ptr *Zval,
	param_count uint32,
	params []Zval,
	no_separation int,
) int {
	var fci ZendFcallInfo
	fci.SetSize(b.SizeOf("fci"))
	if object != nil {
		fci.SetObject(object.GetObj())
	} else {
		fci.SetObject(nil)
	}
	ZVAL_COPY_VALUE(fci.GetFunctionName(), function_name)
	fci.SetRetval(retval_ptr)
	fci.SetParamCount(param_count)
	fci.SetParams(params)
	fci.SetNoSeparation(ZendBool(no_separation))
	return ZendCallFunction(&fci, nil)
}
func ZendCallFunction(fci *ZendFcallInfo, fci_cache *ZendFcallInfoCache) int {
	var i uint32
	var call *ZendExecuteData
	var dummy_execute_data ZendExecuteData
	var fci_cache_local ZendFcallInfoCache
	var func_ *ZendFunction
	var call_info uint32
	var object_or_called_scope any
	fci.GetRetval().SetUndef()
	if EG__().GetActive() == 0 {
		return FAILURE
	}
	if EG__().GetException() != nil {
		return FAILURE
	}
	ZEND_ASSERT(fci.GetSize() == b.SizeOf("zend_fcall_info"))

	/* Initialize executeData */

	if EG__().GetCurrentExecuteData() == nil {

		/* This only happens when we're called outside any execute()'s
		 * It shouldn't be strictly necessary to NULL executeData out,
		 * but it may make bugs easier to spot
		 */

		memset(&dummy_execute_data, 0, b.SizeOf("zend_execute_data"))
		EG__().SetCurrentExecuteData(&dummy_execute_data)
	} else if EG__().GetCurrentExecuteData().GetFunc() != nil && ZEND_USER_CODE(EG__().GetCurrentExecuteData().GetFunc().GetCommonType()) && EG__().GetCurrentExecuteData().GetOpline().GetOpcode() != ZEND_DO_FCALL && EG__().GetCurrentExecuteData().GetOpline().GetOpcode() != ZEND_DO_ICALL && EG__().GetCurrentExecuteData().GetOpline().GetOpcode() != ZEND_DO_UCALL && EG__().GetCurrentExecuteData().GetOpline().GetOpcode() != ZEND_DO_FCALL_BY_NAME {

		/* Insert fake frame in case of include or magic calls */

		dummy_execute_data = (*EG__)().current_execute_data
		dummy_execute_data.SetPrevExecuteData(EG__().GetCurrentExecuteData())
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
				var callable_name *ZendString = ZendGetCallableNameEx(fci.GetFunctionName(), fci.GetObject())
				ZendError(E_WARNING, "Invalid callback %s, %s", callable_name.GetVal(), error)
				Efree(error)
				ZendStringReleaseEx(callable_name, 0)
			}
			if EG__().GetCurrentExecuteData() == &dummy_execute_data {
				EG__().SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
			}
			return FAILURE
		} else if error != nil {

			/* Capitalize the first latter of the error message */

			if error[0] >= 'a' && error[0] <= 'z' {
				error[0] += 'A' - 'a'
			}
			ZendError(E_DEPRECATED, "%s", error)
			Efree(error)
			if EG__().GetException() != nil {
				if EG__().GetCurrentExecuteData() == &dummy_execute_data {
					EG__().SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
				}
				return FAILURE
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
		ZendError(E_DEPRECATED, "Function %s%s%s() is deprecated", b.CondF1(func_.GetScope() != nil, func() []byte { return func_.GetScope().GetName().GetVal() }, ""), b.Cond(func_.GetScope() != nil, "::", ""), func_.GetFunctionName().GetVal())
		if EG__().GetException() != nil {
			ZendVmStackFreeCallFrame(call)
			if EG__().GetCurrentExecuteData() == &dummy_execute_data {
				EG__().SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
				ZendRethrowException(EG__().GetCurrentExecuteData())
			}
			return FAILURE
		}
	}
	for i = 0; i < fci.GetParamCount(); i++ {
		var param *Zval
		var arg *Zval = fci.GetParams()[i]
		var must_wrap ZendBool = 0
		if ARG_SHOULD_BE_SENT_BY_REF(func_, i+1) != 0 {
			if !(arg.IsReference()) {
				if fci.GetNoSeparation() == 0 {

					/* Separation is enabled -- create a ref */

					arg.SetNewRef(arg)

					/* Separation is enabled -- create a ref */

				} else if ARG_MAY_BE_SENT_BY_REF(func_, i+1) == 0 {

					/* By-value send is not allowed -- emit a warning,
					 * and perform the call with the value wrapped in a reference. */

					ZendError(E_WARNING, "Parameter %d to %s%s%s() expected to be a reference, value given", i+1, b.CondF1(func_.GetScope() != nil, func() []byte { return func_.GetScope().GetName().GetVal() }, ""), b.Cond(func_.GetScope() != nil, "::", ""), func_.GetFunctionName().GetVal())
					must_wrap = 1
					if EG__().GetException() != nil {
						ZEND_CALL_NUM_ARGS(call) = i
						ZendVmStackFreeArgs(call)
						ZendVmStackFreeCallFrame(call)
						if EG__().GetCurrentExecuteData() == &dummy_execute_data {
							EG__().SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
						}
						return FAILURE
					}
				}
			}
		} else {
			if arg.IsReference() && !func_.IsCallViaTrampoline() {

				/* don't separate references for __call */

				arg = Z_REFVAL_P(arg)

				/* don't separate references for __call */

			}
		}
		param = ZEND_CALL_ARG(call, i+1)
		if must_wrap == 0 {
			ZVAL_COPY(param, arg)
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
		call.SetPrevExecuteData(EG__().GetCurrentExecuteData())
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
			call.SetPrevExecuteData(EG__().GetCurrentExecuteData())
			EG__().SetCurrentExecuteData(call)
			fci.GetObject().GetHandlers().GetCallMethod()(func_.GetFunctionName(), fci.GetObject(), call, fci.GetRetval())
			EG__().SetCurrentExecuteData(call.GetPrevExecuteData())
		} else {
			ZendThrowError(nil, "Cannot call overloaded function for non-object")
		}
		ZendVmStackFreeArgs(call)
		if func_.GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY {
			ZendStringReleaseEx(func_.GetFunctionName(), 0)
		}
		Efree(func_)
		if EG__().GetException() != nil {
			ZvalPtrDtor(fci.GetRetval())
			fci.GetRetval().SetUndef()
		}
	}
	ZendVmStackFreeCallFrame(call)
	if EG__().GetCurrentExecuteData() == &dummy_execute_data {
		EG__().SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
	}
	if EG__().GetException() != nil {
		if EG__().GetCurrentExecuteData() == nil {
			ZendThrowExceptionInternal(nil)
		} else if EG__().GetCurrentExecuteData().GetFunc() != nil && ZEND_USER_CODE(EG__().GetCurrentExecuteData().GetFunc().GetCommonType()) {
			ZendRethrowException(EG__().GetCurrentExecuteData())
		}
	}
	return SUCCESS
}
func ZendLookupClassEx(name *ZendString, key *ZendString, flags uint32) *ZendClassEntry {
	var ce *ZendClassEntry = nil
	var args []Zval
	var zv *Zval
	var local_retval Zval
	var lc_name *ZendString
	var fcall_info ZendFcallInfo
	var fcall_cache ZendFcallInfoCache
	var orig_fake_scope *ZendClassEntry
	if key != nil {
		lc_name = key
	} else {
		if name == nil || name.GetLen() == 0 {
			return nil
		}
		if name.GetVal()[0] == '\\' {
			lc_name = ZendStringAlloc(name.GetLen()-1, 0)
			ZendStrTolowerCopy(lc_name.GetVal(), name.GetVal()+1, name.GetLen()-1)
		} else {
			lc_name = ZendStringTolower(name)
		}
	}
	zv = EG__().GetClassTable().KeyFind(lc_name.GetStr())
	if zv != nil {
		if key == nil {
			ZendStringReleaseEx(lc_name, 0)
		}
		ce = (*ZendClassEntry)(zv.GetPtr())
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
			ZendStringReleaseEx(lc_name, 0)
		}
		return nil
	}
	if EG__().GetAutoloadFunc() == nil {
		var func_ *ZendFunction = ZendFetchFunction(ZSTR_KNOWN(ZEND_STR_MAGIC_AUTOLOAD))
		if func_ != nil {
			EG__().SetAutoloadFunc(func_)
		} else {
			if key == nil {
				ZendStringReleaseEx(lc_name, 0)
			}
			return nil
		}
	}

	/* Verify class name before passing it to __autoload() */

	if key == nil && strspn(name.GetVal(), "0123456789_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ200201202203204205206207210211212213214215216217220221222223224225226227230231232233234235236237240241242243244245246247250251252253254255256257260261262263264265266267270271272273274275276277300301302303304305306307310311312313314315316317320321322323324325326327330331332333334335336337340341342343344345346347350351352353354355356357360361362363364365366367370371372373374375376377\\") != name.GetLen() {
		ZendStringReleaseEx(lc_name, 0)
		return nil
	}
	if EG__().GetInAutoload() == nil {
		ALLOC_HASHTABLE(EG__().GetInAutoload())
		ZendHashInit(EG__().GetInAutoload(), 8, nil, nil, 0)
	}
	if ZendHashAddEmptyElement(EG__().GetInAutoload(), lc_name) == nil {
		if key == nil {
			ZendStringReleaseEx(lc_name, 0)
		}
		return nil
	}
	local_retval.SetUndef()
	if name.GetVal()[0] == '\\' {
		ZVAL_STRINGL(&args[0], name.GetVal()+1, name.GetLen()-1)
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
	ZendExceptionSave()
	if ZendCallFunction(&fcall_info, &fcall_cache) == SUCCESS && EG__().GetException() == nil {
		ce = ZendHashFindPtr(EG__().GetClassTable(), lc_name)
	}
	ZendExceptionRestore()
	EG__().SetFakeScope(orig_fake_scope)
	ZvalPtrDtor(&args[0])
	ZvalPtrDtorStr(fcall_info.GetFunctionName())
	ZendHashDel(EG__().GetInAutoload(), lc_name)
	ZvalPtrDtor(&local_retval)
	if key == nil {
		ZendStringReleaseEx(lc_name, 0)
	}
	return ce
}
func ZendLookupClass(name *ZendString) *ZendClassEntry { return ZendLookupClassEx(name, nil, 0) }
func ZendGetCalledScope(ex *ZendExecuteData) *ZendClassEntry {
	for ex != nil {
		if ex.GetThis().IsObject() {
			return Z_OBJCE(ex.GetThis())
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
func ZendGetThisObject(ex *ZendExecuteData) *ZendObject {
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
func ZendEvalStringl(str *byte, str_len int, retval_ptr *Zval, string_name *byte) int {
	var pv Zval
	var new_op_array *ZendOpArray
	var original_compiler_options uint32
	var retval int
	if retval_ptr != nil {
		pv.SetString(ZendStringAlloc(str_len+b.SizeOf("\"return ;\"")-1, 0))
		memcpy(Z_STRVAL(pv), "return ", b.SizeOf("\"return \"")-1)
		memcpy(Z_STRVAL(pv)+b.SizeOf("\"return \"")-1, str, str_len)
		Z_STRVAL(pv)[Z_STRLEN(pv)-1] = ';'
		Z_STRVAL(pv)[Z_STRLEN(pv)] = '0'
	} else {
		ZVAL_STRINGL(&pv, str, str_len)
	}

	/*printf("Evaluating '%s'\n", pv.value.str.val);*/

	original_compiler_options = CG__().GetCompilerOptions()
	CG__().SetCompilerOptions(ZEND_COMPILE_DEFAULT_FOR_EVAL)
	new_op_array = ZendCompileString(&pv, string_name)
	CG__().SetCompilerOptions(original_compiler_options)
	if new_op_array != nil {
		var local_retval Zval
		EG__().SetNoExtensions(1)
		new_op_array.SetScope(ZendGetExecutedScope())
		var __orig_bailout *JMP_BUF = EG__().GetBailout()
		var __bailout JMP_BUF
		EG__().SetBailout(&__bailout)
		if SETJMP(__bailout) == 0 {
			local_retval.SetUndef()
			ZendExecute(new_op_array, &local_retval)
		} else {
			EG__().SetBailout(__orig_bailout)
			DestroyOpArray(new_op_array)
			EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
			ZendBailout()
		}
		EG__().SetBailout(__orig_bailout)
		if local_retval.GetType() != IS_UNDEF {
			if retval_ptr != nil {
				ZVAL_COPY_VALUE(retval_ptr, &local_retval)
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
		retval = SUCCESS
	} else {
		retval = FAILURE
	}
	ZvalPtrDtorStr(&pv)
	return retval
}
func ZendEvalString(str *byte, retval_ptr *Zval, string_name *byte) int {
	return ZendEvalStringl(str, strlen(str), retval_ptr, string_name)
}
func ZendEvalStringlEx(str *byte, str_len int, retval_ptr *Zval, string_name *byte, handle_exceptions int) int {
	var result int
	result = ZendEvalStringl(str, str_len, retval_ptr, string_name)
	if handle_exceptions != 0 && EG__().GetException() != nil {
		ZendExceptionError(EG__().GetException(), E_ERROR)
		result = FAILURE
	}
	return result
}
func ZendEvalStringEx(str *byte, retval_ptr *Zval, string_name string, handle_exceptions int) int {
	return ZendEvalStringlEx(str, strlen(str), retval_ptr, string_name, handle_exceptions)
}
func ZendTimeout(dummy int) {
	EG__().SetTimedOut(0)
	ZendSetTimeoutEx(0, 1)
	ZendErrorNoreturn(E_ERROR, "Maximum execution time of "+ZEND_LONG_FMT+" second%s exceeded", EG__().GetTimeoutSeconds(), b.Cond(EG__().GetTimeoutSeconds() == 1, "", "s"))
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
			ZendQuietWrite(2, log_buffer, MIN(output_len, b.SizeOf("log_buffer")))
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
func ZendFetchClass(class_name *ZendString, fetch_type int) *ZendClassEntry {
	var ce *ZendClassEntry
	var scope *ZendClassEntry
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
		ce = ZendGetCalledScope(EG__().GetCurrentExecuteData())
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
func ZendFetchClassByName(class_name *ZendString, key *ZendString, fetch_type int) *ZendClassEntry {
	var ce *ZendClassEntry
	if (fetch_type & ZEND_FETCH_CLASS_NO_AUTOLOAD) != 0 {
		return ZendLookupClassEx(class_name, key, fetch_type)
	} else if b.Assign(&ce, ZendLookupClassEx(class_name, key, fetch_type)) == nil {
		if (fetch_type & ZEND_FETCH_CLASS_SILENT) != 0 {
			return nil
		}
		if EG__().GetException() != nil {
			if (fetch_type & ZEND_FETCH_CLASS_EXCEPTION) == 0 {
				var exception_str *ZendString
				var exception_zv Zval
				exception_zv.SetObject(EG__().GetException())
				exception_zv.AddRefcount()
				ZendClearException()
				exception_str = ZvalGetString(&exception_zv)
				ZendErrorNoreturn(E_ERROR, "During class fetch: Uncaught %s", exception_str.GetVal())
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
func ZendDeleteGlobalVariable(name *ZendString) int {
	return ZendHashDelInd(EG__().GetSymbolTable(), name)
}
func ZendRebuildSymbolTable() *ZendArray {
	var ex *ZendExecuteData
	var symbol_table *ZendArray

	/* Search for last called user function */

	ex = EG__().GetCurrentExecuteData()
	for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetCommonType()))) {
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
		ex.SetSymbolTable(ZendNewArray(ex.GetFunc().GetOpArray().GetLastVar()))
		symbol_table = ex.GetSymbolTable()
		if ex.GetFunc().GetOpArray().GetLastVar() == 0 {
			return symbol_table
		}
		ZendHashRealInitMixed(symbol_table)
	}
	if ex.GetFunc().GetOpArray().GetLastVar() != 0 {
		var str **ZendString = ex.GetFunc().GetOpArray().GetVars()
		var end **ZendString = str + ex.GetFunc().GetOpArray().GetLastVar()
		var var_ *Zval = ZEND_CALL_VAR_NUM(ex, 0)
		for {
			_zendHashAppendInd(symbol_table, *str, var_)
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
	var op_array *ZendOpArray = executeData.GetFunc().GetOpArray()
	var ht *HashTable = executeData.GetSymbolTable()

	/* copy real values from symbol table into CV slots and create
	   INDIRECT references to CV in symbol table  */

	if op_array.GetLastVar() != 0 {
		var str **ZendString = op_array.GetVars()
		var end **ZendString = str + op_array.GetLastVar()
		var var_ *Zval = EX_VAR_NUM(0)
		for {
			var zv *Zval = ht.KeyFind(str.GetStr())
			if zv != nil {
				if zv.IsIndirect() {
					var val *Zval = zv.GetZv()
					ZVAL_COPY_VALUE(var_, val)
				} else {
					ZVAL_COPY_VALUE(var_, zv)
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
	var op_array *ZendOpArray = executeData.GetFunc().GetOpArray()
	var ht *HashTable = executeData.GetSymbolTable()

	/* copy real values from CV slots into symbol table */

	if op_array.GetLastVar() != 0 {
		var str **ZendString = op_array.GetVars()
		var end **ZendString = str + op_array.GetLastVar()
		var var_ *Zval = EX_VAR_NUM(0)
		for {
			if var_.IsUndef() {
				ZendHashDel(ht, *str)
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
func ZendSetLocalVar(name *ZendString, value *Zval, force int) int {
	var executeData *ZendExecuteData = EG__().GetCurrentExecuteData()
	for executeData != nil && (executeData.GetFunc() == nil || !(ZEND_USER_CODE(executeData.GetFunc().GetCommonType()))) {
		executeData = executeData.GetPrevExecuteData()
	}
	if executeData != nil {
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) == 0 {
			var h ZendUlong = name.GetHash()
			var op_array *ZendOpArray = executeData.GetFunc().GetOpArray()
			if op_array.GetLastVar() != 0 {
				var str **ZendString = op_array.GetVars()
				var end **ZendString = str + op_array.GetLastVar()
				for {
					if str.GetH() == h && ZendStringEqualContent(*str, name) != 0 {
						var var_ *Zval = EX_VAR_NUM(str - op_array.GetVars())
						ZVAL_COPY_VALUE(var_, value)
						return SUCCESS
					}
					str++
					if str == end {
						break
					}
				}
			}
			if force != 0 {
				var symbol_table *ZendArray = ZendRebuildSymbolTable()
				if symbol_table != nil {
					symbol_table.KeyUpdate(name.GetStr(), value)
					return SUCCESS
				}
			}
		} else {
			executeData.GetSymbolTable().KeyUpdateIndirect(name.GetStr(), value)
			return SUCCESS
		}
	}
	return FAILURE
}
func ZendSetLocalVarStr(name string, len_ int, value *Zval, force int) int {
	var executeData *ZendExecuteData = EG__().GetCurrentExecuteData()
	for executeData != nil && (executeData.GetFunc() == nil || !(ZEND_USER_CODE(executeData.GetFunc().GetCommonType()))) {
		executeData = executeData.GetPrevExecuteData()
	}
	if executeData != nil {
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) == 0 {
			var h ZendUlong = ZendHashFunc(name, len_)
			var op_array *ZendOpArray = executeData.GetFunc().GetOpArray()
			if op_array.GetLastVar() != 0 {
				var str **ZendString = op_array.GetVars()
				var end **ZendString = str + op_array.GetLastVar()
				for {
					if str.GetH() == h && str.GetLen() == len_ && memcmp(str.GetVal(), name, len_) == 0 {
						var var_ *Zval = EX_VAR_NUM(str - op_array.GetVars())
						ZvalPtrDtor(var_)
						ZVAL_COPY_VALUE(var_, value)
						return SUCCESS
					}
					str++
					if str == end {
						break
					}
				}
			}
			if force != 0 {
				var symbol_table *ZendArray = ZendRebuildSymbolTable()
				if symbol_table != nil {
					symbol_table.KeyUpdate(b.CastStr(name, len_), value)
					return SUCCESS
				}
			}
		} else {
			executeData.GetSymbolTable().KeyUpdateIndirect(b.CastStr(name, len_), value)
			return SUCCESS
		}
	}
	return FAILURE
}
