// <<generate>>

package zend

import g "sik/runtime/grammar"

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

// #define ZEND_H

// #define ZEND_VERSION       "3.4.0"

// #define ZEND_ENGINE_3

// # include "zend_types.h"

// # include "zend_map_ptr.h"

// # include "zend_errors.h"

// # include "zend_alloc.h"

// # include "zend_llist.h"

// # include "zend_string.h"

// # include "zend_hash.h"

// # include "zend_ast.h"

// # include "zend_gc.h"

// # include "zend_variables.h"

// # include "zend_iterators.h"

// # include "zend_stream.h"

// # include "zend_smart_str_public.h"

// # include "zend_smart_string_public.h"

// # include "zend_signal.h"

// #define zend_sprintf       sprintf

// #define HANDLE_BLOCK_INTERRUPTIONS() ZEND_SIGNAL_BLOCK_INTERRUPTIONS ( )

// #define HANDLE_UNBLOCK_INTERRUPTIONS() ZEND_SIGNAL_UNBLOCK_INTERRUPTIONS ( )

// #define INTERNAL_FUNCTION_PARAMETERS       zend_execute_data * execute_data , zval * return_value

// #define INTERNAL_FUNCTION_PARAM_PASSTHRU       execute_data , return_value

// #define USED_RET() ( ! EX ( prev_execute_data ) || ! ZEND_USER_CODE ( EX ( prev_execute_data ) -> func -> common . type ) || ( EX ( prev_execute_data ) -> opline -> result_type != IS_UNUSED ) )

// #define ZEND_TSRMG       TSRMG

// #define ZEND_TSRMG_FAST       TSRMG_FAST

// #define ZEND_TSRMLS_CACHE_EXTERN()

// #define ZEND_TSRMLS_CACHE_DEFINE()

// #define ZEND_TSRMLS_CACHE_UPDATE()

// #define ZEND_TSRMLS_CACHE

type ZendSerializeData = __struct___zend_serialize_data
type ZendUnserializeData = __struct___zend_unserialize_data

// @type ZendClassName struct

// @type ZendTraitMethodReference struct

// @type ZendTraitPrecedence struct

// @type ZendTraitAlias struct

// @type ZendClassEntry struct
// @type ZendUtilityFunctions struct

// @type ZendUtilityValues struct

type ZendWriteFuncT func(str *byte, str_length int) int

// #define zend_bailout() _zend_bailout ( __FILE__ , __LINE__ )

// #define zend_try       { JMP_BUF * __orig_bailout = EG ( bailout ) ; JMP_BUF __bailout ; EG ( bailout ) = & __bailout ; if ( SETJMP ( __bailout ) == 0 ) {

// #define zend_catch       } else { EG ( bailout ) = __orig_bailout ;

// #define zend_end_try() } EG ( bailout ) = __orig_bailout ; }

// #define zend_first_try       EG ( bailout ) = NULL ; zend_try

/* Same as zend_spprintf and zend_strpprintf, without checking of format validity.
 * For use with custom printf specifiers such as %H. */

// #define zend_print_variable(var) zend_print_zval ( ( var ) , 0 )

/* output support */

// #define ZEND_WRITE(str,str_len) zend_write ( ( str ) , ( str_len ) )

// #define ZEND_WRITE_EX(str,str_len) write_func ( ( str ) , ( str_len ) )

// #define ZEND_PUTS(str) zend_write ( ( str ) , strlen ( ( str ) ) )

// #define ZEND_PUTS_EX(str) write_func ( ( str ) , strlen ( ( str ) ) )

// #define ZEND_PUTC(c) zend_write ( & ( c ) , 1 )

var ZendPrintf func(format *byte, _ ...any) int
var ZendWrite ZendWriteFuncT
var ZendFopen func(filename *byte, opened_path **ZendString) *FILE
var ZendTicksFunction func(ticks int)
var ZendInterruptFunction func(execute_data *ZendExecuteData)
var ZendErrorCb func(type_ int, error_filename *byte, error_lineno uint32, format *byte, args va_list)
var ZendOnTimeout func(seconds int)
var ZendStreamOpenFunction func(filename *byte, handle *ZendFileHandle) int
var ZendPrintfToSmartString func(buf *SmartString, format *byte, ap va_list)
var ZendPrintfToSmartStr func(buf *SmartStr, format *byte, ap va_list)
var ZendGetenv func(name *byte, name_len int) *byte
var ZendResolvePath func(filename *byte, filename_len int) *ZendString

/* These two callbacks are especially for opcache */

/* Callback for loading of not preloaded part of the script */

/* If filename is NULL the default filename is used. */

/* The following #define is __special__  used for code duality in PHP for Engine 1 & 2 */

// #define ZEND_STANDARD_CLASS_DEF_PTR       zend_standard_class_def

var ZendUv ZendUtilityValues

/* If DTrace is available and enabled */

var ZendDtraceEnabled ZendBool

// #define ZEND_UV(name) ( zend_uv . name )

/* Messages for applications of Zend */

// #define ZMSG_FAILED_INCLUDE_FOPEN       1L

// #define ZMSG_FAILED_REQUIRE_FOPEN       2L

// #define ZMSG_FAILED_HIGHLIGHT_FOPEN       3L

// #define ZMSG_MEMORY_LEAK_DETECTED       4L

// #define ZMSG_MEMORY_LEAK_REPEATED       5L

// #define ZMSG_LOG_SCRIPT_NAME       6L

// #define ZMSG_MEMORY_LEAKS_GRAND_TOTAL       7L

type ZendErrorHandlingT = int

const (
	EH_NORMAL = 0
	EH_THROW
)

// @type ZendErrorHandling struct

// #define DEBUG_BACKTRACE_PROVIDE_OBJECT       ( 1 << 0 )

// #define DEBUG_BACKTRACE_IGNORE_ARGS       ( 1 << 1 )

// # include "zend_object_handlers.h"

// # include "zend_operators.h"

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

// # include "zend.h"

// # include "zend_extensions.h"

// # include "zend_modules.h"

// # include "zend_constants.h"

// # include "zend_list.h"

// # include "zend_API.h"

// # include "zend_exceptions.h"

// # include "zend_builtin_functions.h"

// # include "zend_ini.h"

// # include "zend_vm.h"

// # include "zend_dtrace.h"

// # include "zend_virtual_cwd.h"

// # include "zend_smart_str.h"

// # include "zend_smart_string.h"

// # include "zend_cpuinfo.h"

var GlobalMapPtrLast int = 0

// #define GLOBAL_FUNCTION_TABLE       CG ( function_table )

// #define GLOBAL_CLASS_TABLE       CG ( class_table )

// #define GLOBAL_AUTO_GLOBALS_TABLE       CG ( auto_globals )

// #define GLOBAL_CONSTANTS_TABLE       EG ( zend_constants )

/* version information */

var ZendVersionInfo *byte
var ZendVersionInfoLength uint32

// #define ZEND_CORE_VERSION_INFO       "Zend Engine v" ZEND_VERSION ", Copyright (c) Zend Technologies\n"

