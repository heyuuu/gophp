package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
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
		core.PhpErrorDocref("", faults.E_WARNING, "Class %s does not exist%s", name, lang.Cond(autoload, " and could not be loaded", ""))
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
		ce = SplFindCeByName(instance.String(), autoload)
		if ce == nil {
			return nil, false
		}
	} else {
		core.PhpErrorDocref("", faults.E_WARNING, "object or string expected")
		return nil, false
	}

	retArr := types.NewArray()
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
		ce = SplFindCeByName(instance.String(), autoload)
		if ce == nil {
			return nil, false
		}
	} else {
		core.PhpErrorDocref("", faults.E_WARNING, "object or string expected")
		return nil, false
	}

	retArr := types.NewArray()
	SplAddInterfaces(retArr, ce, 1, types.AccInterface)
	return retArr, true
}
func ZifClassUses(instance *types.Zval, _ zpp.Opt, autoload_ *bool) (*types.Array, bool) {
	var autoload = b.Option(autoload_, true)

	var ce *types.ClassEntry
	if instance.IsObject() {
		ce = instance.Object().GetCe()
	} else if instance.IsString() {
		ce = SplFindCeByName(instance.String(), autoload)
		if ce == nil {
			return nil, false
		}
	} else {
		core.PhpErrorDocref("", faults.E_WARNING, "object or string expected")
		return nil, false
	}

	retArr := types.NewArray()
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

	list := types.NewArray()
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

//@zif -old "z"
func ZifSplAutoloadCall(executeData zpp.Ex, className string) {
	var retval types.Zval
	var alfi *AutoloadFuncInfo

	if SPL_G__().autoloadFunctions != nil {
		var func_ types.IFunction
		var fci *types.ZendFcallInfo = types.InitFCallInfo(nil, &retval, types.NewZvalString(className))
		var fcic types.ZendFcallInfoCache
		var called_scope *types.ClassEntry = zend.ZendGetCalledScope(executeData)
		var l_autoload_running int = SPL_G__().autoloadRunning
		SPL_G__().SetAutoloadRunning(1)
		lc_name := ascii.StrToLower(className)

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
			zend.ZendCallFunction(fci, &fcic)
			if zend.EG__().GetException() != nil {
				return false
			}
			if zend.EG__().ClassTable().Exists(lc_name) {
				return false
			}
			return true
		})
		SPL_G__().autoloadRunning = l_autoload_running
	} else {
		/* do not use or overwrite &EG(autoload_func) here */
		retval.SetUndef()

		var fcallInfo = types.InitFCallInfo(nil, &retval, types.NewZvalString(className))

		var fcallCache types.ZendFcallInfoCache
		fcallCache.SetFunctionHandler(SplAutoloadFn)
		fcallCache.SetCalledScope(nil)
		fcallCache.SetObject(nil)

		zend.ZendCallFunction(fcallInfo, &fcallCache)
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

//@zif -old "|zbb"
func ZifSplAutoloadRegister(executeData zpp.Ex, _ zpp.Opt, autoloadFunction *types.Zval, throw_ *bool, prepend bool) bool {
	var zcallable *types.Zval = autoloadFunction
	var throw = b.Option(throw_, true)

	var funcName *types.String
	var error_ *byte = nil
	var lcName string = ""
	var splFuncPtr types.IFunction
	var alfi AutoloadFuncInfo
	var objPtr *types.Object
	var fcc types.ZendFcallInfoCache
	if executeData.NumArgs() != 0 {
		if !zend.ZendIsCallableEx(zcallable, nil, zend.IS_CALLABLE_STRICT, &funcName, &fcc, &error_) {
			alfi.SetCe(fcc.GetCallingScope())
			alfi.SetFuncPtr(fcc.GetFunctionHandler())
			objPtr = fcc.GetObject()
			if zcallable.IsType(types.IsArray) {
				if objPtr == nil && alfi.GetFuncPtr() != nil && !alfi.GetFuncPtr().HasFnFlags(types.AccStatic) {
					if throw {
						faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Passed array specifies a non static method but no object (%s)", error_)
					}
					return false
				} else if throw {
					faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Passed array does not specify %s %smethod (%s)", lang.Cond(alfi.GetFuncPtr() != nil, "a callable", "an existing"), lang.Cond(objPtr == nil, "static ", ""), error_)
				}
				return false
			} else if zcallable.IsString() {
				if throw {
					faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Function '%s' not %s (%s)", funcName.GetVal(), lang.Cond(alfi.GetFuncPtr() != nil, "callable", "found"), error_)
				}
				return false
			} else {
				if throw {
					faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Illegal value passed (%s)", error_)
				}
				return false
			}
		} else if interFunc, ok := fcc.GetFunctionHandler().(*types.InternalFunction); ok && interFunc.GetHandler() == ZifSplAutoloadCall {
			if throw {
				faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Function spl_autoload_call() cannot be registered")
			}
			return false
		}
		alfi.SetCe(fcc.GetCallingScope())
		alfi.SetFuncPtr(fcc.GetFunctionHandler())
		objPtr = fcc.GetObject()
		if zcallable.IsType(types.IsObject) {
			types.ZVAL_COPY(alfi.GetClosure(), zcallable)
			lcName = ascii.StrToLower(funcName.GetStr()) + uintToStr(zcallable.Object().GetHandle())
		} else {
			alfi.GetClosure().SetUndef()

			/* Skip leading \ */
			if funcName.GetStr()[0] == '\\' {
				lcName = ascii.StrToLower(funcName.GetStr()[1:])
			} else {
				lcName = ascii.StrToLower(funcName.GetStr())
			}
		}
		if SPL_G__().autoloadFunctions != nil && SPL_G__().autoloadFunctions.KeyExists(lcName) {
			goto skip
		}
		if objPtr != nil && !alfi.GetFuncPtr().HasFnFlags(types.AccStatic) {
			/* add object id to the hash to ensure uniqueness, for more reference look at bug #40091 */
			lcName += uintToStr(objPtr.GetHandle())
			alfi.GetObj().SetObject(objPtr)
		} else {
			alfi.GetObj().SetUndef()
		}
		if SPL_G__().autoloadFunctions == nil {
			SPL_G__().autoloadFunctions = types.NewArrayCap(1)
		}
		splFuncPtr = SplAutoloadFn
		if zend.EG__().GetAutoloadFunc() == splFuncPtr {
			var spl_alfi AutoloadFuncInfo
			spl_alfi.SetFuncPtr(splFuncPtr)
			spl_alfi.GetObj().SetUndef()
			spl_alfi.GetClosure().SetUndef()
			spl_alfi.SetCe(nil)
			types.ZendHashAddMem(SPL_G__().autoloadFunctions, SplAutoloadFn.FunctionName(), &spl_alfi, b.SizeOf("autoload_func_info"))
			if prepend && SPL_G__().autoloadFunctions.Len() > 1 {
				/* Move the newly created element to the head of the hashtable */
				SPL_G__().autoloadFunctions.MoveTailToHead()
			}
		}
		if alfi.GetFuncPtr() == zend.EG__().GetTrampoline() {
			var copy_ types.IFunction = zend.Emalloc(b.SizeOf("zend_op_array"))
			memcpy(copy_, alfi.GetFuncPtr(), b.SizeOf("zend_op_array"))
			alfi.GetFuncPtr().SetFunctionName("")
			alfi.SetFuncPtr(copy_)
		}
		if types.ZendHashAddMem(SPL_G__().autoloadFunctions, lcName, &alfi, b.SizeOf("autoload_func_info")) == nil {
			if alfi.GetFuncPtr().HasFnFlags(types.AccCallViaTrampoline) {
				zend.ZendFreeTrampoline(alfi.GetFuncPtr())
			}
		}
		if prepend && SPL_G__().autoloadFunctions.Len() > 1 {
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
	return true
}

//@zif -old "z"
func ZifSplAutoloadUnregister(autoloadFunction *types.Zval) bool {
	var zcallable *types.Zval = autoloadFunction

	var funcName *types.String = nil
	var error_ *byte = nil
	var lcName string
	var success bool = false
	var splFuncPtr types.IFunction
	var objPtr *types.Object
	var fcc types.ZendFcallInfoCache

	if zend.ZendIsCallableEx(zcallable, nil, zend.IS_CALLABLE_CHECK_SYNTAX_ONLY, &funcName, &fcc, &error_) == 0 {
		faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Unable to unregister invalid function (%s)", error_)
		return false
	}
	objPtr = fcc.GetObject()
	if zcallable.IsType(types.IsObject) {
		lcName = ascii.StrToLower(funcName.GetStr()) + uintToStr(zcallable.Object().GetHandle())
	} else {
		/* Skip leading \ */
		if funcName.GetStr()[0] == '\\' {
			lcName = ascii.StrToLower(funcName.GetStr()[1:])
		} else {
			lcName = ascii.StrToLower(funcName.GetStr())
		}
	}
	if SPL_G__().autoloadFunctions != nil {
		if lcName == SplAutoloadCallFn.FunctionName() {
			/* remove all */
			if SPL_G__().autoloadRunning == 0 {
				SPL_G__().autoloadFunctions.Destroy()
				SPL_G__().autoloadFunctions = nil
				zend.EG__().SetAutoloadFunc(nil)
			} else {
				SPL_G__().autoloadFunctions.Clean()
			}
			success = true
		} else {
			/* remove specific */
			success = SPL_G__().autoloadFunctions.KeyDelete(lcName)
			if !success && objPtr != nil {
				lcName += uintToStr(objPtr.GetHandle())
				success = SPL_G__().autoloadFunctions.KeyDelete(lcName)
			}
		}
	} else if lcName == SplAutoloadFn.FunctionName() {
		/* register single spl_autoload() */
		splFuncPtr = SplAutoloadFn
		if zend.EG__().GetAutoloadFunc() == splFuncPtr {
			success = true
			zend.EG__().SetAutoloadFunc(nil)
		}
	}
	return success
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
			tmp.SetString(types.STR_MAGIC_AUTOLOAD)
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
				tmpArr := types.NewArray()
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
		names = append(names, value.String())
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
