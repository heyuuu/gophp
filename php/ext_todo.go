package php

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"math"
	"path/filepath"
	"strings"
)

const LongMax = math.MaxInt

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

func ZendTryAssignRefDouble(ctx *Context, percent zpp.RefZval, float float64) {

}

func ZendTryAssignRefLong(ctx *Context, count zpp.RefZval, count2 int) {

}
