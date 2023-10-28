package main

import (
	"flag"
	"github.com/heyuuu/gophp/php"
	"log"
)

func main() {
	var code string
	flag.StringVar(&code, "r", "", "code")
	flag.Parse()

	fileHandle := php.NewFileHandleByString(code)

	engine := php.NewEngine()
	ctx := engine.NewContext()
	retval, err := php.ExecuteScript(ctx, fileHandle, true)
	if err != nil {
		panic("Execute failed: " + err.Error())
	}
	log.Printf("Execute succed, retval = %v", retval)
}
