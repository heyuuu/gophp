// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	"sik/zend"
)

func SPL_G(v __auto__) __auto__ { return SplGlobals.v }
func ZmGlobalsCtorSpl(spl_globals *ZendSplGlobals) {
	spl_globals.SetAutoloadExtensions(nil)
	spl_globals.SetAutoloadFunctions(nil)
	spl_globals.SetAutoloadRunning(0)
}
func SplFindCeByName(name *zend.ZendString, autoload zend.ZendBool) *zend.ZendClassEntry {
	var ce *zend.ZendClassEntry
	if autoload == 0 {
		var lc_name *zend.ZendString = zend.ZendStringTolower(name)
		ce = zend.ZendHashFindPtr(zend.ExecutorGlobals.GetClassTable(), lc_name)
		zend.ZendStringRelease(lc_name)
	} else {
		ce = zend.ZendLookupClass(name)
	}
	if ce == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Class %s does not exist%s", name.GetVal(), b.Cond(autoload != 0, " and could not be loaded", ""))
		return nil
	}
	return ce
}
func ZifClassParents(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var parent_class *zend.ZendClassEntry
	var ce *zend.ZendClassEntry
	var autoload zend.ZendBool = 1
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z|b", &obj, &autoload) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	if obj.GetType() != zend.IS_OBJECT && obj.GetType() != zend.IS_STRING {
		core.PhpErrorDocref(nil, zend.E_WARNING, "object or string expected")
		zend.RETVAL_FALSE
		return
	}
	if obj.IsType(zend.IS_STRING) {
		if nil == b.Assign(&ce, SplFindCeByName(obj.GetStr(), autoload)) {
			zend.RETVAL_FALSE
			return
		}
	} else {
		ce = zend.Z_OBJCE_P(obj)
	}
	zend.ArrayInit(return_value)
	parent_class = ce.parent
	for parent_class != nil {
		SplAddClassName(return_value, parent_class, 0, 0)
		parent_class = parent_class.parent
	}
}
func ZifClassImplements(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var autoload zend.ZendBool = 1
	var ce *zend.ZendClassEntry
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z|b", &obj, &autoload) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	if obj.GetType() != zend.IS_OBJECT && obj.GetType() != zend.IS_STRING {
		core.PhpErrorDocref(nil, zend.E_WARNING, "object or string expected")
		zend.RETVAL_FALSE
		return
	}
	if obj.IsType(zend.IS_STRING) {
		if nil == b.Assign(&ce, SplFindCeByName(obj.GetStr(), autoload)) {
			zend.RETVAL_FALSE
			return
		}
	} else {
		ce = zend.Z_OBJCE_P(obj)
	}
	zend.ArrayInit(return_value)
	SplAddInterfaces(return_value, ce, 1, zend.ZEND_ACC_INTERFACE)
}
func ZifClassUses(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var autoload zend.ZendBool = 1
	var ce *zend.ZendClassEntry
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z|b", &obj, &autoload) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	if obj.GetType() != zend.IS_OBJECT && obj.GetType() != zend.IS_STRING {
		core.PhpErrorDocref(nil, zend.E_WARNING, "object or string expected")
		zend.RETVAL_FALSE
		return
	}
	if obj.IsType(zend.IS_STRING) {
		if nil == b.Assign(&ce, SplFindCeByName(obj.GetStr(), autoload)) {
			zend.RETVAL_FALSE
			return
		}
	} else {
		ce = zend.Z_OBJCE_P(obj)
	}
	zend.ArrayInit(return_value)
	SplAddTraits(return_value, ce, 1, zend.ZEND_ACC_TRAIT)
}
func ZifSplClasses(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
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
func SplAutoload(class_name *zend.ZendString, lc_name *zend.ZendString, ext *byte, ext_len int) int {
	var class_file *byte
	var class_file_len int
	var dummy zend.Zval
	var file_handle zend.ZendFileHandle
	var new_op_array *zend.ZendOpArray
	var result zend.Zval
	var ret int
	class_file_len = int(core.Spprintf(&class_file, 0, "%s%.*s", lc_name.GetVal(), ext_len, ext))
	ret = core.PhpStreamOpenForZendEx(class_file, &file_handle, core.USE_PATH|core.STREAM_OPEN_FOR_INCLUDE)
	if ret == zend.SUCCESS {
		var opened_path *zend.ZendString
		if file_handle.GetOpenedPath() == nil {
			file_handle.SetOpenedPath(zend.ZendStringInit(class_file, class_file_len, 0))
		}
		opened_path = zend.ZendStringCopy(file_handle.GetOpenedPath())
		zend.ZVAL_NULL(&dummy)
		if zend.ZendHashAdd(&(zend.ExecutorGlobals.GetIncludedFiles()), opened_path, &dummy) != nil {
			new_op_array = zend.ZendCompileFile(&file_handle, zend.ZEND_REQUIRE)
			zend.ZendDestroyFileHandle(&file_handle)
		} else {
			new_op_array = nil
			zend.ZendFileHandleDtor(&file_handle)
		}
		zend.ZendStringReleaseEx(opened_path, 0)
		if new_op_array != nil {
			zend.ZVAL_UNDEF(&result)
			zend.ZendExecute(new_op_array, &result)
			zend.DestroyOpArray(new_op_array)
			zend.Efree(new_op_array)
			if zend.ExecutorGlobals.GetException() == nil {
				zend.ZvalPtrDtor(&result)
			}
			zend.Efree(class_file)
			return zend.ZendHashExists(zend.ExecutorGlobals.GetClassTable(), lc_name)
		}
	}
	zend.Efree(class_file)
	return 0
}
func ZifSplAutoload(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var pos_len int
	var pos1_len int
	var pos *byte
	var pos1 *byte
	var class_name *zend.ZendString
	var lc_name *zend.ZendString
	var file_exts *zend.ZendString = SPL_G(autoload_extensions)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "S|S", &class_name, &file_exts) == zend.FAILURE {
		zend.RETVAL_FALSE
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
	for pos != nil && (*pos) && zend.ExecutorGlobals.GetException() == nil {
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
	zend.ZendStringRelease(lc_name)
}
func ZifSplAutoloadExtensions(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var file_exts *zend.ZendString = nil
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|S", &file_exts) == zend.FAILURE {
		return
	}
	if file_exts != nil {
		if SPL_G(autoload_extensions) {
			zend.ZendStringReleaseEx(SPL_G(autoload_extensions), 0)
		}
		SPL_G(autoload_extensions) = zend.ZendStringCopy(file_exts)
	}
	if SPL_G(autoload_extensions) == nil {
		zend.RETVAL_STRINGL(SPL_DEFAULT_FILE_EXTENSIONS, b.SizeOf("SPL_DEFAULT_FILE_EXTENSIONS")-1)
		return
	} else {
		zend.ZendStringAddref(SPL_G(autoload_extensions))
		zend.RETVAL_STR(SPL_G(autoload_extensions))
		return
	}
}
func AutoloadFuncInfoDtor(element *zend.Zval) {
	var alfi *AutoloadFuncInfo = (*AutoloadFuncInfo)(element.GetPtr())
	if !(zend.Z_ISUNDEF(alfi.GetObj())) {
		zend.ZvalPtrDtor(alfi.GetObj())
	}
	if alfi.GetFuncPtr() != nil && alfi.GetFuncPtr().HasFnFlags(zend.ZEND_ACC_CALL_VIA_TRAMPOLINE) {
		zend.ZendStringReleaseEx(alfi.GetFuncPtr().GetFunctionName(), 0)
		zend.ZendFreeTrampoline(alfi.GetFuncPtr())
	}
	if !(zend.Z_ISUNDEF(alfi.GetClosure())) {
		zend.ZvalPtrDtor(alfi.GetClosure())
	}
	zend.Efree(alfi)
}
func ZifSplAutoloadCall(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var class_name *zend.Zval
	var retval zend.Zval
	var lc_name *zend.ZendString
	var func_name *zend.ZendString
	var alfi *AutoloadFuncInfo
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &class_name) == zend.FAILURE || class_name.GetType() != zend.IS_STRING {
		return
	}
	if SPL_G(autoload_functions) {
		var pos zend.HashPosition
		var num_idx zend.ZendUlong
		var func_ *zend.ZendFunction
		var fci zend.ZendFcallInfo
		var fcic zend.ZendFcallInfoCache
		var called_scope *zend.ZendClassEntry = zend.ZendGetCalledScope(execute_data)
		var l_autoload_running int = SPL_G(autoload_running)
		SPL_G(autoload_running) = 1
		lc_name = zend.ZendStringTolower(class_name.GetStr())
		fci.SetSize(b.SizeOf("fci"))
		fci.SetRetval(&retval)
		fci.SetParamCount(1)
		fci.SetParams(class_name)
		fci.SetNoSeparation(1)
		zend.ZVAL_UNDEF(fci.GetFunctionName())
		zend.ZendHashInternalPointerResetEx(SPL_G(autoload_functions), &pos)
		for zend.ZendHashGetCurrentKeyEx(SPL_G(autoload_functions), &func_name, &num_idx, &pos) == zend.HASH_KEY_IS_STRING {
			alfi = zend.ZendHashGetCurrentDataPtrEx(SPL_G(autoload_functions), &pos)
			func_ = alfi.GetFuncPtr()
			if func_.HasFnFlags(zend.ZEND_ACC_CALL_VIA_TRAMPOLINE) {
				func_ = zend.Emalloc(b.SizeOf("zend_op_array"))
				memcpy(func_, alfi.GetFuncPtr(), b.SizeOf("zend_op_array"))
				zend.ZendStringAddref(func_.GetOpArray().GetFunctionName())
			}
			zend.ZVAL_UNDEF(&retval)
			fcic.SetFunctionHandler(func_)
			if zend.Z_ISUNDEF(alfi.GetObj()) {
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
				fcic.SetCalledScope(zend.Z_OBJCE(alfi.GetObj()))
			}
			zend.ZendCallFunction(&fci, &fcic)
			zend.ZvalPtrDtor(&retval)
			if zend.ExecutorGlobals.GetException() != nil {
				break
			}
			if pos+1 == SPL_G(autoload_functions).nNumUsed || zend.ZendHashExists(zend.ExecutorGlobals.GetClassTable(), lc_name) != 0 {
				break
			}
			zend.ZendHashMoveForwardEx(SPL_G(autoload_functions), &pos)
		}
		zend.ZendStringReleaseEx(lc_name, 0)
		SPL_G(autoload_running) = l_autoload_running
	} else {

		/* do not use or overwrite &EG(autoload_func) here */

		var fcall_info zend.ZendFcallInfo
		var fcall_cache zend.ZendFcallInfoCache
		zend.ZVAL_UNDEF(&retval)
		fcall_info.SetSize(b.SizeOf("fcall_info"))
		zend.ZVAL_UNDEF(fcall_info.GetFunctionName())
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
func HT_MOVE_TAIL_TO_HEAD(ht *zend.HashTable) {
	var tmp zend.Bucket = ht.GetArData()[ht.GetNNumUsed()-1]
	memmove(ht.GetArData()+1, ht.GetArData(), b.SizeOf("Bucket")*(ht.GetNNumUsed()-1))
	ht.GetArData()[0] = tmp
	zend.ZendHashRehash(ht)
}
func ZifSplAutoloadRegister(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var func_name *zend.ZendString
	var error *byte = nil
	var lc_name *zend.ZendString
	var zcallable *zend.Zval = nil
	var do_throw zend.ZendBool = 1
	var prepend zend.ZendBool = 0
	var spl_func_ptr *zend.ZendFunction
	var alfi AutoloadFuncInfo
	var obj_ptr *zend.ZendObject
	var fcc zend.ZendFcallInfoCache
	if zend.ZendParseParametersEx(zend.ZEND_PARSE_PARAMS_QUIET, zend.ZEND_NUM_ARGS(), "|zbb", &zcallable, &do_throw, &prepend) == zend.FAILURE {
		return
	}
	if zend.ZEND_NUM_ARGS() != 0 {
		if zend.ZendIsCallableEx(zcallable, nil, zend.IS_CALLABLE_STRICT, &func_name, &fcc, &error) == 0 {
			alfi.SetCe(fcc.GetCallingScope())
			alfi.SetFuncPtr(fcc.GetFunctionHandler())
			obj_ptr = fcc.GetObject()
			if zcallable.IsType(zend.IS_ARRAY) {
				if obj_ptr == nil && alfi.GetFuncPtr() != nil && !alfi.GetFuncPtr().HasFnFlags(zend.ZEND_ACC_STATIC) {
					if do_throw != 0 {
						zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Passed array specifies a non static method but no object (%s)", error)
					}
					if error != nil {
						zend.Efree(error)
					}
					zend.ZendStringReleaseEx(func_name, 0)
					zend.RETVAL_FALSE
					return
				} else if do_throw != 0 {
					zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Passed array does not specify %s %smethod (%s)", b.Cond(alfi.GetFuncPtr() != nil, "a callable", "an existing"), b.Cond(obj_ptr == nil, "static ", ""), error)
				}
				if error != nil {
					zend.Efree(error)
				}
				zend.ZendStringReleaseEx(func_name, 0)
				zend.RETVAL_FALSE
				return
			} else if zcallable.IsType(zend.IS_STRING) {
				if do_throw != 0 {
					zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Function '%s' not %s (%s)", func_name.GetVal(), b.Cond(alfi.GetFuncPtr() != nil, "callable", "found"), error)
				}
				if error != nil {
					zend.Efree(error)
				}
				zend.ZendStringReleaseEx(func_name, 0)
				zend.RETVAL_FALSE
				return
			} else {
				if do_throw != 0 {
					zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Illegal value passed (%s)", error)
				}
				if error != nil {
					zend.Efree(error)
				}
				zend.ZendStringReleaseEx(func_name, 0)
				zend.RETVAL_FALSE
				return
			}
		} else if fcc.GetFunctionHandler().GetType() == zend.ZEND_INTERNAL_FUNCTION && fcc.GetFunctionHandler().GetInternalFunction().GetHandler() == ZifSplAutoloadCall {
			if do_throw != 0 {
				zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Function spl_autoload_call() cannot be registered")
			}
			if error != nil {
				zend.Efree(error)
			}
			zend.ZendStringReleaseEx(func_name, 0)
			zend.RETVAL_FALSE
			return
		}
		alfi.SetCe(fcc.GetCallingScope())
		alfi.SetFuncPtr(fcc.GetFunctionHandler())
		obj_ptr = fcc.GetObject()
		if error != nil {
			zend.Efree(error)
		}
		if zcallable.IsType(zend.IS_OBJECT) {
			zend.ZVAL_COPY(alfi.GetClosure(), zcallable)
			lc_name = zend.ZendStringAlloc(func_name.GetLen()+b.SizeOf("uint32_t"), 0)
			zend.ZendStrTolowerCopy(lc_name.GetVal(), func_name.GetVal(), func_name.GetLen())
			memcpy(lc_name.GetVal()+func_name.GetLen(), &(zend.Z_OBJ_HANDLE_P(zcallable)), b.SizeOf("uint32_t"))
			lc_name.GetVal()[lc_name.GetLen()] = '0'
		} else {
			zend.ZVAL_UNDEF(alfi.GetClosure())

			/* Skip leading \ */

			if func_name.GetVal()[0] == '\\' {
				lc_name = zend.ZendStringAlloc(func_name.GetLen()-1, 0)
				zend.ZendStrTolowerCopy(lc_name.GetVal(), func_name.GetVal()+1, func_name.GetLen()-1)
			} else {
				lc_name = zend.ZendStringTolower(func_name)
			}

			/* Skip leading \ */

		}
		zend.ZendStringReleaseEx(func_name, 0)
		if SPL_G(autoload_functions) && zend.ZendHashExists(SPL_G(autoload_functions), lc_name) != 0 {
			if !(zend.Z_ISUNDEF(alfi.GetClosure())) {
				zend.Z_DELREF_P(alfi.GetClosure())
			}
			goto skip
		}
		if obj_ptr != nil && !alfi.GetFuncPtr().HasFnFlags(zend.ZEND_ACC_STATIC) {

			/* add object id to the hash to ensure uniqueness, for more reference look at bug #40091 */

			lc_name = zend.ZendStringExtend(lc_name, lc_name.GetLen()+b.SizeOf("uint32_t"), 0)
			memcpy(lc_name.GetVal()+lc_name.GetLen()-b.SizeOf("uint32_t"), obj_ptr.GetHandle(), b.SizeOf("uint32_t"))
			lc_name.GetVal()[lc_name.GetLen()] = '0'
			zend.ZVAL_OBJ(alfi.GetObj(), obj_ptr)
			zend.Z_ADDREF(alfi.GetObj())
		} else {
			zend.ZVAL_UNDEF(alfi.GetObj())
		}
		if !(SPL_G(autoload_functions)) {
			zend.ALLOC_HASHTABLE(SPL_G(autoload_functions))
			zend.ZendHashInit(SPL_G(autoload_functions), 1, nil, AutoloadFuncInfoDtor, 0)
		}
		spl_func_ptr = SplAutoloadFn
		if zend.ExecutorGlobals.GetAutoloadFunc() == spl_func_ptr {
			var spl_alfi AutoloadFuncInfo
			spl_alfi.SetFuncPtr(spl_func_ptr)
			zend.ZVAL_UNDEF(spl_alfi.GetObj())
			zend.ZVAL_UNDEF(spl_alfi.GetClosure())
			spl_alfi.SetCe(nil)
			zend.ZendHashAddMem(SPL_G(autoload_functions), SplAutoloadFn.GetFunctionName(), &spl_alfi, b.SizeOf("autoload_func_info"))
			if prepend != 0 && SPL_G(autoload_functions).nNumOfElements > 1 {

				/* Move the newly created element to the head of the hashtable */

				HT_MOVE_TAIL_TO_HEAD(SPL_G(autoload_functions))

				/* Move the newly created element to the head of the hashtable */

			}
		}
		if alfi.GetFuncPtr() == &(zend.ExecutorGlobals.GetTrampoline()) {
			var copy *zend.ZendFunction = zend.Emalloc(b.SizeOf("zend_op_array"))
			memcpy(copy, alfi.GetFuncPtr(), b.SizeOf("zend_op_array"))
			alfi.GetFuncPtr().SetFunctionName(nil)
			alfi.SetFuncPtr(copy)
		}
		if zend.ZendHashAddMem(SPL_G(autoload_functions), lc_name, &alfi, b.SizeOf("autoload_func_info")) == nil {
			if obj_ptr != nil && !alfi.GetFuncPtr().HasFnFlags(zend.ZEND_ACC_STATIC) {
				zend.Z_DELREF(alfi.GetObj())
			}
			if !(zend.Z_ISUNDEF(alfi.GetClosure())) {
				zend.Z_DELREF(alfi.GetClosure())
			}
			if alfi.GetFuncPtr().HasFnFlags(zend.ZEND_ACC_CALL_VIA_TRAMPOLINE) {
				zend.ZendStringReleaseEx(alfi.GetFuncPtr().GetFunctionName(), 0)
				zend.ZendFreeTrampoline(alfi.GetFuncPtr())
			}
		}
		if prepend != 0 && SPL_G(autoload_functions).nNumOfElements > 1 {

			/* Move the newly created element to the head of the hashtable */

			HT_MOVE_TAIL_TO_HEAD(SPL_G(autoload_functions))

			/* Move the newly created element to the head of the hashtable */

		}
	skip:
		zend.ZendStringReleaseEx(lc_name, 0)
	}
	if SPL_G(autoload_functions) {
		zend.ExecutorGlobals.SetAutoloadFunc(SplAutoloadCallFn)
	} else {
		zend.ExecutorGlobals.SetAutoloadFunc(SplAutoloadFn)
	}
	zend.RETVAL_TRUE
	return
}
func ZifSplAutoloadUnregister(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var func_name *zend.ZendString = nil
	var error *byte = nil
	var lc_name *zend.ZendString
	var zcallable *zend.Zval
	var success int = zend.FAILURE
	var spl_func_ptr *zend.ZendFunction
	var obj_ptr *zend.ZendObject
	var fcc zend.ZendFcallInfoCache
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &zcallable) == zend.FAILURE {
		return
	}
	if zend.ZendIsCallableEx(zcallable, nil, zend.IS_CALLABLE_CHECK_SYNTAX_ONLY, &func_name, &fcc, &error) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Unable to unregister invalid function (%s)", error)
		if error != nil {
			zend.Efree(error)
		}
		if func_name != nil {
			zend.ZendStringReleaseEx(func_name, 0)
		}
		zend.RETVAL_FALSE
		return
	}
	obj_ptr = fcc.GetObject()
	if error != nil {
		zend.Efree(error)
	}
	if zcallable.IsType(zend.IS_OBJECT) {
		lc_name = zend.ZendStringAlloc(func_name.GetLen()+b.SizeOf("uint32_t"), 0)
		zend.ZendStrTolowerCopy(lc_name.GetVal(), func_name.GetVal(), func_name.GetLen())
		memcpy(lc_name.GetVal()+func_name.GetLen(), &(zend.Z_OBJ_HANDLE_P(zcallable)), b.SizeOf("uint32_t"))
		lc_name.GetVal()[lc_name.GetLen()] = '0'
	} else {

		/* Skip leading \ */

		if func_name.GetVal()[0] == '\\' {
			lc_name = zend.ZendStringAlloc(func_name.GetLen()-1, 0)
			zend.ZendStrTolowerCopy(lc_name.GetVal(), func_name.GetVal()+1, func_name.GetLen()-1)
		} else {
			lc_name = zend.ZendStringTolower(func_name)
		}

		/* Skip leading \ */

	}
	zend.ZendStringReleaseEx(func_name, 0)
	if SPL_G(autoload_functions) {
		if zend.ZendStringEquals(lc_name, SplAutoloadCallFn.GetFunctionName()) != 0 {

			/* remove all */

			if !(SPL_G(autoload_running)) {
				zend.ZendHashDestroy(SPL_G(autoload_functions))
				zend.FREE_HASHTABLE(SPL_G(autoload_functions))
				SPL_G(autoload_functions) = nil
				zend.ExecutorGlobals.SetAutoloadFunc(nil)
			} else {
				zend.ZendHashClean(SPL_G(autoload_functions))
			}
			success = zend.SUCCESS
		} else {

			/* remove specific */

			success = zend.ZendHashDel(SPL_G(autoload_functions), lc_name)
			if success != zend.SUCCESS && obj_ptr != nil {
				lc_name = zend.ZendStringExtend(lc_name, lc_name.GetLen()+b.SizeOf("uint32_t"), 0)
				memcpy(lc_name.GetVal()+lc_name.GetLen()-b.SizeOf("uint32_t"), obj_ptr.GetHandle(), b.SizeOf("uint32_t"))
				lc_name.GetVal()[lc_name.GetLen()] = '0'
				success = zend.ZendHashDel(SPL_G(autoload_functions), lc_name)
			}
		}
	} else if zend.ZendStringEquals(lc_name, SplAutoloadFn.GetFunctionName()) != 0 {

		/* register single spl_autoload() */

		spl_func_ptr = SplAutoloadFn
		if zend.ExecutorGlobals.GetAutoloadFunc() == spl_func_ptr {
			success = zend.SUCCESS
			zend.ExecutorGlobals.SetAutoloadFunc(nil)
		}
	}
	zend.ZendStringReleaseEx(lc_name, 0)
	zend.RETVAL_BOOL(success == zend.SUCCESS)
	return
}
func ZifSplAutoloadFunctions(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var fptr *zend.ZendFunction
	var alfi *AutoloadFuncInfo
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if zend.ExecutorGlobals.GetAutoloadFunc() == nil {
		if b.Assign(&fptr, zend.ZendHashFindPtr(zend.ExecutorGlobals.GetFunctionTable(), zend.ZSTR_KNOWN(zend.ZEND_STR_MAGIC_AUTOLOAD))) {
			var tmp zend.Zval
			zend.ArrayInit(return_value)
			zend.ZVAL_STR_COPY(&tmp, zend.ZSTR_KNOWN(zend.ZEND_STR_MAGIC_AUTOLOAD))
			zend.ZendHashNextIndexInsertNew(return_value.GetArr(), &tmp)
			return
		}
		zend.RETVAL_FALSE
		return
	}
	fptr = SplAutoloadCallFn
	if zend.ExecutorGlobals.GetAutoloadFunc() == fptr {
		var key *zend.ZendString
		zend.ArrayInit(return_value)
		for {
			var __ht *zend.HashTable = SPL_G(autoload_functions)
			var _p *zend.Bucket = __ht.GetArData()
			var _end *zend.Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *zend.Zval = _p.GetVal()

				if _z.IsType(zend.IS_UNDEF) {
					continue
				}
				key = _p.GetKey()
				alfi = _z.GetPtr()
				if !(zend.Z_ISUNDEF(alfi.GetClosure())) {
					zend.Z_ADDREF(alfi.GetClosure())
					zend.AddNextIndexZval(return_value, alfi.GetClosure())
				} else if alfi.GetFuncPtr().GetScope() != nil {
					var tmp zend.Zval
					zend.ArrayInit(&tmp)
					if !(zend.Z_ISUNDEF(alfi.GetObj())) {
						zend.Z_ADDREF(alfi.GetObj())
						zend.AddNextIndexZval(&tmp, alfi.GetObj())
					} else {
						zend.AddNextIndexStr(&tmp, zend.ZendStringCopy(alfi.GetCe().GetName()))
					}
					zend.AddNextIndexStr(&tmp, zend.ZendStringCopy(alfi.GetFuncPtr().GetFunctionName()))
					zend.AddNextIndexZval(return_value, &tmp)
				} else {
					if strncmp(alfi.GetFuncPtr().GetFunctionName().GetVal(), "__lambda_func", b.SizeOf("\"__lambda_func\"")-1) {
						zend.AddNextIndexStr(return_value, zend.ZendStringCopy(alfi.GetFuncPtr().GetFunctionName()))
					} else {
						zend.AddNextIndexStr(return_value, zend.ZendStringCopy(key))
					}
				}
			}
			break
		}
		return
	}
	zend.ArrayInit(return_value)
	zend.AddNextIndexStr(return_value, zend.ZendStringCopy(zend.ExecutorGlobals.GetAutoloadFunc().GetFunctionName()))
}
func ZifSplObjectHash(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "o", &obj) == zend.FAILURE {
		return
	}
	zend.RETVAL_NEW_STR(PhpSplObjectHash(obj))
	return
}
func ZifSplObjectId(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgObject(_arg, &obj, nil, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_OBJECT
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	zend.RETVAL_LONG(zend.ZendLong(zend.Z_OBJ_HANDLE_P(obj)))
	return
}
func PhpSplObjectHash(obj *zend.Zval) *zend.ZendString {
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
func SplBuildClassListString(entry *zend.Zval, list **byte) {
	var res *byte
	core.Spprintf(&res, 0, "%s, %s", *list, zend.Z_STRVAL_P(entry))
	zend.Efree(*list)
	*list = res
}
func ZmInfoSpl(ZEND_MODULE_INFO_FUNC_ARGS) {
	var list zend.Zval
	var zv *zend.Zval
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
	for {
		var __ht *zend.HashTable = list.GetArr()
		var _p *zend.Bucket = __ht.GetArData()
		var _end *zend.Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *zend.Zval = _p.GetVal()

			if _z.IsType(zend.IS_UNDEF) {
				continue
			}
			zv = _z
			SplBuildClassListString(zv, &strg)
		}
		break
	}
	zend.ZendArrayDestroy(list.GetArr())
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
	for {
		var __ht *zend.HashTable = list.GetArr()
		var _p *zend.Bucket = __ht.GetArData()
		var _end *zend.Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *zend.Zval = _p.GetVal()

			if _z.IsType(zend.IS_UNDEF) {
				continue
			}
			zv = _z
			SplBuildClassListString(zv, &strg)
		}
		break
	}
	zend.ZendArrayDestroy(list.GetArr())
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
	SplAutoloadFn = zend.ZendHashStrFindPtr(zend.CompilerGlobals.GetFunctionTable(), "spl_autoload", b.SizeOf("\"spl_autoload\"")-1)
	SplAutoloadCallFn = zend.ZendHashStrFindPtr(zend.CompilerGlobals.GetFunctionTable(), "spl_autoload_call", b.SizeOf("\"spl_autoload_call\"")-1)
	zend.ZEND_ASSERT(SplAutoloadFn != nil && SplAutoloadCallFn != nil)
	return zend.SUCCESS
}
func ZmActivateSpl(type_ int, module_number int) int {
	SPL_G(autoload_extensions) = nil
	SPL_G(autoload_functions) = nil
	SPL_G(hash_mask_init) = 0
	return zend.SUCCESS
}
func ZmDeactivateSpl(type_ int, module_number int) int {
	if SPL_G(autoload_extensions) {
		zend.ZendStringReleaseEx(SPL_G(autoload_extensions), 0)
		SPL_G(autoload_extensions) = nil
	}
	if SPL_G(autoload_functions) {
		zend.ZendHashDestroy(SPL_G(autoload_functions))
		zend.FREE_HASHTABLE(SPL_G(autoload_functions))
		SPL_G(autoload_functions) = nil
	}
	if SPL_G(hash_mask_init) {
		SPL_G(hash_mask_init) = 0
	}
	return zend.SUCCESS
}
