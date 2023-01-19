// <<generate>>

package core

import (
	"sik/zend"
)

/**
 * ArgSeparators
 */
type ArgSeparators struct {
	output *byte
	input  *byte
}

func (this ArgSeparators) GetOutput() *byte       { return this.output }
func (this *ArgSeparators) SetOutput(value *byte) { this.output = value }
func (this ArgSeparators) GetInput() *byte        { return this.input }
func (this *ArgSeparators) SetInput(value *byte)  { this.input = value }

/**
 * PhpCoreGlobals
 */
type PhpCoreGlobals struct {
	implicit_flush              zend.ZendBool
	output_buffering            zend.ZendLong
	enable_dl                   zend.ZendBool
	output_handler              *byte
	unserialize_callback_func   *byte
	serialize_precision         zend.ZendLong
	memory_limit                zend.ZendLong
	max_input_time              zend.ZendLong
	track_errors                zend.ZendBool
	display_errors              zend.ZendBool
	display_startup_errors      zend.ZendBool
	log_errors                  zend.ZendBool
	log_errors_max_len          zend.ZendLong
	ignore_repeated_errors      zend.ZendBool
	ignore_repeated_source      zend.ZendBool
	report_memleaks             zend.ZendBool
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
	rfc1867_protected_variables zend.HashTable
	connection_status           short
	ignore_user_abort           zend.ZendBool
	header_is_being_sent        uint8
	tick_functions              zend.ZendLlist
	http_globals                []zend.Zval
	expose_php                  zend.ZendBool
	register_argc_argv          zend.ZendBool
	auto_globals_jit            zend.ZendBool
	docref_root                 *byte
	docref_ext                  *byte
	html_errors                 zend.ZendBool
	xmlrpc_errors               zend.ZendBool
	xmlrpc_error_number         zend.ZendLong
	activated_auto_globals      []zend.ZendBool
	modules_activated           zend.ZendBool
	file_uploads                zend.ZendBool
	during_request_startup      zend.ZendBool
	allow_url_fopen             zend.ZendBool
	enable_post_data_reading    zend.ZendBool
	report_zend_debug           zend.ZendBool
	last_error_type             int
	last_error_message          *byte
	last_error_file             *byte
	last_error_lineno           int
	php_sys_temp_dir            *byte
	disable_functions           *byte
	disable_classes             *byte
	allow_url_include           zend.ZendBool
	max_input_nesting_level     zend.ZendLong
	max_input_vars              zend.ZendLong
	in_user_include             zend.ZendBool
	user_ini_filename           *byte
	user_ini_cache_ttl          zend.ZendLong
	request_order               *byte
	mail_x_header               zend.ZendBool
	mail_log                    *byte
	in_error_log                zend.ZendBool
	syslog_facility             zend.ZendLong
	syslog_ident                *byte
	have_called_openlog         zend.ZendBool
	syslog_filter               zend.ZendLong
}

