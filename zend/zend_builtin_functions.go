// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
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

// #define ZEND_BUILTIN_FUNCTIONS_H

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

// # include "zend.h"

// # include "zend_API.h"

// # include "zend_gc.h"

// # include "zend_builtin_functions.h"

// # include "zend_constants.h"

// # include "zend_ini.h"

// # include "zend_exceptions.h"

// # include "zend_extensions.h"

// # include "zend_closures.h"

// # include "zend_generators.h"

/* {{{ arginfo */

var ArginfoZendVoid []ZendInternalArgInfo = []ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
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

/* }}} */

var BuiltinFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"zend_version",
		ZifZendVersion,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"func_num_args",
		ZifFuncNumArgs,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"func_get_arg",
		ZifFuncGetArg,
		ArginfoFuncGetArg,
		uint32(g.SizeOf("arginfo_func_get_arg")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"func_get_args",
		ZifFuncGetArgs,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strlen",
		ZifStrlen,
		ArginfoStrlen,
		uint32(g.SizeOf("arginfo_strlen")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strcmp",
		ZifStrcmp,
		ArginfoStrcmp,
		uint32(g.SizeOf("arginfo_strcmp")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strncmp",
		ZifStrncmp,
		ArginfoStrncmp,
		uint32(g.SizeOf("arginfo_strncmp")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strcasecmp",
		ZifStrcasecmp,
		ArginfoStrcmp,
		uint32(g.SizeOf("arginfo_strcmp")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strncasecmp",
		ZifStrncasecmp,
		ArginfoStrncmp,
		uint32(g.SizeOf("arginfo_strncmp")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"each",
		ZifEach,
		ArginfoEach,
		uint32(g.SizeOf("arginfo_each")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"error_reporting",
		ZifErrorReporting,
		ArginfoErrorReporting,
		uint32(g.SizeOf("arginfo_error_reporting")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"define",
		ZifDefine,
		ArginfoDefine,
		uint32(g.SizeOf("arginfo_define")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"defined",
		ZifDefined,
		ArginfoDefined,
		uint32(g.SizeOf("arginfo_defined")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_class",
		ZifGetClass,
		ArginfoGetClass,
		uint32(g.SizeOf("arginfo_get_class")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_called_class",
		ZifGetCalledClass,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_parent_class",
		ZifGetParentClass,
		ArginfoGetClass,
		uint32(g.SizeOf("arginfo_get_class")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"method_exists",
		ZifMethodExists,
		ArginfoMethodExists,
		uint32(g.SizeOf("arginfo_method_exists")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"property_exists",
		ZifPropertyExists,
		ArginfoPropertyExists,
		uint32(g.SizeOf("arginfo_property_exists")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"class_exists",
		ZifClassExists,
		ArginfoClassExists,
		uint32(g.SizeOf("arginfo_class_exists")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"interface_exists",
		ZifInterfaceExists,
		ArginfoClassExists,
		uint32(g.SizeOf("arginfo_class_exists")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"trait_exists",
		ZifTraitExists,
		ArginfoTraitExists,
		uint32(g.SizeOf("arginfo_trait_exists")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"function_exists",
		ZifFunctionExists,
		ArginfoFunctionExists,
		uint32(g.SizeOf("arginfo_function_exists")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"class_alias",
		ZifClassAlias,
		ArginfoClassAlias,
		uint32(g.SizeOf("arginfo_class_alias")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_included_files",
		ZifGetIncludedFiles,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_required_files",
		ZifGetIncludedFiles,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_subclass_of",
		ZifIsSubclassOf,
		ArginfoIsSubclassOf,
		uint32(g.SizeOf("arginfo_is_subclass_of")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_a",
		ZifIsA,
		ArginfoIsSubclassOf,
		uint32(g.SizeOf("arginfo_is_subclass_of")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_class_vars",
		ZifGetClassVars,
		ArginfoGetClassVars,
		uint32(g.SizeOf("arginfo_get_class_vars")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_object_vars",
		ZifGetObjectVars,
		ArginfoGetObjectVars,
		uint32(g.SizeOf("arginfo_get_object_vars")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_mangled_object_vars",
		ZifGetMangledObjectVars,
		ArginfoGetMangledObjectVars,
		uint32(g.SizeOf("arginfo_get_mangled_object_vars")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_class_methods",
		ZifGetClassMethods,
		ArginfoGetClassMethods,
		uint32(g.SizeOf("arginfo_get_class_methods")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"trigger_error",
		ZifTriggerError,
		ArginfoTriggerError,
		uint32(g.SizeOf("arginfo_trigger_error")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"user_error",
		ZifTriggerError,
		ArginfoTriggerError,
		uint32(g.SizeOf("arginfo_trigger_error")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"set_error_handler",
		ZifSetErrorHandler,
		ArginfoSetErrorHandler,
		uint32(g.SizeOf("arginfo_set_error_handler")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"restore_error_handler",
		ZifRestoreErrorHandler,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"set_exception_handler",
		ZifSetExceptionHandler,
		ArginfoSetExceptionHandler,
		uint32(g.SizeOf("arginfo_set_exception_handler")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"restore_exception_handler",
		ZifRestoreExceptionHandler,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_declared_classes",
		ZifGetDeclaredClasses,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_declared_traits",
		ZifGetDeclaredTraits,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_declared_interfaces",
		ZifGetDeclaredInterfaces,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_defined_functions",
		ZifGetDefinedFunctions,
		ArginfoGetDefinedFunctions,
		uint32(g.SizeOf("arginfo_get_defined_functions")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_defined_vars",
		ZifGetDefinedVars,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"create_function",
		ZifCreateFunction,
		ArginfoCreateFunction,
		uint32(g.SizeOf("arginfo_create_function")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 11,
	},
	{
		"get_resource_type",
		ZifGetResourceType,
		ArginfoGetResourceType,
		uint32(g.SizeOf("arginfo_get_resource_type")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_resources",
		ZifGetResources,
		ArginfoGetResources,
		uint32(g.SizeOf("arginfo_get_resources")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_loaded_extensions",
		ZifGetLoadedExtensions,
		ArginfoGetLoadedExtensions,
		uint32(g.SizeOf("arginfo_get_loaded_extensions")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"extension_loaded",
		ZifExtensionLoaded,
		ArginfoExtensionLoaded,
		uint32(g.SizeOf("arginfo_extension_loaded")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_extension_funcs",
		ZifGetExtensionFuncs,
		ArginfoExtensionLoaded,
		uint32(g.SizeOf("arginfo_extension_loaded")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_defined_constants",
		ZifGetDefinedConstants,
		ArginfoGetDefinedConstants,
		uint32(g.SizeOf("arginfo_get_defined_constants")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"debug_backtrace",
		ZifDebugBacktrace,
		ArginfoDebugBacktrace,
		uint32(g.SizeOf("arginfo_debug_backtrace")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"debug_print_backtrace",
		ZifDebugPrintBacktrace,
		ArginfoDebugPrintBacktrace,
		uint32(g.SizeOf("arginfo_debug_print_backtrace")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gc_mem_caches",
		ZifGcMemCaches,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gc_collect_cycles",
		ZifGcCollectCycles,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gc_enabled",
		ZifGcEnabled,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gc_enable",
		ZifGcEnable,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gc_disable",
		ZifGcDisable,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gc_status",
		ZifGcStatus,
		ArginfoZendVoid,
		uint32(g.SizeOf("arginfo_zend__void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

func ZmStartupCore(type_ int, module_number int) int {
	var class_entry ZendClassEntry
	memset(&class_entry, 0, g.SizeOf("zend_class_entry"))
	class_entry.SetName(ZendStringInitInterned("stdClass", g.SizeOf("\"stdClass\"")-1, 1))
	class_entry.SetBuiltinFunctions(nil)
	ZendStandardClassDef = ZendRegisterInternalClass(&class_entry)
	ZendRegisterDefaultClasses()
	return SUCCESS
}

/* }}} */

var ZendBuiltinModule ZendModuleEntry = ZendModuleEntry{g.SizeOf("zend_module_entry"), 20190902, 0, 0, nil, nil, "Core", BuiltinFunctions, ZmStartupCore, nil, nil, nil, nil, "3.4.0", 0, nil, nil, nil, nil, 0, 0, nil, 0, "API" + "20190902" + ",NTS"}

/* }}} */

func ZendStartupBuiltinFunctions() int {
	ZendBuiltinModule.SetModuleNumber(0)
	ZendBuiltinModule.SetType(1)
	if g.Assign(&(EG.GetCurrentModule()), ZendRegisterModuleEx(&ZendBuiltinModule)) == nil {
		return FAILURE
	} else {
		return SUCCESS
	}
}

/* }}} */

func ZifZendVersion(execute_data *ZendExecuteData, return_value *Zval) {
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	var __z *Zval = return_value
	var __s *ZendString = ZendStringInit("3.4.0", g.SizeOf("ZEND_VERSION")-1, 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	return
}

/* }}} */

func ZifGcMemCaches(execute_data *ZendExecuteData, return_value *Zval) {
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	var __z *Zval = return_value
	__z.GetValue().SetLval(ZendMmGc(ZendMmGetHeap()))
	__z.SetTypeInfo(4)
	return
}

/* }}} */

func ZifGcCollectCycles(execute_data *ZendExecuteData, return_value *Zval) {
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	var __z *Zval = return_value
	__z.GetValue().SetLval(GcCollectCycles())
	__z.SetTypeInfo(4)
	return
}

/* }}} */

func ZifGcEnabled(execute_data *ZendExecuteData, return_value *Zval) {
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	if GcEnabled() != 0 {
		return_value.SetTypeInfo(3)
	} else {
		return_value.SetTypeInfo(2)
	}
	return
}

/* }}} */

func ZifGcEnable(execute_data *ZendExecuteData, return_value *Zval) {
	var key *ZendString
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	key = ZendStringInit("zend.enable_gc", g.SizeOf("\"zend.enable_gc\"")-1, 0)
	ZendAlterIniEntryChars(key, "1", g.SizeOf("\"1\"")-1, 1<<0, 1<<4)
	ZendStringReleaseEx(key, 0)
}

/* }}} */

func ZifGcDisable(execute_data *ZendExecuteData, return_value *Zval) {
	var key *ZendString
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	key = ZendStringInit("zend.enable_gc", g.SizeOf("\"zend.enable_gc\"")-1, 0)
	ZendAlterIniEntryChars(key, "0", g.SizeOf("\"0\"")-1, 1<<0, 1<<4)
	ZendStringReleaseEx(key, 0)
}

/* }}} */

func ZifGcStatus(execute_data *ZendExecuteData, return_value *Zval) {
	var status ZendGcStatus
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	ZendGcGetStatus(&status)
	var __arr *ZendArray = _zendNewArray(3)
	var __z *Zval = return_value
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	AddAssocLongEx(return_value, "runs", g.SizeOf("\"runs\"")-1, long(status.GetRuns()))
	AddAssocLongEx(return_value, "collected", g.SizeOf("\"collected\"")-1, long(status.GetCollected()))
	AddAssocLongEx(return_value, "threshold", g.SizeOf("\"threshold\"")-1, long(status.GetThreshold()))
	AddAssocLongEx(return_value, "roots", g.SizeOf("\"roots\"")-1, long(status.GetNumRoots()))
}

/* }}} */

func ZifFuncNumArgs(execute_data *ZendExecuteData, return_value *Zval) {
	var ex *ZendExecuteData = execute_data.GetPrevExecuteData()
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	if (ex.GetThis().GetTypeInfo() & 1 << 16) != 0 {
		ZendError(1<<1, "func_num_args():  Called from the global scope - no function context")
		var __z *Zval = return_value
		__z.GetValue().SetLval(-1)
		__z.SetTypeInfo(4)
		return
	}
	if ZendForbidDynamicCall("func_num_args()") == FAILURE {
		var __z *Zval = return_value
		__z.GetValue().SetLval(-1)
		__z.SetTypeInfo(4)
		return
	}
	var __z *Zval = return_value
	__z.GetValue().SetLval(ex.GetThis().GetNumArgs())
	__z.SetTypeInfo(4)
	return
}

/* }}} */

func ZifFuncGetArg(execute_data *ZendExecuteData, return_value *Zval) {
	var arg_count uint32
	var first_extra_arg uint32
	var arg *Zval
	var requested_offset ZendLong
	var ex *ZendExecuteData
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "l", &requested_offset) == FAILURE {
		return
	}
	if requested_offset < 0 {
		ZendError(1<<1, "func_get_arg():  The argument number should be >= 0")
		return_value.SetTypeInfo(2)
		return
	}
	ex = execute_data.GetPrevExecuteData()
	if (ex.GetThis().GetTypeInfo() & 1 << 16) != 0 {
		ZendError(1<<1, "func_get_arg():  Called from the global scope - no function context")
		return_value.SetTypeInfo(2)
		return
	}
	if ZendForbidDynamicCall("func_get_arg()") == FAILURE {
		return_value.SetTypeInfo(2)
		return
	}
	arg_count = ex.GetThis().GetNumArgs()
	if ZendUlong(requested_offset >= arg_count) != 0 {
		ZendError(1<<1, "func_get_arg():  Argument "+"%"+"lld"+" not passed to function", requested_offset)
		return_value.SetTypeInfo(2)
		return
	}
	first_extra_arg = ex.GetFunc().GetOpArray().GetNumArgs()
	if ZendUlong(requested_offset >= first_extra_arg && ex.GetThis().GetNumArgs() > first_extra_arg) != 0 {
		arg = (*Zval)(ex) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(ex.GetFunc().GetOpArray().GetLastVar()+ex.GetFunc().GetOpArray().GetT())) + (requested_offset - first_extra_arg)
	} else {
		arg = (*Zval)(ex) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(requested_offset+1)-1))
	}
	if arg.GetType() != 0 {
		var _z3 *Zval = arg
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = return_value
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	}
}

/* }}} */

func ZifFuncGetArgs(execute_data *ZendExecuteData, return_value *Zval) {
	var p *Zval
	var q *Zval
	var arg_count uint32
	var first_extra_arg uint32
	var i uint32
	var ex *ZendExecuteData = execute_data.GetPrevExecuteData()
	if (ex.GetThis().GetTypeInfo() & 1 << 16) != 0 {
		ZendError(1<<1, "func_get_args():  Called from the global scope - no function context")
		return_value.SetTypeInfo(2)
		return
	}
	if ZendForbidDynamicCall("func_get_args()") == FAILURE {
		return_value.SetTypeInfo(2)
		return
	}
	arg_count = ex.GetThis().GetNumArgs()
	if arg_count != 0 {
		var __arr *ZendArray = _zendNewArray(arg_count)
		var __z *Zval = return_value
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		first_extra_arg = ex.GetFunc().GetOpArray().GetNumArgs()
		ZendHashRealInitPacked(return_value.GetValue().GetArr())
		var __fill_ht *HashTable = return_value.GetValue().GetArr()
		var __fill_bkt *Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
		var __fill_idx uint32 = __fill_ht.GetNNumUsed()
		assert((__fill_ht.GetUFlags() & 1 << 2) != 0)
		i = 0
		p = (*Zval)(ex) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(1)-1))
		if arg_count > first_extra_arg {
			for i < first_extra_arg {
				q = p
				if q.GetTypeInfo() != 0 {
					if q.GetType() == 10 {
						q = &(*q).value.GetRef().GetVal()
					}
					if (q.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(q)
					}
					var _z1 *Zval = &__fill_bkt.val
					var _z2 *Zval = q
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
				} else {
					&__fill_bkt.val.u1.type_info = 1
				}
				__fill_bkt.SetH(__fill_idx)
				__fill_bkt.SetKey(nil)
				__fill_bkt++
				__fill_idx++
				p++
				i++
			}
			p = (*Zval)(ex) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(ex.GetFunc().GetOpArray().GetLastVar()+ex.GetFunc().GetOpArray().GetT()))
		}
		for i < arg_count {
			q = p
			if q.GetTypeInfo() != 0 {
				if q.GetType() == 10 {
					q = &(*q).value.GetRef().GetVal()
				}
				if (q.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(q)
				}
				var _z1 *Zval = &__fill_bkt.val
				var _z2 *Zval = q
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else {
				&__fill_bkt.val.u1.type_info = 1
			}
			__fill_bkt.SetH(__fill_idx)
			__fill_bkt.SetKey(nil)
			__fill_bkt++
			__fill_idx++
			p++
			i++
		}
		__fill_ht.SetNNumUsed(__fill_idx)
		__fill_ht.SetNNumOfElements(__fill_idx)
		__fill_ht.SetNNextFreeElement(__fill_idx)
		__fill_ht.SetNInternalPointer(0)
		return_value.GetValue().GetArr().SetNNumOfElements(arg_count)
	} else {
		var __z *Zval = return_value
		__z.GetValue().SetArr((*ZendArray)(&ZendEmptyArray))
		__z.SetTypeInfo(7)
		return
	}
}

/* }}} */

func ZifStrlen(execute_data *ZendExecuteData, return_value *Zval) {
	var s *ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &s, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	var __z *Zval = return_value
	__z.GetValue().SetLval(s.GetLen())
	__z.SetTypeInfo(4)
}

/* }}} */

func ZifStrcmp(execute_data *ZendExecuteData, return_value *Zval) {
	var s1 *ZendString
	var s2 *ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &s1, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &s2, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	var __z *Zval = return_value
	__z.GetValue().SetLval(ZendBinaryStrcmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen()))
	__z.SetTypeInfo(4)
	return
}

/* }}} */

func ZifStrncmp(execute_data *ZendExecuteData, return_value *Zval) {
	var s1 *ZendString
	var s2 *ZendString
	var len_ ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &s1, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &s2, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgLong(_arg, &len_, &_dummy, 0, 0) == 0 {
				_expected_type = Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if len_ < 0 {
		ZendError(1<<1, "Length must be greater than or equal to 0")
		return_value.SetTypeInfo(2)
		return
	}
	var __z *Zval = return_value
	__z.GetValue().SetLval(ZendBinaryStrncmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen(), len_))
	__z.SetTypeInfo(4)
	return
}

/* }}} */

func ZifStrcasecmp(execute_data *ZendExecuteData, return_value *Zval) {
	var s1 *ZendString
	var s2 *ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &s1, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &s2, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	var __z *Zval = return_value
	__z.GetValue().SetLval(ZendBinaryStrcasecmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen()))
	__z.SetTypeInfo(4)
	return
}

/* }}} */

func ZifStrncasecmp(execute_data *ZendExecuteData, return_value *Zval) {
	var s1 *ZendString
	var s2 *ZendString
	var len_ ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &s1, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &s2, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgLong(_arg, &len_, &_dummy, 0, 0) == 0 {
				_expected_type = Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if len_ < 0 {
		ZendError(1<<1, "Length must be greater than or equal to 0")
		return_value.SetTypeInfo(2)
		return
	}
	var __z *Zval = return_value
	__z.GetValue().SetLval(ZendBinaryStrncasecmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen(), len_))
	__z.SetTypeInfo(4)
	return
}

/* }}} */

func ZifEach(execute_data *ZendExecuteData, return_value *Zval) {
	var array *Zval
	var entry *Zval
	var tmp Zval
	var num_key ZendUlong
	var target_hash *HashTable
	var key *ZendString
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "z/", &array) == FAILURE {
		return
	}
	if EG.GetEachDeprecationThrown() == 0 {
		ZendError(1<<13, "The each() function is deprecated. This message will be suppressed on further calls")
		EG.SetEachDeprecationThrown(1)
	}
	if array.GetType() == 7 {
		target_hash = array.GetValue().GetArr()
	} else {
		if array.GetType() == 8 {
			target_hash = array.GetValue().GetObj().GetHandlers().GetGetProperties()(array)
		} else {
			target_hash = nil
		}
	}
	if target_hash == nil {
		ZendError(1<<1, "Variable passed to each() is not an array or object")
		return
	}
	for true {
		entry = ZendHashGetCurrentDataEx(target_hash, &target_hash.nInternalPointer)
		if entry == nil {
			return_value.SetTypeInfo(2)
			return
		} else if entry.GetType() == 13 {
			entry = entry.GetValue().GetZv()
			if entry.GetType() == 0 {
				ZendHashMoveForwardEx(target_hash, &target_hash.nInternalPointer)
				continue
			}
		}
		break
	}
	var __arr *ZendArray = _zendNewArray(4)
	var __z *Zval = return_value
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	ZendHashRealInitMixed(return_value.GetValue().GetArr())

	/* add value elements */

	if entry.GetType() == 10 {
		entry = &(*entry).value.GetRef().GetVal()
	}
	if entry.GetTypeFlags() != 0 {
		ZendGcAddrefEx(&(entry.GetValue().GetCounted()).gc, 2)
	}
	ZendHashIndexAddNew(return_value.GetValue().GetArr(), 1, entry)
	ZendHashAddNew(return_value.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_VALUE], entry)

	/* add the key elements */

	if ZendHashGetCurrentKeyEx(target_hash, &key, &num_key, &target_hash.nInternalPointer) == 1 {
		var __z *Zval = &tmp
		var __s *ZendString = key
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		if &tmp.GetTypeFlags() != 0 {
			ZvalAddrefP(&tmp)
		}
	} else {
		var __z *Zval = &tmp
		__z.GetValue().SetLval(num_key)
		__z.SetTypeInfo(4)
	}
	ZendHashIndexAddNew(return_value.GetValue().GetArr(), 0, &tmp)
	ZendHashAddNew(return_value.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_KEY], &tmp)
	ZendHashMoveForwardEx(target_hash, &target_hash.nInternalPointer)
}

/* }}} */

func ZifErrorReporting(execute_data *ZendExecuteData, return_value *Zval) {
	var err *Zval = nil
	var old_error_reporting int
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			ZendParseArgZvalDeref(_arg, &err, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	old_error_reporting = EG.GetErrorReporting()
	if execute_data.GetThis().GetNumArgs() != 0 {
		var new_val *ZendString = ZvalTryGetString(err)
		if new_val == nil {
			return
		}
		for {
			var p *ZendIniEntry = EG.GetErrorReportingIniEntry()
			if p == nil {
				var zv *Zval = ZendHashFindEx(EG.GetIniDirectives(), ZendKnownStrings[ZEND_STR_ERROR_REPORTING], 1)
				if zv != nil {
					EG.SetErrorReportingIniEntry((*ZendIniEntry)(zv.GetValue().GetPtr()))
					p = EG.GetErrorReportingIniEntry()
				} else {
					break
				}
			}
			if p.GetModified() == 0 {
				if EG.GetModifiedIniDirectives() == nil {
					EG.SetModifiedIniDirectives((*HashTable)(_emalloc(g.SizeOf("HashTable"))))
					_zendHashInit(EG.GetModifiedIniDirectives(), 8, nil, 0)
				}
				if ZendHashAddPtr(EG.GetModifiedIniDirectives(), ZendKnownStrings[ZEND_STR_ERROR_REPORTING], p) != nil {
					p.SetOrigValue(p.GetValue())
					p.SetOrigModifiable(p.GetModifiable())
					p.SetModified(1)
				}
			} else if p.GetOrigValue() != p.GetValue() {
				ZendStringReleaseEx(p.GetValue(), 0)
			}
			p.SetValue(new_val)
			if err.GetType() == 4 {
				EG.SetErrorReporting(err.GetValue().GetLval())
			} else {
				EG.SetErrorReporting(atoi(p.GetValue().GetVal()))
			}
			break
		}
	}
	var __z *Zval = return_value
	__z.GetValue().SetLval(old_error_reporting)
	__z.SetTypeInfo(4)
}

/* }}} */

func ValidateConstantArray(ht *HashTable) int {
	var ret int = 1
	var val *Zval
	ht.GetGc().SetTypeInfo(ht.GetGc().GetTypeInfo() | 1<<5<<0)
	for {
		var __ht *HashTable = ht
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val
			if _z.GetType() == 13 {
				_z = _z.GetValue().GetZv()
			}
			if _z.GetType() == 0 {
				continue
			}
			val = _z
			if val.GetType() == 10 {
				val = &(*val).value.GetRef().GetVal()
			}
			if val.GetTypeFlags() != 0 {
				if val.GetType() == 7 {
					if val.GetTypeFlags() != 0 {
						if (ZvalGcFlags(val.GetValue().GetCounted().GetGc().GetTypeInfo()) & 1 << 5) != 0 {
							ZendError(1<<1, "Constants cannot be recursive arrays")
							ret = 0
							break
						} else if ValidateConstantArray(val.GetValue().GetArr()) == 0 {
							ret = 0
							break
						}
					}
				} else if val.GetType() != 6 && val.GetType() != 9 {
					ZendError(1<<1, "Constants may only evaluate to scalar values, arrays or resources")
					ret = 0
					break
				}
			}
		}
		break
	}
	ht.GetGc().SetTypeInfo(ht.GetGc().GetTypeInfo() &^ (1 << 5 << 0))
	return ret
}

/* }}} */

func CopyConstantArray(dst *Zval, src *Zval) {
	var key *ZendString
	var idx ZendUlong
	var new_val *Zval
	var val *Zval
	var __arr *ZendArray = _zendNewArray(src.GetValue().GetArr().GetNNumOfElements())
	var __z *Zval = dst
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	for {
		var __ht *HashTable = src.GetValue().GetArr()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val
			if _z.GetType() == 13 {
				_z = _z.GetValue().GetZv()
			}
			if _z.GetType() == 0 {
				continue
			}
			idx = _p.GetH()
			key = _p.GetKey()
			val = _z

			/* constant arrays can't contain references */

			if val.GetType() == 10 {
				val = &(*val).value.GetRef().GetVal()
			}
			if key != nil {
				new_val = ZendHashAddNew(dst.GetValue().GetArr(), key, val)
			} else {
				new_val = ZendHashIndexAddNew(dst.GetValue().GetArr(), idx, val)
			}
			if val.GetType() == 7 {
				if val.GetTypeFlags() != 0 {
					CopyConstantArray(new_val, val)
				}
			} else {
				if val.GetTypeFlags() != 0 {
					ZvalAddrefP(val)
				}
			}
		}
		break
	}
}

/* }}} */

func ZifDefine(execute_data *ZendExecuteData, return_value *Zval) {
	var name *ZendString
	var val *Zval
	var val_free Zval
	var non_cs ZendBool = 0
	var case_sensitive int = 1 << 0
	var c ZendConstant
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			ZendParseArgZvalDeref(_arg, &val, 0)
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgBool(_arg, &non_cs, &_dummy, 0) == 0 {
				_expected_type = Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if non_cs != 0 {
		case_sensitive = 0
	}
	if ZendMemnstr(name.GetVal(), "::", g.SizeOf("\"::\"")-1, name.GetVal()+name.GetLen()) != nil {
		ZendError(1<<1, "Class constants cannot be defined or redefined")
		return_value.SetTypeInfo(2)
		return
	}
	&val_free.SetTypeInfo(0)
repeat:
	switch val.GetType() {
	case 4:

	case 5:

	case 6:

	case 2:

	case 3:

	case 1:

	case 9:
		break
	case 7:
		if val.GetTypeFlags() != 0 {
			if ValidateConstantArray(val.GetValue().GetArr()) == 0 {
				return_value.SetTypeInfo(2)
				return
			} else {
				CopyConstantArray(&c.value, val)
				goto register_constant
			}
		}
		break
	case 8:
		if val_free.GetType() == 0 {
			if val.GetValue().GetObj().GetHandlers().GetGet() != nil {
				val = val.GetValue().GetObj().GetHandlers().GetGet()(val, &val_free)
				goto repeat
			} else if val.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
				if val.GetValue().GetObj().GetHandlers().GetCastObject()(val, &val_free, 6) == SUCCESS {
					val = &val_free
					break
				}
			}
		}
	default:
		ZendError(1<<1, "Constants may only evaluate to scalar values, arrays or resources")
		ZvalPtrDtor(&val_free)
		return_value.SetTypeInfo(2)
		return
	}
	var _z1 *Zval = &c.value
	var _z2 *Zval = val
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
	ZvalPtrDtor(&val_free)
register_constant:
	if non_cs != 0 {
		ZendError(1<<13, "define(): Declaration of case-insensitive constants is deprecated")
	}

	/* non persistent */

	&c.GetValue().SetConstantFlags(case_sensitive&0xff | 0x7fffff<<8)
	c.SetName(ZendStringCopy(name))
	if ZendRegisterConstant(&c) == SUCCESS {
		return_value.SetTypeInfo(3)
		return
	} else {
		return_value.SetTypeInfo(2)
		return
	}
}

/* }}} */

func ZifDefined(execute_data *ZendExecuteData, return_value *Zval) {
	var name *ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if ZendGetConstantEx(name, ZendGetExecutedScope(), 0x100|0x1000) != nil {
		return_value.SetTypeInfo(3)
		return
	} else {
		return_value.SetTypeInfo(2)
		return
	}
}

/* }}} */

func ZifGetClass(execute_data *ZendExecuteData, return_value *Zval) {
	var obj *Zval = nil
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "|o", &obj) == FAILURE {
		return_value.SetTypeInfo(2)
		return
	}
	if obj == nil {
		var scope *ZendClassEntry = ZendGetExecutedScope()
		if scope != nil {
			var __z *Zval = return_value
			var __s *ZendString = scope.GetName()
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
			return
		} else {
			ZendError(1<<1, "get_class() called without object from outside a class")
			return_value.SetTypeInfo(2)
			return
		}
	}
	var __z *Zval = return_value
	var __s *ZendString = obj.GetValue().GetObj().GetCe().GetName()
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		ZendGcAddref(&__s.gc)
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	return
}

/* }}} */

func ZifGetCalledClass(execute_data *ZendExecuteData, return_value *Zval) {
	var called_scope *ZendClassEntry
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	called_scope = ZendGetCalledScope(execute_data)
	if called_scope != nil {
		var __z *Zval = return_value
		var __s *ZendString = called_scope.GetName()
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		return
	} else {
		var scope *ZendClassEntry = ZendGetExecutedScope()
		if scope == nil {
			ZendError(1<<1, "get_called_class() called from outside a class")
		}
	}
	return_value.SetTypeInfo(2)
	return
}

/* }}} */

func ZifGetParentClass(execute_data *ZendExecuteData, return_value *Zval) {
	var arg *Zval
	var ce *ZendClassEntry = nil
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "|z", &arg) == FAILURE {
		return
	}
	if execute_data.GetThis().GetNumArgs() == 0 {
		ce = ZendGetExecutedScope()
		if ce != nil && ce.parent {
			var __z *Zval = return_value
			var __s *ZendString = ce.parent.name
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
			return
		} else {
			return_value.SetTypeInfo(2)
			return
		}
	}
	if arg.GetType() == 8 {
		ce = arg.GetValue().GetObj().GetCe()
	} else if arg.GetType() == 6 {
		ce = ZendLookupClass(arg.GetValue().GetStr())
	}
	if ce != nil && ce.parent {
		var __z *Zval = return_value
		var __s *ZendString = ce.parent.name
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		return
	} else {
		return_value.SetTypeInfo(2)
		return
	}
}

/* }}} */

func IsAImpl(execute_data *ZendExecuteData, return_value *Zval, only_subclass ZendBool) {
	var obj *Zval
	var class_name *ZendString
	var instance_ce *ZendClassEntry
	var ce *ZendClassEntry
	var allow_string ZendBool = only_subclass
	var retval ZendBool
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			ZendParseArgZvalDeref(_arg, &obj, 0)
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &class_name, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgBool(_arg, &allow_string, &_dummy, 0) == 0 {
				_expected_type = Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/*
	 * allow_string - is_a default is no, is_subclass_of is yes.
	 *   if it's allowed, then the autoloader will be called if the class does not exist.
	 *   default behaviour is different, as 'is_a' used to be used to test mixed return values
	 *   and there is no easy way to deprecate this.
	 */

	if allow_string != 0 && obj.GetType() == 6 {
		instance_ce = ZendLookupClass(obj.GetValue().GetStr())
		if instance_ce == nil {
			return_value.SetTypeInfo(2)
			return
		}
	} else if obj.GetType() == 8 {
		instance_ce = obj.GetValue().GetObj().GetCe()
	} else {
		return_value.SetTypeInfo(2)
		return
	}
	if only_subclass == 0 && ZendStringEquals(instance_ce.GetName(), class_name) != 0 {
		retval = 1
	} else {
		ce = ZendLookupClassEx(class_name, nil, 0x80)
		if ce == nil {
			retval = 0
		} else {
			if only_subclass != 0 && instance_ce == ce {
				retval = 0
			} else {
				retval = InstanceofFunction(instance_ce, ce)
			}
		}
	}
	if retval != 0 {
		return_value.SetTypeInfo(3)
	} else {
		return_value.SetTypeInfo(2)
	}
	return
}

/* }}} */

func ZifIsSubclassOf(execute_data *ZendExecuteData, return_value *Zval) {
	IsAImpl(execute_data, return_value, 1)
}

/* }}} */

func ZifIsA(execute_data *ZendExecuteData, return_value *Zval) {
	IsAImpl(execute_data, return_value, 0)
}

/* }}} */

func AddClassVars(scope *ZendClassEntry, ce *ZendClassEntry, statics int, return_value *Zval) {
	var prop_info *ZendPropertyInfo
	var prop *Zval
	var prop_copy Zval
	var key *ZendString
	for {
		var __ht *HashTable = &ce.properties_info
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			key = _p.GetKey()
			prop_info = _z.GetValue().GetPtr()
			if (prop_info.GetFlags()&1<<1) != 0 && ZendCheckProtected(prop_info.GetCe(), scope) == 0 || (prop_info.GetFlags()&1<<2) != 0 && prop_info.GetCe() != scope {
				continue
			}
			prop = nil
			if statics != 0 && (prop_info.GetFlags()&1<<4) != 0 {
				prop = &ce.default_static_members_table[prop_info.GetOffset()]
				if prop.GetType() == 13 {
					prop = prop.GetValue().GetZv()
				}
			} else if statics == 0 && (prop_info.GetFlags()&1<<4) == 0 {
				prop = &ce.default_properties_table[(prop_info.GetOffset()-uint32(zend_long((*byte)(&((*ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil))+g.SizeOf("zval")*0))/g.SizeOf("zval")]
			}
			if prop == nil {
				continue
			}
			if prop.GetType() == 0 {

				/* Return uninitialized typed properties as a null value */

				&prop_copy.SetTypeInfo(1)

				/* Return uninitialized typed properties as a null value */

			} else {

				/* copy: enforce read only access */

				var _z1 *Zval = &prop_copy
				var _z2 *Zval = prop
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
						ZendGcAddref(&_gc.gc)
					} else {
						ZvalCopyCtorFunc(_z1)
					}
				}

				/* copy: enforce read only access */

			}
			prop = &prop_copy

			/* this is necessary to make it able to work with default array
			 * properties, returned to user */

			if (prop.GetTypeInfo() & 0xff) == 11 {
				if ZvalUpdateConstantEx(prop, nil) != SUCCESS {
					return
				}
			}
			ZendHashAddNew(return_value.GetValue().GetArr(), key, prop)
		}
		break
	}
}

/* }}} */

func ZifGetClassVars(execute_data *ZendExecuteData, return_value *Zval) {
	var class_name *ZendString
	var ce *ZendClassEntry
	var scope *ZendClassEntry
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "S", &class_name) == FAILURE {
		return
	}
	ce = ZendLookupClass(class_name)
	if ce == nil {
		return_value.SetTypeInfo(2)
		return
	} else {
		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = return_value
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		if (ce.GetCeFlags() & 1 << 12) == 0 {
			if ZendUpdateClassConstants(ce) != SUCCESS {
				return
			}
		}
		scope = ZendGetExecutedScope()
		AddClassVars(scope, ce, 0, return_value)
		AddClassVars(scope, ce, 1, return_value)
	}
}

/* }}} */

func ZifGetObjectVars(execute_data *ZendExecuteData, return_value *Zval) {
	var obj *Zval
	var value *Zval
	var properties *HashTable
	var key *ZendString
	var zobj *ZendObject
	var num_key ZendUlong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgObject(_arg, &obj, nil, 0) == 0 {
				_expected_type = Z_EXPECTED_OBJECT
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	properties = obj.GetValue().GetObj().GetHandlers().GetGetProperties()(obj)
	if properties == nil {
		return_value.SetTypeInfo(2)
		return
	}
	zobj = obj.GetValue().GetObj()
	if zobj.GetCe().GetDefaultPropertiesCount() == 0 && properties == zobj.GetProperties() && (ZvalGcFlags(properties.GetGc().GetTypeInfo())&1<<5) == 0 {

		/* fast copy */

		if zobj.GetHandlers() == &StdObjectHandlers {
			var __arr *ZendArray = ZendProptableToSymtable(properties, 0)
			var __z *Zval = return_value
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
			return
		}
		var __arr *ZendArray = ZendProptableToSymtable(properties, 1)
		var __z *Zval = return_value
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		return
	} else {
		var __arr *ZendArray = _zendNewArray(properties.GetNNumOfElements())
		var __z *Zval = return_value
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		for {
			var __ht *HashTable = properties
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				num_key = _p.GetH()
				key = _p.GetKey()
				value = _z
				var is_dynamic ZendBool = 1
				if value.GetType() == 13 {
					value = value.GetValue().GetZv()
					if value.GetType() == 0 {
						continue
					}
					is_dynamic = 0
				}
				if key != nil && ZendCheckPropertyAccess(zobj, key, is_dynamic) == FAILURE {
					continue
				}
				if value.GetType() == 10 && ZvalRefcountP(value) == 1 {
					value = &(*value).value.GetRef().GetVal()
				}
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
				if key == nil {

					/* This case is only possible due to loopholes, e.g. ArrayObject */

					ZendHashIndexAdd(return_value.GetValue().GetArr(), num_key, value)

					/* This case is only possible due to loopholes, e.g. ArrayObject */

				} else if is_dynamic == 0 && key.GetVal()[0] == 0 {
					var prop_name *byte
					var class_name *byte
					var prop_len int
					ZendUnmanglePropertyNameEx(key, &class_name, &prop_name, &prop_len)

					/* We assume here that a mangled property name is never
					 * numeric. This is probably a safe assumption, but
					 * theoretically someone might write an extension with
					 * private, numeric properties. Well, too bad.
					 */

					ZendHashStrAddNew(return_value.GetValue().GetArr(), prop_name, prop_len, value)

					/* We assume here that a mangled property name is never
					 * numeric. This is probably a safe assumption, but
					 * theoretically someone might write an extension with
					 * private, numeric properties. Well, too bad.
					 */

				} else {
					ZendSymtableAddNew(return_value.GetValue().GetArr(), key, value)
				}
			}
			break
		}
	}
}

/* }}} */

func ZifGetMangledObjectVars(execute_data *ZendExecuteData, return_value *Zval) {
	var obj *Zval
	var properties *HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgObject(_arg, &obj, nil, 0) == 0 {
				_expected_type = Z_EXPECTED_OBJECT
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	properties = obj.GetValue().GetObj().GetHandlers().GetGetProperties()(obj)
	if properties == nil {
		var __z *Zval = return_value
		__z.GetValue().SetArr((*ZendArray)(&ZendEmptyArray))
		__z.SetTypeInfo(7)
		return
	}
	properties = ZendProptableToSymtable(properties, obj.GetValue().GetObj().GetCe().GetDefaultPropertiesCount() != 0 || obj.GetValue().GetObj().GetHandlers() != &StdObjectHandlers || (ZvalGcFlags(properties.GetGc().GetTypeInfo())&1<<5) != 0)
	var __arr *ZendArray = properties
	var __z *Zval = return_value
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	return
}

/* }}} */

func SameName(key *ZendString, name *ZendString) int {
	var lcname *ZendString
	var ret int
	if key == name {
		return 1
	}
	if key.GetLen() != name.GetLen() {
		return 0
	}
	lcname = ZendStringTolowerEx(name, 0)
	ret = memcmp(lcname.GetVal(), key.GetVal(), key.GetLen()) == 0
	ZendStringReleaseEx(lcname, 0)
	return ret
}

/* }}} */

func ZifGetClassMethods(execute_data *ZendExecuteData, return_value *Zval) {
	var klass *Zval
	var method_name Zval
	var ce *ZendClassEntry = nil
	var scope *ZendClassEntry
	var mptr *ZendFunction
	var key *ZendString
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "z", &klass) == FAILURE {
		return
	}
	if klass.GetType() == 8 {
		ce = klass.GetValue().GetObj().GetCe()
	} else if klass.GetType() == 6 {
		ce = ZendLookupClass(klass.GetValue().GetStr())
	}
	if ce == nil {
		return_value.SetTypeInfo(1)
		return
	}
	var __arr *ZendArray = _zendNewArray(0)
	var __z *Zval = return_value
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	scope = ZendGetExecutedScope()
	for {
		var __ht *HashTable = &ce.function_table
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			key = _p.GetKey()
			mptr = _z.GetValue().GetPtr()
			if (mptr.GetFnFlags()&1<<0) != 0 || scope != nil && ((mptr.GetFnFlags()&1<<1) != 0 && ZendCheckProtected(mptr.GetScope(), scope) != 0 || (mptr.GetFnFlags()&1<<2) != 0 && scope == mptr.GetScope()) {
				if mptr.GetType() == 2 && (mptr.GetOpArray().GetRefcount() == nil || (*mptr).op_array.refcount > 1) && key != nil && SameName(key, mptr.GetFunctionName()) == 0 {
					var __z *Zval = &method_name
					var __s *ZendString = ZendFindAliasName(mptr.GetScope(), key)
					__z.GetValue().SetStr(__s)
					if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
						__z.SetTypeInfo(6)
					} else {
						ZendGcAddref(&__s.gc)
						__z.SetTypeInfo(6 | 1<<0<<8)
					}
					ZendHashNextIndexInsertNew(return_value.GetValue().GetArr(), &method_name)
				} else {
					var __z *Zval = &method_name
					var __s *ZendString = mptr.GetFunctionName()
					__z.GetValue().SetStr(__s)
					if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
						__z.SetTypeInfo(6)
					} else {
						ZendGcAddref(&__s.gc)
						__z.SetTypeInfo(6 | 1<<0<<8)
					}
					ZendHashNextIndexInsertNew(return_value.GetValue().GetArr(), &method_name)
				}
			}
		}
		break
	}
}

/* }}} */

func ZifMethodExists(execute_data *ZendExecuteData, return_value *Zval) {
	var klass *Zval
	var method_name *ZendString
	var lcname *ZendString
	var ce *ZendClassEntry
	var func_ *ZendFunction
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			ZendParseArgZvalDeref(_arg, &klass, 0)
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &method_name, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if klass.GetType() == 8 {
		ce = klass.GetValue().GetObj().GetCe()
	} else if klass.GetType() == 6 {
		if g.Assign(&ce, ZendLookupClass(klass.GetValue().GetStr())) == nil {
			return_value.SetTypeInfo(2)
			return
		}
	} else {
		return_value.SetTypeInfo(2)
		return
	}
	lcname = ZendStringTolowerEx(method_name, 0)
	func_ = ZendHashFindPtr(&ce.function_table, lcname)
	ZendStringReleaseEx(lcname, 0)
	if func_ != nil {

		/* Exclude shadow properties when checking a method on a specific class. Include
		 * them when checking an object, as method_exists() generally ignores visibility.
		 * TODO: Should we use EG(scope) for the object case instead? */

		if klass.GetType() == 8 || (func_.GetFnFlags()&1<<2) == 0 || func_.GetScope() == ce {
			return_value.SetTypeInfo(3)
		} else {
			return_value.SetTypeInfo(2)
		}
		return
	}
	if klass.GetType() == 8 {
		var obj *ZendObject = klass.GetValue().GetObj()
		func_ = klass.GetValue().GetObj().GetHandlers().GetGetMethod()(&obj, method_name, nil)
		if func_ != nil {
			if (func_.GetFnFlags() & 1 << 18) != 0 {

				/* Returns true to the fake Closure's __invoke */

				if func_.GetScope() == ZendCeClosure && (method_name.GetLen() == g.SizeOf("ZEND_INVOKE_FUNC_NAME")-1 && !(memcmp(method_name.GetVal(), "__invoke", g.SizeOf("ZEND_INVOKE_FUNC_NAME")-1))) {
					return_value.SetTypeInfo(3)
				} else {
					return_value.SetTypeInfo(2)
				}
				ZendStringReleaseEx(func_.GetFunctionName(), 0)
				if func_ == &EG.trampoline {
					EG.GetTrampoline().SetFunctionName(nil)
				} else {
					_efree(func_)
				}
				return
			}
			return_value.SetTypeInfo(3)
			return
		}
	}
	return_value.SetTypeInfo(2)
	return
}

/* }}} */

func ZifPropertyExists(execute_data *ZendExecuteData, return_value *Zval) {
	var object *Zval
	var property *ZendString
	var ce *ZendClassEntry
	var property_info *ZendPropertyInfo
	var property_z Zval
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "zS", &object, &property) == FAILURE {
		return
	}
	if property == nil {
		return_value.SetTypeInfo(2)
		return
	}
	if object.GetType() == 6 {
		ce = ZendLookupClass(object.GetValue().GetStr())
		if ce == nil {
			return_value.SetTypeInfo(2)
			return
		}
	} else if object.GetType() == 8 {
		ce = object.GetValue().GetObj().GetCe()
	} else {
		ZendError(1<<1, "First parameter must either be an object or the name of an existing class")
		return_value.SetTypeInfo(1)
		return
	}
	property_info = ZendHashFindPtr(&ce.properties_info, property)
	if property_info != nil && ((property_info.GetFlags()&1<<2) == 0 || property_info.GetCe() == ce) {
		return_value.SetTypeInfo(3)
		return
	}
	var __z *Zval = &property_z
	var __s *ZendString = property
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	if object.GetType() == 8 && object.GetValue().GetObj().GetHandlers().GetHasProperty()(object, &property_z, 2, nil) != 0 {
		return_value.SetTypeInfo(3)
		return
	}
	return_value.SetTypeInfo(2)
	return
}

/* }}} */

func ClassExistsImpl(execute_data *ZendExecuteData, return_value *Zval, flags int, skip_flags int) {
	var name *ZendString
	var lcname *ZendString
	var ce *ZendClassEntry
	var autoload ZendBool = 1
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgBool(_arg, &autoload, &_dummy, 0) == 0 {
				_expected_type = Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if autoload == 0 {
		if name.GetVal()[0] == '\\' {

			/* Ignore leading "\" */

			lcname = ZendStringAlloc(name.GetLen()-1, 0)
			ZendStrTolowerCopy(lcname.GetVal(), name.GetVal()+1, name.GetLen()-1)
		} else {
			lcname = ZendStringTolowerEx(name, 0)
		}
		ce = ZendHashFindPtr(EG.GetClassTable(), lcname)
		ZendStringReleaseEx(lcname, 0)
	} else {
		ce = ZendLookupClass(name)
	}
	if ce != nil {
		if (ce.GetCeFlags()&flags) == flags && (ce.GetCeFlags()&skip_flags) == 0 {
			return_value.SetTypeInfo(3)
		} else {
			return_value.SetTypeInfo(2)
		}
		return
	} else {
		return_value.SetTypeInfo(2)
		return
	}
}

/* {{{ */

func ZifClassExists(execute_data *ZendExecuteData, return_value *Zval) {
	ClassExistsImpl(execute_data, return_value, 1<<3, 1<<0|1<<1)
}

/* }}} */

func ZifInterfaceExists(execute_data *ZendExecuteData, return_value *Zval) {
	ClassExistsImpl(execute_data, return_value, 1<<3|1<<0, 0)
}

/* }}} */

func ZifTraitExists(execute_data *ZendExecuteData, return_value *Zval) {
	ClassExistsImpl(execute_data, return_value, 1<<1, 0)
}

/* }}} */

func ZifFunctionExists(execute_data *ZendExecuteData, return_value *Zval) {
	var name *ZendString
	var func_ *ZendFunction
	var lcname *ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.GetThis().GetNumArgs()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if name.GetVal()[0] == '\\' {

		/* Ignore leading "\" */

		lcname = ZendStringAlloc(name.GetLen()-1, 0)
		ZendStrTolowerCopy(lcname.GetVal(), name.GetVal()+1, name.GetLen()-1)
	} else {
		lcname = ZendStringTolowerEx(name, 0)
	}
	func_ = ZendHashFindPtr(EG.GetFunctionTable(), lcname)
	ZendStringReleaseEx(lcname, 0)

	/*
	 * A bit of a hack, but not a bad one: we see if the handler of the function
	 * is actually one that displays "function is disabled" message.
	 */

	if func_ != nil && (func_.GetType() != 1 || func_.GetInternalFunction().GetHandler() != ZifDisplayDisabledFunction) {
		return_value.SetTypeInfo(3)
	} else {
		return_value.SetTypeInfo(2)
	}
	return
}

/* }}} */

func ZifClassAlias(execute_data *ZendExecuteData, return_value *Zval) {
	var class_name *ZendString
	var alias_name *byte
	var ce *ZendClassEntry
	var alias_name_len int
	var autoload ZendBool = 1
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "Ss|b", &class_name, &alias_name, &alias_name_len, &autoload) == FAILURE {
		return
	}
	ce = ZendLookupClassEx(class_name, nil, g.Cond(autoload == 0, 0x80, 0))
	if ce != nil {
		if ce.GetType() == 2 {
			if ZendRegisterClassAliasEx(alias_name, alias_name_len, ce, 0) == SUCCESS {
				return_value.SetTypeInfo(3)
				return
			} else {
				ZendError(1<<1, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), alias_name)
				return_value.SetTypeInfo(2)
				return
			}
		} else {
			ZendError(1<<1, "First argument of class_alias() must be a name of user defined class")
			return_value.SetTypeInfo(2)
			return
		}
	} else {
		ZendError(1<<1, "Class '%s' not found", class_name.GetVal())
		return_value.SetTypeInfo(2)
		return
	}
}

/* }}} */

func ZifGetIncludedFiles(execute_data *ZendExecuteData, return_value *Zval) {
	var entry *ZendString
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	var __arr *ZendArray = _zendNewArray(0)
	var __z *Zval = return_value
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	for {
		var __ht *HashTable = &EG.included_files
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			entry = _p.GetKey()
			if entry != nil {
				AddNextIndexStr(return_value, ZendStringCopy(entry))
			}
		}
		break
	}
}

/* }}} */

func ZifTriggerError(execute_data *ZendExecuteData, return_value *Zval) {
	var error_type ZendLong = 1 << 10
	var message *byte
	var message_len int
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "s|l", &message, &message_len, &error_type) == FAILURE {
		return
	}
	switch error_type {
	case 1 << 8:

	case 1 << 9:

	case 1 << 10:

	case 1 << 14:
		break
	default:
		ZendError(1<<1, "Invalid error type specified")
		return_value.SetTypeInfo(2)
		return
		break
	}
	ZendError(int(error_type), "%s", message)
	return_value.SetTypeInfo(3)
	return
}

/* }}} */

func ZifSetErrorHandler(execute_data *ZendExecuteData, return_value *Zval) {
	var error_handler *Zval
	var error_type ZendLong = 1<<0 | 1<<1 | 1<<2 | 1<<3 | 1<<4 | 1<<5 | 1<<6 | 1<<7 | 1<<8 | 1<<9 | 1<<10 | 1<<12 | 1<<13 | 1<<14 | 1<<11
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "z|l", &error_handler, &error_type) == FAILURE {
		return
	}
	if error_handler.GetType() != 1 {
		if ZendIsCallable(error_handler, 0, nil) == 0 {
			var error_handler_name *ZendString = ZendGetCallableName(error_handler)
			ZendError(1<<1, "%s() expects the argument (%s) to be a valid callback", GetActiveFunctionName(), g.CondF1(error_handler_name != nil, func() []byte { return error_handler_name.GetVal() }, "unknown"))
			ZendStringReleaseEx(error_handler_name, 0)
			return
		}
	}
	if EG.GetUserErrorHandler().GetType() != 0 {
		var _z1 *Zval = return_value
		var _z2 *Zval = &EG.user_error_handler
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	ZendStackPush(&EG.user_error_handlers_error_reporting, &EG.user_error_handler_error_reporting)
	ZendStackPush(&EG.user_error_handlers, &EG.user_error_handler)
	if error_handler.GetType() == 1 {
		&EG.user_error_handler.u1.type_info = 0
		return
	}
	var _z1 *Zval = &EG.user_error_handler
	var _z2 *Zval = error_handler
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
	EG.SetUserErrorHandlerErrorReporting(int(error_type))
}

/* }}} */

func ZifRestoreErrorHandler(execute_data *ZendExecuteData, return_value *Zval) {
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	if EG.GetUserErrorHandler().GetType() != 0 {
		var zeh Zval
		var _z1 *Zval = &zeh
		var _z2 *Zval = &EG.user_error_handler
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		&EG.user_error_handler.u1.type_info = 0
		ZvalPtrDtor(&zeh)
	}
	if ZendStackIsEmpty(&EG.user_error_handlers) != 0 {
		&EG.user_error_handler.u1.type_info = 0
	} else {
		var tmp *Zval
		EG.SetUserErrorHandlerErrorReporting(ZendStackIntTop(&EG.user_error_handlers_error_reporting))
		ZendStackDelTop(&EG.user_error_handlers_error_reporting)
		tmp = ZendStackTop(&EG.user_error_handlers)
		var _z1 *Zval = &EG.user_error_handler
		var _z2 *Zval = tmp
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		ZendStackDelTop(&EG.user_error_handlers)
	}
	return_value.SetTypeInfo(3)
	return
}

/* }}} */

func ZifSetExceptionHandler(execute_data *ZendExecuteData, return_value *Zval) {
	var exception_handler *Zval
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "z", &exception_handler) == FAILURE {
		return
	}
	if exception_handler.GetType() != 1 {
		if ZendIsCallable(exception_handler, 0, nil) == 0 {
			var exception_handler_name *ZendString = ZendGetCallableName(exception_handler)
			ZendError(1<<1, "%s() expects the argument (%s) to be a valid callback", GetActiveFunctionName(), g.CondF1(exception_handler_name != nil, func() []byte { return exception_handler_name.GetVal() }, "unknown"))
			ZendStringReleaseEx(exception_handler_name, 0)
			return
		}
	}
	if EG.GetUserExceptionHandler().GetType() != 0 {
		var _z1 *Zval = return_value
		var _z2 *Zval = &EG.user_exception_handler
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	ZendStackPush(&EG.user_exception_handlers, &EG.user_exception_handler)
	if exception_handler.GetType() == 1 {
		&EG.user_exception_handler.u1.type_info = 0
		return
	}
	var _z1 *Zval = &EG.user_exception_handler
	var _z2 *Zval = exception_handler
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
}

/* }}} */

func ZifRestoreExceptionHandler(execute_data *ZendExecuteData, return_value *Zval) {
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	if EG.GetUserExceptionHandler().GetType() != 0 {
		ZvalPtrDtor(&EG.user_exception_handler)
	}
	if ZendStackIsEmpty(&EG.user_exception_handlers) != 0 {
		&EG.user_exception_handler.u1.type_info = 0
	} else {
		var tmp *Zval = ZendStackTop(&EG.user_exception_handlers)
		var _z1 *Zval = &EG.user_exception_handler
		var _z2 *Zval = tmp
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		ZendStackDelTop(&EG.user_exception_handlers)
	}
	return_value.SetTypeInfo(3)
	return
}

/* }}} */

func CopyClassOrInterfaceName(array *Zval, key *ZendString, ce *ZendClassEntry) {
	if ce.GetRefcount() == 1 && (ce.GetCeFlags()&1<<7) == 0 || SameName(key, ce.GetName()) != 0 {
		key = ce.GetName()
	}
	AddNextIndexStr(array, ZendStringCopy(key))
}

/* }}} */

func GetDeclaredClassImpl(execute_data *ZendExecuteData, return_value *Zval, flags int, skip_flags int) {
	var key *ZendString
	var ce *ZendClassEntry
	if g.CondF2(execute_data.GetThis().GetNumArgs() == 0, SUCCESS, func() ZEND_RESULT_CODE {
		ZendWrongParametersNoneError()
		return FAILURE
	}) == FAILURE {
		return
	}
	var __arr *ZendArray = _zendNewArray(0)
	var __z *Zval = return_value
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	for {
		var __ht *HashTable = EG.GetClassTable()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			key = _p.GetKey()
			ce = _z.GetValue().GetPtr()
			if key != nil && key.GetVal()[0] != 0 && (ce.GetCeFlags()&flags) != 0 && (ce.GetCeFlags()&skip_flags) == 0 {
				CopyClassOrInterfaceName(return_value, key, ce)
			}
		}
		break
	}
}

/* {{{ */

func ZifGetDeclaredTraits(execute_data *ZendExecuteData, return_value *Zval) {
	GetDeclaredClassImpl(execute_data, return_value, 1<<1, 0)
}

/* }}} */

func ZifGetDeclaredClasses(execute_data *ZendExecuteData, return_value *Zval) {
	GetDeclaredClassImpl(execute_data, return_value, 1<<3, 1<<0|1<<1)
}

/* }}} */

func ZifGetDeclaredInterfaces(execute_data *ZendExecuteData, return_value *Zval) {
	GetDeclaredClassImpl(execute_data, return_value, 1<<0, 0)
}

/* }}} */

func ZifGetDefinedFunctions(execute_data *ZendExecuteData, return_value *Zval) {
	var internal Zval
	var user Zval
	var key *ZendString
	var func_ *ZendFunction
	var exclude_disabled ZendBool = 0
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "|b", &exclude_disabled) == FAILURE {
		return
	}
	var __arr *ZendArray = _zendNewArray(0)
	var __z *Zval = &internal
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	var __arr *ZendArray = _zendNewArray(0)
	var __z *Zval = &user
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	var __arr *ZendArray = _zendNewArray(0)
	var __z *Zval = return_value
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	for {
		var __ht *HashTable = EG.GetFunctionTable()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			key = _p.GetKey()
			func_ = _z.GetValue().GetPtr()
			if key != nil && key.GetVal()[0] != 0 {
				if func_.GetType() == 1 && (exclude_disabled == 0 || func_.GetInternalFunction().GetHandler() != ZifDisplayDisabledFunction) {
					AddNextIndexStr(&internal, ZendStringCopy(key))
				} else if func_.GetType() == 2 {
					AddNextIndexStr(&user, ZendStringCopy(key))
				}
			}
		}
		break
	}
	ZendHashStrAddNew(return_value.GetValue().GetArr(), "internal", g.SizeOf("\"internal\"")-1, &internal)
	ZendHashStrAddNew(return_value.GetValue().GetArr(), "user", g.SizeOf("\"user\"")-1, &user)
}

/* }}} */

func ZifGetDefinedVars(execute_data *ZendExecuteData, return_value *Zval) {
	var symbol_table *ZendArray
	if ZendForbidDynamicCall("get_defined_vars()") == FAILURE {
		return
	}
	symbol_table = ZendRebuildSymbolTable()
	if symbol_table == nil {
		return
	}
	var __arr *ZendArray = ZendArrayDup(symbol_table)
	var __z *Zval = return_value
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	return
}

/* }}} */

// #define LAMBDA_TEMP_FUNCNAME       "__lambda_func"

/* {{{ proto string create_function(string args, string code)
   Creates an anonymous function, and returns its name (funny, eh?) */

func ZifCreateFunction(execute_data *ZendExecuteData, return_value *Zval) {
	var function_name *ZendString
	var eval_code *byte
	var function_args *byte
	var function_code *byte
	var eval_code_length int
	var function_args_len int
	var function_code_len int
	var retval int
	var eval_name *byte
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "ss", &function_args, &function_args_len, &function_code, &function_code_len) == FAILURE {
		return
	}
	eval_code = (*byte)(_emalloc(g.SizeOf("\"function \" LAMBDA_TEMP_FUNCNAME") + function_args_len + 2 + 2 + function_code_len))
	eval_code_length = g.SizeOf("\"function \" LAMBDA_TEMP_FUNCNAME \"(\"") - 1
	memcpy(eval_code, "function "+"__lambda_func"+"(", eval_code_length)
	memcpy(eval_code+eval_code_length, function_args, function_args_len)
	eval_code_length += function_args_len
	eval_code[g.PostInc(&eval_code_length)] = ')'
	eval_code[g.PostInc(&eval_code_length)] = '{'
	memcpy(eval_code+eval_code_length, function_code, function_code_len)
	eval_code_length += function_code_len
	eval_code[g.PostInc(&eval_code_length)] = '}'
	eval_code[eval_code_length] = '0'
	eval_name = ZendMakeCompiledStringDescription("runtime-created function")
	retval = ZendEvalStringl(eval_code, eval_code_length, nil, eval_name)
	_efree(eval_code)
	_efree(eval_name)
	if retval == SUCCESS {
		var func_ *ZendOpArray
		var static_variables *HashTable
		func_ = ZendHashStrFindPtr(EG.GetFunctionTable(), "__lambda_func", g.SizeOf("LAMBDA_TEMP_FUNCNAME")-1)
		if func_ == nil {
			ZendErrorNoreturn(1<<4, "Unexpected inconsistency in create_function()")
			return_value.SetTypeInfo(2)
			return
		}
		if func_.GetRefcount() != nil {
			(*func_).refcount++
		}
		static_variables = func_.GetStaticVariables()
		func_.SetStaticVariables(nil)
		ZendHashStrDel(EG.GetFunctionTable(), "__lambda_func", g.SizeOf("LAMBDA_TEMP_FUNCNAME")-1)
		func_.SetStaticVariables(static_variables)
		function_name = ZendStringAlloc(g.SizeOf("\"0lambda_\"")+20, 0)
		function_name.GetVal()[0] = '0'
		for {
			function_name.SetLen(snprintf(function_name.GetVal()+1, g.SizeOf("\"lambda_\"")+20, "lambda_%d", g.PreInc(&(EG.GetLambdaCount()))) + 1)
			if ZendHashAddPtr(EG.GetFunctionTable(), function_name, func_) != nil {
				break
			}
		}
		var __z *Zval = return_value
		var __s *ZendString = function_name
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		return
	} else {
		ZendHashStrDel(EG.GetFunctionTable(), "__lambda_func", g.SizeOf("LAMBDA_TEMP_FUNCNAME")-1)
		return_value.SetTypeInfo(2)
		return
	}
}

/* }}} */

/* {{{ proto string get_resource_type(resource res)
   Get the resource type name for a given resource */

func ZifGetResourceType(execute_data *ZendExecuteData, return_value *Zval) {
	var resource_type *byte
	var z_resource_type *Zval
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "r", &z_resource_type) == FAILURE {
		return
	}
	resource_type = ZendRsrcListGetRsrcType(z_resource_type.GetValue().GetRes())
	if resource_type != nil {
		var _s *byte = resource_type
		var __z *Zval = return_value
		var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		return
	} else {
		var _s *byte = "Unknown"
		var __z *Zval = return_value
		var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		return
	}
}

/* }}} */

func ZifGetResources(execute_data *ZendExecuteData, return_value *Zval) {
	var type_ *ZendString = nil
	var key *ZendString
	var index ZendUlong
	var val *Zval
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "|S", &type_) == FAILURE {
		return
	}
	if type_ == nil {
		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = return_value
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		for {
			var __ht *HashTable = &EG.regular_list
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				index = _p.GetH()
				key = _p.GetKey()
				val = _z
				if key == nil {
					ZvalAddrefP(val)
					ZendHashIndexAddNew(return_value.GetValue().GetArr(), index, val)
				}
			}
			break
		}
	} else if type_.GetLen() == g.SizeOf("\"Unknown\"")-1 && !(memcmp(type_.GetVal(), "Unknown", g.SizeOf("\"Unknown\"")-1)) {
		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = return_value
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		for {
			var __ht *HashTable = &EG.regular_list
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				index = _p.GetH()
				key = _p.GetKey()
				val = _z
				if key == nil && val.GetValue().GetRes().GetType() <= 0 {
					ZvalAddrefP(val)
					ZendHashIndexAddNew(return_value.GetValue().GetArr(), index, val)
				}
			}
			break
		}
	} else {
		var id int = ZendFetchListDtorId(type_.GetVal())
		if id <= 0 {
			ZendError(1<<1, "get_resources():  Unknown resource type '%s'", type_.GetVal())
			return_value.SetTypeInfo(2)
			return
		}
		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = return_value
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		for {
			var __ht *HashTable = &EG.regular_list
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				index = _p.GetH()
				key = _p.GetKey()
				val = _z
				if key == nil && val.GetValue().GetRes().GetType() == id {
					ZvalAddrefP(val)
					ZendHashIndexAddNew(return_value.GetValue().GetArr(), index, val)
				}
			}
			break
		}
	}
}

/* }}} */

func AddZendextInfo(ext *ZendExtension, arg any) int {
	var name_array *Zval = (*Zval)(arg)
	AddNextIndexString(name_array, ext.GetName())
	return 0
}

/* }}} */

func ZifGetLoadedExtensions(execute_data *ZendExecuteData, return_value *Zval) {
	var zendext ZendBool = 0
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "|b", &zendext) == FAILURE {
		return
	}
	var __arr *ZendArray = _zendNewArray(0)
	var __z *Zval = return_value
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	if zendext != 0 {
		ZendLlistApplyWithArgument(&ZendExtensions, LlistApplyWithArgFuncT(AddZendextInfo), return_value)
	} else {
		var module *ZendModuleEntry
		for {
			var __ht *HashTable = &ModuleRegistry
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				module = _z.GetValue().GetPtr()
				AddNextIndexString(return_value, module.GetName())
			}
			break
		}
	}
}

/* }}} */

func ZifGetDefinedConstants(execute_data *ZendExecuteData, return_value *Zval) {
	var categorize ZendBool = 0
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "|b", &categorize) == FAILURE {
		return
	}
	var __arr *ZendArray = _zendNewArray(0)
	var __z *Zval = return_value
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	if categorize != 0 {
		var val *ZendConstant
		var module_number int
		var modules *Zval
		var const_val Zval
		var module_names **byte
		var module *ZendModuleEntry
		var i int = 1
		modules = _ecalloc(&ModuleRegistry.GetNNumOfElements()+2, g.SizeOf("zval"))
		module_names = _emalloc((&ModuleRegistry.GetNNumOfElements() + 2) * g.SizeOf("char *"))
		module_names[0] = "internal"
		for {
			var __ht *HashTable = &ModuleRegistry
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				module = _z.GetValue().GetPtr()
				module_names[module.GetModuleNumber()] = (*byte)(module.GetName())
				i++
			}
			break
		}
		module_names[i] = "user"
		for {
			var __ht *HashTable = EG.GetZendConstants()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				val = _z.GetValue().GetPtr()
				if val.GetName() == nil {

					/* skip special constants */

					continue

					/* skip special constants */

				}
				if val.GetValue().GetConstantFlags()>>8 == 0x7fffff {
					module_number = i
				} else if val.GetValue().GetConstantFlags()>>8 > i {

					/* should not happen */

					continue

					/* should not happen */

				} else {
					module_number = val.GetValue().GetConstantFlags() >> 8
				}
				if modules[module_number].GetType() == 0 {
					var __arr *ZendArray = _zendNewArray(0)
					var __z *Zval = &modules[module_number]
					__z.GetValue().SetArr(__arr)
					__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
					AddAssocZvalEx(return_value, module_names[module_number], strlen(module_names[module_number]), &modules[module_number])
				}
				var _z1 *Zval = &const_val
				var _z2 *Zval = &val.value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
						ZendGcAddref(&_gc.gc)
					} else {
						ZvalCopyCtorFunc(_z1)
					}
				}
				ZendHashAddNew(modules[module_number].GetValue().GetArr(), val.GetName(), &const_val)
			}
			break
		}
		_efree(module_names)
		_efree(modules)
	} else {
		var constant *ZendConstant
		var const_val Zval
		for {
			var __ht *HashTable = EG.GetZendConstants()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				constant = _z.GetValue().GetPtr()
				if constant.GetName() == nil {

					/* skip special constants */

					continue

					/* skip special constants */

				}
				var _z1 *Zval = &const_val
				var _z2 *Zval = &constant.value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
						ZendGcAddref(&_gc.gc)
					} else {
						ZvalCopyCtorFunc(_z1)
					}
				}
				ZendHashAddNew(return_value.GetValue().GetArr(), constant.GetName(), &const_val)
			}
			break
		}
	}
}

/* }}} */

func DebugBacktraceGetArgs(call *ZendExecuteData, arg_array *Zval) {
	var num_args uint32 = call.GetThis().GetNumArgs()
	if num_args != 0 {
		var i uint32 = 0
		var p *Zval = (*Zval)(call) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(1)-1))
		var __arr *ZendArray = _zendNewArray(num_args)
		var __z *Zval = arg_array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		ZendHashRealInitPacked(arg_array.GetValue().GetArr())
		var __fill_ht *HashTable = arg_array.GetValue().GetArr()
		var __fill_bkt *Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
		var __fill_idx uint32 = __fill_ht.GetNNumUsed()
		assert((__fill_ht.GetUFlags() & 1 << 2) != 0)
		if call.GetFunc().GetType() == 2 {
			var first_extra_arg uint32 = g.CondF2(num_args < call.GetFunc().GetOpArray().GetNumArgs(), num_args, func() uint32 { return call.GetFunc().GetOpArray().GetNumArgs() })
			if (call.GetThis().GetTypeInfo() & 1 << 20) != 0 {

				/* In case of attached symbol_table, values on stack may be invalid
				 * and we have to access them through symbol_table
				 * See: https://bugs.php.net/bug.php?id=73156
				 */

				var arg_name *ZendString
				var arg *Zval
				for i < first_extra_arg {
					arg_name = call.GetFunc().GetOpArray().GetVars()[i]
					arg = ZendHashFindExInd(call.GetSymbolTable(), arg_name, 1)
					if arg != nil {
						if (arg.GetTypeInfo() & 0xff00) != 0 {
							ZvalAddrefP(arg)
						}
						var _z1 *Zval = &__fill_bkt.val
						var _z2 *Zval = arg
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
					} else {
						&__fill_bkt.val.u1.type_info = 1
					}
					__fill_bkt.SetH(__fill_idx)
					__fill_bkt.SetKey(nil)
					__fill_bkt++
					__fill_idx++
					i++
				}
			} else {
				for i < first_extra_arg {
					if p.GetTypeInfo() != 0 {
						if (p.GetTypeInfo() & 0xff00) != 0 {
							ZvalAddrefP(p)
						}
						var _z1 *Zval = &__fill_bkt.val
						var _z2 *Zval = p
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
					} else {
						&__fill_bkt.val.u1.type_info = 1
					}
					__fill_bkt.SetH(__fill_idx)
					__fill_bkt.SetKey(nil)
					__fill_bkt++
					__fill_idx++
					p++
					i++
				}
			}
			p = (*Zval)(call) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(call.GetFunc().GetOpArray().GetLastVar()+call.GetFunc().GetOpArray().GetT()))
		}
		for i < num_args {
			if p.GetTypeInfo() != 0 {
				if (p.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(p)
				}
				var _z1 *Zval = &__fill_bkt.val
				var _z2 *Zval = p
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else {
				&__fill_bkt.val.u1.type_info = 1
			}
			__fill_bkt.SetH(__fill_idx)
			__fill_bkt.SetKey(nil)
			__fill_bkt++
			__fill_idx++
			p++
			i++
		}
		__fill_ht.SetNNumUsed(__fill_idx)
		__fill_ht.SetNNumOfElements(__fill_idx)
		__fill_ht.SetNNextFreeElement(__fill_idx)
		__fill_ht.SetNInternalPointer(0)
		arg_array.GetValue().GetArr().SetNNumOfElements(num_args)
	} else {
		var __z *Zval = arg_array
		__z.GetValue().SetArr((*ZendArray)(&ZendEmptyArray))
		__z.SetTypeInfo(7)
	}
}

