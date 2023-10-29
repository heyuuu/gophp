package php

func vmEcho(ctx *Context, zv Val) {
	str := ZvalGetString(zv)
	if len(str) > 0 {
		ctx.WriteString(str)
	}
}
