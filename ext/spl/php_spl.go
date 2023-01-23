// <<generate>>

package spl

import (
	"sik/core"
	"sik/ext/standard"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/spl/php_spl.h>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define PHP_SPL_H

// # include "php.h"

// # include < stdarg . h >

// #define PHP_SPL_VERSION       PHP_VERSION

// #define phpext_spl_ptr       & spl_module_entry

// #define SPL_API

var ZmShutdownSpl func(type_ int, module_number int) int

var SplGlobals ZendSplGlobals

// #define SPL_G(v) ZEND_MODULE_GLOBALS_ACCESSOR ( spl , v )

// Source: <ext/spl/php_spl.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_ini.h"

// # include "php_main.h"

// # include "ext/standard/info.h"

// # include "php_spl.h"

// # include "spl_functions.h"

// # include "spl_engine.h"

// # include "spl_array.h"

// # include "spl_directory.h"

// # include "spl_iterators.h"

// # include "spl_exceptions.h"

// # include "spl_observer.h"

// # include "spl_dllist.h"

// # include "spl_fixedarray.h"

// # include "spl_heap.h"

// # include "zend_exceptions.h"

// # include "zend_interfaces.h"

// # include "ext/standard/php_mt_rand.h"

// # include "main/snprintf.h"

// #define SPL_DEFAULT_FILE_EXTENSIONS       ".inc,.php"

var SplAutoloadFn *zend.ZendFunction = nil
var SplAutoloadCallFn *zend.ZendFunction = nil

/* {{{ PHP_GINIT_FUNCTION
 */

func ZmGlobalsCtorSpl(spl_globals *ZendSplGlobals) {
	spl_globals.SetAutoloadExtensions(nil)
	spl_globals.SetAutoloadFunctions(nil)
	spl_globals.SetAutoloadRunning(0)
}

/* }}} */

func SplFindCeByName(name *zend.ZendString, autoload zend.ZendBool) *zend.ZendClassEntry {
	var ce *zend.ZendClassEntry
	if autoload == 0 {
		var lc_name *zend.ZendString = zend.ZendStringTolowerEx(name, 0)
		ce = zend.ZendHashFindPtr(zend.EG.class_table, lc_name)
		zend.ZendStringRelease(lc_name)
	} else {
		ce = zend.ZendLookupClass(name)
	}
	if ce == nil {
		core.PhpErrorDocref(nil, 1<<1, "Class %s does not exist%s", name.val, g.Cond(autoload != 0, " and could not be loaded", ""))
		return nil
	}
	return ce
}

/* {{{ proto array class_parents(object instance [, bool autoload = true])
Return an array containing the names of all parent classes */

func ZifClassParents(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var parent_class *zend.ZendClassEntry
	var ce *zend.ZendClassEntry
	var autoload zend.ZendBool = 1
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z|b", &obj, &autoload) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	if obj.u1.v.type_ != 8 && obj.u1.v.type_ != 6 {
		core.PhpErrorDocref(nil, 1<<1, "object or string expected")
		return_value.u1.type_info = 2
		return
	}
	if obj.u1.v.type_ == 6 {
		if nil == g.Assign(&ce, SplFindCeByName(obj.value.str, autoload)) {
			return_value.u1.type_info = 2
			return
		}
	} else {
		ce = obj.value.obj.ce
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	parent_class = ce.parent
	for parent_class != nil {
		SplAddClassName(return_value, parent_class, 0, 0)
		parent_class = parent_class.parent
	}
}

/* }}} */

func ZifClassImplements(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var autoload zend.ZendBool = 1
	var ce *zend.ZendClassEntry
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z|b", &obj, &autoload) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	if obj.u1.v.type_ != 8 && obj.u1.v.type_ != 6 {
		core.PhpErrorDocref(nil, 1<<1, "object or string expected")
		return_value.u1.type_info = 2
		return
	}
	if obj.u1.v.type_ == 6 {
		if nil == g.Assign(&ce, SplFindCeByName(obj.value.str, autoload)) {
			return_value.u1.type_info = 2
			return
		}
	} else {
		ce = obj.value.obj.ce
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	SplAddInterfaces(return_value, ce, 1, 1<<0)
}

/* }}} */

func ZifClassUses(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var autoload zend.ZendBool = 1
	var ce *zend.ZendClassEntry
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z|b", &obj, &autoload) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	if obj.u1.v.type_ != 8 && obj.u1.v.type_ != 6 {
		core.PhpErrorDocref(nil, 1<<1, "object or string expected")
		return_value.u1.type_info = 2
		return
	}
	if obj.u1.v.type_ == 6 {
		if nil == g.Assign(&ce, SplFindCeByName(obj.value.str, autoload)) {
			return_value.u1.type_info = 2
			return
		}
	} else {
		ce = obj.value.obj.ce
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	SplAddTraits(return_value, ce, 1, 1<<1)
}

/* }}} */

// #define SPL_ADD_CLASS(class_name,z_list,sub,allow,ce_flags) spl_add_classes ( spl_ce_ ## class_name , z_list , sub , allow , ce_flags )

// #define SPL_LIST_CLASSES(z_list,sub,allow,ce_flags) SPL_ADD_CLASS ( AppendIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( ArrayIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( ArrayObject , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( BadFunctionCallException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( BadMethodCallException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( CachingIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( CallbackFilterIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( DirectoryIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( DomainException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( EmptyIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( FilesystemIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( FilterIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( GlobIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( InfiniteIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( InvalidArgumentException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( IteratorIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( LengthException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( LimitIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( LogicException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( MultipleIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( NoRewindIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( OuterIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( OutOfBoundsException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( OutOfRangeException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( OverflowException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( ParentIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RangeException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveArrayIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveCachingIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveCallbackFilterIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveDirectoryIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveFilterIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveIteratorIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveRegexIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RecursiveTreeIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RegexIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( RuntimeException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SeekableIterator , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplDoublyLinkedList , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplFileInfo , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplFileObject , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplFixedArray , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplHeap , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplMinHeap , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplMaxHeap , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplObjectStorage , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplObserver , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplPriorityQueue , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplQueue , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplStack , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplSubject , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( SplTempFileObject , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( UnderflowException , z_list , sub , allow , ce_flags ) ; SPL_ADD_CLASS ( UnexpectedValueException , z_list , sub , allow , ce_flags ) ;

/* {{{ proto array spl_classes()
Return an array containing the names of all clsses and interfaces defined in SPL */

func ZifSplClasses(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
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

/* }}} */

func SplAutoload(class_name *zend.ZendString, lc_name *zend.ZendString, ext *byte, ext_len int) int {
	var class_file *byte
	var class_file_len int
	var dummy zend.Zval
	var file_handle zend.ZendFileHandle
	var new_op_array *zend.ZendOpArray
	var result zend.Zval
	var ret int
	class_file_len = int(zend.ZendSpprintf(&class_file, 0, "%s%.*s", lc_name.val, ext_len, ext))
	ret = core.PhpStreamOpenForZendEx(class_file, &file_handle, 0x1|0x80)
	if ret == zend.SUCCESS {
		var opened_path *zend.ZendString
		if file_handle.opened_path == nil {
			file_handle.opened_path = zend.ZendStringInit(class_file, class_file_len, 0)
		}
		opened_path = zend.ZendStringCopy(file_handle.opened_path)
		&dummy.u1.type_info = 1
		if zend.ZendHashAdd(&zend.EG.included_files, opened_path, &dummy) != nil {
			new_op_array = zend.ZendCompileFile(&file_handle, 1<<3)
			zend.ZendDestroyFileHandle(&file_handle)
		} else {
			new_op_array = nil
			zend.ZendFileHandleDtor(&file_handle)
		}
		zend.ZendStringReleaseEx(opened_path, 0)
		if new_op_array != nil {
			&result.u1.type_info = 0
			zend.ZendExecute(new_op_array, &result)
			zend.DestroyOpArray(new_op_array)
			zend._efree(new_op_array)
			if zend.EG.exception == nil {
				zend.ZvalPtrDtor(&result)
			}
			zend._efree(class_file)
			return zend.ZendHashExists(zend.EG.class_table, lc_name)
		}
	}
	zend._efree(class_file)
	return 0
}

/* {{{ proto void spl_autoload(string class_name [, string file_extensions])
Default implementation for __autoload() */

func ZifSplAutoload(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var pos_len int
	var pos1_len int
	var pos *byte
	var pos1 *byte
	var class_name *zend.ZendString
	var lc_name *zend.ZendString
	var file_exts *zend.ZendString = SplGlobals.GetAutoloadExtensions()
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "S|S", &class_name, &file_exts) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	if file_exts == nil {
		pos = ".inc,.php"
		pos_len = g.SizeOf("SPL_DEFAULT_FILE_EXTENSIONS") - 1
	} else {
		pos = file_exts.val
		pos_len = int(file_exts.len_)
	}
	lc_name = zend.ZendStringTolowerEx(class_name, 0)
	for pos != nil && (*pos) && zend.EG.exception == nil {
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
	zend.ZendStringRelease(lc_name)
}

/* {{{ proto string spl_autoload_extensions([string file_extensions])
Register and return default file extensions for spl_autoload */

func ZifSplAutoloadExtensions(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var file_exts *zend.ZendString = nil
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|S", &file_exts) == zend.FAILURE {
		return
	}
	if file_exts != nil {
		if SplGlobals.GetAutoloadExtensions() != nil {
			zend.ZendStringReleaseEx(SplGlobals.GetAutoloadExtensions(), 0)
		}
		SplGlobals.SetAutoloadExtensions(zend.ZendStringCopy(file_exts))
	}
	if SplGlobals.GetAutoloadExtensions() == nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(".inc,.php", g.SizeOf("SPL_DEFAULT_FILE_EXTENSIONS")-1, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		zend.ZendStringAddref(SplGlobals.GetAutoloadExtensions())
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = SplGlobals.GetAutoloadExtensions()
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	}
}
func AutoloadFuncInfoDtor(element *zend.Zval) {
	var alfi *AutoloadFuncInfo = (*AutoloadFuncInfo)(element.value.ptr)
	if alfi.obj.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&alfi.obj)
	}
	if alfi.GetFuncPtr() != nil && (alfi.GetFuncPtr().common.fn_flags&1<<18) != 0 {
		zend.ZendStringReleaseEx(alfi.GetFuncPtr().common.function_name, 0)
		if alfi.GetFuncPtr() == &zend.EG.trampoline {
			zend.EG.trampoline.common.function_name = nil
		} else {
			zend._efree(alfi.GetFuncPtr())
		}
	}
	if alfi.closure.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&alfi.closure)
	}
	zend._efree(alfi)
}

