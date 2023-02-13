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

var PhpCliServerMaster pid_t
var PhpCliServerWorkers *pid_t
var PhpCliServerWorkersMax zend.ZendLong

var TemplateMap = map[int]string{
	400: "<h1>%s</h1><p>Your browser sent a request that this server could not understand.</p>",
	404: "<h1>%s</h1><p>The requested resource <code class=\"url\">%s</code> was not found on this server.</p>",
	500: "<h1>%s</h1><p>The server is temporarily unavailable.</p>",
	501: "<h1>%s</h1><p>Request method not supported.</p>",
}

const PHP_CLI_SERVER_LOG_PROCESS = 1
const PHP_CLI_SERVER_LOG_ERROR = 2
const PHP_CLI_SERVER_LOG_MESSAGE = 3

var PhpCliServerLogLevel int = 3
var PhpCliServerRequestErrorUnexpectedEof = "Unexpected EOF"

var CliServerModuleEntry = zend.MakeZendModuleEntry(
	b.SizeOf("zend_module_entry"),
	zend.ZEND_MODULE_API_NO,
	0, zend.USING_ZTS,
	nil,
	nil,
	"cli_server",
	nil,
	ZmStartupCliServer,
	ZmShutdownCliServer,
	nil,
	nil,
	ZmInfoCliServer,
	core.PHP_VERSION,
	0,
	nil,
	nil,
	nil,
	nil,
	0,
	0,
	nil,
	0,
	"API"+"ZEND_MODULE_API_NO"+zend.ZEND_BUILD_TS,
)

var ServerAdditionalFunctions = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntryEx("apache_request_headers", 0, ZifApacheRequestHeaders, nil),
	zend.MakeZendFunctionEntryEx("apache_response_headers", 0, ZifApacheResponseHeaders, nil),
	zend.MakeZendFunctionEntryEx("getallheaders", 0, ZifApacheRequestHeaders, nil),
}

var Server PhpCliServer
