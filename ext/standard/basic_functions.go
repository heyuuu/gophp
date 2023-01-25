// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/sapi/cli"
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

// #define BASIC_FUNCTIONS_H

// # include < sys / stat . h >

// # include < wchar . h >

// # include "php_filestat.h"

// # include "zend_highlight.h"

// # include "url_scanner_ex.h"

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

func BG(v *uint32) __auto__ { return BasicGlobals.v }

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

// # include "php.h"

// # include "php_streams.h"

// # include "php_main.h"

// # include "php_globals.h"

// # include "php_variables.h"

// # include "php_ini.h"

// # include "php_standard.h"

// # include "php_math.h"

// # include "php_http.h"

// # include "php_incomplete_class.h"

// # include "php_getopt.h"

// # include "ext/standard/info.h"

// failed # include "ext/session/php_session.h"

// # include "zend_operators.h"

// # include "ext/standard/php_dns.h"

// # include "ext/standard/php_uuencode.h"

// # include "ext/standard/php_mt_rand.h"

type YY_BUFFER_STATE *__struct__yy_buffer_state

// # include "zend.h"

// # include "zend_ini_scanner.h"

// # include "zend_language_scanner.h"

// # include < zend_language_parser . h >

// # include "zend_portability.h"

// # include < stdarg . h >

// # include < stdlib . h >

// # include < math . h >

// # include < time . h >

// # include < stdio . h >

// # include < sys / types . h >

// # include < sys / stat . h >

// # include < netdb . h >

// # include < arpa / inet . h >

// # include < unistd . h >

// # include < string . h >

// # include < locale . h >

// # include < sys / mman . h >

const INADDR_NONE = zend_ulong - 1

// # include "zend_globals.h"

// # include "php_globals.h"

// # include "SAPI.h"

// # include "php_ticks.h"

// # include "php_fopen_wrappers.h"

// # include "streamsfuncs.h"

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

func PhpPutenvDestructor(zv *zend.Zval) {
	var pe *PutenvEntry = zend.Z_PTR_P(zv)
	if pe.GetPreviousValue() != nil {
		putenv(pe.GetPreviousValue())
	} else {
		unsetenv(pe.GetKey())
	}

	/* don't forget to reset the various libc globals that
	 * we might have changed by an earlier call to tzset(). */

	if !(strncmp(pe.GetKey(), "TZ", pe.GetKeyLen())) {
		tzset()
	}
	zend.Efree(pe.GetPutenvString())
	zend.Efree(pe.GetKey())
	zend.Efree(pe)
}

/* }}} */

func BasicGlobalsCtor(basic_globals_p *PhpBasicGlobals) {
	BG(mt_rand_is_seeded) = 0
	BG(mt_rand_mode) = MT_RAND_MT19937
	BG(umask) = -1
	BG(next) = nil
	BG(left) = -1
	BG(user_tick_functions) = nil
	BG(user_filter_map) = nil
	BG(serialize_lock) = 0
	memset(&BG(serialize), 0, b.SizeOf("BG ( serialize )"))
	memset(&BG(unserialize), 0, b.SizeOf("BG ( unserialize )"))
	memset(&BG(url_adapt_session_ex), 0, b.SizeOf("BG ( url_adapt_session_ex )"))
	memset(&BG(url_adapt_output_ex), 0, b.SizeOf("BG ( url_adapt_output_ex )"))
	BG(url_adapt_session_ex).type_ = 1
	BG(url_adapt_output_ex).type_ = 0
	zend.ZendHashInit(&BG(url_adapt_session_hosts_ht), 0, nil, nil, 1)
	zend.ZendHashInit(&BG(url_adapt_output_hosts_ht), 0, nil, nil, 1)
	BG(incomplete_class) = IncompleteClassEntry
	BG(page_uid) = -1
	BG(page_gid) = -1
}

/* }}} */

func BasicGlobalsDtor(basic_globals_p *PhpBasicGlobals) {
	if basic_globals_p.GetUrlAdaptSessionEx().GetTags() != nil {
		zend.ZendHashDestroy(basic_globals_p.GetUrlAdaptSessionEx().GetTags())
		zend.Free(basic_globals_p.GetUrlAdaptSessionEx().GetTags())
	}
	if basic_globals_p.GetUrlAdaptOutputEx().GetTags() != nil {
		zend.ZendHashDestroy(basic_globals_p.GetUrlAdaptOutputEx().GetTags())
		zend.Free(basic_globals_p.GetUrlAdaptOutputEx().GetTags())
	}
	zend.ZendHashDestroy(&basic_globals_p.url_adapt_session_hosts_ht)
	zend.ZendHashDestroy(&basic_globals_p.url_adapt_output_hosts_ht)
}

/* }}} */

func PhpGetNan() float64 { return zend.ZEND_NAN }

/* }}} */

func PhpGetInf() float64 { return zend.ZEND_INFINITY }

/* }}} */

// #define BASIC_MINIT_SUBMODULE(module) if ( PHP_MINIT ( module ) ( INIT_FUNC_ARGS_PASSTHRU ) != SUCCESS ) { return FAILURE ; }

// #define BASIC_RINIT_SUBMODULE(module) PHP_RINIT ( module ) ( INIT_FUNC_ARGS_PASSTHRU ) ;

// #define BASIC_MINFO_SUBMODULE(module) PHP_MINFO ( module ) ( ZEND_MODULE_INFO_FUNC_ARGS_PASSTHRU ) ;

// #define BASIC_RSHUTDOWN_SUBMODULE(module) PHP_RSHUTDOWN ( module ) ( SHUTDOWN_FUNC_ARGS_PASSTHRU ) ;

// #define BASIC_MSHUTDOWN_SUBMODULE(module) PHP_MSHUTDOWN ( module ) ( SHUTDOWN_FUNC_ARGS_PASSTHRU ) ;

