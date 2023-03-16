package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sik/internal/cmd/sikgen/zif"
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
		zif.RunGenFunc(opts.dir)
	case "clear-func":
		zif.RunClearFunc(opts.dir)
	case "":
		log.Fatalln("命令不可为空")
	default:
		log.Fatalln("未定义命令: " + opts.cmd)
	}
}

type opts struct {
	cmd string
	dir string
}

func parseOpts(args []string) opts {
	var cmd string
	var dir string

	// options
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.StringVar(&cmd, "cmd", "", "command")
	flagSet.StringVar(&dir, "d", "", "dir")
	_ = flagSet.Parse(args[1:])

	// workdir
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	return opts{
		cmd: cmd,
		dir: filepath.Join(wd, dir),
	}
}
