package ir

import (
	"fmt"
	"github.com/heyuuu/gophp/php/ast"
	"github.com/heyuuu/gophp/utils/slices"
	"log"
)

func ParseAstFile(astNodes []ast.Stmt) *File {
	var defaultStmts []Stmt

	f := &File{}
	for _, astStmt := range astNodes {
		switch s := astStmt.(type) {
		case *ast.NamespaceStmt:
			f.Segments = append(f.Segments, parseNamespaceStmt(s))
		default:
			defaultStmts = append(defaultStmts, parseAstStmt(s))
		}
	}

	if len(defaultStmts) > 0 {
		f.Segments = append(
			[]Segment{buildSegment("", defaultStmts)},
			f.Segments...,
		)
	}

	return f
}

func parseNamespaceStmt(n *ast.NamespaceStmt) Segment {
	var namespace string
	if n.Name != nil {
		namespace = n.Name.ToString()
	}
	stmts := slices.Map(n.Stmts, parseAstStmt)
	return buildSegment(namespace, stmts)
}

func buildSegment(namespace string, stmts []Stmt) Segment {
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

func ParseAst(node any) any {
	switch n := node.(type) {
	case ast.Node:
		return parseAstNode(n)
	case []ast.Stmt:
		return slices.Map(n, parseAstStmt)
	default:
		return n
	}
}

func parseAssert(cond bool, message string) {
	if !cond {
		log.Fatal(message)
	}
}

// const types
func parseAstFlags(flags ast.Flags) Flags         { return Flags(flags) }
func parseAstUseType(useType ast.UseType) UseType { return UseType(useType) }

// interface types
func parseAstNode(node ast.Node) Node {
	switch n := node.(type) {
	case ast.Expr:
		return parseAstExpr(n)
	case ast.Stmt:
		return parseAstStmt(n)
	case *ast.Ident:
		return parseAstIdent(n)
	case *ast.Name:
		return parseAstName(n)
	case *ast.Arg:
		return parseAstArg(n)
	case *ast.Param:
		return parseAstParam(n)
	case *ast.Attribute:
		return parseAstAttribute(n)
	case *ast.AttributeGroup:
		return parseAstAttributeGroup(n)
	case *ast.Const:
		return parseAstConst(n)
	case *ast.MatchArm:
		return parseAstMatchArm(n)
	case *ast.VariadicPlaceholder:
		return &VariadicPlaceholder{}
	}
	return nil
}

func parseAstExpr(node ast.Expr) Expr {
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
			Items: slices.Map(n.Items, parseAstArrayItemExpr),
		}
	case *ast.ArrayItemExpr:
		return &ArrayItemExpr{
			Key:    parseAstExpr(n.Key),
			Value:  parseAstExpr(n.Value),
			ByRef:  n.ByRef,
			Unpack: n.Unpack,
		}
	case *ast.ClosureExpr:
		return &ClosureExpr{
			Static:     n.Static,
			ByRef:      n.ByRef,
			Params:     slices.Map(n.Params, parseAstParam),
			Uses:       slices.Map(n.Uses, parseAstClosureUseExpr),
			ReturnType: parseAstType(n.ReturnType),
			Stmts:      slices.Map(n.Stmts, parseAstStmt),
			AttrGroups: slices.Map(n.AttrGroups, parseAstAttributeGroup),
		}
	case *ast.ClosureUseExpr:
		return &ClosureUseExpr{
			Var:   parseAstVariableExpr(n.Var),
			ByRef: n.ByRef,
		}
	case *ast.ArrowFunctionExpr:
		return &ArrowFunctionExpr{
			Static:     n.Static,
			ByRef:      n.ByRef,
			Params:     slices.Map(n.Params, parseAstParam),
			ReturnType: parseAstType(n.ReturnType),
			Expr:       parseAstExpr(n.Expr),
			AttrGroups: slices.Map(n.AttrGroups, parseAstAttributeGroup),
		}
	case *ast.IndexExpr:
		return &IndexExpr{
			Var: parseAstExpr(n.Var),
			Dim: parseAstExpr(n.Dim),
		}
	case *ast.CastExpr:
		return &CastExpr{
			Op:   n.Op,
			Expr: parseAstExpr(n.Expr),
		}
	case *ast.UnaryExpr:
		return &UnaryExpr{
			Kind: n.Kind,
			Var:  parseAstExpr(n.Var),
		}
	case *ast.BinaryExpr:
		return &BinaryExpr{
			Op:    n.Op,
			Left:  parseAstExpr(n.Left),
			Right: parseAstExpr(n.Right),
		}
	case *ast.AssignExpr:
		return &AssignExpr{
			Op:   n.Op,
			Var:  parseAstExpr(n.Var),
			Expr: parseAstExpr(n.Expr),
		}
	case *ast.AssignRefExpr:
		return &AssignRefExpr{
			Var:  parseAstExpr(n.Var),
			Expr: parseAstExpr(n.Expr),
		}
	case *ast.InternalCallExpr:
		return &InternalCallExpr{
			Kind: n.Kind,
			Args: slices.Map(n.Args, parseAstExpr),
		}
	case *ast.CloneExpr:
		return &CloneExpr{
			Expr: parseAstExpr(n.Expr),
		}
	case *ast.ErrorSuppressExpr:
		return &ErrorSuppressExpr{
			Expr: parseAstExpr(n.Expr),
		}
	case *ast.ExitExpr:
		return &ExitExpr{
			Expr: parseAstExpr(n.Expr),
		}
	case *ast.ConstFetchExpr:
		return &ConstFetchExpr{
			Name: parseAstName(n.Name),
		}
	case *ast.ClassConstFetchExpr:
		return &ClassConstFetchExpr{
			Class: parseAstNode(n.Class),
			Name:  parseAstIdent(n.Name),
		}
	case *ast.MagicConstExpr:
		return &MagicConstExpr{
			Kind: n.Kind,
		}
	case *ast.MatchExpr:
		return &MatchExpr{
			Cond: parseAstExpr(n.Cond),
			Arms: slices.Map(n.Arms, parseAstMatchArm),
		}
	case *ast.InstanceofExpr:
		return &InstanceofExpr{
			Expr:  parseAstExpr(n.Expr),
			Class: parseAstNode(n.Class),
		}
	case *ast.ListExpr:
		return &ListExpr{
			Items: slices.Map(n.Items, parseAstArrayItemExpr),
		}
	case *ast.PrintExpr:
		return &PrintExpr{
			Expr: parseAstExpr(n.Expr),
		}
	case *ast.PropertyFetchExpr:
		return &PropertyFetchExpr{
			Var:  parseAstExpr(n.Var),
			Name: parseAstNode(n.Name),
		}
	case *ast.NullsafePropertyFetchExpr:
		return &NullsafePropertyFetchExpr{
			Var:  parseAstExpr(n.Var),
			Name: parseAstNode(n.Name),
		}
	case *ast.StaticPropertyFetchExpr:
		return &StaticPropertyFetchExpr{
			Class: parseAstNode(n.Class),
			Name:  parseAstNode(n.Name),
		}
	case *ast.ShellExecExpr:
		return &ShellExecExpr{
			Parts: slices.Map(n.Parts, parseAstExpr),
		}
	case *ast.TernaryExpr:
		return &TernaryExpr{
			Cond: parseAstExpr(n.Cond),
			If:   parseAstExpr(n.If),
			Else: parseAstExpr(n.Else),
		}
	case *ast.ThrowExpr:
		return &ThrowExpr{
			Expr: parseAstExpr(n.Expr),
		}
	case *ast.VariableExpr:
		return &VariableExpr{
			Name: parseAstNode(n.Name),
		}
	case *ast.YieldExpr:
		return &YieldExpr{
			Key:   parseAstExpr(n.Key),
			Value: parseAstExpr(n.Value),
		}
	case *ast.YieldFromExpr:
		return &YieldFromExpr{
			Expr: parseAstExpr(n.Expr),
		}
	case *ast.FuncCallExpr:
		return &FuncCallExpr{
			Name: parseAstNode(n.Name),
			Args: slices.Map(n.Args, parseAstNode),
		}
	case *ast.NewExpr:
		return &NewExpr{
			Class: parseAstNode(n.Class),
			Args:  slices.Map(n.Args, parseAstNode),
		}
	case *ast.MethodCallExpr:
		return &MethodCallExpr{
			Var:  parseAstExpr(n.Var),
			Name: parseAstNode(n.Name),
			Args: slices.Map(n.Args, parseAstNode),
		}
	case *ast.NullsafeMethodCallExpr:
		return &NullsafeMethodCallExpr{
			Var:  parseAstExpr(n.Var),
			Name: parseAstNode(n.Name),
			Args: slices.Map(n.Args, parseAstNode),
		}
	case *ast.StaticCallExpr:
		return &StaticCallExpr{
			Class: parseAstNode(n.Class),
			Name:  parseAstNode(n.Name),
			Args:  slices.Map(n.Args, parseAstNode),
		}
	}
	return nil
}

