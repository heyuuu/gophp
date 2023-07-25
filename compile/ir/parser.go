package ir

import (
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/utils/slices"
	"log"
	"reflect"
)

func ParseAstFile(astFile []ast.Stmt) (file *File, err error) {
	defer func() {
		switch e := recover().(type) {
		case nil:
			return
		case parsingError:
			err = e
		default:
			panic(e)
		}
	}()

	p := &parser{}
	return p.ParseFile(astFile), nil
}

// parsingError
type parsingError string

func (p parsingError) Error() string { return string(p) }

// parsingStmts
type parsingStmts []Stmt

func (p parsingStmts) node()     {}
func (p parsingStmts) stmtNode() {}

// parser
type parser struct {
	file *File
}

func (p *parser) clean() {
	p.file = nil
}

func (p *parser) ParseFile(astFile []ast.Stmt) *File {
	// 拆分 declare 语句、全局代码和命名空间代码
	var declareStmts []*ast.DeclareStmt
	var globalStmts []ast.Stmt
	var namespaceStmts []*ast.NamespaceStmt
	for _, astStmt := range astFile {
		switch s := astStmt.(type) {
		case *ast.DeclareStmt:
			declareStmts = append(declareStmts, s)
		case *ast.NamespaceStmt:
			namespaceStmts = append(namespaceStmts, s)
		default:
			globalStmts = append(globalStmts, s)
		}
	}
	p.assert(len(globalStmts) == 0 || len(namespaceStmts) == 0, "Global code should be enclosed in global namespace declaration")

	// 初始化对象
	p.file = &File{}
	defer func() { p.file = nil }()

	// 优先处理 declare，会影响其他代码
	slices.Each(declareStmts, p.handleDeclareStmt)

	// 区分有无命名空间进行处理
	if len(globalStmts) > 0 {
		p.file.Namespaces = []*Namespace{p.pNamespace("", globalStmts)}
	} else {
		p.file.Namespaces = slices.Map(namespaceStmts, p.pNamespaceStmt)
	}

	return p.file
}

// misc
func (p *parser) fail(message string) {
	panic(parsingError(message))
}
func (p *parser) assert(cond bool, message string) {
	if !cond {
		log.Println(message)
		p.fail(message)
	}
}
func (p *parser) highVersionFeature(feature string) {
	p.fail("high version php feature: " + feature)
}
func (p *parser) lowerVersionFeature(feature string) {
	p.fail("lower version php feature: " + feature)
}
func (p *parser) unsupported(message string) {
	p.fail(message)
}

// const types
func (p *parser) pFlags(flags ast.Flags) Flags         { return Flags(flags) }
func (p *parser) pUseType(useType ast.UseType) UseType { return UseType(useType) }

// special
func (p *parser) handleDeclareStmt(n *ast.DeclareStmt) {
	for _, declare := range n.Declares {
		declareName := declare.Key.Name
		if declareName != "strict_types" {
			p.unsupported("unsupported declare directive: " + declareName)
		}

		valueLit, ok := declare.Value.(*ast.IntLit)
		p.assert(ok && (valueLit.Value == 1 || valueLit.Value == 0), "strict_types declaration must have 0 or 1 as its value")
		p.assert(len(n.Stmts) == 0, "strict_types declaration must not use block mode")
		p.file.StrictTypes = valueLit.Value == 1
	}
}

func (p *parser) pNamespaceStmt(n *ast.NamespaceStmt) *Namespace {
	var name string
	if n.Name != nil {
		name = n.Name.ToString()
	}
	return p.pNamespace(name, n.Stmts)
}

func (p *parser) pNamespace(name string, astStmts []ast.Stmt) *Namespace {
	var initStmts []Stmt
	var segments []Segment

	for _, stmt := range p.pStmtList(astStmts) {
		if declStmt, ok := stmt.(*DeclStmt); ok {
			segments = append(segments, declStmt.Decl)
		} else {
			initStmts = append(initStmts, stmt)
		}
	}

	if len(initStmts) > 0 {
		initFunc := &InitFunc{Stmts: initStmts}
		segments = append([]Segment{initFunc}, segments...)
	}

	return &Namespace{Name: name, Segments: segments}
}

