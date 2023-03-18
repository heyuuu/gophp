// <<generate>>

package zend

import (
	r "sik/runtime"
	"sik/zend/types"
)

const ZEND_VERSION = "3.4.0"

type ZendSerializeData = __struct___zend_serialize_data
type ZendUnserializeData = __struct___zend_unserialize_data

/* output support */
var ZendPrintf func(format string, args ...any) int
var ZendWrite func(str string) int
var ZendFopen func(filename string, opened_path *string) *r.FILE
var ZendTicksFunction func(ticks int)
var ZendInterruptFunction func(executeData *ZendExecuteData)
var ZendErrorCb func(type_ int, error_filename string, error_lineno uint32, format string, args ...any)
var ZendOnTimeout func(seconds int)
var ZendStreamOpenFunction func(filename string, handle *ZendFileHandle) int
var ZendStreamOpenFunctionEx func(filename string, handle *ZendFileHandle) bool
var ZendPrintfToSmartStr func(buf *SmartStr, format string, ap ...any)
var ZendResolvePath func(filename string) *string

var ZendUv ZendUtilityValues

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

/* version information */

var ZendVersionInfo string

const ZEND_CORE_VERSION_INFO = "Zend Engine v" + ZEND_VERSION + ", Copyright (c) Zend Technologies\n"
const PRINT_ZVAL_INDENT = 4

/* true multithread-shared globals */

var ZendStandardClassDef *ZendClassEntry = nil
var ZendPostStartupCb func() int = nil
var ZendPostShutdownCb func() = nil
var ZendPreloadAutoload func(filename *types.ZendString) int = nil
var ZendMessageDispatcherP func(message ZendLong, data any)
var ZendGetConfigurationDirectiveP func(name string) *types.Zval

const SIGNAL_CHECK_DEFAULT = "0"

var IniEntries = []ZendIniEntryDef{
	*NewZendIniEntryDef("error_reporting", ZEND_INI_ALL).OnModify(OnUpdateErrorReportingEx),
	*NewZendIniEntryDef("zend.assertions", ZEND_INI_ALL).Value("1").OnModify(OnUpdateAssertionsEx),
	*NewZendIniEntryDef("zend.signal_check", ZEND_INI_SYSTEM).Value(SIGNAL_CHECK_DEFAULT).
		Displayer(ZendIniBooleanDisplayerCb).
		OnModify(func(entry *ZendIniEntry, new_value *string, stage int) bool {
			ZendSignalGlobals.check = ZendIniStringParseBool(*new_value)
			return true
		}),
	*NewZendIniEntryDef("zend.exception_ignore_args", ZEND_INI_ALL).Value("0").
		Displayer(ZendIniBooleanDisplayerCb).
		OnModify(func(entry *ZendIniEntry, new_value *string, stage int) bool {
			EG__().exception_ignore_args = types.intBool(ZendIniStringParseBool(*new_value))
			return true
		}),
}

const ShortTagsDefault = 1
const CompilerOptionsDefault uint32 = ZEND_COMPILE_DEFAULT
const COMPILED_STRING_DESCRIPTION_FORMAT = "%s(%d) : %s"
