package cli

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/globals"
	"github.com/heyuuu/gophp/zend/types"
	"os"
	"sort"
	"strings"
)

const usage string = `Usage: %s [options] [-f] <file> [--] [args...]
   %s [options] -r <code> [--] [args...]
   %s [options] [-B <begin_code>] -R <code> [-E <end_code>] [--] [args...]
   %s [options] [-B <begin_code>] -F <file> [-E <end_code>] [--] [args...]
   %s [options] -S <addr>:<port> [-t docroot] [router]
   %s [options] -- [args...]
   %s [options] -a

  -a               Run interactively
  -c <path>|<file> Look for php.ini file in this directory
  -n               No configuration (ini) files will be used
  -d foo[=bar]     Define INI entry foo with value 'bar'
  -e               Generate extended information for debugger/profiler
  -f <file>        Parse and execute <file>.
  -h               This help
  -i               PHP information
  -l               Syntax check only (lint)
  -m               Show compiled in modules
  -r <code>        Run PHP <code> without using script tags <?..?>
  -B <begin_code>  Run PHP <begin_code> before processing input lines
  -R <code>        Run PHP <code> for every input line
  -F <file>        Parse and execute <file> for every input line
  -E <end_code>    Run PHP <end_code> after processing all input lines
  -H               Hide any passed arguments from external tools.
  -S <addr>:<port> Run with built-in web server.
  -t <docroot>     Specify document root <docroot> for built-in web server.
  -s               Output HTML syntax highlighted source.
  -v               Version number
  -w               Output source with stripped comments and whitespace.
  -z <file>        Load Zend extension <file>.

  args...          Arguments passed to script. Use -- args when first argument
                   starts with - or script is read from stdin

  --ini            Show configuration file names

  --rf <name>      Show information about function <name>.
  --rc <name>      Show information about class <name>.
  --re <name>      Show information about extension <name>.
  --rz <name>      Show information about Zend extension <name>.
  --ri <name>      Show configuration for extension <name>.
`

