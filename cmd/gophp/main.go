package main

import (
	_ "github.com/heyuuu/gophp/php/boot"
	"github.com/heyuuu/gophp/sapi"
	"os"
)

func main() {
	code := sapi.Run(os.Args)
	os.Exit(code)
}
