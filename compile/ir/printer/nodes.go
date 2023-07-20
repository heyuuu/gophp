package printer

import (
	"fmt"
	"github.com/heyuuu/gophp/compile/ir"
	"github.com/heyuuu/gophp/php/token"
	"log"
)

func (p *printer) arg(n *ir.Arg) {
	if n.Name != nil {
		p.print(n.Name, ": ")
	}
	if n.ByRef {
		p.print("&")
	}
	if n.Unpack {
		p.print("...")
	}
	p.print(n.Value)
}

func (p *printer) param(n *ir.Param) {
	if n.Flags != 0 {
		p.flags(n.Flags)
		p.print(" ")
	}
	if n.Type != nil {
		p.print(n.Type, " ")
	}
	if n.ByRef {
		p.print("&")
	}
	if n.Variadic {
		p.print("...")
	}
	p.print(n.Var)
	if n.Default != nil {
		p.print(" = ", n.Default)
	}
}
func (p *printer) typeHint(n ir.Type) {
	p.typeHint0(n, false)
}

func (p *printer) typeHint0(n ir.Type, wrap bool) {
	switch t := n.(type) {
	case *ir.SimpleType:
		p.print(t.Name)
	case *ir.NullableType:
		p.print('?')
		p.typeHint0(t.Type, true)
	case *ir.IntersectionType:
		if wrap {
			p.print('(')
		}
		for i, typ := range t.Types {
			if i != 0 {
				p.print('&')
			}
			p.typeHint0(typ, true)
		}
		if wrap {
			p.print(')')
		}
	case *ir.UnionType:
		if wrap {
			p.print('(')
		}
		for i, typ := range t.Types {
			if i != 0 {
				p.print('|')
			}
			p.typeHint0(typ, true)
		}
		if wrap {
			p.print(')')
		}
	default:
		panic("unreachable")
	}
}