/* }}} */

func DebugPrintBacktraceArgs(arg_array *Zval) {
	var tmp *Zval
	var i int = 0
	for {
		var __ht *HashTable = arg_array.GetValue().GetArr()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			tmp = _z
			if g.PostInc(&i) {
				ZendWrite(", ", strlen(", "))
			}
			ZendPrintFlatZvalR(tmp)
		}
		break
	}
}

/* }}} */

func SkipInternalHandler(skip *ZendExecuteData) ZendBool {
	return !(skip.GetFunc() != nil && (skip.GetFunc().GetCommonType()&1) == 0) && skip.GetPrevExecuteData() != nil && skip.GetPrevExecuteData().GetFunc() != nil && (skip.GetPrevExecuteData().GetFunc().GetCommonType()&1) == 0 && skip.GetPrevExecuteData().GetOpline().GetOpcode() != 60 && skip.GetPrevExecuteData().GetOpline().GetOpcode() != 129 && skip.GetPrevExecuteData().GetOpline().GetOpcode() != 130 && skip.GetPrevExecuteData().GetOpline().GetOpcode() != 131 && skip.GetPrevExecuteData().GetOpline().GetOpcode() != 73
}

/* {{{ */

func ZifDebugPrintBacktrace(execute_data *ZendExecuteData, return_value *Zval) {
	var call *ZendExecuteData
	var ptr *ZendExecuteData
	var skip *ZendExecuteData
	var object *ZendObject
	var lineno int
	var frameno int = 0
	var func_ *ZendFunction
	var function_name *byte
	var filename *byte
	var class_name *ZendString = nil
	var call_type *byte
	var include_filename *byte = nil
	var arg_array Zval
	var indent int = 0
	var options ZendLong = 0
	var limit ZendLong = 0
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "|ll", &options, &limit) == FAILURE {
		return
	}
	&arg_array.SetTypeInfo(0)
	ptr = execute_data.GetPrevExecuteData()

	/* skip debug_backtrace() */

	call = ptr
	ptr = ptr.GetPrevExecuteData()
	for ptr != nil && (limit == 0 || frameno < limit) {
		frameno++
		class_name = nil
		call_type = nil
		&arg_array.SetTypeInfo(0)
		ptr = ZendGeneratorCheckPlaceholderFrame(ptr)
		skip = ptr

		/* skip internal handler */

		if SkipInternalHandler(skip) != 0 {
			skip = skip.GetPrevExecuteData()
		}
		if skip.GetFunc() != nil && (skip.GetFunc().GetCommonType()&1) == 0 {
			filename = skip.GetFunc().GetOpArray().GetFilename().GetVal()
			if skip.GetOpline().GetOpcode() == 149 {
				if EG.GetOplineBeforeException() != nil {
					lineno = EG.GetOplineBeforeException().GetLineno()
				} else {
					lineno = skip.GetFunc().GetOpArray().GetLineEnd()
				}
			} else {
				lineno = skip.GetOpline().GetLineno()
			}
		} else {
			filename = nil
			lineno = 0
		}

		/* $this may be passed into regular internal functions */

		if call.GetThis().GetType() == 8 {
			object = call.GetThis().GetValue().GetObj()
		} else {
			object = nil
		}
		if call.GetFunc() != nil {
			var zend_function_name *ZendString
			func_ = call.GetFunc()
			if func_.GetScope() != nil && func_.GetScope().GetTraitAliases() != nil {
				zend_function_name = ZendResolveMethodName(g.CondF(object != nil, func() *ZendClassEntry { return object.GetCe() }, func() *ZendClassEntry { return func_.GetScope() }), func_)
			} else {
				zend_function_name = func_.GetFunctionName()
			}
			if zend_function_name != nil {
				function_name = zend_function_name.GetVal()
			} else {
				function_name = nil
			}
		} else {
			func_ = nil
			function_name = nil
		}
		if function_name != nil {
			if object != nil {
				if func_.GetScope() != nil {
					class_name = func_.GetScope().GetName()
				} else if object.GetHandlers().GetGetClassName() == ZendStdGetClassName {
					class_name = object.GetCe().GetName()
				} else {
					class_name = object.GetHandlers().GetGetClassName()(object)
				}
				call_type = "->"
			} else if func_.GetScope() != nil {
				class_name = func_.GetScope().GetName()
				call_type = "::"
			} else {
				class_name = nil
				call_type = nil
			}
			if func_.GetType() != 4 {
				if (options & 1 << 1) == 0 {
					DebugBacktraceGetArgs(call, &arg_array)
				}
			}
		} else {

			/* i know this is kinda ugly, but i'm trying to avoid extra cycles in the main execution loop */

			var build_filename_arg ZendBool = 1
			if ptr.GetFunc() == nil || (ptr.GetFunc().GetCommonType()&1) != 0 || ptr.GetOpline().GetOpcode() != 73 {

				/* can happen when calling eval from a custom sapi */

				function_name = "unknown"
				build_filename_arg = 0
			} else {
				switch ptr.GetOpline().GetExtendedValue() {
				case 1 << 0:
					function_name = "eval"
					build_filename_arg = 0
					break
				case 1 << 1:
					function_name = "include"
					break
				case 1 << 3:
					function_name = "require"
					break
				case 1 << 2:
					function_name = "include_once"
					break
				case 1 << 4:
					function_name = "require_once"
					break
				default:

					/* this can actually happen if you use debug_backtrace() in your error_handler and
					 * you're in the top-scope */

					function_name = "unknown"
					build_filename_arg = 0
					break
				}
			}
			if build_filename_arg != 0 && include_filename != nil {
				var __arr *ZendArray = _zendNewArray(0)
				var __z *Zval = &arg_array
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				AddNextIndexString(&arg_array, (*byte)(include_filename))
			}
			call_type = nil
		}
		ZendPrintf("#%-2d ", indent)
		if class_name != nil {
			ZendWrite(class_name.GetVal(), strlen(class_name.GetVal()))
			ZendWrite(call_type, strlen(call_type))
			if object != nil && func_.GetScope() == nil && object.GetHandlers().GetGetClassName() != ZendStdGetClassName {
				ZendStringReleaseEx(class_name, 0)
			}
		}
		ZendPrintf("%s(", function_name)
		if arg_array.GetType() != 0 {
			DebugPrintBacktraceArgs(&arg_array)
			ZvalPtrDtor(&arg_array)
		}
		if filename != nil {
			ZendPrintf(") called at [%s:%d]\n", filename, lineno)
		} else {
			var prev_call *ZendExecuteData = skip
			var prev *ZendExecuteData = skip.GetPrevExecuteData()
			for prev != nil {
				if prev_call != nil && prev_call.GetFunc() != nil && (prev_call.GetFunc().GetCommonType()&1) != 0 {
					prev = nil
					break
				}
				if prev.GetFunc() != nil && (prev.GetFunc().GetCommonType()&1) == 0 {
					ZendPrintf(") called at [%s:%d]\n", prev.GetFunc().GetOpArray().GetFilename().GetVal(), prev.GetOpline().GetLineno())
					break
				}
				prev_call = prev
				prev = prev.GetPrevExecuteData()
			}
			if prev == nil {
				ZendWrite(")\n", strlen(")\n"))
			}
		}
		include_filename = filename
		call = skip
		ptr = skip.GetPrevExecuteData()
		indent++
	}
}

