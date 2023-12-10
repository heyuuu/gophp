package php

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

/**
 * constants and global variables
 */
const E_ERROR = 1 << 0
const E_WARNING = 1 << 1
const E_PARSE = 1 << 2
const E_NOTICE = 1 << 3
const E_CORE_ERROR = 1 << 4
const E_CORE_WARNING = 1 << 5
const E_COMPILE_ERROR = 1 << 6
const E_COMPILE_WARNING = 1 << 7
const E_USER_ERROR = 1 << 8
const E_USER_WARNING = 1 << 9
const E_USER_NOTICE = 1 << 10
const E_STRICT = 1 << 11
const E_RECOVERABLE_ERROR = 1 << 12
const E_DEPRECATED = 1 << 13
const E_USER_DEPRECATED = 1 << 14
const E_ALL = E_ERROR | E_WARNING | E_PARSE | E_NOTICE | E_CORE_ERROR | E_CORE_WARNING | E_COMPILE_ERROR | E_COMPILE_WARNING | E_USER_ERROR | E_USER_WARNING | E_USER_NOTICE | E_RECOVERABLE_ERROR | E_DEPRECATED | E_USER_DEPRECATED | E_STRICT
const E_CORE = E_CORE_ERROR | E_CORE_WARNING

type PhpError struct {
	typ      int
	message  string
	filename string
	lineno   uint32
}

func (e PhpError) Error() string { return e.message }

func NewError(typ int, message string, filename string, lineno uint32) PhpError {
	return PhpError{typ: typ, message: message, filename: filename, lineno: lineno}
}

// functions
func Error(ctx *Context, typ int, message string) {
	filename, lineno := errorGetFilenameLineno(ctx, typ)
	ErrorTrigger(ctx, NewError(typ, message, filename, lineno))
}
func ErrorNoreturn(ctx *Context, typ int, message string) {
	Error(ctx, typ, message)
	/* Should never reach this. */
	panic(perr.New("unreachable"))
}
func ErrorAt(ctx *Context, typ int, lineno uint32, message string) {
	filename, _ := errorGetFilenameLineno(ctx, typ)
	ErrorTrigger(ctx, NewError(typ, message, filename, lineno))
}
func ErrorAtNoreturn(ctx *Context, typ int, lineno uint32, message string) {
	ErrorAt(ctx, typ, lineno, message)
	/* Should never reach this. */
	panic(perr.New("unreachable"))
}

func ThrowError(exceptionCe *types.Class, message string) {
	// todo
	Error(E_ERROR, message)
}

func errorGetFilenameLineno(ctx *Context, typ int) (string, uint32) {
	// todo
	return "", 0
}

func ErrorTrigger(ctx *Context, err PhpError) {
	errorVaList(ctx, err.typ, err.filename, err.lineno, err.message)
}

func errorVaList(ctx *Context, typ int, errorFilename string, errorLineno uint32, message string) {
	ErrorCb(typ, errorFilename, errorLineno, message)
}

func ErrorCb(typ int, errorFilename string, errorLineno uint32, message string) {
}