// interface types
func (p *parser) pNode(node ast.Node) Node {
	if node == nil || reflect.ValueOf(node).IsNil() {
		return nil
	}

	switch n := node.(type) {
	case ast.Expr:
		return p.pExpr(n)
	case *ast.Ident:
		return p.pIdent(n)
	case *ast.Name:
		return p.pName(n)
	case *ast.Const:
		p.unsupported("unsupported parseNode(*ast.Const), use parseStmt(*ast.ConstStmt) or parseStmt(*ast.ClassConstStmt) instead.")
	case ast.Stmt:
		p.unsupported("unsupported parseNode(ast.Stmt), use parseStmtList() or parseStmt() instead.")
	case *ast.Attribute, *ast.AttributeGroup:
		p.highVersionFeature("php8.0 attribute")
	case *ast.MatchArm:
		p.highVersionFeature("php8.0 match")
	case *ast.VariadicPlaceholder:
		p.highVersionFeature("php8.2 first class callable syntax")
	default:
		p.fail(fmt.Sprintf("unsupported node type for parseNode(node): %T", n))
	}
	panic("unreachable")
}

func (p *parser) pExpr(node ast.Expr) Expr {
	if node == nil || reflect.ValueOf(node).IsNil() {
		return nil
	}

	switch n := node.(type) {
	case *ast.IntLit:
		return &IntLit{
			Value: n.Value,
		}
	case *ast.FloatLit:
		return &FloatLit{
			Value: n.Value,
		}
	case *ast.StringLit:
		return &StringLit{
			Value: n.Value,
		}
	case *ast.ArrayExpr:
		return &ArrayExpr{
			Items: slices.Map(n.Items, p.pArrayItemExpr),
		}
	case *ast.ArrayItemExpr:
		return &ArrayItemExpr{
			Key:    p.pExpr(n.Key),
			Value:  p.pExpr(n.Value),
			ByRef:  n.ByRef,
			Unpack: n.Unpack,
		}
	case *ast.ClosureExpr:
		return &ClosureExpr{
			Static: n.Static,
			ByRef:  n.ByRef,
			Params: slices.Map(n.Params, p.pParam),
			Uses: slices.Map(n.Uses, func(x *ast.ClosureUseExpr) *ClosureUseExpr {
				return &ClosureUseExpr{
					Name:  p.pVariableIdent(x.Var, "ast.ClosureUseExpr.Var"),
					ByRef: n.ByRef,
				}
			}),
			ReturnType: p.pType(n.ReturnType),
			Stmts:      p.pStmtList(n.Stmts),
		}
	case *ast.ArrowFunctionExpr:
		return &ArrowFunctionExpr{
			Static:     n.Static,
			ByRef:      n.ByRef,
			Params:     slices.Map(n.Params, p.pParam),
			ReturnType: p.pType(n.ReturnType),
			Expr:       p.pExpr(n.Expr),
		}
	case *ast.IndexExpr:
		return &IndexExpr{
			Var: p.pExpr(n.Var),
			Dim: p.pExpr(n.Dim),
		}
	case *ast.CastExpr:
		return &CastExpr{
			Op:   n.Op,
			Expr: p.pExpr(n.Expr),
		}
	case *ast.UnaryExpr:
		return &UnaryExpr{
			Kind: n.Kind,
			Var:  p.pExpr(n.Var),
		}
	case *ast.BinaryExpr:
		return &BinaryExpr{
			Op:    n.Op,
			Left:  p.pExpr(n.Left),
			Right: p.pExpr(n.Right),
		}
	case *ast.AssignExpr:
		return &AssignExpr{
			Op:   n.Op,
			Var:  p.pExpr(n.Var),
			Expr: p.pExpr(n.Expr),
		}
	case *ast.AssignRefExpr:
		return &AssignRefExpr{
			Var:  p.pExpr(n.Var),
			Expr: p.pExpr(n.Expr),
		}
	case *ast.InternalCallExpr:
		return &InternalCallExpr{
			Kind: n.Kind,
			Args: slices.Map(n.Args, p.pExpr),
		}
	case *ast.CloneExpr:
		return &CloneExpr{
			Expr: p.pExpr(n.Expr),
		}
	case *ast.ErrorSuppressExpr:
		return &ErrorSuppressExpr{
			Expr: p.pExpr(n.Expr),
		}
	case *ast.ExitExpr:
		return &ExitExpr{
			Expr: p.pExpr(n.Expr),
		}
	case *ast.ConstFetchExpr:
		return &ConstFetchExpr{
			Name: p.pName(n.Name),
		}
	case *ast.ClassConstFetchExpr:
		return &ClassConstFetchExpr{
			Class: p.pNode(n.Class),
			Name:  p.pIdentString(n.Name),
		}
	case *ast.MagicConstExpr:
		return &MagicConstExpr{
			Kind: n.Kind,
		}
	case *ast.MatchExpr:
		p.highVersionFeature("php8.0 match")
	case *ast.InstanceofExpr:
		return &InstanceofExpr{
			Expr:  p.pExpr(n.Expr),
			Class: p.pNode(n.Class),
		}
	case *ast.ListExpr:
		return &ListExpr{
			Items: slices.Map(n.Items, p.pArrayItemExpr),
		}
	case *ast.PrintExpr:
		return &PrintExpr{
			Expr: p.pExpr(n.Expr),
		}
	case *ast.PropertyFetchExpr:
		return &PropertyFetchExpr{
			Var:  p.pExpr(n.Var),
			Name: p.pNode(n.Name),
		}
	case *ast.StaticPropertyFetchExpr:
		return &StaticPropertyFetchExpr{
			Class: p.pNode(n.Class),
			Name:  p.pNode(n.Name),
		}
	case *ast.ShellExecExpr:
		return &ShellExecExpr{
			Parts: slices.Map(n.Parts, p.pExpr),
		}
	case *ast.TernaryExpr:
		return &TernaryExpr{
			Cond: p.pExpr(n.Cond),
			If:   p.pExpr(n.If),
			Else: p.pExpr(n.Else),
		}
	case *ast.ThrowExpr:
		return &ThrowExpr{
			Expr: p.pExpr(n.Expr),
		}
	case *ast.VariableExpr:
		return &VariableExpr{
			Name: p.pNode(n.Name),
		}
	case *ast.YieldExpr:
		return &YieldExpr{
			Key:   p.pExpr(n.Key),
			Value: p.pExpr(n.Value),
		}
	case *ast.YieldFromExpr:
		return &YieldFromExpr{
			Expr: p.pExpr(n.Expr),
		}
	case *ast.FuncCallExpr:
		return &FuncCallExpr{
			Name: p.pNode(n.Name),
			Args: p.pArgs(n.Args),
		}
	case *ast.NewExpr:
		return &NewExpr{
			Class: p.pNode(n.Class),
			Args:  p.pArgs(n.Args),
		}
	case *ast.MethodCallExpr:
		return &MethodCallExpr{
			Var:  p.pExpr(n.Var),
			Name: p.pNode(n.Name),
			Args: p.pArgs(n.Args),
		}
	case *ast.StaticCallExpr:
		return &StaticCallExpr{
			Class: p.pNode(n.Class),
			Name:  p.pNode(n.Name),
			Args:  p.pArgs(n.Args),
		}
	case *ast.NullsafePropertyFetchExpr:
		p.highVersionFeature("php8.0 nullsafe property fetch")
	case *ast.NullsafeMethodCallExpr:
		p.highVersionFeature("php8.0 nullsafe method call")
	default:
		p.fail(fmt.Sprintf("unsupported expr type for parseExpr(node): %T", n))
	}
	panic("unreachable")
}

