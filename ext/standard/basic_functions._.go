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
