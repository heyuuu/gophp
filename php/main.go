package php

import (
	"github.com/heyuuu/gophp/php/perr"
	"strings"
)

var PhpInternalEncodingChanged func() = nil

//const DefaultSendmailPath = PhpProgSendmail + " -t -i"

const DefaultCharset = "UTF-8"
const DefaultMimetype = "text/html"

var MainIniEntries = []*IniEntryDef{
	NewIniEntryDef("highlight.comment", IniAll).Value(HlCommentColor),
	NewIniEntryDef("highlight.default", IniAll).Value(HlDefaultColor),
	NewIniEntryDef("highlight.html", IniAll).Value(HlHtmlColor),
	NewIniEntryDef("highlight.keyword", IniAll).Value(HlKeywordColor),
	NewIniEntryDef("highlight.string", IniAll).Value(HlStringColor),
	//NewIniEntryDef("display_errors", IniAll).Value("1").OnModifyString(func(ctx *Context, s string) { ctx.PG().SetDisplayErrors(PhpGetDisplayErrorsMode(s)) }),
	NewIniEntryDef("display_startup_errors", IniAll).Value("0").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetDisplayStartupErrors(b) }),
	NewIniEntryDef("expose_php", IniSystem).Value("1").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetExposePhp(b) }),
	NewIniEntryDef("docref_root", IniAll).Value("").OnModifyString(func(ctx *Context, s string) { ctx.PG().SetDocrefRoot(s) }),
	NewIniEntryDef("docref_ext", IniAll).Value("").OnModifyString(func(ctx *Context, s string) { ctx.PG().SetDocrefExt(s) }),
	NewIniEntryDef("html_errors", IniAll).Value("1").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetHtmlErrors(b) }),
	NewIniEntryDef("xmlrpc_errors", IniSystem).Value("0").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetXmlrpcErrors(b) }),
	NewIniEntryDef("xmlrpc_error_number", IniAll).Value("0").OnModifyLong(func(ctx *Context, i int) { ctx.PG().SetXmlrpcErrorNumber(i) }),
	NewIniEntryDef("max_input_time", IniSystem|IniPerDir).Value("-1").OnModifyLong(func(ctx *Context, i int) { ctx.PG().SetMaxInputTime(i) }),
	NewIniEntryDef("ignore_user_abort", IniAll).Value("0").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetIgnoreUserAbort(b) }),
	NewIniEntryDef("implicit_flush", IniAll).Value("0").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetImplicitFlush(b) }),
	NewIniEntryDef("log_errors", IniAll).Value("0").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetLogErrors(b) }),
	NewIniEntryDef("log_errors_max_len", IniAll).Value("1024").OnModifyLong(func(ctx *Context, i int) { ctx.PG().SetLogErrorsMaxLen(i) }),
	NewIniEntryDef("ignore_repeated_errors", IniAll).Value("0").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetIgnoreRepeatedErrors(b) }),
	NewIniEntryDef("ignore_repeated_source", IniAll).Value("0").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetIgnoreRepeatedSource(b) }),
	NewIniEntryDef("report_memleaks", IniAll).Value("1").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetReportMemleaks(b) }),
	NewIniEntryDef("report_zend_debug", IniAll).Value("1").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetReportZendDebug(b) }),
	NewIniEntryDef("output_buffering", IniPerDir|IniSystem).Value("0").OnModifyLong(func(ctx *Context, i int) { ctx.PG().SetOutputBuffering(i) }),
	NewIniEntryDef("output_handler", IniPerDir|IniSystem).OnModifyString(func(ctx *Context, s string) { ctx.PG().SetOutputHandler(s) }),
	NewIniEntryDef("register_argc_argv", IniPerDir|IniSystem).Value("1").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetRegisterArgcArgv(b) }),
	NewIniEntryDef("auto_globals_jit", IniPerDir|IniSystem).Value("1").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetAutoGlobalsJit(true) }),
	//NewIniEntryDef("short_open_tag", IniSystem|IniPerDir).Value("1").OnModifyBool(func(ctx *Context, b bool) { ctx.CG().SetShortTags(b) }),
	NewIniEntryDef("track_errors", IniAll).Value("0").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetTrackErrors(b) }),
	NewIniEntryDef("unserialize_callback_func", IniAll).OnModifyString(func(ctx *Context, s string) { ctx.PG().SetUnserializeCallbackFunc(s) }),
	NewIniEntryDef("serialize_precision", IniAll).Value("-1").OnModify(onSetSerializePrecision),
	NewIniEntryDef("arg_separator.output", IniAll).Value("&"),
	NewIniEntryDef("arg_separator.input", IniSystem|IniPerDir).Value("&"),
	NewIniEntryDef("auto_append_file", IniSystem|IniPerDir).OnModifyString(func(ctx *Context, s string) { ctx.PG().SetAutoAppendFile(s) }),
	NewIniEntryDef("auto_prepend_file", IniSystem|IniPerDir).OnModifyString(func(ctx *Context, s string) { ctx.PG().SetAutoPrependFile(s) }),
	NewIniEntryDef("doc_root", IniSystem).OnModifyStringNotEmpty(func(ctx *Context, s string) { ctx.PG().SetDocRoot(s) }),
	NewIniEntryDef("default_charset", IniAll).Value(DefaultCharset).OnModify(onUpdateDefaultCharset),
	NewIniEntryDef("default_mimetype", IniAll).Value(DefaultMimetype).OnModify(onUpdateDefaultMimeType),
	NewIniEntryDef("internal_encoding", IniAll).OnModify(onUpdateInternalEncoding),
	NewIniEntryDef("input_encoding", IniAll).OnModify(onUpdateInputEncoding),
	NewIniEntryDef("output_encoding", IniAll).OnModify(onUpdateOutputEncoding),
	NewIniEntryDef("error_log", IniAll).OnModify(onUpdateErrorLog),
	NewIniEntryDef("extension_dir", IniSystem).Value(PHP_EXTENSION_DIR).OnModifyStringNotEmpty(func(ctx *Context, s string) { ctx.PG().SetExtensionDir(s) }),
	NewIniEntryDef("sys_temp_dir", IniSystem).OnModifyStringNotEmpty(func(ctx *Context, s string) { ctx.PG().SetSysTempDir(s) }),
	NewIniEntryDef("include_path", IniAll).Value(PHP_INCLUDE_PATH).OnModifyStringNotEmpty(func(ctx *Context, s string) { ctx.PG().SetIncludePath(s) }),
	NewIniEntryDef("max_execution_time", IniAll).Value("30").OnModify(onUpdateTimeout),
	//NewIniEntryDef("open_basedir", IniAll).OnModify(OnUpdateBaseDir),
	NewIniEntryDef("file_uploads", IniSystem).Value("1").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetFileUploads(b) }),
	NewIniEntryDef("upload_max_filesize", IniSystem|IniPerDir).Value("2M").OnModifyLong(func(ctx *Context, i int) { ctx.PG().SetUploadMaxFilesize(i) }),
	//NewIniEntryDef("post_max_size", IniSystem|IniPerDir).Value("8M").OnModifyLong(func(ctx *Context, i int) { ctx.SG().SetPostMaxSize(i) }),
	NewIniEntryDef("upload_tmp_dir", IniSystem).OnModifyStringNotEmpty(func(ctx *Context, s string) { ctx.PG().SetUploadTmpDir(s) }),
	NewIniEntryDef("max_input_nesting_level", IniSystem|IniPerDir).Value("64").OnModifyLongGEZero(func(ctx *Context, i int) { ctx.PG().SetMaxInputNestingLevel(i) }),
	NewIniEntryDef("max_input_vars", IniSystem|IniPerDir).Value("1000").OnModifyLongGEZero(func(ctx *Context, i int) { ctx.PG().SetMaxInputVars(i) }),
	NewIniEntryDef("user_dir", IniSystem).OnModifyString(func(ctx *Context, s string) { ctx.PG().SetUserDir(s) }),
	NewIniEntryDef("variables_order", IniSystem|IniPerDir).Value("EGPCS").OnModifyStringNotEmpty(func(ctx *Context, s string) { ctx.PG().SetVariablesOrder(s) }),
	NewIniEntryDef("request_order", IniSystem|IniPerDir).OnModifyString(func(ctx *Context, s string) { ctx.PG().SetRequestOrder(s) }),
	NewIniEntryDef("error_append_string", IniAll).OnModifyString(func(ctx *Context, s string) { ctx.PG().SetErrorAppendString(s) }),
	NewIniEntryDef("error_prepend_string", IniAll).OnModifyString(func(ctx *Context, s string) { ctx.PG().SetErrorPrependString(s) }),
	NewIniEntryDef("SMTP", IniAll).Value("localhost"),
	NewIniEntryDef("smtp_port", IniAll).Value("25"),
	NewIniEntryDef("mail.add_x_header", IniSystem|IniPerDir).Value("0").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetMailXHeader(b) }),
	NewIniEntryDef("mail.log", IniSystem|IniPerDir).OnModify(onUpdateMailLog),
	NewIniEntryDef("browscap", IniSystem).OnModify(onChangeBrowscap),
	NewIniEntryDef("memory_limit", IniAll).Value("128M").OnModify(nil),
	NewIniEntryDef("precision", IniAll).Value("14").OnModify(onSetPrecision),
	NewIniEntryDef("sendmail_from", IniAll),
	//NewIniEntryDef("sendmail_path", IniSystem).Value(DefaultSendmailPath),
	NewIniEntryDef("mail.force_extra_parameters", IniSystem|IniPerDir).OnModify(onChangeMailForceExtra),
	NewIniEntryDef("disable_functions", IniSystem).Value(""),
	NewIniEntryDef("disable_classes", IniSystem).Value(""),
	NewIniEntryDef("max_file_uploads", IniSystem|IniPerDir).Value("20"),
	NewIniEntryDef("allow_url_fopen", IniSystem).Value("1").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetAllowUrlFopen(b) }),
	NewIniEntryDef("allow_url_include", IniSystem).Value("0").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetAllowUrlInclude(b) }),
	NewIniEntryDef("enable_post_data_reading", IniSystem|IniPerDir).Value("1").OnModifyBool(func(ctx *Context, b bool) { ctx.PG().SetEnablePostDataReading(b) }),
	// notice: realpath 缓存功能已取消，获取 realpath 改为直接调用 go 标准库
	//*php.NewIniEntryDef("realpath_cache_size", IniSystem).Value("4096K").OnModifyLong(func(i int) { php.CWDG__().SetRealpathCacheSizeLimit(i) }),
	//*php.NewIniEntryDef("realpath_cache_ttl", IniSystem).Value("120").OnModifyLong(func(i int) { php.CWDG__().SetRealpathCacheTtl(i) }),
	NewIniEntryDef("user_ini.filename", IniSystem).Value(".user.ini").OnModifyString(func(ctx *Context, s string) { ctx.PG().SetUserIniFilename(s) }),
	NewIniEntryDef("user_ini.cache_ttl", IniSystem).Value("300").OnModifyLong(func(ctx *Context, i int) { ctx.PG().SetUserIniCacheTtl(i) }),
	//NewIniEntryDef("hard_timeout", IniSystem).Value("2").OnModifyLong(func(ctx *Context, i int) { ctx.EG().SetHardTimeout(i) }),
	NewIniEntryDef("syslog.facility", IniSystem).Value("LOG_USER").OnModify(onSetFacility),
	NewIniEntryDef("syslog.ident", IniSystem).Value("php").OnModifyString(func(ctx *Context, s string) { ctx.PG().SetSyslogIdent(s) }),
	NewIniEntryDef("syslog.filter", IniAll).Value("no-ctrl").OnModify(onSetLogFilter),

	// -- Zend ini entries --

	NewIniEntryDef("error_reporting", IniAll).OnModify(onUpdateErrorReporting),
	NewIniEntryDef("zend.assertions", IniAll).Value("1").OnModify(onUpdateAssertions),
	NewIniEntryDef("zend.signal_check", IniSystem).Value("0").OnModifyString(func(ctx *Context, newValue string) {
		//SIGG__().check = IniStringParseBool(newValue)
	}),
	NewIniEntryDef("zend.exception_ignore_args", IniAll).Value("0").OnModifyString(func(ctx *Context, newValue string) {
		//ctx.EG().SetExceptionIgnoreArgs(IniStringParseBool(newValue))
	}),
}

