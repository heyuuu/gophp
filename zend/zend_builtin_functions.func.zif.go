package zend

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifZendVersion
var DefZifZendVersion = def.DefFunc("zend_version", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifZendVersion()
	returnValue.SetStringVal(ret)
})

// generate by ZifGcMemCaches
var DefZifGcMemCaches = def.DefFunc("gc_mem_caches", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGcMemCaches()
	returnValue.SetLong(ret)
})

// generate by ZifGcCollectCycles
var DefZifGcCollectCycles = def.DefFunc("gc_collect_cycles", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGcCollectCycles()
	returnValue.SetLong(ret)
})

// generate by ZifGcEnabled
var DefZifGcEnabled = def.DefFunc("gc_enabled", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGcEnabled()
	returnValue.SetBool(ret)
})

// generate by ZifGcEnable
var DefZifGcEnable = def.DefFunc("gc_enable", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGcEnable()
})

// generate by ZifGcDisable
var DefZifGcDisable = def.DefFunc("gc_disable", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGcDisable()
})

// generate by ZifGcStatus
var DefZifGcStatus = def.DefFunc("gc_status", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGcStatus(returnValue)
})

// generate by ZifFuncNumArgs
var DefZifFuncNumArgs = def.DefFunc("func_num_args", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifFuncNumArgs(executeData)
	returnValue.SetLong(ret)
})

// generate by ZifFuncGetArg
var DefZifFuncGetArg = def.DefFunc("func_get_arg", 1, 1, []def.ArgInfo{{Name: "arg_num"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, zpp.FlagOldMode)
	arg_num := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ZifFuncGetArg(executeData, returnValue, arg_num)
})

