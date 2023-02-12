// <<generate>>

package zend

import (
	r "sik/runtime"
)

// Source: <Zend/zend.h>

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

const ZEND_VERSION = "3.4.0"
const ZendSprintf = sprintf
const ZEND_TSRMG = TSRMG
const ZEND_TSRMG_FAST = TSRMG_FAST

type ZendSerializeData = __struct___zend_serialize_data
type ZendUnserializeData = __struct___zend_unserialize_data

type ZendWriteFuncT func(str *byte, str_length int) int

/* Same as zend_spprintf and zend_strpprintf, without checking of format validity.
 * For use with custom printf specifiers such as %H. */

/* output support */

var ZendPrintf func(format *byte, _ ...any) int
var ZendWrite ZendWriteFuncT
var ZendFopen func(filename *byte, opened_path **ZendString) *r.FILE
var ZendFopenEx func(filename string, opened_path *string) *r.FILE
var ZendTicksFunction func(ticks int)
var ZendInterruptFunction func(execute_data *ZendExecuteData)
var ZendErrorCb func(type_ int, error_filename *byte, error_lineno uint32, format *byte, args ...any)
var ZendOnTimeout func(seconds int)
var ZendStreamOpenFunction func(filename string, handle *ZendFileHandle) int
var ZendStreamOpenFunctionEx func(filename string, handle *ZendFileHandle) bool
var ZendPrintfToSmartString func(buf *SmartString, format *byte, ap ...any)
var ZendPrintfToSmartStr func(buf *SmartStr, format *byte, ap ...any)
var ZendGetenv func(name *byte, name_len int) *byte
var ZendResolvePath func(filename *byte, filename_len int) *ZendString

/* These two callbacks are especially for opcache */

/* Callback for loading of not preloaded part of the script */

/* If filename is NULL the default filename is used. */

/* The following #define is __special__  used for code duality in PHP for Engine 1 & 2 */

const ZEND_STANDARD_CLASS_DEF_PTR = ZendStandardClassDef

var ZendUv ZendUtilityValues

/* If DTrace is available and enabled */

var ZendDtraceEnabled ZendBool

/* Messages for applications of Zend */

const ZMSG_FAILED_INCLUDE_FOPEN = 1
const ZMSG_FAILED_REQUIRE_FOPEN = 2
const ZMSG_FAILED_HIGHLIGHT_FOPEN = 3
const ZMSG_MEMORY_LEAK_DETECTED = 4
const ZMSG_MEMORY_LEAK_REPEATED = 5
const ZMSG_LOG_SCRIPT_NAME = 6
const ZMSG_MEMORY_LEAKS_GRAND_TOTAL = 7

type ZendErrorHandlingT = int

const (
	EH_NORMAL = 0
	EH_THROW
)
const DEBUG_BACKTRACE_PROVIDE_OBJECT ZendLong = 1 << 0
const DEBUG_BACKTRACE_IGNORE_ARGS ZendLong = 1 << 1

// Source: <Zend/zend.c>

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

var GlobalMapPtrLast int = 0

const GLOBAL_FUNCTION_TABLE *HashTable = CG__().GetFunctionTable()
const GLOBAL_CLASS_TABLE *HashTable = CG__().GetClassTable()
const GLOBAL_AUTO_GLOBALS_TABLE *HashTable = CG__().GetAutoGlobals()
const GLOBAL_CONSTANTS_TABLE *HashTable = EG__().GetZendConstants()

/* version information */

var ZendVersionInfo *byte
var ZendVersionInfoLength uint32

const ZEND_CORE_VERSION_INFO = "Zend Engine v" + ZEND_VERSION + ", Copyright (c) Zend Technologies\n"
const PRINT_ZVAL_INDENT = 4

/* true multithread-shared globals */

var ZendStandardClassDef *ZendClassEntry = nil
var ZendPostStartupCb func() int = nil
var ZendPostShutdownCb func() = nil
var ZendPreloadAutoload func(filename *ZendString) int = nil
var ZendMessageDispatcherP func(message ZendLong, data any)
var ZendGetConfigurationDirectiveP func(name *ZendString) *Zval

const SIGNAL_CHECK_DEFAULT = "0"

var IniEntries = []ZendIniEntryDef{
	*NewZendIniEntryDef("error_reporting", ZEND_INI_ALL).
		OnModifyArgs(
			OnUpdateErrorReporting, nil, nil, nil,
		),
	*NewZendIniEntryDef("zend.assertions", ZEND_INI_ALL).Value("1").
		OnModifyArgs(
			OnUpdateAssertions, any(zend_long((*byte)(&((*ZendExecutorGlobals)(nil).GetAssertions()))-(*byte)(nil))), any(&ExecutorGlobals), nil,
		),
	*NewZendIniEntryDef("zend.enable_gc", ZEND_INI_ALL).Value("1").
		Displayer(ZendGcEnabledDisplayerCb).
		OnModifyArgs(
			OnUpdateGCEnabled, nil, nil, nil,
		),
	*NewZendIniEntryDef("zend.multibyte", ZEND_INI_PERDIR).Value("0").
		Displayer(ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			OnUpdateBool, any(zend_long((*byte)(&((*ZendCompilerGlobals)(nil).GetMultibyte()))-(*byte)(nil))), any(&CompilerGlobals), nil,
		),
	*NewZendIniEntryDef("zend.script_encoding", ZEND_INI_ALL).
		OnModifyArgs(
			OnUpdateScriptEncoding, nil, nil, nil,
		),
	*NewZendIniEntryDef("zend.detect_unicode", ZEND_INI_ALL).Value("1").
		Displayer(ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			OnUpdateBool, any(zend_long((*byte)(&((*ZendCompilerGlobals)(nil).GetDetectUnicode()))-(*byte)(nil))), any(&CompilerGlobals), nil,
		),
	*NewZendIniEntryDef("zend.signal_check", ZEND_INI_SYSTEM).Value(SIGNAL_CHECK_DEFAULT).
		Displayer(ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			OnUpdateBool, any(zend_long((*byte)(&((*ZendSignalGlobalsT)(nil).GetCheck()))-(*byte)(nil))), any(&ZendSignalGlobals), nil,
		),
	*NewZendIniEntryDef("zend.exception_ignore_args", ZEND_INI_ALL).Value("0").
		Displayer(ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			OnUpdateBool, any(zend_long((*byte)(&((*ZendExecutorGlobals)(nil).GetExceptionIgnoreArgs()))-(*byte)(nil))), any(&ExecutorGlobals), nil,
		),
}

const ShortTagsDefault = 1
const CompilerOptionsDefault uint32 = ZEND_COMPILE_DEFAULT
const COMPILED_STRING_DESCRIPTION_FORMAT = "%s(%d) : %s"
