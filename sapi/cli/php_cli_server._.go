// <<generate>>

package cli

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

const SOCK_EINVAL = EINVAL
const SOCK_EAGAIN = EAGAIN
const SOCK_EINTR = EINTR
const SOCK_EADDRINUSE = EADDRINUSE

// failed # include "ext/date/php_date.h"

const OUTPUT_NOT_CHECKED = -1
const OUTPUT_IS_TTY = 1
const OUTPUT_NOT_TTY = 0

var PhpCliServerMaster pid_t
var PhpCliServerWorkers *pid_t
var PhpCliServerWorkersMax zend.ZendLong

var TemplateMap = []PhpCliServerHttpResponseStatusCodePair{
	MakePhpCliServerHttpResponseStatusCodePair(400, "<h1>%s</h1><p>Your browser sent a request that this server could not understand.</p>"),
	MakePhpCliServerHttpResponseStatusCodePair(404, "<h1>%s</h1><p>The requested resource <code class=\"url\">%s</code> was not found on this server.</p>"),
	MakePhpCliServerHttpResponseStatusCodePair(500, "<h1>%s</h1><p>The server is temporarily unavailable.</p>"),
	MakePhpCliServerHttpResponseStatusCodePair(501, "<h1>%s</h1><p>Request method not supported.</p>"),
}

const PHP_CLI_SERVER_LOG_PROCESS = 1
const PHP_CLI_SERVER_LOG_ERROR = 2
const PHP_CLI_SERVER_LOG_MESSAGE = 3

var PhpCliServerLogLevel int = 3
var PhpCliOutputIsTty int = OUTPUT_NOT_CHECKED
var PhpCliServerRequestErrorUnexpectedEof []byte = "Unexpected EOF"
var CliServerGlobals ZendCliServerGlobals

/* {{{ static char php_cli_server_css[]
 * copied from ext/standard/info.c
 */

var CliServerModuleEntry = zend.MakeZendModuleEntry(b.SizeOf("zend_module_entry"), zend.ZEND_MODULE_API_NO, 0, zend.USING_ZTS, nil, nil, "cli_server", nil, ZmStartupCliServer, ZmShutdownCliServer, nil, nil, ZmInfoCliServer, core.PHP_VERSION, 0, nil, nil, nil, nil, 0, 0, nil, 0, "API"+"ZEND_MODULE_API_NO"+zend.ZEND_BUILD_TS)
var ArginfoNoArgs = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var ServerAdditionalFunctions = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("cli_set_process_title", ZifCliSetProcessTitle, ArginfoCliSetProcessTitle, uint32(b.SizeOf("arginfo_cli_set_process_title")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("cli_get_process_title", ZifCliGetProcessTitle, ArginfoCliGetProcessTitle, uint32(b.SizeOf("arginfo_cli_get_process_title")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("apache_request_headers", ZifApacheRequestHeaders, ArginfoNoArgs, uint32(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("apache_response_headers", ZifApacheResponseHeaders, ArginfoNoArgs, uint32(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getallheaders", ZifApacheRequestHeaders, ArginfoNoArgs, uint32(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}

var CliServerSapiModule = core.MakeSapiModule("cli-server", "Built-in HTTP server", SapiCliServerStartup, core.PhpModuleShutdownWrapper, nil, nil, SapiCliServerUbWrite, SapiCliServerFlush, nil, nil, SapiCliServerSendHeaders, nil, SapiCliServerReadPost, SapiCliServerReadCookies, SapiCliServerRegisterVariables, SapiCliServerLogMessage)

var Server PhpCliServer
