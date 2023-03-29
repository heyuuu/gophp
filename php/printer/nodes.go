package printer

import (
	"fmt"
	"gophp/php/ast"
)

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
	case *ast.ArrayExpr:
	case *ast.ArrayItemExpr:
	case *ast.ClosureExpr:
	case *ast.ClosureUseExpr:
	case *ast.ArrowFunctionExpr:
	case *ast.IndexExpr:
	case *ast.CastExpr:
	case *ast.UnaryExpr:
	case *ast.BinaryExpr:
	case *ast.AssignExpr:
	case *ast.AssignRefExpr:
	case *ast.InternalCallExpr:
	case *ast.CloneExpr:
	case *ast.ErrorSuppressExpr:
	case *ast.ExitExpr:
	case *ast.ConstFetchExpr:
	case *ast.ClassConstFetchExpr:
	case *ast.MagicConstExpr:
	case *ast.MatchExpr:
	case *ast.InstanceofExpr:
	case *ast.ListExpr:
	case *ast.PrintExpr:
	case *ast.PropertyFetchExpr:
	case *ast.NullsafePropertyFetchExpr:
	case *ast.StaticPropertyFetchExpr:
	case *ast.ShellExecExpr:
	case *ast.TernaryExpr:
	case *ast.ThrowExpr:
	case *ast.VariableExpr:
	case *ast.YieldExpr:
	case *ast.YieldFromExpr:
	case *ast.FuncCallExpr:
	case *ast.NewExpr:
	case *ast.MethodCallExpr:
	case *ast.NullsafeMethodCallExpr:
	case *ast.StaticCallExpr:
	default:
		panic("unreachable")
	}
}

func (p *printer) stmt(n ast.Stmt) {
	switch n.(type) {
	case *ast.EmptyStmt:
		p.print(';')
	case *ast.BlockStmt:
	case *ast.ExprStmt:
	case *ast.ReturnStmt:
	case *ast.LabelStmt:
	case *ast.GotoStmt:
	case *ast.IfStmt:
	case *ast.ElseIfStmt:
	case *ast.ElseStmt:
	case *ast.SwitchStmt:
	case *ast.CaseStmt:
	case *ast.ForStmt:
	case *ast.ForeachStmt:
	case *ast.BreakStmt:
	case *ast.ContinueStmt:
	case *ast.WhileStmt:
	case *ast.DoStmt:
	case *ast.TryCatchStmt:
	case *ast.CatchStmt:
	case *ast.FinallyStmt:
	case *ast.ConstStmt:
	case *ast.EchoStmt:
	case *ast.GlobalStmt:
	case *ast.HaltCompilerStmt:
	case *ast.InlineHTMLStmt:
	case *ast.StaticStmt:
	case *ast.StaticVarStmt:
	case *ast.UnsetStmt:
	case *ast.UseStmt:
	case *ast.DeclareStmt:
	case *ast.DeclareDeclareStmt:
	case *ast.NamespaceStmt:
	case *ast.FunctionStmt:
	case *ast.InterfaceStmt:
	case *ast.ClassStmt:
	case *ast.ClassConstStmt:
	case *ast.PropertyStmt:
	case *ast.PropertyPropertyStmt:
	case *ast.ClassMethodStmt:
	case *ast.TraitStmt:
	case *ast.TraitUseStmt:
	case *ast.TraitUseAdaptationAliasStmt:
	case *ast.TraitUseAdaptationPrecedenceStmt:
	case *ast.EnumStmt:
	case *ast.EnumCaseStmt:
	}
}
