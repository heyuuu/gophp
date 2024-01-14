package main

import (
	_ "github.com/heyuuu/gophp/php/boot"
	"github.com/heyuuu/gophp/sapi"
	"os"
)

func main() {
	err := sapi.Run(os.Args)
	if err != nil {
		var code int
		if codeErr, ok := err.(interface{ Code() int }); ok {
			code = codeErr.Code()
		}
		if code == 0 {
			code = 1
		}
		os.Exit(code)
	}
}
