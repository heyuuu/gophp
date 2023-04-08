package core

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
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
type PhpCoreGlobals struct {
	implicit_flush              types.ZendBool
	output_buffering            zend.ZendLong
	enable_dl                   types.ZendBool
	output_handler              *byte
	unserialize_callback_func   *byte
	serialize_precision         zend.ZendLong
	memory_limit                zend.ZendLong
	max_input_time              zend.ZendLong
	track_errors                types.ZendBool
	display_errors              types.ZendBool
	display_startup_errors      types.ZendBool
	log_errors                  types.ZendBool
	log_errors_max_len          zend.ZendLong
	ignore_repeated_errors      types.ZendBool
	ignore_repeated_source      types.ZendBool
	report_memleaks             types.ZendBool
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
	rfc1867_protected_variables types.Array
	connection_status           short
	ignore_user_abort           types.ZendBool
	header_is_being_sent        uint8
	tick_functions              zend.ZendLlist
	http_globals                []types.Zval
	expose_php                  types.ZendBool
	register_argc_argv          types.ZendBool
	auto_globals_jit            types.ZendBool
	docref_root                 *byte
	docref_ext                  *byte
	html_errors                 types.ZendBool
	xmlrpc_errors               types.ZendBool
	xmlrpc_error_number         zend.ZendLong
	activated_auto_globals      []types.ZendBool
	modules_activated           types.ZendBool
	file_uploads                types.ZendBool
	during_request_startup      types.ZendBool
	allow_url_fopen             types.ZendBool
	enable_post_data_reading    types.ZendBool
	report_zend_debug           types.ZendBool
	last_error_type             int
	last_error_message          *byte
	last_error_file             *byte
	last_error_lineno           int
	php_sys_temp_dir            *byte
	disable_functions           *byte
	disable_classes             *byte
	allow_url_include           types.ZendBool
	max_input_nesting_level     zend.ZendLong
	max_input_vars              zend.ZendLong
	in_user_include             types.ZendBool
	user_ini_filename           *byte
	user_ini_cache_ttl          zend.ZendLong
	request_order               *byte
	mail_x_header               types.ZendBool
	mail_log                    *byte
	in_error_log                types.ZendBool
	syslog_facility             zend.ZendLong
	syslog_ident                *byte
	have_called_openlog         types.ZendBool
	syslog_filter               zend.ZendLong
}

