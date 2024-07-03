package transformer

import (
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/ir"
	"github.com/heyuuu/gophp/kits/slicekit"
	"log"
	"runtime"
	"slices"
	"strings"
)

func TransformFile(f *ast.File) (*ir.File, error) {
	t := &transformer{}
	return t.TransformFile(f)
}

type transformer struct{}

func (t *transformer) TransformFile(f *ast.File) (result *ir.File, resultErr error) {
	defer func() {
		if e := recover(); e != nil {
			if err, ok := e.(error); ok {
				resultErr = err
			} else {
				resultErr = fmt.Errorf("transform.transformer.TransformFile() error: %v", e)
			}

			// 打印堆栈
			const size = 64 << 10
			stack := make([]byte, size)
			stack = stack[:runtime.Stack(stack, false)]
			log.Printf(">>> tests.Command.RunBuiltin() panic: %v\n%s", e, stack)
		}
	}()

	return t.file(f), nil
}

func (t *transformer) unexpected(msg string) {
	err := fmt.Errorf("transformer.transformer: %s", msg)
	panic(err)
}

func (t *transformer) unexpectedType(n ast.Node) {
	err := fmt.Errorf("transformer.transformer: unexpected node type %T", n)
	panic(err)
}

// file
func (t *transformer) file(f *ast.File) *ir.File {
	strictTypes := t.declaresToStrictTypes(f.Declares)

	return &ir.File{
		StrictTypes: strictTypes,
		Namespaces:  slicekit.Map(f.Namespaces, t.namespaceStmt),
	}
}

func (t *transformer) declaresToStrictTypes(declares []*ast.DeclareStmt) bool {
	var strictTypes bool
	for _, declare := range declares {
		if len(declare.Stmts) != 0 {
			t.unexpected("declare statement must not have statements")
		}
		for _, kv := range declare.Declares {
			lcKey := strings.ToLower(kv.Key.Name)
			if lcKey != "strict_types" {
				t.unexpected("declare statement only support `strict_types`")
			}

			// strict_types
			if intLit, ok := kv.Value.(*ast.IntLit); ok {
				strictTypes = intLit.Value != 0
			} else {
				t.unexpected("declare statement `strict_types` must be int literal")
			}
		}
	}
	return strictTypes
}

// node
func (t *transformer) node(n ast.Node) ir.Node {
	switch x := n.(type) {
	case nil:
		return nil
	case *ast.Arg:
		return t.arg(x)
	case *ast.Name:
		return t.name(x)
	case *ast.Ident:
		return t.ident(x)
	case *ast.Param:
		return t.param(x)
	case *ast.VariableExpr:
		return t.variableExpr(x)
	}

	// unreachable
	t.unexpectedType(n)
	return nil
}

func (t *transformer) arg(x *ast.Arg) *ir.Arg {
	return &ir.Arg{
		Value:  t.expr(x.Value),
		Unpack: x.Unpack,
	}
}

func (t *transformer) name(x *ast.Name) *ir.Name {
	if x == nil {
		return nil
	}
	return &ir.Name{
		Kind:  x.Kind,
		Parts: slices.Clone(x.Parts),
	}
}

func (t *transformer) ident(x *ast.Ident) *ir.Ident {
	if x == nil {
		return nil
	}
	return &ir.Ident{
		Name:    x.Name,
		VarLike: x.VarLike,
	}
}

func (t *transformer) param(x *ast.Param) *ir.Param {
	return &ir.Param{
		Type:     t.typeHint(x.Type),
		ByRef:    x.ByRef,
		Variadic: x.Variadic,
		Var:      t.variableExpr(x.Var),
		Default:  t.expr(x.Default),
	}
}

// typeHint
func (t *transformer) typeHint(n ast.TypeHint) ir.TypeHint {
	switch x := n.(type) {
	case *ast.SimpleType:
		return t.simpleType(x)
	case *ast.IntersectionType:
		return t.intersectionType(x)
	case *ast.UnionType:
		return t.unionType(x)
	case *ast.NullableType:
		return t.nullableType(x)

	}

	// unreachable
	t.unexpectedType(n)
	return nil
}

