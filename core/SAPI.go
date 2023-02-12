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
}

func MakeSapiModule(
	name string,
	pretty_name string,
	startup func(sapi_module *SapiModule) int,
	shutdown func(sapi_module *SapiModule) int,
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
