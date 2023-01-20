// <<generate>>

package cli

import (
	"sik/core"
	"sik/core/streams"
	"sik/ext/standard"
	g "sik/runtime/grammar"
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

// # include "php.h"

// # include "php_globals.h"

// # include "php_variables.h"

// # include "zend_hash.h"

// # include "zend_modules.h"

// # include "zend_interfaces.h"

// failed # include "ext/reflection/php_reflection.h"

// # include "SAPI.h"

// # include < stdio . h >

// # include "php.h"

// # include < sys / time . h >

// # include < unistd . h >

// # include < signal . h >

// # include < locale . h >

// # include "zend.h"

// # include "zend_extensions.h"

// # include "php_ini.h"

// # include "php_globals.h"

// # include "php_main.h"

// # include "fopen_wrappers.h"

// # include "ext/standard/php_standard.h"

// # include "cli.h"

// # include "zend_compile.h"

// # include "zend_execute.h"

// # include "zend_highlight.h"

// # include "zend_exceptions.h"

// # include "php_getopt.h"

// # include "php_cli_server.h"

// # include "ps_title.h"

// # include "php_cli_process_title.h"

// #define php_select(m,r,w,e,t) select ( m , r , w , e , t )

var PhpIniOpenedPath *byte
var PhpIniScannedPath *byte
var PhpIniScannedFiles *byte

// #define O_BINARY       0

// #define PHP_MODE_STANDARD       1

// #define PHP_MODE_HIGHLIGHT       2

// #define PHP_MODE_LINT       4

// #define PHP_MODE_STRIP       5

// #define PHP_MODE_CLI_DIRECT       6

// #define PHP_MODE_PROCESS_STDIN       7

// #define PHP_MODE_REFLECTION_FUNCTION       8

// #define PHP_MODE_REFLECTION_CLASS       9

// #define PHP_MODE_REFLECTION_EXTENSION       10

// #define PHP_MODE_REFLECTION_EXT_INFO       11

// #define PHP_MODE_REFLECTION_ZEND_EXTENSION       12

// #define PHP_MODE_SHOW_INI_CONFIG       13

var CliShellCallbacks CliShellCallbacksT = CliShellCallbacksT{nil, nil, nil}

func PhpCliGetShellCallbacks() *CliShellCallbacksT { return &CliShellCallbacks }

var HARDCODED_INI []byte = "html_errors=0\n" + "register_argc_argv=1\n" + "implicit_flush=1\n" + "output_buffering=0\n" + "max_execution_time=0\n" + "max_input_time=-1\n0"
var OPTIONS []core.Opt = []core.Opt{{'a', 0, "interactive"}, {'B', 1, "process-begin"}, {'C', 0, "no-chdir"}, {'c', 1, "php-ini"}, {'d', 1, "define"}, {'E', 1, "process-end"}, {'e', 0, "profile-info"}, {'F', 1, "process-file"}, {'f', 1, "file"}, {'h', 0, "help"}, {'i', 0, "info"}, {'l', 0, "syntax-check"}, {'m', 0, "modules"}, {'n', 0, "no-php-ini"}, {'q', 0, "no-header"}, {'R', 1, "process-code"}, {'H', 0, "hide-args"}, {'r', 1, "run"}, {'s', 0, "syntax-highlight"}, {'s', 0, "syntax-highlighting"}, {'S', 1, "server"}, {'t', 1, "docroot"}, {'w', 0, "strip"}, {'?', 0, "usage"}, {'v', 0, "version"}, {'z', 1, "zend-extension"}, {10, 1, "rf"}, {10, 1, "rfunction"}, {11, 1, "rc"}, {11, 1, "rclass"}, {12, 1, "re"}, {12, 1, "rextension"}, {13, 1, "rz"}, {13, 1, "rzendextension"}, {14, 1, "ri"}, {14, 1, "rextinfo"}, {15, 0, "ini"}, {'-', 0, nil}}

func ModuleNameCmp(a any, b any) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	return strcasecmp((*zend.ZendModuleEntry)(f.val.value.ptr).name, (*zend.ZendModuleEntry)(s.val.value.ptr).name)
}

/* }}} */

