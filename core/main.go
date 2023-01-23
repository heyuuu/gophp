// <<generate>>

package core

import (
	"sik/core/streams"
	"sik/ext/standard"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/main.c>

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
   |          Rasmus Lerdorf <rasmus@lerdorf.on.ca>                       |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// #define ZEND_INCLUDE_FULL_WINDOWS_HEADERS

// # include "php.h"

// # include < stdio . h >

// # include < fcntl . h >

// # include < sys / time . h >

// # include < unistd . h >

// # include < signal . h >

// # include < locale . h >

// # include "zend.h"

// # include "zend_types.h"

// # include "zend_extensions.h"

// # include "php_ini.h"

// # include "php_globals.h"

// # include "php_main.h"

// # include "php_syslog.h"

// # include "fopen_wrappers.h"

// # include "ext/standard/php_standard.h"

// # include "ext/standard/php_string.h"

// failed # include "ext/date/php_date.h"

// # include "php_variables.h"

// # include "ext/standard/credits.h"

// # include "php_syslog.h"

// # include "Zend/zend_exceptions.h"

// # include "zend_compile.h"

// # include "zend_execute.h"

// # include "zend_highlight.h"

// # include "zend_extensions.h"

// # include "zend_ini.h"

// # include "zend_dtrace.h"

// # include "php_content_types.h"

// # include "php_ticks.h"

// # include "php_streams.h"

// # include "php_open_temporary_file.h"

// # include "SAPI.h"

// # include "rfc1867.h"

// # include "ext/standard/html_tables.h"

/* }}} */

var PhpRegisterInternalExtensionsFunc func() int = PhpRegisterInternalExtensions
var CoreGlobals PhpCoreGlobals

// #define SAFE_FILENAME(f) ( ( f ) ? ( f ) : "-" )

func GetSafeCharsetHint() *byte {
	var lastHint *byte = nil
	var lastCodeset *byte = nil
	var hint *byte = sapi_globals.GetDefaultCharset()
	var len_ int = strlen(hint)
	var i int = 0
	if lastHint == sapi_globals.GetDefaultCharset() {
		return lastCodeset
	}
	lastHint = hint
	lastCodeset = nil
	for i = 0; i < g.SizeOf("charset_map")/g.SizeOf("charset_map [ 0 ]"); i++ {
		if len_ == standard.CharsetMap[i].codeset_len && zend.ZendBinaryStrcasecmp(hint, len_, standard.CharsetMap[i].codeset, len_) == 0 {
			lastCodeset = (*byte)(standard.CharsetMap[i].codeset)
			break
		}
	}
	return lastCodeset
}

/* {{{ PHP_INI_MH
 */

func OnSetFacility(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var facility *byte = new_value.val
	return zend.FAILURE
}

/* }}} */

func OnSetPrecision(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var i zend.ZendLong
	i = atoll(new_value.val)
	if i >= -1 {
		zend.EG.precision = i
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}

/* }}} */

func OnSetSerializePrecision(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var i zend.ZendLong
	i = atoll(new_value.val)
	if i >= -1 {
		CoreGlobals.SetSerializePrecision(i)
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}

/* }}} */

func OnChangeMemoryLimit(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var value int
	if new_value != nil {
		value = zend.ZendAtol(new_value.val, new_value.len_)
	} else {
		value = 1 << 30
	}
	if zend.ZendSetMemoryLimit(value) == zend.FAILURE {

		/* When the memory limit is reset to the original level during deactivation, we may be
		 * using more memory than the original limit while shutdown is still in progress.
		 * Ignore a failure for now, and set the memory limit when the memory manager has been
		 * shut down and the minimal amount of memory is used. */

		if stage != 1<<3 {
			zend.ZendError(1<<1, "Failed to set memory limit to %zd bytes (Current memory usage is %zd bytes)", value, zend.ZendMemoryUsage(true))
			return zend.FAILURE
		}

		/* When the memory limit is reset to the original level during deactivation, we may be
		 * using more memory than the original limit while shutdown is still in progress.
		 * Ignore a failure for now, and set the memory limit when the memory manager has been
		 * shut down and the minimal amount of memory is used. */

	}
	CoreGlobals.SetMemoryLimit(value)
	return zend.SUCCESS
}

/* }}} */

func OnSetLogFilter(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var filter *byte = new_value.val
	if !(strcmp(filter, "all")) {
		CoreGlobals.SetSyslogFilter(0)
		return zend.SUCCESS
	}
	if !(strcmp(filter, "no-ctrl")) {
		CoreGlobals.SetSyslogFilter(1)
		return zend.SUCCESS
	}
	if !(strcmp(filter, "ascii")) {
		CoreGlobals.SetSyslogFilter(2)
		return zend.SUCCESS
	}
	if !(strcmp(filter, "raw")) {
		CoreGlobals.SetSyslogFilter(3)
		return zend.SUCCESS
	}
	return zend.FAILURE
}

/* }}} */

func PhpDisableFunctions() {
	var s *byte = nil
	var e *byte
	if !(*(zend.ZendIniStringEx("disable_functions", g.SizeOf("\"disable_functions\"")-1, 0, nil))) {
		return
	}
	CoreGlobals.SetDisableFunctions(strdup(zend.ZendIniStringEx("disable_functions", g.SizeOf("\"disable_functions\"")-1, 0, nil)))
	e = CoreGlobals.GetDisableFunctions()
	if e == nil {
		return
	}
	for *e {
		switch *e {
		case ' ':

		case ',':
			if s != nil {
				*e = '0'
				zend.ZendDisableFunction(s, e-s)
				s = nil
			}
			break
		default:
			if s == nil {
				s = e
			}
			break
		}
		e++
	}
	if s != nil {
		zend.ZendDisableFunction(s, e-s)
	}
}

/* }}} */

func PhpDisableClasses() {
	var s *byte = nil
	var e *byte
	if !(*(zend.ZendIniStringEx("disable_classes", g.SizeOf("\"disable_classes\"")-1, 0, nil))) {
		return
	}
	CoreGlobals.SetDisableClasses(strdup(zend.ZendIniStringEx("disable_classes", g.SizeOf("\"disable_classes\"")-1, 0, nil)))
	e = CoreGlobals.GetDisableClasses()
	for *e {
		switch *e {
		case ' ':

		case ',':
			if s != nil {
				*e = '0'
				zend.ZendDisableClass(s, e-s)
				s = nil
			}
			break
		default:
			if s == nil {
				s = e
			}
			break
		}
		e++
	}
	if s != nil {
		zend.ZendDisableClass(s, e-s)
	}
}

/* }}} */

func PhpBinaryInit() {
	var binary_location *byte = nil
	if sapi_module.GetExecutableLocation() != nil {
		binary_location = (*byte)(zend.Malloc(256))
		if binary_location != nil && !(strchr(sapi_module.GetExecutableLocation(), '/')) {
			var envpath *byte
			var path *byte
			var found int = 0
			if g.Assign(&envpath, getenv("PATH")) != nil {
				var search_dir *byte
				var search_path []*byte
				var last *byte = nil
				var s zend.ZendStatT
				path = zend._estrdup(envpath)
				search_dir = strtok_r(path, ":", &last)
				for search_dir != nil {
					ApPhpSnprintf(search_path, 256, "%s/%s", search_dir, sapi_module.GetExecutableLocation())
					if zend.TsrmRealpath(search_path, binary_location) != nil && !(access(binary_location, X_OK)) && stat(binary_location, &s) == 0 && (s.st_mode&S_IFMT) == S_IFREG {
						found = 1
						break
					}
					search_dir = strtok_r(nil, ":", &last)
				}
				zend._efree(path)
			}
			if found == 0 {
				zend.Free(binary_location)
				binary_location = nil
			}
		} else if zend.TsrmRealpath(sapi_module.GetExecutableLocation(), binary_location) == nil || access(binary_location, X_OK) {
			zend.Free(binary_location)
			binary_location = nil
		}
	}
	CoreGlobals.SetPhpBinary(binary_location)
}

/* }}} */

func OnUpdateTimeout(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	if stage == 1<<0 {

		/* Don't set a timeout on startup, only per-request */

		zend.EG.timeout_seconds = atoll(new_value.val)
		return zend.SUCCESS
	}
	zend.ZendUnsetTimeout()
	zend.EG.timeout_seconds = atoll(new_value.val)
	if stage != 1<<3 {

		/*
		 * If we're restoring INI values, we shouldn't reset the timer.
		 * Otherwise, the timer is active when PHP is idle, such as the
		 * the CLI web server or CGI. Running a script will re-activate
		 * the timeout, so it's not needed to do so at script end.
		 */

		zend.ZendSetTimeout(zend.EG.timeout_seconds, 0)

		/*
		 * If we're restoring INI values, we shouldn't reset the timer.
		 * Otherwise, the timer is active when PHP is idle, such as the
		 * the CLI web server or CGI. Running a script will re-activate
		 * the timeout, so it's not needed to do so at script end.
		 */

	}
	return zend.SUCCESS
}

/* }}} */

func PhpGetDisplayErrorsMode(value *byte, value_length int) int {
	var mode int
	if value == nil {
		return 1
	}
	if value_length == 2 && !(strcasecmp("on", value)) {
		mode = 1
	} else if value_length == 3 && !(strcasecmp("yes", value)) {
		mode = 1
	} else if value_length == 4 && !(strcasecmp("true", value)) {
		mode = 1
	} else if value_length == 6 && !(strcasecmp(value, "stderr")) {
		mode = 2
	} else if value_length == 6 && !(strcasecmp(value, "stdout")) {
		mode = 1
	} else {
		mode = atoll(value)
		if mode != 0 && mode != 1 && mode != 2 {
			mode = 1
		}
	}
	return mode
}

/* }}} */

func OnUpdateDisplayErrors(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	CoreGlobals.SetDisplayErrors(zend.ZendBool(PhpGetDisplayErrorsMode(new_value.val, new_value.len_)))
	return zend.SUCCESS
}

/* }}} */

func DisplayErrorsMode(ini_entry *zend.ZendIniEntry, type_ int) {
	var mode int
	var cgi_or_cli int
	var tmp_value_length int
	var tmp_value *byte
	if type_ == 1 && ini_entry.modified != 0 {
		if ini_entry.orig_value != nil {
			tmp_value = ini_entry.orig_value.val
		} else {
			tmp_value = nil
		}
		if ini_entry.orig_value != nil {
			tmp_value_length = ini_entry.orig_value.len_
		} else {
			tmp_value_length = 0
		}
	} else if ini_entry.value != nil {
		tmp_value = ini_entry.value.val
		tmp_value_length = ini_entry.value.len_
	} else {
		tmp_value = nil
		tmp_value_length = 0
	}
	mode = PhpGetDisplayErrorsMode(tmp_value, tmp_value_length)

	/* Display 'On' for other SAPIs instead of STDOUT or STDERR */

	cgi_or_cli = !(strcmp(sapi_module.GetName(), "cli")) || !(strcmp(sapi_module.GetName(), "cgi")) || !(strcmp(sapi_module.GetName(), "phpdbg"))
	switch mode {
	case 2:
		if cgi_or_cli != 0 {
			var __str *byte = "STDERR"
			PhpOutputWrite(__str, strlen(__str))
		} else {
			var __str *byte = "On"
			PhpOutputWrite(__str, strlen(__str))
		}
		break
	case 1:
		if cgi_or_cli != 0 {
			var __str *byte = "STDOUT"
			PhpOutputWrite(__str, strlen(__str))
		} else {
			var __str *byte = "On"
			PhpOutputWrite(__str, strlen(__str))
		}
		break
	default:
		var __str *byte = "Off"
		PhpOutputWrite(__str, strlen(__str))
		break
	}
}

/* }}} */

func PhpGetInternalEncoding() *byte {
	if CoreGlobals.GetInternalEncoding() != nil && CoreGlobals.GetInternalEncoding()[0] {
		return CoreGlobals.GetInternalEncoding()
	} else if sapi_globals.GetDefaultCharset() != nil {
		return sapi_globals.GetDefaultCharset()
	}
	return ""
}
func PhpGetInputEncoding() *byte {
	if CoreGlobals.GetInputEncoding() != nil && CoreGlobals.GetInputEncoding()[0] {
		return CoreGlobals.GetInputEncoding()
	} else if sapi_globals.GetDefaultCharset() != nil {
		return sapi_globals.GetDefaultCharset()
	}
	return ""
}
func PhpGetOutputEncoding() *byte {
	if CoreGlobals.GetOutputEncoding() != nil && CoreGlobals.GetOutputEncoding()[0] {
		return CoreGlobals.GetOutputEncoding()
	} else if sapi_globals.GetDefaultCharset() != nil {
		return sapi_globals.GetDefaultCharset()
	}
	return ""
}

var PhpInternalEncodingChanged func() = nil

/* {{{ PHP_INI_MH
 */

func OnUpdateDefaultCharset(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	if memchr(new_value.val, '0', new_value.len_) || strpbrk(new_value.val, "\r\n") {
		return zend.FAILURE
	}
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	if PhpInternalEncodingChanged != nil {
		PhpInternalEncodingChanged()
	}
	if new_value != nil {

	}
	return zend.SUCCESS
}

/* }}} */

func OnUpdateDefaultMimeTye(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	if memchr(new_value.val, '0', new_value.len_) || strpbrk(new_value.val, "\r\n") {
		return zend.FAILURE
	}
	return zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
}

/* }}} */

func OnUpdateInternalEncoding(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	if PhpInternalEncodingChanged != nil {
		PhpInternalEncodingChanged()
	}
	if new_value != nil {

	}
	return zend.SUCCESS
}

/* }}} */

func OnUpdateInputEncoding(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	if PhpInternalEncodingChanged != nil {
		PhpInternalEncodingChanged()
	}
	if new_value != nil {

	}
	return zend.SUCCESS
}

/* }}} */

func OnUpdateOutputEncoding(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	if PhpInternalEncodingChanged != nil {
		PhpInternalEncodingChanged()
	}
	if new_value != nil {

	}
	return zend.SUCCESS
}

/* }}} */

func OnUpdateErrorLog(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	/* Only do the safemode/open_basedir check at runtime */

	if (stage == 1<<4 || stage == 1<<5) && new_value != nil && strcmp(new_value.val, "syslog") {
		if CoreGlobals.GetOpenBasedir() != nil && PhpCheckOpenBasedir(new_value.val) != 0 {
			return zend.FAILURE
		}
	}
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	return zend.SUCCESS
}

/* }}} */

func OnUpdateMailLog(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	/* Only do the safemode/open_basedir check at runtime */

	if (stage == 1<<4 || stage == 1<<5) && new_value != nil {
		if CoreGlobals.GetOpenBasedir() != nil && PhpCheckOpenBasedir(new_value.val) != 0 {
			return zend.FAILURE
		}
	}
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	return zend.SUCCESS
}

/* }}} */

func OnChangeMailForceExtra(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	/* Don't allow changing it in htaccess */

	if stage == 1<<5 {
		return zend.FAILURE
	}
	return zend.SUCCESS
}

/* }}} */

var OnChangeBrowscap func(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int

/* Need to be read from the environment (?):
 * PHP_AUTO_PREPEND_FILE
 * PHP_AUTO_APPEND_FILE
 * PHP_DOCUMENT_ROOT
 * PHP_USER_DIR
 * PHP_INCLUDE_PATH
 */

// #define DEFAULT_SENDMAIL_PATH       PHP_PROG_SENDMAIL " -t -i"

/* {{{ PHP_INI
 */

var IniEntries []zend.ZendIniEntryDef = []zend.ZendIniEntryDef{
	{"highlight.comment", nil, nil, nil, nil, "#FF8000", zend.ZendIniColorDisplayerCb, g.SizeOf("\"#FF8000\"") - 1, g.SizeOf("\"highlight.comment\"") - 1, 1<<0 | 1<<1 | 1<<2},
	{"highlight.default", nil, nil, nil, nil, "#0000BB", zend.ZendIniColorDisplayerCb, g.SizeOf("\"#0000BB\"") - 1, g.SizeOf("\"highlight.default\"") - 1, 1<<0 | 1<<1 | 1<<2},
	{"highlight.html", nil, nil, nil, nil, "#000000", zend.ZendIniColorDisplayerCb, g.SizeOf("\"#000000\"") - 1, g.SizeOf("\"highlight.html\"") - 1, 1<<0 | 1<<1 | 1<<2},
	{"highlight.keyword", nil, nil, nil, nil, "#007700", zend.ZendIniColorDisplayerCb, g.SizeOf("\"#007700\"") - 1, g.SizeOf("\"highlight.keyword\"") - 1, 1<<0 | 1<<1 | 1<<2},
	{"highlight.string", nil, nil, nil, nil, "#DD0000", zend.ZendIniColorDisplayerCb, g.SizeOf("\"#DD0000\"") - 1, g.SizeOf("\"highlight.string\"") - 1, 1<<0 | 1<<1 | 1<<2},
	{
		"display_errors",
		OnUpdateDisplayErrors,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDisplayErrors())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"1",
		DisplayErrorsMode,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"display_errors\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"display_startup_errors",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDisplayStartupErrors())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"0",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"display_startup_errors\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"enable_dl",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetEnableDl())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"1",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"enable_dl\"") - 1,
		1 << 2,
	},
	{
		"expose_php",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetExposePhp())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"1",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"expose_php\"") - 1,
		1 << 2,
	},
	{
		"docref_root",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDocrefRoot())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"",
		nil,
		g.SizeOf("\"\"") - 1,
		g.SizeOf("\"docref_root\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"docref_ext",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDocrefExt())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"",
		nil,
		g.SizeOf("\"\"") - 1,
		g.SizeOf("\"docref_ext\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"html_errors",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetHtmlErrors())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"1",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"html_errors\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"xmlrpc_errors",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetXmlrpcErrors())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"0",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"xmlrpc_errors\"") - 1,
		1 << 2,
	},
	{
		"xmlrpc_error_number",
		zend.OnUpdateLong,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetXmlrpcErrorNumber())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"0",
		nil,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"xmlrpc_error_number\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"max_input_time",
		zend.OnUpdateLong,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMaxInputTime())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"-1",
		nil,
		g.SizeOf("\"-1\"") - 1,
		g.SizeOf("\"max_input_time\"") - 1,
		1<<2 | 1<<1,
	},
	{
		"ignore_user_abort",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetIgnoreUserAbort())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"0",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"ignore_user_abort\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"implicit_flush",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetImplicitFlush())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"0",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"implicit_flush\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"log_errors",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetLogErrors())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"0",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"log_errors\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"log_errors_max_len",
		zend.OnUpdateLong,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetLogErrorsMaxLen())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"1024",
		nil,
		g.SizeOf("\"1024\"") - 1,
		g.SizeOf("\"log_errors_max_len\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"ignore_repeated_errors",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetIgnoreRepeatedErrors())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"0",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"ignore_repeated_errors\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"ignore_repeated_source",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetIgnoreRepeatedSource())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"0",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"ignore_repeated_source\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"report_memleaks",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetReportMemleaks())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"1",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"report_memleaks\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"report_zend_debug",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetReportZendDebug())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"1",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"report_zend_debug\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"output_buffering",
		zend.OnUpdateLong,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetOutputBuffering())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"0",
		nil,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"output_buffering\"") - 1,
		1<<1 | 1<<2,
	},
	{
		"output_handler",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetOutputHandler())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"output_handler\"") - 1,
		1<<1 | 1<<2,
	},
	{
		"register_argc_argv",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetRegisterArgcArgv())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"1",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"register_argc_argv\"") - 1,
		1<<1 | 1<<2,
	},
	{
		"auto_globals_jit",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAutoGlobalsJit())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"1",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"auto_globals_jit\"") - 1,
		1<<1 | 1<<2,
	},
	{
		"short_open_tag",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*zend.ZendCompilerGlobals)(nil).short_tags)) - (*byte)(nil))),
		any(&zend.CompilerGlobals),
		nil,
		"1",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"short_open_tag\"") - 1,
		1<<2 | 1<<1,
	},
	{
		"track_errors",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetTrackErrors())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"0",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"track_errors\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"unserialize_callback_func",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUnserializeCallbackFunc())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"unserialize_callback_func\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"serialize_precision",
		OnSetSerializePrecision,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSerializePrecision())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"-1",
		nil,
		g.SizeOf("\"-1\"") - 1,
		g.SizeOf("\"serialize_precision\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"arg_separator.output",
		zend.OnUpdateStringUnempty,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetArgSeparator().GetOutput())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"&",
		nil,
		g.SizeOf("\"&\"") - 1,
		g.SizeOf("\"arg_separator.output\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"arg_separator.input",
		zend.OnUpdateStringUnempty,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetArgSeparator().GetInput())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"&",
		nil,
		g.SizeOf("\"&\"") - 1,
		g.SizeOf("\"arg_separator.input\"") - 1,
		1<<2 | 1<<1,
	},
	{
		"auto_append_file",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAutoAppendFile())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"auto_append_file\"") - 1,
		1<<2 | 1<<1,
	},
	{
		"auto_prepend_file",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAutoPrependFile())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"auto_prepend_file\"") - 1,
		1<<2 | 1<<1,
	},
	{
		"doc_root",
		zend.OnUpdateStringUnempty,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDocRoot())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"doc_root\"") - 1,
		1 << 2,
	},
	{
		"default_charset",
		OnUpdateDefaultCharset,
		any(zend_long((*byte)(&((*sapi_globals_struct)(nil).GetDefaultCharset())) - (*byte)(nil))),
		any(&sapi_globals),
		nil,
		"UTF-8",
		nil,
		g.SizeOf("\"UTF-8\"") - 1,
		g.SizeOf("\"default_charset\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"default_mimetype",
		OnUpdateDefaultMimeTye,
		any(zend_long((*byte)(&((*sapi_globals_struct)(nil).GetDefaultMimetype())) - (*byte)(nil))),
		any(&sapi_globals),
		nil,
		"text/html",
		nil,
		g.SizeOf("\"text/html\"") - 1,
		g.SizeOf("\"default_mimetype\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"internal_encoding",
		OnUpdateInternalEncoding,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetInternalEncoding())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"internal_encoding\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"input_encoding",
		OnUpdateInputEncoding,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetInputEncoding())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"input_encoding\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"output_encoding",
		OnUpdateOutputEncoding,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetOutputEncoding())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"output_encoding\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"error_log",
		OnUpdateErrorLog,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetErrorLog())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"error_log\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"extension_dir",
		zend.OnUpdateStringUnempty,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetExtensionDir())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"/usr/local/lib/php/extensions/no-debug-non-zts-20190902",
		nil,
		g.SizeOf("\"/usr/local/lib/php/extensions/no-debug-non-zts-20190902\"") - 1,
		g.SizeOf("\"extension_dir\"") - 1,
		1 << 2,
	},
	{
		"sys_temp_dir",
		zend.OnUpdateStringUnempty,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSysTempDir())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"sys_temp_dir\"") - 1,
		1 << 2,
	},
	{
		"include_path",
		zend.OnUpdateStringUnempty,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetIncludePath())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		".:",
		nil,
		g.SizeOf("\".:\"") - 1,
		g.SizeOf("\"include_path\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{"max_execution_time", OnUpdateTimeout, nil, nil, nil, "30", nil, g.SizeOf("\"30\"") - 1, g.SizeOf("\"max_execution_time\"") - 1, 1<<0 | 1<<1 | 1<<2},
	{
		"open_basedir",
		OnUpdateBaseDir,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetOpenBasedir())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"open_basedir\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"file_uploads",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetFileUploads())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"1",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"file_uploads\"") - 1,
		1 << 2,
	},
	{
		"upload_max_filesize",
		zend.OnUpdateLong,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUploadMaxFilesize())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"2M",
		nil,
		g.SizeOf("\"2M\"") - 1,
		g.SizeOf("\"upload_max_filesize\"") - 1,
		1<<2 | 1<<1,
	},
	{
		"post_max_size",
		zend.OnUpdateLong,
		any(zend_long((*byte)(&((*sapi_globals_struct)(nil).GetPostMaxSize())) - (*byte)(nil))),
		any(&sapi_globals),
		nil,
		"8M",
		nil,
		g.SizeOf("\"8M\"") - 1,
		g.SizeOf("\"post_max_size\"") - 1,
		1<<2 | 1<<1,
	},
	{
		"upload_tmp_dir",
		zend.OnUpdateStringUnempty,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUploadTmpDir())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"upload_tmp_dir\"") - 1,
		1 << 2,
	},
	{
		"max_input_nesting_level",
		zend.OnUpdateLongGEZero,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMaxInputNestingLevel())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"64",
		nil,
		g.SizeOf("\"64\"") - 1,
		g.SizeOf("\"max_input_nesting_level\"") - 1,
		1<<2 | 1<<1,
	},
	{
		"max_input_vars",
		zend.OnUpdateLongGEZero,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMaxInputVars())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"1000",
		nil,
		g.SizeOf("\"1000\"") - 1,
		g.SizeOf("\"max_input_vars\"") - 1,
		1<<2 | 1<<1,
	},
	{
		"user_dir",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUserDir())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"user_dir\"") - 1,
		1 << 2,
	},
	{
		"variables_order",
		zend.OnUpdateStringUnempty,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetVariablesOrder())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"EGPCS",
		nil,
		g.SizeOf("\"EGPCS\"") - 1,
		g.SizeOf("\"variables_order\"") - 1,
		1<<2 | 1<<1,
	},
	{
		"request_order",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetRequestOrder())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"request_order\"") - 1,
		1<<2 | 1<<1,
	},
	{
		"error_append_string",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetErrorAppendString())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"error_append_string\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"error_prepend_string",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetErrorPrependString())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"error_prepend_string\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{"SMTP", nil, nil, nil, nil, "localhost", nil, g.SizeOf("\"localhost\"") - 1, g.SizeOf("\"SMTP\"") - 1, 1<<0 | 1<<1 | 1<<2},
	{"smtp_port", nil, nil, nil, nil, "25", nil, g.SizeOf("\"25\"") - 1, g.SizeOf("\"smtp_port\"") - 1, 1<<0 | 1<<1 | 1<<2},
	{
		"mail.add_x_header",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMailXHeader())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"0",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"mail.add_x_header\"") - 1,
		1<<2 | 1<<1,
	},
	{
		"mail.log",
		OnUpdateMailLog,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMailLog())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"mail.log\"") - 1,
		1<<2 | 1<<1,
	},
	{"browscap", OnChangeBrowscap, nil, nil, nil, nil, nil, g.SizeOf("NULL") - 1, g.SizeOf("\"browscap\"") - 1, 1 << 2},
	{"memory_limit", OnChangeMemoryLimit, nil, nil, nil, "128M", nil, g.SizeOf("\"128M\"") - 1, g.SizeOf("\"memory_limit\"") - 1, 1<<0 | 1<<1 | 1<<2},
	{"precision", OnSetPrecision, nil, nil, nil, "14", nil, g.SizeOf("\"14\"") - 1, g.SizeOf("\"precision\"") - 1, 1<<0 | 1<<1 | 1<<2},
	{"sendmail_from", nil, nil, nil, nil, nil, nil, g.SizeOf("NULL") - 1, g.SizeOf("\"sendmail_from\"") - 1, 1<<0 | 1<<1 | 1<<2},
	{
		"sendmail_path",
		nil,
		nil,
		nil,
		nil,
		"/usr/sbin/sendmail" + " -t -i",
		nil,
		g.SizeOf("\"/usr/sbin/sendmail\" \" -t -i\"") - 1,
		g.SizeOf("\"sendmail_path\"") - 1,
		1 << 2,
	},
	{
		"mail.force_extra_parameters",
		OnChangeMailForceExtra,
		nil,
		nil,
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"mail.force_extra_parameters\"") - 1,
		1<<2 | 1<<1,
	},
	{"disable_functions", nil, nil, nil, nil, "", nil, g.SizeOf("\"\"") - 1, g.SizeOf("\"disable_functions\"") - 1, 1 << 2},
	{"disable_classes", nil, nil, nil, nil, "", nil, g.SizeOf("\"\"") - 1, g.SizeOf("\"disable_classes\"") - 1, 1 << 2},
	{"max_file_uploads", nil, nil, nil, nil, "20", nil, g.SizeOf("\"20\"") - 1, g.SizeOf("\"max_file_uploads\"") - 1, 1<<2 | 1<<1},
	{
		"allow_url_fopen",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAllowUrlFopen())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"1",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"allow_url_fopen\"") - 1,
		1 << 2,
	},
	{
		"allow_url_include",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAllowUrlInclude())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"0",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"allow_url_include\"") - 1,
		1 << 2,
	},
	{
		"enable_post_data_reading",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetEnablePostDataReading())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"1",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"enable_post_data_reading\"") - 1,
		1<<2 | 1<<1,
	},
	{
		"realpath_cache_size",
		zend.OnUpdateLong,
		any(zend_long((*byte)(&((*zend.VirtualCwdGlobals)(nil).realpath_cache_size_limit)) - (*byte)(nil))),
		any(&zend.CwdGlobals),
		nil,
		"4096K",
		nil,
		g.SizeOf("\"4096K\"") - 1,
		g.SizeOf("\"realpath_cache_size\"") - 1,
		1 << 2,
	},
	{
		"realpath_cache_ttl",
		zend.OnUpdateLong,
		any(zend_long((*byte)(&((*zend.VirtualCwdGlobals)(nil).realpath_cache_ttl)) - (*byte)(nil))),
		any(&zend.CwdGlobals),
		nil,
		"120",
		nil,
		g.SizeOf("\"120\"") - 1,
		g.SizeOf("\"realpath_cache_ttl\"") - 1,
		1 << 2,
	},
	{
		"user_ini.filename",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUserIniFilename())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		".user.ini",
		nil,
		g.SizeOf("\".user.ini\"") - 1,
		g.SizeOf("\"user_ini.filename\"") - 1,
		1 << 2,
	},
	{
		"user_ini.cache_ttl",
		zend.OnUpdateLong,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUserIniCacheTtl())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"300",
		nil,
		g.SizeOf("\"300\"") - 1,
		g.SizeOf("\"user_ini.cache_ttl\"") - 1,
		1 << 2,
	},
	{
		"hard_timeout",
		zend.OnUpdateLong,
		any(zend_long((*byte)(&((*zend.ZendExecutorGlobals)(nil).hard_timeout)) - (*byte)(nil))),
		any(&zend.ExecutorGlobals),
		nil,
		"2",
		nil,
		g.SizeOf("\"2\"") - 1,
		g.SizeOf("\"hard_timeout\"") - 1,
		1 << 2,
	},
	{
		"syslog.facility",
		OnSetFacility,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSyslogFacility())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"LOG_USER",
		nil,
		g.SizeOf("\"LOG_USER\"") - 1,
		g.SizeOf("\"syslog.facility\"") - 1,
		1 << 2,
	},
	{
		"syslog.ident",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSyslogIdent())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"php",
		nil,
		g.SizeOf("\"php\"") - 1,
		g.SizeOf("\"syslog.ident\"") - 1,
		1 << 2,
	},
	{
		"syslog.filter",
		OnSetLogFilter,
		any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSyslogFilter())) - (*byte)(nil))),
		any(&CoreGlobals),
		nil,
		"no-ctrl",
		nil,
		g.SizeOf("\"no-ctrl\"") - 1,
		g.SizeOf("\"syslog.filter\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{nil, nil, nil, nil, nil, nil, nil, 0, 0, 0},
}

