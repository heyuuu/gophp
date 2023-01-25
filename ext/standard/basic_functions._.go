// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

// Source: <ext/standard/basic_functions.h>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

const BasicFunctionsModulePtr = &BasicFunctionsModule

/* system functions */

var ZifSetTimeLimit func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var ZifHeaderRegisterCallback func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)

/* From the INI parser */

/* Left for BC (not binary safe!) */

const MT_N = 624

/* Deprecated type aliases -- use the standard types instead */

type PhpUint32 = uint32
type PhpInt32 = int32

var BasicGlobals PhpBasicGlobals

// Source: <ext/standard/basic_functions.c>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// failed # include "ext/session/php_session.h"

type YY_BUFFER_STATE *__struct__yy_buffer_state

const INADDR_NONE = zend_ulong - 1

var IncompleteClassEntry *zend.ZendClassEntry = nil

/* some prototypes for local functions */

/* {{{ arginfo */

var ArginfoSetTimeLimit []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"seconds", 0, 0, 0},
}

/* }}} */

var ArginfoHeaderRegisterCallback []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"callback", 0, 0, 0},
}

/* }}} */

var ArginfoObStart []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"user_function", 0, 0, 0}, {"chunk_size", 0, 0, 0}, {"flags", 0, 0, 0}}
var ArginfoObFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoObClean []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoObEndFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoObEndClean []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoObGetFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoObGetClean []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoObGetContents []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoObGetLevel []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoObGetLength []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoObListHandlers []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoObGetStatus []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"full_status", 0, 0, 0}}
var ArginfoObImplicitFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"flag", 0, 0, 0}}
var ArginfoOutputResetRewriteVars []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoOutputAddRewriteVar []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"name", 0, 0, 0},
	{"value", 0, 0, 0},
}

/* }}} */

var ArginfoStreamWrapperRegister []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"protocol", 0, 0, 0}, {"classname", 0, 0, 0}, {"flags", 0, 0, 0}}
var ArginfoStreamWrapperUnregister []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"protocol", 0, 0, 0},
}
var ArginfoStreamWrapperRestore []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"protocol", 0, 0, 0},
}

/* }}} */

var ArginfoKrsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"sort_flags", 0, 0, 0}}
var ArginfoKsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"sort_flags", 0, 0, 0}}
var ArginfoCount []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var", 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoNatsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 1, 0},
}
var ArginfoNatcasesort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 1, 0},
}
var ArginfoAsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"sort_flags", 0, 0, 0}}
var ArginfoArsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"sort_flags", 0, 0, 0}}
var ArginfoSort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"sort_flags", 0, 0, 0}}
var ArginfoRsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"sort_flags", 0, 0, 0}}
var ArginfoUsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 1, 0},
	{"cmp_function", 0, 0, 0},
}
var ArginfoUasort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 1, 0},
	{"cmp_function", 0, 0, 0},
}
var ArginfoUksort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 1, 0},
	{"cmp_function", 0, 0, 0},
}
var ArginfoEnd []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 1, 0},
}
var ArginfoPrev []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 1, 0},
}
var ArginfoNext []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 1, 0},
}
var ArginfoReset []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 1, 0},
}
var ArginfoCurrent []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 0, 0},
}
var ArginfoKey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 0, 0},
}
var ArginfoMin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"args", 0, 0, 1}}
var ArginfoMax []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"args", 0, 0, 1}}
var ArginfoArrayWalk []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"input", 0, 1, 0}, {"funcname", 0, 0, 0}, {"userdata", 0, 0, 0}}
var ArginfoArrayWalkRecursive []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"input", 0, 1, 0}, {"funcname", 0, 0, 0}, {"userdata", 0, 0, 0}}
var ArginfoInArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"needle", 0, 0, 0}, {"haystack", 0, 0, 0}, {"strict", 0, 0, 0}}
var ArginfoArraySearch []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"needle", 0, 0, 0}, {"haystack", 0, 0, 0}, {"strict", 0, 0, 0}}
var ArginfoExtract []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, zend.ZEND_SEND_PREFER_REF, 0}, {"extract_type", 0, 0, 0}, {"prefix", 0, 0, 0}}
var ArginfoCompact []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var_names", 0, 0, 1}}
var ArginfoArrayFill []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"start_key", 0, 0, 0},
	{"num", 0, 0, 0},
	{"val", 0, 0, 0},
}
var ArginfoArrayFillKeys []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"keys", 0, 0, 0},
	{"val", 0, 0, 0},
}
var ArginfoRange []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"low", 0, 0, 0}, {"high", 0, 0, 0}, {"step", 0, 0, 0}}
var ArginfoShuffle []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 1, 0},
}
var ArginfoArrayPush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"stack", 0, 1, 0}, {"vars", 0, 0, 1}}
var ArginfoArrayPop []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"stack", 0, 1, 0},
}
var ArginfoArrayShift []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"stack", 0, 1, 0},
}
var ArginfoArrayUnshift []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"stack", 0, 1, 0}, {"vars", 0, 0, 1}}
var ArginfoArraySplice []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arg", 0, 1, 0}, {"offset", 0, 0, 0}, {"length", 0, 0, 0}, {"replacement", 0, 0, 0}}
var ArginfoArraySlice []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arg", 0, 0, 0}, {"offset", 0, 0, 0}, {"length", 0, 0, 0}, {"preserve_keys", 0, 0, 0}}
var ArginfoArrayMerge []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayMergeRecursive []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayReplace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayReplaceRecursive []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayKeys []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 0, 0}, {"search_value", 0, 0, 0}, {"strict", 0, 0, 0}}
var ArginfoArrayKeyFirst []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 0, 0},
}
var ArginfoArrayKeyLast []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 0, 0},
}
var ArginfoArrayValues []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 0, 0},
}
var ArginfoArrayCountValues []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 0, 0},
}
var ArginfoArrayColumn []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arg", 0, 0, 0}, {"column_key", 0, 0, 0}, {"index_key", 0, 0, 0}}
var ArginfoArrayReverse []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"input", 0, 0, 0}, {"preserve_keys", 0, 0, 0}}
var ArginfoArrayPad []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 0, 0},
	{"pad_size", 0, 0, 0},
	{"pad_value", 0, 0, 0},
}
var ArginfoArrayFlip []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 0, 0},
}
var ArginfoArrayChangeKeyCase []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"input", 0, 0, 0}, {"case", 0, 0, 0}}
var ArginfoArrayUnique []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 0, 0}, {"flags", 0, 0, 0}}
var ArginfoArrayIntersectKey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayIntersectUkey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arr1", 0, 0, 0},
	{"arr2", 0, 0, 0},
	{"callback_key_compare_func", 0, 0, 0},
}
var ArginfoArrayIntersect []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayUintersect []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arr1", 0, 0, 0},
	{"arr2", 0, 0, 0},
	{"callback_data_compare_func", 0, 0, 0},
}
var ArginfoArrayIntersectAssoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayUintersectAssoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arr1", 0, 0, 0},
	{"arr2", 0, 0, 0},
	{"callback_data_compare_func", 0, 0, 0},
}
var ArginfoArrayIntersectUassoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arr1", 0, 0, 0},
	{"arr2", 0, 0, 0},
	{"callback_key_compare_func", 0, 0, 0},
}
var ArginfoArrayUintersectUassoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arr1", 0, 0, 0},
	{"arr2", 0, 0, 0},
	{"callback_data_compare_func", 0, 0, 0},
	{"callback_key_compare_func", 0, 0, 0},
}
var ArginfoArrayDiffKey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayDiffUkey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arr1", 0, 0, 0},
	{"arr2", 0, 0, 0},
	{"callback_key_comp_func", 0, 0, 0},
}
var ArginfoArrayDiff []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayUdiff []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arr1", 0, 0, 0},
	{"arr2", 0, 0, 0},
	{"callback_data_comp_func", 0, 0, 0},
}
var ArginfoArrayDiffAssoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayDiffUassoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arr1", 0, 0, 0},
	{"arr2", 0, 0, 0},
	{"callback_data_comp_func", 0, 0, 0},
}
var ArginfoArrayUdiffAssoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arr1", 0, 0, 0},
	{"arr2", 0, 0, 0},
	{"callback_key_comp_func", 0, 0, 0},
}
var ArginfoArrayUdiffUassoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arr1", 0, 0, 0},
	{"arr2", 0, 0, 0},
	{"callback_data_comp_func", 0, 0, 0},
	{"callback_key_comp_func", 0, 0, 0},
}
var ArginfoArrayMultisort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{"arr1", 0, zend.ZEND_SEND_PREFER_REF, 0},
	{"sort_order", 0, zend.ZEND_SEND_PREFER_REF, 0},
	{"sort_flags", 0, zend.ZEND_SEND_PREFER_REF, 0},
	{"arr2", 0, zend.ZEND_SEND_PREFER_REF, 1},
}
var ArginfoArrayRand []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 0, 0}, {"num_req", 0, 0, 0}}
var ArginfoArraySum []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 0, 0},
}
var ArginfoArrayProduct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 0, 0},
}
var ArginfoArrayReduce []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arg", 0, 0, 0}, {"callback", 0, 0, 0}, {"initial", 0, 0, 0}}
var ArginfoArrayFilter []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 0, 0}, {"callback", 0, 0, 0}, {"use_keys", 0, 0, 0}}
var ArginfoArrayMap []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"callback", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayKeyExists []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"key", 0, 0, 0},
	{"search", 0, 0, 0},
}
var ArginfoArrayChunk []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arg", 0, 0, 0}, {"size", 0, 0, 0}, {"preserve_keys", 0, 0, 0}}
var ArginfoArrayCombine []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"keys", 0, 0, 0},
	{"values", 0, 0, 0},
}

/* }}} */

var ArginfoGetMagicQuotesGpc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoGetMagicQuotesRuntime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoConstant []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"const_name", 0, 0, 0},
}
var ArginfoInetNtop []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"in_addr", 0, 0, 0},
}
var ArginfoInetPton []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"ip_address", 0, 0, 0},
}
var ArginfoIp2long []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"ip_address", 0, 0, 0},
}
var ArginfoLong2ip []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"proper_address", 0, 0, 0},
}
var ArginfoGetenv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"varname", 0, 0, 0}, {"local_only", 0, 0, 0}}
var ArginfoPutenv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"setting", 0, 0, 0},
}
var ArginfoGetopt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"options", 0, 0, 0}, {"opts", 0, 0, 0}, {"optind", 0, 1, 0}}
var ArginfoFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoSleep []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"seconds", 0, 0, 0},
}
var ArginfoUsleep []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"micro_seconds", 0, 0, 0},
}
var ArginfoTimeNanosleep []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"seconds", 0, 0, 0},
	{"nanoseconds", 0, 0, 0},
}
var ArginfoTimeSleepUntil []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"timestamp", 0, 0, 0},
}
var ArginfoGetCurrentUser []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoGetCfgVar []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"option_name", 0, 0, 0},
}
var ArginfoErrorLog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"message", 0, 0, 0}, {"message_type", 0, 0, 0}, {"destination", 0, 0, 0}, {"extra_headers", 0, 0, 0}}
var ArginfoErrorGetLast []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}}
var ArginfoErrorClearLast []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}}
var ArginfoCallUserFunc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"function_name", 0, 0, 0}, {"parameters", 0, 0, 1}}
var ArginfoCallUserFuncArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"function_name", 0, 0, 0}, {"parameters", 0, 0, 0}}
var ArginfoForwardStaticCall []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"function_name", 0, 0, 0}, {"parameters", 0, 0, 1}}
var ArginfoForwardStaticCallArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"function_name", 0, 0, 0}, {"parameters", 0, 0, 0}}
var ArginfoRegisterShutdownFunction []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"function_name", 0, 0, 0}, {"parameters", 0, 0, 1}}
var ArginfoHighlightFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"file_name", 0, 0, 0}, {"return", 0, 0, 0}}
var ArginfoPhpStripWhitespace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"file_name", 0, 0, 0},
}
var ArginfoHighlightString []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"string", 0, 0, 0}, {"return", 0, 0, 0}}
var ArginfoIniGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"varname", 0, 0, 0},
}
var ArginfoIniGetAll []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"extension", 0, 0, 0}, {"details", 0, 0, 0}}
var ArginfoIniSet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"varname", 0, 0, 0},
	{"newvalue", 0, 0, 0},
}
var ArginfoIniRestore []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"varname", 0, 0, 0},
}
var ArginfoSetIncludePath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"new_include_path", 0, 0, 0},
}
var ArginfoGetIncludePath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoRestoreIncludePath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoPrintR []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var", 0, 0, 0}, {"return", 0, 0, 0}}
var ArginfoConnectionAborted []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoConnectionStatus []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoIgnoreUserAbort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"value", 0, 0, 0}}
var ArginfoGetservbyname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"service", 0, 0, 0},
	{"protocol", 0, 0, 0},
}
var ArginfoGetservbyport []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"port", 0, 0, 0},
	{"protocol", 0, 0, 0},
}
var ArginfoGetprotobyname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"name", 0, 0, 0},
}
var ArginfoGetprotobynumber []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"proto", 0, 0, 0},
}
var ArginfoRegisterTickFunction []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"function_name", 0, 0, 0}, {"parameters", 0, 0, 1}}
var ArginfoUnregisterTickFunction []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"function_name", 0, 0, 0},
}
var ArginfoIsUploadedFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"path", 0, 0, 0},
}
var ArginfoMoveUploadedFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"path", 0, 0, 0},
	{"new_path", 0, 0, 0},
}
var ArginfoParseIniFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"process_sections", 0, 0, 0}, {"scanner_mode", 0, 0, 0}}
var ArginfoParseIniString []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"ini_string", 0, 0, 0}, {"process_sections", 0, 0, 0}, {"scanner_mode", 0, 0, 0}}
var ArginfoSysGetloadavg []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}

/* }}} */

var ArginfoAssert []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"assertion", 0, 0, 0}, {"description", 0, 0, 0}}
var ArginfoAssertOptions []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"what", 0, 0, 0}, {"value", 0, 0, 0}}

/* }}} */

var ArginfoBase64Encode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoBase64Decode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"strict", 0, 0, 0}}

/* }}} */

var ArginfoGetBrowser []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"browser_name", 0, 0, 0}, {"return_array", 0, 0, 0}}

/* }}} */

var ArginfoCrc32 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}

/* }}} */

var ArginfoCrypt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"salt", 0, 0, 0}}

/* }}} */

var ArginfoConvertCyrString []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
	{"from", 0, 0, 0},
	{"to", 0, 0, 0},
}

/* }}} */

var ArginfoStrptime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"timestamp", 0, 0, 0},
	{"format", 0, 0, 0},
}

/* }}} */

var ArginfoOpendir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"path", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoDir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"directory", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoClosedir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"dir_handle", 0, 0, 0}}
var ArginfoChroot []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"directory", 0, 0, 0},
}
var ArginfoChdir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"directory", 0, 0, 0},
}
var ArginfoGetcwd []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoRewinddir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"dir_handle", 0, 0, 0}}
var ArginfoReaddir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"dir_handle", 0, 0, 0}}
var ArginfoGlob []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"pattern", 0, 0, 0}, {"flags", 0, 0, 0}}
var ArginfoScandir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"dir", 0, 0, 0}, {"sorting_order", 0, 0, 0}, {"context", 0, 0, 0}}

/* }}} */

