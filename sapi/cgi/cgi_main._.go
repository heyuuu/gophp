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
var OPTIONS []core.Opt = []core.Opt{
	core.MakeOpt('a', 0, "interactive"),
	core.MakeOpt('b', 1, "bindpath"),
	core.MakeOpt('C', 0, "no-chdir"),
	core.MakeOpt('c', 1, "php-ini"),
	core.MakeOpt('d', 1, "define"),
	core.MakeOpt('e', 0, "profile-info"),
	core.MakeOpt('f', 1, "file"),
	core.MakeOpt('h', 0, "help"),
	core.MakeOpt('i', 0, "info"),
	core.MakeOpt('l', 0, "syntax-check"),
	core.MakeOpt('m', 0, "modules"),
	core.MakeOpt('n', 0, "no-php-ini"),
	core.MakeOpt('q', 0, "no-header"),
	core.MakeOpt('s', 0, "syntax-highlight"),
	core.MakeOpt('s', 0, "syntax-highlighting"),
	core.MakeOpt('w', 0, "strip"),
	core.MakeOpt('?', 0, "usage"),
	core.MakeOpt('v', 0, "version"),
	core.MakeOpt('z', 1, "zend-extension"),
	core.MakeOpt('T', 1, "timing"),
	core.MakeOpt('-', 0, nil),
}

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

var CgiSapiModule = core.MakeSapiModule("cgi-fcgi", "CGI/FastCGI", PhpCgiStartup, core.PhpModuleShutdownWrapper, SapiCgiActivate, SapiCgiDeactivate, SapiCgiUbWrite, SapiCgiFlush, SapiCgiGetenv, nil, SapiCgiSendHeaders, nil, SapiCgiReadPost, SapiCgiReadCookies, SapiCgiRegisterVariables, SapiCgiLogMessage)
var ArginfoDl = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("extension_filename", 0, 0, 0),
}
var AdditionalFunctions = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("dl", standard.ZifDl, ArginfoDl, uint32(b.SizeOf("arginfo_dl")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}

var ArginfoNoArgs = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
}
var CgiFunctions = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("apache_child_terminate", ZifApacheChildTerminate, ArginfoNoArgs, uint32(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("apache_request_headers", ZifApacheRequestHeaders, ArginfoNoArgs, uint32(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("apache_response_headers", ZifApacheResponseHeaders, ArginfoNoArgs, uint32(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("getallheaders", ZifApacheRequestHeaders, ArginfoNoArgs, uint32(b.SizeOf("arginfo_no_args")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var CgiModuleEntry = zend.MakeZendModuleEntry(b.SizeOf("zend_module_entry"), zend.ZEND_MODULE_API_NO, 0, zend.USING_ZTS, nil, nil, "cgi-fcgi", CgiFunctions, ZmStartupCgi, ZmShutdownCgi, nil, nil, ZmInfoCgi, core.PHP_VERSION, 0, nil, nil, nil, nil, 0, 0, nil, 0, "API"+"ZEND_MODULE_API_NO"+zend.ZEND_BUILD_TS)