/* {{{ proto void spl_autoload_call(string class_name)
Try all registered autoload function to load the requested class */

func ZifSplAutoloadCall(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var class_name *zend.Zval
	var retval zend.Zval
	var lc_name *zend.ZendString
	var func_name *zend.ZendString
	var alfi *AutoloadFuncInfo
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &class_name) == zend.FAILURE || class_name.u1.v.type_ != 6 {
		return
	}
	if SplGlobals.GetAutoloadFunctions() != nil {
		var pos zend.HashPosition
		var num_idx zend.ZendUlong
		var func_ *zend.ZendFunction
		var fci zend.ZendFcallInfo
		var fcic zend.ZendFcallInfoCache
		var called_scope *zend.ZendClassEntry = zend.ZendGetCalledScope(execute_data)
		var l_autoload_running int = SplGlobals.GetAutoloadRunning()
		SplGlobals.SetAutoloadRunning(1)
		lc_name = zend.ZendStringTolowerEx(class_name.value.str, 0)
		fci.size = g.SizeOf("fci")
		fci.retval = &retval
		fci.param_count = 1
		fci.params = class_name
		fci.no_separation = 1
		&fci.function_name.u1.type_info = 0
		zend.ZendHashInternalPointerResetEx(SplGlobals.GetAutoloadFunctions(), &pos)
		for zend.ZendHashGetCurrentKeyEx(SplGlobals.GetAutoloadFunctions(), &func_name, &num_idx, &pos) == 1 {
			alfi = zend.ZendHashGetCurrentDataPtrEx(SplGlobals.GetAutoloadFunctions(), &pos)
			func_ = alfi.GetFuncPtr()
			if (func_.common.fn_flags & 1 << 18) != 0 {
				func_ = zend._emalloc(g.SizeOf("zend_op_array"))
				memcpy(func_, alfi.GetFuncPtr(), g.SizeOf("zend_op_array"))
				zend.ZendStringAddref(func_.op_array.function_name)
			}
			&retval.u1.type_info = 0
			fcic.function_handler = func_
			if alfi.obj.u1.v.type_ == 0 {
				fci.object = nil
				fcic.object = nil
				if alfi.GetCe() != nil && (called_scope == nil || zend.InstanceofFunction(called_scope, alfi.GetCe()) == 0) {
					fcic.called_scope = alfi.GetCe()
				} else {
					fcic.called_scope = called_scope
				}
			} else {
				fci.object = alfi.obj.value.obj
				fcic.object = alfi.obj.value.obj
				fcic.called_scope = alfi.obj.value.obj.ce
			}
			zend.ZendCallFunction(&fci, &fcic)
			zend.ZvalPtrDtor(&retval)
			if zend.EG.exception != nil {
				break
			}
			if pos+1 == SplGlobals.GetAutoloadFunctions().nNumUsed || zend.ZendHashExists(zend.EG.class_table, lc_name) != 0 {
				break
			}
			zend.ZendHashMoveForwardEx(SplGlobals.GetAutoloadFunctions(), &pos)
		}
		zend.ZendStringReleaseEx(lc_name, 0)
		SplGlobals.SetAutoloadRunning(l_autoload_running)
	} else {

		/* do not use or overwrite &EG(autoload_func) here */

		var fcall_info zend.ZendFcallInfo
		var fcall_cache zend.ZendFcallInfoCache
		&retval.u1.type_info = 0
		fcall_info.size = g.SizeOf("fcall_info")
		&fcall_info.function_name.u1.type_info = 0
		fcall_info.retval = &retval
		fcall_info.param_count = 1
		fcall_info.params = class_name
		fcall_info.object = nil
		fcall_info.no_separation = 1
		fcall_cache.function_handler = SplAutoloadFn
		fcall_cache.called_scope = nil
		fcall_cache.object = nil
		zend.ZendCallFunction(&fcall_info, &fcall_cache)
		zend.ZvalPtrDtor(&retval)
	}
}