func PrintModules() {
	var sorted_registry zend.HashTable
	var module *zend.ZendModuleEntry
	zend._zendHashInit(&sorted_registry, 50, nil, 0)
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

/* }}} */

func PrintExtensionInfo(ext *zend.ZendExtension, arg any) int {
	core.PhpPrintf("%s\n", ext.name)
	return 0
}

/* }}} */

func ExtensionNameCmp(f **zend.ZendLlistElement, s **zend.ZendLlistElement) int {
	var fe *zend.ZendExtension = (*zend.ZendExtension)((*f).data)
	var se *zend.ZendExtension = (*zend.ZendExtension)((*s).data)
	return strcmp(fe.name, se.name)
}

/* }}} */

func PrintExtensions() {
	var sorted_exts zend.ZendLlist
	zend.ZendLlistCopy(&sorted_exts, &zend.ZendExtensions)
	sorted_exts.dtor = nil
	zend.ZendLlistSort(&sorted_exts, ExtensionNameCmp)
	zend.ZendLlistApply(&sorted_exts, zend.LlistApplyFuncT(PrintExtensionInfo))
	zend.ZendLlistDestroy(&sorted_exts)
}

/* }}} */

// #define STDOUT_FILENO       1

// #define STDERR_FILENO       2

func SapiCliSelect(fd core.PhpSocketT) int {
	var wfd fd_set
	var tv __struct__timeval
	var ret int
	FD_ZERO(&wfd)
	if fd < FD_SETSIZE {
		FD_SET(fd, &wfd)
	}
	tv.tv_sec = long(standard.FileGlobals.default_socket_timeout)
	tv.tv_usec = 0
	ret = select_(fd+1, nil, &wfd, nil, &tv)
	return ret != -1
}
func SapiCliSingleWrite(str *byte, str_length int) ssize_t {
	var ret ssize_t
	if CliShellCallbacks.GetCliShellWrite() != nil {
		CliShellCallbacks.GetCliShellWrite()(str, str_length)
	}
	for {
		ret = write(1, str, str_length)
		if !(ret <= 0 && errno == EAGAIN && SapiCliSelect(1) != 0) {
			break
		}
	}
	return ret
}

/* }}} */

func SapiCliUbWrite(str *byte, str_length int) int {
	var ptr *byte = str
	var remaining int = str_length
	var ret ssize_t
	if str_length == 0 {
		return 0
	}
	if CliShellCallbacks.GetCliShellUbWrite() != nil {
		var ub_wrote int
		ub_wrote = CliShellCallbacks.GetCliShellUbWrite()(str, str_length)
		if ub_wrote != size_t-1 {
			return ub_wrote
		}
	}
	for remaining > 0 {
		ret = SapiCliSingleWrite(ptr, remaining)
		if ret < 0 {
			zend.EG.exit_status = 255
			core.PhpHandleAbortedConnection()
			break
		}
		ptr += ret
		remaining -= ret
	}
	return ptr - str
}

/* }}} */

func SapiCliFlush(server_context any) {
	/* Ignore EBADF here, it's caused by the fact that STDIN/STDOUT/STDERR streams
	 * are/could be closed before fflush() is called.
	 */

	if fflush(stdout) == EOF && errno != EBADF {
		core.PhpHandleAbortedConnection()
	}

	/* Ignore EBADF here, it's caused by the fact that STDIN/STDOUT/STDERR streams
	 * are/could be closed before fflush() is called.
	 */
}

/* }}} */

var PhpSelf *byte = ""
var ScriptFilename *byte = ""

func SapiCliRegisterVariables(track_vars_array *zend.Zval) {
	var len_ int
	var docroot *byte = ""

	/* In CGI mode, we consider the environment to be a part of the server
	 * variables
	 */

	core.PhpImportEnvironmentVariables(track_vars_array)

	/* Build the special-case PHP_SELF variable for the CLI version */

	len_ = strlen(PhpSelf)
	if core.sapi_module.input_filter(5, "PHP_SELF", &PhpSelf, len_, &len_) != 0 {
		core.PhpRegisterVariable("PHP_SELF", PhpSelf, track_vars_array)
	}
	if core.sapi_module.input_filter(5, "SCRIPT_NAME", &PhpSelf, len_, &len_) != 0 {
		core.PhpRegisterVariable("SCRIPT_NAME", PhpSelf, track_vars_array)
	}

	/* filenames are empty for stdin */

	len_ = strlen(ScriptFilename)
	if core.sapi_module.input_filter(5, "SCRIPT_FILENAME", &ScriptFilename, len_, &len_) != 0 {
		core.PhpRegisterVariable("SCRIPT_FILENAME", ScriptFilename, track_vars_array)
	}
	if core.sapi_module.input_filter(5, "PATH_TRANSLATED", &ScriptFilename, len_, &len_) != 0 {
		core.PhpRegisterVariable("PATH_TRANSLATED", ScriptFilename, track_vars_array)
	}

	/* just make it available */

	len_ = 0
	if core.sapi_module.input_filter(5, "DOCUMENT_ROOT", &docroot, len_, &len_) != 0 {
		core.PhpRegisterVariable("DOCUMENT_ROOT", docroot, track_vars_array)
	}
}

/* }}} */

func SapiCliLogMessage(message *byte, syslog_type_int int) { fprintf(stderr, "%s\n", message) }

/* }}} */

func SapiCliDeactivate() int {
	fflush(stdout)
	if core.sapi_globals.request_info.argv0 != nil {
		zend.Free(core.sapi_globals.request_info.argv0)
		core.sapi_globals.request_info.argv0 = nil
	}
	return zend.SUCCESS
}

/* }}} */

func SapiCliReadCookies() *byte { return nil }

/* }}} */

func SapiCliHeaderHandler(h *core.SapiHeader, op core.SapiHeaderOpEnum, s *core.SapiHeaders) int {
	return 0
}

/* }}} */

func SapiCliSendHeaders(sapi_headers *core.SapiHeaders) int {
	/* We do nothing here, this function is needed to prevent that the fallback
	 * header handling is called. */

	return 1

	/* We do nothing here, this function is needed to prevent that the fallback
	 * header handling is called. */
}

/* }}} */

func SapiCliSendHeader(sapi_header *core.SapiHeader, server_context any) {}

/* }}} */

func PhpCliStartup(sapi_module *core.sapi_module_struct) int {
	if core.PhpModuleStartup(sapi_module, nil, 0) == zend.FAILURE {
		return zend.FAILURE
	}
	return zend.SUCCESS
}

/* }}} */

// #define INI_DEFAULT(name,value) ZVAL_NEW_STR ( & tmp , zend_string_init ( value , sizeof ( value ) - 1 , 1 ) ) ; zend_hash_str_update ( configuration_hash , name , sizeof ( name ) - 1 , & tmp ) ;

func SapiCliIniDefaults(configuration_hash *zend.HashTable) {
	var tmp zend.Zval
	var __z *zend.Zval = &tmp
	var __s *zend.ZendString = zend.ZendStringInit("0", g.SizeOf("\"0\"")-1, 1)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	zend.ZendHashStrUpdate(configuration_hash, "report_zend_debug", g.SizeOf("\"report_zend_debug\"")-1, &tmp)
	var __z *zend.Zval = &tmp
	var __s *zend.ZendString = zend.ZendStringInit("1", g.SizeOf("\"1\"")-1, 1)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	zend.ZendHashStrUpdate(configuration_hash, "display_errors", g.SizeOf("\"display_errors\"")-1, &tmp)
}

/* }}} */

var CliSapiModule core.sapi_module_struct = core.sapi_module_struct{"cli", "Command Line Interface", PhpCliStartup, core.PhpModuleShutdownWrapper, nil, SapiCliDeactivate, SapiCliUbWrite, SapiCliFlush, nil, nil, zend.ZendError, SapiCliHeaderHandler, SapiCliSendHeaders, SapiCliSendHeader, nil, SapiCliReadCookies, SapiCliRegisterVariables, SapiCliLogMessage, nil, nil, nil, nil, nil, nil, 0, 0, nil, nil, nil, nil, nil, nil, 0, nil, nil, nil}

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
	{nil, nil, nil, 0, 0},
}