func ZmStartupBasic(type_ int, module_number int) int {
	BasicGlobalsCtor(&BasicGlobals)
	IncompleteClassEntry = PhpCreateIncompleteClass()
	BG(incomplete_class) = IncompleteClassEntry
	zend.REGISTER_LONG_CONSTANT("CONNECTION_ABORTED", core.PHP_CONNECTION_ABORTED, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("CONNECTION_NORMAL", core.PHP_CONNECTION_NORMAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("CONNECTION_TIMEOUT", core.PHP_CONNECTION_TIMEOUT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_USER", zend.ZEND_INI_USER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_PERDIR", zend.ZEND_INI_PERDIR, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_SYSTEM", zend.ZEND_INI_SYSTEM, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_ALL", zend.ZEND_INI_ALL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_SCANNER_NORMAL", zend.ZEND_INI_SCANNER_NORMAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_SCANNER_RAW", zend.ZEND_INI_SCANNER_RAW, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_SCANNER_TYPED", zend.ZEND_INI_SCANNER_TYPED, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_SCHEME", PHP_URL_SCHEME, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_HOST", PHP_URL_HOST, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_PORT", PHP_URL_PORT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_USER", PHP_URL_USER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_PASS", PHP_URL_PASS, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_PATH", PHP_URL_PATH, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_QUERY", PHP_URL_QUERY, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_FRAGMENT", PHP_URL_FRAGMENT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_QUERY_RFC1738", PHP_QUERY_RFC1738, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_QUERY_RFC3986", PHP_QUERY_RFC3986, zend.CONST_CS|zend.CONST_PERSISTENT)

	// #define REGISTER_MATH_CONSTANT(x) REGISTER_DOUBLE_CONSTANT ( # x , x , CONST_CS | CONST_PERSISTENT )

	zend.REGISTER_DOUBLE_CONSTANT("M_E", M_E, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_LOG2E", M_LOG2E, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_LOG10E", M_LOG10E, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_LN2", M_LN2, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_LN10", M_LN10, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_PI", M_PI, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_PI_2", M_PI_2, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_PI_4", M_PI_4, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_1_PI", M_1_PI, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_2_PI", M_2_PI, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_SQRTPI", M_SQRTPI, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_2_SQRTPI", M_2_SQRTPI, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_LNPI", M_LNPI, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_EULER", M_EULER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_SQRT2", M_SQRT2, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_SQRT1_2", M_SQRT1_2, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_SQRT3", M_SQRT3, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("INF", zend.ZEND_INFINITY, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("NAN", zend.ZEND_NAN, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_ROUND_HALF_UP", PHP_ROUND_HALF_UP, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_ROUND_HALF_DOWN", PHP_ROUND_HALF_DOWN, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_ROUND_HALF_EVEN", PHP_ROUND_HALF_EVEN, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_ROUND_HALF_ODD", PHP_ROUND_HALF_ODD, zend.CONST_CS|zend.CONST_PERSISTENT)
	RegisterPhpinfoConstants(type_, module_number)
	RegisterHtmlConstants(type_, module_number)
	RegisterStringConstants(type_, module_number)
	if ZmStartupVar(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupFile(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupPack(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupBrowscap(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupStandardFilters(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupUserFilters(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupPassword(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupMtRand(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if zm_startup_nl_langinfo(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupCrypt(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupLcg(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupDir(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupSyslog(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupArray(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupAssert(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupUrlScannerEx(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupProcOpen(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupExec(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupUserStreams(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupImagetypes(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	streams.PhpRegisterUrlStreamWrapper("php", &PhpStreamPhpWrapper)
	streams.PhpRegisterUrlStreamWrapper("file", &PhpPlainFilesWrapper)
	streams.PhpRegisterUrlStreamWrapper("glob", &streams.PhpGlobStreamWrapper)
	streams.PhpRegisterUrlStreamWrapper("data", &streams.PhpStreamRfc2397Wrapper)
	streams.PhpRegisterUrlStreamWrapper("http", &PhpStreamHttpWrapper)
	streams.PhpRegisterUrlStreamWrapper("ftp", &PhpStreamFtpWrapper)
	if ZmStartupDns(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupRandom(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupHrtime(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	return zend.SUCCESS
}

/* }}} */

func ZmShutdownBasic(type_ int, module_number int) int {
	ZmShutdownSyslog(type_, module_number)
	BasicGlobalsDtor(&BasicGlobals)
	streams.PhpUnregisterUrlStreamWrapper("php")
	streams.PhpUnregisterUrlStreamWrapper("http")
	streams.PhpUnregisterUrlStreamWrapper("ftp")
	ZmShutdownBrowscap(type_, module_number)
	ZmShutdownArray(type_, module_number)
	ZmShutdownAssert(type_, module_number)
	ZmShutdownUrlScannerEx(type_, module_number)
	ZmShutdownFile(type_, module_number)
	ZmShutdownStandardFilters(type_, module_number)
	ZmShutdownCrypt(type_, module_number)
	ZmShutdownRandom(type_, module_number)
	ZmShutdownPassword(type_, module_number)
	return zend.SUCCESS
}

/* }}} */

func ZmActivateBasic(type_ int, module_number int) int {
	memset(BG(strtok_table), 0, 256)
	BG(serialize_lock) = 0
	memset(&BG(serialize), 0, b.SizeOf("BG ( serialize )"))
	memset(&BG(unserialize), 0, b.SizeOf("BG ( unserialize )"))
	BG(strtok_string) = nil
	zend.ZVAL_UNDEF(&BG(strtok_zval))
	BG(strtok_last) = nil
	BG(locale_string) = nil
	BG(locale_changed) = 0
	BG(array_walk_fci) = zend.EmptyFcallInfo
	BG(array_walk_fci_cache) = zend.EmptyFcallInfoCache
	BG(user_compare_fci) = zend.EmptyFcallInfo
	BG(user_compare_fci_cache) = zend.EmptyFcallInfoCache
	BG(page_uid) = -1
	BG(page_gid) = -1
	BG(page_inode) = -1
	BG(page_mtime) = -1
	zend.ZendHashInit(&BG(putenv_ht), 1, nil, PhpPutenvDestructor, 0)
	BG(user_shutdown_function_names) = nil
	ZmActivateFilestat(type_, module_number)
	ZmActivateSyslog(type_, module_number)
	ZmActivateDir(type_, module_number)
	ZmActivateUrlScannerEx(type_, module_number)

	/* Setup default context */

	FG(default_context) = nil

	/* Default to global wrappers only */

	FG(stream_wrappers) = nil

	/* Default to global filters only */

	FG(stream_filters) = nil
	return zend.SUCCESS
}

/* }}} */

func ZmDeactivateBasic(type_ int, module_number int) int {
	zend.ZvalPtrDtor(&BG(strtok_zval))
	zend.ZVAL_UNDEF(&BG(strtok_zval))
	BG(strtok_string) = nil
	tsrm_env_lock()
	zend.ZendHashDestroy(&BG(putenv_ht))
	tsrm_env_unlock()
	BG(mt_rand_is_seeded) = 0
	if BG(umask) != -1 {
		umask(BG(umask))
	}

	/* Check if locale was changed and change it back
	 * to the value in startup environment */

	if BG(locale_changed) {
		setlocale(LC_ALL, "C")
		setlocale(LC_CTYPE, "")
		if BG(locale_string) {
			zend.ZendStringReleaseEx(BG(locale_string), 0)
			BG(locale_string) = nil
		}
	}

	/* FG(stream_wrappers) and FG(stream_filters) are destroyed
	 * during php_request_shutdown() */

	ZmDeactivateFilestat(type_, module_number)
	ZmDeactivateAssert(type_, module_number)
	ZmDeactivateUrlScannerEx(type_, module_number)
	streams.ZmDeactivateStreams(type_, module_number)
	if BG(user_tick_functions) {
		zend.ZendLlistDestroy(BG(user_tick_functions))
		zend.Efree(BG(user_tick_functions))
		BG(user_tick_functions) = nil
	}
	ZmDeactivateUserFilters(type_, module_number)
	ZmDeactivateBrowscap(type_, module_number)
	BG(page_uid) = -1
	BG(page_gid) = -1
	return zend.SUCCESS
}

/* }}} */

func ZmInfoBasic(ZEND_MODULE_INFO_FUNC_ARGS) {
	PhpInfoPrintTableStart()
	ZmInfoDl(zend.ZEND_MODULE_INFO_FUNC_ARGS_PASSTHRU)
	ZmInfoMail(zend.ZEND_MODULE_INFO_FUNC_ARGS_PASSTHRU)
	PhpInfoPrintTableEnd()
	ZmInfoAssert(zend.ZEND_MODULE_INFO_FUNC_ARGS_PASSTHRU)
}

/* }}} */

func ZifConstant(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var const_name *zend.ZendString
	var c *zend.Zval
	var scope *zend.ZendClassEntry
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &const_name, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	scope = zend.ZendGetExecutedScope()
	c = zend.ZendGetConstantEx(const_name, scope, zend.ZEND_FETCH_CLASS_SILENT)
	if c != nil {
		zend.ZVAL_COPY_OR_DUP(return_value, c)
		if zend.Z_TYPE_P(return_value) == zend.IS_CONSTANT_AST {
			if zend.UNEXPECTED(zend.ZvalUpdateConstantEx(return_value, scope) != zend.SUCCESS) {
				return
			}
		}
	} else {
		if zend.ExecutorGlobals.exception == nil {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Couldn't find constant %s", zend.ZSTR_VAL(const_name))
		}
		zend.RETVAL_NULL()
		return
	}
}

/* }}} */

/* {{{ proto string inet_ntop(string in_addr)
   Converts a packed inet address to a human readable IP address string */

func ZifInetNtop(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var address *byte
	var address_len int
	var af int = AF_INET
	var buffer []byte
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &address, &address_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if address_len == 16 {
		af = AF_INET6
	} else if address_len != 4 {
		zend.RETVAL_FALSE
		return
	}
	if !(inet_ntop(af, address, buffer, b.SizeOf("buffer"))) {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STRING(buffer)
	return
}

/* }}} */

/* {{{ proto string inet_pton(string ip_address)
   Converts a human readable IP address to a packed binary string */

func PhpInetPton(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var ret int
	var af int = AF_INET
	var address *byte
	var address_len int
	var buffer []byte
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &address, &address_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	memset(buffer, 0, b.SizeOf("buffer"))
	if strchr(address, ':') {
		af = AF_INET6
	} else if !(strchr(address, '.')) {
		zend.RETVAL_FALSE
		return
	}
	ret = inet_pton(af, address, buffer)
	if ret <= 0 {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STRINGL(buffer, b.Cond(af == AF_INET, 4, 16))
	return
}

/* }}} */

/* {{{ proto int ip2long(string ip_address)
   Converts a string containing an (IPv4) Internet Protocol dotted address into a proper address */

func ZifIp2long(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var addr *byte
	var addr_len int
	var ip __struct__in_addr
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &addr, &addr_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if addr_len == 0 || inet_pton(AF_INET, addr, &ip) != 1 {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(ntohl(ip.s_addr))
	return
}

/* }}} */

func ZifLong2ip(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var ip zend.ZendUlong
	var sip zend.ZendLong
	var myaddr __struct__in_addr
	var str []byte
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &sip, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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

	/* autoboxes on 32bit platforms, but that's expected */

	ip = zend.ZendUlong(sip)
	myaddr.s_addr = htonl(ip)
	if inet_ntop(AF_INET, &myaddr, str, b.SizeOf("str")) {
		zend.RETVAL_STRING(str)
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}

/* }}} */

func ZifGetenv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var ptr *byte
	var str *byte = nil
	var str_len int
	var local_only zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &local_only, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if str == nil {
		zend.ArrayInit(return_value)
		core.PhpImportEnvironmentVariables(return_value)
		return
	}
	if local_only == 0 {

		/* SAPI method returns an emalloc()'d string */

		ptr = core.SapiGetenv(str, str_len)
		if ptr != nil {

			// TODO: avoid realocation ???

			zend.RETVAL_STRING(ptr)
			zend.Efree(ptr)
			return
		}
	}
	tsrm_env_lock()

	/* system method returns a const */

	ptr = getenv(str)
	if ptr != nil {
		zend.RETVAL_STRING(ptr)
	}
	tsrm_env_unlock()
	if ptr != nil {
		return
	}
	zend.RETVAL_FALSE
	return
}

/* }}} */

/* {{{ proto bool putenv(string setting)
   Set the value of an environment variable */

func ZifPutenv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var setting *byte
	var setting_len int
	var p *byte
	var env **byte
	var pe PutenvEntry
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &setting, &setting_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if setting_len == 0 || setting[0] == '=' {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid parameter syntax")
		zend.RETVAL_FALSE
		return
	}
	pe.SetPutenvString(zend.Estrndup(setting, setting_len))
	pe.SetKey(zend.Estrndup(setting, setting_len))
	if b.Assign(&p, strchr(pe.GetKey(), '=')) {
		*p = '0'
	}
	pe.SetKeyLen(strlen(pe.GetKey()))
	tsrm_env_lock()
	zend.ZendHashStrDel(&BG(putenv_ht), pe.GetKey(), pe.GetKeyLen())

	/* find previous value */

	pe.SetPreviousValue(nil)
	for env = cli.Environ; env != nil && (*env) != nil; env++ {
		if !(strncmp(*env, pe.GetKey(), pe.GetKeyLen())) && (*env)[pe.GetKeyLen()] == '=' {
			pe.SetPreviousValue(*env)
			break
		}
	}
	if p == nil {
		unsetenv(pe.GetPutenvString())
	}
	if p == nil || putenv(pe.GetPutenvString()) == 0 {
		zend.ZendHashStrAddMem(&BG(putenv_ht), pe.GetKey(), pe.GetKeyLen(), &pe, b.SizeOf("putenv_entry"))
		if !(strncmp(pe.GetKey(), "TZ", pe.GetKeyLen())) {
			tzset()
		}
		tsrm_env_unlock()
		zend.RETVAL_TRUE
		return
	} else {
		zend.Efree(pe.GetPutenvString())
		zend.Efree(pe.GetKey())
		zend.RETVAL_FALSE
		return
	}
}

/* }}} */

/* {{{ free_argv()
   Free the memory allocated to an argv array. */

func FreeArgv(argv **byte, argc int) {
	var i int
	if argv != nil {
		for i = 0; i < argc; i++ {
			if argv[i] != nil {
				zend.Efree(argv[i])
			}
		}
		zend.Efree(argv)
	}
}

/* }}} */

func FreeLongopts(longopts *core.Opt) {
	var p *core.Opt
	if longopts != nil {
		for p = longopts; p != nil && p.opt_char != '-'; p++ {
			if p.opt_name != nil {
				zend.Efree((*byte)(p.opt_name))
			}
		}
	}
}

/* }}} */

func ParseOpts(opts *byte, result **core.Opt) int {
	var paras *core.Opt = nil
	var i uint
	var count uint = 0
	var opts_len uint = uint(strlen(opts))
	for i = 0; i < opts_len; i++ {
		if opts[i] >= 48 && opts[i] <= 57 || opts[i] >= 65 && opts[i] <= 90 || opts[i] >= 97 && opts[i] <= 122 {
			count++
		}
	}
	paras = zend.SafeEmalloc(b.SizeOf("opt_struct"), count, 0)
	memset(paras, 0, b.SizeOf("opt_struct")*count)
	*result = paras
	for (*opts) >= 48 && (*opts) <= 57 || (*opts) >= 65 && (*opts) <= 90 || (*opts) >= 97 && (*opts) <= 122 {
		paras.opt_char = *opts
		paras.need_param = (*(b.PreInc(&opts))) == ':'
		paras.opt_name = nil
		if paras.need_param == 1 {
			opts++
			if (*opts) == ':' {
				paras.need_param++
				opts++
			}
		}
		paras++
	}
	return count
}

/* }}} */

func ZifGetopt(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var options *byte = nil
	var argv **byte = nil
	var opt []byte = []byte{'0'}
	var optname *byte
	var argc int = 0
	var o int
	var options_len int = 0
	var len_ int
	var php_optarg *byte = nil
	var php_optind int = 1
	var val zend.Zval
	var args *zend.Zval = nil
	var p_longopts *zend.Zval = nil
	var zoptind *zend.Zval = nil
	var optname_len int = 0
	var opts *core.Opt
	var orig_opts *core.Opt
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &options, &options_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &p_longopts, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zoptind, 0)
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}

	/* Init zoptind to 1 */

	if zoptind != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zoptind, 1)
	}

	/* Get argv from the global symbol table. We calculate argc ourselves
	 * in order to be on the safe side, even though it is also available
	 * from the symbol table. */

	if (zend.Z_TYPE(core.PG(http_globals)[core.TRACK_VARS_SERVER]) == zend.IS_ARRAY || zend.ZendIsAutoGlobalStr(zend.ZEND_STRL("_SERVER")) != 0) && (b.Assign(&args, zend.ZendHashFindExInd(zend.Z_ARRVAL_P(&core.PG(http_globals)[core.TRACK_VARS_SERVER]), zend.ZSTR_KNOWN(zend.ZEND_STR_ARGV), 1)) != nil || b.Assign(&args, zend.ZendHashFindExInd(&(zend.ExecutorGlobals.symbol_table), zend.ZSTR_KNOWN(zend.ZEND_STR_ARGV), 1)) != nil) {
		var pos int = 0
		var entry *zend.Zval
		if zend.Z_TYPE_P(args) != zend.IS_ARRAY {
			zend.RETVAL_FALSE
			return
		}
		argc = zend.ZendHashNumElements(zend.Z_ARRVAL_P(args))

		/* Attempt to allocate enough memory to hold all of the arguments
		 * and a trailing NULL */

		argv = (**byte)(zend.SafeEmalloc(b.SizeOf("char *"), argc+1, 0))

		/* Iterate over the hash to construct the argv array. */

		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(args)
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				entry = _z
				var tmp_arg_str *zend.ZendString
				var arg_str *zend.ZendString = zend.ZvalGetTmpString(entry, &tmp_arg_str)
				argv[b.PostInc(&pos)] = zend.Estrdup(zend.ZSTR_VAL(arg_str))
				zend.ZendTmpStringRelease(tmp_arg_str)
			}
			break
		}

		/* The C Standard requires argv[argc] to be NULL - this might
		 * keep some getopt implementations happy. */

		argv[argc] = nil

		/* The C Standard requires argv[argc] to be NULL - this might
		 * keep some getopt implementations happy. */

	} else {

		/* Return false if we can't find argv. */

		zend.RETVAL_FALSE
		return
	}
	len_ = ParseOpts(options, &opts)
	if p_longopts != nil {
		var count int
		var entry *zend.Zval
		count = zend.ZendHashNumElements(zend.Z_ARRVAL_P(p_longopts))

		/* the first <len> slots are filled by the one short ops
		 * we now extend our array and jump to the new added structs */

		opts = (*core.Opt)(zend.Erealloc(opts, b.SizeOf("opt_struct")*(len_+count+1)))
		orig_opts = opts
		opts += len_
		memset(opts, 0, count*b.SizeOf("opt_struct"))

		/* Iterate over the hash to construct the argv array. */

		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(p_longopts)
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				entry = _z
				var tmp_arg_str *zend.ZendString
				var arg_str *zend.ZendString = zend.ZvalGetTmpString(entry, &tmp_arg_str)
				opts.need_param = 0
				opts.opt_name = zend.Estrdup(zend.ZSTR_VAL(arg_str))
				len_ = strlen(opts.opt_name)
				if len_ > 0 && opts.opt_name[len_-1] == ':' {
					opts.need_param++
					opts.opt_name[len_-1] = '0'
					if len_ > 1 && opts.opt_name[len_-2] == ':' {
						opts.need_param++
						opts.opt_name[len_-2] = '0'
					}
				}
				opts.opt_char = 0
				opts++
				zend.ZendTmpStringRelease(tmp_arg_str)
			}
			break
		}

		/* Iterate over the hash to construct the argv array. */

	} else {
		opts = (*core.Opt)(zend.Erealloc(opts, b.SizeOf("opt_struct")*(len_+1)))
		orig_opts = opts
		opts += len_
	}

	/* php_getopt want to identify the last param */

	opts.opt_char = '-'
	opts.need_param = 0
	opts.opt_name = nil

	/* Initialize the return value as an array. */

	zend.ArrayInit(return_value)

	/* after our pointer arithmetic jump back to the first element */

	opts = orig_opts
	for b.Assign(&o, core.PhpGetopt(argc, argv, opts, &php_optarg, &php_optind, 0, 1)) != -1 {

		/* Skip unknown arguments. */

		if o == core.PHP_GETOPT_INVALID_ARG {
			continue
		}

		/* Prepare the option character and the argument string. */

		if o == 0 {
			optname = opts[core.PhpOptidx].opt_name
		} else {
			if o == 1 {
				o = '-'
			}
			opt[0] = o
			optname = opt
		}
		if php_optarg != nil {

			/* keep the arg as binary, since the encoding is not known */

			zend.ZVAL_STRING(&val, php_optarg)

			/* keep the arg as binary, since the encoding is not known */

		} else {
			zend.ZVAL_FALSE(&val)
		}

		/* Add this option / argument pair to the result hash. */

		optname_len = strlen(optname)
		if !(optname_len > 1 && optname[0] == '0') && zend.IsNumericString(optname, optname_len, nil, nil, 0) == zend.IS_LONG {

			/* numeric string */

			var optname_int int = atoi(optname)
			if b.Assign(&args, zend.ZendHashIndexFind(zend.Z_ARRVAL_P(return_value), optname_int)) != nil {
				if zend.Z_TYPE_P(args) != zend.IS_ARRAY {
					zend.ConvertToArrayEx(args)
				}
				zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(args), &val)
			} else {
				zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), optname_int, &val)
			}
		} else {

			/* other strings */

			if b.Assign(&args, zend.ZendHashStrFind(zend.Z_ARRVAL_P(return_value), optname, strlen(optname))) != nil {
				if zend.Z_TYPE_P(args) != zend.IS_ARRAY {
					zend.ConvertToArrayEx(args)
				}
				zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(args), &val)
			} else {
				zend.ZendHashStrAdd(zend.Z_ARRVAL_P(return_value), optname, strlen(optname), &val)
			}

			/* other strings */

		}
		php_optarg = nil
	}

	/* Set zoptind to php_optind */

	if zoptind != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zoptind, php_optind)
	}
	FreeLongopts(orig_opts)
	zend.Efree(orig_opts)
	FreeArgv(argv, argc)
}

/* }}} */

func ZifFlush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	core.SapiFlush()
}

/* }}} */

func ZifSleep(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var num zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &num, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if num < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Number of seconds must be greater than or equal to 0")
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(core.PhpSleep(uint(num)))
	return
}

/* }}} */

func ZifUsleep(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var num zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &num, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if num < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Number of microseconds must be greater than or equal to 0")
		zend.RETVAL_FALSE
		return
	}
	usleep(uint(num))
}

/* }}} */

/* {{{ proto mixed time_nanosleep(int seconds, int nanoseconds)
   Delay for a number of seconds and nano seconds */

func ZifTimeNanosleep(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var tv_sec zend.ZendLong
	var tv_nsec zend.ZendLong
	var php_req __struct__timespec
	var php_rem __struct__timespec
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &tv_sec, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &tv_nsec, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if tv_sec < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The seconds value must be greater than 0")
		zend.RETVAL_FALSE
		return
	}
	if tv_nsec < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The nanoseconds value must be greater than 0")
		zend.RETVAL_FALSE
		return
	}
	php_req.tv_sec = int64(tv_sec)
	php_req.tv_nsec = long(tv_nsec)
	if !(nanosleep(&php_req, &php_rem)) {
		zend.RETVAL_TRUE
		return
	} else if errno == EINTR {
		zend.ArrayInit(return_value)
		zend.AddAssocLongEx(return_value, "seconds", b.SizeOf("\"seconds\"")-1, php_rem.tv_sec)
		zend.AddAssocLongEx(return_value, "nanoseconds", b.SizeOf("\"nanoseconds\"")-1, php_rem.tv_nsec)
		return
	} else if errno == EINVAL {
		core.PhpErrorDocref(nil, zend.E_WARNING, "nanoseconds was not in the range 0 to 999 999 999 or seconds was negative")
	}
	zend.RETVAL_FALSE
	return
}

/* }}} */

func ZifTimeSleepUntil(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var target_secs float64
	var tm __struct__timeval
	var php_req __struct__timespec
	var php_rem __struct__timespec
	var current_ns uint64
	var target_ns uint64
	var diff_ns uint64
	var ns_per_sec uint64 = 1000000000
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgDouble(_arg, &target_secs, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_DOUBLE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if gettimeofday((*__struct__timeval)(&tm), nil) != 0 {
		zend.RETVAL_FALSE
		return
	}
	target_ns = uint64_t(target_secs * ns_per_sec)
	current_ns = uint64(tm.tv_sec)*ns_per_sec + uint64(tm.tv_usec)*1000
	if target_ns < current_ns {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Sleep until to time is less than current time")
		zend.RETVAL_FALSE
		return
	}
	diff_ns = target_ns - current_ns
	php_req.tv_sec = time_t(diff_ns / ns_per_sec)
	php_req.tv_nsec = long(diff_ns % ns_per_sec)
	for nanosleep(&php_req, &php_rem) {
		if errno == EINTR {
			php_req.tv_sec = php_rem.tv_sec
			php_req.tv_nsec = php_rem.tv_nsec
		} else {
			zend.RETVAL_FALSE
			return
		}
	}
	zend.RETVAL_TRUE
	return
}

/* }}} */

/* {{{ proto string get_current_user(void)
   Get the name of the owner of the current PHP script */

func ZifGetCurrentUser(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_STRING(core.PhpGetCurrentUser())
	return
}

/* }}} */

/* {{{ add_config_entry
 */

func AddConfigEntry(h zend.ZendUlong, key *zend.ZendString, entry *zend.Zval, retval *zend.Zval) {
	if zend.Z_TYPE_P(entry) == zend.IS_STRING {
		var str *zend.ZendString = zend.Z_STR_P(entry)
		if zend.ZSTR_IS_INTERNED(str) == 0 {
			if (zend.GC_FLAGS(str) & zend.GC_PERSISTENT) == 0 {
				zend.ZendStringAddref(str)
			} else {
				str = zend.ZendStringInit(zend.ZSTR_VAL(str), zend.ZSTR_LEN(str), 0)
			}
		}
		if key != nil {
			zend.AddAssocStrEx(retval, zend.ZSTR_VAL(key), zend.ZSTR_LEN(key), str)
		} else {
			zend.AddIndexStr(retval, h, str)
		}
	} else if zend.Z_TYPE_P(entry) == zend.IS_ARRAY {
		var tmp zend.Zval
		zend.ArrayInit(&tmp)
		AddConfigEntries(zend.Z_ARRVAL_P(entry), &tmp)
		zend.ZendHashUpdate(zend.Z_ARRVAL_P(retval), key, &tmp)
	}
}

/* }}} */

func AddConfigEntries(hash *zend.HashTable, return_value *zend.Zval) {
	var h zend.ZendUlong
	var key *zend.ZendString
	var zv *zend.Zval
	for {
		var __ht *zend.HashTable = hash
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			h = _p.h
			key = _p.key
			zv = _z
			AddConfigEntry(h, key, zv, return_value)
		}
		break
	}
}

/* }}} */

func ZifGetCfgVar(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var varname *byte
	var varname_len int
	var retval *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &varname, &varname_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	retval = core.CfgGetEntry(varname, uint32(varname_len))
	if retval != nil {
		if zend.Z_TYPE_P(retval) == zend.IS_ARRAY {
			zend.ArrayInit(return_value)
			AddConfigEntries(zend.Z_ARRVAL_P(retval), return_value)
			return
		} else {
			zend.RETVAL_STRING(zend.Z_STRVAL_P(retval))
			return
		}
	} else {
		zend.RETVAL_FALSE
		return
	}
}

/* }}} */

func ZifGetMagicQuotesRuntime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_FALSE
	return
}

/* }}} */

func ZifGetMagicQuotesGpc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_FALSE
	return
}

/* }}} */

func ZifErrorLog(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var message *byte
	var opt *byte = nil
	var headers *byte = nil
	var message_len int
	var opt_len int = 0
	var headers_len int = 0
	var opt_err int = 0
	var argc int = zend.ZEND_NUM_ARGS()
	var erropt zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 4
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &message, &message_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &erropt, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgPath(_arg, &opt, &opt_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &headers, &headers_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if argc > 1 {
		opt_err = int(erropt)
	}
	if _phpErrorLogEx(opt_err, message, message_len, opt, headers) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}

/* }}} */

func _phpErrorLog(opt_err int, message *byte, opt *byte, headers *byte) int {
	return _phpErrorLogEx(opt_err, message, b.CondF1(opt_err == 3, func() __auto__ { return strlen(message) }, 0), opt, headers)
}

/* }}} */

func _phpErrorLogEx(opt_err int, message *byte, message_len int, opt *byte, headers *byte) int {
	var stream *core.PhpStream = nil
	var nbytes int
	switch opt_err {
	case 1:
		if PhpMail(opt, "PHP error_log message", message, headers, nil) == 0 {
			return zend.FAILURE
		}
		break
	case 2:
		core.PhpErrorDocref(nil, zend.E_WARNING, "TCP/IP option not available!")
		return zend.FAILURE
		break
	case 3:
		stream = core.PhpStreamOpenWrapper(opt, "a", core.IGNORE_URL_WIN|core.REPORT_ERRORS, nil)
		if stream == nil {
			return zend.FAILURE
		}
		nbytes = core.PhpStreamWrite(stream, message, message_len)
		core.PhpStreamClose(stream)
		if nbytes != message_len {
			return zend.FAILURE
		}
		break
	case 4:
		if core.sapi_module.log_message != nil {
			core.sapi_module.log_message(message, -1)
		} else {
			return zend.FAILURE
		}
		break
	default:
		core.PhpLogErrWithSeverity(message, LOG_NOTICE)
		break
	}
	return zend.SUCCESS
}

/* }}} */

func ZifErrorGetLast(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if core.PG(last_error_message) {
		zend.ArrayInit(return_value)
		zend.AddAssocLongEx(return_value, "type", b.SizeOf("\"type\"")-1, core.PG(last_error_type))
		zend.AddAssocStringEx(return_value, "message", b.SizeOf("\"message\"")-1, core.PG(last_error_message))
		zend.AddAssocStringEx(return_value, "file", b.SizeOf("\"file\"")-1, b.CondF1(core.PG(last_error_file), func() __auto__ { return core.PG(last_error_file) }, "-"))
		zend.AddAssocLongEx(return_value, "line", b.SizeOf("\"line\"")-1, core.PG(last_error_lineno))
	}
}

/* }}} */

func ZifErrorClearLast(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if core.PG(last_error_message) {
		core.PG(last_error_type) = 0
		core.PG(last_error_lineno) = 0
		zend.Free(core.PG(last_error_message))
		core.PG(last_error_message) = nil
		if core.PG(last_error_file) {
			zend.Free(core.PG(last_error_file))
			core.PG(last_error_file) = nil
		}
	}
}

/* }}} */

func ZifCallUserFunc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var retval zend.Zval
	var fci zend.ZendFcallInfo
	var fci_cache zend.ZendFcallInfoCache
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0) {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if zend.UNEXPECTED(_error != nil) {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				fci.params = _real_arg + 1
				fci.param_count = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				fci.params = nil
				fci.param_count = 0
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	fci.retval = &retval
	if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_UNDEF {
		if zend.Z_ISREF(retval) {
			zend.ZendUnwrapReference(&retval)
		}
		zend.ZVAL_COPY_VALUE(return_value, &retval)
	}
}

/* }}} */

func ZifCallUserFuncArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var params *zend.Zval
	var retval zend.Zval
	var fci zend.ZendFcallInfo
	var fci_cache zend.ZendFcallInfoCache
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0) {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if zend.UNEXPECTED(_error != nil) {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &params, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	zend.ZendFcallInfoArgs(&fci, params)
	fci.retval = &retval
	if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_UNDEF {
		if zend.Z_ISREF(retval) {
			zend.ZendUnwrapReference(&retval)
		}
		zend.ZVAL_COPY_VALUE(return_value, &retval)
	}
	zend.ZendFcallInfoArgsClear(&fci, 1)
}

/* }}} */

func ZifForwardStaticCall(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var retval zend.Zval
	var fci zend.ZendFcallInfo
	var fci_cache zend.ZendFcallInfoCache
	var called_scope *zend.ZendClassEntry
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0) {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if zend.UNEXPECTED(_error != nil) {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				fci.params = _real_arg + 1
				fci.param_count = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				fci.params = nil
				fci.param_count = 0
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if !(zend.EX(prev_execute_data).func_.common.scope) {
		zend.ZendThrowError(nil, "Cannot call forward_static_call() when no class scope is active")
		return
	}
	fci.retval = &retval
	called_scope = zend.ZendGetCalledScope(execute_data)
	if called_scope != nil && fci_cache.calling_scope != nil && zend.InstanceofFunction(called_scope, fci_cache.calling_scope) != 0 {
		fci_cache.called_scope = called_scope
	}
	if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_UNDEF {
		if zend.Z_ISREF(retval) {
			zend.ZendUnwrapReference(&retval)
		}
		zend.ZVAL_COPY_VALUE(return_value, &retval)
	}
}

/* }}} */

func ZifForwardStaticCallArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var params *zend.Zval
	var retval zend.Zval
	var fci zend.ZendFcallInfo
	var fci_cache zend.ZendFcallInfoCache
	var called_scope *zend.ZendClassEntry
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0) {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if zend.UNEXPECTED(_error != nil) {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &params, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	zend.ZendFcallInfoArgs(&fci, params)
	fci.retval = &retval
	called_scope = zend.ZendGetCalledScope(execute_data)
	if called_scope != nil && fci_cache.calling_scope != nil && zend.InstanceofFunction(called_scope, fci_cache.calling_scope) != 0 {
		fci_cache.called_scope = called_scope
	}
	if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_UNDEF {
		if zend.Z_ISREF(retval) {
			zend.ZendUnwrapReference(&retval)
		}
		zend.ZVAL_COPY_VALUE(return_value, &retval)
	}
	zend.ZendFcallInfoArgsClear(&fci, 1)
}

/* }}} */

func UserShutdownFunctionDtor(zv *zend.Zval) {
	var i int
	var shutdown_function_entry *PhpShutdownFunctionEntry = zend.Z_PTR_P(zv)
	for i = 0; i < shutdown_function_entry.GetArgCount(); i++ {
		zend.ZvalPtrDtor(&shutdown_function_entry.arguments[i])
	}
	zend.Efree(shutdown_function_entry.GetArguments())
	zend.Efree(shutdown_function_entry)
}

/* }}} */

func UserTickFunctionDtor(tick_function_entry *UserTickFunctionEntry) {
	var i int
	for i = 0; i < tick_function_entry.GetArgCount(); i++ {
		zend.ZvalPtrDtor(&tick_function_entry.arguments[i])
	}
	zend.Efree(tick_function_entry.GetArguments())
}

/* }}} */

func UserShutdownFunctionCall(zv *zend.Zval) int {
	var shutdown_function_entry *PhpShutdownFunctionEntry = zend.Z_PTR_P(zv)
	var retval zend.Zval
	if zend.ZendIsCallable(&shutdown_function_entry.arguments[0], 0, nil) == 0 {
		var function_name *zend.ZendString = zend.ZendGetCallableName(&shutdown_function_entry.arguments[0])
		core.PhpError(zend.E_WARNING, "(Registered shutdown functions) Unable to call %s() - function does not exist", zend.ZSTR_VAL(function_name))
		zend.ZendStringReleaseEx(function_name, 0)
		return 0
	}
	if zend.CallUserFunction(nil, nil, &shutdown_function_entry.arguments[0], &retval, shutdown_function_entry.GetArgCount()-1, shutdown_function_entry.GetArguments()+1) == zend.SUCCESS {
		zend.ZvalPtrDtor(&retval)
	}
	return 0
}

/* }}} */

func UserTickFunctionCall(tick_fe *UserTickFunctionEntry) {
	var retval zend.Zval
	var function *zend.Zval = &tick_fe.arguments[0]

	/* Prevent reentrant calls to the same user ticks function */

	if tick_fe.GetCalling() == 0 {
		tick_fe.SetCalling(1)
		if zend.CallUserFunction(nil, nil, function, &retval, tick_fe.GetArgCount()-1, tick_fe.GetArguments()+1) == zend.SUCCESS {
			zend.ZvalPtrDtor(&retval)
		} else {
			var obj *zend.Zval
			var method *zend.Zval
			if zend.Z_TYPE_P(function) == zend.IS_STRING {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to call %s() - function does not exist", zend.Z_STRVAL_P(function))
			} else if zend.Z_TYPE_P(function) == zend.IS_ARRAY && b.Assign(&obj, zend.ZendHashIndexFind(zend.Z_ARRVAL_P(function), 0)) != nil && b.Assign(&method, zend.ZendHashIndexFind(zend.Z_ARRVAL_P(function), 1)) != nil && zend.Z_TYPE_P(obj) == zend.IS_OBJECT && zend.Z_TYPE_P(method) == zend.IS_STRING {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to call %s::%s() - function does not exist", zend.ZSTR_VAL(zend.Z_OBJCE_P(obj).name), zend.Z_STRVAL_P(method))
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to call tick function")
			}
		}
		tick_fe.SetCalling(0)
	}

	/* Prevent reentrant calls to the same user ticks function */
}

/* }}} */

func RunUserTickFunctions(tick_count int, arg any) {
	zend.ZendLlistApply(BG(user_tick_functions), zend.LlistApplyFuncT(UserTickFunctionCall))
}

/* }}} */

func UserTickFunctionCompare(tick_fe1 *UserTickFunctionEntry, tick_fe2 *UserTickFunctionEntry) int {
	var func1 *zend.Zval = &tick_fe1.arguments[0]
	var func2 *zend.Zval = &tick_fe2.arguments[0]
	var ret int
	if zend.Z_TYPE_P(func1) == zend.IS_STRING && zend.Z_TYPE_P(func2) == zend.IS_STRING {
		ret = zend.ZendBinaryZvalStrcmp(func1, func2) == 0
	} else if zend.Z_TYPE_P(func1) == zend.IS_ARRAY && zend.Z_TYPE_P(func2) == zend.IS_ARRAY {
		ret = zend.ZendCompareArrays(func1, func2) == 0
	} else if zend.Z_TYPE_P(func1) == zend.IS_OBJECT && zend.Z_TYPE_P(func2) == zend.IS_OBJECT {
		ret = zend.ZendCompareObjects(func1, func2) == 0
	} else {
		ret = 0
	}
	if ret != 0 && tick_fe1.GetCalling() != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to delete tick function executed at the moment")
		return 0
	}
	return ret
}

/* }}} */

func PhpCallShutdownFunctions() {
	if BG(user_shutdown_function_names) {
		var __orig_bailout *JMP_BUF = zend.ExecutorGlobals.bailout
		var __bailout JMP_BUF
		zend.ExecutorGlobals.bailout = &__bailout
		if zend.SETJMP(__bailout) == 0 {
			zend.ZendHashApply(BG(user_shutdown_function_names), UserShutdownFunctionCall)
		}
		zend.ExecutorGlobals.bailout = __orig_bailout
	}
}

/* }}} */

func PhpFreeShutdownFunctions() {
	if BG(user_shutdown_function_names) {
		var __orig_bailout *JMP_BUF = zend.ExecutorGlobals.bailout
		var __bailout JMP_BUF
		zend.ExecutorGlobals.bailout = &__bailout
		if zend.SETJMP(__bailout) == 0 {
			zend.ZendHashDestroy(BG(user_shutdown_function_names))
			zend.FREE_HASHTABLE(BG(user_shutdown_function_names))
			BG(user_shutdown_function_names) = nil
		} else {
			zend.ExecutorGlobals.bailout = __orig_bailout

			/* maybe shutdown method call exit, we just ignore it */

			zend.FREE_HASHTABLE(BG(user_shutdown_function_names))
			BG(user_shutdown_function_names) = nil
		}
		zend.ExecutorGlobals.bailout = __orig_bailout
	}
}

/* }}} */

func ZifRegisterShutdownFunction(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var shutdown_function_entry PhpShutdownFunctionEntry
	var i int
	shutdown_function_entry.SetArgCount(zend.ZEND_NUM_ARGS())
	if shutdown_function_entry.GetArgCount() < 1 {
		zend.WRONG_PARAM_COUNT
	}
	shutdown_function_entry.SetArguments((*zend.Zval)(zend.SafeEmalloc(b.SizeOf("zval"), shutdown_function_entry.GetArgCount(), 0)))
	if zend.ZendGetParametersArray(zend.ZEND_NUM_ARGS(), shutdown_function_entry.GetArgCount(), shutdown_function_entry.GetArguments()) == zend.FAILURE {
		zend.Efree(shutdown_function_entry.GetArguments())
		zend.RETVAL_FALSE
		return
	}

	/* Prevent entering of anything but valid callback (syntax check only!) */

	if zend.ZendIsCallable(&shutdown_function_entry.arguments[0], 0, nil) == 0 {
		var callback_name *zend.ZendString = zend.ZendGetCallableName(&shutdown_function_entry.arguments[0])
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid shutdown callback '%s' passed", zend.ZSTR_VAL(callback_name))
		zend.Efree(shutdown_function_entry.GetArguments())
		zend.ZendStringReleaseEx(callback_name, 0)
		zend.RETVAL_FALSE
	} else {
		if !(BG(user_shutdown_function_names)) {
			zend.ALLOC_HASHTABLE(BG(user_shutdown_function_names))
			zend.ZendHashInit(BG(user_shutdown_function_names), 0, nil, UserShutdownFunctionDtor, 0)
		}
		for i = 0; i < shutdown_function_entry.GetArgCount(); i++ {
			zend.Z_TRY_ADDREF(shutdown_function_entry.GetArguments()[i])
		}
		zend.ZendHashNextIndexInsertMem(BG(user_shutdown_function_names), &shutdown_function_entry, b.SizeOf("php_shutdown_function_entry"))
	}

	/* Prevent entering of anything but valid callback (syntax check only!) */
}

/* }}} */

func RegisterUserShutdownFunction(function_name *byte, function_len int, shutdown_function_entry *PhpShutdownFunctionEntry) zend.ZendBool {
	if !(BG(user_shutdown_function_names)) {
		zend.ALLOC_HASHTABLE(BG(user_shutdown_function_names))
		zend.ZendHashInit(BG(user_shutdown_function_names), 0, nil, UserShutdownFunctionDtor, 0)
	}
	zend.ZendHashStrUpdateMem(BG(user_shutdown_function_names), function_name, function_len, shutdown_function_entry, b.SizeOf("php_shutdown_function_entry"))
	return 1
}

/* }}} */

func RemoveUserShutdownFunction(function_name *byte, function_len int) zend.ZendBool {
	if BG(user_shutdown_function_names) {
		return zend.ZendHashStrDel(BG(user_shutdown_function_names), function_name, function_len) != zend.FAILURE
	}
	return 0
}

/* }}} */

func AppendUserShutdownFunction(shutdown_function_entry PhpShutdownFunctionEntry) zend.ZendBool {
	if !(BG(user_shutdown_function_names)) {
		zend.ALLOC_HASHTABLE(BG(user_shutdown_function_names))
		zend.ZendHashInit(BG(user_shutdown_function_names), 0, nil, UserShutdownFunctionDtor, 0)
	}
	return zend.ZendHashNextIndexInsertMem(BG(user_shutdown_function_names), &shutdown_function_entry, b.SizeOf("php_shutdown_function_entry")) != nil
}

/* }}} */

func PhpGetHighlight(syntax_highlighter_ini *zend.ZendSyntaxHighlighterIni) {
	syntax_highlighter_ini.highlight_comment = zend.INI_STR("highlight.comment")
	syntax_highlighter_ini.highlight_default = zend.INI_STR("highlight.default")
	syntax_highlighter_ini.highlight_html = zend.INI_STR("highlight.html")
	syntax_highlighter_ini.highlight_keyword = zend.INI_STR("highlight.keyword")
	syntax_highlighter_ini.highlight_string = zend.INI_STR("highlight.string")
}

/* }}} */

func ZifHighlightFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var ret int
	var syntax_highlighter_ini zend.ZendSyntaxHighlighterIni
	var i zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &i, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if core.PhpCheckOpenBasedir(filename) != 0 {
		zend.RETVAL_FALSE
		return
	}
	if i != 0 {
		core.PhpOutputStartDefault()
	}
	PhpGetHighlight(&syntax_highlighter_ini)
	ret = zend.HighlightFile(filename, &syntax_highlighter_ini)
	if ret == zend.FAILURE {
		if i != 0 {
			core.PhpOutputEnd()
		}
		zend.RETVAL_FALSE
		return
	}
	if i != 0 {
		core.PhpOutputGetContents(return_value)
		core.PhpOutputDiscard()
	} else {
		zend.RETVAL_TRUE
		return
	}
}

/* }}} */

func ZifPhpStripWhitespace(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var original_lex_state zend.ZendLexState
	var file_handle zend.ZendFileHandle
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	core.PhpOutputStartDefault()
	zend.ZendStreamInitFilename(&file_handle, filename)
	zend.ZendSaveLexicalState(&original_lex_state)
	if zend.OpenFileForScanning(&file_handle) == zend.FAILURE {
		zend.ZendRestoreLexicalState(&original_lex_state)
		core.PhpOutputEnd()
		zend.RETVAL_EMPTY_STRING()
		return
	}
	zend.ZendStrip()
	zend.ZendDestroyFileHandle(&file_handle)
	zend.ZendRestoreLexicalState(&original_lex_state)
	core.PhpOutputGetContents(return_value)
	core.PhpOutputDiscard()
}

/* }}} */

func ZifHighlightString(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var expr *zend.Zval
	var syntax_highlighter_ini zend.ZendSyntaxHighlighterIni
	var hicompiled_string_description *byte
	var i zend.ZendBool = 0
	var old_error_reporting int = zend.ExecutorGlobals.error_reporting
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &expr, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &i, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if zend.TryConvertToString(expr) == 0 {
		return
	}
	if i != 0 {
		core.PhpOutputStartDefault()
	}
	zend.ExecutorGlobals.error_reporting = zend.E_ERROR
	PhpGetHighlight(&syntax_highlighter_ini)
	hicompiled_string_description = zend.ZendMakeCompiledStringDescription("highlighted code")
	if zend.HighlightString(expr, &syntax_highlighter_ini, hicompiled_string_description) == zend.FAILURE {
		zend.Efree(hicompiled_string_description)
		zend.ExecutorGlobals.error_reporting = old_error_reporting
		if i != 0 {
			core.PhpOutputEnd()
		}
		zend.RETVAL_FALSE
		return
	}
	zend.Efree(hicompiled_string_description)
	zend.ExecutorGlobals.error_reporting = old_error_reporting
	if i != 0 {
		core.PhpOutputGetContents(return_value)
		core.PhpOutputDiscard()
	} else {
		zend.RETVAL_TRUE
		return
	}
}

/* }}} */

func ZifIniGet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var varname *zend.ZendString
	var val *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &varname, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	val = zend.ZendIniGetValue(varname)
	if val == nil {
		zend.RETVAL_FALSE
		return
	}
	if zend.ZSTR_IS_INTERNED(val) != 0 {
		zend.RETVAL_INTERNED_STR(val)
	} else if zend.ZSTR_LEN(val) == 0 {
		zend.RETVAL_EMPTY_STRING()
	} else if zend.ZSTR_LEN(val) == 1 {
		zend.RETVAL_INTERNED_STR(zend.ZSTR_CHAR(zend.ZendUchar(zend.ZSTR_VAL(val)[0])))
	} else if (zend.GC_FLAGS(val) & zend.GC_PERSISTENT) == 0 {
		zend.ZVAL_NEW_STR(return_value, zend.ZendStringCopy(val))
	} else {
		zend.ZVAL_NEW_STR(return_value, zend.ZendStringInit(zend.ZSTR_VAL(val), zend.ZSTR_LEN(val), 0))
	}
}

/* }}} */

func ZifIniGetAll(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var extname *byte = nil
	var extname_len int = 0
	var module_number int = 0
	var module *zend.ZendModuleEntry
	var details zend.ZendBool = 1
	var key *zend.ZendString
	var ini_entry *zend.ZendIniEntry
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &extname, &extname_len, 1) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &details, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	zend.ZendIniSortEntries()
	if extname != nil {
		if b.Assign(&module, zend.ZendHashStrFindPtr(&zend.ModuleRegistry, extname, extname_len)) == nil {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to find extension '%s'", extname)
			zend.RETVAL_FALSE
			return
		}
		module_number = module.module_number
	}
	zend.ArrayInit(return_value)
	for {
		var __ht *zend.HashTable = zend.ExecutorGlobals.ini_directives
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			key = _p.key
			ini_entry = zend.Z_PTR_P(_z)
			var option zend.Zval
			if module_number != 0 && ini_entry.module_number != module_number {
				continue
			}
			if key == nil || zend.ZSTR_VAL(key)[0] != 0 {
				if details != 0 {
					zend.ArrayInit(&option)
					if ini_entry.orig_value != nil {
						zend.AddAssocStr(&option, "global_value", zend.ZendStringCopy(ini_entry.orig_value))
					} else if ini_entry.value != nil {
						zend.AddAssocStr(&option, "global_value", zend.ZendStringCopy(ini_entry.value))
					} else {
						zend.AddAssocNull(&option, "global_value")
					}
					if ini_entry.value != nil {
						zend.AddAssocStr(&option, "local_value", zend.ZendStringCopy(ini_entry.value))
					} else {
						zend.AddAssocNull(&option, "local_value")
					}
					zend.AddAssocLong(&option, "access", ini_entry.modifiable)
					zend.ZendSymtableUpdate(zend.Z_ARRVAL_P(return_value), ini_entry.name, &option)
				} else {
					if ini_entry.value != nil {
						var zv zend.Zval
						zend.ZVAL_STR_COPY(&zv, ini_entry.value)
						zend.ZendSymtableUpdate(zend.Z_ARRVAL_P(return_value), ini_entry.name, &zv)
					} else {
						zend.ZendSymtableUpdate(zend.Z_ARRVAL_P(return_value), ini_entry.name, &(zend.ExecutorGlobals.uninitialized_zval))
					}
				}
			}
		}
		break
	}
}

/* }}} */

func PhpIniCheckPath(option_name *byte, option_len int, new_option_name string, new_option_len int) int {
	if option_len+1 != new_option_len {
		return 0
	}
	return !(strncmp(option_name, new_option_name, option_len))
}

/* }}} */

func ZifIniSet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var varname *zend.ZendString
	var new_value *zend.ZendString
	var val *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &varname, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &new_value, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	val = zend.ZendIniGetValue(varname)

	/* copy to return here, because alter might free it! */

	if val != nil {
		if zend.ZSTR_IS_INTERNED(val) != 0 {
			zend.RETVAL_INTERNED_STR(val)
		} else if zend.ZSTR_LEN(val) == 0 {
			zend.RETVAL_EMPTY_STRING()
		} else if zend.ZSTR_LEN(val) == 1 {
			zend.RETVAL_INTERNED_STR(zend.ZSTR_CHAR(zend.ZendUchar(zend.ZSTR_VAL(val)[0])))
		} else if (zend.GC_FLAGS(val) & zend.GC_PERSISTENT) == 0 {
			zend.ZVAL_NEW_STR(return_value, zend.ZendStringCopy(val))
		} else {
			zend.ZVAL_NEW_STR(return_value, zend.ZendStringInit(zend.ZSTR_VAL(val), zend.ZSTR_LEN(val), 0))
		}
	} else {
		zend.RETVAL_FALSE
	}

	// #define _CHECK_PATH(var,var_len,ini) php_ini_check_path ( var , var_len , ini , sizeof ( ini ) )

	/* open basedir check */

	if core.PG(open_basedir) {
		if PhpIniCheckPath(zend.ZSTR_VAL(varname), zend.ZSTR_LEN(varname), "error_log", b.SizeOf("\"error_log\"")) != 0 || PhpIniCheckPath(zend.ZSTR_VAL(varname), zend.ZSTR_LEN(varname), "java.class.path", b.SizeOf("\"java.class.path\"")) != 0 || PhpIniCheckPath(zend.ZSTR_VAL(varname), zend.ZSTR_LEN(varname), "java.home", b.SizeOf("\"java.home\"")) != 0 || PhpIniCheckPath(zend.ZSTR_VAL(varname), zend.ZSTR_LEN(varname), "mail.log", b.SizeOf("\"mail.log\"")) != 0 || PhpIniCheckPath(zend.ZSTR_VAL(varname), zend.ZSTR_LEN(varname), "java.library.path", b.SizeOf("\"java.library.path\"")) != 0 || PhpIniCheckPath(zend.ZSTR_VAL(varname), zend.ZSTR_LEN(varname), "vpopmail.directory", b.SizeOf("\"vpopmail.directory\"")) != 0 {
			if core.PhpCheckOpenBasedir(zend.ZSTR_VAL(new_value)) != 0 {
				zend.ZvalPtrDtorStr(return_value)
				zend.RETVAL_FALSE
				return
			}
		}
	}
	if zend.ZendAlterIniEntryEx(varname, new_value, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0) == zend.FAILURE {
		zend.ZvalPtrDtorStr(return_value)
		zend.RETVAL_FALSE
		return
	}
}

/* }}} */

func ZifIniRestore(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var varname *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &varname, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	zend.ZendRestoreIniEntry(varname, core.PHP_INI_STAGE_RUNTIME)
}

/* }}} */

func ZifSetIncludePath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var new_value *zend.ZendString
	var old_value *byte
	var key *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgPathStr(_arg, &new_value, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	old_value = zend.ZendIniString("include_path", b.SizeOf("\"include_path\"")-1, 0)

	/* copy to return here, because alter might free it! */

	if old_value != nil {
		zend.RETVAL_STRING(old_value)
	} else {
		zend.RETVAL_FALSE
	}
	key = zend.ZendStringInit("include_path", b.SizeOf("\"include_path\"")-1, 0)
	if zend.ZendAlterIniEntryEx(key, new_value, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0) == zend.FAILURE {
		zend.ZendStringReleaseEx(key, 0)
		zend.ZvalPtrDtorStr(return_value)
		zend.RETVAL_FALSE
		return
	}
	zend.ZendStringReleaseEx(key, 0)
}

/* }}} */

func ZifGetIncludePath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	str = zend.ZendIniString("include_path", b.SizeOf("\"include_path\"")-1, 0)
	if str == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STRING(str)
	return
}

/* }}} */

func ZifRestoreIncludePath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var key *zend.ZendString
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	key = zend.ZendStringInit("include_path", b.SizeOf("\"include_path\"")-1, 0)
	zend.ZendRestoreIniEntry(key, core.PHP_INI_STAGE_RUNTIME)
	zend.ZendStringEfree(key)
}

/* }}} */

func ZifPrintR(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var var_ *zend.Zval
	var do_return zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &var_, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &do_return, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if do_return != 0 {
		zend.RETVAL_STR(zend.ZendPrintZvalRToStr(var_, 0))
		return
	} else {
		zend.ZendPrintZvalR(var_, 0)
		zend.RETVAL_TRUE
		return
	}
}

/* }}} */

func ZifConnectionAborted(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	zend.RETVAL_LONG(core.PG(connection_status) & core.PHP_CONNECTION_ABORTED)
	return
}

/* }}} */

func ZifConnectionStatus(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	zend.RETVAL_LONG(core.PG(connection_status))
	return
}

/* }}} */

func ZifIgnoreUserAbort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg zend.ZendBool = 0
	var old_setting int
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &arg, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	old_setting = uint16(core.PG(ignore_user_abort))
	if zend.ZEND_NUM_ARGS() != 0 {
		var key *zend.ZendString = zend.ZendStringInit("ignore_user_abort", b.SizeOf("\"ignore_user_abort\"")-1, 0)
		zend.ZendAlterIniEntryChars(key, b.Cond(arg != 0, "1", "0"), 1, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME)
		zend.ZendStringReleaseEx(key, 0)
	}
	zend.RETVAL_LONG(old_setting)
	return
}

/* }}} */

/* {{{ proto int getservbyname(string service, string protocol)
   Returns port associated with service. Protocol must be "tcp" or "udp" */

func ZifGetservbyname(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var name *byte
	var proto *byte
	var name_len int
	var proto_len int
	var serv *__struct__servent
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &name, &name_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &proto, &proto_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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

	/* empty string behaves like NULL on windows implementation of
	   getservbyname. Let be portable instead. */

	serv = getservbyname(name, proto)
	if serv == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(ntohs(serv.s_port))
	return
}

/* }}} */

/* {{{ proto string getservbyport(int port, string protocol)
   Returns service name associated with port. Protocol must be "tcp" or "udp" */

func ZifGetservbyport(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var proto *byte
	var proto_len int
	var port zend.ZendLong
	var serv *__struct__servent
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &port, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &proto, &proto_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	serv = getservbyport(htons(uint16(port)), proto)
	if serv == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STRING(serv.s_name)
	return
}

/* }}} */

/* {{{ proto int getprotobyname(string name)
   Returns protocol number associated with name as per /etc/protocols */

func ZifGetprotobyname(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var name *byte
	var name_len int
	var ent *__struct__protoent
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &name, &name_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	ent = getprotobyname(name)
	if ent == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(ent.p_proto)
	return
}

/* }}} */

/* {{{ proto string getprotobynumber(int proto)
   Returns protocol name associated with protocol number proto */

func ZifGetprotobynumber(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var proto zend.ZendLong
	var ent *__struct__protoent
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &proto, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	ent = getprotobynumber(int(proto))
	if ent == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STRING(ent.p_name)
	return
}

/* }}} */

/* {{{ proto bool register_tick_function(string function_name [, mixed arg [, mixed ... ]])
   Registers a tick callback function */

func ZifRegisterTickFunction(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var tick_fe UserTickFunctionEntry
	var i int
	var function_name *zend.ZendString = nil
	tick_fe.SetCalling(0)
	tick_fe.SetArgCount(zend.ZEND_NUM_ARGS())
	if tick_fe.GetArgCount() < 1 {
		zend.WRONG_PARAM_COUNT
	}
	tick_fe.SetArguments((*zend.Zval)(zend.SafeEmalloc(b.SizeOf("zval"), tick_fe.GetArgCount(), 0)))
	if zend.ZendGetParametersArray(zend.ZEND_NUM_ARGS(), tick_fe.GetArgCount(), tick_fe.GetArguments()) == zend.FAILURE {
		zend.Efree(tick_fe.GetArguments())
		zend.RETVAL_FALSE
		return
	}
	if zend.ZendIsCallable(&tick_fe.arguments[0], 0, &function_name) == 0 {
		zend.Efree(tick_fe.GetArguments())
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid tick callback '%s' passed", zend.ZSTR_VAL(function_name))
		zend.ZendStringReleaseEx(function_name, 0)
		zend.RETVAL_FALSE
		return
	} else if function_name != nil {
		zend.ZendStringReleaseEx(function_name, 0)
	}
	if zend.Z_TYPE(tick_fe.GetArguments()[0]) != zend.IS_ARRAY && zend.Z_TYPE(tick_fe.GetArguments()[0]) != zend.IS_OBJECT {
		zend.ConvertToStringEx(&tick_fe.arguments[0])
	}
	if !(BG(user_tick_functions)) {
		BG(user_tick_functions) = (*zend.ZendLlist)(zend.Emalloc(b.SizeOf("zend_llist")))
		zend.ZendLlistInit(BG(user_tick_functions), b.SizeOf("user_tick_function_entry"), zend.LlistDtorFuncT(UserTickFunctionDtor), 0)
		core.PhpAddTickFunction(RunUserTickFunctions, nil)
	}
	for i = 0; i < tick_fe.GetArgCount(); i++ {
		zend.Z_TRY_ADDREF(tick_fe.GetArguments()[i])
	}
	zend.ZendLlistAddElement(BG(user_tick_functions), &tick_fe)
	zend.RETVAL_TRUE
	return
}

/* }}} */

func ZifUnregisterTickFunction(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var function *zend.Zval
	var tick_fe UserTickFunctionEntry
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &function, 0)
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if !(BG(user_tick_functions)) {
		return
	}
	if zend.Z_TYPE_P(function) != zend.IS_ARRAY && zend.Z_TYPE_P(function) != zend.IS_OBJECT {
		zend.ConvertToString(function)
	}
	tick_fe.SetArguments((*zend.Zval)(zend.Emalloc(b.SizeOf("zval"))))
	zend.ZVAL_COPY_VALUE(&tick_fe.arguments[0], function)
	tick_fe.SetArgCount(1)
	zend.ZendLlistDelElement(BG(user_tick_functions), &tick_fe, (func(any, any) int)(UserTickFunctionCompare))
	zend.Efree(tick_fe.GetArguments())
}

/* }}} */

func ZifIsUploadedFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var path *byte
	var path_len int
	if !(core.SG(rfc1867_uploaded_files)) {
		zend.RETVAL_FALSE
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &path, &path_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if zend.ZendHashStrExists(core.SG(rfc1867_uploaded_files), path, path_len) != 0 {
		zend.RETVAL_TRUE
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}

/* }}} */

func ZifMoveUploadedFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var path *byte
	var new_path *byte
	var path_len int
	var new_path_len int
	var successful zend.ZendBool = 0
	var oldmask int
	var ret int
	if !(core.SG(rfc1867_uploaded_files)) {
		zend.RETVAL_FALSE
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &path, &path_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgPath(_arg, &new_path, &new_path_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if zend.ZendHashStrExists(core.SG(rfc1867_uploaded_files), path, path_len) == 0 {
		zend.RETVAL_FALSE
		return
	}
	if core.PhpCheckOpenBasedir(new_path) != 0 {
		zend.RETVAL_FALSE
		return
	}
	if zend.VCWD_RENAME(path, new_path) == 0 {
		successful = 1
		oldmask = umask(077)
		umask(oldmask)
		ret = zend.VCWD_CHMOD(new_path, 0666 & ^oldmask)
		if ret == -1 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "%s", strerror(errno))
		}
	} else if PhpCopyFileEx(path, new_path, core.STREAM_DISABLE_OPEN_BASEDIR) == zend.SUCCESS {
		zend.VCWD_UNLINK(path)
		successful = 1
	}
	if successful != 0 {
		zend.ZendHashStrDel(core.SG(rfc1867_uploaded_files), path, path_len)
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to move '%s' to '%s'", path, new_path)
	}
	zend.RETVAL_BOOL(successful != 0)
	return
}

/* }}} */

func PhpSimpleIniParserCb(arg1 *zend.Zval, arg2 *zend.Zval, arg3 *zend.Zval, callback_type int, arr *zend.Zval) {
	switch callback_type {
	case zend.ZEND_INI_PARSER_ENTRY:
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}
		zend.Z_TRY_ADDREF_P(arg2)
		zend.ZendSymtableUpdate(zend.Z_ARRVAL_P(arr), zend.Z_STR_P(arg1), arg2)
		break
	case zend.ZEND_INI_PARSER_POP_ENTRY:
		var hash zend.Zval
		var find_hash *zend.Zval
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}
		if !(zend.Z_STRLEN_P(arg1) > 1 && zend.Z_STRVAL_P(arg1)[0] == '0') && zend.IsNumericString(zend.Z_STRVAL_P(arg1), zend.Z_STRLEN_P(arg1), nil, nil, 0) == zend.IS_LONG {
			var key zend.ZendUlong = zend.ZendUlong(zend.ZendAtol(zend.Z_STRVAL_P(arg1), zend.Z_STRLEN_P(arg1)))
			if b.Assign(&find_hash, zend.ZendHashIndexFind(zend.Z_ARRVAL_P(arr), key)) == nil {
				zend.ArrayInit(&hash)
				find_hash = zend.ZendHashIndexAddNew(zend.Z_ARRVAL_P(arr), key, &hash)
			}
		} else {
			if b.Assign(&find_hash, zend.ZendHashFind(zend.Z_ARRVAL_P(arr), zend.Z_STR_P(arg1))) == nil {
				zend.ArrayInit(&hash)
				find_hash = zend.ZendHashAddNew(zend.Z_ARRVAL_P(arr), zend.Z_STR_P(arg1), &hash)
			}
		}
		if zend.Z_TYPE_P(find_hash) != zend.IS_ARRAY {
			zend.ZvalPtrDtorNogc(find_hash)
			zend.ArrayInit(find_hash)
		}
		if arg3 == nil || zend.Z_TYPE_P(arg3) == zend.IS_STRING && zend.Z_STRLEN_P(arg3) == 0 {
			zend.Z_TRY_ADDREF_P(arg2)
			zend.AddNextIndexZval(find_hash, arg2)
		} else {
			zend.ArraySetZvalKey(zend.Z_ARRVAL_P(find_hash), arg3, arg2)
		}
		break
	case zend.ZEND_INI_PARSER_SECTION:
		break
	}
}

