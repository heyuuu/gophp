package php

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"path/filepath"
	"strings"
)

const PrintZvalIndent = 4

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
	} else if p.IsObject() {
		return p.Object().GetPropertiesArray()
	} else {
		return nil
	}
}

func ThrowException(ctx *Context, ce *types.Class, message string, code int) {
	if ctx.eh != nil {
		ctx.eh.OnException(ce, message, code)
		return
	}

	panic(perr.Todof("ThrowException: message=%s, code=%d", message, code))
}

func ManglePropertyName(src1 string, src2 string) string {
	return "\000" + src1 + "\000" + src2
}

func UnmanglePropertyName(ctx *Context, name string) (className string, propName string, ok bool) {
	if len(name) == 0 || name[0] != '\000' {
		return "", name, true
	}
	if len(name) < 3 || name[1] == '\000' {
		Error(ctx, perr.E_NOTICE, "Illegal member variable name")
		return "", name, false
	}
	/*
	 * 可能的Name结构
	 * -	\0 + {className} + \0 + {$propName}
	 * -	\0 + {className} + \0 + {annoClassSrc} + \0 + {$propName}
	 */
	parts := strings.SplitN(name[1:], "\000", 3)
	switch len(parts) {
	case 2:
		return parts[0], parts[1], true
	case 3:
		return parts[0], parts[2], true
	default:
		Error(ctx, perr.E_NOTICE, "Corrupt member variable name")
		return "", name, false
	}
}

func ZendGetPropertiesFor(v types.Zval, debug types.PropPurposeType) *types.Array {
	return nil
}

func ZendRsrcListGetRsrcTypeEx(res *types.Resource) *string {
	return nil
}
