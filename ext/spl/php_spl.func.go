package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
)

func SPL_G__() *ZendSplGlobals { return &SplGlobals }
func ZmGlobalsCtorSpl(spl_globals *ZendSplGlobals) {
	spl_globals.Ctor()
}
func SplFindCeByName(name string, autoload bool) *types.ClassEntry {
	var ce *types.ClassEntry
	if autoload {
		ce = zend.ZendLookupClassString(name)
	} else {
		ce = zend.EG__().ClassTable().Get(name)
	}
	if ce == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Class %s does not exist%s", name, b.Cond(autoload, " and could not be loaded", ""))
		return nil
	}
	return ce
}
func ZifClassParents(instance *types.Zval, _ zpp.Opt, autoload_ *bool) (*types.Array, bool) {
	var autoload = b.Option(autoload_, true)

	var ce *types.ClassEntry
	if instance.IsObject() {
		ce = instance.Object().GetCe()
	} else if instance.IsString() {
		ce = SplFindCeByName(instance.StringVal(), autoload)
		if ce == nil {
			return nil, false
		}
	} else {
		core.PhpErrorDocref(nil, faults.E_WARNING, "object or string expected")
		return nil, false
	}

	retArr := types.NewArray(0)
	parentClass := ce.GetParent()
	for parentClass != nil {
		SplAddClassName(retArr, parentClass, 0, 0)
		parentClass = parentClass.GetParent()
	}
	return retArr, true
}
func ZifClassImplements(instance *types.Zval, _ zpp.Opt, autoload_ *bool) (*types.Array, bool) {
	var autoload = b.Option(autoload_, true)

	var ce *types.ClassEntry
	if instance.IsObject() {
		ce = instance.Object().GetCe()
	} else if instance.IsString() {
		ce = SplFindCeByName(instance.StringVal(), autoload)
		if ce == nil {
			return nil, false
		}
	} else {
		core.PhpErrorDocref(nil, faults.E_WARNING, "object or string expected")
		return nil, false
	}

	retArr := types.NewArray(0)
	SplAddInterfaces(retArr, ce, 1, types.AccInterface)
	return retArr, true
}
func ZifClassUses(instance *types.Zval, _ zpp.Opt, autoload_ *bool) (*types.Array, bool) {
	var autoload = b.Option(autoload_, true)

	var ce *types.ClassEntry
	if instance.IsObject() {
		ce = instance.Object().GetCe()
	} else if instance.IsString() {
		ce = SplFindCeByName(instance.StringVal(), autoload)
		if ce == nil {
			return nil, false
		}
	} else {
		core.PhpErrorDocref(nil, faults.E_WARNING, "object or string expected")
		return nil, false
	}

	retArr := types.NewArray(0)
	SplAddTraits(retArr, ce, 1, types.AccTrait)
	return retArr, true
}
func splClasses() []*types.ClassEntry {
	return []*types.ClassEntry{
		spl_ce_AppendIterator,
		spl_ce_ArrayIterator,
		spl_ce_ArrayObject,
		spl_ce_BadFunctionCallException,
		spl_ce_BadMethodCallException,
		spl_ce_CachingIterator,
		spl_ce_CallbackFilterIterator,
		spl_ce_DirectoryIterator,
		spl_ce_DomainException,
		spl_ce_EmptyIterator,
		spl_ce_FilesystemIterator,
		spl_ce_FilterIterator,
		spl_ce_GlobIterator,
		spl_ce_InfiniteIterator,
		spl_ce_InvalidArgumentException,
		spl_ce_IteratorIterator,
		spl_ce_LengthException,
		spl_ce_LimitIterator,
		spl_ce_LogicException,
		spl_ce_MultipleIterator,
		spl_ce_NoRewindIterator,
		spl_ce_OuterIterator,
		spl_ce_OutOfBoundsException,
		spl_ce_OutOfRangeException,
		spl_ce_OverflowException,
		spl_ce_ParentIterator,
		spl_ce_RangeException,
		spl_ce_RecursiveArrayIterator,
		spl_ce_RecursiveCachingIterator,
		spl_ce_RecursiveCallbackFilterIterator,
		spl_ce_RecursiveDirectoryIterator,
		spl_ce_RecursiveFilterIterator,
		spl_ce_RecursiveIterator,
		spl_ce_RecursiveIteratorIterator,
		spl_ce_RecursiveRegexIterator,
		spl_ce_RecursiveTreeIterator,
		spl_ce_RegexIterator,
		spl_ce_RuntimeException,
		spl_ce_SeekableIterator,
		spl_ce_SplDoublyLinkedList,
		spl_ce_SplFileInfo,
		spl_ce_SplFileObject,
		spl_ce_SplFixedArray,
		spl_ce_SplHeap,
		spl_ce_SplMinHeap,
		spl_ce_SplMaxHeap,
		spl_ce_SplObjectStorage,
		spl_ce_SplObserver,
		spl_ce_SplPriorityQueue,
		spl_ce_SplQueue,
		spl_ce_SplStack,
		spl_ce_SplSubject,
		spl_ce_SplTempFileObject,
		spl_ce_UnderflowException,
		spl_ce_UnexpectedValueException,
	}
}

