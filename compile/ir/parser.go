package ir

import (
	"errors"
	"fmt"
	"github.com/heyuuu/gophp/php/ast"
	"github.com/heyuuu/gophp/utils/slices"
	"log"
)

func ParseAstFile(astFile []ast.Stmt) (file *File, err error) {
	defer func() {
		switch e := recover().(type) {
		case nil:
			return
		case string:
			err = errors.New(e)
		case error:
			err = e
		default:
			panic(e)
		}
	}()

	p := &parser{}
	return p.ParseFile(astFile), nil
}

// parsingStmts
type parsingStmts []Stmt

func (p parsingStmts) node()     {}
func (p parsingStmts) stmtNode() {}

// parser
type parser struct{}

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

	//
	f := &File{}

	slices.Each(declareStmts, p.parseDeclareStmt)

	if len(globalStmts) > 0 {
		f.Segments = []Segment{p.buildSegment("", p.parseStmtList(globalStmts))}
	} else {
		f.Segments = slices.Map(namespaceStmts, p.parseNamespaceStmt)
	}

	return f
}

// misc
func (p *parser) assert(cond bool, message string) {
	if !cond {
		log.Println(message)
		panic(message)
	}
}
func (p *parser) highVersionFeature(feature string) {
	p.assert(false, "high version php feature: "+feature)
}
func (p *parser) lowerVersionFeature(feature string) {
	p.assert(false, "lower version php feature: "+feature)
}
func (p *parser) unsupported(message string) {
	p.assert(false, message)
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
func (p *parser) parseDeclareStmt(n *ast.DeclareStmt) {
	// todo declare
}

func (p *parser) parseNamespaceStmt(n *ast.NamespaceStmt) Segment {
	var namespace string
	if n.Name != nil {
		namespace = n.Name.ToString()
	}
	stmts := p.parseStmtList(n.Stmts)
	return p.buildSegment(namespace, stmts)
}

// interface types
func (p *parser) parseNode(node ast.Node) Node {
	switch n := node.(type) {
	case ast.Expr:
		return p.parseExpr(n)
	case *ast.Ident:
		return p.parseIdent(n)
	case *ast.Name:
		return p.parseName(n)
	case *ast.Arg:
		return p.parseArg(n)
	case *ast.Param:
		return p.parseParam(n)
	case *ast.Const:
		return p.parseConst(n)
	case *ast.VariadicPlaceholder:
		return &VariadicPlaceholder{}
	case ast.Stmt:
		p.unsupported("unsupported parseNode(ast.Stmt), use parseStmtList() or parseStmt() instead.")
	case *ast.Attribute, *ast.AttributeGroup:
		p.highVersionFeature("php8.0 attribute")
	case *ast.MatchArm:
		p.highVersionFeature("php8.0 match")
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
			Stmts:      p.parseStmtList(n.Stmts),
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
		p.highVersionFeature("php8.0 match")
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
	case *ast.StaticCallExpr:
		return &StaticCallExpr{
			Class: p.parseNode(n.Class),
			Name:  p.parseNode(n.Name),
			Args:  slices.Map(n.Args, p.parseNode),
		}
	case *ast.NullsafePropertyFetchExpr:
		p.highVersionFeature("php8.0 nullsafe property fetch")
	case *ast.NullsafeMethodCallExpr:
		p.highVersionFeature("php8.0 nullsafe method call")
	}
	return nil
}

func (p *parser) parseStmtList(astStmts []ast.Stmt) []Stmt {
	var result []Stmt
	for _, astStmt := range astStmts {
		irStmt := p.parseStmt(astStmt)
		if parsingStmts, ok := irStmt.(parsingStmts); ok {
			result = append(result, parsingStmts...)
		} else {
			result = append(result, irStmt)
		}
	}
	return result
}

func (p *parser) parseStmt(node ast.Stmt) Stmt {
	switch n := node.(type) {
	case *ast.EmptyStmt:
		return &EmptyStmt{}
	case *ast.BlockStmt:
		return &BlockStmt{
			List: p.parseStmtList(n.List),
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
			Stmts:   p.parseStmtList(n.Stmts),
			Elseifs: slices.Map(n.Elseifs, p.parseElseIfStmt),
			Else:    p.parseElseStmt(n.Else),
		}
	case *ast.ElseIfStmt:
		return &ElseIfStmt{
			Cond:  p.parseExpr(n.Cond),
			Stmts: p.parseStmtList(n.Stmts),
		}
	case *ast.ElseStmt:
		return &ElseStmt{
			Stmts: p.parseStmtList(n.Stmts),
		}
	case *ast.SwitchStmt:
		return &SwitchStmt{
			Cond:  p.parseExpr(n.Cond),
			Cases: slices.Map(n.Cases, p.parseCaseStmt),
		}
	case *ast.CaseStmt:
		return &CaseStmt{
			Cond:  p.parseExpr(n.Cond),
			Stmts: p.parseStmtList(n.Stmts),
		}
	case *ast.ForStmt:
		return &ForStmt{
			Init:  slices.Map(n.Init, p.parseExpr),
			Cond:  slices.Map(n.Cond, p.parseExpr),
			Loop:  slices.Map(n.Loop, p.parseExpr),
			Stmts: p.parseStmtList(n.Stmts),
		}
	case *ast.ForeachStmt:
		return &ForeachStmt{
			Expr:     p.parseExpr(n.Expr),
			KeyVar:   p.parseExpr(n.KeyVar),
			ByRef:    n.ByRef,
			ValueVar: p.parseExpr(n.ValueVar),
			Stmts:    p.parseStmtList(n.Stmts),
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
			Stmts: p.parseStmtList(n.Stmts),
		}
	case *ast.DoStmt:
		return &DoStmt{
			Stmts: p.parseStmtList(n.Stmts),
			Cond:  p.parseExpr(n.Cond),
		}
	case *ast.TryCatchStmt:
		return &TryCatchStmt{
			Stmts:   p.parseStmtList(n.Stmts),
			Catches: slices.Map(n.Catches, p.parseCatchStmt),
			Finally: p.parseFinallyStmt(n.Finally),
		}
	case *ast.CatchStmt:
		return &CatchStmt{
			Types: slices.Map(n.Types, p.parseName),
			Var:   p.parseVariableExpr(n.Var),
			Stmts: p.parseStmtList(n.Stmts),
		}
	case *ast.FinallyStmt:
		return &FinallyStmt{
			Stmts: p.parseStmtList(n.Stmts),
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
			Stmts:    p.parseStmtList(n.Stmts),
		}
	case *ast.DeclareDeclareStmt:
		return &DeclareDeclareStmt{
			Key:   p.parseIdent(n.Key),
			Value: p.parseExpr(n.Value),
		}
	case *ast.NamespaceStmt:
		return &NamespaceStmt{
			Name:  p.parseName(n.Name),
			Stmts: p.parseStmtList(n.Stmts),
		}
	case *ast.FunctionStmt:
		p.assert(n.NamespacedName != nil, "FunctionStmt.NamespacedName cannot be nil")

		return &FunctionStmt{
			Name:       p.parseNameAsFQ(n.NamespacedName),
			ByRef:      n.ByRef,
			Params:     slices.Map(n.Params, p.parseParam),
			ReturnType: p.parseType(n.ReturnType),
			Stmts:      p.parseStmtList(n.Stmts),
		}
	case *ast.InterfaceStmt:
		p.assert(n.NamespacedName != nil, "InterfaceStmt.NamespacedName cannot be nil")

		return &InterfaceStmt{
			Name:    p.parseNameAsFQ(n.NamespacedName),
			Extends: slices.Map(n.Extends, p.parseName),
			Stmts:   p.parseStmtList(n.Stmts),
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
			Stmts:      p.parseStmtList(n.Stmts),
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
			Stmts:      p.parseStmtList(n.Stmts),
		}
	case *ast.TraitStmt:
		p.assert(n.NamespacedName != nil, "TraitStmt.NamespacedName cannot be nil")

		return &TraitStmt{
			Name:  p.parseNameAsFQ(n.NamespacedName),
			Stmts: p.parseStmtList(n.Stmts),
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
		p.highVersionFeature("php8.1 enum")
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
	if n.Name != nil {
		p.highVersionFeature("php8.0 named arguments")
	}
	if n.ByRef {
		p.lowerVersionFeature("Call-time pass-by-reference has been removed in PHP 5.4")
	}
	return &Arg{
		Value:  p.parseExpr(n.Value),
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
	if n.Flags != 0 {
		p.highVersionFeature("php8.0 constructor promotion")
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
		Stmts: p.parseStmtList(n.Stmts),
	}
}
func (p *parser) parseElseStmt(n *ast.ElseStmt) *ElseStmt {
	if n == nil {
		return nil
	}
	return &ElseStmt{
		Stmts: p.parseStmtList(n.Stmts),
	}
}
func (p *parser) parseCaseStmt(n *ast.CaseStmt) *CaseStmt {
	if n == nil {
		return nil
	}
	return &CaseStmt{
		Cond:  p.parseExpr(n.Cond),
		Stmts: p.parseStmtList(n.Stmts),
	}
}
func (p *parser) parseCatchStmt(n *ast.CatchStmt) *CatchStmt {
	if n == nil {
		return nil
	}

	return &CatchStmt{
		Types: slices.Map(n.Types, p.parseName),
		Var:   p.parseVariableExpr(n.Var),
		Stmts: p.parseStmtList(n.Stmts),
	}
}
func (p *parser) parseFinallyStmt(n *ast.FinallyStmt) *FinallyStmt {
	if n == nil {
		return nil
	}
	return &FinallyStmt{
		Stmts: p.parseStmtList(n.Stmts),
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
