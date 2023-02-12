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

// func MakeSapiHeader(header *byte, header_len int) SapiHeader {
//     return SapiHeader{
//         header:header,
//         header_len:header_len,
//     }
// }
func (this *SapiHeader) GetHeader() *byte       { return this.header }
func (this *SapiHeader) SetHeader(value *byte)  { this.header = value }
func (this *SapiHeader) GetHeaderLen() int      { return this.header_len }
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

// func MakeSapiHeaders(headers zend.ZendLlist, http_response_code int, send_default_content_type uint8, mimetype *byte, http_status_line *byte) SapiHeaders {
//     return SapiHeaders{
//         headers:headers,
//         http_response_code:http_response_code,
//         send_default_content_type:send_default_content_type,
//         mimetype:mimetype,
//         http_status_line:http_status_line,
//     }
// }
func (this *SapiHeaders) GetHeaders() zend.ZendLlist { return this.headers }

// func (this *SapiHeaders) SetHeaders(value zend.ZendLlist) { this.headers = value }
// func (this *SapiHeaders)  GetHttpResponseCode() int      { return this.http_response_code }
// func (this *SapiHeaders) SetHttpResponseCode(value int) { this.http_response_code = value }
// func (this *SapiHeaders)  GetSendDefaultContentType() uint8      { return this.send_default_content_type }
// func (this *SapiHeaders) SetSendDefaultContentType(value uint8) { this.send_default_content_type = value }
// func (this *SapiHeaders)  GetMimetype() *byte      { return this.mimetype }
// func (this *SapiHeaders) SetMimetype(value *byte) { this.mimetype = value }
// func (this *SapiHeaders)  GetHttpStatusLine() *byte      { return this.http_status_line }
// func (this *SapiHeaders) SetHttpStatusLine(value *byte) { this.http_status_line = value }

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

//             func MakeSapiRequestInfo(
// request_method *byte,
// query_string *byte,
// cookie_data *byte,
// content_length zend.ZendLong,
// path_translated *byte,
// request_uri *byte,
// request_body *PhpStream,
// content_type *byte,
// headers_only zend.ZendBool,
// no_headers zend.ZendBool,
// headers_read zend.ZendBool,
// post_entry *SapiPostEntry,
// content_type_dup *byte,
// auth_user *byte,
// auth_password *byte,
// auth_digest *byte,
// argv0 *byte,
// current_user *byte,
// current_user_length int,
// argc int,
// argv **byte,
// proto_num int,
// ) SapiRequestInfo {
//                 return SapiRequestInfo{
//                     request_method:request_method,
//                     query_string:query_string,
//                     cookie_data:cookie_data,
//                     content_length:content_length,
//                     path_translated:path_translated,
//                     request_uri:request_uri,
//                     request_body:request_body,
//                     content_type:content_type,
//                     headers_only:headers_only,
//                     no_headers:no_headers,
//                     headers_read:headers_read,
//                     post_entry:post_entry,
//                     content_type_dup:content_type_dup,
//                     auth_user:auth_user,
//                     auth_password:auth_password,
//                     auth_digest:auth_digest,
//                     argv0:argv0,
//                     current_user:current_user,
//                     current_user_length:current_user_length,
//                     argc:argc,
//                     argv:argv,
//                     proto_num:proto_num,
//                 }
//             }
// func (this *SapiRequestInfo)  GetRequestMethod() *byte      { return this.request_method }
func (this *SapiRequestInfo) SetRequestMethod(value *byte) { this.request_method = value }

// func (this *SapiRequestInfo)  GetQueryString() *byte      { return this.query_string }
func (this *SapiRequestInfo) SetQueryString(value *byte) { this.query_string = value }

// func (this *SapiRequestInfo)  GetCookieData() *byte      { return this.cookie_data }
// func (this *SapiRequestInfo) SetCookieData(value *byte) { this.cookie_data = value }
// func (this *SapiRequestInfo)  GetContentLength() zend.ZendLong      { return this.content_length }
func (this *SapiRequestInfo) SetContentLength(value zend.ZendLong) { this.content_length = value }