func (p *parser) pStmtList(astStmts []ast.Stmt) []Stmt {
	var result []Stmt
	for _, astStmt := range astStmts {
		irStmt := p.pStmt(astStmt)
		if parsingStmts, ok := irStmt.(parsingStmts); ok {
			result = append(result, parsingStmts...)
		} else {
			result = append(result, irStmt)
		}
	}
	return result
}

func (p *parser) pStmt(node ast.Stmt) Stmt {
	switch n := node.(type) {
	case *ast.EmptyStmt:
		return &EmptyStmt{}
	case *ast.BlockStmt:
		return &BlockStmt{
			List: p.pStmtList(n.List),
		}
	case *ast.ExprStmt:
		return &ExprStmt{
			Expr: p.pExpr(n.Expr),
		}
	case *ast.ReturnStmt:
		return &ReturnStmt{
			Expr: p.pExpr(n.Expr),
		}
	case *ast.LabelStmt:
		return &LabelStmt{
			Name: p.pIdentString(n.Name),
		}
	case *ast.GotoStmt:
		return &GotoStmt{
			Name: p.pIdentString(n.Name),
		}
	case *ast.IfStmt:
		return p.pIfStmt(n)
	case *ast.SwitchStmt:
		return p.pSwitchStmt(n)
	case *ast.ForStmt:
		return &ForStmt{
			Init:  slices.Map(n.Init, p.pExpr),
			Cond:  slices.Map(n.Cond, p.pExpr),
			Loop:  slices.Map(n.Loop, p.pExpr),
			Stmts: p.pStmtList(n.Stmts),
		}
	case *ast.ForeachStmt:
		return &ForeachStmt{
			Expr:     p.pExpr(n.Expr),
			KeyVar:   p.pExpr(n.KeyVar),
			ByRef:    n.ByRef,
			ValueVar: p.pExpr(n.ValueVar),
			Stmts:    p.pStmtList(n.Stmts),
		}
	case *ast.BreakStmt:
		return &BreakStmt{
			Num: p.pExpr(n.Num),
		}
	case *ast.ContinueStmt:
		return &ContinueStmt{
			Num: p.pExpr(n.Num),
		}
	case *ast.WhileStmt:
		return &WhileStmt{
			Cond:  p.pExpr(n.Cond),
			Stmts: p.pStmtList(n.Stmts),
		}
	case *ast.DoStmt:
		return &DoStmt{
			Stmts: p.pStmtList(n.Stmts),
			Cond:  p.pExpr(n.Cond),
		}
	case *ast.TryCatchStmt:
		return p.pTryCatchStmt(n)
	case *ast.ConstStmt:
		return parsingStmts(slices.Map(n.Consts, func(c *ast.Const) Stmt {
			return &ConstStmt{
				Name:  p.pNameAsFQ(c.NamespacedName),
				Value: p.pExpr(c.Value),
			}
		}))
	case *ast.EchoStmt:
		return &EchoStmt{
			Exprs: slices.Map(n.Exprs, p.pExpr),
		}
	case *ast.GlobalStmt:
		return &GlobalStmt{
			Vars: slices.Map(n.Vars, p.pExpr),
		}
	case *ast.HaltCompilerStmt:
		return &HaltCompilerStmt{
			Remaining: n.Remaining,
		}
	case *ast.InlineHTMLStmt:
		return &InlineHTMLStmt{
			Value: n.Value,
		}
	case *ast.StaticStmt:
		return parsingStmts(slices.Map(n.Vars, func(x *ast.StaticVarStmt) Stmt {
			return &StaticStmt{
				Name:    p.pVariableIdent(x.Var, "ast.StaticVarStmt.Var"),
				Default: p.pExpr(x.Default),
			}
		}))
	case *ast.UnsetStmt:
		return &UnsetStmt{
			Vars: slices.Map(n.Vars, p.pExpr),
		}
	case *ast.UseStmt:
		return &UseStmt{
			Type:  p.pUseType(n.Type),
			Name:  p.pNameAsFQ(n.Name),
			Alias: nullsafeOrDefault(n.Alias, p.pIdentString, ""),
		}
	case *ast.FunctionStmt:
		p.assert(n.NamespacedName != nil, "FunctionStmt.NamespacedName cannot be nil")

		fn := &Func{
			Name:       p.pNameAsFQ(n.NamespacedName),
			ByRef:      n.ByRef,
			Params:     slices.Map(n.Params, p.pParam),
			ReturnType: p.pType(n.ReturnType),
			Stmts:      p.pStmtList(n.Stmts),
		}
		return &DeclStmt{fn}
	case *ast.InterfaceStmt:
		p.assert(n.NamespacedName != nil, "InterfaceStmt.NamespacedName cannot be nil")

		it := &Interface{
			Name:    p.pNameAsFQ(n.NamespacedName),
			Extends: slices.Map(n.Extends, p.pNameEx),
			Stmts:   p.pStmtList(n.Stmts),
		}
		return &DeclStmt{it}
	case *ast.ClassStmt:
		if n.NamespacedName != nil {
			c := &Class{
				Name:       p.pNameAsFQ(n.NamespacedName),
				Flags:      p.pFlags(n.Flags),
				Extends:    p.pName(n.Extends),
				Implements: slices.Map(n.Implements, p.pNameEx),
				Stmts:      p.pStmtList(n.Stmts),
			}
			return &DeclStmt{c}
		} else {
			// 匿名类
			return &AnonymousClassStmt{
				Flags:      p.pFlags(n.Flags),
				Extends:    p.pName(n.Extends),
				Implements: slices.Map(n.Implements, p.pNameEx),
				Stmts:      p.pStmtList(n.Stmts),
			}
		}
	case *ast.TraitStmt:
		p.assert(n.NamespacedName != nil, "TraitStmt.NamespacedName cannot be nil")

		t := &Trait{
			Name:  p.pNameAsFQ(n.NamespacedName),
			Stmts: p.pStmtList(n.Stmts),
		}
		return &DeclStmt{t}
	case *ast.ClassConstStmt:
		flags := p.pFlags(n.Flags)
		return parsingStmts(slices.Map(n.Consts, func(x *ast.Const) Stmt {
			return &ClassConstStmt{
				Flags: flags,
				Name:  x.Name.Name,
				Value: p.pExpr(x.Value),
			}
		}))
	case *ast.PropertyStmt:
		flags := p.pFlags(n.Flags)
		typ := p.pType(n.Type)

		return parsingStmts(slices.Map(n.Props, func(x *ast.PropertyPropertyStmt) Stmt {
			return &PropertyStmt{
				Flags:   flags,
				Type:    typ,
				Name:    x.Name.Name,
				Default: p.pExpr(x.Default),
			}
		}))
	case *ast.ClassMethodStmt:
		return &MethodStmt{
			Flags:      p.pFlags(n.Flags),
			ByRef:      n.ByRef,
			Name:       p.pIdentString(n.Name),
			Params:     slices.Map(n.Params, p.pParam),
			ReturnType: p.pType(n.ReturnType),
			Stmts:      p.pStmtList(n.Stmts),
		}
	case *ast.TraitUseStmt:
		return &TraitUseStmt{
			Traits:      slices.Map(n.Traits, p.pName),
			Adaptations: slices.Map(n.Adaptations, p.pTraitUseAdaptationStmt),
		}
	case *ast.TraitUseAdaptationAliasStmt:
		return &TraitUseAdaptationAliasStmt{
			NewModifier: p.pFlags(n.NewModifier),
			NewName:     p.pIdent(n.NewName),
			Trait:       p.pName(n.Trait),
			Method:      p.pIdent(n.Method),
		}
	case *ast.TraitUseAdaptationPrecedenceStmt:
		return &TraitUseAdaptationPrecedenceStmt{
			Insteadof: slices.Map(n.Insteadof, p.pName),
			Trait:     p.pName(n.Trait),
			Method:    p.pIdent(n.Method),
		}
	case *ast.EnumStmt, *ast.EnumCaseStmt:
		p.highVersionFeature("php8.1 enum")
	default:
		p.fail(fmt.Sprintf("parseStmt() cannot support this type: %T", n))
	}
	// unreachable
	return nil
}

