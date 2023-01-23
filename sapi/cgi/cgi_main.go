// <<generate>>

package cgi

import (
	"sik/core"
	"sik/core/streams"
	"sik/ext/standard"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/sapi/cli"
	"sik/zend"
)

// Source: <sapi/cgi/cgi_main.c>

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
   | Authors: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                       |
   |          Stig Bakken <ssb@php.net>                                   |
   |          Zeev Suraski <zeev@php.net>                                 |
   | FastCGI: Ben Mansell <php@slimyhorror.com>                           |
   |          Shane Caraveo <shane@caraveo.com>                           |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_globals.h"

// # include "php_variables.h"

// # include "zend_modules.h"

// # include "SAPI.h"

// # include < stdio . h >

// # include < sys / time . h >

// # include < unistd . h >

// # include < signal . h >

// # include < locale . h >

// # include < sys / types . h >

// # include < sys / wait . h >

// # include "zend.h"

// # include "zend_extensions.h"

// # include "php_ini.h"

// # include "php_globals.h"

// # include "php_main.h"

// # include "fopen_wrappers.h"

// # include "http_status_codes.h"

// # include "ext/standard/php_standard.h"

// # include "ext/standard/url.h"

// # include "zend_compile.h"

// # include "zend_execute.h"

// # include "zend_highlight.h"

// # include "php_getopt.h"

// # include "fastcgi.h"

/* XXX this will need to change later when threaded fastcgi is implemented.  shane */

var Act __struct__sigaction
var OldTerm __struct__sigaction
var OldQuit __struct__sigaction
var OldInt __struct__sigaction
var PhpPhpImportEnvironmentVariables func(array_ptr *zend.Zval)

/* these globals used for forking children on unix systems */

var Children int = 0

/**
 * Set to non-zero if we are the parent process
 */

var Parent int = 1

/* Did parent received exit signals SIG_TERM/SIG_INT/SIG_QUIT */

var ExitSignal int = 0

/* Is Parent waiting for children to exit */

var ParentWaiting int = 0

/**
 * Process group
 */

var Pgroup pid_t

// #define PHP_MODE_STANDARD       1

// #define PHP_MODE_HIGHLIGHT       2

// #define PHP_MODE_LINT       4

// #define PHP_MODE_STRIP       5

var PhpOptarg *byte = nil
var PhpOptind int = 1
var OPTIONS []core.Opt = []core.Opt{{'a', 0, "interactive"}, {'b', 1, "bindpath"}, {'C', 0, "no-chdir"}, {'c', 1, "php-ini"}, {'d', 1, "define"}, {'e', 0, "profile-info"}, {'f', 1, "file"}, {'h', 0, "help"}, {'i', 0, "info"}, {'l', 0, "syntax-check"}, {'m', 0, "modules"}, {'n', 0, "no-php-ini"}, {'q', 0, "no-header"}, {'s', 0, "syntax-highlight"}, {'s', 0, "syntax-highlighting"}, {'w', 0, "strip"}, {'?', 0, "usage"}, {'v', 0, "version"}, {'z', 1, "zend-extension"}, {'T', 1, "timing"}, {'-', 0, nil}}

type _phpCgiGlobals = php_cgi_globals_struct

/* {{{ user_config_cache
 *
 * Key for each cache entry is dirname(PATH_TRANSLATED).
 *
 * NOTE: Each cache entry config_hash contains the combination from all user ini files found in
 *       the path starting from doc_root through to dirname(PATH_TRANSLATED).  There is no point
 *       storing per-file entries as it would not be possible to detect added / deleted entries
 *       between separate files.
 */

func UserConfigCacheEntryDtor(el *zend.Zval) {
	var entry *UserConfigCacheEntry = (*UserConfigCacheEntry)(el.value.ptr)
	zend.ZendHashDestroy(entry.GetUserConfig())
	zend.Free(entry.GetUserConfig())
	zend.Free(entry)
}

/* }}} */

var php_cgi_globals php_cgi_globals_struct

// #define CGIG(v) ( php_cgi_globals . v )

// #define TRANSLATE_SLASHES(path)

func FcgiLog(type_ int, format *byte, _ ...any) {
	var ap va_list
	va_start(ap, format)
	vfprintf(stderr, format, ap)
	va_end(ap)
}
func ModuleNameCmp(a any, b any) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	return strcasecmp((*zend.ZendModuleEntry)(f.val.value.ptr).name, (*zend.ZendModuleEntry)(s.val.value.ptr).name)
}
func PrintModules() {
	var sorted_registry zend.HashTable
	var module *zend.ZendModuleEntry
	zend._zendHashInit(&sorted_registry, 64, nil, 1)
	zend.ZendHashCopy(&sorted_registry, &zend.ModuleRegistry, nil)
	zend.ZendHashSortEx(&sorted_registry, zend.ZendSort, ModuleNameCmp, 0)
	for {
		var __ht *zend.HashTable = &sorted_registry
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			module = _z.value.ptr
			core.PhpPrintf("%s\n", module.name)
		}
		break
	}
	zend.ZendHashDestroy(&sorted_registry)
}
func PrintExtensionInfo(ext *zend.ZendExtension, arg any) int {
	core.PhpPrintf("%s\n", ext.name)
	return 0
}
func ExtensionNameCmp(f **zend.ZendLlistElement, s **zend.ZendLlistElement) int {
	var fe *zend.ZendExtension = (*zend.ZendExtension)((*f).data)
	var se *zend.ZendExtension = (*zend.ZendExtension)((*s).data)
	return strcmp(fe.name, se.name)
}
func PrintExtensions() {
	var sorted_exts zend.ZendLlist
	zend.ZendLlistCopy(&sorted_exts, &zend.ZendExtensions)
	sorted_exts.dtor = nil
	zend.ZendLlistSort(&sorted_exts, ExtensionNameCmp)
	zend.ZendLlistApplyWithArgument(&sorted_exts, zend.LlistApplyWithArgFuncT(PrintExtensionInfo), nil)
	zend.ZendLlistDestroy(&sorted_exts)
}

// #define STDOUT_FILENO       1

func SapiCgiSingleWrite(str *byte, str_length int) int {
	var ret int
	ret = write(1, str, str_length)
	if ret <= 0 {
		return 0
	}
	return ret
}
func SapiCgiUbWrite(str *byte, str_length int) int {
	var ptr *byte = str
	var remaining int = str_length
	var ret int
	for remaining > 0 {
		ret = SapiCgiSingleWrite(ptr, remaining)
		if ret == 0 {
			core.PhpHandleAbortedConnection()
			return str_length - remaining
		}
		ptr += ret
		remaining -= ret
	}
	return str_length
}
func SapiFcgiUbWrite(str *byte, str_length int) int {
	var ptr *byte = str
	var remaining int = str_length
	var request *core.FcgiRequest = (*core.FcgiRequest)(core.sapi_globals.server_context)
	for remaining > 0 {
		var to_write int = g.CondF2(remaining > 2147483647, 2147483647, func() int { return int(remaining) })
		var ret int = core.FcgiWrite(request, core.FCGI_STDOUT, ptr, to_write)
		if ret <= 0 {
			core.PhpHandleAbortedConnection()
			return str_length - remaining
		}
		ptr += ret
		remaining -= ret
	}
	return str_length
}
func SapiCgiFlush(server_context any) {
	if r.Fflush(stdout) == -1 {
		core.PhpHandleAbortedConnection()
	}
}
func SapiFcgiFlush(server_context any) {
	var request *core.FcgiRequest = (*core.FcgiRequest)(server_context)
	if Parent == 0 && request != nil && core.FcgiFlush(request, 0) == 0 {
		core.PhpHandleAbortedConnection()
	}
}

// #define SAPI_CGI_MAX_HEADER_LENGTH       1024

func SapiCgiSendHeaders(sapi_headers *core.SapiHeaders) int {
	var h *core.SapiHeader
	var pos zend.ZendLlistPosition
	var ignore_status zend.ZendBool = 0
	var response_status int = core.sapi_globals.sapi_headers.http_response_code
	if core.sapi_globals.request_info.no_headers == 1 {
		return 1
	}
	if php_cgi_globals.GetNph() != 0 || core.sapi_globals.sapi_headers.http_response_code != 200 {
		var len_ int
		var has_status zend.ZendBool = 0
		var buf []byte
		if php_cgi_globals.GetRfc2616Headers() != 0 && core.sapi_globals.sapi_headers.http_status_line != nil {
			var s *byte
			len_ = core.ApPhpSlprintf(buf, 1024, "%s", core.sapi_globals.sapi_headers.http_status_line)
			if g.Assign(&s, strchr(core.sapi_globals.sapi_headers.http_status_line, ' ')) {
				response_status = atoi(s + 1)
			}
			if len_ > 1024 {
				len_ = 1024
			}
		} else {
			var s *byte
			if core.sapi_globals.sapi_headers.http_status_line != nil && g.Assign(&s, strchr(core.sapi_globals.sapi_headers.http_status_line, ' ')) != 0 && s-core.sapi_globals.sapi_headers.http_status_line >= 5 && strncasecmp(core.sapi_globals.sapi_headers.http_status_line, "HTTP/", 5) == 0 {
				len_ = core.ApPhpSlprintf(buf, g.SizeOf("buf"), "Status:%s", s)
				response_status = atoi(s + 1)
			} else {
				h = (*core.SapiHeader)(zend.ZendLlistGetFirstEx(&sapi_headers.headers, &pos))
				for h != nil {
					if h.header_len > g.SizeOf("\"Status:\"")-1 && strncasecmp(h.header, "Status:", g.SizeOf("\"Status:\"")-1) == 0 {
						has_status = 1
						break
					}
					h = (*core.SapiHeader)(zend.ZendLlistGetNextEx(&sapi_headers.headers, &pos))
				}
				if has_status == 0 {
					var err *core.HttpResponseStatusCodePair = (*core.HttpResponseStatusCodePair)(core.HttpStatusMap)
					for err.code != 0 {
						if err.code == core.sapi_globals.sapi_headers.http_response_code {
							break
						}
						err++
					}
					if err.str != nil {
						len_ = core.ApPhpSlprintf(buf, g.SizeOf("buf"), "Status: %d %s", core.sapi_globals.sapi_headers.http_response_code, err.str)
					} else {
						len_ = core.ApPhpSlprintf(buf, g.SizeOf("buf"), "Status: %d", core.sapi_globals.sapi_headers.http_response_code)
					}
				}
			}
		}
		if has_status == 0 {
			core.PhpOutputWriteUnbuffered(buf, len_)
			core.PhpOutputWriteUnbuffered("\r\n", 2)
			ignore_status = 1
		}
	}
	h = (*core.SapiHeader)(zend.ZendLlistGetFirstEx(&sapi_headers.headers, &pos))
	for h != nil {

		/* prevent CRLFCRLF */

		if h.header_len != 0 {
			if h.header_len > g.SizeOf("\"Status:\"")-1 && strncasecmp(h.header, "Status:", g.SizeOf("\"Status:\"")-1) == 0 {
				if ignore_status == 0 {
					ignore_status = 1
					core.PhpOutputWriteUnbuffered(h.header, h.header_len)
					core.PhpOutputWriteUnbuffered("\r\n", 2)
				}
			} else if response_status == 304 && h.header_len > g.SizeOf("\"Content-Type:\"")-1 && strncasecmp(h.header, "Content-Type:", g.SizeOf("\"Content-Type:\"")-1) == 0 {
				h = (*core.SapiHeader)(zend.ZendLlistGetNextEx(&sapi_headers.headers, &pos))
				continue
			} else {
				core.PhpOutputWriteUnbuffered(h.header, h.header_len)
				core.PhpOutputWriteUnbuffered("\r\n", 2)
			}
		}
		h = (*core.SapiHeader)(zend.ZendLlistGetNextEx(&sapi_headers.headers, &pos))
	}
	core.PhpOutputWriteUnbuffered("\r\n", 2)
	return 1
}

// #define STDIN_FILENO       0

