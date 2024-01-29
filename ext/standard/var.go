package standard

import (
	"github.com/heyuuu/gophp/ext/standard/vari"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
)

func ZifVarDump(ctx *php.Context, vars []types.Zval) {
	p := vari.NewVarDumpPrinter(ctx)
	for _, zv := range vars {
		p.Zval(zv, 1)
	}
}
