package standard

import (
	"github.com/heyuuu/gophp/ext/standard/vari"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
)

func ZifVarDump(ctx *php.Context, vars []types.Zval) {
	p := vari.NewVarDumpPrinter(ctx)
	for _, zv := range vars {
		p.Zval(zv, 1)
	}
}

func ZifDebugZvalDump(ctx *php.Context, vars []types.Zval) {
	p := vari.NewVarDebugPrinter(ctx)
	for _, zval := range vars {
		p.Zval(zval, 1)
	}
}

func ZifVarExport(ctx *php.Context, value types.Zval, _ zpp.Opt, return_ bool) types.Zval {
	p := vari.NewVarExportPrinter(ctx)
	p.Zval(value, 1)
	if return_ {
		return php.String(p.String())
	} else {
		ctx.WriteString(p.String())
		return types.Null
	}
}
