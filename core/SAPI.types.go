// <<generate>>

package core

import (
	"sik/zend"
)

/**
 * SapiHeader
 */
type SapiHeader struct {
	header     *byte
	header_len int
}

func (this SapiHeader) GetHeader() *byte        { return this.header }
func (this *SapiHeader) SetHeader(value *byte)  { this.header = value }
func (this SapiHeader) GetHeaderLen() int       { return this.header_len }
func (this *SapiHeader) SetHeaderLen(value int) { this.header_len = value }

/**
 * SapiHeaders
 */
type SapiHeaders struct {
	headers                   zend.ZendLlist
	http_response_code        int
	send_default_content_type uint8
	mimetype                  *byte
	http_status_line          *byte
}

func (this SapiHeaders) GetHeaders() zend.ZendLlist       { return this.headers }
func (this *SapiHeaders) SetHeaders(value zend.ZendLlist) { this.headers = value }
func (this SapiHeaders) GetHttpResponseCode() int         { return this.http_response_code }
func (this *SapiHeaders) SetHttpResponseCode(value int)   { this.http_response_code = value }
func (this SapiHeaders) GetSendDefaultContentType() uint8 { return this.send_default_content_type }
func (this *SapiHeaders) SetSendDefaultContentType(value uint8) {
	this.send_default_content_type = value
}
func (this SapiHeaders) GetMimetype() *byte             { return this.mimetype }
func (this *SapiHeaders) SetMimetype(value *byte)       { this.mimetype = value }
func (this SapiHeaders) GetHttpStatusLine() *byte       { return this.http_status_line }
func (this *SapiHeaders) SetHttpStatusLine(value *byte) { this.http_status_line = value }

/**
 * SapiRequestInfo
 */
type SapiRequestInfo struct {
	request_method      *byte
	query_string        *byte
	cookie_data         *byte
	content_length      zend.ZendLong
	path_translated     *byte
	request_uri         *byte
	request_body        *PhpStream
	content_type        *byte
	headers_only        zend.ZendBool
	no_headers          zend.ZendBool
	headers_read        zend.ZendBool
	post_entry          *SapiPostEntry
	content_type_dup    *byte
	auth_user           *byte
	auth_password       *byte
	auth_digest         *byte
	argv0               *byte
	current_user        *byte
	current_user_length int
	argc                int
	argv                **byte
	proto_num           int
}