// #define PRINT_ZVAL_INDENT       4

/* true multithread-shared globals */

var ZendStandardClassDef *ZendClassEntry = nil
var ZendPostStartupCb func() int = nil
var ZendPostShutdownCb func() = nil
var ZendPreloadAutoload func(filename *ZendString) int = nil
var ZendMessageDispatcherP func(message ZendLong, data any)
var ZendGetConfigurationDirectiveP func(name *ZendString) *Zval

func OnUpdateErrorReporting(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	if new_value == nil {
		EG.SetErrorReporting((1<<0 | 1<<1 | 1<<2 | 1<<3 | 1<<4 | 1<<5 | 1<<6 | 1<<7 | 1<<8 | 1<<9 | 1<<10 | 1<<12 | 1<<13 | 1<<14 | 1<<11) & ^(1 << 3) & ^(1 << 11) & ^(1 << 13))
	} else {
		EG.SetErrorReporting(atoi(new_value.GetVal()))
	}
	return SUCCESS
}

/* }}} */

func OnUpdateGCEnabled(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var val ZendBool
	val = ZendIniParseBool(new_value)
	GcEnable(val)
	return SUCCESS
}

/* }}} */

func ZendGcEnabledDisplayerCb(ini_entry *ZendIniEntry, type_ int) {
	if GcEnabled() != 0 {
		ZendWrite("On", strlen("On"))
	} else {
		ZendWrite("Off", strlen("Off"))
	}
}

/* }}} */

func OnUpdateScriptEncoding(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	if CG.GetMultibyte() == 0 {
		return FAILURE
	}
	if ZendMultibyteGetFunctions() == nil {
		return SUCCESS
	}
	return ZendMultibyteSetScriptEncodingByString(g.CondF1(new_value != nil, func() []byte { return new_value.GetVal() }, nil), g.CondF1(new_value != nil, func() int { return new_value.GetLen() }, 0))
}

/* }}} */

func OnUpdateAssertions(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p *ZendLong
	var val ZendLong
	var base *byte = (*byte)(mh_arg2)
	p = (*ZendLong)(base + int(mh_arg1))
	val = ZendAtol(new_value.GetVal(), new_value.GetLen())
	if stage != 1<<0 && stage != 1<<1 && (*p) != val && ((*p) < 0 || val < 0) {
		ZendError(1<<1, "zend.assertions may be completely enabled or disabled only in php.ini")
		return FAILURE
	}
	*p = val
	return SUCCESS
}

/* }}} */

// #define SIGNAL_CHECK_DEFAULT       "0"

var IniEntries []ZendIniEntryDef = []ZendIniEntryDef{{"error_reporting", OnUpdateErrorReporting, nil, nil, nil, nil, nil, g.SizeOf("NULL") - 1, g.SizeOf("\"error_reporting\"") - 1, 1<<0 | 1<<1 | 1<<2}, {"zend.assertions", OnUpdateAssertions, any(zend_long((*byte)(&((*ZendExecutorGlobals)(nil).GetAssertions())) - (*byte)(nil))), any(&ExecutorGlobals), nil, "1", nil, g.SizeOf("\"1\"") - 1, g.SizeOf("\"zend.assertions\"") - 1, 1<<0 | 1<<1 | 1<<2}, {"zend.enable_gc", OnUpdateGCEnabled, nil, nil, nil, "1", ZendGcEnabledDisplayerCb, g.SizeOf("\"1\"") - 1, g.SizeOf("\"zend.enable_gc\"") - 1, 1<<0 | 1<<1 | 1<<2}, {"zend.multibyte", OnUpdateBool, any(zend_long((*byte)(&((*ZendCompilerGlobals)(nil).GetMultibyte())) - (*byte)(nil))), any(&CompilerGlobals), nil, "0", ZendIniBooleanDisplayerCb, g.SizeOf("\"0\"") - 1, g.SizeOf("\"zend.multibyte\"") - 1, 1 << 1}, {"zend.script_encoding", OnUpdateScriptEncoding, nil, nil, nil, nil, nil, g.SizeOf("NULL") - 1, g.SizeOf("\"zend.script_encoding\"") - 1, 1<<0 | 1<<1 | 1<<2}, {"zend.detect_unicode", OnUpdateBool, any(zend_long((*byte)(&((*ZendCompilerGlobals)(nil).GetDetectUnicode())) - (*byte)(nil))), any(&CompilerGlobals), nil, "1", ZendIniBooleanDisplayerCb, g.SizeOf("\"1\"") - 1, g.SizeOf("\"zend.detect_unicode\"") - 1, 1<<0 | 1<<1 | 1<<2}, {"zend.signal_check", OnUpdateBool, any(zend_long((*byte)(&((*ZendSignalGlobalsT)(nil).GetCheck())) - (*byte)(nil))), any(&ZendSignalGlobals), nil, "0", ZendIniBooleanDisplayerCb, g.SizeOf("\"0\"") - 1, g.SizeOf("\"zend.signal_check\"") - 1, 1 << 2}, {"zend.exception_ignore_args", OnUpdateBool, any(zend_long((*byte)(&((*ZendExecutorGlobals)(nil).GetExceptionIgnoreArgs())) - (*byte)(nil))), any(&ExecutorGlobals), nil, "0", ZendIniBooleanDisplayerCb, g.SizeOf("\"0\"") - 1, g.SizeOf("\"zend.exception_ignore_args\"") - 1, 1<<0 | 1<<1 | 1<<2}, {nil, nil, nil, nil, nil, nil, nil, 0, 0, 0}}

func ZendVspprintf(pbuf **byte, max_len int, format string, ap va_list) int {
	var buf SmartString = SmartString{0}

	/* since there are places where (v)spprintf called without checking for null,
	   a bit of defensive coding here */

	if pbuf == nil {
		return 0
	}
	ZendPrintfToSmartString(&buf, format, ap)
	if max_len != 0 && buf.GetLen() > max_len {
		buf.SetLen(max_len)
	}
	SmartString0(&buf)
	if buf.GetC() != nil {
		*pbuf = buf.GetC()
		return buf.GetLen()
	} else {
		*pbuf = _estrndup("", 0)
		return 0
	}
}

/* }}} */

func ZendSpprintf(message **byte, max_len int, format string, _ ...any) int {
	var arg va_list
	var len_ int
	va_start(arg, format)
	len_ = ZendVspprintf(message, max_len, format, arg)
	va_end(arg)
	return len_
}

/* }}} */

func ZendSpprintfUnchecked(message **byte, max_len int, format *byte, _ ...any) int {
	var arg va_list
	var len_ int
	va_start(arg, format)
	len_ = ZendVspprintf(message, max_len, format, arg)
	va_end(arg)
	return len_
}

/* }}} */