func PrintModules() {
	var modules = globals.G().GetSortedModules()
	for _, module := range modules {
		core.PhpPrintf("%s\n", module.GetName())
	}
}
func PrintExtensionInfo(ext *zend.ZendExtension, arg any) int {
	core.PhpPrintf("%s\n", ext.GetName())
	return types.ArrayApplyKeep
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
func SapiCliRegisterVariables(track_vars_array *types.Zval) {
	var len_ int
	var docroot *byte = ""

	/* In CGI mode, we consider the environment to be a part of the server
	 * variables
	 */

	core.PhpImportEnvironmentVariables(track_vars_array)

	/* Build the special-case PHP_SELF variable for the CLI version */

	len_ = strlen(PhpSelf)
	if core.SM__().GetInputFilter()(core.PARSE_SERVER, "PHP_SELF", &PhpSelf, len_, &len_) != 0 {
		core.PhpRegisterVariable("PHP_SELF", b.CastStrAuto(PhpSelf), track_vars_array)
	}
	if core.SM__().GetInputFilter()(core.PARSE_SERVER, "SCRIPT_NAME", &PhpSelf, len_, &len_) != 0 {
		core.PhpRegisterVariable("SCRIPT_NAME", b.CastStrAuto(PhpSelf), track_vars_array)
	}

	/* filenames are empty for stdin */

	len_ = strlen(ScriptFilename)
	if core.SM__().GetInputFilter()(core.PARSE_SERVER, "SCRIPT_FILENAME", &ScriptFilename, len_, &len_) != 0 {
		core.PhpRegisterVariable("SCRIPT_FILENAME", b.CastStrAuto(ScriptFilename), track_vars_array)
	}
	if core.SM__().GetInputFilter()(core.PARSE_SERVER, "PATH_TRANSLATED", &ScriptFilename, len_, &len_) != 0 {
		core.PhpRegisterVariable("PATH_TRANSLATED", b.CastStrAuto(ScriptFilename), track_vars_array)
	}

	/* just make it available */

	len_ = 0
	if core.SM__().GetInputFilter()(core.PARSE_SERVER, "DOCUMENT_ROOT", &docroot, len_, &len_) != 0 {
		core.PhpRegisterVariable("DOCUMENT_ROOT", b.CastStrAuto(docroot), track_vars_array)
	}
}
func SapiCliIniDefaults(configuration_hash *types.Array) {
	core.Config().Set("report_zend_debug", "0")
	core.Config().Set("display_errors", "1")
}
func PhpCliUsage(argv0 string) {
	bin := "php"
	if pos := strings.LastIndexByte(argv0, '/'); pos >= 0 {
		bin = argv0[pos+1:]
	}
	fmt.Print(strings.ReplaceAll(usage, "%s", bin))
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

	core.PhpStreamToZval(s_in, ic.Value())
	core.PhpStreamToZval(s_out, oc.Value())
	core.PhpStreamToZval(s_err, ec.Value())

	ic.SetFlags(zend.CONST_CS, 0)
	ic.SetName("STDIN")
	zend.ZendRegisterConstant(&ic)
	oc.SetFlags(zend.CONST_CS, 0)
	oc.SetName("STDOUT")
	zend.ZendRegisterConstant(&oc)
	ec.SetFlags(zend.CONST_CS, 0)
	ec.SetName("STDERR")
	zend.ZendRegisterConstant(&ec)
}
func CliSeekFileBegin(script_file string) *zend.FileHandle {
	fh := zend.NewFileHandleByOpenFile(script_file)
	if fh == nil {
		core.PhpPrintf("Could not open input file: %s\n", script_file)
	}
	return fh
}
func DoCli(argc int, argv **byte, args []string) int {
	var c int
	var file_handle *zend.FileHandle
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
	faults.Try(func() {
		zend.CG__().SetInCompilation(0)
		for b.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &php_optarg, &php_optind, 0, 2)) != -1 {
			switch c {
			case 'i':
				if core.PhpRequestStartup() == types.FAILURE {
					goto err
				}
				request_started = 1
				standard.PhpPrintInfo(standard.PHP_INFO_ALL & ^standard.PHP_INFO_CREDITS)
				core.PhpOutputEndAll()
				exit_status = c == '?' && argc > 1 && !(strchr(argv[1], c))
				goto out
			case 'v':
				core.PhpPrintf("PHP %s (%s) (built: %s %s) ( %s)\nCopyright (c) The PHP Group\n%s", core.PHP_VERSION, CliModule.Name(), __DATE__, __TIME__, "NTS ", zend.GetZendVersion())
				core.SapiDeactivate()
				goto out
			case 'm':
				if core.PhpRequestStartup() == types.FAILURE {
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

		core.SG__().options |= core.SAPI_OPTION_NO_CHDIR
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
			os.Stdout.WriteString("Interactive mode enabled\n\n")
			os.Stdout.Sync()
		}

		/* only set script_file if not set already and not in direct mode and not at end of parameter list */

		if argc > php_optind && script_file == nil && behavior != PHP_MODE_CLI_DIRECT && behavior != PHP_MODE_PROCESS_STDIN && strcmp(argv[php_optind-1], "--") {
			script_file = argv[php_optind]
			php_optind++
		}
		if script_file != nil {
			file_handle = CliSeekFileBegin(script_file)
			if file_handle == nil {
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
			file_handle = zend.NewFileHandleForStdin()

		}
		PhpSelf = (*byte)(file_handle.GetFilename())

		/* before registering argv to module exchange the *new* argv[0] */

		core.SG__().RequestInfo.argc = argc - php_optind + 1
		arg_excp = argv + php_optind - 1
		arg_free = argv[php_optind-1]
		if translated_path != nil {
			core.SG__().RequestInfo.path_translated = translated_path
		} else {
			core.SG__().RequestInfo.path_translated = (*byte)(file_handle.GetFilename())
		}
		argv[php_optind-1] = (*byte)(file_handle.GetFilename())
		core.SG__().RequestInfo.argv = argv + php_optind - 1
		if core.PhpRequestStartup() == types.FAILURE {
			*arg_excp = arg_free
			file_handle.Close()
			core.PUTS("Could not startup.\n")
			goto err
		}
		request_started = 1
		zend.CG__().SetSkipShebang(1)
		zend.RegisterBoolConstant("PHP_CLI_PROCESS_TITLE", false, zend.CONST_CS, 0)
		*arg_excp = arg_free
		if hide_argv != 0 {
			var i int
			for i = 1; i < argc; i++ {
				memset(argv[i], 0, strlen(argv[i]))
			}
		}
		zend.ZendIsAutoGlobalStr("_SERVER")
		core.PG__().during_request_startup = 0
		switch behavior {
		case PHP_MODE_STANDARD:
			if strcmp(file_handle.GetFilename(), "Standard input code") {
				CliRegisterFileHandles()
			}
			core.PhpExecuteScript(file_handle)
			exit_status = zend.EG__().GetExitStatus()
			break
		case PHP_MODE_LINT:
			exit_status = core.PhpLintScript(file_handle)
			if exit_status == types.SUCCESS {
				core.PhpPrintf("No syntax errors detected in %s\n", file_handle.GetFilename())
			} else {
				core.PhpPrintf("Errors parsing %s\n", file_handle.GetFilename())
			}
			break
		case PHP_MODE_STRIP:
			if zend.OpenFileForScanning(file_handle) == types.SUCCESS {
				zend.ZendStrip()
			}
			goto out
			break
		case PHP_MODE_HIGHLIGHT:
			var syntax_highlighter_ini zend.ZendSyntaxHighlighterIni
			if zend.OpenFileForScanning(file_handle) == types.SUCCESS {
				standard.PhpGetHighlight(&syntax_highlighter_ini)
				zend.ZendHighlight(&syntax_highlighter_ini)
			}
			goto out
			break
		case PHP_MODE_CLI_DIRECT:
			CliRegisterFileHandles()
			if zend.ZendEvalStringEx(exec_direct, nil, "Command line code", 1) == types.FAILURE {
				exit_status = 254
			}
			break
		case PHP_MODE_PROCESS_STDIN:
			var input *byte
			var len_ int
			var index int = 0
			var argn types.Zval
			var argi types.Zval
			CliRegisterFileHandles()
			if exec_begin != nil && zend.ZendEvalStringEx(exec_begin, nil, "Command line begin code", 1) == types.FAILURE {
				exit_status = 254
			}
			for exit_status == types.SUCCESS && b.Assign(&input, core.PhpStreamGets(SInProcess, nil, 0)) != nil {
				len_ = strlen(input)
				for len_ > 0 && b.PostDec(&len_) && (input[len_] == '\n' || input[len_] == '\r') {
					input[len_] = '0'
				}
				argn.SetStringVal(b.CastStr(input, len_+1))
				zend.EG__().GetSymbolTable().KeyUpdate("argn", &argn)
				argi.SetLong(b.PreInc(&index))
				zend.EG__().GetSymbolTable().KeyUpdate("argi", &argi)
				if exec_run != nil {
					if zend.ZendEvalStringEx(exec_run, nil, "Command line run code", 1) == types.FAILURE {
						exit_status = 254
					}
				} else {
					if script_file != nil {
						file_handle = CliSeekFileBegin(script_file)
						if file_handle == nil {
							exit_status = 1
						} else {
							zend.CG__().SetSkipShebang(1)
							core.PhpExecuteScript(file_handle)
							exit_status = zend.EG__().GetExitStatus()
						}
					}
				}
				zend.Efree(input)
			}
			if exec_end != nil && zend.ZendEvalStringEx(exec_end, nil, "Command line end code", 1) == types.FAILURE {
				exit_status = 254
			}
			break
		case PHP_MODE_REFLECTION_FUNCTION:

		case PHP_MODE_REFLECTION_CLASS:

		case PHP_MODE_REFLECTION_EXTENSION:

		case PHP_MODE_REFLECTION_ZEND_EXTENSION:
			var pce *types.ClassEntry = nil
			var arg types.Zval
			var ref types.Zval
			var executeData zend.ZendExecuteData
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
			arg.SetStringVal(b.CastStrAuto(reflection_what))
			zend.ObjectInitEx(&ref, pce)
			memset(&executeData, 0, b.SizeOf("zend_execute_data"))
			zend.EG__().SetCurrentExecuteData(&executeData)
			zend.ZendCallMethodWith1Params(&ref, pce, pce.GetConstructor(), "__construct", nil, &arg)
			if zend.EG__().GetException() != nil {
				var tmp types.Zval
				var msg *types.Zval
				var rv types.Zval
				tmp.SetObject(zend.EG__().GetException())
				msg = zend.ZendReadProperty(faults.ZendCeException, &tmp, "message", 0, &rv)
				core.PhpPrintf("Exception: %s\n", msg.String().GetVal())
				// zend.ZvalPtrDtor(&tmp)
				zend.EG__().SetException(nil)
				exit_status = 1
			} else {
				zend.ZendPrintZval(&ref)
				zend.ZendWrite("\n")
			}
			// zend.ZvalPtrDtor(&ref)
			// zend.ZvalPtrDtor(&arg)
			break
		case PHP_MODE_REFLECTION_EXT_INFO:
			if module := globals.G().GetModule(b.CastStrAuto(reflection_what)); module != nil {
				standard.PhpInfoPrintModule(module)
			} else {
				if reflection_what == "main" {
					core.DisplayIniEntries(nil)
				} else {
					core.PhpPrintf("Extension '%s' not present.\n", reflection_what)
					exit_status = 1
				}
			}
			break
		case PHP_MODE_SHOW_INI_CONFIG:
			core.PhpPrintf("Configuration File (php.ini) Path: %s\n", core.PHP_CONFIG_FILE_PATH)
			core.PhpPrintf("Loaded Configuration File:         %s\n", b.Cond(PhpIniOpenedPath != nil, PhpIniOpenedPath, "(none)"))
			core.PhpPrintf("Scan for additional .ini files in: %s\n", b.Cond(PhpIniScannedPath != nil, PhpIniScannedPath, "(none)"))
			core.PhpPrintf("Additional .ini files parsed:      %s\n", b.Cond(PhpIniScannedFiles != nil, PhpIniScannedFiles, "(none)"))
			break
		}
	})
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
