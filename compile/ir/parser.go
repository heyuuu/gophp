package ir

import (
	"fmt"
	"github.com/heyuuu/gophp/php/ast"
	"github.com/heyuuu/gophp/utils/slices"
	"log"
)

func ParseAstFile(astFile []ast.Stmt) *File {
	p := &parser{}
	return p.ParseFile(astFile)
}

// parser
type parser struct{}

func (p *parser) ParseFile(astFile []ast.Stmt) *File {
	var defaultStmts []Stmt

	f := &File{}
	for _, astStmt := range astFile {
		switch s := astStmt.(type) {
		case *ast.NamespaceStmt:
			f.Segments = append(f.Segments, p.parseNamespaceStmt(s))
		default:
			defaultStmts = append(defaultStmts, p.parseStmt(s))
		}
	}

	if len(defaultStmts) > 0 {
		f.Segments = append(
			[]Segment{p.buildSegment("", defaultStmts)},
			f.Segments...,
		)
	}

	return f
}

// misc
func (p *parser) assert(cond bool, message string) {
	if !cond {
		log.Fatal(message)
	}
}
func (p *parser) unsupportedFeature(feature string) {
	p.assert(false, "high version php feature: "+feature)
}

func (p *parser) buildSegment(namespace string, stmts []Stmt) Segment {
	var inits []Stmt
	var decls []Stmt
	for _, irStmt := range stmts {
		switch irStmt.(type) {
		case *FunctionStmt,
			*ClassStmt,
			*InterfaceStmt,
			*TraitStmt:
			decls = append(decls, irStmt)
		default:
			inits = append(inits, irStmt)
		}
	}

	var initStmt *InitStmt
	if len(inits) > 0 {
		initStmt = &InitStmt{Stmts: inits}
	}
	return Segment{Namespace: namespace, Init: initStmt, Decls: decls}
}

// const types
func (p *parser) parseFlags(flags ast.Flags) Flags         { return Flags(flags) }
func (p *parser) parseUseType(useType ast.UseType) UseType { return UseType(useType) }

// special
func (p *parser) parseNamespaceStmt(n *ast.NamespaceStmt) Segment {
	var namespace string
	if n.Name != nil {
		namespace = n.Name.ToString()
	}
	stmts := slices.Map(n.Stmts, p.parseStmt)
	return p.buildSegment(namespace, stmts)
}

// interface types
func (p *parser) parseNode(node ast.Node) Node {
	switch n := node.(type) {
	case ast.Expr:
		return p.parseExpr(n)
	case ast.Stmt:
		return p.parseStmt(n)
	case *ast.Ident:
		return p.parseIdent(n)
	case *ast.Name:
		return p.parseName(n)
	case *ast.Arg:
		return p.parseArg(n)
	case *ast.Param:
		return p.parseParam(n)
	case *ast.Attribute, *ast.AttributeGroup:
		p.unsupportedFeature("php8.0 attribute")
	case *ast.Const:
		return p.parseConst(n)
	case *ast.MatchArm:
		p.unsupportedFeature("php8.0 match")
	case *ast.VariadicPlaceholder:
		return &VariadicPlaceholder{}
	}
	return nil
}

