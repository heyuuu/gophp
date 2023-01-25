// <<generate>>

package zend

import (
	b "sik/builtin"
	r "sik/runtime"
)

const ZEND_VERSION = "3.4.0"
const ZendSprintf = sprintf
const ZEND_TSRMG = TSRMG
const ZEND_TSRMG_FAST = TSRMG_FAST

type ZendSerializeData = __struct___zend_serialize_data
type ZendUnserializeData = __struct___zend_unserialize_data

type ZendWriteFuncT func(str *byte, str_length int) int

var ZendPrintf func(format *byte, _ ...any) int
var ZendWrite ZendWriteFuncT
var ZendFopen func(filename *byte, opened_path **ZendString) *r.FILE
var ZendTicksFunction func(ticks int)
var ZendInterruptFunction func(execute_data *ZendExecuteData)
var ZendErrorCb func(type_ int, error_filename *byte, error_lineno uint32, format *byte, args ...any)
var ZendOnTimeout func(seconds int)
var ZendStreamOpenFunction func(filename *byte, handle *ZendFileHandle) int
var ZendPrintfToSmartString func(buf *SmartString, format *byte, ap ...any)
var ZendPrintfToSmartStr func(buf *SmartStr, format *byte, ap ...any)
var ZendGetenv func(name *byte, name_len int) *byte
var ZendResolvePath func(filename *byte, filename_len int) *ZendString

const ZEND_STANDARD_CLASS_DEF_PTR = ZendStandardClassDef

var ZendUv ZendUtilityValues
var ZendDtraceEnabled ZendBool

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

var GlobalMapPtrLast int = 0

const GLOBAL_FUNCTION_TABLE *HashTable = CompilerGlobals.GetFunctionTable()
const GLOBAL_CLASS_TABLE *HashTable = CompilerGlobals.GetClassTable()
const GLOBAL_AUTO_GLOBALS_TABLE *HashTable = CompilerGlobals.GetAutoGlobals()
const GLOBAL_CONSTANTS_TABLE *HashTable = ExecutorGlobals.GetZendConstants()

var ZendVersionInfo *byte
var ZendVersionInfoLength uint32

const ZEND_CORE_VERSION_INFO = "Zend Engine v" + ZEND_VERSION + ", Copyright (c) Zend Technologies\n"
const PRINT_ZVAL_INDENT = 4

var ZendStandardClassDef *ZendClassEntry = nil
var ZendPostStartupCb func() int = nil
var ZendPostShutdownCb func() = nil
var ZendPreloadAutoload func(filename *ZendString) int = nil
var ZendMessageDispatcherP func(message ZendLong, data any)
var ZendGetConfigurationDirectiveP func(name *ZendString) *Zval

const SIGNAL_CHECK_DEFAULT = "0"

var IniEntries []ZendIniEntryDef = []ZendIniEntryDef{
	{"error_reporting", OnUpdateErrorReporting, nil, nil, nil, nil, nil, b.SizeOf("NULL") - 1, b.SizeOf("\"error_reporting\"") - 1, ZEND_INI_ALL},
	{
		"zend.assertions",
		OnUpdateAssertions,
		any(zend_long((*byte)(&((*ZendExecutorGlobals)(nil).GetAssertions())) - (*byte)(nil))),
		any(&ExecutorGlobals),
		nil,
		"1",
		nil,
		b.SizeOf("\"1\"") - 1,
		b.SizeOf("\"zend.assertions\"") - 1,
		ZEND_INI_ALL,
	},
	{"zend.enable_gc", OnUpdateGCEnabled, nil, nil, nil, "1", ZendGcEnabledDisplayerCb, b.SizeOf("\"1\"") - 1, b.SizeOf("\"zend.enable_gc\"") - 1, ZEND_INI_ALL},
	{
		"zend.multibyte",
		OnUpdateBool,
		any(zend_long((*byte)(&((*ZendCompilerGlobals)(nil).GetMultibyte())) - (*byte)(nil))),
		any(&CompilerGlobals),
		nil,
		"0",
		ZendIniBooleanDisplayerCb,
		b.SizeOf("\"0\"") - 1,
		b.SizeOf("\"zend.multibyte\"") - 1,
		ZEND_INI_PERDIR,
	},
	{"zend.script_encoding", OnUpdateScriptEncoding, nil, nil, nil, nil, nil, b.SizeOf("NULL") - 1, b.SizeOf("\"zend.script_encoding\"") - 1, ZEND_INI_ALL},
	{
		"zend.detect_unicode",
		OnUpdateBool,
		any(zend_long((*byte)(&((*ZendCompilerGlobals)(nil).GetDetectUnicode())) - (*byte)(nil))),
		any(&CompilerGlobals),
		nil,
		"1",
		ZendIniBooleanDisplayerCb,
		b.SizeOf("\"1\"") - 1,
		b.SizeOf("\"zend.detect_unicode\"") - 1,
		ZEND_INI_ALL,
	},
	{
		"zend.signal_check",
		OnUpdateBool,
		any(zend_long((*byte)(&((*ZendSignalGlobalsT)(nil).GetCheck())) - (*byte)(nil))),
		any(&ZendSignalGlobals),
		nil,
		SIGNAL_CHECK_DEFAULT,
		ZendIniBooleanDisplayerCb,
		b.SizeOf("SIGNAL_CHECK_DEFAULT") - 1,
		b.SizeOf("\"zend.signal_check\"") - 1,
		ZEND_INI_SYSTEM,
	},
	{
		"zend.exception_ignore_args",
		OnUpdateBool,
		any(zend_long((*byte)(&((*ZendExecutorGlobals)(nil).GetExceptionIgnoreArgs())) - (*byte)(nil))),
		any(&ExecutorGlobals),
		nil,
		"0",
		ZendIniBooleanDisplayerCb,
		b.SizeOf("\"0\"") - 1,
		b.SizeOf("\"zend.exception_ignore_args\"") - 1,
		ZEND_INI_ALL,
	},
	{nil, nil, nil, nil, nil, nil, nil, 0, 0, 0},
}

const ShortTagsDefault = 1
const CompilerOptionsDefault uint32 = ZEND_COMPILE_DEFAULT
const COMPILED_STRING_DESCRIPTION_FORMAT = "%s(%d) : %s"
