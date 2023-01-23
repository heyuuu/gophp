// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	r "sik/runtime"
	g "sik/runtime/grammar"
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

// #define basic_functions_module_ptr       & basic_functions_module

/* system functions */

var ZifSetTimeLimit func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var ZifHeaderRegisterCallback func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)

/* From the INI parser */

/* Left for BC (not binary safe!) */

// #define MT_N       ( 624 )

/* Deprecated type aliases -- use the standard types instead */

type PhpUint32 = uint32
type PhpInt32 = int32

// #define BG(v) ( basic_globals . v )

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

// #define INADDR_NONE       ( ( zend_ulong ) - 1 )

// # include "zend_globals.h"

// # include "php_globals.h"

// # include "SAPI.h"

// # include "php_ticks.h"

// # include "php_fopen_wrappers.h"

// # include "streamsfuncs.h"

var IncompleteClassEntry *zend.ZendClassEntry = nil

/* some prototypes for local functions */

/* {{{ arginfo */

var ArginfoSetTimeLimit []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"seconds", 0, 0, 0}}

/* }}} */

var ArginfoHeaderRegisterCallback []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"callback", 0, 0, 0}}

/* }}} */

var ArginfoObStart []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"user_function", 0, 0, 0}, {"chunk_size", 0, 0, 0}, {"flags", 0, 0, 0}}
var ArginfoObFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoObClean []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoObEndFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoObEndClean []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoObGetFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoObGetClean []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoObGetContents []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoObGetLevel []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoObGetLength []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoObListHandlers []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoObGetStatus []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"full_status", 0, 0, 0}}
var ArginfoObImplicitFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"flag", 0, 0, 0}}
var ArginfoOutputResetRewriteVars []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoOutputAddRewriteVar []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"name", 0, 0, 0}, {"value", 0, 0, 0}}

/* }}} */

var ArginfoStreamWrapperRegister []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"protocol", 0, 0, 0}, {"classname", 0, 0, 0}, {"flags", 0, 0, 0}}
var ArginfoStreamWrapperUnregister []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"protocol", 0, 0, 0}}
var ArginfoStreamWrapperRestore []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"protocol", 0, 0, 0}}

/* }}} */

var ArginfoKrsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"sort_flags", 0, 0, 0}}
var ArginfoKsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"sort_flags", 0, 0, 0}}
var ArginfoCount []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var", 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoNatsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 1, 0}}
var ArginfoNatcasesort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 1, 0}}
var ArginfoAsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"sort_flags", 0, 0, 0}}
var ArginfoArsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"sort_flags", 0, 0, 0}}
var ArginfoSort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"sort_flags", 0, 0, 0}}
var ArginfoRsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"sort_flags", 0, 0, 0}}
var ArginfoUsort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"cmp_function", 0, 0, 0}}
var ArginfoUasort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"cmp_function", 0, 0, 0}}
var ArginfoUksort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 1, 0}, {"cmp_function", 0, 0, 0}}
var ArginfoEnd []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 1, 0}}
var ArginfoPrev []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 1, 0}}
var ArginfoNext []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 1, 0}}
var ArginfoReset []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 1, 0}}
var ArginfoCurrent []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 0, 0}}
var ArginfoKey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 0, 0}}
var ArginfoMin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"args", 0, 0, 1}}
var ArginfoMax []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"args", 0, 0, 1}}
var ArginfoArrayWalk []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"input", 0, 1, 0}, {"funcname", 0, 0, 0}, {"userdata", 0, 0, 0}}
var ArginfoArrayWalkRecursive []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"input", 0, 1, 0}, {"funcname", 0, 0, 0}, {"userdata", 0, 0, 0}}
var ArginfoInArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"needle", 0, 0, 0}, {"haystack", 0, 0, 0}, {"strict", 0, 0, 0}}
var ArginfoArraySearch []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"needle", 0, 0, 0}, {"haystack", 0, 0, 0}, {"strict", 0, 0, 0}}
var ArginfoExtract []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 2, 0}, {"extract_type", 0, 0, 0}, {"prefix", 0, 0, 0}}
var ArginfoCompact []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var_names", 0, 0, 1}}
var ArginfoArrayFill []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"start_key", 0, 0, 0}, {"num", 0, 0, 0}, {"val", 0, 0, 0}}
var ArginfoArrayFillKeys []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"keys", 0, 0, 0}, {"val", 0, 0, 0}}
var ArginfoRange []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"low", 0, 0, 0}, {"high", 0, 0, 0}, {"step", 0, 0, 0}}
var ArginfoShuffle []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 1, 0}}
var ArginfoArrayPush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"stack", 0, 1, 0}, {"vars", 0, 0, 1}}
var ArginfoArrayPop []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"stack", 0, 1, 0}}
var ArginfoArrayShift []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"stack", 0, 1, 0}}
var ArginfoArrayUnshift []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"stack", 0, 1, 0}, {"vars", 0, 0, 1}}
var ArginfoArraySplice []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arg", 0, 1, 0}, {"offset", 0, 0, 0}, {"length", 0, 0, 0}, {"replacement", 0, 0, 0}}
var ArginfoArraySlice []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arg", 0, 0, 0}, {"offset", 0, 0, 0}, {"length", 0, 0, 0}, {"preserve_keys", 0, 0, 0}}
var ArginfoArrayMerge []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayMergeRecursive []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayReplace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayReplaceRecursive []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayKeys []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 0, 0}, {"search_value", 0, 0, 0}, {"strict", 0, 0, 0}}
var ArginfoArrayKeyFirst []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 0, 0}}
var ArginfoArrayKeyLast []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 0, 0}}
var ArginfoArrayValues []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 0, 0}}
var ArginfoArrayCountValues []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 0, 0}}
var ArginfoArrayColumn []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arg", 0, 0, 0}, {"column_key", 0, 0, 0}, {"index_key", 0, 0, 0}}
var ArginfoArrayReverse []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"input", 0, 0, 0}, {"preserve_keys", 0, 0, 0}}
var ArginfoArrayPad []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 0, 0}, {"pad_size", 0, 0, 0}, {"pad_value", 0, 0, 0}}
var ArginfoArrayFlip []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 0, 0}}
var ArginfoArrayChangeKeyCase []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"input", 0, 0, 0}, {"case", 0, 0, 0}}
var ArginfoArrayUnique []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 0, 0}, {"flags", 0, 0, 0}}
var ArginfoArrayIntersectKey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayIntersectUkey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arr2", 0, 0, 0}, {"callback_key_compare_func", 0, 0, 0}}
var ArginfoArrayIntersect []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayUintersect []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arr2", 0, 0, 0}, {"callback_data_compare_func", 0, 0, 0}}
var ArginfoArrayIntersectAssoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayUintersectAssoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arr2", 0, 0, 0}, {"callback_data_compare_func", 0, 0, 0}}
var ArginfoArrayIntersectUassoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arr2", 0, 0, 0}, {"callback_key_compare_func", 0, 0, 0}}
var ArginfoArrayUintersectUassoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arr2", 0, 0, 0}, {"callback_data_compare_func", 0, 0, 0}, {"callback_key_compare_func", 0, 0, 0}}
var ArginfoArrayDiffKey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayDiffUkey []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arr2", 0, 0, 0}, {"callback_key_comp_func", 0, 0, 0}}
var ArginfoArrayDiff []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayUdiff []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arr2", 0, 0, 0}, {"callback_data_comp_func", 0, 0, 0}}
var ArginfoArrayDiffAssoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayDiffUassoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arr2", 0, 0, 0}, {"callback_data_comp_func", 0, 0, 0}}
var ArginfoArrayUdiffAssoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arr2", 0, 0, 0}, {"callback_key_comp_func", 0, 0, 0}}
var ArginfoArrayUdiffUassoc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arr1", 0, 0, 0}, {"arr2", 0, 0, 0}, {"callback_data_comp_func", 0, 0, 0}, {"callback_key_comp_func", 0, 0, 0}}
var ArginfoArrayMultisort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arr1", 0, 2, 0}, {"sort_order", 0, 2, 0}, {"sort_flags", 0, 2, 0}, {"arr2", 0, 2, 1}}
var ArginfoArrayRand []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 0, 0}, {"num_req", 0, 0, 0}}
var ArginfoArraySum []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 0, 0}}
var ArginfoArrayProduct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 0, 0}}
var ArginfoArrayReduce []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arg", 0, 0, 0}, {"callback", 0, 0, 0}, {"initial", 0, 0, 0}}
var ArginfoArrayFilter []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"arg", 0, 0, 0}, {"callback", 0, 0, 0}, {"use_keys", 0, 0, 0}}
var ArginfoArrayMap []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"callback", 0, 0, 0}, {"arrays", 0, 0, 1}}
var ArginfoArrayKeyExists []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"key", 0, 0, 0}, {"search", 0, 0, 0}}
var ArginfoArrayChunk []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"arg", 0, 0, 0}, {"size", 0, 0, 0}, {"preserve_keys", 0, 0, 0}}
var ArginfoArrayCombine []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"keys", 0, 0, 0}, {"values", 0, 0, 0}}

/* }}} */

var ArginfoGetMagicQuotesGpc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoGetMagicQuotesRuntime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoConstant []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"const_name", 0, 0, 0}}
var ArginfoInetNtop []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"in_addr", 0, 0, 0}}
var ArginfoInetPton []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"ip_address", 0, 0, 0}}
var ArginfoIp2long []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"ip_address", 0, 0, 0}}
var ArginfoLong2ip []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"proper_address", 0, 0, 0}}
var ArginfoGetenv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"varname", 0, 0, 0}, {"local_only", 0, 0, 0}}
var ArginfoPutenv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"setting", 0, 0, 0}}
var ArginfoGetopt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"options", 0, 0, 0}, {"opts", 0, 0, 0}, {"optind", 0, 1, 0}}
var ArginfoFlush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoSleep []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"seconds", 0, 0, 0}}
var ArginfoUsleep []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"micro_seconds", 0, 0, 0}}
var ArginfoTimeNanosleep []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"seconds", 0, 0, 0}, {"nanoseconds", 0, 0, 0}}
var ArginfoTimeSleepUntil []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"timestamp", 0, 0, 0}}
var ArginfoGetCurrentUser []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoGetCfgVar []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"option_name", 0, 0, 0}}
var ArginfoErrorLog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"message", 0, 0, 0}, {"message_type", 0, 0, 0}, {"destination", 0, 0, 0}, {"extra_headers", 0, 0, 0}}
var ArginfoErrorGetLast []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}}
var ArginfoErrorClearLast []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}}
var ArginfoCallUserFunc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"function_name", 0, 0, 0}, {"parameters", 0, 0, 1}}
var ArginfoCallUserFuncArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"function_name", 0, 0, 0}, {"parameters", 0, 0, 0}}
var ArginfoForwardStaticCall []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"function_name", 0, 0, 0}, {"parameters", 0, 0, 1}}
var ArginfoForwardStaticCallArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"function_name", 0, 0, 0}, {"parameters", 0, 0, 0}}
var ArginfoRegisterShutdownFunction []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"function_name", 0, 0, 0}, {"parameters", 0, 0, 1}}
var ArginfoHighlightFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"file_name", 0, 0, 0}, {"return", 0, 0, 0}}
var ArginfoPhpStripWhitespace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"file_name", 0, 0, 0}}
var ArginfoHighlightString []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"string", 0, 0, 0}, {"return", 0, 0, 0}}
var ArginfoIniGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"varname", 0, 0, 0}}
var ArginfoIniGetAll []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"extension", 0, 0, 0}, {"details", 0, 0, 0}}
var ArginfoIniSet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"varname", 0, 0, 0}, {"newvalue", 0, 0, 0}}
var ArginfoIniRestore []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"varname", 0, 0, 0}}
var ArginfoSetIncludePath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"new_include_path", 0, 0, 0}}
var ArginfoGetIncludePath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoRestoreIncludePath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoPrintR []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var", 0, 0, 0}, {"return", 0, 0, 0}}
var ArginfoConnectionAborted []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoConnectionStatus []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoIgnoreUserAbort []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"value", 0, 0, 0}}
var ArginfoGetservbyname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"service", 0, 0, 0}, {"protocol", 0, 0, 0}}
var ArginfoGetservbyport []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"port", 0, 0, 0}, {"protocol", 0, 0, 0}}
var ArginfoGetprotobyname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"name", 0, 0, 0}}
var ArginfoGetprotobynumber []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"proto", 0, 0, 0}}
var ArginfoRegisterTickFunction []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"function_name", 0, 0, 0}, {"parameters", 0, 0, 1}}
var ArginfoUnregisterTickFunction []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"function_name", 0, 0, 0}}
var ArginfoIsUploadedFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"path", 0, 0, 0}}
var ArginfoMoveUploadedFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"path", 0, 0, 0}, {"new_path", 0, 0, 0}}
var ArginfoParseIniFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"process_sections", 0, 0, 0}, {"scanner_mode", 0, 0, 0}}
var ArginfoParseIniString []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"ini_string", 0, 0, 0}, {"process_sections", 0, 0, 0}, {"scanner_mode", 0, 0, 0}}
var ArginfoSysGetloadavg []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}

/* }}} */

var ArginfoAssert []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"assertion", 0, 0, 0}, {"description", 0, 0, 0}}
var ArginfoAssertOptions []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"what", 0, 0, 0}, {"value", 0, 0, 0}}

/* }}} */

var ArginfoBase64Encode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoBase64Decode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"strict", 0, 0, 0}}

/* }}} */

var ArginfoGetBrowser []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"browser_name", 0, 0, 0}, {"return_array", 0, 0, 0}}

/* }}} */

var ArginfoCrc32 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}

/* }}} */

var ArginfoCrypt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"salt", 0, 0, 0}}

/* }}} */

var ArginfoConvertCyrString []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}, {"from", 0, 0, 0}, {"to", 0, 0, 0}}

/* }}} */

var ArginfoStrptime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"timestamp", 0, 0, 0}, {"format", 0, 0, 0}}

/* }}} */

var ArginfoOpendir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"path", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoDir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"directory", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoClosedir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"dir_handle", 0, 0, 0}}
var ArginfoChroot []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"directory", 0, 0, 0}}
var ArginfoChdir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"directory", 0, 0, 0}}
var ArginfoGetcwd []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoRewinddir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"dir_handle", 0, 0, 0}}
var ArginfoReaddir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"dir_handle", 0, 0, 0}}
var ArginfoGlob []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"pattern", 0, 0, 0}, {"flags", 0, 0, 0}}
var ArginfoScandir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"dir", 0, 0, 0}, {"sorting_order", 0, 0, 0}, {"context", 0, 0, 0}}

/* }}} */

var ArginfoGethostbyaddr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"ip_address", 0, 0, 0}}
var ArginfoGethostbyname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"hostname", 0, 0, 0}}
var ArginfoGethostbynamel []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"hostname", 0, 0, 0}}
var ArginfoGethostname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoNetGetInterfaces []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoDnsCheckRecord []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"host", 0, 0, 0}, {"type", 0, 0, 0}}
var ArginfoDnsGetRecord []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{"hostname", 0, 0, 0},
	{"type", 0, 0, 0},
	{"authns", 7<<2 | g.Cond(true, 0x1, 0x0), 1, 0},
	{"addtl", 7<<2 | g.Cond(true, 0x1, 0x0), 1, 0},
	{"raw", 0, 0, 0},
}
var ArginfoDnsGetMx []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"hostname", 0, 0, 0}, {"mxhosts", 0, 1, 0}, {"weight", 0, 1, 0}}

/* }}} */

var ArginfoExec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"command", 0, 0, 0}, {"output", 0, 1, 0}, {"return_value", 0, 1, 0}}
var ArginfoSystem []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"command", 0, 0, 0}, {"return_value", 0, 1, 0}}
var ArginfoPassthru []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"command", 0, 0, 0}, {"return_value", 0, 1, 0}}
var ArginfoEscapeshellcmd []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"command", 0, 0, 0}}
var ArginfoEscapeshellarg []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"arg", 0, 0, 0}}
var ArginfoShellExec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"cmd", 0, 0, 0}}
var ArginfoProcNice []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"priority", 0, 0, 0}}

/* }}} */

var ArginfoFlock []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"fp", 0, 0, 0}, {"operation", 0, 0, 0}, {"wouldblock", 0, 1, 0}}
var ArginfoGetMetaTags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"use_include_path", 0, 0, 0}}
var ArginfoFileGetContents []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"flags", 0, 0, 0}, {"context", 0, 0, 0}, {"offset", 0, 0, 0}, {"maxlen", 0, 0, 0}}
var ArginfoFilePutContents []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"filename", 0, 0, 0}, {"data", 0, 0, 0}, {"flags", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"flags", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoTempnam []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"dir", 0, 0, 0}, {"prefix", 0, 0, 0}}
var ArginfoTmpfile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoFopen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"filename", 0, 0, 0}, {"mode", 0, 0, 0}, {"use_include_path", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoFclose []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}}
var ArginfoPopen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"command", 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoPclose []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}}
var ArginfoFeof []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}}
var ArginfoFgets []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"fp", 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoFgetc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}}
var ArginfoFgetss []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"fp", 0, 0, 0}, {"length", 0, 0, 0}, {"allowable_tags", 0, 0, 0}}
var ArginfoFscanf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"format", 0, 0, 0}, {"vars", 0, 1, 1}}
var ArginfoFwrite []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"fp", 0, 0, 0}, {"str", 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoFflush []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}}
var ArginfoRewind []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}}
var ArginfoFtell []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}}
var ArginfoFseek []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"fp", 0, 0, 0}, {"offset", 0, 0, 0}, {"whence", 0, 0, 0}}
var ArginfoMkdir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"pathname", 0, 0, 0}, {"mode", 0, 0, 0}, {"recursive", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoRmdir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"dirname", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoReadfile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"flags", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoUmask []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"mask", 0, 0, 0}}
var ArginfoFpassthru []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}}
var ArginfoRename []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"old_name", 0, 0, 0}, {"new_name", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoUnlink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoFtruncate []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}, {"size", 0, 0, 0}}
var ArginfoFstat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}}
var ArginfoCopy []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"source_file", 0, 0, 0}, {"destination_file", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoFread []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoFputcsv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"fp", 0, 0, 0}, {"fields", 0, 0, 0}, {"delimiter", 0, 0, 0}, {"enclosure", 0, 0, 0}, {"escape_char", 0, 0, 0}}
var ArginfoFgetcsv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"fp", 0, 0, 0}, {"length", 0, 0, 0}, {"delimiter", 0, 0, 0}, {"enclosure", 0, 0, 0}, {"escape", 0, 0, 0}}
var ArginfoRealpath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"path", 0, 0, 0}}
var ArginfoFnmatch []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"pattern", 0, 0, 0}, {"filename", 0, 0, 0}, {"flags", 0, 0, 0}}
var ArginfoSysGetTempDir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}

/* }}} */

var ArginfoDiskTotalSpace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"path", 0, 0, 0}}
var ArginfoDiskFreeSpace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"path", 0, 0, 0}}
var ArginfoChgrp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"group", 0, 0, 0}}
var ArginfoChown []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"user", 0, 0, 0}}
var ArginfoLchgrp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"group", 0, 0, 0}}
var ArginfoLchown []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"user", 0, 0, 0}}
var ArginfoChmod []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoTouch []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"time", 0, 0, 0}, {"atime", 0, 0, 0}}
var ArginfoClearstatcache []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"clear_realpath_cache", 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoRealpathCacheSize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoRealpathCacheGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoFileperms []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoFileinode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoFilesize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoFileowner []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoFilegroup []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoFileatime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoFilemtime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoFilectime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoFiletype []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoIsWritable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoIsReadable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoIsExecutable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoIsFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoIsDir []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoIsLink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoFileExists []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoLstat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoStat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}

/* }}} */

var ArginfoSprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"format", 0, 0, 0}, {"args", 0, 0, 1}}
var ArginfoVsprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"format", 0, 0, 0}, {"args", 0, 0, 0}}
var ArginfoPrintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"format", 0, 0, 0}, {"args", 0, 0, 1}}
var ArginfoVprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"format", 0, 0, 0}, {"args", 0, 0, 0}}
var ArginfoFprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"format", 0, 0, 0}, {"args", 0, 0, 1}}
var ArginfoVfprintf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"stream", 0, 0, 0}, {"format", 0, 0, 0}, {"args", 0, 0, 0}}

/* }}} */

var ArginfoFsockopen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"hostname", 0, 0, 0}, {"port", 0, 0, 0}, {"errno", 0, 1, 0}, {"errstr", 0, 1, 0}, {"timeout", 0, 0, 0}}
var ArginfoPfsockopen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"hostname", 0, 0, 0}, {"port", 0, 0, 0}, {"errno", 0, 1, 0}, {"errstr", 0, 1, 0}, {"timeout", 0, 0, 0}}

/* }}} */

var ArginfoFtok []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"pathname", 0, 0, 0}, {"proj", 0, 0, 0}}

/* }}} */

var ArginfoHeader []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"header", 0, 0, 0}, {"replace", 0, 0, 0}, {"http_response_code", 0, 0, 0}}
var ArginfoHeaderRemove []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"name", 0, 0, 0}}
var ArginfoSetcookie []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"name", 0, 0, 0}, {"value", 0, 0, 0}, {"expires_or_options", 0, 0, 0}, {"path", 0, 0, 0}, {"domain", 0, 0, 0}, {"secure", 0, 0, 0}, {"httponly", 0, 0, 0}}
var ArginfoSetrawcookie []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"name", 0, 0, 0}, {"value", 0, 0, 0}, {"expires_or_options", 0, 0, 0}, {"path", 0, 0, 0}, {"domain", 0, 0, 0}, {"secure", 0, 0, 0}, {"httponly", 0, 0, 0}}
var ArginfoHeadersSent []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"file", 0, 1, 0}, {"line", 0, 1, 0}}
var ArginfoHeadersList []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoHttpResponseCode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"response_code", 0, 0, 0}}

/* }}} */

var ArginfoHrtime []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"get_as_number", 0, 0, 0}}

/* }}} */