func (this PhpCoreGlobals) GetImplicitFlush() zend.ZendBool         { return this.implicit_flush }
func (this *PhpCoreGlobals) SetImplicitFlush(value zend.ZendBool)   { this.implicit_flush = value }
func (this PhpCoreGlobals) GetOutputBuffering() zend.ZendLong       { return this.output_buffering }
func (this *PhpCoreGlobals) SetOutputBuffering(value zend.ZendLong) { this.output_buffering = value }
func (this PhpCoreGlobals) GetEnableDl() zend.ZendBool              { return this.enable_dl }
func (this *PhpCoreGlobals) SetEnableDl(value zend.ZendBool)        { this.enable_dl = value }
func (this PhpCoreGlobals) GetOutputHandler() *byte                 { return this.output_handler }
func (this *PhpCoreGlobals) SetOutputHandler(value *byte)           { this.output_handler = value }
func (this PhpCoreGlobals) GetUnserializeCallbackFunc() *byte       { return this.unserialize_callback_func }
func (this *PhpCoreGlobals) SetUnserializeCallbackFunc(value *byte) {
	this.unserialize_callback_func = value
}
func (this PhpCoreGlobals) GetSerializePrecision() zend.ZendLong { return this.serialize_precision }
func (this *PhpCoreGlobals) SetSerializePrecision(value zend.ZendLong) {
	this.serialize_precision = value
}
func (this PhpCoreGlobals) GetMemoryLimit() zend.ZendLong         { return this.memory_limit }
func (this *PhpCoreGlobals) SetMemoryLimit(value zend.ZendLong)   { this.memory_limit = value }
func (this PhpCoreGlobals) GetMaxInputTime() zend.ZendLong        { return this.max_input_time }
func (this *PhpCoreGlobals) SetMaxInputTime(value zend.ZendLong)  { this.max_input_time = value }
func (this PhpCoreGlobals) GetTrackErrors() zend.ZendBool         { return this.track_errors }
func (this *PhpCoreGlobals) SetTrackErrors(value zend.ZendBool)   { this.track_errors = value }
func (this PhpCoreGlobals) GetDisplayErrors() zend.ZendBool       { return this.display_errors }
func (this *PhpCoreGlobals) SetDisplayErrors(value zend.ZendBool) { this.display_errors = value }
func (this PhpCoreGlobals) GetDisplayStartupErrors() zend.ZendBool {
	return this.display_startup_errors
}
func (this *PhpCoreGlobals) SetDisplayStartupErrors(value zend.ZendBool) {
	this.display_startup_errors = value
}
func (this PhpCoreGlobals) GetLogErrors() zend.ZendBool             { return this.log_errors }
func (this *PhpCoreGlobals) SetLogErrors(value zend.ZendBool)       { this.log_errors = value }
func (this PhpCoreGlobals) GetLogErrorsMaxLen() zend.ZendLong       { return this.log_errors_max_len }
func (this *PhpCoreGlobals) SetLogErrorsMaxLen(value zend.ZendLong) { this.log_errors_max_len = value }
func (this PhpCoreGlobals) GetIgnoreRepeatedErrors() zend.ZendBool {
	return this.ignore_repeated_errors
}
func (this *PhpCoreGlobals) SetIgnoreRepeatedErrors(value zend.ZendBool) {
	this.ignore_repeated_errors = value
}
func (this PhpCoreGlobals) GetIgnoreRepeatedSource() zend.ZendBool {
	return this.ignore_repeated_source
}
func (this *PhpCoreGlobals) SetIgnoreRepeatedSource(value zend.ZendBool) {
	this.ignore_repeated_source = value
}
func (this PhpCoreGlobals) GetReportMemleaks() zend.ZendBool       { return this.report_memleaks }
func (this *PhpCoreGlobals) SetReportMemleaks(value zend.ZendBool) { this.report_memleaks = value }
func (this PhpCoreGlobals) GetErrorLog() *byte                     { return this.error_log }
func (this *PhpCoreGlobals) SetErrorLog(value *byte)               { this.error_log = value }
func (this PhpCoreGlobals) GetDocRoot() *byte                      { return this.doc_root }
func (this *PhpCoreGlobals) SetDocRoot(value *byte)                { this.doc_root = value }
func (this PhpCoreGlobals) GetUserDir() *byte                      { return this.user_dir }
func (this *PhpCoreGlobals) SetUserDir(value *byte)                { this.user_dir = value }
func (this PhpCoreGlobals) GetIncludePath() *byte                  { return this.include_path }
func (this *PhpCoreGlobals) SetIncludePath(value *byte)            { this.include_path = value }
func (this PhpCoreGlobals) GetOpenBasedir() *byte                  { return this.open_basedir }
func (this *PhpCoreGlobals) SetOpenBasedir(value *byte)            { this.open_basedir = value }
func (this PhpCoreGlobals) GetExtensionDir() *byte                 { return this.extension_dir }
func (this *PhpCoreGlobals) SetExtensionDir(value *byte)           { this.extension_dir = value }
func (this PhpCoreGlobals) GetPhpBinary() *byte                    { return this.php_binary }
func (this *PhpCoreGlobals) SetPhpBinary(value *byte)              { this.php_binary = value }
func (this PhpCoreGlobals) GetSysTempDir() *byte                   { return this.sys_temp_dir }
func (this *PhpCoreGlobals) SetSysTempDir(value *byte)             { this.sys_temp_dir = value }
func (this PhpCoreGlobals) GetUploadTmpDir() *byte                 { return this.upload_tmp_dir }
func (this *PhpCoreGlobals) SetUploadTmpDir(value *byte)           { this.upload_tmp_dir = value }
func (this PhpCoreGlobals) GetUploadMaxFilesize() zend.ZendLong    { return this.upload_max_filesize }
func (this *PhpCoreGlobals) SetUploadMaxFilesize(value zend.ZendLong) {
	this.upload_max_filesize = value
}
func (this PhpCoreGlobals) GetErrorAppendString() *byte          { return this.error_append_string }
func (this *PhpCoreGlobals) SetErrorAppendString(value *byte)    { this.error_append_string = value }
func (this PhpCoreGlobals) GetErrorPrependString() *byte         { return this.error_prepend_string }
func (this *PhpCoreGlobals) SetErrorPrependString(value *byte)   { this.error_prepend_string = value }
func (this PhpCoreGlobals) GetAutoPrependFile() *byte            { return this.auto_prepend_file }
func (this *PhpCoreGlobals) SetAutoPrependFile(value *byte)      { this.auto_prepend_file = value }
func (this PhpCoreGlobals) GetAutoAppendFile() *byte             { return this.auto_append_file }
func (this *PhpCoreGlobals) SetAutoAppendFile(value *byte)       { this.auto_append_file = value }
func (this PhpCoreGlobals) GetInputEncoding() *byte              { return this.input_encoding }
func (this *PhpCoreGlobals) SetInputEncoding(value *byte)        { this.input_encoding = value }
func (this PhpCoreGlobals) GetInternalEncoding() *byte           { return this.internal_encoding }
func (this *PhpCoreGlobals) SetInternalEncoding(value *byte)     { this.internal_encoding = value }
func (this PhpCoreGlobals) GetOutputEncoding() *byte             { return this.output_encoding }
func (this *PhpCoreGlobals) SetOutputEncoding(value *byte)       { this.output_encoding = value }
func (this PhpCoreGlobals) GetArgSeparator() ArgSeparators       { return this.arg_separator }
func (this *PhpCoreGlobals) SetArgSeparator(value ArgSeparators) { this.arg_separator = value }
func (this PhpCoreGlobals) GetVariablesOrder() *byte             { return this.variables_order }
func (this *PhpCoreGlobals) SetVariablesOrder(value *byte)       { this.variables_order = value }
func (this PhpCoreGlobals) GetRfc1867ProtectedVariables() zend.HashTable {
	return this.rfc1867_protected_variables
}
func (this *PhpCoreGlobals) SetRfc1867ProtectedVariables(value zend.HashTable) {
	this.rfc1867_protected_variables = value
}
func (this PhpCoreGlobals) GetConnectionStatus() short               { return this.connection_status }
func (this *PhpCoreGlobals) SetConnectionStatus(value short)         { this.connection_status = value }
func (this PhpCoreGlobals) GetIgnoreUserAbort() zend.ZendBool        { return this.ignore_user_abort }
func (this *PhpCoreGlobals) SetIgnoreUserAbort(value zend.ZendBool)  { this.ignore_user_abort = value }
func (this PhpCoreGlobals) GetHeaderIsBeingSent() uint8              { return this.header_is_being_sent }
func (this *PhpCoreGlobals) SetHeaderIsBeingSent(value uint8)        { this.header_is_being_sent = value }
func (this PhpCoreGlobals) GetTickFunctions() zend.ZendLlist         { return this.tick_functions }
func (this *PhpCoreGlobals) SetTickFunctions(value zend.ZendLlist)   { this.tick_functions = value }
func (this PhpCoreGlobals) GetHttpGlobals() []zend.Zval              { return this.http_globals }
func (this *PhpCoreGlobals) SetHttpGlobals(value []zend.Zval)        { this.http_globals = value }
func (this PhpCoreGlobals) GetExposePhp() zend.ZendBool              { return this.expose_php }
func (this *PhpCoreGlobals) SetExposePhp(value zend.ZendBool)        { this.expose_php = value }
func (this PhpCoreGlobals) GetRegisterArgcArgv() zend.ZendBool       { return this.register_argc_argv }
func (this *PhpCoreGlobals) SetRegisterArgcArgv(value zend.ZendBool) { this.register_argc_argv = value }
func (this PhpCoreGlobals) GetAutoGlobalsJit() zend.ZendBool         { return this.auto_globals_jit }
func (this *PhpCoreGlobals) SetAutoGlobalsJit(value zend.ZendBool)   { this.auto_globals_jit = value }
func (this PhpCoreGlobals) GetDocrefRoot() *byte                     { return this.docref_root }
func (this *PhpCoreGlobals) SetDocrefRoot(value *byte)               { this.docref_root = value }
func (this PhpCoreGlobals) GetDocrefExt() *byte                      { return this.docref_ext }
func (this *PhpCoreGlobals) SetDocrefExt(value *byte)                { this.docref_ext = value }
func (this PhpCoreGlobals) GetHtmlErrors() zend.ZendBool             { return this.html_errors }
func (this *PhpCoreGlobals) SetHtmlErrors(value zend.ZendBool)       { this.html_errors = value }
func (this PhpCoreGlobals) GetXmlrpcErrors() zend.ZendBool           { return this.xmlrpc_errors }
func (this *PhpCoreGlobals) SetXmlrpcErrors(value zend.ZendBool)     { this.xmlrpc_errors = value }
func (this PhpCoreGlobals) GetXmlrpcErrorNumber() zend.ZendLong      { return this.xmlrpc_error_number }
func (this *PhpCoreGlobals) SetXmlrpcErrorNumber(value zend.ZendLong) {
	this.xmlrpc_error_number = value
}
func (this PhpCoreGlobals) GetActivatedAutoGlobals() []zend.ZendBool {
	return this.activated_auto_globals
}
func (this *PhpCoreGlobals) SetActivatedAutoGlobals(value []zend.ZendBool) {
	this.activated_auto_globals = value
}
func (this PhpCoreGlobals) GetModulesActivated() zend.ZendBool       { return this.modules_activated }
func (this *PhpCoreGlobals) SetModulesActivated(value zend.ZendBool) { this.modules_activated = value }
func (this PhpCoreGlobals) GetFileUploads() zend.ZendBool            { return this.file_uploads }
func (this *PhpCoreGlobals) SetFileUploads(value zend.ZendBool)      { this.file_uploads = value }
func (this PhpCoreGlobals) GetDuringRequestStartup() zend.ZendBool {
	return this.during_request_startup
}
func (this *PhpCoreGlobals) SetDuringRequestStartup(value zend.ZendBool) {
	this.during_request_startup = value
}
func (this PhpCoreGlobals) GetAllowUrlFopen() zend.ZendBool       { return this.allow_url_fopen }
func (this *PhpCoreGlobals) SetAllowUrlFopen(value zend.ZendBool) { this.allow_url_fopen = value }
func (this PhpCoreGlobals) GetEnablePostDataReading() zend.ZendBool {
	return this.enable_post_data_reading
}
func (this *PhpCoreGlobals) SetEnablePostDataReading(value zend.ZendBool) {
	this.enable_post_data_reading = value
}
func (this PhpCoreGlobals) GetReportZendDebug() zend.ZendBool       { return this.report_zend_debug }
func (this *PhpCoreGlobals) SetReportZendDebug(value zend.ZendBool) { this.report_zend_debug = value }
func (this PhpCoreGlobals) GetLastErrorType() int                   { return this.last_error_type }
func (this *PhpCoreGlobals) SetLastErrorType(value int)             { this.last_error_type = value }
func (this PhpCoreGlobals) GetLastErrorMessage() *byte              { return this.last_error_message }
func (this *PhpCoreGlobals) SetLastErrorMessage(value *byte)        { this.last_error_message = value }
func (this PhpCoreGlobals) GetLastErrorFile() *byte                 { return this.last_error_file }
func (this *PhpCoreGlobals) SetLastErrorFile(value *byte)           { this.last_error_file = value }
func (this PhpCoreGlobals) GetLastErrorLineno() int                 { return this.last_error_lineno }
func (this *PhpCoreGlobals) SetLastErrorLineno(value int)           { this.last_error_lineno = value }
func (this PhpCoreGlobals) GetPhpSysTempDir() *byte                 { return this.php_sys_temp_dir }
func (this *PhpCoreGlobals) SetPhpSysTempDir(value *byte)           { this.php_sys_temp_dir = value }
func (this PhpCoreGlobals) GetDisableFunctions() *byte              { return this.disable_functions }
func (this *PhpCoreGlobals) SetDisableFunctions(value *byte)        { this.disable_functions = value }
func (this PhpCoreGlobals) GetDisableClasses() *byte                { return this.disable_classes }
func (this *PhpCoreGlobals) SetDisableClasses(value *byte)          { this.disable_classes = value }
func (this PhpCoreGlobals) GetAllowUrlInclude() zend.ZendBool       { return this.allow_url_include }
func (this *PhpCoreGlobals) SetAllowUrlInclude(value zend.ZendBool) { this.allow_url_include = value }
func (this PhpCoreGlobals) GetMaxInputNestingLevel() zend.ZendLong {
	return this.max_input_nesting_level
}
func (this *PhpCoreGlobals) SetMaxInputNestingLevel(value zend.ZendLong) {
	this.max_input_nesting_level = value
}
func (this PhpCoreGlobals) GetMaxInputVars() zend.ZendLong          { return this.max_input_vars }
func (this *PhpCoreGlobals) SetMaxInputVars(value zend.ZendLong)    { this.max_input_vars = value }
func (this PhpCoreGlobals) GetInUserInclude() zend.ZendBool         { return this.in_user_include }
func (this *PhpCoreGlobals) SetInUserInclude(value zend.ZendBool)   { this.in_user_include = value }
func (this PhpCoreGlobals) GetUserIniFilename() *byte               { return this.user_ini_filename }
func (this *PhpCoreGlobals) SetUserIniFilename(value *byte)         { this.user_ini_filename = value }
func (this PhpCoreGlobals) GetUserIniCacheTtl() zend.ZendLong       { return this.user_ini_cache_ttl }
func (this *PhpCoreGlobals) SetUserIniCacheTtl(value zend.ZendLong) { this.user_ini_cache_ttl = value }
func (this PhpCoreGlobals) GetRequestOrder() *byte                  { return this.request_order }
func (this *PhpCoreGlobals) SetRequestOrder(value *byte)            { this.request_order = value }
func (this PhpCoreGlobals) GetMailXHeader() zend.ZendBool           { return this.mail_x_header }
func (this *PhpCoreGlobals) SetMailXHeader(value zend.ZendBool)     { this.mail_x_header = value }
func (this PhpCoreGlobals) GetMailLog() *byte                       { return this.mail_log }
func (this *PhpCoreGlobals) SetMailLog(value *byte)                 { this.mail_log = value }
func (this PhpCoreGlobals) GetInErrorLog() zend.ZendBool            { return this.in_error_log }
func (this *PhpCoreGlobals) SetInErrorLog(value zend.ZendBool)      { this.in_error_log = value }
func (this PhpCoreGlobals) GetSyslogFacility() zend.ZendLong        { return this.syslog_facility }
func (this *PhpCoreGlobals) SetSyslogFacility(value zend.ZendLong)  { this.syslog_facility = value }
func (this PhpCoreGlobals) GetSyslogIdent() *byte                   { return this.syslog_ident }
func (this *PhpCoreGlobals) SetSyslogIdent(value *byte)             { this.syslog_ident = value }
func (this PhpCoreGlobals) GetHaveCalledOpenlog() zend.ZendBool     { return this.have_called_openlog }
func (this *PhpCoreGlobals) SetHaveCalledOpenlog(value zend.ZendBool) {
	this.have_called_openlog = value
}
func (this PhpCoreGlobals) GetSyslogFilter() zend.ZendLong       { return this.syslog_filter }
func (this *PhpCoreGlobals) SetSyslogFilter(value zend.ZendLong) { this.syslog_filter = value }