func (this *PhpCoreGlobals) GetImplicitFlush() types.ZendBool     { return this.implicit_flush }
func (this *PhpCoreGlobals) GetOutputBuffering() zend.ZendLong    { return this.output_buffering }
func (this *PhpCoreGlobals) GetEnableDl() types.ZendBool          { return this.enable_dl }
func (this *PhpCoreGlobals) GetOutputHandler() *byte              { return this.output_handler }
func (this *PhpCoreGlobals) GetUnserializeCallbackFunc() *byte    { return this.unserialize_callback_func }
func (this *PhpCoreGlobals) GetSerializePrecision() zend.ZendLong { return this.serialize_precision }
func (this *PhpCoreGlobals) GetMaxInputTime() zend.ZendLong       { return this.max_input_time }
func (this *PhpCoreGlobals) GetTrackErrors() types.ZendBool       { return this.track_errors }
func (this *PhpCoreGlobals) GetDisplayErrors() types.ZendBool     { return this.display_errors }
func (this *PhpCoreGlobals) GetDisplayStartupErrors() types.ZendBool {
	return this.display_startup_errors
}
func (this *PhpCoreGlobals) GetLogErrors() types.ZendBool      { return this.log_errors }
func (this *PhpCoreGlobals) GetLogErrorsMaxLen() zend.ZendLong { return this.log_errors_max_len }
func (this *PhpCoreGlobals) GetIgnoreRepeatedErrors() types.ZendBool {
	return this.ignore_repeated_errors
}
func (this *PhpCoreGlobals) GetIgnoreRepeatedSource() types.ZendBool {
	return this.ignore_repeated_source
}
func (this *PhpCoreGlobals) GetReportMemleaks() types.ZendBool   { return this.report_memleaks }
func (this *PhpCoreGlobals) GetErrorLog() *byte                  { return this.error_log }
func (this *PhpCoreGlobals) GetDocRoot() *byte                   { return this.doc_root }
func (this *PhpCoreGlobals) GetUserDir() *byte                   { return this.user_dir }
func (this *PhpCoreGlobals) GetIncludePath() *byte               { return this.include_path }
func (this *PhpCoreGlobals) GetOpenBasedir() *byte               { return this.open_basedir }
func (this *PhpCoreGlobals) GetExtensionDir() *byte              { return this.extension_dir }
func (this *PhpCoreGlobals) GetPhpBinary() *byte                 { return this.php_binary }
func (this *PhpCoreGlobals) GetSysTempDir() *byte                { return this.sys_temp_dir }
func (this *PhpCoreGlobals) GetUploadTmpDir() *byte              { return this.upload_tmp_dir }
func (this *PhpCoreGlobals) GetUploadMaxFilesize() zend.ZendLong { return this.upload_max_filesize }
func (this *PhpCoreGlobals) GetErrorAppendString() *byte         { return this.error_append_string }
func (this *PhpCoreGlobals) GetErrorPrependString() *byte        { return this.error_prepend_string }
func (this *PhpCoreGlobals) GetAutoPrependFile() *byte           { return this.auto_prepend_file }
func (this *PhpCoreGlobals) GetAutoAppendFile() *byte            { return this.auto_append_file }
func (this *PhpCoreGlobals) GetInputEncoding() *byte             { return this.input_encoding }
func (this *PhpCoreGlobals) GetInternalEncoding() *byte          { return this.internal_encoding }
func (this *PhpCoreGlobals) GetOutputEncoding() *byte            { return this.output_encoding }
func (this *PhpCoreGlobals) GetArgSeparator() ArgSeparators      { return this.arg_separator }
func (this *PhpCoreGlobals) GetVariablesOrder() *byte            { return this.variables_order }
func (this *PhpCoreGlobals) GetIgnoreUserAbort() types.ZendBool  { return this.ignore_user_abort }
func (this *PhpCoreGlobals) GetExposePhp() types.ZendBool        { return this.expose_php }
func (this *PhpCoreGlobals) GetRegisterArgcArgv() types.ZendBool { return this.register_argc_argv }
func (this *PhpCoreGlobals) GetAutoGlobalsJit() types.ZendBool   { return this.auto_globals_jit }
func (this *PhpCoreGlobals) GetDocrefRoot() *byte                { return this.docref_root }
func (this *PhpCoreGlobals) GetDocrefExt() *byte                 { return this.docref_ext }
func (this *PhpCoreGlobals) GetHtmlErrors() types.ZendBool       { return this.html_errors }
func (this *PhpCoreGlobals) GetXmlrpcErrors() types.ZendBool     { return this.xmlrpc_errors }
func (this *PhpCoreGlobals) GetXmlrpcErrorNumber() zend.ZendLong { return this.xmlrpc_error_number }
func (this *PhpCoreGlobals) GetFileUploads() types.ZendBool      { return this.file_uploads }
func (this *PhpCoreGlobals) GetAllowUrlFopen() types.ZendBool    { return this.allow_url_fopen }
func (this *PhpCoreGlobals) GetEnablePostDataReading() types.ZendBool {
	return this.enable_post_data_reading
}
func (this *PhpCoreGlobals) GetReportZendDebug() types.ZendBool { return this.report_zend_debug }
func (this *PhpCoreGlobals) GetLastErrorMessage() *byte         { return this.last_error_message }
func (this *PhpCoreGlobals) GetLastErrorFile() *byte            { return this.last_error_file }
func (this *PhpCoreGlobals) GetDisableFunctions() *byte         { return this.disable_functions }
func (this *PhpCoreGlobals) GetDisableClasses() *byte           { return this.disable_classes }
func (this *PhpCoreGlobals) GetAllowUrlInclude() types.ZendBool { return this.allow_url_include }
func (this *PhpCoreGlobals) GetMaxInputNestingLevel() zend.ZendLong {
	return this.max_input_nesting_level
}
func (this *PhpCoreGlobals) GetMaxInputVars() zend.ZendLong    { return this.max_input_vars }
func (this *PhpCoreGlobals) GetUserIniFilename() *byte         { return this.user_ini_filename }
func (this *PhpCoreGlobals) GetUserIniCacheTtl() zend.ZendLong { return this.user_ini_cache_ttl }
func (this *PhpCoreGlobals) GetRequestOrder() *byte            { return this.request_order }
func (this *PhpCoreGlobals) GetMailXHeader() types.ZendBool    { return this.mail_x_header }
func (this *PhpCoreGlobals) GetMailLog() *byte                 { return this.mail_log }
func (this *PhpCoreGlobals) GetSyslogFacility() zend.ZendLong  { return this.syslog_facility }
func (this *PhpCoreGlobals) GetSyslogIdent() *byte             { return this.syslog_ident }
func (this *PhpCoreGlobals) GetSyslogFilter() zend.ZendLong    { return this.syslog_filter }
