package render

import (
	"fmt"
	"github.com/heyuuu/gophp/compile/ir"
	"github.com/heyuuu/gophp/php/def"
	"log"
	"runtime"
	"strconv"
	"strings"
)

func RenderFile(f *ir.File) (string, error) {
	r := &render{}
	return r.RenderFile(f)
}

const (
	// packages
	pkgDef pkgName = "github.com/heyuuu/gophp/php/def"
	pkgAst pkgName = `github.com/heyuuu/gophp/compile/ast`

	// variables
	varFile   = "f"
	varDef    = "d"
	varSwitch = "switchVal"
)

type stringLit string

type render struct {
	buf       strings.Builder
	imports   *imports
	indent    int
	indentStr string
	isNewLine bool
}

var _ ir.Visitor = (*render)(nil)

func (r *render) RenderFile(f *ir.File) (result string, resultErr error) {
	defer func() {
		if e := recover(); e != nil {
			if err, ok := e.(error); ok {
				resultErr = fmt.Errorf("render.render.RenderFile() error: %w", err)
			} else {
				resultErr = fmt.Errorf("render.render.RenderFile() error: %v", e)
			}

			// 打印堆栈
			const size = 64 << 10
			stack := make([]byte, size)
			stack = stack[:runtime.Stack(stack, false)]
			log.Printf(">>> tests.Command.RunBuiltin() panic: %v\n%s", e, stack)
		}
	}()

	// reset
	r.buf.Reset()
	r.imports = newImports()
	r.indent = 0

	// render
	r.file(f)

	// header (header 必须执行在 file 后，因为要收集 imports 信息)
	header := r.fileHeader()

	return header + r.buf.String(), nil
}

func (r *render) unexpected(kind string, v any) {
	err := fmt.Errorf("unexpected %s type %T", kind, v)
	panic(err)
}

func (r *render) todo(msg string) {
	panic("todo: " + msg)
}

func (r *render) incIndent() {
	r.indent++
	r.indentStr = strings.Repeat("\t", r.indent)
}

func (r *render) decIndent() {
	r.indent--
	r.indentStr = strings.Repeat("\t", r.indent)
}

func (r *render) indentBlock(f func()) {
	r.incIndent()
	f()
	r.decIndent()
}

func (r *render) string(s string) {
	for s != "" {
		if r.isNewLine {
			r.buf.WriteString(r.indentStr)
			r.isNewLine = false
		}
		if idx := strings.IndexByte(s, '\n'); idx < 0 {
			r.buf.WriteString(s)
			break
		} else {
			r.buf.WriteString(s[:idx+1])
			r.isNewLine = true
			s = s[idx+1:]
		}
	}
}

func (r *render) print(args ...any) {
	for _, arg := range args {
		switch v := arg.(type) {
		// lit
		case nil:
			r.print("nil")
		case bool:
			if v {
				r.string("true")
			} else {
				r.string("false")
			}
		case int:
			r.string(strconv.Itoa(v))
		case string:
			r.string(v)
		case pkgName:
			r.string(r.usePkg(v))
		case stringLit:
			r.string(`"` + string(v) + `"`) // todo escape
		// node
		case ir.Node:
			r.node(v)
		case []ir.Stmt:
			r.stmtList(v, true)
		case []ir.Expr:
			r.exprList(v)
		case func():
			v()
		default:
			r.unexpected("print", arg)
		}
	}
}

func (r *render) printLine(args ...any) {
	r.print(args...)
	r.print("\n")
}

func (r *render) printf(s string, args ...any) {
	if len(args) != 0 {
		s = fmt.Sprintf(s, args...)
	}
	r.string(s)
}

func (r *render) linef(format string, args ...any) {
	r.printf(format, args...)
	r.printf("\n")
}

// imports

func (r *render) usePkg(pkg pkgName) string {
	return r.imports.getOrAdd(pkg)
}

func (r *render) fileHeader() string {
	var buf strings.Builder

	buf.WriteString("package main\n\n")

	if r.imports.Len() > 0 {
		buf.WriteString("import (\n")
		r.imports.SortedEach(func(pkg pkgName, alias string) {
			if alias == pkg.DefaultAlias() {
				_, _ = fmt.Fprintf(&buf, "\t\"%s\"\n", pkg)
			} else {
				_, _ = fmt.Fprintf(&buf, "\t%s \"%s\"\n", pkg, alias)
			}
		})
		buf.WriteString(")\n\n")
	}

	return buf.String()
}

