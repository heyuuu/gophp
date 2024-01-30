package printer

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

func VarDump(ctx *php.Context, vars ...types.Zval) {
	p := NewVarDumpPrinter(ctx, ctx.AsWriter())
	for _, zv := range vars {
		p.Zval(zv, 0)
	}
}

func VarDebugDump(ctx *php.Context, vars ...types.Zval) {
	p := NewVarDebugPrinter(ctx, ctx.AsWriter())
	for _, zval := range vars {
		p.Zval(zval, 0)
	}
}

func VarExport(ctx *php.Context, v types.Zval) string {
	var buf strings.Builder
	p := NewVarExportPrinter(ctx, &buf)
	p.Zval(v, 0)
	return buf.String()
}

func PrintR(ctx *php.Context, v types.Zval) string {
	var buf strings.Builder
	p := NewPrintRPrinter(ctx, &buf)
	p.Zval(v, 0)
	return buf.String()
}

func PrintFlatR(ctx *php.Context, expr types.Zval) {
	p := NewPrintFlatRPrinter(ctx, ctx.AsWriter())
	p.Zval(expr)
}