func (t *transformer) simpleType(x *ast.SimpleType) *ir.SimpleType {
	return &ir.SimpleType{
		Name: t.name(x.Name),
	}
}

func (t *transformer) intersectionType(x *ast.IntersectionType) *ir.IntersectionType {
	return &ir.IntersectionType{
		Types: t.typeHintList(x.Types),
	}
}

func (t *transformer) unionType(x *ast.UnionType) *ir.UnionType {
	return &ir.UnionType{
		Types: t.typeHintList(x.Types),
	}
}

func (t *transformer) nullableType(x *ast.NullableType) *ir.NullableType {
	return &ir.NullableType{
		Type: t.simpleType(x.Type),
	}
}

// expr
func (t *transformer) expr(n ast.Expr) ir.Expr {
	switch x := n.(type) {
	case nil:
		return nil
	case *ast.IntLit:
		return t.intLit(x)
	case *ast.FloatLit:
		return t.floatLit(x)
	case *ast.StringLit:
		return t.stringLit(x)
	case *ast.ArrayExpr:
		return t.arrayExpr(x)
	case *ast.ArrayItemExpr:
		return t.arrayItemExpr(x)
	case *ast.ClosureExpr:
		return t.closureExpr(x)
	case *ast.ClosureUseExpr:
		return t.closureUseExpr(x)
	case *ast.ArrowFunctionExpr:
		return t.arrowFunctionExpr(x)
	case *ast.IndexExpr:
		return t.indexExpr(x)
	case *ast.CastExpr:
		return t.castExpr(x)
	case *ast.UnaryExpr:
		return t.unaryExpr(x)
	case *ast.BinaryOpExpr:
		return t.binaryOpExpr(x)
	case *ast.AssignExpr:
		return t.assignExpr(x)
	case *ast.AssignOpExpr:
		return t.assignOpExpr(x)
	case *ast.AssignRefExpr:
		return t.assignRefExpr(x)
	case *ast.IssetExpr:
		return t.issetExpr(x)
	case *ast.EmptyExpr:
		return t.emptyExpr(x)
	case *ast.EvalExpr:
		return t.evalExpr(x)
	case *ast.IncludeExpr:
		return t.includeExpr(x)
	case *ast.CloneExpr:
		return t.cloneExpr(x)
	case *ast.ErrorSuppressExpr:
		return t.errorSuppressExpr(x)
	case *ast.ExitExpr:
		return t.exitExpr(x)
	case *ast.ConstFetchExpr:
		return t.constFetchExpr(x)
	case *ast.ClassConstFetchExpr:
		return t.classConstFetchExpr(x)
	case *ast.MagicConstExpr:
		return t.magicConstExpr(x)
	case *ast.InstanceofExpr:
		return t.instanceofExpr(x)
	case *ast.ListExpr:
		return t.listExpr(x)
	case *ast.PrintExpr:
		return t.printExpr(x)
	case *ast.PropertyFetchExpr:
		return t.propertyFetchExpr(x)
	case *ast.StaticPropertyFetchExpr:
		return t.staticPropertyFetchExpr(x)
	case *ast.ShellExecExpr:
		return t.shellExecExpr(x)
	case *ast.TernaryExpr:
		return t.ternaryExpr(x)
	case *ast.ThrowExpr:
		return t.throwExpr(x)
	case *ast.VariableExpr:
		return t.variableExpr(x)
	case *ast.YieldExpr:
		return t.yieldExpr(x)
	case *ast.YieldFromExpr:
		return t.yieldFromExpr(x)
	case *ast.FuncCallExpr:
		return t.funcCallExpr(x)
	case *ast.NewExpr:
		return t.newExpr(x)
	case *ast.MethodCallExpr:
		return t.methodCallExpr(x)
	case *ast.StaticCallExpr:
		return t.staticCallExpr(x)
	}

	// unreachable
	t.unexpectedType(n)
	return nil
}

