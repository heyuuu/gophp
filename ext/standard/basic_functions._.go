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

var ArginfoSetTimeLimit []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("seconds"),
}
var ArginfoHeaderRegisterCallback []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("callback"),
}
var ArginfoObStart []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("user_function"),
	zend.MakeArgInfo("chunk_size"),
	zend.MakeArgInfo("flags"),
}
var ArginfoObFlush []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoObClean []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoObEndFlush []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoObEndClean []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoObGetFlush []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoObGetClean []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoObGetContents []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoObGetLevel []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoObGetLength []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoObListHandlers []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoObGetStatus []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("full_status"),
}
var ArginfoObImplicitFlush []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("flag"),
}
var ArginfoOutputResetRewriteVars []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoOutputAddRewriteVar []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("name"),
	zend.MakeArgInfo("value"),
}
var ArginfoStreamWrapperRegister []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("protocol"),
	zend.MakeArgInfo("classname"),
	zend.MakeArgInfo("flags"),
}
var ArginfoStreamWrapperUnregister []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("protocol"),
}
var ArginfoStreamWrapperRestore []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("protocol"),
}
var ArginfoKrsort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
	zend.MakeArgInfo("sort_flags"),
}
var ArginfoKsort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
	zend.MakeArgInfo("sort_flags"),
}
var ArginfoCount []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("var"),
	zend.MakeArgInfo("mode"),
}
var ArginfoNatsort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
}
var ArginfoNatcasesort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
}
var ArginfoAsort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
	zend.MakeArgInfo("sort_flags"),
}
var ArginfoArsort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
	zend.MakeArgInfo("sort_flags"),
}
var ArginfoSort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
	zend.MakeArgInfo("sort_flags"),
}
var ArginfoRsort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
	zend.MakeArgInfo("sort_flags"),
}
var ArginfoUsort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
	zend.MakeArgInfo("cmp_function"),
}
var ArginfoUasort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
	zend.MakeArgInfo("cmp_function"),
}
var ArginfoUksort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
	zend.MakeArgInfo("cmp_function"),
}
var ArginfoEnd []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
}
var ArginfoPrev []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
}
var ArginfoNext []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
}
var ArginfoReset []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
}
var ArginfoCurrent []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg"),
}
var ArginfoKey []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg"),
}
var ArginfoMin []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("args", ArgInfoVariadic()),
}
var ArginfoMax []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("args", ArgInfoVariadic()),
}
var ArginfoArrayWalk []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("input", ArgInfoByRef(1)),
	zend.MakeArgInfo("funcname"),
	zend.MakeArgInfo("userdata"),
}
var ArginfoArrayWalkRecursive []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("input", ArgInfoByRef(1)),
	zend.MakeArgInfo("funcname"),
	zend.MakeArgInfo("userdata"),
}
var ArginfoInArray []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("needle"),
	zend.MakeArgInfo("haystack"),
	zend.MakeArgInfo("strict"),
}
var ArginfoArraySearch []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("needle"),
	zend.MakeArgInfo("haystack"),
	zend.MakeArgInfo("strict"),
}
var ArginfoExtract []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arg", ArgInfoByRef(zend.ZEND_SEND_PREFER_REF)),
	zend.MakeArgInfo("extract_type"),
	zend.MakeArgInfo("prefix"),
}
var ArginfoCompact []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("var_names", ArgInfoVariadic()),
}
var ArginfoArrayFill []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("start_key"),
	zend.MakeArgInfo("num"),
	zend.MakeArgInfo("val"),
}
var ArginfoArrayFillKeys []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("keys"),
	zend.MakeArgInfo("val"),
}
var ArginfoRange []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("low"),
	zend.MakeArgInfo("high"),
	zend.MakeArgInfo("step"),
}
var ArginfoShuffle []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
}
var ArginfoArrayPush []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("stack", ArgInfoByRef(1)),
	zend.MakeArgInfo("vars", ArgInfoVariadic()),
}
var ArginfoArrayPop []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("stack", ArgInfoByRef(1)),
}
var ArginfoArrayShift []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("stack", ArgInfoByRef(1)),
}
var ArginfoArrayUnshift []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("stack", ArgInfoByRef(1)),
	zend.MakeArgInfo("vars", ArgInfoVariadic()),
}
var ArginfoArraySplice []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("arg", ArgInfoByRef(1)),
	zend.MakeArgInfo("offset"),
	zend.MakeArgInfo("length"),
	zend.MakeArgInfo("replacement"),
}
var ArginfoArraySlice []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("arg"),
	zend.MakeArgInfo("offset"),
	zend.MakeArgInfo("length"),
	zend.MakeArgInfo("preserve_keys"),
}
var ArginfoArrayMerge []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("arrays", ArgInfoVariadic()),
}
var ArginfoArrayMergeRecursive []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("arrays", ArgInfoVariadic()),
}
var ArginfoArrayReplace []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arrays", ArgInfoVariadic()),
}
var ArginfoArrayReplaceRecursive []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arrays", ArgInfoVariadic()),
}
var ArginfoArrayKeys []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arg"),
	zend.MakeArgInfo("search_value"),
	zend.MakeArgInfo("strict"),
}
var ArginfoArrayKeyFirst []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg"),
}
var ArginfoArrayKeyLast []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg"),
}
var ArginfoArrayValues []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg"),
}
var ArginfoArrayCountValues []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg"),
}
var ArginfoArrayColumn []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("arg"),
	zend.MakeArgInfo("column_key"),
	zend.MakeArgInfo("index_key"),
}
var ArginfoArrayReverse []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("input"),
	zend.MakeArgInfo("preserve_keys"),
}
var ArginfoArrayPad []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg"),
	zend.MakeArgInfo("pad_size"),
	zend.MakeArgInfo("pad_value"),
}
var ArginfoArrayFlip []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg"),
}
var ArginfoArrayChangeKeyCase []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("input"),
	zend.MakeArgInfo("case"),
}
var ArginfoArrayUnique []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arg"),
	zend.MakeArgInfo("flags"),
}
var ArginfoArrayIntersectKey []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arrays", ArgInfoVariadic()),
}
var ArginfoArrayIntersectUkey []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arr2"),
	zend.MakeArgInfo("callback_key_compare_func"),
}
var ArginfoArrayIntersect []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arrays", ArgInfoVariadic()),
}
var ArginfoArrayUintersect []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arr2"),
	zend.MakeArgInfo("callback_data_compare_func"),
}
var ArginfoArrayIntersectAssoc []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arrays", ArgInfoVariadic()),
}
var ArginfoArrayUintersectAssoc []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arr2"),
	zend.MakeArgInfo("callback_data_compare_func"),
}
var ArginfoArrayIntersectUassoc []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arr2"),
	zend.MakeArgInfo("callback_key_compare_func"),
}
var ArginfoArrayUintersectUassoc []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arr2"),
	zend.MakeArgInfo("callback_data_compare_func"),
	zend.MakeArgInfo("callback_key_compare_func"),
}
var ArginfoArrayDiffKey []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arrays", ArgInfoVariadic()),
}
var ArginfoArrayDiffUkey []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arr2"),
	zend.MakeArgInfo("callback_key_comp_func"),
}
var ArginfoArrayDiff []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arrays", ArgInfoVariadic()),
}
var ArginfoArrayUdiff []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arr2"),
	zend.MakeArgInfo("callback_data_comp_func"),
}
var ArginfoArrayDiffAssoc []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arrays", ArgInfoVariadic()),
}
var ArginfoArrayDiffUassoc []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arr2"),
	zend.MakeArgInfo("callback_data_comp_func"),
}
var ArginfoArrayUdiffAssoc []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arr2"),
	zend.MakeArgInfo("callback_key_comp_func"),
}
var ArginfoArrayUdiffUassoc []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arr1"),
	zend.MakeArgInfo("arr2"),
	zend.MakeArgInfo("callback_data_comp_func"),
	zend.MakeArgInfo("callback_key_comp_func"),
}
var ArginfoArrayMultisort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arr1", ArgInfoByRef(zend.ZEND_SEND_PREFER_REF)),
	zend.MakeArgInfo("sort_order", ArgInfoByRef(zend.ZEND_SEND_PREFER_REF)),
	zend.MakeArgInfo("sort_flags", ArgInfoByRef(zend.ZEND_SEND_PREFER_REF)),
	zend.MakeArgInfo("arr2", ArgInfoByRef(zend.ZEND_SEND_PREFER_REF), ArgInfoVariadic()),
}
var ArginfoArrayRand []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arg"),
	zend.MakeArgInfo("num_req"),
}
var ArginfoArraySum []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg"),
}
var ArginfoArrayProduct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg"),
}
var ArginfoArrayReduce []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("arg"),
	zend.MakeArgInfo("callback"),
	zend.MakeArgInfo("initial"),
}
var ArginfoArrayFilter []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("arg"),
	zend.MakeArgInfo("callback"),
	zend.MakeArgInfo("use_keys"),
}
var ArginfoArrayMap []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("callback"),
	zend.MakeArgInfo("arrays", ArgInfoVariadic()),
}
var ArginfoArrayKeyExists []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("key"),
	zend.MakeArgInfo("search"),
}
var ArginfoArrayChunk []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("arg"),
	zend.MakeArgInfo("size"),
	zend.MakeArgInfo("preserve_keys"),
}
var ArginfoArrayCombine []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("keys"),
	zend.MakeArgInfo("values"),
}
var ArginfoGetMagicQuotesGpc []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoGetMagicQuotesRuntime []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoConstant []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("const_name"),
}
var ArginfoInetNtop []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("in_addr"),
}
var ArginfoInetPton []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("ip_address"),
}
var ArginfoIp2long []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("ip_address"),
}
var ArginfoLong2ip []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("proper_address"),
}
var ArginfoGetenv []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("varname"),
	zend.MakeArgInfo("local_only"),
}
var ArginfoPutenv []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("setting"),
}
var ArginfoGetopt []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("options"),
	zend.MakeArgInfo("opts"),
	zend.MakeArgInfo("optind", ArgInfoByRef(1)),
}
var ArginfoFlush []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoSleep []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("seconds"),
}
var ArginfoUsleep []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("micro_seconds"),
}
var ArginfoTimeNanosleep []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("seconds"),
	zend.MakeArgInfo("nanoseconds"),
}
var ArginfoTimeSleepUntil []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("timestamp"),
}
var ArginfoGetCurrentUser []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoGetCfgVar []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("option_name"),
}
var ArginfoErrorLog []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("message"),
	zend.MakeArgInfo("message_type"),
	zend.MakeArgInfo("destination"),
	zend.MakeArgInfo("extra_headers"),
}
var ArginfoErrorGetLast []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
}
var ArginfoErrorClearLast []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
}
var ArginfoCallUserFunc []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("function_name"),
	zend.MakeArgInfo("parameters", ArgInfoVariadic()),
}
var ArginfoCallUserFuncArray []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("function_name"),
	zend.MakeArgInfo("parameters"),
}
var ArginfoForwardStaticCall []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("function_name"),
	zend.MakeArgInfo("parameters", ArgInfoVariadic()),
}
var ArginfoForwardStaticCallArray []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("function_name"),
	zend.MakeArgInfo("parameters"),
}
var ArginfoRegisterShutdownFunction []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("function_name"),
	zend.MakeArgInfo("parameters", ArgInfoVariadic()),
}
var ArginfoHighlightFile []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("file_name"),
	zend.MakeArgInfo("return"),
}
var ArginfoPhpStripWhitespace []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("file_name"),
}
var ArginfoHighlightString []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("string"),
	zend.MakeArgInfo("return"),
}
var ArginfoIniGet []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("varname"),
}
var ArginfoIniGetAll []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("extension"),
	zend.MakeArgInfo("details"),
}
var ArginfoIniSet []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("varname"),
	zend.MakeArgInfo("newvalue"),
}
var ArginfoIniRestore []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("varname"),
}
var ArginfoSetIncludePath []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("new_include_path"),
}
var ArginfoGetIncludePath []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoRestoreIncludePath []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoPrintR []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("var"),
	zend.MakeArgInfo("return"),
}
var ArginfoConnectionAborted []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoConnectionStatus []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoIgnoreUserAbort []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("value"),
}
var ArginfoGetservbyname []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("service"),
	zend.MakeArgInfo("protocol"),
}
var ArginfoGetservbyport []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("port"),
	zend.MakeArgInfo("protocol"),
}
var ArginfoGetprotobyname []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("name"),
}
var ArginfoGetprotobynumber []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("proto"),
}
var ArginfoRegisterTickFunction []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("function_name"),
	zend.MakeArgInfo("parameters", ArgInfoVariadic()),
}
var ArginfoUnregisterTickFunction []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("function_name"),
}
var ArginfoIsUploadedFile []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("path"),
}
var ArginfoMoveUploadedFile []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("path"),
	zend.MakeArgInfo("new_path"),
}
var ArginfoParseIniFile []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("process_sections"),
	zend.MakeArgInfo("scanner_mode"),
}
var ArginfoParseIniString []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("ini_string"),
	zend.MakeArgInfo("process_sections"),
	zend.MakeArgInfo("scanner_mode"),
}
var ArginfoSysGetloadavg []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoAssert []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("assertion"),
	zend.MakeArgInfo("description"),
}
var ArginfoAssertOptions []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("what"),
	zend.MakeArgInfo("value"),
}
var ArginfoBase64Encode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoBase64Decode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("strict"),
}
var ArginfoGetBrowser []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("browser_name"),
	zend.MakeArgInfo("return_array"),
}
var ArginfoCrc32 []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoCrypt []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("salt"),
}
var ArginfoConvertCyrString []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("from"),
	zend.MakeArgInfo("to"),
}
var ArginfoStrptime []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("timestamp"),
	zend.MakeArgInfo("format"),
}
var ArginfoOpendir []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("path"),
	zend.MakeArgInfo("context"),
}
var ArginfoDir []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("directory"),
	zend.MakeArgInfo("context"),
}
var ArginfoClosedir []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("dir_handle"),
}
var ArginfoChroot []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("directory"),
}
var ArginfoChdir []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("directory"),
}
var ArginfoGetcwd []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoRewinddir []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("dir_handle"),
}
var ArginfoReaddir []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("dir_handle"),
}
var ArginfoGlob []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("pattern"),
	zend.MakeArgInfo("flags"),
}
var ArginfoScandir []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("dir"),
	zend.MakeArgInfo("sorting_order"),
	zend.MakeArgInfo("context"),
}
var ArginfoGethostbyaddr []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("ip_address"),
}
var ArginfoGethostbyname []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("hostname"),
}
var ArginfoGethostbynamel []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("hostname"),
}
var ArginfoGethostname []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoNetGetInterfaces []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoDnsCheckRecord []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("host"),
	zend.MakeArgInfo("type"),
}
var ArginfoDnsGetRecord []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("hostname"),
	zend.MakeArgInfo("type"),
	zend.MakeArgInfo("authns", ArgInfoType(zend.ZEND_TYPE_ENCODE(zend.IS_ARRAY, 1)), ArgInfoByRef(1)),
	zend.MakeArgInfo("addtl", ArgInfoType(zend.ZEND_TYPE_ENCODE(zend.IS_ARRAY, 1)), ArgInfoByRef(1)),
	zend.MakeArgInfo("raw"),
}
var ArginfoDnsGetMx []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("hostname"),
	zend.MakeArgInfo("mxhosts", ArgInfoByRef(1)),
	zend.MakeArgInfo("weight", ArgInfoByRef(1)),
}
var ArginfoExec []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("command"),
	zend.MakeArgInfo("output", ArgInfoByRef(1)),
	zend.MakeArgInfo("return_value", ArgInfoByRef(1)),
}
var ArginfoSystem []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("command"),
	zend.MakeArgInfo("return_value", ArgInfoByRef(1)),
}
var ArginfoPassthru []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("command"),
	zend.MakeArgInfo("return_value", ArgInfoByRef(1)),
}
var ArginfoEscapeshellcmd []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("command"),
}
var ArginfoEscapeshellarg []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("arg"),
}
var ArginfoShellExec []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("cmd"),
}
var ArginfoProcNice []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("priority"),
}
var ArginfoFlock []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("fp"),
	zend.MakeArgInfo("operation"),
	zend.MakeArgInfo("wouldblock", ArgInfoByRef(1)),
}
var ArginfoGetMetaTags []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("use_include_path"),
}
var ArginfoFileGetContents []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("flags"),
	zend.MakeArgInfo("context"),
	zend.MakeArgInfo("offset"),
	zend.MakeArgInfo("maxlen"),
}
var ArginfoFilePutContents []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("data"),
	zend.MakeArgInfo("flags"),
	zend.MakeArgInfo("context"),
}
var ArginfoFile []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("flags"),
	zend.MakeArgInfo("context"),
}
var ArginfoTempnam []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("dir"),
	zend.MakeArgInfo("prefix"),
}
var ArginfoTmpfile []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoFopen []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("mode"),
	zend.MakeArgInfo("use_include_path"),
	zend.MakeArgInfo("context"),
}
var ArginfoFclose []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
}
var ArginfoPopen []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("command"),
	zend.MakeArgInfo("mode"),
}
var ArginfoPclose []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
}
var ArginfoFeof []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
}
var ArginfoFgets []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("fp"),
	zend.MakeArgInfo("length"),
}
var ArginfoFgetc []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
}
var ArginfoFgetss []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("fp"),
	zend.MakeArgInfo("length"),
	zend.MakeArgInfo("allowable_tags"),
}
var ArginfoFscanf []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("stream"),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("vars", ArgInfoByRef(1), ArgInfoVariadic()),
}
var ArginfoFwrite []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("fp"),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("length"),
}
var ArginfoFflush []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
}
var ArginfoRewind []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
}
var ArginfoFtell []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
}
var ArginfoFseek []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("fp"),
	zend.MakeArgInfo("offset"),
	zend.MakeArgInfo("whence"),
}
var ArginfoMkdir []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("pathname"),
	zend.MakeArgInfo("mode"),
	zend.MakeArgInfo("recursive"),
	zend.MakeArgInfo("context"),
}
var ArginfoRmdir []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("dirname"),
	zend.MakeArgInfo("context"),
}
var ArginfoReadfile []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("flags"),
	zend.MakeArgInfo("context"),
}
var ArginfoUmask []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("mask"),
}
var ArginfoFpassthru []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
}
var ArginfoRename []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("old_name"),
	zend.MakeArgInfo("new_name"),
	zend.MakeArgInfo("context"),
}
var ArginfoUnlink []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("context"),
}
var ArginfoFtruncate []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
	zend.MakeArgInfo("size"),
}
var ArginfoFstat []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
}
var ArginfoCopy []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("source_file"),
	zend.MakeArgInfo("destination_file"),
	zend.MakeArgInfo("context"),
}
var ArginfoFread []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
	zend.MakeArgInfo("length"),
}
var ArginfoFputcsv []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("fp"),
	zend.MakeArgInfo("fields"),
	zend.MakeArgInfo("delimiter"),
	zend.MakeArgInfo("enclosure"),
	zend.MakeArgInfo("escape_char"),
}
var ArginfoFgetcsv []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("fp"),
	zend.MakeArgInfo("length"),
	zend.MakeArgInfo("delimiter"),
	zend.MakeArgInfo("enclosure"),
	zend.MakeArgInfo("escape"),
}
var ArginfoRealpath []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("path"),
}
var ArginfoFnmatch []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("pattern"),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("flags"),
}
var ArginfoSysGetTempDir []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoDiskTotalSpace []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("path"),
}
var ArginfoDiskFreeSpace []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("path"),
}
var ArginfoChgrp []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("group"),
}
var ArginfoChown []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("user"),
}
var ArginfoLchgrp []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("group"),
}
var ArginfoLchown []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("user"),
}
var ArginfoChmod []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("mode"),
}
var ArginfoTouch []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("time"),
	zend.MakeArgInfo("atime"),
}
var ArginfoClearstatcache []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("clear_realpath_cache"),
	zend.MakeArgInfo("filename"),
}
var ArginfoRealpathCacheSize []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoRealpathCacheGet []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoFileperms []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoFileinode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoFilesize []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoFileowner []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoFilegroup []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoFileatime []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoFilemtime []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoFilectime []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoFiletype []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoIsWritable []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoIsReadable []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoIsExecutable []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoIsFile []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoIsDir []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoIsLink []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoFileExists []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoLstat []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoStat []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoSprintf []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("args", ArgInfoVariadic()),
}
var ArginfoVsprintf []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("args"),
}
var ArginfoPrintf []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("args", ArgInfoVariadic()),
}
var ArginfoVprintf []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("args"),
}
var ArginfoFprintf []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("stream"),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("args", ArgInfoVariadic()),
}
var ArginfoVfprintf []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("stream"),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("args"),
}
var ArginfoFsockopen []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("hostname"),
	zend.MakeArgInfo("port"),
	zend.MakeArgInfo("errno", ArgInfoByRef(1)),
	zend.MakeArgInfo("errstr", ArgInfoByRef(1)),
	zend.MakeArgInfo("timeout"),
}
var ArginfoPfsockopen []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("hostname"),
	zend.MakeArgInfo("port"),
	zend.MakeArgInfo("errno", ArgInfoByRef(1)),
	zend.MakeArgInfo("errstr", ArgInfoByRef(1)),
	zend.MakeArgInfo("timeout"),
}
var ArginfoFtok []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("pathname"),
	zend.MakeArgInfo("proj"),
}
var ArginfoHeader []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("header"),
	zend.MakeArgInfo("replace"),
	zend.MakeArgInfo("http_response_code"),
}
var ArginfoHeaderRemove []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("name"),
}
var ArginfoSetcookie []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("name"),
	zend.MakeArgInfo("value"),
	zend.MakeArgInfo("expires_or_options"),
	zend.MakeArgInfo("path"),
	zend.MakeArgInfo("domain"),
	zend.MakeArgInfo("secure"),
	zend.MakeArgInfo("httponly"),
}
var ArginfoSetrawcookie []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("name"),
	zend.MakeArgInfo("value"),
	zend.MakeArgInfo("expires_or_options"),
	zend.MakeArgInfo("path"),
	zend.MakeArgInfo("domain"),
	zend.MakeArgInfo("secure"),
	zend.MakeArgInfo("httponly"),
}
var ArginfoHeadersSent []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("file", ArgInfoByRef(1)),
	zend.MakeArgInfo("line", ArgInfoByRef(1)),
}
var ArginfoHeadersList []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoHttpResponseCode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("response_code"),
}
var ArginfoHrtime []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("get_as_number"),
}
var ArginfoHtmlspecialchars []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("string"),
	zend.MakeArgInfo("quote_style"),
	zend.MakeArgInfo("encoding"),
	zend.MakeArgInfo("double_encode"),
}
var ArginfoHtmlspecialcharsDecode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("string"),
	zend.MakeArgInfo("quote_style"),
}
var ArginfoHtmlEntityDecode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("string"),
	zend.MakeArgInfo("quote_style"),
	zend.MakeArgInfo("encoding"),
}
var ArginfoHtmlentities []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("string"),
	zend.MakeArgInfo("quote_style"),
	zend.MakeArgInfo("encoding"),
	zend.MakeArgInfo("double_encode"),
}
var ArginfoGetHtmlTranslationTable []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("table"),
	zend.MakeArgInfo("quote_style"),
	zend.MakeArgInfo("encoding"),
}
var ArginfoHttpBuildQuery []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("formdata"),
	zend.MakeArgInfo("prefix"),
	zend.MakeArgInfo("arg_separator"),
	zend.MakeArgInfo("enc_type"),
}
var ArginfoImageTypeToMimeType []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("imagetype"),
}
var ArginfoImageTypeToExtension []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("imagetype"),
	zend.MakeArgInfo("include_dot"),
}
var ArginfoGetimagesize []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("imagefile"),
	zend.MakeArgInfo("info", ArgInfoByRef(1)),
}
var ArginfoPhpinfo []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("what"),
}
var ArginfoPhpversion []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("extension"),
}
var ArginfoPhpcredits []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("flag"),
}
var ArginfoPhpSapiName []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoPhpUname []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("mode"),
}
var ArginfoPhpIniScannedFiles []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoPhpIniLoadedFile []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoIptcembed []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("iptcdata"),
	zend.MakeArgInfo("jpeg_file_name"),
	zend.MakeArgInfo("spool"),
}
var ArginfoIptcparse []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("iptcdata"),
}
var ArginfoLcgValue []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoLevenshtein []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("str1"),
	zend.MakeArgInfo("str2"),
	zend.MakeArgInfo("cost_ins"),
	zend.MakeArgInfo("cost_rep"),
	zend.MakeArgInfo("cost_del"),
}
var ArginfoReadlink []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoLinkinfo []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoSymlink []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("target"),
	zend.MakeArgInfo("link"),
}
var ArginfoLink []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("target"),
	zend.MakeArgInfo("link"),
}
var ArginfoEzmlmHash []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("addr"),
}
var ArginfoMail []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(3),
	zend.MakeArgInfo("to"),
	zend.MakeArgInfo("subject"),
	zend.MakeArgInfo("message"),
	zend.MakeArgInfo("additional_headers"),
	zend.MakeArgInfo("additional_parameters"),
}
var ArginfoAbs []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoCeil []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoFloor []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoRound []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("number"),
	zend.MakeArgInfo("precision"),
	zend.MakeArgInfo("mode"),
}
var ArginfoSin []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoCos []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoTan []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoAsin []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoAcos []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoAtan []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoAtan2 []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("y"),
	zend.MakeArgInfo("x"),
}
var ArginfoSinh []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoCosh []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoTanh []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoAsinh []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoAcosh []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoAtanh []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoPi []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoIsFinite []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("val"),
}
var ArginfoIsInfinite []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("val"),
}
var ArginfoIsNan []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("val"),
}
var ArginfoPow []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("base"),
	zend.MakeArgInfo("exponent"),
}
var ArginfoExp []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoExpm1 []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoLog1p []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoLog []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("number"),
	zend.MakeArgInfo("base"),
}
var ArginfoLog10 []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoSqrt []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoHypot []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("num1"),
	zend.MakeArgInfo("num2"),
}
var ArginfoDeg2rad []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoRad2deg []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
}
var ArginfoBindec []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("binary_number"),
}
var ArginfoHexdec []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("hexadecimal_number"),
}
var ArginfoOctdec []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("octal_number"),
}
var ArginfoDecbin []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("decimal_number"),
}
var ArginfoDecoct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("decimal_number"),
}
var ArginfoDechex []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("decimal_number"),
}
var ArginfoBaseConvert []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("number"),
	zend.MakeArgInfo("frombase"),
	zend.MakeArgInfo("tobase"),
}
var ArginfoNumberFormat []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("number"),
	zend.MakeArgInfo("num_decimal_places"),
	zend.MakeArgInfo("dec_separator"),
	zend.MakeArgInfo("thousands_separator"),
}
var ArginfoFmod []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("x"),
	zend.MakeArgInfo("y"),
}
var ArginfoIntdiv []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("dividend"),
	zend.MakeArgInfo("divisor"),
}
var ArginfoMd5 []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("raw_output"),
}
var ArginfoMd5File []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("raw_output"),
}
var ArginfoMetaphone []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("text"),
	zend.MakeArgInfo("phones"),
}
var ArginfoMicrotime []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("get_as_float"),
}
var ArginfoGettimeofday []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("get_as_float"),
}
var ArginfoGetrusage []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("who"),
}
var ArginfoPack []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("args", ArgInfoVariadic()),
}
var ArginfoUnpack []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("input"),
	zend.MakeArgInfo("offset"),
}
var ArginfoGetmyuid []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoGetmygid []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoGetmypid []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoGetmyinode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoGetlastmod []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoPasswordHash []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("password"),
	zend.MakeArgInfo("algo"),
	zend.MakeArgInfo("options"),
}
var ArginfoPasswordGetInfo []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("hash"),
}
var ArginfoPasswordNeedsRehash []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("hash"),
	zend.MakeArgInfo("algo"),
	zend.MakeArgInfo("options"),
}
var ArginfoPasswordVerify []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("password"),
	zend.MakeArgInfo("hash"),
}
var ArginfoPasswordAlgos []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoProcTerminate []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("process"),
	zend.MakeArgInfo("signal"),
}
var ArginfoProcClose []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("process"),
}
var ArginfoProcGetStatus []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("process"),
}
var ArginfoProcOpen []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(3),
	zend.MakeArgInfo("command"),
	zend.MakeArgInfo("descriptorspec"),
	zend.MakeArgInfo("pipes", ArgInfoByRef(1)),
	zend.MakeArgInfo("cwd"),
	zend.MakeArgInfo("env"),
	zend.MakeArgInfo("other_options"),
}
var ArginfoQuotedPrintableDecode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoQuotedPrintableEncode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoMtSrand []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("seed"),
	zend.MakeArgInfo("mode"),
}
var ArginfoMtRand []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("min"),
	zend.MakeArgInfo("max"),
}
var ArginfoMtGetrandmax []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoRandomBytes []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("length"),
}
var ArginfoRandomInt []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("min"),
	zend.MakeArgInfo("max"),
}
var ArginfoSha1 []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("raw_output"),
}
var ArginfoSha1File []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("filename"),
	zend.MakeArgInfo("raw_output"),
}
var ArginfoSoundex []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoStreamSocketPair []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("domain"),
	zend.MakeArgInfo("type"),
	zend.MakeArgInfo("protocol"),
}
var ArginfoStreamSocketClient []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("remoteaddress"),
	zend.MakeArgInfo("errcode", ArgInfoByRef(1)),
	zend.MakeArgInfo("errstring", ArgInfoByRef(1)),
	zend.MakeArgInfo("timeout"),
	zend.MakeArgInfo("flags"),
	zend.MakeArgInfo("context"),
}
var ArginfoStreamSocketServer []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("localaddress"),
	zend.MakeArgInfo("errcode", ArgInfoByRef(1)),
	zend.MakeArgInfo("errstring", ArgInfoByRef(1)),
	zend.MakeArgInfo("flags"),
	zend.MakeArgInfo("context"),
}
var ArginfoStreamSocketAccept []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("serverstream"),
	zend.MakeArgInfo("timeout"),
	zend.MakeArgInfo("peername", ArgInfoByRef(1)),
}
var ArginfoStreamSocketGetName []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("stream"),
	zend.MakeArgInfo("want_peer"),
}
var ArginfoStreamSocketSendto []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("stream"),
	zend.MakeArgInfo("data"),
	zend.MakeArgInfo("flags"),
	zend.MakeArgInfo("target_addr"),
}
var ArginfoStreamSocketRecvfrom []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("stream"),
	zend.MakeArgInfo("amount"),
	zend.MakeArgInfo("flags"),
	zend.MakeArgInfo("remote_addr", ArgInfoByRef(1)),
}
var ArginfoStreamGetContents []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("source"),
	zend.MakeArgInfo("maxlen"),
	zend.MakeArgInfo("offset"),
}
var ArginfoStreamCopyToStream []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("source"),
	zend.MakeArgInfo("dest"),
	zend.MakeArgInfo("maxlen"),
	zend.MakeArgInfo("pos"),
}
var ArginfoStreamGetMetaData []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
}
var ArginfoStreamGetTransports []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoStreamGetWrappers []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoStreamResolveIncludePath []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filename"),
}
var ArginfoStreamIsLocal []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("stream"),
}
var ArginfoStreamSupportsLock []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("stream"),
}
var ArginfoStreamIsatty []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("stream"),
}
var ArginfoStreamSelect []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(4),
	zend.MakeArgInfo("read_streams", ArgInfoByRef(1)),
	zend.MakeArgInfo("write_streams", ArgInfoByRef(1)),
	zend.MakeArgInfo("except_streams", ArgInfoByRef(1)),
	zend.MakeArgInfo("tv_sec"),
	zend.MakeArgInfo("tv_usec"),
}
var ArginfoStreamContextGetOptions []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("stream_or_context"),
}
var ArginfoStreamContextSetOption []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("stream_or_context"),
	zend.MakeArgInfo("wrappername"),
	zend.MakeArgInfo("optionname"),
	zend.MakeArgInfo("value"),
}
var ArginfoStreamContextSetParams []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("stream_or_context"),
	zend.MakeArgInfo("options"),
}
var ArginfoStreamContextGetParams []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("stream_or_context"),
}
var ArginfoStreamContextGetDefault []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("options"),
}
var ArginfoStreamContextSetDefault []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("options"),
}
var ArginfoStreamContextCreate []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("options"),
	zend.MakeArgInfo("params"),
}
var ArginfoStreamFilterPrepend []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("stream"),
	zend.MakeArgInfo("filtername"),
	zend.MakeArgInfo("read_write"),
	zend.MakeArgInfo("filterparams"),
}
var ArginfoStreamFilterAppend []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("stream"),
	zend.MakeArgInfo("filtername"),
	zend.MakeArgInfo("read_write"),
	zend.MakeArgInfo("filterparams"),
}
var ArginfoStreamFilterRemove []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("stream_filter"),
}
var ArginfoStreamGetLine []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("stream"),
	zend.MakeArgInfo("maxlen"),
	zend.MakeArgInfo("ending"),
}
var ArginfoStreamSetBlocking []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("socket"),
	zend.MakeArgInfo("mode"),
}
var ArginfoStreamSetTimeout []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("stream"),
	zend.MakeArgInfo("seconds"),
	zend.MakeArgInfo("microseconds"),
}
var ArginfoStreamSetReadBuffer []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
	zend.MakeArgInfo("buffer"),
}
var ArginfoStreamSetWriteBuffer []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
	zend.MakeArgInfo("buffer"),
}
var ArginfoStreamSetChunkSize []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("fp"),
	zend.MakeArgInfo("chunk_size"),
}
var ArginfoStreamSocketEnableCrypto []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("stream"),
	zend.MakeArgInfo("enable"),
	zend.MakeArgInfo("cryptokind"),
	zend.MakeArgInfo("sessionstream"),
}
var ArginfoStreamSocketShutdown []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("stream"),
	zend.MakeArgInfo("how"),
}
var ArginfoBin2hex []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("data"),
}
var ArginfoHex2bin []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("data"),
}
var ArginfoStrspn []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("mask"),
	zend.MakeArgInfo("start"),
	zend.MakeArgInfo("len"),
}
var ArginfoStrcspn []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("mask"),
	zend.MakeArgInfo("start"),
	zend.MakeArgInfo("len"),
}
var ArginfoStrcoll []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str1"),
	zend.MakeArgInfo("str2"),
}
var ArginfoTrim []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("character_mask"),
}
var ArginfoRtrim []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("character_mask"),
}
var ArginfoLtrim []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("character_mask"),
}
var ArginfoWordwrap []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("width"),
	zend.MakeArgInfo("break"),
	zend.MakeArgInfo("cut"),
}
var ArginfoExplode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("separator"),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("limit"),
}
var ArginfoImplode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("glue"),
	zend.MakeArgInfo("pieces"),
}
var ArginfoStrtok []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("token"),
}
var ArginfoStrtoupper []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoStrtolower []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoBasename []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("path"),
	zend.MakeArgInfo("suffix"),
}
var ArginfoDirname []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("path"),
	zend.MakeArgInfo("levels"),
}
var ArginfoPathinfo []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("path"),
	zend.MakeArgInfo("options"),
}
var ArginfoStristr []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("haystack"),
	zend.MakeArgInfo("needle"),
	zend.MakeArgInfo("part"),
}
var ArginfoStrstr []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("haystack"),
	zend.MakeArgInfo("needle"),
	zend.MakeArgInfo("part"),
}
var ArginfoStrpos []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("haystack"),
	zend.MakeArgInfo("needle"),
	zend.MakeArgInfo("offset"),
}
var ArginfoStripos []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("haystack"),
	zend.MakeArgInfo("needle"),
	zend.MakeArgInfo("offset"),
}
var ArginfoStrrpos []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("haystack"),
	zend.MakeArgInfo("needle"),
	zend.MakeArgInfo("offset"),
}
var ArginfoStrripos []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("haystack"),
	zend.MakeArgInfo("needle"),
	zend.MakeArgInfo("offset"),
}
var ArginfoStrrchr []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("haystack"),
	zend.MakeArgInfo("needle"),
}
var ArginfoChunkSplit []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("chunklen"),
	zend.MakeArgInfo("ending"),
}
var ArginfoSubstr []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("start"),
	zend.MakeArgInfo("length"),
}
var ArginfoSubstrReplace []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(3),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("replace"),
	zend.MakeArgInfo("start"),
	zend.MakeArgInfo("length"),
}
var ArginfoQuotemeta []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoOrd []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("character"),
}
var ArginfoChr []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("codepoint"),
}
var ArginfoUcfirst []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoLcfirst []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoUcwords []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("delimiters"),
}
var ArginfoStrtr []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("from"),
	zend.MakeArgInfo("to"),
}
var ArginfoStrrev []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoSimilarText []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("str1"),
	zend.MakeArgInfo("str2"),
	zend.MakeArgInfo("percent", ArgInfoByRef(1)),
}
var ArginfoAddcslashes []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("charlist"),
}
var ArginfoAddslashes []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoStripcslashes []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoStripslashes []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoStrReplace []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(3),
	zend.MakeArgInfo("search"),
	zend.MakeArgInfo("replace"),
	zend.MakeArgInfo("subject"),
	zend.MakeArgInfo("replace_count", ArgInfoByRef(1)),
}
var ArginfoStrIreplace []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(3),
	zend.MakeArgInfo("search"),
	zend.MakeArgInfo("replace"),
	zend.MakeArgInfo("subject"),
	zend.MakeArgInfo("replace_count", ArgInfoByRef(1)),
}
var ArginfoHebrev []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("max_chars_per_line"),
}
var ArginfoHebrevc []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("max_chars_per_line"),
}
var ArginfoNl2br []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("is_xhtml"),
}
var ArginfoStripTags []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("allowable_tags"),
}
var ArginfoSetlocale []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("category"),
	zend.MakeArgInfo("locales", ArgInfoVariadic()),
}
var ArginfoParseStr []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("encoded_string"),
	zend.MakeArgInfo("result", ArgInfoByRef(1)),
}
var ArginfoStrGetcsv []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("string"),
	zend.MakeArgInfo("delimiter"),
	zend.MakeArgInfo("enclosure"),
	zend.MakeArgInfo("escape"),
}
var ArginfoStrRepeat []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("input"),
	zend.MakeArgInfo("mult"),
}
var ArginfoCountChars []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("input"),
	zend.MakeArgInfo("mode"),
}
var ArginfoStrnatcmp []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("s1"),
	zend.MakeArgInfo("s2"),
}
var ArginfoLocaleconv []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoStrnatcasecmp []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("s1"),
	zend.MakeArgInfo("s2"),
}
var ArginfoSubstrCount []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("haystack"),
	zend.MakeArgInfo("needle"),
	zend.MakeArgInfo("offset"),
	zend.MakeArgInfo("length"),
}
var ArginfoStrPad []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("input"),
	zend.MakeArgInfo("pad_length"),
	zend.MakeArgInfo("pad_string"),
	zend.MakeArgInfo("pad_type"),
}
var ArginfoSscanf []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("vars", ArgInfoByRef(1), ArgInfoVariadic()),
}
var ArginfoStrRot13 []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoStrShuffle []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoStrWordCount []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("charlist"),
}
var ArginfoMoneyFormat []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("value"),
}
var ArginfoStrSplit []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("split_length"),
}
var ArginfoStrpbrk []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("haystack"),
	zend.MakeArgInfo("char_list"),
}
var ArginfoSubstrCompare []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(3),
	zend.MakeArgInfo("main_str"),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("offset"),
	zend.MakeArgInfo("length"),
	zend.MakeArgInfo("case_sensitivity"),
}
var ArginfoUtf8Encode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("data"),
}
var ArginfoUtf8Decode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("data"),
}
var ArginfoOpenlog []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("ident"),
	zend.MakeArgInfo("option"),
	zend.MakeArgInfo("facility"),
}
var ArginfoCloselog []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoSyslog []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("priority"),
	zend.MakeArgInfo("message"),
}
var ArginfoGettype []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoSettype []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var", ArgInfoByRef(1)),
	zend.MakeArgInfo("type"),
}
var ArginfoIntval []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("var"),
	zend.MakeArgInfo("base"),
}
var ArginfoFloatval []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoStrval []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoBoolval []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoIsNull []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoIsResource []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoIsBool []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoIsInt []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoIsFloat []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoIsString []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoIsArray []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoIsObject []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoIsNumeric []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("value"),
}
var ArginfoIsScalar []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("value"),
}
var ArginfoIsCallable []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("var"),
	zend.MakeArgInfo("syntax_only"),
	zend.MakeArgInfo("callable_name", ArgInfoByRef(1)),
}
var ArginfoIsIterable []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("var"),
}
var ArginfoIsCountable []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoUniqid []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("prefix"),
	zend.MakeArgInfo("more_entropy"),
}
var ArginfoParseUrl []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("url"),
	zend.MakeArgInfo("component"),
}
var ArginfoUrlencode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoUrldecode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoRawurlencode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoRawurldecode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("str"),
}
var ArginfoGetHeaders []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("url"),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("context"),
}
var ArginfoStreamBucketMakeWriteable []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("brigade"),
}
var ArginfoStreamBucketPrepend []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("brigade"),
	zend.MakeArgInfo("bucket"),
}
var ArginfoStreamBucketAppend []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("brigade"),
	zend.MakeArgInfo("bucket"),
}
var ArginfoStreamBucketNew []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("stream"),
	zend.MakeArgInfo("buffer"),
}
var ArginfoStreamGetFilters []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoStreamFilterRegister []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("filtername"),
	zend.MakeArgInfo("classname"),
}
var ArginfoConvertUuencode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("data"),
}
var ArginfoConvertUudecode []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("data"),
}
var ArginfoVarDump []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("vars", ArgInfoVariadic()),
}
var ArginfoDebugZvalDump []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("vars", ArgInfoVariadic()),
}
var ArginfoVarExport []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("var"),
	zend.MakeArgInfo("return"),
}
var ArginfoSerialize []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("var"),
}
var ArginfoUnserialize []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("variable_representation"),
	zend.MakeArgInfo("allowed_classes"),
}
var ArginfoMemoryGetUsage []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("real_usage"),
}
var ArginfoMemoryGetPeakUsage []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("real_usage"),
}
var ArginfoVersionCompare []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("ver1"),
	zend.MakeArgInfo("ver2"),
	zend.MakeArgInfo("oper"),
}
var BasicFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntryEx("constant", 0, ZifConstant, ArginfoConstant),
	zend.MakeZendFunctionEntryEx("bin2hex", 0, ZifBin2hex, ArginfoBin2hex),
	zend.MakeZendFunctionEntryEx("hex2bin", 0, ZifHex2bin, ArginfoHex2bin),
	zend.MakeZendFunctionEntryEx("sleep", 0, ZifSleep, ArginfoSleep),
	zend.MakeZendFunctionEntryEx("usleep", 0, ZifUsleep, ArginfoUsleep),
	zend.MakeZendFunctionEntryEx("time_nanosleep", 0, ZifTimeNanosleep, ArginfoTimeNanosleep),
	zend.MakeZendFunctionEntryEx("time_sleep_until", 0, ZifTimeSleepUntil, ArginfoTimeSleepUntil),
	zend.MakeZendFunctionEntryEx("strptime", 0, ZifStrptime, ArginfoStrptime),
	zend.MakeZendFunctionEntryEx("flush", 0, ZifFlush, ArginfoFlush),
	zend.MakeZendFunctionEntryEx("wordwrap", 0, ZifWordwrap, ArginfoWordwrap),
	zend.MakeZendFunctionEntryEx("htmlspecialchars", 0, ZifHtmlspecialchars, ArginfoHtmlspecialchars),
	zend.MakeZendFunctionEntryEx("htmlentities", 0, ZifHtmlentities, ArginfoHtmlentities),
	zend.MakeZendFunctionEntryEx("html_entity_decode", 0, ZifHtmlEntityDecode, ArginfoHtmlEntityDecode),
	zend.MakeZendFunctionEntryEx("htmlspecialchars_decode", 0, ZifHtmlspecialcharsDecode, ArginfoHtmlspecialcharsDecode),
	zend.MakeZendFunctionEntryEx("get_html_translation_table", 0, ZifGetHtmlTranslationTable, ArginfoGetHtmlTranslationTable),
	zend.MakeZendFunctionEntryEx("sha1", 0, ZifSha1, ArginfoSha1),
	zend.MakeZendFunctionEntryEx("sha1_file", 0, ZifSha1File, ArginfoSha1File),
	zend.MakeZendFunctionEntryEx("md5", 0, PhpIfMd5, ArginfoMd5),
	zend.MakeZendFunctionEntryEx("md5_file", 0, PhpIfMd5File, ArginfoMd5File),
	zend.MakeZendFunctionEntryEx("crc32", 0, PhpIfCrc32, ArginfoCrc32),
	zend.MakeZendFunctionEntryEx("iptcparse", 0, ZifIptcparse, ArginfoIptcparse),
	zend.MakeZendFunctionEntryEx("iptcembed", 0, ZifIptcembed, ArginfoIptcembed),
	zend.MakeZendFunctionEntryEx("getimagesize", 0, ZifGetimagesize, ArginfoGetimagesize),
	zend.MakeZendFunctionEntryEx("getimagesizefromstring", 0, ZifGetimagesizefromstring, ArginfoGetimagesize),
	zend.MakeZendFunctionEntryEx("image_type_to_mime_type", 0, ZifImageTypeToMimeType, ArginfoImageTypeToMimeType),
	zend.MakeZendFunctionEntryEx("image_type_to_extension", 0, ZifImageTypeToExtension, ArginfoImageTypeToExtension),
	zend.MakeZendFunctionEntryEx("phpversion", 0, ZifPhpversion, ArginfoPhpversion),
	zend.MakeZendFunctionEntryEx("phpcredits", 0, ZifPhpcredits, ArginfoPhpcredits),
	zend.MakeZendFunctionEntryEx("php_sapi_name", 0, ZifPhpSapiName, ArginfoPhpSapiName),
	zend.MakeZendFunctionEntryEx("php_uname", 0, ZifPhpUname, ArginfoPhpUname),
	zend.MakeZendFunctionEntryEx("php_ini_scanned_files", 0, ZifPhpIniScannedFiles, ArginfoPhpIniScannedFiles),
	zend.MakeZendFunctionEntryEx("php_ini_loaded_file", 0, ZifPhpIniLoadedFile, ArginfoPhpIniLoadedFile),
	zend.MakeZendFunctionEntryEx("strnatcmp", 0, ZifStrnatcmp, ArginfoStrnatcmp),
	zend.MakeZendFunctionEntryEx("strnatcasecmp", 0, ZifStrnatcasecmp, ArginfoStrnatcasecmp),
	zend.MakeZendFunctionEntryEx("substr_count", 0, ZifSubstrCount, ArginfoSubstrCount),
	zend.MakeZendFunctionEntryEx("strspn", 0, ZifStrspn, ArginfoStrspn),
	zend.MakeZendFunctionEntryEx("strcspn", 0, ZifStrcspn, ArginfoStrcspn),
	zend.MakeZendFunctionEntryEx("strtok", 0, ZifStrtok, ArginfoStrtok),
	zend.MakeZendFunctionEntryEx("strtoupper", 0, ZifStrtoupper, ArginfoStrtoupper),
	zend.MakeZendFunctionEntryEx("strtolower", 0, ZifStrtolower, ArginfoStrtolower),
	zend.MakeZendFunctionEntryEx("strpos", 0, ZifStrpos, ArginfoStrpos),
	zend.MakeZendFunctionEntryEx("stripos", 0, ZifStripos, ArginfoStripos),
	zend.MakeZendFunctionEntryEx("strrpos", 0, ZifStrrpos, ArginfoStrrpos),
	zend.MakeZendFunctionEntryEx("strripos", 0, ZifStrripos, ArginfoStrripos),
	zend.MakeZendFunctionEntryEx("strrev", 0, ZifStrrev, ArginfoStrrev),
	zend.MakeZendFunctionEntryEx("hebrev", 0, ZifHebrev, ArginfoHebrev),
	zend.MakeZendFunctionEntryEx("hebrevc", zend.ZEND_ACC_DEPRECATED, ZifHebrevc, ArginfoHebrevc),
	zend.MakeZendFunctionEntryEx("nl2br", 0, ZifNl2br, ArginfoNl2br),
	zend.MakeZendFunctionEntryEx("basename", 0, ZifBasename, ArginfoBasename),
	zend.MakeZendFunctionEntryEx("dirname", 0, ZifDirname, ArginfoDirname),
	zend.MakeZendFunctionEntryEx("pathinfo", 0, ZifPathinfo, ArginfoPathinfo),
	zend.MakeZendFunctionEntryEx("stripslashes", 0, ZifStripslashes, ArginfoStripslashes),
	zend.MakeZendFunctionEntryEx("stripcslashes", 0, ZifStripcslashes, ArginfoStripcslashes),
	zend.MakeZendFunctionEntryEx("strstr", 0, ZifStrstr, ArginfoStrstr),
	zend.MakeZendFunctionEntryEx("stristr", 0, ZifStristr, ArginfoStristr),
	zend.MakeZendFunctionEntryEx("strrchr", 0, ZifStrrchr, ArginfoStrrchr),
	zend.MakeZendFunctionEntryEx("str_shuffle", 0, ZifStrShuffle, ArginfoStrShuffle),
	zend.MakeZendFunctionEntryEx("str_word_count", 0, ZifStrWordCount, ArginfoStrWordCount),
	zend.MakeZendFunctionEntryEx("str_split", 0, ZifStrSplit, ArginfoStrSplit),
	zend.MakeZendFunctionEntryEx("strpbrk", 0, ZifStrpbrk, ArginfoStrpbrk),
	zend.MakeZendFunctionEntryEx("substr_compare", 0, ZifSubstrCompare, ArginfoSubstrCompare),
	zend.MakeZendFunctionEntryEx("utf8_encode", 0, ZifUtf8Encode, ArginfoUtf8Encode),
	zend.MakeZendFunctionEntryEx("utf8_decode", 0, ZifUtf8Decode, ArginfoUtf8Decode),
	zend.MakeZendFunctionEntryEx("strcoll", 0, ZifStrcoll, ArginfoStrcoll),
	zend.MakeZendFunctionEntryEx("money_format", zend.ZEND_ACC_DEPRECATED, ZifMoneyFormat, ArginfoMoneyFormat),
	zend.MakeZendFunctionEntryEx("substr", 0, ZifSubstr, ArginfoSubstr),
	zend.MakeZendFunctionEntryEx("substr_replace", 0, ZifSubstrReplace, ArginfoSubstrReplace),
	zend.MakeZendFunctionEntryEx("quotemeta", 0, ZifQuotemeta, ArginfoQuotemeta),
	zend.MakeZendFunctionEntryEx("ucfirst", 0, ZifUcfirst, ArginfoUcfirst),
	zend.MakeZendFunctionEntryEx("lcfirst", 0, ZifLcfirst, ArginfoLcfirst),
	zend.MakeZendFunctionEntryEx("ucwords", 0, ZifUcwords, ArginfoUcwords),
	zend.MakeZendFunctionEntryEx("strtr", 0, ZifStrtr, ArginfoStrtr),
	zend.MakeZendFunctionEntryEx("addslashes", 0, ZifAddslashes, ArginfoAddslashes),
	zend.MakeZendFunctionEntryEx("addcslashes", 0, ZifAddcslashes, ArginfoAddcslashes),
	zend.MakeZendFunctionEntryEx("rtrim", 0, ZifRtrim, ArginfoRtrim),
	zend.MakeZendFunctionEntryEx("str_replace", 0, ZifStrReplace, ArginfoStrReplace),
	zend.MakeZendFunctionEntryEx("str_ireplace", 0, ZifStrIreplace, ArginfoStrIreplace),
	zend.MakeZendFunctionEntryEx("str_repeat", 0, ZifStrRepeat, ArginfoStrRepeat),
	zend.MakeZendFunctionEntryEx("count_chars", 0, ZifCountChars, ArginfoCountChars),
	zend.MakeZendFunctionEntryEx("chunk_split", 0, ZifChunkSplit, ArginfoChunkSplit),
	zend.MakeZendFunctionEntryEx("trim", 0, ZifTrim, ArginfoTrim),
	zend.MakeZendFunctionEntryEx("ltrim", 0, ZifLtrim, ArginfoLtrim),
	zend.MakeZendFunctionEntryEx("strip_tags", 0, ZifStripTags, ArginfoStripTags),
	zend.MakeZendFunctionEntryEx("similar_text", 0, ZifSimilarText, ArginfoSimilarText),
	zend.MakeZendFunctionEntryEx("explode", 0, ZifExplode, ArginfoExplode),
	zend.MakeZendFunctionEntryEx("implode", 0, ZifImplode, ArginfoImplode),
	zend.MakeZendFunctionEntryEx("join", 0, ZifImplode, ArginfoImplode),
	zend.MakeZendFunctionEntryEx("setlocale", 0, ZifSetlocale, ArginfoSetlocale),
	zend.MakeZendFunctionEntryEx("localeconv", 0, ZifLocaleconv, ArginfoLocaleconv),
	zend.MakeZendFunctionEntryEx("soundex", 0, ZifSoundex, ArginfoSoundex),
	zend.MakeZendFunctionEntryEx("levenshtein", 0, ZifLevenshtein, ArginfoLevenshtein),
	zend.MakeZendFunctionEntryEx("chr", 0, ZifChr, ArginfoChr),
	zend.MakeZendFunctionEntryEx("ord", 0, ZifOrd, ArginfoOrd),
	zend.MakeZendFunctionEntryEx("parse_str", 0, ZifParseStr, ArginfoParseStr),
	zend.MakeZendFunctionEntryEx("str_getcsv", 0, ZifStrGetcsv, ArginfoStrGetcsv),
	zend.MakeZendFunctionEntryEx("str_pad", 0, ZifStrPad, ArginfoStrPad),
	zend.MakeZendFunctionEntryEx("chop", 0, ZifRtrim, ArginfoRtrim),
	zend.MakeZendFunctionEntryEx("strchr", 0, ZifStrstr, ArginfoStrstr),
	zend.MakeZendFunctionEntryEx("sprintf", 0, ZifUserSprintf, ArginfoSprintf),
	zend.MakeZendFunctionEntryEx("printf", 0, ZifUserPrintf, ArginfoPrintf),
	zend.MakeZendFunctionEntryEx("vprintf", 0, ZifVprintf, ArginfoVprintf),
	zend.MakeZendFunctionEntryEx("vsprintf", 0, ZifVsprintf, ArginfoVsprintf),
	zend.MakeZendFunctionEntryEx("fprintf", 0, ZifFprintf, ArginfoFprintf),
	zend.MakeZendFunctionEntryEx("vfprintf", 0, ZifVfprintf, ArginfoVfprintf),
	zend.MakeZendFunctionEntryEx("sscanf", 0, ZifSscanf, ArginfoSscanf),
	zend.MakeZendFunctionEntryEx("fscanf", 0, ZifFscanf, ArginfoFscanf),
	zend.MakeZendFunctionEntryEx("parse_url", 0, ZifParseUrl, ArginfoParseUrl),
	zend.MakeZendFunctionEntryEx("urlencode", 0, ZifUrlencode, ArginfoUrlencode),
	zend.MakeZendFunctionEntryEx("urldecode", 0, ZifUrldecode, ArginfoUrldecode),
	zend.MakeZendFunctionEntryEx("rawurlencode", 0, ZifRawurlencode, ArginfoRawurlencode),
	zend.MakeZendFunctionEntryEx("rawurldecode", 0, ZifRawurldecode, ArginfoRawurldecode),
	zend.MakeZendFunctionEntryEx("http_build_query", 0, ZifHttpBuildQuery, ArginfoHttpBuildQuery),
	zend.MakeZendFunctionEntryEx("readlink", 0, ZifReadlink, ArginfoReadlink),
	zend.MakeZendFunctionEntryEx("linkinfo", 0, ZifLinkinfo, ArginfoLinkinfo),
	zend.MakeZendFunctionEntryEx("symlink", 0, ZifSymlink, ArginfoSymlink),
	zend.MakeZendFunctionEntryEx("link", 0, ZifLink, ArginfoLink),
	zend.MakeZendFunctionEntryEx("unlink", 0, ZifUnlink, ArginfoUnlink),
	zend.MakeZendFunctionEntryEx("exec", 0, ZifExec, ArginfoExec),
	zend.MakeZendFunctionEntryEx("system", 0, ZifSystem, ArginfoSystem),
	zend.MakeZendFunctionEntryEx("escapeshellcmd", 0, ZifEscapeshellcmd, ArginfoEscapeshellcmd),
	zend.MakeZendFunctionEntryEx("escapeshellarg", 0, ZifEscapeshellarg, ArginfoEscapeshellarg),
	zend.MakeZendFunctionEntryEx("passthru", 0, ZifPassthru, ArginfoPassthru),
	zend.MakeZendFunctionEntryEx("shell_exec", 0, ZifShellExec, ArginfoShellExec),
	zend.MakeZendFunctionEntryEx("proc_open", 0, ZifProcOpen, ArginfoProcOpen),
	zend.MakeZendFunctionEntryEx("proc_close", 0, ZifProcClose, ArginfoProcClose),
	zend.MakeZendFunctionEntryEx("proc_terminate", 0, ZifProcTerminate, ArginfoProcTerminate),
	zend.MakeZendFunctionEntryEx("proc_get_status", 0, ZifProcGetStatus, ArginfoProcGetStatus),
	zend.MakeZendFunctionEntryEx("proc_nice", 0, ZifProcNice, ArginfoProcNice),
	zend.MakeZendFunctionEntryEx("rand", 0, ZifRand, ArginfoMtRand),
	zend.MakeZendFunctionEntryEx("srand", 0, ZifMtSrand, ArginfoMtSrand),
	zend.MakeZendFunctionEntryEx("getrandmax", 0, ZifMtGetrandmax, ArginfoMtGetrandmax),
	zend.MakeZendFunctionEntryEx("mt_rand", 0, ZifMtRand, ArginfoMtRand),
	zend.MakeZendFunctionEntryEx("mt_srand", 0, ZifMtSrand, ArginfoMtSrand),
	zend.MakeZendFunctionEntryEx("mt_getrandmax", 0, ZifMtGetrandmax, ArginfoMtGetrandmax),
	zend.MakeZendFunctionEntryEx("random_bytes", 0, ZifRandomBytes, ArginfoRandomBytes),
	zend.MakeZendFunctionEntryEx("random_int", 0, ZifRandomInt, ArginfoRandomInt),
	zend.MakeZendFunctionEntryEx("getservbyname", 0, ZifGetservbyname, ArginfoGetservbyname),
	zend.MakeZendFunctionEntryEx("getservbyport", 0, ZifGetservbyport, ArginfoGetservbyport),
	zend.MakeZendFunctionEntryEx("getprotobyname", 0, ZifGetprotobyname, ArginfoGetprotobyname),
	zend.MakeZendFunctionEntryEx("getprotobynumber", 0, ZifGetprotobynumber, ArginfoGetprotobynumber),
	zend.MakeZendFunctionEntryEx("getmyuid", 0, ZifGetmyuid, ArginfoGetmyuid),
	zend.MakeZendFunctionEntryEx("getmygid", 0, ZifGetmygid, ArginfoGetmygid),
	zend.MakeZendFunctionEntryEx("getmypid", 0, ZifGetmypid, ArginfoGetmypid),
	zend.MakeZendFunctionEntryEx("getmyinode", 0, ZifGetmyinode, ArginfoGetmyinode),
	zend.MakeZendFunctionEntryEx("getlastmod", 0, ZifGetlastmod, ArginfoGetlastmod),
	zend.MakeZendFunctionEntryEx("base64_decode", 0, ZifBase64Decode, ArginfoBase64Decode),
	zend.MakeZendFunctionEntryEx("base64_encode", 0, ZifBase64Encode, ArginfoBase64Encode),
	zend.MakeZendFunctionEntryEx("password_hash", 0, ZifPasswordHash, ArginfoPasswordHash),
	zend.MakeZendFunctionEntryEx("password_get_info", 0, ZifPasswordGetInfo, ArginfoPasswordGetInfo),
	zend.MakeZendFunctionEntryEx("password_needs_rehash", 0, ZifPasswordNeedsRehash, ArginfoPasswordNeedsRehash),
	zend.MakeZendFunctionEntryEx("password_verify", 0, ZifPasswordVerify, ArginfoPasswordVerify),
	zend.MakeZendFunctionEntryEx("password_algos", 0, ZifPasswordAlgos, ArginfoPasswordAlgos),
	zend.MakeZendFunctionEntryEx("convert_uuencode", 0, ZifConvertUuencode, ArginfoConvertUuencode),
	zend.MakeZendFunctionEntryEx("convert_uudecode", 0, ZifConvertUudecode, ArginfoConvertUudecode),
	zend.MakeZendFunctionEntryEx("abs", 0, ZifAbs, ArginfoAbs),
	zend.MakeZendFunctionEntryEx("ceil", 0, ZifCeil, ArginfoCeil),
	zend.MakeZendFunctionEntryEx("floor", 0, ZifFloor, ArginfoFloor),
	zend.MakeZendFunctionEntryEx("round", 0, ZifRound, ArginfoRound),
	zend.MakeZendFunctionEntryEx("sin", 0, ZifSin, ArginfoSin),
	zend.MakeZendFunctionEntryEx("cos", 0, ZifCos, ArginfoCos),
	zend.MakeZendFunctionEntryEx("tan", 0, ZifTan, ArginfoTan),
	zend.MakeZendFunctionEntryEx("asin", 0, ZifAsin, ArginfoAsin),
	zend.MakeZendFunctionEntryEx("acos", 0, ZifAcos, ArginfoAcos),
	zend.MakeZendFunctionEntryEx("atan", 0, ZifAtan, ArginfoAtan),
	zend.MakeZendFunctionEntryEx("atanh", 0, ZifAtanh, ArginfoAtanh),
	zend.MakeZendFunctionEntryEx("atan2", 0, ZifAtan2, ArginfoAtan2),
	zend.MakeZendFunctionEntryEx("sinh", 0, ZifSinh, ArginfoSinh),
	zend.MakeZendFunctionEntryEx("cosh", 0, ZifCosh, ArginfoCosh),
	zend.MakeZendFunctionEntryEx("tanh", 0, ZifTanh, ArginfoTanh),
	zend.MakeZendFunctionEntryEx("asinh", 0, ZifAsinh, ArginfoAsinh),
	zend.MakeZendFunctionEntryEx("acosh", 0, ZifAcosh, ArginfoAcosh),
	zend.MakeZendFunctionEntryEx("expm1", 0, ZifExpm1, ArginfoExpm1),
	zend.MakeZendFunctionEntryEx("log1p", 0, ZifLog1p, ArginfoLog1p),
	zend.MakeZendFunctionEntryEx("pi", 0, ZifPi, ArginfoPi),
	zend.MakeZendFunctionEntryEx("is_finite", 0, ZifIsFinite, ArginfoIsFinite),
	zend.MakeZendFunctionEntryEx("is_nan", 0, ZifIsNan, ArginfoIsNan),
	zend.MakeZendFunctionEntryEx("is_infinite", 0, ZifIsInfinite, ArginfoIsInfinite),
	zend.MakeZendFunctionEntryEx("pow", 0, ZifPow, ArginfoPow),
	zend.MakeZendFunctionEntryEx("exp", 0, ZifExp, ArginfoExp),
	zend.MakeZendFunctionEntryEx("log", 0, ZifLog, ArginfoLog),
	zend.MakeZendFunctionEntryEx("log10", 0, ZifLog10, ArginfoLog10),
	zend.MakeZendFunctionEntryEx("sqrt", 0, ZifSqrt, ArginfoSqrt),
	zend.MakeZendFunctionEntryEx("hypot", 0, ZifHypot, ArginfoHypot),
	zend.MakeZendFunctionEntryEx("deg2rad", 0, ZifDeg2rad, ArginfoDeg2rad),
	zend.MakeZendFunctionEntryEx("rad2deg", 0, ZifRad2deg, ArginfoRad2deg),
	zend.MakeZendFunctionEntryEx("bindec", 0, ZifBindec, ArginfoBindec),
	zend.MakeZendFunctionEntryEx("hexdec", 0, ZifHexdec, ArginfoHexdec),
	zend.MakeZendFunctionEntryEx("octdec", 0, ZifOctdec, ArginfoOctdec),
	zend.MakeZendFunctionEntryEx("decbin", 0, ZifDecbin, ArginfoDecbin),
	zend.MakeZendFunctionEntryEx("decoct", 0, ZifDecoct, ArginfoDecoct),
	zend.MakeZendFunctionEntryEx("dechex", 0, ZifDechex, ArginfoDechex),
	zend.MakeZendFunctionEntryEx("base_convert", 0, ZifBaseConvert, ArginfoBaseConvert),
	zend.MakeZendFunctionEntryEx("number_format", 0, ZifNumberFormat, ArginfoNumberFormat),
	zend.MakeZendFunctionEntryEx("fmod", 0, ZifFmod, ArginfoFmod),
	zend.MakeZendFunctionEntryEx("intdiv", 0, ZifIntdiv, ArginfoIntdiv),
	zend.MakeZendFunctionEntryEx("inet_ntop", 0, ZifInetNtop, ArginfoInetNtop),
	zend.MakeZendFunctionEntryEx("inet_pton", 0, PhpInetPton, ArginfoInetPton),
	zend.MakeZendFunctionEntryEx("ip2long", 0, ZifIp2long, ArginfoIp2long),
	zend.MakeZendFunctionEntryEx("long2ip", 0, ZifLong2ip, ArginfoLong2ip),
	zend.MakeZendFunctionEntryEx("getenv", 0, ZifGetenv, ArginfoGetenv),
	zend.MakeZendFunctionEntryEx("putenv", 0, ZifPutenv, ArginfoPutenv),
	zend.MakeZendFunctionEntryEx("getopt", 0, ZifGetopt, ArginfoGetopt),
	zend.MakeZendFunctionEntryEx("sys_getloadavg", 0, ZifSysGetloadavg, ArginfoSysGetloadavg),
	zend.MakeZendFunctionEntryEx("microtime", 0, ZifMicrotime, ArginfoMicrotime),
	zend.MakeZendFunctionEntryEx("gettimeofday", 0, ZifGettimeofday, ArginfoGettimeofday),
	zend.MakeZendFunctionEntryEx("getrusage", 0, ZifGetrusage, ArginfoGetrusage),
	zend.MakeZendFunctionEntryEx("hrtime", 0, ZifHrtime, ArginfoHrtime),
	zend.MakeZendFunctionEntryEx("uniqid", 0, ZifUniqid, ArginfoUniqid),
	zend.MakeZendFunctionEntryEx("quoted_printable_decode", 0, ZifQuotedPrintableDecode, ArginfoQuotedPrintableDecode),
	zend.MakeZendFunctionEntryEx("quoted_printable_encode", 0, ZifQuotedPrintableEncode, ArginfoQuotedPrintableEncode),
	zend.MakeZendFunctionEntryEx("convert_cyr_string", zend.ZEND_ACC_DEPRECATED, ZifConvertCyrString, ArginfoConvertCyrString),
	zend.MakeZendFunctionEntryEx("get_current_user", 0, ZifGetCurrentUser, ArginfoGetCurrentUser),
	zend.MakeZendFunctionEntryEx("set_time_limit", 0, ZifSetTimeLimit, ArginfoSetTimeLimit),
	zend.MakeZendFunctionEntryEx("header_register_callback", 0, ZifHeaderRegisterCallback, ArginfoHeaderRegisterCallback),
	zend.MakeZendFunctionEntryEx("get_cfg_var", 0, ZifGetCfgVar, ArginfoGetCfgVar),
	zend.MakeZendFunctionEntryEx("get_magic_quotes_gpc", zend.ZEND_ACC_DEPRECATED, ZifGetMagicQuotesGpc, ArginfoGetMagicQuotesGpc),
	zend.MakeZendFunctionEntryEx("get_magic_quotes_runtime", zend.ZEND_ACC_DEPRECATED, ZifGetMagicQuotesRuntime, ArginfoGetMagicQuotesRuntime),
	zend.MakeZendFunctionEntryEx("error_log", 0, ZifErrorLog, ArginfoErrorLog),
	zend.MakeZendFunctionEntryEx("error_get_last", 0, ZifErrorGetLast, ArginfoErrorGetLast),
	zend.MakeZendFunctionEntryEx("error_clear_last", 0, ZifErrorClearLast, ArginfoErrorClearLast),
	zend.MakeZendFunctionEntryEx("call_user_func", 0, ZifCallUserFunc, ArginfoCallUserFunc),
	zend.MakeZendFunctionEntryEx("call_user_func_array", 0, ZifCallUserFuncArray, ArginfoCallUserFuncArray),
	zend.MakeZendFunctionEntryEx("forward_static_call", 0, ZifForwardStaticCall, ArginfoForwardStaticCall),
	zend.MakeZendFunctionEntryEx("forward_static_call_array", 0, ZifForwardStaticCallArray, ArginfoForwardStaticCallArray),
	zend.MakeZendFunctionEntryEx("serialize", 0, ZifSerialize, ArginfoSerialize),
	zend.MakeZendFunctionEntryEx("unserialize", 0, ZifUnserialize, ArginfoUnserialize),
	zend.MakeZendFunctionEntryEx("var_dump", 0, ZifVarDump, ArginfoVarDump),
	zend.MakeZendFunctionEntryEx("var_export", 0, ZifVarExport, ArginfoVarExport),
	zend.MakeZendFunctionEntryEx("debug_zval_dump", 0, ZifDebugZvalDump, ArginfoDebugZvalDump),
	zend.MakeZendFunctionEntryEx("print_r", 0, ZifPrintR, ArginfoPrintR),
	zend.MakeZendFunctionEntryEx("memory_get_usage", 0, ZifMemoryGetUsage, ArginfoMemoryGetUsage),
	zend.MakeZendFunctionEntryEx("memory_get_peak_usage", 0, ZifMemoryGetPeakUsage, ArginfoMemoryGetPeakUsage),
	zend.MakeZendFunctionEntryEx("register_shutdown_function", 0, ZifRegisterShutdownFunction, ArginfoRegisterShutdownFunction),
	zend.MakeZendFunctionEntryEx("register_tick_function", 0, ZifRegisterTickFunction, ArginfoRegisterTickFunction),
	zend.MakeZendFunctionEntryEx("unregister_tick_function", 0, ZifUnregisterTickFunction, ArginfoUnregisterTickFunction),
	zend.MakeZendFunctionEntryEx("highlight_file", 0, ZifHighlightFile, ArginfoHighlightFile),
	zend.MakeZendFunctionEntryEx("show_source", 0, ZifHighlightFile, ArginfoHighlightFile),
	zend.MakeZendFunctionEntryEx("highlight_string", 0, ZifHighlightString, ArginfoHighlightString),
	zend.MakeZendFunctionEntryEx("php_strip_whitespace", 0, ZifPhpStripWhitespace, ArginfoPhpStripWhitespace),
	zend.MakeZendFunctionEntryEx("ini_get", 0, ZifIniGet, ArginfoIniGet),
	zend.MakeZendFunctionEntryEx("ini_get_all", 0, ZifIniGetAll, ArginfoIniGetAll),
	zend.MakeZendFunctionEntryEx("ini_set", 0, ZifIniSet, ArginfoIniSet),
	zend.MakeZendFunctionEntryEx("ini_alter", 0, ZifIniSet, ArginfoIniSet),
	zend.MakeZendFunctionEntryEx("ini_restore", 0, ZifIniRestore, ArginfoIniRestore),
	zend.MakeZendFunctionEntryEx("get_include_path", 0, ZifGetIncludePath, ArginfoGetIncludePath),
	zend.MakeZendFunctionEntryEx("set_include_path", 0, ZifSetIncludePath, ArginfoSetIncludePath),
	zend.MakeZendFunctionEntryEx("restore_include_path", zend.ZEND_ACC_DEPRECATED, ZifRestoreIncludePath, ArginfoRestoreIncludePath),
	zend.MakeZendFunctionEntryEx("setcookie", 0, ZifSetcookie, ArginfoSetcookie),
	zend.MakeZendFunctionEntryEx("setrawcookie", 0, ZifSetrawcookie, ArginfoSetrawcookie),
	zend.MakeZendFunctionEntryEx("header", 0, ZifHeader, ArginfoHeader),
	zend.MakeZendFunctionEntryEx("header_remove", 0, ZifHeaderRemove, ArginfoHeaderRemove),
	zend.MakeZendFunctionEntryEx("headers_sent", 0, ZifHeadersSent, ArginfoHeadersSent),
	zend.MakeZendFunctionEntryEx("headers_list", 0, ZifHeadersList, ArginfoHeadersList),
	zend.MakeZendFunctionEntryEx("http_response_code", 0, ZifHttpResponseCode, ArginfoHttpResponseCode),
	zend.MakeZendFunctionEntryEx("connection_aborted", 0, ZifConnectionAborted, ArginfoConnectionAborted),
	zend.MakeZendFunctionEntryEx("connection_status", 0, ZifConnectionStatus, ArginfoConnectionStatus),
	zend.MakeZendFunctionEntryEx("ignore_user_abort", 0, ZifIgnoreUserAbort, ArginfoIgnoreUserAbort),
	zend.MakeZendFunctionEntryEx("parse_ini_file", 0, ZifParseIniFile, ArginfoParseIniFile),
	zend.MakeZendFunctionEntryEx("parse_ini_string", 0, ZifParseIniString, ArginfoParseIniString),
	zend.MakeZendFunctionEntryEx("is_uploaded_file", 0, ZifIsUploadedFile, ArginfoIsUploadedFile),
	zend.MakeZendFunctionEntryEx("move_uploaded_file", 0, ZifMoveUploadedFile, ArginfoMoveUploadedFile),
	zend.MakeZendFunctionEntryEx("gethostbyaddr", 0, ZifGethostbyaddr, ArginfoGethostbyaddr),
	zend.MakeZendFunctionEntryEx("gethostbyname", 0, ZifGethostbyname, ArginfoGethostbyname),
	zend.MakeZendFunctionEntryEx("gethostbynamel", 0, ZifGethostbynamel, ArginfoGethostbynamel),
	zend.MakeZendFunctionEntryEx("gethostname", 0, ZifGethostname, ArginfoGethostname),
	zend.MakeZendFunctionEntryEx("net_get_interfaces", 0, ZifNetGetInterfaces, ArginfoNetGetInterfaces),
	zend.MakeZendFunctionEntryEx("dns_check_record", 0, ZifDnsCheckRecord, ArginfoDnsCheckRecord),
	zend.MakeZendFunctionEntryEx("checkdnsrr", 0, ZifDnsCheckRecord, ArginfoDnsCheckRecord),
	zend.MakeZendFunctionEntryEx("dns_get_mx", 0, ZifDnsGetMx, ArginfoDnsGetMx),
	zend.MakeZendFunctionEntryEx("getmxrr", 0, ZifDnsGetMx, ArginfoDnsGetMx),
	zend.MakeZendFunctionEntryEx("dns_get_record", 0, ZifDnsGetRecord, ArginfoDnsGetRecord),
	zend.MakeZendFunctionEntryEx("intval", 0, ZifIntval, ArginfoIntval),
	zend.MakeZendFunctionEntryEx("floatval", 0, ZifFloatval, ArginfoFloatval),
	zend.MakeZendFunctionEntryEx("doubleval", 0, ZifFloatval, ArginfoFloatval),
	zend.MakeZendFunctionEntryEx("strval", 0, ZifStrval, ArginfoStrval),
	zend.MakeZendFunctionEntryEx("boolval", 0, ZifBoolval, ArginfoBoolval),
	zend.MakeZendFunctionEntryEx("gettype", 0, ZifGettype, ArginfoGettype),
	zend.MakeZendFunctionEntryEx("settype", 0, ZifSettype, ArginfoSettype),
	zend.MakeZendFunctionEntryEx("is_null", 0, ZifIsNull, ArginfoIsNull),
	zend.MakeZendFunctionEntryEx("is_resource", 0, ZifIsResource, ArginfoIsResource),
	zend.MakeZendFunctionEntryEx("is_bool", 0, ZifIsBool, ArginfoIsBool),
	zend.MakeZendFunctionEntryEx("is_int", 0, ZifIsInt, ArginfoIsInt),
	zend.MakeZendFunctionEntryEx("is_float", 0, ZifIsFloat, ArginfoIsFloat),
	zend.MakeZendFunctionEntryEx("is_integer", 0, ZifIsInt, ArginfoIsInt),
	zend.MakeZendFunctionEntryEx("is_long", 0, ZifIsInt, ArginfoIsInt),
	zend.MakeZendFunctionEntryEx("is_double", 0, ZifIsFloat, ArginfoIsFloat),
	zend.MakeZendFunctionEntryEx("is_real", zend.ZEND_ACC_DEPRECATED, ZifIsFloat, ArginfoIsFloat),
	zend.MakeZendFunctionEntryEx("is_numeric", 0, ZifIsNumeric, ArginfoIsNumeric),
	zend.MakeZendFunctionEntryEx("is_string", 0, ZifIsString, ArginfoIsString),
	zend.MakeZendFunctionEntryEx("is_array", 0, ZifIsArray, ArginfoIsArray),
	zend.MakeZendFunctionEntryEx("is_object", 0, ZifIsObject, ArginfoIsObject),
	zend.MakeZendFunctionEntryEx("is_scalar", 0, ZifIsScalar, ArginfoIsScalar),
	zend.MakeZendFunctionEntryEx("is_callable", 0, ZifIsCallable, ArginfoIsCallable),
	zend.MakeZendFunctionEntryEx("is_iterable", 0, ZifIsIterable, ArginfoIsIterable),
	zend.MakeZendFunctionEntryEx("is_countable", 0, ZifIsCountable, ArginfoIsCountable),
	zend.MakeZendFunctionEntryEx("pclose", 0, ZifPclose, ArginfoPclose),
	zend.MakeZendFunctionEntryEx("popen", 0, ZifPopen, ArginfoPopen),
	zend.MakeZendFunctionEntryEx("readfile", 0, ZifReadfile, ArginfoReadfile),
	zend.MakeZendFunctionEntryEx("rewind", 0, ZifRewind, ArginfoRewind),
	zend.MakeZendFunctionEntryEx("rmdir", 0, ZifRmdir, ArginfoRmdir),
	zend.MakeZendFunctionEntryEx("umask", 0, ZifUmask, ArginfoUmask),
	zend.MakeZendFunctionEntryEx("fclose", 0, ZifFclose, ArginfoFclose),
	zend.MakeZendFunctionEntryEx("feof", 0, ZifFeof, ArginfoFeof),
	zend.MakeZendFunctionEntryEx("fgetc", 0, ZifFgetc, ArginfoFgetc),
	zend.MakeZendFunctionEntryEx("fgets", 0, ZifFgets, ArginfoFgets),
	zend.MakeZendFunctionEntryEx("fgetss", zend.ZEND_ACC_DEPRECATED, ZifFgetss, ArginfoFgetss),
	zend.MakeZendFunctionEntryEx("fread", 0, ZifFread, ArginfoFread),
	zend.MakeZendFunctionEntryEx("fopen", 0, PhpIfFopen, ArginfoFopen),
	zend.MakeZendFunctionEntryEx("fpassthru", 0, ZifFpassthru, ArginfoFpassthru),
	zend.MakeZendFunctionEntryEx("ftruncate", 0, PhpIfFtruncate, ArginfoFtruncate),
	zend.MakeZendFunctionEntryEx("fstat", 0, PhpIfFstat, ArginfoFstat),
	zend.MakeZendFunctionEntryEx("fseek", 0, ZifFseek, ArginfoFseek),
	zend.MakeZendFunctionEntryEx("ftell", 0, ZifFtell, ArginfoFtell),
	zend.MakeZendFunctionEntryEx("fflush", 0, ZifFflush, ArginfoFflush),
	zend.MakeZendFunctionEntryEx("fwrite", 0, ZifFwrite, ArginfoFwrite),
	zend.MakeZendFunctionEntryEx("fputs", 0, ZifFwrite, ArginfoFwrite),
	zend.MakeZendFunctionEntryEx("mkdir", 0, ZifMkdir, ArginfoMkdir),
	zend.MakeZendFunctionEntryEx("rename", 0, ZifRename, ArginfoRename),
	zend.MakeZendFunctionEntryEx("copy", 0, ZifCopy, ArginfoCopy),
	zend.MakeZendFunctionEntryEx("tempnam", 0, ZifTempnam, ArginfoTempnam),
	zend.MakeZendFunctionEntryEx("tmpfile", 0, PhpIfTmpfile, ArginfoTmpfile),
	zend.MakeZendFunctionEntryEx("file", 0, ZifFile, ArginfoFile),
	zend.MakeZendFunctionEntryEx("file_get_contents", 0, ZifFileGetContents, ArginfoFileGetContents),
	zend.MakeZendFunctionEntryEx("file_put_contents", 0, ZifFilePutContents, ArginfoFilePutContents),
	zend.MakeZendFunctionEntryEx("stream_select", 0, ZifStreamSelect, ArginfoStreamSelect),
	zend.MakeZendFunctionEntryEx("stream_context_create", 0, ZifStreamContextCreate, ArginfoStreamContextCreate),
	zend.MakeZendFunctionEntryEx("stream_context_set_params", 0, ZifStreamContextSetParams, ArginfoStreamContextSetParams),
	zend.MakeZendFunctionEntryEx("stream_context_get_params", 0, ZifStreamContextGetParams, ArginfoStreamContextGetParams),
	zend.MakeZendFunctionEntryEx("stream_context_set_option", 0, ZifStreamContextSetOption, ArginfoStreamContextSetOption),
	zend.MakeZendFunctionEntryEx("stream_context_get_options", 0, ZifStreamContextGetOptions, ArginfoStreamContextGetOptions),
	zend.MakeZendFunctionEntryEx("stream_context_get_default", 0, ZifStreamContextGetDefault, ArginfoStreamContextGetDefault),
	zend.MakeZendFunctionEntryEx("stream_context_set_default", 0, ZifStreamContextSetDefault, ArginfoStreamContextSetDefault),
	zend.MakeZendFunctionEntryEx("stream_filter_prepend", 0, ZifStreamFilterPrepend, ArginfoStreamFilterPrepend),
	zend.MakeZendFunctionEntryEx("stream_filter_append", 0, ZifStreamFilterAppend, ArginfoStreamFilterAppend),
	zend.MakeZendFunctionEntryEx("stream_filter_remove", 0, ZifStreamFilterRemove, ArginfoStreamFilterRemove),
	zend.MakeZendFunctionEntryEx("stream_socket_client", 0, ZifStreamSocketClient, ArginfoStreamSocketClient),
	zend.MakeZendFunctionEntryEx("stream_socket_server", 0, ZifStreamSocketServer, ArginfoStreamSocketServer),
	zend.MakeZendFunctionEntryEx("stream_socket_accept", 0, ZifStreamSocketAccept, ArginfoStreamSocketAccept),
	zend.MakeZendFunctionEntryEx("stream_socket_get_name", 0, ZifStreamSocketGetName, ArginfoStreamSocketGetName),
	zend.MakeZendFunctionEntryEx("stream_socket_recvfrom", 0, ZifStreamSocketRecvfrom, ArginfoStreamSocketRecvfrom),
	zend.MakeZendFunctionEntryEx("stream_socket_sendto", 0, ZifStreamSocketSendto, ArginfoStreamSocketSendto),
	zend.MakeZendFunctionEntryEx("stream_socket_enable_crypto", 0, ZifStreamSocketEnableCrypto, ArginfoStreamSocketEnableCrypto),
	zend.MakeZendFunctionEntryEx("stream_socket_shutdown", 0, ZifStreamSocketShutdown, ArginfoStreamSocketShutdown),
	zend.MakeZendFunctionEntryEx("stream_socket_pair", 0, ZifStreamSocketPair, ArginfoStreamSocketPair),
	zend.MakeZendFunctionEntryEx("stream_copy_to_stream", 0, ZifStreamCopyToStream, ArginfoStreamCopyToStream),
	zend.MakeZendFunctionEntryEx("stream_get_contents", 0, ZifStreamGetContents, ArginfoStreamGetContents),
	zend.MakeZendFunctionEntryEx("stream_supports_lock", 0, ZifStreamSupportsLock, ArginfoStreamSupportsLock),
	zend.MakeZendFunctionEntryEx("stream_isatty", 0, ZifStreamIsatty, ArginfoStreamIsatty),
	zend.MakeZendFunctionEntryEx("fgetcsv", 0, ZifFgetcsv, ArginfoFgetcsv),
	zend.MakeZendFunctionEntryEx("fputcsv", 0, ZifFputcsv, ArginfoFputcsv),
	zend.MakeZendFunctionEntryEx("flock", 0, ZifFlock, ArginfoFlock),
	zend.MakeZendFunctionEntryEx("get_meta_tags", 0, ZifGetMetaTags, ArginfoGetMetaTags),
	zend.MakeZendFunctionEntryEx("stream_set_read_buffer", 0, ZifStreamSetReadBuffer, ArginfoStreamSetReadBuffer),
	zend.MakeZendFunctionEntryEx("stream_set_write_buffer", 0, ZifStreamSetWriteBuffer, ArginfoStreamSetWriteBuffer),
	zend.MakeZendFunctionEntryEx("set_file_buffer", 0, ZifStreamSetWriteBuffer, ArginfoStreamSetWriteBuffer),
	zend.MakeZendFunctionEntryEx("stream_set_chunk_size", 0, ZifStreamSetChunkSize, ArginfoStreamSetChunkSize),
	zend.MakeZendFunctionEntryEx("stream_set_blocking", 0, ZifStreamSetBlocking, ArginfoStreamSetBlocking),
	zend.MakeZendFunctionEntryEx("socket_set_blocking", 0, ZifStreamSetBlocking, ArginfoStreamSetBlocking),
	zend.MakeZendFunctionEntryEx("stream_get_meta_data", 0, ZifStreamGetMetaData, ArginfoStreamGetMetaData),
	zend.MakeZendFunctionEntryEx("stream_get_line", 0, ZifStreamGetLine, ArginfoStreamGetLine),
	zend.MakeZendFunctionEntryEx("stream_wrapper_register", 0, ZifStreamWrapperRegister, ArginfoStreamWrapperRegister),
	zend.MakeZendFunctionEntryEx("stream_register_wrapper", 0, ZifStreamWrapperRegister, ArginfoStreamWrapperRegister),
	zend.MakeZendFunctionEntryEx("stream_wrapper_unregister", 0, ZifStreamWrapperUnregister, ArginfoStreamWrapperUnregister),
	zend.MakeZendFunctionEntryEx("stream_wrapper_restore", 0, ZifStreamWrapperRestore, ArginfoStreamWrapperRestore),
	zend.MakeZendFunctionEntryEx("stream_get_wrappers", 0, ZifStreamGetWrappers, ArginfoStreamGetWrappers),
	zend.MakeZendFunctionEntryEx("stream_get_transports", 0, ZifStreamGetTransports, ArginfoStreamGetTransports),
	zend.MakeZendFunctionEntryEx("stream_resolve_include_path", 0, ZifStreamResolveIncludePath, ArginfoStreamResolveIncludePath),
	zend.MakeZendFunctionEntryEx("stream_is_local", 0, ZifStreamIsLocal, ArginfoStreamIsLocal),
	zend.MakeZendFunctionEntryEx("get_headers", 0, ZifGetHeaders, ArginfoGetHeaders),
	zend.MakeZendFunctionEntryEx("stream_set_timeout", 0, ZifStreamSetTimeout, ArginfoStreamSetTimeout),
	zend.MakeZendFunctionEntryEx("socket_set_timeout", 0, ZifStreamSetTimeout, ArginfoStreamSetTimeout),
	zend.MakeZendFunctionEntryEx("socket_get_status", 0, ZifStreamGetMetaData, ArginfoStreamGetMetaData),
	zend.MakeZendFunctionEntryEx("realpath", 0, ZifRealpath, ArginfoRealpath),
	zend.MakeZendFunctionEntryEx("fnmatch", 0, ZifFnmatch, ArginfoFnmatch),
	zend.MakeZendFunctionEntryEx("fsockopen", 0, ZifFsockopen, ArginfoFsockopen),
	zend.MakeZendFunctionEntryEx("pfsockopen", 0, ZifPfsockopen, ArginfoPfsockopen),
	zend.MakeZendFunctionEntryEx("pack", 0, ZifPack, ArginfoPack),
	zend.MakeZendFunctionEntryEx("unpack", 0, ZifUnpack, ArginfoUnpack),
	zend.MakeZendFunctionEntryEx("get_browser", 0, ZifGetBrowser, ArginfoGetBrowser),
	zend.MakeZendFunctionEntryEx("crypt", 0, ZifCrypt, ArginfoCrypt),
	zend.MakeZendFunctionEntryEx("opendir", 0, ZifOpendir, ArginfoOpendir),
	zend.MakeZendFunctionEntryEx("closedir", 0, ZifClosedir, ArginfoClosedir),
	zend.MakeZendFunctionEntryEx("chdir", 0, ZifChdir, ArginfoChdir),
	zend.MakeZendFunctionEntryEx("chroot", 0, ZifChroot, ArginfoChroot),
	zend.MakeZendFunctionEntryEx("getcwd", 0, ZifGetcwd, ArginfoGetcwd),
	zend.MakeZendFunctionEntryEx("rewinddir", 0, ZifRewinddir, ArginfoRewinddir),
	zend.MakeZendFunctionEntryEx("readdir", 0, PhpIfReaddir, ArginfoReaddir),
	zend.MakeZendFunctionEntryEx("dir", 0, ZifGetdir, ArginfoDir),
	zend.MakeZendFunctionEntryEx("scandir", 0, ZifScandir, ArginfoScandir),
	zend.MakeZendFunctionEntryEx("glob", 0, ZifGlob, ArginfoGlob),
	zend.MakeZendFunctionEntryEx("fileatime", 0, ZifFileatime, ArginfoFileatime),
	zend.MakeZendFunctionEntryEx("filectime", 0, ZifFilectime, ArginfoFilectime),
	zend.MakeZendFunctionEntryEx("filegroup", 0, ZifFilegroup, ArginfoFilegroup),
	zend.MakeZendFunctionEntryEx("fileinode", 0, ZifFileinode, ArginfoFileinode),
	zend.MakeZendFunctionEntryEx("filemtime", 0, ZifFilemtime, ArginfoFilemtime),
	zend.MakeZendFunctionEntryEx("fileowner", 0, ZifFileowner, ArginfoFileowner),
	zend.MakeZendFunctionEntryEx("fileperms", 0, ZifFileperms, ArginfoFileperms),
	zend.MakeZendFunctionEntryEx("filesize", 0, ZifFilesize, ArginfoFilesize),
	zend.MakeZendFunctionEntryEx("filetype", 0, ZifFiletype, ArginfoFiletype),
	zend.MakeZendFunctionEntryEx("file_exists", 0, ZifFileExists, ArginfoFileExists),
	zend.MakeZendFunctionEntryEx("is_writable", 0, ZifIsWritable, ArginfoIsWritable),
	zend.MakeZendFunctionEntryEx("is_writeable", 0, ZifIsWritable, ArginfoIsWritable),
	zend.MakeZendFunctionEntryEx("is_readable", 0, ZifIsReadable, ArginfoIsReadable),
	zend.MakeZendFunctionEntryEx("is_executable", 0, ZifIsExecutable, ArginfoIsExecutable),
	zend.MakeZendFunctionEntryEx("is_file", 0, ZifIsFile, ArginfoIsFile),
	zend.MakeZendFunctionEntryEx("is_dir", 0, ZifIsDir, ArginfoIsDir),
	zend.MakeZendFunctionEntryEx("is_link", 0, ZifIsLink, ArginfoIsLink),
	zend.MakeZendFunctionEntryEx("stat", 0, PhpIfStat, ArginfoStat),
	zend.MakeZendFunctionEntryEx("lstat", 0, PhpIfLstat, ArginfoLstat),
	zend.MakeZendFunctionEntryEx("chown", 0, ZifChown, ArginfoChown),
	zend.MakeZendFunctionEntryEx("chgrp", 0, ZifChgrp, ArginfoChgrp),
	zend.MakeZendFunctionEntryEx("lchown", 0, ZifLchown, ArginfoLchown),
	zend.MakeZendFunctionEntryEx("lchgrp", 0, ZifLchgrp, ArginfoLchgrp),
	zend.MakeZendFunctionEntryEx("chmod", 0, ZifChmod, ArginfoChmod),
	zend.MakeZendFunctionEntryEx("touch", 0, ZifTouch, ArginfoTouch),
	zend.MakeZendFunctionEntryEx("clearstatcache", 0, ZifClearstatcache, ArginfoClearstatcache),
	zend.MakeZendFunctionEntryEx("disk_total_space", 0, ZifDiskTotalSpace, ArginfoDiskTotalSpace),
	zend.MakeZendFunctionEntryEx("disk_free_space", 0, ZifDiskFreeSpace, ArginfoDiskFreeSpace),
	zend.MakeZendFunctionEntryEx("diskfreespace", 0, ZifDiskFreeSpace, ArginfoDiskFreeSpace),
	zend.MakeZendFunctionEntryEx("realpath_cache_size", 0, ZifRealpathCacheSize, ArginfoRealpathCacheSize),
	zend.MakeZendFunctionEntryEx("realpath_cache_get", 0, ZifRealpathCacheGet, ArginfoRealpathCacheGet),
	zend.MakeZendFunctionEntryEx("mail", 0, ZifMail, ArginfoMail),
	zend.MakeZendFunctionEntryEx("ezmlm_hash", zend.ZEND_ACC_DEPRECATED, ZifEzmlmHash, ArginfoEzmlmHash),
	zend.MakeZendFunctionEntryEx("openlog", 0, ZifOpenlog, ArginfoOpenlog),
	zend.MakeZendFunctionEntryEx("syslog", 0, ZifSyslog, ArginfoSyslog),
	zend.MakeZendFunctionEntryEx("closelog", 0, ZifCloselog, ArginfoCloselog),
	zend.MakeZendFunctionEntryEx("lcg_value", 0, ZifLcgValue, ArginfoLcgValue),
	zend.MakeZendFunctionEntryEx("metaphone", 0, ZifMetaphone, ArginfoMetaphone),
	zend.MakeZendFunctionEntryEx("ob_start", 0, core.ZifObStart, ArginfoObStart),
	zend.MakeZendFunctionEntryEx("ob_flush", 0, core.ZifObFlush, ArginfoObFlush),
	zend.MakeZendFunctionEntryEx("ob_clean", 0, core.ZifObClean, ArginfoObClean),
	zend.MakeZendFunctionEntryEx("ob_end_flush", 0, core.ZifObEndFlush, ArginfoObEndFlush),
	zend.MakeZendFunctionEntryEx("ob_end_clean", 0, core.ZifObEndClean, ArginfoObEndClean),
	zend.MakeZendFunctionEntryEx("ob_get_flush", 0, core.ZifObGetFlush, ArginfoObGetFlush),
	zend.MakeZendFunctionEntryEx("ob_get_clean", 0, core.ZifObGetClean, ArginfoObGetClean),
	zend.MakeZendFunctionEntryEx("ob_get_length", 0, core.ZifObGetLength, ArginfoObGetLength),
	zend.MakeZendFunctionEntryEx("ob_get_level", 0, core.ZifObGetLevel, ArginfoObGetLevel),
	zend.MakeZendFunctionEntryEx("ob_get_status", 0, core.ZifObGetStatus, ArginfoObGetStatus),
	zend.MakeZendFunctionEntryEx("ob_get_contents", 0, core.ZifObGetContents, ArginfoObGetContents),
	zend.MakeZendFunctionEntryEx("ob_implicit_flush", 0, core.ZifObImplicitFlush, ArginfoObImplicitFlush),
	zend.MakeZendFunctionEntryEx("ob_list_handlers", 0, core.ZifObListHandlers, ArginfoObListHandlers),
	zend.MakeZendFunctionEntryEx("ksort", 0, ZifKsort, ArginfoKsort),
	zend.MakeZendFunctionEntryEx("krsort", 0, ZifKrsort, ArginfoKrsort),
	zend.MakeZendFunctionEntryEx("natsort", 0, ZifNatsort, ArginfoNatsort),
	zend.MakeZendFunctionEntryEx("natcasesort", 0, ZifNatcasesort, ArginfoNatcasesort),
	zend.MakeZendFunctionEntryEx("asort", 0, ZifAsort, ArginfoAsort),
	zend.MakeZendFunctionEntryEx("arsort", 0, ZifArsort, ArginfoArsort),
	zend.MakeZendFunctionEntryEx("sort", 0, ZifSort, ArginfoSort),
	zend.MakeZendFunctionEntryEx("rsort", 0, ZifRsort, ArginfoRsort),
	zend.MakeZendFunctionEntryEx("usort", 0, ZifUsort, ArginfoUsort),
	zend.MakeZendFunctionEntryEx("uasort", 0, ZifUasort, ArginfoUasort),
	zend.MakeZendFunctionEntryEx("uksort", 0, ZifUksort, ArginfoUksort),
	zend.MakeZendFunctionEntryEx("shuffle", 0, ZifShuffle, ArginfoShuffle),
	zend.MakeZendFunctionEntryEx("array_walk", 0, ZifArrayWalk, ArginfoArrayWalk),
	zend.MakeZendFunctionEntryEx("array_walk_recursive", 0, ZifArrayWalkRecursive, ArginfoArrayWalkRecursive),
	zend.MakeZendFunctionEntryEx("count", 0, ZifCount, ArginfoCount),
	zend.MakeZendFunctionEntryEx("end", 0, ZifEnd, ArginfoEnd),
	zend.MakeZendFunctionEntryEx("prev", 0, ZifPrev, ArginfoPrev),
	zend.MakeZendFunctionEntryEx("next", 0, ZifNext, ArginfoNext),
	zend.MakeZendFunctionEntryEx("reset", 0, ZifReset, ArginfoReset),
	zend.MakeZendFunctionEntryEx("current", 0, ZifCurrent, ArginfoCurrent),
	zend.MakeZendFunctionEntryEx("key", 0, ZifKey, ArginfoKey),
	zend.MakeZendFunctionEntryEx("min", 0, ZifMin, ArginfoMin),
	zend.MakeZendFunctionEntryEx("max", 0, ZifMax, ArginfoMax),
	zend.MakeZendFunctionEntryEx("in_array", 0, ZifInArray, ArginfoInArray),
	zend.MakeZendFunctionEntryEx("array_search", 0, ZifArraySearch, ArginfoArraySearch),
	zend.MakeZendFunctionEntryEx("extract", 0, ZifExtract, ArginfoExtract),
	zend.MakeZendFunctionEntryEx("compact", 0, ZifCompact, ArginfoCompact),
	zend.MakeZendFunctionEntryEx("array_fill", 0, ZifArrayFill, ArginfoArrayFill),
	zend.MakeZendFunctionEntryEx("array_fill_keys", 0, ZifArrayFillKeys, ArginfoArrayFillKeys),
	zend.MakeZendFunctionEntryEx("range", 0, ZifRange, ArginfoRange),
	zend.MakeZendFunctionEntryEx("array_multisort", 0, ZifArrayMultisort, ArginfoArrayMultisort),
	zend.MakeZendFunctionEntryEx("array_push", 0, ZifArrayPush, ArginfoArrayPush),
	zend.MakeZendFunctionEntryEx("array_pop", 0, ZifArrayPop, ArginfoArrayPop),
	zend.MakeZendFunctionEntryEx("array_shift", 0, ZifArrayShift, ArginfoArrayShift),
	zend.MakeZendFunctionEntryEx("array_unshift", 0, ZifArrayUnshift, ArginfoArrayUnshift),
	zend.MakeZendFunctionEntryEx("array_splice", 0, ZifArraySplice, ArginfoArraySplice),
	zend.MakeZendFunctionEntryEx("array_slice", 0, ZifArraySlice, ArginfoArraySlice),
	zend.MakeZendFunctionEntryEx("array_merge", 0, ZifArrayMerge, ArginfoArrayMerge),
	zend.MakeZendFunctionEntryEx("array_merge_recursive", 0, ZifArrayMergeRecursive, ArginfoArrayMergeRecursive),
	zend.MakeZendFunctionEntryEx("array_replace", 0, ZifArrayReplace, ArginfoArrayReplace),
	zend.MakeZendFunctionEntryEx("array_replace_recursive", 0, ZifArrayReplaceRecursive, ArginfoArrayReplaceRecursive),
	zend.MakeZendFunctionEntryEx("array_keys", 0, ZifArrayKeys, ArginfoArrayKeys),
	zend.MakeZendFunctionEntryEx("array_key_first", 0, ZifArrayKeyFirst, ArginfoArrayKeyFirst),
	zend.MakeZendFunctionEntryEx("array_key_last", 0, ZifArrayKeyLast, ArginfoArrayKeyLast),
	zend.MakeZendFunctionEntryEx("array_values", 0, ZifArrayValues, ArginfoArrayValues),
	zend.MakeZendFunctionEntryEx("array_count_values", 0, ZifArrayCountValues, ArginfoArrayCountValues),
	zend.MakeZendFunctionEntryEx("array_column", 0, ZifArrayColumn, ArginfoArrayColumn),
	zend.MakeZendFunctionEntryEx("array_reverse", 0, ZifArrayReverse, ArginfoArrayReverse),
	zend.MakeZendFunctionEntryEx("array_reduce", 0, ZifArrayReduce, ArginfoArrayReduce),
	zend.MakeZendFunctionEntryEx("array_pad", 0, ZifArrayPad, ArginfoArrayPad),
	zend.MakeZendFunctionEntryEx("array_flip", 0, ZifArrayFlip, ArginfoArrayFlip),
	zend.MakeZendFunctionEntryEx("array_change_key_case", 0, ZifArrayChangeKeyCase, ArginfoArrayChangeKeyCase),
	zend.MakeZendFunctionEntryEx("array_rand", 0, ZifArrayRand, ArginfoArrayRand),
	zend.MakeZendFunctionEntryEx("array_unique", 0, ZifArrayUnique, ArginfoArrayUnique),
	zend.MakeZendFunctionEntryEx("array_intersect", 0, ZifArrayIntersect, ArginfoArrayIntersect),
	zend.MakeZendFunctionEntryEx("array_intersect_key", 0, ZifArrayIntersectKey, ArginfoArrayIntersectKey),
	zend.MakeZendFunctionEntryEx("array_intersect_ukey", 0, ZifArrayIntersectUkey, ArginfoArrayIntersectUkey),
	zend.MakeZendFunctionEntryEx("array_uintersect", 0, ZifArrayUintersect, ArginfoArrayUintersect),
	zend.MakeZendFunctionEntryEx("array_intersect_assoc", 0, ZifArrayIntersectAssoc, ArginfoArrayIntersectAssoc),
	zend.MakeZendFunctionEntryEx("array_uintersect_assoc", 0, ZifArrayUintersectAssoc, ArginfoArrayUintersectAssoc),
	zend.MakeZendFunctionEntryEx("array_intersect_uassoc", 0, ZifArrayIntersectUassoc, ArginfoArrayIntersectUassoc),
	zend.MakeZendFunctionEntryEx("array_uintersect_uassoc", 0, ZifArrayUintersectUassoc, ArginfoArrayUintersectUassoc),
	zend.MakeZendFunctionEntryEx("array_diff", 0, ZifArrayDiff, ArginfoArrayDiff),
	zend.MakeZendFunctionEntryEx("array_diff_key", 0, ZifArrayDiffKey, ArginfoArrayDiffKey),
	zend.MakeZendFunctionEntryEx("array_diff_ukey", 0, ZifArrayDiffUkey, ArginfoArrayDiffUkey),
	zend.MakeZendFunctionEntryEx("array_udiff", 0, ZifArrayUdiff, ArginfoArrayUdiff),
	zend.MakeZendFunctionEntryEx("array_diff_assoc", 0, ZifArrayDiffAssoc, ArginfoArrayDiffAssoc),
	zend.MakeZendFunctionEntryEx("array_udiff_assoc", 0, ZifArrayUdiffAssoc, ArginfoArrayUdiffAssoc),
	zend.MakeZendFunctionEntryEx("array_diff_uassoc", 0, ZifArrayDiffUassoc, ArginfoArrayDiffUassoc),
	zend.MakeZendFunctionEntryEx("array_udiff_uassoc", 0, ZifArrayUdiffUassoc, ArginfoArrayUdiffUassoc),
	zend.MakeZendFunctionEntryEx("array_sum", 0, ZifArraySum, ArginfoArraySum),
	zend.MakeZendFunctionEntryEx("array_product", 0, ZifArrayProduct, ArginfoArrayProduct),
	zend.MakeZendFunctionEntryEx("array_filter", 0, ZifArrayFilter, ArginfoArrayFilter),
	zend.MakeZendFunctionEntryEx("array_map", 0, ZifArrayMap, ArginfoArrayMap),
	zend.MakeZendFunctionEntryEx("array_chunk", 0, ZifArrayChunk, ArginfoArrayChunk),
	zend.MakeZendFunctionEntryEx("array_combine", 0, ZifArrayCombine, ArginfoArrayCombine),
	zend.MakeZendFunctionEntryEx("array_key_exists", 0, ZifArrayKeyExists, ArginfoArrayKeyExists),
	zend.MakeZendFunctionEntryEx("pos", 0, ZifCurrent, ArginfoCurrent),
	zend.MakeZendFunctionEntryEx("sizeof", 0, ZifCount, ArginfoCount),
	zend.MakeZendFunctionEntryEx("key_exists", 0, ZifArrayKeyExists, ArginfoArrayKeyExists),
	zend.MakeZendFunctionEntryEx("assert", 0, ZifAssert, ArginfoAssert),
	zend.MakeZendFunctionEntryEx("assert_options", 0, ZifAssertOptions, ArginfoAssertOptions),
	zend.MakeZendFunctionEntryEx("version_compare", 0, ZifVersionCompare, ArginfoVersionCompare),
	zend.MakeZendFunctionEntryEx("ftok", 0, ZifFtok, ArginfoFtok),
	zend.MakeZendFunctionEntryEx("str_rot13", 0, ZifStrRot13, ArginfoStrRot13),
	zend.MakeZendFunctionEntryEx("stream_get_filters", 0, ZifStreamGetFilters, ArginfoStreamGetFilters),
	zend.MakeZendFunctionEntryEx("stream_filter_register", 0, ZifStreamFilterRegister, ArginfoStreamFilterRegister),
	zend.MakeZendFunctionEntryEx("stream_bucket_make_writeable", 0, ZifStreamBucketMakeWriteable, ArginfoStreamBucketMakeWriteable),
	zend.MakeZendFunctionEntryEx("stream_bucket_prepend", 0, ZifStreamBucketPrepend, ArginfoStreamBucketPrepend),
	zend.MakeZendFunctionEntryEx("stream_bucket_append", 0, ZifStreamBucketAppend, ArginfoStreamBucketAppend),
	zend.MakeZendFunctionEntryEx("stream_bucket_new", 0, ZifStreamBucketNew, ArginfoStreamBucketNew),
	zend.MakeZendFunctionEntryEx("output_add_rewrite_var", 0, core.ZifOutputAddRewriteVar, ArginfoOutputAddRewriteVar),
	zend.MakeZendFunctionEntryEx("output_reset_rewrite_vars", 0, core.ZifOutputResetRewriteVars, ArginfoOutputResetRewriteVars),
	zend.MakeZendFunctionEntryEx("sys_get_temp_dir", 0, ZifSysGetTempDir, ArginfoSysGetTempDir),
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
