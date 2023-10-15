package core

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

var PhpInternalEncodingChanged func() = nil
var OnChangeBrowscap func(
	entry *zend.ZendIniEntry,
	new_value *types.String,
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

var IniEntries []zend.ZendIniEntryDef = []zend.ZendIniEntryDef{
	*zend.NewZendIniEntryDef("highlight.comment", PHP_INI_ALL).Value(zend.HL_COMMENT_COLOR).
		Displayer(PhpIniColorDisplayerCb),
	*zend.NewZendIniEntryDef("highlight.default", PHP_INI_ALL).Value(zend.HL_DEFAULT_COLOR).
		Displayer(PhpIniColorDisplayerCb),
	*zend.NewZendIniEntryDef("highlight.html", PHP_INI_ALL).Value(zend.HL_HTML_COLOR).
		Displayer(PhpIniColorDisplayerCb),
	*zend.NewZendIniEntryDef("highlight.keyword", PHP_INI_ALL).Value(zend.HL_KEYWORD_COLOR).
		Displayer(PhpIniColorDisplayerCb),
	*zend.NewZendIniEntryDef("highlight.string", PHP_INI_ALL).Value(zend.HL_STRING_COLOR).
		Displayer(PhpIniColorDisplayerCb),
	*zend.NewZendIniEntryDef("display_errors", PHP_INI_ALL).Value("1").
		Displayer(DisplayErrorsMode).
		OnModifyArgs(
			OnUpdateDisplayErrors, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDisplayErrors()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("display_startup_errors", PHP_INI_ALL).Value("0").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDisplayStartupErrors()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("enable_dl", PHP_INI_SYSTEM).Value("1").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetEnableDl()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("expose_php", PHP_INI_SYSTEM).Value("1").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetExposePhp()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("docref_root", PHP_INI_ALL).Value("").
		OnModifyArgs(
			zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDocrefRoot()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("docref_ext", PHP_INI_ALL).Value("").
		OnModifyArgs(
			zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDocrefExt()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("html_errors", PHP_INI_ALL).Value("1").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetHtmlErrors()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("xmlrpc_errors", PHP_INI_SYSTEM).Value("0").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetXmlrpcErrors()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("xmlrpc_error_number", PHP_INI_ALL).Value("0").
		OnModifyArgs(
			zend.OnUpdateLong, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetXmlrpcErrorNumber()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("max_input_time", PHP_INI_SYSTEM|PHP_INI_PERDIR).Value("-1").
		OnModifyArgs(
			zend.OnUpdateLong, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMaxInputTime()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("ignore_user_abort", PHP_INI_ALL).Value("0").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetIgnoreUserAbort()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("implicit_flush", PHP_INI_ALL).Value("0").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetImplicitFlush()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("log_errors", PHP_INI_ALL).Value("0").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetLogErrors()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("log_errors_max_len", PHP_INI_ALL).Value("1024").
		OnModifyArgs(
			zend.OnUpdateLong, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetLogErrorsMaxLen()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("ignore_repeated_errors", PHP_INI_ALL).Value("0").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetIgnoreRepeatedErrors()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("ignore_repeated_source", PHP_INI_ALL).Value("0").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetIgnoreRepeatedSource()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("report_memleaks", PHP_INI_ALL).Value("1").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetReportMemleaks()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("report_zend_debug", PHP_INI_ALL).Value("1").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetReportZendDebug()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("output_buffering", PHP_INI_PERDIR|PHP_INI_SYSTEM).Value("0").
		OnModifyArgs(
			zend.OnUpdateLong, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetOutputBuffering()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("output_handler", PHP_INI_PERDIR|PHP_INI_SYSTEM).
		OnModifyArgs(
			zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetOutputHandler()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("register_argc_argv", PHP_INI_PERDIR|PHP_INI_SYSTEM).Value("1").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetRegisterArgcArgv()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("auto_globals_jit", PHP_INI_PERDIR|PHP_INI_SYSTEM).Value("1").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAutoGlobalsJit()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("short_open_tag", PHP_INI_SYSTEM|PHP_INI_PERDIR).Value(DEFAULT_SHORT_OPEN_TAG).
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*zend.ZendCompilerGlobals)(nil).GetShortTags()))-(*byte)(nil))), any(&zend.CompilerGlobals), nil,
		),
	*zend.NewZendIniEntryDef("track_errors", PHP_INI_ALL).Value("0").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetTrackErrors()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("unserialize_callback_func", PHP_INI_ALL).
		OnModifyArgs(
			zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUnserializeCallbackFunc()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("serialize_precision", PHP_INI_ALL).Value("-1").
		OnModifyArgs(
			OnSetSerializePrecision, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSerializePrecision()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("arg_separator.output", PHP_INI_ALL).Value("&").
		OnModifyArgs(
			zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetArgSeparator().GetOutput()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("arg_separator.input", PHP_INI_SYSTEM|PHP_INI_PERDIR).Value("&").
		OnModifyArgs(
			zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetArgSeparator().GetInput()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("auto_append_file", PHP_INI_SYSTEM|PHP_INI_PERDIR).
		OnModifyArgs(
			zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAutoAppendFile()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("auto_prepend_file", PHP_INI_SYSTEM|PHP_INI_PERDIR).
		OnModifyArgs(
			zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAutoPrependFile()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("doc_root", PHP_INI_SYSTEM).
		OnModifyArgs(
			zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetDocRoot()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("default_charset", PHP_INI_ALL).Value(PHP_DEFAULT_CHARSET).
		OnModifyArgs(
			OnUpdateDefaultCharset, any(zend_long((*byte)(&((*SapiGlobals)(nil).DefaultCharset()))-(*byte)(nil))), any(SG__()), nil,
		),
	*zend.NewZendIniEntryDef("default_mimetype", PHP_INI_ALL).Value(SAPI_DEFAULT_MIMETYPE).
		OnModifyArgs(
			OnUpdateDefaultMimeTye, any(zend_long((*byte)(&((*SapiGlobals)(nil).DefaultMimetype()))-(*byte)(nil))), any(SG__()), nil,
		),
	*zend.NewZendIniEntryDef("internal_encoding", PHP_INI_ALL).
		OnModifyArgs(
			OnUpdateInternalEncoding, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetInternalEncoding()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("input_encoding", PHP_INI_ALL).
		OnModifyArgs(
			OnUpdateInputEncoding, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetInputEncoding()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("output_encoding", PHP_INI_ALL).
		OnModifyArgs(
			OnUpdateOutputEncoding, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetOutputEncoding()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("error_log", PHP_INI_ALL).
		OnModifyArgs(
			OnUpdateErrorLog, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetErrorLog()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("extension_dir", PHP_INI_SYSTEM).Value(PHP_EXTENSION_DIR).
		OnModifyArgs(
			zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetExtensionDir()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("sys_temp_dir", PHP_INI_SYSTEM).
		OnModifyArgs(
			zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSysTempDir()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("include_path", PHP_INI_ALL).Value(PHP_INCLUDE_PATH).
		OnModifyArgs(
			zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetIncludePath()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("max_execution_time", PHP_INI_ALL).Value("30").
		OnModifyArgs(
			OnUpdateTimeout, nil, nil, nil,
		),
	*zend.NewZendIniEntryDef("open_basedir", PHP_INI_ALL).
		OnModifyArgs(
			OnUpdateBaseDir, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetOpenBasedir()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("file_uploads", PHP_INI_SYSTEM).Value("1").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetFileUploads()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("upload_max_filesize", PHP_INI_SYSTEM|PHP_INI_PERDIR).Value("2M").
		OnModifyArgs(
			zend.OnUpdateLong, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUploadMaxFilesize()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("post_max_size", PHP_INI_SYSTEM|PHP_INI_PERDIR).Value("8M").
		OnModifyArgs(
			zend.OnUpdateLong, any(zend_long((*byte)(&((*SapiGlobals)(nil).PostMaxSize()))-(*byte)(nil))), any(SG__()), nil,
		),
	*zend.NewZendIniEntryDef("upload_tmp_dir", PHP_INI_SYSTEM).
		OnModifyArgs(
			zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUploadTmpDir()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("max_input_nesting_level", PHP_INI_SYSTEM|PHP_INI_PERDIR).Value("64").
		OnModifyArgs(
			zend.OnUpdateLongGEZero, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMaxInputNestingLevel()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("max_input_vars", PHP_INI_SYSTEM|PHP_INI_PERDIR).Value("1000").
		OnModifyArgs(
			zend.OnUpdateLongGEZero, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMaxInputVars()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("user_dir", PHP_INI_SYSTEM).
		OnModifyArgs(
			zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUserDir()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("variables_order", PHP_INI_SYSTEM|PHP_INI_PERDIR).Value("EGPCS").
		OnModifyArgs(
			zend.OnUpdateStringUnempty, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetVariablesOrder()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("request_order", PHP_INI_SYSTEM|PHP_INI_PERDIR).
		OnModifyArgs(
			zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetRequestOrder()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("error_append_string", PHP_INI_ALL).
		OnModifyArgs(
			zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetErrorAppendString()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("error_prepend_string", PHP_INI_ALL).
		OnModifyArgs(
			zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetErrorPrependString()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("SMTP", PHP_INI_ALL).Value("localhost"),
	*zend.NewZendIniEntryDef("smtp_port", PHP_INI_ALL).Value("25"),
	*zend.NewZendIniEntryDef("mail.add_x_header", PHP_INI_SYSTEM|PHP_INI_PERDIR).Value("0").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMailXHeader()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("mail.log", PHP_INI_SYSTEM|PHP_INI_PERDIR).
		OnModifyArgs(
			OnUpdateMailLog, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetMailLog()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("browscap", PHP_INI_SYSTEM).
		OnModifyArgs(
			OnChangeBrowscap, nil, nil, nil,
		),
	*zend.NewZendIniEntryDef("memory_limit", PHP_INI_ALL).Value("128M").
		OnModifyArgs(
			OnChangeMemoryLimit, nil, nil, nil,
		),
	*zend.NewZendIniEntryDef("precision", PHP_INI_ALL).Value("14").
		OnModifyArgs(
			OnSetPrecision, nil, nil, nil,
		),
	*zend.NewZendIniEntryDef("sendmail_from", PHP_INI_ALL),
	*zend.NewZendIniEntryDef("sendmail_path", PHP_INI_SYSTEM).Value(DEFAULT_SENDMAIL_PATH),
	*zend.NewZendIniEntryDef("mail.force_extra_parameters", PHP_INI_SYSTEM|PHP_INI_PERDIR).
		OnModifyArgs(
			OnChangeMailForceExtra, nil, nil, nil,
		),
	*zend.NewZendIniEntryDef("disable_functions", PHP_INI_SYSTEM).Value(""),
	*zend.NewZendIniEntryDef("disable_classes", PHP_INI_SYSTEM).Value(""),
	*zend.NewZendIniEntryDef("max_file_uploads", PHP_INI_SYSTEM|PHP_INI_PERDIR).Value("20"),
	*zend.NewZendIniEntryDef("allow_url_fopen", PHP_INI_SYSTEM).Value("1").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAllowUrlFopen()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("allow_url_include", PHP_INI_SYSTEM).Value("0").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetAllowUrlInclude()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("enable_post_data_reading", PHP_INI_SYSTEM|PHP_INI_PERDIR).Value("1").
		Displayer(zend.ZendIniBooleanDisplayerCb).
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetEnablePostDataReading()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("realpath_cache_size", PHP_INI_SYSTEM).Value("4096K").
		OnModifyArgs(
			zend.OnUpdateLong, any(zend_long((*byte)(&((*zend.VirtualCwdGlobals)(nil).GetRealpathCacheSizeLimit()))-(*byte)(nil))), any(&zend.CwdGlobals), nil,
		),
	*zend.NewZendIniEntryDef("realpath_cache_ttl", PHP_INI_SYSTEM).Value("120").
		OnModifyArgs(
			zend.OnUpdateLong, any(zend_long((*byte)(&((*zend.VirtualCwdGlobals)(nil).GetRealpathCacheTtl()))-(*byte)(nil))), any(&zend.CwdGlobals), nil,
		),
	*zend.NewZendIniEntryDef("user_ini.filename", PHP_INI_SYSTEM).Value(".user.ini").
		OnModifyArgs(
			zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUserIniFilename()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("user_ini.cache_ttl", PHP_INI_SYSTEM).Value("300").
		OnModifyArgs(
			zend.OnUpdateLong, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetUserIniCacheTtl()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("hard_timeout", PHP_INI_SYSTEM).Value("2").
		OnModifyArgs(
			zend.OnUpdateLong, any(zend_long((*byte)(&((*zend.ZendExecutorGlobals)(nil).GetHardTimeout()))-(*byte)(nil))), zend.EG__(), nil,
		),
	*zend.NewZendIniEntryDef("syslog.facility", PHP_INI_SYSTEM).Value("LOG_USER").
		OnModifyArgs(
			OnSetFacility, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSyslogFacility()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("syslog.ident", PHP_INI_SYSTEM).Value("php").
		OnModifyArgs(
			zend.OnUpdateString, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSyslogIdent()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
	*zend.NewZendIniEntryDef("syslog.filter", PHP_INI_ALL).Value("no-ctrl").
		OnModifyArgs(
			OnSetLogFilter, any(zend_long((*byte)(&((*PhpCoreGlobals)(nil).GetSyslogFilter()))-(*byte)(nil))), any(&CoreGlobals), nil,
		),
}
var ModuleInitialized int = 0
var ModuleStartup int = 1
var ModuleShutdown int = 0
