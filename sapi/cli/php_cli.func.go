// <<generate>>

package cli

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	r "sik/runtime"
	"sik/zend"
)

func PhpSelect(m core.PhpSocketT, r fd_set, w __auto__, e __auto__, t *__struct__timeval) __auto__ {
	return select_(m, r, w, e, t)
}
func ModuleNameCmp(a any, b any) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	return strcasecmp((*zend.ZendModuleEntry)(zend.Z_PTR(f.GetVal())).GetName(), (*zend.ZendModuleEntry)(zend.Z_PTR(s.GetVal())).GetName())
}
func PrintModules() {
	var sorted_registry zend.HashTable
	var module *zend.ZendModuleEntry
	zend.ZendHashInit(&sorted_registry, 50, nil, nil, 0)
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
	return zend.ZEND_HASH_APPLY_KEEP
}
func ExtensionNameCmp(f **zend.ZendLlistElement, s **zend.ZendLlistElement) int {
	var fe *zend.ZendExtension = (*zend.ZendExtension)(f.GetData())
	var se *zend.ZendExtension = (*zend.ZendExtension)(s.GetData())
	return strcmp(fe.GetName(), se.GetName())
}
func PrintExtensions() {
	var sorted_exts zend.ZendLlist
	zend.ZendLlistCopy(&sorted_exts, &zend.ZendExtensions)
	sorted_exts.SetDtor(nil)
	zend.ZendLlistSort(&sorted_exts, ExtensionNameCmp)
	zend.ZendLlistApply(&sorted_exts, zend.LlistApplyFuncT(PrintExtensionInfo))
	zend.ZendLlistDestroy(&sorted_exts)
}
func SapiCliSelect(fd core.PhpSocketT) int {
	var wfd fd_set
	var tv __struct__timeval
	var ret int
	FD_ZERO(&wfd)
	core.PHP_SAFE_FD_SET(fd, &wfd)
	tv.tv_sec = long(standard.FG(default_socket_timeout))
	tv.tv_usec = 0
	ret = PhpSelect(fd+1, nil, &wfd, nil, &tv)
	return ret != -1
}
func SapiCliSingleWrite(str *byte, str_length int) ssize_t {
	var ret ssize_t
	for {
		ret = write(STDOUT_FILENO, str, str_length)
		if !(ret <= 0 && errno == EAGAIN && SapiCliSelect(STDOUT_FILENO) != 0) {
			break
		}
	}
	return ret
}
func SapiCliUbWrite(str *byte, str_length int) int {
	var ptr *byte = str
	var remaining int = str_length
	var ret ssize_t
	if str_length == 0 {
		return 0
	}
	for remaining > 0 {
		ret = SapiCliSingleWrite(ptr, remaining)
		if ret < 0 {
			zend.EG__().SetExitStatus(255)
			core.PhpHandleAbortedConnection()
			break
		}
		ptr += ret
		remaining -= ret
	}
	return ptr - str
}
func SapiCliFlush(server_context any) {
	/* Ignore EBADF here, it's caused by the fact that STDIN/STDOUT/STDERR streams
	 * are/could be closed before fflush() is called.
	 */

	if r.Fflush(stdout) == r.EOF && errno != EBADF {
		core.PhpHandleAbortedConnection()
	}
}
func SapiCliRegisterVariables(track_vars_array *zend.Zval) {
	var len_ int
	var docroot *byte = ""

	/* In CGI mode, we consider the environment to be a part of the server
	 * variables
	 */

	core.PhpImportEnvironmentVariables(track_vars_array)

	/* Build the special-case PHP_SELF variable for the CLI version */

	len_ = strlen(PhpSelf)
	if core.sapi_module.GetInputFilter()(core.PARSE_SERVER, "PHP_SELF", &PhpSelf, len_, &len_) != 0 {
		core.PhpRegisterVariable("PHP_SELF", PhpSelf, track_vars_array)
	}
	if core.sapi_module.GetInputFilter()(core.PARSE_SERVER, "SCRIPT_NAME", &PhpSelf, len_, &len_) != 0 {
		core.PhpRegisterVariable("SCRIPT_NAME", PhpSelf, track_vars_array)
	}

	/* filenames are empty for stdin */

	len_ = strlen(ScriptFilename)
	if core.sapi_module.GetInputFilter()(core.PARSE_SERVER, "SCRIPT_FILENAME", &ScriptFilename, len_, &len_) != 0 {
		core.PhpRegisterVariable("SCRIPT_FILENAME", ScriptFilename, track_vars_array)
	}
	if core.sapi_module.GetInputFilter()(core.PARSE_SERVER, "PATH_TRANSLATED", &ScriptFilename, len_, &len_) != 0 {
		core.PhpRegisterVariable("PATH_TRANSLATED", ScriptFilename, track_vars_array)
	}

	/* just make it available */

	len_ = 0
	if core.sapi_module.GetInputFilter()(core.PARSE_SERVER, "DOCUMENT_ROOT", &docroot, len_, &len_) != 0 {
		core.PhpRegisterVariable("DOCUMENT_ROOT", docroot, track_vars_array)
	}
}
func SapiCliLogMessage(message *byte, syslog_type_int int) { r.Fprintf(stderr, "%s\n", message) }
func SapiCliDeactivate() int {
	r.Fflush(stdout)
	if core.SG(request_info).argv0 {
		zend.Free(core.SG(request_info).argv0)
		core.SG(request_info).argv0 = nil
	}
	return zend.SUCCESS
}
func SapiCliReadCookies() *byte { return nil }
func SapiCliHeaderHandler(h *core.SapiHeader, op core.SapiHeaderOpEnum, s *core.SapiHeaders) int {
	return 0
}
func SapiCliSendHeaders(sapi_headers *core.SapiHeaders) int {
	/* We do nothing here, this function is needed to prevent that the fallback
	 * header handling is called. */

	return core.SAPI_HEADER_SENT_SUCCESSFULLY

	/* We do nothing here, this function is needed to prevent that the fallback
	 * header handling is called. */
}
func SapiCliSendHeader(sapi_header *core.SapiHeader, server_context any) {}
func PhpCliStartup(sapi_module *core.sapi_module_struct) int {
	if core.PhpModuleStartup(sapi_module, nil, 0) == zend.FAILURE {
		return zend.FAILURE
	}
	return zend.SUCCESS
}
func INI_DEFAULT(name string, value string) {
	tmp.SetString(zend.ZendStringInit(value, b.SizeOf("value")-1, 1))
	core.ConfigurationHash.KeyUpdate(b.CastStrAuto(name), &tmp)
}
func SapiCliIniDefaults(configuration_hash *zend.HashTable) {
	var tmp zend.Zval
	INI_DEFAULT("report_zend_debug", "0")
	INI_DEFAULT("display_errors", "1")
}
func PhpCliUsage(argv0 *byte) {
	var prog *byte
	prog = strrchr(argv0, '/')
	if prog != nil {
		prog++
	} else {
		prog = "php"
	}
	r.Printf("Usage: %s [options] [-f] <file> [--] [args...]\n"+"   %s [options] -r <code> [--] [args...]\n"+"   %s [options] [-B <begin_code>] -R <code> [-E <end_code>] [--] [args...]\n"+"   %s [options] [-B <begin_code>] -F <file> [-E <end_code>] [--] [args...]\n"+"   %s [options] -S <addr>:<port> [-t docroot] [router]\n"+"   %s [options] -- [args...]\n"+"   %s [options] -a\n"+"\n"+"  -a               Run interactively\n"+"  -c <path>|<file> Look for php.ini file in this directory\n"+"  -n               No configuration (ini) files will be used\n"+"  -d foo[=bar]     Define INI entry foo with value 'bar'\n"+"  -e               Generate extended information for debugger/profiler\n"+"  -f <file>        Parse and execute <file>.\n"+"  -h               This help\n"+"  -i               PHP information\n"+"  -l               Syntax check only (lint)\n"+"  -m               Show compiled in modules\n"+"  -r <code>        Run PHP <code> without using script tags <?..?>\n"+"  -B <begin_code>  Run PHP <begin_code> before processing input lines\n"+"  -R <code>        Run PHP <code> for every input line\n"+"  -F <file>        Parse and execute <file> for every input line\n"+"  -E <end_code>    Run PHP <end_code> after processing all input lines\n"+"  -H               Hide any passed arguments from external tools.\n"+"  -S <addr>:<port> Run with built-in web server.\n"+"  -t <docroot>     Specify document root <docroot> for built-in web server.\n"+"  -s               Output HTML syntax highlighted source.\n"+"  -v               Version number\n"+"  -w               Output source with stripped comments and whitespace.\n"+"  -z <file>        Load Zend extension <file>.\n"+"\n"+"  args...          Arguments passed to script. Use -- args when first argument\n"+"                   starts with - or script is read from stdin\n"+"\n"+"  --ini            Show configuration file names\n"+"\n"+"  --rf <name>      Show information about function <name>.\n"+"  --rc <name>      Show information about class <name>.\n"+"  --re <name>      Show information about extension <name>.\n"+"  --rz <name>      Show information about Zend extension <name>.\n"+"  --ri <name>      Show configuration for extension <name>.\n"+"\n", prog, prog, prog, prog, prog, prog, prog)
}
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
	s_in = core.PhpStreamOpenWrapperEx("php://stdin", "rb", 0, nil, sc_in)
	s_out = core.PhpStreamOpenWrapperEx("php://stdout", "wb", 0, nil, sc_out)
	s_err = core.PhpStreamOpenWrapperEx("php://stderr", "wb", 0, nil, sc_err)
	if s_in == nil || s_out == nil || s_err == nil {
		if s_in != nil {
			core.PhpStreamClose(s_in)
		}
		if s_out != nil {
			core.PhpStreamClose(s_out)
		}
		if s_err != nil {
			core.PhpStreamClose(s_err)
		}
		return
	}
	SInProcess = s_in
	core.PhpStreamToZval(s_in, ic.GetValue())
	core.PhpStreamToZval(s_out, oc.GetValue())
	core.PhpStreamToZval(s_err, ec.GetValue())
	zend.ZEND_CONSTANT_SET_FLAGS(&ic, zend.CONST_CS, 0)
	ic.SetName(zend.ZendStringInitInterned("STDIN", b.SizeOf("\"STDIN\"")-1, 0))
	zend.ZendRegisterConstant(&ic)
	zend.ZEND_CONSTANT_SET_FLAGS(&oc, zend.CONST_CS, 0)
	oc.SetName(zend.ZendStringInitInterned("STDOUT", b.SizeOf("\"STDOUT\"")-1, 0))
	zend.ZendRegisterConstant(&oc)
	zend.ZEND_CONSTANT_SET_FLAGS(&ec, zend.CONST_CS, 0)
	ec.SetName(zend.ZendStringInitInterned("STDERR", b.SizeOf("\"STDERR\"")-1, 0))
	zend.ZendRegisterConstant(&ec)
}
func CliSeekFileBegin(file_handle *zend.ZendFileHandle, script_file *byte) int {
	var fp *r.FILE = zend.VCWD_FOPEN(script_file, "rb")
	if fp == nil {
		core.PhpPrintf("Could not open input file: %s\n", script_file)
		return zend.FAILURE
	}
	zend.ZendStreamInitFp(file_handle, fp, script_file)
	return zend.SUCCESS
}
func DoCli(argc int, argv **byte) int {
	var c int
	var file_handle zend.ZendFileHandle
	var behavior int = PHP_MODE_STANDARD
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
	var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		zend.CG__().SetInCompilation(0)
		for b.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &php_optarg, &php_optind, 0, 2)) != -1 {
			switch c {
			case 'i':
				if core.PhpRequestStartup() == zend.FAILURE {
					goto err
				}
				request_started = 1
				standard.PhpPrintInfo(standard.PHP_INFO_ALL & ^standard.PHP_INFO_CREDITS)
				core.PhpOutputEndAll()
				exit_status = c == '?' && argc > 1 && !(strchr(argv[1], c))
				goto out
			case 'v':
				core.PhpPrintf("PHP %s (%s) (built: %s %s) ( %s)\nCopyright (c) The PHP Group\n%s", core.PHP_VERSION, CliSapiModule.GetName(), __DATE__, __TIME__, "NTS ", zend.GetZendVersion())
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

		core.SG(options) |= core.SAPI_OPTION_NO_CHDIR
		php_optind = orig_optind
		php_optarg = orig_optarg
		for b.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &php_optarg, &php_optind, 0, 2)) != -1 {
			switch c {
			case 'a':
				if interactive == 0 {
					if behavior != PHP_MODE_STANDARD {
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
				if behavior == PHP_MODE_PROCESS_STDIN {
					if exec_run != nil || script_file != nil {
						param_error = "You can use -R or -F only once.\n"
						break
					}
				} else if behavior != PHP_MODE_STANDARD {
					param_error = ParamModeConflict
					break
				}
				behavior = PHP_MODE_PROCESS_STDIN
				script_file = php_optarg
				break
			case 'f':
				if behavior == PHP_MODE_CLI_DIRECT || behavior == PHP_MODE_PROCESS_STDIN {
					param_error = ParamModeConflict
					break
				} else if script_file != nil {
					param_error = "You can use -f only once.\n"
					break
				}
				script_file = php_optarg
				break
			case 'l':
				if behavior != PHP_MODE_STANDARD {
					break
				}
				behavior = PHP_MODE_LINT
				break
			case 'q':

				/* This is default so NOP */

				break
			case 'r':
				if behavior == PHP_MODE_CLI_DIRECT {
					if exec_direct != nil || script_file != nil {
						param_error = "You can use -r only once.\n"
						break
					}
				} else if behavior != PHP_MODE_STANDARD || interactive != 0 {
					param_error = ParamModeConflict
					break
				}
				behavior = PHP_MODE_CLI_DIRECT
				exec_direct = php_optarg
				break
			case 'R':
				if behavior == PHP_MODE_PROCESS_STDIN {
					if exec_run != nil || script_file != nil {
						param_error = "You can use -R or -F only once.\n"
						break
					}
				} else if behavior != PHP_MODE_STANDARD {
					param_error = ParamModeConflict
					break
				}
				behavior = PHP_MODE_PROCESS_STDIN
				exec_run = php_optarg
				break
			case 'B':
				if behavior == PHP_MODE_PROCESS_STDIN {
					if exec_begin != nil {
						param_error = "You can use -B only once.\n"
						break
					}
				} else if behavior != PHP_MODE_STANDARD || interactive != 0 {
					param_error = ParamModeConflict
					break
				}
				behavior = PHP_MODE_PROCESS_STDIN
				exec_begin = php_optarg
				break
			case 'E':
				if behavior == PHP_MODE_PROCESS_STDIN {
					if exec_end != nil {
						param_error = "You can use -E only once.\n"
						break
					}
				} else if behavior != PHP_MODE_STANDARD || interactive != 0 {
					param_error = ParamModeConflict
					break
				}
				behavior = PHP_MODE_PROCESS_STDIN
				exec_end = php_optarg
				break
			case 's':
				if behavior == PHP_MODE_CLI_DIRECT || behavior == PHP_MODE_PROCESS_STDIN {
					param_error = "Source highlighting only works for files.\n"
					break
				}
				behavior = PHP_MODE_HIGHLIGHT
				break
			case 'w':
				if behavior == PHP_MODE_CLI_DIRECT || behavior == PHP_MODE_PROCESS_STDIN {
					param_error = "Source stripping only works for files.\n"
					break
				}
				behavior = PHP_MODE_STRIP
				break
			case 'z':
				zend.ZendLoadExtension(php_optarg)
				break
			case 'H':
				hide_argv = 1
				break
			case 10:
				behavior = PHP_MODE_REFLECTION_FUNCTION
				reflection_what = php_optarg
				break
			case 11:
				behavior = PHP_MODE_REFLECTION_CLASS
				reflection_what = php_optarg
				break
			case 12:
				behavior = PHP_MODE_REFLECTION_EXTENSION
				reflection_what = php_optarg
				break
			case 13:
				behavior = PHP_MODE_REFLECTION_ZEND_EXTENSION
				reflection_what = php_optarg
				break
			case 14:
				behavior = PHP_MODE_REFLECTION_EXT_INFO
				reflection_what = php_optarg
				break
			case 15:
				behavior = PHP_MODE_SHOW_INI_CONFIG
				break
			default:
				break
			}
		}
		if param_error != nil {
			core.PUTS(param_error)
			exit_status = 1
			goto err
		}
		if interactive != 0 {
			r.Printf("Interactive mode enabled\n\n")
			r.Fflush(stdout)
		}

		/* only set script_file if not set already and not in direct mode and not at end of parameter list */

		if argc > php_optind && script_file == nil && behavior != PHP_MODE_CLI_DIRECT && behavior != PHP_MODE_PROCESS_STDIN && strcmp(argv[php_optind-1], "--") {
			script_file = argv[php_optind]
			php_optind++
		}
		if script_file != nil {
			if CliSeekFileBegin(&file_handle, script_file) != zend.SUCCESS {
				goto err
			} else {
				var real_path []byte
				if zend.VCWD_REALPATH(script_file, real_path) != nil {
					translated_path = strdup(real_path)
				}
				ScriptFilename = script_file
			}
		} else {

			/* We could handle PHP_MODE_PROCESS_STDIN in a different manner  */

			zend.ZendStreamInitFp(&file_handle, stdin, "Standard input code")

			/* We could handle PHP_MODE_PROCESS_STDIN in a different manner  */

		}
		PhpSelf = (*byte)(file_handle.GetFilename())

		/* before registering argv to module exchange the *new* argv[0] */

		core.SG(request_info).argc = argc - php_optind + 1
		arg_excp = argv + php_optind - 1
		arg_free = argv[php_optind-1]
		if translated_path != nil {
			core.SG(request_info).path_translated = translated_path
		} else {
			core.SG(request_info).path_translated = (*byte)(file_handle.GetFilename())
		}
		argv[php_optind-1] = (*byte)(file_handle.GetFilename())
		core.SG(request_info).argv = argv + php_optind - 1
		if core.PhpRequestStartup() == zend.FAILURE {
			*arg_excp = arg_free
			r.Fclose(file_handle.GetFp())
			core.PUTS("Could not startup.\n")
			goto err
		}
		request_started = 1
		zend.CG__().SetSkipShebang(1)
		zend.ZendRegisterBoolConstant(zend.ZEND_STRL("PHP_CLI_PROCESS_TITLE"), IsPsTitleAvailable() == PS_TITLE_SUCCESS, zend.CONST_CS, 0)
		*arg_excp = arg_free
		if hide_argv != 0 {
			var i int
			for i = 1; i < argc; i++ {
				memset(argv[i], 0, strlen(argv[i]))
			}
		}
		zend.ZendIsAutoGlobalStr(zend.ZEND_STRL("_SERVER"))
		core.PG(during_request_startup) = 0
		switch behavior {
		case PHP_MODE_STANDARD:
			if strcmp(file_handle.GetFilename(), "Standard input code") {
				CliRegisterFileHandles()
			}
			core.PhpExecuteScript(&file_handle)
			exit_status = zend.EG__().GetExitStatus()
			break
		case PHP_MODE_LINT:
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
			}
			goto out
			break
		case PHP_MODE_HIGHLIGHT:
			var syntax_highlighter_ini zend.ZendSyntaxHighlighterIni
			if zend.OpenFileForScanning(&file_handle) == zend.SUCCESS {
				standard.PhpGetHighlight(&syntax_highlighter_ini)
				zend.ZendHighlight(&syntax_highlighter_ini)
			}
			goto out
			break
		case PHP_MODE_CLI_DIRECT:
			CliRegisterFileHandles()
			if zend.ZendEvalStringEx(exec_direct, nil, "Command line code", 1) == zend.FAILURE {
				exit_status = 254
			}
			break
		case PHP_MODE_PROCESS_STDIN:
			var input *byte
			var len_ int
			var index int = 0
			var argn zend.Zval
			var argi zend.Zval
			CliRegisterFileHandles()
			if exec_begin != nil && zend.ZendEvalStringEx(exec_begin, nil, "Command line begin code", 1) == zend.FAILURE {
				exit_status = 254
			}
			for exit_status == zend.SUCCESS && b.Assign(&input, core.PhpStreamGets(SInProcess, nil, 0)) != nil {
				len_ = strlen(input)
				for len_ > 0 && b.PostDec(&len_) && (input[len_] == '\n' || input[len_] == '\r') {
					input[len_] = '0'
				}
				zend.ZVAL_STRINGL(&argn, input, len_+1)
				zend.EG__().GetSymbolTable().KeyUpdate("argn", &argn)
				argi.SetLong(b.PreInc(&index))
				zend.EG__().GetSymbolTable().KeyUpdate("argi", &argi)
				if exec_run != nil {
					if zend.ZendEvalStringEx(exec_run, nil, "Command line run code", 1) == zend.FAILURE {
						exit_status = 254
					}
				} else {
					if script_file != nil {
						if CliSeekFileBegin(&file_handle, script_file) != zend.SUCCESS {
							exit_status = 1
						} else {
							zend.CG__().SetSkipShebang(1)
							core.PhpExecuteScript(&file_handle)
							exit_status = zend.EG__().GetExitStatus()
						}
					}
				}
				zend.Efree(input)
			}
			if exec_end != nil && zend.ZendEvalStringEx(exec_end, nil, "Command line end code", 1) == zend.FAILURE {
				exit_status = 254
			}
			break
		case PHP_MODE_REFLECTION_FUNCTION:

		case PHP_MODE_REFLECTION_CLASS:

		case PHP_MODE_REFLECTION_EXTENSION:

		case PHP_MODE_REFLECTION_ZEND_EXTENSION:
			var pce *zend.ZendClassEntry = nil
			var arg zend.Zval
			var ref zend.Zval
			var execute_data zend.ZendExecuteData
			switch behavior {
			default:
				break
			case PHP_MODE_REFLECTION_FUNCTION:
				if strstr(reflection_what, "::") {
					pce = reflection_method_ptr
				} else {
					pce = reflection_function_ptr
				}
				break
			case PHP_MODE_REFLECTION_CLASS:
				pce = reflection_class_ptr
				break
			case PHP_MODE_REFLECTION_EXTENSION:
				pce = reflection_extension_ptr
				break
			case PHP_MODE_REFLECTION_ZEND_EXTENSION:
				pce = reflection_zend_extension_ptr
				break
			}
			zend.ZVAL_STRING(&arg, reflection_what)
			zend.ObjectInitEx(&ref, pce)
			memset(&execute_data, 0, b.SizeOf("zend_execute_data"))
			zend.EG__().SetCurrentExecuteData(&execute_data)
			zend.ZendCallMethodWith1Params(&ref, pce, pce.GetConstructor(), "__construct", nil, &arg)
			if zend.EG__().GetException() != nil {
				var tmp zend.Zval
				var msg *zend.Zval
				var rv zend.Zval
				tmp.SetObject(zend.EG__().GetException())
				msg = zend.ZendReadProperty(zend.ZendCeException, &tmp, "message", b.SizeOf("\"message\"")-1, 0, &rv)
				zend.ZendPrintf("Exception: %s\n", zend.Z_STRVAL_P(msg))
				zend.ZvalPtrDtor(&tmp)
				zend.EG__().SetException(nil)
				exit_status = 1
			} else {
				zend.ZendPrintZval(&ref, 0)
				zend.ZendWrite("\n", 1)
			}
			zend.ZvalPtrDtor(&ref)
			zend.ZvalPtrDtor(&arg)
			break
		case PHP_MODE_REFLECTION_EXT_INFO:
			var len_ int = strlen(reflection_what)
			var lcname *byte = zend.ZendStrTolowerDup(reflection_what, len_)
			var module *zend.ZendModuleEntry
			if b.Assign(&module, zend.ZendHashStrFindPtr(&zend.ModuleRegistry, lcname, len_)) == nil {
				if !(strcmp(reflection_what, "main")) {
					core.DisplayIniEntries(nil)
				} else {
					zend.ZendPrintf("Extension '%s' not present.\n", reflection_what)
					exit_status = 1
				}
			} else {
				standard.PhpInfoPrintModule(module)
			}
			zend.Efree(lcname)
			break
		case PHP_MODE_SHOW_INI_CONFIG:
			zend.ZendPrintf("Configuration File (php.ini) Path: %s\n", core.PHP_CONFIG_FILE_PATH)
			zend.ZendPrintf("Loaded Configuration File:         %s\n", b.Cond(PhpIniOpenedPath != nil, PhpIniOpenedPath, "(none)"))
			zend.ZendPrintf("Scan for additional .ini files in: %s\n", b.Cond(PhpIniScannedPath != nil, PhpIniScannedPath, "(none)"))
			zend.ZendPrintf("Additional .ini files parsed:      %s\n", b.Cond(PhpIniScannedFiles != nil, PhpIniScannedFiles, "(none)"))
			break
		}
	}
	zend.EG__().SetBailout(__orig_bailout)
