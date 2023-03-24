// <<generate>>

package cgi

import (
	"log"
	b "sik/builtin"
	r "sik/builtin/file"
	"sik/core"
	"sik/sapi/cli"
	"sik/zend"
	"sik/zend/types"
	"sik/zend/zpp"
	"sort"
)

func UserConfigCacheEntryDtor(el *types.Zval) {
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
func PrintModules() {
	var modules = zend.CopyRegistryModules()
	sort.Slice(modules, func(i, j int) bool {
		return b.StrCaseCompare(modules[i].GetNameStr(), modules[j].GetNameStr())
	})
	for _, module := range modules {
		core.PhpPrintf("%s\n", module.GetName())
	}
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
	var ignore_status types.ZendBool = 0
	var response_status int = core.SG__().sapi_headers.http_response_code
	if core.SG__().request_info.no_headers == 1 {
		return core.SAPI_HEADER_SENT_SUCCESSFULLY
	}
	if CGIG(nph) || core.SG__().sapi_headers.http_response_code != 200 {
		var len_ int
		var has_status types.ZendBool = 0
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
					code := core.SG__().sapi_headers.http_response_code
					if codeStr, ok := core.HttpStatusMap[code]; ok {
						len_ = core.Slprintf(buf, b.SizeOf("buf"), "Status: %d %s", code, codeStr)
					} else {
						len_ = core.Slprintf(buf, b.SizeOf("buf"), "Status: %d", code)
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
	b.Assert(core.SG__().request_info.content_length >= core.SG__().read_post_bytes)
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
	var array_ptr *types.Zval = (*types.Zval)(arg)
	var filter_arg int = b.Cond(array_ptr.GetArr() == core.PG__().http_globals[core.TRACK_VARS_ENV].GetArr(), core.PARSE_ENV, core.PARSE_SERVER)
	var new_val_len int
	if core.SM__().GetInputFilter()(filter_arg, var_, &val, strlen(val), &new_val_len) != 0 {
		core.PhpRegisterVariableSafe(var_, val, new_val_len, array_ptr)
	}
}
func CgiPhpImportEnvironmentVariables(array_ptr *types.Zval) {
	if core.PG__().variables_order && (strchr(core.PG__().variables_order, 'E') || strchr(core.PG__().variables_order, 'e')) {
		if core.PG__().http_globals[core.TRACK_VARS_ENV].GetType() != types.IS_ARRAY {
			zend.ZendIsAutoGlobalStr("_ENV", b.SizeOf("\"_ENV\"")-1)
		}
		if core.PG__().http_globals[core.TRACK_VARS_ENV].GetType() == types.IS_ARRAY && array_ptr.GetArr() != core.PG__().http_globals[core.TRACK_VARS_ENV].GetArr() {
			array_ptr.GetArr().DestroyEx()
			array_ptr.SetArr(types.ZendArrayDup(core.PG__().http_globals[core.TRACK_VARS_ENV].GetArr()))
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
func SapiCgiRegisterVariables(track_vars_array *types.Zval) {
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

		if core.SM__().GetInputFilter()(core.PARSE_SERVER, "PHP_SELF", &php_self, php_self_len, &php_self_len) != 0 {
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
		if core.SM__().GetInputFilter()(core.PARSE_SERVER, "PHP_SELF", &php_self, php_self_len, &php_self_len) != 0 {
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
			log.Printf("%s\n", message)
		}
	} else {
		log.Printf("%s\n", message)
	}
}
func PhpCgiIniActivateUserConfig(path *byte, path_len int, doc_root *byte, doc_root_len int) {
	var new_entry *UserConfigCacheEntry
	var entry *UserConfigCacheEntry
	var request_time int64 = int64(core.SapiGetRequestTime())

	/* Find cached config entry: If not found, create one */

	if b.Assign(&entry, types.ZendHashStrFindPtr(&(CGIG(user_config_cache)), b.CastStr(path, path_len))) == nil {
		new_entry = zend.Pemalloc(b.SizeOf("user_config_cache_entry"), 1)
		new_entry.SetExpires(0)
		new_entry.SetUserConfig((*types.Array)(zend.Pemalloc(b.SizeOf("HashTable"), 1)))
		new_entry.GetUserConfig() = types.MakeArrayEx(8, types.DtorFuncT(core.ConfigZvalDtor), 1)
		entry = types.ZendHashUpdatePtr(&(CGIG(user_config_cache)), b.CastStr(path, path_len), new_entry)
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
				core.PhpParseUserIniFile(path, core.PG__().user_ini_filename, entry.GetUserConfig())
				*ptr = '/'
				ptr++
			}
		} else {
			core.PhpParseUserIniFile(path, core.PG__().user_ini_filename, entry.GetUserConfig())
		}
		if real_path != nil {
			zend.Efree(real_path)
		}
		entry.SetExpires(request_time + core.PG__().user_ini_cache_ttl)
	}

	/* Activate ini entries with values from the user config hash */

	core.PhpIniActivateConfig(entry.GetUserConfig(), core.PHP_INI_PERDIR, core.PHP_INI_STAGE_HTACCESS)

	/* Activate ini entries with values from the user config hash */
}
func SapiCgiActivate() int {
	/* PATH_TRANSLATED should be defined at this stage but better safe than sorry :) */

	if !(core.SG__().request_info.path_translated) {
		return types.FAILURE
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
	if core.PhpIniHasPerDirConfig() != 0 || core.PG__().user_ini_filename && *core.PG__().user_ini_filename {
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

		if core.PG__().user_ini_filename && *core.PG__().user_ini_filename {
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
	return types.SUCCESS
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
	return types.SUCCESS
}
func PhpCgiStartup(sapi_module *core.sapi_module_struct) int {
	if core.PhpModuleStartup(sapi_module, &CgiModuleEntry, 1) == types.FAILURE {
		return types.FAILURE
	}
	return types.SUCCESS
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
			if env_document_root == nil && core.PG__().doc_root {
				env_document_root = CGI_PUTENV("DOCUMENT_ROOT", core.PG__().doc_root)

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
	php_cgi_globals.GetUserConfigCache() = types.MakeArrayEx(8, UserConfigCacheEntryDtor, 1)
}
func ZmStartupCgi(type_ int, module_number int) int {
	zend.REGISTER_INI_ENTRIES(module_number)
	return types.SUCCESS
}
func ZmShutdownCgi(type_ int, module_number int) int {
	CGIG(user_config_cache).Destroy()
	zend.UNREGISTER_INI_ENTRIES(module_number)
	return types.SUCCESS
}
func ZmInfoCgi(zend_module *zend.ZendModuleEntry) { zend.DISPLAY_INI_ENTRIES() }
func ZifApacheRequestHeaders(executeData zpp.DefEx, return_value zpp.DefReturn) {
	if !executeData.CheckNumArgsNone(false) {
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
			zend.AddAssocStr(return_value, b.CastStr(var_, var_len), b.CastStrAuto(val))
		}
		if t != buf && t != nil {
			zend.Efree(t)
		}
	}
}
func AddResponseHeader(h *core.SapiHeader, return_value *types.Zval) {
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
				zend.AddAssocStringlEx(return_value, b.CastStr(s, len_), b.CastStr(p, h.GetHeaderLen()-(p-h.GetHeader())))
				zend.FreeAlloca(s, use_heap)
			}
		}
	}
}
func ZifApacheResponseHeaders(executeData zpp.DefEx, return_value zpp.DefReturn) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	zend.ArrayInit(return_value)
	zend.ZendLlistApplyWithArgument(core.SG__().sapi_headers.headers, zend.LlistApplyWithArgFuncT(AddResponseHeader), return_value)
}