func (p *parser) parseExpr(node ast.Expr) Expr {
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
			Items: slices.Map(n.Items, p.parseArrayItemExpr),
		}
	case *ast.ArrayItemExpr:
		return &ArrayItemExpr{
			Key:    p.parseExpr(n.Key),
			Value:  p.parseExpr(n.Value),
			ByRef:  n.ByRef,
			Unpack: n.Unpack,
		}
	case *ast.ClosureExpr:
		return &ClosureExpr{
			Static:     n.Static,
			ByRef:      n.ByRef,
			Params:     slices.Map(n.Params, p.parseParam),
			Uses:       slices.Map(n.Uses, p.parseClosureUseExpr),
			ReturnType: p.parseType(n.ReturnType),
			Stmts:      slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.ClosureUseExpr:
		return &ClosureUseExpr{
			Var:   p.parseVariableExpr(n.Var),
			ByRef: n.ByRef,
		}
	case *ast.ArrowFunctionExpr:
		return &ArrowFunctionExpr{
			Static:     n.Static,
			ByRef:      n.ByRef,
			Params:     slices.Map(n.Params, p.parseParam),
			ReturnType: p.parseType(n.ReturnType),
			Expr:       p.parseExpr(n.Expr),
		}
	case *ast.IndexExpr:
		return &IndexExpr{
			Var: p.parseExpr(n.Var),
			Dim: p.parseExpr(n.Dim),
		}
	case *ast.CastExpr:
		return &CastExpr{
			Op:   n.Op,
			Expr: p.parseExpr(n.Expr),
		}
	case *ast.UnaryExpr:
		return &UnaryExpr{
			Kind: n.Kind,
			Var:  p.parseExpr(n.Var),
		}
	case *ast.BinaryExpr:
		return &BinaryExpr{
			Op:    n.Op,
			Left:  p.parseExpr(n.Left),
			Right: p.parseExpr(n.Right),
		}
	case *ast.AssignExpr:
		return &AssignExpr{
			Op:   n.Op,
			Var:  p.parseExpr(n.Var),
			Expr: p.parseExpr(n.Expr),
		}
	case *ast.AssignRefExpr:
		return &AssignRefExpr{
			Var:  p.parseExpr(n.Var),
			Expr: p.parseExpr(n.Expr),
		}
	case *ast.InternalCallExpr:
		return &InternalCallExpr{
			Kind: n.Kind,
			Args: slices.Map(n.Args, p.parseExpr),
		}
	case *ast.CloneExpr:
		return &CloneExpr{
			Expr: p.parseExpr(n.Expr),
		}
	case *ast.ErrorSuppressExpr:
		return &ErrorSuppressExpr{
			Expr: p.parseExpr(n.Expr),
		}
	case *ast.ExitExpr:
		return &ExitExpr{
			Expr: p.parseExpr(n.Expr),
		}
	case *ast.ConstFetchExpr:
		return &ConstFetchExpr{
			Name: p.parseName(n.Name),
		}
	case *ast.ClassConstFetchExpr:
		return &ClassConstFetchExpr{
			Class: p.parseNode(n.Class),
			Name:  p.parseIdent(n.Name),
		}
	case *ast.MagicConstExpr:
		return &MagicConstExpr{
			Kind: n.Kind,
		}
	case *ast.MatchExpr:
		p.unsupportedFeature("php8.0 match")
	case *ast.InstanceofExpr:
		return &InstanceofExpr{
			Expr:  p.parseExpr(n.Expr),
			Class: p.parseNode(n.Class),
		}
	case *ast.ListExpr:
		return &ListExpr{
			Items: slices.Map(n.Items, p.parseArrayItemExpr),
		}
	case *ast.PrintExpr:
		return &PrintExpr{
			Expr: p.parseExpr(n.Expr),
		}
	case *ast.PropertyFetchExpr:
		return &PropertyFetchExpr{
			Var:  p.parseExpr(n.Var),
			Name: p.parseNode(n.Name),
		}
	case *ast.NullsafePropertyFetchExpr:
		return &NullsafePropertyFetchExpr{
			Var:  p.parseExpr(n.Var),
			Name: p.parseNode(n.Name),
		}
	case *ast.StaticPropertyFetchExpr:
		return &StaticPropertyFetchExpr{
			Class: p.parseNode(n.Class),
			Name:  p.parseNode(n.Name),
		}
	case *ast.ShellExecExpr:
		return &ShellExecExpr{
			Parts: slices.Map(n.Parts, p.parseExpr),
		}
	case *ast.TernaryExpr:
		return &TernaryExpr{
			Cond: p.parseExpr(n.Cond),
			If:   p.parseExpr(n.If),
			Else: p.parseExpr(n.Else),
		}
	case *ast.ThrowExpr:
		return &ThrowExpr{
			Expr: p.parseExpr(n.Expr),
		}
	case *ast.VariableExpr:
		return &VariableExpr{
			Name: p.parseNode(n.Name),
		}
	case *ast.YieldExpr:
		return &YieldExpr{
			Key:   p.parseExpr(n.Key),
			Value: p.parseExpr(n.Value),
		}
	case *ast.YieldFromExpr:
		return &YieldFromExpr{
			Expr: p.parseExpr(n.Expr),
		}
	case *ast.FuncCallExpr:
		return &FuncCallExpr{
			Name: p.parseNode(n.Name),
			Args: slices.Map(n.Args, p.parseNode),
		}
	case *ast.NewExpr:
		return &NewExpr{
			Class: p.parseNode(n.Class),
			Args:  slices.Map(n.Args, p.parseNode),
		}
	case *ast.MethodCallExpr:
		return &MethodCallExpr{
			Var:  p.parseExpr(n.Var),
			Name: p.parseNode(n.Name),
			Args: slices.Map(n.Args, p.parseNode),
		}
	case *ast.NullsafeMethodCallExpr:
		return &NullsafeMethodCallExpr{
			Var:  p.parseExpr(n.Var),
			Name: p.parseNode(n.Name),
			Args: slices.Map(n.Args, p.parseNode),
		}
	case *ast.StaticCallExpr:
		return &StaticCallExpr{
			Class: p.parseNode(n.Class),
			Name:  p.parseNode(n.Name),
			Args:  slices.Map(n.Args, p.parseNode),
		}
	}
	return nil
}

