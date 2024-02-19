package php

import "github.com/heyuuu/gophp/php/perr"

const DefaultCharset = "UTF-8"
const DefaultMimetype = "text/html"

var MainIniEntries = []*IniEntryDef{
	//NewIniEntryDef("highlight.comment", IniAll).Value(HlCommentColor),
	//NewIniEntryDef("highlight.default", IniAll).Value(HlDefaultColor),
	//NewIniEntryDef("highlight.html", IniAll).Value(HlHtmlColor),
	//NewIniEntryDef("highlight.keyword", IniAll).Value(HlKeywordColor),
	//NewIniEntryDef("highlight.string", IniAll).Value(HlStringColor),
	//NewIniEntryDef("display_errors", IniAll).Value("1"),
	//NewIniEntryDef("display_startup_errors", IniAll).Value("0"),
	//NewIniEntryDef("expose_php", IniSystem).Value("1"),
	//NewIniEntryDef("docref_root", IniAll).Value(""),
	//NewIniEntryDef("docref_ext", IniAll).Value(""),
	//NewIniEntryDef("html_errors", IniAll).Value("1"),
	//NewIniEntryDef("xmlrpc_errors", IniSystem).Value("0"),
	//NewIniEntryDef("xmlrpc_error_number", IniAll).Value("0"),
	//NewIniEntryDef("max_input_time", IniSystem|IniPerDir).Value("-1"),
	//NewIniEntryDef("ignore_user_abort", IniAll).Value("0"),
	//NewIniEntryDef("implicit_flush", IniAll).Value("0"),
	//NewIniEntryDef("log_errors", IniAll).Value("0"),
	//NewIniEntryDef("log_errors_max_len", IniAll).Value("1024"),
	//NewIniEntryDef("ignore_repeated_errors", IniAll).Value("0"),
	//NewIniEntryDef("ignore_repeated_source", IniAll).Value("0"),
	//NewIniEntryDef("report_memleaks", IniAll).Value("1"),
	//NewIniEntryDef("report_zend_debug", IniAll).Value("1"),
	//NewIniEntryDef("output_buffering", IniPerDir|IniSystem).Value("0"),
	//NewIniEntryDef("output_handler", IniPerDir|IniSystem),
	//NewIniEntryDef("register_argc_argv", IniPerDir|IniSystem).Value("1"),
	//NewIniEntryDef("auto_globals_jit", IniPerDir|IniSystem).Value("1"),
	//NewIniEntryDef("short_open_tag", IniSystem|IniPerDir).Value("1"),
	//NewIniEntryDef("track_errors", IniAll).Value("0"),
	//NewIniEntryDef("unserialize_callback_func", IniAll),
	//NewIniEntryDef("serialize_precision", IniAll).Value("-1"),
	//NewIniEntryDef("arg_separator.output", IniAll).Value("&"),
	//NewIniEntryDef("arg_separator.input", IniSystem|IniPerDir).Value("&"),
	//NewIniEntryDef("auto_append_file", IniSystem|IniPerDir),
	//NewIniEntryDef("auto_prepend_file", IniSystem|IniPerDir),
	//NewIniEntryDef("doc_root", IniSystem),
	//NewIniEntryDef("default_charset", IniAll).Value(DefaultCharset),
	//NewIniEntryDef("default_mimetype", IniAll).Value(DefaultMimetype),
	//NewIniEntryDef("internal_encoding", IniAll),
	//NewIniEntryDef("input_encoding", IniAll),
	//NewIniEntryDef("output_encoding", IniAll),
	//NewIniEntryDef("error_log", IniAll),
	////NewIniEntryDef("extension_dir", IniSystem).Value(PHP_EXTENSION_DIR),
	//NewIniEntryDef("sys_temp_dir", IniSystem),
	////NewIniEntryDef("include_path", IniAll).Value(PHP_INCLUDE_PATH),
	//NewIniEntryDef("max_execution_time", IniAll).Value("30"),
	//NewIniEntryDef("open_basedir", IniAll),
	//NewIniEntryDef("file_uploads", IniSystem).Value("1"),
	//NewIniEntryDef("upload_max_filesize", IniSystem|IniPerDir).Value("2M"),
	//NewIniEntryDef("post_max_size", IniSystem|IniPerDir).Value("8M"),
	//NewIniEntryDef("upload_tmp_dir", IniSystem),
	//NewIniEntryDef("max_input_nesting_level", IniSystem|IniPerDir).Value("64"),
	//NewIniEntryDef("max_input_vars", IniSystem|IniPerDir).Value("1000"),
	//NewIniEntryDef("user_dir", IniSystem),
	//NewIniEntryDef("variables_order", IniSystem|IniPerDir).Value("EGPCS"),
	//NewIniEntryDef("request_order", IniSystem|IniPerDir),
	//NewIniEntryDef("error_append_string", IniAll),
	//NewIniEntryDef("error_prepend_string", IniAll),
	//NewIniEntryDef("SMTP", IniAll).Value("localhost"),
	//NewIniEntryDef("smtp_port", IniAll).Value("25"),
	//NewIniEntryDef("mail.add_x_header", IniSystem|IniPerDir).Value("0"),
	//NewIniEntryDef("mail.log", IniSystem|IniPerDir),
	//NewIniEntryDef("browscap", IniSystem),
	//NewIniEntryDef("memory_limit", IniAll).Value("128M"),
	//NewIniEntryDef("precision", IniAll).Value("14"),
	//NewIniEntryDef("sendmail_from", IniAll),
	////NewIniEntryDef("sendmail_path", IniSystem).Value(DEFAULT_SENDMAIL_PATH),
	//NewIniEntryDef("mail.force_extra_parameters", IniSystem|IniPerDir),
	//NewIniEntryDef("disable_functions", IniSystem).Value(""),
	//NewIniEntryDef("disable_classes", IniSystem).Value(""),
	//NewIniEntryDef("max_file_uploads", IniSystem|IniPerDir).Value("20"),
	//NewIniEntryDef("allow_url_fopen", IniSystem).Value("1"),
	//NewIniEntryDef("allow_url_include", IniSystem).Value("0"),
	//NewIniEntryDef("enable_post_data_reading", IniSystem|IniPerDir).Value("1"),
	//// notice: realpath 缓存功能已取消，获取 realpath 改为直接调用 go 标准库
	////*php.NewIniEntryDef("realpath_cache_size", IniSystem).Value("4096K").OnModifyLong(func(i int) { php.CWDG__().SetRealpathCacheSizeLimit(i) }),
	////*php.NewIniEntryDef("realpath_cache_ttl", IniSystem).Value("120").OnModifyLong(func(i int) { php.CWDG__().SetRealpathCacheTtl(i) }),
	//NewIniEntryDef("user_ini.filename", IniSystem).Value(".user.ini"),
	//NewIniEntryDef("user_ini.cache_ttl", IniSystem).Value("300"),
	//NewIniEntryDef("hard_timeout", IniSystem).Value("2"),
	//NewIniEntryDef("syslog.facility", IniSystem).Value("LOG_USER"),
	//NewIniEntryDef("syslog.ident", IniSystem).Value("php"),
	//NewIniEntryDef("syslog.filter", IniAll).Value("no-ctrl"),

	// zend
	NewIniEntryDef("error_reporting", IniAll).OnModify(onUpdateErrorReporting),
	//NewIniEntryDef("zend.assertions", IniAll).Value("1"),
	//NewIniEntryDef("zend.signal_check", IniSystem).Value("0"),
	//NewIniEntryDef("zend.exception_ignore_args", IniAll).Value("0"),
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