func SapiCgiReadPost(buffer *byte, count_bytes int) int {
	var read_bytes int = 0
	var tmp_read_bytes int
	var remaining_bytes int
	r.Assert(core.sapi_globals.request_info.content_length >= core.sapi_globals.read_post_bytes)
	remaining_bytes = size_t(core.sapi_globals.request_info.content_length - core.sapi_globals.read_post_bytes)
	if count_bytes < remaining_bytes {
		count_bytes = count_bytes
	} else {
		count_bytes = remaining_bytes
	}
	for read_bytes < count_bytes {
		tmp_read_bytes = read(0, buffer+read_bytes, count_bytes-read_bytes)
		if tmp_read_bytes <= 0 {
			break
		}
		read_bytes += tmp_read_bytes
	}
	return read_bytes
}
func SapiFcgiReadPost(buffer *byte, count_bytes int) int {
	var read_bytes int = 0
	var tmp_read_bytes int
	var request *core.FcgiRequest = (*core.FcgiRequest)(core.sapi_globals.server_context)
	var remaining int = core.sapi_globals.request_info.content_length - core.sapi_globals.read_post_bytes
	if remaining < count_bytes {
		count_bytes = remaining
	}
	for read_bytes < count_bytes {
		var diff int = count_bytes - read_bytes
		var to_read int = g.CondF2(diff > 2147483647, 2147483647, func() int { return int(diff) })
		tmp_read_bytes = core.FcgiRead(request, buffer+read_bytes, to_read)
		if tmp_read_bytes <= 0 {
			break
		}
		read_bytes += tmp_read_bytes
	}
	return read_bytes
}
func SapiCgiGetenv(name *byte, name_len int) *byte { return getenv(name) }
func SapiFcgiGetenv(name *byte, name_len int) *byte {
	/* when php is started by mod_fastcgi, no regular environment
	 * is provided to PHP.  It is always sent to PHP at the start
	 * of a request.  So we have to do our own lookup to get env
	 * vars.  This could probably be faster somehow.  */

	var request *core.FcgiRequest = (*core.FcgiRequest)(core.sapi_globals.server_context)
	var ret *byte = core.FcgiGetenv(request, name, int(name_len))
	if ret != nil {
		return ret
	}

	/*  if cgi, or fastcgi and not found in fcgi env
	    check the regular environment */

	return getenv(name)

	/*  if cgi, or fastcgi and not found in fcgi env
	    check the regular environment */
}
func _sapiCgiPutenv(name string, name_len int, value *byte) *byte {
	if value != nil {
		setenv(name, value, 1)
	}
	if value == nil {
		unsetenv(name)
	}
	return getenv(name)
}
func SapiCgiReadCookies() *byte { return getenv("HTTP_COOKIE") }
func SapiFcgiReadCookies() *byte {
	var request *core.FcgiRequest = (*core.FcgiRequest)(core.sapi_globals.server_context)
	return core.FcgiQuickGetenv(request, "HTTP_COOKIE", g.SizeOf("\"HTTP_COOKIE\"")-1, g.CondF(g.SizeOf("\"HTTP_COOKIE\"")-1 < 3, func() uint { return uint(g.SizeOf("\"HTTP_COOKIE\"") - 1) }, func() int {
		return (uint("HTTP_COOKIE"[3]) << 2) + (uint("HTTP_COOKIE"[g.SizeOf("\"HTTP_COOKIE\"")-1-2]) << 4) + (uint("HTTP_COOKIE"[g.SizeOf("\"HTTP_COOKIE\"")-1-1]) << 2) + g.SizeOf("\"HTTP_COOKIE\"") - 1
	}))
}
func CgiPhpLoadEnvVar(var_ *byte, var_len uint, val *byte, val_len uint, arg any) {
	var array_ptr *zend.Zval = (*zend.Zval)(arg)
	var filter_arg int = g.Cond(array_ptr.value.arr == core.CoreGlobals.http_globals[4].value.arr, 4, 5)
	var new_val_len int
	if core.sapi_module.input_filter(filter_arg, var_, &val, strlen(val), &new_val_len) != 0 {
		core.PhpRegisterVariableSafe(var_, val, new_val_len, array_ptr)
	}
}
func CgiPhpImportEnvironmentVariables(array_ptr *zend.Zval) {
	if core.CoreGlobals.variables_order != nil && (strchr(core.CoreGlobals.variables_order, 'E') || strchr(core.CoreGlobals.variables_order, 'e')) {
		if core.CoreGlobals.http_globals[4].u1.v.type_ != 7 {
			zend.ZendIsAutoGlobalStr("_ENV", g.SizeOf("\"_ENV\"")-1)
		}
		if core.CoreGlobals.http_globals[4].u1.v.type_ == 7 && array_ptr.value.arr != core.CoreGlobals.http_globals[4].value.arr {
			zend.ZendArrayDestroy(array_ptr.value.arr)
			array_ptr.value.arr = zend.ZendArrayDup(core.CoreGlobals.http_globals[4].value.arr)
			return
		}
	}

	/* call php's original import as a catch-all */

	PhpPhpImportEnvironmentVariables(array_ptr)
	if core.FcgiIsFastcgi() != 0 {
		var request *core.FcgiRequest = (*core.FcgiRequest)(core.sapi_globals.server_context)
		core.FcgiLoadenv(request, CgiPhpLoadEnvVar, array_ptr)
	}
}
func SapiCgiRegisterVariables(track_vars_array *zend.Zval) {
	var php_self_len int
	var php_self *byte

	/* In CGI mode, we consider the environment to be a part of the server
	 * variables
	 */

	core.PhpImportEnvironmentVariables(track_vars_array)
	if php_cgi_globals.GetFixPathinfo() != 0 {
		var script_name *byte = core.sapi_globals.request_info.request_uri
		var path_info *byte
		var free_php_self int
		if core.FcgiIsFastcgi() != 0 {
			var request *core.FcgiRequest = (*core.FcgiRequest)(core.sapi_globals.server_context)
			path_info = core.FcgiQuickGetenv(request, "PATH_INFO", g.SizeOf("\"PATH_INFO\"")-1, g.CondF(g.SizeOf("\"PATH_INFO\"")-1 < 3, func() uint { return uint(g.SizeOf("\"PATH_INFO\"") - 1) }, func() int {
				return (uint("PATH_INFO"[3]) << 2) + (uint("PATH_INFO"[g.SizeOf("\"PATH_INFO\"")-1-2]) << 4) + (uint("PATH_INFO"[g.SizeOf("\"PATH_INFO\"")-1-1]) << 2) + g.SizeOf("\"PATH_INFO\"") - 1
			}))
		} else {
			path_info = getenv("PATH_INFO")
		}
		if path_info != nil {
			var path_info_len int = strlen(path_info)
			if script_name != nil {
				var script_name_len int = strlen(script_name)
				php_self_len = script_name_len + path_info_len
				php_self = zend._emalloc(php_self_len + 1)
				memcpy(php_self, script_name, script_name_len+1)
				memcpy(php_self+script_name_len, path_info, path_info_len+1)
				free_php_self = 1
			} else {
				php_self = path_info
				php_self_len = path_info_len
				free_php_self = 0
			}
		} else if script_name != nil {
			php_self = script_name
			php_self_len = strlen(script_name)
			free_php_self = 0
		} else {
			php_self = ""
			php_self_len = 0
			free_php_self = 0
		}

		/* Build the special-case PHP_SELF variable for the CGI version */

		if core.sapi_module.input_filter(5, "PHP_SELF", &php_self, php_self_len, &php_self_len) != 0 {
			core.PhpRegisterVariableSafe("PHP_SELF", php_self, php_self_len, track_vars_array)
		}
		if free_php_self != 0 {
			zend._efree(php_self)
		}
	} else {
		if core.sapi_globals.request_info.request_uri != nil {
			php_self = core.sapi_globals.request_info.request_uri
		} else {
			php_self = ""
		}
		php_self_len = strlen(php_self)
		if core.sapi_module.input_filter(5, "PHP_SELF", &php_self, php_self_len, &php_self_len) != 0 {
			core.PhpRegisterVariableSafe("PHP_SELF", php_self, php_self_len, track_vars_array)
		}
	}
}
func SapiCgiLogMessage(message *byte, syslog_type_int int) {
	if core.FcgiIsFastcgi() != 0 && php_cgi_globals.GetFcgiLogging() != 0 {
		var request *core.FcgiRequest
		request = (*core.FcgiRequest)(core.sapi_globals.server_context)
		if request != nil {
			var ret int
			var len_ int = int(strlen(message))
			var buf *byte = zend.Malloc(len_ + 2)
			memcpy(buf, message, len_)
			memcpy(buf+len_, "\n", g.SizeOf("\"\\n\""))
			ret = core.FcgiWrite(request, core.FCGI_STDERR, buf, int(len_+1))
			zend.Free(buf)
			if ret < 0 {
				core.PhpHandleAbortedConnection()
			}
		} else {
			r.Fprintf(stderr, "%s\n", message)
		}
	} else {
		r.Fprintf(stderr, "%s\n", message)
	}
}

/* {{{ php_cgi_ini_activate_user_config
 */

func PhpCgiIniActivateUserConfig(path *byte, path_len int, doc_root *byte, doc_root_len int) {
	var new_entry *UserConfigCacheEntry
	var entry *UserConfigCacheEntry
	var request_time int64 = int64(core.SapiGetRequestTime())

	/* Find cached config entry: If not found, create one */

	if g.Assign(&entry, zend.ZendHashStrFindPtr(&(php_cgi_globals.GetUserConfigCache()), path, path_len)) == nil {
		new_entry = zend.__zendMalloc(g.SizeOf("user_config_cache_entry"))
		new_entry.SetExpires(0)
		new_entry.SetUserConfig((*zend.HashTable)(g.CondF(true, func() any { return zend.__zendMalloc(g.SizeOf("HashTable")) }, func() any { return zend._emalloc(g.SizeOf("HashTable")) })))
		zend._zendHashInit(new_entry.GetUserConfig(), 8, zend.DtorFuncT(core.ConfigZvalDtor), 1)
		entry = zend.ZendHashStrUpdatePtr(&(php_cgi_globals.GetUserConfigCache()), path, path_len, new_entry)
	}

	/* Check whether cache entry has expired and rescan if it is */

	if request_time > entry.GetExpires() {
		var real_path *byte = nil
		var s1 *byte
		var s2 *byte
		var s_len int

		/* Clear the expired config */

		zend.ZendHashClean(entry.GetUserConfig())
		if path[0] != '/' {
			var real_path_len int
			real_path = zend.TsrmRealpath(path, nil)
			if real_path == nil {
				return
			}
			real_path_len = strlen(real_path)
			path = real_path
			path_len = real_path_len
		}
		if path_len > doc_root_len {
			s1 = (*byte)(doc_root)
			s2 = path
			s_len = doc_root_len
		} else {
			s1 = path
			s2 = (*byte)(doc_root)
			s_len = path_len
		}

		/* we have to test if path is part of DOCUMENT_ROOT.
		   if it is inside the docroot, we scan the tree up to the docroot
		     to find more user.ini, if not we only scan the current path.
		*/

		if strncmp(s1, s2, s_len) == 0 {
			var ptr *byte = s2 + doc_root_len
			for g.Assign(&ptr, strchr(ptr, '/')) != nil {
				*ptr = 0
				core.PhpParseUserIniFile(path, core.CoreGlobals.user_ini_filename, entry.GetUserConfig())
				*ptr = '/'
				ptr++
			}
		} else {
			core.PhpParseUserIniFile(path, core.CoreGlobals.user_ini_filename, entry.GetUserConfig())
		}
		if real_path != nil {
			zend._efree(real_path)
		}
		entry.SetExpires(request_time + core.CoreGlobals.user_ini_cache_ttl)
	}

	/* Activate ini entries with values from the user config hash */

	core.PhpIniActivateConfig(entry.GetUserConfig(), 1<<1, 1<<5)

	/* Activate ini entries with values from the user config hash */
}

/* }}} */

