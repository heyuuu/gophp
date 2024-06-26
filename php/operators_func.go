package php

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/kits/cmpkit"
	"github.com/heyuuu/gophp/php/fix"
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

// internal functions
func sign(i int) int {
	if i > 0 {
		return 1
	}
	return 0
}

func fastGetDouble(v types.Zval) float64 {
	if v.IsLong() {
		return float64(v.Long())
	} else if v.IsDouble() {
		return v.Double()
	} else {
		return 0
	}
}

func IsNumericString(str string) bool {
	return ParseNumber(str).IsNotUndef()
}

// @see 替代 FastEqualCheckLong / FastEqualCheckString / FastEqualCheckFunction
func FastEqualFunction(ctx *Context, v1, v2 types.Zval) bool {
	if v, ok := ZvalEqualsEx(ctx, v1, v2); ok {
		return v
	} else {
		return false
	}
}

func StringCompareFunction(ctx *Context, op1 types.Zval, op2 types.Zval) int {
	var str1 = ZvalGetStrVal(ctx, op1)
	var str2 = ZvalGetStrVal(ctx, op2)
	return strings.Compare(str1, str2)
}
func StringCaseCompareFunction(ctx *Context, op1 types.Zval, op2 types.Zval) int {
	var str1 = ZvalGetStrVal(ctx, op1)
	var str2 = ZvalGetStrVal(ctx, op2)
	return ascii.StrCaseCompare(str1, str2)
}
func StringLocaleCompareFunction(ctx *Context, op1 types.Zval, op2 types.Zval) int {
	var str1 = ZvalGetStrVal(ctx, op1)
	var str2 = ZvalGetStrVal(ctx, op2)
	return fix.StrColl(str1, str2)
}
func NumericCompareFunction(ctx *Context, op1 types.Zval, op2 types.Zval) int {
	var d1 float64
	var d2 float64
	d1 = ZvalGetDouble(ctx, op1)
	d2 = ZvalGetDouble(ctx, op2)
	return cmpkit.Normalize(d1 - d2)
}
