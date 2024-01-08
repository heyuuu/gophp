package main

import (
	"bufio"
	"flag"
	"fmt"
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
	var conf tests.Config
	var logFile string

	if len(args) > 1 {
		flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)
		flagSet.StringVar(&conf.SrcDir, "src-dir", "", "")
		flagSet.IntVar(&conf.Limit, "limit", 0, "")
		flagSet.IntVar(&conf.Workers, "j", 0, "")
		flagSet.StringVar(&logFile, "log-file", "", "")
		err := flagSet.Parse(args[1:])
		if err != nil {
			return err
		}
	}

	eventHandler, eventCloser := initEventHandler(logFile)
	conf.Events = eventHandler
	defer eventCloser()

	return tests.Run(conf)
}

func initEventHandler(file string) (handler tests.EventHandler, closer func()) {
	var logFile *os.File
	var logFileWriter *bufio.Writer
	if file != "" {
		var err error
		logFile, err = os.Create(file)
		if err != nil {
			log.Panicln(err)
		}
		logFileWriter = bufio.NewWriter(logFile)
	}

	handler = tests.NewParallelHandler(tests.NewDefaultEventHandler(func(verbose int, message string) {
		fmt.Print(message)
		if logFile != nil {
			logFileWriter.WriteString(message)
		}
	}))
	closer = func() {
		if logFile != nil {
			logFileWriter.Flush()
			logFile.Close()
		}
	}
	return
}
