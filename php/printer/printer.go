package printer

import (
	"fmt"
	"gophp/php/ast"
	"os"
	"strconv"
	"strings"
)

type printer struct {
	buf strings.Builder
	err error
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
		case byte:
			p.writeByte(v)
		case int:
			p.writeString(strconv.Itoa(v))
		case string:
			p.writeString(v)
		case *ast.Ident:
			p.writeString(v.Name)
		case *ast.Name:
			p.writeString(v.ToCodeString())
		default:
			_, _ = fmt.Fprintf(os.Stderr, "print: unsupported argument %v (%T)\n", arg, arg)
			panic("gophp/php/printer type")
		}
	}
}

func (p printer) printNode(node ast.Node) (string, error) {
	switch n := node.(type) {
	case ast.Expr:
		p.expr(n)
	case ast.Stmt:
		p.stmt(n)
	case ast.Type:
		p.typeHint(n)
	case []ast.Stmt:
		for _, stmt := range n {
			p.stmt(stmt)
		}
	default:
		err := fmt.Errorf("printer: unsupported node type %T", node)
		p.checkError(err)
	}

	if p.err != nil {
		return "", p.err
	}
	return p.buf.String(), p.err
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
	var p printer
	return p.printNode(node)
}

func (cfg *Config) Sprint(node any) (string, error) {
	return cfg.sprint(node)
}

func Sprint(node any) (string, error) {
	return (&Config{Tabwidth: 8}).Sprint(node)
}