out:
	if request_started != 0 {
		core.PhpRequestShutdown(any(0))
	}
	if translated_path != nil {
		zend.Free(translated_path)
	}
	if exit_status == 0 {
		exit_status = zend.EG__().GetExitStatus()
	}
	return exit_status
err:
	core.SapiDeactivate()
	zend.ZendIniDeactivate()
	exit_status = 1
	goto out
}
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
	CliSapiModule.SetAdditionalFunctions(AdditionalFunctions)
	zend.ZendSignalStartup()
	for b.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &php_optarg, &php_optind, 1, 2)) != -1 {
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
			if b.Assign(&val, strchr(php_optarg, '=')) {
				val++
				if !(isalnum(*val)) && (*val) != '"' && (*val) != '\'' && (*val) != '0' {
					ini_entries = realloc(ini_entries, ini_entries_len+len_+b.SizeOf("\"\\\"\\\"\\n\\0\""))
					memcpy(ini_entries+ini_entries_len, php_optarg, val-php_optarg)
					ini_entries_len += val - php_optarg
					memcpy(ini_entries+ini_entries_len, "\"", 1)
					ini_entries_len++
					memcpy(ini_entries+ini_entries_len, val, len_-(val-php_optarg))
					ini_entries_len += len_ - (val - php_optarg)
					memcpy(ini_entries+ini_entries_len, "\"\n0", b.SizeOf("\"\\\"\\n\\0\""))
					ini_entries_len += b.SizeOf("\"\\n\\0\\\"\"") - 2
				} else {
					ini_entries = realloc(ini_entries, ini_entries_len+len_+b.SizeOf("\"\\n\\0\""))
					memcpy(ini_entries+ini_entries_len, php_optarg, len_)
					memcpy(ini_entries+ini_entries_len+len_, "\n0", b.SizeOf("\"\\n\\0\""))
					ini_entries_len += len_ + b.SizeOf("\"\\n\\0\"") - 2
				}
			} else {
				ini_entries = realloc(ini_entries, ini_entries_len+len_+b.SizeOf("\"=1\\n\\0\""))
				memcpy(ini_entries+ini_entries_len, php_optarg, len_)
				memcpy(ini_entries+ini_entries_len+len_, "=1\n0", b.SizeOf("\"=1\\n\\0\""))
				ini_entries_len += len_ + b.SizeOf("\"=1\\n\\0\"") - 2
			}
			break
		case 'S':
			sapi_module = &CliServerSapiModule
			CliServerSapiModule.SetAdditionalFunctions(ServerAdditionalFunctions)
			break
		case 'h':

		case '?':
			PhpCliUsage(argv[0])
			goto out
		case core.PHP_GETOPT_INVALID_ARG:
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
	sapi_module.SetIniDefaults(SapiCliIniDefaults)
	sapi_module.SetPhpIniPathOverride(ini_path_override)
	sapi_module.SetPhpinfoAsText(1)
	sapi_module.SetPhpIniIgnoreCwd(1)
	core.SapiStartup(sapi_module)
	sapi_started = 1
	sapi_module.SetPhpIniIgnore(ini_ignore)
	sapi_module.SetExecutableLocation(argv[0])
	if sapi_module == &CliSapiModule {
		if ini_entries != nil {
			ini_entries = realloc(ini_entries, ini_entries_len+b.SizeOf("HARDCODED_INI"))
			memmove(ini_entries+b.SizeOf("HARDCODED_INI")-2, ini_entries, ini_entries_len+1)
			memcpy(ini_entries, HARDCODED_INI, b.SizeOf("HARDCODED_INI")-2)
		} else {
			ini_entries = zend.Malloc(b.SizeOf("HARDCODED_INI"))
			memcpy(ini_entries, HARDCODED_INI, b.SizeOf("HARDCODED_INI"))
		}
		ini_entries_len += b.SizeOf("HARDCODED_INI") - 2
	}
	sapi_module.SetIniEntries(ini_entries)

	/* startup after we get the above ini override se we get things right */

	if sapi_module.GetStartup()(sapi_module) == zend.FAILURE {

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
		zend.CG__().SetCompilerOptions(zend.CG__().GetCompilerOptions() | zend.ZEND_COMPILE_EXTENDED_INFO)
	}
	zend.EG__().SetBailout(nil)
	var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		if sapi_module == &CliSapiModule {
			exit_status = DoCli(argc, argv)
		} else {
			exit_status = DoCliServer(argc, argv)
		}
	}
	zend.EG__().SetBailout(__orig_bailout)
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
