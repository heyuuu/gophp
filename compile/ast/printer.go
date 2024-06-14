package ast

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/**
 *	public
 */
func PrintFile(file *File) (string, error) {
	p := &printer{inCodeScope: false}
	p.file(file)
	return p.result()
}

func PrintNode(node Node) (string, error) {
	p := &printer{inCodeScope: true}
	p.node(node)
	return p.result()
}

/**
 *	private
 */
type printer struct {
	buf            strings.Builder
	indent         int
	err            error
	newLine        bool
	inCodeScope    bool
	multiNamespace bool
}

func (p *printer) enterCodeScope() {
	if !p.inCodeScope {
		p.inCodeScope = true
		p.print("<?php\n")
	}
}

func (p *printer) leaveCodeScope() {
	if p.inCodeScope {
		p.inCodeScope = false
		p.print("\n?>")
	}
}

func (p *printer) checkError(err error) {
	if err != nil {
		p.err = err
	}
}

func (p *printer) result() (string, error) {
	if p.err != nil {
		return "", p.err
	}
	return p.buf.String(), nil
}

func (p *printer) string(s string) {
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

func (p *printer) file(f *File) {
	for _, declare := range f.Declares {
		p.stmt(declare)
		p.print("\n")
	}

	p.multiNamespace = len(f.Namespaces) > 1
	for _, namespace := range f.Namespaces {
		p.namespace(namespace)
		p.print("\n")
	}
}

func (p *printer) namespace(x *NamespaceStmt) {
	if p.multiNamespace {
		if x.Name != nil {
			p.print("namespace ", x.Name, " {\n")
		} else {
			p.print("namespace {\n")
		}
		p.stmtList(x.Stmts, true)
		p.print("}\n")
	} else {
		if x.Name != nil {
			p.print("namespace ", x.Name, ";\n\n")
		}
		p.stmtList(x.Stmts, false)
	}
}

func (p *printer) unexpect(n any) {
	err := fmt.Errorf("printer: unexpected node type %T", n)
	p.checkError(err)
}

func (p *printer) print(args ...any) {
	for _, arg := range args {
		if arg == nil {
			continue
		}

		switch v := arg.(type) {
		case string:
			p.string(v)
		case Node:
			p.node(v)
		case fmt.Stringer:
			p.string(v.String())
		// 以下 case 只是为了加快类型匹配
		case []Stmt:
			p.stmtList(v, true)
		case []Expr:
			printList(p, v, ", ")
		case []Node:
			printList(p, v, ", ")
		default:
			if stmts, ok := convertStmtList(arg); ok {
				p.stmtList(stmts, true)
			} else if nodes, ok := convertNodeList(arg); ok {
				printList(p, nodes, ", ")
			} else {
				p.unexpect(v)
			}
		}
	}
}

func (p *printer) node(node Node) {
	if node == nil || reflect.ValueOf(node).IsNil() {
		p.string("nil")
		return
	}

	switch x := node.(type) {
	case *Ident:
		if x.VarLike {
			p.string("$")
		}
		p.string(x.Name)
	case *Name:
		p.string(x.ToCodeString())
	case Expr:
		p.expr(x)
	case Stmt:
		p.stmt(x)
	case TypeHint:
		p.typeHint(x)
	case *Param:
		p.param(x)
	case *Arg:
		p.arg(x)
	default:
		p.unexpect(x)
	}
}

func printList[T Node](p *printer, list []T, sep string) {
	for i, item := range list {
		if i != 0 {
			p.string(sep)
		}
		p.node(item)
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

func (p *printer) stmtList(stmts []Stmt, indent bool) {
	if indent {
		p.indent++
	}
	for _, stmt := range stmts {
		p.stmt(stmt)
		p.string("\n")
	}
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

func (p *printer) typeHint(n TypeHint) {
	p.typeHint0(n, false)
}

func (p *printer) typeHint0(n TypeHint, wrap bool) {
	switch t := n.(type) {
	case *SimpleType:
		p.print(t.Name)
	case *NullableType:
		p.print("?")
		p.typeHint0(t.Type, true)
	case *IntersectionType:
		if wrap {
			p.print("(")
		}
		for i, typ := range t.Types {
			if i != 0 {
				p.print("&")
			}
			p.typeHint0(typ, true)
		}
		if wrap {
			p.print(")")
		}
	case *UnionType:
		if wrap {
			p.print("(")
		}
		for i, typ := range t.Types {
			if i != 0 {
				p.print("|")
			}
			p.typeHint0(typ, true)
		}
		if wrap {
			p.print(")")
		}
	default:
		p.unexpect(t)
	}
}

func (p *printer) expr(n Expr) {
	switch x := n.(type) {
	case *IntLit:
		if rawValue := MetaRawValue(x); rawValue != "" {
			p.print(rawValue)
			break
		}

		p.print(strconv.Itoa(x.Value))
	case *FloatLit:
		if rawValue := MetaRawValue(x); rawValue != "" {
			p.print(rawValue)
			break
		}

		p.print(fmt.Sprintf("%f", x.Value))
	case *StringLit:
		if rawValue := MetaRawValue(x); rawValue != "" {
			p.print(rawValue)
			break
		}

		// todo escape
		p.print(`"`, x.Value, `"`)
	case *ArrayExpr:
		p.print("[")
		printList(p, x.Items, ", ")
		p.print("]")
	case *ArrayItemExpr:
		if x.Key != nil {
			p.print(x.Key, " => ")
		}
		if x.ByRef {
			p.print("&")
		}
		if x.Unpack {
			p.print("...")
		}
		p.print(x.Value)
	case *ClosureExpr:
		if x.Static {
			p.print("static ")
		}
		p.print("function ")
		if x.ByRef {
			p.print("&")
		}
		p.print("(", x.Params, ")")
		if len(x.Uses) > 0 {
			p.print(" use (", x.Uses, ")")
		}
		if x.ReturnType != nil {
			p.print(" : ", x.ReturnType)
		}
		p.print(" {\n", x.Stmts, "}")
	case *ClosureUseExpr:
		if x.ByRef {
			p.print("&")
		}
		p.print(x.Var)
	case *ArrowFunctionExpr:
		if x.Static {
			p.print("static ")
		}
		p.print("fn")
		if x.ByRef {
			p.print("&")
		}
		p.print("(", x.Params, ")")
		if x.ReturnType != nil {
			p.print(": ", x.ReturnType)
		}
		p.print(" => ", x.Expr)
	case *IndexExpr:
		p.print(x.Var, "[", x.Dim, "]")
	case *CastExpr:
		p.print(x.Kind, x.Expr)
	case *UnaryExpr:
		switch x.Op {
		case UnaryOpPostInc, UnaryOpPostDec:
			p.print(x.Var, x.Op)
		default:
			p.print(x.Op, x.Var)
		}
	case *BinaryOpExpr:
		p.print(x.Left, " ", x.Op, " ", x.Right)
	case *AssignExpr:
		p.print(x.Var, " = ", x.Expr)
	case *AssignOpExpr:
		p.print(x.Var, " ", x.Op, " ", x.Expr)
	case *AssignRefExpr:
		p.print(x.Var, " = &", x.Expr)
	case *IssetExpr:
		p.print("isset(", x.Vars, ")")
	case *EmptyExpr:
		p.print("empty(", x.Expr, ")")
	case *EvalExpr:
		p.print("eval(", x.Expr, ")")
	case *IncludeExpr:
		p.print(x.Kind, " ", x.Expr)
	case *CloneExpr:
		p.print("clone ", x.Expr)
	case *ErrorSuppressExpr:
		p.print("@", x.Expr)
	case *ExitExpr:
		p.print("exit(", x.Expr, ")")
	case *ConstFetchExpr:
		p.print(x.Name)
	case *ClassConstFetchExpr:
		p.print(x.Class, "::", x.Name)
	case *MagicConstExpr:
		p.print(x.Kind)
	case *InstanceofExpr:
		p.print(x.Expr, " instanceof ", x.Class)
	case *ListExpr:
		p.print("list(", x.Items, ")")
	case *PrintExpr:
		p.print("print ", x.Expr)
	case *PropertyFetchExpr:
		if !x.Nullsafe {
			p.print(x.Var, "->", x.Name)
		} else {
			p.print(x.Var, "?->", x.Name)
		}
	case *StaticPropertyFetchExpr:
		p.print(x.Class, "::", x.Name)
	case *ShellExecExpr:
		p.print("`")
		for _, part := range x.Parts {
			if lit, ok := part.(*StringLit); ok {
				// todo escape
				p.print(lit.Value)
			} else {
				p.print("{", part, "}")
			}
		}
		p.print("`")
	case *TernaryExpr:
		if x.If == nil {
			p.print(x.Cond, " ?: ", x.Else)
		} else {
			p.print(x.Cond, " ? ", x.If, " : ", x.Else)
		}
	case *ThrowExpr:
		p.print("throw ", x.Expr)
	case *VariableExpr:
		if nameIdent, ok := x.Name.(*Ident); ok {
			p.print("$", nameIdent.Name)
		} else {
			p.print("${", x.Name, "}")
		}
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
		if !x.Nullsafe {
			p.print(x.Var, "->", x.Name, "(", x.Args, ")")
		} else {
			p.print(x.Var, "?->", x.Name, "(", x.Args, ")")
		}
	case *StaticCallExpr:
		p.print(x.Class, "::", x.Name, "(", x.Args, ")")
	default:
		p.unexpect(n)
	}
}

func (p *printer) stmt(n Stmt) {
	if html, ok := n.(*InlineHTMLStmt); ok {
		p.leaveCodeScope()
		p.print(html.Value)
		return
	}

	p.enterCodeScope()
	switch x := n.(type) {
	case *EmptyStmt:
		p.print(";")
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
		p.print("if (", x.Cond, ") {\n", x.Stmts, "}")
		for _, elseif := range x.Elseifs {
			p.print(" elseif (", elseif.Cond, ") {\n", elseif.Stmts, "}")
		}
		if x.Else != nil {
			p.print(" else {\n", x.Else.Stmts, "}")
		}
	case *SwitchStmt:
		p.print("switch (", x.Cond, ") {\n")
		p.indent++
		for _, caseStmt := range x.Cases {
			if caseStmt.Cond != nil {
				p.print("case ", caseStmt.Cond, ":\n")
			} else {
				p.print("default:\n")
			}
			p.stmtList(caseStmt.Stmts, true)
		}
		p.indent--
		p.print("}")
	case *ForStmt:
		p.print("for (", x.Init, "; ", x.Cond, "; ", x.Loop, ") {\n", x.Stmts, "}")
	case *ForeachStmt:
		p.print("foreach (", x.Expr, " as ")
		if x.KeyVar != nil {
			p.print(x.KeyVar, " => ")
		}
		if x.ByRef {
			p.print("&")
		}
		p.print(x.ValueVar, ") {\n", x.Stmts, "}")
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
		p.print("while (", x.Cond, ") {\n", x.Stmts, "}")
	case *DoStmt:
		p.print("do {\n", x.Stmts, "} while (", x.Cond, ");")
	case *TryCatchStmt:
		p.print("try {\n", x.Stmts, "}")
		for _, catch := range x.Catches {
			p.print(" catch (")
			printList(p, catch.Types, "|")
			if catch.Var != nil {
				p.print(" ", catch.Var)
			}
			p.print(") {\n", catch.Stmts, "}")
		}
		if x.Finally != nil {
			p.print(" finally {\n", x.Finally.Stmts, "}")
		}
	case *ConstStmt:
		p.print("const ", x.Name, " = ", x.Value, ";")
	case *EchoStmt:
		p.print("echo ", x.Exprs, ";")
	case *GlobalStmt:
		p.print("global ", x.Vars, ";")
	case *HaltCompilerStmt:
		p.print("__halt_compiler();", x.Remaining)
	case *StaticStmt:
		p.print("static ", x.Vars, ";")
	case *StaticVarStmt:
		if x.Default != nil {
			p.print(x.Var, " = ", x.Default)
		} else {
			p.print(x.Var)
		}
	case *UnsetStmt:
		p.print("unset(", x.Vars, ");")
	case *UseStmt:
		var useType string
		switch x.Type {
		case UseFunction:
			useType = "function "
		case UseConstant:
			useType = "const "
		}

		if x.Alias != nil {
			p.print("use ", useType, x.Name, " as ", x.Alias, ";")
		} else {
			p.print("use ", useType, x.Name, ";")
		}
	case *DeclareStmt:
		p.print("declare(", x.Declares, ")")
		if len(x.Stmts) == 0 {
			p.print(";")
		} else {
			p.print(" {\n", x.Stmts, "}")
		}
	case *DeclareDeclareStmt:
		p.print(x.Key, "=", x.Value)
	case *FunctionStmt:
		p.print("function ")
		if x.ByRef {
			p.print("&")
		}
		p.print(x.Name, "(", x.Params, ")")
		if x.ReturnType != nil {
			p.print(": ", x.ReturnType)
		}
		p.print(" {\n", x.Stmts, "}")
	case *InterfaceStmt:
		p.print("interface ", x.Name)
		if len(x.Extends) != 0 {
			p.print(" extends ", x.Extends)
		}
		p.print("\n{\n", x.Stmts, "}")
	case *ClassStmt:
		if x.Flags != 0 {
			p.flags(x.Flags)
			p.print(" ")
		}
		p.print("class")
		if x.Name != nil {
			p.print(" ", x.Name)
		}
		if x.Extends != nil {
			p.print(" extends ", x.Extends)
		}
		if len(x.Implements) > 0 {
			p.print(" implements ", x.Implements)
		}
		p.print("\n{\n", x.Stmts, "}")
	case *ClassConstStmt:
		if x.Flags != 0 {
			p.flags(x.Flags)
			p.print(" ")
		}
		p.print("const ")
		if x.Type != nil {
			p.print(x.Type, " ")
		}
		p.print(x.Name, " = ", x.Value, ";")
	case *PropertyStmt:
		if x.Flags != 0 {
			p.flags(x.Flags)
			p.print(" ")
		}
		if x.Type != nil {
			p.print(x.Type, " ")
		}
		if x.Default != nil {
			p.print(x.Name, " = ", x.Default)
		} else {
			p.print(x.Name)
		}
		p.print(";")
	case *ClassMethodStmt:
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
		p.print("\n{\n", x.Stmts, "}")
	case *TraitStmt:
		p.print("trait ", x.Name, "\n{\n", x.Stmts, "}")
	case *TraitUseStmt:
		if len(x.Adaptations) > 0 {
			p.print("use ", x.Traits, " {\n", x.Adaptations, "}")
		} else {
			p.print("use ", x.Traits, ";")
		}
	case *TraitUseAdaptationPrecedenceStmt:
		if x.Trait != nil {
			p.print(x.Trait, "::")
		}
		p.print(x.Method, " insteadof ", x.Insteadof, ";")
	case *TraitUseAdaptationAliasStmt:
		if x.Trait != nil {
			p.print(x.Trait, "::")
		}
		p.print(x.Method, " as")
		if x.NewModifier != 0 {
			p.print(" ")
			p.flags(x.NewModifier)
		}
		if x.NewName != nil {
			p.print(" ", x.NewName)
		}
		p.print(";")
	default:
		p.unexpect(n)
	}
}