func (t *transformer) intLit(x *ast.IntLit) *ir.IntLit {
	return &ir.IntLit{
		Raw:   ast.MetaRawValue(x),
		Value: x.Value,
	}
}

func (t *transformer) floatLit(x *ast.FloatLit) *ir.FloatLit {
	return &ir.FloatLit{
		Raw:   ast.MetaRawValue(x),
		Value: x.Value,
	}
}

func (t *transformer) stringLit(x *ast.StringLit) *ir.StringLit {
	return &ir.StringLit{
		Raw:   ast.MetaRawValue(x),
		Value: x.Value,
	}
}

func (t *transformer) arrayExpr(x *ast.ArrayExpr) *ir.ArrayExpr {
	return &ir.ArrayExpr{
		Items: slicekit.Map(x.Items, t.arrayItemExpr),
	}
}

func (t *transformer) arrayItemExpr(x *ast.ArrayItemExpr) *ir.ArrayItemExpr {
	return &ir.ArrayItemExpr{
		Key:    t.expr(x.Key),
		Value:  t.expr(x.Value),
		ByRef:  x.ByRef,
		Unpack: x.Unpack,
	}
}

func (t *transformer) closureExpr(x *ast.ClosureExpr) *ir.ClosureExpr {
	return &ir.ClosureExpr{
		Static:     x.Static,
		ByRef:      x.ByRef,
		Params:     t.paramList(x.Params),
		Uses:       slicekit.Map(x.Uses, t.closureUseExpr),
		ReturnType: t.typeHint(x.ReturnType),
		Stmts:      t.stmtList(x.Stmts),
	}
}

func (t *transformer) closureUseExpr(x *ast.ClosureUseExpr) *ir.ClosureUseExpr {
	return &ir.ClosureUseExpr{
		Var:   t.variableExpr(x.Var),
		ByRef: x.ByRef,
	}
}

func (t *transformer) arrowFunctionExpr(x *ast.ArrowFunctionExpr) *ir.ArrowFunctionExpr {
	return &ir.ArrowFunctionExpr{
		Static:     x.Static,
		ByRef:      x.ByRef,
		Params:     t.paramList(x.Params),
		ReturnType: t.typeHint(x.ReturnType),
		Expr:       t.expr(x.Expr),
	}
}

func (t *transformer) indexExpr(x *ast.IndexExpr) *ir.IndexExpr {
	return &ir.IndexExpr{
		Var: t.expr(x.Var),
		Dim: t.expr(x.Dim),
	}
}

