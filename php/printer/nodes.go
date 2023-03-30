package printer

import (
	"fmt"
	"github.com/heyuuu/gophp/php/ast"
	"github.com/heyuuu/gophp/php/token"
)

func (p *printer) arg(n *ast.Arg) {
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

func (p *printer) param(n *ast.Param) {
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
func (p *printer) typeHint(n ast.Type) {
	p.typeHint0(n, false)
}

func (p *printer) typeHint0(n ast.Type, wrap bool) {
	switch t := n.(type) {
	case *ast.SimpleType:
		p.print(t.Name)
	case *ast.NullableType:
		p.print('?')
		p.typeHint0(t.Type, true)
	case *ast.IntersectionType:
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
	case *ast.UnionType:
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

func (p *printer) expr(n ast.Expr) {
	switch x := n.(type) {
	case *ast.IntLit:
		p.print(x.Value)
	case *ast.FloatLit:
		p.print(fmt.Printf("%f", x.Value))
	case *ast.StringLit:
		// todo escape
		p.print("\"", x.Value, "\"")
	case *ast.ArrayExpr:
		p.print("[")
		printList(p, x.Items, ", ")
		p.print("]")
	case *ast.ArrayItemExpr:
		if x.Key != nil {
			p.print(x.Key, " => ")
		}
		p.print(x.Value)
	case *ast.ClosureExpr:
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
	case *ast.ClosureUseExpr:
		if x.ByRef {
			p.print("&")
		}
		p.print(x.Var)
	case *ast.ArrowFunctionExpr:
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
	case *ast.IndexExpr:
		p.print(x.Var, "[", x.Dim, "]")
	case *ast.CastExpr:
		p.print(x.Op, x.Expr)
	case *ast.UnaryExpr:
		switch x.Kind {
		case token.PostInc, token.PostDec:
			p.print(x.Var, x.Kind)
		default:
			p.print(x.Kind, x.Var)
		}
	case *ast.BinaryExpr:
		p.print(x.Left, " ", x.Op, " ", x.Right)
	case *ast.AssignExpr:
		p.print(x.Var, " ", x.Op, " ", x.Expr)
	case *ast.AssignRefExpr:
		p.print(x.Var, " = &", x.Expr)
	case *ast.InternalCallExpr:
		p.print(x.Kind)
	case *ast.CloneExpr:
		p.print("clone ", x.Expr)
	case *ast.ErrorSuppressExpr:
		p.print("@", x.Expr)
	case *ast.ExitExpr:
		p.print("exist(", x.Expr, ")")
	case *ast.ConstFetchExpr:
		p.print(x.Name)
	case *ast.ClassConstFetchExpr:
		p.print(x.Class, "::", x.Name)
	case *ast.MagicConstExpr:
		p.print(x.Kind)
	case *ast.MatchExpr:
		p.print("match (", x.Cond, ") {\n")
		p.indent++
		for _, arm := range x.Arms {
			if len(arm.Conds) != 0 {
				p.print(arm.Conds, " => ", arm.Body, "\n")
			} else {
				p.print("default => ", arm.Body, "\n")
			}
		}
		p.indent--
		p.print("}")
	case *ast.InstanceofExpr:
		p.print(x.Expr, " instanceOf ", x.Class)
	case *ast.ListExpr:
		p.print("list(")
		printList(p, x.Items, ", ")
		p.print(")")
	case *ast.PrintExpr:
		p.print("print ", x.Expr)
	case *ast.PropertyFetchExpr:
		p.print(x.Var, "->", x.Name)
	case *ast.NullsafePropertyFetchExpr:
		p.print(x.Var, "?->", x.Name)
	case *ast.StaticPropertyFetchExpr:
		p.print(x.Class, "::", x.Name)
	case *ast.ShellExecExpr:
		p.print('`')
		printList(p, x.Parts, "")
		p.print('`')
	case *ast.TernaryExpr:
		if x.If == nil {
			p.print(x.Cond, " ?: ", x.Else)
		} else {
			p.print(x.Cond, " ? ", x.If, " : ", x.Else)
		}
	case *ast.ThrowExpr:
		p.print("throw ", x.Expr)
	case *ast.VariableExpr:
		p.print("$", x.Name)
	case *ast.YieldExpr:
		if x.Key == nil {
			p.print("yield ", x.Value)
		} else {
			p.print("yield ", x.Key, " => ", x.Value)
		}
	case *ast.YieldFromExpr:
		p.print("yield from ", x.Expr)
	case *ast.FuncCallExpr:
		p.print(x.Name, "(", x.Args, ")")
	case *ast.NewExpr:
		p.print("new ", x.Class, "(", x.Args, ")")
	case *ast.MethodCallExpr:
		p.print(x.Var, "->", x.Name, "(", x.Args, ")")
	case *ast.NullsafeMethodCallExpr:
		p.print(x.Var, "?->", x.Name, "(", x.Args, ")")
	case *ast.StaticCallExpr:
		p.print(x.Class, "::", x.Name, "(", x.Args, ")")
	default:
		panic("unreachable")
	}
}

func (p *printer) stmt(n ast.Stmt) {
	switch x := n.(type) {
	case *ast.EmptyStmt:
		//p.print(";")
	case *ast.BlockStmt:
		p.stmtList(x.List, false)
	case *ast.ExprStmt:
		p.print(x.Expr, ";")
	case *ast.ReturnStmt:
		if x.Expr == nil {
			p.print("return;")
		} else {
			p.print("return ", x.Expr, ";")
		}
	case *ast.LabelStmt:
		p.print(x.Name, ":")
	case *ast.GotoStmt:
		p.print("goto ", x.Name, ";")
	case *ast.IfStmt:
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
	case *ast.SwitchStmt:
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
	case *ast.ForStmt:
		p.print("for (", x.Init, ";", x.Cond, ";", x.Loop, ") {\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ast.ForeachStmt:
		if x.KeyVar != nil {
			p.print("foreach (", x.KeyVar, " as ", x.KeyVar, " => ", x.ValueVar, ") {\n")
		} else {
			p.print("foreach (", x.KeyVar, " as ", x.ValueVar, ") {\n")
		}
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ast.BreakStmt:
		if x.Num != nil {
			p.print("break ", x.Num, ";")
		} else {
			p.print("break;")
		}
	case *ast.ContinueStmt:
		if x.Num != nil {
			p.print("continue ", x.Num, ";")
		} else {
			p.print("continue;")
		}
	case *ast.WhileStmt:
		p.print("while (", x.Cond, ") {\n", x.Stmts, "}")
	case *ast.DoStmt:
		p.print("do {\n", x.Stmts, "} while (", x.Cond, ");")
	case *ast.TryCatchStmt:
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
	case *ast.ConstStmt:
		// todo
	case *ast.EchoStmt:
		p.print("echo ", x.Exprs, ";")
	case *ast.GlobalStmt:
		p.print("global ", x.Vars, ";")
	case *ast.HaltCompilerStmt:
		p.print("__halt_compiler();", x.Remaining)
	case *ast.InlineHTMLStmt:
		p.print("?>", x.Value, "<?php")
	case *ast.StaticStmt:
		p.print("static ", x.Vars, ";")
	case *ast.StaticVarStmt:
		if x.Default != nil {
			p.print(x.Var, " = ", x.Default)
		} else {
			p.print(x.Var)
		}
	case *ast.UnsetStmt:
		p.print("unset(", x.Vars, ")")
	case *ast.UseStmt:
		var useType string
		switch x.Type {
		case ast.UseFunction:
			useType = "function "
		case ast.UseConstant:
			useType = "const "
		}

		if x.Alias != nil {
			p.print("use ", useType, x.Name, " as ", x.Alias, ";")
		} else {
			p.print("use ", useType, x.Name, ";")
		}
	case *ast.DeclareStmt:
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
	case *ast.DeclareDeclareStmt:
		p.print(x.Key, "=", x.Value)
	case *ast.NamespaceStmt:
		p.print("namespace ", x.Name, ";\n")
		p.stmtList(x.Stmts, false)
	case *ast.FunctionStmt:
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
	case *ast.InterfaceStmt:
		p.print("interface ", x.Name)
		if len(x.Extends) != 0 {
			p.print(" extends ", x.Extends)
		}
		p.print("\n{\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ast.ClassStmt:
		if x.Flags != 0 {
			p.flags(x.Flags)
			p.print(" ")
		}
		p.print("class ", x.Name)
		if x.Extends != nil {
			p.print(" extends ", x.Extends)
		}
		if len(x.Implements) != 0 {
			p.print(" implements ", x.Implements)
		}
		p.print("\n{\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ast.ClassConstStmt:
		for _, c := range x.Consts {
			if x.Flags != 0 {
				p.flags(x.Flags)
				p.print(" ")
			}
			p.print("const ", c.Name, " = ", c.Value)
		}
	case *ast.PropertyStmt:
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
	case *ast.ClassMethodStmt:
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
	case *ast.TraitStmt:
		p.print("class ", x.Name, "\n{\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ast.TraitUseStmt:
		if len(x.Adaptations) != 0 {
			p.print("use ", x.Traits, " {\n")
			p.indent++
			p.print(x.Adaptations)
			p.indent--
			p.print("};")
		} else {
			p.print("use ", x.Traits, ";")
		}
	case *ast.TraitUseAdaptationPrecedenceStmt:
		p.print(x.Trait, "::", x.Method, " insteadof ", x.Insteadof, ";")
	case *ast.TraitUseAdaptationAliasStmt:
		p.print(x.Trait, "::", x.Method, " as")
		if x.NewModifier != 0 {
			p.print(" ")
			p.flags(x.NewModifier)
		}
		if x.NewName != nil {
			p.print(" ", x.NewName)
		}
		p.print(";")
	case *ast.EnumStmt:
		p.print("enum ", x.Name)
		if x.ScalarType != nil {
			p.print(": ", x.ScalarType)
		}
		if len(x.Implements) != 0 {
			p.print(" implements ", x.Implements)
		}
		p.print("\n{\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ast.EnumCaseStmt:
		if x.Expr != nil {
			p.print("case ", x.Name, " = ", x.Expr, ";")
		} else {
			p.print("case ", x.Name, ";")
		}
	}
}
