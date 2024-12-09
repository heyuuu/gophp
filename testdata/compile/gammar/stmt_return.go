package main

import (
	"github.com/heyuuu/gophp/php/def"
)

func init() {
	f := def.NewFile("", false)

	f.TopFn("", func(d def.TopDefiner) def.Val {

		return nil
	})
}
