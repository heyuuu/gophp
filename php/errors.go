package php

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

// functions
func Error(ctx *Context, typ perr.ErrorType, message string) {
	filename, lineno := errorGetFilenameLineno(ctx, typ)
	ErrorTrigger(ctx, perr.NewAt(typ, message, filename, lineno))
}
func ErrorNoreturn(ctx *Context, typ perr.ErrorType, message string) {
	Error(ctx, typ, message)
	/* Should never reach this. */
	panic(perr.NewInternal("unreachable"))
}
func ErrorAt(ctx *Context, typ perr.ErrorType, lineno uint32, message string) {
	filename, _ := errorGetFilenameLineno(ctx, typ)
	ErrorTrigger(ctx, perr.NewAt(typ, message, filename, lineno))
}
func ErrorAtNoreturn(ctx *Context, typ perr.ErrorType, lineno uint32, message string) {
	ErrorAt(ctx, typ, lineno, message)
	/* Should never reach this. */
	panic(perr.NewInternal("unreachable"))
}

func ThrowError(ctx *Context, exceptionCe *types.Class, message string) {
	// todo
	Error(ctx, perr.E_ERROR, message)
}

func errorGetFilenameLineno(ctx *Context, typ perr.ErrorType) (string, uint32) {
	// todo
	return "", 0
}

func ErrorTrigger(ctx *Context, err perr.Error) {
	errorVaList(ctx, err.Type, err.Filename, err.Lineno, err.Message)
}

func errorVaList(ctx *Context, typ perr.ErrorType, errorFilename string, errorLineno uint32, message string) {
	ErrorCb(typ, errorFilename, errorLineno, message)
}

func ErrorCb(typ perr.ErrorType, errorFilename string, errorLineno uint32, message string) {
}
