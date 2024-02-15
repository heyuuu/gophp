package php

import (
	"fmt"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/types"
)

const StdClassName = "stdClass"

func ZendThrowOrError(ctx *Context, fetchType uint32, exceptionCe *types.Class, message string) {
	//if (fetchType & ZEND_FETCH_CLASS_EXCEPTION) != 0 {
	ThrowError(ctx, exceptionCe, message)
	//} else {
	//	Error(ctx, perr.E_ERROR, message)
	//}
}

func GetExecutedFilenameEx(ctx *Context) (string, bool) {
	ex := ctx.CurrEX()
	for ex != nil && (ex.Fn() == nil || !ex.Fn().IsUserFunction()) {
		ex = ex.Prev()
	}
	if ex == nil {
		return "", false
	}
	return ex.Fn().Filename(), true
}
func GetExecutedFilename(ctx *Context) string {
	if name, ok := GetExecutedFilenameEx(ctx); ok {
		return name
	}
	return "[no active file]"
}
func GetExecutedFilenameVal(ctx *Context) string {
	name, _ := GetExecutedFilenameEx(ctx)
	return name
}
func GetExecutedLineno(ctx *Context) int {
	return 0
}
func GetExecutedScope(ctx *Context) *types.Class {
	var ex = ctx.CurrEX()
	for {
		if ex == nil {
			return nil
		} else if ex.Fn() != nil && (ex.Fn().IsUserCode() || ex.Fn().Scope() != nil) {
			return ex.Fn().Scope()
		}
		ex = ex.Prev()
	}
}
func IsExecuting(ctx *Context) bool {
	return ctx.CurrEX() != nil
}

func trimClassName(name string) string {
	if name != "" && name[0] == '\\' {
		return name[1:]
	}
	return name
}

func ZendLookupClassEx(ctx *Context, name string, key string, flags uint32) *types.Class {
	if name == "" && key == "" {
		return nil
	}

	var lcName = key
	if lcName == "" {
		lcName = ascii.StrToLower(trimClassName(name))
	}

	if ce := ctx.EG().ClassTable().Get(lcName); ce != nil {
		return ce
	}

	// todo autoload
	return nil
}

func ZendFetchClassByName(ctx *Context, className string, key string, fetchType uint32) *types.Class {
	className = trimClassName(className)
	ce := ZendLookupClassEx(ctx, className, key, fetchType)
	if ce != nil {
		return ce
	}

	if ctx.EG().HasException() {
		return nil
	}
	//if (fetchType & ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_INTERFACE {
	//	ZendThrowOrError(ctx, fetchType, nil, fmt.Sprintf("Interface '%s' not found", className))
	//} else if (fetchType & ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_TRAIT {
	//	ZendThrowOrError(ctx, fetchType, nil, fmt.Sprintf("Trait '%s' not found", className))
	//} else {
	ZendThrowOrError(ctx, fetchType, nil, fmt.Sprintf("Class '%s' not found", className))
	//}
	return nil
}