/* {{{ php_cli_usage
 */

func PhpCliUsage(argv0 *byte) {
	var prog *byte
	prog = strrchr(argv0, '/')
	if prog != nil {
		prog++
	} else {
		prog = "php"
	}
	printf("Usage: %s [options] [-f] <file> [--] [args...]\n"+"   %s [options] -r <code> [--] [args...]\n"+"   %s [options] [-B <begin_code>] -R <code> [-E <end_code>] [--] [args...]\n"+"   %s [options] [-B <begin_code>] -F <file> [-E <end_code>] [--] [args...]\n"+"   %s [options] -S <addr>:<port> [-t docroot] [router]\n"+"   %s [options] -- [args...]\n"+"   %s [options] -a\n"+"\n"+"  -a               Run interactively\n"+"  -c <path>|<file> Look for php.ini file in this directory\n"+"  -n               No configuration (ini) files will be used\n"+"  -d foo[=bar]     Define INI entry foo with value 'bar'\n"+"  -e               Generate extended information for debugger/profiler\n"+"  -f <file>        Parse and execute <file>.\n"+"  -h               This help\n"+"  -i               PHP information\n"+"  -l               Syntax check only (lint)\n"+"  -m               Show compiled in modules\n"+"  -r <code>        Run PHP <code> without using script tags <?..?>\n"+"  -B <begin_code>  Run PHP <begin_code> before processing input lines\n"+"  -R <code>        Run PHP <code> for every input line\n"+"  -F <file>        Parse and execute <file> for every input line\n"+"  -E <end_code>    Run PHP <end_code> after processing all input lines\n"+"  -H               Hide any passed arguments from external tools.\n"+"  -S <addr>:<port> Run with built-in web server.\n"+"  -t <docroot>     Specify document root <docroot> for built-in web server.\n"+"  -s               Output HTML syntax highlighted source.\n"+"  -v               Version number\n"+"  -w               Output source with stripped comments and whitespace.\n"+"  -z <file>        Load Zend extension <file>.\n"+"\n"+"  args...          Arguments passed to script. Use -- args when first argument\n"+"                   starts with - or script is read from stdin\n"+"\n"+"  --ini            Show configuration file names\n"+"\n"+"  --rf <name>      Show information about function <name>.\n"+"  --rc <name>      Show information about class <name>.\n"+"  --re <name>      Show information about extension <name>.\n"+"  --rz <name>      Show information about Zend extension <name>.\n"+"  --ri <name>      Show configuration for extension <name>.\n"+"\n", prog, prog, prog, prog, prog, prog, prog)
}

