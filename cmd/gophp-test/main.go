package main

import "github.com/heyuuu/gophp/tests"

func main() {
	conf := tests.DefaultConfig()
	conf.Verbose = true
	conf.SrcDir = "/Users/heyu/Code/src/php-7.4.33"
	conf.ExtDir = "/__ext__"
	conf.PhpBin = "/opt/homebrew/Cellar/php@7.4/7.4.33_6/bin/php"
	conf.PhpCgiBin = "/opt/homebrew/Cellar/php@7.4/7.4.33_6/bin/php-cgi"

	conf.DumpRoot = "/Users/heyu/Code/sik/gophp/log/dump"

	tests.Run(conf)
}

//func main() {
//	err := run(os.Args)
//	if err != nil {
//		log.Panicln(err)
//	}
//}

//func run(args []string) error {
//	var conf tests.Config
//	var logFile string
//
//	if len(args) > 1 {
//		flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)
//		flagSet.StringVar(&conf.SrcDir, "src-dir", "", "")
//		flagSet.IntVar(&conf.Limit, "limit", 0, "")
//		flagSet.IntVar(&conf.Workers, "j", 0, "")
//		flagSet.StringVar(&logFile, "log-file", "", "")
//
//		flagSet.StringVar(&conf.PhpBin, "php", "", "")
//		flagSet.StringVar(&conf.PhpCgiBin, "php-cgi", "", "")
//
//		err := flagSet.Parse(args[1:])
//		if err != nil {
//			return err
//		}
//	}
//	if conf.SrcDir == "" {
//		return errors.New(`the "--src-dir" must be specified`)
//	}
//
//	conf.Logger = eventHandler
//	defer eventCloser()
//
//	return tests.Run(conf)
//}
