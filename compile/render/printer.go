package render

import (
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/ir"
	"github.com/heyuuu/gophp/kits/mapkit"
	"github.com/heyuuu/gophp/kits/slicekit"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Config
type Config struct{}

/**
 * private
 */
const (
	PkgExecutor = "github.com/heyuuu/gophp/php/executor"
	PkgTypes    = "github.com/heyuuu/gophp/php/types"
)

// printer
type printer struct {
	nameResolver NameResolver
	buf          strings.Builder
	indent       int
	err          error
	newLine      bool
	imports      map[string]bool
}

func newPrinter(config *Config) *printer {
	return &printer{
		nameResolver: newDefaultNameResolver(),
		imports:      map[string]bool{},
	}
}
func defaultPrinter() *printer {
	return newPrinter(nil)
}

func (p *printer) printFile(ns *ir.Namespace) (string, error) {
	p.reset()

	p.pNamespace(ns)
	body, err := p.result()
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	buf.WriteString("/* namespace " + ns.Name + " */\n")
	buf.WriteString("package ir\n\n")

	imports := mapkit.SortedKeys(p.imports)
	if len(imports) > 0 {
		buf.WriteString("import (\n")
		for _, importName := range imports {
			buf.WriteString("\n" + importName + "\"\n")
		}
		buf.WriteString(")\n")
	}
	buf.WriteString(body)

	return buf.String(), nil
}

func (p *printer) reset() {
	p.buf.Reset()
	p.err = nil
	p.newLine = false
	p.imports = make(map[string]bool)
}

func (p *printer) addImport(pkgName string) {
	p.imports[pkgName] = true
}

func (p *printer) pNamespace(ns *ir.Namespace) {
	// ns init func
	nsInitFuncName := "nsInit" + p.nameResolver.Namespace(ns.Name)
	p.write(fmt.Sprintf("func %s(ex executor.Executor) {\n", nsInitFuncName))

	p.write("}\n")

	//
	slicekit.Each(ns.Segments, p.pSegment)
	p.print("\n")
}

func (p *printer) pSegment(seg ir.Segment) {
	switch x := seg.(type) {
	case *ir.InitFunc:
		p.print("// init\n")
		if len(x.Stmts) == 0 {
			p.print("func init() {}")
		} else {
			p.print("func init() {\n")
			p.stmtList(x.Stmts, true)
			p.print("}\n")
		}
	case *ir.Func:
		p.print("// ", x.Name, "\n")
		p.print("func ")
		if x.ByRef {
			p.print("&")
		}
		p.print(x.Name, "(", x.Params, ")")
		if x.ReturnType != nil {
			p.print(" ", x.ReturnType)
		}
		if len(x.Stmts) == 0 {
			p.print(" {}\n")
		} else {
			p.print(" {\n")
			p.stmtList(x.Stmts, true)
			p.print("}\n")
		}
	case *ir.Class:
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
	case *ir.Interface:
		p.print("// interface ", x.Name, "\n")
		p.print("interface ", x.Name)
		if len(x.Extends) != 0 {
			p.print(" extends ", x.Extends)
		}
		p.print("\n{\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ir.Trait:
		p.print("// trait ", x.Name, "\n")
		p.print("trait ", x.Name, "\n{\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	default:
		err := fmt.Errorf("printer: unsupported segment type %T", x)
		p.checkError(err)
	}
	p.print("\n")
}

func (p *printer) PrintFile(f *ir.File) (string, error) {
	p.reset()
	for _, ns := range f.Namespaces {
		p.pNamespace(ns)
	}
	return p.result()
}

// misc
func (p *printer) checkError(err error) {
	if err != nil {
		p.err = err
	}
}

func (p *printer) result() (string, error) {
	if p.err != nil {
		return "", p.err
	}
	return p.buf.String(), p.err
}

func (p *printer) write(s string) {
	if s == "" {
		return
	}

	indentStr := strings.Repeat("    ", p.indent)
	if p.newLine {
		p.buf.WriteString(indentStr)
		p.newLine = false
	}

	l := len(s)
	if s[l-1] != '\n' {
		p.buf.WriteString(strings.ReplaceAll(s, "\n", "\n"+indentStr))
	} else {
		p.buf.WriteString(strings.ReplaceAll(s[:l-1], "\n", "\n"+indentStr))
		p.buf.WriteByte('\n')
		p.newLine = true
	}
}

func (p *printer) print(args ...any) {
	for _, arg := range args {
		if arg == nil {
			continue
		}

		switch v := arg.(type) {
		case int:
			p.write(strconv.Itoa(v))
		case byte:
			p.write(string(v))
		case rune:
			p.write(string(v))
		case string:
			p.write(v)
		case ir.Node:
			p.printNode(v)
		case fmt.Stringer:
			p.print(v.String())
		// 以下 case 只是为了加快类型匹配
		case []ir.Stmt:
			p.stmtList(v, false)
		case []ir.Expr:
			printList(p, v, ", ")
		case []ir.Node:
			printList(p, v, ", ")
		default:
			if stmts, ok := convertStmtList(arg); ok {
				p.stmtList(stmts, false)
			} else if nodes, ok := convertNodeList(arg); ok {
				printList(p, nodes, ", ")
			} else {
				_, _ = fmt.Fprintf(os.Stderr, "print: unsupported argument %v (%T)\n", arg, arg)
				panic("gophp/php/printer type")
			}
		}
	}
}

func (p *printer) printNode(node ir.Node) {
	if isNil(node) {
		p.write("nil")
		return
	}

	switch x := node.(type) {
	case ir.Ident:
		p.write(string(x))
	case *ir.Ident:
		p.write(string(*x))
	case ir.Name:
		p.write(x.ToCodeString())
	case *ir.Name:
		p.write(x.ToCodeString())
	case ir.Expr:
		p.expr(x)
	case ir.Stmt:
		p.stmt(x)
	case ir.Type:
		p.typeHint(x)
	case *ir.Param:
		p.param(x)
	case *ir.Arg:
		p.arg(x)
	default:
		err := fmt.Errorf("printer: unsupported node type %T", node)
		p.checkError(err)
	}
}

func printList[T ir.Node](p *printer, list []T, sep string) {
	for i, item := range list {
		if i != 0 {
			p.print(sep)
		}
		p.print(item)
	}
}

func convertNodeList(data any) ([]ir.Node, bool) {
	if nodes, ok := data.([]ir.Node); ok {
		return nodes, true
	}

	var nodes []ir.Node

	value := reflect.ValueOf(data)
	nodeType := reflect.TypeOf(nodes).Elem()
	if value.Kind() == reflect.Slice && value.Type().Elem().Implements(nodeType) {
		for i := 0; i < value.Len(); i++ {
			nodes = append(nodes, value.Index(i).Interface().(ir.Node))
		}
		return nodes, true
	}
	return nil, false
}

func convertStmtList(data any) ([]ir.Stmt, bool) {
	if nodes, ok := data.([]ir.Stmt); ok {
		return nodes, true
	}

	var nodes []ir.Stmt

	value := reflect.ValueOf(data)
	nodeType := reflect.TypeOf(nodes).Elem()
	if value.Kind() == reflect.Slice && value.Type().Elem().Implements(nodeType) {
		for i := 0; i < value.Len(); i++ {
			nodes = append(nodes, value.Index(i).Interface().(ir.Stmt))
		}
		return nodes, true
	}
	return nil, false
}

func (p *printer) stmtList(stmtList []ir.Stmt, indent bool) {
	if indent {
		p.indent++
	}
	printList(p, stmtList, "\n")
	p.print("\n")
	if indent {
		p.indent--
	}
}

func (p *printer) flags(flags ir.Flags) {
	var names []string
	if flags.Is(ir.FlagPublic) {
		names = append(names, "public")
	}
	if flags.Is(ir.FlagProtected) {
		names = append(names, "protected")
	}
	if flags.Is(ir.FlagPrivate) {
		names = append(names, "private")
	}
	if flags.Is(ir.FlagStatic) {
		names = append(names, "static")
	}
	if flags.Is(ir.FlagAbstract) {
		names = append(names, "abstract")
	}
	if flags.Is(ir.FlagFinal) {
		names = append(names, "final")
	}
	if flags.Is(ir.FlagReadonly) {
		names = append(names, "readonly")
	}
	p.print(strings.Join(names, " "))
}

/**
 * nodes
 */
func (p *printer) arg(n *ir.Arg) {
	if n.Unpack {
		p.print("...")
	}
	p.print(n.Value)
}

func (p *printer) param(n *ir.Param) {
	p.print(n.Name, " ")
	if n.Variadic {
		p.print("...")
	}
	if n.ByRef {
		p.print("*")
	}
	if n.Type != nil {
		p.print(n.Type)
	} else {
		p.print("any") // todo Zval
	}
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
		p.print(x.Name)
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
		p.print(x.Kind, x.Expr)
	case *ir.UnaryExpr:
		switch x.Op {
		case ast.UnaryOpPostInc, ast.UnaryOpPostDec:
			p.print(x.Var, x.Op)
		default:
			p.print(x.Op, x.Var)
		}
	case *ir.BinaryOpExpr:
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
		if len(x.Init) == 0 && len(x.Loop) == 0 {
			p.print("for ", x.Cond, " {\n")
		} else {
			p.print("for ", x.Init, ";", x.Cond, ";", x.Loop, " {\n")
		}
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
		p.print("while ", x.Cond, " {\n", x.Stmts, "}")
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
		p.print("const ", x.Name, " = ", x.Value)
	case *ir.EchoStmt:
		p.print("echo ", x.Exprs, ";")
	case *ir.GlobalStmt:
		p.print("global ", x.Vars, ";")
	case *ir.HaltCompilerStmt:
		p.print("__halt_compiler();", x.Remaining)
	case *ir.InlineHTMLStmt:
		p.print("?>", x.Value, "<?php")
	case *ir.StaticStmt:
		p.print("static ", x.Name)
		if x.Default != nil {
			p.print(" = ", x.Default)
		}
	case *ir.UnsetStmt:
		p.print("unset(", x.Vars, ")")
	case *ir.UseStmt:
		var useType string
		switch x.Type {
		case ir.UseFunc:
			useType = "func "
		case ir.UseConst:
			useType = "const "
		}

		if x.Alias != "" {
			p.print("// use ", useType, x.Name, " as ", x.Alias, ";")
		} else {
			p.print("// use ", useType, x.Name, ";")
		}
	case *ir.DeclStmt:
		p.pSegment(x.Decl)
	case *ir.AnonymousClassStmt:
		if x.Flags != 0 {
			p.flags(x.Flags)
			p.print(" ")
		}
		p.print("class ")
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
		if x.Flags != 0 {
			p.flags(x.Flags)
			p.print(" ")
		}
		p.print("const ", x.Name, " = ", x.Value)
	case *ir.PropertyStmt:
		if x.Flags != 0 {
			p.flags(x.Flags)
		} else {
			p.print("var")
		}
		p.print(" ", x.Name)
		if x.Type != nil {
			p.print(" ", x.Type)
		}
		if x.Default != nil {
			p.print(" = ", x.Default)
		}
	case *ir.MethodStmt:
		if x.Flags != 0 {
			p.flags(x.Flags)
			p.print(" ")
		}
		p.print("func ")
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

func isNil(n any) bool {
	if n == nil {
		return true
	}

	v := reflect.ValueOf(n)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	}
	return false
}
