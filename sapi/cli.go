package sapi

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	"log"
)

func (c *Cmd) runCli(engine *php.Engine, optArgs *OptArgs) (err error) {
	ctx := engine.NewContext(nil, nil)
	engine.HandleContext(ctx, func(ctx *php.Context) {
		// stdio
		if c.Stdout != nil {
			ctx.OG().PushHandler(c.Stdout)
		}

		// file handle
		var fileHandle *php.FileHandle
		var skipShebang bool
		if optArgs.mode == modeCliFile {
			fileHandle, err = php.NewFileHandleByFilename(optArgs.ScriptFile)
			if err != nil {
				log.Panicln("Could not open input file: " + optArgs.ScriptFile)
			}
			skipShebang = false
		} else {
			fileHandle = php.NewFileHandleByCommandLine(optArgs.ScriptCode)
			skipShebang = true
		}

		// run
		_, err = php.ExecuteScript(ctx, fileHandle, skipShebang)
		if err != nil {
			err = withCode(1, fmt.Errorf("Execute failed: %w", err))
		}
	})
	return
}