func SapiCgiActivate() int {
	/* PATH_TRANSLATED should be defined at this stage but better safe than sorry :) */

	if core.sapi_globals.request_info.path_translated == nil {
		return zend.FAILURE
	}
	if core.PhpIniHasPerHostConfig() != 0 {
		var server_name *byte

		/* Activate per-host-system-configuration defined in php.ini and stored into configuration_hash during startup */

		if core.FcgiIsFastcgi() != 0 {
			var request *core.FcgiRequest = (*core.FcgiRequest)(core.sapi_globals.server_context)
			server_name = core.FcgiQuickGetenv(request, "SERVER_NAME", g.SizeOf("\"SERVER_NAME\"")-1, g.CondF(g.SizeOf("\"SERVER_NAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"SERVER_NAME\"") - 1) }, func() int {
				return (uint("SERVER_NAME"[3]) << 2) + (uint("SERVER_NAME"[g.SizeOf("\"SERVER_NAME\"")-1-2]) << 4) + (uint("SERVER_NAME"[g.SizeOf("\"SERVER_NAME\"")-1-1]) << 2) + g.SizeOf("\"SERVER_NAME\"") - 1
			}))
		} else {
			server_name = getenv("SERVER_NAME")
		}

		/* SERVER_NAME should also be defined at this stage..but better check it anyway */

		if server_name != nil {
			var server_name_len int = strlen(server_name)
			server_name = zend._estrndup(server_name, server_name_len)
			zend.ZendStrTolower(server_name, server_name_len)
			core.PhpIniActivatePerHostConfig(server_name, server_name_len)
			zend._efree(server_name)
		}

		/* SERVER_NAME should also be defined at this stage..but better check it anyway */

	}
	if core.PhpIniHasPerDirConfig() != 0 || core.CoreGlobals.user_ini_filename != nil && (*(core.CoreGlobals.user_ini_filename)) {
		var path *byte
		var path_len int

		/* Prepare search path */

		path_len = strlen(core.sapi_globals.request_info.path_translated)

		/* Make sure we have trailing slash! */

		if core.sapi_globals.request_info.path_translated[path_len] != '/' {
			path = zend._emalloc(path_len + 2)
			memcpy(path, core.sapi_globals.request_info.path_translated, path_len+1)
			path_len = zend.ZendDirname(path, path_len)
			path[g.PostInc(&path_len)] = '/'
		} else {
			path = zend._estrndup(core.sapi_globals.request_info.path_translated, path_len)
			path_len = zend.ZendDirname(path, path_len)
		}
		path[path_len] = 0

		/* Activate per-dir-system-configuration defined in php.ini and stored into configuration_hash during startup */

		core.PhpIniActivatePerDirConfig(path, path_len)

		/* Load and activate user ini files in path starting from DOCUMENT_ROOT */

		if core.CoreGlobals.user_ini_filename != nil && (*(core.CoreGlobals.user_ini_filename)) {
			var doc_root *byte
			if core.FcgiIsFastcgi() != 0 {
				var request *core.FcgiRequest = (*core.FcgiRequest)(core.sapi_globals.server_context)
				doc_root = core.FcgiQuickGetenv(request, "DOCUMENT_ROOT", g.SizeOf("\"DOCUMENT_ROOT\"")-1, g.CondF(g.SizeOf("\"DOCUMENT_ROOT\"")-1 < 3, func() uint { return uint(g.SizeOf("\"DOCUMENT_ROOT\"") - 1) }, func() int {
					return (uint("DOCUMENT_ROOT"[3]) << 2) + (uint("DOCUMENT_ROOT"[g.SizeOf("\"DOCUMENT_ROOT\"")-1-2]) << 4) + (uint("DOCUMENT_ROOT"[g.SizeOf("\"DOCUMENT_ROOT\"")-1-1]) << 2) + g.SizeOf("\"DOCUMENT_ROOT\"") - 1
				}))
			} else {
				doc_root = getenv("DOCUMENT_ROOT")
			}

			/* DOCUMENT_ROOT should also be defined at this stage..but better check it anyway */

			if doc_root != nil {
				var doc_root_len int = strlen(doc_root)
				if doc_root_len > 0 && doc_root[doc_root_len-1] == '/' {
					doc_root_len--
				}
				PhpCgiIniActivateUserConfig(path, path_len, doc_root, doc_root_len)
			}

			/* DOCUMENT_ROOT should also be defined at this stage..but better check it anyway */

		}
		zend._efree(path)
	}
	return zend.SUCCESS
}
func SapiCgiDeactivate() int {
	/* flush only when SAPI was started. The reasons are:
	   1. SAPI Deactivate is called from two places: module init and request shutdown
	   2. When the first call occurs and the request is not set up, flush fails on FastCGI.
	*/

	if core.sapi_globals.sapi_started != 0 {
		if core.FcgiIsFastcgi() != 0 {
			if Parent == 0 && core.FcgiFinishRequest((*core.FcgiRequest)(core.sapi_globals.server_context), 0) == 0 {
				core.PhpHandleAbortedConnection()
			}
		} else {
			SapiCgiFlush(core.sapi_globals.server_context)
		}
	}
	return zend.SUCCESS
}
func PhpCgiStartup(sapi_module *core.sapi_module_struct) int {
	if core.PhpModuleStartup(sapi_module, &CgiModuleEntry, 1) == zend.FAILURE {
		return zend.FAILURE
	}
	return zend.SUCCESS
}

/* {{{ sapi_module_struct cgi_sapi_module
 */

var CgiSapiModule core.sapi_module_struct = core.sapi_module_struct{"cgi-fcgi", "CGI/FastCGI", PhpCgiStartup, core.PhpModuleShutdownWrapper, SapiCgiActivate, SapiCgiDeactivate, SapiCgiUbWrite, SapiCgiFlush, nil, SapiCgiGetenv, zend.ZendError, nil, SapiCgiSendHeaders, nil, SapiCgiReadPost, SapiCgiReadCookies, SapiCgiRegisterVariables, SapiCgiLogMessage, nil, nil, nil, nil, nil, nil, 0, 0, nil, nil, nil, nil, nil, nil, 0, nil, nil, nil}

/* }}} */

var ArginfoDl []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"extension_filename", 0, 0, 0}}

/* }}} */

var AdditionalFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"dl",
		standard.ZifDl,
		ArginfoDl,
		uint32(g.SizeOf("arginfo_dl")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

/* {{{ php_cgi_usage
 */

func PhpCgiUsage(argv0 *byte) {
	var prog *byte
	prog = strrchr(argv0, '/')
	if prog != nil {
		prog++
	} else {
		prog = "php"
	}
	core.PhpPrintf("Usage: %s [-q] [-h] [-s] [-v] [-i] [-f <file>]\n"+"       %s <file> [args...]\n"+"  -a               Run interactively\n"+"  -b <address:port>|<port> Bind Path for external FASTCGI Server mode\n"+"  -C               Do not chdir to the script's directory\n"+"  -c <path>|<file> Look for php.ini file in this directory\n"+"  -n               No php.ini file will be used\n"+"  -d foo[=bar]     Define INI entry foo with value 'bar'\n"+"  -e               Generate extended information for debugger/profiler\n"+"  -f <file>        Parse <file>.  Implies `-q'\n"+"  -h               This help\n"+"  -i               PHP information\n"+"  -l               Syntax check only (lint)\n"+"  -m               Show compiled in modules\n"+"  -q               Quiet-mode.  Suppress HTTP Header output.\n"+"  -s               Display colour syntax highlighted source.\n"+"  -v               Version number\n"+"  -w               Display source with stripped comments and whitespace.\n"+"  -z <file>        Load Zend extension <file>.\n"+"  -T <count>       Measure execution time of script repeated <count> times.\n", prog, prog)
}

/* }}} */

func IsValidPath(path *byte) int {
	var p *byte = path
	if p == nil {
		return 0
	}
	if (*p) == '.' && (*(p + 1)) == '.' && (!(*(p + 2)) || (*(p + 2)) == '/') {
		return 0
	}
	for *p {
		if (*p) == '/' {
			p++
			if (*p) == '.' {
				p++
				if (*p) == '.' {
					p++
					if !(*p) || (*p) == '/' {
						return 0
					}
				}
			}
		}
		p++
	}
	return 1
}

/* }}} */

// #define CGI_GETENV(name) ( ( has_env ) ? FCGI_GETENV ( request , name ) : getenv ( name ) )

// #define CGI_PUTENV(name,value) ( ( has_env ) ? FCGI_PUTENV ( request , name , value ) : _sapi_cgi_putenv ( name , sizeof ( name ) - 1 , value ) )

/* {{{ init_request_info

initializes request_info structure

specificly in this section we handle proper translations
for:

PATH_INFO
  derived from the portion of the URI path following
  the script name but preceding any query data
  may be empty

PATH_TRANSLATED
  derived by taking any path-info component of the
  request URI and performing any virtual-to-physical
  translation appropriate to map it onto the server's
  document repository structure

  empty if PATH_INFO is empty

  The env var PATH_TRANSLATED **IS DIFFERENT** than the
  request_info.path_translated variable, the latter should
  match SCRIPT_FILENAME instead.

SCRIPT_NAME
  set to a URL path that could identify the CGI script
  rather than the interpreter.  PHP_SELF is set to this

REQUEST_URI
  uri section following the domain:port part of a URI

SCRIPT_FILENAME
  The virtual-to-physical translation of SCRIPT_NAME (as per
  PATH_TRANSLATED)

These settings are documented at
http://cgi-spec.golux.com/


Based on the following URL request:

http://localhost/info.php/test?a=b

should produce, which btw is the same as if
we were running under mod_cgi on apache (ie. not
using ScriptAlias directives):

PATH_INFO=/test
PATH_TRANSLATED=/docroot/test
SCRIPT_NAME=/info.php
REQUEST_URI=/info.php/test?a=b
SCRIPT_FILENAME=/docroot/info.php
QUERY_STRING=a=b

but what we get is (cgi/mod_fastcgi under apache):

PATH_INFO=/info.php/test
PATH_TRANSLATED=/docroot/info.php/test
SCRIPT_NAME=/php/php-cgi  (from the Action setting I suppose)
REQUEST_URI=/info.php/test?a=b
SCRIPT_FILENAME=/path/to/php/bin/php-cgi  (Action setting translated)
QUERY_STRING=a=b

Comments in the code below refer to using the above URL in a request

*/

func InitRequestInfo(request *core.FcgiRequest) {
	var has_env int = core.FcgiHasEnv(request)
	var env_script_filename *byte = g.CondF(has_env != 0, func() *byte {
		return core.FcgiQuickGetenv(request, "SCRIPT_FILENAME", g.SizeOf("\"SCRIPT_FILENAME\"")-1, g.CondF(g.SizeOf("\"SCRIPT_FILENAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"SCRIPT_FILENAME\"") - 1) }, func() int {
			return (uint("SCRIPT_FILENAME"[3]) << 2) + (uint("SCRIPT_FILENAME"[g.SizeOf("\"SCRIPT_FILENAME\"")-1-2]) << 4) + (uint("SCRIPT_FILENAME"[g.SizeOf("\"SCRIPT_FILENAME\"")-1-1]) << 2) + g.SizeOf("\"SCRIPT_FILENAME\"") - 1
		}))
	}, func() __auto__ { return getenv("SCRIPT_FILENAME") })
	var env_path_translated *byte = g.CondF(has_env != 0, func() *byte {
		return core.FcgiQuickGetenv(request, "PATH_TRANSLATED", g.SizeOf("\"PATH_TRANSLATED\"")-1, g.CondF(g.SizeOf("\"PATH_TRANSLATED\"")-1 < 3, func() uint { return uint(g.SizeOf("\"PATH_TRANSLATED\"") - 1) }, func() int {
			return (uint("PATH_TRANSLATED"[3]) << 2) + (uint("PATH_TRANSLATED"[g.SizeOf("\"PATH_TRANSLATED\"")-1-2]) << 4) + (uint("PATH_TRANSLATED"[g.SizeOf("\"PATH_TRANSLATED\"")-1-1]) << 2) + g.SizeOf("\"PATH_TRANSLATED\"") - 1
		}))
	}, func() __auto__ { return getenv("PATH_TRANSLATED") })
	var script_path_translated *byte = env_script_filename

	/* some broken servers do not have script_filename or argv0
	 * an example, IIS configured in some ways.  then they do more
	 * broken stuff and set path_translated to the cgi script location */

	if script_path_translated == nil && env_path_translated != nil {
		script_path_translated = env_path_translated
	}

	/* initialize the defaults */

	core.sapi_globals.request_info.path_translated = nil
	core.sapi_globals.request_info.request_method = nil
	core.sapi_globals.request_info.proto_num = 1000
	core.sapi_globals.request_info.query_string = nil
	core.sapi_globals.request_info.request_uri = nil
	core.sapi_globals.request_info.content_type = nil
	core.sapi_globals.request_info.content_length = 0
	core.sapi_globals.sapi_headers.http_response_code = 200

	/* script_path_translated being set is a good indication that
	 * we are running in a cgi environment, since it is always
	 * null otherwise.  otherwise, the filename
	 * of the script will be retrieved later via argc/argv */

	if script_path_translated != nil {
		var auth *byte
		var content_length *byte = g.CondF(has_env != 0, func() *byte {
			return core.FcgiQuickGetenv(request, "CONTENT_LENGTH", g.SizeOf("\"CONTENT_LENGTH\"")-1, g.CondF(g.SizeOf("\"CONTENT_LENGTH\"")-1 < 3, func() uint { return uint(g.SizeOf("\"CONTENT_LENGTH\"") - 1) }, func() int {
				return (uint("CONTENT_LENGTH"[3]) << 2) + (uint("CONTENT_LENGTH"[g.SizeOf("\"CONTENT_LENGTH\"")-1-2]) << 4) + (uint("CONTENT_LENGTH"[g.SizeOf("\"CONTENT_LENGTH\"")-1-1]) << 2) + g.SizeOf("\"CONTENT_LENGTH\"") - 1
			}))
		}, func() __auto__ { return getenv("CONTENT_LENGTH") })
		var content_type *byte = g.CondF(has_env != 0, func() *byte {
			return core.FcgiQuickGetenv(request, "CONTENT_TYPE", g.SizeOf("\"CONTENT_TYPE\"")-1, g.CondF(g.SizeOf("\"CONTENT_TYPE\"")-1 < 3, func() uint { return uint(g.SizeOf("\"CONTENT_TYPE\"") - 1) }, func() int {
				return (uint("CONTENT_TYPE"[3]) << 2) + (uint("CONTENT_TYPE"[g.SizeOf("\"CONTENT_TYPE\"")-1-2]) << 4) + (uint("CONTENT_TYPE"[g.SizeOf("\"CONTENT_TYPE\"")-1-1]) << 2) + g.SizeOf("\"CONTENT_TYPE\"") - 1
			}))
		}, func() __auto__ { return getenv("CONTENT_TYPE") })
		var env_path_info *byte = g.CondF(has_env != 0, func() *byte {
			return core.FcgiQuickGetenv(request, "PATH_INFO", g.SizeOf("\"PATH_INFO\"")-1, g.CondF(g.SizeOf("\"PATH_INFO\"")-1 < 3, func() uint { return uint(g.SizeOf("\"PATH_INFO\"") - 1) }, func() int {
				return (uint("PATH_INFO"[3]) << 2) + (uint("PATH_INFO"[g.SizeOf("\"PATH_INFO\"")-1-2]) << 4) + (uint("PATH_INFO"[g.SizeOf("\"PATH_INFO\"")-1-1]) << 2) + g.SizeOf("\"PATH_INFO\"") - 1
			}))
		}, func() __auto__ { return getenv("PATH_INFO") })
		var env_script_name *byte = g.CondF(has_env != 0, func() *byte {
			return core.FcgiQuickGetenv(request, "SCRIPT_NAME", g.SizeOf("\"SCRIPT_NAME\"")-1, g.CondF(g.SizeOf("\"SCRIPT_NAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"SCRIPT_NAME\"") - 1) }, func() int {
				return (uint("SCRIPT_NAME"[3]) << 2) + (uint("SCRIPT_NAME"[g.SizeOf("\"SCRIPT_NAME\"")-1-2]) << 4) + (uint("SCRIPT_NAME"[g.SizeOf("\"SCRIPT_NAME\"")-1-1]) << 2) + g.SizeOf("\"SCRIPT_NAME\"") - 1
			}))
		}, func() __auto__ { return getenv("SCRIPT_NAME") })
		if php_cgi_globals.GetFixPathinfo() != 0 {
			var st zend.ZendStatT
			var real_path *byte = nil
			var env_redirect_url *byte = g.CondF(has_env != 0, func() *byte {
				return core.FcgiQuickGetenv(request, "REDIRECT_URL", g.SizeOf("\"REDIRECT_URL\"")-1, g.CondF(g.SizeOf("\"REDIRECT_URL\"")-1 < 3, func() uint { return uint(g.SizeOf("\"REDIRECT_URL\"") - 1) }, func() int {
					return (uint("REDIRECT_URL"[3]) << 2) + (uint("REDIRECT_URL"[g.SizeOf("\"REDIRECT_URL\"")-1-2]) << 4) + (uint("REDIRECT_URL"[g.SizeOf("\"REDIRECT_URL\"")-1-1]) << 2) + g.SizeOf("\"REDIRECT_URL\"") - 1
				}))
			}, func() __auto__ { return getenv("REDIRECT_URL") })
			var env_document_root *byte = g.CondF(has_env != 0, func() *byte {
				return core.FcgiQuickGetenv(request, "DOCUMENT_ROOT", g.SizeOf("\"DOCUMENT_ROOT\"")-1, g.CondF(g.SizeOf("\"DOCUMENT_ROOT\"")-1 < 3, func() uint { return uint(g.SizeOf("\"DOCUMENT_ROOT\"") - 1) }, func() int {
					return (uint("DOCUMENT_ROOT"[3]) << 2) + (uint("DOCUMENT_ROOT"[g.SizeOf("\"DOCUMENT_ROOT\"")-1-2]) << 4) + (uint("DOCUMENT_ROOT"[g.SizeOf("\"DOCUMENT_ROOT\"")-1-1]) << 2) + g.SizeOf("\"DOCUMENT_ROOT\"") - 1
				}))
			}, func() __auto__ { return getenv("DOCUMENT_ROOT") })
			var orig_path_translated *byte = env_path_translated
			var orig_path_info *byte = env_path_info
			var orig_script_name *byte = env_script_name
			var orig_script_filename *byte = env_script_filename
			var script_path_translated_len int
			if env_document_root == nil && core.CoreGlobals.doc_root != nil {
				if has_env != 0 {
					env_document_root = core.FcgiQuickPutenv(request, "DOCUMENT_ROOT", g.SizeOf("\"DOCUMENT_ROOT\"")-1, g.CondF(g.SizeOf("\"DOCUMENT_ROOT\"")-1 < 3, func() uint { return uint(g.SizeOf("\"DOCUMENT_ROOT\"") - 1) }, func() int {
						return (uint("DOCUMENT_ROOT"[3]) << 2) + (uint("DOCUMENT_ROOT"[g.SizeOf("\"DOCUMENT_ROOT\"")-1-2]) << 4) + (uint("DOCUMENT_ROOT"[g.SizeOf("\"DOCUMENT_ROOT\"")-1-1]) << 2) + g.SizeOf("\"DOCUMENT_ROOT\"") - 1
					}), core.CoreGlobals.doc_root)
				} else {
					env_document_root = _sapiCgiPutenv("DOCUMENT_ROOT", g.SizeOf("\"DOCUMENT_ROOT\"")-1, core.CoreGlobals.doc_root)
				}

				/* fix docroot */

				/* fix docroot */

			}
			if env_path_translated != nil && env_redirect_url != nil && env_path_translated != script_path_translated && strcmp(env_path_translated, script_path_translated) != 0 {

				/*
				 * pretty much apache specific.  If we have a redirect_url
				 * then our script_filename and script_name point to the
				 * php executable
				 */

				script_path_translated = env_path_translated

				/* we correct SCRIPT_NAME now in case we don't have PATH_INFO */

				env_script_name = env_redirect_url

				/* we correct SCRIPT_NAME now in case we don't have PATH_INFO */

			}

			/*
			 * if the file doesn't exist, try to extract PATH_INFO out
			 * of it by stat'ing back through the '/'
			 * this fixes url's like /info.php/test
			 */

			if script_path_translated != nil && g.Assign(&script_path_translated_len, strlen(script_path_translated)) > 0 && (script_path_translated[script_path_translated_len-1] == '/' || g.Assign(&real_path, zend.TsrmRealpath(script_path_translated, nil)) == nil) {
				var pt *byte = zend._estrndup(script_path_translated, script_path_translated_len)
				var len_ int = script_path_translated_len
				var ptr *byte
				for g.Assign(&ptr, strrchr(pt, '/')) || g.Assign(&ptr, strrchr(pt, '\\')) {
					*ptr = 0
					if stat(pt, &st) == 0 && (st.st_mode&S_IFMT) == S_IFREG {

						/*
						 * okay, we found the base script!
						 * work out how many chars we had to strip off;
						 * then we can modify PATH_INFO
						 * accordingly
						 *
						 * we now have the makings of
						 * PATH_INFO=/test
						 * SCRIPT_FILENAME=/docroot/info.php
						 *
						 * we now need to figure out what docroot is.
						 * if DOCUMENT_ROOT is set, this is easy, otherwise,
						 * we have to play the game of hide and seek to figure
						 * out what SCRIPT_NAME should be
						 */

						var slen int = len_ - strlen(pt)
						var pilen int = g.CondF1(env_path_info != nil, func() __auto__ { return strlen(env_path_info) }, 0)
						var path_info *byte = g.Cond(env_path_info != nil, env_path_info+pilen-slen, nil)
						if orig_path_info != path_info {
							if orig_path_info != nil {
								var old byte
								g.CondF(has_env != 0, func() *byte {
									return core.FcgiQuickPutenv(request, "ORIG_PATH_INFO", g.SizeOf("\"ORIG_PATH_INFO\"")-1, g.CondF(g.SizeOf("\"ORIG_PATH_INFO\"")-1 < 3, func() uint { return uint(g.SizeOf("\"ORIG_PATH_INFO\"") - 1) }, func() int {
										return (uint("ORIG_PATH_INFO"[3]) << 2) + (uint("ORIG_PATH_INFO"[g.SizeOf("\"ORIG_PATH_INFO\"")-1-2]) << 4) + (uint("ORIG_PATH_INFO"[g.SizeOf("\"ORIG_PATH_INFO\"")-1-1]) << 2) + g.SizeOf("\"ORIG_PATH_INFO\"") - 1
									}), orig_path_info)
								}, func() *byte {
									return _sapiCgiPutenv("ORIG_PATH_INFO", g.SizeOf("\"ORIG_PATH_INFO\"")-1, orig_path_info)
								})
								old = path_info[0]
								path_info[0] = 0
								if orig_script_name == nil || strcmp(orig_script_name, env_path_info) != 0 {
									if orig_script_name != nil {
										g.CondF(has_env != 0, func() *byte {
											return core.FcgiQuickPutenv(request, "ORIG_SCRIPT_NAME", g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1, g.CondF(g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"ORIG_SCRIPT_NAME\"") - 1) }, func() int {
												return (uint("ORIG_SCRIPT_NAME"[3]) << 2) + (uint("ORIG_SCRIPT_NAME"[g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1-2]) << 4) + (uint("ORIG_SCRIPT_NAME"[g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1-1]) << 2) + g.SizeOf("\"ORIG_SCRIPT_NAME\"") - 1
											}), orig_script_name)
										}, func() *byte {
											return _sapiCgiPutenv("ORIG_SCRIPT_NAME", g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1, orig_script_name)
										})
									}
									if has_env != 0 {
										core.sapi_globals.request_info.request_uri = core.FcgiQuickPutenv(request, "SCRIPT_NAME", g.SizeOf("\"SCRIPT_NAME\"")-1, g.CondF(g.SizeOf("\"SCRIPT_NAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"SCRIPT_NAME\"") - 1) }, func() int {
											return (uint("SCRIPT_NAME"[3]) << 2) + (uint("SCRIPT_NAME"[g.SizeOf("\"SCRIPT_NAME\"")-1-2]) << 4) + (uint("SCRIPT_NAME"[g.SizeOf("\"SCRIPT_NAME\"")-1-1]) << 2) + g.SizeOf("\"SCRIPT_NAME\"") - 1
										}), env_path_info)
									} else {
										core.sapi_globals.request_info.request_uri = _sapiCgiPutenv("SCRIPT_NAME", g.SizeOf("\"SCRIPT_NAME\"")-1, env_path_info)
									}
								} else {
									core.sapi_globals.request_info.request_uri = orig_script_name
								}
								path_info[0] = old
							}
							if has_env != 0 {
								env_path_info = core.FcgiQuickPutenv(request, "PATH_INFO", g.SizeOf("\"PATH_INFO\"")-1, g.CondF(g.SizeOf("\"PATH_INFO\"")-1 < 3, func() uint { return uint(g.SizeOf("\"PATH_INFO\"") - 1) }, func() int {
									return (uint("PATH_INFO"[3]) << 2) + (uint("PATH_INFO"[g.SizeOf("\"PATH_INFO\"")-1-2]) << 4) + (uint("PATH_INFO"[g.SizeOf("\"PATH_INFO\"")-1-1]) << 2) + g.SizeOf("\"PATH_INFO\"") - 1
								}), path_info)
							} else {
								env_path_info = _sapiCgiPutenv("PATH_INFO", g.SizeOf("\"PATH_INFO\"")-1, path_info)
							}
						}
						if orig_script_filename == nil || strcmp(orig_script_filename, pt) != 0 {
							if orig_script_filename != nil {
								g.CondF(has_env != 0, func() *byte {
									return core.FcgiQuickPutenv(request, "ORIG_SCRIPT_FILENAME", g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1, g.CondF(g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"ORIG_SCRIPT_FILENAME\"") - 1) }, func() int {
										return (uint("ORIG_SCRIPT_FILENAME"[3]) << 2) + (uint("ORIG_SCRIPT_FILENAME"[g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1-2]) << 4) + (uint("ORIG_SCRIPT_FILENAME"[g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1-1]) << 2) + g.SizeOf("\"ORIG_SCRIPT_FILENAME\"") - 1
									}), orig_script_filename)
								}, func() *byte {
									return _sapiCgiPutenv("ORIG_SCRIPT_FILENAME", g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1, orig_script_filename)
								})
							}
							if has_env != 0 {
								script_path_translated = core.FcgiQuickPutenv(request, "SCRIPT_FILENAME", g.SizeOf("\"SCRIPT_FILENAME\"")-1, g.CondF(g.SizeOf("\"SCRIPT_FILENAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"SCRIPT_FILENAME\"") - 1) }, func() int {
									return (uint("SCRIPT_FILENAME"[3]) << 2) + (uint("SCRIPT_FILENAME"[g.SizeOf("\"SCRIPT_FILENAME\"")-1-2]) << 4) + (uint("SCRIPT_FILENAME"[g.SizeOf("\"SCRIPT_FILENAME\"")-1-1]) << 2) + g.SizeOf("\"SCRIPT_FILENAME\"") - 1
								}), pt)
							} else {
								script_path_translated = _sapiCgiPutenv("SCRIPT_FILENAME", g.SizeOf("\"SCRIPT_FILENAME\"")-1, pt)
							}
						}

						/* figure out docroot
						 * SCRIPT_FILENAME minus SCRIPT_NAME
						 */

						if env_document_root != nil {
							var l int = strlen(env_document_root)
							var path_translated_len int = 0
							var path_translated *byte = nil
							if l != 0 && env_document_root[l-1] == '/' {
								l--
							}

							/* we have docroot, so we should have:
							 * DOCUMENT_ROOT=/docroot
							 * SCRIPT_FILENAME=/docroot/info.php
							 */

							path_translated_len = l + g.CondF1(env_path_info != nil, func() __auto__ { return strlen(env_path_info) }, 0)
							path_translated = (*byte)(zend._emalloc(path_translated_len + 1))
							memcpy(path_translated, env_document_root, l)
							if env_path_info != nil {
								memcpy(path_translated+l, env_path_info, path_translated_len-l)
							}
							path_translated[path_translated_len] = '0'
							if orig_path_translated != nil {
								g.CondF(has_env != 0, func() *byte {
									return core.FcgiQuickPutenv(request, "ORIG_PATH_TRANSLATED", g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1, g.CondF(g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1 < 3, func() uint { return uint(g.SizeOf("\"ORIG_PATH_TRANSLATED\"") - 1) }, func() int {
										return (uint("ORIG_PATH_TRANSLATED"[3]) << 2) + (uint("ORIG_PATH_TRANSLATED"[g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1-2]) << 4) + (uint("ORIG_PATH_TRANSLATED"[g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1-1]) << 2) + g.SizeOf("\"ORIG_PATH_TRANSLATED\"") - 1
									}), orig_path_translated)
								}, func() *byte {
									return _sapiCgiPutenv("ORIG_PATH_TRANSLATED", g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1, orig_path_translated)
								})
							}
							if has_env != 0 {
								env_path_translated = core.FcgiQuickPutenv(request, "PATH_TRANSLATED", g.SizeOf("\"PATH_TRANSLATED\"")-1, g.CondF(g.SizeOf("\"PATH_TRANSLATED\"")-1 < 3, func() uint { return uint(g.SizeOf("\"PATH_TRANSLATED\"") - 1) }, func() int {
									return (uint("PATH_TRANSLATED"[3]) << 2) + (uint("PATH_TRANSLATED"[g.SizeOf("\"PATH_TRANSLATED\"")-1-2]) << 4) + (uint("PATH_TRANSLATED"[g.SizeOf("\"PATH_TRANSLATED\"")-1-1]) << 2) + g.SizeOf("\"PATH_TRANSLATED\"") - 1
								}), path_translated)
							} else {
								env_path_translated = _sapiCgiPutenv("PATH_TRANSLATED", g.SizeOf("\"PATH_TRANSLATED\"")-1, path_translated)
							}
							zend._efree(path_translated)
						} else if env_script_name != nil && strstr(pt, env_script_name) {

							/* PATH_TRANSLATED = PATH_TRANSLATED - SCRIPT_NAME + PATH_INFO */

							var ptlen int = strlen(pt) - strlen(env_script_name)
							var path_translated_len int = ptlen + g.CondF1(env_path_info != nil, func() __auto__ { return strlen(env_path_info) }, 0)
							var path_translated *byte = (*byte)(zend._emalloc(path_translated_len + 1))
							memcpy(path_translated, pt, ptlen)
							if env_path_info != nil {
								memcpy(path_translated+ptlen, env_path_info, path_translated_len-ptlen)
							}
							path_translated[path_translated_len] = '0'
							if orig_path_translated != nil {
								g.CondF(has_env != 0, func() *byte {
									return core.FcgiQuickPutenv(request, "ORIG_PATH_TRANSLATED", g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1, g.CondF(g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1 < 3, func() uint { return uint(g.SizeOf("\"ORIG_PATH_TRANSLATED\"") - 1) }, func() int {
										return (uint("ORIG_PATH_TRANSLATED"[3]) << 2) + (uint("ORIG_PATH_TRANSLATED"[g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1-2]) << 4) + (uint("ORIG_PATH_TRANSLATED"[g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1-1]) << 2) + g.SizeOf("\"ORIG_PATH_TRANSLATED\"") - 1
									}), orig_path_translated)
								}, func() *byte {
									return _sapiCgiPutenv("ORIG_PATH_TRANSLATED", g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1, orig_path_translated)
								})
							}
							if has_env != 0 {
								env_path_translated = core.FcgiQuickPutenv(request, "PATH_TRANSLATED", g.SizeOf("\"PATH_TRANSLATED\"")-1, g.CondF(g.SizeOf("\"PATH_TRANSLATED\"")-1 < 3, func() uint { return uint(g.SizeOf("\"PATH_TRANSLATED\"") - 1) }, func() int {
									return (uint("PATH_TRANSLATED"[3]) << 2) + (uint("PATH_TRANSLATED"[g.SizeOf("\"PATH_TRANSLATED\"")-1-2]) << 4) + (uint("PATH_TRANSLATED"[g.SizeOf("\"PATH_TRANSLATED\"")-1-1]) << 2) + g.SizeOf("\"PATH_TRANSLATED\"") - 1
								}), path_translated)
							} else {
								env_path_translated = _sapiCgiPutenv("PATH_TRANSLATED", g.SizeOf("\"PATH_TRANSLATED\"")-1, path_translated)
							}
							zend._efree(path_translated)
						}
						break
					}
				}
				if ptr == nil {

					/*
					 * if we stripped out all the '/' and still didn't find
					 * a valid path... we will fail, badly. of course we would
					 * have failed anyway... we output 'no input file' now.
					 */

					if orig_script_filename != nil {
						g.CondF(has_env != 0, func() *byte {
							return core.FcgiQuickPutenv(request, "ORIG_SCRIPT_FILENAME", g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1, g.CondF(g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"ORIG_SCRIPT_FILENAME\"") - 1) }, func() int {
								return (uint("ORIG_SCRIPT_FILENAME"[3]) << 2) + (uint("ORIG_SCRIPT_FILENAME"[g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1-2]) << 4) + (uint("ORIG_SCRIPT_FILENAME"[g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1-1]) << 2) + g.SizeOf("\"ORIG_SCRIPT_FILENAME\"") - 1
							}), orig_script_filename)
						}, func() *byte {
							return _sapiCgiPutenv("ORIG_SCRIPT_FILENAME", g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1, orig_script_filename)
						})
					}
					if has_env != 0 {
						script_path_translated = core.FcgiQuickPutenv(request, "SCRIPT_FILENAME", g.SizeOf("\"SCRIPT_FILENAME\"")-1, g.CondF(g.SizeOf("\"SCRIPT_FILENAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"SCRIPT_FILENAME\"") - 1) }, func() int {
							return (uint("SCRIPT_FILENAME"[3]) << 2) + (uint("SCRIPT_FILENAME"[g.SizeOf("\"SCRIPT_FILENAME\"")-1-2]) << 4) + (uint("SCRIPT_FILENAME"[g.SizeOf("\"SCRIPT_FILENAME\"")-1-1]) << 2) + g.SizeOf("\"SCRIPT_FILENAME\"") - 1
						}), nil)
					} else {
						script_path_translated = _sapiCgiPutenv("SCRIPT_FILENAME", g.SizeOf("\"SCRIPT_FILENAME\"")-1, nil)
					}
					core.sapi_globals.sapi_headers.http_response_code = 404
				}
				if core.sapi_globals.request_info.request_uri == nil {
					if orig_script_name == nil || strcmp(orig_script_name, env_script_name) != 0 {
						if orig_script_name != nil {
							g.CondF(has_env != 0, func() *byte {
								return core.FcgiQuickPutenv(request, "ORIG_SCRIPT_NAME", g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1, g.CondF(g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"ORIG_SCRIPT_NAME\"") - 1) }, func() int {
									return (uint("ORIG_SCRIPT_NAME"[3]) << 2) + (uint("ORIG_SCRIPT_NAME"[g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1-2]) << 4) + (uint("ORIG_SCRIPT_NAME"[g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1-1]) << 2) + g.SizeOf("\"ORIG_SCRIPT_NAME\"") - 1
								}), orig_script_name)
							}, func() *byte {
								return _sapiCgiPutenv("ORIG_SCRIPT_NAME", g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1, orig_script_name)
							})
						}
						if has_env != 0 {
							core.sapi_globals.request_info.request_uri = core.FcgiQuickPutenv(request, "SCRIPT_NAME", g.SizeOf("\"SCRIPT_NAME\"")-1, g.CondF(g.SizeOf("\"SCRIPT_NAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"SCRIPT_NAME\"") - 1) }, func() int {
								return (uint("SCRIPT_NAME"[3]) << 2) + (uint("SCRIPT_NAME"[g.SizeOf("\"SCRIPT_NAME\"")-1-2]) << 4) + (uint("SCRIPT_NAME"[g.SizeOf("\"SCRIPT_NAME\"")-1-1]) << 2) + g.SizeOf("\"SCRIPT_NAME\"") - 1
							}), env_script_name)
						} else {
							core.sapi_globals.request_info.request_uri = _sapiCgiPutenv("SCRIPT_NAME", g.SizeOf("\"SCRIPT_NAME\"")-1, env_script_name)
						}
					} else {
						core.sapi_globals.request_info.request_uri = orig_script_name
					}
				}
				if pt != nil {
					zend._efree(pt)
				}
			} else {

				/* make sure path_info/translated are empty */

				if orig_script_filename == nil || script_path_translated != orig_script_filename && strcmp(script_path_translated, orig_script_filename) != 0 {
					if orig_script_filename != nil {
						g.CondF(has_env != 0, func() *byte {
							return core.FcgiQuickPutenv(request, "ORIG_SCRIPT_FILENAME", g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1, g.CondF(g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"ORIG_SCRIPT_FILENAME\"") - 1) }, func() int {
								return (uint("ORIG_SCRIPT_FILENAME"[3]) << 2) + (uint("ORIG_SCRIPT_FILENAME"[g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1-2]) << 4) + (uint("ORIG_SCRIPT_FILENAME"[g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1-1]) << 2) + g.SizeOf("\"ORIG_SCRIPT_FILENAME\"") - 1
							}), orig_script_filename)
						}, func() *byte {
							return _sapiCgiPutenv("ORIG_SCRIPT_FILENAME", g.SizeOf("\"ORIG_SCRIPT_FILENAME\"")-1, orig_script_filename)
						})
					}
					if has_env != 0 {
						script_path_translated = core.FcgiQuickPutenv(request, "SCRIPT_FILENAME", g.SizeOf("\"SCRIPT_FILENAME\"")-1, g.CondF(g.SizeOf("\"SCRIPT_FILENAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"SCRIPT_FILENAME\"") - 1) }, func() int {
							return (uint("SCRIPT_FILENAME"[3]) << 2) + (uint("SCRIPT_FILENAME"[g.SizeOf("\"SCRIPT_FILENAME\"")-1-2]) << 4) + (uint("SCRIPT_FILENAME"[g.SizeOf("\"SCRIPT_FILENAME\"")-1-1]) << 2) + g.SizeOf("\"SCRIPT_FILENAME\"") - 1
						}), script_path_translated)
					} else {
						script_path_translated = _sapiCgiPutenv("SCRIPT_FILENAME", g.SizeOf("\"SCRIPT_FILENAME\"")-1, script_path_translated)
					}
				}
				if env_redirect_url != nil {
					if orig_path_info != nil {
						g.CondF(has_env != 0, func() *byte {
							return core.FcgiQuickPutenv(request, "ORIG_PATH_INFO", g.SizeOf("\"ORIG_PATH_INFO\"")-1, g.CondF(g.SizeOf("\"ORIG_PATH_INFO\"")-1 < 3, func() uint { return uint(g.SizeOf("\"ORIG_PATH_INFO\"") - 1) }, func() int {
								return (uint("ORIG_PATH_INFO"[3]) << 2) + (uint("ORIG_PATH_INFO"[g.SizeOf("\"ORIG_PATH_INFO\"")-1-2]) << 4) + (uint("ORIG_PATH_INFO"[g.SizeOf("\"ORIG_PATH_INFO\"")-1-1]) << 2) + g.SizeOf("\"ORIG_PATH_INFO\"") - 1
							}), orig_path_info)
						}, func() *byte {
							return _sapiCgiPutenv("ORIG_PATH_INFO", g.SizeOf("\"ORIG_PATH_INFO\"")-1, orig_path_info)
						})
						g.CondF(has_env != 0, func() *byte {
							return core.FcgiQuickPutenv(request, "PATH_INFO", g.SizeOf("\"PATH_INFO\"")-1, g.CondF(g.SizeOf("\"PATH_INFO\"")-1 < 3, func() uint { return uint(g.SizeOf("\"PATH_INFO\"") - 1) }, func() int {
								return (uint("PATH_INFO"[3]) << 2) + (uint("PATH_INFO"[g.SizeOf("\"PATH_INFO\"")-1-2]) << 4) + (uint("PATH_INFO"[g.SizeOf("\"PATH_INFO\"")-1-1]) << 2) + g.SizeOf("\"PATH_INFO\"") - 1
							}), nil)
						}, func() *byte { return _sapiCgiPutenv("PATH_INFO", g.SizeOf("\"PATH_INFO\"")-1, nil) })
					}
					if orig_path_translated != nil {
						g.CondF(has_env != 0, func() *byte {
							return core.FcgiQuickPutenv(request, "ORIG_PATH_TRANSLATED", g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1, g.CondF(g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1 < 3, func() uint { return uint(g.SizeOf("\"ORIG_PATH_TRANSLATED\"") - 1) }, func() int {
								return (uint("ORIG_PATH_TRANSLATED"[3]) << 2) + (uint("ORIG_PATH_TRANSLATED"[g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1-2]) << 4) + (uint("ORIG_PATH_TRANSLATED"[g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1-1]) << 2) + g.SizeOf("\"ORIG_PATH_TRANSLATED\"") - 1
							}), orig_path_translated)
						}, func() *byte {
							return _sapiCgiPutenv("ORIG_PATH_TRANSLATED", g.SizeOf("\"ORIG_PATH_TRANSLATED\"")-1, orig_path_translated)
						})
						g.CondF(has_env != 0, func() *byte {
							return core.FcgiQuickPutenv(request, "PATH_TRANSLATED", g.SizeOf("\"PATH_TRANSLATED\"")-1, g.CondF(g.SizeOf("\"PATH_TRANSLATED\"")-1 < 3, func() uint { return uint(g.SizeOf("\"PATH_TRANSLATED\"") - 1) }, func() int {
								return (uint("PATH_TRANSLATED"[3]) << 2) + (uint("PATH_TRANSLATED"[g.SizeOf("\"PATH_TRANSLATED\"")-1-2]) << 4) + (uint("PATH_TRANSLATED"[g.SizeOf("\"PATH_TRANSLATED\"")-1-1]) << 2) + g.SizeOf("\"PATH_TRANSLATED\"") - 1
							}), nil)
						}, func() *byte { return _sapiCgiPutenv("PATH_TRANSLATED", g.SizeOf("\"PATH_TRANSLATED\"")-1, nil) })
					}
				}
				if env_script_name != orig_script_name {
					if orig_script_name != nil {
						g.CondF(has_env != 0, func() *byte {
							return core.FcgiQuickPutenv(request, "ORIG_SCRIPT_NAME", g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1, g.CondF(g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"ORIG_SCRIPT_NAME\"") - 1) }, func() int {
								return (uint("ORIG_SCRIPT_NAME"[3]) << 2) + (uint("ORIG_SCRIPT_NAME"[g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1-2]) << 4) + (uint("ORIG_SCRIPT_NAME"[g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1-1]) << 2) + g.SizeOf("\"ORIG_SCRIPT_NAME\"") - 1
							}), orig_script_name)
						}, func() *byte {
							return _sapiCgiPutenv("ORIG_SCRIPT_NAME", g.SizeOf("\"ORIG_SCRIPT_NAME\"")-1, orig_script_name)
						})
					}
					if has_env != 0 {
						core.sapi_globals.request_info.request_uri = core.FcgiQuickPutenv(request, "SCRIPT_NAME", g.SizeOf("\"SCRIPT_NAME\"")-1, g.CondF(g.SizeOf("\"SCRIPT_NAME\"")-1 < 3, func() uint { return uint(g.SizeOf("\"SCRIPT_NAME\"") - 1) }, func() int {
							return (uint("SCRIPT_NAME"[3]) << 2) + (uint("SCRIPT_NAME"[g.SizeOf("\"SCRIPT_NAME\"")-1-2]) << 4) + (uint("SCRIPT_NAME"[g.SizeOf("\"SCRIPT_NAME\"")-1-1]) << 2) + g.SizeOf("\"SCRIPT_NAME\"") - 1
						}), env_script_name)
					} else {
						core.sapi_globals.request_info.request_uri = _sapiCgiPutenv("SCRIPT_NAME", g.SizeOf("\"SCRIPT_NAME\"")-1, env_script_name)
					}
				} else {
					core.sapi_globals.request_info.request_uri = env_script_name
				}
				zend._efree(real_path)
			}

			/*
			 * if the file doesn't exist, try to extract PATH_INFO out
			 * of it by stat'ing back through the '/'
			 * this fixes url's like /info.php/test
			 */

		} else {

			/* pre 4.3 behaviour, shouldn't be used but provides BC */

			if env_path_info != nil {
				core.sapi_globals.request_info.request_uri = env_path_info
			} else {
				core.sapi_globals.request_info.request_uri = env_script_name
			}
			if php_cgi_globals.GetDiscardPath() == 0 && env_path_translated != nil {
				script_path_translated = env_path_translated
			}
		}
		if IsValidPath(script_path_translated) != 0 {
			core.sapi_globals.request_info.path_translated = zend._estrdup(script_path_translated)
		}
		if has_env != 0 {
			core.sapi_globals.request_info.request_method = core.FcgiQuickGetenv(request, "REQUEST_METHOD", g.SizeOf("\"REQUEST_METHOD\"")-1, g.CondF(g.SizeOf("\"REQUEST_METHOD\"")-1 < 3, func() uint { return uint(g.SizeOf("\"REQUEST_METHOD\"") - 1) }, func() int {
				return (uint("REQUEST_METHOD"[3]) << 2) + (uint("REQUEST_METHOD"[g.SizeOf("\"REQUEST_METHOD\"")-1-2]) << 4) + (uint("REQUEST_METHOD"[g.SizeOf("\"REQUEST_METHOD\"")-1-1]) << 2) + g.SizeOf("\"REQUEST_METHOD\"") - 1
			}))
		} else {
			core.sapi_globals.request_info.request_method = getenv("REQUEST_METHOD")
		}

		/* FIXME - Work out proto_num here */

		if has_env != 0 {
			core.sapi_globals.request_info.query_string = core.FcgiQuickGetenv(request, "QUERY_STRING", g.SizeOf("\"QUERY_STRING\"")-1, g.CondF(g.SizeOf("\"QUERY_STRING\"")-1 < 3, func() uint { return uint(g.SizeOf("\"QUERY_STRING\"") - 1) }, func() int {
				return (uint("QUERY_STRING"[3]) << 2) + (uint("QUERY_STRING"[g.SizeOf("\"QUERY_STRING\"")-1-2]) << 4) + (uint("QUERY_STRING"[g.SizeOf("\"QUERY_STRING\"")-1-1]) << 2) + g.SizeOf("\"QUERY_STRING\"") - 1
			}))
		} else {
			core.sapi_globals.request_info.query_string = getenv("QUERY_STRING")
		}
		if content_type != nil {
			core.sapi_globals.request_info.content_type = content_type
		} else {
			core.sapi_globals.request_info.content_type = ""
		}
		if content_length != nil {
			core.sapi_globals.request_info.content_length = atol(content_length)
		} else {
			core.sapi_globals.request_info.content_length = 0
		}

		/* The CGI RFC allows servers to pass on unvalidated Authorization data */

		if has_env != 0 {
			auth = core.FcgiQuickGetenv(request, "HTTP_AUTHORIZATION", g.SizeOf("\"HTTP_AUTHORIZATION\"")-1, g.CondF(g.SizeOf("\"HTTP_AUTHORIZATION\"")-1 < 3, func() uint { return uint(g.SizeOf("\"HTTP_AUTHORIZATION\"") - 1) }, func() int {
				return (uint("HTTP_AUTHORIZATION"[3]) << 2) + (uint("HTTP_AUTHORIZATION"[g.SizeOf("\"HTTP_AUTHORIZATION\"")-1-2]) << 4) + (uint("HTTP_AUTHORIZATION"[g.SizeOf("\"HTTP_AUTHORIZATION\"")-1-1]) << 2) + g.SizeOf("\"HTTP_AUTHORIZATION\"") - 1
			}))
		} else {
			auth = getenv("HTTP_AUTHORIZATION")
		}
		core.PhpHandleAuthData(auth)
	}

	/* script_path_translated being set is a good indication that
	 * we are running in a cgi environment, since it is always
	 * null otherwise.  otherwise, the filename
	 * of the script will be retrieved later via argc/argv */
}

