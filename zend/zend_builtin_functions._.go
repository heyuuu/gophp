package zend

import (
	b "sik/builtin"
	"sik/zend/types"
)

var ArginfoEach = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("arr", ArgInfoByRef(1)),
}
var ArginfoDefine = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("constant_name"),
	MakeArgInfo("value"),
	MakeArgInfo("case_insensitive"),
}
var ArginfoDefined = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("constant_name"),
}
var ArginfoGetClass = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("object"),
}
var ArginfoIsSubclassOf = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("object"),
	MakeArgInfo("class_name"),
	MakeArgInfo("allow_string"),
}
var ArginfoGetClassVars = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("class_name"),
}
var ArginfoGetObjectVars = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("obj"),
}
var ArginfoGetMangledObjectVars = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("obj"),
}
var ArginfoGetClassMethods = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("class"),
}
var ArginfoMethodExists = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("object"),
	MakeArgInfo("method"),
}
var ArginfoPropertyExists = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("object_or_class"),
	MakeArgInfo("property_name"),
}
var ArginfoClassExists = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("classname"),
	MakeArgInfo("autoload"),
}
var ArginfoTraitExists = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("traitname"),
	MakeArgInfo("autoload"),
}
var ArginfoFunctionExists = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("function_name"),
}
var ArginfoClassAlias = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("user_class_name"),
	MakeArgInfo("alias_name"),
	MakeArgInfo("autoload"),
}
var ArginfoTriggerError = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("message"),
	MakeArgInfo("error_type"),
}
var ArginfoSetErrorHandler = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("error_handler"),
	MakeArgInfo("error_types"),
}
var ArginfoSetExceptionHandler = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("exception_handler"),
}
var ArginfoGetDefinedFunctions = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("exclude_disabled"),
}
var ArginfoCreateFunction = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("args"),
	MakeArgInfo("code"),
}
var ArginfoGetResourceType = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("res"),
}
var ArginfoGetResources = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("type"),
}
var ArginfoGetLoadedExtensions = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("zend_extensions"),
}
var ArginfoGetDefinedConstants = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("categorize"),
}
var ArginfoDebugBacktrace = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("options"),
	MakeArgInfo("limit"),
}
var ArginfoDebugPrintBacktrace = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("options"),
	MakeArgInfo("limit"),
}
var ArginfoExtensionLoaded = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("extension_name"),
}
var BuiltinFunctions = []types.ZendFunctionEntry{
	DefZifZendVersion,
	DefZifFuncNumArgs,
	DefZifFuncGetArg,
	DefZifFuncGetArgs,
	DefZifStrlen,
	DefZifStrcmp,
	DefZifStrncmp,
	DefZifStrcasecmp,
	DefZifStrncasecmp,
	types.MakeZendFunctionEntryEx("each", 0, ZifEach, ArginfoEach),
	DefZifErrorReporting,
	DefZifDefine,
	DefZifDefined,
	types.MakeZendFunctionEntryEx("get_class", 0, ZifGetClass, ArginfoGetClass),
	DefZifGetCalledClass,
	types.MakeZendFunctionEntryEx("get_parent_class", 0, ZifGetParentClass, ArginfoGetClass),
	types.MakeZendFunctionEntryEx("method_exists", 0, ZifMethodExists, ArginfoMethodExists),
	types.MakeZendFunctionEntryEx("property_exists", 0, ZifPropertyExists, ArginfoPropertyExists),
	types.MakeZendFunctionEntryEx("class_exists", 0, ZifClassExists, ArginfoClassExists),
	types.MakeZendFunctionEntryEx("interface_exists", 0, ZifInterfaceExists, ArginfoClassExists),
	types.MakeZendFunctionEntryEx("trait_exists", 0, ZifTraitExists, ArginfoTraitExists),
	types.MakeZendFunctionEntryEx("function_exists", 0, ZifFunctionExists, ArginfoFunctionExists),
	types.MakeZendFunctionEntryEx("class_alias", 0, ZifClassAlias, ArginfoClassAlias),
	types.MakeZendFunctionEntryEx("get_included_files", 0, ZifGetIncludedFiles, nil),
	types.MakeZendFunctionEntryEx("get_required_files", 0, ZifGetIncludedFiles, nil),
	types.MakeZendFunctionEntryEx("is_subclass_of", 0, ZifIsSubclassOf, ArginfoIsSubclassOf),
	types.MakeZendFunctionEntryEx("is_a", 0, ZifIsA, ArginfoIsSubclassOf),
	types.MakeZendFunctionEntryEx("get_class_vars", 0, ZifGetClassVars, ArginfoGetClassVars),
	types.MakeZendFunctionEntryEx("get_object_vars", 0, ZifGetObjectVars, ArginfoGetObjectVars),
	types.MakeZendFunctionEntryEx("get_mangled_object_vars", 0, ZifGetMangledObjectVars, ArginfoGetMangledObjectVars),
	types.MakeZendFunctionEntryEx("get_class_methods", 0, ZifGetClassMethods, ArginfoGetClassMethods),
	types.MakeZendFunctionEntryEx("trigger_error", 0, ZifTriggerError, ArginfoTriggerError),
	types.MakeZendFunctionEntryEx("user_error", 0, ZifTriggerError, ArginfoTriggerError),
	types.MakeZendFunctionEntryEx("set_error_handler", 0, ZifSetErrorHandler, ArginfoSetErrorHandler),
	types.MakeZendFunctionEntryEx("restore_error_handler", 0, ZifRestoreErrorHandler, nil),
	types.MakeZendFunctionEntryEx("set_exception_handler", 0, ZifSetExceptionHandler, ArginfoSetExceptionHandler),
	types.MakeZendFunctionEntryEx("restore_exception_handler", 0, ZifRestoreExceptionHandler, nil),
	types.MakeZendFunctionEntryEx("get_declared_classes", 0, ZifGetDeclaredClasses, nil),
	types.MakeZendFunctionEntryEx("get_declared_traits", 0, ZifGetDeclaredTraits, nil),
	types.MakeZendFunctionEntryEx("get_declared_interfaces", 0, ZifGetDeclaredInterfaces, nil),
	types.MakeZendFunctionEntryEx("get_defined_functions", 0, ZifGetDefinedFunctions, ArginfoGetDefinedFunctions),
	types.MakeZendFunctionEntryEx("get_defined_vars", 0, ZifGetDefinedVars, nil),
	types.MakeZendFunctionEntryEx("create_function", ZEND_ACC_DEPRECATED, ZifCreateFunction, ArginfoCreateFunction),
	types.MakeZendFunctionEntryEx("get_resource_type", 0, ZifGetResourceType, ArginfoGetResourceType),
	types.MakeZendFunctionEntryEx("get_resources", 0, ZifGetResources, ArginfoGetResources),
	types.MakeZendFunctionEntryEx("get_loaded_extensions", 0, ZifGetLoadedExtensions, ArginfoGetLoadedExtensions),
	types.MakeZendFunctionEntryEx("extension_loaded", 0, ZifExtensionLoaded, ArginfoExtensionLoaded),
	types.MakeZendFunctionEntryEx("get_extension_funcs", 0, ZifGetExtensionFuncs, ArginfoExtensionLoaded),
	types.MakeZendFunctionEntryEx("get_defined_constants", 0, ZifGetDefinedConstants, ArginfoGetDefinedConstants),
	types.MakeZendFunctionEntryEx("debug_backtrace", 0, ZifDebugBacktrace, ArginfoDebugBacktrace),
	types.MakeZendFunctionEntryEx("debug_print_backtrace", 0, ZifDebugPrintBacktrace, ArginfoDebugPrintBacktrace),
	DefZifGcMemCaches,
	DefZifGcCollectCycles,
	DefZifGcEnabled,
	DefZifGcEnable,
	DefZifGcDisable,
	DefZifGcStatus,
}
var ZendBuiltinModule = MakeZendModuleEntry(b.SizeOf("zend_module_entry"), ZEND_MODULE_API_NO, 0, USING_ZTS, nil, nil, "Core", BuiltinFunctions, ZmStartupCore, nil, nil, nil, nil, ZEND_VERSION, 0, nil, nil, nil, nil, 0, 0, nil, 0, "API"+"ZEND_MODULE_API_NO"+ZEND_BUILD_TS)

const LAMBDA_TEMP_FUNCNAME = "__lambda_func"