// file
func (r *render) file(f *ir.File) {
	// check empty
	if len(f.Namespaces) == 0 {
		return
	}

	r.print("func init() {\n")
	r.indentBlock(func() {
		r.print(varFile, ` := `)
		r.pPkgCall(pkgDef, "NewFile", stringLit(f.FilePath), f.StrictTypes)
		for _, namespace := range f.Namespaces {
			r.namespace(namespace)
		}
	})
	r.print("}\n")
}

func (r *render) namespace(n *ir.NamespaceStmt) {
	r.printf("\n")

	var nsName string
	if n.Name != nil {
		nsName = n.Name.ToString()
	}

	r.print(varFile, ".TopFn(", stringLit(nsName), ", func(", varDef, " ", pkgDef, ".TopDefiner) ", pkgDef, ".Val {\n")
	r.stmtList(n.Stmts, true)
	r.print("\treturn nil\n")
	r.print("})\n")
}

// nodes

func (r *render) node(n ir.Node) {
	if n == nil {
		r.print("nil")
		return
	}

	err := ir.Visit(r, n)
	if ir.IsUnexpectedError(err) {
		r.unexpected("node", n)
	}
}

func (r *render) typeHint(n ir.TypeHint) {
	ir.VisitTypeHint(r, n)
}

func (r *render) stmtComments(n ir.Stmt) {
	for _, comment := range n.Comments() {
		r.print(comment.Text, "\n")
	}
}

func (r *render) stmt(n ir.Stmt) {
	r.stmtComments(n)
	err := ir.VisitStmt(r, n)
	if ir.IsUnexpectedError(err) {
		r.unexpected("expr", n)
	}
}

func (r *render) stmtList(stmts []ir.Stmt, indent bool) {
	if indent {
		r.incIndent()
	}
	for _, stmt := range stmts {
		r.stmt(stmt)
		r.printf("\n")
	}
	if indent {
		r.decIndent()
	}
}

func (r *render) expr(n ir.Expr) {
	if n == nil {
		r.print("nil")
		return
	}

	err := ir.VisitExpr(r, n)
	if ir.IsUnexpectedError(err) {
		r.unexpected("expr", n)
	}
}

func (r *render) exprList(exprs []ir.Expr) {
	pList(r, exprs, ", ")
}

// misc helpers
func (r *render) pArgs(args ...any) {
	r.print("(")
	pList(r, args, ", ")
	r.print(")")
}

func (r *render) pPkgCall(pkg pkgName, method string, args ...any) {
	r.print(pkg, ".", method)
	r.pArgs(args...)
}

func (r *render) pVarCall(variable string, method string, args ...any) {
	r.print(variable, ".", method)
	r.pArgs(args...)
}

func (r *render) pDefCall(method string, args ...any) {
	r.print(varDef, ".", method)
	r.pArgs(args...)
}

func (r *render) pAsBool(arg any) {
	r.pDefCall("IsTrue", arg)
}

func (r *render) pAsStr(arg any, stage def.Stage) {
	r.pDefCall("AsStr", arg, func() {
		r.print(pkgDef, ".", stage.String())
	})
}

func (r *render) pIdentOrExprAsStr(n ir.Node, stage def.Stage, expectedType string) {
	switch x := n.(type) {
	case *ir.Ident:
		r.print(stringLit(x.Name))
	case ir.Expr:
		r.pAsStr(x, stage)
	default:
		r.unexpected(expectedType, n)
	}
}

func (r *render) pNameOrExprAsStr(n ir.Node, stage def.Stage, expectedType string) {
	switch x := n.(type) {
	case *ir.Name:
		r.print(stringLit(x.ToCodeString()))
	case ir.Expr:
		r.pAsStr(x, stage)
	default:
		r.unexpected(expectedType, n)
	}
}

func (r *render) pClassName(n ir.Node) {
	r.pNameOrExprAsStr(n, def.StageClassName, "className")
}

func pList[T any](r *render, nodes []T, sep string) {
	for i, node := range nodes {
		if i != 0 {
			r.print(sep)
		}
		r.print(node)
	}
}
