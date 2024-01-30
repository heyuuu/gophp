package printer

import (
	"fmt"
	"io"
)

type basePrinter struct {
	w   io.Writer
	err error
}

func (p *basePrinter) Error() error {
	return p.err
}

func (p *basePrinter) print(str string) {
	if p.err != nil {
		return
	}

	_, p.err = io.WriteString(p.w, str)
}

func (p *basePrinter) printf(format string, a ...any) {
	if p.err != nil {
		return
	}
	_, p.err = fmt.Fprintf(p.w, format, a...)
}

func (p *basePrinter) println(str string) {
	p.print(str)
	p.print("\n")
}

func (p *basePrinter) printIdent(ident int) {
	if ident <= 0 {
		return
	}

	p.printf("%*c", ident, ' ')
}
