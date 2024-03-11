package main

import (
	"errors"
	"flag"
	"github.com/heyuuu/gophp/tests"
	"log"
	"os"
)

func main() {
	err := run(os.Args)
	if err != nil {
		log.Panicln(err)
	}
}

func run(args []string) error {
	conf, err := parseConf(args)
	if err != nil {
		return err
	}

	if conf.SrcDir == "" {
		return errors.New(`the "--src-dir" must be specified`)
	}
	if conf.Logger == nil {
		return errors.New("conf.Logger 不可为 nil")
	}

	return tests.TestAll(conf)
}

func parseConf(args []string) (*tests.Config, error) {
	conf := tests.DefaultConfig()

	var logFile string
	var caseLogRoot string

	if len(args) > 1 {
		flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)
		flagSet.StringVar(&conf.SrcDir, "src-dir", "", "")
		flagSet.IntVar(&conf.Limit, "limit", 0, "")
		flagSet.IntVar(&conf.Workers, "j", 0, "")
		flagSet.StringVar(&conf.PhpBin, "php", "", "")
		flagSet.StringVar(&conf.PhpCgiBin, "php-cgi", "", "")
		flagSet.BoolVar(&conf.Verbose, "verbose", false, "")

		flagSet.StringVar(&logFile, "log-file", "", "")
		flagSet.StringVar(&caseLogRoot, "case-log-root", "", "")

		err := flagSet.Parse(args[1:])
		if err != nil {
			return nil, err
		}
	}
	if logFile != "" || caseLogRoot != "" {
		conf.Logger = tests.NewDumpLogger(logFile, caseLogRoot)
	}

	return conf, nil
}
