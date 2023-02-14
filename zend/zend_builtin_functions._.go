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
	MakeArgInfo("arr",ArgInfoByRef(1)),
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
	MakeZendFunctionEntry("zend_version", ZifZendVersion, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("func_num_args", ZifFuncNumArgs, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("func_get_arg", ZifFuncGetArg, ArginfoFuncGetArg, uint32(b.SizeOf("arginfo_func_get_arg")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("func_get_args", ZifFuncGetArgs, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("strlen", ZifStrlen, ArginfoStrlen, uint32(b.SizeOf("arginfo_strlen")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("strcmp", ZifStrcmp, ArginfoStrcmp, uint32(b.SizeOf("arginfo_strcmp")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("strncmp", ZifStrncmp, ArginfoStrncmp, uint32(b.SizeOf("arginfo_strncmp")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("strcasecmp", ZifStrcasecmp, ArginfoStrcmp, uint32(b.SizeOf("arginfo_strcmp")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("strncasecmp", ZifStrncasecmp, ArginfoStrncmp, uint32(b.SizeOf("arginfo_strncmp")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("each", ZifEach, ArginfoEach, uint32(b.SizeOf("arginfo_each")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("error_reporting", ZifErrorReporting, ArginfoErrorReporting, uint32(b.SizeOf("arginfo_error_reporting")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("define", ZifDefine, ArginfoDefine, uint32(b.SizeOf("arginfo_define")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("defined", ZifDefined, ArginfoDefined, uint32(b.SizeOf("arginfo_defined")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_class", ZifGetClass, ArginfoGetClass, uint32(b.SizeOf("arginfo_get_class")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_called_class", ZifGetCalledClass, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_parent_class", ZifGetParentClass, ArginfoGetClass, uint32(b.SizeOf("arginfo_get_class")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("method_exists", ZifMethodExists, ArginfoMethodExists, uint32(b.SizeOf("arginfo_method_exists")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("property_exists", ZifPropertyExists, ArginfoPropertyExists, uint32(b.SizeOf("arginfo_property_exists")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("class_exists", ZifClassExists, ArginfoClassExists, uint32(b.SizeOf("arginfo_class_exists")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("interface_exists", ZifInterfaceExists, ArginfoClassExists, uint32(b.SizeOf("arginfo_class_exists")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("trait_exists", ZifTraitExists, ArginfoTraitExists, uint32(b.SizeOf("arginfo_trait_exists")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("function_exists", ZifFunctionExists, ArginfoFunctionExists, uint32(b.SizeOf("arginfo_function_exists")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("class_alias", ZifClassAlias, ArginfoClassAlias, uint32(b.SizeOf("arginfo_class_alias")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_included_files", ZifGetIncludedFiles, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_required_files", ZifGetIncludedFiles, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("is_subclass_of", ZifIsSubclassOf, ArginfoIsSubclassOf, uint32(b.SizeOf("arginfo_is_subclass_of")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("is_a", ZifIsA, ArginfoIsSubclassOf, uint32(b.SizeOf("arginfo_is_subclass_of")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_class_vars", ZifGetClassVars, ArginfoGetClassVars, uint32(b.SizeOf("arginfo_get_class_vars")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_object_vars", ZifGetObjectVars, ArginfoGetObjectVars, uint32(b.SizeOf("arginfo_get_object_vars")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_mangled_object_vars", ZifGetMangledObjectVars, ArginfoGetMangledObjectVars, uint32(b.SizeOf("arginfo_get_mangled_object_vars")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_class_methods", ZifGetClassMethods, ArginfoGetClassMethods, uint32(b.SizeOf("arginfo_get_class_methods")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("trigger_error", ZifTriggerError, ArginfoTriggerError, uint32(b.SizeOf("arginfo_trigger_error")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("user_error", ZifTriggerError, ArginfoTriggerError, uint32(b.SizeOf("arginfo_trigger_error")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("set_error_handler", ZifSetErrorHandler, ArginfoSetErrorHandler, uint32(b.SizeOf("arginfo_set_error_handler")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("restore_error_handler", ZifRestoreErrorHandler, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("set_exception_handler", ZifSetExceptionHandler, ArginfoSetExceptionHandler, uint32(b.SizeOf("arginfo_set_exception_handler")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("restore_exception_handler", ZifRestoreExceptionHandler, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_declared_classes", ZifGetDeclaredClasses, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_declared_traits", ZifGetDeclaredTraits, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_declared_interfaces", ZifGetDeclaredInterfaces, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_defined_functions", ZifGetDefinedFunctions, ArginfoGetDefinedFunctions, uint32(b.SizeOf("arginfo_get_defined_functions")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_defined_vars", ZifGetDefinedVars, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("create_function", ZifCreateFunction, ArginfoCreateFunction, uint32(b.SizeOf("arginfo_create_function")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_DEPRECATED),
	MakeZendFunctionEntry("get_resource_type", ZifGetResourceType, ArginfoGetResourceType, uint32(b.SizeOf("arginfo_get_resource_type")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_resources", ZifGetResources, ArginfoGetResources, uint32(b.SizeOf("arginfo_get_resources")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_loaded_extensions", ZifGetLoadedExtensions, ArginfoGetLoadedExtensions, uint32(b.SizeOf("arginfo_get_loaded_extensions")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("extension_loaded", ZifExtensionLoaded, ArginfoExtensionLoaded, uint32(b.SizeOf("arginfo_extension_loaded")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_extension_funcs", ZifGetExtensionFuncs, ArginfoExtensionLoaded, uint32(b.SizeOf("arginfo_extension_loaded")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("get_defined_constants", ZifGetDefinedConstants, ArginfoGetDefinedConstants, uint32(b.SizeOf("arginfo_get_defined_constants")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("debug_backtrace", ZifDebugBacktrace, ArginfoDebugBacktrace, uint32(b.SizeOf("arginfo_debug_backtrace")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("debug_print_backtrace", ZifDebugPrintBacktrace, ArginfoDebugPrintBacktrace, uint32(b.SizeOf("arginfo_debug_print_backtrace")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("gc_mem_caches", ZifGcMemCaches, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("gc_collect_cycles", ZifGcCollectCycles, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("gc_enabled", ZifGcEnabled, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("gc_enable", ZifGcEnable, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("gc_disable", ZifGcDisable, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("gc_status", ZifGcStatus, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry(nil, nil, nil, 0, 0),
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
