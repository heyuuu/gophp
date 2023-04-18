package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
)

func SPL_G(v __auto__) __auto__ { return SplGlobals.v }
func SPL_G__() *ZendSplGlobals  { return &SplGlobals }
func ZmGlobalsCtorSpl(spl_globals *ZendSplGlobals) {
	spl_globals.SetAutoloadExtensions(nil)
	spl_globals.SetAutoloadFunctions(nil)
	spl_globals.SetAutoloadRunning(0)
}
func SplFindCeByName(name *types2.String, autoload types2.ZendBool) *types2.ClassEntry {
	var ce *types2.ClassEntry
	if autoload == 0 {
		ce = zend.EG__().ClassTable().Get(name.GetStr())
	} else {
		ce = zend.ZendLookupClass(name)
	}
	if ce == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Class %s does not exist%s", name.GetVal(), b.Cond(autoload != 0, " and could not be loaded", ""))
		return nil
	}
	return ce
}
func ZifClassParents(executeData zpp.Ex, return_value zpp.Ret, instance *types2.Zval, _ zpp.Opt, autoload *types2.Zval) {
	var obj *types2.Zval
	var parent_class *types2.ClassEntry
	var ce *types2.ClassEntry
	var autoload types2.ZendBool = 1
	if zend.ZendParseParameters(executeData.NumArgs(), "z|b", &obj, &autoload) == types2.FAILURE {
		return_value.SetFalse()
		return
	}
	if obj.GetType() != types2.IS_OBJECT && obj.GetType() != types2.IS_STRING {
		core.PhpErrorDocref(nil, faults.E_WARNING, "object or string expected")
		return_value.SetFalse()
		return
	}
	if obj.IsType(types2.IS_STRING) {
		if nil == b.Assign(&ce, SplFindCeByName(obj.String(), autoload)) {
			return_value.SetFalse()
			return
		}
	} else {
		ce = types2.Z_OBJCE_P(obj)
	}
	zend.ArrayInit(return_value)
	parent_class = ce.GetParent()
	for parent_class != nil {
		SplAddClassName(return_value, parent_class, 0, 0)
		parent_class = parent_class.GetParent()
	}
}
func ZifClassImplements(executeData zpp.Ex, return_value zpp.Ret, what *types2.Zval, _ zpp.Opt, autoload *types2.Zval) {
	var obj *types2.Zval
	var autoload types2.ZendBool = 1
	var ce *types2.ClassEntry
	if zend.ZendParseParameters(executeData.NumArgs(), "z|b", &obj, &autoload) == types2.FAILURE {
		return_value.SetFalse()
		return
	}
	if obj.GetType() != types2.IS_OBJECT && obj.GetType() != types2.IS_STRING {
		core.PhpErrorDocref(nil, faults.E_WARNING, "object or string expected")
		return_value.SetFalse()
		return
	}
	if obj.IsType(types2.IS_STRING) {
		if nil == b.Assign(&ce, SplFindCeByName(obj.String(), autoload)) {
			return_value.SetFalse()
			return
		}
	} else {
		ce = types2.Z_OBJCE_P(obj)
	}
	zend.ArrayInit(return_value)
	SplAddInterfaces(return_value, ce, 1, zend.AccInterface)
}
func ZifClassUses(executeData zpp.Ex, return_value zpp.Ret, what *types2.Zval, _ zpp.Opt, autoload *types2.Zval) {
	var obj *types2.Zval
	var autoload types2.ZendBool = 1
	var ce *types2.ClassEntry
	if zend.ZendParseParameters(executeData.NumArgs(), "z|b", &obj, &autoload) == types2.FAILURE {
		return_value.SetFalse()
		return
	}
	if obj.GetType() != types2.IS_OBJECT && obj.GetType() != types2.IS_STRING {
		core.PhpErrorDocref(nil, faults.E_WARNING, "object or string expected")
		return_value.SetFalse()
		return
	}
	if obj.IsType(types2.IS_STRING) {
		if nil == b.Assign(&ce, SplFindCeByName(obj.String(), autoload)) {
			return_value.SetFalse()
			return
		}
	} else {
		ce = types2.Z_OBJCE_P(obj)
	}
	zend.ArrayInit(return_value)
	SplAddTraits(return_value, ce, 1, zend.AccTrait)
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
func SplAutoload(className string, lcName string, ext *byte, ext_len int) int {
	var dummy types2.Zval
	var opArray *types2.ZendOpArray
	var result types2.Zval

	classFile := lcName + b.CastStr(ext, ext_len)
	fh := core.PhpStreamOpenForZendEx(classFile, core.USE_PATH|core.STREAM_OPEN_FOR_INCLUDE)
	if fh != nil {
		var opened_path *types2.String
		if fh.GetOpenedPath() == "" {
			fh.SetOpenedPath(classFile)
		}
		opened_path = fh.GetOpenedPath().Copy()
		dummy.SetNull()
		if zend.EG__().GetIncludedFiles().KeyAdd(opened_path.GetStr(), &dummy) != nil {
			opArray = zend.CompileFile(fh, zend.ZEND_REQUIRE)
			zend.ZendDestroyFileHandle(fh)
		} else {
			opArray = nil
			fh.Destroy()
		}
		if opArray != nil {
			result.SetUndef()
			zend.ZendExecute(opArray, &result)
			zend.DestroyOpArray(opArray)
			zend.Efree(opArray)
			return types2.IntBool(zend.EG__().ClassTable().Exists(lcName.GetStr()))
		}
	}
	return 0
}
func ZifSplAutoload(executeData zpp.Ex, return_value zpp.Ret, className *types2.Zval, _ zpp.Opt, fileExtensions *types2.Zval) {
	var pos_len int
	var pos1_len int
	var pos *byte
	var pos1 *byte
	var class_name *types2.String
	var lc_name *types2.String
	var file_exts *types2.String = SPL_G__().autoload_extensions
	if zend.ZendParseParameters(executeData.NumArgs(), "S|S", &class_name, &file_exts) == types2.FAILURE {
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
	// types.ZendStringRelease(lc_name)
}
func ZifSplAutoloadExtensions(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, fileExtensions *types2.Zval) {
	var file_exts *types2.String = nil
	if zend.ZendParseParameters(executeData.NumArgs(), "|S", &file_exts) == types2.FAILURE {
		return
	}
	if file_exts != nil {
		if SPL_G__().autoload_extensions {
			// types.ZendStringReleaseEx(SPL_G(autoload_extensions), 0)
		}
		SPL_G__().autoload_extensions = file_exts.Copy()
	}
	if SPL_G__().autoload_extensions == nil {
		return_value.SetStringVal(SPL_DEFAULT_FILE_EXTENSIONS)
		return
	} else {
		//SPL_G(autoload_extensions).AddRefcount()
		return_value.SetString(SPL_G__().autoload_extensions)
		return
	}
}
func AutoloadFuncInfoDtor(element *types2.Zval) {
	var alfi *AutoloadFuncInfo = (*AutoloadFuncInfo)(element.Ptr())
	if !(alfi.GetObj().IsUndef()) {
		// zend.ZvalPtrDtor(alfi.GetObj())
	}
	if alfi.GetFuncPtr() != nil && alfi.GetFuncPtr().HasFnFlags(zend.AccCallViaTrampoline) {
		// types.ZendStringReleaseEx(alfi.GetFuncPtr().GetFunctionName(), 0)
		zend.ZendFreeTrampoline(alfi.GetFuncPtr())
	}
	if !(alfi.GetClosure().IsUndef()) {
		// zend.ZvalPtrDtor(alfi.GetClosure())
	}
	zend.Efree(alfi)
}
func ZifSplAutoloadCall(executeData zpp.Ex, return_value zpp.Ret, className *types2.Zval) {
	var class_name *types2.Zval
	var retval types2.Zval
	var lc_name *types2.String
	var func_name *types2.String
	var alfi *AutoloadFuncInfo
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &class_name) == types2.FAILURE || class_name.GetType() != types2.IS_STRING {
		return
	}
	if SPL_G__().autoload_functions {
		var pos types2.ArrayPosition
		var num_idx zend.ZendUlong
		var func_ types2.IFunction
		var fci types2.ZendFcallInfo
		var fcic types2.ZendFcallInfoCache
		var called_scope *types2.ClassEntry = zend.ZendGetCalledScope(executeData)
		var l_autoload_running int = SPL_G__().autoloadRunning
		SPL_G__().autoloadRunning = 1
		lc_name = zend.ZendStringTolower(class_name.String())
		fci.SetSize(b.SizeOf("fci"))
		fci.SetRetval(&retval)
		fci.SetParamCount(1)
		fci.SetParams(class_name)
		fci.SetNoSeparation(1)
		fci.GetFunctionName().SetUndef()
		types2.ZendHashInternalPointerResetEx(SPL_G__().autoload_functions, &pos)
		for types2.ZendHashGetCurrentKeyEx(SPL_G__().autoload_functions, &func_name, &num_idx, &pos) == types2.HASH_KEY_IS_STRING {
			alfi = types2.ZendHashGetCurrentDataPtrEx(SPL_G__().autoload_functions, &pos)
			func_ = alfi.GetFuncPtr()
			if func_.HasFnFlags(zend.AccCallViaTrampoline) {
				func_ = zend.Emalloc(b.SizeOf("zend_op_array"))
				memcpy(func_, alfi.GetFuncPtr(), b.SizeOf("zend_op_array"))
				//func_.GetOpArray().GetFunctionName().AddRefcount()
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
				fci.SetObject(alfi.GetObj().Object())
				fcic.SetObject(alfi.GetObj().Object())
				fcic.SetCalledScope(types2.Z_OBJCE(alfi.GetObj()))
			}
			zend.ZendCallFunction(&fci, &fcic)
			// zend.ZvalPtrDtor(&retval)
			if zend.EG__().GetException() != nil {
				break
			}
			if pos+1 == SPL_G__().autoload_functions.nNumUsed || zend.EG__().ClassTable().Exists(lc_name.GetStr()) {
				break
			}
			types2.ZendHashMoveForwardEx(SPL_G__().autoload_functions, &pos)
		}
		// types.ZendStringReleaseEx(lc_name, 0)
		SPL_G__().autoloadRunning = l_autoload_running
	} else {

		/* do not use or overwrite &EG(autoload_func) here */

		var fcall_info types2.ZendFcallInfo
		var fcall_cache types2.ZendFcallInfoCache
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
		// zend.ZvalPtrDtor(&retval)
	}
}
func HT_MOVE_TAIL_TO_HEAD(ht *types2.Array) {
	var tmp types2.Bucket = ht.GetArData()[ht.GetNNumUsed()-1]
	memmove(ht.GetArData()+1, ht.GetArData(), b.SizeOf("Bucket")*(ht.GetNNumUsed()-1))
	ht.GetArData()[0] = tmp
	ht.Rehash()
}
func ZifSplAutoloadRegister(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, autoloadFunction *types2.Zval, throw *types2.Zval, prepend *types2.Zval) {
	var func_name *types2.String
	var error *byte = nil
	var lc_name *types2.String
	var zcallable *types2.Zval = nil
	var do_throw types2.ZendBool = 1
	var prepend types2.ZendBool = 0
	var spl_func_ptr types2.IFunction
	var alfi AutoloadFuncInfo
	var obj_ptr *types2.ZendObject
	var fcc types2.ZendFcallInfoCache
	if zend.ZendParseParametersEx(zpp.FlagQuiet, executeData.NumArgs(), "|zbb", &zcallable, &do_throw, &prepend) == types2.FAILURE {
		return
	}
	if executeData.NumArgs() != 0 {
		if zend.ZendIsCallableEx(zcallable, nil, zend.IS_CALLABLE_STRICT, &func_name, &fcc, &error) == 0 {
			alfi.SetCe(fcc.GetCallingScope())
			alfi.SetFuncPtr(fcc.GetFunctionHandler())
			obj_ptr = fcc.GetObject()
			if zcallable.IsType(types2.IS_ARRAY) {
				if obj_ptr == nil && alfi.GetFuncPtr() != nil && !alfi.GetFuncPtr().HasFnFlags(zend.AccStatic) {
					if do_throw != 0 {
						faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Passed array specifies a non static method but no object (%s)", error)
					}
					if error != nil {
						zend.Efree(error)
					}
					// types.ZendStringReleaseEx(func_name, 0)
					return_value.SetFalse()
					return
				} else if do_throw != 0 {
					faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Passed array does not specify %s %smethod (%s)", b.Cond(alfi.GetFuncPtr() != nil, "a callable", "an existing"), b.Cond(obj_ptr == nil, "static ", ""), error)
				}
				if error != nil {
					zend.Efree(error)
				}
				// types.ZendStringReleaseEx(func_name, 0)
				return_value.SetFalse()
				return
			} else if zcallable.IsType(types2.IS_STRING) {
				if do_throw != 0 {
					faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Function '%s' not %s (%s)", func_name.GetVal(), b.Cond(alfi.GetFuncPtr() != nil, "callable", "found"), error)
				}
				if error != nil {
					zend.Efree(error)
				}
				// types.ZendStringReleaseEx(func_name, 0)
				return_value.SetFalse()
				return
			} else {
				if do_throw != 0 {
					faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Illegal value passed (%s)", error)
				}
				if error != nil {
					zend.Efree(error)
				}
				// types.ZendStringReleaseEx(func_name, 0)
				return_value.SetFalse()
				return
			}
		} else if interFunc, ok := fcc.GetFunctionHandler().(*types2.InternalFunction); ok && interFunc.GetHandler() == ZifSplAutoloadCall {
			if do_throw != 0 {
				faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Function spl_autoload_call() cannot be registered")
			}
			if error != nil {
				zend.Efree(error)
			}
			// types.ZendStringReleaseEx(func_name, 0)
			return_value.SetFalse()
			return
		}
		alfi.SetCe(fcc.GetCallingScope())
		alfi.SetFuncPtr(fcc.GetFunctionHandler())
		obj_ptr = fcc.GetObject()
		if error != nil {
			zend.Efree(error)
		}
		if zcallable.IsType(types2.IS_OBJECT) {
			types2.ZVAL_COPY(alfi.GetClosure(), zcallable)
			lc_name = types2.ZendStringAlloc(func_name.GetLen()+b.SizeOf("uint32_t"), 0)
			zend.ZendStrTolowerCopy(lc_name.GetVal(), func_name.GetVal(), func_name.GetLen())
			memcpy(lc_name.GetVal()+func_name.GetLen(), &(zend.Z_OBJ_HANDLE_P(zcallable)), b.SizeOf("uint32_t"))
			lc_name.GetStr()[lc_name.GetLen()] = '0'
		} else {
			alfi.GetClosure().SetUndef()

			/* Skip leading \ */

			if func_name.GetStr()[0] == '\\' {
				lc_name = types2.ZendStringAlloc(func_name.GetLen()-1, 0)
				zend.ZendStrTolowerCopy(lc_name.GetVal(), func_name.GetVal()+1, func_name.GetLen()-1)
			} else {
				lc_name = zend.ZendStringTolower(func_name)
			}

			/* Skip leading \ */

		}
		// types.ZendStringReleaseEx(func_name, 0)
		if SPL_G__().autoload_functions && SPL_G__().autoload_functions.KeyExists(lc_name.GetStr()) {
			if !(alfi.GetClosure().IsUndef()) {
				//alfi.GetClosure().DelRefcount()
			}
			goto skip
		}
		if obj_ptr != nil && !alfi.GetFuncPtr().HasFnFlags(zend.AccStatic) {

			/* add object id to the hash to ensure uniqueness, for more reference look at bug #40091 */

			lc_name = types2.ZendStringExtend(lc_name, lc_name.GetLen()+b.SizeOf("uint32_t"))
			memcpy(lc_name.GetVal()+lc_name.GetLen()-b.SizeOf("uint32_t"), obj_ptr.GetHandle(), b.SizeOf("uint32_t"))
			lc_name.GetStr()[lc_name.GetLen()] = '0'
			alfi.GetObj().SetObject(obj_ptr)
			//alfi.GetObj().AddRefcount()
		} else {
			alfi.GetObj().SetUndef()
		}
		if !(SPL_G__().autoload_functions) {
			zend.ALLOC_HASHTABLE(SPL_G__().autoload_functions)
			SPL_G__().autoload_functions.Init(1, AutoloadFuncInfoDtor)
		}
		spl_func_ptr = SplAutoloadFn
		if zend.EG__().GetAutoloadFunc() == spl_func_ptr {
			var spl_alfi AutoloadFuncInfo
			spl_alfi.SetFuncPtr(spl_func_ptr)
			spl_alfi.GetObj().SetUndef()
			spl_alfi.GetClosure().SetUndef()
			spl_alfi.SetCe(nil)
			types2.ZendHashAddMem(SPL_G__().autoload_functions, SplAutoloadFn.GetFunctionName().GetStr(), &spl_alfi, b.SizeOf("autoload_func_info"))
			if prepend != 0 && SPL_G__().autoload_functions.nNumOfElements > 1 {

				/* Move the newly created element to the head of the hashtable */

				HT_MOVE_TAIL_TO_HEAD(SPL_G__().autoload_functions)

				/* Move the newly created element to the head of the hashtable */

			}
		}
		if alfi.GetFuncPtr() == zend.EG__().GetTrampoline() {
			var copy types2.IFunction = zend.Emalloc(b.SizeOf("zend_op_array"))
			memcpy(copy, alfi.GetFuncPtr(), b.SizeOf("zend_op_array"))
			alfi.GetFuncPtr().SetFunctionName(nil)
			alfi.SetFuncPtr(copy)
		}
		if types2.ZendHashAddMem(SPL_G__().autoload_functions, lc_name.GetStr(), &alfi, b.SizeOf("autoload_func_info")) == nil {
			//if obj_ptr != nil && !alfi.GetFuncPtr().HasFnFlags(zend.AccStatic) {
			//	alfi.GetObj().DelRefcount()
			//}
			//if !(alfi.GetClosure().IsUndef()) {
			//	alfi.GetClosure().DelRefcount()
			//}
			if alfi.GetFuncPtr().HasFnFlags(zend.AccCallViaTrampoline) {
				// types.ZendStringReleaseEx(alfi.GetFuncPtr().GetFunctionName(), 0)
				zend.ZendFreeTrampoline(alfi.GetFuncPtr())
			}
		}
		if prepend != 0 && SPL_G__().autoload_functions.nNumOfElements > 1 {

			/* Move the newly created element to the head of the hashtable */

			HT_MOVE_TAIL_TO_HEAD(SPL_G__().autoload_functions)

			/* Move the newly created element to the head of the hashtable */

		}
	skip:
		// types.ZendStringReleaseEx(lc_name, 0)
	}
	if SPL_G__().autoload_functions {
		zend.EG__().SetAutoloadFunc(SplAutoloadCallFn)
	} else {
		zend.EG__().SetAutoloadFunc(SplAutoloadFn)
	}
	return_value.SetTrue()
	return
}
func ZifSplAutoloadUnregister(executeData zpp.Ex, return_value zpp.Ret, autoloadFunction *types2.Zval) {
	var func_name *types2.String = nil
	var error *byte = nil
	var lc_name *types2.String
	var zcallable *types2.Zval
	var success int = types2.FAILURE
	var spl_func_ptr types2.IFunction
	var obj_ptr *types2.ZendObject
	var fcc types2.ZendFcallInfoCache
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &zcallable) == types2.FAILURE {
		return
	}
	if zend.ZendIsCallableEx(zcallable, nil, zend.IS_CALLABLE_CHECK_SYNTAX_ONLY, &func_name, &fcc, &error) == 0 {
		faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Unable to unregister invalid function (%s)", error)
		if error != nil {
			zend.Efree(error)
		}
		if func_name != nil {
			// types.ZendStringReleaseEx(func_name, 0)
		}
		return_value.SetFalse()
		return
	}
	obj_ptr = fcc.GetObject()
	if error != nil {
		zend.Efree(error)
	}
	if zcallable.IsType(types2.IS_OBJECT) {
		lc_name = types2.ZendStringAlloc(func_name.GetLen()+b.SizeOf("uint32_t"), 0)
		zend.ZendStrTolowerCopy(lc_name.GetVal(), func_name.GetVal(), func_name.GetLen())
		memcpy(lc_name.GetVal()+func_name.GetLen(), &(zend.Z_OBJ_HANDLE_P(zcallable)), b.SizeOf("uint32_t"))
		lc_name.GetStr()[lc_name.GetLen()] = '0'
	} else {

		/* Skip leading \ */

		if func_name.GetStr()[0] == '\\' {
			lc_name = types2.ZendStringAlloc(func_name.GetLen()-1, 0)
			zend.ZendStrTolowerCopy(lc_name.GetVal(), func_name.GetVal()+1, func_name.GetLen()-1)
		} else {
			lc_name = zend.ZendStringTolower(func_name)
		}

		/* Skip leading \ */

	}
	// types.ZendStringReleaseEx(func_name, 0)
	if SPL_G__().autoload_functions {
		if lc_name.GetStr() == SplAutoloadCallFn.GetFunctionName().GetStr() {

			/* remove all */

			if !(SPL_G__().autoloadRunning) {
				SPL_G__().autoload_functions.Destroy()
				zend.FREE_HASHTABLE(SPL_G__().autoload_functions)
				SPL_G__().autoload_functions = nil
				zend.EG__().SetAutoloadFunc(nil)
			} else {
				SPL_G__().autoload_functions.Clean()
			}
			success = types2.SUCCESS
		} else {

			/* remove specific */

			success = types2.ZendHashDel(SPL_G__().autoload_functions, lc_name.GetStr())
			if success != types2.SUCCESS && obj_ptr != nil {
				lc_name = types2.ZendStringExtend(lc_name, lc_name.GetLen()+b.SizeOf("uint32_t"))
				memcpy(lc_name.GetVal()+lc_name.GetLen()-b.SizeOf("uint32_t"), obj_ptr.GetHandle(), b.SizeOf("uint32_t"))
				lc_name.GetStr()[lc_name.GetLen()] = '0'
				success = types2.ZendHashDel(SPL_G__().autoload_functions, lc_name.GetStr())
			}
		}
	} else if lc_name.GetStr() == SplAutoloadFn.GetFunctionName().GetStr() {

		/* register single spl_autoload() */

		spl_func_ptr = SplAutoloadFn
		if zend.EG__().GetAutoloadFunc() == spl_func_ptr {
			success = types2.SUCCESS
			zend.EG__().SetAutoloadFunc(nil)
		}
	}
	// types.ZendStringReleaseEx(lc_name, 0)
	return_value.SetBool(success == types2.SUCCESS)
	return
}
func ZifSplAutoloadFunctions(executeData zpp.Ex, return_value zpp.Ret) {
	var fptr types2.IFunction
	var alfi *AutoloadFuncInfo
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if zend.EG__().GetAutoloadFunc() == nil {
		if fptr := zend.EG__().FunctionTable().Get(types2.STR_MAGIC_AUTOLOAD); fptr != nil {
			var tmp types2.Zval
			zend.ArrayInit(return_value)
			tmp.SetStringVal(types2.STR_MAGIC_AUTOLOAD)
			return_value.Array().NextIndexInsertNew(&tmp)
			return
		}
		return_value.SetFalse()
		return
	}
	fptr = SplAutoloadCallFn
	if zend.EG__().GetAutoloadFunc() == fptr {
		var key *types2.String
		zend.ArrayInit(return_value)
		var __ht *types2.Array = SPL_G__().autoload_functions
		for _, _p := range __ht.ForeachData() {
			var _z *types2.Zval = _p.GetVal()

			key = _p.GetKey()
			alfi = _z.Ptr()
			if !(alfi.GetClosure().IsUndef()) {
				//alfi.GetClosure().AddRefcount()
				zend.AddNextIndexZval(return_value, alfi.GetClosure())
			} else if alfi.GetFuncPtr().GetScope() != nil {
				var tmp types2.Zval
				zend.ArrayInit(&tmp)
				if !(alfi.GetObj().IsUndef()) {
					//alfi.GetObj().AddRefcount()
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
func ZifSplObjectHash(executeData zpp.Ex, return_value zpp.Ret, obj *types2.Zval) {
	var obj *types2.Zval
	if zend.ZendParseParameters(executeData.NumArgs(), "o", &obj) == types2.FAILURE {
		return
	}
	return_value.SetStringVal(PhpSplObjectHash(obj))
	return
}
func ZifSplObjectId(executeData zpp.Ex, return_value zpp.Ret, obj *types2.Zval) {
	var obj *types2.Zval
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
func PhpSplObjectHash(obj *types2.Zval) string {
	return SPL_G__().SplObjectHash(obj.Object().GetHandle())
}
func buildClassListString(arr *types2.Array) string {
	names := make([]string, 0, arr.Len())
	arr.Foreach(func(key types2.ArrayKey, value *types2.Zval) {
		names = append(names, value.StringVal())
	})
	return strings.Join(names, ", ")
}

func SplBuildClassListString(entry *types2.Zval, list **byte) {
	var res *byte
	core.Spprintf(&res, 0, "%s, %s", *list, entry.String().GetVal())
	zend.Efree(*list)
	*list = res
}
func ZmInfoSpl(zend_module *zend.ModuleEntry) {
	var list types2.Zval
	var zv *types2.Zval
	var strg *byte
	standard.PhpInfoPrintTableStart()
	standard.PhpInfoPrintTableHeader(2, "SPL support", "enabled")
	zend.ArrayInit(&list)
	SplAddClasses(spl_ce_AppendIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_ArrayIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_ArrayObject, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_BadFunctionCallException, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_BadMethodCallException, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_CachingIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_CallbackFilterIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_DirectoryIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_DomainException, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_EmptyIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_FilesystemIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_FilterIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_GlobIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_InfiniteIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_InvalidArgumentException, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_IteratorIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_LengthException, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_LimitIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_LogicException, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_MultipleIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_NoRewindIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_OuterIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_OutOfBoundsException, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_OutOfRangeException, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_OverflowException, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_ParentIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_RangeException, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveArrayIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveCachingIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveCallbackFilterIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveDirectoryIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveFilterIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveIteratorIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveRegexIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveTreeIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_RegexIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_RuntimeException, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SeekableIterator, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplDoublyLinkedList, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplFileInfo, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplFileObject, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplFixedArray, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplHeap, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplMinHeap, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplMaxHeap, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplObjectStorage, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplObserver, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplPriorityQueue, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplQueue, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplStack, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplSubject, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_SplTempFileObject, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_UnderflowException, &list, 0, 1, zend.AccInterface)
	SplAddClasses(spl_ce_UnexpectedValueException, &list, 0, 1, zend.AccInterface)

	standard.PhpInfoPrintTableRow(2, "Interfaces", buildClassListString(list.Array()))
	list.Array().DestroyEx()

	zend.ArrayInit(&list)
	SplAddClasses(spl_ce_AppendIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_ArrayIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_ArrayObject, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_BadFunctionCallException, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_BadMethodCallException, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_CachingIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_CallbackFilterIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_DirectoryIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_DomainException, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_EmptyIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_FilesystemIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_FilterIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_GlobIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_InfiniteIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_InvalidArgumentException, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_IteratorIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_LengthException, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_LimitIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_LogicException, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_MultipleIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_NoRewindIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_OuterIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_OutOfBoundsException, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_OutOfRangeException, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_OverflowException, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_ParentIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_RangeException, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveArrayIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveCachingIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveCallbackFilterIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveDirectoryIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveFilterIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveIteratorIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveRegexIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_RecursiveTreeIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_RegexIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_RuntimeException, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SeekableIterator, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplDoublyLinkedList, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplFileInfo, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplFileObject, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplFixedArray, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplHeap, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplMinHeap, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplMaxHeap, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplObjectStorage, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplObserver, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplPriorityQueue, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplQueue, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplStack, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplSubject, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_SplTempFileObject, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_UnderflowException, &list, 0, -1, zend.AccInterface)
	SplAddClasses(spl_ce_UnexpectedValueException, &list, 0, -1, zend.AccInterface)

	standard.PhpInfoPrintTableRow(2, "Classes", buildClassListString(list.Array()))
	list.Array().DestroyEx()

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
	SplAutoloadFn = zend.CG__().FunctionTable().Get("spl_autoload")
	SplAutoloadCallFn = zend.CG__().FunctionTable().Get("spl_autoload_call")
	b.Assert(SplAutoloadFn != nil && SplAutoloadCallFn != nil)
	return types2.SUCCESS
}
func ZmActivateSpl(type_ int, module_number int) int {
	SPL_G__().Reset()
	return types2.SUCCESS
}
func ZmDeactivateSpl(type_ int, module_number int) int {
	SPL_G__().Deactivate()
	return types2.SUCCESS
}