func onSetPrecision(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
	i := ParseLong10(newValue)
	if i < -1 {
		return false
	}

	ctx.EG().SetPrecision(i)
	return true
}
func onSetSerializePrecision(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
	i := ParseLong10(newValue)
	if i < -1 {
		return false
	}

	ctx.PG().SetSerializePrecision(i)
	return true
}

func onSetLogFilter(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
	switch newValue {
	case "all":
		//ctx.PG().SetSyslogFilter(PHP_SYSLOG_FILTER_ALL)
		return true
	case "no-ctrl":
		//ctx.PG().SetSyslogFilter(PHP_SYSLOG_FILTER_NO_CTRL)
		return true
	case "ascii":
		//ctx.PG().SetSyslogFilter(PHP_SYSLOG_FILTER_ASCII)
		return true
	case "raw":
		//ctx.PG().SetSyslogFilter(PHP_SYSLOG_FILTER_RAW)
		return true
	default:
		return false
	}
}

func onSetFacility(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
	var facility = newValue
	_ = facility
	// todo 此处应根据字符串值设置不同的枚举值
	//PG__().SetSyslogFacility(0)
	return false
}

func onUpdateDefaultCharset(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
	if strings.ContainsAny(newValue, "\r\n\000") {
		return false
	}
	ctx.PG().SetDefaultCharset(newValue)
	if PhpInternalEncodingChanged != nil {
		PhpInternalEncodingChanged()
	}
	return true
}
func onUpdateDefaultMimeType(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
	if strings.ContainsAny(newValue, "\r\n\000") {
		return false
	}

	ctx.PG().SetDefaultMimetype(newValue)
	return true
}
func onUpdateInternalEncoding(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
	ctx.PG().SetInternalEncoding(newValue)
	if PhpInternalEncodingChanged != nil {
		PhpInternalEncodingChanged()
	}
	return true
}
func onUpdateInputEncoding(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
	ctx.PG().SetInputEncoding(newValue)
	if PhpInternalEncodingChanged != nil {
		PhpInternalEncodingChanged()
	}
	return true
}
func onUpdateOutputEncoding(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
	ctx.PG().SetOutputEncoding(newValue)
	if PhpInternalEncodingChanged != nil {
		PhpInternalEncodingChanged()
	}
	return true
}
func onUpdateErrorLog(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
	/* Only do the safemode/open_basedir check at runtime */
	if (stage == IniStageRuntime || stage == IniStageHtaccess) && newValue != "" && newValue != "syslog" {
		if ctx.PG().OpenBasedir() != "" && !PhpCheckOpenBasedir(ctx, newValue) {
			return false
		}
	}

	ctx.PG().SetErrorLog(newValue)
	return true
}