func (p *parser) parseStmt(node ast.Stmt) Stmt {
	switch n := node.(type) {
	case *ast.EmptyStmt:
		return &EmptyStmt{}
	case *ast.BlockStmt:
		return &BlockStmt{
			List: slices.Map(n.List, p.parseStmt),
		}
	case *ast.ExprStmt:
		return &ExprStmt{
			Expr: p.parseExpr(n.Expr),
		}
	case *ast.ReturnStmt:
		return &ReturnStmt{
			Expr: p.parseExpr(n.Expr),
		}
	case *ast.LabelStmt:
		return &LabelStmt{
			Name: p.parseIdent(n.Name),
		}
	case *ast.GotoStmt:
		return &GotoStmt{
			Name: p.parseIdent(n.Name),
		}
	case *ast.IfStmt:
		return &IfStmt{
			Cond:    p.parseExpr(n.Cond),
			Stmts:   slices.Map(n.Stmts, p.parseStmt),
			Elseifs: slices.Map(n.Elseifs, p.parseElseIfStmt),
			Else:    p.parseElseStmt(n.Else),
		}
	case *ast.ElseIfStmt:
		return &ElseIfStmt{
			Cond:  p.parseExpr(n.Cond),
			Stmts: slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.ElseStmt:
		return &ElseStmt{
			Stmts: slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.SwitchStmt:
		return &SwitchStmt{
			Cond:  p.parseExpr(n.Cond),
			Cases: slices.Map(n.Cases, p.parseCaseStmt),
		}
	case *ast.CaseStmt:
		return &CaseStmt{
			Cond:  p.parseExpr(n.Cond),
			Stmts: slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.ForStmt:
		return &ForStmt{
			Init:  slices.Map(n.Init, p.parseExpr),
			Cond:  slices.Map(n.Cond, p.parseExpr),
			Loop:  slices.Map(n.Loop, p.parseExpr),
			Stmts: slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.ForeachStmt:
		return &ForeachStmt{
			Expr:     p.parseExpr(n.Expr),
			KeyVar:   p.parseExpr(n.KeyVar),
			ByRef:    n.ByRef,
			ValueVar: p.parseExpr(n.ValueVar),
			Stmts:    slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.BreakStmt:
		return &BreakStmt{
			Num: p.parseExpr(n.Num),
		}
	case *ast.ContinueStmt:
		return &ContinueStmt{
			Num: p.parseExpr(n.Num),
		}
	case *ast.WhileStmt:
		return &WhileStmt{
			Cond:  p.parseExpr(n.Cond),
			Stmts: slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.DoStmt:
		return &DoStmt{
			Stmts: slices.Map(n.Stmts, p.parseStmt),
			Cond:  p.parseExpr(n.Cond),
		}
	case *ast.TryCatchStmt:
		return &TryCatchStmt{
			Stmts:   slices.Map(n.Stmts, p.parseStmt),
			Catches: slices.Map(n.Catches, p.parseCatchStmt),
			Finally: p.parseFinallyStmt(n.Finally),
		}
	case *ast.CatchStmt:
		return &CatchStmt{
			Types: slices.Map(n.Types, p.parseName),
			Var:   p.parseVariableExpr(n.Var),
			Stmts: slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.FinallyStmt:
		return &FinallyStmt{
			Stmts: slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.ConstStmt:
		return &ConstStmt{
			Consts: slices.Map(n.Consts, p.parseConst),
		}
	case *ast.EchoStmt:
		return &EchoStmt{
			Exprs: slices.Map(n.Exprs, p.parseExpr),
		}
	case *ast.GlobalStmt:
		return &GlobalStmt{
			Vars: slices.Map(n.Vars, p.parseExpr),
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
		return &StaticStmt{
			Vars: slices.Map(n.Vars, p.parseStaticVarStmt),
		}
	case *ast.StaticVarStmt:
		return &StaticVarStmt{
			Var:     p.parseVariableExpr(n.Var),
			Default: p.parseExpr(n.Default),
		}
	case *ast.UnsetStmt:
		return &UnsetStmt{
			Vars: slices.Map(n.Vars, p.parseExpr),
		}
	case *ast.UseStmt:
		return &UseStmt{
			Type:  p.parseUseType(n.Type),
			Name:  p.parseName(n.Name),
			Alias: p.parseIdent(n.Alias),
		}
	case *ast.DeclareStmt:
		return &DeclareStmt{
			Declares: slices.Map(n.Declares, p.parseDeclareDeclareStmt),
			Stmts:    slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.DeclareDeclareStmt:
		return &DeclareDeclareStmt{
			Key:   p.parseIdent(n.Key),
			Value: p.parseExpr(n.Value),
		}
	case *ast.NamespaceStmt:
		return &NamespaceStmt{
			Name:  p.parseName(n.Name),
			Stmts: slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.FunctionStmt:
		p.assert(n.NamespacedName != nil, "FunctionStmt.NamespacedName cannot be nil")

		return &FunctionStmt{
			Name:       p.parseNameAsFQ(n.NamespacedName),
			ByRef:      n.ByRef,
			Params:     slices.Map(n.Params, p.parseParam),
			ReturnType: p.parseType(n.ReturnType),
			Stmts:      slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.InterfaceStmt:
		p.assert(n.NamespacedName != nil, "InterfaceStmt.NamespacedName cannot be nil")

		return &InterfaceStmt{
			Name:    p.parseNameAsFQ(n.NamespacedName),
			Extends: slices.Map(n.Extends, p.parseName),
			Stmts:   slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.ClassStmt:
		// todo 将匿名类和实名类定义区分开
		var name *Name
		if n.NamespacedName != nil {
			name = p.parseNameAsFQ(n.NamespacedName)
		}

		return &ClassStmt{
			Name:       name,
			Flags:      p.parseFlags(n.Flags),
			Extends:    p.parseName(n.Extends),
			Implements: slices.Map(n.Implements, p.parseName),
			Stmts:      slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.ClassConstStmt:
		return &ClassConstStmt{
			Flags:  p.parseFlags(n.Flags),
			Consts: slices.Map(n.Consts, p.parseConst),
		}
	case *ast.PropertyStmt:
		return &PropertyStmt{
			Flags: p.parseFlags(n.Flags),
			Props: slices.Map(n.Props, p.parsePropertyPropertyStmt),
			Type:  p.parseType(n.Type),
		}
	case *ast.PropertyPropertyStmt:
		return &PropertyPropertyStmt{
			Name:    p.parseIdent(n.Name),
			Default: p.parseExpr(n.Default),
		}
	case *ast.ClassMethodStmt:
		return &ClassMethodStmt{
			Flags:      p.parseFlags(n.Flags),
			ByRef:      n.ByRef,
			Name:       p.parseIdent(n.Name),
			Params:     slices.Map(n.Params, p.parseParam),
			ReturnType: p.parseType(n.ReturnType),
			Stmts:      slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.TraitStmt:
		p.assert(n.NamespacedName != nil, "TraitStmt.NamespacedName cannot be nil")

		return &TraitStmt{
			Name:  p.parseNameAsFQ(n.NamespacedName),
			Stmts: slices.Map(n.Stmts, p.parseStmt),
		}
	case *ast.TraitUseStmt:
		return &TraitUseStmt{
			Traits:      slices.Map(n.Traits, p.parseName),
			Adaptations: slices.Map(n.Adaptations, p.parseTraitUseAdaptationStmt),
		}
	case *ast.TraitUseAdaptationAliasStmt:
		return &TraitUseAdaptationAliasStmt{
			NewModifier: p.parseFlags(n.NewModifier),
			NewName:     p.parseIdent(n.NewName),
			Trait:       p.parseName(n.Trait),
			Method:      p.parseIdent(n.Method),
		}
	case *ast.TraitUseAdaptationPrecedenceStmt:
		return &TraitUseAdaptationPrecedenceStmt{
			Insteadof: slices.Map(n.Insteadof, p.parseName),
			Trait:     p.parseName(n.Trait),
			Method:    p.parseIdent(n.Method),
		}
	case *ast.EnumStmt, *ast.EnumCaseStmt:
		p.unsupportedFeature("php8.1 enum")
	}
	return nil
}

func (p *parser) parseType(node ast.Type) Type {
	switch n := node.(type) {
	case *ast.SimpleType:
		return &SimpleType{
			Name: p.parseName(n.Name),
		}
	case *ast.IntersectionType:
		return &IntersectionType{
			Types: slices.Map(n.Types, p.parseType),
		}
	case *ast.UnionType:
		return &UnionType{
			Types: slices.Map(n.Types, p.parseType),
		}
	case *ast.NullableType:
		return &NullableType{
			Type: p.parseSimpleType(n.Type),
		}
	}
	return nil
}

func (p *parser) parseTraitUseAdaptationStmt(node ast.TraitUseAdaptationStmt) TraitUseAdaptationStmt {
	switch n := node.(type) {
	case *ast.TraitUseAdaptationAliasStmt:
		return &TraitUseAdaptationAliasStmt{
			NewModifier: p.parseFlags(n.NewModifier),
			NewName:     p.parseIdent(n.NewName),
			Trait:       p.parseName(n.Trait),
			Method:      p.parseIdent(n.Method),
		}
	case *ast.TraitUseAdaptationPrecedenceStmt:
		return &TraitUseAdaptationPrecedenceStmt{
			Insteadof: slices.Map(n.Insteadof, p.parseName),
			Trait:     p.parseName(n.Trait),
			Method:    p.parseIdent(n.Method),
		}
	}
	return nil
}

// struct types
func (p *parser) parseArg(n *ast.Arg) *Arg {
	return &Arg{
		Name:   p.parseIdent(n.Name),
		Value:  p.parseExpr(n.Value),
		ByRef:  n.ByRef,
		Unpack: n.Unpack,
	}
}
func (p *parser) parseConst(n *ast.Const) *Const {
	return &Const{
		Name:  p.parseNameAsFQ(n.NamespacedName),
		Value: p.parseExpr(n.Value),
	}
}
func (p *parser) parseIdent(n *ast.Ident) *Ident {
	if n == nil {
		return nil
	}
	return &Ident{
		Name:    n.Name,
		VarLike: n.VarLike,
	}
}
func (p *parser) parseParam(n *ast.Param) *Param {
	if n == nil {
		return nil
	}
	return &Param{
		Type:     p.parseType(n.Type),
		ByRef:    n.ByRef,
		Variadic: n.Variadic,
		Var:      p.parseVariableExpr(n.Var),
		Default:  p.parseExpr(n.Default),
		Flags:    p.parseFlags(n.Flags),
	}
}
func (p *parser) parseSimpleType(n *ast.SimpleType) *SimpleType {
	if n == nil {
		return nil
	}
	return &SimpleType{
		Name: p.parseName(n.Name),
	}
}

func (p *parser) parseNameAsFQ(n *ast.Name) *Name {
	return NewName(NameFullyQualified, n.Parts)
}

func (p *parser) parseName(n *ast.Name) *Name {
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
		panic(fmt.Sprintf("unexpected ast.Name.Kind: %d", n.Kind))
	}
}
func (p *parser) parseArrayItemExpr(n *ast.ArrayItemExpr) *ArrayItemExpr {
	if n == nil {
		return nil
	}
	return &ArrayItemExpr{
		Key:    p.parseExpr(n.Key),
		Value:  p.parseExpr(n.Value),
		ByRef:  n.ByRef,
		Unpack: n.Unpack,
	}
}
func (p *parser) parseClosureUseExpr(n *ast.ClosureUseExpr) *ClosureUseExpr {
	if n == nil {
		return nil
	}
	return &ClosureUseExpr{
		Var:   p.parseVariableExpr(n.Var),
		ByRef: n.ByRef,
	}
}
func (p *parser) parseVariableExpr(n *ast.VariableExpr) *VariableExpr {
	if n == nil {
		return nil
	}
	return &VariableExpr{
		Name: p.parseNode(n.Name),
	}
}
func (p *parser) parseElseIfStmt(n *ast.ElseIfStmt) *ElseIfStmt {
	if n == nil {
		return nil
	}
	return &ElseIfStmt{
		Cond:  p.parseExpr(n.Cond),
		Stmts: slices.Map(n.Stmts, p.parseStmt),
	}
}
func (p *parser) parseElseStmt(n *ast.ElseStmt) *ElseStmt {
	if n == nil {
		return nil
	}
	return &ElseStmt{
		Stmts: slices.Map(n.Stmts, p.parseStmt),
	}
}
func (p *parser) parseCaseStmt(n *ast.CaseStmt) *CaseStmt {
	if n == nil {
		return nil
	}
	return &CaseStmt{
		Cond:  p.parseExpr(n.Cond),
		Stmts: slices.Map(n.Stmts, p.parseStmt),
	}
}
func (p *parser) parseCatchStmt(n *ast.CatchStmt) *CatchStmt {
	if n == nil {
		return nil
	}

	return &CatchStmt{
		Types: slices.Map(n.Types, p.parseName),
		Var:   p.parseVariableExpr(n.Var),
		Stmts: slices.Map(n.Stmts, p.parseStmt),
	}
}
func (p *parser) parseFinallyStmt(n *ast.FinallyStmt) *FinallyStmt {
	if n == nil {
		return nil
	}
	return &FinallyStmt{
		Stmts: slices.Map(n.Stmts, p.parseStmt),
	}
}
func (p *parser) parseStaticVarStmt(n *ast.StaticVarStmt) *StaticVarStmt {
	if n == nil {
		return nil
	}
	return &StaticVarStmt{
		Var:     p.parseVariableExpr(n.Var),
		Default: p.parseExpr(n.Default),
	}
}
func (p *parser) parseDeclareDeclareStmt(n *ast.DeclareDeclareStmt) *DeclareDeclareStmt {
	if n == nil {
		return nil
	}
	return &DeclareDeclareStmt{
		Key:   p.parseIdent(n.Key),
		Value: p.parseExpr(n.Value),
	}
}
func (p *parser) parsePropertyPropertyStmt(n *ast.PropertyPropertyStmt) *PropertyPropertyStmt {
	if n == nil {
		return nil
	}
	return &PropertyPropertyStmt{
		Name:    p.parseIdent(n.Name),
		Default: p.parseExpr(n.Default),
	}
}