func (t *transformer) castExpr(x *ast.CastExpr) *ir.CastExpr {
	return &ir.CastExpr{
		Kind: x.Kind,
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) unaryExpr(x *ast.UnaryExpr) *ir.UnaryExpr {
	return &ir.UnaryExpr{
		Op:  x.Op,
		Var: t.expr(x.Var),
	}
}

func (t *transformer) binaryOpExpr(x *ast.BinaryOpExpr) *ir.BinaryOpExpr {
	return &ir.BinaryOpExpr{
		Op:    x.Op,
		Left:  t.expr(x.Left),
		Right: t.expr(x.Right),
	}
}

func (t *transformer) assignExpr(x *ast.AssignExpr) *ir.AssignExpr {
	return &ir.AssignExpr{
		Var:  t.expr(x.Var),
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) assignOpExpr(x *ast.AssignOpExpr) *ir.AssignOpExpr {
	return &ir.AssignOpExpr{
		Op:   x.Op,
		Var:  t.expr(x.Var),
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) assignRefExpr(x *ast.AssignRefExpr) *ir.AssignRefExpr {
	return &ir.AssignRefExpr{
		Var:  t.expr(x.Var),
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) issetExpr(x *ast.IssetExpr) *ir.IssetExpr {
	return &ir.IssetExpr{
		Vars: t.exprList(x.Vars),
	}
}

func (t *transformer) emptyExpr(x *ast.EmptyExpr) *ir.EmptyExpr {
	return &ir.EmptyExpr{
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) evalExpr(x *ast.EvalExpr) *ir.EvalExpr {
	return &ir.EvalExpr{
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) includeExpr(x *ast.IncludeExpr) *ir.IncludeExpr {
	return &ir.IncludeExpr{
		Kind: x.Kind,
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) cloneExpr(x *ast.CloneExpr) *ir.CloneExpr {
	return &ir.CloneExpr{
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) errorSuppressExpr(x *ast.ErrorSuppressExpr) *ir.ErrorSuppressExpr {
	return &ir.ErrorSuppressExpr{
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) exitExpr(x *ast.ExitExpr) *ir.ExitExpr {
	return &ir.ExitExpr{
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) constFetchExpr(x *ast.ConstFetchExpr) ir.Expr {
	switch strings.ToLower(x.Name.ToString()) {
	case "null":
		return &ir.NullLit{}
	case "true":
		return &ir.BoolLit{Value: true}
	case "false":
		return &ir.BoolLit{Value: false}
	}

	return &ir.ConstFetchExpr{
		Name: t.name(x.Name),
	}
}

func (t *transformer) classConstFetchExpr(x *ast.ClassConstFetchExpr) *ir.ClassConstFetchExpr {
	return &ir.ClassConstFetchExpr{
		Class: t.node(x.Class),
		Name:  t.ident(x.Name),
	}
}

func (t *transformer) magicConstExpr(x *ast.MagicConstExpr) *ir.MagicConstExpr {
	return &ir.MagicConstExpr{
		Kind: x.Kind,
	}
}

func (t *transformer) instanceofExpr(x *ast.InstanceofExpr) *ir.InstanceofExpr {
	return &ir.InstanceofExpr{
		Expr:  t.expr(x.Expr),
		Class: t.node(x.Class),
	}
}

func (t *transformer) listExpr(x *ast.ListExpr) *ir.ListExpr {
	return &ir.ListExpr{
		Items: slicekit.Map(x.Items, t.arrayItemExpr),
	}
}

func (t *transformer) printExpr(x *ast.PrintExpr) *ir.PrintExpr {
	return &ir.PrintExpr{
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) propertyFetchExpr(x *ast.PropertyFetchExpr) *ir.PropertyFetchExpr {
	return &ir.PropertyFetchExpr{
		Var:      t.expr(x.Var),
		Name:     t.node(x.Name),
		Nullsafe: x.Nullsafe,
	}
}

func (t *transformer) staticPropertyFetchExpr(x *ast.StaticPropertyFetchExpr) *ir.StaticPropertyFetchExpr {
	return &ir.StaticPropertyFetchExpr{
		Class: t.node(x.Class),
		Name:  t.node(x.Name),
	}
}

func (t *transformer) shellExecExpr(x *ast.ShellExecExpr) *ir.ShellExecExpr {
	return &ir.ShellExecExpr{
		Parts: t.exprList(x.Parts),
	}
}

func (t *transformer) ternaryExpr(x *ast.TernaryExpr) *ir.TernaryExpr {
	return &ir.TernaryExpr{
		Cond: t.expr(x.Cond),
		If:   t.expr(x.If),
		Else: t.expr(x.Else),
	}
}

func (t *transformer) throwExpr(x *ast.ThrowExpr) *ir.ThrowExpr {
	return &ir.ThrowExpr{
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) variableExpr(x *ast.VariableExpr) *ir.VariableExpr {
	return &ir.VariableExpr{
		Name: t.node(x.Name),
	}
}

func (t *transformer) yieldExpr(x *ast.YieldExpr) *ir.YieldExpr {
	return &ir.YieldExpr{
		Key:   t.expr(x.Key),
		Value: t.expr(x.Value),
	}
}

func (t *transformer) yieldFromExpr(x *ast.YieldFromExpr) *ir.YieldFromExpr {
	return &ir.YieldFromExpr{
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) funcCallExpr(x *ast.FuncCallExpr) *ir.FuncCallExpr {
	return &ir.FuncCallExpr{
		Name: t.node(x.Name),
		Args: t.argList(x.Args),
	}
}

func (t *transformer) newExpr(x *ast.NewExpr) *ir.NewExpr {
	return &ir.NewExpr{
		Class: t.node(x.Class),
		Args:  t.argList(x.Args),
	}
}

func (t *transformer) methodCallExpr(x *ast.MethodCallExpr) *ir.MethodCallExpr {
	return &ir.MethodCallExpr{
		Var:      t.expr(x.Var),
		Name:     t.node(x.Name),
		Args:     t.argList(x.Args),
		Nullsafe: x.Nullsafe,
	}
}

func (t *transformer) staticCallExpr(x *ast.StaticCallExpr) *ir.StaticCallExpr {
	return &ir.StaticCallExpr{
		Class: t.node(x.Class),
		Name:  t.node(x.Name),
		Args:  t.argList(x.Args),
	}
}

// stmt
func (t *transformer) stmt(n ast.Stmt) (stmt ir.Stmt) {
	defer func() {
		if stmt != nil {
			stmt.SetComments(t.comments(ast.MetaComments(n)))
		}
	}()

	switch x := n.(type) {
	case *ast.EmptyStmt:
		return t.emptyStmt(x)
	case *ast.ExprStmt:
		return t.exprStmt(x)
	case *ast.ReturnStmt:
		return t.returnStmt(x)
	case *ast.LabelStmt:
		return t.labelStmt(x)
	case *ast.GotoStmt:
		return t.gotoStmt(x)
	case *ast.IfStmt:
		return t.ifStmt(x)
	case *ast.SwitchStmt:
		return t.switchStmt(x)
	case *ast.ForStmt:
		return t.forStmt(x)
	case *ast.ForeachStmt:
		return t.foreachStmt(x)
	case *ast.BreakStmt:
		return t.breakStmt(x)
	case *ast.ContinueStmt:
		return t.continueStmt(x)
	case *ast.WhileStmt:
		return t.whileStmt(x)
	case *ast.DoStmt:
		return t.doStmt(x)
	case *ast.TryCatchStmt:
		return t.tryCatchStmt(x)
	case *ast.ConstStmt:
		return t.constStmt(x)
	case *ast.EchoStmt:
		return t.echoStmt(x)
	case *ast.GlobalStmt:
		return t.globalStmt(x)
	case *ast.InlineHTMLStmt:
		return t.inlineHTMLStmt(x)
	case *ast.StaticStmt:
		return t.staticStmt(x)
	case *ast.UnsetStmt:
		return t.unsetStmt(x)
	case *ast.UseStmt:
		return t.useStmt(x)
	case *ast.NamespaceStmt:
		return t.namespaceStmt(x)
	case *ast.FunctionStmt:
		return t.functionStmt(x)
	case *ast.InterfaceStmt:
		return t.interfaceStmt(x)
	case *ast.ClassStmt:
		return t.classStmt(x)
	case *ast.ClassConstStmt:
		return t.classConstStmt(x)
	case *ast.PropertyStmt:
		return t.propertyStmt(x)
	case *ast.ClassMethodStmt:
		return t.classMethodStmt(x)
	case *ast.TraitStmt:
		return t.traitStmt(x)
	case *ast.TraitUseStmt:
		return t.traitUseStmt(x)
	case *ast.TraitUseAdaptationAliasStmt:
		return t.traitUseAdaptationAliasStmt(x)
	case *ast.TraitUseAdaptationPrecedenceStmt:
		return t.traitUseAdaptationPrecedenceStmt(x)
	}

	// unreachable
	t.unexpectedType(n)
	return nil
}

func (t *transformer) emptyStmt(x *ast.EmptyStmt) *ir.EmptyStmt {
	return &ir.EmptyStmt{}
}

func (t *transformer) exprStmt(x *ast.ExprStmt) *ir.ExprStmt {
	return &ir.ExprStmt{
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) returnStmt(x *ast.ReturnStmt) *ir.ReturnStmt {
	return &ir.ReturnStmt{
		Expr: t.expr(x.Expr),
	}
}

func (t *transformer) labelStmt(x *ast.LabelStmt) *ir.LabelStmt {
	return &ir.LabelStmt{
		Name: t.ident(x.Name),
	}
}

func (t *transformer) gotoStmt(x *ast.GotoStmt) *ir.GotoStmt {
	return &ir.GotoStmt{
		Name: t.ident(x.Name),
	}
}

func (t *transformer) ifStmt(x *ast.IfStmt) *ir.IfStmt {
	return &ir.IfStmt{
		Cond:    t.expr(x.Cond),
		Stmts:   t.stmtList(x.Stmts),
		Elseifs: slicekit.Map(x.Elseifs, t.elseIfStmt),
		Else:    t.elseStmt(x.Else),
	}
}

func (t *transformer) elseIfStmt(x *ast.ElseIfStmt) *ir.ElseIfStmt {
	return &ir.ElseIfStmt{
		Cond:  t.expr(x.Cond),
		Stmts: t.stmtList(x.Stmts),
	}
}

func (t *transformer) elseStmt(x *ast.ElseStmt) *ir.ElseStmt {
	if x == nil {
		return nil
	}
	return &ir.ElseStmt{
		Stmts: t.stmtList(x.Stmts),
	}
}

func (t *transformer) switchStmt(x *ast.SwitchStmt) *ir.SwitchStmt {
	return &ir.SwitchStmt{
		Cond:  t.expr(x.Cond),
		Cases: slicekit.Map(x.Cases, t.caseStmt),
	}
}

func (t *transformer) caseStmt(x *ast.CaseStmt) *ir.CaseStmt {
	return &ir.CaseStmt{
		Cond:  t.expr(x.Cond),
		Stmts: t.stmtList(x.Stmts),
	}
}

func (t *transformer) forStmt(x *ast.ForStmt) *ir.ForStmt {
	return &ir.ForStmt{
		Init:  t.exprList(x.Init),
		Cond:  t.exprList(x.Cond),
		Loop:  t.exprList(x.Loop),
		Stmts: t.stmtList(x.Stmts),
	}
}

func (t *transformer) foreachStmt(x *ast.ForeachStmt) *ir.ForeachStmt {
	return &ir.ForeachStmt{
		Expr:     t.expr(x.Expr),
		KeyVar:   t.expr(x.KeyVar),
		ByRef:    x.ByRef,
		ValueVar: t.expr(x.ValueVar),
		Stmts:    t.stmtList(x.Stmts),
	}
}

func (t *transformer) breakStmt(x *ast.BreakStmt) *ir.BreakStmt {
	return &ir.BreakStmt{
		Num: t.expr(x.Num),
	}
}

func (t *transformer) continueStmt(x *ast.ContinueStmt) *ir.ContinueStmt {
	return &ir.ContinueStmt{
		Num: t.expr(x.Num),
	}
}

func (t *transformer) whileStmt(x *ast.WhileStmt) *ir.WhileStmt {
	return &ir.WhileStmt{
		Cond:  t.expr(x.Cond),
		Stmts: t.stmtList(x.Stmts),
	}
}

func (t *transformer) doStmt(x *ast.DoStmt) *ir.DoStmt {
	return &ir.DoStmt{
		Stmts: t.stmtList(x.Stmts),
		Cond:  t.expr(x.Cond),
	}
}

func (t *transformer) tryCatchStmt(x *ast.TryCatchStmt) *ir.TryCatchStmt {
	return &ir.TryCatchStmt{
		Stmts:   t.stmtList(x.Stmts),
		Catches: slicekit.Map(x.Catches, t.catchStmt),
		Finally: t.finallyStmt(x.Finally),
	}
}

func (t *transformer) catchStmt(x *ast.CatchStmt) *ir.CatchStmt {
	return &ir.CatchStmt{
		Types: t.nameList(x.Types),
		Var:   t.variableExpr(x.Var),
		Stmts: t.stmtList(x.Stmts),
	}
}

func (t *transformer) finallyStmt(x *ast.FinallyStmt) *ir.FinallyStmt {
	return &ir.FinallyStmt{
		Stmts: t.stmtList(x.Stmts),
	}
}

func (t *transformer) constStmt(x *ast.ConstStmt) *ir.ConstStmt {
	return &ir.ConstStmt{
		Name:           t.ident(x.Name),
		Value:          t.expr(x.Value),
		NamespacedName: t.name(x.NamespacedName),
	}
}

func (t *transformer) echoStmt(x *ast.EchoStmt) *ir.EchoStmt {
	return &ir.EchoStmt{
		Exprs: t.exprList(x.Exprs),
	}
}

func (t *transformer) globalStmt(x *ast.GlobalStmt) *ir.GlobalStmt {
	return &ir.GlobalStmt{
		Var: t.expr(x.Var),
	}
}

func (t *transformer) inlineHTMLStmt(x *ast.InlineHTMLStmt) ir.Stmt {
	return &ir.EchoStmt{
		Exprs: []ir.Expr{
			&ir.StringLit{Value: x.Value},
		},
	}
}

func (t *transformer) staticStmt(x *ast.StaticStmt) *ir.StaticStmt {
	return &ir.StaticStmt{
		Var:     t.variableExpr(x.Var),
		Default: t.expr(x.Default),
	}
}

func (t *transformer) unsetStmt(x *ast.UnsetStmt) *ir.UnsetStmt {
	return &ir.UnsetStmt{
		Var: t.expr(x.Var),
	}
}

func (t *transformer) useStmt(x *ast.UseStmt) *ir.UseStmt {
	return &ir.UseStmt{
		Type:  x.Type,
		Name:  t.name(x.Name),
		Alias: t.ident(x.Alias),
	}
}

func (t *transformer) namespaceStmt(x *ast.NamespaceStmt) *ir.NamespaceStmt {
	var name *ir.Name
	if x.Name != nil {
		name = t.name(x.Name)
	}
	return &ir.NamespaceStmt{
		Name:  name,
		Stmts: t.stmtList(x.Stmts),
	}
}

func (t *transformer) functionStmt(x *ast.FunctionStmt) *ir.FunctionStmt {
	return &ir.FunctionStmt{
		ByRef:          x.ByRef,
		Name:           t.ident(x.Name),
		Params:         t.paramList(x.Params),
		ReturnType:     t.typeHint(x.ReturnType),
		Stmts:          t.stmtList(x.Stmts),
		NamespacedName: t.name(x.NamespacedName),
	}
}

func (t *transformer) interfaceStmt(x *ast.InterfaceStmt) *ir.InterfaceStmt {
	return &ir.InterfaceStmt{
		Extends:        t.nameList(x.Extends),
		Name:           t.ident(x.Name),
		Stmts:          t.stmtList(x.Stmts),
		NamespacedName: t.name(x.NamespacedName),
	}
}

func (t *transformer) classStmt(x *ast.ClassStmt) *ir.ClassStmt {
	return &ir.ClassStmt{
		Flags:          x.Flags,
		Extends:        t.name(x.Extends),
		Implements:     t.nameList(x.Implements),
		Name:           t.ident(x.Name),
		Stmts:          t.stmtList(x.Stmts),
		NamespacedName: t.name(x.NamespacedName),
	}
}

func (t *transformer) classConstStmt(x *ast.ClassConstStmt) *ir.ClassConstStmt {
	return &ir.ClassConstStmt{
		Flags: x.Flags,
		Type:  t.typeHint(x.Type),
		Name:  t.ident(x.Name),
		Value: t.expr(x.Value),
	}
}

func (t *transformer) propertyStmt(x *ast.PropertyStmt) *ir.PropertyStmt {
	return &ir.PropertyStmt{
		Flags:   x.Flags,
		Type:    t.typeHint(x.Type),
		Name:    t.ident(x.Name),
		Default: t.expr(x.Default),
	}
}

func (t *transformer) classMethodStmt(x *ast.ClassMethodStmt) *ir.ClassMethodStmt {
	return &ir.ClassMethodStmt{
		Flags:      x.Flags,
		ByRef:      x.ByRef,
		Name:       t.ident(x.Name),
		Params:     t.paramList(x.Params),
		ReturnType: t.typeHint(x.ReturnType),
		Stmts:      t.stmtList(x.Stmts),
	}
}

func (t *transformer) traitStmt(x *ast.TraitStmt) *ir.TraitStmt {
	return &ir.TraitStmt{
		Name:           t.ident(x.Name),
		Stmts:          t.stmtList(x.Stmts),
		NamespacedName: t.name(x.NamespacedName),
	}
}

func (t *transformer) traitUseStmt(x *ast.TraitUseStmt) *ir.TraitUseStmt {
	return &ir.TraitUseStmt{
		Traits:      t.nameList(x.Traits),
		Adaptations: slicekit.Map(x.Adaptations, t.traitUseAdaptation),
	}
}

func (t *transformer) traitUseAdaptation(n ast.TraitUseAdaptationStmt) ir.TraitUseAdaptationStmt {
	switch x := n.(type) {
	case *ast.TraitUseAdaptationAliasStmt:
		return t.traitUseAdaptationAliasStmt(x)
	case *ast.TraitUseAdaptationPrecedenceStmt:
		return t.traitUseAdaptationPrecedenceStmt(x)
	}

	// unreachable
	t.unexpectedType(n)
	return nil
}

func (t *transformer) traitUseAdaptationAliasStmt(x *ast.TraitUseAdaptationAliasStmt) *ir.TraitUseAdaptationAliasStmt {
	return &ir.TraitUseAdaptationAliasStmt{
		NewModifier: x.NewModifier,
		NewName:     t.ident(x.NewName),
		Trait:       t.name(x.Trait),
		Method:      t.ident(x.Method),
	}
}

func (t *transformer) traitUseAdaptationPrecedenceStmt(x *ast.TraitUseAdaptationPrecedenceStmt) *ir.TraitUseAdaptationPrecedenceStmt {
	return &ir.TraitUseAdaptationPrecedenceStmt{
		Insteadof: t.nameList(x.Insteadof),
		Trait:     t.name(x.Trait),
		Method:    t.ident(x.Method),
	}
}

// list

func (t *transformer) comments(n []*ast.Comment) []*ir.Comment {
	return slicekit.Map(n, func(c *ast.Comment) *ir.Comment {
		return &ir.Comment{
			Type: c.Type,
			Text: c.Text,
		}
	})
}

func (t *transformer) stringList(n []string) []string {
	return slices.Clone(n)
}

func (t *transformer) argList(n []*ast.Arg) []*ir.Arg {
	return slicekit.Map(n, t.arg)
}
func (t *transformer) nameList(n []*ast.Name) []*ir.Name {
	return slicekit.Map(n, t.name)
}
func (t *transformer) paramList(n []*ast.Param) []*ir.Param {
	return slicekit.Map(n, t.param)
}

func (t *transformer) typeHintList(n []ast.TypeHint) []ir.TypeHint {
	return slicekit.Map(n, t.typeHint)
}

func (t *transformer) exprList(n []ast.Expr) []ir.Expr {
	return slicekit.Map(n, t.expr)
}
func (t *transformer) stmtList(n []ast.Stmt) []ir.Stmt {
	return slicekit.Map(n, t.stmt)
}
