// <<generate>>

package cli

import (
	"sik/core"
	"sik/ext/standard"
	g "sik/runtime/grammar"
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

// #define PHP_CLI_SERVER_H

// # include "SAPI.h"

// @type ZendCliServerGlobals struct

// #define CLI_SERVER_G(v) ( cli_server_globals . v )

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

// # include < stdio . h >

// # include < stdlib . h >

// # include < fcntl . h >

// # include < assert . h >

// # include "php_config.h"

// # include < sys / time . h >

// # include < unistd . h >

// # include < signal . h >

// # include < locale . h >

// # include < dlfcn . h >

// # include "SAPI.h"

// # include "php.h"

// # include "php_ini.h"

// # include "php_main.h"

// # include "php_globals.h"

// # include "php_variables.h"

// # include "zend_hash.h"

// # include "zend_modules.h"

// # include "fopen_wrappers.h"

// # include "http_status_codes.h"

// # include "zend_compile.h"

// # include "zend_execute.h"

// # include "zend_highlight.h"

// # include "zend_exceptions.h"

// # include "php_getopt.h"

// #define php_select(m,r,w,e,t) select ( m , r , w , e , t )

// #define SOCK_EINVAL       EINVAL

// #define SOCK_EAGAIN       EAGAIN

// #define SOCK_EINTR       EINTR

// #define SOCK_EADDRINUSE       EADDRINUSE

// # include "ext/standard/file.h"

// # include "zend_smart_str.h"

// # include "ext/standard/html.h"

// # include "ext/standard/url.h"

// # include "ext/standard/php_string.h"

// failed # include "ext/date/php_date.h"

// # include "php_network.h"

// # include "php_http_parser.h"

// # include "php_cli_server.h"

// # include "mime_type_map.h"

// # include "php_cli_process_title.h"

// #define OUTPUT_NOT_CHECKED       - 1

// #define OUTPUT_IS_TTY       1

// #define OUTPUT_NOT_TTY       0

// # include < sys / wait . h >

var PhpCliServerMaster pid_t
var PhpCliServerWorkers *pid_t
var PhpCliServerWorkersMax zend.ZendLong

// @type PhpCliServerPoller struct

// @type PhpCliServerRequest struct

// @type PhpCliServerChunk struct

// @type PhpCliServerBuffer struct

// @type PhpCliServerContentSender struct

// @type PhpCliServerClient struct

// @type PhpCliServer struct

// @type PhpCliServerHttpResponseStatusCodePair struct

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

// #define PHP_CLI_SERVER_LOG_PROCESS       1

// #define PHP_CLI_SERVER_LOG_ERROR       2

// #define PHP_CLI_SERVER_LOG_MESSAGE       3

var PhpCliServerLogLevel int = 3
var PhpCliOutputIsTty int = -1
var PhpCliServerRequestErrorUnexpectedEof []byte = "Unexpected EOF"
var CliServerGlobals ZendCliServerGlobals

/* {{{ static char php_cli_server_css[]
 * copied from ext/standard/info.c
 */

var PhpCliServerCss []byte = "<style>\n" + "body { background-color: #fcfcfc; color: #333333; margin: 0; padding:0; }\n" + "h1 { font-size: 1.5em; font-weight: normal; background-color: #9999cc; min-height:2em; line-height:2em; border-bottom: 1px inset black; margin: 0; }\n" + "h1, p { padding-left: 10px; }\n" + "code.url { background-color: #eeeeee; font-family:monospace; padding:0 2px;}\n" + "</style>\n"

/* }}} */

func PhpCliServerGetSystemTime(buf *byte) int {
	var tv __struct__timeval
	var tm __struct__tm
	gettimeofday(&tv, nil)

	/* TODO: should be checked for NULL tm/return vaue */

	localtime_r(&tv.tv_sec, &tm)
	asctime_r(&tm, buf)
	return 0
}
func CharPtrDtorP(zv *zend.Zval) {
	g.CondF(true, func() { return zend.Free(zv.value.ptr) }, func() { return zend._efree(zv.value.ptr) })
}
func GetLastError() *byte { return strdup(strerror(errno)) }
func StatusComp(a any, b any) int {
	var pa *core.HttpResponseStatusCodePair = (*core.HttpResponseStatusCodePair)(a)
	var pb *core.HttpResponseStatusCodePair = (*core.HttpResponseStatusCodePair)(b)
	if pa.code < pb.code {
		return -1
	} else if pa.code > pb.code {
		return 1
	}
	return 0
}
func GetStatusString(code int) *byte {
	var needle core.HttpResponseStatusCodePair = core.HttpResponseStatusCodePair{code, nil}
	var result *core.HttpResponseStatusCodePair = nil
	result = bsearch(&needle, core.HttpStatusMap, core.HttpStatusMapLen, g.SizeOf("needle"), StatusComp)
	if result != nil {
		return result.str
	}

	/* Returning NULL would require complicating append_http_status_line() to
	 * not segfault in that case, so let's just return a placeholder, since RFC
	 * 2616 requires a reason phrase. This is basically what a lot of other Web
	 * servers do in this case anyway. */

	return "Unknown Status Code"

	/* Returning NULL would require complicating append_http_status_line() to
	 * not segfault in that case, so let's just return a placeholder, since RFC
	 * 2616 requires a reason phrase. This is basically what a lot of other Web
	 * servers do in this case anyway. */
}
func GetTemplateString(code int) *byte {
	var e int = g.SizeOf("template_map") / g.SizeOf("php_cli_server_http_response_status_code_pair")
	var s int = 0
	for e != s {
		var c int = g.Cond((e+s+1)/2 < e-1, (e+s+1)/2, e-1)
		var d int = TemplateMap[c].GetCode()
		if d > code {
			e = c
		} else if d < code {
			s = c
		} else {
			return TemplateMap[c].GetStr()
		}
	}
	return nil
}
func AppendHttpStatusLine(buffer *zend.SmartStr, protocol_version int, response_code int, persistent int) {
	if response_code == 0 {
		response_code = 200
	}
	zend.SmartStrAppendlEx(buffer, "HTTP", 4, persistent)
	zend.SmartStrAppendcEx(buffer, '/', persistent)
	zend.SmartStrAppendLongEx(buffer, protocol_version/100, persistent)
	zend.SmartStrAppendcEx(buffer, '.', persistent)
	zend.SmartStrAppendLongEx(buffer, protocol_version%100, persistent)
	zend.SmartStrAppendcEx(buffer, ' ', persistent)
	zend.SmartStrAppendLongEx(buffer, response_code, persistent)
	zend.SmartStrAppendcEx(buffer, ' ', persistent)
	zend.SmartStrAppendlEx(buffer, GetStatusString(response_code), strlen(GetStatusString(response_code)), persistent)
	zend.SmartStrAppendlEx(buffer, "\r\n", 2, persistent)
}
func AppendEssentialHeaders(buffer *zend.SmartStr, client *PhpCliServerClient, persistent int) {
	var val *byte
	var tv __struct__timeval = __struct__timeval{0}
	if nil != g.Assign(&val, zend.ZendHashStrFindPtr(&client.request.GetHeaders(), "host", g.SizeOf("\"host\"")-1)) {
		zend.SmartStrAppendlEx(buffer, "Host: ", strlen("Host: "), persistent)
		zend.SmartStrAppendlEx(buffer, val, strlen(val), persistent)
		zend.SmartStrAppendlEx(buffer, "\r\n", strlen("\r\n"), persistent)
	}
	if !(gettimeofday(&tv, nil)) {
		var dt *zend.ZendString = php_format_date("D, d M Y H:i:s", g.SizeOf("\"D, d M Y H:i:s\"")-1, tv.tv_sec, 0)
		zend.SmartStrAppendlEx(buffer, "Date: ", strlen("Date: "), persistent)
		zend.SmartStrAppendlEx(buffer, dt.val, strlen(dt.val), persistent)
		zend.SmartStrAppendlEx(buffer, " GMT\r\n", strlen(" GMT\r\n"), persistent)
		zend.ZendStringReleaseEx(dt, 0)
	}
	zend.SmartStrAppendlEx(buffer, "Connection: close\r\n", g.SizeOf("\"Connection: close\\r\\n\"")-1, persistent)
}
func GetMimeType(server *PhpCliServer, ext *byte, ext_len int) *byte {
	var ret *byte
	var ext_lower *byte = zend._emalloc(ext_len + 1)
	zend.ZendStrTolowerCopy(ext_lower, ext, ext_len)
	ret = zend.ZendHashStrFindPtr(&server.extension_mime_types, ext_lower, ext_len)
	zend._efree(ext_lower)
	return (*byte)(ret)
}
func ZifApacheRequestHeaders(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var client *PhpCliServerClient
	var headers *zend.HashTable
	var key *zend.ZendString
	var value *byte
	var tmp zend.Zval
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	client = core.sapi_globals.server_context
	headers = &client.request.GetHeadersOriginalCase()
	var __arr *zend.ZendArray = zend._zendNewArray(headers.nNumOfElements)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for {
		var __ht *zend.HashTable = headers
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			key = _p.key
			value = _z.value.ptr
			var _s *byte = value
			var __z *zend.Zval = &tmp
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			zend.ZendSymtableUpdate(return_value.value.arr, key, &tmp)
		}
		break
	}
}

/* }}} */

func AddResponseHeader(h *core.SapiHeader, return_value *zend.Zval) {
	var s *byte
	var p *byte
	var len_ ptrdiff_t
	if h.header_len > 0 {
		p = strchr(h.header, ':')
		len_ = p - h.header
		if p != nil && len_ > 0 {
			for len_ > 0 && (h.header[len_-1] == ' ' || h.header[len_-1] == '\t') {
				len_--
			}
			if len_ {
				s = zend._emalloc(len_ + 1)
				memcpy(s, h.header, len_)
				s[len_] = 0
				for {
					p++
					if !((*p) == ' ' || (*p) == '\t') {
						break
					}
				}
				zend.AddAssocStringlEx(return_value, s, uint32(len_), p, h.header_len-(p-h.header))
				zend._efree(s)
			}
		}
	}
}

/* }}} */

func ZifApacheResponseHeaders(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	zend.ZendLlistApplyWithArgument(&(core.sapi_globals.sapi_headers).headers, zend.LlistApplyWithArgFuncT(AddResponseHeader), return_value)
}

/* }}} */

func CliServerInitGlobals(cg *ZendCliServerGlobals) { cg.SetColor(0) }

var IniEntries []zend.ZendIniEntryDef = []zend.ZendIniEntryDef{
	{
		"cli_server.color",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendCliServerGlobals)(nil).GetColor())) - (*byte)(nil))),
		any(&CliServerGlobals),
		nil,
		"0",
		zend.ZendIniBooleanDisplayerCb,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"cli_server.color\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{nil, nil, nil, nil, nil, nil, nil, 0, 0, 0},
}

func ZmStartupCliServer(type_ int, module_number int) int {
	CliServerInitGlobals(&CliServerGlobals)
	zend.ZendRegisterIniEntries(IniEntries, module_number)
	return zend.SUCCESS
}
func ZmShutdownCliServer(type_ int, module_number int) int {
	zend.ZendUnregisterIniEntries(module_number)
	return zend.SUCCESS
}
func ZmInfoCliServer(zend_module *zend.ZendModuleEntry) { core.DisplayIniEntries(zend_module) }

