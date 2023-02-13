package core

import (
	"sik/zend"
)

type ISapiModule interface {
	Name() string
	PrettyName() string
	Startup() bool
	Shutdown() bool
	Activate()
	Deactivate()
	UbWrite(str string) int

	// getter/setter
	SetUbWrite(value func(str *byte, str_length int) int)
	SetFlush(value func(server_context any))
	SetGetenv(value func(name *byte, name_len int) *byte)
	GetSendHeaders() func(sapi_headers *SapiHeaders) int
	SetSendHeaders(value func(sapi_headers *SapiHeaders) int)
	GetSendHeader() func(sapi_header *SapiHeader, server_context any)
	GetReadPost() func(buffer *byte, count_bytes int) int
	SetReadPost(value func(buffer *byte, count_bytes int) int)
	GetReadCookies() func() *byte
	SetReadCookies(value func() *byte)
	GetRegisterServerVariables() func(track_vars_array *zend.Zval)
	GetLogMessage() func(message *byte, syslog_type_int int)
	GetPhpIniPathOverride() *byte
	SetPhpIniPathOverride(value *byte)
	GetDefaultPostReader() func()
	SetDefaultPostReader(value func())
	GetTreatData() func(arg int, str *byte, destArray *zend.Zval)
	SetTreatData(value func(arg int, str *byte, destArray *zend.Zval))
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
	GetIniDefaults() func(configuration_hash *zend.HashTable)
	SetIniDefaults(value func(configuration_hash *zend.HashTable))
	GetPhpinfoAsText() int
	SetPhpinfoAsText(value int)
	GetIniEntries() *byte
	SetIniEntries(value *byte)
	GetAdditionalFunctions() *zend.ZendFunctionEntry
	SetAdditionalFunctions(value *zend.ZendFunctionEntry)
	GetInputFilterInit() func() uint
	SetInputFilterInit(value func() uint)
}

func MakeSapiModule(
	name string,
	pretty_name string,
	startup func(sapi_module ISapiModule) int,
	shutdown func(sapi_module ISapiModule) int,
	activate func() int,
	deactivate func() int,
	ub_write func(str *byte, str_length int) int,
	flush func(server_context any),
	getenv func(name *byte, name_len int) *byte,
	header_handler func(sapi_header *SapiHeader, op SapiHeaderOpEnum, sapi_headers *SapiHeaders) int,
	send_headers func(sapi_headers *SapiHeaders) int,
	send_header func(sapi_header *SapiHeader, server_context any),
	read_post func(buffer *byte, count_bytes int) int,
	read_cookies func() *byte,
	register_server_variables func(track_vars_array *zend.Zval),
	log_message func(message *byte, syslog_type_int int),
) SapiModule {
	return SapiModule{
		name:                      name,
		pretty_name:               pretty_name,
		startup:                   startup,
		shutdown:                  shutdown,
		activate:                  activate,
		deactivate:                deactivate,
		ub_write:                  ub_write,
		flush:                     flush,
		getenv:                    getenv,
		header_handler:            header_handler,
		send_headers:              send_headers,
		send_header:               send_header,
		read_post:                 read_post,
		read_cookies:              read_cookies,
		register_server_variables: register_server_variables,
		log_message:               log_message,
	}
}