func (p *parser) pIfStmt(n *ast.IfStmt) *IfStmt {
	return &IfStmt{
		Cond:  p.pExpr(n.Cond),
		Stmts: p.pStmtList(n.Stmts),
		Elseifs: slices.Map(n.Elseifs, func(x *ast.ElseIfStmt) *ElseIfStmt {
			return &ElseIfStmt{
				Cond:  p.pExpr(x.Cond),
				Stmts: p.pStmtList(x.Stmts),
			}
		}),
		Else: nullsafe(n.Else, func(x *ast.ElseStmt) *ElseStmt {
			return &ElseStmt{
				Stmts: p.pStmtList(x.Stmts),
			}
		}),
	}
}

func (p *parser) pSwitchStmt(n *ast.SwitchStmt) *SwitchStmt {
	return &SwitchStmt{
		Cond: p.pExpr(n.Cond),
		Cases: slices.Map(n.Cases, func(x *ast.CaseStmt) *CaseStmt {
			return &CaseStmt{
				Cond:  p.pExpr(x.Cond),
				Stmts: p.pStmtList(x.Stmts),
			}
		}),
	}
}

func (p *parser) pTryCatchStmt(n *ast.TryCatchStmt) *TryCatchStmt {
	return &TryCatchStmt{
		Stmts: p.pStmtList(n.Stmts),
		Catches: slices.Map(n.Catches, func(x *ast.CatchStmt) *CatchStmt {
			if x.Var == nil {
				p.highVersionFeature("php8.0 catch an exception without storing it in a variable.")
			}
			return &CatchStmt{
				Types: slices.Map(x.Types, p.pNameEx),
				Var:   p.pVariableIdent(x.Var, "ast.CatchStmt.Var"),
				Stmts: p.pStmtList(n.Stmts),
			}
		}),
		Finally: nullsafe(n.Finally, func(x *ast.FinallyStmt) *FinallyStmt {
			return &FinallyStmt{
				Stmts: p.pStmtList(x.Stmts),
			}
		}),
	}
}