var ArginfoGethostbyaddr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"ip_address", 0, 0, 0},
}
var ArginfoGethostbyname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"hostname", 0, 0, 0},
}
var ArginfoGethostbynamel []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"hostname", 0, 0, 0},
}
var ArginfoGethostname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoNetGetInterfaces []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoDnsCheckRecord []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"host", 0, 0, 0}, {"type", 0, 0, 0}}
var ArginfoDnsGetRecord []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{"hostname", 0, 0, 0},
	{"type", 0, 0, 0},
	{"authns", zend.ZEND_TYPE_ENCODE(zend.IS_ARRAY, 1), 1, 0},
	{"addtl", zend.ZEND_TYPE_ENCODE(zend.IS_ARRAY, 1), 1, 0},
	{"raw", 0, 0, 0},
}
var ArginfoDnsGetMx []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"hostname", 0, 0, 0}, {"mxhosts", 0, 1, 0}, {"weight", 0, 1, 0}}

/* }}} */

var ArginfoExec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"command", 0, 0, 0}, {"output", 0, 1, 0}, {"return_value", 0, 1, 0}}
var ArginfoSystem []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"command", 0, 0, 0}, {"return_value", 0, 1, 0}}
var ArginfoPassthru []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"command", 0, 0, 0}, {"return_value", 0, 1, 0}}
var ArginfoEscapeshellcmd []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"command", 0, 0, 0},
}
var ArginfoEscapeshellarg []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"arg", 0, 0, 0},
}
var ArginfoShellExec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"cmd", 0, 0, 0},
}
var ArginfoProcNice []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"priority", 0, 0, 0},
}

/* }}} */

var ArginfoFlock []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"fp", 0, 0, 0}, {"operation", 0, 0, 0}, {"wouldblock", 0, 1, 0}}
var ArginfoGetMetaTags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"use_include_path", 0, 0, 0}}
var ArginfoFileGetContents []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"flags", 0, 0, 0}, {"context", 0, 0, 0}, {"offset", 0, 0, 0}, {"maxlen", 0, 0, 0}}
var ArginfoFilePutContents []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"filename", 0, 0, 0}, {"data", 0, 0, 0}, {"flags", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"flags", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoTempnam []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"dir", 0, 0, 0},
	{"prefix", 0, 0, 0},
}
var ArginfoTmpfile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoFopen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"filename", 0, 0, 0}, {"mode", 0, 0, 0}, {"use_include_path", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoFclose []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
}
var ArginfoPopen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"command", 0, 0, 0},
	{"mode", 0, 0, 0},
}
var ArginfoPclose []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
}
var ArginfoFeof []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
}
var ArginfoFgets []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"fp", 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoFgetc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
}
var ArginfoFgetss []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"fp", 0, 0, 0}, {"length", 0, 0, 0}, {"allowable_tags", 0, 0, 0}}
var ArginfoFscanf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"format", 0, 0, 0}, {"vars", 0, 1, 1}}
var ArginfoFwrite []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"fp", 0, 0, 0}, {"str", 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoFflush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
}
var ArginfoRewind []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
}
var ArginfoFtell []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
}
var ArginfoFseek []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"fp", 0, 0, 0}, {"offset", 0, 0, 0}, {"whence", 0, 0, 0}}
var ArginfoMkdir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"pathname", 0, 0, 0}, {"mode", 0, 0, 0}, {"recursive", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoRmdir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"dirname", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoReadfile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"flags", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoUmask []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"mask", 0, 0, 0}}
var ArginfoFpassthru []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
}
var ArginfoRename []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"old_name", 0, 0, 0}, {"new_name", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoUnlink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoFtruncate []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
	{"size", 0, 0, 0},
}
var ArginfoFstat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
}
var ArginfoCopy []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"source_file", 0, 0, 0}, {"destination_file", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoFread []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
	{"length", 0, 0, 0},
}
var ArginfoFputcsv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"fp", 0, 0, 0}, {"fields", 0, 0, 0}, {"delimiter", 0, 0, 0}, {"enclosure", 0, 0, 0}, {"escape_char", 0, 0, 0}}
var ArginfoFgetcsv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"fp", 0, 0, 0}, {"length", 0, 0, 0}, {"delimiter", 0, 0, 0}, {"enclosure", 0, 0, 0}, {"escape", 0, 0, 0}}
var ArginfoRealpath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"path", 0, 0, 0},
}
var ArginfoFnmatch []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"pattern", 0, 0, 0}, {"filename", 0, 0, 0}, {"flags", 0, 0, 0}}
var ArginfoSysGetTempDir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}

/* }}} */

var ArginfoDiskTotalSpace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"path", 0, 0, 0},
}
var ArginfoDiskFreeSpace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"path", 0, 0, 0},
}
var ArginfoChgrp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
	{"group", 0, 0, 0},
}
var ArginfoChown []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
	{"user", 0, 0, 0},
}
var ArginfoLchgrp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
	{"group", 0, 0, 0},
}
var ArginfoLchown []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
	{"user", 0, 0, 0},
}
var ArginfoChmod []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
	{"mode", 0, 0, 0},
}
var ArginfoTouch []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"time", 0, 0, 0}, {"atime", 0, 0, 0}}
var ArginfoClearstatcache []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"clear_realpath_cache", 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoRealpathCacheSize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoRealpathCacheGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoFileperms []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoFileinode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoFilesize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoFileowner []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoFilegroup []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoFileatime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoFilemtime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoFilectime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoFiletype []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoIsWritable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoIsReadable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoIsExecutable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoIsFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoIsDir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoIsLink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoFileExists []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoLstat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoStat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}

/* }}} */

var ArginfoSprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"format", 0, 0, 0}, {"args", 0, 0, 1}}
var ArginfoVsprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"format", 0, 0, 0},
	{"args", 0, 0, 0},
}
var ArginfoPrintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"format", 0, 0, 0}, {"args", 0, 0, 1}}
var ArginfoVprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"format", 0, 0, 0},
	{"args", 0, 0, 0},
}
var ArginfoFprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"format", 0, 0, 0}, {"args", 0, 0, 1}}
var ArginfoVfprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"stream", 0, 0, 0},
	{"format", 0, 0, 0},
	{"args", 0, 0, 0},
}

/* }}} */

var ArginfoFsockopen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"hostname", 0, 0, 0}, {"port", 0, 0, 0}, {"errno", 0, 1, 0}, {"errstr", 0, 1, 0}, {"timeout", 0, 0, 0}}
var ArginfoPfsockopen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"hostname", 0, 0, 0}, {"port", 0, 0, 0}, {"errno", 0, 1, 0}, {"errstr", 0, 1, 0}, {"timeout", 0, 0, 0}}

/* }}} */

var ArginfoFtok []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"pathname", 0, 0, 0},
	{"proj", 0, 0, 0},
}

/* }}} */

var ArginfoHeader []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"header", 0, 0, 0}, {"replace", 0, 0, 0}, {"http_response_code", 0, 0, 0}}
var ArginfoHeaderRemove []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"name", 0, 0, 0}}
var ArginfoSetcookie []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"name", 0, 0, 0}, {"value", 0, 0, 0}, {"expires_or_options", 0, 0, 0}, {"path", 0, 0, 0}, {"domain", 0, 0, 0}, {"secure", 0, 0, 0}, {"httponly", 0, 0, 0}}
var ArginfoSetrawcookie []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"name", 0, 0, 0}, {"value", 0, 0, 0}, {"expires_or_options", 0, 0, 0}, {"path", 0, 0, 0}, {"domain", 0, 0, 0}, {"secure", 0, 0, 0}, {"httponly", 0, 0, 0}}
var ArginfoHeadersSent []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"file", 0, 1, 0}, {"line", 0, 1, 0}}
var ArginfoHeadersList []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoHttpResponseCode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"response_code", 0, 0, 0}}

/* }}} */

var ArginfoHrtime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"get_as_number", 0, 0, 0},
}

/* }}} */

var ArginfoHtmlspecialchars []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"string", 0, 0, 0}, {"quote_style", 0, 0, 0}, {"encoding", 0, 0, 0}, {"double_encode", 0, 0, 0}}
var ArginfoHtmlspecialcharsDecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"string", 0, 0, 0}, {"quote_style", 0, 0, 0}}
var ArginfoHtmlEntityDecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"string", 0, 0, 0}, {"quote_style", 0, 0, 0}, {"encoding", 0, 0, 0}}
var ArginfoHtmlentities []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"string", 0, 0, 0}, {"quote_style", 0, 0, 0}, {"encoding", 0, 0, 0}, {"double_encode", 0, 0, 0}}
var ArginfoGetHtmlTranslationTable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"table", 0, 0, 0}, {"quote_style", 0, 0, 0}, {"encoding", 0, 0, 0}}

/* }}} */

var ArginfoHttpBuildQuery []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"formdata", 0, 0, 0}, {"prefix", 0, 0, 0}, {"arg_separator", 0, 0, 0}, {"enc_type", 0, 0, 0}}

/* }}} */

var ArginfoImageTypeToMimeType []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"imagetype", 0, 0, 0},
}
var ArginfoImageTypeToExtension []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"imagetype", 0, 0, 0}, {"include_dot", 0, 0, 0}}
var ArginfoGetimagesize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"imagefile", 0, 0, 0}, {"info", 0, 1, 0}}

/* }}} */

var ArginfoPhpinfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"what", 0, 0, 0}}
var ArginfoPhpversion []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"extension", 0, 0, 0}}
var ArginfoPhpcredits []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"flag", 0, 0, 0}}
var ArginfoPhpSapiName []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoPhpUname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoPhpIniScannedFiles []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoPhpIniLoadedFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}

/* }}} */

var ArginfoIptcembed []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"iptcdata", 0, 0, 0}, {"jpeg_file_name", 0, 0, 0}, {"spool", 0, 0, 0}}
var ArginfoIptcparse []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"iptcdata", 0, 0, 0},
}

/* }}} */

var ArginfoLcgValue []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}

/* }}} */

var ArginfoLevenshtein []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str1", 0, 0, 0}, {"str2", 0, 0, 0}, {"cost_ins", 0, 0, 0}, {"cost_rep", 0, 0, 0}, {"cost_del", 0, 0, 0}}

/* }}} */

var ArginfoReadlink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoLinkinfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoSymlink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"target", 0, 0, 0},
	{"link", 0, 0, 0},
}
var ArginfoLink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"target", 0, 0, 0},
	{"link", 0, 0, 0},
}

/* }}} */

var ArginfoEzmlmHash []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"addr", 0, 0, 0},
}
var ArginfoMail []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(3)), 0, 0, 0}, {"to", 0, 0, 0}, {"subject", 0, 0, 0}, {"message", 0, 0, 0}, {"additional_headers", 0, 0, 0}, {"additional_parameters", 0, 0, 0}}

/* }}} */

var ArginfoAbs []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoCeil []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoFloor []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoRound []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"number", 0, 0, 0}, {"precision", 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoSin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoCos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoTan []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoAsin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoAcos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoAtan []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoAtan2 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"y", 0, 0, 0},
	{"x", 0, 0, 0},
}
var ArginfoSinh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoCosh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoTanh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoAsinh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoAcosh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoAtanh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoPi []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoIsFinite []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"val", 0, 0, 0},
}
var ArginfoIsInfinite []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"val", 0, 0, 0},
}
var ArginfoIsNan []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"val", 0, 0, 0},
}
var ArginfoPow []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"base", 0, 0, 0},
	{"exponent", 0, 0, 0},
}
var ArginfoExp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoExpm1 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoLog1p []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoLog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"number", 0, 0, 0}, {"base", 0, 0, 0}}
var ArginfoLog10 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoSqrt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoHypot []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"num1", 0, 0, 0},
	{"num2", 0, 0, 0},
}
var ArginfoDeg2rad []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoRad2deg []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
}
var ArginfoBindec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"binary_number", 0, 0, 0},
}
var ArginfoHexdec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"hexadecimal_number", 0, 0, 0},
}
var ArginfoOctdec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"octal_number", 0, 0, 0},
}
var ArginfoDecbin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"decimal_number", 0, 0, 0},
}
var ArginfoDecoct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"decimal_number", 0, 0, 0},
}
var ArginfoDechex []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"decimal_number", 0, 0, 0},
}
var ArginfoBaseConvert []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"number", 0, 0, 0},
	{"frombase", 0, 0, 0},
	{"tobase", 0, 0, 0},
}
var ArginfoNumberFormat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"number", 0, 0, 0}, {"num_decimal_places", 0, 0, 0}, {"dec_separator", 0, 0, 0}, {"thousands_separator", 0, 0, 0}}
var ArginfoFmod []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"x", 0, 0, 0},
	{"y", 0, 0, 0},
}
var ArginfoIntdiv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"dividend", 0, 0, 0},
	{"divisor", 0, 0, 0},
}

/* }}} */

var ArginfoMd5 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"raw_output", 0, 0, 0}}
var ArginfoMd5File []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"raw_output", 0, 0, 0}}

/* }}} */

var ArginfoMetaphone []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"text", 0, 0, 0}, {"phones", 0, 0, 0}}

/* }}} */

var ArginfoMicrotime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"get_as_float", 0, 0, 0}}
var ArginfoGettimeofday []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"get_as_float", 0, 0, 0}}
var ArginfoGetrusage []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"who", 0, 0, 0}}

/* }}} */

var ArginfoPack []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"format", 0, 0, 0}, {"args", 0, 0, 1}}
var ArginfoUnpack []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"format", 0, 0, 0}, {"input", 0, 0, 0}, {"offset", 0, 0, 0}}

/* }}} */

var ArginfoGetmyuid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoGetmygid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoGetmypid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoGetmyinode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoGetlastmod []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}

/* }}} */

var ArginfoPasswordHash []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"password", 0, 0, 0}, {"algo", 0, 0, 0}, {"options", 0, 0, 0}}
var ArginfoPasswordGetInfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"hash", 0, 0, 0}}
var ArginfoPasswordNeedsRehash []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"hash", 0, 0, 0}, {"algo", 0, 0, 0}, {"options", 0, 0, 0}}
var ArginfoPasswordVerify []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"password", 0, 0, 0}, {"hash", 0, 0, 0}}
var ArginfoPasswordAlgos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}

/* }}} */

var ArginfoProcTerminate []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"process", 0, 0, 0}, {"signal", 0, 0, 0}}
var ArginfoProcClose []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"process", 0, 0, 0},
}
var ArginfoProcGetStatus []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"process", 0, 0, 0},
}
var ArginfoProcOpen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(3)), 0, 0, 0}, {"command", 0, 0, 0}, {"descriptorspec", 0, 0, 0}, {"pipes", 0, 1, 0}, {"cwd", 0, 0, 0}, {"env", 0, 0, 0}, {"other_options", 0, 0, 0}}

/* }}} */

var ArginfoQuotedPrintableDecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}

/* }}} */

var ArginfoQuotedPrintableEncode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}

/* }}} */

var ArginfoMtSrand []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"seed", 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoMtRand []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"min", 0, 0, 0}, {"max", 0, 0, 0}}
var ArginfoMtGetrandmax []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}

/* }}} */

var ArginfoRandomBytes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoRandomInt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"min", 0, 0, 0}, {"max", 0, 0, 0}}

/* }}} */

var ArginfoSha1 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"raw_output", 0, 0, 0}}
var ArginfoSha1File []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"raw_output", 0, 0, 0}}

/* }}} */

var ArginfoSoundex []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}

/* }}} */

