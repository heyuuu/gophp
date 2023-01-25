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

var CliShellCallbacks CliShellCallbacksT = CliShellCallbacksT{nil, nil, nil}
var HARDCODED_INI []byte = "html_errors=0\n" + "register_argc_argv=1\n" + "implicit_flush=1\n" + "output_buffering=0\n" + "max_execution_time=0\n" + "max_input_time=-1\n0"
var OPTIONS []core.Opt = []core.Opt{{'a', 0, "interactive"}, {'B', 1, "process-begin"}, {'C', 0, "no-chdir"}, {'c', 1, "php-ini"}, {'d', 1, "define"}, {'E', 1, "process-end"}, {'e', 0, "profile-info"}, {'F', 1, "process-file"}, {'f', 1, "file"}, {'h', 0, "help"}, {'i', 0, "info"}, {'l', 0, "syntax-check"}, {'m', 0, "modules"}, {'n', 0, "no-php-ini"}, {'q', 0, "no-header"}, {'R', 1, "process-code"}, {'H', 0, "hide-args"}, {'r', 1, "run"}, {'s', 0, "syntax-highlight"}, {'s', 0, "syntax-highlighting"}, {'S', 1, "server"}, {'t', 1, "docroot"}, {'w', 0, "strip"}, {'?', 0, "usage"}, {'v', 0, "version"}, {'z', 1, "zend-extension"}, {10, 1, "rf"}, {10, 1, "rfunction"}, {11, 1, "rc"}, {11, 1, "rclass"}, {12, 1, "re"}, {12, 1, "rextension"}, {13, 1, "rz"}, {13, 1, "rzendextension"}, {14, 1, "ri"}, {14, 1, "rextinfo"}, {15, 0, "ini"}, {'-', 0, nil}}

const STDOUT_FILENO = 1
const STDERR_FILENO = 2

var PhpSelf *byte = ""
var ScriptFilename *byte = ""
var CliSapiModule core.sapi_module_struct = core.sapi_module_struct{"cli", "Command Line Interface", PhpCliStartup, core.PhpModuleShutdownWrapper, nil, SapiCliDeactivate, SapiCliUbWrite, SapiCliFlush, nil, nil, core.PhpError, SapiCliHeaderHandler, SapiCliSendHeaders, SapiCliSendHeader, nil, SapiCliReadCookies, SapiCliRegisterVariables, SapiCliLogMessage, nil, nil, nil, nil, nil, nil, 0, 0, nil, nil, nil, nil, nil, nil, 0, nil, nil, nil}
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
	{
		"cli_set_process_title",
		ZifCliSetProcessTitle,
		ArginfoCliSetProcessTitle,
		uint32_t(b.SizeOf("arginfo_cli_set_process_title")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"cli_get_process_title",
		ZifCliGetProcessTitle,
		ArginfoCliGetProcessTitle,
		uint32_t(b.SizeOf("arginfo_cli_get_process_title")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

/* {{{ php_cli_usage
 */

var SInProcess *core.PhpStream = nil
var ParamModeConflict *byte = "Either execute direct code, process stdin or use a file.\n"

/* {{{ cli_seek_file_begin
 */

/*}}}*/