var ArginfoHtmlspecialchars []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"string", 0, 0, 0}, {"quote_style", 0, 0, 0}, {"encoding", 0, 0, 0}, {"double_encode", 0, 0, 0}}
var ArginfoHtmlspecialcharsDecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"string", 0, 0, 0}, {"quote_style", 0, 0, 0}}
var ArginfoHtmlEntityDecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"string", 0, 0, 0}, {"quote_style", 0, 0, 0}, {"encoding", 0, 0, 0}}
var ArginfoHtmlentities []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"string", 0, 0, 0}, {"quote_style", 0, 0, 0}, {"encoding", 0, 0, 0}, {"double_encode", 0, 0, 0}}
var ArginfoGetHtmlTranslationTable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"table", 0, 0, 0}, {"quote_style", 0, 0, 0}, {"encoding", 0, 0, 0}}

/* }}} */

var ArginfoHttpBuildQuery []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"formdata", 0, 0, 0}, {"prefix", 0, 0, 0}, {"arg_separator", 0, 0, 0}, {"enc_type", 0, 0, 0}}

/* }}} */

var ArginfoImageTypeToMimeType []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"imagetype", 0, 0, 0}}
var ArginfoImageTypeToExtension []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"imagetype", 0, 0, 0}, {"include_dot", 0, 0, 0}}
var ArginfoGetimagesize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"imagefile", 0, 0, 0}, {"info", 0, 1, 0}}

/* }}} */

var ArginfoPhpinfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"what", 0, 0, 0}}
var ArginfoPhpversion []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"extension", 0, 0, 0}}
var ArginfoPhpcredits []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"flag", 0, 0, 0}}
var ArginfoPhpSapiName []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoPhpUname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoPhpIniScannedFiles []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoPhpIniLoadedFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}

/* }}} */

var ArginfoIptcembed []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"iptcdata", 0, 0, 0}, {"jpeg_file_name", 0, 0, 0}, {"spool", 0, 0, 0}}
var ArginfoIptcparse []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"iptcdata", 0, 0, 0}}

/* }}} */

var ArginfoLcgValue []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}

/* }}} */

var ArginfoLevenshtein []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str1", 0, 0, 0}, {"str2", 0, 0, 0}, {"cost_ins", 0, 0, 0}, {"cost_rep", 0, 0, 0}, {"cost_del", 0, 0, 0}}

/* }}} */

var ArginfoReadlink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoLinkinfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoSymlink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"target", 0, 0, 0}, {"link", 0, 0, 0}}
var ArginfoLink []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"target", 0, 0, 0}, {"link", 0, 0, 0}}

/* }}} */

var ArginfoEzmlmHash []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"addr", 0, 0, 0}}
var ArginfoMail []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(3)), 0, 0, 0}, {"to", 0, 0, 0}, {"subject", 0, 0, 0}, {"message", 0, 0, 0}, {"additional_headers", 0, 0, 0}, {"additional_parameters", 0, 0, 0}}

/* }}} */

var ArginfoAbs []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoCeil []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoFloor []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoRound []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"number", 0, 0, 0}, {"precision", 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoSin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoCos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoTan []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoAsin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoAcos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoAtan []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoAtan2 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"y", 0, 0, 0}, {"x", 0, 0, 0}}
var ArginfoSinh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoCosh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoTanh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoAsinh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoAcosh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoAtanh []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoPi []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoIsFinite []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"val", 0, 0, 0}}
var ArginfoIsInfinite []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"val", 0, 0, 0}}
var ArginfoIsNan []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"val", 0, 0, 0}}
var ArginfoPow []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"base", 0, 0, 0}, {"exponent", 0, 0, 0}}
var ArginfoExp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoExpm1 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoLog1p []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoLog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"number", 0, 0, 0}, {"base", 0, 0, 0}}
var ArginfoLog10 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoSqrt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoHypot []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"num1", 0, 0, 0}, {"num2", 0, 0, 0}}
var ArginfoDeg2rad []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoRad2deg []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}}
var ArginfoBindec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"binary_number", 0, 0, 0}}
var ArginfoHexdec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"hexadecimal_number", 0, 0, 0}}
var ArginfoOctdec []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"octal_number", 0, 0, 0}}
var ArginfoDecbin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"decimal_number", 0, 0, 0}}
var ArginfoDecoct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"decimal_number", 0, 0, 0}}
var ArginfoDechex []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"decimal_number", 0, 0, 0}}
var ArginfoBaseConvert []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"number", 0, 0, 0}, {"frombase", 0, 0, 0}, {"tobase", 0, 0, 0}}
var ArginfoNumberFormat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"number", 0, 0, 0}, {"num_decimal_places", 0, 0, 0}, {"dec_separator", 0, 0, 0}, {"thousands_separator", 0, 0, 0}}
var ArginfoFmod []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"x", 0, 0, 0}, {"y", 0, 0, 0}}
var ArginfoIntdiv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"dividend", 0, 0, 0}, {"divisor", 0, 0, 0}}

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

var ArginfoGetmyuid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoGetmygid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoGetmypid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoGetmyinode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoGetlastmod []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}

/* }}} */

var ArginfoPasswordHash []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"password", 0, 0, 0}, {"algo", 0, 0, 0}, {"options", 0, 0, 0}}
var ArginfoPasswordGetInfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"hash", 0, 0, 0}}
var ArginfoPasswordNeedsRehash []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"hash", 0, 0, 0}, {"algo", 0, 0, 0}, {"options", 0, 0, 0}}
var ArginfoPasswordVerify []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"password", 0, 0, 0}, {"hash", 0, 0, 0}}
var ArginfoPasswordAlgos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}

/* }}} */

var ArginfoProcTerminate []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"process", 0, 0, 0}, {"signal", 0, 0, 0}}
var ArginfoProcClose []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"process", 0, 0, 0}}
var ArginfoProcGetStatus []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"process", 0, 0, 0}}
var ArginfoProcOpen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(3)), 0, 0, 0}, {"command", 0, 0, 0}, {"descriptorspec", 0, 0, 0}, {"pipes", 0, 1, 0}, {"cwd", 0, 0, 0}, {"env", 0, 0, 0}, {"other_options", 0, 0, 0}}

/* }}} */

var ArginfoQuotedPrintableDecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}

/* }}} */

var ArginfoQuotedPrintableEncode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}

/* }}} */

var ArginfoMtSrand []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"seed", 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoMtRand []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"min", 0, 0, 0}, {"max", 0, 0, 0}}
var ArginfoMtGetrandmax []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}

/* }}} */

var ArginfoRandomBytes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoRandomInt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"min", 0, 0, 0}, {"max", 0, 0, 0}}

/* }}} */

var ArginfoSha1 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"raw_output", 0, 0, 0}}
var ArginfoSha1File []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"filename", 0, 0, 0}, {"raw_output", 0, 0, 0}}

/* }}} */

var ArginfoSoundex []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}

/* }}} */

var ArginfoStreamSocketPair []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"domain", 0, 0, 0}, {"type", 0, 0, 0}, {"protocol", 0, 0, 0}}
var ArginfoStreamSocketClient []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"remoteaddress", 0, 0, 0}, {"errcode", 0, 1, 0}, {"errstring", 0, 1, 0}, {"timeout", 0, 0, 0}, {"flags", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoStreamSocketServer []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"localaddress", 0, 0, 0}, {"errcode", 0, 1, 0}, {"errstring", 0, 1, 0}, {"flags", 0, 0, 0}, {"context", 0, 0, 0}}
var ArginfoStreamSocketAccept []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"serverstream", 0, 0, 0}, {"timeout", 0, 0, 0}, {"peername", 0, 1, 0}}
var ArginfoStreamSocketGetName []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"stream", 0, 0, 0}, {"want_peer", 0, 0, 0}}
var ArginfoStreamSocketSendto []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"data", 0, 0, 0}, {"flags", 0, 0, 0}, {"target_addr", 0, 0, 0}}
var ArginfoStreamSocketRecvfrom []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"amount", 0, 0, 0}, {"flags", 0, 0, 0}, {"remote_addr", 0, 1, 0}}
var ArginfoStreamGetContents []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"source", 0, 0, 0}, {"maxlen", 0, 0, 0}, {"offset", 0, 0, 0}}
var ArginfoStreamCopyToStream []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"source", 0, 0, 0}, {"dest", 0, 0, 0}, {"maxlen", 0, 0, 0}, {"pos", 0, 0, 0}}
var ArginfoStreamGetMetaData []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}}
var ArginfoStreamGetTransports []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoStreamGetWrappers []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoStreamResolveIncludePath []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filename", 0, 0, 0}}
var ArginfoStreamIsLocal []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"stream", 0, 0, 0}}
var ArginfoStreamSupportsLock []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"stream", 0, 0, 0}}
var ArginfoStreamIsatty []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"stream", 0, 0, 0}}
var ArginfoStreamSelect []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(4)), 0, 0, 0}, {"read_streams", 0, 1, 0}, {"write_streams", 0, 1, 0}, {"except_streams", 0, 1, 0}, {"tv_sec", 0, 0, 0}, {"tv_usec", 0, 0, 0}}
var ArginfoStreamContextGetOptions []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"stream_or_context", 0, 0, 0}}
var ArginfoStreamContextSetOption []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream_or_context", 0, 0, 0}, {"wrappername", 0, 0, 0}, {"optionname", 0, 0, 0}, {"value", 0, 0, 0}}
var ArginfoStreamContextSetParams []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"stream_or_context", 0, 0, 0}, {"options", 0, 0, 0}}
var ArginfoStreamContextGetParams []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"stream_or_context", 0, 0, 0}}
var ArginfoStreamContextGetDefault []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"options", 0, 0, 0}}
var ArginfoStreamContextSetDefault []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"options", 0, 0, 0}}
var ArginfoStreamContextCreate []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"options", 0, 0, 0}, {"params", 0, 0, 0}}
var ArginfoStreamFilterPrepend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"filtername", 0, 0, 0}, {"read_write", 0, 0, 0}, {"filterparams", 0, 0, 0}}
var ArginfoStreamFilterAppend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"filtername", 0, 0, 0}, {"read_write", 0, 0, 0}, {"filterparams", 0, 0, 0}}
var ArginfoStreamFilterRemove []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"stream_filter", 0, 0, 0}}
var ArginfoStreamGetLine []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"maxlen", 0, 0, 0}, {"ending", 0, 0, 0}}
var ArginfoStreamSetBlocking []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"socket", 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoStreamSetTimeout []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"seconds", 0, 0, 0}, {"microseconds", 0, 0, 0}}
var ArginfoStreamSetReadBuffer []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}, {"buffer", 0, 0, 0}}
var ArginfoStreamSetWriteBuffer []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}, {"buffer", 0, 0, 0}}
var ArginfoStreamSetChunkSize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"fp", 0, 0, 0}, {"chunk_size", 0, 0, 0}}
var ArginfoStreamSocketEnableCrypto []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"stream", 0, 0, 0}, {"enable", 0, 0, 0}, {"cryptokind", 0, 0, 0}, {"sessionstream", 0, 0, 0}}
var ArginfoStreamSocketShutdown []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"stream", 0, 0, 0}, {"how", 0, 0, 0}}

/* }}} */

var ArginfoBin2hex []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"data", 0, 0, 0}}
var ArginfoHex2bin []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"data", 0, 0, 0}}
var ArginfoStrspn []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str", 0, 0, 0}, {"mask", 0, 0, 0}, {"start", 0, 0, 0}, {"len", 0, 0, 0}}
var ArginfoStrcspn []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str", 0, 0, 0}, {"mask", 0, 0, 0}, {"start", 0, 0, 0}, {"len", 0, 0, 0}}
var ArginfoStrcoll []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str1", 0, 0, 0}, {"str2", 0, 0, 0}}
var ArginfoTrim []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"character_mask", 0, 0, 0}}
var ArginfoRtrim []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"character_mask", 0, 0, 0}}
var ArginfoLtrim []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"character_mask", 0, 0, 0}}
var ArginfoWordwrap []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"width", 0, 0, 0}, {"break", 0, 0, 0}, {"cut", 0, 0, 0}}
var ArginfoExplode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"separator", 0, 0, 0}, {"str", 0, 0, 0}, {"limit", 0, 0, 0}}
var ArginfoImplode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"glue", 0, 0, 0}, {"pieces", 0, 0, 0}}
var ArginfoStrtok []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"token", 0, 0, 0}}
var ArginfoStrtoupper []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoStrtolower []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoBasename []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"path", 0, 0, 0}, {"suffix", 0, 0, 0}}
var ArginfoDirname []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"path", 0, 0, 0}, {"levels", 0, 0, 0}}
var ArginfoPathinfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"path", 0, 0, 0}, {"options", 0, 0, 0}}
var ArginfoStristr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"part", 0, 0, 0}}
var ArginfoStrstr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"part", 0, 0, 0}}
var ArginfoStrpos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"offset", 0, 0, 0}}
var ArginfoStripos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"offset", 0, 0, 0}}
var ArginfoStrrpos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"offset", 0, 0, 0}}
var ArginfoStrripos []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"offset", 0, 0, 0}}
var ArginfoStrrchr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}}
var ArginfoChunkSplit []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"chunklen", 0, 0, 0}, {"ending", 0, 0, 0}}
var ArginfoSubstr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str", 0, 0, 0}, {"start", 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoSubstrReplace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(3)), 0, 0, 0}, {"str", 0, 0, 0}, {"replace", 0, 0, 0}, {"start", 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoQuotemeta []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoOrd []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"character", 0, 0, 0}}
var ArginfoChr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"codepoint", 0, 0, 0}}
var ArginfoUcfirst []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoLcfirst []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoUcwords []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"delimiters", 0, 0, 0}}
var ArginfoStrtr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str", 0, 0, 0}, {"from", 0, 0, 0}, {"to", 0, 0, 0}}
var ArginfoStrrev []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoSimilarText []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str1", 0, 0, 0}, {"str2", 0, 0, 0}, {"percent", 0, 1, 0}}
var ArginfoAddcslashes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}, {"charlist", 0, 0, 0}}
var ArginfoAddslashes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoStripcslashes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoStripslashes []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoStrReplace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(3)), 0, 0, 0}, {"search", 0, 0, 0}, {"replace", 0, 0, 0}, {"subject", 0, 0, 0}, {"replace_count", 0, 1, 0}}
var ArginfoStrIreplace []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(3)), 0, 0, 0}, {"search", 0, 0, 0}, {"replace", 0, 0, 0}, {"subject", 0, 0, 0}, {"replace_count", 0, 1, 0}}
var ArginfoHebrev []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"max_chars_per_line", 0, 0, 0}}
var ArginfoHebrevc []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"max_chars_per_line", 0, 0, 0}}
var ArginfoNl2br []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"is_xhtml", 0, 0, 0}}
var ArginfoStripTags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"allowable_tags", 0, 0, 0}}
var ArginfoSetlocale []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"category", 0, 0, 0}, {"locales", 0, 0, 1}}
var ArginfoParseStr []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"encoded_string", 0, 0, 0}, {"result", 0, 1, 0}}
var ArginfoStrGetcsv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"string", 0, 0, 0}, {"delimiter", 0, 0, 0}, {"enclosure", 0, 0, 0}, {"escape", 0, 0, 0}}
var ArginfoStrRepeat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"input", 0, 0, 0}, {"mult", 0, 0, 0}}
var ArginfoCountChars []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"input", 0, 0, 0}, {"mode", 0, 0, 0}}
var ArginfoStrnatcmp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"s1", 0, 0, 0}, {"s2", 0, 0, 0}}
var ArginfoLocaleconv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoStrnatcasecmp []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"s1", 0, 0, 0}, {"s2", 0, 0, 0}}
var ArginfoSubstrCount []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"needle", 0, 0, 0}, {"offset", 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoStrPad []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"input", 0, 0, 0}, {"pad_length", 0, 0, 0}, {"pad_string", 0, 0, 0}, {"pad_type", 0, 0, 0}}
var ArginfoSscanf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"str", 0, 0, 0}, {"format", 0, 0, 0}, {"vars", 0, 1, 1}}
var ArginfoStrRot13 []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoStrShuffle []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoStrWordCount []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"format", 0, 0, 0}, {"charlist", 0, 0, 0}}
var ArginfoMoneyFormat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"format", 0, 0, 0}, {"value", 0, 0, 0}}
var ArginfoStrSplit []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"split_length", 0, 0, 0}}
var ArginfoStrpbrk []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(2)), 0, 0, 0}, {"haystack", 0, 0, 0}, {"char_list", 0, 0, 0}}
var ArginfoSubstrCompare []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(3)), 0, 0, 0}, {"main_str", 0, 0, 0}, {"str", 0, 0, 0}, {"offset", 0, 0, 0}, {"length", 0, 0, 0}, {"case_sensitivity", 0, 0, 0}}
var ArginfoUtf8Encode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"data", 0, 0, 0}}
var ArginfoUtf8Decode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"data", 0, 0, 0}}

/* }}} */

var ArginfoOpenlog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"ident", 0, 0, 0}, {"option", 0, 0, 0}, {"facility", 0, 0, 0}}
var ArginfoCloselog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoSyslog []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"priority", 0, 0, 0}, {"message", 0, 0, 0}}

/* }}} */

var ArginfoGettype []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoSettype []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 1, 0}, {"type", 0, 0, 0}}
var ArginfoIntval []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var", 0, 0, 0}, {"base", 0, 0, 0}}
var ArginfoFloatval []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoStrval []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoBoolval []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoIsNull []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoIsResource []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoIsBool []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoIsInt []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoIsFloat []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoIsString []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoIsArray []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoIsObject []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoIsNumeric []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"value", 0, 0, 0}}
var ArginfoIsScalar []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"value", 0, 0, 0}}
var ArginfoIsCallable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var", 0, 0, 0}, {"syntax_only", 0, 0, 0}, {"callable_name", 0, 1, 0}}
var ArginfoIsIterable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var", 0, 0, 0}}
var ArginfoIsCountable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}

/* }}} */

var ArginfoUniqid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"prefix", 0, 0, 0}, {"more_entropy", 0, 0, 0}}

/* }}} */

var ArginfoParseUrl []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"url", 0, 0, 0}, {"component", 0, 0, 0}}
var ArginfoUrlencode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoUrldecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoRawurlencode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoRawurldecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"str", 0, 0, 0}}
var ArginfoGetHeaders []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"url", 0, 0, 0}, {"format", 0, 0, 0}, {"context", 0, 0, 0}}

/* }}} */

var ArginfoStreamBucketMakeWriteable []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"brigade", 0, 0, 0}}
var ArginfoStreamBucketPrepend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"brigade", 0, 0, 0}, {"bucket", 0, 0, 0}}
var ArginfoStreamBucketAppend []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"brigade", 0, 0, 0}, {"bucket", 0, 0, 0}}
var ArginfoStreamBucketNew []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"stream", 0, 0, 0}, {"buffer", 0, 0, 0}}
var ArginfoStreamGetFilters []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ArginfoStreamFilterRegister []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"filtername", 0, 0, 0}, {"classname", 0, 0, 0}}

/* }}} */

var ArginfoConvertUuencode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"data", 0, 0, 0}}
var ArginfoConvertUudecode []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"data", 0, 0, 0}}

/* }}} */