/* }}} */

var SInProcess *core.PhpStream = nil

func CliRegisterFileHandles() {
	var s_in *core.PhpStream
	var s_out *core.PhpStream
	var s_err *core.PhpStream
	var sc_in *core.PhpStreamContext = nil
	var sc_out *core.PhpStreamContext = nil
	var sc_err *core.PhpStreamContext = nil
	var ic zend.ZendConstant
	var oc zend.ZendConstant
	var ec zend.ZendConstant
	s_in = streams._phpStreamOpenWrapperEx("php://stdin", "rb", 0, nil, sc_in)
	s_out = streams._phpStreamOpenWrapperEx("php://stdout", "wb", 0, nil, sc_out)
	s_err = streams._phpStreamOpenWrapperEx("php://stderr", "wb", 0, nil, sc_err)
	if s_in == nil || s_out == nil || s_err == nil {
		if s_in != nil {
			streams._phpStreamFree(s_in, 1|2)
		}
		if s_out != nil {
			streams._phpStreamFree(s_out, 1|2)
		}
		if s_err != nil {
			streams._phpStreamFree(s_err, 1|2)
		}
		return
	}
	SInProcess = s_in
	var __z *zend.Zval = &ic.value
	__z.value.res = s_in.res
	__z.u1.type_info = 9 | 1<<0<<8
	s_in.__exposed = 1
	var __z *zend.Zval = &oc.value
	__z.value.res = s_out.res
	__z.u1.type_info = 9 | 1<<0<<8
	s_out.__exposed = 1
	var __z *zend.Zval = &ec.value
	__z.value.res = s_err.res
	__z.u1.type_info = 9 | 1<<0<<8
	s_err.__exposed = 1
	&ic.value.u2.constant_flags = 1<<0&0xff | 0<<8
	ic.name = zend.ZendStringInitInterned("STDIN", g.SizeOf("\"STDIN\"")-1, 0)
	zend.ZendRegisterConstant(&ic)
	&oc.value.u2.constant_flags = 1<<0&0xff | 0<<8
	oc.name = zend.ZendStringInitInterned("STDOUT", g.SizeOf("\"STDOUT\"")-1, 0)
	zend.ZendRegisterConstant(&oc)
	&ec.value.u2.constant_flags = 1<<0&0xff | 0<<8
	ec.name = zend.ZendStringInitInterned("STDERR", g.SizeOf("\"STDERR\"")-1, 0)
	zend.ZendRegisterConstant(&ec)
}

/* }}} */

var ParamModeConflict *byte = "Either execute direct code, process stdin or use a file.\n"

/* {{{ cli_seek_file_begin
 */

func CliSeekFileBegin(file_handle *zend.ZendFileHandle, script_file *byte) int {
	var fp *FILE = fopen(script_file, "rb")
	if fp == nil {
		core.PhpPrintf("Could not open input file: %s\n", script_file)
		return zend.FAILURE
	}
	zend.ZendStreamInitFp(file_handle, fp, script_file)
	return zend.SUCCESS
}

/* }}} */

/*}}}*/