func onUpdateMailLog(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
	/* Only do the safemode/open_basedir check at runtime */
	if (stage == IniStageRuntime || stage == IniStageHtaccess) && newValue != "" {
		if ctx.PG().OpenBasedir() != "" && !PhpCheckOpenBasedir(ctx, newValue) {
			return false
		}
	}
	ctx.PG().SetMailLog(newValue)
	return true
}
func onChangeMailForceExtra(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
	/* Don't allow changing it in htaccess */
	if stage == IniStageHtaccess {
		return false
	}
	return true
}

func onUpdateTimeout(ctx *Context, entry *IniEntry, newValue string, _ bool, stage IniStage) bool {
	//i := ParseLong10(newValue)

	//if stage == IniStageStartup {
	//	/* Don't set a timeout on startup, only per-request */
	//	ctx.EG().SetTimeoutSeconds(i)
	//	return true
	//}
	//ZendUnsetTimeout(ctx)
	//ctx.EG().SetTimeoutSeconds(i)
	//if stage != IniStageDeactivate {
	//	/*
	//	 * If we're restoring INI values, we shouldn't reset the timer.
	//	 * Otherwise, the timer is active when PHP is idle, such as the
	//	 * the CLI web server or CGI. Running a script will re-activate
	//	 * the timeout, so it's not needed to do so at script end.
	//	 */
	//	ZendSetTimeout(ctx, ctx.EG().TimeoutSeconds(), 0)
	//}
	return true
}