func ZifSplClasses() *types.Array {
	classes := splClasses()

	list := types.NewArray(0)
	for _, ce := range classes {
		if ce != nil {
			SplAddClassName(list, ce, 0, 0)
		}
	}
	return list
}
func SplAutoload(className string, lcName string, ext string) int {
	var dummy types.Zval
	var opArray *types.ZendOpArray
	var result types.Zval

	classFile := lcName + ext
	fh := core.PhpStreamOpenForZendEx(classFile, core.USE_PATH|core.STREAM_OPEN_FOR_INCLUDE)
	if fh != nil {
		if fh.GetOpenedPath() == "" {
			fh.SetOpenedPath(classFile)
		}
		openedPath := fh.GetOpenedPath()
		dummy.SetNull()
		if zend.EG__().GetIncludedFiles().KeyAdd(openedPath, &dummy) != nil {
			opArray = zend.CompileFile(fh, zend.ZEND_REQUIRE)
			zend.ZendDestroyFileHandle(fh)
		} else {
			opArray = nil
			fh.Destroy()
		}
		if opArray != nil {
			result.SetUndef()
			zend.ZendExecute(opArray, &result)
			return types.IntBool(zend.EG__().ClassTable().Exists(lcName))
		}
	}
	return 0
}
func ZifSplAutoload(className string, _ zpp.Opt, fileExtensions *string) {
	var fileExts string
	if fileExtensions != nil {
		fileExts = *fileExtensions
	} else {
		fileExts = SPL_G__().GetAutoloadExtensions()
	}

	lcName := ascii.StrToLower(className)
	exts := strings.Split(fileExts, ",")
	for _, ext := range exts {
		if zend.EG__().GetException() != nil {
			break
		}
		if SplAutoload(className, lcName, ext) != 0 {
			break
		}
	}
}
func ZifSplAutoloadExtensions(_ zpp.Opt, fileExtensions *string) string {
	if fileExtensions != nil {
		SPL_G__().SetAutoloadExtensions(*fileExtensions)
	}
	return SPL_G__().GetAutoloadExtensions()
}
func ZifSplAutoloadCall(executeData zpp.Ex, return_value zpp.Ret, className *types.Zval) {
	var class_name *types.Zval
	var retval types.Zval
	var lc_name *types.String
	var alfi *AutoloadFuncInfo
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &class_name) == types.FAILURE || !class_name.IsString() {
		return
	}
	if SPL_G__().autoloadFunctions != nil {
		var func_ types.IFunction
		var fci types.ZendFcallInfo
		var fcic types.ZendFcallInfoCache
		var called_scope *types.ClassEntry = zend.ZendGetCalledScope(executeData)
		var l_autoload_running int = SPL_G__().autoloadRunning
		SPL_G__().SetAutoloadRunning(1)
		lc_name = operators.ZendStringTolower(class_name.String())
		fci.SetSize(b.SizeOf("fci"))
		fci.SetRetval(&retval)
		fci.SetParamCount(1)
		fci.SetParams(class_name)
		fci.SetNoSeparation(1)
		fci.GetFunctionName().SetUndef()

		SPL_G__().autoloadFunctions.ForeachEx(func(key types.ArrayKey, value *types.Zval) bool {
			if !key.IsStrKey() {
				return false
			}

			alfi = value.Ptr()
			func_ = alfi.GetFuncPtr()
			if func_.HasFnFlags(types.AccCallViaTrampoline) {
				func_ = zend.Emalloc(b.SizeOf("zend_op_array"))
				memcpy(func_, alfi.GetFuncPtr(), b.SizeOf("zend_op_array"))
			}
			retval.SetUndef()
			fcic.SetFunctionHandler(func_)
			if alfi.GetObj().IsUndef() {
				fci.SetObject(nil)
				fcic.SetObject(nil)
				if alfi.GetCe() != nil && (called_scope == nil || operators.InstanceofFunction(called_scope, alfi.GetCe()) == 0) {
					fcic.SetCalledScope(alfi.GetCe())
				} else {
					fcic.SetCalledScope(called_scope)
				}
			} else {
				fci.SetObject(alfi.GetObj().Object())
				fcic.SetObject(alfi.GetObj().Object())
				fcic.SetCalledScope(types.Z_OBJCE(alfi.GetObj()))
			}
			zend.ZendCallFunction(&fci, &fcic)
			if zend.EG__().GetException() != nil {
				return false
			}
			if zend.EG__().ClassTable().Exists(lc_name.GetStr()) {
				return false
			}
			return true
		})
		SPL_G__().autoloadRunning = l_autoload_running
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
	}
}