// func (this *SapiRequestInfo)  GetPathTranslated() *byte      { return this.path_translated }
func (this *SapiRequestInfo) SetPathTranslated(value *byte) { this.path_translated = value }

// func (this *SapiRequestInfo)  GetRequestUri() *byte      { return this.request_uri }
func (this *SapiRequestInfo) SetRequestUri(value *byte) { this.request_uri = value }

// func (this *SapiRequestInfo)  GetRequestBody() *PhpStream      { return this.request_body }
// func (this *SapiRequestInfo) SetRequestBody(value *PhpStream) { this.request_body = value }
// func (this *SapiRequestInfo)  GetContentType() *byte      { return this.content_type }
func (this *SapiRequestInfo) SetContentType(value *byte) { this.content_type = value }

// func (this *SapiRequestInfo)  GetHeadersOnly() zend.ZendBool      { return this.headers_only }
// func (this *SapiRequestInfo) SetHeadersOnly(value zend.ZendBool) { this.headers_only = value }
// func (this *SapiRequestInfo)  GetNoHeaders() zend.ZendBool      { return this.no_headers }
// func (this *SapiRequestInfo) SetNoHeaders(value zend.ZendBool) { this.no_headers = value }
// func (this *SapiRequestInfo)  GetHeadersRead() zend.ZendBool      { return this.headers_read }
// func (this *SapiRequestInfo) SetHeadersRead(value zend.ZendBool) { this.headers_read = value }
// func (this *SapiRequestInfo)  GetPostEntry() *SapiPostEntry      { return this.post_entry }
// func (this *SapiRequestInfo) SetPostEntry(value *SapiPostEntry) { this.post_entry = value }
// func (this *SapiRequestInfo)  GetContentTypeDup() *byte      { return this.content_type_dup }
// func (this *SapiRequestInfo) SetContentTypeDup(value *byte) { this.content_type_dup = value }
// func (this *SapiRequestInfo)  GetAuthUser() *byte      { return this.auth_user }
func (this *SapiRequestInfo) SetAuthUser(value *byte)     { this.auth_user = value }
func (this *SapiRequestInfo) GetAuthPassword() *byte      { return this.auth_password }
func (this *SapiRequestInfo) SetAuthPassword(value *byte) { this.auth_password = value }
func (this *SapiRequestInfo) GetAuthDigest() *byte        { return this.auth_digest }
func (this *SapiRequestInfo) SetAuthDigest(value *byte)   { this.auth_digest = value }

// func (this *SapiRequestInfo)  GetArgv0() *byte      { return this.argv0 }
// func (this *SapiRequestInfo) SetArgv0(value *byte) { this.argv0 = value }
// func (this *SapiRequestInfo)  GetCurrentUser() *byte      { return this.current_user }
// func (this *SapiRequestInfo) SetCurrentUser(value *byte) { this.current_user = value }
// func (this *SapiRequestInfo)  GetCurrentUserLength() int      { return this.current_user_length }
// func (this *SapiRequestInfo) SetCurrentUserLength(value int) { this.current_user_length = value }
// func (this *SapiRequestInfo)  GetArgc() int      { return this.argc }
// func (this *SapiRequestInfo) SetArgc(value int) { this.argc = value }
// func (this *SapiRequestInfo)  GetArgv() **byte      { return this.argv }
// func (this *SapiRequestInfo) SetArgv(value **byte) { this.argv = value }
// func (this *SapiRequestInfo)  GetProtoNum() int      { return this.proto_num }
func (this *SapiRequestInfo) SetProtoNum(value int) { this.proto_num = value }

/**
 * SapiGlobals
 */
