// <<generate>>

package cli

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	"sik/zend"
)

// Source: <sapi/cli/php_cli.c>

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
   | Author: Edin Kadribasic <edink@php.net>                              |
   |         Marcus Boerger <helly@php.net>                               |
   |         Johannes Schlueter <johannes@php.net>                        |
   |         Parts based on CGI SAPI Module by                            |
   |         Rasmus Lerdorf, Stig Bakken and Zeev Suraski                 |
   +----------------------------------------------------------------------+
*/

// failed # include "ext/reflection/php_reflection.h"

var PhpIniOpenedPath *byte
var PhpIniScannedPath *byte
var PhpIniScannedFiles *byte

const O_BINARY = 0
const PHP_MODE_STANDARD = 1
const PHP_MODE_HIGHLIGHT = 2
const PHP_MODE_LINT = 4
const PHP_MODE_STRIP = 5
const PHP_MODE_CLI_DIRECT = 6
const PHP_MODE_PROCESS_STDIN = 7
const PHP_MODE_REFLECTION_FUNCTION = 8
const PHP_MODE_REFLECTION_CLASS = 9
const PHP_MODE_REFLECTION_EXTENSION = 10
const PHP_MODE_REFLECTION_EXT_INFO = 11
const PHP_MODE_REFLECTION_ZEND_EXTENSION = 12
const PHP_MODE_SHOW_INI_CONFIG = 13

var HARDCODED_INI string = "html_errors=0\n" + "register_argc_argv=1\n" + "implicit_flush=1\n" + "output_buffering=0\n" + "max_execution_time=0\n" + "max_input_time=-1\n0"
var OPTIONS []core.Opt = []core.Opt{
	core.MakeOpt('a', 0, "interactive"),
	core.MakeOpt('B', 1, "process-begin"),
	core.MakeOpt('C', 0, "no-chdir"),
	core.MakeOpt('c', 1, "php-ini"),
	core.MakeOpt('d', 1, "define"),
	core.MakeOpt('E', 1, "process-end"),
	core.MakeOpt('e', 0, "profile-info"),
	core.MakeOpt('F', 1, "process-file"),
	core.MakeOpt('f', 1, "file"),
	core.MakeOpt('h', 0, "help"),
	core.MakeOpt('i', 0, "info"),
	core.MakeOpt('l', 0, "syntax-check"),
	core.MakeOpt('m', 0, "modules"),
	core.MakeOpt('n', 0, "no-php-ini"),
	core.MakeOpt('q', 0, "no-header"),
	core.MakeOpt('R', 1, "process-code"),
	core.MakeOpt('H', 0, "hide-args"),
	core.MakeOpt('r', 1, "run"),
	core.MakeOpt('s', 0, "syntax-highlight"),
	core.MakeOpt('s', 0, "syntax-highlighting"),
	core.MakeOpt('S', 1, "server"),
	core.MakeOpt('t', 1, "docroot"),
	core.MakeOpt('w', 0, "strip"),
	core.MakeOpt('?', 0, "usage"),
	core.MakeOpt('v', 0, "version"),
	core.MakeOpt('z', 1, "zend-extension"),
	core.MakeOpt(10, 1, "rf"),
	core.MakeOpt(10, 1, "rfunction"),
	core.MakeOpt(11, 1, "rc"),
	core.MakeOpt(11, 1, "rclass"),
	core.MakeOpt(12, 1, "re"),
	core.MakeOpt(12, 1, "rextension"),
	core.MakeOpt(13, 1, "rz"),
	core.MakeOpt(13, 1, "rzendextension"),
	core.MakeOpt(14, 1, "ri"),
	core.MakeOpt(14, 1, "rextinfo"),
	core.MakeOpt(15, 0, "ini"),
}

const STDOUT_FILENO = 1
const STDERR_FILENO = 2

var PhpSelf *byte = ""
var ScriptFilename *byte = ""
var CliSapiModule core.sapi_module_struct = core.MakeSapiModule("cli", "Command Line Interface", PhpCliStartup, core.PhpModuleShutdownWrapper, nil, SapiCliDeactivate, SapiCliUbWrite, SapiCliFlush, nil, nil, core.PhpError, SapiCliHeaderHandler, SapiCliSendHeaders, SapiCliSendHeader, nil, SapiCliReadCookies, SapiCliRegisterVariables, SapiCliLogMessage, nil, nil, nil, nil, nil, nil, 0, 0, nil, nil, nil, nil, nil, nil, 0, nil, nil, nil)
var ArginfoDl []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	zend.MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0),
	zend.MakeZendInternalArgInfo("extension_filename", 0, 0, 0),
}
var AdditionalFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("dl", standard.ZifDl, ArginfoDl, uint32(b.SizeOf("arginfo_dl")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("cli_set_process_title", ZifCliSetProcessTitle, ArginfoCliSetProcessTitle, uint32(b.SizeOf("arginfo_cli_set_process_title")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry("cli_get_process_title", ZifCliGetProcessTitle, ArginfoCliGetProcessTitle, uint32(b.SizeOf("arginfo_cli_get_process_title")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}

/* {{{ php_cli_usage
 */

var SInProcess *core.PhpStream = nil
var ParamModeConflict *byte = "Either execute direct code, process stdin or use a file.\n"

/* {{{ cli_seek_file_begin
 */

/*}}}*/
