package main

import (
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/php/def"
)

func init() {
	f := def.NewFile("", false)
	f.TopFn("", func(d def.TopDefiner) def.Val {
		d.Use(ast.UseNormal, `anprefix\ClassA`, `A`)
		d.Use(ast.UseNormal, `anprefix\ClassB`, "")
		d.Use(ast.UseFunction, `anprefix\functionC`, `C`)
		d.Use(ast.UseConstant, `anprefix\ConstD`, `D`)

		return nil
	})
}