// #define HT_MOVE_TAIL_TO_HEAD(ht) do { Bucket tmp = ( ht ) -> arData [ ( ht ) -> nNumUsed - 1 ] ; memmove ( ( ht ) -> arData + 1 , ( ht ) -> arData , sizeof ( Bucket ) * ( ( ht ) -> nNumUsed - 1 ) ) ; ( ht ) -> arData [ 0 ] = tmp ; zend_hash_rehash ( ht ) ; } while ( 0 )

/* {{{ proto bool spl_autoload_register([mixed autoload_function [, bool throw [, bool prepend]]])
Register given function as __autoload() implementation */

func ZifSplAutoloadRegister(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var func_name *zend.ZendString
	var error *byte = nil
	var lc_name *zend.ZendString
	var zcallable *zend.Zval = nil
	var do_throw zend.ZendBool = 1
	var prepend zend.ZendBool = 0
	var spl_func_ptr *zend.ZendFunction
	var alfi AutoloadFuncInfo
	var obj_ptr *zend.ZendObject
	var fcc zend.ZendFcallInfoCache
	if zend.ZendParseParametersEx(1<<1, execute_data.This.u2.num_args, "|zbb", &zcallable, &do_throw, &prepend) == zend.FAILURE {
		return
	}
	if execute_data.This.u2.num_args != 0 {
		if zend.ZendIsCallableEx(zcallable, nil, 1<<2, &func_name, &fcc, &error) == 0 {
			alfi.SetCe(fcc.calling_scope)
			alfi.SetFuncPtr(fcc.function_handler)
			obj_ptr = fcc.object
			if zcallable.u1.v.type_ == 7 {
				if obj_ptr == nil && alfi.GetFuncPtr() != nil && (alfi.GetFuncPtr().common.fn_flags&1<<4) == 0 {
					if do_throw != 0 {
						zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Passed array specifies a non static method but no object (%s)", error)
					}
					if error != nil {
						zend._efree(error)
					}
					zend.ZendStringReleaseEx(func_name, 0)
					return_value.u1.type_info = 2
					return
				} else if do_throw != 0 {
					zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Passed array does not specify %s %smethod (%s)", g.Cond(alfi.GetFuncPtr() != nil, "a callable", "an existing"), g.Cond(obj_ptr == nil, "static ", ""), error)
				}
				if error != nil {
					zend._efree(error)
				}
				zend.ZendStringReleaseEx(func_name, 0)
				return_value.u1.type_info = 2
				return
			} else if zcallable.u1.v.type_ == 6 {
				if do_throw != 0 {
					zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Function '%s' not %s (%s)", func_name.val, g.Cond(alfi.GetFuncPtr() != nil, "callable", "found"), error)
				}
				if error != nil {
					zend._efree(error)
				}
				zend.ZendStringReleaseEx(func_name, 0)
				return_value.u1.type_info = 2
				return
			} else {
				if do_throw != 0 {
					zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Illegal value passed (%s)", error)
				}
				if error != nil {
					zend._efree(error)
				}
				zend.ZendStringReleaseEx(func_name, 0)
				return_value.u1.type_info = 2
				return
			}
		} else if fcc.function_handler.type_ == 1 && fcc.function_handler.internal_function.handler == ZifSplAutoloadCall {
			if do_throw != 0 {
				zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Function spl_autoload_call() cannot be registered")
			}
			if error != nil {
				zend._efree(error)
			}
			zend.ZendStringReleaseEx(func_name, 0)
			return_value.u1.type_info = 2
			return
		}
		alfi.SetCe(fcc.calling_scope)
		alfi.SetFuncPtr(fcc.function_handler)
		obj_ptr = fcc.object
		if error != nil {
			zend._efree(error)
		}
		if zcallable.u1.v.type_ == 8 {
			var _z1 *zend.Zval = &alfi.closure
			var _z2 *zend.Zval = zcallable
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
			lc_name = zend.ZendStringAlloc(func_name.len_+g.SizeOf("uint32_t"), 0)
			zend.ZendStrTolowerCopy(lc_name.val, func_name.val, func_name.len_)
			memcpy(lc_name.val+func_name.len_, &(zcallable.value.obj).handle, g.SizeOf("uint32_t"))
			lc_name.val[lc_name.len_] = '0'
		} else {
			&alfi.closure.u1.type_info = 0

			/* Skip leading \ */

			if func_name.val[0] == '\\' {
				lc_name = zend.ZendStringAlloc(func_name.len_-1, 0)
				zend.ZendStrTolowerCopy(lc_name.val, func_name.val+1, func_name.len_-1)
			} else {
				lc_name = zend.ZendStringTolowerEx(func_name, 0)
			}

			/* Skip leading \ */

		}
		zend.ZendStringReleaseEx(func_name, 0)
		if SplGlobals.GetAutoloadFunctions() != nil && zend.ZendHashExists(SplGlobals.GetAutoloadFunctions(), lc_name) != 0 {
			if alfi.closure.u1.v.type_ != 0 {
				zend.ZvalDelrefP(&alfi.closure)
			}
			goto skip
		}
		if obj_ptr != nil && (alfi.GetFuncPtr().common.fn_flags&1<<4) == 0 {

			/* add object id to the hash to ensure uniqueness, for more reference look at bug #40091 */

			lc_name = zend.ZendStringExtend(lc_name, lc_name.len_+g.SizeOf("uint32_t"), 0)
			memcpy(lc_name.val+lc_name.len_-g.SizeOf("uint32_t"), &obj_ptr.handle, g.SizeOf("uint32_t"))
			lc_name.val[lc_name.len_] = '0'
			var __z *zend.Zval = &alfi.obj
			__z.value.obj = obj_ptr
			__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
			zend.ZvalAddrefP(&(alfi.GetObj()))
		} else {
			&alfi.obj.u1.type_info = 0
		}
		if SplGlobals.GetAutoloadFunctions() == nil {
			SplGlobals.SetAutoloadFunctions((*zend.HashTable)(zend._emalloc(g.SizeOf("HashTable"))))
			zend._zendHashInit(SplGlobals.GetAutoloadFunctions(), 1, AutoloadFuncInfoDtor, 0)
		}
		spl_func_ptr = SplAutoloadFn
		if zend.EG.autoload_func == spl_func_ptr {
			var spl_alfi AutoloadFuncInfo
			spl_alfi.SetFuncPtr(spl_func_ptr)
			&spl_alfi.obj.u1.type_info = 0
			&spl_alfi.closure.u1.type_info = 0
			spl_alfi.SetCe(nil)
			zend.ZendHashAddMem(SplGlobals.GetAutoloadFunctions(), SplAutoloadFn.common.function_name, &spl_alfi, g.SizeOf("autoload_func_info"))
			if prepend != 0 && SplGlobals.GetAutoloadFunctions().nNumOfElements > 1 {

				/* Move the newly created element to the head of the hashtable */

				var tmp zend.Bucket = SplGlobals.GetAutoloadFunctions().arData[SplGlobals.GetAutoloadFunctions().nNumUsed-1]
				memmove(SplGlobals.GetAutoloadFunctions().arData+1, SplGlobals.GetAutoloadFunctions().arData, g.SizeOf("Bucket")*(SplGlobals.GetAutoloadFunctions().nNumUsed-1))
				SplGlobals.GetAutoloadFunctions().arData[0] = tmp
				zend.ZendHashRehash(SplGlobals.GetAutoloadFunctions())

				/* Move the newly created element to the head of the hashtable */

			}
		}
		if alfi.GetFuncPtr() == &zend.EG.trampoline {
			var copy *zend.ZendFunction = zend._emalloc(g.SizeOf("zend_op_array"))
			memcpy(copy, alfi.GetFuncPtr(), g.SizeOf("zend_op_array"))
			alfi.GetFuncPtr().common.function_name = nil
			alfi.SetFuncPtr(copy)
		}
		if zend.ZendHashAddMem(SplGlobals.GetAutoloadFunctions(), lc_name, &alfi, g.SizeOf("autoload_func_info")) == nil {
			if obj_ptr != nil && (alfi.GetFuncPtr().common.fn_flags&1<<4) == 0 {
				zend.ZvalDelrefP(&(alfi.GetObj()))
			}
			if alfi.closure.u1.v.type_ != 0 {
				zend.ZvalDelrefP(&(alfi.GetClosure()))
			}
			if (alfi.GetFuncPtr().common.fn_flags & 1 << 18) != 0 {
				zend.ZendStringReleaseEx(alfi.GetFuncPtr().common.function_name, 0)
				if alfi.GetFuncPtr() == &zend.EG.trampoline {
					zend.EG.trampoline.common.function_name = nil
				} else {
					zend._efree(alfi.GetFuncPtr())
				}
			}
		}
		if prepend != 0 && SplGlobals.GetAutoloadFunctions().nNumOfElements > 1 {

			/* Move the newly created element to the head of the hashtable */

			var tmp zend.Bucket = SplGlobals.GetAutoloadFunctions().arData[SplGlobals.GetAutoloadFunctions().nNumUsed-1]
			memmove(SplGlobals.GetAutoloadFunctions().arData+1, SplGlobals.GetAutoloadFunctions().arData, g.SizeOf("Bucket")*(SplGlobals.GetAutoloadFunctions().nNumUsed-1))
			SplGlobals.GetAutoloadFunctions().arData[0] = tmp
			zend.ZendHashRehash(SplGlobals.GetAutoloadFunctions())

			/* Move the newly created element to the head of the hashtable */

		}
	skip:
		zend.ZendStringReleaseEx(lc_name, 0)
	}
	if SplGlobals.GetAutoloadFunctions() != nil {
		zend.EG.autoload_func = SplAutoloadCallFn
	} else {
		zend.EG.autoload_func = SplAutoloadFn
	}
	return_value.u1.type_info = 3
	return
}