var CliServerModuleEntry zend.ZendModuleEntry = zend.ZendModuleEntry{g.SizeOf("zend_module_entry"), 20190902, 0, 0, nil, nil, "cli_server", nil, ZmStartupCliServer, ZmShutdownCliServer, nil, nil, ZmInfoCliServer, "7.4.33", 0, nil, nil, nil, nil, 0, 0, nil, 0, "API" + "20190902" + ",NTS"}

/* }}} */

var ArginfoNoArgs []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var ServerAdditionalFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"cli_set_process_title",
		ZifCliSetProcessTitle,
		ArginfoCliSetProcessTitle,
		uint32(g.SizeOf("arginfo_cli_set_process_title")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"cli_get_process_title",
		ZifCliGetProcessTitle,
		ArginfoCliGetProcessTitle,
		uint32(g.SizeOf("arginfo_cli_get_process_title")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"apache_request_headers",
		ZifApacheRequestHeaders,
		ArginfoNoArgs,
		uint32(g.SizeOf("arginfo_no_args")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"apache_response_headers",
		ZifApacheResponseHeaders,
		ArginfoNoArgs,
		uint32(g.SizeOf("arginfo_no_args")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getallheaders",
		ZifApacheRequestHeaders,
		ArginfoNoArgs,
		uint32(g.SizeOf("arginfo_no_args")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

func SapiCliServerStartup(sapi_module *core.sapi_module_struct) int {
	var workers *byte
	if core.PhpModuleStartup(sapi_module, &CliServerModuleEntry, 1) == zend.FAILURE {
		return zend.FAILURE
	}
	if g.Assign(&workers, getenv("PHP_CLI_SERVER_WORKERS")) {
		fprintf(stderr, "platform does not support SO_REUSEPORT, cannot create workers\n")
	}
	return zend.SUCCESS
}
func SapiCliServerUbWrite(str *byte, str_length int) int {
	var client *PhpCliServerClient = core.sapi_globals.server_context
	if client == nil {
		return 0
	}
	return PhpCliServerClientSendThrough(client, str, str_length)
}
func SapiCliServerFlush(server_context any) {
	var client *PhpCliServerClient = server_context
	if client == nil {
		return
	}
	if client.GetSock() < 0 {
		core.PhpHandleAbortedConnection()
		return
	}
	if core.sapi_globals.headers_sent == 0 {
		core.SapiSendHeaders()
		core.sapi_globals.headers_sent = 1
	}
}
func SapiCliServerDiscardHeaders(sapi_headers *core.SapiHeaders) int { return 1 }

/* }}} */

func SapiCliServerSendHeaders(sapi_headers *core.SapiHeaders) int {
	var client *PhpCliServerClient = core.sapi_globals.server_context
	var buffer zend.SmartStr = zend.SmartStr{0}
	var h *core.SapiHeader
	var pos zend.ZendLlistPosition
	if client == nil || core.sapi_globals.request_info.no_headers != 0 {
		return 1
	}
	if core.sapi_globals.sapi_headers.http_status_line != nil {
		zend.SmartStrAppendlEx(&buffer, core.sapi_globals.sapi_headers.http_status_line, strlen(core.sapi_globals.sapi_headers.http_status_line), 0)
		zend.SmartStrAppendlEx(&buffer, "\r\n", 2, 0)
	} else {
		AppendHttpStatusLine(&buffer, client.GetRequest().GetProtocolVersion(), core.sapi_globals.sapi_headers.http_response_code, 0)
	}
	AppendEssentialHeaders(&buffer, client, 0)
	h = (*core.SapiHeader)(zend.ZendLlistGetFirstEx(&sapi_headers.headers, &pos))
	for h != nil {
		if h.header_len != 0 {
			zend.SmartStrAppendlEx(&buffer, h.header, h.header_len, 0)
			zend.SmartStrAppendlEx(&buffer, "\r\n", 2, 0)
		}
		h = (*core.SapiHeader)(zend.ZendLlistGetNextEx(&sapi_headers.headers, &pos))
	}
	zend.SmartStrAppendlEx(&buffer, "\r\n", 2, 0)
	PhpCliServerClientSendThrough(client, buffer.s.val, buffer.s.len_)
	zend.SmartStrFreeEx(&buffer, 0)
	return 1
}

/* }}} */

func SapiCliServerReadCookies() *byte {
	var client *PhpCliServerClient = core.sapi_globals.server_context
	var val *byte
	if nil == g.Assign(&val, zend.ZendHashStrFindPtr(&client.request.GetHeaders(), "cookie", g.SizeOf("\"cookie\"")-1)) {
		return nil
	}
	return val
}
func SapiCliServerReadPost(buf *byte, count_bytes int) int {
	var client *PhpCliServerClient = core.sapi_globals.server_context
	if client.GetRequest().GetContent() != nil {
		var content_len int = client.GetRequest().GetContentLen()
		var nbytes_copied int = g.CondF1(client.GetPostReadOffset()+count_bytes < content_len, func() int { return client.GetPostReadOffset() + count_bytes }, content_len) - client.GetPostReadOffset()
		memmove(buf, client.GetRequest().GetContent()+client.GetPostReadOffset(), nbytes_copied)
		client.SetPostReadOffset(client.GetPostReadOffset() + nbytes_copied)
		return nbytes_copied
	}
	return 0
}
func SapiCliServerRegisterVariable(track_vars_array *zend.Zval, key *byte, val *byte) {
	var new_val *byte = (*byte)(val)
	var new_val_len int
	if nil == val {
		return
	}
	if core.sapi_module.input_filter(5, (*byte)(key), &new_val, strlen(val), &new_val_len) != 0 {
		core.PhpRegisterVariableSafe((*byte)(key), new_val, new_val_len, track_vars_array)
	}
}
func SapiCliServerRegisterEntryCb(entry **byte, num_args int, args ...any, hash_key *zend.ZendHashKey) int {
	var track_vars_array *zend.Zval = __va_arg(args, (*zend.Zval)(_))
	if hash_key.key != nil {
		var real_key *byte
		var key *byte
		var i uint32
		key = zend._estrndup(hash_key.key.val, hash_key.key.len_)
		for i = 0; i < hash_key.key.len_; i++ {
			if key[i] == '-' {
				key[i] = '_'
			} else {
				key[i] = toupper(key[i])
			}
		}
		zend.ZendSpprintf(&real_key, 0, "%s_%s", "HTTP", key)
		if strcmp(key, "CONTENT_TYPE") == 0 || strcmp(key, "CONTENT_LENGTH") == 0 {
			SapiCliServerRegisterVariable(track_vars_array, key, *entry)
		}
		SapiCliServerRegisterVariable(track_vars_array, real_key, *entry)
		zend._efree(key)
		zend._efree(real_key)
	}
	return 0
}

/* }}} */

func SapiCliServerRegisterVariables(track_vars_array *zend.Zval) {
	var client *PhpCliServerClient = core.sapi_globals.server_context
	SapiCliServerRegisterVariable(track_vars_array, "DOCUMENT_ROOT", client.GetServer().GetDocumentRoot())
	var tmp *byte
	if g.Assign(&tmp, strrchr(client.GetAddrStr(), ':')) {
		var addr []byte
		var port []byte
		var addr_start *byte = client.GetAddrStr()
		var addr_end *byte = tmp
		if addr_start[0] == '[' {
			addr_start++
		}
		if addr_end[-1] == ']' {
			addr_end--
		}
		strncpy(port, tmp+1, 8)
		port[7] = '0'
		strncpy(addr, addr_start, addr_end-addr_start)
		addr[addr_end-addr_start] = '0'
		SapiCliServerRegisterVariable(track_vars_array, "REMOTE_ADDR", addr)
		SapiCliServerRegisterVariable(track_vars_array, "REMOTE_PORT", port)
	} else {
		SapiCliServerRegisterVariable(track_vars_array, "REMOTE_ADDR", client.GetAddrStr())
	}
	var tmp *byte
	zend.ZendSpprintf(&tmp, 0, "PHP %s Development Server", "7.4.33")
	SapiCliServerRegisterVariable(track_vars_array, "SERVER_SOFTWARE", tmp)
	zend._efree(tmp)
	var tmp *byte
	zend.ZendSpprintf(&tmp, 0, "HTTP/%d.%d", client.GetRequest().GetProtocolVersion()/100, client.GetRequest().GetProtocolVersion()%100)
	SapiCliServerRegisterVariable(track_vars_array, "SERVER_PROTOCOL", tmp)
	zend._efree(tmp)
	SapiCliServerRegisterVariable(track_vars_array, "SERVER_NAME", client.GetServer().GetHost())
	var tmp *byte
	zend.ZendSpprintf(&tmp, 0, "%i", client.GetServer().GetPort())
	SapiCliServerRegisterVariable(track_vars_array, "SERVER_PORT", tmp)
	zend._efree(tmp)
	SapiCliServerRegisterVariable(track_vars_array, "REQUEST_URI", client.GetRequest().GetRequestUri())
	SapiCliServerRegisterVariable(track_vars_array, "REQUEST_METHOD", core.sapi_globals.request_info.request_method)
	SapiCliServerRegisterVariable(track_vars_array, "SCRIPT_NAME", client.GetRequest().GetVpath())
	if core.sapi_globals.request_info.path_translated != nil {
		SapiCliServerRegisterVariable(track_vars_array, "SCRIPT_FILENAME", core.sapi_globals.request_info.path_translated)
	} else if client.GetServer().GetRouter() != nil {
		SapiCliServerRegisterVariable(track_vars_array, "SCRIPT_FILENAME", client.GetServer().GetRouter())
	}
	if client.GetRequest().GetPathInfo() != nil {
		SapiCliServerRegisterVariable(track_vars_array, "PATH_INFO", client.GetRequest().GetPathInfo())
	}
	if client.GetRequest().GetPathInfoLen() != 0 {
		var tmp *byte
		zend.ZendSpprintf(&tmp, 0, "%s%s", client.GetRequest().GetVpath(), client.GetRequest().GetPathInfo())
		SapiCliServerRegisterVariable(track_vars_array, "PHP_SELF", tmp)
		zend._efree(tmp)
	} else {
		SapiCliServerRegisterVariable(track_vars_array, "PHP_SELF", client.GetRequest().GetVpath())
	}
	if client.GetRequest().GetQueryString() != nil {
		SapiCliServerRegisterVariable(track_vars_array, "QUERY_STRING", client.GetRequest().GetQueryString())
	}
	zend.ZendHashApplyWithArguments(&client.request.GetHeaders(), zend.ApplyFuncArgsT(SapiCliServerRegisterEntryCb), 1, track_vars_array)
}
func SapiCliServerLogWrite(type_ int, msg *byte) {
	var buf []byte
	if PhpCliServerLogLevel < type_ {
		return
	}
	if PhpCliServerGetSystemTime(buf) != 0 {
		memmove(buf, "unknown time, can't be fetched", g.SizeOf("\"unknown time, can't be fetched\""))
	} else {
		var l int = strlen(buf)
		if l > 0 {
			buf[l-1] = '0'
		} else {
			memmove(buf, "unknown", g.SizeOf("\"unknown\""))
		}
	}
	if PhpCliServerWorkersMax > 1 {
		fprintf(stderr, "[%ld] [%s] %s\n", long(getpid()), buf, msg)
	} else {
		fprintf(stderr, "[%s] %s\n", buf, msg)
	}
}
func SapiCliServerLogMessage(msg *byte, syslog_type_int int) { SapiCliServerLogWrite(3, msg) }

/* {{{ sapi_module_struct cli_server_sapi_module
 */

var CliServerSapiModule core.sapi_module_struct = core.sapi_module_struct{"cli-server", "Built-in HTTP server", SapiCliServerStartup, core.PhpModuleShutdownWrapper, nil, nil, SapiCliServerUbWrite, SapiCliServerFlush, nil, nil, zend.ZendError, nil, SapiCliServerSendHeaders, nil, SapiCliServerReadPost, SapiCliServerReadCookies, SapiCliServerRegisterVariables, SapiCliServerLogMessage, nil, nil, nil, nil, nil, nil, 0, 0, nil, nil, nil, nil, nil, nil, 0, nil, nil, nil}

func PhpCliServerPollerCtor(poller *PhpCliServerPoller) int {
	FD_ZERO(&poller.rfds)
	FD_ZERO(&poller.wfds)
	poller.SetMaxFd(-1)
	return zend.SUCCESS
}
func PhpCliServerPollerAdd(poller *PhpCliServerPoller, mode int, fd core.PhpSocketT) {
	if (mode & POLLIN) != 0 {
		if fd < FD_SETSIZE {
			FD_SET(fd, &poller.rfds)
		}
	}
	if (mode & POLLOUT) != 0 {
		if fd < FD_SETSIZE {
			FD_SET(fd, &poller.wfds)
		}
	}
	if fd > poller.GetMaxFd() {
		poller.SetMaxFd(fd)
	}
}
func PhpCliServerPollerRemove(poller *PhpCliServerPoller, mode int, fd core.PhpSocketT) {
	if (mode & POLLIN) != 0 {
		if fd < FD_SETSIZE {
			FD_CLR(fd, &poller.rfds)
		}
	}
	if (mode & POLLOUT) != 0 {
		if fd < FD_SETSIZE {
			FD_CLR(fd, &poller.wfds)
		}
	}
	if fd == poller.GetMaxFd() {
		for fd > 0 {
			fd--
			if fd < FD_SETSIZE && FD_ISSET(fd, &poller.rfds) || fd < FD_SETSIZE && FD_ISSET(fd, &poller.wfds) {
				break
			}
		}
		poller.SetMaxFd(fd)
	}
}
func PhpCliServerPollerPoll(poller *PhpCliServerPoller, tv *__struct__timeval) int {
	memmove(&poller.active.rfds, &poller.rfds, g.SizeOf("poller -> rfds"))
	memmove(&poller.active.wfds, &poller.wfds, g.SizeOf("poller -> wfds"))
	return select_(poller.GetMaxFd()+1, &poller.active.rfds, &poller.active.wfds, nil, tv)
}
func PhpCliServerPollerIterOnActive(poller *PhpCliServerPoller, opaque any, callback func(_ any, fd core.PhpSocketT, events int) int) int {
	var retval int = zend.SUCCESS
	var fd core.PhpSocketT
	var max_fd core.PhpSocketT = poller.GetMaxFd()
	for fd = 0; fd <= max_fd; fd++ {
		if fd < FD_SETSIZE && FD_ISSET(fd, &poller.active.rfds) {
			if zend.SUCCESS != callback(opaque, fd, POLLIN) {
				retval = zend.FAILURE
			}
		}
		if fd < FD_SETSIZE && FD_ISSET(fd, &poller.active.wfds) {
			if zend.SUCCESS != callback(opaque, fd, POLLOUT) {
				retval = zend.FAILURE
			}
		}
	}
	return retval
}
func PhpCliServerChunkSize(chunk *PhpCliServerChunk) int {
	switch chunk.GetType() {
	case PHP_CLI_SERVER_CHUNK_HEAP:
		return chunk.GetDataHeapLen()
	case PHP_CLI_SERVER_CHUNK_IMMORTAL:
		return chunk.GetDataImmortalLen()
	}
	return 0
}
func PhpCliServerChunkDtor(chunk *PhpCliServerChunk) {
	switch chunk.GetType() {
	case PHP_CLI_SERVER_CHUNK_HEAP:
		if chunk.GetBlock() != chunk {
			g.CondF(true, func() { return zend.Free(chunk.GetBlock()) }, func() { return zend._efree(chunk.GetBlock()) })
		}
		break
	case PHP_CLI_SERVER_CHUNK_IMMORTAL:
		break
	}
}
func PhpCliServerBufferDtor(buffer *PhpCliServerBuffer) {
	var chunk *PhpCliServerChunk
	var next *PhpCliServerChunk
	for chunk = buffer.GetFirst(); chunk != nil; chunk = next {
		next = chunk.GetNext()
		PhpCliServerChunkDtor(chunk)
		g.CondF(true, func() { return zend.Free(chunk) }, func() { return zend._efree(chunk) })
	}
}
func PhpCliServerBufferCtor(buffer *PhpCliServerBuffer) {
	buffer.SetFirst(nil)
	buffer.SetLast(nil)
}
func PhpCliServerBufferAppend(buffer *PhpCliServerBuffer, chunk *PhpCliServerChunk) {
	var last *PhpCliServerChunk
	for last = chunk; last.GetNext() != nil; last = last.GetNext() {

	}
	if buffer.GetLast() == nil {
		buffer.SetFirst(chunk)
	} else {
		buffer.GetLast().SetNext(chunk)
	}
	buffer.SetLast(last)
}
func PhpCliServerBufferPrepend(buffer *PhpCliServerBuffer, chunk *PhpCliServerChunk) {
	var last *PhpCliServerChunk
	for last = chunk; last.GetNext() != nil; last = last.GetNext() {

	}
	last.SetNext(buffer.GetFirst())
	if buffer.GetLast() == nil {
		buffer.SetLast(last)
	}
	buffer.SetFirst(chunk)
}
func PhpCliServerBufferSize(buffer *PhpCliServerBuffer) int {
	var chunk *PhpCliServerChunk
	var retval int = 0
	for chunk = buffer.GetFirst(); chunk != nil; chunk = chunk.GetNext() {
		retval += PhpCliServerChunkSize(chunk)
	}
	return retval
}
func PhpCliServerChunkImmortalNew(buf *byte, len_ int) *PhpCliServerChunk {
	var chunk *PhpCliServerChunk = g.CondF(true, func() any { return zend.__zendMalloc(g.SizeOf("php_cli_server_chunk")) }, func() any { return zend._emalloc(g.SizeOf("php_cli_server_chunk")) })
	chunk.SetType(PHP_CLI_SERVER_CHUNK_IMMORTAL)
	chunk.SetNext(nil)
	chunk.SetDataImmortalP(buf)
	chunk.SetDataImmortalLen(len_)
	return chunk
}
func PhpCliServerChunkHeapNew(block any, buf *byte, len_ int) *PhpCliServerChunk {
	var chunk *PhpCliServerChunk = g.CondF(true, func() any { return zend.__zendMalloc(g.SizeOf("php_cli_server_chunk")) }, func() any { return zend._emalloc(g.SizeOf("php_cli_server_chunk")) })
	chunk.SetType(PHP_CLI_SERVER_CHUNK_HEAP)
	chunk.SetNext(nil)
	chunk.SetBlock(block)
	chunk.SetDataHeapP(buf)
	chunk.SetDataHeapLen(len_)
	return chunk
}
func PhpCliServerChunkHeapNewSelfContained(len_ int) *PhpCliServerChunk {
	var chunk *PhpCliServerChunk = g.CondF(true, func() any { return zend.__zendMalloc(g.SizeOf("php_cli_server_chunk") + len_) }, func() any { return zend._emalloc(g.SizeOf("php_cli_server_chunk") + len_) })
	chunk.SetType(PHP_CLI_SERVER_CHUNK_HEAP)
	chunk.SetNext(nil)
	chunk.SetBlock(chunk)
	chunk.SetDataHeapP((*byte)(chunk + 1))
	chunk.SetDataHeapLen(len_)
	return chunk
}
func PhpCliServerContentSenderDtor(sender *PhpCliServerContentSender) {
	PhpCliServerBufferDtor(&sender.buffer)
}
func PhpCliServerContentSenderCtor(sender *PhpCliServerContentSender) {
	PhpCliServerBufferCtor(&sender.buffer)
}
func PhpCliServerContentSenderSend(sender *PhpCliServerContentSender, fd core.PhpSocketT, nbytes_sent_total *int) int {
	var chunk *PhpCliServerChunk
	var next *PhpCliServerChunk
	var _nbytes_sent_total int = 0
	for chunk = sender.GetBuffer().GetFirst(); chunk != nil; chunk = next {
		var nbytes_sent ssize_t
		next = chunk.GetNext()
		switch chunk.GetType() {
		case PHP_CLI_SERVER_CHUNK_HEAP:
			nbytes_sent = send(fd, chunk.GetDataHeapP(), chunk.GetDataHeapLen(), 0)
			if nbytes_sent < 0 {
				*nbytes_sent_total = _nbytes_sent_total
				return errno
			} else if nbytes_sent == ssize_t(chunk.GetDataHeapLen()) {
				PhpCliServerChunkDtor(chunk)
				g.CondF(true, func() { return zend.Free(chunk) }, func() { return zend._efree(chunk) })
				sender.GetBuffer().SetFirst(next)
				if next == nil {
					sender.GetBuffer().SetLast(nil)
				}
			} else {
				chunk.SetDataHeapP(chunk.GetDataHeapP() + nbytes_sent)
				chunk.SetDataHeapLen(chunk.GetDataHeapLen() - nbytes_sent)
			}
			_nbytes_sent_total += nbytes_sent
			break
		case PHP_CLI_SERVER_CHUNK_IMMORTAL:
			nbytes_sent = send(fd, chunk.GetDataImmortalP(), chunk.GetDataImmortalLen(), 0)
			if nbytes_sent < 0 {
				*nbytes_sent_total = _nbytes_sent_total
				return errno
			} else if nbytes_sent == ssize_t(chunk.GetDataImmortalLen()) {
				PhpCliServerChunkDtor(chunk)
				g.CondF(true, func() { return zend.Free(chunk) }, func() { return zend._efree(chunk) })
				sender.GetBuffer().SetFirst(next)
				if next == nil {
					sender.GetBuffer().SetLast(nil)
				}
			} else {
				chunk.SetDataImmortalP(chunk.GetDataImmortalP() + nbytes_sent)
				chunk.SetDataImmortalLen(chunk.GetDataImmortalLen() - nbytes_sent)
			}
			_nbytes_sent_total += nbytes_sent
			break
		}
	}
	*nbytes_sent_total = _nbytes_sent_total
	return 0
}
func PhpCliServerContentSenderPull(sender *PhpCliServerContentSender, fd int, nbytes_read *int) int {
	var _nbytes_read ssize_t
	var chunk *PhpCliServerChunk = PhpCliServerChunkHeapNewSelfContained(131072)
	_nbytes_read = read(fd, chunk.GetDataHeapP(), chunk.GetDataHeapLen())
	if _nbytes_read < 0 {
		if PhpCliServerLogLevel >= 2 {
			var errstr *byte = GetLastError()
			PhpCliServerLogf(2, "%s", errstr)
			g.CondF(true, func() { return zend.Free(errstr) }, func() { return zend._efree(errstr) })
		}
		PhpCliServerChunkDtor(chunk)
		g.CondF(true, func() { return zend.Free(chunk) }, func() { return zend._efree(chunk) })
		return 1
	}
	chunk.SetDataHeapLen(_nbytes_read)
	PhpCliServerBufferAppend(&sender.buffer, chunk)
	*nbytes_read = _nbytes_read
	return 0
}
func PhpCliIsOutputTty() int {
	if PhpCliOutputIsTty == -1 {
		PhpCliOutputIsTty = zend.Isatty(STDOUT_FILENO)
	}
	return PhpCliOutputIsTty
}
func PhpCliServerLogResponse(client *PhpCliServerClient, status int, message *byte) {
	var color int = 0
	var effective_status int = status
	var basic_buf *byte
	var message_buf *byte = ""
	var error_buf *byte = ""
	var append_error_message zend.ZendBool = 0
	if core.CoreGlobals.last_error_message != nil {
		switch core.CoreGlobals.last_error_type {
		case 1 << 0:

		case 1 << 4:

		case 1 << 6:

		case 1 << 8:

		case 1 << 2:
			if status == 200 {

				/* the status code isn't changed by a fatal error, so fake it */

				effective_status = 500

				/* the status code isn't changed by a fatal error, so fake it */

			}
			append_error_message = 1
			break
		}
	}
	if CliServerGlobals.GetColor() && PhpCliIsOutputTty() == 1 {
		if effective_status >= 500 {

			/* server error: red */

			color = 1

			/* server error: red */

		} else if effective_status >= 400 {

			/* client error: yellow */

			color = 3

			/* client error: yellow */

		} else if effective_status >= 200 {

			/* success: green */

			color = 2

			/* success: green */

		}
	}

	/* basic */

	zend.ZendSpprintf(&basic_buf, 0, "%s [%d]: %s %s", client.GetAddrStr(), status, core.sapi_globals.request_info.request_method, client.GetRequest().GetRequestUri())
	if basic_buf == nil {
		return
	}

	/* message */

	if message != nil {
		zend.ZendSpprintf(&message_buf, 0, " - %s", message)
		if message_buf == nil {
			zend._efree(basic_buf)
			return
		}
	}

	/* error */

	if append_error_message != 0 {
		zend.ZendSpprintf(&error_buf, 0, " - %s in %s on line %d", core.CoreGlobals.last_error_message, core.CoreGlobals.last_error_file, core.CoreGlobals.last_error_lineno)
		if error_buf == nil {
			zend._efree(basic_buf)
			if message != nil {
				zend._efree(message_buf)
			}
			return
		}
	}
	if color != 0 {
		PhpCliServerLogf(3, "x1b[3%dm%s%s%sx1b[0m", color, basic_buf, message_buf, error_buf)
	} else {
		PhpCliServerLogf(3, "%s%s%s", basic_buf, message_buf, error_buf)
	}
	zend._efree(basic_buf)
	if message != nil {
		zend._efree(message_buf)
	}
	if append_error_message != 0 {
		zend._efree(error_buf)
	}
}
func PhpCliServerLogf(type_ int, format string, _ ...any) {
	var buf *byte = nil
	var ap va_list
	if PhpCliServerLogLevel < type_ {
		return
	}
	va_start(ap, format)
	zend.ZendVspprintf(&buf, 0, format, ap)
	va_end(ap)
	if buf == nil {
		return
	}
	SapiCliServerLogWrite(type_, buf)
	zend._efree(buf)
}
func PhpNetworkListenSocket(host *byte, port *int, socktype int, af *int, socklen *socklen_t, errstr **zend.ZendString) core.PhpSocketT {
	var retval core.PhpSocketT = -1
	var err int = 0
	var sa *__struct__sockaddr = nil
	var p **__struct__sockaddr
	var sal **__struct__sockaddr
	var num_addrs int = core.PhpNetworkGetaddresses(host, socktype, &sal, errstr)
	if num_addrs == 0 {
		return -1
	}
	for p = sal; (*p) != nil; p++ {
		if sa != nil {
			g.CondF(true, func() { return zend.Free(sa) }, func() { return zend._efree(sa) })
			sa = nil
		}
		retval = socket((*p).sa_family, socktype, 0)
		if retval == -1 {
			continue
		}
		switch (*p).sa_family {
		case AF_INET6:
			sa = zend.__zendMalloc(g.SizeOf("struct sockaddr_in6"))
			*((*__struct__sockaddr_in6)(sa)) = *((*__struct__sockaddr_in6)(*p))
			(*__struct__sockaddr_in6)(sa).sin6_port = htons(*port)
			*socklen = g.SizeOf("struct sockaddr_in6")
			break
		case AF_INET:
			sa = zend.__zendMalloc(g.SizeOf("struct sockaddr_in"))
			*((*__struct__sockaddr_in)(sa)) = *((*__struct__sockaddr_in)(*p))
			(*__struct__sockaddr_in)(sa).sin_port = htons(*port)
			*socklen = g.SizeOf("struct sockaddr_in")
			break
		default:

			/* Unknown family */

			*socklen = 0
			close(retval)
			continue
		}
		if bind(retval, sa, *socklen) == -1 {
			err = errno
			if err == EINVAL || err == EADDRINUSE {
				goto out
			}
			close(retval)
			retval = -1
			continue
		}
		err = 0
		*af = sa.sa_family
		if (*port) == 0 {
			if getsockname(retval, sa, socklen) {
				err = errno
				goto out
			}
			switch sa.sa_family {
			case AF_INET6:
				*port = ntohs((*__struct__sockaddr_in6)(sa).sin6_port)
				break
			case AF_INET:
				*port = ntohs((*__struct__sockaddr_in)(sa).sin_port)
				break
			}
		}
		break
	}
	if retval == -1 {
		goto out
	}
	if listen(retval, SOMAXCONN) {
		err = errno
		goto out
	}
out:
	if sa != nil {
		g.CondF(true, func() { return zend.Free(sa) }, func() { return zend._efree(sa) })
	}
	if sal != nil {
		core.PhpNetworkFreeaddresses(sal)
	}
	if err != 0 {
		if retval >= 0 {
			close(retval)
		}
		if errstr != nil {
			*errstr = core.PhpSocketErrorStr(err)
		}
		return -1
	}
	return retval
}
func PhpCliServerRequestCtor(req *PhpCliServerRequest) int {
	req.SetProtocolVersion(0)
	req.SetRequestUri(nil)
	req.SetRequestUriLen(0)
	req.SetVpath(nil)
	req.SetVpathLen(0)
	req.SetPathTranslated(nil)
	req.SetPathTranslatedLen(0)
	req.SetPathInfo(nil)
	req.SetPathInfoLen(0)
	req.SetQueryString(nil)
	req.SetQueryStringLen(0)
	zend._zendHashInit(&req.headers, 0, CharPtrDtorP, 1)
	zend._zendHashInit(&req.headers_original_case, 0, nil, 1)
	req.SetContent(nil)
	req.SetContentLen(0)
	req.SetExt(nil)
	req.SetExtLen(0)
	return zend.SUCCESS
}
func PhpCliServerRequestDtor(req *PhpCliServerRequest) {
	if req.GetRequestUri() != nil {
		g.CondF(true, func() { return zend.Free(req.GetRequestUri()) }, func() { return zend._efree(req.GetRequestUri()) })
	}
	if req.GetVpath() != nil {
		g.CondF(true, func() { return zend.Free(req.GetVpath()) }, func() { return zend._efree(req.GetVpath()) })
	}
	if req.GetPathTranslated() != nil {
		g.CondF(true, func() { return zend.Free(req.GetPathTranslated()) }, func() { return zend._efree(req.GetPathTranslated()) })
	}
	if req.GetPathInfo() != nil {
		g.CondF(true, func() { return zend.Free(req.GetPathInfo()) }, func() { return zend._efree(req.GetPathInfo()) })
	}
	if req.GetQueryString() != nil {
		g.CondF(true, func() { return zend.Free(req.GetQueryString()) }, func() { return zend._efree(req.GetQueryString()) })
	}
	zend.ZendHashDestroy(&req.headers)
	zend.ZendHashDestroy(&req.headers_original_case)
	if req.GetContent() != nil {
		g.CondF(true, func() { return zend.Free(req.GetContent()) }, func() { return zend._efree(req.GetContent()) })
	}
}
func PhpCliServerRequestTranslateVpath(request *PhpCliServerRequest, document_root *byte, document_root_len int) {
	var sb zend.ZendStatT
	var index_files []*byte = []*byte{"index.php", "index.html", nil}
	var buf *byte = g.CondF(true, func() any {
		return zend._safeMalloc(1, request.GetVpathLen(), 1+document_root_len+1+g.SizeOf("\"index.html\""))
	}, func() any {
		return zend._safeEmalloc(1, request.GetVpathLen(), 1+document_root_len+1+g.SizeOf("\"index.html\""))
	})
	var p *byte = buf
	var prev_path *byte = nil
	var q *byte
	var vpath *byte
	var prev_path_len int = 0
	var is_static_file int = 0
	memmove(p, document_root, document_root_len)
	p += document_root_len
	vpath = p
	if request.GetVpathLen() > 0 && request.GetVpath()[0] != '/' {
		g.PostInc(&(*p)) = '/'
	}
	q = request.GetVpath() + request.GetVpathLen()
	for q > request.GetVpath() {
		if g.PostDec(&(*q)) == '.' {
			is_static_file = 1
			break
		}
	}
	memmove(p, request.GetVpath(), request.GetVpathLen())
	p += request.GetVpathLen()
	*p = '0'
	q = p
	for q > buf {
		if !(stat(buf, &sb)) {
			if (sb.st_mode & S_IFDIR) != 0 {
				var file **byte = index_files
				if q[-1] != '/' {
					g.PostInc(&(*q)) = '/'
				}
				for (*file) != nil {
					var l int = strlen(*file)
					memmove(q, *file, l+1)
					if !(stat(buf, &sb)) && (sb.st_mode&S_IFREG) != 0 {
						q += l
						break
					}
					file++
				}
				if (*file) == nil || is_static_file != 0 {
					if prev_path != nil {
						g.CondF(true, func() { return zend.Free(prev_path) }, func() { return zend._efree(prev_path) })
					}
					g.CondF(true, func() { return zend.Free(buf) }, func() { return zend._efree(buf) })
					return
				}
			}
			break
		}
		if prev_path != nil {
			g.CondF(true, func() { return zend.Free(prev_path) }, func() { return zend._efree(prev_path) })
			*q = '/'
		}
		for q > buf && (*(g.PreDec(&q))) != '/' {

		}
		prev_path_len = p - q
		prev_path = zend.ZendStrndup(q, prev_path_len)
		*q = '0'
	}
	if prev_path != nil {
		request.SetPathInfoLen(prev_path_len)
		request.SetPathInfo(prev_path)
		g.CondF(true, func() { return zend.Free(request.GetVpath()) }, func() { return zend._efree(request.GetVpath()) })
		request.SetVpath(zend.ZendStrndup(vpath, q-vpath))
		request.SetVpathLen(q - vpath)
		request.SetPathTranslated(buf)
		request.SetPathTranslatedLen(q - buf)
	} else {
		g.CondF(true, func() { return zend.Free(request.GetVpath()) }, func() { return zend._efree(request.GetVpath()) })
		request.SetVpath(zend.ZendStrndup(vpath, q-vpath))
		request.SetVpathLen(q - vpath)
		request.SetPathTranslated(buf)
		request.SetPathTranslatedLen(q - buf)
	}
	request.SetSb(sb)
}
func NormalizeVpath(retval **byte, retval_len *int, vpath *byte, vpath_len int, persistent int) {
	var decoded_vpath *byte = nil
	var decoded_vpath_end *byte
	var p *byte
	*retval = nil
	*retval_len = 0
	if persistent != 0 {
		decoded_vpath = zend.ZendStrndup(vpath, vpath_len)
	} else {
		decoded_vpath = zend._estrndup(vpath, vpath_len)
	}
	if decoded_vpath == nil {
		return
	}
	decoded_vpath_end = decoded_vpath + standard.PhpRawUrlDecode(decoded_vpath, int(vpath_len))
	p = decoded_vpath
	if p < decoded_vpath_end && (*p) == '/' {
		var n *byte = p
		for n < decoded_vpath_end && (*n) == '/' {
			n++
		}
		memmove(g.PreInc(&p), n, decoded_vpath_end-n)
		decoded_vpath_end -= n - p
	}
	for p < decoded_vpath_end {
		var n *byte = p
		for n < decoded_vpath_end && (*n) != '/' {
			n++
		}
		if n-p == 2 && p[0] == '.' && p[1] == '.' {
			if p > decoded_vpath {
				p--
				for {
					if p == decoded_vpath {
						if (*p) == '/' {
							p++
						}
						break
					}
					if (*(g.PreDec(&p))) == '/' {
						p++
						break
					}
				}
			}
			for n < decoded_vpath_end && (*n) == '/' {
				n++
			}
			memmove(p, n, decoded_vpath_end-n)
			decoded_vpath_end -= n - p
		} else if n-p == 1 && p[0] == '.' {
			for n < decoded_vpath_end && (*n) == '/' {
				n++
			}
			memmove(p, n, decoded_vpath_end-n)
			decoded_vpath_end -= n - p
		} else {
			if n < decoded_vpath_end {
				var nn *byte = n
				for nn < decoded_vpath_end && (*nn) == '/' {
					nn++
				}
				p = n + 1
				memmove(p, nn, decoded_vpath_end-nn)
				decoded_vpath_end -= nn - p
			} else {
				p = n
			}
		}
	}
	*decoded_vpath_end = '0'
	*retval = decoded_vpath
	*retval_len = decoded_vpath_end - decoded_vpath
}

/* {{{ php_cli_server_client_read_request */

func PhpCliServerClientReadRequestOnMessageBegin(parser *PhpHttpParser) int { return 0 }
func PhpCliServerClientReadRequestOnPath(parser *PhpHttpParser, at *byte, length int) int {
	var client *PhpCliServerClient = parser.GetData()
	var vpath *byte
	var vpath_len int
	if client.GetRequest().GetVpath() != nil {
		return 1
	}
	NormalizeVpath(&vpath, &vpath_len, at, length, 1)
	client.GetRequest().SetVpath(vpath)
	client.GetRequest().SetVpathLen(vpath_len)
	return 0
}
func PhpCliServerClientReadRequestOnQueryString(parser *PhpHttpParser, at *byte, length int) int {
	var client *PhpCliServerClient = parser.GetData()
	if client.GetRequest().GetQueryString() == nil {
		client.GetRequest().SetQueryString(zend.ZendStrndup(at, length))
		client.GetRequest().SetQueryStringLen(length)
	} else {
		assert(length <= 80*1024 && 80*1024-length >= client.GetRequest().GetQueryStringLen())
		client.GetRequest().SetQueryString(zend.__zendRealloc(client.GetRequest().GetQueryString(), client.GetRequest().GetQueryStringLen()+length+1))
		memcpy(client.GetRequest().GetQueryString()+client.GetRequest().GetQueryStringLen(), at, length)
		client.GetRequest().SetQueryStringLen(client.GetRequest().GetQueryStringLen() + length)
		client.GetRequest().GetQueryString()[client.GetRequest().GetQueryStringLen()] = '0'
	}
	return 0
}
func PhpCliServerClientReadRequestOnUrl(parser *PhpHttpParser, at *byte, length int) int {
	var client *PhpCliServerClient = parser.GetData()
	if client.GetRequest().GetRequestUri() == nil {
		client.GetRequest().SetRequestMethod(parser.GetMethod())
		client.GetRequest().SetRequestUri(zend.ZendStrndup(at, length))
		client.GetRequest().SetRequestUriLen(length)
	} else {
		assert(client.GetRequest().GetRequestMethod() == parser.GetMethod())
		assert(length <= 80*1024 && 80*1024-length >= client.GetRequest().GetQueryStringLen())
		client.GetRequest().SetRequestUri(zend.__zendRealloc(client.GetRequest().GetRequestUri(), client.GetRequest().GetRequestUriLen()+length+1))
		memcpy(client.GetRequest().GetRequestUri()+client.GetRequest().GetRequestUriLen(), at, length)
		client.GetRequest().SetRequestUriLen(client.GetRequest().GetRequestUriLen() + length)
		client.GetRequest().GetRequestUri()[client.GetRequest().GetRequestUriLen()] = '0'
	}
	return 0
}
func PhpCliServerClientReadRequestOnFragment(parser *PhpHttpParser, at *byte, length int) int {
	return 0
}
func PhpCliServerClientSaveHeader(client *PhpCliServerClient) {
	/* strip off the colon */

	var orig_header_name *zend.ZendString = zend.ZendStringInit(client.GetCurrentHeaderName(), client.GetCurrentHeaderNameLen(), 1)
	var lc_header_name *zend.ZendString = zend.ZendStringAlloc(client.GetCurrentHeaderNameLen(), 1)
	zend.ZendStrTolowerCopy(lc_header_name.val, client.GetCurrentHeaderName(), client.GetCurrentHeaderNameLen())

	zend.ZendHashAddPtr(&client.request.GetHeaders(), lc_header_name, client.GetCurrentHeaderValue())
	zend.ZendHashAddPtr(&client.request.GetHeadersOriginalCase(), orig_header_name, client.GetCurrentHeaderValue())
	zend.ZendStringReleaseEx(lc_header_name, 1)
	zend.ZendStringReleaseEx(orig_header_name, 1)
	if client.GetCurrentHeaderNameAllocated() != 0 {
		g.CondF(true, func() { return zend.Free(client.GetCurrentHeaderName()) }, func() { return zend._efree(client.GetCurrentHeaderName()) })
		client.SetCurrentHeaderNameAllocated(0)
	}
	client.SetCurrentHeaderName(nil)
	client.SetCurrentHeaderNameLen(0)
	client.SetCurrentHeaderValue(nil)
	client.SetCurrentHeaderValueLen(0)
}
func PhpCliServerClientReadRequestOnHeaderField(parser *PhpHttpParser, at *byte, length int) int {
	var client *PhpCliServerClient = parser.GetData()
	switch client.GetLastHeaderElement() {
	case HEADER_VALUE:
		PhpCliServerClientSaveHeader(client)
	case HEADER_NONE:
		client.SetCurrentHeaderName((*byte)(at))
		client.SetCurrentHeaderNameLen(length)
		break
	case HEADER_FIELD:
		if client.GetCurrentHeaderNameAllocated() != 0 {
			var new_length int = client.GetCurrentHeaderNameLen() + length
			client.SetCurrentHeaderName(zend.__zendRealloc(client.GetCurrentHeaderName(), new_length+1))
			memcpy(client.GetCurrentHeaderName()+client.GetCurrentHeaderNameLen(), at, length)
			client.GetCurrentHeaderName()[new_length] = '0'
			client.SetCurrentHeaderNameLen(new_length)
		} else {
			var new_length int = client.GetCurrentHeaderNameLen() + length
			var field *byte = g.CondF(true, func() any { return zend.__zendMalloc(new_length + 1) }, func() any { return zend._emalloc(new_length + 1) })
			memcpy(field, client.GetCurrentHeaderName(), client.GetCurrentHeaderNameLen())
			memcpy(field+client.GetCurrentHeaderNameLen(), at, length)
			field[new_length] = '0'
			client.SetCurrentHeaderName(field)
			client.SetCurrentHeaderNameLen(new_length)
			client.SetCurrentHeaderNameAllocated(1)
		}
		break
	}
	client.SetLastHeaderElement(HEADER_FIELD)
	return 0
}
func PhpCliServerClientReadRequestOnHeaderValue(parser *PhpHttpParser, at *byte, length int) int {
	var client *PhpCliServerClient = parser.GetData()
	switch client.GetLastHeaderElement() {
	case HEADER_FIELD:
		client.SetCurrentHeaderValue(zend.ZendStrndup(at, length))
		client.SetCurrentHeaderValueLen(length)
		break
	case HEADER_VALUE:
		var new_length int = client.GetCurrentHeaderValueLen() + length
		client.SetCurrentHeaderValue(zend.__zendRealloc(client.GetCurrentHeaderValue(), new_length+1))
		memcpy(client.GetCurrentHeaderValue()+client.GetCurrentHeaderValueLen(), at, length)
		client.GetCurrentHeaderValue()[new_length] = '0'
		client.SetCurrentHeaderValueLen(new_length)
		break
	case HEADER_NONE:

		// can't happen

		assert(false)
		break
	}
	client.SetLastHeaderElement(HEADER_VALUE)
	return 0
}
func PhpCliServerClientReadRequestOnHeadersComplete(parser *PhpHttpParser) int {
	var client *PhpCliServerClient = parser.GetData()
	switch client.GetLastHeaderElement() {
	case HEADER_NONE:
		break
	case HEADER_FIELD:
		client.SetCurrentHeaderValue(zend.__zendMalloc(1))
		(*client).current_header_value = '0'
		client.SetCurrentHeaderValueLen(0)
	case HEADER_VALUE:
		PhpCliServerClientSaveHeader(client)
		break
	}
	client.SetLastHeaderElement(HEADER_NONE)
	return 0
}
func PhpCliServerClientReadRequestOnBody(parser *PhpHttpParser, at *byte, length int) int {
	var client *PhpCliServerClient = parser.GetData()
	if client.GetRequest().GetContent() == nil {
		client.GetRequest().SetContent(zend.__zendMalloc(parser.GetContentLength()))
		client.GetRequest().SetContentLen(0)
	}
	client.GetRequest().SetContent(zend.__zendRealloc(client.GetRequest().GetContent(), client.GetRequest().GetContentLen()+length))
	memmove(client.GetRequest().GetContent()+client.GetRequest().GetContentLen(), at, length)
	client.GetRequest().SetContentLen(client.GetRequest().GetContentLen() + length)
	return 0
}
func PhpCliServerClientReadRequestOnMessageComplete(parser *PhpHttpParser) int {
	var client *PhpCliServerClient = parser.GetData()
	client.GetRequest().SetProtocolVersion(parser.GetHttpMajor()*100 + parser.GetHttpMinor())
	PhpCliServerRequestTranslateVpath(&client.request, client.GetServer().GetDocumentRoot(), client.GetServer().GetDocumentRootLen())
	var vpath *byte = client.GetRequest().GetVpath()
	var end *byte = vpath + client.GetRequest().GetVpathLen()
	var p *byte = end
	client.GetRequest().SetExt(end)
	client.GetRequest().SetExtLen(0)
	for p > vpath {
		p--
		if (*p) == '.' {
			p++
			client.GetRequest().SetExt(p)
			client.GetRequest().SetExtLen(end - p)
			break
		}
	}
	client.SetRequestRead(1)
	return 0
}
func PhpCliServerClientReadRequest(client *PhpCliServerClient, errstr **byte) int {
	var buf []byte
	var settings PhpHttpParserSettings = PhpHttpParserSettings{
		PhpCliServerClientReadRequestOnMessageBegin,
		PhpCliServerClientReadRequestOnPath,
		PhpCliServerClientReadRequestOnQueryString,
		PhpCliServerClientReadRequestOnUrl,
		PhpCliServerClientReadRequestOnFragment,
		PhpCliServerClientReadRequestOnHeaderField,
		PhpCliServerClientReadRequestOnHeaderValue,
		PhpCliServerClientReadRequestOnHeadersComplete,
		PhpCliServerClientReadRequestOnBody,
		PhpCliServerClientReadRequestOnMessageComplete,
	}
	var nbytes_consumed int
	var nbytes_read int
	if client.GetRequestRead() != 0 {
		return 1
	}
	nbytes_read = recv(client.GetSock(), buf, g.SizeOf("buf")-1, 0)
	if nbytes_read < 0 {
		var err int = errno
		if err == EAGAIN {
			return 0
		}
		if PhpCliServerLogLevel >= 2 {
			*errstr = core.PhpSocketStrerror(err, nil, 0)
		}
		return -1
	} else if nbytes_read == 0 {
		if PhpCliServerLogLevel >= 2 {
			*errstr = zend._estrdup(PhpCliServerRequestErrorUnexpectedEof)
		}
		return -1
	}
	client.GetParser().SetData(client)
	nbytes_consumed = PhpHttpParserExecute(&client.parser, &settings, buf, nbytes_read)
	if nbytes_consumed != int(nbytes_read) {
		if PhpCliServerLogLevel >= 2 {
			if (buf[0]&0x80) != 0 || buf[0] == 0x16 {
				*errstr = zend._estrdup("Unsupported SSL request")
			} else {
				*errstr = zend._estrdup("Malformed HTTP request")
			}
		}
		return -1
	}
	if client.GetCurrentHeaderName() != nil {
		var header_name *byte = g.CondF(true, func() any { return zend._safeMalloc(client.GetCurrentHeaderNameLen(), 1, 1) }, func() any { return zend._safeEmalloc(client.GetCurrentHeaderNameLen(), 1, 1) })
		memmove(header_name, client.GetCurrentHeaderName(), client.GetCurrentHeaderNameLen())
		client.SetCurrentHeaderName(header_name)
		client.SetCurrentHeaderNameAllocated(1)
	}
	if client.GetRequestRead() != 0 {
		return 1
	} else {
		return 0
	}
}

/* }}} */

func PhpCliServerClientSendThrough(client *PhpCliServerClient, str *byte, str_len int) int {
	var tv __struct__timeval = __struct__timeval{10, 0}
	var nbytes_left ssize_t = ssize_t(str_len)
	for {
		var nbytes_sent ssize_t
		nbytes_sent = send(client.GetSock(), str+str_len-nbytes_left, nbytes_left, 0)
		if nbytes_sent < 0 {
			var err int = errno
			if err == EAGAIN {
				var nfds int = core.PhpPollfdFor(client.GetSock(), POLLOUT, &tv)
				if nfds > 0 {
					continue
				} else if nfds < 0 {

					/* error */

					core.PhpHandleAbortedConnection()
					return nbytes_left
				} else {

					/* timeout */

					core.PhpHandleAbortedConnection()
					return nbytes_left
				}
			} else {
				core.PhpHandleAbortedConnection()
				return nbytes_left
			}
		}
		nbytes_left -= nbytes_sent
		if nbytes_left <= 0 {
			break
		}
	}
	return str_len
}
func PhpCliServerClientPopulateRequestInfo(client *PhpCliServerClient, request_info *core.SapiRequestInfo) {
	var val *byte
	request_info.request_method = PhpHttpMethodStr(client.GetRequest().GetRequestMethod())
	request_info.proto_num = client.GetRequest().GetProtocolVersion()
	request_info.request_uri = client.GetRequest().GetRequestUri()
	request_info.path_translated = client.GetRequest().GetPathTranslated()
	request_info.query_string = client.GetRequest().GetQueryString()
	request_info.content_length = client.GetRequest().GetContentLen()
	request_info.auth_digest = nil
	request_info.auth_password = request_info.auth_digest
	request_info.auth_user = request_info.auth_password
	if nil != g.Assign(&val, zend.ZendHashStrFindPtr(&client.request.GetHeaders(), "content-type", g.SizeOf("\"content-type\"")-1)) {
		request_info.content_type = val
	}
}
func DestroyRequestInfo(request_info *core.SapiRequestInfo) {}
func PhpCliServerClientCtor(client *PhpCliServerClient, server *PhpCliServer, client_sock core.PhpSocketT, addr *__struct__sockaddr, addr_len socklen_t) int {
	client.SetServer(server)
	client.SetSock(client_sock)
	client.SetAddr(addr)
	client.SetAddrLen(addr_len)
	var addr_str *zend.ZendString = 0
	core.PhpNetworkPopulateNameFromSockaddr(addr, addr_len, &addr_str, nil, 0)
	client.SetAddrStr(zend.ZendStrndup(addr_str.val, addr_str.len_))
	client.SetAddrStrLen(addr_str.len_)
	zend.ZendStringReleaseEx(addr_str, 0)
	PhpHttpParserInit(&client.parser, PHP_HTTP_REQUEST)
	client.SetRequestRead(0)
	client.SetLastHeaderElement(HEADER_NONE)
	client.SetCurrentHeaderName(nil)
	client.SetCurrentHeaderNameLen(0)
	client.SetCurrentHeaderNameAllocated(0)
	client.SetCurrentHeaderValue(nil)
	client.SetCurrentHeaderValueLen(0)
	client.SetPostReadOffset(0)
	if zend.FAILURE == PhpCliServerRequestCtor(&client.request) {
		return zend.FAILURE
	}
	client.SetContentSenderInitialized(0)
	client.SetFileFd(-1)
	return zend.SUCCESS
}
func PhpCliServerClientDtor(client *PhpCliServerClient) {
	PhpCliServerRequestDtor(&client.request)
	if client.GetFileFd() >= 0 {
		close(client.GetFileFd())
		client.SetFileFd(-1)
	}
	g.CondF(true, func() { return zend.Free(client.GetAddr()) }, func() { return zend._efree(client.GetAddr()) })
	g.CondF(true, func() { return zend.Free(client.GetAddrStr()) }, func() { return zend._efree(client.GetAddrStr()) })
	if client.GetContentSenderInitialized() != 0 {
		PhpCliServerContentSenderDtor(&client.content_sender)
	}
}
func PhpCliServerCloseConnection(server *PhpCliServer, client *PhpCliServerClient) {
	PhpCliServerLogf(3, "%s Closing", client.GetAddrStr())
	zend.ZendHashIndexDel(&server.clients, client.GetSock())
}
func PhpCliServerSendErrorPage(server *PhpCliServer, client *PhpCliServerClient, status int) int {
	var escaped_request_uri *zend.ZendString = nil
	var status_string *byte = GetStatusString(status)
	var content_template *byte = GetTemplateString(status)
	var errstr *byte = GetLastError()
	assert(status_string != nil && content_template != nil)
	PhpCliServerContentSenderCtor(&client.content_sender)
	client.SetContentSenderInitialized(1)
	escaped_request_uri = standard.PhpEscapeHtmlEntitiesEx((*uint8)(client.GetRequest().GetRequestUri()), client.GetRequest().GetRequestUriLen(), 0, 2|1, nil, 0)
	var prologue_template []byte = "<!doctype html><html><head><title>%d %s</title>"
	var chunk *php_cli_server_chunk = php_cli_server_chunk_heap_new_self_contained(strlen(prologue_template) + 3 + strlen(status_string) + 1)
	if chunk == nil {
		goto fail
	}
	core.ApPhpSnprintf(chunk.GetDataHeapP(), chunk.GetDataHeapLen(), prologue_template, status, status_string)
	chunk.SetDataHeapLen(strlen(chunk.GetDataHeapP()))
	PhpCliServerBufferAppend(&client.content_sender.GetBuffer(), chunk)
	var chunk *php_cli_server_chunk = php_cli_server_chunk_immortal_new(php_cli_server_css, g.SizeOf("php_cli_server_css")-1)
	if chunk == nil {
		goto fail
	}
	PhpCliServerBufferAppend(&client.content_sender.GetBuffer(), chunk)
	var template []byte = "</head><body>"
	var chunk *php_cli_server_chunk = php_cli_server_chunk_immortal_new(template, g.SizeOf("template")-1)
	if chunk == nil {
		goto fail
	}
	PhpCliServerBufferAppend(&client.content_sender.GetBuffer(), chunk)
	var chunk *php_cli_server_chunk = php_cli_server_chunk_heap_new_self_contained(strlen(content_template) + escaped_request_uri.len + 3 + strlen(status_string) + 1)
	if chunk == nil {
		goto fail
	}
	core.ApPhpSnprintf(chunk.GetDataHeapP(), chunk.GetDataHeapLen(), content_template, status_string, escaped_request_uri.val)
	chunk.SetDataHeapLen(strlen(chunk.GetDataHeapP()))
	PhpCliServerBufferAppend(&client.content_sender.GetBuffer(), chunk)
	var epilogue_template []byte = "</body></html>"
	var chunk *php_cli_server_chunk = php_cli_server_chunk_immortal_new(epilogue_template, g.SizeOf("epilogue_template")-1)
	if chunk == nil {
		goto fail
	}
	PhpCliServerBufferAppend(&client.content_sender.GetBuffer(), chunk)
	var chunk *PhpCliServerChunk
	var buffer zend.SmartStr = zend.SmartStr{0}
	AppendHttpStatusLine(&buffer, client.GetRequest().GetProtocolVersion(), status, 1)
	if buffer.s == nil {

		/* out of memory */

		goto fail

		/* out of memory */

	}
	AppendEssentialHeaders(&buffer, client, 1)
	zend.SmartStrAppendlEx(&buffer, "Content-Type: text/html; charset=UTF-8\r\n", strlen("Content-Type: text/html; charset=UTF-8\r\n"), 1)
	zend.SmartStrAppendlEx(&buffer, "Content-Length: ", strlen("Content-Length: "), 1)
	zend.SmartStrAppendUnsignedEx(&buffer, PhpCliServerBufferSize(&client.content_sender.GetBuffer()), 1)
	zend.SmartStrAppendlEx(&buffer, "\r\n", 2, 1)
	zend.SmartStrAppendlEx(&buffer, "\r\n", 2, 1)
	chunk = PhpCliServerChunkHeapNew(buffer.s, buffer.s.val, buffer.s.len_)
	if chunk == nil {
		zend.SmartStrFreeEx(&buffer, 1)
		goto fail
	}
	PhpCliServerBufferPrepend(&client.content_sender.GetBuffer(), chunk)
	PhpCliServerLogResponse(client, status, g.Cond(errstr != nil, errstr, "?"))
	PhpCliServerPollerAdd(&server.poller, POLLOUT, client.GetSock())
	if errstr != nil {
		g.CondF(true, func() { return zend.Free(errstr) }, func() { return zend._efree(errstr) })
	}
	zend.ZendStringFree(escaped_request_uri)
	return zend.SUCCESS
fail:
	if errstr != nil {
		g.CondF(true, func() { return zend.Free(errstr) }, func() { return zend._efree(errstr) })
	}
	zend.ZendStringFree(escaped_request_uri)
	return zend.FAILURE
}
func PhpCliServerDispatchScript(server *PhpCliServer, client *PhpCliServerClient) int {
	if strlen(client.GetRequest().GetPathTranslated()) != client.GetRequest().GetPathTranslatedLen() {

		/* can't handle paths that contain nul bytes */

		return PhpCliServerSendErrorPage(server, client, 400)

		/* can't handle paths that contain nul bytes */

	}
	var zfd zend.ZendFileHandle
	zend.ZendStreamInitFilename(&zfd, core.sapi_globals.request_info.path_translated)
	var __orig_bailout *sigjmp_buf = zend.EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		core.PhpExecuteScript(&zfd)
	}
	zend.EG.bailout = __orig_bailout
	PhpCliServerLogResponse(client, core.sapi_globals.sapi_headers.http_response_code, nil)
	return zend.SUCCESS
}
func PhpCliServerBeginSendStatic(server *PhpCliServer, client *PhpCliServerClient) int {
	var fd int
	var status int = 200
	if client.GetRequest().GetPathTranslated() != nil && strlen(client.GetRequest().GetPathTranslated()) != client.GetRequest().GetPathTranslatedLen() {

		/* can't handle paths that contain nul bytes */

		return PhpCliServerSendErrorPage(server, client, 400)

		/* can't handle paths that contain nul bytes */

	}
	if client.GetRequest().GetPathTranslated() != nil {
		fd = open(client.GetRequest().GetPathTranslated(), O_RDONLY)
	} else {
		fd = -1
	}
	if fd < 0 {
		return PhpCliServerSendErrorPage(server, client, 404)
	}
	PhpCliServerContentSenderCtor(&client.content_sender)
	client.SetContentSenderInitialized(1)
	client.SetFileFd(fd)
	var chunk *PhpCliServerChunk
	var buffer zend.SmartStr = zend.SmartStr{0}
	var mime_type *byte = GetMimeType(server, client.GetRequest().GetExt(), client.GetRequest().GetExtLen())
	AppendHttpStatusLine(&buffer, client.GetRequest().GetProtocolVersion(), status, 1)
	if buffer.s == nil {

		/* out of memory */

		PhpCliServerLogResponse(client, 500, nil)
		return zend.FAILURE
	}
	AppendEssentialHeaders(&buffer, client, 1)
	if mime_type != nil {
		zend.SmartStrAppendlEx(&buffer, "Content-Type: ", g.SizeOf("\"Content-Type: \"")-1, 1)
		zend.SmartStrAppendlEx(&buffer, mime_type, strlen(mime_type), 1)
		if strncmp(mime_type, "text/", 5) == 0 {
			zend.SmartStrAppendlEx(&buffer, "; charset=UTF-8", strlen("; charset=UTF-8"), 1)
		}
		zend.SmartStrAppendlEx(&buffer, "\r\n", 2, 1)
	}
	zend.SmartStrAppendlEx(&buffer, "Content-Length: ", strlen("Content-Length: "), 1)
	zend.SmartStrAppendUnsignedEx(&buffer, client.request.sb.st_size, 1)
	zend.SmartStrAppendlEx(&buffer, "\r\n", 2, 1)
	zend.SmartStrAppendlEx(&buffer, "\r\n", 2, 1)
	chunk = PhpCliServerChunkHeapNew(buffer.s, buffer.s.val, buffer.s.len_)
	if chunk == nil {
		zend.SmartStrFreeEx(&buffer, 1)
		PhpCliServerLogResponse(client, 500, nil)
		return zend.FAILURE
	}
	PhpCliServerBufferAppend(&client.content_sender.GetBuffer(), chunk)
	PhpCliServerLogResponse(client, 200, nil)
	PhpCliServerPollerAdd(&server.poller, POLLOUT, client.GetSock())
	return zend.SUCCESS
}

/* }}} */

func PhpCliServerRequestStartup(server *PhpCliServer, client *PhpCliServerClient) int {
	var auth *byte
	PhpCliServerClientPopulateRequestInfo(client, &(core.sapi_globals.request_info))
	if nil != g.Assign(&auth, zend.ZendHashStrFindPtr(&client.request.GetHeaders(), "authorization", g.SizeOf("\"authorization\"")-1)) {
		core.PhpHandleAuthData(auth)
	}
	core.sapi_globals.sapi_headers.http_response_code = 200
	if zend.FAILURE == core.PhpRequestStartup() {

		/* should never be happen */

		DestroyRequestInfo(&(core.sapi_globals.request_info))
		return zend.FAILURE
	}
	core.CoreGlobals.during_request_startup = 0
	return zend.SUCCESS
}

/* }}} */

func PhpCliServerRequestShutdown(server *PhpCliServer, client *PhpCliServerClient) int {
	core.PhpRequestShutdown(0)
	PhpCliServerCloseConnection(server, client)
	DestroyRequestInfo(&(core.sapi_globals.request_info))
	core.sapi_globals.server_context = nil
	core.sapi_globals.rfc1867_uploaded_files = nil
	return zend.SUCCESS
}

/* }}} */

func PhpCliServerDispatchRouter(server *PhpCliServer, client *PhpCliServerClient) int {
	var decline int = 0
	var zfd zend.ZendFileHandle
	var old_cwd *byte
	old_cwd = zend._emalloc(256)
	old_cwd[0] = '0'
	void(getcwd(old_cwd, 256-1))
	zend.ZendStreamInitFilename(&zfd, server.GetRouter())
	var __orig_bailout *sigjmp_buf = zend.EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		var retval zend.Zval
		&retval.u1.type_info = 0
		if zend.SUCCESS == zend.ZendExecuteScripts(1<<3, &retval, 1, &zfd) {
			if retval.u1.v.type_ != 0 {
				decline = retval.u1.v.type_ == 2
				zend.ZvalPtrDtor(&retval)
			}
		} else {
			decline = 1
		}
	}
	zend.EG.bailout = __orig_bailout
	if old_cwd[0] != '0' {
		void(chdir(old_cwd))
	}
	zend._efree(old_cwd)
	return decline
}

/* }}} */

func PhpCliServerDispatch(server *PhpCliServer, client *PhpCliServerClient) int {
	var is_static_file int = 0
	var ext *byte = client.GetRequest().GetExt()
	core.sapi_globals.server_context = client
	if client.GetRequest().GetExtLen() != 3 || ext[0] != 'p' && ext[0] != 'P' || ext[1] != 'h' && ext[1] != 'H' || ext[2] != 'p' && ext[2] != 'P' || client.GetRequest().GetPathTranslated() == nil {
		is_static_file = 1
	}
	if server.GetRouter() != nil || is_static_file == 0 {
		if zend.FAILURE == PhpCliServerRequestStartup(server, client) {
			core.sapi_globals.server_context = nil
			PhpCliServerCloseConnection(server, client)
			DestroyRequestInfo(&(core.sapi_globals.request_info))
			return zend.SUCCESS
		}
	}
	if server.GetRouter() != nil {
		if PhpCliServerDispatchRouter(server, client) == 0 {
			PhpCliServerRequestShutdown(server, client)
			return zend.SUCCESS
		}
	}
	if is_static_file == 0 {
		if zend.SUCCESS == PhpCliServerDispatchScript(server, client) || zend.SUCCESS != PhpCliServerSendErrorPage(server, client, 500) {
			if core.sapi_globals.sapi_headers.http_response_code == 304 {
				core.sapi_globals.sapi_headers.send_default_content_type = 0
			}
			PhpCliServerRequestShutdown(server, client)
			return zend.SUCCESS
		}
	} else {
		if server.GetRouter() != nil {
			var send_header_func func(*core.SapiHeaders) int
			send_header_func = core.sapi_module.send_headers

			/* do not generate default content type header */

			core.sapi_globals.sapi_headers.send_default_content_type = 0

			/* we don't want headers to be sent */

			core.sapi_module.send_headers = SapiCliServerDiscardHeaders
			core.PhpRequestShutdown(0)
			core.sapi_module.send_headers = send_header_func
			core.sapi_globals.sapi_headers.send_default_content_type = 1
			core.sapi_globals.rfc1867_uploaded_files = nil
		}
		if zend.SUCCESS != PhpCliServerBeginSendStatic(server, client) {
			PhpCliServerCloseConnection(server, client)
		}
		core.sapi_globals.server_context = nil
		return zend.SUCCESS
	}
	core.sapi_globals.server_context = nil
	DestroyRequestInfo(&(core.sapi_globals.request_info))
	return zend.SUCCESS
}

/* }}} */

func PhpCliServerMimeTypeCtor(server *PhpCliServer, mime_type_map *PhpCliServerExtMimeTypePair) int {
	var pair *PhpCliServerExtMimeTypePair
	zend._zendHashInit(&server.extension_mime_types, 0, nil, 1)
	for pair = mime_type_map; pair.GetExt() != nil; pair++ {
		var ext_len int = strlen(pair.GetExt())
		zend.ZendHashStrAddPtr(&server.extension_mime_types, pair.GetExt(), ext_len, any(pair.GetMimeType()))
	}
	return zend.SUCCESS
}
func PhpCliServerDtor(server *PhpCliServer) {
	zend.ZendHashDestroy(&server.clients)
	zend.ZendHashDestroy(&server.extension_mime_types)
	if server.GetServerSock() >= 0 {
		close(server.GetServerSock())
	}
	if server.GetHost() != nil {
		g.CondF(true, func() { return zend.Free(server.GetHost()) }, func() { return zend._efree(server.GetHost()) })
	}
	if server.GetDocumentRoot() != nil {
		g.CondF(true, func() { return zend.Free(server.GetDocumentRoot()) }, func() { return zend._efree(server.GetDocumentRoot()) })
	}
	if server.GetRouter() != nil {
		g.CondF(true, func() { return zend.Free(server.GetRouter()) }, func() { return zend._efree(server.GetRouter()) })
	}
	if PhpCliServerWorkersMax > 1 && PhpCliServerWorkers != nil && getpid() == PhpCliServerMaster {
		var php_cli_server_worker zend.ZendLong
		for php_cli_server_worker = 0; php_cli_server_worker < PhpCliServerWorkersMax; php_cli_server_worker++ {
			var php_cli_server_worker_status int
			for {
				if waitpid(PhpCliServerWorkers[php_cli_server_worker], &php_cli_server_worker_status, 0) == zend.FAILURE {

					/* an extremely bad thing happened */

					break

					/* an extremely bad thing happened */

				}
				if !(!(WIFEXITED(php_cli_server_worker_status)) && !(WIFSIGNALED(php_cli_server_worker_status))) {
					break
				}
			}
		}
		zend.Free(PhpCliServerWorkers)
	}
}
func PhpCliServerClientDtorWrapper(zv *zend.Zval) {
	var p *PhpCliServerClient = zv.value.ptr
	shutdown(p.GetSock(), 2)
	close(p.GetSock())
	PhpCliServerPollerRemove(&p.server.GetPoller(), POLLIN|POLLOUT, p.GetSock())
	PhpCliServerClientDtor(p)
	g.CondF(true, func() { return zend.Free(p) }, func() { return zend._efree(p) })
}
func PhpCliServerCtor(server *PhpCliServer, addr *byte, document_root *byte, router *byte) int {
	var retval int = zend.SUCCESS
	var host *byte = nil
	var errstr *zend.ZendString = nil
	var _document_root *byte = nil
	var _router *byte = nil
	var err int = 0
	var port int = 3000
	var server_sock core.PhpSocketT = -1
	var p *byte = nil
	if addr[0] == '[' {
		host = strdup(addr + 1)
		if host == nil {
			return zend.FAILURE
		}
		p = strchr(host, ']')
		if p != nil {
			g.PostInc(&(*p)) = '0'
			if (*p) == ':' {
				port = strtol(p+1, &p, 10)
				if port <= 0 || port > 65535 {
					p = nil
				}
			} else if (*p) != '0' {
				p = nil
			}
		}
	} else {
		host = strdup(addr)
		if host == nil {
			return zend.FAILURE
		}
		p = strchr(host, ':')
		if p != nil {
			g.PostInc(&(*p)) = '0'
			port = strtol(p, &p, 10)
			if port <= 0 || port > 65535 {
				p = nil
			}
		}
	}
	if p == nil {
		fprintf(stderr, "Invalid address: %s\n", addr)
		retval = zend.FAILURE
		goto out
	}
	server_sock = PhpNetworkListenSocket(host, &port, SOCK_STREAM, &server.address_family, &server.socklen, &errstr)
	if server_sock == -1 {
		PhpCliServerLogf(2, "Failed to listen on %s:%d (reason: %s)", host, port, g.CondF1(errstr != nil, func() []byte { return errstr.val }, "?"))
		if errstr != nil {
			zend.ZendStringReleaseEx(errstr, 0)
		}
		retval = zend.FAILURE
		goto out
	}
	server.SetServerSock(server_sock)
	err = PhpCliServerPollerCtor(&server.poller)
	if zend.SUCCESS != err {
		goto out
	}
	PhpCliServerPollerAdd(&server.poller, POLLIN, server_sock)
	server.SetHost(host)
	server.SetPort(port)
	zend._zendHashInit(&server.clients, 0, PhpCliServerClientDtorWrapper, 1)
	var document_root_len int = strlen(document_root)
	_document_root = zend.ZendStrndup(document_root, document_root_len)
	if _document_root == nil {
		retval = zend.FAILURE
		goto out
	}
	server.SetDocumentRoot(_document_root)
	server.SetDocumentRootLen(document_root_len)
	if router != nil {
		var router_len int = strlen(router)
		_router = zend.ZendStrndup(router, router_len)
		if _router == nil {
			retval = zend.FAILURE
			goto out
		}
		server.SetRouter(_router)
		server.SetRouterLen(router_len)
	} else {
		server.SetRouter(nil)
		server.SetRouterLen(0)
	}
	if PhpCliServerMimeTypeCtor(server, MimeTypeMap) == zend.FAILURE {
		retval = zend.FAILURE
		goto out
	}
	server.SetIsRunning(1)
out:
	if retval != zend.SUCCESS {
		if host != nil {
			g.CondF(true, func() { return zend.Free(host) }, func() { return zend._efree(host) })
		}
		if _document_root != nil {
			g.CondF(true, func() { return zend.Free(_document_root) }, func() { return zend._efree(_document_root) })
		}
		if _router != nil {
			g.CondF(true, func() { return zend.Free(_router) }, func() { return zend._efree(_router) })
		}
		if server_sock > -1 {
			close(server_sock)
		}
	}
	return retval
}
func PhpCliServerRecvEventReadRequest(server *PhpCliServer, client *PhpCliServerClient) int {
	var errstr *byte = nil
	var status int = PhpCliServerClientReadRequest(client, &errstr)
	if status < 0 {
		if errstr != nil {
			if strcmp(errstr, PhpCliServerRequestErrorUnexpectedEof) == 0 && client.GetParser().GetState() == SStartReq {
				PhpCliServerLogf(3, "%s Closed without sending a request; it was probably just an unused speculative preconnection", client.GetAddrStr())
			} else {
				PhpCliServerLogf(2, "%s Invalid request (%s)", client.GetAddrStr(), errstr)
			}
			zend._efree(errstr)
		}
		PhpCliServerCloseConnection(server, client)
		return zend.FAILURE
	} else if status == 1 && client.GetRequest().GetRequestMethod() == PHP_HTTP_NOT_IMPLEMENTED {
		return PhpCliServerSendErrorPage(server, client, 501)
	} else if status == 1 {
		PhpCliServerPollerRemove(&server.poller, POLLIN, client.GetSock())
		PhpCliServerDispatch(server, client)
	} else {
		PhpCliServerPollerAdd(&server.poller, POLLIN, client.GetSock())
	}
	return zend.SUCCESS
}
func PhpCliServerSendEvent(server *PhpCliServer, client *PhpCliServerClient) int {
	if client.GetContentSenderInitialized() != 0 {
		if client.GetFileFd() >= 0 && client.GetContentSender().GetBuffer().GetFirst() == nil {
			var nbytes_read int
			if PhpCliServerContentSenderPull(&client.content_sender, client.GetFileFd(), &nbytes_read) != 0 {
				PhpCliServerCloseConnection(server, client)
				return zend.FAILURE
			}
			if nbytes_read == 0 {
				close(client.GetFileFd())
				client.SetFileFd(-1)
			}
		}
		var nbytes_sent int
		var err int = PhpCliServerContentSenderSend(&client.content_sender, client.GetSock(), &nbytes_sent)
		if err != 0 && err != EAGAIN {
			PhpCliServerCloseConnection(server, client)
			return zend.FAILURE
		}
		if client.GetContentSender().GetBuffer().GetFirst() == nil && client.GetFileFd() < 0 {
			PhpCliServerCloseConnection(server, client)
		}
	}
	return zend.SUCCESS
}

/* }}} */

// @type PhpCliServerDoEventForEachFdCallbackParams struct

func PhpCliServerDoEventForEachFdCallback(_params any, fd core.PhpSocketT, event int) int {
	var params *PhpCliServerDoEventForEachFdCallbackParams = _params
	var server *PhpCliServer = params.GetServer()
	if server.GetServerSock() == fd {
		var client *PhpCliServerClient = nil
		var client_sock core.PhpSocketT
		var socklen socklen_t = server.GetSocklen()
		var sa *__struct__sockaddr = g.CondF(true, func() any { return zend.__zendMalloc(server.GetSocklen()) }, func() any { return zend._emalloc(server.GetSocklen()) })
		client_sock = accept(server.GetServerSock(), sa, &socklen)
		if client_sock < 0 {
			if PhpCliServerLogLevel >= 2 {
				var errstr *byte = core.PhpSocketStrerror(errno, nil, 0)
				PhpCliServerLogf(2, "Failed to accept a client (reason: %s)", errstr)
				zend._efree(errstr)
			}
			g.CondF(true, func() { return zend.Free(sa) }, func() { return zend._efree(sa) })
			return zend.SUCCESS
		}
		if zend.SUCCESS != core.PhpSetSockBlocking(client_sock, 0) {
			g.CondF(true, func() { return zend.Free(sa) }, func() { return zend._efree(sa) })
			close(client_sock)
			return zend.SUCCESS
		}
		client = zend.__zendMalloc(g.SizeOf("php_cli_server_client"))
		if zend.FAILURE == PhpCliServerClientCtor(client, server, client_sock, sa, socklen) {
			PhpCliServerLogf(2, "Failed to create a new request object")
			g.CondF(true, func() { return zend.Free(sa) }, func() { return zend._efree(sa) })
			close(client_sock)
			return zend.SUCCESS
		}
		PhpCliServerLogf(3, "%s Accepted", client.GetAddrStr())
		zend.ZendHashIndexUpdatePtr(&server.clients, client_sock, client)
		PhpCliServerPollerAdd(&server.poller, POLLIN, client.GetSock())
	} else {
		var client *PhpCliServerClient
		if nil != g.Assign(&client, zend.ZendHashIndexFindPtr(&server.clients, fd)) {
			if (event & POLLIN) != 0 {
				params.GetRhandler()(server, client)
			}
			if (event & POLLOUT) != 0 {
				params.GetWhandler()(server, client)
			}
		}
	}
	return zend.SUCCESS
}
func PhpCliServerDoEventForEachFd(server *PhpCliServer, rhandler func(*PhpCliServer, *PhpCliServerClient) int, whandler func(*PhpCliServer, *PhpCliServerClient) int) {
	var params PhpCliServerDoEventForEachFdCallbackParams = PhpCliServerDoEventForEachFdCallbackParams{server, rhandler, whandler}
	PhpCliServerPollerIterOnActive(&server.poller, &params, PhpCliServerDoEventForEachFdCallback)
}
func PhpCliServerDoEventLoop(server *PhpCliServer) int {
	var retval int = zend.SUCCESS
	for server.GetIsRunning() != 0 {
		var tv __struct__timeval = __struct__timeval{1, 0}
		var n int = PhpCliServerPollerPoll(&server.poller, &tv)
		if n > 0 {
			PhpCliServerDoEventForEachFd(server, PhpCliServerRecvEventReadRequest, PhpCliServerSendEvent)
		} else if n == 0 {

		} else {
			var err int = errno
			if err != EINTR {
				if PhpCliServerLogLevel >= 2 {
					var errstr *byte = core.PhpSocketStrerror(err, nil, 0)
					PhpCliServerLogf(2, "%s", errstr)
					zend._efree(errstr)
				}
				retval = zend.FAILURE
				goto out
			}
		}
	}
out:
	return retval
}

var Server PhpCliServer

func PhpCliServerSigintHandler(sig int) { Server.SetIsRunning(0) }

/* }}} */

func DoCliServer(argc int, argv **byte) int {
	var php_optarg *byte = nil
	var php_optind int = 1
	var c int
	var server_bind_address *byte = nil
	var OPTIONS []core.Opt
	var document_root *byte = nil
	var router *byte = nil
	var document_root_buf []byte
	for g.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &php_optarg, &php_optind, 0, 2)) != -1 {
		switch c {
		case 'S':
			server_bind_address = php_optarg
			break
		case 't':
			document_root = php_optarg
			break
		case 'q':
			if PhpCliServerLogLevel > 1 {
				PhpCliServerLogLevel--
			}
			break
		}
	}
	if document_root != nil {
		var sb zend.ZendStatT
		if stat(document_root, &sb) {
			fprintf(stderr, "Directory %s does not exist.\n", document_root)
			return 1
		}
		if (sb.st_mode & S_IFMT) != S_IFDIR {
			fprintf(stderr, "%s is not a directory.\n", document_root)
			return 1
		}
		if zend.TsrmRealpath(document_root, document_root_buf) != nil {
			document_root = document_root_buf
		}
	} else {
		var ret *byte = nil
		ret = getcwd(document_root_buf, 256)
		if ret != nil {
			document_root = document_root_buf
		} else {
			document_root = "."
		}
	}
	if argc > php_optind {
		router = argv[php_optind]
	}
	if zend.FAILURE == PhpCliServerCtor(&Server, server_bind_address, document_root, router) {
		return 1
	}
	core.sapi_module.phpinfo_as_text = 0
	PhpCliServerLogf(1, "PHP %s Development Server (http://%s) started", "7.4.33", server_bind_address)
	zend.ZendSignalInit()
	PhpCliServerDoEventLoop(&Server)
	PhpCliServerDtor(&Server)
	return 0
}
