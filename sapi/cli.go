package sapi

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	"log"
)

func RunCli(engine *php.Engine, optArgs *OptArgs) int {
	fmt.Printf("run cli:%+v\n", optArgs)
	if optArgs.mode == modeCliFile {
		panic("todo run cli file")
	}

	fileHandle := php.NewFileHandleByString(optArgs.ScriptCode)
	ctx := engine.NewContext(nil, nil)
	retval, err := php.ExecuteScript(ctx, fileHandle, true)
	if err != nil {
		log.Println("Execute failed: " + err.Error())
		return fail
	}
	log.Printf("Execute succed, retval = %v", retval)

	return ok
}