/* {{{ proto bool spl_autoload_unregister(mixed autoload_function)
Unregister given function as __autoload() implementation */

func ZifSplAutoloadUnregister(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var func_name *zend.ZendString = nil
	var error *byte = nil
	var lc_name *zend.ZendString
	var zcallable *zend.Zval
	var success int = zend.FAILURE
	var spl_func_ptr *zend.ZendFunction
	var obj_ptr *zend.ZendObject
	var fcc zend.ZendFcallInfoCache
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &zcallable) == zend.FAILURE {
		return
	}
	if zend.ZendIsCallableEx(zcallable, nil, 1<<0, &func_name, &fcc, &error) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Unable to unregister invalid function (%s)", error)
		if error != nil {
			zend._efree(error)
		}
		if func_name != nil {
			zend.ZendStringReleaseEx(func_name, 0)
		}
		return_value.u1.type_info = 2
		return
	}
	obj_ptr = fcc.object
	if error != nil {
		zend._efree(error)
	}
	if zcallable.u1.v.type_ == 8 {
		lc_name = zend.ZendStringAlloc(func_name.len_+g.SizeOf("uint32_t"), 0)
		zend.ZendStrTolowerCopy(lc_name.val, func_name.val, func_name.len_)
		memcpy(lc_name.val+func_name.len_, &(zcallable.value.obj).handle, g.SizeOf("uint32_t"))
		lc_name.val[lc_name.len_] = '0'
	} else {

		/* Skip leading \ */

		if func_name.val[0] == '\\' {
			lc_name = zend.ZendStringAlloc(func_name.len_-1, 0)
			zend.ZendStrTolowerCopy(lc_name.val, func_name.val+1, func_name.len_-1)
		} else {
			lc_name = zend.ZendStringTolowerEx(func_name, 0)
		}

		/* Skip leading \ */

	}
	zend.ZendStringReleaseEx(func_name, 0)
	if SplGlobals.GetAutoloadFunctions() != nil {
		if zend.ZendStringEquals(lc_name, SplAutoloadCallFn.common.function_name) != 0 {

			/* remove all */

			if SplGlobals.GetAutoloadRunning() == 0 {
				zend.ZendHashDestroy(SplGlobals.GetAutoloadFunctions())
				zend._efree(SplGlobals.GetAutoloadFunctions())
				SplGlobals.SetAutoloadFunctions(nil)
				zend.EG.autoload_func = nil
			} else {
				zend.ZendHashClean(SplGlobals.GetAutoloadFunctions())
			}
			success = zend.SUCCESS
		} else {

			/* remove specific */

			success = zend.ZendHashDel(SplGlobals.GetAutoloadFunctions(), lc_name)
			if success != zend.SUCCESS && obj_ptr != nil {
				lc_name = zend.ZendStringExtend(lc_name, lc_name.len_+g.SizeOf("uint32_t"), 0)
				memcpy(lc_name.val+lc_name.len_-g.SizeOf("uint32_t"), &obj_ptr.handle, g.SizeOf("uint32_t"))
				lc_name.val[lc_name.len_] = '0'
				success = zend.ZendHashDel(SplGlobals.GetAutoloadFunctions(), lc_name)
			}
		}
	} else if zend.ZendStringEquals(lc_name, SplAutoloadFn.common.function_name) != 0 {

		/* register single spl_autoload() */

		spl_func_ptr = SplAutoloadFn
		if zend.EG.autoload_func == spl_func_ptr {
			success = zend.SUCCESS
			zend.EG.autoload_func = nil
		}
	}
	zend.ZendStringReleaseEx(lc_name, 0)
	if success == zend.SUCCESS {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto false|array spl_autoload_functions()
Return all registered __autoload() functionns */

func ZifSplAutoloadFunctions(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var fptr *zend.ZendFunction
	var alfi *AutoloadFuncInfo
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if zend.EG.autoload_func == nil {
		if g.Assign(&fptr, zend.ZendHashFindPtr(zend.EG.function_table, zend.ZendKnownStrings[zend.ZEND_STR_MAGIC_AUTOLOAD])) {
			var tmp zend.Zval
			var __arr *zend.ZendArray = zend._zendNewArray(0)
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			var __z *zend.Zval = &tmp
			var __s *zend.ZendString = zend.ZendKnownStrings[zend.ZEND_STR_MAGIC_AUTOLOAD]
			__z.value.str = __s
			if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
				__z.u1.type_info = 6
			} else {
				zend.ZendGcAddref(&__s.gc)
				__z.u1.type_info = 6 | 1<<0<<8
			}
			zend.ZendHashNextIndexInsertNew(return_value.value.arr, &tmp)
			return
		}
		return_value.u1.type_info = 2
		return
	}
	fptr = SplAutoloadCallFn
	if zend.EG.autoload_func == fptr {
		var key *zend.ZendString
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		for {
			var __ht *zend.HashTable = SplGlobals.GetAutoloadFunctions()
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				key = _p.key
				alfi = _z.value.ptr
				if alfi.closure.u1.v.type_ != 0 {
					zend.ZvalAddrefP(&(alfi.GetClosure()))
					zend.AddNextIndexZval(return_value, &alfi.closure)
				} else if alfi.GetFuncPtr().common.scope != nil {
					var tmp zend.Zval
					var __arr *zend.ZendArray = zend._zendNewArray(0)
					var __z *zend.Zval = &tmp
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
					if alfi.obj.u1.v.type_ != 0 {
						zend.ZvalAddrefP(&(alfi.GetObj()))
						zend.AddNextIndexZval(&tmp, &alfi.obj)
					} else {
						zend.AddNextIndexStr(&tmp, zend.ZendStringCopy(alfi.GetCe().name))
					}
					zend.AddNextIndexStr(&tmp, zend.ZendStringCopy(alfi.GetFuncPtr().common.function_name))
					zend.AddNextIndexZval(return_value, &tmp)
				} else {
					if strncmp(alfi.GetFuncPtr().common.function_name.val, "__lambda_func", g.SizeOf("\"__lambda_func\"")-1) {
						zend.AddNextIndexStr(return_value, zend.ZendStringCopy(alfi.GetFuncPtr().common.function_name))
					} else {
						zend.AddNextIndexStr(return_value, zend.ZendStringCopy(key))
					}
				}
			}
			break
		}
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	zend.AddNextIndexStr(return_value, zend.ZendStringCopy(zend.EG.autoload_func.common.function_name))
}

