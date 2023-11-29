package main

import (
	"github.com/heyuuu/gophp/sapi"
	"os"
)

func main() {
	code := sapi.Run(os.Args)
	os.Exit(code)
}
