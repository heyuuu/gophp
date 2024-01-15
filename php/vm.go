package php

func vmEcho(ctx *Context, zv Val) {
	str := ZvalGetStrVal(ctx, zv)
	if len(str) > 0 {
		ctx.WriteString(str)
	}
}