/* {{{ proto string spl_object_hash(object obj)
Return hash id for given object */

func ZifSplObjectHash(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "o", &obj) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpSplObjectHash(obj)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func ZifSplObjectId(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgObject(_arg, &obj, nil, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_OBJECT
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	var __z *zend.Zval = return_value
	__z.value.lval = zend_long(obj.value.obj).handle
	__z.u1.type_info = 4
	return
}

/* }}} */

func PhpSplObjectHash(obj *zend.Zval) *zend.ZendString {
	var hash_handle intPtr
	var hash_handlers intPtr
	if SplGlobals.GetHashMaskInit() == 0 {
		SplGlobals.SetHashMaskHandle(intptr_t(standard.PhpMtRand() >> 1))
		SplGlobals.SetHashMaskHandlers(intptr_t(standard.PhpMtRand() >> 1))
		SplGlobals.SetHashMaskInit(1)
	}
	hash_handle = SplGlobals.GetHashMaskHandle() ^ intptr_t(obj.value.obj).handle
	hash_handlers = SplGlobals.GetHashMaskHandlers()
	return zend.ZendStrpprintf(32, "%016zx%016zx", hash_handle, hash_handlers)
}

/* }}} */

func SplBuildClassListString(entry *zend.Zval, list **byte) {
	var res *byte
	zend.ZendSpprintf(&res, 0, "%s, %s", *list, entry.value.str.val)
	zend._efree(*list)
	*list = res
}

/* {{{ PHP_MINFO(spl)
 */

func ZmInfoSpl(zend_module *zend.ZendModuleEntry) {
	var list zend.Zval
	var zv *zend.Zval
	var strg *byte
	standard.PhpInfoPrintTableStart()
	standard.PhpInfoPrintTableHeader(2, "SPL support", "enabled")
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = &list
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	SplAddClasses(spl_ce_AppendIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_ArrayIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_ArrayObject, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_BadFunctionCallException, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_BadMethodCallException, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_CachingIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_CallbackFilterIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_DirectoryIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_DomainException, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_EmptyIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_FilesystemIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_FilterIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_GlobIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_InfiniteIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_InvalidArgumentException, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_IteratorIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_LengthException, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_LimitIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_LogicException, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_MultipleIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_NoRewindIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_OuterIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_OutOfBoundsException, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_OutOfRangeException, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_OverflowException, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_ParentIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_RangeException, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_RecursiveArrayIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_RecursiveCachingIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_RecursiveCallbackFilterIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_RecursiveDirectoryIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_RecursiveFilterIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_RecursiveIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_RecursiveIteratorIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_RecursiveRegexIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_RecursiveTreeIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_RegexIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_RuntimeException, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SeekableIterator, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplDoublyLinkedList, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplFileInfo, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplFileObject, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplFixedArray, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplHeap, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplMinHeap, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplMaxHeap, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplObjectStorage, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplObserver, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplPriorityQueue, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplQueue, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplStack, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplSubject, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_SplTempFileObject, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_UnderflowException, &list, 0, 1, 1<<0)
	SplAddClasses(spl_ce_UnexpectedValueException, &list, 0, 1, 1<<0)
	strg = zend._estrdup("")
	for {
		var __ht *zend.HashTable = &list.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			zv = _z
			SplBuildClassListString(zv, &strg)
		}
		break
	}
	zend.ZendArrayDestroy(list.value.arr)
	standard.PhpInfoPrintTableRow(2, "Interfaces", strg+2)
	zend._efree(strg)
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = &list
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	SplAddClasses(spl_ce_AppendIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_ArrayIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_ArrayObject, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_BadFunctionCallException, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_BadMethodCallException, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_CachingIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_CallbackFilterIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_DirectoryIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_DomainException, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_EmptyIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_FilesystemIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_FilterIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_GlobIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_InfiniteIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_InvalidArgumentException, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_IteratorIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_LengthException, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_LimitIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_LogicException, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_MultipleIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_NoRewindIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_OuterIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_OutOfBoundsException, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_OutOfRangeException, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_OverflowException, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_ParentIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_RangeException, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_RecursiveArrayIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_RecursiveCachingIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_RecursiveCallbackFilterIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_RecursiveDirectoryIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_RecursiveFilterIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_RecursiveIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_RecursiveIteratorIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_RecursiveRegexIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_RecursiveTreeIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_RegexIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_RuntimeException, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SeekableIterator, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplDoublyLinkedList, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplFileInfo, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplFileObject, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplFixedArray, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplHeap, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplMinHeap, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplMaxHeap, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplObjectStorage, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplObserver, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplPriorityQueue, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplQueue, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplStack, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplSubject, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_SplTempFileObject, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_UnderflowException, &list, 0, -1, 1<<0)
	SplAddClasses(spl_ce_UnexpectedValueException, &list, 0, -1, 1<<0)
	strg = zend._estrdup("")
	for {
		var __ht *zend.HashTable = &list.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			zv = _z
			SplBuildClassListString(zv, &strg)
		}
		break
	}
	zend.ZendArrayDestroy(list.value.arr)
	standard.PhpInfoPrintTableRow(2, "Classes", strg+2)
	zend._efree(strg)
	standard.PhpInfoPrintTableEnd()
}

