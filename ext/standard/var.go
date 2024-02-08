package standard

import (
	"github.com/heyuuu/gophp/ext/standard/printer"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
)

func ZifVarDump(ctx *php.Context, vars []types.Zval) {
	printer.VarDump(ctx, vars...)
}

func ZifDebugZvalDump(ctx *php.Context, vars []types.Zval) {
	printer.VarDebugDump(ctx, vars...)
}

func ZifVarExport(ctx *php.Context, value types.Zval, _ zpp.Opt, return_ bool) types.Zval {
	s := printer.VarExport(ctx, value)
	if return_ {
		return php.String(s)
	} else {
		ctx.WriteString(s)
		return types.Null
	}
}
func ZifSerialize(ctx *php.Context, var_ types.Zval) types.Zval {
	serializer := InitSerializer(ctx)
	serializer.Serialize(var_)
	serializer.DestroyData()

	if ctx.EG().HasException() {
		return types.False
	}

	serializerStr := serializer.String()
	if serializerStr != "" {
		return types.ZvalString(serializerStr)
	} else {
		return types.Null
	}
}