/* }}} */

func ZendFetchDebugBacktrace(return_value *Zval, skip_last int, options int, limit int) {
	var ptr *ZendExecuteData
	var skip *ZendExecuteData
	var call *ZendExecuteData = nil
	var object *ZendObject
	var lineno int
	var frameno int = 0
	var func_ *ZendFunction
	var function_name *ZendString
	var filename *ZendString
	var include_filename *ZendString = nil
	var stack_frame Zval
	var tmp Zval
	var __arr *ZendArray = _zendNewArray(0)
	var __z *Zval = return_value
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	if !(g.Assign(&ptr, EG.GetCurrentExecuteData())) {
		return
	}
	if ptr.GetFunc() == nil || (ptr.GetFunc().GetCommonType()&1) != 0 {
		call = ptr
		ptr = ptr.GetPrevExecuteData()
	}
	if ptr != nil {
		if skip_last != 0 {

			/* skip debug_backtrace() */

			call = ptr
			ptr = ptr.GetPrevExecuteData()
		} else {

			/* skip "new Exception()" */

			if ptr.GetFunc() != nil && (ptr.GetFunc().GetCommonType()&1) == 0 && ptr.GetOpline().GetOpcode() == 68 {
				call = ptr
				ptr = ptr.GetPrevExecuteData()
			}

			/* skip "new Exception()" */

		}
		if call == nil {
			call = ptr
			ptr = ptr.GetPrevExecuteData()
		}
	}
	for ptr != nil && (limit == 0 || frameno < limit) {
		frameno++
		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = &stack_frame
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		ptr = ZendGeneratorCheckPlaceholderFrame(ptr)
		skip = ptr

		/* skip internal handler */

		if SkipInternalHandler(skip) != 0 {
			skip = skip.GetPrevExecuteData()
		}
		if skip.GetFunc() != nil && (skip.GetFunc().GetCommonType()&1) == 0 {
			filename = skip.GetFunc().GetOpArray().GetFilename()
			if skip.GetOpline().GetOpcode() == 149 {
				if EG.GetOplineBeforeException() != nil {
					lineno = EG.GetOplineBeforeException().GetLineno()
				} else {
					lineno = skip.GetFunc().GetOpArray().GetLineEnd()
				}
			} else {
				lineno = skip.GetOpline().GetLineno()
			}
			var __z *Zval = &tmp
			var __s *ZendString = filename
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
			ZendHashAddNew(stack_frame.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_FILE], &tmp)
			var __z *Zval = &tmp
			__z.GetValue().SetLval(lineno)
			__z.SetTypeInfo(4)
			ZendHashAddNew(stack_frame.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_LINE], &tmp)
		} else {
			var prev_call *ZendExecuteData = skip
			var prev *ZendExecuteData = skip.GetPrevExecuteData()
			for prev != nil {
				if prev_call != nil && prev_call.GetFunc() != nil && (prev_call.GetFunc().GetCommonType()&1) != 0 && (prev_call.GetFunc().GetFnFlags()&1<<18) == 0 {
					break
				}
				if prev.GetFunc() != nil && (prev.GetFunc().GetCommonType()&1) == 0 {
					var __z *Zval = &tmp
					var __s *ZendString = prev.GetFunc().GetOpArray().GetFilename()
					__z.GetValue().SetStr(__s)
					if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
						__z.SetTypeInfo(6)
					} else {
						ZendGcAddref(&__s.gc)
						__z.SetTypeInfo(6 | 1<<0<<8)
					}
					ZendHashAddNew(stack_frame.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_FILE], &tmp)
					var __z *Zval = &tmp
					__z.GetValue().SetLval(prev.GetOpline().GetLineno())
					__z.SetTypeInfo(4)
					ZendHashAddNew(stack_frame.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_LINE], &tmp)
					break
				}
				prev_call = prev
				prev = prev.GetPrevExecuteData()
			}
			filename = nil
		}

		/* $this may be passed into regular internal functions */

		if call != nil && call.GetThis().GetType() == 8 {
			object = call.GetThis().GetValue().GetObj()
		} else {
			object = nil
		}
		if call != nil && call.GetFunc() != nil {
			func_ = call.GetFunc()
			if func_.GetScope() != nil && func_.GetScope().GetTraitAliases() != nil {
				function_name = ZendResolveMethodName(g.CondF(object != nil, func() *ZendClassEntry { return object.GetCe() }, func() *ZendClassEntry { return func_.GetScope() }), func_)
			} else {
				function_name = func_.GetFunctionName()
			}
		} else {
			func_ = nil
			function_name = nil
		}
		if function_name != nil {
			var __z *Zval = &tmp
			var __s *ZendString = function_name
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
			ZendHashAddNew(stack_frame.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_FUNCTION], &tmp)
			if object != nil {
				if func_.GetScope() != nil {
					var __z *Zval = &tmp
					var __s *ZendString = func_.GetScope().GetName()
					__z.GetValue().SetStr(__s)
					if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
						__z.SetTypeInfo(6)
					} else {
						ZendGcAddref(&__s.gc)
						__z.SetTypeInfo(6 | 1<<0<<8)
					}
				} else if object.GetHandlers().GetGetClassName() == ZendStdGetClassName {
					var __z *Zval = &tmp
					var __s *ZendString = object.GetCe().GetName()
					__z.GetValue().SetStr(__s)
					if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
						__z.SetTypeInfo(6)
					} else {
						ZendGcAddref(&__s.gc)
						__z.SetTypeInfo(6 | 1<<0<<8)
					}
				} else {
					var __z *Zval = &tmp
					var __s *ZendString = object.GetHandlers().GetGetClassName()(object)
					__z.GetValue().SetStr(__s)
					if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
						__z.SetTypeInfo(6)
					} else {
						__z.SetTypeInfo(6 | 1<<0<<8)
					}
				}
				ZendHashAddNew(stack_frame.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_CLASS], &tmp)
				if (options & 1 << 0) != 0 {
					var __z *Zval = &tmp
					__z.GetValue().SetObj(object)
					__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
					ZendHashAddNew(stack_frame.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_OBJECT], &tmp)
					ZvalAddrefP(&tmp)
				}
				var __z *Zval = &tmp
				var __s *ZendString = ZendKnownStrings[ZEND_STR_OBJECT_OPERATOR]
				__z.GetValue().SetStr(__s)
				__z.SetTypeInfo(6)
				ZendHashAddNew(stack_frame.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_TYPE], &tmp)
			} else if func_.GetScope() != nil {
				var __z *Zval = &tmp
				var __s *ZendString = func_.GetScope().GetName()
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				ZendHashAddNew(stack_frame.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_CLASS], &tmp)
				var __z *Zval = &tmp
				var __s *ZendString = ZendKnownStrings[ZEND_STR_PAAMAYIM_NEKUDOTAYIM]
				__z.GetValue().SetStr(__s)
				__z.SetTypeInfo(6)
				ZendHashAddNew(stack_frame.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_TYPE], &tmp)
			}
			if (options&1<<1) == 0 && func_.GetType() != 4 {
				DebugBacktraceGetArgs(call, &tmp)
				ZendHashAddNew(stack_frame.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_ARGS], &tmp)
			}
		} else {

			/* i know this is kinda ugly, but i'm trying to avoid extra cycles in the main execution loop */

			var build_filename_arg ZendBool = 1
			var pseudo_function_name *ZendString
			if ptr.GetFunc() == nil || (ptr.GetFunc().GetCommonType()&1) != 0 || ptr.GetOpline().GetOpcode() != 73 {

				/* can happen when calling eval from a custom sapi */

				pseudo_function_name = ZendKnownStrings[ZEND_STR_UNKNOWN]
				build_filename_arg = 0
			} else {
				switch ptr.GetOpline().GetExtendedValue() {
				case 1 << 0:
					pseudo_function_name = ZendKnownStrings[ZEND_STR_EVAL]
					build_filename_arg = 0
					break
				case 1 << 1:
					pseudo_function_name = ZendKnownStrings[ZEND_STR_INCLUDE]
					break
				case 1 << 3:
					pseudo_function_name = ZendKnownStrings[ZEND_STR_REQUIRE]
					break
				case 1 << 2:
					pseudo_function_name = ZendKnownStrings[ZEND_STR_INCLUDE_ONCE]
					break
				case 1 << 4:
					pseudo_function_name = ZendKnownStrings[ZEND_STR_REQUIRE_ONCE]
					break
				default:

					/* this can actually happen if you use debug_backtrace() in your error_handler and
					 * you're in the top-scope */

					pseudo_function_name = ZendKnownStrings[ZEND_STR_UNKNOWN]
					build_filename_arg = 0
					break
				}
			}
			if build_filename_arg != 0 && include_filename != nil {
				var arg_array Zval
				var __arr *ZendArray = _zendNewArray(0)
				var __z *Zval = &arg_array
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)

				/* include_filename always points to the last filename of the last last called-function.
				   if we have called include in the frame above - this is the file we have included.
				*/

				var __z *Zval = &tmp
				var __s *ZendString = include_filename
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				ZendHashNextIndexInsertNew(arg_array.GetValue().GetArr(), &tmp)
				ZendHashAddNew(stack_frame.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_ARGS], &arg_array)
			}
			var __z *Zval = &tmp
			var __s *ZendString = pseudo_function_name
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6)
			ZendHashAddNew(stack_frame.GetValue().GetArr(), ZendKnownStrings[ZEND_STR_FUNCTION], &tmp)
		}
		ZendHashNextIndexInsertNew(return_value.GetValue().GetArr(), &stack_frame)
		include_filename = filename
		call = skip
		ptr = skip.GetPrevExecuteData()
	}
}