/* }}} */

var ModuleInitialized int = 0
var ModuleStartup int = 1
var ModuleShutdown int = 0

/* {{{ php_during_module_startup */

func PhpDuringModuleStartup() int { return ModuleStartup }

/* }}} */

func PhpDuringModuleShutdown() int { return ModuleShutdown }

/* }}} */

func PhpGetModuleInitialized() int { return ModuleInitialized }

/* }}} */

func PhpLogErrWithSeverity(log_message *byte, syslog_type_int int) {
	var fd int = -1
	var error_time int64
	if CoreGlobals.GetInErrorLog() != 0 {

		/* prevent recursive invocation */

		return

		/* prevent recursive invocation */

	}
	CoreGlobals.SetInErrorLog(1)

	/* Try to use the specified logging location. */

	if CoreGlobals.GetErrorLog() != nil {
		if !(strcmp(CoreGlobals.GetErrorLog(), "syslog")) {
			PhpSyslog(syslog_type_int, "%s", log_message)
			CoreGlobals.SetInErrorLog(0)
			return
		}
		fd = open(CoreGlobals.GetErrorLog(), O_CREAT|O_APPEND|O_WRONLY, 0644)
		if fd != -1 {
			var tmp *byte
			var len_ int
			var error_time_str *zend.ZendString
			time(&error_time)
			error_time_str = php_format_date("d-M-Y H:i:s e", 13, error_time, 1)
			len_ = zend.ZendSpprintf(&tmp, 0, "[%s] %s%s", error_time_str.val, log_message, "\n")
			void(write(fd, tmp, len_))
			zend._efree(tmp)
			zend.ZendStringFree(error_time_str)
			close(fd)
			CoreGlobals.SetInErrorLog(0)
			return
		}
	}

	/* Otherwise fall back to the default logging location, if we have one */

	if sapi_module.GetLogMessage() != nil {
		sapi_module.GetLogMessage()(log_message, syslog_type_int)
	}
	CoreGlobals.SetInErrorLog(0)
}

