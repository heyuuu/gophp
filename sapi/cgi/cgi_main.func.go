// <<generate>>

package cgi

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/ext/standard"
	r "sik/runtime"
	"sik/sapi/cli"
	"sik/zend"
	"sort"
)

func UserConfigCacheEntryDtor(el *zend.Zval) {
	var entry *UserConfigCacheEntry = (*UserConfigCacheEntry)(el.GetPtr())
	entry.GetUserConfig().Destroy()
	zend.Free(entry.GetUserConfig())
	zend.Free(entry)
}
func CGIG(v __auto__) __auto__ { return php_cgi_globals.v }
func FcgiLog(type_ int, format *byte, _ ...any) {
	var ap va_list
	va_start(ap, format)
	vfprintf(stderr, format, ap)
	va_end(ap)
}
func ModuleNameCmp(a any, b any) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	return strcasecmp((*zend.ZendModuleEntry)(zend.Z_PTR(f.GetVal())).GetName(), (*zend.ZendModuleEntry)(zend.Z_PTR(s.GetVal())).GetName())
}
func PrintModules() {
	var sorted_registry zend.HashTable
	var module *zend.ZendModuleEntry
	zend.ZendHashInit(&sorted_registry, 64, nil, nil, 1)
	zend.ZendHashCopy(&sorted_registry, &zend.ModuleRegistry, nil)
	sorted_registry.SortCompatible(ModuleNameCmp, 0)
	var __ht *zend.HashTable = &sorted_registry
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		module = _z.GetPtr()
		core.PhpPrintf("%s\n", module.GetName())
	}
	sorted_registry.Destroy()
}
func PrintExtensionInfo(ext *zend.ZendExtension, arg any) int {
	core.PhpPrintf("%s\n", ext.GetName())
	return 0
}
func PrintExtensions() {
	elements := zend.ZendExtensions.ElementsData()
	sort.Slice(elements, func(i, j int) bool {
		ext1 := elements[i].(*zend.ZendExtension)
		ext2 := elements[j].(*zend.ZendExtension)
		return ext1.GetName() < ext2.GetName()
	})

	for _, element := range elements {
		ext := element.(*zend.ZendExtension)
		PrintExtensionInfo(ext, nil)
	}
}