/* }}} */

var ArginfoIteratorToArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{"iterator", zend.ZendType("Traversable"), 0, 0},
	{"use_keys", 0, 0, 0},
}
var ArginfoIterator []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, 0, 0},
	{"iterator", zend.ZendType("Traversable"), 0, 0},
}
var ArginfoIteratorApply []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(2)), 0, 0, 0},
	{"iterator", zend.ZendType("Traversable"), 0, 0},
	{"function", 0, 0, 0},
	{"args", 7<<2 | g.Cond(true, 0x1, 0x0), 0, 0},
}
var ArginfoClassParents []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"instance", 0, 0, 0}, {"autoload", 0, 0, 0}}
var ArginfoClassImplements []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"what", 0, 0, 0}, {"autoload", 0, 0, 0}}
var ArginfoClassUses []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"what", 0, 0, 0}, {"autoload", 0, 0, 0}}
var ArginfoSplClasses []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoSplAutoloadFunctions []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoSplAutoload []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"class_name", 0, 0, 0}, {"file_extensions", 0, 0, 0}}
var ArginfoSplAutoloadExtensions []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"file_extensions", 0, 0, 0}}
var ArginfoSplAutoloadCall []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"class_name", 0, 0, 0}}
var ArginfoSplAutoloadRegister []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"autoload_function", 0, 0, 0}, {"throw", 0, 0, 0}, {"prepend", 0, 0, 0}}
var ArginfoSplAutoloadUnregister []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"autoload_function", 0, 0, 0}}
var ArginfoSplObjectHash []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"obj", 0, 0, 0}}
var ArginfoSplObjectId []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"obj", 0, 0, 0}}