func DoCli(argc int, argv **byte) int {
	var c int
	var file_handle zend.ZendFileHandle
	var behavior int = 1
	var reflection_what *byte = nil
	var request_started int = 0
	var exit_status int = 0
	var php_optarg *byte = nil
	var orig_optarg *byte = nil
	var php_optind int = 1
	var orig_optind int = 1
	var exec_direct *byte = nil
	var exec_run *byte = nil
	var exec_begin *byte = nil
	var exec_end *byte = nil
	var arg_free *byte = nil
	var arg_excp **byte = &arg_free
	var script_file *byte = nil
	var translated_path *byte = nil
	var interactive int = 0
	var param_error *byte = nil
	var hide_argv int = 0
	var __orig_bailout *sigjmp_buf = zend.EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		zend.CG.in_compilation = 0
		for g.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &php_optarg, &php_optind, 0, 2)) != -1 {
			switch c {
			case 'i':
				if core.PhpRequestStartup() == zend.FAILURE {
					goto err
				}
				request_started = 1
				standard.PhpPrintInfo(0xffffffff & ^(1 << 1))
				core.PhpOutputEndAll()
				exit_status = c == '?' && argc > 1 && !(strchr(argv[1], c))
				goto out
			case 'v':
				core.PhpPrintf("PHP %s (%s) (built: %s %s) ( %s)\nCopyright (c) The PHP Group\n%s", "7.4.33", CliSapiModule.name, __DATE__, __TIME__, "NTS ", zend.GetZendVersion())
				core.SapiDeactivate()
				goto out
			case 'm':
				if core.PhpRequestStartup() == zend.FAILURE {
					goto err
				}
				request_started = 1
				core.PhpPrintf("[PHP Modules]\n")
				PrintModules()
				core.PhpPrintf("\n[Zend Modules]\n")
				PrintExtensions()
				core.PhpPrintf("\n")
				core.PhpOutputEndAll()
				exit_status = 0
				goto out
			default:
				break
			}
		}

		/* Set some CLI defaults */

		core.sapi_globals.options |= 1
		php_optind = orig_optind
		php_optarg = orig_optarg
		for g.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &php_optarg, &php_optind, 0, 2)) != -1 {
			switch c {
			case 'a':
				if interactive == 0 {
					if behavior != 1 {
						param_error = ParamModeConflict
						break
					}
					interactive = 1
				}
				break
			case 'C':

				/* This is default so NOP */

				break
			case 'F':
				if behavior == 7 {
					if exec_run != nil || script_file != nil {
						param_error = "You can use -R or -F only once.\n"
						break
					}
				} else if behavior != 1 {
					param_error = ParamModeConflict
					break
				}
				behavior = 7
				script_file = php_optarg
				break
			case 'f':
				if behavior == 6 || behavior == 7 {
					param_error = ParamModeConflict
					break
				} else if script_file != nil {
					param_error = "You can use -f only once.\n"
					break
				}
				script_file = php_optarg
				break
			case 'l':
				if behavior != 1 {
					break
				}
				behavior = 4
				break
			case 'q':

				/* This is default so NOP */

				break
			case 'r':
				if behavior == 6 {
					if exec_direct != nil || script_file != nil {
						param_error = "You can use -r only once.\n"
						break
					}
				} else if behavior != 1 || interactive != 0 {
					param_error = ParamModeConflict
					break
				}
				behavior = 6
				exec_direct = php_optarg
				break
			case 'R':
				if behavior == 7 {
					if exec_run != nil || script_file != nil {
						param_error = "You can use -R or -F only once.\n"
						break
					}
				} else if behavior != 1 {
					param_error = ParamModeConflict
					break
				}
				behavior = 7
				exec_run = php_optarg
				break
			case 'B':
				if behavior == 7 {
					if exec_begin != nil {
						param_error = "You can use -B only once.\n"
						break
					}
				} else if behavior != 1 || interactive != 0 {
					param_error = ParamModeConflict
					break
				}
				behavior = 7
				exec_begin = php_optarg
				break
			case 'E':
				if behavior == 7 {
					if exec_end != nil {
						param_error = "You can use -E only once.\n"
						break
					}
				} else if behavior != 1 || interactive != 0 {
					param_error = ParamModeConflict
					break
				}
				behavior = 7
				exec_end = php_optarg
				break
			case 's':
				if behavior == 6 || behavior == 7 {
					param_error = "Source highlighting only works for files.\n"
					break
				}
				behavior = 2
				break
			case 'w':
				if behavior == 6 || behavior == 7 {
					param_error = "Source stripping only works for files.\n"
					break
				}
				behavior = 5
				break
			case 'z':
				zend.ZendLoadExtension(php_optarg)
				break
			case 'H':
				hide_argv = 1
				break
			case 10:
				behavior = 8
				reflection_what = php_optarg
				break
			case 11:
				behavior = 9
				reflection_what = php_optarg
				break
			case 12:
				behavior = 10
				reflection_what = php_optarg
				break
			case 13:
				behavior = 12
				reflection_what = php_optarg
				break
			case 14:
				behavior = 11
				reflection_what = php_optarg
				break
			case 15:
				behavior = 13
				break
			default:
				break
			}
		}
		if param_error != nil {
			var __str *byte = param_error
			core.PhpOutputWrite(__str, strlen(__str))
			exit_status = 1
			goto err
		}
		if interactive != 0 {
			printf("Interactive mode enabled\n\n")
			fflush(stdout)
		}

		/* only set script_file if not set already and not in direct mode and not at end of parameter list */

		if argc > php_optind && script_file == nil && behavior != 6 && behavior != 7 && strcmp(argv[php_optind-1], "--") {
			script_file = argv[php_optind]
			php_optind++
		}
		if script_file != nil {
			if CliSeekFileBegin(&file_handle, script_file) != zend.SUCCESS {
				goto err
			} else {
				var real_path []byte
				if zend.TsrmRealpath(script_file, real_path) != nil {
					translated_path = strdup(real_path)
				}
				ScriptFilename = script_file
			}
		} else {

			/* We could handle PHP_MODE_PROCESS_STDIN in a different manner  */

			zend.ZendStreamInitFp(&file_handle, stdin, "Standard input code")

			/* We could handle PHP_MODE_PROCESS_STDIN in a different manner  */

		}
		PhpSelf = (*byte)(file_handle.filename)

		/* before registering argv to module exchange the *new* argv[0] */

		core.sapi_globals.request_info.argc = argc - php_optind + 1
		arg_excp = argv + php_optind - 1
		arg_free = argv[php_optind-1]
		if translated_path != nil {
			core.sapi_globals.request_info.path_translated = translated_path
		} else {
			core.sapi_globals.request_info.path_translated = (*byte)(file_handle.filename)
		}
		argv[php_optind-1] = (*byte)(file_handle.filename)
		core.sapi_globals.request_info.argv = argv + php_optind - 1
		if core.PhpRequestStartup() == zend.FAILURE {
			*arg_excp = arg_free
			fclose(file_handle.handle.fp)
			var __str *byte = "Could not startup.\n"
			core.PhpOutputWrite(__str, strlen(__str))
			goto err
		}
		request_started = 1
		zend.CG.skip_shebang = 1
		zend.ZendRegisterBoolConstant("PHP_CLI_PROCESS_TITLE", g.SizeOf("\"PHP_CLI_PROCESS_TITLE\"")-1, IsPsTitleAvailable() == 0, 1<<0, 0)
		*arg_excp = arg_free
		if hide_argv != 0 {
			var i int
			for i = 1; i < argc; i++ {
				memset(argv[i], 0, strlen(argv[i]))
			}
		}
		zend.ZendIsAutoGlobalStr("_SERVER", g.SizeOf("\"_SERVER\"")-1)
		core.CoreGlobals.during_request_startup = 0
		switch behavior {
		case 1:
			if strcmp(file_handle.filename, "Standard input code") {
				CliRegisterFileHandles()
			}
			if interactive != 0 && CliShellCallbacks.GetCliShellRun() != nil {
				exit_status = CliShellCallbacks.GetCliShellRun()()
			} else {
				core.PhpExecuteScript(&file_handle)
				exit_status = zend.EG.exit_status
			}
			break
		case 4:
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
			}
			goto out
			break
		case 2:
			var syntax_highlighter_ini zend.ZendSyntaxHighlighterIni
			if zend.OpenFileForScanning(&file_handle) == zend.SUCCESS {
				standard.PhpGetHighlight(&syntax_highlighter_ini)
				zend.ZendHighlight(&syntax_highlighter_ini)
			}
			goto out
			break
		case 6:
			CliRegisterFileHandles()
			if zend.ZendEvalStringEx(exec_direct, nil, "Command line code", 1) == zend.FAILURE {
				exit_status = 254
			}
			break
		case 7:
			var input *byte
			var len_ int
			var index int = 0
			var argn zend.Zval
			var argi zend.Zval
			CliRegisterFileHandles()
			if exec_begin != nil && zend.ZendEvalStringEx(exec_begin, nil, "Command line begin code", 1) == zend.FAILURE {
				exit_status = 254
			}
			for exit_status == zend.SUCCESS && g.Assign(&input, streams._phpStreamGetLine(SInProcess, nil, 0, nil)) != nil {
				len_ = strlen(input)
				for len_ > 0 && g.PostDec(&len_) && (input[len_] == '\n' || input[len_] == '\r') {
					input[len_] = '0'
				}
				var __z *zend.Zval = &argn
				var __s *zend.ZendString = zend.ZendStringInit(input, len_+1, 0)
				__z.value.str = __s
				__z.u1.type_info = 6 | 1<<0<<8
				zend.ZendHashStrUpdate(&zend.EG.symbol_table, "argn", g.SizeOf("\"argn\"")-1, &argn)
				var __z *zend.Zval = &argi
				index++
				__z.value.lval = index
				__z.u1.type_info = 4
				zend.ZendHashStrUpdate(&zend.EG.symbol_table, "argi", g.SizeOf("\"argi\"")-1, &argi)
				if exec_run != nil {
					if zend.ZendEvalStringEx(exec_run, nil, "Command line run code", 1) == zend.FAILURE {
						exit_status = 254
					}
				} else {
					if script_file != nil {
						if CliSeekFileBegin(&file_handle, script_file) != zend.SUCCESS {
							exit_status = 1
						} else {
							zend.CG.skip_shebang = 1
							core.PhpExecuteScript(&file_handle)
							exit_status = zend.EG.exit_status
						}
					}
				}
				zend._efree(input)
			}
			if exec_end != nil && zend.ZendEvalStringEx(exec_end, nil, "Command line end code", 1) == zend.FAILURE {
				exit_status = 254
			}
			break
		case 8:

		case 9:

		case 10:

		case 12:
			var pce *zend.ZendClassEntry = nil
			var arg zend.Zval
			var ref zend.Zval
			var execute_data zend.ZendExecuteData
			switch behavior {
			default:
				break
			case 8:
				if strstr(reflection_what, "::") {
					pce = reflection_method_ptr
				} else {
					pce = reflection_function_ptr
				}
				break
			case 9:
				pce = reflection_class_ptr
				break
			case 10:
				pce = reflection_extension_ptr
				break
			case 12:
				pce = reflection_zend_extension_ptr
				break
			}
			var _s *byte = reflection_what
			var __z *zend.Zval = &arg
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			zend.ObjectInitEx(&ref, pce)
			memset(&execute_data, 0, g.SizeOf("zend_execute_data"))
			zend.EG.current_execute_data = &execute_data
			zend.ZendCallMethod(&ref, pce, &pce.constructor, "__construct", g.SizeOf("\"__construct\"")-1, nil, 1, &arg, nil)
			if zend.EG.exception != nil {
				var tmp zend.Zval
				var msg *zend.Zval
				var rv zend.Zval
				var __z *zend.Zval = &tmp
				__z.value.obj = zend.EG.exception
				__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
				msg = zend.ZendReadProperty(zend.ZendCeException, &tmp, "message", g.SizeOf("\"message\"")-1, 0, &rv)
				zend.ZendPrintf("Exception: %s\n", msg.value.str.val)
				zend.ZvalPtrDtor(&tmp)
				zend.EG.exception = nil
				exit_status = 1
			} else {
				zend.ZendPrintZval(&ref, 0)
				zend.ZendWrite("\n", 1)
			}
			zend.ZvalPtrDtor(&ref)
			zend.ZvalPtrDtor(&arg)
			break
		case 11:
			var len_ int = strlen(reflection_what)
			var lcname *byte = zend.ZendStrTolowerDup(reflection_what, len_)
			var module *zend.ZendModuleEntry
			if g.Assign(&module, zend.ZendHashStrFindPtr(&zend.ModuleRegistry, lcname, len_)) == nil {
				if !(strcmp(reflection_what, "main")) {
					core.DisplayIniEntries(nil)
				} else {
					zend.ZendPrintf("Extension '%s' not present.\n", reflection_what)
					exit_status = 1
				}
			} else {
				standard.PhpInfoPrintModule(module)
			}
			zend._efree(lcname)
			break
		case 13:
			zend.ZendPrintf("Configuration File (php.ini) Path: %s\n", "/usr/local/lib")
			zend.ZendPrintf("Loaded Configuration File:         %s\n", g.Cond(PhpIniOpenedPath != nil, PhpIniOpenedPath, "(none)"))
			zend.ZendPrintf("Scan for additional .ini files in: %s\n", g.Cond(PhpIniScannedPath != nil, PhpIniScannedPath, "(none)"))
			zend.ZendPrintf("Additional .ini files parsed:      %s\n", g.Cond(PhpIniScannedFiles != nil, PhpIniScannedFiles, "(none)"))
			break
		}
	}
	zend.EG.bailout = __orig_bailout
