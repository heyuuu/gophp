package main

import (
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/php/def"
)

func init() {
	f := def.NewFile("", false)
	f.TopFn("", func(d def.TopDefiner) def.Val {
		// lit
		d.Null()
		d.Bool(false)
		d.Bool(true)
		d.Int(1)
		d.Float(1.0)
		d.String("one string")
		d.String("another string")
		d.Array(
			d.ArrayItem(nil, d.Int(1), false, false),
			d.ArrayItem(nil, d.Int(2), false, false),
			d.ArrayItem(nil, d.Float(1.0), false, false),
			d.ArrayItem(nil, d.String("abc"), false, false),
		)
		// variable
		d.VarValue("var")
		d.VarValueEx(d.VarValue("var"))
		// index
		d.Index(d.Index(d.Index(d.VarValue("var"), d.String("a")), d.String("b")), d.String("c"))
		// cast
		d.Cast(ast.CastBool, d.VarValue("var"))
		d.Cast(ast.CastInt, d.VarValue("var"))
		d.Cast(ast.CastDouble, d.VarValue("var"))
		d.Cast(ast.CastString, d.VarValue("var"))
		d.Cast(ast.CastArray, d.VarValue("var"))
		d.Cast(ast.CastObject, d.VarValue("var"))
		d.Cast(ast.CastUnset, d.VarValue("var"))
		// new
		d.New("\\stdObject")
		d.NewEx(d.VarValue("class"))
		// unary op


		return nil
	})
}
