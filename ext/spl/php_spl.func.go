package spl

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func SPL_G(v __auto__) __auto__ { return SplGlobals.v }
func ZmGlobalsCtorSpl(spl_globals *ZendSplGlobals) {
	spl_globals.SetAutoloadExtensions(nil)
	spl_globals.SetAutoloadFunctions(nil)
	spl_globals.SetAutoloadRunning(0)
}
func SplFindCeByName(name *types.String, autoload types.ZendBool) *types.ClassEntry {
	var ce *types.ClassEntry
	if autoload == 0 {
		var lc_name *types.String = zend.ZendStringTolower(name)
		ce = types.ZendHashFindPtr(zend.EG__().GetClassTable(), lc_name.GetStr())
		types.ZendStringRelease(lc_name)
	} else {
		ce = zend.ZendLookupClass(name)
	}
	if ce == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Class %s does not exist%s", name.GetVal(), b.Cond(autoload != 0, " and could not be loaded", ""))
		return nil
	}
	return ce
}
func ZifClassParents(executeData zpp.Ex, return_value zpp.Ret, instance *types.Zval, _ zpp.Opt, autoload *types.Zval) {
	var obj *types.Zval
	var parent_class *types.ClassEntry
	var ce *types.ClassEntry
	var autoload types.ZendBool = 1
	if zend.ZendParseParameters(executeData.NumArgs(), "z|b", &obj, &autoload) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	if obj.GetType() != types.IS_OBJECT && obj.GetType() != types.IS_STRING {
		core.PhpErrorDocref(nil, faults.E_WARNING, "object or string expected")
		return_value.SetFalse()
		return
	}
	if obj.IsType(types.IS_STRING) {
		if nil == b.Assign(&ce, SplFindCeByName(obj.GetStr(), autoload)) {
			return_value.SetFalse()
			return
		}
	} else {
		ce = types.Z_OBJCE_P(obj)
	}
	zend.ArrayInit(return_value)
	parent_class = ce.GetParent()
	for parent_class != nil {
		SplAddClassName(return_value, parent_class, 0, 0)
		parent_class = parent_class.GetParent()
	}
}
func ZifClassImplements(executeData zpp.Ex, return_value zpp.Ret, what *types.Zval, _ zpp.Opt, autoload *types.Zval) {
	var obj *types.Zval
	var autoload types.ZendBool = 1
	var ce *types.ClassEntry
	if zend.ZendParseParameters(executeData.NumArgs(), "z|b", &obj, &autoload) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	if obj.GetType() != types.IS_OBJECT && obj.GetType() != types.IS_STRING {
		core.PhpErrorDocref(nil, faults.E_WARNING, "object or string expected")
		return_value.SetFalse()
		return
	}
	if obj.IsType(types.IS_STRING) {
		if nil == b.Assign(&ce, SplFindCeByName(obj.GetStr(), autoload)) {
			return_value.SetFalse()
			return
		}
	} else {
		ce = types.Z_OBJCE_P(obj)
	}
	zend.ArrayInit(return_value)
	SplAddInterfaces(return_value, ce, 1, zend.ZEND_ACC_INTERFACE)
}
func ZifClassUses(executeData zpp.Ex, return_value zpp.Ret, what *types.Zval, _ zpp.Opt, autoload *types.Zval) {
	var obj *types.Zval
	var autoload types.ZendBool = 1
	var ce *types.ClassEntry
	if zend.ZendParseParameters(executeData.NumArgs(), "z|b", &obj, &autoload) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	if obj.GetType() != types.IS_OBJECT && obj.GetType() != types.IS_STRING {
		core.PhpErrorDocref(nil, faults.E_WARNING, "object or string expected")
		return_value.SetFalse()
		return
	}
	if obj.IsType(types.IS_STRING) {
		if nil == b.Assign(&ce, SplFindCeByName(obj.GetStr(), autoload)) {
			return_value.SetFalse()
			return
		}
	} else {
		ce = types.Z_OBJCE_P(obj)
	}
	zend.ArrayInit(return_value)
	SplAddTraits(return_value, ce, 1, zend.ZEND_ACC_TRAIT)
}
func ZifSplClasses(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	zend.ArrayInit(return_value)
	SplAddClasses(spl_ce_AppendIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_ArrayIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_ArrayObject, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_BadFunctionCallException, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_BadMethodCallException, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_CachingIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_CallbackFilterIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_DirectoryIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_DomainException, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_EmptyIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_FilesystemIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_FilterIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_GlobIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_InfiniteIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_InvalidArgumentException, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_IteratorIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_LengthException, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_LimitIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_LogicException, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_MultipleIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_NoRewindIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_OuterIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_OutOfBoundsException, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_OutOfRangeException, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_OverflowException, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_ParentIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_RangeException, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_RecursiveArrayIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_RecursiveCachingIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_RecursiveCallbackFilterIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_RecursiveDirectoryIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_RecursiveFilterIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_RecursiveIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_RecursiveIteratorIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_RecursiveRegexIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_RecursiveTreeIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_RegexIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_RuntimeException, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SeekableIterator, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplDoublyLinkedList, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplFileInfo, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplFileObject, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplFixedArray, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplHeap, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplMinHeap, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplMaxHeap, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplObjectStorage, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplObserver, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplPriorityQueue, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplQueue, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplStack, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplSubject, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_SplTempFileObject, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_UnderflowException, return_value, 0, 0, 0)
	SplAddClasses(spl_ce_UnexpectedValueException, return_value, 0, 0, 0)
}
func SplAutoload(class_name *types.String, lc_name *types.String, ext *byte, ext_len int) int {
	var class_file *byte
	var class_file_len int
	var dummy types.Zval
	var file_handle zend.ZendFileHandle
	var new_op_array *zend.ZendOpArray
	var result types.Zval
	var ret int
	class_file_len = int(core.Spprintf(&class_file, 0, "%s%.*s", lc_name.GetVal(), ext_len, ext))
	ret = core.PhpStreamOpenForZendEx(class_file, &file_handle, core.USE_PATH|core.STREAM_OPEN_FOR_INCLUDE)
	if ret == types.SUCCESS {
		var opened_path *types.String
		if file_handle.GetOpenedPath() == nil {
			file_handle.SetOpenedPath(types.NewString(b.CastStr(class_file, class_file_len)))
		}
		opened_path = file_handle.GetOpenedPath().Copy()
		dummy.SetNull()
		if zend.EG__().GetIncludedFiles().KeyAdd(opened_path.GetStr(), &dummy) != nil {
			new_op_array = zend.ZendCompileFile(&file_handle, zend.ZEND_REQUIRE)
			zend.ZendDestroyFileHandle(&file_handle)
		} else {
			new_op_array = nil
			file_handle.Destroy()
		}
		types.ZendStringReleaseEx(opened_path, 0)
		if new_op_array != nil {
			result.SetUndef()
			zend.ZendExecute(new_op_array, &result)
			zend.DestroyOpArray(new_op_array)
			zend.Efree(new_op_array)
			if zend.EG__().GetException() == nil {
				zend.ZvalPtrDtor(&result)
			}
			zend.Efree(class_file)
			return types.IntBool(zend.EG__().GetClassTable().KeyExists(lc_name.GetStr()))
		}
	}
	zend.Efree(class_file)
	return 0
}
func ZifSplAutoload(executeData zpp.Ex, return_value zpp.Ret, className *types.Zval, _ zpp.Opt, fileExtensions *types.Zval) {
	var pos_len int
	var pos1_len int
	var pos *byte
	var pos1 *byte
	var class_name *types.String
	var lc_name *types.String
	var file_exts *types.String = SPL_G(autoload_extensions)
	if zend.ZendParseParameters(executeData.NumArgs(), "S|S", &class_name, &file_exts) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	if file_exts == nil {
		pos = SPL_DEFAULT_FILE_EXTENSIONS
		pos_len = b.SizeOf("SPL_DEFAULT_FILE_EXTENSIONS") - 1
	} else {
		pos = file_exts.GetVal()
		pos_len = int(file_exts.GetLen())
	}
	lc_name = zend.ZendStringTolower(class_name)
	for pos != nil && (*pos) && zend.EG__().GetException() == nil {
		pos1 = strchr(pos, ',')
		if pos1 != nil {
			pos1_len = int(pos1 - pos)
		} else {
			pos1_len = pos_len
		}
		if SplAutoload(class_name, lc_name, pos, pos1_len) != 0 {
			break
		}
		if pos1 != nil {
			pos = pos1 + 1
		} else {
			pos = nil
		}
		if pos1 != nil {
			pos_len = pos_len - pos1_len - 1
		} else {
			pos_len = 0
		}
	}
	types.ZendStringRelease(lc_name)
}
func ZifSplAutoloadExtensions(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, fileExtensions *types.Zval) {
	var file_exts *types.String = nil
	if zend.ZendParseParameters(executeData.NumArgs(), "|S", &file_exts) == types.FAILURE {
		return
	}
	if file_exts != nil {
		if SPL_G(autoload_extensions) {
			types.ZendStringReleaseEx(SPL_G(autoload_extensions), 0)
		}
		SPL_G(autoload_extensions) = file_exts.Copy()
	}
	if SPL_G(autoload_extensions) == nil {
		return_value.SetRawString(SPL_DEFAULT_FILE_EXTENSIONS)
		return
	} else {
		SPL_G(autoload_extensions).AddRefcount()
		return_value.SetString(SPL_G(autoload_extensions))
		return
	}
}
func AutoloadFuncInfoDtor(element *types.Zval) {
	var alfi *AutoloadFuncInfo = (*AutoloadFuncInfo)(element.GetPtr())
	if !(alfi.GetObj().IsUndef()) {
		zend.ZvalPtrDtor(alfi.GetObj())
	}
	if alfi.GetFuncPtr() != nil && alfi.GetFuncPtr().HasFnFlags(zend.ZEND_ACC_CALL_VIA_TRAMPOLINE) {
		types.ZendStringReleaseEx(alfi.GetFuncPtr().GetFunctionName(), 0)
		zend.ZendFreeTrampoline(alfi.GetFuncPtr())
	}
	if !(alfi.GetClosure().IsUndef()) {
		zend.ZvalPtrDtor(alfi.GetClosure())
	}
	zend.Efree(alfi)
}
func ZifSplAutoloadCall(executeData zpp.Ex, return_value zpp.Ret, className *types.Zval) {
	var class_name *types.Zval
	var retval types.Zval
	var lc_name *types.String
	var func_name *types.String
	var alfi *AutoloadFuncInfo
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &class_name) == types.FAILURE || class_name.GetType() != types.IS_STRING {
		return
	}
	if SPL_G(autoload_functions) {
		var pos types.ArrayPosition
		var num_idx zend.ZendUlong
		var func_ *zend.ZendFunction
		var fci types.ZendFcallInfo
		var fcic types.ZendFcallInfoCache
		var called_scope *types.ClassEntry = zend.ZendGetCalledScope(executeData)
		var l_autoload_running int = SPL_G(autoload_running)
		SPL_G(autoload_running) = 1
		lc_name = zend.ZendStringTolower(class_name.GetStr())
		fci.SetSize(b.SizeOf("fci"))
		fci.SetRetval(&retval)
		fci.SetParamCount(1)
		fci.SetParams(class_name)
		fci.SetNoSeparation(1)
		fci.GetFunctionName().SetUndef()
		types.ZendHashInternalPointerResetEx(SPL_G(autoload_functions), &pos)
		for types.ZendHashGetCurrentKeyEx(SPL_G(autoload_functions), &func_name, &num_idx, &pos) == types.HASH_KEY_IS_STRING {
			alfi = types.ZendHashGetCurrentDataPtrEx(SPL_G(autoload_functions), &pos)
			func_ = alfi.GetFuncPtr()
			if func_.HasFnFlags(zend.ZEND_ACC_CALL_VIA_TRAMPOLINE) {
				func_ = zend.Emalloc(b.SizeOf("zend_op_array"))
				memcpy(func_, alfi.GetFuncPtr(), b.SizeOf("zend_op_array"))
				func_.GetOpArray().GetFunctionName().AddRefcount()
			}
			retval.SetUndef()
			fcic.SetFunctionHandler(func_)
			if alfi.GetObj().IsUndef() {
				fci.SetObject(nil)
				fcic.SetObject(nil)
				if alfi.GetCe() != nil && (called_scope == nil || zend.InstanceofFunction(called_scope, alfi.GetCe()) == 0) {
					fcic.SetCalledScope(alfi.GetCe())
				} else {
					fcic.SetCalledScope(called_scope)
				}
			} else {
				fci.SetObject(alfi.GetObj().GetObj())
				fcic.SetObject(alfi.GetObj().GetObj())
				fcic.SetCalledScope(types.Z_OBJCE(alfi.GetObj()))
			}
			zend.ZendCallFunction(&fci, &fcic)
			zend.ZvalPtrDtor(&retval)
			if zend.EG__().GetException() != nil {
				break
			}
			if pos+1 == SPL_G(autoload_functions).nNumUsed || zend.EG__().GetClassTable().KeyExists(lc_name.GetStr()) {
				break
			}
			types.ZendHashMoveForwardEx(SPL_G(autoload_functions), &pos)
		}
		types.ZendStringReleaseEx(lc_name, 0)
		SPL_G(autoload_running) = l_autoload_running
	} else {

		/* do not use or overwrite &EG(autoload_func) here */

		var fcall_info types.ZendFcallInfo
		var fcall_cache types.ZendFcallInfoCache
		retval.SetUndef()
		fcall_info.SetSize(b.SizeOf("fcall_info"))
		fcall_info.GetFunctionName().SetUndef()
		fcall_info.SetRetval(&retval)
		fcall_info.SetParamCount(1)
		fcall_info.SetParams(class_name)
		fcall_info.SetObject(nil)
		fcall_info.SetNoSeparation(1)
		fcall_cache.SetFunctionHandler(SplAutoloadFn)
		fcall_cache.SetCalledScope(nil)
		fcall_cache.SetObject(nil)
		zend.ZendCallFunction(&fcall_info, &fcall_cache)
		zend.ZvalPtrDtor(&retval)
	}
}
func HT_MOVE_TAIL_TO_HEAD(ht *types.Array) {
	var tmp types.Bucket = ht.GetArData()[ht.GetNNumUsed()-1]
	memmove(ht.GetArData()+1, ht.GetArData(), b.SizeOf("Bucket")*(ht.GetNNumUsed()-1))
	ht.GetArData()[0] = tmp
	ht.Rehash()
}
func ZifSplAutoloadRegister(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, autoloadFunction *types.Zval, throw *types.Zval, prepend *types.Zval) {
	var func_name *types.String
	var error *byte = nil
	var lc_name *types.String
	var zcallable *types.Zval = nil
	var do_throw types.ZendBool = 1
	var prepend types.ZendBool = 0
	var spl_func_ptr *zend.ZendFunction
	var alfi AutoloadFuncInfo
	var obj_ptr *types.ZendObject
	var fcc types.ZendFcallInfoCache
	if zend.ZendParseParametersEx(zpp.FlagQuiet, executeData.NumArgs(), "|zbb", &zcallable, &do_throw, &prepend) == types.FAILURE {
		return
	}
	if executeData.NumArgs() != 0 {
		if zend.ZendIsCallableEx(zcallable, nil, zend.IS_CALLABLE_STRICT, &func_name, &fcc, &error) == 0 {
			alfi.SetCe(fcc.GetCallingScope())
			alfi.SetFuncPtr(fcc.GetFunctionHandler())
			obj_ptr = fcc.GetObject()
			if zcallable.IsType(types.IS_ARRAY) {
				if obj_ptr == nil && alfi.GetFuncPtr() != nil && !alfi.GetFuncPtr().HasFnFlags(zend.ZEND_ACC_STATIC) {
					if do_throw != 0 {
						faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Passed array specifies a non static method but no object (%s)", error)
					}
					if error != nil {
						zend.Efree(error)
					}
					types.ZendStringReleaseEx(func_name, 0)
					return_value.SetFalse()
					return
				} else if do_throw != 0 {
					faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Passed array does not specify %s %smethod (%s)", b.Cond(alfi.GetFuncPtr() != nil, "a callable", "an existing"), b.Cond(obj_ptr == nil, "static ", ""), error)
				}
				if error != nil {
					zend.Efree(error)
				}
				types.ZendStringReleaseEx(func_name, 0)
				return_value.SetFalse()
				return
			} else if zcallable.IsType(types.IS_STRING) {
				if do_throw != 0 {
					faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Function '%s' not %s (%s)", func_name.GetVal(), b.Cond(alfi.GetFuncPtr() != nil, "callable", "found"), error)
				}
				if error != nil {
					zend.Efree(error)
				}
				types.ZendStringReleaseEx(func_name, 0)
				return_value.SetFalse()
				return
			} else {
				if do_throw != 0 {
					faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Illegal value passed (%s)", error)
				}
				if error != nil {
					zend.Efree(error)
				}
				types.ZendStringReleaseEx(func_name, 0)
				return_value.SetFalse()
				return
			}
		} else if fcc.GetFunctionHandler().GetType() == zend.ZEND_INTERNAL_FUNCTION && fcc.GetFunctionHandler().GetInternalFunction().GetHandler() == ZifSplAutoloadCall {
			if do_throw != 0 {
				faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Function spl_autoload_call() cannot be registered")
			}
			if error != nil {
				zend.Efree(error)
			}
			types.ZendStringReleaseEx(func_name, 0)
			return_value.SetFalse()
			return
		}
		alfi.SetCe(fcc.GetCallingScope())
		alfi.SetFuncPtr(fcc.GetFunctionHandler())
		obj_ptr = fcc.GetObject()
		if error != nil {
			zend.Efree(error)
		}
		if zcallable.IsType(types.IS_OBJECT) {
			types.ZVAL_COPY(alfi.GetClosure(), zcallable)
			lc_name = types.ZendStringAlloc(func_name.GetLen()+b.SizeOf("uint32_t"), 0)
			zend.ZendStrTolowerCopy(lc_name.GetVal(), func_name.GetVal(), func_name.GetLen())
			memcpy(lc_name.GetVal()+func_name.GetLen(), &(zend.Z_OBJ_HANDLE_P(zcallable)), b.SizeOf("uint32_t"))
			lc_name.GetVal()[lc_name.GetLen()] = '0'
		} else {
			alfi.GetClosure().SetUndef()

			/* Skip leading \ */

			if func_name.GetVal()[0] == '\\' {
				lc_name = types.ZendStringAlloc(func_name.GetLen()-1, 0)
				zend.ZendStrTolowerCopy(lc_name.GetVal(), func_name.GetVal()+1, func_name.GetLen()-1)
			} else {
				lc_name = zend.ZendStringTolower(func_name)
			}

			/* Skip leading \ */

		}
		types.ZendStringReleaseEx(func_name, 0)
		if SPL_G(autoload_functions) && SPL_G(autoload_functions).KeyExists(lc_name.GetStr()) {
			if !(alfi.GetClosure().IsUndef()) {
				alfi.GetClosure().DelRefcount()
			}
			goto skip
		}
		if obj_ptr != nil && !alfi.GetFuncPtr().HasFnFlags(zend.ZEND_ACC_STATIC) {

			/* add object id to the hash to ensure uniqueness, for more reference look at bug #40091 */

			lc_name = types.ZendStringExtend(lc_name, lc_name.GetLen()+b.SizeOf("uint32_t"), 0)
			memcpy(lc_name.GetVal()+lc_name.GetLen()-b.SizeOf("uint32_t"), obj_ptr.GetHandle(), b.SizeOf("uint32_t"))
			lc_name.GetVal()[lc_name.GetLen()] = '0'
			alfi.GetObj().SetObject(obj_ptr)
			alfi.GetObj().AddRefcount()
		} else {
			alfi.GetObj().SetUndef()
		}
		if !(SPL_G(autoload_functions)) {
			zend.ALLOC_HASHTABLE(SPL_G(autoload_functions))
			SPL_G(autoload_functions) = types.MakeArrayEx(1, AutoloadFuncInfoDtor, 0)
		}
		spl_func_ptr = SplAutoloadFn
		if zend.EG__().GetAutoloadFunc() == spl_func_ptr {
			var spl_alfi AutoloadFuncInfo
			spl_alfi.SetFuncPtr(spl_func_ptr)
			spl_alfi.GetObj().SetUndef()
			spl_alfi.GetClosure().SetUndef()
			spl_alfi.SetCe(nil)
			types.ZendHashAddMem(SPL_G(autoload_functions), SplAutoloadFn.GetFunctionName().GetStr(), &spl_alfi, b.SizeOf("autoload_func_info"))
			if prepend != 0 && SPL_G(autoload_functions).nNumOfElements > 1 {

				/* Move the newly created element to the head of the hashtable */

				HT_MOVE_TAIL_TO_HEAD(SPL_G(autoload_functions))

				/* Move the newly created element to the head of the hashtable */

			}
		}
		if alfi.GetFuncPtr() == zend.EG__().GetTrampoline() {
			var copy *zend.ZendFunction = zend.Emalloc(b.SizeOf("zend_op_array"))
			memcpy(copy, alfi.GetFuncPtr(), b.SizeOf("zend_op_array"))
			alfi.GetFuncPtr().SetFunctionName(nil)
			alfi.SetFuncPtr(copy)
		}
		if types.ZendHashAddMem(SPL_G(autoload_functions), lc_name.GetStr(), &alfi, b.SizeOf("autoload_func_info")) == nil {
			if obj_ptr != nil && !alfi.GetFuncPtr().HasFnFlags(zend.ZEND_ACC_STATIC) {
				alfi.GetObj().DelRefcount()
			}
			if !(alfi.GetClosure().IsUndef()) {
				alfi.GetClosure().DelRefcount()
			}
			if alfi.GetFuncPtr().HasFnFlags(zend.ZEND_ACC_CALL_VIA_TRAMPOLINE) {
				types.ZendStringReleaseEx(alfi.GetFuncPtr().GetFunctionName(), 0)
				zend.ZendFreeTrampoline(alfi.GetFuncPtr())
			}
		}
		if prepend != 0 && SPL_G(autoload_functions).nNumOfElements > 1 {

			/* Move the newly created element to the head of the hashtable */

			HT_MOVE_TAIL_TO_HEAD(SPL_G(autoload_functions))

			/* Move the newly created element to the head of the hashtable */

		}
	skip:
		types.ZendStringReleaseEx(lc_name, 0)
	}
	if SPL_G(autoload_functions) {
		zend.EG__().SetAutoloadFunc(SplAutoloadCallFn)
	} else {
		zend.EG__().SetAutoloadFunc(SplAutoloadFn)
	}
	return_value.SetTrue()
	return
}
func ZifSplAutoloadUnregister(executeData zpp.Ex, return_value zpp.Ret, autoloadFunction *types.Zval) {
	var func_name *types.String = nil
	var error *byte = nil
	var lc_name *types.String
	var zcallable *types.Zval
	var success int = types.FAILURE
	var spl_func_ptr *zend.ZendFunction
	var obj_ptr *types.ZendObject
	var fcc types.ZendFcallInfoCache
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &zcallable) == types.FAILURE {
		return
	}
	if zend.ZendIsCallableEx(zcallable, nil, zend.IS_CALLABLE_CHECK_SYNTAX_ONLY, &func_name, &fcc, &error) == 0 {
		faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Unable to unregister invalid function (%s)", error)
		if error != nil {
			zend.Efree(error)
		}
		if func_name != nil {
			types.ZendStringReleaseEx(func_name, 0)
		}
		return_value.SetFalse()
		return
	}
	obj_ptr = fcc.GetObject()
	if error != nil {
		zend.Efree(error)
	}
	if zcallable.IsType(types.IS_OBJECT) {
		lc_name = types.ZendStringAlloc(func_name.GetLen()+b.SizeOf("uint32_t"), 0)
		zend.ZendStrTolowerCopy(lc_name.GetVal(), func_name.GetVal(), func_name.GetLen())
		memcpy(lc_name.GetVal()+func_name.GetLen(), &(zend.Z_OBJ_HANDLE_P(zcallable)), b.SizeOf("uint32_t"))
		lc_name.GetVal()[lc_name.GetLen()] = '0'
	} else {

		/* Skip leading \ */

		if func_name.GetVal()[0] == '\\' {
			lc_name = types.ZendStringAlloc(func_name.GetLen()-1, 0)
			zend.ZendStrTolowerCopy(lc_name.GetVal(), func_name.GetVal()+1, func_name.GetLen()-1)
		} else {
			lc_name = zend.ZendStringTolower(func_name)
		}

		/* Skip leading \ */

	}
	types.ZendStringReleaseEx(func_name, 0)
	if SPL_G(autoload_functions) {
		if types.ZendStringEquals(lc_name, SplAutoloadCallFn.GetFunctionName()) != 0 {

			/* remove all */

			if !(SPL_G(autoload_running)) {
				SPL_G(autoload_functions).Destroy()
				zend.FREE_HASHTABLE(SPL_G(autoload_functions))
				SPL_G(autoload_functions) = nil
				zend.EG__().SetAutoloadFunc(nil)
			} else {
				SPL_G(autoload_functions).Clean()
			}
			success = types.SUCCESS
		} else {

			/* remove specific */

			success = types.ZendHashDel(SPL_G(autoload_functions), lc_name.GetStr())
			if success != types.SUCCESS && obj_ptr != nil {
				lc_name = types.ZendStringExtend(lc_name, lc_name.GetLen()+b.SizeOf("uint32_t"), 0)
				memcpy(lc_name.GetVal()+lc_name.GetLen()-b.SizeOf("uint32_t"), obj_ptr.GetHandle(), b.SizeOf("uint32_t"))
				lc_name.GetVal()[lc_name.GetLen()] = '0'
				success = types.ZendHashDel(SPL_G(autoload_functions), lc_name.GetStr())
			}
		}
	} else if types.ZendStringEquals(lc_name, SplAutoloadFn.GetFunctionName()) != 0 {

		/* register single spl_autoload() */

		spl_func_ptr = SplAutoloadFn
		if zend.EG__().GetAutoloadFunc() == spl_func_ptr {
			success = types.SUCCESS
			zend.EG__().SetAutoloadFunc(nil)
		}
	}
	types.ZendStringReleaseEx(lc_name, 0)
	types.ZVAL_BOOL(return_value, success == types.SUCCESS)
	return
}
func ZifSplAutoloadFunctions(executeData zpp.Ex, return_value zpp.Ret) {
	var fptr *zend.ZendFunction
	var alfi *AutoloadFuncInfo
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if zend.EG__().GetAutoloadFunc() == nil {
		if b.Assign(&fptr, types.ZendHashFindPtr(zend.EG__().GetFunctionTable(), types.ZSTR_MAGIC_AUTOLOAD.GetStr())) {
			var tmp types.Zval
			zend.ArrayInit(return_value)
			tmp.SetStringCopy(types.ZSTR_MAGIC_AUTOLOAD)
			return_value.GetArr().NextIndexInsertNew(&tmp)
			return
		}
		return_value.SetFalse()
		return
	}
	fptr = SplAutoloadCallFn
	if zend.EG__().GetAutoloadFunc() == fptr {
		var key *types.String
		zend.ArrayInit(return_value)
		var __ht *types.Array = SPL_G(autoload_functions)
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			alfi = _z.GetPtr()
			if !(alfi.GetClosure().IsUndef()) {
				alfi.GetClosure().AddRefcount()
				zend.AddNextIndexZval(return_value, alfi.GetClosure())
			} else if alfi.GetFuncPtr().GetScope() != nil {
				var tmp types.Zval
				zend.ArrayInit(&tmp)
				if !(alfi.GetObj().IsUndef()) {
					alfi.GetObj().AddRefcount()
					zend.AddNextIndexZval(&tmp, alfi.GetObj())
				} else {
					zend.AddNextIndexStr(&tmp, alfi.GetCe().GetName().Copy())
				}
				zend.AddNextIndexStr(&tmp, alfi.GetFuncPtr().GetFunctionName().Copy())
				zend.AddNextIndexZval(return_value, &tmp)
			} else {
				if strncmp(alfi.GetFuncPtr().GetFunctionName().GetVal(), "__lambda_func", b.SizeOf("\"__lambda_func\"")-1) {
					zend.AddNextIndexStr(return_value, alfi.GetFuncPtr().GetFunctionName().Copy())
				} else {
					zend.AddNextIndexStr(return_value, key.Copy())
				}
			}
		}
		return
	}
	zend.ArrayInit(return_value)
	zend.AddNextIndexStr(return_value, zend.EG__().GetAutoloadFunc().GetFunctionName().Copy())
}
func ZifSplObjectHash(executeData zpp.Ex, return_value zpp.Ret, obj *types.Zval) {
	var obj *types.Zval
	if zend.ZendParseParameters(executeData.NumArgs(), "o", &obj) == types.FAILURE {
		return
	}
	return_value.SetString(PhpSplObjectHash(obj))
	return
}
func ZifSplObjectId(executeData zpp.Ex, return_value zpp.Ret, obj *types.Zval) {
	var obj *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			obj = fp.ParseObject()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetLong(zend.ZendLong(zend.Z_OBJ_HANDLE_P(obj)))
	return
}
func PhpSplObjectHash(obj *types.Zval) *types.String {
	var hash_handle intPtr
	var hash_handlers intPtr
	if !(SPL_G(hash_mask_init)) {
		SPL_G(hash_mask_handle) = intptr_t(standard.PhpMtRand() >> 1)
		SPL_G(hash_mask_handlers) = intptr_t(standard.PhpMtRand() >> 1)
		SPL_G(hash_mask_init) = 1
	}
	hash_handle = SPL_G(hash_mask_handle) ^ intPtr(zend.Z_OBJ_HANDLE_P(obj))
	hash_handlers = SPL_G(hash_mask_handlers)
	return core.Strpprintf(32, "%016zx%016zx", hash_handle, hash_handlers)
}
func SplBuildClassListString(entry *types.Zval, list **byte) {
	var res *byte
	core.Spprintf(&res, 0, "%s, %s", *list, entry.GetStr().GetVal())
	zend.Efree(*list)
	*list = res
}
func ZmInfoSpl(zend_module *zend.ZendModuleEntry) {
	var list types.Zval
	var zv *types.Zval
	var strg *byte
	standard.PhpInfoPrintTableStart()
	standard.PhpInfoPrintTableHeader(2, "SPL support", "enabled")
	zend.ArrayInit(&list)
	SplAddClasses(spl_ce_AppendIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_ArrayIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_ArrayObject, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_BadFunctionCallException, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_BadMethodCallException, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_CachingIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_CallbackFilterIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_DirectoryIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_DomainException, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_EmptyIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_FilesystemIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_FilterIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_GlobIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_InfiniteIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_InvalidArgumentException, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_IteratorIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_LengthException, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_LimitIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_LogicException, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_MultipleIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_NoRewindIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_OuterIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_OutOfBoundsException, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_OutOfRangeException, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_OverflowException, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_ParentIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RangeException, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveArrayIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveCachingIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveCallbackFilterIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveDirectoryIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveFilterIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveIteratorIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveRegexIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveTreeIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RegexIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RuntimeException, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SeekableIterator, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplDoublyLinkedList, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplFileInfo, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplFileObject, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplFixedArray, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplHeap, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplMinHeap, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplMaxHeap, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplObjectStorage, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplObserver, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplPriorityQueue, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplQueue, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplStack, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplSubject, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplTempFileObject, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_UnderflowException, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_UnexpectedValueException, &list, 0, 1, zend.ZEND_ACC_INTERFACE)
	strg = zend.Estrdup("")
	var __ht *types.Array = list.GetArr()
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		zv = _z
		SplBuildClassListString(zv, &strg)
	}
	list.GetArr().DestroyEx()
	standard.PhpInfoPrintTableRow(2, "Interfaces", strg+2)
	zend.Efree(strg)
	zend.ArrayInit(&list)
	SplAddClasses(spl_ce_AppendIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_ArrayIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_ArrayObject, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_BadFunctionCallException, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_BadMethodCallException, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_CachingIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_CallbackFilterIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_DirectoryIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_DomainException, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_EmptyIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_FilesystemIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_FilterIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_GlobIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_InfiniteIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_InvalidArgumentException, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_IteratorIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_LengthException, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_LimitIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_LogicException, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_MultipleIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_NoRewindIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_OuterIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_OutOfBoundsException, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_OutOfRangeException, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_OverflowException, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_ParentIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RangeException, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveArrayIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveCachingIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveCallbackFilterIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveDirectoryIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveFilterIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveIteratorIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveRegexIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RecursiveTreeIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RegexIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_RuntimeException, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SeekableIterator, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplDoublyLinkedList, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplFileInfo, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplFileObject, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplFixedArray, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplHeap, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplMinHeap, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplMaxHeap, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplObjectStorage, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplObserver, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplPriorityQueue, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplQueue, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplStack, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplSubject, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_SplTempFileObject, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_UnderflowException, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	SplAddClasses(spl_ce_UnexpectedValueException, &list, 0, -1, zend.ZEND_ACC_INTERFACE)
	strg = zend.Estrdup("")
	var __ht__1 *types.Array = list.GetArr()
	for _, _p := range __ht__1.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		zv = _z
		SplBuildClassListString(zv, &strg)
	}
	list.GetArr().DestroyEx()
	standard.PhpInfoPrintTableRow(2, "Classes", strg+2)
	zend.Efree(strg)
	standard.PhpInfoPrintTableEnd()
}
func ZmStartupSpl(type_ int, module_number int) int {
	ZmStartupSplExceptions(type_, module_number)
	ZmStartupSplIterators(type_, module_number)
	ZmStartupSplArray(type_, module_number)
	ZmStartupSplDirectory(type_, module_number)
	ZmStartupSplDllist(type_, module_number)
	ZmStartupSplHeap(type_, module_number)
	ZmStartupSplFixedarray(type_, module_number)
	ZmStartupSplObserver(type_, module_number)
	SplAutoloadFn = types.ZendHashStrFindPtr(zend.CG__().GetFunctionTable(), "spl_autoload")
	SplAutoloadCallFn = types.ZendHashStrFindPtr(zend.CG__().GetFunctionTable(), "spl_autoload_call")
	b.Assert(SplAutoloadFn != nil && SplAutoloadCallFn != nil)
	return types.SUCCESS
}
func ZmActivateSpl(type_ int, module_number int) int {
	SPL_G(autoload_extensions) = nil
	SPL_G(autoload_functions) = nil
	SPL_G(hash_mask_init) = 0
	return types.SUCCESS
}
func ZmDeactivateSpl(type_ int, module_number int) int {
	if SPL_G(autoload_extensions) {
		types.ZendStringReleaseEx(SPL_G(autoload_extensions), 0)
		SPL_G(autoload_extensions) = nil
	}
	if SPL_G(autoload_functions) {
		SPL_G(autoload_functions).Destroy()
		zend.FREE_HASHTABLE(SPL_G(autoload_functions))
		SPL_G(autoload_functions) = nil
	}
	if SPL_G(hash_mask_init) {
		SPL_G(hash_mask_init) = 0
	}
	return types.SUCCESS
}