/* }}} */

func PhpWrite(buf any, size int) int { return PhpOutputWrite(buf, size) }

/* }}} */

func PhpPrintf(format string, _ ...any) int {
	var args va_list
	var ret int
	var buffer *byte
	var size int
	va_start(args, format)
	size = zend.ZendVspprintf(&buffer, 0, format, args)
	ret = PhpOutputWrite(buffer, size)
	zend._efree(buffer)
	va_end(args)
	return ret
}

/* }}} */

func PhpVerror(docref *byte, params *byte, type_ int, format *byte, args ...any) {
	var replace_buffer *zend.ZendString = nil
	var replace_origin *zend.ZendString = nil
	var buffer *byte = nil
	var docref_buf *byte = nil
	var target *byte = nil
	var docref_target *byte = ""
	var docref_root *byte = ""
	var p *byte
	var buffer_len int = 0
	var space *byte = ""
	var class_name *byte = ""
	var function *byte
	var origin_len int
	var origin *byte
	var message *byte
	var is_function int = 0

	/* get error text into buffer and escape for html if necessary */

	buffer_len = int(zend.ZendVspprintf(&buffer, 0, format, args))
	if CoreGlobals.GetHtmlErrors() != 0 {
		replace_buffer = standard.PhpEscapeHtmlEntities((*uint8)(buffer), buffer_len, 0, 2, GetSafeCharsetHint())

		/* Retry with substituting invalid chars on fail. */

		if replace_buffer == nil || replace_buffer.len_ < 1 {
			replace_buffer = standard.PhpEscapeHtmlEntities((*uint8)(buffer), buffer_len, 0, 2|8, GetSafeCharsetHint())
		}
		zend._efree(buffer)
		if replace_buffer != nil {
			buffer = replace_buffer.val
			buffer_len = int(replace_buffer.len_)
		} else {
			buffer = ""
			buffer_len = 0
		}
	}

	/* which function caused the problem if any at all */

	if PhpDuringModuleStartup() != 0 {
		function = "PHP Startup"
	} else if PhpDuringModuleShutdown() != 0 {
		function = "PHP Shutdown"
	} else if zend.EG.current_execute_data != nil && zend.EG.current_execute_data.func_ != nil && (zend.EG.current_execute_data.func_.common.type_&1) == 0 && zend.EG.current_execute_data.opline != nil && zend.EG.current_execute_data.opline.opcode == 73 {
		switch zend.EG.current_execute_data.opline.extended_value {
		case 1 << 0:
			function = "eval"
			is_function = 1
			break
		case 1 << 1:
			function = "include"
			is_function = 1
			break
		case 1 << 2:
			function = "include_once"
			is_function = 1
			break
		case 1 << 3:
			function = "require"
			is_function = 1
			break
		case 1 << 4:
			function = "require_once"
			is_function = 1
			break
		default:
			function = "Unknown"
		}
	} else {
		function = zend.GetActiveFunctionName()
		if function == nil || !(strlen(function)) {
			function = "Unknown"
		} else {
			is_function = 1
			class_name = zend.GetActiveClassName(&space)
		}
	}

	/* if we still have memory then format the origin */

	if is_function != 0 {
		origin_len = int(zend.ZendSpprintf(&origin, 0, "%s%s%s(%s)", class_name, space, function, params))
	} else {
		origin_len = int(zend.ZendSpprintf(&origin, 0, "%s", function))
	}
	if CoreGlobals.GetHtmlErrors() != 0 {
		replace_origin = standard.PhpEscapeHtmlEntities((*uint8)(origin), origin_len, 0, 2, GetSafeCharsetHint())
		zend._efree(origin)
		origin = replace_origin.val
	}

	/* origin and buffer available, so lets come up with the error message */

	if docref != nil && docref[0] == '#' {
		docref_target = strchr(docref, '#')
		docref = nil
	}

	/* no docref given but function is known (the default) */

	if docref == nil && is_function != 0 {
		var doclen int
		for (*function) == '_' {
			function++
		}
		if space[0] == '0' {
			doclen = int(zend.ZendSpprintf(&docref_buf, 0, "function.%s", function))
		} else {
			doclen = int(zend.ZendSpprintf(&docref_buf, 0, "%s.%s", class_name, function))
		}
		for g.Assign(&p, strchr(docref_buf, '_')) != nil {
			*p = '-'
		}
		docref = standard.PhpStrtolower(docref_buf, doclen)
	}

	/* we have a docref for a function AND
	 * - we show errors in html mode AND
	 * - the user wants to see the links
	 */

	if docref != nil && is_function != 0 && CoreGlobals.GetHtmlErrors() != 0 && strlen(CoreGlobals.GetDocrefRoot()) {
		if strncmp(docref, "http://", 7) {

			/* We don't have 'http://' so we use docref_root */

			var ref *byte
			docref_root = CoreGlobals.GetDocrefRoot()
			ref = zend._estrdup(docref)
			if docref_buf != nil {
				zend._efree(docref_buf)
			}
			docref_buf = ref

			/* strip of the target if any */

			p = strrchr(ref, '#')
			if p != nil {
				target = zend._estrdup(p)
				if target != nil {
					docref_target = target
					*p = '0'
				}
			}

			/* add the extension if it is set in ini */

			if CoreGlobals.GetDocrefExt() != nil && strlen(CoreGlobals.GetDocrefExt()) {
				zend.ZendSpprintf(&docref_buf, 0, "%s%s", ref, CoreGlobals.GetDocrefExt())
				zend._efree(ref)
			}
			docref = docref_buf
		}

		/* display html formatted or only show the additional links */

		if CoreGlobals.GetHtmlErrors() != 0 {
			zend.ZendSpprintf(&message, 0, "%s [<a href='%s%s%s'>%s</a>]: %s", origin, docref_root, docref, docref_target, docref, buffer)
		} else {
			zend.ZendSpprintf(&message, 0, "%s [%s%s%s]: %s", origin, docref_root, docref, docref_target, buffer)
		}
		if target != nil {
			zend._efree(target)
		}
	} else {
		zend.ZendSpprintf(&message, 0, "%s: %s", origin, buffer)
	}
	if replace_origin != nil {
		zend.ZendStringFree(replace_origin)
	} else {
		zend._efree(origin)
	}
	if docref_buf != nil {
		zend._efree(docref_buf)
	}
	if CoreGlobals.GetTrackErrors() != 0 && ModuleInitialized != 0 && zend.EG.active != 0 && (zend.EG.user_error_handler.u1.v.type_ == 0 || (zend.EG.user_error_handler_error_reporting&type_) == 0) {
		var tmp zend.Zval
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = zend.ZendStringInit(buffer, buffer_len, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		if zend.EG.current_execute_data != nil {
			if zend.ZendSetLocalVarStr("php_errormsg", g.SizeOf("\"php_errormsg\"")-1, &tmp, 0) == zend.FAILURE {
				zend.ZvalPtrDtor(&tmp)
			}
		} else {
			zend.ZendHashStrUpdateInd(&zend.EG.symbol_table, "php_errormsg", g.SizeOf("\"php_errormsg\"")-1, &tmp)
		}
	}
	if replace_buffer != nil {
		zend.ZendStringFree(replace_buffer)
	} else {
		zend._efree(buffer)
	}
	zend.ZendError(type_, "%s", message)
	zend._efree(message)
}

/* }}} */

func PhpErrorDocref(docref string, type_ int, format string, _ ...any) {
	var args va_list
	va_start(args, format)
	PhpVerror(docref, "", type_, format, args)
	va_end(args)
}

/* }}} */

func PhpErrorDocref1(docref *byte, param1 *byte, type_ int, format string, _ ...any) {
	var args va_list
	va_start(args, format)
	PhpVerror(docref, param1, type_, format, args)
	va_end(args)
}

/* }}} */

func PhpErrorDocref2(docref *byte, param1 *byte, param2 *byte, type_ int, format string, _ ...any) {
	var params *byte
	var args va_list
	zend.ZendSpprintf(&params, 0, "%s,%s", param1, param2)
	va_start(args, format)
	PhpVerror(docref, g.Cond(params != nil, params, "..."), type_, format, args)
	va_end(args)
	if params != nil {
		zend._efree(params)
	}
}

/* }}} */

/* {{{ php_html_puts */

func PhpHtmlPuts(str *byte, size int) { zend.ZendHtmlPuts(str, size) }

/* }}} */

func PhpErrorCb(type_ int, error_filename *byte, error_lineno uint32, format *byte, args ...any) {
	var buffer *byte
	var buffer_len int
	var display int
	buffer_len = int(zend.ZendVspprintf(&buffer, CoreGlobals.GetLogErrorsMaxLen(), format, args))

	/* check for repeated errors to be ignored */

	if CoreGlobals.GetIgnoreRepeatedErrors() != 0 && CoreGlobals.GetLastErrorMessage() != nil {

		/* no check for PG(last_error_file) is needed since it cannot
		 * be NULL if PG(last_error_message) is not NULL */

		if strcmp(CoreGlobals.GetLastErrorMessage(), buffer) || CoreGlobals.GetIgnoreRepeatedSource() == 0 && (CoreGlobals.GetLastErrorLineno() != int(error_lineno) || strcmp(CoreGlobals.GetLastErrorFile(), error_filename)) {
			display = 1
		} else {
			display = 0
		}

		/* no check for PG(last_error_file) is needed since it cannot
		 * be NULL if PG(last_error_message) is not NULL */

	} else {
		display = 1
	}

	/* according to error handling mode, throw exception or show it */

	if zend.EG.error_handling == zend.EH_THROW {
		switch type_ {
		case 1 << 0:

		case 1 << 4:

		case 1 << 6:

		case 1 << 8:

		case 1 << 2:

			/* fatal errors are real errors and cannot be made exceptions */

			break
		case 1 << 11:

		case 1 << 13:

		case 1 << 14:

			/* for the sake of BC to old damaged code */

			break
		case 1 << 3:

		case 1 << 10:

			/* notices are no errors and are not treated as such like E_WARNINGS */

			break
		default:

			/* throw an exception if we are in EH_THROW mode
			 * but DO NOT overwrite a pending exception
			 */

			if zend.EG.exception == nil {
				zend.ZendThrowErrorException(zend.EG.exception_class, buffer, 0, type_)
			}
			zend._efree(buffer)
			return
		}
	}

	/* store the error if it has changed */

	if display != 0 {
		if CoreGlobals.GetLastErrorMessage() != nil {
			var s *byte = CoreGlobals.GetLastErrorMessage()
			CoreGlobals.SetLastErrorMessage(nil)
			zend.Free(s)
		}
		if CoreGlobals.GetLastErrorFile() != nil {
			var s *byte = CoreGlobals.GetLastErrorFile()
			CoreGlobals.SetLastErrorFile(nil)
			zend.Free(s)
		}
		if error_filename == nil {
			error_filename = "Unknown"
		}
		CoreGlobals.SetLastErrorType(type_)
		CoreGlobals.SetLastErrorMessage(strdup(buffer))
		CoreGlobals.SetLastErrorFile(strdup(error_filename))
		CoreGlobals.SetLastErrorLineno(error_lineno)
	}

	/* display/log the error if necessary */

	if display != 0 && ((zend.EG.error_reporting&type_) != 0 || (type_&(1<<4|1<<5)) != 0) && (CoreGlobals.GetLogErrors() != 0 || CoreGlobals.GetDisplayErrors() != 0 || ModuleInitialized == 0) {
		var error_type_str *byte
		var syslog_type_int int = LOG_NOTICE
		switch type_ {
		case 1 << 0:

		case 1 << 4:

		case 1 << 6:

		case 1 << 8:
			error_type_str = "Fatal error"
			syslog_type_int = LOG_ERR
			break
		case 1 << 12:
			error_type_str = "Recoverable fatal error"
			syslog_type_int = LOG_ERR
			break
		case 1 << 1:

		case 1 << 5:

		case 1 << 7:

		case 1 << 9:
			error_type_str = "Warning"
			syslog_type_int = LOG_WARNING
			break
		case 1 << 2:
			error_type_str = "Parse error"
			syslog_type_int = LOG_ERR
			break
		case 1 << 3:

		case 1 << 10:
			error_type_str = "Notice"
			syslog_type_int = LOG_NOTICE
			break
		case 1 << 11:
			error_type_str = "Strict Standards"
			syslog_type_int = LOG_INFO
			break
		case 1 << 13:

		case 1 << 14:
			error_type_str = "Deprecated"
			syslog_type_int = LOG_INFO
			break
		default:
			error_type_str = "Unknown error"
			break
		}
		if ModuleInitialized == 0 || CoreGlobals.GetLogErrors() != 0 {
			var log_buffer *byte
			zend.ZendSpprintf(&log_buffer, 0, "PHP %s:  %s in %s on line %"+"u", error_type_str, buffer, error_filename, error_lineno)
			PhpLogErrWithSeverity(log_buffer, syslog_type_int)
			zend._efree(log_buffer)
		}
		if CoreGlobals.GetDisplayErrors() != 0 && (ModuleInitialized != 0 && CoreGlobals.GetDuringRequestStartup() == 0 || CoreGlobals.GetDisplayStartupErrors() != 0) {
			if CoreGlobals.GetXmlrpcErrors() != 0 {
				PhpPrintf("<?xml version=\"1.0\"?><methodResponse><fault><value><struct><member><name>faultCode</name><value><int>"+"%"+"lld"+"</int></value></member><member><name>faultString</name><value><string>%s:%s in %s on line %"+"u"+"</string></value></member></struct></value></fault></methodResponse>", CoreGlobals.GetXmlrpcErrorNumber(), error_type_str, buffer, error_filename, error_lineno)
			} else {
				var prepend_string *byte = zend.ZendIniStringEx("error_prepend_string", g.SizeOf("\"error_prepend_string\"")-1, 0, nil)
				var append_string *byte = zend.ZendIniStringEx("error_append_string", g.SizeOf("\"error_append_string\"")-1, 0, nil)
				if CoreGlobals.GetHtmlErrors() != 0 {
					if type_ == 1<<0 || type_ == 1<<2 {
						var buf *zend.ZendString = standard.PhpEscapeHtmlEntities((*uint8)(buffer), buffer_len, 0, 2, GetSafeCharsetHint())
						PhpPrintf("%s<br />\n<b>%s</b>:  %s in <b>%s</b> on line <b>%"+"u"+"</b><br />\n%s", g.Cond(prepend_string != nil, prepend_string, ""), error_type_str, buf.val, error_filename, error_lineno, g.Cond(append_string != nil, append_string, ""))
						zend.ZendStringFree(buf)
					} else {
						PhpPrintf("%s<br />\n<b>%s</b>:  %s in <b>%s</b> on line <b>%"+"u"+"</b><br />\n%s", g.Cond(prepend_string != nil, prepend_string, ""), error_type_str, buffer, error_filename, error_lineno, g.Cond(append_string != nil, append_string, ""))
					}
				} else {

					/* Write CLI/CGI errors to stderr if display_errors = "stderr" */

					if (!(strcmp(sapi_module.GetName(), "cli")) || !(strcmp(sapi_module.GetName(), "cgi")) || !(strcmp(sapi_module.GetName(), "phpdbg"))) && CoreGlobals.GetDisplayErrors() == 2 {
						r.Fprintf(stderr, "%s: %s in %s on line %"+"u"+"\n", error_type_str, buffer, error_filename, error_lineno)
					} else {
						PhpPrintf("%s\n%s: %s in %s on line %"+"u"+"\n%s", g.Cond(prepend_string != nil, prepend_string, ""), error_type_str, buffer, error_filename, error_lineno, g.Cond(append_string != nil, append_string, ""))
					}

					/* Write CLI/CGI errors to stderr if display_errors = "stderr" */

				}
			}
		}
	}

	/* Bail out if we can't recover */

	switch type_ {
	case 1 << 4:
		if ModuleInitialized == 0 {

			/* bad error in module startup - no way we can live with this */

			exit(-2)

			/* bad error in module startup - no way we can live with this */

		}
	case 1 << 0:

	case 1 << 12:

	case 1 << 2:

	case 1 << 6:

	case 1 << 8:
		zend.EG.exit_status = 255
		if ModuleInitialized != 0 {
			if CoreGlobals.GetDisplayErrors() == 0 && sapi_globals.GetHeadersSent() == 0 && sapi_globals.GetSapiHeaders().GetHttpResponseCode() == 200 {
				var ctr SapiHeaderLine = SapiHeaderLine{0}
				ctr.SetLine("HTTP/1.0 500 Internal Server Error")
				ctr.SetLineLen(g.SizeOf("\"HTTP/1.0 500 Internal Server Error\"") - 1)
				SapiHeaderOp(SAPI_HEADER_REPLACE, &ctr)
			}

			/* the parser would return 1 (failure), we can bail out nicely */

			if type_ != 1<<2 {

				/* restore memory limit */

				zend.ZendSetMemoryLimit(CoreGlobals.GetMemoryLimit())
				zend._efree(buffer)
				zend.ZendObjectsStoreMarkDestructed(&zend.EG.objects_store)
				zend._zendBailout(__FILE__, __LINE__)
				return
			}

			/* the parser would return 1 (failure), we can bail out nicely */

		}
		break
	}

	/* Log if necessary */

	if display == 0 {
		zend._efree(buffer)
		return
	}
	if CoreGlobals.GetTrackErrors() != 0 && ModuleInitialized != 0 && zend.EG.active != 0 {
		var tmp zend.Zval
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = zend.ZendStringInit(buffer, buffer_len, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		if zend.EG.current_execute_data != nil {
			if zend.ZendSetLocalVarStr("php_errormsg", g.SizeOf("\"php_errormsg\"")-1, &tmp, 0) == zend.FAILURE {
				zend.ZvalPtrDtor(&tmp)
			}
		} else {
			zend.ZendHashStrUpdateInd(&zend.EG.symbol_table, "php_errormsg", g.SizeOf("\"php_errormsg\"")-1, &tmp)
		}
	}
	zend._efree(buffer)
}

/* }}} */

func PhpGetCurrentUser() *byte {
	var pstat *zend.ZendStatT
	if sapi_globals.GetRequestInfo().GetCurrentUser() != nil {
		return sapi_globals.GetRequestInfo().GetCurrentUser()
	}

	/* FIXME: I need to have this somehow handled if
	   USE_SAPI is defined, because cgi will also be
	   interfaced in USE_SAPI */

	pstat = SapiGetStat()
	if pstat == nil {
		return ""
	} else {
		var pwd *__struct__passwd
		if g.Assign(&pwd, getpwuid(pstat.st_uid)) == nil {
			return ""
		}
		sapi_globals.GetRequestInfo().SetCurrentUserLength(strlen(pwd.pw_name))
		sapi_globals.GetRequestInfo().SetCurrentUser(zend._estrndup(pwd.pw_name, sapi_globals.GetRequestInfo().GetCurrentUserLength()))
		return sapi_globals.GetRequestInfo().GetCurrentUser()
	}
}

/* }}} */

func ZifSetTimeLimit(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var new_timeout zend.ZendLong
	var new_timeout_str *byte
	var new_timeout_strlen int
	var key *zend.ZendString
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &new_timeout) == zend.FAILURE {
		return
	}
	new_timeout_strlen = int(zend.ZendSpprintf(&new_timeout_str, 0, "%"+"lld", new_timeout))
	key = zend.ZendStringInit("max_execution_time", g.SizeOf("\"max_execution_time\"")-1, 0)
	if zend.ZendAlterIniEntryCharsEx(key, new_timeout_str, new_timeout_strlen, 1<<0, 1<<4, 0) == zend.SUCCESS {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	zend.ZendStringReleaseEx(key, 0)
	zend._efree(new_timeout_str)
}

/* }}} */

func PhpFopenWrapperForZend(filename *byte, opened_path **zend.ZendString) *r.FILE {
	return streams._phpStreamOpenWrapperAsFile((*byte)(filename), "rb", 0x1|0|0x8|0x80, opened_path)
}

/* }}} */

func PhpZendStreamCloser(handle any) {
	_phpStreamFree((*PhpStream)(handle), 1|2)
}

/* }}} */

func PhpZendStreamFsizer(handle any) int {
	var stream *PhpStream = handle
	var ssb PhpStreamStatbuf

	/* File size reported by stat() may be inaccurate if stream filters are used.
	 * TODO: Should stat() be generally disabled if filters are used? */

	if stream.readfilters.head != nil {
		return 0
	}
	if _phpStreamStat(stream, &ssb) == 0 {
		return ssb.sb.st_size
	}
	return 0
}

/* }}} */

func PhpStreamOpenForZend(filename *byte, handle *zend.ZendFileHandle) int {
	return PhpStreamOpenForZendEx(filename, handle, 0x1|0x8|0x80)
}

/* }}} */

func PhpStreamOpenForZendEx(filename *byte, handle *zend.ZendFileHandle, mode int) int {
	var opened_path *zend.ZendString
	var stream *PhpStream = _phpStreamOpenWrapperEx((*byte)(filename), "rb", mode, &opened_path, nil)
	if stream != nil {
		memset(handle, 0, g.SizeOf("zend_file_handle"))
		handle.type_ = zend.ZEND_HANDLE_STREAM
		handle.filename = (*byte)(filename)
		handle.opened_path = opened_path
		handle.handle.stream.handle = stream
		handle.handle.stream.reader = zend.ZendStreamReaderT(_phpStreamRead)
		handle.handle.stream.fsizer = PhpZendStreamFsizer
		handle.handle.stream.isatty = 0
		handle.handle.stream.closer = PhpZendStreamCloser

		/* suppress warning if this stream is not explicitly closed */

		stream.SetExposed(1)

		/* Disable buffering to avoid double buffering between PHP and Zend streams. */

		_phpStreamSetOption(stream, 2, 0, nil)
		return zend.SUCCESS
	}
	return zend.FAILURE
}

/* }}} */

func PhpResolvePathForZend(filename *byte, filename_len int) *zend.ZendString {
	return PhpResolvePath(filename, filename_len, CoreGlobals.GetIncludePath())
}

/* }}} */

func PhpGetConfigurationDirectiveForZend(name *zend.ZendString) *zend.Zval {
	return CfgGetEntryEx(name)
}

/* }}} */

func PhpFreeRequestGlobals() {
	if CoreGlobals.GetLastErrorMessage() != nil {
		zend.Free(CoreGlobals.GetLastErrorMessage())
		CoreGlobals.SetLastErrorMessage(nil)
	}
	if CoreGlobals.GetLastErrorFile() != nil {
		zend.Free(CoreGlobals.GetLastErrorFile())
		CoreGlobals.SetLastErrorFile(nil)
	}
	if CoreGlobals.GetPhpSysTempDir() != nil {
		zend._efree(CoreGlobals.GetPhpSysTempDir())
		CoreGlobals.SetPhpSysTempDir(nil)
	}
}

/* }}} */

func PhpMessageHandlerForZend(message zend.ZendLong, data any) {
	switch message {
	case 1:
		PhpErrorDocref("function.include", 1<<1, "Failed opening '%s' for inclusion (include_path='%s')", PhpStripUrlPasswd((*byte)(data)), g.CondF1(CoreGlobals.GetIncludePath() != nil, func() *byte { return CoreGlobals.GetIncludePath() }, ""))
		break
	case 2:
		PhpErrorDocref("function.require", 1<<6, "Failed opening required '%s' (include_path='%s')", PhpStripUrlPasswd((*byte)(data)), g.CondF1(CoreGlobals.GetIncludePath() != nil, func() *byte { return CoreGlobals.GetIncludePath() }, ""))
		break
	case 3:
		PhpErrorDocref(nil, 1<<1, "Failed opening '%s' for highlighting", PhpStripUrlPasswd((*byte)(data)))
		break
	case 4:

	case 5:
		break
	case 7:
		break
	case 6:
		var ta *__struct__tm
		var tmbuf __struct__tm
		var curtime int64
		var datetime_str *byte
		var asctimebuf []*byte
		var memory_leak_buf []byte
		time(&curtime)
		ta = localtime_r(&curtime, &tmbuf)
		datetime_str = asctime_r(ta, asctimebuf)
		if datetime_str != nil {
			datetime_str[strlen(datetime_str)-1] = 0
			ApPhpSnprintf(memory_leak_buf, g.SizeOf("memory_leak_buf"), "[%s]  Script:  '%s'\n", datetime_str, g.CondF1(sapi_globals.GetRequestInfo().GetPathTranslated() != nil, func() *byte { return sapi_globals.GetRequestInfo().GetPathTranslated() }, "-"))
		} else {
			ApPhpSnprintf(memory_leak_buf, g.SizeOf("memory_leak_buf"), "[null]  Script:  '%s'\n", g.CondF1(sapi_globals.GetRequestInfo().GetPathTranslated() != nil, func() *byte { return sapi_globals.GetRequestInfo().GetPathTranslated() }, "-"))
		}
		r.Fprintf(stderr, "%s", memory_leak_buf)
		break
	}
}

/* }}} */

func PhpOnTimeout(seconds int) {
	CoreGlobals.SetConnectionStatus(CoreGlobals.GetConnectionStatus() | 2)
}

/* {{{ php_request_startup
 */

func PhpRequestStartup() int {
	var retval int = zend.SUCCESS
	zend.ZendInternedStringsActivate()
	var __orig_bailout *sigjmp_buf = zend.EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		CoreGlobals.SetInErrorLog(0)
		CoreGlobals.SetDuringRequestStartup(1)
		PhpOutputActivate()

		/* initialize global variables */

		CoreGlobals.SetModulesActivated(0)
		CoreGlobals.SetHeaderIsBeingSent(0)
		CoreGlobals.SetConnectionStatus(0)
		CoreGlobals.SetInUserInclude(0)
		zend.ZendActivate()
		SapiActivate()
		zend.ZendSignalActivate()
		if CoreGlobals.GetMaxInputTime() == -1 {
			zend.ZendSetTimeout(zend.EG.timeout_seconds, 1)
		} else {
			zend.ZendSetTimeout(CoreGlobals.GetMaxInputTime(), 1)
		}

		/* Disable realpath cache if an open_basedir is set */

		if CoreGlobals.GetOpenBasedir() != nil && (*(CoreGlobals.GetOpenBasedir())) {
			zend.CwdGlobals.realpath_cache_size_limit = 0
		}
		if CoreGlobals.GetExposePhp() != 0 {
			SapiAddHeaderEx("X-Powered-By: PHP/"+"7.4.33", g.SizeOf("SAPI_PHP_VERSION_HEADER")-1, 1, 1)
		}
		if CoreGlobals.GetOutputHandler() != nil && CoreGlobals.GetOutputHandler()[0] {
			var oh zend.Zval
			var _s *byte = CoreGlobals.GetOutputHandler()
			var __z *zend.Zval = &oh
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			PhpOutputStartUser(&oh, 0, 0x70)
			zend.ZvalPtrDtor(&oh)
		} else if CoreGlobals.GetOutputBuffering() != 0 {
			PhpOutputStartUser(nil, g.CondF1(CoreGlobals.GetOutputBuffering() > 1, func() zend.ZendLong { return CoreGlobals.GetOutputBuffering() }, 0), 0x70)
		} else if CoreGlobals.GetImplicitFlush() != 0 {
			PhpOutputSetImplicitFlush(1)
		}

		/* We turn this off in php_execute_script() */

		PhpHashEnvironment()
		zend.ZendActivateModules()
		CoreGlobals.SetModulesActivated(1)
	} else {
		zend.EG.bailout = __orig_bailout
		retval = zend.FAILURE
	}
	zend.EG.bailout = __orig_bailout
	sapi_globals.SetSapiStarted(1)
	return retval
}

/* }}} */

func PhpRequestShutdown(dummy any) {
	var report_memleaks zend.ZendBool
	zend.EG.flags |= 1 << 0
	report_memleaks = CoreGlobals.GetReportMemleaks()

	/* EG(current_execute_data) points into nirvana and therefore cannot be safely accessed
	 * inside zend_executor callback functions.
	 */

	zend.EG.current_execute_data = nil
	PhpDeactivateTicks()

	/* 1. Call all possible shutdown functions registered with register_shutdown_function() */

	if CoreGlobals.GetModulesActivated() != 0 {
		var __orig_bailout *sigjmp_buf = zend.EG.bailout
		var __bailout sigjmp_buf
		zend.EG.bailout = &__bailout
		if sigsetjmp(__bailout, 0) == 0 {
			standard.PhpCallShutdownFunctions()
		}
		zend.EG.bailout = __orig_bailout
	}

	/* 2. Call all possible __destruct() functions */

	var __orig_bailout *sigjmp_buf = EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		zend.ZendCallDestructors()
	}
	zend.EG.bailout = __orig_bailout

	/* 3. Flush all output buffers */

	var __orig_bailout *sigjmp_buf = EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		var send_buffer zend.ZendBool = g.Cond(sapi_globals.GetRequestInfo().GetHeadersOnly() != 0, 0, 1)
		if zend.CG.unclean_shutdown != 0 && CoreGlobals.GetLastErrorType() == 1<<0 && size_t(CoreGlobals.GetMemoryLimit()) < zend.ZendMemoryUsage(1) {
			send_buffer = 0
		}
		if send_buffer == 0 {
			PhpOutputDiscardAll()
		} else {
			PhpOutputEndAll()
		}
	}
	zend.EG.bailout = __orig_bailout

	/* 4. Reset max_execution_time (no longer executing php code after response sent) */

	var __orig_bailout *sigjmp_buf = EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		zend.ZendUnsetTimeout()
	}
	zend.EG.bailout = __orig_bailout

	/* 5. Call all extensions RSHUTDOWN functions */

	if CoreGlobals.GetModulesActivated() != 0 {
		zend.ZendDeactivateModules()
	}

	/* 6. Shutdown output layer (send the set HTTP headers, cleanup output handlers, etc.) */

	var __orig_bailout *sigjmp_buf = EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		PhpOutputDeactivate()
	}
	zend.EG.bailout = __orig_bailout

	/* 7. Free shutdown functions */

	if CoreGlobals.GetModulesActivated() != 0 {
		standard.PhpFreeShutdownFunctions()
	}

	/* 8. Destroy super-globals */

	var __orig_bailout *sigjmp_buf = EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		var i int
		for i = 0; i < 6; i++ {
			zend.ZvalPtrDtor(&CoreGlobals.GetHttpGlobals()[i])
		}
	}
	zend.EG.bailout = __orig_bailout

	/* 9. free request-bound globals */

	PhpFreeRequestGlobals()

	/* 10. Shutdown scanner/executor/compiler and restore ini entries */

	zend.ZendDeactivate()

	/* 11. Call all extensions post-RSHUTDOWN functions */

	var __orig_bailout *sigjmp_buf = EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		zend.ZendPostDeactivateModules()
	}
	zend.EG.bailout = __orig_bailout

	/* 12. SAPI related shutdown (free stuff) */

	var __orig_bailout *sigjmp_buf = EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		SapiDeactivate()
	}
	zend.EG.bailout = __orig_bailout

	/* 13. free virtual CWD memory */

	zend.VirtualCwdDeactivate()

	/* 14. Destroy stream hashes */

	var __orig_bailout *sigjmp_buf = EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		PhpShutdownStreamHashes()
	}
	zend.EG.bailout = __orig_bailout

	/* 15. Free Willy (here be crashes) */

	zend.ZendInternedStringsDeactivate()
	var __orig_bailout *sigjmp_buf = zend.EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		zend.ShutdownMemoryManager(zend.CG.unclean_shutdown != 0 || report_memleaks == 0, 0)
	}
	zend.EG.bailout = __orig_bailout

	/* Reset memory limit, as the reset during INI_STAGE_DEACTIVATE may have failed.
	 * At this point, no memory beyond a single chunk should be in use. */

	zend.ZendSetMemoryLimit(CoreGlobals.GetMemoryLimit())

	/* 16. Deactivate Zend signals */

	zend.ZendSignalDeactivate()

	/* 16. Deactivate Zend signals */
}

