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
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("seconds", 0, 0, 0),
}
var ArginfoHeaderRegisterCallback []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("callback", 0, 0, 0),
}
var ArginfoObStart []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("user_function", 0, 0, 0),
	zend.MakeZendInternalArgInfo("chunk_size", 0, 0, 0),
	zend.MakeZendInternalArgInfo("flags", 0, 0, 0),
}
var ArginfoObFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoObClean []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoObEndFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoObEndClean []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoObGetFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoObGetClean []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoObGetContents []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoObGetLevel []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoObGetLength []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoObListHandlers []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoObGetStatus []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("full_status", 0, 0, 0),
}
var ArginfoObImplicitFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("flag", 0, 0, 0),
}
var ArginfoOutputResetRewriteVars []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoOutputAddRewriteVar []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("value", 0, 0, 0),
}
var ArginfoStreamWrapperRegister []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("protocol", 0, 0, 0),
	zend.MakeZendInternalArgInfo("classname", 0, 0, 0),
	zend.MakeZendInternalArgInfo("flags", 0, 0, 0),
}
var ArginfoStreamWrapperUnregister []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("protocol", 0, 0, 0),
}
var ArginfoStreamWrapperRestore []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("protocol", 0, 0, 0),
}
var ArginfoKrsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
	zend.MakeZendInternalArgInfo("sort_flags", 0, 0, 0),
}
var ArginfoKsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
	zend.MakeZendInternalArgInfo("sort_flags", 0, 0, 0),
}
var ArginfoCount []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
	zend.MakeZendInternalArgInfo("mode", 0, 0, 0),
}
var ArginfoNatsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
}
var ArginfoNatcasesort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
}
var ArginfoAsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
	zend.MakeZendInternalArgInfo("sort_flags", 0, 0, 0),
}
var ArginfoArsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
	zend.MakeZendInternalArgInfo("sort_flags", 0, 0, 0),
}
var ArginfoSort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
	zend.MakeZendInternalArgInfo("sort_flags", 0, 0, 0),
}
var ArginfoRsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
	zend.MakeZendInternalArgInfo("sort_flags", 0, 0, 0),
}
var ArginfoUsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
	zend.MakeZendInternalArgInfo("cmp_function", 0, 0, 0),
}
var ArginfoUasort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
	zend.MakeZendInternalArgInfo("cmp_function", 0, 0, 0),
}
var ArginfoUksort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
	zend.MakeZendInternalArgInfo("cmp_function", 0, 0, 0),
}
var ArginfoEnd []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
}
var ArginfoPrev []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
}
var ArginfoNext []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
}
var ArginfoReset []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
}
var ArginfoCurrent []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
}
var ArginfoKey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
}
var ArginfoMin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("args", 0, 0, 1),
}
var ArginfoMax []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("args", 0, 0, 1),
}
var ArginfoArrayWalk []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("input", 0, 1, 0),
	zend.MakeZendInternalArgInfo("funcname", 0, 0, 0),
	zend.MakeZendInternalArgInfo("userdata", 0, 0, 0),
}
var ArginfoArrayWalkRecursive []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("input", 0, 1, 0),
	zend.MakeZendInternalArgInfo("funcname", 0, 0, 0),
	zend.MakeZendInternalArgInfo("userdata", 0, 0, 0),
}
var ArginfoInArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("needle", 0, 0, 0),
	zend.MakeZendInternalArgInfo("haystack", 0, 0, 0),
	zend.MakeZendInternalArgInfo("strict", 0, 0, 0),
}
var ArginfoArraySearch []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("needle", 0, 0, 0),
	zend.MakeZendInternalArgInfo("haystack", 0, 0, 0),
	zend.MakeZendInternalArgInfo("strict", 0, 0, 0),
}
var ArginfoExtract []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, zend.ZEND_SEND_PREFER_REF, 0),
	zend.MakeZendInternalArgInfo("extract_type", 0, 0, 0),
	zend.MakeZendInternalArgInfo("prefix", 0, 0, 0),
}
var ArginfoCompact []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("var_names", 0, 0, 1),
}
var ArginfoArrayFill []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("start_key", 0, 0, 0),
	zend.MakeZendInternalArgInfo("num", 0, 0, 0),
	zend.MakeZendInternalArgInfo("val", 0, 0, 0),
}
var ArginfoArrayFillKeys []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("keys", 0, 0, 0),
	zend.MakeZendInternalArgInfo("val", 0, 0, 0),
}
var ArginfoRange []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("low", 0, 0, 0),
	zend.MakeZendInternalArgInfo("high", 0, 0, 0),
	zend.MakeZendInternalArgInfo("step", 0, 0, 0),
}
var ArginfoShuffle []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
}
var ArginfoArrayPush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stack", 0, 1, 0),
	zend.MakeZendInternalArgInfo("vars", 0, 0, 1),
}
var ArginfoArrayPop []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("stack", 0, 1, 0),
}
var ArginfoArrayShift []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("stack", 0, 1, 0),
}
var ArginfoArrayUnshift []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stack", 0, 1, 0),
	zend.MakeZendInternalArgInfo("vars", 0, 0, 1),
}
var ArginfoArraySplice []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 1, 0),
	zend.MakeZendInternalArgInfo("offset", 0, 0, 0),
	zend.MakeZendInternalArgInfo("length", 0, 0, 0),
	zend.MakeZendInternalArgInfo("replacement", 0, 0, 0),
}
var ArginfoArraySlice []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
	zend.MakeZendInternalArgInfo("offset", 0, 0, 0),
	zend.MakeZendInternalArgInfo("length", 0, 0, 0),
	zend.MakeZendInternalArgInfo("preserve_keys", 0, 0, 0),
}
var ArginfoArrayMerge []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arrays", 0, 0, 1),
}
var ArginfoArrayMergeRecursive []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arrays", 0, 0, 1),
}
var ArginfoArrayReplace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arrays", 0, 0, 1),
}
var ArginfoArrayReplaceRecursive []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arrays", 0, 0, 1),
}
var ArginfoArrayKeys []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
	zend.MakeZendInternalArgInfo("search_value", 0, 0, 0),
	zend.MakeZendInternalArgInfo("strict", 0, 0, 0),
}
var ArginfoArrayKeyFirst []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
}
var ArginfoArrayKeyLast []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
}
var ArginfoArrayValues []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
}
var ArginfoArrayCountValues []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
}
var ArginfoArrayColumn []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
	zend.MakeZendInternalArgInfo("column_key", 0, 0, 0),
	zend.MakeZendInternalArgInfo("index_key", 0, 0, 0),
}
var ArginfoArrayReverse []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("input", 0, 0, 0),
	zend.MakeZendInternalArgInfo("preserve_keys", 0, 0, 0),
}
var ArginfoArrayPad []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
	zend.MakeZendInternalArgInfo("pad_size", 0, 0, 0),
	zend.MakeZendInternalArgInfo("pad_value", 0, 0, 0),
}
var ArginfoArrayFlip []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
}
var ArginfoArrayChangeKeyCase []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("input", 0, 0, 0),
	zend.MakeZendInternalArgInfo("case", 0, 0, 0),
}
var ArginfoArrayUnique []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
	zend.MakeZendInternalArgInfo("flags", 0, 0, 0),
}
var ArginfoArrayIntersectKey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arrays", 0, 0, 1),
}
var ArginfoArrayIntersectUkey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr2", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback_key_compare_func", 0, 0, 0),
}
var ArginfoArrayIntersect []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arrays", 0, 0, 1),
}
var ArginfoArrayUintersect []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr2", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback_data_compare_func", 0, 0, 0),
}
var ArginfoArrayIntersectAssoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arrays", 0, 0, 1),
}
var ArginfoArrayUintersectAssoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr2", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback_data_compare_func", 0, 0, 0),
}
var ArginfoArrayIntersectUassoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr2", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback_key_compare_func", 0, 0, 0),
}
var ArginfoArrayUintersectUassoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr2", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback_data_compare_func", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback_key_compare_func", 0, 0, 0),
}
var ArginfoArrayDiffKey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arrays", 0, 0, 1),
}
var ArginfoArrayDiffUkey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr2", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback_key_comp_func", 0, 0, 0),
}
var ArginfoArrayDiff []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arrays", 0, 0, 1),
}
var ArginfoArrayUdiff []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr2", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback_data_comp_func", 0, 0, 0),
}
var ArginfoArrayDiffAssoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arrays", 0, 0, 1),
}
var ArginfoArrayDiffUassoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr2", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback_data_comp_func", 0, 0, 0),
}
var ArginfoArrayUdiffAssoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr2", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback_key_comp_func", 0, 0, 0),
}
var ArginfoArrayUdiffUassoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr2", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback_data_comp_func", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback_key_comp_func", 0, 0, 0),
}
var ArginfoArrayMultisort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arr1", 0, zend.ZEND_SEND_PREFER_REF, 0),
	zend.MakeZendInternalArgInfo("sort_order", 0, zend.ZEND_SEND_PREFER_REF, 0),
	zend.MakeZendInternalArgInfo("sort_flags", 0, zend.ZEND_SEND_PREFER_REF, 0),
	zend.MakeZendInternalArgInfo("arr2", 0, zend.ZEND_SEND_PREFER_REF, 1),
}
var ArginfoArrayRand []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
	zend.MakeZendInternalArgInfo("num_req", 0, 0, 0),
}
var ArginfoArraySum []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
}
var ArginfoArrayProduct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
}
var ArginfoArrayReduce []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback", 0, 0, 0),
	zend.MakeZendInternalArgInfo("initial", 0, 0, 0),
}
var ArginfoArrayFilter []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback", 0, 0, 0),
	zend.MakeZendInternalArgInfo("use_keys", 0, 0, 0),
}
var ArginfoArrayMap []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("callback", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arrays", 0, 0, 1),
}
var ArginfoArrayKeyExists []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("key", 0, 0, 0),
	zend.MakeZendInternalArgInfo("search", 0, 0, 0),
}
var ArginfoArrayChunk []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
	zend.MakeZendInternalArgInfo("size", 0, 0, 0),
	zend.MakeZendInternalArgInfo("preserve_keys", 0, 0, 0),
}
var ArginfoArrayCombine []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("keys", 0, 0, 0),
	zend.MakeZendInternalArgInfo("values", 0, 0, 0),
}
var ArginfoGetMagicQuotesGpc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoGetMagicQuotesRuntime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoConstant []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("const_name", 0, 0, 0),
}
var ArginfoInetNtop []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("in_addr", 0, 0, 0),
}
var ArginfoInetPton []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("ip_address", 0, 0, 0),
}
var ArginfoIp2long []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("ip_address", 0, 0, 0),
}
var ArginfoLong2ip []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("proper_address", 0, 0, 0),
}
var ArginfoGetenv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("varname", 0, 0, 0),
	zend.MakeZendInternalArgInfo("local_only", 0, 0, 0),
}
var ArginfoPutenv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("setting", 0, 0, 0),
}
var ArginfoGetopt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("options", 0, 0, 0),
	zend.MakeZendInternalArgInfo("opts", 0, 0, 0),
	zend.MakeZendInternalArgInfo("optind", 0, 1, 0),
}
var ArginfoFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoSleep []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("seconds", 0, 0, 0),
}
var ArginfoUsleep []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("micro_seconds", 0, 0, 0),
}
var ArginfoTimeNanosleep []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("seconds", 0, 0, 0),
	zend.MakeZendInternalArgInfo("nanoseconds", 0, 0, 0),
}
var ArginfoTimeSleepUntil []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("timestamp", 0, 0, 0),
}
var ArginfoGetCurrentUser []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoGetCfgVar []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("option_name", 0, 0, 0),
}
var ArginfoErrorLog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("message", 0, 0, 0),
	zend.MakeZendInternalArgInfo("message_type", 0, 0, 0),
	zend.MakeZendInternalArgInfo("destination", 0, 0, 0),
	zend.MakeZendInternalArgInfo("extra_headers", 0, 0, 0),
}
var ArginfoErrorGetLast []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
}
var ArginfoErrorClearLast []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
}
var ArginfoCallUserFunc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("function_name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("parameters", 0, 0, 1),
}
var ArginfoCallUserFuncArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("function_name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("parameters", 0, 0, 0),
}
var ArginfoForwardStaticCall []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("function_name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("parameters", 0, 0, 1),
}
var ArginfoForwardStaticCallArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("function_name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("parameters", 0, 0, 0),
}
var ArginfoRegisterShutdownFunction []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("function_name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("parameters", 0, 0, 1),
}
var ArginfoHighlightFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("file_name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("return", 0, 0, 0),
}
var ArginfoPhpStripWhitespace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("file_name", 0, 0, 0),
}
var ArginfoHighlightString []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("string", 0, 0, 0),
	zend.MakeZendInternalArgInfo("return", 0, 0, 0),
}
var ArginfoIniGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("varname", 0, 0, 0),
}
var ArginfoIniGetAll []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("extension", 0, 0, 0),
	zend.MakeZendInternalArgInfo("details", 0, 0, 0),
}
var ArginfoIniSet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("varname", 0, 0, 0),
	zend.MakeZendInternalArgInfo("newvalue", 0, 0, 0),
}
var ArginfoIniRestore []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("varname", 0, 0, 0),
}
var ArginfoSetIncludePath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("new_include_path", 0, 0, 0),
}
var ArginfoGetIncludePath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoRestoreIncludePath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoPrintR []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
	zend.MakeZendInternalArgInfo("return", 0, 0, 0),
}
var ArginfoConnectionAborted []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoConnectionStatus []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoIgnoreUserAbort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("value", 0, 0, 0),
}
var ArginfoGetservbyname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("service", 0, 0, 0),
	zend.MakeZendInternalArgInfo("protocol", 0, 0, 0),
}
var ArginfoGetservbyport []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("port", 0, 0, 0),
	zend.MakeZendInternalArgInfo("protocol", 0, 0, 0),
}
var ArginfoGetprotobyname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("name", 0, 0, 0),
}
var ArginfoGetprotobynumber []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("proto", 0, 0, 0),
}
var ArginfoRegisterTickFunction []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("function_name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("parameters", 0, 0, 1),
}
var ArginfoUnregisterTickFunction []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("function_name", 0, 0, 0),
}
var ArginfoIsUploadedFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("path", 0, 0, 0),
}
var ArginfoMoveUploadedFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("path", 0, 0, 0),
	zend.MakeZendInternalArgInfo("new_path", 0, 0, 0),
}
var ArginfoParseIniFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("process_sections", 0, 0, 0),
	zend.MakeZendInternalArgInfo("scanner_mode", 0, 0, 0),
}
var ArginfoParseIniString []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("ini_string", 0, 0, 0),
	zend.MakeZendInternalArgInfo("process_sections", 0, 0, 0),
	zend.MakeZendInternalArgInfo("scanner_mode", 0, 0, 0),
}
var ArginfoSysGetloadavg []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoAssert []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("assertion", 0, 0, 0),
	zend.MakeZendInternalArgInfo("description", 0, 0, 0),
}
var ArginfoAssertOptions []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("what", 0, 0, 0),
	zend.MakeZendInternalArgInfo("value", 0, 0, 0),
}
var ArginfoBase64Encode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoBase64Decode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("strict", 0, 0, 0),
}
var ArginfoGetBrowser []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("browser_name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("return_array", 0, 0, 0),
}
var ArginfoCrc32 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoCrypt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("salt", 0, 0, 0),
}
var ArginfoConvertCyrString []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("from", 0, 0, 0),
	zend.MakeZendInternalArgInfo("to", 0, 0, 0),
}
var ArginfoStrptime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("timestamp", 0, 0, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
}
var ArginfoOpendir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("path", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoDir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("directory", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoClosedir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("dir_handle", 0, 0, 0),
}
var ArginfoChroot []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("directory", 0, 0, 0),
}
var ArginfoChdir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("directory", 0, 0, 0),
}
var ArginfoGetcwd []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoRewinddir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("dir_handle", 0, 0, 0),
}
var ArginfoReaddir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("dir_handle", 0, 0, 0),
}
var ArginfoGlob []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("pattern", 0, 0, 0),
	zend.MakeZendInternalArgInfo("flags", 0, 0, 0),
}
var ArginfoScandir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("dir", 0, 0, 0),
	zend.MakeZendInternalArgInfo("sorting_order", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoGethostbyaddr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("ip_address", 0, 0, 0),
}
var ArginfoGethostbyname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("hostname", 0, 0, 0),
}
var ArginfoGethostbynamel []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("hostname", 0, 0, 0),
}
var ArginfoGethostname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoNetGetInterfaces []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoDnsCheckRecord []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("host", 0, 0, 0),
	zend.MakeZendInternalArgInfo("type", 0, 0, 0),
}
var ArginfoDnsGetRecord []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("hostname", 0, 0, 0),
	zend.MakeZendInternalArgInfo("type", 0, 0, 0),
	zend.MakeZendInternalArgInfo("authns", zend.ZEND_TYPE_ENCODE(zend.IS_ARRAY, 1), 1, 0),
	zend.MakeZendInternalArgInfo("addtl", zend.ZEND_TYPE_ENCODE(zend.IS_ARRAY, 1), 1, 0),
	zend.MakeZendInternalArgInfo("raw", 0, 0, 0),
}
var ArginfoDnsGetMx []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("hostname", 0, 0, 0),
	zend.MakeZendInternalArgInfo("mxhosts", 0, 1, 0),
	zend.MakeZendInternalArgInfo("weight", 0, 1, 0),
}
var ArginfoExec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("command", 0, 0, 0),
	zend.MakeZendInternalArgInfo("output", 0, 1, 0),
	zend.MakeZendInternalArgInfo("return_value", 0, 1, 0),
}
var ArginfoSystem []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("command", 0, 0, 0),
	zend.MakeZendInternalArgInfo("return_value", 0, 1, 0),
}
var ArginfoPassthru []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("command", 0, 0, 0),
	zend.MakeZendInternalArgInfo("return_value", 0, 1, 0),
}
var ArginfoEscapeshellcmd []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("command", 0, 0, 0),
}
var ArginfoEscapeshellarg []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("arg", 0, 0, 0),
}
var ArginfoShellExec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("cmd", 0, 0, 0),
}
var ArginfoProcNice []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("priority", 0, 0, 0),
}
var ArginfoFlock []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
	zend.MakeZendInternalArgInfo("operation", 0, 0, 0),
	zend.MakeZendInternalArgInfo("wouldblock", 0, 1, 0),
}
var ArginfoGetMetaTags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("use_include_path", 0, 0, 0),
}
var ArginfoFileGetContents []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("flags", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
	zend.MakeZendInternalArgInfo("offset", 0, 0, 0),
	zend.MakeZendInternalArgInfo("maxlen", 0, 0, 0),
}
var ArginfoFilePutContents []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("data", 0, 0, 0),
	zend.MakeZendInternalArgInfo("flags", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("flags", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoTempnam []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("dir", 0, 0, 0),
	zend.MakeZendInternalArgInfo("prefix", 0, 0, 0),
}
var ArginfoTmpfile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoFopen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("mode", 0, 0, 0),
	zend.MakeZendInternalArgInfo("use_include_path", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoFclose []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
}
var ArginfoPopen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("command", 0, 0, 0),
	zend.MakeZendInternalArgInfo("mode", 0, 0, 0),
}
var ArginfoPclose []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
}
var ArginfoFeof []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
}
var ArginfoFgets []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
	zend.MakeZendInternalArgInfo("length", 0, 0, 0),
}
var ArginfoFgetc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
}
var ArginfoFgetss []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
	zend.MakeZendInternalArgInfo("length", 0, 0, 0),
	zend.MakeZendInternalArgInfo("allowable_tags", 0, 0, 0),
}
var ArginfoFscanf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
	zend.MakeZendInternalArgInfo("vars", 0, 1, 1),
}
var ArginfoFwrite []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("length", 0, 0, 0),
}
var ArginfoFflush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
}
var ArginfoRewind []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
}
var ArginfoFtell []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
}
var ArginfoFseek []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
	zend.MakeZendInternalArgInfo("offset", 0, 0, 0),
	zend.MakeZendInternalArgInfo("whence", 0, 0, 0),
}
var ArginfoMkdir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("pathname", 0, 0, 0),
	zend.MakeZendInternalArgInfo("mode", 0, 0, 0),
	zend.MakeZendInternalArgInfo("recursive", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoRmdir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("dirname", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoReadfile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("flags", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoUmask []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("mask", 0, 0, 0),
}
var ArginfoFpassthru []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
}
var ArginfoRename []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("old_name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("new_name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoUnlink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoFtruncate []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
	zend.MakeZendInternalArgInfo("size", 0, 0, 0),
}
var ArginfoFstat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
}
var ArginfoCopy []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("source_file", 0, 0, 0),
	zend.MakeZendInternalArgInfo("destination_file", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoFread []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
	zend.MakeZendInternalArgInfo("length", 0, 0, 0),
}
var ArginfoFputcsv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
	zend.MakeZendInternalArgInfo("fields", 0, 0, 0),
	zend.MakeZendInternalArgInfo("delimiter", 0, 0, 0),
	zend.MakeZendInternalArgInfo("enclosure", 0, 0, 0),
	zend.MakeZendInternalArgInfo("escape_char", 0, 0, 0),
}
var ArginfoFgetcsv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
	zend.MakeZendInternalArgInfo("length", 0, 0, 0),
	zend.MakeZendInternalArgInfo("delimiter", 0, 0, 0),
	zend.MakeZendInternalArgInfo("enclosure", 0, 0, 0),
	zend.MakeZendInternalArgInfo("escape", 0, 0, 0),
}
var ArginfoRealpath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("path", 0, 0, 0),
}
var ArginfoFnmatch []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("pattern", 0, 0, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("flags", 0, 0, 0),
}
var ArginfoSysGetTempDir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoDiskTotalSpace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("path", 0, 0, 0),
}
var ArginfoDiskFreeSpace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("path", 0, 0, 0),
}
var ArginfoChgrp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("group", 0, 0, 0),
}
var ArginfoChown []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("user", 0, 0, 0),
}
var ArginfoLchgrp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("group", 0, 0, 0),
}
var ArginfoLchown []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("user", 0, 0, 0),
}
var ArginfoChmod []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("mode", 0, 0, 0),
}
var ArginfoTouch []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("time", 0, 0, 0),
	zend.MakeZendInternalArgInfo("atime", 0, 0, 0),
}
var ArginfoClearstatcache []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("clear_realpath_cache", 0, 0, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoRealpathCacheSize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoRealpathCacheGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoFileperms []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoFileinode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoFilesize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoFileowner []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoFilegroup []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoFileatime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoFilemtime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoFilectime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoFiletype []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoIsWritable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoIsReadable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoIsExecutable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoIsFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoIsDir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoIsLink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoFileExists []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoLstat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoStat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoSprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
	zend.MakeZendInternalArgInfo("args", 0, 0, 1),
}
var ArginfoVsprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
	zend.MakeZendInternalArgInfo("args", 0, 0, 0),
}
var ArginfoPrintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
	zend.MakeZendInternalArgInfo("args", 0, 0, 1),
}
var ArginfoVprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
	zend.MakeZendInternalArgInfo("args", 0, 0, 0),
}
var ArginfoFprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
	zend.MakeZendInternalArgInfo("args", 0, 0, 1),
}
var ArginfoVfprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
	zend.MakeZendInternalArgInfo("args", 0, 0, 0),
}
var ArginfoFsockopen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("hostname", 0, 0, 0),
	zend.MakeZendInternalArgInfo("port", 0, 0, 0),
	zend.MakeZendInternalArgInfo("errno", 0, 1, 0),
	zend.MakeZendInternalArgInfo("errstr", 0, 1, 0),
	zend.MakeZendInternalArgInfo("timeout", 0, 0, 0),
}
var ArginfoPfsockopen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("hostname", 0, 0, 0),
	zend.MakeZendInternalArgInfo("port", 0, 0, 0),
	zend.MakeZendInternalArgInfo("errno", 0, 1, 0),
	zend.MakeZendInternalArgInfo("errstr", 0, 1, 0),
	zend.MakeZendInternalArgInfo("timeout", 0, 0, 0),
}
var ArginfoFtok []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("pathname", 0, 0, 0),
	zend.MakeZendInternalArgInfo("proj", 0, 0, 0),
}
var ArginfoHeader []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("header", 0, 0, 0),
	zend.MakeZendInternalArgInfo("replace", 0, 0, 0),
	zend.MakeZendInternalArgInfo("http_response_code", 0, 0, 0),
}
var ArginfoHeaderRemove []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("name", 0, 0, 0),
}
var ArginfoSetcookie []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("value", 0, 0, 0),
	zend.MakeZendInternalArgInfo("expires_or_options", 0, 0, 0),
	zend.MakeZendInternalArgInfo("path", 0, 0, 0),
	zend.MakeZendInternalArgInfo("domain", 0, 0, 0),
	zend.MakeZendInternalArgInfo("secure", 0, 0, 0),
	zend.MakeZendInternalArgInfo("httponly", 0, 0, 0),
}
var ArginfoSetrawcookie []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("value", 0, 0, 0),
	zend.MakeZendInternalArgInfo("expires_or_options", 0, 0, 0),
	zend.MakeZendInternalArgInfo("path", 0, 0, 0),
	zend.MakeZendInternalArgInfo("domain", 0, 0, 0),
	zend.MakeZendInternalArgInfo("secure", 0, 0, 0),
	zend.MakeZendInternalArgInfo("httponly", 0, 0, 0),
}
var ArginfoHeadersSent []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("file", 0, 1, 0),
	zend.MakeZendInternalArgInfo("line", 0, 1, 0),
}
var ArginfoHeadersList []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoHttpResponseCode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("response_code", 0, 0, 0),
}
var ArginfoHrtime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("get_as_number", 0, 0, 0),
}
var ArginfoHtmlspecialchars []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("string", 0, 0, 0),
	zend.MakeZendInternalArgInfo("quote_style", 0, 0, 0),
	zend.MakeZendInternalArgInfo("encoding", 0, 0, 0),
	zend.MakeZendInternalArgInfo("double_encode", 0, 0, 0),
}
var ArginfoHtmlspecialcharsDecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("string", 0, 0, 0),
	zend.MakeZendInternalArgInfo("quote_style", 0, 0, 0),
}
var ArginfoHtmlEntityDecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("string", 0, 0, 0),
	zend.MakeZendInternalArgInfo("quote_style", 0, 0, 0),
	zend.MakeZendInternalArgInfo("encoding", 0, 0, 0),
}
var ArginfoHtmlentities []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("string", 0, 0, 0),
	zend.MakeZendInternalArgInfo("quote_style", 0, 0, 0),
	zend.MakeZendInternalArgInfo("encoding", 0, 0, 0),
	zend.MakeZendInternalArgInfo("double_encode", 0, 0, 0),
}
var ArginfoGetHtmlTranslationTable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("table", 0, 0, 0),
	zend.MakeZendInternalArgInfo("quote_style", 0, 0, 0),
	zend.MakeZendInternalArgInfo("encoding", 0, 0, 0),
}
var ArginfoHttpBuildQuery []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("formdata", 0, 0, 0),
	zend.MakeZendInternalArgInfo("prefix", 0, 0, 0),
	zend.MakeZendInternalArgInfo("arg_separator", 0, 0, 0),
	zend.MakeZendInternalArgInfo("enc_type", 0, 0, 0),
}
var ArginfoImageTypeToMimeType []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("imagetype", 0, 0, 0),
}
var ArginfoImageTypeToExtension []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("imagetype", 0, 0, 0),
	zend.MakeZendInternalArgInfo("include_dot", 0, 0, 0),
}
var ArginfoGetimagesize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("imagefile", 0, 0, 0),
	zend.MakeZendInternalArgInfo("info", 0, 1, 0),
}
var ArginfoPhpinfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("what", 0, 0, 0),
}
var ArginfoPhpversion []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("extension", 0, 0, 0),
}
var ArginfoPhpcredits []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("flag", 0, 0, 0),
}
var ArginfoPhpSapiName []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoPhpUname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("mode", 0, 0, 0),
}
var ArginfoPhpIniScannedFiles []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoPhpIniLoadedFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoIptcembed []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("iptcdata", 0, 0, 0),
	zend.MakeZendInternalArgInfo("jpeg_file_name", 0, 0, 0),
	zend.MakeZendInternalArgInfo("spool", 0, 0, 0),
}
var ArginfoIptcparse []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("iptcdata", 0, 0, 0),
}
var ArginfoLcgValue []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoLevenshtein []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("str2", 0, 0, 0),
	zend.MakeZendInternalArgInfo("cost_ins", 0, 0, 0),
	zend.MakeZendInternalArgInfo("cost_rep", 0, 0, 0),
	zend.MakeZendInternalArgInfo("cost_del", 0, 0, 0),
}
var ArginfoReadlink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoLinkinfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoSymlink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("target", 0, 0, 0),
	zend.MakeZendInternalArgInfo("link", 0, 0, 0),
}
var ArginfoLink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("target", 0, 0, 0),
	zend.MakeZendInternalArgInfo("link", 0, 0, 0),
}
var ArginfoEzmlmHash []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("addr", 0, 0, 0),
}
var ArginfoMail []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(3)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("to", 0, 0, 0),
	zend.MakeZendInternalArgInfo("subject", 0, 0, 0),
	zend.MakeZendInternalArgInfo("message", 0, 0, 0),
	zend.MakeZendInternalArgInfo("additional_headers", 0, 0, 0),
	zend.MakeZendInternalArgInfo("additional_parameters", 0, 0, 0),
}
var ArginfoAbs []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoCeil []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoFloor []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoRound []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
	zend.MakeZendInternalArgInfo("precision", 0, 0, 0),
	zend.MakeZendInternalArgInfo("mode", 0, 0, 0),
}
var ArginfoSin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoCos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoTan []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoAsin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoAcos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoAtan []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoAtan2 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("y", 0, 0, 0),
	zend.MakeZendInternalArgInfo("x", 0, 0, 0),
}
var ArginfoSinh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoCosh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoTanh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoAsinh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoAcosh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoAtanh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoPi []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoIsFinite []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("val", 0, 0, 0),
}
var ArginfoIsInfinite []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("val", 0, 0, 0),
}
var ArginfoIsNan []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("val", 0, 0, 0),
}
var ArginfoPow []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("base", 0, 0, 0),
	zend.MakeZendInternalArgInfo("exponent", 0, 0, 0),
}
var ArginfoExp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoExpm1 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoLog1p []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoLog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
	zend.MakeZendInternalArgInfo("base", 0, 0, 0),
}
var ArginfoLog10 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoSqrt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoHypot []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("num1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("num2", 0, 0, 0),
}
var ArginfoDeg2rad []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoRad2deg []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
}
var ArginfoBindec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("binary_number", 0, 0, 0),
}
var ArginfoHexdec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("hexadecimal_number", 0, 0, 0),
}
var ArginfoOctdec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("octal_number", 0, 0, 0),
}
var ArginfoDecbin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("decimal_number", 0, 0, 0),
}
var ArginfoDecoct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("decimal_number", 0, 0, 0),
}
var ArginfoDechex []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("decimal_number", 0, 0, 0),
}
var ArginfoBaseConvert []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
	zend.MakeZendInternalArgInfo("frombase", 0, 0, 0),
	zend.MakeZendInternalArgInfo("tobase", 0, 0, 0),
}
var ArginfoNumberFormat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("number", 0, 0, 0),
	zend.MakeZendInternalArgInfo("num_decimal_places", 0, 0, 0),
	zend.MakeZendInternalArgInfo("dec_separator", 0, 0, 0),
	zend.MakeZendInternalArgInfo("thousands_separator", 0, 0, 0),
}
var ArginfoFmod []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("x", 0, 0, 0),
	zend.MakeZendInternalArgInfo("y", 0, 0, 0),
}
var ArginfoIntdiv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("dividend", 0, 0, 0),
	zend.MakeZendInternalArgInfo("divisor", 0, 0, 0),
}
var ArginfoMd5 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("raw_output", 0, 0, 0),
}
var ArginfoMd5File []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("raw_output", 0, 0, 0),
}
var ArginfoMetaphone []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("text", 0, 0, 0),
	zend.MakeZendInternalArgInfo("phones", 0, 0, 0),
}
var ArginfoMicrotime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("get_as_float", 0, 0, 0),
}
var ArginfoGettimeofday []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("get_as_float", 0, 0, 0),
}
var ArginfoGetrusage []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("who", 0, 0, 0),
}
var ArginfoPack []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
	zend.MakeZendInternalArgInfo("args", 0, 0, 1),
}
var ArginfoUnpack []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
	zend.MakeZendInternalArgInfo("input", 0, 0, 0),
	zend.MakeZendInternalArgInfo("offset", 0, 0, 0),
}
var ArginfoGetmyuid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoGetmygid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoGetmypid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoGetmyinode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoGetlastmod []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoPasswordHash []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("password", 0, 0, 0),
	zend.MakeZendInternalArgInfo("algo", 0, 0, 0),
	zend.MakeZendInternalArgInfo("options", 0, 0, 0),
}
var ArginfoPasswordGetInfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("hash", 0, 0, 0),
}
var ArginfoPasswordNeedsRehash []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("hash", 0, 0, 0),
	zend.MakeZendInternalArgInfo("algo", 0, 0, 0),
	zend.MakeZendInternalArgInfo("options", 0, 0, 0),
}
var ArginfoPasswordVerify []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("password", 0, 0, 0),
	zend.MakeZendInternalArgInfo("hash", 0, 0, 0),
}
var ArginfoPasswordAlgos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoProcTerminate []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("process", 0, 0, 0),
	zend.MakeZendInternalArgInfo("signal", 0, 0, 0),
}
var ArginfoProcClose []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("process", 0, 0, 0),
}
var ArginfoProcGetStatus []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("process", 0, 0, 0),
}
var ArginfoProcOpen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(3)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("command", 0, 0, 0),
	zend.MakeZendInternalArgInfo("descriptorspec", 0, 0, 0),
	zend.MakeZendInternalArgInfo("pipes", 0, 1, 0),
	zend.MakeZendInternalArgInfo("cwd", 0, 0, 0),
	zend.MakeZendInternalArgInfo("env", 0, 0, 0),
	zend.MakeZendInternalArgInfo("other_options", 0, 0, 0),
}
var ArginfoQuotedPrintableDecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoQuotedPrintableEncode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoMtSrand []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("seed", 0, 0, 0),
	zend.MakeZendInternalArgInfo("mode", 0, 0, 0),
}
var ArginfoMtRand []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("min", 0, 0, 0),
	zend.MakeZendInternalArgInfo("max", 0, 0, 0),
}
var ArginfoMtGetrandmax []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoRandomBytes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("length", 0, 0, 0),
}
var ArginfoRandomInt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("min", 0, 0, 0),
	zend.MakeZendInternalArgInfo("max", 0, 0, 0),
}
var ArginfoSha1 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("raw_output", 0, 0, 0),
}
var ArginfoSha1File []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
	zend.MakeZendInternalArgInfo("raw_output", 0, 0, 0),
}
var ArginfoSoundex []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoStreamSocketPair []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("domain", 0, 0, 0),
	zend.MakeZendInternalArgInfo("type", 0, 0, 0),
	zend.MakeZendInternalArgInfo("protocol", 0, 0, 0),
}
var ArginfoStreamSocketClient []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("remoteaddress", 0, 0, 0),
	zend.MakeZendInternalArgInfo("errcode", 0, 1, 0),
	zend.MakeZendInternalArgInfo("errstring", 0, 1, 0),
	zend.MakeZendInternalArgInfo("timeout", 0, 0, 0),
	zend.MakeZendInternalArgInfo("flags", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoStreamSocketServer []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("localaddress", 0, 0, 0),
	zend.MakeZendInternalArgInfo("errcode", 0, 1, 0),
	zend.MakeZendInternalArgInfo("errstring", 0, 1, 0),
	zend.MakeZendInternalArgInfo("flags", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoStreamSocketAccept []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("serverstream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("timeout", 0, 0, 0),
	zend.MakeZendInternalArgInfo("peername", 0, 1, 0),
}
var ArginfoStreamSocketGetName []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("want_peer", 0, 0, 0),
}
var ArginfoStreamSocketSendto []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("data", 0, 0, 0),
	zend.MakeZendInternalArgInfo("flags", 0, 0, 0),
	zend.MakeZendInternalArgInfo("target_addr", 0, 0, 0),
}
var ArginfoStreamSocketRecvfrom []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("amount", 0, 0, 0),
	zend.MakeZendInternalArgInfo("flags", 0, 0, 0),
	zend.MakeZendInternalArgInfo("remote_addr", 0, 1, 0),
}
var ArginfoStreamGetContents []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("source", 0, 0, 0),
	zend.MakeZendInternalArgInfo("maxlen", 0, 0, 0),
	zend.MakeZendInternalArgInfo("offset", 0, 0, 0),
}
var ArginfoStreamCopyToStream []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("source", 0, 0, 0),
	zend.MakeZendInternalArgInfo("dest", 0, 0, 0),
	zend.MakeZendInternalArgInfo("maxlen", 0, 0, 0),
	zend.MakeZendInternalArgInfo("pos", 0, 0, 0),
}
var ArginfoStreamGetMetaData []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
}
var ArginfoStreamGetTransports []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoStreamGetWrappers []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoStreamResolveIncludePath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filename", 0, 0, 0),
}
var ArginfoStreamIsLocal []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
}
var ArginfoStreamSupportsLock []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
}
var ArginfoStreamIsatty []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
}
var ArginfoStreamSelect []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(4)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("read_streams", 0, 1, 0),
	zend.MakeZendInternalArgInfo("write_streams", 0, 1, 0),
	zend.MakeZendInternalArgInfo("except_streams", 0, 1, 0),
	zend.MakeZendInternalArgInfo("tv_sec", 0, 0, 0),
	zend.MakeZendInternalArgInfo("tv_usec", 0, 0, 0),
}
var ArginfoStreamContextGetOptions []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("stream_or_context", 0, 0, 0),
}
var ArginfoStreamContextSetOption []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stream_or_context", 0, 0, 0),
	zend.MakeZendInternalArgInfo("wrappername", 0, 0, 0),
	zend.MakeZendInternalArgInfo("optionname", 0, 0, 0),
	zend.MakeZendInternalArgInfo("value", 0, 0, 0),
}
var ArginfoStreamContextSetParams []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("stream_or_context", 0, 0, 0),
	zend.MakeZendInternalArgInfo("options", 0, 0, 0),
}
var ArginfoStreamContextGetParams []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("stream_or_context", 0, 0, 0),
}
var ArginfoStreamContextGetDefault []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("options", 0, 0, 0),
}
var ArginfoStreamContextSetDefault []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("options", 0, 0, 0),
}
var ArginfoStreamContextCreate []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("options", 0, 0, 0),
	zend.MakeZendInternalArgInfo("params", 0, 0, 0),
}
var ArginfoStreamFilterPrepend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("filtername", 0, 0, 0),
	zend.MakeZendInternalArgInfo("read_write", 0, 0, 0),
	zend.MakeZendInternalArgInfo("filterparams", 0, 0, 0),
}
var ArginfoStreamFilterAppend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("filtername", 0, 0, 0),
	zend.MakeZendInternalArgInfo("read_write", 0, 0, 0),
	zend.MakeZendInternalArgInfo("filterparams", 0, 0, 0),
}
var ArginfoStreamFilterRemove []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("stream_filter", 0, 0, 0),
}
var ArginfoStreamGetLine []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("maxlen", 0, 0, 0),
	zend.MakeZendInternalArgInfo("ending", 0, 0, 0),
}
var ArginfoStreamSetBlocking []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("socket", 0, 0, 0),
	zend.MakeZendInternalArgInfo("mode", 0, 0, 0),
}
var ArginfoStreamSetTimeout []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("seconds", 0, 0, 0),
	zend.MakeZendInternalArgInfo("microseconds", 0, 0, 0),
}
var ArginfoStreamSetReadBuffer []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
	zend.MakeZendInternalArgInfo("buffer", 0, 0, 0),
}
var ArginfoStreamSetWriteBuffer []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
	zend.MakeZendInternalArgInfo("buffer", 0, 0, 0),
}
var ArginfoStreamSetChunkSize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("fp", 0, 0, 0),
	zend.MakeZendInternalArgInfo("chunk_size", 0, 0, 0),
}
var ArginfoStreamSocketEnableCrypto []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("enable", 0, 0, 0),
	zend.MakeZendInternalArgInfo("cryptokind", 0, 0, 0),
	zend.MakeZendInternalArgInfo("sessionstream", 0, 0, 0),
}
var ArginfoStreamSocketShutdown []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("how", 0, 0, 0),
}
var ArginfoBin2hex []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("data", 0, 0, 0),
}
var ArginfoHex2bin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("data", 0, 0, 0),
}
var ArginfoStrspn []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("mask", 0, 0, 0),
	zend.MakeZendInternalArgInfo("start", 0, 0, 0),
	zend.MakeZendInternalArgInfo("len", 0, 0, 0),
}
var ArginfoStrcspn []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("mask", 0, 0, 0),
	zend.MakeZendInternalArgInfo("start", 0, 0, 0),
	zend.MakeZendInternalArgInfo("len", 0, 0, 0),
}
var ArginfoStrcoll []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("str2", 0, 0, 0),
}
var ArginfoTrim []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("character_mask", 0, 0, 0),
}
var ArginfoRtrim []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("character_mask", 0, 0, 0),
}
var ArginfoLtrim []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("character_mask", 0, 0, 0),
}
var ArginfoWordwrap []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("width", 0, 0, 0),
	zend.MakeZendInternalArgInfo("break", 0, 0, 0),
	zend.MakeZendInternalArgInfo("cut", 0, 0, 0),
}
var ArginfoExplode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("separator", 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("limit", 0, 0, 0),
}
var ArginfoImplode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("glue", 0, 0, 0),
	zend.MakeZendInternalArgInfo("pieces", 0, 0, 0),
}
var ArginfoStrtok []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("token", 0, 0, 0),
}
var ArginfoStrtoupper []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoStrtolower []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoBasename []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("path", 0, 0, 0),
	zend.MakeZendInternalArgInfo("suffix", 0, 0, 0),
}
var ArginfoDirname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("path", 0, 0, 0),
	zend.MakeZendInternalArgInfo("levels", 0, 0, 0),
}
var ArginfoPathinfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("path", 0, 0, 0),
	zend.MakeZendInternalArgInfo("options", 0, 0, 0),
}
var ArginfoStristr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("haystack", 0, 0, 0),
	zend.MakeZendInternalArgInfo("needle", 0, 0, 0),
	zend.MakeZendInternalArgInfo("part", 0, 0, 0),
}
var ArginfoStrstr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("haystack", 0, 0, 0),
	zend.MakeZendInternalArgInfo("needle", 0, 0, 0),
	zend.MakeZendInternalArgInfo("part", 0, 0, 0),
}
var ArginfoStrpos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("haystack", 0, 0, 0),
	zend.MakeZendInternalArgInfo("needle", 0, 0, 0),
	zend.MakeZendInternalArgInfo("offset", 0, 0, 0),
}
var ArginfoStripos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("haystack", 0, 0, 0),
	zend.MakeZendInternalArgInfo("needle", 0, 0, 0),
	zend.MakeZendInternalArgInfo("offset", 0, 0, 0),
}
var ArginfoStrrpos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("haystack", 0, 0, 0),
	zend.MakeZendInternalArgInfo("needle", 0, 0, 0),
	zend.MakeZendInternalArgInfo("offset", 0, 0, 0),
}
var ArginfoStrripos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("haystack", 0, 0, 0),
	zend.MakeZendInternalArgInfo("needle", 0, 0, 0),
	zend.MakeZendInternalArgInfo("offset", 0, 0, 0),
}
var ArginfoStrrchr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("haystack", 0, 0, 0),
	zend.MakeZendInternalArgInfo("needle", 0, 0, 0),
}
var ArginfoChunkSplit []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("chunklen", 0, 0, 0),
	zend.MakeZendInternalArgInfo("ending", 0, 0, 0),
}
var ArginfoSubstr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("start", 0, 0, 0),
	zend.MakeZendInternalArgInfo("length", 0, 0, 0),
}
var ArginfoSubstrReplace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(3)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("replace", 0, 0, 0),
	zend.MakeZendInternalArgInfo("start", 0, 0, 0),
	zend.MakeZendInternalArgInfo("length", 0, 0, 0),
}
var ArginfoQuotemeta []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoOrd []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("character", 0, 0, 0),
}
var ArginfoChr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("codepoint", 0, 0, 0),
}
var ArginfoUcfirst []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoLcfirst []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoUcwords []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("delimiters", 0, 0, 0),
}
var ArginfoStrtr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("from", 0, 0, 0),
	zend.MakeZendInternalArgInfo("to", 0, 0, 0),
}
var ArginfoStrrev []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoSimilarText []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("str2", 0, 0, 0),
	zend.MakeZendInternalArgInfo("percent", 0, 1, 0),
}
var ArginfoAddcslashes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("charlist", 0, 0, 0),
}
var ArginfoAddslashes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoStripcslashes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoStripslashes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoStrReplace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(3)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("search", 0, 0, 0),
	zend.MakeZendInternalArgInfo("replace", 0, 0, 0),
	zend.MakeZendInternalArgInfo("subject", 0, 0, 0),
	zend.MakeZendInternalArgInfo("replace_count", 0, 1, 0),
}
var ArginfoStrIreplace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(3)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("search", 0, 0, 0),
	zend.MakeZendInternalArgInfo("replace", 0, 0, 0),
	zend.MakeZendInternalArgInfo("subject", 0, 0, 0),
	zend.MakeZendInternalArgInfo("replace_count", 0, 1, 0),
}
var ArginfoHebrev []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("max_chars_per_line", 0, 0, 0),
}
var ArginfoHebrevc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("max_chars_per_line", 0, 0, 0),
}
var ArginfoNl2br []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("is_xhtml", 0, 0, 0),
}
var ArginfoStripTags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("allowable_tags", 0, 0, 0),
}
var ArginfoSetlocale []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("category", 0, 0, 0),
	zend.MakeZendInternalArgInfo("locales", 0, 0, 1),
}
var ArginfoParseStr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("encoded_string", 0, 0, 0),
	zend.MakeZendInternalArgInfo("result", 0, 1, 0),
}
var ArginfoStrGetcsv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("string", 0, 0, 0),
	zend.MakeZendInternalArgInfo("delimiter", 0, 0, 0),
	zend.MakeZendInternalArgInfo("enclosure", 0, 0, 0),
	zend.MakeZendInternalArgInfo("escape", 0, 0, 0),
}
var ArginfoStrRepeat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("input", 0, 0, 0),
	zend.MakeZendInternalArgInfo("mult", 0, 0, 0),
}
var ArginfoCountChars []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("input", 0, 0, 0),
	zend.MakeZendInternalArgInfo("mode", 0, 0, 0),
}
var ArginfoStrnatcmp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("s1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("s2", 0, 0, 0),
}
var ArginfoLocaleconv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoStrnatcasecmp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("s1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("s2", 0, 0, 0),
}
var ArginfoSubstrCount []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("haystack", 0, 0, 0),
	zend.MakeZendInternalArgInfo("needle", 0, 0, 0),
	zend.MakeZendInternalArgInfo("offset", 0, 0, 0),
	zend.MakeZendInternalArgInfo("length", 0, 0, 0),
}
var ArginfoStrPad []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("input", 0, 0, 0),
	zend.MakeZendInternalArgInfo("pad_length", 0, 0, 0),
	zend.MakeZendInternalArgInfo("pad_string", 0, 0, 0),
	zend.MakeZendInternalArgInfo("pad_type", 0, 0, 0),
}
var ArginfoSscanf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
	zend.MakeZendInternalArgInfo("vars", 0, 1, 1),
}
var ArginfoStrRot13 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoStrShuffle []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoStrWordCount []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
	zend.MakeZendInternalArgInfo("charlist", 0, 0, 0),
}
var ArginfoMoneyFormat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
	zend.MakeZendInternalArgInfo("value", 0, 0, 0),
}
var ArginfoStrSplit []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("split_length", 0, 0, 0),
}
var ArginfoStrpbrk []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("haystack", 0, 0, 0),
	zend.MakeZendInternalArgInfo("char_list", 0, 0, 0),
}
var ArginfoSubstrCompare []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(3)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("main_str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
	zend.MakeZendInternalArgInfo("offset", 0, 0, 0),
	zend.MakeZendInternalArgInfo("length", 0, 0, 0),
	zend.MakeZendInternalArgInfo("case_sensitivity", 0, 0, 0),
}
var ArginfoUtf8Encode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("data", 0, 0, 0),
}
var ArginfoUtf8Decode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("data", 0, 0, 0),
}
var ArginfoOpenlog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("ident", 0, 0, 0),
	zend.MakeZendInternalArgInfo("option", 0, 0, 0),
	zend.MakeZendInternalArgInfo("facility", 0, 0, 0),
}
var ArginfoCloselog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoSyslog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("priority", 0, 0, 0),
	zend.MakeZendInternalArgInfo("message", 0, 0, 0),
}
var ArginfoGettype []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoSettype []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 1, 0),
	zend.MakeZendInternalArgInfo("type", 0, 0, 0),
}
var ArginfoIntval []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
	zend.MakeZendInternalArgInfo("base", 0, 0, 0),
}
var ArginfoFloatval []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoStrval []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoBoolval []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoIsNull []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoIsResource []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoIsBool []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoIsInt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoIsFloat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoIsString []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoIsArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoIsObject []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoIsNumeric []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("value", 0, 0, 0),
}
var ArginfoIsScalar []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("value", 0, 0, 0),
}
var ArginfoIsCallable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
	zend.MakeZendInternalArgInfo("syntax_only", 0, 0, 0),
	zend.MakeZendInternalArgInfo("callable_name", 0, 1, 0),
}
var ArginfoIsIterable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoIsCountable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoUniqid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("prefix", 0, 0, 0),
	zend.MakeZendInternalArgInfo("more_entropy", 0, 0, 0),
}
var ArginfoParseUrl []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("url", 0, 0, 0),
	zend.MakeZendInternalArgInfo("component", 0, 0, 0),
}
var ArginfoUrlencode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoUrldecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoRawurlencode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoRawurldecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("str", 0, 0, 0),
}
var ArginfoGetHeaders []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("url", 0, 0, 0),
	zend.MakeZendInternalArgInfo("format", 0, 0, 0),
	zend.MakeZendInternalArgInfo("context", 0, 0, 0),
}
var ArginfoStreamBucketMakeWriteable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("brigade", 0, 0, 0),
}
var ArginfoStreamBucketPrepend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("brigade", 0, 0, 0),
	zend.MakeZendInternalArgInfo("bucket", 0, 0, 0),
}
var ArginfoStreamBucketAppend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("brigade", 0, 0, 0),
	zend.MakeZendInternalArgInfo("bucket", 0, 0, 0),
}
var ArginfoStreamBucketNew []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("stream", 0, 0, 0),
	zend.MakeZendInternalArgInfo("buffer", 0, 0, 0),
}
var ArginfoStreamGetFilters []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ArginfoStreamFilterRegister []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("filtername", 0, 0, 0),
	zend.MakeZendInternalArgInfo("classname", 0, 0, 0),
}
var ArginfoConvertUuencode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("data", 0, 0, 0),
}
var ArginfoConvertUudecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("data", 0, 0, 0),
}
var ArginfoVarDump []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("vars", 0, 0, 1),
}
var ArginfoDebugZvalDump []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("vars", 0, 0, 1),
}
var ArginfoVarExport []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
	zend.MakeZendInternalArgInfo("return", 0, 0, 0),
}
var ArginfoSerialize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("var", 0, 0, 0),
}
var ArginfoUnserialize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("variable_representation", 0, 0, 0),
	zend.MakeZendInternalArgInfo("allowed_classes", 0, 0, 0),
}
var ArginfoMemoryGetUsage []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("real_usage", 0, 0, 0),
}
var ArginfoMemoryGetPeakUsage []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(0)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("real_usage", 0, 0, 0),
}
var ArginfoVersionCompare []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	zend.MakeZendInternalArgInfo("ver1", 0, 0, 0),
	zend.MakeZendInternalArgInfo("ver2", 0, 0, 0),
	zend.MakeZendInternalArgInfo("oper", 0, 0, 0),
}
var BasicFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("constant", ZifConstant, ArginfoConstant, uint32(b.SizeOf("arginfo_constant")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("bin2hex", ZifBin2hex, ArginfoBin2hex, uint32(b.SizeOf("arginfo_bin2hex")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("hex2bin", ZifHex2bin, ArginfoHex2bin, uint32(b.SizeOf("arginfo_hex2bin")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("sleep", ZifSleep, ArginfoSleep, uint32(b.SizeOf("arginfo_sleep")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("usleep", ZifUsleep, ArginfoUsleep, uint32(b.SizeOf("arginfo_usleep")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("time_nanosleep", ZifTimeNanosleep, ArginfoTimeNanosleep, uint32(b.SizeOf("arginfo_time_nanosleep")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("time_sleep_until", ZifTimeSleepUntil, ArginfoTimeSleepUntil, uint32(b.SizeOf("arginfo_time_sleep_until")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strptime", ZifStrptime, ArginfoStrptime, uint32(b.SizeOf("arginfo_strptime")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("flush", ZifFlush, ArginfoFlush, uint32(b.SizeOf("arginfo_flush")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("wordwrap", ZifWordwrap, ArginfoWordwrap, uint32(b.SizeOf("arginfo_wordwrap")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("htmlspecialchars", ZifHtmlspecialchars, ArginfoHtmlspecialchars, uint32(b.SizeOf("arginfo_htmlspecialchars")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("htmlentities", ZifHtmlentities, ArginfoHtmlentities, uint32(b.SizeOf("arginfo_htmlentities")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("html_entity_decode", ZifHtmlEntityDecode, ArginfoHtmlEntityDecode, uint32(b.SizeOf("arginfo_html_entity_decode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("htmlspecialchars_decode", ZifHtmlspecialcharsDecode, ArginfoHtmlspecialcharsDecode, uint32(b.SizeOf("arginfo_htmlspecialchars_decode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("get_html_translation_table", ZifGetHtmlTranslationTable, ArginfoGetHtmlTranslationTable, uint32(b.SizeOf("arginfo_get_html_translation_table")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("sha1", ZifSha1, ArginfoSha1, uint32(b.SizeOf("arginfo_sha1")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("sha1_file", ZifSha1File, ArginfoSha1File, uint32(b.SizeOf("arginfo_sha1_file")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("md5", PhpIfMd5, ArginfoMd5, uint32(b.SizeOf("arginfo_md5")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("md5_file", PhpIfMd5File, ArginfoMd5File, uint32(b.SizeOf("arginfo_md5_file")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("crc32", PhpIfCrc32, ArginfoCrc32, uint32(b.SizeOf("arginfo_crc32")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("iptcparse", ZifIptcparse, ArginfoIptcparse, uint32(b.SizeOf("arginfo_iptcparse")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("iptcembed", ZifIptcembed, ArginfoIptcembed, uint32(b.SizeOf("arginfo_iptcembed")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getimagesize", ZifGetimagesize, ArginfoGetimagesize, uint32(b.SizeOf("arginfo_getimagesize")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getimagesizefromstring", ZifGetimagesizefromstring, ArginfoGetimagesize, uint32(b.SizeOf("arginfo_getimagesize")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("image_type_to_mime_type", ZifImageTypeToMimeType, ArginfoImageTypeToMimeType, uint32(b.SizeOf("arginfo_image_type_to_mime_type")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("image_type_to_extension", ZifImageTypeToExtension, ArginfoImageTypeToExtension, uint32(b.SizeOf("arginfo_image_type_to_extension")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("phpversion", ZifPhpversion, ArginfoPhpversion, uint32(b.SizeOf("arginfo_phpversion")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("phpcredits", ZifPhpcredits, ArginfoPhpcredits, uint32(b.SizeOf("arginfo_phpcredits")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("php_sapi_name", ZifPhpSapiName, ArginfoPhpSapiName, uint32(b.SizeOf("arginfo_php_sapi_name")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("php_uname", ZifPhpUname, ArginfoPhpUname, uint32(b.SizeOf("arginfo_php_uname")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("php_ini_scanned_files", ZifPhpIniScannedFiles, ArginfoPhpIniScannedFiles, uint32(b.SizeOf("arginfo_php_ini_scanned_files")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("php_ini_loaded_file", ZifPhpIniLoadedFile, ArginfoPhpIniLoadedFile, uint32(b.SizeOf("arginfo_php_ini_loaded_file")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strnatcmp", ZifStrnatcmp, ArginfoStrnatcmp, uint32(b.SizeOf("arginfo_strnatcmp")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strnatcasecmp", ZifStrnatcasecmp, ArginfoStrnatcasecmp, uint32(b.SizeOf("arginfo_strnatcasecmp")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("substr_count", ZifSubstrCount, ArginfoSubstrCount, uint32(b.SizeOf("arginfo_substr_count")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strspn", ZifStrspn, ArginfoStrspn, uint32(b.SizeOf("arginfo_strspn")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strcspn", ZifStrcspn, ArginfoStrcspn, uint32(b.SizeOf("arginfo_strcspn")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strtok", ZifStrtok, ArginfoStrtok, uint32(b.SizeOf("arginfo_strtok")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strtoupper", ZifStrtoupper, ArginfoStrtoupper, uint32(b.SizeOf("arginfo_strtoupper")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strtolower", ZifStrtolower, ArginfoStrtolower, uint32(b.SizeOf("arginfo_strtolower")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strpos", ZifStrpos, ArginfoStrpos, uint32(b.SizeOf("arginfo_strpos")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stripos", ZifStripos, ArginfoStripos, uint32(b.SizeOf("arginfo_stripos")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strrpos", ZifStrrpos, ArginfoStrrpos, uint32(b.SizeOf("arginfo_strrpos")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strripos", ZifStrripos, ArginfoStrripos, uint32(b.SizeOf("arginfo_strripos")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strrev", ZifStrrev, ArginfoStrrev, uint32(b.SizeOf("arginfo_strrev")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("hebrev", ZifHebrev, ArginfoHebrev, uint32(b.SizeOf("arginfo_hebrev")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("hebrevc", ZifHebrevc, ArginfoHebrevc, uint32(b.SizeOf("arginfo_hebrevc")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_DEPRECATED),
	zend.MakeZendFunctionEntry("nl2br", ZifNl2br, ArginfoNl2br, uint32(b.SizeOf("arginfo_nl2br")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("basename", ZifBasename, ArginfoBasename, uint32(b.SizeOf("arginfo_basename")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("dirname", ZifDirname, ArginfoDirname, uint32(b.SizeOf("arginfo_dirname")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("pathinfo", ZifPathinfo, ArginfoPathinfo, uint32(b.SizeOf("arginfo_pathinfo")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stripslashes", ZifStripslashes, ArginfoStripslashes, uint32(b.SizeOf("arginfo_stripslashes")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stripcslashes", ZifStripcslashes, ArginfoStripcslashes, uint32(b.SizeOf("arginfo_stripcslashes")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strstr", ZifStrstr, ArginfoStrstr, uint32(b.SizeOf("arginfo_strstr")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stristr", ZifStristr, ArginfoStristr, uint32(b.SizeOf("arginfo_stristr")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strrchr", ZifStrrchr, ArginfoStrrchr, uint32(b.SizeOf("arginfo_strrchr")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("str_shuffle", ZifStrShuffle, ArginfoStrShuffle, uint32(b.SizeOf("arginfo_str_shuffle")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("str_word_count", ZifStrWordCount, ArginfoStrWordCount, uint32(b.SizeOf("arginfo_str_word_count")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("str_split", ZifStrSplit, ArginfoStrSplit, uint32(b.SizeOf("arginfo_str_split")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strpbrk", ZifStrpbrk, ArginfoStrpbrk, uint32(b.SizeOf("arginfo_strpbrk")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("substr_compare", ZifSubstrCompare, ArginfoSubstrCompare, uint32(b.SizeOf("arginfo_substr_compare")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("utf8_encode", ZifUtf8Encode, ArginfoUtf8Encode, uint32(b.SizeOf("arginfo_utf8_encode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("utf8_decode", ZifUtf8Decode, ArginfoUtf8Decode, uint32(b.SizeOf("arginfo_utf8_decode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strcoll", ZifStrcoll, ArginfoStrcoll, uint32(b.SizeOf("arginfo_strcoll")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("money_format", ZifMoneyFormat, ArginfoMoneyFormat, uint32(b.SizeOf("arginfo_money_format")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_DEPRECATED),
	zend.MakeZendFunctionEntry("substr", ZifSubstr, ArginfoSubstr, uint32(b.SizeOf("arginfo_substr")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("substr_replace", ZifSubstrReplace, ArginfoSubstrReplace, uint32(b.SizeOf("arginfo_substr_replace")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("quotemeta", ZifQuotemeta, ArginfoQuotemeta, uint32(b.SizeOf("arginfo_quotemeta")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ucfirst", ZifUcfirst, ArginfoUcfirst, uint32(b.SizeOf("arginfo_ucfirst")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("lcfirst", ZifLcfirst, ArginfoLcfirst, uint32(b.SizeOf("arginfo_lcfirst")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ucwords", ZifUcwords, ArginfoUcwords, uint32(b.SizeOf("arginfo_ucwords")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strtr", ZifStrtr, ArginfoStrtr, uint32(b.SizeOf("arginfo_strtr")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("addslashes", ZifAddslashes, ArginfoAddslashes, uint32(b.SizeOf("arginfo_addslashes")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("addcslashes", ZifAddcslashes, ArginfoAddcslashes, uint32(b.SizeOf("arginfo_addcslashes")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("rtrim", ZifRtrim, ArginfoRtrim, uint32(b.SizeOf("arginfo_rtrim")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("str_replace", ZifStrReplace, ArginfoStrReplace, uint32(b.SizeOf("arginfo_str_replace")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("str_ireplace", ZifStrIreplace, ArginfoStrIreplace, uint32(b.SizeOf("arginfo_str_ireplace")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("str_repeat", ZifStrRepeat, ArginfoStrRepeat, uint32(b.SizeOf("arginfo_str_repeat")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("count_chars", ZifCountChars, ArginfoCountChars, uint32(b.SizeOf("arginfo_count_chars")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("chunk_split", ZifChunkSplit, ArginfoChunkSplit, uint32(b.SizeOf("arginfo_chunk_split")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("trim", ZifTrim, ArginfoTrim, uint32(b.SizeOf("arginfo_trim")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ltrim", ZifLtrim, ArginfoLtrim, uint32(b.SizeOf("arginfo_ltrim")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strip_tags", ZifStripTags, ArginfoStripTags, uint32(b.SizeOf("arginfo_strip_tags")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("similar_text", ZifSimilarText, ArginfoSimilarText, uint32(b.SizeOf("arginfo_similar_text")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("explode", ZifExplode, ArginfoExplode, uint32(b.SizeOf("arginfo_explode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("implode", ZifImplode, ArginfoImplode, uint32(b.SizeOf("arginfo_implode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("join", ZifImplode, ArginfoImplode, uint32(b.SizeOf("arginfo_implode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("setlocale", ZifSetlocale, ArginfoSetlocale, uint32(b.SizeOf("arginfo_setlocale")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("localeconv", ZifLocaleconv, ArginfoLocaleconv, uint32(b.SizeOf("arginfo_localeconv")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("soundex", ZifSoundex, ArginfoSoundex, uint32(b.SizeOf("arginfo_soundex")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("levenshtein", ZifLevenshtein, ArginfoLevenshtein, uint32(b.SizeOf("arginfo_levenshtein")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("chr", ZifChr, ArginfoChr, uint32(b.SizeOf("arginfo_chr")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ord", ZifOrd, ArginfoOrd, uint32(b.SizeOf("arginfo_ord")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("parse_str", ZifParseStr, ArginfoParseStr, uint32(b.SizeOf("arginfo_parse_str")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("str_getcsv", ZifStrGetcsv, ArginfoStrGetcsv, uint32(b.SizeOf("arginfo_str_getcsv")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("str_pad", ZifStrPad, ArginfoStrPad, uint32(b.SizeOf("arginfo_str_pad")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("chop", ZifRtrim, ArginfoRtrim, uint32(b.SizeOf("arginfo_rtrim")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strchr", ZifStrstr, ArginfoStrstr, uint32(b.SizeOf("arginfo_strstr")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("sprintf", ZifUserSprintf, ArginfoSprintf, uint32(b.SizeOf("arginfo_sprintf")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("printf", ZifUserPrintf, ArginfoPrintf, uint32(b.SizeOf("arginfo_printf")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("vprintf", ZifVprintf, ArginfoVprintf, uint32(b.SizeOf("arginfo_vprintf")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("vsprintf", ZifVsprintf, ArginfoVsprintf, uint32(b.SizeOf("arginfo_vsprintf")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fprintf", ZifFprintf, ArginfoFprintf, uint32(b.SizeOf("arginfo_fprintf")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("vfprintf", ZifVfprintf, ArginfoVfprintf, uint32(b.SizeOf("arginfo_vfprintf")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("sscanf", ZifSscanf, ArginfoSscanf, uint32(b.SizeOf("arginfo_sscanf")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fscanf", ZifFscanf, ArginfoFscanf, uint32(b.SizeOf("arginfo_fscanf")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("parse_url", ZifParseUrl, ArginfoParseUrl, uint32(b.SizeOf("arginfo_parse_url")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("urlencode", ZifUrlencode, ArginfoUrlencode, uint32(b.SizeOf("arginfo_urlencode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("urldecode", ZifUrldecode, ArginfoUrldecode, uint32(b.SizeOf("arginfo_urldecode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("rawurlencode", ZifRawurlencode, ArginfoRawurlencode, uint32(b.SizeOf("arginfo_rawurlencode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("rawurldecode", ZifRawurldecode, ArginfoRawurldecode, uint32(b.SizeOf("arginfo_rawurldecode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("http_build_query", ZifHttpBuildQuery, ArginfoHttpBuildQuery, uint32(b.SizeOf("arginfo_http_build_query")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("readlink", ZifReadlink, ArginfoReadlink, uint32(b.SizeOf("arginfo_readlink")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("linkinfo", ZifLinkinfo, ArginfoLinkinfo, uint32(b.SizeOf("arginfo_linkinfo")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("symlink", ZifSymlink, ArginfoSymlink, uint32(b.SizeOf("arginfo_symlink")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("link", ZifLink, ArginfoLink, uint32(b.SizeOf("arginfo_link")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("unlink", ZifUnlink, ArginfoUnlink, uint32(b.SizeOf("arginfo_unlink")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("exec", ZifExec, ArginfoExec, uint32(b.SizeOf("arginfo_exec")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("system", ZifSystem, ArginfoSystem, uint32(b.SizeOf("arginfo_system")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("escapeshellcmd", ZifEscapeshellcmd, ArginfoEscapeshellcmd, uint32(b.SizeOf("arginfo_escapeshellcmd")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("escapeshellarg", ZifEscapeshellarg, ArginfoEscapeshellarg, uint32(b.SizeOf("arginfo_escapeshellarg")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("passthru", ZifPassthru, ArginfoPassthru, uint32(b.SizeOf("arginfo_passthru")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("shell_exec", ZifShellExec, ArginfoShellExec, uint32(b.SizeOf("arginfo_shell_exec")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("proc_open", ZifProcOpen, ArginfoProcOpen, uint32(b.SizeOf("arginfo_proc_open")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("proc_close", ZifProcClose, ArginfoProcClose, uint32(b.SizeOf("arginfo_proc_close")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("proc_terminate", ZifProcTerminate, ArginfoProcTerminate, uint32(b.SizeOf("arginfo_proc_terminate")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("proc_get_status", ZifProcGetStatus, ArginfoProcGetStatus, uint32(b.SizeOf("arginfo_proc_get_status")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("proc_nice", ZifProcNice, ArginfoProcNice, uint32(b.SizeOf("arginfo_proc_nice")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("rand", ZifRand, ArginfoMtRand, uint32(b.SizeOf("arginfo_mt_rand")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("srand", ZifMtSrand, ArginfoMtSrand, uint32(b.SizeOf("arginfo_mt_srand")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getrandmax", ZifMtGetrandmax, ArginfoMtGetrandmax, uint32(b.SizeOf("arginfo_mt_getrandmax")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("mt_rand", ZifMtRand, ArginfoMtRand, uint32(b.SizeOf("arginfo_mt_rand")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("mt_srand", ZifMtSrand, ArginfoMtSrand, uint32(b.SizeOf("arginfo_mt_srand")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("mt_getrandmax", ZifMtGetrandmax, ArginfoMtGetrandmax, uint32(b.SizeOf("arginfo_mt_getrandmax")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("random_bytes", ZifRandomBytes, ArginfoRandomBytes, uint32(b.SizeOf("arginfo_random_bytes")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("random_int", ZifRandomInt, ArginfoRandomInt, uint32(b.SizeOf("arginfo_random_int")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getservbyname", ZifGetservbyname, ArginfoGetservbyname, uint32(b.SizeOf("arginfo_getservbyname")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getservbyport", ZifGetservbyport, ArginfoGetservbyport, uint32(b.SizeOf("arginfo_getservbyport")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getprotobyname", ZifGetprotobyname, ArginfoGetprotobyname, uint32(b.SizeOf("arginfo_getprotobyname")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getprotobynumber", ZifGetprotobynumber, ArginfoGetprotobynumber, uint32(b.SizeOf("arginfo_getprotobynumber")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getmyuid", ZifGetmyuid, ArginfoGetmyuid, uint32(b.SizeOf("arginfo_getmyuid")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getmygid", ZifGetmygid, ArginfoGetmygid, uint32(b.SizeOf("arginfo_getmygid")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getmypid", ZifGetmypid, ArginfoGetmypid, uint32(b.SizeOf("arginfo_getmypid")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getmyinode", ZifGetmyinode, ArginfoGetmyinode, uint32(b.SizeOf("arginfo_getmyinode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getlastmod", ZifGetlastmod, ArginfoGetlastmod, uint32(b.SizeOf("arginfo_getlastmod")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("base64_decode", ZifBase64Decode, ArginfoBase64Decode, uint32(b.SizeOf("arginfo_base64_decode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("base64_encode", ZifBase64Encode, ArginfoBase64Encode, uint32(b.SizeOf("arginfo_base64_encode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("password_hash", ZifPasswordHash, ArginfoPasswordHash, uint32(b.SizeOf("arginfo_password_hash")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("password_get_info", ZifPasswordGetInfo, ArginfoPasswordGetInfo, uint32(b.SizeOf("arginfo_password_get_info")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("password_needs_rehash", ZifPasswordNeedsRehash, ArginfoPasswordNeedsRehash, uint32(b.SizeOf("arginfo_password_needs_rehash")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("password_verify", ZifPasswordVerify, ArginfoPasswordVerify, uint32(b.SizeOf("arginfo_password_verify")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("password_algos", ZifPasswordAlgos, ArginfoPasswordAlgos, uint32(b.SizeOf("arginfo_password_algos")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("convert_uuencode", ZifConvertUuencode, ArginfoConvertUuencode, uint32(b.SizeOf("arginfo_convert_uuencode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("convert_uudecode", ZifConvertUudecode, ArginfoConvertUudecode, uint32(b.SizeOf("arginfo_convert_uudecode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("abs", ZifAbs, ArginfoAbs, uint32(b.SizeOf("arginfo_abs")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ceil", ZifCeil, ArginfoCeil, uint32(b.SizeOf("arginfo_ceil")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("floor", ZifFloor, ArginfoFloor, uint32(b.SizeOf("arginfo_floor")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("round", ZifRound, ArginfoRound, uint32(b.SizeOf("arginfo_round")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("sin", ZifSin, ArginfoSin, uint32(b.SizeOf("arginfo_sin")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("cos", ZifCos, ArginfoCos, uint32(b.SizeOf("arginfo_cos")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("tan", ZifTan, ArginfoTan, uint32(b.SizeOf("arginfo_tan")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("asin", ZifAsin, ArginfoAsin, uint32(b.SizeOf("arginfo_asin")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("acos", ZifAcos, ArginfoAcos, uint32(b.SizeOf("arginfo_acos")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("atan", ZifAtan, ArginfoAtan, uint32(b.SizeOf("arginfo_atan")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("atanh", ZifAtanh, ArginfoAtanh, uint32(b.SizeOf("arginfo_atanh")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("atan2", ZifAtan2, ArginfoAtan2, uint32(b.SizeOf("arginfo_atan2")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("sinh", ZifSinh, ArginfoSinh, uint32(b.SizeOf("arginfo_sinh")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("cosh", ZifCosh, ArginfoCosh, uint32(b.SizeOf("arginfo_cosh")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("tanh", ZifTanh, ArginfoTanh, uint32(b.SizeOf("arginfo_tanh")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("asinh", ZifAsinh, ArginfoAsinh, uint32(b.SizeOf("arginfo_asinh")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("acosh", ZifAcosh, ArginfoAcosh, uint32(b.SizeOf("arginfo_acosh")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("expm1", ZifExpm1, ArginfoExpm1, uint32(b.SizeOf("arginfo_expm1")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("log1p", ZifLog1p, ArginfoLog1p, uint32(b.SizeOf("arginfo_log1p")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("pi", ZifPi, ArginfoPi, uint32(b.SizeOf("arginfo_pi")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_finite", ZifIsFinite, ArginfoIsFinite, uint32(b.SizeOf("arginfo_is_finite")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_nan", ZifIsNan, ArginfoIsNan, uint32(b.SizeOf("arginfo_is_nan")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_infinite", ZifIsInfinite, ArginfoIsInfinite, uint32(b.SizeOf("arginfo_is_infinite")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("pow", ZifPow, ArginfoPow, uint32(b.SizeOf("arginfo_pow")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("exp", ZifExp, ArginfoExp, uint32(b.SizeOf("arginfo_exp")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("log", ZifLog, ArginfoLog, uint32(b.SizeOf("arginfo_log")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("log10", ZifLog10, ArginfoLog10, uint32(b.SizeOf("arginfo_log10")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("sqrt", ZifSqrt, ArginfoSqrt, uint32(b.SizeOf("arginfo_sqrt")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("hypot", ZifHypot, ArginfoHypot, uint32(b.SizeOf("arginfo_hypot")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("deg2rad", ZifDeg2rad, ArginfoDeg2rad, uint32(b.SizeOf("arginfo_deg2rad")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("rad2deg", ZifRad2deg, ArginfoRad2deg, uint32(b.SizeOf("arginfo_rad2deg")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("bindec", ZifBindec, ArginfoBindec, uint32(b.SizeOf("arginfo_bindec")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("hexdec", ZifHexdec, ArginfoHexdec, uint32(b.SizeOf("arginfo_hexdec")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("octdec", ZifOctdec, ArginfoOctdec, uint32(b.SizeOf("arginfo_octdec")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("decbin", ZifDecbin, ArginfoDecbin, uint32(b.SizeOf("arginfo_decbin")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("decoct", ZifDecoct, ArginfoDecoct, uint32(b.SizeOf("arginfo_decoct")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("dechex", ZifDechex, ArginfoDechex, uint32(b.SizeOf("arginfo_dechex")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("base_convert", ZifBaseConvert, ArginfoBaseConvert, uint32(b.SizeOf("arginfo_base_convert")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("number_format", ZifNumberFormat, ArginfoNumberFormat, uint32(b.SizeOf("arginfo_number_format")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fmod", ZifFmod, ArginfoFmod, uint32(b.SizeOf("arginfo_fmod")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("intdiv", ZifIntdiv, ArginfoIntdiv, uint32(b.SizeOf("arginfo_intdiv")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("inet_ntop", ZifInetNtop, ArginfoInetNtop, uint32(b.SizeOf("arginfo_inet_ntop")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("inet_pton", PhpInetPton, ArginfoInetPton, uint32(b.SizeOf("arginfo_inet_pton")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ip2long", ZifIp2long, ArginfoIp2long, uint32(b.SizeOf("arginfo_ip2long")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("long2ip", ZifLong2ip, ArginfoLong2ip, uint32(b.SizeOf("arginfo_long2ip")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getenv", ZifGetenv, ArginfoGetenv, uint32(b.SizeOf("arginfo_getenv")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("putenv", ZifPutenv, ArginfoPutenv, uint32(b.SizeOf("arginfo_putenv")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getopt", ZifGetopt, ArginfoGetopt, uint32(b.SizeOf("arginfo_getopt")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("sys_getloadavg", ZifSysGetloadavg, ArginfoSysGetloadavg, uint32(b.SizeOf("arginfo_sys_getloadavg")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("microtime", ZifMicrotime, ArginfoMicrotime, uint32(b.SizeOf("arginfo_microtime")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("gettimeofday", ZifGettimeofday, ArginfoGettimeofday, uint32(b.SizeOf("arginfo_gettimeofday")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getrusage", ZifGetrusage, ArginfoGetrusage, uint32(b.SizeOf("arginfo_getrusage")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("hrtime", ZifHrtime, ArginfoHrtime, uint32(b.SizeOf("arginfo_hrtime")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("uniqid", ZifUniqid, ArginfoUniqid, uint32(b.SizeOf("arginfo_uniqid")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("quoted_printable_decode", ZifQuotedPrintableDecode, ArginfoQuotedPrintableDecode, uint32(b.SizeOf("arginfo_quoted_printable_decode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("quoted_printable_encode", ZifQuotedPrintableEncode, ArginfoQuotedPrintableEncode, uint32(b.SizeOf("arginfo_quoted_printable_encode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("convert_cyr_string", ZifConvertCyrString, ArginfoConvertCyrString, uint32(b.SizeOf("arginfo_convert_cyr_string")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_DEPRECATED),
	zend.MakeZendFunctionEntry("get_current_user", ZifGetCurrentUser, ArginfoGetCurrentUser, uint32(b.SizeOf("arginfo_get_current_user")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("set_time_limit", ZifSetTimeLimit, ArginfoSetTimeLimit, uint32(b.SizeOf("arginfo_set_time_limit")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("header_register_callback", ZifHeaderRegisterCallback, ArginfoHeaderRegisterCallback, uint32(b.SizeOf("arginfo_header_register_callback")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("get_cfg_var", ZifGetCfgVar, ArginfoGetCfgVar, uint32(b.SizeOf("arginfo_get_cfg_var")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("get_magic_quotes_gpc", ZifGetMagicQuotesGpc, ArginfoGetMagicQuotesGpc, uint32(b.SizeOf("arginfo_get_magic_quotes_gpc")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_DEPRECATED),
	zend.MakeZendFunctionEntry("get_magic_quotes_runtime", ZifGetMagicQuotesRuntime, ArginfoGetMagicQuotesRuntime, uint32(b.SizeOf("arginfo_get_magic_quotes_runtime")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_DEPRECATED),
	zend.MakeZendFunctionEntry("error_log", ZifErrorLog, ArginfoErrorLog, uint32(b.SizeOf("arginfo_error_log")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("error_get_last", ZifErrorGetLast, ArginfoErrorGetLast, uint32(b.SizeOf("arginfo_error_get_last")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("error_clear_last", ZifErrorClearLast, ArginfoErrorClearLast, uint32(b.SizeOf("arginfo_error_clear_last")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("call_user_func", ZifCallUserFunc, ArginfoCallUserFunc, uint32(b.SizeOf("arginfo_call_user_func")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("call_user_func_array", ZifCallUserFuncArray, ArginfoCallUserFuncArray, uint32(b.SizeOf("arginfo_call_user_func_array")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("forward_static_call", ZifForwardStaticCall, ArginfoForwardStaticCall, uint32(b.SizeOf("arginfo_forward_static_call")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("forward_static_call_array", ZifForwardStaticCallArray, ArginfoForwardStaticCallArray, uint32(b.SizeOf("arginfo_forward_static_call_array")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("serialize", ZifSerialize, ArginfoSerialize, uint32(b.SizeOf("arginfo_serialize")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("unserialize", ZifUnserialize, ArginfoUnserialize, uint32(b.SizeOf("arginfo_unserialize")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("var_dump", ZifVarDump, ArginfoVarDump, uint32(b.SizeOf("arginfo_var_dump")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("var_export", ZifVarExport, ArginfoVarExport, uint32(b.SizeOf("arginfo_var_export")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("debug_zval_dump", ZifDebugZvalDump, ArginfoDebugZvalDump, uint32(b.SizeOf("arginfo_debug_zval_dump")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("print_r", ZifPrintR, ArginfoPrintR, uint32(b.SizeOf("arginfo_print_r")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("memory_get_usage", ZifMemoryGetUsage, ArginfoMemoryGetUsage, uint32(b.SizeOf("arginfo_memory_get_usage")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("memory_get_peak_usage", ZifMemoryGetPeakUsage, ArginfoMemoryGetPeakUsage, uint32(b.SizeOf("arginfo_memory_get_peak_usage")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("register_shutdown_function", ZifRegisterShutdownFunction, ArginfoRegisterShutdownFunction, uint32(b.SizeOf("arginfo_register_shutdown_function")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("register_tick_function", ZifRegisterTickFunction, ArginfoRegisterTickFunction, uint32(b.SizeOf("arginfo_register_tick_function")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("unregister_tick_function", ZifUnregisterTickFunction, ArginfoUnregisterTickFunction, uint32(b.SizeOf("arginfo_unregister_tick_function")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("highlight_file", ZifHighlightFile, ArginfoHighlightFile, uint32(b.SizeOf("arginfo_highlight_file")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("show_source", ZifHighlightFile, ArginfoHighlightFile, uint32(b.SizeOf("arginfo_highlight_file")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("highlight_string", ZifHighlightString, ArginfoHighlightString, uint32(b.SizeOf("arginfo_highlight_string")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("php_strip_whitespace", ZifPhpStripWhitespace, ArginfoPhpStripWhitespace, uint32(b.SizeOf("arginfo_php_strip_whitespace")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ini_get", ZifIniGet, ArginfoIniGet, uint32(b.SizeOf("arginfo_ini_get")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ini_get_all", ZifIniGetAll, ArginfoIniGetAll, uint32(b.SizeOf("arginfo_ini_get_all")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ini_set", ZifIniSet, ArginfoIniSet, uint32(b.SizeOf("arginfo_ini_set")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ini_alter", ZifIniSet, ArginfoIniSet, uint32(b.SizeOf("arginfo_ini_set")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ini_restore", ZifIniRestore, ArginfoIniRestore, uint32(b.SizeOf("arginfo_ini_restore")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("get_include_path", ZifGetIncludePath, ArginfoGetIncludePath, uint32(b.SizeOf("arginfo_get_include_path")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("set_include_path", ZifSetIncludePath, ArginfoSetIncludePath, uint32(b.SizeOf("arginfo_set_include_path")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("restore_include_path", ZifRestoreIncludePath, ArginfoRestoreIncludePath, uint32(b.SizeOf("arginfo_restore_include_path")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_DEPRECATED),
	zend.MakeZendFunctionEntry("setcookie", ZifSetcookie, ArginfoSetcookie, uint32(b.SizeOf("arginfo_setcookie")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("setrawcookie", ZifSetrawcookie, ArginfoSetrawcookie, uint32(b.SizeOf("arginfo_setrawcookie")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("header", ZifHeader, ArginfoHeader, uint32(b.SizeOf("arginfo_header")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("header_remove", ZifHeaderRemove, ArginfoHeaderRemove, uint32(b.SizeOf("arginfo_header_remove")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("headers_sent", ZifHeadersSent, ArginfoHeadersSent, uint32(b.SizeOf("arginfo_headers_sent")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("headers_list", ZifHeadersList, ArginfoHeadersList, uint32(b.SizeOf("arginfo_headers_list")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("http_response_code", ZifHttpResponseCode, ArginfoHttpResponseCode, uint32(b.SizeOf("arginfo_http_response_code")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("connection_aborted", ZifConnectionAborted, ArginfoConnectionAborted, uint32(b.SizeOf("arginfo_connection_aborted")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("connection_status", ZifConnectionStatus, ArginfoConnectionStatus, uint32(b.SizeOf("arginfo_connection_status")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ignore_user_abort", ZifIgnoreUserAbort, ArginfoIgnoreUserAbort, uint32(b.SizeOf("arginfo_ignore_user_abort")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("parse_ini_file", ZifParseIniFile, ArginfoParseIniFile, uint32(b.SizeOf("arginfo_parse_ini_file")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("parse_ini_string", ZifParseIniString, ArginfoParseIniString, uint32(b.SizeOf("arginfo_parse_ini_string")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_uploaded_file", ZifIsUploadedFile, ArginfoIsUploadedFile, uint32(b.SizeOf("arginfo_is_uploaded_file")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("move_uploaded_file", ZifMoveUploadedFile, ArginfoMoveUploadedFile, uint32(b.SizeOf("arginfo_move_uploaded_file")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("gethostbyaddr", ZifGethostbyaddr, ArginfoGethostbyaddr, uint32(b.SizeOf("arginfo_gethostbyaddr")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("gethostbyname", ZifGethostbyname, ArginfoGethostbyname, uint32(b.SizeOf("arginfo_gethostbyname")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("gethostbynamel", ZifGethostbynamel, ArginfoGethostbynamel, uint32(b.SizeOf("arginfo_gethostbynamel")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("gethostname", ZifGethostname, ArginfoGethostname, uint32(b.SizeOf("arginfo_gethostname")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("net_get_interfaces", ZifNetGetInterfaces, ArginfoNetGetInterfaces, uint32(b.SizeOf("arginfo_net_get_interfaces")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("dns_check_record", ZifDnsCheckRecord, ArginfoDnsCheckRecord, uint32(b.SizeOf("arginfo_dns_check_record")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("checkdnsrr", ZifDnsCheckRecord, ArginfoDnsCheckRecord, uint32(b.SizeOf("arginfo_dns_check_record")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("dns_get_mx", ZifDnsGetMx, ArginfoDnsGetMx, uint32(b.SizeOf("arginfo_dns_get_mx")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getmxrr", ZifDnsGetMx, ArginfoDnsGetMx, uint32(b.SizeOf("arginfo_dns_get_mx")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("dns_get_record", ZifDnsGetRecord, ArginfoDnsGetRecord, uint32(b.SizeOf("arginfo_dns_get_record")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("intval", ZifIntval, ArginfoIntval, uint32(b.SizeOf("arginfo_intval")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("floatval", ZifFloatval, ArginfoFloatval, uint32(b.SizeOf("arginfo_floatval")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("doubleval", ZifFloatval, ArginfoFloatval, uint32(b.SizeOf("arginfo_floatval")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("strval", ZifStrval, ArginfoStrval, uint32(b.SizeOf("arginfo_strval")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("boolval", ZifBoolval, ArginfoBoolval, uint32(b.SizeOf("arginfo_boolval")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("gettype", ZifGettype, ArginfoGettype, uint32(b.SizeOf("arginfo_gettype")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("settype", ZifSettype, ArginfoSettype, uint32(b.SizeOf("arginfo_settype")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_null", ZifIsNull, ArginfoIsNull, uint32(b.SizeOf("arginfo_is_null")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_resource", ZifIsResource, ArginfoIsResource, uint32(b.SizeOf("arginfo_is_resource")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_bool", ZifIsBool, ArginfoIsBool, uint32(b.SizeOf("arginfo_is_bool")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_int", ZifIsInt, ArginfoIsInt, uint32(b.SizeOf("arginfo_is_int")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_float", ZifIsFloat, ArginfoIsFloat, uint32(b.SizeOf("arginfo_is_float")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_integer", ZifIsInt, ArginfoIsInt, uint32(b.SizeOf("arginfo_is_int")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_long", ZifIsInt, ArginfoIsInt, uint32(b.SizeOf("arginfo_is_int")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_double", ZifIsFloat, ArginfoIsFloat, uint32(b.SizeOf("arginfo_is_float")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_real", ZifIsFloat, ArginfoIsFloat, uint32(b.SizeOf("arginfo_is_float")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_DEPRECATED),
	zend.MakeZendFunctionEntry("is_numeric", ZifIsNumeric, ArginfoIsNumeric, uint32(b.SizeOf("arginfo_is_numeric")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_string", ZifIsString, ArginfoIsString, uint32(b.SizeOf("arginfo_is_string")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_array", ZifIsArray, ArginfoIsArray, uint32(b.SizeOf("arginfo_is_array")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_object", ZifIsObject, ArginfoIsObject, uint32(b.SizeOf("arginfo_is_object")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_scalar", ZifIsScalar, ArginfoIsScalar, uint32(b.SizeOf("arginfo_is_scalar")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_callable", ZifIsCallable, ArginfoIsCallable, uint32(b.SizeOf("arginfo_is_callable")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_iterable", ZifIsIterable, ArginfoIsIterable, uint32(b.SizeOf("arginfo_is_iterable")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_countable", ZifIsCountable, ArginfoIsCountable, uint32(b.SizeOf("arginfo_is_countable")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("pclose", ZifPclose, ArginfoPclose, uint32(b.SizeOf("arginfo_pclose")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("popen", ZifPopen, ArginfoPopen, uint32(b.SizeOf("arginfo_popen")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("readfile", ZifReadfile, ArginfoReadfile, uint32(b.SizeOf("arginfo_readfile")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("rewind", ZifRewind, ArginfoRewind, uint32(b.SizeOf("arginfo_rewind")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("rmdir", ZifRmdir, ArginfoRmdir, uint32(b.SizeOf("arginfo_rmdir")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("umask", ZifUmask, ArginfoUmask, uint32(b.SizeOf("arginfo_umask")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fclose", ZifFclose, ArginfoFclose, uint32(b.SizeOf("arginfo_fclose")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("feof", ZifFeof, ArginfoFeof, uint32(b.SizeOf("arginfo_feof")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fgetc", ZifFgetc, ArginfoFgetc, uint32(b.SizeOf("arginfo_fgetc")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fgets", ZifFgets, ArginfoFgets, uint32(b.SizeOf("arginfo_fgets")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fgetss", ZifFgetss, ArginfoFgetss, uint32(b.SizeOf("arginfo_fgetss")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_DEPRECATED),
	zend.MakeZendFunctionEntry("fread", ZifFread, ArginfoFread, uint32(b.SizeOf("arginfo_fread")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fopen", PhpIfFopen, ArginfoFopen, uint32(b.SizeOf("arginfo_fopen")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fpassthru", ZifFpassthru, ArginfoFpassthru, uint32(b.SizeOf("arginfo_fpassthru")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ftruncate", PhpIfFtruncate, ArginfoFtruncate, uint32(b.SizeOf("arginfo_ftruncate")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fstat", PhpIfFstat, ArginfoFstat, uint32(b.SizeOf("arginfo_fstat")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fseek", ZifFseek, ArginfoFseek, uint32(b.SizeOf("arginfo_fseek")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ftell", ZifFtell, ArginfoFtell, uint32(b.SizeOf("arginfo_ftell")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fflush", ZifFflush, ArginfoFflush, uint32(b.SizeOf("arginfo_fflush")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fwrite", ZifFwrite, ArginfoFwrite, uint32(b.SizeOf("arginfo_fwrite")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fputs", ZifFwrite, ArginfoFwrite, uint32(b.SizeOf("arginfo_fwrite")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("mkdir", ZifMkdir, ArginfoMkdir, uint32(b.SizeOf("arginfo_mkdir")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("rename", ZifRename, ArginfoRename, uint32(b.SizeOf("arginfo_rename")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("copy", ZifCopy, ArginfoCopy, uint32(b.SizeOf("arginfo_copy")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("tempnam", ZifTempnam, ArginfoTempnam, uint32(b.SizeOf("arginfo_tempnam")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("tmpfile", PhpIfTmpfile, ArginfoTmpfile, uint32(b.SizeOf("arginfo_tmpfile")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("file", ZifFile, ArginfoFile, uint32(b.SizeOf("arginfo_file")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("file_get_contents", ZifFileGetContents, ArginfoFileGetContents, uint32(b.SizeOf("arginfo_file_get_contents")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("file_put_contents", ZifFilePutContents, ArginfoFilePutContents, uint32(b.SizeOf("arginfo_file_put_contents")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_select", ZifStreamSelect, ArginfoStreamSelect, uint32(b.SizeOf("arginfo_stream_select")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_context_create", ZifStreamContextCreate, ArginfoStreamContextCreate, uint32(b.SizeOf("arginfo_stream_context_create")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_context_set_params", ZifStreamContextSetParams, ArginfoStreamContextSetParams, uint32(b.SizeOf("arginfo_stream_context_set_params")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_context_get_params", ZifStreamContextGetParams, ArginfoStreamContextGetParams, uint32(b.SizeOf("arginfo_stream_context_get_params")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_context_set_option", ZifStreamContextSetOption, ArginfoStreamContextSetOption, uint32(b.SizeOf("arginfo_stream_context_set_option")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_context_get_options", ZifStreamContextGetOptions, ArginfoStreamContextGetOptions, uint32(b.SizeOf("arginfo_stream_context_get_options")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_context_get_default", ZifStreamContextGetDefault, ArginfoStreamContextGetDefault, uint32(b.SizeOf("arginfo_stream_context_get_default")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_context_set_default", ZifStreamContextSetDefault, ArginfoStreamContextSetDefault, uint32(b.SizeOf("arginfo_stream_context_set_default")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_filter_prepend", ZifStreamFilterPrepend, ArginfoStreamFilterPrepend, uint32(b.SizeOf("arginfo_stream_filter_prepend")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_filter_append", ZifStreamFilterAppend, ArginfoStreamFilterAppend, uint32(b.SizeOf("arginfo_stream_filter_append")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_filter_remove", ZifStreamFilterRemove, ArginfoStreamFilterRemove, uint32(b.SizeOf("arginfo_stream_filter_remove")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_socket_client", ZifStreamSocketClient, ArginfoStreamSocketClient, uint32(b.SizeOf("arginfo_stream_socket_client")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_socket_server", ZifStreamSocketServer, ArginfoStreamSocketServer, uint32(b.SizeOf("arginfo_stream_socket_server")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_socket_accept", ZifStreamSocketAccept, ArginfoStreamSocketAccept, uint32(b.SizeOf("arginfo_stream_socket_accept")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_socket_get_name", ZifStreamSocketGetName, ArginfoStreamSocketGetName, uint32(b.SizeOf("arginfo_stream_socket_get_name")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_socket_recvfrom", ZifStreamSocketRecvfrom, ArginfoStreamSocketRecvfrom, uint32(b.SizeOf("arginfo_stream_socket_recvfrom")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_socket_sendto", ZifStreamSocketSendto, ArginfoStreamSocketSendto, uint32(b.SizeOf("arginfo_stream_socket_sendto")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_socket_enable_crypto", ZifStreamSocketEnableCrypto, ArginfoStreamSocketEnableCrypto, uint32(b.SizeOf("arginfo_stream_socket_enable_crypto")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_socket_shutdown", ZifStreamSocketShutdown, ArginfoStreamSocketShutdown, uint32(b.SizeOf("arginfo_stream_socket_shutdown")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_socket_pair", ZifStreamSocketPair, ArginfoStreamSocketPair, uint32(b.SizeOf("arginfo_stream_socket_pair")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_copy_to_stream", ZifStreamCopyToStream, ArginfoStreamCopyToStream, uint32(b.SizeOf("arginfo_stream_copy_to_stream")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_get_contents", ZifStreamGetContents, ArginfoStreamGetContents, uint32(b.SizeOf("arginfo_stream_get_contents")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_supports_lock", ZifStreamSupportsLock, ArginfoStreamSupportsLock, uint32(b.SizeOf("arginfo_stream_supports_lock")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_isatty", ZifStreamIsatty, ArginfoStreamIsatty, uint32(b.SizeOf("arginfo_stream_isatty")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fgetcsv", ZifFgetcsv, ArginfoFgetcsv, uint32(b.SizeOf("arginfo_fgetcsv")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fputcsv", ZifFputcsv, ArginfoFputcsv, uint32(b.SizeOf("arginfo_fputcsv")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("flock", ZifFlock, ArginfoFlock, uint32(b.SizeOf("arginfo_flock")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("get_meta_tags", ZifGetMetaTags, ArginfoGetMetaTags, uint32(b.SizeOf("arginfo_get_meta_tags")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_set_read_buffer", ZifStreamSetReadBuffer, ArginfoStreamSetReadBuffer, uint32(b.SizeOf("arginfo_stream_set_read_buffer")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_set_write_buffer", ZifStreamSetWriteBuffer, ArginfoStreamSetWriteBuffer, uint32(b.SizeOf("arginfo_stream_set_write_buffer")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("set_file_buffer", ZifStreamSetWriteBuffer, ArginfoStreamSetWriteBuffer, uint32(b.SizeOf("arginfo_stream_set_write_buffer")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_set_chunk_size", ZifStreamSetChunkSize, ArginfoStreamSetChunkSize, uint32(b.SizeOf("arginfo_stream_set_chunk_size")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_set_blocking", ZifStreamSetBlocking, ArginfoStreamSetBlocking, uint32(b.SizeOf("arginfo_stream_set_blocking")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("socket_set_blocking", ZifStreamSetBlocking, ArginfoStreamSetBlocking, uint32(b.SizeOf("arginfo_stream_set_blocking")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_get_meta_data", ZifStreamGetMetaData, ArginfoStreamGetMetaData, uint32(b.SizeOf("arginfo_stream_get_meta_data")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_get_line", ZifStreamGetLine, ArginfoStreamGetLine, uint32(b.SizeOf("arginfo_stream_get_line")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_wrapper_register", ZifStreamWrapperRegister, ArginfoStreamWrapperRegister, uint32(b.SizeOf("arginfo_stream_wrapper_register")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_register_wrapper", ZifStreamWrapperRegister, ArginfoStreamWrapperRegister, uint32(b.SizeOf("arginfo_stream_wrapper_register")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_wrapper_unregister", ZifStreamWrapperUnregister, ArginfoStreamWrapperUnregister, uint32(b.SizeOf("arginfo_stream_wrapper_unregister")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_wrapper_restore", ZifStreamWrapperRestore, ArginfoStreamWrapperRestore, uint32(b.SizeOf("arginfo_stream_wrapper_restore")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_get_wrappers", ZifStreamGetWrappers, ArginfoStreamGetWrappers, uint32(b.SizeOf("arginfo_stream_get_wrappers")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_get_transports", ZifStreamGetTransports, ArginfoStreamGetTransports, uint32(b.SizeOf("arginfo_stream_get_transports")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_resolve_include_path", ZifStreamResolveIncludePath, ArginfoStreamResolveIncludePath, uint32(b.SizeOf("arginfo_stream_resolve_include_path")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_is_local", ZifStreamIsLocal, ArginfoStreamIsLocal, uint32(b.SizeOf("arginfo_stream_is_local")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("get_headers", ZifGetHeaders, ArginfoGetHeaders, uint32(b.SizeOf("arginfo_get_headers")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_set_timeout", ZifStreamSetTimeout, ArginfoStreamSetTimeout, uint32(b.SizeOf("arginfo_stream_set_timeout")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("socket_set_timeout", ZifStreamSetTimeout, ArginfoStreamSetTimeout, uint32(b.SizeOf("arginfo_stream_set_timeout")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("socket_get_status", ZifStreamGetMetaData, ArginfoStreamGetMetaData, uint32(b.SizeOf("arginfo_stream_get_meta_data")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("realpath", ZifRealpath, ArginfoRealpath, uint32(b.SizeOf("arginfo_realpath")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fnmatch", ZifFnmatch, ArginfoFnmatch, uint32(b.SizeOf("arginfo_fnmatch")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fsockopen", ZifFsockopen, ArginfoFsockopen, uint32(b.SizeOf("arginfo_fsockopen")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("pfsockopen", ZifPfsockopen, ArginfoPfsockopen, uint32(b.SizeOf("arginfo_pfsockopen")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("pack", ZifPack, ArginfoPack, uint32(b.SizeOf("arginfo_pack")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("unpack", ZifUnpack, ArginfoUnpack, uint32(b.SizeOf("arginfo_unpack")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("get_browser", ZifGetBrowser, ArginfoGetBrowser, uint32(b.SizeOf("arginfo_get_browser")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("crypt", ZifCrypt, ArginfoCrypt, uint32(b.SizeOf("arginfo_crypt")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("opendir", ZifOpendir, ArginfoOpendir, uint32(b.SizeOf("arginfo_opendir")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("closedir", ZifClosedir, ArginfoClosedir, uint32(b.SizeOf("arginfo_closedir")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("chdir", ZifChdir, ArginfoChdir, uint32(b.SizeOf("arginfo_chdir")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("chroot", ZifChroot, ArginfoChroot, uint32(b.SizeOf("arginfo_chroot")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getcwd", ZifGetcwd, ArginfoGetcwd, uint32(b.SizeOf("arginfo_getcwd")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("rewinddir", ZifRewinddir, ArginfoRewinddir, uint32(b.SizeOf("arginfo_rewinddir")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("readdir", PhpIfReaddir, ArginfoReaddir, uint32(b.SizeOf("arginfo_readdir")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("dir", ZifGetdir, ArginfoDir, uint32(b.SizeOf("arginfo_dir")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("scandir", ZifScandir, ArginfoScandir, uint32(b.SizeOf("arginfo_scandir")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("glob", ZifGlob, ArginfoGlob, uint32(b.SizeOf("arginfo_glob")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fileatime", ZifFileatime, ArginfoFileatime, uint32(b.SizeOf("arginfo_fileatime")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("filectime", ZifFilectime, ArginfoFilectime, uint32(b.SizeOf("arginfo_filectime")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("filegroup", ZifFilegroup, ArginfoFilegroup, uint32(b.SizeOf("arginfo_filegroup")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fileinode", ZifFileinode, ArginfoFileinode, uint32(b.SizeOf("arginfo_fileinode")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("filemtime", ZifFilemtime, ArginfoFilemtime, uint32(b.SizeOf("arginfo_filemtime")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fileowner", ZifFileowner, ArginfoFileowner, uint32(b.SizeOf("arginfo_fileowner")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("fileperms", ZifFileperms, ArginfoFileperms, uint32(b.SizeOf("arginfo_fileperms")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("filesize", ZifFilesize, ArginfoFilesize, uint32(b.SizeOf("arginfo_filesize")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("filetype", ZifFiletype, ArginfoFiletype, uint32(b.SizeOf("arginfo_filetype")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("file_exists", ZifFileExists, ArginfoFileExists, uint32(b.SizeOf("arginfo_file_exists")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_writable", ZifIsWritable, ArginfoIsWritable, uint32(b.SizeOf("arginfo_is_writable")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_writeable", ZifIsWritable, ArginfoIsWritable, uint32(b.SizeOf("arginfo_is_writable")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_readable", ZifIsReadable, ArginfoIsReadable, uint32(b.SizeOf("arginfo_is_readable")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_executable", ZifIsExecutable, ArginfoIsExecutable, uint32(b.SizeOf("arginfo_is_executable")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_file", ZifIsFile, ArginfoIsFile, uint32(b.SizeOf("arginfo_is_file")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_dir", ZifIsDir, ArginfoIsDir, uint32(b.SizeOf("arginfo_is_dir")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("is_link", ZifIsLink, ArginfoIsLink, uint32(b.SizeOf("arginfo_is_link")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stat", PhpIfStat, ArginfoStat, uint32(b.SizeOf("arginfo_stat")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("lstat", PhpIfLstat, ArginfoLstat, uint32(b.SizeOf("arginfo_lstat")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("chown", ZifChown, ArginfoChown, uint32(b.SizeOf("arginfo_chown")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("chgrp", ZifChgrp, ArginfoChgrp, uint32(b.SizeOf("arginfo_chgrp")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("lchown", ZifLchown, ArginfoLchown, uint32(b.SizeOf("arginfo_lchown")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("lchgrp", ZifLchgrp, ArginfoLchgrp, uint32(b.SizeOf("arginfo_lchgrp")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("chmod", ZifChmod, ArginfoChmod, uint32(b.SizeOf("arginfo_chmod")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("touch", ZifTouch, ArginfoTouch, uint32(b.SizeOf("arginfo_touch")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("clearstatcache", ZifClearstatcache, ArginfoClearstatcache, uint32(b.SizeOf("arginfo_clearstatcache")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("disk_total_space", ZifDiskTotalSpace, ArginfoDiskTotalSpace, uint32(b.SizeOf("arginfo_disk_total_space")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("disk_free_space", ZifDiskFreeSpace, ArginfoDiskFreeSpace, uint32(b.SizeOf("arginfo_disk_free_space")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("diskfreespace", ZifDiskFreeSpace, ArginfoDiskFreeSpace, uint32(b.SizeOf("arginfo_disk_free_space")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("realpath_cache_size", ZifRealpathCacheSize, ArginfoRealpathCacheSize, uint32(b.SizeOf("arginfo_realpath_cache_size")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("realpath_cache_get", ZifRealpathCacheGet, ArginfoRealpathCacheGet, uint32(b.SizeOf("arginfo_realpath_cache_get")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("mail", ZifMail, ArginfoMail, uint32(b.SizeOf("arginfo_mail")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ezmlm_hash", ZifEzmlmHash, ArginfoEzmlmHash, uint32(b.SizeOf("arginfo_ezmlm_hash")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_DEPRECATED),
	zend.MakeZendFunctionEntry("openlog", ZifOpenlog, ArginfoOpenlog, uint32(b.SizeOf("arginfo_openlog")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("syslog", ZifSyslog, ArginfoSyslog, uint32(b.SizeOf("arginfo_syslog")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("closelog", ZifCloselog, ArginfoCloselog, uint32(b.SizeOf("arginfo_closelog")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("lcg_value", ZifLcgValue, ArginfoLcgValue, uint32(b.SizeOf("arginfo_lcg_value")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("metaphone", ZifMetaphone, ArginfoMetaphone, uint32(b.SizeOf("arginfo_metaphone")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ob_start", core.ZifObStart, ArginfoObStart, uint32(b.SizeOf("arginfo_ob_start")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ob_flush", core.ZifObFlush, ArginfoObFlush, uint32(b.SizeOf("arginfo_ob_flush")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ob_clean", core.ZifObClean, ArginfoObClean, uint32(b.SizeOf("arginfo_ob_clean")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ob_end_flush", core.ZifObEndFlush, ArginfoObEndFlush, uint32(b.SizeOf("arginfo_ob_end_flush")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ob_end_clean", core.ZifObEndClean, ArginfoObEndClean, uint32(b.SizeOf("arginfo_ob_end_clean")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ob_get_flush", core.ZifObGetFlush, ArginfoObGetFlush, uint32(b.SizeOf("arginfo_ob_get_flush")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ob_get_clean", core.ZifObGetClean, ArginfoObGetClean, uint32(b.SizeOf("arginfo_ob_get_clean")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ob_get_length", core.ZifObGetLength, ArginfoObGetLength, uint32(b.SizeOf("arginfo_ob_get_length")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ob_get_level", core.ZifObGetLevel, ArginfoObGetLevel, uint32(b.SizeOf("arginfo_ob_get_level")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ob_get_status", core.ZifObGetStatus, ArginfoObGetStatus, uint32(b.SizeOf("arginfo_ob_get_status")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ob_get_contents", core.ZifObGetContents, ArginfoObGetContents, uint32(b.SizeOf("arginfo_ob_get_contents")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ob_implicit_flush", core.ZifObImplicitFlush, ArginfoObImplicitFlush, uint32(b.SizeOf("arginfo_ob_implicit_flush")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ob_list_handlers", core.ZifObListHandlers, ArginfoObListHandlers, uint32(b.SizeOf("arginfo_ob_list_handlers")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ksort", ZifKsort, ArginfoKsort, uint32(b.SizeOf("arginfo_ksort")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("krsort", ZifKrsort, ArginfoKrsort, uint32(b.SizeOf("arginfo_krsort")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("natsort", ZifNatsort, ArginfoNatsort, uint32(b.SizeOf("arginfo_natsort")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("natcasesort", ZifNatcasesort, ArginfoNatcasesort, uint32(b.SizeOf("arginfo_natcasesort")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("asort", ZifAsort, ArginfoAsort, uint32(b.SizeOf("arginfo_asort")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("arsort", ZifArsort, ArginfoArsort, uint32(b.SizeOf("arginfo_arsort")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("sort", ZifSort, ArginfoSort, uint32(b.SizeOf("arginfo_sort")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("rsort", ZifRsort, ArginfoRsort, uint32(b.SizeOf("arginfo_rsort")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("usort", ZifUsort, ArginfoUsort, uint32(b.SizeOf("arginfo_usort")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("uasort", ZifUasort, ArginfoUasort, uint32(b.SizeOf("arginfo_uasort")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("uksort", ZifUksort, ArginfoUksort, uint32(b.SizeOf("arginfo_uksort")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("shuffle", ZifShuffle, ArginfoShuffle, uint32(b.SizeOf("arginfo_shuffle")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_walk", ZifArrayWalk, ArginfoArrayWalk, uint32(b.SizeOf("arginfo_array_walk")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_walk_recursive", ZifArrayWalkRecursive, ArginfoArrayWalkRecursive, uint32(b.SizeOf("arginfo_array_walk_recursive")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("count", ZifCount, ArginfoCount, uint32(b.SizeOf("arginfo_count")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("end", ZifEnd, ArginfoEnd, uint32(b.SizeOf("arginfo_end")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("prev", ZifPrev, ArginfoPrev, uint32(b.SizeOf("arginfo_prev")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("next", ZifNext, ArginfoNext, uint32(b.SizeOf("arginfo_next")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("reset", ZifReset, ArginfoReset, uint32(b.SizeOf("arginfo_reset")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("current", ZifCurrent, ArginfoCurrent, uint32(b.SizeOf("arginfo_current")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("key", ZifKey, ArginfoKey, uint32(b.SizeOf("arginfo_key")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("min", ZifMin, ArginfoMin, uint32(b.SizeOf("arginfo_min")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("max", ZifMax, ArginfoMax, uint32(b.SizeOf("arginfo_max")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("in_array", ZifInArray, ArginfoInArray, uint32(b.SizeOf("arginfo_in_array")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_search", ZifArraySearch, ArginfoArraySearch, uint32(b.SizeOf("arginfo_array_search")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("extract", ZifExtract, ArginfoExtract, uint32(b.SizeOf("arginfo_extract")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("compact", ZifCompact, ArginfoCompact, uint32(b.SizeOf("arginfo_compact")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_fill", ZifArrayFill, ArginfoArrayFill, uint32(b.SizeOf("arginfo_array_fill")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_fill_keys", ZifArrayFillKeys, ArginfoArrayFillKeys, uint32(b.SizeOf("arginfo_array_fill_keys")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("range", ZifRange, ArginfoRange, uint32(b.SizeOf("arginfo_range")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_multisort", ZifArrayMultisort, ArginfoArrayMultisort, uint32(b.SizeOf("arginfo_array_multisort")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_push", ZifArrayPush, ArginfoArrayPush, uint32(b.SizeOf("arginfo_array_push")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_pop", ZifArrayPop, ArginfoArrayPop, uint32(b.SizeOf("arginfo_array_pop")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_shift", ZifArrayShift, ArginfoArrayShift, uint32(b.SizeOf("arginfo_array_shift")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_unshift", ZifArrayUnshift, ArginfoArrayUnshift, uint32(b.SizeOf("arginfo_array_unshift")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_splice", ZifArraySplice, ArginfoArraySplice, uint32(b.SizeOf("arginfo_array_splice")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_slice", ZifArraySlice, ArginfoArraySlice, uint32(b.SizeOf("arginfo_array_slice")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_merge", ZifArrayMerge, ArginfoArrayMerge, uint32(b.SizeOf("arginfo_array_merge")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_merge_recursive", ZifArrayMergeRecursive, ArginfoArrayMergeRecursive, uint32(b.SizeOf("arginfo_array_merge_recursive")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_replace", ZifArrayReplace, ArginfoArrayReplace, uint32(b.SizeOf("arginfo_array_replace")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_replace_recursive", ZifArrayReplaceRecursive, ArginfoArrayReplaceRecursive, uint32(b.SizeOf("arginfo_array_replace_recursive")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_keys", ZifArrayKeys, ArginfoArrayKeys, uint32(b.SizeOf("arginfo_array_keys")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_key_first", ZifArrayKeyFirst, ArginfoArrayKeyFirst, uint32(b.SizeOf("arginfo_array_key_first")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_key_last", ZifArrayKeyLast, ArginfoArrayKeyLast, uint32(b.SizeOf("arginfo_array_key_last")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_values", ZifArrayValues, ArginfoArrayValues, uint32(b.SizeOf("arginfo_array_values")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_count_values", ZifArrayCountValues, ArginfoArrayCountValues, uint32(b.SizeOf("arginfo_array_count_values")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_column", ZifArrayColumn, ArginfoArrayColumn, uint32(b.SizeOf("arginfo_array_column")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_reverse", ZifArrayReverse, ArginfoArrayReverse, uint32(b.SizeOf("arginfo_array_reverse")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_reduce", ZifArrayReduce, ArginfoArrayReduce, uint32(b.SizeOf("arginfo_array_reduce")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_pad", ZifArrayPad, ArginfoArrayPad, uint32(b.SizeOf("arginfo_array_pad")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_flip", ZifArrayFlip, ArginfoArrayFlip, uint32(b.SizeOf("arginfo_array_flip")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_change_key_case", ZifArrayChangeKeyCase, ArginfoArrayChangeKeyCase, uint32(b.SizeOf("arginfo_array_change_key_case")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_rand", ZifArrayRand, ArginfoArrayRand, uint32(b.SizeOf("arginfo_array_rand")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_unique", ZifArrayUnique, ArginfoArrayUnique, uint32(b.SizeOf("arginfo_array_unique")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_intersect", ZifArrayIntersect, ArginfoArrayIntersect, uint32(b.SizeOf("arginfo_array_intersect")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_intersect_key", ZifArrayIntersectKey, ArginfoArrayIntersectKey, uint32(b.SizeOf("arginfo_array_intersect_key")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_intersect_ukey", ZifArrayIntersectUkey, ArginfoArrayIntersectUkey, uint32(b.SizeOf("arginfo_array_intersect_ukey")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_uintersect", ZifArrayUintersect, ArginfoArrayUintersect, uint32(b.SizeOf("arginfo_array_uintersect")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_intersect_assoc", ZifArrayIntersectAssoc, ArginfoArrayIntersectAssoc, uint32(b.SizeOf("arginfo_array_intersect_assoc")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_uintersect_assoc", ZifArrayUintersectAssoc, ArginfoArrayUintersectAssoc, uint32(b.SizeOf("arginfo_array_uintersect_assoc")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_intersect_uassoc", ZifArrayIntersectUassoc, ArginfoArrayIntersectUassoc, uint32(b.SizeOf("arginfo_array_intersect_uassoc")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_uintersect_uassoc", ZifArrayUintersectUassoc, ArginfoArrayUintersectUassoc, uint32(b.SizeOf("arginfo_array_uintersect_uassoc")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_diff", ZifArrayDiff, ArginfoArrayDiff, uint32(b.SizeOf("arginfo_array_diff")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_diff_key", ZifArrayDiffKey, ArginfoArrayDiffKey, uint32(b.SizeOf("arginfo_array_diff_key")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_diff_ukey", ZifArrayDiffUkey, ArginfoArrayDiffUkey, uint32(b.SizeOf("arginfo_array_diff_ukey")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_udiff", ZifArrayUdiff, ArginfoArrayUdiff, uint32(b.SizeOf("arginfo_array_udiff")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_diff_assoc", ZifArrayDiffAssoc, ArginfoArrayDiffAssoc, uint32(b.SizeOf("arginfo_array_diff_assoc")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_udiff_assoc", ZifArrayUdiffAssoc, ArginfoArrayUdiffAssoc, uint32(b.SizeOf("arginfo_array_udiff_assoc")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_diff_uassoc", ZifArrayDiffUassoc, ArginfoArrayDiffUassoc, uint32(b.SizeOf("arginfo_array_diff_uassoc")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_udiff_uassoc", ZifArrayUdiffUassoc, ArginfoArrayUdiffUassoc, uint32(b.SizeOf("arginfo_array_udiff_uassoc")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_sum", ZifArraySum, ArginfoArraySum, uint32(b.SizeOf("arginfo_array_sum")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_product", ZifArrayProduct, ArginfoArrayProduct, uint32(b.SizeOf("arginfo_array_product")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_filter", ZifArrayFilter, ArginfoArrayFilter, uint32(b.SizeOf("arginfo_array_filter")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_map", ZifArrayMap, ArginfoArrayMap, uint32(b.SizeOf("arginfo_array_map")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_chunk", ZifArrayChunk, ArginfoArrayChunk, uint32(b.SizeOf("arginfo_array_chunk")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_combine", ZifArrayCombine, ArginfoArrayCombine, uint32(b.SizeOf("arginfo_array_combine")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("array_key_exists", ZifArrayKeyExists, ArginfoArrayKeyExists, uint32(b.SizeOf("arginfo_array_key_exists")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("pos", ZifCurrent, ArginfoCurrent, uint32(b.SizeOf("arginfo_current")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("sizeof", ZifCount, ArginfoCount, uint32(b.SizeOf("arginfo_count")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("key_exists", ZifArrayKeyExists, ArginfoArrayKeyExists, uint32(b.SizeOf("arginfo_array_key_exists")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("assert", ZifAssert, ArginfoAssert, uint32(b.SizeOf("arginfo_assert")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("assert_options", ZifAssertOptions, ArginfoAssertOptions, uint32(b.SizeOf("arginfo_assert_options")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("version_compare", ZifVersionCompare, ArginfoVersionCompare, uint32(b.SizeOf("arginfo_version_compare")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("ftok", ZifFtok, ArginfoFtok, uint32(b.SizeOf("arginfo_ftok")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("str_rot13", ZifStrRot13, ArginfoStrRot13, uint32(b.SizeOf("arginfo_str_rot13")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_get_filters", ZifStreamGetFilters, ArginfoStreamGetFilters, uint32(b.SizeOf("arginfo_stream_get_filters")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_filter_register", ZifStreamFilterRegister, ArginfoStreamFilterRegister, uint32(b.SizeOf("arginfo_stream_filter_register")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_bucket_make_writeable", ZifStreamBucketMakeWriteable, ArginfoStreamBucketMakeWriteable, uint32(b.SizeOf("arginfo_stream_bucket_make_writeable")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_bucket_prepend", ZifStreamBucketPrepend, ArginfoStreamBucketPrepend, uint32(b.SizeOf("arginfo_stream_bucket_prepend")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_bucket_append", ZifStreamBucketAppend, ArginfoStreamBucketAppend, uint32(b.SizeOf("arginfo_stream_bucket_append")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("stream_bucket_new", ZifStreamBucketNew, ArginfoStreamBucketNew, uint32(b.SizeOf("arginfo_stream_bucket_new")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("output_add_rewrite_var", core.ZifOutputAddRewriteVar, ArginfoOutputAddRewriteVar, uint32(b.SizeOf("arginfo_output_add_rewrite_var")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("output_reset_rewrite_vars", core.ZifOutputResetRewriteVars, ArginfoOutputResetRewriteVars, uint32(b.SizeOf("arginfo_output_reset_rewrite_vars")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("sys_get_temp_dir", ZifSysGetTempDir, ArginfoSysGetTempDir, uint32(b.SizeOf("arginfo_sys_get_temp_dir")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var StandardDeps []zend.ZendModuleDep = []zend.ZendModuleDep{
	zend.MakeZendModuleDep("session", nil, nil, zend.MODULE_DEP_OPTIONAL),
	zend.MakeZendModuleDep(nil, nil, nil, 0),
}
var BasicFunctionsModule zend.ZendModuleEntry = zend.MakeZendModuleEntry(b.SizeOf("zend_module_entry"), zend.ZEND_MODULE_API_NO, 0, zend.USING_ZTS, nil, StandardDeps, "standard", BasicFunctions, ZmStartupBasic, ZmShutdownBasic, ZmActivateBasic, ZmDeactivateBasic, ZmInfoBasic, PHP_STANDARD_VERSION, 0, nil, nil, nil, nil, 0, 0, nil, 0, "API"+"ZEND_MODULE_API_NO"+zend.ZEND_BUILD_TS)

/* {{{ proto string inet_ntop(string in_addr)
   Converts a packed inet address to a human readable IP address string */

/* {{{ proto string inet_pton(string ip_address)
   Converts a human readable IP address to a packed binary string */

/* {{{ proto int ip2long(string ip_address)
   Converts a string containing an (IPv4) Internet Protocol dotted address into a proper address */

/* {{{ proto bool putenv(string setting)
   Set the value of an environment variable */

/* {{{ free_argv()
   Free the memory allocated to an argv array. */

/* {{{ proto mixed time_nanosleep(int seconds, int nanoseconds)
   Delay for a number of seconds and nano seconds */

/* {{{ proto string get_current_user(void)
   Get the name of the owner of the current PHP script */

/* {{{ add_config_entry
 */

/* {{{ proto int getservbyname(string service, string protocol)
   Returns port associated with service. Protocol must be "tcp" or "udp" */

/* {{{ proto string getservbyport(int port, string protocol)
   Returns service name associated with port. Protocol must be "tcp" or "udp" */

/* {{{ proto int getprotobyname(string name)
   Returns protocol number associated with name as per /etc/protocols */

/* {{{ proto string getprotobynumber(int proto)
   Returns protocol name associated with protocol number proto */

/* {{{ proto bool register_tick_function(string function_name [, mixed arg [, mixed ... ]])
   Registers a tick callback function */

/* {{{ proto array sys_getloadavg()
 */
