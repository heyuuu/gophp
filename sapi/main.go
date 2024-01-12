package sapi

import (
	"errors"
	"fmt"
	"github.com/heyuuu/gophp/php"
	_ "github.com/heyuuu/gophp/php/boot"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/shim/slices"
)

const usage = `Usage: php [options] [-f] <file> [--] [args...]
   php [options] -r <code> [--] [args...]
   php [options] -S <addr>:<port> [-t docroot] [router]
   php [options] -- [args...]

  -c <path>|<file> Look for php.ini file in this directory
  -n               No configuration (ini) files will be used
  -d foo[=bar]     Define INI entry foo with value 'bar'
  -f <file>        Parse and execute <file>.
  -h               This help
  -i               PHP information
  -m               Show compiled in modules
  -r <code>        Run PHP <code> without using script tags <?..?>
  -H               Hide any passed arguments from external tools.
  -S <addr>:<port> Run with built-in web server.
  -t <docroot>     Specify document root <docroot> for built-in web server.
  -v               Version number

  args...          Arguments passed to script. Use -- args when first argument
                   starts with - or script is read from stdin

  --ini            Show configuration file names
`

var Options = []php.Opt{
	php.MakeOpt('C', 0, "no-chdir"),
	php.MakeOpt('c', 1, "php-ini"),
	php.MakeOpt('d', 1, "define"),
	php.MakeOpt('f', 1, "file"),
	php.MakeOpt('h', 0, "help"),
	php.MakeOpt('i', 0, "info"),
	php.MakeOpt('m', 0, "modules"),
	php.MakeOpt('n', 0, "no-php-ini"),
	php.MakeOpt('H', 0, "hide-args"),
	php.MakeOpt('r', 1, "run"),
	php.MakeOpt('S', 1, "server"),
	php.MakeOpt('t', 1, "docroot"),
	php.MakeOpt('?', 0, "usage"),
	php.MakeOpt('v', 0, "version"),
	php.MakeOpt(15, 0, "ini"),
	php.MakeOpt(16, 0, "cgi"),
}

type runMode uint8

const (
	modeUnknown = iota
	modeHelp
	modeVersion
	modeInfo
	modeModules
	modeIni
	modeCliCode
	modeCliFile
	modeCliServer
	modeCgiServer
)

const (
	ok   = 0
	fail = 1
)

type OptArgs struct {
	mode runMode
	// ini opts
	IniPath   string
	IniIgnore bool
	IniAppend []string
	// cli opts
	ScriptFile string
	ScriptCode string
	hideArgv   bool
	// server opts
	Address      string
	DocumentRoot string
	// other
	RemainArgs []string
}

func parseArgs(args []string) (*OptArgs, error) {
	var optArgs OptArgs
	optsParser := php.NewOptsParser(args, Options, 1)
	err := optsParser.EachEx(true, func(opt *php.Opt, optArg string) error {
		switch opt.Char() {
		// ini opts
		case 'c':
			optArgs.IniPath = optArg
		case 'n':
			optArgs.IniIgnore = true
		case 'd':
			/* define ini __special__  entries on command line */
			optArgs.IniAppend = append(optArgs.IniAppend, optArg)
		// cli opts
		case 'f':
			if optArgs.mode == modeCliFile {
				return errors.New("You can use -f only once.\n")
			} else if optArgs.mode == modeCliCode {
				return errors.New("Either execute direct code, process stdin or use a file.\n")
			}
			optArgs.mode = modeCliFile
			optArgs.ScriptFile = optArg
		case 'r':
			if optArgs.mode == modeCliCode {
				return errors.New("You can use -r only once.\n")
			} else if optArgs.mode == modeCliFile {
				return errors.New("Either execute direct code, process stdin or use a file.\n")
			}
			optArgs.mode = modeCliCode
			optArgs.ScriptCode = optArg
		case 'H':
			optArgs.hideArgv = true
		// server opts
		case 'S':
			optArgs.Address = optArg
		case 't':
			optArgs.DocumentRoot = optArg
		// modes
		case 'v':
			optArgs.mode = modeVersion
			return lang.BreakErr
		case 'i':
			optArgs.mode = modeInfo
			return lang.BreakErr
		case 'm':
			optArgs.mode = modeModules
			return lang.BreakErr
		case 'h', '?':
			optArgs.mode = modeHelp
			return lang.BreakErr
		case 15:
			optArgs.mode = modeIni
			return lang.BreakErr
		case 16:
			optArgs.mode = modeCgiServer
		}
		return nil
	})
	if err == lang.BreakErr {
		err = nil
	}
	if err != nil {
		return nil, err
	}

	remainArgs := slices.Clone(optsParser.RemainArgs())
	if optArgs.mode == modeUnknown && len(remainArgs) > 0 && remainArgs[0] != "--" {
		optArgs.mode = modeCliFile
		optArgs.ScriptFile = remainArgs[0]
		remainArgs = remainArgs[1:]
	}

	optArgs.RemainArgs = remainArgs

	return &optArgs, nil
}

func Run(args []string) int {
	if len(args) == 0 {
		return fail
	}

	optArgs, err := parseArgs(args)
	if err != nil {
		fmt.Println(err.Error())
		showHelp()
		return fail
	}

	// prepare engine
	engine := php.NewEngine()

	switch optArgs.mode {
	case modeVersion:
		showVersion()
	case modeInfo:
		return showInfo(engine)
	case modeModules:
		return showModules(engine)
	case modeIni:
		return showIni(engine)
	case modeCliCode, modeCliFile:
		return RunCli(engine, optArgs)
	case modeCliServer, modeCgiServer:
		return RunServer(engine, optArgs)
	case modeHelp:
		fallthrough
	default:
		showHelp()
	}
	return ok
}

func showHelp() {
	fmt.Print(usage)
}

func showVersion() {
	fmt.Printf("gophp (php version %s)", "7.4.33")
}

func showInfo(engine *php.Engine) int {
	// todo show info
	return ok
}

func showIni(engine *php.Engine) int {
	// todo show ini
	return ok
}

func showModules(engine *php.Engine) int {
	fmt.Println("[PHP Modules]")
	// todo show modules
	fmt.Println("")
	return ok
}