func (p *parser) pType(node ast.Type) Type {
	switch n := node.(type) {
	case *ast.SimpleType:
		return p.pSimpleType(n)
	case *ast.IntersectionType:
		return &IntersectionType{
			Types: slices.Map(n.Types, p.pType),
		}
	case *ast.UnionType:
		return &UnionType{
			Types: slices.Map(n.Types, p.pType),
		}
	case *ast.NullableType:
		return &NullableType{
			Type: p.pSimpleType(n.Type),
		}
	}
	return nil
}

func (p *parser) pSimpleType(n *ast.SimpleType) *SimpleType {
	if n == nil {
		return nil
	}
	return &SimpleType{
		Name: p.pNameAsFQ(n.Name),
	}
}

func (p *parser) pTraitUseAdaptationStmt(node ast.TraitUseAdaptationStmt) TraitUseAdaptationStmt {
	switch n := node.(type) {
	case *ast.TraitUseAdaptationAliasStmt:
		return &TraitUseAdaptationAliasStmt{
			NewModifier: p.pFlags(n.NewModifier),
			NewName:     p.pIdent(n.NewName),
			Trait:       p.pName(n.Trait),
			Method:      p.pIdent(n.Method),
		}
	case *ast.TraitUseAdaptationPrecedenceStmt:
		return &TraitUseAdaptationPrecedenceStmt{
			Insteadof: slices.Map(n.Insteadof, p.pName),
			Trait:     p.pName(n.Trait),
			Method:    p.pIdent(n.Method),
		}
	}
	return nil
}