func ZendVstrpprintf(max_len int, format *byte, ap va_list) *ZendString {
	var buf SmartStr = SmartStr{0}
	ZendPrintfToSmartStr(&buf, format, ap)
	if buf.GetS() == nil {
		return ZendEmptyString
	}
	if max_len != 0 && buf.GetS().GetLen() > max_len {
		buf.GetS().SetLen(max_len)
	}
	SmartStr0(&buf)
	return buf.GetS()
}

/* }}} */

func ZendStrpprintf(max_len int, format string, _ ...any) *ZendString {
	var arg va_list
	var str *ZendString
	va_start(arg, format)
	str = ZendVstrpprintf(max_len, format, arg)
	va_end(arg)
	return str
}

/* }}} */

func ZendStrpprintfUnchecked(max_len int, format string, _ ...any) *ZendString {
	var arg va_list
	var str *ZendString
	va_start(arg, format)
	str = ZendVstrpprintf(max_len, format, arg)
	va_end(arg)
	return str
}

/* }}} */

func PrintHash(buf *SmartStr, ht *HashTable, indent int, is_object ZendBool) {
	var tmp *Zval
	var string_key *ZendString
	var num_key ZendUlong
	var i int
	for i = 0; i < indent; i++ {
		SmartStrAppendcEx(buf, ' ', 0)
	}
	SmartStrAppendlEx(buf, "(\n", strlen("(\n"), 0)
	indent += 4
	for {
		var __ht *HashTable = ht
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val
			if _z.GetType() == 13 {
				_z = _z.GetValue().GetZv()
			}
			if _z.GetType() == 0 {
				continue
			}
			num_key = _p.GetH()
			string_key = _p.GetKey()
			tmp = _z
			for i = 0; i < indent; i++ {
				SmartStrAppendcEx(buf, ' ', 0)
			}
			SmartStrAppendcEx(buf, '[', 0)
			if string_key != nil {
				if is_object != 0 {
					var prop_name *byte
					var class_name *byte
					var prop_len int
					var mangled int = ZendUnmanglePropertyNameEx(string_key, &class_name, &prop_name, &prop_len)
					SmartStrAppendlEx(buf, prop_name, prop_len, 0)
					if class_name != nil && mangled == SUCCESS {
						if class_name[0] == '*' {
							SmartStrAppendlEx(buf, ":protected", strlen(":protected"), 0)
						} else {
							SmartStrAppendlEx(buf, ":", strlen(":"), 0)
							SmartStrAppendlEx(buf, class_name, strlen(class_name), 0)
							SmartStrAppendlEx(buf, ":private", strlen(":private"), 0)
						}
					}
				} else {
					SmartStrAppendEx(buf, string_key, 0)
				}
			} else {
				SmartStrAppendLongEx(buf, num_key, 0)
			}
			SmartStrAppendlEx(buf, "] => ", strlen("] => "), 0)
			ZendPrintZvalRToBuf(buf, tmp, indent+4)
			SmartStrAppendlEx(buf, "\n", strlen("\n"), 0)
		}
		break
	}
	indent -= 4
	for i = 0; i < indent; i++ {
		SmartStrAppendcEx(buf, ' ', 0)
	}
	SmartStrAppendlEx(buf, ")\n", strlen(")\n"), 0)
}

/* }}} */

func PrintFlatHash(ht *HashTable) {
	var tmp *Zval
	var string_key *ZendString
	var num_key ZendUlong
	var i int = 0
	for {
		var __ht *HashTable = ht
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val
			if _z.GetType() == 13 {
				_z = _z.GetValue().GetZv()
			}
			if _z.GetType() == 0 {
				continue
			}
			num_key = _p.GetH()
			string_key = _p.GetKey()
			tmp = _z
			if g.PostInc(&i) > 0 {
				ZendWrite(",", strlen(","))
			}
			ZendWrite("[", strlen("["))
			if string_key != nil {
				ZendWrite(string_key.GetVal(), string_key.GetLen())
			} else {
				ZendPrintf("%"+"llu", num_key)
			}
			ZendWrite("] => ", strlen("] => "))
			ZendPrintFlatZvalR(tmp)
		}
		break
	}
}

/* }}} */

func ZendMakePrintableZval(expr *Zval, expr_copy *Zval) int {
	if expr.GetType() == 6 {
		return 0
	} else {
		var __z *Zval = expr_copy
		var __s *ZendString = ZvalGetStringFunc(expr)
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		return 1
	}
}

/* }}} */

func ZendPrintZval(expr *Zval, indent int) int {
	var tmp_str *ZendString
	var str *ZendString = ZvalGetTmpString(expr, &tmp_str)
	var len_ int = str.GetLen()
	if len_ != 0 {
		ZendWrite(str.GetVal(), len_)
	}
	ZendTmpStringRelease(tmp_str)
	return len_
}

/* }}} */