/* }}} */

func PhpComInitialize() {}

/* }}} */

/* {{{ core_globals_dtor
 */

func CoreGlobalsDtor(core_globals *PhpCoreGlobals) {
	if core_globals.GetLastErrorMessage() != nil {
		zend.Free(core_globals.GetLastErrorMessage())
	}
	if core_globals.GetLastErrorFile() != nil {
		zend.Free(core_globals.GetLastErrorFile())
	}
	if core_globals.GetDisableFunctions() != nil {
		zend.Free(core_globals.GetDisableFunctions())
	}
	if core_globals.GetDisableClasses() != nil {
		zend.Free(core_globals.GetDisableClasses())
	}
	if core_globals.GetPhpBinary() != nil {
		zend.Free(core_globals.GetPhpBinary())
	}
	PhpShutdownTicks()
}

/* }}} */

func ZmInfoPhpCore(zend_module *zend.ZendModuleEntry) {
	standard.PhpInfoPrintTableStart()
	standard.PhpInfoPrintTableRow(2, "PHP Version", "7.4.33")
	standard.PhpInfoPrintTableEnd()
	DisplayIniEntries(zend_module)
}

/* }}} */

func PhpRegisterExtensions(ptr **zend.ZendModuleEntry, count int) int {
	var end **zend.ZendModuleEntry = ptr + count
	for ptr < end {
		if (*ptr) != nil {
			if zend.ZendRegisterInternalModule(*ptr) == nil {
				return zend.FAILURE
			}
		}
		ptr++
	}
	return zend.SUCCESS
}