func parseAstStmt(node ast.Stmt) Stmt {
	switch n := node.(type) {
	case *ast.EmptyStmt:
		return &EmptyStmt{}
	case *ast.BlockStmt:
		return &BlockStmt{
			List: slices.Map(n.List, parseAstStmt),
		}
	case *ast.ExprStmt:
		return &ExprStmt{
			Expr: parseAstExpr(n.Expr),
		}
	case *ast.ReturnStmt:
		return &ReturnStmt{
			Expr: parseAstExpr(n.Expr),
		}
	case *ast.LabelStmt:
		return &LabelStmt{
			Name: parseAstIdent(n.Name),
		}
	case *ast.GotoStmt:
		return &GotoStmt{
			Name: parseAstIdent(n.Name),
		}
	case *ast.IfStmt:
		return &IfStmt{
			Cond:    parseAstExpr(n.Cond),
			Stmts:   slices.Map(n.Stmts, parseAstStmt),
			Elseifs: slices.Map(n.Elseifs, parseAstElseIfStmt),
			Else:    parseAstElseStmt(n.Else),
		}
	case *ast.ElseIfStmt:
		return &ElseIfStmt{
			Cond:  parseAstExpr(n.Cond),
			Stmts: slices.Map(n.Stmts, parseAstStmt),
		}
	case *ast.ElseStmt:
		return &ElseStmt{
			Stmts: slices.Map(n.Stmts, parseAstStmt),
		}
	case *ast.SwitchStmt:
		return &SwitchStmt{
			Cond:  parseAstExpr(n.Cond),
			Cases: slices.Map(n.Cases, parseAstCaseStmt),
		}
	case *ast.CaseStmt:
		return &CaseStmt{
			Cond:  parseAstExpr(n.Cond),
			Stmts: slices.Map(n.Stmts, parseAstStmt),
		}
	case *ast.ForStmt:
		return &ForStmt{
			Init:  slices.Map(n.Init, parseAstExpr),
			Cond:  slices.Map(n.Cond, parseAstExpr),
			Loop:  slices.Map(n.Loop, parseAstExpr),
			Stmts: slices.Map(n.Stmts, parseAstStmt),
		}
	case *ast.ForeachStmt:
		return &ForeachStmt{
			Expr:     parseAstExpr(n.Expr),
			KeyVar:   parseAstExpr(n.KeyVar),
			ByRef:    n.ByRef,
			ValueVar: parseAstExpr(n.ValueVar),
			Stmts:    slices.Map(n.Stmts, parseAstStmt),
		}
	case *ast.BreakStmt:
		return &BreakStmt{
			Num: parseAstExpr(n.Num),
		}
	case *ast.ContinueStmt:
		return &ContinueStmt{
			Num: parseAstExpr(n.Num),
		}
	case *ast.WhileStmt:
		return &WhileStmt{
			Cond:  parseAstExpr(n.Cond),
			Stmts: slices.Map(n.Stmts, parseAstStmt),
		}
	case *ast.DoStmt:
		return &DoStmt{
			Stmts: slices.Map(n.Stmts, parseAstStmt),
			Cond:  parseAstExpr(n.Cond),
		}
	case *ast.TryCatchStmt:
		return &TryCatchStmt{
			Stmts:   slices.Map(n.Stmts, parseAstStmt),
			Catches: slices.Map(n.Catches, parseAstCatchStmt),
			Finally: parseAstFinallyStmt(n.Finally),
		}
	case *ast.CatchStmt:
		return &CatchStmt{
			Types: slices.Map(n.Types, parseAstName),
			Var:   parseAstVariableExpr(n.Var),
			Stmts: slices.Map(n.Stmts, parseAstStmt),
		}
	case *ast.FinallyStmt:
		return &FinallyStmt{
			Stmts: slices.Map(n.Stmts, parseAstStmt),
		}
	case *ast.ConstStmt:
		return &ConstStmt{
			Consts: slices.Map(n.Consts, parseAstConst),
		}
	case *ast.EchoStmt:
		return &EchoStmt{
			Exprs: slices.Map(n.Exprs, parseAstExpr),
		}
	case *ast.GlobalStmt:
		return &GlobalStmt{
			Vars: slices.Map(n.Vars, parseAstExpr),
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
			Vars: slices.Map(n.Vars, parseAstStaticVarStmt),
		}
	case *ast.StaticVarStmt:
		return &StaticVarStmt{
			Var:     parseAstVariableExpr(n.Var),
			Default: parseAstExpr(n.Default),
		}
	case *ast.UnsetStmt:
		return &UnsetStmt{
			Vars: slices.Map(n.Vars, parseAstExpr),
		}
	case *ast.UseStmt:
		return &UseStmt{
			Type:  parseAstUseType(n.Type),
			Name:  parseAstName(n.Name),
			Alias: parseAstIdent(n.Alias),
		}
	case *ast.DeclareStmt:
		return &DeclareStmt{
			Declares: slices.Map(n.Declares, parseAstDeclareDeclareStmt),
			Stmts:    slices.Map(n.Stmts, parseAstStmt),
		}
	case *ast.DeclareDeclareStmt:
		return &DeclareDeclareStmt{
			Key:   parseAstIdent(n.Key),
			Value: parseAstExpr(n.Value),
		}
	case *ast.NamespaceStmt:
		return &NamespaceStmt{
			Name:  parseAstName(n.Name),
			Stmts: slices.Map(n.Stmts, parseAstStmt),
		}
	case *ast.FunctionStmt:
		parseAssert(n.NamespacedName != nil, "FunctionStmt.NamespacedName cannot be nil")

		return &FunctionStmt{
			Name:       parseAstNameAsFQ(n.NamespacedName),
			ByRef:      n.ByRef,
			Params:     slices.Map(n.Params, parseAstParam),
			ReturnType: parseAstType(n.ReturnType),
			Stmts:      slices.Map(n.Stmts, parseAstStmt),
			AttrGroups: slices.Map(n.AttrGroups, parseAstAttributeGroup),
		}
	case *ast.InterfaceStmt:
		parseAssert(n.NamespacedName != nil, "InterfaceStmt.NamespacedName cannot be nil")

		return &InterfaceStmt{
			Name:       parseAstNameAsFQ(n.NamespacedName),
			Extends:    slices.Map(n.Extends, parseAstName),
			Stmts:      slices.Map(n.Stmts, parseAstStmt),
			AttrGroups: slices.Map(n.AttrGroups, parseAstAttributeGroup),
		}
	case *ast.ClassStmt:
		// todo 将匿名类和实名类定义区分开
		var name *Name
		if n.NamespacedName != nil {
			name = parseAstNameAsFQ(n.NamespacedName)
		}

		return &ClassStmt{
			Name:       name,
			Flags:      parseAstFlags(n.Flags),
			Extends:    parseAstName(n.Extends),
			Implements: slices.Map(n.Implements, parseAstName),
			Stmts:      slices.Map(n.Stmts, parseAstStmt),
			AttrGroups: slices.Map(n.AttrGroups, parseAstAttributeGroup),
		}
	case *ast.ClassConstStmt:
		return &ClassConstStmt{
			Flags:      parseAstFlags(n.Flags),
			Consts:     slices.Map(n.Consts, parseAstConst),
			AttrGroups: slices.Map(n.AttrGroups, parseAstAttributeGroup),
		}
	case *ast.PropertyStmt:
		return &PropertyStmt{
			Flags:      parseAstFlags(n.Flags),
			Props:      slices.Map(n.Props, parseAstPropertyPropertyStmt),
			Type:       parseAstType(n.Type),
			AttrGroups: slices.Map(n.AttrGroups, parseAstAttributeGroup),
		}
	case *ast.PropertyPropertyStmt:
		return &PropertyPropertyStmt{
			Name:    parseAstIdent(n.Name),
			Default: parseAstExpr(n.Default),
		}
	case *ast.ClassMethodStmt:
		return &ClassMethodStmt{
			Flags:      parseAstFlags(n.Flags),
			ByRef:      n.ByRef,
			Name:       parseAstIdent(n.Name),
			Params:     slices.Map(n.Params, parseAstParam),
			ReturnType: parseAstType(n.ReturnType),
			Stmts:      slices.Map(n.Stmts, parseAstStmt),
			AttrGroups: slices.Map(n.AttrGroups, parseAstAttributeGroup),
		}
	case *ast.TraitStmt:
		parseAssert(n.NamespacedName != nil, "TraitStmt.NamespacedName cannot be nil")

		return &TraitStmt{
			Name:       parseAstNameAsFQ(n.NamespacedName),
			Stmts:      slices.Map(n.Stmts, parseAstStmt),
			AttrGroups: slices.Map(n.AttrGroups, parseAstAttributeGroup),
		}
	case *ast.TraitUseStmt:
		return &TraitUseStmt{
			Traits:      slices.Map(n.Traits, parseAstName),
			Adaptations: slices.Map(n.Adaptations, parseAstTraitUseAdaptationStmt),
		}
	case *ast.TraitUseAdaptationAliasStmt:
		return &TraitUseAdaptationAliasStmt{
			NewModifier: parseAstFlags(n.NewModifier),
			NewName:     parseAstIdent(n.NewName),
			Trait:       parseAstName(n.Trait),
			Method:      parseAstIdent(n.Method),
		}
	case *ast.TraitUseAdaptationPrecedenceStmt:
		return &TraitUseAdaptationPrecedenceStmt{
			Insteadof: slices.Map(n.Insteadof, parseAstName),
			Trait:     parseAstName(n.Trait),
			Method:    parseAstIdent(n.Method),
		}
	case *ast.EnumStmt:
		return &EnumStmt{
			ScalarType:     parseAstIdent(n.ScalarType),
			Implements:     slices.Map(n.Implements, parseAstName),
			Name:           parseAstIdent(n.Name),
			Stmts:          slices.Map(n.Stmts, parseAstStmt),
			AttrGroups:     slices.Map(n.AttrGroups, parseAstAttributeGroup),
			NamespacedName: parseAstName(n.NamespacedName),
		}
	case *ast.EnumCaseStmt:
		return &EnumCaseStmt{
			Name:       parseAstIdent(n.Name),
			Expr:       parseAstExpr(n.Expr),
			AttrGroups: slices.Map(n.AttrGroups, parseAstAttributeGroup),
		}
	}
	return nil
}

