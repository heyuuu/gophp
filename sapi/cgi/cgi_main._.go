// <<generate>>

package cgi

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
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

const PHP_MODE_STANDARD = 1
const PHP_MODE_HIGHLIGHT = 2
const PHP_MODE_LINT = 4
const PHP_MODE_STRIP = 5

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

var php_cgi_globals php_cgi_globals_struct

const STDOUT_FILENO = 1
const SAPI_CGI_MAX_HEADER_LENGTH = 1024
const STDIN_FILENO = 0

/* {{{ php_cgi_ini_activate_user_config
 */

/* {{{ sapi_module_struct cgi_sapi_module
 */

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
		uint32(b.SizeOf("arginfo_dl")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

/* {{{ php_cgi_usage
 */

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

/**
 * Clean up child processes upon exit
 */

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

/* {{{ php_cgi_globals_ctor
 */

var ArginfoNoArgs []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var CgiFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"apache_child_terminate",
		ZifApacheChildTerminate,
		ArginfoNoArgs,
		uint32(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info") - 1),
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

/* {{{ main
 */