/* A very long time ago php_module_startup() was refactored in a way
 * which broke calling it with more than one additional module.
 * This alternative to php_register_extensions() works around that
 * by walking the shallower structure.
 *
 * See algo: https://bugs.php.net/bug.php?id=63159
 */

func PhpRegisterExtensionsBc(ptr *zend.ZendModuleEntry, count int) int {
	for g.PostDec(&count) {
		if zend.ZendRegisterInternalModule(g.PostInc(&ptr)) == nil {
			return zend.FAILURE
		}
	}
	return zend.SUCCESS
}

/* }}} */

/* {{{ php_module_startup
 */

func PhpModuleStartup(sf *sapi_module_struct, additional_modules *zend.ZendModuleEntry, num_additional_modules uint32) int {
	var zuf zend.ZendUtilityFunctions
	var zuv zend.ZendUtilityValues
	var retval int = zend.SUCCESS
	var module_number int = 0
	var php_os *byte
	var module *zend.ZendModuleEntry
	php_os = "Darwin"
	ModuleShutdown = 0
	ModuleStartup = 1
	SapiInitializeEmptyRequest()
	SapiActivate()
	if ModuleInitialized != 0 {
		return zend.SUCCESS
	}
	sapi_module = *sf
	PhpOutputStartup()
	memset(&CoreGlobals, 0, g.SizeOf("core_globals"))
	PhpStartupTicks()
	zend.GcGlobalsCtor()
	zuf.error_function = PhpErrorCb
	zuf.printf_function = PhpPrintf
	zuf.write_function = PhpOutputWrite
	zuf.fopen_function = PhpFopenWrapperForZend
	zuf.message_handler = PhpMessageHandlerForZend
	zuf.get_configuration_directive = PhpGetConfigurationDirectiveForZend
	zuf.ticks_function = PhpRunTicks
	zuf.on_timeout = PhpOnTimeout
	zuf.stream_open_function = PhpStreamOpenForZend
	zuf.printf_to_smart_string_function = PhpPrintfToSmartString
	zuf.printf_to_smart_str_function = PhpPrintfToSmartStr
	zuf.getenv_function = SapiGetenv
	zuf.resolve_path_function = PhpResolvePathForZend
	zend.ZendStartup(&zuf)
	setlocale(LC_CTYPE, "")
	tzset()
	zend.LeIndexPtr = zend.ZendRegisterListDestructorsEx(nil, nil, "index pointer", 0)

	/* Register constants */

	zend.ZendRegisterStringlConstant("PHP_VERSION", g.SizeOf("\"PHP_VERSION\"")-1, "7.4.33", g.SizeOf("PHP_VERSION")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterLongConstant("PHP_MAJOR_VERSION", g.SizeOf("\"PHP_MAJOR_VERSION\"")-1, 7, 1<<1|1<<0, 0)
	zend.ZendRegisterLongConstant("PHP_MINOR_VERSION", g.SizeOf("\"PHP_MINOR_VERSION\"")-1, 4, 1<<1|1<<0, 0)
	zend.ZendRegisterLongConstant("PHP_RELEASE_VERSION", g.SizeOf("\"PHP_RELEASE_VERSION\"")-1, 33, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_EXTRA_VERSION", g.SizeOf("\"PHP_EXTRA_VERSION\"")-1, "", g.SizeOf("PHP_EXTRA_VERSION")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterLongConstant("PHP_VERSION_ID", g.SizeOf("\"PHP_VERSION_ID\"")-1, 70433, 1<<1|1<<0, 0)
	zend.ZendRegisterLongConstant("PHP_ZTS", g.SizeOf("\"PHP_ZTS\"")-1, 0, 1<<1|1<<0, 0)
	zend.ZendRegisterLongConstant("PHP_DEBUG", g.SizeOf("\"PHP_DEBUG\"")-1, 0, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_OS", g.SizeOf("\"PHP_OS\"")-1, php_os, strlen(php_os), 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_OS_FAMILY", g.SizeOf("\"PHP_OS_FAMILY\"")-1, "Unknown", g.SizeOf("PHP_OS_FAMILY")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_SAPI", g.SizeOf("\"PHP_SAPI\"")-1, sapi_module.GetName(), strlen(sapi_module.GetName()), 1<<1|1<<0|1<<3, 0)
	zend.ZendRegisterStringlConstant("DEFAULT_INCLUDE_PATH", g.SizeOf("\"DEFAULT_INCLUDE_PATH\"")-1, ".:", g.SizeOf("PHP_INCLUDE_PATH")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PEAR_INSTALL_DIR", g.SizeOf("\"PEAR_INSTALL_DIR\"")-1, "", g.SizeOf("PEAR_INSTALLDIR")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PEAR_EXTENSION_DIR", g.SizeOf("\"PEAR_EXTENSION_DIR\"")-1, "/usr/local/lib/php/extensions/no-debug-non-zts-20190902", g.SizeOf("PHP_EXTENSION_DIR")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_EXTENSION_DIR", g.SizeOf("\"PHP_EXTENSION_DIR\"")-1, "/usr/local/lib/php/extensions/no-debug-non-zts-20190902", g.SizeOf("PHP_EXTENSION_DIR")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_PREFIX", g.SizeOf("\"PHP_PREFIX\"")-1, "/usr/local", g.SizeOf("PHP_PREFIX")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_BINDIR", g.SizeOf("\"PHP_BINDIR\"")-1, "/usr/local/bin", g.SizeOf("PHP_BINDIR")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_MANDIR", g.SizeOf("\"PHP_MANDIR\"")-1, "/usr/local/php/man", g.SizeOf("PHP_MANDIR")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_LIBDIR", g.SizeOf("\"PHP_LIBDIR\"")-1, "/usr/local/lib/php", g.SizeOf("PHP_LIBDIR")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_DATADIR", g.SizeOf("\"PHP_DATADIR\"")-1, "/usr/local/share/php", g.SizeOf("PHP_DATADIR")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_SYSCONFDIR", g.SizeOf("\"PHP_SYSCONFDIR\"")-1, "/usr/local/etc", g.SizeOf("PHP_SYSCONFDIR")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_LOCALSTATEDIR", g.SizeOf("\"PHP_LOCALSTATEDIR\"")-1, "/usr/local/var", g.SizeOf("PHP_LOCALSTATEDIR")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_CONFIG_FILE_PATH", g.SizeOf("\"PHP_CONFIG_FILE_PATH\"")-1, "/usr/local/lib", strlen("/usr/local/lib"), 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_CONFIG_FILE_SCAN_DIR", g.SizeOf("\"PHP_CONFIG_FILE_SCAN_DIR\"")-1, "", g.SizeOf("PHP_CONFIG_FILE_SCAN_DIR")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_SHLIB_SUFFIX", g.SizeOf("\"PHP_SHLIB_SUFFIX\"")-1, "so", g.SizeOf("PHP_SHLIB_SUFFIX")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterStringlConstant("PHP_EOL", g.SizeOf("\"PHP_EOL\"")-1, "\n", g.SizeOf("PHP_EOL")-1, 1<<1|1<<0, 0)
	zend.ZendRegisterLongConstant("PHP_MAXPATHLEN", g.SizeOf("\"PHP_MAXPATHLEN\"")-1, 256, 1<<1|1<<0, 0)
	zend.ZendRegisterLongConstant("PHP_INT_MAX", g.SizeOf("\"PHP_INT_MAX\"")-1, INT64_MAX, 1<<1|1<<0, 0)
	zend.ZendRegisterLongConstant("PHP_INT_MIN", g.SizeOf("\"PHP_INT_MIN\"")-1, INT64_MIN, 1<<1|1<<0, 0)
	zend.ZendRegisterLongConstant("PHP_INT_SIZE", g.SizeOf("\"PHP_INT_SIZE\"")-1, 8, 1<<1|1<<0, 0)
	zend.ZendRegisterLongConstant("PHP_FD_SETSIZE", g.SizeOf("\"PHP_FD_SETSIZE\"")-1, FD_SETSIZE, 1<<1|1<<0, 0)
	zend.ZendRegisterLongConstant("PHP_FLOAT_DIG", g.SizeOf("\"PHP_FLOAT_DIG\"")-1, DBL_DIG, 1<<1|1<<0, 0)
	zend.ZendRegisterDoubleConstant("PHP_FLOAT_EPSILON", g.SizeOf("\"PHP_FLOAT_EPSILON\"")-1, DBL_EPSILON, 1<<1|1<<0, 0)
	zend.ZendRegisterDoubleConstant("PHP_FLOAT_MAX", g.SizeOf("\"PHP_FLOAT_MAX\"")-1, DBL_MAX, 1<<1|1<<0, 0)
	zend.ZendRegisterDoubleConstant("PHP_FLOAT_MIN", g.SizeOf("\"PHP_FLOAT_MIN\"")-1, DBL_MIN, 1<<1|1<<0, 0)
	PhpBinaryInit()
	if CoreGlobals.GetPhpBinary() != nil {
		zend.ZendRegisterStringlConstant("PHP_BINARY", g.SizeOf("\"PHP_BINARY\"")-1, CoreGlobals.GetPhpBinary(), strlen(CoreGlobals.GetPhpBinary()), 1<<1|1<<0|1<<3, 0)
	} else {
		zend.ZendRegisterStringlConstant("PHP_BINARY", g.SizeOf("\"PHP_BINARY\"")-1, "", 0, 1<<1|1<<0|1<<3, 0)
	}
	PhpOutputRegisterConstants()
	PhpRfc1867RegisterConstants()

	/* this will read in php.ini, set up the configuration parameters,
	   load zend extensions and register php function extensions
	   to be loaded later */

	if PhpInitConfig() == zend.FAILURE {
		return zend.FAILURE
	}

	/* Register PHP core ini entries */

	zend.ZendRegisterIniEntries(IniEntries, module_number)

	/* Register Zend ini entries */

	zend.ZendRegisterStandardIniEntries()

	/* Disable realpath cache if an open_basedir is set */

	if CoreGlobals.GetOpenBasedir() != nil && (*(CoreGlobals.GetOpenBasedir())) {
		zend.CwdGlobals.realpath_cache_size_limit = 0
	}
	CoreGlobals.SetHaveCalledOpenlog(0)

	/* initialize stream wrappers registry
	 * (this uses configuration parameters from php.ini)
	 */

	if PhpInitStreamWrappers(module_number) == zend.FAILURE {
		PhpPrintf("PHP:  Unable to initialize stream url wrappers.\n")
		return zend.FAILURE
	}
	zuv.html_errors = 1
	PhpStartupAutoGlobals()
	zend.ZendSetUtilityValues(&zuv)
	PhpStartupSapiContentTypes()

	/* startup extensions statically compiled in */

	if PhpRegisterInternalExtensionsFunc() == zend.FAILURE {
		PhpPrintf("Unable to start builtin modules\n")
		return zend.FAILURE
	}

	/* start additional PHP extensions */

	PhpRegisterExtensionsBc(additional_modules, num_additional_modules)

	/* load and startup extensions compiled as shared objects (aka DLLs)
	   as requested by php.ini entries
	   these are loaded after initialization of internal extensions
	   as extensions *might* rely on things from ext/standard
	   which is always an internal extension and to be initialized
	   ahead of all other internals
	*/

	PhpIniRegisterExtensions()
	zend.ZendStartupModules()

	/* start Zend extensions */

	zend.ZendStartupExtensions()
	zend.ZendCollectModuleHandlers()

	/* register additional functions */

	if sapi_module.GetAdditionalFunctions() != nil {
		if g.Assign(&module, zend.ZendHashStrFindPtr(&zend.ModuleRegistry, "standard", g.SizeOf("\"standard\"")-1)) != nil {
			zend.EG.current_module = module
			zend.ZendRegisterFunctions(nil, sapi_module.GetAdditionalFunctions(), nil, 1)
			zend.EG.current_module = nil
		}
	}

	/* disable certain classes and functions as requested by php.ini */

	PhpDisableFunctions()
	PhpDisableClasses()

	/* make core report what it should */

	if g.Assign(&module, zend.ZendHashStrFindPtr(&zend.ModuleRegistry, "core", g.SizeOf("\"core\"")-1)) != nil {
		module.version = "7.4.33"
		module.info_func = ZmInfoPhpCore
	}
	ModuleInitialized = 1
	if zend.ZendPostStartup() != zend.SUCCESS {
		return zend.FAILURE
	}

	/* Check for deprecated directives */

	var directives []struct {
		error_level long
		phrase      *byte
		directives  []*byte
	} = []struct {
		error_level long
		phrase      *byte
		directives  []*byte
	}{
		{1 << 13, "Directive '%s' is deprecated", {"track_errors", "allow_url_include", nil}},
		{
			1 << 4,
			"Directive '%s' is no longer available in PHP",
			{"allow_call_time_pass_reference", "asp_tags", "define_syslog_variables", "highlight.bg", "magic_quotes_gpc", "magic_quotes_runtime", "magic_quotes_sybase", "register_globals", "register_long_arrays", "safe_mode", "safe_mode_gid", "safe_mode_include_dir", "safe_mode_exec_dir", "safe_mode_allowed_env_vars", "safe_mode_protected_env_vars", "zend.ze1_compatibility_mode", nil},
		},
	}
	var i uint
	var __orig_bailout *sigjmp_buf = zend.EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {

		/* 2 = Count of deprecation structs */

		for i = 0; i < 2; i++ {
			var p **byte = directives[i].directives
			for (*p) != nil {
				var value zend.ZendLong
				if CfgGetLong((*byte)(*p), &value) == zend.SUCCESS && value != 0 {
					zend.ZendError(directives[i].error_level, directives[i].phrase, *p)
				}
				p++
			}
		}

		/* 2 = Count of deprecation structs */

	} else {
		zend.EG.bailout = __orig_bailout
		retval = zend.FAILURE
	}
	zend.EG.bailout = __orig_bailout
	zend.VirtualCwdDeactivate()
	SapiDeactivate()
	ModuleStartup = 0
	zend.ShutdownMemoryManager(1, 0)
	zend.VirtualCwdActivate()
	zend.ZendInternedStringsSwitchStorage(1)

	/* we're done */

	return retval

	/* we're done */
}