func (this SapiRequestInfo) GetRequestMethod() *byte               { return this.request_method }
func (this *SapiRequestInfo) SetRequestMethod(value *byte)         { this.request_method = value }
func (this SapiRequestInfo) GetQueryString() *byte                 { return this.query_string }
func (this *SapiRequestInfo) SetQueryString(value *byte)           { this.query_string = value }
func (this SapiRequestInfo) GetCookieData() *byte                  { return this.cookie_data }
func (this *SapiRequestInfo) SetCookieData(value *byte)            { this.cookie_data = value }
func (this SapiRequestInfo) GetContentLength() zend.ZendLong       { return this.content_length }
func (this *SapiRequestInfo) SetContentLength(value zend.ZendLong) { this.content_length = value }
func (this SapiRequestInfo) GetPathTranslated() *byte              { return this.path_translated }
func (this *SapiRequestInfo) SetPathTranslated(value *byte)        { this.path_translated = value }
func (this SapiRequestInfo) GetRequestUri() *byte                  { return this.request_uri }
func (this *SapiRequestInfo) SetRequestUri(value *byte)            { this.request_uri = value }
func (this SapiRequestInfo) GetRequestBody() *PhpStream            { return this.request_body }
func (this *SapiRequestInfo) SetRequestBody(value *PhpStream)      { this.request_body = value }
func (this SapiRequestInfo) GetContentType() *byte                 { return this.content_type }
func (this *SapiRequestInfo) SetContentType(value *byte)           { this.content_type = value }
func (this SapiRequestInfo) GetHeadersOnly() zend.ZendBool         { return this.headers_only }
func (this *SapiRequestInfo) SetHeadersOnly(value zend.ZendBool)   { this.headers_only = value }
func (this SapiRequestInfo) GetNoHeaders() zend.ZendBool           { return this.no_headers }
func (this *SapiRequestInfo) SetNoHeaders(value zend.ZendBool)     { this.no_headers = value }
func (this SapiRequestInfo) GetHeadersRead() zend.ZendBool         { return this.headers_read }
func (this *SapiRequestInfo) SetHeadersRead(value zend.ZendBool)   { this.headers_read = value }
func (this SapiRequestInfo) GetPostEntry() *SapiPostEntry          { return this.post_entry }
func (this *SapiRequestInfo) SetPostEntry(value *SapiPostEntry)    { this.post_entry = value }
func (this SapiRequestInfo) GetContentTypeDup() *byte              { return this.content_type_dup }
func (this *SapiRequestInfo) SetContentTypeDup(value *byte)        { this.content_type_dup = value }
func (this SapiRequestInfo) GetAuthUser() *byte                    { return this.auth_user }
func (this *SapiRequestInfo) SetAuthUser(value *byte)              { this.auth_user = value }
func (this SapiRequestInfo) GetAuthPassword() *byte                { return this.auth_password }
func (this *SapiRequestInfo) SetAuthPassword(value *byte)          { this.auth_password = value }
func (this SapiRequestInfo) GetAuthDigest() *byte                  { return this.auth_digest }
func (this *SapiRequestInfo) SetAuthDigest(value *byte)            { this.auth_digest = value }
func (this SapiRequestInfo) GetArgv0() *byte                       { return this.argv0 }
func (this *SapiRequestInfo) SetArgv0(value *byte)                 { this.argv0 = value }
func (this SapiRequestInfo) GetCurrentUser() *byte                 { return this.current_user }
func (this *SapiRequestInfo) SetCurrentUser(value *byte)           { this.current_user = value }
func (this SapiRequestInfo) GetCurrentUserLength() int             { return this.current_user_length }
func (this *SapiRequestInfo) SetCurrentUserLength(value int)       { this.current_user_length = value }
func (this SapiRequestInfo) GetArgc() int                          { return this.argc }
func (this *SapiRequestInfo) SetArgc(value int)                    { this.argc = value }
func (this SapiRequestInfo) GetArgv() **byte                       { return this.argv }
func (this *SapiRequestInfo) SetArgv(value **byte)                 { this.argv = value }
func (this SapiRequestInfo) GetProtoNum() int                      { return this.proto_num }
func (this *SapiRequestInfo) SetProtoNum(value int)                { this.proto_num = value }

/**
 * sapi_globals_struct
 */
type sapi_globals_struct struct {
	server_context           any
	request_info             SapiRequestInfo
	sapi_headers             SapiHeaders
	read_post_bytes          int64
	post_read                uint8
	headers_sent             uint8
	global_stat              zend.ZendStatT
	default_mimetype         *byte
	default_charset          *byte
	rfc1867_uploaded_files   *zend.HashTable
	post_max_size            zend.ZendLong
	options                  int
	sapi_started             zend.ZendBool
	global_request_time      float64
	known_post_content_types zend.HashTable
	callback_func            zend.Zval
	fci_cache                zend.ZendFcallInfoCache
}

