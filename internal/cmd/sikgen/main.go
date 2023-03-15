package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Args 不可为空")
	}

	opts := parseOpts()
	switch opts.cmd {
	case "gen-func":
		runGenFunc(opts.dir)
	case "clear-func":
		runClearFunc(opts.dir)
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

func parseOpts() opts {
	var cmd string
	var dir string

	// arguments
	if len(os.Args) >= 2 {
		cmd = os.Args[1]
	}

	// options
	flag.StringVar(&dir, "d", "", "mode")
	flag.Parse()

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