/* }}} */

func PhpModuleShutdownWrapper(sapi_globals *sapi_module_struct) int {
	PhpModuleShutdown()
	return zend.SUCCESS
}

/* }}} */

func PhpModuleShutdown() {
	var module_number int = 0
	ModuleShutdown = 1
	if ModuleInitialized == 0 {
		return
	}
	zend.ZendInternedStringsSwitchStorage(0)
	SapiFlush()
	zend.ZendShutdown()

	/* Destroys filter & transport registries too */

	PhpShutdownStreamWrappers(module_number)
	zend.ZendUnregisterIniEntries(module_number)

	/* close down the ini config */

	PhpShutdownConfig()
	zend.ZendIniShutdown()
	zend.ShutdownMemoryManager(zend.CG.unclean_shutdown, 1)
	PhpOutputShutdown()
	zend.ZendInternedStringsDtor()
	if zend.ZendPostShutdownCb != nil {
		var cb func() = zend.ZendPostShutdownCb
		zend.ZendPostShutdownCb = nil
		cb()
	}
	ModuleInitialized = 0
	CoreGlobalsDtor(&CoreGlobals)
	zend.GcGlobalsDtor()
}

/* }}} */

func PhpExecuteScript(primary_file *zend.ZendFileHandle) int {
	var prepend_file_p *zend.ZendFileHandle
	var append_file_p *zend.ZendFileHandle
	var prepend_file zend.ZendFileHandle
	var append_file zend.ZendFileHandle
	var old_cwd *byte
	var retval int = 0
	zend.EG.exit_status = 0

	// #define OLD_CWD_SIZE       4096

	old_cwd = zend._emalloc(4096)
	old_cwd[0] = '0'
	var __orig_bailout *sigjmp_buf = zend.EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		var realfile []byte
		CoreGlobals.SetDuringRequestStartup(0)
		if primary_file.filename != nil && (sapi_globals.GetOptions()&1) == 0 {
			void(getcwd(old_cwd, 4096-1))
			zend.VirtualChdirFile(primary_file.filename, chdir)
		}

		/* Only lookup the real file path and add it to the included_files list if already opened
		 *   otherwise it will get opened and added to the included_files list in zend_execute_scripts
		 */

		if primary_file.filename != nil && strcmp("Standard input code", primary_file.filename) && primary_file.opened_path == nil && primary_file.type_ != zend.ZEND_HANDLE_FILENAME {
			if ExpandFilepath(primary_file.filename, realfile) != nil {
				primary_file.opened_path = zend.ZendStringInit(realfile, strlen(realfile), 0)
				zend.ZendHashAddEmptyElement(&zend.EG.included_files, primary_file.opened_path)
			}
		}
		if CoreGlobals.GetAutoPrependFile() != nil && CoreGlobals.GetAutoPrependFile()[0] {
			zend.ZendStreamInitFilename(&prepend_file, CoreGlobals.GetAutoPrependFile())
			prepend_file_p = &prepend_file
		} else {
			prepend_file_p = nil
		}
		if CoreGlobals.GetAutoAppendFile() != nil && CoreGlobals.GetAutoAppendFile()[0] {
			zend.ZendStreamInitFilename(&append_file, CoreGlobals.GetAutoAppendFile())
			append_file_p = &append_file
		} else {
			append_file_p = nil
		}
		if CoreGlobals.GetMaxInputTime() != -1 {
			zend.ZendSetTimeout(zend.ZendIniLong("max_execution_time", g.SizeOf("\"max_execution_time\"")-1, 0), 0)
		}

		/*
		   If cli primary file has shabang line and there is a prepend file,
		   the `skip_shebang` will be used by prepend file but not primary file,
		   save it and restore after prepend file been executed.
		*/

		if zend.CG.skip_shebang != 0 && prepend_file_p != nil {
			zend.CG.skip_shebang = 0
			if zend.ZendExecuteScripts(1<<3, nil, 1, prepend_file_p) == zend.SUCCESS {
				zend.CG.skip_shebang = 1
				retval = zend.ZendExecuteScripts(1<<3, nil, 2, primary_file, append_file_p) == zend.SUCCESS
			}
		} else {
			retval = zend.ZendExecuteScripts(1<<3, nil, 3, prepend_file_p, primary_file, append_file_p) == zend.SUCCESS
		}

		/*
		   If cli primary file has shabang line and there is a prepend file,
		   the `skip_shebang` will be used by prepend file but not primary file,
		   save it and restore after prepend file been executed.
		*/

	}
	zend.EG.bailout = __orig_bailout
	if zend.EG.exception != nil {
		var __orig_bailout *sigjmp_buf = zend.EG.bailout
		var __bailout sigjmp_buf
		zend.EG.bailout = &__bailout
		if sigsetjmp(__bailout, 0) == 0 {
			zend.ZendExceptionError(zend.EG.exception, 1<<0)
		}
		zend.EG.bailout = __orig_bailout
	}
	if old_cwd[0] != '0' {
		void(chdir(old_cwd))
	}
	zend._efree(old_cwd)
	return retval
}