func (this sapi_globals_struct) GetServerContext() any                 { return this.server_context }
func (this *sapi_globals_struct) SetServerContext(value any)           { this.server_context = value }
func (this sapi_globals_struct) GetRequestInfo() SapiRequestInfo       { return this.request_info }
func (this *sapi_globals_struct) SetRequestInfo(value SapiRequestInfo) { this.request_info = value }
func (this sapi_globals_struct) GetSapiHeaders() SapiHeaders           { return this.sapi_headers }
func (this *sapi_globals_struct) SetSapiHeaders(value SapiHeaders)     { this.sapi_headers = value }
func (this sapi_globals_struct) GetReadPostBytes() int64               { return this.read_post_bytes }
func (this *sapi_globals_struct) SetReadPostBytes(value int64)         { this.read_post_bytes = value }
func (this sapi_globals_struct) GetPostRead() uint8                    { return this.post_read }
func (this *sapi_globals_struct) SetPostRead(value uint8)              { this.post_read = value }
func (this sapi_globals_struct) GetHeadersSent() uint8                 { return this.headers_sent }
func (this *sapi_globals_struct) SetHeadersSent(value uint8)           { this.headers_sent = value }
func (this sapi_globals_struct) GetGlobalStat() zend.ZendStatT         { return this.global_stat }
func (this *sapi_globals_struct) SetGlobalStat(value zend.ZendStatT)   { this.global_stat = value }
func (this sapi_globals_struct) GetDefaultMimetype() *byte             { return this.default_mimetype }
func (this *sapi_globals_struct) SetDefaultMimetype(value *byte)       { this.default_mimetype = value }
func (this sapi_globals_struct) GetDefaultCharset() *byte              { return this.default_charset }
func (this *sapi_globals_struct) SetDefaultCharset(value *byte)        { this.default_charset = value }
func (this sapi_globals_struct) GetRfc1867UploadedFiles() *zend.HashTable {
	return this.rfc1867_uploaded_files
}
func (this *sapi_globals_struct) SetRfc1867UploadedFiles(value *zend.HashTable) {
	this.rfc1867_uploaded_files = value
}
func (this sapi_globals_struct) GetPostMaxSize() zend.ZendLong       { return this.post_max_size }
func (this *sapi_globals_struct) SetPostMaxSize(value zend.ZendLong) { this.post_max_size = value }
func (this sapi_globals_struct) GetOptions() int                     { return this.options }
func (this *sapi_globals_struct) SetOptions(value int)               { this.options = value }
func (this sapi_globals_struct) GetSapiStarted() zend.ZendBool       { return this.sapi_started }
func (this *sapi_globals_struct) SetSapiStarted(value zend.ZendBool) { this.sapi_started = value }
func (this sapi_globals_struct) GetGlobalRequestTime() float64       { return this.global_request_time }
func (this *sapi_globals_struct) SetGlobalRequestTime(value float64) {
	this.global_request_time = value
}
func (this sapi_globals_struct) GetKnownPostContentTypes() zend.HashTable {
	return this.known_post_content_types
}
func (this *sapi_globals_struct) SetKnownPostContentTypes(value zend.HashTable) {
	this.known_post_content_types = value
}
func (this sapi_globals_struct) GetCallbackFunc() zend.Zval                 { return this.callback_func }
func (this *sapi_globals_struct) SetCallbackFunc(value zend.Zval)           { this.callback_func = value }
func (this sapi_globals_struct) GetFciCache() zend.ZendFcallInfoCache       { return this.fci_cache }
func (this *sapi_globals_struct) SetFciCache(value zend.ZendFcallInfoCache) { this.fci_cache = value }

/**
 * SapiHeaderLine
 */
type SapiHeaderLine struct {
	line          *byte
	line_len      int
	response_code zend.ZendLong
}

func (this SapiHeaderLine) GetLine() *byte                       { return this.line }
func (this *SapiHeaderLine) SetLine(value *byte)                 { this.line = value }
func (this SapiHeaderLine) GetLineLen() int                      { return this.line_len }
func (this *SapiHeaderLine) SetLineLen(value int)                { this.line_len = value }
func (this SapiHeaderLine) GetResponseCode() zend.ZendLong       { return this.response_code }
func (this *SapiHeaderLine) SetResponseCode(value zend.ZendLong) { this.response_code = value }

/**
 * _sapiModule
 */