// struct types
func (p *parser) pArg(n *ast.Arg) *Arg {
	if n.Name != nil {
		p.highVersionFeature("php8.0 named arguments")
	}
	if n.ByRef {
		p.lowerVersionFeature("Call-time pass-by-reference has been removed in PHP 5.4")
	}
	return &Arg{
		Value:  p.pExpr(n.Value),
		Unpack: n.Unpack,
	}
}
func (p *parser) pArgs(args []ast.Node) []*Arg {
	return slices.Map(args, func(n ast.Node) *Arg {
		switch arg := n.(type) {
		case *ast.Arg:
			return p.pArg(arg)
		case *ast.VariadicPlaceholder:
			p.highVersionFeature("php8.2 first class callable syntax")
		default:
			p.fail(fmt.Sprintf("expected type of arg must be *ast.Arg, provide is %T", arg))
		}
		panic("unreachable")
	})
}

func (p *parser) pIdentString(n *ast.Ident) string {
	p.assert(n != nil, "*ast.Ident cannot be nil")
	return n.Name
}
func (p *parser) pIdent(n *ast.Ident) *Ident {
	if n == nil {
		return nil
	}
	ident := Ident(n.Name)
	return &ident
}
func (p *parser) pParam(n *ast.Param) *Param {
	if n == nil {
		return nil
	}
	if n.Flags != 0 {
		p.highVersionFeature("php8.0 constructor promotion")
	}

	return &Param{
		Name:     p.pVariableIdent(n.Var, "ast.Param.Var"),
		Type:     p.pType(n.Type),
		ByRef:    n.ByRef,
		Variadic: n.Variadic,
		Default:  p.pExpr(n.Default),
	}
}