var ArginfoVarDump []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"vars", 0, 0, 1}}
var ArginfoDebugZvalDump []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"vars", 0, 0, 1}}
var ArginfoVarExport []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"var", 0, 0, 0}, {"return", 0, 0, 0}}
var ArginfoSerialize []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"var", 0, 0, 0}}
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
		uint32(g.SizeOf("arginfo_constant")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"bin2hex",
		ZifBin2hex,
		ArginfoBin2hex,
		uint32(g.SizeOf("arginfo_bin2hex")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"hex2bin",
		ZifHex2bin,
		ArginfoHex2bin,
		uint32(g.SizeOf("arginfo_hex2bin")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sleep",
		ZifSleep,
		ArginfoSleep,
		uint32(g.SizeOf("arginfo_sleep")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"usleep",
		ZifUsleep,
		ArginfoUsleep,
		uint32(g.SizeOf("arginfo_usleep")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"time_nanosleep",
		ZifTimeNanosleep,
		ArginfoTimeNanosleep,
		uint32(g.SizeOf("arginfo_time_nanosleep")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"time_sleep_until",
		ZifTimeSleepUntil,
		ArginfoTimeSleepUntil,
		uint32(g.SizeOf("arginfo_time_sleep_until")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strptime",
		ZifStrptime,
		ArginfoStrptime,
		uint32(g.SizeOf("arginfo_strptime")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"flush",
		ZifFlush,
		ArginfoFlush,
		uint32(g.SizeOf("arginfo_flush")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"wordwrap",
		ZifWordwrap,
		ArginfoWordwrap,
		uint32(g.SizeOf("arginfo_wordwrap")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"htmlspecialchars",
		ZifHtmlspecialchars,
		ArginfoHtmlspecialchars,
		uint32(g.SizeOf("arginfo_htmlspecialchars")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"htmlentities",
		ZifHtmlentities,
		ArginfoHtmlentities,
		uint32(g.SizeOf("arginfo_htmlentities")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"html_entity_decode",
		ZifHtmlEntityDecode,
		ArginfoHtmlEntityDecode,
		uint32(g.SizeOf("arginfo_html_entity_decode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"htmlspecialchars_decode",
		ZifHtmlspecialcharsDecode,
		ArginfoHtmlspecialcharsDecode,
		uint32(g.SizeOf("arginfo_htmlspecialchars_decode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_html_translation_table",
		ZifGetHtmlTranslationTable,
		ArginfoGetHtmlTranslationTable,
		uint32(g.SizeOf("arginfo_get_html_translation_table")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sha1",
		ZifSha1,
		ArginfoSha1,
		uint32(g.SizeOf("arginfo_sha1")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sha1_file",
		ZifSha1File,
		ArginfoSha1File,
		uint32(g.SizeOf("arginfo_sha1_file")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"md5",
		PhpIfMd5,
		ArginfoMd5,
		uint32(g.SizeOf("arginfo_md5")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"md5_file",
		PhpIfMd5File,
		ArginfoMd5File,
		uint32(g.SizeOf("arginfo_md5_file")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"crc32",
		PhpIfCrc32,
		ArginfoCrc32,
		uint32(g.SizeOf("arginfo_crc32")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"iptcparse",
		ZifIptcparse,
		ArginfoIptcparse,
		uint32(g.SizeOf("arginfo_iptcparse")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"iptcembed",
		ZifIptcembed,
		ArginfoIptcembed,
		uint32(g.SizeOf("arginfo_iptcembed")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getimagesize",
		ZifGetimagesize,
		ArginfoGetimagesize,
		uint32(g.SizeOf("arginfo_getimagesize")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getimagesizefromstring",
		ZifGetimagesizefromstring,
		ArginfoGetimagesize,
		uint32(g.SizeOf("arginfo_getimagesize")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"image_type_to_mime_type",
		ZifImageTypeToMimeType,
		ArginfoImageTypeToMimeType,
		uint32(g.SizeOf("arginfo_image_type_to_mime_type")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"image_type_to_extension",
		ZifImageTypeToExtension,
		ArginfoImageTypeToExtension,
		uint32(g.SizeOf("arginfo_image_type_to_extension")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"phpinfo",
		ZifPhpinfo,
		ArginfoPhpinfo,
		uint32(g.SizeOf("arginfo_phpinfo")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"phpversion",
		ZifPhpversion,
		ArginfoPhpversion,
		uint32(g.SizeOf("arginfo_phpversion")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"phpcredits",
		ZifPhpcredits,
		ArginfoPhpcredits,
		uint32(g.SizeOf("arginfo_phpcredits")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"php_sapi_name",
		ZifPhpSapiName,
		ArginfoPhpSapiName,
		uint32(g.SizeOf("arginfo_php_sapi_name")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"php_uname",
		ZifPhpUname,
		ArginfoPhpUname,
		uint32(g.SizeOf("arginfo_php_uname")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"php_ini_scanned_files",
		ZifPhpIniScannedFiles,
		ArginfoPhpIniScannedFiles,
		uint32(g.SizeOf("arginfo_php_ini_scanned_files")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"php_ini_loaded_file",
		ZifPhpIniLoadedFile,
		ArginfoPhpIniLoadedFile,
		uint32(g.SizeOf("arginfo_php_ini_loaded_file")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strnatcmp",
		ZifStrnatcmp,
		ArginfoStrnatcmp,
		uint32(g.SizeOf("arginfo_strnatcmp")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strnatcasecmp",
		ZifStrnatcasecmp,
		ArginfoStrnatcasecmp,
		uint32(g.SizeOf("arginfo_strnatcasecmp")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"substr_count",
		ZifSubstrCount,
		ArginfoSubstrCount,
		uint32(g.SizeOf("arginfo_substr_count")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strspn",
		ZifStrspn,
		ArginfoStrspn,
		uint32(g.SizeOf("arginfo_strspn")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strcspn",
		ZifStrcspn,
		ArginfoStrcspn,
		uint32(g.SizeOf("arginfo_strcspn")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strtok",
		ZifStrtok,
		ArginfoStrtok,
		uint32(g.SizeOf("arginfo_strtok")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strtoupper",
		ZifStrtoupper,
		ArginfoStrtoupper,
		uint32(g.SizeOf("arginfo_strtoupper")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strtolower",
		ZifStrtolower,
		ArginfoStrtolower,
		uint32(g.SizeOf("arginfo_strtolower")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strpos",
		ZifStrpos,
		ArginfoStrpos,
		uint32(g.SizeOf("arginfo_strpos")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stripos",
		ZifStripos,
		ArginfoStripos,
		uint32(g.SizeOf("arginfo_stripos")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strrpos",
		ZifStrrpos,
		ArginfoStrrpos,
		uint32(g.SizeOf("arginfo_strrpos")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strripos",
		ZifStrripos,
		ArginfoStrripos,
		uint32(g.SizeOf("arginfo_strripos")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strrev",
		ZifStrrev,
		ArginfoStrrev,
		uint32(g.SizeOf("arginfo_strrev")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"hebrev",
		ZifHebrev,
		ArginfoHebrev,
		uint32(g.SizeOf("arginfo_hebrev")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"hebrevc",
		ZifHebrevc,
		ArginfoHebrevc,
		uint32(g.SizeOf("arginfo_hebrevc")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 11,
	},
	{
		"nl2br",
		ZifNl2br,
		ArginfoNl2br,
		uint32(g.SizeOf("arginfo_nl2br")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"basename",
		ZifBasename,
		ArginfoBasename,
		uint32(g.SizeOf("arginfo_basename")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"dirname",
		ZifDirname,
		ArginfoDirname,
		uint32(g.SizeOf("arginfo_dirname")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pathinfo",
		ZifPathinfo,
		ArginfoPathinfo,
		uint32(g.SizeOf("arginfo_pathinfo")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stripslashes",
		ZifStripslashes,
		ArginfoStripslashes,
		uint32(g.SizeOf("arginfo_stripslashes")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stripcslashes",
		ZifStripcslashes,
		ArginfoStripcslashes,
		uint32(g.SizeOf("arginfo_stripcslashes")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strstr",
		ZifStrstr,
		ArginfoStrstr,
		uint32(g.SizeOf("arginfo_strstr")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stristr",
		ZifStristr,
		ArginfoStristr,
		uint32(g.SizeOf("arginfo_stristr")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strrchr",
		ZifStrrchr,
		ArginfoStrrchr,
		uint32(g.SizeOf("arginfo_strrchr")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_shuffle",
		ZifStrShuffle,
		ArginfoStrShuffle,
		uint32(g.SizeOf("arginfo_str_shuffle")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_word_count",
		ZifStrWordCount,
		ArginfoStrWordCount,
		uint32(g.SizeOf("arginfo_str_word_count")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_split",
		ZifStrSplit,
		ArginfoStrSplit,
		uint32(g.SizeOf("arginfo_str_split")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strpbrk",
		ZifStrpbrk,
		ArginfoStrpbrk,
		uint32(g.SizeOf("arginfo_strpbrk")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"substr_compare",
		ZifSubstrCompare,
		ArginfoSubstrCompare,
		uint32(g.SizeOf("arginfo_substr_compare")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"utf8_encode",
		ZifUtf8Encode,
		ArginfoUtf8Encode,
		uint32(g.SizeOf("arginfo_utf8_encode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"utf8_decode",
		ZifUtf8Decode,
		ArginfoUtf8Decode,
		uint32(g.SizeOf("arginfo_utf8_decode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strcoll",
		ZifStrcoll,
		ArginfoStrcoll,
		uint32(g.SizeOf("arginfo_strcoll")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"money_format",
		ZifMoneyFormat,
		ArginfoMoneyFormat,
		uint32(g.SizeOf("arginfo_money_format")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 11,
	},
	{
		"substr",
		ZifSubstr,
		ArginfoSubstr,
		uint32(g.SizeOf("arginfo_substr")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"substr_replace",
		ZifSubstrReplace,
		ArginfoSubstrReplace,
		uint32(g.SizeOf("arginfo_substr_replace")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"quotemeta",
		ZifQuotemeta,
		ArginfoQuotemeta,
		uint32(g.SizeOf("arginfo_quotemeta")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ucfirst",
		ZifUcfirst,
		ArginfoUcfirst,
		uint32(g.SizeOf("arginfo_ucfirst")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"lcfirst",
		ZifLcfirst,
		ArginfoLcfirst,
		uint32(g.SizeOf("arginfo_lcfirst")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ucwords",
		ZifUcwords,
		ArginfoUcwords,
		uint32(g.SizeOf("arginfo_ucwords")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strtr",
		ZifStrtr,
		ArginfoStrtr,
		uint32(g.SizeOf("arginfo_strtr")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"addslashes",
		ZifAddslashes,
		ArginfoAddslashes,
		uint32(g.SizeOf("arginfo_addslashes")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"addcslashes",
		ZifAddcslashes,
		ArginfoAddcslashes,
		uint32(g.SizeOf("arginfo_addcslashes")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rtrim",
		ZifRtrim,
		ArginfoRtrim,
		uint32(g.SizeOf("arginfo_rtrim")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_replace",
		ZifStrReplace,
		ArginfoStrReplace,
		uint32(g.SizeOf("arginfo_str_replace")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_ireplace",
		ZifStrIreplace,
		ArginfoStrIreplace,
		uint32(g.SizeOf("arginfo_str_ireplace")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_repeat",
		ZifStrRepeat,
		ArginfoStrRepeat,
		uint32(g.SizeOf("arginfo_str_repeat")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"count_chars",
		ZifCountChars,
		ArginfoCountChars,
		uint32(g.SizeOf("arginfo_count_chars")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chunk_split",
		ZifChunkSplit,
		ArginfoChunkSplit,
		uint32(g.SizeOf("arginfo_chunk_split")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"trim",
		ZifTrim,
		ArginfoTrim,
		uint32(g.SizeOf("arginfo_trim")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ltrim",
		ZifLtrim,
		ArginfoLtrim,
		uint32(g.SizeOf("arginfo_ltrim")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strip_tags",
		ZifStripTags,
		ArginfoStripTags,
		uint32(g.SizeOf("arginfo_strip_tags")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"similar_text",
		ZifSimilarText,
		ArginfoSimilarText,
		uint32(g.SizeOf("arginfo_similar_text")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"explode",
		ZifExplode,
		ArginfoExplode,
		uint32(g.SizeOf("arginfo_explode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"implode",
		ZifImplode,
		ArginfoImplode,
		uint32(g.SizeOf("arginfo_implode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"join",
		ZifImplode,
		ArginfoImplode,
		uint32(g.SizeOf("arginfo_implode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"setlocale",
		ZifSetlocale,
		ArginfoSetlocale,
		uint32(g.SizeOf("arginfo_setlocale")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"localeconv",
		ZifLocaleconv,
		ArginfoLocaleconv,
		uint32(g.SizeOf("arginfo_localeconv")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"soundex",
		ZifSoundex,
		ArginfoSoundex,
		uint32(g.SizeOf("arginfo_soundex")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"levenshtein",
		ZifLevenshtein,
		ArginfoLevenshtein,
		uint32(g.SizeOf("arginfo_levenshtein")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chr",
		ZifChr,
		ArginfoChr,
		uint32(g.SizeOf("arginfo_chr")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ord",
		ZifOrd,
		ArginfoOrd,
		uint32(g.SizeOf("arginfo_ord")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"parse_str",
		ZifParseStr,
		ArginfoParseStr,
		uint32(g.SizeOf("arginfo_parse_str")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_getcsv",
		ZifStrGetcsv,
		ArginfoStrGetcsv,
		uint32(g.SizeOf("arginfo_str_getcsv")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_pad",
		ZifStrPad,
		ArginfoStrPad,
		uint32(g.SizeOf("arginfo_str_pad")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chop",
		ZifRtrim,
		ArginfoRtrim,
		uint32(g.SizeOf("arginfo_rtrim")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strchr",
		ZifStrstr,
		ArginfoStrstr,
		uint32(g.SizeOf("arginfo_strstr")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sprintf",
		ZifUserSprintf,
		ArginfoSprintf,
		uint32(g.SizeOf("arginfo_sprintf")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"printf",
		ZifUserPrintf,
		ArginfoPrintf,
		uint32(g.SizeOf("arginfo_printf")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"vprintf",
		ZifVprintf,
		ArginfoVprintf,
		uint32(g.SizeOf("arginfo_vprintf")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"vsprintf",
		ZifVsprintf,
		ArginfoVsprintf,
		uint32(g.SizeOf("arginfo_vsprintf")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fprintf",
		ZifFprintf,
		ArginfoFprintf,
		uint32(g.SizeOf("arginfo_fprintf")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"vfprintf",
		ZifVfprintf,
		ArginfoVfprintf,
		uint32(g.SizeOf("arginfo_vfprintf")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sscanf",
		ZifSscanf,
		ArginfoSscanf,
		uint32(g.SizeOf("arginfo_sscanf")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fscanf",
		ZifFscanf,
		ArginfoFscanf,
		uint32(g.SizeOf("arginfo_fscanf")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"parse_url",
		ZifParseUrl,
		ArginfoParseUrl,
		uint32(g.SizeOf("arginfo_parse_url")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"urlencode",
		ZifUrlencode,
		ArginfoUrlencode,
		uint32(g.SizeOf("arginfo_urlencode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"urldecode",
		ZifUrldecode,
		ArginfoUrldecode,
		uint32(g.SizeOf("arginfo_urldecode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rawurlencode",
		ZifRawurlencode,
		ArginfoRawurlencode,
		uint32(g.SizeOf("arginfo_rawurlencode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rawurldecode",
		ZifRawurldecode,
		ArginfoRawurldecode,
		uint32(g.SizeOf("arginfo_rawurldecode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"http_build_query",
		ZifHttpBuildQuery,
		ArginfoHttpBuildQuery,
		uint32(g.SizeOf("arginfo_http_build_query")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"readlink",
		ZifReadlink,
		ArginfoReadlink,
		uint32(g.SizeOf("arginfo_readlink")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"linkinfo",
		ZifLinkinfo,
		ArginfoLinkinfo,
		uint32(g.SizeOf("arginfo_linkinfo")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"symlink",
		ZifSymlink,
		ArginfoSymlink,
		uint32(g.SizeOf("arginfo_symlink")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"link",
		ZifLink,
		ArginfoLink,
		uint32(g.SizeOf("arginfo_link")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"unlink",
		ZifUnlink,
		ArginfoUnlink,
		uint32(g.SizeOf("arginfo_unlink")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"exec",
		ZifExec,
		ArginfoExec,
		uint32(g.SizeOf("arginfo_exec")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"system",
		ZifSystem,
		ArginfoSystem,
		uint32(g.SizeOf("arginfo_system")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"escapeshellcmd",
		ZifEscapeshellcmd,
		ArginfoEscapeshellcmd,
		uint32(g.SizeOf("arginfo_escapeshellcmd")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"escapeshellarg",
		ZifEscapeshellarg,
		ArginfoEscapeshellarg,
		uint32(g.SizeOf("arginfo_escapeshellarg")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"passthru",
		ZifPassthru,
		ArginfoPassthru,
		uint32(g.SizeOf("arginfo_passthru")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"shell_exec",
		ZifShellExec,
		ArginfoShellExec,
		uint32(g.SizeOf("arginfo_shell_exec")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"proc_open",
		ZifProcOpen,
		ArginfoProcOpen,
		uint32(g.SizeOf("arginfo_proc_open")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"proc_close",
		ZifProcClose,
		ArginfoProcClose,
		uint32(g.SizeOf("arginfo_proc_close")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"proc_terminate",
		ZifProcTerminate,
		ArginfoProcTerminate,
		uint32(g.SizeOf("arginfo_proc_terminate")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"proc_get_status",
		ZifProcGetStatus,
		ArginfoProcGetStatus,
		uint32(g.SizeOf("arginfo_proc_get_status")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"proc_nice",
		ZifProcNice,
		ArginfoProcNice,
		uint32(g.SizeOf("arginfo_proc_nice")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rand",
		ZifRand,
		ArginfoMtRand,
		uint32(g.SizeOf("arginfo_mt_rand")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"srand",
		ZifMtSrand,
		ArginfoMtSrand,
		uint32(g.SizeOf("arginfo_mt_srand")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getrandmax",
		ZifMtGetrandmax,
		ArginfoMtGetrandmax,
		uint32(g.SizeOf("arginfo_mt_getrandmax")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"mt_rand",
		ZifMtRand,
		ArginfoMtRand,
		uint32(g.SizeOf("arginfo_mt_rand")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"mt_srand",
		ZifMtSrand,
		ArginfoMtSrand,
		uint32(g.SizeOf("arginfo_mt_srand")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"mt_getrandmax",
		ZifMtGetrandmax,
		ArginfoMtGetrandmax,
		uint32(g.SizeOf("arginfo_mt_getrandmax")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"random_bytes",
		ZifRandomBytes,
		ArginfoRandomBytes,
		uint32(g.SizeOf("arginfo_random_bytes")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"random_int",
		ZifRandomInt,
		ArginfoRandomInt,
		uint32(g.SizeOf("arginfo_random_int")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getservbyname",
		ZifGetservbyname,
		ArginfoGetservbyname,
		uint32(g.SizeOf("arginfo_getservbyname")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getservbyport",
		ZifGetservbyport,
		ArginfoGetservbyport,
		uint32(g.SizeOf("arginfo_getservbyport")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getprotobyname",
		ZifGetprotobyname,
		ArginfoGetprotobyname,
		uint32(g.SizeOf("arginfo_getprotobyname")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getprotobynumber",
		ZifGetprotobynumber,
		ArginfoGetprotobynumber,
		uint32(g.SizeOf("arginfo_getprotobynumber")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getmyuid",
		ZifGetmyuid,
		ArginfoGetmyuid,
		uint32(g.SizeOf("arginfo_getmyuid")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getmygid",
		ZifGetmygid,
		ArginfoGetmygid,
		uint32(g.SizeOf("arginfo_getmygid")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getmypid",
		ZifGetmypid,
		ArginfoGetmypid,
		uint32(g.SizeOf("arginfo_getmypid")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getmyinode",
		ZifGetmyinode,
		ArginfoGetmyinode,
		uint32(g.SizeOf("arginfo_getmyinode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getlastmod",
		ZifGetlastmod,
		ArginfoGetlastmod,
		uint32(g.SizeOf("arginfo_getlastmod")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"base64_decode",
		ZifBase64Decode,
		ArginfoBase64Decode,
		uint32(g.SizeOf("arginfo_base64_decode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"base64_encode",
		ZifBase64Encode,
		ArginfoBase64Encode,
		uint32(g.SizeOf("arginfo_base64_encode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"password_hash",
		ZifPasswordHash,
		ArginfoPasswordHash,
		uint32(g.SizeOf("arginfo_password_hash")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"password_get_info",
		ZifPasswordGetInfo,
		ArginfoPasswordGetInfo,
		uint32(g.SizeOf("arginfo_password_get_info")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"password_needs_rehash",
		ZifPasswordNeedsRehash,
		ArginfoPasswordNeedsRehash,
		uint32(g.SizeOf("arginfo_password_needs_rehash")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"password_verify",
		ZifPasswordVerify,
		ArginfoPasswordVerify,
		uint32(g.SizeOf("arginfo_password_verify")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"password_algos",
		ZifPasswordAlgos,
		ArginfoPasswordAlgos,
		uint32(g.SizeOf("arginfo_password_algos")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"convert_uuencode",
		ZifConvertUuencode,
		ArginfoConvertUuencode,
		uint32(g.SizeOf("arginfo_convert_uuencode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"convert_uudecode",
		ZifConvertUudecode,
		ArginfoConvertUudecode,
		uint32(g.SizeOf("arginfo_convert_uudecode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"abs",
		ZifAbs,
		ArginfoAbs,
		uint32(g.SizeOf("arginfo_abs")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ceil",
		ZifCeil,
		ArginfoCeil,
		uint32(g.SizeOf("arginfo_ceil")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"floor",
		ZifFloor,
		ArginfoFloor,
		uint32(g.SizeOf("arginfo_floor")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"round",
		ZifRound,
		ArginfoRound,
		uint32(g.SizeOf("arginfo_round")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sin",
		ZifSin,
		ArginfoSin,
		uint32(g.SizeOf("arginfo_sin")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"cos",
		ZifCos,
		ArginfoCos,
		uint32(g.SizeOf("arginfo_cos")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"tan",
		ZifTan,
		ArginfoTan,
		uint32(g.SizeOf("arginfo_tan")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"asin",
		ZifAsin,
		ArginfoAsin,
		uint32(g.SizeOf("arginfo_asin")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"acos",
		ZifAcos,
		ArginfoAcos,
		uint32(g.SizeOf("arginfo_acos")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"atan",
		ZifAtan,
		ArginfoAtan,
		uint32(g.SizeOf("arginfo_atan")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"atanh",
		ZifAtanh,
		ArginfoAtanh,
		uint32(g.SizeOf("arginfo_atanh")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"atan2",
		ZifAtan2,
		ArginfoAtan2,
		uint32(g.SizeOf("arginfo_atan2")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sinh",
		ZifSinh,
		ArginfoSinh,
		uint32(g.SizeOf("arginfo_sinh")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"cosh",
		ZifCosh,
		ArginfoCosh,
		uint32(g.SizeOf("arginfo_cosh")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"tanh",
		ZifTanh,
		ArginfoTanh,
		uint32(g.SizeOf("arginfo_tanh")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"asinh",
		ZifAsinh,
		ArginfoAsinh,
		uint32(g.SizeOf("arginfo_asinh")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"acosh",
		ZifAcosh,
		ArginfoAcosh,
		uint32(g.SizeOf("arginfo_acosh")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"expm1",
		ZifExpm1,
		ArginfoExpm1,
		uint32(g.SizeOf("arginfo_expm1")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"log1p",
		ZifLog1p,
		ArginfoLog1p,
		uint32(g.SizeOf("arginfo_log1p")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pi",
		ZifPi,
		ArginfoPi,
		uint32(g.SizeOf("arginfo_pi")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_finite",
		ZifIsFinite,
		ArginfoIsFinite,
		uint32(g.SizeOf("arginfo_is_finite")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_nan",
		ZifIsNan,
		ArginfoIsNan,
		uint32(g.SizeOf("arginfo_is_nan")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_infinite",
		ZifIsInfinite,
		ArginfoIsInfinite,
		uint32(g.SizeOf("arginfo_is_infinite")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pow",
		ZifPow,
		ArginfoPow,
		uint32(g.SizeOf("arginfo_pow")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"exp",
		ZifExp,
		ArginfoExp,
		uint32(g.SizeOf("arginfo_exp")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"log",
		ZifLog,
		ArginfoLog,
		uint32(g.SizeOf("arginfo_log")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"log10",
		ZifLog10,
		ArginfoLog10,
		uint32(g.SizeOf("arginfo_log10")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sqrt",
		ZifSqrt,
		ArginfoSqrt,
		uint32(g.SizeOf("arginfo_sqrt")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"hypot",
		ZifHypot,
		ArginfoHypot,
		uint32(g.SizeOf("arginfo_hypot")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"deg2rad",
		ZifDeg2rad,
		ArginfoDeg2rad,
		uint32(g.SizeOf("arginfo_deg2rad")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rad2deg",
		ZifRad2deg,
		ArginfoRad2deg,
		uint32(g.SizeOf("arginfo_rad2deg")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"bindec",
		ZifBindec,
		ArginfoBindec,
		uint32(g.SizeOf("arginfo_bindec")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"hexdec",
		ZifHexdec,
		ArginfoHexdec,
		uint32(g.SizeOf("arginfo_hexdec")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"octdec",
		ZifOctdec,
		ArginfoOctdec,
		uint32(g.SizeOf("arginfo_octdec")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"decbin",
		ZifDecbin,
		ArginfoDecbin,
		uint32(g.SizeOf("arginfo_decbin")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"decoct",
		ZifDecoct,
		ArginfoDecoct,
		uint32(g.SizeOf("arginfo_decoct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"dechex",
		ZifDechex,
		ArginfoDechex,
		uint32(g.SizeOf("arginfo_dechex")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"base_convert",
		ZifBaseConvert,
		ArginfoBaseConvert,
		uint32(g.SizeOf("arginfo_base_convert")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"number_format",
		ZifNumberFormat,
		ArginfoNumberFormat,
		uint32(g.SizeOf("arginfo_number_format")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fmod",
		ZifFmod,
		ArginfoFmod,
		uint32(g.SizeOf("arginfo_fmod")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"intdiv",
		ZifIntdiv,
		ArginfoIntdiv,
		uint32(g.SizeOf("arginfo_intdiv")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"inet_ntop",
		ZifInetNtop,
		ArginfoInetNtop,
		uint32(g.SizeOf("arginfo_inet_ntop")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"inet_pton",
		PhpInetPton,
		ArginfoInetPton,
		uint32(g.SizeOf("arginfo_inet_pton")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ip2long",
		ZifIp2long,
		ArginfoIp2long,
		uint32(g.SizeOf("arginfo_ip2long")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"long2ip",
		ZifLong2ip,
		ArginfoLong2ip,
		uint32(g.SizeOf("arginfo_long2ip")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getenv",
		ZifGetenv,
		ArginfoGetenv,
		uint32(g.SizeOf("arginfo_getenv")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"putenv",
		ZifPutenv,
		ArginfoPutenv,
		uint32(g.SizeOf("arginfo_putenv")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getopt",
		ZifGetopt,
		ArginfoGetopt,
		uint32(g.SizeOf("arginfo_getopt")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sys_getloadavg",
		ZifSysGetloadavg,
		ArginfoSysGetloadavg,
		uint32(g.SizeOf("arginfo_sys_getloadavg")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"microtime",
		ZifMicrotime,
		ArginfoMicrotime,
		uint32(g.SizeOf("arginfo_microtime")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gettimeofday",
		ZifGettimeofday,
		ArginfoGettimeofday,
		uint32(g.SizeOf("arginfo_gettimeofday")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getrusage",
		ZifGetrusage,
		ArginfoGetrusage,
		uint32(g.SizeOf("arginfo_getrusage")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"hrtime",
		ZifHrtime,
		ArginfoHrtime,
		uint32(g.SizeOf("arginfo_hrtime")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"uniqid",
		ZifUniqid,
		ArginfoUniqid,
		uint32(g.SizeOf("arginfo_uniqid")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"quoted_printable_decode",
		ZifQuotedPrintableDecode,
		ArginfoQuotedPrintableDecode,
		uint32(g.SizeOf("arginfo_quoted_printable_decode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"quoted_printable_encode",
		ZifQuotedPrintableEncode,
		ArginfoQuotedPrintableEncode,
		uint32(g.SizeOf("arginfo_quoted_printable_encode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"convert_cyr_string",
		ZifConvertCyrString,
		ArginfoConvertCyrString,
		uint32(g.SizeOf("arginfo_convert_cyr_string")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 11,
	},
	{
		"get_current_user",
		ZifGetCurrentUser,
		ArginfoGetCurrentUser,
		uint32(g.SizeOf("arginfo_get_current_user")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"set_time_limit",
		ZifSetTimeLimit,
		ArginfoSetTimeLimit,
		uint32(g.SizeOf("arginfo_set_time_limit")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"header_register_callback",
		ZifHeaderRegisterCallback,
		ArginfoHeaderRegisterCallback,
		uint32(g.SizeOf("arginfo_header_register_callback")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_cfg_var",
		ZifGetCfgVar,
		ArginfoGetCfgVar,
		uint32(g.SizeOf("arginfo_get_cfg_var")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_magic_quotes_gpc",
		ZifGetMagicQuotesGpc,
		ArginfoGetMagicQuotesGpc,
		uint32(g.SizeOf("arginfo_get_magic_quotes_gpc")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 11,
	},
	{
		"get_magic_quotes_runtime",
		ZifGetMagicQuotesRuntime,
		ArginfoGetMagicQuotesRuntime,
		uint32(g.SizeOf("arginfo_get_magic_quotes_runtime")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 11,
	},
	{
		"error_log",
		ZifErrorLog,
		ArginfoErrorLog,
		uint32(g.SizeOf("arginfo_error_log")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"error_get_last",
		ZifErrorGetLast,
		ArginfoErrorGetLast,
		uint32(g.SizeOf("arginfo_error_get_last")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"error_clear_last",
		ZifErrorClearLast,
		ArginfoErrorClearLast,
		uint32(g.SizeOf("arginfo_error_clear_last")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"call_user_func",
		ZifCallUserFunc,
		ArginfoCallUserFunc,
		uint32(g.SizeOf("arginfo_call_user_func")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"call_user_func_array",
		ZifCallUserFuncArray,
		ArginfoCallUserFuncArray,
		uint32(g.SizeOf("arginfo_call_user_func_array")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"forward_static_call",
		ZifForwardStaticCall,
		ArginfoForwardStaticCall,
		uint32(g.SizeOf("arginfo_forward_static_call")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"forward_static_call_array",
		ZifForwardStaticCallArray,
		ArginfoForwardStaticCallArray,
		uint32(g.SizeOf("arginfo_forward_static_call_array")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"serialize",
		ZifSerialize,
		ArginfoSerialize,
		uint32(g.SizeOf("arginfo_serialize")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"unserialize",
		ZifUnserialize,
		ArginfoUnserialize,
		uint32(g.SizeOf("arginfo_unserialize")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"var_dump",
		ZifVarDump,
		ArginfoVarDump,
		uint32(g.SizeOf("arginfo_var_dump")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"var_export",
		ZifVarExport,
		ArginfoVarExport,
		uint32(g.SizeOf("arginfo_var_export")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"debug_zval_dump",
		ZifDebugZvalDump,
		ArginfoDebugZvalDump,
		uint32(g.SizeOf("arginfo_debug_zval_dump")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"print_r",
		ZifPrintR,
		ArginfoPrintR,
		uint32(g.SizeOf("arginfo_print_r")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"memory_get_usage",
		ZifMemoryGetUsage,
		ArginfoMemoryGetUsage,
		uint32(g.SizeOf("arginfo_memory_get_usage")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"memory_get_peak_usage",
		ZifMemoryGetPeakUsage,
		ArginfoMemoryGetPeakUsage,
		uint32(g.SizeOf("arginfo_memory_get_peak_usage")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"register_shutdown_function",
		ZifRegisterShutdownFunction,
		ArginfoRegisterShutdownFunction,
		uint32(g.SizeOf("arginfo_register_shutdown_function")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"register_tick_function",
		ZifRegisterTickFunction,
		ArginfoRegisterTickFunction,
		uint32(g.SizeOf("arginfo_register_tick_function")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"unregister_tick_function",
		ZifUnregisterTickFunction,
		ArginfoUnregisterTickFunction,
		uint32(g.SizeOf("arginfo_unregister_tick_function")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"highlight_file",
		ZifHighlightFile,
		ArginfoHighlightFile,
		uint32(g.SizeOf("arginfo_highlight_file")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"show_source",
		ZifHighlightFile,
		ArginfoHighlightFile,
		uint32(g.SizeOf("arginfo_highlight_file")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"highlight_string",
		ZifHighlightString,
		ArginfoHighlightString,
		uint32(g.SizeOf("arginfo_highlight_string")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"php_strip_whitespace",
		ZifPhpStripWhitespace,
		ArginfoPhpStripWhitespace,
		uint32(g.SizeOf("arginfo_php_strip_whitespace")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ini_get",
		ZifIniGet,
		ArginfoIniGet,
		uint32(g.SizeOf("arginfo_ini_get")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ini_get_all",
		ZifIniGetAll,
		ArginfoIniGetAll,
		uint32(g.SizeOf("arginfo_ini_get_all")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ini_set",
		ZifIniSet,
		ArginfoIniSet,
		uint32(g.SizeOf("arginfo_ini_set")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ini_alter",
		ZifIniSet,
		ArginfoIniSet,
		uint32(g.SizeOf("arginfo_ini_set")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ini_restore",
		ZifIniRestore,
		ArginfoIniRestore,
		uint32(g.SizeOf("arginfo_ini_restore")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_include_path",
		ZifGetIncludePath,
		ArginfoGetIncludePath,
		uint32(g.SizeOf("arginfo_get_include_path")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"set_include_path",
		ZifSetIncludePath,
		ArginfoSetIncludePath,
		uint32(g.SizeOf("arginfo_set_include_path")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"restore_include_path",
		ZifRestoreIncludePath,
		ArginfoRestoreIncludePath,
		uint32(g.SizeOf("arginfo_restore_include_path")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 11,
	},
	{
		"setcookie",
		ZifSetcookie,
		ArginfoSetcookie,
		uint32(g.SizeOf("arginfo_setcookie")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"setrawcookie",
		ZifSetrawcookie,
		ArginfoSetrawcookie,
		uint32(g.SizeOf("arginfo_setrawcookie")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"header",
		ZifHeader,
		ArginfoHeader,
		uint32(g.SizeOf("arginfo_header")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"header_remove",
		ZifHeaderRemove,
		ArginfoHeaderRemove,
		uint32(g.SizeOf("arginfo_header_remove")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"headers_sent",
		ZifHeadersSent,
		ArginfoHeadersSent,
		uint32(g.SizeOf("arginfo_headers_sent")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"headers_list",
		ZifHeadersList,
		ArginfoHeadersList,
		uint32(g.SizeOf("arginfo_headers_list")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"http_response_code",
		ZifHttpResponseCode,
		ArginfoHttpResponseCode,
		uint32(g.SizeOf("arginfo_http_response_code")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"connection_aborted",
		ZifConnectionAborted,
		ArginfoConnectionAborted,
		uint32(g.SizeOf("arginfo_connection_aborted")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"connection_status",
		ZifConnectionStatus,
		ArginfoConnectionStatus,
		uint32(g.SizeOf("arginfo_connection_status")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ignore_user_abort",
		ZifIgnoreUserAbort,
		ArginfoIgnoreUserAbort,
		uint32(g.SizeOf("arginfo_ignore_user_abort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"parse_ini_file",
		ZifParseIniFile,
		ArginfoParseIniFile,
		uint32(g.SizeOf("arginfo_parse_ini_file")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"parse_ini_string",
		ZifParseIniString,
		ArginfoParseIniString,
		uint32(g.SizeOf("arginfo_parse_ini_string")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_uploaded_file",
		ZifIsUploadedFile,
		ArginfoIsUploadedFile,
		uint32(g.SizeOf("arginfo_is_uploaded_file")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"move_uploaded_file",
		ZifMoveUploadedFile,
		ArginfoMoveUploadedFile,
		uint32(g.SizeOf("arginfo_move_uploaded_file")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gethostbyaddr",
		ZifGethostbyaddr,
		ArginfoGethostbyaddr,
		uint32(g.SizeOf("arginfo_gethostbyaddr")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gethostbyname",
		ZifGethostbyname,
		ArginfoGethostbyname,
		uint32(g.SizeOf("arginfo_gethostbyname")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gethostbynamel",
		ZifGethostbynamel,
		ArginfoGethostbynamel,
		uint32(g.SizeOf("arginfo_gethostbynamel")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gethostname",
		ZifGethostname,
		ArginfoGethostname,
		uint32(g.SizeOf("arginfo_gethostname")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"net_get_interfaces",
		ZifNetGetInterfaces,
		ArginfoNetGetInterfaces,
		uint32(g.SizeOf("arginfo_net_get_interfaces")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"dns_check_record",
		ZifDnsCheckRecord,
		ArginfoDnsCheckRecord,
		uint32(g.SizeOf("arginfo_dns_check_record")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"checkdnsrr",
		ZifDnsCheckRecord,
		ArginfoDnsCheckRecord,
		uint32(g.SizeOf("arginfo_dns_check_record")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"dns_get_mx",
		ZifDnsGetMx,
		ArginfoDnsGetMx,
		uint32(g.SizeOf("arginfo_dns_get_mx")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getmxrr",
		ZifDnsGetMx,
		ArginfoDnsGetMx,
		uint32(g.SizeOf("arginfo_dns_get_mx")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"dns_get_record",
		ZifDnsGetRecord,
		ArginfoDnsGetRecord,
		uint32(g.SizeOf("arginfo_dns_get_record")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"intval",
		ZifIntval,
		ArginfoIntval,
		uint32(g.SizeOf("arginfo_intval")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"floatval",
		ZifFloatval,
		ArginfoFloatval,
		uint32(g.SizeOf("arginfo_floatval")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"doubleval",
		ZifFloatval,
		ArginfoFloatval,
		uint32(g.SizeOf("arginfo_floatval")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"strval",
		ZifStrval,
		ArginfoStrval,
		uint32(g.SizeOf("arginfo_strval")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"boolval",
		ZifBoolval,
		ArginfoBoolval,
		uint32(g.SizeOf("arginfo_boolval")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"gettype",
		ZifGettype,
		ArginfoGettype,
		uint32(g.SizeOf("arginfo_gettype")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"settype",
		ZifSettype,
		ArginfoSettype,
		uint32(g.SizeOf("arginfo_settype")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_null",
		ZifIsNull,
		ArginfoIsNull,
		uint32(g.SizeOf("arginfo_is_null")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_resource",
		ZifIsResource,
		ArginfoIsResource,
		uint32(g.SizeOf("arginfo_is_resource")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_bool",
		ZifIsBool,
		ArginfoIsBool,
		uint32(g.SizeOf("arginfo_is_bool")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_int",
		ZifIsInt,
		ArginfoIsInt,
		uint32(g.SizeOf("arginfo_is_int")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_float",
		ZifIsFloat,
		ArginfoIsFloat,
		uint32(g.SizeOf("arginfo_is_float")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_integer",
		ZifIsInt,
		ArginfoIsInt,
		uint32(g.SizeOf("arginfo_is_int")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_long",
		ZifIsInt,
		ArginfoIsInt,
		uint32(g.SizeOf("arginfo_is_int")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_double",
		ZifIsFloat,
		ArginfoIsFloat,
		uint32(g.SizeOf("arginfo_is_float")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_real",
		ZifIsFloat,
		ArginfoIsFloat,
		uint32(g.SizeOf("arginfo_is_float")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 11,
	},
	{
		"is_numeric",
		ZifIsNumeric,
		ArginfoIsNumeric,
		uint32(g.SizeOf("arginfo_is_numeric")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_string",
		ZifIsString,
		ArginfoIsString,
		uint32(g.SizeOf("arginfo_is_string")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_array",
		ZifIsArray,
		ArginfoIsArray,
		uint32(g.SizeOf("arginfo_is_array")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_object",
		ZifIsObject,
		ArginfoIsObject,
		uint32(g.SizeOf("arginfo_is_object")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_scalar",
		ZifIsScalar,
		ArginfoIsScalar,
		uint32(g.SizeOf("arginfo_is_scalar")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_callable",
		ZifIsCallable,
		ArginfoIsCallable,
		uint32(g.SizeOf("arginfo_is_callable")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_iterable",
		ZifIsIterable,
		ArginfoIsIterable,
		uint32(g.SizeOf("arginfo_is_iterable")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_countable",
		ZifIsCountable,
		ArginfoIsCountable,
		uint32(g.SizeOf("arginfo_is_countable")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pclose",
		ZifPclose,
		ArginfoPclose,
		uint32(g.SizeOf("arginfo_pclose")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"popen",
		ZifPopen,
		ArginfoPopen,
		uint32(g.SizeOf("arginfo_popen")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"readfile",
		ZifReadfile,
		ArginfoReadfile,
		uint32(g.SizeOf("arginfo_readfile")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rewind",
		ZifRewind,
		ArginfoRewind,
		uint32(g.SizeOf("arginfo_rewind")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rmdir",
		ZifRmdir,
		ArginfoRmdir,
		uint32(g.SizeOf("arginfo_rmdir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"umask",
		ZifUmask,
		ArginfoUmask,
		uint32(g.SizeOf("arginfo_umask")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fclose",
		ZifFclose,
		ArginfoFclose,
		uint32(g.SizeOf("arginfo_fclose")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"feof",
		ZifFeof,
		ArginfoFeof,
		uint32(g.SizeOf("arginfo_feof")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fgetc",
		ZifFgetc,
		ArginfoFgetc,
		uint32(g.SizeOf("arginfo_fgetc")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fgets",
		ZifFgets,
		ArginfoFgets,
		uint32(g.SizeOf("arginfo_fgets")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fgetss",
		ZifFgetss,
		ArginfoFgetss,
		uint32(g.SizeOf("arginfo_fgetss")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 11,
	},
	{
		"fread",
		ZifFread,
		ArginfoFread,
		uint32(g.SizeOf("arginfo_fread")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fopen",
		PhpIfFopen,
		ArginfoFopen,
		uint32(g.SizeOf("arginfo_fopen")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fpassthru",
		ZifFpassthru,
		ArginfoFpassthru,
		uint32(g.SizeOf("arginfo_fpassthru")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ftruncate",
		PhpIfFtruncate,
		ArginfoFtruncate,
		uint32(g.SizeOf("arginfo_ftruncate")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fstat",
		PhpIfFstat,
		ArginfoFstat,
		uint32(g.SizeOf("arginfo_fstat")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fseek",
		ZifFseek,
		ArginfoFseek,
		uint32(g.SizeOf("arginfo_fseek")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ftell",
		ZifFtell,
		ArginfoFtell,
		uint32(g.SizeOf("arginfo_ftell")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fflush",
		ZifFflush,
		ArginfoFflush,
		uint32(g.SizeOf("arginfo_fflush")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fwrite",
		ZifFwrite,
		ArginfoFwrite,
		uint32(g.SizeOf("arginfo_fwrite")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fputs",
		ZifFwrite,
		ArginfoFwrite,
		uint32(g.SizeOf("arginfo_fwrite")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"mkdir",
		ZifMkdir,
		ArginfoMkdir,
		uint32(g.SizeOf("arginfo_mkdir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rename",
		ZifRename,
		ArginfoRename,
		uint32(g.SizeOf("arginfo_rename")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"copy",
		ZifCopy,
		ArginfoCopy,
		uint32(g.SizeOf("arginfo_copy")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"tempnam",
		ZifTempnam,
		ArginfoTempnam,
		uint32(g.SizeOf("arginfo_tempnam")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"tmpfile",
		PhpIfTmpfile,
		ArginfoTmpfile,
		uint32(g.SizeOf("arginfo_tmpfile")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"file",
		ZifFile,
		ArginfoFile,
		uint32(g.SizeOf("arginfo_file")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"file_get_contents",
		ZifFileGetContents,
		ArginfoFileGetContents,
		uint32(g.SizeOf("arginfo_file_get_contents")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"file_put_contents",
		ZifFilePutContents,
		ArginfoFilePutContents,
		uint32(g.SizeOf("arginfo_file_put_contents")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_select",
		ZifStreamSelect,
		ArginfoStreamSelect,
		uint32(g.SizeOf("arginfo_stream_select")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_create",
		ZifStreamContextCreate,
		ArginfoStreamContextCreate,
		uint32(g.SizeOf("arginfo_stream_context_create")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_set_params",
		ZifStreamContextSetParams,
		ArginfoStreamContextSetParams,
		uint32(g.SizeOf("arginfo_stream_context_set_params")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_get_params",
		ZifStreamContextGetParams,
		ArginfoStreamContextGetParams,
		uint32(g.SizeOf("arginfo_stream_context_get_params")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_set_option",
		ZifStreamContextSetOption,
		ArginfoStreamContextSetOption,
		uint32(g.SizeOf("arginfo_stream_context_set_option")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_get_options",
		ZifStreamContextGetOptions,
		ArginfoStreamContextGetOptions,
		uint32(g.SizeOf("arginfo_stream_context_get_options")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_get_default",
		ZifStreamContextGetDefault,
		ArginfoStreamContextGetDefault,
		uint32(g.SizeOf("arginfo_stream_context_get_default")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_context_set_default",
		ZifStreamContextSetDefault,
		ArginfoStreamContextSetDefault,
		uint32(g.SizeOf("arginfo_stream_context_set_default")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_filter_prepend",
		ZifStreamFilterPrepend,
		ArginfoStreamFilterPrepend,
		uint32(g.SizeOf("arginfo_stream_filter_prepend")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_filter_append",
		ZifStreamFilterAppend,
		ArginfoStreamFilterAppend,
		uint32(g.SizeOf("arginfo_stream_filter_append")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_filter_remove",
		ZifStreamFilterRemove,
		ArginfoStreamFilterRemove,
		uint32(g.SizeOf("arginfo_stream_filter_remove")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_client",
		ZifStreamSocketClient,
		ArginfoStreamSocketClient,
		uint32(g.SizeOf("arginfo_stream_socket_client")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_server",
		ZifStreamSocketServer,
		ArginfoStreamSocketServer,
		uint32(g.SizeOf("arginfo_stream_socket_server")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_accept",
		ZifStreamSocketAccept,
		ArginfoStreamSocketAccept,
		uint32(g.SizeOf("arginfo_stream_socket_accept")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_get_name",
		ZifStreamSocketGetName,
		ArginfoStreamSocketGetName,
		uint32(g.SizeOf("arginfo_stream_socket_get_name")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_recvfrom",
		ZifStreamSocketRecvfrom,
		ArginfoStreamSocketRecvfrom,
		uint32(g.SizeOf("arginfo_stream_socket_recvfrom")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_sendto",
		ZifStreamSocketSendto,
		ArginfoStreamSocketSendto,
		uint32(g.SizeOf("arginfo_stream_socket_sendto")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_enable_crypto",
		ZifStreamSocketEnableCrypto,
		ArginfoStreamSocketEnableCrypto,
		uint32(g.SizeOf("arginfo_stream_socket_enable_crypto")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_shutdown",
		ZifStreamSocketShutdown,
		ArginfoStreamSocketShutdown,
		uint32(g.SizeOf("arginfo_stream_socket_shutdown")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_socket_pair",
		ZifStreamSocketPair,
		ArginfoStreamSocketPair,
		uint32(g.SizeOf("arginfo_stream_socket_pair")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_copy_to_stream",
		ZifStreamCopyToStream,
		ArginfoStreamCopyToStream,
		uint32(g.SizeOf("arginfo_stream_copy_to_stream")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_get_contents",
		ZifStreamGetContents,
		ArginfoStreamGetContents,
		uint32(g.SizeOf("arginfo_stream_get_contents")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_supports_lock",
		ZifStreamSupportsLock,
		ArginfoStreamSupportsLock,
		uint32(g.SizeOf("arginfo_stream_supports_lock")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_isatty",
		ZifStreamIsatty,
		ArginfoStreamIsatty,
		uint32(g.SizeOf("arginfo_stream_isatty")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fgetcsv",
		ZifFgetcsv,
		ArginfoFgetcsv,
		uint32(g.SizeOf("arginfo_fgetcsv")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fputcsv",
		ZifFputcsv,
		ArginfoFputcsv,
		uint32(g.SizeOf("arginfo_fputcsv")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"flock",
		ZifFlock,
		ArginfoFlock,
		uint32(g.SizeOf("arginfo_flock")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_meta_tags",
		ZifGetMetaTags,
		ArginfoGetMetaTags,
		uint32(g.SizeOf("arginfo_get_meta_tags")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_set_read_buffer",
		ZifStreamSetReadBuffer,
		ArginfoStreamSetReadBuffer,
		uint32(g.SizeOf("arginfo_stream_set_read_buffer")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_set_write_buffer",
		ZifStreamSetWriteBuffer,
		ArginfoStreamSetWriteBuffer,
		uint32(g.SizeOf("arginfo_stream_set_write_buffer")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"set_file_buffer",
		ZifStreamSetWriteBuffer,
		ArginfoStreamSetWriteBuffer,
		uint32(g.SizeOf("arginfo_stream_set_write_buffer")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_set_chunk_size",
		ZifStreamSetChunkSize,
		ArginfoStreamSetChunkSize,
		uint32(g.SizeOf("arginfo_stream_set_chunk_size")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_set_blocking",
		ZifStreamSetBlocking,
		ArginfoStreamSetBlocking,
		uint32(g.SizeOf("arginfo_stream_set_blocking")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"socket_set_blocking",
		ZifStreamSetBlocking,
		ArginfoStreamSetBlocking,
		uint32(g.SizeOf("arginfo_stream_set_blocking")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_get_meta_data",
		ZifStreamGetMetaData,
		ArginfoStreamGetMetaData,
		uint32(g.SizeOf("arginfo_stream_get_meta_data")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_get_line",
		ZifStreamGetLine,
		ArginfoStreamGetLine,
		uint32(g.SizeOf("arginfo_stream_get_line")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_wrapper_register",
		ZifStreamWrapperRegister,
		ArginfoStreamWrapperRegister,
		uint32(g.SizeOf("arginfo_stream_wrapper_register")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_register_wrapper",
		ZifStreamWrapperRegister,
		ArginfoStreamWrapperRegister,
		uint32(g.SizeOf("arginfo_stream_wrapper_register")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_wrapper_unregister",
		ZifStreamWrapperUnregister,
		ArginfoStreamWrapperUnregister,
		uint32(g.SizeOf("arginfo_stream_wrapper_unregister")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_wrapper_restore",
		ZifStreamWrapperRestore,
		ArginfoStreamWrapperRestore,
		uint32(g.SizeOf("arginfo_stream_wrapper_restore")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_get_wrappers",
		ZifStreamGetWrappers,
		ArginfoStreamGetWrappers,
		uint32(g.SizeOf("arginfo_stream_get_wrappers")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_get_transports",
		ZifStreamGetTransports,
		ArginfoStreamGetTransports,
		uint32(g.SizeOf("arginfo_stream_get_transports")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_resolve_include_path",
		ZifStreamResolveIncludePath,
		ArginfoStreamResolveIncludePath,
		uint32(g.SizeOf("arginfo_stream_resolve_include_path")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_is_local",
		ZifStreamIsLocal,
		ArginfoStreamIsLocal,
		uint32(g.SizeOf("arginfo_stream_is_local")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_headers",
		ZifGetHeaders,
		ArginfoGetHeaders,
		uint32(g.SizeOf("arginfo_get_headers")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_set_timeout",
		ZifStreamSetTimeout,
		ArginfoStreamSetTimeout,
		uint32(g.SizeOf("arginfo_stream_set_timeout")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"socket_set_timeout",
		ZifStreamSetTimeout,
		ArginfoStreamSetTimeout,
		uint32(g.SizeOf("arginfo_stream_set_timeout")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"socket_get_status",
		ZifStreamGetMetaData,
		ArginfoStreamGetMetaData,
		uint32(g.SizeOf("arginfo_stream_get_meta_data")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"realpath",
		ZifRealpath,
		ArginfoRealpath,
		uint32(g.SizeOf("arginfo_realpath")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fnmatch",
		ZifFnmatch,
		ArginfoFnmatch,
		uint32(g.SizeOf("arginfo_fnmatch")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fsockopen",
		ZifFsockopen,
		ArginfoFsockopen,
		uint32(g.SizeOf("arginfo_fsockopen")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pfsockopen",
		ZifPfsockopen,
		ArginfoPfsockopen,
		uint32(g.SizeOf("arginfo_pfsockopen")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pack",
		ZifPack,
		ArginfoPack,
		uint32(g.SizeOf("arginfo_pack")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"unpack",
		ZifUnpack,
		ArginfoUnpack,
		uint32(g.SizeOf("arginfo_unpack")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"get_browser",
		ZifGetBrowser,
		ArginfoGetBrowser,
		uint32(g.SizeOf("arginfo_get_browser")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"crypt",
		ZifCrypt,
		ArginfoCrypt,
		uint32(g.SizeOf("arginfo_crypt")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"opendir",
		ZifOpendir,
		ArginfoOpendir,
		uint32(g.SizeOf("arginfo_opendir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"closedir",
		ZifClosedir,
		ArginfoClosedir,
		uint32(g.SizeOf("arginfo_closedir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chdir",
		ZifChdir,
		ArginfoChdir,
		uint32(g.SizeOf("arginfo_chdir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chroot",
		ZifChroot,
		ArginfoChroot,
		uint32(g.SizeOf("arginfo_chroot")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getcwd",
		ZifGetcwd,
		ArginfoGetcwd,
		uint32(g.SizeOf("arginfo_getcwd")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rewinddir",
		ZifRewinddir,
		ArginfoRewinddir,
		uint32(g.SizeOf("arginfo_rewinddir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"readdir",
		PhpIfReaddir,
		ArginfoReaddir,
		uint32(g.SizeOf("arginfo_readdir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"dir",
		ZifGetdir,
		ArginfoDir,
		uint32(g.SizeOf("arginfo_dir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"scandir",
		ZifScandir,
		ArginfoScandir,
		uint32(g.SizeOf("arginfo_scandir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"glob",
		ZifGlob,
		ArginfoGlob,
		uint32(g.SizeOf("arginfo_glob")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fileatime",
		ZifFileatime,
		ArginfoFileatime,
		uint32(g.SizeOf("arginfo_fileatime")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"filectime",
		ZifFilectime,
		ArginfoFilectime,
		uint32(g.SizeOf("arginfo_filectime")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"filegroup",
		ZifFilegroup,
		ArginfoFilegroup,
		uint32(g.SizeOf("arginfo_filegroup")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fileinode",
		ZifFileinode,
		ArginfoFileinode,
		uint32(g.SizeOf("arginfo_fileinode")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"filemtime",
		ZifFilemtime,
		ArginfoFilemtime,
		uint32(g.SizeOf("arginfo_filemtime")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fileowner",
		ZifFileowner,
		ArginfoFileowner,
		uint32(g.SizeOf("arginfo_fileowner")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"fileperms",
		ZifFileperms,
		ArginfoFileperms,
		uint32(g.SizeOf("arginfo_fileperms")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"filesize",
		ZifFilesize,
		ArginfoFilesize,
		uint32(g.SizeOf("arginfo_filesize")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"filetype",
		ZifFiletype,
		ArginfoFiletype,
		uint32(g.SizeOf("arginfo_filetype")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"file_exists",
		ZifFileExists,
		ArginfoFileExists,
		uint32(g.SizeOf("arginfo_file_exists")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_writable",
		ZifIsWritable,
		ArginfoIsWritable,
		uint32(g.SizeOf("arginfo_is_writable")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_writeable",
		ZifIsWritable,
		ArginfoIsWritable,
		uint32(g.SizeOf("arginfo_is_writable")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_readable",
		ZifIsReadable,
		ArginfoIsReadable,
		uint32(g.SizeOf("arginfo_is_readable")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_executable",
		ZifIsExecutable,
		ArginfoIsExecutable,
		uint32(g.SizeOf("arginfo_is_executable")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_file",
		ZifIsFile,
		ArginfoIsFile,
		uint32(g.SizeOf("arginfo_is_file")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_dir",
		ZifIsDir,
		ArginfoIsDir,
		uint32(g.SizeOf("arginfo_is_dir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"is_link",
		ZifIsLink,
		ArginfoIsLink,
		uint32(g.SizeOf("arginfo_is_link")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stat",
		PhpIfStat,
		ArginfoStat,
		uint32(g.SizeOf("arginfo_stat")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"lstat",
		PhpIfLstat,
		ArginfoLstat,
		uint32(g.SizeOf("arginfo_lstat")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chown",
		ZifChown,
		ArginfoChown,
		uint32(g.SizeOf("arginfo_chown")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chgrp",
		ZifChgrp,
		ArginfoChgrp,
		uint32(g.SizeOf("arginfo_chgrp")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"lchown",
		ZifLchown,
		ArginfoLchown,
		uint32(g.SizeOf("arginfo_lchown")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"lchgrp",
		ZifLchgrp,
		ArginfoLchgrp,
		uint32(g.SizeOf("arginfo_lchgrp")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"chmod",
		ZifChmod,
		ArginfoChmod,
		uint32(g.SizeOf("arginfo_chmod")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"touch",
		ZifTouch,
		ArginfoTouch,
		uint32(g.SizeOf("arginfo_touch")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"clearstatcache",
		ZifClearstatcache,
		ArginfoClearstatcache,
		uint32(g.SizeOf("arginfo_clearstatcache")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"disk_total_space",
		ZifDiskTotalSpace,
		ArginfoDiskTotalSpace,
		uint32(g.SizeOf("arginfo_disk_total_space")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"disk_free_space",
		ZifDiskFreeSpace,
		ArginfoDiskFreeSpace,
		uint32(g.SizeOf("arginfo_disk_free_space")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"diskfreespace",
		ZifDiskFreeSpace,
		ArginfoDiskFreeSpace,
		uint32(g.SizeOf("arginfo_disk_free_space")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"realpath_cache_size",
		ZifRealpathCacheSize,
		ArginfoRealpathCacheSize,
		uint32(g.SizeOf("arginfo_realpath_cache_size")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"realpath_cache_get",
		ZifRealpathCacheGet,
		ArginfoRealpathCacheGet,
		uint32(g.SizeOf("arginfo_realpath_cache_get")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"mail",
		ZifMail,
		ArginfoMail,
		uint32(g.SizeOf("arginfo_mail")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ezmlm_hash",
		ZifEzmlmHash,
		ArginfoEzmlmHash,
		uint32(g.SizeOf("arginfo_ezmlm_hash")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 11,
	},
	{
		"openlog",
		ZifOpenlog,
		ArginfoOpenlog,
		uint32(g.SizeOf("arginfo_openlog")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"syslog",
		ZifSyslog,
		ArginfoSyslog,
		uint32(g.SizeOf("arginfo_syslog")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"closelog",
		ZifCloselog,
		ArginfoCloselog,
		uint32(g.SizeOf("arginfo_closelog")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"lcg_value",
		ZifLcgValue,
		ArginfoLcgValue,
		uint32(g.SizeOf("arginfo_lcg_value")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"metaphone",
		ZifMetaphone,
		ArginfoMetaphone,
		uint32(g.SizeOf("arginfo_metaphone")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_start",
		core.ZifObStart,
		ArginfoObStart,
		uint32(g.SizeOf("arginfo_ob_start")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_flush",
		core.ZifObFlush,
		ArginfoObFlush,
		uint32(g.SizeOf("arginfo_ob_flush")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_clean",
		core.ZifObClean,
		ArginfoObClean,
		uint32(g.SizeOf("arginfo_ob_clean")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_end_flush",
		core.ZifObEndFlush,
		ArginfoObEndFlush,
		uint32(g.SizeOf("arginfo_ob_end_flush")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_end_clean",
		core.ZifObEndClean,
		ArginfoObEndClean,
		uint32(g.SizeOf("arginfo_ob_end_clean")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_get_flush",
		core.ZifObGetFlush,
		ArginfoObGetFlush,
		uint32(g.SizeOf("arginfo_ob_get_flush")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_get_clean",
		core.ZifObGetClean,
		ArginfoObGetClean,
		uint32(g.SizeOf("arginfo_ob_get_clean")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_get_length",
		core.ZifObGetLength,
		ArginfoObGetLength,
		uint32(g.SizeOf("arginfo_ob_get_length")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_get_level",
		core.ZifObGetLevel,
		ArginfoObGetLevel,
		uint32(g.SizeOf("arginfo_ob_get_level")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_get_status",
		core.ZifObGetStatus,
		ArginfoObGetStatus,
		uint32(g.SizeOf("arginfo_ob_get_status")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_get_contents",
		core.ZifObGetContents,
		ArginfoObGetContents,
		uint32(g.SizeOf("arginfo_ob_get_contents")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_implicit_flush",
		core.ZifObImplicitFlush,
		ArginfoObImplicitFlush,
		uint32(g.SizeOf("arginfo_ob_implicit_flush")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ob_list_handlers",
		core.ZifObListHandlers,
		ArginfoObListHandlers,
		uint32(g.SizeOf("arginfo_ob_list_handlers")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ksort",
		ZifKsort,
		ArginfoKsort,
		uint32(g.SizeOf("arginfo_ksort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"krsort",
		ZifKrsort,
		ArginfoKrsort,
		uint32(g.SizeOf("arginfo_krsort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"natsort",
		ZifNatsort,
		ArginfoNatsort,
		uint32(g.SizeOf("arginfo_natsort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"natcasesort",
		ZifNatcasesort,
		ArginfoNatcasesort,
		uint32(g.SizeOf("arginfo_natcasesort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"asort",
		ZifAsort,
		ArginfoAsort,
		uint32(g.SizeOf("arginfo_asort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"arsort",
		ZifArsort,
		ArginfoArsort,
		uint32(g.SizeOf("arginfo_arsort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sort",
		ZifSort,
		ArginfoSort,
		uint32(g.SizeOf("arginfo_sort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rsort",
		ZifRsort,
		ArginfoRsort,
		uint32(g.SizeOf("arginfo_rsort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"usort",
		ZifUsort,
		ArginfoUsort,
		uint32(g.SizeOf("arginfo_usort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"uasort",
		ZifUasort,
		ArginfoUasort,
		uint32(g.SizeOf("arginfo_uasort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"uksort",
		ZifUksort,
		ArginfoUksort,
		uint32(g.SizeOf("arginfo_uksort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"shuffle",
		ZifShuffle,
		ArginfoShuffle,
		uint32(g.SizeOf("arginfo_shuffle")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_walk",
		ZifArrayWalk,
		ArginfoArrayWalk,
		uint32(g.SizeOf("arginfo_array_walk")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_walk_recursive",
		ZifArrayWalkRecursive,
		ArginfoArrayWalkRecursive,
		uint32(g.SizeOf("arginfo_array_walk_recursive")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"count",
		ZifCount,
		ArginfoCount,
		uint32(g.SizeOf("arginfo_count")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"end",
		ZifEnd,
		ArginfoEnd,
		uint32(g.SizeOf("arginfo_end")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"prev",
		ZifPrev,
		ArginfoPrev,
		uint32(g.SizeOf("arginfo_prev")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"next",
		ZifNext,
		ArginfoNext,
		uint32(g.SizeOf("arginfo_next")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"reset",
		ZifReset,
		ArginfoReset,
		uint32(g.SizeOf("arginfo_reset")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"current",
		ZifCurrent,
		ArginfoCurrent,
		uint32(g.SizeOf("arginfo_current")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"key",
		ZifKey,
		ArginfoKey,
		uint32(g.SizeOf("arginfo_key")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"min",
		ZifMin,
		ArginfoMin,
		uint32(g.SizeOf("arginfo_min")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"max",
		ZifMax,
		ArginfoMax,
		uint32(g.SizeOf("arginfo_max")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"in_array",
		ZifInArray,
		ArginfoInArray,
		uint32(g.SizeOf("arginfo_in_array")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_search",
		ZifArraySearch,
		ArginfoArraySearch,
		uint32(g.SizeOf("arginfo_array_search")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"extract",
		ZifExtract,
		ArginfoExtract,
		uint32(g.SizeOf("arginfo_extract")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"compact",
		ZifCompact,
		ArginfoCompact,
		uint32(g.SizeOf("arginfo_compact")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_fill",
		ZifArrayFill,
		ArginfoArrayFill,
		uint32(g.SizeOf("arginfo_array_fill")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_fill_keys",
		ZifArrayFillKeys,
		ArginfoArrayFillKeys,
		uint32(g.SizeOf("arginfo_array_fill_keys")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"range",
		ZifRange,
		ArginfoRange,
		uint32(g.SizeOf("arginfo_range")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_multisort",
		ZifArrayMultisort,
		ArginfoArrayMultisort,
		uint32(g.SizeOf("arginfo_array_multisort")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_push",
		ZifArrayPush,
		ArginfoArrayPush,
		uint32(g.SizeOf("arginfo_array_push")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_pop",
		ZifArrayPop,
		ArginfoArrayPop,
		uint32(g.SizeOf("arginfo_array_pop")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_shift",
		ZifArrayShift,
		ArginfoArrayShift,
		uint32(g.SizeOf("arginfo_array_shift")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_unshift",
		ZifArrayUnshift,
		ArginfoArrayUnshift,
		uint32(g.SizeOf("arginfo_array_unshift")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_splice",
		ZifArraySplice,
		ArginfoArraySplice,
		uint32(g.SizeOf("arginfo_array_splice")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_slice",
		ZifArraySlice,
		ArginfoArraySlice,
		uint32(g.SizeOf("arginfo_array_slice")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_merge",
		ZifArrayMerge,
		ArginfoArrayMerge,
		uint32(g.SizeOf("arginfo_array_merge")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_merge_recursive",
		ZifArrayMergeRecursive,
		ArginfoArrayMergeRecursive,
		uint32(g.SizeOf("arginfo_array_merge_recursive")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_replace",
		ZifArrayReplace,
		ArginfoArrayReplace,
		uint32(g.SizeOf("arginfo_array_replace")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_replace_recursive",
		ZifArrayReplaceRecursive,
		ArginfoArrayReplaceRecursive,
		uint32(g.SizeOf("arginfo_array_replace_recursive")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_keys",
		ZifArrayKeys,
		ArginfoArrayKeys,
		uint32(g.SizeOf("arginfo_array_keys")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_key_first",
		ZifArrayKeyFirst,
		ArginfoArrayKeyFirst,
		uint32(g.SizeOf("arginfo_array_key_first")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_key_last",
		ZifArrayKeyLast,
		ArginfoArrayKeyLast,
		uint32(g.SizeOf("arginfo_array_key_last")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_values",
		ZifArrayValues,
		ArginfoArrayValues,
		uint32(g.SizeOf("arginfo_array_values")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_count_values",
		ZifArrayCountValues,
		ArginfoArrayCountValues,
		uint32(g.SizeOf("arginfo_array_count_values")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_column",
		ZifArrayColumn,
		ArginfoArrayColumn,
		uint32(g.SizeOf("arginfo_array_column")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_reverse",
		ZifArrayReverse,
		ArginfoArrayReverse,
		uint32(g.SizeOf("arginfo_array_reverse")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_reduce",
		ZifArrayReduce,
		ArginfoArrayReduce,
		uint32(g.SizeOf("arginfo_array_reduce")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_pad",
		ZifArrayPad,
		ArginfoArrayPad,
		uint32(g.SizeOf("arginfo_array_pad")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_flip",
		ZifArrayFlip,
		ArginfoArrayFlip,
		uint32(g.SizeOf("arginfo_array_flip")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_change_key_case",
		ZifArrayChangeKeyCase,
		ArginfoArrayChangeKeyCase,
		uint32(g.SizeOf("arginfo_array_change_key_case")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_rand",
		ZifArrayRand,
		ArginfoArrayRand,
		uint32(g.SizeOf("arginfo_array_rand")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_unique",
		ZifArrayUnique,
		ArginfoArrayUnique,
		uint32(g.SizeOf("arginfo_array_unique")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_intersect",
		ZifArrayIntersect,
		ArginfoArrayIntersect,
		uint32(g.SizeOf("arginfo_array_intersect")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_intersect_key",
		ZifArrayIntersectKey,
		ArginfoArrayIntersectKey,
		uint32(g.SizeOf("arginfo_array_intersect_key")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_intersect_ukey",
		ZifArrayIntersectUkey,
		ArginfoArrayIntersectUkey,
		uint32(g.SizeOf("arginfo_array_intersect_ukey")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_uintersect",
		ZifArrayUintersect,
		ArginfoArrayUintersect,
		uint32(g.SizeOf("arginfo_array_uintersect")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_intersect_assoc",
		ZifArrayIntersectAssoc,
		ArginfoArrayIntersectAssoc,
		uint32(g.SizeOf("arginfo_array_intersect_assoc")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_uintersect_assoc",
		ZifArrayUintersectAssoc,
		ArginfoArrayUintersectAssoc,
		uint32(g.SizeOf("arginfo_array_uintersect_assoc")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_intersect_uassoc",
		ZifArrayIntersectUassoc,
		ArginfoArrayIntersectUassoc,
		uint32(g.SizeOf("arginfo_array_intersect_uassoc")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_uintersect_uassoc",
		ZifArrayUintersectUassoc,
		ArginfoArrayUintersectUassoc,
		uint32(g.SizeOf("arginfo_array_uintersect_uassoc")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_diff",
		ZifArrayDiff,
		ArginfoArrayDiff,
		uint32(g.SizeOf("arginfo_array_diff")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_diff_key",
		ZifArrayDiffKey,
		ArginfoArrayDiffKey,
		uint32(g.SizeOf("arginfo_array_diff_key")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_diff_ukey",
		ZifArrayDiffUkey,
		ArginfoArrayDiffUkey,
		uint32(g.SizeOf("arginfo_array_diff_ukey")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_udiff",
		ZifArrayUdiff,
		ArginfoArrayUdiff,
		uint32(g.SizeOf("arginfo_array_udiff")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_diff_assoc",
		ZifArrayDiffAssoc,
		ArginfoArrayDiffAssoc,
		uint32(g.SizeOf("arginfo_array_diff_assoc")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_udiff_assoc",
		ZifArrayUdiffAssoc,
		ArginfoArrayUdiffAssoc,
		uint32(g.SizeOf("arginfo_array_udiff_assoc")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_diff_uassoc",
		ZifArrayDiffUassoc,
		ArginfoArrayDiffUassoc,
		uint32(g.SizeOf("arginfo_array_diff_uassoc")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_udiff_uassoc",
		ZifArrayUdiffUassoc,
		ArginfoArrayUdiffUassoc,
		uint32(g.SizeOf("arginfo_array_udiff_uassoc")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_sum",
		ZifArraySum,
		ArginfoArraySum,
		uint32(g.SizeOf("arginfo_array_sum")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_product",
		ZifArrayProduct,
		ArginfoArrayProduct,
		uint32(g.SizeOf("arginfo_array_product")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_filter",
		ZifArrayFilter,
		ArginfoArrayFilter,
		uint32(g.SizeOf("arginfo_array_filter")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_map",
		ZifArrayMap,
		ArginfoArrayMap,
		uint32(g.SizeOf("arginfo_array_map")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_chunk",
		ZifArrayChunk,
		ArginfoArrayChunk,
		uint32(g.SizeOf("arginfo_array_chunk")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_combine",
		ZifArrayCombine,
		ArginfoArrayCombine,
		uint32(g.SizeOf("arginfo_array_combine")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"array_key_exists",
		ZifArrayKeyExists,
		ArginfoArrayKeyExists,
		uint32(g.SizeOf("arginfo_array_key_exists")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"pos",
		ZifCurrent,
		ArginfoCurrent,
		uint32(g.SizeOf("arginfo_current")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sizeof",
		ZifCount,
		ArginfoCount,
		uint32(g.SizeOf("arginfo_count")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"key_exists",
		ZifArrayKeyExists,
		ArginfoArrayKeyExists,
		uint32(g.SizeOf("arginfo_array_key_exists")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"assert",
		ZifAssert,
		ArginfoAssert,
		uint32(g.SizeOf("arginfo_assert")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"assert_options",
		ZifAssertOptions,
		ArginfoAssertOptions,
		uint32(g.SizeOf("arginfo_assert_options")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"version_compare",
		ZifVersionCompare,
		ArginfoVersionCompare,
		uint32(g.SizeOf("arginfo_version_compare")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"ftok",
		ZifFtok,
		ArginfoFtok,
		uint32(g.SizeOf("arginfo_ftok")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"str_rot13",
		ZifStrRot13,
		ArginfoStrRot13,
		uint32(g.SizeOf("arginfo_str_rot13")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_get_filters",
		ZifStreamGetFilters,
		ArginfoStreamGetFilters,
		uint32(g.SizeOf("arginfo_stream_get_filters")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_filter_register",
		ZifStreamFilterRegister,
		ArginfoStreamFilterRegister,
		uint32(g.SizeOf("arginfo_stream_filter_register")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_bucket_make_writeable",
		ZifStreamBucketMakeWriteable,
		ArginfoStreamBucketMakeWriteable,
		uint32(g.SizeOf("arginfo_stream_bucket_make_writeable")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_bucket_prepend",
		ZifStreamBucketPrepend,
		ArginfoStreamBucketPrepend,
		uint32(g.SizeOf("arginfo_stream_bucket_prepend")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_bucket_append",
		ZifStreamBucketAppend,
		ArginfoStreamBucketAppend,
		uint32(g.SizeOf("arginfo_stream_bucket_append")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"stream_bucket_new",
		ZifStreamBucketNew,
		ArginfoStreamBucketNew,
		uint32(g.SizeOf("arginfo_stream_bucket_new")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"output_add_rewrite_var",
		core.ZifOutputAddRewriteVar,
		ArginfoOutputAddRewriteVar,
		uint32(g.SizeOf("arginfo_output_add_rewrite_var")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"output_reset_rewrite_vars",
		core.ZifOutputResetRewriteVars,
		ArginfoOutputResetRewriteVars,
		uint32(g.SizeOf("arginfo_output_reset_rewrite_vars")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"sys_get_temp_dir",
		ZifSysGetTempDir,
		ArginfoSysGetTempDir,
		uint32(g.SizeOf("arginfo_sys_get_temp_dir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

var StandardDeps []zend.ZendModuleDep = []zend.ZendModuleDep{{"session", nil, nil, 3}, {nil, nil, nil, 0}}

/* }}} */

var BasicFunctionsModule zend.ZendModuleEntry = zend.ZendModuleEntry{g.SizeOf("zend_module_entry"), 20190902, 0, 0, nil, StandardDeps, "standard", BasicFunctions, ZmStartupBasic, ZmShutdownBasic, ZmActivateBasic, ZmDeactivateBasic, ZmInfoBasic, "7.4.33", 0, nil, nil, nil, nil, 0, 0, nil, 0, "API" + "20190902" + ",NTS"}

/* }}} */

func PhpPutenvDestructor(zv *zend.Zval) {
	var pe *PutenvEntry = zv.value.ptr
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
	zend._efree(pe.GetPutenvString())
	zend._efree(pe.GetKey())
	zend._efree(pe)
}

/* }}} */

func BasicGlobalsCtor(basic_globals_p *PhpBasicGlobals) {
	BasicGlobals.SetMtRandIsSeeded(0)
	BasicGlobals.SetMtRandMode(0)
	BasicGlobals.SetUmask(-1)
	BasicGlobals.SetNext(nil)
	BasicGlobals.SetLeft(-1)
	BasicGlobals.SetUserTickFunctions(nil)
	BasicGlobals.SetUserFilterMap(nil)
	BasicGlobals.SetSerializeLock(0)
	memset(&(BasicGlobals.GetSerialize()), 0, g.SizeOf("BG ( serialize )"))
	memset(&(BasicGlobals.GetUnserialize()), 0, g.SizeOf("BG ( unserialize )"))
	memset(&(BasicGlobals.GetUrlAdaptSessionEx()), 0, g.SizeOf("BG ( url_adapt_session_ex )"))
	memset(&(BasicGlobals.GetUrlAdaptOutputEx()), 0, g.SizeOf("BG ( url_adapt_output_ex )"))
	BasicGlobals.GetUrlAdaptSessionEx().SetType(1)
	BasicGlobals.GetUrlAdaptOutputEx().SetType(0)
	zend._zendHashInit(&(BasicGlobals.GetUrlAdaptSessionHostsHt()), 0, nil, 1)
	zend._zendHashInit(&(BasicGlobals.GetUrlAdaptOutputHostsHt()), 0, nil, 1)
	BasicGlobals.SetIncompleteClass(IncompleteClassEntry)
	BasicGlobals.SetPageUid(-1)
	BasicGlobals.SetPageGid(-1)
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

func PhpGetNan() float64 { return zend._zendGetNan() }

/* }}} */

func PhpGetInf() float64 { return zend._zendGetInf() }

/* }}} */

// #define BASIC_MINIT_SUBMODULE(module) if ( PHP_MINIT ( module ) ( INIT_FUNC_ARGS_PASSTHRU ) != SUCCESS ) { return FAILURE ; }

// #define BASIC_RINIT_SUBMODULE(module) PHP_RINIT ( module ) ( INIT_FUNC_ARGS_PASSTHRU ) ;

// #define BASIC_MINFO_SUBMODULE(module) PHP_MINFO ( module ) ( ZEND_MODULE_INFO_FUNC_ARGS_PASSTHRU ) ;

// #define BASIC_RSHUTDOWN_SUBMODULE(module) PHP_RSHUTDOWN ( module ) ( SHUTDOWN_FUNC_ARGS_PASSTHRU ) ;

// #define BASIC_MSHUTDOWN_SUBMODULE(module) PHP_MSHUTDOWN ( module ) ( SHUTDOWN_FUNC_ARGS_PASSTHRU ) ;

func ZmStartupBasic(type_ int, module_number int) int {
	BasicGlobalsCtor(&BasicGlobals)
	IncompleteClassEntry = PhpCreateIncompleteClass()
	BasicGlobals.SetIncompleteClass(IncompleteClassEntry)
	zend.ZendRegisterLongConstant("CONNECTION_ABORTED", g.SizeOf("\"CONNECTION_ABORTED\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("CONNECTION_NORMAL", g.SizeOf("\"CONNECTION_NORMAL\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("CONNECTION_TIMEOUT", g.SizeOf("\"CONNECTION_TIMEOUT\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("INI_USER", g.SizeOf("\"INI_USER\"")-1, 1<<0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("INI_PERDIR", g.SizeOf("\"INI_PERDIR\"")-1, 1<<1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("INI_SYSTEM", g.SizeOf("\"INI_SYSTEM\"")-1, 1<<2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("INI_ALL", g.SizeOf("\"INI_ALL\"")-1, 1<<0|1<<1|1<<2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("INI_SCANNER_NORMAL", g.SizeOf("\"INI_SCANNER_NORMAL\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("INI_SCANNER_RAW", g.SizeOf("\"INI_SCANNER_RAW\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("INI_SCANNER_TYPED", g.SizeOf("\"INI_SCANNER_TYPED\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_URL_SCHEME", g.SizeOf("\"PHP_URL_SCHEME\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_URL_HOST", g.SizeOf("\"PHP_URL_HOST\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_URL_PORT", g.SizeOf("\"PHP_URL_PORT\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_URL_USER", g.SizeOf("\"PHP_URL_USER\"")-1, 3, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_URL_PASS", g.SizeOf("\"PHP_URL_PASS\"")-1, 4, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_URL_PATH", g.SizeOf("\"PHP_URL_PATH\"")-1, 5, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_URL_QUERY", g.SizeOf("\"PHP_URL_QUERY\"")-1, 6, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_URL_FRAGMENT", g.SizeOf("\"PHP_URL_FRAGMENT\"")-1, 7, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_QUERY_RFC1738", g.SizeOf("\"PHP_QUERY_RFC1738\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_QUERY_RFC3986", g.SizeOf("\"PHP_QUERY_RFC3986\"")-1, 2, 1<<0|1<<1, module_number)

	// #define REGISTER_MATH_CONSTANT(x) REGISTER_DOUBLE_CONSTANT ( # x , x , CONST_CS | CONST_PERSISTENT )

	zend.ZendRegisterDoubleConstant("M_E", g.SizeOf("\"M_E\"")-1, 2.7182817, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_LOG2E", g.SizeOf("\"M_LOG2E\"")-1, 1.442695, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_LOG10E", g.SizeOf("\"M_LOG10E\"")-1, 0.4342945, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_LN2", g.SizeOf("\"M_LN2\"")-1, 0.6931472, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_LN10", g.SizeOf("\"M_LN10\"")-1, 2.3025851, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_PI", g.SizeOf("\"M_PI\"")-1, 3.1415927, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_PI_2", g.SizeOf("\"M_PI_2\"")-1, 1.5707964, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_PI_4", g.SizeOf("\"M_PI_4\"")-1, 0.7853982, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_1_PI", g.SizeOf("\"M_1_PI\"")-1, 0.31830987, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_2_PI", g.SizeOf("\"M_2_PI\"")-1, 0.63661975, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_SQRTPI", g.SizeOf("\"M_SQRTPI\"")-1, 1.7724539, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_2_SQRTPI", g.SizeOf("\"M_2_SQRTPI\"")-1, 1.1283792, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_LNPI", g.SizeOf("\"M_LNPI\"")-1, 1.1447299, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_EULER", g.SizeOf("\"M_EULER\"")-1, 0.5772157, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_SQRT2", g.SizeOf("\"M_SQRT2\"")-1, 1.4142135, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_SQRT1_2", g.SizeOf("\"M_SQRT1_2\"")-1, 0.70710677, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("M_SQRT3", g.SizeOf("\"M_SQRT3\"")-1, 1.7320508, 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("INF", g.SizeOf("\"INF\"")-1, zend._zendGetInf(), 1<<0|1<<1, module_number)
	zend.ZendRegisterDoubleConstant("NAN", g.SizeOf("\"NAN\"")-1, zend._zendGetNan(), 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_ROUND_HALF_UP", g.SizeOf("\"PHP_ROUND_HALF_UP\"")-1, 0x1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_ROUND_HALF_DOWN", g.SizeOf("\"PHP_ROUND_HALF_DOWN\"")-1, 0x2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_ROUND_HALF_EVEN", g.SizeOf("\"PHP_ROUND_HALF_EVEN\"")-1, 0x3, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PHP_ROUND_HALF_ODD", g.SizeOf("\"PHP_ROUND_HALF_ODD\"")-1, 0x4, 1<<0|1<<1, module_number)
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
	memset(BasicGlobals.GetStrtokTable(), 0, 256)
	BasicGlobals.SetSerializeLock(0)
	memset(&(BasicGlobals.GetSerialize()), 0, g.SizeOf("BG ( serialize )"))
	memset(&(BasicGlobals.GetUnserialize()), 0, g.SizeOf("BG ( unserialize )"))
	BasicGlobals.SetStrtokString(nil)
	&(BasicGlobals.GetStrtokZval()).u1.type_info = 0
	BasicGlobals.SetStrtokLast(nil)
	BasicGlobals.SetLocaleString(nil)
	BasicGlobals.SetLocaleChanged(0)
	BasicGlobals.SetArrayWalkFci(zend.EmptyFcallInfo)
	BasicGlobals.SetArrayWalkFciCache(zend.EmptyFcallInfoCache)
	BasicGlobals.SetUserCompareFci(zend.EmptyFcallInfo)
	BasicGlobals.SetUserCompareFciCache(zend.EmptyFcallInfoCache)
	BasicGlobals.SetPageUid(-1)
	BasicGlobals.SetPageGid(-1)
	BasicGlobals.SetPageInode(-1)
	BasicGlobals.SetPageMtime(-1)
	zend._zendHashInit(&(BasicGlobals.GetPutenvHt()), 1, PhpPutenvDestructor, 0)
	BasicGlobals.SetUserShutdownFunctionNames(nil)
	ZmActivateFilestat(type_, module_number)
	ZmActivateSyslog(type_, module_number)
	ZmActivateDir(type_, module_number)
	ZmActivateUrlScannerEx(type_, module_number)

	/* Setup default context */

	FileGlobals.SetDefaultContext(nil)

	/* Default to global wrappers only */

	FileGlobals.SetStreamWrappers(nil)

	/* Default to global filters only */

	FileGlobals.SetStreamFilters(nil)
	return zend.SUCCESS
}

/* }}} */

func ZmDeactivateBasic(type_ int, module_number int) int {
	zend.ZvalPtrDtor(&(BasicGlobals.GetStrtokZval()))
	&(BasicGlobals.GetStrtokZval()).u1.type_info = 0
	BasicGlobals.SetStrtokString(nil)
	tsrm_env_lock()
	zend.ZendHashDestroy(&(BasicGlobals.GetPutenvHt()))
	tsrm_env_unlock()
	BasicGlobals.SetMtRandIsSeeded(0)
	if BasicGlobals.GetUmask() != -1 {
		umask(BasicGlobals.GetUmask())
	}

	/* Check if locale was changed and change it back
	 * to the value in startup environment */

	if BasicGlobals.GetLocaleChanged() != 0 {
		setlocale(LC_ALL, "C")
		setlocale(LC_CTYPE, "")
		if BasicGlobals.GetLocaleString() != nil {
			zend.ZendStringReleaseEx(BasicGlobals.GetLocaleString(), 0)
			BasicGlobals.SetLocaleString(nil)
		}
	}

	/* FG(stream_wrappers) and FG(stream_filters) are destroyed
	 * during php_request_shutdown() */

	ZmDeactivateFilestat(type_, module_number)
	ZmDeactivateAssert(type_, module_number)
	ZmDeactivateUrlScannerEx(type_, module_number)
	streams.ZmDeactivateStreams(type_, module_number)
	if BasicGlobals.GetUserTickFunctions() != nil {
		zend.ZendLlistDestroy(BasicGlobals.GetUserTickFunctions())
		zend._efree(BasicGlobals.GetUserTickFunctions())
		BasicGlobals.SetUserTickFunctions(nil)
	}
	ZmDeactivateUserFilters(type_, module_number)
	ZmDeactivateBrowscap(type_, module_number)
	BasicGlobals.SetPageUid(-1)
	BasicGlobals.SetPageGid(-1)
	return zend.SUCCESS
}

/* }}} */

func ZmInfoBasic(zend_module *zend.ZendModuleEntry) {
	PhpInfoPrintTableStart()
	ZmInfoDl(zend_module)
	ZmInfoMail(zend_module)
	PhpInfoPrintTableEnd()
	ZmInfoAssert(zend_module)
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

			if zend.ZendParseArgStr(_arg, &const_name, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	scope = zend.ZendGetExecutedScope()
	c = zend.ZendGetConstantEx(const_name, scope, 0x100)
	if c != nil {
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = c
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			if (zend.ZvalGcFlags(_gc.gc.u.type_info) & 1 << 7) == 0 {
				zend.ZendGcAddref(&_gc.gc)
			} else {
				zend.ZvalCopyCtorFunc(_z1)
			}
		}
		if return_value.u1.v.type_ == 11 {
			if zend.ZvalUpdateConstantEx(return_value, scope) != zend.SUCCESS {
				return
			}
		}
	} else {
		if zend.EG.exception == nil {
			core.PhpErrorDocref(nil, 1<<1, "Couldn't find constant %s", const_name.val)
		}
		return_value.u1.type_info = 1
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

			if zend.ZendParseArgString(_arg, &address, &address_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if address_len == 16 {
		af = AF_INET6
	} else if address_len != 4 {
		return_value.u1.type_info = 2
		return
	}
	if !(inet_ntop(af, address, buffer, g.SizeOf("buffer"))) {
		return_value.u1.type_info = 2
		return
	}
	var _s *byte = buffer
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
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

			if zend.ZendParseArgString(_arg, &address, &address_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	memset(buffer, 0, g.SizeOf("buffer"))
	if strchr(address, ':') {
		af = AF_INET6
	} else if !(strchr(address, '.')) {
		return_value.u1.type_info = 2
		return
	}
	ret = inet_pton(af, address, buffer)
	if ret <= 0 {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(buffer, g.Cond(af == AF_INET, 4, 16), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
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

			if zend.ZendParseArgString(_arg, &addr, &addr_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	if addr_len == 0 || inet_pton(AF_INET, addr, &ip) != 1 {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = ntohl(ip.s_addr)
	__z.u1.type_info = 4
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

			if zend.ZendParseArgLong(_arg, &sip, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

	/* autoboxes on 32bit platforms, but that's expected */

	ip = zend.ZendUlong(sip)
	myaddr.s_addr = htonl(ip)
	if inet_ntop(AF_INET, &myaddr, str, g.SizeOf("str")) {
		var _s *byte = str
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		return_value.u1.type_info = 2
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
			_optional = 1
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

			if zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgBool(_arg, &local_only, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if str == nil {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		core.PhpImportEnvironmentVariables(return_value)
		return
	}
	if local_only == 0 {

		/* SAPI method returns an emalloc()'d string */

		ptr = core.SapiGetenv(str, str_len)
		if ptr != nil {

			// TODO: avoid realocation ???

			var _s *byte = ptr
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			zend._efree(ptr)
			return
		}
	}
	tsrm_env_lock()

	/* system method returns a const */

	ptr = getenv(str)
	if ptr != nil {
		var _s *byte = ptr
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
	}
	tsrm_env_unlock()
	if ptr != nil {
		return
	}
	return_value.u1.type_info = 2
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

			if zend.ZendParseArgString(_arg, &setting, &setting_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	if setting_len == 0 || setting[0] == '=' {
		core.PhpErrorDocref(nil, 1<<1, "Invalid parameter syntax")
		return_value.u1.type_info = 2
		return
	}
	pe.SetPutenvString(zend._estrndup(setting, setting_len))
	pe.SetKey(zend._estrndup(setting, setting_len))
	if g.Assign(&p, strchr(pe.GetKey(), '=')) {
		*p = '0'
	}
	pe.SetKeyLen(strlen(pe.GetKey()))
	tsrm_env_lock()
	zend.ZendHashStrDel(&(BasicGlobals.GetPutenvHt()), pe.GetKey(), pe.GetKeyLen())

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
		zend.ZendHashStrAddMem(&(BasicGlobals.GetPutenvHt()), pe.GetKey(), pe.GetKeyLen(), &pe, g.SizeOf("putenv_entry"))
		if !(strncmp(pe.GetKey(), "TZ", pe.GetKeyLen())) {
			tzset()
		}
		tsrm_env_unlock()
		return_value.u1.type_info = 3
		return
	} else {
		zend._efree(pe.GetPutenvString())
		zend._efree(pe.GetKey())
		return_value.u1.type_info = 2
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
				zend._efree(argv[i])
			}
		}
		zend._efree(argv)
	}
}

/* }}} */

func FreeLongopts(longopts *core.Opt) {
	var p *core.Opt
	if longopts != nil {
		for p = longopts; p != nil && p.opt_char != '-'; p++ {
			if p.opt_name != nil {
				zend._efree((*byte)(p.opt_name))
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
	paras = zend._safeEmalloc(g.SizeOf("opt_struct"), count, 0)
	memset(paras, 0, g.SizeOf("opt_struct")*count)
	*result = paras
	for (*opts) >= 48 && (*opts) <= 57 || (*opts) >= 65 && (*opts) <= 90 || (*opts) >= 97 && (*opts) <= 122 {
		paras.opt_char = *opts
		paras.need_param = (*(g.PreInc(&opts))) == ':'
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

			if zend.ZendParseArgString(_arg, &options, &options_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgArray(_arg, &p_longopts, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
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

			zend.ZendParseArgZvalDeref(_arg, &zoptind, 0)
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
			return_value.u1.type_info = 2
			return
		}
		break
	}

	/* Init zoptind to 1 */

	if zoptind != nil {
		for {
			r.Assert(zoptind.u1.v.type_ == 10)
			for {
				var _zv *zend.Zval = zoptind
				var ref *zend.ZendReference = _zv.value.ref
				if ref.sources.ptr != nil {
					zend.ZendTryAssignTypedRefLong(ref, 1)
					break
				}
				_zv = &ref.val
				zend.ZvalPtrDtor(_zv)
				var __z *zend.Zval = _zv
				__z.value.lval = 1
				__z.u1.type_info = 4
				break
			}
			break
		}
	}

	/* Get argv from the global symbol table. We calculate argc ourselves
	 * in order to be on the safe side, even though it is also available
	 * from the symbol table. */

	if (core.CoreGlobals.http_globals[3].u1.v.type_ == 7 || zend.ZendIsAutoGlobalStr("_SERVER", g.SizeOf("\"_SERVER\"")-1) != 0) && (g.Assign(&args, zend.ZendHashFindExInd(&core.CoreGlobals.http_globals[3].value.arr, zend.ZendKnownStrings[zend.ZEND_STR_ARGV], 1)) != nil || g.Assign(&args, zend.ZendHashFindExInd(&zend.EG.symbol_table, zend.ZendKnownStrings[zend.ZEND_STR_ARGV], 1)) != nil) {
		var pos int = 0
		var entry *zend.Zval
		if args.u1.v.type_ != 7 {
			return_value.u1.type_info = 2
			return
		}
		argc = args.value.arr.nNumOfElements

		/* Attempt to allocate enough memory to hold all of the arguments
		 * and a trailing NULL */

		argv = (**byte)(zend._safeEmalloc(g.SizeOf("char *"), argc+1, 0))

		/* Iterate over the hash to construct the argv array. */

		for {
			var __ht *zend.HashTable = args.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				entry = _z
				var tmp_arg_str *zend.ZendString
				var arg_str *zend.ZendString = zend.ZvalGetTmpString(entry, &tmp_arg_str)
				argv[g.PostInc(&pos)] = zend._estrdup(arg_str.val)
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

		return_value.u1.type_info = 2
		return
	}
	len_ = ParseOpts(options, &opts)
	if p_longopts != nil {
		var count int
		var entry *zend.Zval
		count = p_longopts.value.arr.nNumOfElements

		/* the first <len> slots are filled by the one short ops
		 * we now extend our array and jump to the new added structs */

		opts = (*core.Opt)(zend._erealloc(opts, g.SizeOf("opt_struct")*(len_+count+1)))
		orig_opts = opts
		opts += len_
		memset(opts, 0, count*g.SizeOf("opt_struct"))

		/* Iterate over the hash to construct the argv array. */

		for {
			var __ht *zend.HashTable = p_longopts.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				entry = _z
				var tmp_arg_str *zend.ZendString
				var arg_str *zend.ZendString = zend.ZvalGetTmpString(entry, &tmp_arg_str)
				opts.need_param = 0
				opts.opt_name = zend._estrdup(arg_str.val)
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
		opts = (*core.Opt)(zend._erealloc(opts, g.SizeOf("opt_struct")*(len_+1)))
		orig_opts = opts
		opts += len_
	}

	/* php_getopt want to identify the last param */

	opts.opt_char = '-'
	opts.need_param = 0
	opts.opt_name = nil

	/* Initialize the return value as an array. */

	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

	/* after our pointer arithmetic jump back to the first element */

	opts = orig_opts
	for g.Assign(&o, core.PhpGetopt(argc, argv, opts, &php_optarg, &php_optind, 0, 1)) != -1 {

		/* Skip unknown arguments. */

		if o == -2 {
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

			var _s *byte = php_optarg
			var __z *zend.Zval = &val
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8

			/* keep the arg as binary, since the encoding is not known */

		} else {
			&val.u1.type_info = 2
		}

		/* Add this option / argument pair to the result hash. */

		optname_len = strlen(optname)
		if !(optname_len > 1 && optname[0] == '0') && zend.IsNumericString(optname, optname_len, nil, nil, 0) == 4 {

			/* numeric string */

			var optname_int int = atoi(optname)
			if g.Assign(&args, zend.ZendHashIndexFind(return_value.value.arr, optname_int)) != nil {
				if args.u1.v.type_ != 7 {
					if args.u1.v.type_ != 7 {
						zend.ConvertToArray(args)
					}
				}
				zend.ZendHashNextIndexInsert(args.value.arr, &val)
			} else {
				zend.ZendHashIndexUpdate(return_value.value.arr, optname_int, &val)
			}
		} else {

			/* other strings */

			if g.Assign(&args, zend.ZendHashStrFind(return_value.value.arr, optname, strlen(optname))) != nil {
				if args.u1.v.type_ != 7 {
					if args.u1.v.type_ != 7 {
						zend.ConvertToArray(args)
					}
				}
				zend.ZendHashNextIndexInsert(args.value.arr, &val)
			} else {
				zend.ZendHashStrAdd(return_value.value.arr, optname, strlen(optname), &val)
			}

			/* other strings */

		}
		php_optarg = nil
	}

	/* Set zoptind to php_optind */

	if zoptind != nil {
		for {
			r.Assert(zoptind.u1.v.type_ == 10)
			for {
				var _zv *zend.Zval = zoptind
				var ref *zend.ZendReference = _zv.value.ref
				if ref.sources.ptr != nil {
					zend.ZendTryAssignTypedRefLong(ref, php_optind)
					break
				}
				_zv = &ref.val
				zend.ZvalPtrDtor(_zv)
				var __z *zend.Zval = _zv
				__z.value.lval = php_optind
				__z.u1.type_info = 4
				break
			}
			break
		}
	}
	FreeLongopts(orig_opts)
	zend._efree(orig_opts)
	FreeArgv(argv, argc)
}

/* }}} */

func ZifFlush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
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

			if zend.ZendParseArgLong(_arg, &num, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if num < 0 {
		core.PhpErrorDocref(nil, 1<<1, "Number of seconds must be greater than or equal to 0")
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = sleep(uint(num))
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifUsleep(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var num zend.ZendLong
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

			if zend.ZendParseArgLong(_arg, &num, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if num < 0 {
		core.PhpErrorDocref(nil, 1<<1, "Number of microseconds must be greater than or equal to 0")
		return_value.u1.type_info = 2
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

			if zend.ZendParseArgLong(_arg, &tv_sec, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgLong(_arg, &tv_nsec, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if tv_sec < 0 {
		core.PhpErrorDocref(nil, 1<<1, "The seconds value must be greater than 0")
		return_value.u1.type_info = 2
		return
	}
	if tv_nsec < 0 {
		core.PhpErrorDocref(nil, 1<<1, "The nanoseconds value must be greater than 0")
		return_value.u1.type_info = 2
		return
	}
	php_req.tv_sec = int64(tv_sec)
	php_req.tv_nsec = long(tv_nsec)
	if !(nanosleep(&php_req, &php_rem)) {
		return_value.u1.type_info = 3
		return
	} else if errno == EINTR {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		zend.AddAssocLongEx(return_value, "seconds", g.SizeOf("\"seconds\"")-1, php_rem.tv_sec)
		zend.AddAssocLongEx(return_value, "nanoseconds", g.SizeOf("\"nanoseconds\"")-1, php_rem.tv_nsec)
		return
	} else if errno == EINVAL {
		core.PhpErrorDocref(nil, 1<<1, "nanoseconds was not in the range 0 to 999 999 999 or seconds was negative")
	}
	return_value.u1.type_info = 2
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

			if zend.ZendParseArgDouble(_arg, &target_secs, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_DOUBLE
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
	if gettimeofday((*__struct__timeval)(&tm), nil) != 0 {
		return_value.u1.type_info = 2
		return
	}
	target_ns = uint64(target_secs * ns_per_sec)
	current_ns = uint64(tm.tv_sec)*ns_per_sec + uint64(tm.tv_usec)*1000
	if target_ns < current_ns {
		core.PhpErrorDocref(nil, 1<<1, "Sleep until to time is less than current time")
		return_value.u1.type_info = 2
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
			return_value.u1.type_info = 2
			return
		}
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

/* {{{ proto string get_current_user(void)
   Get the name of the owner of the current PHP script */

func ZifGetCurrentUser(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var _s *byte = core.PhpGetCurrentUser()
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

/* {{{ add_config_entry
 */

func AddConfigEntry(h zend.ZendUlong, key *zend.ZendString, entry *zend.Zval, retval *zend.Zval) {
	if entry.u1.v.type_ == 6 {
		var str *zend.ZendString = entry.value.str
		if (zend.ZvalGcFlags(str.gc.u.type_info) & 1 << 6) == 0 {
			if (zend.ZvalGcFlags(str.gc.u.type_info) & 1 << 7) == 0 {
				zend.ZendStringAddref(str)
			} else {
				str = zend.ZendStringInit(str.val, str.len_, 0)
			}
		}
		if key != nil {
			zend.AddAssocStrEx(retval, key.val, key.len_, str)
		} else {
			zend.AddIndexStr(retval, h, str)
		}
	} else if entry.u1.v.type_ == 7 {
		var tmp zend.Zval
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &tmp
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		AddConfigEntries(entry.value.arr, &tmp)
		zend.ZendHashUpdate(retval.value.arr, key, &tmp)
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

			if _z.u1.v.type_ == 0 {
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

			if zend.ZendParseArgString(_arg, &varname, &varname_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	retval = core.CfgGetEntry(varname, uint32(varname_len))
	if retval != nil {
		if retval.u1.v.type_ == 7 {
			var __arr *zend.ZendArray = zend._zendNewArray(0)
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			AddConfigEntries(retval.value.arr, return_value)
			return
		} else {
			var _s *byte = retval.value.str.val
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return
		}
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifGetMagicQuotesRuntime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifGetMagicQuotesGpc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	return_value.u1.type_info = 2
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
	var argc int = execute_data.This.u2.num_args
	var erropt zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 4
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

			if zend.ZendParseArgString(_arg, &message, &message_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgLong(_arg, &erropt, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgPath(_arg, &opt, &opt_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgString(_arg, &headers, &headers_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	if argc > 1 {
		opt_err = int(erropt)
	}
	if _phpErrorLogEx(opt_err, message, message_len, opt, headers) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func _phpErrorLog(opt_err int, message *byte, opt *byte, headers *byte) int {
	return _phpErrorLogEx(opt_err, message, g.CondF1(opt_err == 3, func() __auto__ { return strlen(message) }, 0), opt, headers)
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
		core.PhpErrorDocref(nil, 1<<1, "TCP/IP option not available!")
		return zend.FAILURE
		break
	case 3:
		stream = streams._phpStreamOpenWrapperEx(opt, "a", 0|0x8, nil, nil)
		if stream == nil {
			return zend.FAILURE
		}
		nbytes = streams._phpStreamWrite(stream, message, message_len)
		streams._phpStreamFree(stream, 1|2)
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
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if core.CoreGlobals.last_error_message != nil {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		zend.AddAssocLongEx(return_value, "type", g.SizeOf("\"type\"")-1, core.CoreGlobals.last_error_type)
		zend.AddAssocStringEx(return_value, "message", g.SizeOf("\"message\"")-1, core.CoreGlobals.last_error_message)
		zend.AddAssocStringEx(return_value, "file", g.SizeOf("\"file\"")-1, g.CondF1(core.CoreGlobals.last_error_file != nil, func() *byte { return core.CoreGlobals.last_error_file }, "-"))
		zend.AddAssocLongEx(return_value, "line", g.SizeOf("\"line\"")-1, core.CoreGlobals.last_error_lineno)
	}
}

/* }}} */

func ZifErrorClearLast(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if core.CoreGlobals.last_error_message != nil {
		core.CoreGlobals.last_error_type = 0
		core.CoreGlobals.last_error_lineno = 0
		zend.Free(core.CoreGlobals.last_error_message)
		core.CoreGlobals.last_error_message = nil
		if core.CoreGlobals.last_error_file != nil {
			zend.Free(core.CoreGlobals.last_error_file)
			core.CoreGlobals.last_error_file = nil
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

			if zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = 4
					break
				} else {
					_error_code = 2
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
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
	fci.retval = &retval
	if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && retval.u1.v.type_ != 0 {
		if retval.u1.v.type_ == 10 {
			zend.ZendUnwrapReference(&retval)
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = &retval
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
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

			if zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = 4
					break
				} else {
					_error_code = 2
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
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

			if zend.ZendParseArgArray(_arg, &params, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
	zend.ZendFcallInfoArgs(&fci, params)
	fci.retval = &retval
	if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && retval.u1.v.type_ != 0 {
		if retval.u1.v.type_ == 10 {
			zend.ZendUnwrapReference(&retval)
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = &retval
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
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

			if zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = 4
					break
				} else {
					_error_code = 2
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
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
	if execute_data.prev_execute_data.func_.common.scope == nil {
		zend.ZendThrowError(nil, "Cannot call forward_static_call() when no class scope is active")
		return
	}
	fci.retval = &retval
	called_scope = zend.ZendGetCalledScope(execute_data)
	if called_scope != nil && fci_cache.calling_scope != nil && zend.InstanceofFunction(called_scope, fci_cache.calling_scope) != 0 {
		fci_cache.called_scope = called_scope
	}
	if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && retval.u1.v.type_ != 0 {
		if retval.u1.v.type_ == 10 {
			zend.ZendUnwrapReference(&retval)
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = &retval
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
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

			if zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = 4
					break
				} else {
					_error_code = 2
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
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

			if zend.ZendParseArgArray(_arg, &params, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
	zend.ZendFcallInfoArgs(&fci, params)
	fci.retval = &retval
	called_scope = zend.ZendGetCalledScope(execute_data)
	if called_scope != nil && fci_cache.calling_scope != nil && zend.InstanceofFunction(called_scope, fci_cache.calling_scope) != 0 {
		fci_cache.called_scope = called_scope
	}
	if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && retval.u1.v.type_ != 0 {
		if retval.u1.v.type_ == 10 {
			zend.ZendUnwrapReference(&retval)
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = &retval
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	}
	zend.ZendFcallInfoArgsClear(&fci, 1)
}

/* }}} */

func UserShutdownFunctionDtor(zv *zend.Zval) {
	var i int
	var shutdown_function_entry *PhpShutdownFunctionEntry = zv.value.ptr
	for i = 0; i < shutdown_function_entry.GetArgCount(); i++ {
		zend.ZvalPtrDtor(&shutdown_function_entry.arguments[i])
	}
	zend._efree(shutdown_function_entry.GetArguments())
	zend._efree(shutdown_function_entry)
}

/* }}} */

func UserTickFunctionDtor(tick_function_entry *UserTickFunctionEntry) {
	var i int
	for i = 0; i < tick_function_entry.GetArgCount(); i++ {
		zend.ZvalPtrDtor(&tick_function_entry.arguments[i])
	}
	zend._efree(tick_function_entry.GetArguments())
}

/* }}} */

func UserShutdownFunctionCall(zv *zend.Zval) int {
	var shutdown_function_entry *PhpShutdownFunctionEntry = zv.value.ptr
	var retval zend.Zval
	if zend.ZendIsCallable(&shutdown_function_entry.arguments[0], 0, nil) == 0 {
		var function_name *zend.ZendString = zend.ZendGetCallableName(&shutdown_function_entry.arguments[0])
		zend.ZendError(1<<1, "(Registered shutdown functions) Unable to call %s() - function does not exist", function_name.val)
		zend.ZendStringReleaseEx(function_name, 0)
		return 0
	}
	if zend._callUserFunctionEx(nil, &shutdown_function_entry.arguments[0], &retval, shutdown_function_entry.GetArgCount()-1, shutdown_function_entry.GetArguments()+1, 1) == zend.SUCCESS {
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
		if zend._callUserFunctionEx(nil, function, &retval, tick_fe.GetArgCount()-1, tick_fe.GetArguments()+1, 1) == zend.SUCCESS {
			zend.ZvalPtrDtor(&retval)
		} else {
			var obj *zend.Zval
			var method *zend.Zval
			if function.u1.v.type_ == 6 {
				core.PhpErrorDocref(nil, 1<<1, "Unable to call %s() - function does not exist", function.value.str.val)
			} else if function.u1.v.type_ == 7 && g.Assign(&obj, zend.ZendHashIndexFind(function.value.arr, 0)) != nil && g.Assign(&method, zend.ZendHashIndexFind(function.value.arr, 1)) != nil && obj.u1.v.type_ == 8 && method.u1.v.type_ == 6 {
				core.PhpErrorDocref(nil, 1<<1, "Unable to call %s::%s() - function does not exist", obj.value.obj.ce.name.val, method.value.str.val)
			} else {
				core.PhpErrorDocref(nil, 1<<1, "Unable to call tick function")
			}
		}
		tick_fe.SetCalling(0)
	}

	/* Prevent reentrant calls to the same user ticks function */
}

/* }}} */

func RunUserTickFunctions(tick_count int, arg any) {
	zend.ZendLlistApply(BasicGlobals.GetUserTickFunctions(), zend.LlistApplyFuncT(UserTickFunctionCall))
}

/* }}} */

func UserTickFunctionCompare(tick_fe1 *UserTickFunctionEntry, tick_fe2 *UserTickFunctionEntry) int {
	var func1 *zend.Zval = &tick_fe1.arguments[0]
	var func2 *zend.Zval = &tick_fe2.arguments[0]
	var ret int
	if func1.u1.v.type_ == 6 && func2.u1.v.type_ == 6 {
		ret = zend.ZendBinaryZvalStrcmp(func1, func2) == 0
	} else if func1.u1.v.type_ == 7 && func2.u1.v.type_ == 7 {
		ret = zend.ZendCompareArrays(func1, func2) == 0
	} else if func1.u1.v.type_ == 8 && func2.u1.v.type_ == 8 {
		ret = zend.ZendCompareObjects(func1, func2) == 0
	} else {
		ret = 0
	}
	if ret != 0 && tick_fe1.GetCalling() != 0 {
		core.PhpErrorDocref(nil, 1<<1, "Unable to delete tick function executed at the moment")
		return 0
	}
	return ret
}

/* }}} */

func PhpCallShutdownFunctions() {
	if BasicGlobals.GetUserShutdownFunctionNames() != nil {
		var __orig_bailout *sigjmp_buf = zend.EG.bailout
		var __bailout sigjmp_buf
		zend.EG.bailout = &__bailout
		if sigsetjmp(__bailout, 0) == 0 {
			zend.ZendHashApply(BasicGlobals.GetUserShutdownFunctionNames(), UserShutdownFunctionCall)
		}
		zend.EG.bailout = __orig_bailout
	}
}

/* }}} */

func PhpFreeShutdownFunctions() {
	if BasicGlobals.GetUserShutdownFunctionNames() != nil {
		var __orig_bailout *sigjmp_buf = zend.EG.bailout
		var __bailout sigjmp_buf
		zend.EG.bailout = &__bailout
		if sigsetjmp(__bailout, 0) == 0 {
			zend.ZendHashDestroy(BasicGlobals.GetUserShutdownFunctionNames())
			zend._efree(BasicGlobals.GetUserShutdownFunctionNames())
			BasicGlobals.SetUserShutdownFunctionNames(nil)
		} else {
			zend.EG.bailout = __orig_bailout

			/* maybe shutdown method call exit, we just ignore it */

			zend._efree(BasicGlobals.GetUserShutdownFunctionNames())
			BasicGlobals.SetUserShutdownFunctionNames(nil)
		}
		zend.EG.bailout = __orig_bailout
	}
}

/* }}} */

func ZifRegisterShutdownFunction(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var shutdown_function_entry PhpShutdownFunctionEntry
	var i int
	shutdown_function_entry.SetArgCount(execute_data.This.u2.num_args)
	if shutdown_function_entry.GetArgCount() < 1 {
		zend.ZendWrongParamCount()
		return
	}
	shutdown_function_entry.SetArguments((*zend.Zval)(zend._safeEmalloc(g.SizeOf("zval"), shutdown_function_entry.GetArgCount(), 0)))
	if zend._zendGetParametersArrayEx(shutdown_function_entry.GetArgCount(), shutdown_function_entry.GetArguments()) == zend.FAILURE {
		zend._efree(shutdown_function_entry.GetArguments())
		return_value.u1.type_info = 2
		return
	}

	/* Prevent entering of anything but valid callback (syntax check only!) */

	if zend.ZendIsCallable(&shutdown_function_entry.arguments[0], 0, nil) == 0 {
		var callback_name *zend.ZendString = zend.ZendGetCallableName(&shutdown_function_entry.arguments[0])
		core.PhpErrorDocref(nil, 1<<1, "Invalid shutdown callback '%s' passed", callback_name.val)
		zend._efree(shutdown_function_entry.GetArguments())
		zend.ZendStringReleaseEx(callback_name, 0)
		return_value.u1.type_info = 2
	} else {
		if BasicGlobals.GetUserShutdownFunctionNames() == nil {
			BasicGlobals.SetUserShutdownFunctionNames((*zend.HashTable)(zend._emalloc(g.SizeOf("HashTable"))))
			zend._zendHashInit(BasicGlobals.GetUserShutdownFunctionNames(), 0, UserShutdownFunctionDtor, 0)
		}
		for i = 0; i < shutdown_function_entry.GetArgCount(); i++ {
			if &shutdown_function_entry.GetArguments()[i].u1.v.type_flags != 0 {
				zend.ZvalAddrefP(&shutdown_function_entry.GetArguments()[i])
			}
		}
		zend.ZendHashNextIndexInsertMem(BasicGlobals.GetUserShutdownFunctionNames(), &shutdown_function_entry, g.SizeOf("php_shutdown_function_entry"))
	}

	/* Prevent entering of anything but valid callback (syntax check only!) */
}

/* }}} */

func RegisterUserShutdownFunction(function_name *byte, function_len int, shutdown_function_entry *PhpShutdownFunctionEntry) zend.ZendBool {
	if BasicGlobals.GetUserShutdownFunctionNames() == nil {
		BasicGlobals.SetUserShutdownFunctionNames((*zend.HashTable)(zend._emalloc(g.SizeOf("HashTable"))))
		zend._zendHashInit(BasicGlobals.GetUserShutdownFunctionNames(), 0, UserShutdownFunctionDtor, 0)
	}
	zend.ZendHashStrUpdateMem(BasicGlobals.GetUserShutdownFunctionNames(), function_name, function_len, shutdown_function_entry, g.SizeOf("php_shutdown_function_entry"))
	return 1
}

/* }}} */

func RemoveUserShutdownFunction(function_name *byte, function_len int) zend.ZendBool {
	if BasicGlobals.GetUserShutdownFunctionNames() != nil {
		return zend.ZendHashStrDel(BasicGlobals.GetUserShutdownFunctionNames(), function_name, function_len) != zend.FAILURE
	}
	return 0
}

/* }}} */

func AppendUserShutdownFunction(shutdown_function_entry PhpShutdownFunctionEntry) zend.ZendBool {
	if BasicGlobals.GetUserShutdownFunctionNames() == nil {
		BasicGlobals.SetUserShutdownFunctionNames((*zend.HashTable)(zend._emalloc(g.SizeOf("HashTable"))))
		zend._zendHashInit(BasicGlobals.GetUserShutdownFunctionNames(), 0, UserShutdownFunctionDtor, 0)
	}
	return zend.ZendHashNextIndexInsertMem(BasicGlobals.GetUserShutdownFunctionNames(), &shutdown_function_entry, g.SizeOf("php_shutdown_function_entry")) != nil
}

/* }}} */

func PhpGetHighlight(syntax_highlighter_ini *zend.ZendSyntaxHighlighterIni) {
	syntax_highlighter_ini.highlight_comment = zend.ZendIniStringEx("highlight.comment", g.SizeOf("\"highlight.comment\"")-1, 0, nil)
	syntax_highlighter_ini.highlight_default = zend.ZendIniStringEx("highlight.default", g.SizeOf("\"highlight.default\"")-1, 0, nil)
	syntax_highlighter_ini.highlight_html = zend.ZendIniStringEx("highlight.html", g.SizeOf("\"highlight.html\"")-1, 0, nil)
	syntax_highlighter_ini.highlight_keyword = zend.ZendIniStringEx("highlight.keyword", g.SizeOf("\"highlight.keyword\"")-1, 0, nil)
	syntax_highlighter_ini.highlight_string = zend.ZendIniStringEx("highlight.string", g.SizeOf("\"highlight.string\"")-1, 0, nil)
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

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgBool(_arg, &i, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if core.PhpCheckOpenBasedir(filename) != 0 {
		return_value.u1.type_info = 2
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
		return_value.u1.type_info = 2
		return
	}
	if i != 0 {
		core.PhpOutputGetContents(return_value)
		core.PhpOutputDiscard()
	} else {
		return_value.u1.type_info = 3
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

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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
			return_value.u1.type_info = 2
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
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
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
	var old_error_reporting int = zend.EG.error_reporting
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
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

			zend.ZendParseArgZvalDeref(_arg, &expr, 0)
			_optional = 1
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

			if zend.ZendParseArgBool(_arg, &i, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
			return_value.u1.type_info = 2
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
	zend.EG.error_reporting = 1 << 0
	PhpGetHighlight(&syntax_highlighter_ini)
	hicompiled_string_description = zend.ZendMakeCompiledStringDescription("highlighted code")
	if zend.HighlightString(expr, &syntax_highlighter_ini, hicompiled_string_description) == zend.FAILURE {
		zend._efree(hicompiled_string_description)
		zend.EG.error_reporting = old_error_reporting
		if i != 0 {
			core.PhpOutputEnd()
		}
		return_value.u1.type_info = 2
		return
	}
	zend._efree(hicompiled_string_description)
	zend.EG.error_reporting = old_error_reporting
	if i != 0 {
		core.PhpOutputGetContents(return_value)
		core.PhpOutputDiscard()
	} else {
		return_value.u1.type_info = 3
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

			if zend.ZendParseArgStr(_arg, &varname, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	val = zend.ZendIniGetValue(varname)
	if val == nil {
		return_value.u1.type_info = 2
		return
	}
	if (zend.ZvalGcFlags(val.gc.u.type_info) & 1 << 6) != 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = val
		__z.value.str = __s
		__z.u1.type_info = 6
	} else if val.len_ == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
	} else if val.len_ == 1 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendOneCharString[zend_uchar(val).val[0]]
		__z.value.str = __s
		__z.u1.type_info = 6
	} else if (zend.ZvalGcFlags(val.gc.u.type_info) & 1 << 7) == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringCopy(val)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(val.val, val.len_, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
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
			_optional = 1
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

			if zend.ZendParseArgString(_arg, &extname, &extname_len, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgBool(_arg, &details, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
	zend.ZendIniSortEntries()
	if extname != nil {
		if g.Assign(&module, zend.ZendHashStrFindPtr(&zend.ModuleRegistry, extname, extname_len)) == nil {
			core.PhpErrorDocref(nil, 1<<1, "Unable to find extension '%s'", extname)
			return_value.u1.type_info = 2
			return
		}
		module_number = module.module_number
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for {
		var __ht *zend.HashTable = zend.EG.ini_directives
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			key = _p.key
			ini_entry = _z.value.ptr
			var option zend.Zval
			if module_number != 0 && ini_entry.module_number != module_number {
				continue
			}
			if key == nil || key.val[0] != 0 {
				if details != 0 {
					var __arr *zend.ZendArray = zend._zendNewArray(0)
					var __z *zend.Zval = &option
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
					if ini_entry.orig_value != nil {
						zend.AddAssocStrEx(&option, "global_value", strlen("global_value"), zend.ZendStringCopy(ini_entry.orig_value))
					} else if ini_entry.value != nil {
						zend.AddAssocStrEx(&option, "global_value", strlen("global_value"), zend.ZendStringCopy(ini_entry.value))
					} else {
						zend.AddAssocNullEx(&option, "global_value", strlen("global_value"))
					}
					if ini_entry.value != nil {
						zend.AddAssocStrEx(&option, "local_value", strlen("local_value"), zend.ZendStringCopy(ini_entry.value))
					} else {
						zend.AddAssocNullEx(&option, "local_value", strlen("local_value"))
					}
					zend.AddAssocLongEx(&option, "access", strlen("access"), ini_entry.modifiable)
					zend.ZendSymtableUpdate(return_value.value.arr, ini_entry.name, &option)
				} else {
					if ini_entry.value != nil {
						var zv zend.Zval
						var __z *zend.Zval = &zv
						var __s *zend.ZendString = ini_entry.value
						__z.value.str = __s
						if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
							__z.u1.type_info = 6
						} else {
							zend.ZendGcAddref(&__s.gc)
							__z.u1.type_info = 6 | 1<<0<<8
						}
						zend.ZendSymtableUpdate(return_value.value.arr, ini_entry.name, &zv)
					} else {
						zend.ZendSymtableUpdate(return_value.value.arr, ini_entry.name, &zend.EG.uninitialized_zval)
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

			if zend.ZendParseArgStr(_arg, &varname, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgStr(_arg, &new_value, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	val = zend.ZendIniGetValue(varname)

	/* copy to return here, because alter might free it! */

	if val != nil {
		if (zend.ZvalGcFlags(val.gc.u.type_info) & 1 << 6) != 0 {
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = val
			__z.value.str = __s
			__z.u1.type_info = 6
		} else if val.len_ == 0 {
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendEmptyString
			__z.value.str = __s
			__z.u1.type_info = 6
		} else if val.len_ == 1 {
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendOneCharString[zend_uchar(val).val[0]]
			__z.value.str = __s
			__z.u1.type_info = 6
		} else if (zend.ZvalGcFlags(val.gc.u.type_info) & 1 << 7) == 0 {
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringCopy(val)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
		} else {
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(val.val, val.len_, 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
		}
	} else {
		return_value.u1.type_info = 2
	}

	// #define _CHECK_PATH(var,var_len,ini) php_ini_check_path ( var , var_len , ini , sizeof ( ini ) )

	/* open basedir check */

	if core.CoreGlobals.open_basedir != nil {
		if PhpIniCheckPath(varname.val, varname.len_, "error_log", g.SizeOf("\"error_log\"")) != 0 || PhpIniCheckPath(varname.val, varname.len_, "java.class.path", g.SizeOf("\"java.class.path\"")) != 0 || PhpIniCheckPath(varname.val, varname.len_, "java.home", g.SizeOf("\"java.home\"")) != 0 || PhpIniCheckPath(varname.val, varname.len_, "mail.log", g.SizeOf("\"mail.log\"")) != 0 || PhpIniCheckPath(varname.val, varname.len_, "java.library.path", g.SizeOf("\"java.library.path\"")) != 0 || PhpIniCheckPath(varname.val, varname.len_, "vpopmail.directory", g.SizeOf("\"vpopmail.directory\"")) != 0 {
			if core.PhpCheckOpenBasedir(new_value.val) != 0 {
				zend.ZvalPtrDtorStr(return_value)
				return_value.u1.type_info = 2
				return
			}
		}
	}
	if zend.ZendAlterIniEntryEx(varname, new_value, 1<<0, 1<<4, 0) == zend.FAILURE {
		zend.ZvalPtrDtorStr(return_value)
		return_value.u1.type_info = 2
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

			if zend.ZendParseArgStr(_arg, &varname, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	zend.ZendRestoreIniEntry(varname, 1<<4)
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

			if zend.ZendParseArgPathStr(_arg, &new_value, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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
	old_value = zend.ZendIniString("include_path", g.SizeOf("\"include_path\"")-1, 0)

	/* copy to return here, because alter might free it! */

	if old_value != nil {
		var _s *byte = old_value
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
	} else {
		return_value.u1.type_info = 2
	}
	key = zend.ZendStringInit("include_path", g.SizeOf("\"include_path\"")-1, 0)
	if zend.ZendAlterIniEntryEx(key, new_value, 1<<0, 1<<4, 0) == zend.FAILURE {
		zend.ZendStringReleaseEx(key, 0)
		zend.ZvalPtrDtorStr(return_value)
		return_value.u1.type_info = 2
		return
	}
	zend.ZendStringReleaseEx(key, 0)
}

/* }}} */

func ZifGetIncludePath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	str = zend.ZendIniString("include_path", g.SizeOf("\"include_path\"")-1, 0)
	if str == nil {
		return_value.u1.type_info = 2
		return
	}
	var _s *byte = str
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func ZifRestoreIncludePath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var key *zend.ZendString
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	key = zend.ZendStringInit("include_path", g.SizeOf("\"include_path\"")-1, 0)
	zend.ZendRestoreIniEntry(key, 1<<4)
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

			zend.ZendParseArgZvalDeref(_arg, &var_, 0)
			_optional = 1
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

			if zend.ZendParseArgBool(_arg, &do_return, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if do_return != 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendPrintZvalRToStr(var_, 0)
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	} else {
		zend.ZendPrintZvalR(var_, 0)
		return_value.u1.type_info = 3
		return
	}
}

/* }}} */

func ZifConnectionAborted(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var __z *zend.Zval = return_value
	__z.value.lval = core.CoreGlobals.connection_status & 1
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifConnectionStatus(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var __z *zend.Zval = return_value
	__z.value.lval = core.CoreGlobals.connection_status
	__z.u1.type_info = 4
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
			_optional = 1
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

			if zend.ZendParseArgBool(_arg, &arg, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
	old_setting = uint16(core.CoreGlobals.ignore_user_abort)
	if execute_data.This.u2.num_args != 0 {
		var key *zend.ZendString = zend.ZendStringInit("ignore_user_abort", g.SizeOf("\"ignore_user_abort\"")-1, 0)
		zend.ZendAlterIniEntryChars(key, g.Cond(arg != 0, "1", "0"), 1, 1<<0, 1<<4)
		zend.ZendStringReleaseEx(key, 0)
	}
	var __z *zend.Zval = return_value
	__z.value.lval = old_setting
	__z.u1.type_info = 4
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

			if zend.ZendParseArgString(_arg, &name, &name_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgString(_arg, &proto, &proto_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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

	/* empty string behaves like NULL on windows implementation of
	   getservbyname. Let be portable instead. */

	serv = getservbyname(name, proto)
	if serv == nil {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = ntohs(serv.s_port)
	__z.u1.type_info = 4
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

			if zend.ZendParseArgLong(_arg, &port, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgString(_arg, &proto, &proto_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	serv = getservbyport(htons(uint16(port)), proto)
	if serv == nil {
		return_value.u1.type_info = 2
		return
	}
	var _s *byte = serv.s_name
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
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

			if zend.ZendParseArgString(_arg, &name, &name_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	ent = getprotobyname(name)
	if ent == nil {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = ent.p_proto
	__z.u1.type_info = 4
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

			if zend.ZendParseArgLong(_arg, &proto, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
	ent = getprotobynumber(int(proto))
	if ent == nil {
		return_value.u1.type_info = 2
		return
	}
	var _s *byte = ent.p_name
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
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
	tick_fe.SetArgCount(execute_data.This.u2.num_args)
	if tick_fe.GetArgCount() < 1 {
		zend.ZendWrongParamCount()
		return
	}
	tick_fe.SetArguments((*zend.Zval)(zend._safeEmalloc(g.SizeOf("zval"), tick_fe.GetArgCount(), 0)))
	if zend._zendGetParametersArrayEx(tick_fe.GetArgCount(), tick_fe.GetArguments()) == zend.FAILURE {
		zend._efree(tick_fe.GetArguments())
		return_value.u1.type_info = 2
		return
	}
	if zend.ZendIsCallable(&tick_fe.arguments[0], 0, &function_name) == 0 {
		zend._efree(tick_fe.GetArguments())
		core.PhpErrorDocref(nil, 1<<1, "Invalid tick callback '%s' passed", function_name.val)
		zend.ZendStringReleaseEx(function_name, 0)
		return_value.u1.type_info = 2
		return
	} else if function_name != nil {
		zend.ZendStringReleaseEx(function_name, 0)
	}
	if tick_fe.GetArguments()[0].u1.v.type_ != 7 && tick_fe.GetArguments()[0].u1.v.type_ != 8 {
		if &tick_fe.arguments[0].u1.v.type_ != 6 {
			if &tick_fe.arguments[0].u1.v.type_ != 6 {
				zend._convertToString(&tick_fe.arguments[0])
			}
		}
	}
	if BasicGlobals.GetUserTickFunctions() == nil {
		BasicGlobals.SetUserTickFunctions((*zend.ZendLlist)(zend._emalloc(g.SizeOf("zend_llist"))))
		zend.ZendLlistInit(BasicGlobals.GetUserTickFunctions(), g.SizeOf("user_tick_function_entry"), zend.LlistDtorFuncT(UserTickFunctionDtor), 0)
		core.PhpAddTickFunction(RunUserTickFunctions, nil)
	}
	for i = 0; i < tick_fe.GetArgCount(); i++ {
		if &tick_fe.GetArguments()[i].u1.v.type_flags != 0 {
			zend.ZvalAddrefP(&tick_fe.GetArguments()[i])
		}
	}
	zend.ZendLlistAddElement(BasicGlobals.GetUserTickFunctions(), &tick_fe)
	return_value.u1.type_info = 3
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

			zend.ZendParseArgZvalDeref(_arg, &function, 0)
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
	if BasicGlobals.GetUserTickFunctions() == nil {
		return
	}
	if function.u1.v.type_ != 7 && function.u1.v.type_ != 8 {
		if function.u1.v.type_ != 6 {
			zend._convertToString(function)
		}
	}
	tick_fe.SetArguments((*zend.Zval)(zend._emalloc(g.SizeOf("zval"))))
	var _z1 *zend.Zval = &tick_fe.arguments[0]
	var _z2 *zend.Zval = function
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	tick_fe.SetArgCount(1)
	zend.ZendLlistDelElement(BasicGlobals.GetUserTickFunctions(), &tick_fe, (func(any, any) int)(UserTickFunctionCompare))
	zend._efree(tick_fe.GetArguments())
}

/* }}} */

func ZifIsUploadedFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var path *byte
	var path_len int
	if core.sapi_globals.rfc1867_uploaded_files == nil {
		return_value.u1.type_info = 2
		return
	}
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

			if zend.ZendParseArgString(_arg, &path, &path_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	if zend.ZendHashStrExists(core.sapi_globals.rfc1867_uploaded_files, path, path_len) != 0 {
		return_value.u1.type_info = 3
		return
	} else {
		return_value.u1.type_info = 2
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
	if core.sapi_globals.rfc1867_uploaded_files == nil {
		return_value.u1.type_info = 2
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
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

			if zend.ZendParseArgString(_arg, &path, &path_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgPath(_arg, &new_path, &new_path_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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
	if zend.ZendHashStrExists(core.sapi_globals.rfc1867_uploaded_files, path, path_len) == 0 {
		return_value.u1.type_info = 2
		return
	}
	if core.PhpCheckOpenBasedir(new_path) != 0 {
		return_value.u1.type_info = 2
		return
	}
	if r.Rename(path, new_path) == 0 {
		successful = 1
		oldmask = umask(077)
		umask(oldmask)
		ret = chmod(new_path, 0666 & ^oldmask)
		if ret == -1 {
			core.PhpErrorDocref(nil, 1<<1, "%s", strerror(errno))
		}
	} else if PhpCopyFileEx(path, new_path, 0x400) == zend.SUCCESS {
		unlink(path)
		successful = 1
	}
	if successful != 0 {
		zend.ZendHashStrDel(core.sapi_globals.rfc1867_uploaded_files, path, path_len)
	} else {
		core.PhpErrorDocref(nil, 1<<1, "Unable to move '%s' to '%s'", path, new_path)
	}
	if successful != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func PhpSimpleIniParserCb(arg1 *zend.Zval, arg2 *zend.Zval, arg3 *zend.Zval, callback_type int, arr *zend.Zval) {
	switch callback_type {
	case 1:
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}
		if arg2.u1.v.type_flags != 0 {
			zend.ZvalAddrefP(arg2)
		}
		zend.ZendSymtableUpdate(arr.value.arr, arg1.value.str, arg2)
		break
	case 3:
		var hash zend.Zval
		var find_hash *zend.Zval
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}
		if !(arg1.value.str.len_ > 1 && arg1.value.str.val[0] == '0') && zend.IsNumericString(arg1.value.str.val, arg1.value.str.len_, nil, nil, 0) == 4 {
			var key zend.ZendUlong = zend.ZendUlong(zend.ZendAtol(arg1.value.str.val, arg1.value.str.len_))
			if g.Assign(&find_hash, zend.ZendHashIndexFind(arr.value.arr, key)) == nil {
				var __arr *zend.ZendArray = zend._zendNewArray(0)
				var __z *zend.Zval = &hash
				__z.value.arr = __arr
				__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				find_hash = zend.ZendHashIndexAddNew(arr.value.arr, key, &hash)
			}
		} else {
			if g.Assign(&find_hash, zend.ZendHashFind(arr.value.arr, arg1.value.str)) == nil {
				var __arr *zend.ZendArray = zend._zendNewArray(0)
				var __z *zend.Zval = &hash
				__z.value.arr = __arr
				__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				find_hash = zend.ZendHashAddNew(arr.value.arr, arg1.value.str, &hash)
			}
		}
		if find_hash.u1.v.type_ != 7 {
			zend.ZvalPtrDtorNogc(find_hash)
			var __arr *zend.ZendArray = zend._zendNewArray(0)
			var __z *zend.Zval = find_hash
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		}
		if arg3 == nil || arg3.u1.v.type_ == 6 && arg3.value.str.len_ == 0 {
			if arg2.u1.v.type_flags != 0 {
				zend.ZvalAddrefP(arg2)
			}
			zend.AddNextIndexZval(find_hash, arg2)
		} else {
			zend.ArraySetZvalKey(find_hash.value.arr, arg3, arg2)
		}
		break
	case 2:
		break
	}
}

/* }}} */

func PhpIniParserCbWithSections(arg1 *zend.Zval, arg2 *zend.Zval, arg3 *zend.Zval, callback_type int, arr *zend.Zval) {
	if callback_type == 2 {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &(BasicGlobals.GetActiveIniFileSection())
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		zend.ZendSymtableUpdate(arr.value.arr, arg1.value.str, &(BasicGlobals.GetActiveIniFileSection()))
	} else if arg2 != nil {
		var active_arr *zend.Zval
		if BasicGlobals.active_ini_file_section.u1.v.type_ != 0 {
			active_arr = &(BasicGlobals.GetActiveIniFileSection())
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
	var scanner_mode zend.ZendLong = 0
	var fh zend.ZendFileHandle
	var ini_parser_cb zend.ZendIniParserCbT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
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

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgBool(_arg, &process_sections, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgLong(_arg, &scanner_mode, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if filename_len == 0 {
		core.PhpErrorDocref(nil, 1<<1, "Filename cannot be empty!")
		return_value.u1.type_info = 2
		return
	}

	/* Set callback function */

	if process_sections != 0 {
		&(BasicGlobals.GetActiveIniFileSection()).u1.type_info = 0
		ini_parser_cb = zend.ZendIniParserCbT(PhpIniParserCbWithSections)
	} else {
		ini_parser_cb = zend.ZendIniParserCbT(PhpSimpleIniParserCb)
	}

	/* Setup filehandle */

	zend.ZendStreamInitFilename(&fh, filename)
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if zend.ZendParseIniFile(&fh, 0, int(scanner_mode), ini_parser_cb, return_value) == zend.FAILURE {
		zend.ZendArrayDestroy(return_value.value.arr)
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifParseIniString(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var string *byte = nil
	var str *byte = nil
	var str_len int = 0
	var process_sections zend.ZendBool = 0
	var scanner_mode zend.ZendLong = 0
	var ini_parser_cb zend.ZendIniParserCbT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
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

			if zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgBool(_arg, &process_sections, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgLong(_arg, &scanner_mode, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if 2147483647-str_len < 32 {
		return_value.u1.type_info = 2
	}

	/* Set callback function */

	if process_sections != 0 {
		&(BasicGlobals.GetActiveIniFileSection()).u1.type_info = 0
		ini_parser_cb = zend.ZendIniParserCbT(PhpIniParserCbWithSections)
	} else {
		ini_parser_cb = zend.ZendIniParserCbT(PhpSimpleIniParserCb)
	}

	/* Setup string */

	string = (*byte)(zend._emalloc(str_len + 32))
	memcpy(string, str, str_len)
	memset(string+str_len, 0, 32)
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if zend.ZendParseIniString(string, 0, int(scanner_mode), ini_parser_cb, return_value) == zend.FAILURE {
		zend.ZendArrayDestroy(return_value.value.arr)
		return_value.u1.type_info = 2
	}
	zend._efree(string)
}
