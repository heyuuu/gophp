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

/* {{{ arginfo */

var ArginfoZendVoid []ZendInternalArgInfo = []ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, ZEND_RETURN_VALUE, 0},
}
var ArginfoFuncGetArg []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg_num", 0, 0, 0}}
var ArginfoStrlen []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoStrcmp []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str1", 0, 0, 0}, {"str2", 0, 0, 0}}
var ArginfoStrncmp []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(3)), 0, 0, 0}, {"str1", 0, 0, 0}, {"str2", 0, 0, 0}, {"len", 0, 0, 0}}
var ArginfoEach []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arr", 0, 1, 0}}
var ArginfoErrorReporting []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"new_error_level", 0, 0, 0}}
var ArginfoDefine []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"constant_name", 0, 0, 0}, {"value", 0, 0, 0}, {"case_insensitive", 0, 0, 0}}
var ArginfoDefined []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"constant_name", 0, 0, 0}}
var ArginfoGetClass []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"object", 0, 0, 0}}
var ArginfoIsSubclassOf []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"object", 0, 0, 0}, {"class_name", 0, 0, 0}, {"allow_string", 0, 0, 0}}
var ArginfoGetClassVars []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"class_name", 0, 0, 0}}
var ArginfoGetObjectVars []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"obj", 0, 0, 0}}
var ArginfoGetMangledObjectVars []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"obj", 0, 0, 0}}
var ArginfoGetClassMethods []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"class", 0, 0, 0}}
var ArginfoMethodExists []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"object", 0, 0, 0}, {"method", 0, 0, 0}}
var ArginfoPropertyExists []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"object_or_class", 0, 0, 0}, {"property_name", 0, 0, 0}}
var ArginfoClassExists []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"classname", 0, 0, 0}, {"autoload", 0, 0, 0}}
var ArginfoTraitExists []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"traitname", 0, 0, 0}, {"autoload", 0, 0, 0}}
var ArginfoFunctionExists []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"function_name", 0, 0, 0}}
var ArginfoClassAlias []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"user_class_name", 0, 0, 0}, {"alias_name", 0, 0, 0}, {"autoload", 0, 0, 0}}
var ArginfoTriggerError []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"message", 0, 0, 0}, {"error_type", 0, 0, 0}}
var ArginfoSetErrorHandler []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"error_handler", 0, 0, 0}, {"error_types", 0, 0, 0}}
var ArginfoSetExceptionHandler []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"exception_handler", 0, 0, 0}}
var ArginfoGetDefinedFunctions []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"exclude_disabled", 0, 0, 0}}
var ArginfoCreateFunction []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"args", 0, 0, 0}, {"code", 0, 0, 0}}
var ArginfoGetResourceType []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"res", 0, 0, 0}}
var ArginfoGetResources []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"type", 0, 0, 0}}
var ArginfoGetLoadedExtensions []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"zend_extensions", 0, 0, 0}}
var ArginfoGetDefinedConstants []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"categorize", 0, 0, 0}}
var ArginfoDebugBacktrace []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"options", 0, 0, 0}, {"limit", 0, 0, 0}}
var ArginfoDebugPrintBacktrace []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"options", 0, 0, 0}, {"limit", 0, 0, 0}}
var ArginfoExtensionLoaded []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"extension_name", 0, 0, 0}}
var BuiltinFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"zend_version",
		ZifZendVersion,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"func_num_args",
		ZifFuncNumArgs,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"func_get_arg",
		ZifFuncGetArg,
		ArginfoFuncGetArg,
		uint32(b.SizeOf("arginfo_func_get_arg")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"func_get_args",
		ZifFuncGetArgs,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strlen",
		ZifStrlen,
		ArginfoStrlen,
		uint32(b.SizeOf("arginfo_strlen")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strcmp",
		ZifStrcmp,
		ArginfoStrcmp,
		uint32(b.SizeOf("arginfo_strcmp")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strncmp",
		ZifStrncmp,
		ArginfoStrncmp,
		uint32(b.SizeOf("arginfo_strncmp")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strcasecmp",
		ZifStrcasecmp,
		ArginfoStrcmp,
		uint32(b.SizeOf("arginfo_strcmp")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strncasecmp",
		ZifStrncasecmp,
		ArginfoStrncmp,
		uint32(b.SizeOf("arginfo_strncmp")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"each",
		ZifEach,
		ArginfoEach,
		uint32(b.SizeOf("arginfo_each")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"error_reporting",
		ZifErrorReporting,
		ArginfoErrorReporting,
		uint32(b.SizeOf("arginfo_error_reporting")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"define",
		ZifDefine,
		ArginfoDefine,
		uint32(b.SizeOf("arginfo_define")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"defined",
		ZifDefined,
		ArginfoDefined,
		uint32(b.SizeOf("arginfo_defined")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_class",
		ZifGetClass,
		ArginfoGetClass,
		uint32(b.SizeOf("arginfo_get_class")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_called_class",
		ZifGetCalledClass,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_parent_class",
		ZifGetParentClass,
		ArginfoGetClass,
		uint32(b.SizeOf("arginfo_get_class")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"method_exists",
		ZifMethodExists,
		ArginfoMethodExists,
		uint32(b.SizeOf("arginfo_method_exists")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"property_exists",
		ZifPropertyExists,
		ArginfoPropertyExists,
		uint32(b.SizeOf("arginfo_property_exists")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"class_exists",
		ZifClassExists,
		ArginfoClassExists,
		uint32(b.SizeOf("arginfo_class_exists")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"interface_exists",
		ZifInterfaceExists,
		ArginfoClassExists,
		uint32(b.SizeOf("arginfo_class_exists")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"trait_exists",
		ZifTraitExists,
		ArginfoTraitExists,
		uint32(b.SizeOf("arginfo_trait_exists")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"function_exists",
		ZifFunctionExists,
		ArginfoFunctionExists,
		uint32(b.SizeOf("arginfo_function_exists")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"class_alias",
		ZifClassAlias,
		ArginfoClassAlias,
		uint32(b.SizeOf("arginfo_class_alias")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_included_files",
		ZifGetIncludedFiles,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_required_files",
		ZifGetIncludedFiles,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_subclass_of",
		ZifIsSubclassOf,
		ArginfoIsSubclassOf,
		uint32(b.SizeOf("arginfo_is_subclass_of")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_a",
		ZifIsA,
		ArginfoIsSubclassOf,
		uint32(b.SizeOf("arginfo_is_subclass_of")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_class_vars",
		ZifGetClassVars,
		ArginfoGetClassVars,
		uint32(b.SizeOf("arginfo_get_class_vars")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_object_vars",
		ZifGetObjectVars,
		ArginfoGetObjectVars,
		uint32(b.SizeOf("arginfo_get_object_vars")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_mangled_object_vars",
		ZifGetMangledObjectVars,
		ArginfoGetMangledObjectVars,
		uint32(b.SizeOf("arginfo_get_mangled_object_vars")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_class_methods",
		ZifGetClassMethods,
		ArginfoGetClassMethods,
		uint32(b.SizeOf("arginfo_get_class_methods")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"trigger_error",
		ZifTriggerError,
		ArginfoTriggerError,
		uint32(b.SizeOf("arginfo_trigger_error")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"user_error",
		ZifTriggerError,
		ArginfoTriggerError,
		uint32(b.SizeOf("arginfo_trigger_error")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"set_error_handler",
		ZifSetErrorHandler,
		ArginfoSetErrorHandler,
		uint32(b.SizeOf("arginfo_set_error_handler")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"restore_error_handler",
		ZifRestoreErrorHandler,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"set_exception_handler",
		ZifSetExceptionHandler,
		ArginfoSetExceptionHandler,
		uint32(b.SizeOf("arginfo_set_exception_handler")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"restore_exception_handler",
		ZifRestoreExceptionHandler,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_declared_classes",
		ZifGetDeclaredClasses,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_declared_traits",
		ZifGetDeclaredTraits,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_declared_interfaces",
		ZifGetDeclaredInterfaces,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_defined_functions",
		ZifGetDefinedFunctions,
		ArginfoGetDefinedFunctions,
		uint32(b.SizeOf("arginfo_get_defined_functions")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_defined_vars",
		ZifGetDefinedVars,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"create_function",
		ZifCreateFunction,
		ArginfoCreateFunction,
		uint32(b.SizeOf("arginfo_create_function")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_DEPRECATED,
	},
	{
		"get_resource_type",
		ZifGetResourceType,
		ArginfoGetResourceType,
		uint32(b.SizeOf("arginfo_get_resource_type")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_resources",
		ZifGetResources,
		ArginfoGetResources,
		uint32(b.SizeOf("arginfo_get_resources")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_loaded_extensions",
		ZifGetLoadedExtensions,
		ArginfoGetLoadedExtensions,
		uint32(b.SizeOf("arginfo_get_loaded_extensions")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"extension_loaded",
		ZifExtensionLoaded,
		ArginfoExtensionLoaded,
		uint32(b.SizeOf("arginfo_extension_loaded")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_extension_funcs",
		ZifGetExtensionFuncs,
		ArginfoExtensionLoaded,
		uint32(b.SizeOf("arginfo_extension_loaded")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_defined_constants",
		ZifGetDefinedConstants,
		ArginfoGetDefinedConstants,
		uint32(b.SizeOf("arginfo_get_defined_constants")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"debug_backtrace",
		ZifDebugBacktrace,
		ArginfoDebugBacktrace,
		uint32(b.SizeOf("arginfo_debug_backtrace")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"debug_print_backtrace",
		ZifDebugPrintBacktrace,
		ArginfoDebugPrintBacktrace,
		uint32(b.SizeOf("arginfo_debug_print_backtrace")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gc_mem_caches",
		ZifGcMemCaches,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gc_collect_cycles",
		ZifGcCollectCycles,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gc_enabled",
		ZifGcEnabled,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gc_enable",
		ZifGcEnable,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gc_disable",
		ZifGcDisable,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gc_status",
		ZifGcStatus,
		ArginfoZendVoid,
		uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}
var ZendBuiltinModule ZendModuleEntry = ZendModuleEntry{
	b.SizeOf("zend_module_entry"),
	ZEND_MODULE_API_NO,
	0,
	USING_ZTS,
	nil,
	nil,
	"Core",
	BuiltinFunctions,
	ZEND_MINIT(core),
	nil,
	nil,
	nil,
	nil,
	ZEND_VERSION,
	0,
	nil,
	nil,
	nil,
	nil,
	0,
	0,
	nil,
	0,
	"API" + "ZEND_MODULE_API_NO" + ZEND_BUILD_TS,
}

/* {{{ */

/* {{{ */

const LAMBDA_TEMP_FUNCNAME = "__lambda_func"

/* {{{ proto string create_function(string args, string code)
   Creates an anonymous function, and returns its name (funny, eh?) */

/* {{{ proto string get_resource_type(resource res)
   Get the resource type name for a given resource */

/* {{{ */
