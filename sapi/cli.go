package sapi

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	"log"
)

func RunCli(engine *php.Engine, optArgs *OptArgs) int {
	fmt.Printf("run cli:%+v\n", optArgs)
	var fileHandle *php.FileHandle
	var err error
	var skipShebang bool

	if optArgs.mode == modeCliFile {
		fileHandle, err = php.NewFileHandleByFilename(optArgs.ScriptFile)
		if err != nil {
			log.Panicln("Could not open input file: " + optArgs.ScriptFile)
		}
		skipShebang = false
	} else {
		fileHandle = php.NewFileHandleByString(optArgs.ScriptCode)
		skipShebang = true
	}

	ctx := engine.NewContext(nil, nil)
	retval, err := php.ExecuteScript(ctx, fileHandle, skipShebang)
	if err != nil {
		log.Println("Execute failed: " + err.Error())
		return fail
	}
	log.Printf("Execute succed, retval = %v", retval)

	return ok
}
