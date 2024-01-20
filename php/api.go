package php

import "github.com/heyuuu/gophp/php/types"

func Print(ctx *Context, v types.Zval) {
	str := ZvalGetStrVal(ctx, v)
	if len(str) > 0 {
		ctx.WriteString(str)
	}
}