/* }}} */

/**
 * Clean up child processes upon exit
 */

func FastcgiCleanup(signal int) {
	sigaction(SIGTERM, &OldTerm, 0)

	/* Kill all the processes in our process group */

	kill(-Pgroup, SIGTERM)
	if Parent != 0 && ParentWaiting != 0 {
		ExitSignal = 1
	} else {
		exit(0)
	}
}

var IniEntries []zend.ZendIniEntryDef = []zend.ZendIniEntryDef{
	{
		"cgi.rfc2616_headers",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetRfc2616Headers())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		"0",
		nil,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"cgi.rfc2616_headers\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"cgi.nph",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetNph())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		"0",
		nil,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"cgi.nph\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"cgi.check_shebang_line",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetCheckShebangLine())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		"1",
		nil,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"cgi.check_shebang_line\"") - 1,
		1 << 2,
	},
	{
		"cgi.force_redirect",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetForceRedirect())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		"1",
		nil,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"cgi.force_redirect\"") - 1,
		1 << 2,
	},
	{
		"cgi.redirect_status_env",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetRedirectStatusEnv())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		nil,
		nil,
		g.SizeOf("NULL") - 1,
		g.SizeOf("\"cgi.redirect_status_env\"") - 1,
		1 << 2,
	},
	{
		"cgi.fix_pathinfo",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetFixPathinfo())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		"1",
		nil,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"cgi.fix_pathinfo\"") - 1,
		1 << 2,
	},
	{
		"cgi.discard_path",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetDiscardPath())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		"0",
		nil,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"cgi.discard_path\"") - 1,
		1 << 2,
	},
	{
		"fastcgi.logging",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetFcgiLogging())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		"1",
		nil,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"fastcgi.logging\"") - 1,
		1 << 2,
	},
	{nil, nil, nil, nil, nil, nil, nil, 0, 0, 0},
}