func (p *parser) pNameAsFQ(n *ast.Name) Name {
	return MakeName(NameFullyQualified, n.Parts)
}

func (p *parser) pNameEx(n *ast.Name) Name {
	return *p.pName(n)
}

func (p *parser) pName(n *ast.Name) *Name {
	if n == nil {
		return nil
	}
	if n.Kind != ast.NameFullyQualified {
		log.Println("ast.Name.Kind is not FQ")
	}

	switch n.Kind {
	case ast.NameNormal:
		return NewName(NameNormal, n.Parts)
	case ast.NameFullyQualified:
		return NewName(NameFullyQualified, n.Parts)
	case ast.NameRelative:
		return NewName(NameRelative, n.Parts)
	default:
		p.fail(fmt.Sprintf("unexpected ast.Name.Kind: %d", n.Kind))
	}
	panic("unreachable")
}
func (p *parser) pArrayItemExpr(n *ast.ArrayItemExpr) *ArrayItemExpr {
	if n == nil {
		return nil
	}
	return &ArrayItemExpr{
		Key:    p.pExpr(n.Key),
		Value:  p.pExpr(n.Value),
		ByRef:  n.ByRef,
		Unpack: n.Unpack,
	}
}
func (p *parser) pVariableIdent(n *ast.VariableExpr, typ string) string {
	p.assert(n != nil, typ+" must not nil")

	nameIdent := n.Name.(*ast.Ident)
	p.assert(nameIdent != nil, typ+".Name must be a Ident")

	return nameIdent.Name
}
func (p *parser) pVariableExpr(n *ast.VariableExpr) *VariableExpr {
	if n == nil {
		return nil
	}
	return &VariableExpr{
		Name: p.pNode(n.Name),
	}
}