func parseAstType(node ast.Type) Type {
	switch n := node.(type) {
	case *ast.SimpleType:
		return &SimpleType{
			Name: parseAstName(n.Name),
		}
	case *ast.IntersectionType:
		return &IntersectionType{
			Types: slices.Map(n.Types, parseAstType),
		}
	case *ast.UnionType:
		return &UnionType{
			Types: slices.Map(n.Types, parseAstType),
		}
	case *ast.NullableType:
		return &NullableType{
			Type: parseAstSimpleType(n.Type),
		}
	}
	return nil
}

func parseAstTraitUseAdaptationStmt(node ast.TraitUseAdaptationStmt) TraitUseAdaptationStmt {
	switch n := node.(type) {
	case *ast.TraitUseAdaptationAliasStmt:
		return &TraitUseAdaptationAliasStmt{
			NewModifier: parseAstFlags(n.NewModifier),
			NewName:     parseAstIdent(n.NewName),
			Trait:       parseAstName(n.Trait),
			Method:      parseAstIdent(n.Method),
		}
	case *ast.TraitUseAdaptationPrecedenceStmt:
		return &TraitUseAdaptationPrecedenceStmt{
			Insteadof: slices.Map(n.Insteadof, parseAstName),
			Trait:     parseAstName(n.Trait),
			Method:    parseAstIdent(n.Method),
		}
	}
	return nil
}

