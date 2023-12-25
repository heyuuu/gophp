package standard

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
)

func ZifVarDump(ctx zpp.Ctx, ex zpp.Ex, value *types.Zval) {
	// todo
	ctx.WriteString("has call var_dump()")
}
