package php

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"math"
	"path/filepath"
	"strconv"
	"strings"
)

const LongMax = math.MaxInt
const LongMin = math.MinInt

func ForbidDynamicCall(ctx *Context, s string) bool {
	// todo
	return true
}

func ZendDirname(path string) string {
	if path == "" {
		return ""
	}

	/* Strip trailing slashes */
	path = strings.TrimRight(path, "/")
	if path == "" {
		/* The path only contained slashes */
		return string(filepath.Separator)
	}

	/* Strip filename */
	if pos := strings.LastIndexByte(path, '/'); pos >= 0 {
		path = path[:pos]
	} else {
		/* No slash found, therefore return '.' */
		return "."
	}

	/* Strip slashes which came before the file name */
	path = strings.TrimRight(path, "/")
	if path == "" {
		return string(filepath.Separator)
	}
	return path
}

func ConvertToString(ctx *Context, zv *types.Zval) {
	if v, ok := ZvalTryGetStr(ctx, *zv); ok {
		zv.SetString(v)
	}
}

func ZendTryAssignRefLong(ctx *Context, refVal zpp.RefZval, value int) {
	refVal.SetVal(types.ZvalLong(value))
}

func ZendTryAssignRefDouble(ctx *Context, refVal zpp.RefZval, value float64) {
	refVal.SetVal(types.ZvalDouble(value))
}

func NewZvalZval(zv types.Zval, copy bool, dtor bool) types.Zval {
	if zv.IsRef() {
		return zv.DeRef()
	} else if copy && !dtor {
		return zv
	} else {
		return zv
	}
}

func HashOf(p types.Zval) *types.Array {
	if p.IsArray() {
		return p.Array()
		//} else if p.IsObject() {
		//	return p.Object().GetPropertiesArray()
	} else {
		return nil
	}
}

// strtoll(s0, s1, base)
func ZendStrToLong(s string, base int) int {
	value, _ := ZendStrToLongN(s, base)
	return value
}
func ZendStrToLongN(s string, base int) (value int, n int) {
	i, err := strconv.ParseInt(s, base, 64)
	if err != nil {
		return 0, 0
	}
	return int(i), len(s)
}

func ThrowException(ctx *Context, ce *types.Class, message string, code int) {
	if ctx.eh != nil {
		ctx.eh.OnException(ce, message, code)
		return
	}

	panic(perr.Todof("ThrowException: message=%s, code=%d", message, code))
}