/* }}} */

func PhpExecuteSimpleScript(primary_file *zend.ZendFileHandle, ret *zend.Zval) int {
	var old_cwd *byte
	zend.EG.exit_status = 0

	// #define OLD_CWD_SIZE       4096

	old_cwd = zend._emalloc(4096)
	old_cwd[0] = '0'
	var __orig_bailout *sigjmp_buf = zend.EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		CoreGlobals.SetDuringRequestStartup(0)
		if primary_file.filename != nil && (sapi_globals.GetOptions()&1) == 0 {
			void(getcwd(old_cwd, 4096-1))
			zend.VirtualChdirFile(primary_file.filename, chdir)
		}
		zend.ZendExecuteScripts(1<<3, ret, 1, primary_file)
	}
	zend.EG.bailout = __orig_bailout
	if old_cwd[0] != '0' {
		void(chdir(old_cwd))
	}
	zend._efree(old_cwd)
	return zend.EG.exit_status
}

/* }}} */

func PhpHandleAbortedConnection() {
	CoreGlobals.SetConnectionStatus(1)
	PhpOutputSetStatus(0x2)
	if CoreGlobals.GetIgnoreUserAbort() == 0 {
		zend._zendBailout(__FILE__, __LINE__)
	}
}

/* }}} */

func PhpHandleAuthData(auth *byte) int {
	var ret int = -1
	var auth_len int = g.CondF1(auth != nil, func() __auto__ { return strlen(auth) }, 0)
	if auth != nil && auth_len > 0 && zend.ZendBinaryStrncasecmp(auth, auth_len, "Basic ", g.SizeOf("\"Basic \"")-1, g.SizeOf("\"Basic \"")-1) == 0 {
		var pass *byte
		var user *zend.ZendString
		user = standard.PhpBase64Decode((*uint8)(auth+6), auth_len-6)
		if user != nil {
			pass = strchr(user.val, ':')
			if pass != nil {
				g.PostInc(&(*pass)) = '0'
				sapi_globals.GetRequestInfo().SetAuthUser(zend._estrndup(user.val, user.len_))
				sapi_globals.GetRequestInfo().SetAuthPassword(zend._estrdup(pass))
				ret = 0
			}
			zend.ZendStringFree(user)
		}
	}
	if ret == -1 {
		sapi_globals.GetRequestInfo().SetAuthPassword(nil)
		sapi_globals.GetRequestInfo().SetAuthUser(sapi_globals.GetRequestInfo().GetAuthPassword())
	} else {
		sapi_globals.GetRequestInfo().SetAuthDigest(nil)
	}
	if ret == -1 && auth != nil && auth_len > 0 && zend.ZendBinaryStrncasecmp(auth, auth_len, "Digest ", g.SizeOf("\"Digest \"")-1, g.SizeOf("\"Digest \"")-1) == 0 {
		sapi_globals.GetRequestInfo().SetAuthDigest(zend._estrdup(auth + 7))
		ret = 0
	}
	if ret == -1 {
		sapi_globals.GetRequestInfo().SetAuthDigest(nil)
	}
	return ret
}

/* }}} */

func PhpLintScript(file *zend.ZendFileHandle) int {
	var op_array *zend.ZendOpArray
	var retval int = zend.FAILURE
	var __orig_bailout *sigjmp_buf = zend.EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		op_array = zend.ZendCompileFile(file, 1<<1)
		zend.ZendDestroyFileHandle(file)
		if op_array != nil {
			zend.DestroyOpArray(op_array)
			zend._efree(op_array)
			retval = zend.SUCCESS
		}
	}
	zend.EG.bailout = __orig_bailout
	if zend.EG.exception != nil {
		zend.ZendExceptionError(zend.EG.exception, 1<<0)
	}
	return retval
}

/* }}} */