func uintToStr(i uint) string {
	return string([]byte{
		byte(i & 0xff),
		byte(i >> 8 & 0xff),
		byte(i >> 16 & 0xff),
		byte(i >> 24 & 0xff),
		byte(i >> 32 & 0xff),
		byte(i >> 40 & 0xff),
		byte(i >> 48 & 0xff),
		byte(i >> 56 & 0xff),
	})
}

func ZifSplAutoloadRegister(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, autoloadFunction *types.Zval, throw *types.Zval, prepend *types.Zval) {
	var func_name *types.String
	var error *byte = nil
	var lc_name string = ""
	var zcallable *types.Zval = nil
	var do_throw types.ZendBool = 1
	var prepend types.ZendBool = 0
	var spl_func_ptr types.IFunction
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
				if obj_ptr == nil && alfi.GetFuncPtr() != nil && !alfi.GetFuncPtr().HasFnFlags(types.AccStatic) {
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
			} else if zcallable.IsString() {
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
		} else if interFunc, ok := fcc.GetFunctionHandler().(*types.InternalFunction); ok && interFunc.GetHandler() == ZifSplAutoloadCall {
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
		if zcallable.IsType(types.IS_OBJECT) {
			types.ZVAL_COPY(alfi.GetClosure(), zcallable)
			lc_name = ascii.StrToLower(func_name.GetStr()) + uintToStr(zcallable.Object().GetHandle())
		} else {
			alfi.GetClosure().SetUndef()

			/* Skip leading \ */
			if func_name.GetStr()[0] == '\\' {
				lc_name = ascii.StrToLower(func_name.GetStr()[1:])
			} else {
				lc_name = ascii.StrToLower(func_name.GetStr())
			}
		}
		// types.ZendStringReleaseEx(func_name, 0)
		if SPL_G__().autoloadFunctions != nil && SPL_G__().autoloadFunctions.KeyExists(lc_name) {
			goto skip
		}
		if obj_ptr != nil && !alfi.GetFuncPtr().HasFnFlags(types.AccStatic) {

			/* add object id to the hash to ensure uniqueness, for more reference look at bug #40091 */
			lc_name += uintToStr(obj_ptr.GetHandle())
			alfi.GetObj().SetObject(obj_ptr)
		} else {
			alfi.GetObj().SetUndef()
		}
		if SPL_G__().autoloadFunctions == nil {
			SPL_G__().autoloadFunctions = types.NewArray(1)
		}
		spl_func_ptr = SplAutoloadFn
		if zend.EG__().GetAutoloadFunc() == spl_func_ptr {
			var spl_alfi AutoloadFuncInfo
			spl_alfi.SetFuncPtr(spl_func_ptr)
			spl_alfi.GetObj().SetUndef()
			spl_alfi.GetClosure().SetUndef()
			spl_alfi.SetCe(nil)
			types.ZendHashAddMem(SPL_G__().autoloadFunctions, SplAutoloadFn.FunctionName(), &spl_alfi, b.SizeOf("autoload_func_info"))
			if prepend != 0 && SPL_G__().autoloadFunctions.Len() > 1 {
				/* Move the newly created element to the head of the hashtable */
				SPL_G__().autoloadFunctions.MoveTailToHead()
			}
		}
		if alfi.GetFuncPtr() == zend.EG__().GetTrampoline() {
			var copy types.IFunction = zend.Emalloc(b.SizeOf("zend_op_array"))
			memcpy(copy, alfi.GetFuncPtr(), b.SizeOf("zend_op_array"))
			alfi.GetFuncPtr().SetFunctionName("")
			alfi.SetFuncPtr(copy)
		}
		if types.ZendHashAddMem(SPL_G__().autoloadFunctions, lc_name, &alfi, b.SizeOf("autoload_func_info")) == nil {
			if alfi.GetFuncPtr().HasFnFlags(types.AccCallViaTrampoline) {
				zend.ZendFreeTrampoline(alfi.GetFuncPtr())
			}
		}
		if prepend != 0 && SPL_G__().autoloadFunctions.Len() > 1 {
			/* Move the newly created element to the head of the hashtable */
			SPL_G__().autoloadFunctions.MoveTailToHead()
		}
	skip:
		// types.ZendStringReleaseEx(lc_name, 0)
	}
	if SPL_G__().autoloadFunctions {
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
	var lc_name string
	var zcallable *types.Zval
	var success bool = false
	var spl_func_ptr types.IFunction
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
			// types.ZendStringReleaseEx(func_name, 0)
		}
		return_value.SetFalse()
		return
	}
	obj_ptr = fcc.GetObject()
	if error != nil {
		zend.Efree(error)
	}
	if zcallable.IsType(types.IS_OBJECT) {
		lc_name = ascii.StrToLower(func_name.GetStr()) + uintToStr(zcallable.Object().GetHandle())
	} else {

		/* Skip leading \ */

		if func_name.GetStr()[0] == '\\' {
			lc_name = ascii.StrToLower(func_name.GetStr()[1:])
		} else {
			lc_name = ascii.StrToLower(func_name.GetStr())
		}
	}
	// types.ZendStringReleaseEx(func_name, 0)
	if SPL_G__().autoloadFunctions {
		if lc_name == SplAutoloadCallFn.FunctionName() {
			/* remove all */
			if !(SPL_G__().autoloadRunning) {
				SPL_G__().autoloadFunctions.Destroy()
				SPL_G__().autoloadFunctions = nil
				zend.EG__().SetAutoloadFunc(nil)
			} else {
				SPL_G__().autoloadFunctions.Clean()
			}
			success = true
		} else {
			/* remove specific */
			success = SPL_G__().autoloadFunctions.KeyDelete(lc_name)
			if !success && obj_ptr != nil {
				lc_name += uintToStr(obj_ptr.GetHandle())
				success = SPL_G__().autoloadFunctions.KeyDelete(lc_name)
			}
		}
	} else if lc_name == SplAutoloadFn.FunctionName() {

		/* register single spl_autoload() */

		spl_func_ptr = SplAutoloadFn
		if zend.EG__().GetAutoloadFunc() == spl_func_ptr {
			success = true
			zend.EG__().SetAutoloadFunc(nil)
		}
	}
	// types.ZendStringReleaseEx(lc_name, 0)
	return_value.SetBool(success)
	return
}
func ZifSplAutoloadFunctions(executeData zpp.Ex, return_value zpp.Ret) {
	var fptr types.IFunction
	var alfi *AutoloadFuncInfo
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if zend.EG__().GetAutoloadFunc() == nil {
		if fptr := zend.EG__().FunctionTable().Get(types.STR_MAGIC_AUTOLOAD); fptr != nil {
			var tmp types.Zval
			zend.ArrayInit(return_value)
			tmp.SetStringVal(types.STR_MAGIC_AUTOLOAD)
			return_value.Array().AppendNew(&tmp)
			return
		}
		return_value.SetFalse()
		return
	}
	fptr = SplAutoloadCallFn
	if zend.EG__().GetAutoloadFunc() == fptr {
		zend.ArrayInit(return_value)
		SPL_G__().autoloadFunctions.Foreach(func(key_ types.ArrayKey, value *types.Zval) {
			alfi = value.Ptr()
			if !(alfi.GetClosure().IsUndef()) {
				zend.AddNextIndexZval(return_value, alfi.GetClosure())
			} else if alfi.GetFuncPtr().GetScope() != nil {
				tmpArr := types.NewArray(0)
				if !(alfi.GetObj().IsUndef()) {
					tmpArr.Append(alfi.GetObj())
				} else {
					tmpArr.Append(types.NewZvalString(alfi.GetCe().Name()))
				}
				tmpArr.Append(types.NewZvalString(alfi.GetFuncPtr().FunctionName()))

				return_value.Array().Append(types.NewZvalArray(tmpArr))
			} else {
				if alfi.GetFuncPtr().FunctionName() != "__lambda_func" {
					zend.AddNextIndexStrEx(return_value, alfi.GetFuncPtr().FunctionName())
				} else {
					zend.AddNextIndexStrEx(return_value, key_.StrKey())
				}
			}
		})
		return
	}
	zend.ArrayInit(return_value)
	return_value.Array().Append(types.NewZvalString(zend.EG__().GetAutoloadFunc().FunctionName()))
}