// struct types
func parseAstArg(n *ast.Arg) *Arg {
	return &Arg{
		Name:   parseAstIdent(n.Name),
		Value:  parseAstExpr(n.Value),
		ByRef:  n.ByRef,
		Unpack: n.Unpack,
	}
}
func parseAstAttribute(n *ast.Attribute) *Attribute {
	return &Attribute{
		Name: parseAstName(n.Name),
		Args: slices.Map(n.Args, parseAstArg),
	}
}
func parseAstAttributeGroup(n *ast.AttributeGroup) *AttributeGroup {
	return &AttributeGroup{
		Attrs: slices.Map(n.Attrs, parseAstAttribute),
	}
}
func parseAstConst(n *ast.Const) *Const {
	return &Const{
		Name:           parseAstIdent(n.Name),
		Value:          parseAstExpr(n.Value),
		NamespacedName: parseAstName(n.NamespacedName),
	}
}
func parseAstIdent(n *ast.Ident) *Ident {
	if n == nil {
		return nil
	}
	return &Ident{
		Name:    n.Name,
		VarLike: n.VarLike,
	}
}
func parseAstMatchArm(n *ast.MatchArm) *MatchArm {
	if n == nil {
		return nil
	}
	return &MatchArm{
		Conds: slices.Map(n.Conds, parseAstExpr),
		Body:  parseAstExpr(n.Body),
	}
}
func parseAstParam(n *ast.Param) *Param {
	if n == nil {
		return nil
	}
	return &Param{
		Type:       parseAstType(n.Type),
		ByRef:      n.ByRef,
		Variadic:   n.Variadic,
		Var:        parseAstVariableExpr(n.Var),
		Default:    parseAstExpr(n.Default),
		Flags:      parseAstFlags(n.Flags),
		AttrGroups: slices.Map(n.AttrGroups, parseAstAttributeGroup),
	}
}
func parseAstSimpleType(n *ast.SimpleType) *SimpleType {
	if n == nil {
		return nil
	}
	return &SimpleType{
		Name: parseAstName(n.Name),
	}
}

