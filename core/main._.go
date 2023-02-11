// <<generate>>

package core

import (
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

// failed # include "ext/date/php_date.h"

var PhpRegisterInternalExtensionsFunc func() int = PhpRegisterInternalExtensions
var CoreGlobals PhpCoreGlobals

/* {{{ PHP_INI_MH
 */

var PhpInternalEncodingChanged func() = nil

/* {{{ PHP_INI_MH
 */

var OnChangeBrowscap func(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int

/* Need to be read from the environment (?):
 * PHP_AUTO_PREPEND_FILE
 * PHP_AUTO_APPEND_FILE
 * PHP_DOCUMENT_ROOT
 * PHP_USER_DIR
 * PHP_INCLUDE_PATH
 */

const DEFAULT_SENDMAIL_PATH = PHP_PROG_SENDMAIL + " -t -i"

/* {{{ PHP_INI
 */

var IniEntries []zend.ZendIniEntryDef = []zend.ZendIniEntryDef{
	zend.MakeZendIniEntryDef("highlight.comment", nil, nil, nil, nil, zend.HL_COMMENT_COLOR, PhpIniColorDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("highlight.default", nil, nil, nil, nil, zend.HL_DEFAULT_COLOR, PhpIniColorDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("highlight.html", nil, nil, nil, nil, zend.HL_HTML_COLOR, PhpIniColorDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("highlight.keyword", nil, nil, nil, nil, zend.HL_KEYWORD_COLOR, PhpIniColorDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("highlight.string", nil, nil, nil, nil, zend.HL_STRING_COLOR, PhpIniColorDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("display_errors", OnUpdateDisplayErrors, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDisplayErrors()))-(*byte)(nil))), any(&CoreGlobals), nil, "1", DisplayErrorsMode, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("display_startup_errors", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDisplayStartupErrors()))-(*byte)(nil))), any(&CoreGlobals), nil, "0", zend.ZendIniBooleanDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("enable_dl", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetEnableDl()))-(*byte)(nil))), any(&CoreGlobals), nil, "1", zend.ZendIniBooleanDisplayerCb, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("expose_php", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetExposePhp()))-(*byte)(nil))), any(&CoreGlobals), nil, "1", zend.ZendIniBooleanDisplayerCb, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("docref_root", zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDocrefRoot()))-(*byte)(nil))), any(&CoreGlobals), nil, "", nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("docref_ext", zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDocrefExt()))-(*byte)(nil))), any(&CoreGlobals), nil, "", nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("html_errors", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetHtmlErrors()))-(*byte)(nil))), any(&CoreGlobals), nil, "1", zend.ZendIniBooleanDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("xmlrpc_errors", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetXmlrpcErrors()))-(*byte)(nil))), any(&CoreGlobals), nil, "0", zend.ZendIniBooleanDisplayerCb, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("xmlrpc_error_number", zend.OnUpdateLong, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetXmlrpcErrorNumber()))-(*byte)(nil))), any(&CoreGlobals), nil, "0", nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("max_input_time", zend.OnUpdateLong, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMaxInputTime()))-(*byte)(nil))), any(&CoreGlobals), nil, "-1", nil, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("ignore_user_abort", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetIgnoreUserAbort()))-(*byte)(nil))), any(&CoreGlobals), nil, "0", zend.ZendIniBooleanDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("implicit_flush", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetImplicitFlush()))-(*byte)(nil))), any(&CoreGlobals), nil, "0", zend.ZendIniBooleanDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("log_errors", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetLogErrors()))-(*byte)(nil))), any(&CoreGlobals), nil, "0", zend.ZendIniBooleanDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("log_errors_max_len", zend.OnUpdateLong, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetLogErrorsMaxLen()))-(*byte)(nil))), any(&CoreGlobals), nil, "1024", nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("ignore_repeated_errors", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetIgnoreRepeatedErrors()))-(*byte)(nil))), any(&CoreGlobals), nil, "0", zend.ZendIniBooleanDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("ignore_repeated_source", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetIgnoreRepeatedSource()))-(*byte)(nil))), any(&CoreGlobals), nil, "0", zend.ZendIniBooleanDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("report_memleaks", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetReportMemleaks()))-(*byte)(nil))), any(&CoreGlobals), nil, "1", zend.ZendIniBooleanDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("report_zend_debug", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetReportZendDebug()))-(*byte)(nil))), any(&CoreGlobals), nil, "1", zend.ZendIniBooleanDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("output_buffering", zend.OnUpdateLong, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetOutputBuffering()))-(*byte)(nil))), any(&CoreGlobals), nil, "0", nil, PHP_INI_PERDIR|PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("output_handler", zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetOutputHandler()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_PERDIR|PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("register_argc_argv", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetRegisterArgcArgv()))-(*byte)(nil))), any(&CoreGlobals), nil, "1", zend.ZendIniBooleanDisplayerCb, PHP_INI_PERDIR|PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("auto_globals_jit", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAutoGlobalsJit()))-(*byte)(nil))), any(&CoreGlobals), nil, "1", zend.ZendIniBooleanDisplayerCb, PHP_INI_PERDIR|PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("short_open_tag", zend.OnUpdateBool, any(zend_long((*byte)(&((*zend.ZendCompilerGlobals)(nil).GetShortTags()))-(*byte)(nil))), any(&zend.CompilerGlobals), nil, DEFAULT_SHORT_OPEN_TAG, zend.ZendIniBooleanDisplayerCb, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("track_errors", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetTrackErrors()))-(*byte)(nil))), any(&CoreGlobals), nil, "0", zend.ZendIniBooleanDisplayerCb, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("unserialize_callback_func", zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUnserializeCallbackFunc()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("serialize_precision", OnSetSerializePrecision, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSerializePrecision()))-(*byte)(nil))), any(&CoreGlobals), nil, "-1", nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("arg_separator.output", zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetArgSeparator().GetOutput()))-(*byte)(nil))), any(&CoreGlobals), nil, "&", nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("arg_separator.input", zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetArgSeparator().GetInput()))-(*byte)(nil))), any(&CoreGlobals), nil, "&", nil, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("auto_append_file", zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAutoAppendFile()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("auto_prepend_file", zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAutoPrependFile()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("doc_root", zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDocRoot()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("default_charset", OnUpdateDefaultCharset, any(zend_long((*byte)(&((*sapi_globals_struct)(nil).GetDefaultCharset()))-(*byte)(nil))), any(&sapi_globals), nil, PHP_DEFAULT_CHARSET, nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("default_mimetype", OnUpdateDefaultMimeTye, any(zend_long((*byte)(&((*sapi_globals_struct)(nil).GetDefaultMimetype()))-(*byte)(nil))), any(&sapi_globals), nil, SAPI_DEFAULT_MIMETYPE, nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("internal_encoding", OnUpdateInternalEncoding, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetInternalEncoding()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("input_encoding", OnUpdateInputEncoding, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetInputEncoding()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("output_encoding", OnUpdateOutputEncoding, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetOutputEncoding()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("error_log", OnUpdateErrorLog, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetErrorLog()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("extension_dir", zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetExtensionDir()))-(*byte)(nil))), any(&CoreGlobals), nil, PHP_EXTENSION_DIR, nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("sys_temp_dir", zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSysTempDir()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("include_path", zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetIncludePath()))-(*byte)(nil))), any(&CoreGlobals), nil, PHP_INCLUDE_PATH, nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("max_execution_time", OnUpdateTimeout, nil, nil, nil, "30", nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("open_basedir", OnUpdateBaseDir, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetOpenBasedir()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("file_uploads", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetFileUploads()))-(*byte)(nil))), any(&CoreGlobals), nil, "1", zend.ZendIniBooleanDisplayerCb, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("upload_max_filesize", zend.OnUpdateLong, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUploadMaxFilesize()))-(*byte)(nil))), any(&CoreGlobals), nil, "2M", nil, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("post_max_size", zend.OnUpdateLong, any(zend_long((*byte)(&((*sapi_globals_struct)(nil).GetPostMaxSize()))-(*byte)(nil))), any(&sapi_globals), nil, "8M", nil, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("upload_tmp_dir", zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUploadTmpDir()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("max_input_nesting_level", zend.OnUpdateLongGEZero, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMaxInputNestingLevel()))-(*byte)(nil))), any(&CoreGlobals), nil, "64", nil, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("max_input_vars", zend.OnUpdateLongGEZero, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMaxInputVars()))-(*byte)(nil))), any(&CoreGlobals), nil, "1000", nil, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("user_dir", zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUserDir()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("variables_order", zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetVariablesOrder()))-(*byte)(nil))), any(&CoreGlobals), nil, "EGPCS", nil, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("request_order", zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetRequestOrder()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("error_append_string", zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetErrorAppendString()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("error_prepend_string", zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetErrorPrependString()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("SMTP", nil, nil, nil, nil, "localhost", nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("smtp_port", nil, nil, nil, nil, "25", nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("mail.add_x_header", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMailXHeader()))-(*byte)(nil))), any(&CoreGlobals), nil, "0", zend.ZendIniBooleanDisplayerCb, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("mail.log", OnUpdateMailLog, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMailLog()))-(*byte)(nil))), any(&CoreGlobals), nil, nil, nil, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("browscap", OnChangeBrowscap, nil, nil, nil, nil, nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("memory_limit", OnChangeMemoryLimit, nil, nil, nil, "128M", nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("precision", OnSetPrecision, nil, nil, nil, "14", nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("sendmail_from", nil, nil, nil, nil, nil, nil, PHP_INI_ALL),
	zend.MakeZendIniEntryDef("sendmail_path", nil, nil, nil, nil, DEFAULT_SENDMAIL_PATH, nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("mail.force_extra_parameters", OnChangeMailForceExtra, nil, nil, nil, nil, nil, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("disable_functions", nil, nil, nil, nil, "", nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("disable_classes", nil, nil, nil, nil, "", nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("max_file_uploads", nil, nil, nil, nil, "20", nil, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("allow_url_fopen", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAllowUrlFopen()))-(*byte)(nil))), any(&CoreGlobals), nil, "1", zend.ZendIniBooleanDisplayerCb, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("allow_url_include", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAllowUrlInclude()))-(*byte)(nil))), any(&CoreGlobals), nil, "0", zend.ZendIniBooleanDisplayerCb, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("enable_post_data_reading", zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetEnablePostDataReading()))-(*byte)(nil))), any(&CoreGlobals), nil, "1", zend.ZendIniBooleanDisplayerCb, PHP_INI_SYSTEM|PHP_INI_PERDIR),
	zend.MakeZendIniEntryDef("realpath_cache_size", zend.OnUpdateLong, any(zend_long((*byte)(&((*zend.VirtualCwdGlobals)(nil).GetRealpathCacheSizeLimit()))-(*byte)(nil))), any(&zend.CwdGlobals), nil, "4096K", nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("realpath_cache_ttl", zend.OnUpdateLong, any(zend_long((*byte)(&((*zend.VirtualCwdGlobals)(nil).GetRealpathCacheTtl()))-(*byte)(nil))), any(&zend.CwdGlobals), nil, "120", nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("user_ini.filename", zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUserIniFilename()))-(*byte)(nil))), any(&CoreGlobals), nil, ".user.ini", nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("user_ini.cache_ttl", zend.OnUpdateLong, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUserIniCacheTtl()))-(*byte)(nil))), any(&CoreGlobals), nil, "300", nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("hard_timeout", zend.OnUpdateLong, any(zend_long((*byte)(&((*zend.ZendExecutorGlobals)(nil).GetHardTimeout()))-(*byte)(nil))), any(&zend.ExecutorGlobals), nil, "2", nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("syslog.facility", OnSetFacility, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSyslogFacility()))-(*byte)(nil))), any(&CoreGlobals), nil, "LOG_USER", nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("syslog.ident", zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSyslogIdent()))-(*byte)(nil))), any(&CoreGlobals), nil, "php", nil, PHP_INI_SYSTEM),
	zend.MakeZendIniEntryDef("syslog.filter", OnSetLogFilter, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSyslogFilter()))-(*byte)(nil))), any(&CoreGlobals), nil, "no-ctrl", nil, PHP_INI_ALL),
}
var ModuleInitialized int = 0
var ModuleStartup int = 1
var ModuleShutdown int = 0

/* {{{ php_during_module_startup */

/* {{{ php_html_puts */

/* {{{ php_request_startup
 */

/* {{{ core_globals_dtor
 */

/* A very long time ago php_module_startup() was refactored in a way
 * which broke calling it with more than one additional module.
 * This alternative to php_register_extensions() works around that
 * by walking the shallower structure.
 *
 * See algo: https://bugs.php.net/bug.php?id=63159
 */

/* {{{ php_module_startup
 */