/* }}} */

var SplFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"spl_classes",
		ZifSplClasses,
		ArginfoSplClasses,
		uint32(g.SizeOf("arginfo_spl_classes")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"spl_autoload",
		ZifSplAutoload,
		ArginfoSplAutoload,
		uint32(g.SizeOf("arginfo_spl_autoload")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"spl_autoload_extensions",
		ZifSplAutoloadExtensions,
		ArginfoSplAutoloadExtensions,
		uint32(g.SizeOf("arginfo_spl_autoload_extensions")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"spl_autoload_register",
		ZifSplAutoloadRegister,
		ArginfoSplAutoloadRegister,
		uint32(g.SizeOf("arginfo_spl_autoload_register")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"spl_autoload_unregister",
		ZifSplAutoloadUnregister,
		ArginfoSplAutoloadUnregister,
		uint32(g.SizeOf("arginfo_spl_autoload_unregister")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"spl_autoload_functions",
		ZifSplAutoloadFunctions,
		ArginfoSplAutoloadFunctions,
		uint32(g.SizeOf("arginfo_spl_autoload_functions")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"spl_autoload_call",
		ZifSplAutoloadCall,
		ArginfoSplAutoloadCall,
		uint32(g.SizeOf("arginfo_spl_autoload_call")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"class_parents",
		ZifClassParents,
		ArginfoClassParents,
		uint32(g.SizeOf("arginfo_class_parents")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"class_implements",
		ZifClassImplements,
		ArginfoClassImplements,
		uint32(g.SizeOf("arginfo_class_implements")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"class_uses",
		ZifClassUses,
		ArginfoClassUses,
		uint32(g.SizeOf("arginfo_class_uses")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"spl_object_hash",
		ZifSplObjectHash,
		ArginfoSplObjectHash,
		uint32(g.SizeOf("arginfo_spl_object_hash")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"spl_object_id",
		ZifSplObjectId,
		ArginfoSplObjectId,
		uint32(g.SizeOf("arginfo_spl_object_id")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"iterator_to_array",
		ZifIteratorToArray,
		ArginfoIteratorToArray,
		uint32(g.SizeOf("arginfo_iterator_to_array")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"iterator_count",
		ZifIteratorCount,
		ArginfoIterator,
		uint32(g.SizeOf("arginfo_iterator")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"iterator_apply",
		ZifIteratorApply,
		ArginfoIteratorApply,
		uint32(g.SizeOf("arginfo_iterator_apply")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

func ZmStartupSpl(type_ int, module_number int) int {
	ZmStartupSplExceptions(type_, module_number)
	ZmStartupSplIterators(type_, module_number)
	ZmStartupSplArray(type_, module_number)
	ZmStartupSplDirectory(type_, module_number)
	ZmStartupSplDllist(type_, module_number)
	ZmStartupSplHeap(type_, module_number)
	ZmStartupSplFixedarray(type_, module_number)
	ZmStartupSplObserver(type_, module_number)
	SplAutoloadFn = zend.ZendHashStrFindPtr(zend.CG.function_table, "spl_autoload", g.SizeOf("\"spl_autoload\"")-1)
	SplAutoloadCallFn = zend.ZendHashStrFindPtr(zend.CG.function_table, "spl_autoload_call", g.SizeOf("\"spl_autoload_call\"")-1)
	r.Assert(SplAutoloadFn != nil && SplAutoloadCallFn != nil)
	return zend.SUCCESS
}

/* }}} */

func ZmActivateSpl(type_ int, module_number int) int {
	SplGlobals.SetAutoloadExtensions(nil)
	SplGlobals.SetAutoloadFunctions(nil)
	SplGlobals.SetHashMaskInit(0)
	return zend.SUCCESS
}
func ZmDeactivateSpl(type_ int, module_number int) int {
	if SplGlobals.GetAutoloadExtensions() != nil {
		zend.ZendStringReleaseEx(SplGlobals.GetAutoloadExtensions(), 0)
		SplGlobals.SetAutoloadExtensions(nil)
	}
	if SplGlobals.GetAutoloadFunctions() != nil {
		zend.ZendHashDestroy(SplGlobals.GetAutoloadFunctions())
		zend._efree(SplGlobals.GetAutoloadFunctions())
		SplGlobals.SetAutoloadFunctions(nil)
	}
	if SplGlobals.GetHashMaskInit() != 0 {
		SplGlobals.SetHashMaskInit(0)
	}
	return zend.SUCCESS
}

/* {{{ spl_module_entry
 */

var SplModuleEntry zend.ZendModuleEntry = zend.ZendModuleEntry{g.SizeOf("zend_module_entry"), 20190902, 0, 0, nil, nil, "SPL", SplFunctions, ZmStartupSpl, nil, ZmActivateSpl, ZmDeactivateSpl, ZmInfoSpl, "7.4.33", g.SizeOf("zend_spl_globals"), &SplGlobals, (func(any))(ZmGlobalsCtorSpl), nil, nil, 0, 0, nil, 0, "API" + "20190902" + ",NTS"}

/* }}} */