type _sapiModule struct {
	name                      *byte
	pretty_name               *byte
	startup                   func(sapi_module *_sapiModule) int
	shutdown                  func(sapi_module *_sapiModule) int
	activate                  func() int
	deactivate                func() int
	ub_write                  func(str *byte, str_length int) int
	flush                     func(server_context any)
	get_stat                  func() *zend.ZendStatT
	getenv                    func(name *byte, name_len int) *byte
	sapi_error                func(type_ int, error_msg *byte, _ ...any)
	header_handler            func(sapi_header *SapiHeader, op SapiHeaderOpEnum, sapi_headers *SapiHeaders) int
	send_headers              func(sapi_headers *SapiHeaders) int
	send_header               func(sapi_header *SapiHeader, server_context any)
	read_post                 func(buffer *byte, count_bytes int) int
	read_cookies              func() *byte
	register_server_variables func(track_vars_array *zend.Zval)
	log_message               func(message *byte, syslog_type_int int)
	get_request_time          func() float64
	terminate_process         func()
	php_ini_path_override     *byte
	default_post_reader       func()
	treat_data                func(arg int, str *byte, destArray *zend.Zval)
	executable_location       *byte
	php_ini_ignore            int
	php_ini_ignore_cwd        int
	get_fd                    func(fd *int) int
	force_http_10             func() int
	get_target_uid            func(*uid_t) int
	get_target_gid            func(*gid_t) int
	input_filter              func(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint
	ini_defaults              func(configuration_hash *zend.HashTable)
	phpinfo_as_text           int
	ini_entries               *byte
	additional_functions      *zend.ZendFunctionEntry
	input_filter_init         func() uint
}

func (this _sapiModule) GetName() *byte                                        { return this.name }
func (this *_sapiModule) SetName(value *byte)                                  { this.name = value }
func (this _sapiModule) GetPrettyName() *byte                                  { return this.pretty_name }
func (this *_sapiModule) SetPrettyName(value *byte)                            { this.pretty_name = value }
func (this _sapiModule) GetStartup() func(sapi_module *_sapiModule) int        { return this.startup }
func (this *_sapiModule) SetStartup(value func(sapi_module *_sapiModule) int)  { this.startup = value }
func (this _sapiModule) GetShutdown() func(sapi_module *_sapiModule) int       { return this.shutdown }
func (this *_sapiModule) SetShutdown(value func(sapi_module *_sapiModule) int) { this.shutdown = value }
func (this _sapiModule) GetActivate() func() int                               { return this.activate }
func (this *_sapiModule) SetActivate(value func() int)                         { this.activate = value }
func (this _sapiModule) GetDeactivate() func() int                             { return this.deactivate }
func (this *_sapiModule) SetDeactivate(value func() int)                       { this.deactivate = value }
func (this _sapiModule) GetUbWrite() func(str *byte, str_length int) int       { return this.ub_write }
func (this *_sapiModule) SetUbWrite(value func(str *byte, str_length int) int) { this.ub_write = value }
func (this _sapiModule) GetFlush() func(server_context any)                    { return this.flush }
func (this *_sapiModule) SetFlush(value func(server_context any))              { this.flush = value }
func (this _sapiModule) GetGetStat() func() *zend.ZendStatT                    { return this.get_stat }
func (this *_sapiModule) SetGetStat(value func() *zend.ZendStatT)              { this.get_stat = value }
func (this _sapiModule) GetGetenv() func(name *byte, name_len int) *byte       { return this.getenv }
func (this *_sapiModule) SetGetenv(value func(name *byte, name_len int) *byte) { this.getenv = value }
func (this _sapiModule) GetSapiError() func(type_ int, error_msg *byte, _ ...any) {
	return this.sapi_error
}
func (this *_sapiModule) SetSapiError(value func(type_ int, error_msg *byte, _ ...any)) {
	this.sapi_error = value
}
func (this _sapiModule) GetHeaderHandler() func(sapi_header *SapiHeader, op SapiHeaderOpEnum, sapi_headers *SapiHeaders) int {
	return this.header_handler
}
func (this *_sapiModule) SetHeaderHandler(value func(sapi_header *SapiHeader, op SapiHeaderOpEnum, sapi_headers *SapiHeaders) int) {
	this.header_handler = value
}
func (this _sapiModule) GetSendHeaders() func(sapi_headers *SapiHeaders) int {
	return this.send_headers
}
func (this *_sapiModule) SetSendHeaders(value func(sapi_headers *SapiHeaders) int) {
	this.send_headers = value
}
func (this _sapiModule) GetSendHeader() func(sapi_header *SapiHeader, server_context any) {
	return this.send_header
}
func (this *_sapiModule) SetSendHeader(value func(sapi_header *SapiHeader, server_context any)) {
	this.send_header = value
}
func (this _sapiModule) GetReadPost() func(buffer *byte, count_bytes int) int { return this.read_post }
func (this *_sapiModule) SetReadPost(value func(buffer *byte, count_bytes int) int) {
	this.read_post = value
}
func (this _sapiModule) GetReadCookies() func() *byte       { return this.read_cookies }
func (this *_sapiModule) SetReadCookies(value func() *byte) { this.read_cookies = value }
func (this _sapiModule) GetRegisterServerVariables() func(track_vars_array *zend.Zval) {
	return this.register_server_variables
}
func (this *_sapiModule) SetRegisterServerVariables(value func(track_vars_array *zend.Zval)) {
	this.register_server_variables = value
}
func (this _sapiModule) GetLogMessage() func(message *byte, syslog_type_int int) {
	return this.log_message
}
func (this *_sapiModule) SetLogMessage(value func(message *byte, syslog_type_int int)) {
	this.log_message = value
}
func (this _sapiModule) GetGetRequestTime() func() float64       { return this.get_request_time }
func (this *_sapiModule) SetGetRequestTime(value func() float64) { this.get_request_time = value }
func (this _sapiModule) GetTerminateProcess() func()             { return this.terminate_process }
func (this *_sapiModule) SetTerminateProcess(value func())       { this.terminate_process = value }
func (this _sapiModule) GetPhpIniPathOverride() *byte            { return this.php_ini_path_override }
func (this *_sapiModule) SetPhpIniPathOverride(value *byte)      { this.php_ini_path_override = value }
func (this _sapiModule) GetDefaultPostReader() func()            { return this.default_post_reader }
func (this *_sapiModule) SetDefaultPostReader(value func())      { this.default_post_reader = value }
func (this _sapiModule) GetTreatData() func(arg int, str *byte, destArray *zend.Zval) {
	return this.treat_data
}
func (this *_sapiModule) SetTreatData(value func(arg int, str *byte, destArray *zend.Zval)) {
	this.treat_data = value
}
func (this _sapiModule) GetExecutableLocation() *byte            { return this.executable_location }
func (this *_sapiModule) SetExecutableLocation(value *byte)      { this.executable_location = value }
func (this _sapiModule) GetPhpIniIgnore() int                    { return this.php_ini_ignore }
func (this *_sapiModule) SetPhpIniIgnore(value int)              { this.php_ini_ignore = value }
func (this _sapiModule) GetPhpIniIgnoreCwd() int                 { return this.php_ini_ignore_cwd }
func (this *_sapiModule) SetPhpIniIgnoreCwd(value int)           { this.php_ini_ignore_cwd = value }
func (this _sapiModule) GetGetFd() func(fd *int) int             { return this.get_fd }
func (this *_sapiModule) SetGetFd(value func(fd *int) int)       { this.get_fd = value }
func (this _sapiModule) GetForceHttp10() func() int              { return this.force_http_10 }
func (this *_sapiModule) SetForceHttp10(value func() int)        { this.force_http_10 = value }
func (this _sapiModule) GetGetTargetUid() func(*uid_t) int       { return this.get_target_uid }
func (this *_sapiModule) SetGetTargetUid(value func(*uid_t) int) { this.get_target_uid = value }
func (this _sapiModule) GetGetTargetGid() func(*gid_t) int       { return this.get_target_gid }
func (this *_sapiModule) SetGetTargetGid(value func(*gid_t) int) { this.get_target_gid = value }
func (this _sapiModule) GetInputFilter() func(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint {
	return this.input_filter
}
func (this *_sapiModule) SetInputFilter(value func(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint) {
	this.input_filter = value
}
func (this _sapiModule) GetIniDefaults() func(configuration_hash *zend.HashTable) {
	return this.ini_defaults
}
func (this *_sapiModule) SetIniDefaults(value func(configuration_hash *zend.HashTable)) {
	this.ini_defaults = value
}
func (this _sapiModule) GetPhpinfoAsText() int       { return this.phpinfo_as_text }
func (this *_sapiModule) SetPhpinfoAsText(value int) { this.phpinfo_as_text = value }
func (this _sapiModule) GetIniEntries() *byte        { return this.ini_entries }
func (this *_sapiModule) SetIniEntries(value *byte)  { this.ini_entries = value }
func (this _sapiModule) GetAdditionalFunctions() *zend.ZendFunctionEntry {
	return this.additional_functions
}
func (this *_sapiModule) SetAdditionalFunctions(value *zend.ZendFunctionEntry) {
	this.additional_functions = value
}
func (this _sapiModule) GetInputFilterInit() func() uint       { return this.input_filter_init }
func (this *_sapiModule) SetInputFilterInit(value func() uint) { this.input_filter_init = value }

/**
 * SapiPostEntry
 */
type SapiPostEntry struct {
	content_type     *byte
	content_type_len uint32
	post_reader      func()
	post_handler     func(content_type_dup *byte, arg any)
}

func (this SapiPostEntry) GetContentType() *byte           { return this.content_type }
func (this *SapiPostEntry) SetContentType(value *byte)     { this.content_type = value }
func (this SapiPostEntry) GetContentTypeLen() uint32       { return this.content_type_len }
func (this *SapiPostEntry) SetContentTypeLen(value uint32) { this.content_type_len = value }
func (this SapiPostEntry) GetPostReader() func()           { return this.post_reader }
func (this *SapiPostEntry) SetPostReader(value func())     { this.post_reader = value }
func (this SapiPostEntry) GetPostHandler() func(content_type_dup *byte, arg any) {
	return this.post_handler
}
func (this *SapiPostEntry) SetPostHandler(value func(content_type_dup *byte, arg any)) {
	this.post_handler = value
}
