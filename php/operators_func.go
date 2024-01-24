package php

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/kits/cmpkit"
	"github.com/heyuuu/gophp/php/fix"
	"strings"
)

// @see 替代 FastEqualCheckLong / FastEqualCheckString / FastEqualCheckFunction
func FastEqualFunction(ctx *Context, v1, v2 Val) bool {
	if v, ok := ZvalEqualsEx(ctx, v1, v2); ok {
		return v
	} else {
		return false
	}
}

func StringCompareFunction(ctx *Context, op1 Val, op2 Val) int {
	var str1 = ZvalGetStrVal(ctx, op1)
	var str2 = ZvalGetStrVal(ctx, op2)
	return strings.Compare(str1, str2)
}
func StringCaseCompareFunction(ctx *Context, op1 Val, op2 Val) int {
	var str1 = ZvalGetStrVal(ctx, op1)
	var str2 = ZvalGetStrVal(ctx, op2)
	return ascii.StrCaseCompare(str1, str2)
}
func StringLocaleCompareFunction(ctx *Context, op1 Val, op2 Val) int {
	var str1 = ZvalGetStrVal(ctx, op1)
	var str2 = ZvalGetStrVal(ctx, op2)
	return fix.StrColl(str1, str2)
}
func NumericCompareFunction(ctx *Context, op1 Val, op2 Val) int {
	var d1 float64
	var d2 float64
	d1 = ZvalGetDouble(ctx, op1)
	d2 = ZvalGetDouble(ctx, op2)
	return cmpkit.Normalize(d1 - d2)
}
