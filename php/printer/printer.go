package printer

import (
	"fmt"
	"github.com/heyuuu/gophp/php/ast"
	"github.com/heyuuu/gophp/php/token"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type printer struct {
	buf     strings.Builder
	indent  int
	err     error
	newLine bool
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
		case string:
			p.write(v)
		case token.Token:
			p.write(token.TokenName(v))
		case ast.Node:
			p.printNode(v)
		// 以下 case 只是为了加快类型匹配
		case []ast.Stmt:
			p.stmtList(v, false)
		case []ast.Expr:
			printList(p, v, ", ")
		case []ast.Node:
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

func (p *printer) printNode(node ast.Node) {
	switch x := node.(type) {
	case *ast.Ident:
		if x.VarLike {
			p.write("$")
		}
		p.write(x.Name)
	case *ast.Name:
		p.write(x.ToCodeString())
	case ast.Expr:
		p.expr(x)
	case ast.Stmt:
		p.stmt(x)
	case ast.Type:
		p.typeHint(x)
	case *ast.Param:
		p.param(x)
	case *ast.Arg:
		p.arg(x)
	default:
		err := fmt.Errorf("printer: unsupported node type %T", node)
		p.checkError(err)
	}
}

func printList[T ast.Node](p *printer, list []T, sep string) {
	for i, item := range list {
		if i != 0 {
			p.print(sep)
		}
		p.print(item)
	}
}

func convertNodeList(data any) ([]ast.Node, bool) {
	if nodes, ok := data.([]ast.Node); ok {
		return nodes, true
	}

	var nodes []ast.Node

	value := reflect.ValueOf(data)
	nodeType := reflect.TypeOf(nodes).Elem()
	if value.Kind() == reflect.Slice && value.Type().Elem().Implements(nodeType) {
		for i := 0; i < value.Len(); i++ {
			nodes = append(nodes, value.Index(i).Interface().(ast.Node))
		}
		return nodes, true
	}
	return nil, false
}

func convertStmtList(data any) ([]ast.Stmt, bool) {
	if nodes, ok := data.([]ast.Stmt); ok {
		return nodes, true
	}

	var nodes []ast.Stmt

	value := reflect.ValueOf(data)
	nodeType := reflect.TypeOf(nodes).Elem()
	if value.Kind() == reflect.Slice && value.Type().Elem().Implements(nodeType) {
		for i := 0; i < value.Len(); i++ {
			nodes = append(nodes, value.Index(i).Interface().(ast.Stmt))
		}
		return nodes, true
	}
	return nil, false
}

func (p *printer) stmtList(stmtList []ast.Stmt, indent bool) {
	if indent {
		p.indent++
	}
	printList(p, stmtList, "\n")
	p.print("\n")
	if indent {
		p.indent--
	}
}

func (p *printer) flags(flags ast.Flags) {
	var names []string
	if flags.Is(ast.FlagPublic) {
		names = append(names, "public")
	}
	if flags.Is(ast.FlagProtected) {
		names = append(names, "protected")
	}
	if flags.Is(ast.FlagPrivate) {
		names = append(names, "private")
	}
	if flags.Is(ast.FlagStatic) {
		names = append(names, "static")
	}
	if flags.Is(ast.FlagAbstract) {
		names = append(names, "abstract")
	}
	if flags.Is(ast.FlagFinal) {
		names = append(names, "final")
	}
	if flags.Is(ast.FlagReadonly) {
		names = append(names, "readonly")
	}
	p.print(strings.Join(names, " "))
}

// ----------------------------------------------------------------------------
// Public interface

// A Mode value is a set of flags (or 0). They control printing.
type Mode uint

const (
	RawFormat Mode = 1 << iota // do not use a tabwriter; if set, UseSpaces is ignored
	TabIndent                  // use tabs for indentation independent of UseSpaces
	UseSpaces                  // use spaces instead of tabs for alignment
	SourcePos                  // emit //line directives to preserve original source positions
)

// A Config node controls the output of Fprint.
type Config struct {
	Mode     Mode // default: 0
	Tabwidth int  // default: 8
	Indent   int  // default: 0 (all code is indented at least by this much)
}

func (cfg *Config) sprint(node any) (string, error) {
	var p = &printer{}
	// todo 需要验证 node 为 print 可以打印的类型范围
	p.print(node)
	return p.result()
}

func (cfg *Config) Sprint(node any) (string, error) {
	return cfg.sprint(node)
}

func (cfg *Config) SprintFile(node any) (string, error) {
	code, err := cfg.sprint(node)
	if err != nil {
		return "", err
	}
	return "<?php\n" + code, nil
}

func Sprint(node any) (string, error) {
	return (&Config{Tabwidth: 8}).Sprint(node)
}

func SprintFile(node any) (string, error) {
	return (&Config{Tabwidth: 8}).SprintFile(node)
}
