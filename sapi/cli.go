package sapi

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	"log"
)

func RunCli(engine *php.Engine, optArgs *OptArgs) error {
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
	_, err = php.ExecuteScript(ctx, fileHandle, skipShebang)
	if err != nil {
		return withCode(1, fmt.Errorf("Execute failed: %w", err))
	}

	return nil
}