out:
	if request_started != 0 {
		core.PhpRequestShutdown(any(0))
	}
	if translated_path != nil {
		zend.Free(translated_path)
	}
	if exit_status == 0 {
		exit_status = zend.EG.exit_status
	}
	return exit_status
err:
	core.SapiDeactivate()
	zend.ZendIniDeactivate()
	exit_status = 1
	goto out
}

/* }}} */

func Main(argc int, argv []*byte) int {
	var c int
	var exit_status int = zend.SUCCESS
	var module_started int = 0
	var sapi_started int = 0
	var php_optarg *byte = nil
	var php_optind int = 1
	var use_extended_info int = 0
	var ini_path_override *byte = nil
	var ini_entries *byte = nil
	var ini_entries_len int = 0
	var ini_ignore int = 0
	var sapi_module *core.sapi_module_struct = &CliSapiModule

	/*
	 * Do not move this initialization. It needs to happen before argv is used
	 * in any way.
	 */

	argv = SavePsArgs(argc, argv)
	CliSapiModule.additional_functions = AdditionalFunctions
	zend.ZendSignalStartup()
	for g.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &php_optarg, &php_optind, 1, 2)) != -1 {
		switch c {
		case 'c':
			if ini_path_override != nil {
				zend.Free(ini_path_override)
			}
			ini_path_override = strdup(php_optarg)
			break
		case 'n':
			ini_ignore = 1
			break
		case 'd':

			/* define ini __special__  entries on command line */

			var len_ int = strlen(php_optarg)
			var val *byte
			if g.Assign(&val, strchr(php_optarg, '=')) {
				val++
				if !(isalnum(*val)) && (*val) != '"' && (*val) != '\'' && (*val) != '0' {
					ini_entries = realloc(ini_entries, ini_entries_len+len_+g.SizeOf("\"\\\"\\\"\\n\\0\""))
					memcpy(ini_entries+ini_entries_len, php_optarg, val-php_optarg)
					ini_entries_len += val - php_optarg
					memcpy(ini_entries+ini_entries_len, "\"", 1)
					ini_entries_len++
					memcpy(ini_entries+ini_entries_len, val, len_-(val-php_optarg))
					ini_entries_len += len_ - (val - php_optarg)
					memcpy(ini_entries+ini_entries_len, "\"\n0", g.SizeOf("\"\\\"\\n\\0\""))
					ini_entries_len += g.SizeOf("\"\\n\\0\\\"\"") - 2
				} else {
					ini_entries = realloc(ini_entries, ini_entries_len+len_+g.SizeOf("\"\\n\\0\""))
					memcpy(ini_entries+ini_entries_len, php_optarg, len_)
					memcpy(ini_entries+ini_entries_len+len_, "\n0", g.SizeOf("\"\\n\\0\""))
					ini_entries_len += len_ + g.SizeOf("\"\\n\\0\"") - 2
				}
			} else {
				ini_entries = realloc(ini_entries, ini_entries_len+len_+g.SizeOf("\"=1\\n\\0\""))
				memcpy(ini_entries+ini_entries_len, php_optarg, len_)
				memcpy(ini_entries+ini_entries_len+len_, "=1\n0", g.SizeOf("\"=1\\n\\0\""))
				ini_entries_len += len_ + g.SizeOf("\"=1\\n\\0\"") - 2
			}
			break
		case 'S':
			sapi_module = &CliServerSapiModule
			CliServerSapiModule.additional_functions = ServerAdditionalFunctions
			break
		case 'h':

		case '?':
			PhpCliUsage(argv[0])
			goto out
		case -2:
			PhpCliUsage(argv[0])
			exit_status = 1
			goto out
		case 'i':

		case 'v':

		case 'm':
			sapi_module = &CliSapiModule
			goto exit_loop
		case 'e':
			use_extended_info = 1
			break
		}
	}