/* }}} */

func ZifDebugBacktrace(execute_data *ZendExecuteData, return_value *Zval) {
	var options ZendLong = 1 << 0
	var limit ZendLong = 0
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "|ll", &options, &limit) == FAILURE {
		return
	}
	ZendFetchDebugBacktrace(return_value, 1, options, limit)
}

/* }}} */

func ZifExtensionLoaded(execute_data *ZendExecuteData, return_value *Zval) {
	var extension_name *ZendString
	var lcname *ZendString
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "S", &extension_name) == FAILURE {
		return
	}
	lcname = ZendStringTolowerEx(extension_name, 0)
	if ZendHashExists(&ModuleRegistry, lcname) != 0 {
		return_value.SetTypeInfo(3)
	} else {
		return_value.SetTypeInfo(2)
	}
	ZendStringReleaseEx(lcname, 0)
}

/* }}} */

func ZifGetExtensionFuncs(execute_data *ZendExecuteData, return_value *Zval) {
	var extension_name *ZendString
	var lcname *ZendString
	var array int
	var module *ZendModuleEntry
	var zif *ZendFunction
	if ZendParseParameters(execute_data.GetThis().GetNumArgs(), "S", &extension_name) == FAILURE {
		return
	}
	if strncasecmp(extension_name.GetVal(), "zend", g.SizeOf("\"zend\"")) {
		lcname = ZendStringTolowerEx(extension_name, 0)
		module = ZendHashFindPtr(&ModuleRegistry, lcname)
		ZendStringReleaseEx(lcname, 0)
	} else {
		module = ZendHashStrFindPtr(&ModuleRegistry, "core", g.SizeOf("\"core\"")-1)
	}
	if module == nil {
		return_value.SetTypeInfo(2)
		return
	}
	if module.GetFunctions() != nil {

		/* avoid BC break, if functions list is empty, will return an empty array */

		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = return_value
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		array = 1
	} else {
		array = 0
	}
	for {
		var __ht *HashTable = CG.GetFunctionTable()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			zif = _z.GetValue().GetPtr()
			if zif.GetCommonType() == 1 && zif.GetInternalFunction().GetModule() == module {
				if array == 0 {
					var __arr *ZendArray = _zendNewArray(0)
					var __z *Zval = return_value
					__z.GetValue().SetArr(__arr)
					__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
					array = 1
				}
				AddNextIndexStr(return_value, ZendStringCopy(zif.GetFunctionName()))
			}
		}
		break
	}
	if array == 0 {
		return_value.SetTypeInfo(2)
		return
	}
}

/* }}} */
