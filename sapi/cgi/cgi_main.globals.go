// <<generate>>

package cgi

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	"sik/zend"
)

var Act __struct__sigaction
var OldTerm __struct__sigaction
var OldQuit __struct__sigaction
var OldInt __struct__sigaction
var PhpPhpImportEnvironmentVariables func(array_ptr *zend.Zval)
var Children int = 0
var Parent int = 1
var ExitSignal int = 0
var ParentWaiting int = 0
var Pgroup pid_t

const PHP_MODE_STANDARD = 1
const PHP_MODE_HIGHLIGHT = 2
const PHP_MODE_LINT = 4
const PHP_MODE_STRIP = 5

var PhpOptarg *byte = nil
var PhpOptind int = 1
var OPTIONS []core.Opt = []core.Opt{{'a', 0, "interactive"}, {'b', 1, "bindpath"}, {'C', 0, "no-chdir"}, {'c', 1, "php-ini"}, {'d', 1, "define"}, {'e', 0, "profile-info"}, {'f', 1, "file"}, {'h', 0, "help"}, {'i', 0, "info"}, {'l', 0, "syntax-check"}, {'m', 0, "modules"}, {'n', 0, "no-php-ini"}, {'q', 0, "no-header"}, {'s', 0, "syntax-highlight"}, {'s', 0, "syntax-highlighting"}, {'w', 0, "strip"}, {'?', 0, "usage"}, {'v', 0, "version"}, {'z', 1, "zend-extension"}, {'T', 1, "timing"}, {'-', 0, nil}}

type _phpCgiGlobals = php_cgi_globals_struct

var php_cgi_globals php_cgi_globals_struct

const STDOUT_FILENO = 1
const SAPI_CGI_MAX_HEADER_LENGTH = 1024
const STDIN_FILENO = 0

var CgiSapiModule core.sapi_module_struct = core.sapi_module_struct{"cgi-fcgi", "CGI/FastCGI", PhpCgiStartup, core.PhpModuleShutdownWrapper, SapiCgiActivate, SapiCgiDeactivate, SapiCgiUbWrite, SapiCgiFlush, nil, SapiCgiGetenv, core.PhpError, nil, SapiCgiSendHeaders, nil, SapiCgiReadPost, SapiCgiReadCookies, SapiCgiRegisterVariables, SapiCgiLogMessage, nil, nil, nil, nil, nil, nil, 0, 0, nil, nil, nil, nil, nil, nil, 0, nil, nil, nil}
var ArginfoDl []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"extension_filename", 0, 0, 0},
}
var AdditionalFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"dl",
		standard.ZifDl,
		ArginfoDl,
		uint32_t(b.SizeOf("arginfo_dl")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
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
		b.SizeOf("\"0\"") - 1,
		b.SizeOf("\"cgi.rfc2616_headers\"") - 1,
		core.PHP_INI_ALL,
	},
	{
		"cgi.nph",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetNph())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		"0",
		nil,
		b.SizeOf("\"0\"") - 1,
		b.SizeOf("\"cgi.nph\"") - 1,
		core.PHP_INI_ALL,
	},
	{
		"cgi.check_shebang_line",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetCheckShebangLine())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		"1",
		nil,
		b.SizeOf("\"1\"") - 1,
		b.SizeOf("\"cgi.check_shebang_line\"") - 1,
		core.PHP_INI_SYSTEM,
	},
	{
		"cgi.force_redirect",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetForceRedirect())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		"1",
		nil,
		b.SizeOf("\"1\"") - 1,
		b.SizeOf("\"cgi.force_redirect\"") - 1,
		core.PHP_INI_SYSTEM,
	},
	{
		"cgi.redirect_status_env",
		zend.OnUpdateString,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetRedirectStatusEnv())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		nil,
		nil,
		b.SizeOf("NULL") - 1,
		b.SizeOf("\"cgi.redirect_status_env\"") - 1,
		core.PHP_INI_SYSTEM,
	},
	{
		"cgi.fix_pathinfo",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetFixPathinfo())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		"1",
		nil,
		b.SizeOf("\"1\"") - 1,
		b.SizeOf("\"cgi.fix_pathinfo\"") - 1,
		core.PHP_INI_SYSTEM,
	},
	{
		"cgi.discard_path",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetDiscardPath())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		"0",
		nil,
		b.SizeOf("\"0\"") - 1,
		b.SizeOf("\"cgi.discard_path\"") - 1,
		core.PHP_INI_SYSTEM,
	},
	{
		"fastcgi.logging",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*php_cgi_globals_struct)(nil).GetFcgiLogging())) - (*byte)(nil))),
		any(&php_cgi_globals),
		nil,
		"1",
		nil,
		b.SizeOf("\"1\"") - 1,
		b.SizeOf("\"fastcgi.logging\"") - 1,
		core.PHP_INI_SYSTEM,
	},
	{nil, nil, nil, nil, nil, nil, nil, 0, 0, 0},
}
var ArginfoNoArgs []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var CgiFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"apache_child_terminate",
		ZifApacheChildTerminate,
		ArginfoNoArgs,
		uint32_t(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"apache_request_headers",
		ZifApacheRequestHeaders,
		ArginfoNoArgs,
		uint32_t(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"apache_response_headers",
		ZifApacheResponseHeaders,
		ArginfoNoArgs,
		uint32_t(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getallheaders",
		ZifApacheRequestHeaders,
		ArginfoNoArgs,
		uint32_t(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}
var CgiModuleEntry zend.ZendModuleEntry = zend.ZendModuleEntry{
	b.SizeOf("zend_module_entry"),
	zend.ZEND_MODULE_API_NO,
	core.ZEND_DEBUG,
	zend.USING_ZTS,
	nil,
	nil,
	"cgi-fcgi",
	CgiFunctions,
	ZmStartupCgi,
	ZmShutdownCgi,
	nil,
	nil,
	ZmInfoCgi,
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