func SapiCgiSingleWrite(str *byte, str_length int) int {
	var ret int
	ret = write(STDOUT_FILENO, str, str_length)
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
	var request *core.FcgiRequest = (*core.FcgiRequest)(core.SG__().server_context)
	for remaining > 0 {
		var to_write int = b.CondF2(remaining > core.INT_MAX, core.INT_MAX, func() int { return int(remaining) })
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
	if r.Fflush(stdout) == r.EOF {
		core.PhpHandleAbortedConnection()
	}
}
func SapiFcgiFlush(server_context any) {
	var request *core.FcgiRequest = (*core.FcgiRequest)(server_context)
	if Parent == 0 && request != nil && core.FcgiFlush(request, 0) == 0 {
		core.PhpHandleAbortedConnection()
	}
}
func SapiCgiSendHeaders(sapi_headers *core.SapiHeaders) int {
	var h *core.SapiHeader
	var pos zend.ZendLlistPosition
	var ignore_status zend.ZendBool = 0
	var response_status int = core.SG__().sapi_headers.http_response_code
	if core.SG__().request_info.no_headers == 1 {
		return core.SAPI_HEADER_SENT_SUCCESSFULLY
	}
	if CGIG(nph) || core.SG__().sapi_headers.http_response_code != 200 {
		var len_ int
		var has_status zend.ZendBool = 0
		var buf []byte
		if CGIG(rfc2616_headers) && core.SG__().sapi_headers.http_status_line {
			var s *byte
			len_ = core.Slprintf(buf, SAPI_CGI_MAX_HEADER_LENGTH, "%s", core.SG__().sapi_headers.http_status_line)
			if b.Assign(&s, strchr(core.SG__().sapi_headers.http_status_line, ' ')) {
				response_status = atoi(s + 1)
			}
			if len_ > SAPI_CGI_MAX_HEADER_LENGTH {
				len_ = SAPI_CGI_MAX_HEADER_LENGTH
			}
		} else {
			var s *byte
			if core.SG__().sapi_headers.http_status_line && b.Assign(&s, strchr(core.SG__().sapi_headers.http_status_line, ' ')) != 0 && s-core.SG__().sapi_headers.http_status_line >= 5 && strncasecmp(core.SG__().sapi_headers.http_status_line, "HTTP/", 5) == 0 {
				len_ = core.Slprintf(buf, b.SizeOf("buf"), "Status:%s", s)
				response_status = atoi(s + 1)
			} else {
				h = (*core.SapiHeader)(zend.ZendLlistGetFirstEx(sapi_headers.GetHeaders(), &pos))
				for h != nil {
					if h.GetHeaderLen() > b.SizeOf("\"Status:\"")-1 && strncasecmp(h.GetHeader(), "Status:", b.SizeOf("\"Status:\"")-1) == 0 {
						has_status = 1
						break
					}
					h = (*core.SapiHeader)(zend.ZendLlistGetNextEx(sapi_headers.GetHeaders(), &pos))
				}
				if has_status == 0 {
					var err *core.HttpResponseStatusCodePair = (*core.HttpResponseStatusCodePair)(core.HttpStatusMap)
					for err.GetCode() != 0 {
						if err.GetCode() == core.SG__().sapi_headers.http_response_code {
							break
						}
						err++
					}
					if err.GetStr() != nil {
						len_ = core.Slprintf(buf, b.SizeOf("buf"), "Status: %d %s", core.SG__().sapi_headers.http_response_code, err.GetStr())
					} else {
						len_ = core.Slprintf(buf, b.SizeOf("buf"), "Status: %d", core.SG__().sapi_headers.http_response_code)
					}
				}
			}
		}
		if has_status == 0 {
			core.PHPWRITE_H(buf, len_)
			core.PHPWRITE_H("\r\n", 2)
			ignore_status = 1
		}
	}
	h = (*core.SapiHeader)(zend.ZendLlistGetFirstEx(sapi_headers.GetHeaders(), &pos))
	for h != nil {

		/* prevent CRLFCRLF */

		if h.GetHeaderLen() != 0 {
			if h.GetHeaderLen() > b.SizeOf("\"Status:\"")-1 && strncasecmp(h.GetHeader(), "Status:", b.SizeOf("\"Status:\"")-1) == 0 {
				if ignore_status == 0 {
					ignore_status = 1
					core.PHPWRITE_H(h.GetHeader(), h.GetHeaderLen())
					core.PHPWRITE_H("\r\n", 2)
				}
			} else if response_status == 304 && h.GetHeaderLen() > b.SizeOf("\"Content-Type:\"")-1 && strncasecmp(h.GetHeader(), "Content-Type:", b.SizeOf("\"Content-Type:\"")-1) == 0 {
				h = (*core.SapiHeader)(zend.ZendLlistGetNextEx(sapi_headers.GetHeaders(), &pos))
				continue
			} else {
				core.PHPWRITE_H(h.GetHeader(), h.GetHeaderLen())
				core.PHPWRITE_H("\r\n", 2)
			}
		}
		h = (*core.SapiHeader)(zend.ZendLlistGetNextEx(sapi_headers.GetHeaders(), &pos))
	}
	core.PHPWRITE_H("\r\n", 2)
	return core.SAPI_HEADER_SENT_SUCCESSFULLY
}
func SapiCgiReadPost(buffer *byte, count_bytes int) int {
	var read_bytes int = 0
	var tmp_read_bytes int
	var remaining_bytes int
	r.Assert(core.SG__().request_info.content_length >= core.SG__().read_post_bytes)
	remaining_bytes = size_t(core.SG__().request_info.content_length - core.SG__().read_post_bytes)
	count_bytes = cli.MIN(count_bytes, remaining_bytes)
	for read_bytes < count_bytes {
		tmp_read_bytes = read(STDIN_FILENO, buffer+read_bytes, count_bytes-read_bytes)
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
	var request *core.FcgiRequest = (*core.FcgiRequest)(core.SG__().server_context)
	var remaining int = core.SG__().request_info.content_length - core.SG__().read_post_bytes
	if remaining < count_bytes {
		count_bytes = remaining
	}
	for read_bytes < count_bytes {
		var diff int = count_bytes - read_bytes
		var to_read int = b.CondF2(diff > core.INT_MAX, core.INT_MAX, func() int { return int(diff) })
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

	var request *core.FcgiRequest = (*core.FcgiRequest)(core.SG__().server_context)
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
	var request *core.FcgiRequest = (*core.FcgiRequest)(core.SG__().server_context)
	return core.FCGI_GETENV(request, "HTTP_COOKIE")
}
func CgiPhpLoadEnvVar(var_ *byte, var_len uint, val *byte, val_len uint, arg any) {
	var array_ptr *zend.Zval = (*zend.Zval)(arg)
	var filter_arg int = b.Cond(array_ptr.GetArr() == core.PG(http_globals)[core.TRACK_VARS_ENV].GetArr(), core.PARSE_ENV, core.PARSE_SERVER)
	var new_val_len int
	if core.sapi_module.GetInputFilter()(filter_arg, var_, &val, strlen(val), &new_val_len) != 0 {
		core.PhpRegisterVariableSafe(var_, val, new_val_len, array_ptr)
	}
}
func CgiPhpImportEnvironmentVariables(array_ptr *zend.Zval) {
	if core.PG(variables_order) && (strchr(core.PG(variables_order), 'E') || strchr(core.PG(variables_order), 'e')) {
		if core.PG(http_globals)[core.TRACK_VARS_ENV].u1.v.type_ != zend.IS_ARRAY {
			zend.ZendIsAutoGlobalStr("_ENV", b.SizeOf("\"_ENV\"")-1)
		}
		if core.PG(http_globals)[core.TRACK_VARS_ENV].u1.v.type_ == zend.IS_ARRAY && array_ptr.GetArr() != core.PG(http_globals)[core.TRACK_VARS_ENV].GetArr() {
			array_ptr.GetArr().DestroyEx()
			array_ptr.SetArr(zend.ZendArrayDup(core.PG(http_globals)[core.TRACK_VARS_ENV].GetArr()))
			return
		}
	}

	/* call php's original import as a catch-all */

	PhpPhpImportEnvironmentVariables(array_ptr)
	if core.FcgiIsFastcgi() != 0 {
		var request *core.FcgiRequest = (*core.FcgiRequest)(core.SG__().server_context)
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
	if CGIG(fix_pathinfo) {
		var script_name *byte = core.SG__().request_info.request_uri
		var path_info *byte
		var free_php_self int
		if core.FcgiIsFastcgi() != 0 {
			var request *core.FcgiRequest = (*core.FcgiRequest)(core.SG__().server_context)
			path_info = core.FCGI_GETENV(request, "PATH_INFO")
		} else {
			path_info = getenv("PATH_INFO")
		}
		if path_info != nil {
			var path_info_len int = strlen(path_info)
			if script_name != nil {
				var script_name_len int = strlen(script_name)
				php_self_len = script_name_len + path_info_len
				php_self = zend.DoAlloca(php_self_len+1, use_heap)
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

		if core.sapi_module.GetInputFilter()(core.PARSE_SERVER, "PHP_SELF", &php_self, php_self_len, &php_self_len) != 0 {
			core.PhpRegisterVariableSafe("PHP_SELF", php_self, php_self_len, track_vars_array)
		}
		if free_php_self != 0 {
			zend.FreeAlloca(php_self, use_heap)
		}
	} else {
		if core.SG__().request_info.request_uri {
			php_self = core.SG__().request_info.request_uri
		} else {
			php_self = ""
		}
		php_self_len = strlen(php_self)
		if core.sapi_module.GetInputFilter()(core.PARSE_SERVER, "PHP_SELF", &php_self, php_self_len, &php_self_len) != 0 {
			core.PhpRegisterVariableSafe("PHP_SELF", php_self, php_self_len, track_vars_array)
		}
	}
}
func SapiCgiLogMessage(message *byte, syslog_type_int int) {
	if core.FcgiIsFastcgi() != 0 && CGIG(fcgi_logging) {
		var request *core.FcgiRequest
		request = (*core.FcgiRequest)(core.SG__().server_context)
		if request != nil {
			var ret int
			var len_ int = int(strlen(message))
			var buf *byte = zend.Malloc(len_ + 2)
			memcpy(buf, message, len_)
			memcpy(buf+len_, "\n", b.SizeOf("\"\\n\""))
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
func PhpCgiIniActivateUserConfig(path *byte, path_len int, doc_root *byte, doc_root_len int) {
	var new_entry *UserConfigCacheEntry
	var entry *UserConfigCacheEntry
	var request_time int64 = int64(core.SapiGetRequestTime())

	/* Find cached config entry: If not found, create one */

	if b.Assign(&entry, zend.ZendHashStrFindPtr(&(CGIG(user_config_cache)), path, path_len)) == nil {
		new_entry = zend.Pemalloc(b.SizeOf("user_config_cache_entry"), 1)
		new_entry.SetExpires(0)
		new_entry.SetUserConfig((*zend.HashTable)(zend.Pemalloc(b.SizeOf("HashTable"), 1)))
		zend.ZendHashInit(new_entry.GetUserConfig(), 8, nil, zend.DtorFuncT(core.ConfigZvalDtor), 1)
		entry = zend.ZendHashStrUpdatePtr(&(CGIG(user_config_cache)), path, path_len, new_entry)
	}

	/* Check whether cache entry has expired and rescan if it is */

	if request_time > entry.GetExpires() {
		var real_path *byte = nil
		var s1 *byte
		var s2 *byte
		var s_len int

		/* Clear the expired config */

		entry.GetUserConfig().Clean()
		if !(zend.IS_ABSOLUTE_PATH(path, path_len)) {
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
			for b.Assign(&ptr, strchr(ptr, zend.DEFAULT_SLASH)) != nil {
				*ptr = 0
				core.PhpParseUserIniFile(path, core.PG(user_ini_filename), entry.GetUserConfig())
				*ptr = '/'
				ptr++
			}
		} else {
			core.PhpParseUserIniFile(path, core.PG(user_ini_filename), entry.GetUserConfig())
		}
		if real_path != nil {
			zend.Efree(real_path)
		}
		entry.SetExpires(request_time + core.PG(user_ini_cache_ttl))
	}

	/* Activate ini entries with values from the user config hash */

	core.PhpIniActivateConfig(entry.GetUserConfig(), core.PHP_INI_PERDIR, core.PHP_INI_STAGE_HTACCESS)

	/* Activate ini entries with values from the user config hash */
}
func SapiCgiActivate() int {
	/* PATH_TRANSLATED should be defined at this stage but better safe than sorry :) */

	if !(core.SG__().request_info.path_translated) {
		return zend.FAILURE
	}
	if core.PhpIniHasPerHostConfig() != 0 {
		var server_name *byte

		/* Activate per-host-system-configuration defined in php.ini and stored into configuration_hash during startup */

		if core.FcgiIsFastcgi() != 0 {
			var request *core.FcgiRequest = (*core.FcgiRequest)(core.SG__().server_context)
			server_name = core.FCGI_GETENV(request, "SERVER_NAME")
		} else {
			server_name = getenv("SERVER_NAME")
		}

		/* SERVER_NAME should also be defined at this stage..but better check it anyway */

		if server_name != nil {
			var server_name_len int = strlen(server_name)
			server_name = zend.Estrndup(server_name, server_name_len)
			zend.ZendStrTolower(server_name, server_name_len)
			core.PhpIniActivatePerHostConfig(server_name, server_name_len)
			zend.Efree(server_name)
		}

		/* SERVER_NAME should also be defined at this stage..but better check it anyway */

	}
	if core.PhpIniHasPerDirConfig() != 0 || core.PG(user_ini_filename) && (*core.PG)(user_ini_filename) {
		var path *byte
		var path_len int

		/* Prepare search path */

		path_len = strlen(core.SG__().request_info.path_translated)

		/* Make sure we have trailing slash! */

		if !(zend.IS_SLASH(core.SG__().request_info.path_translated[path_len])) {
			path = zend.Emalloc(path_len + 2)
			memcpy(path, core.SG__().request_info.path_translated, path_len+1)
			path_len = zend.ZendDirname(path, path_len)
			path[b.PostInc(&path_len)] = zend.DEFAULT_SLASH
		} else {
			path = zend.Estrndup(core.SG__().request_info.path_translated, path_len)
			path_len = zend.ZendDirname(path, path_len)
		}
		path[path_len] = 0

		/* Activate per-dir-system-configuration defined in php.ini and stored into configuration_hash during startup */

		core.PhpIniActivatePerDirConfig(path, path_len)

		/* Load and activate user ini files in path starting from DOCUMENT_ROOT */

		if core.PG(user_ini_filename) && (*core.PG)(user_ini_filename) {
			var doc_root *byte
			if core.FcgiIsFastcgi() != 0 {
				var request *core.FcgiRequest = (*core.FcgiRequest)(core.SG__().server_context)
				doc_root = core.FCGI_GETENV(request, "DOCUMENT_ROOT")
			} else {
				doc_root = getenv("DOCUMENT_ROOT")
			}

			/* DOCUMENT_ROOT should also be defined at this stage..but better check it anyway */

			if doc_root != nil {
				var doc_root_len int = strlen(doc_root)
				if doc_root_len > 0 && zend.IS_SLASH(doc_root[doc_root_len-1]) {
					doc_root_len--
				}
				PhpCgiIniActivateUserConfig(path, path_len, doc_root, doc_root_len)
			}

			/* DOCUMENT_ROOT should also be defined at this stage..but better check it anyway */

		}
		zend.Efree(path)
	}
	return zend.SUCCESS
}
func SapiCgiDeactivate() int {
	/* flush only when SAPI was started. The reasons are:
	   1. SAPI Deactivate is called from two places: module init and request shutdown
	   2. When the first call occurs and the request is not set up, flush fails on FastCGI.
	*/

	if core.SG__().sapi_started {
		if core.FcgiIsFastcgi() != 0 {
			if Parent == 0 && core.FcgiFinishRequest((*core.FcgiRequest)(core.SG__().server_context), 0) == 0 {
				core.PhpHandleAbortedConnection()
			}
		} else {
			SapiCgiFlush(core.SG__().server_context)
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
func IsValidPath(path *byte) int {
	var p *byte = path
	if p == nil {
		return 0
	}
	if (*p) == '.' && (*(p + 1)) == '.' && (!(*(p + 2)) || zend.IS_SLASH(*(p + 2))) {
		return 0
	}
	for *p {
		if zend.IS_SLASH(*p) {
			p++
			if (*p) == '.' {
				p++
				if (*p) == '.' {
					p++
					if !(*p) || zend.IS_SLASH(*p) {
						return 0
					}
				}
			}
		}
		p++
	}
	return 1
}
func CGI_GETENV(name string) *byte {
	if has_env {
		return core.FCGI_GETENV(request, name)
	} else {
		return getenv(name)
	}
}
func CGI_PUTENV(name string, value *byte) *byte {
	if has_env {
		return core.FCGI_PUTENV(request, name, value)
	} else {
		return _sapiCgiPutenv(name, b.SizeOf("name")-1, value)
	}
}
func InitRequestInfo(request *core.FcgiRequest) {
	var has_env int = core.FcgiHasEnv(request)
	var env_script_filename *byte = CGI_GETENV("SCRIPT_FILENAME")
	var env_path_translated *byte = CGI_GETENV("PATH_TRANSLATED")
	var script_path_translated *byte = env_script_filename

	/* some broken servers do not have script_filename or argv0
	 * an example, IIS configured in some ways.  then they do more
	 * broken stuff and set path_translated to the cgi script location */

	if script_path_translated == nil && env_path_translated != nil {
		script_path_translated = env_path_translated
	}

	/* initialize the defaults */

	core.SG__().request_info.path_translated = nil
	core.SG__().request_info.request_method = nil
	core.SG__().request_info.proto_num = 1000
	core.SG__().request_info.query_string = nil
	core.SG__().request_info.request_uri = nil
	core.SG__().request_info.content_type = nil
	core.SG__().request_info.content_length = 0
	core.SG__().sapi_headers.http_response_code = 200

	/* script_path_translated being set is a good indication that
	 * we are running in a cgi environment, since it is always
	 * null otherwise.  otherwise, the filename
	 * of the script will be retrieved later via argc/argv */

	if script_path_translated != nil {
		var auth *byte
		var content_length *byte = CGI_GETENV("CONTENT_LENGTH")
		var content_type *byte = CGI_GETENV("CONTENT_TYPE")
		var env_path_info *byte = CGI_GETENV("PATH_INFO")
		var env_script_name *byte = CGI_GETENV("SCRIPT_NAME")
		if CGIG(fix_pathinfo) {
			var st zend.ZendStatT
			var real_path *byte = nil
			var env_redirect_url *byte = CGI_GETENV("REDIRECT_URL")
			var env_document_root *byte = CGI_GETENV("DOCUMENT_ROOT")
			var orig_path_translated *byte = env_path_translated
			var orig_path_info *byte = env_path_info
			var orig_script_name *byte = env_script_name
			var orig_script_filename *byte = env_script_filename
			var script_path_translated_len int
			if env_document_root == nil && core.PG(doc_root) {
				env_document_root = CGI_PUTENV("DOCUMENT_ROOT", core.PG(doc_root))

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

			if script_path_translated != nil && b.Assign(&script_path_translated_len, strlen(script_path_translated)) > 0 && (script_path_translated[script_path_translated_len-1] == '/' || b.Assign(&real_path, zend.TsrmRealpath(script_path_translated, nil)) == nil) {
				var pt *byte = zend.Estrndup(script_path_translated, script_path_translated_len)
				var len_ int = script_path_translated_len
				var ptr *byte
				for b.Assign(&ptr, strrchr(pt, '/')) || b.Assign(&ptr, strrchr(pt, '\\')) {
					*ptr = 0
					if zend.ZendStat(pt, &st) == 0 && zend.S_ISREG(st.st_mode) {

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
						var pilen int = b.CondF1(env_path_info != nil, func() __auto__ { return strlen(env_path_info) }, 0)
						var path_info *byte = b.Cond(env_path_info != nil, env_path_info+pilen-slen, nil)
						if orig_path_info != path_info {
							if orig_path_info != nil {
								var old byte
								CGI_PUTENV("ORIG_PATH_INFO", orig_path_info)
								old = path_info[0]
								path_info[0] = 0
								if orig_script_name == nil || strcmp(orig_script_name, env_path_info) != 0 {
									if orig_script_name != nil {
										CGI_PUTENV("ORIG_SCRIPT_NAME", orig_script_name)
									}
									core.SG__().request_info.request_uri = CGI_PUTENV("SCRIPT_NAME", env_path_info)
								} else {
									core.SG__().request_info.request_uri = orig_script_name
								}
								path_info[0] = old
							}
							env_path_info = CGI_PUTENV("PATH_INFO", path_info)
						}
						if orig_script_filename == nil || strcmp(orig_script_filename, pt) != 0 {
							if orig_script_filename != nil {
								CGI_PUTENV("ORIG_SCRIPT_FILENAME", orig_script_filename)
							}
							script_path_translated = CGI_PUTENV("SCRIPT_FILENAME", pt)
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

							path_translated_len = l + b.CondF1(env_path_info != nil, func() __auto__ { return strlen(env_path_info) }, 0)
							path_translated = (*byte)(zend.Emalloc(path_translated_len + 1))
							memcpy(path_translated, env_document_root, l)
							if env_path_info != nil {
								memcpy(path_translated+l, env_path_info, path_translated_len-l)
							}
							path_translated[path_translated_len] = '0'
							if orig_path_translated != nil {
								CGI_PUTENV("ORIG_PATH_TRANSLATED", orig_path_translated)
							}
							env_path_translated = CGI_PUTENV("PATH_TRANSLATED", path_translated)
							zend.Efree(path_translated)
						} else if env_script_name != nil && strstr(pt, env_script_name) {

							/* PATH_TRANSLATED = PATH_TRANSLATED - SCRIPT_NAME + PATH_INFO */

							var ptlen int = strlen(pt) - strlen(env_script_name)
							var path_translated_len int = ptlen + b.CondF1(env_path_info != nil, func() __auto__ { return strlen(env_path_info) }, 0)
							var path_translated *byte = (*byte)(zend.Emalloc(path_translated_len + 1))
							memcpy(path_translated, pt, ptlen)
							if env_path_info != nil {
								memcpy(path_translated+ptlen, env_path_info, path_translated_len-ptlen)
							}
							path_translated[path_translated_len] = '0'
							if orig_path_translated != nil {
								CGI_PUTENV("ORIG_PATH_TRANSLATED", orig_path_translated)
							}
							env_path_translated = CGI_PUTENV("PATH_TRANSLATED", path_translated)
							zend.Efree(path_translated)
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
						CGI_PUTENV("ORIG_SCRIPT_FILENAME", orig_script_filename)
					}
					script_path_translated = CGI_PUTENV("SCRIPT_FILENAME", nil)
					core.SG__().sapi_headers.http_response_code = 404
				}
				if !(core.SG__().request_info.request_uri) {
					if orig_script_name == nil || strcmp(orig_script_name, env_script_name) != 0 {
						if orig_script_name != nil {
							CGI_PUTENV("ORIG_SCRIPT_NAME", orig_script_name)
						}
						core.SG__().request_info.request_uri = CGI_PUTENV("SCRIPT_NAME", env_script_name)
					} else {
						core.SG__().request_info.request_uri = orig_script_name
					}
				}
				if pt != nil {
					zend.Efree(pt)
				}
			} else {

				/* make sure path_info/translated are empty */

				if orig_script_filename == nil || script_path_translated != orig_script_filename && strcmp(script_path_translated, orig_script_filename) != 0 {
					if orig_script_filename != nil {
						CGI_PUTENV("ORIG_SCRIPT_FILENAME", orig_script_filename)
					}
					script_path_translated = CGI_PUTENV("SCRIPT_FILENAME", script_path_translated)
				}
				if env_redirect_url != nil {
					if orig_path_info != nil {
						CGI_PUTENV("ORIG_PATH_INFO", orig_path_info)
						CGI_PUTENV("PATH_INFO", nil)
					}
					if orig_path_translated != nil {
						CGI_PUTENV("ORIG_PATH_TRANSLATED", orig_path_translated)
						CGI_PUTENV("PATH_TRANSLATED", nil)
					}
				}
				if env_script_name != orig_script_name {
					if orig_script_name != nil {
						CGI_PUTENV("ORIG_SCRIPT_NAME", orig_script_name)
					}
					core.SG__().request_info.request_uri = CGI_PUTENV("SCRIPT_NAME", env_script_name)
				} else {
					core.SG__().request_info.request_uri = env_script_name
				}
				zend.Efree(real_path)
			}

			/*
			 * if the file doesn't exist, try to extract PATH_INFO out
			 * of it by stat'ing back through the '/'
			 * this fixes url's like /info.php/test
			 */

		} else {

			/* pre 4.3 behaviour, shouldn't be used but provides BC */

			if env_path_info != nil {
				core.SG__().request_info.request_uri = env_path_info
			} else {
				core.SG__().request_info.request_uri = env_script_name
			}
			if !(CGIG(discard_path)) && env_path_translated != nil {
				script_path_translated = env_path_translated
			}
		}
		if IsValidPath(script_path_translated) != 0 {
			core.SG__().request_info.path_translated = zend.Estrdup(script_path_translated)
		}
		core.SG__().request_info.request_method = CGI_GETENV("REQUEST_METHOD")

		/* FIXME - Work out proto_num here */

		core.SG__().request_info.query_string = CGI_GETENV("QUERY_STRING")
		if content_type != nil {
			core.SG__().request_info.content_type = content_type
		} else {
			core.SG__().request_info.content_type = ""
		}
		if content_length != nil {
			core.SG__().request_info.content_length = atol(content_length)
		} else {
			core.SG__().request_info.content_length = 0
		}

		/* The CGI RFC allows servers to pass on unvalidated Authorization data */

		auth = CGI_GETENV("HTTP_AUTHORIZATION")
		core.PhpHandleAuthData(auth)
	}

	/* script_path_translated being set is a good indication that
	 * we are running in a cgi environment, since it is always
	 * null otherwise.  otherwise, the filename
	 * of the script will be retrieved later via argc/argv */
}
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
func PhpCgiGlobalsCtor(php_cgi_globals *php_cgi_globals_struct) {
	php_cgi_globals.SetRfc2616Headers(0)
	php_cgi_globals.SetNph(0)
	php_cgi_globals.SetCheckShebangLine(1)
	php_cgi_globals.SetForceRedirect(1)
	php_cgi_globals.SetRedirectStatusEnv(nil)
	php_cgi_globals.SetFixPathinfo(1)
	php_cgi_globals.SetDiscardPath(0)
	php_cgi_globals.SetFcgiLogging(1)
	zend.ZendHashInit(php_cgi_globals.GetUserConfigCache(), 8, nil, UserConfigCacheEntryDtor, 1)
}
func ZmStartupCgi(type_ int, module_number int) int {
	zend.REGISTER_INI_ENTRIES()
	return zend.SUCCESS
}
func ZmShutdownCgi(type_ int, module_number int) int {
	CGIG(user_config_cache).Destroy()
	zend.UNREGISTER_INI_ENTRIES()
	return zend.SUCCESS
}
func ZmInfoCgi(zend_module *zend.ZendModuleEntry) { zend.DISPLAY_INI_ENTRIES() }
func ZifApacheChildTerminate(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() != 0 {
		return
	}
	if core.FcgiIsFastcgi() != 0 {
		core.FcgiTerminate()
	}
}
func ZifApacheRequestHeaders(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() != 0 {
		return
	}
	zend.ArrayInit(return_value)
	if core.FcgiIsFastcgi() != 0 {
		var request *core.FcgiRequest = (*core.FcgiRequest)(core.SG__().server_context)
		core.FcgiLoadenv(request, core.SapiAddRequestHeader, return_value)
	} else {
		var buf []byte
		var env **byte
		var p **byte
		var q **byte
		var var_ **byte
		var val **byte
		var t **byte = buf
		var alloc_size int = b.SizeOf("buf")
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
					t = zend.Emalloc(alloc_size)
				} else {
					t = zend.Erealloc(t, alloc_size)
				}
			}
			var_ = *env
			if var_len > 5 && var_[0] == 'H' && var_[1] == 'T' && var_[2] == 'T' && var_[3] == 'P' && var_[4] == '_' {
				var_len -= 5
				if var_len >= alloc_size {
					alloc_size = var_len + 64
					if t == buf {
						t = zend.Emalloc(alloc_size)
					} else {
						t = zend.Erealloc(t, alloc_size)
					}
				}
				p = var_ + 5
				q = t
				var_ = q

				/* First char keep uppercase */

				*p++
				b.PostInc(&(*q)) = (*p) - 1
				for (*p) != nil {
					if (*p) == '=' {

						/* End of name */

						break

						/* End of name */

					} else if (*p) == '_' {
						b.PostInc(&(*q)) = '-'
						p++

						/* First char after - keep uppercase */

						if (*p) != nil && (*p) != '=' {
							*p++
							b.PostInc(&(*q)) = (*p) - 1
						}

						/* First char after - keep uppercase */

					} else if (*p) >= 'A' && (*p) <= 'Z' {

						/* lowercase */

						b.PostInc(&(*q)) = b.PostInc(&(*p)) - 'A' + 'a'

						/* lowercase */

					} else {
						*p++
						b.PostInc(&(*q)) = (*p) - 1
					}
				}
				*q = 0
			} else if var_len == b.SizeOf("\"CONTENT_TYPE\"")-1 && memcmp(var_, "CONTENT_TYPE", b.SizeOf("\"CONTENT_TYPE\"")-1) == 0 {
				var_ = "Content-Type"
			} else if var_len == b.SizeOf("\"CONTENT_LENGTH\"")-1 && memcmp(var_, "CONTENT_LENGTH", b.SizeOf("\"CONTENT_LENGTH\"")-1) == 0 {
				var_ = "Content-Length"
			} else {
				continue
			}
			val++
			zend.AddAssocStringEx(return_value, var_, var_len, val)
		}
		if t != buf && t != nil {
			zend.Efree(t)
		}
	}
}
func AddResponseHeader(h *core.SapiHeader, return_value *zend.Zval) {
	if h.GetHeaderLen() > 0 {
		var s *byte
		var len_ int = 0
		var p *byte = strchr(h.GetHeader(), ':')
		if nil != p {
			len_ = p - h.GetHeader()
		}
		if len_ > 0 {
			for len_ != 0 && (h.GetHeader()[len_-1] == ' ' || h.GetHeader()[len_-1] == '\t') {
				len_--
			}
			if len_ != 0 {
				s = zend.DoAlloca(len_+1, use_heap)
				memcpy(s, h.GetHeader(), len_)
				s[len_] = 0
				for {
					p++
					if !((*p) == ' ' || (*p) == '\t') {
						break
					}
				}
				zend.AddAssocStringlEx(return_value, s, len_, p, h.GetHeaderLen()-(p-h.GetHeader()))
				zend.FreeAlloca(s, use_heap)
			}
		}
	}
}
func ZifApacheResponseHeaders(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ArrayInit(return_value)
	zend.ZendLlistApplyWithArgument(core.SG__().sapi_headers.headers, zend.LlistApplyWithArgFuncT(AddResponseHeader), return_value)
}
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

	var behavior int = PHP_MODE_STANDARD
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

	app := core.NewApp()

	zend.ZendSignalStartup()
	PhpCgiGlobalsCtor(&php_cgi_globals)
	app.Startup(&CgiSapiModule)
	fastcgi = core.FcgiIsFastcgi()
	CgiSapiModule.SetPhpIniPathOverride(nil)
	if fastcgi == 0 {

		/* Make sure we detect we are a cgi - a bit redundancy here,
		 * but the default case is that we have to check only the first one. */

		if getenv("SERVER_SOFTWARE") || getenv("SERVER_NAME") || getenv("GATEWAY_INTERFACE") || getenv("REQUEST_METHOD") {
			cgi = 1
		}

		/* Make sure we detect we are a cgi - a bit redundancy here,
		 * but the default case is that we have to check only the first one. */

	}
	if b.Assign(&query_string, getenv("QUERY_STRING")) != nil && strchr(query_string, '=') == nil {

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
	for skip_getopt == 0 && b.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &PhpOptarg, &PhpOptind, 0, 2)) != -1 {
		switch c {
		case 'c':
			if CgiSapiModule.GetPhpIniPathOverride() != nil {
				zend.Free(CgiSapiModule.GetPhpIniPathOverride())
			}
			CgiSapiModule.SetPhpIniPathOverride(strdup(PhpOptarg))
			break
		case 'n':
			CgiSapiModule.SetPhpIniIgnore(1)
			break
		case 'd':

			/* define ini __special__  entries on command line */

			var len_ int = strlen(PhpOptarg)
			var val *byte
			if b.Assign(&val, strchr(PhpOptarg, '=')) {
				val++
				if !(isalnum(*val)) && (*val) != '"' && (*val) != '\'' && (*val) != '0' {
					CgiSapiModule.SetIniEntries(realloc(CgiSapiModule.GetIniEntries(), ini_entries_len+len_+b.SizeOf("\"\\\"\\\"\\n\\0\"")))
					memcpy(CgiSapiModule.GetIniEntries()+ini_entries_len, PhpOptarg, val-PhpOptarg)
					ini_entries_len += val - PhpOptarg
					memcpy(CgiSapiModule.GetIniEntries()+ini_entries_len, "\"", 1)
					ini_entries_len++
					memcpy(CgiSapiModule.GetIniEntries()+ini_entries_len, val, len_-(val-PhpOptarg))
					ini_entries_len += len_ - (val - PhpOptarg)
					memcpy(CgiSapiModule.GetIniEntries()+ini_entries_len, "\"\n0", b.SizeOf("\"\\\"\\n\\0\""))
					ini_entries_len += b.SizeOf("\"\\n\\0\\\"\"") - 2
				} else {
					CgiSapiModule.SetIniEntries(realloc(CgiSapiModule.GetIniEntries(), ini_entries_len+len_+b.SizeOf("\"\\n\\0\"")))
					memcpy(CgiSapiModule.GetIniEntries()+ini_entries_len, PhpOptarg, len_)
					memcpy(CgiSapiModule.GetIniEntries()+ini_entries_len+len_, "\n0", b.SizeOf("\"\\n\\0\""))
					ini_entries_len += len_ + b.SizeOf("\"\\n\\0\"") - 2
				}
			} else {
				CgiSapiModule.SetIniEntries(realloc(CgiSapiModule.GetIniEntries(), ini_entries_len+len_+b.SizeOf("\"=1\\n\\0\"")))
				memcpy(CgiSapiModule.GetIniEntries()+ini_entries_len, PhpOptarg, len_)
				memcpy(CgiSapiModule.GetIniEntries()+ini_entries_len+len_, "=1\n0", b.SizeOf("\"=1\\n\\0\""))
				ini_entries_len += len_ + b.SizeOf("\"=1\\n\\0\"") - 2
			}
			break
		case 'b':
			if fastcgi == 0 {
				bindpath = strdup(PhpOptarg)
			}
			break
		case 's':
			behavior = PHP_MODE_HIGHLIGHT
			break
		}
	}
	PhpOptind = orig_optind
	PhpOptarg = orig_optarg
	if fastcgi != 0 || bindpath != nil {

		/* Override SAPI callbacks */

		CgiSapiModule.SetUbWrite(SapiFcgiUbWrite)
		CgiSapiModule.SetFlush(SapiFcgiFlush)
		CgiSapiModule.SetReadPost(SapiFcgiReadPost)
		CgiSapiModule.SetGetenv(SapiFcgiGetenv)
		CgiSapiModule.SetReadCookies(SapiFcgiReadCookies)
	}
	CgiSapiModule.SetExecutableLocation(argv[0])
	if cgi == 0 && fastcgi == 0 && bindpath == nil {
		CgiSapiModule.SetAdditionalFunctions(AdditionalFunctions)
	}

	/* startup after we get the above ini override se we get things right */

	if !CgiSapiModule.Startup() {
		zend.Free(bindpath)
		return zend.FAILURE
	}

	/* check force_cgi after startup, so we have proper output */

	if cgi != 0 && CGIG(force_redirect) {

		/* Apache will generate REDIRECT_STATUS,
		 * Netscape and redirect.so will generate HTTP_REDIRECT_STATUS.
		 * redirect.so and installation instructions available from
		 * http://www.koehntopp.de/php.
		 *   -- kk@netuse.de
		 */

		if !(getenv("REDIRECT_STATUS")) && !(getenv("HTTP_REDIRECT_STATUS")) && (!(CGIG(redirect_status_env)) || !(getenv(CGIG(redirect_status_env)))) {
			var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
			var __bailout JMP_BUF
			zend.EG__().SetBailout(&__bailout)
			if zend.SETJMP(__bailout) == 0 {
				core.SG__().sapi_headers.http_response_code = 400
				core.PUTS("<b>Security Alert!</b> The PHP CGI cannot be accessed directly.\n\n\n<p>This PHP CGI binary was compiled with force-cgi-redirect enabled.  This\n\nmeans that a page will only be served up if the REDIRECT_STATUS CGI variable is\n\nset, e.g. via an Apache Action directive.</p>\n\n<p>For more information as to <i>why</i> this behaviour exists, see the <a href=\"http://php.net/security.cgi-bin\">\nmanual page for CGI security</a>.</p>\n\n<p>For more information about changing this behaviour or re-enabling this webserver,\n\nconsult the installation file that came with this distribution, or visit \n\n<a href=\"http://php.net/install.windows\">the manual page</a>.</p>\n")
			} else {
				zend.EG__().SetBailout(__orig_bailout)
			}
			zend.EG__().SetBailout(__orig_bailout)
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
			core.FcgiSetMgmtVar("FCGI_MAX_CONNS", b.SizeOf("\"FCGI_MAX_CONNS\"")-1, children_str, strlen(children_str))

			/* This is the number of concurrent requests, equals FCGI_MAX_CONNS */

			core.FcgiSetMgmtVar("FCGI_MAX_REQS", b.SizeOf("\"FCGI_MAX_REQS\"")-1, children_str, strlen(children_str))

			/* This is the number of concurrent requests, equals FCGI_MAX_CONNS */

		} else {
			core.FcgiSetMgmtVar("FCGI_MAX_CONNS", b.SizeOf("\"FCGI_MAX_CONNS\"")-1, "1", b.SizeOf("\"1\"")-1)
			core.FcgiSetMgmtVar("FCGI_MAX_REQS", b.SizeOf("\"FCGI_MAX_REQS\"")-1, "1", b.SizeOf("\"1\"")-1)
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
	zend.EG__().SetBailout(nil)
	var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		for skip_getopt == 0 && b.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &PhpOptarg, &PhpOptind, 1, 2)) != -1 {
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

			case core.PHP_GETOPT_INVALID_ARG:
				if request != nil {
					core.FcgiDestroyRequest(request)
				}
				core.FcgiShutdown()
				no_headers = 1
				core.SG__().headers_sent = 1
				PhpCgiUsage(argv[0])
				core.PhpOutputEndAll()
				exit_status = 0
				if c == core.PHP_GETOPT_INVALID_ARG {
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
				core.SG__().server_context = any(request)
			} else {
				core.SG__().server_context = any(1)
			}
			InitRequestInfo(request)
			if cgi == 0 && fastcgi == 0 {
				for b.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &PhpOptarg, &PhpOptind, 0, 2)) != -1 {
					switch c {
					case 'a':
						r.Printf("Interactive mode enabled\n\n")
						break
					case 'C':
						core.SG__().options |= core.SAPI_OPTION_NO_CHDIR
						break
					case 'e':
						zend.CG__().SetCompilerOptions(zend.CG__().GetCompilerOptions() | zend.ZEND_COMPILE_EXTENDED_INFO)
						break
					case 'f':
						if script_file != nil {
							zend.Efree(script_file)
						}
						script_file = zend.Estrdup(PhpOptarg)
						no_headers = 1
						break
					case 'i':
						if script_file != nil {
							zend.Efree(script_file)
						}
						if core.PhpRequestStartup() == zend.FAILURE {
							core.SG__().server_context = nil
							core.PhpModuleShutdown()
							zend.Free(bindpath)
							return zend.FAILURE
						}
						if no_headers != 0 {
							core.SG__().headers_sent = 1
							core.SG__().request_info.no_headers = 1
						}
						standard.PhpPrintInfo(0xffffffff)
						core.PhpRequestShutdown(any(0))
						core.FcgiShutdown()
						exit_status = 0
						goto out
					case 'l':
						no_headers = 1
						behavior = PHP_MODE_LINT
						break
					case 'm':
						if script_file != nil {
							zend.Efree(script_file)
						}
						core.SG__().headers_sent = 1
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
							zend.Efree(script_file)
						}
						no_headers = 1
						if core.PhpRequestStartup() == zend.FAILURE {
							core.SG__().server_context = nil
							core.PhpModuleShutdown()
							zend.Free(bindpath)
							return zend.FAILURE
						}
						core.SG__().headers_sent = 1
						core.SG__().request_info.no_headers = 1
						core.PhpPrintf("PHP %s (%s) (built: %s %s)\nCopyright (c) The PHP Group\n%s", core.PHP_VERSION, core.SM__().Name(), __DATE__, __TIME__, zend.GetZendVersion())
						core.PhpRequestShutdown(any(0))
						core.FcgiShutdown()
						exit_status = 0
						goto out
					case 'w':
						behavior = PHP_MODE_STRIP
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

					if core.SG__().request_info.path_translated {
						zend.Efree(core.SG__().request_info.path_translated)
					}
					core.SG__().request_info.path_translated = script_file

					/* before registering argv to module exchange the *new* argv[0] */

					core.SG__().request_info.argc = argc - (PhpOptind - 1)
					core.SG__().request_info.argv = &argv[PhpOptind-1]
					core.SG__().request_info.argv[0] = script_file
				} else if argc > PhpOptind {

					/* file is on command line, but not in -f opt */

					if core.SG__().request_info.path_translated {
						zend.Efree(core.SG__().request_info.path_translated)
					}
					core.SG__().request_info.path_translated = zend.Estrdup(argv[PhpOptind])

					/* arguments after the file are considered script args */

					core.SG__().request_info.argc = argc - PhpOptind
					core.SG__().request_info.argv = &argv[PhpOptind]
				}
				if no_headers != 0 {
					core.SG__().headers_sent = 1
					core.SG__().request_info.no_headers = 1
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

				if !(core.SG__().request_info.query_string) && argc > PhpOptind {
					var slen int = strlen(core.PG(arg_separator).input)
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
							strlcat(s, core.PG(arg_separator).input, len_)
						}
					}
					core.SG__().request_info.query_string = s
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

			if core.SG__().request_info.path_translated || cgi != 0 || fastcgi != 0 {
				zend.ZendStreamInitFilename(&file_handle, core.SG__().request_info.path_translated)
			} else {
				zend.ZendStreamInitFp(&file_handle, stdin, "Standard input code")
			}

			/* request startup only after we've done all we can to
			 * get path_translated */

			if core.PhpRequestStartup() == zend.FAILURE {
				if fastcgi != 0 {
					core.FcgiFinishRequest(request, 1)
				}
				core.SG__().server_context = nil
				core.PhpModuleShutdown()
				return zend.FAILURE
			}
			if no_headers != 0 {
				core.SG__().headers_sent = 1
				core.SG__().request_info.no_headers = 1
			}

			/*
			   at this point path_translated will be set if:
			   1. we are running from shell and got filename was there
			   2. we are running as cgi or fastcgi
			*/

			if cgi != 0 || fastcgi != 0 || core.SG__().request_info.path_translated {
				if core.PhpFopenPrimaryScript(&file_handle) == zend.FAILURE {
					var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
					var __bailout JMP_BUF
					zend.EG__().SetBailout(&__bailout)
					if zend.SETJMP(__bailout) == 0 {
						if errno == EACCES {
							core.SG__().sapi_headers.http_response_code = 403
							core.PUTS("Access denied.\n")
						} else {
							core.SG__().sapi_headers.http_response_code = 404
							core.PUTS("No input file specified.\n")
						}
					} else {
						zend.EG__().SetBailout(__orig_bailout)
					}
					zend.EG__().SetBailout(__orig_bailout)

					/* we want to serve more requests if this is fastcgi
					 * so cleanup and continue, request shutdown is
					 * handled later */

					if fastcgi != 0 {
						goto fastcgi_request_done
					}
					if core.SG__().request_info.path_translated {
						zend.Efree(core.SG__().request_info.path_translated)
						core.SG__().request_info.path_translated = nil
					}
					if free_query_string != 0 && core.SG__().request_info.query_string {
						zend.Free(core.SG__().request_info.query_string)
						core.SG__().request_info.query_string = nil
					}
					core.PhpRequestShutdown(any(0))
					core.SG__().server_context = nil
					core.PhpModuleShutdown()
					app.Shutdown()
					zend.Free(bindpath)
					return zend.FAILURE
				}
			}
			if CGIG(check_shebang_line) {
				zend.CG__().SetSkipShebang(1)
			}
			switch behavior {
			case PHP_MODE_STANDARD:
				core.PhpExecuteScript(&file_handle)
				break
			case PHP_MODE_LINT:
				core.PG(during_request_startup) = 0
				exit_status = core.PhpLintScript(&file_handle)
				if exit_status == zend.SUCCESS {
					zend.ZendPrintf("No syntax errors detected in %s\n", file_handle.GetFilename())
				} else {
					zend.ZendPrintf("Errors parsing %s\n", file_handle.GetFilename())
				}
				break
			case PHP_MODE_STRIP:
				if zend.OpenFileForScanning(&file_handle) == zend.SUCCESS {
					zend.ZendStrip()
					zend.ZendFileHandleDtor(&file_handle)
					core.PhpOutputTeardown()
				}
				return zend.SUCCESS
				break
			case PHP_MODE_HIGHLIGHT:
				var syntax_highlighter_ini zend.ZendSyntaxHighlighterIni
				if zend.OpenFileForScanning(&file_handle) == zend.SUCCESS {
					standard.PhpGetHighlight(&syntax_highlighter_ini)
					zend.ZendHighlight(&syntax_highlighter_ini)
					if fastcgi != 0 {
						goto fastcgi_request_done
					}
					zend.ZendFileHandleDtor(&file_handle)
					core.PhpOutputTeardown()
				}
				return zend.SUCCESS
				break
			}
		fastcgi_request_done:
			if core.SG__().request_info.path_translated {
				zend.Efree(core.SG__().request_info.path_translated)
				core.SG__().request_info.path_translated = nil
			}
			core.PhpRequestShutdown(any(0))
			if exit_status == 0 {
				exit_status = zend.EG__().GetExitStatus()
			}
			if free_query_string != 0 && core.SG__().request_info.query_string {
				zend.Free(core.SG__().request_info.query_string)
				core.SG__().request_info.query_string = nil
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
		if CgiSapiModule.GetPhpIniPathOverride() != nil {
			zend.Free(CgiSapiModule.GetPhpIniPathOverride())
		}
		if CgiSapiModule.GetIniEntries() != nil {
			zend.Free(CgiSapiModule.GetIniEntries())
		}
	} else {
		zend.EG__().SetBailout(__orig_bailout)
		exit_status = 255
	}
	zend.EG__().SetBailout(__orig_bailout)
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
	core.SG__().server_context = nil
	core.PhpModuleShutdown()
	app.Shutdown()
	return exit_status
}
