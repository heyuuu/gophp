package ir

import (
	"fmt"
	"github.com/heyuuu/gophp/compile/token"
	"github.com/heyuuu/gophp/kits/slicekit"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/**
 *	public
 */
func PrintProject(proj *Project) (map[string]string, error) {
	return defaultPrinter().PrintProject(proj)
}

func PrintFile(file *File) (string, error) {
	return defaultPrinter().PrintFile(file)
}

type Config struct {
	TabWidth int // default: 8
	Indent   int // default: 0 (all code is indented at least by this much)
}

/**
 *	private
 */
type printer struct {
	buf     strings.Builder
	indent  int
	err     error
	newLine bool
}

func newPrinter(config *Config) *printer {
	return &printer{}
}
func defaultPrinter() *printer {
	return newPrinter(nil)
}

func (p *printer) reset() {
	p.buf.Reset()
	p.err = nil
	p.newLine = false
}

func (p *printer) PrintProject(proj *Project) (map[string]string, error) {
	result := make(map[string]string, len(proj.namespaces))
	for _, namespace := range proj.namespaces {
		content, err := p.PrintNamespace(namespace)
		if err != nil {
			return nil, err
		}

		result[namespace.Name] = content
	}
	return result, nil
}

func (p *printer) PrintNamespace(ns *Namespace) (string, error) {
	p.reset()
	p.pNamespace(ns)
	return p.result()
}

func (p *printer) pNamespace(ns *Namespace) {
	nsName := ns.Name
	if nsName == "" {
		nsName = "_"
	}
	p.print("/**\n * namespace " + nsName + "\n */\n")
	p.print("package ir\n\n")
	slicekit.Each(ns.Segments, p.pSegment)
	p.print("\n")
}

func (p *printer) pSegment(seg Segment) {
	switch x := seg.(type) {
	case *InitFunc:
		p.print("// init\n")
		if len(x.Stmts) == 0 {
			p.print("func init() {}")
		} else {
			p.print("func init() {\n")
			p.stmtList(x.Stmts, true)
			p.print("}\n")
		}
	case *Func:
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
	case *Class:
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
	case *Interface:
		p.print("// interface ", x.Name, "\n")
		p.print("interface ", x.Name)
		if len(x.Extends) != 0 {
			p.print(" extends ", x.Extends)
		}
		p.print("\n{\n")
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *Trait:
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

func (p *printer) PrintFile(f *File) (string, error) {
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
		case token.Token:
			p.write(token.TokenName(v))
		case Node:
			p.printNode(v)
		// 以下 case 只是为了加快类型匹配
		case []Stmt:
			p.stmtList(v, false)
		case []Expr:
			printList(p, v, ", ")
		case []Node:
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

func (p *printer) printNode(node Node) {
	if isNil(node) {
		p.write("nil")
		return
	}

	switch x := node.(type) {
	case Ident:
		p.write(string(x))
	case *Ident:
		p.write(string(*x))
	case Name:
		p.write(x.ToCodeString())
	case *Name:
		p.write(x.ToCodeString())
	case Expr:
		p.expr(x)
	case Stmt:
		p.stmt(x)
	case Type:
		p.typeHint(x)
	case *Param:
		p.param(x)
	case *Arg:
		p.arg(x)
	default:
		err := fmt.Errorf("printer: unsupported node type %T", node)
		p.checkError(err)
	}
}

func printList[T Node](p *printer, list []T, sep string) {
	for i, item := range list {
		if i != 0 {
			p.print(sep)
		}
		p.print(item)
	}
}

func convertNodeList(data any) ([]Node, bool) {
	if nodes, ok := data.([]Node); ok {
		return nodes, true
	}

	var nodes []Node

	value := reflect.ValueOf(data)
	nodeType := reflect.TypeOf(nodes).Elem()
	if value.Kind() == reflect.Slice && value.Type().Elem().Implements(nodeType) {
		for i := 0; i < value.Len(); i++ {
			nodes = append(nodes, value.Index(i).Interface().(Node))
		}
		return nodes, true
	}
	return nil, false
}

func convertStmtList(data any) ([]Stmt, bool) {
	if nodes, ok := data.([]Stmt); ok {
		return nodes, true
	}

	var nodes []Stmt

	value := reflect.ValueOf(data)
	nodeType := reflect.TypeOf(nodes).Elem()
	if value.Kind() == reflect.Slice && value.Type().Elem().Implements(nodeType) {
		for i := 0; i < value.Len(); i++ {
			nodes = append(nodes, value.Index(i).Interface().(Stmt))
		}
		return nodes, true
	}
	return nil, false
}

func (p *printer) stmtList(stmtList []Stmt, indent bool) {
	if indent {
		p.indent++
	}
	printList(p, stmtList, "\n")
	p.print("\n")
	if indent {
		p.indent--
	}
}

func (p *printer) flags(flags Flags) {
	var names []string
	if flags.Is(FlagPublic) {
		names = append(names, "public")
	}
	if flags.Is(FlagProtected) {
		names = append(names, "protected")
	}
	if flags.Is(FlagPrivate) {
		names = append(names, "private")
	}
	if flags.Is(FlagStatic) {
		names = append(names, "static")
	}
	if flags.Is(FlagAbstract) {
		names = append(names, "abstract")
	}
	if flags.Is(FlagFinal) {
		names = append(names, "final")
	}
	if flags.Is(FlagReadonly) {
		names = append(names, "readonly")
	}
	p.print(strings.Join(names, " "))
}

/**
 * nodes
 */
func (p *printer) arg(n *Arg) {
	if n.Unpack {
		p.print("...")
	}
	p.print(n.Value)
}

func (p *printer) param(n *Param) {
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
func (p *printer) typeHint(n Type) {
	p.typeHint0(n, false)
}

func (p *printer) typeHint0(n Type, wrap bool) {
	switch t := n.(type) {
	case *SimpleType:
		p.print(t.Name)
	case *NullableType:
		p.print('?')
		p.typeHint0(t.Type, true)
	case *IntersectionType:
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
	case *UnionType:
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

func (p *printer) expr(n Expr) {
	switch x := n.(type) {
	case *IntLit:
		p.print(x.Value)
	case *FloatLit:
		p.print(fmt.Printf("%f", x.Value))
	case *StringLit:
		// todo escape
		p.print("\"", x.Value, "\"")
	case *ArrayExpr:
		p.print("[")
		printList(p, x.Items, ", ")
		p.print("]")
	case *ArrayItemExpr:
		if x.Key != nil {
			p.print(x.Key, " => ")
		}
		p.print(x.Value)
	case *ClosureExpr:
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
	case *ClosureUseExpr:
		if x.ByRef {
			p.print("&")
		}
		p.print(x.Name)
	case *ArrowFunctionExpr:
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
	case *IndexExpr:
		p.print(x.Var, "[", x.Dim, "]")
	case *CastExpr:
		p.print(x.Op, x.Expr)
	case *UnaryExpr:
		switch x.Kind {
		case token.PostInc, token.PostDec:
			p.print(x.Var, x.Kind)
		default:
			p.print(x.Kind, x.Var)
		}
	case *BinaryExpr:
		p.print(x.Left, " ", x.Op, " ", x.Right)
	case *AssignExpr:
		p.print(x.Var, " ", x.Op, " ", x.Expr)
	case *AssignRefExpr:
		p.print(x.Var, " = &", x.Expr)
	case *InternalCallExpr:
		p.print(x.Kind)
	case *CloneExpr:
		p.print("clone ", x.Expr)
	case *ErrorSuppressExpr:
		p.print("@", x.Expr)
	case *ExitExpr:
		p.print("exist(", x.Expr, ")")
	case *ConstFetchExpr:
		p.print(x.Name)
	case *ClassConstFetchExpr:
		p.print(x.Class, "::", x.Name)
	case *MagicConstExpr:
		p.print(x.Kind)
	case *InstanceofExpr:
		p.print(x.Expr, " instanceOf ", x.Class)
	case *ListExpr:
		p.print("list(")
		printList(p, x.Items, ", ")
		p.print(")")
	case *PrintExpr:
		p.print("print ", x.Expr)
	case *PropertyFetchExpr:
		p.print(x.Var, "->", x.Name)
	case *StaticPropertyFetchExpr:
		p.print(x.Class, "::", x.Name)
	case *ShellExecExpr:
		p.print('`')
		printList(p, x.Parts, "")
		p.print('`')
	case *TernaryExpr:
		if x.If == nil {
			p.print(x.Cond, " ?: ", x.Else)
		} else {
			p.print(x.Cond, " ? ", x.If, " : ", x.Else)
		}
	case *ThrowExpr:
		p.print("throw ", x.Expr)
	case *VariableExpr:
		p.print("$", x.Name)
	case *YieldExpr:
		if x.Key == nil {
			p.print("yield ", x.Value)
		} else {
			p.print("yield ", x.Key, " => ", x.Value)
		}
	case *YieldFromExpr:
		p.print("yield from ", x.Expr)
	case *FuncCallExpr:
		p.print(x.Name, "(", x.Args, ")")
	case *NewExpr:
		p.print("new ", x.Class, "(", x.Args, ")")
	case *MethodCallExpr:
		p.print(x.Var, "->", x.Name, "(", x.Args, ")")
	case *StaticCallExpr:
		p.print(x.Class, "::", x.Name, "(", x.Args, ")")
	default:
		panic("unreachable")
	}
}

func (p *printer) stmt(n Stmt) {
	switch x := n.(type) {
	case *EmptyStmt:
		//p.print(";")
	case *BlockStmt:
		p.stmtList(x.List, false)
	case *ExprStmt:
		p.print(x.Expr, ";")
	case *ReturnStmt:
		if x.Expr == nil {
			p.print("return;")
		} else {
			p.print("return ", x.Expr, ";")
		}
	case *LabelStmt:
		p.print(x.Name, ":")
	case *GotoStmt:
		p.print("goto ", x.Name, ";")
	case *IfStmt:
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
	case *SwitchStmt:
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
	case *ForStmt:
		if len(x.Init) == 0 && len(x.Loop) == 0 {
			p.print("for ", x.Cond, " {\n")
		} else {
			p.print("for ", x.Init, ";", x.Cond, ";", x.Loop, " {\n")
		}
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *ForeachStmt:
		if x.KeyVar != nil {
			p.print("foreach (", x.KeyVar, " as ", x.KeyVar, " => ", x.ValueVar, ") {\n")
		} else {
			p.print("foreach (", x.KeyVar, " as ", x.ValueVar, ") {\n")
		}
		p.stmtList(x.Stmts, true)
		p.print("}")
	case *BreakStmt:
		if x.Num != nil {
			p.print("break ", x.Num, ";")
		} else {
			p.print("break;")
		}
	case *ContinueStmt:
		if x.Num != nil {
			p.print("continue ", x.Num, ";")
		} else {
			p.print("continue;")
		}
	case *WhileStmt:
		p.print("while ", x.Cond, " {\n", x.Stmts, "}")
	case *DoStmt:
		p.print("do {\n", x.Stmts, "} while (", x.Cond, ");")
	case *TryCatchStmt:
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
	case *ConstStmt:
		p.print("const ", x.Name, " = ", x.Value)
	case *EchoStmt:
		p.print("echo ", x.Exprs, ";")
	case *GlobalStmt:
		p.print("global ", x.Vars, ";")
	case *HaltCompilerStmt:
		p.print("__halt_compiler();", x.Remaining)
	case *InlineHTMLStmt:
		p.print("?>", x.Value, "<?php")
	case *StaticStmt:
		p.print("static ", x.Name)
		if x.Default != nil {
			p.print(" = ", x.Default)
		}
	case *UnsetStmt:
		p.print("unset(", x.Vars, ")")
	case *UseStmt:
		var useType string
		switch x.Type {
		case UseFunc:
			useType = "func "
		case UseConst:
			useType = "const "
		}

		if x.Alias != "" {
			p.print("use ", useType, x.Name, " as ", x.Alias, ";")
		} else {
			p.print("use ", useType, x.Name, ";")
		}
	case *DeclStmt:
		p.pSegment(x.Decl)
	case *AnonymousClassStmt:
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
	case *ClassConstStmt:
		if x.Flags != 0 {
			p.flags(x.Flags)
			p.print(" ")
		}
		p.print("const ", x.Name, " = ", x.Value)
	case *PropertyStmt:
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
	case *MethodStmt:
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
	case *TraitUseStmt:
		if len(x.Adaptations) != 0 {
			p.print("use ", x.Traits, " {\n")
			p.indent++
			p.print(x.Adaptations)
			p.indent--
			p.print("};")
		} else {
			p.print("use ", x.Traits, ";")
		}
	case *TraitUseAdaptationPrecedenceStmt:
		p.print(x.Trait, "::", x.Method, " insteadof ", x.Insteadof, ";")
	case *TraitUseAdaptationAliasStmt:
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
