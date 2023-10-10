package main

import (
	"flag"
	"fmt"
	"github.com/heyuuu/gophp/internal/cmd/sikgen/zif"
	"log"
	"os"
	"path/filepath"
)

func main() {
	run(os.Args)
}

func run(args []string) {
	if len(args) < 2 {
		log.Fatalln("Args 不可为空")
	}

	opts := parseOpts(args)
	fmt.Printf("%+v\n", opts)
	switch opts.cmd {
	case "gen-func":
		zif.RunGenFunc(opts.dir, opts.mode)
	case "clear-func":
		zif.RunClearFunc(opts.dir)
	case "":
		log.Fatalln("命令不可为空")
	default:
		log.Fatalln("未定义命令: " + opts.cmd)
	}
}

type optsType struct {
	cmd  string
	dir  string
	mode string
}

func parseOpts(args []string) (opts optsType) {
	// options
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.StringVar(&opts.cmd, "cmd", "", "command")
	flagSet.StringVar(&opts.dir, "d", "", "dir")
	flagSet.StringVar(&opts.mode, "mode", "file", "mode: file | pkg, default file")
	_ = flagSet.Parse(args[1:])

	// workdir
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	opts.dir = filepath.Join(wd, opts.dir)

	return
}