func ZendPrintFlatZvalR(expr *Zval) {
	switch expr.GetType() {
	case 7:
		ZendWrite("Array (", strlen("Array ("))
		if (ZvalGcFlags(expr.GetValue().GetArr().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
			if (ZvalGcFlags(expr.GetValue().GetArr().GetGc().GetTypeInfo()) & 1 << 5) != 0 {
				ZendWrite(" *RECURSION*", strlen(" *RECURSION*"))
				return
			}
			expr.GetValue().GetArr().GetGc().SetTypeInfo(expr.GetValue().GetArr().GetGc().GetTypeInfo() | 1<<5<<0)
		}
		PrintFlatHash(expr.GetValue().GetArr())
		ZendWrite(")", strlen(")"))
		if (ZvalGcFlags(expr.GetValue().GetArr().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
			expr.GetValue().GetArr().GetGc().SetTypeInfo(expr.GetValue().GetArr().GetGc().GetTypeInfo() &^ (1 << 5 << 0))
		}
		break
	case 8:
		var properties *HashTable
		var class_name *ZendString = expr.GetValue().GetObj().GetHandlers().GetGetClassName()(expr.GetValue().GetObj())
		ZendPrintf("%s Object (", class_name.GetVal())
		ZendStringReleaseEx(class_name, 0)
		if (ZvalGcFlags(expr.GetValue().GetCounted().GetGc().GetTypeInfo()) & 1 << 5) != 0 {
			ZendWrite(" *RECURSION*", strlen(" *RECURSION*"))
			return
		}
		properties = expr.GetValue().GetObj().GetHandlers().GetGetProperties()(&(*expr))
		if properties != nil {
			expr.GetValue().GetObj().GetGc().SetTypeInfo(expr.GetValue().GetObj().GetGc().GetTypeInfo() | 1<<5<<0)
			PrintFlatHash(properties)
			expr.GetValue().GetObj().GetGc().SetTypeInfo(expr.GetValue().GetObj().GetGc().GetTypeInfo() &^ (1 << 5 << 0))
		}
		ZendWrite(")", strlen(")"))
		break
	case 10:
		ZendPrintFlatZvalR(&(*expr).value.GetRef().GetVal())
		break
	default:
		ZendPrintZval(expr, 0)
		break
	}
}

/* }}} */

func ZendPrintZvalRToBuf(buf *SmartStr, expr *Zval, indent int) {
	switch expr.GetType() {
	case 7:
		SmartStrAppendlEx(buf, "Array\n", strlen("Array\n"), 0)
		if (ZvalGcFlags(expr.GetValue().GetArr().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
			if (ZvalGcFlags(expr.GetValue().GetArr().GetGc().GetTypeInfo()) & 1 << 5) != 0 {
				SmartStrAppendlEx(buf, " *RECURSION*", strlen(" *RECURSION*"), 0)
				return
			}
			expr.GetValue().GetArr().GetGc().SetTypeInfo(expr.GetValue().GetArr().GetGc().GetTypeInfo() | 1<<5<<0)
		}
		PrintHash(buf, expr.GetValue().GetArr(), indent, 0)
		if (ZvalGcFlags(expr.GetValue().GetArr().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
			expr.GetValue().GetArr().GetGc().SetTypeInfo(expr.GetValue().GetArr().GetGc().GetTypeInfo() &^ (1 << 5 << 0))
		}
		break
	case 8:
		var properties *HashTable
		var class_name *ZendString = expr.GetValue().GetObj().GetHandlers().GetGetClassName()(expr.GetValue().GetObj())
		SmartStrAppendlEx(buf, class_name.GetVal(), strlen(class_name.GetVal()), 0)
		ZendStringReleaseEx(class_name, 0)
		SmartStrAppendlEx(buf, " Object\n", strlen(" Object\n"), 0)
		if (ZvalGcFlags(expr.GetValue().GetObj().GetGc().GetTypeInfo()) & 1 << 5) != 0 {
			SmartStrAppendlEx(buf, " *RECURSION*", strlen(" *RECURSION*"), 0)
			return
		}
		if g.Assign(&properties, ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_DEBUG)) == nil {
			break
		}
		expr.GetValue().GetObj().GetGc().SetTypeInfo(expr.GetValue().GetObj().GetGc().GetTypeInfo() | 1<<5<<0)
		PrintHash(buf, properties, indent, 1)
		expr.GetValue().GetObj().GetGc().SetTypeInfo(expr.GetValue().GetObj().GetGc().GetTypeInfo() &^ (1 << 5 << 0))
		if properties != nil && (ZvalGcFlags(properties.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcDelref(&properties.gc) == 0 {
			ZendArrayDestroy(properties)
		}
		break
	case 4:
		SmartStrAppendLongEx(buf, expr.GetValue().GetLval(), 0)
		break
	case 10:
		ZendPrintZvalRToBuf(buf, &(*expr).value.GetRef().GetVal(), indent)
		break
	case 6:
		SmartStrAppendEx(buf, expr.GetValue().GetStr(), 0)
		break
	default:
		var str *ZendString = ZvalGetStringFunc(expr)
		SmartStrAppendEx(buf, str, 0)
		ZendStringReleaseEx(str, 0)
		break
	}
}

/* }}} */

func ZendPrintZvalRToStr(expr *Zval, indent int) *ZendString {
	var buf SmartStr = SmartStr{0}
	ZendPrintZvalRToBuf(&buf, expr, indent)
	SmartStr0(&buf)
	return buf.GetS()
}

/* }}} */

func ZendPrintZvalR(expr *Zval, indent int) {
	var str *ZendString = ZendPrintZvalRToStr(expr, indent)
	ZendWrite(str.GetVal(), str.GetLen())
	ZendStringReleaseEx(str, 0)
}

/* }}} */

func ZendFopenWrapper(filename *byte, opened_path **ZendString) *FILE {
	if opened_path != nil {
		*opened_path = ZendStringInit(filename, strlen(filename), 0)
	}
	return fopen(filename, "rb")
}

/* }}} */

// #define short_tags_default       1

// #define compiler_options_default       ZEND_COMPILE_DEFAULT

func ZendSetDefaultCompileTimeValues() {
	/* default compile-time values */

	CG.SetShortTags(1)
	CG.SetCompilerOptions(1 << 2)
	CG.SetRtdKeyCounter(0)
}

/* }}} */

func ZendInitExceptionOp() {
	memset(EG.GetExceptionOp(), 0, g.SizeOf("EG ( exception_op )"))
	EG.GetExceptionOp()[0].SetOpcode(149)
	ZendVmSetOpcodeHandler(EG.GetExceptionOp())
	EG.GetExceptionOp()[1].SetOpcode(149)
	ZendVmSetOpcodeHandler(EG.GetExceptionOp() + 1)
	EG.GetExceptionOp()[2].SetOpcode(149)
	ZendVmSetOpcodeHandler(EG.GetExceptionOp() + 2)
}

/* }}} */

func ZendInitCallTrampolineOp() {
	memset(&EG.call_trampoline_op, 0, g.SizeOf("EG ( call_trampoline_op )"))
	EG.GetCallTrampolineOp().SetOpcode(158)
	ZendVmSetOpcodeHandler(&EG.call_trampoline_op)
}

/* }}} */

func AutoGlobalDtor(zv *Zval) { Free(zv.GetValue().GetPtr()) }

/* }}} */

func IniScannerGlobalsCtor(scanner_globals_p *ZendIniScannerGlobals) {
	memset(scanner_globals_p, 0, g.SizeOf("* scanner_globals_p"))
}

/* }}} */

func PhpScannerGlobalsCtor(scanner_globals_p *ZendPhpScannerGlobals) {
	memset(scanner_globals_p, 0, g.SizeOf("* scanner_globals_p"))
}

/* }}} */

func ModuleDestructorZval(zv *Zval) {
	var module *ZendModuleEntry = (*ZendModuleEntry)(zv.GetValue().GetPtr())
	ModuleDestructor(module)
	Free(module)
}

/* }}} */

func PhpAutoGlobalsCreateGlobals(name *ZendString) ZendBool {
	var globals Zval

	/* IS_ARRAY, but with ref-counter 1 and not IS_TYPE_REFCOUNTED */

	var __arr *ZendArray = &EG.symbol_table
	var __z *Zval = &globals
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	&globals.SetTypeFlags(0)
	var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
	ZendGcSetRefcount(&_ref.gc, 1)
	_ref.GetGc().SetTypeInfo(10)
	var _z1 *Zval = &_ref.val
	var _z2 *Zval = &globals
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	_ref.GetSources().SetPtr(nil)
	&globals.GetValue().SetRef(_ref)
	&globals.SetTypeInfo(10 | 1<<0<<8)
	ZendHashUpdate(&EG.symbol_table, name, &globals)
	return 0
}

/* }}} */

func ZendStartup(utility_functions *ZendUtilityFunctions) int {
	var ini_scanner_globals ZendIniScannerGlobals
	var language_scanner_globals ZendPhpScannerGlobals
	ZendCpuStartup()
	StartMemoryManager()
	VirtualCwdStartup()
	ZendStartupStrtod()
	ZendStartupExtensionsMechanism()

	/* Set up utility functions and values */

	ZendErrorCb = utility_functions.GetErrorFunction()
	ZendPrintf = utility_functions.GetPrintfFunction()
	ZendWrite = ZendWriteFuncT(utility_functions.GetWriteFunction())
	ZendFopen = utility_functions.GetFopenFunction()
	if ZendFopen == nil {
		ZendFopen = ZendFopenWrapper
	}
	ZendStreamOpenFunction = utility_functions.GetStreamOpenFunction()
	ZendMessageDispatcherP = utility_functions.GetMessageHandler()
	ZendGetConfigurationDirectiveP = utility_functions.GetGetConfigurationDirective()
	ZendTicksFunction = utility_functions.GetTicksFunction()
	ZendOnTimeout = utility_functions.GetOnTimeout()
	ZendPrintfToSmartString = utility_functions.GetPrintfToSmartStringFunction()
	ZendPrintfToSmartStr = utility_functions.GetPrintfToSmartStrFunction()
	ZendGetenv = utility_functions.GetGetenvFunction()
	ZendResolvePath = utility_functions.GetResolvePathFunction()
	ZendInterruptFunction = nil
	ZendCompileFile = CompileFile
	ZendExecuteEx = ExecuteEx
	ZendExecuteInternal = nil
	ZendCompileString = CompileString
	ZendThrowExceptionHook = nil

	/* Set up the default garbage collection implementation. */

	GcCollectCycles = ZendGcCollectCycles
	ZendVmInit()

	/* set up version */

	ZendVersionInfo = strdup("Zend Engine v" + "3.4.0" + ", Copyright (c) Zend Technologies\n")
	ZendVersionInfoLength = g.SizeOf("ZEND_CORE_VERSION_INFO") - 1
	CG.SetFunctionTable((*HashTable)(Malloc(g.SizeOf("HashTable"))))
	CG.SetClassTable((*HashTable)(Malloc(g.SizeOf("HashTable"))))
	CG.SetAutoGlobals((*HashTable)(Malloc(g.SizeOf("HashTable"))))
	EG.SetZendConstants((*HashTable)(Malloc(g.SizeOf("HashTable"))))
	_zendHashInit(CG.GetFunctionTable(), 1024, ZendFunctionDtor, 1)
	_zendHashInit(CG.GetClassTable(), 64, DestroyZendClass, 1)
	_zendHashInit(CG.GetAutoGlobals(), 8, AutoGlobalDtor, 1)
	_zendHashInit(EG.GetZendConstants(), 128, FreeZendConstant, 1)
	_zendHashInit(&ModuleRegistry, 32, ModuleDestructorZval, 1)
	ZendInitRsrcListDtors()
	IniScannerGlobalsCtor(&ini_scanner_globals)
	PhpScannerGlobalsCtor(&language_scanner_globals)
	ZendSetDefaultCompileTimeValues()

	/* Map region is going to be created and resized at run-time. */

	CG.SetMapPtrBase(nil)
	CG.SetMapPtrSize(0)
	CG.SetMapPtrLast(0)
	EG.SetErrorReporting((1<<0 | 1<<1 | 1<<2 | 1<<3 | 1<<4 | 1<<5 | 1<<6 | 1<<7 | 1<<8 | 1<<9 | 1<<10 | 1<<12 | 1<<13 | 1<<14 | 1<<11) & ^(1 << 3))
	ZendInternedStringsInit()
	ZendStartupBuiltinFunctions()
	ZendRegisterStandardConstants()
	ZendRegisterAutoGlobal(ZendStringInitInterned("GLOBALS", g.SizeOf("\"GLOBALS\"")-1, 1), 1, PhpAutoGlobalsCreateGlobals)
	ZendInitRsrcPlist()
	ZendInitExceptionOp()
	ZendInitCallTrampolineOp()
	ZendIniStartup()
	return SUCCESS
}

/* }}} */

func ZendRegisterStandardIniEntries() {
	var module_number int = 0
	ZendRegisterIniEntries(IniEntries, module_number)
}

/* }}} */

func ZendResolvePropertyTypes() {
	var ce *ZendClassEntry
	var prop_info *ZendPropertyInfo
	for {
		var __ht *HashTable = CG.GetClassTable()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			ce = _z.GetValue().GetPtr()
			if ce.GetType() != 1 {
				continue
			}
			if (ce.GetCeFlags() & 1 << 8) == 1<<8 {
				for {
					var __ht *HashTable = &ce.properties_info
					var _p *Bucket = __ht.GetArData()
					var _end *Bucket = _p + __ht.GetNNumUsed()
					for ; _p != _end; _p++ {
						var _z *Zval = &_p.val

						if _z.GetType() == 0 {
							continue
						}
						prop_info = _z.GetValue().GetPtr()
						if prop_info.GetType() > 0x3ff && (prop_info.GetType()&0x2) == 0 {
							var type_name *ZendString = (*ZendString)(prop_info.GetType() & ^0x3)
							var lc_type_name *ZendString = ZendStringTolowerEx(type_name, 0)
							var prop_ce *ZendClassEntry = ZendHashFindPtr(CG.GetClassTable(), lc_type_name)
							assert(prop_ce != nil && prop_ce.GetType() == 1)
							prop_info.SetType(uintptr_t(prop_ce) | g.Cond((prop_info.GetType()&0x1) != 0, 0x3, 0x2))
							ZendStringRelease(lc_type_name)
							ZendStringRelease(type_name)
						}
					}
					break
				}
			}
			ce.SetCeFlags(ce.GetCeFlags() | 1<<17)
		}
		break
	}
}

/* }}} */

func ZendPostStartup() int {
	ZendResolvePropertyTypes()
	if ZendPostStartupCb != nil {
		var cb func() int = ZendPostStartupCb
		ZendPostStartupCb = nil
		if cb() != SUCCESS {
			return FAILURE
		}
	}
	GlobalMapPtrLast = CG.GetMapPtrLast()
	return SUCCESS
}

/* }}} */

func ZendShutdown() {
	ZendVmDtor()
	ZendDestroyRsrcList(&EG.persistent_list)
	ZendDestroyModules()
	VirtualCwdDeactivate()
	VirtualCwdShutdown()
	ZendHashDestroy(CG.GetFunctionTable())
	ZendHashDestroy(CG.GetClassTable())
	ZendHashDestroy(CG.GetAutoGlobals())
	Free(CG.GetAutoGlobals())
	ZendShutdownExtensions()
	Free(ZendVersionInfo)
	Free(CG.GetFunctionTable())
	Free(CG.GetClassTable())
	ZendHashDestroy(EG.GetZendConstants())
	Free(EG.GetZendConstants())
	ZendShutdownStrtod()
	if CG.GetMapPtrBase() {
		Free(CG.GetMapPtrBase())
		CG.SetMapPtrBase(nil)
		CG.SetMapPtrSize(0)
	}
	if CG.GetScriptEncodingList() != nil {
		Free(CG.GetScriptEncodingList())
		CG.SetScriptEncodingList(nil)
		CG.SetScriptEncodingListSize(0)
	}
	ZendDestroyRsrcListDtors()
}

/* }}} */

func ZendSetUtilityValues(utility_values *ZendUtilityValues) { ZendUv = *utility_values }

/* }}} */

func Zenderror(error *byte) {
	CG.SetParseError(0)
	if EG.GetException() != nil {

		/* An exception was thrown in the lexer, don't throw another in the parser. */

		return

		/* An exception was thrown in the lexer, don't throw another in the parser. */

	}
	ZendThrowException(ZendCeParseError, error, 0)
}

/* }}} */

func _zendBailout(filename *byte, lineno uint32) {
	if EG.GetBailout() == nil {
		ZendOutputDebugString(1, "%s(%d) : Bailed out without a bailout address!", filename, lineno)
		exit(-1)
	}
	GcProtect(1)
	CG.SetUncleanShutdown(1)
	CG.SetActiveClassEntry(nil)
	CG.SetInCompilation(0)
	EG.SetCurrentExecuteData(nil)
	siglongjmp(EG.bailout, FAILURE)
}

/* }}} */

func ZendAppendVersionInfo(extension *ZendExtension) {
	var new_info *byte
	var new_info_length uint32
	new_info_length = uint32(g.SizeOf("\"    with  v, , by \\n\"") + strlen(extension.GetName()) + strlen(extension.GetVersion()) + strlen(extension.GetCopyright()) + strlen(extension.GetAuthor()))
	new_info = (*byte)(Malloc(new_info_length + 1))
	snprintf(new_info, new_info_length, "    with %s v%s, %s, by %s\n", extension.GetName(), extension.GetVersion(), extension.GetCopyright(), extension.GetAuthor())
	ZendVersionInfo = (*byte)(realloc(ZendVersionInfo, ZendVersionInfoLength+new_info_length+1))
	strncat(ZendVersionInfo, new_info, new_info_length)
	ZendVersionInfoLength += new_info_length
	Free(new_info)
}

/* }}} */

func GetZendVersion() *byte { return ZendVersionInfo }

/* }}} */

func ZendActivate() {
	GcReset()
	InitCompiler()
	InitExecutor()
	StartupScanner()
	if CG.GetMapPtrLast() != 0 {
		memset(CG.GetMapPtrBase(), 0, CG.GetMapPtrLast()*g.SizeOf("void *"))
	}
}

/* }}} */

func ZendCallDestructors() {
	var __orig_bailout *sigjmp_buf = EG.GetBailout()
	var __bailout sigjmp_buf
	EG.SetBailout(&__bailout)
	if sigsetjmp(__bailout, 0) == 0 {
		ShutdownDestructors()
	}
	EG.SetBailout(__orig_bailout)
}

/* }}} */

func ZendDeactivate() {
	/* we're no longer executing anything */

	EG.SetCurrentExecuteData(nil)
	var __orig_bailout *sigjmp_buf = EG.bailout
	var __bailout sigjmp_buf
	EG.SetBailout(&__bailout)
	if sigsetjmp(__bailout, 0) == 0 {
		ShutdownScanner()
	}
	EG.SetBailout(__orig_bailout)

	/* shutdown_executor() takes care of its own bailout handling */

	ShutdownExecutor()
	var __orig_bailout *sigjmp_buf = EG.bailout
	var __bailout sigjmp_buf
	EG.SetBailout(&__bailout)
	if sigsetjmp(__bailout, 0) == 0 {
		ZendIniDeactivate()
	}
	EG.SetBailout(__orig_bailout)
	var __orig_bailout *sigjmp_buf = EG.GetBailout()
	var __bailout sigjmp_buf
	EG.SetBailout(&__bailout)
	if sigsetjmp(__bailout, 0) == 0 {
		ShutdownCompiler()
	}
	EG.SetBailout(__orig_bailout)
	ZendDestroyRsrcList(&EG.regular_list)
}

/* }}} */

func ZendMessageDispatcher(message ZendLong, data any) {
	if ZendMessageDispatcherP != nil {
		ZendMessageDispatcherP(message, data)
	}
}

/* }}} */

func ZendGetConfigurationDirective(name *ZendString) *Zval {
	if ZendGetConfigurationDirectiveP != nil {
		return ZendGetConfigurationDirectiveP(name)
	} else {
		return nil
	}
}

/* }}} */

// #define SAVE_STACK(stack) do { if ( CG ( stack ) . top ) { memcpy ( & stack , & CG ( stack ) , sizeof ( zend_stack ) ) ; CG ( stack ) . top = CG ( stack ) . max = 0 ; CG ( stack ) . elements = NULL ; } else { stack . top = 0 ; } } while ( 0 )

// #define RESTORE_STACK(stack) do { if ( stack . top ) { zend_stack_destroy ( & CG ( stack ) ) ; memcpy ( & CG ( stack ) , & stack , sizeof ( zend_stack ) ) ; } } while ( 0 )

func ZendErrorVaList(type_ int, error_filename *byte, error_lineno uint32, format *byte, args va_list) {
	var usr_copy va_list
	var params []Zval
	var retval Zval
	var orig_user_error_handler Zval
	var in_compilation ZendBool
	var saved_class_entry *ZendClassEntry
	var loop_var_stack ZendStack
	var delayed_oplines_stack ZendStack
	var symbol_table *ZendArray
	var orig_fake_scope *ZendClassEntry

	/* Report about uncaught exception in case of fatal errors */

	if EG.GetException() != nil {
		var ex *ZendExecuteData
		var opline *ZendOp
		switch type_ {
		case 1 << 4:

		case 1 << 0:

		case 1 << 12:

		case 1 << 2:

		case 1 << 6:

		case 1 << 8:
			ex = EG.GetCurrentExecuteData()
			opline = nil
			for ex != nil && (ex.GetFunc() == nil || (ex.GetFunc().GetType()&1) != 0) {
				ex = ex.GetPrevExecuteData()
			}
			if ex != nil && ex.GetOpline().GetOpcode() == 149 && EG.GetOplineBeforeException() != nil {
				opline = EG.GetOplineBeforeException()
			}
			ZendExceptionError(EG.GetException(), 1<<1)
			EG.SetException(nil)
			if opline != nil {
				ex.SetOpline(opline)
			}
			break
		default:
			break
		}
	}

	/* if we don't have a user defined error handler */

	if EG.GetUserErrorHandler().GetType() == 0 || (EG.GetUserErrorHandlerErrorReporting()&type_) == 0 || EG.GetErrorHandling() != EH_NORMAL {
		ZendErrorCb(type_, error_filename, error_lineno, format, args)
	} else {
		switch type_ {
		case 1 << 0:

		case 1 << 2:

		case 1 << 4:

		case 1 << 5:

		case 1 << 6:

		case 1 << 7:

			/* The error may not be safe to handle in user-space */

			ZendErrorCb(type_, error_filename, error_lineno, format, args)
			break
		default:

			/* Handle the error in user space */

			memcpy(&usr_copy, &args, g.SizeOf("va_list"))
			var __z *Zval = &params[1]
			var __s *ZendString = ZendVstrpprintf(0, format, usr_copy)
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
			va_end(usr_copy)
			var __z *zval = &params[0]
			__z.GetValue().SetLval(type_)
			__z.SetTypeInfo(4)
			if error_filename != nil {
				var _s *byte = error_filename
				var __z *Zval = &params[2]
				var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
				__z.GetValue().SetStr(__s)
				__z.SetTypeInfo(6 | 1<<0<<8)
			} else {
				&params[2].SetTypeInfo(1)
			}
			var __z *Zval = &params[3]
			__z.GetValue().SetLval(error_lineno)
			__z.SetTypeInfo(4)
			symbol_table = ZendRebuildSymbolTable()

			/* during shutdown the symbol table table can be still null */

			if symbol_table == nil {
				&params[4].SetTypeInfo(1)
			} else {
				var __arr *ZendArray = ZendArrayDup(symbol_table)
				var __z *Zval = &params[4]
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
			}
			var _z1 *Zval = &orig_user_error_handler
			var _z2 *Zval = &EG.user_error_handler
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			&EG.user_error_handler.u1.type_info = 0

			/* User error handler may include() additinal PHP files.
			 * If an error was generated during comilation PHP will compile
			 * such scripts recursively, but some CG() variables may be
			 * inconsistent. */

			in_compilation = CG.GetInCompilation()
			if in_compilation != 0 {
				saved_class_entry = CG.GetActiveClassEntry()
				CG.SetActiveClassEntry(nil)
				if CG.GetLoopVarStack().GetTop() != 0 {
					memcpy(&loop_var_stack, &CG.loop_var_stack, g.SizeOf("zend_stack"))
					CG.GetLoopVarStack().SetMax(0)
					CG.GetLoopVarStack().SetTop(CG.GetLoopVarStack().GetMax())
					CG.GetLoopVarStack().SetElements(nil)
				} else {
					loop_var_stack.SetTop(0)
				}
				if CG.GetDelayedOplinesStack().GetTop() != 0 {
					memcpy(&delayed_oplines_stack, &CG.delayed_oplines_stack, g.SizeOf("zend_stack"))
					CG.GetDelayedOplinesStack().SetMax(0)
					CG.GetDelayedOplinesStack().SetTop(CG.GetDelayedOplinesStack().GetMax())
					CG.GetDelayedOplinesStack().SetElements(nil)
				} else {
					delayed_oplines_stack.SetTop(0)
				}
				CG.SetInCompilation(0)
			}
			orig_fake_scope = EG.GetFakeScope()
			EG.SetFakeScope(nil)
			if _callUserFunctionEx(nil, &orig_user_error_handler, &retval, 5, params, 1) == SUCCESS {
				if retval.GetType() != 0 {
					if retval.GetType() == 2 {
						ZendErrorCb(type_, error_filename, error_lineno, format, args)
					}
					ZvalPtrDtor(&retval)
				}
			} else if EG.GetException() == nil {

				/* The user error handler failed, use built-in error handler */

				ZendErrorCb(type_, error_filename, error_lineno, format, args)

				/* The user error handler failed, use built-in error handler */

			}
			EG.SetFakeScope(orig_fake_scope)
			if in_compilation != 0 {
				CG.SetActiveClassEntry(saved_class_entry)
				if loop_var_stack.GetTop() != 0 {
					ZendStackDestroy(&CG.loop_var_stack)
					memcpy(&CG.loop_var_stack, &loop_var_stack, g.SizeOf("zend_stack"))
				}
				if delayed_oplines_stack.GetTop() != 0 {
					ZendStackDestroy(&CG.delayed_oplines_stack)
					memcpy(&CG.delayed_oplines_stack, &delayed_oplines_stack, g.SizeOf("zend_stack"))
				}
				CG.SetInCompilation(1)
			}
			ZvalPtrDtor(&params[4])
			ZvalPtrDtor(&params[2])
			ZvalPtrDtor(&params[1])
			if EG.GetUserErrorHandler().GetType() == 0 {
				var _z1 *Zval = &EG.user_error_handler
				var _z2 *Zval = &orig_user_error_handler
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else {
				ZvalPtrDtor(&orig_user_error_handler)
			}
			break
		}
	}
	if type_ == 1<<2 {

		/* eval() errors do not affect exit_status */

		if !(EG.GetCurrentExecuteData() != nil && EG.GetCurrentExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetFunc().GetType()&1) == 0 && EG.GetCurrentExecuteData().GetOpline().GetOpcode() == 73 && EG.GetCurrentExecuteData().GetOpline().GetExtendedValue() == 1<<0) {
			EG.SetExitStatus(255)
		}

		/* eval() errors do not affect exit_status */

	}
}

/* }}} */

func GetFilenameLineno(type_ int, filename **byte, lineno *uint32) {
	/* Obtain relevant filename and lineno */

	switch type_ {
	case 1 << 4:

	case 1 << 5:
		*filename = nil
		*lineno = 0
		break
	case 1 << 2:

	case 1 << 6:

	case 1 << 7:

	case 1 << 0:

	case 1 << 3:

	case 1 << 11:

	case 1 << 13:

	case 1 << 1:

	case 1 << 8:

	case 1 << 9:

	case 1 << 10:

	case 1 << 14:

	case 1 << 12:
		if ZendIsCompiling() != 0 {
			*filename = ZendGetCompiledFilename().GetVal()
			*lineno = ZendGetCompiledLineno()
		} else if ZendIsExecuting() != 0 {
			*filename = ZendGetExecutedFilename()
			if (*filename)[0] == '[' {
				*filename = nil
				*lineno = 0
			} else {
				*lineno = ZendGetExecutedLineno()
			}
		} else {
			*filename = nil
			*lineno = 0
		}
		break
	default:
		*filename = nil
		*lineno = 0
		break
	}
	if (*filename) == nil {
		*filename = "Unknown"
	}
}
func ZendErrorAt(type_ int, filename *byte, lineno uint32, format string, _ ...any) {
	var args va_list
	if filename == nil {
		var dummy_lineno uint32
		GetFilenameLineno(type_, &filename, &dummy_lineno)
	}
	va_start(args, format)
	ZendErrorVaList(type_, filename, lineno, format, args)
	va_end(args)
}
func ZendError(type_ int, format string, _ ...any) {
	var filename *byte
	var lineno uint32
	var args va_list
	GetFilenameLineno(type_, &filename, &lineno)
	va_start(args, format)
	ZendErrorVaList(type_, filename, lineno, format, args)
	va_end(args)
}
func ZendErrorAtNoreturn(type_ int, filename *byte, lineno uint32, format string, _ ...any) {
	var args va_list
	if filename == nil {
		var dummy_lineno uint32
		GetFilenameLineno(type_, &filename, &dummy_lineno)
	}
	va_start(args, format)
	ZendErrorVaList(type_, filename, lineno, format, args)
	va_end(args)

	/* Should never reach this. */

	abort()

	/* Should never reach this. */
}

/* }}} */

func ZendErrorNoreturn(type_ int, format string, _ ...any) {
	var filename *byte
	var lineno uint32
	var args va_list
	GetFilenameLineno(type_, &filename, &lineno)
	va_start(args, format)
	ZendErrorVaList(type_, filename, lineno, format, args)
	va_end(args)

	/* Should never reach this. */

	abort()

	/* Should never reach this. */
}

/* }}} */

func ZendThrowError(exception_ce *ZendClassEntry, format string, _ ...any) {
	var va va_list
	var message *byte = nil
	if exception_ce != nil {
		if InstanceofFunction(exception_ce, ZendCeError) == 0 {
			ZendError(1<<3, "Error exceptions must be derived from Error")
			exception_ce = ZendCeError
		}
	} else {
		exception_ce = ZendCeError
	}

	/* Marker used to disable exception generation during preloading. */

	if EG.GetException() == any(uintptr_t-1) {
		return
	}
	va_start(va, format)
	ZendVspprintf(&message, 0, format, va)

	//TODO: we can't convert compile-time errors to exceptions yet???

	if EG.GetCurrentExecuteData() != nil && CG.GetInCompilation() == 0 {
		ZendThrowException(exception_ce, message, 0)
	} else {
		ZendError(1<<0, "%s", message)
	}
	_efree(message)
	va_end(va)
}

/* }}} */

func ZendTypeError(format string, _ ...any) {
	var va va_list
	var message *byte = nil
	va_start(va, format)
	ZendVspprintf(&message, 0, format, va)
	ZendThrowException(ZendCeTypeError, message, 0)
	_efree(message)
	va_end(va)
}
func ZendInternalTypeError(throw_exception ZendBool, format string, _ ...any) {
	var va va_list
	var message *byte = nil
	va_start(va, format)
	ZendVspprintf(&message, 0, format, va)
	if throw_exception != 0 {
		ZendThrowException(ZendCeTypeError, message, 0)
	} else {
		ZendError(1<<1, "%s", message)
	}
	_efree(message)
	va_end(va)
}
func ZendInternalArgumentCountError(throw_exception ZendBool, format string, _ ...any) {
	var va va_list
	var message *byte = nil
	va_start(va, format)
	ZendVspprintf(&message, 0, format, va)
	if throw_exception != 0 {
		ZendThrowException(ZendCeArgumentCountError, message, 0)
	} else {
		ZendError(1<<1, "%s", message)
	}
	_efree(message)
	va_end(va)
}
func ZendOutputDebugString(trigger_break ZendBool, format string, _ ...any) {}

/* }}} */

func ZendUserExceptionHandler() {
	var orig_user_exception_handler Zval
	var params []Zval
	var retval2 Zval
	var old_exception *ZendObject
	old_exception = EG.GetException()
	EG.SetException(nil)
	var __z *Zval = &params[0]
	__z.GetValue().SetObj(old_exception)
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	var _z1 *Zval = &orig_user_exception_handler
	var _z2 *Zval = &EG.user_exception_handler
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if _callUserFunctionEx(nil, &orig_user_exception_handler, &retval2, 1, params, 1) == SUCCESS {
		ZvalPtrDtor(&retval2)
		if EG.GetException() != nil {
			ZendObjectRelease(EG.GetException())
			EG.SetException(nil)
		}
		ZendObjectRelease(old_exception)
	} else {
		EG.SetException(old_exception)
	}
}
func ZendExecuteScripts(type_ int, retval *Zval, file_count int, _ ...any) int {
	var files va_list
	var i int
	var file_handle *ZendFileHandle
	var op_array *ZendOpArray
	va_start(files, file_count)
	for i = 0; i < file_count; i++ {
		file_handle = __va_arg(files, (*ZendFileHandle)(_))
		if file_handle == nil {
			continue
		}
		op_array = ZendCompileFile(file_handle, type_)
		if file_handle.GetOpenedPath() != nil {
			ZendHashAddEmptyElement(&EG.included_files, file_handle.GetOpenedPath())
		}
		ZendDestroyFileHandle(file_handle)
		if op_array != nil {
			ZendExecute(op_array, retval)
			ZendExceptionRestore()
			if EG.GetException() != nil {
				if EG.GetUserExceptionHandler().GetType() != 0 {
					ZendUserExceptionHandler()
				}
				if EG.GetException() != nil {
					ZendExceptionError(EG.GetException(), 1<<0)
				}
			}
			DestroyOpArray(op_array)
			_efree(op_array)
		} else if type_ == 1<<3 {
			va_end(files)
			return FAILURE
		}
	}
	va_end(files)
	return SUCCESS
}

/* }}} */

// #define COMPILED_STRING_DESCRIPTION_FORMAT       "%s(%d) : %s"

func ZendMakeCompiledStringDescription(name string) *byte {
	var cur_filename *byte
	var cur_lineno int
	var compiled_string_description *byte
	if ZendIsCompiling() != 0 {
		cur_filename = ZendGetCompiledFilename().GetVal()
		cur_lineno = ZendGetCompiledLineno()
	} else if ZendIsExecuting() != 0 {
		cur_filename = ZendGetExecutedFilename()
		cur_lineno = ZendGetExecutedLineno()
	} else {
		cur_filename = "Unknown"
		cur_lineno = 0
	}
	ZendSpprintf(&compiled_string_description, 0, "%s(%d) : %s", cur_filename, cur_lineno, name)
	return compiled_string_description
}

/* }}} */

func FreeEstring(str_p **byte) { _efree(*str_p) }

/* }}} */

func ZendMapPtrReset() { CG.SetMapPtrLast(GlobalMapPtrLast) }
func ZendMapPtrNew() any {
	var ptr *any
	if CG.GetMapPtrLast() >= CG.GetMapPtrSize() {

		/* Grow map_ptr table */

		CG.SetMapPtrSize(CG.GetMapPtrLast() + 1 + (4096-1) & ^(4096-1))
		CG.SetMapPtrBase(__zendRealloc(CG.GetMapPtrBase(), CG.GetMapPtrSize()*g.SizeOf("void *")))
	}
	ptr = (*any)(CG.GetMapPtrBase() + CG.GetMapPtrLast())
	*ptr = nil
	CG.GetMapPtrLast()++
	return any(uintptr_t((*byte)(ptr)-(*byte)(CG.GetMapPtrBase())) | 1)
}
func ZendMapPtrExtend(last int) {
	if last > CG.GetMapPtrLast() {
		var ptr *any
		if last >= CG.GetMapPtrSize() {

			/* Grow map_ptr table */

			CG.SetMapPtrSize(last + (4096-1) & ^(4096-1))
			CG.SetMapPtrBase(__zendRealloc(CG.GetMapPtrBase(), CG.GetMapPtrSize()*g.SizeOf("void *")))
		}
		ptr = (*any)(CG.GetMapPtrBase() + CG.GetMapPtrLast())
		memset(ptr, 0, (last-CG.GetMapPtrLast())*g.SizeOf("void *"))
		CG.SetMapPtrLast(last)
	}
}
