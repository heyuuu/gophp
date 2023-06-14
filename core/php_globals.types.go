package core

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

/**
 * ArgSeparators
 */
type ArgSeparators struct {
	output *byte
	input  *byte
}

func (this *ArgSeparators) GetOutput() *byte { return this.output }
func (this *ArgSeparators) GetInput() *byte  { return this.input }

/**
 * PhpCoreGlobals
 */
type LastError struct {
	Type    int
	Message string
	File    string
	Lineno  int
}

type PhpCoreGlobals struct {
	implicit_flush              bool
	output_buffering            zend.ZendLong
	enable_dl                   bool
	output_handler              *byte
	unserialize_callback_func   *byte
	serialize_precision         zend.ZendLong
	memory_limit                zend.ZendLong
	max_input_time              zend.ZendLong
	track_errors                bool
	display_errors              bool
	display_startup_errors      bool
	log_errors                  bool
	log_errors_max_len          zend.ZendLong
	ignore_repeated_errors      bool
	ignore_repeated_source      bool
	report_memleaks             bool
	error_log                   *byte
	doc_root                    *byte
	user_dir                    *byte
	include_path                *byte
	open_basedir                *byte
	extension_dir               *byte
	php_binary                  *byte
	sys_temp_dir                *byte
	upload_tmp_dir              *byte
	upload_max_filesize         zend.ZendLong
	error_append_string         *byte
	error_prepend_string        *byte
	auto_prepend_file           *byte
	auto_append_file            *byte
	input_encoding              *byte
	internal_encoding           *byte
	output_encoding             *byte
	arg_separator               ArgSeparators
	variables_order             *byte
	rfc1867_protected_variables *types.Array
	connection_status           int16
	ignore_user_abort           bool
	header_is_being_sent        uint8
	tick_functions              zend.ZendLlist
	http_globals                []types.Zval
	expose_php                  bool
	register_argc_argv          bool
	auto_globals_jit            bool
	docref_root                 *byte
	docref_ext                  *byte
	html_errors                 bool
	xmlrpc_errors               bool
	xmlrpc_error_number         zend.ZendLong
	activated_auto_globals      []bool
	modules_activated           bool
	file_uploads                bool
	during_request_startup      bool
	allow_url_fopen             bool
	enable_post_data_reading    bool
	report_zend_debug           bool
	lastError                   *LastError
	php_sys_temp_dir            *byte
	disable_functions           *byte
	disable_classes             *byte
	allow_url_include           bool
	max_input_nesting_level     zend.ZendLong
	max_input_vars              zend.ZendLong
	in_user_include             bool
	user_ini_filename           *byte
	user_ini_cache_ttl          zend.ZendLong
	request_order               *byte
	mail_x_header               bool
	mail_log                    *byte
	in_error_log                bool
	syslog_facility             zend.ZendLong
	syslog_ident                *byte
	have_called_openlog         bool
	syslog_filter               zend.ZendLong
}

// last error
func (pg *PhpCoreGlobals) LastError() *LastError { return pg.lastError }
func (pg *PhpCoreGlobals) AddLastError(typ int, message string, file string, lineno int) {
	pg.lastError = &LastError{Type: typ, Message: message, File: file, Lineno: lineno}
}
func (pg *PhpCoreGlobals) ClearLastError() { pg.lastError = nil }

