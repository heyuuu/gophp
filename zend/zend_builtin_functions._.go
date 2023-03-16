// <<generate>>

package zend

import (
	b "sik/builtin"
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
var BuiltinFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	ZifZendVersionDef,
	MakeZendFunctionEntryEx("func_num_args", 0, ZifFuncNumArgs, ArginfoZendVoid),
	MakeZendFunctionEntryEx("func_get_arg", 0, ZifFuncGetArg, ArginfoFuncGetArg),
	MakeZendFunctionEntryEx("func_get_args", 0, ZifFuncGetArgs, ArginfoZendVoid),
	MakeZendFunctionEntryEx("strlen", 0, ZifStrlen, ArginfoStrlen),
	MakeZendFunctionEntryEx("strcmp", 0, ZifStrcmp, ArginfoStrcmp),
	MakeZendFunctionEntryEx("strncmp", 0, ZifStrncmp, ArginfoStrncmp),
	MakeZendFunctionEntryEx("strcasecmp", 0, ZifStrcasecmp, ArginfoStrcmp),
	MakeZendFunctionEntryEx("strncasecmp", 0, ZifStrncasecmp, ArginfoStrncmp),
	MakeZendFunctionEntryEx("each", 0, ZifEach, ArginfoEach),
	MakeZendFunctionEntryEx("error_reporting", 0, ZifErrorReporting, ArginfoErrorReporting),
	MakeZendFunctionEntryEx("define", 0, ZifDefine, ArginfoDefine),
	MakeZendFunctionEntryEx("defined", 0, ZifDefined, ArginfoDefined),
	MakeZendFunctionEntryEx("get_class", 0, ZifGetClass, ArginfoGetClass),
	MakeZendFunctionEntryEx("get_called_class", 0, ZifGetCalledClass, ArginfoZendVoid),
	MakeZendFunctionEntryEx("get_parent_class", 0, ZifGetParentClass, ArginfoGetClass),
	MakeZendFunctionEntryEx("method_exists", 0, ZifMethodExists, ArginfoMethodExists),
	MakeZendFunctionEntryEx("property_exists", 0, ZifPropertyExists, ArginfoPropertyExists),
	MakeZendFunctionEntryEx("class_exists", 0, ZifClassExists, ArginfoClassExists),
	MakeZendFunctionEntryEx("interface_exists", 0, ZifInterfaceExists, ArginfoClassExists),
	MakeZendFunctionEntryEx("trait_exists", 0, ZifTraitExists, ArginfoTraitExists),
	MakeZendFunctionEntryEx("function_exists", 0, ZifFunctionExists, ArginfoFunctionExists),
	MakeZendFunctionEntryEx("class_alias", 0, ZifClassAlias, ArginfoClassAlias),
	MakeZendFunctionEntryEx("get_included_files", 0, ZifGetIncludedFiles, ArginfoZendVoid),
	MakeZendFunctionEntryEx("get_required_files", 0, ZifGetIncludedFiles, ArginfoZendVoid),
	MakeZendFunctionEntryEx("is_subclass_of", 0, ZifIsSubclassOf, ArginfoIsSubclassOf),
	MakeZendFunctionEntryEx("is_a", 0, ZifIsA, ArginfoIsSubclassOf),
	MakeZendFunctionEntryEx("get_class_vars", 0, ZifGetClassVars, ArginfoGetClassVars),
	MakeZendFunctionEntryEx("get_object_vars", 0, ZifGetObjectVars, ArginfoGetObjectVars),
	MakeZendFunctionEntryEx("get_mangled_object_vars", 0, ZifGetMangledObjectVars, ArginfoGetMangledObjectVars),
	MakeZendFunctionEntryEx("get_class_methods", 0, ZifGetClassMethods, ArginfoGetClassMethods),
	MakeZendFunctionEntryEx("trigger_error", 0, ZifTriggerError, ArginfoTriggerError),
	MakeZendFunctionEntryEx("user_error", 0, ZifTriggerError, ArginfoTriggerError),
	MakeZendFunctionEntryEx("set_error_handler", 0, ZifSetErrorHandler, ArginfoSetErrorHandler),
	MakeZendFunctionEntryEx("restore_error_handler", 0, ZifRestoreErrorHandler, ArginfoZendVoid),
	MakeZendFunctionEntryEx("set_exception_handler", 0, ZifSetExceptionHandler, ArginfoSetExceptionHandler),
	MakeZendFunctionEntryEx("restore_exception_handler", 0, ZifRestoreExceptionHandler, ArginfoZendVoid),
	MakeZendFunctionEntryEx("get_declared_classes", 0, ZifGetDeclaredClasses, ArginfoZendVoid),
	MakeZendFunctionEntryEx("get_declared_traits", 0, ZifGetDeclaredTraits, ArginfoZendVoid),
	MakeZendFunctionEntryEx("get_declared_interfaces", 0, ZifGetDeclaredInterfaces, ArginfoZendVoid),
	MakeZendFunctionEntryEx("get_defined_functions", 0, ZifGetDefinedFunctions, ArginfoGetDefinedFunctions),
	MakeZendFunctionEntryEx("get_defined_vars", 0, ZifGetDefinedVars, ArginfoZendVoid),
	MakeZendFunctionEntryEx("create_function", ZEND_ACC_DEPRECATED, ZifCreateFunction, ArginfoCreateFunction),
	MakeZendFunctionEntryEx("get_resource_type", 0, ZifGetResourceType, ArginfoGetResourceType),
	MakeZendFunctionEntryEx("get_resources", 0, ZifGetResources, ArginfoGetResources),
	MakeZendFunctionEntryEx("get_loaded_extensions", 0, ZifGetLoadedExtensions, ArginfoGetLoadedExtensions),
	MakeZendFunctionEntryEx("extension_loaded", 0, ZifExtensionLoaded, ArginfoExtensionLoaded),
	MakeZendFunctionEntryEx("get_extension_funcs", 0, ZifGetExtensionFuncs, ArginfoExtensionLoaded),
	MakeZendFunctionEntryEx("get_defined_constants", 0, ZifGetDefinedConstants, ArginfoGetDefinedConstants),
	MakeZendFunctionEntryEx("debug_backtrace", 0, ZifDebugBacktrace, ArginfoDebugBacktrace),
	MakeZendFunctionEntryEx("debug_print_backtrace", 0, ZifDebugPrintBacktrace, ArginfoDebugPrintBacktrace),
	MakeZendFunctionEntryEx("gc_mem_caches", 0, ZifGcMemCaches, ArginfoZendVoid),
	MakeZendFunctionEntryEx("gc_collect_cycles", 0, ZifGcCollectCycles, ArginfoZendVoid),
	MakeZendFunctionEntryEx("gc_enabled", 0, ZifGcEnabled, ArginfoZendVoid),
	MakeZendFunctionEntryEx("gc_enable", 0, ZifGcEnable, ArginfoZendVoid),
	MakeZendFunctionEntryEx("gc_disable", 0, ZifGcDisable, ArginfoZendVoid),
	MakeZendFunctionEntryEx("gc_status", 0, ZifGcStatus, ArginfoZendVoid),
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
