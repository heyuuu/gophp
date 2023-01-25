// <<generate>>

package cli

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

// Source: <sapi/cli/php_cli_server.h>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Author: Moriyoshi Koizumi <moriyoshi@php.net>                        |
   +----------------------------------------------------------------------+
*/

// Source: <sapi/cli/php_cli_server.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Author: Moriyoshi Koizumi <moriyoshi@php.net>                        |
   |         Xinchen Hui       <laruence@php.net>                         |
   +----------------------------------------------------------------------+
*/

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

var TemplateMap []PhpCliServerHttpResponseStatusCodePair = []PhpCliServerHttpResponseStatusCodePair{
	{
		400,
		"<h1>%s</h1><p>Your browser sent a request that this server could not understand.</p>",
	},
	{
		404,
		"<h1>%s</h1><p>The requested resource <code class=\"url\">%s</code> was not found on this server.</p>",
	},
	{
		500,
		"<h1>%s</h1><p>The server is temporarily unavailable.</p>",
	},
	{
		501,
		"<h1>%s</h1><p>Request method not supported.</p>",
	},
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

var PhpCliServerCss []byte = "<style>\n" + "body { background-color: #fcfcfc; color: #333333; margin: 0; padding:0; }\n" + "h1 { font-size: 1.5em; font-weight: normal; background-color: #9999cc; min-height:2em; line-height:2em; border-bottom: 1px inset black; margin: 0; }\n" + "h1, p { padding-left: 10px; }\n" + "code.url { background-color: #eeeeee; font-family:monospace; padding:0 2px;}\n" + "</style>\n"
var IniEntries []zend.ZendIniEntryDef = []zend.ZendIniEntryDef{
	{
		"cli_server.color",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendCliServerGlobals)(nil).GetColor())) - (*byte)(nil))),
		any(&CliServerGlobals),
		nil,
		"0",
		zend.ZendIniBooleanDisplayerCb,
		b.SizeOf("\"0\"") - 1,
		b.SizeOf("\"cli_server.color\"") - 1,
		core.PHP_INI_ALL,
	},
	{nil, nil, nil, nil, nil, nil, nil, 0, 0, 0},
}
var CliServerModuleEntry zend.ZendModuleEntry = zend.ZendModuleEntry{
	b.SizeOf("zend_module_entry"),
	zend.ZEND_MODULE_API_NO,
	core.ZEND_DEBUG,
	zend.USING_ZTS,
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
	"API" + "ZEND_MODULE_API_NO" + zend.ZEND_BUILD_TS,
}
var ArginfoNoArgs []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var ServerAdditionalFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"cli_set_process_title",
		ZifCliSetProcessTitle,
		ArginfoCliSetProcessTitle,
		uint32(b.SizeOf("arginfo_cli_set_process_title")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"cli_get_process_title",
		ZifCliGetProcessTitle,
		ArginfoCliGetProcessTitle,
		uint32(b.SizeOf("arginfo_cli_get_process_title")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"apache_request_headers",
		ZifApacheRequestHeaders,
		ArginfoNoArgs,
		uint32(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"apache_response_headers",
		ZifApacheResponseHeaders,
		ArginfoNoArgs,
		uint32(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getallheaders",
		ZifApacheRequestHeaders,
		ArginfoNoArgs,
		uint32(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

/* {{{ sapi_module_struct cli_server_sapi_module
 */

var CliServerSapiModule core.sapi_module_struct = core.sapi_module_struct{"cli-server", "Built-in HTTP server", SapiCliServerStartup, core.PhpModuleShutdownWrapper, nil, nil, SapiCliServerUbWrite, SapiCliServerFlush, nil, nil, core.PhpError, nil, SapiCliServerSendHeaders, nil, SapiCliServerReadPost, SapiCliServerReadCookies, SapiCliServerRegisterVariables, SapiCliServerLogMessage, nil, nil, nil, nil, nil, nil, 0, 0, nil, nil, nil, nil, nil, nil, 0, nil, nil, nil}

/* {{{ php_cli_server_client_read_request */

var Server PhpCliServer
