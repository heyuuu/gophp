package printer

import (
	"fmt"
	"gophp/php/ast"
	"gophp/php/token"
	"os"
	"strconv"
	"strings"
)

type printer struct {
	buf    strings.Builder
	indent int
	err    error
}

func (p *printer) checkError(err error) {
	if err != nil {
		p.err = err
	}
}

func (p *printer) write(data []byte) {
	_, err := p.buf.Write(data)
	p.checkError(err)
}

func (p *printer) writeByte(c byte) {
	err := p.buf.WriteByte(c)
	p.checkError(err)
}

func (p *printer) writeRune(c rune) {
	_, err := p.buf.WriteRune(c)
	p.checkError(err)
}

func (p *printer) writeString(s string) {
	_, err := p.buf.WriteString(s)
	p.checkError(err)
}

func (p *printer) print(args ...any) {
	for _, arg := range args {
		if arg == nil {
			continue
		}

		switch v := arg.(type) {
		case int:
			p.writeString(strconv.Itoa(v))
		case string:
			p.writeString(v)
		case token.Token:
			p.writeString(token.TokenName(v))
		case ast.Node:
			p.printNode(v)
		case []ast.Stmt:
			p.printStmtList(v)
		case []ast.Expr:
			p.printExprList(v)
		default:
			_, _ = fmt.Fprintf(os.Stderr, "print: unsupported argument %v (%T)\n", arg, arg)
			panic("gophp/php/printer type")
		}
	}
}

func (p *printer) printNode(node ast.Node) {
	switch x := node.(type) {
	case *ast.Ident:
		p.writeString(x.Name)
	case *ast.Name:
		p.writeString(x.ToCodeString())
	case ast.Expr:
		p.expr(x)
	case ast.Stmt:
		p.stmt(x)
	case ast.Type:
		p.typeHint(x)
	default:
		err := fmt.Errorf("printer: unsupported node type %T", node)
		p.checkError(err)
	}
}

func (p *printer) result() (string, error) {
	if p.err != nil {
		return "", p.err
	}
	return p.buf.String(), p.err
}

func printList[T ast.Node](p *printer, list []T, sep string) {
	for i, item := range list {
		if i != 0 {
			p.print(sep)
		}
		p.print(item)
	}
}

func (p *printer) printStmtList(stmtList []ast.Stmt) {
	printList(p, stmtList, "\n")
}

func (p *printer) printExprList(exprList []ast.Expr) {
	printList(p, exprList, ", ")
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