type SapiGlobals struct {
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

func (this *SapiGlobals) Init() {
	this.known_post_content_types = *zend.NewZendArrayEx(8, _typeDtor, true)
	PhpSetupSapiContentTypes()
}

func (this *SapiGlobals) Destroy() {
	this.known_post_content_types.Destroy()
}

/**
 * generate
 */
func (this *SapiGlobals) GetDefaultMimetype() *byte     { return this.default_mimetype }
func (this *SapiGlobals) GetDefaultCharset() *byte      { return this.default_charset }
func (this *SapiGlobals) GetPostMaxSize() zend.ZendLong { return this.post_max_size }
func (this *SapiGlobals) GetKnownPostContentTypes() zend.HashTable {
	return this.known_post_content_types
}

/**
 * SapiHeaderLine
 */
type SapiHeaderLine struct {
	line          *byte
	line_len      int
	response_code zend.ZendLong
}

func MakeSapiHeaderLineEx(line string) SapiHeaderLine {
	return SapiHeaderLine{}
}

func MakeSapiHeaderLine(line *byte, line_len int, response_code zend.ZendLong) SapiHeaderLine {
	return SapiHeaderLine{
		line:          line,
		line_len:      line_len,
		response_code: response_code,
	}
}
func (this *SapiHeaderLine) GetLine() *byte                 { return this.line }
func (this *SapiHeaderLine) SetLine(value *byte)            { this.line = value }
func (this *SapiHeaderLine) GetLineLen() int                { return this.line_len }
func (this *SapiHeaderLine) SetLineLen(value int)           { this.line_len = value }
func (this *SapiHeaderLine) GetResponseCode() zend.ZendLong { return this.response_code }

/**
 * SapiModule
 */
type SapiModule struct {
	name                      *byte
	pretty_name               *byte
	startup                   func(sapi_module *SapiModule) int
	shutdown                  func(sapi_module *SapiModule) int
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

var _ ISapiModule = (*SapiModule)(nil)

/**
 * generate
 */
func (this *SapiModule) GetName() *byte                                       { return this.name }
func (this *SapiModule) GetPrettyName() *byte                                 { return this.pretty_name }
func (this *SapiModule) GetStartup() func(sapi_module *SapiModule) int        { return this.startup }
func (this *SapiModule) GetActivate() func() int                              { return this.activate }
func (this *SapiModule) GetDeactivate() func() int                            { return this.deactivate }
func (this *SapiModule) GetUbWrite() func(str *byte, str_length int) int      { return this.ub_write }
func (this *SapiModule) SetUbWrite(value func(str *byte, str_length int) int) { this.ub_write = value }
func (this *SapiModule) GetFlush() func(server_context any)                   { return this.flush }
func (this *SapiModule) SetFlush(value func(server_context any))              { this.flush = value }
func (this *SapiModule) GetGetStat() func() *zend.ZendStatT                   { return this.get_stat }
func (this *SapiModule) GetGetenv() func(name *byte, name_len int) *byte      { return this.getenv }
func (this *SapiModule) SetGetenv(value func(name *byte, name_len int) *byte) { this.getenv = value }
func (this *SapiModule) GetSapiError() func(type_ int, error_msg *byte, _ ...any) {
	return this.sapi_error
}
func (this *SapiModule) GetHeaderHandler() func(sapi_header *SapiHeader, op SapiHeaderOpEnum, sapi_headers *SapiHeaders) int {
	return this.header_handler
}
func (this *SapiModule) GetSendHeaders() func(sapi_headers *SapiHeaders) int {
	return this.send_headers
}
func (this *SapiModule) SetSendHeaders(value func(sapi_headers *SapiHeaders) int) {
	this.send_headers = value
}
func (this *SapiModule) GetSendHeader() func(sapi_header *SapiHeader, server_context any) {
	return this.send_header
}
func (this *SapiModule) GetReadPost() func(buffer *byte, count_bytes int) int { return this.read_post }
func (this *SapiModule) SetReadPost(value func(buffer *byte, count_bytes int) int) {
	this.read_post = value
}
func (this *SapiModule) GetReadCookies() func() *byte      { return this.read_cookies }
func (this *SapiModule) SetReadCookies(value func() *byte) { this.read_cookies = value }
func (this *SapiModule) GetRegisterServerVariables() func(track_vars_array *zend.Zval) {
	return this.register_server_variables
}
func (this *SapiModule) GetLogMessage() func(message *byte, syslog_type_int int) {
	return this.log_message
}
func (this *SapiModule) GetGetRequestTime() func() float64 { return this.get_request_time }
func (this *SapiModule) GetTerminateProcess() func()       { return this.terminate_process }
func (this *SapiModule) GetPhpIniPathOverride() *byte      { return this.php_ini_path_override }
func (this *SapiModule) SetPhpIniPathOverride(value *byte) { this.php_ini_path_override = value }
func (this *SapiModule) GetDefaultPostReader() func()      { return this.default_post_reader }
func (this *SapiModule) SetDefaultPostReader(value func()) { this.default_post_reader = value }
func (this *SapiModule) GetTreatData() func(arg int, str *byte, destArray *zend.Zval) {
	return this.treat_data
}
func (this *SapiModule) SetTreatData(value func(arg int, str *byte, destArray *zend.Zval)) {
	this.treat_data = value
}
func (this *SapiModule) GetExecutableLocation() *byte      { return this.executable_location }
func (this *SapiModule) SetExecutableLocation(value *byte) { this.executable_location = value }
func (this *SapiModule) GetPhpIniIgnore() int              { return this.php_ini_ignore }
func (this *SapiModule) SetPhpIniIgnore(value int)         { this.php_ini_ignore = value }
func (this *SapiModule) GetPhpIniIgnoreCwd() int           { return this.php_ini_ignore_cwd }
func (this *SapiModule) SetPhpIniIgnoreCwd(value int)      { this.php_ini_ignore_cwd = value }
func (this *SapiModule) GetGetFd() func(fd *int) int       { return this.get_fd }
func (this *SapiModule) GetForceHttp10() func() int        { return this.force_http_10 }
func (this *SapiModule) GetGetTargetUid() func(*uid_t) int { return this.get_target_uid }
func (this *SapiModule) GetGetTargetGid() func(*gid_t) int { return this.get_target_gid }
func (this *SapiModule) GetInputFilter() func(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint {
	return this.input_filter
}
func (this *SapiModule) SetInputFilter(value func(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint) {
	this.input_filter = value
}
func (this *SapiModule) GetIniDefaults() func(configuration_hash *zend.HashTable) {
	return this.ini_defaults
}
func (this *SapiModule) SetIniDefaults(value func(configuration_hash *zend.HashTable)) {
	this.ini_defaults = value
}
func (this *SapiModule) GetPhpinfoAsText() int      { return this.phpinfo_as_text }
func (this *SapiModule) SetPhpinfoAsText(value int) { this.phpinfo_as_text = value }
func (this *SapiModule) GetIniEntries() *byte       { return this.ini_entries }
func (this *SapiModule) SetIniEntries(value *byte)  { this.ini_entries = value }
func (this *SapiModule) GetAdditionalFunctions() *zend.ZendFunctionEntry {
	return this.additional_functions
}
func (this *SapiModule) SetAdditionalFunctions(value *zend.ZendFunctionEntry) {
	this.additional_functions = value
}
func (this *SapiModule) GetInputFilterInit() func() uint      { return this.input_filter_init }
func (this *SapiModule) SetInputFilterInit(value func() uint) { this.input_filter_init = value }

/**
 * SapiPostEntry
 */
type SapiPostEntry struct {
	content_type     *byte
	content_type_len uint32
	post_reader      func()
	post_handler     func(content_type_dup *byte, arg any)
}

func MakeSapiPostEntry(content_type *byte, content_type_len uint32, post_reader func(), post_handler func(content_type_dup *byte, arg any)) SapiPostEntry {
	return SapiPostEntry{
		content_type:     content_type,
		content_type_len: content_type_len,
		post_reader:      post_reader,
		post_handler:     post_handler,
	}
}
func (this *SapiPostEntry) GetContentType() *byte { return this.content_type }

// func (this *SapiPostEntry) SetContentType(value *byte) { this.content_type = value }
func (this *SapiPostEntry) GetContentTypeLen() uint32 { return this.content_type_len }

// func (this *SapiPostEntry) SetContentTypeLen(value uint32) { this.content_type_len = value }
func (this *SapiPostEntry) GetPostReader() func() { return this.post_reader }

// func (this *SapiPostEntry) SetPostReader(value func()) { this.post_reader = value }
// func (this *SapiPostEntry)  GetPostHandler() func(content_type_dup *byte, arg any)      { return this.post_handler }
// func (this *SapiPostEntry) SetPostHandler(value func(content_type_dup *byte, arg any)) { this.post_handler = value }