// generate by ZifFuncGetArgs
var DefZifFuncGetArgs = def.DefFunc("func_get_args", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret, ok := ZifFuncGetArgs(executeData)
	if ok {
		returnValue.SetArray(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrlen
var DefZifStrlen = def.DefFunc("strlen", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStrlen(str)
	returnValue.SetLong(ret)
})

// generate by ZifStrcmp
var DefZifStrcmp = def.DefFunc("strcmp", 2, 2, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	str1 := fp.ParseStringVal()
	str2 := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStrcmp(str1, str2)
	returnValue.SetLong(ret)
})

// generate by ZifStrncmp
var DefZifStrncmp = def.DefFunc("strncmp", 3, 3, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}, {Name: "len_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	str1 := fp.ParseStringVal()
	str2 := fp.ParseStringVal()
	len_ := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrncmp(str1, str2, len_)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStrcasecmp
var DefZifStrcasecmp = def.DefFunc("strcasecmp", 2, 2, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	str1 := fp.ParseStringVal()
	str2 := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStrcasecmp(str1, str2)
	returnValue.SetLong(ret)
})

// generate by ZifStrncasecmp
var DefZifStrncasecmp = def.DefFunc("strncasecmp", 3, 3, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}, {Name: "len_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	str1 := fp.ParseStringVal()
	str2 := fp.ParseStringVal()
	len_ := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifStrncasecmp(str1, str2, len_)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifEach
var DefZifEach = def.DefFunc("each", 1, 1, []def.ArgInfo{{Name: "arr"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arr := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifEach(executeData, returnValue, arr)
})

// generate by ZifErrorReporting
var DefZifErrorReporting = def.DefFunc("error_reporting", 0, 1, []def.ArgInfo{{Name: "new_error_level"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	new_error_level := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifErrorReporting(returnValue, nil, new_error_level)
})

// generate by ZifDefine
var DefZifDefine = def.DefFunc("define", 2, 3, []def.ArgInfo{{Name: "constant_name"}, {Name: "value"}, {Name: "case_insensitive"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	constant_name := fp.ParseStringVal()
	value := fp.ParseZval()
	fp.StartOptional()
	case_insensitive := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifDefine(constant_name, value, nil, case_insensitive)
	returnValue.SetBool(ret)
})

// generate by ZifDefined
var DefZifDefined = def.DefFunc("defined", 1, 1, []def.ArgInfo{{Name: "constant_name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	constant_name := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifDefined(constant_name)
	returnValue.SetBool(ret)
})

// generate by ZifGetClass
var DefZifGetClass = def.DefFunc("get_class", 0, 1, []def.ArgInfo{{Name: "object"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	object := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetClass(executeData, returnValue, nil, object)
})

// generate by ZifGetCalledClass
var DefZifGetCalledClass = def.DefFunc("get_called_class", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret, ok := ZifGetCalledClass(executeData)
	if ok {
		returnValue.SetStringVal(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifGetParentClass
var DefZifGetParentClass = def.DefFunc("get_parent_class", 0, 1, []def.ArgInfo{{Name: "object"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	object := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetParentClass(executeData, returnValue, nil, object)
})

// generate by ZifIsSubclassOf
var DefZifIsSubclassOf = def.DefFunc("is_subclass_of", 2, 3, []def.ArgInfo{{Name: "object"}, {Name: "class_name"}, {Name: "allow_string"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	object := fp.ParseZval()
	class_name := fp.ParseZval()
	fp.StartOptional()
	allow_string := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsSubclassOf(executeData, returnValue, object, class_name, nil, allow_string)
})

// generate by ZifIsA
var DefZifIsA = def.DefFunc("is_a", 2, 3, []def.ArgInfo{{Name: "object"}, {Name: "class_name"}, {Name: "allow_string"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	object := fp.ParseZval()
	class_name := fp.ParseZval()
	fp.StartOptional()
	allow_string := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsA(executeData, returnValue, object, class_name, nil, allow_string)
})

// generate by ZifGetClassVars
var DefZifGetClassVars = def.DefFunc("get_class_vars", 1, 1, []def.ArgInfo{{Name: "class_name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	class_name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetClassVars(executeData, returnValue, class_name)
})

// generate by ZifGetObjectVars
var DefZifGetObjectVars = def.DefFunc("get_object_vars", 1, 1, []def.ArgInfo{{Name: "obj"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	obj := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetObjectVars(executeData, returnValue, obj)
})

// generate by ZifGetMangledObjectVars
var DefZifGetMangledObjectVars = def.DefFunc("get_mangled_object_vars", 1, 1, []def.ArgInfo{{Name: "obj"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	obj := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetMangledObjectVars(executeData, returnValue, obj)
})

// generate by ZifGetClassMethods
var DefZifGetClassMethods = def.DefFunc("get_class_methods", 1, 1, []def.ArgInfo{{Name: "class"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	class := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetClassMethods(executeData, returnValue, class)
})

// generate by ZifMethodExists
var DefZifMethodExists = def.DefFunc("method_exists", 2, 2, []def.ArgInfo{{Name: "object"}, {Name: "method"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	object := fp.ParseZval()
	method := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMethodExists(executeData, returnValue, object, method)
})

// generate by ZifPropertyExists
var DefZifPropertyExists = def.DefFunc("property_exists", 2, 2, []def.ArgInfo{{Name: "object_or_class"}, {Name: "property_name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	object_or_class := fp.ParseZval()
	property_name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPropertyExists(executeData, returnValue, object_or_class, property_name)
})

// generate by ZifClassExists
var DefZifClassExists = def.DefFunc("class_exists", 1, 2, []def.ArgInfo{{Name: "classname"}, {Name: "autoload"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	classname := fp.ParseZval()
	fp.StartOptional()
	autoload := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifClassExists(executeData, returnValue, classname, nil, autoload)
})

// generate by ZifInterfaceExists
var DefZifInterfaceExists = def.DefFunc("interface_exists", 1, 2, []def.ArgInfo{{Name: "classname"}, {Name: "autoload"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	classname := fp.ParseZval()
	fp.StartOptional()
	autoload := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifInterfaceExists(executeData, returnValue, classname, nil, autoload)
})

// generate by ZifTraitExists
var DefZifTraitExists = def.DefFunc("trait_exists", 1, 2, []def.ArgInfo{{Name: "traitname"}, {Name: "autoload"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	traitname := fp.ParseZval()
	fp.StartOptional()
	autoload := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTraitExists(executeData, returnValue, traitname, nil, autoload)
})

// generate by ZifFunctionExists
var DefZifFunctionExists = def.DefFunc("function_exists", 1, 1, []def.ArgInfo{{Name: "function_name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	function_name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFunctionExists(executeData, returnValue, function_name)
})

// generate by ZifClassAlias
var DefZifClassAlias = def.DefFunc("class_alias", 2, 3, []def.ArgInfo{{Name: "user_class_name"}, {Name: "alias_name"}, {Name: "autoload"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	user_class_name := fp.ParseZval()
	alias_name := fp.ParseZval()
	fp.StartOptional()
	autoload := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifClassAlias(executeData, returnValue, user_class_name, alias_name, nil, autoload)
})

// generate by ZifGetIncludedFiles
var DefZifGetIncludedFiles = def.DefFunc("get_included_files", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetIncludedFiles(executeData, returnValue)
})

// generate by ZifGetIncludedFiles
var DefZifGetRequiredFiles = def.DefFunc("get_required_files", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetIncludedFiles(executeData, returnValue)
})

// generate by ZifTriggerError
var DefZifTriggerError = def.DefFunc("trigger_error", 1, 2, []def.ArgInfo{{Name: "message"}, {Name: "error_type"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	message := fp.ParseZval()
	fp.StartOptional()
	error_type := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTriggerError(executeData, returnValue, message, nil, error_type)
})

// generate by ZifTriggerError
var DefZifUserError = def.DefFunc("user_error", 1, 2, []def.ArgInfo{{Name: "message"}, {Name: "error_type"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	message := fp.ParseZval()
	fp.StartOptional()
	error_type := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTriggerError(executeData, returnValue, message, nil, error_type)
})

// generate by ZifSetErrorHandler
var DefZifSetErrorHandler = def.DefFunc("set_error_handler", 1, 2, []def.ArgInfo{{Name: "error_handler"}, {Name: "error_types"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	error_handler := fp.ParseZval()
	fp.StartOptional()
	error_types := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSetErrorHandler(executeData, returnValue, error_handler, nil, error_types)
})

// generate by ZifRestoreErrorHandler
var DefZifRestoreErrorHandler = def.DefFunc("restore_error_handler", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifRestoreErrorHandler(executeData, returnValue)
})

// generate by ZifSetExceptionHandler
var DefZifSetExceptionHandler = def.DefFunc("set_exception_handler", 1, 1, []def.ArgInfo{{Name: "exception_handler"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	exception_handler := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSetExceptionHandler(executeData, returnValue, exception_handler)
})

// generate by ZifRestoreExceptionHandler
var DefZifRestoreExceptionHandler = def.DefFunc("restore_exception_handler", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifRestoreExceptionHandler(executeData, returnValue)
})

// generate by ZifGetDeclaredTraits
var DefZifGetDeclaredTraits = def.DefFunc("get_declared_traits", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetDeclaredTraits(executeData, returnValue)
})

// generate by ZifGetDeclaredClasses
var DefZifGetDeclaredClasses = def.DefFunc("get_declared_classes", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetDeclaredClasses(executeData, returnValue)
})

// generate by ZifGetDeclaredInterfaces
var DefZifGetDeclaredInterfaces = def.DefFunc("get_declared_interfaces", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetDeclaredInterfaces(executeData, returnValue)
})

// generate by ZifGetDefinedFunctions
var DefZifGetDefinedFunctions = def.DefFunc("get_defined_functions", 0, 1, []def.ArgInfo{{Name: "exclude_disabled"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	exclude_disabled := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetDefinedFunctions(executeData, returnValue, nil, exclude_disabled)
})

// generate by ZifGetDefinedVars
var DefZifGetDefinedVars = def.DefFunc("get_defined_vars", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetDefinedVars(executeData, returnValue)
})

// generate by ZifCreateFunction
var DefZifCreateFunction = def.DefFunc("create_function", 2, 2, []def.ArgInfo{{Name: "args"}, {Name: "code"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	args := fp.ParseZval()
	code := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCreateFunction(executeData, returnValue, args, code)
})

// generate by ZifGetResourceType
var DefZifGetResourceType = def.DefFunc("get_resource_type", 1, 1, []def.ArgInfo{{Name: "res"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	res := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetResourceType(executeData, returnValue, res)
})

// generate by ZifGetResources
var DefZifGetResources = def.DefFunc("get_resources", 0, 1, []def.ArgInfo{{Name: "type_"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	type_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetResources(executeData, returnValue, nil, type_)
})

// generate by ZifGetLoadedExtensions
var DefZifGetLoadedExtensions = def.DefFunc("get_loaded_extensions", 0, 1, []def.ArgInfo{{Name: "zend_extensions"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	zend_extensions := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetLoadedExtensions(executeData, returnValue, nil, zend_extensions)
})

// generate by ZifGetDefinedConstants
var DefZifGetDefinedConstants = def.DefFunc("get_defined_constants", 0, 1, []def.ArgInfo{{Name: "categorize"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	categorize := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetDefinedConstants(executeData, returnValue, nil, categorize)
})

// generate by ZifDebugPrintBacktrace
var DefZifDebugPrintBacktrace = def.DefFunc("debug_print_backtrace", 0, 2, []def.ArgInfo{{Name: "options"}, {Name: "limit"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	options := fp.ParseZval()
	limit := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDebugPrintBacktrace(executeData, returnValue, nil, options, limit)
})

// generate by ZifDebugBacktrace
var DefZifDebugBacktrace = def.DefFunc("debug_backtrace", 0, 2, []def.ArgInfo{{Name: "options"}, {Name: "limit"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	options := fp.ParseZval()
	limit := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDebugBacktrace(executeData, returnValue, nil, options, limit)
})

// generate by ZifExtensionLoaded
var DefZifExtensionLoaded = def.DefFunc("extension_loaded", 1, 1, []def.ArgInfo{{Name: "extension_name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	extension_name := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifExtensionLoaded(extension_name)
	returnValue.SetBool(ret)
})

// generate by ZifGetExtensionFuncs
var DefZifGetExtensionFuncs = def.DefFunc("get_extension_funcs", 1, 1, []def.ArgInfo{{Name: "extension_name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	extension_name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetExtensionFuncs(executeData, returnValue, extension_name)
})