func parseAstNameAsFQ(n *ast.Name) *Name {
	return NewName(NameFullyQualified, n.Parts)
}

func parseAstName(n *ast.Name) *Name {
	if n == nil {
		return nil
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
func parseAstArrayItemExpr(n *ast.ArrayItemExpr) *ArrayItemExpr {
	if n == nil {
		return nil
	}
	return &ArrayItemExpr{
		Key:    parseAstExpr(n.Key),
		Value:  parseAstExpr(n.Value),
		ByRef:  n.ByRef,
		Unpack: n.Unpack,
	}
}
func parseAstClosureUseExpr(n *ast.ClosureUseExpr) *ClosureUseExpr {
	if n == nil {
		return nil
	}
	return &ClosureUseExpr{
		Var:   parseAstVariableExpr(n.Var),
		ByRef: n.ByRef,
	}
}
func parseAstVariableExpr(n *ast.VariableExpr) *VariableExpr {
	if n == nil {
		return nil
	}
	return &VariableExpr{
		Name: parseAstNode(n.Name),
	}
}
func parseAstElseIfStmt(n *ast.ElseIfStmt) *ElseIfStmt {
	if n == nil {
		return nil
	}
	return &ElseIfStmt{
		Cond:  parseAstExpr(n.Cond),
		Stmts: slices.Map(n.Stmts, parseAstStmt),
	}
}
func parseAstElseStmt(n *ast.ElseStmt) *ElseStmt {
	if n == nil {
		return nil
	}
	return &ElseStmt{
		Stmts: slices.Map(n.Stmts, parseAstStmt),
	}
}
func parseAstCaseStmt(n *ast.CaseStmt) *CaseStmt {
	if n == nil {
		return nil
	}
	return &CaseStmt{
		Cond:  parseAstExpr(n.Cond),
		Stmts: slices.Map(n.Stmts, parseAstStmt),
	}
}
func parseAstCatchStmt(n *ast.CatchStmt) *CatchStmt {
	if n == nil {
		return nil
	}
	return &CatchStmt{
		Types: slices.Map(n.Types, parseAstName),
		Var:   parseAstVariableExpr(n.Var),
		Stmts: slices.Map(n.Stmts, parseAstStmt),
	}
}
func parseAstFinallyStmt(n *ast.FinallyStmt) *FinallyStmt {
	if n == nil {
		return nil
	}
	return &FinallyStmt{
		Stmts: slices.Map(n.Stmts, parseAstStmt),
	}
}
func parseAstStaticVarStmt(n *ast.StaticVarStmt) *StaticVarStmt {
	if n == nil {
		return nil
	}
	return &StaticVarStmt{
		Var:     parseAstVariableExpr(n.Var),
		Default: parseAstExpr(n.Default),
	}
}
func parseAstDeclareDeclareStmt(n *ast.DeclareDeclareStmt) *DeclareDeclareStmt {
	if n == nil {
		return nil
	}
	return &DeclareDeclareStmt{
		Key:   parseAstIdent(n.Key),
		Value: parseAstExpr(n.Value),
	}
}
func parseAstPropertyPropertyStmt(n *ast.PropertyPropertyStmt) *PropertyPropertyStmt {
	if n == nil {
		return nil
	}
	return &PropertyPropertyStmt{
		Name:    parseAstIdent(n.Name),
		Default: parseAstExpr(n.Default),
	}
}