//@zif -old
func ZifSplObjectHash(obj zpp.Object) string {
	return PhpSplObjectHash(obj)
}
func ZifSplObjectId(obj zpp.Object) int {
	return int(obj.Object().GetHandle())
}
func PhpSplObjectHash(obj *types.Zval) string {
	return SPL_G__().SplObjectHash(obj.Object().GetHandle())
}
func buildClassListString(arr *types.Array) string {
	names := make([]string, 0, arr.Len())
	arr.Foreach(func(key types.ArrayKey, value *types.Zval) {
		names = append(names, value.StringVal())
	})
	return strings.Join(names, ", ")
}

func ZmInfoSpl(zend_module *zend.ModuleEntry) {
	classes := splClasses()

	standard.PhpInfoPrintTableStart()
	standard.PhpInfoPrintTableHeader(2, "SPL support", "enabled")

	var interfaceNames strings.Builder
	var classNames strings.Builder
	for _, ce := range classes {
		if ce.HasCeFlags(types.AccInterface) {
			if interfaceNames.Len() > 0 {
				interfaceNames.WriteString(", ")
			}
			interfaceNames.WriteString(ce.Name())
		} else {
			if classNames.Len() > 0 {
				classNames.WriteString(", ")
			}
			classNames.WriteString(ce.Name())
		}
	}
	standard.PhpInfoPrintTableRow(2, "Interfaces", interfaceNames.String())
	standard.PhpInfoPrintTableRow(2, "Classes", classNames.String())

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
	return types.SUCCESS
}
func ZmActivateSpl(type_ int, module_number int) int {
	SPL_G__().Reset()
	return types.SUCCESS
}
func ZmDeactivateSpl(type_ int, module_number int) int {
	SPL_G__().Deactivate()
	return types.SUCCESS
}