exit_loop:
	sapi_module.ini_defaults = SapiCliIniDefaults
	sapi_module.php_ini_path_override = ini_path_override
	sapi_module.phpinfo_as_text = 1
	sapi_module.php_ini_ignore_cwd = 1
	core.SapiStartup(sapi_module)
	sapi_started = 1
	sapi_module.php_ini_ignore = ini_ignore
	sapi_module.executable_location = argv[0]
	if sapi_module == &CliSapiModule {
		if ini_entries != nil {
			ini_entries = realloc(ini_entries, ini_entries_len+g.SizeOf("HARDCODED_INI"))
			memmove(ini_entries+g.SizeOf("HARDCODED_INI")-2, ini_entries, ini_entries_len+1)
			memcpy(ini_entries, HARDCODED_INI, g.SizeOf("HARDCODED_INI")-2)
		} else {
			ini_entries = zend.Malloc(g.SizeOf("HARDCODED_INI"))
			memcpy(ini_entries, HARDCODED_INI, g.SizeOf("HARDCODED_INI"))
		}
		ini_entries_len += g.SizeOf("HARDCODED_INI") - 2
	}
	sapi_module.ini_entries = ini_entries

	/* startup after we get the above ini override se we get things right */

	if sapi_module.startup(sapi_module) == zend.FAILURE {

		/* there is no way to see if we must call zend_ini_deactivate()
		 * since we cannot check if EG(ini_directives) has been initialised
		 * because the executor's constructor does not set initialize it.
		 * Apart from that there seems no need for zend_ini_deactivate() yet.
		 * So we goto out_err.*/

		exit_status = 1
		goto out
	}
	module_started = 1

	/* -e option */

	if use_extended_info != 0 {
		zend.CG.compiler_options |= 1<<0 | 1<<1
	}
	zend.EG.bailout = nil
	var __orig_bailout *sigjmp_buf = zend.EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		if sapi_module == &CliSapiModule {
			exit_status = DoCli(argc, argv)
		} else {
			exit_status = DoCliServer(argc, argv)
		}
	}
	zend.EG.bailout = __orig_bailout
out:
	if ini_path_override != nil {
		zend.Free(ini_path_override)
	}
	if ini_entries != nil {
		zend.Free(ini_entries)
	}
	if module_started != 0 {
		core.PhpModuleShutdown()
	}
	if sapi_started != 0 {
		core.SapiShutdown()
	}

	/*
	 * Do not move this de-initialization. It needs to happen right before
	 * exiting.
	 */

	CleanupPsArgs(argv)
	exit(exit_status)
}

/* }}} */