/* }}} */

func PhpIniParserCbWithSections(arg1 *zend.Zval, arg2 *zend.Zval, arg3 *zend.Zval, callback_type int, arr *zend.Zval) {
	if callback_type == zend.ZEND_INI_PARSER_SECTION {
		zend.ArrayInit(&BG(active_ini_file_section))
		zend.ZendSymtableUpdate(zend.Z_ARRVAL_P(arr), zend.Z_STR_P(arg1), &BG(active_ini_file_section))
	} else if arg2 != nil {
		var active_arr *zend.Zval
		if zend.Z_TYPE(BG(active_ini_file_section)) != zend.IS_UNDEF {
			active_arr = &BG(active_ini_file_section)
		} else {
			active_arr = arr
		}
		PhpSimpleIniParserCb(arg1, arg2, arg3, callback_type, active_arr)
	}
}

/* }}} */

func ZifParseIniFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte = nil
	var filename_len int = 0
	var process_sections zend.ZendBool = 0
	var scanner_mode zend.ZendLong = zend.ZEND_INI_SCANNER_NORMAL
	var fh zend.ZendFileHandle
	var ini_parser_cb zend.ZendIniParserCbT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &process_sections, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &scanner_mode, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if filename_len == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Filename cannot be empty!")
		zend.RETVAL_FALSE
		return
	}

	/* Set callback function */

	if process_sections != 0 {
		zend.ZVAL_UNDEF(&BG(active_ini_file_section))
		ini_parser_cb = zend.ZendIniParserCbT(PhpIniParserCbWithSections)
	} else {
		ini_parser_cb = zend.ZendIniParserCbT(PhpSimpleIniParserCb)
	}

	/* Setup filehandle */

	zend.ZendStreamInitFilename(&fh, filename)
	zend.ArrayInit(return_value)
	if zend.ZendParseIniFile(&fh, 0, int(scanner_mode), ini_parser_cb, return_value) == zend.FAILURE {
		zend.ZendArrayDestroy(zend.Z_ARR_P(return_value))
		zend.RETVAL_FALSE
		return
	}
}