/* {{{ php_cgi_globals_ctor
 */

func PhpCgiGlobalsCtor(php_cgi_globals *php_cgi_globals_struct) {
	php_cgi_globals.SetRfc2616Headers(0)
	php_cgi_globals.SetNph(0)
	php_cgi_globals.SetCheckShebangLine(1)
	php_cgi_globals.SetForceRedirect(1)
	php_cgi_globals.SetRedirectStatusEnv(nil)
	php_cgi_globals.SetFixPathinfo(1)
	php_cgi_globals.SetDiscardPath(0)
	php_cgi_globals.SetFcgiLogging(1)
	zend._zendHashInit(&php_cgi_globals.user_config_cache, 8, UserConfigCacheEntryDtor, 1)
}

/* }}} */

func ZmStartupCgi(type_ int, module_number int) int {
	zend.ZendRegisterIniEntries(IniEntries, module_number)
	return zend.SUCCESS
}

/* }}} */

func ZmShutdownCgi(type_ int, module_number int) int {
	zend.ZendHashDestroy(&(php_cgi_globals.GetUserConfigCache()))
	zend.ZendUnregisterIniEntries(module_number)
	return zend.SUCCESS
}

/* }}} */

func ZmInfoCgi(zend_module *zend.ZendModuleEntry) { core.DisplayIniEntries(zend_module) }

/* }}} */

func ZifApacheChildTerminate(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) {
		return
	}
	if core.FcgiIsFastcgi() != 0 {
		core.FcgiTerminate()
	}
}

/* }}} */

func ZifApacheRequestHeaders(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) {
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if core.FcgiIsFastcgi() != 0 {
		var request *core.FcgiRequest = (*core.FcgiRequest)(core.sapi_globals.server_context)
		core.FcgiLoadenv(request, core.SapiAddRequestHeader, return_value)
	} else {
		var buf []byte
		var env **byte
		var p **byte
		var q **byte
		var var_ **byte
		var val **byte
		var t **byte = buf
		var alloc_size int = g.SizeOf("buf")
		var var_len zend.ZendUlong
		for env = cli.Environ; env != nil && (*env) != nil; env++ {
			val = strchr(*env, '=')
			if val == nil {
				continue
			}
			var_len = val - (*env)
			if var_len >= alloc_size {
				alloc_size = var_len + 64
				if t == buf {
					t = zend._emalloc(alloc_size)
				} else {
					t = zend._erealloc(t, alloc_size)
				}
			}
			var_ = *env
			if var_len > 5 && var_[0] == 'H' && var_[1] == 'T' && var_[2] == 'T' && var_[3] == 'P' && var_[4] == '_' {
				var_len -= 5
				if var_len >= alloc_size {
					alloc_size = var_len + 64
					if t == buf {
						t = zend._emalloc(alloc_size)
					} else {
						t = zend._erealloc(t, alloc_size)
					}
				}
				p = var_ + 5
				q = t
				var_ = q

				/* First char keep uppercase */

				*p++
				g.PostInc(&(*q)) = (*p) - 1
				for (*p) != nil {
					if (*p) == '=' {

						/* End of name */

						break

						/* End of name */

					} else if (*p) == '_' {
						g.PostInc(&(*q)) = '-'
						p++

						/* First char after - keep uppercase */

						if (*p) != nil && (*p) != '=' {
							*p++
							g.PostInc(&(*q)) = (*p) - 1
						}

						/* First char after - keep uppercase */

					} else if (*p) >= 'A' && (*p) <= 'Z' {

						/* lowercase */

						g.PostInc(&(*q)) = g.PostInc(&(*p)) - 'A' + 'a'

						/* lowercase */

					} else {
						*p++
						g.PostInc(&(*q)) = (*p) - 1
					}
				}
				*q = 0
			} else if var_len == g.SizeOf("\"CONTENT_TYPE\"")-1 && memcmp(var_, "CONTENT_TYPE", g.SizeOf("\"CONTENT_TYPE\"")-1) == 0 {
				var_ = "Content-Type"
			} else if var_len == g.SizeOf("\"CONTENT_LENGTH\"")-1 && memcmp(var_, "CONTENT_LENGTH", g.SizeOf("\"CONTENT_LENGTH\"")-1) == 0 {
				var_ = "Content-Length"
			} else {
				continue
			}
			val++
			zend.AddAssocStringEx(return_value, var_, var_len, val)
		}
		if t != buf && t != nil {
			zend._efree(t)
		}
	}
}

/* }}} */

