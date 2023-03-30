package core

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
)

type ISapiModule interface {
	Name() string
	PrettyName() string
	Startup() bool
	Shutdown() bool
	Activate()
	Deactivate()
	UbWrite(str string) (int, error)
	Flush(serverContext any)
	GetStat() bool
	GetEnv(name string) (string, bool)
	HeaderHandler(header *SapiHeader, op SapiHeaderOpEnum, headers *SapiHeaders) int
	SendHeaders(headers *SapiHeaders) int
	SendHeader(header *SapiHeader, serverContext any)
	ReadPost(buffer *byte, count_bytes int) int
	ReadCookies() (string, bool)
	RegisterServerVariables(trackVarsArray []types.Zval)
	LogMessage(message string, syslogType int)
	InputFilter(arg int, name string, value string) string

	// getter/setter
	GetSendHeaders() func(sapi_headers *SapiHeaders) int
	GetSendHeader() func(sapi_header *SapiHeader, server_context any)
	GetReadPost() func(buffer *byte, count_bytes int) int
	GetReadCookies() func() *byte
	GetRegisterServerVariables() func(track_vars_array *types.Zval)
	GetLogMessage() func(message *byte, syslog_type_int int)
	GetPhpIniPathOverride() *byte
	SetPhpIniPathOverride(value *byte)
	GetDefaultPostReader() func()
	SetDefaultPostReader(value func())
	GetTreatData() func(arg int, str *byte, destArray *types.Zval)
	SetTreatData(value func(arg int, str *byte, destArray *types.Zval))
	GetExecutableLocation() *byte
	SetExecutableLocation(value *byte)
	GetPhpIniIgnore() int
	SetPhpIniIgnore(value int)
	GetPhpIniIgnoreCwd() int
	SetPhpIniIgnoreCwd(value int)
	GetGetFd() func(fd *int) int
	GetForceHttp10() func() int
	GetInputFilter() func(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint
	SetInputFilter(value func(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint)
	GetIniDefaults() func(configuration_hash *types.Array)
	SetIniDefaults(value func(configuration_hash *types.Array))
	GetPhpinfoAsText() int
	SetPhpinfoAsText(value int)
	GetIniEntries() *byte
	SetIniEntries(value *byte)
	GetAdditionalFunctions() *types.ZendFunctionEntry
	SetAdditionalFunctions(value *types.ZendFunctionEntry)
	GetInputFilterInit() func() uint
	SetInputFilterInit(value func() uint)
}

/**
 * BaseSapiModule
 */
type BaseSapiModule struct {
	ub_write                  func(str *byte, str_length int) int
	flush                     func(server_context any)
	get_stat                  func() *zend.ZendStatT
	getenv                    func(name *byte, name_len int) *byte
	header_handler            func(sapi_header *SapiHeader, op SapiHeaderOpEnum, sapi_headers *SapiHeaders) int
	send_headers              func(sapi_headers *SapiHeaders) int
	send_header               func(sapi_header *SapiHeader, server_context any)
	read_post                 func(buffer *byte, count_bytes int) int
	read_cookies              func() *byte
	register_server_variables func(track_vars_array *types.Zval)
	log_message               func(message *byte, syslog_type_int int)

	php_ini_path_override *byte
	default_post_reader   func()
	treat_data            func(arg int, str *byte, destArray *types.Zval)
	executable_location   *byte
	php_ini_ignore        int
	php_ini_ignore_cwd    int
	get_fd                func(fd *int) int
	force_http_10         func() int
	input_filter          func(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint
	ini_defaults          func(configuration_hash *types.Array)
	phpinfo_as_text       int
	ini_entries           *byte
	additional_functions  *types.ZendFunctionEntry
	input_filter_init     func() uint
}

var _ ISapiModule = (*BaseSapiModule)(nil)

func (this *BaseSapiModule) Name() string                      { panic("implement me") }
func (this *BaseSapiModule) PrettyName() string                { panic("implement me") }
func (this *BaseSapiModule) Startup() bool                     { panic("implement me") }
func (this *BaseSapiModule) Shutdown() bool                    { panic("implement me") }
func (this *BaseSapiModule) Activate()                         { panic("implement me") }
func (this *BaseSapiModule) Deactivate()                       { panic("implement me") }
func (this *BaseSapiModule) UbWrite(str string) (int, error)   { panic("implement me") }
func (this *BaseSapiModule) Flush(serverContext any)           {}
func (this *BaseSapiModule) GetStat() bool                     { return false }
func (this *BaseSapiModule) GetEnv(name string) (string, bool) { return "", false }
func (this *BaseSapiModule) HeaderHandler(header *SapiHeader, op SapiHeaderOpEnum, headers *SapiHeaders) int {
	return 0
}
func (this *BaseSapiModule) SendHeaders(headers *SapiHeaders) int {
	panic("implement me")
}

func (this *BaseSapiModule) SendHeader(header *SapiHeader, serverContext any) {}

func (this *BaseSapiModule) ReadPost(buffer *byte, count_bytes int) int {
	panic("implement me")
}

func (this *BaseSapiModule) ReadCookies() (string, bool) {
	panic("implement me")
}

func (this *BaseSapiModule) RegisterServerVariables(trackVarsArray []types.Zval) {
	panic("implement me")
}

func (this *BaseSapiModule) LogMessage(message string, syslogType int) {
	panic("implement me")
}

func (this *BaseSapiModule) InputFilter(arg int, name string, value string) string {
	if this.input_filter != nil {
		this.input_filter(arg, name, &value, len(value), nil)
		return value
	}
	return value
}

func (this *BaseSapiModule) SapiError(type_ int, error_msg string, args ...any) {
	PhpError(type_, error_msg, args...)
}

func (this *BaseSapiModule) InputFilterInit() {
	if this.input_filter_init != nil {
		this.input_filter_init()
	}
}

/**
 * generate
 */

func (this *BaseSapiModule) GetSendHeaders() func(sapi_headers *SapiHeaders) int {
	return this.send_headers
}
func (this *BaseSapiModule) GetSendHeader() func(sapi_header *SapiHeader, server_context any) {
	return this.send_header
}
func (this *BaseSapiModule) GetReadPost() func(buffer *byte, count_bytes int) int {
	return this.read_post
}
func (this *BaseSapiModule) GetReadCookies() func() *byte      { return this.read_cookies }
func (this *BaseSapiModule) SetReadCookies(value func() *byte) { this.read_cookies = value }
func (this *BaseSapiModule) GetRegisterServerVariables() func(track_vars_array *types.Zval) {
	return this.register_server_variables
}
func (this *BaseSapiModule) GetLogMessage() func(message *byte, syslog_type_int int) {
	return this.log_message
}
func (this *BaseSapiModule) GetPhpIniPathOverride() *byte      { return this.php_ini_path_override }
func (this *BaseSapiModule) SetPhpIniPathOverride(value *byte) { this.php_ini_path_override = value }
func (this *BaseSapiModule) GetDefaultPostReader() func()      { return this.default_post_reader }
func (this *BaseSapiModule) SetDefaultPostReader(value func()) { this.default_post_reader = value }
func (this *BaseSapiModule) GetTreatData() func(arg int, str *byte, destArray *types.Zval) {
	return this.treat_data
}
func (this *BaseSapiModule) SetTreatData(value func(arg int, str *byte, destArray *types.Zval)) {
	this.treat_data = value
}
func (this *BaseSapiModule) GetExecutableLocation() *byte      { return this.executable_location }
func (this *BaseSapiModule) SetExecutableLocation(value *byte) { this.executable_location = value }
func (this *BaseSapiModule) GetPhpIniIgnore() int              { return this.php_ini_ignore }
func (this *BaseSapiModule) SetPhpIniIgnore(value int)         { this.php_ini_ignore = value }
func (this *BaseSapiModule) GetPhpIniIgnoreCwd() int           { return this.php_ini_ignore_cwd }
func (this *BaseSapiModule) SetPhpIniIgnoreCwd(value int)      { this.php_ini_ignore_cwd = value }
func (this *BaseSapiModule) GetGetFd() func(fd *int) int       { return this.get_fd }
func (this *BaseSapiModule) GetForceHttp10() func() int        { return this.force_http_10 }
func (this *BaseSapiModule) GetInputFilter() func(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint {
	return this.input_filter
}
func (this *BaseSapiModule) SetInputFilter(value func(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint) {
	this.input_filter = value
}
func (this *BaseSapiModule) GetIniDefaults() func(configuration_hash *types.Array) {
	return this.ini_defaults
}
func (this *BaseSapiModule) SetIniDefaults(value func(configuration_hash *types.Array)) {
	this.ini_defaults = value
}
func (this *BaseSapiModule) GetPhpinfoAsText() int      { return this.phpinfo_as_text }
func (this *BaseSapiModule) SetPhpinfoAsText(value int) { this.phpinfo_as_text = value }
func (this *BaseSapiModule) GetIniEntries() *byte       { return this.ini_entries }
func (this *BaseSapiModule) SetIniEntries(value *byte)  { this.ini_entries = value }
func (this *BaseSapiModule) GetAdditionalFunctions() *types.ZendFunctionEntry {
	return this.additional_functions
}
func (this *BaseSapiModule) SetAdditionalFunctions(value *types.ZendFunctionEntry) {
	this.additional_functions = value
}
func (this *BaseSapiModule) GetInputFilterInit() func() uint      { return this.input_filter_init }
func (this *BaseSapiModule) SetInputFilterInit(value func() uint) { this.input_filter_init = value }