func onChangeBrowscap(ctx *Context, entry *IniEntry, newValue string, hasValue bool, stage IniStage) bool {
	if stage == IniStageStartup {
		/* value handled in browscap.c's MINIT */
		return true
	} else if stage == IniStageActivate {
		// todo 待迁移 @see: PHP_INI_MH(OnChangeBrowscap)
	}
	return false
}

func onUpdateErrorReporting(ctx *Context, entry *IniEntry, newValue string, hasValue bool, stage IniStage) bool {
	if !hasValue {
		ctx.EG().SetErrorReporting(perr.E_ALL &^ perr.E_NOTICE &^ perr.E_STRICT &^ perr.E_DEPRECATED)
	} else {
		intVal := ParseLong10(newValue)
		ctx.EG().SetErrorReporting(perr.ErrorType(intVal))
	}
	return true
}

func onUpdateAssertions(ctx *Context, entry *IniEntry, newValue string, hasValue bool, stage IniStage) bool {
	if !hasValue {
		return true
	}

	//assertions := ctx.EG().Assertions()
	//val, ok := ParseLongWithUnit(newValue)
	//if !ok {
	//	return false
	//}
	//if stage != IniStageStartup && stage != IniStageShutdown && assertions != val && (assertions < 0 || val < 0) {
	//	Error(ctx, perr.E_WARNING, "zend.assertions may be completely enabled or disabled only in php.ini")
	//	return false
	//}
	//ctx.EG().SetAssertions(val)
	return true
}