func (p *printer) expr(n ir.Expr) {
	switch x := n.(type) {
	case *ir.IntLit:
		p.print(x.Value)
	case *ir.FloatLit:
		p.print(fmt.Printf("%f", x.Value))
	case *ir.StringLit:
		// todo escape
		p.print("\"", x.Value, "\"")
	case *ir.ArrayExpr:
		p.print("[")
		printList(p, x.Items, ", ")
		p.print("]")
	case *ir.ArrayItemExpr:
		if x.Key != nil {
			p.print(x.Key, " => ")
		}
		p.print(x.Value)
	case *ir.ClosureExpr:
		p.print("function (")
		printList(p, x.Params, ", ")
		p.print(") ")
		if len(x.Uses) != 0 {
			p.print("use (")
			printList(p, x.Uses, ", ")
			p.print(") ")
		}
		p.print("{\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ir.ClosureUseExpr:
		if x.ByRef {
			p.print("&")
		}
		p.print(x.Var)
	case *ir.ArrowFunctionExpr:
		if x.Static {
			p.print("static ")
		}
		p.print("fn")
		if x.ByRef {
			p.print("&")
		}
		p.print("(")
		printList(p, x.Params, ", ")
		p.print(")")
		if x.ReturnType != nil {
			p.print(": ", x.ReturnType)
		}
		p.print(" => ", x.Expr)
	case *ir.IndexExpr:
		p.print(x.Var, "[", x.Dim, "]")
	case *ir.CastExpr:
		p.print(x.Op, x.Expr)
	case *ir.UnaryExpr:
		switch x.Kind {
		case token.PostInc, token.PostDec:
			p.print(x.Var, x.Kind)
		default:
			p.print(x.Kind, x.Var)
		}
	case *ir.BinaryExpr:
		p.print(x.Left, " ", x.Op, " ", x.Right)
	case *ir.AssignExpr:
		p.print(x.Var, " ", x.Op, " ", x.Expr)
	case *ir.AssignRefExpr:
		p.print(x.Var, " = &", x.Expr)
	case *ir.InternalCallExpr:
		p.print(x.Kind)
	case *ir.CloneExpr:
		p.print("clone ", x.Expr)
	case *ir.ErrorSuppressExpr:
		p.print("@", x.Expr)
	case *ir.ExitExpr:
		p.print("exist(", x.Expr, ")")
	case *ir.ConstFetchExpr:
		p.print(x.Name)
	case *ir.ClassConstFetchExpr:
		p.print(x.Class, "::", x.Name)
	case *ir.MagicConstExpr:
		p.print(x.Kind)
	case *ir.InstanceofExpr:
		p.print(x.Expr, " instanceOf ", x.Class)
	case *ir.ListExpr:
		p.print("list(")
		printList(p, x.Items, ", ")
		p.print(")")
	case *ir.PrintExpr:
		p.print("print ", x.Expr)
	case *ir.PropertyFetchExpr:
		p.print(x.Var, "->", x.Name)
	case *ir.NullsafePropertyFetchExpr:
		p.print(x.Var, "?->", x.Name)
	case *ir.StaticPropertyFetchExpr:
		p.print(x.Class, "::", x.Name)
	case *ir.ShellExecExpr:
		p.print('`')
		printList(p, x.Parts, "")
		p.print('`')
	case *ir.TernaryExpr:
		if x.If == nil {
			p.print(x.Cond, " ?: ", x.Else)
		} else {
			p.print(x.Cond, " ? ", x.If, " : ", x.Else)
		}
	case *ir.ThrowExpr:
		p.print("throw ", x.Expr)
	case *ir.VariableExpr:
		p.print("$", x.Name)
	case *ir.YieldExpr:
		if x.Key == nil {
			p.print("yield ", x.Value)
		} else {
			p.print("yield ", x.Key, " => ", x.Value)
		}
	case *ir.YieldFromExpr:
		p.print("yield from ", x.Expr)
	case *ir.FuncCallExpr:
		p.print(x.Name, "(", x.Args, ")")
	case *ir.NewExpr:
		p.print("new ", x.Class, "(", x.Args, ")")
	case *ir.MethodCallExpr:
		p.print(x.Var, "->", x.Name, "(", x.Args, ")")
	case *ir.NullsafeMethodCallExpr:
		p.print(x.Var, "?->", x.Name, "(", x.Args, ")")
	case *ir.StaticCallExpr:
		p.print(x.Class, "::", x.Name, "(", x.Args, ")")
	default:
		panic("unreachable")
	}
}

func (p *printer) stmt(n ir.Stmt) {
	switch x := n.(type) {
	case *ir.EmptyStmt:
		//p.print(";")
	case *ir.BlockStmt:
		p.stmtList(x.List, false)
	case *ir.ExprStmt:
		p.print(x.Expr, ";")
	case *ir.ReturnStmt:
		if x.Expr == nil {
			p.print("return;")
		} else {
			p.print("return ", x.Expr, ";")
		}
	case *ir.LabelStmt:
		p.print(x.Name, ":")
	case *ir.GotoStmt:
		p.print("goto ", x.Name, ";")
	case *ir.IfStmt:
		p.print("if (", x.Cond, ") {\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
		for _, elseif := range x.Elseifs {
			p.print(" elseif (", elseif.Cond, ") {\n")
			p.stmtList(elseif.Stmts, true)
			p.print("}")
		}
		if x.Else != nil {
			p.print(" else {\n")
			p.stmtList(x.Else.Stmts, true)
			p.print("}")
		}
	case *ir.SwitchStmt:
		p.print("switch (", x.Cond, ") {\n")
		for _, caseStmt := range x.Cases {
			if caseStmt.Cond != nil {
				p.print("case ", caseStmt.Cond, ":\n")
				p.stmtList(caseStmt.Stmts, true)
			} else {
				p.print("default:\n")
				p.stmtList(caseStmt.Stmts, true)
			}
		}
		p.print("}")
	case *ir.ForStmt:
		p.print("for (", x.Init, ";", x.Cond, ";", x.Loop, ") {\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ir.ForeachStmt:
		if x.KeyVar != nil {
			p.print("foreach (", x.KeyVar, " as ", x.KeyVar, " => ", x.ValueVar, ") {\n")
		} else {
			p.print("foreach (", x.KeyVar, " as ", x.ValueVar, ") {\n")
		}
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ir.BreakStmt:
		if x.Num != nil {
			p.print("break ", x.Num, ";")
		} else {
			p.print("break;")
		}
	case *ir.ContinueStmt:
		if x.Num != nil {
			p.print("continue ", x.Num, ";")
		} else {
			p.print("continue;")
		}
	case *ir.WhileStmt:
		p.print("while (", x.Cond, ") {\n", x.Stmts, "}")
	case *ir.DoStmt:
		p.print("do {\n", x.Stmts, "} while (", x.Cond, ");")
	case *ir.TryCatchStmt:
		p.print("try {\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
		for _, catch := range x.Catches {
			p.print(" catch (")
			printList(p, catch.Types, "|")
			p.print(" ", catch.Var, ") {\n")
			p.stmtList(catch.Stmts, true)
			p.print("}")
		}
		if x.Finally != nil {
			p.print(" finally {\n", x.Finally.Stmts, "}")
		}
	case *ir.ConstStmt:
		// todo
	case *ir.EchoStmt:
		p.print("echo ", x.Exprs, ";")
	case *ir.GlobalStmt:
		p.print("global ", x.Vars, ";")
	case *ir.HaltCompilerStmt:
		p.print("__halt_compiler();", x.Remaining)
	case *ir.InlineHTMLStmt:
		p.print("?>", x.Value, "<?php")
	case *ir.StaticStmt:
		p.print("static ", x.Vars, ";")
	case *ir.StaticVarStmt:
		if x.Default != nil {
			p.print(x.Var, " = ", x.Default)
		} else {
			p.print(x.Var)
		}
	case *ir.UnsetStmt:
		p.print("unset(", x.Vars, ")")
	case *ir.UseStmt:
		var useType string
		switch x.Type {
		case ir.UseFunction:
			useType = "function "
		case ir.UseConstant:
			useType = "const "
		}

		if x.Alias != nil {
			p.print("use ", useType, x.Name, " as ", x.Alias, ";")
		} else {
			p.print("use ", useType, x.Name, ";")
		}
	case *ir.DeclareStmt:
		p.print("declare(")
		printList(p, x.Declares, ", ")
		p.print(")")
		if len(x.Stmts) == 0 {
			p.print(";")
		} else {
			p.print("{\n")
			p.stmtList(x.Stmts, true)
			p.print("}")
		}
	case *ir.DeclareDeclareStmt:
		p.print(x.Key, "=", x.Value)
	case *ir.NamespaceStmt:
		p.print("namespace ", x.Name, ";\n")
		p.stmtList(x.Stmts, false)
	case *ir.InitStmt:
		p.print("// init\n")
		if len(x.Stmts) == 0 {
			p.print("func init() {}")
		} else {
			p.print("func init() {\n")
			p.stmtList(x.Stmts, true)
			p.print("}\n")
		}
	case *ir.FunctionStmt:
		p.print("// ", x.Name, "\n")
		p.print("func ")
		if x.ByRef {
			p.print("&")
		}
		p.print(x.Name, "(", x.Params, ")")
		if x.ReturnType != nil {
			p.print(": ", x.ReturnType)
		}
		if len(x.Stmts) == 0 {
			p.print(" {}\n")
		} else {
			p.print(" {\n")
			p.stmtList(x.Stmts, true)
			p.print("}\n")
		}
	case *ir.InterfaceStmt:
		p.print("interface ", x.Name)
		if len(x.Extends) != 0 {
			p.print(" extends ", x.Extends)
		}
		p.print("\n{\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ir.ClassStmt:
		p.print("// class ", x.Name, "\n")
		if x.Flags != 0 {
			p.flags(x.Flags)
			p.print(" ")
		}
		p.print("class ", x.Name)
		if x.Extends != nil {
			p.print(" < ", x.Extends)
		}
		if len(x.Implements) != 0 {
			p.print(" : ", x.Implements)
		}
		p.print(" {\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ir.ClassConstStmt:
		for _, c := range x.Consts {
			if x.Flags != 0 {
				p.flags(x.Flags)
				p.print(" ")
			}
			p.print("const ", c.Name, " = ", c.Value)
		}
	case *ir.PropertyStmt:
		for _, prop := range x.Props {
			if x.Flags != 0 {
				p.flags(x.Flags)
				p.print(" ")
			}
			if x.Type != nil {
				p.print(x.Type, " ")
			}
			if prop.Default != nil {
				p.print(prop.Name, " = ", prop.Default, ";")
			} else {
				p.print(prop.Name, ";")
			}
		}
	case *ir.ClassMethodStmt:
		if x.Flags != 0 {
			p.flags(x.Flags)
			p.print(" ")
		}
		p.print("function ")
		if x.ByRef {
			p.print("&")
		}
		p.print(x.Name, "(", x.Params, ")")
		if x.ReturnType != nil {
			p.print(": ", x.ReturnType)
		}
		p.print("\n{\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ir.TraitStmt:
		p.print("trait ", x.Name, "\n{\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ir.TraitUseStmt:
		if len(x.Adaptations) != 0 {
			p.print("use ", x.Traits, " {\n")
			p.indent++
			p.print(x.Adaptations)
			p.indent--
			p.print("};")
		} else {
			p.print("use ", x.Traits, ";")
		}
	case *ir.TraitUseAdaptationPrecedenceStmt:
		p.print(x.Trait, "::", x.Method, " insteadof ", x.Insteadof, ";")
	case *ir.TraitUseAdaptationAliasStmt:
		p.print(x.Trait, "::", x.Method, " as")
		if x.NewModifier != 0 {
			p.print(" ")
			p.flags(x.NewModifier)
		}
		if x.NewName != nil {
			p.print(" ", x.NewName)
		}
		p.print(";")
	default:
		log.Fatalf("unsupported type of ir.stmt: %v", x)
	}
}
