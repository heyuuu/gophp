// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/types"
)

// Source: <Zend/zend_builtin_functions.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// Source: <Zend/zend_builtin_functions.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

var ArginfoZendVoid []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(-1),
}
var ArginfoFuncGetArg []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("arg_num"),
}
var ArginfoStrlen []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("str"),
}
var ArginfoStrcmp []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("str1"),
	MakeArgInfo("str2"),
}
var ArginfoStrncmp []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(3),
	MakeArgInfo("str1"),
	MakeArgInfo("str2"),
	MakeArgInfo("len"),
}
var ArginfoEach []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("arr", ArgInfoByRef(1)),
}
var ArginfoErrorReporting []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("new_error_level"),
}
var ArginfoDefine []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("constant_name"),
	MakeArgInfo("value"),
	MakeArgInfo("case_insensitive"),
}
var ArginfoDefined []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("constant_name"),
}
var ArginfoGetClass []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("object"),
}
var ArginfoIsSubclassOf []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("object"),
	MakeArgInfo("class_name"),
	MakeArgInfo("allow_string"),
}
var ArginfoGetClassVars []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("class_name"),
}
var ArginfoGetObjectVars []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("obj"),
}
var ArginfoGetMangledObjectVars []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("obj"),
}
var ArginfoGetClassMethods []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("class"),
}
var ArginfoMethodExists []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("object"),
	MakeArgInfo("method"),
}
var ArginfoPropertyExists []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("object_or_class"),
	MakeArgInfo("property_name"),
}
var ArginfoClassExists []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("classname"),
	MakeArgInfo("autoload"),
}
var ArginfoTraitExists []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("traitname"),
	MakeArgInfo("autoload"),
}
var ArginfoFunctionExists []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("function_name"),
}
var ArginfoClassAlias []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("user_class_name"),
	MakeArgInfo("alias_name"),
	MakeArgInfo("autoload"),
}
var ArginfoTriggerError []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("message"),
	MakeArgInfo("error_type"),
}
var ArginfoSetErrorHandler []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("error_handler"),
	MakeArgInfo("error_types"),
}
var ArginfoSetExceptionHandler []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("exception_handler"),
}
var ArginfoGetDefinedFunctions []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("exclude_disabled"),
}
var ArginfoCreateFunction []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("args"),
	MakeArgInfo("code"),
}
var ArginfoGetResourceType []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("res"),
}
var ArginfoGetResources []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("type"),
}
var ArginfoGetLoadedExtensions []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("zend_extensions"),
}
var ArginfoGetDefinedConstants []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("categorize"),
}
var ArginfoDebugBacktrace []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("options"),
	MakeArgInfo("limit"),
}
var ArginfoDebugPrintBacktrace []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(0),
	MakeArgInfo("options"),
	MakeArgInfo("limit"),
}
var ArginfoExtensionLoaded []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("extension_name"),
}
var BuiltinFunctions []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	ZifZendVersionDef,
	types.MakeZendFunctionEntryEx("func_num_args", 0, ZifFuncNumArgs, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("func_get_arg", 0, ZifFuncGetArg, ArginfoFuncGetArg),
	types.MakeZendFunctionEntryEx("func_get_args", 0, ZifFuncGetArgs, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("strlen", 0, ZifStrlen, ArginfoStrlen),
	types.MakeZendFunctionEntryEx("strcmp", 0, ZifStrcmp, ArginfoStrcmp),
	types.MakeZendFunctionEntryEx("strncmp", 0, ZifStrncmp, ArginfoStrncmp),
	types.MakeZendFunctionEntryEx("strcasecmp", 0, ZifStrcasecmp, ArginfoStrcmp),
	types.MakeZendFunctionEntryEx("strncasecmp", 0, ZifStrncasecmp, ArginfoStrncmp),
	types.MakeZendFunctionEntryEx("each", 0, ZifEach, ArginfoEach),
	types.MakeZendFunctionEntryEx("error_reporting", 0, ZifErrorReporting, ArginfoErrorReporting),
	types.MakeZendFunctionEntryEx("define", 0, ZifDefine, ArginfoDefine),
	types.MakeZendFunctionEntryEx("defined", 0, ZifDefined, ArginfoDefined),
	types.MakeZendFunctionEntryEx("get_class", 0, ZifGetClass, ArginfoGetClass),
	types.MakeZendFunctionEntryEx("get_called_class", 0, ZifGetCalledClass, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("get_parent_class", 0, ZifGetParentClass, ArginfoGetClass),
	types.MakeZendFunctionEntryEx("method_exists", 0, ZifMethodExists, ArginfoMethodExists),
	types.MakeZendFunctionEntryEx("property_exists", 0, ZifPropertyExists, ArginfoPropertyExists),
	types.MakeZendFunctionEntryEx("class_exists", 0, ZifClassExists, ArginfoClassExists),
	types.MakeZendFunctionEntryEx("interface_exists", 0, ZifInterfaceExists, ArginfoClassExists),
	types.MakeZendFunctionEntryEx("trait_exists", 0, ZifTraitExists, ArginfoTraitExists),
	types.MakeZendFunctionEntryEx("function_exists", 0, ZifFunctionExists, ArginfoFunctionExists),
	types.MakeZendFunctionEntryEx("class_alias", 0, ZifClassAlias, ArginfoClassAlias),
	types.MakeZendFunctionEntryEx("get_included_files", 0, ZifGetIncludedFiles, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("get_required_files", 0, ZifGetIncludedFiles, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("is_subclass_of", 0, ZifIsSubclassOf, ArginfoIsSubclassOf),
	types.MakeZendFunctionEntryEx("is_a", 0, ZifIsA, ArginfoIsSubclassOf),
	types.MakeZendFunctionEntryEx("get_class_vars", 0, ZifGetClassVars, ArginfoGetClassVars),
	types.MakeZendFunctionEntryEx("get_object_vars", 0, ZifGetObjectVars, ArginfoGetObjectVars),
	types.MakeZendFunctionEntryEx("get_mangled_object_vars", 0, ZifGetMangledObjectVars, ArginfoGetMangledObjectVars),
	types.MakeZendFunctionEntryEx("get_class_methods", 0, ZifGetClassMethods, ArginfoGetClassMethods),
	types.MakeZendFunctionEntryEx("trigger_error", 0, ZifTriggerError, ArginfoTriggerError),
	types.MakeZendFunctionEntryEx("user_error", 0, ZifTriggerError, ArginfoTriggerError),
	types.MakeZendFunctionEntryEx("set_error_handler", 0, ZifSetErrorHandler, ArginfoSetErrorHandler),
	types.MakeZendFunctionEntryEx("restore_error_handler", 0, ZifRestoreErrorHandler, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("set_exception_handler", 0, ZifSetExceptionHandler, ArginfoSetExceptionHandler),
	types.MakeZendFunctionEntryEx("restore_exception_handler", 0, ZifRestoreExceptionHandler, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("get_declared_classes", 0, ZifGetDeclaredClasses, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("get_declared_traits", 0, ZifGetDeclaredTraits, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("get_declared_interfaces", 0, ZifGetDeclaredInterfaces, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("get_defined_functions", 0, ZifGetDefinedFunctions, ArginfoGetDefinedFunctions),
	types.MakeZendFunctionEntryEx("get_defined_vars", 0, ZifGetDefinedVars, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("create_function", ZEND_ACC_DEPRECATED, ZifCreateFunction, ArginfoCreateFunction),
	types.MakeZendFunctionEntryEx("get_resource_type", 0, ZifGetResourceType, ArginfoGetResourceType),
	types.MakeZendFunctionEntryEx("get_resources", 0, ZifGetResources, ArginfoGetResources),
	types.MakeZendFunctionEntryEx("get_loaded_extensions", 0, ZifGetLoadedExtensions, ArginfoGetLoadedExtensions),
	types.MakeZendFunctionEntryEx("extension_loaded", 0, ZifExtensionLoaded, ArginfoExtensionLoaded),
	types.MakeZendFunctionEntryEx("get_extension_funcs", 0, ZifGetExtensionFuncs, ArginfoExtensionLoaded),
	types.MakeZendFunctionEntryEx("get_defined_constants", 0, ZifGetDefinedConstants, ArginfoGetDefinedConstants),
	types.MakeZendFunctionEntryEx("debug_backtrace", 0, ZifDebugBacktrace, ArginfoDebugBacktrace),
	types.MakeZendFunctionEntryEx("debug_print_backtrace", 0, ZifDebugPrintBacktrace, ArginfoDebugPrintBacktrace),
	types.MakeZendFunctionEntryEx("gc_mem_caches", 0, ZifGcMemCaches, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("gc_collect_cycles", 0, ZifGcCollectCycles, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("gc_enabled", 0, ZifGcEnabled, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("gc_enable", 0, ZifGcEnable, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("gc_disable", 0, ZifGcDisable, ArginfoZendVoid),
	types.MakeZendFunctionEntryEx("gc_status", 0, ZifGcStatus, ArginfoZendVoid),
}
var ZendBuiltinModule ZendModuleEntry = MakeZendModuleEntry(b.SizeOf("zend_module_entry"), ZEND_MODULE_API_NO, 0, USING_ZTS, nil, nil, "Core", BuiltinFunctions, ZmStartupCore, nil, nil, nil, nil, ZEND_VERSION, 0, nil, nil, nil, nil, 0, 0, nil, 0, "API"+"ZEND_MODULE_API_NO"+ZEND_BUILD_TS)

/* {{{ */

/* {{{ */

const LAMBDA_TEMP_FUNCNAME = "__lambda_func"

/* {{{ proto string create_function(string args, string code)
   Creates an anonymous function, and returns its name (funny, eh?) */

/* {{{ proto string get_resource_type(resource res)
   Get the resource type name for a given resource */

/* {{{ */
