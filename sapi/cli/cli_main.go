package cli

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func main(argc int, argv []*byte) int {
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
	var sapi_module *core.SapiModule = &CliSapiModule

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
	app := core.NewApp()

	sapi_module.SetIniDefaults(SapiCliIniDefaults)
	sapi_module.SetPhpIniPathOverride(ini_path_override)
	sapi_module.SetPhpinfoAsText(1)
	sapi_module.SetPhpIniIgnoreCwd(1)

	app.Startup(sapi_module)
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