/* }}} */

func ZifParseIniString(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var string *byte = nil
	var str *byte = nil
	var str_len int = 0
	var process_sections zend.ZendBool = 0
	var scanner_mode zend.ZendLong = zend.ZEND_INI_SCANNER_NORMAL
	var ini_parser_cb zend.ZendIniParserCbT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &process_sections, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &scanner_mode, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if core.INT_MAX-str_len < zend.ZEND_MMAP_AHEAD {
		zend.RETVAL_FALSE
	}

	/* Set callback function */

	if process_sections != 0 {
		zend.ZVAL_UNDEF(&BG(active_ini_file_section))
		ini_parser_cb = zend.ZendIniParserCbT(PhpIniParserCbWithSections)
	} else {
		ini_parser_cb = zend.ZendIniParserCbT(PhpSimpleIniParserCb)
	}

	/* Setup string */

	string = (*byte)(zend.Emalloc(str_len + zend.ZEND_MMAP_AHEAD))
	memcpy(string, str, str_len)
	memset(string+str_len, 0, zend.ZEND_MMAP_AHEAD)
	zend.ArrayInit(return_value)
	if zend.ZendParseIniString(string, 0, int(scanner_mode), ini_parser_cb, return_value) == zend.FAILURE {
		zend.ZendArrayDestroy(zend.Z_ARR_P(return_value))
		zend.RETVAL_FALSE
	}
	zend.Efree(string)
}

/* }}} */

/* {{{ proto array sys_getloadavg()
 */

func ZifSysGetloadavg(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var load []float64
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if getloadavg(load, 3) == -1 {
		zend.RETVAL_FALSE
		return
	} else {
		zend.ArrayInit(return_value)
		zend.AddIndexDouble(return_value, 0, load[0])
		zend.AddIndexDouble(return_value, 1, load[1])
		zend.AddIndexDouble(return_value, 2, load[2])
	}
}

/* }}} */
