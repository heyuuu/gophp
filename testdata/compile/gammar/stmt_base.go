package main

import "github.com/heyuuu/gophp/php/def"

func init() {
	f := def.NewFile("", false)

	f.TopFn("", func(d def.TopDefiner) def.Val {
		// empty stmt
		// expr stmt
		d.Int(1)
		// const stmt
		d.DeclConst("a", d.Int(1))
		// echo stmt
		d.Echo("abc")
		d.EchoVal(d.VarValue("a"))
		// global stmt
		d.DeclGlobal("a")
		d.DeclGlobalVal(d.VarValue("a"))
		// static var stmt
		d.DeclStatic("a", nil)
		d.DeclStatic("b", d.Int(1))
		// unset stmt
		d.Unset(d.Var("a"))

		// if stmt
		if d.IsTrue(d.Equal(d.VarValue("a"), d.Int(1))) {
			// branch if
			d.Echo("if")
		} else if d.IsTrue(d.Equal(d.VarValue("b"), d.Int(2))) {
			// branch elseif
			d.Echo("elseif")
		} else {
			// branch else
			d.Echo("else")
		}

		// switch stmt
		switch switchVar := d.Equal(d.VarValue("a"), d.Int(2)); {
		case d.IsTrue(d.Equal(switchVar, d.Int(1))):
			// branch case 1
			break
		case d.IsTrue(d.Equal(switchVar, d.Int(2))):
			// branch case 2
			fallthrough
		case d.IsTrue(d.Equal(switchVar, d.Int(2))):
			// branch case 2
			break
		default:
			// branch default
		}

		// for stmt
		for d.Assign(d.Var("i"), d.Int(0)); d.IsTrue(d.Smaller(d.VarValue("i"), d.Int(10))); d.PostInc(d.Var("i")) {
			// branch for
			d.Echo("for")
		}

		// foreach stmt
		for iter := d.ForeachIterator(d.VarValue("arr")); iter.Valid(); iter.Next() {
			// todo
			d.Echo("foreach")
		}

		// while stmt
	_break_0:
		for {
			for {
				goto _break_0
			}
		}

		// do-while stmt
		for {
			continue
			break
		}

		return nil
	})
}
