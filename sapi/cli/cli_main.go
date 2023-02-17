package cli

import (
	"fmt"
	"log"
	"os"
	"sik/core"
	"sik/zend"
	"strings"
)

func isAlphaNum(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func main() int {
	args := os.Args

	var use_extended_info int = 0
	var ini_path_override string = ""
	var ini_entries string = ""
	var ini_ignore int = 0
	var sapiModule ICliSapiModule = CliModule

	zend.ZendSignalStartup()

	optArgs, err := core.GetOpts(args[1:], OPTIONS)
	if err != nil {
		log.Printf(err.Error() + "\n")
		PhpCliUsage(args[0])
		return 0
	}

loop:
	for _, optVal := range optArgs.OptionValues {
		//for b.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &php_optarg, &php_optind, 1, 2)) != -1 {
		switch optVal.Char {
		case 'c':
			ini_path_override = optVal.Value
		case 'n':
			ini_ignore = 1
		case 'd':
			/* define ini __special__  entries on command line */
			if pos := strings.IndexByte(optVal.Value, '='); pos >= 0 {
				defName := optVal.Value[0:pos]
				defValue := optVal.Value[pos+1:]
				if len(defValue) > 0 && !isAlphaNum(defValue[0]) && defValue[0] != '"' && defValue[0] != '\'' {
					ini_entries += defName + "\"" + defValue + "\"\n"
				} else {
					ini_entries += optVal.Value + "\n"
				}
			} else {
				ini_entries += optVal.Value + "=1\n"
			}
		case 'S':
			sapiModule = CliServerModule
			CliServerModule.SetAdditionalFunctions(ServerAdditionalFunctions)
		case 'h', '?':
			PhpCliUsage(args[0])
			return 0
		case 'i', 'v', 'm':
			sapiModule = CliModule
			break loop
		case 'e':
			use_extended_info = 1
		}
	}

	app := core.NewApp()
	sapiModule.SetIniDefaults(SapiCliIniDefaults)
	sapiModule.SetPhpIniPathOverride(ini_path_override)
	sapiModule.SetPhpinfoAsText(1)
	sapiModule.SetPhpIniIgnoreCwd(1)

	app.SapiStartup(sapiModule)
	defer app.SapiShutdown()

	sapiModule.SetPhpIniIgnore(ini_ignore)
	sapiModule.SetExecutableLocation(args[0])

	if sapiModule == CliModule {
		ini_entries += HARDCODED_INI
	}
	sapiModule.SetIniEntries(ini_entries)

	/* startup after we get the above ini override se we get things right */

	if !sapiModule.Startup() {
		/* there is no way to see if we must call zend_ini_deactivate()
		 * since we cannot check if EG(ini_directives) has been initialised
		 * because the executor's constructor does not set initialize it.
		 * Apart from that there seems no need for zend_ini_deactivate() yet.
		 * So we goto out_err.*/
		return 1
	}
	defer core.PhpModuleShutdown()

	/* -e option */

	if use_extended_info != 0 {
		zend.CG__().SetCompilerOptions(zend.CG__().GetCompilerOptions() | zend.ZEND_COMPILE_EXTENDED_INFO)
	}

	// try-catch
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
	}()

	if sapiModule == CliModule {
		return DoCli(nil, nil, args)
	} else {
		return DoCliServer(optArgs)
	}
}