func AddResponseHeader(h *core.SapiHeader, return_value *zend.Zval) {
	if h.header_len > 0 {
		var s *byte
		var len_ int = 0
		var p *byte = strchr(h.header, ':')
		if nil != p {
			len_ = p - h.header
		}
		if len_ > 0 {
			for len_ != 0 && (h.header[len_-1] == ' ' || h.header[len_-1] == '\t') {
				len_--
			}
			if len_ != 0 {
				s = zend._emalloc(len_ + 1)
				memcpy(s, h.header, len_)
				s[len_] = 0
				for {
					p++
					if !((*p) == ' ' || (*p) == '\t') {
						break
					}
				}
				zend.AddAssocStringlEx(return_value, s, len_, p, h.header_len-(p-h.header))
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

var ArginfoNoArgs []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var CgiFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"apache_child_terminate",
		ZifApacheChildTerminate,
		ArginfoNoArgs,
		uint32(g.SizeOf("arginfo_no_args")/g.SizeOf("struct _zend_internal_arg_info") - 1),
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
var CgiModuleEntry zend.ZendModuleEntry = zend.ZendModuleEntry{g.SizeOf("zend_module_entry"), 20190902, 0, 0, nil, nil, "cgi-fcgi", CgiFunctions, ZmStartupCgi, ZmShutdownCgi, nil, nil, ZmInfoCgi, "7.4.33", 0, nil, nil, nil, nil, 0, 0, nil, 0, "API" + "20190902" + ",NTS"}

/* {{{ main
 */

func Main(argc int, argv []*byte) int {
	var free_query_string int = 0
	var exit_status int = zend.SUCCESS
	var cgi int = 0
	var c int
	var i int
	var len_ int
	var file_handle zend.ZendFileHandle
	var s *byte

	/* temporary locals */

	var behavior int = 1
	var no_headers int = 0
	var orig_optind int = PhpOptind
	var orig_optarg *byte = PhpOptarg
	var script_file *byte = nil
	var ini_entries_len int = 0

	/* end of temporary locals */

	var max_requests int = 500
	var requests int = 0
	var fastcgi int
	var bindpath *byte = nil
	var fcgi_fd int = 0
	var request *core.FcgiRequest = nil
	var warmup_repeats int = 0
	var repeats int = 1
	var benchmark int = 0
	var start __struct__timeval
	var end __struct__timeval
	var status int = 0
	var query_string *byte
	var decoded_query_string *byte
	var skip_getopt int = 0
	zend.ZendSignalStartup()
	PhpCgiGlobalsCtor(&php_cgi_globals)
	core.SapiStartup(&CgiSapiModule)
	fastcgi = core.FcgiIsFastcgi()
	CgiSapiModule.php_ini_path_override = nil
	if fastcgi == 0 {

		/* Make sure we detect we are a cgi - a bit redundancy here,
		 * but the default case is that we have to check only the first one. */

		if getenv("SERVER_SOFTWARE") || getenv("SERVER_NAME") || getenv("GATEWAY_INTERFACE") || getenv("REQUEST_METHOD") {
			cgi = 1
		}

		/* Make sure we detect we are a cgi - a bit redundancy here,
		 * but the default case is that we have to check only the first one. */

	}
	if g.Assign(&query_string, getenv("QUERY_STRING")) != nil && strchr(query_string, '=') == nil {

		/* we've got query string that has no = - apache CGI will pass it to command line */

		var p *uint8
		decoded_query_string = strdup(query_string)
		streams.PhpUrlDecode(decoded_query_string, strlen(decoded_query_string))
		for p = (*uint8)(decoded_query_string); (*p) != 0 && (*p) <= ' '; p++ {

		}
		if (*p) == '-' {
			skip_getopt = 1
		}
		zend.Free(decoded_query_string)
	}
	for skip_getopt == 0 && g.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &PhpOptarg, &PhpOptind, 0, 2)) != -1 {
		switch c {
		case 'c':
			if CgiSapiModule.php_ini_path_override != nil {
				zend.Free(CgiSapiModule.php_ini_path_override)
			}
			CgiSapiModule.php_ini_path_override = strdup(PhpOptarg)
			break
		case 'n':
			CgiSapiModule.php_ini_ignore = 1
			break
		case 'd':

			/* define ini __special__  entries on command line */

			var len_ int = strlen(PhpOptarg)
			var val *byte
			if g.Assign(&val, strchr(PhpOptarg, '=')) {
				val++
				if !(isalnum(*val)) && (*val) != '"' && (*val) != '\'' && (*val) != '0' {
					CgiSapiModule.ini_entries = realloc(CgiSapiModule.ini_entries, ini_entries_len+len_+g.SizeOf("\"\\\"\\\"\\n\\0\""))
					memcpy(CgiSapiModule.ini_entries+ini_entries_len, PhpOptarg, val-PhpOptarg)
					ini_entries_len += val - PhpOptarg
					memcpy(CgiSapiModule.ini_entries+ini_entries_len, "\"", 1)
					ini_entries_len++
					memcpy(CgiSapiModule.ini_entries+ini_entries_len, val, len_-(val-PhpOptarg))
					ini_entries_len += len_ - (val - PhpOptarg)
					memcpy(CgiSapiModule.ini_entries+ini_entries_len, "\"\n0", g.SizeOf("\"\\\"\\n\\0\""))
					ini_entries_len += g.SizeOf("\"\\n\\0\\\"\"") - 2
				} else {
					CgiSapiModule.ini_entries = realloc(CgiSapiModule.ini_entries, ini_entries_len+len_+g.SizeOf("\"\\n\\0\""))
					memcpy(CgiSapiModule.ini_entries+ini_entries_len, PhpOptarg, len_)
					memcpy(CgiSapiModule.ini_entries+ini_entries_len+len_, "\n0", g.SizeOf("\"\\n\\0\""))
					ini_entries_len += len_ + g.SizeOf("\"\\n\\0\"") - 2
				}
			} else {
				CgiSapiModule.ini_entries = realloc(CgiSapiModule.ini_entries, ini_entries_len+len_+g.SizeOf("\"=1\\n\\0\""))
				memcpy(CgiSapiModule.ini_entries+ini_entries_len, PhpOptarg, len_)
				memcpy(CgiSapiModule.ini_entries+ini_entries_len+len_, "=1\n0", g.SizeOf("\"=1\\n\\0\""))
				ini_entries_len += len_ + g.SizeOf("\"=1\\n\\0\"") - 2
			}
			break
		case 'b':
			if fastcgi == 0 {
				bindpath = strdup(PhpOptarg)
			}
			break
		case 's':
			behavior = 2
			break
		}
	}
	PhpOptind = orig_optind
	PhpOptarg = orig_optarg
	if fastcgi != 0 || bindpath != nil {

		/* Override SAPI callbacks */

		CgiSapiModule.ub_write = SapiFcgiUbWrite
		CgiSapiModule.flush = SapiFcgiFlush
		CgiSapiModule.read_post = SapiFcgiReadPost
		CgiSapiModule.getenv = SapiFcgiGetenv
		CgiSapiModule.read_cookies = SapiFcgiReadCookies
	}
	CgiSapiModule.executable_location = argv[0]
	if cgi == 0 && fastcgi == 0 && bindpath == nil {
		CgiSapiModule.additional_functions = AdditionalFunctions
	}

	/* startup after we get the above ini override se we get things right */

	if CgiSapiModule.startup(&CgiSapiModule) == zend.FAILURE {
		zend.Free(bindpath)
		return zend.FAILURE
	}

	/* check force_cgi after startup, so we have proper output */

	if cgi != 0 && php_cgi_globals.GetForceRedirect() != 0 {

		/* Apache will generate REDIRECT_STATUS,
		 * Netscape and redirect.so will generate HTTP_REDIRECT_STATUS.
		 * redirect.so and installation instructions available from
		 * http://www.koehntopp.de/php.
		 *   -- kk@netuse.de
		 */

		if !(getenv("REDIRECT_STATUS")) && !(getenv("HTTP_REDIRECT_STATUS")) && (php_cgi_globals.GetRedirectStatusEnv() == nil || !(getenv(php_cgi_globals.GetRedirectStatusEnv()))) {
			var __orig_bailout *sigjmp_buf = zend.EG.bailout
			var __bailout sigjmp_buf
			zend.EG.bailout = &__bailout
			if sigsetjmp(__bailout, 0) == 0 {
				core.sapi_globals.sapi_headers.http_response_code = 400
				var __str *byte = "<b>Security Alert!</b> The PHP CGI cannot be accessed directly.\n\n\n<p>This PHP CGI binary was compiled with force-cgi-redirect enabled.  This\n\nmeans that a page will only be served up if the REDIRECT_STATUS CGI variable is\n\nset, e.g. via an Apache Action directive.</p>\n\n<p>For more information as to <i>why</i> this behaviour exists, see the <a href=\"http://php.net/security.cgi-bin\">\nmanual page for CGI security</a>.</p>\n\n<p>For more information about changing this behaviour or re-enabling this webserver,\n\nconsult the installation file that came with this distribution, or visit \n\n<a href=\"http://php.net/install.windows\">the manual page</a>.</p>\n"
				core.PhpOutputWrite(__str, strlen(__str))
			} else {
				zend.EG.bailout = __orig_bailout
			}
			zend.EG.bailout = __orig_bailout
			zend.Free(bindpath)
			return zend.FAILURE
		}

		/* Apache will generate REDIRECT_STATUS,
		 * Netscape and redirect.so will generate HTTP_REDIRECT_STATUS.
		 * redirect.so and installation instructions available from
		 * http://www.koehntopp.de/php.
		 *   -- kk@netuse.de
		 */

	}
	core.FcgiSetLogger(FcgiLog)
	if bindpath != nil {
		var backlog int = 128
		if getenv("PHP_FCGI_BACKLOG") {
			backlog = atoi(getenv("PHP_FCGI_BACKLOG"))
		}
		fcgi_fd = core.FcgiListen(bindpath, backlog)
		if fcgi_fd < 0 {
			r.Fprintf(stderr, "Couldn't create FastCGI listen socket on port %s\n", bindpath)
			return zend.FAILURE
		}
		fastcgi = core.FcgiIsFastcgi()
	}

	/* make php call us to get _ENV vars */

	PhpPhpImportEnvironmentVariables = core.PhpImportEnvironmentVariables
	core.PhpImportEnvironmentVariables = CgiPhpImportEnvironmentVariables
	if fastcgi != 0 {

		/* How many times to run PHP scripts before dying */

		if getenv("PHP_FCGI_MAX_REQUESTS") {
			max_requests = atoi(getenv("PHP_FCGI_MAX_REQUESTS"))
			if max_requests < 0 {
				r.Fprintf(stderr, "PHP_FCGI_MAX_REQUESTS is not valid\n")
				return zend.FAILURE
			}
		}

		/* library is already initialized, now init our request */

		request = core.FcgiInitRequest(fcgi_fd, nil, nil, nil)

		/* Pre-fork or spawn, if required */

		if getenv("PHP_FCGI_CHILDREN") {
			var children_str *byte = getenv("PHP_FCGI_CHILDREN")
			Children = atoi(children_str)
			if Children < 0 {
				r.Fprintf(stderr, "PHP_FCGI_CHILDREN is not valid\n")
				return zend.FAILURE
			}
			core.FcgiSetMgmtVar("FCGI_MAX_CONNS", g.SizeOf("\"FCGI_MAX_CONNS\"")-1, children_str, strlen(children_str))

			/* This is the number of concurrent requests, equals FCGI_MAX_CONNS */

			core.FcgiSetMgmtVar("FCGI_MAX_REQS", g.SizeOf("\"FCGI_MAX_REQS\"")-1, children_str, strlen(children_str))

			/* This is the number of concurrent requests, equals FCGI_MAX_CONNS */

		} else {
			core.FcgiSetMgmtVar("FCGI_MAX_CONNS", g.SizeOf("\"FCGI_MAX_CONNS\"")-1, "1", g.SizeOf("\"1\"")-1)
			core.FcgiSetMgmtVar("FCGI_MAX_REQS", g.SizeOf("\"FCGI_MAX_REQS\"")-1, "1", g.SizeOf("\"1\"")-1)
		}
		if Children != 0 {
			var running int = 0
			var pid pid_t

			/* Create a process group for ourself & children */

			setsid()
			Pgroup = getpgrp()

			/* Set up handler to kill children upon exit */

			Act.sa_flags = 0
			Act.sa_handler = FastcgiCleanup
			if sigaction(SIGTERM, &Act, &OldTerm) || sigaction(SIGINT, &Act, &OldInt) || sigaction(SIGQUIT, &Act, &OldQuit) {
				r.Perror("Can't set signals")
				exit(1)
			}
			if core.FcgiInShutdown() != 0 {
				goto parent_out
			}
			for Parent != 0 {
				for {
					pid = fork()
					switch pid {
					case 0:

						/* One of the children.
						 * Make sure we don't go round the
						 * fork loop any more
						 */

						Parent = 0

						/* don't catch our signals */

						sigaction(SIGTERM, &OldTerm, 0)
						sigaction(SIGQUIT, &OldQuit, 0)
						sigaction(SIGINT, &OldInt, 0)
						zend.ZendSignalInit()
						break
					case -1:
						r.Perror("php (pre-forking)")
						exit(1)
						break
					default:

						/* Fine */

						running++
						break
					}
					if !(Parent != 0 && running < Children) {
						break
					}
				}
				if Parent != 0 {
					ParentWaiting = 1
					for true {
						if wait(&status) >= 0 {
							running--
							break
						} else if ExitSignal != 0 {
							break
						}
					}
					if ExitSignal != 0 {
						goto parent_out
					}
				}
			}
		} else {
			Parent = 0
			zend.ZendSignalInit()
		}
	}
	zend.EG.bailout = nil
	var __orig_bailout *sigjmp_buf = zend.EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		for skip_getopt == 0 && g.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &PhpOptarg, &PhpOptind, 1, 2)) != -1 {
			switch c {
			case 'T':
				benchmark = 1
				var comma *byte = strchr(PhpOptarg, ',')
				if comma != nil {
					warmup_repeats = atoi(PhpOptarg)
					repeats = atoi(comma + 1)
				} else {
					repeats = atoi(PhpOptarg)
				}
				gettimeofday(&start, nil)
				break
			case 'h':

			case '?':

			case -2:
				if request != nil {
					core.FcgiDestroyRequest(request)
				}
				core.FcgiShutdown()
				no_headers = 1
				core.sapi_globals.headers_sent = 1
				PhpCgiUsage(argv[0])
				core.PhpOutputEndAll()
				exit_status = 0
				if c == -2 {
					exit_status = 1
				}
				goto out
			}
		}
		PhpOptind = orig_optind
		PhpOptarg = orig_optarg

		/* start of FAST CGI loop */

		for fastcgi == 0 || core.FcgiAcceptRequest(request) >= 0 {
			if fastcgi != 0 {
				core.sapi_globals.server_context = any(request)
			} else {
				core.sapi_globals.server_context = any(1)
			}
			InitRequestInfo(request)
			if cgi == 0 && fastcgi == 0 {
				for g.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &PhpOptarg, &PhpOptind, 0, 2)) != -1 {
					switch c {
					case 'a':
						r.Printf("Interactive mode enabled\n\n")
						break
					case 'C':
						core.sapi_globals.options |= 1
						break
					case 'e':
						zend.CG.compiler_options |= 1<<0 | 1<<1
						break
					case 'f':
						if script_file != nil {
							zend._efree(script_file)
						}
						script_file = zend._estrdup(PhpOptarg)
						no_headers = 1
						break
					case 'i':
						if script_file != nil {
							zend._efree(script_file)
						}
						if core.PhpRequestStartup() == zend.FAILURE {
							core.sapi_globals.server_context = nil
							core.PhpModuleShutdown()
							zend.Free(bindpath)
							return zend.FAILURE
						}
						if no_headers != 0 {
							core.sapi_globals.headers_sent = 1
							core.sapi_globals.request_info.no_headers = 1
						}
						standard.PhpPrintInfo(0xffffffff)
						core.PhpRequestShutdown(any(0))
						core.FcgiShutdown()
						exit_status = 0
						goto out
					case 'l':
						no_headers = 1
						behavior = 4
						break
					case 'm':
						if script_file != nil {
							zend._efree(script_file)
						}
						core.sapi_globals.headers_sent = 1
						core.PhpPrintf("[PHP Modules]\n")
						PrintModules()
						core.PhpPrintf("\n[Zend Modules]\n")
						PrintExtensions()
						core.PhpPrintf("\n")
						core.PhpOutputEndAll()
						core.FcgiShutdown()
						exit_status = 0
						goto out
					case 'q':
						no_headers = 1
						break
					case 'v':
						if script_file != nil {
							zend._efree(script_file)
						}
						no_headers = 1
						if core.PhpRequestStartup() == zend.FAILURE {
							core.sapi_globals.server_context = nil
							core.PhpModuleShutdown()
							zend.Free(bindpath)
							return zend.FAILURE
						}
						core.sapi_globals.headers_sent = 1
						core.sapi_globals.request_info.no_headers = 1
						core.PhpPrintf("PHP %s (%s) (built: %s %s)\nCopyright (c) The PHP Group\n%s", "7.4.33", core.sapi_module.name, __DATE__, __TIME__, zend.GetZendVersion())
						core.PhpRequestShutdown(any(0))
						core.FcgiShutdown()
						exit_status = 0
						goto out
					case 'w':
						behavior = 5
						break
					case 'z':
						zend.ZendLoadExtension(PhpOptarg)
						break
					default:
						break
					}
				}
				if script_file != nil {

					/* override path_translated if -f on command line */

					if core.sapi_globals.request_info.path_translated != nil {
						zend._efree(core.sapi_globals.request_info.path_translated)
					}
					core.sapi_globals.request_info.path_translated = script_file

					/* before registering argv to module exchange the *new* argv[0] */

					core.sapi_globals.request_info.argc = argc - (PhpOptind - 1)
					core.sapi_globals.request_info.argv = &argv[PhpOptind-1]
					core.sapi_globals.request_info.argv[0] = script_file
				} else if argc > PhpOptind {

					/* file is on command line, but not in -f opt */

					if core.sapi_globals.request_info.path_translated != nil {
						zend._efree(core.sapi_globals.request_info.path_translated)
					}
					core.sapi_globals.request_info.path_translated = zend._estrdup(argv[PhpOptind])

					/* arguments after the file are considered script args */

					core.sapi_globals.request_info.argc = argc - PhpOptind
					core.sapi_globals.request_info.argv = &argv[PhpOptind]
				}
				if no_headers != 0 {
					core.sapi_globals.headers_sent = 1
					core.sapi_globals.request_info.no_headers = 1
				}

				/* all remaining arguments are part of the query string
				 * this section of code concatenates all remaining arguments
				 * into a single string, separating args with a &
				 * this allows command lines like:
				 *
				 *  test.php v1=test v2=hello+world!
				 *  test.php "v1=test&v2=hello world!"
				 *  test.php v1=test "v2=hello world!"
				 */

				if core.sapi_globals.request_info.query_string == nil && argc > PhpOptind {
					var slen int = strlen(core.CoreGlobals.arg_separator.input)
					len_ = 0
					for i = PhpOptind; i < argc; i++ {
						if i < argc-1 {
							len_ += strlen(argv[i]) + slen
						} else {
							len_ += strlen(argv[i])
						}
					}
					len_ += 2
					s = zend.Malloc(len_)
					*s = '0'
					for i = PhpOptind; i < argc; i++ {
						strlcat(s, argv[i], len_)
						if i < argc-1 {
							strlcat(s, core.CoreGlobals.arg_separator.input, len_)
						}
					}
					core.sapi_globals.request_info.query_string = s
					free_query_string = 1
				}

				/* all remaining arguments are part of the query string
				 * this section of code concatenates all remaining arguments
				 * into a single string, separating args with a &
				 * this allows command lines like:
				 *
				 *  test.php v1=test v2=hello+world!
				 *  test.php "v1=test&v2=hello world!"
				 *  test.php v1=test "v2=hello world!"
				 */

			}

			/*
			   we never take stdin if we're (f)cgi, always
			   rely on the web server giving us the info
			   we need in the environment.
			*/

			if core.sapi_globals.request_info.path_translated != nil || cgi != 0 || fastcgi != 0 {
				zend.ZendStreamInitFilename(&file_handle, core.sapi_globals.request_info.path_translated)
			} else {
				zend.ZendStreamInitFp(&file_handle, stdin, "Standard input code")
			}

			/* request startup only after we've done all we can to
			 * get path_translated */

			if core.PhpRequestStartup() == zend.FAILURE {
				if fastcgi != 0 {
					core.FcgiFinishRequest(request, 1)
				}
				core.sapi_globals.server_context = nil
				core.PhpModuleShutdown()
				return zend.FAILURE
			}
			if no_headers != 0 {
				core.sapi_globals.headers_sent = 1
				core.sapi_globals.request_info.no_headers = 1
			}

			/*
			   at this point path_translated will be set if:
			   1. we are running from shell and got filename was there
			   2. we are running as cgi or fastcgi
			*/

			if cgi != 0 || fastcgi != 0 || core.sapi_globals.request_info.path_translated != nil {
				if core.PhpFopenPrimaryScript(&file_handle) == zend.FAILURE {
					var __orig_bailout *sigjmp_buf = zend.EG.bailout
					var __bailout sigjmp_buf
					zend.EG.bailout = &__bailout
					if sigsetjmp(__bailout, 0) == 0 {
						if errno == EACCES {
							core.sapi_globals.sapi_headers.http_response_code = 403
							var __str *byte = "Access denied.\n"
							core.PhpOutputWrite(__str, strlen(__str))
						} else {
							core.sapi_globals.sapi_headers.http_response_code = 404
							var __str *byte = "No input file specified.\n"
							core.PhpOutputWrite(__str, strlen(__str))
						}
					} else {
						zend.EG.bailout = __orig_bailout
					}
					zend.EG.bailout = __orig_bailout

					/* we want to serve more requests if this is fastcgi
					 * so cleanup and continue, request shutdown is
					 * handled later */

					if fastcgi != 0 {
						goto fastcgi_request_done
					}
					if core.sapi_globals.request_info.path_translated != nil {
						zend._efree(core.sapi_globals.request_info.path_translated)
						core.sapi_globals.request_info.path_translated = nil
					}
					if free_query_string != 0 && core.sapi_globals.request_info.query_string != nil {
						zend.Free(core.sapi_globals.request_info.query_string)
						core.sapi_globals.request_info.query_string = nil
					}
					core.PhpRequestShutdown(any(0))
					core.sapi_globals.server_context = nil
					core.PhpModuleShutdown()
					core.SapiShutdown()
					zend.Free(bindpath)
					return zend.FAILURE
				}
			}
			if php_cgi_globals.GetCheckShebangLine() != 0 {
				zend.CG.skip_shebang = 1
			}
			switch behavior {
			case 1:
				core.PhpExecuteScript(&file_handle)
				break
			case 4:
				core.CoreGlobals.during_request_startup = 0
				exit_status = core.PhpLintScript(&file_handle)
				if exit_status == zend.SUCCESS {
					zend.ZendPrintf("No syntax errors detected in %s\n", file_handle.filename)
				} else {
					zend.ZendPrintf("Errors parsing %s\n", file_handle.filename)
				}
				break
			case 5:
				if zend.OpenFileForScanning(&file_handle) == zend.SUCCESS {
					zend.ZendStrip()
					zend.ZendFileHandleDtor(&file_handle)
					core.PhpOutputEndAll()
					core.PhpOutputDeactivate()
					core.PhpOutputShutdown()
				}
				return zend.SUCCESS
				break
			case 2:
				var syntax_highlighter_ini zend.ZendSyntaxHighlighterIni
				if zend.OpenFileForScanning(&file_handle) == zend.SUCCESS {
					standard.PhpGetHighlight(&syntax_highlighter_ini)
					zend.ZendHighlight(&syntax_highlighter_ini)
					if fastcgi != 0 {
						goto fastcgi_request_done
					}
					zend.ZendFileHandleDtor(&file_handle)
					core.PhpOutputEndAll()
					core.PhpOutputDeactivate()
					core.PhpOutputShutdown()
				}
				return zend.SUCCESS
				break
			}
		fastcgi_request_done:
			if core.sapi_globals.request_info.path_translated != nil {
				zend._efree(core.sapi_globals.request_info.path_translated)
				core.sapi_globals.request_info.path_translated = nil
			}
			core.PhpRequestShutdown(any(0))
			if exit_status == 0 {
				exit_status = zend.EG.exit_status
			}
			if free_query_string != 0 && core.sapi_globals.request_info.query_string != nil {
				zend.Free(core.sapi_globals.request_info.query_string)
				core.sapi_globals.request_info.query_string = nil
			}
			if fastcgi == 0 {
				if benchmark != 0 {
					if warmup_repeats != 0 {
						warmup_repeats--
						if warmup_repeats == 0 {
							gettimeofday(&start, nil)
						}
						continue
					} else {
						repeats--
						if repeats > 0 {
							script_file = nil
							PhpOptind = orig_optind
							PhpOptarg = orig_optarg
							continue
						}
					}
				}
				break
			}

			/* only fastcgi will get here */

			requests++
			if max_requests != 0 && requests == max_requests {
				core.FcgiFinishRequest(request, 1)
				zend.Free(bindpath)
				if max_requests != 1 {

					/* no need to return exit_status of the last request */

					exit_status = 0

					/* no need to return exit_status of the last request */

				}
				break
			}
		}
		if request != nil {
			core.FcgiDestroyRequest(request)
		}
		core.FcgiShutdown()
		if CgiSapiModule.php_ini_path_override != nil {
			zend.Free(CgiSapiModule.php_ini_path_override)
		}
		if CgiSapiModule.ini_entries != nil {
			zend.Free(CgiSapiModule.ini_entries)
		}
	} else {
		zend.EG.bailout = __orig_bailout
		exit_status = 255
	}
	zend.EG.bailout = __orig_bailout
out:
	if benchmark != 0 {
		var sec int
		var usec int
		gettimeofday(&end, nil)
		sec = int(end.tv_sec - start.tv_sec)
		if end.tv_usec >= start.tv_usec {
			usec = int(end.tv_usec - start.tv_usec)
		} else {
			sec -= 1
			usec = int(end.tv_usec + 1000000 - start.tv_usec)
		}
		r.Fprintf(stderr, "\nElapsed time: %d.%06d sec\n", sec, usec)
	}
parent_out:
	core.sapi_globals.server_context = nil
	core.PhpModuleShutdown()
	core.SapiShutdown()
	return exit_status
}

/* }}} */