func (pg *PhpCoreGlobals) GetImplicitFlush() bool               { return pg.implicit_flush }
func (pg *PhpCoreGlobals) GetOutputBuffering() zend.ZendLong    { return pg.output_buffering }
func (pg *PhpCoreGlobals) GetEnableDl() bool                    { return pg.enable_dl }
func (pg *PhpCoreGlobals) GetOutputHandler() *byte              { return pg.output_handler }
func (pg *PhpCoreGlobals) GetUnserializeCallbackFunc() *byte    { return pg.unserialize_callback_func }
func (pg *PhpCoreGlobals) GetSerializePrecision() zend.ZendLong { return pg.serialize_precision }
func (pg *PhpCoreGlobals) GetMaxInputTime() zend.ZendLong       { return pg.max_input_time }
func (pg *PhpCoreGlobals) GetTrackErrors() bool                 { return pg.track_errors }
func (pg *PhpCoreGlobals) GetDisplayErrors() bool               { return pg.display_errors }
func (pg *PhpCoreGlobals) GetDisplayStartupErrors() bool {
	return pg.display_startup_errors
}
func (pg *PhpCoreGlobals) GetLogErrors() bool                { return pg.log_errors }
func (pg *PhpCoreGlobals) GetLogErrorsMaxLen() zend.ZendLong { return pg.log_errors_max_len }
func (pg *PhpCoreGlobals) GetIgnoreRepeatedErrors() bool {
	return pg.ignore_repeated_errors
}
func (pg *PhpCoreGlobals) GetIgnoreRepeatedSource() bool {
	return pg.ignore_repeated_source
}
func (pg *PhpCoreGlobals) GetReportMemleaks() bool             { return pg.report_memleaks }
func (pg *PhpCoreGlobals) GetErrorLog() *byte                  { return pg.error_log }
func (pg *PhpCoreGlobals) GetDocRoot() *byte                   { return pg.doc_root }
func (pg *PhpCoreGlobals) GetUserDir() *byte                   { return pg.user_dir }
func (pg *PhpCoreGlobals) GetIncludePath() *byte               { return pg.include_path }
func (pg *PhpCoreGlobals) GetOpenBasedir() *byte               { return pg.open_basedir }
func (pg *PhpCoreGlobals) GetExtensionDir() *byte              { return pg.extension_dir }
func (pg *PhpCoreGlobals) GetPhpBinary() *byte                 { return pg.php_binary }
func (pg *PhpCoreGlobals) GetSysTempDir() *byte                { return pg.sys_temp_dir }
func (pg *PhpCoreGlobals) GetUploadTmpDir() *byte              { return pg.upload_tmp_dir }
func (pg *PhpCoreGlobals) GetUploadMaxFilesize() zend.ZendLong { return pg.upload_max_filesize }
func (pg *PhpCoreGlobals) GetErrorAppendString() *byte         { return pg.error_append_string }
func (pg *PhpCoreGlobals) GetErrorPrependString() *byte        { return pg.error_prepend_string }
func (pg *PhpCoreGlobals) GetAutoPrependFile() *byte           { return pg.auto_prepend_file }
func (pg *PhpCoreGlobals) GetAutoAppendFile() *byte            { return pg.auto_append_file }
func (pg *PhpCoreGlobals) GetInputEncoding() *byte             { return pg.input_encoding }
func (pg *PhpCoreGlobals) GetInternalEncoding() *byte          { return pg.internal_encoding }
func (pg *PhpCoreGlobals) GetOutputEncoding() *byte            { return pg.output_encoding }
func (pg *PhpCoreGlobals) GetArgSeparator() ArgSeparators      { return pg.arg_separator }
func (pg *PhpCoreGlobals) GetVariablesOrder() *byte            { return pg.variables_order }
func (pg *PhpCoreGlobals) GetIgnoreUserAbort() bool            { return pg.ignore_user_abort }
func (pg *PhpCoreGlobals) GetExposePhp() bool                  { return pg.expose_php }
func (pg *PhpCoreGlobals) GetRegisterArgcArgv() bool           { return pg.register_argc_argv }
func (pg *PhpCoreGlobals) GetAutoGlobalsJit() bool             { return pg.auto_globals_jit }
func (pg *PhpCoreGlobals) GetDocrefRoot() *byte                { return pg.docref_root }
func (pg *PhpCoreGlobals) GetDocrefExt() *byte                 { return pg.docref_ext }
func (pg *PhpCoreGlobals) GetHtmlErrors() bool                 { return pg.html_errors }
func (pg *PhpCoreGlobals) GetXmlrpcErrors() bool               { return pg.xmlrpc_errors }
func (pg *PhpCoreGlobals) GetXmlrpcErrorNumber() zend.ZendLong { return pg.xmlrpc_error_number }
func (pg *PhpCoreGlobals) GetFileUploads() bool                { return pg.file_uploads }
func (pg *PhpCoreGlobals) GetAllowUrlFopen() bool              { return pg.allow_url_fopen }
func (pg *PhpCoreGlobals) GetEnablePostDataReading() bool {
	return pg.enable_post_data_reading
}
func (pg *PhpCoreGlobals) GetReportZendDebug() bool   { return pg.report_zend_debug }
func (pg *PhpCoreGlobals) GetDisableFunctions() *byte { return pg.disable_functions }
func (pg *PhpCoreGlobals) GetDisableClasses() *byte   { return pg.disable_classes }
func (pg *PhpCoreGlobals) GetAllowUrlInclude() bool   { return pg.allow_url_include }
func (pg *PhpCoreGlobals) GetMaxInputNestingLevel() zend.ZendLong {
	return pg.max_input_nesting_level
}
func (pg *PhpCoreGlobals) GetMaxInputVars() zend.ZendLong    { return pg.max_input_vars }
func (pg *PhpCoreGlobals) GetUserIniFilename() *byte         { return pg.user_ini_filename }
func (pg *PhpCoreGlobals) GetUserIniCacheTtl() zend.ZendLong { return pg.user_ini_cache_ttl }
func (pg *PhpCoreGlobals) GetRequestOrder() *byte            { return pg.request_order }
func (pg *PhpCoreGlobals) GetMailXHeader() bool              { return pg.mail_x_header }
func (pg *PhpCoreGlobals) GetMailLog() *byte                 { return pg.mail_log }
func (pg *PhpCoreGlobals) GetSyslogFacility() zend.ZendLong  { return pg.syslog_facility }
func (pg *PhpCoreGlobals) GetSyslogIdent() *byte             { return pg.syslog_ident }
func (pg *PhpCoreGlobals) GetSyslogFilter() zend.ZendLong    { return pg.syslog_filter }