var ArginfoStreamSocketPair []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"domain", 0, 0, 0},
	{"type", 0, 0, 0},
	{"protocol", 0, 0, 0},
}
var ArginfoStreamSocketClient []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"remoteaddress", 0, 0, 0}, {"errcode", 0, 1, 0}, {"errstring", 0, 1, 0}, {"timeout", 0, 0, 0}, {"flags", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoStreamSocketServer []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"localaddress", 0, 0, 0}, {"errcode", 0, 1, 0}, {"errstring", 0, 1, 0}, {"flags", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoStreamSocketAccept []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"serverstream", 0, 0, 0}, {"timeout", 0, 0, 0}, {"peername", 0, 1, 0}}
var ArginfoStreamSocketGetName []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"stream", 0, 0, 0},
	{"want_peer", 0, 0, 0},
}
var ArginfoStreamSocketSendto []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"data", 0, 0, 0}, {"flags", 0, 0, 0}, {"target_addr", 0, 0, 0}}
var ArginfoStreamSocketRecvfrom []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"amount", 0, 0, 0}, {"flags", 0, 0, 0}, {"remote_addr", 0, 1, 0}}
var ArginfoStreamGetContents []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"source", 0, 0, 0}, {"maxlen", 0, 0, 0}, {"offset", 0, 0, 0}}
var ArginfoStreamCopyToStream []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"source", 0, 0, 0}, {"dest", 0, 0, 0}, {"maxlen", 0, 0, 0}, {"pos", 0, 0, 0}}
var ArginfoStreamGetMetaData []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
}
var ArginfoStreamGetTransports []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoStreamGetWrappers []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoStreamResolveIncludePath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filename", 0, 0, 0},
}
var ArginfoStreamIsLocal []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"stream", 0, 0, 0},
}
var ArginfoStreamSupportsLock []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"stream", 0, 0, 0}}
var ArginfoStreamIsatty []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"stream", 0, 0, 0}}
var ArginfoStreamSelect []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(4)), 0, 0, 0}, {"read_streams", 0, 1, 0}, {"write_streams", 0, 1, 0}, {"except_streams", 0, 1, 0}, {"tv_sec", 0, 0, 0}, {"tv_usec", 0, 0, 0}}
var ArginfoStreamContextGetOptions []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"stream_or_context", 0, 0, 0},
}
var ArginfoStreamContextSetOption []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream_or_context", 0, 0, 0}, {"wrappername", 0, 0, 0}, {"optionname", 0, 0, 0}, {"value", 0, 0, 0}}
var ArginfoStreamContextSetParams []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"stream_or_context", 0, 0, 0},
	{"options", 0, 0, 0},
}
var ArginfoStreamContextGetParams []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"stream_or_context", 0, 0, 0},
}
var ArginfoStreamContextGetDefault []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"options", 0, 0, 0}}
var ArginfoStreamContextSetDefault []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"options", 0, 0, 0},
}
var ArginfoStreamContextCreate []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"options", 0, 0, 0}, {"params", 0, 0, 0}}
var ArginfoStreamFilterPrepend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"filtername", 0, 0, 0}, {"read_write", 0, 0, 0}, {"filterparams", 0, 0, 0}}
var ArginfoStreamFilterAppend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"filtername", 0, 0, 0}, {"read_write", 0, 0, 0}, {"filterparams", 0, 0, 0}}
var ArginfoStreamFilterRemove []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"stream_filter", 0, 0, 0},
}
var ArginfoStreamGetLine []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"maxlen", 0, 0, 0}, {"ending", 0, 0, 0}}
var ArginfoStreamSetBlocking []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"socket", 0, 0, 0},
	{"mode", 0, 0, 0},
}
var ArginfoStreamSetTimeout []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"seconds", 0, 0, 0}, {"microseconds", 0, 0, 0}}
var ArginfoStreamSetReadBuffer []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
	{"buffer", 0, 0, 0},
}
var ArginfoStreamSetWriteBuffer []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
	{"buffer", 0, 0, 0},
}
var ArginfoStreamSetChunkSize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"fp", 0, 0, 0},
	{"chunk_size", 0, 0, 0},
}
var ArginfoStreamSocketEnableCrypto []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"enable", 0, 0, 0}, {"cryptokind", 0, 0, 0}, {"sessionstream", 0, 0, 0}}
var ArginfoStreamSocketShutdown []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"stream", 0, 0, 0},
	{"how", 0, 0, 0},
}

/* }}} */

var ArginfoBin2hex []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"data", 0, 0, 0},
}
var ArginfoHex2bin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"data", 0, 0, 0},
}
var ArginfoStrspn []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str", 0, 0, 0}, {"mask", 0, 0, 0}, {"start", 0, 0, 0}, {"len", 0, 0, 0}}
var ArginfoStrcspn []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str", 0, 0, 0}, {"mask", 0, 0, 0}, {"start", 0, 0, 0}, {"len", 0, 0, 0}}
var ArginfoStrcoll []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str1", 0, 0, 0},
	{"str2", 0, 0, 0},
}
var ArginfoTrim []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"character_mask", 0, 0, 0}}
var ArginfoRtrim []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"character_mask", 0, 0, 0}}
var ArginfoLtrim []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"character_mask", 0, 0, 0}}
var ArginfoWordwrap []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"width", 0, 0, 0}, {"break", 0, 0, 0}, {"cut", 0, 0, 0}}
var ArginfoExplode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"separator", 0, 0, 0}, {"str", 0, 0, 0}, {"limit", 0, 0, 0}}
var ArginfoImplode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"glue", 0, 0, 0},
	{"pieces", 0, 0, 0},
}
var ArginfoStrtok []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"token", 0, 0, 0}}
var ArginfoStrtoupper []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoStrtolower []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoBasename []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"path", 0, 0, 0}, {"suffix", 0, 0, 0}}
var ArginfoDirname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"path", 0, 0, 0}, {"levels", 0, 0, 0}}
var ArginfoPathinfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"path", 0, 0, 0}, {"options", 0, 0, 0}}
var ArginfoStristr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"part", 0, 0, 0}}
var ArginfoStrstr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"part", 0, 0, 0}}
var ArginfoStrpos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"offset", 0, 0, 0}}
var ArginfoStripos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"offset", 0, 0, 0}}
var ArginfoStrrpos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"offset", 0, 0, 0}}
var ArginfoStrripos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"offset", 0, 0, 0}}
var ArginfoStrrchr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"haystack", 0, 0, 0},
	{"needle", 0, 0, 0},
}
var ArginfoChunkSplit []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"chunklen", 0, 0, 0}, {"ending", 0, 0, 0}}
var ArginfoSubstr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str", 0, 0, 0}, {"start", 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoSubstrReplace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(3)), 0, 0, 0}, {"str", 0, 0, 0}, {"replace", 0, 0, 0}, {"start", 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoQuotemeta []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoOrd []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"character", 0, 0, 0},
}
var ArginfoChr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"codepoint", 0, 0, 0},
}
var ArginfoUcfirst []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoLcfirst []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoUcwords []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"delimiters", 0, 0, 0}}
var ArginfoStrtr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str", 0, 0, 0}, {"from", 0, 0, 0}, {"to", 0, 0, 0}}
var ArginfoStrrev []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoSimilarText []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str1", 0, 0, 0}, {"str2", 0, 0, 0}, {"percent", 0, 1, 0}}
var ArginfoAddcslashes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
	{"charlist", 0, 0, 0},
}
var ArginfoAddslashes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoStripcslashes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoStripslashes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoStrReplace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(3)), 0, 0, 0}, {"search", 0, 0, 0}, {"replace", 0, 0, 0}, {"subject", 0, 0, 0}, {"replace_count", 0, 1, 0}}
var ArginfoStrIreplace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(3)), 0, 0, 0}, {"search", 0, 0, 0}, {"replace", 0, 0, 0}, {"subject", 0, 0, 0}, {"replace_count", 0, 1, 0}}
var ArginfoHebrev []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"max_chars_per_line", 0, 0, 0}}
var ArginfoHebrevc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"max_chars_per_line", 0, 0, 0}}
var ArginfoNl2br []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"is_xhtml", 0, 0, 0}}
var ArginfoStripTags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"allowable_tags", 0, 0, 0}}
var ArginfoSetlocale []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"category", 0, 0, 0}, {"locales", 0, 0, 1}}
var ArginfoParseStr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"encoded_string", 0, 0, 0}, {"result", 0, 1, 0}}
var ArginfoStrGetcsv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"string", 0, 0, 0}, {"delimiter", 0, 0, 0}, {"enclosure", 0, 0, 0}, {"escape", 0, 0, 0}}
var ArginfoStrRepeat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"input", 0, 0, 0},
	{"mult", 0, 0, 0},
}
var ArginfoCountChars []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"input", 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoStrnatcmp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"s1", 0, 0, 0},
	{"s2", 0, 0, 0},
}
var ArginfoLocaleconv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoStrnatcasecmp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"s1", 0, 0, 0},
	{"s2", 0, 0, 0},
}
var ArginfoSubstrCount []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"offset", 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoStrPad []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"input", 0, 0, 0}, {"pad_length", 0, 0, 0}, {"pad_string", 0, 0, 0}, {"pad_type", 0, 0, 0}}
var ArginfoSscanf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str", 0, 0, 0}, {"format", 0, 0, 0}, {"vars", 0, 1, 1}}
var ArginfoStrRot13 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoStrShuffle []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoStrWordCount []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"format", 0, 0, 0}, {"charlist", 0, 0, 0}}
var ArginfoMoneyFormat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"format", 0, 0, 0},
	{"value", 0, 0, 0},
}
var ArginfoStrSplit []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"split_length", 0, 0, 0}}
var ArginfoStrpbrk []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"char_list", 0, 0, 0}}
var ArginfoSubstrCompare []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(3)), 0, 0, 0}, {"main_str", 0, 0, 0}, {"str", 0, 0, 0}, {"offset", 0, 0, 0}, {"length", 0, 0, 0}, {"case_sensitivity", 0, 0, 0}}
var ArginfoUtf8Encode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"data", 0, 0, 0}}
var ArginfoUtf8Decode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"data", 0, 0, 0}}

/* }}} */

var ArginfoOpenlog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"ident", 0, 0, 0},
	{"option", 0, 0, 0},
	{"facility", 0, 0, 0},
}
var ArginfoCloselog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoSyslog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"priority", 0, 0, 0},
	{"message", 0, 0, 0},
}

/* }}} */

var ArginfoGettype []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}
var ArginfoSettype []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 1, 0},
	{"type", 0, 0, 0},
}
var ArginfoIntval []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var", 0, 0, 0}, {"base", 0, 0, 0}}
var ArginfoFloatval []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}
var ArginfoStrval []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}
var ArginfoBoolval []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}
var ArginfoIsNull []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}
var ArginfoIsResource []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}
var ArginfoIsBool []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}
var ArginfoIsInt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}
var ArginfoIsFloat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}
var ArginfoIsString []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}
var ArginfoIsArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}
var ArginfoIsObject []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}
var ArginfoIsNumeric []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"value", 0, 0, 0},
}
var ArginfoIsScalar []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"value", 0, 0, 0},
}
var ArginfoIsCallable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var", 0, 0, 0}, {"syntax_only", 0, 0, 0}, {"callable_name", 0, 1, 0}}
var ArginfoIsIterable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoIsCountable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}

/* }}} */

var ArginfoUniqid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"prefix", 0, 0, 0}, {"more_entropy", 0, 0, 0}}

/* }}} */

var ArginfoParseUrl []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"url", 0, 0, 0}, {"component", 0, 0, 0}}
var ArginfoUrlencode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoUrldecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoRawurlencode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoRawurldecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"str", 0, 0, 0},
}
var ArginfoGetHeaders []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"url", 0, 0, 0}, {"format", 0, 0, 0}, {"context", 0, 0, 0}}

/* }}} */

var ArginfoStreamBucketMakeWriteable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"brigade", 0, 0, 0},
}
var ArginfoStreamBucketPrepend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"brigade", 0, 0, 0},
	{"bucket", 0, 0, 0},
}
var ArginfoStreamBucketAppend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"brigade", 0, 0, 0},
	{"bucket", 0, 0, 0},
}
var ArginfoStreamBucketNew []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"stream", 0, 0, 0},
	{"buffer", 0, 0, 0},
}
var ArginfoStreamGetFilters []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ArginfoStreamFilterRegister []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"filtername", 0, 0, 0},
	{"classname", 0, 0, 0},
}

/* }}} */

var ArginfoConvertUuencode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"data", 0, 0, 0},
}
var ArginfoConvertUudecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"data", 0, 0, 0},
}

/* }}} */

var ArginfoVarDump []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"vars", 0, 0, 1}}
var ArginfoDebugZvalDump []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"vars", 0, 0, 1}}
var ArginfoVarExport []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var", 0, 0, 0}, {"return", 0, 0, 0}}
var ArginfoSerialize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"var", 0, 0, 0},
}
var ArginfoUnserialize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"variable_representation", 0, 0, 0}, {"allowed_classes", 0, 0, 0}}
var ArginfoMemoryGetUsage []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"real_usage", 0, 0, 0}}
var ArginfoMemoryGetPeakUsage []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"real_usage", 0, 0, 0}}

/* }}} */

var ArginfoVersionCompare []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"ver1", 0, 0, 0}, {"ver2", 0, 0, 0}, {"oper", 0, 0, 0}}

/* }}} */

/* }}} */

var BasicFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"constant",
		ZifConstant,
		ArginfoConstant,
		uint32_t(b.SizeOf("arginfo_constant")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"bin2hex",
		ZifBin2hex,
		ArginfoBin2hex,
		uint32_t(b.SizeOf("arginfo_bin2hex")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"hex2bin",
		ZifHex2bin,
		ArginfoHex2bin,
		uint32_t(b.SizeOf("arginfo_hex2bin")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sleep",
		ZifSleep,
		ArginfoSleep,
		uint32_t(b.SizeOf("arginfo_sleep")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"usleep",
		ZifUsleep,
		ArginfoUsleep,
		uint32_t(b.SizeOf("arginfo_usleep")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"time_nanosleep",
		ZifTimeNanosleep,
		ArginfoTimeNanosleep,
		uint32_t(b.SizeOf("arginfo_time_nanosleep")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"time_sleep_until",
		ZifTimeSleepUntil,
		ArginfoTimeSleepUntil,
		uint32_t(b.SizeOf("arginfo_time_sleep_until")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strptime",
		ZifStrptime,
		ArginfoStrptime,
		uint32_t(b.SizeOf("arginfo_strptime")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"flush",
		ZifFlush,
		ArginfoFlush,
		uint32_t(b.SizeOf("arginfo_flush")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"wordwrap",
		ZifWordwrap,
		ArginfoWordwrap,
		uint32_t(b.SizeOf("arginfo_wordwrap")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"htmlspecialchars",
		ZifHtmlspecialchars,
		ArginfoHtmlspecialchars,
		uint32_t(b.SizeOf("arginfo_htmlspecialchars")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"htmlentities",
		ZifHtmlentities,
		ArginfoHtmlentities,
		uint32_t(b.SizeOf("arginfo_htmlentities")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"html_entity_decode",
		ZifHtmlEntityDecode,
		ArginfoHtmlEntityDecode,
		uint32_t(b.SizeOf("arginfo_html_entity_decode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"htmlspecialchars_decode",
		ZifHtmlspecialcharsDecode,
		ArginfoHtmlspecialcharsDecode,
		uint32_t(b.SizeOf("arginfo_htmlspecialchars_decode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_html_translation_table",
		ZifGetHtmlTranslationTable,
		ArginfoGetHtmlTranslationTable,
		uint32_t(b.SizeOf("arginfo_get_html_translation_table")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sha1",
		ZifSha1,
		ArginfoSha1,
		uint32_t(b.SizeOf("arginfo_sha1")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sha1_file",
		ZifSha1File,
		ArginfoSha1File,
		uint32_t(b.SizeOf("arginfo_sha1_file")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"md5",
		PhpIfMd5,
		ArginfoMd5,
		uint32_t(b.SizeOf("arginfo_md5")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"md5_file",
		PhpIfMd5File,
		ArginfoMd5File,
		uint32_t(b.SizeOf("arginfo_md5_file")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"crc32",
		PhpIfCrc32,
		ArginfoCrc32,
		uint32_t(b.SizeOf("arginfo_crc32")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"iptcparse",
		ZifIptcparse,
		ArginfoIptcparse,
		uint32_t(b.SizeOf("arginfo_iptcparse")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"iptcembed",
		ZifIptcembed,
		ArginfoIptcembed,
		uint32_t(b.SizeOf("arginfo_iptcembed")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getimagesize",
		ZifGetimagesize,
		ArginfoGetimagesize,
		uint32_t(b.SizeOf("arginfo_getimagesize")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getimagesizefromstring",
		ZifGetimagesizefromstring,
		ArginfoGetimagesize,
		uint32_t(b.SizeOf("arginfo_getimagesize")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"image_type_to_mime_type",
		ZifImageTypeToMimeType,
		ArginfoImageTypeToMimeType,
		uint32_t(b.SizeOf("arginfo_image_type_to_mime_type")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"image_type_to_extension",
		ZifImageTypeToExtension,
		ArginfoImageTypeToExtension,
		uint32_t(b.SizeOf("arginfo_image_type_to_extension")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"phpinfo",
		ZifPhpinfo,
		ArginfoPhpinfo,
		uint32_t(b.SizeOf("arginfo_phpinfo")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"phpversion",
		ZifPhpversion,
		ArginfoPhpversion,
		uint32_t(b.SizeOf("arginfo_phpversion")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"phpcredits",
		ZifPhpcredits,
		ArginfoPhpcredits,
		uint32_t(b.SizeOf("arginfo_phpcredits")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"php_sapi_name",
		ZifPhpSapiName,
		ArginfoPhpSapiName,
		uint32_t(b.SizeOf("arginfo_php_sapi_name")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"php_uname",
		ZifPhpUname,
		ArginfoPhpUname,
		uint32_t(b.SizeOf("arginfo_php_uname")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"php_ini_scanned_files",
		ZifPhpIniScannedFiles,
		ArginfoPhpIniScannedFiles,
		uint32_t(b.SizeOf("arginfo_php_ini_scanned_files")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"php_ini_loaded_file",
		ZifPhpIniLoadedFile,
		ArginfoPhpIniLoadedFile,
		uint32_t(b.SizeOf("arginfo_php_ini_loaded_file")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strnatcmp",
		ZifStrnatcmp,
		ArginfoStrnatcmp,
		uint32_t(b.SizeOf("arginfo_strnatcmp")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strnatcasecmp",
		ZifStrnatcasecmp,
		ArginfoStrnatcasecmp,
		uint32_t(b.SizeOf("arginfo_strnatcasecmp")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"substr_count",
		ZifSubstrCount,
		ArginfoSubstrCount,
		uint32_t(b.SizeOf("arginfo_substr_count")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strspn",
		ZifStrspn,
		ArginfoStrspn,
		uint32_t(b.SizeOf("arginfo_strspn")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strcspn",
		ZifStrcspn,
		ArginfoStrcspn,
		uint32_t(b.SizeOf("arginfo_strcspn")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strtok",
		ZifStrtok,
		ArginfoStrtok,
		uint32_t(b.SizeOf("arginfo_strtok")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strtoupper",
		ZifStrtoupper,
		ArginfoStrtoupper,
		uint32_t(b.SizeOf("arginfo_strtoupper")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strtolower",
		ZifStrtolower,
		ArginfoStrtolower,
		uint32_t(b.SizeOf("arginfo_strtolower")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strpos",
		ZifStrpos,
		ArginfoStrpos,
		uint32_t(b.SizeOf("arginfo_strpos")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stripos",
		ZifStripos,
		ArginfoStripos,
		uint32_t(b.SizeOf("arginfo_stripos")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strrpos",
		ZifStrrpos,
		ArginfoStrrpos,
		uint32_t(b.SizeOf("arginfo_strrpos")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strripos",
		ZifStrripos,
		ArginfoStrripos,
		uint32_t(b.SizeOf("arginfo_strripos")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strrev",
		ZifStrrev,
		ArginfoStrrev,
		uint32_t(b.SizeOf("arginfo_strrev")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"hebrev",
		ZifHebrev,
		ArginfoHebrev,
		uint32_t(b.SizeOf("arginfo_hebrev")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"hebrevc",
		ZifHebrevc,
		ArginfoHebrevc,
		uint32_t(b.SizeOf("arginfo_hebrevc")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_DEPRECATED,
	},
	{
		"nl2br",
		ZifNl2br,
		ArginfoNl2br,
		uint32_t(b.SizeOf("arginfo_nl2br")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"basename",
		ZifBasename,
		ArginfoBasename,
		uint32_t(b.SizeOf("arginfo_basename")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"dirname",
		ZifDirname,
		ArginfoDirname,
		uint32_t(b.SizeOf("arginfo_dirname")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pathinfo",
		ZifPathinfo,
		ArginfoPathinfo,
		uint32_t(b.SizeOf("arginfo_pathinfo")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stripslashes",
		ZifStripslashes,
		ArginfoStripslashes,
		uint32_t(b.SizeOf("arginfo_stripslashes")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stripcslashes",
		ZifStripcslashes,
		ArginfoStripcslashes,
		uint32_t(b.SizeOf("arginfo_stripcslashes")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strstr",
		ZifStrstr,
		ArginfoStrstr,
		uint32_t(b.SizeOf("arginfo_strstr")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stristr",
		ZifStristr,
		ArginfoStristr,
		uint32_t(b.SizeOf("arginfo_stristr")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strrchr",
		ZifStrrchr,
		ArginfoStrrchr,
		uint32_t(b.SizeOf("arginfo_strrchr")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_shuffle",
		ZifStrShuffle,
		ArginfoStrShuffle,
		uint32_t(b.SizeOf("arginfo_str_shuffle")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_word_count",
		ZifStrWordCount,
		ArginfoStrWordCount,
		uint32_t(b.SizeOf("arginfo_str_word_count")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_split",
		ZifStrSplit,
		ArginfoStrSplit,
		uint32_t(b.SizeOf("arginfo_str_split")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strpbrk",
		ZifStrpbrk,
		ArginfoStrpbrk,
		uint32_t(b.SizeOf("arginfo_strpbrk")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"substr_compare",
		ZifSubstrCompare,
		ArginfoSubstrCompare,
		uint32_t(b.SizeOf("arginfo_substr_compare")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"utf8_encode",
		ZifUtf8Encode,
		ArginfoUtf8Encode,
		uint32_t(b.SizeOf("arginfo_utf8_encode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"utf8_decode",
		ZifUtf8Decode,
		ArginfoUtf8Decode,
		uint32_t(b.SizeOf("arginfo_utf8_decode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strcoll",
		ZifStrcoll,
		ArginfoStrcoll,
		uint32_t(b.SizeOf("arginfo_strcoll")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"money_format",
		ZifMoneyFormat,
		ArginfoMoneyFormat,
		uint32_t(b.SizeOf("arginfo_money_format")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_DEPRECATED,
	},
	{
		"substr",
		ZifSubstr,
		ArginfoSubstr,
		uint32_t(b.SizeOf("arginfo_substr")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"substr_replace",
		ZifSubstrReplace,
		ArginfoSubstrReplace,
		uint32_t(b.SizeOf("arginfo_substr_replace")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"quotemeta",
		ZifQuotemeta,
		ArginfoQuotemeta,
		uint32_t(b.SizeOf("arginfo_quotemeta")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ucfirst",
		ZifUcfirst,
		ArginfoUcfirst,
		uint32_t(b.SizeOf("arginfo_ucfirst")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"lcfirst",
		ZifLcfirst,
		ArginfoLcfirst,
		uint32_t(b.SizeOf("arginfo_lcfirst")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ucwords",
		ZifUcwords,
		ArginfoUcwords,
		uint32_t(b.SizeOf("arginfo_ucwords")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strtr",
		ZifStrtr,
		ArginfoStrtr,
		uint32_t(b.SizeOf("arginfo_strtr")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"addslashes",
		ZifAddslashes,
		ArginfoAddslashes,
		uint32_t(b.SizeOf("arginfo_addslashes")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"addcslashes",
		ZifAddcslashes,
		ArginfoAddcslashes,
		uint32_t(b.SizeOf("arginfo_addcslashes")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rtrim",
		ZifRtrim,
		ArginfoRtrim,
		uint32_t(b.SizeOf("arginfo_rtrim")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_replace",
		ZifStrReplace,
		ArginfoStrReplace,
		uint32_t(b.SizeOf("arginfo_str_replace")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_ireplace",
		ZifStrIreplace,
		ArginfoStrIreplace,
		uint32_t(b.SizeOf("arginfo_str_ireplace")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_repeat",
		ZifStrRepeat,
		ArginfoStrRepeat,
		uint32_t(b.SizeOf("arginfo_str_repeat")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"count_chars",
		ZifCountChars,
		ArginfoCountChars,
		uint32_t(b.SizeOf("arginfo_count_chars")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chunk_split",
		ZifChunkSplit,
		ArginfoChunkSplit,
		uint32_t(b.SizeOf("arginfo_chunk_split")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"trim",
		ZifTrim,
		ArginfoTrim,
		uint32_t(b.SizeOf("arginfo_trim")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ltrim",
		ZifLtrim,
		ArginfoLtrim,
		uint32_t(b.SizeOf("arginfo_ltrim")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strip_tags",
		ZifStripTags,
		ArginfoStripTags,
		uint32_t(b.SizeOf("arginfo_strip_tags")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"similar_text",
		ZifSimilarText,
		ArginfoSimilarText,
		uint32_t(b.SizeOf("arginfo_similar_text")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"explode",
		ZifExplode,
		ArginfoExplode,
		uint32_t(b.SizeOf("arginfo_explode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"implode",
		ZifImplode,
		ArginfoImplode,
		uint32_t(b.SizeOf("arginfo_implode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"join",
		ZifImplode,
		ArginfoImplode,
		uint32_t(b.SizeOf("arginfo_implode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"setlocale",
		ZifSetlocale,
		ArginfoSetlocale,
		uint32_t(b.SizeOf("arginfo_setlocale")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"localeconv",
		ZifLocaleconv,
		ArginfoLocaleconv,
		uint32_t(b.SizeOf("arginfo_localeconv")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"soundex",
		ZifSoundex,
		ArginfoSoundex,
		uint32_t(b.SizeOf("arginfo_soundex")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"levenshtein",
		ZifLevenshtein,
		ArginfoLevenshtein,
		uint32_t(b.SizeOf("arginfo_levenshtein")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chr",
		ZifChr,
		ArginfoChr,
		uint32_t(b.SizeOf("arginfo_chr")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ord",
		ZifOrd,
		ArginfoOrd,
		uint32_t(b.SizeOf("arginfo_ord")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"parse_str",
		ZifParseStr,
		ArginfoParseStr,
		uint32_t(b.SizeOf("arginfo_parse_str")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_getcsv",
		ZifStrGetcsv,
		ArginfoStrGetcsv,
		uint32_t(b.SizeOf("arginfo_str_getcsv")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_pad",
		ZifStrPad,
		ArginfoStrPad,
		uint32_t(b.SizeOf("arginfo_str_pad")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chop",
		ZifRtrim,
		ArginfoRtrim,
		uint32_t(b.SizeOf("arginfo_rtrim")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strchr",
		ZifStrstr,
		ArginfoStrstr,
		uint32_t(b.SizeOf("arginfo_strstr")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sprintf",
		ZifUserSprintf,
		ArginfoSprintf,
		uint32_t(b.SizeOf("arginfo_sprintf")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"printf",
		ZifUserPrintf,
		ArginfoPrintf,
		uint32_t(b.SizeOf("arginfo_printf")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"vprintf",
		ZifVprintf,
		ArginfoVprintf,
		uint32_t(b.SizeOf("arginfo_vprintf")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"vsprintf",
		ZifVsprintf,
		ArginfoVsprintf,
		uint32_t(b.SizeOf("arginfo_vsprintf")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fprintf",
		ZifFprintf,
		ArginfoFprintf,
		uint32_t(b.SizeOf("arginfo_fprintf")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"vfprintf",
		ZifVfprintf,
		ArginfoVfprintf,
		uint32_t(b.SizeOf("arginfo_vfprintf")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sscanf",
		ZifSscanf,
		ArginfoSscanf,
		uint32_t(b.SizeOf("arginfo_sscanf")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fscanf",
		ZifFscanf,
		ArginfoFscanf,
		uint32_t(b.SizeOf("arginfo_fscanf")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"parse_url",
		ZifParseUrl,
		ArginfoParseUrl,
		uint32_t(b.SizeOf("arginfo_parse_url")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"urlencode",
		ZifUrlencode,
		ArginfoUrlencode,
		uint32_t(b.SizeOf("arginfo_urlencode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"urldecode",
		ZifUrldecode,
		ArginfoUrldecode,
		uint32_t(b.SizeOf("arginfo_urldecode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rawurlencode",
		ZifRawurlencode,
		ArginfoRawurlencode,
		uint32_t(b.SizeOf("arginfo_rawurlencode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rawurldecode",
		ZifRawurldecode,
		ArginfoRawurldecode,
		uint32_t(b.SizeOf("arginfo_rawurldecode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"http_build_query",
		ZifHttpBuildQuery,
		ArginfoHttpBuildQuery,
		uint32_t(b.SizeOf("arginfo_http_build_query")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"readlink",
		ZifReadlink,
		ArginfoReadlink,
		uint32_t(b.SizeOf("arginfo_readlink")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"linkinfo",
		ZifLinkinfo,
		ArginfoLinkinfo,
		uint32_t(b.SizeOf("arginfo_linkinfo")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"symlink",
		ZifSymlink,
		ArginfoSymlink,
		uint32_t(b.SizeOf("arginfo_symlink")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"link",
		ZifLink,
		ArginfoLink,
		uint32_t(b.SizeOf("arginfo_link")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"unlink",
		ZifUnlink,
		ArginfoUnlink,
		uint32_t(b.SizeOf("arginfo_unlink")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"exec",
		ZifExec,
		ArginfoExec,
		uint32_t(b.SizeOf("arginfo_exec")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"system",
		ZifSystem,
		ArginfoSystem,
		uint32_t(b.SizeOf("arginfo_system")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"escapeshellcmd",
		ZifEscapeshellcmd,
		ArginfoEscapeshellcmd,
		uint32_t(b.SizeOf("arginfo_escapeshellcmd")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"escapeshellarg",
		ZifEscapeshellarg,
		ArginfoEscapeshellarg,
		uint32_t(b.SizeOf("arginfo_escapeshellarg")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"passthru",
		ZifPassthru,
		ArginfoPassthru,
		uint32_t(b.SizeOf("arginfo_passthru")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"shell_exec",
		ZifShellExec,
		ArginfoShellExec,
		uint32_t(b.SizeOf("arginfo_shell_exec")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"proc_open",
		ZifProcOpen,
		ArginfoProcOpen,
		uint32_t(b.SizeOf("arginfo_proc_open")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"proc_close",
		ZifProcClose,
		ArginfoProcClose,
		uint32_t(b.SizeOf("arginfo_proc_close")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"proc_terminate",
		ZifProcTerminate,
		ArginfoProcTerminate,
		uint32_t(b.SizeOf("arginfo_proc_terminate")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"proc_get_status",
		ZifProcGetStatus,
		ArginfoProcGetStatus,
		uint32_t(b.SizeOf("arginfo_proc_get_status")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"proc_nice",
		ZifProcNice,
		ArginfoProcNice,
		uint32_t(b.SizeOf("arginfo_proc_nice")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rand",
		ZifRand,
		ArginfoMtRand,
		uint32_t(b.SizeOf("arginfo_mt_rand")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"srand",
		ZifMtSrand,
		ArginfoMtSrand,
		uint32_t(b.SizeOf("arginfo_mt_srand")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getrandmax",
		ZifMtGetrandmax,
		ArginfoMtGetrandmax,
		uint32_t(b.SizeOf("arginfo_mt_getrandmax")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"mt_rand",
		ZifMtRand,
		ArginfoMtRand,
		uint32_t(b.SizeOf("arginfo_mt_rand")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"mt_srand",
		ZifMtSrand,
		ArginfoMtSrand,
		uint32_t(b.SizeOf("arginfo_mt_srand")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"mt_getrandmax",
		ZifMtGetrandmax,
		ArginfoMtGetrandmax,
		uint32_t(b.SizeOf("arginfo_mt_getrandmax")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"random_bytes",
		ZifRandomBytes,
		ArginfoRandomBytes,
		uint32_t(b.SizeOf("arginfo_random_bytes")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"random_int",
		ZifRandomInt,
		ArginfoRandomInt,
		uint32_t(b.SizeOf("arginfo_random_int")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getservbyname",
		ZifGetservbyname,
		ArginfoGetservbyname,
		uint32_t(b.SizeOf("arginfo_getservbyname")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getservbyport",
		ZifGetservbyport,
		ArginfoGetservbyport,
		uint32_t(b.SizeOf("arginfo_getservbyport")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getprotobyname",
		ZifGetprotobyname,
		ArginfoGetprotobyname,
		uint32_t(b.SizeOf("arginfo_getprotobyname")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getprotobynumber",
		ZifGetprotobynumber,
		ArginfoGetprotobynumber,
		uint32_t(b.SizeOf("arginfo_getprotobynumber")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getmyuid",
		ZifGetmyuid,
		ArginfoGetmyuid,
		uint32_t(b.SizeOf("arginfo_getmyuid")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getmygid",
		ZifGetmygid,
		ArginfoGetmygid,
		uint32_t(b.SizeOf("arginfo_getmygid")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getmypid",
		ZifGetmypid,
		ArginfoGetmypid,
		uint32_t(b.SizeOf("arginfo_getmypid")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getmyinode",
		ZifGetmyinode,
		ArginfoGetmyinode,
		uint32_t(b.SizeOf("arginfo_getmyinode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getlastmod",
		ZifGetlastmod,
		ArginfoGetlastmod,
		uint32_t(b.SizeOf("arginfo_getlastmod")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"base64_decode",
		ZifBase64Decode,
		ArginfoBase64Decode,
		uint32_t(b.SizeOf("arginfo_base64_decode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"base64_encode",
		ZifBase64Encode,
		ArginfoBase64Encode,
		uint32_t(b.SizeOf("arginfo_base64_encode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"password_hash",
		ZifPasswordHash,
		ArginfoPasswordHash,
		uint32_t(b.SizeOf("arginfo_password_hash")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"password_get_info",
		ZifPasswordGetInfo,
		ArginfoPasswordGetInfo,
		uint32_t(b.SizeOf("arginfo_password_get_info")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"password_needs_rehash",
		ZifPasswordNeedsRehash,
		ArginfoPasswordNeedsRehash,
		uint32_t(b.SizeOf("arginfo_password_needs_rehash")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"password_verify",
		ZifPasswordVerify,
		ArginfoPasswordVerify,
		uint32_t(b.SizeOf("arginfo_password_verify")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"password_algos",
		ZifPasswordAlgos,
		ArginfoPasswordAlgos,
		uint32_t(b.SizeOf("arginfo_password_algos")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"convert_uuencode",
		ZifConvertUuencode,
		ArginfoConvertUuencode,
		uint32_t(b.SizeOf("arginfo_convert_uuencode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"convert_uudecode",
		ZifConvertUudecode,
		ArginfoConvertUudecode,
		uint32_t(b.SizeOf("arginfo_convert_uudecode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"abs",
		ZifAbs,
		ArginfoAbs,
		uint32_t(b.SizeOf("arginfo_abs")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ceil",
		ZifCeil,
		ArginfoCeil,
		uint32_t(b.SizeOf("arginfo_ceil")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"floor",
		ZifFloor,
		ArginfoFloor,
		uint32_t(b.SizeOf("arginfo_floor")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"round",
		ZifRound,
		ArginfoRound,
		uint32_t(b.SizeOf("arginfo_round")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sin",
		ZifSin,
		ArginfoSin,
		uint32_t(b.SizeOf("arginfo_sin")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"cos",
		ZifCos,
		ArginfoCos,
		uint32_t(b.SizeOf("arginfo_cos")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"tan",
		ZifTan,
		ArginfoTan,
		uint32_t(b.SizeOf("arginfo_tan")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"asin",
		ZifAsin,
		ArginfoAsin,
		uint32_t(b.SizeOf("arginfo_asin")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"acos",
		ZifAcos,
		ArginfoAcos,
		uint32_t(b.SizeOf("arginfo_acos")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"atan",
		ZifAtan,
		ArginfoAtan,
		uint32_t(b.SizeOf("arginfo_atan")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"atanh",
		ZifAtanh,
		ArginfoAtanh,
		uint32_t(b.SizeOf("arginfo_atanh")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"atan2",
		ZifAtan2,
		ArginfoAtan2,
		uint32_t(b.SizeOf("arginfo_atan2")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sinh",
		ZifSinh,
		ArginfoSinh,
		uint32_t(b.SizeOf("arginfo_sinh")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"cosh",
		ZifCosh,
		ArginfoCosh,
		uint32_t(b.SizeOf("arginfo_cosh")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"tanh",
		ZifTanh,
		ArginfoTanh,
		uint32_t(b.SizeOf("arginfo_tanh")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"asinh",
		ZifAsinh,
		ArginfoAsinh,
		uint32_t(b.SizeOf("arginfo_asinh")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"acosh",
		ZifAcosh,
		ArginfoAcosh,
		uint32_t(b.SizeOf("arginfo_acosh")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"expm1",
		ZifExpm1,
		ArginfoExpm1,
		uint32_t(b.SizeOf("arginfo_expm1")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"log1p",
		ZifLog1p,
		ArginfoLog1p,
		uint32_t(b.SizeOf("arginfo_log1p")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pi",
		ZifPi,
		ArginfoPi,
		uint32_t(b.SizeOf("arginfo_pi")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_finite",
		ZifIsFinite,
		ArginfoIsFinite,
		uint32_t(b.SizeOf("arginfo_is_finite")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_nan",
		ZifIsNan,
		ArginfoIsNan,
		uint32_t(b.SizeOf("arginfo_is_nan")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_infinite",
		ZifIsInfinite,
		ArginfoIsInfinite,
		uint32_t(b.SizeOf("arginfo_is_infinite")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pow",
		ZifPow,
		ArginfoPow,
		uint32_t(b.SizeOf("arginfo_pow")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"exp",
		ZifExp,
		ArginfoExp,
		uint32_t(b.SizeOf("arginfo_exp")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"log",
		ZifLog,
		ArginfoLog,
		uint32_t(b.SizeOf("arginfo_log")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"log10",
		ZifLog10,
		ArginfoLog10,
		uint32_t(b.SizeOf("arginfo_log10")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sqrt",
		ZifSqrt,
		ArginfoSqrt,
		uint32_t(b.SizeOf("arginfo_sqrt")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"hypot",
		ZifHypot,
		ArginfoHypot,
		uint32_t(b.SizeOf("arginfo_hypot")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"deg2rad",
		ZifDeg2rad,
		ArginfoDeg2rad,
		uint32_t(b.SizeOf("arginfo_deg2rad")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rad2deg",
		ZifRad2deg,
		ArginfoRad2deg,
		uint32_t(b.SizeOf("arginfo_rad2deg")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"bindec",
		ZifBindec,
		ArginfoBindec,
		uint32_t(b.SizeOf("arginfo_bindec")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"hexdec",
		ZifHexdec,
		ArginfoHexdec,
		uint32_t(b.SizeOf("arginfo_hexdec")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"octdec",
		ZifOctdec,
		ArginfoOctdec,
		uint32_t(b.SizeOf("arginfo_octdec")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"decbin",
		ZifDecbin,
		ArginfoDecbin,
		uint32_t(b.SizeOf("arginfo_decbin")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"decoct",
		ZifDecoct,
		ArginfoDecoct,
		uint32_t(b.SizeOf("arginfo_decoct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"dechex",
		ZifDechex,
		ArginfoDechex,
		uint32_t(b.SizeOf("arginfo_dechex")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"base_convert",
		ZifBaseConvert,
		ArginfoBaseConvert,
		uint32_t(b.SizeOf("arginfo_base_convert")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"number_format",
		ZifNumberFormat,
		ArginfoNumberFormat,
		uint32_t(b.SizeOf("arginfo_number_format")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fmod",
		ZifFmod,
		ArginfoFmod,
		uint32_t(b.SizeOf("arginfo_fmod")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"intdiv",
		ZifIntdiv,
		ArginfoIntdiv,
		uint32_t(b.SizeOf("arginfo_intdiv")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"inet_ntop",
		ZifInetNtop,
		ArginfoInetNtop,
		uint32_t(b.SizeOf("arginfo_inet_ntop")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"inet_pton",
		PhpInetPton,
		ArginfoInetPton,
		uint32_t(b.SizeOf("arginfo_inet_pton")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ip2long",
		ZifIp2long,
		ArginfoIp2long,
		uint32_t(b.SizeOf("arginfo_ip2long")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"long2ip",
		ZifLong2ip,
		ArginfoLong2ip,
		uint32_t(b.SizeOf("arginfo_long2ip")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getenv",
		ZifGetenv,
		ArginfoGetenv,
		uint32_t(b.SizeOf("arginfo_getenv")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"putenv",
		ZifPutenv,
		ArginfoPutenv,
		uint32_t(b.SizeOf("arginfo_putenv")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getopt",
		ZifGetopt,
		ArginfoGetopt,
		uint32_t(b.SizeOf("arginfo_getopt")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sys_getloadavg",
		ZifSysGetloadavg,
		ArginfoSysGetloadavg,
		uint32_t(b.SizeOf("arginfo_sys_getloadavg")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"microtime",
		ZifMicrotime,
		ArginfoMicrotime,
		uint32_t(b.SizeOf("arginfo_microtime")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gettimeofday",
		ZifGettimeofday,
		ArginfoGettimeofday,
		uint32_t(b.SizeOf("arginfo_gettimeofday")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getrusage",
		ZifGetrusage,
		ArginfoGetrusage,
		uint32_t(b.SizeOf("arginfo_getrusage")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"hrtime",
		ZifHrtime,
		ArginfoHrtime,
		uint32_t(b.SizeOf("arginfo_hrtime")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"uniqid",
		ZifUniqid,
		ArginfoUniqid,
		uint32_t(b.SizeOf("arginfo_uniqid")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"quoted_printable_decode",
		ZifQuotedPrintableDecode,
		ArginfoQuotedPrintableDecode,
		uint32_t(b.SizeOf("arginfo_quoted_printable_decode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"quoted_printable_encode",
		ZifQuotedPrintableEncode,
		ArginfoQuotedPrintableEncode,
		uint32_t(b.SizeOf("arginfo_quoted_printable_encode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"convert_cyr_string",
		ZifConvertCyrString,
		ArginfoConvertCyrString,
		uint32_t(b.SizeOf("arginfo_convert_cyr_string")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_DEPRECATED,
	},
	{
		"get_current_user",
		ZifGetCurrentUser,
		ArginfoGetCurrentUser,
		uint32_t(b.SizeOf("arginfo_get_current_user")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"set_time_limit",
		ZifSetTimeLimit,
		ArginfoSetTimeLimit,
		uint32_t(b.SizeOf("arginfo_set_time_limit")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"header_register_callback",
		ZifHeaderRegisterCallback,
		ArginfoHeaderRegisterCallback,
		uint32_t(b.SizeOf("arginfo_header_register_callback")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_cfg_var",
		ZifGetCfgVar,
		ArginfoGetCfgVar,
		uint32_t(b.SizeOf("arginfo_get_cfg_var")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_magic_quotes_gpc",
		ZifGetMagicQuotesGpc,
		ArginfoGetMagicQuotesGpc,
		uint32_t(b.SizeOf("arginfo_get_magic_quotes_gpc")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_DEPRECATED,
	},
	{
		"get_magic_quotes_runtime",
		ZifGetMagicQuotesRuntime,
		ArginfoGetMagicQuotesRuntime,
		uint32_t(b.SizeOf("arginfo_get_magic_quotes_runtime")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_DEPRECATED,
	},
	{
		"error_log",
		ZifErrorLog,
		ArginfoErrorLog,
		uint32_t(b.SizeOf("arginfo_error_log")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"error_get_last",
		ZifErrorGetLast,
		ArginfoErrorGetLast,
		uint32_t(b.SizeOf("arginfo_error_get_last")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"error_clear_last",
		ZifErrorClearLast,
		ArginfoErrorClearLast,
		uint32_t(b.SizeOf("arginfo_error_clear_last")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"call_user_func",
		ZifCallUserFunc,
		ArginfoCallUserFunc,
		uint32_t(b.SizeOf("arginfo_call_user_func")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"call_user_func_array",
		ZifCallUserFuncArray,
		ArginfoCallUserFuncArray,
		uint32_t(b.SizeOf("arginfo_call_user_func_array")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"forward_static_call",
		ZifForwardStaticCall,
		ArginfoForwardStaticCall,
		uint32_t(b.SizeOf("arginfo_forward_static_call")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"forward_static_call_array",
		ZifForwardStaticCallArray,
		ArginfoForwardStaticCallArray,
		uint32_t(b.SizeOf("arginfo_forward_static_call_array")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"serialize",
		ZifSerialize,
		ArginfoSerialize,
		uint32_t(b.SizeOf("arginfo_serialize")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"unserialize",
		ZifUnserialize,
		ArginfoUnserialize,
		uint32_t(b.SizeOf("arginfo_unserialize")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"var_dump",
		ZifVarDump,
		ArginfoVarDump,
		uint32_t(b.SizeOf("arginfo_var_dump")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"var_export",
		ZifVarExport,
		ArginfoVarExport,
		uint32_t(b.SizeOf("arginfo_var_export")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"debug_zval_dump",
		ZifDebugZvalDump,
		ArginfoDebugZvalDump,
		uint32_t(b.SizeOf("arginfo_debug_zval_dump")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"print_r",
		ZifPrintR,
		ArginfoPrintR,
		uint32_t(b.SizeOf("arginfo_print_r")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"memory_get_usage",
		ZifMemoryGetUsage,
		ArginfoMemoryGetUsage,
		uint32_t(b.SizeOf("arginfo_memory_get_usage")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"memory_get_peak_usage",
		ZifMemoryGetPeakUsage,
		ArginfoMemoryGetPeakUsage,
		uint32_t(b.SizeOf("arginfo_memory_get_peak_usage")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"register_shutdown_function",
		ZifRegisterShutdownFunction,
		ArginfoRegisterShutdownFunction,
		uint32_t(b.SizeOf("arginfo_register_shutdown_function")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"register_tick_function",
		ZifRegisterTickFunction,
		ArginfoRegisterTickFunction,
		uint32_t(b.SizeOf("arginfo_register_tick_function")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"unregister_tick_function",
		ZifUnregisterTickFunction,
		ArginfoUnregisterTickFunction,
		uint32_t(b.SizeOf("arginfo_unregister_tick_function")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"highlight_file",
		ZifHighlightFile,
		ArginfoHighlightFile,
		uint32_t(b.SizeOf("arginfo_highlight_file")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"show_source",
		ZifHighlightFile,
		ArginfoHighlightFile,
		uint32_t(b.SizeOf("arginfo_highlight_file")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"highlight_string",
		ZifHighlightString,
		ArginfoHighlightString,
		uint32_t(b.SizeOf("arginfo_highlight_string")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"php_strip_whitespace",
		ZifPhpStripWhitespace,
		ArginfoPhpStripWhitespace,
		uint32_t(b.SizeOf("arginfo_php_strip_whitespace")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ini_get",
		ZifIniGet,
		ArginfoIniGet,
		uint32_t(b.SizeOf("arginfo_ini_get")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ini_get_all",
		ZifIniGetAll,
		ArginfoIniGetAll,
		uint32_t(b.SizeOf("arginfo_ini_get_all")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ini_set",
		ZifIniSet,
		ArginfoIniSet,
		uint32_t(b.SizeOf("arginfo_ini_set")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ini_alter",
		ZifIniSet,
		ArginfoIniSet,
		uint32_t(b.SizeOf("arginfo_ini_set")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ini_restore",
		ZifIniRestore,
		ArginfoIniRestore,
		uint32_t(b.SizeOf("arginfo_ini_restore")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_include_path",
		ZifGetIncludePath,
		ArginfoGetIncludePath,
		uint32_t(b.SizeOf("arginfo_get_include_path")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"set_include_path",
		ZifSetIncludePath,
		ArginfoSetIncludePath,
		uint32_t(b.SizeOf("arginfo_set_include_path")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"restore_include_path",
		ZifRestoreIncludePath,
		ArginfoRestoreIncludePath,
		uint32_t(b.SizeOf("arginfo_restore_include_path")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_DEPRECATED,
	},
	{
		"setcookie",
		ZifSetcookie,
		ArginfoSetcookie,
		uint32_t(b.SizeOf("arginfo_setcookie")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"setrawcookie",
		ZifSetrawcookie,
		ArginfoSetrawcookie,
		uint32_t(b.SizeOf("arginfo_setrawcookie")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"header",
		ZifHeader,
		ArginfoHeader,
		uint32_t(b.SizeOf("arginfo_header")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"header_remove",
		ZifHeaderRemove,
		ArginfoHeaderRemove,
		uint32_t(b.SizeOf("arginfo_header_remove")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"headers_sent",
		ZifHeadersSent,
		ArginfoHeadersSent,
		uint32_t(b.SizeOf("arginfo_headers_sent")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"headers_list",
		ZifHeadersList,
		ArginfoHeadersList,
		uint32_t(b.SizeOf("arginfo_headers_list")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"http_response_code",
		ZifHttpResponseCode,
		ArginfoHttpResponseCode,
		uint32_t(b.SizeOf("arginfo_http_response_code")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"connection_aborted",
		ZifConnectionAborted,
		ArginfoConnectionAborted,
		uint32_t(b.SizeOf("arginfo_connection_aborted")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"connection_status",
		ZifConnectionStatus,
		ArginfoConnectionStatus,
		uint32_t(b.SizeOf("arginfo_connection_status")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ignore_user_abort",
		ZifIgnoreUserAbort,
		ArginfoIgnoreUserAbort,
		uint32_t(b.SizeOf("arginfo_ignore_user_abort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"parse_ini_file",
		ZifParseIniFile,
		ArginfoParseIniFile,
		uint32_t(b.SizeOf("arginfo_parse_ini_file")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"parse_ini_string",
		ZifParseIniString,
		ArginfoParseIniString,
		uint32_t(b.SizeOf("arginfo_parse_ini_string")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_uploaded_file",
		ZifIsUploadedFile,
		ArginfoIsUploadedFile,
		uint32_t(b.SizeOf("arginfo_is_uploaded_file")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"move_uploaded_file",
		ZifMoveUploadedFile,
		ArginfoMoveUploadedFile,
		uint32_t(b.SizeOf("arginfo_move_uploaded_file")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gethostbyaddr",
		ZifGethostbyaddr,
		ArginfoGethostbyaddr,
		uint32_t(b.SizeOf("arginfo_gethostbyaddr")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gethostbyname",
		ZifGethostbyname,
		ArginfoGethostbyname,
		uint32_t(b.SizeOf("arginfo_gethostbyname")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gethostbynamel",
		ZifGethostbynamel,
		ArginfoGethostbynamel,
		uint32_t(b.SizeOf("arginfo_gethostbynamel")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gethostname",
		ZifGethostname,
		ArginfoGethostname,
		uint32_t(b.SizeOf("arginfo_gethostname")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"net_get_interfaces",
		ZifNetGetInterfaces,
		ArginfoNetGetInterfaces,
		uint32_t(b.SizeOf("arginfo_net_get_interfaces")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"dns_check_record",
		ZifDnsCheckRecord,
		ArginfoDnsCheckRecord,
		uint32_t(b.SizeOf("arginfo_dns_check_record")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"checkdnsrr",
		ZifDnsCheckRecord,
		ArginfoDnsCheckRecord,
		uint32_t(b.SizeOf("arginfo_dns_check_record")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"dns_get_mx",
		ZifDnsGetMx,
		ArginfoDnsGetMx,
		uint32_t(b.SizeOf("arginfo_dns_get_mx")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getmxrr",
		ZifDnsGetMx,
		ArginfoDnsGetMx,
		uint32_t(b.SizeOf("arginfo_dns_get_mx")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"dns_get_record",
		ZifDnsGetRecord,
		ArginfoDnsGetRecord,
		uint32_t(b.SizeOf("arginfo_dns_get_record")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"intval",
		ZifIntval,
		ArginfoIntval,
		uint32_t(b.SizeOf("arginfo_intval")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"floatval",
		ZifFloatval,
		ArginfoFloatval,
		uint32_t(b.SizeOf("arginfo_floatval")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"doubleval",
		ZifFloatval,
		ArginfoFloatval,
		uint32_t(b.SizeOf("arginfo_floatval")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strval",
		ZifStrval,
		ArginfoStrval,
		uint32_t(b.SizeOf("arginfo_strval")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"boolval",
		ZifBoolval,
		ArginfoBoolval,
		uint32_t(b.SizeOf("arginfo_boolval")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gettype",
		ZifGettype,
		ArginfoGettype,
		uint32_t(b.SizeOf("arginfo_gettype")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"settype",
		ZifSettype,
		ArginfoSettype,
		uint32_t(b.SizeOf("arginfo_settype")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_null",
		ZifIsNull,
		ArginfoIsNull,
		uint32_t(b.SizeOf("arginfo_is_null")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_resource",
		ZifIsResource,
		ArginfoIsResource,
		uint32_t(b.SizeOf("arginfo_is_resource")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_bool",
		ZifIsBool,
		ArginfoIsBool,
		uint32_t(b.SizeOf("arginfo_is_bool")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_int",
		ZifIsInt,
		ArginfoIsInt,
		uint32_t(b.SizeOf("arginfo_is_int")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_float",
		ZifIsFloat,
		ArginfoIsFloat,
		uint32_t(b.SizeOf("arginfo_is_float")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_integer",
		ZifIsInt,
		ArginfoIsInt,
		uint32_t(b.SizeOf("arginfo_is_int")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_long",
		ZifIsInt,
		ArginfoIsInt,
		uint32_t(b.SizeOf("arginfo_is_int")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_double",
		ZifIsFloat,
		ArginfoIsFloat,
		uint32_t(b.SizeOf("arginfo_is_float")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_real",
		ZifIsFloat,
		ArginfoIsFloat,
		uint32_t(b.SizeOf("arginfo_is_float")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_DEPRECATED,
	},
	{
		"is_numeric",
		ZifIsNumeric,
		ArginfoIsNumeric,
		uint32_t(b.SizeOf("arginfo_is_numeric")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_string",
		ZifIsString,
		ArginfoIsString,
		uint32_t(b.SizeOf("arginfo_is_string")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_array",
		ZifIsArray,
		ArginfoIsArray,
		uint32_t(b.SizeOf("arginfo_is_array")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_object",
		ZifIsObject,
		ArginfoIsObject,
		uint32_t(b.SizeOf("arginfo_is_object")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_scalar",
		ZifIsScalar,
		ArginfoIsScalar,
		uint32_t(b.SizeOf("arginfo_is_scalar")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_callable",
		ZifIsCallable,
		ArginfoIsCallable,
		uint32_t(b.SizeOf("arginfo_is_callable")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_iterable",
		ZifIsIterable,
		ArginfoIsIterable,
		uint32_t(b.SizeOf("arginfo_is_iterable")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_countable",
		ZifIsCountable,
		ArginfoIsCountable,
		uint32_t(b.SizeOf("arginfo_is_countable")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pclose",
		ZifPclose,
		ArginfoPclose,
		uint32_t(b.SizeOf("arginfo_pclose")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"popen",
		ZifPopen,
		ArginfoPopen,
		uint32_t(b.SizeOf("arginfo_popen")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"readfile",
		ZifReadfile,
		ArginfoReadfile,
		uint32_t(b.SizeOf("arginfo_readfile")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rewind",
		ZifRewind,
		ArginfoRewind,
		uint32_t(b.SizeOf("arginfo_rewind")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rmdir",
		ZifRmdir,
		ArginfoRmdir,
		uint32_t(b.SizeOf("arginfo_rmdir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"umask",
		ZifUmask,
		ArginfoUmask,
		uint32_t(b.SizeOf("arginfo_umask")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fclose",
		ZifFclose,
		ArginfoFclose,
		uint32_t(b.SizeOf("arginfo_fclose")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"feof",
		ZifFeof,
		ArginfoFeof,
		uint32_t(b.SizeOf("arginfo_feof")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fgetc",
		ZifFgetc,
		ArginfoFgetc,
		uint32_t(b.SizeOf("arginfo_fgetc")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fgets",
		ZifFgets,
		ArginfoFgets,
		uint32_t(b.SizeOf("arginfo_fgets")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fgetss",
		ZifFgetss,
		ArginfoFgetss,
		uint32_t(b.SizeOf("arginfo_fgetss")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_DEPRECATED,
	},
	{
		"fread",
		ZifFread,
		ArginfoFread,
		uint32_t(b.SizeOf("arginfo_fread")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fopen",
		PhpIfFopen,
		ArginfoFopen,
		uint32_t(b.SizeOf("arginfo_fopen")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fpassthru",
		ZifFpassthru,
		ArginfoFpassthru,
		uint32_t(b.SizeOf("arginfo_fpassthru")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ftruncate",
		PhpIfFtruncate,
		ArginfoFtruncate,
		uint32_t(b.SizeOf("arginfo_ftruncate")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fstat",
		PhpIfFstat,
		ArginfoFstat,
		uint32_t(b.SizeOf("arginfo_fstat")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fseek",
		ZifFseek,
		ArginfoFseek,
		uint32_t(b.SizeOf("arginfo_fseek")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ftell",
		ZifFtell,
		ArginfoFtell,
		uint32_t(b.SizeOf("arginfo_ftell")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fflush",
		ZifFflush,
		ArginfoFflush,
		uint32_t(b.SizeOf("arginfo_fflush")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fwrite",
		ZifFwrite,
		ArginfoFwrite,
		uint32_t(b.SizeOf("arginfo_fwrite")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fputs",
		ZifFwrite,
		ArginfoFwrite,
		uint32_t(b.SizeOf("arginfo_fwrite")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"mkdir",
		ZifMkdir,
		ArginfoMkdir,
		uint32_t(b.SizeOf("arginfo_mkdir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rename",
		ZifRename,
		ArginfoRename,
		uint32_t(b.SizeOf("arginfo_rename")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"copy",
		ZifCopy,
		ArginfoCopy,
		uint32_t(b.SizeOf("arginfo_copy")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"tempnam",
		ZifTempnam,
		ArginfoTempnam,
		uint32_t(b.SizeOf("arginfo_tempnam")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"tmpfile",
		PhpIfTmpfile,
		ArginfoTmpfile,
		uint32_t(b.SizeOf("arginfo_tmpfile")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"file",
		ZifFile,
		ArginfoFile,
		uint32_t(b.SizeOf("arginfo_file")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"file_get_contents",
		ZifFileGetContents,
		ArginfoFileGetContents,
		uint32_t(b.SizeOf("arginfo_file_get_contents")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"file_put_contents",
		ZifFilePutContents,
		ArginfoFilePutContents,
		uint32_t(b.SizeOf("arginfo_file_put_contents")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_select",
		ZifStreamSelect,
		ArginfoStreamSelect,
		uint32_t(b.SizeOf("arginfo_stream_select")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_create",
		ZifStreamContextCreate,
		ArginfoStreamContextCreate,
		uint32_t(b.SizeOf("arginfo_stream_context_create")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_set_params",
		ZifStreamContextSetParams,
		ArginfoStreamContextSetParams,
		uint32_t(b.SizeOf("arginfo_stream_context_set_params")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_get_params",
		ZifStreamContextGetParams,
		ArginfoStreamContextGetParams,
		uint32_t(b.SizeOf("arginfo_stream_context_get_params")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_set_option",
		ZifStreamContextSetOption,
		ArginfoStreamContextSetOption,
		uint32_t(b.SizeOf("arginfo_stream_context_set_option")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_get_options",
		ZifStreamContextGetOptions,
		ArginfoStreamContextGetOptions,
		uint32_t(b.SizeOf("arginfo_stream_context_get_options")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_get_default",
		ZifStreamContextGetDefault,
		ArginfoStreamContextGetDefault,
		uint32_t(b.SizeOf("arginfo_stream_context_get_default")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_set_default",
		ZifStreamContextSetDefault,
		ArginfoStreamContextSetDefault,
		uint32_t(b.SizeOf("arginfo_stream_context_set_default")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_filter_prepend",
		ZifStreamFilterPrepend,
		ArginfoStreamFilterPrepend,
		uint32_t(b.SizeOf("arginfo_stream_filter_prepend")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_filter_append",
		ZifStreamFilterAppend,
		ArginfoStreamFilterAppend,
		uint32_t(b.SizeOf("arginfo_stream_filter_append")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_filter_remove",
		ZifStreamFilterRemove,
		ArginfoStreamFilterRemove,
		uint32_t(b.SizeOf("arginfo_stream_filter_remove")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_client",
		ZifStreamSocketClient,
		ArginfoStreamSocketClient,
		uint32_t(b.SizeOf("arginfo_stream_socket_client")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_server",
		ZifStreamSocketServer,
		ArginfoStreamSocketServer,
		uint32_t(b.SizeOf("arginfo_stream_socket_server")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_accept",
		ZifStreamSocketAccept,
		ArginfoStreamSocketAccept,
		uint32_t(b.SizeOf("arginfo_stream_socket_accept")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_get_name",
		ZifStreamSocketGetName,
		ArginfoStreamSocketGetName,
		uint32_t(b.SizeOf("arginfo_stream_socket_get_name")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_recvfrom",
		ZifStreamSocketRecvfrom,
		ArginfoStreamSocketRecvfrom,
		uint32_t(b.SizeOf("arginfo_stream_socket_recvfrom")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_sendto",
		ZifStreamSocketSendto,
		ArginfoStreamSocketSendto,
		uint32_t(b.SizeOf("arginfo_stream_socket_sendto")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_enable_crypto",
		ZifStreamSocketEnableCrypto,
		ArginfoStreamSocketEnableCrypto,
		uint32_t(b.SizeOf("arginfo_stream_socket_enable_crypto")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_shutdown",
		ZifStreamSocketShutdown,
		ArginfoStreamSocketShutdown,
		uint32_t(b.SizeOf("arginfo_stream_socket_shutdown")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_pair",
		ZifStreamSocketPair,
		ArginfoStreamSocketPair,
		uint32_t(b.SizeOf("arginfo_stream_socket_pair")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_copy_to_stream",
		ZifStreamCopyToStream,
		ArginfoStreamCopyToStream,
		uint32_t(b.SizeOf("arginfo_stream_copy_to_stream")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_get_contents",
		ZifStreamGetContents,
		ArginfoStreamGetContents,
		uint32_t(b.SizeOf("arginfo_stream_get_contents")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_supports_lock",
		ZifStreamSupportsLock,
		ArginfoStreamSupportsLock,
		uint32_t(b.SizeOf("arginfo_stream_supports_lock")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_isatty",
		ZifStreamIsatty,
		ArginfoStreamIsatty,
		uint32_t(b.SizeOf("arginfo_stream_isatty")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fgetcsv",
		ZifFgetcsv,
		ArginfoFgetcsv,
		uint32_t(b.SizeOf("arginfo_fgetcsv")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fputcsv",
		ZifFputcsv,
		ArginfoFputcsv,
		uint32_t(b.SizeOf("arginfo_fputcsv")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"flock",
		ZifFlock,
		ArginfoFlock,
		uint32_t(b.SizeOf("arginfo_flock")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_meta_tags",
		ZifGetMetaTags,
		ArginfoGetMetaTags,
		uint32_t(b.SizeOf("arginfo_get_meta_tags")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_set_read_buffer",
		ZifStreamSetReadBuffer,
		ArginfoStreamSetReadBuffer,
		uint32_t(b.SizeOf("arginfo_stream_set_read_buffer")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_set_write_buffer",
		ZifStreamSetWriteBuffer,
		ArginfoStreamSetWriteBuffer,
		uint32_t(b.SizeOf("arginfo_stream_set_write_buffer")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"set_file_buffer",
		ZifStreamSetWriteBuffer,
		ArginfoStreamSetWriteBuffer,
		uint32_t(b.SizeOf("arginfo_stream_set_write_buffer")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_set_chunk_size",
		ZifStreamSetChunkSize,
		ArginfoStreamSetChunkSize,
		uint32_t(b.SizeOf("arginfo_stream_set_chunk_size")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_set_blocking",
		ZifStreamSetBlocking,
		ArginfoStreamSetBlocking,
		uint32_t(b.SizeOf("arginfo_stream_set_blocking")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"socket_set_blocking",
		ZifStreamSetBlocking,
		ArginfoStreamSetBlocking,
		uint32_t(b.SizeOf("arginfo_stream_set_blocking")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_get_meta_data",
		ZifStreamGetMetaData,
		ArginfoStreamGetMetaData,
		uint32_t(b.SizeOf("arginfo_stream_get_meta_data")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_get_line",
		ZifStreamGetLine,
		ArginfoStreamGetLine,
		uint32_t(b.SizeOf("arginfo_stream_get_line")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_wrapper_register",
		ZifStreamWrapperRegister,
		ArginfoStreamWrapperRegister,
		uint32_t(b.SizeOf("arginfo_stream_wrapper_register")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_register_wrapper",
		ZifStreamWrapperRegister,
		ArginfoStreamWrapperRegister,
		uint32_t(b.SizeOf("arginfo_stream_wrapper_register")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_wrapper_unregister",
		ZifStreamWrapperUnregister,
		ArginfoStreamWrapperUnregister,
		uint32_t(b.SizeOf("arginfo_stream_wrapper_unregister")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_wrapper_restore",
		ZifStreamWrapperRestore,
		ArginfoStreamWrapperRestore,
		uint32_t(b.SizeOf("arginfo_stream_wrapper_restore")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_get_wrappers",
		ZifStreamGetWrappers,
		ArginfoStreamGetWrappers,
		uint32_t(b.SizeOf("arginfo_stream_get_wrappers")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_get_transports",
		ZifStreamGetTransports,
		ArginfoStreamGetTransports,
		uint32_t(b.SizeOf("arginfo_stream_get_transports")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_resolve_include_path",
		ZifStreamResolveIncludePath,
		ArginfoStreamResolveIncludePath,
		uint32_t(b.SizeOf("arginfo_stream_resolve_include_path")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_is_local",
		ZifStreamIsLocal,
		ArginfoStreamIsLocal,
		uint32_t(b.SizeOf("arginfo_stream_is_local")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_headers",
		ZifGetHeaders,
		ArginfoGetHeaders,
		uint32_t(b.SizeOf("arginfo_get_headers")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_set_timeout",
		ZifStreamSetTimeout,
		ArginfoStreamSetTimeout,
		uint32_t(b.SizeOf("arginfo_stream_set_timeout")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"socket_set_timeout",
		ZifStreamSetTimeout,
		ArginfoStreamSetTimeout,
		uint32_t(b.SizeOf("arginfo_stream_set_timeout")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"socket_get_status",
		ZifStreamGetMetaData,
		ArginfoStreamGetMetaData,
		uint32_t(b.SizeOf("arginfo_stream_get_meta_data")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"realpath",
		ZifRealpath,
		ArginfoRealpath,
		uint32_t(b.SizeOf("arginfo_realpath")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fnmatch",
		ZifFnmatch,
		ArginfoFnmatch,
		uint32_t(b.SizeOf("arginfo_fnmatch")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fsockopen",
		ZifFsockopen,
		ArginfoFsockopen,
		uint32_t(b.SizeOf("arginfo_fsockopen")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pfsockopen",
		ZifPfsockopen,
		ArginfoPfsockopen,
		uint32_t(b.SizeOf("arginfo_pfsockopen")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pack",
		ZifPack,
		ArginfoPack,
		uint32_t(b.SizeOf("arginfo_pack")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"unpack",
		ZifUnpack,
		ArginfoUnpack,
		uint32_t(b.SizeOf("arginfo_unpack")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_browser",
		ZifGetBrowser,
		ArginfoGetBrowser,
		uint32_t(b.SizeOf("arginfo_get_browser")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"crypt",
		ZifCrypt,
		ArginfoCrypt,
		uint32_t(b.SizeOf("arginfo_crypt")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"opendir",
		ZifOpendir,
		ArginfoOpendir,
		uint32_t(b.SizeOf("arginfo_opendir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"closedir",
		ZifClosedir,
		ArginfoClosedir,
		uint32_t(b.SizeOf("arginfo_closedir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chdir",
		ZifChdir,
		ArginfoChdir,
		uint32_t(b.SizeOf("arginfo_chdir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chroot",
		ZifChroot,
		ArginfoChroot,
		uint32_t(b.SizeOf("arginfo_chroot")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getcwd",
		ZifGetcwd,
		ArginfoGetcwd,
		uint32_t(b.SizeOf("arginfo_getcwd")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rewinddir",
		ZifRewinddir,
		ArginfoRewinddir,
		uint32_t(b.SizeOf("arginfo_rewinddir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"readdir",
		PhpIfReaddir,
		ArginfoReaddir,
		uint32_t(b.SizeOf("arginfo_readdir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"dir",
		ZifGetdir,
		ArginfoDir,
		uint32_t(b.SizeOf("arginfo_dir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"scandir",
		ZifScandir,
		ArginfoScandir,
		uint32_t(b.SizeOf("arginfo_scandir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"glob",
		ZifGlob,
		ArginfoGlob,
		uint32_t(b.SizeOf("arginfo_glob")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fileatime",
		ZifFileatime,
		ArginfoFileatime,
		uint32_t(b.SizeOf("arginfo_fileatime")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"filectime",
		ZifFilectime,
		ArginfoFilectime,
		uint32_t(b.SizeOf("arginfo_filectime")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"filegroup",
		ZifFilegroup,
		ArginfoFilegroup,
		uint32_t(b.SizeOf("arginfo_filegroup")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fileinode",
		ZifFileinode,
		ArginfoFileinode,
		uint32_t(b.SizeOf("arginfo_fileinode")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"filemtime",
		ZifFilemtime,
		ArginfoFilemtime,
		uint32_t(b.SizeOf("arginfo_filemtime")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fileowner",
		ZifFileowner,
		ArginfoFileowner,
		uint32_t(b.SizeOf("arginfo_fileowner")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fileperms",
		ZifFileperms,
		ArginfoFileperms,
		uint32_t(b.SizeOf("arginfo_fileperms")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"filesize",
		ZifFilesize,
		ArginfoFilesize,
		uint32_t(b.SizeOf("arginfo_filesize")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"filetype",
		ZifFiletype,
		ArginfoFiletype,
		uint32_t(b.SizeOf("arginfo_filetype")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"file_exists",
		ZifFileExists,
		ArginfoFileExists,
		uint32_t(b.SizeOf("arginfo_file_exists")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_writable",
		ZifIsWritable,
		ArginfoIsWritable,
		uint32_t(b.SizeOf("arginfo_is_writable")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_writeable",
		ZifIsWritable,
		ArginfoIsWritable,
		uint32_t(b.SizeOf("arginfo_is_writable")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_readable",
		ZifIsReadable,
		ArginfoIsReadable,
		uint32_t(b.SizeOf("arginfo_is_readable")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_executable",
		ZifIsExecutable,
		ArginfoIsExecutable,
		uint32_t(b.SizeOf("arginfo_is_executable")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_file",
		ZifIsFile,
		ArginfoIsFile,
		uint32_t(b.SizeOf("arginfo_is_file")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_dir",
		ZifIsDir,
		ArginfoIsDir,
		uint32_t(b.SizeOf("arginfo_is_dir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_link",
		ZifIsLink,
		ArginfoIsLink,
		uint32_t(b.SizeOf("arginfo_is_link")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stat",
		PhpIfStat,
		ArginfoStat,
		uint32_t(b.SizeOf("arginfo_stat")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"lstat",
		PhpIfLstat,
		ArginfoLstat,
		uint32_t(b.SizeOf("arginfo_lstat")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chown",
		ZifChown,
		ArginfoChown,
		uint32_t(b.SizeOf("arginfo_chown")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chgrp",
		ZifChgrp,
		ArginfoChgrp,
		uint32_t(b.SizeOf("arginfo_chgrp")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"lchown",
		ZifLchown,
		ArginfoLchown,
		uint32_t(b.SizeOf("arginfo_lchown")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"lchgrp",
		ZifLchgrp,
		ArginfoLchgrp,
		uint32_t(b.SizeOf("arginfo_lchgrp")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chmod",
		ZifChmod,
		ArginfoChmod,
		uint32_t(b.SizeOf("arginfo_chmod")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"touch",
		ZifTouch,
		ArginfoTouch,
		uint32_t(b.SizeOf("arginfo_touch")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"clearstatcache",
		ZifClearstatcache,
		ArginfoClearstatcache,
		uint32_t(b.SizeOf("arginfo_clearstatcache")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"disk_total_space",
		ZifDiskTotalSpace,
		ArginfoDiskTotalSpace,
		uint32_t(b.SizeOf("arginfo_disk_total_space")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"disk_free_space",
		ZifDiskFreeSpace,
		ArginfoDiskFreeSpace,
		uint32_t(b.SizeOf("arginfo_disk_free_space")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"diskfreespace",
		ZifDiskFreeSpace,
		ArginfoDiskFreeSpace,
		uint32_t(b.SizeOf("arginfo_disk_free_space")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"realpath_cache_size",
		ZifRealpathCacheSize,
		ArginfoRealpathCacheSize,
		uint32_t(b.SizeOf("arginfo_realpath_cache_size")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"realpath_cache_get",
		ZifRealpathCacheGet,
		ArginfoRealpathCacheGet,
		uint32_t(b.SizeOf("arginfo_realpath_cache_get")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"mail",
		ZifMail,
		ArginfoMail,
		uint32_t(b.SizeOf("arginfo_mail")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ezmlm_hash",
		ZifEzmlmHash,
		ArginfoEzmlmHash,
		uint32_t(b.SizeOf("arginfo_ezmlm_hash")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_DEPRECATED,
	},
	{
		"openlog",
		ZifOpenlog,
		ArginfoOpenlog,
		uint32_t(b.SizeOf("arginfo_openlog")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"syslog",
		ZifSyslog,
		ArginfoSyslog,
		uint32_t(b.SizeOf("arginfo_syslog")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"closelog",
		ZifCloselog,
		ArginfoCloselog,
		uint32_t(b.SizeOf("arginfo_closelog")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"lcg_value",
		ZifLcgValue,
		ArginfoLcgValue,
		uint32_t(b.SizeOf("arginfo_lcg_value")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"metaphone",
		ZifMetaphone,
		ArginfoMetaphone,
		uint32_t(b.SizeOf("arginfo_metaphone")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_start",
		core.ZifObStart,
		ArginfoObStart,
		uint32_t(b.SizeOf("arginfo_ob_start")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_flush",
		core.ZifObFlush,
		ArginfoObFlush,
		uint32_t(b.SizeOf("arginfo_ob_flush")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_clean",
		core.ZifObClean,
		ArginfoObClean,
		uint32_t(b.SizeOf("arginfo_ob_clean")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_end_flush",
		core.ZifObEndFlush,
		ArginfoObEndFlush,
		uint32_t(b.SizeOf("arginfo_ob_end_flush")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_end_clean",
		core.ZifObEndClean,
		ArginfoObEndClean,
		uint32_t(b.SizeOf("arginfo_ob_end_clean")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_get_flush",
		core.ZifObGetFlush,
		ArginfoObGetFlush,
		uint32_t(b.SizeOf("arginfo_ob_get_flush")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_get_clean",
		core.ZifObGetClean,
		ArginfoObGetClean,
		uint32_t(b.SizeOf("arginfo_ob_get_clean")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_get_length",
		core.ZifObGetLength,
		ArginfoObGetLength,
		uint32_t(b.SizeOf("arginfo_ob_get_length")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_get_level",
		core.ZifObGetLevel,
		ArginfoObGetLevel,
		uint32_t(b.SizeOf("arginfo_ob_get_level")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_get_status",
		core.ZifObGetStatus,
		ArginfoObGetStatus,
		uint32_t(b.SizeOf("arginfo_ob_get_status")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_get_contents",
		core.ZifObGetContents,
		ArginfoObGetContents,
		uint32_t(b.SizeOf("arginfo_ob_get_contents")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_implicit_flush",
		core.ZifObImplicitFlush,
		ArginfoObImplicitFlush,
		uint32_t(b.SizeOf("arginfo_ob_implicit_flush")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_list_handlers",
		core.ZifObListHandlers,
		ArginfoObListHandlers,
		uint32_t(b.SizeOf("arginfo_ob_list_handlers")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ksort",
		ZifKsort,
		ArginfoKsort,
		uint32_t(b.SizeOf("arginfo_ksort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"krsort",
		ZifKrsort,
		ArginfoKrsort,
		uint32_t(b.SizeOf("arginfo_krsort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"natsort",
		ZifNatsort,
		ArginfoNatsort,
		uint32_t(b.SizeOf("arginfo_natsort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"natcasesort",
		ZifNatcasesort,
		ArginfoNatcasesort,
		uint32_t(b.SizeOf("arginfo_natcasesort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"asort",
		ZifAsort,
		ArginfoAsort,
		uint32_t(b.SizeOf("arginfo_asort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"arsort",
		ZifArsort,
		ArginfoArsort,
		uint32_t(b.SizeOf("arginfo_arsort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sort",
		ZifSort,
		ArginfoSort,
		uint32_t(b.SizeOf("arginfo_sort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rsort",
		ZifRsort,
		ArginfoRsort,
		uint32_t(b.SizeOf("arginfo_rsort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"usort",
		ZifUsort,
		ArginfoUsort,
		uint32_t(b.SizeOf("arginfo_usort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"uasort",
		ZifUasort,
		ArginfoUasort,
		uint32_t(b.SizeOf("arginfo_uasort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"uksort",
		ZifUksort,
		ArginfoUksort,
		uint32_t(b.SizeOf("arginfo_uksort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"shuffle",
		ZifShuffle,
		ArginfoShuffle,
		uint32_t(b.SizeOf("arginfo_shuffle")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_walk",
		ZifArrayWalk,
		ArginfoArrayWalk,
		uint32_t(b.SizeOf("arginfo_array_walk")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_walk_recursive",
		ZifArrayWalkRecursive,
		ArginfoArrayWalkRecursive,
		uint32_t(b.SizeOf("arginfo_array_walk_recursive")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"count",
		ZifCount,
		ArginfoCount,
		uint32_t(b.SizeOf("arginfo_count")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"end",
		ZifEnd,
		ArginfoEnd,
		uint32_t(b.SizeOf("arginfo_end")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"prev",
		ZifPrev,
		ArginfoPrev,
		uint32_t(b.SizeOf("arginfo_prev")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"next",
		ZifNext,
		ArginfoNext,
		uint32_t(b.SizeOf("arginfo_next")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"reset",
		ZifReset,
		ArginfoReset,
		uint32_t(b.SizeOf("arginfo_reset")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"current",
		ZifCurrent,
		ArginfoCurrent,
		uint32_t(b.SizeOf("arginfo_current")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"key",
		ZifKey,
		ArginfoKey,
		uint32_t(b.SizeOf("arginfo_key")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"min",
		ZifMin,
		ArginfoMin,
		uint32_t(b.SizeOf("arginfo_min")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"max",
		ZifMax,
		ArginfoMax,
		uint32_t(b.SizeOf("arginfo_max")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"in_array",
		ZifInArray,
		ArginfoInArray,
		uint32_t(b.SizeOf("arginfo_in_array")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_search",
		ZifArraySearch,
		ArginfoArraySearch,
		uint32_t(b.SizeOf("arginfo_array_search")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"extract",
		ZifExtract,
		ArginfoExtract,
		uint32_t(b.SizeOf("arginfo_extract")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"compact",
		ZifCompact,
		ArginfoCompact,
		uint32_t(b.SizeOf("arginfo_compact")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_fill",
		ZifArrayFill,
		ArginfoArrayFill,
		uint32_t(b.SizeOf("arginfo_array_fill")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_fill_keys",
		ZifArrayFillKeys,
		ArginfoArrayFillKeys,
		uint32_t(b.SizeOf("arginfo_array_fill_keys")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"range",
		ZifRange,
		ArginfoRange,
		uint32_t(b.SizeOf("arginfo_range")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_multisort",
		ZifArrayMultisort,
		ArginfoArrayMultisort,
		uint32_t(b.SizeOf("arginfo_array_multisort")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_push",
		ZifArrayPush,
		ArginfoArrayPush,
		uint32_t(b.SizeOf("arginfo_array_push")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_pop",
		ZifArrayPop,
		ArginfoArrayPop,
		uint32_t(b.SizeOf("arginfo_array_pop")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_shift",
		ZifArrayShift,
		ArginfoArrayShift,
		uint32_t(b.SizeOf("arginfo_array_shift")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_unshift",
		ZifArrayUnshift,
		ArginfoArrayUnshift,
		uint32_t(b.SizeOf("arginfo_array_unshift")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_splice",
		ZifArraySplice,
		ArginfoArraySplice,
		uint32_t(b.SizeOf("arginfo_array_splice")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_slice",
		ZifArraySlice,
		ArginfoArraySlice,
		uint32_t(b.SizeOf("arginfo_array_slice")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_merge",
		ZifArrayMerge,
		ArginfoArrayMerge,
		uint32_t(b.SizeOf("arginfo_array_merge")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_merge_recursive",
		ZifArrayMergeRecursive,
		ArginfoArrayMergeRecursive,
		uint32_t(b.SizeOf("arginfo_array_merge_recursive")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_replace",
		ZifArrayReplace,
		ArginfoArrayReplace,
		uint32_t(b.SizeOf("arginfo_array_replace")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_replace_recursive",
		ZifArrayReplaceRecursive,
		ArginfoArrayReplaceRecursive,
		uint32_t(b.SizeOf("arginfo_array_replace_recursive")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_keys",
		ZifArrayKeys,
		ArginfoArrayKeys,
		uint32_t(b.SizeOf("arginfo_array_keys")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_key_first",
		ZifArrayKeyFirst,
		ArginfoArrayKeyFirst,
		uint32_t(b.SizeOf("arginfo_array_key_first")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_key_last",
		ZifArrayKeyLast,
		ArginfoArrayKeyLast,
		uint32_t(b.SizeOf("arginfo_array_key_last")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_values",
		ZifArrayValues,
		ArginfoArrayValues,
		uint32_t(b.SizeOf("arginfo_array_values")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_count_values",
		ZifArrayCountValues,
		ArginfoArrayCountValues,
		uint32_t(b.SizeOf("arginfo_array_count_values")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_column",
		ZifArrayColumn,
		ArginfoArrayColumn,
		uint32_t(b.SizeOf("arginfo_array_column")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_reverse",
		ZifArrayReverse,
		ArginfoArrayReverse,
		uint32_t(b.SizeOf("arginfo_array_reverse")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_reduce",
		ZifArrayReduce,
		ArginfoArrayReduce,
		uint32_t(b.SizeOf("arginfo_array_reduce")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_pad",
		ZifArrayPad,
		ArginfoArrayPad,
		uint32_t(b.SizeOf("arginfo_array_pad")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_flip",
		ZifArrayFlip,
		ArginfoArrayFlip,
		uint32_t(b.SizeOf("arginfo_array_flip")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_change_key_case",
		ZifArrayChangeKeyCase,
		ArginfoArrayChangeKeyCase,
		uint32_t(b.SizeOf("arginfo_array_change_key_case")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_rand",
		ZifArrayRand,
		ArginfoArrayRand,
		uint32_t(b.SizeOf("arginfo_array_rand")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_unique",
		ZifArrayUnique,
		ArginfoArrayUnique,
		uint32_t(b.SizeOf("arginfo_array_unique")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_intersect",
		ZifArrayIntersect,
		ArginfoArrayIntersect,
		uint32_t(b.SizeOf("arginfo_array_intersect")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_intersect_key",
		ZifArrayIntersectKey,
		ArginfoArrayIntersectKey,
		uint32_t(b.SizeOf("arginfo_array_intersect_key")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_intersect_ukey",
		ZifArrayIntersectUkey,
		ArginfoArrayIntersectUkey,
		uint32_t(b.SizeOf("arginfo_array_intersect_ukey")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_uintersect",
		ZifArrayUintersect,
		ArginfoArrayUintersect,
		uint32_t(b.SizeOf("arginfo_array_uintersect")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_intersect_assoc",
		ZifArrayIntersectAssoc,
		ArginfoArrayIntersectAssoc,
		uint32_t(b.SizeOf("arginfo_array_intersect_assoc")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_uintersect_assoc",
		ZifArrayUintersectAssoc,
		ArginfoArrayUintersectAssoc,
		uint32_t(b.SizeOf("arginfo_array_uintersect_assoc")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_intersect_uassoc",
		ZifArrayIntersectUassoc,
		ArginfoArrayIntersectUassoc,
		uint32_t(b.SizeOf("arginfo_array_intersect_uassoc")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_uintersect_uassoc",
		ZifArrayUintersectUassoc,
		ArginfoArrayUintersectUassoc,
		uint32_t(b.SizeOf("arginfo_array_uintersect_uassoc")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_diff",
		ZifArrayDiff,
		ArginfoArrayDiff,
		uint32_t(b.SizeOf("arginfo_array_diff")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_diff_key",
		ZifArrayDiffKey,
		ArginfoArrayDiffKey,
		uint32_t(b.SizeOf("arginfo_array_diff_key")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_diff_ukey",
		ZifArrayDiffUkey,
		ArginfoArrayDiffUkey,
		uint32_t(b.SizeOf("arginfo_array_diff_ukey")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_udiff",
		ZifArrayUdiff,
		ArginfoArrayUdiff,
		uint32_t(b.SizeOf("arginfo_array_udiff")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_diff_assoc",
		ZifArrayDiffAssoc,
		ArginfoArrayDiffAssoc,
		uint32_t(b.SizeOf("arginfo_array_diff_assoc")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_udiff_assoc",
		ZifArrayUdiffAssoc,
		ArginfoArrayUdiffAssoc,
		uint32_t(b.SizeOf("arginfo_array_udiff_assoc")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_diff_uassoc",
		ZifArrayDiffUassoc,
		ArginfoArrayDiffUassoc,
		uint32_t(b.SizeOf("arginfo_array_diff_uassoc")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_udiff_uassoc",
		ZifArrayUdiffUassoc,
		ArginfoArrayUdiffUassoc,
		uint32_t(b.SizeOf("arginfo_array_udiff_uassoc")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_sum",
		ZifArraySum,
		ArginfoArraySum,
		uint32_t(b.SizeOf("arginfo_array_sum")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_product",
		ZifArrayProduct,
		ArginfoArrayProduct,
		uint32_t(b.SizeOf("arginfo_array_product")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_filter",
		ZifArrayFilter,
		ArginfoArrayFilter,
		uint32_t(b.SizeOf("arginfo_array_filter")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_map",
		ZifArrayMap,
		ArginfoArrayMap,
		uint32_t(b.SizeOf("arginfo_array_map")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_chunk",
		ZifArrayChunk,
		ArginfoArrayChunk,
		uint32_t(b.SizeOf("arginfo_array_chunk")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_combine",
		ZifArrayCombine,
		ArginfoArrayCombine,
		uint32_t(b.SizeOf("arginfo_array_combine")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_key_exists",
		ZifArrayKeyExists,
		ArginfoArrayKeyExists,
		uint32_t(b.SizeOf("arginfo_array_key_exists")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pos",
		ZifCurrent,
		ArginfoCurrent,
		uint32_t(b.SizeOf("arginfo_current")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sizeof",
		ZifCount,
		ArginfoCount,
		uint32_t(b.SizeOf("arginfo_count")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"key_exists",
		ZifArrayKeyExists,
		ArginfoArrayKeyExists,
		uint32_t(b.SizeOf("arginfo_array_key_exists")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"assert",
		ZifAssert,
		ArginfoAssert,
		uint32_t(b.SizeOf("arginfo_assert")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"assert_options",
		ZifAssertOptions,
		ArginfoAssertOptions,
		uint32_t(b.SizeOf("arginfo_assert_options")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"version_compare",
		ZifVersionCompare,
		ArginfoVersionCompare,
		uint32_t(b.SizeOf("arginfo_version_compare")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ftok",
		ZifFtok,
		ArginfoFtok,
		uint32_t(b.SizeOf("arginfo_ftok")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_rot13",
		ZifStrRot13,
		ArginfoStrRot13,
		uint32_t(b.SizeOf("arginfo_str_rot13")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_get_filters",
		ZifStreamGetFilters,
		ArginfoStreamGetFilters,
		uint32_t(b.SizeOf("arginfo_stream_get_filters")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_filter_register",
		ZifStreamFilterRegister,
		ArginfoStreamFilterRegister,
		uint32_t(b.SizeOf("arginfo_stream_filter_register")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_bucket_make_writeable",
		ZifStreamBucketMakeWriteable,
		ArginfoStreamBucketMakeWriteable,
		uint32_t(b.SizeOf("arginfo_stream_bucket_make_writeable")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_bucket_prepend",
		ZifStreamBucketPrepend,
		ArginfoStreamBucketPrepend,
		uint32_t(b.SizeOf("arginfo_stream_bucket_prepend")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_bucket_append",
		ZifStreamBucketAppend,
		ArginfoStreamBucketAppend,
		uint32_t(b.SizeOf("arginfo_stream_bucket_append")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_bucket_new",
		ZifStreamBucketNew,
		ArginfoStreamBucketNew,
		uint32_t(b.SizeOf("arginfo_stream_bucket_new")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"output_add_rewrite_var",
		core.ZifOutputAddRewriteVar,
		ArginfoOutputAddRewriteVar,
		uint32_t(b.SizeOf("arginfo_output_add_rewrite_var")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"output_reset_rewrite_vars",
		core.ZifOutputResetRewriteVars,
		ArginfoOutputResetRewriteVars,
		uint32_t(b.SizeOf("arginfo_output_reset_rewrite_vars")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sys_get_temp_dir",
		ZifSysGetTempDir,
		ArginfoSysGetTempDir,
		uint32_t(b.SizeOf("arginfo_sys_get_temp_dir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

var StandardDeps []zend.ZendModuleDep = []zend.ZendModuleDep{
	{"session", nil, nil, zend.MODULE_DEP_OPTIONAL},
	{nil, nil, nil, 0},
}

/* }}} */

var BasicFunctionsModule zend.ZendModuleEntry = zend.ZendModuleEntry{
	b.SizeOf("zend_module_entry"),
	zend.ZEND_MODULE_API_NO,
	core.ZEND_DEBUG,
	zend.USING_ZTS,
	nil,
	StandardDeps,
	"standard",
	BasicFunctions,
	ZmStartupBasic,
	ZmShutdownBasic,
	ZmActivateBasic,
	ZmDeactivateBasic,
	ZmInfoBasic,
	PHP_STANDARD_VERSION,
	0,
	nil,
	nil,
	nil,
	nil,
	0,
	0,
	nil,
	0,
	"API" + "ZEND_MODULE_API_NO" + zend.ZEND_BUILD_TS,
}

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ proto string inet_ntop(string in_addr)
   Converts a packed inet address to a human readable IP address string */

/* }}} */

/* {{{ proto string inet_pton(string ip_address)
   Converts a human readable IP address to a packed binary string */

/* }}} */

/* {{{ proto int ip2long(string ip_address)
   Converts a string containing an (IPv4) Internet Protocol dotted address into a proper address */

/* }}} */

/* }}} */

/* }}} */

/* {{{ proto bool putenv(string setting)
   Set the value of an environment variable */

/* }}} */

/* {{{ free_argv()
   Free the memory allocated to an argv array. */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ proto mixed time_nanosleep(int seconds, int nanoseconds)
   Delay for a number of seconds and nano seconds */

/* }}} */

/* }}} */

/* {{{ proto string get_current_user(void)
   Get the name of the owner of the current PHP script */

/* }}} */

/* {{{ add_config_entry
 */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ proto int getservbyname(string service, string protocol)
   Returns port associated with service. Protocol must be "tcp" or "udp" */

/* }}} */

/* {{{ proto string getservbyport(int port, string protocol)
   Returns service name associated with port. Protocol must be "tcp" or "udp" */

/* }}} */

/* {{{ proto int getprotobyname(string name)
   Returns protocol number associated with name as per /etc/protocols */

/* }}} */

/* {{{ proto string getprotobynumber(int proto)
   Returns protocol name associated with protocol number proto */

/* }}} */

/* {{{ proto bool register_tick_function(string function_name [, mixed arg [, mixed ... ]])
   Registers a tick callback function */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ proto array sys_getloadavg()
 */

/* }}} */
